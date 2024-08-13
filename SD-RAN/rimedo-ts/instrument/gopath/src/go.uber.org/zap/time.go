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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/time.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/time.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/time.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/time.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/time.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/time.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/time.go:21
)

import "time"

func timeToMillis(t time.Time) int64 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/time.go:25
	_go_fuzz_dep_.CoverTab[131922]++
									return t.UnixNano() / int64(time.Millisecond)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/time.go:26
	// _ = "end of CoverTab[131922]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/time.go:27
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/time.go:27
var _ = _go_fuzz_dep_.CoverTab
