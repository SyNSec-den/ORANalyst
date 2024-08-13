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
// description      Pool of referenced schemas.
//
// created          25-06-2013

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:26
package gojsonschema

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:26
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:26
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:26
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:26
)

import (
	"fmt"
)

type schemaReferencePool struct {
	documents map[string]*subSchema
}

func newSchemaReferencePool() *schemaReferencePool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:36
	_go_fuzz_dep_.CoverTab[195692]++

													p := &schemaReferencePool{}
													p.documents = make(map[string]*subSchema)

													return p
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:41
	// _ = "end of CoverTab[195692]"
}

func (p *schemaReferencePool) Get(ref string) (r *subSchema, o bool) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:44
	_go_fuzz_dep_.CoverTab[195693]++

													if internalLogEnabled {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:46
		_go_fuzz_dep_.CoverTab[195696]++
														internalLog(fmt.Sprintf("Schema Reference ( %s )", ref))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:47
		// _ = "end of CoverTab[195696]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:48
		_go_fuzz_dep_.CoverTab[195697]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:48
		// _ = "end of CoverTab[195697]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:48
	// _ = "end of CoverTab[195693]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:48
	_go_fuzz_dep_.CoverTab[195694]++

													if sch, ok := p.documents[ref]; ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:50
		_go_fuzz_dep_.CoverTab[195698]++
														if internalLogEnabled {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:51
			_go_fuzz_dep_.CoverTab[195700]++
															internalLog(fmt.Sprintf(" From pool"))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:52
			// _ = "end of CoverTab[195700]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:53
			_go_fuzz_dep_.CoverTab[195701]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:53
			// _ = "end of CoverTab[195701]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:53
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:53
		// _ = "end of CoverTab[195698]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:53
		_go_fuzz_dep_.CoverTab[195699]++
														return sch, true
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:54
		// _ = "end of CoverTab[195699]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:55
		_go_fuzz_dep_.CoverTab[195702]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:55
		// _ = "end of CoverTab[195702]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:55
	// _ = "end of CoverTab[195694]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:55
	_go_fuzz_dep_.CoverTab[195695]++

													return nil, false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:57
	// _ = "end of CoverTab[195695]"
}

func (p *schemaReferencePool) Add(ref string, sch *subSchema) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:60
	_go_fuzz_dep_.CoverTab[195703]++

													if internalLogEnabled {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:62
		_go_fuzz_dep_.CoverTab[195705]++
														internalLog(fmt.Sprintf("Add Schema Reference %s to pool", ref))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:63
		// _ = "end of CoverTab[195705]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:64
		_go_fuzz_dep_.CoverTab[195706]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:64
		// _ = "end of CoverTab[195706]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:64
	// _ = "end of CoverTab[195703]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:64
	_go_fuzz_dep_.CoverTab[195704]++
													if _, ok := p.documents[ref]; !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:65
		_go_fuzz_dep_.CoverTab[195707]++
														p.documents[ref] = sch
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:66
		// _ = "end of CoverTab[195707]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:67
		_go_fuzz_dep_.CoverTab[195708]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:67
		// _ = "end of CoverTab[195708]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:67
	// _ = "end of CoverTab[195704]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:68
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/schemaReferencePool.go:68
var _ = _go_fuzz_dep_.CoverTab
