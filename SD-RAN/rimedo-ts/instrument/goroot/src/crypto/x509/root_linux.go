// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/x509/root_linux.go:5
package x509

//line /usr/local/go/src/crypto/x509/root_linux.go:5
import (
//line /usr/local/go/src/crypto/x509/root_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/x509/root_linux.go:5
)
//line /usr/local/go/src/crypto/x509/root_linux.go:5
import (
//line /usr/local/go/src/crypto/x509/root_linux.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/x509/root_linux.go:5
)

// Possible certificate files; stop after finding one.
var certFiles = []string{
	"/etc/ssl/certs/ca-certificates.crt",
	"/etc/pki/tls/certs/ca-bundle.crt",
	"/etc/ssl/ca-bundle.pem",
	"/etc/pki/tls/cacert.pem",
	"/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem",
	"/etc/ssl/cert.pem",
}

// Possible directories with certificate files; all will be read.
var certDirectories = []string{
	"/etc/ssl/certs",
	"/etc/pki/tls/certs",
	"/system/etc/security/cacerts",
}
//line /usr/local/go/src/crypto/x509/root_linux.go:22
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/x509/root_linux.go:22
var _ = _go_fuzz_dep_.CoverTab
