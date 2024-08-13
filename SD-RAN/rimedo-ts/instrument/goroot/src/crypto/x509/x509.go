// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/x509/x509.go:5
// Package x509 implements a subset of the X.509 standard.
//line /usr/local/go/src/crypto/x509/x509.go:5
//
//line /usr/local/go/src/crypto/x509/x509.go:5
// It allows parsing and generating certificates, certificate signing
//line /usr/local/go/src/crypto/x509/x509.go:5
// requests, certificate revocation lists, and encoded public and private keys.
//line /usr/local/go/src/crypto/x509/x509.go:5
// It provides a certificate verifier, complete with a chain builder.
//line /usr/local/go/src/crypto/x509/x509.go:5
//
//line /usr/local/go/src/crypto/x509/x509.go:5
// The package targets the X.509 technical profile defined by the IETF (RFC
//line /usr/local/go/src/crypto/x509/x509.go:5
// 2459/3280/5280), and as further restricted by the CA/Browser Forum Baseline
//line /usr/local/go/src/crypto/x509/x509.go:5
// Requirements. There is minimal support for features outside of these
//line /usr/local/go/src/crypto/x509/x509.go:5
// profiles, as the primary goal of the package is to provide compatibility
//line /usr/local/go/src/crypto/x509/x509.go:5
// with the publicly trusted TLS certificate ecosystem and its policies and
//line /usr/local/go/src/crypto/x509/x509.go:5
// constraints.
//line /usr/local/go/src/crypto/x509/x509.go:5
//
//line /usr/local/go/src/crypto/x509/x509.go:5
// On macOS and Windows, certificate verification is handled by system APIs, but
//line /usr/local/go/src/crypto/x509/x509.go:5
// the package aims to apply consistent validation rules across operating
//line /usr/local/go/src/crypto/x509/x509.go:5
// systems.
//line /usr/local/go/src/crypto/x509/x509.go:21
package x509

//line /usr/local/go/src/crypto/x509/x509.go:21
import (
//line /usr/local/go/src/crypto/x509/x509.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/x509/x509.go:21
)
//line /usr/local/go/src/crypto/x509/x509.go:21
import (
//line /usr/local/go/src/crypto/x509/x509.go:21
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/x509/x509.go:21
)

import (
	"bytes"
	"crypto"
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"internal/godebug"
	"io"
	"math/big"
	"net"
	"net/url"
	"strconv"
	"time"
	"unicode"

	// Explicitly import these for their crypto.RegisterHash init side-effects.
	// Keep these as blank imports, even if they're imported above.
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"

	"golang.org/x/crypto/cryptobyte"
	cryptobyte_asn1 "golang.org/x/crypto/cryptobyte/asn1"
)

// pkixPublicKey reflects a PKIX public key structure. See SubjectPublicKeyInfo
//line /usr/local/go/src/crypto/x509/x509.go:56
// in RFC 3280.
//line /usr/local/go/src/crypto/x509/x509.go:58
type pkixPublicKey struct {
	Algo		pkix.AlgorithmIdentifier
	BitString	asn1.BitString
}

// ParsePKIXPublicKey parses a public key in PKIX, ASN.1 DER form. The encoded
//line /usr/local/go/src/crypto/x509/x509.go:63
// public key is a SubjectPublicKeyInfo structure (see RFC 5280, Section 4.1).
//line /usr/local/go/src/crypto/x509/x509.go:63
//
//line /usr/local/go/src/crypto/x509/x509.go:63
// It returns a *rsa.PublicKey, *dsa.PublicKey, *ecdsa.PublicKey,
//line /usr/local/go/src/crypto/x509/x509.go:63
// ed25519.PublicKey (not a pointer), or *ecdh.PublicKey (for X25519).
//line /usr/local/go/src/crypto/x509/x509.go:63
// More types might be supported in the future.
//line /usr/local/go/src/crypto/x509/x509.go:63
//
//line /usr/local/go/src/crypto/x509/x509.go:63
// This kind of key is commonly encoded in PEM blocks of type "PUBLIC KEY".
//line /usr/local/go/src/crypto/x509/x509.go:71
func ParsePKIXPublicKey(derBytes []byte) (pub any, err error) {
//line /usr/local/go/src/crypto/x509/x509.go:71
	_go_fuzz_dep_.CoverTab[19909]++
							var pki publicKeyInfo
							if rest, err := asn1.Unmarshal(derBytes, &pki); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:73
		_go_fuzz_dep_.CoverTab[19911]++
								if _, err := asn1.Unmarshal(derBytes, &pkcs1PublicKey{}); err == nil {
//line /usr/local/go/src/crypto/x509/x509.go:74
			_go_fuzz_dep_.CoverTab[19913]++
									return nil, errors.New("x509: failed to parse public key (use ParsePKCS1PublicKey instead for this key format)")
//line /usr/local/go/src/crypto/x509/x509.go:75
			// _ = "end of CoverTab[19913]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:76
			_go_fuzz_dep_.CoverTab[19914]++
//line /usr/local/go/src/crypto/x509/x509.go:76
			// _ = "end of CoverTab[19914]"
//line /usr/local/go/src/crypto/x509/x509.go:76
		}
//line /usr/local/go/src/crypto/x509/x509.go:76
		// _ = "end of CoverTab[19911]"
//line /usr/local/go/src/crypto/x509/x509.go:76
		_go_fuzz_dep_.CoverTab[19912]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:77
		// _ = "end of CoverTab[19912]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:78
		_go_fuzz_dep_.CoverTab[19915]++
//line /usr/local/go/src/crypto/x509/x509.go:78
		if len(rest) != 0 {
//line /usr/local/go/src/crypto/x509/x509.go:78
			_go_fuzz_dep_.CoverTab[19916]++
									return nil, errors.New("x509: trailing data after ASN.1 of public-key")
//line /usr/local/go/src/crypto/x509/x509.go:79
			// _ = "end of CoverTab[19916]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:80
			_go_fuzz_dep_.CoverTab[19917]++
//line /usr/local/go/src/crypto/x509/x509.go:80
			// _ = "end of CoverTab[19917]"
//line /usr/local/go/src/crypto/x509/x509.go:80
		}
//line /usr/local/go/src/crypto/x509/x509.go:80
		// _ = "end of CoverTab[19915]"
//line /usr/local/go/src/crypto/x509/x509.go:80
	}
//line /usr/local/go/src/crypto/x509/x509.go:80
	// _ = "end of CoverTab[19909]"
//line /usr/local/go/src/crypto/x509/x509.go:80
	_go_fuzz_dep_.CoverTab[19910]++
							return parsePublicKey(&pki)
//line /usr/local/go/src/crypto/x509/x509.go:81
	// _ = "end of CoverTab[19910]"
}

func marshalPublicKey(pub any) (publicKeyBytes []byte, publicKeyAlgorithm pkix.AlgorithmIdentifier, err error) {
//line /usr/local/go/src/crypto/x509/x509.go:84
	_go_fuzz_dep_.CoverTab[19918]++
							switch pub := pub.(type) {
	case *rsa.PublicKey:
//line /usr/local/go/src/crypto/x509/x509.go:86
		_go_fuzz_dep_.CoverTab[19920]++
								publicKeyBytes, err = asn1.Marshal(pkcs1PublicKey{
			N:	pub.N,
			E:	pub.E,
		})
		if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:91
			_go_fuzz_dep_.CoverTab[19929]++
									return nil, pkix.AlgorithmIdentifier{}, err
//line /usr/local/go/src/crypto/x509/x509.go:92
			// _ = "end of CoverTab[19929]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:93
			_go_fuzz_dep_.CoverTab[19930]++
//line /usr/local/go/src/crypto/x509/x509.go:93
			// _ = "end of CoverTab[19930]"
//line /usr/local/go/src/crypto/x509/x509.go:93
		}
//line /usr/local/go/src/crypto/x509/x509.go:93
		// _ = "end of CoverTab[19920]"
//line /usr/local/go/src/crypto/x509/x509.go:93
		_go_fuzz_dep_.CoverTab[19921]++
								publicKeyAlgorithm.Algorithm = oidPublicKeyRSA

//line /usr/local/go/src/crypto/x509/x509.go:97
		publicKeyAlgorithm.Parameters = asn1.NullRawValue
//line /usr/local/go/src/crypto/x509/x509.go:97
		// _ = "end of CoverTab[19921]"
	case *ecdsa.PublicKey:
//line /usr/local/go/src/crypto/x509/x509.go:98
		_go_fuzz_dep_.CoverTab[19922]++
								oid, ok := oidFromNamedCurve(pub.Curve)
								if !ok {
//line /usr/local/go/src/crypto/x509/x509.go:100
			_go_fuzz_dep_.CoverTab[19931]++
									return nil, pkix.AlgorithmIdentifier{}, errors.New("x509: unsupported elliptic curve")
//line /usr/local/go/src/crypto/x509/x509.go:101
			// _ = "end of CoverTab[19931]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:102
			_go_fuzz_dep_.CoverTab[19932]++
//line /usr/local/go/src/crypto/x509/x509.go:102
			// _ = "end of CoverTab[19932]"
//line /usr/local/go/src/crypto/x509/x509.go:102
		}
//line /usr/local/go/src/crypto/x509/x509.go:102
		// _ = "end of CoverTab[19922]"
//line /usr/local/go/src/crypto/x509/x509.go:102
		_go_fuzz_dep_.CoverTab[19923]++
								if !pub.Curve.IsOnCurve(pub.X, pub.Y) {
//line /usr/local/go/src/crypto/x509/x509.go:103
			_go_fuzz_dep_.CoverTab[19933]++
									return nil, pkix.AlgorithmIdentifier{}, errors.New("x509: invalid elliptic curve public key")
//line /usr/local/go/src/crypto/x509/x509.go:104
			// _ = "end of CoverTab[19933]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:105
			_go_fuzz_dep_.CoverTab[19934]++
//line /usr/local/go/src/crypto/x509/x509.go:105
			// _ = "end of CoverTab[19934]"
//line /usr/local/go/src/crypto/x509/x509.go:105
		}
//line /usr/local/go/src/crypto/x509/x509.go:105
		// _ = "end of CoverTab[19923]"
//line /usr/local/go/src/crypto/x509/x509.go:105
		_go_fuzz_dep_.CoverTab[19924]++
								publicKeyBytes = elliptic.Marshal(pub.Curve, pub.X, pub.Y)
								publicKeyAlgorithm.Algorithm = oidPublicKeyECDSA
								var paramBytes []byte
								paramBytes, err = asn1.Marshal(oid)
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:110
			_go_fuzz_dep_.CoverTab[19935]++
									return
//line /usr/local/go/src/crypto/x509/x509.go:111
			// _ = "end of CoverTab[19935]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:112
			_go_fuzz_dep_.CoverTab[19936]++
//line /usr/local/go/src/crypto/x509/x509.go:112
			// _ = "end of CoverTab[19936]"
//line /usr/local/go/src/crypto/x509/x509.go:112
		}
//line /usr/local/go/src/crypto/x509/x509.go:112
		// _ = "end of CoverTab[19924]"
//line /usr/local/go/src/crypto/x509/x509.go:112
		_go_fuzz_dep_.CoverTab[19925]++
								publicKeyAlgorithm.Parameters.FullBytes = paramBytes
//line /usr/local/go/src/crypto/x509/x509.go:113
		// _ = "end of CoverTab[19925]"
	case ed25519.PublicKey:
//line /usr/local/go/src/crypto/x509/x509.go:114
		_go_fuzz_dep_.CoverTab[19926]++
								publicKeyBytes = pub
								publicKeyAlgorithm.Algorithm = oidPublicKeyEd25519
//line /usr/local/go/src/crypto/x509/x509.go:116
		// _ = "end of CoverTab[19926]"
	case *ecdh.PublicKey:
//line /usr/local/go/src/crypto/x509/x509.go:117
		_go_fuzz_dep_.CoverTab[19927]++
								publicKeyBytes = pub.Bytes()
								if pub.Curve() == ecdh.X25519() {
//line /usr/local/go/src/crypto/x509/x509.go:119
			_go_fuzz_dep_.CoverTab[19937]++
									publicKeyAlgorithm.Algorithm = oidPublicKeyX25519
//line /usr/local/go/src/crypto/x509/x509.go:120
			// _ = "end of CoverTab[19937]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:121
			_go_fuzz_dep_.CoverTab[19938]++
									oid, ok := oidFromECDHCurve(pub.Curve())
									if !ok {
//line /usr/local/go/src/crypto/x509/x509.go:123
				_go_fuzz_dep_.CoverTab[19941]++
										return nil, pkix.AlgorithmIdentifier{}, errors.New("x509: unsupported elliptic curve")
//line /usr/local/go/src/crypto/x509/x509.go:124
				// _ = "end of CoverTab[19941]"
			} else {
//line /usr/local/go/src/crypto/x509/x509.go:125
				_go_fuzz_dep_.CoverTab[19942]++
//line /usr/local/go/src/crypto/x509/x509.go:125
				// _ = "end of CoverTab[19942]"
//line /usr/local/go/src/crypto/x509/x509.go:125
			}
//line /usr/local/go/src/crypto/x509/x509.go:125
			// _ = "end of CoverTab[19938]"
//line /usr/local/go/src/crypto/x509/x509.go:125
			_go_fuzz_dep_.CoverTab[19939]++
									publicKeyAlgorithm.Algorithm = oidPublicKeyECDSA
									var paramBytes []byte
									paramBytes, err = asn1.Marshal(oid)
									if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:129
				_go_fuzz_dep_.CoverTab[19943]++
										return
//line /usr/local/go/src/crypto/x509/x509.go:130
				// _ = "end of CoverTab[19943]"
			} else {
//line /usr/local/go/src/crypto/x509/x509.go:131
				_go_fuzz_dep_.CoverTab[19944]++
//line /usr/local/go/src/crypto/x509/x509.go:131
				// _ = "end of CoverTab[19944]"
//line /usr/local/go/src/crypto/x509/x509.go:131
			}
//line /usr/local/go/src/crypto/x509/x509.go:131
			// _ = "end of CoverTab[19939]"
//line /usr/local/go/src/crypto/x509/x509.go:131
			_go_fuzz_dep_.CoverTab[19940]++
									publicKeyAlgorithm.Parameters.FullBytes = paramBytes
//line /usr/local/go/src/crypto/x509/x509.go:132
			// _ = "end of CoverTab[19940]"
		}
//line /usr/local/go/src/crypto/x509/x509.go:133
		// _ = "end of CoverTab[19927]"
	default:
//line /usr/local/go/src/crypto/x509/x509.go:134
		_go_fuzz_dep_.CoverTab[19928]++
								return nil, pkix.AlgorithmIdentifier{}, fmt.Errorf("x509: unsupported public key type: %T", pub)
//line /usr/local/go/src/crypto/x509/x509.go:135
		// _ = "end of CoverTab[19928]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:136
	// _ = "end of CoverTab[19918]"
//line /usr/local/go/src/crypto/x509/x509.go:136
	_go_fuzz_dep_.CoverTab[19919]++

							return publicKeyBytes, publicKeyAlgorithm, nil
//line /usr/local/go/src/crypto/x509/x509.go:138
	// _ = "end of CoverTab[19919]"
}

// MarshalPKIXPublicKey converts a public key to PKIX, ASN.1 DER form.
//line /usr/local/go/src/crypto/x509/x509.go:141
// The encoded public key is a SubjectPublicKeyInfo structure
//line /usr/local/go/src/crypto/x509/x509.go:141
// (see RFC 5280, Section 4.1).
//line /usr/local/go/src/crypto/x509/x509.go:141
//
//line /usr/local/go/src/crypto/x509/x509.go:141
// The following key types are currently supported: *rsa.PublicKey,
//line /usr/local/go/src/crypto/x509/x509.go:141
// *ecdsa.PublicKey, ed25519.PublicKey (not a pointer), and *ecdh.PublicKey.
//line /usr/local/go/src/crypto/x509/x509.go:141
// Unsupported key types result in an error.
//line /usr/local/go/src/crypto/x509/x509.go:141
//
//line /usr/local/go/src/crypto/x509/x509.go:141
// This kind of key is commonly encoded in PEM blocks of type "PUBLIC KEY".
//line /usr/local/go/src/crypto/x509/x509.go:150
func MarshalPKIXPublicKey(pub any) ([]byte, error) {
//line /usr/local/go/src/crypto/x509/x509.go:150
	_go_fuzz_dep_.CoverTab[19945]++
							var publicKeyBytes []byte
							var publicKeyAlgorithm pkix.AlgorithmIdentifier
							var err error

							if publicKeyBytes, publicKeyAlgorithm, err = marshalPublicKey(pub); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:155
		_go_fuzz_dep_.CoverTab[19947]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:156
		// _ = "end of CoverTab[19947]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:157
		_go_fuzz_dep_.CoverTab[19948]++
//line /usr/local/go/src/crypto/x509/x509.go:157
		// _ = "end of CoverTab[19948]"
//line /usr/local/go/src/crypto/x509/x509.go:157
	}
//line /usr/local/go/src/crypto/x509/x509.go:157
	// _ = "end of CoverTab[19945]"
//line /usr/local/go/src/crypto/x509/x509.go:157
	_go_fuzz_dep_.CoverTab[19946]++

							pkix := pkixPublicKey{
		Algo:	publicKeyAlgorithm,
		BitString: asn1.BitString{
			Bytes:		publicKeyBytes,
			BitLength:	8 * len(publicKeyBytes),
		},
	}

							ret, _ := asn1.Marshal(pkix)
							return ret, nil
//line /usr/local/go/src/crypto/x509/x509.go:168
	// _ = "end of CoverTab[19946]"
}

//line /usr/local/go/src/crypto/x509/x509.go:173
type certificate struct {
	TBSCertificate		tbsCertificate
	SignatureAlgorithm	pkix.AlgorithmIdentifier
	SignatureValue		asn1.BitString
}

type tbsCertificate struct {
	Raw			asn1.RawContent
	Version			int	`asn1:"optional,explicit,default:0,tag:0"`
	SerialNumber		*big.Int
	SignatureAlgorithm	pkix.AlgorithmIdentifier
	Issuer			asn1.RawValue
	Validity		validity
	Subject			asn1.RawValue
	PublicKey		publicKeyInfo
	UniqueId		asn1.BitString		`asn1:"optional,tag:1"`
	SubjectUniqueId		asn1.BitString		`asn1:"optional,tag:2"`
	Extensions		[]pkix.Extension	`asn1:"omitempty,optional,explicit,tag:3"`
}

type dsaAlgorithmParameters struct {
	P, Q, G *big.Int
}

type validity struct {
	NotBefore, NotAfter time.Time
}

type publicKeyInfo struct {
	Raw		asn1.RawContent
	Algorithm	pkix.AlgorithmIdentifier
	PublicKey	asn1.BitString
}

// RFC 5280,  4.2.1.1
type authKeyId struct {
	Id []byte `asn1:"optional,tag:0"`
}

type SignatureAlgorithm int

const (
	UnknownSignatureAlgorithm	SignatureAlgorithm	= iota

	MD2WithRSA	// Unsupported.
	MD5WithRSA	// Only supported for signing, not verification.
	SHA1WithRSA	// Only supported for signing, and verification of CRLs, CSRs, and OCSP responses.
	SHA256WithRSA
	SHA384WithRSA
	SHA512WithRSA
	DSAWithSHA1	// Unsupported.
	DSAWithSHA256	// Unsupported.
	ECDSAWithSHA1	// Only supported for signing, and verification of CRLs, CSRs, and OCSP responses.
	ECDSAWithSHA256
	ECDSAWithSHA384
	ECDSAWithSHA512
	SHA256WithRSAPSS
	SHA384WithRSAPSS
	SHA512WithRSAPSS
	PureEd25519
)

func (algo SignatureAlgorithm) isRSAPSS() bool {
//line /usr/local/go/src/crypto/x509/x509.go:235
	_go_fuzz_dep_.CoverTab[19949]++
							switch algo {
	case SHA256WithRSAPSS, SHA384WithRSAPSS, SHA512WithRSAPSS:
//line /usr/local/go/src/crypto/x509/x509.go:237
		_go_fuzz_dep_.CoverTab[19950]++
								return true
//line /usr/local/go/src/crypto/x509/x509.go:238
		// _ = "end of CoverTab[19950]"
	default:
//line /usr/local/go/src/crypto/x509/x509.go:239
		_go_fuzz_dep_.CoverTab[19951]++
								return false
//line /usr/local/go/src/crypto/x509/x509.go:240
		// _ = "end of CoverTab[19951]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:241
	// _ = "end of CoverTab[19949]"
}

func (algo SignatureAlgorithm) String() string {
//line /usr/local/go/src/crypto/x509/x509.go:244
	_go_fuzz_dep_.CoverTab[19952]++
							for _, details := range signatureAlgorithmDetails {
//line /usr/local/go/src/crypto/x509/x509.go:245
		_go_fuzz_dep_.CoverTab[19954]++
								if details.algo == algo {
//line /usr/local/go/src/crypto/x509/x509.go:246
			_go_fuzz_dep_.CoverTab[19955]++
									return details.name
//line /usr/local/go/src/crypto/x509/x509.go:247
			// _ = "end of CoverTab[19955]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:248
			_go_fuzz_dep_.CoverTab[19956]++
//line /usr/local/go/src/crypto/x509/x509.go:248
			// _ = "end of CoverTab[19956]"
//line /usr/local/go/src/crypto/x509/x509.go:248
		}
//line /usr/local/go/src/crypto/x509/x509.go:248
		// _ = "end of CoverTab[19954]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:249
	// _ = "end of CoverTab[19952]"
//line /usr/local/go/src/crypto/x509/x509.go:249
	_go_fuzz_dep_.CoverTab[19953]++
							return strconv.Itoa(int(algo))
//line /usr/local/go/src/crypto/x509/x509.go:250
	// _ = "end of CoverTab[19953]"
}

type PublicKeyAlgorithm int

const (
	UnknownPublicKeyAlgorithm	PublicKeyAlgorithm	= iota
	RSA
	DSA	// Only supported for parsing.
	ECDSA
	Ed25519
)

var publicKeyAlgoName = [...]string{
	RSA:		"RSA",
	DSA:		"DSA",
	ECDSA:		"ECDSA",
	Ed25519:	"Ed25519",
}

func (algo PublicKeyAlgorithm) String() string {
//line /usr/local/go/src/crypto/x509/x509.go:270
	_go_fuzz_dep_.CoverTab[19957]++
							if 0 < algo && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:271
		_go_fuzz_dep_.CoverTab[19959]++
//line /usr/local/go/src/crypto/x509/x509.go:271
		return int(algo) < len(publicKeyAlgoName)
//line /usr/local/go/src/crypto/x509/x509.go:271
		// _ = "end of CoverTab[19959]"
//line /usr/local/go/src/crypto/x509/x509.go:271
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:271
		_go_fuzz_dep_.CoverTab[19960]++
								return publicKeyAlgoName[algo]
//line /usr/local/go/src/crypto/x509/x509.go:272
		// _ = "end of CoverTab[19960]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:273
		_go_fuzz_dep_.CoverTab[19961]++
//line /usr/local/go/src/crypto/x509/x509.go:273
		// _ = "end of CoverTab[19961]"
//line /usr/local/go/src/crypto/x509/x509.go:273
	}
//line /usr/local/go/src/crypto/x509/x509.go:273
	// _ = "end of CoverTab[19957]"
//line /usr/local/go/src/crypto/x509/x509.go:273
	_go_fuzz_dep_.CoverTab[19958]++
							return strconv.Itoa(int(algo))
//line /usr/local/go/src/crypto/x509/x509.go:274
	// _ = "end of CoverTab[19958]"
}

// OIDs for signature algorithms
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
//	pkcs-1 OBJECT IDENTIFIER ::= {
//line /usr/local/go/src/crypto/x509/x509.go:277
//		iso(1) member-body(2) us(840) rsadsi(113549) pkcs(1) 1 }
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
// RFC 3279 2.2.1 RSA Signature Algorithms
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
//	md2WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 2 }
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
//	md5WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 4 }
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
//	sha-1WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 5 }
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
//	dsaWithSha1 OBJECT IDENTIFIER ::= {
//line /usr/local/go/src/crypto/x509/x509.go:277
//		iso(1) member-body(2) us(840) x9-57(10040) x9cm(4) 3 }
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
// RFC 3279 2.2.3 ECDSA Signature Algorithm
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
//	ecdsa-with-SHA1 OBJECT IDENTIFIER ::= {
//line /usr/local/go/src/crypto/x509/x509.go:277
//		iso(1) member-body(2) us(840) ansi-x962(10045)
//line /usr/local/go/src/crypto/x509/x509.go:277
//		signatures(4) ecdsa-with-SHA1(1)}
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
// RFC 4055 5 PKCS #1 Version 1.5
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
//	sha256WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 11 }
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
//	sha384WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 12 }
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
//	sha512WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 13 }
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
// RFC 5758 3.1 DSA Signature Algorithms
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
//	dsaWithSha256 OBJECT IDENTIFIER ::= {
//line /usr/local/go/src/crypto/x509/x509.go:277
//		joint-iso-ccitt(2) country(16) us(840) organization(1) gov(101)
//line /usr/local/go/src/crypto/x509/x509.go:277
//		csor(3) algorithms(4) id-dsa-with-sha2(3) 2}
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
// RFC 5758 3.2 ECDSA Signature Algorithm
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
//	ecdsa-with-SHA256 OBJECT IDENTIFIER ::= { iso(1) member-body(2)
//line /usr/local/go/src/crypto/x509/x509.go:277
//		us(840) ansi-X9-62(10045) signatures(4) ecdsa-with-SHA2(3) 2 }
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
//	ecdsa-with-SHA384 OBJECT IDENTIFIER ::= { iso(1) member-body(2)
//line /usr/local/go/src/crypto/x509/x509.go:277
//		us(840) ansi-X9-62(10045) signatures(4) ecdsa-with-SHA2(3) 3 }
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
//	ecdsa-with-SHA512 OBJECT IDENTIFIER ::= { iso(1) member-body(2)
//line /usr/local/go/src/crypto/x509/x509.go:277
//		us(840) ansi-X9-62(10045) signatures(4) ecdsa-with-SHA2(3) 4 }
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
// RFC 8410 3 Curve25519 and Curve448 Algorithm Identifiers
//line /usr/local/go/src/crypto/x509/x509.go:277
//
//line /usr/local/go/src/crypto/x509/x509.go:277
//	id-Ed25519   OBJECT IDENTIFIER ::= { 1 3 101 112 }
//line /usr/local/go/src/crypto/x509/x509.go:327
var (
	oidSignatureMD2WithRSA		= asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 2}
	oidSignatureMD5WithRSA		= asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 4}
	oidSignatureSHA1WithRSA		= asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 5}
	oidSignatureSHA256WithRSA	= asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 11}
	oidSignatureSHA384WithRSA	= asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 12}
	oidSignatureSHA512WithRSA	= asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 13}
	oidSignatureRSAPSS		= asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 10}
	oidSignatureDSAWithSHA1		= asn1.ObjectIdentifier{1, 2, 840, 10040, 4, 3}
	oidSignatureDSAWithSHA256	= asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 3, 2}
	oidSignatureECDSAWithSHA1	= asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 1}
	oidSignatureECDSAWithSHA256	= asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 3, 2}
	oidSignatureECDSAWithSHA384	= asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 3, 3}
	oidSignatureECDSAWithSHA512	= asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 3, 4}
	oidSignatureEd25519		= asn1.ObjectIdentifier{1, 3, 101, 112}

	oidSHA256	= asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 2, 1}
	oidSHA384	= asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 2, 2}
	oidSHA512	= asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 2, 3}

	oidMGF1	= asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 8}

	// oidISOSignatureSHA1WithRSA means the same as oidSignatureSHA1WithRSA
	// but it's specified by ISO. Microsoft's makecert.exe has been known
	// to produce certificates with this OID.
	oidISOSignatureSHA1WithRSA	= asn1.ObjectIdentifier{1, 3, 14, 3, 2, 29}
)

var signatureAlgorithmDetails = []struct {
	algo		SignatureAlgorithm
	name		string
	oid		asn1.ObjectIdentifier
	pubKeyAlgo	PublicKeyAlgorithm
	hash		crypto.Hash
}{
	{MD2WithRSA, "MD2-RSA", oidSignatureMD2WithRSA, RSA, crypto.Hash(0)},
	{MD5WithRSA, "MD5-RSA", oidSignatureMD5WithRSA, RSA, crypto.MD5},
	{SHA1WithRSA, "SHA1-RSA", oidSignatureSHA1WithRSA, RSA, crypto.SHA1},
	{SHA1WithRSA, "SHA1-RSA", oidISOSignatureSHA1WithRSA, RSA, crypto.SHA1},
	{SHA256WithRSA, "SHA256-RSA", oidSignatureSHA256WithRSA, RSA, crypto.SHA256},
	{SHA384WithRSA, "SHA384-RSA", oidSignatureSHA384WithRSA, RSA, crypto.SHA384},
	{SHA512WithRSA, "SHA512-RSA", oidSignatureSHA512WithRSA, RSA, crypto.SHA512},
	{SHA256WithRSAPSS, "SHA256-RSAPSS", oidSignatureRSAPSS, RSA, crypto.SHA256},
	{SHA384WithRSAPSS, "SHA384-RSAPSS", oidSignatureRSAPSS, RSA, crypto.SHA384},
	{SHA512WithRSAPSS, "SHA512-RSAPSS", oidSignatureRSAPSS, RSA, crypto.SHA512},
	{DSAWithSHA1, "DSA-SHA1", oidSignatureDSAWithSHA1, DSA, crypto.SHA1},
	{DSAWithSHA256, "DSA-SHA256", oidSignatureDSAWithSHA256, DSA, crypto.SHA256},
	{ECDSAWithSHA1, "ECDSA-SHA1", oidSignatureECDSAWithSHA1, ECDSA, crypto.SHA1},
	{ECDSAWithSHA256, "ECDSA-SHA256", oidSignatureECDSAWithSHA256, ECDSA, crypto.SHA256},
	{ECDSAWithSHA384, "ECDSA-SHA384", oidSignatureECDSAWithSHA384, ECDSA, crypto.SHA384},
	{ECDSAWithSHA512, "ECDSA-SHA512", oidSignatureECDSAWithSHA512, ECDSA, crypto.SHA512},
	{PureEd25519, "Ed25519", oidSignatureEd25519, Ed25519, crypto.Hash(0)},
}

// hashToPSSParameters contains the DER encoded RSA PSS parameters for the
//line /usr/local/go/src/crypto/x509/x509.go:381
// SHA256, SHA384, and SHA512 hashes as defined in RFC 3447, Appendix A.2.3.
//line /usr/local/go/src/crypto/x509/x509.go:381
// The parameters contain the following values:
//line /usr/local/go/src/crypto/x509/x509.go:381
//   - hashAlgorithm contains the associated hash identifier with NULL parameters
//line /usr/local/go/src/crypto/x509/x509.go:381
//   - maskGenAlgorithm always contains the default mgf1SHA1 identifier
//line /usr/local/go/src/crypto/x509/x509.go:381
//   - saltLength contains the length of the associated hash
//line /usr/local/go/src/crypto/x509/x509.go:381
//   - trailerField always contains the default trailerFieldBC value
//line /usr/local/go/src/crypto/x509/x509.go:388
var hashToPSSParameters = map[crypto.Hash]asn1.RawValue{
	crypto.SHA256:	asn1.RawValue{FullBytes: []byte{48, 52, 160, 15, 48, 13, 6, 9, 96, 134, 72, 1, 101, 3, 4, 2, 1, 5, 0, 161, 28, 48, 26, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 8, 48, 13, 6, 9, 96, 134, 72, 1, 101, 3, 4, 2, 1, 5, 0, 162, 3, 2, 1, 32}},
	crypto.SHA384:	asn1.RawValue{FullBytes: []byte{48, 52, 160, 15, 48, 13, 6, 9, 96, 134, 72, 1, 101, 3, 4, 2, 2, 5, 0, 161, 28, 48, 26, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 8, 48, 13, 6, 9, 96, 134, 72, 1, 101, 3, 4, 2, 2, 5, 0, 162, 3, 2, 1, 48}},
	crypto.SHA512:	asn1.RawValue{FullBytes: []byte{48, 52, 160, 15, 48, 13, 6, 9, 96, 134, 72, 1, 101, 3, 4, 2, 3, 5, 0, 161, 28, 48, 26, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 8, 48, 13, 6, 9, 96, 134, 72, 1, 101, 3, 4, 2, 3, 5, 0, 162, 3, 2, 1, 64}},
}

// pssParameters reflects the parameters in an AlgorithmIdentifier that
//line /usr/local/go/src/crypto/x509/x509.go:394
// specifies RSA PSS. See RFC 3447, Appendix A.2.3.
//line /usr/local/go/src/crypto/x509/x509.go:396
type pssParameters struct {
	// The following three fields are not marked as
	// optional because the default values specify SHA-1,
	// which is no longer suitable for use in signatures.
	Hash		pkix.AlgorithmIdentifier	`asn1:"explicit,tag:0"`
	MGF		pkix.AlgorithmIdentifier	`asn1:"explicit,tag:1"`
	SaltLength	int				`asn1:"explicit,tag:2"`
	TrailerField	int				`asn1:"optional,explicit,tag:3,default:1"`
}

