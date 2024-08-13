// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:5
)

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func newListConverter(t reflect.Type, fd protoreflect.FieldDescriptor) Converter {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:14
	_go_fuzz_dep_.CoverTab[57085]++
													switch {
	case t.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:16
		_go_fuzz_dep_.CoverTab[57090]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:16
		return t.Elem().Kind() == reflect.Slice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:16
		// _ = "end of CoverTab[57090]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:16
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:16
		_go_fuzz_dep_.CoverTab[57087]++
														return &listPtrConverter{t, newSingularConverter(t.Elem().Elem(), fd)}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:17
		// _ = "end of CoverTab[57087]"
	case t.Kind() == reflect.Slice:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:18
		_go_fuzz_dep_.CoverTab[57088]++
														return &listConverter{t, newSingularConverter(t.Elem(), fd)}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:19
		// _ = "end of CoverTab[57088]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:19
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:19
		_go_fuzz_dep_.CoverTab[57089]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:19
		// _ = "end of CoverTab[57089]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:20
	// _ = "end of CoverTab[57085]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:20
	_go_fuzz_dep_.CoverTab[57086]++
													panic(fmt.Sprintf("invalid Go type %v for field %v", t, fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:21
	// _ = "end of CoverTab[57086]"
}

type listConverter struct {
	goType	reflect.Type	// []T
	c	Converter
}

func (c *listConverter) PBValueOf(v reflect.Value) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:29
	_go_fuzz_dep_.CoverTab[57091]++
													if v.Type() != c.goType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:30
		_go_fuzz_dep_.CoverTab[57093]++
														panic(fmt.Sprintf("invalid type: got %v, want %v", v.Type(), c.goType))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:31
		// _ = "end of CoverTab[57093]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:32
		_go_fuzz_dep_.CoverTab[57094]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:32
		// _ = "end of CoverTab[57094]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:32
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:32
	// _ = "end of CoverTab[57091]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:32
	_go_fuzz_dep_.CoverTab[57092]++
													pv := reflect.New(c.goType)
													pv.Elem().Set(v)
													return protoreflect.ValueOfList(&listReflect{pv, c.c})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:35
	// _ = "end of CoverTab[57092]"
}

func (c *listConverter) GoValueOf(v protoreflect.Value) reflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:38
	_go_fuzz_dep_.CoverTab[57095]++
													rv := v.List().(*listReflect).v
													if rv.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:40
		_go_fuzz_dep_.CoverTab[57097]++
														return reflect.Zero(c.goType)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:41
		// _ = "end of CoverTab[57097]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:42
		_go_fuzz_dep_.CoverTab[57098]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:42
		// _ = "end of CoverTab[57098]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:42
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:42
	// _ = "end of CoverTab[57095]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:42
	_go_fuzz_dep_.CoverTab[57096]++
													return rv.Elem()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:43
	// _ = "end of CoverTab[57096]"
}

func (c *listConverter) IsValidPB(v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:46
	_go_fuzz_dep_.CoverTab[57099]++
													list, ok := v.Interface().(*listReflect)
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:48
		_go_fuzz_dep_.CoverTab[57101]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:49
		// _ = "end of CoverTab[57101]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:50
		_go_fuzz_dep_.CoverTab[57102]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:50
		// _ = "end of CoverTab[57102]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:50
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:50
	// _ = "end of CoverTab[57099]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:50
	_go_fuzz_dep_.CoverTab[57100]++
													return list.v.Type().Elem() == c.goType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:51
	// _ = "end of CoverTab[57100]"
}

func (c *listConverter) IsValidGo(v reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:54
	_go_fuzz_dep_.CoverTab[57103]++
													return v.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:55
		_go_fuzz_dep_.CoverTab[57104]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:55
		return v.Type() == c.goType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:55
		// _ = "end of CoverTab[57104]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:55
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:55
	// _ = "end of CoverTab[57103]"
}

func (c *listConverter) New() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:58
	_go_fuzz_dep_.CoverTab[57105]++
													return protoreflect.ValueOfList(&listReflect{reflect.New(c.goType), c.c})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:59
	// _ = "end of CoverTab[57105]"
}

func (c *listConverter) Zero() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:62
	_go_fuzz_dep_.CoverTab[57106]++
													return protoreflect.ValueOfList(&listReflect{reflect.Zero(reflect.PtrTo(c.goType)), c.c})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:63
	// _ = "end of CoverTab[57106]"
}

type listPtrConverter struct {
	goType	reflect.Type	// *[]T
	c	Converter
}

func (c *listPtrConverter) PBValueOf(v reflect.Value) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:71
	_go_fuzz_dep_.CoverTab[57107]++
													if v.Type() != c.goType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:72
		_go_fuzz_dep_.CoverTab[57109]++
														panic(fmt.Sprintf("invalid type: got %v, want %v", v.Type(), c.goType))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:73
		// _ = "end of CoverTab[57109]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:74
		_go_fuzz_dep_.CoverTab[57110]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:74
		// _ = "end of CoverTab[57110]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:74
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:74
	// _ = "end of CoverTab[57107]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:74
	_go_fuzz_dep_.CoverTab[57108]++
													return protoreflect.ValueOfList(&listReflect{v, c.c})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:75
	// _ = "end of CoverTab[57108]"
}

