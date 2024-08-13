//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:1
package jwt

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:1
)

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var (
	ErrKeyMustBePEMEncoded	= errors.New("Invalid Key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	ErrNotRSAPrivateKey	= errors.New("Key is not a valid RSA private key")
	ErrNotRSAPublicKey	= errors.New("Key is not a valid RSA public key")
)

// Parse PEM encoded PKCS1 or PKCS8 private key
func ParseRSAPrivateKeyFromPEM(key []byte) (*rsa.PrivateKey, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:17
	_go_fuzz_dep_.CoverTab[187273]++
													var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:22
		_go_fuzz_dep_.CoverTab[187277]++
														return nil, ErrKeyMustBePEMEncoded
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:23
		// _ = "end of CoverTab[187277]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:24
		_go_fuzz_dep_.CoverTab[187278]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:24
		// _ = "end of CoverTab[187278]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:24
	// _ = "end of CoverTab[187273]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:24
	_go_fuzz_dep_.CoverTab[187274]++

													var parsedKey interface{}
													if parsedKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:27
		_go_fuzz_dep_.CoverTab[187279]++
														if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:28
			_go_fuzz_dep_.CoverTab[187280]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:29
			// _ = "end of CoverTab[187280]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:30
			_go_fuzz_dep_.CoverTab[187281]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:30
			// _ = "end of CoverTab[187281]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:30
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:30
		// _ = "end of CoverTab[187279]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:31
		_go_fuzz_dep_.CoverTab[187282]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:31
		// _ = "end of CoverTab[187282]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:31
	// _ = "end of CoverTab[187274]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:31
	_go_fuzz_dep_.CoverTab[187275]++

													var pkey *rsa.PrivateKey
													var ok bool
													if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:35
		_go_fuzz_dep_.CoverTab[187283]++
														return nil, ErrNotRSAPrivateKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:36
		// _ = "end of CoverTab[187283]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:37
		_go_fuzz_dep_.CoverTab[187284]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:37
		// _ = "end of CoverTab[187284]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:37
	// _ = "end of CoverTab[187275]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:37
	_go_fuzz_dep_.CoverTab[187276]++

													return pkey, nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:39
	// _ = "end of CoverTab[187276]"
}

// Parse PEM encoded PKCS1 or PKCS8 private key protected with password
func ParseRSAPrivateKeyFromPEMWithPassword(key []byte, password string) (*rsa.PrivateKey, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:43
	_go_fuzz_dep_.CoverTab[187285]++
													var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:48
		_go_fuzz_dep_.CoverTab[187290]++
														return nil, ErrKeyMustBePEMEncoded
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:49
		// _ = "end of CoverTab[187290]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:50
		_go_fuzz_dep_.CoverTab[187291]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:50
		// _ = "end of CoverTab[187291]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:50
	// _ = "end of CoverTab[187285]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:50
	_go_fuzz_dep_.CoverTab[187286]++

													var parsedKey interface{}

													var blockDecrypted []byte
													if blockDecrypted, err = x509.DecryptPEMBlock(block, []byte(password)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:55
		_go_fuzz_dep_.CoverTab[187292]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:56
		// _ = "end of CoverTab[187292]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:57
		_go_fuzz_dep_.CoverTab[187293]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:57
		// _ = "end of CoverTab[187293]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:57
	// _ = "end of CoverTab[187286]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:57
	_go_fuzz_dep_.CoverTab[187287]++

													if parsedKey, err = x509.ParsePKCS1PrivateKey(blockDecrypted); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:59
		_go_fuzz_dep_.CoverTab[187294]++
														if parsedKey, err = x509.ParsePKCS8PrivateKey(blockDecrypted); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:60
			_go_fuzz_dep_.CoverTab[187295]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:61
			// _ = "end of CoverTab[187295]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:62
			_go_fuzz_dep_.CoverTab[187296]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:62
			// _ = "end of CoverTab[187296]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:62
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:62
		// _ = "end of CoverTab[187294]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:63
		_go_fuzz_dep_.CoverTab[187297]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:63
		// _ = "end of CoverTab[187297]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:63
	// _ = "end of CoverTab[187287]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:63
	_go_fuzz_dep_.CoverTab[187288]++

													var pkey *rsa.PrivateKey
													var ok bool
													if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:67
		_go_fuzz_dep_.CoverTab[187298]++
														return nil, ErrNotRSAPrivateKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:68
		// _ = "end of CoverTab[187298]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:69
		_go_fuzz_dep_.CoverTab[187299]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:69
		// _ = "end of CoverTab[187299]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:69
	// _ = "end of CoverTab[187288]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:69
	_go_fuzz_dep_.CoverTab[187289]++

													return pkey, nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:71
	// _ = "end of CoverTab[187289]"
}

// Parse PEM encoded PKCS1 or PKCS8 public key
func ParseRSAPublicKeyFromPEM(key []byte) (*rsa.PublicKey, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:75
	_go_fuzz_dep_.CoverTab[187300]++
													var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:80
		_go_fuzz_dep_.CoverTab[187304]++
														return nil, ErrKeyMustBePEMEncoded
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:81
		// _ = "end of CoverTab[187304]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:82
		_go_fuzz_dep_.CoverTab[187305]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:82
		// _ = "end of CoverTab[187305]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:82
	// _ = "end of CoverTab[187300]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:82
	_go_fuzz_dep_.CoverTab[187301]++

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:86
		_go_fuzz_dep_.CoverTab[187306]++
														if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:87
			_go_fuzz_dep_.CoverTab[187307]++
															parsedKey = cert.PublicKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:88
			// _ = "end of CoverTab[187307]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:89
			_go_fuzz_dep_.CoverTab[187308]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:90
			// _ = "end of CoverTab[187308]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:91
		// _ = "end of CoverTab[187306]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:92
		_go_fuzz_dep_.CoverTab[187309]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:92
		// _ = "end of CoverTab[187309]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:92
	// _ = "end of CoverTab[187301]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:92
	_go_fuzz_dep_.CoverTab[187302]++

													var pkey *rsa.PublicKey
													var ok bool
													if pkey, ok = parsedKey.(*rsa.PublicKey); !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:96
		_go_fuzz_dep_.CoverTab[187310]++
														return nil, ErrNotRSAPublicKey
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:97
		// _ = "end of CoverTab[187310]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:98
		_go_fuzz_dep_.CoverTab[187311]++
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:98
		// _ = "end of CoverTab[187311]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:98
	// _ = "end of CoverTab[187302]"
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:98
	_go_fuzz_dep_.CoverTab[187303]++

													return pkey, nil
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:100
	// _ = "end of CoverTab[187303]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:101
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang-jwt/jwt@v3.2.2+incompatible/rsa_utils.go:101
var _ = _go_fuzz_dep_.CoverTab
