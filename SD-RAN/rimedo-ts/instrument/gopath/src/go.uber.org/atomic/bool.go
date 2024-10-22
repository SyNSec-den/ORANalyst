// @generated Code generated by gen-atomicwrapper.

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

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:23
package atomic

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:23
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:23
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:23
)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:23
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:23
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:23
)

import (
	"encoding/json"
)

// Bool is an atomic type-safe wrapper for bool values.
type Bool struct {
	_	nocmp	// disallow non-atomic comparison

	v	Uint32
}

var _zeroBool bool

// NewBool creates a new Bool.
func NewBool(v bool) *Bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:39
	_go_fuzz_dep_.CoverTab[130345]++
									x := &Bool{}
									if v != _zeroBool {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:41
		_go_fuzz_dep_.CoverTab[130347]++
										x.Store(v)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:42
		// _ = "end of CoverTab[130347]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:43
		_go_fuzz_dep_.CoverTab[130348]++
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:43
		// _ = "end of CoverTab[130348]"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:43
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:43
	// _ = "end of CoverTab[130345]"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:43
	_go_fuzz_dep_.CoverTab[130346]++
									return x
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:44
	// _ = "end of CoverTab[130346]"
}

// Load atomically loads the wrapped bool.
func (x *Bool) Load() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:48
	_go_fuzz_dep_.CoverTab[130349]++
									return truthy(x.v.Load())
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:49
	// _ = "end of CoverTab[130349]"
}

// Store atomically stores the passed bool.
func (x *Bool) Store(v bool) {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:53
	_go_fuzz_dep_.CoverTab[130350]++
									x.v.Store(boolToInt(v))
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:54
	// _ = "end of CoverTab[130350]"
}

// CAS is an atomic compare-and-swap for bool values.
func (x *Bool) CAS(o, n bool) bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:58
	_go_fuzz_dep_.CoverTab[130351]++
									return x.v.CAS(boolToInt(o), boolToInt(n))
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:59
	// _ = "end of CoverTab[130351]"
}

// Swap atomically stores the given bool and returns the old
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:62
// value.
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:64
func (x *Bool) Swap(o bool) bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:64
	_go_fuzz_dep_.CoverTab[130352]++
									return truthy(x.v.Swap(boolToInt(o)))
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:65
	// _ = "end of CoverTab[130352]"
}

// MarshalJSON encodes the wrapped bool into JSON.
func (x *Bool) MarshalJSON() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:69
	_go_fuzz_dep_.CoverTab[130353]++
									return json.Marshal(x.Load())
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:70
	// _ = "end of CoverTab[130353]"
}

// UnmarshalJSON decodes a bool from JSON.
func (x *Bool) UnmarshalJSON(b []byte) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:74
	_go_fuzz_dep_.CoverTab[130354]++
									var v bool
									if err := json.Unmarshal(b, &v); err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:76
		_go_fuzz_dep_.CoverTab[130356]++
										return err
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:77
		// _ = "end of CoverTab[130356]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:78
		_go_fuzz_dep_.CoverTab[130357]++
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:78
		// _ = "end of CoverTab[130357]"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:78
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:78
	// _ = "end of CoverTab[130354]"
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:78
	_go_fuzz_dep_.CoverTab[130355]++
									x.Store(v)
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:80
	// _ = "end of CoverTab[130355]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:81
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/atomic@v1.7.0/bool.go:81
var _ = _go_fuzz_dep_.CoverTab
