// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego && !appengine
// +build !purego,!appengine

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:8
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:8
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:8
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:8
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:8
)

import (
	"reflect"
	"sync/atomic"
	"unsafe"
)

const UnsafeEnabled = true

// Pointer is an opaque pointer type.
type Pointer unsafe.Pointer

// offset represents the offset to a struct field, accessible from a pointer.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:21
// The offset is the byte offset to the field from the start of the struct.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:23
type offset uintptr

// offsetOf returns a field offset for the struct field.
func offsetOf(f reflect.StructField, x exporter) offset {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:26
	_go_fuzz_dep_.CoverTab[58663]++
														return offset(f.Offset)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:27
	// _ = "end of CoverTab[58663]"
}

// IsValid reports whether the offset is valid.
func (f offset) IsValid() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:31
	_go_fuzz_dep_.CoverTab[58664]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:31
	return f != invalidOffset
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:31
	// _ = "end of CoverTab[58664]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:31
}

// invalidOffset is an invalid field offset.
var invalidOffset = ^offset(0)

// zeroOffset is a noop when calling pointer.Apply.
var zeroOffset = offset(0)

// pointer is a pointer to a message struct or field.
type pointer struct{ p unsafe.Pointer }

// pointerOf returns p as a pointer.
func pointerOf(p Pointer) pointer {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:43
	_go_fuzz_dep_.CoverTab[58665]++
														return pointer{p: unsafe.Pointer(p)}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:44
	// _ = "end of CoverTab[58665]"
}

// pointerOfValue returns v as a pointer.
func pointerOfValue(v reflect.Value) pointer {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:48
	_go_fuzz_dep_.CoverTab[58666]++
														return pointer{p: unsafe.Pointer(v.Pointer())}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:49
	// _ = "end of CoverTab[58666]"
}

// pointerOfIface returns the pointer portion of an interface.
func pointerOfIface(v interface{}) pointer {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:53
	_go_fuzz_dep_.CoverTab[58667]++
														type ifaceHeader struct {
		Type	unsafe.Pointer
		Data	unsafe.Pointer
	}
														return pointer{p: (*ifaceHeader)(unsafe.Pointer(&v)).Data}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:58
	// _ = "end of CoverTab[58667]"
}

// IsNil reports whether the pointer is nil.
func (p pointer) IsNil() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:62
	_go_fuzz_dep_.CoverTab[58668]++
														return p.p == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:63
	// _ = "end of CoverTab[58668]"
}

// Apply adds an offset to the pointer to derive a new pointer
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:66
// to a specified field. The pointer must be valid and pointing at a struct.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:68
func (p pointer) Apply(f offset) pointer {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:68
	_go_fuzz_dep_.CoverTab[58669]++
														if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:69
		_go_fuzz_dep_.CoverTab[58671]++
															panic("invalid nil pointer")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:70
		// _ = "end of CoverTab[58671]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:71
		_go_fuzz_dep_.CoverTab[58672]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:71
		// _ = "end of CoverTab[58672]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:71
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:71
	// _ = "end of CoverTab[58669]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:71
	_go_fuzz_dep_.CoverTab[58670]++
														return pointer{p: unsafe.Pointer(uintptr(p.p) + uintptr(f))}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:72
	// _ = "end of CoverTab[58670]"
}

// AsValueOf treats p as a pointer to an object of type t and returns the value.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:75
// It is equivalent to reflect.ValueOf(p.AsIfaceOf(t))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:77
func (p pointer) AsValueOf(t reflect.Type) reflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:77
	_go_fuzz_dep_.CoverTab[58673]++
														return reflect.NewAt(t, p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:78
	// _ = "end of CoverTab[58673]"
}

// AsIfaceOf treats p as a pointer to an object of type t and returns the value.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:81
// It is equivalent to p.AsValueOf(t).Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:83
func (p pointer) AsIfaceOf(t reflect.Type) interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:83
	_go_fuzz_dep_.CoverTab[58674]++

														return p.AsValueOf(t).Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:85
	// _ = "end of CoverTab[58674]"
}

