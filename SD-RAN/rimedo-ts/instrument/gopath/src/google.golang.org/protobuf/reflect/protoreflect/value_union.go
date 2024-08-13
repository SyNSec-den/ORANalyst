// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:5
package protoreflect

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:5
)

import (
	"fmt"
	"math"
)

// Value is a union where only one Go type may be set at a time.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// The Value is used to represent all possible values a field may take.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// The following shows which Go type is used to represent each proto Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	╔════════════╤═════════════════════════════════════╗
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	║ Go type    │ Protobuf kind                       ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	╠════════════╪═════════════════════════════════════╣
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	║ bool       │ BoolKind                            ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	║ int32      │ Int32Kind, Sint32Kind, Sfixed32Kind ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	║ int64      │ Int64Kind, Sint64Kind, Sfixed64Kind ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	║ uint32     │ Uint32Kind, Fixed32Kind             ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	║ uint64     │ Uint64Kind, Fixed64Kind             ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	║ float32    │ FloatKind                           ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	║ float64    │ DoubleKind                          ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	║ string     │ StringKind                          ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	║ []byte     │ BytesKind                           ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	║ EnumNumber │ EnumKind                            ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	║ Message    │ MessageKind, GroupKind              ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	╚════════════╧═════════════════════════════════════╝
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// Multiple protobuf Kinds may be represented by a single Go type if the type
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// can losslessly represent the information for the proto kind. For example,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// Int64Kind, Sint64Kind, and Sfixed64Kind are all represented by int64,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// but use different integer encoding methods.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// The List or Map types are used if the field cardinality is repeated.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// A field is a List if FieldDescriptor.IsList reports true.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// A field is a Map if FieldDescriptor.IsMap reports true.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// Converting to/from a Value and a concrete Go value panics on type mismatch.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// For example, ValueOf("hello").Int() panics because this attempts to
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// retrieve an int64 from a string.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// List, Map, and Message Values are called "composite" values.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// A composite Value may alias (reference) memory at some location,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// such that changes to the Value updates the that location.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// A composite value acquired with a Mutable method, such as Message.Mutable,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// always references the source object.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// For example:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	// Append a 0 to a "repeated int32" field.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	// Since the Value returned by Mutable is guaranteed to alias
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	// the source message, modifying the Value modifies the message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	message.Mutable(fieldDesc).(List).Append(protoreflect.ValueOfInt32(0))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	// Assign [0] to a "repeated int32" field by creating a new Value,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	// modifying it, and assigning it.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	list := message.NewField(fieldDesc).(List)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	list.Append(protoreflect.ValueOfInt32(0))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	message.Set(fieldDesc, list)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	// ERROR: Since it is not defined whether Set aliases the source,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	// appending to the List here may or may not modify the message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//	list.Append(protoreflect.ValueOfInt32(0))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// Some operations, such as Message.Get, may return an "empty, read-only"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:12
// composite Value. Modifying an empty, read-only value panics.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:70
type Value value

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:85
// ValueOf returns a Value initialized with the concrete value stored in v.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:85
// This panics if the type does not match one of the allowed types in the
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:85
// Value union.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:88
func ValueOf(v interface{}) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:88
	_go_fuzz_dep_.CoverTab[48903]++
														switch v := v.(type) {
	case nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:90
		_go_fuzz_dep_.CoverTab[48904]++
															return Value{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:91
		// _ = "end of CoverTab[48904]"
	case bool:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:92
		_go_fuzz_dep_.CoverTab[48905]++
															return ValueOfBool(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:93
		// _ = "end of CoverTab[48905]"
	case int32:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:94
		_go_fuzz_dep_.CoverTab[48906]++
															return ValueOfInt32(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:95
		// _ = "end of CoverTab[48906]"
	case int64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:96
		_go_fuzz_dep_.CoverTab[48907]++
															return ValueOfInt64(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:97
		// _ = "end of CoverTab[48907]"
	case uint32:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:98
		_go_fuzz_dep_.CoverTab[48908]++
															return ValueOfUint32(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:99
		// _ = "end of CoverTab[48908]"
	case uint64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:100
		_go_fuzz_dep_.CoverTab[48909]++
															return ValueOfUint64(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:101
		// _ = "end of CoverTab[48909]"
	case float32:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:102
		_go_fuzz_dep_.CoverTab[48910]++
															return ValueOfFloat32(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:103
		// _ = "end of CoverTab[48910]"
	case float64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:104
		_go_fuzz_dep_.CoverTab[48911]++
															return ValueOfFloat64(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:105
		// _ = "end of CoverTab[48911]"
	case string:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:106
		_go_fuzz_dep_.CoverTab[48912]++
															return ValueOfString(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:107
		// _ = "end of CoverTab[48912]"
	case []byte:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:108
		_go_fuzz_dep_.CoverTab[48913]++
															return ValueOfBytes(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:109
		// _ = "end of CoverTab[48913]"
	case EnumNumber:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:110
		_go_fuzz_dep_.CoverTab[48914]++
															return ValueOfEnum(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:111
		// _ = "end of CoverTab[48914]"
	case Message, List, Map:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:112
		_go_fuzz_dep_.CoverTab[48915]++
															return valueOfIface(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:113
		// _ = "end of CoverTab[48915]"
	case ProtoMessage:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:114
		_go_fuzz_dep_.CoverTab[48916]++
															panic(fmt.Sprintf("invalid proto.Message(%T) type, expected a protoreflect.Message type", v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:115
		// _ = "end of CoverTab[48916]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:116
		_go_fuzz_dep_.CoverTab[48917]++
															panic(fmt.Sprintf("invalid type: %T", v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:117
		// _ = "end of CoverTab[48917]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:118
	// _ = "end of CoverTab[48903]"
}

// ValueOfBool returns a new boolean value.
func ValueOfBool(v bool) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:122
	_go_fuzz_dep_.CoverTab[48918]++
														if v {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:123
		_go_fuzz_dep_.CoverTab[48919]++
															return Value{typ: boolType, num: 1}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:124
		// _ = "end of CoverTab[48919]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:125
		_go_fuzz_dep_.CoverTab[48920]++
															return Value{typ: boolType, num: 0}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:126
		// _ = "end of CoverTab[48920]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:127
	// _ = "end of CoverTab[48918]"
}

// ValueOfInt32 returns a new int32 value.
func ValueOfInt32(v int32) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:131
	_go_fuzz_dep_.CoverTab[48921]++
														return Value{typ: int32Type, num: uint64(v)}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:132
	// _ = "end of CoverTab[48921]"
}

// ValueOfInt64 returns a new int64 value.
func ValueOfInt64(v int64) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:136
	_go_fuzz_dep_.CoverTab[48922]++
														return Value{typ: int64Type, num: uint64(v)}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:137
	// _ = "end of CoverTab[48922]"
}

// ValueOfUint32 returns a new uint32 value.
func ValueOfUint32(v uint32) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:141
	_go_fuzz_dep_.CoverTab[48923]++
														return Value{typ: uint32Type, num: uint64(v)}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:142
	// _ = "end of CoverTab[48923]"
}

// ValueOfUint64 returns a new uint64 value.
func ValueOfUint64(v uint64) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:146
	_go_fuzz_dep_.CoverTab[48924]++
														return Value{typ: uint64Type, num: v}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:147
	// _ = "end of CoverTab[48924]"
}

// ValueOfFloat32 returns a new float32 value.
func ValueOfFloat32(v float32) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:151
	_go_fuzz_dep_.CoverTab[48925]++
														return Value{typ: float32Type, num: uint64(math.Float64bits(float64(v)))}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:152
	// _ = "end of CoverTab[48925]"
}

// ValueOfFloat64 returns a new float64 value.
func ValueOfFloat64(v float64) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:156
	_go_fuzz_dep_.CoverTab[48926]++
														return Value{typ: float64Type, num: uint64(math.Float64bits(float64(v)))}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:157
	// _ = "end of CoverTab[48926]"
}

// ValueOfString returns a new string value.
func ValueOfString(v string) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:161
	_go_fuzz_dep_.CoverTab[48927]++
														return valueOfString(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:162
	// _ = "end of CoverTab[48927]"
}

// ValueOfBytes returns a new bytes value.
func ValueOfBytes(v []byte) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:166
	_go_fuzz_dep_.CoverTab[48928]++
														return valueOfBytes(v[:len(v):len(v)])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:167
	// _ = "end of CoverTab[48928]"
}

// ValueOfEnum returns a new enum value.
func ValueOfEnum(v EnumNumber) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:171
	_go_fuzz_dep_.CoverTab[48929]++
														return Value{typ: enumType, num: uint64(v)}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:172
	// _ = "end of CoverTab[48929]"
}

// ValueOfMessage returns a new Message value.
func ValueOfMessage(v Message) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:176
	_go_fuzz_dep_.CoverTab[48930]++
														return valueOfIface(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:177
	// _ = "end of CoverTab[48930]"
}

// ValueOfList returns a new List value.
func ValueOfList(v List) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:181
	_go_fuzz_dep_.CoverTab[48931]++
														return valueOfIface(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:182
	// _ = "end of CoverTab[48931]"
}

// ValueOfMap returns a new Map value.
func ValueOfMap(v Map) Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:186
	_go_fuzz_dep_.CoverTab[48932]++
														return valueOfIface(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:187
	// _ = "end of CoverTab[48932]"
}

// IsValid reports whether v is populated with a value.
func (v Value) IsValid() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:191
	_go_fuzz_dep_.CoverTab[48933]++
														return v.typ != nilType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:192
	// _ = "end of CoverTab[48933]"
}

// Interface returns v as an interface{}.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:195
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:195
// Invariant: v == ValueOf(v).Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:198
func (v Value) Interface() interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:198
	_go_fuzz_dep_.CoverTab[48934]++
														switch v.typ {
	case nilType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:200
		_go_fuzz_dep_.CoverTab[48935]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:201
		// _ = "end of CoverTab[48935]"
	case boolType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:202
		_go_fuzz_dep_.CoverTab[48936]++
															return v.Bool()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:203
		// _ = "end of CoverTab[48936]"
	case int32Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:204
		_go_fuzz_dep_.CoverTab[48937]++
															return int32(v.Int())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:205
		// _ = "end of CoverTab[48937]"
	case int64Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:206
		_go_fuzz_dep_.CoverTab[48938]++
															return int64(v.Int())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:207
		// _ = "end of CoverTab[48938]"
	case uint32Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:208
		_go_fuzz_dep_.CoverTab[48939]++
															return uint32(v.Uint())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:209
		// _ = "end of CoverTab[48939]"
	case uint64Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:210
		_go_fuzz_dep_.CoverTab[48940]++
															return uint64(v.Uint())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:211
		// _ = "end of CoverTab[48940]"
	case float32Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:212
		_go_fuzz_dep_.CoverTab[48941]++
															return float32(v.Float())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:213
		// _ = "end of CoverTab[48941]"
	case float64Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:214
		_go_fuzz_dep_.CoverTab[48942]++
															return float64(v.Float())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:215
		// _ = "end of CoverTab[48942]"
	case stringType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:216
		_go_fuzz_dep_.CoverTab[48943]++
															return v.String()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:217
		// _ = "end of CoverTab[48943]"
	case bytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:218
		_go_fuzz_dep_.CoverTab[48944]++
															return v.Bytes()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:219
		// _ = "end of CoverTab[48944]"
	case enumType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:220
		_go_fuzz_dep_.CoverTab[48945]++
															return v.Enum()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:221
		// _ = "end of CoverTab[48945]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:222
		_go_fuzz_dep_.CoverTab[48946]++
															return v.getIface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:223
		// _ = "end of CoverTab[48946]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:224
	// _ = "end of CoverTab[48934]"
}

func (v Value) typeName() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:227
	_go_fuzz_dep_.CoverTab[48947]++
														switch v.typ {
	case nilType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:229
		_go_fuzz_dep_.CoverTab[48948]++
															return "nil"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:230
		// _ = "end of CoverTab[48948]"
	case boolType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:231
		_go_fuzz_dep_.CoverTab[48949]++
															return "bool"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:232
		// _ = "end of CoverTab[48949]"
	case int32Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:233
		_go_fuzz_dep_.CoverTab[48950]++
															return "int32"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:234
		// _ = "end of CoverTab[48950]"
	case int64Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:235
		_go_fuzz_dep_.CoverTab[48951]++
															return "int64"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:236
		// _ = "end of CoverTab[48951]"
	case uint32Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:237
		_go_fuzz_dep_.CoverTab[48952]++
															return "uint32"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:238
		// _ = "end of CoverTab[48952]"
	case uint64Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:239
		_go_fuzz_dep_.CoverTab[48953]++
															return "uint64"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:240
		// _ = "end of CoverTab[48953]"
	case float32Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:241
		_go_fuzz_dep_.CoverTab[48954]++
															return "float32"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:242
		// _ = "end of CoverTab[48954]"
	case float64Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:243
		_go_fuzz_dep_.CoverTab[48955]++
															return "float64"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:244
		// _ = "end of CoverTab[48955]"
	case stringType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:245
		_go_fuzz_dep_.CoverTab[48956]++
															return "string"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:246
		// _ = "end of CoverTab[48956]"
	case bytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:247
		_go_fuzz_dep_.CoverTab[48957]++
															return "bytes"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:248
		// _ = "end of CoverTab[48957]"
	case enumType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:249
		_go_fuzz_dep_.CoverTab[48958]++
															return "enum"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:250
		// _ = "end of CoverTab[48958]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:251
		_go_fuzz_dep_.CoverTab[48959]++
															switch v := v.getIface().(type) {
		case Message:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:253
			_go_fuzz_dep_.CoverTab[48960]++
																return "message"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:254
			// _ = "end of CoverTab[48960]"
		case List:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:255
			_go_fuzz_dep_.CoverTab[48961]++
																return "list"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:256
			// _ = "end of CoverTab[48961]"
		case Map:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:257
			_go_fuzz_dep_.CoverTab[48962]++
																return "map"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:258
			// _ = "end of CoverTab[48962]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:259
			_go_fuzz_dep_.CoverTab[48963]++
																return fmt.Sprintf("<unknown: %T>", v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:260
			// _ = "end of CoverTab[48963]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:261
		// _ = "end of CoverTab[48959]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:262
	// _ = "end of CoverTab[48947]"
}

func (v Value) panicMessage(what string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:265
	_go_fuzz_dep_.CoverTab[48964]++
														return fmt.Sprintf("type mismatch: cannot convert %v to %s", v.typeName(), what)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:266
	// _ = "end of CoverTab[48964]"
}

// Bool returns v as a bool and panics if the type is not a bool.
func (v Value) Bool() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:270
	_go_fuzz_dep_.CoverTab[48965]++
														switch v.typ {
	case boolType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:272
		_go_fuzz_dep_.CoverTab[48966]++
															return v.num > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:273
		// _ = "end of CoverTab[48966]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:274
		_go_fuzz_dep_.CoverTab[48967]++
															panic(v.panicMessage("bool"))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:275
		// _ = "end of CoverTab[48967]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:276
	// _ = "end of CoverTab[48965]"
}

// Int returns v as a int64 and panics if the type is not a int32 or int64.
func (v Value) Int() int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:280
	_go_fuzz_dep_.CoverTab[48968]++
														switch v.typ {
	case int32Type, int64Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:282
		_go_fuzz_dep_.CoverTab[48969]++
															return int64(v.num)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:283
		// _ = "end of CoverTab[48969]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:284
		_go_fuzz_dep_.CoverTab[48970]++
															panic(v.panicMessage("int"))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:285
		// _ = "end of CoverTab[48970]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:286
	// _ = "end of CoverTab[48968]"
}

// Uint returns v as a uint64 and panics if the type is not a uint32 or uint64.
func (v Value) Uint() uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:290
	_go_fuzz_dep_.CoverTab[48971]++
														switch v.typ {
	case uint32Type, uint64Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:292
		_go_fuzz_dep_.CoverTab[48972]++
															return uint64(v.num)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:293
		// _ = "end of CoverTab[48972]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:294
		_go_fuzz_dep_.CoverTab[48973]++
															panic(v.panicMessage("uint"))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:295
		// _ = "end of CoverTab[48973]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:296
	// _ = "end of CoverTab[48971]"
}

// Float returns v as a float64 and panics if the type is not a float32 or float64.
func (v Value) Float() float64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:300
	_go_fuzz_dep_.CoverTab[48974]++
														switch v.typ {
	case float32Type, float64Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:302
		_go_fuzz_dep_.CoverTab[48975]++
															return math.Float64frombits(uint64(v.num))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:303
		// _ = "end of CoverTab[48975]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:304
		_go_fuzz_dep_.CoverTab[48976]++
															panic(v.panicMessage("float"))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:305
		// _ = "end of CoverTab[48976]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:306
	// _ = "end of CoverTab[48974]"
}

// String returns v as a string. Since this method implements fmt.Stringer,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:309
// this returns the formatted string value for any non-string type.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:311
func (v Value) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:311
	_go_fuzz_dep_.CoverTab[48977]++
														switch v.typ {
	case stringType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:313
		_go_fuzz_dep_.CoverTab[48978]++
															return v.getString()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:314
		// _ = "end of CoverTab[48978]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:315
		_go_fuzz_dep_.CoverTab[48979]++
															return fmt.Sprint(v.Interface())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:316
		// _ = "end of CoverTab[48979]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:317
	// _ = "end of CoverTab[48977]"
}

// Bytes returns v as a []byte and panics if the type is not a []byte.
func (v Value) Bytes() []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:321
	_go_fuzz_dep_.CoverTab[48980]++
														switch v.typ {
	case bytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:323
		_go_fuzz_dep_.CoverTab[48981]++
															return v.getBytes()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:324
		// _ = "end of CoverTab[48981]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:325
		_go_fuzz_dep_.CoverTab[48982]++
															panic(v.panicMessage("bytes"))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:326
		// _ = "end of CoverTab[48982]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:327
	// _ = "end of CoverTab[48980]"
}

// Enum returns v as a EnumNumber and panics if the type is not a EnumNumber.
func (v Value) Enum() EnumNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:331
	_go_fuzz_dep_.CoverTab[48983]++
														switch v.typ {
	case enumType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:333
		_go_fuzz_dep_.CoverTab[48984]++
															return EnumNumber(v.num)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:334
		// _ = "end of CoverTab[48984]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:335
		_go_fuzz_dep_.CoverTab[48985]++
															panic(v.panicMessage("enum"))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:336
		// _ = "end of CoverTab[48985]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:337
	// _ = "end of CoverTab[48983]"
}

// Message returns v as a Message and panics if the type is not a Message.
func (v Value) Message() Message {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:341
	_go_fuzz_dep_.CoverTab[48986]++
														switch vi := v.getIface().(type) {
	case Message:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:343
		_go_fuzz_dep_.CoverTab[48987]++
															return vi
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:344
		// _ = "end of CoverTab[48987]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:345
		_go_fuzz_dep_.CoverTab[48988]++
															panic(v.panicMessage("message"))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:346
		// _ = "end of CoverTab[48988]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:347
	// _ = "end of CoverTab[48986]"
}

// List returns v as a List and panics if the type is not a List.
func (v Value) List() List {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:351
	_go_fuzz_dep_.CoverTab[48989]++
														switch vi := v.getIface().(type) {
	case List:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:353
		_go_fuzz_dep_.CoverTab[48990]++
															return vi
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:354
		// _ = "end of CoverTab[48990]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:355
		_go_fuzz_dep_.CoverTab[48991]++
															panic(v.panicMessage("list"))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:356
		// _ = "end of CoverTab[48991]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:357
	// _ = "end of CoverTab[48989]"
}

// Map returns v as a Map and panics if the type is not a Map.
func (v Value) Map() Map {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:361
	_go_fuzz_dep_.CoverTab[48992]++
														switch vi := v.getIface().(type) {
	case Map:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:363
		_go_fuzz_dep_.CoverTab[48993]++
															return vi
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:364
		// _ = "end of CoverTab[48993]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:365
		_go_fuzz_dep_.CoverTab[48994]++
															panic(v.panicMessage("map"))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:366
		// _ = "end of CoverTab[48994]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:367
	// _ = "end of CoverTab[48992]"
}

// MapKey returns v as a MapKey and panics for invalid MapKey types.
func (v Value) MapKey() MapKey {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:371
	_go_fuzz_dep_.CoverTab[48995]++
														switch v.typ {
	case boolType, int32Type, int64Type, uint32Type, uint64Type, stringType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:373
		_go_fuzz_dep_.CoverTab[48996]++
															return MapKey(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:374
		// _ = "end of CoverTab[48996]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:375
		_go_fuzz_dep_.CoverTab[48997]++
															panic(v.panicMessage("map key"))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:376
		// _ = "end of CoverTab[48997]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:377
	// _ = "end of CoverTab[48995]"
}

// MapKey is used to index maps, where the Go type of the MapKey must match
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
// the specified key Kind (see MessageDescriptor.IsMapEntry).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
// The following shows what Go type is used to represent each proto Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//	╔═════════╤═════════════════════════════════════╗
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//	║ Go type │ Protobuf kind                       ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//	╠═════════╪═════════════════════════════════════╣
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//	║ bool    │ BoolKind                            ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//	║ int32   │ Int32Kind, Sint32Kind, Sfixed32Kind ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//	║ int64   │ Int64Kind, Sint64Kind, Sfixed64Kind ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//	║ uint32  │ Uint32Kind, Fixed32Kind             ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//	║ uint64  │ Uint64Kind, Fixed64Kind             ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//	║ string  │ StringKind                          ║
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//	╚═════════╧═════════════════════════════════════╝
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
// A MapKey is constructed and accessed through a Value:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//	k := ValueOf("hash").MapKey() // convert string to MapKey
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//	s := k.String()               // convert MapKey to string
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
// The MapKey is a strict subset of valid types used in Value;
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:380
// converting a Value to a MapKey with an invalid type panics.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:402
type MapKey value

// IsValid reports whether k is populated with a value.
func (k MapKey) IsValid() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:405
	_go_fuzz_dep_.CoverTab[48998]++
														return Value(k).IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:406
	// _ = "end of CoverTab[48998]"
}

// Interface returns k as an interface{}.
func (k MapKey) Interface() interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:410
	_go_fuzz_dep_.CoverTab[48999]++
														return Value(k).Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:411
	// _ = "end of CoverTab[48999]"
}

// Bool returns k as a bool and panics if the type is not a bool.
func (k MapKey) Bool() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:415
	_go_fuzz_dep_.CoverTab[49000]++
														return Value(k).Bool()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:416
	// _ = "end of CoverTab[49000]"
}

// Int returns k as a int64 and panics if the type is not a int32 or int64.
func (k MapKey) Int() int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:420
	_go_fuzz_dep_.CoverTab[49001]++
														return Value(k).Int()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:421
	// _ = "end of CoverTab[49001]"
}

// Uint returns k as a uint64 and panics if the type is not a uint32 or uint64.
func (k MapKey) Uint() uint64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:425
	_go_fuzz_dep_.CoverTab[49002]++
														return Value(k).Uint()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:426
	// _ = "end of CoverTab[49002]"
}

// String returns k as a string. Since this method implements fmt.Stringer,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:429
// this returns the formatted string value for any non-string type.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:431
func (k MapKey) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:431
	_go_fuzz_dep_.CoverTab[49003]++
														return Value(k).String()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:432
	// _ = "end of CoverTab[49003]"
}

// Value returns k as a Value.
func (k MapKey) Value() Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:436
	_go_fuzz_dep_.CoverTab[49004]++
														return Value(k)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:437
	// _ = "end of CoverTab[49004]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:438
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protoreflect/value_union.go:438
var _ = _go_fuzz_dep_.CoverTab
