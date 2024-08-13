// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !boringcrypto

//line /usr/local/go/src/crypto/tls/notboring.go:7
package tls

//line /usr/local/go/src/crypto/tls/notboring.go:7
import (
//line /usr/local/go/src/crypto/tls/notboring.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/notboring.go:7
)
//line /usr/local/go/src/crypto/tls/notboring.go:7
import (
//line /usr/local/go/src/crypto/tls/notboring.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/notboring.go:7
)

func needFIPS() bool	{ _go_fuzz_dep_.CoverTab[24922]++; return false; // _ = "end of CoverTab[24922]" }

func supportedSignatureAlgorithms() []SignatureScheme {
//line /usr/local/go/src/crypto/tls/notboring.go:11
	_go_fuzz_dep_.CoverTab[24923]++
							return defaultSupportedSignatureAlgorithms
//line /usr/local/go/src/crypto/tls/notboring.go:12
	// _ = "end of CoverTab[24923]"
}

func fipsMinVersion(c *Config) uint16 {
//line /usr/local/go/src/crypto/tls/notboring.go:15
	_go_fuzz_dep_.CoverTab[24924]++
//line /usr/local/go/src/crypto/tls/notboring.go:15
	panic("fipsMinVersion")
//line /usr/local/go/src/crypto/tls/notboring.go:15
	// _ = "end of CoverTab[24924]"
//line /usr/local/go/src/crypto/tls/notboring.go:15
}
func fipsMaxVersion(c *Config) uint16 {
//line /usr/local/go/src/crypto/tls/notboring.go:16
	_go_fuzz_dep_.CoverTab[24925]++
//line /usr/local/go/src/crypto/tls/notboring.go:16
	panic("fipsMaxVersion")
//line /usr/local/go/src/crypto/tls/notboring.go:16
	// _ = "end of CoverTab[24925]"
//line /usr/local/go/src/crypto/tls/notboring.go:16
}
func fipsCurvePreferences(c *Config) []CurveID {
//line /usr/local/go/src/crypto/tls/notboring.go:17
	_go_fuzz_dep_.CoverTab[24926]++
//line /usr/local/go/src/crypto/tls/notboring.go:17
	panic("fipsCurvePreferences")
//line /usr/local/go/src/crypto/tls/notboring.go:17
	// _ = "end of CoverTab[24926]"
//line /usr/local/go/src/crypto/tls/notboring.go:17
}
func fipsCipherSuites(c *Config) []uint16 {
//line /usr/local/go/src/crypto/tls/notboring.go:18
	_go_fuzz_dep_.CoverTab[24927]++
//line /usr/local/go/src/crypto/tls/notboring.go:18
	panic("fipsCipherSuites")
//line /usr/local/go/src/crypto/tls/notboring.go:18
	// _ = "end of CoverTab[24927]"
//line /usr/local/go/src/crypto/tls/notboring.go:18
}

var fipsSupportedSignatureAlgorithms []SignatureScheme
//line /usr/local/go/src/crypto/tls/notboring.go:20
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/notboring.go:20
var _ = _go_fuzz_dep_.CoverTab
