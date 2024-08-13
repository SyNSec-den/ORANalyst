// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/tls/key_agreement.go:5
package tls

//line /usr/local/go/src/crypto/tls/key_agreement.go:5
import (
//line /usr/local/go/src/crypto/tls/key_agreement.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/key_agreement.go:5
)
//line /usr/local/go/src/crypto/tls/key_agreement.go:5
import (
//line /usr/local/go/src/crypto/tls/key_agreement.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/key_agreement.go:5
)

import (
	"crypto"
	"crypto/ecdh"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
)

// a keyAgreement implements the client and server side of a TLS key agreement
//line /usr/local/go/src/crypto/tls/key_agreement.go:19
// protocol by generating and processing key exchange messages.
//line /usr/local/go/src/crypto/tls/key_agreement.go:21
type keyAgreement interface {

//line /usr/local/go/src/crypto/tls/key_agreement.go:24
	// In the case that the key agreement protocol doesn't use a
	// ServerKeyExchange message, generateServerKeyExchange can return nil,
	// nil.
	generateServerKeyExchange(*Config, *Certificate, *clientHelloMsg, *serverHelloMsg) (*serverKeyExchangeMsg, error)
	processClientKeyExchange(*Config, *Certificate, *clientKeyExchangeMsg, uint16) ([]byte, error)

//line /usr/local/go/src/crypto/tls/key_agreement.go:32
	// This method may not be called if the server doesn't send a
	// ServerKeyExchange message.
	processServerKeyExchange(*Config, *clientHelloMsg, *serverHelloMsg, *x509.Certificate, *serverKeyExchangeMsg) error
	generateClientKeyExchange(*Config, *clientHelloMsg, *x509.Certificate) ([]byte, *clientKeyExchangeMsg, error)
}

var errClientKeyExchange = errors.New("tls: invalid ClientKeyExchange message")
var errServerKeyExchange = errors.New("tls: invalid ServerKeyExchange message")

// rsaKeyAgreement implements the standard TLS key agreement where the client
//line /usr/local/go/src/crypto/tls/key_agreement.go:41
// encrypts the pre-master secret to the server's public key.
//line /usr/local/go/src/crypto/tls/key_agreement.go:43
type rsaKeyAgreement struct{}

func (ka rsaKeyAgreement) generateServerKeyExchange(config *Config, cert *Certificate, clientHello *clientHelloMsg, hello *serverHelloMsg) (*serverKeyExchangeMsg, error) {
//line /usr/local/go/src/crypto/tls/key_agreement.go:45
	_go_fuzz_dep_.CoverTab[24730]++
								return nil, nil
//line /usr/local/go/src/crypto/tls/key_agreement.go:46
	// _ = "end of CoverTab[24730]"
}

