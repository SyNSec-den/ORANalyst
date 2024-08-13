// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:5
)

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/internal/detrand"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type reflectMessageInfo struct {
	fields	map[protoreflect.FieldNumber]*fieldInfo
	oneofs	map[protoreflect.Name]*oneofInfo

	// fieldTypes contains the zero value of an enum or message field.
	// For lists, it contains the element type.
	// For maps, it contains the entry value type.
	fieldTypes	map[protoreflect.FieldNumber]interface{}

	// denseFields is a subset of fields where:
	//	0 < fieldDesc.Number() < len(denseFields)
	// It provides faster access to the fieldInfo, but may be incomplete.
	denseFields	[]*fieldInfo

	// rangeInfos is a list of all fields (not belonging to a oneof) and oneofs.
	rangeInfos	[]interface{}	// either *fieldInfo or *oneofInfo

	getUnknown	func(pointer) protoreflect.RawFields
	setUnknown	func(pointer, protoreflect.RawFields)
	extensionMap	func(pointer) *extensionMap

	nilMessage	atomicNilMessage
}

// makeReflectFuncs generates the set of functions to support reflection.
func (mi *MessageInfo) makeReflectFuncs(t reflect.Type, si structInfo) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:41
	_go_fuzz_dep_.CoverTab[58095]++
														mi.makeKnownFieldsFunc(si)
														mi.makeUnknownFieldsFunc(t, si)
														mi.makeExtensionFieldsFunc(t, si)
														mi.makeFieldTypes(si)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:45
	// _ = "end of CoverTab[58095]"
}

