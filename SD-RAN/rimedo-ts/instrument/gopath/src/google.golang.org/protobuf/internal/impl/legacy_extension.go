// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:5
)

import (
	"reflect"

	"google.golang.org/protobuf/internal/descopts"
	"google.golang.org/protobuf/internal/encoding/messageset"
	ptag "google.golang.org/protobuf/internal/encoding/tag"
	"google.golang.org/protobuf/internal/filedesc"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoiface"
)

func (xi *ExtensionInfo) initToLegacy() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:20
	_go_fuzz_dep_.CoverTab[57542]++
														xd := xi.desc
														var parent protoiface.MessageV1
														messageName := xd.ContainingMessage().FullName()
														if mt, _ := protoregistry.GlobalTypes.FindMessageByName(messageName); mt != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:24
		_go_fuzz_dep_.CoverTab[57548]++

															mv := mt.New().Interface()
															t := reflect.TypeOf(mv)
															if mv, ok := mv.(unwrapper); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:28
			_go_fuzz_dep_.CoverTab[57550]++
																t = reflect.TypeOf(mv.protoUnwrap())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:29
			// _ = "end of CoverTab[57550]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:30
			_go_fuzz_dep_.CoverTab[57551]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:30
			// _ = "end of CoverTab[57551]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:30
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:30
		// _ = "end of CoverTab[57548]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:30
		_go_fuzz_dep_.CoverTab[57549]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:33
		mz := reflect.Zero(t).Interface()
		if mz, ok := mz.(protoiface.MessageV1); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:34
			_go_fuzz_dep_.CoverTab[57552]++
																parent = mz
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:35
			// _ = "end of CoverTab[57552]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:36
			_go_fuzz_dep_.CoverTab[57553]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:36
			// _ = "end of CoverTab[57553]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:36
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:36
		// _ = "end of CoverTab[57549]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:37
		_go_fuzz_dep_.CoverTab[57554]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:37
		// _ = "end of CoverTab[57554]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:37
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:37
	// _ = "end of CoverTab[57542]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:37
	_go_fuzz_dep_.CoverTab[57543]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:41
	extType := xi.goType
	switch extType.Kind() {
	case reflect.Bool, reflect.Int32, reflect.Int64, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.String:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:43
		_go_fuzz_dep_.CoverTab[57555]++
															extType = reflect.PtrTo(extType)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:44
		// _ = "end of CoverTab[57555]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:44
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:44
		_go_fuzz_dep_.CoverTab[57556]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:44
		// _ = "end of CoverTab[57556]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:45
	// _ = "end of CoverTab[57543]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:45
	_go_fuzz_dep_.CoverTab[57544]++

	// Reconstruct the legacy enum full name.
	var enumName string
	if xd.Kind() == protoreflect.EnumKind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:49
		_go_fuzz_dep_.CoverTab[57557]++
															enumName = legacyEnumName(xd.Enum())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:50
		// _ = "end of CoverTab[57557]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:51
		_go_fuzz_dep_.CoverTab[57558]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:51
		// _ = "end of CoverTab[57558]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:51
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:51
	// _ = "end of CoverTab[57544]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:51
	_go_fuzz_dep_.CoverTab[57545]++

	// Derive the proto file that the extension was declared within.
	var filename string
	if fd := xd.ParentFile(); fd != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:55
		_go_fuzz_dep_.CoverTab[57559]++
															filename = fd.Path()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:56
		// _ = "end of CoverTab[57559]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:57
		_go_fuzz_dep_.CoverTab[57560]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:57
		// _ = "end of CoverTab[57560]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:57
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:57
	// _ = "end of CoverTab[57545]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:57
	_go_fuzz_dep_.CoverTab[57546]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:60
	name := xd.FullName()
	if messageset.IsMessageSetExtension(xd) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:61
		_go_fuzz_dep_.CoverTab[57561]++
															name = name.Parent()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:62
		// _ = "end of CoverTab[57561]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:63
		_go_fuzz_dep_.CoverTab[57562]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:63
		// _ = "end of CoverTab[57562]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:63
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:63
	// _ = "end of CoverTab[57546]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:63
	_go_fuzz_dep_.CoverTab[57547]++

														xi.ExtendedType = parent
														xi.ExtensionType = reflect.Zero(extType).Interface()
														xi.Field = int32(xd.Number())
														xi.Name = string(name)
														xi.Tag = ptag.Marshal(xd, enumName)
														xi.Filename = filename
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:70
	// _ = "end of CoverTab[57547]"
}

