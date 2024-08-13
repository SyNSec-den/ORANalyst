// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:5
)

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// pointerCoderFuncs is a set of pointer encoding functions.
type pointerCoderFuncs struct {
	mi		*MessageInfo
	size		func(p pointer, f *coderFieldInfo, opts marshalOptions) int
	marshal		func(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error)
	unmarshal	func(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error)
	isInit		func(p pointer, f *coderFieldInfo) error
	merge		func(dst, src pointer, f *coderFieldInfo, opts mergeOptions)
}

// valueCoderFuncs is a set of protoreflect.Value encoding functions.
type valueCoderFuncs struct {
	size		func(v protoreflect.Value, tagsize int, opts marshalOptions) int
	marshal		func(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error)
	unmarshal	func(b []byte, v protoreflect.Value, num protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (protoreflect.Value, unmarshalOutput, error)
	isInit		func(v protoreflect.Value) error
	merge		func(dst, src protoreflect.Value, opts mergeOptions) protoreflect.Value
}

// fieldCoder returns pointer functions for a field, used for operating on
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:35
// struct fields.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:37
func fieldCoder(fd protoreflect.FieldDescriptor, ft reflect.Type) (*MessageInfo, pointerCoderFuncs) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:37
	_go_fuzz_dep_.CoverTab[56504]++
													switch {
	case fd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:39
		_go_fuzz_dep_.CoverTab[56506]++
														return encoderFuncsForMap(fd, ft)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:40
		// _ = "end of CoverTab[56506]"
	case fd.Cardinality() == protoreflect.Repeated && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:41
		_go_fuzz_dep_.CoverTab[56516]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:41
		return !fd.IsPacked()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:41
		// _ = "end of CoverTab[56516]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:41
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:41
		_go_fuzz_dep_.CoverTab[56507]++

														if ft.Kind() != reflect.Slice {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:43
			_go_fuzz_dep_.CoverTab[56517]++
															break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:44
			// _ = "end of CoverTab[56517]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:45
			_go_fuzz_dep_.CoverTab[56518]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:45
			// _ = "end of CoverTab[56518]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:45
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:45
		// _ = "end of CoverTab[56507]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:45
		_go_fuzz_dep_.CoverTab[56508]++
														ft := ft.Elem()
														switch fd.Kind() {
		case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:48
			_go_fuzz_dep_.CoverTab[56519]++
															if ft.Kind() == reflect.Bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:49
				_go_fuzz_dep_.CoverTab[56542]++
																return nil, coderBoolSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:50
				// _ = "end of CoverTab[56542]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:51
				_go_fuzz_dep_.CoverTab[56543]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:51
				// _ = "end of CoverTab[56543]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:51
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:51
			// _ = "end of CoverTab[56519]"
		case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:52
			_go_fuzz_dep_.CoverTab[56520]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:53
				_go_fuzz_dep_.CoverTab[56544]++
																return nil, coderEnumSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:54
				// _ = "end of CoverTab[56544]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:55
				_go_fuzz_dep_.CoverTab[56545]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:55
				// _ = "end of CoverTab[56545]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:55
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:55
			// _ = "end of CoverTab[56520]"
		case protoreflect.Int32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:56
			_go_fuzz_dep_.CoverTab[56521]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:57
				_go_fuzz_dep_.CoverTab[56546]++
																return nil, coderInt32Slice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:58
				// _ = "end of CoverTab[56546]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:59
				_go_fuzz_dep_.CoverTab[56547]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:59
				// _ = "end of CoverTab[56547]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:59
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:59
			// _ = "end of CoverTab[56521]"
		case protoreflect.Sint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:60
			_go_fuzz_dep_.CoverTab[56522]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:61
				_go_fuzz_dep_.CoverTab[56548]++
																return nil, coderSint32Slice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:62
				// _ = "end of CoverTab[56548]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:63
				_go_fuzz_dep_.CoverTab[56549]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:63
				// _ = "end of CoverTab[56549]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:63
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:63
			// _ = "end of CoverTab[56522]"
		case protoreflect.Uint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:64
			_go_fuzz_dep_.CoverTab[56523]++
															if ft.Kind() == reflect.Uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:65
				_go_fuzz_dep_.CoverTab[56550]++
																return nil, coderUint32Slice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:66
				// _ = "end of CoverTab[56550]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:67
				_go_fuzz_dep_.CoverTab[56551]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:67
				// _ = "end of CoverTab[56551]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:67
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:67
			// _ = "end of CoverTab[56523]"
		case protoreflect.Int64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:68
			_go_fuzz_dep_.CoverTab[56524]++
															if ft.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:69
				_go_fuzz_dep_.CoverTab[56552]++
																return nil, coderInt64Slice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:70
				// _ = "end of CoverTab[56552]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:71
				_go_fuzz_dep_.CoverTab[56553]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:71
				// _ = "end of CoverTab[56553]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:71
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:71
			// _ = "end of CoverTab[56524]"
		case protoreflect.Sint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:72
			_go_fuzz_dep_.CoverTab[56525]++
															if ft.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:73
				_go_fuzz_dep_.CoverTab[56554]++
																return nil, coderSint64Slice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:74
				// _ = "end of CoverTab[56554]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:75
				_go_fuzz_dep_.CoverTab[56555]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:75
				// _ = "end of CoverTab[56555]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:75
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:75
			// _ = "end of CoverTab[56525]"
		case protoreflect.Uint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:76
			_go_fuzz_dep_.CoverTab[56526]++
															if ft.Kind() == reflect.Uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:77
				_go_fuzz_dep_.CoverTab[56556]++
																return nil, coderUint64Slice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:78
				// _ = "end of CoverTab[56556]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:79
				_go_fuzz_dep_.CoverTab[56557]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:79
				// _ = "end of CoverTab[56557]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:79
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:79
			// _ = "end of CoverTab[56526]"
		case protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:80
			_go_fuzz_dep_.CoverTab[56527]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:81
				_go_fuzz_dep_.CoverTab[56558]++
																return nil, coderSfixed32Slice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:82
				// _ = "end of CoverTab[56558]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:83
				_go_fuzz_dep_.CoverTab[56559]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:83
				// _ = "end of CoverTab[56559]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:83
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:83
			// _ = "end of CoverTab[56527]"
		case protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:84
			_go_fuzz_dep_.CoverTab[56528]++
															if ft.Kind() == reflect.Uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:85
				_go_fuzz_dep_.CoverTab[56560]++
																return nil, coderFixed32Slice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:86
				// _ = "end of CoverTab[56560]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:87
				_go_fuzz_dep_.CoverTab[56561]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:87
				// _ = "end of CoverTab[56561]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:87
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:87
			// _ = "end of CoverTab[56528]"
		case protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:88
			_go_fuzz_dep_.CoverTab[56529]++
															if ft.Kind() == reflect.Float32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:89
				_go_fuzz_dep_.CoverTab[56562]++
																return nil, coderFloatSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:90
				// _ = "end of CoverTab[56562]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:91
				_go_fuzz_dep_.CoverTab[56563]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:91
				// _ = "end of CoverTab[56563]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:91
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:91
			// _ = "end of CoverTab[56529]"
		case protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:92
			_go_fuzz_dep_.CoverTab[56530]++
															if ft.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:93
				_go_fuzz_dep_.CoverTab[56564]++
																return nil, coderSfixed64Slice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:94
				// _ = "end of CoverTab[56564]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:95
				_go_fuzz_dep_.CoverTab[56565]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:95
				// _ = "end of CoverTab[56565]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:95
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:95
			// _ = "end of CoverTab[56530]"
		case protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:96
			_go_fuzz_dep_.CoverTab[56531]++
															if ft.Kind() == reflect.Uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:97
				_go_fuzz_dep_.CoverTab[56566]++
																return nil, coderFixed64Slice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:98
				// _ = "end of CoverTab[56566]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:99
				_go_fuzz_dep_.CoverTab[56567]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:99
				// _ = "end of CoverTab[56567]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:99
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:99
			// _ = "end of CoverTab[56531]"
		case protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:100
			_go_fuzz_dep_.CoverTab[56532]++
															if ft.Kind() == reflect.Float64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:101
				_go_fuzz_dep_.CoverTab[56568]++
																return nil, coderDoubleSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:102
				// _ = "end of CoverTab[56568]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:103
				_go_fuzz_dep_.CoverTab[56569]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:103
				// _ = "end of CoverTab[56569]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:103
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:103
			// _ = "end of CoverTab[56532]"
		case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:104
			_go_fuzz_dep_.CoverTab[56533]++
															if ft.Kind() == reflect.String && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:105
				_go_fuzz_dep_.CoverTab[56570]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:105
				return strs.EnforceUTF8(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:105
				// _ = "end of CoverTab[56570]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:105
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:105
				_go_fuzz_dep_.CoverTab[56571]++
																return nil, coderStringSliceValidateUTF8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:106
				// _ = "end of CoverTab[56571]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:107
				_go_fuzz_dep_.CoverTab[56572]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:107
				// _ = "end of CoverTab[56572]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:107
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:107
			// _ = "end of CoverTab[56533]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:107
			_go_fuzz_dep_.CoverTab[56534]++
															if ft.Kind() == reflect.String {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:108
				_go_fuzz_dep_.CoverTab[56573]++
																return nil, coderStringSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:109
				// _ = "end of CoverTab[56573]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:110
				_go_fuzz_dep_.CoverTab[56574]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:110
				// _ = "end of CoverTab[56574]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:110
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:110
			// _ = "end of CoverTab[56534]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:110
			_go_fuzz_dep_.CoverTab[56535]++
															if ft.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:111
				_go_fuzz_dep_.CoverTab[56575]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:111
				return ft.Elem().Kind() == reflect.Uint8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:111
				// _ = "end of CoverTab[56575]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:111
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:111
				_go_fuzz_dep_.CoverTab[56576]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:111
				return strs.EnforceUTF8(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:111
				// _ = "end of CoverTab[56576]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:111
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:111
				_go_fuzz_dep_.CoverTab[56577]++
																return nil, coderBytesSliceValidateUTF8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:112
				// _ = "end of CoverTab[56577]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:113
				_go_fuzz_dep_.CoverTab[56578]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:113
				// _ = "end of CoverTab[56578]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:113
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:113
			// _ = "end of CoverTab[56535]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:113
			_go_fuzz_dep_.CoverTab[56536]++
															if ft.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:114
				_go_fuzz_dep_.CoverTab[56579]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:114
				return ft.Elem().Kind() == reflect.Uint8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:114
				// _ = "end of CoverTab[56579]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:114
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:114
				_go_fuzz_dep_.CoverTab[56580]++
																return nil, coderBytesSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:115
				// _ = "end of CoverTab[56580]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:116
				_go_fuzz_dep_.CoverTab[56581]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:116
				// _ = "end of CoverTab[56581]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:116
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:116
			// _ = "end of CoverTab[56536]"
		case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:117
			_go_fuzz_dep_.CoverTab[56537]++
															if ft.Kind() == reflect.String {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:118
				_go_fuzz_dep_.CoverTab[56582]++
																return nil, coderStringSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:119
				// _ = "end of CoverTab[56582]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:120
				_go_fuzz_dep_.CoverTab[56583]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:120
				// _ = "end of CoverTab[56583]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:120
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:120
			// _ = "end of CoverTab[56537]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:120
			_go_fuzz_dep_.CoverTab[56538]++
															if ft.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:121
				_go_fuzz_dep_.CoverTab[56584]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:121
				return ft.Elem().Kind() == reflect.Uint8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:121
				// _ = "end of CoverTab[56584]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:121
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:121
				_go_fuzz_dep_.CoverTab[56585]++
																return nil, coderBytesSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:122
				// _ = "end of CoverTab[56585]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:123
				_go_fuzz_dep_.CoverTab[56586]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:123
				// _ = "end of CoverTab[56586]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:123
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:123
			// _ = "end of CoverTab[56538]"
		case protoreflect.MessageKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:124
			_go_fuzz_dep_.CoverTab[56539]++
															return getMessageInfo(ft), makeMessageSliceFieldCoder(fd, ft)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:125
			// _ = "end of CoverTab[56539]"
		case protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:126
			_go_fuzz_dep_.CoverTab[56540]++
															return getMessageInfo(ft), makeGroupSliceFieldCoder(fd, ft)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:127
			// _ = "end of CoverTab[56540]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:127
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:127
			_go_fuzz_dep_.CoverTab[56541]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:127
			// _ = "end of CoverTab[56541]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:128
		// _ = "end of CoverTab[56508]"
	case fd.Cardinality() == protoreflect.Repeated && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:129
		_go_fuzz_dep_.CoverTab[56587]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:129
		return fd.IsPacked()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:129
		// _ = "end of CoverTab[56587]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:129
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:129
		_go_fuzz_dep_.CoverTab[56509]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:134
		if ft.Kind() != reflect.Slice {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:134
			_go_fuzz_dep_.CoverTab[56588]++
															break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:135
			// _ = "end of CoverTab[56588]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:136
			_go_fuzz_dep_.CoverTab[56589]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:136
			// _ = "end of CoverTab[56589]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:136
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:136
		// _ = "end of CoverTab[56509]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:136
		_go_fuzz_dep_.CoverTab[56510]++
														ft := ft.Elem()
														switch fd.Kind() {
		case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:139
			_go_fuzz_dep_.CoverTab[56590]++
															if ft.Kind() == reflect.Bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:140
				_go_fuzz_dep_.CoverTab[56605]++
																return nil, coderBoolPackedSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:141
				// _ = "end of CoverTab[56605]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:142
				_go_fuzz_dep_.CoverTab[56606]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:142
				// _ = "end of CoverTab[56606]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:142
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:142
			// _ = "end of CoverTab[56590]"
		case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:143
			_go_fuzz_dep_.CoverTab[56591]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:144
				_go_fuzz_dep_.CoverTab[56607]++
																return nil, coderEnumPackedSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:145
				// _ = "end of CoverTab[56607]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:146
				_go_fuzz_dep_.CoverTab[56608]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:146
				// _ = "end of CoverTab[56608]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:146
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:146
			// _ = "end of CoverTab[56591]"
		case protoreflect.Int32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:147
			_go_fuzz_dep_.CoverTab[56592]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:148
				_go_fuzz_dep_.CoverTab[56609]++
																return nil, coderInt32PackedSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:149
				// _ = "end of CoverTab[56609]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:150
				_go_fuzz_dep_.CoverTab[56610]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:150
				// _ = "end of CoverTab[56610]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:150
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:150
			// _ = "end of CoverTab[56592]"
		case protoreflect.Sint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:151
			_go_fuzz_dep_.CoverTab[56593]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:152
				_go_fuzz_dep_.CoverTab[56611]++
																return nil, coderSint32PackedSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:153
				// _ = "end of CoverTab[56611]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:154
				_go_fuzz_dep_.CoverTab[56612]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:154
				// _ = "end of CoverTab[56612]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:154
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:154
			// _ = "end of CoverTab[56593]"
		case protoreflect.Uint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:155
			_go_fuzz_dep_.CoverTab[56594]++
															if ft.Kind() == reflect.Uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:156
				_go_fuzz_dep_.CoverTab[56613]++
																return nil, coderUint32PackedSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:157
				// _ = "end of CoverTab[56613]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:158
				_go_fuzz_dep_.CoverTab[56614]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:158
				// _ = "end of CoverTab[56614]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:158
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:158
			// _ = "end of CoverTab[56594]"
		case protoreflect.Int64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:159
			_go_fuzz_dep_.CoverTab[56595]++
															if ft.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:160
				_go_fuzz_dep_.CoverTab[56615]++
																return nil, coderInt64PackedSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:161
				// _ = "end of CoverTab[56615]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:162
				_go_fuzz_dep_.CoverTab[56616]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:162
				// _ = "end of CoverTab[56616]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:162
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:162
			// _ = "end of CoverTab[56595]"
		case protoreflect.Sint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:163
			_go_fuzz_dep_.CoverTab[56596]++
															if ft.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:164
				_go_fuzz_dep_.CoverTab[56617]++
																return nil, coderSint64PackedSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:165
				// _ = "end of CoverTab[56617]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:166
				_go_fuzz_dep_.CoverTab[56618]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:166
				// _ = "end of CoverTab[56618]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:166
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:166
			// _ = "end of CoverTab[56596]"
		case protoreflect.Uint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:167
			_go_fuzz_dep_.CoverTab[56597]++
															if ft.Kind() == reflect.Uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:168
				_go_fuzz_dep_.CoverTab[56619]++
																return nil, coderUint64PackedSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:169
				// _ = "end of CoverTab[56619]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:170
				_go_fuzz_dep_.CoverTab[56620]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:170
				// _ = "end of CoverTab[56620]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:170
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:170
			// _ = "end of CoverTab[56597]"
		case protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:171
			_go_fuzz_dep_.CoverTab[56598]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:172
				_go_fuzz_dep_.CoverTab[56621]++
																return nil, coderSfixed32PackedSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:173
				// _ = "end of CoverTab[56621]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:174
				_go_fuzz_dep_.CoverTab[56622]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:174
				// _ = "end of CoverTab[56622]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:174
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:174
			// _ = "end of CoverTab[56598]"
		case protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:175
			_go_fuzz_dep_.CoverTab[56599]++
															if ft.Kind() == reflect.Uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:176
				_go_fuzz_dep_.CoverTab[56623]++
																return nil, coderFixed32PackedSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:177
				// _ = "end of CoverTab[56623]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:178
				_go_fuzz_dep_.CoverTab[56624]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:178
				// _ = "end of CoverTab[56624]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:178
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:178
			// _ = "end of CoverTab[56599]"
		case protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:179
			_go_fuzz_dep_.CoverTab[56600]++
															if ft.Kind() == reflect.Float32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:180
				_go_fuzz_dep_.CoverTab[56625]++
																return nil, coderFloatPackedSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:181
				// _ = "end of CoverTab[56625]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:182
				_go_fuzz_dep_.CoverTab[56626]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:182
				// _ = "end of CoverTab[56626]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:182
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:182
			// _ = "end of CoverTab[56600]"
		case protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:183
			_go_fuzz_dep_.CoverTab[56601]++
															if ft.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:184
				_go_fuzz_dep_.CoverTab[56627]++
																return nil, coderSfixed64PackedSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:185
				// _ = "end of CoverTab[56627]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:186
				_go_fuzz_dep_.CoverTab[56628]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:186
				// _ = "end of CoverTab[56628]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:186
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:186
			// _ = "end of CoverTab[56601]"
		case protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:187
			_go_fuzz_dep_.CoverTab[56602]++
															if ft.Kind() == reflect.Uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:188
				_go_fuzz_dep_.CoverTab[56629]++
																return nil, coderFixed64PackedSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:189
				// _ = "end of CoverTab[56629]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:190
				_go_fuzz_dep_.CoverTab[56630]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:190
				// _ = "end of CoverTab[56630]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:190
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:190
			// _ = "end of CoverTab[56602]"
		case protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:191
			_go_fuzz_dep_.CoverTab[56603]++
															if ft.Kind() == reflect.Float64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:192
				_go_fuzz_dep_.CoverTab[56631]++
																return nil, coderDoublePackedSlice
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:193
				// _ = "end of CoverTab[56631]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:194
				_go_fuzz_dep_.CoverTab[56632]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:194
				// _ = "end of CoverTab[56632]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:194
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:194
			// _ = "end of CoverTab[56603]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:194
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:194
			_go_fuzz_dep_.CoverTab[56604]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:194
			// _ = "end of CoverTab[56604]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:195
		// _ = "end of CoverTab[56510]"
	case fd.Kind() == protoreflect.MessageKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:196
		_go_fuzz_dep_.CoverTab[56511]++
														return getMessageInfo(ft), makeMessageFieldCoder(fd, ft)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:197
		// _ = "end of CoverTab[56511]"
	case fd.Kind() == protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:198
		_go_fuzz_dep_.CoverTab[56512]++
														return getMessageInfo(ft), makeGroupFieldCoder(fd, ft)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:199
		// _ = "end of CoverTab[56512]"
	case fd.Syntax() == protoreflect.Proto3 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:200
		_go_fuzz_dep_.CoverTab[56633]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:200
		return fd.ContainingOneof() == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:200
		// _ = "end of CoverTab[56633]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:200
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:200
		_go_fuzz_dep_.CoverTab[56513]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:203
		switch fd.Kind() {
		case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:204
			_go_fuzz_dep_.CoverTab[56634]++
															if ft.Kind() == reflect.Bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:205
				_go_fuzz_dep_.CoverTab[56655]++
																return nil, coderBoolNoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:206
				// _ = "end of CoverTab[56655]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:207
				_go_fuzz_dep_.CoverTab[56656]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:207
				// _ = "end of CoverTab[56656]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:207
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:207
			// _ = "end of CoverTab[56634]"
		case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:208
			_go_fuzz_dep_.CoverTab[56635]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:209
				_go_fuzz_dep_.CoverTab[56657]++
																return nil, coderEnumNoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:210
				// _ = "end of CoverTab[56657]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:211
				_go_fuzz_dep_.CoverTab[56658]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:211
				// _ = "end of CoverTab[56658]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:211
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:211
			// _ = "end of CoverTab[56635]"
		case protoreflect.Int32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:212
			_go_fuzz_dep_.CoverTab[56636]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:213
				_go_fuzz_dep_.CoverTab[56659]++
																return nil, coderInt32NoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:214
				// _ = "end of CoverTab[56659]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:215
				_go_fuzz_dep_.CoverTab[56660]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:215
				// _ = "end of CoverTab[56660]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:215
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:215
			// _ = "end of CoverTab[56636]"
		case protoreflect.Sint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:216
			_go_fuzz_dep_.CoverTab[56637]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:217
				_go_fuzz_dep_.CoverTab[56661]++
																return nil, coderSint32NoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:218
				// _ = "end of CoverTab[56661]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:219
				_go_fuzz_dep_.CoverTab[56662]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:219
				// _ = "end of CoverTab[56662]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:219
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:219
			// _ = "end of CoverTab[56637]"
		case protoreflect.Uint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:220
			_go_fuzz_dep_.CoverTab[56638]++
															if ft.Kind() == reflect.Uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:221
				_go_fuzz_dep_.CoverTab[56663]++
																return nil, coderUint32NoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:222
				// _ = "end of CoverTab[56663]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:223
				_go_fuzz_dep_.CoverTab[56664]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:223
				// _ = "end of CoverTab[56664]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:223
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:223
			// _ = "end of CoverTab[56638]"
		case protoreflect.Int64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:224
			_go_fuzz_dep_.CoverTab[56639]++
															if ft.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:225
				_go_fuzz_dep_.CoverTab[56665]++
																return nil, coderInt64NoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:226
				// _ = "end of CoverTab[56665]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:227
				_go_fuzz_dep_.CoverTab[56666]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:227
				// _ = "end of CoverTab[56666]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:227
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:227
			// _ = "end of CoverTab[56639]"
		case protoreflect.Sint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:228
			_go_fuzz_dep_.CoverTab[56640]++
															if ft.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:229
				_go_fuzz_dep_.CoverTab[56667]++
																return nil, coderSint64NoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:230
				// _ = "end of CoverTab[56667]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:231
				_go_fuzz_dep_.CoverTab[56668]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:231
				// _ = "end of CoverTab[56668]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:231
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:231
			// _ = "end of CoverTab[56640]"
		case protoreflect.Uint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:232
			_go_fuzz_dep_.CoverTab[56641]++
															if ft.Kind() == reflect.Uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:233
				_go_fuzz_dep_.CoverTab[56669]++
																return nil, coderUint64NoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:234
				// _ = "end of CoverTab[56669]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:235
				_go_fuzz_dep_.CoverTab[56670]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:235
				// _ = "end of CoverTab[56670]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:235
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:235
			// _ = "end of CoverTab[56641]"
		case protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:236
			_go_fuzz_dep_.CoverTab[56642]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:237
				_go_fuzz_dep_.CoverTab[56671]++
																return nil, coderSfixed32NoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:238
				// _ = "end of CoverTab[56671]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:239
				_go_fuzz_dep_.CoverTab[56672]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:239
				// _ = "end of CoverTab[56672]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:239
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:239
			// _ = "end of CoverTab[56642]"
		case protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:240
			_go_fuzz_dep_.CoverTab[56643]++
															if ft.Kind() == reflect.Uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:241
				_go_fuzz_dep_.CoverTab[56673]++
																return nil, coderFixed32NoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:242
				// _ = "end of CoverTab[56673]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:243
				_go_fuzz_dep_.CoverTab[56674]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:243
				// _ = "end of CoverTab[56674]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:243
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:243
			// _ = "end of CoverTab[56643]"
		case protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:244
			_go_fuzz_dep_.CoverTab[56644]++
															if ft.Kind() == reflect.Float32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:245
				_go_fuzz_dep_.CoverTab[56675]++
																return nil, coderFloatNoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:246
				// _ = "end of CoverTab[56675]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:247
				_go_fuzz_dep_.CoverTab[56676]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:247
				// _ = "end of CoverTab[56676]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:247
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:247
			// _ = "end of CoverTab[56644]"
		case protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:248
			_go_fuzz_dep_.CoverTab[56645]++
															if ft.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:249
				_go_fuzz_dep_.CoverTab[56677]++
																return nil, coderSfixed64NoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:250
				// _ = "end of CoverTab[56677]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:251
				_go_fuzz_dep_.CoverTab[56678]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:251
				// _ = "end of CoverTab[56678]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:251
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:251
			// _ = "end of CoverTab[56645]"
		case protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:252
			_go_fuzz_dep_.CoverTab[56646]++
															if ft.Kind() == reflect.Uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:253
				_go_fuzz_dep_.CoverTab[56679]++
																return nil, coderFixed64NoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:254
				// _ = "end of CoverTab[56679]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:255
				_go_fuzz_dep_.CoverTab[56680]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:255
				// _ = "end of CoverTab[56680]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:255
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:255
			// _ = "end of CoverTab[56646]"
		case protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:256
			_go_fuzz_dep_.CoverTab[56647]++
															if ft.Kind() == reflect.Float64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:257
				_go_fuzz_dep_.CoverTab[56681]++
																return nil, coderDoubleNoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:258
				// _ = "end of CoverTab[56681]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:259
				_go_fuzz_dep_.CoverTab[56682]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:259
				// _ = "end of CoverTab[56682]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:259
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:259
			// _ = "end of CoverTab[56647]"
		case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:260
			_go_fuzz_dep_.CoverTab[56648]++
															if ft.Kind() == reflect.String && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:261
				_go_fuzz_dep_.CoverTab[56683]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:261
				return strs.EnforceUTF8(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:261
				// _ = "end of CoverTab[56683]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:261
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:261
				_go_fuzz_dep_.CoverTab[56684]++
																return nil, coderStringNoZeroValidateUTF8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:262
				// _ = "end of CoverTab[56684]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:263
				_go_fuzz_dep_.CoverTab[56685]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:263
				// _ = "end of CoverTab[56685]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:263
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:263
			// _ = "end of CoverTab[56648]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:263
			_go_fuzz_dep_.CoverTab[56649]++
															if ft.Kind() == reflect.String {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:264
				_go_fuzz_dep_.CoverTab[56686]++
																return nil, coderStringNoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:265
				// _ = "end of CoverTab[56686]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:266
				_go_fuzz_dep_.CoverTab[56687]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:266
				// _ = "end of CoverTab[56687]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:266
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:266
			// _ = "end of CoverTab[56649]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:266
			_go_fuzz_dep_.CoverTab[56650]++
															if ft.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:267
				_go_fuzz_dep_.CoverTab[56688]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:267
				return ft.Elem().Kind() == reflect.Uint8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:267
				// _ = "end of CoverTab[56688]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:267
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:267
				_go_fuzz_dep_.CoverTab[56689]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:267
				return strs.EnforceUTF8(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:267
				// _ = "end of CoverTab[56689]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:267
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:267
				_go_fuzz_dep_.CoverTab[56690]++
																return nil, coderBytesNoZeroValidateUTF8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:268
				// _ = "end of CoverTab[56690]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:269
				_go_fuzz_dep_.CoverTab[56691]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:269
				// _ = "end of CoverTab[56691]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:269
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:269
			// _ = "end of CoverTab[56650]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:269
			_go_fuzz_dep_.CoverTab[56651]++
															if ft.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:270
				_go_fuzz_dep_.CoverTab[56692]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:270
				return ft.Elem().Kind() == reflect.Uint8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:270
				// _ = "end of CoverTab[56692]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:270
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:270
				_go_fuzz_dep_.CoverTab[56693]++
																return nil, coderBytesNoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:271
				// _ = "end of CoverTab[56693]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:272
				_go_fuzz_dep_.CoverTab[56694]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:272
				// _ = "end of CoverTab[56694]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:272
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:272
			// _ = "end of CoverTab[56651]"
		case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:273
			_go_fuzz_dep_.CoverTab[56652]++
															if ft.Kind() == reflect.String {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:274
				_go_fuzz_dep_.CoverTab[56695]++
																return nil, coderStringNoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:275
				// _ = "end of CoverTab[56695]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:276
				_go_fuzz_dep_.CoverTab[56696]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:276
				// _ = "end of CoverTab[56696]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:276
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:276
			// _ = "end of CoverTab[56652]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:276
			_go_fuzz_dep_.CoverTab[56653]++
															if ft.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:277
				_go_fuzz_dep_.CoverTab[56697]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:277
				return ft.Elem().Kind() == reflect.Uint8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:277
				// _ = "end of CoverTab[56697]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:277
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:277
				_go_fuzz_dep_.CoverTab[56698]++
																return nil, coderBytesNoZero
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:278
				// _ = "end of CoverTab[56698]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:279
				_go_fuzz_dep_.CoverTab[56699]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:279
				// _ = "end of CoverTab[56699]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:279
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:279
			// _ = "end of CoverTab[56653]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:279
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:279
			_go_fuzz_dep_.CoverTab[56654]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:279
			// _ = "end of CoverTab[56654]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:280
		// _ = "end of CoverTab[56513]"
	case ft.Kind() == reflect.Ptr:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:281
		_go_fuzz_dep_.CoverTab[56514]++
														ft := ft.Elem()
														switch fd.Kind() {
		case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:284
			_go_fuzz_dep_.CoverTab[56700]++
															if ft.Kind() == reflect.Bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:285
				_go_fuzz_dep_.CoverTab[56718]++
																return nil, coderBoolPtr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:286
				// _ = "end of CoverTab[56718]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:287
				_go_fuzz_dep_.CoverTab[56719]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:287
				// _ = "end of CoverTab[56719]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:287
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:287
			// _ = "end of CoverTab[56700]"
		case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:288
			_go_fuzz_dep_.CoverTab[56701]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:289
				_go_fuzz_dep_.CoverTab[56720]++
																return nil, coderEnumPtr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:290
				// _ = "end of CoverTab[56720]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:291
				_go_fuzz_dep_.CoverTab[56721]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:291
				// _ = "end of CoverTab[56721]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:291
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:291
			// _ = "end of CoverTab[56701]"
		case protoreflect.Int32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:292
			_go_fuzz_dep_.CoverTab[56702]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:293
				_go_fuzz_dep_.CoverTab[56722]++
																return nil, coderInt32Ptr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:294
				// _ = "end of CoverTab[56722]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:295
				_go_fuzz_dep_.CoverTab[56723]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:295
				// _ = "end of CoverTab[56723]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:295
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:295
			// _ = "end of CoverTab[56702]"
		case protoreflect.Sint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:296
			_go_fuzz_dep_.CoverTab[56703]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:297
				_go_fuzz_dep_.CoverTab[56724]++
																return nil, coderSint32Ptr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:298
				// _ = "end of CoverTab[56724]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:299
				_go_fuzz_dep_.CoverTab[56725]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:299
				// _ = "end of CoverTab[56725]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:299
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:299
			// _ = "end of CoverTab[56703]"
		case protoreflect.Uint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:300
			_go_fuzz_dep_.CoverTab[56704]++
															if ft.Kind() == reflect.Uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:301
				_go_fuzz_dep_.CoverTab[56726]++
																return nil, coderUint32Ptr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:302
				// _ = "end of CoverTab[56726]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:303
				_go_fuzz_dep_.CoverTab[56727]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:303
				// _ = "end of CoverTab[56727]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:303
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:303
			// _ = "end of CoverTab[56704]"
		case protoreflect.Int64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:304
			_go_fuzz_dep_.CoverTab[56705]++
															if ft.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:305
				_go_fuzz_dep_.CoverTab[56728]++
																return nil, coderInt64Ptr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:306
				// _ = "end of CoverTab[56728]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:307
				_go_fuzz_dep_.CoverTab[56729]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:307
				// _ = "end of CoverTab[56729]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:307
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:307
			// _ = "end of CoverTab[56705]"
		case protoreflect.Sint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:308
			_go_fuzz_dep_.CoverTab[56706]++
															if ft.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:309
				_go_fuzz_dep_.CoverTab[56730]++
																return nil, coderSint64Ptr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:310
				// _ = "end of CoverTab[56730]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:311
				_go_fuzz_dep_.CoverTab[56731]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:311
				// _ = "end of CoverTab[56731]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:311
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:311
			// _ = "end of CoverTab[56706]"
		case protoreflect.Uint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:312
			_go_fuzz_dep_.CoverTab[56707]++
															if ft.Kind() == reflect.Uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:313
				_go_fuzz_dep_.CoverTab[56732]++
																return nil, coderUint64Ptr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:314
				// _ = "end of CoverTab[56732]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:315
				_go_fuzz_dep_.CoverTab[56733]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:315
				// _ = "end of CoverTab[56733]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:315
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:315
			// _ = "end of CoverTab[56707]"
		case protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:316
			_go_fuzz_dep_.CoverTab[56708]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:317
				_go_fuzz_dep_.CoverTab[56734]++
																return nil, coderSfixed32Ptr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:318
				// _ = "end of CoverTab[56734]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:319
				_go_fuzz_dep_.CoverTab[56735]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:319
				// _ = "end of CoverTab[56735]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:319
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:319
			// _ = "end of CoverTab[56708]"
		case protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:320
			_go_fuzz_dep_.CoverTab[56709]++
															if ft.Kind() == reflect.Uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:321
				_go_fuzz_dep_.CoverTab[56736]++
																return nil, coderFixed32Ptr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:322
				// _ = "end of CoverTab[56736]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:323
				_go_fuzz_dep_.CoverTab[56737]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:323
				// _ = "end of CoverTab[56737]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:323
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:323
			// _ = "end of CoverTab[56709]"
		case protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:324
			_go_fuzz_dep_.CoverTab[56710]++
															if ft.Kind() == reflect.Float32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:325
				_go_fuzz_dep_.CoverTab[56738]++
																return nil, coderFloatPtr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:326
				// _ = "end of CoverTab[56738]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:327
				_go_fuzz_dep_.CoverTab[56739]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:327
				// _ = "end of CoverTab[56739]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:327
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:327
			// _ = "end of CoverTab[56710]"
		case protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:328
			_go_fuzz_dep_.CoverTab[56711]++
															if ft.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:329
				_go_fuzz_dep_.CoverTab[56740]++
																return nil, coderSfixed64Ptr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:330
				// _ = "end of CoverTab[56740]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:331
				_go_fuzz_dep_.CoverTab[56741]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:331
				// _ = "end of CoverTab[56741]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:331
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:331
			// _ = "end of CoverTab[56711]"
		case protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:332
			_go_fuzz_dep_.CoverTab[56712]++
															if ft.Kind() == reflect.Uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:333
				_go_fuzz_dep_.CoverTab[56742]++
																return nil, coderFixed64Ptr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:334
				// _ = "end of CoverTab[56742]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:335
				_go_fuzz_dep_.CoverTab[56743]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:335
				// _ = "end of CoverTab[56743]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:335
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:335
			// _ = "end of CoverTab[56712]"
		case protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:336
			_go_fuzz_dep_.CoverTab[56713]++
															if ft.Kind() == reflect.Float64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:337
				_go_fuzz_dep_.CoverTab[56744]++
																return nil, coderDoublePtr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:338
				// _ = "end of CoverTab[56744]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:339
				_go_fuzz_dep_.CoverTab[56745]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:339
				// _ = "end of CoverTab[56745]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:339
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:339
			// _ = "end of CoverTab[56713]"
		case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:340
			_go_fuzz_dep_.CoverTab[56714]++
															if ft.Kind() == reflect.String && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:341
				_go_fuzz_dep_.CoverTab[56746]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:341
				return strs.EnforceUTF8(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:341
				// _ = "end of CoverTab[56746]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:341
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:341
				_go_fuzz_dep_.CoverTab[56747]++
																return nil, coderStringPtrValidateUTF8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:342
				// _ = "end of CoverTab[56747]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:343
				_go_fuzz_dep_.CoverTab[56748]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:343
				// _ = "end of CoverTab[56748]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:343
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:343
			// _ = "end of CoverTab[56714]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:343
			_go_fuzz_dep_.CoverTab[56715]++
															if ft.Kind() == reflect.String {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:344
				_go_fuzz_dep_.CoverTab[56749]++
																return nil, coderStringPtr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:345
				// _ = "end of CoverTab[56749]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:346
				_go_fuzz_dep_.CoverTab[56750]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:346
				// _ = "end of CoverTab[56750]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:346
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:346
			// _ = "end of CoverTab[56715]"
		case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:347
			_go_fuzz_dep_.CoverTab[56716]++
															if ft.Kind() == reflect.String {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:348
				_go_fuzz_dep_.CoverTab[56751]++
																return nil, coderStringPtr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:349
				// _ = "end of CoverTab[56751]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:350
				_go_fuzz_dep_.CoverTab[56752]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:350
				// _ = "end of CoverTab[56752]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:350
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:350
			// _ = "end of CoverTab[56716]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:350
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:350
			_go_fuzz_dep_.CoverTab[56717]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:350
			// _ = "end of CoverTab[56717]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:351
		// _ = "end of CoverTab[56514]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:352
		_go_fuzz_dep_.CoverTab[56515]++
														switch fd.Kind() {
		case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:354
			_go_fuzz_dep_.CoverTab[56753]++
															if ft.Kind() == reflect.Bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:355
				_go_fuzz_dep_.CoverTab[56774]++
																return nil, coderBool
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:356
				// _ = "end of CoverTab[56774]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:357
				_go_fuzz_dep_.CoverTab[56775]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:357
				// _ = "end of CoverTab[56775]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:357
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:357
			// _ = "end of CoverTab[56753]"
		case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:358
			_go_fuzz_dep_.CoverTab[56754]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:359
				_go_fuzz_dep_.CoverTab[56776]++
																return nil, coderEnum
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:360
				// _ = "end of CoverTab[56776]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:361
				_go_fuzz_dep_.CoverTab[56777]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:361
				// _ = "end of CoverTab[56777]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:361
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:361
			// _ = "end of CoverTab[56754]"
		case protoreflect.Int32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:362
			_go_fuzz_dep_.CoverTab[56755]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:363
				_go_fuzz_dep_.CoverTab[56778]++
																return nil, coderInt32
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:364
				// _ = "end of CoverTab[56778]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:365
				_go_fuzz_dep_.CoverTab[56779]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:365
				// _ = "end of CoverTab[56779]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:365
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:365
			// _ = "end of CoverTab[56755]"
		case protoreflect.Sint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:366
			_go_fuzz_dep_.CoverTab[56756]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:367
				_go_fuzz_dep_.CoverTab[56780]++
																return nil, coderSint32
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:368
				// _ = "end of CoverTab[56780]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:369
				_go_fuzz_dep_.CoverTab[56781]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:369
				// _ = "end of CoverTab[56781]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:369
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:369
			// _ = "end of CoverTab[56756]"
		case protoreflect.Uint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:370
			_go_fuzz_dep_.CoverTab[56757]++
															if ft.Kind() == reflect.Uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:371
				_go_fuzz_dep_.CoverTab[56782]++
																return nil, coderUint32
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:372
				// _ = "end of CoverTab[56782]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:373
				_go_fuzz_dep_.CoverTab[56783]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:373
				// _ = "end of CoverTab[56783]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:373
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:373
			// _ = "end of CoverTab[56757]"
		case protoreflect.Int64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:374
			_go_fuzz_dep_.CoverTab[56758]++
															if ft.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:375
				_go_fuzz_dep_.CoverTab[56784]++
																return nil, coderInt64
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:376
				// _ = "end of CoverTab[56784]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:377
				_go_fuzz_dep_.CoverTab[56785]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:377
				// _ = "end of CoverTab[56785]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:377
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:377
			// _ = "end of CoverTab[56758]"
		case protoreflect.Sint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:378
			_go_fuzz_dep_.CoverTab[56759]++
															if ft.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:379
				_go_fuzz_dep_.CoverTab[56786]++
																return nil, coderSint64
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:380
				// _ = "end of CoverTab[56786]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:381
				_go_fuzz_dep_.CoverTab[56787]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:381
				// _ = "end of CoverTab[56787]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:381
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:381
			// _ = "end of CoverTab[56759]"
		case protoreflect.Uint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:382
			_go_fuzz_dep_.CoverTab[56760]++
															if ft.Kind() == reflect.Uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:383
				_go_fuzz_dep_.CoverTab[56788]++
																return nil, coderUint64
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:384
				// _ = "end of CoverTab[56788]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:385
				_go_fuzz_dep_.CoverTab[56789]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:385
				// _ = "end of CoverTab[56789]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:385
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:385
			// _ = "end of CoverTab[56760]"
		case protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:386
			_go_fuzz_dep_.CoverTab[56761]++
															if ft.Kind() == reflect.Int32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:387
				_go_fuzz_dep_.CoverTab[56790]++
																return nil, coderSfixed32
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:388
				// _ = "end of CoverTab[56790]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:389
				_go_fuzz_dep_.CoverTab[56791]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:389
				// _ = "end of CoverTab[56791]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:389
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:389
			// _ = "end of CoverTab[56761]"
		case protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:390
			_go_fuzz_dep_.CoverTab[56762]++
															if ft.Kind() == reflect.Uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:391
				_go_fuzz_dep_.CoverTab[56792]++
																return nil, coderFixed32
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:392
				// _ = "end of CoverTab[56792]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:393
				_go_fuzz_dep_.CoverTab[56793]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:393
				// _ = "end of CoverTab[56793]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:393
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:393
			// _ = "end of CoverTab[56762]"
		case protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:394
			_go_fuzz_dep_.CoverTab[56763]++
															if ft.Kind() == reflect.Float32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:395
				_go_fuzz_dep_.CoverTab[56794]++
																return nil, coderFloat
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:396
				// _ = "end of CoverTab[56794]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:397
				_go_fuzz_dep_.CoverTab[56795]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:397
				// _ = "end of CoverTab[56795]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:397
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:397
			// _ = "end of CoverTab[56763]"
		case protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:398
			_go_fuzz_dep_.CoverTab[56764]++
															if ft.Kind() == reflect.Int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:399
				_go_fuzz_dep_.CoverTab[56796]++
																return nil, coderSfixed64
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:400
				// _ = "end of CoverTab[56796]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:401
				_go_fuzz_dep_.CoverTab[56797]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:401
				// _ = "end of CoverTab[56797]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:401
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:401
			// _ = "end of CoverTab[56764]"
		case protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:402
			_go_fuzz_dep_.CoverTab[56765]++
															if ft.Kind() == reflect.Uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:403
				_go_fuzz_dep_.CoverTab[56798]++
																return nil, coderFixed64
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:404
				// _ = "end of CoverTab[56798]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:405
				_go_fuzz_dep_.CoverTab[56799]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:405
				// _ = "end of CoverTab[56799]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:405
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:405
			// _ = "end of CoverTab[56765]"
		case protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:406
			_go_fuzz_dep_.CoverTab[56766]++
															if ft.Kind() == reflect.Float64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:407
				_go_fuzz_dep_.CoverTab[56800]++
																return nil, coderDouble
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:408
				// _ = "end of CoverTab[56800]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:409
				_go_fuzz_dep_.CoverTab[56801]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:409
				// _ = "end of CoverTab[56801]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:409
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:409
			// _ = "end of CoverTab[56766]"
		case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:410
			_go_fuzz_dep_.CoverTab[56767]++
															if ft.Kind() == reflect.String && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:411
				_go_fuzz_dep_.CoverTab[56802]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:411
				return strs.EnforceUTF8(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:411
				// _ = "end of CoverTab[56802]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:411
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:411
				_go_fuzz_dep_.CoverTab[56803]++
																return nil, coderStringValidateUTF8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:412
				// _ = "end of CoverTab[56803]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:413
				_go_fuzz_dep_.CoverTab[56804]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:413
				// _ = "end of CoverTab[56804]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:413
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:413
			// _ = "end of CoverTab[56767]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:413
			_go_fuzz_dep_.CoverTab[56768]++
															if ft.Kind() == reflect.String {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:414
				_go_fuzz_dep_.CoverTab[56805]++
																return nil, coderString
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:415
				// _ = "end of CoverTab[56805]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:416
				_go_fuzz_dep_.CoverTab[56806]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:416
				// _ = "end of CoverTab[56806]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:416
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:416
			// _ = "end of CoverTab[56768]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:416
			_go_fuzz_dep_.CoverTab[56769]++
															if ft.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:417
				_go_fuzz_dep_.CoverTab[56807]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:417
				return ft.Elem().Kind() == reflect.Uint8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:417
				// _ = "end of CoverTab[56807]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:417
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:417
				_go_fuzz_dep_.CoverTab[56808]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:417
				return strs.EnforceUTF8(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:417
				// _ = "end of CoverTab[56808]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:417
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:417
				_go_fuzz_dep_.CoverTab[56809]++
																return nil, coderBytesValidateUTF8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:418
				// _ = "end of CoverTab[56809]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:419
				_go_fuzz_dep_.CoverTab[56810]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:419
				// _ = "end of CoverTab[56810]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:419
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:419
			// _ = "end of CoverTab[56769]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:419
			_go_fuzz_dep_.CoverTab[56770]++
															if ft.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:420
				_go_fuzz_dep_.CoverTab[56811]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:420
				return ft.Elem().Kind() == reflect.Uint8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:420
				// _ = "end of CoverTab[56811]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:420
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:420
				_go_fuzz_dep_.CoverTab[56812]++
																return nil, coderBytes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:421
				// _ = "end of CoverTab[56812]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:422
				_go_fuzz_dep_.CoverTab[56813]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:422
				// _ = "end of CoverTab[56813]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:422
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:422
			// _ = "end of CoverTab[56770]"
		case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:423
			_go_fuzz_dep_.CoverTab[56771]++
															if ft.Kind() == reflect.String {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:424
				_go_fuzz_dep_.CoverTab[56814]++
																return nil, coderString
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:425
				// _ = "end of CoverTab[56814]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:426
				_go_fuzz_dep_.CoverTab[56815]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:426
				// _ = "end of CoverTab[56815]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:426
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:426
			// _ = "end of CoverTab[56771]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:426
			_go_fuzz_dep_.CoverTab[56772]++
															if ft.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:427
				_go_fuzz_dep_.CoverTab[56816]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:427
				return ft.Elem().Kind() == reflect.Uint8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:427
				// _ = "end of CoverTab[56816]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:427
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:427
				_go_fuzz_dep_.CoverTab[56817]++
																return nil, coderBytes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:428
				// _ = "end of CoverTab[56817]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:429
				_go_fuzz_dep_.CoverTab[56818]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:429
				// _ = "end of CoverTab[56818]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:429
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:429
			// _ = "end of CoverTab[56772]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:429
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:429
			_go_fuzz_dep_.CoverTab[56773]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:429
			// _ = "end of CoverTab[56773]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:430
		// _ = "end of CoverTab[56515]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:431
	// _ = "end of CoverTab[56504]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:431
	_go_fuzz_dep_.CoverTab[56505]++
													panic(fmt.Sprintf("invalid type: no encoder for %v %v %v/%v", fd.FullName(), fd.Cardinality(), fd.Kind(), ft))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:432
	// _ = "end of CoverTab[56505]"
}

// encoderFuncsForValue returns value functions for a field, used for
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:435
// extension values and map encoding.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:437
func encoderFuncsForValue(fd protoreflect.FieldDescriptor) valueCoderFuncs {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:437
	_go_fuzz_dep_.CoverTab[56819]++
													switch {
	case fd.Cardinality() == protoreflect.Repeated && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:439
		_go_fuzz_dep_.CoverTab[56824]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:439
		return !fd.IsPacked()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:439
		// _ = "end of CoverTab[56824]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:439
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:439
		_go_fuzz_dep_.CoverTab[56821]++
														switch fd.Kind() {
		case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:441
			_go_fuzz_dep_.CoverTab[56825]++
															return coderBoolSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:442
			// _ = "end of CoverTab[56825]"
		case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:443
			_go_fuzz_dep_.CoverTab[56826]++
															return coderEnumSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:444
			// _ = "end of CoverTab[56826]"
		case protoreflect.Int32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:445
			_go_fuzz_dep_.CoverTab[56827]++
															return coderInt32SliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:446
			// _ = "end of CoverTab[56827]"
		case protoreflect.Sint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:447
			_go_fuzz_dep_.CoverTab[56828]++
															return coderSint32SliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:448
			// _ = "end of CoverTab[56828]"
		case protoreflect.Uint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:449
			_go_fuzz_dep_.CoverTab[56829]++
															return coderUint32SliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:450
			// _ = "end of CoverTab[56829]"
		case protoreflect.Int64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:451
			_go_fuzz_dep_.CoverTab[56830]++
															return coderInt64SliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:452
			// _ = "end of CoverTab[56830]"
		case protoreflect.Sint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:453
			_go_fuzz_dep_.CoverTab[56831]++
															return coderSint64SliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:454
			// _ = "end of CoverTab[56831]"
		case protoreflect.Uint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:455
			_go_fuzz_dep_.CoverTab[56832]++
															return coderUint64SliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:456
			// _ = "end of CoverTab[56832]"
		case protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:457
			_go_fuzz_dep_.CoverTab[56833]++
															return coderSfixed32SliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:458
			// _ = "end of CoverTab[56833]"
		case protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:459
			_go_fuzz_dep_.CoverTab[56834]++
															return coderFixed32SliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:460
			// _ = "end of CoverTab[56834]"
		case protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:461
			_go_fuzz_dep_.CoverTab[56835]++
															return coderFloatSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:462
			// _ = "end of CoverTab[56835]"
		case protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:463
			_go_fuzz_dep_.CoverTab[56836]++
															return coderSfixed64SliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:464
			// _ = "end of CoverTab[56836]"
		case protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:465
			_go_fuzz_dep_.CoverTab[56837]++
															return coderFixed64SliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:466
			// _ = "end of CoverTab[56837]"
		case protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:467
			_go_fuzz_dep_.CoverTab[56838]++
															return coderDoubleSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:468
			// _ = "end of CoverTab[56838]"
		case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:469
			_go_fuzz_dep_.CoverTab[56839]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:473
			return coderStringSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:473
			// _ = "end of CoverTab[56839]"
		case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:474
			_go_fuzz_dep_.CoverTab[56840]++
															return coderBytesSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:475
			// _ = "end of CoverTab[56840]"
		case protoreflect.MessageKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:476
			_go_fuzz_dep_.CoverTab[56841]++
															return coderMessageSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:477
			// _ = "end of CoverTab[56841]"
		case protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:478
			_go_fuzz_dep_.CoverTab[56842]++
															return coderGroupSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:479
			// _ = "end of CoverTab[56842]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:479
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:479
			_go_fuzz_dep_.CoverTab[56843]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:479
			// _ = "end of CoverTab[56843]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:480
		// _ = "end of CoverTab[56821]"
	case fd.Cardinality() == protoreflect.Repeated && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:481
		_go_fuzz_dep_.CoverTab[56844]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:481
		return fd.IsPacked()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:481
		// _ = "end of CoverTab[56844]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:481
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:481
		_go_fuzz_dep_.CoverTab[56822]++
														switch fd.Kind() {
		case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:483
			_go_fuzz_dep_.CoverTab[56845]++
															return coderBoolPackedSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:484
			// _ = "end of CoverTab[56845]"
		case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:485
			_go_fuzz_dep_.CoverTab[56846]++
															return coderEnumPackedSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:486
			// _ = "end of CoverTab[56846]"
		case protoreflect.Int32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:487
			_go_fuzz_dep_.CoverTab[56847]++
															return coderInt32PackedSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:488
			// _ = "end of CoverTab[56847]"
		case protoreflect.Sint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:489
			_go_fuzz_dep_.CoverTab[56848]++
															return coderSint32PackedSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:490
			// _ = "end of CoverTab[56848]"
		case protoreflect.Uint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:491
			_go_fuzz_dep_.CoverTab[56849]++
															return coderUint32PackedSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:492
			// _ = "end of CoverTab[56849]"
		case protoreflect.Int64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:493
			_go_fuzz_dep_.CoverTab[56850]++
															return coderInt64PackedSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:494
			// _ = "end of CoverTab[56850]"
		case protoreflect.Sint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:495
			_go_fuzz_dep_.CoverTab[56851]++
															return coderSint64PackedSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:496
			// _ = "end of CoverTab[56851]"
		case protoreflect.Uint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:497
			_go_fuzz_dep_.CoverTab[56852]++
															return coderUint64PackedSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:498
			// _ = "end of CoverTab[56852]"
		case protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:499
			_go_fuzz_dep_.CoverTab[56853]++
															return coderSfixed32PackedSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:500
			// _ = "end of CoverTab[56853]"
		case protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:501
			_go_fuzz_dep_.CoverTab[56854]++
															return coderFixed32PackedSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:502
			// _ = "end of CoverTab[56854]"
		case protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:503
			_go_fuzz_dep_.CoverTab[56855]++
															return coderFloatPackedSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:504
			// _ = "end of CoverTab[56855]"
		case protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:505
			_go_fuzz_dep_.CoverTab[56856]++
															return coderSfixed64PackedSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:506
			// _ = "end of CoverTab[56856]"
		case protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:507
			_go_fuzz_dep_.CoverTab[56857]++
															return coderFixed64PackedSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:508
			// _ = "end of CoverTab[56857]"
		case protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:509
			_go_fuzz_dep_.CoverTab[56858]++
															return coderDoublePackedSliceValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:510
			// _ = "end of CoverTab[56858]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:510
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:510
			_go_fuzz_dep_.CoverTab[56859]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:510
			// _ = "end of CoverTab[56859]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:511
		// _ = "end of CoverTab[56822]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:512
		_go_fuzz_dep_.CoverTab[56823]++
														switch fd.Kind() {
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:514
			_go_fuzz_dep_.CoverTab[56860]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:514
			// _ = "end of CoverTab[56860]"
		case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:515
			_go_fuzz_dep_.CoverTab[56861]++
															return coderBoolValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:516
			// _ = "end of CoverTab[56861]"
		case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:517
			_go_fuzz_dep_.CoverTab[56862]++
															return coderEnumValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:518
			// _ = "end of CoverTab[56862]"
		case protoreflect.Int32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:519
			_go_fuzz_dep_.CoverTab[56863]++
															return coderInt32Value
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:520
			// _ = "end of CoverTab[56863]"
		case protoreflect.Sint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:521
			_go_fuzz_dep_.CoverTab[56864]++
															return coderSint32Value
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:522
			// _ = "end of CoverTab[56864]"
		case protoreflect.Uint32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:523
			_go_fuzz_dep_.CoverTab[56865]++
															return coderUint32Value
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:524
			// _ = "end of CoverTab[56865]"
		case protoreflect.Int64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:525
			_go_fuzz_dep_.CoverTab[56866]++
															return coderInt64Value
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:526
			// _ = "end of CoverTab[56866]"
		case protoreflect.Sint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:527
			_go_fuzz_dep_.CoverTab[56867]++
															return coderSint64Value
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:528
			// _ = "end of CoverTab[56867]"
		case protoreflect.Uint64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:529
			_go_fuzz_dep_.CoverTab[56868]++
															return coderUint64Value
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:530
			// _ = "end of CoverTab[56868]"
		case protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:531
			_go_fuzz_dep_.CoverTab[56869]++
															return coderSfixed32Value
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:532
			// _ = "end of CoverTab[56869]"
		case protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:533
			_go_fuzz_dep_.CoverTab[56870]++
															return coderFixed32Value
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:534
			// _ = "end of CoverTab[56870]"
		case protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:535
			_go_fuzz_dep_.CoverTab[56871]++
															return coderFloatValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:536
			// _ = "end of CoverTab[56871]"
		case protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:537
			_go_fuzz_dep_.CoverTab[56872]++
															return coderSfixed64Value
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:538
			// _ = "end of CoverTab[56872]"
		case protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:539
			_go_fuzz_dep_.CoverTab[56873]++
															return coderFixed64Value
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:540
			// _ = "end of CoverTab[56873]"
		case protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:541
			_go_fuzz_dep_.CoverTab[56874]++
															return coderDoubleValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:542
			// _ = "end of CoverTab[56874]"
		case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:543
			_go_fuzz_dep_.CoverTab[56875]++
															if strs.EnforceUTF8(fd) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:544
				_go_fuzz_dep_.CoverTab[56880]++
																return coderStringValueValidateUTF8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:545
				// _ = "end of CoverTab[56880]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:546
				_go_fuzz_dep_.CoverTab[56881]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:546
				// _ = "end of CoverTab[56881]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:546
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:546
			// _ = "end of CoverTab[56875]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:546
			_go_fuzz_dep_.CoverTab[56876]++
															return coderStringValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:547
			// _ = "end of CoverTab[56876]"
		case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:548
			_go_fuzz_dep_.CoverTab[56877]++
															return coderBytesValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:549
			// _ = "end of CoverTab[56877]"
		case protoreflect.MessageKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:550
			_go_fuzz_dep_.CoverTab[56878]++
															return coderMessageValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:551
			// _ = "end of CoverTab[56878]"
		case protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:552
			_go_fuzz_dep_.CoverTab[56879]++
															return coderGroupValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:553
			// _ = "end of CoverTab[56879]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:554
		// _ = "end of CoverTab[56823]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:555
	// _ = "end of CoverTab[56819]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:555
	_go_fuzz_dep_.CoverTab[56820]++
													panic(fmt.Sprintf("invalid field: no encoder for %v %v %v", fd.FullName(), fd.Cardinality(), fd.Kind()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:556
	// _ = "end of CoverTab[56820]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:557
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_tables.go:557
var _ = _go_fuzz_dep_.CoverTab
