// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/rsa/pss.go:5
package rsa

//line /usr/local/go/src/crypto/rsa/pss.go:5
import (
//line /usr/local/go/src/crypto/rsa/pss.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/rsa/pss.go:5
)
//line /usr/local/go/src/crypto/rsa/pss.go:5
import (
//line /usr/local/go/src/crypto/rsa/pss.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/rsa/pss.go:5
)

//line /usr/local/go/src/crypto/rsa/pss.go:9
import (
	"bytes"
	"crypto"
	"crypto/internal/boring"
	"errors"
	"hash"
	"io"
)

//line /usr/local/go/src/crypto/rsa/pss.go:31
func emsaPSSEncode(mHash []byte, emBits int, salt []byte, hash hash.Hash) ([]byte, error) {
//line /usr/local/go/src/crypto/rsa/pss.go:31
	_go_fuzz_dep_.CoverTab[9752]++

//line /usr/local/go/src/crypto/rsa/pss.go:34
	hLen := hash.Size()
						sLen := len(salt)
						emLen := (emBits + 7) / 8

//line /usr/local/go/src/crypto/rsa/pss.go:44
	if len(mHash) != hLen {
//line /usr/local/go/src/crypto/rsa/pss.go:44
		_go_fuzz_dep_.CoverTab[9755]++
							return nil, errors.New("crypto/rsa: input must be hashed with given hash")
//line /usr/local/go/src/crypto/rsa/pss.go:45
		// _ = "end of CoverTab[9755]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:46
		_go_fuzz_dep_.CoverTab[9756]++
//line /usr/local/go/src/crypto/rsa/pss.go:46
		// _ = "end of CoverTab[9756]"
//line /usr/local/go/src/crypto/rsa/pss.go:46
	}
//line /usr/local/go/src/crypto/rsa/pss.go:46
	// _ = "end of CoverTab[9752]"
//line /usr/local/go/src/crypto/rsa/pss.go:46
	_go_fuzz_dep_.CoverTab[9753]++

//line /usr/local/go/src/crypto/rsa/pss.go:50
	if emLen < hLen+sLen+2 {
//line /usr/local/go/src/crypto/rsa/pss.go:50
		_go_fuzz_dep_.CoverTab[9757]++
							return nil, ErrMessageTooLong
//line /usr/local/go/src/crypto/rsa/pss.go:51
		// _ = "end of CoverTab[9757]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:52
		_go_fuzz_dep_.CoverTab[9758]++
//line /usr/local/go/src/crypto/rsa/pss.go:52
		// _ = "end of CoverTab[9758]"
//line /usr/local/go/src/crypto/rsa/pss.go:52
	}
//line /usr/local/go/src/crypto/rsa/pss.go:52
	// _ = "end of CoverTab[9753]"
//line /usr/local/go/src/crypto/rsa/pss.go:52
	_go_fuzz_dep_.CoverTab[9754]++

						em := make([]byte, emLen)
						psLen := emLen - sLen - hLen - 2
						db := em[:psLen+1+sLen]
						h := em[psLen+1+sLen : emLen-1]

//line /usr/local/go/src/crypto/rsa/pss.go:70
	var prefix [8]byte

						hash.Write(prefix[:])
						hash.Write(mHash)
						hash.Write(salt)

						h = hash.Sum(h[:0])
						hash.Reset()

//line /usr/local/go/src/crypto/rsa/pss.go:85
	db[psLen] = 0x01
						copy(db[psLen+1:], salt)

//line /usr/local/go/src/crypto/rsa/pss.go:92
	mgf1XOR(db, hash, h)

//line /usr/local/go/src/crypto/rsa/pss.go:97
	db[0] &= 0xff >> (8*emLen - emBits)

//line /usr/local/go/src/crypto/rsa/pss.go:100
	em[emLen-1] = 0xbc

//line /usr/local/go/src/crypto/rsa/pss.go:103
	return em, nil
//line /usr/local/go/src/crypto/rsa/pss.go:103
	// _ = "end of CoverTab[9754]"
}

