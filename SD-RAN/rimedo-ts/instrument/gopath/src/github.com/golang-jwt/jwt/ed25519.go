//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:1
package jwt

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:1
)

import (
	"errors"

	"crypto/ed25519"
)

var (
	ErrEd25519Verification = errors.New("ed25519: verification error")
)

// Implements the EdDSA family
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:13
// Expects ed25519.PrivateKey for signing and ed25519.PublicKey for verification
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:15
type SigningMethodEd25519 struct{}

// Specific instance for EdDSA
var (
	SigningMethodEdDSA *SigningMethodEd25519
)

func init() {
	SigningMethodEdDSA = &SigningMethodEd25519{}
	RegisterSigningMethod(SigningMethodEdDSA.Alg(), func() SigningMethod {
		return SigningMethodEdDSA
	})
}

func (m *SigningMethodEd25519) Alg() string {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:29
	_go_fuzz_dep_.CoverTab[187040]++
												return "EdDSA"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:30
	// _ = "end of CoverTab[187040]"
}

// Implements the Verify method from SigningMethod
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:33
// For this verify method, key must be an ed25519.PublicKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:35
func (m *SigningMethodEd25519) Verify(signingString, signature string, key interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:35
	_go_fuzz_dep_.CoverTab[187041]++
												var err error
												var ed25519Key ed25519.PublicKey
												var ok bool

												if ed25519Key, ok = key.(ed25519.PublicKey); !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:40
		_go_fuzz_dep_.CoverTab[187046]++
													return ErrInvalidKeyType
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:41
		// _ = "end of CoverTab[187046]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:42
		_go_fuzz_dep_.CoverTab[187047]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:42
		// _ = "end of CoverTab[187047]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:42
	// _ = "end of CoverTab[187041]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:42
	_go_fuzz_dep_.CoverTab[187042]++

												if len(ed25519Key) != ed25519.PublicKeySize {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:44
		_go_fuzz_dep_.CoverTab[187048]++
													return ErrInvalidKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:45
		// _ = "end of CoverTab[187048]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:46
		_go_fuzz_dep_.CoverTab[187049]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:46
		// _ = "end of CoverTab[187049]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:46
	// _ = "end of CoverTab[187042]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:46
	_go_fuzz_dep_.CoverTab[187043]++

	// Decode the signature
	var sig []byte
	if sig, err = DecodeSegment(signature); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:50
		_go_fuzz_dep_.CoverTab[187050]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:51
		// _ = "end of CoverTab[187050]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:52
		_go_fuzz_dep_.CoverTab[187051]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:52
		// _ = "end of CoverTab[187051]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:52
	// _ = "end of CoverTab[187043]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:52
	_go_fuzz_dep_.CoverTab[187044]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:55
	if !ed25519.Verify(ed25519Key, []byte(signingString), sig) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:55
		_go_fuzz_dep_.CoverTab[187052]++
													return ErrEd25519Verification
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:56
		// _ = "end of CoverTab[187052]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:57
		_go_fuzz_dep_.CoverTab[187053]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:57
		// _ = "end of CoverTab[187053]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:57
	// _ = "end of CoverTab[187044]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:57
	_go_fuzz_dep_.CoverTab[187045]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:59
	// _ = "end of CoverTab[187045]"
}

// Implements the Sign method from SigningMethod
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:62
// For this signing method, key must be an ed25519.PrivateKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:64
func (m *SigningMethodEd25519) Sign(signingString string, key interface{}) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:64
	_go_fuzz_dep_.CoverTab[187054]++
												var ed25519Key ed25519.PrivateKey
												var ok bool

												if ed25519Key, ok = key.(ed25519.PrivateKey); !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:68
		_go_fuzz_dep_.CoverTab[187057]++
													return "", ErrInvalidKeyType
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:69
		// _ = "end of CoverTab[187057]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:70
		_go_fuzz_dep_.CoverTab[187058]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:70
		// _ = "end of CoverTab[187058]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:70
	// _ = "end of CoverTab[187054]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:70
	_go_fuzz_dep_.CoverTab[187055]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:74
	if len(ed25519Key) != ed25519.PrivateKeySize {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:74
		_go_fuzz_dep_.CoverTab[187059]++
													return "", ErrInvalidKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:75
		// _ = "end of CoverTab[187059]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:76
		_go_fuzz_dep_.CoverTab[187060]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:76
		// _ = "end of CoverTab[187060]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:76
	// _ = "end of CoverTab[187055]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:76
	_go_fuzz_dep_.CoverTab[187056]++

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:79
	sig := ed25519.Sign(ed25519Key, []byte(signingString))
												return EncodeSegment(sig), nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:80
	// _ = "end of CoverTab[187056]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:81
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519.go:81
var _ = _go_fuzz_dep_.CoverTab
