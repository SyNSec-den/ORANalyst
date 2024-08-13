// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/x509/sec1.go:5
package x509

//line /usr/local/go/src/crypto/x509/sec1.go:5
import (
//line /usr/local/go/src/crypto/x509/sec1.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/x509/sec1.go:5
)
//line /usr/local/go/src/crypto/x509/sec1.go:5
import (
//line /usr/local/go/src/crypto/x509/sec1.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/x509/sec1.go:5
)

import (
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/asn1"
	"errors"
	"fmt"
	"math/big"
)

const ecPrivKeyVersion = 1

// ecPrivateKey reflects an ASN.1 Elliptic Curve Private Key Structure.
//line /usr/local/go/src/crypto/x509/sec1.go:19
// References:
//line /usr/local/go/src/crypto/x509/sec1.go:19
//
//line /usr/local/go/src/crypto/x509/sec1.go:19
//	RFC 5915
//line /usr/local/go/src/crypto/x509/sec1.go:19
//	SEC1 - http://www.secg.org/sec1-v2.pdf
//line /usr/local/go/src/crypto/x509/sec1.go:19
//
//line /usr/local/go/src/crypto/x509/sec1.go:19
// Per RFC 5915 the NamedCurveOID is marked as ASN.1 OPTIONAL, however in
//line /usr/local/go/src/crypto/x509/sec1.go:19
// most cases it is not.
//line /usr/local/go/src/crypto/x509/sec1.go:27
type ecPrivateKey struct {
	Version		int
	PrivateKey	[]byte
	NamedCurveOID	asn1.ObjectIdentifier	`asn1:"optional,explicit,tag:0"`
	PublicKey	asn1.BitString		`asn1:"optional,explicit,tag:1"`
}

// ParseECPrivateKey parses an EC private key in SEC 1, ASN.1 DER form.
//line /usr/local/go/src/crypto/x509/sec1.go:34
//
//line /usr/local/go/src/crypto/x509/sec1.go:34
// This kind of key is commonly encoded in PEM blocks of type "EC PRIVATE KEY".
//line /usr/local/go/src/crypto/x509/sec1.go:37
func ParseECPrivateKey(der []byte) (*ecdsa.PrivateKey, error) {
//line /usr/local/go/src/crypto/x509/sec1.go:37
	_go_fuzz_dep_.CoverTab[19312]++
							return parseECPrivateKey(nil, der)
//line /usr/local/go/src/crypto/x509/sec1.go:38
	// _ = "end of CoverTab[19312]"
}

// MarshalECPrivateKey converts an EC private key to SEC 1, ASN.1 DER form.
//line /usr/local/go/src/crypto/x509/sec1.go:41
//
//line /usr/local/go/src/crypto/x509/sec1.go:41
// This kind of key is commonly encoded in PEM blocks of type "EC PRIVATE KEY".
//line /usr/local/go/src/crypto/x509/sec1.go:41
// For a more flexible key format which is not EC specific, use
//line /usr/local/go/src/crypto/x509/sec1.go:41
// MarshalPKCS8PrivateKey.
//line /usr/local/go/src/crypto/x509/sec1.go:46
func MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error) {
//line /usr/local/go/src/crypto/x509/sec1.go:46
	_go_fuzz_dep_.CoverTab[19313]++
							oid, ok := oidFromNamedCurve(key.Curve)
							if !ok {
//line /usr/local/go/src/crypto/x509/sec1.go:48
		_go_fuzz_dep_.CoverTab[19315]++
								return nil, errors.New("x509: unknown elliptic curve")
//line /usr/local/go/src/crypto/x509/sec1.go:49
		// _ = "end of CoverTab[19315]"
	} else {
//line /usr/local/go/src/crypto/x509/sec1.go:50
		_go_fuzz_dep_.CoverTab[19316]++
//line /usr/local/go/src/crypto/x509/sec1.go:50
		// _ = "end of CoverTab[19316]"
//line /usr/local/go/src/crypto/x509/sec1.go:50
	}
//line /usr/local/go/src/crypto/x509/sec1.go:50
	// _ = "end of CoverTab[19313]"
//line /usr/local/go/src/crypto/x509/sec1.go:50
	_go_fuzz_dep_.CoverTab[19314]++

							return marshalECPrivateKeyWithOID(key, oid)
//line /usr/local/go/src/crypto/x509/sec1.go:52
	// _ = "end of CoverTab[19314]"
}

