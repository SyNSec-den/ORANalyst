// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/cipher/cipher.go:5
// Package cipher implements standard block cipher modes that can be wrapped
//line /usr/local/go/src/crypto/cipher/cipher.go:5
// around low-level block cipher implementations.
//line /usr/local/go/src/crypto/cipher/cipher.go:5
// See https://csrc.nist.gov/groups/ST/toolkit/BCM/current_modes.html
//line /usr/local/go/src/crypto/cipher/cipher.go:5
// and NIST Special Publication 800-38A.
//line /usr/local/go/src/crypto/cipher/cipher.go:9
package cipher

//line /usr/local/go/src/crypto/cipher/cipher.go:9
import (
//line /usr/local/go/src/crypto/cipher/cipher.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/cipher/cipher.go:9
)
//line /usr/local/go/src/crypto/cipher/cipher.go:9
import (
//line /usr/local/go/src/crypto/cipher/cipher.go:9
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/cipher/cipher.go:9
)

// A Block represents an implementation of block cipher
//line /usr/local/go/src/crypto/cipher/cipher.go:11
// using a given key. It provides the capability to encrypt
//line /usr/local/go/src/crypto/cipher/cipher.go:11
// or decrypt individual blocks. The mode implementations
//line /usr/local/go/src/crypto/cipher/cipher.go:11
// extend that capability to streams of blocks.
//line /usr/local/go/src/crypto/cipher/cipher.go:15
type Block interface {
	// BlockSize returns the cipher's block size.
	BlockSize() int

	// Encrypt encrypts the first block in src into dst.
	// Dst and src must overlap entirely or not at all.
	Encrypt(dst, src []byte)

	// Decrypt decrypts the first block in src into dst.
	// Dst and src must overlap entirely or not at all.
	Decrypt(dst, src []byte)
}

// A Stream represents a stream cipher.
type Stream interface {
	// XORKeyStream XORs each byte in the given slice with a byte from the
	// cipher's key stream. Dst and src must overlap entirely or not at all.
	//
	// If len(dst) < len(src), XORKeyStream should panic. It is acceptable
	// to pass a dst bigger than src, and in that case, XORKeyStream will
	// only update dst[:len(src)] and will not touch the rest of dst.
	//
	// Multiple calls to XORKeyStream behave as if the concatenation of
	// the src buffers was passed in a single run. That is, Stream
	// maintains state and does not reset at each XORKeyStream call.
	XORKeyStream(dst, src []byte)
}

// A BlockMode represents a block cipher running in a block-based mode (CBC,
//line /usr/local/go/src/crypto/cipher/cipher.go:43
// ECB etc).
//line /usr/local/go/src/crypto/cipher/cipher.go:45
type BlockMode interface {
	// BlockSize returns the mode's block size.
	BlockSize() int

	// CryptBlocks encrypts or decrypts a number of blocks. The length of
	// src must be a multiple of the block size. Dst and src must overlap
	// entirely or not at all.
	//
	// If len(dst) < len(src), CryptBlocks should panic. It is acceptable
	// to pass a dst bigger than src, and in that case, CryptBlocks will
	// only update dst[:len(src)] and will not touch the rest of dst.
	//
	// Multiple calls to CryptBlocks behave as if the concatenation of
	// the src buffers was passed in a single run. That is, BlockMode
	// maintains state and does not reset at each CryptBlocks call.
	CryptBlocks(dst, src []byte)
}

//line /usr/local/go/src/crypto/cipher/cipher.go:61
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/cipher/cipher.go:61
var _ = _go_fuzz_dep_.CoverTab
