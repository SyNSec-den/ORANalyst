// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:5
)

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type mapConverter struct {
	goType			reflect.Type	// map[K]V
	keyConv, valConv	Converter
}

func newMapConverter(t reflect.Type, fd protoreflect.FieldDescriptor) *mapConverter {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:19
	_go_fuzz_dep_.CoverTab[57135]++
													if t.Kind() != reflect.Map {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:20
		_go_fuzz_dep_.CoverTab[57137]++
														panic(fmt.Sprintf("invalid Go type %v for field %v", t, fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:21
		// _ = "end of CoverTab[57137]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:22
		_go_fuzz_dep_.CoverTab[57138]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:22
		// _ = "end of CoverTab[57138]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:22
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:22
	// _ = "end of CoverTab[57135]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:22
	_go_fuzz_dep_.CoverTab[57136]++
													return &mapConverter{
		goType:		t,
		keyConv:	newSingularConverter(t.Key(), fd.MapKey()),
		valConv:	newSingularConverter(t.Elem(), fd.MapValue()),
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:27
	// _ = "end of CoverTab[57136]"
}

func (c *mapConverter) PBValueOf(v reflect.Value) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:30
	_go_fuzz_dep_.CoverTab[57139]++
													if v.Type() != c.goType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:31
		_go_fuzz_dep_.CoverTab[57141]++
														panic(fmt.Sprintf("invalid type: got %v, want %v", v.Type(), c.goType))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:32
		// _ = "end of CoverTab[57141]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:33
		_go_fuzz_dep_.CoverTab[57142]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:33
		// _ = "end of CoverTab[57142]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:33
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:33
	// _ = "end of CoverTab[57139]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:33
	_go_fuzz_dep_.CoverTab[57140]++
													return protoreflect.ValueOfMap(&mapReflect{v, c.keyConv, c.valConv})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:34
	// _ = "end of CoverTab[57140]"
}

func (c *mapConverter) GoValueOf(v protoreflect.Value) reflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:37
	_go_fuzz_dep_.CoverTab[57143]++
													return v.Map().(*mapReflect).v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:38
	// _ = "end of CoverTab[57143]"
}

func (c *mapConverter) IsValidPB(v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:41
	_go_fuzz_dep_.CoverTab[57144]++
													mapv, ok := v.Interface().(*mapReflect)
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:43
		_go_fuzz_dep_.CoverTab[57146]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:44
		// _ = "end of CoverTab[57146]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:45
		_go_fuzz_dep_.CoverTab[57147]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:45
		// _ = "end of CoverTab[57147]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:45
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:45
	// _ = "end of CoverTab[57144]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:45
	_go_fuzz_dep_.CoverTab[57145]++
													return mapv.v.Type() == c.goType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:46
	// _ = "end of CoverTab[57145]"
}

func (c *mapConverter) IsValidGo(v reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:49
	_go_fuzz_dep_.CoverTab[57148]++
													return v.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:50
		_go_fuzz_dep_.CoverTab[57149]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:50
		return v.Type() == c.goType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:50
		// _ = "end of CoverTab[57149]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:50
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:50
	// _ = "end of CoverTab[57148]"
}

func (c *mapConverter) New() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:53
	_go_fuzz_dep_.CoverTab[57150]++
													return c.PBValueOf(reflect.MakeMap(c.goType))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:54
	// _ = "end of CoverTab[57150]"
}

func (c *mapConverter) Zero() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:57
	_go_fuzz_dep_.CoverTab[57151]++
													return c.PBValueOf(reflect.Zero(c.goType))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:58
	// _ = "end of CoverTab[57151]"
}

type mapReflect struct {
	v	reflect.Value	// map[K]V
	keyConv	Converter
	valConv	Converter
}