func (p pointer) Bool() *bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:88
	_go_fuzz_dep_.CoverTab[58675]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:88
	return (*bool)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:88
	// _ = "end of CoverTab[58675]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:88
}
func (p pointer) BoolPtr() **bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:89
	_go_fuzz_dep_.CoverTab[58676]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:89
	return (**bool)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:89
	// _ = "end of CoverTab[58676]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:89
}
func (p pointer) BoolSlice() *[]bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:90
	_go_fuzz_dep_.CoverTab[58677]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:90
	return (*[]bool)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:90
	// _ = "end of CoverTab[58677]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:90
}
func (p pointer) Int32() *int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:91
	_go_fuzz_dep_.CoverTab[58678]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:91
	return (*int32)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:91
	// _ = "end of CoverTab[58678]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:91
}
func (p pointer) Int32Ptr() **int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:92
	_go_fuzz_dep_.CoverTab[58679]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:92
	return (**int32)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:92
	// _ = "end of CoverTab[58679]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:92
}
func (p pointer) Int32Slice() *[]int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:93
	_go_fuzz_dep_.CoverTab[58680]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:93
	return (*[]int32)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:93
	// _ = "end of CoverTab[58680]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:93
}
func (p pointer) Int64() *int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:94
	_go_fuzz_dep_.CoverTab[58681]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:94
	return (*int64)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:94
	// _ = "end of CoverTab[58681]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:94
}
func (p pointer) Int64Ptr() **int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:95
	_go_fuzz_dep_.CoverTab[58682]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:95
	return (**int64)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:95
	// _ = "end of CoverTab[58682]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:95
}
func (p pointer) Int64Slice() *[]int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:96
	_go_fuzz_dep_.CoverTab[58683]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:96
	return (*[]int64)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:96
	// _ = "end of CoverTab[58683]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:96
}
func (p pointer) Uint32() *uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:97
	_go_fuzz_dep_.CoverTab[58684]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:97
	return (*uint32)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:97
	// _ = "end of CoverTab[58684]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:97
}
func (p pointer) Uint32Ptr() **uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:98
	_go_fuzz_dep_.CoverTab[58685]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:98
	return (**uint32)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:98
	// _ = "end of CoverTab[58685]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:98
}
func (p pointer) Uint32Slice() *[]uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:99
	_go_fuzz_dep_.CoverTab[58686]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:99
	return (*[]uint32)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:99
	// _ = "end of CoverTab[58686]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:99
}
func (p pointer) Uint64() *uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:100
	_go_fuzz_dep_.CoverTab[58687]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:100
	return (*uint64)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:100
	// _ = "end of CoverTab[58687]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:100
}
func (p pointer) Uint64Ptr() **uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:101
	_go_fuzz_dep_.CoverTab[58688]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:101
	return (**uint64)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:101
	// _ = "end of CoverTab[58688]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:101
}
func (p pointer) Uint64Slice() *[]uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:102
	_go_fuzz_dep_.CoverTab[58689]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:102
	return (*[]uint64)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:102
	// _ = "end of CoverTab[58689]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:102
}
func (p pointer) Float32() *float32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:103
	_go_fuzz_dep_.CoverTab[58690]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:103
	return (*float32)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:103
	// _ = "end of CoverTab[58690]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:103
}
func (p pointer) Float32Ptr() **float32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:104
	_go_fuzz_dep_.CoverTab[58691]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:104
	return (**float32)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:104
	// _ = "end of CoverTab[58691]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:104
}
func (p pointer) Float32Slice() *[]float32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:105
	_go_fuzz_dep_.CoverTab[58692]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:105
	return (*[]float32)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:105
	// _ = "end of CoverTab[58692]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:105
}
func (p pointer) Float64() *float64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:106
	_go_fuzz_dep_.CoverTab[58693]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:106
	return (*float64)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:106
	// _ = "end of CoverTab[58693]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:106
}
func (p pointer) Float64Ptr() **float64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:107
	_go_fuzz_dep_.CoverTab[58694]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:107
	return (**float64)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:107
	// _ = "end of CoverTab[58694]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:107
}
func (p pointer) Float64Slice() *[]float64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:108
	_go_fuzz_dep_.CoverTab[58695]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:108
	return (*[]float64)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:108
	// _ = "end of CoverTab[58695]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:108
}
func (p pointer) String() *string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:109
	_go_fuzz_dep_.CoverTab[58696]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:109
	return (*string)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:109
	// _ = "end of CoverTab[58696]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:109
}
func (p pointer) StringPtr() **string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:110
	_go_fuzz_dep_.CoverTab[58697]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:110
	return (**string)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:110
	// _ = "end of CoverTab[58697]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:110
}
func (p pointer) StringSlice() *[]string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:111
	_go_fuzz_dep_.CoverTab[58698]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:111
	return (*[]string)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:111
	// _ = "end of CoverTab[58698]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:111
}
func (p pointer) Bytes() *[]byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:112
	_go_fuzz_dep_.CoverTab[58699]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:112
	return (*[]byte)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:112
	// _ = "end of CoverTab[58699]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:112
}
func (p pointer) BytesPtr() **[]byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:113
	_go_fuzz_dep_.CoverTab[58700]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:113
	return (**[]byte)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:113
	// _ = "end of CoverTab[58700]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:113
}
func (p pointer) BytesSlice() *[][]byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:114
	_go_fuzz_dep_.CoverTab[58701]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:114
	return (*[][]byte)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:114
	// _ = "end of CoverTab[58701]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:114
}
func (p pointer) WeakFields() *weakFields {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:115
	_go_fuzz_dep_.CoverTab[58702]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:115
	return (*weakFields)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:115
	// _ = "end of CoverTab[58702]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:115
}
func (p pointer) Extensions() *map[int32]ExtensionField {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:116
	_go_fuzz_dep_.CoverTab[58703]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:116
	return (*map[int32]ExtensionField)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:116
	// _ = "end of CoverTab[58703]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:116
}

