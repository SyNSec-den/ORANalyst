// Copyright (c) 2017 Uber Technologies, Inc.
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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:21
)

import (
	"sync"

	"go.uber.org/zap/zapcore"
)

var _errArrayElemPool = sync.Pool{New: func() interface{} {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:29
	_go_fuzz_dep_.CoverTab[131416]++
									return &errArrayElem{}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:30
	// _ = "end of CoverTab[131416]"
}}

// Error is shorthand for the common idiom NamedError("error", err).
func Error(err error) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:34
	_go_fuzz_dep_.CoverTab[131417]++
									return NamedError("error", err)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:35
	// _ = "end of CoverTab[131417]"
}

// NamedError constructs a field that lazily stores err.Error() under the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:38
// provided key. Errors which also implement fmt.Formatter (like those produced
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:38
// by github.com/pkg/errors) will also have their verbose representation stored
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:38
// under key+"Verbose". If passed a nil error, the field is a no-op.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:38
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:38
// For the common case in which the key is simply "error", the Error function
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:38
// is shorter and less repetitive.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:45
func NamedError(key string, err error) Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:45
	_go_fuzz_dep_.CoverTab[131418]++
									if err == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:46
		_go_fuzz_dep_.CoverTab[131420]++
										return Skip()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:47
		// _ = "end of CoverTab[131420]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:48
		_go_fuzz_dep_.CoverTab[131421]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:48
		// _ = "end of CoverTab[131421]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:48
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:48
	// _ = "end of CoverTab[131418]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:48
	_go_fuzz_dep_.CoverTab[131419]++
									return Field{Key: key, Type: zapcore.ErrorType, Interface: err}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:49
	// _ = "end of CoverTab[131419]"
}

type errArray []error

func (errs errArray) MarshalLogArray(arr zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:54
	_go_fuzz_dep_.CoverTab[131422]++
									for i := range errs {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:55
		_go_fuzz_dep_.CoverTab[131424]++
										if errs[i] == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:56
			_go_fuzz_dep_.CoverTab[131426]++
											continue
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:57
			// _ = "end of CoverTab[131426]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:58
			_go_fuzz_dep_.CoverTab[131427]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:58
			// _ = "end of CoverTab[131427]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:58
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:58
		// _ = "end of CoverTab[131424]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:58
		_go_fuzz_dep_.CoverTab[131425]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:63
		elem := _errArrayElemPool.Get().(*errArrayElem)
										elem.error = errs[i]
										arr.AppendObject(elem)
										elem.error = nil
										_errArrayElemPool.Put(elem)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:67
		// _ = "end of CoverTab[131425]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:68
	// _ = "end of CoverTab[131422]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:68
	_go_fuzz_dep_.CoverTab[131423]++
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:69
	// _ = "end of CoverTab[131423]"
}

type errArrayElem struct {
	error
}

func (e *errArrayElem) MarshalLogObject(enc zapcore.ObjectEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:76
	_go_fuzz_dep_.CoverTab[131428]++

									Error(e.error).AddTo(enc)
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:79
	// _ = "end of CoverTab[131428]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:80
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/error.go:80
var _ = _go_fuzz_dep_.CoverTab