func (ms *mapReflect) Len() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:67
	_go_fuzz_dep_.CoverTab[57152]++
													return ms.v.Len()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:68
	// _ = "end of CoverTab[57152]"
}
func (ms *mapReflect) Has(k protoreflect.MapKey) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:70
	_go_fuzz_dep_.CoverTab[57153]++
													rk := ms.keyConv.GoValueOf(k.Value())
													rv := ms.v.MapIndex(rk)
													return rv.IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:73
	// _ = "end of CoverTab[57153]"
}
func (ms *mapReflect) Get(k protoreflect.MapKey) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:75
	_go_fuzz_dep_.CoverTab[57154]++
													rk := ms.keyConv.GoValueOf(k.Value())
													rv := ms.v.MapIndex(rk)
													if !rv.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:78
		_go_fuzz_dep_.CoverTab[57156]++
														return protoreflect.Value{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:79
		// _ = "end of CoverTab[57156]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:80
		_go_fuzz_dep_.CoverTab[57157]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:80
		// _ = "end of CoverTab[57157]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:80
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:80
	// _ = "end of CoverTab[57154]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:80
	_go_fuzz_dep_.CoverTab[57155]++
													return ms.valConv.PBValueOf(rv)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:81
	// _ = "end of CoverTab[57155]"
}
func (ms *mapReflect) Set(k protoreflect.MapKey, v protoreflect.Value) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:83
	_go_fuzz_dep_.CoverTab[57158]++
													rk := ms.keyConv.GoValueOf(k.Value())
													rv := ms.valConv.GoValueOf(v)
													ms.v.SetMapIndex(rk, rv)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:86
	// _ = "end of CoverTab[57158]"
}
func (ms *mapReflect) Clear(k protoreflect.MapKey) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:88
	_go_fuzz_dep_.CoverTab[57159]++
													rk := ms.keyConv.GoValueOf(k.Value())
													ms.v.SetMapIndex(rk, reflect.Value{})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:90
	// _ = "end of CoverTab[57159]"
}
func (ms *mapReflect) Mutable(k protoreflect.MapKey) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:92
	_go_fuzz_dep_.CoverTab[57160]++
													if _, ok := ms.valConv.(*messageConverter); !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:93
		_go_fuzz_dep_.CoverTab[57163]++
														panic("invalid Mutable on map with non-message value type")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:94
		// _ = "end of CoverTab[57163]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:95
		_go_fuzz_dep_.CoverTab[57164]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:95
		// _ = "end of CoverTab[57164]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:95
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:95
	// _ = "end of CoverTab[57160]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:95
	_go_fuzz_dep_.CoverTab[57161]++
													v := ms.Get(k)
													if !v.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:97
		_go_fuzz_dep_.CoverTab[57165]++
														v = ms.NewValue()
														ms.Set(k, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:99
		// _ = "end of CoverTab[57165]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:100
		_go_fuzz_dep_.CoverTab[57166]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:100
		// _ = "end of CoverTab[57166]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:100
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:100
	// _ = "end of CoverTab[57161]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:100
	_go_fuzz_dep_.CoverTab[57162]++
													return v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:101
	// _ = "end of CoverTab[57162]"
}
func (ms *mapReflect) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:103
	_go_fuzz_dep_.CoverTab[57167]++
													iter := mapRange(ms.v)
													for iter.Next() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:105
		_go_fuzz_dep_.CoverTab[57168]++
														k := ms.keyConv.PBValueOf(iter.Key()).MapKey()
														v := ms.valConv.PBValueOf(iter.Value())
														if !f(k, v) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:108
			_go_fuzz_dep_.CoverTab[57169]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:109
			// _ = "end of CoverTab[57169]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:110
			_go_fuzz_dep_.CoverTab[57170]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:110
			// _ = "end of CoverTab[57170]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:110
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:110
		// _ = "end of CoverTab[57168]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:111
	// _ = "end of CoverTab[57167]"
}
func (ms *mapReflect) NewValue() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:113
	_go_fuzz_dep_.CoverTab[57171]++
													return ms.valConv.New()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:114
	// _ = "end of CoverTab[57171]"
}
func (ms *mapReflect) IsValid() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:116
	_go_fuzz_dep_.CoverTab[57172]++
													return !ms.v.IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:117
	// _ = "end of CoverTab[57172]"
}
func (ms *mapReflect) protoUnwrap() interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:119
	_go_fuzz_dep_.CoverTab[57173]++
													return ms.v.Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:120
	// _ = "end of CoverTab[57173]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:121
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_map.go:121
var _ = _go_fuzz_dep_.CoverTab
