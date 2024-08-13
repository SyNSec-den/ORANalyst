// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:29
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:29
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:29
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:29
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:29
)

import (
	"fmt"
	"reflect"
)

func (tm *TextMarshaler) writeEnum(w *textWriter, v reflect.Value, props *Properties) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:36
	_go_fuzz_dep_.CoverTab[112984]++
												m, ok := enumStringMaps[props.Enum]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:38
		_go_fuzz_dep_.CoverTab[112988]++
													if err := tm.writeAny(w, v, props); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:39
			_go_fuzz_dep_.CoverTab[112989]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:40
			// _ = "end of CoverTab[112989]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:41
			_go_fuzz_dep_.CoverTab[112990]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:41
			// _ = "end of CoverTab[112990]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:41
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:41
		// _ = "end of CoverTab[112988]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:42
		_go_fuzz_dep_.CoverTab[112991]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:42
		// _ = "end of CoverTab[112991]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:42
	// _ = "end of CoverTab[112984]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:42
	_go_fuzz_dep_.CoverTab[112985]++
												key := int32(0)
												if v.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:44
		_go_fuzz_dep_.CoverTab[112992]++
													key = int32(v.Elem().Int())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:45
		// _ = "end of CoverTab[112992]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:46
		_go_fuzz_dep_.CoverTab[112993]++
													key = int32(v.Int())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:47
		// _ = "end of CoverTab[112993]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:48
	// _ = "end of CoverTab[112985]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:48
	_go_fuzz_dep_.CoverTab[112986]++
												s, ok := m[key]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:50
		_go_fuzz_dep_.CoverTab[112994]++
													if err := tm.writeAny(w, v, props); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:51
			_go_fuzz_dep_.CoverTab[112995]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:52
			// _ = "end of CoverTab[112995]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:53
			_go_fuzz_dep_.CoverTab[112996]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:53
			// _ = "end of CoverTab[112996]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:53
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:53
		// _ = "end of CoverTab[112994]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:54
		_go_fuzz_dep_.CoverTab[112997]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:54
		// _ = "end of CoverTab[112997]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:54
	// _ = "end of CoverTab[112986]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:54
	_go_fuzz_dep_.CoverTab[112987]++
												_, err := fmt.Fprint(w, s)
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:56
	// _ = "end of CoverTab[112987]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:57
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/text_gogo.go:57
var _ = _go_fuzz_dep_.CoverTab
