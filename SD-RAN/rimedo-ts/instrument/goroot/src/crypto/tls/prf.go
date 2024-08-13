// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/tls/prf.go:5
package tls

//line /usr/local/go/src/crypto/tls/prf.go:5
import (
//line /usr/local/go/src/crypto/tls/prf.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/prf.go:5
)
//line /usr/local/go/src/crypto/tls/prf.go:5
import (
//line /usr/local/go/src/crypto/tls/prf.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/prf.go:5
)

import (
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"hash"
)

// Split a premaster secret in two as specified in RFC 4346, Section 5.
func splitPreMasterSecret(secret []byte) (s1, s2 []byte) {
//line /usr/local/go/src/crypto/tls/prf.go:20
	_go_fuzz_dep_.CoverTab[24928]++
						s1 = secret[0 : (len(secret)+1)/2]
						s2 = secret[len(secret)/2:]
						return
//line /usr/local/go/src/crypto/tls/prf.go:23
	// _ = "end of CoverTab[24928]"
}

// pHash implements the P_hash function, as defined in RFC 4346, Section 5.
func pHash(result, secret, seed []byte, hash func() hash.Hash) {
//line /usr/local/go/src/crypto/tls/prf.go:27
	_go_fuzz_dep_.CoverTab[24929]++
						h := hmac.New(hash, secret)
						h.Write(seed)
						a := h.Sum(nil)

						j := 0
						for j < len(result) {
//line /usr/local/go/src/crypto/tls/prf.go:33
		_go_fuzz_dep_.CoverTab[24930]++
							h.Reset()
							h.Write(a)
							h.Write(seed)
							b := h.Sum(nil)
							copy(result[j:], b)
							j += len(b)

							h.Reset()
							h.Write(a)
							a = h.Sum(nil)
//line /usr/local/go/src/crypto/tls/prf.go:43
		// _ = "end of CoverTab[24930]"
	}
//line /usr/local/go/src/crypto/tls/prf.go:44
	// _ = "end of CoverTab[24929]"
}

// prf10 implements the TLS 1.0 pseudo-random function, as defined in RFC 2246, Section 5.
func prf10(result, secret, label, seed []byte) {
//line /usr/local/go/src/crypto/tls/prf.go:48
	_go_fuzz_dep_.CoverTab[24931]++
						hashSHA1 := sha1.New
						hashMD5 := md5.New

						labelAndSeed := make([]byte, len(label)+len(seed))
						copy(labelAndSeed, label)
						copy(labelAndSeed[len(label):], seed)

						s1, s2 := splitPreMasterSecret(secret)
						pHash(result, s1, labelAndSeed, hashMD5)
						result2 := make([]byte, len(result))
						pHash(result2, s2, labelAndSeed, hashSHA1)

						for i, b := range result2 {
//line /usr/local/go/src/crypto/tls/prf.go:61
		_go_fuzz_dep_.CoverTab[24932]++
							result[i] ^= b
//line /usr/local/go/src/crypto/tls/prf.go:62
		// _ = "end of CoverTab[24932]"
	}
//line /usr/local/go/src/crypto/tls/prf.go:63
	// _ = "end of CoverTab[24931]"
}

// prf12 implements the TLS 1.2 pseudo-random function, as defined in RFC 5246, Section 5.
func prf12(hashFunc func() hash.Hash) func(result, secret, label, seed []byte) {
//line /usr/local/go/src/crypto/tls/prf.go:67
	_go_fuzz_dep_.CoverTab[24933]++
						return func(result, secret, label, seed []byte) {
//line /usr/local/go/src/crypto/tls/prf.go:68
		_go_fuzz_dep_.CoverTab[24934]++
							labelAndSeed := make([]byte, len(label)+len(seed))
							copy(labelAndSeed, label)
							copy(labelAndSeed[len(label):], seed)

							pHash(result, secret, labelAndSeed, hashFunc)
//line /usr/local/go/src/crypto/tls/prf.go:73
		// _ = "end of CoverTab[24934]"
	}
//line /usr/local/go/src/crypto/tls/prf.go:74
	// _ = "end of CoverTab[24933]"
}

