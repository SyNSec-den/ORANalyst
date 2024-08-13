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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:21
)

import (
	"fmt"
	"math"
	"time"

	"go.uber.org/zap/zapcore"
)

// Field is an alias for Field. Aliasing this type dramatically
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:31
// improves the navigability of this package's API documentation.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:33
type Field = zapcore.Field

var (
	_minTimeInt64	= time.Unix(0, math.MinInt64)
	_maxTimeInt64	= time.Unix(0, math.MaxInt64)
)

// Skip constructs a no-op field, which is often useful when handling invalid
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:40
// inputs in other Field constructors.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:42
func Skip() Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:42
	_go_fuzz_dep_.CoverTab[131429]++
									return Field{Type: zapcore.SkipType}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:43
	// _ = "end of CoverTab[131429]"
}

// nilField returns a field which will marshal explicitly as nil. See motivation
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:46
// in https://github.com/uber-go/zap/issues/753 . If we ever make breaking
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:46
// changes and add zapcore.NilType and zapcore.ObjectEncoder.AddNil, the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:46
// implementation here should be changed to reflect that.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:50
func nilField(key string) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:50
	_go_fuzz_dep_.CoverTab[131430]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:50
	return Reflect(key, nil)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:50
	// _ = "end of CoverTab[131430]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:50
}

// Binary constructs a field that carries an opaque binary blob.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:52
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:52
// Binary data is serialized in an encoding-appropriate format. For example,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:52
// zap's JSON encoder base64-encodes binary blobs. To log UTF-8 encoded text,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:52
// use ByteString.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:57
func Binary(key string, val []byte) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:57
	_go_fuzz_dep_.CoverTab[131431]++
									return Field{Key: key, Type: zapcore.BinaryType, Interface: val}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:58
	// _ = "end of CoverTab[131431]"
}

// Bool constructs a field that carries a bool.
func Bool(key string, val bool) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:62
	_go_fuzz_dep_.CoverTab[131432]++
									var ival int64
									if val {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:64
		_go_fuzz_dep_.CoverTab[131434]++
										ival = 1
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:65
		// _ = "end of CoverTab[131434]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:66
		_go_fuzz_dep_.CoverTab[131435]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:66
		// _ = "end of CoverTab[131435]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:66
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:66
	// _ = "end of CoverTab[131432]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:66
	_go_fuzz_dep_.CoverTab[131433]++
									return Field{Key: key, Type: zapcore.BoolType, Integer: ival}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:67
	// _ = "end of CoverTab[131433]"
}

// Boolp constructs a field that carries a *bool. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:70
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:72
func Boolp(key string, val *bool) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:72
	_go_fuzz_dep_.CoverTab[131436]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:73
		_go_fuzz_dep_.CoverTab[131438]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:74
		// _ = "end of CoverTab[131438]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:75
		_go_fuzz_dep_.CoverTab[131439]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:75
		// _ = "end of CoverTab[131439]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:75
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:75
	// _ = "end of CoverTab[131436]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:75
	_go_fuzz_dep_.CoverTab[131437]++
									return Bool(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:76
	// _ = "end of CoverTab[131437]"
}

// ByteString constructs a field that carries UTF-8 encoded text as a []byte.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:79
// To log opaque binary blobs (which aren't necessarily valid UTF-8), use
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:79
// Binary.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:82
func ByteString(key string, val []byte) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:82
	_go_fuzz_dep_.CoverTab[131440]++
									return Field{Key: key, Type: zapcore.ByteStringType, Interface: val}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:83
	// _ = "end of CoverTab[131440]"
}

// Complex128 constructs a field that carries a complex number. Unlike most
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:86
// numeric fields, this costs an allocation (to convert the complex128 to
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:86
// interface{}).
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:89
func Complex128(key string, val complex128) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:89
	_go_fuzz_dep_.CoverTab[131441]++
									return Field{Key: key, Type: zapcore.Complex128Type, Interface: val}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:90
	// _ = "end of CoverTab[131441]"
}

// Complex128p constructs a field that carries a *complex128. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:93
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:95
func Complex128p(key string, val *complex128) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:95
	_go_fuzz_dep_.CoverTab[131442]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:96
		_go_fuzz_dep_.CoverTab[131444]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:97
		// _ = "end of CoverTab[131444]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:98
		_go_fuzz_dep_.CoverTab[131445]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:98
		// _ = "end of CoverTab[131445]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:98
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:98
	// _ = "end of CoverTab[131442]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:98
	_go_fuzz_dep_.CoverTab[131443]++
									return Complex128(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:99
	// _ = "end of CoverTab[131443]"
}

