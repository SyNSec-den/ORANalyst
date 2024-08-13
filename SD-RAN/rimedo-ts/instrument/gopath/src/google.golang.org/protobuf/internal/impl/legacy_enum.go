// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:5
)

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"google.golang.org/protobuf/internal/filedesc"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// legacyEnumName returns the name of enums used in legacy code.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:18
// It is neither the protobuf full name nor the qualified Go name,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:18
// but rather an odd hybrid of both.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:21
func legacyEnumName(ed protoreflect.EnumDescriptor) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:21
	_go_fuzz_dep_.CoverTab[57451]++
													var protoPkg string
													enumName := string(ed.FullName())
													if fd := ed.ParentFile(); fd != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:24
		_go_fuzz_dep_.CoverTab[57454]++
														protoPkg = string(fd.Package())
														enumName = strings.TrimPrefix(enumName, protoPkg+".")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:26
		// _ = "end of CoverTab[57454]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:27
		_go_fuzz_dep_.CoverTab[57455]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:27
		// _ = "end of CoverTab[57455]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:27
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:27
	// _ = "end of CoverTab[57451]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:27
	_go_fuzz_dep_.CoverTab[57452]++
													if protoPkg == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:28
		_go_fuzz_dep_.CoverTab[57456]++
														return strs.GoCamelCase(enumName)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:29
		// _ = "end of CoverTab[57456]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:30
		_go_fuzz_dep_.CoverTab[57457]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:30
		// _ = "end of CoverTab[57457]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:30
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:30
	// _ = "end of CoverTab[57452]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:30
	_go_fuzz_dep_.CoverTab[57453]++
													return protoPkg + "." + strs.GoCamelCase(enumName)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:31
	// _ = "end of CoverTab[57453]"
}

// legacyWrapEnum wraps v as a protoreflect.Enum,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:34
// where v must be a int32 kind and not implement the v2 API already.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:36
func legacyWrapEnum(v reflect.Value) protoreflect.Enum {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:36
	_go_fuzz_dep_.CoverTab[57458]++
													et := legacyLoadEnumType(v.Type())
													return et.New(protoreflect.EnumNumber(v.Int()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:38
	// _ = "end of CoverTab[57458]"
}

var legacyEnumTypeCache sync.Map	// map[reflect.Type]protoreflect.EnumType

// legacyLoadEnumType dynamically loads a protoreflect.EnumType for t,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:43
// where t must be an int32 kind and not implement the v2 API already.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:45
func legacyLoadEnumType(t reflect.Type) protoreflect.EnumType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:45
	_go_fuzz_dep_.CoverTab[57459]++

													if et, ok := legacyEnumTypeCache.Load(t); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:47
		_go_fuzz_dep_.CoverTab[57462]++
														return et.(protoreflect.EnumType)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:48
		// _ = "end of CoverTab[57462]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:49
		_go_fuzz_dep_.CoverTab[57463]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:49
		// _ = "end of CoverTab[57463]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:49
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:49
	// _ = "end of CoverTab[57459]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:49
	_go_fuzz_dep_.CoverTab[57460]++

	// Slow-path: derive enum descriptor and initialize EnumType.
	var et protoreflect.EnumType
	ed := LegacyLoadEnumDesc(t)
	et = &legacyEnumType{
		desc:	ed,
		goType:	t,
	}
	if et, ok := legacyEnumTypeCache.LoadOrStore(t, et); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:58
		_go_fuzz_dep_.CoverTab[57464]++
														return et.(protoreflect.EnumType)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:59
		// _ = "end of CoverTab[57464]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:60
		_go_fuzz_dep_.CoverTab[57465]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:60
		// _ = "end of CoverTab[57465]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:60
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:60
	// _ = "end of CoverTab[57460]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:60
	_go_fuzz_dep_.CoverTab[57461]++
													return et
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:61
	// _ = "end of CoverTab[57461]"
}

type legacyEnumType struct {
	desc	protoreflect.EnumDescriptor
	goType	reflect.Type
	m	sync.Map	// map[protoreflect.EnumNumber]proto.Enum
}

func (t *legacyEnumType) New(n protoreflect.EnumNumber) protoreflect.Enum {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:70
	_go_fuzz_dep_.CoverTab[57466]++
													if e, ok := t.m.Load(n); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:71
		_go_fuzz_dep_.CoverTab[57468]++
														return e.(protoreflect.Enum)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:72
		// _ = "end of CoverTab[57468]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:73
		_go_fuzz_dep_.CoverTab[57469]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:73
		// _ = "end of CoverTab[57469]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:73
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:73
	// _ = "end of CoverTab[57466]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:73
	_go_fuzz_dep_.CoverTab[57467]++
													e := &legacyEnumWrapper{num: n, pbTyp: t, goTyp: t.goType}
													t.m.Store(n, e)
													return e
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:76
	// _ = "end of CoverTab[57467]"
}
func (t *legacyEnumType) Descriptor() protoreflect.EnumDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:78
	_go_fuzz_dep_.CoverTab[57470]++
													return t.desc
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:79
	// _ = "end of CoverTab[57470]"
}

