// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/tls/auth.go:5
package tls

//line /usr/local/go/src/crypto/tls/auth.go:5
import (
//line /usr/local/go/src/crypto/tls/auth.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/auth.go:5
)
//line /usr/local/go/src/crypto/tls/auth.go:5
import (
//line /usr/local/go/src/crypto/tls/auth.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/auth.go:5
)

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rsa"
	"errors"
	"fmt"
	"hash"
	"io"
)

// verifyHandshakeSignature verifies a signature against pre-hashed
//line /usr/local/go/src/crypto/tls/auth.go:20
// (if required) handshake contents.
//line /usr/local/go/src/crypto/tls/auth.go:22
func verifyHandshakeSignature(sigType uint8, pubkey crypto.PublicKey, hashFunc crypto.Hash, signed, sig []byte) error {
//line /usr/local/go/src/crypto/tls/auth.go:22
	_go_fuzz_dep_.CoverTab[21011]++
						switch sigType {
	case signatureECDSA:
//line /usr/local/go/src/crypto/tls/auth.go:24
		_go_fuzz_dep_.CoverTab[21013]++
							pubKey, ok := pubkey.(*ecdsa.PublicKey)
							if !ok {
//line /usr/local/go/src/crypto/tls/auth.go:26
			_go_fuzz_dep_.CoverTab[21022]++
								return fmt.Errorf("expected an ECDSA public key, got %T", pubkey)
//line /usr/local/go/src/crypto/tls/auth.go:27
			// _ = "end of CoverTab[21022]"
		} else {
//line /usr/local/go/src/crypto/tls/auth.go:28
			_go_fuzz_dep_.CoverTab[21023]++
//line /usr/local/go/src/crypto/tls/auth.go:28
			// _ = "end of CoverTab[21023]"
//line /usr/local/go/src/crypto/tls/auth.go:28
		}
//line /usr/local/go/src/crypto/tls/auth.go:28
		// _ = "end of CoverTab[21013]"
//line /usr/local/go/src/crypto/tls/auth.go:28
		_go_fuzz_dep_.CoverTab[21014]++
							if !ecdsa.VerifyASN1(pubKey, signed, sig) {
//line /usr/local/go/src/crypto/tls/auth.go:29
			_go_fuzz_dep_.CoverTab[21024]++
								return errors.New("ECDSA verification failure")
//line /usr/local/go/src/crypto/tls/auth.go:30
			// _ = "end of CoverTab[21024]"
		} else {
//line /usr/local/go/src/crypto/tls/auth.go:31
			_go_fuzz_dep_.CoverTab[21025]++
//line /usr/local/go/src/crypto/tls/auth.go:31
			// _ = "end of CoverTab[21025]"
//line /usr/local/go/src/crypto/tls/auth.go:31
		}
//line /usr/local/go/src/crypto/tls/auth.go:31
		// _ = "end of CoverTab[21014]"
	case signatureEd25519:
//line /usr/local/go/src/crypto/tls/auth.go:32
		_go_fuzz_dep_.CoverTab[21015]++
							pubKey, ok := pubkey.(ed25519.PublicKey)
							if !ok {
//line /usr/local/go/src/crypto/tls/auth.go:34
			_go_fuzz_dep_.CoverTab[21026]++
								return fmt.Errorf("expected an Ed25519 public key, got %T", pubkey)
//line /usr/local/go/src/crypto/tls/auth.go:35
			// _ = "end of CoverTab[21026]"
		} else {
//line /usr/local/go/src/crypto/tls/auth.go:36
			_go_fuzz_dep_.CoverTab[21027]++
//line /usr/local/go/src/crypto/tls/auth.go:36
			// _ = "end of CoverTab[21027]"
//line /usr/local/go/src/crypto/tls/auth.go:36
		}
//line /usr/local/go/src/crypto/tls/auth.go:36
		// _ = "end of CoverTab[21015]"
//line /usr/local/go/src/crypto/tls/auth.go:36
		_go_fuzz_dep_.CoverTab[21016]++
							if !ed25519.Verify(pubKey, signed, sig) {
//line /usr/local/go/src/crypto/tls/auth.go:37
			_go_fuzz_dep_.CoverTab[21028]++
								return errors.New("Ed25519 verification failure")
//line /usr/local/go/src/crypto/tls/auth.go:38
			// _ = "end of CoverTab[21028]"
		} else {
//line /usr/local/go/src/crypto/tls/auth.go:39
			_go_fuzz_dep_.CoverTab[21029]++
//line /usr/local/go/src/crypto/tls/auth.go:39
			// _ = "end of CoverTab[21029]"
//line /usr/local/go/src/crypto/tls/auth.go:39
		}
//line /usr/local/go/src/crypto/tls/auth.go:39
		// _ = "end of CoverTab[21016]"
	case signaturePKCS1v15:
//line /usr/local/go/src/crypto/tls/auth.go:40
		_go_fuzz_dep_.CoverTab[21017]++
							pubKey, ok := pubkey.(*rsa.PublicKey)
							if !ok {
//line /usr/local/go/src/crypto/tls/auth.go:42
			_go_fuzz_dep_.CoverTab[21030]++
								return fmt.Errorf("expected an RSA public key, got %T", pubkey)
//line /usr/local/go/src/crypto/tls/auth.go:43
			// _ = "end of CoverTab[21030]"
		} else {
//line /usr/local/go/src/crypto/tls/auth.go:44
			_go_fuzz_dep_.CoverTab[21031]++
//line /usr/local/go/src/crypto/tls/auth.go:44
			// _ = "end of CoverTab[21031]"
//line /usr/local/go/src/crypto/tls/auth.go:44
		}
//line /usr/local/go/src/crypto/tls/auth.go:44
		// _ = "end of CoverTab[21017]"
//line /usr/local/go/src/crypto/tls/auth.go:44
		_go_fuzz_dep_.CoverTab[21018]++
							if err := rsa.VerifyPKCS1v15(pubKey, hashFunc, signed, sig); err != nil {
//line /usr/local/go/src/crypto/tls/auth.go:45
			_go_fuzz_dep_.CoverTab[21032]++
								return err
//line /usr/local/go/src/crypto/tls/auth.go:46
			// _ = "end of CoverTab[21032]"
		} else {
//line /usr/local/go/src/crypto/tls/auth.go:47
			_go_fuzz_dep_.CoverTab[21033]++
//line /usr/local/go/src/crypto/tls/auth.go:47
			// _ = "end of CoverTab[21033]"
//line /usr/local/go/src/crypto/tls/auth.go:47
		}
//line /usr/local/go/src/crypto/tls/auth.go:47
		// _ = "end of CoverTab[21018]"
	case signatureRSAPSS:
//line /usr/local/go/src/crypto/tls/auth.go:48
		_go_fuzz_dep_.CoverTab[21019]++
							pubKey, ok := pubkey.(*rsa.PublicKey)
							if !ok {
//line /usr/local/go/src/crypto/tls/auth.go:50
			_go_fuzz_dep_.CoverTab[21034]++
								return fmt.Errorf("expected an RSA public key, got %T", pubkey)
//line /usr/local/go/src/crypto/tls/auth.go:51
			// _ = "end of CoverTab[21034]"
		} else {
//line /usr/local/go/src/crypto/tls/auth.go:52
			_go_fuzz_dep_.CoverTab[21035]++
//line /usr/local/go/src/crypto/tls/auth.go:52
			// _ = "end of CoverTab[21035]"
//line /usr/local/go/src/crypto/tls/auth.go:52
		}
//line /usr/local/go/src/crypto/tls/auth.go:52
		// _ = "end of CoverTab[21019]"
//line /usr/local/go/src/crypto/tls/auth.go:52
		_go_fuzz_dep_.CoverTab[21020]++
							signOpts := &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash}
							if err := rsa.VerifyPSS(pubKey, hashFunc, signed, sig, signOpts); err != nil {
//line /usr/local/go/src/crypto/tls/auth.go:54
			_go_fuzz_dep_.CoverTab[21036]++
								return err
//line /usr/local/go/src/crypto/tls/auth.go:55
			// _ = "end of CoverTab[21036]"
		} else {
//line /usr/local/go/src/crypto/tls/auth.go:56
			_go_fuzz_dep_.CoverTab[21037]++
//line /usr/local/go/src/crypto/tls/auth.go:56
			// _ = "end of CoverTab[21037]"
//line /usr/local/go/src/crypto/tls/auth.go:56
		}
//line /usr/local/go/src/crypto/tls/auth.go:56
		// _ = "end of CoverTab[21020]"
	default:
//line /usr/local/go/src/crypto/tls/auth.go:57
		_go_fuzz_dep_.CoverTab[21021]++
							return errors.New("internal error: unknown signature type")
//line /usr/local/go/src/crypto/tls/auth.go:58
		// _ = "end of CoverTab[21021]"
	}