// Complex64 constructs a field that carries a complex number. Unlike most
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:102
// numeric fields, this costs an allocation (to convert the complex64 to
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:102
// interface{}).
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:105
func Complex64(key string, val complex64) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:105
	_go_fuzz_dep_.CoverTab[131446]++
									return Field{Key: key, Type: zapcore.Complex64Type, Interface: val}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:106
	// _ = "end of CoverTab[131446]"
}

// Complex64p constructs a field that carries a *complex64. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:109
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:111
func Complex64p(key string, val *complex64) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:111
	_go_fuzz_dep_.CoverTab[131447]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:112
		_go_fuzz_dep_.CoverTab[131449]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:113
		// _ = "end of CoverTab[131449]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:114
		_go_fuzz_dep_.CoverTab[131450]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:114
		// _ = "end of CoverTab[131450]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:114
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:114
	// _ = "end of CoverTab[131447]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:114
	_go_fuzz_dep_.CoverTab[131448]++
									return Complex64(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:115
	// _ = "end of CoverTab[131448]"
}

// Float64 constructs a field that carries a float64. The way the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:118
// floating-point value is represented is encoder-dependent, so marshaling is
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:118
// necessarily lazy.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:121
func Float64(key string, val float64) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:121
	_go_fuzz_dep_.CoverTab[131451]++
									return Field{Key: key, Type: zapcore.Float64Type, Integer: int64(math.Float64bits(val))}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:122
	// _ = "end of CoverTab[131451]"
}

// Float64p constructs a field that carries a *float64. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:125
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:127
func Float64p(key string, val *float64) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:127
	_go_fuzz_dep_.CoverTab[131452]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:128
		_go_fuzz_dep_.CoverTab[131454]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:129
		// _ = "end of CoverTab[131454]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:130
		_go_fuzz_dep_.CoverTab[131455]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:130
		// _ = "end of CoverTab[131455]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:130
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:130
	// _ = "end of CoverTab[131452]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:130
	_go_fuzz_dep_.CoverTab[131453]++
									return Float64(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:131
	// _ = "end of CoverTab[131453]"
}

// Float32 constructs a field that carries a float32. The way the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:134
// floating-point value is represented is encoder-dependent, so marshaling is
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:134
// necessarily lazy.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:137
func Float32(key string, val float32) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:137
	_go_fuzz_dep_.CoverTab[131456]++
									return Field{Key: key, Type: zapcore.Float32Type, Integer: int64(math.Float32bits(val))}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:138
	// _ = "end of CoverTab[131456]"
}

// Float32p constructs a field that carries a *float32. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:141
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:143
func Float32p(key string, val *float32) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:143
	_go_fuzz_dep_.CoverTab[131457]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:144
		_go_fuzz_dep_.CoverTab[131459]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:145
		// _ = "end of CoverTab[131459]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:146
		_go_fuzz_dep_.CoverTab[131460]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:146
		// _ = "end of CoverTab[131460]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:146
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:146
	// _ = "end of CoverTab[131457]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:146
	_go_fuzz_dep_.CoverTab[131458]++
									return Float32(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:147
	// _ = "end of CoverTab[131458]"
}

// Int constructs a field with the given key and value.
func Int(key string, val int) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:151
	_go_fuzz_dep_.CoverTab[131461]++
									return Int64(key, int64(val))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:152
	// _ = "end of CoverTab[131461]"
}

// Intp constructs a field that carries a *int. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:155
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:157
func Intp(key string, val *int) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:157
	_go_fuzz_dep_.CoverTab[131462]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:158
		_go_fuzz_dep_.CoverTab[131464]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:159
		// _ = "end of CoverTab[131464]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:160
		_go_fuzz_dep_.CoverTab[131465]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:160
		// _ = "end of CoverTab[131465]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:160
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:160
	// _ = "end of CoverTab[131462]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:160
	_go_fuzz_dep_.CoverTab[131463]++
									return Int(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:161
	// _ = "end of CoverTab[131463]"
}

// Int64 constructs a field with the given key and value.
func Int64(key string, val int64) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:165
	_go_fuzz_dep_.CoverTab[131466]++
									return Field{Key: key, Type: zapcore.Int64Type, Integer: val}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:166
	// _ = "end of CoverTab[131466]"
}