func emsaPSSVerify(mHash, em []byte, emBits, sLen int, hash hash.Hash) error {
//line /usr/local/go/src/crypto/rsa/pss.go:106
	_go_fuzz_dep_.CoverTab[9759]++

//line /usr/local/go/src/crypto/rsa/pss.go:109
	hLen := hash.Size()
	if sLen == PSSSaltLengthEqualsHash {
//line /usr/local/go/src/crypto/rsa/pss.go:110
		_go_fuzz_dep_.CoverTab[9770]++
							sLen = hLen
//line /usr/local/go/src/crypto/rsa/pss.go:111
		// _ = "end of CoverTab[9770]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:112
		_go_fuzz_dep_.CoverTab[9771]++
//line /usr/local/go/src/crypto/rsa/pss.go:112
		// _ = "end of CoverTab[9771]"
//line /usr/local/go/src/crypto/rsa/pss.go:112
	}
//line /usr/local/go/src/crypto/rsa/pss.go:112
	// _ = "end of CoverTab[9759]"
//line /usr/local/go/src/crypto/rsa/pss.go:112
	_go_fuzz_dep_.CoverTab[9760]++
						emLen := (emBits + 7) / 8
						if emLen != len(em) {
//line /usr/local/go/src/crypto/rsa/pss.go:114
		_go_fuzz_dep_.CoverTab[9772]++
							return errors.New("rsa: internal error: inconsistent length")
//line /usr/local/go/src/crypto/rsa/pss.go:115
		// _ = "end of CoverTab[9772]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:116
		_go_fuzz_dep_.CoverTab[9773]++
//line /usr/local/go/src/crypto/rsa/pss.go:116
		// _ = "end of CoverTab[9773]"
//line /usr/local/go/src/crypto/rsa/pss.go:116
	}
//line /usr/local/go/src/crypto/rsa/pss.go:116
	// _ = "end of CoverTab[9760]"
//line /usr/local/go/src/crypto/rsa/pss.go:116
	_go_fuzz_dep_.CoverTab[9761]++

//line /usr/local/go/src/crypto/rsa/pss.go:123
	if hLen != len(mHash) {
//line /usr/local/go/src/crypto/rsa/pss.go:123
		_go_fuzz_dep_.CoverTab[9774]++
							return ErrVerification
//line /usr/local/go/src/crypto/rsa/pss.go:124
		// _ = "end of CoverTab[9774]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:125
		_go_fuzz_dep_.CoverTab[9775]++
//line /usr/local/go/src/crypto/rsa/pss.go:125
		// _ = "end of CoverTab[9775]"
//line /usr/local/go/src/crypto/rsa/pss.go:125
	}
//line /usr/local/go/src/crypto/rsa/pss.go:125
	// _ = "end of CoverTab[9761]"
//line /usr/local/go/src/crypto/rsa/pss.go:125
	_go_fuzz_dep_.CoverTab[9762]++

//line /usr/local/go/src/crypto/rsa/pss.go:128
	if emLen < hLen+sLen+2 {
//line /usr/local/go/src/crypto/rsa/pss.go:128
		_go_fuzz_dep_.CoverTab[9776]++
							return ErrVerification
//line /usr/local/go/src/crypto/rsa/pss.go:129
		// _ = "end of CoverTab[9776]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:130
		_go_fuzz_dep_.CoverTab[9777]++
//line /usr/local/go/src/crypto/rsa/pss.go:130
		// _ = "end of CoverTab[9777]"
//line /usr/local/go/src/crypto/rsa/pss.go:130
	}
//line /usr/local/go/src/crypto/rsa/pss.go:130
	// _ = "end of CoverTab[9762]"
//line /usr/local/go/src/crypto/rsa/pss.go:130
	_go_fuzz_dep_.CoverTab[9763]++

//line /usr/local/go/src/crypto/rsa/pss.go:134
	if em[emLen-1] != 0xbc {
//line /usr/local/go/src/crypto/rsa/pss.go:134
		_go_fuzz_dep_.CoverTab[9778]++
							return ErrVerification
//line /usr/local/go/src/crypto/rsa/pss.go:135
		// _ = "end of CoverTab[9778]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:136
		_go_fuzz_dep_.CoverTab[9779]++
//line /usr/local/go/src/crypto/rsa/pss.go:136
		// _ = "end of CoverTab[9779]"
//line /usr/local/go/src/crypto/rsa/pss.go:136
	}
//line /usr/local/go/src/crypto/rsa/pss.go:136
	// _ = "end of CoverTab[9763]"
//line /usr/local/go/src/crypto/rsa/pss.go:136
	_go_fuzz_dep_.CoverTab[9764]++

//line /usr/local/go/src/crypto/rsa/pss.go:140
	db := em[:emLen-hLen-1]
	h := em[emLen-hLen-1 : emLen-1]

	// 6.  If the leftmost 8 * emLen - emBits bits of the leftmost octet in
	//     maskedDB are not all equal to zero, output "inconsistent" and
	//     stop.
	var bitMask byte = 0xff >> (8*emLen - emBits)
	if em[0] & ^bitMask != 0 {
//line /usr/local/go/src/crypto/rsa/pss.go:147
		_go_fuzz_dep_.CoverTab[9780]++
							return ErrVerification
//line /usr/local/go/src/crypto/rsa/pss.go:148
		// _ = "end of CoverTab[9780]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:149
		_go_fuzz_dep_.CoverTab[9781]++
//line /usr/local/go/src/crypto/rsa/pss.go:149
		// _ = "end of CoverTab[9781]"
//line /usr/local/go/src/crypto/rsa/pss.go:149
	}
//line /usr/local/go/src/crypto/rsa/pss.go:149
	// _ = "end of CoverTab[9764]"
//line /usr/local/go/src/crypto/rsa/pss.go:149
	_go_fuzz_dep_.CoverTab[9765]++

//line /usr/local/go/src/crypto/rsa/pss.go:154
	mgf1XOR(db, hash, h)

//line /usr/local/go/src/crypto/rsa/pss.go:158
	db[0] &= bitMask

//line /usr/local/go/src/crypto/rsa/pss.go:161
	if sLen == PSSSaltLengthAuto {
//line /usr/local/go/src/crypto/rsa/pss.go:161
		_go_fuzz_dep_.CoverTab[9782]++
							psLen := bytes.IndexByte(db, 0x01)
							if psLen < 0 {
//line /usr/local/go/src/crypto/rsa/pss.go:163
			_go_fuzz_dep_.CoverTab[9784]++
								return ErrVerification
//line /usr/local/go/src/crypto/rsa/pss.go:164
			// _ = "end of CoverTab[9784]"
		} else {
//line /usr/local/go/src/crypto/rsa/pss.go:165
			_go_fuzz_dep_.CoverTab[9785]++
//line /usr/local/go/src/crypto/rsa/pss.go:165
			// _ = "end of CoverTab[9785]"
//line /usr/local/go/src/crypto/rsa/pss.go:165
		}
//line /usr/local/go/src/crypto/rsa/pss.go:165
		// _ = "end of CoverTab[9782]"
//line /usr/local/go/src/crypto/rsa/pss.go:165
		_go_fuzz_dep_.CoverTab[9783]++
							sLen = len(db) - psLen - 1
//line /usr/local/go/src/crypto/rsa/pss.go:166
		// _ = "end of CoverTab[9783]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:167
		_go_fuzz_dep_.CoverTab[9786]++
//line /usr/local/go/src/crypto/rsa/pss.go:167
		// _ = "end of CoverTab[9786]"
//line /usr/local/go/src/crypto/rsa/pss.go:167
	}
//line /usr/local/go/src/crypto/rsa/pss.go:167
	// _ = "end of CoverTab[9765]"
//line /usr/local/go/src/crypto/rsa/pss.go:167
	_go_fuzz_dep_.CoverTab[9766]++

//line /usr/local/go/src/crypto/rsa/pss.go:173
	psLen := emLen - hLen - sLen - 2
	for _, e := range db[:psLen] {
//line /usr/local/go/src/crypto/rsa/pss.go:174
		_go_fuzz_dep_.CoverTab[9787]++
							if e != 0x00 {
//line /usr/local/go/src/crypto/rsa/pss.go:175
			_go_fuzz_dep_.CoverTab[9788]++
								return ErrVerification
//line /usr/local/go/src/crypto/rsa/pss.go:176
			// _ = "end of CoverTab[9788]"
		} else {
//line /usr/local/go/src/crypto/rsa/pss.go:177
			_go_fuzz_dep_.CoverTab[9789]++
//line /usr/local/go/src/crypto/rsa/pss.go:177
			// _ = "end of CoverTab[9789]"
//line /usr/local/go/src/crypto/rsa/pss.go:177
		}
//line /usr/local/go/src/crypto/rsa/pss.go:177
		// _ = "end of CoverTab[9787]"
	}
//line /usr/local/go/src/crypto/rsa/pss.go:178
	// _ = "end of CoverTab[9766]"
//line /usr/local/go/src/crypto/rsa/pss.go:178
	_go_fuzz_dep_.CoverTab[9767]++
						if db[psLen] != 0x01 {
//line /usr/local/go/src/crypto/rsa/pss.go:179
		_go_fuzz_dep_.CoverTab[9790]++
							return ErrVerification
//line /usr/local/go/src/crypto/rsa/pss.go:180
		// _ = "end of CoverTab[9790]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:181
		_go_fuzz_dep_.CoverTab[9791]++
//line /usr/local/go/src/crypto/rsa/pss.go:181
		// _ = "end of CoverTab[9791]"
//line /usr/local/go/src/crypto/rsa/pss.go:181
	}
//line /usr/local/go/src/crypto/rsa/pss.go:181
	// _ = "end of CoverTab[9767]"
//line /usr/local/go/src/crypto/rsa/pss.go:181
	_go_fuzz_dep_.CoverTab[9768]++

//line /usr/local/go/src/crypto/rsa/pss.go:184
	salt := db[len(db)-sLen:]

	// 12.  Let
	//          M' = (0x)00 00 00 00 00 00 00 00 || mHash || salt ;
	//     M' is an octet string of length 8 + hLen + sLen with eight
	//     initial zero octets.
	//
						// 13. Let H' = Hash(M'), an octet string of length hLen.
						var prefix [8]byte
						hash.Write(prefix[:])
						hash.Write(mHash)
						hash.Write(salt)

						h0 := hash.Sum(nil)

//line /usr/local/go/src/crypto/rsa/pss.go:200
	if !bytes.Equal(h0, h) {
//line /usr/local/go/src/crypto/rsa/pss.go:200
		_go_fuzz_dep_.CoverTab[9792]++
							return ErrVerification
//line /usr/local/go/src/crypto/rsa/pss.go:201
		// _ = "end of CoverTab[9792]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:202
		_go_fuzz_dep_.CoverTab[9793]++
//line /usr/local/go/src/crypto/rsa/pss.go:202
		// _ = "end of CoverTab[9793]"
//line /usr/local/go/src/crypto/rsa/pss.go:202
	}
//line /usr/local/go/src/crypto/rsa/pss.go:202
	// _ = "end of CoverTab[9768]"
//line /usr/local/go/src/crypto/rsa/pss.go:202
	_go_fuzz_dep_.CoverTab[9769]++
						return nil
//line /usr/local/go/src/crypto/rsa/pss.go:203
	// _ = "end of CoverTab[9769]"
}

