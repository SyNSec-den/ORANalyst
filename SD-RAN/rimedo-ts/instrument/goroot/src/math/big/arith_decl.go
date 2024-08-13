// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !math_big_pure_go
// +build !math_big_pure_go

//line /usr/local/go/src/math/big/arith_decl.go:8
package big

//line /usr/local/go/src/math/big/arith_decl.go:8
import (
//line /usr/local/go/src/math/big/arith_decl.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/arith_decl.go:8
)
//line /usr/local/go/src/math/big/arith_decl.go:8
import (
//line /usr/local/go/src/math/big/arith_decl.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/arith_decl.go:8
)

//line /usr/local/go/src/math/big/arith_decl.go:12
//go:noescape
func addVV(z, x, y []Word) (c Word)

//go:noescape
func subVV(z, x, y []Word) (c Word)

//go:noescape
func addVW(z, x []Word, y Word) (c Word)

//go:noescape
func subVW(z, x []Word, y Word) (c Word)

//go:noescape
func shlVU(z, x []Word, s uint) (c Word)

//go:noescape
func shrVU(z, x []Word, s uint) (c Word)

//go:noescape
func mulAddVWW(z, x []Word, y, r Word) (c Word)

//go:noescape
func addMulVVW(z, x []Word, y Word) (c Word)

//line /usr/local/go/src/math/big/arith_decl.go:34
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/arith_decl.go:34
var _ = _go_fuzz_dep_.CoverTab
