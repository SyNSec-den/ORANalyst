// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:5
)

import (
	"reflect"
	"sync"
	"sync/atomic"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

// ExtensionInfo implements ExtensionType.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:16
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:16
// This type contains a number of exported fields for legacy compatibility.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:16
// The only non-deprecated use of this type is through the methods of the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:16
// ExtensionType interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:21
type ExtensionInfo struct {
	// An ExtensionInfo may exist in several stages of initialization.
	//
	// extensionInfoUninitialized: Some or all of the legacy exported
	// fields may be set, but none of the unexported fields have been
	// initialized. This is the starting state for an ExtensionInfo
	// in legacy generated code.
	//
	// extensionInfoDescInit: The desc field is set, but other unexported fields
	// may not be initialized. Legacy exported fields may or may not be set.
	// This is the starting state for an ExtensionInfo in newly generated code.
	//
	// extensionInfoFullInit: The ExtensionInfo is fully initialized.
	// This state is only entered after lazy initialization is complete.
	init	uint32
	mu	sync.Mutex

	goType	reflect.Type
	desc	extensionTypeDescriptor
	conv	Converter
	info	*extensionFieldInfo	// for fast-path method implementations

	// ExtendedType is a typed nil-pointer to the parent message type that
	// is being extended. It is possible for this to be unpopulated in v2
	// since the message may no longer implement the MessageV1 interface.
	//
	// Deprecated: Use the ExtendedType method instead.
	ExtendedType	protoiface.MessageV1

	// ExtensionType is the zero value of the extension type.
	//
	// For historical reasons, reflect.TypeOf(ExtensionType) and the
	// type returned by InterfaceOf may not be identical.
	//
	// Deprecated: Use InterfaceOf(xt.Zero()) instead.
	ExtensionType	interface{}

	// Field is the field number of the extension.
	//
	// Deprecated: Use the Descriptor().Number method instead.
	Field	int32

	// Name is the fully qualified name of extension.
	//
	// Deprecated: Use the Descriptor().FullName method instead.
	Name	string

	// Tag is the protobuf struct tag used in the v1 API.
	//
	// Deprecated: Do not use.
	Tag	string

	// Filename is the proto filename in which the extension is defined.
	//
	// Deprecated: Use Descriptor().ParentFile().Path() instead.
	Filename	string
}

// Stages of initialization: See the ExtensionInfo.init field.
const (
	extensionInfoUninitialized	= 0
	extensionInfoDescInit		= 1
	extensionInfoFullInit		= 2
)

func InitExtensionInfo(xi *ExtensionInfo, xd protoreflect.ExtensionDescriptor, goType reflect.Type) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:86
	_go_fuzz_dep_.CoverTab[57422]++
													xi.goType = goType
													xi.desc = extensionTypeDescriptor{xd, xi}
													xi.init = extensionInfoDescInit
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:89
	// _ = "end of CoverTab[57422]"
}

