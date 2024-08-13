// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/x509/pkcs1.go:5
package x509

//line /usr/local/go/src/crypto/x509/pkcs1.go:5
import (
//line /usr/local/go/src/crypto/x509/pkcs1.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/x509/pkcs1.go:5
)
//line /usr/local/go/src/crypto/x509/pkcs1.go:5
import (
//line /usr/local/go/src/crypto/x509/pkcs1.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/x509/pkcs1.go:5
)

import (
	"crypto/rsa"
	"encoding/asn1"
	"errors"
	"math/big"
)

// pkcs1PrivateKey is a structure which mirrors the PKCS #1 ASN.1 for an RSA private key.
type pkcs1PrivateKey struct {
	Version	int
	N	*big.Int
	E	int
	D	*big.Int
	P	*big.Int
	Q	*big.Int
	// We ignore these values, if present, because rsa will calculate them.
	Dp	*big.Int	`asn1:"optional"`
	Dq	*big.Int	`asn1:"optional"`
	Qinv	*big.Int	`asn1:"optional"`

	AdditionalPrimes	[]pkcs1AdditionalRSAPrime	`asn1:"optional,omitempty"`
}

type pkcs1AdditionalRSAPrime struct {
	Prime	*big.Int

	// We ignore these values because rsa will calculate them.
	Exp	*big.Int
	Coeff	*big.Int
}

// pkcs1PublicKey reflects the ASN.1 structure of a PKCS #1 public key.
type pkcs1PublicKey struct {
	N	*big.Int
	E	int
}