func getSignatureAlgorithmFromAI(ai pkix.AlgorithmIdentifier) SignatureAlgorithm {
//line /usr/local/go/src/crypto/x509/x509.go:406
	_go_fuzz_dep_.CoverTab[19962]++
							if ai.Algorithm.Equal(oidSignatureEd25519) {
//line /usr/local/go/src/crypto/x509/x509.go:407
		_go_fuzz_dep_.CoverTab[19969]++

//line /usr/local/go/src/crypto/x509/x509.go:410
		if len(ai.Parameters.FullBytes) != 0 {
//line /usr/local/go/src/crypto/x509/x509.go:410
			_go_fuzz_dep_.CoverTab[19970]++
									return UnknownSignatureAlgorithm
//line /usr/local/go/src/crypto/x509/x509.go:411
			// _ = "end of CoverTab[19970]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:412
			_go_fuzz_dep_.CoverTab[19971]++
//line /usr/local/go/src/crypto/x509/x509.go:412
			// _ = "end of CoverTab[19971]"
//line /usr/local/go/src/crypto/x509/x509.go:412
		}
//line /usr/local/go/src/crypto/x509/x509.go:412
		// _ = "end of CoverTab[19969]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:413
		_go_fuzz_dep_.CoverTab[19972]++
//line /usr/local/go/src/crypto/x509/x509.go:413
		// _ = "end of CoverTab[19972]"
//line /usr/local/go/src/crypto/x509/x509.go:413
	}
//line /usr/local/go/src/crypto/x509/x509.go:413
	// _ = "end of CoverTab[19962]"
//line /usr/local/go/src/crypto/x509/x509.go:413
	_go_fuzz_dep_.CoverTab[19963]++

							if !ai.Algorithm.Equal(oidSignatureRSAPSS) {
//line /usr/local/go/src/crypto/x509/x509.go:415
		_go_fuzz_dep_.CoverTab[19973]++
								for _, details := range signatureAlgorithmDetails {
//line /usr/local/go/src/crypto/x509/x509.go:416
			_go_fuzz_dep_.CoverTab[19975]++
									if ai.Algorithm.Equal(details.oid) {
//line /usr/local/go/src/crypto/x509/x509.go:417
				_go_fuzz_dep_.CoverTab[19976]++
										return details.algo
//line /usr/local/go/src/crypto/x509/x509.go:418
				// _ = "end of CoverTab[19976]"
			} else {
//line /usr/local/go/src/crypto/x509/x509.go:419
				_go_fuzz_dep_.CoverTab[19977]++
//line /usr/local/go/src/crypto/x509/x509.go:419
				// _ = "end of CoverTab[19977]"
//line /usr/local/go/src/crypto/x509/x509.go:419
			}
//line /usr/local/go/src/crypto/x509/x509.go:419
			// _ = "end of CoverTab[19975]"
		}
//line /usr/local/go/src/crypto/x509/x509.go:420
		// _ = "end of CoverTab[19973]"
//line /usr/local/go/src/crypto/x509/x509.go:420
		_go_fuzz_dep_.CoverTab[19974]++
								return UnknownSignatureAlgorithm
//line /usr/local/go/src/crypto/x509/x509.go:421
		// _ = "end of CoverTab[19974]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:422
		_go_fuzz_dep_.CoverTab[19978]++
//line /usr/local/go/src/crypto/x509/x509.go:422
		// _ = "end of CoverTab[19978]"
//line /usr/local/go/src/crypto/x509/x509.go:422
	}
//line /usr/local/go/src/crypto/x509/x509.go:422
	// _ = "end of CoverTab[19963]"
//line /usr/local/go/src/crypto/x509/x509.go:422
	_go_fuzz_dep_.CoverTab[19964]++

//line /usr/local/go/src/crypto/x509/x509.go:427
	var params pssParameters
	if _, err := asn1.Unmarshal(ai.Parameters.FullBytes, &params); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:428
		_go_fuzz_dep_.CoverTab[19979]++
								return UnknownSignatureAlgorithm
//line /usr/local/go/src/crypto/x509/x509.go:429
		// _ = "end of CoverTab[19979]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:430
		_go_fuzz_dep_.CoverTab[19980]++
//line /usr/local/go/src/crypto/x509/x509.go:430
		// _ = "end of CoverTab[19980]"
//line /usr/local/go/src/crypto/x509/x509.go:430
	}
//line /usr/local/go/src/crypto/x509/x509.go:430
	// _ = "end of CoverTab[19964]"
//line /usr/local/go/src/crypto/x509/x509.go:430
	_go_fuzz_dep_.CoverTab[19965]++

							var mgf1HashFunc pkix.AlgorithmIdentifier
							if _, err := asn1.Unmarshal(params.MGF.Parameters.FullBytes, &mgf1HashFunc); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:433
		_go_fuzz_dep_.CoverTab[19981]++
								return UnknownSignatureAlgorithm
//line /usr/local/go/src/crypto/x509/x509.go:434
		// _ = "end of CoverTab[19981]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:435
		_go_fuzz_dep_.CoverTab[19982]++
//line /usr/local/go/src/crypto/x509/x509.go:435
		// _ = "end of CoverTab[19982]"
//line /usr/local/go/src/crypto/x509/x509.go:435
	}
//line /usr/local/go/src/crypto/x509/x509.go:435
	// _ = "end of CoverTab[19965]"
//line /usr/local/go/src/crypto/x509/x509.go:435
	_go_fuzz_dep_.CoverTab[19966]++

//line /usr/local/go/src/crypto/x509/x509.go:442
	if (len(params.Hash.Parameters.FullBytes) != 0 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:442
		_go_fuzz_dep_.CoverTab[19983]++
//line /usr/local/go/src/crypto/x509/x509.go:442
		return !bytes.Equal(params.Hash.Parameters.FullBytes, asn1.NullBytes)
//line /usr/local/go/src/crypto/x509/x509.go:442
		// _ = "end of CoverTab[19983]"
//line /usr/local/go/src/crypto/x509/x509.go:442
	}()) || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:442
		_go_fuzz_dep_.CoverTab[19984]++
//line /usr/local/go/src/crypto/x509/x509.go:442
		return !params.MGF.Algorithm.Equal(oidMGF1)
								// _ = "end of CoverTab[19984]"
//line /usr/local/go/src/crypto/x509/x509.go:443
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:443
		_go_fuzz_dep_.CoverTab[19985]++
//line /usr/local/go/src/crypto/x509/x509.go:443
		return !mgf1HashFunc.Algorithm.Equal(params.Hash.Algorithm)
								// _ = "end of CoverTab[19985]"
//line /usr/local/go/src/crypto/x509/x509.go:444
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:444
		_go_fuzz_dep_.CoverTab[19986]++
//line /usr/local/go/src/crypto/x509/x509.go:444
		return (len(mgf1HashFunc.Parameters.FullBytes) != 0 && func() bool {
									_go_fuzz_dep_.CoverTab[19987]++
//line /usr/local/go/src/crypto/x509/x509.go:445
			return !bytes.Equal(mgf1HashFunc.Parameters.FullBytes, asn1.NullBytes)
//line /usr/local/go/src/crypto/x509/x509.go:445
			// _ = "end of CoverTab[19987]"
//line /usr/local/go/src/crypto/x509/x509.go:445
		}())
//line /usr/local/go/src/crypto/x509/x509.go:445
		// _ = "end of CoverTab[19986]"
//line /usr/local/go/src/crypto/x509/x509.go:445
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:445
		_go_fuzz_dep_.CoverTab[19988]++
//line /usr/local/go/src/crypto/x509/x509.go:445
		return params.TrailerField != 1
								// _ = "end of CoverTab[19988]"
//line /usr/local/go/src/crypto/x509/x509.go:446
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:446
		_go_fuzz_dep_.CoverTab[19989]++
								return UnknownSignatureAlgorithm
//line /usr/local/go/src/crypto/x509/x509.go:447
		// _ = "end of CoverTab[19989]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:448
		_go_fuzz_dep_.CoverTab[19990]++
//line /usr/local/go/src/crypto/x509/x509.go:448
		// _ = "end of CoverTab[19990]"
//line /usr/local/go/src/crypto/x509/x509.go:448
	}
//line /usr/local/go/src/crypto/x509/x509.go:448
	// _ = "end of CoverTab[19966]"
//line /usr/local/go/src/crypto/x509/x509.go:448
	_go_fuzz_dep_.CoverTab[19967]++

							switch {
	case params.Hash.Algorithm.Equal(oidSHA256) && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:451
		_go_fuzz_dep_.CoverTab[19995]++
//line /usr/local/go/src/crypto/x509/x509.go:451
		return params.SaltLength == 32
//line /usr/local/go/src/crypto/x509/x509.go:451
		// _ = "end of CoverTab[19995]"
//line /usr/local/go/src/crypto/x509/x509.go:451
	}():
//line /usr/local/go/src/crypto/x509/x509.go:451
		_go_fuzz_dep_.CoverTab[19991]++
								return SHA256WithRSAPSS
//line /usr/local/go/src/crypto/x509/x509.go:452
		// _ = "end of CoverTab[19991]"
	case params.Hash.Algorithm.Equal(oidSHA384) && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:453
		_go_fuzz_dep_.CoverTab[19996]++
//line /usr/local/go/src/crypto/x509/x509.go:453
		return params.SaltLength == 48
//line /usr/local/go/src/crypto/x509/x509.go:453
		// _ = "end of CoverTab[19996]"
//line /usr/local/go/src/crypto/x509/x509.go:453
	}():
//line /usr/local/go/src/crypto/x509/x509.go:453
		_go_fuzz_dep_.CoverTab[19992]++
								return SHA384WithRSAPSS
//line /usr/local/go/src/crypto/x509/x509.go:454
		// _ = "end of CoverTab[19992]"
	case params.Hash.Algorithm.Equal(oidSHA512) && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:455
		_go_fuzz_dep_.CoverTab[19997]++
//line /usr/local/go/src/crypto/x509/x509.go:455
		return params.SaltLength == 64
//line /usr/local/go/src/crypto/x509/x509.go:455
		// _ = "end of CoverTab[19997]"
//line /usr/local/go/src/crypto/x509/x509.go:455
	}():
//line /usr/local/go/src/crypto/x509/x509.go:455
		_go_fuzz_dep_.CoverTab[19993]++
								return SHA512WithRSAPSS
//line /usr/local/go/src/crypto/x509/x509.go:456
		// _ = "end of CoverTab[19993]"
//line /usr/local/go/src/crypto/x509/x509.go:456
	default:
//line /usr/local/go/src/crypto/x509/x509.go:456
		_go_fuzz_dep_.CoverTab[19994]++
//line /usr/local/go/src/crypto/x509/x509.go:456
		// _ = "end of CoverTab[19994]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:457
	// _ = "end of CoverTab[19967]"
//line /usr/local/go/src/crypto/x509/x509.go:457
	_go_fuzz_dep_.CoverTab[19968]++

							return UnknownSignatureAlgorithm
//line /usr/local/go/src/crypto/x509/x509.go:459
	// _ = "end of CoverTab[19968]"
}

var (
	// RFC 3279, 2.3 Public Key Algorithms
	//
	//	pkcs-1 OBJECT IDENTIFIER ::== { iso(1) member-body(2) us(840)
	//		rsadsi(113549) pkcs(1) 1 }
	//
	// rsaEncryption OBJECT IDENTIFIER ::== { pkcs1-1 1 }
	//
	//	id-dsa OBJECT IDENTIFIER ::== { iso(1) member-body(2) us(840)
	//		x9-57(10040) x9cm(4) 1 }
	oidPublicKeyRSA	= asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	oidPublicKeyDSA	= asn1.ObjectIdentifier{1, 2, 840, 10040, 4, 1}
	// RFC 5480, 2.1.1 Unrestricted Algorithm Identifier and Parameters
	//
	//	id-ecPublicKey OBJECT IDENTIFIER ::= {
	//		iso(1) member-body(2) us(840) ansi-X9-62(10045) keyType(2) 1 }
	oidPublicKeyECDSA	= asn1.ObjectIdentifier{1, 2, 840, 10045, 2, 1}
	// RFC 8410, Section 3
	//
	//	id-X25519    OBJECT IDENTIFIER ::= { 1 3 101 110 }
	//	id-Ed25519   OBJECT IDENTIFIER ::= { 1 3 101 112 }
	oidPublicKeyX25519	= asn1.ObjectIdentifier{1, 3, 101, 110}
	oidPublicKeyEd25519	= asn1.ObjectIdentifier{1, 3, 101, 112}
)

// getPublicKeyAlgorithmFromOID returns the exposed PublicKeyAlgorithm
//line /usr/local/go/src/crypto/x509/x509.go:487
// identifier for public key types supported in certificates and CSRs. Marshal
//line /usr/local/go/src/crypto/x509/x509.go:487
// and Parse functions may support a different set of public key types.
//line /usr/local/go/src/crypto/x509/x509.go:490
func getPublicKeyAlgorithmFromOID(oid asn1.ObjectIdentifier) PublicKeyAlgorithm {
//line /usr/local/go/src/crypto/x509/x509.go:490
	_go_fuzz_dep_.CoverTab[19998]++
							switch {
	case oid.Equal(oidPublicKeyRSA):
//line /usr/local/go/src/crypto/x509/x509.go:492
		_go_fuzz_dep_.CoverTab[20000]++
								return RSA
//line /usr/local/go/src/crypto/x509/x509.go:493
		// _ = "end of CoverTab[20000]"
	case oid.Equal(oidPublicKeyDSA):
//line /usr/local/go/src/crypto/x509/x509.go:494
		_go_fuzz_dep_.CoverTab[20001]++
								return DSA
//line /usr/local/go/src/crypto/x509/x509.go:495
		// _ = "end of CoverTab[20001]"
	case oid.Equal(oidPublicKeyECDSA):
//line /usr/local/go/src/crypto/x509/x509.go:496
		_go_fuzz_dep_.CoverTab[20002]++
								return ECDSA
//line /usr/local/go/src/crypto/x509/x509.go:497
		// _ = "end of CoverTab[20002]"
	case oid.Equal(oidPublicKeyEd25519):
//line /usr/local/go/src/crypto/x509/x509.go:498
		_go_fuzz_dep_.CoverTab[20003]++
								return Ed25519
//line /usr/local/go/src/crypto/x509/x509.go:499
		// _ = "end of CoverTab[20003]"
//line /usr/local/go/src/crypto/x509/x509.go:499
	default:
//line /usr/local/go/src/crypto/x509/x509.go:499
		_go_fuzz_dep_.CoverTab[20004]++
//line /usr/local/go/src/crypto/x509/x509.go:499
		// _ = "end of CoverTab[20004]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:500
	// _ = "end of CoverTab[19998]"
//line /usr/local/go/src/crypto/x509/x509.go:500
	_go_fuzz_dep_.CoverTab[19999]++
							return UnknownPublicKeyAlgorithm
//line /usr/local/go/src/crypto/x509/x509.go:501
	// _ = "end of CoverTab[19999]"
}

// RFC 5480, 2.1.1.1. Named Curve
//line /usr/local/go/src/crypto/x509/x509.go:504
//
//line /usr/local/go/src/crypto/x509/x509.go:504
//	secp224r1 OBJECT IDENTIFIER ::= {
//line /usr/local/go/src/crypto/x509/x509.go:504
//	  iso(1) identified-organization(3) certicom(132) curve(0) 33 }
//line /usr/local/go/src/crypto/x509/x509.go:504
//
//line /usr/local/go/src/crypto/x509/x509.go:504
//	secp256r1 OBJECT IDENTIFIER ::= {
//line /usr/local/go/src/crypto/x509/x509.go:504
//	  iso(1) member-body(2) us(840) ansi-X9-62(10045) curves(3)
//line /usr/local/go/src/crypto/x509/x509.go:504
//	  prime(1) 7 }
//line /usr/local/go/src/crypto/x509/x509.go:504
//
//line /usr/local/go/src/crypto/x509/x509.go:504
//	secp384r1 OBJECT IDENTIFIER ::= {
//line /usr/local/go/src/crypto/x509/x509.go:504
//	  iso(1) identified-organization(3) certicom(132) curve(0) 34 }
//line /usr/local/go/src/crypto/x509/x509.go:504
//
//line /usr/local/go/src/crypto/x509/x509.go:504
//	secp521r1 OBJECT IDENTIFIER ::= {
//line /usr/local/go/src/crypto/x509/x509.go:504
//	  iso(1) identified-organization(3) certicom(132) curve(0) 35 }
//line /usr/local/go/src/crypto/x509/x509.go:504
//
//line /usr/local/go/src/crypto/x509/x509.go:504
// NB: secp256r1 is equivalent to prime256v1
//line /usr/local/go/src/crypto/x509/x509.go:520
var (
	oidNamedCurveP224	= asn1.ObjectIdentifier{1, 3, 132, 0, 33}
	oidNamedCurveP256	= asn1.ObjectIdentifier{1, 2, 840, 10045, 3, 1, 7}
	oidNamedCurveP384	= asn1.ObjectIdentifier{1, 3, 132, 0, 34}
	oidNamedCurveP521	= asn1.ObjectIdentifier{1, 3, 132, 0, 35}
)

func namedCurveFromOID(oid asn1.ObjectIdentifier) elliptic.Curve {
//line /usr/local/go/src/crypto/x509/x509.go:527
	_go_fuzz_dep_.CoverTab[20005]++
							switch {
	case oid.Equal(oidNamedCurveP224):
//line /usr/local/go/src/crypto/x509/x509.go:529
		_go_fuzz_dep_.CoverTab[20007]++
								return elliptic.P224()
//line /usr/local/go/src/crypto/x509/x509.go:530
		// _ = "end of CoverTab[20007]"
	case oid.Equal(oidNamedCurveP256):
//line /usr/local/go/src/crypto/x509/x509.go:531
		_go_fuzz_dep_.CoverTab[20008]++
								return elliptic.P256()
//line /usr/local/go/src/crypto/x509/x509.go:532
		// _ = "end of CoverTab[20008]"
	case oid.Equal(oidNamedCurveP384):
//line /usr/local/go/src/crypto/x509/x509.go:533
		_go_fuzz_dep_.CoverTab[20009]++
								return elliptic.P384()
//line /usr/local/go/src/crypto/x509/x509.go:534
		// _ = "end of CoverTab[20009]"
	case oid.Equal(oidNamedCurveP521):
//line /usr/local/go/src/crypto/x509/x509.go:535
		_go_fuzz_dep_.CoverTab[20010]++
								return elliptic.P521()
//line /usr/local/go/src/crypto/x509/x509.go:536
		// _ = "end of CoverTab[20010]"
//line /usr/local/go/src/crypto/x509/x509.go:536
	default:
//line /usr/local/go/src/crypto/x509/x509.go:536
		_go_fuzz_dep_.CoverTab[20011]++
//line /usr/local/go/src/crypto/x509/x509.go:536
		// _ = "end of CoverTab[20011]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:537
	// _ = "end of CoverTab[20005]"
//line /usr/local/go/src/crypto/x509/x509.go:537
	_go_fuzz_dep_.CoverTab[20006]++
							return nil
//line /usr/local/go/src/crypto/x509/x509.go:538
	// _ = "end of CoverTab[20006]"
}

func oidFromNamedCurve(curve elliptic.Curve) (asn1.ObjectIdentifier, bool) {
//line /usr/local/go/src/crypto/x509/x509.go:541
	_go_fuzz_dep_.CoverTab[20012]++
							switch curve {
	case elliptic.P224():
//line /usr/local/go/src/crypto/x509/x509.go:543
		_go_fuzz_dep_.CoverTab[20014]++
								return oidNamedCurveP224, true
//line /usr/local/go/src/crypto/x509/x509.go:544
		// _ = "end of CoverTab[20014]"
	case elliptic.P256():
//line /usr/local/go/src/crypto/x509/x509.go:545
		_go_fuzz_dep_.CoverTab[20015]++
								return oidNamedCurveP256, true
//line /usr/local/go/src/crypto/x509/x509.go:546
		// _ = "end of CoverTab[20015]"
	case elliptic.P384():
//line /usr/local/go/src/crypto/x509/x509.go:547
		_go_fuzz_dep_.CoverTab[20016]++
								return oidNamedCurveP384, true
//line /usr/local/go/src/crypto/x509/x509.go:548
		// _ = "end of CoverTab[20016]"
	case elliptic.P521():
//line /usr/local/go/src/crypto/x509/x509.go:549
		_go_fuzz_dep_.CoverTab[20017]++
								return oidNamedCurveP521, true
//line /usr/local/go/src/crypto/x509/x509.go:550
		// _ = "end of CoverTab[20017]"
//line /usr/local/go/src/crypto/x509/x509.go:550
	default:
//line /usr/local/go/src/crypto/x509/x509.go:550
		_go_fuzz_dep_.CoverTab[20018]++
//line /usr/local/go/src/crypto/x509/x509.go:550
		// _ = "end of CoverTab[20018]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:551
	// _ = "end of CoverTab[20012]"
//line /usr/local/go/src/crypto/x509/x509.go:551
	_go_fuzz_dep_.CoverTab[20013]++

							return nil, false
//line /usr/local/go/src/crypto/x509/x509.go:553
	// _ = "end of CoverTab[20013]"
}

func oidFromECDHCurve(curve ecdh.Curve) (asn1.ObjectIdentifier, bool) {
//line /usr/local/go/src/crypto/x509/x509.go:556
	_go_fuzz_dep_.CoverTab[20019]++
							switch curve {
	case ecdh.X25519():
//line /usr/local/go/src/crypto/x509/x509.go:558
		_go_fuzz_dep_.CoverTab[20021]++
								return oidPublicKeyX25519, true
//line /usr/local/go/src/crypto/x509/x509.go:559
		// _ = "end of CoverTab[20021]"
	case ecdh.P256():
//line /usr/local/go/src/crypto/x509/x509.go:560
		_go_fuzz_dep_.CoverTab[20022]++
								return oidNamedCurveP256, true
//line /usr/local/go/src/crypto/x509/x509.go:561
		// _ = "end of CoverTab[20022]"
	case ecdh.P384():
//line /usr/local/go/src/crypto/x509/x509.go:562
		_go_fuzz_dep_.CoverTab[20023]++
								return oidNamedCurveP384, true
//line /usr/local/go/src/crypto/x509/x509.go:563
		// _ = "end of CoverTab[20023]"
	case ecdh.P521():
//line /usr/local/go/src/crypto/x509/x509.go:564
		_go_fuzz_dep_.CoverTab[20024]++
								return oidNamedCurveP521, true
//line /usr/local/go/src/crypto/x509/x509.go:565
		// _ = "end of CoverTab[20024]"
//line /usr/local/go/src/crypto/x509/x509.go:565
	default:
//line /usr/local/go/src/crypto/x509/x509.go:565
		_go_fuzz_dep_.CoverTab[20025]++
//line /usr/local/go/src/crypto/x509/x509.go:565
		// _ = "end of CoverTab[20025]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:566
	// _ = "end of CoverTab[20019]"
//line /usr/local/go/src/crypto/x509/x509.go:566
	_go_fuzz_dep_.CoverTab[20020]++

							return nil, false
//line /usr/local/go/src/crypto/x509/x509.go:568
	// _ = "end of CoverTab[20020]"
}

// KeyUsage represents the set of actions that are valid for a given key. It's
//line /usr/local/go/src/crypto/x509/x509.go:571
// a bitmap of the KeyUsage* constants.
//line /usr/local/go/src/crypto/x509/x509.go:573
type KeyUsage int

const (
	KeyUsageDigitalSignature	KeyUsage	= 1 << iota
	KeyUsageContentCommitment
	KeyUsageKeyEncipherment
	KeyUsageDataEncipherment
	KeyUsageKeyAgreement
	KeyUsageCertSign
	KeyUsageCRLSign
	KeyUsageEncipherOnly
	KeyUsageDecipherOnly
)

// RFC 5280, 4.2.1.12  Extended Key Usage
//line /usr/local/go/src/crypto/x509/x509.go:587
//
//line /usr/local/go/src/crypto/x509/x509.go:587
//	anyExtendedKeyUsage OBJECT IDENTIFIER ::= { id-ce-extKeyUsage 0 }
//line /usr/local/go/src/crypto/x509/x509.go:587
//
//line /usr/local/go/src/crypto/x509/x509.go:587
//	id-kp OBJECT IDENTIFIER ::= { id-pkix 3 }
//line /usr/local/go/src/crypto/x509/x509.go:587
//
//line /usr/local/go/src/crypto/x509/x509.go:587
//	id-kp-serverAuth             OBJECT IDENTIFIER ::= { id-kp 1 }
//line /usr/local/go/src/crypto/x509/x509.go:587
//	id-kp-clientAuth             OBJECT IDENTIFIER ::= { id-kp 2 }
//line /usr/local/go/src/crypto/x509/x509.go:587
//	id-kp-codeSigning            OBJECT IDENTIFIER ::= { id-kp 3 }
//line /usr/local/go/src/crypto/x509/x509.go:587
//	id-kp-emailProtection        OBJECT IDENTIFIER ::= { id-kp 4 }
//line /usr/local/go/src/crypto/x509/x509.go:587
//	id-kp-timeStamping           OBJECT IDENTIFIER ::= { id-kp 8 }
//line /usr/local/go/src/crypto/x509/x509.go:587
//	id-kp-OCSPSigning            OBJECT IDENTIFIER ::= { id-kp 9 }
//line /usr/local/go/src/crypto/x509/x509.go:599
var (
	oidExtKeyUsageAny				= asn1.ObjectIdentifier{2, 5, 29, 37, 0}
	oidExtKeyUsageServerAuth			= asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 3, 1}
	oidExtKeyUsageClientAuth			= asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 3, 2}
	oidExtKeyUsageCodeSigning			= asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 3, 3}
	oidExtKeyUsageEmailProtection			= asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 3, 4}
	oidExtKeyUsageIPSECEndSystem			= asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 3, 5}
	oidExtKeyUsageIPSECTunnel			= asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 3, 6}
	oidExtKeyUsageIPSECUser				= asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 3, 7}
	oidExtKeyUsageTimeStamping			= asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 3, 8}
	oidExtKeyUsageOCSPSigning			= asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 3, 9}
	oidExtKeyUsageMicrosoftServerGatedCrypto	= asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 311, 10, 3, 3}
	oidExtKeyUsageNetscapeServerGatedCrypto		= asn1.ObjectIdentifier{2, 16, 840, 1, 113730, 4, 1}
	oidExtKeyUsageMicrosoftCommercialCodeSigning	= asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 311, 2, 1, 22}
	oidExtKeyUsageMicrosoftKernelCodeSigning	= asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 311, 61, 1, 1}
)

// ExtKeyUsage represents an extended set of actions that are valid for a given key.
//line /usr/local/go/src/crypto/x509/x509.go:616
// Each of the ExtKeyUsage* constants define a unique action.
//line /usr/local/go/src/crypto/x509/x509.go:618
type ExtKeyUsage int

const (
	ExtKeyUsageAny	ExtKeyUsage	= iota
	ExtKeyUsageServerAuth
	ExtKeyUsageClientAuth
	ExtKeyUsageCodeSigning
	ExtKeyUsageEmailProtection
	ExtKeyUsageIPSECEndSystem
	ExtKeyUsageIPSECTunnel
	ExtKeyUsageIPSECUser
	ExtKeyUsageTimeStamping
	ExtKeyUsageOCSPSigning
	ExtKeyUsageMicrosoftServerGatedCrypto
	ExtKeyUsageNetscapeServerGatedCrypto
	ExtKeyUsageMicrosoftCommercialCodeSigning
	ExtKeyUsageMicrosoftKernelCodeSigning
)

// extKeyUsageOIDs contains the mapping between an ExtKeyUsage and its OID.
var extKeyUsageOIDs = []struct {
	extKeyUsage	ExtKeyUsage
	oid		asn1.ObjectIdentifier
}{
	{ExtKeyUsageAny, oidExtKeyUsageAny},
	{ExtKeyUsageServerAuth, oidExtKeyUsageServerAuth},
	{ExtKeyUsageClientAuth, oidExtKeyUsageClientAuth},
	{ExtKeyUsageCodeSigning, oidExtKeyUsageCodeSigning},
	{ExtKeyUsageEmailProtection, oidExtKeyUsageEmailProtection},
	{ExtKeyUsageIPSECEndSystem, oidExtKeyUsageIPSECEndSystem},
	{ExtKeyUsageIPSECTunnel, oidExtKeyUsageIPSECTunnel},
	{ExtKeyUsageIPSECUser, oidExtKeyUsageIPSECUser},
	{ExtKeyUsageTimeStamping, oidExtKeyUsageTimeStamping},
	{ExtKeyUsageOCSPSigning, oidExtKeyUsageOCSPSigning},
	{ExtKeyUsageMicrosoftServerGatedCrypto, oidExtKeyUsageMicrosoftServerGatedCrypto},
	{ExtKeyUsageNetscapeServerGatedCrypto, oidExtKeyUsageNetscapeServerGatedCrypto},
	{ExtKeyUsageMicrosoftCommercialCodeSigning, oidExtKeyUsageMicrosoftCommercialCodeSigning},
	{ExtKeyUsageMicrosoftKernelCodeSigning, oidExtKeyUsageMicrosoftKernelCodeSigning},
}

