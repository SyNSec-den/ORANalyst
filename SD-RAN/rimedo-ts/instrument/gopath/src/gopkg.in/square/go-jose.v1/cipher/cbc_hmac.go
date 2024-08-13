//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:17
package josecipher

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:17
)

import (
	"bytes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/subtle"
	"encoding/binary"
	"errors"
	"hash"
)

const (
	nonceBytes = 16
)

// NewCBCHMAC instantiates a new AEAD based on CBC+HMAC.
func NewCBCHMAC(key []byte, newBlockCipher func([]byte) (cipher.Block, error)) (cipher.AEAD, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:36
	_go_fuzz_dep_.CoverTab[184293]++
												keySize := len(key) / 2
												integrityKey := key[:keySize]
												encryptionKey := key[keySize:]

												blockCipher, err := newBlockCipher(encryptionKey)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:42
		_go_fuzz_dep_.CoverTab[184296]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:43
		// _ = "end of CoverTab[184296]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:44
		_go_fuzz_dep_.CoverTab[184297]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:44
		// _ = "end of CoverTab[184297]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:44
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:44
	// _ = "end of CoverTab[184293]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:44
	_go_fuzz_dep_.CoverTab[184294]++

												var hash func() hash.Hash
												switch keySize {
	case 16:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:48
		_go_fuzz_dep_.CoverTab[184298]++
													hash = sha256.New
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:49
		// _ = "end of CoverTab[184298]"
	case 24:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:50
		_go_fuzz_dep_.CoverTab[184299]++
													hash = sha512.New384
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:51
		// _ = "end of CoverTab[184299]"
	case 32:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:52
		_go_fuzz_dep_.CoverTab[184300]++
													hash = sha512.New
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:53
		// _ = "end of CoverTab[184300]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:53
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:53
		_go_fuzz_dep_.CoverTab[184301]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:53
		// _ = "end of CoverTab[184301]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:54
	// _ = "end of CoverTab[184294]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:54
	_go_fuzz_dep_.CoverTab[184295]++

												return &cbcAEAD{
		hash:		hash,
		blockCipher:	blockCipher,
		authtagBytes:	keySize,
		integrityKey:	integrityKey,
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:61
	// _ = "end of CoverTab[184295]"
}

// An AEAD based on CBC+HMAC
type cbcAEAD struct {
	hash		func() hash.Hash
	authtagBytes	int
	integrityKey	[]byte
	blockCipher	cipher.Block
}

func (ctx *cbcAEAD) NonceSize() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:72
	_go_fuzz_dep_.CoverTab[184302]++
												return nonceBytes
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:73
	// _ = "end of CoverTab[184302]"
}

func (ctx *cbcAEAD) Overhead() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:76
	_go_fuzz_dep_.CoverTab[184303]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:79
	return ctx.blockCipher.BlockSize() + ctx.authtagBytes
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:79
	// _ = "end of CoverTab[184303]"
}

// Seal encrypts and authenticates the plaintext.
func (ctx *cbcAEAD) Seal(dst, nonce, plaintext, data []byte) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:83
	_go_fuzz_dep_.CoverTab[184304]++

												ciphertext := make([]byte, uint64(len(plaintext))+uint64(ctx.Overhead()))[:len(plaintext)]
												copy(ciphertext, plaintext)
												ciphertext = padBuffer(ciphertext, ctx.blockCipher.BlockSize())

												cbc := cipher.NewCBCEncrypter(ctx.blockCipher, nonce)

												cbc.CryptBlocks(ciphertext, ciphertext)
												authtag := ctx.computeAuthTag(data, nonce, ciphertext)

												ret, out := resize(dst, uint64(len(dst))+uint64(len(ciphertext))+uint64(len(authtag)))
												copy(out, ciphertext)
												copy(out[len(ciphertext):], authtag)

												return ret
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:98
	// _ = "end of CoverTab[184304]"
}