// marshalECPrivateKeyWithOID marshals an EC private key into ASN.1, DER format and
//line /usr/local/go/src/crypto/x509/sec1.go:55
// sets the curve ID to the given OID, or omits it if OID is nil.
//line /usr/local/go/src/crypto/x509/sec1.go:57
func marshalECPrivateKeyWithOID(key *ecdsa.PrivateKey, oid asn1.ObjectIdentifier) ([]byte, error) {
//line /usr/local/go/src/crypto/x509/sec1.go:57
	_go_fuzz_dep_.CoverTab[19317]++
							if !key.Curve.IsOnCurve(key.X, key.Y) {
//line /usr/local/go/src/crypto/x509/sec1.go:58
		_go_fuzz_dep_.CoverTab[19319]++
								return nil, errors.New("invalid elliptic key public key")
//line /usr/local/go/src/crypto/x509/sec1.go:59
		// _ = "end of CoverTab[19319]"
	} else {
//line /usr/local/go/src/crypto/x509/sec1.go:60
		_go_fuzz_dep_.CoverTab[19320]++
//line /usr/local/go/src/crypto/x509/sec1.go:60
		// _ = "end of CoverTab[19320]"
//line /usr/local/go/src/crypto/x509/sec1.go:60
	}
//line /usr/local/go/src/crypto/x509/sec1.go:60
	// _ = "end of CoverTab[19317]"
//line /usr/local/go/src/crypto/x509/sec1.go:60
	_go_fuzz_dep_.CoverTab[19318]++
							privateKey := make([]byte, (key.Curve.Params().N.BitLen()+7)/8)
							return asn1.Marshal(ecPrivateKey{
		Version:	1,
		PrivateKey:	key.D.FillBytes(privateKey),
		NamedCurveOID:	oid,
		PublicKey:	asn1.BitString{Bytes: elliptic.Marshal(key.Curve, key.X, key.Y)},
	})
//line /usr/local/go/src/crypto/x509/sec1.go:67
	// _ = "end of CoverTab[19318]"
}

// marshalECPrivateKeyWithOID marshals an EC private key into ASN.1, DER format
//line /usr/local/go/src/crypto/x509/sec1.go:70
// suitable for NIST curves.
//line /usr/local/go/src/crypto/x509/sec1.go:72
func marshalECDHPrivateKey(key *ecdh.PrivateKey) ([]byte, error) {
//line /usr/local/go/src/crypto/x509/sec1.go:72
	_go_fuzz_dep_.CoverTab[19321]++
							return asn1.Marshal(ecPrivateKey{
		Version:	1,
		PrivateKey:	key.Bytes(),
		PublicKey:	asn1.BitString{Bytes: key.PublicKey().Bytes()},
	})
//line /usr/local/go/src/crypto/x509/sec1.go:77
	// _ = "end of CoverTab[19321]"
}