func extKeyUsageFromOID(oid asn1.ObjectIdentifier) (eku ExtKeyUsage, ok bool) {
//line /usr/local/go/src/crypto/x509/x509.go:658
	_go_fuzz_dep_.CoverTab[20026]++
							for _, pair := range extKeyUsageOIDs {
//line /usr/local/go/src/crypto/x509/x509.go:659
		_go_fuzz_dep_.CoverTab[20028]++
								if oid.Equal(pair.oid) {
//line /usr/local/go/src/crypto/x509/x509.go:660
			_go_fuzz_dep_.CoverTab[20029]++
									return pair.extKeyUsage, true
//line /usr/local/go/src/crypto/x509/x509.go:661
			// _ = "end of CoverTab[20029]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:662
			_go_fuzz_dep_.CoverTab[20030]++
//line /usr/local/go/src/crypto/x509/x509.go:662
			// _ = "end of CoverTab[20030]"
//line /usr/local/go/src/crypto/x509/x509.go:662
		}
//line /usr/local/go/src/crypto/x509/x509.go:662
		// _ = "end of CoverTab[20028]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:663
	// _ = "end of CoverTab[20026]"
//line /usr/local/go/src/crypto/x509/x509.go:663
	_go_fuzz_dep_.CoverTab[20027]++
							return
//line /usr/local/go/src/crypto/x509/x509.go:664
	// _ = "end of CoverTab[20027]"
}

func oidFromExtKeyUsage(eku ExtKeyUsage) (oid asn1.ObjectIdentifier, ok bool) {
//line /usr/local/go/src/crypto/x509/x509.go:667
	_go_fuzz_dep_.CoverTab[20031]++
							for _, pair := range extKeyUsageOIDs {
//line /usr/local/go/src/crypto/x509/x509.go:668
		_go_fuzz_dep_.CoverTab[20033]++
								if eku == pair.extKeyUsage {
//line /usr/local/go/src/crypto/x509/x509.go:669
			_go_fuzz_dep_.CoverTab[20034]++
									return pair.oid, true
//line /usr/local/go/src/crypto/x509/x509.go:670
			// _ = "end of CoverTab[20034]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:671
			_go_fuzz_dep_.CoverTab[20035]++
//line /usr/local/go/src/crypto/x509/x509.go:671
			// _ = "end of CoverTab[20035]"
//line /usr/local/go/src/crypto/x509/x509.go:671
		}
//line /usr/local/go/src/crypto/x509/x509.go:671
		// _ = "end of CoverTab[20033]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:672
	// _ = "end of CoverTab[20031]"
//line /usr/local/go/src/crypto/x509/x509.go:672
	_go_fuzz_dep_.CoverTab[20032]++
							return
//line /usr/local/go/src/crypto/x509/x509.go:673
	// _ = "end of CoverTab[20032]"
}

// A Certificate represents an X.509 certificate.
type Certificate struct {
	Raw			[]byte	// Complete ASN.1 DER content (certificate, signature algorithm and signature).
	RawTBSCertificate	[]byte	// Certificate part of raw ASN.1 DER content.
	RawSubjectPublicKeyInfo	[]byte	// DER encoded SubjectPublicKeyInfo.
	RawSubject		[]byte	// DER encoded Subject
	RawIssuer		[]byte	// DER encoded Issuer

	Signature		[]byte
	SignatureAlgorithm	SignatureAlgorithm

	PublicKeyAlgorithm	PublicKeyAlgorithm
	PublicKey		any

	Version			int
	SerialNumber		*big.Int
	Issuer			pkix.Name
	Subject			pkix.Name
	NotBefore, NotAfter	time.Time	// Validity bounds.
	KeyUsage		KeyUsage

	// Extensions contains raw X.509 extensions. When parsing certificates,
	// this can be used to extract non-critical extensions that are not
	// parsed by this package. When marshaling certificates, the Extensions
	// field is ignored, see ExtraExtensions.
	Extensions	[]pkix.Extension

	// ExtraExtensions contains extensions to be copied, raw, into any
	// marshaled certificates. Values override any extensions that would
	// otherwise be produced based on the other fields. The ExtraExtensions
	// field is not populated when parsing certificates, see Extensions.
	ExtraExtensions	[]pkix.Extension

	// UnhandledCriticalExtensions contains a list of extension IDs that
	// were not (fully) processed when parsing. Verify will fail if this
	// slice is non-empty, unless verification is delegated to an OS
	// library which understands all the critical extensions.
	//
	// Users can access these extensions using Extensions and can remove
	// elements from this slice if they believe that they have been
	// handled.
	UnhandledCriticalExtensions	[]asn1.ObjectIdentifier

	ExtKeyUsage		[]ExtKeyUsage		// Sequence of extended key usages.
	UnknownExtKeyUsage	[]asn1.ObjectIdentifier	// Encountered extended key usages unknown to this package.

	// BasicConstraintsValid indicates whether IsCA, MaxPathLen,
	// and MaxPathLenZero are valid.
	BasicConstraintsValid	bool
	IsCA			bool

	// MaxPathLen and MaxPathLenZero indicate the presence and
	// value of the BasicConstraints' "pathLenConstraint".
	//
	// When parsing a certificate, a positive non-zero MaxPathLen
	// means that the field was specified, -1 means it was unset,
	// and MaxPathLenZero being true mean that the field was
	// explicitly set to zero. The case of MaxPathLen==0 with MaxPathLenZero==false
	// should be treated equivalent to -1 (unset).
	//
	// When generating a certificate, an unset pathLenConstraint
	// can be requested with either MaxPathLen == -1 or using the
	// zero value for both MaxPathLen and MaxPathLenZero.
	MaxPathLen	int
	// MaxPathLenZero indicates that BasicConstraintsValid==true
	// and MaxPathLen==0 should be interpreted as an actual
	// maximum path length of zero. Otherwise, that combination is
	// interpreted as MaxPathLen not being set.
	MaxPathLenZero	bool

	SubjectKeyId	[]byte
	AuthorityKeyId	[]byte

	// RFC 5280, 4.2.2.1 (Authority Information Access)
	OCSPServer		[]string
	IssuingCertificateURL	[]string

	// Subject Alternate Name values. (Note that these values may not be valid
	// if invalid values were contained within a parsed certificate. For
	// example, an element of DNSNames may not be a valid DNS domain name.)
	DNSNames	[]string
	EmailAddresses	[]string
	IPAddresses	[]net.IP
	URIs		[]*url.URL

	// Name constraints
	PermittedDNSDomainsCritical	bool	// if true then the name constraints are marked critical.
	PermittedDNSDomains		[]string
	ExcludedDNSDomains		[]string
	PermittedIPRanges		[]*net.IPNet
	ExcludedIPRanges		[]*net.IPNet
	PermittedEmailAddresses		[]string
	ExcludedEmailAddresses		[]string
	PermittedURIDomains		[]string
	ExcludedURIDomains		[]string

	// CRL Distribution Points
	CRLDistributionPoints	[]string

	PolicyIdentifiers	[]asn1.ObjectIdentifier
}

// ErrUnsupportedAlgorithm results from attempting to perform an operation that
//line /usr/local/go/src/crypto/x509/x509.go:778
// involves algorithms that are not currently implemented.
//line /usr/local/go/src/crypto/x509/x509.go:780
var ErrUnsupportedAlgorithm = errors.New("x509: cannot verify signature: algorithm unimplemented")

// An InsecureAlgorithmError indicates that the SignatureAlgorithm used to
//line /usr/local/go/src/crypto/x509/x509.go:782
// generate the signature is not secure, and the signature has been rejected.
//line /usr/local/go/src/crypto/x509/x509.go:782
//
//line /usr/local/go/src/crypto/x509/x509.go:782
// To temporarily restore support for SHA-1 signatures, include the value
//line /usr/local/go/src/crypto/x509/x509.go:782
// "x509sha1=1" in the GODEBUG environment variable. Note that this option will
//line /usr/local/go/src/crypto/x509/x509.go:782
// be removed in a future release.
//line /usr/local/go/src/crypto/x509/x509.go:788
type InsecureAlgorithmError SignatureAlgorithm

func (e InsecureAlgorithmError) Error() string {
//line /usr/local/go/src/crypto/x509/x509.go:790
	_go_fuzz_dep_.CoverTab[20036]++
							var override string
							if SignatureAlgorithm(e) == SHA1WithRSA || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:792
		_go_fuzz_dep_.CoverTab[20038]++
//line /usr/local/go/src/crypto/x509/x509.go:792
		return SignatureAlgorithm(e) == ECDSAWithSHA1
//line /usr/local/go/src/crypto/x509/x509.go:792
		// _ = "end of CoverTab[20038]"
//line /usr/local/go/src/crypto/x509/x509.go:792
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:792
		_go_fuzz_dep_.CoverTab[20039]++
								override = " (temporarily override with GODEBUG=x509sha1=1)"
//line /usr/local/go/src/crypto/x509/x509.go:793
		// _ = "end of CoverTab[20039]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:794
		_go_fuzz_dep_.CoverTab[20040]++
//line /usr/local/go/src/crypto/x509/x509.go:794
		// _ = "end of CoverTab[20040]"
//line /usr/local/go/src/crypto/x509/x509.go:794
	}
//line /usr/local/go/src/crypto/x509/x509.go:794
	// _ = "end of CoverTab[20036]"
//line /usr/local/go/src/crypto/x509/x509.go:794
	_go_fuzz_dep_.CoverTab[20037]++
							return fmt.Sprintf("x509: cannot verify signature: insecure algorithm %v", SignatureAlgorithm(e)) + override
//line /usr/local/go/src/crypto/x509/x509.go:795
	// _ = "end of CoverTab[20037]"
}

// ConstraintViolationError results when a requested usage is not permitted by
//line /usr/local/go/src/crypto/x509/x509.go:798
// a certificate. For example: checking a signature when the public key isn't a
//line /usr/local/go/src/crypto/x509/x509.go:798
// certificate signing key.
//line /usr/local/go/src/crypto/x509/x509.go:801
type ConstraintViolationError struct{}

func (ConstraintViolationError) Error() string {
//line /usr/local/go/src/crypto/x509/x509.go:803
	_go_fuzz_dep_.CoverTab[20041]++
							return "x509: invalid signature: parent certificate cannot sign this kind of certificate"
//line /usr/local/go/src/crypto/x509/x509.go:804
	// _ = "end of CoverTab[20041]"
}

func (c *Certificate) Equal(other *Certificate) bool {
//line /usr/local/go/src/crypto/x509/x509.go:807
	_go_fuzz_dep_.CoverTab[20042]++
							if c == nil || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:808
		_go_fuzz_dep_.CoverTab[20044]++
//line /usr/local/go/src/crypto/x509/x509.go:808
		return other == nil
//line /usr/local/go/src/crypto/x509/x509.go:808
		// _ = "end of CoverTab[20044]"
//line /usr/local/go/src/crypto/x509/x509.go:808
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:808
		_go_fuzz_dep_.CoverTab[20045]++
								return c == other
//line /usr/local/go/src/crypto/x509/x509.go:809
		// _ = "end of CoverTab[20045]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:810
		_go_fuzz_dep_.CoverTab[20046]++
//line /usr/local/go/src/crypto/x509/x509.go:810
		// _ = "end of CoverTab[20046]"
//line /usr/local/go/src/crypto/x509/x509.go:810
	}
//line /usr/local/go/src/crypto/x509/x509.go:810
	// _ = "end of CoverTab[20042]"
//line /usr/local/go/src/crypto/x509/x509.go:810
	_go_fuzz_dep_.CoverTab[20043]++
							return bytes.Equal(c.Raw, other.Raw)
//line /usr/local/go/src/crypto/x509/x509.go:811
	// _ = "end of CoverTab[20043]"
}

func (c *Certificate) hasSANExtension() bool {
//line /usr/local/go/src/crypto/x509/x509.go:814
	_go_fuzz_dep_.CoverTab[20047]++
							return oidInExtensions(oidExtensionSubjectAltName, c.Extensions)
//line /usr/local/go/src/crypto/x509/x509.go:815
	// _ = "end of CoverTab[20047]"
}

// CheckSignatureFrom verifies that the signature on c is a valid signature from parent.
//line /usr/local/go/src/crypto/x509/x509.go:818
//
//line /usr/local/go/src/crypto/x509/x509.go:818
// This is a low-level API that performs very limited checks, and not a full
//line /usr/local/go/src/crypto/x509/x509.go:818
// path verifier. Most users should use [Certificate.Verify] instead.
//line /usr/local/go/src/crypto/x509/x509.go:822
func (c *Certificate) CheckSignatureFrom(parent *Certificate) error {
//line /usr/local/go/src/crypto/x509/x509.go:822
	_go_fuzz_dep_.CoverTab[20048]++

//line /usr/local/go/src/crypto/x509/x509.go:828
	if parent.Version == 3 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:828
		_go_fuzz_dep_.CoverTab[20052]++
//line /usr/local/go/src/crypto/x509/x509.go:828
		return !parent.BasicConstraintsValid
//line /usr/local/go/src/crypto/x509/x509.go:828
		// _ = "end of CoverTab[20052]"
//line /usr/local/go/src/crypto/x509/x509.go:828
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:828
		_go_fuzz_dep_.CoverTab[20053]++
//line /usr/local/go/src/crypto/x509/x509.go:828
		return parent.BasicConstraintsValid && func() bool {
									_go_fuzz_dep_.CoverTab[20054]++
//line /usr/local/go/src/crypto/x509/x509.go:829
			return !parent.IsCA
//line /usr/local/go/src/crypto/x509/x509.go:829
			// _ = "end of CoverTab[20054]"
//line /usr/local/go/src/crypto/x509/x509.go:829
		}()
//line /usr/local/go/src/crypto/x509/x509.go:829
		// _ = "end of CoverTab[20053]"
//line /usr/local/go/src/crypto/x509/x509.go:829
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:829
		_go_fuzz_dep_.CoverTab[20055]++
								return ConstraintViolationError{}
//line /usr/local/go/src/crypto/x509/x509.go:830
		// _ = "end of CoverTab[20055]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:831
		_go_fuzz_dep_.CoverTab[20056]++
//line /usr/local/go/src/crypto/x509/x509.go:831
		// _ = "end of CoverTab[20056]"
//line /usr/local/go/src/crypto/x509/x509.go:831
	}
//line /usr/local/go/src/crypto/x509/x509.go:831
	// _ = "end of CoverTab[20048]"
//line /usr/local/go/src/crypto/x509/x509.go:831
	_go_fuzz_dep_.CoverTab[20049]++

							if parent.KeyUsage != 0 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:833
		_go_fuzz_dep_.CoverTab[20057]++
//line /usr/local/go/src/crypto/x509/x509.go:833
		return parent.KeyUsage&KeyUsageCertSign == 0
//line /usr/local/go/src/crypto/x509/x509.go:833
		// _ = "end of CoverTab[20057]"
//line /usr/local/go/src/crypto/x509/x509.go:833
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:833
		_go_fuzz_dep_.CoverTab[20058]++
								return ConstraintViolationError{}
//line /usr/local/go/src/crypto/x509/x509.go:834
		// _ = "end of CoverTab[20058]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:835
		_go_fuzz_dep_.CoverTab[20059]++
//line /usr/local/go/src/crypto/x509/x509.go:835
		// _ = "end of CoverTab[20059]"
//line /usr/local/go/src/crypto/x509/x509.go:835
	}
//line /usr/local/go/src/crypto/x509/x509.go:835
	// _ = "end of CoverTab[20049]"
//line /usr/local/go/src/crypto/x509/x509.go:835
	_go_fuzz_dep_.CoverTab[20050]++

							if parent.PublicKeyAlgorithm == UnknownPublicKeyAlgorithm {
//line /usr/local/go/src/crypto/x509/x509.go:837
		_go_fuzz_dep_.CoverTab[20060]++
								return ErrUnsupportedAlgorithm
//line /usr/local/go/src/crypto/x509/x509.go:838
		// _ = "end of CoverTab[20060]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:839
		_go_fuzz_dep_.CoverTab[20061]++
//line /usr/local/go/src/crypto/x509/x509.go:839
		// _ = "end of CoverTab[20061]"
//line /usr/local/go/src/crypto/x509/x509.go:839
	}
//line /usr/local/go/src/crypto/x509/x509.go:839
	// _ = "end of CoverTab[20050]"
//line /usr/local/go/src/crypto/x509/x509.go:839
	_go_fuzz_dep_.CoverTab[20051]++

							return checkSignature(c.SignatureAlgorithm, c.RawTBSCertificate, c.Signature, parent.PublicKey, false)
//line /usr/local/go/src/crypto/x509/x509.go:841
	// _ = "end of CoverTab[20051]"
}

// CheckSignature verifies that signature is a valid signature over signed from
//line /usr/local/go/src/crypto/x509/x509.go:844
// c's public key.
//line /usr/local/go/src/crypto/x509/x509.go:844
//
//line /usr/local/go/src/crypto/x509/x509.go:844
// This is a low-level API that performs no validity checks on the certificate.
//line /usr/local/go/src/crypto/x509/x509.go:844
//
//line /usr/local/go/src/crypto/x509/x509.go:844
// [MD5WithRSA] signatures are rejected, while [SHA1WithRSA] and [ECDSAWithSHA1]
//line /usr/local/go/src/crypto/x509/x509.go:844
// signatures are currently accepted.
//line /usr/local/go/src/crypto/x509/x509.go:851
func (c *Certificate) CheckSignature(algo SignatureAlgorithm, signed, signature []byte) error {
//line /usr/local/go/src/crypto/x509/x509.go:851
	_go_fuzz_dep_.CoverTab[20062]++
							return checkSignature(algo, signed, signature, c.PublicKey, true)
//line /usr/local/go/src/crypto/x509/x509.go:852
	// _ = "end of CoverTab[20062]"
}

func (c *Certificate) hasNameConstraints() bool {
//line /usr/local/go/src/crypto/x509/x509.go:855
	_go_fuzz_dep_.CoverTab[20063]++
							return oidInExtensions(oidExtensionNameConstraints, c.Extensions)
//line /usr/local/go/src/crypto/x509/x509.go:856
	// _ = "end of CoverTab[20063]"
}

func (c *Certificate) getSANExtension() []byte {
//line /usr/local/go/src/crypto/x509/x509.go:859
	_go_fuzz_dep_.CoverTab[20064]++
							for _, e := range c.Extensions {
//line /usr/local/go/src/crypto/x509/x509.go:860
		_go_fuzz_dep_.CoverTab[20066]++
								if e.Id.Equal(oidExtensionSubjectAltName) {
//line /usr/local/go/src/crypto/x509/x509.go:861
			_go_fuzz_dep_.CoverTab[20067]++
									return e.Value
//line /usr/local/go/src/crypto/x509/x509.go:862
			// _ = "end of CoverTab[20067]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:863
			_go_fuzz_dep_.CoverTab[20068]++
//line /usr/local/go/src/crypto/x509/x509.go:863
			// _ = "end of CoverTab[20068]"
//line /usr/local/go/src/crypto/x509/x509.go:863
		}
//line /usr/local/go/src/crypto/x509/x509.go:863
		// _ = "end of CoverTab[20066]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:864
	// _ = "end of CoverTab[20064]"
//line /usr/local/go/src/crypto/x509/x509.go:864
	_go_fuzz_dep_.CoverTab[20065]++
							return nil
//line /usr/local/go/src/crypto/x509/x509.go:865
	// _ = "end of CoverTab[20065]"
}

func signaturePublicKeyAlgoMismatchError(expectedPubKeyAlgo PublicKeyAlgorithm, pubKey any) error {
//line /usr/local/go/src/crypto/x509/x509.go:868
	_go_fuzz_dep_.CoverTab[20069]++
							return fmt.Errorf("x509: signature algorithm specifies an %s public key, but have public key of type %T", expectedPubKeyAlgo.String(), pubKey)
//line /usr/local/go/src/crypto/x509/x509.go:869
	// _ = "end of CoverTab[20069]"
}

var x509sha1 = godebug.New("x509sha1")

// checkSignature verifies that signature is a valid signature over signed from
//line /usr/local/go/src/crypto/x509/x509.go:874
// a crypto.PublicKey.
//line /usr/local/go/src/crypto/x509/x509.go:876
func checkSignature(algo SignatureAlgorithm, signed, signature []byte, publicKey crypto.PublicKey, allowSHA1 bool) (err error) {
//line /usr/local/go/src/crypto/x509/x509.go:876
	_go_fuzz_dep_.CoverTab[20070]++
							var hashType crypto.Hash
							var pubKeyAlgo PublicKeyAlgorithm

							for _, details := range signatureAlgorithmDetails {
//line /usr/local/go/src/crypto/x509/x509.go:880
		_go_fuzz_dep_.CoverTab[20074]++
								if details.algo == algo {
//line /usr/local/go/src/crypto/x509/x509.go:881
			_go_fuzz_dep_.CoverTab[20075]++
									hashType = details.hash
									pubKeyAlgo = details.pubKeyAlgo
//line /usr/local/go/src/crypto/x509/x509.go:883
			// _ = "end of CoverTab[20075]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:884
			_go_fuzz_dep_.CoverTab[20076]++
//line /usr/local/go/src/crypto/x509/x509.go:884
			// _ = "end of CoverTab[20076]"
//line /usr/local/go/src/crypto/x509/x509.go:884
		}
//line /usr/local/go/src/crypto/x509/x509.go:884
		// _ = "end of CoverTab[20074]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:885
	// _ = "end of CoverTab[20070]"
//line /usr/local/go/src/crypto/x509/x509.go:885
	_go_fuzz_dep_.CoverTab[20071]++

							switch hashType {
	case crypto.Hash(0):
//line /usr/local/go/src/crypto/x509/x509.go:888
		_go_fuzz_dep_.CoverTab[20077]++
								if pubKeyAlgo != Ed25519 {
//line /usr/local/go/src/crypto/x509/x509.go:889
			_go_fuzz_dep_.CoverTab[20083]++
									return ErrUnsupportedAlgorithm
//line /usr/local/go/src/crypto/x509/x509.go:890
			// _ = "end of CoverTab[20083]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:891
			_go_fuzz_dep_.CoverTab[20084]++
//line /usr/local/go/src/crypto/x509/x509.go:891
			// _ = "end of CoverTab[20084]"
//line /usr/local/go/src/crypto/x509/x509.go:891
		}
//line /usr/local/go/src/crypto/x509/x509.go:891
		// _ = "end of CoverTab[20077]"
	case crypto.MD5:
//line /usr/local/go/src/crypto/x509/x509.go:892
		_go_fuzz_dep_.CoverTab[20078]++
								return InsecureAlgorithmError(algo)
//line /usr/local/go/src/crypto/x509/x509.go:893
		// _ = "end of CoverTab[20078]"
	case crypto.SHA1:
//line /usr/local/go/src/crypto/x509/x509.go:894
		_go_fuzz_dep_.CoverTab[20079]++

								if !allowSHA1 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:896
			_go_fuzz_dep_.CoverTab[20085]++
//line /usr/local/go/src/crypto/x509/x509.go:896
			return x509sha1.Value() != "1"
//line /usr/local/go/src/crypto/x509/x509.go:896
			// _ = "end of CoverTab[20085]"
//line /usr/local/go/src/crypto/x509/x509.go:896
		}() {
//line /usr/local/go/src/crypto/x509/x509.go:896
			_go_fuzz_dep_.CoverTab[20086]++
									return InsecureAlgorithmError(algo)
//line /usr/local/go/src/crypto/x509/x509.go:897
			// _ = "end of CoverTab[20086]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:898
			_go_fuzz_dep_.CoverTab[20087]++
//line /usr/local/go/src/crypto/x509/x509.go:898
			// _ = "end of CoverTab[20087]"
//line /usr/local/go/src/crypto/x509/x509.go:898
		}
//line /usr/local/go/src/crypto/x509/x509.go:898
		// _ = "end of CoverTab[20079]"
//line /usr/local/go/src/crypto/x509/x509.go:898
		_go_fuzz_dep_.CoverTab[20080]++
								fallthrough
//line /usr/local/go/src/crypto/x509/x509.go:899
		// _ = "end of CoverTab[20080]"
	default:
//line /usr/local/go/src/crypto/x509/x509.go:900
		_go_fuzz_dep_.CoverTab[20081]++
								if !hashType.Available() {
//line /usr/local/go/src/crypto/x509/x509.go:901
			_go_fuzz_dep_.CoverTab[20088]++
									return ErrUnsupportedAlgorithm
//line /usr/local/go/src/crypto/x509/x509.go:902
			// _ = "end of CoverTab[20088]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:903
			_go_fuzz_dep_.CoverTab[20089]++
//line /usr/local/go/src/crypto/x509/x509.go:903
			// _ = "end of CoverTab[20089]"
//line /usr/local/go/src/crypto/x509/x509.go:903
		}
//line /usr/local/go/src/crypto/x509/x509.go:903
		// _ = "end of CoverTab[20081]"
//line /usr/local/go/src/crypto/x509/x509.go:903
		_go_fuzz_dep_.CoverTab[20082]++
								h := hashType.New()
								h.Write(signed)
								signed = h.Sum(nil)
//line /usr/local/go/src/crypto/x509/x509.go:906
		// _ = "end of CoverTab[20082]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:907
	// _ = "end of CoverTab[20071]"
//line /usr/local/go/src/crypto/x509/x509.go:907
	_go_fuzz_dep_.CoverTab[20072]++

							switch pub := publicKey.(type) {
	case *rsa.PublicKey:
//line /usr/local/go/src/crypto/x509/x509.go:910
		_go_fuzz_dep_.CoverTab[20090]++
								if pubKeyAlgo != RSA {
//line /usr/local/go/src/crypto/x509/x509.go:911
			_go_fuzz_dep_.CoverTab[20098]++
									return signaturePublicKeyAlgoMismatchError(pubKeyAlgo, pub)
//line /usr/local/go/src/crypto/x509/x509.go:912
			// _ = "end of CoverTab[20098]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:913
			_go_fuzz_dep_.CoverTab[20099]++
//line /usr/local/go/src/crypto/x509/x509.go:913
			// _ = "end of CoverTab[20099]"
//line /usr/local/go/src/crypto/x509/x509.go:913
		}
//line /usr/local/go/src/crypto/x509/x509.go:913
		// _ = "end of CoverTab[20090]"
//line /usr/local/go/src/crypto/x509/x509.go:913
		_go_fuzz_dep_.CoverTab[20091]++
								if algo.isRSAPSS() {
//line /usr/local/go/src/crypto/x509/x509.go:914
			_go_fuzz_dep_.CoverTab[20100]++
									return rsa.VerifyPSS(pub, hashType, signed, signature, &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash})
//line /usr/local/go/src/crypto/x509/x509.go:915
			// _ = "end of CoverTab[20100]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:916
			_go_fuzz_dep_.CoverTab[20101]++
									return rsa.VerifyPKCS1v15(pub, hashType, signed, signature)
//line /usr/local/go/src/crypto/x509/x509.go:917
			// _ = "end of CoverTab[20101]"
		}
//line /usr/local/go/src/crypto/x509/x509.go:918
		// _ = "end of CoverTab[20091]"
	case *ecdsa.PublicKey:
//line /usr/local/go/src/crypto/x509/x509.go:919
		_go_fuzz_dep_.CoverTab[20092]++
								if pubKeyAlgo != ECDSA {
//line /usr/local/go/src/crypto/x509/x509.go:920
			_go_fuzz_dep_.CoverTab[20102]++
									return signaturePublicKeyAlgoMismatchError(pubKeyAlgo, pub)
//line /usr/local/go/src/crypto/x509/x509.go:921
			// _ = "end of CoverTab[20102]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:922
			_go_fuzz_dep_.CoverTab[20103]++
//line /usr/local/go/src/crypto/x509/x509.go:922
			// _ = "end of CoverTab[20103]"
//line /usr/local/go/src/crypto/x509/x509.go:922
		}
//line /usr/local/go/src/crypto/x509/x509.go:922
		// _ = "end of CoverTab[20092]"
//line /usr/local/go/src/crypto/x509/x509.go:922
		_go_fuzz_dep_.CoverTab[20093]++
								if !ecdsa.VerifyASN1(pub, signed, signature) {
//line /usr/local/go/src/crypto/x509/x509.go:923
			_go_fuzz_dep_.CoverTab[20104]++
									return errors.New("x509: ECDSA verification failure")
//line /usr/local/go/src/crypto/x509/x509.go:924
			// _ = "end of CoverTab[20104]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:925
			_go_fuzz_dep_.CoverTab[20105]++
//line /usr/local/go/src/crypto/x509/x509.go:925
			// _ = "end of CoverTab[20105]"
//line /usr/local/go/src/crypto/x509/x509.go:925
		}
//line /usr/local/go/src/crypto/x509/x509.go:925
		// _ = "end of CoverTab[20093]"
//line /usr/local/go/src/crypto/x509/x509.go:925
		_go_fuzz_dep_.CoverTab[20094]++
								return
//line /usr/local/go/src/crypto/x509/x509.go:926
		// _ = "end of CoverTab[20094]"
	case ed25519.PublicKey:
//line /usr/local/go/src/crypto/x509/x509.go:927
		_go_fuzz_dep_.CoverTab[20095]++
								if pubKeyAlgo != Ed25519 {
//line /usr/local/go/src/crypto/x509/x509.go:928
			_go_fuzz_dep_.CoverTab[20106]++
									return signaturePublicKeyAlgoMismatchError(pubKeyAlgo, pub)
//line /usr/local/go/src/crypto/x509/x509.go:929
			// _ = "end of CoverTab[20106]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:930
			_go_fuzz_dep_.CoverTab[20107]++
//line /usr/local/go/src/crypto/x509/x509.go:930
			// _ = "end of CoverTab[20107]"
//line /usr/local/go/src/crypto/x509/x509.go:930
		}
//line /usr/local/go/src/crypto/x509/x509.go:930
		// _ = "end of CoverTab[20095]"
//line /usr/local/go/src/crypto/x509/x509.go:930
		_go_fuzz_dep_.CoverTab[20096]++
								if !ed25519.Verify(pub, signed, signature) {
//line /usr/local/go/src/crypto/x509/x509.go:931
			_go_fuzz_dep_.CoverTab[20108]++
									return errors.New("x509: Ed25519 verification failure")
//line /usr/local/go/src/crypto/x509/x509.go:932
			// _ = "end of CoverTab[20108]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:933
			_go_fuzz_dep_.CoverTab[20109]++
//line /usr/local/go/src/crypto/x509/x509.go:933
			// _ = "end of CoverTab[20109]"
//line /usr/local/go/src/crypto/x509/x509.go:933
		}
//line /usr/local/go/src/crypto/x509/x509.go:933
		// _ = "end of CoverTab[20096]"
//line /usr/local/go/src/crypto/x509/x509.go:933
		_go_fuzz_dep_.CoverTab[20097]++
								return
//line /usr/local/go/src/crypto/x509/x509.go:934
		// _ = "end of CoverTab[20097]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:935
	// _ = "end of CoverTab[20072]"
//line /usr/local/go/src/crypto/x509/x509.go:935
	_go_fuzz_dep_.CoverTab[20073]++
							return ErrUnsupportedAlgorithm
//line /usr/local/go/src/crypto/x509/x509.go:936
	// _ = "end of CoverTab[20073]"
}

// CheckCRLSignature checks that the signature in crl is from c.
//line /usr/local/go/src/crypto/x509/x509.go:939
//
//line /usr/local/go/src/crypto/x509/x509.go:939
// Deprecated: Use RevocationList.CheckSignatureFrom instead.
//line /usr/local/go/src/crypto/x509/x509.go:942
func (c *Certificate) CheckCRLSignature(crl *pkix.CertificateList) error {
//line /usr/local/go/src/crypto/x509/x509.go:942
	_go_fuzz_dep_.CoverTab[20110]++
							algo := getSignatureAlgorithmFromAI(crl.SignatureAlgorithm)
							return c.CheckSignature(algo, crl.TBSCertList.Raw, crl.SignatureValue.RightAlign())
//line /usr/local/go/src/crypto/x509/x509.go:944
	// _ = "end of CoverTab[20110]"
}

type UnhandledCriticalExtension struct{}

func (h UnhandledCriticalExtension) Error() string {
//line /usr/local/go/src/crypto/x509/x509.go:949
	_go_fuzz_dep_.CoverTab[20111]++
							return "x509: unhandled critical extension"
//line /usr/local/go/src/crypto/x509/x509.go:950
	// _ = "end of CoverTab[20111]"
}

type basicConstraints struct {
	IsCA		bool	`asn1:"optional"`
	MaxPathLen	int	`asn1:"optional,default:-1"`
}

// RFC 5280 4.2.1.4
type policyInformation struct {
	Policy asn1.ObjectIdentifier
//line /usr/local/go/src/crypto/x509/x509.go:962
}

const (
	nameTypeEmail	= 1
	nameTypeDNS	= 2
	nameTypeURI	= 6
	nameTypeIP	= 7
)

// RFC 5280, 4.2.2.1
type authorityInfoAccess struct {
	Method		asn1.ObjectIdentifier
	Location	asn1.RawValue
}

// RFC 5280, 4.2.1.14
type distributionPoint struct {
	DistributionPoint	distributionPointName	`asn1:"optional,tag:0"`
	Reason			asn1.BitString		`asn1:"optional,tag:1"`
	CRLIssuer		asn1.RawValue		`asn1:"optional,tag:2"`
}

type distributionPointName struct {
	FullName	[]asn1.RawValue		`asn1:"optional,tag:0"`
	RelativeName	pkix.RDNSequence	`asn1:"optional,tag:1"`
}

func reverseBitsInAByte(in byte) byte {
//line /usr/local/go/src/crypto/x509/x509.go:989
	_go_fuzz_dep_.CoverTab[20112]++
							b1 := in>>4 | in<<4
							b2 := b1>>2&0x33 | b1<<2&0xcc
							b3 := b2>>1&0x55 | b2<<1&0xaa
							return b3
//line /usr/local/go/src/crypto/x509/x509.go:993
	// _ = "end of CoverTab[20112]"
}

// asn1BitLength returns the bit-length of bitString by considering the
//line /usr/local/go/src/crypto/x509/x509.go:996
// most-significant bit in a byte to be the "first" bit. This convention
//line /usr/local/go/src/crypto/x509/x509.go:996
// matches ASN.1, but differs from almost everything else.
//line /usr/local/go/src/crypto/x509/x509.go:999
func asn1BitLength(bitString []byte) int {
//line /usr/local/go/src/crypto/x509/x509.go:999
	_go_fuzz_dep_.CoverTab[20113]++
							bitLen := len(bitString) * 8

							for i := range bitString {
//line /usr/local/go/src/crypto/x509/x509.go:1002
		_go_fuzz_dep_.CoverTab[20115]++
								b := bitString[len(bitString)-i-1]

								for bit := uint(0); bit < 8; bit++ {
//line /usr/local/go/src/crypto/x509/x509.go:1005
			_go_fuzz_dep_.CoverTab[20116]++
									if (b>>bit)&1 == 1 {
//line /usr/local/go/src/crypto/x509/x509.go:1006
				_go_fuzz_dep_.CoverTab[20118]++
										return bitLen
//line /usr/local/go/src/crypto/x509/x509.go:1007
				// _ = "end of CoverTab[20118]"
			} else {
//line /usr/local/go/src/crypto/x509/x509.go:1008
				_go_fuzz_dep_.CoverTab[20119]++
//line /usr/local/go/src/crypto/x509/x509.go:1008
				// _ = "end of CoverTab[20119]"
//line /usr/local/go/src/crypto/x509/x509.go:1008
			}
//line /usr/local/go/src/crypto/x509/x509.go:1008
			// _ = "end of CoverTab[20116]"
//line /usr/local/go/src/crypto/x509/x509.go:1008
			_go_fuzz_dep_.CoverTab[20117]++
									bitLen--
//line /usr/local/go/src/crypto/x509/x509.go:1009
			// _ = "end of CoverTab[20117]"
		}
//line /usr/local/go/src/crypto/x509/x509.go:1010
		// _ = "end of CoverTab[20115]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:1011
	// _ = "end of CoverTab[20113]"
//line /usr/local/go/src/crypto/x509/x509.go:1011
	_go_fuzz_dep_.CoverTab[20114]++

							return 0
//line /usr/local/go/src/crypto/x509/x509.go:1013
	// _ = "end of CoverTab[20114]"
}

var (
	oidExtensionSubjectKeyId		= []int{2, 5, 29, 14}
	oidExtensionKeyUsage			= []int{2, 5, 29, 15}
	oidExtensionExtendedKeyUsage		= []int{2, 5, 29, 37}
	oidExtensionAuthorityKeyId		= []int{2, 5, 29, 35}
	oidExtensionBasicConstraints		= []int{2, 5, 29, 19}
	oidExtensionSubjectAltName		= []int{2, 5, 29, 17}
	oidExtensionCertificatePolicies		= []int{2, 5, 29, 32}
	oidExtensionNameConstraints		= []int{2, 5, 29, 30}
	oidExtensionCRLDistributionPoints	= []int{2, 5, 29, 31}
	oidExtensionAuthorityInfoAccess		= []int{1, 3, 6, 1, 5, 5, 7, 1, 1}
	oidExtensionCRLNumber			= []int{2, 5, 29, 20}
)

var (
	oidAuthorityInfoAccessOcsp	= asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 48, 1}
	oidAuthorityInfoAccessIssuers	= asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 48, 2}
)

// oidInExtensions reports whether an extension with the given oid exists in
//line /usr/local/go/src/crypto/x509/x509.go:1035
// extensions.
//line /usr/local/go/src/crypto/x509/x509.go:1037
func oidInExtensions(oid asn1.ObjectIdentifier, extensions []pkix.Extension) bool {
//line /usr/local/go/src/crypto/x509/x509.go:1037
	_go_fuzz_dep_.CoverTab[20120]++
							for _, e := range extensions {
//line /usr/local/go/src/crypto/x509/x509.go:1038
		_go_fuzz_dep_.CoverTab[20122]++
								if e.Id.Equal(oid) {
//line /usr/local/go/src/crypto/x509/x509.go:1039
			_go_fuzz_dep_.CoverTab[20123]++
									return true
//line /usr/local/go/src/crypto/x509/x509.go:1040
			// _ = "end of CoverTab[20123]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1041
			_go_fuzz_dep_.CoverTab[20124]++
//line /usr/local/go/src/crypto/x509/x509.go:1041
			// _ = "end of CoverTab[20124]"
//line /usr/local/go/src/crypto/x509/x509.go:1041
		}
//line /usr/local/go/src/crypto/x509/x509.go:1041
		// _ = "end of CoverTab[20122]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:1042
	// _ = "end of CoverTab[20120]"
//line /usr/local/go/src/crypto/x509/x509.go:1042
	_go_fuzz_dep_.CoverTab[20121]++
							return false
//line /usr/local/go/src/crypto/x509/x509.go:1043
	// _ = "end of CoverTab[20121]"
}

// marshalSANs marshals a list of addresses into a the contents of an X.509
//line /usr/local/go/src/crypto/x509/x509.go:1046
// SubjectAlternativeName extension.
//line /usr/local/go/src/crypto/x509/x509.go:1048
func marshalSANs(dnsNames, emailAddresses []string, ipAddresses []net.IP, uris []*url.URL) (derBytes []byte, err error) {
//line /usr/local/go/src/crypto/x509/x509.go:1048
	_go_fuzz_dep_.CoverTab[20125]++
							var rawValues []asn1.RawValue
							for _, name := range dnsNames {
//line /usr/local/go/src/crypto/x509/x509.go:1050
		_go_fuzz_dep_.CoverTab[20130]++
								if err := isIA5String(name); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1051
			_go_fuzz_dep_.CoverTab[20132]++
									return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1052
			// _ = "end of CoverTab[20132]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1053
			_go_fuzz_dep_.CoverTab[20133]++
//line /usr/local/go/src/crypto/x509/x509.go:1053
			// _ = "end of CoverTab[20133]"
//line /usr/local/go/src/crypto/x509/x509.go:1053
		}
//line /usr/local/go/src/crypto/x509/x509.go:1053
		// _ = "end of CoverTab[20130]"
//line /usr/local/go/src/crypto/x509/x509.go:1053
		_go_fuzz_dep_.CoverTab[20131]++
								rawValues = append(rawValues, asn1.RawValue{Tag: nameTypeDNS, Class: 2, Bytes: []byte(name)})
//line /usr/local/go/src/crypto/x509/x509.go:1054
		// _ = "end of CoverTab[20131]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:1055
	// _ = "end of CoverTab[20125]"
//line /usr/local/go/src/crypto/x509/x509.go:1055
	_go_fuzz_dep_.CoverTab[20126]++
							for _, email := range emailAddresses {
//line /usr/local/go/src/crypto/x509/x509.go:1056
		_go_fuzz_dep_.CoverTab[20134]++
								if err := isIA5String(email); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1057
			_go_fuzz_dep_.CoverTab[20136]++
									return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1058
			// _ = "end of CoverTab[20136]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1059
			_go_fuzz_dep_.CoverTab[20137]++
//line /usr/local/go/src/crypto/x509/x509.go:1059
			// _ = "end of CoverTab[20137]"
//line /usr/local/go/src/crypto/x509/x509.go:1059
		}
//line /usr/local/go/src/crypto/x509/x509.go:1059
		// _ = "end of CoverTab[20134]"
//line /usr/local/go/src/crypto/x509/x509.go:1059
		_go_fuzz_dep_.CoverTab[20135]++
								rawValues = append(rawValues, asn1.RawValue{Tag: nameTypeEmail, Class: 2, Bytes: []byte(email)})
//line /usr/local/go/src/crypto/x509/x509.go:1060
		// _ = "end of CoverTab[20135]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:1061
	// _ = "end of CoverTab[20126]"
//line /usr/local/go/src/crypto/x509/x509.go:1061
	_go_fuzz_dep_.CoverTab[20127]++
							for _, rawIP := range ipAddresses {
//line /usr/local/go/src/crypto/x509/x509.go:1062
		_go_fuzz_dep_.CoverTab[20138]++

								ip := rawIP.To4()
								if ip == nil {
//line /usr/local/go/src/crypto/x509/x509.go:1065
			_go_fuzz_dep_.CoverTab[20140]++
									ip = rawIP
//line /usr/local/go/src/crypto/x509/x509.go:1066
			// _ = "end of CoverTab[20140]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1067
			_go_fuzz_dep_.CoverTab[20141]++
//line /usr/local/go/src/crypto/x509/x509.go:1067
			// _ = "end of CoverTab[20141]"
//line /usr/local/go/src/crypto/x509/x509.go:1067
		}
//line /usr/local/go/src/crypto/x509/x509.go:1067
		// _ = "end of CoverTab[20138]"
//line /usr/local/go/src/crypto/x509/x509.go:1067
		_go_fuzz_dep_.CoverTab[20139]++
								rawValues = append(rawValues, asn1.RawValue{Tag: nameTypeIP, Class: 2, Bytes: ip})
//line /usr/local/go/src/crypto/x509/x509.go:1068
		// _ = "end of CoverTab[20139]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:1069
	// _ = "end of CoverTab[20127]"
//line /usr/local/go/src/crypto/x509/x509.go:1069
	_go_fuzz_dep_.CoverTab[20128]++
							for _, uri := range uris {
//line /usr/local/go/src/crypto/x509/x509.go:1070
		_go_fuzz_dep_.CoverTab[20142]++
								uriStr := uri.String()
								if err := isIA5String(uriStr); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1072
			_go_fuzz_dep_.CoverTab[20144]++
									return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1073
			// _ = "end of CoverTab[20144]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1074
			_go_fuzz_dep_.CoverTab[20145]++
//line /usr/local/go/src/crypto/x509/x509.go:1074
			// _ = "end of CoverTab[20145]"
//line /usr/local/go/src/crypto/x509/x509.go:1074
		}
//line /usr/local/go/src/crypto/x509/x509.go:1074
		// _ = "end of CoverTab[20142]"
//line /usr/local/go/src/crypto/x509/x509.go:1074
		_go_fuzz_dep_.CoverTab[20143]++
								rawValues = append(rawValues, asn1.RawValue{Tag: nameTypeURI, Class: 2, Bytes: []byte(uriStr)})
//line /usr/local/go/src/crypto/x509/x509.go:1075
		// _ = "end of CoverTab[20143]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:1076
	// _ = "end of CoverTab[20128]"
//line /usr/local/go/src/crypto/x509/x509.go:1076
	_go_fuzz_dep_.CoverTab[20129]++
							return asn1.Marshal(rawValues)
//line /usr/local/go/src/crypto/x509/x509.go:1077
	// _ = "end of CoverTab[20129]"
}

func isIA5String(s string) error {
//line /usr/local/go/src/crypto/x509/x509.go:1080
	_go_fuzz_dep_.CoverTab[20146]++
							for _, r := range s {
//line /usr/local/go/src/crypto/x509/x509.go:1081
		_go_fuzz_dep_.CoverTab[20148]++

								if r > unicode.MaxASCII {
//line /usr/local/go/src/crypto/x509/x509.go:1083
			_go_fuzz_dep_.CoverTab[20149]++
									return fmt.Errorf("x509: %q cannot be encoded as an IA5String", s)
//line /usr/local/go/src/crypto/x509/x509.go:1084
			// _ = "end of CoverTab[20149]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1085
			_go_fuzz_dep_.CoverTab[20150]++
//line /usr/local/go/src/crypto/x509/x509.go:1085
			// _ = "end of CoverTab[20150]"
//line /usr/local/go/src/crypto/x509/x509.go:1085
		}
//line /usr/local/go/src/crypto/x509/x509.go:1085
		// _ = "end of CoverTab[20148]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:1086
	// _ = "end of CoverTab[20146]"
//line /usr/local/go/src/crypto/x509/x509.go:1086
	_go_fuzz_dep_.CoverTab[20147]++

							return nil
//line /usr/local/go/src/crypto/x509/x509.go:1088
	// _ = "end of CoverTab[20147]"
}

func buildCertExtensions(template *Certificate, subjectIsEmpty bool, authorityKeyId []byte, subjectKeyId []byte) (ret []pkix.Extension, err error) {
//line /usr/local/go/src/crypto/x509/x509.go:1091
	_go_fuzz_dep_.CoverTab[20151]++
							ret = make([]pkix.Extension, 10)
							n := 0

							if template.KeyUsage != 0 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1095
		_go_fuzz_dep_.CoverTab[20162]++
//line /usr/local/go/src/crypto/x509/x509.go:1095
		return !oidInExtensions(oidExtensionKeyUsage, template.ExtraExtensions)
								// _ = "end of CoverTab[20162]"
//line /usr/local/go/src/crypto/x509/x509.go:1096
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1096
		_go_fuzz_dep_.CoverTab[20163]++
								ret[n], err = marshalKeyUsage(template.KeyUsage)
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1098
			_go_fuzz_dep_.CoverTab[20165]++
									return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1099
			// _ = "end of CoverTab[20165]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1100
			_go_fuzz_dep_.CoverTab[20166]++
//line /usr/local/go/src/crypto/x509/x509.go:1100
			// _ = "end of CoverTab[20166]"
//line /usr/local/go/src/crypto/x509/x509.go:1100
		}
//line /usr/local/go/src/crypto/x509/x509.go:1100
		// _ = "end of CoverTab[20163]"
//line /usr/local/go/src/crypto/x509/x509.go:1100
		_go_fuzz_dep_.CoverTab[20164]++
								n++
//line /usr/local/go/src/crypto/x509/x509.go:1101
		// _ = "end of CoverTab[20164]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1102
		_go_fuzz_dep_.CoverTab[20167]++
//line /usr/local/go/src/crypto/x509/x509.go:1102
		// _ = "end of CoverTab[20167]"
//line /usr/local/go/src/crypto/x509/x509.go:1102
	}
//line /usr/local/go/src/crypto/x509/x509.go:1102
	// _ = "end of CoverTab[20151]"
//line /usr/local/go/src/crypto/x509/x509.go:1102
	_go_fuzz_dep_.CoverTab[20152]++

							if (len(template.ExtKeyUsage) > 0 || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1104
		_go_fuzz_dep_.CoverTab[20168]++
//line /usr/local/go/src/crypto/x509/x509.go:1104
		return len(template.UnknownExtKeyUsage) > 0
//line /usr/local/go/src/crypto/x509/x509.go:1104
		// _ = "end of CoverTab[20168]"
//line /usr/local/go/src/crypto/x509/x509.go:1104
	}()) && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1104
		_go_fuzz_dep_.CoverTab[20169]++
//line /usr/local/go/src/crypto/x509/x509.go:1104
		return !oidInExtensions(oidExtensionExtendedKeyUsage, template.ExtraExtensions)
								// _ = "end of CoverTab[20169]"
//line /usr/local/go/src/crypto/x509/x509.go:1105
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1105
		_go_fuzz_dep_.CoverTab[20170]++
								ret[n], err = marshalExtKeyUsage(template.ExtKeyUsage, template.UnknownExtKeyUsage)
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1107
			_go_fuzz_dep_.CoverTab[20172]++
									return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1108
			// _ = "end of CoverTab[20172]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1109
			_go_fuzz_dep_.CoverTab[20173]++
//line /usr/local/go/src/crypto/x509/x509.go:1109
			// _ = "end of CoverTab[20173]"
//line /usr/local/go/src/crypto/x509/x509.go:1109
		}
//line /usr/local/go/src/crypto/x509/x509.go:1109
		// _ = "end of CoverTab[20170]"
//line /usr/local/go/src/crypto/x509/x509.go:1109
		_go_fuzz_dep_.CoverTab[20171]++
								n++
//line /usr/local/go/src/crypto/x509/x509.go:1110
		// _ = "end of CoverTab[20171]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1111
		_go_fuzz_dep_.CoverTab[20174]++
//line /usr/local/go/src/crypto/x509/x509.go:1111
		// _ = "end of CoverTab[20174]"
//line /usr/local/go/src/crypto/x509/x509.go:1111
	}
