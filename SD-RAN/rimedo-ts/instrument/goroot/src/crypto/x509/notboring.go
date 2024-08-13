// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !boringcrypto

//line /usr/local/go/src/crypto/x509/notboring.go:7
package x509

//line /usr/local/go/src/crypto/x509/notboring.go:7
import (
//line /usr/local/go/src/crypto/x509/notboring.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/x509/notboring.go:7
)
//line /usr/local/go/src/crypto/x509/notboring.go:7
import (
//line /usr/local/go/src/crypto/x509/notboring.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/x509/notboring.go:7
)

func boringAllowCert(c *Certificate) bool {
//line /usr/local/go/src/crypto/x509/notboring.go:9
	_go_fuzz_dep_.CoverTab[18396]++
//line /usr/local/go/src/crypto/x509/notboring.go:9
	return true
//line /usr/local/go/src/crypto/x509/notboring.go:9
	// _ = "end of CoverTab[18396]"
//line /usr/local/go/src/crypto/x509/notboring.go:9
}

//line /usr/local/go/src/crypto/x509/notboring.go:9
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/x509/notboring.go:9
var _ = _go_fuzz_dep_.CoverTab
