//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:17
package jose

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:17
)

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/subtle"
	"errors"
	"fmt"
	"hash"
	"io"

	"golang.org/x/crypto/pbkdf2"
	"gopkg.in/square/go-jose.v2/cipher"
)

// Random reader (stubbed out in tests)
var RandReader = rand.Reader

const (
	// RFC7518 recommends a minimum of 1,000 iterations:
	// https://tools.ietf.org/html/rfc7518#section-4.8.1.2
	// NIST recommends a minimum of 10,000:
	// https://pages.nist.gov/800-63-3/sp800-63b.html
	// 1Password uses 100,000:
	// https://support.1password.com/pbkdf2/
	defaultP2C	= 100000
	// Default salt size: 128 bits
	defaultP2SSize	= 16
)

// Dummy key cipher for shared symmetric key mode
type symmetricKeyCipher struct {
	key	[]byte	// Pre-shared content-encryption key
	p2c	int	// PBES2 Count
	p2s	[]byte	// PBES2 Salt Input
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
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:87
	_go_fuzz_dep_.CoverTab[190201]++
											return &aeadContentCipher{
		keyBytes:	keySize,
		authtagBytes:	16,
		getAead: func(key []byte) (cipher.AEAD, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:91
			_go_fuzz_dep_.CoverTab[190202]++
													aes, err := aes.NewCipher(key)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:93
				_go_fuzz_dep_.CoverTab[190204]++
														return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:94
				// _ = "end of CoverTab[190204]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:95
				_go_fuzz_dep_.CoverTab[190205]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:95
				// _ = "end of CoverTab[190205]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:95
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:95
			// _ = "end of CoverTab[190202]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:95
			_go_fuzz_dep_.CoverTab[190203]++

													return cipher.NewGCM(aes)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:97
			// _ = "end of CoverTab[190203]"
		},
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:99
	// _ = "end of CoverTab[190201]"
}

// Create a new content cipher based on AES-CBC+HMAC
func newAESCBC(keySize int) contentCipher {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:103
	_go_fuzz_dep_.CoverTab[190206]++
											return &aeadContentCipher{
		keyBytes:	keySize * 2,
		authtagBytes:	keySize,
		getAead: func(key []byte) (cipher.AEAD, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:107
			_go_fuzz_dep_.CoverTab[190207]++
													return josecipher.NewCBCHMAC(key, aes.NewCipher)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:108
			// _ = "end of CoverTab[190207]"
		},
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:110
	// _ = "end of CoverTab[190206]"
}

// Get an AEAD cipher object for the given content encryption algorithm
func getContentCipher(alg ContentEncryption) contentCipher {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:114
	_go_fuzz_dep_.CoverTab[190208]++
											switch alg {
	case A128GCM:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:116
		_go_fuzz_dep_.CoverTab[190209]++
												return newAESGCM(16)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:117
		// _ = "end of CoverTab[190209]"
	case A192GCM:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:118
		_go_fuzz_dep_.CoverTab[190210]++
												return newAESGCM(24)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:119
		// _ = "end of CoverTab[190210]"
	case A256GCM:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:120
		_go_fuzz_dep_.CoverTab[190211]++
												return newAESGCM(32)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:121
		// _ = "end of CoverTab[190211]"
	case A128CBC_HS256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:122
		_go_fuzz_dep_.CoverTab[190212]++
												return newAESCBC(16)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:123
		// _ = "end of CoverTab[190212]"
	case A192CBC_HS384:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:124
		_go_fuzz_dep_.CoverTab[190213]++
												return newAESCBC(24)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:125
		// _ = "end of CoverTab[190213]"
	case A256CBC_HS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:126
		_go_fuzz_dep_.CoverTab[190214]++
												return newAESCBC(32)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:127
		// _ = "end of CoverTab[190214]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:128
		_go_fuzz_dep_.CoverTab[190215]++
												return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:129
		// _ = "end of CoverTab[190215]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:130
	// _ = "end of CoverTab[190208]"
}

// getPbkdf2Params returns the key length and hash function used in
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:133
// pbkdf2.Key.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:135
func getPbkdf2Params(alg KeyAlgorithm) (int, func() hash.Hash) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:135
	_go_fuzz_dep_.CoverTab[190216]++
											switch alg {
	case PBES2_HS256_A128KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:137
		_go_fuzz_dep_.CoverTab[190217]++
												return 16, sha256.New
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:138
		// _ = "end of CoverTab[190217]"
	case PBES2_HS384_A192KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:139
		_go_fuzz_dep_.CoverTab[190218]++
												return 24, sha512.New384
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:140
		// _ = "end of CoverTab[190218]"
	case PBES2_HS512_A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:141
		_go_fuzz_dep_.CoverTab[190219]++
												return 32, sha512.New
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:142
		// _ = "end of CoverTab[190219]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:143
		_go_fuzz_dep_.CoverTab[190220]++
												panic("invalid algorithm")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:144
		// _ = "end of CoverTab[190220]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:145
	// _ = "end of CoverTab[190216]"
}

// getRandomSalt generates a new salt of the given size.
func getRandomSalt(size int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:149
	_go_fuzz_dep_.CoverTab[190221]++
											salt := make([]byte, size)
											_, err := io.ReadFull(RandReader, salt)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:152
		_go_fuzz_dep_.CoverTab[190223]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:153
		// _ = "end of CoverTab[190223]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:154
		_go_fuzz_dep_.CoverTab[190224]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:154
		// _ = "end of CoverTab[190224]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:154
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:154
	// _ = "end of CoverTab[190221]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:154
	_go_fuzz_dep_.CoverTab[190222]++

											return salt, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:156
	// _ = "end of CoverTab[190222]"
}

// newSymmetricRecipient creates a JWE encrypter based on AES-GCM key wrap.
func newSymmetricRecipient(keyAlg KeyAlgorithm, key []byte) (recipientKeyInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:160
	_go_fuzz_dep_.CoverTab[190225]++
											switch keyAlg {
	case DIRECT, A128GCMKW, A192GCMKW, A256GCMKW, A128KW, A192KW, A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:162
		_go_fuzz_dep_.CoverTab[190227]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:162
		// _ = "end of CoverTab[190227]"
	case PBES2_HS256_A128KW, PBES2_HS384_A192KW, PBES2_HS512_A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:163
		_go_fuzz_dep_.CoverTab[190228]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:163
		// _ = "end of CoverTab[190228]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:164
		_go_fuzz_dep_.CoverTab[190229]++
												return recipientKeyInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:165
		// _ = "end of CoverTab[190229]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:166
	// _ = "end of CoverTab[190225]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:166
	_go_fuzz_dep_.CoverTab[190226]++

											return recipientKeyInfo{
		keyAlg:	keyAlg,
		keyEncrypter: &symmetricKeyCipher{
			key: key,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:173
	// _ = "end of CoverTab[190226]"
}

// newSymmetricSigner creates a recipientSigInfo based on the given key.
func newSymmetricSigner(sigAlg SignatureAlgorithm, key []byte) (recipientSigInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:177
	_go_fuzz_dep_.CoverTab[190230]++

											switch sigAlg {
	case HS256, HS384, HS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:180
		_go_fuzz_dep_.CoverTab[190232]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:180
		// _ = "end of CoverTab[190232]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:181
		_go_fuzz_dep_.CoverTab[190233]++
												return recipientSigInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:182
		// _ = "end of CoverTab[190233]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:183
	// _ = "end of CoverTab[190230]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:183
	_go_fuzz_dep_.CoverTab[190231]++

											return recipientSigInfo{
		sigAlg:	sigAlg,
		signer: &symmetricMac{
			key: key,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:190
	// _ = "end of CoverTab[190231]"
}

// Generate a random key for the given content cipher
func (ctx randomKeyGenerator) genKey() ([]byte, rawHeader, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:194
	_go_fuzz_dep_.CoverTab[190234]++
											key := make([]byte, ctx.size)
											_, err := io.ReadFull(RandReader, key)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:197
		_go_fuzz_dep_.CoverTab[190236]++
												return nil, rawHeader{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:198
		// _ = "end of CoverTab[190236]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:199
		_go_fuzz_dep_.CoverTab[190237]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:199
		// _ = "end of CoverTab[190237]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:199
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:199
	// _ = "end of CoverTab[190234]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:199
	_go_fuzz_dep_.CoverTab[190235]++

											return key, rawHeader{}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:201
	// _ = "end of CoverTab[190235]"
}

// Key size for random generator
func (ctx randomKeyGenerator) keySize() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:205
	_go_fuzz_dep_.CoverTab[190238]++
											return ctx.size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:206
	// _ = "end of CoverTab[190238]"
}

// Generate a static key (for direct mode)
func (ctx staticKeyGenerator) genKey() ([]byte, rawHeader, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:210
	_go_fuzz_dep_.CoverTab[190239]++
											cek := make([]byte, len(ctx.key))
											copy(cek, ctx.key)
											return cek, rawHeader{}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:213
	// _ = "end of CoverTab[190239]"
}

// Key size for static generator
func (ctx staticKeyGenerator) keySize() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:217
	_go_fuzz_dep_.CoverTab[190240]++
											return len(ctx.key)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:218
	// _ = "end of CoverTab[190240]"
}

// Get key size for this cipher
func (ctx aeadContentCipher) keySize() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:222
	_go_fuzz_dep_.CoverTab[190241]++
											return ctx.keyBytes
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:223
	// _ = "end of CoverTab[190241]"
}

// Encrypt some data
func (ctx aeadContentCipher) encrypt(key, aad, pt []byte) (*aeadParts, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:227
	_go_fuzz_dep_.CoverTab[190242]++

											aead, err := ctx.getAead(key)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:230
		_go_fuzz_dep_.CoverTab[190245]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:231
		// _ = "end of CoverTab[190245]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:232
		_go_fuzz_dep_.CoverTab[190246]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:232
		// _ = "end of CoverTab[190246]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:232
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:232
	// _ = "end of CoverTab[190242]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:232
	_go_fuzz_dep_.CoverTab[190243]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:235
	iv := make([]byte, aead.NonceSize())
	_, err = io.ReadFull(RandReader, iv)
	if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:237
		_go_fuzz_dep_.CoverTab[190247]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:238
		// _ = "end of CoverTab[190247]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:239
		_go_fuzz_dep_.CoverTab[190248]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:239
		// _ = "end of CoverTab[190248]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:239
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:239
	// _ = "end of CoverTab[190243]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:239
	_go_fuzz_dep_.CoverTab[190244]++

											ciphertextAndTag := aead.Seal(nil, iv, pt, aad)
											offset := len(ciphertextAndTag) - ctx.authtagBytes

											return &aeadParts{
		iv:		iv,
		ciphertext:	ciphertextAndTag[:offset],
		tag:		ciphertextAndTag[offset:],
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:248
	// _ = "end of CoverTab[190244]"
}

// Decrypt some data
func (ctx aeadContentCipher) decrypt(key, aad []byte, parts *aeadParts) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:252
	_go_fuzz_dep_.CoverTab[190249]++
											aead, err := ctx.getAead(key)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:254
		_go_fuzz_dep_.CoverTab[190252]++
												return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:255
		// _ = "end of CoverTab[190252]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:256
		_go_fuzz_dep_.CoverTab[190253]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:256
		// _ = "end of CoverTab[190253]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:256
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:256
	// _ = "end of CoverTab[190249]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:256
	_go_fuzz_dep_.CoverTab[190250]++

											if len(parts.iv) != aead.NonceSize() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:258
		_go_fuzz_dep_.CoverTab[190254]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:258
		return len(parts.tag) < ctx.authtagBytes
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:258
		// _ = "end of CoverTab[190254]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:258
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:258
		_go_fuzz_dep_.CoverTab[190255]++
												return nil, ErrCryptoFailure
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:259
		// _ = "end of CoverTab[190255]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:260
		_go_fuzz_dep_.CoverTab[190256]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:260
		// _ = "end of CoverTab[190256]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:260
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:260
	// _ = "end of CoverTab[190250]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:260
	_go_fuzz_dep_.CoverTab[190251]++

											return aead.Open(nil, parts.iv, append(parts.ciphertext, parts.tag...), aad)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:262
	// _ = "end of CoverTab[190251]"
}

// Encrypt the content encryption key.
func (ctx *symmetricKeyCipher) encryptKey(cek []byte, alg KeyAlgorithm) (recipientInfo, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:266
	_go_fuzz_dep_.CoverTab[190257]++
											switch alg {
	case DIRECT:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:268
		_go_fuzz_dep_.CoverTab[190259]++
												return recipientInfo{
			header: &rawHeader{},
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:271
		// _ = "end of CoverTab[190259]"
	case A128GCMKW, A192GCMKW, A256GCMKW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:272
		_go_fuzz_dep_.CoverTab[190260]++
												aead := newAESGCM(len(ctx.key))

												parts, err := aead.encrypt(ctx.key, []byte{}, cek)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:276
			_go_fuzz_dep_.CoverTab[190271]++
													return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:277
			// _ = "end of CoverTab[190271]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:278
			_go_fuzz_dep_.CoverTab[190272]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:278
			// _ = "end of CoverTab[190272]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:278
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:278
		// _ = "end of CoverTab[190260]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:278
		_go_fuzz_dep_.CoverTab[190261]++

												header := &rawHeader{}
												header.set(headerIV, newBuffer(parts.iv))
												header.set(headerTag, newBuffer(parts.tag))

												return recipientInfo{
			header:		header,
			encryptedKey:	parts.ciphertext,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:287
		// _ = "end of CoverTab[190261]"
	case A128KW, A192KW, A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:288
		_go_fuzz_dep_.CoverTab[190262]++
												block, err := aes.NewCipher(ctx.key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:290
			_go_fuzz_dep_.CoverTab[190273]++
													return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:291
			// _ = "end of CoverTab[190273]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:292
			_go_fuzz_dep_.CoverTab[190274]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:292
			// _ = "end of CoverTab[190274]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:292
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:292
		// _ = "end of CoverTab[190262]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:292
		_go_fuzz_dep_.CoverTab[190263]++

												jek, err := josecipher.KeyWrap(block, cek)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:295
			_go_fuzz_dep_.CoverTab[190275]++
													return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:296
			// _ = "end of CoverTab[190275]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:297
			_go_fuzz_dep_.CoverTab[190276]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:297
			// _ = "end of CoverTab[190276]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:297
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:297
		// _ = "end of CoverTab[190263]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:297
		_go_fuzz_dep_.CoverTab[190264]++

												return recipientInfo{
			encryptedKey:	jek,
			header:		&rawHeader{},
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:302
		// _ = "end of CoverTab[190264]"
	case PBES2_HS256_A128KW, PBES2_HS384_A192KW, PBES2_HS512_A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:303
		_go_fuzz_dep_.CoverTab[190265]++
												if len(ctx.p2s) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:304
			_go_fuzz_dep_.CoverTab[190277]++
													salt, err := getRandomSalt(defaultP2SSize)
													if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:306
				_go_fuzz_dep_.CoverTab[190279]++
														return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:307
				// _ = "end of CoverTab[190279]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:308
				_go_fuzz_dep_.CoverTab[190280]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:308
				// _ = "end of CoverTab[190280]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:308
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:308
			// _ = "end of CoverTab[190277]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:308
			_go_fuzz_dep_.CoverTab[190278]++
													ctx.p2s = salt
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:309
			// _ = "end of CoverTab[190278]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:310
			_go_fuzz_dep_.CoverTab[190281]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:310
			// _ = "end of CoverTab[190281]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:310
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:310
		// _ = "end of CoverTab[190265]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:310
		_go_fuzz_dep_.CoverTab[190266]++

												if ctx.p2c <= 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:312
			_go_fuzz_dep_.CoverTab[190282]++
													ctx.p2c = defaultP2C
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:313
			// _ = "end of CoverTab[190282]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:314
			_go_fuzz_dep_.CoverTab[190283]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:314
			// _ = "end of CoverTab[190283]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:314
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:314
		// _ = "end of CoverTab[190266]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:314
		_go_fuzz_dep_.CoverTab[190267]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:317
		salt := bytes.Join([][]byte{[]byte(alg), ctx.p2s}, []byte{0x00})

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:320
		keyLen, h := getPbkdf2Params(alg)
												key := pbkdf2.Key(ctx.key, salt, ctx.p2c, keyLen, h)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:324
		block, err := aes.NewCipher(key)
		if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:325
			_go_fuzz_dep_.CoverTab[190284]++
													return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:326
			// _ = "end of CoverTab[190284]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:327
			_go_fuzz_dep_.CoverTab[190285]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:327
			// _ = "end of CoverTab[190285]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:327
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:327
		// _ = "end of CoverTab[190267]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:327
		_go_fuzz_dep_.CoverTab[190268]++

												jek, err := josecipher.KeyWrap(block, cek)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:330
			_go_fuzz_dep_.CoverTab[190286]++
													return recipientInfo{}, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:331
			// _ = "end of CoverTab[190286]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:332
			_go_fuzz_dep_.CoverTab[190287]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:332
			// _ = "end of CoverTab[190287]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:332
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:332
		// _ = "end of CoverTab[190268]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:332
		_go_fuzz_dep_.CoverTab[190269]++

												header := &rawHeader{}
												header.set(headerP2C, ctx.p2c)
												header.set(headerP2S, newBuffer(ctx.p2s))

												return recipientInfo{
			encryptedKey:	jek,
			header:		header,
		}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:341
		// _ = "end of CoverTab[190269]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:341
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:341
		_go_fuzz_dep_.CoverTab[190270]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:341
		// _ = "end of CoverTab[190270]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:342
	// _ = "end of CoverTab[190257]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:342
	_go_fuzz_dep_.CoverTab[190258]++

											return recipientInfo{}, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:344
	// _ = "end of CoverTab[190258]"
}

// Decrypt the content encryption key.
func (ctx *symmetricKeyCipher) decryptKey(headers rawHeader, recipient *recipientInfo, generator keyGenerator) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:348
	_go_fuzz_dep_.CoverTab[190288]++
											switch headers.getAlgorithm() {
	case DIRECT:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:350
		_go_fuzz_dep_.CoverTab[190290]++
												cek := make([]byte, len(ctx.key))
												copy(cek, ctx.key)
												return cek, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:353
		// _ = "end of CoverTab[190290]"
	case A128GCMKW, A192GCMKW, A256GCMKW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:354
		_go_fuzz_dep_.CoverTab[190291]++
												aead := newAESGCM(len(ctx.key))

												iv, err := headers.getIV()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:358
			_go_fuzz_dep_.CoverTab[190306]++
													return nil, fmt.Errorf("square/go-jose: invalid IV: %v", err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:359
			// _ = "end of CoverTab[190306]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:360
			_go_fuzz_dep_.CoverTab[190307]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:360
			// _ = "end of CoverTab[190307]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:360
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:360
		// _ = "end of CoverTab[190291]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:360
		_go_fuzz_dep_.CoverTab[190292]++
												tag, err := headers.getTag()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:362
			_go_fuzz_dep_.CoverTab[190308]++
													return nil, fmt.Errorf("square/go-jose: invalid tag: %v", err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:363
			// _ = "end of CoverTab[190308]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:364
			_go_fuzz_dep_.CoverTab[190309]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:364
			// _ = "end of CoverTab[190309]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:364
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:364
		// _ = "end of CoverTab[190292]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:364
		_go_fuzz_dep_.CoverTab[190293]++

												parts := &aeadParts{
			iv:		iv.bytes(),
			ciphertext:	recipient.encryptedKey,
			tag:		tag.bytes(),
		}

		cek, err := aead.decrypt(ctx.key, []byte{}, parts)
		if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:373
			_go_fuzz_dep_.CoverTab[190310]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:374
			// _ = "end of CoverTab[190310]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:375
			_go_fuzz_dep_.CoverTab[190311]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:375
			// _ = "end of CoverTab[190311]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:375
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:375
		// _ = "end of CoverTab[190293]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:375
		_go_fuzz_dep_.CoverTab[190294]++

												return cek, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:377
		// _ = "end of CoverTab[190294]"
	case A128KW, A192KW, A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:378
		_go_fuzz_dep_.CoverTab[190295]++
												block, err := aes.NewCipher(ctx.key)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:380
			_go_fuzz_dep_.CoverTab[190312]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:381
			// _ = "end of CoverTab[190312]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:382
			_go_fuzz_dep_.CoverTab[190313]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:382
			// _ = "end of CoverTab[190313]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:382
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:382
		// _ = "end of CoverTab[190295]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:382
		_go_fuzz_dep_.CoverTab[190296]++

												cek, err := josecipher.KeyUnwrap(block, recipient.encryptedKey)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:385
			_go_fuzz_dep_.CoverTab[190314]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:386
			// _ = "end of CoverTab[190314]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:387
			_go_fuzz_dep_.CoverTab[190315]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:387
			// _ = "end of CoverTab[190315]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:387
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:387
		// _ = "end of CoverTab[190296]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:387
		_go_fuzz_dep_.CoverTab[190297]++
												return cek, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:388
		// _ = "end of CoverTab[190297]"
	case PBES2_HS256_A128KW, PBES2_HS384_A192KW, PBES2_HS512_A256KW:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:389
		_go_fuzz_dep_.CoverTab[190298]++
												p2s, err := headers.getP2S()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:391
			_go_fuzz_dep_.CoverTab[190316]++
													return nil, fmt.Errorf("square/go-jose: invalid P2S: %v", err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:392
			// _ = "end of CoverTab[190316]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:393
			_go_fuzz_dep_.CoverTab[190317]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:393
			// _ = "end of CoverTab[190317]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:393
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:393
		// _ = "end of CoverTab[190298]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:393
		_go_fuzz_dep_.CoverTab[190299]++
												if p2s == nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:394
			_go_fuzz_dep_.CoverTab[190318]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:394
			return len(p2s.data) == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:394
			// _ = "end of CoverTab[190318]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:394
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:394
			_go_fuzz_dep_.CoverTab[190319]++
													return nil, fmt.Errorf("square/go-jose: invalid P2S: must be present")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:395
			// _ = "end of CoverTab[190319]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:396
			_go_fuzz_dep_.CoverTab[190320]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:396
			// _ = "end of CoverTab[190320]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:396
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:396
		// _ = "end of CoverTab[190299]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:396
		_go_fuzz_dep_.CoverTab[190300]++

												p2c, err := headers.getP2C()
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:399
			_go_fuzz_dep_.CoverTab[190321]++
													return nil, fmt.Errorf("square/go-jose: invalid P2C: %v", err)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:400
			// _ = "end of CoverTab[190321]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:401
			_go_fuzz_dep_.CoverTab[190322]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:401
			// _ = "end of CoverTab[190322]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:401
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:401
		// _ = "end of CoverTab[190300]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:401
		_go_fuzz_dep_.CoverTab[190301]++
												if p2c <= 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:402
			_go_fuzz_dep_.CoverTab[190323]++
													return nil, fmt.Errorf("square/go-jose: invalid P2C: must be a positive integer")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:403
			// _ = "end of CoverTab[190323]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:404
			_go_fuzz_dep_.CoverTab[190324]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:404
			// _ = "end of CoverTab[190324]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:404
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:404
		// _ = "end of CoverTab[190301]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:404
		_go_fuzz_dep_.CoverTab[190302]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:407
		alg := headers.getAlgorithm()
												salt := bytes.Join([][]byte{[]byte(alg), p2s.bytes()}, []byte{0x00})

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:411
		keyLen, h := getPbkdf2Params(alg)
												key := pbkdf2.Key(ctx.key, salt, p2c, keyLen, h)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:415
		block, err := aes.NewCipher(key)
		if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:416
			_go_fuzz_dep_.CoverTab[190325]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:417
			// _ = "end of CoverTab[190325]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:418
			_go_fuzz_dep_.CoverTab[190326]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:418
			// _ = "end of CoverTab[190326]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:418
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:418
		// _ = "end of CoverTab[190302]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:418
		_go_fuzz_dep_.CoverTab[190303]++

												cek, err := josecipher.KeyUnwrap(block, recipient.encryptedKey)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:421
			_go_fuzz_dep_.CoverTab[190327]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:422
			// _ = "end of CoverTab[190327]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:423
			_go_fuzz_dep_.CoverTab[190328]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:423
			// _ = "end of CoverTab[190328]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:423
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:423
		// _ = "end of CoverTab[190303]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:423
		_go_fuzz_dep_.CoverTab[190304]++
												return cek, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:424
		// _ = "end of CoverTab[190304]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:424
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:424
		_go_fuzz_dep_.CoverTab[190305]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:424
		// _ = "end of CoverTab[190305]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:425
	// _ = "end of CoverTab[190288]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:425
	_go_fuzz_dep_.CoverTab[190289]++

											return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:427
	// _ = "end of CoverTab[190289]"
}

// Sign the given payload
func (ctx symmetricMac) signPayload(payload []byte, alg SignatureAlgorithm) (Signature, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:431
	_go_fuzz_dep_.CoverTab[190329]++
											mac, err := ctx.hmac(payload, alg)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:433
		_go_fuzz_dep_.CoverTab[190331]++
												return Signature{}, errors.New("square/go-jose: failed to compute hmac")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:434
		// _ = "end of CoverTab[190331]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:435
		_go_fuzz_dep_.CoverTab[190332]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:435
		// _ = "end of CoverTab[190332]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:435
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:435
	// _ = "end of CoverTab[190329]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:435
	_go_fuzz_dep_.CoverTab[190330]++

											return Signature{
		Signature:	mac,
		protected:	&rawHeader{},
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:440
	// _ = "end of CoverTab[190330]"
}

// Verify the given payload
func (ctx symmetricMac) verifyPayload(payload []byte, mac []byte, alg SignatureAlgorithm) error {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:444
	_go_fuzz_dep_.CoverTab[190333]++
											expected, err := ctx.hmac(payload, alg)
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:446
		_go_fuzz_dep_.CoverTab[190337]++
												return errors.New("square/go-jose: failed to compute hmac")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:447
		// _ = "end of CoverTab[190337]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:448
		_go_fuzz_dep_.CoverTab[190338]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:448
		// _ = "end of CoverTab[190338]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:448
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:448
	// _ = "end of CoverTab[190333]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:448
	_go_fuzz_dep_.CoverTab[190334]++

											if len(mac) != len(expected) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:450
		_go_fuzz_dep_.CoverTab[190339]++
												return errors.New("square/go-jose: invalid hmac")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:451
		// _ = "end of CoverTab[190339]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:452
		_go_fuzz_dep_.CoverTab[190340]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:452
		// _ = "end of CoverTab[190340]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:452
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:452
	// _ = "end of CoverTab[190334]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:452
	_go_fuzz_dep_.CoverTab[190335]++

											match := subtle.ConstantTimeCompare(mac, expected)
											if match != 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:455
		_go_fuzz_dep_.CoverTab[190341]++
												return errors.New("square/go-jose: invalid hmac")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:456
		// _ = "end of CoverTab[190341]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:457
		_go_fuzz_dep_.CoverTab[190342]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:457
		// _ = "end of CoverTab[190342]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:457
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:457
	// _ = "end of CoverTab[190335]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:457
	_go_fuzz_dep_.CoverTab[190336]++

											return nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:459
	// _ = "end of CoverTab[190336]"
}

// Compute the HMAC based on the given alg value
func (ctx symmetricMac) hmac(payload []byte, alg SignatureAlgorithm) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:463
	_go_fuzz_dep_.CoverTab[190343]++
											var hash func() hash.Hash

											switch alg {
	case HS256:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:467
		_go_fuzz_dep_.CoverTab[190345]++
												hash = sha256.New
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:468
		// _ = "end of CoverTab[190345]"
	case HS384:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:469
		_go_fuzz_dep_.CoverTab[190346]++
												hash = sha512.New384
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:470
		// _ = "end of CoverTab[190346]"
	case HS512:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:471
		_go_fuzz_dep_.CoverTab[190347]++
												hash = sha512.New
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:472
		// _ = "end of CoverTab[190347]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:473
		_go_fuzz_dep_.CoverTab[190348]++
												return nil, ErrUnsupportedAlgorithm
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:474
		// _ = "end of CoverTab[190348]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:475
	// _ = "end of CoverTab[190343]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:475
	_go_fuzz_dep_.CoverTab[190344]++

											hmac := hmac.New(hash, ctx.key)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:480
	_, _ = hmac.Write(payload)
											return hmac.Sum(nil), nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:481
	// _ = "end of CoverTab[190344]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:482
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/symmetric.go:482
var _ = _go_fuzz_dep_.CoverTab