const (
	masterSecretLength	= 48	// Length of a master secret in TLS 1.1.
	finishedVerifyLength	= 12	// Length of verify_data in a Finished message.
)

var masterSecretLabel = []byte("master secret")
var keyExpansionLabel = []byte("key expansion")
var clientFinishedLabel = []byte("client finished")
var serverFinishedLabel = []byte("server finished")

func prfAndHashForVersion(version uint16, suite *cipherSuite) (func(result, secret, label, seed []byte), crypto.Hash) {
//line /usr/local/go/src/crypto/tls/prf.go:87
	_go_fuzz_dep_.CoverTab[24935]++
						switch version {
	case VersionTLS10, VersionTLS11:
//line /usr/local/go/src/crypto/tls/prf.go:89
		_go_fuzz_dep_.CoverTab[24936]++
							return prf10, crypto.Hash(0)
//line /usr/local/go/src/crypto/tls/prf.go:90
		// _ = "end of CoverTab[24936]"
	case VersionTLS12:
//line /usr/local/go/src/crypto/tls/prf.go:91
		_go_fuzz_dep_.CoverTab[24937]++
							if suite.flags&suiteSHA384 != 0 {
//line /usr/local/go/src/crypto/tls/prf.go:92
			_go_fuzz_dep_.CoverTab[24940]++
								return prf12(sha512.New384), crypto.SHA384
//line /usr/local/go/src/crypto/tls/prf.go:93
			// _ = "end of CoverTab[24940]"
		} else {
//line /usr/local/go/src/crypto/tls/prf.go:94
			_go_fuzz_dep_.CoverTab[24941]++
//line /usr/local/go/src/crypto/tls/prf.go:94
			// _ = "end of CoverTab[24941]"
//line /usr/local/go/src/crypto/tls/prf.go:94
		}
//line /usr/local/go/src/crypto/tls/prf.go:94
		// _ = "end of CoverTab[24937]"
//line /usr/local/go/src/crypto/tls/prf.go:94
		_go_fuzz_dep_.CoverTab[24938]++
							return prf12(sha256.New), crypto.SHA256
//line /usr/local/go/src/crypto/tls/prf.go:95
		// _ = "end of CoverTab[24938]"
	default:
//line /usr/local/go/src/crypto/tls/prf.go:96
		_go_fuzz_dep_.CoverTab[24939]++
							panic("unknown version")
//line /usr/local/go/src/crypto/tls/prf.go:97
		// _ = "end of CoverTab[24939]"
	}
//line /usr/local/go/src/crypto/tls/prf.go:98
	// _ = "end of CoverTab[24935]"
}

func prfForVersion(version uint16, suite *cipherSuite) func(result, secret, label, seed []byte) {
//line /usr/local/go/src/crypto/tls/prf.go:101
	_go_fuzz_dep_.CoverTab[24942]++
						prf, _ := prfAndHashForVersion(version, suite)
						return prf
//line /usr/local/go/src/crypto/tls/prf.go:103
	// _ = "end of CoverTab[24942]"
}

// masterFromPreMasterSecret generates the master secret from the pre-master
//line /usr/local/go/src/crypto/tls/prf.go:106
// secret. See RFC 5246, Section 8.1.
//line /usr/local/go/src/crypto/tls/prf.go:108
func masterFromPreMasterSecret(version uint16, suite *cipherSuite, preMasterSecret, clientRandom, serverRandom []byte) []byte {
//line /usr/local/go/src/crypto/tls/prf.go:108
	_go_fuzz_dep_.CoverTab[24943]++
						seed := make([]byte, 0, len(clientRandom)+len(serverRandom))
						seed = append(seed, clientRandom...)
						seed = append(seed, serverRandom...)

						masterSecret := make([]byte, masterSecretLength)
						prfForVersion(version, suite)(masterSecret, preMasterSecret, masterSecretLabel, seed)
						return masterSecret
//line /usr/local/go/src/crypto/tls/prf.go:115
	// _ = "end of CoverTab[24943]"
}