// Int64p constructs a field that carries a *int64. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:169
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:171
func Int64p(key string, val *int64) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:171
	_go_fuzz_dep_.CoverTab[131467]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:172
		_go_fuzz_dep_.CoverTab[131469]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:173
		// _ = "end of CoverTab[131469]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:174
		_go_fuzz_dep_.CoverTab[131470]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:174
		// _ = "end of CoverTab[131470]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:174
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:174
	// _ = "end of CoverTab[131467]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:174
	_go_fuzz_dep_.CoverTab[131468]++
									return Int64(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:175
	// _ = "end of CoverTab[131468]"
}

// Int32 constructs a field with the given key and value.
func Int32(key string, val int32) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:179
	_go_fuzz_dep_.CoverTab[131471]++
									return Field{Key: key, Type: zapcore.Int32Type, Integer: int64(val)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:180
	// _ = "end of CoverTab[131471]"
}

// Int32p constructs a field that carries a *int32. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:183
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:185
func Int32p(key string, val *int32) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:185
	_go_fuzz_dep_.CoverTab[131472]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:186
		_go_fuzz_dep_.CoverTab[131474]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:187
		// _ = "end of CoverTab[131474]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:188
		_go_fuzz_dep_.CoverTab[131475]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:188
		// _ = "end of CoverTab[131475]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:188
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:188
	// _ = "end of CoverTab[131472]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:188
	_go_fuzz_dep_.CoverTab[131473]++
									return Int32(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:189
	// _ = "end of CoverTab[131473]"
}

// Int16 constructs a field with the given key and value.
func Int16(key string, val int16) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:193
	_go_fuzz_dep_.CoverTab[131476]++
									return Field{Key: key, Type: zapcore.Int16Type, Integer: int64(val)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:194
	// _ = "end of CoverTab[131476]"
}

// Int16p constructs a field that carries a *int16. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:197
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:199
func Int16p(key string, val *int16) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:199
	_go_fuzz_dep_.CoverTab[131477]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:200
		_go_fuzz_dep_.CoverTab[131479]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:201
		// _ = "end of CoverTab[131479]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:202
		_go_fuzz_dep_.CoverTab[131480]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:202
		// _ = "end of CoverTab[131480]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:202
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:202
	// _ = "end of CoverTab[131477]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:202
	_go_fuzz_dep_.CoverTab[131478]++
									return Int16(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:203
	// _ = "end of CoverTab[131478]"
}

// Int8 constructs a field with the given key and value.
func Int8(key string, val int8) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:207
	_go_fuzz_dep_.CoverTab[131481]++
									return Field{Key: key, Type: zapcore.Int8Type, Integer: int64(val)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:208
	// _ = "end of CoverTab[131481]"
}

// Int8p constructs a field that carries a *int8. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:211
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:213
func Int8p(key string, val *int8) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:213
	_go_fuzz_dep_.CoverTab[131482]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:214
		_go_fuzz_dep_.CoverTab[131484]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:215
		// _ = "end of CoverTab[131484]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:216
		_go_fuzz_dep_.CoverTab[131485]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:216
		// _ = "end of CoverTab[131485]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:216
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:216
	// _ = "end of CoverTab[131482]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:216
	_go_fuzz_dep_.CoverTab[131483]++
									return Int8(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:217
	// _ = "end of CoverTab[131483]"
}

// String constructs a field with the given key and value.
func String(key string, val string) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:221
	_go_fuzz_dep_.CoverTab[131486]++
									return Field{Key: key, Type: zapcore.StringType, String: val}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:222
	// _ = "end of CoverTab[131486]"
}

// Stringp constructs a field that carries a *string. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:225
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:227
func Stringp(key string, val *string) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:227
	_go_fuzz_dep_.CoverTab[131487]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:228
		_go_fuzz_dep_.CoverTab[131489]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:229
		// _ = "end of CoverTab[131489]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:230
		_go_fuzz_dep_.CoverTab[131490]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:230
		// _ = "end of CoverTab[131490]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:230
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:230
	// _ = "end of CoverTab[131487]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:230
	_go_fuzz_dep_.CoverTab[131488]++
									return String(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:231
	// _ = "end of CoverTab[131488]"
}

// Uint constructs a field with the given key and value.
func Uint(key string, val uint) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:235
	_go_fuzz_dep_.CoverTab[131491]++
									return Uint64(key, uint64(val))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:236
	// _ = "end of CoverTab[131491]"
}

