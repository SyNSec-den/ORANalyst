//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:17
package josecipher

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:17
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
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:36
	_go_fuzz_dep_.CoverTab[187344]++
												keySize := len(key) / 2
												integrityKey := key[:keySize]
												encryptionKey := key[keySize:]

												blockCipher, err := newBlockCipher(encryptionKey)
												if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:42
		_go_fuzz_dep_.CoverTab[187347]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:43
		// _ = "end of CoverTab[187347]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:44
		_go_fuzz_dep_.CoverTab[187348]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:44
		// _ = "end of CoverTab[187348]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:44
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:44
	// _ = "end of CoverTab[187344]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:44
	_go_fuzz_dep_.CoverTab[187345]++

												var hash func() hash.Hash
												switch keySize {
	case 16:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:48
		_go_fuzz_dep_.CoverTab[187349]++
													hash = sha256.New
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:49
		// _ = "end of CoverTab[187349]"
	case 24:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:50
		_go_fuzz_dep_.CoverTab[187350]++
													hash = sha512.New384
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:51
		// _ = "end of CoverTab[187350]"
	case 32:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:52
		_go_fuzz_dep_.CoverTab[187351]++
													hash = sha512.New
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:53
		// _ = "end of CoverTab[187351]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:53
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:53
		_go_fuzz_dep_.CoverTab[187352]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:53
		// _ = "end of CoverTab[187352]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:54
	// _ = "end of CoverTab[187345]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:54
	_go_fuzz_dep_.CoverTab[187346]++

												return &cbcAEAD{
		hash:		hash,
		blockCipher:	blockCipher,
		authtagBytes:	keySize,
		integrityKey:	integrityKey,
	}, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:61
	// _ = "end of CoverTab[187346]"
}

// An AEAD based on CBC+HMAC
type cbcAEAD struct {
	hash		func() hash.Hash
	authtagBytes	int
	integrityKey	[]byte
	blockCipher	cipher.Block
}

func (ctx *cbcAEAD) NonceSize() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:72
	_go_fuzz_dep_.CoverTab[187353]++
												return nonceBytes
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:73
	// _ = "end of CoverTab[187353]"
}

func (ctx *cbcAEAD) Overhead() int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:76
	_go_fuzz_dep_.CoverTab[187354]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:79
	return ctx.blockCipher.BlockSize() + ctx.authtagBytes
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:79
	// _ = "end of CoverTab[187354]"
}

// Seal encrypts and authenticates the plaintext.
func (ctx *cbcAEAD) Seal(dst, nonce, plaintext, data []byte) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:83
	_go_fuzz_dep_.CoverTab[187355]++

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
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:98
	// _ = "end of CoverTab[187355]"
}

// Open decrypts and authenticates the ciphertext.
func (ctx *cbcAEAD) Open(dst, nonce, ciphertext, data []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:102
	_go_fuzz_dep_.CoverTab[187356]++
												if len(ciphertext) < ctx.authtagBytes {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:103
		_go_fuzz_dep_.CoverTab[187361]++
													return nil, errors.New("square/go-jose: invalid ciphertext (too short)")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:104
		// _ = "end of CoverTab[187361]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:105
		_go_fuzz_dep_.CoverTab[187362]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:105
		// _ = "end of CoverTab[187362]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:105
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:105
	// _ = "end of CoverTab[187356]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:105
	_go_fuzz_dep_.CoverTab[187357]++

												offset := len(ciphertext) - ctx.authtagBytes
												expectedTag := ctx.computeAuthTag(data, nonce, ciphertext[:offset])
												match := subtle.ConstantTimeCompare(expectedTag, ciphertext[offset:])
												if match != 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:110
		_go_fuzz_dep_.CoverTab[187363]++
													return nil, errors.New("square/go-jose: invalid ciphertext (auth tag mismatch)")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:111
		// _ = "end of CoverTab[187363]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:112
		_go_fuzz_dep_.CoverTab[187364]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:112
		// _ = "end of CoverTab[187364]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:112
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:112
	// _ = "end of CoverTab[187357]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:112
	_go_fuzz_dep_.CoverTab[187358]++

												cbc := cipher.NewCBCDecrypter(ctx.blockCipher, nonce)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:117
	buffer := append([]byte{}, []byte(ciphertext[:offset])...)

	if len(buffer)%ctx.blockCipher.BlockSize() > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:119
		_go_fuzz_dep_.CoverTab[187365]++
													return nil, errors.New("square/go-jose: invalid ciphertext (invalid length)")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:120
		// _ = "end of CoverTab[187365]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:121
		_go_fuzz_dep_.CoverTab[187366]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:121
		// _ = "end of CoverTab[187366]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:121
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:121
	// _ = "end of CoverTab[187358]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:121
	_go_fuzz_dep_.CoverTab[187359]++

												cbc.CryptBlocks(buffer, buffer)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:126
	plaintext, err := unpadBuffer(buffer, ctx.blockCipher.BlockSize())
	if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:127
		_go_fuzz_dep_.CoverTab[187367]++
													return nil, err
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:128
		// _ = "end of CoverTab[187367]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:129
		_go_fuzz_dep_.CoverTab[187368]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:129
		// _ = "end of CoverTab[187368]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:129
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:129
	// _ = "end of CoverTab[187359]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:129
	_go_fuzz_dep_.CoverTab[187360]++

												ret, out := resize(dst, uint64(len(dst))+uint64(len(plaintext)))
												copy(out, plaintext)

												return ret, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:134
	// _ = "end of CoverTab[187360]"
}

// Compute an authentication tag
func (ctx *cbcAEAD) computeAuthTag(aad, nonce, ciphertext []byte) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:138
	_go_fuzz_dep_.CoverTab[187369]++
												buffer := make([]byte, uint64(len(aad))+uint64(len(nonce))+uint64(len(ciphertext))+8)
												n := 0
												n += copy(buffer, aad)
												n += copy(buffer[n:], nonce)
												n += copy(buffer[n:], ciphertext)
												binary.BigEndian.PutUint64(buffer[n:], uint64(len(aad))*8)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:147
	hmac := hmac.New(ctx.hash, ctx.integrityKey)
												_, _ = hmac.Write(buffer)

												return hmac.Sum(nil)[:ctx.authtagBytes]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:150
	// _ = "end of CoverTab[187369]"
}

