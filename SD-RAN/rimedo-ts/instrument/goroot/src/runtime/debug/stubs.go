// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/runtime/debug/stubs.go:5
package debug

//line /usr/local/go/src/runtime/debug/stubs.go:5
import (
//line /usr/local/go/src/runtime/debug/stubs.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/runtime/debug/stubs.go:5
)
//line /usr/local/go/src/runtime/debug/stubs.go:5
import (
//line /usr/local/go/src/runtime/debug/stubs.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/runtime/debug/stubs.go:5
)

import (
	"time"
)

// Implemented in package runtime.
func readGCStats(*[]time.Duration)
func freeOSMemory()
func setMaxStack(int) int
func setGCPercent(int32) int32
func setPanicOnFault(bool) bool
func setMaxThreads(int) int
func setMemoryLimit(int64) int64

//line /usr/local/go/src/runtime/debug/stubs.go:18
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/runtime/debug/stubs.go:18
var _ = _go_fuzz_dep_.CoverTab