// Uintp constructs a field that carries a *uint. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:239
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:241
func Uintp(key string, val *uint) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:241
	_go_fuzz_dep_.CoverTab[131492]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:242
		_go_fuzz_dep_.CoverTab[131494]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:243
		// _ = "end of CoverTab[131494]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:244
		_go_fuzz_dep_.CoverTab[131495]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:244
		// _ = "end of CoverTab[131495]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:244
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:244
	// _ = "end of CoverTab[131492]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:244
	_go_fuzz_dep_.CoverTab[131493]++
									return Uint(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:245
	// _ = "end of CoverTab[131493]"
}

// Uint64 constructs a field with the given key and value.
func Uint64(key string, val uint64) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:249
	_go_fuzz_dep_.CoverTab[131496]++
									return Field{Key: key, Type: zapcore.Uint64Type, Integer: int64(val)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:250
	// _ = "end of CoverTab[131496]"
}

// Uint64p constructs a field that carries a *uint64. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:253
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:255
func Uint64p(key string, val *uint64) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:255
	_go_fuzz_dep_.CoverTab[131497]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:256
		_go_fuzz_dep_.CoverTab[131499]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:257
		// _ = "end of CoverTab[131499]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:258
		_go_fuzz_dep_.CoverTab[131500]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:258
		// _ = "end of CoverTab[131500]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:258
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:258
	// _ = "end of CoverTab[131497]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:258
	_go_fuzz_dep_.CoverTab[131498]++
									return Uint64(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:259
	// _ = "end of CoverTab[131498]"
}

// Uint32 constructs a field with the given key and value.
func Uint32(key string, val uint32) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:263
	_go_fuzz_dep_.CoverTab[131501]++
									return Field{Key: key, Type: zapcore.Uint32Type, Integer: int64(val)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:264
	// _ = "end of CoverTab[131501]"
}

// Uint32p constructs a field that carries a *uint32. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:267
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:269
func Uint32p(key string, val *uint32) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:269
	_go_fuzz_dep_.CoverTab[131502]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:270
		_go_fuzz_dep_.CoverTab[131504]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:271
		// _ = "end of CoverTab[131504]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:272
		_go_fuzz_dep_.CoverTab[131505]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:272
		// _ = "end of CoverTab[131505]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:272
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:272
	// _ = "end of CoverTab[131502]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:272
	_go_fuzz_dep_.CoverTab[131503]++
									return Uint32(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:273
	// _ = "end of CoverTab[131503]"
}

// Uint16 constructs a field with the given key and value.
func Uint16(key string, val uint16) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:277
	_go_fuzz_dep_.CoverTab[131506]++
									return Field{Key: key, Type: zapcore.Uint16Type, Integer: int64(val)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:278
	// _ = "end of CoverTab[131506]"
}

// Uint16p constructs a field that carries a *uint16. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:281
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:283
func Uint16p(key string, val *uint16) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:283
	_go_fuzz_dep_.CoverTab[131507]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:284
		_go_fuzz_dep_.CoverTab[131509]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:285
		// _ = "end of CoverTab[131509]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:286
		_go_fuzz_dep_.CoverTab[131510]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:286
		// _ = "end of CoverTab[131510]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:286
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:286
	// _ = "end of CoverTab[131507]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:286
	_go_fuzz_dep_.CoverTab[131508]++
									return Uint16(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:287
	// _ = "end of CoverTab[131508]"
}

// Uint8 constructs a field with the given key and value.
func Uint8(key string, val uint8) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:291
	_go_fuzz_dep_.CoverTab[131511]++
									return Field{Key: key, Type: zapcore.Uint8Type, Integer: int64(val)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:292
	// _ = "end of CoverTab[131511]"
}

// Uint8p constructs a field that carries a *uint8. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:295
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:297
func Uint8p(key string, val *uint8) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:297
	_go_fuzz_dep_.CoverTab[131512]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:298
		_go_fuzz_dep_.CoverTab[131514]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:299
		// _ = "end of CoverTab[131514]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:300
		_go_fuzz_dep_.CoverTab[131515]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:300
		// _ = "end of CoverTab[131515]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:300
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:300
	// _ = "end of CoverTab[131512]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:300
	_go_fuzz_dep_.CoverTab[131513]++
									return Uint8(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:301
	// _ = "end of CoverTab[131513]"
}

