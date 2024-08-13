// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64

//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:7
package sha512

//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:7
import (
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:7
)
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:7
import (
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:7
)

import "internal/cpu"

//go:noescape
func blockAVX2(dig *digest, p []byte)

//go:noescape
func blockAMD64(dig *digest, p []byte)

var useAVX2 = cpu.X86.HasAVX2 && func() bool {
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:17
	_go_fuzz_dep_.CoverTab[7435]++
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:17
	return cpu.X86.HasBMI1
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:17
	// _ = "end of CoverTab[7435]"
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:17
}() && func() bool {
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:17
	_go_fuzz_dep_.CoverTab[7436]++
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:17
	return cpu.X86.HasBMI2
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:17
	// _ = "end of CoverTab[7436]"
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:17
}()

func block(dig *digest, p []byte) {
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:19
	_go_fuzz_dep_.CoverTab[7437]++
								if useAVX2 {
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:20
		_go_fuzz_dep_.CoverTab[7438]++
									blockAVX2(dig, p)
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:21
		// _ = "end of CoverTab[7438]"
	} else {
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:22
		_go_fuzz_dep_.CoverTab[7439]++
									blockAMD64(dig, p)
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:23
		// _ = "end of CoverTab[7439]"
	}
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:24
	// _ = "end of CoverTab[7437]"
}

//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:25
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/sha512/sha512block_amd64.go:25
var _ = _go_fuzz_dep_.CoverTab
