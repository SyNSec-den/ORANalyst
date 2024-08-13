// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:5
package uuid

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:5
)

import (
	"io"
)

// randomBits completely fills slice b with random data.
func randomBits(b []byte) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:12
	_go_fuzz_dep_.CoverTab[179436]++
										if _, err := io.ReadFull(rander, b); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:13
		_go_fuzz_dep_.CoverTab[179437]++
											panic(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:14
		// _ = "end of CoverTab[179437]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:15
		_go_fuzz_dep_.CoverTab[179438]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:15
		// _ = "end of CoverTab[179438]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:15
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:15
	// _ = "end of CoverTab[179436]"
}

// xvalues returns the value of a byte as a hexadecimal digit or 255.
var xvalues = [256]byte{
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
}

// xtob converts hex characters x1 and x2 into a byte.
func xtob(x1, x2 byte) (byte, bool) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:39
	_go_fuzz_dep_.CoverTab[179439]++
										b1 := xvalues[x1]
										b2 := xvalues[x2]
										return (b1 << 4) | b2, b1 != 255 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:42
		_go_fuzz_dep_.CoverTab[179440]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:42
		return b2 != 255
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:42
		// _ = "end of CoverTab[179440]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:42
	}()
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:42
	// _ = "end of CoverTab[179439]"
}

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:43
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/util.go:43
var _ = _go_fuzz_dep_.CoverTab
