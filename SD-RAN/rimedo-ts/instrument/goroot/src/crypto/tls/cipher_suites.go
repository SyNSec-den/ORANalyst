// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/tls/cipher_suites.go:5
package tls

//line /usr/local/go/src/crypto/tls/cipher_suites.go:5
import (
//line /usr/local/go/src/crypto/tls/cipher_suites.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:5
)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:5
import (
//line /usr/local/go/src/crypto/tls/cipher_suites.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:5
)

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/internal/boring"
	"crypto/rc4"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"internal/cpu"
	"runtime"

	"golang.org/x/crypto/chacha20poly1305"
)

// CipherSuite is a TLS cipher suite. Note that most functions in this package
//line /usr/local/go/src/crypto/tls/cipher_suites.go:25
// accept and expose cipher suite IDs instead of this type.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:27
type CipherSuite struct {
	ID	uint16
	Name	string

	// Supported versions is the list of TLS protocol versions that can
	// negotiate this cipher suite.
	SupportedVersions	[]uint16

	// Insecure is true if the cipher suite has known security issues
	// due to its primitives, design, or implementation.
	Insecure	bool
}

var (
	supportedUpToTLS12	= []uint16{VersionTLS10, VersionTLS11, VersionTLS12}
	supportedOnlyTLS12	= []uint16{VersionTLS12}
	supportedOnlyTLS13	= []uint16{VersionTLS13}
)

