// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/x509/pkcs8.go:5
package x509

//line /usr/local/go/src/crypto/x509/pkcs8.go:5
import (
//line /usr/local/go/src/crypto/x509/pkcs8.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/x509/pkcs8.go:5
)
//line /usr/local/go/src/crypto/x509/pkcs8.go:5
import (
//line /usr/local/go/src/crypto/x509/pkcs8.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/x509/pkcs8.go:5
)

import (
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509/pkix"
	"encoding/asn1"
	"errors"
	"fmt"
)

// pkcs8 reflects an ASN.1, PKCS #8 PrivateKey. See
//line /usr/local/go/src/crypto/x509/pkcs8.go:18
// ftp://ftp.rsasecurity.com/pub/pkcs/pkcs-8/pkcs-8v1_2.asn
//line /usr/local/go/src/crypto/x509/pkcs8.go:18
// and RFC 5208.
//line /usr/local/go/src/crypto/x509/pkcs8.go:21
type pkcs8 struct {
	Version		int
	Algo		pkix.AlgorithmIdentifier
	PrivateKey	[]byte
//line /usr/local/go/src/crypto/x509/pkcs8.go:26
}

// ParsePKCS8PrivateKey parses an unencrypted private key in PKCS #8, ASN.1 DER form.
//line /usr/local/go/src/crypto/x509/pkcs8.go:28
//
//line /usr/local/go/src/crypto/x509/pkcs8.go:28
// It returns a *rsa.PrivateKey, a *ecdsa.PrivateKey, a ed25519.PrivateKey (not
//line /usr/local/go/src/crypto/x509/pkcs8.go:28
// a pointer), or a *ecdh.PrivateKey (for X25519). More types might be supported
//line /usr/local/go/src/crypto/x509/pkcs8.go:28
// in the future.
//line /usr/local/go/src/crypto/x509/pkcs8.go:28
//
//line /usr/local/go/src/crypto/x509/pkcs8.go:28
// This kind of key is commonly encoded in PEM blocks of type "PRIVATE KEY".
//line /usr/local/go/src/crypto/x509/pkcs8.go:35
func ParsePKCS8PrivateKey(der []byte) (key any, err error) {
//line /usr/local/go/src/crypto/x509/pkcs8.go:35
	_go_fuzz_dep_.CoverTab[19181]++
							var privKey pkcs8
							if _, err := asn1.Unmarshal(der, &privKey); err != nil {
//line /usr/local/go/src/crypto/x509/pkcs8.go:37
		_go_fuzz_dep_.CoverTab[19183]++
								if _, err := asn1.Unmarshal(der, &ecPrivateKey{}); err == nil {
//line /usr/local/go/src/crypto/x509/pkcs8.go:38
			_go_fuzz_dep_.CoverTab[19186]++
									return nil, errors.New("x509: failed to parse private key (use ParseECPrivateKey instead for this key format)")
//line /usr/local/go/src/crypto/x509/pkcs8.go:39
			// _ = "end of CoverTab[19186]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:40
			_go_fuzz_dep_.CoverTab[19187]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:40
			// _ = "end of CoverTab[19187]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:40
		}
//line /usr/local/go/src/crypto/x509/pkcs8.go:40
		// _ = "end of CoverTab[19183]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:40
		_go_fuzz_dep_.CoverTab[19184]++
								if _, err := asn1.Unmarshal(der, &pkcs1PrivateKey{}); err == nil {
//line /usr/local/go/src/crypto/x509/pkcs8.go:41
			_go_fuzz_dep_.CoverTab[19188]++
									return nil, errors.New("x509: failed to parse private key (use ParsePKCS1PrivateKey instead for this key format)")
//line /usr/local/go/src/crypto/x509/pkcs8.go:42
			// _ = "end of CoverTab[19188]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:43
			_go_fuzz_dep_.CoverTab[19189]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:43
			// _ = "end of CoverTab[19189]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:43
		}
//line /usr/local/go/src/crypto/x509/pkcs8.go:43
		// _ = "end of CoverTab[19184]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:43
		_go_fuzz_dep_.CoverTab[19185]++
								return nil, err
//line /usr/local/go/src/crypto/x509/pkcs8.go:44
		// _ = "end of CoverTab[19185]"
	} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:45
		_go_fuzz_dep_.CoverTab[19190]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:45
		// _ = "end of CoverTab[19190]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:45
	}
