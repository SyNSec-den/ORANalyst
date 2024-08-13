// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego && !appengine
// +build !purego,!appengine

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:8
package protoreflect

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:8
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:8
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:8
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:8
)

import (
	"unsafe"

	"google.golang.org/protobuf/internal/pragma"
)

type (
	stringHeader	struct {
		Data	unsafe.Pointer
		Len	int
	}
	sliceHeader	struct {
		Data	unsafe.Pointer
		Len	int
		Cap	int
	}
	ifaceHeader	struct {
		Type	unsafe.Pointer
		Data	unsafe.Pointer
	}
)

var (
	nilType		= typeOf(nil)
	boolType	= typeOf(*new(bool))
	int32Type	= typeOf(*new(int32))
	int64Type	= typeOf(*new(int64))
	uint32Type	= typeOf(*new(uint32))
	uint64Type	= typeOf(*new(uint64))
	float32Type	= typeOf(*new(float32))
	float64Type	= typeOf(*new(float64))
	stringType	= typeOf(*new(string))
	bytesType	= typeOf(*new([]byte))
	enumType	= typeOf(*new(EnumNumber))
)

// typeOf returns a pointer to the Go type information.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:46
// The pointer is comparable and equal if and only if the types are identical.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:48
func typeOf(t interface{}) unsafe.Pointer {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:48
	_go_fuzz_dep_.CoverTab[49005]++
														return (*ifaceHeader)(unsafe.Pointer(&t)).Type
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:49
	// _ = "end of CoverTab[49005]"
}

// value is a union where only one type can be represented at a time.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:52
// The struct is 24B large on 64-bit systems and requires the minimum storage
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:52
// necessary to represent each possible type.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:52
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:52
// The Go GC needs to be able to scan variables containing pointers.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:52
// As such, pointers and non-pointers cannot be intermixed.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:58
type value struct {
	pragma.DoNotCompare	// 0B

	// typ stores the type of the value as a pointer to the Go type.
	typ	unsafe.Pointer	// 8B

	// ptr stores the data pointer for a String, Bytes, or interface value.
	ptr	unsafe.Pointer	// 8B

	// num stores a Bool, Int32, Int64, Uint32, Uint64, Float32, Float64, or
	// Enum value as a raw uint64.
	//
	// It is also used to store the length of a String or Bytes value;
	// the capacity is ignored.
	num	uint64	// 8B
}

func valueOfString(v string) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:75
	_go_fuzz_dep_.CoverTab[49006]++
														p := (*stringHeader)(unsafe.Pointer(&v))
														return Value{typ: stringType, ptr: p.Data, num: uint64(len(v))}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:77
	// _ = "end of CoverTab[49006]"
}
func valueOfBytes(v []byte) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:79
	_go_fuzz_dep_.CoverTab[49007]++
														p := (*sliceHeader)(unsafe.Pointer(&v))
														return Value{typ: bytesType, ptr: p.Data, num: uint64(len(v))}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:81
	// _ = "end of CoverTab[49007]"
}
func valueOfIface(v interface{}) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:83
	_go_fuzz_dep_.CoverTab[49008]++
														p := (*ifaceHeader)(unsafe.Pointer(&v))
														return Value{typ: p.Type, ptr: p.Data}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:85
	// _ = "end of CoverTab[49008]"
}

func (v Value) getString() (x string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:88
	_go_fuzz_dep_.CoverTab[49009]++
														*(*stringHeader)(unsafe.Pointer(&x)) = stringHeader{Data: v.ptr, Len: int(v.num)}
														return x
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:90
	// _ = "end of CoverTab[49009]"
}
func (v Value) getBytes() (x []byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:92
	_go_fuzz_dep_.CoverTab[49010]++
														*(*sliceHeader)(unsafe.Pointer(&x)) = sliceHeader{Data: v.ptr, Len: int(v.num), Cap: int(v.num)}
														return x
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:94
	// _ = "end of CoverTab[49010]"
}
func (v Value) getIface() (x interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:96
	_go_fuzz_dep_.CoverTab[49011]++
														*(*ifaceHeader)(unsafe.Pointer(&x)) = ifaceHeader{Type: v.typ, Data: v.ptr}
														return x
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:98
	// _ = "end of CoverTab[49011]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:99
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_unsafe.go:99
var _ = _go_fuzz_dep_.CoverTab
