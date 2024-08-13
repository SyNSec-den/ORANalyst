//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:17
)

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"errors"
	"fmt"
	"reflect"

	"gopkg.in/square/go-jose.v2/json"
)

// Encrypter represents an encrypter which produces an encrypted JWE object.
type Encrypter interface {
	Encrypt(plaintext []byte) (*JSONWebEncryption, error)
	EncryptWithAuthData(plaintext []byte, aad []byte) (*JSONWebEncryption, error)
	Options() EncrypterOptions
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
	extraHeaders	map[HeaderKey]interface{}
}

type recipientKeyInfo struct {
	keyID		string
	keyAlg		KeyAlgorithm
	keyEncrypter	keyEncrypter
}

// EncrypterOptions represents options that can be set on new encrypters.
type EncrypterOptions struct {
	Compression	CompressionAlgorithm

	// Optional map of additional keys to be inserted into the protected header
	// of a JWS object. Some specifications which make use of JWS like to insert
	// additional values here. All values must be JSON-serializable.
	ExtraHeaders	map[HeaderKey]interface{}
}

// WithHeader adds an arbitrary value to the ExtraHeaders map, initializing it
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:85
// if necessary. It returns itself and so can be used in a fluent style.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:87
func (eo *EncrypterOptions) WithHeader(k HeaderKey, v interface{}) *EncrypterOptions {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:87
	_go_fuzz_dep_.CoverTab[189106]++
											if eo.ExtraHeaders == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:88
		_go_fuzz_dep_.CoverTab[189108]++
												eo.ExtraHeaders = map[HeaderKey]interface{}{}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:89
		// _ = "end of CoverTab[189108]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:90
		_go_fuzz_dep_.CoverTab[189109]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:90
		// _ = "end of CoverTab[189109]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:90
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:90
	// _ = "end of CoverTab[189106]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:90
	_go_fuzz_dep_.CoverTab[189107]++
											eo.ExtraHeaders[k] = v
											return eo
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:92
	// _ = "end of CoverTab[189107]"
}

// WithContentType adds a content type ("cty") header and returns the updated
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:95
// EncrypterOptions.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:97
func (eo *EncrypterOptions) WithContentType(contentType ContentType) *EncrypterOptions {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:97
	_go_fuzz_dep_.CoverTab[189110]++
											return eo.WithHeader(HeaderContentType, contentType)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:98
	// _ = "end of CoverTab[189110]"
}

// WithType adds a type ("typ") header and returns the updated EncrypterOptions.
func (eo *EncrypterOptions) WithType(typ ContentType) *EncrypterOptions {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:102
	_go_fuzz_dep_.CoverTab[189111]++
											return eo.WithHeader(HeaderType, typ)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:103
	// _ = "end of CoverTab[189111]"
}

// Recipient represents an algorithm/key to encrypt messages to.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:106
//
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:106
// PBES2Count and PBES2Salt correspond with the  "p2c" and "p2s" headers used
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:106
// on the password-based encryption algorithms PBES2-HS256+A128KW,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:106
// PBES2-HS384+A192KW, and PBES2-HS512+A256KW. If they are not provided a safe
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:106
// default of 100000 will be used for the count and a 128-bit random salt will
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:106
// be generated.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:113
type Recipient struct {
	Algorithm	KeyAlgorithm
	Key		interface{}
	KeyID		string
	PBES2Count	int
	PBES2Salt	[]byte
}

