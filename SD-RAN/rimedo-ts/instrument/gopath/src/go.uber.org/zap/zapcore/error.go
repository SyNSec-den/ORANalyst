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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:21
package zapcore

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:21
)

import (
	"fmt"
	"reflect"
	"sync"
)

// Encodes the given error into fields of an object. A field with the given
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
// name is added for the error message.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
// If the error implements fmt.Formatter, a field with the name ${key}Verbose
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
// is also added with the full verbose error message.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
// Finally, if the error implements errorGroup (from go.uber.org/multierr) or
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
// causer (from github.com/pkg/errors), a ${key}Causes field is added with an
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
// array of objects containing the errors this error was comprised of.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
//	{
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
//	  "error": err.Error(),
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
//	  "errorVerbose": fmt.Sprintf("%+v", err),
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
//	  "errorCauses": [
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
//	    ...
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
//	  ],
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:29
//	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:46
func encodeError(key string, err error, enc ObjectEncoder) (retErr error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:46
	_go_fuzz_dep_.CoverTab[130811]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:49
	defer func() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:49
		_go_fuzz_dep_.CoverTab[130814]++
											if rerr := recover(); rerr != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:50
			_go_fuzz_dep_.CoverTab[130815]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:54
			if v := reflect.ValueOf(err); v.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:54
				_go_fuzz_dep_.CoverTab[130817]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:54
				return v.IsNil()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:54
				// _ = "end of CoverTab[130817]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:54
			}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:54
				_go_fuzz_dep_.CoverTab[130818]++
													enc.AddString(key, "<nil>")
													return
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:56
				// _ = "end of CoverTab[130818]"
			} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:57
				_go_fuzz_dep_.CoverTab[130819]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:57
				// _ = "end of CoverTab[130819]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:57
			}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:57
			// _ = "end of CoverTab[130815]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:57
			_go_fuzz_dep_.CoverTab[130816]++

												retErr = fmt.Errorf("PANIC=%v", rerr)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:59
			// _ = "end of CoverTab[130816]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:60
			_go_fuzz_dep_.CoverTab[130820]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:60
			// _ = "end of CoverTab[130820]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:60
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:60
		// _ = "end of CoverTab[130814]"
	}()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:61
	// _ = "end of CoverTab[130811]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:61
	_go_fuzz_dep_.CoverTab[130812]++

										basic := err.Error()
										enc.AddString(key, basic)

										switch e := err.(type) {
	case errorGroup:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:67
		_go_fuzz_dep_.CoverTab[130821]++
											return enc.AddArray(key+"Causes", errArray(e.Errors()))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:68
		// _ = "end of CoverTab[130821]"
	case fmt.Formatter:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:69
		_go_fuzz_dep_.CoverTab[130822]++
											verbose := fmt.Sprintf("%+v", e)
											if verbose != basic {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:71
			_go_fuzz_dep_.CoverTab[130823]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:74
			enc.AddString(key+"Verbose", verbose)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:74
			// _ = "end of CoverTab[130823]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:75
			_go_fuzz_dep_.CoverTab[130824]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:75
			// _ = "end of CoverTab[130824]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:75
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:75
		// _ = "end of CoverTab[130822]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:76
	// _ = "end of CoverTab[130812]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:76
	_go_fuzz_dep_.CoverTab[130813]++
										return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:77
	// _ = "end of CoverTab[130813]"
}

type errorGroup interface {
	// Provides read-only access to the underlying list of errors, preferably
	// without causing any allocs.
	Errors() []error
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:90
// Encodes a list of errors using the standard error encoding logic.
type errArray []error

func (errs errArray) MarshalLogArray(arr ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:93
	_go_fuzz_dep_.CoverTab[130825]++
										for i := range errs {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:94
		_go_fuzz_dep_.CoverTab[130827]++
											if errs[i] == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:95
			_go_fuzz_dep_.CoverTab[130829]++
												continue
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:96
			// _ = "end of CoverTab[130829]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:97
			_go_fuzz_dep_.CoverTab[130830]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:97
			// _ = "end of CoverTab[130830]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:97
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:97
		// _ = "end of CoverTab[130827]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:97
		_go_fuzz_dep_.CoverTab[130828]++

											el := newErrArrayElem(errs[i])
											arr.AppendObject(el)
											el.Free()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:101
		// _ = "end of CoverTab[130828]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:102
	// _ = "end of CoverTab[130825]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:102
	_go_fuzz_dep_.CoverTab[130826]++
										return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:103
	// _ = "end of CoverTab[130826]"
}

var _errArrayElemPool = sync.Pool{New: func() interface{} {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:106
	_go_fuzz_dep_.CoverTab[130831]++
										return &errArrayElem{}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:107
	// _ = "end of CoverTab[130831]"
}}

// Encodes any error into a {"error": ...} re-using the same errors logic.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:110
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:110
// May be passed in place of an array to build a single-element array.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:113
type errArrayElem struct{ err error }

func newErrArrayElem(err error) *errArrayElem {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:115
	_go_fuzz_dep_.CoverTab[130832]++
										e := _errArrayElemPool.Get().(*errArrayElem)
										e.err = err
										return e
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:118
	// _ = "end of CoverTab[130832]"
}

func (e *errArrayElem) MarshalLogArray(arr ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:121
	_go_fuzz_dep_.CoverTab[130833]++
										return arr.AppendObject(e)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:122
	// _ = "end of CoverTab[130833]"
}

func (e *errArrayElem) MarshalLogObject(enc ObjectEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:125
	_go_fuzz_dep_.CoverTab[130834]++
										return encodeError("error", e.err, enc)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:126
	// _ = "end of CoverTab[130834]"
}

func (e *errArrayElem) Free() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:129
	_go_fuzz_dep_.CoverTab[130835]++
										e.err = nil
										_errArrayElemPool.Put(e)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:131
	// _ = "end of CoverTab[130835]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:132
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/error.go:132
var _ = _go_fuzz_dep_.CoverTab
