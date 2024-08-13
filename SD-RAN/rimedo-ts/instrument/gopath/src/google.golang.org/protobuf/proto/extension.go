// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:5
package proto

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:5
)

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

// HasExtension reports whether an extension field is populated.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:11
// It returns false if m is invalid or if xt does not extend m.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:13
func HasExtension(m Message, xt protoreflect.ExtensionType) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:13
	_go_fuzz_dep_.CoverTab[51406]++

												if m == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:15
		_go_fuzz_dep_.CoverTab[51409]++
													return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:16
		// _ = "end of CoverTab[51409]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:17
		_go_fuzz_dep_.CoverTab[51410]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:17
		// _ = "end of CoverTab[51410]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:17
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:17
	// _ = "end of CoverTab[51406]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:17
	_go_fuzz_dep_.CoverTab[51407]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:21
	if xt == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:21
		_go_fuzz_dep_.CoverTab[51411]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:21
		return m.ProtoReflect().Descriptor() != xt.TypeDescriptor().ContainingMessage()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:21
		// _ = "end of CoverTab[51411]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:21
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:21
		_go_fuzz_dep_.CoverTab[51412]++
													return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:22
		// _ = "end of CoverTab[51412]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:23
		_go_fuzz_dep_.CoverTab[51413]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:23
		// _ = "end of CoverTab[51413]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:23
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:23
	// _ = "end of CoverTab[51407]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:23
	_go_fuzz_dep_.CoverTab[51408]++

												return m.ProtoReflect().Has(xt.TypeDescriptor())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:25
	// _ = "end of CoverTab[51408]"
}

// ClearExtension clears an extension field such that subsequent
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:28
// HasExtension calls return false.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:28
// It panics if m is invalid or if xt does not extend m.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:31
func ClearExtension(m Message, xt protoreflect.ExtensionType) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:31
	_go_fuzz_dep_.CoverTab[51414]++
												m.ProtoReflect().Clear(xt.TypeDescriptor())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:32
	// _ = "end of CoverTab[51414]"
}

// GetExtension retrieves the value for an extension field.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:35
// If the field is unpopulated, it returns the default value for
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:35
// scalars and an immutable, empty value for lists or messages.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:35
// It panics if xt does not extend m.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:39
func GetExtension(m Message, xt protoreflect.ExtensionType) interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:39
	_go_fuzz_dep_.CoverTab[51415]++

												if m == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:41
		_go_fuzz_dep_.CoverTab[51417]++
													return xt.InterfaceOf(xt.Zero())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:42
		// _ = "end of CoverTab[51417]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:43
		_go_fuzz_dep_.CoverTab[51418]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:43
		// _ = "end of CoverTab[51418]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:43
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:43
	// _ = "end of CoverTab[51415]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:43
	_go_fuzz_dep_.CoverTab[51416]++

												return xt.InterfaceOf(m.ProtoReflect().Get(xt.TypeDescriptor()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:45
	// _ = "end of CoverTab[51416]"
}

// SetExtension stores the value of an extension field.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:48
// It panics if m is invalid, xt does not extend m, or if type of v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:48
// is invalid for the specified extension field.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:51
func SetExtension(m Message, xt protoreflect.ExtensionType, v interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:51
	_go_fuzz_dep_.CoverTab[51419]++
												xd := xt.TypeDescriptor()
												pv := xt.ValueOf(v)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:56
	isValid := true
	switch {
	case xd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:58
		_go_fuzz_dep_.CoverTab[51422]++
													isValid = pv.List().IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:59
		// _ = "end of CoverTab[51422]"
	case xd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:60
		_go_fuzz_dep_.CoverTab[51423]++
													isValid = pv.Map().IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:61
		// _ = "end of CoverTab[51423]"
	case xd.Message() != nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:62
		_go_fuzz_dep_.CoverTab[51424]++
													isValid = pv.Message().IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:63
		// _ = "end of CoverTab[51424]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:63
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:63
		_go_fuzz_dep_.CoverTab[51425]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:63
		// _ = "end of CoverTab[51425]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:64
	// _ = "end of CoverTab[51419]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:64
	_go_fuzz_dep_.CoverTab[51420]++
												if !isValid {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:65
		_go_fuzz_dep_.CoverTab[51426]++
													m.ProtoReflect().Clear(xd)
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:67
		// _ = "end of CoverTab[51426]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:68
		_go_fuzz_dep_.CoverTab[51427]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:68
		// _ = "end of CoverTab[51427]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:68
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:68
	// _ = "end of CoverTab[51420]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:68
	_go_fuzz_dep_.CoverTab[51421]++

												m.ProtoReflect().Set(xd, pv)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:70
	// _ = "end of CoverTab[51421]"
}

// RangeExtensions iterates over every populated extension field in m in an
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:73
// undefined order, calling f for each extension type and value encountered.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:73
// It returns immediately if f returns false.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:73
// While iterating, mutating operations may only be performed
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:73
// on the current extension field.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:78
func RangeExtensions(m Message, f func(protoreflect.ExtensionType, interface{}) bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:78
	_go_fuzz_dep_.CoverTab[51428]++

												if m == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:80
		_go_fuzz_dep_.CoverTab[51430]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:81
		// _ = "end of CoverTab[51430]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:82
		_go_fuzz_dep_.CoverTab[51431]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:82
		// _ = "end of CoverTab[51431]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:82
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:82
	// _ = "end of CoverTab[51428]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:82
	_go_fuzz_dep_.CoverTab[51429]++

												m.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:84
		_go_fuzz_dep_.CoverTab[51432]++
													if fd.IsExtension() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:85
			_go_fuzz_dep_.CoverTab[51434]++
														xt := fd.(protoreflect.ExtensionTypeDescriptor).Type()
														vi := xt.InterfaceOf(v)
														return f(xt, vi)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:88
			// _ = "end of CoverTab[51434]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:89
			_go_fuzz_dep_.CoverTab[51435]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:89
			// _ = "end of CoverTab[51435]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:89
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:89
		// _ = "end of CoverTab[51432]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:89
		_go_fuzz_dep_.CoverTab[51433]++
													return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:90
		// _ = "end of CoverTab[51433]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:91
	// _ = "end of CoverTab[51429]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:92
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/extension.go:92
var _ = _go_fuzz_dep_.CoverTab
