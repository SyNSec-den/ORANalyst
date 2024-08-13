//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:17
)

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/subtle"
	"errors"
	"hash"
	"io"

	"gopkg.in/square/go-jose.v1/cipher"
)

// Random reader (stubbed out in tests)
var randReader = rand.Reader

// Dummy key cipher for shared symmetric key mode
type symmetricKeyCipher struct {
	key []byte	// Pre-shared content-encryption key
}

// Signer/verifier for MAC modes
type symmetricMac struct {
	key []byte
}

// Input/output from an AEAD operation
type aeadParts struct {
	iv, ciphertext, tag []byte
}

// A content cipher based on an AEAD construction
type aeadContentCipher struct {
	keyBytes	int
	authtagBytes	int
	getAead		func(key []byte) (cipher.AEAD, error)
}

// Random key generator
type randomKeyGenerator struct {
	size int
}

// Static key generator
type staticKeyGenerator struct {
	key []byte
}

// Create a new content cipher based on AES-GCM
func newAESGCM(keySize int) contentCipher {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:70
	_go_fuzz_dep_.CoverTab[186652]++
											return &aeadContentCipher{
		keyBytes:	keySize,
		authtagBytes:	16,
		getAead: func(key []byte) (cipher.AEAD, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:74
			_go_fuzz_dep_.CoverTab[186653]++
													aes, err := aes.NewCipher(key)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:76
				_go_fuzz_dep_.CoverTab[186655]++
														return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:77
				// _ = "end of CoverTab[186655]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:78
				_go_fuzz_dep_.CoverTab[186656]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:78
				// _ = "end of CoverTab[186656]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:78
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:78
			// _ = "end of CoverTab[186653]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:78
			_go_fuzz_dep_.CoverTab[186654]++

													return cipher.NewGCM(aes)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:80
			// _ = "end of CoverTab[186654]"
		},
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:82
	// _ = "end of CoverTab[186652]"
}

// Create a new content cipher based on AES-CBC+HMAC
func newAESCBC(keySize int) contentCipher {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:86
	_go_fuzz_dep_.CoverTab[186657]++
											return &aeadContentCipher{
		keyBytes:	keySize * 2,
		authtagBytes:	16,
		getAead: func(key []byte) (cipher.AEAD, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:90
			_go_fuzz_dep_.CoverTab[186658]++
													return josecipher.NewCBCHMAC(key, aes.NewCipher)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:91
			// _ = "end of CoverTab[186658]"
		},
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:93
	// _ = "end of CoverTab[186657]"
}

// Get an AEAD cipher object for the given content encryption algorithm
func getContentCipher(alg ContentEncryption) contentCipher {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:97
	_go_fuzz_dep_.CoverTab[186659]++
											switch alg {
	case A128GCM:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:99
		_go_fuzz_dep_.CoverTab[186660]++
												return newAESGCM(16)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:100
		// _ = "end of CoverTab[186660]"
	case A192GCM:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:101
		_go_fuzz_dep_.CoverTab[186661]++
												return newAESGCM(24)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:102
		// _ = "end of CoverTab[186661]"
	case A256GCM:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:103
		_go_fuzz_dep_.CoverTab[186662]++
												return newAESGCM(32)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:104
		// _ = "end of CoverTab[186662]"
	case A128CBC_HS256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:105
		_go_fuzz_dep_.CoverTab[186663]++
												return newAESCBC(16)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:106
		// _ = "end of CoverTab[186663]"
	case A192CBC_HS384:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:107
		_go_fuzz_dep_.CoverTab[186664]++
												return newAESCBC(24)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:108
		// _ = "end of CoverTab[186664]"
	case A256CBC_HS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:109
		_go_fuzz_dep_.CoverTab[186665]++
												return newAESCBC(32)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:110
		// _ = "end of CoverTab[186665]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:111
		_go_fuzz_dep_.CoverTab[186666]++
												return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:112
		// _ = "end of CoverTab[186666]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:113
	// _ = "end of CoverTab[186659]"
}

// newSymmetricRecipient creates a JWE encrypter based on AES-GCM key wrap.
func newSymmetricRecipient(keyAlg KeyAlgorithm, key []byte) (recipientKeyInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:117
	_go_fuzz_dep_.CoverTab[186667]++
											switch keyAlg {
	case DIRECT, A128GCMKW, A192GCMKW, A256GCMKW, A128KW, A192KW, A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:119
		_go_fuzz_dep_.CoverTab[186669]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:119
		// _ = "end of CoverTab[186669]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:120
		_go_fuzz_dep_.CoverTab[186670]++
												return recipientKeyInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:121
		// _ = "end of CoverTab[186670]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:122
	// _ = "end of CoverTab[186667]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:122
	_go_fuzz_dep_.CoverTab[186668]++

											return recipientKeyInfo{
		keyAlg:	keyAlg,
		keyEncrypter: &symmetricKeyCipher{
			key: key,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:129
	// _ = "end of CoverTab[186668]"
}

// newSymmetricSigner creates a recipientSigInfo based on the given key.
func newSymmetricSigner(sigAlg SignatureAlgorithm, key []byte) (recipientSigInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:133
	_go_fuzz_dep_.CoverTab[186671]++

											switch sigAlg {
	case HS256, HS384, HS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:136
		_go_fuzz_dep_.CoverTab[186673]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:136
		// _ = "end of CoverTab[186673]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:137
		_go_fuzz_dep_.CoverTab[186674]++
												return recipientSigInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:138
		// _ = "end of CoverTab[186674]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:139
	// _ = "end of CoverTab[186671]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:139
	_go_fuzz_dep_.CoverTab[186672]++

											return recipientSigInfo{
		sigAlg:	sigAlg,
		signer: &symmetricMac{
			key: key,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:146
	// _ = "end of CoverTab[186672]"
}

// Generate a random key for the given content cipher
func (ctx randomKeyGenerator) genKey() ([]byte, rawHeader, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:150
	_go_fuzz_dep_.CoverTab[186675]++
											key := make([]byte, ctx.size)
											_, err := io.ReadFull(randReader, key)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:153
		_go_fuzz_dep_.CoverTab[186677]++
												return nil, rawHeader{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:154
		// _ = "end of CoverTab[186677]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:155
		_go_fuzz_dep_.CoverTab[186678]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:155
		// _ = "end of CoverTab[186678]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:155
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:155
	// _ = "end of CoverTab[186675]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:155
	_go_fuzz_dep_.CoverTab[186676]++

											return key, rawHeader{}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:157
	// _ = "end of CoverTab[186676]"
}

// Key size for random generator
func (ctx randomKeyGenerator) keySize() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:161
	_go_fuzz_dep_.CoverTab[186679]++
											return ctx.size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:162
	// _ = "end of CoverTab[186679]"
}

// Generate a static key (for direct mode)
func (ctx staticKeyGenerator) genKey() ([]byte, rawHeader, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:166
	_go_fuzz_dep_.CoverTab[186680]++
											cek := make([]byte, len(ctx.key))
											copy(cek, ctx.key)
											return cek, rawHeader{}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:169
	// _ = "end of CoverTab[186680]"
}

// Key size for static generator
func (ctx staticKeyGenerator) keySize() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:173
	_go_fuzz_dep_.CoverTab[186681]++
											return len(ctx.key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:174
	// _ = "end of CoverTab[186681]"
}

// Get key size for this cipher
func (ctx aeadContentCipher) keySize() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:178
	_go_fuzz_dep_.CoverTab[186682]++
											return ctx.keyBytes
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:179
	// _ = "end of CoverTab[186682]"
}

// Encrypt some data
func (ctx aeadContentCipher) encrypt(key, aad, pt []byte) (*aeadParts, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:183
	_go_fuzz_dep_.CoverTab[186683]++

											aead, err := ctx.getAead(key)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:186
		_go_fuzz_dep_.CoverTab[186686]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:187
		// _ = "end of CoverTab[186686]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:188
		_go_fuzz_dep_.CoverTab[186687]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:188
		// _ = "end of CoverTab[186687]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:188
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:188
	// _ = "end of CoverTab[186683]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:188
	_go_fuzz_dep_.CoverTab[186684]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:191
	iv := make([]byte, aead.NonceSize())
	_, err = io.ReadFull(randReader, iv)
	if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:193
		_go_fuzz_dep_.CoverTab[186688]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:194
		// _ = "end of CoverTab[186688]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:195
		_go_fuzz_dep_.CoverTab[186689]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:195
		// _ = "end of CoverTab[186689]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:195
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:195
	// _ = "end of CoverTab[186684]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:195
	_go_fuzz_dep_.CoverTab[186685]++

											ciphertextAndTag := aead.Seal(nil, iv, pt, aad)
											offset := len(ciphertextAndTag) - ctx.authtagBytes

											return &aeadParts{
		iv:		iv,
		ciphertext:	ciphertextAndTag[:offset],
		tag:		ciphertextAndTag[offset:],
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:204
	// _ = "end of CoverTab[186685]"
}

// Decrypt some data
func (ctx aeadContentCipher) decrypt(key, aad []byte, parts *aeadParts) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:208
	_go_fuzz_dep_.CoverTab[186690]++
											aead, err := ctx.getAead(key)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:210
		_go_fuzz_dep_.CoverTab[186693]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:211
		// _ = "end of CoverTab[186693]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:212
		_go_fuzz_dep_.CoverTab[186694]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:212
		// _ = "end of CoverTab[186694]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:212
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:212
	// _ = "end of CoverTab[186690]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:212
	_go_fuzz_dep_.CoverTab[186691]++

											if len(parts.iv) < aead.NonceSize() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:214
		_go_fuzz_dep_.CoverTab[186695]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:214
		return len(parts.tag) < ctx.authtagBytes
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:214
		// _ = "end of CoverTab[186695]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:214
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:214
		_go_fuzz_dep_.CoverTab[186696]++
												return nil, ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:215
		// _ = "end of CoverTab[186696]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:216
		_go_fuzz_dep_.CoverTab[186697]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:216
		// _ = "end of CoverTab[186697]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:216
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:216
	// _ = "end of CoverTab[186691]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:216
	_go_fuzz_dep_.CoverTab[186692]++

											return aead.Open(nil, parts.iv, append(parts.ciphertext, parts.tag...), aad)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:218
	// _ = "end of CoverTab[186692]"
}

// Encrypt the content encryption key.
func (ctx *symmetricKeyCipher) encryptKey(cek []byte, alg KeyAlgorithm) (recipientInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:222
	_go_fuzz_dep_.CoverTab[186698]++
											switch alg {
	case DIRECT:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:224
		_go_fuzz_dep_.CoverTab[186700]++
												return recipientInfo{
			header: &rawHeader{},
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:227
		// _ = "end of CoverTab[186700]"
	case A128GCMKW, A192GCMKW, A256GCMKW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:228
		_go_fuzz_dep_.CoverTab[186701]++
												aead := newAESGCM(len(ctx.key))

												parts, err := aead.encrypt(ctx.key, []byte{}, cek)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:232
			_go_fuzz_dep_.CoverTab[186707]++
													return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:233
			// _ = "end of CoverTab[186707]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:234
			_go_fuzz_dep_.CoverTab[186708]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:234
			// _ = "end of CoverTab[186708]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:234
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:234
		// _ = "end of CoverTab[186701]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:234
		_go_fuzz_dep_.CoverTab[186702]++

												return recipientInfo{
			header: &rawHeader{
				Iv:	newBuffer(parts.iv),
				Tag:	newBuffer(parts.tag),
			},
			encryptedKey:	parts.ciphertext,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:242
		// _ = "end of CoverTab[186702]"
	case A128KW, A192KW, A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:243
		_go_fuzz_dep_.CoverTab[186703]++
												block, err := aes.NewCipher(ctx.key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:245
			_go_fuzz_dep_.CoverTab[186709]++
													return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:246
			// _ = "end of CoverTab[186709]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:247
			_go_fuzz_dep_.CoverTab[186710]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:247
			// _ = "end of CoverTab[186710]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:247
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:247
		// _ = "end of CoverTab[186703]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:247
		_go_fuzz_dep_.CoverTab[186704]++

												jek, err := josecipher.KeyWrap(block, cek)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:250
			_go_fuzz_dep_.CoverTab[186711]++
													return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:251
			// _ = "end of CoverTab[186711]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:252
			_go_fuzz_dep_.CoverTab[186712]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:252
			// _ = "end of CoverTab[186712]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:252
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:252
		// _ = "end of CoverTab[186704]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:252
		_go_fuzz_dep_.CoverTab[186705]++

												return recipientInfo{
			encryptedKey:	jek,
			header:		&rawHeader{},
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:257
		// _ = "end of CoverTab[186705]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:257
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:257
		_go_fuzz_dep_.CoverTab[186706]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:257
		// _ = "end of CoverTab[186706]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:258
	// _ = "end of CoverTab[186698]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:258
	_go_fuzz_dep_.CoverTab[186699]++

											return recipientInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:260
	// _ = "end of CoverTab[186699]"
}

// Decrypt the content encryption key.
func (ctx *symmetricKeyCipher) decryptKey(headers rawHeader, recipient *recipientInfo, generator keyGenerator) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:264
	_go_fuzz_dep_.CoverTab[186713]++
											switch KeyAlgorithm(headers.Alg) {
	case DIRECT:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:266
		_go_fuzz_dep_.CoverTab[186715]++
												cek := make([]byte, len(ctx.key))
												copy(cek, ctx.key)
												return cek, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:269
		// _ = "end of CoverTab[186715]"
	case A128GCMKW, A192GCMKW, A256GCMKW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:270
		_go_fuzz_dep_.CoverTab[186716]++
												aead := newAESGCM(len(ctx.key))

												parts := &aeadParts{
			iv:		headers.Iv.bytes(),
			ciphertext:	recipient.encryptedKey,
			tag:		headers.Tag.bytes(),
		}

		cek, err := aead.decrypt(ctx.key, []byte{}, parts)
		if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:280
			_go_fuzz_dep_.CoverTab[186722]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:281
			// _ = "end of CoverTab[186722]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:282
			_go_fuzz_dep_.CoverTab[186723]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:282
			// _ = "end of CoverTab[186723]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:282
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:282
		// _ = "end of CoverTab[186716]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:282
		_go_fuzz_dep_.CoverTab[186717]++

												return cek, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:284
		// _ = "end of CoverTab[186717]"
	case A128KW, A192KW, A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:285
		_go_fuzz_dep_.CoverTab[186718]++
												block, err := aes.NewCipher(ctx.key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:287
			_go_fuzz_dep_.CoverTab[186724]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:288
			// _ = "end of CoverTab[186724]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:289
			_go_fuzz_dep_.CoverTab[186725]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:289
			// _ = "end of CoverTab[186725]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:289
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:289
		// _ = "end of CoverTab[186718]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:289
		_go_fuzz_dep_.CoverTab[186719]++

												cek, err := josecipher.KeyUnwrap(block, recipient.encryptedKey)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:292
			_go_fuzz_dep_.CoverTab[186726]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:293
			// _ = "end of CoverTab[186726]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:294
			_go_fuzz_dep_.CoverTab[186727]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:294
			// _ = "end of CoverTab[186727]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:294
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:294
		// _ = "end of CoverTab[186719]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:294
		_go_fuzz_dep_.CoverTab[186720]++
												return cek, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:295
		// _ = "end of CoverTab[186720]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:295
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:295
		_go_fuzz_dep_.CoverTab[186721]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:295
		// _ = "end of CoverTab[186721]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:296
	// _ = "end of CoverTab[186713]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:296
	_go_fuzz_dep_.CoverTab[186714]++

											return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:298
	// _ = "end of CoverTab[186714]"
}

// Sign the given payload
func (ctx symmetricMac) signPayload(payload []byte, alg SignatureAlgorithm) (Signature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:302
	_go_fuzz_dep_.CoverTab[186728]++
											mac, err := ctx.hmac(payload, alg)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:304
		_go_fuzz_dep_.CoverTab[186730]++
												return Signature{}, errors.New("square/go-jose: failed to compute hmac")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:305
		// _ = "end of CoverTab[186730]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:306
		_go_fuzz_dep_.CoverTab[186731]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:306
		// _ = "end of CoverTab[186731]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:306
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:306
	// _ = "end of CoverTab[186728]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:306
	_go_fuzz_dep_.CoverTab[186729]++

											return Signature{
		Signature:	mac,
		protected:	&rawHeader{},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:311
	// _ = "end of CoverTab[186729]"
}

// Verify the given payload
func (ctx symmetricMac) verifyPayload(payload []byte, mac []byte, alg SignatureAlgorithm) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:315
	_go_fuzz_dep_.CoverTab[186732]++
											expected, err := ctx.hmac(payload, alg)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:317
		_go_fuzz_dep_.CoverTab[186736]++
												return errors.New("square/go-jose: failed to compute hmac")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:318
		// _ = "end of CoverTab[186736]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:319
		_go_fuzz_dep_.CoverTab[186737]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:319
		// _ = "end of CoverTab[186737]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:319
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:319
	// _ = "end of CoverTab[186732]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:319
	_go_fuzz_dep_.CoverTab[186733]++

											if len(mac) != len(expected) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:321
		_go_fuzz_dep_.CoverTab[186738]++
												return errors.New("square/go-jose: invalid hmac")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:322
		// _ = "end of CoverTab[186738]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:323
		_go_fuzz_dep_.CoverTab[186739]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:323
		// _ = "end of CoverTab[186739]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:323
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:323
	// _ = "end of CoverTab[186733]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:323
	_go_fuzz_dep_.CoverTab[186734]++

											match := subtle.ConstantTimeCompare(mac, expected)
											if match != 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:326
		_go_fuzz_dep_.CoverTab[186740]++
												return errors.New("square/go-jose: invalid hmac")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:327
		// _ = "end of CoverTab[186740]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:328
		_go_fuzz_dep_.CoverTab[186741]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:328
		// _ = "end of CoverTab[186741]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:328
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:328
	// _ = "end of CoverTab[186734]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:328
	_go_fuzz_dep_.CoverTab[186735]++

											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:330
	// _ = "end of CoverTab[186735]"
}

// Compute the HMAC based on the given alg value
func (ctx symmetricMac) hmac(payload []byte, alg SignatureAlgorithm) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:334
	_go_fuzz_dep_.CoverTab[186742]++
											var hash func() hash.Hash

											switch alg {
	case HS256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:338
		_go_fuzz_dep_.CoverTab[186744]++
												hash = sha256.New
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:339
		// _ = "end of CoverTab[186744]"
	case HS384:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:340
		_go_fuzz_dep_.CoverTab[186745]++
												hash = sha512.New384
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:341
		// _ = "end of CoverTab[186745]"
	case HS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:342
		_go_fuzz_dep_.CoverTab[186746]++
												hash = sha512.New
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:343
		// _ = "end of CoverTab[186746]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:344
		_go_fuzz_dep_.CoverTab[186747]++
												return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:345
		// _ = "end of CoverTab[186747]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:346
	// _ = "end of CoverTab[186742]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:346
	_go_fuzz_dep_.CoverTab[186743]++

											hmac := hmac.New(hash, ctx.key)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:351
	_, _ = hmac.Write(payload)
											return hmac.Sum(nil), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:352
	// _ = "end of CoverTab[186743]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:353
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/symmetric.go:353
var _ = _go_fuzz_dep_.CoverTab
