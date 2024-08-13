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

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:21
package atomic

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:21
)

import "strconv"

//go:generate bin/gen-atomicwrapper -name=Float64 -type=float64 -wrapped=Uint64 -pack=math.Float64bits -unpack=math.Float64frombits -cas -json -imports math -file=float64.go

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:28
func (f *Float64) Add(s float64) float64 {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:28
	_go_fuzz_dep_.CoverTab[130407]++
										for {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:29
		_go_fuzz_dep_.CoverTab[130408]++
											old := f.Load()
											new := old + s
											if f.CAS(old, new) {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:32
			_go_fuzz_dep_.CoverTab[130409]++
												return new
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:33
			// _ = "end of CoverTab[130409]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:34
			_go_fuzz_dep_.CoverTab[130410]++
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:34
			// _ = "end of CoverTab[130410]"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:34
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:34
		// _ = "end of CoverTab[130408]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:35
	// _ = "end of CoverTab[130407]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:39
func (f *Float64) Sub(s float64) float64 {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:39
	_go_fuzz_dep_.CoverTab[130411]++
										return f.Add(-s)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:40
	// _ = "end of CoverTab[130411]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:44
func (f *Float64) String() string {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:44
	_go_fuzz_dep_.CoverTab[130412]++

										return strconv.FormatFloat(f.Load(), 'g', -1, 64)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:46
	// _ = "end of CoverTab[130412]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:47
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/float64_ext.go:47
var _ = _go_fuzz_dep_.CoverTab
