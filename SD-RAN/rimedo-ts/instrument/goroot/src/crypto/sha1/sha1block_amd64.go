// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:5
package sha1

//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:5
import (
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:5
)
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:5
import (
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:5
)

import "internal/cpu"

//go:noescape
func blockAVX2(dig *digest, p []byte)

//go:noescape
func blockAMD64(dig *digest, p []byte)

var useAVX2 = cpu.X86.HasAVX2 && func() bool {
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:15
	_go_fuzz_dep_.CoverTab[10158]++
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:15
	return cpu.X86.HasBMI1
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:15
	// _ = "end of CoverTab[10158]"
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:15
}() && func() bool {
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:15
	_go_fuzz_dep_.CoverTab[10159]++
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:15
	return cpu.X86.HasBMI2
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:15
	// _ = "end of CoverTab[10159]"
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:15
}()

func block(dig *digest, p []byte) {
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:17
	_go_fuzz_dep_.CoverTab[10160]++
								if useAVX2 && func() bool {
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:18
		_go_fuzz_dep_.CoverTab[10161]++
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:18
		return len(p) >= 256
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:18
		// _ = "end of CoverTab[10161]"
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:18
	}() {
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:18
		_go_fuzz_dep_.CoverTab[10162]++

//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:25
		safeLen := len(p) - 128
		if safeLen%128 != 0 {
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:26
			_go_fuzz_dep_.CoverTab[10164]++
										safeLen -= 64
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:27
			// _ = "end of CoverTab[10164]"
		} else {
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:28
			_go_fuzz_dep_.CoverTab[10165]++
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:28
			// _ = "end of CoverTab[10165]"
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:28
		}
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:28
		// _ = "end of CoverTab[10162]"
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:28
		_go_fuzz_dep_.CoverTab[10163]++
									blockAVX2(dig, p[:safeLen])
									blockAMD64(dig, p[safeLen:])
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:30
		// _ = "end of CoverTab[10163]"
	} else {
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:31
		_go_fuzz_dep_.CoverTab[10166]++
									blockAMD64(dig, p)
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:32
		// _ = "end of CoverTab[10166]"
	}
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:33
	// _ = "end of CoverTab[10160]"
}

//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:34
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/sha1/sha1block_amd64.go:34
var _ = _go_fuzz_dep_.CoverTab
