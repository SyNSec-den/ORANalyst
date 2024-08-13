// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/x509/pem_decrypt.go:5
package x509

//line /usr/local/go/src/crypto/x509/pem_decrypt.go:5
import (
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:5
)
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:5
import (
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:5
)

//line /usr/local/go/src/crypto/x509/pem_decrypt.go:11
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"io"
	"strings"
)

type PEMCipher int

// Possible values for the EncryptPEMBlock encryption algorithm.
const (
	_	PEMCipher	= iota
	PEMCipherDES
	PEMCipher3DES
	PEMCipherAES128
	PEMCipherAES192
	PEMCipherAES256
)

// rfc1423Algo holds a method for enciphering a PEM block.
type rfc1423Algo struct {
	cipher		PEMCipher
	name		string
	cipherFunc	func(key []byte) (cipher.Block, error)
	keySize		int
	blockSize	int
}

// rfc1423Algos holds a slice of the possible ways to encrypt a PEM
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:44
// block. The ivSize numbers were taken from the OpenSSL source.
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:46
var rfc1423Algos = []rfc1423Algo{{
	cipher:		PEMCipherDES,
	name:		"DES-CBC",
	cipherFunc:	des.NewCipher,
	keySize:	8,
	blockSize:	des.BlockSize,
}, {
	cipher:		PEMCipher3DES,
	name:		"DES-EDE3-CBC",
	cipherFunc:	des.NewTripleDESCipher,
	keySize:	24,
	blockSize:	des.BlockSize,
}, {
	cipher:		PEMCipherAES128,
	name:		"AES-128-CBC",
	cipherFunc:	aes.NewCipher,
	keySize:	16,
	blockSize:	aes.BlockSize,
}, {
	cipher:		PEMCipherAES192,
	name:		"AES-192-CBC",
	cipherFunc:	aes.NewCipher,
	keySize:	24,
	blockSize:	aes.BlockSize,
}, {
	cipher:		PEMCipherAES256,
	name:		"AES-256-CBC",
	cipherFunc:	aes.NewCipher,
	keySize:	32,
	blockSize:	aes.BlockSize,
},
}