func (p pointer) Elem() pointer {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:118
	_go_fuzz_dep_.CoverTab[58704]++
														return pointer{p: *(*unsafe.Pointer)(p.p)}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:119
	// _ = "end of CoverTab[58704]"
}

// PointerSlice loads []*T from p as a []pointer.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:122
// The value returned is aliased with the original slice.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:122
// This behavior differs from the implementation in pointer_reflect.go.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:125
func (p pointer) PointerSlice() []pointer {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:125
	_go_fuzz_dep_.CoverTab[58705]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:128
	return *(*[]pointer)(p.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:128
	// _ = "end of CoverTab[58705]"
}

// AppendPointerSlice appends v to p, which must be a []*T.
func (p pointer) AppendPointerSlice(v pointer) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:132
	_go_fuzz_dep_.CoverTab[58706]++
														*(*[]pointer)(p.p) = append(*(*[]pointer)(p.p), v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:133
	// _ = "end of CoverTab[58706]"
}

// SetPointer sets *p to v.
func (p pointer) SetPointer(v pointer) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:137
	_go_fuzz_dep_.CoverTab[58707]++
														*(*unsafe.Pointer)(p.p) = (unsafe.Pointer)(v.p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:138
	// _ = "end of CoverTab[58707]"
}

// Static check that MessageState does not exceed the size of a pointer.
const _ = uint(unsafe.Sizeof(unsafe.Pointer(nil)) - unsafe.Sizeof(MessageState{}))

func (Export) MessageStateOf(p Pointer) *messageState {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:144
	_go_fuzz_dep_.CoverTab[58708]++

														return (*messageState)(unsafe.Pointer(p))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:146
	// _ = "end of CoverTab[58708]"
}
func (ms *messageState) pointer() pointer {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:148
	_go_fuzz_dep_.CoverTab[58709]++

														return pointer{p: unsafe.Pointer(ms)}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:150
	// _ = "end of CoverTab[58709]"
}
func (ms *messageState) messageInfo() *MessageInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:152
	_go_fuzz_dep_.CoverTab[58710]++
														mi := ms.LoadMessageInfo()
														if mi == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:154
		_go_fuzz_dep_.CoverTab[58712]++
															panic("invalid nil message info; this suggests memory corruption due to a race or shallow copy on the message struct")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:155
		// _ = "end of CoverTab[58712]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:156
		_go_fuzz_dep_.CoverTab[58713]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:156
		// _ = "end of CoverTab[58713]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:156
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:156
	// _ = "end of CoverTab[58710]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:156
	_go_fuzz_dep_.CoverTab[58711]++
														return mi
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:157
	// _ = "end of CoverTab[58711]"
}
func (ms *messageState) LoadMessageInfo() *MessageInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:159
	_go_fuzz_dep_.CoverTab[58714]++
														return (*MessageInfo)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&ms.atomicMessageInfo))))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:160
	// _ = "end of CoverTab[58714]"
}
func (ms *messageState) StoreMessageInfo(mi *MessageInfo) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:162
	_go_fuzz_dep_.CoverTab[58715]++
														atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&ms.atomicMessageInfo)), unsafe.Pointer(mi))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:163
	// _ = "end of CoverTab[58715]"
}

type atomicNilMessage struct{ p unsafe.Pointer }	// p is a *messageReflectWrapper

func (m *atomicNilMessage) Init(mi *MessageInfo) *messageReflectWrapper {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:168
	_go_fuzz_dep_.CoverTab[58716]++
														if p := atomic.LoadPointer(&m.p); p != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:169
		_go_fuzz_dep_.CoverTab[58718]++
															return (*messageReflectWrapper)(p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:170
		// _ = "end of CoverTab[58718]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:171
		_go_fuzz_dep_.CoverTab[58719]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:171
		// _ = "end of CoverTab[58719]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:171
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:171
	// _ = "end of CoverTab[58716]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:171
	_go_fuzz_dep_.CoverTab[58717]++
														w := &messageReflectWrapper{mi: mi}
														atomic.CompareAndSwapPointer(&m.p, nil, (unsafe.Pointer)(w))
														return (*messageReflectWrapper)(atomic.LoadPointer(&m.p))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:174
	// _ = "end of CoverTab[58717]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:175
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/pointer_unsafe.go:175
var _ = _go_fuzz_dep_.CoverTab
