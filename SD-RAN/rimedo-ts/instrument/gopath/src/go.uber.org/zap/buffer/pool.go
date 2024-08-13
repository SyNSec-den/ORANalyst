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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:21
package buffer

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:21
)

import "sync"

// A Pool is a type-safe wrapper around a sync.Pool.
type Pool struct {
	p *sync.Pool
}

// NewPool constructs a new Pool.
func NewPool() Pool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:31
	_go_fuzz_dep_.CoverTab[130596]++
										return Pool{p: &sync.Pool{
		New: func() interface{} {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:33
			_go_fuzz_dep_.CoverTab[130597]++
												return &Buffer{bs: make([]byte, 0, _size)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:34
			// _ = "end of CoverTab[130597]"
		},
	}}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:36
	// _ = "end of CoverTab[130596]"
}

// Get retrieves a Buffer from the pool, creating one if necessary.
func (p Pool) Get() *Buffer {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:40
	_go_fuzz_dep_.CoverTab[130598]++
										buf := p.p.Get().(*Buffer)
										buf.Reset()
										buf.pool = p
										return buf
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:44
	// _ = "end of CoverTab[130598]"
}

func (p Pool) put(buf *Buffer) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:47
	_go_fuzz_dep_.CoverTab[130599]++
										p.p.Put(buf)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:48
	// _ = "end of CoverTab[130599]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:49
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/pool.go:49
var _ = _go_fuzz_dep_.CoverTab
