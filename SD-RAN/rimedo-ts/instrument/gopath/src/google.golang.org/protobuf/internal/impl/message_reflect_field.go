// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:5
)

import (
	"fmt"
	"math"
	"reflect"
	"sync"

	"google.golang.org/protobuf/internal/flags"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type fieldInfo struct {
	fieldDesc	protoreflect.FieldDescriptor

	// These fields are used for protobuf reflection support.
	has		func(pointer) bool
	clear		func(pointer)
	get		func(pointer) protoreflect.Value
	set		func(pointer, protoreflect.Value)
	mutable		func(pointer) protoreflect.Value
	newMessage	func() protoreflect.Message
	newField	func() protoreflect.Value
}

func fieldInfoForMissing(fd protoreflect.FieldDescriptor) fieldInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:31
	_go_fuzz_dep_.CoverTab[58305]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:35
	return fieldInfo{
		fieldDesc:	fd,
		has: func(p pointer) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:37
			_go_fuzz_dep_.CoverTab[58306]++
																return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:38
			// _ = "end of CoverTab[58306]"
		},
		clear: func(p pointer) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:40
			_go_fuzz_dep_.CoverTab[58307]++
																panic("missing Go struct field for " + string(fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:41
			// _ = "end of CoverTab[58307]"
		},
		get: func(p pointer) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:43
			_go_fuzz_dep_.CoverTab[58308]++
																return fd.Default()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:44
			// _ = "end of CoverTab[58308]"
		},
		set: func(p pointer, v protoreflect.Value) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:46
			_go_fuzz_dep_.CoverTab[58309]++
																panic("missing Go struct field for " + string(fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:47
			// _ = "end of CoverTab[58309]"
		},
		mutable: func(p pointer) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:49
			_go_fuzz_dep_.CoverTab[58310]++
																panic("missing Go struct field for " + string(fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:50
			// _ = "end of CoverTab[58310]"
		},
		newMessage: func() protoreflect.Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:52
			_go_fuzz_dep_.CoverTab[58311]++
																panic("missing Go struct field for " + string(fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:53
			// _ = "end of CoverTab[58311]"
		},
		newField: func() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:55
			_go_fuzz_dep_.CoverTab[58312]++
																if v := fd.Default(); v.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:56
				_go_fuzz_dep_.CoverTab[58314]++
																	return v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:57
				// _ = "end of CoverTab[58314]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:58
				_go_fuzz_dep_.CoverTab[58315]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:58
				// _ = "end of CoverTab[58315]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:58
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:58
			// _ = "end of CoverTab[58312]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:58
			_go_fuzz_dep_.CoverTab[58313]++
																panic("missing Go struct field for " + string(fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:59
			// _ = "end of CoverTab[58313]"
		},
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:61
	// _ = "end of CoverTab[58305]"
}

func fieldInfoForOneof(fd protoreflect.FieldDescriptor, fs reflect.StructField, x exporter, ot reflect.Type) fieldInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:64
	_go_fuzz_dep_.CoverTab[58316]++
														ft := fs.Type
														if ft.Kind() != reflect.Interface {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:66
		_go_fuzz_dep_.CoverTab[58320]++
															panic(fmt.Sprintf("field %v has invalid type: got %v, want interface kind", fd.FullName(), ft))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:67
		// _ = "end of CoverTab[58320]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:68
		_go_fuzz_dep_.CoverTab[58321]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:68
		// _ = "end of CoverTab[58321]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:68
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:68
	// _ = "end of CoverTab[58316]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:68
	_go_fuzz_dep_.CoverTab[58317]++
														if ot.Kind() != reflect.Struct {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:69
		_go_fuzz_dep_.CoverTab[58322]++
															panic(fmt.Sprintf("field %v has invalid type: got %v, want struct kind", fd.FullName(), ot))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:70
		// _ = "end of CoverTab[58322]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:71
		_go_fuzz_dep_.CoverTab[58323]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:71
		// _ = "end of CoverTab[58323]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:71
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:71
	// _ = "end of CoverTab[58317]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:71
	_go_fuzz_dep_.CoverTab[58318]++
														if !reflect.PtrTo(ot).Implements(ft) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:72
		_go_fuzz_dep_.CoverTab[58324]++
															panic(fmt.Sprintf("field %v has invalid type: %v does not implement %v", fd.FullName(), ot, ft))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:73
		// _ = "end of CoverTab[58324]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:74
		_go_fuzz_dep_.CoverTab[58325]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:74
		// _ = "end of CoverTab[58325]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:74
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:74
	// _ = "end of CoverTab[58318]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:74
	_go_fuzz_dep_.CoverTab[58319]++
														conv := NewConverter(ot.Field(0).Type, fd)
														isMessage := fd.Message() != nil

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:79
	fieldOffset := offsetOf(fs, x)
	return fieldInfo{

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:85
		fieldDesc:	fd,
		has: func(p pointer) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:86
			_go_fuzz_dep_.CoverTab[58326]++
																if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:87
				_go_fuzz_dep_.CoverTab[58329]++
																	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:88
				// _ = "end of CoverTab[58329]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:89
				_go_fuzz_dep_.CoverTab[58330]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:89
				// _ = "end of CoverTab[58330]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:89
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:89
			// _ = "end of CoverTab[58326]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:89
			_go_fuzz_dep_.CoverTab[58327]++
																rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																if rv.IsNil() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:91
				_go_fuzz_dep_.CoverTab[58331]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:91
				return rv.Elem().Type().Elem() != ot
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:91
				// _ = "end of CoverTab[58331]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:91
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:91
				_go_fuzz_dep_.CoverTab[58332]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:91
				return rv.Elem().IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:91
				// _ = "end of CoverTab[58332]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:91
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:91
				_go_fuzz_dep_.CoverTab[58333]++
																	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:92
				// _ = "end of CoverTab[58333]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:93
				_go_fuzz_dep_.CoverTab[58334]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:93
				// _ = "end of CoverTab[58334]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:93
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:93
			// _ = "end of CoverTab[58327]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:93
			_go_fuzz_dep_.CoverTab[58328]++
																return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:94
			// _ = "end of CoverTab[58328]"
		},
		clear: func(p pointer) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:96
			_go_fuzz_dep_.CoverTab[58335]++
																rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																if rv.IsNil() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:98
				_go_fuzz_dep_.CoverTab[58337]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:98
				return rv.Elem().Type().Elem() != ot
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:98
				// _ = "end of CoverTab[58337]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:98
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:98
					_go_fuzz_dep_.CoverTab[58338]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:101
				return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:101
				// _ = "end of CoverTab[58338]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:102
				_go_fuzz_dep_.CoverTab[58339]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:102
				// _ = "end of CoverTab[58339]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:102
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:102
			// _ = "end of CoverTab[58335]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:102
			_go_fuzz_dep_.CoverTab[58336]++
																	rv.Set(reflect.Zero(rv.Type()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:103
			// _ = "end of CoverTab[58336]"
		},
		get: func(p pointer) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:105
			_go_fuzz_dep_.CoverTab[58340]++
																	if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:106
				_go_fuzz_dep_.CoverTab[58343]++
																		return conv.Zero()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:107
				// _ = "end of CoverTab[58343]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:108
				_go_fuzz_dep_.CoverTab[58344]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:108
				// _ = "end of CoverTab[58344]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:108
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:108
			// _ = "end of CoverTab[58340]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:108
			_go_fuzz_dep_.CoverTab[58341]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	if rv.IsNil() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:110
				_go_fuzz_dep_.CoverTab[58345]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:110
				return rv.Elem().Type().Elem() != ot
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:110
				// _ = "end of CoverTab[58345]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:110
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:110
				_go_fuzz_dep_.CoverTab[58346]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:110
				return rv.Elem().IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:110
				// _ = "end of CoverTab[58346]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:110
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:110
				_go_fuzz_dep_.CoverTab[58347]++
																		return conv.Zero()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:111
				// _ = "end of CoverTab[58347]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:112
				_go_fuzz_dep_.CoverTab[58348]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:112
				// _ = "end of CoverTab[58348]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:112
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:112
			// _ = "end of CoverTab[58341]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:112
			_go_fuzz_dep_.CoverTab[58342]++
																	rv = rv.Elem().Elem().Field(0)
																	return conv.PBValueOf(rv)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:114
			// _ = "end of CoverTab[58342]"
		},
		set: func(p pointer, v protoreflect.Value) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:116
			_go_fuzz_dep_.CoverTab[58349]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	if rv.IsNil() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:118
				_go_fuzz_dep_.CoverTab[58351]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:118
				return rv.Elem().Type().Elem() != ot
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:118
				// _ = "end of CoverTab[58351]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:118
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:118
				_go_fuzz_dep_.CoverTab[58352]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:118
				return rv.Elem().IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:118
				// _ = "end of CoverTab[58352]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:118
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:118
				_go_fuzz_dep_.CoverTab[58353]++
																		rv.Set(reflect.New(ot))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:119
				// _ = "end of CoverTab[58353]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:120
				_go_fuzz_dep_.CoverTab[58354]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:120
				// _ = "end of CoverTab[58354]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:120
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:120
			// _ = "end of CoverTab[58349]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:120
			_go_fuzz_dep_.CoverTab[58350]++
																	rv = rv.Elem().Elem().Field(0)
																	rv.Set(conv.GoValueOf(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:122
			// _ = "end of CoverTab[58350]"
		},
		mutable: func(p pointer) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:124
			_go_fuzz_dep_.CoverTab[58355]++
																	if !isMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:125
				_go_fuzz_dep_.CoverTab[58359]++
																		panic(fmt.Sprintf("field %v with invalid Mutable call on field with non-composite type", fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:126
				// _ = "end of CoverTab[58359]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:127
				_go_fuzz_dep_.CoverTab[58360]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:127
				// _ = "end of CoverTab[58360]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:127
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:127
			// _ = "end of CoverTab[58355]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:127
			_go_fuzz_dep_.CoverTab[58356]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	if rv.IsNil() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:129
				_go_fuzz_dep_.CoverTab[58361]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:129
				return rv.Elem().Type().Elem() != ot
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:129
				// _ = "end of CoverTab[58361]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:129
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:129
				_go_fuzz_dep_.CoverTab[58362]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:129
				return rv.Elem().IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:129
				// _ = "end of CoverTab[58362]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:129
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:129
				_go_fuzz_dep_.CoverTab[58363]++
																		rv.Set(reflect.New(ot))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:130
				// _ = "end of CoverTab[58363]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:131
				_go_fuzz_dep_.CoverTab[58364]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:131
				// _ = "end of CoverTab[58364]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:131
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:131
			// _ = "end of CoverTab[58356]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:131
			_go_fuzz_dep_.CoverTab[58357]++
																	rv = rv.Elem().Elem().Field(0)
																	if rv.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:133
				_go_fuzz_dep_.CoverTab[58365]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:133
				return rv.IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:133
				// _ = "end of CoverTab[58365]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:133
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:133
				_go_fuzz_dep_.CoverTab[58366]++
																		rv.Set(conv.GoValueOf(protoreflect.ValueOfMessage(conv.New().Message())))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:134
				// _ = "end of CoverTab[58366]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:135
				_go_fuzz_dep_.CoverTab[58367]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:135
				// _ = "end of CoverTab[58367]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:135
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:135
			// _ = "end of CoverTab[58357]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:135
			_go_fuzz_dep_.CoverTab[58358]++
																	return conv.PBValueOf(rv)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:136
			// _ = "end of CoverTab[58358]"
		},
		newMessage: func() protoreflect.Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:138
			_go_fuzz_dep_.CoverTab[58368]++
																	return conv.New().Message()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:139
			// _ = "end of CoverTab[58368]"
		},
		newField: func() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:141
			_go_fuzz_dep_.CoverTab[58369]++
																	return conv.New()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:142
			// _ = "end of CoverTab[58369]"
		},
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:144
	// _ = "end of CoverTab[58319]"
}

func fieldInfoForMap(fd protoreflect.FieldDescriptor, fs reflect.StructField, x exporter) fieldInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:147
	_go_fuzz_dep_.CoverTab[58370]++
															ft := fs.Type
															if ft.Kind() != reflect.Map {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:149
		_go_fuzz_dep_.CoverTab[58372]++
																panic(fmt.Sprintf("field %v has invalid type: got %v, want map kind", fd.FullName(), ft))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:150
		// _ = "end of CoverTab[58372]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:151
		_go_fuzz_dep_.CoverTab[58373]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:151
		// _ = "end of CoverTab[58373]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:151
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:151
	// _ = "end of CoverTab[58370]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:151
	_go_fuzz_dep_.CoverTab[58371]++
															conv := NewConverter(ft, fd)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:155
	fieldOffset := offsetOf(fs, x)
	return fieldInfo{
		fieldDesc:	fd,
		has: func(p pointer) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:158
			_go_fuzz_dep_.CoverTab[58374]++
																	if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:159
				_go_fuzz_dep_.CoverTab[58376]++
																		return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:160
				// _ = "end of CoverTab[58376]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:161
				_go_fuzz_dep_.CoverTab[58377]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:161
				// _ = "end of CoverTab[58377]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:161
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:161
			// _ = "end of CoverTab[58374]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:161
			_go_fuzz_dep_.CoverTab[58375]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	return rv.Len() > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:163
			// _ = "end of CoverTab[58375]"
		},
		clear: func(p pointer) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:165
			_go_fuzz_dep_.CoverTab[58378]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	rv.Set(reflect.Zero(rv.Type()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:167
			// _ = "end of CoverTab[58378]"
		},
		get: func(p pointer) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:169
			_go_fuzz_dep_.CoverTab[58379]++
																	if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:170
				_go_fuzz_dep_.CoverTab[58382]++
																		return conv.Zero()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:171
				// _ = "end of CoverTab[58382]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:172
				_go_fuzz_dep_.CoverTab[58383]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:172
				// _ = "end of CoverTab[58383]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:172
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:172
			// _ = "end of CoverTab[58379]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:172
			_go_fuzz_dep_.CoverTab[58380]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	if rv.Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:174
				_go_fuzz_dep_.CoverTab[58384]++
																		return conv.Zero()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:175
				// _ = "end of CoverTab[58384]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:176
				_go_fuzz_dep_.CoverTab[58385]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:176
				// _ = "end of CoverTab[58385]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:176
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:176
			// _ = "end of CoverTab[58380]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:176
			_go_fuzz_dep_.CoverTab[58381]++
																	return conv.PBValueOf(rv)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:177
			// _ = "end of CoverTab[58381]"
		},
		set: func(p pointer, v protoreflect.Value) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:179
			_go_fuzz_dep_.CoverTab[58386]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	pv := conv.GoValueOf(v)
																	if pv.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:182
				_go_fuzz_dep_.CoverTab[58388]++
																		panic(fmt.Sprintf("map field %v cannot be set with read-only value", fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:183
				// _ = "end of CoverTab[58388]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:184
				_go_fuzz_dep_.CoverTab[58389]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:184
				// _ = "end of CoverTab[58389]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:184
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:184
			// _ = "end of CoverTab[58386]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:184
			_go_fuzz_dep_.CoverTab[58387]++
																	rv.Set(pv)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:185
			// _ = "end of CoverTab[58387]"
		},
		mutable: func(p pointer) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:187
			_go_fuzz_dep_.CoverTab[58390]++
																	v := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	if v.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:189
				_go_fuzz_dep_.CoverTab[58392]++
																		v.Set(reflect.MakeMap(fs.Type))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:190
				// _ = "end of CoverTab[58392]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:191
				_go_fuzz_dep_.CoverTab[58393]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:191
				// _ = "end of CoverTab[58393]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:191
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:191
			// _ = "end of CoverTab[58390]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:191
			_go_fuzz_dep_.CoverTab[58391]++
																	return conv.PBValueOf(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:192
			// _ = "end of CoverTab[58391]"
		},
		newField: func() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:194
			_go_fuzz_dep_.CoverTab[58394]++
																	return conv.New()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:195
			// _ = "end of CoverTab[58394]"
		},
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:197
	// _ = "end of CoverTab[58371]"
}

func fieldInfoForList(fd protoreflect.FieldDescriptor, fs reflect.StructField, x exporter) fieldInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:200
	_go_fuzz_dep_.CoverTab[58395]++
															ft := fs.Type
															if ft.Kind() != reflect.Slice {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:202
		_go_fuzz_dep_.CoverTab[58397]++
																panic(fmt.Sprintf("field %v has invalid type: got %v, want slice kind", fd.FullName(), ft))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:203
		// _ = "end of CoverTab[58397]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:204
		_go_fuzz_dep_.CoverTab[58398]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:204
		// _ = "end of CoverTab[58398]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:204
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:204
	// _ = "end of CoverTab[58395]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:204
	_go_fuzz_dep_.CoverTab[58396]++
															conv := NewConverter(reflect.PtrTo(ft), fd)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:208
	fieldOffset := offsetOf(fs, x)
	return fieldInfo{
		fieldDesc:	fd,
		has: func(p pointer) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:211
			_go_fuzz_dep_.CoverTab[58399]++
																	if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:212
				_go_fuzz_dep_.CoverTab[58401]++
																		return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:213
				// _ = "end of CoverTab[58401]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:214
				_go_fuzz_dep_.CoverTab[58402]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:214
				// _ = "end of CoverTab[58402]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:214
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:214
			// _ = "end of CoverTab[58399]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:214
			_go_fuzz_dep_.CoverTab[58400]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	return rv.Len() > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:216
			// _ = "end of CoverTab[58400]"
		},
		clear: func(p pointer) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:218
			_go_fuzz_dep_.CoverTab[58403]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	rv.Set(reflect.Zero(rv.Type()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:220
			// _ = "end of CoverTab[58403]"
		},
		get: func(p pointer) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:222
			_go_fuzz_dep_.CoverTab[58404]++
																	if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:223
				_go_fuzz_dep_.CoverTab[58407]++
																		return conv.Zero()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:224
				// _ = "end of CoverTab[58407]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:225
				_go_fuzz_dep_.CoverTab[58408]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:225
				// _ = "end of CoverTab[58408]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:225
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:225
			// _ = "end of CoverTab[58404]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:225
			_go_fuzz_dep_.CoverTab[58405]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type)
																	if rv.Elem().Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:227
				_go_fuzz_dep_.CoverTab[58409]++
																		return conv.Zero()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:228
				// _ = "end of CoverTab[58409]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:229
				_go_fuzz_dep_.CoverTab[58410]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:229
				// _ = "end of CoverTab[58410]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:229
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:229
			// _ = "end of CoverTab[58405]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:229
			_go_fuzz_dep_.CoverTab[58406]++
																	return conv.PBValueOf(rv)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:230
			// _ = "end of CoverTab[58406]"
		},
		set: func(p pointer, v protoreflect.Value) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:232
			_go_fuzz_dep_.CoverTab[58411]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	pv := conv.GoValueOf(v)
																	if pv.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:235
				_go_fuzz_dep_.CoverTab[58413]++
																		panic(fmt.Sprintf("list field %v cannot be set with read-only value", fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:236
				// _ = "end of CoverTab[58413]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:237
				_go_fuzz_dep_.CoverTab[58414]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:237
				// _ = "end of CoverTab[58414]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:237
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:237
			// _ = "end of CoverTab[58411]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:237
			_go_fuzz_dep_.CoverTab[58412]++
																	rv.Set(pv.Elem())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:238
			// _ = "end of CoverTab[58412]"
		},
		mutable: func(p pointer) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:240
			_go_fuzz_dep_.CoverTab[58415]++
																	v := p.Apply(fieldOffset).AsValueOf(fs.Type)
																	return conv.PBValueOf(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:242
			// _ = "end of CoverTab[58415]"
		},
		newField: func() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:244
			_go_fuzz_dep_.CoverTab[58416]++
																	return conv.New()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:245
			// _ = "end of CoverTab[58416]"
		},
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:247
	// _ = "end of CoverTab[58396]"
}

var (
	nilBytes	= reflect.ValueOf([]byte(nil))
	emptyBytes	= reflect.ValueOf([]byte{})
)

func fieldInfoForScalar(fd protoreflect.FieldDescriptor, fs reflect.StructField, x exporter) fieldInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:255
	_go_fuzz_dep_.CoverTab[58417]++
															ft := fs.Type
															nullable := fd.HasPresence()
															isBytes := ft.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:258
		_go_fuzz_dep_.CoverTab[58419]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:258
		return ft.Elem().Kind() == reflect.Uint8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:258
		// _ = "end of CoverTab[58419]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:258
	}()
															if nullable {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:259
		_go_fuzz_dep_.CoverTab[58420]++
																if ft.Kind() != reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:260
			_go_fuzz_dep_.CoverTab[58422]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:260
			return ft.Kind() != reflect.Slice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:260
			// _ = "end of CoverTab[58422]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:260
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:260
			_go_fuzz_dep_.CoverTab[58423]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:264
			nullable = false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:264
			// _ = "end of CoverTab[58423]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:265
			_go_fuzz_dep_.CoverTab[58424]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:265
			// _ = "end of CoverTab[58424]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:265
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:265
		// _ = "end of CoverTab[58420]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:265
		_go_fuzz_dep_.CoverTab[58421]++
																if ft.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:266
			_go_fuzz_dep_.CoverTab[58425]++
																	ft = ft.Elem()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:267
			// _ = "end of CoverTab[58425]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:268
			_go_fuzz_dep_.CoverTab[58426]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:268
			// _ = "end of CoverTab[58426]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:268
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:268
		// _ = "end of CoverTab[58421]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:269
		_go_fuzz_dep_.CoverTab[58427]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:269
		// _ = "end of CoverTab[58427]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:269
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:269
	// _ = "end of CoverTab[58417]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:269
	_go_fuzz_dep_.CoverTab[58418]++
															conv := NewConverter(ft, fd)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:273
	fieldOffset := offsetOf(fs, x)
	return fieldInfo{
		fieldDesc:	fd,
		has: func(p pointer) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:276
			_go_fuzz_dep_.CoverTab[58428]++
																	if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:277
				_go_fuzz_dep_.CoverTab[58431]++
																		return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:278
				// _ = "end of CoverTab[58431]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:279
				_go_fuzz_dep_.CoverTab[58432]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:279
				// _ = "end of CoverTab[58432]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:279
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:279
			// _ = "end of CoverTab[58428]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:279
			_go_fuzz_dep_.CoverTab[58429]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	if nullable {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:281
				_go_fuzz_dep_.CoverTab[58433]++
																		return !rv.IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:282
				// _ = "end of CoverTab[58433]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:283
				_go_fuzz_dep_.CoverTab[58434]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:283
				// _ = "end of CoverTab[58434]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:283
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:283
			// _ = "end of CoverTab[58429]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:283
			_go_fuzz_dep_.CoverTab[58430]++
																	switch rv.Kind() {
			case reflect.Bool:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:285
				_go_fuzz_dep_.CoverTab[58435]++
																		return rv.Bool()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:286
				// _ = "end of CoverTab[58435]"
			case reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:287
				_go_fuzz_dep_.CoverTab[58436]++
																		return rv.Int() != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:288
				// _ = "end of CoverTab[58436]"
			case reflect.Uint32, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:289
				_go_fuzz_dep_.CoverTab[58437]++
																		return rv.Uint() != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:290
				// _ = "end of CoverTab[58437]"
			case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:291
				_go_fuzz_dep_.CoverTab[58438]++
																		return rv.Float() != 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:292
					_go_fuzz_dep_.CoverTab[58441]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:292
					return math.Signbit(rv.Float())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:292
					// _ = "end of CoverTab[58441]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:292
				}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:292
				// _ = "end of CoverTab[58438]"
			case reflect.String, reflect.Slice:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:293
				_go_fuzz_dep_.CoverTab[58439]++
																		return rv.Len() > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:294
				// _ = "end of CoverTab[58439]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:295
				_go_fuzz_dep_.CoverTab[58440]++
																		panic(fmt.Sprintf("field %v has invalid type: %v", fd.FullName(), rv.Type()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:296
				// _ = "end of CoverTab[58440]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:297
			// _ = "end of CoverTab[58430]"
		},
		clear: func(p pointer) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:299
			_go_fuzz_dep_.CoverTab[58442]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	rv.Set(reflect.Zero(rv.Type()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:301
			// _ = "end of CoverTab[58442]"
		},
		get: func(p pointer) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:303
			_go_fuzz_dep_.CoverTab[58443]++
																	if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:304
				_go_fuzz_dep_.CoverTab[58446]++
																		return conv.Zero()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:305
				// _ = "end of CoverTab[58446]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:306
				_go_fuzz_dep_.CoverTab[58447]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:306
				// _ = "end of CoverTab[58447]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:306
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:306
			// _ = "end of CoverTab[58443]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:306
			_go_fuzz_dep_.CoverTab[58444]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	if nullable {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:308
				_go_fuzz_dep_.CoverTab[58448]++
																		if rv.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:309
					_go_fuzz_dep_.CoverTab[58450]++
																			return conv.Zero()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:310
					// _ = "end of CoverTab[58450]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:311
					_go_fuzz_dep_.CoverTab[58451]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:311
					// _ = "end of CoverTab[58451]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:311
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:311
				// _ = "end of CoverTab[58448]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:311
				_go_fuzz_dep_.CoverTab[58449]++
																		if rv.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:312
					_go_fuzz_dep_.CoverTab[58452]++
																			rv = rv.Elem()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:313
					// _ = "end of CoverTab[58452]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:314
					_go_fuzz_dep_.CoverTab[58453]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:314
					// _ = "end of CoverTab[58453]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:314
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:314
				// _ = "end of CoverTab[58449]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:315
				_go_fuzz_dep_.CoverTab[58454]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:315
				// _ = "end of CoverTab[58454]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:315
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:315
			// _ = "end of CoverTab[58444]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:315
			_go_fuzz_dep_.CoverTab[58445]++
																	return conv.PBValueOf(rv)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:316
			// _ = "end of CoverTab[58445]"
		},
		set: func(p pointer, v protoreflect.Value) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:318
			_go_fuzz_dep_.CoverTab[58455]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	if nullable && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:320
				_go_fuzz_dep_.CoverTab[58457]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:320
				return rv.Kind() == reflect.Ptr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:320
				// _ = "end of CoverTab[58457]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:320
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:320
				_go_fuzz_dep_.CoverTab[58458]++
																		if rv.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:321
					_go_fuzz_dep_.CoverTab[58460]++
																			rv.Set(reflect.New(ft))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:322
					// _ = "end of CoverTab[58460]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:323
					_go_fuzz_dep_.CoverTab[58461]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:323
					// _ = "end of CoverTab[58461]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:323
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:323
				// _ = "end of CoverTab[58458]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:323
				_go_fuzz_dep_.CoverTab[58459]++
																		rv = rv.Elem()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:324
				// _ = "end of CoverTab[58459]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:325
				_go_fuzz_dep_.CoverTab[58462]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:325
				// _ = "end of CoverTab[58462]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:325
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:325
			// _ = "end of CoverTab[58455]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:325
			_go_fuzz_dep_.CoverTab[58456]++
																	rv.Set(conv.GoValueOf(v))
																	if isBytes && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:327
				_go_fuzz_dep_.CoverTab[58463]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:327
				return rv.Len() == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:327
				// _ = "end of CoverTab[58463]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:327
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:327
				_go_fuzz_dep_.CoverTab[58464]++
																		if nullable {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:328
					_go_fuzz_dep_.CoverTab[58465]++
																			rv.Set(emptyBytes)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:329
					// _ = "end of CoverTab[58465]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:330
					_go_fuzz_dep_.CoverTab[58466]++
																			rv.Set(nilBytes)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:331
					// _ = "end of CoverTab[58466]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:332
				// _ = "end of CoverTab[58464]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:333
				_go_fuzz_dep_.CoverTab[58467]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:333
				// _ = "end of CoverTab[58467]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:333
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:333
			// _ = "end of CoverTab[58456]"
		},
		newField: func() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:335
			_go_fuzz_dep_.CoverTab[58468]++
																	return conv.New()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:336
			// _ = "end of CoverTab[58468]"
		},
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:338
	// _ = "end of CoverTab[58418]"
}

func fieldInfoForWeakMessage(fd protoreflect.FieldDescriptor, weakOffset offset) fieldInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:341
	_go_fuzz_dep_.CoverTab[58469]++
															if !flags.ProtoLegacy {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:342
		_go_fuzz_dep_.CoverTab[58472]++
																panic("no support for proto1 weak fields")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:343
		// _ = "end of CoverTab[58472]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:344
		_go_fuzz_dep_.CoverTab[58473]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:344
		// _ = "end of CoverTab[58473]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:344
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:344
	// _ = "end of CoverTab[58469]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:344
	_go_fuzz_dep_.CoverTab[58470]++

															var once sync.Once
															var messageType protoreflect.MessageType
															lazyInit := func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:348
		_go_fuzz_dep_.CoverTab[58474]++
																once.Do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:349
			_go_fuzz_dep_.CoverTab[58475]++
																	messageName := fd.Message().FullName()
																	messageType, _ = protoregistry.GlobalTypes.FindMessageByName(messageName)
																	if messageType == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:352
				_go_fuzz_dep_.CoverTab[58476]++
																		panic(fmt.Sprintf("weak message %v for field %v is not linked in", messageName, fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:353
				// _ = "end of CoverTab[58476]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:354
				_go_fuzz_dep_.CoverTab[58477]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:354
				// _ = "end of CoverTab[58477]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:354
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:354
			// _ = "end of CoverTab[58475]"
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:355
		// _ = "end of CoverTab[58474]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:356
	// _ = "end of CoverTab[58470]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:356
	_go_fuzz_dep_.CoverTab[58471]++

															num := fd.Number()
															return fieldInfo{
		fieldDesc:	fd,
		has: func(p pointer) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:361
			_go_fuzz_dep_.CoverTab[58478]++
																	if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:362
				_go_fuzz_dep_.CoverTab[58480]++
																		return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:363
				// _ = "end of CoverTab[58480]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:364
				_go_fuzz_dep_.CoverTab[58481]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:364
				// _ = "end of CoverTab[58481]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:364
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:364
			// _ = "end of CoverTab[58478]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:364
			_go_fuzz_dep_.CoverTab[58479]++
																	_, ok := p.Apply(weakOffset).WeakFields().get(num)
																	return ok
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:366
			// _ = "end of CoverTab[58479]"
		},
		clear: func(p pointer) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:368
			_go_fuzz_dep_.CoverTab[58482]++
																	p.Apply(weakOffset).WeakFields().clear(num)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:369
			// _ = "end of CoverTab[58482]"
		},
		get: func(p pointer) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:371
			_go_fuzz_dep_.CoverTab[58483]++
																	lazyInit()
																	if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:373
				_go_fuzz_dep_.CoverTab[58486]++
																		return protoreflect.ValueOfMessage(messageType.Zero())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:374
				// _ = "end of CoverTab[58486]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:375
				_go_fuzz_dep_.CoverTab[58487]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:375
				// _ = "end of CoverTab[58487]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:375
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:375
			// _ = "end of CoverTab[58483]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:375
			_go_fuzz_dep_.CoverTab[58484]++
																	m, ok := p.Apply(weakOffset).WeakFields().get(num)
																	if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:377
				_go_fuzz_dep_.CoverTab[58488]++
																		return protoreflect.ValueOfMessage(messageType.Zero())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:378
				// _ = "end of CoverTab[58488]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:379
				_go_fuzz_dep_.CoverTab[58489]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:379
				// _ = "end of CoverTab[58489]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:379
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:379
			// _ = "end of CoverTab[58484]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:379
			_go_fuzz_dep_.CoverTab[58485]++
																	return protoreflect.ValueOfMessage(m.ProtoReflect())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:380
			// _ = "end of CoverTab[58485]"
		},
		set: func(p pointer, v protoreflect.Value) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:382
			_go_fuzz_dep_.CoverTab[58490]++
																	lazyInit()
																	m := v.Message()
																	if m.Descriptor() != messageType.Descriptor() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:385
				_go_fuzz_dep_.CoverTab[58492]++
																		if got, want := m.Descriptor().FullName(), messageType.Descriptor().FullName(); got != want {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:386
					_go_fuzz_dep_.CoverTab[58494]++
																			panic(fmt.Sprintf("field %v has mismatching message descriptor: got %v, want %v", fd.FullName(), got, want))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:387
					// _ = "end of CoverTab[58494]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:388
					_go_fuzz_dep_.CoverTab[58495]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:388
					// _ = "end of CoverTab[58495]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:388
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:388
				// _ = "end of CoverTab[58492]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:388
				_go_fuzz_dep_.CoverTab[58493]++
																		panic(fmt.Sprintf("field %v has mismatching message descriptor: %v", fd.FullName(), m.Descriptor().FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:389
				// _ = "end of CoverTab[58493]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:390
				_go_fuzz_dep_.CoverTab[58496]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:390
				// _ = "end of CoverTab[58496]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:390
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:390
			// _ = "end of CoverTab[58490]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:390
			_go_fuzz_dep_.CoverTab[58491]++
																	p.Apply(weakOffset).WeakFields().set(num, m.Interface())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:391
			// _ = "end of CoverTab[58491]"
		},
		mutable: func(p pointer) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:393
			_go_fuzz_dep_.CoverTab[58497]++
																	lazyInit()
																	fs := p.Apply(weakOffset).WeakFields()
																	m, ok := fs.get(num)
																	if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:397
				_go_fuzz_dep_.CoverTab[58499]++
																		m = messageType.New().Interface()
																		fs.set(num, m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:399
				// _ = "end of CoverTab[58499]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:400
				_go_fuzz_dep_.CoverTab[58500]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:400
				// _ = "end of CoverTab[58500]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:400
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:400
			// _ = "end of CoverTab[58497]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:400
			_go_fuzz_dep_.CoverTab[58498]++
																	return protoreflect.ValueOfMessage(m.ProtoReflect())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:401
			// _ = "end of CoverTab[58498]"
		},
		newMessage: func() protoreflect.Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:403
			_go_fuzz_dep_.CoverTab[58501]++
																	lazyInit()
																	return messageType.New()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:405
			// _ = "end of CoverTab[58501]"
		},
		newField: func() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:407
			_go_fuzz_dep_.CoverTab[58502]++
																	lazyInit()
																	return protoreflect.ValueOfMessage(messageType.New())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:409
			// _ = "end of CoverTab[58502]"
		},
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:411
	// _ = "end of CoverTab[58471]"
}

func fieldInfoForMessage(fd protoreflect.FieldDescriptor, fs reflect.StructField, x exporter) fieldInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:414
	_go_fuzz_dep_.CoverTab[58503]++
															ft := fs.Type
															conv := NewConverter(ft, fd)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:419
	fieldOffset := offsetOf(fs, x)
	return fieldInfo{
		fieldDesc:	fd,
		has: func(p pointer) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:422
			_go_fuzz_dep_.CoverTab[58504]++
																	if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:423
				_go_fuzz_dep_.CoverTab[58507]++
																		return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:424
				// _ = "end of CoverTab[58507]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:425
				_go_fuzz_dep_.CoverTab[58508]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:425
				// _ = "end of CoverTab[58508]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:425
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:425
			// _ = "end of CoverTab[58504]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:425
			_go_fuzz_dep_.CoverTab[58505]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	if fs.Type.Kind() != reflect.Ptr {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:427
				_go_fuzz_dep_.CoverTab[58509]++
																		return !isZero(rv)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:428
				// _ = "end of CoverTab[58509]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:429
				_go_fuzz_dep_.CoverTab[58510]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:429
				// _ = "end of CoverTab[58510]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:429
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:429
			// _ = "end of CoverTab[58505]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:429
			_go_fuzz_dep_.CoverTab[58506]++
																	return !rv.IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:430
			// _ = "end of CoverTab[58506]"
		},
		clear: func(p pointer) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:432
			_go_fuzz_dep_.CoverTab[58511]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	rv.Set(reflect.Zero(rv.Type()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:434
			// _ = "end of CoverTab[58511]"
		},
		get: func(p pointer) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:436
			_go_fuzz_dep_.CoverTab[58512]++
																	if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:437
				_go_fuzz_dep_.CoverTab[58514]++
																		return conv.Zero()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:438
				// _ = "end of CoverTab[58514]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:439
				_go_fuzz_dep_.CoverTab[58515]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:439
				// _ = "end of CoverTab[58515]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:439
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:439
			// _ = "end of CoverTab[58512]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:439
			_go_fuzz_dep_.CoverTab[58513]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	return conv.PBValueOf(rv)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:441
			// _ = "end of CoverTab[58513]"
		},
		set: func(p pointer, v protoreflect.Value) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:443
			_go_fuzz_dep_.CoverTab[58516]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	rv.Set(conv.GoValueOf(v))
																	if fs.Type.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:446
				_go_fuzz_dep_.CoverTab[58517]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:446
				return rv.IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:446
				// _ = "end of CoverTab[58517]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:446
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:446
				_go_fuzz_dep_.CoverTab[58518]++
																		panic(fmt.Sprintf("field %v has invalid nil pointer", fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:447
				// _ = "end of CoverTab[58518]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:448
				_go_fuzz_dep_.CoverTab[58519]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:448
				// _ = "end of CoverTab[58519]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:448
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:448
			// _ = "end of CoverTab[58516]"
		},
		mutable: func(p pointer) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:450
			_go_fuzz_dep_.CoverTab[58520]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	if fs.Type.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:452
				_go_fuzz_dep_.CoverTab[58522]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:452
				return rv.IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:452
				// _ = "end of CoverTab[58522]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:452
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:452
				_go_fuzz_dep_.CoverTab[58523]++
																		rv.Set(conv.GoValueOf(conv.New()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:453
				// _ = "end of CoverTab[58523]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:454
				_go_fuzz_dep_.CoverTab[58524]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:454
				// _ = "end of CoverTab[58524]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:454
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:454
			// _ = "end of CoverTab[58520]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:454
			_go_fuzz_dep_.CoverTab[58521]++
																	return conv.PBValueOf(rv)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:455
			// _ = "end of CoverTab[58521]"
		},
		newMessage: func() protoreflect.Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:457
			_go_fuzz_dep_.CoverTab[58525]++
																	return conv.New().Message()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:458
			// _ = "end of CoverTab[58525]"
		},
		newField: func() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:460
			_go_fuzz_dep_.CoverTab[58526]++
																	return conv.New()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:461
			// _ = "end of CoverTab[58526]"
		},
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:463
	// _ = "end of CoverTab[58503]"
}

type oneofInfo struct {
	oneofDesc	protoreflect.OneofDescriptor
	which		func(pointer) protoreflect.FieldNumber
}

func makeOneofInfo(od protoreflect.OneofDescriptor, si structInfo, x exporter) *oneofInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:471
	_go_fuzz_dep_.CoverTab[58527]++
															oi := &oneofInfo{oneofDesc: od}
															if od.IsSynthetic() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:473
		_go_fuzz_dep_.CoverTab[58529]++
																fs := si.fieldsByNumber[od.Fields().Get(0).Number()]
																fieldOffset := offsetOf(fs, x)
																oi.which = func(p pointer) protoreflect.FieldNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:476
			_go_fuzz_dep_.CoverTab[58530]++
																	if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:477
				_go_fuzz_dep_.CoverTab[58533]++
																		return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:478
				// _ = "end of CoverTab[58533]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:479
				_go_fuzz_dep_.CoverTab[58534]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:479
				// _ = "end of CoverTab[58534]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:479
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:479
			// _ = "end of CoverTab[58530]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:479
			_go_fuzz_dep_.CoverTab[58531]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	if rv.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:481
				_go_fuzz_dep_.CoverTab[58535]++
																		return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:482
				// _ = "end of CoverTab[58535]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:483
				_go_fuzz_dep_.CoverTab[58536]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:483
				// _ = "end of CoverTab[58536]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:483
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:483
			// _ = "end of CoverTab[58531]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:483
			_go_fuzz_dep_.CoverTab[58532]++
																	return od.Fields().Get(0).Number()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:484
			// _ = "end of CoverTab[58532]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:485
		// _ = "end of CoverTab[58529]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:486
		_go_fuzz_dep_.CoverTab[58537]++
																fs := si.oneofsByName[od.Name()]
																fieldOffset := offsetOf(fs, x)
																oi.which = func(p pointer) protoreflect.FieldNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:489
			_go_fuzz_dep_.CoverTab[58538]++
																	if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:490
				_go_fuzz_dep_.CoverTab[58542]++
																		return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:491
				// _ = "end of CoverTab[58542]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:492
				_go_fuzz_dep_.CoverTab[58543]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:492
				// _ = "end of CoverTab[58543]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:492
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:492
			// _ = "end of CoverTab[58538]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:492
			_go_fuzz_dep_.CoverTab[58539]++
																	rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
																	if rv.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:494
				_go_fuzz_dep_.CoverTab[58544]++
																		return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:495
				// _ = "end of CoverTab[58544]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:496
				_go_fuzz_dep_.CoverTab[58545]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:496
				// _ = "end of CoverTab[58545]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:496
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:496
			// _ = "end of CoverTab[58539]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:496
			_go_fuzz_dep_.CoverTab[58540]++
																	rv = rv.Elem()
																	if rv.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:498
				_go_fuzz_dep_.CoverTab[58546]++
																		return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:499
				// _ = "end of CoverTab[58546]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:500
				_go_fuzz_dep_.CoverTab[58547]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:500
				// _ = "end of CoverTab[58547]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:500
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:500
			// _ = "end of CoverTab[58540]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:500
			_go_fuzz_dep_.CoverTab[58541]++
																	return si.oneofWrappersByType[rv.Type().Elem()]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:501
			// _ = "end of CoverTab[58541]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:502
		// _ = "end of CoverTab[58537]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:503
	// _ = "end of CoverTab[58527]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:503
	_go_fuzz_dep_.CoverTab[58528]++
															return oi
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:504
	// _ = "end of CoverTab[58528]"
}

// isZero is identical to reflect.Value.IsZero.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:507
// TODO: Remove this when Go1.13 is the minimally supported Go version.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:509
func isZero(v reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:509
	_go_fuzz_dep_.CoverTab[58548]++
															switch v.Kind() {
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:511
		_go_fuzz_dep_.CoverTab[58549]++
																return !v.Bool()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:512
		// _ = "end of CoverTab[58549]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:513
		_go_fuzz_dep_.CoverTab[58550]++
																return v.Int() == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:514
		// _ = "end of CoverTab[58550]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:515
		_go_fuzz_dep_.CoverTab[58551]++
																return v.Uint() == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:516
		// _ = "end of CoverTab[58551]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:517
		_go_fuzz_dep_.CoverTab[58552]++
																return math.Float64bits(v.Float()) == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:518
		// _ = "end of CoverTab[58552]"
	case reflect.Complex64, reflect.Complex128:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:519
		_go_fuzz_dep_.CoverTab[58553]++
																c := v.Complex()
																return math.Float64bits(real(c)) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:521
			_go_fuzz_dep_.CoverTab[58561]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:521
			return math.Float64bits(imag(c)) == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:521
			// _ = "end of CoverTab[58561]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:521
		}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:521
		// _ = "end of CoverTab[58553]"
	case reflect.Array:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:522
		_go_fuzz_dep_.CoverTab[58554]++
																for i := 0; i < v.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:523
			_go_fuzz_dep_.CoverTab[58562]++
																	if !isZero(v.Index(i)) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:524
				_go_fuzz_dep_.CoverTab[58563]++
																		return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:525
				// _ = "end of CoverTab[58563]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:526
				_go_fuzz_dep_.CoverTab[58564]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:526
				// _ = "end of CoverTab[58564]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:526
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:526
			// _ = "end of CoverTab[58562]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:527
		// _ = "end of CoverTab[58554]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:527
		_go_fuzz_dep_.CoverTab[58555]++
																return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:528
		// _ = "end of CoverTab[58555]"
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:529
		_go_fuzz_dep_.CoverTab[58556]++
																return v.IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:530
		// _ = "end of CoverTab[58556]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:531
		_go_fuzz_dep_.CoverTab[58557]++
																return v.Len() == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:532
		// _ = "end of CoverTab[58557]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:533
		_go_fuzz_dep_.CoverTab[58558]++
																for i := 0; i < v.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:534
			_go_fuzz_dep_.CoverTab[58565]++
																	if !isZero(v.Field(i)) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:535
				_go_fuzz_dep_.CoverTab[58566]++
																		return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:536
				// _ = "end of CoverTab[58566]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:537
				_go_fuzz_dep_.CoverTab[58567]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:537
				// _ = "end of CoverTab[58567]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:537
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:537
			// _ = "end of CoverTab[58565]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:538
		// _ = "end of CoverTab[58558]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:538
		_go_fuzz_dep_.CoverTab[58559]++
																return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:539
		// _ = "end of CoverTab[58559]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:540
		_go_fuzz_dep_.CoverTab[58560]++
																panic(&reflect.ValueError{"reflect.Value.IsZero", v.Kind()})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:541
		// _ = "end of CoverTab[58560]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:542
	// _ = "end of CoverTab[58548]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:543
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect_field.go:543
var _ = _go_fuzz_dep_.CoverTab