//line /usr/local/go/src/crypto/tls/auth.go:59
	// _ = "end of CoverTab[21011]"
//line /usr/local/go/src/crypto/tls/auth.go:59
	_go_fuzz_dep_.CoverTab[21012]++
						return nil
//line /usr/local/go/src/crypto/tls/auth.go:60
	// _ = "end of CoverTab[21012]"
}

const (
	serverSignatureContext	= "TLS 1.3, server CertificateVerify\x00"
	clientSignatureContext	= "TLS 1.3, client CertificateVerify\x00"
)

var signaturePadding = []byte{
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
}

// signedMessage returns the pre-hashed (if necessary) message to be signed by
//line /usr/local/go/src/crypto/tls/auth.go:79
// certificate keys in TLS 1.3. See RFC 8446, Section 4.4.3.
//line /usr/local/go/src/crypto/tls/auth.go:81
func signedMessage(sigHash crypto.Hash, context string, transcript hash.Hash) []byte {
//line /usr/local/go/src/crypto/tls/auth.go:81
	_go_fuzz_dep_.CoverTab[21038]++
						if sigHash == directSigning {
//line /usr/local/go/src/crypto/tls/auth.go:82
		_go_fuzz_dep_.CoverTab[21040]++
							b := &bytes.Buffer{}
							b.Write(signaturePadding)
							io.WriteString(b, context)
							b.Write(transcript.Sum(nil))
							return b.Bytes()
//line /usr/local/go/src/crypto/tls/auth.go:87
		// _ = "end of CoverTab[21040]"
	} else {
//line /usr/local/go/src/crypto/tls/auth.go:88
		_go_fuzz_dep_.CoverTab[21041]++
//line /usr/local/go/src/crypto/tls/auth.go:88
		// _ = "end of CoverTab[21041]"
//line /usr/local/go/src/crypto/tls/auth.go:88
	}
//line /usr/local/go/src/crypto/tls/auth.go:88
	// _ = "end of CoverTab[21038]"
//line /usr/local/go/src/crypto/tls/auth.go:88
	_go_fuzz_dep_.CoverTab[21039]++
						h := sigHash.New()
						h.Write(signaturePadding)
						io.WriteString(h, context)
						h.Write(transcript.Sum(nil))
						return h.Sum(nil)
//line /usr/local/go/src/crypto/tls/auth.go:93
	// _ = "end of CoverTab[21039]"
}