// signPSSWithSalt calculates the signature of hashed using PSS with specified salt.
//line /usr/local/go/src/crypto/rsa/pss.go:206
// Note that hashed must be the result of hashing the input message using the
//line /usr/local/go/src/crypto/rsa/pss.go:206
// given hash function. salt is a random sequence of bytes whose length will be
//line /usr/local/go/src/crypto/rsa/pss.go:206
// later used to verify the signature.
//line /usr/local/go/src/crypto/rsa/pss.go:210
func signPSSWithSalt(priv *PrivateKey, hash crypto.Hash, hashed, salt []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/rsa/pss.go:210
	_go_fuzz_dep_.CoverTab[9794]++
						emBits := priv.N.BitLen() - 1
						em, err := emsaPSSEncode(hashed, emBits, salt, hash.New())
						if err != nil {
//line /usr/local/go/src/crypto/rsa/pss.go:213
		_go_fuzz_dep_.CoverTab[9798]++
							return nil, err
//line /usr/local/go/src/crypto/rsa/pss.go:214
		// _ = "end of CoverTab[9798]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:215
		_go_fuzz_dep_.CoverTab[9799]++
//line /usr/local/go/src/crypto/rsa/pss.go:215
		// _ = "end of CoverTab[9799]"
//line /usr/local/go/src/crypto/rsa/pss.go:215
	}
//line /usr/local/go/src/crypto/rsa/pss.go:215
	// _ = "end of CoverTab[9794]"
//line /usr/local/go/src/crypto/rsa/pss.go:215
	_go_fuzz_dep_.CoverTab[9795]++

						if boring.Enabled {
//line /usr/local/go/src/crypto/rsa/pss.go:217
		_go_fuzz_dep_.CoverTab[9800]++
							bkey, err := boringPrivateKey(priv)
							if err != nil {
//line /usr/local/go/src/crypto/rsa/pss.go:219
			_go_fuzz_dep_.CoverTab[9803]++
								return nil, err
//line /usr/local/go/src/crypto/rsa/pss.go:220
			// _ = "end of CoverTab[9803]"
		} else {
//line /usr/local/go/src/crypto/rsa/pss.go:221
			_go_fuzz_dep_.CoverTab[9804]++
//line /usr/local/go/src/crypto/rsa/pss.go:221
			// _ = "end of CoverTab[9804]"
//line /usr/local/go/src/crypto/rsa/pss.go:221
		}
//line /usr/local/go/src/crypto/rsa/pss.go:221
		// _ = "end of CoverTab[9800]"
//line /usr/local/go/src/crypto/rsa/pss.go:221
		_go_fuzz_dep_.CoverTab[9801]++

//line /usr/local/go/src/crypto/rsa/pss.go:224
		s, err := boring.DecryptRSANoPadding(bkey, em)
		if err != nil {
//line /usr/local/go/src/crypto/rsa/pss.go:225
			_go_fuzz_dep_.CoverTab[9805]++
								return nil, err
//line /usr/local/go/src/crypto/rsa/pss.go:226
			// _ = "end of CoverTab[9805]"
		} else {
//line /usr/local/go/src/crypto/rsa/pss.go:227
			_go_fuzz_dep_.CoverTab[9806]++
//line /usr/local/go/src/crypto/rsa/pss.go:227
			// _ = "end of CoverTab[9806]"
//line /usr/local/go/src/crypto/rsa/pss.go:227
		}
//line /usr/local/go/src/crypto/rsa/pss.go:227
		// _ = "end of CoverTab[9801]"
//line /usr/local/go/src/crypto/rsa/pss.go:227
		_go_fuzz_dep_.CoverTab[9802]++
							return s, nil
//line /usr/local/go/src/crypto/rsa/pss.go:228
		// _ = "end of CoverTab[9802]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:229
		_go_fuzz_dep_.CoverTab[9807]++
//line /usr/local/go/src/crypto/rsa/pss.go:229
		// _ = "end of CoverTab[9807]"
//line /usr/local/go/src/crypto/rsa/pss.go:229
	}
//line /usr/local/go/src/crypto/rsa/pss.go:229
	// _ = "end of CoverTab[9795]"
//line /usr/local/go/src/crypto/rsa/pss.go:229
	_go_fuzz_dep_.CoverTab[9796]++

//line /usr/local/go/src/crypto/rsa/pss.go:238
	if emLen, k := len(em), priv.Size(); emLen < k {
//line /usr/local/go/src/crypto/rsa/pss.go:238
		_go_fuzz_dep_.CoverTab[9808]++
							emNew := make([]byte, k)
							copy(emNew[k-emLen:], em)
							em = emNew
//line /usr/local/go/src/crypto/rsa/pss.go:241
		// _ = "end of CoverTab[9808]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:242
		_go_fuzz_dep_.CoverTab[9809]++
//line /usr/local/go/src/crypto/rsa/pss.go:242
		// _ = "end of CoverTab[9809]"
//line /usr/local/go/src/crypto/rsa/pss.go:242
	}
//line /usr/local/go/src/crypto/rsa/pss.go:242
	// _ = "end of CoverTab[9796]"
//line /usr/local/go/src/crypto/rsa/pss.go:242
	_go_fuzz_dep_.CoverTab[9797]++

						return decrypt(priv, em, withCheck)
//line /usr/local/go/src/crypto/rsa/pss.go:244
	// _ = "end of CoverTab[9797]"
}

