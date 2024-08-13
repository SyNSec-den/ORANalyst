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
// description      Helper structure to handle schema types, and the combination of them.
//
// created          28-02-2013

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:26
package gojsonschema

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:26
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:26
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:26
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:26
)

import (
	"errors"
	"fmt"
	"strings"
)

type jsonSchemaType struct {
	types []string
}

// Is the schema typed ? that is containing at least one type
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:38
// When not typed, the schema does not need any type validation
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:40
func (t *jsonSchemaType) IsTyped() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:40
	_go_fuzz_dep_.CoverTab[195709]++
												return len(t.types) > 0
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:41
	// _ = "end of CoverTab[195709]"
}

func (t *jsonSchemaType) Add(etype string) error {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:44
	_go_fuzz_dep_.CoverTab[195710]++

												if !isStringInSlice(JSON_TYPES, etype) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:46
		_go_fuzz_dep_.CoverTab[195713]++
													return errors.New(formatErrorDescription(Locale.NotAValidType(), ErrorDetails{"given": "/" + etype + "/", "expected": JSON_TYPES}))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:47
		// _ = "end of CoverTab[195713]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:48
		_go_fuzz_dep_.CoverTab[195714]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:48
		// _ = "end of CoverTab[195714]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:48
	// _ = "end of CoverTab[195710]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:48
	_go_fuzz_dep_.CoverTab[195711]++

												if t.Contains(etype) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:50
		_go_fuzz_dep_.CoverTab[195715]++
													return errors.New(formatErrorDescription(Locale.Duplicated(), ErrorDetails{"type": etype}))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:51
		// _ = "end of CoverTab[195715]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:52
		_go_fuzz_dep_.CoverTab[195716]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:52
		// _ = "end of CoverTab[195716]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:52
	// _ = "end of CoverTab[195711]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:52
	_go_fuzz_dep_.CoverTab[195712]++

												t.types = append(t.types, etype)

												return nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:56
	// _ = "end of CoverTab[195712]"
}

func (t *jsonSchemaType) Contains(etype string) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:59
	_go_fuzz_dep_.CoverTab[195717]++

												for _, v := range t.types {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:61
		_go_fuzz_dep_.CoverTab[195719]++
													if v == etype {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:62
			_go_fuzz_dep_.CoverTab[195720]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:63
			// _ = "end of CoverTab[195720]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:64
			_go_fuzz_dep_.CoverTab[195721]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:64
			// _ = "end of CoverTab[195721]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:64
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:64
		// _ = "end of CoverTab[195719]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:65
	// _ = "end of CoverTab[195717]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:65
	_go_fuzz_dep_.CoverTab[195718]++

												return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:67
	// _ = "end of CoverTab[195718]"
}

func (t *jsonSchemaType) String() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:70
	_go_fuzz_dep_.CoverTab[195722]++

												if len(t.types) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:72
		_go_fuzz_dep_.CoverTab[195725]++
													return STRING_UNDEFINED
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:73
		// _ = "end of CoverTab[195725]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:74
		_go_fuzz_dep_.CoverTab[195726]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:74
		// _ = "end of CoverTab[195726]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:74
	// _ = "end of CoverTab[195722]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:74
	_go_fuzz_dep_.CoverTab[195723]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:77
	if len(t.types) > 1 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:77
		_go_fuzz_dep_.CoverTab[195727]++
													return fmt.Sprintf("[%s]", strings.Join(t.types, ","))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:78
		// _ = "end of CoverTab[195727]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:79
		_go_fuzz_dep_.CoverTab[195728]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:79
		// _ = "end of CoverTab[195728]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:79
	// _ = "end of CoverTab[195723]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:79
	_go_fuzz_dep_.CoverTab[195724]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:82
	return t.types[0]
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:82
	// _ = "end of CoverTab[195724]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:83
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaType.go:83
var _ = _go_fuzz_dep_.CoverTab