// parseECPrivateKey parses an ASN.1 Elliptic Curve Private Key Structure.
//line /usr/local/go/src/crypto/x509/sec1.go:80
// The OID for the named curve may be provided from another source (such as
//line /usr/local/go/src/crypto/x509/sec1.go:80
// the PKCS8 container) - if it is provided then use this instead of the OID
//line /usr/local/go/src/crypto/x509/sec1.go:80
// that may exist in the EC private key structure.
//line /usr/local/go/src/crypto/x509/sec1.go:84
func parseECPrivateKey(namedCurveOID *asn1.ObjectIdentifier, der []byte) (key *ecdsa.PrivateKey, err error) {
//line /usr/local/go/src/crypto/x509/sec1.go:84
	_go_fuzz_dep_.CoverTab[19322]++
							var privKey ecPrivateKey
							if _, err := asn1.Unmarshal(der, &privKey); err != nil {
//line /usr/local/go/src/crypto/x509/sec1.go:86
		_go_fuzz_dep_.CoverTab[19329]++
								if _, err := asn1.Unmarshal(der, &pkcs8{}); err == nil {
//line /usr/local/go/src/crypto/x509/sec1.go:87
			_go_fuzz_dep_.CoverTab[19332]++
									return nil, errors.New("x509: failed to parse private key (use ParsePKCS8PrivateKey instead for this key format)")
//line /usr/local/go/src/crypto/x509/sec1.go:88
			// _ = "end of CoverTab[19332]"
		} else {
//line /usr/local/go/src/crypto/x509/sec1.go:89
			_go_fuzz_dep_.CoverTab[19333]++
//line /usr/local/go/src/crypto/x509/sec1.go:89
			// _ = "end of CoverTab[19333]"
//line /usr/local/go/src/crypto/x509/sec1.go:89
		}
//line /usr/local/go/src/crypto/x509/sec1.go:89
		// _ = "end of CoverTab[19329]"
//line /usr/local/go/src/crypto/x509/sec1.go:89
		_go_fuzz_dep_.CoverTab[19330]++
								if _, err := asn1.Unmarshal(der, &pkcs1PrivateKey{}); err == nil {
//line /usr/local/go/src/crypto/x509/sec1.go:90
			_go_fuzz_dep_.CoverTab[19334]++
									return nil, errors.New("x509: failed to parse private key (use ParsePKCS1PrivateKey instead for this key format)")
//line /usr/local/go/src/crypto/x509/sec1.go:91
			// _ = "end of CoverTab[19334]"
		} else {
//line /usr/local/go/src/crypto/x509/sec1.go:92
			_go_fuzz_dep_.CoverTab[19335]++
//line /usr/local/go/src/crypto/x509/sec1.go:92
			// _ = "end of CoverTab[19335]"
//line /usr/local/go/src/crypto/x509/sec1.go:92
		}
//line /usr/local/go/src/crypto/x509/sec1.go:92
		// _ = "end of CoverTab[19330]"
//line /usr/local/go/src/crypto/x509/sec1.go:92
		_go_fuzz_dep_.CoverTab[19331]++
								return nil, errors.New("x509: failed to parse EC private key: " + err.Error())
//line /usr/local/go/src/crypto/x509/sec1.go:93
		// _ = "end of CoverTab[19331]"
	} else {
//line /usr/local/go/src/crypto/x509/sec1.go:94
		_go_fuzz_dep_.CoverTab[19336]++
//line /usr/local/go/src/crypto/x509/sec1.go:94
		// _ = "end of CoverTab[19336]"
//line /usr/local/go/src/crypto/x509/sec1.go:94
	}
//line /usr/local/go/src/crypto/x509/sec1.go:94
	// _ = "end of CoverTab[19322]"
//line /usr/local/go/src/crypto/x509/sec1.go:94
	_go_fuzz_dep_.CoverTab[19323]++
							if privKey.Version != ecPrivKeyVersion {
//line /usr/local/go/src/crypto/x509/sec1.go:95
		_go_fuzz_dep_.CoverTab[19337]++
								return nil, fmt.Errorf("x509: unknown EC private key version %d", privKey.Version)
//line /usr/local/go/src/crypto/x509/sec1.go:96
		// _ = "end of CoverTab[19337]"
	} else {
//line /usr/local/go/src/crypto/x509/sec1.go:97
		_go_fuzz_dep_.CoverTab[19338]++
//line /usr/local/go/src/crypto/x509/sec1.go:97
		// _ = "end of CoverTab[19338]"
//line /usr/local/go/src/crypto/x509/sec1.go:97
	}
//line /usr/local/go/src/crypto/x509/sec1.go:97
	// _ = "end of CoverTab[19323]"
//line /usr/local/go/src/crypto/x509/sec1.go:97
	_go_fuzz_dep_.CoverTab[19324]++

							var curve elliptic.Curve
							if namedCurveOID != nil {
//line /usr/local/go/src/crypto/x509/sec1.go:100
		_go_fuzz_dep_.CoverTab[19339]++
								curve = namedCurveFromOID(*namedCurveOID)
//line /usr/local/go/src/crypto/x509/sec1.go:101
		// _ = "end of CoverTab[19339]"
	} else {
//line /usr/local/go/src/crypto/x509/sec1.go:102
		_go_fuzz_dep_.CoverTab[19340]++
								curve = namedCurveFromOID(privKey.NamedCurveOID)
//line /usr/local/go/src/crypto/x509/sec1.go:103
		// _ = "end of CoverTab[19340]"
	}
//line /usr/local/go/src/crypto/x509/sec1.go:104
	// _ = "end of CoverTab[19324]"
//line /usr/local/go/src/crypto/x509/sec1.go:104
	_go_fuzz_dep_.CoverTab[19325]++
							if curve == nil {
//line /usr/local/go/src/crypto/x509/sec1.go:105
		_go_fuzz_dep_.CoverTab[19341]++
								return nil, errors.New("x509: unknown elliptic curve")
//line /usr/local/go/src/crypto/x509/sec1.go:106
		// _ = "end of CoverTab[19341]"
	} else {
//line /usr/local/go/src/crypto/x509/sec1.go:107
		_go_fuzz_dep_.CoverTab[19342]++
//line /usr/local/go/src/crypto/x509/sec1.go:107
		// _ = "end of CoverTab[19342]"
//line /usr/local/go/src/crypto/x509/sec1.go:107
	}
//line /usr/local/go/src/crypto/x509/sec1.go:107
	// _ = "end of CoverTab[19325]"
//line /usr/local/go/src/crypto/x509/sec1.go:107
	_go_fuzz_dep_.CoverTab[19326]++

							k := new(big.Int).SetBytes(privKey.PrivateKey)
							curveOrder := curve.Params().N
							if k.Cmp(curveOrder) >= 0 {
//line /usr/local/go/src/crypto/x509/sec1.go:111
		_go_fuzz_dep_.CoverTab[19343]++
								return nil, errors.New("x509: invalid elliptic curve private key value")
//line /usr/local/go/src/crypto/x509/sec1.go:112
		// _ = "end of CoverTab[19343]"
	} else {
//line /usr/local/go/src/crypto/x509/sec1.go:113
		_go_fuzz_dep_.CoverTab[19344]++
//line /usr/local/go/src/crypto/x509/sec1.go:113
		// _ = "end of CoverTab[19344]"
//line /usr/local/go/src/crypto/x509/sec1.go:113
	}
//line /usr/local/go/src/crypto/x509/sec1.go:113
	// _ = "end of CoverTab[19326]"
//line /usr/local/go/src/crypto/x509/sec1.go:113
	_go_fuzz_dep_.CoverTab[19327]++
							priv := new(ecdsa.PrivateKey)
							priv.Curve = curve
							priv.D = k

							privateKey := make([]byte, (curveOrder.BitLen()+7)/8)

//line /usr/local/go/src/crypto/x509/sec1.go:122
	for len(privKey.PrivateKey) > len(privateKey) {
//line /usr/local/go/src/crypto/x509/sec1.go:122
		_go_fuzz_dep_.CoverTab[19345]++
								if privKey.PrivateKey[0] != 0 {
//line /usr/local/go/src/crypto/x509/sec1.go:123
			_go_fuzz_dep_.CoverTab[19347]++
									return nil, errors.New("x509: invalid private key length")
//line /usr/local/go/src/crypto/x509/sec1.go:124
			// _ = "end of CoverTab[19347]"
		} else {
//line /usr/local/go/src/crypto/x509/sec1.go:125
			_go_fuzz_dep_.CoverTab[19348]++
//line /usr/local/go/src/crypto/x509/sec1.go:125
			// _ = "end of CoverTab[19348]"
//line /usr/local/go/src/crypto/x509/sec1.go:125
		}
//line /usr/local/go/src/crypto/x509/sec1.go:125
		// _ = "end of CoverTab[19345]"
//line /usr/local/go/src/crypto/x509/sec1.go:125
		_go_fuzz_dep_.CoverTab[19346]++
								privKey.PrivateKey = privKey.PrivateKey[1:]
//line /usr/local/go/src/crypto/x509/sec1.go:126
		// _ = "end of CoverTab[19346]"
	}
//line /usr/local/go/src/crypto/x509/sec1.go:127
	// _ = "end of CoverTab[19327]"
//line /usr/local/go/src/crypto/x509/sec1.go:127
	_go_fuzz_dep_.CoverTab[19328]++

//line /usr/local/go/src/crypto/x509/sec1.go:132
	copy(privateKey[len(privateKey)-len(privKey.PrivateKey):], privKey.PrivateKey)
							priv.X, priv.Y = curve.ScalarBaseMult(privateKey)

							return priv, nil
//line /usr/local/go/src/crypto/x509/sec1.go:135
	// _ = "end of CoverTab[19328]"
}

//line /usr/local/go/src/crypto/x509/sec1.go:136
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/x509/sec1.go:136
var _ = _go_fuzz_dep_.CoverTab