//line /usr/local/go/src/crypto/x509/pkcs8.go:45
	// _ = "end of CoverTab[19181]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:45
	_go_fuzz_dep_.CoverTab[19182]++
							switch {
	case privKey.Algo.Algorithm.Equal(oidPublicKeyRSA):
//line /usr/local/go/src/crypto/x509/pkcs8.go:47
		_go_fuzz_dep_.CoverTab[19191]++
								key, err = ParsePKCS1PrivateKey(privKey.PrivateKey)
								if err != nil {
//line /usr/local/go/src/crypto/x509/pkcs8.go:49
			_go_fuzz_dep_.CoverTab[19204]++
									return nil, errors.New("x509: failed to parse RSA private key embedded in PKCS#8: " + err.Error())
//line /usr/local/go/src/crypto/x509/pkcs8.go:50
			// _ = "end of CoverTab[19204]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:51
			_go_fuzz_dep_.CoverTab[19205]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:51
			// _ = "end of CoverTab[19205]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:51
		}
//line /usr/local/go/src/crypto/x509/pkcs8.go:51
		// _ = "end of CoverTab[19191]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:51
		_go_fuzz_dep_.CoverTab[19192]++
								return key, nil
//line /usr/local/go/src/crypto/x509/pkcs8.go:52
		// _ = "end of CoverTab[19192]"

	case privKey.Algo.Algorithm.Equal(oidPublicKeyECDSA):
//line /usr/local/go/src/crypto/x509/pkcs8.go:54
		_go_fuzz_dep_.CoverTab[19193]++
								bytes := privKey.Algo.Parameters.FullBytes
								namedCurveOID := new(asn1.ObjectIdentifier)
								if _, err := asn1.Unmarshal(bytes, namedCurveOID); err != nil {
//line /usr/local/go/src/crypto/x509/pkcs8.go:57
			_go_fuzz_dep_.CoverTab[19206]++
									namedCurveOID = nil
//line /usr/local/go/src/crypto/x509/pkcs8.go:58
			// _ = "end of CoverTab[19206]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:59
			_go_fuzz_dep_.CoverTab[19207]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:59
			// _ = "end of CoverTab[19207]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:59
		}
//line /usr/local/go/src/crypto/x509/pkcs8.go:59
		// _ = "end of CoverTab[19193]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:59
		_go_fuzz_dep_.CoverTab[19194]++
								key, err = parseECPrivateKey(namedCurveOID, privKey.PrivateKey)
								if err != nil {
//line /usr/local/go/src/crypto/x509/pkcs8.go:61
			_go_fuzz_dep_.CoverTab[19208]++
									return nil, errors.New("x509: failed to parse EC private key embedded in PKCS#8: " + err.Error())
//line /usr/local/go/src/crypto/x509/pkcs8.go:62
			// _ = "end of CoverTab[19208]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:63
			_go_fuzz_dep_.CoverTab[19209]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:63
			// _ = "end of CoverTab[19209]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:63
		}
//line /usr/local/go/src/crypto/x509/pkcs8.go:63
		// _ = "end of CoverTab[19194]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:63
		_go_fuzz_dep_.CoverTab[19195]++
								return key, nil
//line /usr/local/go/src/crypto/x509/pkcs8.go:64
		// _ = "end of CoverTab[19195]"

	case privKey.Algo.Algorithm.Equal(oidPublicKeyEd25519):
//line /usr/local/go/src/crypto/x509/pkcs8.go:66
		_go_fuzz_dep_.CoverTab[19196]++
								if l := len(privKey.Algo.Parameters.FullBytes); l != 0 {
//line /usr/local/go/src/crypto/x509/pkcs8.go:67
			_go_fuzz_dep_.CoverTab[19210]++
									return nil, errors.New("x509: invalid Ed25519 private key parameters")
//line /usr/local/go/src/crypto/x509/pkcs8.go:68
			// _ = "end of CoverTab[19210]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:69
			_go_fuzz_dep_.CoverTab[19211]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:69
			// _ = "end of CoverTab[19211]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:69
		}
//line /usr/local/go/src/crypto/x509/pkcs8.go:69
		// _ = "end of CoverTab[19196]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:69
		_go_fuzz_dep_.CoverTab[19197]++
								var curvePrivateKey []byte
								if _, err := asn1.Unmarshal(privKey.PrivateKey, &curvePrivateKey); err != nil {
//line /usr/local/go/src/crypto/x509/pkcs8.go:71
			_go_fuzz_dep_.CoverTab[19212]++
									return nil, fmt.Errorf("x509: invalid Ed25519 private key: %v", err)
//line /usr/local/go/src/crypto/x509/pkcs8.go:72
			// _ = "end of CoverTab[19212]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:73
			_go_fuzz_dep_.CoverTab[19213]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:73
			// _ = "end of CoverTab[19213]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:73
		}
//line /usr/local/go/src/crypto/x509/pkcs8.go:73
		// _ = "end of CoverTab[19197]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:73
		_go_fuzz_dep_.CoverTab[19198]++
								if l := len(curvePrivateKey); l != ed25519.SeedSize {
//line /usr/local/go/src/crypto/x509/pkcs8.go:74
			_go_fuzz_dep_.CoverTab[19214]++
									return nil, fmt.Errorf("x509: invalid Ed25519 private key length: %d", l)
//line /usr/local/go/src/crypto/x509/pkcs8.go:75
			// _ = "end of CoverTab[19214]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:76
			_go_fuzz_dep_.CoverTab[19215]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:76
			// _ = "end of CoverTab[19215]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:76
		}
//line /usr/local/go/src/crypto/x509/pkcs8.go:76
		// _ = "end of CoverTab[19198]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:76
		_go_fuzz_dep_.CoverTab[19199]++
								return ed25519.NewKeyFromSeed(curvePrivateKey), nil
//line /usr/local/go/src/crypto/x509/pkcs8.go:77
		// _ = "end of CoverTab[19199]"

	case privKey.Algo.Algorithm.Equal(oidPublicKeyX25519):
//line /usr/local/go/src/crypto/x509/pkcs8.go:79
		_go_fuzz_dep_.CoverTab[19200]++
								if l := len(privKey.Algo.Parameters.FullBytes); l != 0 {
//line /usr/local/go/src/crypto/x509/pkcs8.go:80
			_go_fuzz_dep_.CoverTab[19216]++
									return nil, errors.New("x509: invalid X25519 private key parameters")
//line /usr/local/go/src/crypto/x509/pkcs8.go:81
			// _ = "end of CoverTab[19216]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:82
			_go_fuzz_dep_.CoverTab[19217]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:82
			// _ = "end of CoverTab[19217]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:82
		}
//line /usr/local/go/src/crypto/x509/pkcs8.go:82
		// _ = "end of CoverTab[19200]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:82
		_go_fuzz_dep_.CoverTab[19201]++
								var curvePrivateKey []byte
								if _, err := asn1.Unmarshal(privKey.PrivateKey, &curvePrivateKey); err != nil {
//line /usr/local/go/src/crypto/x509/pkcs8.go:84
			_go_fuzz_dep_.CoverTab[19218]++
									return nil, fmt.Errorf("x509: invalid X25519 private key: %v", err)
//line /usr/local/go/src/crypto/x509/pkcs8.go:85
			// _ = "end of CoverTab[19218]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:86
			_go_fuzz_dep_.CoverTab[19219]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:86
			// _ = "end of CoverTab[19219]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:86
		}
//line /usr/local/go/src/crypto/x509/pkcs8.go:86
		// _ = "end of CoverTab[19201]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:86
		_go_fuzz_dep_.CoverTab[19202]++
								return ecdh.X25519().NewPrivateKey(curvePrivateKey)
//line /usr/local/go/src/crypto/x509/pkcs8.go:87
		// _ = "end of CoverTab[19202]"

	default:
//line /usr/local/go/src/crypto/x509/pkcs8.go:89
		_go_fuzz_dep_.CoverTab[19203]++
								return nil, fmt.Errorf("x509: PKCS#8 wrapping contained private key with unknown algorithm: %v", privKey.Algo.Algorithm)
//line /usr/local/go/src/crypto/x509/pkcs8.go:90
		// _ = "end of CoverTab[19203]"
	}