// keysFromMasterSecret generates the connection keys from the master
//line /usr/local/go/src/crypto/tls/prf.go:118
// secret, given the lengths of the MAC key, cipher key and IV, as defined in
//line /usr/local/go/src/crypto/tls/prf.go:118
// RFC 2246, Section 6.3.
//line /usr/local/go/src/crypto/tls/prf.go:121
func keysFromMasterSecret(version uint16, suite *cipherSuite, masterSecret, clientRandom, serverRandom []byte, macLen, keyLen, ivLen int) (clientMAC, serverMAC, clientKey, serverKey, clientIV, serverIV []byte) {
//line /usr/local/go/src/crypto/tls/prf.go:121
	_go_fuzz_dep_.CoverTab[24944]++
						seed := make([]byte, 0, len(serverRandom)+len(clientRandom))
						seed = append(seed, serverRandom...)
						seed = append(seed, clientRandom...)

						n := 2*macLen + 2*keyLen + 2*ivLen
						keyMaterial := make([]byte, n)
						prfForVersion(version, suite)(keyMaterial, masterSecret, keyExpansionLabel, seed)
						clientMAC = keyMaterial[:macLen]
						keyMaterial = keyMaterial[macLen:]
						serverMAC = keyMaterial[:macLen]
						keyMaterial = keyMaterial[macLen:]
						clientKey = keyMaterial[:keyLen]
						keyMaterial = keyMaterial[keyLen:]
						serverKey = keyMaterial[:keyLen]
						keyMaterial = keyMaterial[keyLen:]
						clientIV = keyMaterial[:ivLen]
						keyMaterial = keyMaterial[ivLen:]
						serverIV = keyMaterial[:ivLen]
						return
//line /usr/local/go/src/crypto/tls/prf.go:140
	// _ = "end of CoverTab[24944]"
}

func newFinishedHash(version uint16, cipherSuite *cipherSuite) finishedHash {
//line /usr/local/go/src/crypto/tls/prf.go:143
	_go_fuzz_dep_.CoverTab[24945]++
						var buffer []byte
						if version >= VersionTLS12 {
//line /usr/local/go/src/crypto/tls/prf.go:145
		_go_fuzz_dep_.CoverTab[24948]++
							buffer = []byte{}
//line /usr/local/go/src/crypto/tls/prf.go:146
		// _ = "end of CoverTab[24948]"
	} else {
//line /usr/local/go/src/crypto/tls/prf.go:147
		_go_fuzz_dep_.CoverTab[24949]++
//line /usr/local/go/src/crypto/tls/prf.go:147
		// _ = "end of CoverTab[24949]"
//line /usr/local/go/src/crypto/tls/prf.go:147
	}
//line /usr/local/go/src/crypto/tls/prf.go:147
	// _ = "end of CoverTab[24945]"
//line /usr/local/go/src/crypto/tls/prf.go:147
	_go_fuzz_dep_.CoverTab[24946]++

						prf, hash := prfAndHashForVersion(version, cipherSuite)
						if hash != 0 {
//line /usr/local/go/src/crypto/tls/prf.go:150
		_go_fuzz_dep_.CoverTab[24950]++
							return finishedHash{hash.New(), hash.New(), nil, nil, buffer, version, prf}
//line /usr/local/go/src/crypto/tls/prf.go:151
		// _ = "end of CoverTab[24950]"
	} else {
//line /usr/local/go/src/crypto/tls/prf.go:152
		_go_fuzz_dep_.CoverTab[24951]++
//line /usr/local/go/src/crypto/tls/prf.go:152
		// _ = "end of CoverTab[24951]"
//line /usr/local/go/src/crypto/tls/prf.go:152
	}
//line /usr/local/go/src/crypto/tls/prf.go:152
	// _ = "end of CoverTab[24946]"
//line /usr/local/go/src/crypto/tls/prf.go:152
	_go_fuzz_dep_.CoverTab[24947]++

						return finishedHash{sha1.New(), sha1.New(), md5.New(), md5.New(), buffer, version, prf}
//line /usr/local/go/src/crypto/tls/prf.go:154
	// _ = "end of CoverTab[24947]"
}

