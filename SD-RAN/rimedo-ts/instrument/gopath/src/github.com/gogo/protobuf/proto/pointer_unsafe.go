// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2012 The Go Authors.  All rights reserved.
// https://github.com/golang/protobuf
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
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:36
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:36
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:36
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:36
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:36
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:36
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:36
)

import (
	"reflect"
	"sync/atomic"
	"unsafe"
)

const unsafeAllowed = true

// A field identifies a field in a struct, accessible from a pointer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:46
// In this implementation, a field is identified by its byte offset from the start of the struct.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:48
type field uintptr

// toField returns a field equivalent to the given reflect field.
func toField(f *reflect.StructField) field {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:51
	_go_fuzz_dep_.CoverTab[108851]++
												return field(f.Offset)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:52
	// _ = "end of CoverTab[108851]"
}

// invalidField is an invalid field identifier.
const invalidField = ^field(0)

// zeroField is a noop when calling pointer.offset.
const zeroField = field(0)

// IsValid reports whether the field identifier is valid.
func (f field) IsValid() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:62
	_go_fuzz_dep_.CoverTab[108852]++
												return f != invalidField
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:63
	// _ = "end of CoverTab[108852]"
}

// The pointer type below is for the new table-driven encoder/decoder.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:66
// The implementation here uses unsafe.Pointer to create a generic pointer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:66
// In pointer_reflect.go we use reflect instead of unsafe to implement
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:66
// the same (but slower) interface.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:70
type pointer struct {
	p unsafe.Pointer
}

// size of pointer
var ptrSize = unsafe.Sizeof(uintptr(0))

// toPointer converts an interface of pointer type to a pointer
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:77
// that points to the same target.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:79
func toPointer(i *Message) pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:79
	_go_fuzz_dep_.CoverTab[108853]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:83
	return pointer{p: (*[2]unsafe.Pointer)(unsafe.Pointer(i))[1]}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:83
	// _ = "end of CoverTab[108853]"
}

// toAddrPointer converts an interface to a pointer that points to
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:86
// the interface data.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:88
func toAddrPointer(i *interface{}, isptr bool) pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:88
	_go_fuzz_dep_.CoverTab[108854]++

												if isptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:90
		_go_fuzz_dep_.CoverTab[108856]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:93
		return pointer{p: unsafe.Pointer(uintptr(unsafe.Pointer(i)) + ptrSize)}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:93
		// _ = "end of CoverTab[108856]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:94
		_go_fuzz_dep_.CoverTab[108857]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:94
		// _ = "end of CoverTab[108857]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:94
	// _ = "end of CoverTab[108854]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:94
	_go_fuzz_dep_.CoverTab[108855]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:97
	return pointer{p: (*[2]unsafe.Pointer)(unsafe.Pointer(i))[1]}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:97
	// _ = "end of CoverTab[108855]"
}

// valToPointer converts v to a pointer. v must be of pointer type.
func valToPointer(v reflect.Value) pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:101
	_go_fuzz_dep_.CoverTab[108858]++
												return pointer{p: unsafe.Pointer(v.Pointer())}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:102
	// _ = "end of CoverTab[108858]"
}

// offset converts from a pointer to a structure to a pointer to
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:105
// one of its fields.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:107
func (p pointer) offset(f field) pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:107
	_go_fuzz_dep_.CoverTab[108859]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:115
	return pointer{p: unsafe.Pointer(uintptr(p.p) + uintptr(f))}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:115
	// _ = "end of CoverTab[108859]"
}

func (p pointer) isNil() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:118
	_go_fuzz_dep_.CoverTab[108860]++
												return p.p == nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:119
	// _ = "end of CoverTab[108860]"
}

func (p pointer) toInt64() *int64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:122
	_go_fuzz_dep_.CoverTab[108861]++
												return (*int64)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:123
	// _ = "end of CoverTab[108861]"
}
func (p pointer) toInt64Ptr() **int64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:125
	_go_fuzz_dep_.CoverTab[108862]++
												return (**int64)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:126
	// _ = "end of CoverTab[108862]"
}
func (p pointer) toInt64Slice() *[]int64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:128
	_go_fuzz_dep_.CoverTab[108863]++
												return (*[]int64)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:129
	// _ = "end of CoverTab[108863]"
}
func (p pointer) toInt32() *int32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:131
	_go_fuzz_dep_.CoverTab[108864]++
												return (*int32)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:132
	// _ = "end of CoverTab[108864]"
}

// See pointer_reflect.go for why toInt32Ptr/Slice doesn't exist.
/*
	func (p pointer) toInt32Ptr() **int32 {
		return (**int32)(p.p)
	}
	func (p pointer) toInt32Slice() *[]int32 {
		return (*[]int32)(p.p)
	}
*/
func (p pointer) getInt32Ptr() *int32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:144
	_go_fuzz_dep_.CoverTab[108865]++
												return *(**int32)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:145
	// _ = "end of CoverTab[108865]"
}
func (p pointer) setInt32Ptr(v int32) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:147
	_go_fuzz_dep_.CoverTab[108866]++
												*(**int32)(p.p) = &v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:148
	// _ = "end of CoverTab[108866]"
}

