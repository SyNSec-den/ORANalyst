//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:17
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

	"golang.org/x/crypto/ed25519"
	josecipher "gopkg.in/square/go-jose.v2/cipher"
	"gopkg.in/square/go-jose.v2/json"
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

type edEncrypterVerifier struct {
	publicKey ed25519.PublicKey
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

type edDecrypterSigner struct {
	privateKey ed25519.PrivateKey
}

// newRSARecipient creates recipientKeyInfo based on the given key.
func newRSARecipient(keyAlg KeyAlgorithm, publicKey *rsa.PublicKey) (recipientKeyInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:72
	_go_fuzz_dep_.CoverTab[188923]++

											switch keyAlg {
	case RSA1_5, RSA_OAEP, RSA_OAEP_256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:75
		_go_fuzz_dep_.CoverTab[188926]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:75
		// _ = "end of CoverTab[188926]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:76
		_go_fuzz_dep_.CoverTab[188927]++
												return recipientKeyInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:77
		// _ = "end of CoverTab[188927]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:78
	// _ = "end of CoverTab[188923]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:78
	_go_fuzz_dep_.CoverTab[188924]++

											if publicKey == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:80
		_go_fuzz_dep_.CoverTab[188928]++
												return recipientKeyInfo{}, errors.New("invalid public key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:81
		// _ = "end of CoverTab[188928]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:82
		_go_fuzz_dep_.CoverTab[188929]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:82
		// _ = "end of CoverTab[188929]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:82
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:82
	// _ = "end of CoverTab[188924]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:82
	_go_fuzz_dep_.CoverTab[188925]++

											return recipientKeyInfo{
		keyAlg:	keyAlg,
		keyEncrypter: &rsaEncrypterVerifier{
			publicKey: publicKey,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:89
	// _ = "end of CoverTab[188925]"
}

// newRSASigner creates a recipientSigInfo based on the given key.
func newRSASigner(sigAlg SignatureAlgorithm, privateKey *rsa.PrivateKey) (recipientSigInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:93
	_go_fuzz_dep_.CoverTab[188930]++

											switch sigAlg {
	case RS256, RS384, RS512, PS256, PS384, PS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:96
		_go_fuzz_dep_.CoverTab[188933]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:96
		// _ = "end of CoverTab[188933]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:97
		_go_fuzz_dep_.CoverTab[188934]++
												return recipientSigInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:98
		// _ = "end of CoverTab[188934]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:99
	// _ = "end of CoverTab[188930]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:99
	_go_fuzz_dep_.CoverTab[188931]++

											if privateKey == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:101
		_go_fuzz_dep_.CoverTab[188935]++
												return recipientSigInfo{}, errors.New("invalid private key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:102
		// _ = "end of CoverTab[188935]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:103
		_go_fuzz_dep_.CoverTab[188936]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:103
		// _ = "end of CoverTab[188936]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:103
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:103
	// _ = "end of CoverTab[188931]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:103
	_go_fuzz_dep_.CoverTab[188932]++

											return recipientSigInfo{
		sigAlg:	sigAlg,
		publicKey: staticPublicKey(&JSONWebKey{
			Key: privateKey.Public(),
		}),
		signer: &rsaDecrypterSigner{
			privateKey: privateKey,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:113
	// _ = "end of CoverTab[188932]"
}

func newEd25519Signer(sigAlg SignatureAlgorithm, privateKey ed25519.PrivateKey) (recipientSigInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:116
	_go_fuzz_dep_.CoverTab[188937]++
											if sigAlg != EdDSA {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:117
		_go_fuzz_dep_.CoverTab[188940]++
												return recipientSigInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:118
		// _ = "end of CoverTab[188940]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:119
		_go_fuzz_dep_.CoverTab[188941]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:119
		// _ = "end of CoverTab[188941]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:119
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:119
	// _ = "end of CoverTab[188937]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:119
	_go_fuzz_dep_.CoverTab[188938]++

											if privateKey == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:121
		_go_fuzz_dep_.CoverTab[188942]++
												return recipientSigInfo{}, errors.New("invalid private key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:122
		// _ = "end of CoverTab[188942]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:123
		_go_fuzz_dep_.CoverTab[188943]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:123
		// _ = "end of CoverTab[188943]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:123
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:123
	// _ = "end of CoverTab[188938]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:123
	_go_fuzz_dep_.CoverTab[188939]++
											return recipientSigInfo{
		sigAlg:	sigAlg,
		publicKey: staticPublicKey(&JSONWebKey{
			Key: privateKey.Public(),
		}),
		signer: &edDecrypterSigner{
			privateKey: privateKey,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:132
	// _ = "end of CoverTab[188939]"
}

// newECDHRecipient creates recipientKeyInfo based on the given key.
func newECDHRecipient(keyAlg KeyAlgorithm, publicKey *ecdsa.PublicKey) (recipientKeyInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:136
	_go_fuzz_dep_.CoverTab[188944]++

											switch keyAlg {
	case ECDH_ES, ECDH_ES_A128KW, ECDH_ES_A192KW, ECDH_ES_A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:139
		_go_fuzz_dep_.CoverTab[188947]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:139
		// _ = "end of CoverTab[188947]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:140
		_go_fuzz_dep_.CoverTab[188948]++
												return recipientKeyInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:141
		// _ = "end of CoverTab[188948]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:142
	// _ = "end of CoverTab[188944]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:142
	_go_fuzz_dep_.CoverTab[188945]++

											if publicKey == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:144
		_go_fuzz_dep_.CoverTab[188949]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:144
		return !publicKey.Curve.IsOnCurve(publicKey.X, publicKey.Y)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:144
		// _ = "end of CoverTab[188949]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:144
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:144
		_go_fuzz_dep_.CoverTab[188950]++
												return recipientKeyInfo{}, errors.New("invalid public key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:145
		// _ = "end of CoverTab[188950]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:146
		_go_fuzz_dep_.CoverTab[188951]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:146
		// _ = "end of CoverTab[188951]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:146
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:146
	// _ = "end of CoverTab[188945]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:146
	_go_fuzz_dep_.CoverTab[188946]++

											return recipientKeyInfo{
		keyAlg:	keyAlg,
		keyEncrypter: &ecEncrypterVerifier{
			publicKey: publicKey,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:153
	// _ = "end of CoverTab[188946]"
}

// newECDSASigner creates a recipientSigInfo based on the given key.
func newECDSASigner(sigAlg SignatureAlgorithm, privateKey *ecdsa.PrivateKey) (recipientSigInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:157
	_go_fuzz_dep_.CoverTab[188952]++

											switch sigAlg {
	case ES256, ES384, ES512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:160
		_go_fuzz_dep_.CoverTab[188955]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:160
		// _ = "end of CoverTab[188955]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:161
		_go_fuzz_dep_.CoverTab[188956]++
												return recipientSigInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:162
		// _ = "end of CoverTab[188956]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:163
	// _ = "end of CoverTab[188952]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:163
	_go_fuzz_dep_.CoverTab[188953]++

											if privateKey == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:165
		_go_fuzz_dep_.CoverTab[188957]++
												return recipientSigInfo{}, errors.New("invalid private key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:166
		// _ = "end of CoverTab[188957]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:167
		_go_fuzz_dep_.CoverTab[188958]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:167
		// _ = "end of CoverTab[188958]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:167
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:167
	// _ = "end of CoverTab[188953]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:167
	_go_fuzz_dep_.CoverTab[188954]++

											return recipientSigInfo{
		sigAlg:	sigAlg,
		publicKey: staticPublicKey(&JSONWebKey{
			Key: privateKey.Public(),
		}),
		signer: &ecDecrypterSigner{
			privateKey: privateKey,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:177
	// _ = "end of CoverTab[188954]"
}

// Encrypt the given payload and update the object.
func (ctx rsaEncrypterVerifier) encryptKey(cek []byte, alg KeyAlgorithm) (recipientInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:181
	_go_fuzz_dep_.CoverTab[188959]++
											encryptedKey, err := ctx.encrypt(cek, alg)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:183
		_go_fuzz_dep_.CoverTab[188961]++
												return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:184
		// _ = "end of CoverTab[188961]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:185
		_go_fuzz_dep_.CoverTab[188962]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:185
		// _ = "end of CoverTab[188962]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:185
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:185
	// _ = "end of CoverTab[188959]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:185
	_go_fuzz_dep_.CoverTab[188960]++

											return recipientInfo{
		encryptedKey:	encryptedKey,
		header:		&rawHeader{},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:190
	// _ = "end of CoverTab[188960]"
}

// Encrypt the given payload. Based on the key encryption algorithm,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:193
// this will either use RSA-PKCS1v1.5 or RSA-OAEP (with SHA-1 or SHA-256).
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:195
func (ctx rsaEncrypterVerifier) encrypt(cek []byte, alg KeyAlgorithm) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:195
	_go_fuzz_dep_.CoverTab[188963]++
											switch alg {
	case RSA1_5:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:197
		_go_fuzz_dep_.CoverTab[188965]++
												return rsa.EncryptPKCS1v15(RandReader, ctx.publicKey, cek)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:198
		// _ = "end of CoverTab[188965]"
	case RSA_OAEP:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:199
		_go_fuzz_dep_.CoverTab[188966]++
												return rsa.EncryptOAEP(sha1.New(), RandReader, ctx.publicKey, cek, []byte{})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:200
		// _ = "end of CoverTab[188966]"
	case RSA_OAEP_256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:201
		_go_fuzz_dep_.CoverTab[188967]++
												return rsa.EncryptOAEP(sha256.New(), RandReader, ctx.publicKey, cek, []byte{})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:202
		// _ = "end of CoverTab[188967]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:202
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:202
		_go_fuzz_dep_.CoverTab[188968]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:202
		// _ = "end of CoverTab[188968]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:203
	// _ = "end of CoverTab[188963]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:203
	_go_fuzz_dep_.CoverTab[188964]++

											return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:205
	// _ = "end of CoverTab[188964]"
}

// Decrypt the given payload and return the content encryption key.
func (ctx rsaDecrypterSigner) decryptKey(headers rawHeader, recipient *recipientInfo, generator keyGenerator) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:209
	_go_fuzz_dep_.CoverTab[188969]++
											return ctx.decrypt(recipient.encryptedKey, headers.getAlgorithm(), generator)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:210
	// _ = "end of CoverTab[188969]"
}

// Decrypt the given payload. Based on the key encryption algorithm,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:213
// this will either use RSA-PKCS1v1.5 or RSA-OAEP (with SHA-1 or SHA-256).
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:215
func (ctx rsaDecrypterSigner) decrypt(jek []byte, alg KeyAlgorithm, generator keyGenerator) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:215
	_go_fuzz_dep_.CoverTab[188970]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:218
	switch alg {
	case RSA1_5:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:219
		_go_fuzz_dep_.CoverTab[188972]++
												defer func() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:220
			_go_fuzz_dep_.CoverTab[188979]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:227
			_ = recover()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:227
			// _ = "end of CoverTab[188979]"
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:228
		// _ = "end of CoverTab[188972]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:228
		_go_fuzz_dep_.CoverTab[188973]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:231
		keyBytes := ctx.privateKey.PublicKey.N.BitLen() / 8
		if keyBytes != len(jek) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:232
			_go_fuzz_dep_.CoverTab[188980]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:236
			return nil, ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:236
			// _ = "end of CoverTab[188980]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:237
			_go_fuzz_dep_.CoverTab[188981]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:237
			// _ = "end of CoverTab[188981]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:237
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:237
		// _ = "end of CoverTab[188973]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:237
		_go_fuzz_dep_.CoverTab[188974]++

												cek, _, err := generator.genKey()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:240
			_go_fuzz_dep_.CoverTab[188982]++
													return nil, ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:241
			// _ = "end of CoverTab[188982]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:242
			_go_fuzz_dep_.CoverTab[188983]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:242
			// _ = "end of CoverTab[188983]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:242
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:242
		// _ = "end of CoverTab[188974]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:242
		_go_fuzz_dep_.CoverTab[188975]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:248
		_ = rsa.DecryptPKCS1v15SessionKey(rand.Reader, ctx.privateKey, jek, cek)

												return cek, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:250
		// _ = "end of CoverTab[188975]"
	case RSA_OAEP:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:251
		_go_fuzz_dep_.CoverTab[188976]++

												return rsa.DecryptOAEP(sha1.New(), rand.Reader, ctx.privateKey, jek, []byte{})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:253
		// _ = "end of CoverTab[188976]"
	case RSA_OAEP_256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:254
		_go_fuzz_dep_.CoverTab[188977]++

												return rsa.DecryptOAEP(sha256.New(), rand.Reader, ctx.privateKey, jek, []byte{})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:256
		// _ = "end of CoverTab[188977]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:256
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:256
		_go_fuzz_dep_.CoverTab[188978]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:256
		// _ = "end of CoverTab[188978]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:257
	// _ = "end of CoverTab[188970]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:257
	_go_fuzz_dep_.CoverTab[188971]++

											return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:259
	// _ = "end of CoverTab[188971]"
}

// Sign the given payload
func (ctx rsaDecrypterSigner) signPayload(payload []byte, alg SignatureAlgorithm) (Signature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:263
	_go_fuzz_dep_.CoverTab[188984]++
											var hash crypto.Hash

											switch alg {
	case RS256, PS256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:267
		_go_fuzz_dep_.CoverTab[188988]++
												hash = crypto.SHA256
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:268
		// _ = "end of CoverTab[188988]"
	case RS384, PS384:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:269
		_go_fuzz_dep_.CoverTab[188989]++
												hash = crypto.SHA384
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:270
		// _ = "end of CoverTab[188989]"
	case RS512, PS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:271
		_go_fuzz_dep_.CoverTab[188990]++
												hash = crypto.SHA512
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:272
		// _ = "end of CoverTab[188990]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:273
		_go_fuzz_dep_.CoverTab[188991]++
												return Signature{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:274
		// _ = "end of CoverTab[188991]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:275
	// _ = "end of CoverTab[188984]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:275
	_go_fuzz_dep_.CoverTab[188985]++

											hasher := hash.New()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:280
	_, _ = hasher.Write(payload)
	hashed := hasher.Sum(nil)

	var out []byte
	var err error

	switch alg {
	case RS256, RS384, RS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:287
		_go_fuzz_dep_.CoverTab[188992]++
												out, err = rsa.SignPKCS1v15(RandReader, ctx.privateKey, hash, hashed)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:288
		// _ = "end of CoverTab[188992]"
	case PS256, PS384, PS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:289
		_go_fuzz_dep_.CoverTab[188993]++
												out, err = rsa.SignPSS(RandReader, ctx.privateKey, hash, hashed, &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthEqualsHash,
		})
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:292
		// _ = "end of CoverTab[188993]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:292
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:292
		_go_fuzz_dep_.CoverTab[188994]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:292
		// _ = "end of CoverTab[188994]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:293
	// _ = "end of CoverTab[188985]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:293
	_go_fuzz_dep_.CoverTab[188986]++

											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:295
		_go_fuzz_dep_.CoverTab[188995]++
												return Signature{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:296
		// _ = "end of CoverTab[188995]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:297
		_go_fuzz_dep_.CoverTab[188996]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:297
		// _ = "end of CoverTab[188996]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:297
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:297
	// _ = "end of CoverTab[188986]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:297
	_go_fuzz_dep_.CoverTab[188987]++

											return Signature{
		Signature:	out,
		protected:	&rawHeader{},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:302
	// _ = "end of CoverTab[188987]"
}

// Verify the given payload
func (ctx rsaEncrypterVerifier) verifyPayload(payload []byte, signature []byte, alg SignatureAlgorithm) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:306
	_go_fuzz_dep_.CoverTab[188997]++
											var hash crypto.Hash

											switch alg {
	case RS256, PS256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:310
		_go_fuzz_dep_.CoverTab[189000]++
												hash = crypto.SHA256
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:311
		// _ = "end of CoverTab[189000]"
	case RS384, PS384:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:312
		_go_fuzz_dep_.CoverTab[189001]++
												hash = crypto.SHA384
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:313
		// _ = "end of CoverTab[189001]"
	case RS512, PS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:314
		_go_fuzz_dep_.CoverTab[189002]++
												hash = crypto.SHA512
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:315
		// _ = "end of CoverTab[189002]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:316
		_go_fuzz_dep_.CoverTab[189003]++
												return ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:317
		// _ = "end of CoverTab[189003]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:318
	// _ = "end of CoverTab[188997]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:318
	_go_fuzz_dep_.CoverTab[188998]++

											hasher := hash.New()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:323
	_, _ = hasher.Write(payload)
	hashed := hasher.Sum(nil)

	switch alg {
	case RS256, RS384, RS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:327
		_go_fuzz_dep_.CoverTab[189004]++
												return rsa.VerifyPKCS1v15(ctx.publicKey, hash, hashed, signature)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:328
		// _ = "end of CoverTab[189004]"
	case PS256, PS384, PS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:329
		_go_fuzz_dep_.CoverTab[189005]++
												return rsa.VerifyPSS(ctx.publicKey, hash, hashed, signature, nil)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:330
		// _ = "end of CoverTab[189005]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:330
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:330
		_go_fuzz_dep_.CoverTab[189006]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:330
		// _ = "end of CoverTab[189006]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:331
	// _ = "end of CoverTab[188998]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:331
	_go_fuzz_dep_.CoverTab[188999]++

											return ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:333
	// _ = "end of CoverTab[188999]"
}

// Encrypt the given payload and update the object.
func (ctx ecEncrypterVerifier) encryptKey(cek []byte, alg KeyAlgorithm) (recipientInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:337
	_go_fuzz_dep_.CoverTab[189007]++
											switch alg {
	case ECDH_ES:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:339
		_go_fuzz_dep_.CoverTab[189013]++

												return recipientInfo{
			header: &rawHeader{},
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:343
		// _ = "end of CoverTab[189013]"
	case ECDH_ES_A128KW, ECDH_ES_A192KW, ECDH_ES_A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:344
		_go_fuzz_dep_.CoverTab[189014]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:344
		// _ = "end of CoverTab[189014]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:345
		_go_fuzz_dep_.CoverTab[189015]++
												return recipientInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:346
		// _ = "end of CoverTab[189015]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:347
	// _ = "end of CoverTab[189007]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:347
	_go_fuzz_dep_.CoverTab[189008]++

											generator := ecKeyGenerator{
		algID:		string(alg),
		publicKey:	ctx.publicKey,
	}

	switch alg {
	case ECDH_ES_A128KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:355
		_go_fuzz_dep_.CoverTab[189016]++
												generator.size = 16
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:356
		// _ = "end of CoverTab[189016]"
	case ECDH_ES_A192KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:357
		_go_fuzz_dep_.CoverTab[189017]++
												generator.size = 24
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:358
		// _ = "end of CoverTab[189017]"
	case ECDH_ES_A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:359
		_go_fuzz_dep_.CoverTab[189018]++
												generator.size = 32
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:360
		// _ = "end of CoverTab[189018]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:360
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:360
		_go_fuzz_dep_.CoverTab[189019]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:360
		// _ = "end of CoverTab[189019]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:361
	// _ = "end of CoverTab[189008]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:361
	_go_fuzz_dep_.CoverTab[189009]++

											kek, header, err := generator.genKey()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:364
		_go_fuzz_dep_.CoverTab[189020]++
												return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:365
		// _ = "end of CoverTab[189020]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:366
		_go_fuzz_dep_.CoverTab[189021]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:366
		// _ = "end of CoverTab[189021]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:366
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:366
	// _ = "end of CoverTab[189009]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:366
	_go_fuzz_dep_.CoverTab[189010]++

											block, err := aes.NewCipher(kek)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:369
		_go_fuzz_dep_.CoverTab[189022]++
												return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:370
		// _ = "end of CoverTab[189022]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:371
		_go_fuzz_dep_.CoverTab[189023]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:371
		// _ = "end of CoverTab[189023]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:371
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:371
	// _ = "end of CoverTab[189010]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:371
	_go_fuzz_dep_.CoverTab[189011]++

											jek, err := josecipher.KeyWrap(block, cek)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:374
		_go_fuzz_dep_.CoverTab[189024]++
												return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:375
		// _ = "end of CoverTab[189024]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:376
		_go_fuzz_dep_.CoverTab[189025]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:376
		// _ = "end of CoverTab[189025]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:376
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:376
	// _ = "end of CoverTab[189011]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:376
	_go_fuzz_dep_.CoverTab[189012]++

											return recipientInfo{
		encryptedKey:	jek,
		header:		&header,
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:381
	// _ = "end of CoverTab[189012]"
}

// Get key size for EC key generator
func (ctx ecKeyGenerator) keySize() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:385
	_go_fuzz_dep_.CoverTab[189026]++
											return ctx.size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:386
	// _ = "end of CoverTab[189026]"
}

// Get a content encryption key for ECDH-ES
func (ctx ecKeyGenerator) genKey() ([]byte, rawHeader, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:390
	_go_fuzz_dep_.CoverTab[189027]++
											priv, err := ecdsa.GenerateKey(ctx.publicKey.Curve, RandReader)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:392
		_go_fuzz_dep_.CoverTab[189030]++
												return nil, rawHeader{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:393
		// _ = "end of CoverTab[189030]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:394
		_go_fuzz_dep_.CoverTab[189031]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:394
		// _ = "end of CoverTab[189031]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:394
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:394
	// _ = "end of CoverTab[189027]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:394
	_go_fuzz_dep_.CoverTab[189028]++

											out := josecipher.DeriveECDHES(ctx.algID, []byte{}, []byte{}, priv, ctx.publicKey, ctx.size)

											b, err := json.Marshal(&JSONWebKey{
		Key: &priv.PublicKey,
	})
	if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:401
		_go_fuzz_dep_.CoverTab[189032]++
												return nil, nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:402
		// _ = "end of CoverTab[189032]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:403
		_go_fuzz_dep_.CoverTab[189033]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:403
		// _ = "end of CoverTab[189033]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:403
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:403
	// _ = "end of CoverTab[189028]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:403
	_go_fuzz_dep_.CoverTab[189029]++

											headers := rawHeader{
		headerEPK: makeRawMessage(b),
	}

											return out, headers, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:409
	// _ = "end of CoverTab[189029]"
}

// Decrypt the given payload and return the content encryption key.
func (ctx ecDecrypterSigner) decryptKey(headers rawHeader, recipient *recipientInfo, generator keyGenerator) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:413
	_go_fuzz_dep_.CoverTab[189034]++
											epk, err := headers.getEPK()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:415
		_go_fuzz_dep_.CoverTab[189044]++
												return nil, errors.New("square/go-jose: invalid epk header")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:416
		// _ = "end of CoverTab[189044]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:417
		_go_fuzz_dep_.CoverTab[189045]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:417
		// _ = "end of CoverTab[189045]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:417
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:417
	// _ = "end of CoverTab[189034]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:417
	_go_fuzz_dep_.CoverTab[189035]++
											if epk == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:418
		_go_fuzz_dep_.CoverTab[189046]++
												return nil, errors.New("square/go-jose: missing epk header")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:419
		// _ = "end of CoverTab[189046]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:420
		_go_fuzz_dep_.CoverTab[189047]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:420
		// _ = "end of CoverTab[189047]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:420
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:420
	// _ = "end of CoverTab[189035]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:420
	_go_fuzz_dep_.CoverTab[189036]++

											publicKey, ok := epk.Key.(*ecdsa.PublicKey)
											if publicKey == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:423
		_go_fuzz_dep_.CoverTab[189048]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:423
		return !ok
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:423
		// _ = "end of CoverTab[189048]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:423
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:423
		_go_fuzz_dep_.CoverTab[189049]++
												return nil, errors.New("square/go-jose: invalid epk header")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:424
		// _ = "end of CoverTab[189049]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:425
		_go_fuzz_dep_.CoverTab[189050]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:425
		// _ = "end of CoverTab[189050]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:425
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:425
	// _ = "end of CoverTab[189036]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:425
	_go_fuzz_dep_.CoverTab[189037]++

											if !ctx.privateKey.Curve.IsOnCurve(publicKey.X, publicKey.Y) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:427
		_go_fuzz_dep_.CoverTab[189051]++
												return nil, errors.New("square/go-jose: invalid public key in epk header")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:428
		// _ = "end of CoverTab[189051]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:429
		_go_fuzz_dep_.CoverTab[189052]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:429
		// _ = "end of CoverTab[189052]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:429
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:429
	// _ = "end of CoverTab[189037]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:429
	_go_fuzz_dep_.CoverTab[189038]++

											apuData, err := headers.getAPU()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:432
		_go_fuzz_dep_.CoverTab[189053]++
												return nil, errors.New("square/go-jose: invalid apu header")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:433
		// _ = "end of CoverTab[189053]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:434
		_go_fuzz_dep_.CoverTab[189054]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:434
		// _ = "end of CoverTab[189054]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:434
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:434
	// _ = "end of CoverTab[189038]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:434
	_go_fuzz_dep_.CoverTab[189039]++
											apvData, err := headers.getAPV()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:436
		_go_fuzz_dep_.CoverTab[189055]++
												return nil, errors.New("square/go-jose: invalid apv header")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:437
		// _ = "end of CoverTab[189055]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:438
		_go_fuzz_dep_.CoverTab[189056]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:438
		// _ = "end of CoverTab[189056]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:438
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:438
	// _ = "end of CoverTab[189039]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:438
	_go_fuzz_dep_.CoverTab[189040]++

											deriveKey := func(algID string, size int) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:440
		_go_fuzz_dep_.CoverTab[189057]++
												return josecipher.DeriveECDHES(algID, apuData.bytes(), apvData.bytes(), ctx.privateKey, publicKey, size)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:441
		// _ = "end of CoverTab[189057]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:442
	// _ = "end of CoverTab[189040]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:442
	_go_fuzz_dep_.CoverTab[189041]++

											var keySize int

											algorithm := headers.getAlgorithm()
											switch algorithm {
	case ECDH_ES:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:448
		_go_fuzz_dep_.CoverTab[189058]++

												return deriveKey(string(headers.getEncryption()), generator.keySize()), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:450
		// _ = "end of CoverTab[189058]"
	case ECDH_ES_A128KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:451
		_go_fuzz_dep_.CoverTab[189059]++
												keySize = 16
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:452
		// _ = "end of CoverTab[189059]"
	case ECDH_ES_A192KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:453
		_go_fuzz_dep_.CoverTab[189060]++
												keySize = 24
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:454
		// _ = "end of CoverTab[189060]"
	case ECDH_ES_A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:455
		_go_fuzz_dep_.CoverTab[189061]++
												keySize = 32
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:456
		// _ = "end of CoverTab[189061]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:457
		_go_fuzz_dep_.CoverTab[189062]++
												return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:458
		// _ = "end of CoverTab[189062]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:459
	// _ = "end of CoverTab[189041]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:459
	_go_fuzz_dep_.CoverTab[189042]++

											key := deriveKey(string(algorithm), keySize)
											block, err := aes.NewCipher(key)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:463
		_go_fuzz_dep_.CoverTab[189063]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:464
		// _ = "end of CoverTab[189063]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:465
		_go_fuzz_dep_.CoverTab[189064]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:465
		// _ = "end of CoverTab[189064]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:465
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:465
	// _ = "end of CoverTab[189042]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:465
	_go_fuzz_dep_.CoverTab[189043]++

											return josecipher.KeyUnwrap(block, recipient.encryptedKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:467
	// _ = "end of CoverTab[189043]"
}

func (ctx edDecrypterSigner) signPayload(payload []byte, alg SignatureAlgorithm) (Signature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:470
	_go_fuzz_dep_.CoverTab[189065]++
											if alg != EdDSA {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:471
		_go_fuzz_dep_.CoverTab[189068]++
												return Signature{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:472
		// _ = "end of CoverTab[189068]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:473
		_go_fuzz_dep_.CoverTab[189069]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:473
		// _ = "end of CoverTab[189069]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:473
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:473
	// _ = "end of CoverTab[189065]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:473
	_go_fuzz_dep_.CoverTab[189066]++

											sig, err := ctx.privateKey.Sign(RandReader, payload, crypto.Hash(0))
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:476
		_go_fuzz_dep_.CoverTab[189070]++
												return Signature{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:477
		// _ = "end of CoverTab[189070]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:478
		_go_fuzz_dep_.CoverTab[189071]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:478
		// _ = "end of CoverTab[189071]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:478
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:478
	// _ = "end of CoverTab[189066]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:478
	_go_fuzz_dep_.CoverTab[189067]++

											return Signature{
		Signature:	sig,
		protected:	&rawHeader{},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:483
	// _ = "end of CoverTab[189067]"
}

func (ctx edEncrypterVerifier) verifyPayload(payload []byte, signature []byte, alg SignatureAlgorithm) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:486
	_go_fuzz_dep_.CoverTab[189072]++
											if alg != EdDSA {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:487
		_go_fuzz_dep_.CoverTab[189075]++
												return ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:488
		// _ = "end of CoverTab[189075]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:489
		_go_fuzz_dep_.CoverTab[189076]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:489
		// _ = "end of CoverTab[189076]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:489
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:489
	// _ = "end of CoverTab[189072]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:489
	_go_fuzz_dep_.CoverTab[189073]++
											ok := ed25519.Verify(ctx.publicKey, payload, signature)
											if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:491
		_go_fuzz_dep_.CoverTab[189077]++
												return errors.New("square/go-jose: ed25519 signature failed to verify")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:492
		// _ = "end of CoverTab[189077]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:493
		_go_fuzz_dep_.CoverTab[189078]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:493
		// _ = "end of CoverTab[189078]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:493
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:493
	// _ = "end of CoverTab[189073]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:493
	_go_fuzz_dep_.CoverTab[189074]++
											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:494
	// _ = "end of CoverTab[189074]"
}

// Sign the given payload
func (ctx ecDecrypterSigner) signPayload(payload []byte, alg SignatureAlgorithm) (Signature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:498
	_go_fuzz_dep_.CoverTab[189079]++
											var expectedBitSize int
											var hash crypto.Hash

											switch alg {
	case ES256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:503
		_go_fuzz_dep_.CoverTab[189084]++
												expectedBitSize = 256
												hash = crypto.SHA256
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:505
		// _ = "end of CoverTab[189084]"
	case ES384:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:506
		_go_fuzz_dep_.CoverTab[189085]++
												expectedBitSize = 384
												hash = crypto.SHA384
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:508
		// _ = "end of CoverTab[189085]"
	case ES512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:509
		_go_fuzz_dep_.CoverTab[189086]++
												expectedBitSize = 521
												hash = crypto.SHA512
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:511
		// _ = "end of CoverTab[189086]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:511
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:511
		_go_fuzz_dep_.CoverTab[189087]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:511
		// _ = "end of CoverTab[189087]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:512
	// _ = "end of CoverTab[189079]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:512
	_go_fuzz_dep_.CoverTab[189080]++

											curveBits := ctx.privateKey.Curve.Params().BitSize
											if expectedBitSize != curveBits {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:515
		_go_fuzz_dep_.CoverTab[189088]++
												return Signature{}, fmt.Errorf("square/go-jose: expected %d bit key, got %d bits instead", expectedBitSize, curveBits)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:516
		// _ = "end of CoverTab[189088]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:517
		_go_fuzz_dep_.CoverTab[189089]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:517
		// _ = "end of CoverTab[189089]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:517
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:517
	// _ = "end of CoverTab[189080]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:517
	_go_fuzz_dep_.CoverTab[189081]++

											hasher := hash.New()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:522
	_, _ = hasher.Write(payload)
	hashed := hasher.Sum(nil)

	r, s, err := ecdsa.Sign(RandReader, ctx.privateKey, hashed)
	if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:526
		_go_fuzz_dep_.CoverTab[189090]++
												return Signature{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:527
		// _ = "end of CoverTab[189090]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:528
		_go_fuzz_dep_.CoverTab[189091]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:528
		// _ = "end of CoverTab[189091]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:528
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:528
	// _ = "end of CoverTab[189081]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:528
	_go_fuzz_dep_.CoverTab[189082]++

											keyBytes := curveBits / 8
											if curveBits%8 > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:531
		_go_fuzz_dep_.CoverTab[189092]++
												keyBytes++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:532
		// _ = "end of CoverTab[189092]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:533
		_go_fuzz_dep_.CoverTab[189093]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:533
		// _ = "end of CoverTab[189093]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:533
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:533
	// _ = "end of CoverTab[189082]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:533
	_go_fuzz_dep_.CoverTab[189083]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:538
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
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:551
	// _ = "end of CoverTab[189083]"
}

// Verify the given payload
func (ctx ecEncrypterVerifier) verifyPayload(payload []byte, signature []byte, alg SignatureAlgorithm) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:555
	_go_fuzz_dep_.CoverTab[189094]++
											var keySize int
											var hash crypto.Hash

											switch alg {
	case ES256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:560
		_go_fuzz_dep_.CoverTab[189098]++
												keySize = 32
												hash = crypto.SHA256
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:562
		// _ = "end of CoverTab[189098]"
	case ES384:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:563
		_go_fuzz_dep_.CoverTab[189099]++
												keySize = 48
												hash = crypto.SHA384
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:565
		// _ = "end of CoverTab[189099]"
	case ES512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:566
		_go_fuzz_dep_.CoverTab[189100]++
												keySize = 66
												hash = crypto.SHA512
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:568
		// _ = "end of CoverTab[189100]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:569
		_go_fuzz_dep_.CoverTab[189101]++
												return ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:570
		// _ = "end of CoverTab[189101]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:571
	// _ = "end of CoverTab[189094]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:571
	_go_fuzz_dep_.CoverTab[189095]++

											if len(signature) != 2*keySize {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:573
		_go_fuzz_dep_.CoverTab[189102]++
												return fmt.Errorf("square/go-jose: invalid signature size, have %d bytes, wanted %d", len(signature), 2*keySize)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:574
		// _ = "end of CoverTab[189102]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:575
		_go_fuzz_dep_.CoverTab[189103]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:575
		// _ = "end of CoverTab[189103]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:575
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:575
	// _ = "end of CoverTab[189095]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:575
	_go_fuzz_dep_.CoverTab[189096]++

											hasher := hash.New()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:580
	_, _ = hasher.Write(payload)
	hashed := hasher.Sum(nil)

	r := big.NewInt(0).SetBytes(signature[:keySize])
	s := big.NewInt(0).SetBytes(signature[keySize:])

	match := ecdsa.Verify(ctx.publicKey, hashed, r, s)
	if !match {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:587
		_go_fuzz_dep_.CoverTab[189104]++
												return errors.New("square/go-jose: ecdsa signature failed to verify")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:588
		// _ = "end of CoverTab[189104]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:589
		_go_fuzz_dep_.CoverTab[189105]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:589
		// _ = "end of CoverTab[189105]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:589
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:589
	// _ = "end of CoverTab[189096]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:589
	_go_fuzz_dep_.CoverTab[189097]++

											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:591
	// _ = "end of CoverTab[189097]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:592
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/asymmetric.go:592
var _ = _go_fuzz_dep_.CoverTab
