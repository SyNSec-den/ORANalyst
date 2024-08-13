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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:21
// Package exit provides stubs so that unit tests can exercise code that calls
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:21
// os.Exit(1).
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:23
package exit

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:23
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:23
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:23
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:23
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:23
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:23
)

import "os"

var real = func() { _go_fuzz_dep_.CoverTab[130601]++; os.Exit(1); // _ = "end of CoverTab[130601]" }

// Exit normally terminates the process by calling os.Exit(1). If the package
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:29
// is stubbed, it instead records a call in the testing spy.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:31
func Exit() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:31
	_go_fuzz_dep_.CoverTab[130602]++
											real()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:32
	// _ = "end of CoverTab[130602]"
}

// A StubbedExit is a testing fake for os.Exit.
type StubbedExit struct {
	Exited	bool
	prev	func()
}

// Stub substitutes a fake for the call to os.Exit(1).
func Stub() *StubbedExit {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:42
	_go_fuzz_dep_.CoverTab[130603]++
											s := &StubbedExit{prev: real}
											real = s.exit
											return s
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:45
	// _ = "end of CoverTab[130603]"
}

// WithStub runs the supplied function with Exit stubbed. It returns the stub
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:48
// used, so that users can test whether the process would have crashed.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:50
func WithStub(f func()) *StubbedExit {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:50
	_go_fuzz_dep_.CoverTab[130604]++
											s := Stub()
											defer s.Unstub()
											f()
											return s
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:54
	// _ = "end of CoverTab[130604]"
}

// Unstub restores the previous exit function.
func (se *StubbedExit) Unstub() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:58
	_go_fuzz_dep_.CoverTab[130605]++
											real = se.prev
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:59
	// _ = "end of CoverTab[130605]"
}

func (se *StubbedExit) exit() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:62
	_go_fuzz_dep_.CoverTab[130606]++
											se.Exited = true
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:63
	// _ = "end of CoverTab[130606]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:64
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/internal/exit/exit.go:64
var _ = _go_fuzz_dep_.CoverTab
