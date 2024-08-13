//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:1
package jwt

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:1
)

import (
	"crypto"
	"crypto/hmac"
	"errors"
)

// Implements the HMAC-SHA family of signing methods signing methods
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:9
// Expects key type of []byte for both signing and validation
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:11
type SigningMethodHMAC struct {
	Name	string
	Hash	crypto.Hash
}

// Specific instances for HS256 and company
var (
	SigningMethodHS256	*SigningMethodHMAC
	SigningMethodHS384	*SigningMethodHMAC
	SigningMethodHS512	*SigningMethodHMAC
	ErrSignatureInvalid	= errors.New("signature is invalid")
)

func init() {

	SigningMethodHS256 = &SigningMethodHMAC{"HS256", crypto.SHA256}
	RegisterSigningMethod(SigningMethodHS256.Alg(), func() SigningMethod {
		return SigningMethodHS256
	})

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:32
	SigningMethodHS384 = &SigningMethodHMAC{"HS384", crypto.SHA384}
	RegisterSigningMethod(SigningMethodHS384.Alg(), func() SigningMethod {
		return SigningMethodHS384
	})

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:38
	SigningMethodHS512 = &SigningMethodHMAC{"HS512", crypto.SHA512}
	RegisterSigningMethod(SigningMethodHS512.Alg(), func() SigningMethod {
		return SigningMethodHS512
	})
}

func (m *SigningMethodHMAC) Alg() string {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:44
	_go_fuzz_dep_.CoverTab[187088]++
												return m.Name
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:45
	// _ = "end of CoverTab[187088]"
}

// Verify the signature of HSXXX tokens.  Returns nil if the signature is valid.
func (m *SigningMethodHMAC) Verify(signingString, signature string, key interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:49
	_go_fuzz_dep_.CoverTab[187089]++

												keyBytes, ok := key.([]byte)
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:52
		_go_fuzz_dep_.CoverTab[187094]++
													return ErrInvalidKeyType
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:53
		// _ = "end of CoverTab[187094]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:54
		_go_fuzz_dep_.CoverTab[187095]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:54
		// _ = "end of CoverTab[187095]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:54
	// _ = "end of CoverTab[187089]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:54
	_go_fuzz_dep_.CoverTab[187090]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:57
	sig, err := DecodeSegment(signature)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:58
		_go_fuzz_dep_.CoverTab[187096]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:59
		// _ = "end of CoverTab[187096]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:60
		_go_fuzz_dep_.CoverTab[187097]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:60
		// _ = "end of CoverTab[187097]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:60
	// _ = "end of CoverTab[187090]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:60
	_go_fuzz_dep_.CoverTab[187091]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:63
	if !m.Hash.Available() {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:63
		_go_fuzz_dep_.CoverTab[187098]++
													return ErrHashUnavailable
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:64
		// _ = "end of CoverTab[187098]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:65
		_go_fuzz_dep_.CoverTab[187099]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:65
		// _ = "end of CoverTab[187099]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:65
	// _ = "end of CoverTab[187091]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:65
	_go_fuzz_dep_.CoverTab[187092]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:70
	hasher := hmac.New(m.Hash.New, keyBytes)
	hasher.Write([]byte(signingString))
	if !hmac.Equal(sig, hasher.Sum(nil)) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:72
		_go_fuzz_dep_.CoverTab[187100]++
													return ErrSignatureInvalid
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:73
		// _ = "end of CoverTab[187100]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:74
		_go_fuzz_dep_.CoverTab[187101]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:74
		// _ = "end of CoverTab[187101]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:74
	// _ = "end of CoverTab[187092]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:74
	_go_fuzz_dep_.CoverTab[187093]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:77
	return nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:77
	// _ = "end of CoverTab[187093]"
}

// Implements the Sign method from SigningMethod for this signing method.
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:80
// Key must be []byte
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:82
func (m *SigningMethodHMAC) Sign(signingString string, key interface{}) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:82
	_go_fuzz_dep_.CoverTab[187102]++
												if keyBytes, ok := key.([]byte); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:83
		_go_fuzz_dep_.CoverTab[187104]++
													if !m.Hash.Available() {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:84
			_go_fuzz_dep_.CoverTab[187106]++
														return "", ErrHashUnavailable
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:85
			// _ = "end of CoverTab[187106]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:86
			_go_fuzz_dep_.CoverTab[187107]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:86
			// _ = "end of CoverTab[187107]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:86
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:86
		// _ = "end of CoverTab[187104]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:86
		_go_fuzz_dep_.CoverTab[187105]++

													hasher := hmac.New(m.Hash.New, keyBytes)
													hasher.Write([]byte(signingString))

													return EncodeSegment(hasher.Sum(nil)), nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:91
		// _ = "end of CoverTab[187105]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:92
		_go_fuzz_dep_.CoverTab[187108]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:92
		// _ = "end of CoverTab[187108]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:92
	// _ = "end of CoverTab[187102]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:92
	_go_fuzz_dep_.CoverTab[187103]++

												return "", ErrInvalidKeyType
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:94
	// _ = "end of CoverTab[187103]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:95
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/hmac.go:95
var _ = _go_fuzz_dep_.CoverTab
