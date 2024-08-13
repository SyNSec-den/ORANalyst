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
// description      Defines Schema, the main entry to every subSchema.
//                  Contains the parsing logic and error checking.
//
// created          26-02-2013

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:27
package gojsonschema

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:27
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:27
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:27
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:27
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:27
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:27
)

import (
	"errors"
	"math/big"
	"reflect"
	"regexp"
	"text/template"

	"github.com/xeipuuv/gojsonreference"
)

var (
	// Locale is the default locale to use
	// Library users can overwrite with their own implementation
	Locale	locale	= DefaultLocale{}

	// ErrorTemplateFuncs allows you to define custom template funcs for use in localization.
	ErrorTemplateFuncs	template.FuncMap
)

// NewSchema instances a schema using the given JSONLoader
func NewSchema(l JSONLoader) (*Schema, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:49
	_go_fuzz_dep_.CoverTab[195085]++
											return NewSchemaLoader().Compile(l)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:50
	// _ = "end of CoverTab[195085]"
}

// Schema holds a schema
type Schema struct {
	documentReference	gojsonreference.JsonReference
	rootSchema		*subSchema
	pool			*schemaPool
	referencePool		*schemaReferencePool
}

func (d *Schema) parse(document interface{}, draft Draft) error {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:61
	_go_fuzz_dep_.CoverTab[195086]++
											d.rootSchema = &subSchema{property: STRING_ROOT_SCHEMA_PROPERTY, draft: &draft}
											return d.parseSchema(document, d.rootSchema)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:63
	// _ = "end of CoverTab[195086]"
}

// SetRootSchemaName sets the root-schema name
func (d *Schema) SetRootSchemaName(name string) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:67
	_go_fuzz_dep_.CoverTab[195087]++
											d.rootSchema.property = name
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:68
	// _ = "end of CoverTab[195087]"
}