// makeKnownFieldsFunc generates functions for operations that can be performed
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:48
// on each protobuf message field. It takes in a reflect.Type representing the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:48
// Go struct and matches message fields with struct fields.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:48
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:48
// This code assumes that the struct is well-formed and panics if there are
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:48
// any discrepancies.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:54
func (mi *MessageInfo) makeKnownFieldsFunc(si structInfo) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:54
	_go_fuzz_dep_.CoverTab[58096]++
														mi.fields = map[protoreflect.FieldNumber]*fieldInfo{}
														md := mi.Desc
														fds := md.Fields()
														for i := 0; i < fds.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:58
		_go_fuzz_dep_.CoverTab[58101]++
															fd := fds.Get(i)
															fs := si.fieldsByNumber[fd.Number()]
															isOneof := fd.ContainingOneof() != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:61
			_go_fuzz_dep_.CoverTab[58104]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:61
			return !fd.ContainingOneof().IsSynthetic()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:61
			// _ = "end of CoverTab[58104]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:61
		}()
															if isOneof {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:62
			_go_fuzz_dep_.CoverTab[58105]++
																fs = si.oneofsByName[fd.ContainingOneof().Name()]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:63
			// _ = "end of CoverTab[58105]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:64
			_go_fuzz_dep_.CoverTab[58106]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:64
			// _ = "end of CoverTab[58106]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:64
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:64
		// _ = "end of CoverTab[58101]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:64
		_go_fuzz_dep_.CoverTab[58102]++
															var fi fieldInfo
															switch {
		case fs.Type == nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:67
			_go_fuzz_dep_.CoverTab[58107]++
																fi = fieldInfoForMissing(fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:68
			// _ = "end of CoverTab[58107]"
		case isOneof:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:69
			_go_fuzz_dep_.CoverTab[58108]++
																fi = fieldInfoForOneof(fd, fs, mi.Exporter, si.oneofWrappersByNumber[fd.Number()])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:70
			// _ = "end of CoverTab[58108]"
		case fd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:71
			_go_fuzz_dep_.CoverTab[58109]++
																fi = fieldInfoForMap(fd, fs, mi.Exporter)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:72
			// _ = "end of CoverTab[58109]"
		case fd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:73
			_go_fuzz_dep_.CoverTab[58110]++
																fi = fieldInfoForList(fd, fs, mi.Exporter)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:74
			// _ = "end of CoverTab[58110]"
		case fd.IsWeak():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:75
			_go_fuzz_dep_.CoverTab[58111]++
																fi = fieldInfoForWeakMessage(fd, si.weakOffset)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:76
			// _ = "end of CoverTab[58111]"
		case fd.Message() != nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:77
			_go_fuzz_dep_.CoverTab[58112]++
																fi = fieldInfoForMessage(fd, fs, mi.Exporter)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:78
			// _ = "end of CoverTab[58112]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:79
			_go_fuzz_dep_.CoverTab[58113]++
																fi = fieldInfoForScalar(fd, fs, mi.Exporter)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:80
			// _ = "end of CoverTab[58113]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:81
		// _ = "end of CoverTab[58102]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:81
		_go_fuzz_dep_.CoverTab[58103]++
															mi.fields[fd.Number()] = &fi
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:82
		// _ = "end of CoverTab[58103]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:83
	// _ = "end of CoverTab[58096]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:83
	_go_fuzz_dep_.CoverTab[58097]++

														mi.oneofs = map[protoreflect.Name]*oneofInfo{}
														for i := 0; i < md.Oneofs().Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:86
		_go_fuzz_dep_.CoverTab[58114]++
															od := md.Oneofs().Get(i)
															mi.oneofs[od.Name()] = makeOneofInfo(od, si, mi.Exporter)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:88
		// _ = "end of CoverTab[58114]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:89
	// _ = "end of CoverTab[58097]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:89
	_go_fuzz_dep_.CoverTab[58098]++

														mi.denseFields = make([]*fieldInfo, fds.Len()*2)
														for i := 0; i < fds.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:92
		_go_fuzz_dep_.CoverTab[58115]++
															if fd := fds.Get(i); int(fd.Number()) < len(mi.denseFields) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:93
			_go_fuzz_dep_.CoverTab[58116]++
																mi.denseFields[fd.Number()] = mi.fields[fd.Number()]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:94
			// _ = "end of CoverTab[58116]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:95
			_go_fuzz_dep_.CoverTab[58117]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:95
			// _ = "end of CoverTab[58117]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:95
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:95
		// _ = "end of CoverTab[58115]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:96
	// _ = "end of CoverTab[58098]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:96
	_go_fuzz_dep_.CoverTab[58099]++

														for i := 0; i < fds.Len(); {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:98
		_go_fuzz_dep_.CoverTab[58118]++
															fd := fds.Get(i)
															if od := fd.ContainingOneof(); od != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:100
			_go_fuzz_dep_.CoverTab[58119]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:100
			return !od.IsSynthetic()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:100
			// _ = "end of CoverTab[58119]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:100
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:100
			_go_fuzz_dep_.CoverTab[58120]++
																mi.rangeInfos = append(mi.rangeInfos, mi.oneofs[od.Name()])
																i += od.Fields().Len()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:102
			// _ = "end of CoverTab[58120]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:103
			_go_fuzz_dep_.CoverTab[58121]++
																mi.rangeInfos = append(mi.rangeInfos, mi.fields[fd.Number()])
																i++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:105
			// _ = "end of CoverTab[58121]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:106
		// _ = "end of CoverTab[58118]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:107
	// _ = "end of CoverTab[58099]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:107
	_go_fuzz_dep_.CoverTab[58100]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:110
	if len(mi.rangeInfos) > 1 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:110
		_go_fuzz_dep_.CoverTab[58122]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:110
		return detrand.Bool()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:110
		// _ = "end of CoverTab[58122]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:110
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:110
		_go_fuzz_dep_.CoverTab[58123]++
															i := detrand.Intn(len(mi.rangeInfos) - 1)
															mi.rangeInfos[i], mi.rangeInfos[i+1] = mi.rangeInfos[i+1], mi.rangeInfos[i]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:112
		// _ = "end of CoverTab[58123]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:113
		_go_fuzz_dep_.CoverTab[58124]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:113
		// _ = "end of CoverTab[58124]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:113
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:113
	// _ = "end of CoverTab[58100]"
}

func (mi *MessageInfo) makeUnknownFieldsFunc(t reflect.Type, si structInfo) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:116
	_go_fuzz_dep_.CoverTab[58125]++
														switch {
	case si.unknownOffset.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:118
		_go_fuzz_dep_.CoverTab[58132]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:118
		return si.unknownType == unknownFieldsAType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:118
		// _ = "end of CoverTab[58132]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:118
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:118
		_go_fuzz_dep_.CoverTab[58126]++

															mi.getUnknown = func(p pointer) protoreflect.RawFields {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:120
			_go_fuzz_dep_.CoverTab[58133]++
																if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:121
				_go_fuzz_dep_.CoverTab[58135]++
																	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:122
				// _ = "end of CoverTab[58135]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:123
				_go_fuzz_dep_.CoverTab[58136]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:123
				// _ = "end of CoverTab[58136]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:123
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:123
			// _ = "end of CoverTab[58133]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:123
			_go_fuzz_dep_.CoverTab[58134]++
																return *p.Apply(mi.unknownOffset).Bytes()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:124
			// _ = "end of CoverTab[58134]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:125
		// _ = "end of CoverTab[58126]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:125
		_go_fuzz_dep_.CoverTab[58127]++
															mi.setUnknown = func(p pointer, b protoreflect.RawFields) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:126
			_go_fuzz_dep_.CoverTab[58137]++
																if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:127
				_go_fuzz_dep_.CoverTab[58139]++
																	panic("invalid SetUnknown on nil Message")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:128
				// _ = "end of CoverTab[58139]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:129
				_go_fuzz_dep_.CoverTab[58140]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:129
				// _ = "end of CoverTab[58140]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:129
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:129
			// _ = "end of CoverTab[58137]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:129
			_go_fuzz_dep_.CoverTab[58138]++
																*p.Apply(mi.unknownOffset).Bytes() = b
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:130
			// _ = "end of CoverTab[58138]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:131
		// _ = "end of CoverTab[58127]"
	case si.unknownOffset.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:132
		_go_fuzz_dep_.CoverTab[58141]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:132
		return si.unknownType == unknownFieldsBType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:132
		// _ = "end of CoverTab[58141]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:132
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:132
		_go_fuzz_dep_.CoverTab[58128]++

															mi.getUnknown = func(p pointer) protoreflect.RawFields {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:134
			_go_fuzz_dep_.CoverTab[58142]++
																if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:135
				_go_fuzz_dep_.CoverTab[58145]++
																	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:136
				// _ = "end of CoverTab[58145]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:137
				_go_fuzz_dep_.CoverTab[58146]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:137
				// _ = "end of CoverTab[58146]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:137
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:137
			// _ = "end of CoverTab[58142]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:137
			_go_fuzz_dep_.CoverTab[58143]++
																bp := p.Apply(mi.unknownOffset).BytesPtr()
																if *bp == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:139
				_go_fuzz_dep_.CoverTab[58147]++
																	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:140
				// _ = "end of CoverTab[58147]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:141
				_go_fuzz_dep_.CoverTab[58148]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:141
				// _ = "end of CoverTab[58148]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:141
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:141
			// _ = "end of CoverTab[58143]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:141
			_go_fuzz_dep_.CoverTab[58144]++
																return **bp
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:142
			// _ = "end of CoverTab[58144]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:143
		// _ = "end of CoverTab[58128]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:143
		_go_fuzz_dep_.CoverTab[58129]++
															mi.setUnknown = func(p pointer, b protoreflect.RawFields) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:144
			_go_fuzz_dep_.CoverTab[58149]++
																if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:145
				_go_fuzz_dep_.CoverTab[58152]++
																	panic("invalid SetUnknown on nil Message")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:146
				// _ = "end of CoverTab[58152]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:147
				_go_fuzz_dep_.CoverTab[58153]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:147
				// _ = "end of CoverTab[58153]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:147
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:147
			// _ = "end of CoverTab[58149]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:147
			_go_fuzz_dep_.CoverTab[58150]++
																bp := p.Apply(mi.unknownOffset).BytesPtr()
																if *bp == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:149
				_go_fuzz_dep_.CoverTab[58154]++
																	*bp = new([]byte)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:150
				// _ = "end of CoverTab[58154]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:151
				_go_fuzz_dep_.CoverTab[58155]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:151
				// _ = "end of CoverTab[58155]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:151
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:151
			// _ = "end of CoverTab[58150]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:151
			_go_fuzz_dep_.CoverTab[58151]++
																**bp = b
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:152
			// _ = "end of CoverTab[58151]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:153
		// _ = "end of CoverTab[58129]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:154
		_go_fuzz_dep_.CoverTab[58130]++
															mi.getUnknown = func(pointer) protoreflect.RawFields {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:155
			_go_fuzz_dep_.CoverTab[58156]++
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:156
			// _ = "end of CoverTab[58156]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:157
		// _ = "end of CoverTab[58130]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:157
		_go_fuzz_dep_.CoverTab[58131]++
															mi.setUnknown = func(p pointer, _ protoreflect.RawFields) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:158
			_go_fuzz_dep_.CoverTab[58157]++
																if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:159
				_go_fuzz_dep_.CoverTab[58158]++
																	panic("invalid SetUnknown on nil Message")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:160
				// _ = "end of CoverTab[58158]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:161
				_go_fuzz_dep_.CoverTab[58159]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:161
				// _ = "end of CoverTab[58159]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:161
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:161
			// _ = "end of CoverTab[58157]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:162
		// _ = "end of CoverTab[58131]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:163
	// _ = "end of CoverTab[58125]"
}

func (mi *MessageInfo) makeExtensionFieldsFunc(t reflect.Type, si structInfo) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:166
	_go_fuzz_dep_.CoverTab[58160]++
														if si.extensionOffset.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:167
		_go_fuzz_dep_.CoverTab[58161]++
															mi.extensionMap = func(p pointer) *extensionMap {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:168
			_go_fuzz_dep_.CoverTab[58162]++
																if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:169
				_go_fuzz_dep_.CoverTab[58164]++
																	return (*extensionMap)(nil)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:170
				// _ = "end of CoverTab[58164]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:171
				_go_fuzz_dep_.CoverTab[58165]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:171
				// _ = "end of CoverTab[58165]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:171
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:171
			// _ = "end of CoverTab[58162]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:171
			_go_fuzz_dep_.CoverTab[58163]++
																v := p.Apply(si.extensionOffset).AsValueOf(extensionFieldsType)
																return (*extensionMap)(v.Interface().(*map[int32]ExtensionField))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:173
			// _ = "end of CoverTab[58163]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:174
		// _ = "end of CoverTab[58161]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:175
		_go_fuzz_dep_.CoverTab[58166]++
															mi.extensionMap = func(pointer) *extensionMap {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:176
			_go_fuzz_dep_.CoverTab[58167]++
																return (*extensionMap)(nil)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:177
			// _ = "end of CoverTab[58167]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:178
		// _ = "end of CoverTab[58166]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:179
	// _ = "end of CoverTab[58160]"
}
func (mi *MessageInfo) makeFieldTypes(si structInfo) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:181
	_go_fuzz_dep_.CoverTab[58168]++
														md := mi.Desc
														fds := md.Fields()
														for i := 0; i < fds.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:184
		_go_fuzz_dep_.CoverTab[58169]++
															var ft reflect.Type
															fd := fds.Get(i)
															fs := si.fieldsByNumber[fd.Number()]
															isOneof := fd.ContainingOneof() != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:188
			_go_fuzz_dep_.CoverTab[58173]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:188
			return !fd.ContainingOneof().IsSynthetic()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:188
			// _ = "end of CoverTab[58173]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:188
		}()
															if isOneof {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:189
			_go_fuzz_dep_.CoverTab[58174]++
																fs = si.oneofsByName[fd.ContainingOneof().Name()]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:190
			// _ = "end of CoverTab[58174]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:191
			_go_fuzz_dep_.CoverTab[58175]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:191
			// _ = "end of CoverTab[58175]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:191
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:191
		// _ = "end of CoverTab[58169]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:191
		_go_fuzz_dep_.CoverTab[58170]++
															var isMessage bool
															switch {
		case fs.Type == nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:194
			_go_fuzz_dep_.CoverTab[58176]++
																continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:195
			// _ = "end of CoverTab[58176]"
		case isOneof:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:196
			_go_fuzz_dep_.CoverTab[58177]++
																if fd.Enum() != nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:197
				_go_fuzz_dep_.CoverTab[58186]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:197
				return fd.Message() != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:197
				// _ = "end of CoverTab[58186]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:197
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:197
				_go_fuzz_dep_.CoverTab[58187]++
																	ft = si.oneofWrappersByNumber[fd.Number()].Field(0).Type
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:198
				// _ = "end of CoverTab[58187]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:199
				_go_fuzz_dep_.CoverTab[58188]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:199
				// _ = "end of CoverTab[58188]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:199
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:199
			// _ = "end of CoverTab[58177]"
		case fd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:200
			_go_fuzz_dep_.CoverTab[58178]++
																if fd.MapValue().Enum() != nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:201
				_go_fuzz_dep_.CoverTab[58189]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:201
				return fd.MapValue().Message() != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:201
				// _ = "end of CoverTab[58189]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:201
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:201
				_go_fuzz_dep_.CoverTab[58190]++
																	ft = fs.Type.Elem()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:202
				// _ = "end of CoverTab[58190]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:203
				_go_fuzz_dep_.CoverTab[58191]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:203
				// _ = "end of CoverTab[58191]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:203
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:203
			// _ = "end of CoverTab[58178]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:203
			_go_fuzz_dep_.CoverTab[58179]++
																isMessage = fd.MapValue().Message() != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:204
			// _ = "end of CoverTab[58179]"
		case fd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:205
			_go_fuzz_dep_.CoverTab[58180]++
																if fd.Enum() != nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:206
				_go_fuzz_dep_.CoverTab[58192]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:206
				return fd.Message() != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:206
				// _ = "end of CoverTab[58192]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:206
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:206
				_go_fuzz_dep_.CoverTab[58193]++
																	ft = fs.Type.Elem()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:207
				// _ = "end of CoverTab[58193]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:208
				_go_fuzz_dep_.CoverTab[58194]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:208
				// _ = "end of CoverTab[58194]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:208
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:208
			// _ = "end of CoverTab[58180]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:208
			_go_fuzz_dep_.CoverTab[58181]++
																isMessage = fd.Message() != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:209
			// _ = "end of CoverTab[58181]"
		case fd.Enum() != nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:210
			_go_fuzz_dep_.CoverTab[58182]++
																ft = fs.Type
																if fd.HasPresence() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:212
				_go_fuzz_dep_.CoverTab[58195]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:212
				return ft.Kind() == reflect.Ptr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:212
				// _ = "end of CoverTab[58195]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:212
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:212
				_go_fuzz_dep_.CoverTab[58196]++
																	ft = ft.Elem()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:213
				// _ = "end of CoverTab[58196]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:214
				_go_fuzz_dep_.CoverTab[58197]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:214
				// _ = "end of CoverTab[58197]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:214
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:214
			// _ = "end of CoverTab[58182]"
		case fd.Message() != nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:215
			_go_fuzz_dep_.CoverTab[58183]++
																ft = fs.Type
																if fd.IsWeak() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:217
				_go_fuzz_dep_.CoverTab[58198]++
																	ft = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:218
				// _ = "end of CoverTab[58198]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:219
				_go_fuzz_dep_.CoverTab[58199]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:219
				// _ = "end of CoverTab[58199]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:219
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:219
			// _ = "end of CoverTab[58183]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:219
			_go_fuzz_dep_.CoverTab[58184]++
																isMessage = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:220
			// _ = "end of CoverTab[58184]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:220
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:220
			_go_fuzz_dep_.CoverTab[58185]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:220
			// _ = "end of CoverTab[58185]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:221
		// _ = "end of CoverTab[58170]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:221
		_go_fuzz_dep_.CoverTab[58171]++
															if isMessage && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:222
			_go_fuzz_dep_.CoverTab[58200]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:222
			return ft != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:222
			// _ = "end of CoverTab[58200]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:222
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:222
			_go_fuzz_dep_.CoverTab[58201]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:222
			return ft.Kind() != reflect.Ptr
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:222
			// _ = "end of CoverTab[58201]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:222
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:222
			_go_fuzz_dep_.CoverTab[58202]++
																ft = reflect.PtrTo(ft)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:223
			// _ = "end of CoverTab[58202]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:224
			_go_fuzz_dep_.CoverTab[58203]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:224
			// _ = "end of CoverTab[58203]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:224
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:224
		// _ = "end of CoverTab[58171]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:224
		_go_fuzz_dep_.CoverTab[58172]++
															if ft != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:225
			_go_fuzz_dep_.CoverTab[58204]++
																if mi.fieldTypes == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:226
				_go_fuzz_dep_.CoverTab[58206]++
																	mi.fieldTypes = make(map[protoreflect.FieldNumber]interface{})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:227
				// _ = "end of CoverTab[58206]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:228
				_go_fuzz_dep_.CoverTab[58207]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:228
				// _ = "end of CoverTab[58207]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:228
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:228
			// _ = "end of CoverTab[58204]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:228
			_go_fuzz_dep_.CoverTab[58205]++
																mi.fieldTypes[fd.Number()] = reflect.Zero(ft).Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:229
			// _ = "end of CoverTab[58205]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:230
			_go_fuzz_dep_.CoverTab[58208]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:230
			// _ = "end of CoverTab[58208]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:230
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:230
		// _ = "end of CoverTab[58172]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:231
	// _ = "end of CoverTab[58168]"
}

type extensionMap map[int32]ExtensionField

func (m *extensionMap) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:236
	_go_fuzz_dep_.CoverTab[58209]++
														if m != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:237
		_go_fuzz_dep_.CoverTab[58210]++
															for _, x := range *m {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:238
			_go_fuzz_dep_.CoverTab[58211]++
																xd := x.Type().TypeDescriptor()
																v := x.Value()
																if xd.IsList() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:241
				_go_fuzz_dep_.CoverTab[58213]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:241
				return v.List().Len() == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:241
				// _ = "end of CoverTab[58213]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:241
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:241
				_go_fuzz_dep_.CoverTab[58214]++
																	continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:242
				// _ = "end of CoverTab[58214]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:243
				_go_fuzz_dep_.CoverTab[58215]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:243
				// _ = "end of CoverTab[58215]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:243
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:243
			// _ = "end of CoverTab[58211]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:243
			_go_fuzz_dep_.CoverTab[58212]++
																if !f(xd, v) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:244
				_go_fuzz_dep_.CoverTab[58216]++
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:245
				// _ = "end of CoverTab[58216]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:246
				_go_fuzz_dep_.CoverTab[58217]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:246
				// _ = "end of CoverTab[58217]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:246
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:246
			// _ = "end of CoverTab[58212]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:247
		// _ = "end of CoverTab[58210]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:248
		_go_fuzz_dep_.CoverTab[58218]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:248
		// _ = "end of CoverTab[58218]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:248
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:248
	// _ = "end of CoverTab[58209]"
}
func (m *extensionMap) Has(xt protoreflect.ExtensionType) (ok bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:250
	_go_fuzz_dep_.CoverTab[58219]++
														if m == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:251
		_go_fuzz_dep_.CoverTab[58223]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:252
		// _ = "end of CoverTab[58223]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:253
		_go_fuzz_dep_.CoverTab[58224]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:253
		// _ = "end of CoverTab[58224]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:253
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:253
	// _ = "end of CoverTab[58219]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:253
	_go_fuzz_dep_.CoverTab[58220]++
														xd := xt.TypeDescriptor()
														x, ok := (*m)[int32(xd.Number())]
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:256
		_go_fuzz_dep_.CoverTab[58225]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:257
		// _ = "end of CoverTab[58225]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:258
		_go_fuzz_dep_.CoverTab[58226]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:258
		// _ = "end of CoverTab[58226]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:258
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:258
	// _ = "end of CoverTab[58220]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:258
	_go_fuzz_dep_.CoverTab[58221]++
														switch {
	case xd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:260
		_go_fuzz_dep_.CoverTab[58227]++
															return x.Value().List().Len() > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:261
		// _ = "end of CoverTab[58227]"
	case xd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:262
		_go_fuzz_dep_.CoverTab[58228]++
															return x.Value().Map().Len() > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:263
		// _ = "end of CoverTab[58228]"
	case xd.Message() != nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:264
		_go_fuzz_dep_.CoverTab[58229]++
															return x.Value().Message().IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:265
		// _ = "end of CoverTab[58229]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:265
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:265
		_go_fuzz_dep_.CoverTab[58230]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:265
		// _ = "end of CoverTab[58230]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:266
	// _ = "end of CoverTab[58221]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:266
	_go_fuzz_dep_.CoverTab[58222]++
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:267
	// _ = "end of CoverTab[58222]"
}
func (m *extensionMap) Clear(xt protoreflect.ExtensionType) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:269
	_go_fuzz_dep_.CoverTab[58231]++
														delete(*m, int32(xt.TypeDescriptor().Number()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:270
	// _ = "end of CoverTab[58231]"
}
func (m *extensionMap) Get(xt protoreflect.ExtensionType) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:272
	_go_fuzz_dep_.CoverTab[58232]++
														xd := xt.TypeDescriptor()
														if m != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:274
		_go_fuzz_dep_.CoverTab[58234]++
															if x, ok := (*m)[int32(xd.Number())]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:275
			_go_fuzz_dep_.CoverTab[58235]++
																return x.Value()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:276
			// _ = "end of CoverTab[58235]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:277
			_go_fuzz_dep_.CoverTab[58236]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:277
			// _ = "end of CoverTab[58236]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:277
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:277
		// _ = "end of CoverTab[58234]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:278
		_go_fuzz_dep_.CoverTab[58237]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:278
		// _ = "end of CoverTab[58237]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:278
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:278
	// _ = "end of CoverTab[58232]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:278
	_go_fuzz_dep_.CoverTab[58233]++
														return xt.Zero()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:279
	// _ = "end of CoverTab[58233]"
}
func (m *extensionMap) Set(xt protoreflect.ExtensionType, v protoreflect.Value) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:281
	_go_fuzz_dep_.CoverTab[58238]++
														xd := xt.TypeDescriptor()
														isValid := true
														switch {
	case !xt.IsValidValue(v):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:285
		_go_fuzz_dep_.CoverTab[58242]++
															isValid = false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:286
		// _ = "end of CoverTab[58242]"
	case xd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:287
		_go_fuzz_dep_.CoverTab[58243]++
															isValid = v.List().IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:288
		// _ = "end of CoverTab[58243]"
	case xd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:289
		_go_fuzz_dep_.CoverTab[58244]++
															isValid = v.Map().IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:290
		// _ = "end of CoverTab[58244]"
	case xd.Message() != nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:291
		_go_fuzz_dep_.CoverTab[58245]++
															isValid = v.Message().IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:292
		// _ = "end of CoverTab[58245]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:292
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:292
		_go_fuzz_dep_.CoverTab[58246]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:292
		// _ = "end of CoverTab[58246]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:293
	// _ = "end of CoverTab[58238]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:293
	_go_fuzz_dep_.CoverTab[58239]++
														if !isValid {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:294
		_go_fuzz_dep_.CoverTab[58247]++
															panic(fmt.Sprintf("%v: assigning invalid value", xt.TypeDescriptor().FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:295
		// _ = "end of CoverTab[58247]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:296
		_go_fuzz_dep_.CoverTab[58248]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:296
		// _ = "end of CoverTab[58248]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:296
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:296
	// _ = "end of CoverTab[58239]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:296
	_go_fuzz_dep_.CoverTab[58240]++

														if *m == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:298
		_go_fuzz_dep_.CoverTab[58249]++
															*m = make(map[int32]ExtensionField)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:299
		// _ = "end of CoverTab[58249]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:300
		_go_fuzz_dep_.CoverTab[58250]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:300
		// _ = "end of CoverTab[58250]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:300
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:300
	// _ = "end of CoverTab[58240]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:300
	_go_fuzz_dep_.CoverTab[58241]++
														var x ExtensionField
														x.Set(xt, v)
														(*m)[int32(xd.Number())] = x
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:303
	// _ = "end of CoverTab[58241]"
}
func (m *extensionMap) Mutable(xt protoreflect.ExtensionType) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:305
	_go_fuzz_dep_.CoverTab[58251]++
														xd := xt.TypeDescriptor()
														if xd.Kind() != protoreflect.MessageKind && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:307
		_go_fuzz_dep_.CoverTab[58254]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:307
		return xd.Kind() != protoreflect.GroupKind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:307
		// _ = "end of CoverTab[58254]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:307
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:307
		_go_fuzz_dep_.CoverTab[58255]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:307
		return !xd.IsList()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:307
		// _ = "end of CoverTab[58255]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:307
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:307
		_go_fuzz_dep_.CoverTab[58256]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:307
		return !xd.IsMap()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:307
		// _ = "end of CoverTab[58256]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:307
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:307
		_go_fuzz_dep_.CoverTab[58257]++
															panic("invalid Mutable on field with non-composite type")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:308
		// _ = "end of CoverTab[58257]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:309
		_go_fuzz_dep_.CoverTab[58258]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:309
		// _ = "end of CoverTab[58258]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:309
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:309
	// _ = "end of CoverTab[58251]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:309
	_go_fuzz_dep_.CoverTab[58252]++
														if x, ok := (*m)[int32(xd.Number())]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:310
		_go_fuzz_dep_.CoverTab[58259]++
															return x.Value()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:311
		// _ = "end of CoverTab[58259]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:312
		_go_fuzz_dep_.CoverTab[58260]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:312
		// _ = "end of CoverTab[58260]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:312
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:312
	// _ = "end of CoverTab[58252]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:312
	_go_fuzz_dep_.CoverTab[58253]++
														v := xt.New()
														m.Set(xt, v)
														return v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:315
	// _ = "end of CoverTab[58253]"
}

// MessageState is a data structure that is nested as the first field in a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
// concrete message. It provides a way to implement the ProtoReflect method
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
// in an allocation-free way without needing to have a shadow Go type generated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
// for every message type. This technique only works using unsafe.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
// Example generated code:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//	type M struct {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//		state protoimpl.MessageState
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//		Field1 int32
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//		Field2 string
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//		Field3 *BarMessage
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//		...
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//	func (m *M) ProtoReflect() protoreflect.Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//		mi := &file_fizz_buzz_proto_msgInfos[5]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//		if protoimpl.UnsafeEnabled && m != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//			ms := protoimpl.X.MessageStateOf(Pointer(m))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//			if ms.LoadMessageInfo() == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//				ms.StoreMessageInfo(mi)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//			return ms
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//		return mi.MessageOf(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
// The MessageState type holds a *MessageInfo, which must be atomically set to
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
// the message info associated with a given message instance.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
// By unsafely converting a *M into a *MessageState, the MessageState object
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
// has access to all the information needed to implement protobuf reflection.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
// It has access to the message info as its first field, and a pointer to the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
// MessageState is identical to a pointer to the concrete message value.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
// Requirements:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//   - The type M must implement protoreflect.ProtoMessage.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//   - The address of m must not be nil.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//   - The address of m and the address of m.state must be equal,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:318
//     even though they are different Go types.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:358
type MessageState struct {
	pragma.NoUnkeyedLiterals
	pragma.DoNotCompare
	pragma.DoNotCopy

	atomicMessageInfo	*MessageInfo
}

type messageState MessageState

var (
	_	protoreflect.Message	= (*messageState)(nil)
	_	unwrapper		= (*messageState)(nil)
)

// messageDataType is a tuple of a pointer to the message data and
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:373
// a pointer to the message type. It is a generalized way of providing a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:373
// reflective view over a message instance. The disadvantage of this approach
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:373
// is the need to allocate this tuple of 16B.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:377
type messageDataType struct {
	p	pointer
	mi	*MessageInfo
}

type (
	messageReflectWrapper	messageDataType
	messageIfaceWrapper	messageDataType
)

var (
	_	protoreflect.Message		= (*messageReflectWrapper)(nil)
	_	unwrapper			= (*messageReflectWrapper)(nil)
	_	protoreflect.ProtoMessage	= (*messageIfaceWrapper)(nil)
	_	unwrapper			= (*messageIfaceWrapper)(nil)
)

// MessageOf returns a reflective view over a message. The input must be a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:394
// pointer to a named Go struct. If the provided type has a ProtoReflect method,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:394
// it must be implemented by calling this method.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:397
func (mi *MessageInfo) MessageOf(m interface{}) protoreflect.Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:397
	_go_fuzz_dep_.CoverTab[58261]++
														if reflect.TypeOf(m) != mi.GoReflectType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:398
		_go_fuzz_dep_.CoverTab[58264]++
															panic(fmt.Sprintf("type mismatch: got %T, want %v", m, mi.GoReflectType))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:399
		// _ = "end of CoverTab[58264]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:400
		_go_fuzz_dep_.CoverTab[58265]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:400
		// _ = "end of CoverTab[58265]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:400
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:400
	// _ = "end of CoverTab[58261]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:400
	_go_fuzz_dep_.CoverTab[58262]++
														p := pointerOfIface(m)
														if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:402
		_go_fuzz_dep_.CoverTab[58266]++
															return mi.nilMessage.Init(mi)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:403
		// _ = "end of CoverTab[58266]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:404
		_go_fuzz_dep_.CoverTab[58267]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:404
		// _ = "end of CoverTab[58267]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:404
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:404
	// _ = "end of CoverTab[58262]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:404
	_go_fuzz_dep_.CoverTab[58263]++
														return &messageReflectWrapper{p, mi}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:405
	// _ = "end of CoverTab[58263]"
}

func (m *messageReflectWrapper) pointer() pointer {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:408
	_go_fuzz_dep_.CoverTab[58268]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:408
	return m.p
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:408
	// _ = "end of CoverTab[58268]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:408
}
func (m *messageReflectWrapper) messageInfo() *MessageInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:409
	_go_fuzz_dep_.CoverTab[58269]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:409
	return m.mi
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:409
	// _ = "end of CoverTab[58269]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:409
}

// Reset implements the v1 proto.Message.Reset method.
func (m *messageIfaceWrapper) Reset() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:412
	_go_fuzz_dep_.CoverTab[58270]++
														if mr, ok := m.protoUnwrap().(interface{ Reset() }); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:413
		_go_fuzz_dep_.CoverTab[58272]++
															mr.Reset()
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:415
		// _ = "end of CoverTab[58272]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:416
		_go_fuzz_dep_.CoverTab[58273]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:416
		// _ = "end of CoverTab[58273]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:416
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:416
	// _ = "end of CoverTab[58270]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:416
	_go_fuzz_dep_.CoverTab[58271]++
														rv := reflect.ValueOf(m.protoUnwrap())
														if rv.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:418
		_go_fuzz_dep_.CoverTab[58274]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:418
		return !rv.IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:418
		// _ = "end of CoverTab[58274]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:418
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:418
		_go_fuzz_dep_.CoverTab[58275]++
															rv.Elem().Set(reflect.Zero(rv.Type().Elem()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:419
		// _ = "end of CoverTab[58275]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:420
		_go_fuzz_dep_.CoverTab[58276]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:420
		// _ = "end of CoverTab[58276]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:420
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:420
	// _ = "end of CoverTab[58271]"
}
func (m *messageIfaceWrapper) ProtoReflect() protoreflect.Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:422
	_go_fuzz_dep_.CoverTab[58277]++
														return (*messageReflectWrapper)(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:423
	// _ = "end of CoverTab[58277]"
}
func (m *messageIfaceWrapper) protoUnwrap() interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:425
	_go_fuzz_dep_.CoverTab[58278]++
														return m.p.AsIfaceOf(m.mi.GoReflectType.Elem())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:426
	// _ = "end of CoverTab[58278]"
}

// checkField verifies that the provided field descriptor is valid.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:429
// Exactly one of the returned values is populated.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:431
func (mi *MessageInfo) checkField(fd protoreflect.FieldDescriptor) (*fieldInfo, protoreflect.ExtensionType) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:431
	_go_fuzz_dep_.CoverTab[58279]++
														var fi *fieldInfo
														if n := fd.Number(); 0 < n && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:433
		_go_fuzz_dep_.CoverTab[58283]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:433
		return int(n) < len(mi.denseFields)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:433
		// _ = "end of CoverTab[58283]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:433
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:433
		_go_fuzz_dep_.CoverTab[58284]++
															fi = mi.denseFields[n]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:434
		// _ = "end of CoverTab[58284]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:435
		_go_fuzz_dep_.CoverTab[58285]++
															fi = mi.fields[n]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:436
		// _ = "end of CoverTab[58285]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:437
	// _ = "end of CoverTab[58279]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:437
	_go_fuzz_dep_.CoverTab[58280]++
														if fi != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:438
		_go_fuzz_dep_.CoverTab[58286]++
															if fi.fieldDesc != fd {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:439
			_go_fuzz_dep_.CoverTab[58288]++
																if got, want := fd.FullName(), fi.fieldDesc.FullName(); got != want {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:440
				_go_fuzz_dep_.CoverTab[58290]++
																	panic(fmt.Sprintf("mismatching field: got %v, want %v", got, want))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:441
				// _ = "end of CoverTab[58290]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:442
				_go_fuzz_dep_.CoverTab[58291]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:442
				// _ = "end of CoverTab[58291]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:442
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:442
			// _ = "end of CoverTab[58288]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:442
			_go_fuzz_dep_.CoverTab[58289]++
																panic(fmt.Sprintf("mismatching field: %v", fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:443
			// _ = "end of CoverTab[58289]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:444
			_go_fuzz_dep_.CoverTab[58292]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:444
			// _ = "end of CoverTab[58292]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:444
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:444
		// _ = "end of CoverTab[58286]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:444
		_go_fuzz_dep_.CoverTab[58287]++
															return fi, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:445
		// _ = "end of CoverTab[58287]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:446
		_go_fuzz_dep_.CoverTab[58293]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:446
		// _ = "end of CoverTab[58293]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:446
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:446
	// _ = "end of CoverTab[58280]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:446
	_go_fuzz_dep_.CoverTab[58281]++

														if fd.IsExtension() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:448
		_go_fuzz_dep_.CoverTab[58294]++
															if got, want := fd.ContainingMessage().FullName(), mi.Desc.FullName(); got != want {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:449
			_go_fuzz_dep_.CoverTab[58298]++

																panic(fmt.Sprintf("extension %v has mismatching containing message: got %v, want %v", fd.FullName(), got, want))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:451
			// _ = "end of CoverTab[58298]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:452
			_go_fuzz_dep_.CoverTab[58299]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:452
			// _ = "end of CoverTab[58299]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:452
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:452
		// _ = "end of CoverTab[58294]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:452
		_go_fuzz_dep_.CoverTab[58295]++
															if !mi.Desc.ExtensionRanges().Has(fd.Number()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:453
			_go_fuzz_dep_.CoverTab[58300]++
																panic(fmt.Sprintf("extension %v extends %v outside the extension range", fd.FullName(), mi.Desc.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:454
			// _ = "end of CoverTab[58300]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:455
			_go_fuzz_dep_.CoverTab[58301]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:455
			// _ = "end of CoverTab[58301]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:455
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:455
		// _ = "end of CoverTab[58295]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:455
		_go_fuzz_dep_.CoverTab[58296]++
															xtd, ok := fd.(protoreflect.ExtensionTypeDescriptor)
															if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:457
			_go_fuzz_dep_.CoverTab[58302]++
																panic(fmt.Sprintf("extension %v does not implement protoreflect.ExtensionTypeDescriptor", fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:458
			// _ = "end of CoverTab[58302]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:459
			_go_fuzz_dep_.CoverTab[58303]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:459
			// _ = "end of CoverTab[58303]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:459
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:459
		// _ = "end of CoverTab[58296]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:459
		_go_fuzz_dep_.CoverTab[58297]++
															return nil, xtd.Type()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:460
		// _ = "end of CoverTab[58297]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:461
		_go_fuzz_dep_.CoverTab[58304]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:461
		// _ = "end of CoverTab[58304]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:461
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:461
	// _ = "end of CoverTab[58281]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:461
	_go_fuzz_dep_.CoverTab[58282]++
														panic(fmt.Sprintf("field %v is invalid", fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:462
	// _ = "end of CoverTab[58282]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:463
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/message_reflect.go:463
var _ = _go_fuzz_dep_.CoverTab
