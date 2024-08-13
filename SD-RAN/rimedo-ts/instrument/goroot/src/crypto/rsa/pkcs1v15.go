// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:5
package rsa

//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:5
import (
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:5
)
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:5
import (
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:5
)

import (
	"crypto"
	"crypto/internal/boring"
	"crypto/internal/randutil"
	"crypto/subtle"
	"errors"
	"io"
)

//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:18
// PKCS1v15DecryptOptions is for passing options to PKCS #1 v1.5 decryption using
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:18
// the crypto.Decrypter interface.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:20
type PKCS1v15DecryptOptions struct {
	// SessionKeyLen is the length of the session key that is being
	// decrypted. If not zero, then a padding error during decryption will
	// cause a random plaintext of this length to be returned rather than
	// an error. These alternatives happen in constant time.
	SessionKeyLen int
}

// EncryptPKCS1v15 encrypts the given message with RSA and the padding
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:28
// scheme from PKCS #1 v1.5.  The message must be no longer than the
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:28
// length of the public modulus minus 11 bytes.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:28
//
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:28
// The random parameter is used as a source of entropy to ensure that
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:28
// encrypting the same message twice doesn't result in the same
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:28
// ciphertext.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:28
//
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:28
// WARNING: use of this function to encrypt plaintexts other than
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:28
// session keys is dangerous. Use RSA OAEP in new protocols.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:38
func EncryptPKCS1v15(random io.Reader, pub *PublicKey, msg []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:38
	_go_fuzz_dep_.CoverTab[9619]++
							randutil.MaybeReadByte(random)

							if err := checkPub(pub); err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:41
		_go_fuzz_dep_.CoverTab[9625]++
								return nil, err
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:42
		// _ = "end of CoverTab[9625]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:43
		_go_fuzz_dep_.CoverTab[9626]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:43
		// _ = "end of CoverTab[9626]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:43
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:43
	// _ = "end of CoverTab[9619]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:43
	_go_fuzz_dep_.CoverTab[9620]++
							k := pub.Size()
							if len(msg) > k-11 {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:45
		_go_fuzz_dep_.CoverTab[9627]++
								return nil, ErrMessageTooLong
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:46
		// _ = "end of CoverTab[9627]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:47
		_go_fuzz_dep_.CoverTab[9628]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:47
		// _ = "end of CoverTab[9628]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:47
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:47
	// _ = "end of CoverTab[9620]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:47
	_go_fuzz_dep_.CoverTab[9621]++

							if boring.Enabled && func() bool {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:49
		_go_fuzz_dep_.CoverTab[9629]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:49
		return random == boring.RandReader
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:49
		// _ = "end of CoverTab[9629]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:49
	}() {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:49
		_go_fuzz_dep_.CoverTab[9630]++
								bkey, err := boringPublicKey(pub)
								if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:51
			_go_fuzz_dep_.CoverTab[9632]++
									return nil, err
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:52
			// _ = "end of CoverTab[9632]"
		} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:53
			_go_fuzz_dep_.CoverTab[9633]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:53
			// _ = "end of CoverTab[9633]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:53
		}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:53
		// _ = "end of CoverTab[9630]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:53
		_go_fuzz_dep_.CoverTab[9631]++
								return boring.EncryptRSAPKCS1(bkey, msg)
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:54
		// _ = "end of CoverTab[9631]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:55
		_go_fuzz_dep_.CoverTab[9634]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:55
		// _ = "end of CoverTab[9634]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:55
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:55
	// _ = "end of CoverTab[9621]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:55
	_go_fuzz_dep_.CoverTab[9622]++
							boring.UnreachableExceptTests()

//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:59
	em := make([]byte, k)
	em[1] = 2
	ps, mm := em[2:len(em)-len(msg)-1], em[len(em)-len(msg):]
	err := nonZeroRandomBytes(ps, random)
	if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:63
		_go_fuzz_dep_.CoverTab[9635]++
								return nil, err
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:64
		// _ = "end of CoverTab[9635]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:65
		_go_fuzz_dep_.CoverTab[9636]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:65
		// _ = "end of CoverTab[9636]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:65
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:65
	// _ = "end of CoverTab[9622]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:65
	_go_fuzz_dep_.CoverTab[9623]++
							em[len(em)-len(msg)-1] = 0
							copy(mm, msg)

							if boring.Enabled {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:69
		_go_fuzz_dep_.CoverTab[9637]++
								var bkey *boring.PublicKeyRSA
								bkey, err = boringPublicKey(pub)
								if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:72
			_go_fuzz_dep_.CoverTab[9639]++
									return nil, err
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:73
			// _ = "end of CoverTab[9639]"
		} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:74
			_go_fuzz_dep_.CoverTab[9640]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:74
			// _ = "end of CoverTab[9640]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:74
		}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:74
		// _ = "end of CoverTab[9637]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:74
		_go_fuzz_dep_.CoverTab[9638]++
								return boring.EncryptRSANoPadding(bkey, em)
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:75
		// _ = "end of CoverTab[9638]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:76
		_go_fuzz_dep_.CoverTab[9641]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:76
		// _ = "end of CoverTab[9641]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:76
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:76
	// _ = "end of CoverTab[9623]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:76
	_go_fuzz_dep_.CoverTab[9624]++

							return encrypt(pub, em)
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:78
	// _ = "end of CoverTab[9624]"
}