// A finishedHash calculates the hash of a set of handshake messages suitable
//line /usr/local/go/src/crypto/tls/prf.go:157
// for including in a Finished message.
//line /usr/local/go/src/crypto/tls/prf.go:159
type finishedHash struct {
	client	hash.Hash
	server	hash.Hash

	// Prior to TLS 1.2, an additional MD5 hash is required.
	clientMD5	hash.Hash
	serverMD5	hash.Hash

	// In TLS 1.2, a full buffer is sadly required.
	buffer	[]byte

	version	uint16
	prf	func(result, secret, label, seed []byte)
}

func (h *finishedHash) Write(msg []byte) (n int, err error) {
//line /usr/local/go/src/crypto/tls/prf.go:174
	_go_fuzz_dep_.CoverTab[24952]++
						h.client.Write(msg)
						h.server.Write(msg)

						if h.version < VersionTLS12 {
//line /usr/local/go/src/crypto/tls/prf.go:178
		_go_fuzz_dep_.CoverTab[24955]++
							h.clientMD5.Write(msg)
							h.serverMD5.Write(msg)
//line /usr/local/go/src/crypto/tls/prf.go:180
		// _ = "end of CoverTab[24955]"
	} else {
//line /usr/local/go/src/crypto/tls/prf.go:181
		_go_fuzz_dep_.CoverTab[24956]++
//line /usr/local/go/src/crypto/tls/prf.go:181
		// _ = "end of CoverTab[24956]"
//line /usr/local/go/src/crypto/tls/prf.go:181
	}
//line /usr/local/go/src/crypto/tls/prf.go:181
	// _ = "end of CoverTab[24952]"
//line /usr/local/go/src/crypto/tls/prf.go:181
	_go_fuzz_dep_.CoverTab[24953]++

						if h.buffer != nil {
//line /usr/local/go/src/crypto/tls/prf.go:183
		_go_fuzz_dep_.CoverTab[24957]++
							h.buffer = append(h.buffer, msg...)
//line /usr/local/go/src/crypto/tls/prf.go:184
		// _ = "end of CoverTab[24957]"
	} else {
//line /usr/local/go/src/crypto/tls/prf.go:185
		_go_fuzz_dep_.CoverTab[24958]++
//line /usr/local/go/src/crypto/tls/prf.go:185
		// _ = "end of CoverTab[24958]"
//line /usr/local/go/src/crypto/tls/prf.go:185
	}
//line /usr/local/go/src/crypto/tls/prf.go:185
	// _ = "end of CoverTab[24953]"
//line /usr/local/go/src/crypto/tls/prf.go:185
	_go_fuzz_dep_.CoverTab[24954]++

						return len(msg), nil
//line /usr/local/go/src/crypto/tls/prf.go:187
	// _ = "end of CoverTab[24954]"
}

func (h finishedHash) Sum() []byte {
//line /usr/local/go/src/crypto/tls/prf.go:190
	_go_fuzz_dep_.CoverTab[24959]++
						if h.version >= VersionTLS12 {
//line /usr/local/go/src/crypto/tls/prf.go:191
		_go_fuzz_dep_.CoverTab[24961]++
							return h.client.Sum(nil)
//line /usr/local/go/src/crypto/tls/prf.go:192
		// _ = "end of CoverTab[24961]"
	} else {
//line /usr/local/go/src/crypto/tls/prf.go:193
		_go_fuzz_dep_.CoverTab[24962]++
//line /usr/local/go/src/crypto/tls/prf.go:193
		// _ = "end of CoverTab[24962]"
//line /usr/local/go/src/crypto/tls/prf.go:193
	}
//line /usr/local/go/src/crypto/tls/prf.go:193
	// _ = "end of CoverTab[24959]"
//line /usr/local/go/src/crypto/tls/prf.go:193
	_go_fuzz_dep_.CoverTab[24960]++

						out := make([]byte, 0, md5.Size+sha1.Size)
						out = h.clientMD5.Sum(out)
						return h.client.Sum(out)
//line /usr/local/go/src/crypto/tls/prf.go:197
	// _ = "end of CoverTab[24960]"
}

