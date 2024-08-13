//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:17
)

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"errors"
	"fmt"
	"reflect"
)

// Encrypter represents an encrypter which produces an encrypted JWE object.
type Encrypter interface {
	Encrypt(plaintext []byte) (*JsonWebEncryption, error)
	EncryptWithAuthData(plaintext []byte, aad []byte) (*JsonWebEncryption, error)
	SetCompression(alg CompressionAlgorithm)
}

// MultiEncrypter represents an encrypter which supports multiple recipients.
type MultiEncrypter interface {
	Encrypt(plaintext []byte) (*JsonWebEncryption, error)
	EncryptWithAuthData(plaintext []byte, aad []byte) (*JsonWebEncryption, error)
	SetCompression(alg CompressionAlgorithm)
	AddRecipient(alg KeyAlgorithm, encryptionKey interface{}) error
}

// A generic content cipher
type contentCipher interface {
	keySize() int
	encrypt(cek []byte, aad, plaintext []byte) (*aeadParts, error)
	decrypt(cek []byte, aad []byte, parts *aeadParts) ([]byte, error)
}

// A key generator (for generating/getting a CEK)
type keyGenerator interface {
	keySize() int
	genKey() ([]byte, rawHeader, error)
}

// A generic key encrypter
type keyEncrypter interface {
	encryptKey(cek []byte, alg KeyAlgorithm) (recipientInfo, error)	// Encrypt a key
}

// A generic key decrypter
type keyDecrypter interface {
	decryptKey(headers rawHeader, recipient *recipientInfo, generator keyGenerator) ([]byte, error)	// Decrypt a key
}

// A generic encrypter based on the given key encrypter and content cipher.
type genericEncrypter struct {
	contentAlg	ContentEncryption
	compressionAlg	CompressionAlgorithm
	cipher		contentCipher
	recipients	[]recipientKeyInfo
	keyGenerator	keyGenerator
}

type recipientKeyInfo struct {
	keyID		string
	keyAlg		KeyAlgorithm
	keyEncrypter	keyEncrypter
}

// SetCompression sets a compression algorithm to be applied before encryption.
func (ctx *genericEncrypter) SetCompression(compressionAlg CompressionAlgorithm) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:81
	_go_fuzz_dep_.CoverTab[186015]++
											ctx.compressionAlg = compressionAlg
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:82
	// _ = "end of CoverTab[186015]"
}

