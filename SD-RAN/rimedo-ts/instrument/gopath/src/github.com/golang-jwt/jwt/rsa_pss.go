// +build go1.4

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:3
package jwt

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:3
)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:3
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:3
)

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
)

// Implements the RSAPSS family of signing methods signing methods
type SigningMethodRSAPSS struct {
	*SigningMethodRSA
	Options	*rsa.PSSOptions
	// VerifyOptions is optional. If set overrides Options for rsa.VerifyPPS.
	// Used to accept tokens signed with rsa.PSSSaltLengthAuto, what doesn't follow
	// https://tools.ietf.org/html/rfc7518#section-3.5 but was used previously.
	// See https://github.com/dgrijalva/jwt-go/issues/285#issuecomment-437451244 for details.
	VerifyOptions	*rsa.PSSOptions
}

// Specific instances for RS/PS and company.
var (
	SigningMethodPS256	*SigningMethodRSAPSS
	SigningMethodPS384	*SigningMethodRSAPSS
	SigningMethodPS512	*SigningMethodRSAPSS
)

func init() {

	SigningMethodPS256 = &SigningMethodRSAPSS{
		SigningMethodRSA: &SigningMethodRSA{
			Name:	"PS256",
			Hash:	crypto.SHA256,
		},
		Options: &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthEqualsHash,
		},
		VerifyOptions: &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthAuto,
		},
	}
	RegisterSigningMethod(SigningMethodPS256.Alg(), func() SigningMethod {
		return SigningMethodPS256
	})

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:48
	SigningMethodPS384 = &SigningMethodRSAPSS{
		SigningMethodRSA: &SigningMethodRSA{
			Name:	"PS384",
			Hash:	crypto.SHA384,
		},
		Options: &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthEqualsHash,
		},
		VerifyOptions: &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthAuto,
		},
	}
	RegisterSigningMethod(SigningMethodPS384.Alg(), func() SigningMethod {
		return SigningMethodPS384
	})

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:65
	SigningMethodPS512 = &SigningMethodRSAPSS{
		SigningMethodRSA: &SigningMethodRSA{
			Name:	"PS512",
			Hash:	crypto.SHA512,
		},
		Options: &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthEqualsHash,
		},
		VerifyOptions: &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthAuto,
		},
	}
	RegisterSigningMethod(SigningMethodPS512.Alg(), func() SigningMethod {
		return SigningMethodPS512
	})
}

// Implements the Verify method from SigningMethod
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:82
// For this verify method, key must be an rsa.PublicKey struct
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:84
func (m *SigningMethodRSAPSS) Verify(signingString, signature string, key interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:84
	_go_fuzz_dep_.CoverTab[187251]++
												var err error

	// Decode the signature
	var sig []byte
	if sig, err = DecodeSegment(signature); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:89
		_go_fuzz_dep_.CoverTab[187256]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:90
		// _ = "end of CoverTab[187256]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:91
		_go_fuzz_dep_.CoverTab[187257]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:91
		// _ = "end of CoverTab[187257]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:91
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:91
	// _ = "end of CoverTab[187251]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:91
	_go_fuzz_dep_.CoverTab[187252]++

												var rsaKey *rsa.PublicKey
												switch k := key.(type) {
	case *rsa.PublicKey:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:95
		_go_fuzz_dep_.CoverTab[187258]++
													rsaKey = k
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:96
		// _ = "end of CoverTab[187258]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:97
		_go_fuzz_dep_.CoverTab[187259]++
													return ErrInvalidKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:98
		// _ = "end of CoverTab[187259]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:99
	// _ = "end of CoverTab[187252]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:99
	_go_fuzz_dep_.CoverTab[187253]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:102
	if !m.Hash.Available() {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:102
		_go_fuzz_dep_.CoverTab[187260]++
													return ErrHashUnavailable
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:103
		// _ = "end of CoverTab[187260]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:104
		_go_fuzz_dep_.CoverTab[187261]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:104
		// _ = "end of CoverTab[187261]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:104
	// _ = "end of CoverTab[187253]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:104
	_go_fuzz_dep_.CoverTab[187254]++
												hasher := m.Hash.New()
												hasher.Write([]byte(signingString))

												opts := m.Options
												if m.VerifyOptions != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:109
		_go_fuzz_dep_.CoverTab[187262]++
													opts = m.VerifyOptions
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:110
		// _ = "end of CoverTab[187262]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:111
		_go_fuzz_dep_.CoverTab[187263]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:111
		// _ = "end of CoverTab[187263]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:111
	// _ = "end of CoverTab[187254]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:111
	_go_fuzz_dep_.CoverTab[187255]++

												return rsa.VerifyPSS(rsaKey, m.Hash, hasher.Sum(nil), sig, opts)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:113
	// _ = "end of CoverTab[187255]"
}

// Implements the Sign method from SigningMethod
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:116
// For this signing method, key must be an rsa.PrivateKey struct
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:118
func (m *SigningMethodRSAPSS) Sign(signingString string, key interface{}) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:118
	_go_fuzz_dep_.CoverTab[187264]++
												var rsaKey *rsa.PrivateKey

												switch k := key.(type) {
	case *rsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:122
		_go_fuzz_dep_.CoverTab[187267]++
													rsaKey = k
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:123
		// _ = "end of CoverTab[187267]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:124
		_go_fuzz_dep_.CoverTab[187268]++
													return "", ErrInvalidKeyType
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:125
		// _ = "end of CoverTab[187268]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:126
	// _ = "end of CoverTab[187264]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:126
	_go_fuzz_dep_.CoverTab[187265]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:129
	if !m.Hash.Available() {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:129
		_go_fuzz_dep_.CoverTab[187269]++
													return "", ErrHashUnavailable
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:130
		// _ = "end of CoverTab[187269]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:131
		_go_fuzz_dep_.CoverTab[187270]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:131
		// _ = "end of CoverTab[187270]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:131
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:131
	// _ = "end of CoverTab[187265]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:131
	_go_fuzz_dep_.CoverTab[187266]++

												hasher := m.Hash.New()
												hasher.Write([]byte(signingString))

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:137
	if sigBytes, err := rsa.SignPSS(rand.Reader, rsaKey, m.Hash, hasher.Sum(nil), m.Options); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:137
		_go_fuzz_dep_.CoverTab[187271]++
													return EncodeSegment(sigBytes), nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:138
		// _ = "end of CoverTab[187271]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:139
		_go_fuzz_dep_.CoverTab[187272]++
													return "", err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:140
		// _ = "end of CoverTab[187272]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:141
	// _ = "end of CoverTab[187266]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:142
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_pss.go:142
var _ = _go_fuzz_dep_.CoverTab
