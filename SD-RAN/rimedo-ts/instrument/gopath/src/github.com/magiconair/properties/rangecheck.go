// Copyright 2018 Frank Schroeder. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:5
package properties

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:5
)

import (
	"fmt"
	"math"
)

// make this a var to overwrite it in a test
var is32Bit = ^uint(0) == math.MaxUint32

// intRangeCheck checks if the value fits into the int type and
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:15
// panics if it does not.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:17
func intRangeCheck(key string, v int64) int {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:17
	_go_fuzz_dep_.CoverTab[116163]++
												if is32Bit && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:18
		_go_fuzz_dep_.CoverTab[116165]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:18
		return (v < math.MinInt32 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:18
			_go_fuzz_dep_.CoverTab[116166]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:18
			return v > math.MaxInt32
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:18
			// _ = "end of CoverTab[116166]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:18
		}())
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:18
		// _ = "end of CoverTab[116165]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:18
	}() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:18
		_go_fuzz_dep_.CoverTab[116167]++
													panic(fmt.Sprintf("Value %d for key %s out of range", v, key))
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:19
		// _ = "end of CoverTab[116167]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:20
		_go_fuzz_dep_.CoverTab[116168]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:20
		// _ = "end of CoverTab[116168]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:20
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:20
	// _ = "end of CoverTab[116163]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:20
	_go_fuzz_dep_.CoverTab[116164]++
												return int(v)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:21
	// _ = "end of CoverTab[116164]"
}

// uintRangeCheck checks if the value fits into the uint type and
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:24
// panics if it does not.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:26
func uintRangeCheck(key string, v uint64) uint {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:26
	_go_fuzz_dep_.CoverTab[116169]++
												if is32Bit && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:27
		_go_fuzz_dep_.CoverTab[116171]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:27
		return v > math.MaxUint32
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:27
		// _ = "end of CoverTab[116171]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:27
	}() {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:27
		_go_fuzz_dep_.CoverTab[116172]++
													panic(fmt.Sprintf("Value %d for key %s out of range", v, key))
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:28
		// _ = "end of CoverTab[116172]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:29
		_go_fuzz_dep_.CoverTab[116173]++
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:29
		// _ = "end of CoverTab[116173]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:29
	// _ = "end of CoverTab[116169]"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:29
	_go_fuzz_dep_.CoverTab[116170]++
												return uint(v)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:30
	// _ = "end of CoverTab[116170]"
}

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:31
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/rangecheck.go:31
var _ = _go_fuzz_dep_.CoverTab
