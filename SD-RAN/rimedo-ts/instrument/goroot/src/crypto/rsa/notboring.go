// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !boringcrypto

//line /usr/local/go/src/crypto/rsa/notboring.go:7
package rsa

//line /usr/local/go/src/crypto/rsa/notboring.go:7
import (
//line /usr/local/go/src/crypto/rsa/notboring.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/rsa/notboring.go:7
)
//line /usr/local/go/src/crypto/rsa/notboring.go:7
import (
//line /usr/local/go/src/crypto/rsa/notboring.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/rsa/notboring.go:7
)

import "crypto/internal/boring"

func boringPublicKey(*PublicKey) (*boring.PublicKeyRSA, error) {
//line /usr/local/go/src/crypto/rsa/notboring.go:11
	_go_fuzz_dep_.CoverTab[9617]++
							panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/rsa/notboring.go:12
	// _ = "end of CoverTab[9617]"
}
func boringPrivateKey(*PrivateKey) (*boring.PrivateKeyRSA, error) {
//line /usr/local/go/src/crypto/rsa/notboring.go:14
	_go_fuzz_dep_.CoverTab[9618]++
							panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/rsa/notboring.go:15
	// _ = "end of CoverTab[9618]"
}

//line /usr/local/go/src/crypto/rsa/notboring.go:16
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/rsa/notboring.go:16
var _ = _go_fuzz_dep_.CoverTab
