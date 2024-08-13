//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:1
package jwt

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:1
)

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var (
	ErrNotECPublicKey	= errors.New("Key is not a valid ECDSA public key")
	ErrNotECPrivateKey	= errors.New("Key is not a valid ECDSA private key")
)

// Parse PEM encoded Elliptic Curve Private Key Structure
func ParseECPrivateKeyFromPEM(key []byte) (*ecdsa.PrivateKey, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:16
	_go_fuzz_dep_.CoverTab[187016]++
													var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:21
		_go_fuzz_dep_.CoverTab[187020]++
														return nil, ErrKeyMustBePEMEncoded
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:22
		// _ = "end of CoverTab[187020]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:23
		_go_fuzz_dep_.CoverTab[187021]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:23
		// _ = "end of CoverTab[187021]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:23
	// _ = "end of CoverTab[187016]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:23
	_go_fuzz_dep_.CoverTab[187017]++

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParseECPrivateKey(block.Bytes); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:27
		_go_fuzz_dep_.CoverTab[187022]++
														if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:28
			_go_fuzz_dep_.CoverTab[187023]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:29
			// _ = "end of CoverTab[187023]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:30
			_go_fuzz_dep_.CoverTab[187024]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:30
			// _ = "end of CoverTab[187024]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:30
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:30
		// _ = "end of CoverTab[187022]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:31
		_go_fuzz_dep_.CoverTab[187025]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:31
		// _ = "end of CoverTab[187025]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:31
	// _ = "end of CoverTab[187017]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:31
	_go_fuzz_dep_.CoverTab[187018]++

													var pkey *ecdsa.PrivateKey
													var ok bool
													if pkey, ok = parsedKey.(*ecdsa.PrivateKey); !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:35
		_go_fuzz_dep_.CoverTab[187026]++
														return nil, ErrNotECPrivateKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:36
		// _ = "end of CoverTab[187026]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:37
		_go_fuzz_dep_.CoverTab[187027]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:37
		// _ = "end of CoverTab[187027]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:37
	// _ = "end of CoverTab[187018]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:37
	_go_fuzz_dep_.CoverTab[187019]++

													return pkey, nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:39
	// _ = "end of CoverTab[187019]"
}

// Parse PEM encoded PKCS1 or PKCS8 public key
func ParseECPublicKeyFromPEM(key []byte) (*ecdsa.PublicKey, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:43
	_go_fuzz_dep_.CoverTab[187028]++
													var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:48
		_go_fuzz_dep_.CoverTab[187032]++
														return nil, ErrKeyMustBePEMEncoded
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:49
		// _ = "end of CoverTab[187032]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:50
		_go_fuzz_dep_.CoverTab[187033]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:50
		// _ = "end of CoverTab[187033]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:50
	// _ = "end of CoverTab[187028]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:50
	_go_fuzz_dep_.CoverTab[187029]++

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:54
		_go_fuzz_dep_.CoverTab[187034]++
														if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:55
			_go_fuzz_dep_.CoverTab[187035]++
															parsedKey = cert.PublicKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:56
			// _ = "end of CoverTab[187035]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:57
			_go_fuzz_dep_.CoverTab[187036]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:58
			// _ = "end of CoverTab[187036]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:59
		// _ = "end of CoverTab[187034]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:60
		_go_fuzz_dep_.CoverTab[187037]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:60
		// _ = "end of CoverTab[187037]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:60
	// _ = "end of CoverTab[187029]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:60
	_go_fuzz_dep_.CoverTab[187030]++

													var pkey *ecdsa.PublicKey
													var ok bool
													if pkey, ok = parsedKey.(*ecdsa.PublicKey); !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:64
		_go_fuzz_dep_.CoverTab[187038]++
														return nil, ErrNotECPublicKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:65
		// _ = "end of CoverTab[187038]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:66
		_go_fuzz_dep_.CoverTab[187039]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:66
		// _ = "end of CoverTab[187039]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:66
	// _ = "end of CoverTab[187030]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:66
	_go_fuzz_dep_.CoverTab[187031]++

													return pkey, nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:68
	// _ = "end of CoverTab[187031]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:69
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ecdsa_utils.go:69
var _ = _go_fuzz_dep_.CoverTab