// clientSum returns the contents of the verify_data member of a client's
//line /usr/local/go/src/crypto/tls/prf.go:200
// Finished message.
//line /usr/local/go/src/crypto/tls/prf.go:202
func (h finishedHash) clientSum(masterSecret []byte) []byte {
//line /usr/local/go/src/crypto/tls/prf.go:202
	_go_fuzz_dep_.CoverTab[24963]++
						out := make([]byte, finishedVerifyLength)
						h.prf(out, masterSecret, clientFinishedLabel, h.Sum())
						return out
//line /usr/local/go/src/crypto/tls/prf.go:205
	// _ = "end of CoverTab[24963]"
}

// serverSum returns the contents of the verify_data member of a server's
//line /usr/local/go/src/crypto/tls/prf.go:208
// Finished message.
//line /usr/local/go/src/crypto/tls/prf.go:210
func (h finishedHash) serverSum(masterSecret []byte) []byte {
//line /usr/local/go/src/crypto/tls/prf.go:210
	_go_fuzz_dep_.CoverTab[24964]++
						out := make([]byte, finishedVerifyLength)
						h.prf(out, masterSecret, serverFinishedLabel, h.Sum())
						return out
//line /usr/local/go/src/crypto/tls/prf.go:213
	// _ = "end of CoverTab[24964]"
}

// hashForClientCertificate returns the handshake messages so far, pre-hashed if
//line /usr/local/go/src/crypto/tls/prf.go:216
// necessary, suitable for signing by a TLS client certificate.
//line /usr/local/go/src/crypto/tls/prf.go:218
func (h finishedHash) hashForClientCertificate(sigType uint8, hashAlg crypto.Hash) []byte {
//line /usr/local/go/src/crypto/tls/prf.go:218
	_go_fuzz_dep_.CoverTab[24965]++
						if (h.version >= VersionTLS12 || func() bool {
//line /usr/local/go/src/crypto/tls/prf.go:219
		_go_fuzz_dep_.CoverTab[24970]++
//line /usr/local/go/src/crypto/tls/prf.go:219
		return sigType == signatureEd25519
//line /usr/local/go/src/crypto/tls/prf.go:219
		// _ = "end of CoverTab[24970]"
//line /usr/local/go/src/crypto/tls/prf.go:219
	}()) && func() bool {
//line /usr/local/go/src/crypto/tls/prf.go:219
		_go_fuzz_dep_.CoverTab[24971]++
//line /usr/local/go/src/crypto/tls/prf.go:219
		return h.buffer == nil
//line /usr/local/go/src/crypto/tls/prf.go:219
		// _ = "end of CoverTab[24971]"
//line /usr/local/go/src/crypto/tls/prf.go:219
	}() {
//line /usr/local/go/src/crypto/tls/prf.go:219
		_go_fuzz_dep_.CoverTab[24972]++
							panic("tls: handshake hash for a client certificate requested after discarding the handshake buffer")
//line /usr/local/go/src/crypto/tls/prf.go:220
		// _ = "end of CoverTab[24972]"
	} else {
//line /usr/local/go/src/crypto/tls/prf.go:221
		_go_fuzz_dep_.CoverTab[24973]++
//line /usr/local/go/src/crypto/tls/prf.go:221
		// _ = "end of CoverTab[24973]"
//line /usr/local/go/src/crypto/tls/prf.go:221
	}
//line /usr/local/go/src/crypto/tls/prf.go:221
	// _ = "end of CoverTab[24965]"
//line /usr/local/go/src/crypto/tls/prf.go:221
	_go_fuzz_dep_.CoverTab[24966]++

						if sigType == signatureEd25519 {
//line /usr/local/go/src/crypto/tls/prf.go:223
		_go_fuzz_dep_.CoverTab[24974]++
							return h.buffer
//line /usr/local/go/src/crypto/tls/prf.go:224
		// _ = "end of CoverTab[24974]"
	} else {
//line /usr/local/go/src/crypto/tls/prf.go:225
		_go_fuzz_dep_.CoverTab[24975]++
//line /usr/local/go/src/crypto/tls/prf.go:225
		// _ = "end of CoverTab[24975]"
//line /usr/local/go/src/crypto/tls/prf.go:225
	}
//line /usr/local/go/src/crypto/tls/prf.go:225
	// _ = "end of CoverTab[24966]"
//line /usr/local/go/src/crypto/tls/prf.go:225
	_go_fuzz_dep_.CoverTab[24967]++

						if h.version >= VersionTLS12 {
//line /usr/local/go/src/crypto/tls/prf.go:227
		_go_fuzz_dep_.CoverTab[24976]++
							hash := hashAlg.New()
							hash.Write(h.buffer)
							return hash.Sum(nil)
//line /usr/local/go/src/crypto/tls/prf.go:230
		// _ = "end of CoverTab[24976]"
	} else {
//line /usr/local/go/src/crypto/tls/prf.go:231
		_go_fuzz_dep_.CoverTab[24977]++
//line /usr/local/go/src/crypto/tls/prf.go:231
		// _ = "end of CoverTab[24977]"
//line /usr/local/go/src/crypto/tls/prf.go:231
	}
//line /usr/local/go/src/crypto/tls/prf.go:231
	// _ = "end of CoverTab[24967]"
//line /usr/local/go/src/crypto/tls/prf.go:231
	_go_fuzz_dep_.CoverTab[24968]++

						if sigType == signatureECDSA {
//line /usr/local/go/src/crypto/tls/prf.go:233
		_go_fuzz_dep_.CoverTab[24978]++
							return h.server.Sum(nil)
//line /usr/local/go/src/crypto/tls/prf.go:234
		// _ = "end of CoverTab[24978]"
	} else {
//line /usr/local/go/src/crypto/tls/prf.go:235
		_go_fuzz_dep_.CoverTab[24979]++
//line /usr/local/go/src/crypto/tls/prf.go:235
		// _ = "end of CoverTab[24979]"
//line /usr/local/go/src/crypto/tls/prf.go:235
	}
//line /usr/local/go/src/crypto/tls/prf.go:235
	// _ = "end of CoverTab[24968]"
//line /usr/local/go/src/crypto/tls/prf.go:235
	_go_fuzz_dep_.CoverTab[24969]++

						return h.Sum()
//line /usr/local/go/src/crypto/tls/prf.go:237
	// _ = "end of CoverTab[24969]"
}

