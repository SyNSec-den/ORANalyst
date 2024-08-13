// Copyright (c) 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !arm64 || !gc || purego

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_arm64_noasm.go:7
package field

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_arm64_noasm.go:7
import (
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_arm64_noasm.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_arm64_noasm.go:7
)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_arm64_noasm.go:7
import (
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_arm64_noasm.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_arm64_noasm.go:7
)

func (v *Element) carryPropagate() *Element {
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_arm64_noasm.go:9
	_go_fuzz_dep_.CoverTab[2046]++
											return v.carryPropagateGeneric()
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_arm64_noasm.go:10
	// _ = "end of CoverTab[2046]"
}

//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_arm64_noasm.go:11
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/edwards25519/field/fe_arm64_noasm.go:11
var _ = _go_fuzz_dep_.CoverTab