// NewEncrypter creates an appropriate encrypter based on the key type
func NewEncrypter(enc ContentEncryption, rcpt Recipient, opts *EncrypterOptions) (Encrypter, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:122
	_go_fuzz_dep_.CoverTab[189112]++
											encrypter := &genericEncrypter{
		contentAlg:	enc,
		recipients:	[]recipientKeyInfo{},
		cipher:		getContentCipher(enc),
	}
	if opts != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:128
		_go_fuzz_dep_.CoverTab[189116]++
												encrypter.compressionAlg = opts.Compression
												encrypter.extraHeaders = opts.ExtraHeaders
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:130
		// _ = "end of CoverTab[189116]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:131
		_go_fuzz_dep_.CoverTab[189117]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:131
		// _ = "end of CoverTab[189117]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:131
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:131
	// _ = "end of CoverTab[189112]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:131
	_go_fuzz_dep_.CoverTab[189113]++

											if encrypter.cipher == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:133
		_go_fuzz_dep_.CoverTab[189118]++
												return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:134
		// _ = "end of CoverTab[189118]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:135
		_go_fuzz_dep_.CoverTab[189119]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:135
		// _ = "end of CoverTab[189119]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:135
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:135
	// _ = "end of CoverTab[189113]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:135
	_go_fuzz_dep_.CoverTab[189114]++

											var keyID string
											var rawKey interface{}
											switch encryptionKey := rcpt.Key.(type) {
	case JSONWebKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:140
		_go_fuzz_dep_.CoverTab[189120]++
												keyID, rawKey = encryptionKey.KeyID, encryptionKey.Key
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:141
		// _ = "end of CoverTab[189120]"
	case *JSONWebKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:142
		_go_fuzz_dep_.CoverTab[189121]++
												keyID, rawKey = encryptionKey.KeyID, encryptionKey.Key
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:143
		// _ = "end of CoverTab[189121]"
	case OpaqueKeyEncrypter:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:144
		_go_fuzz_dep_.CoverTab[189122]++
												keyID, rawKey = encryptionKey.KeyID(), encryptionKey
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:145
		// _ = "end of CoverTab[189122]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:146
		_go_fuzz_dep_.CoverTab[189123]++
												rawKey = encryptionKey
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:147
		// _ = "end of CoverTab[189123]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:148
	// _ = "end of CoverTab[189114]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:148
	_go_fuzz_dep_.CoverTab[189115]++

											switch rcpt.Algorithm {
	case DIRECT:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:151
		_go_fuzz_dep_.CoverTab[189124]++

												if reflect.TypeOf(rawKey) != reflect.TypeOf([]byte{}) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:153
			_go_fuzz_dep_.CoverTab[189132]++
													return nil, ErrUnsupportedKeyType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:154
			// _ = "end of CoverTab[189132]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:155
			_go_fuzz_dep_.CoverTab[189133]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:155
			// _ = "end of CoverTab[189133]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:155
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:155
		// _ = "end of CoverTab[189124]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:155
		_go_fuzz_dep_.CoverTab[189125]++
												if encrypter.cipher.keySize() != len(rawKey.([]byte)) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:156
			_go_fuzz_dep_.CoverTab[189134]++
													return nil, ErrInvalidKeySize
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:157
			// _ = "end of CoverTab[189134]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:158
			_go_fuzz_dep_.CoverTab[189135]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:158
			// _ = "end of CoverTab[189135]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:158
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:158
		// _ = "end of CoverTab[189125]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:158
		_go_fuzz_dep_.CoverTab[189126]++
												encrypter.keyGenerator = staticKeyGenerator{
			key: rawKey.([]byte),
		}
		recipientInfo, _ := newSymmetricRecipient(rcpt.Algorithm, rawKey.([]byte))
		recipientInfo.keyID = keyID
		if rcpt.KeyID != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:164
			_go_fuzz_dep_.CoverTab[189136]++
													recipientInfo.keyID = rcpt.KeyID
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:165
			// _ = "end of CoverTab[189136]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:166
			_go_fuzz_dep_.CoverTab[189137]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:166
			// _ = "end of CoverTab[189137]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:166
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:166
		// _ = "end of CoverTab[189126]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:166
		_go_fuzz_dep_.CoverTab[189127]++
												encrypter.recipients = []recipientKeyInfo{recipientInfo}
												return encrypter, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:168
		// _ = "end of CoverTab[189127]"
	case ECDH_ES:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:169
		_go_fuzz_dep_.CoverTab[189128]++

												typeOf := reflect.TypeOf(rawKey)
												if typeOf != reflect.TypeOf(&ecdsa.PublicKey{}) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:172
			_go_fuzz_dep_.CoverTab[189138]++
													return nil, ErrUnsupportedKeyType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:173
			// _ = "end of CoverTab[189138]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:174
			_go_fuzz_dep_.CoverTab[189139]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:174
			// _ = "end of CoverTab[189139]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:174
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:174
		// _ = "end of CoverTab[189128]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:174
		_go_fuzz_dep_.CoverTab[189129]++
												encrypter.keyGenerator = ecKeyGenerator{
			size:		encrypter.cipher.keySize(),
			algID:		string(enc),
			publicKey:	rawKey.(*ecdsa.PublicKey),
		}
		recipientInfo, _ := newECDHRecipient(rcpt.Algorithm, rawKey.(*ecdsa.PublicKey))
		recipientInfo.keyID = keyID
		if rcpt.KeyID != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:182
			_go_fuzz_dep_.CoverTab[189140]++
													recipientInfo.keyID = rcpt.KeyID
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:183
			// _ = "end of CoverTab[189140]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:184
			_go_fuzz_dep_.CoverTab[189141]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:184
			// _ = "end of CoverTab[189141]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:184
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:184
		// _ = "end of CoverTab[189129]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:184
		_go_fuzz_dep_.CoverTab[189130]++
												encrypter.recipients = []recipientKeyInfo{recipientInfo}
												return encrypter, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:186
		// _ = "end of CoverTab[189130]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:187
		_go_fuzz_dep_.CoverTab[189131]++

												encrypter.keyGenerator = randomKeyGenerator{
			size: encrypter.cipher.keySize(),
		}
												err := encrypter.addRecipient(rcpt)
												return encrypter, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:193
		// _ = "end of CoverTab[189131]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:194
	// _ = "end of CoverTab[189115]"
}

// NewMultiEncrypter creates a multi-encrypter based on the given parameters
func NewMultiEncrypter(enc ContentEncryption, rcpts []Recipient, opts *EncrypterOptions) (Encrypter, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:198
	_go_fuzz_dep_.CoverTab[189142]++
											cipher := getContentCipher(enc)

											if cipher == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:201
		_go_fuzz_dep_.CoverTab[189147]++
												return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:202
		// _ = "end of CoverTab[189147]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:203
		_go_fuzz_dep_.CoverTab[189148]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:203
		// _ = "end of CoverTab[189148]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:203
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:203
	// _ = "end of CoverTab[189142]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:203
	_go_fuzz_dep_.CoverTab[189143]++
											if rcpts == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:204
		_go_fuzz_dep_.CoverTab[189149]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:204
		return len(rcpts) == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:204
		// _ = "end of CoverTab[189149]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:204
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:204
		_go_fuzz_dep_.CoverTab[189150]++
												return nil, fmt.Errorf("square/go-jose: recipients is nil or empty")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:205
		// _ = "end of CoverTab[189150]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:206
		_go_fuzz_dep_.CoverTab[189151]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:206
		// _ = "end of CoverTab[189151]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:206
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:206
	// _ = "end of CoverTab[189143]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:206
	_go_fuzz_dep_.CoverTab[189144]++

											encrypter := &genericEncrypter{
		contentAlg:	enc,
		recipients:	[]recipientKeyInfo{},
		cipher:		cipher,
		keyGenerator: randomKeyGenerator{
			size: cipher.keySize(),
		},
	}

	if opts != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:217
		_go_fuzz_dep_.CoverTab[189152]++
												encrypter.compressionAlg = opts.Compression
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:218
		// _ = "end of CoverTab[189152]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:219
		_go_fuzz_dep_.CoverTab[189153]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:219
		// _ = "end of CoverTab[189153]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:219
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:219
	// _ = "end of CoverTab[189144]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:219
	_go_fuzz_dep_.CoverTab[189145]++

											for _, recipient := range rcpts {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:221
		_go_fuzz_dep_.CoverTab[189154]++
												err := encrypter.addRecipient(recipient)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:223
			_go_fuzz_dep_.CoverTab[189155]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:224
			// _ = "end of CoverTab[189155]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:225
			_go_fuzz_dep_.CoverTab[189156]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:225
			// _ = "end of CoverTab[189156]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:225
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:225
		// _ = "end of CoverTab[189154]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:226
	// _ = "end of CoverTab[189145]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:226
	_go_fuzz_dep_.CoverTab[189146]++

											return encrypter, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:228
	// _ = "end of CoverTab[189146]"
}

func (ctx *genericEncrypter) addRecipient(recipient Recipient) (err error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:231
	_go_fuzz_dep_.CoverTab[189157]++
											var recipientInfo recipientKeyInfo

											switch recipient.Algorithm {
	case DIRECT, ECDH_ES:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:235
		_go_fuzz_dep_.CoverTab[189162]++
												return fmt.Errorf("square/go-jose: key algorithm '%s' not supported in multi-recipient mode", recipient.Algorithm)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:236
		// _ = "end of CoverTab[189162]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:236
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:236
		_go_fuzz_dep_.CoverTab[189163]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:236
		// _ = "end of CoverTab[189163]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:237
	// _ = "end of CoverTab[189157]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:237
	_go_fuzz_dep_.CoverTab[189158]++

											recipientInfo, err = makeJWERecipient(recipient.Algorithm, recipient.Key)
											if recipient.KeyID != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:240
		_go_fuzz_dep_.CoverTab[189164]++
												recipientInfo.keyID = recipient.KeyID
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:241
		// _ = "end of CoverTab[189164]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:242
		_go_fuzz_dep_.CoverTab[189165]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:242
		// _ = "end of CoverTab[189165]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:242
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:242
	// _ = "end of CoverTab[189158]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:242
	_go_fuzz_dep_.CoverTab[189159]++

											switch recipient.Algorithm {
	case PBES2_HS256_A128KW, PBES2_HS384_A192KW, PBES2_HS512_A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:245
		_go_fuzz_dep_.CoverTab[189166]++
												if sr, ok := recipientInfo.keyEncrypter.(*symmetricKeyCipher); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:246
			_go_fuzz_dep_.CoverTab[189168]++
													sr.p2c = recipient.PBES2Count
													sr.p2s = recipient.PBES2Salt
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:248
			// _ = "end of CoverTab[189168]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:249
			_go_fuzz_dep_.CoverTab[189169]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:249
			// _ = "end of CoverTab[189169]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:249
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:249
		// _ = "end of CoverTab[189166]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:249
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:249
		_go_fuzz_dep_.CoverTab[189167]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:249
		// _ = "end of CoverTab[189167]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:250
	// _ = "end of CoverTab[189159]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:250
	_go_fuzz_dep_.CoverTab[189160]++

											if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:252
		_go_fuzz_dep_.CoverTab[189170]++
												ctx.recipients = append(ctx.recipients, recipientInfo)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:253
		// _ = "end of CoverTab[189170]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:254
		_go_fuzz_dep_.CoverTab[189171]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:254
		// _ = "end of CoverTab[189171]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:254
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:254
	// _ = "end of CoverTab[189160]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:254
	_go_fuzz_dep_.CoverTab[189161]++
											return err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:255
	// _ = "end of CoverTab[189161]"
}

func makeJWERecipient(alg KeyAlgorithm, encryptionKey interface{}) (recipientKeyInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:258
	_go_fuzz_dep_.CoverTab[189172]++
											switch encryptionKey := encryptionKey.(type) {
	case *rsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:260
		_go_fuzz_dep_.CoverTab[189175]++
												return newRSARecipient(alg, encryptionKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:261
		// _ = "end of CoverTab[189175]"
	case *ecdsa.PublicKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:262
		_go_fuzz_dep_.CoverTab[189176]++
												return newECDHRecipient(alg, encryptionKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:263
		// _ = "end of CoverTab[189176]"
	case []byte:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:264
		_go_fuzz_dep_.CoverTab[189177]++
												return newSymmetricRecipient(alg, encryptionKey)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:265
		// _ = "end of CoverTab[189177]"
	case string:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:266
		_go_fuzz_dep_.CoverTab[189178]++
												return newSymmetricRecipient(alg, []byte(encryptionKey))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:267
		// _ = "end of CoverTab[189178]"
	case *JSONWebKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:268
		_go_fuzz_dep_.CoverTab[189179]++
												recipient, err := makeJWERecipient(alg, encryptionKey.Key)
												recipient.keyID = encryptionKey.KeyID
												return recipient, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:271
		// _ = "end of CoverTab[189179]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:272
	// _ = "end of CoverTab[189172]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:272
	_go_fuzz_dep_.CoverTab[189173]++
											if encrypter, ok := encryptionKey.(OpaqueKeyEncrypter); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:273
		_go_fuzz_dep_.CoverTab[189180]++
												return newOpaqueKeyEncrypter(alg, encrypter)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:274
		// _ = "end of CoverTab[189180]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:275
		_go_fuzz_dep_.CoverTab[189181]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:275
		// _ = "end of CoverTab[189181]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:275
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:275
	// _ = "end of CoverTab[189173]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:275
	_go_fuzz_dep_.CoverTab[189174]++
											return recipientKeyInfo{}, ErrUnsupportedKeyType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:276
	// _ = "end of CoverTab[189174]"
}

// newDecrypter creates an appropriate decrypter based on the key type
func newDecrypter(decryptionKey interface{}) (keyDecrypter, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:280
	_go_fuzz_dep_.CoverTab[189182]++
											switch decryptionKey := decryptionKey.(type) {
	case *rsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:282
		_go_fuzz_dep_.CoverTab[189185]++
												return &rsaDecrypterSigner{
			privateKey: decryptionKey,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:285
		// _ = "end of CoverTab[189185]"
	case *ecdsa.PrivateKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:286
		_go_fuzz_dep_.CoverTab[189186]++
												return &ecDecrypterSigner{
			privateKey: decryptionKey,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:289
		// _ = "end of CoverTab[189186]"
	case []byte:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:290
		_go_fuzz_dep_.CoverTab[189187]++
												return &symmetricKeyCipher{
			key: decryptionKey,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:293
		// _ = "end of CoverTab[189187]"
	case string:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:294
		_go_fuzz_dep_.CoverTab[189188]++
												return &symmetricKeyCipher{
			key: []byte(decryptionKey),
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:297
		// _ = "end of CoverTab[189188]"
	case JSONWebKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:298
		_go_fuzz_dep_.CoverTab[189189]++
												return newDecrypter(decryptionKey.Key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:299
		// _ = "end of CoverTab[189189]"
	case *JSONWebKey:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:300
		_go_fuzz_dep_.CoverTab[189190]++
												return newDecrypter(decryptionKey.Key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:301
		// _ = "end of CoverTab[189190]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:302
	// _ = "end of CoverTab[189182]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:302
	_go_fuzz_dep_.CoverTab[189183]++
											if okd, ok := decryptionKey.(OpaqueKeyDecrypter); ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:303
		_go_fuzz_dep_.CoverTab[189191]++
												return &opaqueKeyDecrypter{decrypter: okd}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:304
		// _ = "end of CoverTab[189191]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:305
		_go_fuzz_dep_.CoverTab[189192]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:305
		// _ = "end of CoverTab[189192]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:305
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:305
	// _ = "end of CoverTab[189183]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:305
	_go_fuzz_dep_.CoverTab[189184]++
											return nil, ErrUnsupportedKeyType
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:306
	// _ = "end of CoverTab[189184]"
}

// Implementation of encrypt method producing a JWE object.
func (ctx *genericEncrypter) Encrypt(plaintext []byte) (*JSONWebEncryption, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:310
	_go_fuzz_dep_.CoverTab[189193]++
											return ctx.EncryptWithAuthData(plaintext, nil)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:311
	// _ = "end of CoverTab[189193]"
}

// Implementation of encrypt method producing a JWE object.
func (ctx *genericEncrypter) EncryptWithAuthData(plaintext, aad []byte) (*JSONWebEncryption, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:315
	_go_fuzz_dep_.CoverTab[189194]++
											obj := &JSONWebEncryption{}
											obj.aad = aad

											obj.protected = &rawHeader{}
											err := obj.protected.set(headerEncryption, ctx.contentAlg)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:321
		_go_fuzz_dep_.CoverTab[189203]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:322
		// _ = "end of CoverTab[189203]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:323
		_go_fuzz_dep_.CoverTab[189204]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:323
		// _ = "end of CoverTab[189204]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:323
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:323
	// _ = "end of CoverTab[189194]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:323
	_go_fuzz_dep_.CoverTab[189195]++

											obj.recipients = make([]recipientInfo, len(ctx.recipients))

											if len(ctx.recipients) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:327
		_go_fuzz_dep_.CoverTab[189205]++
												return nil, fmt.Errorf("square/go-jose: no recipients to encrypt to")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:328
		// _ = "end of CoverTab[189205]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:329
		_go_fuzz_dep_.CoverTab[189206]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:329
		// _ = "end of CoverTab[189206]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:329
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:329
	// _ = "end of CoverTab[189195]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:329
	_go_fuzz_dep_.CoverTab[189196]++

											cek, headers, err := ctx.keyGenerator.genKey()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:332
		_go_fuzz_dep_.CoverTab[189207]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:333
		// _ = "end of CoverTab[189207]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:334
		_go_fuzz_dep_.CoverTab[189208]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:334
		// _ = "end of CoverTab[189208]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:334
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:334
	// _ = "end of CoverTab[189196]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:334
	_go_fuzz_dep_.CoverTab[189197]++

											obj.protected.merge(&headers)

											for i, info := range ctx.recipients {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:338
		_go_fuzz_dep_.CoverTab[189209]++
												recipient, err := info.keyEncrypter.encryptKey(cek, info.keyAlg)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:340
			_go_fuzz_dep_.CoverTab[189213]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:341
			// _ = "end of CoverTab[189213]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:342
			_go_fuzz_dep_.CoverTab[189214]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:342
			// _ = "end of CoverTab[189214]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:342
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:342
		// _ = "end of CoverTab[189209]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:342
		_go_fuzz_dep_.CoverTab[189210]++

												err = recipient.header.set(headerAlgorithm, info.keyAlg)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:345
			_go_fuzz_dep_.CoverTab[189215]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:346
			// _ = "end of CoverTab[189215]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:347
			_go_fuzz_dep_.CoverTab[189216]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:347
			// _ = "end of CoverTab[189216]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:347
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:347
		// _ = "end of CoverTab[189210]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:347
		_go_fuzz_dep_.CoverTab[189211]++

												if info.keyID != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:349
			_go_fuzz_dep_.CoverTab[189217]++
													err = recipient.header.set(headerKeyID, info.keyID)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:351
				_go_fuzz_dep_.CoverTab[189218]++
														return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:352
				// _ = "end of CoverTab[189218]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:353
				_go_fuzz_dep_.CoverTab[189219]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:353
				// _ = "end of CoverTab[189219]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:353
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:353
			// _ = "end of CoverTab[189217]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:354
			_go_fuzz_dep_.CoverTab[189220]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:354
			// _ = "end of CoverTab[189220]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:354
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:354
		// _ = "end of CoverTab[189211]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:354
		_go_fuzz_dep_.CoverTab[189212]++
												obj.recipients[i] = recipient
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:355
		// _ = "end of CoverTab[189212]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:356
	// _ = "end of CoverTab[189197]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:356
	_go_fuzz_dep_.CoverTab[189198]++

											if len(ctx.recipients) == 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:358
		_go_fuzz_dep_.CoverTab[189221]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:361
		obj.protected.merge(obj.recipients[0].header)
												obj.recipients[0].header = nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:362
		// _ = "end of CoverTab[189221]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:363
		_go_fuzz_dep_.CoverTab[189222]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:363
		// _ = "end of CoverTab[189222]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:363
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:363
	// _ = "end of CoverTab[189198]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:363
	_go_fuzz_dep_.CoverTab[189199]++

											if ctx.compressionAlg != NONE {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:365
		_go_fuzz_dep_.CoverTab[189223]++
												plaintext, err = compress(ctx.compressionAlg, plaintext)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:367
			_go_fuzz_dep_.CoverTab[189225]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:368
			// _ = "end of CoverTab[189225]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:369
			_go_fuzz_dep_.CoverTab[189226]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:369
			// _ = "end of CoverTab[189226]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:369
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:369
		// _ = "end of CoverTab[189223]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:369
		_go_fuzz_dep_.CoverTab[189224]++

												err = obj.protected.set(headerCompression, ctx.compressionAlg)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:372
			_go_fuzz_dep_.CoverTab[189227]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:373
			// _ = "end of CoverTab[189227]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:374
			_go_fuzz_dep_.CoverTab[189228]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:374
			// _ = "end of CoverTab[189228]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:374
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:374
		// _ = "end of CoverTab[189224]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:375
		_go_fuzz_dep_.CoverTab[189229]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:375
		// _ = "end of CoverTab[189229]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:375
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:375
	// _ = "end of CoverTab[189199]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:375
	_go_fuzz_dep_.CoverTab[189200]++

											for k, v := range ctx.extraHeaders {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:377
		_go_fuzz_dep_.CoverTab[189230]++
												b, err := json.Marshal(v)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:379
			_go_fuzz_dep_.CoverTab[189232]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:380
			// _ = "end of CoverTab[189232]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:381
			_go_fuzz_dep_.CoverTab[189233]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:381
			// _ = "end of CoverTab[189233]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:381
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:381
		// _ = "end of CoverTab[189230]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:381
		_go_fuzz_dep_.CoverTab[189231]++
												(*obj.protected)[k] = makeRawMessage(b)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:382
		// _ = "end of CoverTab[189231]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:383
	// _ = "end of CoverTab[189200]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:383
	_go_fuzz_dep_.CoverTab[189201]++

											authData := obj.computeAuthData()
											parts, err := ctx.cipher.encrypt(cek, authData, plaintext)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:387
		_go_fuzz_dep_.CoverTab[189234]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:388
		// _ = "end of CoverTab[189234]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:389
		_go_fuzz_dep_.CoverTab[189235]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:389
		// _ = "end of CoverTab[189235]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:389
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:389
	// _ = "end of CoverTab[189201]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:389
	_go_fuzz_dep_.CoverTab[189202]++

											obj.iv = parts.iv
											obj.ciphertext = parts.ciphertext
											obj.tag = parts.tag

											return obj, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:395
	// _ = "end of CoverTab[189202]"
}

func (ctx *genericEncrypter) Options() EncrypterOptions {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:398
	_go_fuzz_dep_.CoverTab[189236]++
											return EncrypterOptions{
		Compression:	ctx.compressionAlg,
		ExtraHeaders:	ctx.extraHeaders,
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:402
	// _ = "end of CoverTab[189236]"
}

// Decrypt and validate the object and return the plaintext. Note that this
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:405
// function does not support multi-recipient, if you desire multi-recipient
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:405
// decryption use DecryptMulti instead.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:408
func (obj JSONWebEncryption) Decrypt(decryptionKey interface{}) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:408
	_go_fuzz_dep_.CoverTab[189237]++
											headers := obj.mergedHeaders(nil)

											if len(obj.recipients) > 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:411
		_go_fuzz_dep_.CoverTab[189246]++
												return nil, errors.New("square/go-jose: too many recipients in payload; expecting only one")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:412
		// _ = "end of CoverTab[189246]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:413
		_go_fuzz_dep_.CoverTab[189247]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:413
		// _ = "end of CoverTab[189247]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:413
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:413
	// _ = "end of CoverTab[189237]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:413
	_go_fuzz_dep_.CoverTab[189238]++

											critical, err := headers.getCritical()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:416
		_go_fuzz_dep_.CoverTab[189248]++
												return nil, fmt.Errorf("square/go-jose: invalid crit header")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:417
		// _ = "end of CoverTab[189248]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:418
		_go_fuzz_dep_.CoverTab[189249]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:418
		// _ = "end of CoverTab[189249]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:418
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:418
	// _ = "end of CoverTab[189238]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:418
	_go_fuzz_dep_.CoverTab[189239]++

											if len(critical) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:420
		_go_fuzz_dep_.CoverTab[189250]++
												return nil, fmt.Errorf("square/go-jose: unsupported crit header")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:421
		// _ = "end of CoverTab[189250]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:422
		_go_fuzz_dep_.CoverTab[189251]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:422
		// _ = "end of CoverTab[189251]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:422
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:422
	// _ = "end of CoverTab[189239]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:422
	_go_fuzz_dep_.CoverTab[189240]++

											decrypter, err := newDecrypter(decryptionKey)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:425
		_go_fuzz_dep_.CoverTab[189252]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:426
		// _ = "end of CoverTab[189252]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:427
		_go_fuzz_dep_.CoverTab[189253]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:427
		// _ = "end of CoverTab[189253]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:427
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:427
	// _ = "end of CoverTab[189240]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:427
	_go_fuzz_dep_.CoverTab[189241]++

											cipher := getContentCipher(headers.getEncryption())
											if cipher == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:430
		_go_fuzz_dep_.CoverTab[189254]++
												return nil, fmt.Errorf("square/go-jose: unsupported enc value '%s'", string(headers.getEncryption()))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:431
		// _ = "end of CoverTab[189254]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:432
		_go_fuzz_dep_.CoverTab[189255]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:432
		// _ = "end of CoverTab[189255]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:432
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:432
	// _ = "end of CoverTab[189241]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:432
	_go_fuzz_dep_.CoverTab[189242]++

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
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:451
		_go_fuzz_dep_.CoverTab[189256]++

												plaintext, err = cipher.decrypt(cek, authData, parts)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:453
		// _ = "end of CoverTab[189256]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:454
		_go_fuzz_dep_.CoverTab[189257]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:454
		// _ = "end of CoverTab[189257]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:454
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:454
	// _ = "end of CoverTab[189242]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:454
	_go_fuzz_dep_.CoverTab[189243]++

											if plaintext == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:456
		_go_fuzz_dep_.CoverTab[189258]++
												return nil, ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:457
		// _ = "end of CoverTab[189258]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:458
		_go_fuzz_dep_.CoverTab[189259]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:458
		// _ = "end of CoverTab[189259]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:458
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:458
	// _ = "end of CoverTab[189243]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:458
	_go_fuzz_dep_.CoverTab[189244]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:461
	if comp := obj.protected.getCompression(); comp != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:461
		_go_fuzz_dep_.CoverTab[189260]++
												plaintext, err = decompress(comp, plaintext)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:462
		// _ = "end of CoverTab[189260]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:463
		_go_fuzz_dep_.CoverTab[189261]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:463
		// _ = "end of CoverTab[189261]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:463
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:463
	// _ = "end of CoverTab[189244]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:463
	_go_fuzz_dep_.CoverTab[189245]++

											return plaintext, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:465
	// _ = "end of CoverTab[189245]"
}

// DecryptMulti decrypts and validates the object and returns the plaintexts,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:468
// with support for multiple recipients. It returns the index of the recipient
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:468
// for which the decryption was successful, the merged headers for that recipient,
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:468
// and the plaintext.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:472
func (obj JSONWebEncryption) DecryptMulti(decryptionKey interface{}) (int, Header, []byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:472
	_go_fuzz_dep_.CoverTab[189262]++
											globalHeaders := obj.mergedHeaders(nil)

											critical, err := globalHeaders.getCritical()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:476
		_go_fuzz_dep_.CoverTab[189271]++
												return -1, Header{}, nil, fmt.Errorf("square/go-jose: invalid crit header")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:477
		// _ = "end of CoverTab[189271]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:478
		_go_fuzz_dep_.CoverTab[189272]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:478
		// _ = "end of CoverTab[189272]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:478
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:478
	// _ = "end of CoverTab[189262]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:478
	_go_fuzz_dep_.CoverTab[189263]++

											if len(critical) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:480
		_go_fuzz_dep_.CoverTab[189273]++
												return -1, Header{}, nil, fmt.Errorf("square/go-jose: unsupported crit header")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:481
		// _ = "end of CoverTab[189273]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:482
		_go_fuzz_dep_.CoverTab[189274]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:482
		// _ = "end of CoverTab[189274]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:482
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:482
	// _ = "end of CoverTab[189263]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:482
	_go_fuzz_dep_.CoverTab[189264]++

											decrypter, err := newDecrypter(decryptionKey)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:485
		_go_fuzz_dep_.CoverTab[189275]++
												return -1, Header{}, nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:486
		// _ = "end of CoverTab[189275]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:487
		_go_fuzz_dep_.CoverTab[189276]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:487
		// _ = "end of CoverTab[189276]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:487
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:487
	// _ = "end of CoverTab[189264]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:487
	_go_fuzz_dep_.CoverTab[189265]++

											encryption := globalHeaders.getEncryption()
											cipher := getContentCipher(encryption)
											if cipher == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:491
		_go_fuzz_dep_.CoverTab[189277]++
												return -1, Header{}, nil, fmt.Errorf("square/go-jose: unsupported enc value '%s'", string(encryption))
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:492
		// _ = "end of CoverTab[189277]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:493
		_go_fuzz_dep_.CoverTab[189278]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:493
		// _ = "end of CoverTab[189278]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:493
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:493
	// _ = "end of CoverTab[189265]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:493
	_go_fuzz_dep_.CoverTab[189266]++

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
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:511
		_go_fuzz_dep_.CoverTab[189279]++
												recipientHeaders := obj.mergedHeaders(&recipient)

												cek, err := decrypter.decryptKey(recipientHeaders, &recipient, generator)
												if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:515
			_go_fuzz_dep_.CoverTab[189280]++

													plaintext, err = cipher.decrypt(cek, authData, parts)
													if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:518
				_go_fuzz_dep_.CoverTab[189281]++
														index = i
														headers = recipientHeaders
														break
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:521
				// _ = "end of CoverTab[189281]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:522
				_go_fuzz_dep_.CoverTab[189282]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:522
				// _ = "end of CoverTab[189282]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:522
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:522
			// _ = "end of CoverTab[189280]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:523
			_go_fuzz_dep_.CoverTab[189283]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:523
			// _ = "end of CoverTab[189283]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:523
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:523
		// _ = "end of CoverTab[189279]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:524
	// _ = "end of CoverTab[189266]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:524
	_go_fuzz_dep_.CoverTab[189267]++

											if plaintext == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:526
		_go_fuzz_dep_.CoverTab[189284]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:526
		return err != nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:526
		// _ = "end of CoverTab[189284]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:526
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:526
		_go_fuzz_dep_.CoverTab[189285]++
												return -1, Header{}, nil, ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:527
		// _ = "end of CoverTab[189285]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:528
		_go_fuzz_dep_.CoverTab[189286]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:528
		// _ = "end of CoverTab[189286]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:528
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:528
	// _ = "end of CoverTab[189267]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:528
	_go_fuzz_dep_.CoverTab[189268]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:531
	if comp := obj.protected.getCompression(); comp != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:531
		_go_fuzz_dep_.CoverTab[189287]++
												plaintext, err = decompress(comp, plaintext)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:532
		// _ = "end of CoverTab[189287]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:533
		_go_fuzz_dep_.CoverTab[189288]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:533
		// _ = "end of CoverTab[189288]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:533
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:533
	// _ = "end of CoverTab[189268]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:533
	_go_fuzz_dep_.CoverTab[189269]++

											sanitized, err := headers.sanitized()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:536
		_go_fuzz_dep_.CoverTab[189289]++
												return -1, Header{}, nil, fmt.Errorf("square/go-jose: failed to sanitize header: %v", err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:537
		// _ = "end of CoverTab[189289]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:538
		_go_fuzz_dep_.CoverTab[189290]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:538
		// _ = "end of CoverTab[189290]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:538
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:538
	// _ = "end of CoverTab[189269]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:538
	_go_fuzz_dep_.CoverTab[189270]++

											return index, sanitized, plaintext, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:540
	// _ = "end of CoverTab[189270]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:541
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/crypter.go:541
var _ = _go_fuzz_dep_.CoverTab
