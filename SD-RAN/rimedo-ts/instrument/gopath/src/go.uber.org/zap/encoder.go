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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:21
)

import (
	"errors"
	"fmt"
	"sync"

	"go.uber.org/zap/zapcore"
)

var (
	errNoEncoderNameSpecified	= errors.New("no encoder name specified")

	_encoderNameToConstructor	= map[string]func(zapcore.EncoderConfig) (zapcore.Encoder, error){
		"console": func(encoderConfig zapcore.EncoderConfig) (zapcore.Encoder, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:35
			_go_fuzz_dep_.CoverTab[131396]++
												return zapcore.NewConsoleEncoder(encoderConfig), nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:36
			// _ = "end of CoverTab[131396]"
		},
		"json": func(encoderConfig zapcore.EncoderConfig) (zapcore.Encoder, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:38
			_go_fuzz_dep_.CoverTab[131397]++
												return zapcore.NewJSONEncoder(encoderConfig), nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:39
			// _ = "end of CoverTab[131397]"
		},
	}
	_encoderMutex	sync.RWMutex
)

// RegisterEncoder registers an encoder constructor, which the Config struct
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:45
// can then reference. By default, the "json" and "console" encoders are
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:45
// registered.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:45
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:45
// Attempting to register an encoder whose name is already taken returns an
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:45
// error.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:51
func RegisterEncoder(name string, constructor func(zapcore.EncoderConfig) (zapcore.Encoder, error)) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:51
	_go_fuzz_dep_.CoverTab[131398]++
										_encoderMutex.Lock()
										defer _encoderMutex.Unlock()
										if name == "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:54
		_go_fuzz_dep_.CoverTab[131401]++
											return errNoEncoderNameSpecified
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:55
		// _ = "end of CoverTab[131401]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:56
		_go_fuzz_dep_.CoverTab[131402]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:56
		// _ = "end of CoverTab[131402]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:56
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:56
	// _ = "end of CoverTab[131398]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:56
	_go_fuzz_dep_.CoverTab[131399]++
										if _, ok := _encoderNameToConstructor[name]; ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:57
		_go_fuzz_dep_.CoverTab[131403]++
											return fmt.Errorf("encoder already registered for name %q", name)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:58
		// _ = "end of CoverTab[131403]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:59
		_go_fuzz_dep_.CoverTab[131404]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:59
		// _ = "end of CoverTab[131404]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:59
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:59
	// _ = "end of CoverTab[131399]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:59
	_go_fuzz_dep_.CoverTab[131400]++
										_encoderNameToConstructor[name] = constructor
										return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:61
	// _ = "end of CoverTab[131400]"
}

func newEncoder(name string, encoderConfig zapcore.EncoderConfig) (zapcore.Encoder, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:64
	_go_fuzz_dep_.CoverTab[131405]++
										if encoderConfig.TimeKey != "" && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:65
		_go_fuzz_dep_.CoverTab[131409]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:65
		return encoderConfig.EncodeTime == nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:65
		// _ = "end of CoverTab[131409]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:65
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:65
		_go_fuzz_dep_.CoverTab[131410]++
											return nil, fmt.Errorf("missing EncodeTime in EncoderConfig")
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:66
		// _ = "end of CoverTab[131410]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:67
		_go_fuzz_dep_.CoverTab[131411]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:67
		// _ = "end of CoverTab[131411]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:67
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:67
	// _ = "end of CoverTab[131405]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:67
	_go_fuzz_dep_.CoverTab[131406]++

										_encoderMutex.RLock()
										defer _encoderMutex.RUnlock()
										if name == "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:71
		_go_fuzz_dep_.CoverTab[131412]++
											return nil, errNoEncoderNameSpecified
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:72
		// _ = "end of CoverTab[131412]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:73
		_go_fuzz_dep_.CoverTab[131413]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:73
		// _ = "end of CoverTab[131413]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:73
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:73
	// _ = "end of CoverTab[131406]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:73
	_go_fuzz_dep_.CoverTab[131407]++
										constructor, ok := _encoderNameToConstructor[name]
										if !ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:75
		_go_fuzz_dep_.CoverTab[131414]++
											return nil, fmt.Errorf("no encoder registered for name %q", name)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:76
		// _ = "end of CoverTab[131414]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:77
		_go_fuzz_dep_.CoverTab[131415]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:77
		// _ = "end of CoverTab[131415]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:77
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:77
	// _ = "end of CoverTab[131407]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:77
	_go_fuzz_dep_.CoverTab[131408]++
										return constructor(encoderConfig)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:78
	// _ = "end of CoverTab[131408]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:79
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/encoder.go:79
var _ = _go_fuzz_dep_.CoverTab