// discardHandshakeBuffer is called when there is no more need to
//line /usr/local/go/src/crypto/tls/prf.go:240
// buffer the entirety of the handshake messages.
//line /usr/local/go/src/crypto/tls/prf.go:242
func (h *finishedHash) discardHandshakeBuffer() {
//line /usr/local/go/src/crypto/tls/prf.go:242
	_go_fuzz_dep_.CoverTab[24980]++
						h.buffer = nil
//line /usr/local/go/src/crypto/tls/prf.go:243
	// _ = "end of CoverTab[24980]"
}

// noExportedKeyingMaterial is used as a value of
//line /usr/local/go/src/crypto/tls/prf.go:246
// ConnectionState.ekm when renegotiation is enabled and thus
//line /usr/local/go/src/crypto/tls/prf.go:246
// we wish to fail all key-material export requests.
//line /usr/local/go/src/crypto/tls/prf.go:249
func noExportedKeyingMaterial(label string, context []byte, length int) ([]byte, error) {
//line /usr/local/go/src/crypto/tls/prf.go:249
	_go_fuzz_dep_.CoverTab[24981]++
						return nil, errors.New("crypto/tls: ExportKeyingMaterial is unavailable when renegotiation is enabled")
//line /usr/local/go/src/crypto/tls/prf.go:250
	// _ = "end of CoverTab[24981]"
}

