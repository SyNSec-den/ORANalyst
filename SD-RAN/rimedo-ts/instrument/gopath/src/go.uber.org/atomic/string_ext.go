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

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:21
package atomic

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:21
)

//go:generate bin/gen-atomicwrapper -name=String -type=string -wrapped=Value -file=string.go

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:26
func (s *String) String() string {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:26
	_go_fuzz_dep_.CoverTab[130452]++
										return s.Load()
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:27
	// _ = "end of CoverTab[130452]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:33
func (s *String) MarshalText() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:33
	_go_fuzz_dep_.CoverTab[130453]++
										return []byte(s.Load()), nil
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:34
	// _ = "end of CoverTab[130453]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:40
func (s *String) UnmarshalText(b []byte) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:40
	_go_fuzz_dep_.CoverTab[130454]++
										s.Store(string(b))
										return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:42
	// _ = "end of CoverTab[130454]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:43
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/string_ext.go:43
var _ = _go_fuzz_dep_.CoverTab
