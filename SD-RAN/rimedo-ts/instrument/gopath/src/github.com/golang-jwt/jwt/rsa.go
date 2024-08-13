//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:1
package jwt

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:1
)

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
)

// Implements the RSA family of signing methods signing methods
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:9
// Expects *rsa.PrivateKey for signing and *rsa.PublicKey for validation
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:11
type SigningMethodRSA struct {
	Name	string
	Hash	crypto.Hash
}

// Specific instances for RS256 and company
var (
	SigningMethodRS256	*SigningMethodRSA
	SigningMethodRS384	*SigningMethodRSA
	SigningMethodRS512	*SigningMethodRSA
)

func init() {

	SigningMethodRS256 = &SigningMethodRSA{"RS256", crypto.SHA256}
	RegisterSigningMethod(SigningMethodRS256.Alg(), func() SigningMethod {
		return SigningMethodRS256
	})

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:31
	SigningMethodRS384 = &SigningMethodRSA{"RS384", crypto.SHA384}
	RegisterSigningMethod(SigningMethodRS384.Alg(), func() SigningMethod {
		return SigningMethodRS384
	})

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:37
	SigningMethodRS512 = &SigningMethodRSA{"RS512", crypto.SHA512}
	RegisterSigningMethod(SigningMethodRS512.Alg(), func() SigningMethod {
		return SigningMethodRS512
	})
}

func (m *SigningMethodRSA) Alg() string {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:43
	_go_fuzz_dep_.CoverTab[187231]++
												return m.Name
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:44
	// _ = "end of CoverTab[187231]"
}

// Implements the Verify method from SigningMethod
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:47
// For this signing method, must be an *rsa.PublicKey structure.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:49
func (m *SigningMethodRSA) Verify(signingString, signature string, key interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:49
	_go_fuzz_dep_.CoverTab[187232]++
												var err error

	// Decode the signature
	var sig []byte
	if sig, err = DecodeSegment(signature); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:54
		_go_fuzz_dep_.CoverTab[187236]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:55
		// _ = "end of CoverTab[187236]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:56
		_go_fuzz_dep_.CoverTab[187237]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:56
		// _ = "end of CoverTab[187237]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:56
	// _ = "end of CoverTab[187232]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:56
	_go_fuzz_dep_.CoverTab[187233]++

												var rsaKey *rsa.PublicKey
												var ok bool

												if rsaKey, ok = key.(*rsa.PublicKey); !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:61
		_go_fuzz_dep_.CoverTab[187238]++
													return ErrInvalidKeyType
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:62
		// _ = "end of CoverTab[187238]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:63
		_go_fuzz_dep_.CoverTab[187239]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:63
		// _ = "end of CoverTab[187239]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:63
	// _ = "end of CoverTab[187233]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:63
	_go_fuzz_dep_.CoverTab[187234]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:66
	if !m.Hash.Available() {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:66
		_go_fuzz_dep_.CoverTab[187240]++
													return ErrHashUnavailable
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:67
		// _ = "end of CoverTab[187240]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:68
		_go_fuzz_dep_.CoverTab[187241]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:68
		// _ = "end of CoverTab[187241]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:68
	// _ = "end of CoverTab[187234]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:68
	_go_fuzz_dep_.CoverTab[187235]++
												hasher := m.Hash.New()
												hasher.Write([]byte(signingString))

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:73
	return rsa.VerifyPKCS1v15(rsaKey, m.Hash, hasher.Sum(nil), sig)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:73
	// _ = "end of CoverTab[187235]"
}

// Implements the Sign method from SigningMethod
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:76
// For this signing method, must be an *rsa.PrivateKey structure.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:78
func (m *SigningMethodRSA) Sign(signingString string, key interface{}) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:78
	_go_fuzz_dep_.CoverTab[187242]++
												var rsaKey *rsa.PrivateKey
												var ok bool

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:83
	if rsaKey, ok = key.(*rsa.PrivateKey); !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:83
		_go_fuzz_dep_.CoverTab[187245]++
													return "", ErrInvalidKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:84
		// _ = "end of CoverTab[187245]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:85
		_go_fuzz_dep_.CoverTab[187246]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:85
		// _ = "end of CoverTab[187246]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:85
	// _ = "end of CoverTab[187242]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:85
	_go_fuzz_dep_.CoverTab[187243]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:88
	if !m.Hash.Available() {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:88
		_go_fuzz_dep_.CoverTab[187247]++
													return "", ErrHashUnavailable
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:89
		// _ = "end of CoverTab[187247]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:90
		_go_fuzz_dep_.CoverTab[187248]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:90
		// _ = "end of CoverTab[187248]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:90
	// _ = "end of CoverTab[187243]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:90
	_go_fuzz_dep_.CoverTab[187244]++

												hasher := m.Hash.New()
												hasher.Write([]byte(signingString))

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:96
	if sigBytes, err := rsa.SignPKCS1v15(rand.Reader, rsaKey, m.Hash, hasher.Sum(nil)); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:96
		_go_fuzz_dep_.CoverTab[187249]++
													return EncodeSegment(sigBytes), nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:97
		// _ = "end of CoverTab[187249]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:98
		_go_fuzz_dep_.CoverTab[187250]++
													return "", err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:99
		// _ = "end of CoverTab[187250]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:100
	// _ = "end of CoverTab[187244]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:101
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa.go:101
var _ = _go_fuzz_dep_.CoverTab