// typeAndHashFromSignatureScheme returns the corresponding signature type and
//line /usr/local/go/src/crypto/tls/auth.go:96
// crypto.Hash for a given TLS SignatureScheme.
//line /usr/local/go/src/crypto/tls/auth.go:98
func typeAndHashFromSignatureScheme(signatureAlgorithm SignatureScheme) (sigType uint8, hash crypto.Hash, err error) {
//line /usr/local/go/src/crypto/tls/auth.go:98
	_go_fuzz_dep_.CoverTab[21042]++
						switch signatureAlgorithm {
	case PKCS1WithSHA1, PKCS1WithSHA256, PKCS1WithSHA384, PKCS1WithSHA512:
//line /usr/local/go/src/crypto/tls/auth.go:100
		_go_fuzz_dep_.CoverTab[21045]++
								sigType = signaturePKCS1v15
//line /usr/local/go/src/crypto/tls/auth.go:101
		// _ = "end of CoverTab[21045]"
	case PSSWithSHA256, PSSWithSHA384, PSSWithSHA512:
//line /usr/local/go/src/crypto/tls/auth.go:102
		_go_fuzz_dep_.CoverTab[21046]++
								sigType = signatureRSAPSS
//line /usr/local/go/src/crypto/tls/auth.go:103
		// _ = "end of CoverTab[21046]"
	case ECDSAWithSHA1, ECDSAWithP256AndSHA256, ECDSAWithP384AndSHA384, ECDSAWithP521AndSHA512:
//line /usr/local/go/src/crypto/tls/auth.go:104
		_go_fuzz_dep_.CoverTab[21047]++
								sigType = signatureECDSA
//line /usr/local/go/src/crypto/tls/auth.go:105
		// _ = "end of CoverTab[21047]"
	case Ed25519:
//line /usr/local/go/src/crypto/tls/auth.go:106
		_go_fuzz_dep_.CoverTab[21048]++
								sigType = signatureEd25519
//line /usr/local/go/src/crypto/tls/auth.go:107
		// _ = "end of CoverTab[21048]"
	default:
//line /usr/local/go/src/crypto/tls/auth.go:108
		_go_fuzz_dep_.CoverTab[21049]++
								return 0, 0, fmt.Errorf("unsupported signature algorithm: %v", signatureAlgorithm)
//line /usr/local/go/src/crypto/tls/auth.go:109
		// _ = "end of CoverTab[21049]"
	}
//line /usr/local/go/src/crypto/tls/auth.go:110
	// _ = "end of CoverTab[21042]"
//line /usr/local/go/src/crypto/tls/auth.go:110
	_go_fuzz_dep_.CoverTab[21043]++
							switch signatureAlgorithm {
	case PKCS1WithSHA1, ECDSAWithSHA1:
//line /usr/local/go/src/crypto/tls/auth.go:112
		_go_fuzz_dep_.CoverTab[21050]++
								hash = crypto.SHA1
//line /usr/local/go/src/crypto/tls/auth.go:113
		// _ = "end of CoverTab[21050]"
	case PKCS1WithSHA256, PSSWithSHA256, ECDSAWithP256AndSHA256:
//line /usr/local/go/src/crypto/tls/auth.go:114
		_go_fuzz_dep_.CoverTab[21051]++
								hash = crypto.SHA256
//line /usr/local/go/src/crypto/tls/auth.go:115
		// _ = "end of CoverTab[21051]"
	case PKCS1WithSHA384, PSSWithSHA384, ECDSAWithP384AndSHA384:
//line /usr/local/go/src/crypto/tls/auth.go:116
		_go_fuzz_dep_.CoverTab[21052]++
								hash = crypto.SHA384
//line /usr/local/go/src/crypto/tls/auth.go:117
		// _ = "end of CoverTab[21052]"
	case PKCS1WithSHA512, PSSWithSHA512, ECDSAWithP521AndSHA512:
//line /usr/local/go/src/crypto/tls/auth.go:118
		_go_fuzz_dep_.CoverTab[21053]++
								hash = crypto.SHA512
//line /usr/local/go/src/crypto/tls/auth.go:119
		// _ = "end of CoverTab[21053]"
	case Ed25519:
//line /usr/local/go/src/crypto/tls/auth.go:120
		_go_fuzz_dep_.CoverTab[21054]++
								hash = directSigning
//line /usr/local/go/src/crypto/tls/auth.go:121
		// _ = "end of CoverTab[21054]"
	default:
//line /usr/local/go/src/crypto/tls/auth.go:122
		_go_fuzz_dep_.CoverTab[21055]++
								return 0, 0, fmt.Errorf("unsupported signature algorithm: %v", signatureAlgorithm)
//line /usr/local/go/src/crypto/tls/auth.go:123
		// _ = "end of CoverTab[21055]"
	}
//line /usr/local/go/src/crypto/tls/auth.go:124
	// _ = "end of CoverTab[21043]"
//line /usr/local/go/src/crypto/tls/auth.go:124
	_go_fuzz_dep_.CoverTab[21044]++
							return sigType, hash, nil
//line /usr/local/go/src/crypto/tls/auth.go:125
	// _ = "end of CoverTab[21044]"
}

