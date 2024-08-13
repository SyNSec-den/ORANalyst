//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:17
)

import (
	"crypto"
	"crypto/aes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"

	"gopkg.in/square/go-jose.v1/cipher"
)

// A generic RSA-based encrypter/verifier
type rsaEncrypterVerifier struct {
	publicKey *rsa.PublicKey
}

// A generic RSA-based decrypter/signer
type rsaDecrypterSigner struct {
	privateKey *rsa.PrivateKey
}

// A generic EC-based encrypter/verifier
type ecEncrypterVerifier struct {
	publicKey *ecdsa.PublicKey
}

// A key generator for ECDH-ES
type ecKeyGenerator struct {
	size		int
	algID		string
	publicKey	*ecdsa.PublicKey
}

// A generic EC-based decrypter/signer
type ecDecrypterSigner struct {
	privateKey *ecdsa.PrivateKey
}

// newRSARecipient creates recipientKeyInfo based on the given key.
func newRSARecipient(keyAlg KeyAlgorithm, publicKey *rsa.PublicKey) (recipientKeyInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:62
	_go_fuzz_dep_.CoverTab[185865]++

											switch keyAlg {
	case RSA1_5, RSA_OAEP, RSA_OAEP_256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:65
		_go_fuzz_dep_.CoverTab[185868]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:65
		// _ = "end of CoverTab[185868]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:66
		_go_fuzz_dep_.CoverTab[185869]++
												return recipientKeyInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:67
		// _ = "end of CoverTab[185869]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:68
	// _ = "end of CoverTab[185865]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:68
	_go_fuzz_dep_.CoverTab[185866]++

											if publicKey == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:70
		_go_fuzz_dep_.CoverTab[185870]++
												return recipientKeyInfo{}, errors.New("invalid public key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:71
		// _ = "end of CoverTab[185870]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:72
		_go_fuzz_dep_.CoverTab[185871]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:72
		// _ = "end of CoverTab[185871]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:72
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:72
	// _ = "end of CoverTab[185866]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:72
	_go_fuzz_dep_.CoverTab[185867]++

											return recipientKeyInfo{
		keyAlg:	keyAlg,
		keyEncrypter: &rsaEncrypterVerifier{
			publicKey: publicKey,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:79
	// _ = "end of CoverTab[185867]"
}

// newRSASigner creates a recipientSigInfo based on the given key.
func newRSASigner(sigAlg SignatureAlgorithm, privateKey *rsa.PrivateKey) (recipientSigInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:83
	_go_fuzz_dep_.CoverTab[185872]++

											switch sigAlg {
	case RS256, RS384, RS512, PS256, PS384, PS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:86
		_go_fuzz_dep_.CoverTab[185875]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:86
		// _ = "end of CoverTab[185875]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:87
		_go_fuzz_dep_.CoverTab[185876]++
												return recipientSigInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:88
		// _ = "end of CoverTab[185876]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:89
	// _ = "end of CoverTab[185872]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:89
	_go_fuzz_dep_.CoverTab[185873]++

											if privateKey == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:91
		_go_fuzz_dep_.CoverTab[185877]++
												return recipientSigInfo{}, errors.New("invalid private key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:92
		// _ = "end of CoverTab[185877]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:93
		_go_fuzz_dep_.CoverTab[185878]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:93
		// _ = "end of CoverTab[185878]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:93
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:93
	// _ = "end of CoverTab[185873]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:93
	_go_fuzz_dep_.CoverTab[185874]++

											return recipientSigInfo{
		sigAlg:	sigAlg,
		publicKey: &JsonWebKey{
			Key: &privateKey.PublicKey,
		},
		signer: &rsaDecrypterSigner{
			privateKey: privateKey,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:103
	// _ = "end of CoverTab[185874]"
}

// newECDHRecipient creates recipientKeyInfo based on the given key.
func newECDHRecipient(keyAlg KeyAlgorithm, publicKey *ecdsa.PublicKey) (recipientKeyInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:107
	_go_fuzz_dep_.CoverTab[185879]++

											switch keyAlg {
	case ECDH_ES, ECDH_ES_A128KW, ECDH_ES_A192KW, ECDH_ES_A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:110
		_go_fuzz_dep_.CoverTab[185882]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:110
		// _ = "end of CoverTab[185882]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:111
		_go_fuzz_dep_.CoverTab[185883]++
												return recipientKeyInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:112
		// _ = "end of CoverTab[185883]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:113
	// _ = "end of CoverTab[185879]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:113
	_go_fuzz_dep_.CoverTab[185880]++

											if publicKey == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:115
		_go_fuzz_dep_.CoverTab[185884]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:115
		return !publicKey.Curve.IsOnCurve(publicKey.X, publicKey.Y)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:115
		// _ = "end of CoverTab[185884]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:115
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:115
		_go_fuzz_dep_.CoverTab[185885]++
												return recipientKeyInfo{}, errors.New("invalid public key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:116
		// _ = "end of CoverTab[185885]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:117
		_go_fuzz_dep_.CoverTab[185886]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:117
		// _ = "end of CoverTab[185886]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:117
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:117
	// _ = "end of CoverTab[185880]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:117
	_go_fuzz_dep_.CoverTab[185881]++

											return recipientKeyInfo{
		keyAlg:	keyAlg,
		keyEncrypter: &ecEncrypterVerifier{
			publicKey: publicKey,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:124
	// _ = "end of CoverTab[185881]"
}

// newECDSASigner creates a recipientSigInfo based on the given key.
func newECDSASigner(sigAlg SignatureAlgorithm, privateKey *ecdsa.PrivateKey) (recipientSigInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:128
	_go_fuzz_dep_.CoverTab[185887]++

											switch sigAlg {
	case ES256, ES384, ES512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:131
		_go_fuzz_dep_.CoverTab[185890]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:131
		// _ = "end of CoverTab[185890]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:132
		_go_fuzz_dep_.CoverTab[185891]++
												return recipientSigInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:133
		// _ = "end of CoverTab[185891]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:134
	// _ = "end of CoverTab[185887]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:134
	_go_fuzz_dep_.CoverTab[185888]++

											if privateKey == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:136
		_go_fuzz_dep_.CoverTab[185892]++
												return recipientSigInfo{}, errors.New("invalid private key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:137
		// _ = "end of CoverTab[185892]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:138
		_go_fuzz_dep_.CoverTab[185893]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:138
		// _ = "end of CoverTab[185893]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:138
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:138
	// _ = "end of CoverTab[185888]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:138
	_go_fuzz_dep_.CoverTab[185889]++

											return recipientSigInfo{
		sigAlg:	sigAlg,
		publicKey: &JsonWebKey{
			Key: &privateKey.PublicKey,
		},
		signer: &ecDecrypterSigner{
			privateKey: privateKey,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:148
	// _ = "end of CoverTab[185889]"
}

// Encrypt the given payload and update the object.
func (ctx rsaEncrypterVerifier) encryptKey(cek []byte, alg KeyAlgorithm) (recipientInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:152
	_go_fuzz_dep_.CoverTab[185894]++
											encryptedKey, err := ctx.encrypt(cek, alg)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:154
		_go_fuzz_dep_.CoverTab[185896]++
												return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:155
		// _ = "end of CoverTab[185896]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:156
		_go_fuzz_dep_.CoverTab[185897]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:156
		// _ = "end of CoverTab[185897]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:156
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:156
	// _ = "end of CoverTab[185894]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:156
	_go_fuzz_dep_.CoverTab[185895]++

											return recipientInfo{
		encryptedKey:	encryptedKey,
		header:		&rawHeader{},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:161
	// _ = "end of CoverTab[185895]"
}

// Encrypt the given payload. Based on the key encryption algorithm,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:164
// this will either use RSA-PKCS1v1.5 or RSA-OAEP (with SHA-1 or SHA-256).
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:166
func (ctx rsaEncrypterVerifier) encrypt(cek []byte, alg KeyAlgorithm) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:166
	_go_fuzz_dep_.CoverTab[185898]++
											switch alg {
	case RSA1_5:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:168
		_go_fuzz_dep_.CoverTab[185900]++
												return rsa.EncryptPKCS1v15(randReader, ctx.publicKey, cek)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:169
		// _ = "end of CoverTab[185900]"
	case RSA_OAEP:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:170
		_go_fuzz_dep_.CoverTab[185901]++
												return rsa.EncryptOAEP(sha1.New(), randReader, ctx.publicKey, cek, []byte{})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:171
		// _ = "end of CoverTab[185901]"
	case RSA_OAEP_256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:172
		_go_fuzz_dep_.CoverTab[185902]++
												return rsa.EncryptOAEP(sha256.New(), randReader, ctx.publicKey, cek, []byte{})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:173
		// _ = "end of CoverTab[185902]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:173
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:173
		_go_fuzz_dep_.CoverTab[185903]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:173
		// _ = "end of CoverTab[185903]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:174
	// _ = "end of CoverTab[185898]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:174
	_go_fuzz_dep_.CoverTab[185899]++

											return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:176
	// _ = "end of CoverTab[185899]"
}

// Decrypt the given payload and return the content encryption key.
func (ctx rsaDecrypterSigner) decryptKey(headers rawHeader, recipient *recipientInfo, generator keyGenerator) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:180
	_go_fuzz_dep_.CoverTab[185904]++
											return ctx.decrypt(recipient.encryptedKey, KeyAlgorithm(headers.Alg), generator)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:181
	// _ = "end of CoverTab[185904]"
}

// Decrypt the given payload. Based on the key encryption algorithm,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:184
// this will either use RSA-PKCS1v1.5 or RSA-OAEP (with SHA-1 or SHA-256).
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:186
func (ctx rsaDecrypterSigner) decrypt(jek []byte, alg KeyAlgorithm, generator keyGenerator) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:186
	_go_fuzz_dep_.CoverTab[185905]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:189
	switch alg {
	case RSA1_5:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:190
		_go_fuzz_dep_.CoverTab[185907]++
												defer func() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:191
			_go_fuzz_dep_.CoverTab[185914]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:198
			_ = recover()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:198
			// _ = "end of CoverTab[185914]"
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:199
		// _ = "end of CoverTab[185907]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:199
		_go_fuzz_dep_.CoverTab[185908]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:202
		keyBytes := ctx.privateKey.PublicKey.N.BitLen() / 8
		if keyBytes != len(jek) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:203
			_go_fuzz_dep_.CoverTab[185915]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:207
			return nil, ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:207
			// _ = "end of CoverTab[185915]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:208
			_go_fuzz_dep_.CoverTab[185916]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:208
			// _ = "end of CoverTab[185916]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:208
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:208
		// _ = "end of CoverTab[185908]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:208
		_go_fuzz_dep_.CoverTab[185909]++

												cek, _, err := generator.genKey()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:211
			_go_fuzz_dep_.CoverTab[185917]++
													return nil, ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:212
			// _ = "end of CoverTab[185917]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:213
			_go_fuzz_dep_.CoverTab[185918]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:213
			// _ = "end of CoverTab[185918]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:213
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:213
		// _ = "end of CoverTab[185909]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:213
		_go_fuzz_dep_.CoverTab[185910]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:219
		_ = rsa.DecryptPKCS1v15SessionKey(rand.Reader, ctx.privateKey, jek, cek)

												return cek, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:221
		// _ = "end of CoverTab[185910]"
	case RSA_OAEP:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:222
		_go_fuzz_dep_.CoverTab[185911]++

												return rsa.DecryptOAEP(sha1.New(), rand.Reader, ctx.privateKey, jek, []byte{})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:224
		// _ = "end of CoverTab[185911]"
	case RSA_OAEP_256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:225
		_go_fuzz_dep_.CoverTab[185912]++

												return rsa.DecryptOAEP(sha256.New(), rand.Reader, ctx.privateKey, jek, []byte{})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:227
		// _ = "end of CoverTab[185912]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:227
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:227
		_go_fuzz_dep_.CoverTab[185913]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:227
		// _ = "end of CoverTab[185913]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:228
	// _ = "end of CoverTab[185905]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:228
	_go_fuzz_dep_.CoverTab[185906]++

											return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:230
	// _ = "end of CoverTab[185906]"
}

// Sign the given payload
func (ctx rsaDecrypterSigner) signPayload(payload []byte, alg SignatureAlgorithm) (Signature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:234
	_go_fuzz_dep_.CoverTab[185919]++
											var hash crypto.Hash

											switch alg {
	case RS256, PS256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:238
		_go_fuzz_dep_.CoverTab[185923]++
												hash = crypto.SHA256
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:239
		// _ = "end of CoverTab[185923]"
	case RS384, PS384:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:240
		_go_fuzz_dep_.CoverTab[185924]++
												hash = crypto.SHA384
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:241
		// _ = "end of CoverTab[185924]"
	case RS512, PS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:242
		_go_fuzz_dep_.CoverTab[185925]++
												hash = crypto.SHA512
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:243
		// _ = "end of CoverTab[185925]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:244
		_go_fuzz_dep_.CoverTab[185926]++
												return Signature{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:245
		// _ = "end of CoverTab[185926]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:246
	// _ = "end of CoverTab[185919]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:246
	_go_fuzz_dep_.CoverTab[185920]++

											hasher := hash.New()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:251
	_, _ = hasher.Write(payload)
	hashed := hasher.Sum(nil)

	var out []byte
	var err error

	switch alg {
	case RS256, RS384, RS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:258
		_go_fuzz_dep_.CoverTab[185927]++
												out, err = rsa.SignPKCS1v15(randReader, ctx.privateKey, hash, hashed)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:259
		// _ = "end of CoverTab[185927]"
	case PS256, PS384, PS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:260
		_go_fuzz_dep_.CoverTab[185928]++
												out, err = rsa.SignPSS(randReader, ctx.privateKey, hash, hashed, &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthAuto,
		})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:263
		// _ = "end of CoverTab[185928]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:263
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:263
		_go_fuzz_dep_.CoverTab[185929]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:263
		// _ = "end of CoverTab[185929]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:264
	// _ = "end of CoverTab[185920]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:264
	_go_fuzz_dep_.CoverTab[185921]++

											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:266
		_go_fuzz_dep_.CoverTab[185930]++
												return Signature{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:267
		// _ = "end of CoverTab[185930]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:268
		_go_fuzz_dep_.CoverTab[185931]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:268
		// _ = "end of CoverTab[185931]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:268
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:268
	// _ = "end of CoverTab[185921]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:268
	_go_fuzz_dep_.CoverTab[185922]++

											return Signature{
		Signature:	out,
		protected:	&rawHeader{},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:273
	// _ = "end of CoverTab[185922]"
}

// Verify the given payload
func (ctx rsaEncrypterVerifier) verifyPayload(payload []byte, signature []byte, alg SignatureAlgorithm) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:277
	_go_fuzz_dep_.CoverTab[185932]++
											var hash crypto.Hash

											switch alg {
	case RS256, PS256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:281
		_go_fuzz_dep_.CoverTab[185935]++
												hash = crypto.SHA256
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:282
		// _ = "end of CoverTab[185935]"
	case RS384, PS384:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:283
		_go_fuzz_dep_.CoverTab[185936]++
												hash = crypto.SHA384
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:284
		// _ = "end of CoverTab[185936]"
	case RS512, PS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:285
		_go_fuzz_dep_.CoverTab[185937]++
												hash = crypto.SHA512
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:286
		// _ = "end of CoverTab[185937]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:287
		_go_fuzz_dep_.CoverTab[185938]++
												return ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:288
		// _ = "end of CoverTab[185938]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:289
	// _ = "end of CoverTab[185932]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:289
	_go_fuzz_dep_.CoverTab[185933]++

											hasher := hash.New()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:294
	_, _ = hasher.Write(payload)
	hashed := hasher.Sum(nil)

	switch alg {
	case RS256, RS384, RS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:298
		_go_fuzz_dep_.CoverTab[185939]++
												return rsa.VerifyPKCS1v15(ctx.publicKey, hash, hashed, signature)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:299
		// _ = "end of CoverTab[185939]"
	case PS256, PS384, PS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:300
		_go_fuzz_dep_.CoverTab[185940]++
												return rsa.VerifyPSS(ctx.publicKey, hash, hashed, signature, nil)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:301
		// _ = "end of CoverTab[185940]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:301
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:301
		_go_fuzz_dep_.CoverTab[185941]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:301
		// _ = "end of CoverTab[185941]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:302
	// _ = "end of CoverTab[185933]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:302
	_go_fuzz_dep_.CoverTab[185934]++

											return ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:304
	// _ = "end of CoverTab[185934]"
}

// Encrypt the given payload and update the object.
func (ctx ecEncrypterVerifier) encryptKey(cek []byte, alg KeyAlgorithm) (recipientInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:308
	_go_fuzz_dep_.CoverTab[185942]++
											switch alg {
	case ECDH_ES:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:310
		_go_fuzz_dep_.CoverTab[185948]++

												return recipientInfo{
			header: &rawHeader{},
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:314
		// _ = "end of CoverTab[185948]"
	case ECDH_ES_A128KW, ECDH_ES_A192KW, ECDH_ES_A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:315
		_go_fuzz_dep_.CoverTab[185949]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:315
		// _ = "end of CoverTab[185949]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:316
		_go_fuzz_dep_.CoverTab[185950]++
												return recipientInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:317
		// _ = "end of CoverTab[185950]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:318
	// _ = "end of CoverTab[185942]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:318
	_go_fuzz_dep_.CoverTab[185943]++

											generator := ecKeyGenerator{
		algID:		string(alg),
		publicKey:	ctx.publicKey,
	}

	switch alg {
	case ECDH_ES_A128KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:326
		_go_fuzz_dep_.CoverTab[185951]++
												generator.size = 16
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:327
		// _ = "end of CoverTab[185951]"
	case ECDH_ES_A192KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:328
		_go_fuzz_dep_.CoverTab[185952]++
												generator.size = 24
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:329
		// _ = "end of CoverTab[185952]"
	case ECDH_ES_A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:330
		_go_fuzz_dep_.CoverTab[185953]++
												generator.size = 32
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:331
		// _ = "end of CoverTab[185953]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:331
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:331
		_go_fuzz_dep_.CoverTab[185954]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:331
		// _ = "end of CoverTab[185954]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:332
	// _ = "end of CoverTab[185943]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:332
	_go_fuzz_dep_.CoverTab[185944]++

											kek, header, err := generator.genKey()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:335
		_go_fuzz_dep_.CoverTab[185955]++
												return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:336
		// _ = "end of CoverTab[185955]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:337
		_go_fuzz_dep_.CoverTab[185956]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:337
		// _ = "end of CoverTab[185956]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:337
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:337
	// _ = "end of CoverTab[185944]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:337
	_go_fuzz_dep_.CoverTab[185945]++

											block, err := aes.NewCipher(kek)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:340
		_go_fuzz_dep_.CoverTab[185957]++
												return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:341
		// _ = "end of CoverTab[185957]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:342
		_go_fuzz_dep_.CoverTab[185958]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:342
		// _ = "end of CoverTab[185958]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:342
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:342
	// _ = "end of CoverTab[185945]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:342
	_go_fuzz_dep_.CoverTab[185946]++

											jek, err := josecipher.KeyWrap(block, cek)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:345
		_go_fuzz_dep_.CoverTab[185959]++
												return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:346
		// _ = "end of CoverTab[185959]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:347
		_go_fuzz_dep_.CoverTab[185960]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:347
		// _ = "end of CoverTab[185960]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:347
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:347
	// _ = "end of CoverTab[185946]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:347
	_go_fuzz_dep_.CoverTab[185947]++

											return recipientInfo{
		encryptedKey:	jek,
		header:		&header,
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:352
	// _ = "end of CoverTab[185947]"
}

// Get key size for EC key generator
func (ctx ecKeyGenerator) keySize() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:356
	_go_fuzz_dep_.CoverTab[185961]++
											return ctx.size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:357
	// _ = "end of CoverTab[185961]"
}

// Get a content encryption key for ECDH-ES
func (ctx ecKeyGenerator) genKey() ([]byte, rawHeader, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:361
	_go_fuzz_dep_.CoverTab[185962]++
											priv, err := ecdsa.GenerateKey(ctx.publicKey.Curve, randReader)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:363
		_go_fuzz_dep_.CoverTab[185964]++
												return nil, rawHeader{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:364
		// _ = "end of CoverTab[185964]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:365
		_go_fuzz_dep_.CoverTab[185965]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:365
		// _ = "end of CoverTab[185965]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:365
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:365
	// _ = "end of CoverTab[185962]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:365
	_go_fuzz_dep_.CoverTab[185963]++

											out := josecipher.DeriveECDHES(ctx.algID, []byte{}, []byte{}, priv, ctx.publicKey, ctx.size)

											headers := rawHeader{
		Epk: &JsonWebKey{
			Key: &priv.PublicKey,
		},
	}

											return out, headers, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:375
	// _ = "end of CoverTab[185963]"
}

// Decrypt the given payload and return the content encryption key.
func (ctx ecDecrypterSigner) decryptKey(headers rawHeader, recipient *recipientInfo, generator keyGenerator) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:379
	_go_fuzz_dep_.CoverTab[185966]++
											if headers.Epk == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:380
		_go_fuzz_dep_.CoverTab[185973]++
												return nil, errors.New("square/go-jose: missing epk header")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:381
		// _ = "end of CoverTab[185973]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:382
		_go_fuzz_dep_.CoverTab[185974]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:382
		// _ = "end of CoverTab[185974]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:382
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:382
	// _ = "end of CoverTab[185966]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:382
	_go_fuzz_dep_.CoverTab[185967]++

											publicKey, ok := headers.Epk.Key.(*ecdsa.PublicKey)
											if publicKey == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:385
		_go_fuzz_dep_.CoverTab[185975]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:385
		return !ok
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:385
		// _ = "end of CoverTab[185975]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:385
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:385
		_go_fuzz_dep_.CoverTab[185976]++
												return nil, errors.New("square/go-jose: invalid epk header")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:386
		// _ = "end of CoverTab[185976]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:387
		_go_fuzz_dep_.CoverTab[185977]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:387
		// _ = "end of CoverTab[185977]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:387
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:387
	// _ = "end of CoverTab[185967]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:387
	_go_fuzz_dep_.CoverTab[185968]++

											if !ctx.privateKey.Curve.IsOnCurve(publicKey.X, publicKey.Y) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:389
		_go_fuzz_dep_.CoverTab[185978]++
												return nil, errors.New("square/go-jose: invalid public key in epk header")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:390
		// _ = "end of CoverTab[185978]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:391
		_go_fuzz_dep_.CoverTab[185979]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:391
		// _ = "end of CoverTab[185979]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:391
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:391
	// _ = "end of CoverTab[185968]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:391
	_go_fuzz_dep_.CoverTab[185969]++

											apuData := headers.Apu.bytes()
											apvData := headers.Apv.bytes()

											deriveKey := func(algID string, size int) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:396
		_go_fuzz_dep_.CoverTab[185980]++
												return josecipher.DeriveECDHES(algID, apuData, apvData, ctx.privateKey, publicKey, size)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:397
		// _ = "end of CoverTab[185980]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:398
	// _ = "end of CoverTab[185969]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:398
	_go_fuzz_dep_.CoverTab[185970]++

											var keySize int

											switch KeyAlgorithm(headers.Alg) {
	case ECDH_ES:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:403
		_go_fuzz_dep_.CoverTab[185981]++

												return deriveKey(string(headers.Enc), generator.keySize()), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:405
		// _ = "end of CoverTab[185981]"
	case ECDH_ES_A128KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:406
		_go_fuzz_dep_.CoverTab[185982]++
												keySize = 16
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:407
		// _ = "end of CoverTab[185982]"
	case ECDH_ES_A192KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:408
		_go_fuzz_dep_.CoverTab[185983]++
												keySize = 24
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:409
		// _ = "end of CoverTab[185983]"
	case ECDH_ES_A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:410
		_go_fuzz_dep_.CoverTab[185984]++
												keySize = 32
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:411
		// _ = "end of CoverTab[185984]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:412
		_go_fuzz_dep_.CoverTab[185985]++
												return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:413
		// _ = "end of CoverTab[185985]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:414
	// _ = "end of CoverTab[185970]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:414
	_go_fuzz_dep_.CoverTab[185971]++

											key := deriveKey(headers.Alg, keySize)
											block, err := aes.NewCipher(key)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:418
		_go_fuzz_dep_.CoverTab[185986]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:419
		// _ = "end of CoverTab[185986]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:420
		_go_fuzz_dep_.CoverTab[185987]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:420
		// _ = "end of CoverTab[185987]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:420
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:420
	// _ = "end of CoverTab[185971]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:420
	_go_fuzz_dep_.CoverTab[185972]++

											return josecipher.KeyUnwrap(block, recipient.encryptedKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:422
	// _ = "end of CoverTab[185972]"
}

// Sign the given payload
func (ctx ecDecrypterSigner) signPayload(payload []byte, alg SignatureAlgorithm) (Signature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:426
	_go_fuzz_dep_.CoverTab[185988]++
											var expectedBitSize int
											var hash crypto.Hash

											switch alg {
	case ES256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:431
		_go_fuzz_dep_.CoverTab[185993]++
												expectedBitSize = 256
												hash = crypto.SHA256
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:433
		// _ = "end of CoverTab[185993]"
	case ES384:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:434
		_go_fuzz_dep_.CoverTab[185994]++
												expectedBitSize = 384
												hash = crypto.SHA384
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:436
		// _ = "end of CoverTab[185994]"
	case ES512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:437
		_go_fuzz_dep_.CoverTab[185995]++
												expectedBitSize = 521
												hash = crypto.SHA512
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:439
		// _ = "end of CoverTab[185995]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:439
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:439
		_go_fuzz_dep_.CoverTab[185996]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:439
		// _ = "end of CoverTab[185996]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:440
	// _ = "end of CoverTab[185988]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:440
	_go_fuzz_dep_.CoverTab[185989]++

											curveBits := ctx.privateKey.Curve.Params().BitSize
											if expectedBitSize != curveBits {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:443
		_go_fuzz_dep_.CoverTab[185997]++
												return Signature{}, fmt.Errorf("square/go-jose: expected %d bit key, got %d bits instead", expectedBitSize, curveBits)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:444
		// _ = "end of CoverTab[185997]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:445
		_go_fuzz_dep_.CoverTab[185998]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:445
		// _ = "end of CoverTab[185998]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:445
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:445
	// _ = "end of CoverTab[185989]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:445
	_go_fuzz_dep_.CoverTab[185990]++

											hasher := hash.New()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:450
	_, _ = hasher.Write(payload)
	hashed := hasher.Sum(nil)

	r, s, err := ecdsa.Sign(randReader, ctx.privateKey, hashed)
	if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:454
		_go_fuzz_dep_.CoverTab[185999]++
												return Signature{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:455
		// _ = "end of CoverTab[185999]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:456
		_go_fuzz_dep_.CoverTab[186000]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:456
		// _ = "end of CoverTab[186000]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:456
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:456
	// _ = "end of CoverTab[185990]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:456
	_go_fuzz_dep_.CoverTab[185991]++

											keyBytes := curveBits / 8
											if curveBits%8 > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:459
		_go_fuzz_dep_.CoverTab[186001]++
												keyBytes += 1
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:460
		// _ = "end of CoverTab[186001]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:461
		_go_fuzz_dep_.CoverTab[186002]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:461
		// _ = "end of CoverTab[186002]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:461
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:461
	// _ = "end of CoverTab[185991]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:461
	_go_fuzz_dep_.CoverTab[185992]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:466
	rBytes := r.Bytes()
	rBytesPadded := make([]byte, keyBytes)
	copy(rBytesPadded[keyBytes-len(rBytes):], rBytes)

	sBytes := s.Bytes()
	sBytesPadded := make([]byte, keyBytes)
	copy(sBytesPadded[keyBytes-len(sBytes):], sBytes)

	out := append(rBytesPadded, sBytesPadded...)

	return Signature{
		Signature:	out,
		protected:	&rawHeader{},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:479
	// _ = "end of CoverTab[185992]"
}

// Verify the given payload
func (ctx ecEncrypterVerifier) verifyPayload(payload []byte, signature []byte, alg SignatureAlgorithm) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:483
	_go_fuzz_dep_.CoverTab[186003]++
											var keySize int
											var hash crypto.Hash

											switch alg {
	case ES256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:488
		_go_fuzz_dep_.CoverTab[186007]++
												keySize = 32
												hash = crypto.SHA256
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:490
		// _ = "end of CoverTab[186007]"
	case ES384:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:491
		_go_fuzz_dep_.CoverTab[186008]++
												keySize = 48
												hash = crypto.SHA384
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:493
		// _ = "end of CoverTab[186008]"
	case ES512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:494
		_go_fuzz_dep_.CoverTab[186009]++
												keySize = 66
												hash = crypto.SHA512
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:496
		// _ = "end of CoverTab[186009]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:497
		_go_fuzz_dep_.CoverTab[186010]++
												return ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:498
		// _ = "end of CoverTab[186010]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:499
	// _ = "end of CoverTab[186003]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:499
	_go_fuzz_dep_.CoverTab[186004]++

											if len(signature) != 2*keySize {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:501
		_go_fuzz_dep_.CoverTab[186011]++
												return fmt.Errorf("square/go-jose: invalid signature size, have %d bytes, wanted %d", len(signature), 2*keySize)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:502
		// _ = "end of CoverTab[186011]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:503
		_go_fuzz_dep_.CoverTab[186012]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:503
		// _ = "end of CoverTab[186012]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:503
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:503
	// _ = "end of CoverTab[186004]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:503
	_go_fuzz_dep_.CoverTab[186005]++

											hasher := hash.New()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:508
	_, _ = hasher.Write(payload)
	hashed := hasher.Sum(nil)

	r := big.NewInt(0).SetBytes(signature[:keySize])
	s := big.NewInt(0).SetBytes(signature[keySize:])

	match := ecdsa.Verify(ctx.publicKey, hashed, r, s)
	if !match {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:515
		_go_fuzz_dep_.CoverTab[186013]++
												return errors.New("square/go-jose: ecdsa signature failed to verify")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:516
		// _ = "end of CoverTab[186013]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:517
		_go_fuzz_dep_.CoverTab[186014]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:517
		// _ = "end of CoverTab[186014]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:517
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:517
	// _ = "end of CoverTab[186005]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:517
	_go_fuzz_dep_.CoverTab[186006]++

											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:519
	// _ = "end of CoverTab[186006]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:520
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/asymmetric.go:520
var _ = _go_fuzz_dep_.CoverTab