// getInt32Slice loads a []int32 from p.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:151
// The value returned is aliased with the original slice.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:151
// This behavior differs from the implementation in pointer_reflect.go.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:154
func (p pointer) getInt32Slice() []int32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:154
	_go_fuzz_dep_.CoverTab[108867]++
												return *(*[]int32)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:155
	// _ = "end of CoverTab[108867]"
}

// setInt32Slice stores a []int32 to p.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:158
// The value set is aliased with the input slice.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:158
// This behavior differs from the implementation in pointer_reflect.go.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:161
func (p pointer) setInt32Slice(v []int32) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:161
	_go_fuzz_dep_.CoverTab[108868]++
												*(*[]int32)(p.p) = v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:162
	// _ = "end of CoverTab[108868]"
}

// TODO: Can we get rid of appendInt32Slice and use setInt32Slice instead?
func (p pointer) appendInt32Slice(v int32) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:166
	_go_fuzz_dep_.CoverTab[108869]++
												s := (*[]int32)(p.p)
												*s = append(*s, v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:168
	// _ = "end of CoverTab[108869]"
}

func (p pointer) toUint64() *uint64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:171
	_go_fuzz_dep_.CoverTab[108870]++
												return (*uint64)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:172
	// _ = "end of CoverTab[108870]"
}
func (p pointer) toUint64Ptr() **uint64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:174
	_go_fuzz_dep_.CoverTab[108871]++
												return (**uint64)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:175
	// _ = "end of CoverTab[108871]"
}
func (p pointer) toUint64Slice() *[]uint64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:177
	_go_fuzz_dep_.CoverTab[108872]++
												return (*[]uint64)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:178
	// _ = "end of CoverTab[108872]"
}
func (p pointer) toUint32() *uint32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:180
	_go_fuzz_dep_.CoverTab[108873]++
												return (*uint32)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:181
	// _ = "end of CoverTab[108873]"
}
func (p pointer) toUint32Ptr() **uint32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:183
	_go_fuzz_dep_.CoverTab[108874]++
												return (**uint32)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:184
	// _ = "end of CoverTab[108874]"
}
func (p pointer) toUint32Slice() *[]uint32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:186
	_go_fuzz_dep_.CoverTab[108875]++
												return (*[]uint32)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:187
	// _ = "end of CoverTab[108875]"
}
func (p pointer) toBool() *bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:189
	_go_fuzz_dep_.CoverTab[108876]++
												return (*bool)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:190
	// _ = "end of CoverTab[108876]"
}
func (p pointer) toBoolPtr() **bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:192
	_go_fuzz_dep_.CoverTab[108877]++
												return (**bool)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:193
	// _ = "end of CoverTab[108877]"
}
func (p pointer) toBoolSlice() *[]bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:195
	_go_fuzz_dep_.CoverTab[108878]++
												return (*[]bool)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:196
	// _ = "end of CoverTab[108878]"
}
func (p pointer) toFloat64() *float64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:198
	_go_fuzz_dep_.CoverTab[108879]++
												return (*float64)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:199
	// _ = "end of CoverTab[108879]"
}
func (p pointer) toFloat64Ptr() **float64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:201
	_go_fuzz_dep_.CoverTab[108880]++
												return (**float64)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:202
	// _ = "end of CoverTab[108880]"
}
func (p pointer) toFloat64Slice() *[]float64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:204
	_go_fuzz_dep_.CoverTab[108881]++
												return (*[]float64)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:205
	// _ = "end of CoverTab[108881]"
}
func (p pointer) toFloat32() *float32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:207
	_go_fuzz_dep_.CoverTab[108882]++
												return (*float32)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:208
	// _ = "end of CoverTab[108882]"
}
func (p pointer) toFloat32Ptr() **float32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:210
	_go_fuzz_dep_.CoverTab[108883]++
												return (**float32)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:211
	// _ = "end of CoverTab[108883]"
}
func (p pointer) toFloat32Slice() *[]float32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:213
	_go_fuzz_dep_.CoverTab[108884]++
												return (*[]float32)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:214
	// _ = "end of CoverTab[108884]"
}
func (p pointer) toString() *string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:216
	_go_fuzz_dep_.CoverTab[108885]++
												return (*string)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:217
	// _ = "end of CoverTab[108885]"
}
func (p pointer) toStringPtr() **string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:219
	_go_fuzz_dep_.CoverTab[108886]++
												return (**string)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:220
	// _ = "end of CoverTab[108886]"
}
func (p pointer) toStringSlice() *[]string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:222
	_go_fuzz_dep_.CoverTab[108887]++
												return (*[]string)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:223
	// _ = "end of CoverTab[108887]"
}
func (p pointer) toBytes() *[]byte {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:225
	_go_fuzz_dep_.CoverTab[108888]++
												return (*[]byte)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:226
	// _ = "end of CoverTab[108888]"
}
func (p pointer) toBytesSlice() *[][]byte {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:228
	_go_fuzz_dep_.CoverTab[108889]++
												return (*[][]byte)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:229
	// _ = "end of CoverTab[108889]"
}
func (p pointer) toExtensions() *XXX_InternalExtensions {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:231
	_go_fuzz_dep_.CoverTab[108890]++
												return (*XXX_InternalExtensions)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:232
	// _ = "end of CoverTab[108890]"
}
func (p pointer) toOldExtensions() *map[int32]Extension {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:234
	_go_fuzz_dep_.CoverTab[108891]++
												return (*map[int32]Extension)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:235
	// _ = "end of CoverTab[108891]"
}

