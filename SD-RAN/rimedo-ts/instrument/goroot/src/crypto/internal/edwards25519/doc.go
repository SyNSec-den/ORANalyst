// Copyright (c) 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
// Package edwards25519 implements group logic for the twisted Edwards curve
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
//
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
//	-x^2 + y^2 = 1 + -(121665/121666)*x^2*y^2
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
//
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
// This is better known as the Edwards curve equivalent to Curve25519, and is
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
// the curve used by the Ed25519 signature scheme.
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
//
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
// Most users don't need this package, and should instead use crypto/ed25519 for
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
// signatures, golang.org/x/crypto/curve25519 for Diffie-Hellman, or
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
// github.com/gtank/ristretto255 for prime order group logic.
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
//
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
// However, developers who do need to interact with low-level edwards25519
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
// operations can use filippo.io/edwards25519, an extended version of this
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
// package repackaged as an importable module.
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
//
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
// (Note that filippo.io/edwards25519 and github.com/gtank/ristretto255 are not
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:5
// maintained by the Go team and are not covered by the Go 1 Compatibility Promise.)
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:22
package edwards25519

//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:22
import (
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:22
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:22
)
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:22
import (
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:22
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:22
)

//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:22
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/edwards25519/doc.go:22
var _ = _go_fuzz_dep_.CoverTab