type legacyEnumWrapper struct {
	num	protoreflect.EnumNumber
	pbTyp	protoreflect.EnumType
	goTyp	reflect.Type
}

func (e *legacyEnumWrapper) Descriptor() protoreflect.EnumDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:88
	_go_fuzz_dep_.CoverTab[57471]++
													return e.pbTyp.Descriptor()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:89
	// _ = "end of CoverTab[57471]"
}
func (e *legacyEnumWrapper) Type() protoreflect.EnumType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:91
	_go_fuzz_dep_.CoverTab[57472]++
													return e.pbTyp
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:92
	// _ = "end of CoverTab[57472]"
}
func (e *legacyEnumWrapper) Number() protoreflect.EnumNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:94
	_go_fuzz_dep_.CoverTab[57473]++
													return e.num
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:95
	// _ = "end of CoverTab[57473]"
}
func (e *legacyEnumWrapper) ProtoReflect() protoreflect.Enum {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:97
	_go_fuzz_dep_.CoverTab[57474]++
													return e
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:98
	// _ = "end of CoverTab[57474]"
}
func (e *legacyEnumWrapper) protoUnwrap() interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:100
	_go_fuzz_dep_.CoverTab[57475]++
													v := reflect.New(e.goTyp).Elem()
													v.SetInt(int64(e.num))
													return v.Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:103
	// _ = "end of CoverTab[57475]"
}

var (
	_	protoreflect.Enum	= (*legacyEnumWrapper)(nil)
	_	unwrapper		= (*legacyEnumWrapper)(nil)
)

var legacyEnumDescCache sync.Map	// map[reflect.Type]protoreflect.EnumDescriptor