// getPointerSlice loads []*T from p as a []pointer.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:238
// The value returned is aliased with the original slice.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:238
// This behavior differs from the implementation in pointer_reflect.go.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:241
func (p pointer) getPointerSlice() []pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:241
	_go_fuzz_dep_.CoverTab[108892]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:244
	return *(*[]pointer)(p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:244
	// _ = "end of CoverTab[108892]"
}

// setPointerSlice stores []pointer into p as a []*T.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:247
// The value set is aliased with the input slice.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:247
// This behavior differs from the implementation in pointer_reflect.go.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:250
func (p pointer) setPointerSlice(v []pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:250
	_go_fuzz_dep_.CoverTab[108893]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:253
	*(*[]pointer)(p.p) = v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:253
	// _ = "end of CoverTab[108893]"
}

// getPointer loads the pointer at p and returns it.
func (p pointer) getPointer() pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:257
	_go_fuzz_dep_.CoverTab[108894]++
												return pointer{p: *(*unsafe.Pointer)(p.p)}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:258
	// _ = "end of CoverTab[108894]"
}

// setPointer stores the pointer q at p.
func (p pointer) setPointer(q pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:262
	_go_fuzz_dep_.CoverTab[108895]++
												*(*unsafe.Pointer)(p.p) = q.p
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:263
	// _ = "end of CoverTab[108895]"
}

// append q to the slice pointed to by p.
func (p pointer) appendPointer(q pointer) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:267
	_go_fuzz_dep_.CoverTab[108896]++
												s := (*[]unsafe.Pointer)(p.p)
												*s = append(*s, q.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:269
	// _ = "end of CoverTab[108896]"
}

// getInterfacePointer returns a pointer that points to the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:272
// interface data of the interface pointed by p.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:274
func (p pointer) getInterfacePointer() pointer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:274
	_go_fuzz_dep_.CoverTab[108897]++

												return pointer{p: (*(*[2]unsafe.Pointer)(p.p))[1]}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:276
	// _ = "end of CoverTab[108897]"
}

// asPointerTo returns a reflect.Value that is a pointer to an
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:279
// object of type t stored at p.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:281
func (p pointer) asPointerTo(t reflect.Type) reflect.Value {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:281
	_go_fuzz_dep_.CoverTab[108898]++
												return reflect.NewAt(t, p.p)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:282
	// _ = "end of CoverTab[108898]"
}

func atomicLoadUnmarshalInfo(p **unmarshalInfo) *unmarshalInfo {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:285
	_go_fuzz_dep_.CoverTab[108899]++
												return (*unmarshalInfo)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(p))))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:286
	// _ = "end of CoverTab[108899]"
}
func atomicStoreUnmarshalInfo(p **unmarshalInfo, v *unmarshalInfo) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:288
	_go_fuzz_dep_.CoverTab[108900]++
												atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(p)), unsafe.Pointer(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:289
	// _ = "end of CoverTab[108900]"
}
func atomicLoadMarshalInfo(p **marshalInfo) *marshalInfo {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:291
	_go_fuzz_dep_.CoverTab[108901]++
												return (*marshalInfo)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(p))))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:292
	// _ = "end of CoverTab[108901]"
}
func atomicStoreMarshalInfo(p **marshalInfo, v *marshalInfo) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:294
	_go_fuzz_dep_.CoverTab[108902]++
												atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(p)), unsafe.Pointer(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:295
	// _ = "end of CoverTab[108902]"
}
func atomicLoadMergeInfo(p **mergeInfo) *mergeInfo {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:297
	_go_fuzz_dep_.CoverTab[108903]++
												return (*mergeInfo)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(p))))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:298
	// _ = "end of CoverTab[108903]"
}
func atomicStoreMergeInfo(p **mergeInfo, v *mergeInfo) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:300
	_go_fuzz_dep_.CoverTab[108904]++
												atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(p)), unsafe.Pointer(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:301
	// _ = "end of CoverTab[108904]"
}
func atomicLoadDiscardInfo(p **discardInfo) *discardInfo {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:303
	_go_fuzz_dep_.CoverTab[108905]++
												return (*discardInfo)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(p))))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:304
	// _ = "end of CoverTab[108905]"
}
func atomicStoreDiscardInfo(p **discardInfo, v *discardInfo) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:306
	_go_fuzz_dep_.CoverTab[108906]++
												atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(p)), unsafe.Pointer(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:307
	// _ = "end of CoverTab[108906]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:308
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/pointer_unsafe.go:308
var _ = _go_fuzz_dep_.CoverTab
