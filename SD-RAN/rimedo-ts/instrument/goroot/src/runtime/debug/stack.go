// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/runtime/debug/stack.go:5
// Package debug contains facilities for programs to debug themselves while
//line /usr/local/go/src/runtime/debug/stack.go:5
// they are running.
//line /usr/local/go/src/runtime/debug/stack.go:7
package debug

//line /usr/local/go/src/runtime/debug/stack.go:7
import (
//line /usr/local/go/src/runtime/debug/stack.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/runtime/debug/stack.go:7
)
//line /usr/local/go/src/runtime/debug/stack.go:7
import (
//line /usr/local/go/src/runtime/debug/stack.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/runtime/debug/stack.go:7
)

import (
	"os"
	"runtime"
)

// PrintStack prints to standard error the stack trace returned by runtime.Stack.
func PrintStack() {
//line /usr/local/go/src/runtime/debug/stack.go:15
	_go_fuzz_dep_.CoverTab[90835]++
							os.Stderr.Write(Stack())
//line /usr/local/go/src/runtime/debug/stack.go:16
	// _ = "end of CoverTab[90835]"
}

// Stack returns a formatted stack trace of the goroutine that calls it.
//line /usr/local/go/src/runtime/debug/stack.go:19
// It calls runtime.Stack with a large enough buffer to capture the entire trace.
//line /usr/local/go/src/runtime/debug/stack.go:21
func Stack() []byte {
//line /usr/local/go/src/runtime/debug/stack.go:21
	_go_fuzz_dep_.CoverTab[90836]++
							buf := make([]byte, 1024)
							for {
//line /usr/local/go/src/runtime/debug/stack.go:23
		_go_fuzz_dep_.CoverTab[90837]++
								n := runtime.Stack(buf, false)
								if n < len(buf) {
//line /usr/local/go/src/runtime/debug/stack.go:25
			_go_fuzz_dep_.CoverTab[90839]++
									return buf[:n]
//line /usr/local/go/src/runtime/debug/stack.go:26
			// _ = "end of CoverTab[90839]"
		} else {
//line /usr/local/go/src/runtime/debug/stack.go:27
			_go_fuzz_dep_.CoverTab[90840]++
//line /usr/local/go/src/runtime/debug/stack.go:27
			// _ = "end of CoverTab[90840]"
//line /usr/local/go/src/runtime/debug/stack.go:27
		}
//line /usr/local/go/src/runtime/debug/stack.go:27
		// _ = "end of CoverTab[90837]"
//line /usr/local/go/src/runtime/debug/stack.go:27
		_go_fuzz_dep_.CoverTab[90838]++
								buf = make([]byte, 2*len(buf))
//line /usr/local/go/src/runtime/debug/stack.go:28
		// _ = "end of CoverTab[90838]"
	}
//line /usr/local/go/src/runtime/debug/stack.go:29
	// _ = "end of CoverTab[90836]"
}

//line /usr/local/go/src/runtime/debug/stack.go:30
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/runtime/debug/stack.go:30
var _ = _go_fuzz_dep_.CoverTab