// ParsePKCS1PrivateKey parses an RSA private key in PKCS #1, ASN.1 DER form.
//line /usr/local/go/src/crypto/x509/pkcs1.go:44
//
//line /usr/local/go/src/crypto/x509/pkcs1.go:44
// This kind of key is commonly encoded in PEM blocks of type "RSA PRIVATE KEY".
//line /usr/local/go/src/crypto/x509/pkcs1.go:47
func ParsePKCS1PrivateKey(der []byte) (*rsa.PrivateKey, error) {
//line /usr/local/go/src/crypto/x509/pkcs1.go:47
	_go_fuzz_dep_.CoverTab[19127]++
							var priv pkcs1PrivateKey
							rest, err := asn1.Unmarshal(der, &priv)
							if len(rest) > 0 {
//line /usr/local/go/src/crypto/x509/pkcs1.go:50
		_go_fuzz_dep_.CoverTab[19134]++
								return nil, asn1.SyntaxError{Msg: "trailing data"}
//line /usr/local/go/src/crypto/x509/pkcs1.go:51
		// _ = "end of CoverTab[19134]"
	} else {
//line /usr/local/go/src/crypto/x509/pkcs1.go:52
		_go_fuzz_dep_.CoverTab[19135]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:52
		// _ = "end of CoverTab[19135]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:52
	}
//line /usr/local/go/src/crypto/x509/pkcs1.go:52
	// _ = "end of CoverTab[19127]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:52
	_go_fuzz_dep_.CoverTab[19128]++
							if err != nil {
//line /usr/local/go/src/crypto/x509/pkcs1.go:53
		_go_fuzz_dep_.CoverTab[19136]++
								if _, err := asn1.Unmarshal(der, &ecPrivateKey{}); err == nil {
//line /usr/local/go/src/crypto/x509/pkcs1.go:54
			_go_fuzz_dep_.CoverTab[19139]++
									return nil, errors.New("x509: failed to parse private key (use ParseECPrivateKey instead for this key format)")
//line /usr/local/go/src/crypto/x509/pkcs1.go:55
			// _ = "end of CoverTab[19139]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs1.go:56
			_go_fuzz_dep_.CoverTab[19140]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:56
			// _ = "end of CoverTab[19140]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:56
		}
//line /usr/local/go/src/crypto/x509/pkcs1.go:56
		// _ = "end of CoverTab[19136]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:56
		_go_fuzz_dep_.CoverTab[19137]++
								if _, err := asn1.Unmarshal(der, &pkcs8{}); err == nil {
//line /usr/local/go/src/crypto/x509/pkcs1.go:57
			_go_fuzz_dep_.CoverTab[19141]++
									return nil, errors.New("x509: failed to parse private key (use ParsePKCS8PrivateKey instead for this key format)")
//line /usr/local/go/src/crypto/x509/pkcs1.go:58
			// _ = "end of CoverTab[19141]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs1.go:59
			_go_fuzz_dep_.CoverTab[19142]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:59
			// _ = "end of CoverTab[19142]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:59
		}
//line /usr/local/go/src/crypto/x509/pkcs1.go:59
		// _ = "end of CoverTab[19137]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:59
		_go_fuzz_dep_.CoverTab[19138]++
								return nil, err
//line /usr/local/go/src/crypto/x509/pkcs1.go:60
		// _ = "end of CoverTab[19138]"
	} else {
//line /usr/local/go/src/crypto/x509/pkcs1.go:61
		_go_fuzz_dep_.CoverTab[19143]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:61
		// _ = "end of CoverTab[19143]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:61
	}
//line /usr/local/go/src/crypto/x509/pkcs1.go:61
	// _ = "end of CoverTab[19128]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:61
	_go_fuzz_dep_.CoverTab[19129]++

							if priv.Version > 1 {
//line /usr/local/go/src/crypto/x509/pkcs1.go:63
		_go_fuzz_dep_.CoverTab[19144]++
								return nil, errors.New("x509: unsupported private key version")
//line /usr/local/go/src/crypto/x509/pkcs1.go:64
		// _ = "end of CoverTab[19144]"
	} else {
//line /usr/local/go/src/crypto/x509/pkcs1.go:65
		_go_fuzz_dep_.CoverTab[19145]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:65
		// _ = "end of CoverTab[19145]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:65
	}
//line /usr/local/go/src/crypto/x509/pkcs1.go:65
	// _ = "end of CoverTab[19129]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:65
	_go_fuzz_dep_.CoverTab[19130]++

							if priv.N.Sign() <= 0 || func() bool {
//line /usr/local/go/src/crypto/x509/pkcs1.go:67
		_go_fuzz_dep_.CoverTab[19146]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:67
		return priv.D.Sign() <= 0
//line /usr/local/go/src/crypto/x509/pkcs1.go:67
		// _ = "end of CoverTab[19146]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:67
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/pkcs1.go:67
		_go_fuzz_dep_.CoverTab[19147]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:67
		return priv.P.Sign() <= 0
//line /usr/local/go/src/crypto/x509/pkcs1.go:67
		// _ = "end of CoverTab[19147]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:67
	}() || func() bool {
//line /usr/local/go/src/crypto/x509/pkcs1.go:67
		_go_fuzz_dep_.CoverTab[19148]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:67
		return priv.Q.Sign() <= 0
//line /usr/local/go/src/crypto/x509/pkcs1.go:67
		// _ = "end of CoverTab[19148]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:67
	}() {
//line /usr/local/go/src/crypto/x509/pkcs1.go:67
		_go_fuzz_dep_.CoverTab[19149]++
								return nil, errors.New("x509: private key contains zero or negative value")
//line /usr/local/go/src/crypto/x509/pkcs1.go:68
		// _ = "end of CoverTab[19149]"
	} else {
//line /usr/local/go/src/crypto/x509/pkcs1.go:69
		_go_fuzz_dep_.CoverTab[19150]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:69
		// _ = "end of CoverTab[19150]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:69
	}
//line /usr/local/go/src/crypto/x509/pkcs1.go:69
	// _ = "end of CoverTab[19130]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:69
	_go_fuzz_dep_.CoverTab[19131]++

							key := new(rsa.PrivateKey)
							key.PublicKey = rsa.PublicKey{
		E:	priv.E,
		N:	priv.N,
	}

	key.D = priv.D
	key.Primes = make([]*big.Int, 2+len(priv.AdditionalPrimes))
	key.Primes[0] = priv.P
	key.Primes[1] = priv.Q
	for i, a := range priv.AdditionalPrimes {
//line /usr/local/go/src/crypto/x509/pkcs1.go:81
		_go_fuzz_dep_.CoverTab[19151]++
								if a.Prime.Sign() <= 0 {
//line /usr/local/go/src/crypto/x509/pkcs1.go:82
			_go_fuzz_dep_.CoverTab[19153]++
									return nil, errors.New("x509: private key contains zero or negative prime")
//line /usr/local/go/src/crypto/x509/pkcs1.go:83
			// _ = "end of CoverTab[19153]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs1.go:84
			_go_fuzz_dep_.CoverTab[19154]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:84
			// _ = "end of CoverTab[19154]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:84
		}
//line /usr/local/go/src/crypto/x509/pkcs1.go:84
		// _ = "end of CoverTab[19151]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:84
		_go_fuzz_dep_.CoverTab[19152]++
								key.Primes[i+2] = a.Prime
//line /usr/local/go/src/crypto/x509/pkcs1.go:85
		// _ = "end of CoverTab[19152]"

//line /usr/local/go/src/crypto/x509/pkcs1.go:88
	}
//line /usr/local/go/src/crypto/x509/pkcs1.go:88
	// _ = "end of CoverTab[19131]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:88
	_go_fuzz_dep_.CoverTab[19132]++

							err = key.Validate()
							if err != nil {
//line /usr/local/go/src/crypto/x509/pkcs1.go:91
		_go_fuzz_dep_.CoverTab[19155]++
								return nil, err
//line /usr/local/go/src/crypto/x509/pkcs1.go:92
		// _ = "end of CoverTab[19155]"
	} else {
//line /usr/local/go/src/crypto/x509/pkcs1.go:93
		_go_fuzz_dep_.CoverTab[19156]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:93
		// _ = "end of CoverTab[19156]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:93
	}
//line /usr/local/go/src/crypto/x509/pkcs1.go:93
	// _ = "end of CoverTab[19132]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:93
	_go_fuzz_dep_.CoverTab[19133]++
							key.Precompute()

							return key, nil
//line /usr/local/go/src/crypto/x509/pkcs1.go:96
	// _ = "end of CoverTab[19133]"
}