// legacyTypeAndHashFromPublicKey returns the fixed signature type and crypto.Hash for
//line /usr/local/go/src/crypto/tls/auth.go:128
// a given public key used with TLS 1.0 and 1.1, before the introduction of
//line /usr/local/go/src/crypto/tls/auth.go:128
// signature algorithm negotiation.
//line /usr/local/go/src/crypto/tls/auth.go:131
func legacyTypeAndHashFromPublicKey(pub crypto.PublicKey) (sigType uint8, hash crypto.Hash, err error) {
//line /usr/local/go/src/crypto/tls/auth.go:131
	_go_fuzz_dep_.CoverTab[21056]++
							switch pub.(type) {
	case *rsa.PublicKey:
//line /usr/local/go/src/crypto/tls/auth.go:133
		_go_fuzz_dep_.CoverTab[21057]++
								return signaturePKCS1v15, crypto.MD5SHA1, nil
//line /usr/local/go/src/crypto/tls/auth.go:134
		// _ = "end of CoverTab[21057]"
	case *ecdsa.PublicKey:
//line /usr/local/go/src/crypto/tls/auth.go:135
		_go_fuzz_dep_.CoverTab[21058]++
								return signatureECDSA, crypto.SHA1, nil
//line /usr/local/go/src/crypto/tls/auth.go:136
		// _ = "end of CoverTab[21058]"
	case ed25519.PublicKey:
//line /usr/local/go/src/crypto/tls/auth.go:137
		_go_fuzz_dep_.CoverTab[21059]++

//line /usr/local/go/src/crypto/tls/auth.go:142
		return 0, 0, fmt.Errorf("tls: Ed25519 public keys are not supported before TLS 1.2")
//line /usr/local/go/src/crypto/tls/auth.go:142
		// _ = "end of CoverTab[21059]"
	default:
//line /usr/local/go/src/crypto/tls/auth.go:143
		_go_fuzz_dep_.CoverTab[21060]++
								return 0, 0, fmt.Errorf("tls: unsupported public key: %T", pub)
//line /usr/local/go/src/crypto/tls/auth.go:144
		// _ = "end of CoverTab[21060]"
	}
//line /usr/local/go/src/crypto/tls/auth.go:145
	// _ = "end of CoverTab[21056]"
}

var rsaSignatureSchemes = []struct {
	scheme		SignatureScheme
	minModulusBytes	int
	maxVersion	uint16
}{

//line /usr/local/go/src/crypto/tls/auth.go:155
	{PSSWithSHA256, crypto.SHA256.Size()*2 + 2, VersionTLS13},
							{PSSWithSHA384, crypto.SHA384.Size()*2 + 2, VersionTLS13},
							{PSSWithSHA512, crypto.SHA512.Size()*2 + 2, VersionTLS13},

//line /usr/local/go/src/crypto/tls/auth.go:161
	{PKCS1WithSHA256, 19 + crypto.SHA256.Size() + 11, VersionTLS12},
							{PKCS1WithSHA384, 19 + crypto.SHA384.Size() + 11, VersionTLS12},
							{PKCS1WithSHA512, 19 + crypto.SHA512.Size() + 11, VersionTLS12},
							{PKCS1WithSHA1, 15 + crypto.SHA1.Size() + 11, VersionTLS12},
}

