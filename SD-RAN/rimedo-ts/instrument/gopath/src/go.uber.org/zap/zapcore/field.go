// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:21
package zapcore

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:21
)

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"time"
)

// A FieldType indicates which member of the Field union struct should be used
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:31
// and how it should be serialized.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:33
type FieldType uint8

const (
	// UnknownType is the default field type. Attempting to add it to an encoder will panic.
	UnknownType	FieldType	= iota
	// ArrayMarshalerType indicates that the field carries an ArrayMarshaler.
	ArrayMarshalerType
	// ObjectMarshalerType indicates that the field carries an ObjectMarshaler.
	ObjectMarshalerType
	// BinaryType indicates that the field carries an opaque binary blob.
	BinaryType
	// BoolType indicates that the field carries a bool.
	BoolType
	// ByteStringType indicates that the field carries UTF-8 encoded bytes.
	ByteStringType
	// Complex128Type indicates that the field carries a complex128.
	Complex128Type
	// Complex64Type indicates that the field carries a complex128.
	Complex64Type
	// DurationType indicates that the field carries a time.Duration.
	DurationType
	// Float64Type indicates that the field carries a float64.
	Float64Type
	// Float32Type indicates that the field carries a float32.
	Float32Type
	// Int64Type indicates that the field carries an int64.
	Int64Type
	// Int32Type indicates that the field carries an int32.
	Int32Type
	// Int16Type indicates that the field carries an int16.
	Int16Type
	// Int8Type indicates that the field carries an int8.
	Int8Type
	// StringType indicates that the field carries a string.
	StringType
	// TimeType indicates that the field carries a time.Time that is
	// representable by a UnixNano() stored as an int64.
	TimeType
	// TimeFullType indicates that the field carries a time.Time stored as-is.
	TimeFullType
	// Uint64Type indicates that the field carries a uint64.
	Uint64Type
	// Uint32Type indicates that the field carries a uint32.
	Uint32Type
	// Uint16Type indicates that the field carries a uint16.
	Uint16Type
	// Uint8Type indicates that the field carries a uint8.
	Uint8Type
	// UintptrType indicates that the field carries a uintptr.
	UintptrType
	// ReflectType indicates that the field carries an interface{}, which should
	// be serialized using reflection.
	ReflectType
	// NamespaceType signals the beginning of an isolated namespace. All
	// subsequent fields should be added to the new namespace.
	NamespaceType
	// StringerType indicates that the field carries a fmt.Stringer.
	StringerType
	// ErrorType indicates that the field carries an error.
	ErrorType
	// SkipType indicates that the field is a no-op.
	SkipType

	// InlineMarshalerType indicates that the field carries an ObjectMarshaler
	// that should be inlined.
	InlineMarshalerType
)

// A Field is a marshaling operation used to add a key-value pair to a logger's
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:101
// context. Most fields are lazily marshaled, so it's inexpensive to add fields
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:101
// to disabled debug-level log statements.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:104
type Field struct {
	Key		string
	Type		FieldType
	Integer		int64
	String		string
	Interface	interface{}
}

