// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (!arm64 && !s390x && !ppc64le) || (arm64 && !go1.11) || !gc || purego
// +build !arm64,!s390x,!ppc64le arm64,!go1.11 !gc purego

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_noasm.go:8
package chacha20

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_noasm.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_noasm.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_noasm.go:8
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_noasm.go:8
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_noasm.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_noasm.go:8
)

const bufSize = blockSize

func (s *Cipher) xorKeyStreamBlocks(dst, src []byte) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_noasm.go:12
	_go_fuzz_dep_.CoverTab[20732]++
											s.xorKeyStreamBlocksGeneric(dst, src)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_noasm.go:13
	// _ = "end of CoverTab[20732]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_noasm.go:14
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/chacha20/chacha_noasm.go:14
var _ = _go_fuzz_dep_.CoverTab
