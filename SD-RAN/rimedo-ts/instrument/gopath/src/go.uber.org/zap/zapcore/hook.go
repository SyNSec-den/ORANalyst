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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:21
package zapcore

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:21
)

import "go.uber.org/multierr"

type hooked struct {
	Core
	funcs	[]func(Entry) error
}

// RegisterHooks wraps a Core and runs a collection of user-defined callback
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:30
// hooks each time a message is logged. Execution of the callbacks is blocking.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:30
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:30
// This offers users an easy way to register simple callbacks (e.g., metrics
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:30
// collection) without implementing the full Core interface.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:35
func RegisterHooks(core Core, hooks ...func(Entry) error) Core {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:35
	_go_fuzz_dep_.CoverTab[130892]++
										funcs := append([]func(Entry) error{}, hooks...)
										return &hooked{
		Core:	core,
		funcs:	funcs,
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:40
	// _ = "end of CoverTab[130892]"
}

func (h *hooked) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:43
	_go_fuzz_dep_.CoverTab[130893]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:47
	if downstream := h.Core.Check(ent, ce); downstream != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:47
		_go_fuzz_dep_.CoverTab[130895]++
											return downstream.AddCore(ent, h)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:48
		// _ = "end of CoverTab[130895]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:49
		_go_fuzz_dep_.CoverTab[130896]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:49
		// _ = "end of CoverTab[130896]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:49
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:49
	// _ = "end of CoverTab[130893]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:49
	_go_fuzz_dep_.CoverTab[130894]++
										return ce
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:50
	// _ = "end of CoverTab[130894]"
}

func (h *hooked) With(fields []Field) Core {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:53
	_go_fuzz_dep_.CoverTab[130897]++
										return &hooked{
		Core:	h.Core.With(fields),
		funcs:	h.funcs,
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:57
	// _ = "end of CoverTab[130897]"
}

func (h *hooked) Write(ent Entry, _ []Field) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:60
	_go_fuzz_dep_.CoverTab[130898]++
	// Since our downstream had a chance to register itself directly with the
	// CheckedMessage, we don't need to call it here.
	var err error
	for i := range h.funcs {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:64
		_go_fuzz_dep_.CoverTab[130900]++
											err = multierr.Append(err, h.funcs[i](ent))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:65
		// _ = "end of CoverTab[130900]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:66
	// _ = "end of CoverTab[130898]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:66
	_go_fuzz_dep_.CoverTab[130899]++
										return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:67
	// _ = "end of CoverTab[130899]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:68
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/hook.go:68
var _ = _go_fuzz_dep_.CoverTab