func (ka rsaKeyAgreement) processClientKeyExchange(config *Config, cert *Certificate, ckx *clientKeyExchangeMsg, version uint16) ([]byte, error) {
//line /usr/local/go/src/crypto/tls/key_agreement.go:49
	_go_fuzz_dep_.CoverTab[24731]++
								if len(ckx.ciphertext) < 2 {
//line /usr/local/go/src/crypto/tls/key_agreement.go:50
		_go_fuzz_dep_.CoverTab[24736]++
									return nil, errClientKeyExchange
//line /usr/local/go/src/crypto/tls/key_agreement.go:51
		// _ = "end of CoverTab[24736]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:52
		_go_fuzz_dep_.CoverTab[24737]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:52
		// _ = "end of CoverTab[24737]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:52
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:52
	// _ = "end of CoverTab[24731]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:52
	_go_fuzz_dep_.CoverTab[24732]++
								ciphertextLen := int(ckx.ciphertext[0])<<8 | int(ckx.ciphertext[1])
								if ciphertextLen != len(ckx.ciphertext)-2 {
//line /usr/local/go/src/crypto/tls/key_agreement.go:54
		_go_fuzz_dep_.CoverTab[24738]++
									return nil, errClientKeyExchange
//line /usr/local/go/src/crypto/tls/key_agreement.go:55
		// _ = "end of CoverTab[24738]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:56
		_go_fuzz_dep_.CoverTab[24739]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:56
		// _ = "end of CoverTab[24739]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:56
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:56
	// _ = "end of CoverTab[24732]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:56
	_go_fuzz_dep_.CoverTab[24733]++
								ciphertext := ckx.ciphertext[2:]

								priv, ok := cert.PrivateKey.(crypto.Decrypter)
								if !ok {
//line /usr/local/go/src/crypto/tls/key_agreement.go:60
		_go_fuzz_dep_.CoverTab[24740]++
									return nil, errors.New("tls: certificate private key does not implement crypto.Decrypter")
//line /usr/local/go/src/crypto/tls/key_agreement.go:61
		// _ = "end of CoverTab[24740]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:62
		_go_fuzz_dep_.CoverTab[24741]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:62
		// _ = "end of CoverTab[24741]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:62
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:62
	// _ = "end of CoverTab[24733]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:62
	_go_fuzz_dep_.CoverTab[24734]++

								preMasterSecret, err := priv.Decrypt(config.rand(), ciphertext, &rsa.PKCS1v15DecryptOptions{SessionKeyLen: 48})
								if err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:65
		_go_fuzz_dep_.CoverTab[24742]++
									return nil, err
//line /usr/local/go/src/crypto/tls/key_agreement.go:66
		// _ = "end of CoverTab[24742]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:67
		_go_fuzz_dep_.CoverTab[24743]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:67
		// _ = "end of CoverTab[24743]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:67
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:67
	// _ = "end of CoverTab[24734]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:67
	_go_fuzz_dep_.CoverTab[24735]++

//line /usr/local/go/src/crypto/tls/key_agreement.go:74
	return preMasterSecret, nil
//line /usr/local/go/src/crypto/tls/key_agreement.go:74
	// _ = "end of CoverTab[24735]"
}

func (ka rsaKeyAgreement) processServerKeyExchange(config *Config, clientHello *clientHelloMsg, serverHello *serverHelloMsg, cert *x509.Certificate, skx *serverKeyExchangeMsg) error {
//line /usr/local/go/src/crypto/tls/key_agreement.go:77
	_go_fuzz_dep_.CoverTab[24744]++
								return errors.New("tls: unexpected ServerKeyExchange")
//line /usr/local/go/src/crypto/tls/key_agreement.go:78
	// _ = "end of CoverTab[24744]"
}

func (ka rsaKeyAgreement) generateClientKeyExchange(config *Config, clientHello *clientHelloMsg, cert *x509.Certificate) ([]byte, *clientKeyExchangeMsg, error) {
//line /usr/local/go/src/crypto/tls/key_agreement.go:81
	_go_fuzz_dep_.CoverTab[24745]++
								preMasterSecret := make([]byte, 48)
								preMasterSecret[0] = byte(clientHello.vers >> 8)
								preMasterSecret[1] = byte(clientHello.vers)
								_, err := io.ReadFull(config.rand(), preMasterSecret[2:])
								if err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:86
		_go_fuzz_dep_.CoverTab[24749]++
									return nil, nil, err
//line /usr/local/go/src/crypto/tls/key_agreement.go:87
		// _ = "end of CoverTab[24749]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:88
		_go_fuzz_dep_.CoverTab[24750]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:88
		// _ = "end of CoverTab[24750]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:88
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:88
	// _ = "end of CoverTab[24745]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:88
	_go_fuzz_dep_.CoverTab[24746]++

								rsaKey, ok := cert.PublicKey.(*rsa.PublicKey)
								if !ok {
//line /usr/local/go/src/crypto/tls/key_agreement.go:91
		_go_fuzz_dep_.CoverTab[24751]++
									return nil, nil, errors.New("tls: server certificate contains incorrect key type for selected ciphersuite")
//line /usr/local/go/src/crypto/tls/key_agreement.go:92
		// _ = "end of CoverTab[24751]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:93
		_go_fuzz_dep_.CoverTab[24752]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:93
		// _ = "end of CoverTab[24752]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:93
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:93
	// _ = "end of CoverTab[24746]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:93
	_go_fuzz_dep_.CoverTab[24747]++
								encrypted, err := rsa.EncryptPKCS1v15(config.rand(), rsaKey, preMasterSecret)
								if err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:95
		_go_fuzz_dep_.CoverTab[24753]++
									return nil, nil, err
//line /usr/local/go/src/crypto/tls/key_agreement.go:96
		// _ = "end of CoverTab[24753]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:97
		_go_fuzz_dep_.CoverTab[24754]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:97
		// _ = "end of CoverTab[24754]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:97
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:97
	// _ = "end of CoverTab[24747]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:97
	_go_fuzz_dep_.CoverTab[24748]++
								ckx := new(clientKeyExchangeMsg)
								ckx.ciphertext = make([]byte, len(encrypted)+2)
								ckx.ciphertext[0] = byte(len(encrypted) >> 8)
								ckx.ciphertext[1] = byte(len(encrypted))
								copy(ckx.ciphertext[2:], encrypted)
								return preMasterSecret, ckx, nil
//line /usr/local/go/src/crypto/tls/key_agreement.go:103
	// _ = "end of CoverTab[24748]"
}

// sha1Hash calculates a SHA1 hash over the given byte slices.
func sha1Hash(slices [][]byte) []byte {
//line /usr/local/go/src/crypto/tls/key_agreement.go:107
	_go_fuzz_dep_.CoverTab[24755]++
								hsha1 := sha1.New()
								for _, slice := range slices {
//line /usr/local/go/src/crypto/tls/key_agreement.go:109
		_go_fuzz_dep_.CoverTab[24757]++
									hsha1.Write(slice)
//line /usr/local/go/src/crypto/tls/key_agreement.go:110
		// _ = "end of CoverTab[24757]"
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:111
	// _ = "end of CoverTab[24755]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:111
	_go_fuzz_dep_.CoverTab[24756]++
								return hsha1.Sum(nil)
//line /usr/local/go/src/crypto/tls/key_agreement.go:112
	// _ = "end of CoverTab[24756]"
}

// md5SHA1Hash implements TLS 1.0's hybrid hash function which consists of the
//line /usr/local/go/src/crypto/tls/key_agreement.go:115
// concatenation of an MD5 and SHA1 hash.
//line /usr/local/go/src/crypto/tls/key_agreement.go:117
func md5SHA1Hash(slices [][]byte) []byte {
//line /usr/local/go/src/crypto/tls/key_agreement.go:117
	_go_fuzz_dep_.CoverTab[24758]++
								md5sha1 := make([]byte, md5.Size+sha1.Size)
								hmd5 := md5.New()
								for _, slice := range slices {
//line /usr/local/go/src/crypto/tls/key_agreement.go:120
		_go_fuzz_dep_.CoverTab[24760]++
									hmd5.Write(slice)
//line /usr/local/go/src/crypto/tls/key_agreement.go:121
		// _ = "end of CoverTab[24760]"
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:122
	// _ = "end of CoverTab[24758]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:122
	_go_fuzz_dep_.CoverTab[24759]++
								copy(md5sha1, hmd5.Sum(nil))
								copy(md5sha1[md5.Size:], sha1Hash(slices))
								return md5sha1
//line /usr/local/go/src/crypto/tls/key_agreement.go:125
	// _ = "end of CoverTab[24759]"
}

// hashForServerKeyExchange hashes the given slices and returns their digest
//line /usr/local/go/src/crypto/tls/key_agreement.go:128
// using the given hash function (for >= TLS 1.2) or using a default based on
//line /usr/local/go/src/crypto/tls/key_agreement.go:128
// the sigType (for earlier TLS versions). For Ed25519 signatures, which don't
//line /usr/local/go/src/crypto/tls/key_agreement.go:128
// do pre-hashing, it returns the concatenation of the slices.
//line /usr/local/go/src/crypto/tls/key_agreement.go:132
func hashForServerKeyExchange(sigType uint8, hashFunc crypto.Hash, version uint16, slices ...[]byte) []byte {
//line /usr/local/go/src/crypto/tls/key_agreement.go:132
	_go_fuzz_dep_.CoverTab[24761]++
								if sigType == signatureEd25519 {
//line /usr/local/go/src/crypto/tls/key_agreement.go:133
		_go_fuzz_dep_.CoverTab[24765]++
									var signed []byte
									for _, slice := range slices {
//line /usr/local/go/src/crypto/tls/key_agreement.go:135
			_go_fuzz_dep_.CoverTab[24767]++
										signed = append(signed, slice...)
//line /usr/local/go/src/crypto/tls/key_agreement.go:136
			// _ = "end of CoverTab[24767]"
		}
//line /usr/local/go/src/crypto/tls/key_agreement.go:137
		// _ = "end of CoverTab[24765]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:137
		_go_fuzz_dep_.CoverTab[24766]++
									return signed
//line /usr/local/go/src/crypto/tls/key_agreement.go:138
		// _ = "end of CoverTab[24766]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:139
		_go_fuzz_dep_.CoverTab[24768]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:139
		// _ = "end of CoverTab[24768]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:139
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:139
	// _ = "end of CoverTab[24761]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:139
	_go_fuzz_dep_.CoverTab[24762]++
								if version >= VersionTLS12 {
//line /usr/local/go/src/crypto/tls/key_agreement.go:140
		_go_fuzz_dep_.CoverTab[24769]++
									h := hashFunc.New()
									for _, slice := range slices {
//line /usr/local/go/src/crypto/tls/key_agreement.go:142
			_go_fuzz_dep_.CoverTab[24771]++
										h.Write(slice)
//line /usr/local/go/src/crypto/tls/key_agreement.go:143
			// _ = "end of CoverTab[24771]"
		}
//line /usr/local/go/src/crypto/tls/key_agreement.go:144
		// _ = "end of CoverTab[24769]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:144
		_go_fuzz_dep_.CoverTab[24770]++
									digest := h.Sum(nil)
									return digest
//line /usr/local/go/src/crypto/tls/key_agreement.go:146
		// _ = "end of CoverTab[24770]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:147
		_go_fuzz_dep_.CoverTab[24772]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:147
		// _ = "end of CoverTab[24772]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:147
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:147
	// _ = "end of CoverTab[24762]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:147
	_go_fuzz_dep_.CoverTab[24763]++
								if sigType == signatureECDSA {
//line /usr/local/go/src/crypto/tls/key_agreement.go:148
		_go_fuzz_dep_.CoverTab[24773]++
									return sha1Hash(slices)
//line /usr/local/go/src/crypto/tls/key_agreement.go:149
		// _ = "end of CoverTab[24773]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:150
		_go_fuzz_dep_.CoverTab[24774]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:150
		// _ = "end of CoverTab[24774]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:150
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:150
	// _ = "end of CoverTab[24763]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:150
	_go_fuzz_dep_.CoverTab[24764]++
								return md5SHA1Hash(slices)
//line /usr/local/go/src/crypto/tls/key_agreement.go:151
	// _ = "end of CoverTab[24764]"
}

// ecdheKeyAgreement implements a TLS key agreement where the server
//line /usr/local/go/src/crypto/tls/key_agreement.go:154
// generates an ephemeral EC public/private key pair and signs it. The
//line /usr/local/go/src/crypto/tls/key_agreement.go:154
// pre-master secret is then calculated using ECDH. The signature may
//line /usr/local/go/src/crypto/tls/key_agreement.go:154
// be ECDSA, Ed25519 or RSA.
//line /usr/local/go/src/crypto/tls/key_agreement.go:158
type ecdheKeyAgreement struct {
	version	uint16
	isRSA	bool
	key	*ecdh.PrivateKey

	// ckx and preMasterSecret are generated in processServerKeyExchange
	// and returned in generateClientKeyExchange.
	ckx		*clientKeyExchangeMsg
	preMasterSecret	[]byte
}

func (ka *ecdheKeyAgreement) generateServerKeyExchange(config *Config, cert *Certificate, clientHello *clientHelloMsg, hello *serverHelloMsg) (*serverKeyExchangeMsg, error) {
//line /usr/local/go/src/crypto/tls/key_agreement.go:169
	_go_fuzz_dep_.CoverTab[24775]++
								var curveID CurveID
								for _, c := range clientHello.supportedCurves {
//line /usr/local/go/src/crypto/tls/key_agreement.go:171
		_go_fuzz_dep_.CoverTab[24787]++
									if config.supportsCurve(c) {
//line /usr/local/go/src/crypto/tls/key_agreement.go:172
			_go_fuzz_dep_.CoverTab[24788]++
										curveID = c
										break
//line /usr/local/go/src/crypto/tls/key_agreement.go:174
			// _ = "end of CoverTab[24788]"
		} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:175
			_go_fuzz_dep_.CoverTab[24789]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:175
			// _ = "end of CoverTab[24789]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:175
		}
//line /usr/local/go/src/crypto/tls/key_agreement.go:175
		// _ = "end of CoverTab[24787]"
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:176
	// _ = "end of CoverTab[24775]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:176
	_go_fuzz_dep_.CoverTab[24776]++

								if curveID == 0 {
//line /usr/local/go/src/crypto/tls/key_agreement.go:178
		_go_fuzz_dep_.CoverTab[24790]++
									return nil, errors.New("tls: no supported elliptic curves offered")
//line /usr/local/go/src/crypto/tls/key_agreement.go:179
		// _ = "end of CoverTab[24790]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:180
		_go_fuzz_dep_.CoverTab[24791]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:180
		// _ = "end of CoverTab[24791]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:180
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:180
	// _ = "end of CoverTab[24776]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:180
	_go_fuzz_dep_.CoverTab[24777]++
								if _, ok := curveForCurveID(curveID); !ok {
//line /usr/local/go/src/crypto/tls/key_agreement.go:181
		_go_fuzz_dep_.CoverTab[24792]++
									return nil, errors.New("tls: CurvePreferences includes unsupported curve")
//line /usr/local/go/src/crypto/tls/key_agreement.go:182
		// _ = "end of CoverTab[24792]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:183
		_go_fuzz_dep_.CoverTab[24793]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:183
		// _ = "end of CoverTab[24793]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:183
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:183
	// _ = "end of CoverTab[24777]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:183
	_go_fuzz_dep_.CoverTab[24778]++

								key, err := generateECDHEKey(config.rand(), curveID)
								if err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:186
		_go_fuzz_dep_.CoverTab[24794]++
									return nil, err
//line /usr/local/go/src/crypto/tls/key_agreement.go:187
		// _ = "end of CoverTab[24794]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:188
		_go_fuzz_dep_.CoverTab[24795]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:188
		// _ = "end of CoverTab[24795]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:188
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:188
	// _ = "end of CoverTab[24778]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:188
	_go_fuzz_dep_.CoverTab[24779]++
								ka.key = key

//line /usr/local/go/src/crypto/tls/key_agreement.go:192
	ecdhePublic := key.PublicKey().Bytes()
	serverECDHEParams := make([]byte, 1+2+1+len(ecdhePublic))
	serverECDHEParams[0] = 3
	serverECDHEParams[1] = byte(curveID >> 8)
	serverECDHEParams[2] = byte(curveID)
	serverECDHEParams[3] = byte(len(ecdhePublic))
	copy(serverECDHEParams[4:], ecdhePublic)

	priv, ok := cert.PrivateKey.(crypto.Signer)
	if !ok {
//line /usr/local/go/src/crypto/tls/key_agreement.go:201
		_go_fuzz_dep_.CoverTab[24796]++
									return nil, fmt.Errorf("tls: certificate private key of type %T does not implement crypto.Signer", cert.PrivateKey)
//line /usr/local/go/src/crypto/tls/key_agreement.go:202
		// _ = "end of CoverTab[24796]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:203
		_go_fuzz_dep_.CoverTab[24797]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:203
		// _ = "end of CoverTab[24797]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:203
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:203
	// _ = "end of CoverTab[24779]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:203
	_go_fuzz_dep_.CoverTab[24780]++

								var signatureAlgorithm SignatureScheme
								var sigType uint8
								var sigHash crypto.Hash
								if ka.version >= VersionTLS12 {
//line /usr/local/go/src/crypto/tls/key_agreement.go:208
		_go_fuzz_dep_.CoverTab[24798]++
									signatureAlgorithm, err = selectSignatureScheme(ka.version, cert, clientHello.supportedSignatureAlgorithms)
									if err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:210
			_go_fuzz_dep_.CoverTab[24800]++
										return nil, err
//line /usr/local/go/src/crypto/tls/key_agreement.go:211
			// _ = "end of CoverTab[24800]"
		} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:212
			_go_fuzz_dep_.CoverTab[24801]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:212
			// _ = "end of CoverTab[24801]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:212
		}
//line /usr/local/go/src/crypto/tls/key_agreement.go:212
		// _ = "end of CoverTab[24798]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:212
		_go_fuzz_dep_.CoverTab[24799]++
									sigType, sigHash, err = typeAndHashFromSignatureScheme(signatureAlgorithm)
									if err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:214
			_go_fuzz_dep_.CoverTab[24802]++
										return nil, err
//line /usr/local/go/src/crypto/tls/key_agreement.go:215
			// _ = "end of CoverTab[24802]"
		} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:216
			_go_fuzz_dep_.CoverTab[24803]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:216
			// _ = "end of CoverTab[24803]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:216
		}
//line /usr/local/go/src/crypto/tls/key_agreement.go:216
		// _ = "end of CoverTab[24799]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:217
		_go_fuzz_dep_.CoverTab[24804]++
									sigType, sigHash, err = legacyTypeAndHashFromPublicKey(priv.Public())
									if err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:219
			_go_fuzz_dep_.CoverTab[24805]++
										return nil, err
//line /usr/local/go/src/crypto/tls/key_agreement.go:220
			// _ = "end of CoverTab[24805]"
		} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:221
			_go_fuzz_dep_.CoverTab[24806]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:221
			// _ = "end of CoverTab[24806]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:221
		}
//line /usr/local/go/src/crypto/tls/key_agreement.go:221
		// _ = "end of CoverTab[24804]"
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:222
	// _ = "end of CoverTab[24780]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:222
	_go_fuzz_dep_.CoverTab[24781]++
								if (sigType == signaturePKCS1v15 || func() bool {
//line /usr/local/go/src/crypto/tls/key_agreement.go:223
		_go_fuzz_dep_.CoverTab[24807]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:223
		return sigType == signatureRSAPSS
//line /usr/local/go/src/crypto/tls/key_agreement.go:223
		// _ = "end of CoverTab[24807]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:223
	}()) != ka.isRSA {
//line /usr/local/go/src/crypto/tls/key_agreement.go:223
		_go_fuzz_dep_.CoverTab[24808]++
									return nil, errors.New("tls: certificate cannot be used with the selected cipher suite")
//line /usr/local/go/src/crypto/tls/key_agreement.go:224
		// _ = "end of CoverTab[24808]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:225
		_go_fuzz_dep_.CoverTab[24809]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:225
		// _ = "end of CoverTab[24809]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:225
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:225
	// _ = "end of CoverTab[24781]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:225
	_go_fuzz_dep_.CoverTab[24782]++

								signed := hashForServerKeyExchange(sigType, sigHash, ka.version, clientHello.random, hello.random, serverECDHEParams)

								signOpts := crypto.SignerOpts(sigHash)
								if sigType == signatureRSAPSS {
//line /usr/local/go/src/crypto/tls/key_agreement.go:230
		_go_fuzz_dep_.CoverTab[24810]++
									signOpts = &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash, Hash: sigHash}
//line /usr/local/go/src/crypto/tls/key_agreement.go:231
		// _ = "end of CoverTab[24810]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:232
		_go_fuzz_dep_.CoverTab[24811]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:232
		// _ = "end of CoverTab[24811]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:232
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:232
	// _ = "end of CoverTab[24782]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:232
	_go_fuzz_dep_.CoverTab[24783]++
								sig, err := priv.Sign(config.rand(), signed, signOpts)
								if err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:234
		_go_fuzz_dep_.CoverTab[24812]++
									return nil, errors.New("tls: failed to sign ECDHE parameters: " + err.Error())
//line /usr/local/go/src/crypto/tls/key_agreement.go:235
		// _ = "end of CoverTab[24812]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:236
		_go_fuzz_dep_.CoverTab[24813]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:236
		// _ = "end of CoverTab[24813]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:236
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:236
	// _ = "end of CoverTab[24783]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:236
	_go_fuzz_dep_.CoverTab[24784]++

								skx := new(serverKeyExchangeMsg)
								sigAndHashLen := 0
								if ka.version >= VersionTLS12 {
//line /usr/local/go/src/crypto/tls/key_agreement.go:240
		_go_fuzz_dep_.CoverTab[24814]++
									sigAndHashLen = 2
//line /usr/local/go/src/crypto/tls/key_agreement.go:241
		// _ = "end of CoverTab[24814]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:242
		_go_fuzz_dep_.CoverTab[24815]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:242
		// _ = "end of CoverTab[24815]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:242
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:242
	// _ = "end of CoverTab[24784]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:242
	_go_fuzz_dep_.CoverTab[24785]++
								skx.key = make([]byte, len(serverECDHEParams)+sigAndHashLen+2+len(sig))
								copy(skx.key, serverECDHEParams)
								k := skx.key[len(serverECDHEParams):]
								if ka.version >= VersionTLS12 {
//line /usr/local/go/src/crypto/tls/key_agreement.go:246
		_go_fuzz_dep_.CoverTab[24816]++
									k[0] = byte(signatureAlgorithm >> 8)
									k[1] = byte(signatureAlgorithm)
									k = k[2:]
//line /usr/local/go/src/crypto/tls/key_agreement.go:249
		// _ = "end of CoverTab[24816]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:250
		_go_fuzz_dep_.CoverTab[24817]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:250
		// _ = "end of CoverTab[24817]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:250
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:250
	// _ = "end of CoverTab[24785]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:250
	_go_fuzz_dep_.CoverTab[24786]++
								k[0] = byte(len(sig) >> 8)
								k[1] = byte(len(sig))
								copy(k[2:], sig)

								return skx, nil
//line /usr/local/go/src/crypto/tls/key_agreement.go:255
	// _ = "end of CoverTab[24786]"
}

func (ka *ecdheKeyAgreement) processClientKeyExchange(config *Config, cert *Certificate, ckx *clientKeyExchangeMsg, version uint16) ([]byte, error) {
//line /usr/local/go/src/crypto/tls/key_agreement.go:258
	_go_fuzz_dep_.CoverTab[24818]++
								if len(ckx.ciphertext) == 0 || func() bool {
//line /usr/local/go/src/crypto/tls/key_agreement.go:259
		_go_fuzz_dep_.CoverTab[24822]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:259
		return int(ckx.ciphertext[0]) != len(ckx.ciphertext)-1
//line /usr/local/go/src/crypto/tls/key_agreement.go:259
		// _ = "end of CoverTab[24822]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:259
	}() {
//line /usr/local/go/src/crypto/tls/key_agreement.go:259
		_go_fuzz_dep_.CoverTab[24823]++
									return nil, errClientKeyExchange
//line /usr/local/go/src/crypto/tls/key_agreement.go:260
		// _ = "end of CoverTab[24823]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:261
		_go_fuzz_dep_.CoverTab[24824]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:261
		// _ = "end of CoverTab[24824]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:261
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:261
	// _ = "end of CoverTab[24818]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:261
	_go_fuzz_dep_.CoverTab[24819]++

								peerKey, err := ka.key.Curve().NewPublicKey(ckx.ciphertext[1:])
								if err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:264
		_go_fuzz_dep_.CoverTab[24825]++
									return nil, errClientKeyExchange
//line /usr/local/go/src/crypto/tls/key_agreement.go:265
		// _ = "end of CoverTab[24825]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:266
		_go_fuzz_dep_.CoverTab[24826]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:266
		// _ = "end of CoverTab[24826]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:266
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:266
	// _ = "end of CoverTab[24819]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:266
	_go_fuzz_dep_.CoverTab[24820]++
								preMasterSecret, err := ka.key.ECDH(peerKey)
								if err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:268
		_go_fuzz_dep_.CoverTab[24827]++
									return nil, errClientKeyExchange
//line /usr/local/go/src/crypto/tls/key_agreement.go:269
		// _ = "end of CoverTab[24827]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:270
		_go_fuzz_dep_.CoverTab[24828]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:270
		// _ = "end of CoverTab[24828]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:270
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:270
	// _ = "end of CoverTab[24820]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:270
	_go_fuzz_dep_.CoverTab[24821]++

								return preMasterSecret, nil
//line /usr/local/go/src/crypto/tls/key_agreement.go:272
	// _ = "end of CoverTab[24821]"
}

func (ka *ecdheKeyAgreement) processServerKeyExchange(config *Config, clientHello *clientHelloMsg, serverHello *serverHelloMsg, cert *x509.Certificate, skx *serverKeyExchangeMsg) error {
//line /usr/local/go/src/crypto/tls/key_agreement.go:275
	_go_fuzz_dep_.CoverTab[24829]++
								if len(skx.key) < 4 {
//line /usr/local/go/src/crypto/tls/key_agreement.go:276
		_go_fuzz_dep_.CoverTab[24842]++
									return errServerKeyExchange
//line /usr/local/go/src/crypto/tls/key_agreement.go:277
		// _ = "end of CoverTab[24842]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:278
		_go_fuzz_dep_.CoverTab[24843]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:278
		// _ = "end of CoverTab[24843]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:278
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:278
	// _ = "end of CoverTab[24829]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:278
	_go_fuzz_dep_.CoverTab[24830]++
								if skx.key[0] != 3 {
//line /usr/local/go/src/crypto/tls/key_agreement.go:279
		_go_fuzz_dep_.CoverTab[24844]++
									return errors.New("tls: server selected unsupported curve")
//line /usr/local/go/src/crypto/tls/key_agreement.go:280
		// _ = "end of CoverTab[24844]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:281
		_go_fuzz_dep_.CoverTab[24845]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:281
		// _ = "end of CoverTab[24845]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:281
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:281
	// _ = "end of CoverTab[24830]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:281
	_go_fuzz_dep_.CoverTab[24831]++
								curveID := CurveID(skx.key[1])<<8 | CurveID(skx.key[2])

								publicLen := int(skx.key[3])
								if publicLen+4 > len(skx.key) {
//line /usr/local/go/src/crypto/tls/key_agreement.go:285
		_go_fuzz_dep_.CoverTab[24846]++
									return errServerKeyExchange
//line /usr/local/go/src/crypto/tls/key_agreement.go:286
		// _ = "end of CoverTab[24846]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:287
		_go_fuzz_dep_.CoverTab[24847]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:287
		// _ = "end of CoverTab[24847]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:287
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:287
	// _ = "end of CoverTab[24831]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:287
	_go_fuzz_dep_.CoverTab[24832]++
								serverECDHEParams := skx.key[:4+publicLen]
								publicKey := serverECDHEParams[4:]

								sig := skx.key[4+publicLen:]
								if len(sig) < 2 {
//line /usr/local/go/src/crypto/tls/key_agreement.go:292
		_go_fuzz_dep_.CoverTab[24848]++
									return errServerKeyExchange
//line /usr/local/go/src/crypto/tls/key_agreement.go:293
		// _ = "end of CoverTab[24848]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:294
		_go_fuzz_dep_.CoverTab[24849]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:294
		// _ = "end of CoverTab[24849]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:294
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:294
	// _ = "end of CoverTab[24832]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:294
	_go_fuzz_dep_.CoverTab[24833]++

								if _, ok := curveForCurveID(curveID); !ok {
//line /usr/local/go/src/crypto/tls/key_agreement.go:296
		_go_fuzz_dep_.CoverTab[24850]++
									return errors.New("tls: server selected unsupported curve")
//line /usr/local/go/src/crypto/tls/key_agreement.go:297
		// _ = "end of CoverTab[24850]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:298
		_go_fuzz_dep_.CoverTab[24851]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:298
		// _ = "end of CoverTab[24851]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:298
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:298
	// _ = "end of CoverTab[24833]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:298
	_go_fuzz_dep_.CoverTab[24834]++

								key, err := generateECDHEKey(config.rand(), curveID)
								if err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:301
		_go_fuzz_dep_.CoverTab[24852]++
									return err
//line /usr/local/go/src/crypto/tls/key_agreement.go:302
		// _ = "end of CoverTab[24852]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:303
		_go_fuzz_dep_.CoverTab[24853]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:303
		// _ = "end of CoverTab[24853]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:303
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:303
	// _ = "end of CoverTab[24834]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:303
	_go_fuzz_dep_.CoverTab[24835]++
								ka.key = key

								peerKey, err := key.Curve().NewPublicKey(publicKey)
								if err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:307
		_go_fuzz_dep_.CoverTab[24854]++
									return errServerKeyExchange
//line /usr/local/go/src/crypto/tls/key_agreement.go:308
		// _ = "end of CoverTab[24854]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:309
		_go_fuzz_dep_.CoverTab[24855]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:309
		// _ = "end of CoverTab[24855]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:309
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:309
	// _ = "end of CoverTab[24835]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:309
	_go_fuzz_dep_.CoverTab[24836]++
								ka.preMasterSecret, err = key.ECDH(peerKey)
								if err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:311
		_go_fuzz_dep_.CoverTab[24856]++
									return errServerKeyExchange
//line /usr/local/go/src/crypto/tls/key_agreement.go:312
		// _ = "end of CoverTab[24856]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:313
		_go_fuzz_dep_.CoverTab[24857]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:313
		// _ = "end of CoverTab[24857]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:313
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:313
	// _ = "end of CoverTab[24836]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:313
	_go_fuzz_dep_.CoverTab[24837]++

								ourPublicKey := key.PublicKey().Bytes()
								ka.ckx = new(clientKeyExchangeMsg)
								ka.ckx.ciphertext = make([]byte, 1+len(ourPublicKey))
								ka.ckx.ciphertext[0] = byte(len(ourPublicKey))
								copy(ka.ckx.ciphertext[1:], ourPublicKey)

								var sigType uint8
								var sigHash crypto.Hash
								if ka.version >= VersionTLS12 {
//line /usr/local/go/src/crypto/tls/key_agreement.go:323
		_go_fuzz_dep_.CoverTab[24858]++
									signatureAlgorithm := SignatureScheme(sig[0])<<8 | SignatureScheme(sig[1])
									sig = sig[2:]
									if len(sig) < 2 {
//line /usr/local/go/src/crypto/tls/key_agreement.go:326
			_go_fuzz_dep_.CoverTab[24861]++
										return errServerKeyExchange
//line /usr/local/go/src/crypto/tls/key_agreement.go:327
			// _ = "end of CoverTab[24861]"
		} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:328
			_go_fuzz_dep_.CoverTab[24862]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:328
			// _ = "end of CoverTab[24862]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:328
		}
//line /usr/local/go/src/crypto/tls/key_agreement.go:328
		// _ = "end of CoverTab[24858]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:328
		_go_fuzz_dep_.CoverTab[24859]++

									if !isSupportedSignatureAlgorithm(signatureAlgorithm, clientHello.supportedSignatureAlgorithms) {
//line /usr/local/go/src/crypto/tls/key_agreement.go:330
			_go_fuzz_dep_.CoverTab[24863]++
										return errors.New("tls: certificate used with invalid signature algorithm")
//line /usr/local/go/src/crypto/tls/key_agreement.go:331
			// _ = "end of CoverTab[24863]"
		} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:332
			_go_fuzz_dep_.CoverTab[24864]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:332
			// _ = "end of CoverTab[24864]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:332
		}
//line /usr/local/go/src/crypto/tls/key_agreement.go:332
		// _ = "end of CoverTab[24859]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:332
		_go_fuzz_dep_.CoverTab[24860]++
									sigType, sigHash, err = typeAndHashFromSignatureScheme(signatureAlgorithm)
									if err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:334
			_go_fuzz_dep_.CoverTab[24865]++
										return err
//line /usr/local/go/src/crypto/tls/key_agreement.go:335
			// _ = "end of CoverTab[24865]"
		} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:336
			_go_fuzz_dep_.CoverTab[24866]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:336
			// _ = "end of CoverTab[24866]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:336
		}
//line /usr/local/go/src/crypto/tls/key_agreement.go:336
		// _ = "end of CoverTab[24860]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:337
		_go_fuzz_dep_.CoverTab[24867]++
									sigType, sigHash, err = legacyTypeAndHashFromPublicKey(cert.PublicKey)
									if err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:339
			_go_fuzz_dep_.CoverTab[24868]++
										return err
//line /usr/local/go/src/crypto/tls/key_agreement.go:340
			// _ = "end of CoverTab[24868]"
		} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:341
			_go_fuzz_dep_.CoverTab[24869]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:341
			// _ = "end of CoverTab[24869]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:341
		}
//line /usr/local/go/src/crypto/tls/key_agreement.go:341
		// _ = "end of CoverTab[24867]"
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:342
	// _ = "end of CoverTab[24837]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:342
	_go_fuzz_dep_.CoverTab[24838]++
								if (sigType == signaturePKCS1v15 || func() bool {
//line /usr/local/go/src/crypto/tls/key_agreement.go:343
		_go_fuzz_dep_.CoverTab[24870]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:343
		return sigType == signatureRSAPSS
//line /usr/local/go/src/crypto/tls/key_agreement.go:343
		// _ = "end of CoverTab[24870]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:343
	}()) != ka.isRSA {
//line /usr/local/go/src/crypto/tls/key_agreement.go:343
		_go_fuzz_dep_.CoverTab[24871]++
									return errServerKeyExchange
//line /usr/local/go/src/crypto/tls/key_agreement.go:344
		// _ = "end of CoverTab[24871]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:345
		_go_fuzz_dep_.CoverTab[24872]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:345
		// _ = "end of CoverTab[24872]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:345
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:345
	// _ = "end of CoverTab[24838]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:345
	_go_fuzz_dep_.CoverTab[24839]++

								sigLen := int(sig[0])<<8 | int(sig[1])
								if sigLen+2 != len(sig) {
//line /usr/local/go/src/crypto/tls/key_agreement.go:348
		_go_fuzz_dep_.CoverTab[24873]++
									return errServerKeyExchange
//line /usr/local/go/src/crypto/tls/key_agreement.go:349
		// _ = "end of CoverTab[24873]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:350
		_go_fuzz_dep_.CoverTab[24874]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:350
		// _ = "end of CoverTab[24874]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:350
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:350
	// _ = "end of CoverTab[24839]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:350
	_go_fuzz_dep_.CoverTab[24840]++
								sig = sig[2:]

								signed := hashForServerKeyExchange(sigType, sigHash, ka.version, clientHello.random, serverHello.random, serverECDHEParams)
								if err := verifyHandshakeSignature(sigType, cert.PublicKey, sigHash, signed, sig); err != nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:354
		_go_fuzz_dep_.CoverTab[24875]++
									return errors.New("tls: invalid signature by the server certificate: " + err.Error())
//line /usr/local/go/src/crypto/tls/key_agreement.go:355
		// _ = "end of CoverTab[24875]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:356
		_go_fuzz_dep_.CoverTab[24876]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:356
		// _ = "end of CoverTab[24876]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:356
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:356
	// _ = "end of CoverTab[24840]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:356
	_go_fuzz_dep_.CoverTab[24841]++
								return nil
//line /usr/local/go/src/crypto/tls/key_agreement.go:357
	// _ = "end of CoverTab[24841]"
}

func (ka *ecdheKeyAgreement) generateClientKeyExchange(config *Config, clientHello *clientHelloMsg, cert *x509.Certificate) ([]byte, *clientKeyExchangeMsg, error) {
//line /usr/local/go/src/crypto/tls/key_agreement.go:360
	_go_fuzz_dep_.CoverTab[24877]++
								if ka.ckx == nil {
//line /usr/local/go/src/crypto/tls/key_agreement.go:361
		_go_fuzz_dep_.CoverTab[24879]++
									return nil, nil, errors.New("tls: missing ServerKeyExchange message")
//line /usr/local/go/src/crypto/tls/key_agreement.go:362
		// _ = "end of CoverTab[24879]"
	} else {
//line /usr/local/go/src/crypto/tls/key_agreement.go:363
		_go_fuzz_dep_.CoverTab[24880]++
//line /usr/local/go/src/crypto/tls/key_agreement.go:363
		// _ = "end of CoverTab[24880]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:363
	}
//line /usr/local/go/src/crypto/tls/key_agreement.go:363
	// _ = "end of CoverTab[24877]"
//line /usr/local/go/src/crypto/tls/key_agreement.go:363
	_go_fuzz_dep_.CoverTab[24878]++

								return ka.preMasterSecret, ka.ckx, nil
//line /usr/local/go/src/crypto/tls/key_agreement.go:365
	// _ = "end of CoverTab[24878]"
}

//line /usr/local/go/src/crypto/tls/key_agreement.go:366
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/key_agreement.go:366
var _ = _go_fuzz_dep_.CoverTab
