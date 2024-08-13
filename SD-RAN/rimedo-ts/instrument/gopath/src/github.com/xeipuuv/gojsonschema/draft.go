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

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:15
package gojsonschema

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:15
)

import (
	"errors"
	"math"
	"reflect"

	"github.com/xeipuuv/gojsonreference"
)

// Draft is a JSON-schema draft version
type Draft int

// Supported Draft versions
const (
	Draft4	Draft	= 4
	Draft6	Draft	= 6
	Draft7	Draft	= 7
	Hybrid	Draft	= math.MaxInt32
)

type draftConfig struct {
	Version		Draft
	MetaSchemaURL	string
	MetaSchema	string
}
type draftConfigs []draftConfig

var drafts draftConfigs

func init() {
	drafts = []draftConfig{
		{
			Version:	Draft4,
			MetaSchemaURL:	"http://json-schema.org/draft-04/schema",
			MetaSchema:	`{"id":"http://json-schema.org/draft-04/schema#","$schema":"http://json-schema.org/draft-04/schema#","description":"Core schema meta-schema","definitions":{"schemaArray":{"type":"array","minItems":1,"items":{"$ref":"#"}},"positiveInteger":{"type":"integer","minimum":0},"positiveIntegerDefault0":{"allOf":[{"$ref":"#/definitions/positiveInteger"},{"default":0}]},"simpleTypes":{"enum":["array","boolean","integer","null","number","object","string"]},"stringArray":{"type":"array","items":{"type":"string"},"minItems":1,"uniqueItems":true}},"type":"object","properties":{"id":{"type":"string"},"$schema":{"type":"string"},"title":{"type":"string"},"description":{"type":"string"},"default":{},"multipleOf":{"type":"number","minimum":0,"exclusiveMinimum":true},"maximum":{"type":"number"},"exclusiveMaximum":{"type":"boolean","default":false},"minimum":{"type":"number"},"exclusiveMinimum":{"type":"boolean","default":false},"maxLength":{"$ref":"#/definitions/positiveInteger"},"minLength":{"$ref":"#/definitions/positiveIntegerDefault0"},"pattern":{"type":"string","format":"regex"},"additionalItems":{"anyOf":[{"type":"boolean"},{"$ref":"#"}],"default":{}},"items":{"anyOf":[{"$ref":"#"},{"$ref":"#/definitions/schemaArray"}],"default":{}},"maxItems":{"$ref":"#/definitions/positiveInteger"},"minItems":{"$ref":"#/definitions/positiveIntegerDefault0"},"uniqueItems":{"type":"boolean","default":false},"maxProperties":{"$ref":"#/definitions/positiveInteger"},"minProperties":{"$ref":"#/definitions/positiveIntegerDefault0"},"required":{"$ref":"#/definitions/stringArray"},"additionalProperties":{"anyOf":[{"type":"boolean"},{"$ref":"#"}],"default":{}},"definitions":{"type":"object","additionalProperties":{"$ref":"#"},"default":{}},"properties":{"type":"object","additionalProperties":{"$ref":"#"},"default":{}},"patternProperties":{"type":"object","additionalProperties":{"$ref":"#"},"default":{}},"dependencies":{"type":"object","additionalProperties":{"anyOf":[{"$ref":"#"},{"$ref":"#/definitions/stringArray"}]}},"enum":{"type":"array","minItems":1,"uniqueItems":true},"type":{"anyOf":[{"$ref":"#/definitions/simpleTypes"},{"type":"array","items":{"$ref":"#/definitions/simpleTypes"},"minItems":1,"uniqueItems":true}]},"format":{"type":"string"},"allOf":{"$ref":"#/definitions/schemaArray"},"anyOf":{"$ref":"#/definitions/schemaArray"},"oneOf":{"$ref":"#/definitions/schemaArray"},"not":{"$ref":"#"}},"dependencies":{"exclusiveMaximum":["maximum"],"exclusiveMinimum":["minimum"]},"default":{}}`,
		},
		{
			Version:	Draft6,
			MetaSchemaURL:	"http://json-schema.org/draft-06/schema",
			MetaSchema:	`{"$schema":"http://json-schema.org/draft-06/schema#","$id":"http://json-schema.org/draft-06/schema#","title":"Core schema meta-schema","definitions":{"schemaArray":{"type":"array","minItems":1,"items":{"$ref":"#"}},"nonNegativeInteger":{"type":"integer","minimum":0},"nonNegativeIntegerDefault0":{"allOf":[{"$ref":"#/definitions/nonNegativeInteger"},{"default":0}]},"simpleTypes":{"enum":["array","boolean","integer","null","number","object","string"]},"stringArray":{"type":"array","items":{"type":"string"},"uniqueItems":true,"default":[]}},"type":["object","boolean"],"properties":{"$id":{"type":"string","format":"uri-reference"},"$schema":{"type":"string","format":"uri"},"$ref":{"type":"string","format":"uri-reference"},"title":{"type":"string"},"description":{"type":"string"},"default":{},"examples":{"type":"array","items":{}},"multipleOf":{"type":"number","exclusiveMinimum":0},"maximum":{"type":"number"},"exclusiveMaximum":{"type":"number"},"minimum":{"type":"number"},"exclusiveMinimum":{"type":"number"},"maxLength":{"$ref":"#/definitions/nonNegativeInteger"},"minLength":{"$ref":"#/definitions/nonNegativeIntegerDefault0"},"pattern":{"type":"string","format":"regex"},"additionalItems":{"$ref":"#"},"items":{"anyOf":[{"$ref":"#"},{"$ref":"#/definitions/schemaArray"}],"default":{}},"maxItems":{"$ref":"#/definitions/nonNegativeInteger"},"minItems":{"$ref":"#/definitions/nonNegativeIntegerDefault0"},"uniqueItems":{"type":"boolean","default":false},"contains":{"$ref":"#"},"maxProperties":{"$ref":"#/definitions/nonNegativeInteger"},"minProperties":{"$ref":"#/definitions/nonNegativeIntegerDefault0"},"required":{"$ref":"#/definitions/stringArray"},"additionalProperties":{"$ref":"#"},"definitions":{"type":"object","additionalProperties":{"$ref":"#"},"default":{}},"properties":{"type":"object","additionalProperties":{"$ref":"#"},"default":{}},"patternProperties":{"type":"object","additionalProperties":{"$ref":"#"},"default":{}},"dependencies":{"type":"object","additionalProperties":{"anyOf":[{"$ref":"#"},{"$ref":"#/definitions/stringArray"}]}},"propertyNames":{"$ref":"#"},"const":{},"enum":{"type":"array","minItems":1,"uniqueItems":true},"type":{"anyOf":[{"$ref":"#/definitions/simpleTypes"},{"type":"array","items":{"$ref":"#/definitions/simpleTypes"},"minItems":1,"uniqueItems":true}]},"format":{"type":"string"},"allOf":{"$ref":"#/definitions/schemaArray"},"anyOf":{"$ref":"#/definitions/schemaArray"},"oneOf":{"$ref":"#/definitions/schemaArray"},"not":{"$ref":"#"}},"default":{}}`,
		},
		{
			Version:	Draft7,
			MetaSchemaURL:	"http://json-schema.org/draft-07/schema",
			MetaSchema:	`{"$schema":"http://json-schema.org/draft-07/schema#","$id":"http://json-schema.org/draft-07/schema#","title":"Core schema meta-schema","definitions":{"schemaArray":{"type":"array","minItems":1,"items":{"$ref":"#"}},"nonNegativeInteger":{"type":"integer","minimum":0},"nonNegativeIntegerDefault0":{"allOf":[{"$ref":"#/definitions/nonNegativeInteger"},{"default":0}]},"simpleTypes":{"enum":["array","boolean","integer","null","number","object","string"]},"stringArray":{"type":"array","items":{"type":"string"},"uniqueItems":true,"default":[]}},"type":["object","boolean"],"properties":{"$id":{"type":"string","format":"uri-reference"},"$schema":{"type":"string","format":"uri"},"$ref":{"type":"string","format":"uri-reference"},"$comment":{"type":"string"},"title":{"type":"string"},"description":{"type":"string"},"default":true,"readOnly":{"type":"boolean","default":false},"examples":{"type":"array","items":true},"multipleOf":{"type":"number","exclusiveMinimum":0},"maximum":{"type":"number"},"exclusiveMaximum":{"type":"number"},"minimum":{"type":"number"},"exclusiveMinimum":{"type":"number"},"maxLength":{"$ref":"#/definitions/nonNegativeInteger"},"minLength":{"$ref":"#/definitions/nonNegativeIntegerDefault0"},"pattern":{"type":"string","format":"regex"},"additionalItems":{"$ref":"#"},"items":{"anyOf":[{"$ref":"#"},{"$ref":"#/definitions/schemaArray"}],"default":true},"maxItems":{"$ref":"#/definitions/nonNegativeInteger"},"minItems":{"$ref":"#/definitions/nonNegativeIntegerDefault0"},"uniqueItems":{"type":"boolean","default":false},"contains":{"$ref":"#"},"maxProperties":{"$ref":"#/definitions/nonNegativeInteger"},"minProperties":{"$ref":"#/definitions/nonNegativeIntegerDefault0"},"required":{"$ref":"#/definitions/stringArray"},"additionalProperties":{"$ref":"#"},"definitions":{"type":"object","additionalProperties":{"$ref":"#"},"default":{}},"properties":{"type":"object","additionalProperties":{"$ref":"#"},"default":{}},"patternProperties":{"type":"object","additionalProperties":{"$ref":"#"},"propertyNames":{"format":"regex"},"default":{}},"dependencies":{"type":"object","additionalProperties":{"anyOf":[{"$ref":"#"},{"$ref":"#/definitions/stringArray"}]}},"propertyNames":{"$ref":"#"},"const":true,"enum":{"type":"array","items":true,"minItems":1,"uniqueItems":true},"type":{"anyOf":[{"$ref":"#/definitions/simpleTypes"},{"type":"array","items":{"$ref":"#/definitions/simpleTypes"},"minItems":1,"uniqueItems":true}]},"format":{"type":"string"},"contentMediaType":{"type":"string"},"contentEncoding":{"type":"string"},"if":{"$ref":"#"},"then":{"$ref":"#"},"else":{"$ref":"#"},"allOf":{"$ref":"#/definitions/schemaArray"},"anyOf":{"$ref":"#/definitions/schemaArray"},"oneOf":{"$ref":"#/definitions/schemaArray"},"not":{"$ref":"#"}},"default":true}`,
		},
	}
}

