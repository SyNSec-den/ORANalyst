// Copyright 2013 MongoDB, Inc.
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

// author           tolsen
// author-github    https://github.com/tolsen
//
// repository-name  gojsonschema
// repository-desc  An implementation of JSON Schema, based on IETF's draft v4 - Go language.
//
// description      Implements a persistent (immutable w/ shared structure) singly-linked list of strings for the purpose of storing a json context
//
// created          04-09-2013

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:25
package gojsonschema

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:25
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:25
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:25
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:25
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:25
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:25
)

import "bytes"

// JsonContext implements a persistent linked-list of strings
type JsonContext struct {
	head	string
	tail	*JsonContext
}

// NewJsonContext creates a new JsonContext
func NewJsonContext(head string, tail *JsonContext) *JsonContext {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:36
	_go_fuzz_dep_.CoverTab[194912]++
												return &JsonContext{head, tail}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:37
	// _ = "end of CoverTab[194912]"
}

// String displays the context in reverse.
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:40
// This plays well with the data structure's persistent nature with
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:40
// Cons and a json document's tree structure.
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:43
func (c *JsonContext) String(del ...string) string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:43
	_go_fuzz_dep_.CoverTab[194913]++
												byteArr := make([]byte, 0, c.stringLen())
												buf := bytes.NewBuffer(byteArr)
												c.writeStringToBuffer(buf, del)

												return buf.String()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:48
	// _ = "end of CoverTab[194913]"
}

func (c *JsonContext) stringLen() int {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:51
	_go_fuzz_dep_.CoverTab[194914]++
												length := 0
												if c.tail != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:53
		_go_fuzz_dep_.CoverTab[194916]++
													length = c.tail.stringLen() + 1
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:54
		// _ = "end of CoverTab[194916]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:55
		_go_fuzz_dep_.CoverTab[194917]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:55
		// _ = "end of CoverTab[194917]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:55
	// _ = "end of CoverTab[194914]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:55
	_go_fuzz_dep_.CoverTab[194915]++

												length += len(c.head)
												return length
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:58
	// _ = "end of CoverTab[194915]"
}

func (c *JsonContext) writeStringToBuffer(buf *bytes.Buffer, del []string) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:61
	_go_fuzz_dep_.CoverTab[194918]++
												if c.tail != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:62
		_go_fuzz_dep_.CoverTab[194920]++
													c.tail.writeStringToBuffer(buf, del)

													if len(del) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:65
			_go_fuzz_dep_.CoverTab[194921]++
														buf.WriteString(del[0])
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:66
			// _ = "end of CoverTab[194921]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:67
			_go_fuzz_dep_.CoverTab[194922]++
														buf.WriteString(".")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:68
			// _ = "end of CoverTab[194922]"
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:69
		// _ = "end of CoverTab[194920]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:70
		_go_fuzz_dep_.CoverTab[194923]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:70
		// _ = "end of CoverTab[194923]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:70
	// _ = "end of CoverTab[194918]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:70
	_go_fuzz_dep_.CoverTab[194919]++

												buf.WriteString(c.head)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:72
	// _ = "end of CoverTab[194919]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:73
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/jsonContext.go:73
var _ = _go_fuzz_dep_.CoverTab