func (xi *ExtensionInfo) New() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:92
	_go_fuzz_dep_.CoverTab[57423]++
													return xi.lazyInit().New()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:93
	// _ = "end of CoverTab[57423]"
}
func (xi *ExtensionInfo) Zero() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:95
	_go_fuzz_dep_.CoverTab[57424]++
													return xi.lazyInit().Zero()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:96
	// _ = "end of CoverTab[57424]"
}
func (xi *ExtensionInfo) ValueOf(v interface{}) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:98
	_go_fuzz_dep_.CoverTab[57425]++
													return xi.lazyInit().PBValueOf(reflect.ValueOf(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:99
	// _ = "end of CoverTab[57425]"
}
func (xi *ExtensionInfo) InterfaceOf(v protoreflect.Value) interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:101
	_go_fuzz_dep_.CoverTab[57426]++
													return xi.lazyInit().GoValueOf(v).Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:102
	// _ = "end of CoverTab[57426]"
}
func (xi *ExtensionInfo) IsValidValue(v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:104
	_go_fuzz_dep_.CoverTab[57427]++
													return xi.lazyInit().IsValidPB(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:105
	// _ = "end of CoverTab[57427]"
}
func (xi *ExtensionInfo) IsValidInterface(v interface{}) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:107
	_go_fuzz_dep_.CoverTab[57428]++
													return xi.lazyInit().IsValidGo(reflect.ValueOf(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:108
	// _ = "end of CoverTab[57428]"
}
func (xi *ExtensionInfo) TypeDescriptor() protoreflect.ExtensionTypeDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:110
	_go_fuzz_dep_.CoverTab[57429]++
													if atomic.LoadUint32(&xi.init) < extensionInfoDescInit {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:111
		_go_fuzz_dep_.CoverTab[57431]++
														xi.lazyInitSlow()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:112
		// _ = "end of CoverTab[57431]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:113
		_go_fuzz_dep_.CoverTab[57432]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:113
		// _ = "end of CoverTab[57432]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:113
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:113
	// _ = "end of CoverTab[57429]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:113
	_go_fuzz_dep_.CoverTab[57430]++
													return &xi.desc
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:114
	// _ = "end of CoverTab[57430]"
}

func (xi *ExtensionInfo) lazyInit() Converter {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:117
	_go_fuzz_dep_.CoverTab[57433]++
													if atomic.LoadUint32(&xi.init) < extensionInfoFullInit {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:118
		_go_fuzz_dep_.CoverTab[57435]++
														xi.lazyInitSlow()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:119
		// _ = "end of CoverTab[57435]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:120
		_go_fuzz_dep_.CoverTab[57436]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:120
		// _ = "end of CoverTab[57436]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:120
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:120
	// _ = "end of CoverTab[57433]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:120
	_go_fuzz_dep_.CoverTab[57434]++
													return xi.conv
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:121
	// _ = "end of CoverTab[57434]"
}

func (xi *ExtensionInfo) lazyInitSlow() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:124
	_go_fuzz_dep_.CoverTab[57437]++
													xi.mu.Lock()
													defer xi.mu.Unlock()

													if xi.init == extensionInfoFullInit {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:128
		_go_fuzz_dep_.CoverTab[57440]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:129
		// _ = "end of CoverTab[57440]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:130
		_go_fuzz_dep_.CoverTab[57441]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:130
		// _ = "end of CoverTab[57441]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:130
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:130
	// _ = "end of CoverTab[57437]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:130
	_go_fuzz_dep_.CoverTab[57438]++
													defer atomic.StoreUint32(&xi.init, extensionInfoFullInit)

													if xi.desc.ExtensionDescriptor == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:133
		_go_fuzz_dep_.CoverTab[57442]++
														xi.initFromLegacy()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:134
		// _ = "end of CoverTab[57442]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:135
		_go_fuzz_dep_.CoverTab[57443]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:135
		// _ = "end of CoverTab[57443]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:135
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:135
	// _ = "end of CoverTab[57438]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:135
	_go_fuzz_dep_.CoverTab[57439]++
													if !xi.desc.ExtensionDescriptor.IsPlaceholder() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:136
		_go_fuzz_dep_.CoverTab[57444]++
														if xi.ExtensionType == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:137
			_go_fuzz_dep_.CoverTab[57446]++
															xi.initToLegacy()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:138
			// _ = "end of CoverTab[57446]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:139
			_go_fuzz_dep_.CoverTab[57447]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:139
			// _ = "end of CoverTab[57447]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:139
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:139
		// _ = "end of CoverTab[57444]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:139
		_go_fuzz_dep_.CoverTab[57445]++
														xi.conv = NewConverter(xi.goType, xi.desc.ExtensionDescriptor)
														xi.info = makeExtensionFieldInfo(xi.desc.ExtensionDescriptor)
														xi.info.validation = newValidationInfo(xi.desc.ExtensionDescriptor, xi.goType)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:142
		// _ = "end of CoverTab[57445]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:143
		_go_fuzz_dep_.CoverTab[57448]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:143
		// _ = "end of CoverTab[57448]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:143
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:143
	// _ = "end of CoverTab[57439]"
}

type extensionTypeDescriptor struct {
	protoreflect.ExtensionDescriptor
	xi	*ExtensionInfo
}

func (xtd *extensionTypeDescriptor) Type() protoreflect.ExtensionType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:151
	_go_fuzz_dep_.CoverTab[57449]++
													return xtd.xi
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:152
	// _ = "end of CoverTab[57449]"
}
func (xtd *extensionTypeDescriptor) Descriptor() protoreflect.ExtensionDescriptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:154
	_go_fuzz_dep_.CoverTab[57450]++
													return xtd.ExtensionDescriptor
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:155
	// _ = "end of CoverTab[57450]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:156
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/extension.go:156
var _ = _go_fuzz_dep_.CoverTab