// initFromLegacy initializes an ExtensionInfo from
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:73
// the contents of the deprecated exported fields of the type.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:75
func (xi *ExtensionInfo) initFromLegacy() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:75
	_go_fuzz_dep_.CoverTab[57563]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:78
	if xi.ExtendedType == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:78
		_go_fuzz_dep_.CoverTab[57570]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:78
		return xi.ExtensionType == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:78
		// _ = "end of CoverTab[57570]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:78
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:78
		_go_fuzz_dep_.CoverTab[57571]++
															xd := placeholderExtension{
			name:	protoreflect.FullName(xi.Name),
			number:	protoreflect.FieldNumber(xi.Field),
		}
															xi.desc = extensionTypeDescriptor{xd, xi}
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:84
		// _ = "end of CoverTab[57571]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:85
		_go_fuzz_dep_.CoverTab[57572]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:85
		// _ = "end of CoverTab[57572]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:85
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:85
	// _ = "end of CoverTab[57563]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:85
	_go_fuzz_dep_.CoverTab[57564]++

	// Resolve enum or message dependencies.
	var ed protoreflect.EnumDescriptor
	var md protoreflect.MessageDescriptor
	t := reflect.TypeOf(xi.ExtensionType)
	isOptional := t.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:91
		_go_fuzz_dep_.CoverTab[57573]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:91
		return t.Elem().Kind() != reflect.Struct
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:91
		// _ = "end of CoverTab[57573]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:91
	}()
														isRepeated := t.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:92
		_go_fuzz_dep_.CoverTab[57574]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:92
		return t.Elem().Kind() != reflect.Uint8
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:92
		// _ = "end of CoverTab[57574]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:92
	}()
														if isOptional || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:93
		_go_fuzz_dep_.CoverTab[57575]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:93
		return isRepeated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:93
		// _ = "end of CoverTab[57575]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:93
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:93
		_go_fuzz_dep_.CoverTab[57576]++
															t = t.Elem()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:94
		// _ = "end of CoverTab[57576]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:95
		_go_fuzz_dep_.CoverTab[57577]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:95
		// _ = "end of CoverTab[57577]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:95
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:95
	// _ = "end of CoverTab[57564]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:95
	_go_fuzz_dep_.CoverTab[57565]++
														switch v := reflect.Zero(t).Interface().(type) {
	case protoreflect.Enum:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:97
		_go_fuzz_dep_.CoverTab[57578]++
															ed = v.Descriptor()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:98
		// _ = "end of CoverTab[57578]"
	case enumV1:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:99
		_go_fuzz_dep_.CoverTab[57579]++
															ed = LegacyLoadEnumDesc(t)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:100
		// _ = "end of CoverTab[57579]"
	case protoreflect.ProtoMessage:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:101
		_go_fuzz_dep_.CoverTab[57580]++
															md = v.ProtoReflect().Descriptor()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:102
		// _ = "end of CoverTab[57580]"
	case messageV1:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:103
		_go_fuzz_dep_.CoverTab[57581]++
															md = LegacyLoadMessageDesc(t)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:104
		// _ = "end of CoverTab[57581]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:105
	// _ = "end of CoverTab[57565]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:105
	_go_fuzz_dep_.CoverTab[57566]++

	// Derive basic field information from the struct tag.
	var evs protoreflect.EnumValueDescriptors
	if ed != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:109
		_go_fuzz_dep_.CoverTab[57582]++
															evs = ed.Values()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:110
		// _ = "end of CoverTab[57582]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:111
		_go_fuzz_dep_.CoverTab[57583]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:111
		// _ = "end of CoverTab[57583]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:111
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:111
	// _ = "end of CoverTab[57566]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:111
	_go_fuzz_dep_.CoverTab[57567]++
														fd := ptag.Unmarshal(xi.Tag, t, evs).(*filedesc.Field)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:115
	xd := &filedesc.Extension{L2: new(filedesc.ExtensionL2)}
														xd.L0.ParentFile = filedesc.SurrogateProto2
														xd.L0.FullName = protoreflect.FullName(xi.Name)
														xd.L1.Number = protoreflect.FieldNumber(xi.Field)
														xd.L1.Cardinality = fd.L1.Cardinality
														xd.L1.Kind = fd.L1.Kind
														xd.L2.IsPacked = fd.L1.IsPacked
														xd.L2.Default = fd.L1.Default
														xd.L1.Extendee = Export{}.MessageDescriptorOf(xi.ExtendedType)
														xd.L2.Enum = ed
														xd.L2.Message = md

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:128
	if messageset.IsMessageSet(xd.L1.Extendee) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:128
		_go_fuzz_dep_.CoverTab[57584]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:128
		return md.FullName() == xd.L0.FullName
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:128
		// _ = "end of CoverTab[57584]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:128
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:128
		_go_fuzz_dep_.CoverTab[57585]++
															xd.L0.FullName = xd.L0.FullName.Append(messageset.ExtensionName)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:129
		// _ = "end of CoverTab[57585]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:130
		_go_fuzz_dep_.CoverTab[57586]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:130
		// _ = "end of CoverTab[57586]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:130
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:130
	// _ = "end of CoverTab[57567]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:130
	_go_fuzz_dep_.CoverTab[57568]++

														tt := reflect.TypeOf(xi.ExtensionType)
														if isOptional {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:133
		_go_fuzz_dep_.CoverTab[57587]++
															tt = tt.Elem()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:134
		// _ = "end of CoverTab[57587]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:135
		_go_fuzz_dep_.CoverTab[57588]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:135
		// _ = "end of CoverTab[57588]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:135
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:135
	// _ = "end of CoverTab[57568]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:135
	_go_fuzz_dep_.CoverTab[57569]++
														xi.goType = tt
														xi.desc = extensionTypeDescriptor{xd, xi}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:137
	// _ = "end of CoverTab[57569]"
}

type placeholderExtension struct {
	name	protoreflect.FullName
	number	protoreflect.FieldNumber
}

func (x placeholderExtension) ParentFile() protoreflect.FileDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:145
	_go_fuzz_dep_.CoverTab[57589]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:145
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:145
	// _ = "end of CoverTab[57589]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:145
}
func (x placeholderExtension) Parent() protoreflect.Descriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:146
	_go_fuzz_dep_.CoverTab[57590]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:146
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:146
	// _ = "end of CoverTab[57590]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:146
}
func (x placeholderExtension) Index() int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:147
	_go_fuzz_dep_.CoverTab[57591]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:147
	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:147
	// _ = "end of CoverTab[57591]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:147
}
func (x placeholderExtension) Syntax() protoreflect.Syntax {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:148
	_go_fuzz_dep_.CoverTab[57592]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:148
	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:148
	// _ = "end of CoverTab[57592]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:148
}
func (x placeholderExtension) Name() protoreflect.Name {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:149
	_go_fuzz_dep_.CoverTab[57593]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:149
	return x.name.Name()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:149
	// _ = "end of CoverTab[57593]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:149
}
func (x placeholderExtension) FullName() protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:150
	_go_fuzz_dep_.CoverTab[57594]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:150
	return x.name
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:150
	// _ = "end of CoverTab[57594]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:150
}
func (x placeholderExtension) IsPlaceholder() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:151
	_go_fuzz_dep_.CoverTab[57595]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:151
	return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:151
	// _ = "end of CoverTab[57595]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:151
}
func (x placeholderExtension) Options() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:152
	_go_fuzz_dep_.CoverTab[57596]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:152
	return descopts.Field
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:152
	// _ = "end of CoverTab[57596]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:152
}
func (x placeholderExtension) Number() protoreflect.FieldNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:153
	_go_fuzz_dep_.CoverTab[57597]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:153
	return x.number
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:153
	// _ = "end of CoverTab[57597]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:153
}
func (x placeholderExtension) Cardinality() protoreflect.Cardinality {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:154
	_go_fuzz_dep_.CoverTab[57598]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:154
	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:154
	// _ = "end of CoverTab[57598]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:154
}
func (x placeholderExtension) Kind() protoreflect.Kind {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:155
	_go_fuzz_dep_.CoverTab[57599]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:155
	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:155
	// _ = "end of CoverTab[57599]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:155
}
func (x placeholderExtension) HasJSONName() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:156
	_go_fuzz_dep_.CoverTab[57600]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:156
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:156
	// _ = "end of CoverTab[57600]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:156
}
func (x placeholderExtension) JSONName() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:157
	_go_fuzz_dep_.CoverTab[57601]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:157
	return "[" + string(x.name) + "]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:157
	// _ = "end of CoverTab[57601]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:157
}
func (x placeholderExtension) TextName() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:158
	_go_fuzz_dep_.CoverTab[57602]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:158
	return "[" + string(x.name) + "]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:158
	// _ = "end of CoverTab[57602]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:158
}
func (x placeholderExtension) HasPresence() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:159
	_go_fuzz_dep_.CoverTab[57603]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:159
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:159
	// _ = "end of CoverTab[57603]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:159
}
func (x placeholderExtension) HasOptionalKeyword() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:160
	_go_fuzz_dep_.CoverTab[57604]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:160
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:160
	// _ = "end of CoverTab[57604]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:160
}
func (x placeholderExtension) IsExtension() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:161
	_go_fuzz_dep_.CoverTab[57605]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:161
	return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:161
	// _ = "end of CoverTab[57605]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:161
}
func (x placeholderExtension) IsWeak() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:162
	_go_fuzz_dep_.CoverTab[57606]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:162
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:162
	// _ = "end of CoverTab[57606]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:162
}
func (x placeholderExtension) IsPacked() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:163
	_go_fuzz_dep_.CoverTab[57607]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:163
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:163
	// _ = "end of CoverTab[57607]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:163
}
func (x placeholderExtension) IsList() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:164
	_go_fuzz_dep_.CoverTab[57608]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:164
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:164
	// _ = "end of CoverTab[57608]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:164
}
func (x placeholderExtension) IsMap() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:165
	_go_fuzz_dep_.CoverTab[57609]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:165
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:165
	// _ = "end of CoverTab[57609]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:165
}
func (x placeholderExtension) MapKey() protoreflect.FieldDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:166
	_go_fuzz_dep_.CoverTab[57610]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:166
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:166
	// _ = "end of CoverTab[57610]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:166
}
func (x placeholderExtension) MapValue() protoreflect.FieldDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:167
	_go_fuzz_dep_.CoverTab[57611]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:167
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:167
	// _ = "end of CoverTab[57611]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:167
}
func (x placeholderExtension) HasDefault() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:168
	_go_fuzz_dep_.CoverTab[57612]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:168
	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:168
	// _ = "end of CoverTab[57612]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:168
}
func (x placeholderExtension) Default() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:169
	_go_fuzz_dep_.CoverTab[57613]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:169
	return protoreflect.Value{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:169
	// _ = "end of CoverTab[57613]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:169
}
func (x placeholderExtension) DefaultEnumValue() protoreflect.EnumValueDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:170
	_go_fuzz_dep_.CoverTab[57614]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:170
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:170
	// _ = "end of CoverTab[57614]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:170
}
func (x placeholderExtension) ContainingOneof() protoreflect.OneofDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:171
	_go_fuzz_dep_.CoverTab[57615]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:171
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:171
	// _ = "end of CoverTab[57615]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:171
}
func (x placeholderExtension) ContainingMessage() protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:172
	_go_fuzz_dep_.CoverTab[57616]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:172
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:172
	// _ = "end of CoverTab[57616]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:172
}
func (x placeholderExtension) Enum() protoreflect.EnumDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:173
	_go_fuzz_dep_.CoverTab[57617]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:173
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:173
	// _ = "end of CoverTab[57617]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:173
}
func (x placeholderExtension) Message() protoreflect.MessageDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:174
	_go_fuzz_dep_.CoverTab[57618]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:174
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:174
	// _ = "end of CoverTab[57618]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:174
}
func (x placeholderExtension) ProtoType(protoreflect.FieldDescriptor) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:175
	_go_fuzz_dep_.CoverTab[57619]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:175
	return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:175
	// _ = "end of CoverTab[57619]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:175
}
func (x placeholderExtension) ProtoInternal(pragma.DoNotImplement) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:176
	_go_fuzz_dep_.CoverTab[57620]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:176
	return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:176
	// _ = "end of CoverTab[57620]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:176
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:176
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_extension.go:176
var _ = _go_fuzz_dep_.CoverTab