const (
	// PSSSaltLengthAuto causes the salt in a PSS signature to be as large
	// as possible when signing, and to be auto-detected when verifying.
	PSSSaltLengthAuto	= 0
	// PSSSaltLengthEqualsHash causes the salt length to equal the length
	// of the hash used in the signature.
	PSSSaltLengthEqualsHash	= -1
)

// PSSOptions contains options for creating and verifying PSS signatures.
type PSSOptions struct {
	// SaltLength controls the length of the salt used in the PSS signature. It
	// can either be a positive number of bytes, or one of the special
	// PSSSaltLength constants.
	SaltLength	int

	// Hash is the hash function used to generate the message digest. If not
	// zero, it overrides the hash function passed to SignPSS. It's required
	// when using PrivateKey.Sign.
	Hash	crypto.Hash
}

// HashFunc returns opts.Hash so that PSSOptions implements crypto.SignerOpts.
func (opts *PSSOptions) HashFunc() crypto.Hash {
//line /usr/local/go/src/crypto/rsa/pss.go:270
	_go_fuzz_dep_.CoverTab[9810]++
						return opts.Hash
//line /usr/local/go/src/crypto/rsa/pss.go:271
	// _ = "end of CoverTab[9810]"
}

func (opts *PSSOptions) saltLength() int {
//line /usr/local/go/src/crypto/rsa/pss.go:274
	_go_fuzz_dep_.CoverTab[9811]++
						if opts == nil {
//line /usr/local/go/src/crypto/rsa/pss.go:275
		_go_fuzz_dep_.CoverTab[9813]++
							return PSSSaltLengthAuto
//line /usr/local/go/src/crypto/rsa/pss.go:276
		// _ = "end of CoverTab[9813]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:277
		_go_fuzz_dep_.CoverTab[9814]++
//line /usr/local/go/src/crypto/rsa/pss.go:277
		// _ = "end of CoverTab[9814]"
//line /usr/local/go/src/crypto/rsa/pss.go:277
	}
//line /usr/local/go/src/crypto/rsa/pss.go:277
	// _ = "end of CoverTab[9811]"
//line /usr/local/go/src/crypto/rsa/pss.go:277
	_go_fuzz_dep_.CoverTab[9812]++
						return opts.SaltLength
//line /usr/local/go/src/crypto/rsa/pss.go:278
	// _ = "end of CoverTab[9812]"
}

