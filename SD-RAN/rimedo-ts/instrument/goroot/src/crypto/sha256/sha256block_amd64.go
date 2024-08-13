// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/sha256/sha256block_amd64.go:5
package sha256

//line /usr/local/go/src/crypto/sha256/sha256block_amd64.go:5
import (
//line /usr/local/go/src/crypto/sha256/sha256block_amd64.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/sha256/sha256block_amd64.go:5
)
//line /usr/local/go/src/crypto/sha256/sha256block_amd64.go:5
import (
//line /usr/local/go/src/crypto/sha256/sha256block_amd64.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/sha256/sha256block_amd64.go:5
)

import "internal/cpu"

var useAVX2 = cpu.X86.HasAVX2 && func() bool {
//line /usr/local/go/src/crypto/sha256/sha256block_amd64.go:9
	_go_fuzz_dep_.CoverTab[10245]++
//line /usr/local/go/src/crypto/sha256/sha256block_amd64.go:9
	return cpu.X86.HasBMI2
//line /usr/local/go/src/crypto/sha256/sha256block_amd64.go:9
	// _ = "end of CoverTab[10245]"
//line /usr/local/go/src/crypto/sha256/sha256block_amd64.go:9
}()
//line /usr/local/go/src/crypto/sha256/sha256block_amd64.go:9
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/sha256/sha256block_amd64.go:9
var _ = _go_fuzz_dep_.CoverTab
