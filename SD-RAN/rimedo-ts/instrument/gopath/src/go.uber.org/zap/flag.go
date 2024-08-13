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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:21
)

import (
	"flag"

	"go.uber.org/zap/zapcore"
)

// LevelFlag uses the standard library's flag.Var to declare a global flag
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:29
// with the specified name, default, and usage guidance. The returned value is
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:29
// a pointer to the value of the flag.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:29
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:29
// If you don't want to use the flag package's global state, you can use any
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:29
// non-nil *Level as a flag.Value with your own *flag.FlagSet.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:35
func LevelFlag(name string, defaultLevel zapcore.Level, usage string) *zapcore.Level {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:35
	_go_fuzz_dep_.CoverTab[131606]++
									lvl := defaultLevel
									flag.Var(&lvl, name, usage)
									return &lvl
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:38
	// _ = "end of CoverTab[131606]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:39
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/flag.go:39
var _ = _go_fuzz_dep_.CoverTab