var invalidSaltLenErr = errors.New("crypto/rsa: PSSOptions.SaltLength cannot be negative")

// SignPSS calculates the signature of digest using PSS.
//line /usr/local/go/src/crypto/rsa/pss.go:283
//
//line /usr/local/go/src/crypto/rsa/pss.go:283
// digest must be the result of hashing the input message using the given hash
//line /usr/local/go/src/crypto/rsa/pss.go:283
// function. The opts argument may be nil, in which case sensible defaults are
//line /usr/local/go/src/crypto/rsa/pss.go:283
// used. If opts.Hash is set, it overrides hash.
//line /usr/local/go/src/crypto/rsa/pss.go:288
func SignPSS(rand io.Reader, priv *PrivateKey, hash crypto.Hash, digest []byte, opts *PSSOptions) ([]byte, error) {
//line /usr/local/go/src/crypto/rsa/pss.go:288
	_go_fuzz_dep_.CoverTab[9815]++
						if boring.Enabled && func() bool {
//line /usr/local/go/src/crypto/rsa/pss.go:289
		_go_fuzz_dep_.CoverTab[9820]++
//line /usr/local/go/src/crypto/rsa/pss.go:289
		return rand == boring.RandReader
//line /usr/local/go/src/crypto/rsa/pss.go:289
		// _ = "end of CoverTab[9820]"
//line /usr/local/go/src/crypto/rsa/pss.go:289
	}() {
//line /usr/local/go/src/crypto/rsa/pss.go:289
		_go_fuzz_dep_.CoverTab[9821]++
							bkey, err := boringPrivateKey(priv)
							if err != nil {
//line /usr/local/go/src/crypto/rsa/pss.go:291
			_go_fuzz_dep_.CoverTab[9823]++
								return nil, err
//line /usr/local/go/src/crypto/rsa/pss.go:292
			// _ = "end of CoverTab[9823]"
		} else {
//line /usr/local/go/src/crypto/rsa/pss.go:293
			_go_fuzz_dep_.CoverTab[9824]++
//line /usr/local/go/src/crypto/rsa/pss.go:293
			// _ = "end of CoverTab[9824]"
//line /usr/local/go/src/crypto/rsa/pss.go:293
		}
//line /usr/local/go/src/crypto/rsa/pss.go:293
		// _ = "end of CoverTab[9821]"
//line /usr/local/go/src/crypto/rsa/pss.go:293
		_go_fuzz_dep_.CoverTab[9822]++
							return boring.SignRSAPSS(bkey, hash, digest, opts.saltLength())
//line /usr/local/go/src/crypto/rsa/pss.go:294
		// _ = "end of CoverTab[9822]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:295
		_go_fuzz_dep_.CoverTab[9825]++
//line /usr/local/go/src/crypto/rsa/pss.go:295
		// _ = "end of CoverTab[9825]"
//line /usr/local/go/src/crypto/rsa/pss.go:295
	}
//line /usr/local/go/src/crypto/rsa/pss.go:295
	// _ = "end of CoverTab[9815]"
//line /usr/local/go/src/crypto/rsa/pss.go:295
	_go_fuzz_dep_.CoverTab[9816]++
						boring.UnreachableExceptTests()

						if opts != nil && func() bool {
//line /usr/local/go/src/crypto/rsa/pss.go:298
		_go_fuzz_dep_.CoverTab[9826]++
//line /usr/local/go/src/crypto/rsa/pss.go:298
		return opts.Hash != 0
//line /usr/local/go/src/crypto/rsa/pss.go:298
		// _ = "end of CoverTab[9826]"
//line /usr/local/go/src/crypto/rsa/pss.go:298
	}() {
//line /usr/local/go/src/crypto/rsa/pss.go:298
		_go_fuzz_dep_.CoverTab[9827]++
							hash = opts.Hash
//line /usr/local/go/src/crypto/rsa/pss.go:299
		// _ = "end of CoverTab[9827]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:300
		_go_fuzz_dep_.CoverTab[9828]++
//line /usr/local/go/src/crypto/rsa/pss.go:300
		// _ = "end of CoverTab[9828]"
//line /usr/local/go/src/crypto/rsa/pss.go:300
	}
//line /usr/local/go/src/crypto/rsa/pss.go:300
	// _ = "end of CoverTab[9816]"
//line /usr/local/go/src/crypto/rsa/pss.go:300
	_go_fuzz_dep_.CoverTab[9817]++

						saltLength := opts.saltLength()
						switch saltLength {
	case PSSSaltLengthAuto:
//line /usr/local/go/src/crypto/rsa/pss.go:304
		_go_fuzz_dep_.CoverTab[9829]++
							saltLength = (priv.N.BitLen()-1+7)/8 - 2 - hash.Size()
							if saltLength < 0 {
//line /usr/local/go/src/crypto/rsa/pss.go:306
			_go_fuzz_dep_.CoverTab[9832]++
								return nil, ErrMessageTooLong
//line /usr/local/go/src/crypto/rsa/pss.go:307
			// _ = "end of CoverTab[9832]"
		} else {
//line /usr/local/go/src/crypto/rsa/pss.go:308
			_go_fuzz_dep_.CoverTab[9833]++
//line /usr/local/go/src/crypto/rsa/pss.go:308
			// _ = "end of CoverTab[9833]"
//line /usr/local/go/src/crypto/rsa/pss.go:308
		}
//line /usr/local/go/src/crypto/rsa/pss.go:308
		// _ = "end of CoverTab[9829]"
	case PSSSaltLengthEqualsHash:
//line /usr/local/go/src/crypto/rsa/pss.go:309
		_go_fuzz_dep_.CoverTab[9830]++
							saltLength = hash.Size()
//line /usr/local/go/src/crypto/rsa/pss.go:310
		// _ = "end of CoverTab[9830]"
	default:
//line /usr/local/go/src/crypto/rsa/pss.go:311
		_go_fuzz_dep_.CoverTab[9831]++

//line /usr/local/go/src/crypto/rsa/pss.go:314
		if saltLength <= 0 {
//line /usr/local/go/src/crypto/rsa/pss.go:314
			_go_fuzz_dep_.CoverTab[9834]++
								return nil, invalidSaltLenErr
//line /usr/local/go/src/crypto/rsa/pss.go:315
			// _ = "end of CoverTab[9834]"
		} else {
//line /usr/local/go/src/crypto/rsa/pss.go:316
			_go_fuzz_dep_.CoverTab[9835]++
//line /usr/local/go/src/crypto/rsa/pss.go:316
			// _ = "end of CoverTab[9835]"
//line /usr/local/go/src/crypto/rsa/pss.go:316
		}
//line /usr/local/go/src/crypto/rsa/pss.go:316
		// _ = "end of CoverTab[9831]"
	}
//line /usr/local/go/src/crypto/rsa/pss.go:317
	// _ = "end of CoverTab[9817]"
//line /usr/local/go/src/crypto/rsa/pss.go:317
	_go_fuzz_dep_.CoverTab[9818]++
						salt := make([]byte, saltLength)
						if _, err := io.ReadFull(rand, salt); err != nil {
//line /usr/local/go/src/crypto/rsa/pss.go:319
		_go_fuzz_dep_.CoverTab[9836]++
							return nil, err
//line /usr/local/go/src/crypto/rsa/pss.go:320
		// _ = "end of CoverTab[9836]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:321
		_go_fuzz_dep_.CoverTab[9837]++
//line /usr/local/go/src/crypto/rsa/pss.go:321
		// _ = "end of CoverTab[9837]"
//line /usr/local/go/src/crypto/rsa/pss.go:321
	}
//line /usr/local/go/src/crypto/rsa/pss.go:321
	// _ = "end of CoverTab[9818]"
//line /usr/local/go/src/crypto/rsa/pss.go:321
	_go_fuzz_dep_.CoverTab[9819]++
						return signPSSWithSalt(priv, hash, digest, salt)
//line /usr/local/go/src/crypto/rsa/pss.go:322
	// _ = "end of CoverTab[9819]"
}

