//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:1
// Package aescts provides AES CBC CipherText Stealing encryption and decryption methods
package aescts

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:2
)

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

// Encrypt the message with the key and the initial vector.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:11
// Returns: next iv, ciphertext bytes, error
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:13
func Encrypt(key, iv, plaintext []byte) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:13
	_go_fuzz_dep_.CoverTab[85643]++
											l := len(plaintext)

											block, err := aes.NewCipher(key)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:17
		_go_fuzz_dep_.CoverTab[85649]++
												return []byte{}, []byte{}, fmt.Errorf("error creating cipher: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:18
		// _ = "end of CoverTab[85649]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:19
		_go_fuzz_dep_.CoverTab[85650]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:19
		// _ = "end of CoverTab[85650]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:19
	// _ = "end of CoverTab[85643]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:19
	_go_fuzz_dep_.CoverTab[85644]++
											mode := cipher.NewCBCEncrypter(block, iv)

											m := make([]byte, len(plaintext))
											copy(m, plaintext)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:32
	if l <= aes.BlockSize {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:32
		_go_fuzz_dep_.CoverTab[85651]++
												m, _ = zeroPad(m, aes.BlockSize)
												mode.CryptBlocks(m, m)
												return m, m, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:35
		// _ = "end of CoverTab[85651]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:36
		_go_fuzz_dep_.CoverTab[85652]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:36
		// _ = "end of CoverTab[85652]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:36
	// _ = "end of CoverTab[85644]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:36
	_go_fuzz_dep_.CoverTab[85645]++
											if l%aes.BlockSize == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:37
		_go_fuzz_dep_.CoverTab[85653]++
												mode.CryptBlocks(m, m)
												iv = m[len(m)-aes.BlockSize:]
												rb, _ := swapLastTwoBlocks(m, aes.BlockSize)
												return iv, rb, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:41
		// _ = "end of CoverTab[85653]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:42
		_go_fuzz_dep_.CoverTab[85654]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:42
		// _ = "end of CoverTab[85654]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:42
	// _ = "end of CoverTab[85645]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:42
	_go_fuzz_dep_.CoverTab[85646]++
											m, _ = zeroPad(m, aes.BlockSize)
											rb, pb, lb, err := tailBlocks(m, aes.BlockSize)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:45
		_go_fuzz_dep_.CoverTab[85655]++
												return []byte{}, []byte{}, fmt.Errorf("error tailing blocks: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:46
		// _ = "end of CoverTab[85655]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:47
		_go_fuzz_dep_.CoverTab[85656]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:47
		// _ = "end of CoverTab[85656]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:47
	// _ = "end of CoverTab[85646]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:47
	_go_fuzz_dep_.CoverTab[85647]++
											var ct []byte
											if rb != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:49
		_go_fuzz_dep_.CoverTab[85657]++

												mode.CryptBlocks(rb, rb)
												iv = rb[len(rb)-aes.BlockSize:]
												mode = cipher.NewCBCEncrypter(block, iv)
												ct = append(ct, rb...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:54
		// _ = "end of CoverTab[85657]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:55
		_go_fuzz_dep_.CoverTab[85658]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:55
		// _ = "end of CoverTab[85658]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:55
	// _ = "end of CoverTab[85647]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:55
	_go_fuzz_dep_.CoverTab[85648]++
											mode.CryptBlocks(pb, pb)
											mode = cipher.NewCBCEncrypter(block, pb)
											mode.CryptBlocks(lb, lb)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:62
	ct = append(ct, lb...)
											ct = append(ct, pb...)
											return lb, ct[:l], nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:64
	// _ = "end of CoverTab[85648]"
}

// Decrypt the ciphertext with the key and the initial vector.
func Decrypt(key, iv, ciphertext []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:68
	_go_fuzz_dep_.CoverTab[85659]++

											ct := make([]byte, len(ciphertext))
											copy(ct, ciphertext)
											if len(ct) < aes.BlockSize {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:72
		_go_fuzz_dep_.CoverTab[85664]++
												return []byte{}, fmt.Errorf("ciphertext is not large enough. It is less that one block size. Blocksize:%v; Ciphertext:%v", aes.BlockSize, len(ct))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:73
		// _ = "end of CoverTab[85664]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:74
		_go_fuzz_dep_.CoverTab[85665]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:74
		// _ = "end of CoverTab[85665]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:74
	// _ = "end of CoverTab[85659]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:74
	_go_fuzz_dep_.CoverTab[85660]++

											block, err := aes.NewCipher(key)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:77
		_go_fuzz_dep_.CoverTab[85666]++
												return nil, fmt.Errorf("error creating cipher: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:78
		// _ = "end of CoverTab[85666]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:79
		_go_fuzz_dep_.CoverTab[85667]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:79
		// _ = "end of CoverTab[85667]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:79
	// _ = "end of CoverTab[85660]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:79
	_go_fuzz_dep_.CoverTab[85661]++
											var mode cipher.BlockMode

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:84
	if len(ct)%aes.BlockSize == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:84
		_go_fuzz_dep_.CoverTab[85668]++
												if len(ct) > aes.BlockSize {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:85
			_go_fuzz_dep_.CoverTab[85670]++
													ct, _ = swapLastTwoBlocks(ct, aes.BlockSize)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:86
			// _ = "end of CoverTab[85670]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:87
			_go_fuzz_dep_.CoverTab[85671]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:87
			// _ = "end of CoverTab[85671]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:87
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:87
		// _ = "end of CoverTab[85668]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:87
		_go_fuzz_dep_.CoverTab[85669]++
												mode = cipher.NewCBCDecrypter(block, iv)
												message := make([]byte, len(ct))
												mode.CryptBlocks(message, ct)
												return message[:len(ct)], nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:91
		// _ = "end of CoverTab[85669]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:92
		_go_fuzz_dep_.CoverTab[85672]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:92
		// _ = "end of CoverTab[85672]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:92
	// _ = "end of CoverTab[85661]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:92
	_go_fuzz_dep_.CoverTab[85662]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:96
	crb, cpb, clb, _ := tailBlocks(ct, aes.BlockSize)
	v := make([]byte, len(iv), len(iv))
	copy(v, iv)
	var message []byte
	if crb != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:100
		_go_fuzz_dep_.CoverTab[85673]++

												rb := make([]byte, len(crb))
												mode = cipher.NewCBCDecrypter(block, v)
												v = crb[len(crb)-aes.BlockSize:]
												mode.CryptBlocks(rb, crb)
												message = append(message, rb...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:106
		// _ = "end of CoverTab[85673]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:107
		_go_fuzz_dep_.CoverTab[85674]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:107
		// _ = "end of CoverTab[85674]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:107
	// _ = "end of CoverTab[85662]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:107
	_go_fuzz_dep_.CoverTab[85663]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:111
	pb := make([]byte, aes.BlockSize)
											mode = cipher.NewCBCDecrypter(block, iv)
											mode.CryptBlocks(pb, cpb)

											npb := aes.BlockSize - len(ct)%aes.BlockSize

											clb = append(clb, pb[len(pb)-npb:]...)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:121
	lb := make([]byte, aes.BlockSize)
											mode = cipher.NewCBCDecrypter(block, v)
											v = clb
											mode.CryptBlocks(lb, clb)
											message = append(message, lb...)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:128
	mode = cipher.NewCBCDecrypter(block, v)
											mode.CryptBlocks(cpb, cpb)
											message = append(message, cpb...)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:133
	return message[:len(ct)], nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:133
	// _ = "end of CoverTab[85663]"
}

func tailBlocks(b []byte, c int) ([]byte, []byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:136
	_go_fuzz_dep_.CoverTab[85675]++
											if len(b) <= c {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:137
		_go_fuzz_dep_.CoverTab[85679]++
												return []byte{}, []byte{}, []byte{}, errors.New("bytes slice is not larger than one block so cannot tail")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:138
		// _ = "end of CoverTab[85679]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:139
		_go_fuzz_dep_.CoverTab[85680]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:139
		// _ = "end of CoverTab[85680]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:139
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:139
	// _ = "end of CoverTab[85675]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:139
	_go_fuzz_dep_.CoverTab[85676]++
	// Get size of last block
	var lbs int
	if l := len(b) % aes.BlockSize; l == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:142
		_go_fuzz_dep_.CoverTab[85681]++
												lbs = aes.BlockSize
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:143
		// _ = "end of CoverTab[85681]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:144
		_go_fuzz_dep_.CoverTab[85682]++
												lbs = l
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:145
		// _ = "end of CoverTab[85682]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:146
	// _ = "end of CoverTab[85676]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:146
	_go_fuzz_dep_.CoverTab[85677]++

											lb := b[len(b)-lbs:]

											pb := b[len(b)-lbs-c : len(b)-lbs]
											if len(b) > 2*c {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:151
		_go_fuzz_dep_.CoverTab[85683]++
												rb := b[:len(b)-lbs-c]
												return rb, pb, lb, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:153
		// _ = "end of CoverTab[85683]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:154
		_go_fuzz_dep_.CoverTab[85684]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:154
		// _ = "end of CoverTab[85684]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:154
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:154
	// _ = "end of CoverTab[85677]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:154
	_go_fuzz_dep_.CoverTab[85678]++
											return nil, pb, lb, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:155
	// _ = "end of CoverTab[85678]"
}

func swapLastTwoBlocks(b []byte, c int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:158
	_go_fuzz_dep_.CoverTab[85685]++
											rb, pb, lb, err := tailBlocks(b, c)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:160
		_go_fuzz_dep_.CoverTab[85688]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:161
		// _ = "end of CoverTab[85688]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:162
		_go_fuzz_dep_.CoverTab[85689]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:162
		// _ = "end of CoverTab[85689]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:162
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:162
	// _ = "end of CoverTab[85685]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:162
	_go_fuzz_dep_.CoverTab[85686]++
											var out []byte
											if rb != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:164
		_go_fuzz_dep_.CoverTab[85690]++
												out = append(out, rb...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:165
		// _ = "end of CoverTab[85690]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:166
		_go_fuzz_dep_.CoverTab[85691]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:166
		// _ = "end of CoverTab[85691]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:166
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:166
	// _ = "end of CoverTab[85686]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:166
	_go_fuzz_dep_.CoverTab[85687]++
											out = append(out, lb...)
											out = append(out, pb...)
											return out, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:169
	// _ = "end of CoverTab[85687]"
}

// zeroPad pads bytes with zeros to nearest multiple of message size m.
func zeroPad(b []byte, m int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:173
	_go_fuzz_dep_.CoverTab[85692]++
											if m <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:174
		_go_fuzz_dep_.CoverTab[85696]++
												return nil, errors.New("invalid message block size when padding")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:175
		// _ = "end of CoverTab[85696]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:176
		_go_fuzz_dep_.CoverTab[85697]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:176
		// _ = "end of CoverTab[85697]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:176
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:176
	// _ = "end of CoverTab[85692]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:176
	_go_fuzz_dep_.CoverTab[85693]++
											if b == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:177
		_go_fuzz_dep_.CoverTab[85698]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:177
		return len(b) == 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:177
		// _ = "end of CoverTab[85698]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:177
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:177
		_go_fuzz_dep_.CoverTab[85699]++
												return nil, errors.New("data not valid to pad: Zero size")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:178
		// _ = "end of CoverTab[85699]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:179
		_go_fuzz_dep_.CoverTab[85700]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:179
		// _ = "end of CoverTab[85700]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:179
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:179
	// _ = "end of CoverTab[85693]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:179
	_go_fuzz_dep_.CoverTab[85694]++
											if l := len(b) % m; l != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:180
		_go_fuzz_dep_.CoverTab[85701]++
												n := m - l
												z := make([]byte, n)
												b = append(b, z...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:183
		// _ = "end of CoverTab[85701]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:184
		_go_fuzz_dep_.CoverTab[85702]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:184
		// _ = "end of CoverTab[85702]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:184
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:184
	// _ = "end of CoverTab[85694]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:184
	_go_fuzz_dep_.CoverTab[85695]++
											return b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:185
	// _ = "end of CoverTab[85695]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:186
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/aescts/v2@v2.0.0/aescts.go:186
var _ = _go_fuzz_dep_.CoverTab
