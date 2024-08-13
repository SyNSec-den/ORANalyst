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

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:21
package atomic

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:21
)

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:26
//go:generate bin/gen-atomicwrapper -name=Error -type=error -wrapped=Value -pack=packError -unpack=unpackError -file=error.go

type packedError struct{ Value error }

func packError(v error) interface{} {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:30
	_go_fuzz_dep_.CoverTab[130390]++
										return packedError{v}
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:31
	// _ = "end of CoverTab[130390]"
}

func unpackError(v interface{}) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:34
	_go_fuzz_dep_.CoverTab[130391]++
										if err, ok := v.(packedError); ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:35
		_go_fuzz_dep_.CoverTab[130393]++
											return err.Value
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:36
		// _ = "end of CoverTab[130393]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:37
		_go_fuzz_dep_.CoverTab[130394]++
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:37
		// _ = "end of CoverTab[130394]"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:37
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:37
	// _ = "end of CoverTab[130391]"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:37
	_go_fuzz_dep_.CoverTab[130392]++
										return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:38
	// _ = "end of CoverTab[130392]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:39
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/error_ext.go:39
var _ = _go_fuzz_dep_.CoverTab
