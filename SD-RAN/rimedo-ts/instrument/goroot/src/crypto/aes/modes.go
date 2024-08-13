// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/aes/modes.go:5
package aes

//line /usr/local/go/src/crypto/aes/modes.go:5
import (
//line /usr/local/go/src/crypto/aes/modes.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/aes/modes.go:5
)
//line /usr/local/go/src/crypto/aes/modes.go:5
import (
//line /usr/local/go/src/crypto/aes/modes.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/aes/modes.go:5
)

import (
	"crypto/cipher"
)

// gcmAble is implemented by cipher.Blocks that can provide an optimized
//line /usr/local/go/src/crypto/aes/modes.go:11
// implementation of GCM through the AEAD interface.
//line /usr/local/go/src/crypto/aes/modes.go:11
// See crypto/cipher/gcm.go.
//line /usr/local/go/src/crypto/aes/modes.go:14
type gcmAble interface {
	NewGCM(nonceSize, tagSize int) (cipher.AEAD, error)
}

// cbcEncAble is implemented by cipher.Blocks that can provide an optimized
//line /usr/local/go/src/crypto/aes/modes.go:18
// implementation of CBC encryption through the cipher.BlockMode interface.
//line /usr/local/go/src/crypto/aes/modes.go:18
// See crypto/cipher/cbc.go.
//line /usr/local/go/src/crypto/aes/modes.go:21
type cbcEncAble interface {
	NewCBCEncrypter(iv []byte) cipher.BlockMode
}

// cbcDecAble is implemented by cipher.Blocks that can provide an optimized
//line /usr/local/go/src/crypto/aes/modes.go:25
// implementation of CBC decryption through the cipher.BlockMode interface.
//line /usr/local/go/src/crypto/aes/modes.go:25
// See crypto/cipher/cbc.go.
//line /usr/local/go/src/crypto/aes/modes.go:28
type cbcDecAble interface {
	NewCBCDecrypter(iv []byte) cipher.BlockMode
}

// ctrAble is implemented by cipher.Blocks that can provide an optimized
//line /usr/local/go/src/crypto/aes/modes.go:32
// implementation of CTR through the cipher.Stream interface.
//line /usr/local/go/src/crypto/aes/modes.go:32
// See crypto/cipher/ctr.go.
//line /usr/local/go/src/crypto/aes/modes.go:35
type ctrAble interface {
	NewCTR(iv []byte) cipher.Stream
}

//line /usr/local/go/src/crypto/aes/modes.go:37
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/aes/modes.go:37
var _ = _go_fuzz_dep_.CoverTab