// DecryptPKCS1v15 decrypts a plaintext using RSA and the padding scheme from PKCS #1 v1.5.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:81
// The random parameter is legacy and ignored, and it can be as nil.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:81
//
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:81
// Note that whether this function returns an error or not discloses secret
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:81
// information. If an attacker can cause this function to run repeatedly and
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:81
// learn whether each instance returned an error then they can decrypt and
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:81
// forge signatures as if they had the private key. See
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:81
// DecryptPKCS1v15SessionKey for a way of solving this problem.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:89
func DecryptPKCS1v15(random io.Reader, priv *PrivateKey, ciphertext []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:89
	_go_fuzz_dep_.CoverTab[9642]++
							if err := checkPub(&priv.PublicKey); err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:90
		_go_fuzz_dep_.CoverTab[9647]++
								return nil, err
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:91
		// _ = "end of CoverTab[9647]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:92
		_go_fuzz_dep_.CoverTab[9648]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:92
		// _ = "end of CoverTab[9648]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:92
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:92
	// _ = "end of CoverTab[9642]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:92
	_go_fuzz_dep_.CoverTab[9643]++

							if boring.Enabled {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:94
		_go_fuzz_dep_.CoverTab[9649]++
								bkey, err := boringPrivateKey(priv)
								if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:96
			_go_fuzz_dep_.CoverTab[9652]++
									return nil, err
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:97
			// _ = "end of CoverTab[9652]"
		} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:98
			_go_fuzz_dep_.CoverTab[9653]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:98
			// _ = "end of CoverTab[9653]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:98
		}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:98
		// _ = "end of CoverTab[9649]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:98
		_go_fuzz_dep_.CoverTab[9650]++
								out, err := boring.DecryptRSAPKCS1(bkey, ciphertext)
								if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:100
			_go_fuzz_dep_.CoverTab[9654]++
									return nil, ErrDecryption
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:101
			// _ = "end of CoverTab[9654]"
		} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:102
			_go_fuzz_dep_.CoverTab[9655]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:102
			// _ = "end of CoverTab[9655]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:102
		}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:102
		// _ = "end of CoverTab[9650]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:102
		_go_fuzz_dep_.CoverTab[9651]++
								return out, nil
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:103
		// _ = "end of CoverTab[9651]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:104
		_go_fuzz_dep_.CoverTab[9656]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:104
		// _ = "end of CoverTab[9656]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:104
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:104
	// _ = "end of CoverTab[9643]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:104
	_go_fuzz_dep_.CoverTab[9644]++

							valid, out, index, err := decryptPKCS1v15(priv, ciphertext)
							if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:107
		_go_fuzz_dep_.CoverTab[9657]++
								return nil, err
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:108
		// _ = "end of CoverTab[9657]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:109
		_go_fuzz_dep_.CoverTab[9658]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:109
		// _ = "end of CoverTab[9658]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:109
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:109
	// _ = "end of CoverTab[9644]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:109
	_go_fuzz_dep_.CoverTab[9645]++
							if valid == 0 {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:110
		_go_fuzz_dep_.CoverTab[9659]++
								return nil, ErrDecryption
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:111
		// _ = "end of CoverTab[9659]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:112
		_go_fuzz_dep_.CoverTab[9660]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:112
		// _ = "end of CoverTab[9660]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:112
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:112
	// _ = "end of CoverTab[9645]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:112
	_go_fuzz_dep_.CoverTab[9646]++
							return out[index:], nil
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:113
	// _ = "end of CoverTab[9646]"
}

