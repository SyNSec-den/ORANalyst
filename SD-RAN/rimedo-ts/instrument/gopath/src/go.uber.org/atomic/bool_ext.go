// Copyright (c) 2020 Uber Technologies, Inc.
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

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:21
package atomic

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:21
)

import (
	"strconv"
)

//go:generate bin/gen-atomicwrapper -name=Bool -type=bool -wrapped=Uint32 -pack=boolToInt -unpack=truthy -cas -swap -json -file=bool.go

func truthy(n uint32) bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:29
	_go_fuzz_dep_.CoverTab[130358]++
										return n == 1
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:30
	// _ = "end of CoverTab[130358]"
}

func boolToInt(b bool) uint32 {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:33
	_go_fuzz_dep_.CoverTab[130359]++
										if b {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:34
		_go_fuzz_dep_.CoverTab[130361]++
											return 1
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:35
		// _ = "end of CoverTab[130361]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:36
		_go_fuzz_dep_.CoverTab[130362]++
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:36
		// _ = "end of CoverTab[130362]"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:36
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:36
	// _ = "end of CoverTab[130359]"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:36
	_go_fuzz_dep_.CoverTab[130360]++
										return 0
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:37
	// _ = "end of CoverTab[130360]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:41
func (b *Bool) Toggle() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:41
	_go_fuzz_dep_.CoverTab[130363]++
										for {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:42
		_go_fuzz_dep_.CoverTab[130364]++
											old := b.Load()
											if b.CAS(old, !old) {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:44
			_go_fuzz_dep_.CoverTab[130365]++
												return old
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:45
			// _ = "end of CoverTab[130365]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:46
			_go_fuzz_dep_.CoverTab[130366]++
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:46
			// _ = "end of CoverTab[130366]"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:46
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:46
		// _ = "end of CoverTab[130364]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:47
	// _ = "end of CoverTab[130363]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:51
func (b *Bool) String() string {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:51
	_go_fuzz_dep_.CoverTab[130367]++
										return strconv.FormatBool(b.Load())
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:52
	// _ = "end of CoverTab[130367]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:53
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool_ext.go:53
var _ = _go_fuzz_dep_.CoverTab
