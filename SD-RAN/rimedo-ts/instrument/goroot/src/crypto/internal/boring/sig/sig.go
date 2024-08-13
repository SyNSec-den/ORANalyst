// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/internal/boring/sig/sig.go:5
// Package sig holds “code signatures” that can be called
//line /usr/local/go/src/crypto/internal/boring/sig/sig.go:5
// and will result in certain code sequences being linked into
//line /usr/local/go/src/crypto/internal/boring/sig/sig.go:5
// the final binary. The functions themselves are no-ops.
//line /usr/local/go/src/crypto/internal/boring/sig/sig.go:8
package sig

//line /usr/local/go/src/crypto/internal/boring/sig/sig.go:8
import (
//line /usr/local/go/src/crypto/internal/boring/sig/sig.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/boring/sig/sig.go:8
)
//line /usr/local/go/src/crypto/internal/boring/sig/sig.go:8
import (
//line /usr/local/go/src/crypto/internal/boring/sig/sig.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/boring/sig/sig.go:8
)

// BoringCrypto indicates that the BoringCrypto module is present.
func BoringCrypto()

// FIPSOnly indicates that package crypto/tls/fipsonly is present.
func FIPSOnly()

// StandardCrypto indicates that standard Go crypto is present.
func StandardCrypto()

//line /usr/local/go/src/crypto/internal/boring/sig/sig.go:17
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/boring/sig/sig.go:17
var _ = _go_fuzz_dep_.CoverTab