//line /usr/local/go/src/crypto/x509/x509.go:1111
	// _ = "end of CoverTab[20152]"
//line /usr/local/go/src/crypto/x509/x509.go:1111
	_go_fuzz_dep_.CoverTab[20153]++

							if template.BasicConstraintsValid && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1113
		_go_fuzz_dep_.CoverTab[20175]++
//line /usr/local/go/src/crypto/x509/x509.go:1113
		return !oidInExtensions(oidExtensionBasicConstraints, template.ExtraExtensions)
//line /usr/local/go/src/crypto/x509/x509.go:1113
		// _ = "end of CoverTab[20175]"
//line /usr/local/go/src/crypto/x509/x509.go:1113
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1113
		_go_fuzz_dep_.CoverTab[20176]++
								ret[n], err = marshalBasicConstraints(template.IsCA, template.MaxPathLen, template.MaxPathLenZero)
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1115
			_go_fuzz_dep_.CoverTab[20178]++
									return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1116
			// _ = "end of CoverTab[20178]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1117
			_go_fuzz_dep_.CoverTab[20179]++
//line /usr/local/go/src/crypto/x509/x509.go:1117
			// _ = "end of CoverTab[20179]"
//line /usr/local/go/src/crypto/x509/x509.go:1117
		}
//line /usr/local/go/src/crypto/x509/x509.go:1117
		// _ = "end of CoverTab[20176]"
//line /usr/local/go/src/crypto/x509/x509.go:1117
		_go_fuzz_dep_.CoverTab[20177]++
								n++
//line /usr/local/go/src/crypto/x509/x509.go:1118
		// _ = "end of CoverTab[20177]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1119
		_go_fuzz_dep_.CoverTab[20180]++
//line /usr/local/go/src/crypto/x509/x509.go:1119
		// _ = "end of CoverTab[20180]"
//line /usr/local/go/src/crypto/x509/x509.go:1119
	}
//line /usr/local/go/src/crypto/x509/x509.go:1119
	// _ = "end of CoverTab[20153]"
//line /usr/local/go/src/crypto/x509/x509.go:1119
	_go_fuzz_dep_.CoverTab[20154]++

							if len(subjectKeyId) > 0 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1121
		_go_fuzz_dep_.CoverTab[20181]++
//line /usr/local/go/src/crypto/x509/x509.go:1121
		return !oidInExtensions(oidExtensionSubjectKeyId, template.ExtraExtensions)
//line /usr/local/go/src/crypto/x509/x509.go:1121
		// _ = "end of CoverTab[20181]"
//line /usr/local/go/src/crypto/x509/x509.go:1121
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1121
		_go_fuzz_dep_.CoverTab[20182]++
								ret[n].Id = oidExtensionSubjectKeyId
								ret[n].Value, err = asn1.Marshal(subjectKeyId)
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1124
			_go_fuzz_dep_.CoverTab[20184]++
									return
//line /usr/local/go/src/crypto/x509/x509.go:1125
			// _ = "end of CoverTab[20184]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1126
			_go_fuzz_dep_.CoverTab[20185]++
//line /usr/local/go/src/crypto/x509/x509.go:1126
			// _ = "end of CoverTab[20185]"
//line /usr/local/go/src/crypto/x509/x509.go:1126
		}
//line /usr/local/go/src/crypto/x509/x509.go:1126
		// _ = "end of CoverTab[20182]"
//line /usr/local/go/src/crypto/x509/x509.go:1126
		_go_fuzz_dep_.CoverTab[20183]++
								n++
//line /usr/local/go/src/crypto/x509/x509.go:1127
		// _ = "end of CoverTab[20183]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1128
		_go_fuzz_dep_.CoverTab[20186]++
//line /usr/local/go/src/crypto/x509/x509.go:1128
		// _ = "end of CoverTab[20186]"
//line /usr/local/go/src/crypto/x509/x509.go:1128
	}
//line /usr/local/go/src/crypto/x509/x509.go:1128
	// _ = "end of CoverTab[20154]"
//line /usr/local/go/src/crypto/x509/x509.go:1128
	_go_fuzz_dep_.CoverTab[20155]++

							if len(authorityKeyId) > 0 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1130
		_go_fuzz_dep_.CoverTab[20187]++
//line /usr/local/go/src/crypto/x509/x509.go:1130
		return !oidInExtensions(oidExtensionAuthorityKeyId, template.ExtraExtensions)
//line /usr/local/go/src/crypto/x509/x509.go:1130
		// _ = "end of CoverTab[20187]"
//line /usr/local/go/src/crypto/x509/x509.go:1130
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1130
		_go_fuzz_dep_.CoverTab[20188]++
								ret[n].Id = oidExtensionAuthorityKeyId
								ret[n].Value, err = asn1.Marshal(authKeyId{authorityKeyId})
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1133
			_go_fuzz_dep_.CoverTab[20190]++
									return
//line /usr/local/go/src/crypto/x509/x509.go:1134
			// _ = "end of CoverTab[20190]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1135
			_go_fuzz_dep_.CoverTab[20191]++
//line /usr/local/go/src/crypto/x509/x509.go:1135
			// _ = "end of CoverTab[20191]"
//line /usr/local/go/src/crypto/x509/x509.go:1135
		}
//line /usr/local/go/src/crypto/x509/x509.go:1135
		// _ = "end of CoverTab[20188]"
//line /usr/local/go/src/crypto/x509/x509.go:1135
		_go_fuzz_dep_.CoverTab[20189]++
								n++
//line /usr/local/go/src/crypto/x509/x509.go:1136
		// _ = "end of CoverTab[20189]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1137
		_go_fuzz_dep_.CoverTab[20192]++
//line /usr/local/go/src/crypto/x509/x509.go:1137
		// _ = "end of CoverTab[20192]"
//line /usr/local/go/src/crypto/x509/x509.go:1137
	}
//line /usr/local/go/src/crypto/x509/x509.go:1137
	// _ = "end of CoverTab[20155]"
//line /usr/local/go/src/crypto/x509/x509.go:1137
	_go_fuzz_dep_.CoverTab[20156]++

							if (len(template.OCSPServer) > 0 || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1139
		_go_fuzz_dep_.CoverTab[20193]++
//line /usr/local/go/src/crypto/x509/x509.go:1139
		return len(template.IssuingCertificateURL) > 0
//line /usr/local/go/src/crypto/x509/x509.go:1139
		// _ = "end of CoverTab[20193]"
//line /usr/local/go/src/crypto/x509/x509.go:1139
	}()) && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1139
		_go_fuzz_dep_.CoverTab[20194]++
//line /usr/local/go/src/crypto/x509/x509.go:1139
		return !oidInExtensions(oidExtensionAuthorityInfoAccess, template.ExtraExtensions)
								// _ = "end of CoverTab[20194]"
//line /usr/local/go/src/crypto/x509/x509.go:1140
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1140
		_go_fuzz_dep_.CoverTab[20195]++
								ret[n].Id = oidExtensionAuthorityInfoAccess
								var aiaValues []authorityInfoAccess
								for _, name := range template.OCSPServer {
//line /usr/local/go/src/crypto/x509/x509.go:1143
			_go_fuzz_dep_.CoverTab[20199]++
									aiaValues = append(aiaValues, authorityInfoAccess{
				Method:		oidAuthorityInfoAccessOcsp,
				Location:	asn1.RawValue{Tag: 6, Class: 2, Bytes: []byte(name)},
			})
//line /usr/local/go/src/crypto/x509/x509.go:1147
			// _ = "end of CoverTab[20199]"
		}
//line /usr/local/go/src/crypto/x509/x509.go:1148
		// _ = "end of CoverTab[20195]"
//line /usr/local/go/src/crypto/x509/x509.go:1148
		_go_fuzz_dep_.CoverTab[20196]++
								for _, name := range template.IssuingCertificateURL {
//line /usr/local/go/src/crypto/x509/x509.go:1149
			_go_fuzz_dep_.CoverTab[20200]++
									aiaValues = append(aiaValues, authorityInfoAccess{
				Method:		oidAuthorityInfoAccessIssuers,
				Location:	asn1.RawValue{Tag: 6, Class: 2, Bytes: []byte(name)},
			})
//line /usr/local/go/src/crypto/x509/x509.go:1153
			// _ = "end of CoverTab[20200]"
		}
//line /usr/local/go/src/crypto/x509/x509.go:1154
		// _ = "end of CoverTab[20196]"
//line /usr/local/go/src/crypto/x509/x509.go:1154
		_go_fuzz_dep_.CoverTab[20197]++
								ret[n].Value, err = asn1.Marshal(aiaValues)
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1156
			_go_fuzz_dep_.CoverTab[20201]++
									return
//line /usr/local/go/src/crypto/x509/x509.go:1157
			// _ = "end of CoverTab[20201]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1158
			_go_fuzz_dep_.CoverTab[20202]++
//line /usr/local/go/src/crypto/x509/x509.go:1158
			// _ = "end of CoverTab[20202]"
//line /usr/local/go/src/crypto/x509/x509.go:1158
		}
//line /usr/local/go/src/crypto/x509/x509.go:1158
		// _ = "end of CoverTab[20197]"
//line /usr/local/go/src/crypto/x509/x509.go:1158
		_go_fuzz_dep_.CoverTab[20198]++
								n++
//line /usr/local/go/src/crypto/x509/x509.go:1159
		// _ = "end of CoverTab[20198]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1160
		_go_fuzz_dep_.CoverTab[20203]++
//line /usr/local/go/src/crypto/x509/x509.go:1160
		// _ = "end of CoverTab[20203]"
//line /usr/local/go/src/crypto/x509/x509.go:1160
	}
//line /usr/local/go/src/crypto/x509/x509.go:1160
	// _ = "end of CoverTab[20156]"
//line /usr/local/go/src/crypto/x509/x509.go:1160
	_go_fuzz_dep_.CoverTab[20157]++

							if (len(template.DNSNames) > 0 || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1162
		_go_fuzz_dep_.CoverTab[20204]++
//line /usr/local/go/src/crypto/x509/x509.go:1162
		return len(template.EmailAddresses) > 0
//line /usr/local/go/src/crypto/x509/x509.go:1162
		// _ = "end of CoverTab[20204]"
//line /usr/local/go/src/crypto/x509/x509.go:1162
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1162
		_go_fuzz_dep_.CoverTab[20205]++
//line /usr/local/go/src/crypto/x509/x509.go:1162
		return len(template.IPAddresses) > 0
//line /usr/local/go/src/crypto/x509/x509.go:1162
		// _ = "end of CoverTab[20205]"
//line /usr/local/go/src/crypto/x509/x509.go:1162
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1162
		_go_fuzz_dep_.CoverTab[20206]++
//line /usr/local/go/src/crypto/x509/x509.go:1162
		return len(template.URIs) > 0
//line /usr/local/go/src/crypto/x509/x509.go:1162
		// _ = "end of CoverTab[20206]"
//line /usr/local/go/src/crypto/x509/x509.go:1162
	}()) && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1162
		_go_fuzz_dep_.CoverTab[20207]++
//line /usr/local/go/src/crypto/x509/x509.go:1162
		return !oidInExtensions(oidExtensionSubjectAltName, template.ExtraExtensions)
								// _ = "end of CoverTab[20207]"
//line /usr/local/go/src/crypto/x509/x509.go:1163
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1163
		_go_fuzz_dep_.CoverTab[20208]++
								ret[n].Id = oidExtensionSubjectAltName

//line /usr/local/go/src/crypto/x509/x509.go:1168
		ret[n].Critical = subjectIsEmpty
		ret[n].Value, err = marshalSANs(template.DNSNames, template.EmailAddresses, template.IPAddresses, template.URIs)
		if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1170
			_go_fuzz_dep_.CoverTab[20210]++
									return
//line /usr/local/go/src/crypto/x509/x509.go:1171
			// _ = "end of CoverTab[20210]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1172
			_go_fuzz_dep_.CoverTab[20211]++
//line /usr/local/go/src/crypto/x509/x509.go:1172
			// _ = "end of CoverTab[20211]"
//line /usr/local/go/src/crypto/x509/x509.go:1172
		}
//line /usr/local/go/src/crypto/x509/x509.go:1172
		// _ = "end of CoverTab[20208]"
//line /usr/local/go/src/crypto/x509/x509.go:1172
		_go_fuzz_dep_.CoverTab[20209]++
								n++
//line /usr/local/go/src/crypto/x509/x509.go:1173
		// _ = "end of CoverTab[20209]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1174
		_go_fuzz_dep_.CoverTab[20212]++
//line /usr/local/go/src/crypto/x509/x509.go:1174
		// _ = "end of CoverTab[20212]"
//line /usr/local/go/src/crypto/x509/x509.go:1174
	}
//line /usr/local/go/src/crypto/x509/x509.go:1174
	// _ = "end of CoverTab[20157]"
//line /usr/local/go/src/crypto/x509/x509.go:1174
	_go_fuzz_dep_.CoverTab[20158]++

							if len(template.PolicyIdentifiers) > 0 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1176
		_go_fuzz_dep_.CoverTab[20213]++
//line /usr/local/go/src/crypto/x509/x509.go:1176
		return !oidInExtensions(oidExtensionCertificatePolicies, template.ExtraExtensions)
								// _ = "end of CoverTab[20213]"
//line /usr/local/go/src/crypto/x509/x509.go:1177
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1177
		_go_fuzz_dep_.CoverTab[20214]++
								ret[n], err = marshalCertificatePolicies(template.PolicyIdentifiers)
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1179
			_go_fuzz_dep_.CoverTab[20216]++
									return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1180
			// _ = "end of CoverTab[20216]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1181
			_go_fuzz_dep_.CoverTab[20217]++
//line /usr/local/go/src/crypto/x509/x509.go:1181
			// _ = "end of CoverTab[20217]"
//line /usr/local/go/src/crypto/x509/x509.go:1181
		}
//line /usr/local/go/src/crypto/x509/x509.go:1181
		// _ = "end of CoverTab[20214]"
//line /usr/local/go/src/crypto/x509/x509.go:1181
		_go_fuzz_dep_.CoverTab[20215]++
								n++
//line /usr/local/go/src/crypto/x509/x509.go:1182
		// _ = "end of CoverTab[20215]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1183
		_go_fuzz_dep_.CoverTab[20218]++
//line /usr/local/go/src/crypto/x509/x509.go:1183
		// _ = "end of CoverTab[20218]"
//line /usr/local/go/src/crypto/x509/x509.go:1183
	}
//line /usr/local/go/src/crypto/x509/x509.go:1183
	// _ = "end of CoverTab[20158]"
//line /usr/local/go/src/crypto/x509/x509.go:1183
	_go_fuzz_dep_.CoverTab[20159]++

							if (len(template.PermittedDNSDomains) > 0 || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1185
		_go_fuzz_dep_.CoverTab[20219]++
//line /usr/local/go/src/crypto/x509/x509.go:1185
		return len(template.ExcludedDNSDomains) > 0
//line /usr/local/go/src/crypto/x509/x509.go:1185
		// _ = "end of CoverTab[20219]"
//line /usr/local/go/src/crypto/x509/x509.go:1185
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1185
		_go_fuzz_dep_.CoverTab[20220]++
//line /usr/local/go/src/crypto/x509/x509.go:1185
		return len(template.PermittedIPRanges) > 0
								// _ = "end of CoverTab[20220]"
//line /usr/local/go/src/crypto/x509/x509.go:1186
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1186
		_go_fuzz_dep_.CoverTab[20221]++
//line /usr/local/go/src/crypto/x509/x509.go:1186
		return len(template.ExcludedIPRanges) > 0
//line /usr/local/go/src/crypto/x509/x509.go:1186
		// _ = "end of CoverTab[20221]"
//line /usr/local/go/src/crypto/x509/x509.go:1186
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1186
		_go_fuzz_dep_.CoverTab[20222]++
//line /usr/local/go/src/crypto/x509/x509.go:1186
		return len(template.PermittedEmailAddresses) > 0
								// _ = "end of CoverTab[20222]"
//line /usr/local/go/src/crypto/x509/x509.go:1187
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1187
		_go_fuzz_dep_.CoverTab[20223]++
//line /usr/local/go/src/crypto/x509/x509.go:1187
		return len(template.ExcludedEmailAddresses) > 0
//line /usr/local/go/src/crypto/x509/x509.go:1187
		// _ = "end of CoverTab[20223]"
//line /usr/local/go/src/crypto/x509/x509.go:1187
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1187
		_go_fuzz_dep_.CoverTab[20224]++
//line /usr/local/go/src/crypto/x509/x509.go:1187
		return len(template.PermittedURIDomains) > 0
								// _ = "end of CoverTab[20224]"
//line /usr/local/go/src/crypto/x509/x509.go:1188
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1188
		_go_fuzz_dep_.CoverTab[20225]++
//line /usr/local/go/src/crypto/x509/x509.go:1188
		return len(template.ExcludedURIDomains) > 0
//line /usr/local/go/src/crypto/x509/x509.go:1188
		// _ = "end of CoverTab[20225]"
//line /usr/local/go/src/crypto/x509/x509.go:1188
	}()) && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1188
		_go_fuzz_dep_.CoverTab[20226]++
//line /usr/local/go/src/crypto/x509/x509.go:1188
		return !oidInExtensions(oidExtensionNameConstraints, template.ExtraExtensions)
								// _ = "end of CoverTab[20226]"
//line /usr/local/go/src/crypto/x509/x509.go:1189
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1189
		_go_fuzz_dep_.CoverTab[20227]++
								ret[n].Id = oidExtensionNameConstraints
								ret[n].Critical = template.PermittedDNSDomainsCritical

								ipAndMask := func(ipNet *net.IPNet) []byte {
//line /usr/local/go/src/crypto/x509/x509.go:1193
			_go_fuzz_dep_.CoverTab[20234]++
									maskedIP := ipNet.IP.Mask(ipNet.Mask)
									ipAndMask := make([]byte, 0, len(maskedIP)+len(ipNet.Mask))
									ipAndMask = append(ipAndMask, maskedIP...)
									ipAndMask = append(ipAndMask, ipNet.Mask...)
									return ipAndMask
//line /usr/local/go/src/crypto/x509/x509.go:1198
			// _ = "end of CoverTab[20234]"
		}
//line /usr/local/go/src/crypto/x509/x509.go:1199
		// _ = "end of CoverTab[20227]"
//line /usr/local/go/src/crypto/x509/x509.go:1199
		_go_fuzz_dep_.CoverTab[20228]++

								serialiseConstraints := func(dns []string, ips []*net.IPNet, emails []string, uriDomains []string) (der []byte, err error) {
//line /usr/local/go/src/crypto/x509/x509.go:1201
			_go_fuzz_dep_.CoverTab[20235]++
									var b cryptobyte.Builder

									for _, name := range dns {
//line /usr/local/go/src/crypto/x509/x509.go:1204
				_go_fuzz_dep_.CoverTab[20240]++
										if err = isIA5String(name); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1205
					_go_fuzz_dep_.CoverTab[20242]++
											return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1206
					// _ = "end of CoverTab[20242]"
				} else {
//line /usr/local/go/src/crypto/x509/x509.go:1207
					_go_fuzz_dep_.CoverTab[20243]++
//line /usr/local/go/src/crypto/x509/x509.go:1207
					// _ = "end of CoverTab[20243]"
//line /usr/local/go/src/crypto/x509/x509.go:1207
				}
//line /usr/local/go/src/crypto/x509/x509.go:1207
				// _ = "end of CoverTab[20240]"
//line /usr/local/go/src/crypto/x509/x509.go:1207
				_go_fuzz_dep_.CoverTab[20241]++

										b.AddASN1(cryptobyte_asn1.SEQUENCE, func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/x509/x509.go:1209
					_go_fuzz_dep_.CoverTab[20244]++
											b.AddASN1(cryptobyte_asn1.Tag(2).ContextSpecific(), func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/x509/x509.go:1210
						_go_fuzz_dep_.CoverTab[20245]++
												b.AddBytes([]byte(name))
//line /usr/local/go/src/crypto/x509/x509.go:1211
						// _ = "end of CoverTab[20245]"
					})
//line /usr/local/go/src/crypto/x509/x509.go:1212
					// _ = "end of CoverTab[20244]"
				})
//line /usr/local/go/src/crypto/x509/x509.go:1213
				// _ = "end of CoverTab[20241]"
			}
//line /usr/local/go/src/crypto/x509/x509.go:1214
			// _ = "end of CoverTab[20235]"
//line /usr/local/go/src/crypto/x509/x509.go:1214
			_go_fuzz_dep_.CoverTab[20236]++

									for _, ipNet := range ips {
//line /usr/local/go/src/crypto/x509/x509.go:1216
				_go_fuzz_dep_.CoverTab[20246]++
										b.AddASN1(cryptobyte_asn1.SEQUENCE, func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/x509/x509.go:1217
					_go_fuzz_dep_.CoverTab[20247]++
											b.AddASN1(cryptobyte_asn1.Tag(7).ContextSpecific(), func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/x509/x509.go:1218
						_go_fuzz_dep_.CoverTab[20248]++
												b.AddBytes(ipAndMask(ipNet))
//line /usr/local/go/src/crypto/x509/x509.go:1219
						// _ = "end of CoverTab[20248]"
					})
//line /usr/local/go/src/crypto/x509/x509.go:1220
					// _ = "end of CoverTab[20247]"
				})
//line /usr/local/go/src/crypto/x509/x509.go:1221
				// _ = "end of CoverTab[20246]"
			}
//line /usr/local/go/src/crypto/x509/x509.go:1222
			// _ = "end of CoverTab[20236]"
//line /usr/local/go/src/crypto/x509/x509.go:1222
			_go_fuzz_dep_.CoverTab[20237]++

									for _, email := range emails {
//line /usr/local/go/src/crypto/x509/x509.go:1224
				_go_fuzz_dep_.CoverTab[20249]++
										if err = isIA5String(email); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1225
					_go_fuzz_dep_.CoverTab[20251]++
											return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1226
					// _ = "end of CoverTab[20251]"
				} else {
//line /usr/local/go/src/crypto/x509/x509.go:1227
					_go_fuzz_dep_.CoverTab[20252]++
//line /usr/local/go/src/crypto/x509/x509.go:1227
					// _ = "end of CoverTab[20252]"
//line /usr/local/go/src/crypto/x509/x509.go:1227
				}
//line /usr/local/go/src/crypto/x509/x509.go:1227
				// _ = "end of CoverTab[20249]"
//line /usr/local/go/src/crypto/x509/x509.go:1227
				_go_fuzz_dep_.CoverTab[20250]++

										b.AddASN1(cryptobyte_asn1.SEQUENCE, func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/x509/x509.go:1229
					_go_fuzz_dep_.CoverTab[20253]++
											b.AddASN1(cryptobyte_asn1.Tag(1).ContextSpecific(), func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/x509/x509.go:1230
						_go_fuzz_dep_.CoverTab[20254]++
												b.AddBytes([]byte(email))
//line /usr/local/go/src/crypto/x509/x509.go:1231
						// _ = "end of CoverTab[20254]"
					})
//line /usr/local/go/src/crypto/x509/x509.go:1232
					// _ = "end of CoverTab[20253]"
				})
//line /usr/local/go/src/crypto/x509/x509.go:1233
				// _ = "end of CoverTab[20250]"
			}
//line /usr/local/go/src/crypto/x509/x509.go:1234
			// _ = "end of CoverTab[20237]"
