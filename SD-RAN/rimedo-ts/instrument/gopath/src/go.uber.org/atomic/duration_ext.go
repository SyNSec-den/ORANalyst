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

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:21
package atomic

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:21
)

import "time"

//go:generate bin/gen-atomicwrapper -name=Duration -type=time.Duration -wrapped=Int64 -pack=int64 -unpack=time.Duration -cas -swap -json -imports time -file=duration.go

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:28
func (d *Duration) Add(n time.Duration) time.Duration {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:28
	_go_fuzz_dep_.CoverTab[130381]++
										return time.Duration(d.v.Add(int64(n)))
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:29
	// _ = "end of CoverTab[130381]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:33
func (d *Duration) Sub(n time.Duration) time.Duration {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:33
	_go_fuzz_dep_.CoverTab[130382]++
										return time.Duration(d.v.Sub(int64(n)))
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:34
	// _ = "end of CoverTab[130382]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:38
func (d *Duration) String() string {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:38
	_go_fuzz_dep_.CoverTab[130383]++
										return d.Load().String()
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:39
	// _ = "end of CoverTab[130383]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:40
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/duration_ext.go:40
var _ = _go_fuzz_dep_.CoverTab