func (c *listPtrConverter) GoValueOf(v protoreflect.Value) reflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:78
	_go_fuzz_dep_.CoverTab[57111]++
													return v.List().(*listReflect).v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:79
	// _ = "end of CoverTab[57111]"
}

func (c *listPtrConverter) IsValidPB(v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:82
	_go_fuzz_dep_.CoverTab[57112]++
													list, ok := v.Interface().(*listReflect)
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:84
		_go_fuzz_dep_.CoverTab[57114]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:85
		// _ = "end of CoverTab[57114]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:86
		_go_fuzz_dep_.CoverTab[57115]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:86
		// _ = "end of CoverTab[57115]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:86
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:86
	// _ = "end of CoverTab[57112]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:86
	_go_fuzz_dep_.CoverTab[57113]++
													return list.v.Type() == c.goType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:87
	// _ = "end of CoverTab[57113]"
}

func (c *listPtrConverter) IsValidGo(v reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:90
	_go_fuzz_dep_.CoverTab[57116]++
													return v.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:91
		_go_fuzz_dep_.CoverTab[57117]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:91
		return v.Type() == c.goType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:91
		// _ = "end of CoverTab[57117]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:91
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:91
	// _ = "end of CoverTab[57116]"
}

func (c *listPtrConverter) New() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:94
	_go_fuzz_dep_.CoverTab[57118]++
													return c.PBValueOf(reflect.New(c.goType.Elem()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:95
	// _ = "end of CoverTab[57118]"
}

func (c *listPtrConverter) Zero() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:98
	_go_fuzz_dep_.CoverTab[57119]++
													return c.PBValueOf(reflect.Zero(c.goType))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:99
	// _ = "end of CoverTab[57119]"
}

type listReflect struct {
	v	reflect.Value	// *[]T
	conv	Converter
}

func (ls *listReflect) Len() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:107
	_go_fuzz_dep_.CoverTab[57120]++
													if ls.v.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:108
		_go_fuzz_dep_.CoverTab[57122]++
														return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:109
		// _ = "end of CoverTab[57122]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:110
		_go_fuzz_dep_.CoverTab[57123]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:110
		// _ = "end of CoverTab[57123]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:110
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:110
	// _ = "end of CoverTab[57120]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:110
	_go_fuzz_dep_.CoverTab[57121]++
													return ls.v.Elem().Len()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:111
	// _ = "end of CoverTab[57121]"
}
func (ls *listReflect) Get(i int) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:113
	_go_fuzz_dep_.CoverTab[57124]++
													return ls.conv.PBValueOf(ls.v.Elem().Index(i))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:114
	// _ = "end of CoverTab[57124]"
}
func (ls *listReflect) Set(i int, v protoreflect.Value) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:116
	_go_fuzz_dep_.CoverTab[57125]++
													ls.v.Elem().Index(i).Set(ls.conv.GoValueOf(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:117
	// _ = "end of CoverTab[57125]"
}
func (ls *listReflect) Append(v protoreflect.Value) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:119
	_go_fuzz_dep_.CoverTab[57126]++
													ls.v.Elem().Set(reflect.Append(ls.v.Elem(), ls.conv.GoValueOf(v)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:120
	// _ = "end of CoverTab[57126]"
}
func (ls *listReflect) AppendMutable() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:122
	_go_fuzz_dep_.CoverTab[57127]++
													if _, ok := ls.conv.(*messageConverter); !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:123
		_go_fuzz_dep_.CoverTab[57129]++
														panic("invalid AppendMutable on list with non-message type")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:124
		// _ = "end of CoverTab[57129]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:125
		_go_fuzz_dep_.CoverTab[57130]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:125
		// _ = "end of CoverTab[57130]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:125
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:125
	// _ = "end of CoverTab[57127]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:125
	_go_fuzz_dep_.CoverTab[57128]++
													v := ls.NewElement()
													ls.Append(v)
													return v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:128
	// _ = "end of CoverTab[57128]"
}
func (ls *listReflect) Truncate(i int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:130
	_go_fuzz_dep_.CoverTab[57131]++
													ls.v.Elem().Set(ls.v.Elem().Slice(0, i))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:131
	// _ = "end of CoverTab[57131]"
}
func (ls *listReflect) NewElement() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:133
	_go_fuzz_dep_.CoverTab[57132]++
													return ls.conv.New()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:134
	// _ = "end of CoverTab[57132]"
}
func (ls *listReflect) IsValid() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:136
	_go_fuzz_dep_.CoverTab[57133]++
													return !ls.v.IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:137
	// _ = "end of CoverTab[57133]"
}
func (ls *listReflect) protoUnwrap() interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:139
	_go_fuzz_dep_.CoverTab[57134]++
													return ls.v.Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:140
	// _ = "end of CoverTab[57134]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:141
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/convert_list.go:141
var _ = _go_fuzz_dep_.CoverTab