// signatureSchemesForCertificate returns the list of supported SignatureSchemes
//line /usr/local/go/src/crypto/tls/auth.go:167
// for a given certificate, based on the public key and the protocol version,
//line /usr/local/go/src/crypto/tls/auth.go:167
// and optionally filtered by its explicit SupportedSignatureAlgorithms.
//line /usr/local/go/src/crypto/tls/auth.go:167
//
//line /usr/local/go/src/crypto/tls/auth.go:167
// This function must be kept in sync with supportedSignatureAlgorithms.
//line /usr/local/go/src/crypto/tls/auth.go:167
// FIPS filtering is applied in the caller, selectSignatureScheme.
//line /usr/local/go/src/crypto/tls/auth.go:173
func signatureSchemesForCertificate(version uint16, cert *Certificate) []SignatureScheme {
//line /usr/local/go/src/crypto/tls/auth.go:173
	_go_fuzz_dep_.CoverTab[21061]++
							priv, ok := cert.PrivateKey.(crypto.Signer)
							if !ok {
//line /usr/local/go/src/crypto/tls/auth.go:175
		_go_fuzz_dep_.CoverTab[21065]++
								return nil
//line /usr/local/go/src/crypto/tls/auth.go:176
		// _ = "end of CoverTab[21065]"
	} else {
//line /usr/local/go/src/crypto/tls/auth.go:177
		_go_fuzz_dep_.CoverTab[21066]++
//line /usr/local/go/src/crypto/tls/auth.go:177
		// _ = "end of CoverTab[21066]"
//line /usr/local/go/src/crypto/tls/auth.go:177
	}
//line /usr/local/go/src/crypto/tls/auth.go:177
	// _ = "end of CoverTab[21061]"
//line /usr/local/go/src/crypto/tls/auth.go:177
	_go_fuzz_dep_.CoverTab[21062]++

							var sigAlgs []SignatureScheme
							switch pub := priv.Public().(type) {
	case *ecdsa.PublicKey:
//line /usr/local/go/src/crypto/tls/auth.go:181
		_go_fuzz_dep_.CoverTab[21067]++
								if version != VersionTLS13 {
//line /usr/local/go/src/crypto/tls/auth.go:182
			_go_fuzz_dep_.CoverTab[21072]++

//line /usr/local/go/src/crypto/tls/auth.go:185
			sigAlgs = []SignatureScheme{
				ECDSAWithP256AndSHA256,
				ECDSAWithP384AndSHA384,
				ECDSAWithP521AndSHA512,
				ECDSAWithSHA1,
			}
									break
//line /usr/local/go/src/crypto/tls/auth.go:191
			// _ = "end of CoverTab[21072]"
		} else {
//line /usr/local/go/src/crypto/tls/auth.go:192
			_go_fuzz_dep_.CoverTab[21073]++
//line /usr/local/go/src/crypto/tls/auth.go:192
			// _ = "end of CoverTab[21073]"
//line /usr/local/go/src/crypto/tls/auth.go:192
		}
//line /usr/local/go/src/crypto/tls/auth.go:192
		// _ = "end of CoverTab[21067]"
//line /usr/local/go/src/crypto/tls/auth.go:192
		_go_fuzz_dep_.CoverTab[21068]++
								switch pub.Curve {
		case elliptic.P256():
//line /usr/local/go/src/crypto/tls/auth.go:194
			_go_fuzz_dep_.CoverTab[21074]++
									sigAlgs = []SignatureScheme{ECDSAWithP256AndSHA256}
//line /usr/local/go/src/crypto/tls/auth.go:195
			// _ = "end of CoverTab[21074]"
		case elliptic.P384():
//line /usr/local/go/src/crypto/tls/auth.go:196
			_go_fuzz_dep_.CoverTab[21075]++
									sigAlgs = []SignatureScheme{ECDSAWithP384AndSHA384}
//line /usr/local/go/src/crypto/tls/auth.go:197
			// _ = "end of CoverTab[21075]"
		case elliptic.P521():
//line /usr/local/go/src/crypto/tls/auth.go:198
			_go_fuzz_dep_.CoverTab[21076]++
									sigAlgs = []SignatureScheme{ECDSAWithP521AndSHA512}
//line /usr/local/go/src/crypto/tls/auth.go:199
			// _ = "end of CoverTab[21076]"
		default:
//line /usr/local/go/src/crypto/tls/auth.go:200
			_go_fuzz_dep_.CoverTab[21077]++
									return nil
//line /usr/local/go/src/crypto/tls/auth.go:201
			// _ = "end of CoverTab[21077]"
		}
//line /usr/local/go/src/crypto/tls/auth.go:202
		// _ = "end of CoverTab[21068]"
	case *rsa.PublicKey:
//line /usr/local/go/src/crypto/tls/auth.go:203
		_go_fuzz_dep_.CoverTab[21069]++
								size := pub.Size()
								sigAlgs = make([]SignatureScheme, 0, len(rsaSignatureSchemes))
								for _, candidate := range rsaSignatureSchemes {
//line /usr/local/go/src/crypto/tls/auth.go:206
			_go_fuzz_dep_.CoverTab[21078]++
									if size >= candidate.minModulusBytes && func() bool {
//line /usr/local/go/src/crypto/tls/auth.go:207
				_go_fuzz_dep_.CoverTab[21079]++
//line /usr/local/go/src/crypto/tls/auth.go:207
				return version <= candidate.maxVersion
//line /usr/local/go/src/crypto/tls/auth.go:207
				// _ = "end of CoverTab[21079]"
//line /usr/local/go/src/crypto/tls/auth.go:207
			}() {
//line /usr/local/go/src/crypto/tls/auth.go:207
				_go_fuzz_dep_.CoverTab[21080]++
										sigAlgs = append(sigAlgs, candidate.scheme)
//line /usr/local/go/src/crypto/tls/auth.go:208
				// _ = "end of CoverTab[21080]"
			} else {
//line /usr/local/go/src/crypto/tls/auth.go:209
				_go_fuzz_dep_.CoverTab[21081]++
//line /usr/local/go/src/crypto/tls/auth.go:209
				// _ = "end of CoverTab[21081]"
//line /usr/local/go/src/crypto/tls/auth.go:209
			}
//line /usr/local/go/src/crypto/tls/auth.go:209
			// _ = "end of CoverTab[21078]"
		}
//line /usr/local/go/src/crypto/tls/auth.go:210
		// _ = "end of CoverTab[21069]"
	case ed25519.PublicKey:
//line /usr/local/go/src/crypto/tls/auth.go:211
		_go_fuzz_dep_.CoverTab[21070]++
								sigAlgs = []SignatureScheme{Ed25519}
//line /usr/local/go/src/crypto/tls/auth.go:212
		// _ = "end of CoverTab[21070]"
	default:
//line /usr/local/go/src/crypto/tls/auth.go:213
		_go_fuzz_dep_.CoverTab[21071]++
								return nil
//line /usr/local/go/src/crypto/tls/auth.go:214
		// _ = "end of CoverTab[21071]"
	}
//line /usr/local/go/src/crypto/tls/auth.go:215
	// _ = "end of CoverTab[21062]"
//line /usr/local/go/src/crypto/tls/auth.go:215
	_go_fuzz_dep_.CoverTab[21063]++

							if cert.SupportedSignatureAlgorithms != nil {
//line /usr/local/go/src/crypto/tls/auth.go:217
		_go_fuzz_dep_.CoverTab[21082]++
								var filteredSigAlgs []SignatureScheme
								for _, sigAlg := range sigAlgs {
//line /usr/local/go/src/crypto/tls/auth.go:219
			_go_fuzz_dep_.CoverTab[21084]++
									if isSupportedSignatureAlgorithm(sigAlg, cert.SupportedSignatureAlgorithms) {
//line /usr/local/go/src/crypto/tls/auth.go:220
				_go_fuzz_dep_.CoverTab[21085]++
										filteredSigAlgs = append(filteredSigAlgs, sigAlg)
//line /usr/local/go/src/crypto/tls/auth.go:221
				// _ = "end of CoverTab[21085]"
			} else {
//line /usr/local/go/src/crypto/tls/auth.go:222
				_go_fuzz_dep_.CoverTab[21086]++
//line /usr/local/go/src/crypto/tls/auth.go:222
				// _ = "end of CoverTab[21086]"
//line /usr/local/go/src/crypto/tls/auth.go:222
			}
//line /usr/local/go/src/crypto/tls/auth.go:222
			// _ = "end of CoverTab[21084]"
		}
//line /usr/local/go/src/crypto/tls/auth.go:223
		// _ = "end of CoverTab[21082]"
//line /usr/local/go/src/crypto/tls/auth.go:223
		_go_fuzz_dep_.CoverTab[21083]++
								return filteredSigAlgs
//line /usr/local/go/src/crypto/tls/auth.go:224
		// _ = "end of CoverTab[21083]"
	} else {
//line /usr/local/go/src/crypto/tls/auth.go:225
		_go_fuzz_dep_.CoverTab[21087]++
//line /usr/local/go/src/crypto/tls/auth.go:225
		// _ = "end of CoverTab[21087]"
//line /usr/local/go/src/crypto/tls/auth.go:225
	}
//line /usr/local/go/src/crypto/tls/auth.go:225
	// _ = "end of CoverTab[21063]"
//line /usr/local/go/src/crypto/tls/auth.go:225
	_go_fuzz_dep_.CoverTab[21064]++
							return sigAlgs
//line /usr/local/go/src/crypto/tls/auth.go:226
	// _ = "end of CoverTab[21064]"
}

