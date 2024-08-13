// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2018, The GoGo Authors. All rights reserved.
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

// +build !purego,!appengine,!js

// This file contains the implementation of the proto field accesses using package unsafe.

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe_gogo.go:33
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe_gogo.go:33
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe_gogo.go:33
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe_gogo.go:33
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe_gogo.go:33
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe_gogo.go:33
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe_gogo.go:33
)

import (
	"reflect"
	"unsafe"
)

func (p pointer) getRef() pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe_gogo.go:40
	_go_fuzz_dep_.CoverTab[108907]++
													return pointer{p: (unsafe.Pointer)(&p.p)}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe_gogo.go:41
	// _ = "end of CoverTab[108907]"
}

func (p pointer) appendRef(v pointer, typ reflect.Type) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe_gogo.go:44
	_go_fuzz_dep_.CoverTab[108908]++
													slice := p.getSlice(typ)
													elem := v.asPointerTo(typ).Elem()
													newSlice := reflect.Append(slice, elem)
													slice.Set(newSlice)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe_gogo.go:48
	// _ = "end of CoverTab[108908]"
}

func (p pointer) getSlice(typ reflect.Type) reflect.Value {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe_gogo.go:51
	_go_fuzz_dep_.CoverTab[108909]++
													sliceTyp := reflect.SliceOf(typ)
													slice := p.asPointerTo(sliceTyp)
													slice = slice.Elem()
													return slice
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe_gogo.go:55
	// _ = "end of CoverTab[108909]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe_gogo.go:56
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe_gogo.go:56
var _ = _go_fuzz_dep_.CoverTab