//line /usr/local/go/src/crypto/x509/pkcs8.go:91
	// _ = "end of CoverTab[19182]"
}

// MarshalPKCS8PrivateKey converts a private key to PKCS #8, ASN.1 DER form.
//line /usr/local/go/src/crypto/x509/pkcs8.go:94
//
//line /usr/local/go/src/crypto/x509/pkcs8.go:94
// The following key types are currently supported: *rsa.PrivateKey,
//line /usr/local/go/src/crypto/x509/pkcs8.go:94
// *ecdsa.PrivateKey, ed25519.PrivateKey (not a pointer), and *ecdh.PrivateKey.
//line /usr/local/go/src/crypto/x509/pkcs8.go:94
// Unsupported key types result in an error.
//line /usr/local/go/src/crypto/x509/pkcs8.go:94
//
//line /usr/local/go/src/crypto/x509/pkcs8.go:94
// This kind of key is commonly encoded in PEM blocks of type "PRIVATE KEY".
//line /usr/local/go/src/crypto/x509/pkcs8.go:101
func MarshalPKCS8PrivateKey(key any) ([]byte, error) {
//line /usr/local/go/src/crypto/x509/pkcs8.go:101
	_go_fuzz_dep_.CoverTab[19220]++
							var privKey pkcs8

							switch k := key.(type) {
	case *rsa.PrivateKey:
//line /usr/local/go/src/crypto/x509/pkcs8.go:105
		_go_fuzz_dep_.CoverTab[19222]++
								privKey.Algo = pkix.AlgorithmIdentifier{
			Algorithm:	oidPublicKeyRSA,
			Parameters:	asn1.NullRawValue,
		}
								privKey.PrivateKey = MarshalPKCS1PrivateKey(k)
//line /usr/local/go/src/crypto/x509/pkcs8.go:110
		// _ = "end of CoverTab[19222]"

	case *ecdsa.PrivateKey:
//line /usr/local/go/src/crypto/x509/pkcs8.go:112
		_go_fuzz_dep_.CoverTab[19223]++
								oid, ok := oidFromNamedCurve(k.Curve)
								if !ok {
//line /usr/local/go/src/crypto/x509/pkcs8.go:114
			_go_fuzz_dep_.CoverTab[19230]++
									return nil, errors.New("x509: unknown curve while marshaling to PKCS#8")
//line /usr/local/go/src/crypto/x509/pkcs8.go:115
			// _ = "end of CoverTab[19230]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:116
			_go_fuzz_dep_.CoverTab[19231]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:116
			// _ = "end of CoverTab[19231]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:116
		}
//line /usr/local/go/src/crypto/x509/pkcs8.go:116
		// _ = "end of CoverTab[19223]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:116
		_go_fuzz_dep_.CoverTab[19224]++
								oidBytes, err := asn1.Marshal(oid)
								if err != nil {
//line /usr/local/go/src/crypto/x509/pkcs8.go:118
			_go_fuzz_dep_.CoverTab[19232]++
									return nil, errors.New("x509: failed to marshal curve OID: " + err.Error())
//line /usr/local/go/src/crypto/x509/pkcs8.go:119
			// _ = "end of CoverTab[19232]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:120
			_go_fuzz_dep_.CoverTab[19233]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:120
			// _ = "end of CoverTab[19233]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:120
		}
//line /usr/local/go/src/crypto/x509/pkcs8.go:120
		// _ = "end of CoverTab[19224]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:120
		_go_fuzz_dep_.CoverTab[19225]++
								privKey.Algo = pkix.AlgorithmIdentifier{
			Algorithm:	oidPublicKeyECDSA,
			Parameters: asn1.RawValue{
				FullBytes: oidBytes,
			},
		}
		if privKey.PrivateKey, err = marshalECPrivateKeyWithOID(k, nil); err != nil {
//line /usr/local/go/src/crypto/x509/pkcs8.go:127
			_go_fuzz_dep_.CoverTab[19234]++
									return nil, errors.New("x509: failed to marshal EC private key while building PKCS#8: " + err.Error())
//line /usr/local/go/src/crypto/x509/pkcs8.go:128
			// _ = "end of CoverTab[19234]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:129
			_go_fuzz_dep_.CoverTab[19235]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:129
			// _ = "end of CoverTab[19235]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:129
		}
//line /usr/local/go/src/crypto/x509/pkcs8.go:129
		// _ = "end of CoverTab[19225]"

	case ed25519.PrivateKey:
//line /usr/local/go/src/crypto/x509/pkcs8.go:131
		_go_fuzz_dep_.CoverTab[19226]++
								privKey.Algo = pkix.AlgorithmIdentifier{
			Algorithm: oidPublicKeyEd25519,
		}
		curvePrivateKey, err := asn1.Marshal(k.Seed())
		if err != nil {
//line /usr/local/go/src/crypto/x509/pkcs8.go:136
			_go_fuzz_dep_.CoverTab[19236]++
									return nil, fmt.Errorf("x509: failed to marshal private key: %v", err)
//line /usr/local/go/src/crypto/x509/pkcs8.go:137
			// _ = "end of CoverTab[19236]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:138
			_go_fuzz_dep_.CoverTab[19237]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:138
			// _ = "end of CoverTab[19237]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:138
		}
//line /usr/local/go/src/crypto/x509/pkcs8.go:138
		// _ = "end of CoverTab[19226]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:138
		_go_fuzz_dep_.CoverTab[19227]++
								privKey.PrivateKey = curvePrivateKey
//line /usr/local/go/src/crypto/x509/pkcs8.go:139
		// _ = "end of CoverTab[19227]"

	case *ecdh.PrivateKey:
//line /usr/local/go/src/crypto/x509/pkcs8.go:141
		_go_fuzz_dep_.CoverTab[19228]++
								if k.Curve() == ecdh.X25519() {
//line /usr/local/go/src/crypto/x509/pkcs8.go:142
			_go_fuzz_dep_.CoverTab[19238]++
									privKey.Algo = pkix.AlgorithmIdentifier{
				Algorithm: oidPublicKeyX25519,
			}
			var err error
			if privKey.PrivateKey, err = asn1.Marshal(k.Bytes()); err != nil {
//line /usr/local/go/src/crypto/x509/pkcs8.go:147
				_go_fuzz_dep_.CoverTab[19239]++
										return nil, fmt.Errorf("x509: failed to marshal private key: %v", err)
//line /usr/local/go/src/crypto/x509/pkcs8.go:148
				// _ = "end of CoverTab[19239]"
			} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:149
				_go_fuzz_dep_.CoverTab[19240]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:149
				// _ = "end of CoverTab[19240]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:149
			}
//line /usr/local/go/src/crypto/x509/pkcs8.go:149
			// _ = "end of CoverTab[19238]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:150
			_go_fuzz_dep_.CoverTab[19241]++
									oid, ok := oidFromECDHCurve(k.Curve())
									if !ok {
//line /usr/local/go/src/crypto/x509/pkcs8.go:152
				_go_fuzz_dep_.CoverTab[19244]++
										return nil, errors.New("x509: unknown curve while marshaling to PKCS#8")
//line /usr/local/go/src/crypto/x509/pkcs8.go:153
				// _ = "end of CoverTab[19244]"
			} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:154
				_go_fuzz_dep_.CoverTab[19245]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:154
				// _ = "end of CoverTab[19245]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:154
			}
//line /usr/local/go/src/crypto/x509/pkcs8.go:154
			// _ = "end of CoverTab[19241]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:154
			_go_fuzz_dep_.CoverTab[19242]++
									oidBytes, err := asn1.Marshal(oid)
									if err != nil {
//line /usr/local/go/src/crypto/x509/pkcs8.go:156
				_go_fuzz_dep_.CoverTab[19246]++
										return nil, errors.New("x509: failed to marshal curve OID: " + err.Error())
//line /usr/local/go/src/crypto/x509/pkcs8.go:157
				// _ = "end of CoverTab[19246]"
			} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:158
				_go_fuzz_dep_.CoverTab[19247]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:158
				// _ = "end of CoverTab[19247]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:158
			}