// Uintptr constructs a field with the given key and value.
func Uintptr(key string, val uintptr) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:305
	_go_fuzz_dep_.CoverTab[131516]++
									return Field{Key: key, Type: zapcore.UintptrType, Integer: int64(val)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:306
	// _ = "end of CoverTab[131516]"
}

// Uintptrp constructs a field that carries a *uintptr. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:309
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:311
func Uintptrp(key string, val *uintptr) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:311
	_go_fuzz_dep_.CoverTab[131517]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:312
		_go_fuzz_dep_.CoverTab[131519]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:313
		// _ = "end of CoverTab[131519]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:314
		_go_fuzz_dep_.CoverTab[131520]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:314
		// _ = "end of CoverTab[131520]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:314
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:314
	// _ = "end of CoverTab[131517]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:314
	_go_fuzz_dep_.CoverTab[131518]++
									return Uintptr(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:315
	// _ = "end of CoverTab[131518]"
}

// Reflect constructs a field with the given key and an arbitrary object. It uses
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:318
// an encoding-appropriate, reflection-based function to lazily serialize nearly
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:318
// any object into the logging context, but it's relatively slow and
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:318
// allocation-heavy. Outside tests, Any is always a better choice.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:318
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:318
// If encoding fails (e.g., trying to serialize a map[int]string to JSON), Reflect
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:318
// includes the error message in the final log output.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:325
func Reflect(key string, val interface{}) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:325
	_go_fuzz_dep_.CoverTab[131521]++
									return Field{Key: key, Type: zapcore.ReflectType, Interface: val}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:326
	// _ = "end of CoverTab[131521]"
}

// Namespace creates a named, isolated scope within the logger's context. All
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:329
// subsequent fields will be added to the new namespace.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:329
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:329
// This helps prevent key collisions when injecting loggers into sub-components
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:329
// or third-party libraries.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:334
func Namespace(key string) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:334
	_go_fuzz_dep_.CoverTab[131522]++
									return Field{Key: key, Type: zapcore.NamespaceType}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:335
	// _ = "end of CoverTab[131522]"
}

// Stringer constructs a field with the given key and the output of the value's
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:338
// String method. The Stringer's String method is called lazily.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:340
func Stringer(key string, val fmt.Stringer) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:340
	_go_fuzz_dep_.CoverTab[131523]++
									return Field{Key: key, Type: zapcore.StringerType, Interface: val}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:341
	// _ = "end of CoverTab[131523]"
}

// Time constructs a Field with the given key and value. The encoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:344
// controls how the time is serialized.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:346
func Time(key string, val time.Time) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:346
	_go_fuzz_dep_.CoverTab[131524]++
									if val.Before(_minTimeInt64) || func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:347
		_go_fuzz_dep_.CoverTab[131526]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:347
		return val.After(_maxTimeInt64)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:347
		// _ = "end of CoverTab[131526]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:347
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:347
		_go_fuzz_dep_.CoverTab[131527]++
										return Field{Key: key, Type: zapcore.TimeFullType, Interface: val}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:348
		// _ = "end of CoverTab[131527]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:349
		_go_fuzz_dep_.CoverTab[131528]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:349
		// _ = "end of CoverTab[131528]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:349
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:349
	// _ = "end of CoverTab[131524]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:349
	_go_fuzz_dep_.CoverTab[131525]++
									return Field{Key: key, Type: zapcore.TimeType, Integer: val.UnixNano(), Interface: val.Location()}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:350
	// _ = "end of CoverTab[131525]"
}

// Timep constructs a field that carries a *time.Time. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:353
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:355
func Timep(key string, val *time.Time) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:355
	_go_fuzz_dep_.CoverTab[131529]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:356
		_go_fuzz_dep_.CoverTab[131531]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:357
		// _ = "end of CoverTab[131531]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:358
		_go_fuzz_dep_.CoverTab[131532]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:358
		// _ = "end of CoverTab[131532]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:358
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:358
	// _ = "end of CoverTab[131529]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:358
	_go_fuzz_dep_.CoverTab[131530]++
									return Time(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:359
	// _ = "end of CoverTab[131530]"
}

// Stack constructs a field that stores a stacktrace of the current goroutine
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:362
// under provided key. Keep in mind that taking a stacktrace is eager and
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:362
// expensive (relatively speaking); this function both makes an allocation and
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:362
// takes about two microseconds.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:366
func Stack(key string) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:366
	_go_fuzz_dep_.CoverTab[131533]++
									return StackSkip(key, 1)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:367
	// _ = "end of CoverTab[131533]"
}

