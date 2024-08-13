// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/hash/hash.go:5
// Package hash provides interfaces for hash functions.
package hash

//line /usr/local/go/src/hash/hash.go:6
import (
//line /usr/local/go/src/hash/hash.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/hash/hash.go:6
)
//line /usr/local/go/src/hash/hash.go:6
import (
//line /usr/local/go/src/hash/hash.go:6
	_atomic_ "sync/atomic"
//line /usr/local/go/src/hash/hash.go:6
)

import "io"

// Hash is the common interface implemented by all hash functions.
//line /usr/local/go/src/hash/hash.go:10
//
//line /usr/local/go/src/hash/hash.go:10
// Hash implementations in the standard library (e.g. hash/crc32 and
//line /usr/local/go/src/hash/hash.go:10
// crypto/sha256) implement the encoding.BinaryMarshaler and
//line /usr/local/go/src/hash/hash.go:10
// encoding.BinaryUnmarshaler interfaces. Marshaling a hash implementation
//line /usr/local/go/src/hash/hash.go:10
// allows its internal state to be saved and used for additional processing
//line /usr/local/go/src/hash/hash.go:10
// later, without having to re-write the data previously written to the hash.
//line /usr/local/go/src/hash/hash.go:10
// The hash state may contain portions of the input in its original form,
//line /usr/local/go/src/hash/hash.go:10
// which users are expected to handle for any possible security implications.
//line /usr/local/go/src/hash/hash.go:10
//
//line /usr/local/go/src/hash/hash.go:10
// Compatibility: Any future changes to hash or crypto packages will endeavor
//line /usr/local/go/src/hash/hash.go:10
// to maintain compatibility with state encoded using previous versions.
//line /usr/local/go/src/hash/hash.go:10
// That is, any released versions of the packages should be able to
//line /usr/local/go/src/hash/hash.go:10
// decode data written with any previously released version,
//line /usr/local/go/src/hash/hash.go:10
// subject to issues such as security fixes.
//line /usr/local/go/src/hash/hash.go:10
// See the Go compatibility document for background: https://golang.org/doc/go1compat
//line /usr/local/go/src/hash/hash.go:26
type Hash interface {
	// Write (via the embedded io.Writer interface) adds more data to the running hash.
	// It never returns an error.
	io.Writer

	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	Sum(b []byte) []byte

	// Reset resets the Hash to its initial state.
	Reset()

	// Size returns the number of bytes Sum will return.
	Size() int

	// BlockSize returns the hash's underlying block size.
	// The Write method must be able to accept any amount
	// of data, but it may operate more efficiently if all writes
	// are a multiple of the block size.
	BlockSize() int
}

// Hash32 is the common interface implemented by all 32-bit hash functions.
type Hash32 interface {
	Hash
	Sum32() uint32
}

// Hash64 is the common interface implemented by all 64-bit hash functions.
type Hash64 interface {
	Hash
	Sum64() uint64
}

//line /usr/local/go/src/hash/hash.go:58
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/hash/hash.go:58
var _ = _go_fuzz_dep_.CoverTab