// MarshalPKCS1PrivateKey converts an RSA private key to PKCS #1, ASN.1 DER form.
//line /usr/local/go/src/crypto/x509/pkcs1.go:99
//
//line /usr/local/go/src/crypto/x509/pkcs1.go:99
// This kind of key is commonly encoded in PEM blocks of type "RSA PRIVATE KEY".
//line /usr/local/go/src/crypto/x509/pkcs1.go:99
// For a more flexible key format which is not RSA specific, use
//line /usr/local/go/src/crypto/x509/pkcs1.go:99
// MarshalPKCS8PrivateKey.
//line /usr/local/go/src/crypto/x509/pkcs1.go:104
func MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte {
//line /usr/local/go/src/crypto/x509/pkcs1.go:104
	_go_fuzz_dep_.CoverTab[19157]++
							key.Precompute()

							version := 0
							if len(key.Primes) > 2 {
//line /usr/local/go/src/crypto/x509/pkcs1.go:108
		_go_fuzz_dep_.CoverTab[19160]++
								version = 1
//line /usr/local/go/src/crypto/x509/pkcs1.go:109
		// _ = "end of CoverTab[19160]"
	} else {
//line /usr/local/go/src/crypto/x509/pkcs1.go:110
		_go_fuzz_dep_.CoverTab[19161]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:110
		// _ = "end of CoverTab[19161]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:110
	}
//line /usr/local/go/src/crypto/x509/pkcs1.go:110
	// _ = "end of CoverTab[19157]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:110
	_go_fuzz_dep_.CoverTab[19158]++

							priv := pkcs1PrivateKey{
		Version:	version,
		N:		key.N,
		E:		key.PublicKey.E,
		D:		key.D,
		P:		key.Primes[0],
		Q:		key.Primes[1],
		Dp:		key.Precomputed.Dp,
		Dq:		key.Precomputed.Dq,
		Qinv:		key.Precomputed.Qinv,
	}

	priv.AdditionalPrimes = make([]pkcs1AdditionalRSAPrime, len(key.Precomputed.CRTValues))
	for i, values := range key.Precomputed.CRTValues {
//line /usr/local/go/src/crypto/x509/pkcs1.go:125
		_go_fuzz_dep_.CoverTab[19162]++
								priv.AdditionalPrimes[i].Prime = key.Primes[2+i]
								priv.AdditionalPrimes[i].Exp = values.Exp
								priv.AdditionalPrimes[i].Coeff = values.Coeff
//line /usr/local/go/src/crypto/x509/pkcs1.go:128
		// _ = "end of CoverTab[19162]"
	}
//line /usr/local/go/src/crypto/x509/pkcs1.go:129
	// _ = "end of CoverTab[19158]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:129
	_go_fuzz_dep_.CoverTab[19159]++

							b, _ := asn1.Marshal(priv)
							return b
//line /usr/local/go/src/crypto/x509/pkcs1.go:132
	// _ = "end of CoverTab[19159]"
}