// StackSkip constructs a field similarly to Stack, but also skips the given
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:370
// number of frames from the top of the stacktrace.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:372
func StackSkip(key string, skip int) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:372
	_go_fuzz_dep_.CoverTab[131534]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:377
	return String(key, takeStacktrace(skip+1))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:377
	// _ = "end of CoverTab[131534]"
}

// Duration constructs a field with the given key and value. The encoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:380
// controls how the duration is serialized.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:382
func Duration(key string, val time.Duration) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:382
	_go_fuzz_dep_.CoverTab[131535]++
									return Field{Key: key, Type: zapcore.DurationType, Integer: int64(val)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:383
	// _ = "end of CoverTab[131535]"
}

// Durationp constructs a field that carries a *time.Duration. The returned Field will safely
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:386
// and explicitly represent `nil` when appropriate.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:388
func Durationp(key string, val *time.Duration) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:388
	_go_fuzz_dep_.CoverTab[131536]++
									if val == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:389
		_go_fuzz_dep_.CoverTab[131538]++
										return nilField(key)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:390
		// _ = "end of CoverTab[131538]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:391
		_go_fuzz_dep_.CoverTab[131539]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:391
		// _ = "end of CoverTab[131539]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:391
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:391
	// _ = "end of CoverTab[131536]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:391
	_go_fuzz_dep_.CoverTab[131537]++
									return Duration(key, *val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:392
	// _ = "end of CoverTab[131537]"
}

// Object constructs a field with the given key and ObjectMarshaler. It
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:395
// provides a flexible, but still type-safe and efficient, way to add map- or
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:395
// struct-like user-defined types to the logging context. The struct's
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:395
// MarshalLogObject method is called lazily.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:399
func Object(key string, val zapcore.ObjectMarshaler) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:399
	_go_fuzz_dep_.CoverTab[131540]++
									return Field{Key: key, Type: zapcore.ObjectMarshalerType, Interface: val}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:400
	// _ = "end of CoverTab[131540]"
}

// Inline constructs a Field that is similar to Object, but it
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:403
// will add the elements of the provided ObjectMarshaler to the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:403
// current namespace.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:406
func Inline(val zapcore.ObjectMarshaler) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:406
	_go_fuzz_dep_.CoverTab[131541]++
									return zapcore.Field{
		Type:		zapcore.InlineMarshalerType,
		Interface:	val,
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:410
	// _ = "end of CoverTab[131541]"
}