// deriveKey uses a key derivation function to stretch the password into a key
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:79
// with the number of bits our cipher requires. This algorithm was derived from
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:79
// the OpenSSL source.
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:82
func (c rfc1423Algo) deriveKey(password, salt []byte) []byte {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:82
	_go_fuzz_dep_.CoverTab[19064]++
							hash := md5.New()
							out := make([]byte, c.keySize)
							var digest []byte

							for i := 0; i < len(out); i += len(digest) {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:87
		_go_fuzz_dep_.CoverTab[19066]++
								hash.Reset()
								hash.Write(digest)
								hash.Write(password)
								hash.Write(salt)
								digest = hash.Sum(digest[:0])
								copy(out[i:], digest)
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:93
		// _ = "end of CoverTab[19066]"
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:94
	// _ = "end of CoverTab[19064]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:94
	_go_fuzz_dep_.CoverTab[19065]++
							return out
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:95
	// _ = "end of CoverTab[19065]"
}

// IsEncryptedPEMBlock returns whether the PEM block is password encrypted
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:98
// according to RFC 1423.
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:98
//
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:98
// Deprecated: Legacy PEM encryption as specified in RFC 1423 is insecure by
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:98
// design. Since it does not authenticate the ciphertext, it is vulnerable to
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:98
// padding oracle attacks that can let an attacker recover the plaintext.
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:104
func IsEncryptedPEMBlock(b *pem.Block) bool {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:104
	_go_fuzz_dep_.CoverTab[19067]++
								_, ok := b.Headers["DEK-Info"]
								return ok
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:106
	// _ = "end of CoverTab[19067]"
}

// IncorrectPasswordError is returned when an incorrect password is detected.
var IncorrectPasswordError = errors.New("x509: decryption password incorrect")

// DecryptPEMBlock takes a PEM block encrypted according to RFC 1423 and the
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:112
// password used to encrypt it and returns a slice of decrypted DER encoded
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:112
// bytes. It inspects the DEK-Info header to determine the algorithm used for
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:112
// decryption. If no DEK-Info header is present, an error is returned. If an
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:112
// incorrect password is detected an IncorrectPasswordError is returned. Because
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:112
// of deficiencies in the format, it's not always possible to detect an
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:112
// incorrect password. In these cases no error will be returned but the
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:112
// decrypted DER bytes will be random noise.
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:112
//
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:112
// Deprecated: Legacy PEM encryption as specified in RFC 1423 is insecure by
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:112
// design. Since it does not authenticate the ciphertext, it is vulnerable to
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:112
// padding oracle attacks that can let an attacker recover the plaintext.
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:124
func DecryptPEMBlock(b *pem.Block, password []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:124
	_go_fuzz_dep_.CoverTab[19068]++
								dek, ok := b.Headers["DEK-Info"]
								if !ok {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:126
		_go_fuzz_dep_.CoverTab[19080]++
									return nil, errors.New("x509: no DEK-Info header in block")
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:127
		// _ = "end of CoverTab[19080]"
	} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:128
		_go_fuzz_dep_.CoverTab[19081]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:128
		// _ = "end of CoverTab[19081]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:128
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:128
	// _ = "end of CoverTab[19068]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:128
	_go_fuzz_dep_.CoverTab[19069]++

								mode, hexIV, ok := strings.Cut(dek, ",")
								if !ok {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:131
		_go_fuzz_dep_.CoverTab[19082]++
									return nil, errors.New("x509: malformed DEK-Info header")
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:132
		// _ = "end of CoverTab[19082]"
	} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:133
		_go_fuzz_dep_.CoverTab[19083]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:133
		// _ = "end of CoverTab[19083]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:133
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:133
	// _ = "end of CoverTab[19069]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:133
	_go_fuzz_dep_.CoverTab[19070]++

								ciph := cipherByName(mode)
								if ciph == nil {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:136
		_go_fuzz_dep_.CoverTab[19084]++
									return nil, errors.New("x509: unknown encryption mode")
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:137
		// _ = "end of CoverTab[19084]"
	} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:138
		_go_fuzz_dep_.CoverTab[19085]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:138
		// _ = "end of CoverTab[19085]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:138
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:138
	// _ = "end of CoverTab[19070]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:138
	_go_fuzz_dep_.CoverTab[19071]++
								iv, err := hex.DecodeString(hexIV)
								if err != nil {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:140
		_go_fuzz_dep_.CoverTab[19086]++
									return nil, err
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:141
		// _ = "end of CoverTab[19086]"
	} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:142
		_go_fuzz_dep_.CoverTab[19087]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:142
		// _ = "end of CoverTab[19087]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:142
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:142
	// _ = "end of CoverTab[19071]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:142
	_go_fuzz_dep_.CoverTab[19072]++
								if len(iv) != ciph.blockSize {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:143
		_go_fuzz_dep_.CoverTab[19088]++
									return nil, errors.New("x509: incorrect IV size")
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:144
		// _ = "end of CoverTab[19088]"
	} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:145
		_go_fuzz_dep_.CoverTab[19089]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:145
		// _ = "end of CoverTab[19089]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:145
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:145
	// _ = "end of CoverTab[19072]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:145
	_go_fuzz_dep_.CoverTab[19073]++

//line /usr/local/go/src/crypto/x509/pem_decrypt.go:149
	key := ciph.deriveKey(password, iv[:8])
	block, err := ciph.cipherFunc(key)
	if err != nil {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:151
		_go_fuzz_dep_.CoverTab[19090]++
									return nil, err
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:152
		// _ = "end of CoverTab[19090]"
	} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:153
		_go_fuzz_dep_.CoverTab[19091]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:153
		// _ = "end of CoverTab[19091]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:153
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:153
	// _ = "end of CoverTab[19073]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:153
	_go_fuzz_dep_.CoverTab[19074]++

								if len(b.Bytes)%block.BlockSize() != 0 {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:155
		_go_fuzz_dep_.CoverTab[19092]++
									return nil, errors.New("x509: encrypted PEM data is not a multiple of the block size")
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:156
		// _ = "end of CoverTab[19092]"
	} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:157
		_go_fuzz_dep_.CoverTab[19093]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:157
		// _ = "end of CoverTab[19093]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:157
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:157
	// _ = "end of CoverTab[19074]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:157
	_go_fuzz_dep_.CoverTab[19075]++

								data := make([]byte, len(b.Bytes))
								dec := cipher.NewCBCDecrypter(block, iv)
								dec.CryptBlocks(data, b.Bytes)

//line /usr/local/go/src/crypto/x509/pem_decrypt.go:169
	dlen := len(data)
	if dlen == 0 || func() bool {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:170
		_go_fuzz_dep_.CoverTab[19094]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:170
		return dlen%ciph.blockSize != 0
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:170
		// _ = "end of CoverTab[19094]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:170
	}() {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:170
		_go_fuzz_dep_.CoverTab[19095]++
									return nil, errors.New("x509: invalid padding")
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:171
		// _ = "end of CoverTab[19095]"
	} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:172
		_go_fuzz_dep_.CoverTab[19096]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:172
		// _ = "end of CoverTab[19096]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:172
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:172
	// _ = "end of CoverTab[19075]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:172
	_go_fuzz_dep_.CoverTab[19076]++
								last := int(data[dlen-1])
								if dlen < last {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:174
		_go_fuzz_dep_.CoverTab[19097]++
									return nil, IncorrectPasswordError
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:175
		// _ = "end of CoverTab[19097]"
	} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:176
		_go_fuzz_dep_.CoverTab[19098]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:176
		// _ = "end of CoverTab[19098]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:176
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:176
	// _ = "end of CoverTab[19076]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:176
	_go_fuzz_dep_.CoverTab[19077]++
								if last == 0 || func() bool {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:177
		_go_fuzz_dep_.CoverTab[19099]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:177
		return last > ciph.blockSize
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:177
		// _ = "end of CoverTab[19099]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:177
	}() {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:177
		_go_fuzz_dep_.CoverTab[19100]++
									return nil, IncorrectPasswordError
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:178
		// _ = "end of CoverTab[19100]"
	} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:179
		_go_fuzz_dep_.CoverTab[19101]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:179
		// _ = "end of CoverTab[19101]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:179
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:179
	// _ = "end of CoverTab[19077]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:179
	_go_fuzz_dep_.CoverTab[19078]++
								for _, val := range data[dlen-last:] {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:180
		_go_fuzz_dep_.CoverTab[19102]++
									if int(val) != last {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:181
			_go_fuzz_dep_.CoverTab[19103]++
										return nil, IncorrectPasswordError
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:182
			// _ = "end of CoverTab[19103]"
		} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:183
			_go_fuzz_dep_.CoverTab[19104]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:183
			// _ = "end of CoverTab[19104]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:183
		}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:183
		// _ = "end of CoverTab[19102]"
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:184
	// _ = "end of CoverTab[19078]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:184
	_go_fuzz_dep_.CoverTab[19079]++
								return data[:dlen-last], nil
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:185
	// _ = "end of CoverTab[19079]"
}

// EncryptPEMBlock returns a PEM block of the specified type holding the
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:188
// given DER encoded data encrypted with the specified algorithm and
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:188
// password according to RFC 1423.
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:188
//
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:188
// Deprecated: Legacy PEM encryption as specified in RFC 1423 is insecure by
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:188
// design. Since it does not authenticate the ciphertext, it is vulnerable to
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:188
// padding oracle attacks that can let an attacker recover the plaintext.
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:195
func EncryptPEMBlock(rand io.Reader, blockType string, data, password []byte, alg PEMCipher) (*pem.Block, error) {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:195
	_go_fuzz_dep_.CoverTab[19105]++
								ciph := cipherByKey(alg)
								if ciph == nil {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:197
		_go_fuzz_dep_.CoverTab[19110]++
									return nil, errors.New("x509: unknown encryption mode")
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:198
		// _ = "end of CoverTab[19110]"
	} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:199
		_go_fuzz_dep_.CoverTab[19111]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:199
		// _ = "end of CoverTab[19111]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:199
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:199
	// _ = "end of CoverTab[19105]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:199
	_go_fuzz_dep_.CoverTab[19106]++
								iv := make([]byte, ciph.blockSize)
								if _, err := io.ReadFull(rand, iv); err != nil {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:201
		_go_fuzz_dep_.CoverTab[19112]++
									return nil, errors.New("x509: cannot generate IV: " + err.Error())
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:202
		// _ = "end of CoverTab[19112]"
	} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:203
		_go_fuzz_dep_.CoverTab[19113]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:203
		// _ = "end of CoverTab[19113]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:203
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:203
	// _ = "end of CoverTab[19106]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:203
	_go_fuzz_dep_.CoverTab[19107]++

//line /usr/local/go/src/crypto/x509/pem_decrypt.go:206
	key := ciph.deriveKey(password, iv[:8])
	block, err := ciph.cipherFunc(key)
	if err != nil {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:208
		_go_fuzz_dep_.CoverTab[19114]++
									return nil, err
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:209
		// _ = "end of CoverTab[19114]"
	} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:210
		_go_fuzz_dep_.CoverTab[19115]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:210
		// _ = "end of CoverTab[19115]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:210
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:210
	// _ = "end of CoverTab[19107]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:210
	_go_fuzz_dep_.CoverTab[19108]++
								enc := cipher.NewCBCEncrypter(block, iv)
								pad := ciph.blockSize - len(data)%ciph.blockSize
								encrypted := make([]byte, len(data), len(data)+pad)

//line /usr/local/go/src/crypto/x509/pem_decrypt.go:217
	copy(encrypted, data)

	for i := 0; i < pad; i++ {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:219
		_go_fuzz_dep_.CoverTab[19116]++
									encrypted = append(encrypted, byte(pad))
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:220
		// _ = "end of CoverTab[19116]"
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:221
	// _ = "end of CoverTab[19108]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:221
	_go_fuzz_dep_.CoverTab[19109]++
								enc.CryptBlocks(encrypted, encrypted)

								return &pem.Block{
		Type:	blockType,
		Headers: map[string]string{
			"Proc-Type":	"4,ENCRYPTED",
			"DEK-Info":	ciph.name + "," + hex.EncodeToString(iv),
		},
		Bytes:	encrypted,
	}, nil
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:231
	// _ = "end of CoverTab[19109]"
}

func cipherByName(name string) *rfc1423Algo {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:234
	_go_fuzz_dep_.CoverTab[19117]++
								for i := range rfc1423Algos {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:235
		_go_fuzz_dep_.CoverTab[19119]++
									alg := &rfc1423Algos[i]
									if alg.name == name {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:237
			_go_fuzz_dep_.CoverTab[19120]++
										return alg
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:238
			// _ = "end of CoverTab[19120]"
		} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:239
			_go_fuzz_dep_.CoverTab[19121]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:239
			// _ = "end of CoverTab[19121]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:239
		}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:239
		// _ = "end of CoverTab[19119]"
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:240
	// _ = "end of CoverTab[19117]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:240
	_go_fuzz_dep_.CoverTab[19118]++
								return nil
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:241
	// _ = "end of CoverTab[19118]"
}

func cipherByKey(key PEMCipher) *rfc1423Algo {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:244
	_go_fuzz_dep_.CoverTab[19122]++
								for i := range rfc1423Algos {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:245
		_go_fuzz_dep_.CoverTab[19124]++
									alg := &rfc1423Algos[i]
									if alg.cipher == key {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:247
			_go_fuzz_dep_.CoverTab[19125]++
										return alg
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:248
			// _ = "end of CoverTab[19125]"
		} else {
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:249
			_go_fuzz_dep_.CoverTab[19126]++
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:249
			// _ = "end of CoverTab[19126]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:249
		}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:249
		// _ = "end of CoverTab[19124]"
	}
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:250
	// _ = "end of CoverTab[19122]"
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:250
	_go_fuzz_dep_.CoverTab[19123]++
								return nil
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:251
	// _ = "end of CoverTab[19123]"
}

//line /usr/local/go/src/crypto/x509/pem_decrypt.go:252
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/x509/pem_decrypt.go:252
var _ = _go_fuzz_dep_.CoverTab