// NewEncrypter creates an appropriate encrypter based on the key type
func NewEncrypter(alg KeyAlgorithm, enc ContentEncryption, encryptionKey interface{}) (Encrypter, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:86
	_go_fuzz_dep_.CoverTab[186016]++
											encrypter := &genericEncrypter{
		contentAlg:	enc,
		compressionAlg:	NONE,
		recipients:	[]recipientKeyInfo{},
		cipher:		getContentCipher(enc),
	}

	if encrypter.cipher == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:94
		_go_fuzz_dep_.CoverTab[186019]++
												return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:95
		// _ = "end of CoverTab[186019]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:96
		_go_fuzz_dep_.CoverTab[186020]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:96
		// _ = "end of CoverTab[186020]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:96
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:96
	// _ = "end of CoverTab[186016]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:96
	_go_fuzz_dep_.CoverTab[186017]++

											var keyID string
											var rawKey interface{}
											switch encryptionKey := encryptionKey.(type) {
	case *JsonWebKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:101
		_go_fuzz_dep_.CoverTab[186021]++
												keyID = encryptionKey.KeyID
												rawKey = encryptionKey.Key
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:103
		// _ = "end of CoverTab[186021]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:104
		_go_fuzz_dep_.CoverTab[186022]++
												rawKey = encryptionKey
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:105
		// _ = "end of CoverTab[186022]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:106
	// _ = "end of CoverTab[186017]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:106
	_go_fuzz_dep_.CoverTab[186018]++

											switch alg {
	case DIRECT:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:109
		_go_fuzz_dep_.CoverTab[186023]++

												if reflect.TypeOf(rawKey) != reflect.TypeOf([]byte{}) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:111
			_go_fuzz_dep_.CoverTab[186030]++
													return nil, ErrUnsupportedKeyType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:112
			// _ = "end of CoverTab[186030]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:113
			_go_fuzz_dep_.CoverTab[186031]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:113
			// _ = "end of CoverTab[186031]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:113
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:113
		// _ = "end of CoverTab[186023]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:113
		_go_fuzz_dep_.CoverTab[186024]++
												encrypter.keyGenerator = staticKeyGenerator{
			key: rawKey.([]byte),
		}
		recipient, _ := newSymmetricRecipient(alg, rawKey.([]byte))
		if keyID != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:118
			_go_fuzz_dep_.CoverTab[186032]++
													recipient.keyID = keyID
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:119
			// _ = "end of CoverTab[186032]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:120
			_go_fuzz_dep_.CoverTab[186033]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:120
			// _ = "end of CoverTab[186033]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:120
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:120
		// _ = "end of CoverTab[186024]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:120
		_go_fuzz_dep_.CoverTab[186025]++
												encrypter.recipients = []recipientKeyInfo{recipient}
												return encrypter, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:122
		// _ = "end of CoverTab[186025]"
	case ECDH_ES:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:123
		_go_fuzz_dep_.CoverTab[186026]++

												typeOf := reflect.TypeOf(rawKey)
												if typeOf != reflect.TypeOf(&ecdsa.PublicKey{}) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:126
			_go_fuzz_dep_.CoverTab[186034]++
													return nil, ErrUnsupportedKeyType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:127
			// _ = "end of CoverTab[186034]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:128
			_go_fuzz_dep_.CoverTab[186035]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:128
			// _ = "end of CoverTab[186035]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:128
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:128
		// _ = "end of CoverTab[186026]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:128
		_go_fuzz_dep_.CoverTab[186027]++
												encrypter.keyGenerator = ecKeyGenerator{
			size:		encrypter.cipher.keySize(),
			algID:		string(enc),
			publicKey:	rawKey.(*ecdsa.PublicKey),
		}
		recipient, _ := newECDHRecipient(alg, rawKey.(*ecdsa.PublicKey))
		if keyID != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:135
			_go_fuzz_dep_.CoverTab[186036]++
													recipient.keyID = keyID
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:136
			// _ = "end of CoverTab[186036]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:137
			_go_fuzz_dep_.CoverTab[186037]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:137
			// _ = "end of CoverTab[186037]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:137
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:137
		// _ = "end of CoverTab[186027]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:137
		_go_fuzz_dep_.CoverTab[186028]++
												encrypter.recipients = []recipientKeyInfo{recipient}
												return encrypter, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:139
		// _ = "end of CoverTab[186028]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:140
		_go_fuzz_dep_.CoverTab[186029]++

												encrypter.keyGenerator = randomKeyGenerator{
			size: encrypter.cipher.keySize(),
		}
												err := encrypter.AddRecipient(alg, encryptionKey)
												return encrypter, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:146
		// _ = "end of CoverTab[186029]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:147
	// _ = "end of CoverTab[186018]"
}

// NewMultiEncrypter creates a multi-encrypter based on the given parameters
func NewMultiEncrypter(enc ContentEncryption) (MultiEncrypter, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:151
	_go_fuzz_dep_.CoverTab[186038]++
											cipher := getContentCipher(enc)

											if cipher == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:154
		_go_fuzz_dep_.CoverTab[186040]++
												return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:155
		// _ = "end of CoverTab[186040]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:156
		_go_fuzz_dep_.CoverTab[186041]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:156
		// _ = "end of CoverTab[186041]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:156
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:156
	// _ = "end of CoverTab[186038]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:156
	_go_fuzz_dep_.CoverTab[186039]++

											encrypter := &genericEncrypter{
		contentAlg:	enc,
		compressionAlg:	NONE,
		recipients:	[]recipientKeyInfo{},
		cipher:		cipher,
		keyGenerator: randomKeyGenerator{
			size: cipher.keySize(),
		},
	}

											return encrypter, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:168
	// _ = "end of CoverTab[186039]"
}

func (ctx *genericEncrypter) AddRecipient(alg KeyAlgorithm, encryptionKey interface{}) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:171
	_go_fuzz_dep_.CoverTab[186042]++
											var recipient recipientKeyInfo

											switch alg {
	case DIRECT, ECDH_ES:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:175
		_go_fuzz_dep_.CoverTab[186045]++
												return fmt.Errorf("square/go-jose: key algorithm '%s' not supported in multi-recipient mode", alg)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:176
		// _ = "end of CoverTab[186045]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:176
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:176
		_go_fuzz_dep_.CoverTab[186046]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:176
		// _ = "end of CoverTab[186046]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:177
	// _ = "end of CoverTab[186042]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:177
	_go_fuzz_dep_.CoverTab[186043]++

											recipient, err = makeJWERecipient(alg, encryptionKey)

											if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:181
		_go_fuzz_dep_.CoverTab[186047]++
												ctx.recipients = append(ctx.recipients, recipient)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:182
		// _ = "end of CoverTab[186047]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:183
		_go_fuzz_dep_.CoverTab[186048]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:183
		// _ = "end of CoverTab[186048]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:183
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:183
	// _ = "end of CoverTab[186043]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:183
	_go_fuzz_dep_.CoverTab[186044]++
											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:184
	// _ = "end of CoverTab[186044]"
}

func makeJWERecipient(alg KeyAlgorithm, encryptionKey interface{}) (recipientKeyInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:187
	_go_fuzz_dep_.CoverTab[186049]++
											switch encryptionKey := encryptionKey.(type) {
	case *rsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:189
		_go_fuzz_dep_.CoverTab[186050]++
												return newRSARecipient(alg, encryptionKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:190
		// _ = "end of CoverTab[186050]"
	case *ecdsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:191
		_go_fuzz_dep_.CoverTab[186051]++
												return newECDHRecipient(alg, encryptionKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:192
		// _ = "end of CoverTab[186051]"
	case []byte:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:193
		_go_fuzz_dep_.CoverTab[186052]++
												return newSymmetricRecipient(alg, encryptionKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:194
		// _ = "end of CoverTab[186052]"
	case *JsonWebKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:195
		_go_fuzz_dep_.CoverTab[186053]++
												recipient, err := makeJWERecipient(alg, encryptionKey.Key)
												if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:197
			_go_fuzz_dep_.CoverTab[186056]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:197
			return encryptionKey.KeyID != ""
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:197
			// _ = "end of CoverTab[186056]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:197
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:197
			_go_fuzz_dep_.CoverTab[186057]++
													recipient.keyID = encryptionKey.KeyID
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:198
			// _ = "end of CoverTab[186057]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:199
			_go_fuzz_dep_.CoverTab[186058]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:199
			// _ = "end of CoverTab[186058]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:199
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:199
		// _ = "end of CoverTab[186053]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:199
		_go_fuzz_dep_.CoverTab[186054]++
												return recipient, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:200
		// _ = "end of CoverTab[186054]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:201
		_go_fuzz_dep_.CoverTab[186055]++
												return recipientKeyInfo{}, ErrUnsupportedKeyType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:202
		// _ = "end of CoverTab[186055]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:203
	// _ = "end of CoverTab[186049]"
}

// newDecrypter creates an appropriate decrypter based on the key type
func newDecrypter(decryptionKey interface{}) (keyDecrypter, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:207
	_go_fuzz_dep_.CoverTab[186059]++
											switch decryptionKey := decryptionKey.(type) {
	case *rsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:209
		_go_fuzz_dep_.CoverTab[186060]++
												return &rsaDecrypterSigner{
			privateKey: decryptionKey,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:212
		// _ = "end of CoverTab[186060]"
	case *ecdsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:213
		_go_fuzz_dep_.CoverTab[186061]++
												return &ecDecrypterSigner{
			privateKey: decryptionKey,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:216
		// _ = "end of CoverTab[186061]"
	case []byte:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:217
		_go_fuzz_dep_.CoverTab[186062]++
												return &symmetricKeyCipher{
			key: decryptionKey,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:220
		// _ = "end of CoverTab[186062]"
	case *JsonWebKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:221
		_go_fuzz_dep_.CoverTab[186063]++
												return newDecrypter(decryptionKey.Key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:222
		// _ = "end of CoverTab[186063]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:223
		_go_fuzz_dep_.CoverTab[186064]++
												return nil, ErrUnsupportedKeyType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:224
		// _ = "end of CoverTab[186064]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:225
	// _ = "end of CoverTab[186059]"
}

// Implementation of encrypt method producing a JWE object.
func (ctx *genericEncrypter) Encrypt(plaintext []byte) (*JsonWebEncryption, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:229
	_go_fuzz_dep_.CoverTab[186065]++
											return ctx.EncryptWithAuthData(plaintext, nil)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:230
	// _ = "end of CoverTab[186065]"
}

// Implementation of encrypt method producing a JWE object.
func (ctx *genericEncrypter) EncryptWithAuthData(plaintext, aad []byte) (*JsonWebEncryption, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:234
	_go_fuzz_dep_.CoverTab[186066]++
											obj := &JsonWebEncryption{}
											obj.aad = aad

											obj.protected = &rawHeader{
		Enc: ctx.contentAlg,
	}
	obj.recipients = make([]recipientInfo, len(ctx.recipients))

	if len(ctx.recipients) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:243
		_go_fuzz_dep_.CoverTab[186073]++
												return nil, fmt.Errorf("square/go-jose: no recipients to encrypt to")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:244
		// _ = "end of CoverTab[186073]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:245
		_go_fuzz_dep_.CoverTab[186074]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:245
		// _ = "end of CoverTab[186074]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:245
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:245
	// _ = "end of CoverTab[186066]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:245
	_go_fuzz_dep_.CoverTab[186067]++

											cek, headers, err := ctx.keyGenerator.genKey()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:248
		_go_fuzz_dep_.CoverTab[186075]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:249
		// _ = "end of CoverTab[186075]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:250
		_go_fuzz_dep_.CoverTab[186076]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:250
		// _ = "end of CoverTab[186076]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:250
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:250
	// _ = "end of CoverTab[186067]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:250
	_go_fuzz_dep_.CoverTab[186068]++

											obj.protected.merge(&headers)

											for i, info := range ctx.recipients {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:254
		_go_fuzz_dep_.CoverTab[186077]++
												recipient, err := info.keyEncrypter.encryptKey(cek, info.keyAlg)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:256
			_go_fuzz_dep_.CoverTab[186080]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:257
			// _ = "end of CoverTab[186080]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:258
			_go_fuzz_dep_.CoverTab[186081]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:258
			// _ = "end of CoverTab[186081]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:258
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:258
		// _ = "end of CoverTab[186077]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:258
		_go_fuzz_dep_.CoverTab[186078]++

												recipient.header.Alg = string(info.keyAlg)
												if info.keyID != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:261
			_go_fuzz_dep_.CoverTab[186082]++
													recipient.header.Kid = info.keyID
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:262
			// _ = "end of CoverTab[186082]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:263
			_go_fuzz_dep_.CoverTab[186083]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:263
			// _ = "end of CoverTab[186083]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:263
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:263
		// _ = "end of CoverTab[186078]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:263
		_go_fuzz_dep_.CoverTab[186079]++
												obj.recipients[i] = recipient
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:264
		// _ = "end of CoverTab[186079]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:265
	// _ = "end of CoverTab[186068]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:265
	_go_fuzz_dep_.CoverTab[186069]++

											if len(ctx.recipients) == 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:267
		_go_fuzz_dep_.CoverTab[186084]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:270
		obj.protected.merge(obj.recipients[0].header)
												obj.recipients[0].header = nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:271
		// _ = "end of CoverTab[186084]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:272
		_go_fuzz_dep_.CoverTab[186085]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:272
		// _ = "end of CoverTab[186085]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:272
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:272
	// _ = "end of CoverTab[186069]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:272
	_go_fuzz_dep_.CoverTab[186070]++

											if ctx.compressionAlg != NONE {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:274
		_go_fuzz_dep_.CoverTab[186086]++
												plaintext, err = compress(ctx.compressionAlg, plaintext)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:276
			_go_fuzz_dep_.CoverTab[186088]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:277
			// _ = "end of CoverTab[186088]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:278
			_go_fuzz_dep_.CoverTab[186089]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:278
			// _ = "end of CoverTab[186089]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:278
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:278
		// _ = "end of CoverTab[186086]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:278
		_go_fuzz_dep_.CoverTab[186087]++

												obj.protected.Zip = ctx.compressionAlg
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:280
		// _ = "end of CoverTab[186087]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:281
		_go_fuzz_dep_.CoverTab[186090]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:281
		// _ = "end of CoverTab[186090]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:281
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:281
	// _ = "end of CoverTab[186070]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:281
	_go_fuzz_dep_.CoverTab[186071]++

											authData := obj.computeAuthData()
											parts, err := ctx.cipher.encrypt(cek, authData, plaintext)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:285
		_go_fuzz_dep_.CoverTab[186091]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:286
		// _ = "end of CoverTab[186091]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:287
		_go_fuzz_dep_.CoverTab[186092]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:287
		// _ = "end of CoverTab[186092]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:287
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:287
	// _ = "end of CoverTab[186071]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:287
	_go_fuzz_dep_.CoverTab[186072]++

											obj.iv = parts.iv
											obj.ciphertext = parts.ciphertext
											obj.tag = parts.tag

											return obj, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:293
	// _ = "end of CoverTab[186072]"
}

// Decrypt and validate the object and return the plaintext. Note that this
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:296
// function does not support multi-recipient, if you desire multi-recipient
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:296
// decryption use DecryptMulti instead.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:299
func (obj JsonWebEncryption) Decrypt(decryptionKey interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:299
	_go_fuzz_dep_.CoverTab[186093]++
											headers := obj.mergedHeaders(nil)

											if len(obj.recipients) > 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:302
		_go_fuzz_dep_.CoverTab[186101]++
												return nil, errors.New("square/go-jose: too many recipients in payload; expecting only one")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:303
		// _ = "end of CoverTab[186101]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:304
		_go_fuzz_dep_.CoverTab[186102]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:304
		// _ = "end of CoverTab[186102]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:304
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:304
	// _ = "end of CoverTab[186093]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:304
	_go_fuzz_dep_.CoverTab[186094]++

											if len(headers.Crit) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:306
		_go_fuzz_dep_.CoverTab[186103]++
												return nil, fmt.Errorf("square/go-jose: unsupported crit header")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:307
		// _ = "end of CoverTab[186103]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:308
		_go_fuzz_dep_.CoverTab[186104]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:308
		// _ = "end of CoverTab[186104]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:308
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:308
	// _ = "end of CoverTab[186094]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:308
	_go_fuzz_dep_.CoverTab[186095]++

											decrypter, err := newDecrypter(decryptionKey)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:311
		_go_fuzz_dep_.CoverTab[186105]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:312
		// _ = "end of CoverTab[186105]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:313
		_go_fuzz_dep_.CoverTab[186106]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:313
		// _ = "end of CoverTab[186106]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:313
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:313
	// _ = "end of CoverTab[186095]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:313
	_go_fuzz_dep_.CoverTab[186096]++

											cipher := getContentCipher(headers.Enc)
											if cipher == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:316
		_go_fuzz_dep_.CoverTab[186107]++
												return nil, fmt.Errorf("square/go-jose: unsupported enc value '%s'", string(headers.Enc))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:317
		// _ = "end of CoverTab[186107]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:318
		_go_fuzz_dep_.CoverTab[186108]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:318
		// _ = "end of CoverTab[186108]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:318
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:318
	// _ = "end of CoverTab[186096]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:318
	_go_fuzz_dep_.CoverTab[186097]++

											generator := randomKeyGenerator{
		size: cipher.keySize(),
	}

	parts := &aeadParts{
		iv:		obj.iv,
		ciphertext:	obj.ciphertext,
		tag:		obj.tag,
	}

	authData := obj.computeAuthData()

	var plaintext []byte
	recipient := obj.recipients[0]
	recipientHeaders := obj.mergedHeaders(&recipient)

	cek, err := decrypter.decryptKey(recipientHeaders, &recipient, generator)
	if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:337
		_go_fuzz_dep_.CoverTab[186109]++

												plaintext, err = cipher.decrypt(cek, authData, parts)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:339
		// _ = "end of CoverTab[186109]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:340
		_go_fuzz_dep_.CoverTab[186110]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:340
		// _ = "end of CoverTab[186110]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:340
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:340
	// _ = "end of CoverTab[186097]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:340
	_go_fuzz_dep_.CoverTab[186098]++

											if plaintext == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:342
		_go_fuzz_dep_.CoverTab[186111]++
												return nil, ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:343
		// _ = "end of CoverTab[186111]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:344
		_go_fuzz_dep_.CoverTab[186112]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:344
		// _ = "end of CoverTab[186112]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:344
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:344
	// _ = "end of CoverTab[186098]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:344
	_go_fuzz_dep_.CoverTab[186099]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:347
	if obj.protected.Zip != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:347
		_go_fuzz_dep_.CoverTab[186113]++
												plaintext, err = decompress(obj.protected.Zip, plaintext)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:348
		// _ = "end of CoverTab[186113]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:349
		_go_fuzz_dep_.CoverTab[186114]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:349
		// _ = "end of CoverTab[186114]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:349
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:349
	// _ = "end of CoverTab[186099]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:349
	_go_fuzz_dep_.CoverTab[186100]++

											return plaintext, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:351
	// _ = "end of CoverTab[186100]"
}

// DecryptMulti decrypts and validates the object and returns the plaintexts,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:354
// with support for multiple recipients. It returns the index of the recipient
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:354
// for which the decryption was successful, the merged headers for that recipient,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:354
// and the plaintext.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:358
func (obj JsonWebEncryption) DecryptMulti(decryptionKey interface{}) (int, JoseHeader, []byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:358
	_go_fuzz_dep_.CoverTab[186115]++
											globalHeaders := obj.mergedHeaders(nil)

											if len(globalHeaders.Crit) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:361
		_go_fuzz_dep_.CoverTab[186122]++
												return -1, JoseHeader{}, nil, fmt.Errorf("square/go-jose: unsupported crit header")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:362
		// _ = "end of CoverTab[186122]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:363
		_go_fuzz_dep_.CoverTab[186123]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:363
		// _ = "end of CoverTab[186123]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:363
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:363
	// _ = "end of CoverTab[186115]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:363
	_go_fuzz_dep_.CoverTab[186116]++

											decrypter, err := newDecrypter(decryptionKey)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:366
		_go_fuzz_dep_.CoverTab[186124]++
												return -1, JoseHeader{}, nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:367
		// _ = "end of CoverTab[186124]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:368
		_go_fuzz_dep_.CoverTab[186125]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:368
		// _ = "end of CoverTab[186125]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:368
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:368
	// _ = "end of CoverTab[186116]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:368
	_go_fuzz_dep_.CoverTab[186117]++

											cipher := getContentCipher(globalHeaders.Enc)
											if cipher == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:371
		_go_fuzz_dep_.CoverTab[186126]++
												return -1, JoseHeader{}, nil, fmt.Errorf("square/go-jose: unsupported enc value '%s'", string(globalHeaders.Enc))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:372
		// _ = "end of CoverTab[186126]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:373
		_go_fuzz_dep_.CoverTab[186127]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:373
		// _ = "end of CoverTab[186127]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:373
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:373
	// _ = "end of CoverTab[186117]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:373
	_go_fuzz_dep_.CoverTab[186118]++

											generator := randomKeyGenerator{
		size: cipher.keySize(),
	}

	parts := &aeadParts{
		iv:		obj.iv,
		ciphertext:	obj.ciphertext,
		tag:		obj.tag,
	}

	authData := obj.computeAuthData()

	index := -1
	var plaintext []byte
	var headers rawHeader

	for i, recipient := range obj.recipients {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:391
		_go_fuzz_dep_.CoverTab[186128]++
												recipientHeaders := obj.mergedHeaders(&recipient)

												cek, err := decrypter.decryptKey(recipientHeaders, &recipient, generator)
												if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:395
			_go_fuzz_dep_.CoverTab[186129]++

													plaintext, err = cipher.decrypt(cek, authData, parts)
													if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:398
				_go_fuzz_dep_.CoverTab[186130]++
														index = i
														headers = recipientHeaders
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:401
				// _ = "end of CoverTab[186130]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:402
				_go_fuzz_dep_.CoverTab[186131]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:402
				// _ = "end of CoverTab[186131]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:402
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:402
			// _ = "end of CoverTab[186129]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:403
			_go_fuzz_dep_.CoverTab[186132]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:403
			// _ = "end of CoverTab[186132]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:403
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:403
		// _ = "end of CoverTab[186128]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:404
	// _ = "end of CoverTab[186118]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:404
	_go_fuzz_dep_.CoverTab[186119]++

											if plaintext == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:406
		_go_fuzz_dep_.CoverTab[186133]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:406
		return err != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:406
		// _ = "end of CoverTab[186133]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:406
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:406
		_go_fuzz_dep_.CoverTab[186134]++
												return -1, JoseHeader{}, nil, ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:407
		// _ = "end of CoverTab[186134]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:408
		_go_fuzz_dep_.CoverTab[186135]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:408
		// _ = "end of CoverTab[186135]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:408
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:408
	// _ = "end of CoverTab[186119]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:408
	_go_fuzz_dep_.CoverTab[186120]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:411
	if obj.protected.Zip != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:411
		_go_fuzz_dep_.CoverTab[186136]++
												plaintext, err = decompress(obj.protected.Zip, plaintext)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:412
		// _ = "end of CoverTab[186136]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:413
		_go_fuzz_dep_.CoverTab[186137]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:413
		// _ = "end of CoverTab[186137]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:413
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:413
	// _ = "end of CoverTab[186120]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:413
	_go_fuzz_dep_.CoverTab[186121]++

											return index, headers.sanitized(), plaintext, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:415
	// _ = "end of CoverTab[186121]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:416
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/crypter.go:416
var _ = _go_fuzz_dep_.CoverTab