// DecryptPKCS1v15SessionKey decrypts a session key using RSA and the padding scheme from PKCS #1 v1.5.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// The random parameter is legacy and ignored, and it can be as nil.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// It returns an error if the ciphertext is the wrong length or if the
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// ciphertext is greater than the public modulus. Otherwise, no error is
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// returned. If the padding is valid, the resulting plaintext message is copied
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// into key. Otherwise, key is unchanged. These alternatives occur in constant
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// time. It is intended that the user of this function generate a random
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// session key beforehand and continue the protocol with the resulting value.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// This will remove any possibility that an attacker can learn any information
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// about the plaintext.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// See “Chosen Ciphertext Attacks Against Protocols Based on the RSA
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// Encryption Standard PKCS #1”, Daniel Bleichenbacher, Advances in Cryptology
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// (Crypto '98).
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
//
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// Note that if the session key is too small then it may be possible for an
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// attacker to brute-force it. If they can do that then they can learn whether
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// a random value was used (because it'll be different for the same ciphertext)
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// and thus whether the padding was correct. This defeats the point of this
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:116
// function. Using at least a 16-byte key will protect against this attack.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:135
func DecryptPKCS1v15SessionKey(random io.Reader, priv *PrivateKey, ciphertext []byte, key []byte) error {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:135
	_go_fuzz_dep_.CoverTab[9661]++
							if err := checkPub(&priv.PublicKey); err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:136
		_go_fuzz_dep_.CoverTab[9666]++
								return err
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:137
		// _ = "end of CoverTab[9666]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:138
		_go_fuzz_dep_.CoverTab[9667]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:138
		// _ = "end of CoverTab[9667]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:138
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:138
	// _ = "end of CoverTab[9661]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:138
	_go_fuzz_dep_.CoverTab[9662]++
							k := priv.Size()
							if k-(len(key)+3+8) < 0 {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:140
		_go_fuzz_dep_.CoverTab[9668]++
								return ErrDecryption
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:141
		// _ = "end of CoverTab[9668]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:142
		_go_fuzz_dep_.CoverTab[9669]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:142
		// _ = "end of CoverTab[9669]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:142
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:142
	// _ = "end of CoverTab[9662]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:142
	_go_fuzz_dep_.CoverTab[9663]++

							valid, em, index, err := decryptPKCS1v15(priv, ciphertext)
							if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:145
		_go_fuzz_dep_.CoverTab[9670]++
								return err
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:146
		// _ = "end of CoverTab[9670]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:147
		_go_fuzz_dep_.CoverTab[9671]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:147
		// _ = "end of CoverTab[9671]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:147
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:147
	// _ = "end of CoverTab[9663]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:147
	_go_fuzz_dep_.CoverTab[9664]++

							if len(em) != k {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:149
		_go_fuzz_dep_.CoverTab[9672]++

//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:152
		return ErrDecryption
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:152
		// _ = "end of CoverTab[9672]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:153
		_go_fuzz_dep_.CoverTab[9673]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:153
		// _ = "end of CoverTab[9673]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:153
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:153
	// _ = "end of CoverTab[9664]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:153
	_go_fuzz_dep_.CoverTab[9665]++

							valid &= subtle.ConstantTimeEq(int32(len(em)-index), int32(len(key)))
							subtle.ConstantTimeCopy(valid, key, em[len(em)-len(key):])
							return nil
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:157
	// _ = "end of CoverTab[9665]"
}

// decryptPKCS1v15 decrypts ciphertext using priv. It returns one or zero in
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:160
// valid that indicates whether the plaintext was correctly structured.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:160
// In either case, the plaintext is returned in em so that it may be read
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:160
// independently of whether it was valid in order to maintain constant memory
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:160
// access patterns. If the plaintext was valid then index contains the index of
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:160
// the original message in em, to allow constant time padding removal.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:166
func decryptPKCS1v15(priv *PrivateKey, ciphertext []byte) (valid int, em []byte, index int, err error) {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:166
	_go_fuzz_dep_.CoverTab[9674]++
							k := priv.Size()
							if k < 11 {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:168
		_go_fuzz_dep_.CoverTab[9678]++
								err = ErrDecryption
								return
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:170
		// _ = "end of CoverTab[9678]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:171
		_go_fuzz_dep_.CoverTab[9679]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:171
		// _ = "end of CoverTab[9679]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:171
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:171
	// _ = "end of CoverTab[9674]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:171
	_go_fuzz_dep_.CoverTab[9675]++

							if boring.Enabled {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:173
		_go_fuzz_dep_.CoverTab[9680]++
								var bkey *boring.PrivateKeyRSA
								bkey, err = boringPrivateKey(priv)
								if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:176
			_go_fuzz_dep_.CoverTab[9682]++
									return
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:177
			// _ = "end of CoverTab[9682]"
		} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:178
			_go_fuzz_dep_.CoverTab[9683]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:178
			// _ = "end of CoverTab[9683]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:178
		}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:178
		// _ = "end of CoverTab[9680]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:178
		_go_fuzz_dep_.CoverTab[9681]++
								em, err = boring.DecryptRSANoPadding(bkey, ciphertext)
								if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:180
			_go_fuzz_dep_.CoverTab[9684]++
									return
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:181
			// _ = "end of CoverTab[9684]"
		} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:182
			_go_fuzz_dep_.CoverTab[9685]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:182
			// _ = "end of CoverTab[9685]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:182
		}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:182
		// _ = "end of CoverTab[9681]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:183
		_go_fuzz_dep_.CoverTab[9686]++
								em, err = decrypt(priv, ciphertext, noCheck)
								if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:185
			_go_fuzz_dep_.CoverTab[9687]++
									return
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:186
			// _ = "end of CoverTab[9687]"
		} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:187
			_go_fuzz_dep_.CoverTab[9688]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:187
			// _ = "end of CoverTab[9688]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:187
		}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:187
		// _ = "end of CoverTab[9686]"
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:188
	// _ = "end of CoverTab[9675]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:188
	_go_fuzz_dep_.CoverTab[9676]++

							firstByteIsZero := subtle.ConstantTimeByteEq(em[0], 0)
							secondByteIsTwo := subtle.ConstantTimeByteEq(em[1], 2)