// selectSignatureScheme picks a SignatureScheme from the peer's preference list
//line /usr/local/go/src/crypto/tls/auth.go:229
// that works with the selected certificate. It's only called for protocol
//line /usr/local/go/src/crypto/tls/auth.go:229
// versions that support signature algorithms, so TLS 1.2 and 1.3.
//line /usr/local/go/src/crypto/tls/auth.go:232
func selectSignatureScheme(vers uint16, c *Certificate, peerAlgs []SignatureScheme) (SignatureScheme, error) {
//line /usr/local/go/src/crypto/tls/auth.go:232
	_go_fuzz_dep_.CoverTab[21088]++
							supportedAlgs := signatureSchemesForCertificate(vers, c)
							if len(supportedAlgs) == 0 {
//line /usr/local/go/src/crypto/tls/auth.go:234
		_go_fuzz_dep_.CoverTab[21092]++
								return 0, unsupportedCertificateError(c)
//line /usr/local/go/src/crypto/tls/auth.go:235
		// _ = "end of CoverTab[21092]"
	} else {
//line /usr/local/go/src/crypto/tls/auth.go:236
		_go_fuzz_dep_.CoverTab[21093]++
//line /usr/local/go/src/crypto/tls/auth.go:236
		// _ = "end of CoverTab[21093]"
//line /usr/local/go/src/crypto/tls/auth.go:236
	}
//line /usr/local/go/src/crypto/tls/auth.go:236
	// _ = "end of CoverTab[21088]"
//line /usr/local/go/src/crypto/tls/auth.go:236
	_go_fuzz_dep_.CoverTab[21089]++
							if len(peerAlgs) == 0 && func() bool {
//line /usr/local/go/src/crypto/tls/auth.go:237
		_go_fuzz_dep_.CoverTab[21094]++
//line /usr/local/go/src/crypto/tls/auth.go:237
		return vers == VersionTLS12
//line /usr/local/go/src/crypto/tls/auth.go:237
		// _ = "end of CoverTab[21094]"
//line /usr/local/go/src/crypto/tls/auth.go:237
	}() {
//line /usr/local/go/src/crypto/tls/auth.go:237
		_go_fuzz_dep_.CoverTab[21095]++

//line /usr/local/go/src/crypto/tls/auth.go:240
		peerAlgs = []SignatureScheme{PKCS1WithSHA1, ECDSAWithSHA1}
//line /usr/local/go/src/crypto/tls/auth.go:240
		// _ = "end of CoverTab[21095]"
	} else {
//line /usr/local/go/src/crypto/tls/auth.go:241
		_go_fuzz_dep_.CoverTab[21096]++
//line /usr/local/go/src/crypto/tls/auth.go:241
		// _ = "end of CoverTab[21096]"
//line /usr/local/go/src/crypto/tls/auth.go:241
	}
//line /usr/local/go/src/crypto/tls/auth.go:241
	// _ = "end of CoverTab[21089]"
//line /usr/local/go/src/crypto/tls/auth.go:241
	_go_fuzz_dep_.CoverTab[21090]++

//line /usr/local/go/src/crypto/tls/auth.go:244
	for _, preferredAlg := range peerAlgs {
//line /usr/local/go/src/crypto/tls/auth.go:244
		_go_fuzz_dep_.CoverTab[21097]++
								if needFIPS() && func() bool {
//line /usr/local/go/src/crypto/tls/auth.go:245
			_go_fuzz_dep_.CoverTab[21099]++
//line /usr/local/go/src/crypto/tls/auth.go:245
			return !isSupportedSignatureAlgorithm(preferredAlg, fipsSupportedSignatureAlgorithms)
//line /usr/local/go/src/crypto/tls/auth.go:245
			// _ = "end of CoverTab[21099]"
//line /usr/local/go/src/crypto/tls/auth.go:245
		}() {
//line /usr/local/go/src/crypto/tls/auth.go:245
			_go_fuzz_dep_.CoverTab[21100]++
									continue
//line /usr/local/go/src/crypto/tls/auth.go:246
			// _ = "end of CoverTab[21100]"
		} else {
//line /usr/local/go/src/crypto/tls/auth.go:247
			_go_fuzz_dep_.CoverTab[21101]++
//line /usr/local/go/src/crypto/tls/auth.go:247
			// _ = "end of CoverTab[21101]"
//line /usr/local/go/src/crypto/tls/auth.go:247
		}
//line /usr/local/go/src/crypto/tls/auth.go:247
		// _ = "end of CoverTab[21097]"
//line /usr/local/go/src/crypto/tls/auth.go:247
		_go_fuzz_dep_.CoverTab[21098]++
								if isSupportedSignatureAlgorithm(preferredAlg, supportedAlgs) {
//line /usr/local/go/src/crypto/tls/auth.go:248
			_go_fuzz_dep_.CoverTab[21102]++
									return preferredAlg, nil
//line /usr/local/go/src/crypto/tls/auth.go:249
			// _ = "end of CoverTab[21102]"
		} else {
//line /usr/local/go/src/crypto/tls/auth.go:250
			_go_fuzz_dep_.CoverTab[21103]++
//line /usr/local/go/src/crypto/tls/auth.go:250
			// _ = "end of CoverTab[21103]"
//line /usr/local/go/src/crypto/tls/auth.go:250
		}
//line /usr/local/go/src/crypto/tls/auth.go:250
		// _ = "end of CoverTab[21098]"
	}
//line /usr/local/go/src/crypto/tls/auth.go:251
	// _ = "end of CoverTab[21090]"
//line /usr/local/go/src/crypto/tls/auth.go:251
	_go_fuzz_dep_.CoverTab[21091]++
							return 0, errors.New("tls: peer doesn't support any of the certificate's signature algorithms")
//line /usr/local/go/src/crypto/tls/auth.go:252
	// _ = "end of CoverTab[21091]"
}