//line /usr/local/go/src/crypto/x509/x509.go:1234
			_go_fuzz_dep_.CoverTab[20238]++

									for _, uriDomain := range uriDomains {
//line /usr/local/go/src/crypto/x509/x509.go:1236
				_go_fuzz_dep_.CoverTab[20255]++
										if err = isIA5String(uriDomain); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1237
					_go_fuzz_dep_.CoverTab[20257]++
											return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1238
					// _ = "end of CoverTab[20257]"
				} else {
//line /usr/local/go/src/crypto/x509/x509.go:1239
					_go_fuzz_dep_.CoverTab[20258]++
//line /usr/local/go/src/crypto/x509/x509.go:1239
					// _ = "end of CoverTab[20258]"
//line /usr/local/go/src/crypto/x509/x509.go:1239
				}
//line /usr/local/go/src/crypto/x509/x509.go:1239
				// _ = "end of CoverTab[20255]"
//line /usr/local/go/src/crypto/x509/x509.go:1239
				_go_fuzz_dep_.CoverTab[20256]++

										b.AddASN1(cryptobyte_asn1.SEQUENCE, func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/x509/x509.go:1241
					_go_fuzz_dep_.CoverTab[20259]++
											b.AddASN1(cryptobyte_asn1.Tag(6).ContextSpecific(), func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/x509/x509.go:1242
						_go_fuzz_dep_.CoverTab[20260]++
												b.AddBytes([]byte(uriDomain))
//line /usr/local/go/src/crypto/x509/x509.go:1243
						// _ = "end of CoverTab[20260]"
					})
//line /usr/local/go/src/crypto/x509/x509.go:1244
					// _ = "end of CoverTab[20259]"
				})
//line /usr/local/go/src/crypto/x509/x509.go:1245
				// _ = "end of CoverTab[20256]"
			}
//line /usr/local/go/src/crypto/x509/x509.go:1246
			// _ = "end of CoverTab[20238]"
//line /usr/local/go/src/crypto/x509/x509.go:1246
			_go_fuzz_dep_.CoverTab[20239]++

									return b.Bytes()
//line /usr/local/go/src/crypto/x509/x509.go:1248
			// _ = "end of CoverTab[20239]"
		}
//line /usr/local/go/src/crypto/x509/x509.go:1249
		// _ = "end of CoverTab[20228]"
//line /usr/local/go/src/crypto/x509/x509.go:1249
		_go_fuzz_dep_.CoverTab[20229]++

								permitted, err := serialiseConstraints(template.PermittedDNSDomains, template.PermittedIPRanges, template.PermittedEmailAddresses, template.PermittedURIDomains)
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1252
			_go_fuzz_dep_.CoverTab[20261]++
									return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1253
			// _ = "end of CoverTab[20261]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1254
			_go_fuzz_dep_.CoverTab[20262]++
//line /usr/local/go/src/crypto/x509/x509.go:1254
			// _ = "end of CoverTab[20262]"
//line /usr/local/go/src/crypto/x509/x509.go:1254
		}
//line /usr/local/go/src/crypto/x509/x509.go:1254
		// _ = "end of CoverTab[20229]"
//line /usr/local/go/src/crypto/x509/x509.go:1254
		_go_fuzz_dep_.CoverTab[20230]++

								excluded, err := serialiseConstraints(template.ExcludedDNSDomains, template.ExcludedIPRanges, template.ExcludedEmailAddresses, template.ExcludedURIDomains)
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1257
			_go_fuzz_dep_.CoverTab[20263]++
									return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1258
			// _ = "end of CoverTab[20263]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1259
			_go_fuzz_dep_.CoverTab[20264]++
//line /usr/local/go/src/crypto/x509/x509.go:1259
			// _ = "end of CoverTab[20264]"
//line /usr/local/go/src/crypto/x509/x509.go:1259
		}
//line /usr/local/go/src/crypto/x509/x509.go:1259
		// _ = "end of CoverTab[20230]"
//line /usr/local/go/src/crypto/x509/x509.go:1259
		_go_fuzz_dep_.CoverTab[20231]++

								var b cryptobyte.Builder
								b.AddASN1(cryptobyte_asn1.SEQUENCE, func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/x509/x509.go:1262
			_go_fuzz_dep_.CoverTab[20265]++
									if len(permitted) > 0 {
//line /usr/local/go/src/crypto/x509/x509.go:1263
				_go_fuzz_dep_.CoverTab[20267]++
										b.AddASN1(cryptobyte_asn1.Tag(0).ContextSpecific().Constructed(), func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/x509/x509.go:1264
					_go_fuzz_dep_.CoverTab[20268]++
											b.AddBytes(permitted)
//line /usr/local/go/src/crypto/x509/x509.go:1265
					// _ = "end of CoverTab[20268]"
				})
//line /usr/local/go/src/crypto/x509/x509.go:1266
				// _ = "end of CoverTab[20267]"
			} else {
//line /usr/local/go/src/crypto/x509/x509.go:1267
				_go_fuzz_dep_.CoverTab[20269]++
//line /usr/local/go/src/crypto/x509/x509.go:1267
				// _ = "end of CoverTab[20269]"
//line /usr/local/go/src/crypto/x509/x509.go:1267
			}
//line /usr/local/go/src/crypto/x509/x509.go:1267
			// _ = "end of CoverTab[20265]"
//line /usr/local/go/src/crypto/x509/x509.go:1267
			_go_fuzz_dep_.CoverTab[20266]++

									if len(excluded) > 0 {
//line /usr/local/go/src/crypto/x509/x509.go:1269
				_go_fuzz_dep_.CoverTab[20270]++
										b.AddASN1(cryptobyte_asn1.Tag(1).ContextSpecific().Constructed(), func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/x509/x509.go:1270
					_go_fuzz_dep_.CoverTab[20271]++
											b.AddBytes(excluded)
//line /usr/local/go/src/crypto/x509/x509.go:1271
					// _ = "end of CoverTab[20271]"
				})
//line /usr/local/go/src/crypto/x509/x509.go:1272
				// _ = "end of CoverTab[20270]"
			} else {
//line /usr/local/go/src/crypto/x509/x509.go:1273
				_go_fuzz_dep_.CoverTab[20272]++
//line /usr/local/go/src/crypto/x509/x509.go:1273
				// _ = "end of CoverTab[20272]"
//line /usr/local/go/src/crypto/x509/x509.go:1273
			}
//line /usr/local/go/src/crypto/x509/x509.go:1273
			// _ = "end of CoverTab[20266]"
		})
//line /usr/local/go/src/crypto/x509/x509.go:1274
		// _ = "end of CoverTab[20231]"
//line /usr/local/go/src/crypto/x509/x509.go:1274
		_go_fuzz_dep_.CoverTab[20232]++

								ret[n].Value, err = b.Bytes()
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1277
			_go_fuzz_dep_.CoverTab[20273]++
									return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1278
			// _ = "end of CoverTab[20273]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1279
			_go_fuzz_dep_.CoverTab[20274]++
//line /usr/local/go/src/crypto/x509/x509.go:1279
			// _ = "end of CoverTab[20274]"
//line /usr/local/go/src/crypto/x509/x509.go:1279
		}
//line /usr/local/go/src/crypto/x509/x509.go:1279
		// _ = "end of CoverTab[20232]"
//line /usr/local/go/src/crypto/x509/x509.go:1279
		_go_fuzz_dep_.CoverTab[20233]++
								n++
//line /usr/local/go/src/crypto/x509/x509.go:1280
		// _ = "end of CoverTab[20233]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1281
		_go_fuzz_dep_.CoverTab[20275]++
//line /usr/local/go/src/crypto/x509/x509.go:1281
		// _ = "end of CoverTab[20275]"
//line /usr/local/go/src/crypto/x509/x509.go:1281
	}
//line /usr/local/go/src/crypto/x509/x509.go:1281
	// _ = "end of CoverTab[20159]"
//line /usr/local/go/src/crypto/x509/x509.go:1281
	_go_fuzz_dep_.CoverTab[20160]++

							if len(template.CRLDistributionPoints) > 0 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1283
		_go_fuzz_dep_.CoverTab[20276]++
//line /usr/local/go/src/crypto/x509/x509.go:1283
		return !oidInExtensions(oidExtensionCRLDistributionPoints, template.ExtraExtensions)
								// _ = "end of CoverTab[20276]"
//line /usr/local/go/src/crypto/x509/x509.go:1284
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1284
		_go_fuzz_dep_.CoverTab[20277]++
								ret[n].Id = oidExtensionCRLDistributionPoints

								var crlDp []distributionPoint
								for _, name := range template.CRLDistributionPoints {
//line /usr/local/go/src/crypto/x509/x509.go:1288
			_go_fuzz_dep_.CoverTab[20280]++
									dp := distributionPoint{
				DistributionPoint: distributionPointName{
					FullName: []asn1.RawValue{
						{Tag: 6, Class: 2, Bytes: []byte(name)},
					},
				},
			}
									crlDp = append(crlDp, dp)
//line /usr/local/go/src/crypto/x509/x509.go:1296
			// _ = "end of CoverTab[20280]"
		}
//line /usr/local/go/src/crypto/x509/x509.go:1297
		// _ = "end of CoverTab[20277]"
//line /usr/local/go/src/crypto/x509/x509.go:1297
		_go_fuzz_dep_.CoverTab[20278]++

								ret[n].Value, err = asn1.Marshal(crlDp)
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1300
			_go_fuzz_dep_.CoverTab[20281]++
									return
//line /usr/local/go/src/crypto/x509/x509.go:1301
			// _ = "end of CoverTab[20281]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1302
			_go_fuzz_dep_.CoverTab[20282]++
//line /usr/local/go/src/crypto/x509/x509.go:1302
			// _ = "end of CoverTab[20282]"
//line /usr/local/go/src/crypto/x509/x509.go:1302
		}
//line /usr/local/go/src/crypto/x509/x509.go:1302
		// _ = "end of CoverTab[20278]"
//line /usr/local/go/src/crypto/x509/x509.go:1302
		_go_fuzz_dep_.CoverTab[20279]++
								n++
//line /usr/local/go/src/crypto/x509/x509.go:1303
		// _ = "end of CoverTab[20279]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1304
		_go_fuzz_dep_.CoverTab[20283]++
//line /usr/local/go/src/crypto/x509/x509.go:1304
		// _ = "end of CoverTab[20283]"
//line /usr/local/go/src/crypto/x509/x509.go:1304
	}
//line /usr/local/go/src/crypto/x509/x509.go:1304
	// _ = "end of CoverTab[20160]"
//line /usr/local/go/src/crypto/x509/x509.go:1304
	_go_fuzz_dep_.CoverTab[20161]++

//line /usr/local/go/src/crypto/x509/x509.go:1310
	return append(ret[:n], template.ExtraExtensions...), nil
//line /usr/local/go/src/crypto/x509/x509.go:1310
	// _ = "end of CoverTab[20161]"
}

func marshalKeyUsage(ku KeyUsage) (pkix.Extension, error) {
//line /usr/local/go/src/crypto/x509/x509.go:1313
	_go_fuzz_dep_.CoverTab[20284]++
							ext := pkix.Extension{Id: oidExtensionKeyUsage, Critical: true}

							var a [2]byte
							a[0] = reverseBitsInAByte(byte(ku))
							a[1] = reverseBitsInAByte(byte(ku >> 8))

							l := 1
							if a[1] != 0 {
//line /usr/local/go/src/crypto/x509/x509.go:1321
		_go_fuzz_dep_.CoverTab[20286]++
								l = 2
//line /usr/local/go/src/crypto/x509/x509.go:1322
		// _ = "end of CoverTab[20286]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1323
		_go_fuzz_dep_.CoverTab[20287]++
//line /usr/local/go/src/crypto/x509/x509.go:1323
		// _ = "end of CoverTab[20287]"
//line /usr/local/go/src/crypto/x509/x509.go:1323
	}
//line /usr/local/go/src/crypto/x509/x509.go:1323
	// _ = "end of CoverTab[20284]"
//line /usr/local/go/src/crypto/x509/x509.go:1323
	_go_fuzz_dep_.CoverTab[20285]++

							bitString := a[:l]
							var err error
							ext.Value, err = asn1.Marshal(asn1.BitString{Bytes: bitString, BitLength: asn1BitLength(bitString)})
							return ext, err
//line /usr/local/go/src/crypto/x509/x509.go:1328
	// _ = "end of CoverTab[20285]"
}

func marshalExtKeyUsage(extUsages []ExtKeyUsage, unknownUsages []asn1.ObjectIdentifier) (pkix.Extension, error) {
//line /usr/local/go/src/crypto/x509/x509.go:1331
	_go_fuzz_dep_.CoverTab[20288]++
							ext := pkix.Extension{Id: oidExtensionExtendedKeyUsage}

							oids := make([]asn1.ObjectIdentifier, len(extUsages)+len(unknownUsages))
							for i, u := range extUsages {
//line /usr/local/go/src/crypto/x509/x509.go:1335
		_go_fuzz_dep_.CoverTab[20290]++
								if oid, ok := oidFromExtKeyUsage(u); ok {
//line /usr/local/go/src/crypto/x509/x509.go:1336
			_go_fuzz_dep_.CoverTab[20291]++
									oids[i] = oid
//line /usr/local/go/src/crypto/x509/x509.go:1337
			// _ = "end of CoverTab[20291]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1338
			_go_fuzz_dep_.CoverTab[20292]++
									return ext, errors.New("x509: unknown extended key usage")
//line /usr/local/go/src/crypto/x509/x509.go:1339
			// _ = "end of CoverTab[20292]"
		}
//line /usr/local/go/src/crypto/x509/x509.go:1340
		// _ = "end of CoverTab[20290]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:1341
	// _ = "end of CoverTab[20288]"
//line /usr/local/go/src/crypto/x509/x509.go:1341
	_go_fuzz_dep_.CoverTab[20289]++

							copy(oids[len(extUsages):], unknownUsages)

							var err error
							ext.Value, err = asn1.Marshal(oids)
							return ext, err
//line /usr/local/go/src/crypto/x509/x509.go:1347
	// _ = "end of CoverTab[20289]"
}

func marshalBasicConstraints(isCA bool, maxPathLen int, maxPathLenZero bool) (pkix.Extension, error) {
//line /usr/local/go/src/crypto/x509/x509.go:1350
	_go_fuzz_dep_.CoverTab[20293]++
							ext := pkix.Extension{Id: oidExtensionBasicConstraints, Critical: true}

//line /usr/local/go/src/crypto/x509/x509.go:1355
	if maxPathLen == 0 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1355
		_go_fuzz_dep_.CoverTab[20295]++
//line /usr/local/go/src/crypto/x509/x509.go:1355
		return !maxPathLenZero
//line /usr/local/go/src/crypto/x509/x509.go:1355
		// _ = "end of CoverTab[20295]"
//line /usr/local/go/src/crypto/x509/x509.go:1355
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1355
		_go_fuzz_dep_.CoverTab[20296]++
								maxPathLen = -1
//line /usr/local/go/src/crypto/x509/x509.go:1356
		// _ = "end of CoverTab[20296]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1357
		_go_fuzz_dep_.CoverTab[20297]++
//line /usr/local/go/src/crypto/x509/x509.go:1357
		// _ = "end of CoverTab[20297]"
//line /usr/local/go/src/crypto/x509/x509.go:1357
	}
//line /usr/local/go/src/crypto/x509/x509.go:1357
	// _ = "end of CoverTab[20293]"
//line /usr/local/go/src/crypto/x509/x509.go:1357
	_go_fuzz_dep_.CoverTab[20294]++
							var err error
							ext.Value, err = asn1.Marshal(basicConstraints{isCA, maxPathLen})
							return ext, err
//line /usr/local/go/src/crypto/x509/x509.go:1360
	// _ = "end of CoverTab[20294]"
}

func marshalCertificatePolicies(policyIdentifiers []asn1.ObjectIdentifier) (pkix.Extension, error) {
//line /usr/local/go/src/crypto/x509/x509.go:1363
	_go_fuzz_dep_.CoverTab[20298]++
							ext := pkix.Extension{Id: oidExtensionCertificatePolicies}
							policies := make([]policyInformation, len(policyIdentifiers))
							for i, policy := range policyIdentifiers {
//line /usr/local/go/src/crypto/x509/x509.go:1366
		_go_fuzz_dep_.CoverTab[20300]++
								policies[i].Policy = policy
//line /usr/local/go/src/crypto/x509/x509.go:1367
		// _ = "end of CoverTab[20300]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:1368
	// _ = "end of CoverTab[20298]"
//line /usr/local/go/src/crypto/x509/x509.go:1368
	_go_fuzz_dep_.CoverTab[20299]++
							var err error
							ext.Value, err = asn1.Marshal(policies)
							return ext, err
//line /usr/local/go/src/crypto/x509/x509.go:1371
	// _ = "end of CoverTab[20299]"
}

func buildCSRExtensions(template *CertificateRequest) ([]pkix.Extension, error) {
//line /usr/local/go/src/crypto/x509/x509.go:1374
	_go_fuzz_dep_.CoverTab[20301]++
							var ret []pkix.Extension

							if (len(template.DNSNames) > 0 || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1377
		_go_fuzz_dep_.CoverTab[20303]++
//line /usr/local/go/src/crypto/x509/x509.go:1377
		return len(template.EmailAddresses) > 0
//line /usr/local/go/src/crypto/x509/x509.go:1377
		// _ = "end of CoverTab[20303]"
//line /usr/local/go/src/crypto/x509/x509.go:1377
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1377
		_go_fuzz_dep_.CoverTab[20304]++
//line /usr/local/go/src/crypto/x509/x509.go:1377
		return len(template.IPAddresses) > 0
//line /usr/local/go/src/crypto/x509/x509.go:1377
		// _ = "end of CoverTab[20304]"
//line /usr/local/go/src/crypto/x509/x509.go:1377
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1377
		_go_fuzz_dep_.CoverTab[20305]++
//line /usr/local/go/src/crypto/x509/x509.go:1377
		return len(template.URIs) > 0
//line /usr/local/go/src/crypto/x509/x509.go:1377
		// _ = "end of CoverTab[20305]"
//line /usr/local/go/src/crypto/x509/x509.go:1377
	}()) && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1377
		_go_fuzz_dep_.CoverTab[20306]++
//line /usr/local/go/src/crypto/x509/x509.go:1377
		return !oidInExtensions(oidExtensionSubjectAltName, template.ExtraExtensions)
								// _ = "end of CoverTab[20306]"
//line /usr/local/go/src/crypto/x509/x509.go:1378
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1378
		_go_fuzz_dep_.CoverTab[20307]++
								sanBytes, err := marshalSANs(template.DNSNames, template.EmailAddresses, template.IPAddresses, template.URIs)
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1380
			_go_fuzz_dep_.CoverTab[20309]++
									return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1381
			// _ = "end of CoverTab[20309]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1382
			_go_fuzz_dep_.CoverTab[20310]++
//line /usr/local/go/src/crypto/x509/x509.go:1382
			// _ = "end of CoverTab[20310]"
//line /usr/local/go/src/crypto/x509/x509.go:1382
		}
//line /usr/local/go/src/crypto/x509/x509.go:1382
		// _ = "end of CoverTab[20307]"
//line /usr/local/go/src/crypto/x509/x509.go:1382
		_go_fuzz_dep_.CoverTab[20308]++

								ret = append(ret, pkix.Extension{
			Id:	oidExtensionSubjectAltName,
			Value:	sanBytes,
		})
//line /usr/local/go/src/crypto/x509/x509.go:1387
		// _ = "end of CoverTab[20308]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1388
		_go_fuzz_dep_.CoverTab[20311]++
//line /usr/local/go/src/crypto/x509/x509.go:1388
		// _ = "end of CoverTab[20311]"
//line /usr/local/go/src/crypto/x509/x509.go:1388
	}
//line /usr/local/go/src/crypto/x509/x509.go:1388
	// _ = "end of CoverTab[20301]"
//line /usr/local/go/src/crypto/x509/x509.go:1388
	_go_fuzz_dep_.CoverTab[20302]++

							return append(ret, template.ExtraExtensions...), nil
//line /usr/local/go/src/crypto/x509/x509.go:1390
	// _ = "end of CoverTab[20302]"
}

func subjectBytes(cert *Certificate) ([]byte, error) {
//line /usr/local/go/src/crypto/x509/x509.go:1393
	_go_fuzz_dep_.CoverTab[20312]++
							if len(cert.RawSubject) > 0 {
//line /usr/local/go/src/crypto/x509/x509.go:1394
		_go_fuzz_dep_.CoverTab[20314]++
								return cert.RawSubject, nil
//line /usr/local/go/src/crypto/x509/x509.go:1395
		// _ = "end of CoverTab[20314]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1396
		_go_fuzz_dep_.CoverTab[20315]++
//line /usr/local/go/src/crypto/x509/x509.go:1396
		// _ = "end of CoverTab[20315]"
//line /usr/local/go/src/crypto/x509/x509.go:1396
	}
//line /usr/local/go/src/crypto/x509/x509.go:1396
	// _ = "end of CoverTab[20312]"
//line /usr/local/go/src/crypto/x509/x509.go:1396
	_go_fuzz_dep_.CoverTab[20313]++

							return asn1.Marshal(cert.Subject.ToRDNSequence())
//line /usr/local/go/src/crypto/x509/x509.go:1398
	// _ = "end of CoverTab[20313]"
}

// signingParamsForPublicKey returns the parameters to use for signing with
//line /usr/local/go/src/crypto/x509/x509.go:1401
// priv. If requestedSigAlgo is not zero then it overrides the default
//line /usr/local/go/src/crypto/x509/x509.go:1401
// signature algorithm.
//line /usr/local/go/src/crypto/x509/x509.go:1404
func signingParamsForPublicKey(pub any, requestedSigAlgo SignatureAlgorithm) (hashFunc crypto.Hash, sigAlgo pkix.AlgorithmIdentifier, err error) {
//line /usr/local/go/src/crypto/x509/x509.go:1404
	_go_fuzz_dep_.CoverTab[20316]++
							var pubType PublicKeyAlgorithm

							switch pub := pub.(type) {
	case *rsa.PublicKey:
//line /usr/local/go/src/crypto/x509/x509.go:1408
		_go_fuzz_dep_.CoverTab[20322]++
								pubType = RSA
								hashFunc = crypto.SHA256
								sigAlgo.Algorithm = oidSignatureSHA256WithRSA
								sigAlgo.Parameters = asn1.NullRawValue
//line /usr/local/go/src/crypto/x509/x509.go:1412
		// _ = "end of CoverTab[20322]"

	case *ecdsa.PublicKey:
//line /usr/local/go/src/crypto/x509/x509.go:1414
		_go_fuzz_dep_.CoverTab[20323]++
								pubType = ECDSA

								switch pub.Curve {
		case elliptic.P224(), elliptic.P256():
//line /usr/local/go/src/crypto/x509/x509.go:1418
			_go_fuzz_dep_.CoverTab[20326]++
									hashFunc = crypto.SHA256
									sigAlgo.Algorithm = oidSignatureECDSAWithSHA256
//line /usr/local/go/src/crypto/x509/x509.go:1420
			// _ = "end of CoverTab[20326]"
		case elliptic.P384():
//line /usr/local/go/src/crypto/x509/x509.go:1421
			_go_fuzz_dep_.CoverTab[20327]++
									hashFunc = crypto.SHA384
									sigAlgo.Algorithm = oidSignatureECDSAWithSHA384
//line /usr/local/go/src/crypto/x509/x509.go:1423
			// _ = "end of CoverTab[20327]"
		case elliptic.P521():
//line /usr/local/go/src/crypto/x509/x509.go:1424
			_go_fuzz_dep_.CoverTab[20328]++
									hashFunc = crypto.SHA512
									sigAlgo.Algorithm = oidSignatureECDSAWithSHA512
//line /usr/local/go/src/crypto/x509/x509.go:1426
			// _ = "end of CoverTab[20328]"
		default:
//line /usr/local/go/src/crypto/x509/x509.go:1427
			_go_fuzz_dep_.CoverTab[20329]++
									err = errors.New("x509: unknown elliptic curve")
//line /usr/local/go/src/crypto/x509/x509.go:1428
			// _ = "end of CoverTab[20329]"
		}
//line /usr/local/go/src/crypto/x509/x509.go:1429
		// _ = "end of CoverTab[20323]"

	case ed25519.PublicKey:
//line /usr/local/go/src/crypto/x509/x509.go:1431
		_go_fuzz_dep_.CoverTab[20324]++
								pubType = Ed25519
								sigAlgo.Algorithm = oidSignatureEd25519
//line /usr/local/go/src/crypto/x509/x509.go:1433
		// _ = "end of CoverTab[20324]"

	default:
//line /usr/local/go/src/crypto/x509/x509.go:1435
		_go_fuzz_dep_.CoverTab[20325]++
								err = errors.New("x509: only RSA, ECDSA and Ed25519 keys supported")
//line /usr/local/go/src/crypto/x509/x509.go:1436
		// _ = "end of CoverTab[20325]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:1437
	// _ = "end of CoverTab[20316]"
//line /usr/local/go/src/crypto/x509/x509.go:1437
	_go_fuzz_dep_.CoverTab[20317]++

							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1439
		_go_fuzz_dep_.CoverTab[20330]++
								return
//line /usr/local/go/src/crypto/x509/x509.go:1440
		// _ = "end of CoverTab[20330]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1441
		_go_fuzz_dep_.CoverTab[20331]++
//line /usr/local/go/src/crypto/x509/x509.go:1441
		// _ = "end of CoverTab[20331]"
//line /usr/local/go/src/crypto/x509/x509.go:1441
	}
//line /usr/local/go/src/crypto/x509/x509.go:1441
	// _ = "end of CoverTab[20317]"
//line /usr/local/go/src/crypto/x509/x509.go:1441
	_go_fuzz_dep_.CoverTab[20318]++

							if requestedSigAlgo == 0 {
//line /usr/local/go/src/crypto/x509/x509.go:1443
		_go_fuzz_dep_.CoverTab[20332]++
								return
//line /usr/local/go/src/crypto/x509/x509.go:1444
		// _ = "end of CoverTab[20332]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1445
		_go_fuzz_dep_.CoverTab[20333]++
//line /usr/local/go/src/crypto/x509/x509.go:1445
		// _ = "end of CoverTab[20333]"
//line /usr/local/go/src/crypto/x509/x509.go:1445
	}
//line /usr/local/go/src/crypto/x509/x509.go:1445
	// _ = "end of CoverTab[20318]"
//line /usr/local/go/src/crypto/x509/x509.go:1445
	_go_fuzz_dep_.CoverTab[20319]++

							found := false
							for _, details := range signatureAlgorithmDetails {
//line /usr/local/go/src/crypto/x509/x509.go:1448
		_go_fuzz_dep_.CoverTab[20334]++
								if details.algo == requestedSigAlgo {
//line /usr/local/go/src/crypto/x509/x509.go:1449
			_go_fuzz_dep_.CoverTab[20335]++
									if details.pubKeyAlgo != pubType {
//line /usr/local/go/src/crypto/x509/x509.go:1450
				_go_fuzz_dep_.CoverTab[20340]++
										err = errors.New("x509: requested SignatureAlgorithm does not match private key type")
										return
//line /usr/local/go/src/crypto/x509/x509.go:1452
				// _ = "end of CoverTab[20340]"
			} else {
//line /usr/local/go/src/crypto/x509/x509.go:1453
				_go_fuzz_dep_.CoverTab[20341]++
//line /usr/local/go/src/crypto/x509/x509.go:1453
				// _ = "end of CoverTab[20341]"
//line /usr/local/go/src/crypto/x509/x509.go:1453
			}
//line /usr/local/go/src/crypto/x509/x509.go:1453
			// _ = "end of CoverTab[20335]"
//line /usr/local/go/src/crypto/x509/x509.go:1453
			_go_fuzz_dep_.CoverTab[20336]++
									sigAlgo.Algorithm, hashFunc = details.oid, details.hash
									if hashFunc == 0 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1455
				_go_fuzz_dep_.CoverTab[20342]++
//line /usr/local/go/src/crypto/x509/x509.go:1455
				return pubType != Ed25519
//line /usr/local/go/src/crypto/x509/x509.go:1455
				// _ = "end of CoverTab[20342]"
//line /usr/local/go/src/crypto/x509/x509.go:1455
			}() {
//line /usr/local/go/src/crypto/x509/x509.go:1455
				_go_fuzz_dep_.CoverTab[20343]++
										err = errors.New("x509: cannot sign with hash function requested")
										return
//line /usr/local/go/src/crypto/x509/x509.go:1457
				// _ = "end of CoverTab[20343]"
			} else {
//line /usr/local/go/src/crypto/x509/x509.go:1458
				_go_fuzz_dep_.CoverTab[20344]++
//line /usr/local/go/src/crypto/x509/x509.go:1458
				// _ = "end of CoverTab[20344]"
//line /usr/local/go/src/crypto/x509/x509.go:1458
			}
//line /usr/local/go/src/crypto/x509/x509.go:1458
			// _ = "end of CoverTab[20336]"
//line /usr/local/go/src/crypto/x509/x509.go:1458
			_go_fuzz_dep_.CoverTab[20337]++
									if hashFunc == crypto.MD5 {
//line /usr/local/go/src/crypto/x509/x509.go:1459
				_go_fuzz_dep_.CoverTab[20345]++
										err = errors.New("x509: signing with MD5 is not supported")
										return
//line /usr/local/go/src/crypto/x509/x509.go:1461
				// _ = "end of CoverTab[20345]"
			} else {
//line /usr/local/go/src/crypto/x509/x509.go:1462
				_go_fuzz_dep_.CoverTab[20346]++
//line /usr/local/go/src/crypto/x509/x509.go:1462
				// _ = "end of CoverTab[20346]"
//line /usr/local/go/src/crypto/x509/x509.go:1462
			}
//line /usr/local/go/src/crypto/x509/x509.go:1462
			// _ = "end of CoverTab[20337]"
//line /usr/local/go/src/crypto/x509/x509.go:1462
			_go_fuzz_dep_.CoverTab[20338]++
									if requestedSigAlgo.isRSAPSS() {
//line /usr/local/go/src/crypto/x509/x509.go:1463
				_go_fuzz_dep_.CoverTab[20347]++
										sigAlgo.Parameters = hashToPSSParameters[hashFunc]
//line /usr/local/go/src/crypto/x509/x509.go:1464
				// _ = "end of CoverTab[20347]"
			} else {
//line /usr/local/go/src/crypto/x509/x509.go:1465
				_go_fuzz_dep_.CoverTab[20348]++
//line /usr/local/go/src/crypto/x509/x509.go:1465
				// _ = "end of CoverTab[20348]"
//line /usr/local/go/src/crypto/x509/x509.go:1465
			}
//line /usr/local/go/src/crypto/x509/x509.go:1465
			// _ = "end of CoverTab[20338]"
//line /usr/local/go/src/crypto/x509/x509.go:1465
			_go_fuzz_dep_.CoverTab[20339]++
									found = true
									break
//line /usr/local/go/src/crypto/x509/x509.go:1467
			// _ = "end of CoverTab[20339]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1468
			_go_fuzz_dep_.CoverTab[20349]++
//line /usr/local/go/src/crypto/x509/x509.go:1468
			// _ = "end of CoverTab[20349]"
//line /usr/local/go/src/crypto/x509/x509.go:1468
		}
//line /usr/local/go/src/crypto/x509/x509.go:1468
		// _ = "end of CoverTab[20334]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:1469
	// _ = "end of CoverTab[20319]"
//line /usr/local/go/src/crypto/x509/x509.go:1469
	_go_fuzz_dep_.CoverTab[20320]++

							if !found {
//line /usr/local/go/src/crypto/x509/x509.go:1471
		_go_fuzz_dep_.CoverTab[20350]++
								err = errors.New("x509: unknown SignatureAlgorithm")
//line /usr/local/go/src/crypto/x509/x509.go:1472
		// _ = "end of CoverTab[20350]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1473
		_go_fuzz_dep_.CoverTab[20351]++
//line /usr/local/go/src/crypto/x509/x509.go:1473
		// _ = "end of CoverTab[20351]"
//line /usr/local/go/src/crypto/x509/x509.go:1473
	}
//line /usr/local/go/src/crypto/x509/x509.go:1473
	// _ = "end of CoverTab[20320]"
//line /usr/local/go/src/crypto/x509/x509.go:1473
	_go_fuzz_dep_.CoverTab[20321]++

							return
//line /usr/local/go/src/crypto/x509/x509.go:1475
	// _ = "end of CoverTab[20321]"
}

// emptyASN1Subject is the ASN.1 DER encoding of an empty Subject, which is
//line /usr/local/go/src/crypto/x509/x509.go:1478
// just an empty SEQUENCE.
//line /usr/local/go/src/crypto/x509/x509.go:1480
var emptyASN1Subject = []byte{0x30, 0}

