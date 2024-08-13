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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:21
)

import (
	"time"

	"go.uber.org/zap/zapcore"
)

// Array constructs a field with the given key and ArrayMarshaler. It provides
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:29
// a flexible, but still type-safe and efficient, way to add array-like types
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:29
// to the logging context. The struct's MarshalLogArray method is called lazily.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:32
func Array(key string, val zapcore.ArrayMarshaler) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:32
	_go_fuzz_dep_.CoverTab[131262]++
									return Field{Key: key, Type: zapcore.ArrayMarshalerType, Interface: val}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:33
	// _ = "end of CoverTab[131262]"
}

// Bools constructs a field that carries a slice of bools.
func Bools(key string, bs []bool) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:37
	_go_fuzz_dep_.CoverTab[131263]++
									return Array(key, bools(bs))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:38
	// _ = "end of CoverTab[131263]"
}

// ByteStrings constructs a field that carries a slice of []byte, each of which
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:41
// must be UTF-8 encoded text.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:43
func ByteStrings(key string, bss [][]byte) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:43
	_go_fuzz_dep_.CoverTab[131264]++
									return Array(key, byteStringsArray(bss))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:44
	// _ = "end of CoverTab[131264]"
}

// Complex128s constructs a field that carries a slice of complex numbers.
func Complex128s(key string, nums []complex128) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:48
	_go_fuzz_dep_.CoverTab[131265]++
									return Array(key, complex128s(nums))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:49
	// _ = "end of CoverTab[131265]"
}

// Complex64s constructs a field that carries a slice of complex numbers.
func Complex64s(key string, nums []complex64) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:53
	_go_fuzz_dep_.CoverTab[131266]++
									return Array(key, complex64s(nums))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:54
	// _ = "end of CoverTab[131266]"
}

// Durations constructs a field that carries a slice of time.Durations.
func Durations(key string, ds []time.Duration) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:58
	_go_fuzz_dep_.CoverTab[131267]++
									return Array(key, durations(ds))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:59
	// _ = "end of CoverTab[131267]"
}

// Float64s constructs a field that carries a slice of floats.
func Float64s(key string, nums []float64) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:63
	_go_fuzz_dep_.CoverTab[131268]++
									return Array(key, float64s(nums))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:64
	// _ = "end of CoverTab[131268]"
}

// Float32s constructs a field that carries a slice of floats.
func Float32s(key string, nums []float32) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:68
	_go_fuzz_dep_.CoverTab[131269]++
									return Array(key, float32s(nums))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:69
	// _ = "end of CoverTab[131269]"
}

// Ints constructs a field that carries a slice of integers.
func Ints(key string, nums []int) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:73
	_go_fuzz_dep_.CoverTab[131270]++
									return Array(key, ints(nums))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:74
	// _ = "end of CoverTab[131270]"
}

// Int64s constructs a field that carries a slice of integers.
func Int64s(key string, nums []int64) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:78
	_go_fuzz_dep_.CoverTab[131271]++
									return Array(key, int64s(nums))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:79
	// _ = "end of CoverTab[131271]"
}

// Int32s constructs a field that carries a slice of integers.
func Int32s(key string, nums []int32) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:83
	_go_fuzz_dep_.CoverTab[131272]++
									return Array(key, int32s(nums))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:84
	// _ = "end of CoverTab[131272]"
}

// Int16s constructs a field that carries a slice of integers.
func Int16s(key string, nums []int16) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:88
	_go_fuzz_dep_.CoverTab[131273]++
									return Array(key, int16s(nums))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:89
	// _ = "end of CoverTab[131273]"
}

// Int8s constructs a field that carries a slice of integers.
func Int8s(key string, nums []int8) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:93
	_go_fuzz_dep_.CoverTab[131274]++
									return Array(key, int8s(nums))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:94
	// _ = "end of CoverTab[131274]"
}

