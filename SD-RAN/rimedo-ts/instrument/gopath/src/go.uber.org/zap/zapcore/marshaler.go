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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:21
package zapcore

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:21
)

// ObjectMarshaler allows user-defined types to efficiently add themselves to the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:23
// logging context, and to selectively omit information which shouldn't be
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:23
// included in logs (e.g., passwords).
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:23
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:23
// Note: ObjectMarshaler is only used when zap.Object is used or when
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:23
// passed directly to zap.Any. It is not used when reflection-based
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:23
// encoding is used.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:30
type ObjectMarshaler interface {
	MarshalLogObject(ObjectEncoder) error
}

// ObjectMarshalerFunc is a type adapter that turns a function into an
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:34
// ObjectMarshaler.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:36
type ObjectMarshalerFunc func(ObjectEncoder) error

// MarshalLogObject calls the underlying function.
func (f ObjectMarshalerFunc) MarshalLogObject(enc ObjectEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:39
	_go_fuzz_dep_.CoverTab[131136]++
											return f(enc)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:40
	// _ = "end of CoverTab[131136]"
}

// ArrayMarshaler allows user-defined types to efficiently add themselves to the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:43
// logging context, and to selectively omit information which shouldn't be
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:43
// included in logs (e.g., passwords).
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:43
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:43
// Note: ArrayMarshaler is only used when zap.Array is used or when
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:43
// passed directly to zap.Any. It is not used when reflection-based
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:43
// encoding is used.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:50
type ArrayMarshaler interface {
	MarshalLogArray(ArrayEncoder) error
}

// ArrayMarshalerFunc is a type adapter that turns a function into an
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:54
// ArrayMarshaler.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:56
type ArrayMarshalerFunc func(ArrayEncoder) error

// MarshalLogArray calls the underlying function.
func (f ArrayMarshalerFunc) MarshalLogArray(enc ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:59
	_go_fuzz_dep_.CoverTab[131137]++
											return f(enc)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:60
	// _ = "end of CoverTab[131137]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:61
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/marshaler.go:61
var _ = _go_fuzz_dep_.CoverTab
