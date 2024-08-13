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
// description      Very simple log wrapper.
//					Used for debugging/testing purposes.
//
// created          01-01-2015

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/internalLog.go:27
package gojsonschema

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/internalLog.go:27
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/internalLog.go:27
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/internalLog.go:27
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/internalLog.go:27
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/internalLog.go:27
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/internalLog.go:27
)

import (
	"log"
)

const internalLogEnabled = false

func internalLog(format string, v ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/internalLog.go:35
	_go_fuzz_dep_.CoverTab[194911]++
												log.Printf(format, v...)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/internalLog.go:36
	// _ = "end of CoverTab[194911]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/internalLog.go:37
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/internalLog.go:37
var _ = _go_fuzz_dep_.CoverTab
