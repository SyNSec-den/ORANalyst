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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:21
// Package buffer provides a thin wrapper around a byte slice. Unlike the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:21
// standard library's bytes.Buffer, it supports a portion of the strconv
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:21
// package's zero-allocation formatters.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:24
package buffer

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:24
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:24
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:24
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:24
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:24
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:24
)

import (
	"strconv"
	"time"
)

const _size = 1024	// by default, create 1 KiB buffers

// Buffer is a thin wrapper around a byte slice. It's intended to be pooled, so
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:33
// the only way to construct one is via a Pool.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:35
type Buffer struct {
	bs	[]byte
	pool	Pool
}

// AppendByte writes a single byte to the Buffer.
func (b *Buffer) AppendByte(v byte) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:41
	_go_fuzz_dep_.CoverTab[130577]++
										b.bs = append(b.bs, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:42
	// _ = "end of CoverTab[130577]"
}

// AppendString writes a string to the Buffer.
func (b *Buffer) AppendString(s string) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:46
	_go_fuzz_dep_.CoverTab[130578]++
										b.bs = append(b.bs, s...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:47
	// _ = "end of CoverTab[130578]"
}

// AppendInt appends an integer to the underlying buffer (assuming base 10).
func (b *Buffer) AppendInt(i int64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:51
	_go_fuzz_dep_.CoverTab[130579]++
										b.bs = strconv.AppendInt(b.bs, i, 10)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:52
	// _ = "end of CoverTab[130579]"
}

// AppendTime appends the time formatted using the specified layout.
func (b *Buffer) AppendTime(t time.Time, layout string) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:56
	_go_fuzz_dep_.CoverTab[130580]++
										b.bs = t.AppendFormat(b.bs, layout)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:57
	// _ = "end of CoverTab[130580]"
}

// AppendUint appends an unsigned integer to the underlying buffer (assuming
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:60
// base 10).
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:62
func (b *Buffer) AppendUint(i uint64) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:62
	_go_fuzz_dep_.CoverTab[130581]++
										b.bs = strconv.AppendUint(b.bs, i, 10)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:63
	// _ = "end of CoverTab[130581]"
}

// AppendBool appends a bool to the underlying buffer.
func (b *Buffer) AppendBool(v bool) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:67
	_go_fuzz_dep_.CoverTab[130582]++
										b.bs = strconv.AppendBool(b.bs, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:68
	// _ = "end of CoverTab[130582]"
}

// AppendFloat appends a float to the underlying buffer. It doesn't quote NaN
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:71
// or +/- Inf.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:73
func (b *Buffer) AppendFloat(f float64, bitSize int) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:73
	_go_fuzz_dep_.CoverTab[130583]++
										b.bs = strconv.AppendFloat(b.bs, f, 'f', -1, bitSize)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:74
	// _ = "end of CoverTab[130583]"
}

// Len returns the length of the underlying byte slice.
func (b *Buffer) Len() int {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:78
	_go_fuzz_dep_.CoverTab[130584]++
										return len(b.bs)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:79
	// _ = "end of CoverTab[130584]"
}

// Cap returns the capacity of the underlying byte slice.
func (b *Buffer) Cap() int {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:83
	_go_fuzz_dep_.CoverTab[130585]++
										return cap(b.bs)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:84
	// _ = "end of CoverTab[130585]"
}

// Bytes returns a mutable reference to the underlying byte slice.
func (b *Buffer) Bytes() []byte {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:88
	_go_fuzz_dep_.CoverTab[130586]++
										return b.bs
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:89
	// _ = "end of CoverTab[130586]"
}

// String returns a string copy of the underlying byte slice.
func (b *Buffer) String() string {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:93
	_go_fuzz_dep_.CoverTab[130587]++
										return string(b.bs)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:94
	// _ = "end of CoverTab[130587]"
}

// Reset resets the underlying byte slice. Subsequent writes re-use the slice's
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:97
// backing array.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:99
func (b *Buffer) Reset() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:99
	_go_fuzz_dep_.CoverTab[130588]++
										b.bs = b.bs[:0]
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:100
	// _ = "end of CoverTab[130588]"
}

// Write implements io.Writer.
func (b *Buffer) Write(bs []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:104
	_go_fuzz_dep_.CoverTab[130589]++
										b.bs = append(b.bs, bs...)
										return len(bs), nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:106
	// _ = "end of CoverTab[130589]"
}

// TrimNewline trims any final "\n" byte from the end of the buffer.
func (b *Buffer) TrimNewline() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:110
	_go_fuzz_dep_.CoverTab[130590]++
										if i := len(b.bs) - 1; i >= 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:111
		_go_fuzz_dep_.CoverTab[130591]++
											if b.bs[i] == '\n' {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:112
			_go_fuzz_dep_.CoverTab[130592]++
												b.bs = b.bs[:i]
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:113
			// _ = "end of CoverTab[130592]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:114
			_go_fuzz_dep_.CoverTab[130593]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:114
			// _ = "end of CoverTab[130593]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:114
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:114
		// _ = "end of CoverTab[130591]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:115
		_go_fuzz_dep_.CoverTab[130594]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:115
		// _ = "end of CoverTab[130594]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:115
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:115
	// _ = "end of CoverTab[130590]"
}

// Free returns the Buffer to its Pool.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:118
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:118
// Callers must not retain references to the Buffer after calling Free.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:121
func (b *Buffer) Free() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:121
	_go_fuzz_dep_.CoverTab[130595]++
										b.pool.put(b)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:122
	// _ = "end of CoverTab[130595]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:123
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/buffer/buffer.go:123
var _ = _go_fuzz_dep_.CoverTab
