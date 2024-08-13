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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:21
package zapcore

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:21
)

import "go.uber.org/multierr"

type multiCore []Core

// NewTee creates a Core that duplicates log entries into two or more
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:27
// underlying Cores.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:27
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:27
// Calling it with a single Core returns the input unchanged, and calling
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:27
// it with no input returns a no-op Core.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:32
func NewTee(cores ...Core) Core {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:32
	_go_fuzz_dep_.CoverTab[131216]++
										switch len(cores) {
	case 0:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:34
		_go_fuzz_dep_.CoverTab[131217]++
											return NewNopCore()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:35
		// _ = "end of CoverTab[131217]"
	case 1:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:36
		_go_fuzz_dep_.CoverTab[131218]++
											return cores[0]
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:37
		// _ = "end of CoverTab[131218]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:38
		_go_fuzz_dep_.CoverTab[131219]++
											return multiCore(cores)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:39
		// _ = "end of CoverTab[131219]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:40
	// _ = "end of CoverTab[131216]"
}

func (mc multiCore) With(fields []Field) Core {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:43
	_go_fuzz_dep_.CoverTab[131220]++
										clone := make(multiCore, len(mc))
										for i := range mc {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:45
		_go_fuzz_dep_.CoverTab[131222]++
											clone[i] = mc[i].With(fields)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:46
		// _ = "end of CoverTab[131222]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:47
	// _ = "end of CoverTab[131220]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:47
	_go_fuzz_dep_.CoverTab[131221]++
										return clone
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:48
	// _ = "end of CoverTab[131221]"
}

func (mc multiCore) Enabled(lvl Level) bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:51
	_go_fuzz_dep_.CoverTab[131223]++
										for i := range mc {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:52
		_go_fuzz_dep_.CoverTab[131225]++
											if mc[i].Enabled(lvl) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:53
			_go_fuzz_dep_.CoverTab[131226]++
												return true
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:54
			// _ = "end of CoverTab[131226]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:55
			_go_fuzz_dep_.CoverTab[131227]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:55
			// _ = "end of CoverTab[131227]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:55
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:55
		// _ = "end of CoverTab[131225]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:56
	// _ = "end of CoverTab[131223]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:56
	_go_fuzz_dep_.CoverTab[131224]++
										return false
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:57
	// _ = "end of CoverTab[131224]"
}

func (mc multiCore) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:60
	_go_fuzz_dep_.CoverTab[131228]++
										for i := range mc {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:61
		_go_fuzz_dep_.CoverTab[131230]++
											ce = mc[i].Check(ent, ce)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:62
		// _ = "end of CoverTab[131230]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:63
	// _ = "end of CoverTab[131228]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:63
	_go_fuzz_dep_.CoverTab[131229]++
										return ce
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:64
	// _ = "end of CoverTab[131229]"
}

func (mc multiCore) Write(ent Entry, fields []Field) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:67
	_go_fuzz_dep_.CoverTab[131231]++
										var err error
										for i := range mc {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:69
		_go_fuzz_dep_.CoverTab[131233]++
											err = multierr.Append(err, mc[i].Write(ent, fields))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:70
		// _ = "end of CoverTab[131233]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:71
	// _ = "end of CoverTab[131231]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:71
	_go_fuzz_dep_.CoverTab[131232]++
										return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:72
	// _ = "end of CoverTab[131232]"
}

func (mc multiCore) Sync() error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:75
	_go_fuzz_dep_.CoverTab[131234]++
										var err error
										for i := range mc {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:77
		_go_fuzz_dep_.CoverTab[131236]++
											err = multierr.Append(err, mc[i].Sync())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:78
		// _ = "end of CoverTab[131236]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:79
	// _ = "end of CoverTab[131234]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:79
	_go_fuzz_dep_.CoverTab[131235]++
										return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:80
	// _ = "end of CoverTab[131235]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:81
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/tee.go:81
var _ = _go_fuzz_dep_.CoverTab