// AddTo exports a field through the ObjectEncoder interface. It's primarily
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:112
// useful to library authors, and shouldn't be necessary in most applications.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:114
func (f Field) AddTo(enc ObjectEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:114
	_go_fuzz_dep_.CoverTab[130836]++
										var err error

										switch f.Type {
	case ArrayMarshalerType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:118
		_go_fuzz_dep_.CoverTab[130838]++
											err = enc.AddArray(f.Key, f.Interface.(ArrayMarshaler))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:119
		// _ = "end of CoverTab[130838]"
	case ObjectMarshalerType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:120
		_go_fuzz_dep_.CoverTab[130839]++
											err = enc.AddObject(f.Key, f.Interface.(ObjectMarshaler))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:121
		// _ = "end of CoverTab[130839]"
	case InlineMarshalerType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:122
		_go_fuzz_dep_.CoverTab[130840]++
											err = f.Interface.(ObjectMarshaler).MarshalLogObject(enc)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:123
		// _ = "end of CoverTab[130840]"
	case BinaryType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:124
		_go_fuzz_dep_.CoverTab[130841]++
											enc.AddBinary(f.Key, f.Interface.([]byte))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:125
		// _ = "end of CoverTab[130841]"
	case BoolType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:126
		_go_fuzz_dep_.CoverTab[130842]++
											enc.AddBool(f.Key, f.Integer == 1)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:127
		// _ = "end of CoverTab[130842]"
	case ByteStringType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:128
		_go_fuzz_dep_.CoverTab[130843]++
											enc.AddByteString(f.Key, f.Interface.([]byte))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:129
		// _ = "end of CoverTab[130843]"
	case Complex128Type:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:130
		_go_fuzz_dep_.CoverTab[130844]++
											enc.AddComplex128(f.Key, f.Interface.(complex128))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:131
		// _ = "end of CoverTab[130844]"
	case Complex64Type:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:132
		_go_fuzz_dep_.CoverTab[130845]++
											enc.AddComplex64(f.Key, f.Interface.(complex64))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:133
		// _ = "end of CoverTab[130845]"
	case DurationType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:134
		_go_fuzz_dep_.CoverTab[130846]++
											enc.AddDuration(f.Key, time.Duration(f.Integer))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:135
		// _ = "end of CoverTab[130846]"
	case Float64Type:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:136
		_go_fuzz_dep_.CoverTab[130847]++
											enc.AddFloat64(f.Key, math.Float64frombits(uint64(f.Integer)))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:137
		// _ = "end of CoverTab[130847]"
	case Float32Type:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:138
		_go_fuzz_dep_.CoverTab[130848]++
											enc.AddFloat32(f.Key, math.Float32frombits(uint32(f.Integer)))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:139
		// _ = "end of CoverTab[130848]"
	case Int64Type:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:140
		_go_fuzz_dep_.CoverTab[130849]++
											enc.AddInt64(f.Key, f.Integer)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:141
		// _ = "end of CoverTab[130849]"
	case Int32Type:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:142
		_go_fuzz_dep_.CoverTab[130850]++
											enc.AddInt32(f.Key, int32(f.Integer))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:143
		// _ = "end of CoverTab[130850]"
	case Int16Type:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:144
		_go_fuzz_dep_.CoverTab[130851]++
											enc.AddInt16(f.Key, int16(f.Integer))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:145
		// _ = "end of CoverTab[130851]"
	case Int8Type:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:146
		_go_fuzz_dep_.CoverTab[130852]++
											enc.AddInt8(f.Key, int8(f.Integer))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:147
		// _ = "end of CoverTab[130852]"
	case StringType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:148
		_go_fuzz_dep_.CoverTab[130853]++
											enc.AddString(f.Key, f.String)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:149
		// _ = "end of CoverTab[130853]"
	case TimeType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:150
		_go_fuzz_dep_.CoverTab[130854]++
											if f.Interface != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:151
			_go_fuzz_dep_.CoverTab[130867]++
												enc.AddTime(f.Key, time.Unix(0, f.Integer).In(f.Interface.(*time.Location)))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:152
			// _ = "end of CoverTab[130867]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:153
			_go_fuzz_dep_.CoverTab[130868]++

												enc.AddTime(f.Key, time.Unix(0, f.Integer))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:155
			// _ = "end of CoverTab[130868]"
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:156
		// _ = "end of CoverTab[130854]"
	case TimeFullType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:157
		_go_fuzz_dep_.CoverTab[130855]++
											enc.AddTime(f.Key, f.Interface.(time.Time))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:158
		// _ = "end of CoverTab[130855]"
	case Uint64Type:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:159
		_go_fuzz_dep_.CoverTab[130856]++
											enc.AddUint64(f.Key, uint64(f.Integer))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:160
		// _ = "end of CoverTab[130856]"
	case Uint32Type:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:161
		_go_fuzz_dep_.CoverTab[130857]++
											enc.AddUint32(f.Key, uint32(f.Integer))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:162
		// _ = "end of CoverTab[130857]"
	case Uint16Type:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:163
		_go_fuzz_dep_.CoverTab[130858]++
											enc.AddUint16(f.Key, uint16(f.Integer))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:164
		// _ = "end of CoverTab[130858]"
	case Uint8Type:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:165
		_go_fuzz_dep_.CoverTab[130859]++
											enc.AddUint8(f.Key, uint8(f.Integer))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:166
		// _ = "end of CoverTab[130859]"
	case UintptrType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:167
		_go_fuzz_dep_.CoverTab[130860]++
											enc.AddUintptr(f.Key, uintptr(f.Integer))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:168
		// _ = "end of CoverTab[130860]"
	case ReflectType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:169
		_go_fuzz_dep_.CoverTab[130861]++
											err = enc.AddReflected(f.Key, f.Interface)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:170
		// _ = "end of CoverTab[130861]"
	case NamespaceType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:171
		_go_fuzz_dep_.CoverTab[130862]++
											enc.OpenNamespace(f.Key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:172
		// _ = "end of CoverTab[130862]"
	case StringerType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:173
		_go_fuzz_dep_.CoverTab[130863]++
											err = encodeStringer(f.Key, f.Interface, enc)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:174
		// _ = "end of CoverTab[130863]"
	case ErrorType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:175
		_go_fuzz_dep_.CoverTab[130864]++
											err = encodeError(f.Key, f.Interface.(error), enc)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:176
		// _ = "end of CoverTab[130864]"
	case SkipType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:177
		_go_fuzz_dep_.CoverTab[130865]++
											break
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:178
		// _ = "end of CoverTab[130865]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:179
		_go_fuzz_dep_.CoverTab[130866]++
											panic(fmt.Sprintf("unknown field type: %v", f))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:180
		// _ = "end of CoverTab[130866]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:181
	// _ = "end of CoverTab[130836]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:181
	_go_fuzz_dep_.CoverTab[130837]++

										if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:183
		_go_fuzz_dep_.CoverTab[130869]++
											enc.AddString(fmt.Sprintf("%sError", f.Key), err.Error())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:184
		// _ = "end of CoverTab[130869]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:185
		_go_fuzz_dep_.CoverTab[130870]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:185
		// _ = "end of CoverTab[130870]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:185
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:185
	// _ = "end of CoverTab[130837]"
}

// Equals returns whether two fields are equal. For non-primitive types such as
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:188
// errors, marshalers, or reflect types, it uses reflect.DeepEqual.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:190
func (f Field) Equals(other Field) bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:190
	_go_fuzz_dep_.CoverTab[130871]++
										if f.Type != other.Type {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:191
		_go_fuzz_dep_.CoverTab[130874]++
											return false
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:192
		// _ = "end of CoverTab[130874]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:193
		_go_fuzz_dep_.CoverTab[130875]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:193
		// _ = "end of CoverTab[130875]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:193
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:193
	// _ = "end of CoverTab[130871]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:193
	_go_fuzz_dep_.CoverTab[130872]++
										if f.Key != other.Key {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:194
		_go_fuzz_dep_.CoverTab[130876]++
											return false
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:195
		// _ = "end of CoverTab[130876]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:196
		_go_fuzz_dep_.CoverTab[130877]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:196
		// _ = "end of CoverTab[130877]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:196
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:196
	// _ = "end of CoverTab[130872]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:196
	_go_fuzz_dep_.CoverTab[130873]++

										switch f.Type {
	case BinaryType, ByteStringType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:199
		_go_fuzz_dep_.CoverTab[130878]++
											return bytes.Equal(f.Interface.([]byte), other.Interface.([]byte))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:200
		// _ = "end of CoverTab[130878]"
	case ArrayMarshalerType, ObjectMarshalerType, ErrorType, ReflectType:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:201
		_go_fuzz_dep_.CoverTab[130879]++
											return reflect.DeepEqual(f.Interface, other.Interface)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:202
		// _ = "end of CoverTab[130879]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:203
		_go_fuzz_dep_.CoverTab[130880]++
											return f == other
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:204
		// _ = "end of CoverTab[130880]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:205
	// _ = "end of CoverTab[130873]"
}

func addFields(enc ObjectEncoder, fields []Field) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:208
	_go_fuzz_dep_.CoverTab[130881]++
										for i := range fields {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:209
		_go_fuzz_dep_.CoverTab[130882]++
											fields[i].AddTo(enc)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:210
		// _ = "end of CoverTab[130882]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:211
	// _ = "end of CoverTab[130881]"
}

func encodeStringer(key string, stringer interface{}, enc ObjectEncoder) (retErr error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:214
	_go_fuzz_dep_.CoverTab[130883]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:217
	defer func() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:217
		_go_fuzz_dep_.CoverTab[130885]++
											if err := recover(); err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:218
			_go_fuzz_dep_.CoverTab[130886]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:222
			if v := reflect.ValueOf(stringer); v.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:222
				_go_fuzz_dep_.CoverTab[130888]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:222
				return v.IsNil()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:222
				// _ = "end of CoverTab[130888]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:222
			}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:222
				_go_fuzz_dep_.CoverTab[130889]++
													enc.AddString(key, "<nil>")
													return
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:224
				// _ = "end of CoverTab[130889]"
			} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:225
				_go_fuzz_dep_.CoverTab[130890]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:225
				// _ = "end of CoverTab[130890]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:225
			}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:225
			// _ = "end of CoverTab[130886]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:225
			_go_fuzz_dep_.CoverTab[130887]++

												retErr = fmt.Errorf("PANIC=%v", err)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:227
			// _ = "end of CoverTab[130887]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:228
			_go_fuzz_dep_.CoverTab[130891]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:228
			// _ = "end of CoverTab[130891]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:228
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:228
		// _ = "end of CoverTab[130885]"
	}()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:229
	// _ = "end of CoverTab[130883]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:229
	_go_fuzz_dep_.CoverTab[130884]++

										enc.AddString(key, stringer.(fmt.Stringer).String())
										return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:232
	// _ = "end of CoverTab[130884]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:233
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/field.go:233
var _ = _go_fuzz_dep_.CoverTab