//line /usr/local/go/src/crypto/x509/pkcs8.go:158
			// _ = "end of CoverTab[19242]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:158
			_go_fuzz_dep_.CoverTab[19243]++
									privKey.Algo = pkix.AlgorithmIdentifier{
				Algorithm:	oidPublicKeyECDSA,
				Parameters: asn1.RawValue{
					FullBytes: oidBytes,
				},
			}
			if privKey.PrivateKey, err = marshalECDHPrivateKey(k); err != nil {
//line /usr/local/go/src/crypto/x509/pkcs8.go:165
				_go_fuzz_dep_.CoverTab[19248]++
										return nil, errors.New("x509: failed to marshal EC private key while building PKCS#8: " + err.Error())
//line /usr/local/go/src/crypto/x509/pkcs8.go:166
				// _ = "end of CoverTab[19248]"
			} else {
//line /usr/local/go/src/crypto/x509/pkcs8.go:167
				_go_fuzz_dep_.CoverTab[19249]++
//line /usr/local/go/src/crypto/x509/pkcs8.go:167
				// _ = "end of CoverTab[19249]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:167
			}
//line /usr/local/go/src/crypto/x509/pkcs8.go:167
			// _ = "end of CoverTab[19243]"
		}
//line /usr/local/go/src/crypto/x509/pkcs8.go:168
		// _ = "end of CoverTab[19228]"

	default:
//line /usr/local/go/src/crypto/x509/pkcs8.go:170
		_go_fuzz_dep_.CoverTab[19229]++
								return nil, fmt.Errorf("x509: unknown key type while marshaling PKCS#8: %T", key)
//line /usr/local/go/src/crypto/x509/pkcs8.go:171
		// _ = "end of CoverTab[19229]"
	}
//line /usr/local/go/src/crypto/x509/pkcs8.go:172
	// _ = "end of CoverTab[19220]"
//line /usr/local/go/src/crypto/x509/pkcs8.go:172
	_go_fuzz_dep_.CoverTab[19221]++

							return asn1.Marshal(privKey)
//line /usr/local/go/src/crypto/x509/pkcs8.go:174
	// _ = "end of CoverTab[19221]"
}

//line /usr/local/go/src/crypto/x509/pkcs8.go:175
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/x509/pkcs8.go:175
var _ = _go_fuzz_dep_.CoverTab
