// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !s390x

//line /usr/local/go/src/crypto/ecdsa/ecdsa_noasm.go:7
package ecdsa

//line /usr/local/go/src/crypto/ecdsa/ecdsa_noasm.go:7
import (
//line /usr/local/go/src/crypto/ecdsa/ecdsa_noasm.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_noasm.go:7
)
//line /usr/local/go/src/crypto/ecdsa/ecdsa_noasm.go:7
import (
//line /usr/local/go/src/crypto/ecdsa/ecdsa_noasm.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/ecdsa/ecdsa_noasm.go:7
)

import "io"

func verifyAsm(pub *PublicKey, hash []byte, sig []byte) error {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_noasm.go:11
	_go_fuzz_dep_.CoverTab[9149]++
								return errNoAsm
//line /usr/local/go/src/crypto/ecdsa/ecdsa_noasm.go:12
	// _ = "end of CoverTab[9149]"
}

func signAsm(priv *PrivateKey, csprng io.Reader, hash []byte) (sig []byte, err error) {
//line /usr/local/go/src/crypto/ecdsa/ecdsa_noasm.go:15
	_go_fuzz_dep_.CoverTab[9150]++
								return nil, errNoAsm
//line /usr/local/go/src/crypto/ecdsa/ecdsa_noasm.go:16
	// _ = "end of CoverTab[9150]"
}

//line /usr/local/go/src/crypto/ecdsa/ecdsa_noasm.go:17
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/ecdsa/ecdsa_noasm.go:17
var _ = _go_fuzz_dep_.CoverTab