func (dc draftConfigs) GetMetaSchema(url string) string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:65
	_go_fuzz_dep_.CoverTab[194745]++
											for _, config := range dc {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:66
		_go_fuzz_dep_.CoverTab[194747]++
												if config.MetaSchemaURL == url {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:67
			_go_fuzz_dep_.CoverTab[194748]++
													return config.MetaSchema
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:68
			// _ = "end of CoverTab[194748]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:69
			_go_fuzz_dep_.CoverTab[194749]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:69
			// _ = "end of CoverTab[194749]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:69
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:69
		// _ = "end of CoverTab[194747]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:70
	// _ = "end of CoverTab[194745]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:70
	_go_fuzz_dep_.CoverTab[194746]++
											return ""
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:71
	// _ = "end of CoverTab[194746]"
}
func (dc draftConfigs) GetDraftVersion(url string) *Draft {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:73
	_go_fuzz_dep_.CoverTab[194750]++
											for _, config := range dc {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:74
		_go_fuzz_dep_.CoverTab[194752]++
												if config.MetaSchemaURL == url {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:75
			_go_fuzz_dep_.CoverTab[194753]++
													return &config.Version
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:76
			// _ = "end of CoverTab[194753]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:77
			_go_fuzz_dep_.CoverTab[194754]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:77
			// _ = "end of CoverTab[194754]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:77
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:77
		// _ = "end of CoverTab[194752]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:78
	// _ = "end of CoverTab[194750]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:78
	_go_fuzz_dep_.CoverTab[194751]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:79
	// _ = "end of CoverTab[194751]"
}
func (dc draftConfigs) GetSchemaURL(draft Draft) string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:81
	_go_fuzz_dep_.CoverTab[194755]++
											for _, config := range dc {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:82
		_go_fuzz_dep_.CoverTab[194757]++
												if config.Version == draft {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:83
			_go_fuzz_dep_.CoverTab[194758]++
													return config.MetaSchemaURL
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:84
			// _ = "end of CoverTab[194758]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:85
			_go_fuzz_dep_.CoverTab[194759]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:85
			// _ = "end of CoverTab[194759]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:85
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:85
		// _ = "end of CoverTab[194757]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:86
	// _ = "end of CoverTab[194755]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:86
	_go_fuzz_dep_.CoverTab[194756]++
											return ""
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:87
	// _ = "end of CoverTab[194756]"
}

func parseSchemaURL(documentNode interface{}) (string, *Draft, error) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:90
	_go_fuzz_dep_.CoverTab[194760]++

											if isKind(documentNode, reflect.Bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:92
		_go_fuzz_dep_.CoverTab[194764]++
												return "", nil, nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:93
		// _ = "end of CoverTab[194764]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:94
		_go_fuzz_dep_.CoverTab[194765]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:94
		// _ = "end of CoverTab[194765]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:94
	// _ = "end of CoverTab[194760]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:94
	_go_fuzz_dep_.CoverTab[194761]++

											if !isKind(documentNode, reflect.Map) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:96
		_go_fuzz_dep_.CoverTab[194766]++
												return "", nil, errors.New("schema is invalid")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:97
		// _ = "end of CoverTab[194766]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:98
		_go_fuzz_dep_.CoverTab[194767]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:98
		// _ = "end of CoverTab[194767]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:98
	// _ = "end of CoverTab[194761]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:98
	_go_fuzz_dep_.CoverTab[194762]++

											m := documentNode.(map[string]interface{})

											if existsMapKey(m, KEY_SCHEMA) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:102
		_go_fuzz_dep_.CoverTab[194768]++
												if !isKind(m[KEY_SCHEMA], reflect.String) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:103
			_go_fuzz_dep_.CoverTab[194771]++
													return "", nil, errors.New(formatErrorDescription(
				Locale.MustBeOfType(),
				ErrorDetails{
					"key":	KEY_SCHEMA,
					"type":	TYPE_STRING,
				},
			))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:110
			// _ = "end of CoverTab[194771]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:111
			_go_fuzz_dep_.CoverTab[194772]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:111
			// _ = "end of CoverTab[194772]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:111
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:111
		// _ = "end of CoverTab[194768]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:111
		_go_fuzz_dep_.CoverTab[194769]++

												schemaReference, err := gojsonreference.NewJsonReference(m[KEY_SCHEMA].(string))

												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:115
			_go_fuzz_dep_.CoverTab[194773]++
													return "", nil, err
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:116
			// _ = "end of CoverTab[194773]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:117
			_go_fuzz_dep_.CoverTab[194774]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:117
			// _ = "end of CoverTab[194774]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:117
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:117
		// _ = "end of CoverTab[194769]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:117
		_go_fuzz_dep_.CoverTab[194770]++

												schema := schemaReference.String()

												return schema, drafts.GetDraftVersion(schema), nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:121
		// _ = "end of CoverTab[194770]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:122
		_go_fuzz_dep_.CoverTab[194775]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:122
		// _ = "end of CoverTab[194775]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:122
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:122
	// _ = "end of CoverTab[194762]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:122
	_go_fuzz_dep_.CoverTab[194763]++

											return "", nil, nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:124
	// _ = "end of CoverTab[194763]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:125
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/draft.go:125
var _ = _go_fuzz_dep_.CoverTab