// CipherSuites returns a list of cipher suites currently implemented by this
//line /usr/local/go/src/crypto/tls/cipher_suites.go:46
// package, excluding those with security issues, which are returned by
//line /usr/local/go/src/crypto/tls/cipher_suites.go:46
// InsecureCipherSuites.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:46
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:46
// The list is sorted by ID. Note that the default cipher suites selected by
//line /usr/local/go/src/crypto/tls/cipher_suites.go:46
// this package might depend on logic that can't be captured by a static list,
//line /usr/local/go/src/crypto/tls/cipher_suites.go:46
// and might not match those returned by this function.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:53
func CipherSuites() []*CipherSuite {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:53
	_go_fuzz_dep_.CoverTab[21139]++
								return []*CipherSuite{
		{TLS_RSA_WITH_AES_128_CBC_SHA, "TLS_RSA_WITH_AES_128_CBC_SHA", supportedUpToTLS12, false},
		{TLS_RSA_WITH_AES_256_CBC_SHA, "TLS_RSA_WITH_AES_256_CBC_SHA", supportedUpToTLS12, false},
		{TLS_RSA_WITH_AES_128_GCM_SHA256, "TLS_RSA_WITH_AES_128_GCM_SHA256", supportedOnlyTLS12, false},
		{TLS_RSA_WITH_AES_256_GCM_SHA384, "TLS_RSA_WITH_AES_256_GCM_SHA384", supportedOnlyTLS12, false},

		{TLS_AES_128_GCM_SHA256, "TLS_AES_128_GCM_SHA256", supportedOnlyTLS13, false},
		{TLS_AES_256_GCM_SHA384, "TLS_AES_256_GCM_SHA384", supportedOnlyTLS13, false},
		{TLS_CHACHA20_POLY1305_SHA256, "TLS_CHACHA20_POLY1305_SHA256", supportedOnlyTLS13, false},

		{TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA", supportedUpToTLS12, false},
		{TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA", supportedUpToTLS12, false},
		{TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA", supportedUpToTLS12, false},
		{TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA", supportedUpToTLS12, false},
		{TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", supportedOnlyTLS12, false},
		{TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", supportedOnlyTLS12, false},
		{TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", supportedOnlyTLS12, false},
		{TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384", supportedOnlyTLS12, false},
		{TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256, "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256", supportedOnlyTLS12, false},
		{TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256, "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256", supportedOnlyTLS12, false},
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:74
	// _ = "end of CoverTab[21139]"
}

// InsecureCipherSuites returns a list of cipher suites currently implemented by
//line /usr/local/go/src/crypto/tls/cipher_suites.go:77
// this package and which have security issues.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:77
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:77
// Most applications should not use the cipher suites in this list, and should
//line /usr/local/go/src/crypto/tls/cipher_suites.go:77
// only use those returned by CipherSuites.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:82
func InsecureCipherSuites() []*CipherSuite {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:82
	_go_fuzz_dep_.CoverTab[21140]++

//line /usr/local/go/src/crypto/tls/cipher_suites.go:85
	return []*CipherSuite{
		{TLS_RSA_WITH_RC4_128_SHA, "TLS_RSA_WITH_RC4_128_SHA", supportedUpToTLS12, true},
		{TLS_RSA_WITH_3DES_EDE_CBC_SHA, "TLS_RSA_WITH_3DES_EDE_CBC_SHA", supportedUpToTLS12, true},
		{TLS_RSA_WITH_AES_128_CBC_SHA256, "TLS_RSA_WITH_AES_128_CBC_SHA256", supportedOnlyTLS12, true},
		{TLS_ECDHE_ECDSA_WITH_RC4_128_SHA, "TLS_ECDHE_ECDSA_WITH_RC4_128_SHA", supportedUpToTLS12, true},
		{TLS_ECDHE_RSA_WITH_RC4_128_SHA, "TLS_ECDHE_RSA_WITH_RC4_128_SHA", supportedUpToTLS12, true},
		{TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA, "TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA", supportedUpToTLS12, true},
		{TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256, "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256", supportedOnlyTLS12, true},
		{TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256, "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256", supportedOnlyTLS12, true},
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:94
	// _ = "end of CoverTab[21140]"
}

// CipherSuiteName returns the standard name for the passed cipher suite ID
//line /usr/local/go/src/crypto/tls/cipher_suites.go:97
// (e.g. "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"), or a fallback representation
//line /usr/local/go/src/crypto/tls/cipher_suites.go:97
// of the ID value if the cipher suite is not implemented by this package.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:100
func CipherSuiteName(id uint16) string {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:100
	_go_fuzz_dep_.CoverTab[21141]++
								for _, c := range CipherSuites() {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:101
		_go_fuzz_dep_.CoverTab[21144]++
									if c.ID == id {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:102
			_go_fuzz_dep_.CoverTab[21145]++
										return c.Name
//line /usr/local/go/src/crypto/tls/cipher_suites.go:103
			// _ = "end of CoverTab[21145]"
		} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:104
			_go_fuzz_dep_.CoverTab[21146]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:104
			// _ = "end of CoverTab[21146]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:104
		}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:104
		// _ = "end of CoverTab[21144]"
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:105
	// _ = "end of CoverTab[21141]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:105
	_go_fuzz_dep_.CoverTab[21142]++
								for _, c := range InsecureCipherSuites() {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:106
		_go_fuzz_dep_.CoverTab[21147]++
									if c.ID == id {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:107
			_go_fuzz_dep_.CoverTab[21148]++
										return c.Name
//line /usr/local/go/src/crypto/tls/cipher_suites.go:108
			// _ = "end of CoverTab[21148]"
		} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:109
			_go_fuzz_dep_.CoverTab[21149]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:109
			// _ = "end of CoverTab[21149]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:109
		}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:109
		// _ = "end of CoverTab[21147]"
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:110
	// _ = "end of CoverTab[21142]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:110
	_go_fuzz_dep_.CoverTab[21143]++
								return fmt.Sprintf("0x%04X", id)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:111
	// _ = "end of CoverTab[21143]"
}

const (
	// suiteECDHE indicates that the cipher suite involves elliptic curve
	// Diffie-Hellman. This means that it should only be selected when the
	// client indicates that it supports ECC with a curve and point format
	// that we're happy with.
	suiteECDHE	= 1 << iota
	// suiteECSign indicates that the cipher suite involves an ECDSA or
	// EdDSA signature and therefore may only be selected when the server's
	// certificate is ECDSA or EdDSA. If this is not set then the cipher suite
	// is RSA based.
	suiteECSign
	// suiteTLS12 indicates that the cipher suite should only be advertised
	// and accepted when using TLS 1.2.
	suiteTLS12
	// suiteSHA384 indicates that the cipher suite uses SHA384 as the
	// handshake hash.
	suiteSHA384
)

// A cipherSuite is a TLS 1.0–1.2 cipher suite, and defines the key exchange
//line /usr/local/go/src/crypto/tls/cipher_suites.go:133
// mechanism, as well as the cipher+MAC pair or the AEAD.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:135
type cipherSuite struct {
	id	uint16
	// the lengths, in bytes, of the key material needed for each component.
	keyLen	int
	macLen	int
	ivLen	int
	ka	func(version uint16) keyAgreement
	// flags is a bitmask of the suite* values, above.
	flags	int
	cipher	func(key, iv []byte, isRead bool) any
	mac	func(key []byte) hash.Hash
	aead	func(key, fixedNonce []byte) aead
}

var cipherSuites = []*cipherSuite{
	{TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305, 32, 0, 12, ecdheRSAKA, suiteECDHE | suiteTLS12, nil, nil, aeadChaCha20Poly1305},
	{TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, 32, 0, 12, ecdheECDSAKA, suiteECDHE | suiteECSign | suiteTLS12, nil, nil, aeadChaCha20Poly1305},
	{TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, 16, 0, 4, ecdheRSAKA, suiteECDHE | suiteTLS12, nil, nil, aeadAESGCM},
	{TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, 16, 0, 4, ecdheECDSAKA, suiteECDHE | suiteECSign | suiteTLS12, nil, nil, aeadAESGCM},
	{TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, 32, 0, 4, ecdheRSAKA, suiteECDHE | suiteTLS12 | suiteSHA384, nil, nil, aeadAESGCM},
	{TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, 32, 0, 4, ecdheECDSAKA, suiteECDHE | suiteECSign | suiteTLS12 | suiteSHA384, nil, nil, aeadAESGCM},
	{TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256, 16, 32, 16, ecdheRSAKA, suiteECDHE | suiteTLS12, cipherAES, macSHA256, nil},
	{TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, 16, 20, 16, ecdheRSAKA, suiteECDHE, cipherAES, macSHA1, nil},
	{TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256, 16, 32, 16, ecdheECDSAKA, suiteECDHE | suiteECSign | suiteTLS12, cipherAES, macSHA256, nil},
	{TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, 16, 20, 16, ecdheECDSAKA, suiteECDHE | suiteECSign, cipherAES, macSHA1, nil},
	{TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, 32, 20, 16, ecdheRSAKA, suiteECDHE, cipherAES, macSHA1, nil},
	{TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, 32, 20, 16, ecdheECDSAKA, suiteECDHE | suiteECSign, cipherAES, macSHA1, nil},
	{TLS_RSA_WITH_AES_128_GCM_SHA256, 16, 0, 4, rsaKA, suiteTLS12, nil, nil, aeadAESGCM},
	{TLS_RSA_WITH_AES_256_GCM_SHA384, 32, 0, 4, rsaKA, suiteTLS12 | suiteSHA384, nil, nil, aeadAESGCM},
	{TLS_RSA_WITH_AES_128_CBC_SHA256, 16, 32, 16, rsaKA, suiteTLS12, cipherAES, macSHA256, nil},
	{TLS_RSA_WITH_AES_128_CBC_SHA, 16, 20, 16, rsaKA, 0, cipherAES, macSHA1, nil},
	{TLS_RSA_WITH_AES_256_CBC_SHA, 32, 20, 16, rsaKA, 0, cipherAES, macSHA1, nil},
	{TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA, 24, 20, 8, ecdheRSAKA, suiteECDHE, cipher3DES, macSHA1, nil},
	{TLS_RSA_WITH_3DES_EDE_CBC_SHA, 24, 20, 8, rsaKA, 0, cipher3DES, macSHA1, nil},
	{TLS_RSA_WITH_RC4_128_SHA, 16, 20, 0, rsaKA, 0, cipherRC4, macSHA1, nil},
	{TLS_ECDHE_RSA_WITH_RC4_128_SHA, 16, 20, 0, ecdheRSAKA, suiteECDHE, cipherRC4, macSHA1, nil},
	{TLS_ECDHE_ECDSA_WITH_RC4_128_SHA, 16, 20, 0, ecdheECDSAKA, suiteECDHE | suiteECSign, cipherRC4, macSHA1, nil},
}

// selectCipherSuite returns the first TLS 1.0–1.2 cipher suite from ids which
//line /usr/local/go/src/crypto/tls/cipher_suites.go:174
// is also in supportedIDs and passes the ok filter.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:176
func selectCipherSuite(ids, supportedIDs []uint16, ok func(*cipherSuite) bool) *cipherSuite {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:176
	_go_fuzz_dep_.CoverTab[21150]++
								for _, id := range ids {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:177
		_go_fuzz_dep_.CoverTab[21152]++
									candidate := cipherSuiteByID(id)
									if candidate == nil || func() bool {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:179
			_go_fuzz_dep_.CoverTab[21154]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:179
			return !ok(candidate)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:179
			// _ = "end of CoverTab[21154]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:179
		}() {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:179
			_go_fuzz_dep_.CoverTab[21155]++
										continue
//line /usr/local/go/src/crypto/tls/cipher_suites.go:180
			// _ = "end of CoverTab[21155]"
		} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:181
			_go_fuzz_dep_.CoverTab[21156]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:181
			// _ = "end of CoverTab[21156]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:181
		}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:181
		// _ = "end of CoverTab[21152]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:181
		_go_fuzz_dep_.CoverTab[21153]++

									for _, suppID := range supportedIDs {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:183
			_go_fuzz_dep_.CoverTab[21157]++
										if id == suppID {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:184
				_go_fuzz_dep_.CoverTab[21158]++
											return candidate
//line /usr/local/go/src/crypto/tls/cipher_suites.go:185
				// _ = "end of CoverTab[21158]"
			} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:186
				_go_fuzz_dep_.CoverTab[21159]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:186
				// _ = "end of CoverTab[21159]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:186
			}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:186
			// _ = "end of CoverTab[21157]"
		}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:187
		// _ = "end of CoverTab[21153]"
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:188
	// _ = "end of CoverTab[21150]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:188
	_go_fuzz_dep_.CoverTab[21151]++
								return nil
//line /usr/local/go/src/crypto/tls/cipher_suites.go:189
	// _ = "end of CoverTab[21151]"
}

// A cipherSuiteTLS13 defines only the pair of the AEAD algorithm and hash
//line /usr/local/go/src/crypto/tls/cipher_suites.go:192
// algorithm to be used with HKDF. See RFC 8446, Appendix B.4.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:194
type cipherSuiteTLS13 struct {
	id	uint16
	keyLen	int
	aead	func(key, fixedNonce []byte) aead
	hash	crypto.Hash
}

var cipherSuitesTLS13 = []*cipherSuiteTLS13{
	{TLS_AES_128_GCM_SHA256, 16, aeadAESGCMTLS13, crypto.SHA256},
	{TLS_CHACHA20_POLY1305_SHA256, 32, aeadChaCha20Poly1305, crypto.SHA256},
	{TLS_AES_256_GCM_SHA384, 32, aeadAESGCMTLS13, crypto.SHA384},
}

// cipherSuitesPreferenceOrder is the order in which we'll select (on the
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
// server) or advertise (on the client) TLS 1.0–1.2 cipher suites.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
// Cipher suites are filtered but not reordered based on the application and
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
// peer's preferences, meaning we'll never select a suite lower in this list if
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
// any higher one is available. This makes it more defensible to keep weaker
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
// cipher suites enabled, especially on the server side where we get the last
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
// word, since there are no known downgrade attacks on cipher suites selection.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
// The list is sorted by applying the following priority rules, stopping at the
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
// first (most important) applicable one:
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//   - Anything else comes before RC4
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     RC4 has practically exploitable biases. See https://www.rc4nomore.com.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//   - Anything else comes before CBC_SHA256
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     SHA-256 variants of the CBC ciphersuites don't implement any Lucky13
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     countermeasures. See http://www.isg.rhul.ac.uk/tls/Lucky13.html and
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     https://www.imperialviolet.org/2013/02/04/luckythirteen.html.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//   - Anything else comes before 3DES
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     3DES has 64-bit blocks, which makes it fundamentally susceptible to
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     birthday attacks. See https://sweet32.info.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//   - ECDHE comes before anything else
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     Once we got the broken stuff out of the way, the most important
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     property a cipher suite can have is forward secrecy. We don't
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     implement FFDHE, so that means ECDHE.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//   - AEADs come before CBC ciphers
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     Even with Lucky13 countermeasures, MAC-then-Encrypt CBC cipher suites
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     are fundamentally fragile, and suffered from an endless sequence of
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     padding oracle attacks. See https://eprint.iacr.org/2015/1129,
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     https://www.imperialviolet.org/2014/12/08/poodleagain.html, and
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     https://blog.cloudflare.com/yet-another-padding-oracle-in-openssl-cbc-ciphersuites/.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//   - AES comes before ChaCha20
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     When AES hardware is available, AES-128-GCM and AES-256-GCM are faster
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     than ChaCha20Poly1305.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     When AES hardware is not available, AES-128-GCM is one or more of: much
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     slower, way more complex, and less safe (because not constant time)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     than ChaCha20Poly1305.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     We use this list if we think both peers have AES hardware, and
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     cipherSuitesPreferenceOrderNoAES otherwise.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//   - AES-128 comes before AES-256
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     The only potential advantages of AES-256 are better multi-target
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     margins, and hypothetical post-quantum properties. Neither apply to
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     TLS, and AES-256 is slower due to its four extra rounds (which don't
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     contribute to the advantages above).
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//   - ECDSA comes before RSA
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     The relative order of ECDSA and RSA cipher suites doesn't matter,
//line /usr/local/go/src/crypto/tls/cipher_suites.go:207
//     as they depend on the certificate. Pick one to get a stable order.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:271
var cipherSuitesPreferenceOrder = []uint16{

	TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,

//line /usr/local/go/src/crypto/tls/cipher_suites.go:278
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
	TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,

//line /usr/local/go/src/crypto/tls/cipher_suites.go:282
	TLS_RSA_WITH_AES_128_GCM_SHA256,
								TLS_RSA_WITH_AES_256_GCM_SHA384,

//line /usr/local/go/src/crypto/tls/cipher_suites.go:286
	TLS_RSA_WITH_AES_128_CBC_SHA,
								TLS_RSA_WITH_AES_256_CBC_SHA,

//line /usr/local/go/src/crypto/tls/cipher_suites.go:290
	TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
								TLS_RSA_WITH_3DES_EDE_CBC_SHA,

//line /usr/local/go/src/crypto/tls/cipher_suites.go:294
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
								TLS_RSA_WITH_AES_128_CBC_SHA256,

//line /usr/local/go/src/crypto/tls/cipher_suites.go:298
	TLS_ECDHE_ECDSA_WITH_RC4_128_SHA, TLS_ECDHE_RSA_WITH_RC4_128_SHA,
	TLS_RSA_WITH_RC4_128_SHA,
}

var cipherSuitesPreferenceOrderNoAES = []uint16{

	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,

//line /usr/local/go/src/crypto/tls/cipher_suites.go:307
	TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,

//line /usr/local/go/src/crypto/tls/cipher_suites.go:311
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
	TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	TLS_RSA_WITH_AES_128_GCM_SHA256,
	TLS_RSA_WITH_AES_256_GCM_SHA384,
	TLS_RSA_WITH_AES_128_CBC_SHA,
	TLS_RSA_WITH_AES_256_CBC_SHA,
	TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
	TLS_RSA_WITH_3DES_EDE_CBC_SHA,
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
	TLS_RSA_WITH_AES_128_CBC_SHA256,
	TLS_ECDHE_ECDSA_WITH_RC4_128_SHA, TLS_ECDHE_RSA_WITH_RC4_128_SHA,
	TLS_RSA_WITH_RC4_128_SHA,
}

// disabledCipherSuites are not used unless explicitly listed in
//line /usr/local/go/src/crypto/tls/cipher_suites.go:325
// Config.CipherSuites. They MUST be at the end of cipherSuitesPreferenceOrder.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:327
var disabledCipherSuites = []uint16{

	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
								TLS_RSA_WITH_AES_128_CBC_SHA256,

//line /usr/local/go/src/crypto/tls/cipher_suites.go:333
	TLS_ECDHE_ECDSA_WITH_RC4_128_SHA, TLS_ECDHE_RSA_WITH_RC4_128_SHA,
	TLS_RSA_WITH_RC4_128_SHA,
}

var (
	defaultCipherSuitesLen	= len(cipherSuitesPreferenceOrder) - len(disabledCipherSuites)
	defaultCipherSuites	= cipherSuitesPreferenceOrder[:defaultCipherSuitesLen]
)

// defaultCipherSuitesTLS13 is also the preference order, since there are no
//line /usr/local/go/src/crypto/tls/cipher_suites.go:342
// disabled by default TLS 1.3 cipher suites. The same AES vs ChaCha20 logic as
//line /usr/local/go/src/crypto/tls/cipher_suites.go:342
// cipherSuitesPreferenceOrder applies.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:345
var defaultCipherSuitesTLS13 = []uint16{
	TLS_AES_128_GCM_SHA256,
	TLS_AES_256_GCM_SHA384,
	TLS_CHACHA20_POLY1305_SHA256,
}

var defaultCipherSuitesTLS13NoAES = []uint16{
	TLS_CHACHA20_POLY1305_SHA256,
	TLS_AES_128_GCM_SHA256,
	TLS_AES_256_GCM_SHA384,
}

var (
	hasGCMAsmAMD64	= cpu.X86.HasAES && func() bool {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:358
		_go_fuzz_dep_.CoverTab[21160]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:358
		return cpu.X86.HasPCLMULQDQ
//line /usr/local/go/src/crypto/tls/cipher_suites.go:358
		// _ = "end of CoverTab[21160]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:358
	}()
								hasGCMAsmARM64	= cpu.ARM64.HasAES && func() bool {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:359
		_go_fuzz_dep_.CoverTab[21161]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:359
		return cpu.ARM64.HasPMULL
//line /usr/local/go/src/crypto/tls/cipher_suites.go:359
		// _ = "end of CoverTab[21161]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:359
	}()
	// Keep in sync with crypto/aes/cipher_s390x.go.
	hasGCMAsmS390X	= cpu.S390X.HasAES && func() bool {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:361
		_go_fuzz_dep_.CoverTab[21162]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:361
		return cpu.S390X.HasAESCBC
//line /usr/local/go/src/crypto/tls/cipher_suites.go:361
		// _ = "end of CoverTab[21162]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:361
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:361
		_go_fuzz_dep_.CoverTab[21163]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:361
		return cpu.S390X.HasAESCTR
//line /usr/local/go/src/crypto/tls/cipher_suites.go:361
		// _ = "end of CoverTab[21163]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:361
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:361
		_go_fuzz_dep_.CoverTab[21164]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:361
		return (cpu.S390X.HasGHASH || func() bool {
										_go_fuzz_dep_.CoverTab[21165]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:362
			return cpu.S390X.HasAESGCM
//line /usr/local/go/src/crypto/tls/cipher_suites.go:362
			// _ = "end of CoverTab[21165]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:362
		}())
//line /usr/local/go/src/crypto/tls/cipher_suites.go:362
		// _ = "end of CoverTab[21164]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:362
	}()

	hasAESGCMHardwareSupport	= runtime.GOARCH == "amd64" && func() bool {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:364
		_go_fuzz_dep_.CoverTab[21166]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:364
		return hasGCMAsmAMD64
//line /usr/local/go/src/crypto/tls/cipher_suites.go:364
		// _ = "end of CoverTab[21166]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:364
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:364
		_go_fuzz_dep_.CoverTab[21167]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:364
		return runtime.GOARCH == "arm64" && func() bool {
										_go_fuzz_dep_.CoverTab[21168]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:365
			return hasGCMAsmARM64
//line /usr/local/go/src/crypto/tls/cipher_suites.go:365
			// _ = "end of CoverTab[21168]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:365
		}()
//line /usr/local/go/src/crypto/tls/cipher_suites.go:365
		// _ = "end of CoverTab[21167]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:365
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:365
		_go_fuzz_dep_.CoverTab[21169]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:365
		return runtime.GOARCH == "s390x" && func() bool {
										_go_fuzz_dep_.CoverTab[21170]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:366
			return hasGCMAsmS390X
//line /usr/local/go/src/crypto/tls/cipher_suites.go:366
			// _ = "end of CoverTab[21170]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:366
		}()
//line /usr/local/go/src/crypto/tls/cipher_suites.go:366
		// _ = "end of CoverTab[21169]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:366
	}()
)

var aesgcmCiphers = map[uint16]bool{

	TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:		true,
	TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384:		true,
	TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:	true,
	TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384:	true,

	TLS_AES_128_GCM_SHA256:	true,
	TLS_AES_256_GCM_SHA384:	true,
}

var nonAESGCMAEADCiphers = map[uint16]bool{

	TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305:	true,
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305:	true,

	TLS_CHACHA20_POLY1305_SHA256:	true,
}

// aesgcmPreferred returns whether the first known cipher in the preference list
//line /usr/local/go/src/crypto/tls/cipher_suites.go:388
// is an AES-GCM cipher, implying the peer has hardware support for it.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:390
func aesgcmPreferred(ciphers []uint16) bool {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:390
	_go_fuzz_dep_.CoverTab[21171]++
								for _, cID := range ciphers {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:391
		_go_fuzz_dep_.CoverTab[21173]++
									if c := cipherSuiteByID(cID); c != nil {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:392
			_go_fuzz_dep_.CoverTab[21175]++
										return aesgcmCiphers[cID]
//line /usr/local/go/src/crypto/tls/cipher_suites.go:393
			// _ = "end of CoverTab[21175]"
		} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:394
			_go_fuzz_dep_.CoverTab[21176]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:394
			// _ = "end of CoverTab[21176]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:394
		}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:394
		// _ = "end of CoverTab[21173]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:394
		_go_fuzz_dep_.CoverTab[21174]++
									if c := cipherSuiteTLS13ByID(cID); c != nil {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:395
			_go_fuzz_dep_.CoverTab[21177]++
										return aesgcmCiphers[cID]
//line /usr/local/go/src/crypto/tls/cipher_suites.go:396
			// _ = "end of CoverTab[21177]"
		} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:397
			_go_fuzz_dep_.CoverTab[21178]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:397
			// _ = "end of CoverTab[21178]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:397
		}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:397
		// _ = "end of CoverTab[21174]"
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:398
	// _ = "end of CoverTab[21171]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:398
	_go_fuzz_dep_.CoverTab[21172]++
								return false
//line /usr/local/go/src/crypto/tls/cipher_suites.go:399
	// _ = "end of CoverTab[21172]"
}

func cipherRC4(key, iv []byte, isRead bool) any {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:402
	_go_fuzz_dep_.CoverTab[21179]++
								cipher, _ := rc4.NewCipher(key)
								return cipher
//line /usr/local/go/src/crypto/tls/cipher_suites.go:404
	// _ = "end of CoverTab[21179]"
}

func cipher3DES(key, iv []byte, isRead bool) any {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:407
	_go_fuzz_dep_.CoverTab[21180]++
								block, _ := des.NewTripleDESCipher(key)
								if isRead {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:409
		_go_fuzz_dep_.CoverTab[21182]++
									return cipher.NewCBCDecrypter(block, iv)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:410
		// _ = "end of CoverTab[21182]"
	} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:411
		_go_fuzz_dep_.CoverTab[21183]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:411
		// _ = "end of CoverTab[21183]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:411
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:411
	// _ = "end of CoverTab[21180]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:411
	_go_fuzz_dep_.CoverTab[21181]++
								return cipher.NewCBCEncrypter(block, iv)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:412
	// _ = "end of CoverTab[21181]"
}

func cipherAES(key, iv []byte, isRead bool) any {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:415
	_go_fuzz_dep_.CoverTab[21184]++
								block, _ := aes.NewCipher(key)
								if isRead {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:417
		_go_fuzz_dep_.CoverTab[21186]++
									return cipher.NewCBCDecrypter(block, iv)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:418
		// _ = "end of CoverTab[21186]"
	} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:419
		_go_fuzz_dep_.CoverTab[21187]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:419
		// _ = "end of CoverTab[21187]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:419
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:419
	// _ = "end of CoverTab[21184]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:419
	_go_fuzz_dep_.CoverTab[21185]++
								return cipher.NewCBCEncrypter(block, iv)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:420
	// _ = "end of CoverTab[21185]"
}

// macSHA1 returns a SHA-1 based constant time MAC.
func macSHA1(key []byte) hash.Hash {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:424
	_go_fuzz_dep_.CoverTab[21188]++
								h := sha1.New

//line /usr/local/go/src/crypto/tls/cipher_suites.go:428
	if !boring.Enabled {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:428
		_go_fuzz_dep_.CoverTab[21190]++
									h = newConstantTimeHash(h)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:429
		// _ = "end of CoverTab[21190]"
	} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:430
		_go_fuzz_dep_.CoverTab[21191]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:430
		// _ = "end of CoverTab[21191]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:430
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:430
	// _ = "end of CoverTab[21188]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:430
	_go_fuzz_dep_.CoverTab[21189]++
								return hmac.New(h, key)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:431
	// _ = "end of CoverTab[21189]"
}

// macSHA256 returns a SHA-256 based MAC. This is only supported in TLS 1.2 and
//line /usr/local/go/src/crypto/tls/cipher_suites.go:434
// is currently only used in disabled-by-default cipher suites.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:436
func macSHA256(key []byte) hash.Hash {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:436
	_go_fuzz_dep_.CoverTab[21192]++
								return hmac.New(sha256.New, key)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:437
	// _ = "end of CoverTab[21192]"
}

type aead interface {
	cipher.AEAD

	// explicitNonceLen returns the number of bytes of explicit nonce
	// included in each record. This is eight for older AEADs and
	// zero for modern ones.
	explicitNonceLen() int
}

const (
	aeadNonceLength		= 12
	noncePrefixLength	= 4
)

// prefixNonceAEAD wraps an AEAD and prefixes a fixed portion of the nonce to
//line /usr/local/go/src/crypto/tls/cipher_suites.go:454
// each call.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:456
type prefixNonceAEAD struct {
	// nonce contains the fixed part of the nonce in the first four bytes.
	nonce	[aeadNonceLength]byte
	aead	cipher.AEAD
}

func (f *prefixNonceAEAD) NonceSize() int {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:462
	_go_fuzz_dep_.CoverTab[21193]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:462
	return aeadNonceLength - noncePrefixLength
//line /usr/local/go/src/crypto/tls/cipher_suites.go:462
	// _ = "end of CoverTab[21193]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:462
}
func (f *prefixNonceAEAD) Overhead() int {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:463
	_go_fuzz_dep_.CoverTab[21194]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:463
	return f.aead.Overhead()
//line /usr/local/go/src/crypto/tls/cipher_suites.go:463
	// _ = "end of CoverTab[21194]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:463
}
func (f *prefixNonceAEAD) explicitNonceLen() int {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:464
	_go_fuzz_dep_.CoverTab[21195]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:464
	return f.NonceSize()
//line /usr/local/go/src/crypto/tls/cipher_suites.go:464
	// _ = "end of CoverTab[21195]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:464
}

func (f *prefixNonceAEAD) Seal(out, nonce, plaintext, additionalData []byte) []byte {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:466
	_go_fuzz_dep_.CoverTab[21196]++
								copy(f.nonce[4:], nonce)
								return f.aead.Seal(out, f.nonce[:], plaintext, additionalData)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:468
	// _ = "end of CoverTab[21196]"
}

func (f *prefixNonceAEAD) Open(out, nonce, ciphertext, additionalData []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:471
	_go_fuzz_dep_.CoverTab[21197]++
								copy(f.nonce[4:], nonce)
								return f.aead.Open(out, f.nonce[:], ciphertext, additionalData)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:473
	// _ = "end of CoverTab[21197]"
}

// xorNonceAEAD wraps an AEAD by XORing in a fixed pattern to the nonce
//line /usr/local/go/src/crypto/tls/cipher_suites.go:476
// before each call.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:478
type xorNonceAEAD struct {
	nonceMask	[aeadNonceLength]byte
	aead		cipher.AEAD
}

func (f *xorNonceAEAD) NonceSize() int {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:483
	_go_fuzz_dep_.CoverTab[21198]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:483
	return 8
//line /usr/local/go/src/crypto/tls/cipher_suites.go:483
	// _ = "end of CoverTab[21198]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:483
}
func (f *xorNonceAEAD) Overhead() int {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:484
	_go_fuzz_dep_.CoverTab[21199]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:484
	return f.aead.Overhead()
//line /usr/local/go/src/crypto/tls/cipher_suites.go:484
	// _ = "end of CoverTab[21199]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:484
}
func (f *xorNonceAEAD) explicitNonceLen() int {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:485
	_go_fuzz_dep_.CoverTab[21200]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:485
	return 0
//line /usr/local/go/src/crypto/tls/cipher_suites.go:485
	// _ = "end of CoverTab[21200]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:485
}

func (f *xorNonceAEAD) Seal(out, nonce, plaintext, additionalData []byte) []byte {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:487
	_go_fuzz_dep_.CoverTab[21201]++
								for i, b := range nonce {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:488
		_go_fuzz_dep_.CoverTab[21204]++
									f.nonceMask[4+i] ^= b
//line /usr/local/go/src/crypto/tls/cipher_suites.go:489
		// _ = "end of CoverTab[21204]"
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:490
	// _ = "end of CoverTab[21201]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:490
	_go_fuzz_dep_.CoverTab[21202]++
								result := f.aead.Seal(out, f.nonceMask[:], plaintext, additionalData)
								for i, b := range nonce {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:492
		_go_fuzz_dep_.CoverTab[21205]++
									f.nonceMask[4+i] ^= b
//line /usr/local/go/src/crypto/tls/cipher_suites.go:493
		// _ = "end of CoverTab[21205]"
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:494
	// _ = "end of CoverTab[21202]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:494
	_go_fuzz_dep_.CoverTab[21203]++

								return result
//line /usr/local/go/src/crypto/tls/cipher_suites.go:496
	// _ = "end of CoverTab[21203]"
}

func (f *xorNonceAEAD) Open(out, nonce, ciphertext, additionalData []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:499
	_go_fuzz_dep_.CoverTab[21206]++
								for i, b := range nonce {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:500
		_go_fuzz_dep_.CoverTab[21209]++
									f.nonceMask[4+i] ^= b
//line /usr/local/go/src/crypto/tls/cipher_suites.go:501
		// _ = "end of CoverTab[21209]"
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:502
	// _ = "end of CoverTab[21206]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:502
	_go_fuzz_dep_.CoverTab[21207]++
								result, err := f.aead.Open(out, f.nonceMask[:], ciphertext, additionalData)
								for i, b := range nonce {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:504
		_go_fuzz_dep_.CoverTab[21210]++
									f.nonceMask[4+i] ^= b
//line /usr/local/go/src/crypto/tls/cipher_suites.go:505
		// _ = "end of CoverTab[21210]"
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:506
	// _ = "end of CoverTab[21207]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:506
	_go_fuzz_dep_.CoverTab[21208]++

								return result, err
//line /usr/local/go/src/crypto/tls/cipher_suites.go:508
	// _ = "end of CoverTab[21208]"
}

func aeadAESGCM(key, noncePrefix []byte) aead {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:511
	_go_fuzz_dep_.CoverTab[21211]++
								if len(noncePrefix) != noncePrefixLength {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:512
		_go_fuzz_dep_.CoverTab[21216]++
									panic("tls: internal error: wrong nonce length")
//line /usr/local/go/src/crypto/tls/cipher_suites.go:513
		// _ = "end of CoverTab[21216]"
	} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:514
		_go_fuzz_dep_.CoverTab[21217]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:514
		// _ = "end of CoverTab[21217]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:514
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:514
	// _ = "end of CoverTab[21211]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:514
	_go_fuzz_dep_.CoverTab[21212]++
								aes, err := aes.NewCipher(key)
								if err != nil {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:516
		_go_fuzz_dep_.CoverTab[21218]++
									panic(err)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:517
		// _ = "end of CoverTab[21218]"
	} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:518
		_go_fuzz_dep_.CoverTab[21219]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:518
		// _ = "end of CoverTab[21219]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:518
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:518
	// _ = "end of CoverTab[21212]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:518
	_go_fuzz_dep_.CoverTab[21213]++
								var aead cipher.AEAD
								if boring.Enabled {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:520
		_go_fuzz_dep_.CoverTab[21220]++
									aead, err = boring.NewGCMTLS(aes)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:521
		// _ = "end of CoverTab[21220]"
	} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:522
		_go_fuzz_dep_.CoverTab[21221]++
									boring.Unreachable()
									aead, err = cipher.NewGCM(aes)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:524
		// _ = "end of CoverTab[21221]"
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:525
	// _ = "end of CoverTab[21213]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:525
	_go_fuzz_dep_.CoverTab[21214]++
								if err != nil {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:526
		_go_fuzz_dep_.CoverTab[21222]++
									panic(err)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:527
		// _ = "end of CoverTab[21222]"
	} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:528
		_go_fuzz_dep_.CoverTab[21223]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:528
		// _ = "end of CoverTab[21223]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:528
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:528
	// _ = "end of CoverTab[21214]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:528
	_go_fuzz_dep_.CoverTab[21215]++

								ret := &prefixNonceAEAD{aead: aead}
								copy(ret.nonce[:], noncePrefix)
								return ret
//line /usr/local/go/src/crypto/tls/cipher_suites.go:532
	// _ = "end of CoverTab[21215]"
}

func aeadAESGCMTLS13(key, nonceMask []byte) aead {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:535
	_go_fuzz_dep_.CoverTab[21224]++
								if len(nonceMask) != aeadNonceLength {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:536
		_go_fuzz_dep_.CoverTab[21228]++
									panic("tls: internal error: wrong nonce length")
//line /usr/local/go/src/crypto/tls/cipher_suites.go:537
		// _ = "end of CoverTab[21228]"
	} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:538
		_go_fuzz_dep_.CoverTab[21229]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:538
		// _ = "end of CoverTab[21229]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:538
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:538
	// _ = "end of CoverTab[21224]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:538
	_go_fuzz_dep_.CoverTab[21225]++
								aes, err := aes.NewCipher(key)
								if err != nil {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:540
		_go_fuzz_dep_.CoverTab[21230]++
									panic(err)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:541
		// _ = "end of CoverTab[21230]"
	} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:542
		_go_fuzz_dep_.CoverTab[21231]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:542
		// _ = "end of CoverTab[21231]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:542
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:542
	// _ = "end of CoverTab[21225]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:542
	_go_fuzz_dep_.CoverTab[21226]++
								aead, err := cipher.NewGCM(aes)
								if err != nil {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:544
		_go_fuzz_dep_.CoverTab[21232]++
									panic(err)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:545
		// _ = "end of CoverTab[21232]"
	} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:546
		_go_fuzz_dep_.CoverTab[21233]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:546
		// _ = "end of CoverTab[21233]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:546
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:546
	// _ = "end of CoverTab[21226]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:546
	_go_fuzz_dep_.CoverTab[21227]++

								ret := &xorNonceAEAD{aead: aead}
								copy(ret.nonceMask[:], nonceMask)
								return ret
//line /usr/local/go/src/crypto/tls/cipher_suites.go:550
	// _ = "end of CoverTab[21227]"
}

func aeadChaCha20Poly1305(key, nonceMask []byte) aead {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:553
	_go_fuzz_dep_.CoverTab[21234]++
								if len(nonceMask) != aeadNonceLength {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:554
		_go_fuzz_dep_.CoverTab[21237]++
									panic("tls: internal error: wrong nonce length")
//line /usr/local/go/src/crypto/tls/cipher_suites.go:555
		// _ = "end of CoverTab[21237]"
	} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:556
		_go_fuzz_dep_.CoverTab[21238]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:556
		// _ = "end of CoverTab[21238]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:556
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:556
	// _ = "end of CoverTab[21234]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:556
	_go_fuzz_dep_.CoverTab[21235]++
								aead, err := chacha20poly1305.New(key)
								if err != nil {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:558
		_go_fuzz_dep_.CoverTab[21239]++
									panic(err)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:559
		// _ = "end of CoverTab[21239]"
	} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:560
		_go_fuzz_dep_.CoverTab[21240]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:560
		// _ = "end of CoverTab[21240]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:560
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:560
	// _ = "end of CoverTab[21235]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:560
	_go_fuzz_dep_.CoverTab[21236]++

								ret := &xorNonceAEAD{aead: aead}
								copy(ret.nonceMask[:], nonceMask)
								return ret
//line /usr/local/go/src/crypto/tls/cipher_suites.go:564
	// _ = "end of CoverTab[21236]"
}

type constantTimeHash interface {
	hash.Hash
	ConstantTimeSum(b []byte) []byte
}

// cthWrapper wraps any hash.Hash that implements ConstantTimeSum, and replaces
//line /usr/local/go/src/crypto/tls/cipher_suites.go:572
// with that all calls to Sum. It's used to obtain a ConstantTimeSum-based HMAC.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:574
type cthWrapper struct {
	h constantTimeHash
}

func (c *cthWrapper) Size() int {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:578
	_go_fuzz_dep_.CoverTab[21241]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:578
	return c.h.Size()
//line /usr/local/go/src/crypto/tls/cipher_suites.go:578
	// _ = "end of CoverTab[21241]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:578
}
func (c *cthWrapper) BlockSize() int {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:579
	_go_fuzz_dep_.CoverTab[21242]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:579
	return c.h.BlockSize()
//line /usr/local/go/src/crypto/tls/cipher_suites.go:579
	// _ = "end of CoverTab[21242]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:579
}
func (c *cthWrapper) Reset() {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:580
	_go_fuzz_dep_.CoverTab[21243]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:580
	c.h.Reset()
//line /usr/local/go/src/crypto/tls/cipher_suites.go:580
	// _ = "end of CoverTab[21243]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:580
}
func (c *cthWrapper) Write(p []byte) (int, error) {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:581
	_go_fuzz_dep_.CoverTab[21244]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:581
	return c.h.Write(p)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:581
	// _ = "end of CoverTab[21244]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:581
}
func (c *cthWrapper) Sum(b []byte) []byte {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:582
	_go_fuzz_dep_.CoverTab[21245]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:582
	return c.h.ConstantTimeSum(b)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:582
	// _ = "end of CoverTab[21245]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:582
}

func newConstantTimeHash(h func() hash.Hash) func() hash.Hash {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:584
	_go_fuzz_dep_.CoverTab[21246]++
								boring.Unreachable()
								return func() hash.Hash {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:586
		_go_fuzz_dep_.CoverTab[21247]++
									return &cthWrapper{h().(constantTimeHash)}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:587
		// _ = "end of CoverTab[21247]"
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:588
	// _ = "end of CoverTab[21246]"
}

// tls10MAC implements the TLS 1.0 MAC function. RFC 2246, Section 6.2.3.
func tls10MAC(h hash.Hash, out, seq, header, data, extra []byte) []byte {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:592
	_go_fuzz_dep_.CoverTab[21248]++
								h.Reset()
								h.Write(seq)
								h.Write(header)
								h.Write(data)
								res := h.Sum(out)
								if extra != nil {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:598
		_go_fuzz_dep_.CoverTab[21250]++
									h.Write(extra)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:599
		// _ = "end of CoverTab[21250]"
	} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:600
		_go_fuzz_dep_.CoverTab[21251]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:600
		// _ = "end of CoverTab[21251]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:600
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:600
	// _ = "end of CoverTab[21248]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:600
	_go_fuzz_dep_.CoverTab[21249]++
								return res
//line /usr/local/go/src/crypto/tls/cipher_suites.go:601
	// _ = "end of CoverTab[21249]"
}

func rsaKA(version uint16) keyAgreement {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:604
	_go_fuzz_dep_.CoverTab[21252]++
								return rsaKeyAgreement{}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:605
	// _ = "end of CoverTab[21252]"
}

func ecdheECDSAKA(version uint16) keyAgreement {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:608
	_go_fuzz_dep_.CoverTab[21253]++
								return &ecdheKeyAgreement{
		isRSA:		false,
		version:	version,
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:612
	// _ = "end of CoverTab[21253]"
}

func ecdheRSAKA(version uint16) keyAgreement {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:615
	_go_fuzz_dep_.CoverTab[21254]++
								return &ecdheKeyAgreement{
		isRSA:		true,
		version:	version,
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:619
	// _ = "end of CoverTab[21254]"
}

// mutualCipherSuite returns a cipherSuite given a list of supported
//line /usr/local/go/src/crypto/tls/cipher_suites.go:622
// ciphersuites and the id requested by the peer.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:624
func mutualCipherSuite(have []uint16, want uint16) *cipherSuite {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:624
	_go_fuzz_dep_.CoverTab[21255]++
								for _, id := range have {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:625
		_go_fuzz_dep_.CoverTab[21257]++
									if id == want {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:626
			_go_fuzz_dep_.CoverTab[21258]++
										return cipherSuiteByID(id)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:627
			// _ = "end of CoverTab[21258]"
		} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:628
			_go_fuzz_dep_.CoverTab[21259]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:628
			// _ = "end of CoverTab[21259]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:628
		}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:628
		// _ = "end of CoverTab[21257]"
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:629
	// _ = "end of CoverTab[21255]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:629
	_go_fuzz_dep_.CoverTab[21256]++
								return nil
//line /usr/local/go/src/crypto/tls/cipher_suites.go:630
	// _ = "end of CoverTab[21256]"
}

func cipherSuiteByID(id uint16) *cipherSuite {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:633
	_go_fuzz_dep_.CoverTab[21260]++
								for _, cipherSuite := range cipherSuites {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:634
		_go_fuzz_dep_.CoverTab[21262]++
									if cipherSuite.id == id {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:635
			_go_fuzz_dep_.CoverTab[21263]++
										return cipherSuite
//line /usr/local/go/src/crypto/tls/cipher_suites.go:636
			// _ = "end of CoverTab[21263]"
		} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:637
			_go_fuzz_dep_.CoverTab[21264]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:637
			// _ = "end of CoverTab[21264]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:637
		}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:637
		// _ = "end of CoverTab[21262]"
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:638
	// _ = "end of CoverTab[21260]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:638
	_go_fuzz_dep_.CoverTab[21261]++
								return nil
//line /usr/local/go/src/crypto/tls/cipher_suites.go:639
	// _ = "end of CoverTab[21261]"
}

func mutualCipherSuiteTLS13(have []uint16, want uint16) *cipherSuiteTLS13 {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:642
	_go_fuzz_dep_.CoverTab[21265]++
								for _, id := range have {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:643
		_go_fuzz_dep_.CoverTab[21267]++
									if id == want {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:644
			_go_fuzz_dep_.CoverTab[21268]++
										return cipherSuiteTLS13ByID(id)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:645
			// _ = "end of CoverTab[21268]"
		} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:646
			_go_fuzz_dep_.CoverTab[21269]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:646
			// _ = "end of CoverTab[21269]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:646
		}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:646
		// _ = "end of CoverTab[21267]"
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:647
	// _ = "end of CoverTab[21265]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:647
	_go_fuzz_dep_.CoverTab[21266]++
								return nil
//line /usr/local/go/src/crypto/tls/cipher_suites.go:648
	// _ = "end of CoverTab[21266]"
}

func cipherSuiteTLS13ByID(id uint16) *cipherSuiteTLS13 {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:651
	_go_fuzz_dep_.CoverTab[21270]++
								for _, cipherSuite := range cipherSuitesTLS13 {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:652
		_go_fuzz_dep_.CoverTab[21272]++
									if cipherSuite.id == id {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:653
			_go_fuzz_dep_.CoverTab[21273]++
										return cipherSuite
//line /usr/local/go/src/crypto/tls/cipher_suites.go:654
			// _ = "end of CoverTab[21273]"
		} else {
//line /usr/local/go/src/crypto/tls/cipher_suites.go:655
			_go_fuzz_dep_.CoverTab[21274]++
//line /usr/local/go/src/crypto/tls/cipher_suites.go:655
			// _ = "end of CoverTab[21274]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:655
		}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:655
		// _ = "end of CoverTab[21272]"
	}
//line /usr/local/go/src/crypto/tls/cipher_suites.go:656
	// _ = "end of CoverTab[21270]"
//line /usr/local/go/src/crypto/tls/cipher_suites.go:656
	_go_fuzz_dep_.CoverTab[21271]++
								return nil
//line /usr/local/go/src/crypto/tls/cipher_suites.go:657
	// _ = "end of CoverTab[21271]"
}

// A list of cipher suite IDs that are, or have been, implemented by this
//line /usr/local/go/src/crypto/tls/cipher_suites.go:660
// package.
//line /usr/local/go/src/crypto/tls/cipher_suites.go:660
//
//line /usr/local/go/src/crypto/tls/cipher_suites.go:660
// See https://www.iana.org/assignments/tls-parameters/tls-parameters.xml
//line /usr/local/go/src/crypto/tls/cipher_suites.go:664
const (
	// TLS 1.0 - 1.2 cipher suites.
	TLS_RSA_WITH_RC4_128_SHA			uint16	= 0x0005
	TLS_RSA_WITH_3DES_EDE_CBC_SHA			uint16	= 0x000a
	TLS_RSA_WITH_AES_128_CBC_SHA			uint16	= 0x002f
	TLS_RSA_WITH_AES_256_CBC_SHA			uint16	= 0x0035
	TLS_RSA_WITH_AES_128_CBC_SHA256			uint16	= 0x003c
	TLS_RSA_WITH_AES_128_GCM_SHA256			uint16	= 0x009c
	TLS_RSA_WITH_AES_256_GCM_SHA384			uint16	= 0x009d
	TLS_ECDHE_ECDSA_WITH_RC4_128_SHA		uint16	= 0xc007
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA		uint16	= 0xc009
	TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA		uint16	= 0xc00a
	TLS_ECDHE_RSA_WITH_RC4_128_SHA			uint16	= 0xc011
	TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA		uint16	= 0xc012
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA		uint16	= 0xc013
	TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA		uint16	= 0xc014
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256		uint16	= 0xc023
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256		uint16	= 0xc027
	TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256		uint16	= 0xc02f
	TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256		uint16	= 0xc02b
	TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384		uint16	= 0xc030
	TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384		uint16	= 0xc02c
	TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256	uint16	= 0xcca8
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256	uint16	= 0xcca9

	// TLS 1.3 cipher suites.
	TLS_AES_128_GCM_SHA256		uint16	= 0x1301
	TLS_AES_256_GCM_SHA384		uint16	= 0x1302
	TLS_CHACHA20_POLY1305_SHA256	uint16	= 0x1303

	// TLS_FALLBACK_SCSV isn't a standard cipher suite but an indicator
	// that the client is doing version fallback. See RFC 7507.
	TLS_FALLBACK_SCSV	uint16	= 0x5600

	// Legacy names for the corresponding cipher suites with the correct _SHA256
	// suffix, retained for backward compatibility.
	TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305	= TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305	= TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
)

//line /usr/local/go/src/crypto/tls/cipher_suites.go:702
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/cipher_suites.go:702
var _ = _go_fuzz_dep_.CoverTab