//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:197
	lookingForIndex := 1

	for i := 2; i < len(em); i++ {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:199
		_go_fuzz_dep_.CoverTab[9689]++
								equals0 := subtle.ConstantTimeByteEq(em[i], 0)
								index = subtle.ConstantTimeSelect(lookingForIndex&equals0, i, index)
								lookingForIndex = subtle.ConstantTimeSelect(equals0, 0, lookingForIndex)
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:202
		// _ = "end of CoverTab[9689]"
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:203
	// _ = "end of CoverTab[9676]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:203
	_go_fuzz_dep_.CoverTab[9677]++

//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:207
	validPS := subtle.ConstantTimeLessOrEq(2+8, index)

							valid = firstByteIsZero & secondByteIsTwo & (^lookingForIndex & 1) & validPS
							index = subtle.ConstantTimeSelect(valid, index+1, 0)
							return valid, em, index, nil
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:211
	// _ = "end of CoverTab[9677]"
}

// nonZeroRandomBytes fills the given slice with non-zero random octets.
func nonZeroRandomBytes(s []byte, random io.Reader) (err error) {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:215
	_go_fuzz_dep_.CoverTab[9690]++
							_, err = io.ReadFull(random, s)
							if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:217
		_go_fuzz_dep_.CoverTab[9693]++
								return
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:218
		// _ = "end of CoverTab[9693]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:219
		_go_fuzz_dep_.CoverTab[9694]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:219
		// _ = "end of CoverTab[9694]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:219
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:219
	// _ = "end of CoverTab[9690]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:219
	_go_fuzz_dep_.CoverTab[9691]++

							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:221
		_go_fuzz_dep_.CoverTab[9695]++
								for s[i] == 0 {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:222
			_go_fuzz_dep_.CoverTab[9696]++
									_, err = io.ReadFull(random, s[i:i+1])
									if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:224
				_go_fuzz_dep_.CoverTab[9698]++
										return
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:225
				// _ = "end of CoverTab[9698]"
			} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:226
				_go_fuzz_dep_.CoverTab[9699]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:226
				// _ = "end of CoverTab[9699]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:226
			}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:226
			// _ = "end of CoverTab[9696]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:226
			_go_fuzz_dep_.CoverTab[9697]++