// Open decrypts and authenticates the ciphertext.
func (ctx *cbcAEAD) Open(dst, nonce, ciphertext, data []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:102
	_go_fuzz_dep_.CoverTab[184305]++
												if len(ciphertext) < ctx.authtagBytes {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:103
		_go_fuzz_dep_.CoverTab[184310]++
													return nil, errors.New("square/go-jose: invalid ciphertext (too short)")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:104
		// _ = "end of CoverTab[184310]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:105
		_go_fuzz_dep_.CoverTab[184311]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:105
		// _ = "end of CoverTab[184311]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:105
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:105
	// _ = "end of CoverTab[184305]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:105
	_go_fuzz_dep_.CoverTab[184306]++

												offset := len(ciphertext) - ctx.authtagBytes
												expectedTag := ctx.computeAuthTag(data, nonce, ciphertext[:offset])
												match := subtle.ConstantTimeCompare(expectedTag, ciphertext[offset:])
												if match != 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:110
		_go_fuzz_dep_.CoverTab[184312]++
													return nil, errors.New("square/go-jose: invalid ciphertext (auth tag mismatch)")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:111
		// _ = "end of CoverTab[184312]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:112
		_go_fuzz_dep_.CoverTab[184313]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:112
		// _ = "end of CoverTab[184313]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:112
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:112
	// _ = "end of CoverTab[184306]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:112
	_go_fuzz_dep_.CoverTab[184307]++

												cbc := cipher.NewCBCDecrypter(ctx.blockCipher, nonce)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:117
	buffer := append([]byte{}, []byte(ciphertext[:offset])...)

	if len(buffer)%ctx.blockCipher.BlockSize() > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:119
		_go_fuzz_dep_.CoverTab[184314]++
													return nil, errors.New("square/go-jose: invalid ciphertext (invalid length)")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:120
		// _ = "end of CoverTab[184314]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:121
		_go_fuzz_dep_.CoverTab[184315]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:121
		// _ = "end of CoverTab[184315]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:121
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:121
	// _ = "end of CoverTab[184307]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:121
	_go_fuzz_dep_.CoverTab[184308]++

												cbc.CryptBlocks(buffer, buffer)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:126
	plaintext, err := unpadBuffer(buffer, ctx.blockCipher.BlockSize())
	if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:127
		_go_fuzz_dep_.CoverTab[184316]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:128
		// _ = "end of CoverTab[184316]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:129
		_go_fuzz_dep_.CoverTab[184317]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:129
		// _ = "end of CoverTab[184317]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:129
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:129
	// _ = "end of CoverTab[184308]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:129
	_go_fuzz_dep_.CoverTab[184309]++

												ret, out := resize(dst, uint64(len(dst))+uint64(len(plaintext)))
												copy(out, plaintext)

												return ret, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:134
	// _ = "end of CoverTab[184309]"
}

// Compute an authentication tag
func (ctx *cbcAEAD) computeAuthTag(aad, nonce, ciphertext []byte) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:138
	_go_fuzz_dep_.CoverTab[184318]++
												buffer := make([]byte, uint64(len(aad))+uint64(len(nonce))+uint64(len(ciphertext))+8)
												n := 0
												n += copy(buffer, aad)
												n += copy(buffer[n:], nonce)
												n += copy(buffer[n:], ciphertext)
												binary.BigEndian.PutUint64(buffer[n:], uint64(len(aad))*8)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:147
	hmac := hmac.New(ctx.hash, ctx.integrityKey)
												_, _ = hmac.Write(buffer)

												return hmac.Sum(nil)[:ctx.authtagBytes]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:150
	// _ = "end of CoverTab[184318]"
}

// resize ensures the the given slice has a capacity of at least n bytes.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:153
// If the capacity of the slice is less than n, a new slice is allocated
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:153
// and the existing data will be copied.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:156
func resize(in []byte, n uint64) (head, tail []byte) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:156
	_go_fuzz_dep_.CoverTab[184319]++
												if uint64(cap(in)) >= n {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:157
		_go_fuzz_dep_.CoverTab[184321]++
													head = in[:n]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:158
		// _ = "end of CoverTab[184321]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:159
		_go_fuzz_dep_.CoverTab[184322]++
													head = make([]byte, n)
													copy(head, in)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:161
		// _ = "end of CoverTab[184322]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:162
	// _ = "end of CoverTab[184319]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:162
	_go_fuzz_dep_.CoverTab[184320]++

												tail = head[len(in):]
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:165
	// _ = "end of CoverTab[184320]"
}

// Apply padding
func padBuffer(buffer []byte, blockSize int) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:169
	_go_fuzz_dep_.CoverTab[184323]++
												missing := blockSize - (len(buffer) % blockSize)
												ret, out := resize(buffer, uint64(len(buffer))+uint64(missing))
												padding := bytes.Repeat([]byte{byte(missing)}, missing)
												copy(out, padding)
												return ret
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:174
	// _ = "end of CoverTab[184323]"
}

// Remove padding
func unpadBuffer(buffer []byte, blockSize int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:178
	_go_fuzz_dep_.CoverTab[184324]++
												if len(buffer)%blockSize != 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:179
		_go_fuzz_dep_.CoverTab[184328]++
													return nil, errors.New("square/go-jose: invalid padding")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:180
		// _ = "end of CoverTab[184328]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:181
		_go_fuzz_dep_.CoverTab[184329]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:181
		// _ = "end of CoverTab[184329]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:181
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:181
	// _ = "end of CoverTab[184324]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:181
	_go_fuzz_dep_.CoverTab[184325]++

												last := buffer[len(buffer)-1]
												count := int(last)

												if count == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:186
		_go_fuzz_dep_.CoverTab[184330]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:186
		return count > blockSize
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:186
		// _ = "end of CoverTab[184330]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:186
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:186
		_go_fuzz_dep_.CoverTab[184331]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:186
		return count > len(buffer)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:186
		// _ = "end of CoverTab[184331]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:186
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:186
		_go_fuzz_dep_.CoverTab[184332]++
													return nil, errors.New("square/go-jose: invalid padding")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:187
		// _ = "end of CoverTab[184332]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:188
		_go_fuzz_dep_.CoverTab[184333]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:188
		// _ = "end of CoverTab[184333]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:188
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:188
	// _ = "end of CoverTab[184325]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:188
	_go_fuzz_dep_.CoverTab[184326]++

												padding := bytes.Repeat([]byte{last}, count)
												if !bytes.HasSuffix(buffer, padding) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:191
		_go_fuzz_dep_.CoverTab[184334]++
													return nil, errors.New("square/go-jose: invalid padding")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:192
		// _ = "end of CoverTab[184334]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:193
		_go_fuzz_dep_.CoverTab[184335]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:193
		// _ = "end of CoverTab[184335]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:193
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:193
	// _ = "end of CoverTab[184326]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:193
	_go_fuzz_dep_.CoverTab[184327]++

												return buffer[:len(buffer)-count], nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:195
	// _ = "end of CoverTab[184327]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:196
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/cbc_hmac.go:196
var _ = _go_fuzz_dep_.CoverTab