// ParsePKCS1PublicKey parses an RSA public key in PKCS #1, ASN.1 DER form.
//line /usr/local/go/src/crypto/x509/pkcs1.go:135
//
//line /usr/local/go/src/crypto/x509/pkcs1.go:135
// This kind of key is commonly encoded in PEM blocks of type "RSA PUBLIC KEY".
//line /usr/local/go/src/crypto/x509/pkcs1.go:138
func ParsePKCS1PublicKey(der []byte) (*rsa.PublicKey, error) {
//line /usr/local/go/src/crypto/x509/pkcs1.go:138
	_go_fuzz_dep_.CoverTab[19163]++
							var pub pkcs1PublicKey
							rest, err := asn1.Unmarshal(der, &pub)
							if err != nil {
//line /usr/local/go/src/crypto/x509/pkcs1.go:141
		_go_fuzz_dep_.CoverTab[19168]++
								if _, err := asn1.Unmarshal(der, &publicKeyInfo{}); err == nil {
//line /usr/local/go/src/crypto/x509/pkcs1.go:142
			_go_fuzz_dep_.CoverTab[19170]++
									return nil, errors.New("x509: failed to parse public key (use ParsePKIXPublicKey instead for this key format)")
//line /usr/local/go/src/crypto/x509/pkcs1.go:143
			// _ = "end of CoverTab[19170]"
		} else {
//line /usr/local/go/src/crypto/x509/pkcs1.go:144
			_go_fuzz_dep_.CoverTab[19171]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:144
			// _ = "end of CoverTab[19171]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:144
		}
//line /usr/local/go/src/crypto/x509/pkcs1.go:144
		// _ = "end of CoverTab[19168]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:144
		_go_fuzz_dep_.CoverTab[19169]++
								return nil, err
//line /usr/local/go/src/crypto/x509/pkcs1.go:145
		// _ = "end of CoverTab[19169]"
	} else {
//line /usr/local/go/src/crypto/x509/pkcs1.go:146
		_go_fuzz_dep_.CoverTab[19172]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:146
		// _ = "end of CoverTab[19172]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:146
	}
//line /usr/local/go/src/crypto/x509/pkcs1.go:146
	// _ = "end of CoverTab[19163]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:146
	_go_fuzz_dep_.CoverTab[19164]++
							if len(rest) > 0 {
//line /usr/local/go/src/crypto/x509/pkcs1.go:147
		_go_fuzz_dep_.CoverTab[19173]++
								return nil, asn1.SyntaxError{Msg: "trailing data"}
//line /usr/local/go/src/crypto/x509/pkcs1.go:148
		// _ = "end of CoverTab[19173]"
	} else {
//line /usr/local/go/src/crypto/x509/pkcs1.go:149
		_go_fuzz_dep_.CoverTab[19174]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:149
		// _ = "end of CoverTab[19174]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:149
	}
//line /usr/local/go/src/crypto/x509/pkcs1.go:149
	// _ = "end of CoverTab[19164]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:149
	_go_fuzz_dep_.CoverTab[19165]++

							if pub.N.Sign() <= 0 || func() bool {
//line /usr/local/go/src/crypto/x509/pkcs1.go:151
		_go_fuzz_dep_.CoverTab[19175]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:151
		return pub.E <= 0
//line /usr/local/go/src/crypto/x509/pkcs1.go:151
		// _ = "end of CoverTab[19175]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:151
	}() {
//line /usr/local/go/src/crypto/x509/pkcs1.go:151
		_go_fuzz_dep_.CoverTab[19176]++
								return nil, errors.New("x509: public key contains zero or negative value")
//line /usr/local/go/src/crypto/x509/pkcs1.go:152
		// _ = "end of CoverTab[19176]"
	} else {
//line /usr/local/go/src/crypto/x509/pkcs1.go:153
		_go_fuzz_dep_.CoverTab[19177]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:153
		// _ = "end of CoverTab[19177]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:153
	}
//line /usr/local/go/src/crypto/x509/pkcs1.go:153
	// _ = "end of CoverTab[19165]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:153
	_go_fuzz_dep_.CoverTab[19166]++
							if pub.E > 1<<31-1 {
//line /usr/local/go/src/crypto/x509/pkcs1.go:154
		_go_fuzz_dep_.CoverTab[19178]++
								return nil, errors.New("x509: public key contains large public exponent")
//line /usr/local/go/src/crypto/x509/pkcs1.go:155
		// _ = "end of CoverTab[19178]"
	} else {
//line /usr/local/go/src/crypto/x509/pkcs1.go:156
		_go_fuzz_dep_.CoverTab[19179]++
//line /usr/local/go/src/crypto/x509/pkcs1.go:156
		// _ = "end of CoverTab[19179]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:156
	}
//line /usr/local/go/src/crypto/x509/pkcs1.go:156
	// _ = "end of CoverTab[19166]"
//line /usr/local/go/src/crypto/x509/pkcs1.go:156
	_go_fuzz_dep_.CoverTab[19167]++

							return &rsa.PublicKey{
		E:	pub.E,
		N:	pub.N,
	}, nil
//line /usr/local/go/src/crypto/x509/pkcs1.go:161
	// _ = "end of CoverTab[19167]"
}

// MarshalPKCS1PublicKey converts an RSA public key to PKCS #1, ASN.1 DER form.
//line /usr/local/go/src/crypto/x509/pkcs1.go:164
//
//line /usr/local/go/src/crypto/x509/pkcs1.go:164
// This kind of key is commonly encoded in PEM blocks of type "RSA PUBLIC KEY".
//line /usr/local/go/src/crypto/x509/pkcs1.go:167
func MarshalPKCS1PublicKey(key *rsa.PublicKey) []byte {
//line /usr/local/go/src/crypto/x509/pkcs1.go:167
	_go_fuzz_dep_.CoverTab[19180]++
							derBytes, _ := asn1.Marshal(pkcs1PublicKey{
		N:	key.N,
		E:	key.E,
	})
							return derBytes
//line /usr/local/go/src/crypto/x509/pkcs1.go:172
	// _ = "end of CoverTab[19180]"
}

//line /usr/local/go/src/crypto/x509/pkcs1.go:173
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/x509/pkcs1.go:173
var _ = _go_fuzz_dep_.CoverTab