// unsupportedCertificateError returns a helpful error for certificates with
//line /usr/local/go/src/crypto/tls/auth.go:255
// an unsupported private key.
//line /usr/local/go/src/crypto/tls/auth.go:257
func unsupportedCertificateError(cert *Certificate) error {
//line /usr/local/go/src/crypto/tls/auth.go:257
	_go_fuzz_dep_.CoverTab[21104]++
							switch cert.PrivateKey.(type) {
	case rsa.PrivateKey, ecdsa.PrivateKey:
//line /usr/local/go/src/crypto/tls/auth.go:259
		_go_fuzz_dep_.CoverTab[21109]++
								return fmt.Errorf("tls: unsupported certificate: private key is %T, expected *%T",
			cert.PrivateKey, cert.PrivateKey)
//line /usr/local/go/src/crypto/tls/auth.go:261
		// _ = "end of CoverTab[21109]"
	case *ed25519.PrivateKey:
//line /usr/local/go/src/crypto/tls/auth.go:262
		_go_fuzz_dep_.CoverTab[21110]++
								return fmt.Errorf("tls: unsupported certificate: private key is *ed25519.PrivateKey, expected ed25519.PrivateKey")
//line /usr/local/go/src/crypto/tls/auth.go:263
		// _ = "end of CoverTab[21110]"
	}
//line /usr/local/go/src/crypto/tls/auth.go:264
	// _ = "end of CoverTab[21104]"
//line /usr/local/go/src/crypto/tls/auth.go:264
	_go_fuzz_dep_.CoverTab[21105]++

							signer, ok := cert.PrivateKey.(crypto.Signer)
							if !ok {
//line /usr/local/go/src/crypto/tls/auth.go:267
		_go_fuzz_dep_.CoverTab[21111]++
								return fmt.Errorf("tls: certificate private key (%T) does not implement crypto.Signer",
			cert.PrivateKey)
//line /usr/local/go/src/crypto/tls/auth.go:269
		// _ = "end of CoverTab[21111]"
	} else {
//line /usr/local/go/src/crypto/tls/auth.go:270
		_go_fuzz_dep_.CoverTab[21112]++
//line /usr/local/go/src/crypto/tls/auth.go:270
		// _ = "end of CoverTab[21112]"
//line /usr/local/go/src/crypto/tls/auth.go:270
	}
//line /usr/local/go/src/crypto/tls/auth.go:270
	// _ = "end of CoverTab[21105]"
//line /usr/local/go/src/crypto/tls/auth.go:270
	_go_fuzz_dep_.CoverTab[21106]++

							switch pub := signer.Public().(type) {
	case *ecdsa.PublicKey:
//line /usr/local/go/src/crypto/tls/auth.go:273
		_go_fuzz_dep_.CoverTab[21113]++
								switch pub.Curve {
		case elliptic.P256():
//line /usr/local/go/src/crypto/tls/auth.go:275
			_go_fuzz_dep_.CoverTab[21117]++
//line /usr/local/go/src/crypto/tls/auth.go:275
			// _ = "end of CoverTab[21117]"
		case elliptic.P384():
//line /usr/local/go/src/crypto/tls/auth.go:276
			_go_fuzz_dep_.CoverTab[21118]++
//line /usr/local/go/src/crypto/tls/auth.go:276
			// _ = "end of CoverTab[21118]"
		case elliptic.P521():
//line /usr/local/go/src/crypto/tls/auth.go:277
			_go_fuzz_dep_.CoverTab[21119]++
//line /usr/local/go/src/crypto/tls/auth.go:277
			// _ = "end of CoverTab[21119]"
		default:
//line /usr/local/go/src/crypto/tls/auth.go:278
			_go_fuzz_dep_.CoverTab[21120]++
									return fmt.Errorf("tls: unsupported certificate curve (%s)", pub.Curve.Params().Name)
//line /usr/local/go/src/crypto/tls/auth.go:279
			// _ = "end of CoverTab[21120]"
		}
//line /usr/local/go/src/crypto/tls/auth.go:280
		// _ = "end of CoverTab[21113]"
	case *rsa.PublicKey:
//line /usr/local/go/src/crypto/tls/auth.go:281
		_go_fuzz_dep_.CoverTab[21114]++
								return fmt.Errorf("tls: certificate RSA key size too small for supported signature algorithms")
//line /usr/local/go/src/crypto/tls/auth.go:282
		// _ = "end of CoverTab[21114]"
	case ed25519.PublicKey:
//line /usr/local/go/src/crypto/tls/auth.go:283
		_go_fuzz_dep_.CoverTab[21115]++
//line /usr/local/go/src/crypto/tls/auth.go:283
		// _ = "end of CoverTab[21115]"
	default:
//line /usr/local/go/src/crypto/tls/auth.go:284
		_go_fuzz_dep_.CoverTab[21116]++
								return fmt.Errorf("tls: unsupported certificate key (%T)", pub)
//line /usr/local/go/src/crypto/tls/auth.go:285
		// _ = "end of CoverTab[21116]"
	}
//line /usr/local/go/src/crypto/tls/auth.go:286
	// _ = "end of CoverTab[21106]"
//line /usr/local/go/src/crypto/tls/auth.go:286
	_go_fuzz_dep_.CoverTab[21107]++

							if cert.SupportedSignatureAlgorithms != nil {
//line /usr/local/go/src/crypto/tls/auth.go:288
		_go_fuzz_dep_.CoverTab[21121]++
								return fmt.Errorf("tls: peer doesn't support the certificate custom signature algorithms")
//line /usr/local/go/src/crypto/tls/auth.go:289
		// _ = "end of CoverTab[21121]"
	} else {
//line /usr/local/go/src/crypto/tls/auth.go:290
		_go_fuzz_dep_.CoverTab[21122]++
//line /usr/local/go/src/crypto/tls/auth.go:290
		// _ = "end of CoverTab[21122]"
//line /usr/local/go/src/crypto/tls/auth.go:290
	}
//line /usr/local/go/src/crypto/tls/auth.go:290
	// _ = "end of CoverTab[21107]"
//line /usr/local/go/src/crypto/tls/auth.go:290
	_go_fuzz_dep_.CoverTab[21108]++

							return fmt.Errorf("tls: internal error: unsupported key (%T)", cert.PrivateKey)
//line /usr/local/go/src/crypto/tls/auth.go:292
	// _ = "end of CoverTab[21108]"
}

//line /usr/local/go/src/crypto/tls/auth.go:293
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/auth.go:293
var _ = _go_fuzz_dep_.CoverTab
