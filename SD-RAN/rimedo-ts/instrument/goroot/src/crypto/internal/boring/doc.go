// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/internal/boring/doc.go:5
// Package boring provides access to BoringCrypto implementation functions.
//line /usr/local/go/src/crypto/internal/boring/doc.go:5
// Check the constant Enabled to find out whether BoringCrypto is available.
//line /usr/local/go/src/crypto/internal/boring/doc.go:5
// If BoringCrypto is not available, the functions in this package all panic.
//line /usr/local/go/src/crypto/internal/boring/doc.go:8
package boring

//line /usr/local/go/src/crypto/internal/boring/doc.go:8
import (
//line /usr/local/go/src/crypto/internal/boring/doc.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/boring/doc.go:8
)
//line /usr/local/go/src/crypto/internal/boring/doc.go:8
import (
//line /usr/local/go/src/crypto/internal/boring/doc.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/boring/doc.go:8
)

// Enabled reports whether BoringCrypto is available.
//line /usr/local/go/src/crypto/internal/boring/doc.go:10
// When enabled is false, all functions in this package panic.
//line /usr/local/go/src/crypto/internal/boring/doc.go:10
//
//line /usr/local/go/src/crypto/internal/boring/doc.go:10
// BoringCrypto is only available on linux/amd64 systems.
//line /usr/local/go/src/crypto/internal/boring/doc.go:14
const Enabled = available

// A BigInt is the raw words from a BigInt.
//line /usr/local/go/src/crypto/internal/boring/doc.go:16
// This definition allows us to avoid importing math/big.
//line /usr/local/go/src/crypto/internal/boring/doc.go:16
// Conversion between BigInt and *big.Int is in crypto/internal/boring/bbig.
//line /usr/local/go/src/crypto/internal/boring/doc.go:19
type BigInt []uint

//line /usr/local/go/src/crypto/internal/boring/doc.go:19
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/boring/doc.go:19
var _ = _go_fuzz_dep_.CoverTab