// ekmFromMasterSecret generates exported keying material as defined in RFC 5705.
func ekmFromMasterSecret(version uint16, suite *cipherSuite, masterSecret, clientRandom, serverRandom []byte) func(string, []byte, int) ([]byte, error) {
//line /usr/local/go/src/crypto/tls/prf.go:254
	_go_fuzz_dep_.CoverTab[24982]++
						return func(label string, context []byte, length int) ([]byte, error) {
//line /usr/local/go/src/crypto/tls/prf.go:255
		_go_fuzz_dep_.CoverTab[24983]++
							switch label {
		case "client finished", "server finished", "master secret", "key expansion":
//line /usr/local/go/src/crypto/tls/prf.go:257
			_go_fuzz_dep_.CoverTab[24987]++

								return nil, fmt.Errorf("crypto/tls: reserved ExportKeyingMaterial label: %s", label)
//line /usr/local/go/src/crypto/tls/prf.go:259
			// _ = "end of CoverTab[24987]"
//line /usr/local/go/src/crypto/tls/prf.go:259
		default:
//line /usr/local/go/src/crypto/tls/prf.go:259
			_go_fuzz_dep_.CoverTab[24988]++
//line /usr/local/go/src/crypto/tls/prf.go:259
			// _ = "end of CoverTab[24988]"
		}
//line /usr/local/go/src/crypto/tls/prf.go:260
		// _ = "end of CoverTab[24983]"
//line /usr/local/go/src/crypto/tls/prf.go:260
		_go_fuzz_dep_.CoverTab[24984]++

							seedLen := len(serverRandom) + len(clientRandom)
							if context != nil {
//line /usr/local/go/src/crypto/tls/prf.go:263
			_go_fuzz_dep_.CoverTab[24989]++
								seedLen += 2 + len(context)
//line /usr/local/go/src/crypto/tls/prf.go:264
			// _ = "end of CoverTab[24989]"
		} else {
//line /usr/local/go/src/crypto/tls/prf.go:265
			_go_fuzz_dep_.CoverTab[24990]++
//line /usr/local/go/src/crypto/tls/prf.go:265
			// _ = "end of CoverTab[24990]"
//line /usr/local/go/src/crypto/tls/prf.go:265
		}
//line /usr/local/go/src/crypto/tls/prf.go:265
		// _ = "end of CoverTab[24984]"
//line /usr/local/go/src/crypto/tls/prf.go:265
		_go_fuzz_dep_.CoverTab[24985]++
							seed := make([]byte, 0, seedLen)

							seed = append(seed, clientRandom...)
							seed = append(seed, serverRandom...)

							if context != nil {
//line /usr/local/go/src/crypto/tls/prf.go:271
			_go_fuzz_dep_.CoverTab[24991]++
								if len(context) >= 1<<16 {
//line /usr/local/go/src/crypto/tls/prf.go:272
				_go_fuzz_dep_.CoverTab[24993]++
									return nil, fmt.Errorf("crypto/tls: ExportKeyingMaterial context too long")
//line /usr/local/go/src/crypto/tls/prf.go:273
				// _ = "end of CoverTab[24993]"
			} else {
//line /usr/local/go/src/crypto/tls/prf.go:274
				_go_fuzz_dep_.CoverTab[24994]++
//line /usr/local/go/src/crypto/tls/prf.go:274
				// _ = "end of CoverTab[24994]"
//line /usr/local/go/src/crypto/tls/prf.go:274
			}
//line /usr/local/go/src/crypto/tls/prf.go:274
			// _ = "end of CoverTab[24991]"
//line /usr/local/go/src/crypto/tls/prf.go:274
			_go_fuzz_dep_.CoverTab[24992]++
								seed = append(seed, byte(len(context)>>8), byte(len(context)))
								seed = append(seed, context...)
//line /usr/local/go/src/crypto/tls/prf.go:276
			// _ = "end of CoverTab[24992]"
		} else {
//line /usr/local/go/src/crypto/tls/prf.go:277
			_go_fuzz_dep_.CoverTab[24995]++
//line /usr/local/go/src/crypto/tls/prf.go:277
			// _ = "end of CoverTab[24995]"
//line /usr/local/go/src/crypto/tls/prf.go:277
		}
//line /usr/local/go/src/crypto/tls/prf.go:277
		// _ = "end of CoverTab[24985]"
//line /usr/local/go/src/crypto/tls/prf.go:277
		_go_fuzz_dep_.CoverTab[24986]++

							keyMaterial := make([]byte, length)
							prfForVersion(version, suite)(keyMaterial, masterSecret, []byte(label), seed)
							return keyMaterial, nil
//line /usr/local/go/src/crypto/tls/prf.go:281
		// _ = "end of CoverTab[24986]"
	}
//line /usr/local/go/src/crypto/tls/prf.go:282
	// _ = "end of CoverTab[24982]"
}

//line /usr/local/go/src/crypto/tls/prf.go:283
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/prf.go:283
var _ = _go_fuzz_dep_.CoverTab
