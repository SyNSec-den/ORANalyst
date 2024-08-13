// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !math_big_pure_go
// +build !math_big_pure_go

//line /usr/local/go/src/math/big/arith_amd64.go:8
package big

//line /usr/local/go/src/math/big/arith_amd64.go:8
import (
//line /usr/local/go/src/math/big/arith_amd64.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/arith_amd64.go:8
)
//line /usr/local/go/src/math/big/arith_amd64.go:8
import (
//line /usr/local/go/src/math/big/arith_amd64.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/arith_amd64.go:8
)

import "internal/cpu"

var support_adx = cpu.X86.HasADX && func() bool {
//line /usr/local/go/src/math/big/arith_amd64.go:12
	_go_fuzz_dep_.CoverTab[4120]++
//line /usr/local/go/src/math/big/arith_amd64.go:12
	return cpu.X86.HasBMI2
//line /usr/local/go/src/math/big/arith_amd64.go:12
	// _ = "end of CoverTab[4120]"
//line /usr/local/go/src/math/big/arith_amd64.go:12
}()
//line /usr/local/go/src/math/big/arith_amd64.go:12
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/arith_amd64.go:12
var _ = _go_fuzz_dep_.CoverTab
