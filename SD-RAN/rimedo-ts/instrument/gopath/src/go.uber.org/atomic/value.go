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

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/value.go:21
package atomic

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/value.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/value.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/value.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/value.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/value.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/value.go:21
)

import "sync/atomic"

// Value shadows the type of the same name from sync/atomic
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/value.go:25
// https://godoc.org/sync/atomic#Value
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/value.go:27
type Value struct {
	atomic.Value

	_	nocmp	// disallow non-atomic comparison
}

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/value.go:31
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/value.go:31
var _ = _go_fuzz_dep_.CoverTab