// CreateCertificate creates a new X.509 v3 certificate based on a template.
//line /usr/local/go/src/crypto/x509/x509.go:1482
// The following members of template are currently used:
//line /usr/local/go/src/crypto/x509/x509.go:1482
//
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - AuthorityKeyId
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - BasicConstraintsValid
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - CRLDistributionPoints
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - DNSNames
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - EmailAddresses
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - ExcludedDNSDomains
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - ExcludedEmailAddresses
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - ExcludedIPRanges
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - ExcludedURIDomains
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - ExtKeyUsage
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - ExtraExtensions
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - IPAddresses
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - IsCA
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - IssuingCertificateURL
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - KeyUsage
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - MaxPathLen
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - MaxPathLenZero
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - NotAfter
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - NotBefore
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - OCSPServer
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - PermittedDNSDomains
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - PermittedDNSDomainsCritical
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - PermittedEmailAddresses
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - PermittedIPRanges
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - PermittedURIDomains
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - PolicyIdentifiers
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - SerialNumber
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - SignatureAlgorithm
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - Subject
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - SubjectKeyId
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - URIs
//line /usr/local/go/src/crypto/x509/x509.go:1482
//   - UnknownExtKeyUsage
//line /usr/local/go/src/crypto/x509/x509.go:1482
//
//line /usr/local/go/src/crypto/x509/x509.go:1482
// The certificate is signed by parent. If parent is equal to template then the
//line /usr/local/go/src/crypto/x509/x509.go:1482
// certificate is self-signed. The parameter pub is the public key of the
//line /usr/local/go/src/crypto/x509/x509.go:1482
// certificate to be generated and priv is the private key of the signer.
//line /usr/local/go/src/crypto/x509/x509.go:1482
//
//line /usr/local/go/src/crypto/x509/x509.go:1482
// The returned slice is the certificate in DER encoding.
//line /usr/local/go/src/crypto/x509/x509.go:1482
//
//line /usr/local/go/src/crypto/x509/x509.go:1482
// The currently supported key types are *rsa.PublicKey, *ecdsa.PublicKey and
//line /usr/local/go/src/crypto/x509/x509.go:1482
// ed25519.PublicKey. pub must be a supported key type, and priv must be a
//line /usr/local/go/src/crypto/x509/x509.go:1482
// crypto.Signer with a supported public key.
//line /usr/local/go/src/crypto/x509/x509.go:1482
//
//line /usr/local/go/src/crypto/x509/x509.go:1482
// The AuthorityKeyId will be taken from the SubjectKeyId of parent, if any,
//line /usr/local/go/src/crypto/x509/x509.go:1482
// unless the resulting certificate is self-signed. Otherwise the value from
//line /usr/local/go/src/crypto/x509/x509.go:1482
// template will be used.
//line /usr/local/go/src/crypto/x509/x509.go:1482
//
//line /usr/local/go/src/crypto/x509/x509.go:1482
// If SubjectKeyId from template is empty and the template is a CA, SubjectKeyId
//line /usr/local/go/src/crypto/x509/x509.go:1482
// will be generated from the hash of the public key.
//line /usr/local/go/src/crypto/x509/x509.go:1534
func CreateCertificate(rand io.Reader, template, parent *Certificate, pub, priv any) ([]byte, error) {
//line /usr/local/go/src/crypto/x509/x509.go:1534
	_go_fuzz_dep_.CoverTab[20352]++
							key, ok := priv.(crypto.Signer)
							if !ok {
//line /usr/local/go/src/crypto/x509/x509.go:1536
		_go_fuzz_dep_.CoverTab[20372]++
								return nil, errors.New("x509: certificate private key does not implement crypto.Signer")
//line /usr/local/go/src/crypto/x509/x509.go:1537
		// _ = "end of CoverTab[20372]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1538
		_go_fuzz_dep_.CoverTab[20373]++
//line /usr/local/go/src/crypto/x509/x509.go:1538
		// _ = "end of CoverTab[20373]"
//line /usr/local/go/src/crypto/x509/x509.go:1538
	}
//line /usr/local/go/src/crypto/x509/x509.go:1538
	// _ = "end of CoverTab[20352]"
//line /usr/local/go/src/crypto/x509/x509.go:1538
	_go_fuzz_dep_.CoverTab[20353]++

							if template.SerialNumber == nil {
//line /usr/local/go/src/crypto/x509/x509.go:1540
		_go_fuzz_dep_.CoverTab[20374]++
								return nil, errors.New("x509: no SerialNumber given")
//line /usr/local/go/src/crypto/x509/x509.go:1541
		// _ = "end of CoverTab[20374]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1542
		_go_fuzz_dep_.CoverTab[20375]++
//line /usr/local/go/src/crypto/x509/x509.go:1542
		// _ = "end of CoverTab[20375]"
//line /usr/local/go/src/crypto/x509/x509.go:1542
	}
//line /usr/local/go/src/crypto/x509/x509.go:1542
	// _ = "end of CoverTab[20353]"
//line /usr/local/go/src/crypto/x509/x509.go:1542
	_go_fuzz_dep_.CoverTab[20354]++

//line /usr/local/go/src/crypto/x509/x509.go:1549
	if template.SerialNumber.Sign() == -1 {
//line /usr/local/go/src/crypto/x509/x509.go:1549
		_go_fuzz_dep_.CoverTab[20376]++
								return nil, errors.New("x509: serial number must be positive")
//line /usr/local/go/src/crypto/x509/x509.go:1550
		// _ = "end of CoverTab[20376]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1551
		_go_fuzz_dep_.CoverTab[20377]++
//line /usr/local/go/src/crypto/x509/x509.go:1551
		// _ = "end of CoverTab[20377]"
//line /usr/local/go/src/crypto/x509/x509.go:1551
	}
//line /usr/local/go/src/crypto/x509/x509.go:1551
	// _ = "end of CoverTab[20354]"
//line /usr/local/go/src/crypto/x509/x509.go:1551
	_go_fuzz_dep_.CoverTab[20355]++

							if template.BasicConstraintsValid && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1553
		_go_fuzz_dep_.CoverTab[20378]++
//line /usr/local/go/src/crypto/x509/x509.go:1553
		return !template.IsCA
//line /usr/local/go/src/crypto/x509/x509.go:1553
		// _ = "end of CoverTab[20378]"
//line /usr/local/go/src/crypto/x509/x509.go:1553
	}() && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1553
		_go_fuzz_dep_.CoverTab[20379]++
//line /usr/local/go/src/crypto/x509/x509.go:1553
		return template.MaxPathLen != -1
//line /usr/local/go/src/crypto/x509/x509.go:1553
		// _ = "end of CoverTab[20379]"
//line /usr/local/go/src/crypto/x509/x509.go:1553
	}() && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1553
		_go_fuzz_dep_.CoverTab[20380]++
//line /usr/local/go/src/crypto/x509/x509.go:1553
		return (template.MaxPathLen != 0 || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1553
			_go_fuzz_dep_.CoverTab[20381]++
//line /usr/local/go/src/crypto/x509/x509.go:1553
			return template.MaxPathLenZero
//line /usr/local/go/src/crypto/x509/x509.go:1553
			// _ = "end of CoverTab[20381]"
//line /usr/local/go/src/crypto/x509/x509.go:1553
		}())
//line /usr/local/go/src/crypto/x509/x509.go:1553
		// _ = "end of CoverTab[20380]"
//line /usr/local/go/src/crypto/x509/x509.go:1553
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1553
		_go_fuzz_dep_.CoverTab[20382]++
								return nil, errors.New("x509: only CAs are allowed to specify MaxPathLen")
//line /usr/local/go/src/crypto/x509/x509.go:1554
		// _ = "end of CoverTab[20382]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1555
		_go_fuzz_dep_.CoverTab[20383]++
//line /usr/local/go/src/crypto/x509/x509.go:1555
		// _ = "end of CoverTab[20383]"
//line /usr/local/go/src/crypto/x509/x509.go:1555
	}
//line /usr/local/go/src/crypto/x509/x509.go:1555
	// _ = "end of CoverTab[20355]"
//line /usr/local/go/src/crypto/x509/x509.go:1555
	_go_fuzz_dep_.CoverTab[20356]++

							hashFunc, signatureAlgorithm, err := signingParamsForPublicKey(key.Public(), template.SignatureAlgorithm)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1558
		_go_fuzz_dep_.CoverTab[20384]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1559
		// _ = "end of CoverTab[20384]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1560
		_go_fuzz_dep_.CoverTab[20385]++
//line /usr/local/go/src/crypto/x509/x509.go:1560
		// _ = "end of CoverTab[20385]"
//line /usr/local/go/src/crypto/x509/x509.go:1560
	}
//line /usr/local/go/src/crypto/x509/x509.go:1560
	// _ = "end of CoverTab[20356]"
//line /usr/local/go/src/crypto/x509/x509.go:1560
	_go_fuzz_dep_.CoverTab[20357]++

							publicKeyBytes, publicKeyAlgorithm, err := marshalPublicKey(pub)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1563
		_go_fuzz_dep_.CoverTab[20386]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1564
		// _ = "end of CoverTab[20386]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1565
		_go_fuzz_dep_.CoverTab[20387]++
//line /usr/local/go/src/crypto/x509/x509.go:1565
		// _ = "end of CoverTab[20387]"
//line /usr/local/go/src/crypto/x509/x509.go:1565
	}
//line /usr/local/go/src/crypto/x509/x509.go:1565
	// _ = "end of CoverTab[20357]"
//line /usr/local/go/src/crypto/x509/x509.go:1565
	_go_fuzz_dep_.CoverTab[20358]++
							if getPublicKeyAlgorithmFromOID(publicKeyAlgorithm.Algorithm) == UnknownPublicKeyAlgorithm {
//line /usr/local/go/src/crypto/x509/x509.go:1566
		_go_fuzz_dep_.CoverTab[20388]++
								return nil, fmt.Errorf("x509: unsupported public key type: %T", pub)
//line /usr/local/go/src/crypto/x509/x509.go:1567
		// _ = "end of CoverTab[20388]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1568
		_go_fuzz_dep_.CoverTab[20389]++
//line /usr/local/go/src/crypto/x509/x509.go:1568
		// _ = "end of CoverTab[20389]"
//line /usr/local/go/src/crypto/x509/x509.go:1568
	}
//line /usr/local/go/src/crypto/x509/x509.go:1568
	// _ = "end of CoverTab[20358]"
//line /usr/local/go/src/crypto/x509/x509.go:1568
	_go_fuzz_dep_.CoverTab[20359]++

							asn1Issuer, err := subjectBytes(parent)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1571
		_go_fuzz_dep_.CoverTab[20390]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1572
		// _ = "end of CoverTab[20390]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1573
		_go_fuzz_dep_.CoverTab[20391]++
//line /usr/local/go/src/crypto/x509/x509.go:1573
		// _ = "end of CoverTab[20391]"
//line /usr/local/go/src/crypto/x509/x509.go:1573
	}
//line /usr/local/go/src/crypto/x509/x509.go:1573
	// _ = "end of CoverTab[20359]"
//line /usr/local/go/src/crypto/x509/x509.go:1573
	_go_fuzz_dep_.CoverTab[20360]++

							asn1Subject, err := subjectBytes(template)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1576
		_go_fuzz_dep_.CoverTab[20392]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1577
		// _ = "end of CoverTab[20392]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1578
		_go_fuzz_dep_.CoverTab[20393]++
//line /usr/local/go/src/crypto/x509/x509.go:1578
		// _ = "end of CoverTab[20393]"
//line /usr/local/go/src/crypto/x509/x509.go:1578
	}
//line /usr/local/go/src/crypto/x509/x509.go:1578
	// _ = "end of CoverTab[20360]"
//line /usr/local/go/src/crypto/x509/x509.go:1578
	_go_fuzz_dep_.CoverTab[20361]++

							authorityKeyId := template.AuthorityKeyId
							if !bytes.Equal(asn1Issuer, asn1Subject) && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1581
		_go_fuzz_dep_.CoverTab[20394]++
//line /usr/local/go/src/crypto/x509/x509.go:1581
		return len(parent.SubjectKeyId) > 0
//line /usr/local/go/src/crypto/x509/x509.go:1581
		// _ = "end of CoverTab[20394]"
//line /usr/local/go/src/crypto/x509/x509.go:1581
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1581
		_go_fuzz_dep_.CoverTab[20395]++
								authorityKeyId = parent.SubjectKeyId
//line /usr/local/go/src/crypto/x509/x509.go:1582
		// _ = "end of CoverTab[20395]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1583
		_go_fuzz_dep_.CoverTab[20396]++
//line /usr/local/go/src/crypto/x509/x509.go:1583
		// _ = "end of CoverTab[20396]"
//line /usr/local/go/src/crypto/x509/x509.go:1583
	}
//line /usr/local/go/src/crypto/x509/x509.go:1583
	// _ = "end of CoverTab[20361]"
//line /usr/local/go/src/crypto/x509/x509.go:1583
	_go_fuzz_dep_.CoverTab[20362]++

							subjectKeyId := template.SubjectKeyId
							if len(subjectKeyId) == 0 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1586
		_go_fuzz_dep_.CoverTab[20397]++
//line /usr/local/go/src/crypto/x509/x509.go:1586
		return template.IsCA
//line /usr/local/go/src/crypto/x509/x509.go:1586
		// _ = "end of CoverTab[20397]"
//line /usr/local/go/src/crypto/x509/x509.go:1586
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1586
		_go_fuzz_dep_.CoverTab[20398]++

//line /usr/local/go/src/crypto/x509/x509.go:1591
		h := sha1.Sum(publicKeyBytes)
								subjectKeyId = h[:]
//line /usr/local/go/src/crypto/x509/x509.go:1592
		// _ = "end of CoverTab[20398]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1593
		_go_fuzz_dep_.CoverTab[20399]++
//line /usr/local/go/src/crypto/x509/x509.go:1593
		// _ = "end of CoverTab[20399]"
//line /usr/local/go/src/crypto/x509/x509.go:1593
	}
//line /usr/local/go/src/crypto/x509/x509.go:1593
	// _ = "end of CoverTab[20362]"
//line /usr/local/go/src/crypto/x509/x509.go:1593
	_go_fuzz_dep_.CoverTab[20363]++

	// Check that the signer's public key matches the private key, if available.
	type privateKey interface {
		Equal(crypto.PublicKey) bool
	}
	if privPub, ok := key.Public().(privateKey); !ok {
//line /usr/local/go/src/crypto/x509/x509.go:1599
		_go_fuzz_dep_.CoverTab[20400]++
								return nil, errors.New("x509: internal error: supported public key does not implement Equal")
//line /usr/local/go/src/crypto/x509/x509.go:1600
		// _ = "end of CoverTab[20400]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1601
		_go_fuzz_dep_.CoverTab[20401]++
//line /usr/local/go/src/crypto/x509/x509.go:1601
		if parent.PublicKey != nil && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1601
			_go_fuzz_dep_.CoverTab[20402]++
//line /usr/local/go/src/crypto/x509/x509.go:1601
			return !privPub.Equal(parent.PublicKey)
//line /usr/local/go/src/crypto/x509/x509.go:1601
			// _ = "end of CoverTab[20402]"
//line /usr/local/go/src/crypto/x509/x509.go:1601
		}() {
//line /usr/local/go/src/crypto/x509/x509.go:1601
			_go_fuzz_dep_.CoverTab[20403]++
									return nil, errors.New("x509: provided PrivateKey doesn't match parent's PublicKey")
//line /usr/local/go/src/crypto/x509/x509.go:1602
			// _ = "end of CoverTab[20403]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1603
			_go_fuzz_dep_.CoverTab[20404]++
//line /usr/local/go/src/crypto/x509/x509.go:1603
			// _ = "end of CoverTab[20404]"
//line /usr/local/go/src/crypto/x509/x509.go:1603
		}
//line /usr/local/go/src/crypto/x509/x509.go:1603
		// _ = "end of CoverTab[20401]"
//line /usr/local/go/src/crypto/x509/x509.go:1603
	}
//line /usr/local/go/src/crypto/x509/x509.go:1603
	// _ = "end of CoverTab[20363]"
//line /usr/local/go/src/crypto/x509/x509.go:1603
	_go_fuzz_dep_.CoverTab[20364]++

							extensions, err := buildCertExtensions(template, bytes.Equal(asn1Subject, emptyASN1Subject), authorityKeyId, subjectKeyId)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1606
		_go_fuzz_dep_.CoverTab[20405]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1607
		// _ = "end of CoverTab[20405]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1608
		_go_fuzz_dep_.CoverTab[20406]++
//line /usr/local/go/src/crypto/x509/x509.go:1608
		// _ = "end of CoverTab[20406]"
//line /usr/local/go/src/crypto/x509/x509.go:1608
	}
//line /usr/local/go/src/crypto/x509/x509.go:1608
	// _ = "end of CoverTab[20364]"
//line /usr/local/go/src/crypto/x509/x509.go:1608
	_go_fuzz_dep_.CoverTab[20365]++

							encodedPublicKey := asn1.BitString{BitLength: len(publicKeyBytes) * 8, Bytes: publicKeyBytes}
							c := tbsCertificate{
		Version:		2,
		SerialNumber:		template.SerialNumber,
		SignatureAlgorithm:	signatureAlgorithm,
		Issuer:			asn1.RawValue{FullBytes: asn1Issuer},
		Validity:		validity{template.NotBefore.UTC(), template.NotAfter.UTC()},
		Subject:		asn1.RawValue{FullBytes: asn1Subject},
		PublicKey:		publicKeyInfo{nil, publicKeyAlgorithm, encodedPublicKey},
		Extensions:		extensions,
	}

	tbsCertContents, err := asn1.Marshal(c)
	if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1623
		_go_fuzz_dep_.CoverTab[20407]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1624
		// _ = "end of CoverTab[20407]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1625
		_go_fuzz_dep_.CoverTab[20408]++
//line /usr/local/go/src/crypto/x509/x509.go:1625
		// _ = "end of CoverTab[20408]"
//line /usr/local/go/src/crypto/x509/x509.go:1625
	}
//line /usr/local/go/src/crypto/x509/x509.go:1625
	// _ = "end of CoverTab[20365]"
//line /usr/local/go/src/crypto/x509/x509.go:1625
	_go_fuzz_dep_.CoverTab[20366]++
							c.Raw = tbsCertContents

							signed := tbsCertContents
							if hashFunc != 0 {
//line /usr/local/go/src/crypto/x509/x509.go:1629
		_go_fuzz_dep_.CoverTab[20409]++
								h := hashFunc.New()
								h.Write(signed)
								signed = h.Sum(nil)
//line /usr/local/go/src/crypto/x509/x509.go:1632
		// _ = "end of CoverTab[20409]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1633
		_go_fuzz_dep_.CoverTab[20410]++
//line /usr/local/go/src/crypto/x509/x509.go:1633
		// _ = "end of CoverTab[20410]"
//line /usr/local/go/src/crypto/x509/x509.go:1633
	}
//line /usr/local/go/src/crypto/x509/x509.go:1633
	// _ = "end of CoverTab[20366]"
//line /usr/local/go/src/crypto/x509/x509.go:1633
	_go_fuzz_dep_.CoverTab[20367]++

							var signerOpts crypto.SignerOpts = hashFunc
							if template.SignatureAlgorithm != 0 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1636
		_go_fuzz_dep_.CoverTab[20411]++
//line /usr/local/go/src/crypto/x509/x509.go:1636
		return template.SignatureAlgorithm.isRSAPSS()
//line /usr/local/go/src/crypto/x509/x509.go:1636
		// _ = "end of CoverTab[20411]"
//line /usr/local/go/src/crypto/x509/x509.go:1636
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:1636
		_go_fuzz_dep_.CoverTab[20412]++
								signerOpts = &rsa.PSSOptions{
			SaltLength:	rsa.PSSSaltLengthEqualsHash,
			Hash:		hashFunc,
		}
//line /usr/local/go/src/crypto/x509/x509.go:1640
		// _ = "end of CoverTab[20412]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1641
		_go_fuzz_dep_.CoverTab[20413]++
//line /usr/local/go/src/crypto/x509/x509.go:1641
		// _ = "end of CoverTab[20413]"
//line /usr/local/go/src/crypto/x509/x509.go:1641
	}
//line /usr/local/go/src/crypto/x509/x509.go:1641
	// _ = "end of CoverTab[20367]"
//line /usr/local/go/src/crypto/x509/x509.go:1641
	_go_fuzz_dep_.CoverTab[20368]++

							var signature []byte
							signature, err = key.Sign(rand, signed, signerOpts)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1645
		_go_fuzz_dep_.CoverTab[20414]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1646
		// _ = "end of CoverTab[20414]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1647
		_go_fuzz_dep_.CoverTab[20415]++
//line /usr/local/go/src/crypto/x509/x509.go:1647
		// _ = "end of CoverTab[20415]"
//line /usr/local/go/src/crypto/x509/x509.go:1647
	}
//line /usr/local/go/src/crypto/x509/x509.go:1647
	// _ = "end of CoverTab[20368]"
//line /usr/local/go/src/crypto/x509/x509.go:1647
	_go_fuzz_dep_.CoverTab[20369]++

							signedCert, err := asn1.Marshal(certificate{
		c,
		signatureAlgorithm,
		asn1.BitString{Bytes: signature, BitLength: len(signature) * 8},
	})
	if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1654
		_go_fuzz_dep_.CoverTab[20416]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1655
		// _ = "end of CoverTab[20416]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1656
		_go_fuzz_dep_.CoverTab[20417]++
//line /usr/local/go/src/crypto/x509/x509.go:1656
		// _ = "end of CoverTab[20417]"
//line /usr/local/go/src/crypto/x509/x509.go:1656
	}
//line /usr/local/go/src/crypto/x509/x509.go:1656
	// _ = "end of CoverTab[20369]"
//line /usr/local/go/src/crypto/x509/x509.go:1656
	_go_fuzz_dep_.CoverTab[20370]++

//line /usr/local/go/src/crypto/x509/x509.go:1659
	if err := checkSignature(getSignatureAlgorithmFromAI(signatureAlgorithm), c.Raw, signature, key.Public(), true); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1659
		_go_fuzz_dep_.CoverTab[20418]++
								return nil, fmt.Errorf("x509: signature over certificate returned by signer is invalid: %w", err)
//line /usr/local/go/src/crypto/x509/x509.go:1660
		// _ = "end of CoverTab[20418]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1661
		_go_fuzz_dep_.CoverTab[20419]++
//line /usr/local/go/src/crypto/x509/x509.go:1661
		// _ = "end of CoverTab[20419]"
//line /usr/local/go/src/crypto/x509/x509.go:1661
	}
//line /usr/local/go/src/crypto/x509/x509.go:1661
	// _ = "end of CoverTab[20370]"
//line /usr/local/go/src/crypto/x509/x509.go:1661
	_go_fuzz_dep_.CoverTab[20371]++

							return signedCert, nil
//line /usr/local/go/src/crypto/x509/x509.go:1663
	// _ = "end of CoverTab[20371]"
}

// pemCRLPrefix is the magic string that indicates that we have a PEM encoded
//line /usr/local/go/src/crypto/x509/x509.go:1666
// CRL.
//line /usr/local/go/src/crypto/x509/x509.go:1668
var pemCRLPrefix = []byte("-----BEGIN X509 CRL")

// pemType is the type of a PEM encoded CRL.
var pemType = "X509 CRL"

// ParseCRL parses a CRL from the given bytes. It's often the case that PEM
//line /usr/local/go/src/crypto/x509/x509.go:1673
// encoded CRLs will appear where they should be DER encoded, so this function
//line /usr/local/go/src/crypto/x509/x509.go:1673
// will transparently handle PEM encoding as long as there isn't any leading
//line /usr/local/go/src/crypto/x509/x509.go:1673
// garbage.
//line /usr/local/go/src/crypto/x509/x509.go:1673
//
//line /usr/local/go/src/crypto/x509/x509.go:1673
// Deprecated: Use ParseRevocationList instead.
//line /usr/local/go/src/crypto/x509/x509.go:1679
func ParseCRL(crlBytes []byte) (*pkix.CertificateList, error) {
//line /usr/local/go/src/crypto/x509/x509.go:1679
	_go_fuzz_dep_.CoverTab[20420]++
							if bytes.HasPrefix(crlBytes, pemCRLPrefix) {
//line /usr/local/go/src/crypto/x509/x509.go:1680
		_go_fuzz_dep_.CoverTab[20422]++
								block, _ := pem.Decode(crlBytes)
								if block != nil && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1682
			_go_fuzz_dep_.CoverTab[20423]++
//line /usr/local/go/src/crypto/x509/x509.go:1682
			return block.Type == pemType
//line /usr/local/go/src/crypto/x509/x509.go:1682
			// _ = "end of CoverTab[20423]"
//line /usr/local/go/src/crypto/x509/x509.go:1682
		}() {
//line /usr/local/go/src/crypto/x509/x509.go:1682
			_go_fuzz_dep_.CoverTab[20424]++
									crlBytes = block.Bytes
//line /usr/local/go/src/crypto/x509/x509.go:1683
			// _ = "end of CoverTab[20424]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1684
			_go_fuzz_dep_.CoverTab[20425]++
//line /usr/local/go/src/crypto/x509/x509.go:1684
			// _ = "end of CoverTab[20425]"
//line /usr/local/go/src/crypto/x509/x509.go:1684
		}
//line /usr/local/go/src/crypto/x509/x509.go:1684
		// _ = "end of CoverTab[20422]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1685
		_go_fuzz_dep_.CoverTab[20426]++
//line /usr/local/go/src/crypto/x509/x509.go:1685
		// _ = "end of CoverTab[20426]"
//line /usr/local/go/src/crypto/x509/x509.go:1685
	}
//line /usr/local/go/src/crypto/x509/x509.go:1685
	// _ = "end of CoverTab[20420]"
//line /usr/local/go/src/crypto/x509/x509.go:1685
	_go_fuzz_dep_.CoverTab[20421]++
							return ParseDERCRL(crlBytes)
//line /usr/local/go/src/crypto/x509/x509.go:1686
	// _ = "end of CoverTab[20421]"
}

// ParseDERCRL parses a DER encoded CRL from the given bytes.
//line /usr/local/go/src/crypto/x509/x509.go:1689
//
//line /usr/local/go/src/crypto/x509/x509.go:1689
// Deprecated: Use ParseRevocationList instead.
//line /usr/local/go/src/crypto/x509/x509.go:1692
func ParseDERCRL(derBytes []byte) (*pkix.CertificateList, error) {
//line /usr/local/go/src/crypto/x509/x509.go:1692
	_go_fuzz_dep_.CoverTab[20427]++
							certList := new(pkix.CertificateList)
							if rest, err := asn1.Unmarshal(derBytes, certList); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1694
		_go_fuzz_dep_.CoverTab[20429]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1695
		// _ = "end of CoverTab[20429]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1696
		_go_fuzz_dep_.CoverTab[20430]++
//line /usr/local/go/src/crypto/x509/x509.go:1696
		if len(rest) != 0 {
//line /usr/local/go/src/crypto/x509/x509.go:1696
			_go_fuzz_dep_.CoverTab[20431]++
									return nil, errors.New("x509: trailing data after CRL")
//line /usr/local/go/src/crypto/x509/x509.go:1697
			// _ = "end of CoverTab[20431]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1698
			_go_fuzz_dep_.CoverTab[20432]++
//line /usr/local/go/src/crypto/x509/x509.go:1698
			// _ = "end of CoverTab[20432]"
//line /usr/local/go/src/crypto/x509/x509.go:1698
		}
//line /usr/local/go/src/crypto/x509/x509.go:1698
		// _ = "end of CoverTab[20430]"
//line /usr/local/go/src/crypto/x509/x509.go:1698
	}
//line /usr/local/go/src/crypto/x509/x509.go:1698
	// _ = "end of CoverTab[20427]"
//line /usr/local/go/src/crypto/x509/x509.go:1698
	_go_fuzz_dep_.CoverTab[20428]++
							return certList, nil
//line /usr/local/go/src/crypto/x509/x509.go:1699
	// _ = "end of CoverTab[20428]"
}

// CreateCRL returns a DER encoded CRL, signed by this Certificate, that
//line /usr/local/go/src/crypto/x509/x509.go:1702
// contains the given list of revoked certificates.
//line /usr/local/go/src/crypto/x509/x509.go:1702
//
//line /usr/local/go/src/crypto/x509/x509.go:1702
// Deprecated: this method does not generate an RFC 5280 conformant X.509 v2 CRL.
//line /usr/local/go/src/crypto/x509/x509.go:1702
// To generate a standards compliant CRL, use CreateRevocationList instead.
//line /usr/local/go/src/crypto/x509/x509.go:1707
func (c *Certificate) CreateCRL(rand io.Reader, priv any, revokedCerts []pkix.RevokedCertificate, now, expiry time.Time) (crlBytes []byte, err error) {
//line /usr/local/go/src/crypto/x509/x509.go:1707
	_go_fuzz_dep_.CoverTab[20433]++
							key, ok := priv.(crypto.Signer)
							if !ok {
//line /usr/local/go/src/crypto/x509/x509.go:1709
		_go_fuzz_dep_.CoverTab[20441]++
								return nil, errors.New("x509: certificate private key does not implement crypto.Signer")
//line /usr/local/go/src/crypto/x509/x509.go:1710
		// _ = "end of CoverTab[20441]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1711
		_go_fuzz_dep_.CoverTab[20442]++
//line /usr/local/go/src/crypto/x509/x509.go:1711
		// _ = "end of CoverTab[20442]"
//line /usr/local/go/src/crypto/x509/x509.go:1711
	}
//line /usr/local/go/src/crypto/x509/x509.go:1711
	// _ = "end of CoverTab[20433]"
//line /usr/local/go/src/crypto/x509/x509.go:1711
	_go_fuzz_dep_.CoverTab[20434]++

							hashFunc, signatureAlgorithm, err := signingParamsForPublicKey(key.Public(), 0)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1714
		_go_fuzz_dep_.CoverTab[20443]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1715
		// _ = "end of CoverTab[20443]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1716
		_go_fuzz_dep_.CoverTab[20444]++
//line /usr/local/go/src/crypto/x509/x509.go:1716
		// _ = "end of CoverTab[20444]"
//line /usr/local/go/src/crypto/x509/x509.go:1716
	}
//line /usr/local/go/src/crypto/x509/x509.go:1716
	// _ = "end of CoverTab[20434]"
//line /usr/local/go/src/crypto/x509/x509.go:1716
	_go_fuzz_dep_.CoverTab[20435]++

//line /usr/local/go/src/crypto/x509/x509.go:1719
	revokedCertsUTC := make([]pkix.RevokedCertificate, len(revokedCerts))
	for i, rc := range revokedCerts {
//line /usr/local/go/src/crypto/x509/x509.go:1720
		_go_fuzz_dep_.CoverTab[20445]++
								rc.RevocationTime = rc.RevocationTime.UTC()
								revokedCertsUTC[i] = rc
//line /usr/local/go/src/crypto/x509/x509.go:1722
		// _ = "end of CoverTab[20445]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:1723
	// _ = "end of CoverTab[20435]"
//line /usr/local/go/src/crypto/x509/x509.go:1723
	_go_fuzz_dep_.CoverTab[20436]++

							tbsCertList := pkix.TBSCertificateList{
		Version:		1,
		Signature:		signatureAlgorithm,
		Issuer:			c.Subject.ToRDNSequence(),
		ThisUpdate:		now.UTC(),
		NextUpdate:		expiry.UTC(),
		RevokedCertificates:	revokedCertsUTC,
	}

//line /usr/local/go/src/crypto/x509/x509.go:1735
	if len(c.SubjectKeyId) > 0 {
//line /usr/local/go/src/crypto/x509/x509.go:1735
		_go_fuzz_dep_.CoverTab[20446]++
								var aki pkix.Extension
								aki.Id = oidExtensionAuthorityKeyId
								aki.Value, err = asn1.Marshal(authKeyId{Id: c.SubjectKeyId})
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1739
			_go_fuzz_dep_.CoverTab[20448]++
									return
//line /usr/local/go/src/crypto/x509/x509.go:1740
			// _ = "end of CoverTab[20448]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1741
			_go_fuzz_dep_.CoverTab[20449]++
//line /usr/local/go/src/crypto/x509/x509.go:1741
			// _ = "end of CoverTab[20449]"
//line /usr/local/go/src/crypto/x509/x509.go:1741
		}
//line /usr/local/go/src/crypto/x509/x509.go:1741
		// _ = "end of CoverTab[20446]"
//line /usr/local/go/src/crypto/x509/x509.go:1741
		_go_fuzz_dep_.CoverTab[20447]++
								tbsCertList.Extensions = append(tbsCertList.Extensions, aki)
//line /usr/local/go/src/crypto/x509/x509.go:1742
		// _ = "end of CoverTab[20447]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1743
		_go_fuzz_dep_.CoverTab[20450]++
//line /usr/local/go/src/crypto/x509/x509.go:1743
		// _ = "end of CoverTab[20450]"
//line /usr/local/go/src/crypto/x509/x509.go:1743
	}
//line /usr/local/go/src/crypto/x509/x509.go:1743
	// _ = "end of CoverTab[20436]"
//line /usr/local/go/src/crypto/x509/x509.go:1743
	_go_fuzz_dep_.CoverTab[20437]++

							tbsCertListContents, err := asn1.Marshal(tbsCertList)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1746
		_go_fuzz_dep_.CoverTab[20451]++
								return
//line /usr/local/go/src/crypto/x509/x509.go:1747
		// _ = "end of CoverTab[20451]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1748
		_go_fuzz_dep_.CoverTab[20452]++
//line /usr/local/go/src/crypto/x509/x509.go:1748
		// _ = "end of CoverTab[20452]"
//line /usr/local/go/src/crypto/x509/x509.go:1748
	}
//line /usr/local/go/src/crypto/x509/x509.go:1748
	// _ = "end of CoverTab[20437]"
//line /usr/local/go/src/crypto/x509/x509.go:1748
	_go_fuzz_dep_.CoverTab[20438]++

							signed := tbsCertListContents
							if hashFunc != 0 {
//line /usr/local/go/src/crypto/x509/x509.go:1751
		_go_fuzz_dep_.CoverTab[20453]++
								h := hashFunc.New()
								h.Write(signed)
								signed = h.Sum(nil)
//line /usr/local/go/src/crypto/x509/x509.go:1754
		// _ = "end of CoverTab[20453]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1755
		_go_fuzz_dep_.CoverTab[20454]++
//line /usr/local/go/src/crypto/x509/x509.go:1755
		// _ = "end of CoverTab[20454]"
//line /usr/local/go/src/crypto/x509/x509.go:1755
	}
//line /usr/local/go/src/crypto/x509/x509.go:1755
	// _ = "end of CoverTab[20438]"
//line /usr/local/go/src/crypto/x509/x509.go:1755
	_go_fuzz_dep_.CoverTab[20439]++

							var signature []byte
							signature, err = key.Sign(rand, signed, hashFunc)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1759
		_go_fuzz_dep_.CoverTab[20455]++
								return
//line /usr/local/go/src/crypto/x509/x509.go:1760
		// _ = "end of CoverTab[20455]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1761
		_go_fuzz_dep_.CoverTab[20456]++
//line /usr/local/go/src/crypto/x509/x509.go:1761
		// _ = "end of CoverTab[20456]"
//line /usr/local/go/src/crypto/x509/x509.go:1761
	}