// Any takes a key and an arbitrary value and chooses the best way to represent
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:413
// them as a field, falling back to a reflection-based approach only if
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:413
// necessary.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:413
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:413
// Since byte/uint8 and rune/int32 are aliases, Any can't differentiate between
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:413
// them. To minimize surprises, []byte values are treated as binary blobs, byte
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:413
// values are treated as uint8, and runes are always treated as integers.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:420
func Any(key string, value interface{}) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:420
	_go_fuzz_dep_.CoverTab[131542]++
									switch val := value.(type) {
	case zapcore.ObjectMarshaler:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:422
		_go_fuzz_dep_.CoverTab[131543]++
										return Object(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:423
		// _ = "end of CoverTab[131543]"
	case zapcore.ArrayMarshaler:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:424
		_go_fuzz_dep_.CoverTab[131544]++
										return Array(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:425
		// _ = "end of CoverTab[131544]"
	case bool:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:426
		_go_fuzz_dep_.CoverTab[131545]++
										return Bool(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:427
		// _ = "end of CoverTab[131545]"
	case *bool:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:428
		_go_fuzz_dep_.CoverTab[131546]++
										return Boolp(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:429
		// _ = "end of CoverTab[131546]"
	case []bool:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:430
		_go_fuzz_dep_.CoverTab[131547]++
										return Bools(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:431
		// _ = "end of CoverTab[131547]"
	case complex128:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:432
		_go_fuzz_dep_.CoverTab[131548]++
										return Complex128(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:433
		// _ = "end of CoverTab[131548]"
	case *complex128:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:434
		_go_fuzz_dep_.CoverTab[131549]++
										return Complex128p(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:435
		// _ = "end of CoverTab[131549]"
	case []complex128:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:436
		_go_fuzz_dep_.CoverTab[131550]++
										return Complex128s(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:437
		// _ = "end of CoverTab[131550]"
	case complex64:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:438
		_go_fuzz_dep_.CoverTab[131551]++
										return Complex64(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:439
		// _ = "end of CoverTab[131551]"
	case *complex64:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:440
		_go_fuzz_dep_.CoverTab[131552]++
										return Complex64p(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:441
		// _ = "end of CoverTab[131552]"
	case []complex64:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:442
		_go_fuzz_dep_.CoverTab[131553]++
										return Complex64s(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:443
		// _ = "end of CoverTab[131553]"
	case float64:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:444
		_go_fuzz_dep_.CoverTab[131554]++
										return Float64(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:445
		// _ = "end of CoverTab[131554]"
	case *float64:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:446
		_go_fuzz_dep_.CoverTab[131555]++
										return Float64p(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:447
		// _ = "end of CoverTab[131555]"
	case []float64:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:448
		_go_fuzz_dep_.CoverTab[131556]++
										return Float64s(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:449
		// _ = "end of CoverTab[131556]"
	case float32:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:450
		_go_fuzz_dep_.CoverTab[131557]++
										return Float32(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:451
		// _ = "end of CoverTab[131557]"
	case *float32:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:452
		_go_fuzz_dep_.CoverTab[131558]++
										return Float32p(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:453
		// _ = "end of CoverTab[131558]"
	case []float32:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:454
		_go_fuzz_dep_.CoverTab[131559]++
										return Float32s(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:455
		// _ = "end of CoverTab[131559]"
	case int:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:456
		_go_fuzz_dep_.CoverTab[131560]++
										return Int(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:457
		// _ = "end of CoverTab[131560]"
	case *int:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:458
		_go_fuzz_dep_.CoverTab[131561]++
										return Intp(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:459
		// _ = "end of CoverTab[131561]"
	case []int:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:460
		_go_fuzz_dep_.CoverTab[131562]++
										return Ints(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:461
		// _ = "end of CoverTab[131562]"
	case int64:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:462
		_go_fuzz_dep_.CoverTab[131563]++
										return Int64(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:463
		// _ = "end of CoverTab[131563]"
	case *int64:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:464
		_go_fuzz_dep_.CoverTab[131564]++
										return Int64p(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:465
		// _ = "end of CoverTab[131564]"
	case []int64:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:466
		_go_fuzz_dep_.CoverTab[131565]++
										return Int64s(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:467
		// _ = "end of CoverTab[131565]"
	case int32:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:468
		_go_fuzz_dep_.CoverTab[131566]++
										return Int32(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:469
		// _ = "end of CoverTab[131566]"
	case *int32:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:470
		_go_fuzz_dep_.CoverTab[131567]++
										return Int32p(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:471
		// _ = "end of CoverTab[131567]"
	case []int32:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:472
		_go_fuzz_dep_.CoverTab[131568]++
										return Int32s(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:473
		// _ = "end of CoverTab[131568]"
	case int16:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:474
		_go_fuzz_dep_.CoverTab[131569]++
										return Int16(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:475
		// _ = "end of CoverTab[131569]"
	case *int16:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:476
		_go_fuzz_dep_.CoverTab[131570]++
										return Int16p(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:477
		// _ = "end of CoverTab[131570]"
	case []int16:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:478
		_go_fuzz_dep_.CoverTab[131571]++
										return Int16s(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:479
		// _ = "end of CoverTab[131571]"
	case int8:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:480
		_go_fuzz_dep_.CoverTab[131572]++
										return Int8(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:481
		// _ = "end of CoverTab[131572]"
	case *int8:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:482
		_go_fuzz_dep_.CoverTab[131573]++
										return Int8p(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:483
		// _ = "end of CoverTab[131573]"
	case []int8:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:484
		_go_fuzz_dep_.CoverTab[131574]++
										return Int8s(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:485
		// _ = "end of CoverTab[131574]"
	case string:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:486
		_go_fuzz_dep_.CoverTab[131575]++
										return String(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:487
		// _ = "end of CoverTab[131575]"
	case *string:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:488
		_go_fuzz_dep_.CoverTab[131576]++
										return Stringp(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:489
		// _ = "end of CoverTab[131576]"
	case []string:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:490
		_go_fuzz_dep_.CoverTab[131577]++
										return Strings(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:491
		// _ = "end of CoverTab[131577]"
	case uint:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:492
		_go_fuzz_dep_.CoverTab[131578]++
										return Uint(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:493
		// _ = "end of CoverTab[131578]"
	case *uint:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:494
		_go_fuzz_dep_.CoverTab[131579]++
										return Uintp(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:495
		// _ = "end of CoverTab[131579]"
	case []uint:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:496
		_go_fuzz_dep_.CoverTab[131580]++
										return Uints(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:497
		// _ = "end of CoverTab[131580]"
	case uint64:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:498
		_go_fuzz_dep_.CoverTab[131581]++
										return Uint64(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:499
		// _ = "end of CoverTab[131581]"
	case *uint64:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:500
		_go_fuzz_dep_.CoverTab[131582]++
										return Uint64p(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:501
		// _ = "end of CoverTab[131582]"
	case []uint64:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:502
		_go_fuzz_dep_.CoverTab[131583]++
										return Uint64s(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:503
		// _ = "end of CoverTab[131583]"
	case uint32:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:504
		_go_fuzz_dep_.CoverTab[131584]++
										return Uint32(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:505
		// _ = "end of CoverTab[131584]"
	case *uint32:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:506
		_go_fuzz_dep_.CoverTab[131585]++
										return Uint32p(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:507
		// _ = "end of CoverTab[131585]"
	case []uint32:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:508
		_go_fuzz_dep_.CoverTab[131586]++
										return Uint32s(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:509
		// _ = "end of CoverTab[131586]"
	case uint16:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:510
		_go_fuzz_dep_.CoverTab[131587]++
										return Uint16(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:511
		// _ = "end of CoverTab[131587]"
	case *uint16:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:512
		_go_fuzz_dep_.CoverTab[131588]++
										return Uint16p(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:513
		// _ = "end of CoverTab[131588]"
	case []uint16:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:514
		_go_fuzz_dep_.CoverTab[131589]++
										return Uint16s(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:515
		// _ = "end of CoverTab[131589]"
	case uint8:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:516
		_go_fuzz_dep_.CoverTab[131590]++
										return Uint8(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:517
		// _ = "end of CoverTab[131590]"
	case *uint8:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:518
		_go_fuzz_dep_.CoverTab[131591]++
										return Uint8p(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:519
		// _ = "end of CoverTab[131591]"
	case []byte:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:520
		_go_fuzz_dep_.CoverTab[131592]++
										return Binary(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:521
		// _ = "end of CoverTab[131592]"
	case uintptr:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:522
		_go_fuzz_dep_.CoverTab[131593]++
										return Uintptr(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:523
		// _ = "end of CoverTab[131593]"
	case *uintptr:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:524
		_go_fuzz_dep_.CoverTab[131594]++
										return Uintptrp(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:525
		// _ = "end of CoverTab[131594]"
	case []uintptr:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:526
		_go_fuzz_dep_.CoverTab[131595]++
										return Uintptrs(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:527
		// _ = "end of CoverTab[131595]"
	case time.Time:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:528
		_go_fuzz_dep_.CoverTab[131596]++
										return Time(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:529
		// _ = "end of CoverTab[131596]"
	case *time.Time:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:530
		_go_fuzz_dep_.CoverTab[131597]++
										return Timep(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:531
		// _ = "end of CoverTab[131597]"
	case []time.Time:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:532
		_go_fuzz_dep_.CoverTab[131598]++
										return Times(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:533
		// _ = "end of CoverTab[131598]"
	case time.Duration:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:534
		_go_fuzz_dep_.CoverTab[131599]++
										return Duration(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:535
		// _ = "end of CoverTab[131599]"
	case *time.Duration:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:536
		_go_fuzz_dep_.CoverTab[131600]++
										return Durationp(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:537
		// _ = "end of CoverTab[131600]"
	case []time.Duration:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:538
		_go_fuzz_dep_.CoverTab[131601]++
										return Durations(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:539
		// _ = "end of CoverTab[131601]"
	case error:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:540
		_go_fuzz_dep_.CoverTab[131602]++
										return NamedError(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:541
		// _ = "end of CoverTab[131602]"
	case []error:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:542
		_go_fuzz_dep_.CoverTab[131603]++
										return Errors(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:543
		// _ = "end of CoverTab[131603]"
	case fmt.Stringer:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:544
		_go_fuzz_dep_.CoverTab[131604]++
										return Stringer(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:545
		// _ = "end of CoverTab[131604]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:546
		_go_fuzz_dep_.CoverTab[131605]++
										return Reflect(key, val)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:547
		// _ = "end of CoverTab[131605]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:548
	// _ = "end of CoverTab[131542]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:549
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/field.go:549
var _ = _go_fuzz_dep_.CoverTab
