//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:1
package jwt

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:1
)

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"math/big"
)

var (
	// Sadly this is missing from crypto/ecdsa compared to crypto/rsa
	ErrECDSAVerification = errors.New("crypto/ecdsa: verification error")
)

// Implements the ECDSA family of signing methods signing methods
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:16
// Expects *ecdsa.PrivateKey for signing and *ecdsa.PublicKey for verification
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:18
type SigningMethodECDSA struct {
	Name		string
	Hash		crypto.Hash
	KeySize		int
	CurveBits	int
}

// Specific instances for EC256 and company
var (
	SigningMethodES256	*SigningMethodECDSA
	SigningMethodES384	*SigningMethodECDSA
	SigningMethodES512	*SigningMethodECDSA
)

func init() {

	SigningMethodES256 = &SigningMethodECDSA{"ES256", crypto.SHA256, 32, 256}
	RegisterSigningMethod(SigningMethodES256.Alg(), func() SigningMethod {
		return SigningMethodES256
	})

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:40
	SigningMethodES384 = &SigningMethodECDSA{"ES384", crypto.SHA384, 48, 384}
	RegisterSigningMethod(SigningMethodES384.Alg(), func() SigningMethod {
		return SigningMethodES384
	})

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:46
	SigningMethodES512 = &SigningMethodECDSA{"ES512", crypto.SHA512, 66, 521}
	RegisterSigningMethod(SigningMethodES512.Alg(), func() SigningMethod {
		return SigningMethodES512
	})
}

func (m *SigningMethodECDSA) Alg() string {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:52
	_go_fuzz_dep_.CoverTab[186984]++
												return m.Name
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:53
	// _ = "end of CoverTab[186984]"
}