// Strings constructs a field that carries a slice of strings.
func Strings(key string, ss []string) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:98
	_go_fuzz_dep_.CoverTab[131275]++
									return Array(key, stringArray(ss))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:99
	// _ = "end of CoverTab[131275]"
}

// Times constructs a field that carries a slice of time.Times.
func Times(key string, ts []time.Time) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:103
	_go_fuzz_dep_.CoverTab[131276]++
									return Array(key, times(ts))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:104
	// _ = "end of CoverTab[131276]"
}

// Uints constructs a field that carries a slice of unsigned integers.
func Uints(key string, nums []uint) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:108
	_go_fuzz_dep_.CoverTab[131277]++
									return Array(key, uints(nums))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:109
	// _ = "end of CoverTab[131277]"
}

// Uint64s constructs a field that carries a slice of unsigned integers.
func Uint64s(key string, nums []uint64) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:113
	_go_fuzz_dep_.CoverTab[131278]++
									return Array(key, uint64s(nums))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:114
	// _ = "end of CoverTab[131278]"
}

// Uint32s constructs a field that carries a slice of unsigned integers.
func Uint32s(key string, nums []uint32) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:118
	_go_fuzz_dep_.CoverTab[131279]++
									return Array(key, uint32s(nums))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:119
	// _ = "end of CoverTab[131279]"
}

// Uint16s constructs a field that carries a slice of unsigned integers.
func Uint16s(key string, nums []uint16) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:123
	_go_fuzz_dep_.CoverTab[131280]++
									return Array(key, uint16s(nums))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:124
	// _ = "end of CoverTab[131280]"
}

// Uint8s constructs a field that carries a slice of unsigned integers.
func Uint8s(key string, nums []uint8) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:128
	_go_fuzz_dep_.CoverTab[131281]++
									return Array(key, uint8s(nums))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:129
	// _ = "end of CoverTab[131281]"
}

// Uintptrs constructs a field that carries a slice of pointer addresses.
func Uintptrs(key string, us []uintptr) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:133
	_go_fuzz_dep_.CoverTab[131282]++
									return Array(key, uintptrs(us))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:134
	// _ = "end of CoverTab[131282]"
}

// Errors constructs a field that carries a slice of errors.
func Errors(key string, errs []error) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:138
	_go_fuzz_dep_.CoverTab[131283]++
									return Array(key, errArray(errs))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:139
	// _ = "end of CoverTab[131283]"
}

type bools []bool

