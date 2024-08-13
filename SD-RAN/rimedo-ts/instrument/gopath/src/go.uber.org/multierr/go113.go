// Copyright (c) 2019 Uber Technologies, Inc.
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

// +build go1.13

//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:23
package multierr

//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:23
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:23
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:23
)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:23
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:23
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:23
)

import "errors"

// As attempts to find the first error in the error list that matches the type
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:27
// of the value that target points to.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:27
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:27
// This function allows errors.As to traverse the values stored on the
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:27
// multierr error.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:32
func (merr *multiError) As(target interface{}) bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:32
	_go_fuzz_dep_.CoverTab[130567]++
										for _, err := range merr.Errors() {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:33
		_go_fuzz_dep_.CoverTab[130569]++
											if errors.As(err, target) {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:34
			_go_fuzz_dep_.CoverTab[130570]++
												return true
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:35
			// _ = "end of CoverTab[130570]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:36
			_go_fuzz_dep_.CoverTab[130571]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:36
			// _ = "end of CoverTab[130571]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:36
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:36
		// _ = "end of CoverTab[130569]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:37
	// _ = "end of CoverTab[130567]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:37
	_go_fuzz_dep_.CoverTab[130568]++
										return false
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:38
	// _ = "end of CoverTab[130568]"
}

// Is attempts to match the provided error against errors in the error list.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:41
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:41
// This function allows errors.Is to traverse the values stored on the
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:41
// multierr error.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:45
func (merr *multiError) Is(target error) bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:45
	_go_fuzz_dep_.CoverTab[130572]++
										for _, err := range merr.Errors() {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:46
		_go_fuzz_dep_.CoverTab[130574]++
											if errors.Is(err, target) {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:47
			_go_fuzz_dep_.CoverTab[130575]++
												return true
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:48
			// _ = "end of CoverTab[130575]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:49
			_go_fuzz_dep_.CoverTab[130576]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:49
			// _ = "end of CoverTab[130576]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:49
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:49
		// _ = "end of CoverTab[130574]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:50
	// _ = "end of CoverTab[130572]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:50
	_go_fuzz_dep_.CoverTab[130573]++
										return false
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:51
	// _ = "end of CoverTab[130573]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:52
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/go113.go:52
var _ = _go_fuzz_dep_.CoverTab
