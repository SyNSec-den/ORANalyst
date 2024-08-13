//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:1
package jwt

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:1
)

import (
	"crypto"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var (
	ErrNotEdPrivateKey	= errors.New("Key is not a valid Ed25519 private key")
	ErrNotEdPublicKey	= errors.New("Key is not a valid Ed25519 public key")
)

// Parse PEM-encoded Edwards curve private key
func ParseEdPrivateKeyFromPEM(key []byte) (crypto.PrivateKey, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:17
	_go_fuzz_dep_.CoverTab[187061]++
													var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:22
		_go_fuzz_dep_.CoverTab[187065]++
														return nil, ErrKeyMustBePEMEncoded
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:23
		// _ = "end of CoverTab[187065]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:24
		_go_fuzz_dep_.CoverTab[187066]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:24
		// _ = "end of CoverTab[187066]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:24
	// _ = "end of CoverTab[187061]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:24
	_go_fuzz_dep_.CoverTab[187062]++

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:28
		_go_fuzz_dep_.CoverTab[187067]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:29
		// _ = "end of CoverTab[187067]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:30
		_go_fuzz_dep_.CoverTab[187068]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:30
		// _ = "end of CoverTab[187068]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:30
	// _ = "end of CoverTab[187062]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:30
	_go_fuzz_dep_.CoverTab[187063]++

													var pkey ed25519.PrivateKey
													var ok bool
													if pkey, ok = parsedKey.(ed25519.PrivateKey); !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:34
		_go_fuzz_dep_.CoverTab[187069]++
														return nil, ErrNotEdPrivateKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:35
		// _ = "end of CoverTab[187069]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:36
		_go_fuzz_dep_.CoverTab[187070]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:36
		// _ = "end of CoverTab[187070]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:36
	// _ = "end of CoverTab[187063]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:36
	_go_fuzz_dep_.CoverTab[187064]++

													return pkey, nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:38
	// _ = "end of CoverTab[187064]"
}

// Parse PEM-encoded Edwards curve public key
func ParseEdPublicKeyFromPEM(key []byte) (crypto.PublicKey, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:42
	_go_fuzz_dep_.CoverTab[187071]++
													var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:47
		_go_fuzz_dep_.CoverTab[187075]++
														return nil, ErrKeyMustBePEMEncoded
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:48
		// _ = "end of CoverTab[187075]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:49
		_go_fuzz_dep_.CoverTab[187076]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:49
		// _ = "end of CoverTab[187076]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:49
	// _ = "end of CoverTab[187071]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:49
	_go_fuzz_dep_.CoverTab[187072]++

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:53
		_go_fuzz_dep_.CoverTab[187077]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:54
		// _ = "end of CoverTab[187077]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:55
		_go_fuzz_dep_.CoverTab[187078]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:55
		// _ = "end of CoverTab[187078]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:55
	// _ = "end of CoverTab[187072]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:55
	_go_fuzz_dep_.CoverTab[187073]++

													var pkey ed25519.PublicKey
													var ok bool
													if pkey, ok = parsedKey.(ed25519.PublicKey); !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:59
		_go_fuzz_dep_.CoverTab[187079]++
														return nil, ErrNotEdPublicKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:60
		// _ = "end of CoverTab[187079]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:61
		_go_fuzz_dep_.CoverTab[187080]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:61
		// _ = "end of CoverTab[187080]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:61
	// _ = "end of CoverTab[187073]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:61
	_go_fuzz_dep_.CoverTab[187074]++

													return pkey, nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:63
	// _ = "end of CoverTab[187074]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:64
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/ed25519_utils.go:64
var _ = _go_fuzz_dep_.CoverTab