// Parses a subSchema
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:71
//
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:71
// Pretty long function ( sorry :) )... but pretty straight forward, repetitive and boring
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:71
// Not much magic involved here, most of the job is to validate the key names and their values,
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:71
// then the values are copied into subSchema struct
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:77
func (d *Schema) parseSchema(documentNode interface{}, currentSchema *subSchema) error {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:77
	_go_fuzz_dep_.CoverTab[195088]++

											if currentSchema.draft == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:79
		_go_fuzz_dep_.CoverTab[195137]++
												if currentSchema.parent == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:80
			_go_fuzz_dep_.CoverTab[195139]++
													return errors.New("Draft not set")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:81
			// _ = "end of CoverTab[195139]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:82
			_go_fuzz_dep_.CoverTab[195140]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:82
			// _ = "end of CoverTab[195140]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:82
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:82
		// _ = "end of CoverTab[195137]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:82
		_go_fuzz_dep_.CoverTab[195138]++
												currentSchema.draft = currentSchema.parent.draft
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:83
		// _ = "end of CoverTab[195138]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:84
		_go_fuzz_dep_.CoverTab[195141]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:84
		// _ = "end of CoverTab[195141]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:84
	// _ = "end of CoverTab[195088]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:84
	_go_fuzz_dep_.CoverTab[195089]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:87
	if *currentSchema.draft >= Draft6 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:87
		_go_fuzz_dep_.CoverTab[195142]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:87
		return isKind(documentNode, reflect.Bool)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:87
		// _ = "end of CoverTab[195142]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:87
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:87
		_go_fuzz_dep_.CoverTab[195143]++
												b := documentNode.(bool)
												currentSchema.pass = &b
												return nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:90
		// _ = "end of CoverTab[195143]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:91
		_go_fuzz_dep_.CoverTab[195144]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:91
		// _ = "end of CoverTab[195144]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:91
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:91
	// _ = "end of CoverTab[195089]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:91
	_go_fuzz_dep_.CoverTab[195090]++

											if !isKind(documentNode, reflect.Map) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:93
		_go_fuzz_dep_.CoverTab[195145]++
												return errors.New(formatErrorDescription(
			Locale.ParseError(),
			ErrorDetails{
				"expected": STRING_SCHEMA,
			},
		))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:99
		// _ = "end of CoverTab[195145]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:100
		_go_fuzz_dep_.CoverTab[195146]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:100
		// _ = "end of CoverTab[195146]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:100
	// _ = "end of CoverTab[195090]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:100
	_go_fuzz_dep_.CoverTab[195091]++

											m := documentNode.(map[string]interface{})

											if currentSchema.parent == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:104
		_go_fuzz_dep_.CoverTab[195147]++
												currentSchema.ref = &d.documentReference
												currentSchema.id = &d.documentReference
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:106
		// _ = "end of CoverTab[195147]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:107
		_go_fuzz_dep_.CoverTab[195148]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:107
		// _ = "end of CoverTab[195148]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:107
	// _ = "end of CoverTab[195091]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:107
	_go_fuzz_dep_.CoverTab[195092]++

											if currentSchema.id == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:109
		_go_fuzz_dep_.CoverTab[195149]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:109
		return currentSchema.parent != nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:109
		// _ = "end of CoverTab[195149]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:109
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:109
		_go_fuzz_dep_.CoverTab[195150]++
												currentSchema.id = currentSchema.parent.id
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:110
		// _ = "end of CoverTab[195150]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:111
		_go_fuzz_dep_.CoverTab[195151]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:111
		// _ = "end of CoverTab[195151]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:111
	// _ = "end of CoverTab[195092]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:111
	_go_fuzz_dep_.CoverTab[195093]++

	// In draft 6 the id keyword was renamed to $id
	// Hybrid mode uses the old id by default
	var keyID string

	switch *currentSchema.draft {
	case Draft4:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:118
		_go_fuzz_dep_.CoverTab[195152]++
												keyID = KEY_ID
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:119
		// _ = "end of CoverTab[195152]"
	case Hybrid:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:120
		_go_fuzz_dep_.CoverTab[195153]++
												keyID = KEY_ID_NEW
												if existsMapKey(m, KEY_ID) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:122
			_go_fuzz_dep_.CoverTab[195155]++
													keyID = KEY_ID
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:123
			// _ = "end of CoverTab[195155]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:124
			_go_fuzz_dep_.CoverTab[195156]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:124
			// _ = "end of CoverTab[195156]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:124
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:124
		// _ = "end of CoverTab[195153]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:125
		_go_fuzz_dep_.CoverTab[195154]++
												keyID = KEY_ID_NEW
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:126
		// _ = "end of CoverTab[195154]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:127
	// _ = "end of CoverTab[195093]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:127
	_go_fuzz_dep_.CoverTab[195094]++
											if existsMapKey(m, keyID) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:128
		_go_fuzz_dep_.CoverTab[195157]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:128
		return !isKind(m[keyID], reflect.String)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:128
		// _ = "end of CoverTab[195157]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:128
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:128
		_go_fuzz_dep_.CoverTab[195158]++
												return errors.New(formatErrorDescription(
			Locale.InvalidType(),
			ErrorDetails{
				"expected":	TYPE_STRING,
				"given":	keyID,
			},
		))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:135
		// _ = "end of CoverTab[195158]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:136
		_go_fuzz_dep_.CoverTab[195159]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:136
		// _ = "end of CoverTab[195159]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:136
	// _ = "end of CoverTab[195094]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:136
	_go_fuzz_dep_.CoverTab[195095]++
											if k, ok := m[keyID].(string); ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:137
		_go_fuzz_dep_.CoverTab[195160]++
												jsonReference, err := gojsonreference.NewJsonReference(k)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:139
			_go_fuzz_dep_.CoverTab[195162]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:140
			// _ = "end of CoverTab[195162]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:141
			_go_fuzz_dep_.CoverTab[195163]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:141
			// _ = "end of CoverTab[195163]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:141
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:141
		// _ = "end of CoverTab[195160]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:141
		_go_fuzz_dep_.CoverTab[195161]++
												if currentSchema == d.rootSchema {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:142
			_go_fuzz_dep_.CoverTab[195164]++
													currentSchema.id = &jsonReference
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:143
			// _ = "end of CoverTab[195164]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:144
			_go_fuzz_dep_.CoverTab[195165]++
													ref, err := currentSchema.parent.id.Inherits(jsonReference)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:146
				_go_fuzz_dep_.CoverTab[195167]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:147
				// _ = "end of CoverTab[195167]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:148
				_go_fuzz_dep_.CoverTab[195168]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:148
				// _ = "end of CoverTab[195168]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:148
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:148
			// _ = "end of CoverTab[195165]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:148
			_go_fuzz_dep_.CoverTab[195166]++
													currentSchema.id = ref
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:149
			// _ = "end of CoverTab[195166]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:150
		// _ = "end of CoverTab[195161]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:151
		_go_fuzz_dep_.CoverTab[195169]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:151
		// _ = "end of CoverTab[195169]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:151
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:151
	// _ = "end of CoverTab[195095]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:151
	_go_fuzz_dep_.CoverTab[195096]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:154
	if existsMapKey(m, KEY_DEFINITIONS) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:154
		_go_fuzz_dep_.CoverTab[195170]++
												if isKind(m[KEY_DEFINITIONS], reflect.Map, reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:155
			_go_fuzz_dep_.CoverTab[195171]++
													for _, dv := range m[KEY_DEFINITIONS].(map[string]interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:156
				_go_fuzz_dep_.CoverTab[195172]++
														if isKind(dv, reflect.Map, reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:157
					_go_fuzz_dep_.CoverTab[195173]++

															newSchema := &subSchema{property: KEY_DEFINITIONS, parent: currentSchema}

															err := d.parseSchema(dv, newSchema)

															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:163
						_go_fuzz_dep_.CoverTab[195174]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:164
						// _ = "end of CoverTab[195174]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:165
						_go_fuzz_dep_.CoverTab[195175]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:165
						// _ = "end of CoverTab[195175]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:165
					}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:165
					// _ = "end of CoverTab[195173]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:166
					_go_fuzz_dep_.CoverTab[195176]++
															return errors.New(formatErrorDescription(
						Locale.InvalidType(),
						ErrorDetails{
							"expected":	STRING_ARRAY_OF_SCHEMAS,
							"given":	KEY_DEFINITIONS,
						},
					))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:173
					// _ = "end of CoverTab[195176]"
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:174
				// _ = "end of CoverTab[195172]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:175
			// _ = "end of CoverTab[195171]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:176
			_go_fuzz_dep_.CoverTab[195177]++
													return errors.New(formatErrorDescription(
				Locale.InvalidType(),
				ErrorDetails{
					"expected":	STRING_ARRAY_OF_SCHEMAS,
					"given":	KEY_DEFINITIONS,
				},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:183
			// _ = "end of CoverTab[195177]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:184
		// _ = "end of CoverTab[195170]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:186
		_go_fuzz_dep_.CoverTab[195178]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:186
		// _ = "end of CoverTab[195178]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:186
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:186
	// _ = "end of CoverTab[195096]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:186
	_go_fuzz_dep_.CoverTab[195097]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:189
	if existsMapKey(m, KEY_TITLE) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:189
		_go_fuzz_dep_.CoverTab[195179]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:189
		return !isKind(m[KEY_TITLE], reflect.String)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:189
		// _ = "end of CoverTab[195179]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:189
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:189
		_go_fuzz_dep_.CoverTab[195180]++
												return errors.New(formatErrorDescription(
			Locale.InvalidType(),
			ErrorDetails{
				"expected":	TYPE_STRING,
				"given":	KEY_TITLE,
			},
		))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:196
		// _ = "end of CoverTab[195180]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:197
		_go_fuzz_dep_.CoverTab[195181]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:197
		// _ = "end of CoverTab[195181]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:197
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:197
	// _ = "end of CoverTab[195097]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:197
	_go_fuzz_dep_.CoverTab[195098]++
											if k, ok := m[KEY_TITLE].(string); ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:198
		_go_fuzz_dep_.CoverTab[195182]++
												currentSchema.title = &k
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:199
		// _ = "end of CoverTab[195182]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:200
		_go_fuzz_dep_.CoverTab[195183]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:200
		// _ = "end of CoverTab[195183]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:200
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:200
	// _ = "end of CoverTab[195098]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:200
	_go_fuzz_dep_.CoverTab[195099]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:203
	if existsMapKey(m, KEY_DESCRIPTION) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:203
		_go_fuzz_dep_.CoverTab[195184]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:203
		return !isKind(m[KEY_DESCRIPTION], reflect.String)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:203
		// _ = "end of CoverTab[195184]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:203
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:203
		_go_fuzz_dep_.CoverTab[195185]++
												return errors.New(formatErrorDescription(
			Locale.InvalidType(),
			ErrorDetails{
				"expected":	TYPE_STRING,
				"given":	KEY_DESCRIPTION,
			},
		))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:210
		// _ = "end of CoverTab[195185]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:211
		_go_fuzz_dep_.CoverTab[195186]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:211
		// _ = "end of CoverTab[195186]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:211
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:211
	// _ = "end of CoverTab[195099]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:211
	_go_fuzz_dep_.CoverTab[195100]++
											if k, ok := m[KEY_DESCRIPTION].(string); ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:212
		_go_fuzz_dep_.CoverTab[195187]++
												currentSchema.description = &k
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:213
		// _ = "end of CoverTab[195187]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:214
		_go_fuzz_dep_.CoverTab[195188]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:214
		// _ = "end of CoverTab[195188]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:214
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:214
	// _ = "end of CoverTab[195100]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:214
	_go_fuzz_dep_.CoverTab[195101]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:217
	if existsMapKey(m, KEY_REF) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:217
		_go_fuzz_dep_.CoverTab[195189]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:217
		return !isKind(m[KEY_REF], reflect.String)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:217
		// _ = "end of CoverTab[195189]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:217
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:217
		_go_fuzz_dep_.CoverTab[195190]++
												return errors.New(formatErrorDescription(
			Locale.InvalidType(),
			ErrorDetails{
				"expected":	TYPE_STRING,
				"given":	KEY_REF,
			},
		))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:224
		// _ = "end of CoverTab[195190]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:225
		_go_fuzz_dep_.CoverTab[195191]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:225
		// _ = "end of CoverTab[195191]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:225
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:225
	// _ = "end of CoverTab[195101]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:225
	_go_fuzz_dep_.CoverTab[195102]++

											if k, ok := m[KEY_REF].(string); ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:227
		_go_fuzz_dep_.CoverTab[195192]++

												jsonReference, err := gojsonreference.NewJsonReference(k)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:230
			_go_fuzz_dep_.CoverTab[195194]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:231
			// _ = "end of CoverTab[195194]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:232
			_go_fuzz_dep_.CoverTab[195195]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:232
			// _ = "end of CoverTab[195195]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:232
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:232
		// _ = "end of CoverTab[195192]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:232
		_go_fuzz_dep_.CoverTab[195193]++

												currentSchema.ref = &jsonReference

												if sch, ok := d.referencePool.Get(currentSchema.ref.String()); ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:236
			_go_fuzz_dep_.CoverTab[195196]++
													currentSchema.refSchema = sch
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:237
			// _ = "end of CoverTab[195196]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:238
			_go_fuzz_dep_.CoverTab[195197]++
													err := d.parseReference(documentNode, currentSchema)

													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:241
				_go_fuzz_dep_.CoverTab[195199]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:242
				// _ = "end of CoverTab[195199]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:243
				_go_fuzz_dep_.CoverTab[195200]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:243
				// _ = "end of CoverTab[195200]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:243
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:243
			// _ = "end of CoverTab[195197]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:243
			_go_fuzz_dep_.CoverTab[195198]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:245
			// _ = "end of CoverTab[195198]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:246
		// _ = "end of CoverTab[195193]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:247
		_go_fuzz_dep_.CoverTab[195201]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:247
		// _ = "end of CoverTab[195201]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:247
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:247
	// _ = "end of CoverTab[195102]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:247
	_go_fuzz_dep_.CoverTab[195103]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:250
	if existsMapKey(m, KEY_TYPE) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:250
		_go_fuzz_dep_.CoverTab[195202]++
												if isKind(m[KEY_TYPE], reflect.String) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:251
			_go_fuzz_dep_.CoverTab[195203]++
													if k, ok := m[KEY_TYPE].(string); ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:252
				_go_fuzz_dep_.CoverTab[195204]++
														err := currentSchema.types.Add(k)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:254
					_go_fuzz_dep_.CoverTab[195205]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:255
					// _ = "end of CoverTab[195205]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:256
					_go_fuzz_dep_.CoverTab[195206]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:256
					// _ = "end of CoverTab[195206]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:256
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:256
				// _ = "end of CoverTab[195204]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:257
				_go_fuzz_dep_.CoverTab[195207]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:257
				// _ = "end of CoverTab[195207]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:257
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:257
			// _ = "end of CoverTab[195203]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:258
			_go_fuzz_dep_.CoverTab[195208]++
													if isKind(m[KEY_TYPE], reflect.Slice) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:259
				_go_fuzz_dep_.CoverTab[195209]++
														arrayOfTypes := m[KEY_TYPE].([]interface{})
														for _, typeInArray := range arrayOfTypes {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:261
					_go_fuzz_dep_.CoverTab[195210]++
															if reflect.ValueOf(typeInArray).Kind() != reflect.String {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:262
						_go_fuzz_dep_.CoverTab[195212]++
																return errors.New(formatErrorDescription(
							Locale.InvalidType(),
							ErrorDetails{
								"expected":	TYPE_STRING + "/" + STRING_ARRAY_OF_STRINGS,
								"given":	KEY_TYPE,
							},
						))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:269
						// _ = "end of CoverTab[195212]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:270
						_go_fuzz_dep_.CoverTab[195213]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:270
						// _ = "end of CoverTab[195213]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:270
					}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:270
					// _ = "end of CoverTab[195210]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:270
					_go_fuzz_dep_.CoverTab[195211]++
															if err := currentSchema.types.Add(typeInArray.(string)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:271
						_go_fuzz_dep_.CoverTab[195214]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:272
						// _ = "end of CoverTab[195214]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:273
						_go_fuzz_dep_.CoverTab[195215]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:273
						// _ = "end of CoverTab[195215]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:273
					}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:273
					// _ = "end of CoverTab[195211]"
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:274
				// _ = "end of CoverTab[195209]"

			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:276
				_go_fuzz_dep_.CoverTab[195216]++
														return errors.New(formatErrorDescription(
					Locale.InvalidType(),
					ErrorDetails{
						"expected":	TYPE_STRING + "/" + STRING_ARRAY_OF_STRINGS,
						"given":	KEY_TYPE,
					},
				))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:283
				// _ = "end of CoverTab[195216]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:284
			// _ = "end of CoverTab[195208]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:285
		// _ = "end of CoverTab[195202]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:286
		_go_fuzz_dep_.CoverTab[195217]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:286
		// _ = "end of CoverTab[195217]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:286
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:286
	// _ = "end of CoverTab[195103]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:286
	_go_fuzz_dep_.CoverTab[195104]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:289
	if existsMapKey(m, KEY_PROPERTIES) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:289
		_go_fuzz_dep_.CoverTab[195218]++
												err := d.parseProperties(m[KEY_PROPERTIES], currentSchema)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:291
			_go_fuzz_dep_.CoverTab[195219]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:292
			// _ = "end of CoverTab[195219]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:293
			_go_fuzz_dep_.CoverTab[195220]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:293
			// _ = "end of CoverTab[195220]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:293
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:293
		// _ = "end of CoverTab[195218]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:294
		_go_fuzz_dep_.CoverTab[195221]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:294
		// _ = "end of CoverTab[195221]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:294
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:294
	// _ = "end of CoverTab[195104]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:294
	_go_fuzz_dep_.CoverTab[195105]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:297
	if existsMapKey(m, KEY_ADDITIONAL_PROPERTIES) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:297
		_go_fuzz_dep_.CoverTab[195222]++
												if isKind(m[KEY_ADDITIONAL_PROPERTIES], reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:298
			_go_fuzz_dep_.CoverTab[195223]++
													currentSchema.additionalProperties = m[KEY_ADDITIONAL_PROPERTIES].(bool)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:299
			// _ = "end of CoverTab[195223]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:300
			_go_fuzz_dep_.CoverTab[195224]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:300
			if isKind(m[KEY_ADDITIONAL_PROPERTIES], reflect.Map) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:300
				_go_fuzz_dep_.CoverTab[195225]++
														newSchema := &subSchema{property: KEY_ADDITIONAL_PROPERTIES, parent: currentSchema, ref: currentSchema.ref}
														currentSchema.additionalProperties = newSchema
														err := d.parseSchema(m[KEY_ADDITIONAL_PROPERTIES], newSchema)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:304
					_go_fuzz_dep_.CoverTab[195226]++
															return errors.New(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:305
					// _ = "end of CoverTab[195226]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:306
					_go_fuzz_dep_.CoverTab[195227]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:306
					// _ = "end of CoverTab[195227]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:306
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:306
				// _ = "end of CoverTab[195225]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:307
				_go_fuzz_dep_.CoverTab[195228]++
														return errors.New(formatErrorDescription(
					Locale.InvalidType(),
					ErrorDetails{
						"expected":	TYPE_BOOLEAN + "/" + STRING_SCHEMA,
						"given":	KEY_ADDITIONAL_PROPERTIES,
					},
				))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:314
				// _ = "end of CoverTab[195228]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:315
			// _ = "end of CoverTab[195224]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:315
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:315
		// _ = "end of CoverTab[195222]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:316
		_go_fuzz_dep_.CoverTab[195229]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:316
		// _ = "end of CoverTab[195229]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:316
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:316
	// _ = "end of CoverTab[195105]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:316
	_go_fuzz_dep_.CoverTab[195106]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:319
	if existsMapKey(m, KEY_PATTERN_PROPERTIES) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:319
		_go_fuzz_dep_.CoverTab[195230]++
												if isKind(m[KEY_PATTERN_PROPERTIES], reflect.Map) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:320
			_go_fuzz_dep_.CoverTab[195231]++
													patternPropertiesMap := m[KEY_PATTERN_PROPERTIES].(map[string]interface{})
													if len(patternPropertiesMap) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:322
				_go_fuzz_dep_.CoverTab[195232]++
														currentSchema.patternProperties = make(map[string]*subSchema)
														for k, v := range patternPropertiesMap {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:324
					_go_fuzz_dep_.CoverTab[195233]++
															_, err := regexp.MatchString(k, "")
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:326
						_go_fuzz_dep_.CoverTab[195236]++
																return errors.New(formatErrorDescription(
							Locale.RegexPattern(),
							ErrorDetails{"pattern": k},
						))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:330
						// _ = "end of CoverTab[195236]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:331
						_go_fuzz_dep_.CoverTab[195237]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:331
						// _ = "end of CoverTab[195237]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:331
					}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:331
					// _ = "end of CoverTab[195233]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:331
					_go_fuzz_dep_.CoverTab[195234]++
															newSchema := &subSchema{property: k, parent: currentSchema, ref: currentSchema.ref}
															err = d.parseSchema(v, newSchema)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:334
						_go_fuzz_dep_.CoverTab[195238]++
																return errors.New(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:335
						// _ = "end of CoverTab[195238]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:336
						_go_fuzz_dep_.CoverTab[195239]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:336
						// _ = "end of CoverTab[195239]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:336
					}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:336
					// _ = "end of CoverTab[195234]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:336
					_go_fuzz_dep_.CoverTab[195235]++
															currentSchema.patternProperties[k] = newSchema
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:337
					// _ = "end of CoverTab[195235]"
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:338
				// _ = "end of CoverTab[195232]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:339
				_go_fuzz_dep_.CoverTab[195240]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:339
				// _ = "end of CoverTab[195240]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:339
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:339
			// _ = "end of CoverTab[195231]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:340
			_go_fuzz_dep_.CoverTab[195241]++
													return errors.New(formatErrorDescription(
				Locale.InvalidType(),
				ErrorDetails{
					"expected":	STRING_SCHEMA,
					"given":	KEY_PATTERN_PROPERTIES,
				},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:347
			// _ = "end of CoverTab[195241]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:348
		// _ = "end of CoverTab[195230]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:349
		_go_fuzz_dep_.CoverTab[195242]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:349
		// _ = "end of CoverTab[195242]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:349
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:349
	// _ = "end of CoverTab[195106]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:349
	_go_fuzz_dep_.CoverTab[195107]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:352
	if existsMapKey(m, KEY_PROPERTY_NAMES) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:352
		_go_fuzz_dep_.CoverTab[195243]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:352
		return *currentSchema.draft >= Draft6
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:352
		// _ = "end of CoverTab[195243]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:352
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:352
		_go_fuzz_dep_.CoverTab[195244]++
												if isKind(m[KEY_PROPERTY_NAMES], reflect.Map, reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:353
			_go_fuzz_dep_.CoverTab[195245]++
													newSchema := &subSchema{property: KEY_PROPERTY_NAMES, parent: currentSchema, ref: currentSchema.ref}
													currentSchema.propertyNames = newSchema
													err := d.parseSchema(m[KEY_PROPERTY_NAMES], newSchema)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:357
				_go_fuzz_dep_.CoverTab[195246]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:358
				// _ = "end of CoverTab[195246]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:359
				_go_fuzz_dep_.CoverTab[195247]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:359
				// _ = "end of CoverTab[195247]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:359
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:359
			// _ = "end of CoverTab[195245]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:360
			_go_fuzz_dep_.CoverTab[195248]++
													return errors.New(formatErrorDescription(
				Locale.InvalidType(),
				ErrorDetails{
					"expected":	STRING_SCHEMA,
					"given":	KEY_PATTERN_PROPERTIES,
				},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:367
			// _ = "end of CoverTab[195248]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:368
		// _ = "end of CoverTab[195244]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:369
		_go_fuzz_dep_.CoverTab[195249]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:369
		// _ = "end of CoverTab[195249]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:369
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:369
	// _ = "end of CoverTab[195107]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:369
	_go_fuzz_dep_.CoverTab[195108]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:372
	if existsMapKey(m, KEY_DEPENDENCIES) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:372
		_go_fuzz_dep_.CoverTab[195250]++
												err := d.parseDependencies(m[KEY_DEPENDENCIES], currentSchema)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:374
			_go_fuzz_dep_.CoverTab[195251]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:375
			// _ = "end of CoverTab[195251]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:376
			_go_fuzz_dep_.CoverTab[195252]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:376
			// _ = "end of CoverTab[195252]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:376
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:376
		// _ = "end of CoverTab[195250]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:377
		_go_fuzz_dep_.CoverTab[195253]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:377
		// _ = "end of CoverTab[195253]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:377
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:377
	// _ = "end of CoverTab[195108]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:377
	_go_fuzz_dep_.CoverTab[195109]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:380
	if existsMapKey(m, KEY_ITEMS) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:380
		_go_fuzz_dep_.CoverTab[195254]++
												if isKind(m[KEY_ITEMS], reflect.Slice) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:381
			_go_fuzz_dep_.CoverTab[195255]++
													for _, itemElement := range m[KEY_ITEMS].([]interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:382
				_go_fuzz_dep_.CoverTab[195256]++
														if isKind(itemElement, reflect.Map, reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:383
					_go_fuzz_dep_.CoverTab[195258]++
															newSchema := &subSchema{parent: currentSchema, property: KEY_ITEMS}
															newSchema.ref = currentSchema.ref
															currentSchema.itemsChildren = append(currentSchema.itemsChildren, newSchema)
															err := d.parseSchema(itemElement, newSchema)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:388
						_go_fuzz_dep_.CoverTab[195259]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:389
						// _ = "end of CoverTab[195259]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:390
						_go_fuzz_dep_.CoverTab[195260]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:390
						// _ = "end of CoverTab[195260]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:390
					}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:390
					// _ = "end of CoverTab[195258]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:391
					_go_fuzz_dep_.CoverTab[195261]++
															return errors.New(formatErrorDescription(
						Locale.InvalidType(),
						ErrorDetails{
							"expected":	STRING_SCHEMA + "/" + STRING_ARRAY_OF_SCHEMAS,
							"given":	KEY_ITEMS,
						},
					))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:398
					// _ = "end of CoverTab[195261]"
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:399
				// _ = "end of CoverTab[195256]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:399
				_go_fuzz_dep_.CoverTab[195257]++
														currentSchema.itemsChildrenIsSingleSchema = false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:400
				// _ = "end of CoverTab[195257]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:401
			// _ = "end of CoverTab[195255]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:402
			_go_fuzz_dep_.CoverTab[195262]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:402
			if isKind(m[KEY_ITEMS], reflect.Map, reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:402
				_go_fuzz_dep_.CoverTab[195263]++
														newSchema := &subSchema{parent: currentSchema, property: KEY_ITEMS}
														newSchema.ref = currentSchema.ref
														currentSchema.itemsChildren = append(currentSchema.itemsChildren, newSchema)
														err := d.parseSchema(m[KEY_ITEMS], newSchema)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:407
					_go_fuzz_dep_.CoverTab[195265]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:408
					// _ = "end of CoverTab[195265]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:409
					_go_fuzz_dep_.CoverTab[195266]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:409
					// _ = "end of CoverTab[195266]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:409
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:409
				// _ = "end of CoverTab[195263]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:409
				_go_fuzz_dep_.CoverTab[195264]++
														currentSchema.itemsChildrenIsSingleSchema = true
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:410
				// _ = "end of CoverTab[195264]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:411
				_go_fuzz_dep_.CoverTab[195267]++
														return errors.New(formatErrorDescription(
					Locale.InvalidType(),
					ErrorDetails{
						"expected":	STRING_SCHEMA + "/" + STRING_ARRAY_OF_SCHEMAS,
						"given":	KEY_ITEMS,
					},
				))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:418
				// _ = "end of CoverTab[195267]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:419
			// _ = "end of CoverTab[195262]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:419
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:419
		// _ = "end of CoverTab[195254]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:420
		_go_fuzz_dep_.CoverTab[195268]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:420
		// _ = "end of CoverTab[195268]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:420
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:420
	// _ = "end of CoverTab[195109]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:420
	_go_fuzz_dep_.CoverTab[195110]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:423
	if existsMapKey(m, KEY_ADDITIONAL_ITEMS) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:423
		_go_fuzz_dep_.CoverTab[195269]++
												if isKind(m[KEY_ADDITIONAL_ITEMS], reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:424
			_go_fuzz_dep_.CoverTab[195270]++
													currentSchema.additionalItems = m[KEY_ADDITIONAL_ITEMS].(bool)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:425
			// _ = "end of CoverTab[195270]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:426
			_go_fuzz_dep_.CoverTab[195271]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:426
			if isKind(m[KEY_ADDITIONAL_ITEMS], reflect.Map) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:426
				_go_fuzz_dep_.CoverTab[195272]++
														newSchema := &subSchema{property: KEY_ADDITIONAL_ITEMS, parent: currentSchema, ref: currentSchema.ref}
														currentSchema.additionalItems = newSchema
														err := d.parseSchema(m[KEY_ADDITIONAL_ITEMS], newSchema)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:430
					_go_fuzz_dep_.CoverTab[195273]++
															return errors.New(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:431
					// _ = "end of CoverTab[195273]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:432
					_go_fuzz_dep_.CoverTab[195274]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:432
					// _ = "end of CoverTab[195274]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:432
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:432
				// _ = "end of CoverTab[195272]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:433
				_go_fuzz_dep_.CoverTab[195275]++
														return errors.New(formatErrorDescription(
					Locale.InvalidType(),
					ErrorDetails{
						"expected":	TYPE_BOOLEAN + "/" + STRING_SCHEMA,
						"given":	KEY_ADDITIONAL_ITEMS,
					},
				))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:440
				// _ = "end of CoverTab[195275]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:441
			// _ = "end of CoverTab[195271]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:441
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:441
		// _ = "end of CoverTab[195269]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:442
		_go_fuzz_dep_.CoverTab[195276]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:442
		// _ = "end of CoverTab[195276]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:442
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:442
	// _ = "end of CoverTab[195110]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:442
	_go_fuzz_dep_.CoverTab[195111]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:446
	if existsMapKey(m, KEY_MULTIPLE_OF) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:446
		_go_fuzz_dep_.CoverTab[195277]++
												multipleOfValue := mustBeNumber(m[KEY_MULTIPLE_OF])
												if multipleOfValue == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:448
			_go_fuzz_dep_.CoverTab[195280]++
													return errors.New(formatErrorDescription(
				Locale.InvalidType(),
				ErrorDetails{
					"expected":	STRING_NUMBER,
					"given":	KEY_MULTIPLE_OF,
				},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:455
			// _ = "end of CoverTab[195280]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:456
			_go_fuzz_dep_.CoverTab[195281]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:456
			// _ = "end of CoverTab[195281]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:456
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:456
		// _ = "end of CoverTab[195277]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:456
		_go_fuzz_dep_.CoverTab[195278]++
												if multipleOfValue.Cmp(big.NewRat(0, 1)) <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:457
			_go_fuzz_dep_.CoverTab[195282]++
													return errors.New(formatErrorDescription(
				Locale.GreaterThanZero(),
				ErrorDetails{"number": KEY_MULTIPLE_OF},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:461
			// _ = "end of CoverTab[195282]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:462
			_go_fuzz_dep_.CoverTab[195283]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:462
			// _ = "end of CoverTab[195283]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:462
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:462
		// _ = "end of CoverTab[195278]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:462
		_go_fuzz_dep_.CoverTab[195279]++
												currentSchema.multipleOf = multipleOfValue
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:463
		// _ = "end of CoverTab[195279]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:464
		_go_fuzz_dep_.CoverTab[195284]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:464
		// _ = "end of CoverTab[195284]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:464
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:464
	// _ = "end of CoverTab[195111]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:464
	_go_fuzz_dep_.CoverTab[195112]++

											if existsMapKey(m, KEY_MINIMUM) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:466
		_go_fuzz_dep_.CoverTab[195285]++
												minimumValue := mustBeNumber(m[KEY_MINIMUM])
												if minimumValue == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:468
			_go_fuzz_dep_.CoverTab[195287]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfA(),
				ErrorDetails{"x": KEY_MINIMUM, "y": STRING_NUMBER},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:472
			// _ = "end of CoverTab[195287]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:473
			_go_fuzz_dep_.CoverTab[195288]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:473
			// _ = "end of CoverTab[195288]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:473
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:473
		// _ = "end of CoverTab[195285]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:473
		_go_fuzz_dep_.CoverTab[195286]++
												currentSchema.minimum = minimumValue
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:474
		// _ = "end of CoverTab[195286]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:475
		_go_fuzz_dep_.CoverTab[195289]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:475
		// _ = "end of CoverTab[195289]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:475
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:475
	// _ = "end of CoverTab[195112]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:475
	_go_fuzz_dep_.CoverTab[195113]++

											if existsMapKey(m, KEY_EXCLUSIVE_MINIMUM) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:477
		_go_fuzz_dep_.CoverTab[195290]++
												switch *currentSchema.draft {
		case Draft4:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:479
			_go_fuzz_dep_.CoverTab[195291]++
													if !isKind(m[KEY_EXCLUSIVE_MINIMUM], reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:480
				_go_fuzz_dep_.CoverTab[195296]++
														return errors.New(formatErrorDescription(
					Locale.InvalidType(),
					ErrorDetails{
						"expected":	TYPE_BOOLEAN,
						"given":	KEY_EXCLUSIVE_MINIMUM,
					},
				))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:487
				// _ = "end of CoverTab[195296]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:488
				_go_fuzz_dep_.CoverTab[195297]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:488
				// _ = "end of CoverTab[195297]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:488
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:488
			// _ = "end of CoverTab[195291]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:488
			_go_fuzz_dep_.CoverTab[195292]++
													if currentSchema.minimum == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:489
				_go_fuzz_dep_.CoverTab[195298]++
														return errors.New(formatErrorDescription(
					Locale.CannotBeUsedWithout(),
					ErrorDetails{"x": KEY_EXCLUSIVE_MINIMUM, "y": KEY_MINIMUM},
				))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:493
				// _ = "end of CoverTab[195298]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:494
				_go_fuzz_dep_.CoverTab[195299]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:494
				// _ = "end of CoverTab[195299]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:494
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:494
			// _ = "end of CoverTab[195292]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:494
			_go_fuzz_dep_.CoverTab[195293]++
													if m[KEY_EXCLUSIVE_MINIMUM].(bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:495
				_go_fuzz_dep_.CoverTab[195300]++
														currentSchema.exclusiveMinimum = currentSchema.minimum
														currentSchema.minimum = nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:497
				// _ = "end of CoverTab[195300]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:498
				_go_fuzz_dep_.CoverTab[195301]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:498
				// _ = "end of CoverTab[195301]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:498
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:498
			// _ = "end of CoverTab[195293]"
		case Hybrid:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:499
			_go_fuzz_dep_.CoverTab[195294]++
													if isKind(m[KEY_EXCLUSIVE_MINIMUM], reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:500
				_go_fuzz_dep_.CoverTab[195302]++
														if currentSchema.minimum == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:501
					_go_fuzz_dep_.CoverTab[195304]++
															return errors.New(formatErrorDescription(
						Locale.CannotBeUsedWithout(),
						ErrorDetails{"x": KEY_EXCLUSIVE_MINIMUM, "y": KEY_MINIMUM},
					))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:505
					// _ = "end of CoverTab[195304]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:506
					_go_fuzz_dep_.CoverTab[195305]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:506
					// _ = "end of CoverTab[195305]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:506
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:506
				// _ = "end of CoverTab[195302]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:506
				_go_fuzz_dep_.CoverTab[195303]++
														if m[KEY_EXCLUSIVE_MINIMUM].(bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:507
					_go_fuzz_dep_.CoverTab[195306]++
															currentSchema.exclusiveMinimum = currentSchema.minimum
															currentSchema.minimum = nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:509
					// _ = "end of CoverTab[195306]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:510
					_go_fuzz_dep_.CoverTab[195307]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:510
					// _ = "end of CoverTab[195307]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:510
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:510
				// _ = "end of CoverTab[195303]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:511
				_go_fuzz_dep_.CoverTab[195308]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:511
				if isJSONNumber(m[KEY_EXCLUSIVE_MINIMUM]) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:511
					_go_fuzz_dep_.CoverTab[195309]++
															currentSchema.exclusiveMinimum = mustBeNumber(m[KEY_EXCLUSIVE_MINIMUM])
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:512
					// _ = "end of CoverTab[195309]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:513
					_go_fuzz_dep_.CoverTab[195310]++
															return errors.New(formatErrorDescription(
						Locale.InvalidType(),
						ErrorDetails{
							"expected":	TYPE_BOOLEAN + "/" + TYPE_NUMBER,
							"given":	KEY_EXCLUSIVE_MINIMUM,
						},
					))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:520
					// _ = "end of CoverTab[195310]"
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:521
				// _ = "end of CoverTab[195308]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:521
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:521
			// _ = "end of CoverTab[195294]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:522
			_go_fuzz_dep_.CoverTab[195295]++
													if isJSONNumber(m[KEY_EXCLUSIVE_MINIMUM]) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:523
				_go_fuzz_dep_.CoverTab[195311]++
														currentSchema.exclusiveMinimum = mustBeNumber(m[KEY_EXCLUSIVE_MINIMUM])
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:524
				// _ = "end of CoverTab[195311]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:525
				_go_fuzz_dep_.CoverTab[195312]++
														return errors.New(formatErrorDescription(
					Locale.InvalidType(),
					ErrorDetails{
						"expected":	TYPE_NUMBER,
						"given":	KEY_EXCLUSIVE_MINIMUM,
					},
				))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:532
				// _ = "end of CoverTab[195312]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:533
			// _ = "end of CoverTab[195295]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:534
		// _ = "end of CoverTab[195290]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:535
		_go_fuzz_dep_.CoverTab[195313]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:535
		// _ = "end of CoverTab[195313]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:535
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:535
	// _ = "end of CoverTab[195113]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:535
	_go_fuzz_dep_.CoverTab[195114]++

											if existsMapKey(m, KEY_MAXIMUM) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:537
		_go_fuzz_dep_.CoverTab[195314]++
												maximumValue := mustBeNumber(m[KEY_MAXIMUM])
												if maximumValue == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:539
			_go_fuzz_dep_.CoverTab[195316]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfA(),
				ErrorDetails{"x": KEY_MAXIMUM, "y": STRING_NUMBER},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:543
			// _ = "end of CoverTab[195316]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:544
			_go_fuzz_dep_.CoverTab[195317]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:544
			// _ = "end of CoverTab[195317]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:544
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:544
		// _ = "end of CoverTab[195314]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:544
		_go_fuzz_dep_.CoverTab[195315]++
												currentSchema.maximum = maximumValue
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:545
		// _ = "end of CoverTab[195315]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:546
		_go_fuzz_dep_.CoverTab[195318]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:546
		// _ = "end of CoverTab[195318]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:546
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:546
	// _ = "end of CoverTab[195114]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:546
	_go_fuzz_dep_.CoverTab[195115]++

											if existsMapKey(m, KEY_EXCLUSIVE_MAXIMUM) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:548
		_go_fuzz_dep_.CoverTab[195319]++
												switch *currentSchema.draft {
		case Draft4:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:550
			_go_fuzz_dep_.CoverTab[195320]++
													if !isKind(m[KEY_EXCLUSIVE_MAXIMUM], reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:551
				_go_fuzz_dep_.CoverTab[195325]++
														return errors.New(formatErrorDescription(
					Locale.InvalidType(),
					ErrorDetails{
						"expected":	TYPE_BOOLEAN,
						"given":	KEY_EXCLUSIVE_MAXIMUM,
					},
				))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:558
				// _ = "end of CoverTab[195325]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:559
				_go_fuzz_dep_.CoverTab[195326]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:559
				// _ = "end of CoverTab[195326]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:559
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:559
			// _ = "end of CoverTab[195320]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:559
			_go_fuzz_dep_.CoverTab[195321]++
													if currentSchema.maximum == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:560
				_go_fuzz_dep_.CoverTab[195327]++
														return errors.New(formatErrorDescription(
					Locale.CannotBeUsedWithout(),
					ErrorDetails{"x": KEY_EXCLUSIVE_MAXIMUM, "y": KEY_MAXIMUM},
				))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:564
				// _ = "end of CoverTab[195327]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:565
				_go_fuzz_dep_.CoverTab[195328]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:565
				// _ = "end of CoverTab[195328]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:565
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:565
			// _ = "end of CoverTab[195321]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:565
			_go_fuzz_dep_.CoverTab[195322]++
													if m[KEY_EXCLUSIVE_MAXIMUM].(bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:566
				_go_fuzz_dep_.CoverTab[195329]++
														currentSchema.exclusiveMaximum = currentSchema.maximum
														currentSchema.maximum = nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:568
				// _ = "end of CoverTab[195329]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:569
				_go_fuzz_dep_.CoverTab[195330]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:569
				// _ = "end of CoverTab[195330]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:569
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:569
			// _ = "end of CoverTab[195322]"
		case Hybrid:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:570
			_go_fuzz_dep_.CoverTab[195323]++
													if isKind(m[KEY_EXCLUSIVE_MAXIMUM], reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:571
				_go_fuzz_dep_.CoverTab[195331]++
														if currentSchema.maximum == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:572
					_go_fuzz_dep_.CoverTab[195333]++
															return errors.New(formatErrorDescription(
						Locale.CannotBeUsedWithout(),
						ErrorDetails{"x": KEY_EXCLUSIVE_MAXIMUM, "y": KEY_MAXIMUM},
					))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:576
					// _ = "end of CoverTab[195333]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:577
					_go_fuzz_dep_.CoverTab[195334]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:577
					// _ = "end of CoverTab[195334]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:577
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:577
				// _ = "end of CoverTab[195331]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:577
				_go_fuzz_dep_.CoverTab[195332]++
														if m[KEY_EXCLUSIVE_MAXIMUM].(bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:578
					_go_fuzz_dep_.CoverTab[195335]++
															currentSchema.exclusiveMaximum = currentSchema.maximum
															currentSchema.maximum = nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:580
					// _ = "end of CoverTab[195335]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:581
					_go_fuzz_dep_.CoverTab[195336]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:581
					// _ = "end of CoverTab[195336]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:581
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:581
				// _ = "end of CoverTab[195332]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:582
				_go_fuzz_dep_.CoverTab[195337]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:582
				if isJSONNumber(m[KEY_EXCLUSIVE_MAXIMUM]) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:582
					_go_fuzz_dep_.CoverTab[195338]++
															currentSchema.exclusiveMaximum = mustBeNumber(m[KEY_EXCLUSIVE_MAXIMUM])
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:583
					// _ = "end of CoverTab[195338]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:584
					_go_fuzz_dep_.CoverTab[195339]++
															return errors.New(formatErrorDescription(
						Locale.InvalidType(),
						ErrorDetails{
							"expected":	TYPE_BOOLEAN + "/" + TYPE_NUMBER,
							"given":	KEY_EXCLUSIVE_MAXIMUM,
						},
					))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:591
					// _ = "end of CoverTab[195339]"
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:592
				// _ = "end of CoverTab[195337]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:592
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:592
			// _ = "end of CoverTab[195323]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:593
			_go_fuzz_dep_.CoverTab[195324]++
													if isJSONNumber(m[KEY_EXCLUSIVE_MAXIMUM]) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:594
				_go_fuzz_dep_.CoverTab[195340]++
														currentSchema.exclusiveMaximum = mustBeNumber(m[KEY_EXCLUSIVE_MAXIMUM])
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:595
				// _ = "end of CoverTab[195340]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:596
				_go_fuzz_dep_.CoverTab[195341]++
														return errors.New(formatErrorDescription(
					Locale.InvalidType(),
					ErrorDetails{
						"expected":	TYPE_NUMBER,
						"given":	KEY_EXCLUSIVE_MAXIMUM,
					},
				))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:603
				// _ = "end of CoverTab[195341]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:604
			// _ = "end of CoverTab[195324]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:605
		// _ = "end of CoverTab[195319]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:606
		_go_fuzz_dep_.CoverTab[195342]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:606
		// _ = "end of CoverTab[195342]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:606
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:606
	// _ = "end of CoverTab[195115]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:606
	_go_fuzz_dep_.CoverTab[195116]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:610
	if existsMapKey(m, KEY_MIN_LENGTH) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:610
		_go_fuzz_dep_.CoverTab[195343]++
												minLengthIntegerValue := mustBeInteger(m[KEY_MIN_LENGTH])
												if minLengthIntegerValue == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:612
			_go_fuzz_dep_.CoverTab[195346]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfAn(),
				ErrorDetails{"x": KEY_MIN_LENGTH, "y": TYPE_INTEGER},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:616
			// _ = "end of CoverTab[195346]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:617
			_go_fuzz_dep_.CoverTab[195347]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:617
			// _ = "end of CoverTab[195347]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:617
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:617
		// _ = "end of CoverTab[195343]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:617
		_go_fuzz_dep_.CoverTab[195344]++
												if *minLengthIntegerValue < 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:618
			_go_fuzz_dep_.CoverTab[195348]++
													return errors.New(formatErrorDescription(
				Locale.MustBeGTEZero(),
				ErrorDetails{"key": KEY_MIN_LENGTH},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:622
			// _ = "end of CoverTab[195348]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:623
			_go_fuzz_dep_.CoverTab[195349]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:623
			// _ = "end of CoverTab[195349]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:623
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:623
		// _ = "end of CoverTab[195344]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:623
		_go_fuzz_dep_.CoverTab[195345]++
												currentSchema.minLength = minLengthIntegerValue
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:624
		// _ = "end of CoverTab[195345]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:625
		_go_fuzz_dep_.CoverTab[195350]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:625
		// _ = "end of CoverTab[195350]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:625
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:625
	// _ = "end of CoverTab[195116]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:625
	_go_fuzz_dep_.CoverTab[195117]++

											if existsMapKey(m, KEY_MAX_LENGTH) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:627
		_go_fuzz_dep_.CoverTab[195351]++
												maxLengthIntegerValue := mustBeInteger(m[KEY_MAX_LENGTH])
												if maxLengthIntegerValue == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:629
			_go_fuzz_dep_.CoverTab[195354]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfAn(),
				ErrorDetails{"x": KEY_MAX_LENGTH, "y": TYPE_INTEGER},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:633
			// _ = "end of CoverTab[195354]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:634
			_go_fuzz_dep_.CoverTab[195355]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:634
			// _ = "end of CoverTab[195355]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:634
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:634
		// _ = "end of CoverTab[195351]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:634
		_go_fuzz_dep_.CoverTab[195352]++
												if *maxLengthIntegerValue < 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:635
			_go_fuzz_dep_.CoverTab[195356]++
													return errors.New(formatErrorDescription(
				Locale.MustBeGTEZero(),
				ErrorDetails{"key": KEY_MAX_LENGTH},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:639
			// _ = "end of CoverTab[195356]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:640
			_go_fuzz_dep_.CoverTab[195357]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:640
			// _ = "end of CoverTab[195357]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:640
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:640
		// _ = "end of CoverTab[195352]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:640
		_go_fuzz_dep_.CoverTab[195353]++
												currentSchema.maxLength = maxLengthIntegerValue
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:641
		// _ = "end of CoverTab[195353]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:642
		_go_fuzz_dep_.CoverTab[195358]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:642
		// _ = "end of CoverTab[195358]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:642
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:642
	// _ = "end of CoverTab[195117]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:642
	_go_fuzz_dep_.CoverTab[195118]++

											if currentSchema.minLength != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:644
		_go_fuzz_dep_.CoverTab[195359]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:644
		return currentSchema.maxLength != nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:644
		// _ = "end of CoverTab[195359]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:644
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:644
		_go_fuzz_dep_.CoverTab[195360]++
												if *currentSchema.minLength > *currentSchema.maxLength {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:645
			_go_fuzz_dep_.CoverTab[195361]++
													return errors.New(formatErrorDescription(
				Locale.CannotBeGT(),
				ErrorDetails{"x": KEY_MIN_LENGTH, "y": KEY_MAX_LENGTH},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:649
			// _ = "end of CoverTab[195361]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:650
			_go_fuzz_dep_.CoverTab[195362]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:650
			// _ = "end of CoverTab[195362]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:650
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:650
		// _ = "end of CoverTab[195360]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:651
		_go_fuzz_dep_.CoverTab[195363]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:651
		// _ = "end of CoverTab[195363]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:651
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:651
	// _ = "end of CoverTab[195118]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:651
	_go_fuzz_dep_.CoverTab[195119]++

											if existsMapKey(m, KEY_PATTERN) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:653
		_go_fuzz_dep_.CoverTab[195364]++
												if isKind(m[KEY_PATTERN], reflect.String) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:654
			_go_fuzz_dep_.CoverTab[195365]++
													regexpObject, err := regexp.Compile(m[KEY_PATTERN].(string))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:656
				_go_fuzz_dep_.CoverTab[195367]++
														return errors.New(formatErrorDescription(
					Locale.MustBeValidRegex(),
					ErrorDetails{"key": KEY_PATTERN},
				))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:660
				// _ = "end of CoverTab[195367]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:661
				_go_fuzz_dep_.CoverTab[195368]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:661
				// _ = "end of CoverTab[195368]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:661
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:661
			// _ = "end of CoverTab[195365]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:661
			_go_fuzz_dep_.CoverTab[195366]++
													currentSchema.pattern = regexpObject
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:662
			// _ = "end of CoverTab[195366]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:663
			_go_fuzz_dep_.CoverTab[195369]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfA(),
				ErrorDetails{"x": KEY_PATTERN, "y": TYPE_STRING},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:667
			// _ = "end of CoverTab[195369]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:668
		// _ = "end of CoverTab[195364]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:669
		_go_fuzz_dep_.CoverTab[195370]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:669
		// _ = "end of CoverTab[195370]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:669
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:669
	// _ = "end of CoverTab[195119]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:669
	_go_fuzz_dep_.CoverTab[195120]++

											if existsMapKey(m, KEY_FORMAT) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:671
		_go_fuzz_dep_.CoverTab[195371]++
												formatString, ok := m[KEY_FORMAT].(string)
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:673
			_go_fuzz_dep_.CoverTab[195373]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfType(),
				ErrorDetails{"key": KEY_FORMAT, "type": TYPE_STRING},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:677
			// _ = "end of CoverTab[195373]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:678
			_go_fuzz_dep_.CoverTab[195374]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:678
			// _ = "end of CoverTab[195374]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:678
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:678
		// _ = "end of CoverTab[195371]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:678
		_go_fuzz_dep_.CoverTab[195372]++
												currentSchema.format = formatString
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:679
		// _ = "end of CoverTab[195372]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:680
		_go_fuzz_dep_.CoverTab[195375]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:680
		// _ = "end of CoverTab[195375]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:680
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:680
	// _ = "end of CoverTab[195120]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:680
	_go_fuzz_dep_.CoverTab[195121]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:684
	if existsMapKey(m, KEY_MIN_PROPERTIES) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:684
		_go_fuzz_dep_.CoverTab[195376]++
												minPropertiesIntegerValue := mustBeInteger(m[KEY_MIN_PROPERTIES])
												if minPropertiesIntegerValue == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:686
			_go_fuzz_dep_.CoverTab[195379]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfAn(),
				ErrorDetails{"x": KEY_MIN_PROPERTIES, "y": TYPE_INTEGER},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:690
			// _ = "end of CoverTab[195379]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:691
			_go_fuzz_dep_.CoverTab[195380]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:691
			// _ = "end of CoverTab[195380]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:691
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:691
		// _ = "end of CoverTab[195376]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:691
		_go_fuzz_dep_.CoverTab[195377]++
												if *minPropertiesIntegerValue < 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:692
			_go_fuzz_dep_.CoverTab[195381]++
													return errors.New(formatErrorDescription(
				Locale.MustBeGTEZero(),
				ErrorDetails{"key": KEY_MIN_PROPERTIES},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:696
			// _ = "end of CoverTab[195381]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:697
			_go_fuzz_dep_.CoverTab[195382]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:697
			// _ = "end of CoverTab[195382]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:697
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:697
		// _ = "end of CoverTab[195377]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:697
		_go_fuzz_dep_.CoverTab[195378]++
												currentSchema.minProperties = minPropertiesIntegerValue
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:698
		// _ = "end of CoverTab[195378]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:699
		_go_fuzz_dep_.CoverTab[195383]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:699
		// _ = "end of CoverTab[195383]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:699
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:699
	// _ = "end of CoverTab[195121]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:699
	_go_fuzz_dep_.CoverTab[195122]++

											if existsMapKey(m, KEY_MAX_PROPERTIES) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:701
		_go_fuzz_dep_.CoverTab[195384]++
												maxPropertiesIntegerValue := mustBeInteger(m[KEY_MAX_PROPERTIES])
												if maxPropertiesIntegerValue == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:703
			_go_fuzz_dep_.CoverTab[195387]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfAn(),
				ErrorDetails{"x": KEY_MAX_PROPERTIES, "y": TYPE_INTEGER},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:707
			// _ = "end of CoverTab[195387]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:708
			_go_fuzz_dep_.CoverTab[195388]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:708
			// _ = "end of CoverTab[195388]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:708
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:708
		// _ = "end of CoverTab[195384]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:708
		_go_fuzz_dep_.CoverTab[195385]++
												if *maxPropertiesIntegerValue < 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:709
			_go_fuzz_dep_.CoverTab[195389]++
													return errors.New(formatErrorDescription(
				Locale.MustBeGTEZero(),
				ErrorDetails{"key": KEY_MAX_PROPERTIES},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:713
			// _ = "end of CoverTab[195389]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:714
			_go_fuzz_dep_.CoverTab[195390]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:714
			// _ = "end of CoverTab[195390]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:714
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:714
		// _ = "end of CoverTab[195385]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:714
		_go_fuzz_dep_.CoverTab[195386]++
												currentSchema.maxProperties = maxPropertiesIntegerValue
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:715
		// _ = "end of CoverTab[195386]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:716
		_go_fuzz_dep_.CoverTab[195391]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:716
		// _ = "end of CoverTab[195391]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:716
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:716
	// _ = "end of CoverTab[195122]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:716
	_go_fuzz_dep_.CoverTab[195123]++

											if currentSchema.minProperties != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:718
		_go_fuzz_dep_.CoverTab[195392]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:718
		return currentSchema.maxProperties != nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:718
		// _ = "end of CoverTab[195392]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:718
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:718
		_go_fuzz_dep_.CoverTab[195393]++
												if *currentSchema.minProperties > *currentSchema.maxProperties {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:719
			_go_fuzz_dep_.CoverTab[195394]++
													return errors.New(formatErrorDescription(
				Locale.KeyCannotBeGreaterThan(),
				ErrorDetails{"key": KEY_MIN_PROPERTIES, "y": KEY_MAX_PROPERTIES},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:723
			// _ = "end of CoverTab[195394]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:724
			_go_fuzz_dep_.CoverTab[195395]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:724
			// _ = "end of CoverTab[195395]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:724
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:724
		// _ = "end of CoverTab[195393]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:725
		_go_fuzz_dep_.CoverTab[195396]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:725
		// _ = "end of CoverTab[195396]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:725
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:725
	// _ = "end of CoverTab[195123]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:725
	_go_fuzz_dep_.CoverTab[195124]++

											if existsMapKey(m, KEY_REQUIRED) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:727
		_go_fuzz_dep_.CoverTab[195397]++
												if isKind(m[KEY_REQUIRED], reflect.Slice) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:728
			_go_fuzz_dep_.CoverTab[195398]++
													requiredValues := m[KEY_REQUIRED].([]interface{})
													for _, requiredValue := range requiredValues {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:730
				_go_fuzz_dep_.CoverTab[195399]++
														if isKind(requiredValue, reflect.String) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:731
					_go_fuzz_dep_.CoverTab[195400]++
															if isStringInSlice(currentSchema.required, requiredValue.(string)) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:732
						_go_fuzz_dep_.CoverTab[195402]++
																return errors.New(formatErrorDescription(
							Locale.KeyItemsMustBeUnique(),
							ErrorDetails{"key": KEY_REQUIRED},
						))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:736
						// _ = "end of CoverTab[195402]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:737
						_go_fuzz_dep_.CoverTab[195403]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:737
						// _ = "end of CoverTab[195403]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:737
					}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:737
					// _ = "end of CoverTab[195400]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:737
					_go_fuzz_dep_.CoverTab[195401]++
															currentSchema.required = append(currentSchema.required, requiredValue.(string))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:738
					// _ = "end of CoverTab[195401]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:739
					_go_fuzz_dep_.CoverTab[195404]++
															return errors.New(formatErrorDescription(
						Locale.KeyItemsMustBeOfType(),
						ErrorDetails{"key": KEY_REQUIRED, "type": TYPE_STRING},
					))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:743
					// _ = "end of CoverTab[195404]"
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:744
				// _ = "end of CoverTab[195399]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:745
			// _ = "end of CoverTab[195398]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:746
			_go_fuzz_dep_.CoverTab[195405]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfAn(),
				ErrorDetails{"x": KEY_REQUIRED, "y": TYPE_ARRAY},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:750
			// _ = "end of CoverTab[195405]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:751
		// _ = "end of CoverTab[195397]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:752
		_go_fuzz_dep_.CoverTab[195406]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:752
		// _ = "end of CoverTab[195406]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:752
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:752
	// _ = "end of CoverTab[195124]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:752
	_go_fuzz_dep_.CoverTab[195125]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:756
	if existsMapKey(m, KEY_MIN_ITEMS) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:756
		_go_fuzz_dep_.CoverTab[195407]++
												minItemsIntegerValue := mustBeInteger(m[KEY_MIN_ITEMS])
												if minItemsIntegerValue == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:758
			_go_fuzz_dep_.CoverTab[195410]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfAn(),
				ErrorDetails{"x": KEY_MIN_ITEMS, "y": TYPE_INTEGER},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:762
			// _ = "end of CoverTab[195410]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:763
			_go_fuzz_dep_.CoverTab[195411]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:763
			// _ = "end of CoverTab[195411]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:763
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:763
		// _ = "end of CoverTab[195407]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:763
		_go_fuzz_dep_.CoverTab[195408]++
												if *minItemsIntegerValue < 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:764
			_go_fuzz_dep_.CoverTab[195412]++
													return errors.New(formatErrorDescription(
				Locale.MustBeGTEZero(),
				ErrorDetails{"key": KEY_MIN_ITEMS},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:768
			// _ = "end of CoverTab[195412]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:769
			_go_fuzz_dep_.CoverTab[195413]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:769
			// _ = "end of CoverTab[195413]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:769
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:769
		// _ = "end of CoverTab[195408]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:769
		_go_fuzz_dep_.CoverTab[195409]++
												currentSchema.minItems = minItemsIntegerValue
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:770
		// _ = "end of CoverTab[195409]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:771
		_go_fuzz_dep_.CoverTab[195414]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:771
		// _ = "end of CoverTab[195414]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:771
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:771
	// _ = "end of CoverTab[195125]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:771
	_go_fuzz_dep_.CoverTab[195126]++

											if existsMapKey(m, KEY_MAX_ITEMS) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:773
		_go_fuzz_dep_.CoverTab[195415]++
												maxItemsIntegerValue := mustBeInteger(m[KEY_MAX_ITEMS])
												if maxItemsIntegerValue == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:775
			_go_fuzz_dep_.CoverTab[195418]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfAn(),
				ErrorDetails{"x": KEY_MAX_ITEMS, "y": TYPE_INTEGER},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:779
			// _ = "end of CoverTab[195418]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:780
			_go_fuzz_dep_.CoverTab[195419]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:780
			// _ = "end of CoverTab[195419]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:780
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:780
		// _ = "end of CoverTab[195415]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:780
		_go_fuzz_dep_.CoverTab[195416]++
												if *maxItemsIntegerValue < 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:781
			_go_fuzz_dep_.CoverTab[195420]++
													return errors.New(formatErrorDescription(
				Locale.MustBeGTEZero(),
				ErrorDetails{"key": KEY_MAX_ITEMS},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:785
			// _ = "end of CoverTab[195420]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:786
			_go_fuzz_dep_.CoverTab[195421]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:786
			// _ = "end of CoverTab[195421]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:786
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:786
		// _ = "end of CoverTab[195416]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:786
		_go_fuzz_dep_.CoverTab[195417]++
												currentSchema.maxItems = maxItemsIntegerValue
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:787
		// _ = "end of CoverTab[195417]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:788
		_go_fuzz_dep_.CoverTab[195422]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:788
		// _ = "end of CoverTab[195422]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:788
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:788
	// _ = "end of CoverTab[195126]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:788
	_go_fuzz_dep_.CoverTab[195127]++

											if existsMapKey(m, KEY_UNIQUE_ITEMS) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:790
		_go_fuzz_dep_.CoverTab[195423]++
												if isKind(m[KEY_UNIQUE_ITEMS], reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:791
			_go_fuzz_dep_.CoverTab[195424]++
													currentSchema.uniqueItems = m[KEY_UNIQUE_ITEMS].(bool)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:792
			// _ = "end of CoverTab[195424]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:793
			_go_fuzz_dep_.CoverTab[195425]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfA(),
				ErrorDetails{"x": KEY_UNIQUE_ITEMS, "y": TYPE_BOOLEAN},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:797
			// _ = "end of CoverTab[195425]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:798
		// _ = "end of CoverTab[195423]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:799
		_go_fuzz_dep_.CoverTab[195426]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:799
		// _ = "end of CoverTab[195426]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:799
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:799
	// _ = "end of CoverTab[195127]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:799
	_go_fuzz_dep_.CoverTab[195128]++

											if existsMapKey(m, KEY_CONTAINS) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:801
		_go_fuzz_dep_.CoverTab[195427]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:801
		return *currentSchema.draft >= Draft6
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:801
		// _ = "end of CoverTab[195427]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:801
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:801
		_go_fuzz_dep_.CoverTab[195428]++
												newSchema := &subSchema{property: KEY_CONTAINS, parent: currentSchema, ref: currentSchema.ref}
												currentSchema.contains = newSchema
												err := d.parseSchema(m[KEY_CONTAINS], newSchema)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:805
			_go_fuzz_dep_.CoverTab[195429]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:806
			// _ = "end of CoverTab[195429]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:807
			_go_fuzz_dep_.CoverTab[195430]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:807
			// _ = "end of CoverTab[195430]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:807
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:807
		// _ = "end of CoverTab[195428]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:808
		_go_fuzz_dep_.CoverTab[195431]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:808
		// _ = "end of CoverTab[195431]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:808
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:808
	// _ = "end of CoverTab[195128]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:808
	_go_fuzz_dep_.CoverTab[195129]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:812
	if existsMapKey(m, KEY_CONST) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:812
		_go_fuzz_dep_.CoverTab[195432]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:812
		return *currentSchema.draft >= Draft6
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:812
		// _ = "end of CoverTab[195432]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:812
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:812
		_go_fuzz_dep_.CoverTab[195433]++
												is, err := marshalWithoutNumber(m[KEY_CONST])
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:814
			_go_fuzz_dep_.CoverTab[195435]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:815
			// _ = "end of CoverTab[195435]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:816
			_go_fuzz_dep_.CoverTab[195436]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:816
			// _ = "end of CoverTab[195436]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:816
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:816
		// _ = "end of CoverTab[195433]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:816
		_go_fuzz_dep_.CoverTab[195434]++
												currentSchema._const = is
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:817
		// _ = "end of CoverTab[195434]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:818
		_go_fuzz_dep_.CoverTab[195437]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:818
		// _ = "end of CoverTab[195437]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:818
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:818
	// _ = "end of CoverTab[195129]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:818
	_go_fuzz_dep_.CoverTab[195130]++

											if existsMapKey(m, KEY_ENUM) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:820
		_go_fuzz_dep_.CoverTab[195438]++
												if isKind(m[KEY_ENUM], reflect.Slice) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:821
			_go_fuzz_dep_.CoverTab[195439]++
													for _, v := range m[KEY_ENUM].([]interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:822
				_go_fuzz_dep_.CoverTab[195440]++
														is, err := marshalWithoutNumber(v)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:824
					_go_fuzz_dep_.CoverTab[195443]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:825
					// _ = "end of CoverTab[195443]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:826
					_go_fuzz_dep_.CoverTab[195444]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:826
					// _ = "end of CoverTab[195444]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:826
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:826
				// _ = "end of CoverTab[195440]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:826
				_go_fuzz_dep_.CoverTab[195441]++
														if isStringInSlice(currentSchema.enum, *is) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:827
					_go_fuzz_dep_.CoverTab[195445]++
															return errors.New(formatErrorDescription(
						Locale.KeyItemsMustBeUnique(),
						ErrorDetails{"key": KEY_ENUM},
					))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:831
					// _ = "end of CoverTab[195445]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:832
					_go_fuzz_dep_.CoverTab[195446]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:832
					// _ = "end of CoverTab[195446]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:832
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:832
				// _ = "end of CoverTab[195441]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:832
				_go_fuzz_dep_.CoverTab[195442]++
														currentSchema.enum = append(currentSchema.enum, *is)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:833
				// _ = "end of CoverTab[195442]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:834
			// _ = "end of CoverTab[195439]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:835
			_go_fuzz_dep_.CoverTab[195447]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfAn(),
				ErrorDetails{"x": KEY_ENUM, "y": TYPE_ARRAY},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:839
			// _ = "end of CoverTab[195447]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:840
		// _ = "end of CoverTab[195438]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:841
		_go_fuzz_dep_.CoverTab[195448]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:841
		// _ = "end of CoverTab[195448]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:841
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:841
	// _ = "end of CoverTab[195130]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:841
	_go_fuzz_dep_.CoverTab[195131]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:845
	if existsMapKey(m, KEY_ONE_OF) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:845
		_go_fuzz_dep_.CoverTab[195449]++
												if isKind(m[KEY_ONE_OF], reflect.Slice) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:846
			_go_fuzz_dep_.CoverTab[195450]++
													for _, v := range m[KEY_ONE_OF].([]interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:847
				_go_fuzz_dep_.CoverTab[195451]++
														newSchema := &subSchema{property: KEY_ONE_OF, parent: currentSchema, ref: currentSchema.ref}
														currentSchema.oneOf = append(currentSchema.oneOf, newSchema)
														err := d.parseSchema(v, newSchema)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:851
					_go_fuzz_dep_.CoverTab[195452]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:852
					// _ = "end of CoverTab[195452]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:853
					_go_fuzz_dep_.CoverTab[195453]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:853
					// _ = "end of CoverTab[195453]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:853
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:853
				// _ = "end of CoverTab[195451]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:854
			// _ = "end of CoverTab[195450]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:855
			_go_fuzz_dep_.CoverTab[195454]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfAn(),
				ErrorDetails{"x": KEY_ONE_OF, "y": TYPE_ARRAY},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:859
			// _ = "end of CoverTab[195454]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:860
		// _ = "end of CoverTab[195449]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:861
		_go_fuzz_dep_.CoverTab[195455]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:861
		// _ = "end of CoverTab[195455]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:861
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:861
	// _ = "end of CoverTab[195131]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:861
	_go_fuzz_dep_.CoverTab[195132]++

											if existsMapKey(m, KEY_ANY_OF) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:863
		_go_fuzz_dep_.CoverTab[195456]++
												if isKind(m[KEY_ANY_OF], reflect.Slice) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:864
			_go_fuzz_dep_.CoverTab[195457]++
													for _, v := range m[KEY_ANY_OF].([]interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:865
				_go_fuzz_dep_.CoverTab[195458]++
														newSchema := &subSchema{property: KEY_ANY_OF, parent: currentSchema, ref: currentSchema.ref}
														currentSchema.anyOf = append(currentSchema.anyOf, newSchema)
														err := d.parseSchema(v, newSchema)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:869
					_go_fuzz_dep_.CoverTab[195459]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:870
					// _ = "end of CoverTab[195459]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:871
					_go_fuzz_dep_.CoverTab[195460]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:871
					// _ = "end of CoverTab[195460]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:871
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:871
				// _ = "end of CoverTab[195458]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:872
			// _ = "end of CoverTab[195457]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:873
			_go_fuzz_dep_.CoverTab[195461]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfAn(),
				ErrorDetails{"x": KEY_ANY_OF, "y": TYPE_ARRAY},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:877
			// _ = "end of CoverTab[195461]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:878
		// _ = "end of CoverTab[195456]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:879
		_go_fuzz_dep_.CoverTab[195462]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:879
		// _ = "end of CoverTab[195462]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:879
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:879
	// _ = "end of CoverTab[195132]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:879
	_go_fuzz_dep_.CoverTab[195133]++

											if existsMapKey(m, KEY_ALL_OF) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:881
		_go_fuzz_dep_.CoverTab[195463]++
												if isKind(m[KEY_ALL_OF], reflect.Slice) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:882
			_go_fuzz_dep_.CoverTab[195464]++
													for _, v := range m[KEY_ALL_OF].([]interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:883
				_go_fuzz_dep_.CoverTab[195465]++
														newSchema := &subSchema{property: KEY_ALL_OF, parent: currentSchema, ref: currentSchema.ref}
														currentSchema.allOf = append(currentSchema.allOf, newSchema)
														err := d.parseSchema(v, newSchema)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:887
					_go_fuzz_dep_.CoverTab[195466]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:888
					// _ = "end of CoverTab[195466]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:889
					_go_fuzz_dep_.CoverTab[195467]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:889
					// _ = "end of CoverTab[195467]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:889
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:889
				// _ = "end of CoverTab[195465]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:890
			// _ = "end of CoverTab[195464]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:891
			_go_fuzz_dep_.CoverTab[195468]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfAn(),
				ErrorDetails{"x": KEY_ANY_OF, "y": TYPE_ARRAY},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:895
			// _ = "end of CoverTab[195468]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:896
		// _ = "end of CoverTab[195463]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:897
		_go_fuzz_dep_.CoverTab[195469]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:897
		// _ = "end of CoverTab[195469]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:897
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:897
	// _ = "end of CoverTab[195133]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:897
	_go_fuzz_dep_.CoverTab[195134]++

											if existsMapKey(m, KEY_NOT) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:899
		_go_fuzz_dep_.CoverTab[195470]++
												if isKind(m[KEY_NOT], reflect.Map, reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:900
			_go_fuzz_dep_.CoverTab[195471]++
													newSchema := &subSchema{property: KEY_NOT, parent: currentSchema, ref: currentSchema.ref}
													currentSchema.not = newSchema
													err := d.parseSchema(m[KEY_NOT], newSchema)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:904
				_go_fuzz_dep_.CoverTab[195472]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:905
				// _ = "end of CoverTab[195472]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:906
				_go_fuzz_dep_.CoverTab[195473]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:906
				// _ = "end of CoverTab[195473]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:906
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:906
			// _ = "end of CoverTab[195471]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:907
			_go_fuzz_dep_.CoverTab[195474]++
													return errors.New(formatErrorDescription(
				Locale.MustBeOfAn(),
				ErrorDetails{"x": KEY_NOT, "y": TYPE_OBJECT},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:911
			// _ = "end of CoverTab[195474]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:912
		// _ = "end of CoverTab[195470]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:913
		_go_fuzz_dep_.CoverTab[195475]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:913
		// _ = "end of CoverTab[195475]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:913
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:913
	// _ = "end of CoverTab[195134]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:913
	_go_fuzz_dep_.CoverTab[195135]++

											if *currentSchema.draft >= Draft7 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:915
		_go_fuzz_dep_.CoverTab[195476]++
												if existsMapKey(m, KEY_IF) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:916
			_go_fuzz_dep_.CoverTab[195479]++
													if isKind(m[KEY_IF], reflect.Map, reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:917
				_go_fuzz_dep_.CoverTab[195480]++
														newSchema := &subSchema{property: KEY_IF, parent: currentSchema, ref: currentSchema.ref}
														currentSchema._if = newSchema
														err := d.parseSchema(m[KEY_IF], newSchema)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:921
					_go_fuzz_dep_.CoverTab[195481]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:922
					// _ = "end of CoverTab[195481]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:923
					_go_fuzz_dep_.CoverTab[195482]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:923
					// _ = "end of CoverTab[195482]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:923
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:923
				// _ = "end of CoverTab[195480]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:924
				_go_fuzz_dep_.CoverTab[195483]++
														return errors.New(formatErrorDescription(
					Locale.MustBeOfAn(),
					ErrorDetails{"x": KEY_IF, "y": TYPE_OBJECT},
				))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:928
				// _ = "end of CoverTab[195483]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:929
			// _ = "end of CoverTab[195479]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:930
			_go_fuzz_dep_.CoverTab[195484]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:930
			// _ = "end of CoverTab[195484]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:930
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:930
		// _ = "end of CoverTab[195476]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:930
		_go_fuzz_dep_.CoverTab[195477]++

												if existsMapKey(m, KEY_THEN) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:932
			_go_fuzz_dep_.CoverTab[195485]++
													if isKind(m[KEY_THEN], reflect.Map, reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:933
				_go_fuzz_dep_.CoverTab[195486]++
														newSchema := &subSchema{property: KEY_THEN, parent: currentSchema, ref: currentSchema.ref}
														currentSchema._then = newSchema
														err := d.parseSchema(m[KEY_THEN], newSchema)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:937
					_go_fuzz_dep_.CoverTab[195487]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:938
					// _ = "end of CoverTab[195487]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:939
					_go_fuzz_dep_.CoverTab[195488]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:939
					// _ = "end of CoverTab[195488]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:939
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:939
				// _ = "end of CoverTab[195486]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:940
				_go_fuzz_dep_.CoverTab[195489]++
														return errors.New(formatErrorDescription(
					Locale.MustBeOfAn(),
					ErrorDetails{"x": KEY_THEN, "y": TYPE_OBJECT},
				))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:944
				// _ = "end of CoverTab[195489]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:945
			// _ = "end of CoverTab[195485]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:946
			_go_fuzz_dep_.CoverTab[195490]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:946
			// _ = "end of CoverTab[195490]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:946
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:946
		// _ = "end of CoverTab[195477]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:946
		_go_fuzz_dep_.CoverTab[195478]++

												if existsMapKey(m, KEY_ELSE) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:948
			_go_fuzz_dep_.CoverTab[195491]++
													if isKind(m[KEY_ELSE], reflect.Map, reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:949
				_go_fuzz_dep_.CoverTab[195492]++
														newSchema := &subSchema{property: KEY_ELSE, parent: currentSchema, ref: currentSchema.ref}
														currentSchema._else = newSchema
														err := d.parseSchema(m[KEY_ELSE], newSchema)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:953
					_go_fuzz_dep_.CoverTab[195493]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:954
					// _ = "end of CoverTab[195493]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:955
					_go_fuzz_dep_.CoverTab[195494]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:955
					// _ = "end of CoverTab[195494]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:955
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:955
				// _ = "end of CoverTab[195492]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:956
				_go_fuzz_dep_.CoverTab[195495]++
														return errors.New(formatErrorDescription(
					Locale.MustBeOfAn(),
					ErrorDetails{"x": KEY_ELSE, "y": TYPE_OBJECT},
				))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:960
				// _ = "end of CoverTab[195495]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:961
			// _ = "end of CoverTab[195491]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:962
			_go_fuzz_dep_.CoverTab[195496]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:962
			// _ = "end of CoverTab[195496]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:962
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:962
		// _ = "end of CoverTab[195478]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:963
		_go_fuzz_dep_.CoverTab[195497]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:963
		// _ = "end of CoverTab[195497]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:963
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:963
	// _ = "end of CoverTab[195135]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:963
	_go_fuzz_dep_.CoverTab[195136]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:965
	// _ = "end of CoverTab[195136]"
}

func (d *Schema) parseReference(documentNode interface{}, currentSchema *subSchema) error {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:968
	_go_fuzz_dep_.CoverTab[195498]++
											var (
		refdDocumentNode	interface{}
		dsp			*schemaPoolDocument
		err			error
	)

	newSchema := &subSchema{property: KEY_REF, parent: currentSchema, ref: currentSchema.ref}

	d.referencePool.Add(currentSchema.ref.String(), newSchema)

	dsp, err = d.pool.GetDocument(*currentSchema.ref)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:980
		_go_fuzz_dep_.CoverTab[195503]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:981
		// _ = "end of CoverTab[195503]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:982
		_go_fuzz_dep_.CoverTab[195504]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:982
		// _ = "end of CoverTab[195504]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:982
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:982
	// _ = "end of CoverTab[195498]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:982
	_go_fuzz_dep_.CoverTab[195499]++
											newSchema.id = currentSchema.ref

											refdDocumentNode = dsp.Document
											newSchema.draft = dsp.Draft

											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:988
		_go_fuzz_dep_.CoverTab[195505]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:989
		// _ = "end of CoverTab[195505]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:990
		_go_fuzz_dep_.CoverTab[195506]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:990
		// _ = "end of CoverTab[195506]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:990
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:990
	// _ = "end of CoverTab[195499]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:990
	_go_fuzz_dep_.CoverTab[195500]++

											if !isKind(refdDocumentNode, reflect.Map, reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:992
		_go_fuzz_dep_.CoverTab[195507]++
												return errors.New(formatErrorDescription(
			Locale.MustBeOfType(),
			ErrorDetails{"key": STRING_SCHEMA, "type": TYPE_OBJECT},
		))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:996
		// _ = "end of CoverTab[195507]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:997
		_go_fuzz_dep_.CoverTab[195508]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:997
		// _ = "end of CoverTab[195508]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:997
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:997
	// _ = "end of CoverTab[195500]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:997
	_go_fuzz_dep_.CoverTab[195501]++

											err = d.parseSchema(refdDocumentNode, newSchema)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1000
		_go_fuzz_dep_.CoverTab[195509]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1001
		// _ = "end of CoverTab[195509]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1002
		_go_fuzz_dep_.CoverTab[195510]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1002
		// _ = "end of CoverTab[195510]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1002
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1002
	// _ = "end of CoverTab[195501]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1002
	_go_fuzz_dep_.CoverTab[195502]++

												currentSchema.refSchema = newSchema

												return nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1006
	// _ = "end of CoverTab[195502]"

}

func (d *Schema) parseProperties(documentNode interface{}, currentSchema *subSchema) error {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1010
	_go_fuzz_dep_.CoverTab[195511]++

												if !isKind(documentNode, reflect.Map) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1012
		_go_fuzz_dep_.CoverTab[195514]++
													return errors.New(formatErrorDescription(
			Locale.MustBeOfType(),
			ErrorDetails{"key": STRING_PROPERTIES, "type": TYPE_OBJECT},
		))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1016
		// _ = "end of CoverTab[195514]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1017
		_go_fuzz_dep_.CoverTab[195515]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1017
		// _ = "end of CoverTab[195515]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1017
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1017
	// _ = "end of CoverTab[195511]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1017
	_go_fuzz_dep_.CoverTab[195512]++

												m := documentNode.(map[string]interface{})
												for k := range m {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1020
		_go_fuzz_dep_.CoverTab[195516]++
													schemaProperty := k
													newSchema := &subSchema{property: schemaProperty, parent: currentSchema, ref: currentSchema.ref}
													currentSchema.propertiesChildren = append(currentSchema.propertiesChildren, newSchema)
													err := d.parseSchema(m[k], newSchema)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1025
			_go_fuzz_dep_.CoverTab[195517]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1026
			// _ = "end of CoverTab[195517]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1027
			_go_fuzz_dep_.CoverTab[195518]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1027
			// _ = "end of CoverTab[195518]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1027
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1027
		// _ = "end of CoverTab[195516]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1028
	// _ = "end of CoverTab[195512]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1028
	_go_fuzz_dep_.CoverTab[195513]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1030
	// _ = "end of CoverTab[195513]"
}

func (d *Schema) parseDependencies(documentNode interface{}, currentSchema *subSchema) error {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1033
	_go_fuzz_dep_.CoverTab[195519]++

												if !isKind(documentNode, reflect.Map) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1035
		_go_fuzz_dep_.CoverTab[195522]++
													return errors.New(formatErrorDescription(
			Locale.MustBeOfType(),
			ErrorDetails{"key": KEY_DEPENDENCIES, "type": TYPE_OBJECT},
		))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1039
		// _ = "end of CoverTab[195522]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1040
		_go_fuzz_dep_.CoverTab[195523]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1040
		// _ = "end of CoverTab[195523]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1040
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1040
	// _ = "end of CoverTab[195519]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1040
	_go_fuzz_dep_.CoverTab[195520]++

												m := documentNode.(map[string]interface{})
												currentSchema.dependencies = make(map[string]interface{})

												for k := range m {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1045
		_go_fuzz_dep_.CoverTab[195524]++
													switch reflect.ValueOf(m[k]).Kind() {

		case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1048
			_go_fuzz_dep_.CoverTab[195525]++
														values := m[k].([]interface{})
														var valuesToRegister []string

														for _, value := range values {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1052
				_go_fuzz_dep_.CoverTab[195529]++
															if !isKind(value, reflect.String) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1053
					_go_fuzz_dep_.CoverTab[195531]++
																return errors.New(formatErrorDescription(
						Locale.MustBeOfType(),
						ErrorDetails{
							"key":	STRING_DEPENDENCY,
							"type":	STRING_SCHEMA_OR_ARRAY_OF_STRINGS,
						},
					))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1060
					// _ = "end of CoverTab[195531]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1061
					_go_fuzz_dep_.CoverTab[195532]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1061
					// _ = "end of CoverTab[195532]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1061
				}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1061
				// _ = "end of CoverTab[195529]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1061
				_go_fuzz_dep_.CoverTab[195530]++
															valuesToRegister = append(valuesToRegister, value.(string))
															currentSchema.dependencies[k] = valuesToRegister
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1063
				// _ = "end of CoverTab[195530]"
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1064
			// _ = "end of CoverTab[195525]"

		case reflect.Map, reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1066
			_go_fuzz_dep_.CoverTab[195526]++
														depSchema := &subSchema{property: k, parent: currentSchema, ref: currentSchema.ref}
														err := d.parseSchema(m[k], depSchema)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1069
				_go_fuzz_dep_.CoverTab[195533]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1070
				// _ = "end of CoverTab[195533]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1071
				_go_fuzz_dep_.CoverTab[195534]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1071
				// _ = "end of CoverTab[195534]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1071
			}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1071
			// _ = "end of CoverTab[195526]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1071
			_go_fuzz_dep_.CoverTab[195527]++
														currentSchema.dependencies[k] = depSchema
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1072
			// _ = "end of CoverTab[195527]"

		default:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1074
			_go_fuzz_dep_.CoverTab[195528]++
														return errors.New(formatErrorDescription(
				Locale.MustBeOfType(),
				ErrorDetails{
					"key":	STRING_DEPENDENCY,
					"type":	STRING_SCHEMA_OR_ARRAY_OF_STRINGS,
				},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1081
			// _ = "end of CoverTab[195528]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1082
		// _ = "end of CoverTab[195524]"

	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1084
	// _ = "end of CoverTab[195520]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1084
	_go_fuzz_dep_.CoverTab[195521]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1086
	// _ = "end of CoverTab[195521]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1087
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schema.go:1087
var _ = _go_fuzz_dep_.CoverTab