//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:229
			s[i] ^= 0x42
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:229
			// _ = "end of CoverTab[9697]"
		}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:230
		// _ = "end of CoverTab[9695]"
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:231
	// _ = "end of CoverTab[9691]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:231
	_go_fuzz_dep_.CoverTab[9692]++

							return
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:233
	// _ = "end of CoverTab[9692]"
}

// These are ASN1 DER structures:
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:236
//
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:236
//	DigestInfo ::= SEQUENCE {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:236
//	  digestAlgorithm AlgorithmIdentifier,
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:236
//	  digest OCTET STRING
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:236
//	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:236
//
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:236
// For performance, we don't use the generic ASN1 encoder. Rather, we
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:236
// precompute a prefix of the digest value that makes a valid ASN1 DER string
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:236
// with the correct contents.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:246
var hashPrefixes = map[crypto.Hash][]byte{
	crypto.MD5:		{0x30, 0x20, 0x30, 0x0c, 0x06, 0x08, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x02, 0x05, 0x05, 0x00, 0x04, 0x10},
	crypto.SHA1:		{0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14},
	crypto.SHA224:		{0x30, 0x2d, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x04, 0x05, 0x00, 0x04, 0x1c},
	crypto.SHA256:		{0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20},
	crypto.SHA384:		{0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30},
	crypto.SHA512:		{0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40},
	crypto.MD5SHA1:		{},
	crypto.RIPEMD160:	{0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14},
}

// SignPKCS1v15 calculates the signature of hashed using
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:257
// RSASSA-PKCS1-V1_5-SIGN from RSA PKCS #1 v1.5.  Note that hashed must
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:257
// be the result of hashing the input message using the given hash
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:257
// function. If hash is zero, hashed is signed directly. This isn't
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:257
// advisable except for interoperability.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:257
//
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:257
// The random parameter is legacy and ignored, and it can be as nil.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:257
//
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:257
// This function is deterministic. Thus, if the set of possible
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:257
// messages is small, an attacker may be able to build a map from
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:257
// messages to signatures and identify the signed messages. As ever,
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:257
// signatures provide authenticity, not confidentiality.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:269
func SignPKCS1v15(random io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:269
	_go_fuzz_dep_.CoverTab[9700]++
							hashLen, prefix, err := pkcs1v15HashInfo(hash, len(hashed))
							if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:271
		_go_fuzz_dep_.CoverTab[9705]++
								return nil, err
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:272
		// _ = "end of CoverTab[9705]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:273
		_go_fuzz_dep_.CoverTab[9706]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:273
		// _ = "end of CoverTab[9706]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:273
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:273
	// _ = "end of CoverTab[9700]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:273
	_go_fuzz_dep_.CoverTab[9701]++

							tLen := len(prefix) + hashLen
							k := priv.Size()
							if k < tLen+11 {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:277
		_go_fuzz_dep_.CoverTab[9707]++
								return nil, ErrMessageTooLong
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:278
		// _ = "end of CoverTab[9707]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:279
		_go_fuzz_dep_.CoverTab[9708]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:279
		// _ = "end of CoverTab[9708]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:279
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:279
	// _ = "end of CoverTab[9701]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:279
	_go_fuzz_dep_.CoverTab[9702]++

							if boring.Enabled {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:281
		_go_fuzz_dep_.CoverTab[9709]++
								bkey, err := boringPrivateKey(priv)
								if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:283
			_go_fuzz_dep_.CoverTab[9711]++
									return nil, err
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:284
			// _ = "end of CoverTab[9711]"
		} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:285
			_go_fuzz_dep_.CoverTab[9712]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:285
			// _ = "end of CoverTab[9712]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:285
		}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:285
		// _ = "end of CoverTab[9709]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:285
		_go_fuzz_dep_.CoverTab[9710]++
								return boring.SignRSAPKCS1v15(bkey, hash, hashed)
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:286
		// _ = "end of CoverTab[9710]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:287
		_go_fuzz_dep_.CoverTab[9713]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:287
		// _ = "end of CoverTab[9713]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:287
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:287
	// _ = "end of CoverTab[9702]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:287
	_go_fuzz_dep_.CoverTab[9703]++

//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:290
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:292
		_go_fuzz_dep_.CoverTab[9714]++
								em[i] = 0xff
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:293
		// _ = "end of CoverTab[9714]"
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:294
	// _ = "end of CoverTab[9703]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:294
	_go_fuzz_dep_.CoverTab[9704]++
							copy(em[k-tLen:k-hashLen], prefix)
							copy(em[k-hashLen:k], hashed)

							return decrypt(priv, em, withCheck)
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:298
	// _ = "end of CoverTab[9704]"
}