func (bs bools) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:144
	_go_fuzz_dep_.CoverTab[131284]++
									for i := range bs {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:145
		_go_fuzz_dep_.CoverTab[131286]++
										arr.AppendBool(bs[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:146
		// _ = "end of CoverTab[131286]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:147
	// _ = "end of CoverTab[131284]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:147
	_go_fuzz_dep_.CoverTab[131285]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:148
	// _ = "end of CoverTab[131285]"
}

type byteStringsArray [][]byte

func (bss byteStringsArray) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:153
	_go_fuzz_dep_.CoverTab[131287]++
									for i := range bss {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:154
		_go_fuzz_dep_.CoverTab[131289]++
										arr.AppendByteString(bss[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:155
		// _ = "end of CoverTab[131289]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:156
	// _ = "end of CoverTab[131287]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:156
	_go_fuzz_dep_.CoverTab[131288]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:157
	// _ = "end of CoverTab[131288]"
}

type complex128s []complex128

func (nums complex128s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:162
	_go_fuzz_dep_.CoverTab[131290]++
									for i := range nums {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:163
		_go_fuzz_dep_.CoverTab[131292]++
										arr.AppendComplex128(nums[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:164
		// _ = "end of CoverTab[131292]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:165
	// _ = "end of CoverTab[131290]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:165
	_go_fuzz_dep_.CoverTab[131291]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:166
	// _ = "end of CoverTab[131291]"
}

type complex64s []complex64

func (nums complex64s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:171
	_go_fuzz_dep_.CoverTab[131293]++
									for i := range nums {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:172
		_go_fuzz_dep_.CoverTab[131295]++
										arr.AppendComplex64(nums[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:173
		// _ = "end of CoverTab[131295]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:174
	// _ = "end of CoverTab[131293]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:174
	_go_fuzz_dep_.CoverTab[131294]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:175
	// _ = "end of CoverTab[131294]"
}

type durations []time.Duration

func (ds durations) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:180
	_go_fuzz_dep_.CoverTab[131296]++
									for i := range ds {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:181
		_go_fuzz_dep_.CoverTab[131298]++
										arr.AppendDuration(ds[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:182
		// _ = "end of CoverTab[131298]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:183
	// _ = "end of CoverTab[131296]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:183
	_go_fuzz_dep_.CoverTab[131297]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:184
	// _ = "end of CoverTab[131297]"
}

type float64s []float64

func (nums float64s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:189
	_go_fuzz_dep_.CoverTab[131299]++
									for i := range nums {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:190
		_go_fuzz_dep_.CoverTab[131301]++
										arr.AppendFloat64(nums[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:191
		// _ = "end of CoverTab[131301]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:192
	// _ = "end of CoverTab[131299]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:192
	_go_fuzz_dep_.CoverTab[131300]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:193
	// _ = "end of CoverTab[131300]"
}

type float32s []float32

func (nums float32s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:198
	_go_fuzz_dep_.CoverTab[131302]++
									for i := range nums {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:199
		_go_fuzz_dep_.CoverTab[131304]++
										arr.AppendFloat32(nums[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:200
		// _ = "end of CoverTab[131304]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:201
	// _ = "end of CoverTab[131302]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:201
	_go_fuzz_dep_.CoverTab[131303]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:202
	// _ = "end of CoverTab[131303]"
}

type ints []int

func (nums ints) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:207
	_go_fuzz_dep_.CoverTab[131305]++
									for i := range nums {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:208
		_go_fuzz_dep_.CoverTab[131307]++
										arr.AppendInt(nums[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:209
		// _ = "end of CoverTab[131307]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:210
	// _ = "end of CoverTab[131305]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:210
	_go_fuzz_dep_.CoverTab[131306]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:211
	// _ = "end of CoverTab[131306]"
}

type int64s []int64

func (nums int64s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:216
	_go_fuzz_dep_.CoverTab[131308]++
									for i := range nums {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:217
		_go_fuzz_dep_.CoverTab[131310]++
										arr.AppendInt64(nums[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:218
		// _ = "end of CoverTab[131310]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:219
	// _ = "end of CoverTab[131308]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:219
	_go_fuzz_dep_.CoverTab[131309]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:220
	// _ = "end of CoverTab[131309]"
}

type int32s []int32

func (nums int32s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:225
	_go_fuzz_dep_.CoverTab[131311]++
									for i := range nums {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:226
		_go_fuzz_dep_.CoverTab[131313]++
										arr.AppendInt32(nums[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:227
		// _ = "end of CoverTab[131313]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:228
	// _ = "end of CoverTab[131311]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:228
	_go_fuzz_dep_.CoverTab[131312]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:229
	// _ = "end of CoverTab[131312]"
}

type int16s []int16

func (nums int16s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:234
	_go_fuzz_dep_.CoverTab[131314]++
									for i := range nums {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:235
		_go_fuzz_dep_.CoverTab[131316]++
										arr.AppendInt16(nums[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:236
		// _ = "end of CoverTab[131316]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:237
	// _ = "end of CoverTab[131314]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:237
	_go_fuzz_dep_.CoverTab[131315]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:238
	// _ = "end of CoverTab[131315]"
}

type int8s []int8

func (nums int8s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:243
	_go_fuzz_dep_.CoverTab[131317]++
									for i := range nums {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:244
		_go_fuzz_dep_.CoverTab[131319]++
										arr.AppendInt8(nums[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:245
		// _ = "end of CoverTab[131319]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:246
	// _ = "end of CoverTab[131317]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:246
	_go_fuzz_dep_.CoverTab[131318]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:247
	// _ = "end of CoverTab[131318]"
}

type stringArray []string

func (ss stringArray) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:252
	_go_fuzz_dep_.CoverTab[131320]++
									for i := range ss {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:253
		_go_fuzz_dep_.CoverTab[131322]++
										arr.AppendString(ss[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:254
		// _ = "end of CoverTab[131322]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:255
	// _ = "end of CoverTab[131320]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:255
	_go_fuzz_dep_.CoverTab[131321]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:256
	// _ = "end of CoverTab[131321]"
}

type times []time.Time

func (ts times) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:261
	_go_fuzz_dep_.CoverTab[131323]++
									for i := range ts {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:262
		_go_fuzz_dep_.CoverTab[131325]++
										arr.AppendTime(ts[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:263
		// _ = "end of CoverTab[131325]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:264
	// _ = "end of CoverTab[131323]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:264
	_go_fuzz_dep_.CoverTab[131324]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:265
	// _ = "end of CoverTab[131324]"
}

type uints []uint

func (nums uints) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:270
	_go_fuzz_dep_.CoverTab[131326]++
									for i := range nums {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:271
		_go_fuzz_dep_.CoverTab[131328]++
										arr.AppendUint(nums[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:272
		// _ = "end of CoverTab[131328]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:273
	// _ = "end of CoverTab[131326]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:273
	_go_fuzz_dep_.CoverTab[131327]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:274
	// _ = "end of CoverTab[131327]"
}

type uint64s []uint64

func (nums uint64s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:279
	_go_fuzz_dep_.CoverTab[131329]++
									for i := range nums {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:280
		_go_fuzz_dep_.CoverTab[131331]++
										arr.AppendUint64(nums[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:281
		// _ = "end of CoverTab[131331]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:282
	// _ = "end of CoverTab[131329]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:282
	_go_fuzz_dep_.CoverTab[131330]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:283
	// _ = "end of CoverTab[131330]"
}

type uint32s []uint32

func (nums uint32s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:288
	_go_fuzz_dep_.CoverTab[131332]++
									for i := range nums {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:289
		_go_fuzz_dep_.CoverTab[131334]++
										arr.AppendUint32(nums[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:290
		// _ = "end of CoverTab[131334]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:291
	// _ = "end of CoverTab[131332]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:291
	_go_fuzz_dep_.CoverTab[131333]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:292
	// _ = "end of CoverTab[131333]"
}

type uint16s []uint16

func (nums uint16s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:297
	_go_fuzz_dep_.CoverTab[131335]++
									for i := range nums {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:298
		_go_fuzz_dep_.CoverTab[131337]++
										arr.AppendUint16(nums[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:299
		// _ = "end of CoverTab[131337]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:300
	// _ = "end of CoverTab[131335]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:300
	_go_fuzz_dep_.CoverTab[131336]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:301
	// _ = "end of CoverTab[131336]"
}

type uint8s []uint8

func (nums uint8s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:306
	_go_fuzz_dep_.CoverTab[131338]++
									for i := range nums {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:307
		_go_fuzz_dep_.CoverTab[131340]++
										arr.AppendUint8(nums[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:308
		// _ = "end of CoverTab[131340]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:309
	// _ = "end of CoverTab[131338]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:309
	_go_fuzz_dep_.CoverTab[131339]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:310
	// _ = "end of CoverTab[131339]"
}

type uintptrs []uintptr

func (nums uintptrs) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:315
	_go_fuzz_dep_.CoverTab[131341]++
									for i := range nums {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:316
		_go_fuzz_dep_.CoverTab[131343]++
										arr.AppendUintptr(nums[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:317
		// _ = "end of CoverTab[131343]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:318
	// _ = "end of CoverTab[131341]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:318
	_go_fuzz_dep_.CoverTab[131342]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:319
	// _ = "end of CoverTab[131342]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:320
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/array.go:320
var _ = _go_fuzz_dep_.CoverTab