// LegacyLoadEnumDesc returns an EnumDescriptor derived from the Go type,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:113
// which must be an int32 kind and not implement the v2 API already.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:113
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:113
// This is exported for testing purposes.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:117
func LegacyLoadEnumDesc(t reflect.Type) protoreflect.EnumDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:117
	_go_fuzz_dep_.CoverTab[57476]++

													if ed, ok := legacyEnumDescCache.Load(t); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:119
		_go_fuzz_dep_.CoverTab[57482]++
														return ed.(protoreflect.EnumDescriptor)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:120
		// _ = "end of CoverTab[57482]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:121
		_go_fuzz_dep_.CoverTab[57483]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:121
		// _ = "end of CoverTab[57483]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:121
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:121
	// _ = "end of CoverTab[57476]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:121
	_go_fuzz_dep_.CoverTab[57477]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:124
	ev := reflect.Zero(t).Interface()
	if _, ok := ev.(protoreflect.Enum); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:125
		_go_fuzz_dep_.CoverTab[57484]++
														panic(fmt.Sprintf("%v already implements proto.Enum", t))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:126
		// _ = "end of CoverTab[57484]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:127
		_go_fuzz_dep_.CoverTab[57485]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:127
		// _ = "end of CoverTab[57485]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:127
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:127
	// _ = "end of CoverTab[57477]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:127
	_go_fuzz_dep_.CoverTab[57478]++
													edV1, ok := ev.(enumV1)
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:129
		_go_fuzz_dep_.CoverTab[57486]++
														return aberrantLoadEnumDesc(t)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:130
		// _ = "end of CoverTab[57486]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:131
		_go_fuzz_dep_.CoverTab[57487]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:131
		// _ = "end of CoverTab[57487]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:131
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:131
	// _ = "end of CoverTab[57478]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:131
	_go_fuzz_dep_.CoverTab[57479]++
													b, idxs := edV1.EnumDescriptor()

													var ed protoreflect.EnumDescriptor
													if len(idxs) == 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:135
		_go_fuzz_dep_.CoverTab[57488]++
														ed = legacyLoadFileDesc(b).Enums().Get(idxs[0])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:136
		// _ = "end of CoverTab[57488]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:137
		_go_fuzz_dep_.CoverTab[57489]++
														md := legacyLoadFileDesc(b).Messages().Get(idxs[0])
														for _, i := range idxs[1 : len(idxs)-1] {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:139
			_go_fuzz_dep_.CoverTab[57491]++
															md = md.Messages().Get(i)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:140
			// _ = "end of CoverTab[57491]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:141
		// _ = "end of CoverTab[57489]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:141
		_go_fuzz_dep_.CoverTab[57490]++
														ed = md.Enums().Get(idxs[len(idxs)-1])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:142
		// _ = "end of CoverTab[57490]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:143
	// _ = "end of CoverTab[57479]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:143
	_go_fuzz_dep_.CoverTab[57480]++
													if ed, ok := legacyEnumDescCache.LoadOrStore(t, ed); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:144
		_go_fuzz_dep_.CoverTab[57492]++
														return ed.(protoreflect.EnumDescriptor)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:145
		// _ = "end of CoverTab[57492]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:146
		_go_fuzz_dep_.CoverTab[57493]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:146
		// _ = "end of CoverTab[57493]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:146
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:146
	// _ = "end of CoverTab[57480]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:146
	_go_fuzz_dep_.CoverTab[57481]++
													return ed
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:147
	// _ = "end of CoverTab[57481]"
}

var aberrantEnumDescCache sync.Map	// map[reflect.Type]protoreflect.EnumDescriptor

// aberrantLoadEnumDesc returns an EnumDescriptor derived from the Go type,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:152
// which must not implement protoreflect.Enum or enumV1.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:152
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:152
// If the type does not implement enumV1, then there is no reliable
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:152
// way to derive the original protobuf type information.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:152
// We are unable to use the global enum registry since it is
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:152
// unfortunately keyed by the protobuf full name, which we also do not know.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:152
// Thus, this produces some bogus enum descriptor based on the Go type name.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:160
func aberrantLoadEnumDesc(t reflect.Type) protoreflect.EnumDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:160
	_go_fuzz_dep_.CoverTab[57494]++

													if ed, ok := aberrantEnumDescCache.Load(t); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:162
		_go_fuzz_dep_.CoverTab[57497]++
														return ed.(protoreflect.EnumDescriptor)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:163
		// _ = "end of CoverTab[57497]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:164
		_go_fuzz_dep_.CoverTab[57498]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:164
		// _ = "end of CoverTab[57498]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:164
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:164
	// _ = "end of CoverTab[57494]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:164
	_go_fuzz_dep_.CoverTab[57495]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:167
	ed := &filedesc.Enum{L2: new(filedesc.EnumL2)}
													ed.L0.FullName = AberrantDeriveFullName(t)
													ed.L0.ParentFile = filedesc.SurrogateProto3
													ed.L2.Values.List = append(ed.L2.Values.List, filedesc.EnumValue{})

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:174
	vd := &ed.L2.Values.List[0]
													vd.L0.FullName = ed.L0.FullName + "_UNKNOWN"
													vd.L0.ParentFile = ed.L0.ParentFile
													vd.L0.Parent = ed

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:183
	if ed, ok := aberrantEnumDescCache.LoadOrStore(t, ed); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:183
		_go_fuzz_dep_.CoverTab[57499]++
														return ed.(protoreflect.EnumDescriptor)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:184
		// _ = "end of CoverTab[57499]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:185
		_go_fuzz_dep_.CoverTab[57500]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:185
		// _ = "end of CoverTab[57500]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:185
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:185
	// _ = "end of CoverTab[57495]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:185
	_go_fuzz_dep_.CoverTab[57496]++
													return ed
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:186
	// _ = "end of CoverTab[57496]"
}