// VerifyPKCS1v15 verifies an RSA PKCS #1 v1.5 signature.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:301
// hashed is the result of hashing the input message using the given hash
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:301
// function and sig is the signature. A valid signature is indicated by
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:301
// returning a nil error. If hash is zero then hashed is used directly. This
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:301
// isn't advisable except for interoperability.
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:306
func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:306
	_go_fuzz_dep_.CoverTab[9715]++
							if boring.Enabled {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:307
		_go_fuzz_dep_.CoverTab[9723]++
								bkey, err := boringPublicKey(pub)
								if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:309
			_go_fuzz_dep_.CoverTab[9726]++
									return err
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:310
			// _ = "end of CoverTab[9726]"
		} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:311
			_go_fuzz_dep_.CoverTab[9727]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:311
			// _ = "end of CoverTab[9727]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:311
		}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:311
		// _ = "end of CoverTab[9723]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:311
		_go_fuzz_dep_.CoverTab[9724]++
								if err := boring.VerifyRSAPKCS1v15(bkey, hash, hashed, sig); err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:312
			_go_fuzz_dep_.CoverTab[9728]++
									return ErrVerification
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:313
			// _ = "end of CoverTab[9728]"
		} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:314
			_go_fuzz_dep_.CoverTab[9729]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:314
			// _ = "end of CoverTab[9729]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:314
		}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:314
		// _ = "end of CoverTab[9724]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:314
		_go_fuzz_dep_.CoverTab[9725]++
								return nil
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:315
		// _ = "end of CoverTab[9725]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:316
		_go_fuzz_dep_.CoverTab[9730]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:316
		// _ = "end of CoverTab[9730]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:316
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:316
	// _ = "end of CoverTab[9715]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:316
	_go_fuzz_dep_.CoverTab[9716]++

							hashLen, prefix, err := pkcs1v15HashInfo(hash, len(hashed))
							if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:319
		_go_fuzz_dep_.CoverTab[9731]++
								return err
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:320
		// _ = "end of CoverTab[9731]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:321
		_go_fuzz_dep_.CoverTab[9732]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:321
		// _ = "end of CoverTab[9732]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:321
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:321
	// _ = "end of CoverTab[9716]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:321
	_go_fuzz_dep_.CoverTab[9717]++

							tLen := len(prefix) + hashLen
							k := pub.Size()
							if k < tLen+11 {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:325
		_go_fuzz_dep_.CoverTab[9733]++
								return ErrVerification
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:326
		// _ = "end of CoverTab[9733]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:327
		_go_fuzz_dep_.CoverTab[9734]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:327
		// _ = "end of CoverTab[9734]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:327
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:327
	// _ = "end of CoverTab[9717]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:327
	_go_fuzz_dep_.CoverTab[9718]++

//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:332
	if k != len(sig) {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:332
		_go_fuzz_dep_.CoverTab[9735]++
								return ErrVerification
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:333
		// _ = "end of CoverTab[9735]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:334
		_go_fuzz_dep_.CoverTab[9736]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:334
		// _ = "end of CoverTab[9736]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:334
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:334
	// _ = "end of CoverTab[9718]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:334
	_go_fuzz_dep_.CoverTab[9719]++

							em, err := encrypt(pub, sig)
							if err != nil {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:337
		_go_fuzz_dep_.CoverTab[9737]++
								return ErrVerification
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:338
		// _ = "end of CoverTab[9737]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:339
		_go_fuzz_dep_.CoverTab[9738]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:339
		// _ = "end of CoverTab[9738]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:339
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:339
	// _ = "end of CoverTab[9719]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:339
	_go_fuzz_dep_.CoverTab[9720]++

//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:342
	ok := subtle.ConstantTimeByteEq(em[0], 0)
	ok &= subtle.ConstantTimeByteEq(em[1], 1)
	ok &= subtle.ConstantTimeCompare(em[k-hashLen:k], hashed)
	ok &= subtle.ConstantTimeCompare(em[k-tLen:k-hashLen], prefix)
	ok &= subtle.ConstantTimeByteEq(em[k-tLen-1], 0)

	for i := 2; i < k-tLen-1; i++ {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:348
		_go_fuzz_dep_.CoverTab[9739]++
								ok &= subtle.ConstantTimeByteEq(em[i], 0xff)
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:349
		// _ = "end of CoverTab[9739]"
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:350
	// _ = "end of CoverTab[9720]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:350
	_go_fuzz_dep_.CoverTab[9721]++

							if ok != 1 {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:352
		_go_fuzz_dep_.CoverTab[9740]++
								return ErrVerification
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:353
		// _ = "end of CoverTab[9740]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:354
		_go_fuzz_dep_.CoverTab[9741]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:354
		// _ = "end of CoverTab[9741]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:354
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:354
	// _ = "end of CoverTab[9721]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:354
	_go_fuzz_dep_.CoverTab[9722]++

							return nil
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:356
	// _ = "end of CoverTab[9722]"
}

func pkcs1v15HashInfo(hash crypto.Hash, inLen int) (hashLen int, prefix []byte, err error) {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:359
	_go_fuzz_dep_.CoverTab[9742]++

//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:362
	if hash == 0 {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:362
		_go_fuzz_dep_.CoverTab[9746]++
								return inLen, nil, nil
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:363
		// _ = "end of CoverTab[9746]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:364
		_go_fuzz_dep_.CoverTab[9747]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:364
		// _ = "end of CoverTab[9747]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:364
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:364
	// _ = "end of CoverTab[9742]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:364
	_go_fuzz_dep_.CoverTab[9743]++

							hashLen = hash.Size()
							if inLen != hashLen {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:367
		_go_fuzz_dep_.CoverTab[9748]++
								return 0, nil, errors.New("crypto/rsa: input must be hashed message")
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:368
		// _ = "end of CoverTab[9748]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:369
		_go_fuzz_dep_.CoverTab[9749]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:369
		// _ = "end of CoverTab[9749]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:369
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:369
	// _ = "end of CoverTab[9743]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:369
	_go_fuzz_dep_.CoverTab[9744]++
							prefix, ok := hashPrefixes[hash]
							if !ok {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:371
		_go_fuzz_dep_.CoverTab[9750]++
								return 0, nil, errors.New("crypto/rsa: unsupported hash function")
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:372
		// _ = "end of CoverTab[9750]"
	} else {
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:373
		_go_fuzz_dep_.CoverTab[9751]++
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:373
		// _ = "end of CoverTab[9751]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:373
	}
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:373
	// _ = "end of CoverTab[9744]"
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:373
	_go_fuzz_dep_.CoverTab[9745]++
							return
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:374
	// _ = "end of CoverTab[9745]"
}

//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:375
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/rsa/pkcs1v15.go:375
var _ = _go_fuzz_dep_.CoverTab