// Implements the Verify method from SigningMethod
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:56
// For this verify method, key must be an ecdsa.PublicKey struct
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:58
func (m *SigningMethodECDSA) Verify(signingString, signature string, key interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:58
	_go_fuzz_dep_.CoverTab[186985]++
												var err error

	// Decode the signature
	var sig []byte
	if sig, err = DecodeSegment(signature); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:63
		_go_fuzz_dep_.CoverTab[186991]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:64
		// _ = "end of CoverTab[186991]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:65
		_go_fuzz_dep_.CoverTab[186992]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:65
		// _ = "end of CoverTab[186992]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:65
	// _ = "end of CoverTab[186985]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:65
	_go_fuzz_dep_.CoverTab[186986]++

	// Get the key
	var ecdsaKey *ecdsa.PublicKey
	switch k := key.(type) {
	case *ecdsa.PublicKey:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:70
		_go_fuzz_dep_.CoverTab[186993]++
													ecdsaKey = k
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:71
		// _ = "end of CoverTab[186993]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:72
		_go_fuzz_dep_.CoverTab[186994]++
													return ErrInvalidKeyType
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:73
		// _ = "end of CoverTab[186994]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:74
	// _ = "end of CoverTab[186986]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:74
	_go_fuzz_dep_.CoverTab[186987]++

												if len(sig) != 2*m.KeySize {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:76
		_go_fuzz_dep_.CoverTab[186995]++
													return ErrECDSAVerification
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:77
		// _ = "end of CoverTab[186995]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:78
		_go_fuzz_dep_.CoverTab[186996]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:78
		// _ = "end of CoverTab[186996]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:78
	// _ = "end of CoverTab[186987]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:78
	_go_fuzz_dep_.CoverTab[186988]++

												r := big.NewInt(0).SetBytes(sig[:m.KeySize])
												s := big.NewInt(0).SetBytes(sig[m.KeySize:])

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:84
	if !m.Hash.Available() {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:84
		_go_fuzz_dep_.CoverTab[186997]++
													return ErrHashUnavailable
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:85
		// _ = "end of CoverTab[186997]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:86
		_go_fuzz_dep_.CoverTab[186998]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:86
		// _ = "end of CoverTab[186998]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:86
	// _ = "end of CoverTab[186988]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:86
	_go_fuzz_dep_.CoverTab[186989]++
												hasher := m.Hash.New()
												hasher.Write([]byte(signingString))

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:91
	if verifystatus := ecdsa.Verify(ecdsaKey, hasher.Sum(nil), r, s); verifystatus {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:91
		_go_fuzz_dep_.CoverTab[186999]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:92
		// _ = "end of CoverTab[186999]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:93
		_go_fuzz_dep_.CoverTab[187000]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:93
		// _ = "end of CoverTab[187000]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:93
	// _ = "end of CoverTab[186989]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:93
	_go_fuzz_dep_.CoverTab[186990]++

												return ErrECDSAVerification
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:95
	// _ = "end of CoverTab[186990]"
}

// Implements the Sign method from SigningMethod
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:98
// For this signing method, key must be an ecdsa.PrivateKey struct
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:100
func (m *SigningMethodECDSA) Sign(signingString string, key interface{}) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:100
	_go_fuzz_dep_.CoverTab[187001]++
	// Get the key
	var ecdsaKey *ecdsa.PrivateKey
	switch k := key.(type) {
	case *ecdsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:104
		_go_fuzz_dep_.CoverTab[187004]++
													ecdsaKey = k
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:105
		// _ = "end of CoverTab[187004]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:106
		_go_fuzz_dep_.CoverTab[187005]++
													return "", ErrInvalidKeyType
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:107
		// _ = "end of CoverTab[187005]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:108
	// _ = "end of CoverTab[187001]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:108
	_go_fuzz_dep_.CoverTab[187002]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:111
	if !m.Hash.Available() {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:111
		_go_fuzz_dep_.CoverTab[187006]++
													return "", ErrHashUnavailable
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:112
		// _ = "end of CoverTab[187006]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:113
		_go_fuzz_dep_.CoverTab[187007]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:113
		// _ = "end of CoverTab[187007]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:113
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:113
	// _ = "end of CoverTab[187002]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:113
	_go_fuzz_dep_.CoverTab[187003]++

												hasher := m.Hash.New()
												hasher.Write([]byte(signingString))

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:119
	if r, s, err := ecdsa.Sign(rand.Reader, ecdsaKey, hasher.Sum(nil)); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:119
		_go_fuzz_dep_.CoverTab[187008]++
													curveBits := ecdsaKey.Curve.Params().BitSize

													if m.CurveBits != curveBits {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:122
			_go_fuzz_dep_.CoverTab[187011]++
														return "", ErrInvalidKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:123
			// _ = "end of CoverTab[187011]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:124
			_go_fuzz_dep_.CoverTab[187012]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:124
			// _ = "end of CoverTab[187012]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:124
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:124
		// _ = "end of CoverTab[187008]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:124
		_go_fuzz_dep_.CoverTab[187009]++

													keyBytes := curveBits / 8
													if curveBits%8 > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:127
			_go_fuzz_dep_.CoverTab[187013]++
														keyBytes += 1
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:128
			// _ = "end of CoverTab[187013]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:129
			_go_fuzz_dep_.CoverTab[187014]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:129
			// _ = "end of CoverTab[187014]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:129
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:129
		// _ = "end of CoverTab[187009]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:129
		_go_fuzz_dep_.CoverTab[187010]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:134
		out := make([]byte, 2*keyBytes)
													r.FillBytes(out[0:keyBytes])
													s.FillBytes(out[keyBytes:])

													return EncodeSegment(out), nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:138
		// _ = "end of CoverTab[187010]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:139
		_go_fuzz_dep_.CoverTab[187015]++
													return "", err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:140
		// _ = "end of CoverTab[187015]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:141
	// _ = "end of CoverTab[187003]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:142
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa.go:142
var _ = _go_fuzz_dep_.CoverTab
