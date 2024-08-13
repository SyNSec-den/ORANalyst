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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:21
)

import (
	"runtime"
	"sync"

	"go.uber.org/zap/internal/bufferpool"
)

var (
	_stacktracePool = sync.Pool{
		New: func() interface{} {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:32
			_go_fuzz_dep_.CoverTab[131840]++
												return newProgramCounters(64)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:33
			// _ = "end of CoverTab[131840]"
		},
	}
)

func takeStacktrace(skip int) string {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:38
	_go_fuzz_dep_.CoverTab[131841]++
										buffer := bufferpool.Get()
										defer buffer.Free()
										programCounters := _stacktracePool.Get().(*programCounters)
										defer _stacktracePool.Put(programCounters)

										var numFrames int
										for {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:45
		_go_fuzz_dep_.CoverTab[131844]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:48
		numFrames = runtime.Callers(skip+2, programCounters.pcs)
		if numFrames < len(programCounters.pcs) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:49
			_go_fuzz_dep_.CoverTab[131846]++
												break
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:50
			// _ = "end of CoverTab[131846]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:51
			_go_fuzz_dep_.CoverTab[131847]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:51
			// _ = "end of CoverTab[131847]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:51
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:51
		// _ = "end of CoverTab[131844]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:51
		_go_fuzz_dep_.CoverTab[131845]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:54
		programCounters = newProgramCounters(len(programCounters.pcs) * 2)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:54
		// _ = "end of CoverTab[131845]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:55
	// _ = "end of CoverTab[131841]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:55
	_go_fuzz_dep_.CoverTab[131842]++

										i := 0
										frames := runtime.CallersFrames(programCounters.pcs[:numFrames])

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:63
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:63
		_go_fuzz_dep_.CoverTab[131848]++
											if i != 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:64
			_go_fuzz_dep_.CoverTab[131850]++
												buffer.AppendByte('\n')
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:65
			// _ = "end of CoverTab[131850]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:66
			_go_fuzz_dep_.CoverTab[131851]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:66
			// _ = "end of CoverTab[131851]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:66
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:66
		// _ = "end of CoverTab[131848]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:66
		_go_fuzz_dep_.CoverTab[131849]++
											i++
											buffer.AppendString(frame.Function)
											buffer.AppendByte('\n')
											buffer.AppendByte('\t')
											buffer.AppendString(frame.File)
											buffer.AppendByte(':')
											buffer.AppendInt(int64(frame.Line))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:73
		// _ = "end of CoverTab[131849]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:74
	// _ = "end of CoverTab[131842]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:74
	_go_fuzz_dep_.CoverTab[131843]++

										return buffer.String()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:76
	// _ = "end of CoverTab[131843]"
}

type programCounters struct {
	pcs []uintptr
}

func newProgramCounters(size int) *programCounters {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:83
	_go_fuzz_dep_.CoverTab[131852]++
										return &programCounters{make([]uintptr, size)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:84
	// _ = "end of CoverTab[131852]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:85
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/stacktrace.go:85
var _ = _go_fuzz_dep_.CoverTab