//line /usr/local/go/src/crypto/x509/x509.go:1761
	// _ = "end of CoverTab[20439]"
//line /usr/local/go/src/crypto/x509/x509.go:1761
	_go_fuzz_dep_.CoverTab[20440]++

							return asn1.Marshal(pkix.CertificateList{
		TBSCertList:		tbsCertList,
		SignatureAlgorithm:	signatureAlgorithm,
		SignatureValue:		asn1.BitString{Bytes: signature, BitLength: len(signature) * 8},
	})
//line /usr/local/go/src/crypto/x509/x509.go:1767
	// _ = "end of CoverTab[20440]"
}

// CertificateRequest represents a PKCS #10, certificate signature request.
type CertificateRequest struct {
	Raw				[]byte	// Complete ASN.1 DER content (CSR, signature algorithm and signature).
	RawTBSCertificateRequest	[]byte	// Certificate request info part of raw ASN.1 DER content.
	RawSubjectPublicKeyInfo		[]byte	// DER encoded SubjectPublicKeyInfo.
	RawSubject			[]byte	// DER encoded Subject.

	Version			int
	Signature		[]byte
	SignatureAlgorithm	SignatureAlgorithm

	PublicKeyAlgorithm	PublicKeyAlgorithm
	PublicKey		any

	Subject	pkix.Name

	// Attributes contains the CSR attributes that can parse as
	// pkix.AttributeTypeAndValueSET.
	//
	// Deprecated: Use Extensions and ExtraExtensions instead for parsing and
	// generating the requestedExtensions attribute.
	Attributes	[]pkix.AttributeTypeAndValueSET

	// Extensions contains all requested extensions, in raw form. When parsing
	// CSRs, this can be used to extract extensions that are not parsed by this
	// package.
	Extensions	[]pkix.Extension

	// ExtraExtensions contains extensions to be copied, raw, into any CSR
	// marshaled by CreateCertificateRequest. Values override any extensions
	// that would otherwise be produced based on the other fields but are
	// overridden by any extensions specified in Attributes.
	//
	// The ExtraExtensions field is not populated by ParseCertificateRequest,
	// see Extensions instead.
	ExtraExtensions	[]pkix.Extension

	// Subject Alternate Name values.
	DNSNames	[]string
	EmailAddresses	[]string
	IPAddresses	[]net.IP
	URIs		[]*url.URL
}

//line /usr/local/go/src/crypto/x509/x509.go:1817
type tbsCertificateRequest struct {
	Raw		asn1.RawContent
	Version		int
	Subject		asn1.RawValue
	PublicKey	publicKeyInfo
	RawAttributes	[]asn1.RawValue	`asn1:"tag:0"`
}

type certificateRequest struct {
	Raw			asn1.RawContent
	TBSCSR			tbsCertificateRequest
	SignatureAlgorithm	pkix.AlgorithmIdentifier
	SignatureValue		asn1.BitString
}

// oidExtensionRequest is a PKCS #9 OBJECT IDENTIFIER that indicates requested
//line /usr/local/go/src/crypto/x509/x509.go:1832
// extensions in a CSR.
//line /usr/local/go/src/crypto/x509/x509.go:1834
var oidExtensionRequest = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 14}

// newRawAttributes converts AttributeTypeAndValueSETs from a template
//line /usr/local/go/src/crypto/x509/x509.go:1836
// CertificateRequest's Attributes into tbsCertificateRequest RawAttributes.
//line /usr/local/go/src/crypto/x509/x509.go:1838
func newRawAttributes(attributes []pkix.AttributeTypeAndValueSET) ([]asn1.RawValue, error) {
//line /usr/local/go/src/crypto/x509/x509.go:1838
	_go_fuzz_dep_.CoverTab[20457]++
							var rawAttributes []asn1.RawValue
							b, err := asn1.Marshal(attributes)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1841
		_go_fuzz_dep_.CoverTab[20461]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1842
		// _ = "end of CoverTab[20461]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1843
		_go_fuzz_dep_.CoverTab[20462]++
//line /usr/local/go/src/crypto/x509/x509.go:1843
		// _ = "end of CoverTab[20462]"
//line /usr/local/go/src/crypto/x509/x509.go:1843
	}
//line /usr/local/go/src/crypto/x509/x509.go:1843
	// _ = "end of CoverTab[20457]"
//line /usr/local/go/src/crypto/x509/x509.go:1843
	_go_fuzz_dep_.CoverTab[20458]++
							rest, err := asn1.Unmarshal(b, &rawAttributes)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1845
		_go_fuzz_dep_.CoverTab[20463]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1846
		// _ = "end of CoverTab[20463]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1847
		_go_fuzz_dep_.CoverTab[20464]++
//line /usr/local/go/src/crypto/x509/x509.go:1847
		// _ = "end of CoverTab[20464]"
//line /usr/local/go/src/crypto/x509/x509.go:1847
	}
//line /usr/local/go/src/crypto/x509/x509.go:1847
	// _ = "end of CoverTab[20458]"
//line /usr/local/go/src/crypto/x509/x509.go:1847
	_go_fuzz_dep_.CoverTab[20459]++
							if len(rest) != 0 {
//line /usr/local/go/src/crypto/x509/x509.go:1848
		_go_fuzz_dep_.CoverTab[20465]++
								return nil, errors.New("x509: failed to unmarshal raw CSR Attributes")
//line /usr/local/go/src/crypto/x509/x509.go:1849
		// _ = "end of CoverTab[20465]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1850
		_go_fuzz_dep_.CoverTab[20466]++
//line /usr/local/go/src/crypto/x509/x509.go:1850
		// _ = "end of CoverTab[20466]"
//line /usr/local/go/src/crypto/x509/x509.go:1850
	}
//line /usr/local/go/src/crypto/x509/x509.go:1850
	// _ = "end of CoverTab[20459]"
//line /usr/local/go/src/crypto/x509/x509.go:1850
	_go_fuzz_dep_.CoverTab[20460]++
							return rawAttributes, nil
//line /usr/local/go/src/crypto/x509/x509.go:1851
	// _ = "end of CoverTab[20460]"
}

// parseRawAttributes Unmarshals RawAttributes into AttributeTypeAndValueSETs.
func parseRawAttributes(rawAttributes []asn1.RawValue) []pkix.AttributeTypeAndValueSET {
//line /usr/local/go/src/crypto/x509/x509.go:1855
	_go_fuzz_dep_.CoverTab[20467]++
							var attributes []pkix.AttributeTypeAndValueSET
							for _, rawAttr := range rawAttributes {
//line /usr/local/go/src/crypto/x509/x509.go:1857
		_go_fuzz_dep_.CoverTab[20469]++
								var attr pkix.AttributeTypeAndValueSET
								rest, err := asn1.Unmarshal(rawAttr.FullBytes, &attr)

//line /usr/local/go/src/crypto/x509/x509.go:1862
		if err == nil && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1862
			_go_fuzz_dep_.CoverTab[20470]++
//line /usr/local/go/src/crypto/x509/x509.go:1862
			return len(rest) == 0
//line /usr/local/go/src/crypto/x509/x509.go:1862
			// _ = "end of CoverTab[20470]"
//line /usr/local/go/src/crypto/x509/x509.go:1862
		}() {
//line /usr/local/go/src/crypto/x509/x509.go:1862
			_go_fuzz_dep_.CoverTab[20471]++
									attributes = append(attributes, attr)
//line /usr/local/go/src/crypto/x509/x509.go:1863
			// _ = "end of CoverTab[20471]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1864
			_go_fuzz_dep_.CoverTab[20472]++
//line /usr/local/go/src/crypto/x509/x509.go:1864
			// _ = "end of CoverTab[20472]"
//line /usr/local/go/src/crypto/x509/x509.go:1864
		}
//line /usr/local/go/src/crypto/x509/x509.go:1864
		// _ = "end of CoverTab[20469]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:1865
	// _ = "end of CoverTab[20467]"
//line /usr/local/go/src/crypto/x509/x509.go:1865
	_go_fuzz_dep_.CoverTab[20468]++
							return attributes
//line /usr/local/go/src/crypto/x509/x509.go:1866
	// _ = "end of CoverTab[20468]"
}

// parseCSRExtensions parses the attributes from a CSR and extracts any
//line /usr/local/go/src/crypto/x509/x509.go:1869
// requested extensions.
//line /usr/local/go/src/crypto/x509/x509.go:1871
func parseCSRExtensions(rawAttributes []asn1.RawValue) ([]pkix.Extension, error) {
//line /usr/local/go/src/crypto/x509/x509.go:1871
	_go_fuzz_dep_.CoverTab[20473]++
	// pkcs10Attribute reflects the Attribute structure from RFC 2986, Section 4.1.
	type pkcs10Attribute struct {
		Id	asn1.ObjectIdentifier
		Values	[]asn1.RawValue	`asn1:"set"`
	}

	var ret []pkix.Extension
	requestedExts := make(map[string]bool)
	for _, rawAttr := range rawAttributes {
//line /usr/local/go/src/crypto/x509/x509.go:1880
		_go_fuzz_dep_.CoverTab[20475]++
								var attr pkcs10Attribute
								if rest, err := asn1.Unmarshal(rawAttr.FullBytes, &attr); err != nil || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1882
			_go_fuzz_dep_.CoverTab[20480]++
//line /usr/local/go/src/crypto/x509/x509.go:1882
			return len(rest) != 0
//line /usr/local/go/src/crypto/x509/x509.go:1882
			// _ = "end of CoverTab[20480]"
//line /usr/local/go/src/crypto/x509/x509.go:1882
		}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1882
			_go_fuzz_dep_.CoverTab[20481]++
//line /usr/local/go/src/crypto/x509/x509.go:1882
			return len(attr.Values) == 0
//line /usr/local/go/src/crypto/x509/x509.go:1882
			// _ = "end of CoverTab[20481]"
//line /usr/local/go/src/crypto/x509/x509.go:1882
		}() {
//line /usr/local/go/src/crypto/x509/x509.go:1882
			_go_fuzz_dep_.CoverTab[20482]++

									continue
//line /usr/local/go/src/crypto/x509/x509.go:1884
			// _ = "end of CoverTab[20482]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1885
			_go_fuzz_dep_.CoverTab[20483]++
//line /usr/local/go/src/crypto/x509/x509.go:1885
			// _ = "end of CoverTab[20483]"
//line /usr/local/go/src/crypto/x509/x509.go:1885
		}
//line /usr/local/go/src/crypto/x509/x509.go:1885
		// _ = "end of CoverTab[20475]"
//line /usr/local/go/src/crypto/x509/x509.go:1885
		_go_fuzz_dep_.CoverTab[20476]++

								if !attr.Id.Equal(oidExtensionRequest) {
//line /usr/local/go/src/crypto/x509/x509.go:1887
			_go_fuzz_dep_.CoverTab[20484]++
									continue
//line /usr/local/go/src/crypto/x509/x509.go:1888
			// _ = "end of CoverTab[20484]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1889
			_go_fuzz_dep_.CoverTab[20485]++
//line /usr/local/go/src/crypto/x509/x509.go:1889
			// _ = "end of CoverTab[20485]"
//line /usr/local/go/src/crypto/x509/x509.go:1889
		}
//line /usr/local/go/src/crypto/x509/x509.go:1889
		// _ = "end of CoverTab[20476]"
//line /usr/local/go/src/crypto/x509/x509.go:1889
		_go_fuzz_dep_.CoverTab[20477]++

								var extensions []pkix.Extension
								if _, err := asn1.Unmarshal(attr.Values[0].FullBytes, &extensions); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1892
			_go_fuzz_dep_.CoverTab[20486]++
									return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1893
			// _ = "end of CoverTab[20486]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:1894
			_go_fuzz_dep_.CoverTab[20487]++
//line /usr/local/go/src/crypto/x509/x509.go:1894
			// _ = "end of CoverTab[20487]"
//line /usr/local/go/src/crypto/x509/x509.go:1894
		}
//line /usr/local/go/src/crypto/x509/x509.go:1894
		// _ = "end of CoverTab[20477]"
//line /usr/local/go/src/crypto/x509/x509.go:1894
		_go_fuzz_dep_.CoverTab[20478]++
								for _, ext := range extensions {
//line /usr/local/go/src/crypto/x509/x509.go:1895
			_go_fuzz_dep_.CoverTab[20488]++
									oidStr := ext.Id.String()
									if requestedExts[oidStr] {
//line /usr/local/go/src/crypto/x509/x509.go:1897
				_go_fuzz_dep_.CoverTab[20490]++
										return nil, errors.New("x509: certificate request contains duplicate requested extensions")
//line /usr/local/go/src/crypto/x509/x509.go:1898
				// _ = "end of CoverTab[20490]"
			} else {
//line /usr/local/go/src/crypto/x509/x509.go:1899
				_go_fuzz_dep_.CoverTab[20491]++
//line /usr/local/go/src/crypto/x509/x509.go:1899
				// _ = "end of CoverTab[20491]"
//line /usr/local/go/src/crypto/x509/x509.go:1899
			}
//line /usr/local/go/src/crypto/x509/x509.go:1899
			// _ = "end of CoverTab[20488]"
//line /usr/local/go/src/crypto/x509/x509.go:1899
			_go_fuzz_dep_.CoverTab[20489]++
									requestedExts[oidStr] = true
//line /usr/local/go/src/crypto/x509/x509.go:1900
			// _ = "end of CoverTab[20489]"
		}
//line /usr/local/go/src/crypto/x509/x509.go:1901
		// _ = "end of CoverTab[20478]"
//line /usr/local/go/src/crypto/x509/x509.go:1901
		_go_fuzz_dep_.CoverTab[20479]++
								ret = append(ret, extensions...)
//line /usr/local/go/src/crypto/x509/x509.go:1902
		// _ = "end of CoverTab[20479]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:1903
	// _ = "end of CoverTab[20473]"
//line /usr/local/go/src/crypto/x509/x509.go:1903
	_go_fuzz_dep_.CoverTab[20474]++

							return ret, nil
//line /usr/local/go/src/crypto/x509/x509.go:1905
	// _ = "end of CoverTab[20474]"
}

// CreateCertificateRequest creates a new certificate request based on a
//line /usr/local/go/src/crypto/x509/x509.go:1908
// template. The following members of template are used:
//line /usr/local/go/src/crypto/x509/x509.go:1908
//
//line /usr/local/go/src/crypto/x509/x509.go:1908
//   - SignatureAlgorithm
//line /usr/local/go/src/crypto/x509/x509.go:1908
//   - Subject
//line /usr/local/go/src/crypto/x509/x509.go:1908
//   - DNSNames
//line /usr/local/go/src/crypto/x509/x509.go:1908
//   - EmailAddresses
//line /usr/local/go/src/crypto/x509/x509.go:1908
//   - IPAddresses
//line /usr/local/go/src/crypto/x509/x509.go:1908
//   - URIs
//line /usr/local/go/src/crypto/x509/x509.go:1908
//   - ExtraExtensions
//line /usr/local/go/src/crypto/x509/x509.go:1908
//   - Attributes (deprecated)
//line /usr/local/go/src/crypto/x509/x509.go:1908
//
//line /usr/local/go/src/crypto/x509/x509.go:1908
// priv is the private key to sign the CSR with, and the corresponding public
//line /usr/local/go/src/crypto/x509/x509.go:1908
// key will be included in the CSR. It must implement crypto.Signer and its
//line /usr/local/go/src/crypto/x509/x509.go:1908
// Public() method must return a *rsa.PublicKey or a *ecdsa.PublicKey or a
//line /usr/local/go/src/crypto/x509/x509.go:1908
// ed25519.PublicKey. (A *rsa.PrivateKey, *ecdsa.PrivateKey or
//line /usr/local/go/src/crypto/x509/x509.go:1908
// ed25519.PrivateKey satisfies this.)
//line /usr/local/go/src/crypto/x509/x509.go:1908
//
//line /usr/local/go/src/crypto/x509/x509.go:1908
// The returned slice is the certificate request in DER encoding.
//line /usr/local/go/src/crypto/x509/x509.go:1927
func CreateCertificateRequest(rand io.Reader, template *CertificateRequest, priv any) (csr []byte, err error) {
//line /usr/local/go/src/crypto/x509/x509.go:1927
	_go_fuzz_dep_.CoverTab[20492]++
							key, ok := priv.(crypto.Signer)
							if !ok {
//line /usr/local/go/src/crypto/x509/x509.go:1929
		_go_fuzz_dep_.CoverTab[20505]++
								return nil, errors.New("x509: certificate private key does not implement crypto.Signer")
//line /usr/local/go/src/crypto/x509/x509.go:1930
		// _ = "end of CoverTab[20505]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1931
		_go_fuzz_dep_.CoverTab[20506]++
//line /usr/local/go/src/crypto/x509/x509.go:1931
		// _ = "end of CoverTab[20506]"
//line /usr/local/go/src/crypto/x509/x509.go:1931
	}
//line /usr/local/go/src/crypto/x509/x509.go:1931
	// _ = "end of CoverTab[20492]"
//line /usr/local/go/src/crypto/x509/x509.go:1931
	_go_fuzz_dep_.CoverTab[20493]++

							var hashFunc crypto.Hash
							var sigAlgo pkix.AlgorithmIdentifier
							hashFunc, sigAlgo, err = signingParamsForPublicKey(key.Public(), template.SignatureAlgorithm)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1936
		_go_fuzz_dep_.CoverTab[20507]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1937
		// _ = "end of CoverTab[20507]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1938
		_go_fuzz_dep_.CoverTab[20508]++
//line /usr/local/go/src/crypto/x509/x509.go:1938
		// _ = "end of CoverTab[20508]"
//line /usr/local/go/src/crypto/x509/x509.go:1938
	}
//line /usr/local/go/src/crypto/x509/x509.go:1938
	// _ = "end of CoverTab[20493]"
//line /usr/local/go/src/crypto/x509/x509.go:1938
	_go_fuzz_dep_.CoverTab[20494]++

							var publicKeyBytes []byte
							var publicKeyAlgorithm pkix.AlgorithmIdentifier
							publicKeyBytes, publicKeyAlgorithm, err = marshalPublicKey(key.Public())
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1943
		_go_fuzz_dep_.CoverTab[20509]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1944
		// _ = "end of CoverTab[20509]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1945
		_go_fuzz_dep_.CoverTab[20510]++
//line /usr/local/go/src/crypto/x509/x509.go:1945
		// _ = "end of CoverTab[20510]"
//line /usr/local/go/src/crypto/x509/x509.go:1945
	}
//line /usr/local/go/src/crypto/x509/x509.go:1945
	// _ = "end of CoverTab[20494]"
//line /usr/local/go/src/crypto/x509/x509.go:1945
	_go_fuzz_dep_.CoverTab[20495]++

							extensions, err := buildCSRExtensions(template)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:1948
		_go_fuzz_dep_.CoverTab[20511]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:1949
		// _ = "end of CoverTab[20511]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:1950
		_go_fuzz_dep_.CoverTab[20512]++
//line /usr/local/go/src/crypto/x509/x509.go:1950
		// _ = "end of CoverTab[20512]"
//line /usr/local/go/src/crypto/x509/x509.go:1950
	}
//line /usr/local/go/src/crypto/x509/x509.go:1950
	// _ = "end of CoverTab[20495]"
//line /usr/local/go/src/crypto/x509/x509.go:1950
	_go_fuzz_dep_.CoverTab[20496]++

//line /usr/local/go/src/crypto/x509/x509.go:1953
	attributes := make([]pkix.AttributeTypeAndValueSET, 0, len(template.Attributes))
	for _, attr := range template.Attributes {
//line /usr/local/go/src/crypto/x509/x509.go:1954
		_go_fuzz_dep_.CoverTab[20513]++
								values := make([][]pkix.AttributeTypeAndValue, len(attr.Value))
								copy(values, attr.Value)
								attributes = append(attributes, pkix.AttributeTypeAndValueSET{
			Type:	attr.Type,
			Value:	values,
		})
//line /usr/local/go/src/crypto/x509/x509.go:1960
		// _ = "end of CoverTab[20513]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:1961
	// _ = "end of CoverTab[20496]"
//line /usr/local/go/src/crypto/x509/x509.go:1961
	_go_fuzz_dep_.CoverTab[20497]++

							extensionsAppended := false
							if len(extensions) > 0 {
//line /usr/local/go/src/crypto/x509/x509.go:1964
		_go_fuzz_dep_.CoverTab[20514]++

								for _, atvSet := range attributes {
//line /usr/local/go/src/crypto/x509/x509.go:1966
			_go_fuzz_dep_.CoverTab[20515]++
									if !atvSet.Type.Equal(oidExtensionRequest) || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:1967
				_go_fuzz_dep_.CoverTab[20519]++
//line /usr/local/go/src/crypto/x509/x509.go:1967
				return len(atvSet.Value) == 0
//line /usr/local/go/src/crypto/x509/x509.go:1967
				// _ = "end of CoverTab[20519]"
//line /usr/local/go/src/crypto/x509/x509.go:1967
			}() {
//line /usr/local/go/src/crypto/x509/x509.go:1967
				_go_fuzz_dep_.CoverTab[20520]++
										continue
//line /usr/local/go/src/crypto/x509/x509.go:1968
				// _ = "end of CoverTab[20520]"
			} else {
//line /usr/local/go/src/crypto/x509/x509.go:1969
				_go_fuzz_dep_.CoverTab[20521]++
//line /usr/local/go/src/crypto/x509/x509.go:1969
				// _ = "end of CoverTab[20521]"
//line /usr/local/go/src/crypto/x509/x509.go:1969
			}
//line /usr/local/go/src/crypto/x509/x509.go:1969
			// _ = "end of CoverTab[20515]"
//line /usr/local/go/src/crypto/x509/x509.go:1969
			_go_fuzz_dep_.CoverTab[20516]++

//line /usr/local/go/src/crypto/x509/x509.go:1973
			specifiedExtensions := make(map[string]bool)

			for _, atvs := range atvSet.Value {
//line /usr/local/go/src/crypto/x509/x509.go:1975
				_go_fuzz_dep_.CoverTab[20522]++
										for _, atv := range atvs {
//line /usr/local/go/src/crypto/x509/x509.go:1976
					_go_fuzz_dep_.CoverTab[20523]++
											specifiedExtensions[atv.Type.String()] = true
//line /usr/local/go/src/crypto/x509/x509.go:1977
					// _ = "end of CoverTab[20523]"
				}
//line /usr/local/go/src/crypto/x509/x509.go:1978
				// _ = "end of CoverTab[20522]"
			}
//line /usr/local/go/src/crypto/x509/x509.go:1979
			// _ = "end of CoverTab[20516]"
//line /usr/local/go/src/crypto/x509/x509.go:1979
			_go_fuzz_dep_.CoverTab[20517]++

									newValue := make([]pkix.AttributeTypeAndValue, 0, len(atvSet.Value[0])+len(extensions))
									newValue = append(newValue, atvSet.Value[0]...)

									for _, e := range extensions {
//line /usr/local/go/src/crypto/x509/x509.go:1984
				_go_fuzz_dep_.CoverTab[20524]++
										if specifiedExtensions[e.Id.String()] {
//line /usr/local/go/src/crypto/x509/x509.go:1985
					_go_fuzz_dep_.CoverTab[20526]++

//line /usr/local/go/src/crypto/x509/x509.go:1988
					continue
//line /usr/local/go/src/crypto/x509/x509.go:1988
					// _ = "end of CoverTab[20526]"
				} else {
//line /usr/local/go/src/crypto/x509/x509.go:1989
					_go_fuzz_dep_.CoverTab[20527]++
//line /usr/local/go/src/crypto/x509/x509.go:1989
					// _ = "end of CoverTab[20527]"
//line /usr/local/go/src/crypto/x509/x509.go:1989
				}
//line /usr/local/go/src/crypto/x509/x509.go:1989
				// _ = "end of CoverTab[20524]"
//line /usr/local/go/src/crypto/x509/x509.go:1989
				_go_fuzz_dep_.CoverTab[20525]++

										newValue = append(newValue, pkix.AttributeTypeAndValue{

//line /usr/local/go/src/crypto/x509/x509.go:1994
					Type:	e.Id,
											Value:	e.Value,
				})
//line /usr/local/go/src/crypto/x509/x509.go:1996
				// _ = "end of CoverTab[20525]"
			}
//line /usr/local/go/src/crypto/x509/x509.go:1997
			// _ = "end of CoverTab[20517]"
//line /usr/local/go/src/crypto/x509/x509.go:1997
			_go_fuzz_dep_.CoverTab[20518]++

									atvSet.Value[0] = newValue
									extensionsAppended = true
									break
//line /usr/local/go/src/crypto/x509/x509.go:2001
			// _ = "end of CoverTab[20518]"
		}
//line /usr/local/go/src/crypto/x509/x509.go:2002
		// _ = "end of CoverTab[20514]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2003
		_go_fuzz_dep_.CoverTab[20528]++
//line /usr/local/go/src/crypto/x509/x509.go:2003
		// _ = "end of CoverTab[20528]"
//line /usr/local/go/src/crypto/x509/x509.go:2003
	}
//line /usr/local/go/src/crypto/x509/x509.go:2003
	// _ = "end of CoverTab[20497]"
//line /usr/local/go/src/crypto/x509/x509.go:2003
	_go_fuzz_dep_.CoverTab[20498]++

							rawAttributes, err := newRawAttributes(attributes)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2006
		_go_fuzz_dep_.CoverTab[20529]++
								return
//line /usr/local/go/src/crypto/x509/x509.go:2007
		// _ = "end of CoverTab[20529]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2008
		_go_fuzz_dep_.CoverTab[20530]++
//line /usr/local/go/src/crypto/x509/x509.go:2008
		// _ = "end of CoverTab[20530]"
//line /usr/local/go/src/crypto/x509/x509.go:2008
	}
//line /usr/local/go/src/crypto/x509/x509.go:2008
	// _ = "end of CoverTab[20498]"
//line /usr/local/go/src/crypto/x509/x509.go:2008
	_go_fuzz_dep_.CoverTab[20499]++

//line /usr/local/go/src/crypto/x509/x509.go:2012
	if len(extensions) > 0 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:2012
		_go_fuzz_dep_.CoverTab[20531]++
//line /usr/local/go/src/crypto/x509/x509.go:2012
		return !extensionsAppended
//line /usr/local/go/src/crypto/x509/x509.go:2012
		// _ = "end of CoverTab[20531]"
//line /usr/local/go/src/crypto/x509/x509.go:2012
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:2012
		_go_fuzz_dep_.CoverTab[20532]++
								attr := struct {
			Type	asn1.ObjectIdentifier
			Value	[][]pkix.Extension	`asn1:"set"`
		}{
			Type:	oidExtensionRequest,
			Value:	[][]pkix.Extension{extensions},
		}

		b, err := asn1.Marshal(attr)
		if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2022
			_go_fuzz_dep_.CoverTab[20535]++
									return nil, errors.New("x509: failed to serialise extensions attribute: " + err.Error())
//line /usr/local/go/src/crypto/x509/x509.go:2023
			// _ = "end of CoverTab[20535]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:2024
			_go_fuzz_dep_.CoverTab[20536]++
//line /usr/local/go/src/crypto/x509/x509.go:2024
			// _ = "end of CoverTab[20536]"
//line /usr/local/go/src/crypto/x509/x509.go:2024
		}
//line /usr/local/go/src/crypto/x509/x509.go:2024
		// _ = "end of CoverTab[20532]"
//line /usr/local/go/src/crypto/x509/x509.go:2024
		_go_fuzz_dep_.CoverTab[20533]++

								var rawValue asn1.RawValue
								if _, err := asn1.Unmarshal(b, &rawValue); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2027
			_go_fuzz_dep_.CoverTab[20537]++
									return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:2028
			// _ = "end of CoverTab[20537]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:2029
			_go_fuzz_dep_.CoverTab[20538]++
//line /usr/local/go/src/crypto/x509/x509.go:2029
			// _ = "end of CoverTab[20538]"
//line /usr/local/go/src/crypto/x509/x509.go:2029
		}
//line /usr/local/go/src/crypto/x509/x509.go:2029
		// _ = "end of CoverTab[20533]"
//line /usr/local/go/src/crypto/x509/x509.go:2029
		_go_fuzz_dep_.CoverTab[20534]++

								rawAttributes = append(rawAttributes, rawValue)
//line /usr/local/go/src/crypto/x509/x509.go:2031
		// _ = "end of CoverTab[20534]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2032
		_go_fuzz_dep_.CoverTab[20539]++
//line /usr/local/go/src/crypto/x509/x509.go:2032
		// _ = "end of CoverTab[20539]"
//line /usr/local/go/src/crypto/x509/x509.go:2032
	}
//line /usr/local/go/src/crypto/x509/x509.go:2032
	// _ = "end of CoverTab[20499]"
//line /usr/local/go/src/crypto/x509/x509.go:2032
	_go_fuzz_dep_.CoverTab[20500]++

							asn1Subject := template.RawSubject
							if len(asn1Subject) == 0 {
//line /usr/local/go/src/crypto/x509/x509.go:2035
		_go_fuzz_dep_.CoverTab[20540]++
								asn1Subject, err = asn1.Marshal(template.Subject.ToRDNSequence())
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2037
			_go_fuzz_dep_.CoverTab[20541]++
									return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:2038
			// _ = "end of CoverTab[20541]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:2039
			_go_fuzz_dep_.CoverTab[20542]++
//line /usr/local/go/src/crypto/x509/x509.go:2039
			// _ = "end of CoverTab[20542]"
//line /usr/local/go/src/crypto/x509/x509.go:2039
		}
//line /usr/local/go/src/crypto/x509/x509.go:2039
		// _ = "end of CoverTab[20540]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2040
		_go_fuzz_dep_.CoverTab[20543]++
//line /usr/local/go/src/crypto/x509/x509.go:2040
		// _ = "end of CoverTab[20543]"
//line /usr/local/go/src/crypto/x509/x509.go:2040
	}
//line /usr/local/go/src/crypto/x509/x509.go:2040
	// _ = "end of CoverTab[20500]"
//line /usr/local/go/src/crypto/x509/x509.go:2040
	_go_fuzz_dep_.CoverTab[20501]++

							tbsCSR := tbsCertificateRequest{
		Version:	0,
		Subject:	asn1.RawValue{FullBytes: asn1Subject},
		PublicKey: publicKeyInfo{
			Algorithm:	publicKeyAlgorithm,
			PublicKey: asn1.BitString{
				Bytes:		publicKeyBytes,
				BitLength:	len(publicKeyBytes) * 8,
			},
		},
		RawAttributes:	rawAttributes,
	}

	tbsCSRContents, err := asn1.Marshal(tbsCSR)
	if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2056
		_go_fuzz_dep_.CoverTab[20544]++
								return
//line /usr/local/go/src/crypto/x509/x509.go:2057
		// _ = "end of CoverTab[20544]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2058
		_go_fuzz_dep_.CoverTab[20545]++
//line /usr/local/go/src/crypto/x509/x509.go:2058
		// _ = "end of CoverTab[20545]"
//line /usr/local/go/src/crypto/x509/x509.go:2058
	}
//line /usr/local/go/src/crypto/x509/x509.go:2058
	// _ = "end of CoverTab[20501]"
//line /usr/local/go/src/crypto/x509/x509.go:2058
	_go_fuzz_dep_.CoverTab[20502]++
							tbsCSR.Raw = tbsCSRContents

							signed := tbsCSRContents
							if hashFunc != 0 {
//line /usr/local/go/src/crypto/x509/x509.go:2062
		_go_fuzz_dep_.CoverTab[20546]++
								h := hashFunc.New()
								h.Write(signed)
								signed = h.Sum(nil)
//line /usr/local/go/src/crypto/x509/x509.go:2065
		// _ = "end of CoverTab[20546]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2066
		_go_fuzz_dep_.CoverTab[20547]++
//line /usr/local/go/src/crypto/x509/x509.go:2066
		// _ = "end of CoverTab[20547]"
//line /usr/local/go/src/crypto/x509/x509.go:2066
	}
//line /usr/local/go/src/crypto/x509/x509.go:2066
	// _ = "end of CoverTab[20502]"
//line /usr/local/go/src/crypto/x509/x509.go:2066
	_go_fuzz_dep_.CoverTab[20503]++

							var signature []byte
							signature, err = key.Sign(rand, signed, hashFunc)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2070
		_go_fuzz_dep_.CoverTab[20548]++
								return
//line /usr/local/go/src/crypto/x509/x509.go:2071
		// _ = "end of CoverTab[20548]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2072
		_go_fuzz_dep_.CoverTab[20549]++
//line /usr/local/go/src/crypto/x509/x509.go:2072
		// _ = "end of CoverTab[20549]"
//line /usr/local/go/src/crypto/x509/x509.go:2072
	}
//line /usr/local/go/src/crypto/x509/x509.go:2072
	// _ = "end of CoverTab[20503]"
//line /usr/local/go/src/crypto/x509/x509.go:2072
	_go_fuzz_dep_.CoverTab[20504]++

							return asn1.Marshal(certificateRequest{
		TBSCSR:			tbsCSR,
		SignatureAlgorithm:	sigAlgo,
		SignatureValue: asn1.BitString{
			Bytes:		signature,
			BitLength:	len(signature) * 8,
		},
	})
//line /usr/local/go/src/crypto/x509/x509.go:2081
	// _ = "end of CoverTab[20504]"
}