// AberrantDeriveFullName derives a fully qualified protobuf name for the given Go type
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:189
// The provided name is not guaranteed to be stable nor universally unique.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:189
// It should be sufficiently unique within a program.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:189
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:189
// This is exported for testing purposes.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:194
func AberrantDeriveFullName(t reflect.Type) protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:194
	_go_fuzz_dep_.CoverTab[57501]++
													sanitize := func(r rune) rune {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:195
		_go_fuzz_dep_.CoverTab[57505]++
														switch {
		case r == '/':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:197
			_go_fuzz_dep_.CoverTab[57506]++
															return '.'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:198
			// _ = "end of CoverTab[57506]"
		case 'a' <= r && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:199
			_go_fuzz_dep_.CoverTab[57509]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:199
			return r <= 'z'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:199
			// _ = "end of CoverTab[57509]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:199
		}(), 'A' <= r && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:199
			_go_fuzz_dep_.CoverTab[57510]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:199
			return r <= 'Z'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:199
			// _ = "end of CoverTab[57510]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:199
		}(), '0' <= r && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:199
			_go_fuzz_dep_.CoverTab[57511]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:199
			return r <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:199
			// _ = "end of CoverTab[57511]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:199
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:199
			_go_fuzz_dep_.CoverTab[57507]++
															return r
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:200
			// _ = "end of CoverTab[57507]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:201
			_go_fuzz_dep_.CoverTab[57508]++
															return '_'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:202
			// _ = "end of CoverTab[57508]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:203
		// _ = "end of CoverTab[57505]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:204
	// _ = "end of CoverTab[57501]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:204
	_go_fuzz_dep_.CoverTab[57502]++
													prefix := strings.Map(sanitize, t.PkgPath())
													suffix := strings.Map(sanitize, t.Name())
													if suffix == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:207
		_go_fuzz_dep_.CoverTab[57512]++
														suffix = fmt.Sprintf("UnknownX%X", reflect.ValueOf(t).Pointer())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:208
		// _ = "end of CoverTab[57512]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:209
		_go_fuzz_dep_.CoverTab[57513]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:209
		// _ = "end of CoverTab[57513]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:209
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:209
	// _ = "end of CoverTab[57502]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:209
	_go_fuzz_dep_.CoverTab[57503]++

													ss := append(strings.Split(prefix, "."), suffix)
													for i, s := range ss {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:212
		_go_fuzz_dep_.CoverTab[57514]++
														if s == "" || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:213
			_go_fuzz_dep_.CoverTab[57515]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:213
			return ('0' <= s[0] && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:213
				_go_fuzz_dep_.CoverTab[57516]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:213
				return s[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:213
				// _ = "end of CoverTab[57516]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:213
			}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:213
			// _ = "end of CoverTab[57515]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:213
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:213
			_go_fuzz_dep_.CoverTab[57517]++
															ss[i] = "x" + s
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:214
			// _ = "end of CoverTab[57517]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:215
			_go_fuzz_dep_.CoverTab[57518]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:215
			// _ = "end of CoverTab[57518]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:215
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:215
		// _ = "end of CoverTab[57514]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:216
	// _ = "end of CoverTab[57503]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:216
	_go_fuzz_dep_.CoverTab[57504]++
													return protoreflect.FullName(strings.Join(ss, "."))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:217
	// _ = "end of CoverTab[57504]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:218
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/legacy_enum.go:218
var _ = _go_fuzz_dep_.CoverTab