// resize ensures the the given slice has a capacity of at least n bytes.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:153
// If the capacity of the slice is less than n, a new slice is allocated
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:153
// and the existing data will be copied.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:156
func resize(in []byte, n uint64) (head, tail []byte) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:156
	_go_fuzz_dep_.CoverTab[187370]++
												if uint64(cap(in)) >= n {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:157
		_go_fuzz_dep_.CoverTab[187372]++
													head = in[:n]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:158
		// _ = "end of CoverTab[187372]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:159
		_go_fuzz_dep_.CoverTab[187373]++
													head = make([]byte, n)
													copy(head, in)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:161
		// _ = "end of CoverTab[187373]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:162
	// _ = "end of CoverTab[187370]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:162
	_go_fuzz_dep_.CoverTab[187371]++

												tail = head[len(in):]
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:165
	// _ = "end of CoverTab[187371]"
}

// Apply padding
func padBuffer(buffer []byte, blockSize int) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:169
	_go_fuzz_dep_.CoverTab[187374]++
												missing := blockSize - (len(buffer) % blockSize)
												ret, out := resize(buffer, uint64(len(buffer))+uint64(missing))
												padding := bytes.Repeat([]byte{byte(missing)}, missing)
												copy(out, padding)
												return ret
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:174
	// _ = "end of CoverTab[187374]"
}

// Remove padding
func unpadBuffer(buffer []byte, blockSize int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:178
	_go_fuzz_dep_.CoverTab[187375]++
												if len(buffer)%blockSize != 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:179
		_go_fuzz_dep_.CoverTab[187379]++
													return nil, errors.New("square/go-jose: invalid padding")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:180
		// _ = "end of CoverTab[187379]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:181
		_go_fuzz_dep_.CoverTab[187380]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:181
		// _ = "end of CoverTab[187380]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:181
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:181
	// _ = "end of CoverTab[187375]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:181
	_go_fuzz_dep_.CoverTab[187376]++

												last := buffer[len(buffer)-1]
												count := int(last)

												if count == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:186
		_go_fuzz_dep_.CoverTab[187381]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:186
		return count > blockSize
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:186
		// _ = "end of CoverTab[187381]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:186
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:186
		_go_fuzz_dep_.CoverTab[187382]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:186
		return count > len(buffer)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:186
		// _ = "end of CoverTab[187382]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:186
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:186
		_go_fuzz_dep_.CoverTab[187383]++
													return nil, errors.New("square/go-jose: invalid padding")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:187
		// _ = "end of CoverTab[187383]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:188
		_go_fuzz_dep_.CoverTab[187384]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:188
		// _ = "end of CoverTab[187384]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:188
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:188
	// _ = "end of CoverTab[187376]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:188
	_go_fuzz_dep_.CoverTab[187377]++

												padding := bytes.Repeat([]byte{last}, count)
												if !bytes.HasSuffix(buffer, padding) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:191
		_go_fuzz_dep_.CoverTab[187385]++
													return nil, errors.New("square/go-jose: invalid padding")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:192
		// _ = "end of CoverTab[187385]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:193
		_go_fuzz_dep_.CoverTab[187386]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:193
		// _ = "end of CoverTab[187386]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:193
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:193
	// _ = "end of CoverTab[187377]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:193
	_go_fuzz_dep_.CoverTab[187378]++

												return buffer[:len(buffer)-count], nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:195
	// _ = "end of CoverTab[187378]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:196
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/cbc_hmac.go:196
var _ = _go_fuzz_dep_.CoverTab
