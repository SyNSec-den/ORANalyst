//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:1
package rfc3961

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:1
)

import (
	"bytes"

	"github.com/jcmturner/gokrb5/v8/crypto/etype"
)

const (
	prfconstant = "prf"
)

// DeriveRandom implements the RFC 3961 defined function: DR(Key, Constant) = k-truncate(E(Key, Constant, initial-cipher-state)).
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:13
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:13
// key: base key or protocol key. Likely to be a key from a keytab file.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:13
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:13
// usage: a constant.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:13
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:13
// n: block size in bits (not bytes) - note if you use something like aes.BlockSize this is in bytes.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:13
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:13
// k: key length / key seed length in bits. Eg. for AES256 this value is 256.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:13
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:13
// e: the encryption etype function to use.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:24
func DeriveRandom(key, usage []byte, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:24
	_go_fuzz_dep_.CoverTab[85568]++
														n := e.GetCypherBlockBitLength()
														k := e.GetKeySeedBitLength()

														nFoldUsage := Nfold(usage, n)

														out := make([]byte, k/8)

														_, K, err := e.EncryptData(key, nFoldUsage)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:33
		_go_fuzz_dep_.CoverTab[85571]++
															return out, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:34
		// _ = "end of CoverTab[85571]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:35
		_go_fuzz_dep_.CoverTab[85572]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:35
		// _ = "end of CoverTab[85572]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:35
	// _ = "end of CoverTab[85568]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:35
	_go_fuzz_dep_.CoverTab[85569]++
														for i := copy(out, K); i < len(out); {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:36
		_go_fuzz_dep_.CoverTab[85573]++
															_, K, _ = e.EncryptData(key, K)
															i = i + copy(out[i:], K)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:38
		// _ = "end of CoverTab[85573]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:39
	// _ = "end of CoverTab[85569]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:39
	_go_fuzz_dep_.CoverTab[85570]++
														return out, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:40
	// _ = "end of CoverTab[85570]"
}

// DeriveKey derives a key from the protocol key based on the usage and the etype's specific methods.
func DeriveKey(protocolKey, usage []byte, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:44
	_go_fuzz_dep_.CoverTab[85574]++
														r, err := e.DeriveRandom(protocolKey, usage)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:46
		_go_fuzz_dep_.CoverTab[85576]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:47
		// _ = "end of CoverTab[85576]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:48
		_go_fuzz_dep_.CoverTab[85577]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:48
		// _ = "end of CoverTab[85577]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:48
	// _ = "end of CoverTab[85574]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:48
	_go_fuzz_dep_.CoverTab[85575]++
														return e.RandomToKey(r), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:49
	// _ = "end of CoverTab[85575]"
}

// RandomToKey returns a key from the bytes provided according to the definition in RFC 3961.
func RandomToKey(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:53
	_go_fuzz_dep_.CoverTab[85578]++
														return b
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:54
	// _ = "end of CoverTab[85578]"
}

// DES3RandomToKey returns a key from the bytes provided according to the definition in RFC 3961 for DES3 etypes.
func DES3RandomToKey(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:58
	_go_fuzz_dep_.CoverTab[85579]++
														r := fixWeakKey(stretch56Bits(b[:7]))
														r2 := fixWeakKey(stretch56Bits(b[7:14]))
														r = append(r, r2...)
														r3 := fixWeakKey(stretch56Bits(b[14:21]))
														r = append(r, r3...)
														return r
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:64
	// _ = "end of CoverTab[85579]"
}

// DES3StringToKey returns a key derived from the string provided according to the definition in RFC 3961 for DES3 etypes.
func DES3StringToKey(secret, salt string, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:68
	_go_fuzz_dep_.CoverTab[85580]++
														s := secret + salt
														tkey := e.RandomToKey(Nfold([]byte(s), e.GetKeySeedBitLength()))
														return e.DeriveKey(tkey, []byte("kerberos"))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:71
	// _ = "end of CoverTab[85580]"
}

// PseudoRandom function as defined in RFC 3961
func PseudoRandom(key, b []byte, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:75
	_go_fuzz_dep_.CoverTab[85581]++
														h := e.GetHashFunc()()
														h.Write(b)
														tmp := h.Sum(nil)[:e.GetMessageBlockByteSize()]
														k, err := e.DeriveKey(key, []byte(prfconstant))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:80
		_go_fuzz_dep_.CoverTab[85584]++
															return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:81
		// _ = "end of CoverTab[85584]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:82
		_go_fuzz_dep_.CoverTab[85585]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:82
		// _ = "end of CoverTab[85585]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:82
	// _ = "end of CoverTab[85581]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:82
	_go_fuzz_dep_.CoverTab[85582]++
														_, prf, err := e.EncryptData(k, tmp)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:84
		_go_fuzz_dep_.CoverTab[85586]++
															return []byte{}, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:85
		// _ = "end of CoverTab[85586]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:86
		_go_fuzz_dep_.CoverTab[85587]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:86
		// _ = "end of CoverTab[85587]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:86
	// _ = "end of CoverTab[85582]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:86
	_go_fuzz_dep_.CoverTab[85583]++
														return prf, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:87
	// _ = "end of CoverTab[85583]"
}

func stretch56Bits(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:90
	_go_fuzz_dep_.CoverTab[85588]++
														d := make([]byte, len(b), len(b))
														copy(d, b)
														var lb byte
														for i, v := range d {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:94
		_go_fuzz_dep_.CoverTab[85590]++
															bv, nb := calcEvenParity(v)
															d[i] = nb
															if bv != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:97
			_go_fuzz_dep_.CoverTab[85591]++
																lb = lb | (1 << uint(i+1))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:98
			// _ = "end of CoverTab[85591]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:99
			_go_fuzz_dep_.CoverTab[85592]++
																lb = lb &^ (1 << uint(i+1))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:100
			// _ = "end of CoverTab[85592]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:101
		// _ = "end of CoverTab[85590]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:102
	// _ = "end of CoverTab[85588]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:102
	_go_fuzz_dep_.CoverTab[85589]++
														_, lb = calcEvenParity(lb)
														d = append(d, lb)
														return d
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:105
	// _ = "end of CoverTab[85589]"
}

func calcEvenParity(b byte) (uint8, uint8) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:108
	_go_fuzz_dep_.CoverTab[85593]++
														lowestbit := b & 0x01
	// c counter of 1s in the first 7 bits of the byte
	var c int

	for p := 1; p < 8; p++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:113
		_go_fuzz_dep_.CoverTab[85596]++
															v := b & (1 << uint(p))
															if v != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:115
			_go_fuzz_dep_.CoverTab[85597]++
																c++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:116
			// _ = "end of CoverTab[85597]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:117
			_go_fuzz_dep_.CoverTab[85598]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:117
			// _ = "end of CoverTab[85598]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:117
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:117
		// _ = "end of CoverTab[85596]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:118
	// _ = "end of CoverTab[85593]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:118
	_go_fuzz_dep_.CoverTab[85594]++
														if c%2 == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:119
		_go_fuzz_dep_.CoverTab[85599]++

															b = b | 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:121
		// _ = "end of CoverTab[85599]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:122
		_go_fuzz_dep_.CoverTab[85600]++

															b = b &^ 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:124
		// _ = "end of CoverTab[85600]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:125
	// _ = "end of CoverTab[85594]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:125
	_go_fuzz_dep_.CoverTab[85595]++
														return lowestbit, b
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:126
	// _ = "end of CoverTab[85595]"
}

func fixWeakKey(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:129
	_go_fuzz_dep_.CoverTab[85601]++
														if weak(b) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:130
		_go_fuzz_dep_.CoverTab[85603]++
															b[7] ^= 0xF0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:131
		// _ = "end of CoverTab[85603]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:132
		_go_fuzz_dep_.CoverTab[85604]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:132
		// _ = "end of CoverTab[85604]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:132
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:132
	// _ = "end of CoverTab[85601]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:132
	_go_fuzz_dep_.CoverTab[85602]++
														return b
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:133
	// _ = "end of CoverTab[85602]"
}

func weak(b []byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:136
	_go_fuzz_dep_.CoverTab[85605]++

														weakKeys := [4][]byte{
		{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01},
		{0xFE, 0xFE, 0xFE, 0xFE, 0xFE, 0xFE, 0xFE, 0xFE},
		{0xE0, 0xE0, 0xE0, 0xE0, 0xF1, 0xF1, 0xF1, 0xF1},
		{0x1F, 0x1F, 0x1F, 0x1F, 0x0E, 0x0E, 0x0E, 0x0E},
	}
	semiWeakKeys := [12][]byte{
		{0x01, 0x1F, 0x01, 0x1F, 0x01, 0x0E, 0x01, 0x0E},
		{0x1F, 0x01, 0x1F, 0x01, 0x0E, 0x01, 0x0E, 0x01},
		{0x01, 0xE0, 0x01, 0xE0, 0x01, 0xF1, 0x01, 0xF1},
		{0xE0, 0x01, 0xE0, 0x01, 0xF1, 0x01, 0xF1, 0x01},
		{0x01, 0xFE, 0x01, 0xFE, 0x01, 0xFE, 0x01, 0xFE},
		{0xFE, 0x01, 0xFE, 0x01, 0xFE, 0x01, 0xFE, 0x01},
		{0x1F, 0xE0, 0x1F, 0xE0, 0x0E, 0xF1, 0x0E, 0xF1},
		{0xE0, 0x1F, 0xE0, 0x1F, 0xF1, 0x0E, 0xF1, 0x0E},
		{0x1F, 0xFE, 0x1F, 0xFE, 0x0E, 0xFE, 0x0E, 0xFE},
		{0xFE, 0x1F, 0xFE, 0x1F, 0xFE, 0x0E, 0xFE, 0x0E},
		{0xE0, 0xFE, 0xE0, 0xFE, 0xF1, 0xFE, 0xF1, 0xFE},
		{0xFE, 0xE0, 0xFE, 0xE0, 0xFE, 0xF1, 0xFE, 0xF1},
	}
	for _, k := range weakKeys {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:158
		_go_fuzz_dep_.CoverTab[85608]++
															if bytes.Equal(b, k) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:159
			_go_fuzz_dep_.CoverTab[85609]++
																return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:160
			// _ = "end of CoverTab[85609]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:161
			_go_fuzz_dep_.CoverTab[85610]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:161
			// _ = "end of CoverTab[85610]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:161
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:161
		// _ = "end of CoverTab[85608]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:162
	// _ = "end of CoverTab[85605]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:162
	_go_fuzz_dep_.CoverTab[85606]++
														for _, k := range semiWeakKeys {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:163
		_go_fuzz_dep_.CoverTab[85611]++
															if bytes.Equal(b, k) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:164
			_go_fuzz_dep_.CoverTab[85612]++
																return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:165
			// _ = "end of CoverTab[85612]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:166
			_go_fuzz_dep_.CoverTab[85613]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:166
			// _ = "end of CoverTab[85613]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:166
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:166
		// _ = "end of CoverTab[85611]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:167
	// _ = "end of CoverTab[85606]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:167
	_go_fuzz_dep_.CoverTab[85607]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:168
	// _ = "end of CoverTab[85607]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:169
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/keyDerivation.go:169
var _ = _go_fuzz_dep_.CoverTab