// ParseCertificateRequest parses a single certificate request from the
//line /usr/local/go/src/crypto/x509/x509.go:2084
// given ASN.1 DER data.
//line /usr/local/go/src/crypto/x509/x509.go:2086
func ParseCertificateRequest(asn1Data []byte) (*CertificateRequest, error) {
//line /usr/local/go/src/crypto/x509/x509.go:2086
	_go_fuzz_dep_.CoverTab[20550]++
							var csr certificateRequest

							rest, err := asn1.Unmarshal(asn1Data, &csr)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2090
		_go_fuzz_dep_.CoverTab[20552]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:2091
		// _ = "end of CoverTab[20552]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2092
		_go_fuzz_dep_.CoverTab[20553]++
//line /usr/local/go/src/crypto/x509/x509.go:2092
		if len(rest) != 0 {
//line /usr/local/go/src/crypto/x509/x509.go:2092
			_go_fuzz_dep_.CoverTab[20554]++
									return nil, asn1.SyntaxError{Msg: "trailing data"}
//line /usr/local/go/src/crypto/x509/x509.go:2093
			// _ = "end of CoverTab[20554]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:2094
			_go_fuzz_dep_.CoverTab[20555]++
//line /usr/local/go/src/crypto/x509/x509.go:2094
			// _ = "end of CoverTab[20555]"
//line /usr/local/go/src/crypto/x509/x509.go:2094
		}
//line /usr/local/go/src/crypto/x509/x509.go:2094
		// _ = "end of CoverTab[20553]"
//line /usr/local/go/src/crypto/x509/x509.go:2094
	}
//line /usr/local/go/src/crypto/x509/x509.go:2094
	// _ = "end of CoverTab[20550]"
//line /usr/local/go/src/crypto/x509/x509.go:2094
	_go_fuzz_dep_.CoverTab[20551]++

							return parseCertificateRequest(&csr)
//line /usr/local/go/src/crypto/x509/x509.go:2096
	// _ = "end of CoverTab[20551]"
}

func parseCertificateRequest(in *certificateRequest) (*CertificateRequest, error) {
//line /usr/local/go/src/crypto/x509/x509.go:2099
	_go_fuzz_dep_.CoverTab[20556]++
							out := &CertificateRequest{
		Raw:				in.Raw,
		RawTBSCertificateRequest:	in.TBSCSR.Raw,
		RawSubjectPublicKeyInfo:	in.TBSCSR.PublicKey.Raw,
		RawSubject:			in.TBSCSR.Subject.FullBytes,

		Signature:		in.SignatureValue.RightAlign(),
		SignatureAlgorithm:	getSignatureAlgorithmFromAI(in.SignatureAlgorithm),

		PublicKeyAlgorithm:	getPublicKeyAlgorithmFromOID(in.TBSCSR.PublicKey.Algorithm.Algorithm),

		Version:	in.TBSCSR.Version,
		Attributes:	parseRawAttributes(in.TBSCSR.RawAttributes),
	}

	var err error
	if out.PublicKeyAlgorithm != UnknownPublicKeyAlgorithm {
//line /usr/local/go/src/crypto/x509/x509.go:2116
		_go_fuzz_dep_.CoverTab[20561]++
								out.PublicKey, err = parsePublicKey(&in.TBSCSR.PublicKey)
								if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2118
			_go_fuzz_dep_.CoverTab[20562]++
									return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:2119
			// _ = "end of CoverTab[20562]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:2120
			_go_fuzz_dep_.CoverTab[20563]++
//line /usr/local/go/src/crypto/x509/x509.go:2120
			// _ = "end of CoverTab[20563]"
//line /usr/local/go/src/crypto/x509/x509.go:2120
		}
//line /usr/local/go/src/crypto/x509/x509.go:2120
		// _ = "end of CoverTab[20561]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2121
		_go_fuzz_dep_.CoverTab[20564]++
//line /usr/local/go/src/crypto/x509/x509.go:2121
		// _ = "end of CoverTab[20564]"
//line /usr/local/go/src/crypto/x509/x509.go:2121
	}
//line /usr/local/go/src/crypto/x509/x509.go:2121
	// _ = "end of CoverTab[20556]"
//line /usr/local/go/src/crypto/x509/x509.go:2121
	_go_fuzz_dep_.CoverTab[20557]++

							var subject pkix.RDNSequence
							if rest, err := asn1.Unmarshal(in.TBSCSR.Subject.FullBytes, &subject); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2124
		_go_fuzz_dep_.CoverTab[20565]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:2125
		// _ = "end of CoverTab[20565]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2126
		_go_fuzz_dep_.CoverTab[20566]++
//line /usr/local/go/src/crypto/x509/x509.go:2126
		if len(rest) != 0 {
//line /usr/local/go/src/crypto/x509/x509.go:2126
			_go_fuzz_dep_.CoverTab[20567]++
									return nil, errors.New("x509: trailing data after X.509 Subject")
//line /usr/local/go/src/crypto/x509/x509.go:2127
			// _ = "end of CoverTab[20567]"
		} else {
//line /usr/local/go/src/crypto/x509/x509.go:2128
			_go_fuzz_dep_.CoverTab[20568]++
//line /usr/local/go/src/crypto/x509/x509.go:2128
			// _ = "end of CoverTab[20568]"
//line /usr/local/go/src/crypto/x509/x509.go:2128
		}
//line /usr/local/go/src/crypto/x509/x509.go:2128
		// _ = "end of CoverTab[20566]"
//line /usr/local/go/src/crypto/x509/x509.go:2128
	}
//line /usr/local/go/src/crypto/x509/x509.go:2128
	// _ = "end of CoverTab[20557]"
//line /usr/local/go/src/crypto/x509/x509.go:2128
	_go_fuzz_dep_.CoverTab[20558]++

							out.Subject.FillFromRDNSequence(&subject)

							if out.Extensions, err = parseCSRExtensions(in.TBSCSR.RawAttributes); err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2132
		_go_fuzz_dep_.CoverTab[20569]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:2133
		// _ = "end of CoverTab[20569]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2134
		_go_fuzz_dep_.CoverTab[20570]++
//line /usr/local/go/src/crypto/x509/x509.go:2134
		// _ = "end of CoverTab[20570]"
//line /usr/local/go/src/crypto/x509/x509.go:2134
	}
//line /usr/local/go/src/crypto/x509/x509.go:2134
	// _ = "end of CoverTab[20558]"
//line /usr/local/go/src/crypto/x509/x509.go:2134
	_go_fuzz_dep_.CoverTab[20559]++

							for _, extension := range out.Extensions {
//line /usr/local/go/src/crypto/x509/x509.go:2136
		_go_fuzz_dep_.CoverTab[20571]++
								switch {
		case extension.Id.Equal(oidExtensionSubjectAltName):
//line /usr/local/go/src/crypto/x509/x509.go:2138
			_go_fuzz_dep_.CoverTab[20572]++
									out.DNSNames, out.EmailAddresses, out.IPAddresses, out.URIs, err = parseSANExtension(extension.Value)
									if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2140
				_go_fuzz_dep_.CoverTab[20574]++
										return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:2141
				// _ = "end of CoverTab[20574]"
			} else {
//line /usr/local/go/src/crypto/x509/x509.go:2142
				_go_fuzz_dep_.CoverTab[20575]++
//line /usr/local/go/src/crypto/x509/x509.go:2142
				// _ = "end of CoverTab[20575]"
//line /usr/local/go/src/crypto/x509/x509.go:2142
			}
//line /usr/local/go/src/crypto/x509/x509.go:2142
			// _ = "end of CoverTab[20572]"
//line /usr/local/go/src/crypto/x509/x509.go:2142
		default:
//line /usr/local/go/src/crypto/x509/x509.go:2142
			_go_fuzz_dep_.CoverTab[20573]++
//line /usr/local/go/src/crypto/x509/x509.go:2142
			// _ = "end of CoverTab[20573]"
		}
//line /usr/local/go/src/crypto/x509/x509.go:2143
		// _ = "end of CoverTab[20571]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:2144
	// _ = "end of CoverTab[20559]"
//line /usr/local/go/src/crypto/x509/x509.go:2144
	_go_fuzz_dep_.CoverTab[20560]++

							return out, nil
//line /usr/local/go/src/crypto/x509/x509.go:2146
	// _ = "end of CoverTab[20560]"
}

// CheckSignature reports whether the signature on c is valid.
func (c *CertificateRequest) CheckSignature() error {
//line /usr/local/go/src/crypto/x509/x509.go:2150
	_go_fuzz_dep_.CoverTab[20576]++
							return checkSignature(c.SignatureAlgorithm, c.RawTBSCertificateRequest, c.Signature, c.PublicKey, true)
//line /usr/local/go/src/crypto/x509/x509.go:2151
	// _ = "end of CoverTab[20576]"
}

// RevocationList contains the fields used to create an X.509 v2 Certificate
//line /usr/local/go/src/crypto/x509/x509.go:2154
// Revocation list with CreateRevocationList.
//line /usr/local/go/src/crypto/x509/x509.go:2156
type RevocationList struct {
	// Raw contains the complete ASN.1 DER content of the CRL (tbsCertList,
	// signatureAlgorithm, and signatureValue.)
	Raw	[]byte
	// RawTBSRevocationList contains just the tbsCertList portion of the ASN.1
	// DER.
	RawTBSRevocationList	[]byte
	// RawIssuer contains the DER encoded Issuer.
	RawIssuer	[]byte

	// Issuer contains the DN of the issuing certificate.
	Issuer	pkix.Name
	// AuthorityKeyId is used to identify the public key associated with the
	// issuing certificate. It is populated from the authorityKeyIdentifier
	// extension when parsing a CRL. It is ignored when creating a CRL; the
	// extension is populated from the issuing certificate itself.
	AuthorityKeyId	[]byte

	Signature	[]byte
	// SignatureAlgorithm is used to determine the signature algorithm to be
	// used when signing the CRL. If 0 the default algorithm for the signing
	// key will be used.
	SignatureAlgorithm	SignatureAlgorithm

	// RevokedCertificates is used to populate the revokedCertificates
	// sequence in the CRL, it may be empty. RevokedCertificates may be nil,
	// in which case an empty CRL will be created.
	RevokedCertificates	[]pkix.RevokedCertificate

	// Number is used to populate the X.509 v2 cRLNumber extension in the CRL,
	// which should be a monotonically increasing sequence number for a given
	// CRL scope and CRL issuer. It is also populated from the cRLNumber
	// extension when parsing a CRL.
	Number	*big.Int

	// ThisUpdate is used to populate the thisUpdate field in the CRL, which
	// indicates the issuance date of the CRL.
	ThisUpdate	time.Time
	// NextUpdate is used to populate the nextUpdate field in the CRL, which
	// indicates the date by which the next CRL will be issued. NextUpdate
	// must be greater than ThisUpdate.
	NextUpdate	time.Time

	// Extensions contains raw X.509 extensions. When creating a CRL,
	// the Extensions field is ignored, see ExtraExtensions.
	Extensions	[]pkix.Extension

	// ExtraExtensions contains any additional extensions to add directly to
	// the CRL.
	ExtraExtensions	[]pkix.Extension
}

// These structures reflect the ASN.1 structure of X.509 CRLs better than
//line /usr/local/go/src/crypto/x509/x509.go:2208
// the existing crypto/x509/pkix variants do. These mirror the existing
//line /usr/local/go/src/crypto/x509/x509.go:2208
// certificate structs in this file.
//line /usr/local/go/src/crypto/x509/x509.go:2208
//
//line /usr/local/go/src/crypto/x509/x509.go:2208
// Notably, we include issuer as an asn1.RawValue, mirroring the behavior of
//line /usr/local/go/src/crypto/x509/x509.go:2208
// tbsCertificate and allowing raw (unparsed) subjects to be passed cleanly.
//line /usr/local/go/src/crypto/x509/x509.go:2214
type certificateList struct {
	TBSCertList		tbsCertificateList
	SignatureAlgorithm	pkix.AlgorithmIdentifier
	SignatureValue		asn1.BitString
}

type tbsCertificateList struct {
	Raw			asn1.RawContent
	Version			int	`asn1:"optional,default:0"`
	Signature		pkix.AlgorithmIdentifier
	Issuer			asn1.RawValue
	ThisUpdate		time.Time
	NextUpdate		time.Time			`asn1:"optional"`
	RevokedCertificates	[]pkix.RevokedCertificate	`asn1:"optional"`
	Extensions		[]pkix.Extension		`asn1:"tag:0,optional,explicit"`
}

// CreateRevocationList creates a new X.509 v2 Certificate Revocation List,
//line /usr/local/go/src/crypto/x509/x509.go:2231
// according to RFC 5280, based on template.
//line /usr/local/go/src/crypto/x509/x509.go:2231
//
//line /usr/local/go/src/crypto/x509/x509.go:2231
// The CRL is signed by priv which should be the private key associated with
//line /usr/local/go/src/crypto/x509/x509.go:2231
// the public key in the issuer certificate.
//line /usr/local/go/src/crypto/x509/x509.go:2231
//
//line /usr/local/go/src/crypto/x509/x509.go:2231
// The issuer may not be nil, and the crlSign bit must be set in KeyUsage in
//line /usr/local/go/src/crypto/x509/x509.go:2231
// order to use it as a CRL issuer.
//line /usr/local/go/src/crypto/x509/x509.go:2231
//
//line /usr/local/go/src/crypto/x509/x509.go:2231
// The issuer distinguished name CRL field and authority key identifier
//line /usr/local/go/src/crypto/x509/x509.go:2231
// extension are populated using the issuer certificate. issuer must have
//line /usr/local/go/src/crypto/x509/x509.go:2231
// SubjectKeyId set.
//line /usr/local/go/src/crypto/x509/x509.go:2243
func CreateRevocationList(rand io.Reader, template *RevocationList, issuer *Certificate, priv crypto.Signer) ([]byte, error) {
//line /usr/local/go/src/crypto/x509/x509.go:2243
	_go_fuzz_dep_.CoverTab[20577]++
							if template == nil {
//line /usr/local/go/src/crypto/x509/x509.go:2244
		_go_fuzz_dep_.CoverTab[20596]++
								return nil, errors.New("x509: template can not be nil")
//line /usr/local/go/src/crypto/x509/x509.go:2245
		// _ = "end of CoverTab[20596]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2246
		_go_fuzz_dep_.CoverTab[20597]++
//line /usr/local/go/src/crypto/x509/x509.go:2246
		// _ = "end of CoverTab[20597]"
//line /usr/local/go/src/crypto/x509/x509.go:2246
	}
//line /usr/local/go/src/crypto/x509/x509.go:2246
	// _ = "end of CoverTab[20577]"
//line /usr/local/go/src/crypto/x509/x509.go:2246
	_go_fuzz_dep_.CoverTab[20578]++
							if issuer == nil {
//line /usr/local/go/src/crypto/x509/x509.go:2247
		_go_fuzz_dep_.CoverTab[20598]++
								return nil, errors.New("x509: issuer can not be nil")
//line /usr/local/go/src/crypto/x509/x509.go:2248
		// _ = "end of CoverTab[20598]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2249
		_go_fuzz_dep_.CoverTab[20599]++
//line /usr/local/go/src/crypto/x509/x509.go:2249
		// _ = "end of CoverTab[20599]"
//line /usr/local/go/src/crypto/x509/x509.go:2249
	}
//line /usr/local/go/src/crypto/x509/x509.go:2249
	// _ = "end of CoverTab[20578]"
//line /usr/local/go/src/crypto/x509/x509.go:2249
	_go_fuzz_dep_.CoverTab[20579]++
							if (issuer.KeyUsage & KeyUsageCRLSign) == 0 {
//line /usr/local/go/src/crypto/x509/x509.go:2250
		_go_fuzz_dep_.CoverTab[20600]++
								return nil, errors.New("x509: issuer must have the crlSign key usage bit set")
//line /usr/local/go/src/crypto/x509/x509.go:2251
		// _ = "end of CoverTab[20600]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2252
		_go_fuzz_dep_.CoverTab[20601]++
//line /usr/local/go/src/crypto/x509/x509.go:2252
		// _ = "end of CoverTab[20601]"
//line /usr/local/go/src/crypto/x509/x509.go:2252
	}
//line /usr/local/go/src/crypto/x509/x509.go:2252
	// _ = "end of CoverTab[20579]"
//line /usr/local/go/src/crypto/x509/x509.go:2252
	_go_fuzz_dep_.CoverTab[20580]++
							if len(issuer.SubjectKeyId) == 0 {
//line /usr/local/go/src/crypto/x509/x509.go:2253
		_go_fuzz_dep_.CoverTab[20602]++
								return nil, errors.New("x509: issuer certificate doesn't contain a subject key identifier")
//line /usr/local/go/src/crypto/x509/x509.go:2254
		// _ = "end of CoverTab[20602]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2255
		_go_fuzz_dep_.CoverTab[20603]++
//line /usr/local/go/src/crypto/x509/x509.go:2255
		// _ = "end of CoverTab[20603]"
//line /usr/local/go/src/crypto/x509/x509.go:2255
	}
//line /usr/local/go/src/crypto/x509/x509.go:2255
	// _ = "end of CoverTab[20580]"
//line /usr/local/go/src/crypto/x509/x509.go:2255
	_go_fuzz_dep_.CoverTab[20581]++
							if template.NextUpdate.Before(template.ThisUpdate) {
//line /usr/local/go/src/crypto/x509/x509.go:2256
		_go_fuzz_dep_.CoverTab[20604]++
								return nil, errors.New("x509: template.ThisUpdate is after template.NextUpdate")
//line /usr/local/go/src/crypto/x509/x509.go:2257
		// _ = "end of CoverTab[20604]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2258
		_go_fuzz_dep_.CoverTab[20605]++
//line /usr/local/go/src/crypto/x509/x509.go:2258
		// _ = "end of CoverTab[20605]"
//line /usr/local/go/src/crypto/x509/x509.go:2258
	}
//line /usr/local/go/src/crypto/x509/x509.go:2258
	// _ = "end of CoverTab[20581]"
//line /usr/local/go/src/crypto/x509/x509.go:2258
	_go_fuzz_dep_.CoverTab[20582]++
							if template.Number == nil {
//line /usr/local/go/src/crypto/x509/x509.go:2259
		_go_fuzz_dep_.CoverTab[20606]++
								return nil, errors.New("x509: template contains nil Number field")
//line /usr/local/go/src/crypto/x509/x509.go:2260
		// _ = "end of CoverTab[20606]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2261
		_go_fuzz_dep_.CoverTab[20607]++
//line /usr/local/go/src/crypto/x509/x509.go:2261
		// _ = "end of CoverTab[20607]"
//line /usr/local/go/src/crypto/x509/x509.go:2261
	}
//line /usr/local/go/src/crypto/x509/x509.go:2261
	// _ = "end of CoverTab[20582]"
//line /usr/local/go/src/crypto/x509/x509.go:2261
	_go_fuzz_dep_.CoverTab[20583]++

							hashFunc, signatureAlgorithm, err := signingParamsForPublicKey(priv.Public(), template.SignatureAlgorithm)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2264
		_go_fuzz_dep_.CoverTab[20608]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:2265
		// _ = "end of CoverTab[20608]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2266
		_go_fuzz_dep_.CoverTab[20609]++
//line /usr/local/go/src/crypto/x509/x509.go:2266
		// _ = "end of CoverTab[20609]"
//line /usr/local/go/src/crypto/x509/x509.go:2266
	}
//line /usr/local/go/src/crypto/x509/x509.go:2266
	// _ = "end of CoverTab[20583]"
//line /usr/local/go/src/crypto/x509/x509.go:2266
	_go_fuzz_dep_.CoverTab[20584]++

//line /usr/local/go/src/crypto/x509/x509.go:2269
	revokedCertsUTC := make([]pkix.RevokedCertificate, len(template.RevokedCertificates))
	for i, rc := range template.RevokedCertificates {
//line /usr/local/go/src/crypto/x509/x509.go:2270
		_go_fuzz_dep_.CoverTab[20610]++
								rc.RevocationTime = rc.RevocationTime.UTC()
								revokedCertsUTC[i] = rc
//line /usr/local/go/src/crypto/x509/x509.go:2272
		// _ = "end of CoverTab[20610]"
	}
//line /usr/local/go/src/crypto/x509/x509.go:2273
	// _ = "end of CoverTab[20584]"
//line /usr/local/go/src/crypto/x509/x509.go:2273
	_go_fuzz_dep_.CoverTab[20585]++

							aki, err := asn1.Marshal(authKeyId{Id: issuer.SubjectKeyId})
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2276
		_go_fuzz_dep_.CoverTab[20611]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:2277
		// _ = "end of CoverTab[20611]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2278
		_go_fuzz_dep_.CoverTab[20612]++
//line /usr/local/go/src/crypto/x509/x509.go:2278
		// _ = "end of CoverTab[20612]"
//line /usr/local/go/src/crypto/x509/x509.go:2278
	}
//line /usr/local/go/src/crypto/x509/x509.go:2278
	// _ = "end of CoverTab[20585]"
//line /usr/local/go/src/crypto/x509/x509.go:2278
	_go_fuzz_dep_.CoverTab[20586]++

							if numBytes := template.Number.Bytes(); len(numBytes) > 20 || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:2280
		_go_fuzz_dep_.CoverTab[20613]++
//line /usr/local/go/src/crypto/x509/x509.go:2280
		return (len(numBytes) == 20 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:2280
			_go_fuzz_dep_.CoverTab[20614]++
//line /usr/local/go/src/crypto/x509/x509.go:2280
			return numBytes[0]&0x80 != 0
//line /usr/local/go/src/crypto/x509/x509.go:2280
			// _ = "end of CoverTab[20614]"
//line /usr/local/go/src/crypto/x509/x509.go:2280
		}())
//line /usr/local/go/src/crypto/x509/x509.go:2280
		// _ = "end of CoverTab[20613]"
//line /usr/local/go/src/crypto/x509/x509.go:2280
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:2280
		_go_fuzz_dep_.CoverTab[20615]++
								return nil, errors.New("x509: CRL number exceeds 20 octets")
//line /usr/local/go/src/crypto/x509/x509.go:2281
		// _ = "end of CoverTab[20615]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2282
		_go_fuzz_dep_.CoverTab[20616]++
//line /usr/local/go/src/crypto/x509/x509.go:2282
		// _ = "end of CoverTab[20616]"
//line /usr/local/go/src/crypto/x509/x509.go:2282
	}
//line /usr/local/go/src/crypto/x509/x509.go:2282
	// _ = "end of CoverTab[20586]"
//line /usr/local/go/src/crypto/x509/x509.go:2282
	_go_fuzz_dep_.CoverTab[20587]++
							crlNum, err := asn1.Marshal(template.Number)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2284
		_go_fuzz_dep_.CoverTab[20617]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:2285
		// _ = "end of CoverTab[20617]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2286
		_go_fuzz_dep_.CoverTab[20618]++
//line /usr/local/go/src/crypto/x509/x509.go:2286
		// _ = "end of CoverTab[20618]"
//line /usr/local/go/src/crypto/x509/x509.go:2286
	}
//line /usr/local/go/src/crypto/x509/x509.go:2286
	// _ = "end of CoverTab[20587]"
//line /usr/local/go/src/crypto/x509/x509.go:2286
	_go_fuzz_dep_.CoverTab[20588]++

//line /usr/local/go/src/crypto/x509/x509.go:2289
	issuerSubject, err := subjectBytes(issuer)
	if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2290
		_go_fuzz_dep_.CoverTab[20619]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:2291
		// _ = "end of CoverTab[20619]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2292
		_go_fuzz_dep_.CoverTab[20620]++
//line /usr/local/go/src/crypto/x509/x509.go:2292
		// _ = "end of CoverTab[20620]"
//line /usr/local/go/src/crypto/x509/x509.go:2292
	}
//line /usr/local/go/src/crypto/x509/x509.go:2292
	// _ = "end of CoverTab[20588]"
//line /usr/local/go/src/crypto/x509/x509.go:2292
	_go_fuzz_dep_.CoverTab[20589]++

							tbsCertList := tbsCertificateList{
		Version:	1,
		Signature:	signatureAlgorithm,
		Issuer:		asn1.RawValue{FullBytes: issuerSubject},
		ThisUpdate:	template.ThisUpdate.UTC(),
		NextUpdate:	template.NextUpdate.UTC(),
		Extensions: []pkix.Extension{
			{
				Id:	oidExtensionAuthorityKeyId,
				Value:	aki,
			},
			{
				Id:	oidExtensionCRLNumber,
				Value:	crlNum,
			},
		},
	}
	if len(revokedCertsUTC) > 0 {
//line /usr/local/go/src/crypto/x509/x509.go:2311
		_go_fuzz_dep_.CoverTab[20621]++
								tbsCertList.RevokedCertificates = revokedCertsUTC
//line /usr/local/go/src/crypto/x509/x509.go:2312
		// _ = "end of CoverTab[20621]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2313
		_go_fuzz_dep_.CoverTab[20622]++
//line /usr/local/go/src/crypto/x509/x509.go:2313
		// _ = "end of CoverTab[20622]"
//line /usr/local/go/src/crypto/x509/x509.go:2313
	}
//line /usr/local/go/src/crypto/x509/x509.go:2313
	// _ = "end of CoverTab[20589]"
//line /usr/local/go/src/crypto/x509/x509.go:2313
	_go_fuzz_dep_.CoverTab[20590]++

							if len(template.ExtraExtensions) > 0 {
//line /usr/local/go/src/crypto/x509/x509.go:2315
		_go_fuzz_dep_.CoverTab[20623]++
								tbsCertList.Extensions = append(tbsCertList.Extensions, template.ExtraExtensions...)
//line /usr/local/go/src/crypto/x509/x509.go:2316
		// _ = "end of CoverTab[20623]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2317
		_go_fuzz_dep_.CoverTab[20624]++
//line /usr/local/go/src/crypto/x509/x509.go:2317
		// _ = "end of CoverTab[20624]"
//line /usr/local/go/src/crypto/x509/x509.go:2317
	}
//line /usr/local/go/src/crypto/x509/x509.go:2317
	// _ = "end of CoverTab[20590]"
//line /usr/local/go/src/crypto/x509/x509.go:2317
	_go_fuzz_dep_.CoverTab[20591]++

							tbsCertListContents, err := asn1.Marshal(tbsCertList)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2320
		_go_fuzz_dep_.CoverTab[20625]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:2321
		// _ = "end of CoverTab[20625]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2322
		_go_fuzz_dep_.CoverTab[20626]++
//line /usr/local/go/src/crypto/x509/x509.go:2322
		// _ = "end of CoverTab[20626]"
//line /usr/local/go/src/crypto/x509/x509.go:2322
	}
//line /usr/local/go/src/crypto/x509/x509.go:2322
	// _ = "end of CoverTab[20591]"
//line /usr/local/go/src/crypto/x509/x509.go:2322
	_go_fuzz_dep_.CoverTab[20592]++

//line /usr/local/go/src/crypto/x509/x509.go:2326
	tbsCertList.Raw = tbsCertListContents

	input := tbsCertListContents
	if hashFunc != 0 {
//line /usr/local/go/src/crypto/x509/x509.go:2329
		_go_fuzz_dep_.CoverTab[20627]++
								h := hashFunc.New()
								h.Write(tbsCertListContents)
								input = h.Sum(nil)
//line /usr/local/go/src/crypto/x509/x509.go:2332
		// _ = "end of CoverTab[20627]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2333
		_go_fuzz_dep_.CoverTab[20628]++
//line /usr/local/go/src/crypto/x509/x509.go:2333
		// _ = "end of CoverTab[20628]"
//line /usr/local/go/src/crypto/x509/x509.go:2333
	}
//line /usr/local/go/src/crypto/x509/x509.go:2333
	// _ = "end of CoverTab[20592]"
//line /usr/local/go/src/crypto/x509/x509.go:2333
	_go_fuzz_dep_.CoverTab[20593]++
							var signerOpts crypto.SignerOpts = hashFunc
							if template.SignatureAlgorithm.isRSAPSS() {
//line /usr/local/go/src/crypto/x509/x509.go:2335
		_go_fuzz_dep_.CoverTab[20629]++
								signerOpts = &rsa.PSSOptions{
			SaltLength:	rsa.PSSSaltLengthEqualsHash,
			Hash:		hashFunc,
		}
//line /usr/local/go/src/crypto/x509/x509.go:2339
		// _ = "end of CoverTab[20629]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2340
		_go_fuzz_dep_.CoverTab[20630]++
//line /usr/local/go/src/crypto/x509/x509.go:2340
		// _ = "end of CoverTab[20630]"
//line /usr/local/go/src/crypto/x509/x509.go:2340
	}
//line /usr/local/go/src/crypto/x509/x509.go:2340
	// _ = "end of CoverTab[20593]"
//line /usr/local/go/src/crypto/x509/x509.go:2340
	_go_fuzz_dep_.CoverTab[20594]++

							signature, err := priv.Sign(rand, input, signerOpts)
							if err != nil {
//line /usr/local/go/src/crypto/x509/x509.go:2343
		_go_fuzz_dep_.CoverTab[20631]++
								return nil, err
//line /usr/local/go/src/crypto/x509/x509.go:2344
		// _ = "end of CoverTab[20631]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2345
		_go_fuzz_dep_.CoverTab[20632]++
//line /usr/local/go/src/crypto/x509/x509.go:2345
		// _ = "end of CoverTab[20632]"
//line /usr/local/go/src/crypto/x509/x509.go:2345
	}
//line /usr/local/go/src/crypto/x509/x509.go:2345
	// _ = "end of CoverTab[20594]"
//line /usr/local/go/src/crypto/x509/x509.go:2345
	_go_fuzz_dep_.CoverTab[20595]++

							return asn1.Marshal(certificateList{
		TBSCertList:		tbsCertList,
		SignatureAlgorithm:	signatureAlgorithm,
		SignatureValue:		asn1.BitString{Bytes: signature, BitLength: len(signature) * 8},
	})
//line /usr/local/go/src/crypto/x509/x509.go:2351
	// _ = "end of CoverTab[20595]"
}

// CheckSignatureFrom verifies that the signature on rl is a valid signature
//line /usr/local/go/src/crypto/x509/x509.go:2354
// from issuer.
//line /usr/local/go/src/crypto/x509/x509.go:2356
func (rl *RevocationList) CheckSignatureFrom(parent *Certificate) error {
//line /usr/local/go/src/crypto/x509/x509.go:2356
	_go_fuzz_dep_.CoverTab[20633]++
							if parent.Version == 3 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:2357
		_go_fuzz_dep_.CoverTab[20637]++
//line /usr/local/go/src/crypto/x509/x509.go:2357
		return !parent.BasicConstraintsValid
//line /usr/local/go/src/crypto/x509/x509.go:2357
		// _ = "end of CoverTab[20637]"
//line /usr/local/go/src/crypto/x509/x509.go:2357
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:2357
		_go_fuzz_dep_.CoverTab[20638]++
//line /usr/local/go/src/crypto/x509/x509.go:2357
		return parent.BasicConstraintsValid && func() bool {
									_go_fuzz_dep_.CoverTab[20639]++
//line /usr/local/go/src/crypto/x509/x509.go:2358
			return !parent.IsCA
//line /usr/local/go/src/crypto/x509/x509.go:2358
			// _ = "end of CoverTab[20639]"
//line /usr/local/go/src/crypto/x509/x509.go:2358
		}()
//line /usr/local/go/src/crypto/x509/x509.go:2358
		// _ = "end of CoverTab[20638]"
//line /usr/local/go/src/crypto/x509/x509.go:2358
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:2358
		_go_fuzz_dep_.CoverTab[20640]++
								return ConstraintViolationError{}
//line /usr/local/go/src/crypto/x509/x509.go:2359
		// _ = "end of CoverTab[20640]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2360
		_go_fuzz_dep_.CoverTab[20641]++
//line /usr/local/go/src/crypto/x509/x509.go:2360
		// _ = "end of CoverTab[20641]"
//line /usr/local/go/src/crypto/x509/x509.go:2360
	}
//line /usr/local/go/src/crypto/x509/x509.go:2360
	// _ = "end of CoverTab[20633]"
//line /usr/local/go/src/crypto/x509/x509.go:2360
	_go_fuzz_dep_.CoverTab[20634]++

							if parent.KeyUsage != 0 && func() bool {
//line /usr/local/go/src/crypto/x509/x509.go:2362
		_go_fuzz_dep_.CoverTab[20642]++
//line /usr/local/go/src/crypto/x509/x509.go:2362
		return parent.KeyUsage&KeyUsageCRLSign == 0
//line /usr/local/go/src/crypto/x509/x509.go:2362
		// _ = "end of CoverTab[20642]"
//line /usr/local/go/src/crypto/x509/x509.go:2362
	}() {
//line /usr/local/go/src/crypto/x509/x509.go:2362
		_go_fuzz_dep_.CoverTab[20643]++
								return ConstraintViolationError{}
//line /usr/local/go/src/crypto/x509/x509.go:2363
		// _ = "end of CoverTab[20643]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2364
		_go_fuzz_dep_.CoverTab[20644]++
//line /usr/local/go/src/crypto/x509/x509.go:2364
		// _ = "end of CoverTab[20644]"
//line /usr/local/go/src/crypto/x509/x509.go:2364
	}
//line /usr/local/go/src/crypto/x509/x509.go:2364
	// _ = "end of CoverTab[20634]"
//line /usr/local/go/src/crypto/x509/x509.go:2364
	_go_fuzz_dep_.CoverTab[20635]++

							if parent.PublicKeyAlgorithm == UnknownPublicKeyAlgorithm {
//line /usr/local/go/src/crypto/x509/x509.go:2366
		_go_fuzz_dep_.CoverTab[20645]++
								return ErrUnsupportedAlgorithm
//line /usr/local/go/src/crypto/x509/x509.go:2367
		// _ = "end of CoverTab[20645]"
	} else {
//line /usr/local/go/src/crypto/x509/x509.go:2368
		_go_fuzz_dep_.CoverTab[20646]++
//line /usr/local/go/src/crypto/x509/x509.go:2368
		// _ = "end of CoverTab[20646]"
//line /usr/local/go/src/crypto/x509/x509.go:2368
	}
//line /usr/local/go/src/crypto/x509/x509.go:2368
	// _ = "end of CoverTab[20635]"
//line /usr/local/go/src/crypto/x509/x509.go:2368
	_go_fuzz_dep_.CoverTab[20636]++

							return parent.CheckSignature(rl.SignatureAlgorithm, rl.RawTBSRevocationList, rl.Signature)
//line /usr/local/go/src/crypto/x509/x509.go:2370
	// _ = "end of CoverTab[20636]"
}

//line /usr/local/go/src/crypto/x509/x509.go:2371
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/x509/x509.go:2371
var _ = _go_fuzz_dep_.CoverTab