// VerifyPSS verifies a PSS signature.
//line /usr/local/go/src/crypto/rsa/pss.go:325
//
//line /usr/local/go/src/crypto/rsa/pss.go:325
// A valid signature is indicated by returning a nil error. digest must be the
//line /usr/local/go/src/crypto/rsa/pss.go:325
// result of hashing the input message using the given hash function. The opts
//line /usr/local/go/src/crypto/rsa/pss.go:325
// argument may be nil, in which case sensible defaults are used. opts.Hash is
//line /usr/local/go/src/crypto/rsa/pss.go:325
// ignored.
//line /usr/local/go/src/crypto/rsa/pss.go:331
func VerifyPSS(pub *PublicKey, hash crypto.Hash, digest []byte, sig []byte, opts *PSSOptions) error {
//line /usr/local/go/src/crypto/rsa/pss.go:331
	_go_fuzz_dep_.CoverTab[9838]++
						if boring.Enabled {
//line /usr/local/go/src/crypto/rsa/pss.go:332
		_go_fuzz_dep_.CoverTab[9844]++
							bkey, err := boringPublicKey(pub)
							if err != nil {
//line /usr/local/go/src/crypto/rsa/pss.go:334
			_go_fuzz_dep_.CoverTab[9847]++
								return err
//line /usr/local/go/src/crypto/rsa/pss.go:335
			// _ = "end of CoverTab[9847]"
		} else {
//line /usr/local/go/src/crypto/rsa/pss.go:336
			_go_fuzz_dep_.CoverTab[9848]++
//line /usr/local/go/src/crypto/rsa/pss.go:336
			// _ = "end of CoverTab[9848]"
//line /usr/local/go/src/crypto/rsa/pss.go:336
		}
//line /usr/local/go/src/crypto/rsa/pss.go:336
		// _ = "end of CoverTab[9844]"
//line /usr/local/go/src/crypto/rsa/pss.go:336
		_go_fuzz_dep_.CoverTab[9845]++
							if err := boring.VerifyRSAPSS(bkey, hash, digest, sig, opts.saltLength()); err != nil {
//line /usr/local/go/src/crypto/rsa/pss.go:337
			_go_fuzz_dep_.CoverTab[9849]++
								return ErrVerification
//line /usr/local/go/src/crypto/rsa/pss.go:338
			// _ = "end of CoverTab[9849]"
		} else {
//line /usr/local/go/src/crypto/rsa/pss.go:339
			_go_fuzz_dep_.CoverTab[9850]++
//line /usr/local/go/src/crypto/rsa/pss.go:339
			// _ = "end of CoverTab[9850]"
//line /usr/local/go/src/crypto/rsa/pss.go:339
		}
//line /usr/local/go/src/crypto/rsa/pss.go:339
		// _ = "end of CoverTab[9845]"
//line /usr/local/go/src/crypto/rsa/pss.go:339
		_go_fuzz_dep_.CoverTab[9846]++
							return nil
//line /usr/local/go/src/crypto/rsa/pss.go:340
		// _ = "end of CoverTab[9846]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:341
		_go_fuzz_dep_.CoverTab[9851]++
//line /usr/local/go/src/crypto/rsa/pss.go:341
		// _ = "end of CoverTab[9851]"
//line /usr/local/go/src/crypto/rsa/pss.go:341
	}
//line /usr/local/go/src/crypto/rsa/pss.go:341
	// _ = "end of CoverTab[9838]"
//line /usr/local/go/src/crypto/rsa/pss.go:341
	_go_fuzz_dep_.CoverTab[9839]++
						if len(sig) != pub.Size() {
//line /usr/local/go/src/crypto/rsa/pss.go:342
		_go_fuzz_dep_.CoverTab[9852]++
							return ErrVerification
//line /usr/local/go/src/crypto/rsa/pss.go:343
		// _ = "end of CoverTab[9852]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:344
		_go_fuzz_dep_.CoverTab[9853]++
//line /usr/local/go/src/crypto/rsa/pss.go:344
		// _ = "end of CoverTab[9853]"
//line /usr/local/go/src/crypto/rsa/pss.go:344
	}
//line /usr/local/go/src/crypto/rsa/pss.go:344
	// _ = "end of CoverTab[9839]"
//line /usr/local/go/src/crypto/rsa/pss.go:344
	_go_fuzz_dep_.CoverTab[9840]++

//line /usr/local/go/src/crypto/rsa/pss.go:348
	if opts.saltLength() < PSSSaltLengthEqualsHash {
//line /usr/local/go/src/crypto/rsa/pss.go:348
		_go_fuzz_dep_.CoverTab[9854]++
							return invalidSaltLenErr
//line /usr/local/go/src/crypto/rsa/pss.go:349
		// _ = "end of CoverTab[9854]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:350
		_go_fuzz_dep_.CoverTab[9855]++
//line /usr/local/go/src/crypto/rsa/pss.go:350
		// _ = "end of CoverTab[9855]"
//line /usr/local/go/src/crypto/rsa/pss.go:350
	}
//line /usr/local/go/src/crypto/rsa/pss.go:350
	// _ = "end of CoverTab[9840]"
//line /usr/local/go/src/crypto/rsa/pss.go:350
	_go_fuzz_dep_.CoverTab[9841]++

						emBits := pub.N.BitLen() - 1
						emLen := (emBits + 7) / 8
						em, err := encrypt(pub, sig)
						if err != nil {
//line /usr/local/go/src/crypto/rsa/pss.go:355
		_go_fuzz_dep_.CoverTab[9856]++
							return ErrVerification
//line /usr/local/go/src/crypto/rsa/pss.go:356
		// _ = "end of CoverTab[9856]"
	} else {
//line /usr/local/go/src/crypto/rsa/pss.go:357
		_go_fuzz_dep_.CoverTab[9857]++
//line /usr/local/go/src/crypto/rsa/pss.go:357
		// _ = "end of CoverTab[9857]"
//line /usr/local/go/src/crypto/rsa/pss.go:357
	}
//line /usr/local/go/src/crypto/rsa/pss.go:357
	// _ = "end of CoverTab[9841]"
//line /usr/local/go/src/crypto/rsa/pss.go:357
	_go_fuzz_dep_.CoverTab[9842]++

//line /usr/local/go/src/crypto/rsa/pss.go:364
	for len(em) > emLen && func() bool {
//line /usr/local/go/src/crypto/rsa/pss.go:364
		_go_fuzz_dep_.CoverTab[9858]++
//line /usr/local/go/src/crypto/rsa/pss.go:364
		return len(em) > 0
//line /usr/local/go/src/crypto/rsa/pss.go:364
		// _ = "end of CoverTab[9858]"
//line /usr/local/go/src/crypto/rsa/pss.go:364
	}() {
//line /usr/local/go/src/crypto/rsa/pss.go:364
		_go_fuzz_dep_.CoverTab[9859]++
							if em[0] != 0 {
//line /usr/local/go/src/crypto/rsa/pss.go:365
			_go_fuzz_dep_.CoverTab[9861]++
								return ErrVerification
//line /usr/local/go/src/crypto/rsa/pss.go:366
			// _ = "end of CoverTab[9861]"
		} else {
//line /usr/local/go/src/crypto/rsa/pss.go:367
			_go_fuzz_dep_.CoverTab[9862]++
//line /usr/local/go/src/crypto/rsa/pss.go:367
			// _ = "end of CoverTab[9862]"
//line /usr/local/go/src/crypto/rsa/pss.go:367
		}
//line /usr/local/go/src/crypto/rsa/pss.go:367
		// _ = "end of CoverTab[9859]"
//line /usr/local/go/src/crypto/rsa/pss.go:367
		_go_fuzz_dep_.CoverTab[9860]++
							em = em[1:]
//line /usr/local/go/src/crypto/rsa/pss.go:368
		// _ = "end of CoverTab[9860]"
	}
//line /usr/local/go/src/crypto/rsa/pss.go:369
	// _ = "end of CoverTab[9842]"
//line /usr/local/go/src/crypto/rsa/pss.go:369
	_go_fuzz_dep_.CoverTab[9843]++

						return emsaPSSVerify(digest, em, emBits, opts.saltLength(), hash.New())
//line /usr/local/go/src/crypto/rsa/pss.go:371
	// _ = "end of CoverTab[9843]"
}

//line /usr/local/go/src/crypto/rsa/pss.go:372
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/rsa/pss.go:372
var _ = _go_fuzz_dep_.CoverTab
