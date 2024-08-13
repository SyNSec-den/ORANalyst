//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:1
package rfc8009

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:1
)

import (
	"crypto/hmac"
	"encoding/binary"
	"encoding/hex"
	"errors"

	"github.com/jcmturner/gokrb5/v8/crypto/etype"
	"github.com/jcmturner/gokrb5/v8/iana/etypeID"
	"golang.org/x/crypto/pbkdf2"
)

const (
	s2kParamsZero = 32768
)

// DeriveRandom for key derivation as defined in RFC 8009
func DeriveRandom(protocolKey, usage []byte, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:19
	_go_fuzz_dep_.CoverTab[85898]++
														h := e.GetHashFunc()()
														return KDF_HMAC_SHA2(protocolKey, []byte("prf"), usage, h.Size(), e), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:21
	// _ = "end of CoverTab[85898]"
}

// DeriveKey derives a key from the protocol key based on the usage and the etype's specific methods.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:24
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:24
// https://tools.ietf.org/html/rfc8009#section-5
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:27
func DeriveKey(protocolKey, label []byte, e etype.EType) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:27
	_go_fuzz_dep_.CoverTab[85899]++
														var context []byte
														var kl int

														if e.GetETypeID() == etypeID.AES256_CTS_HMAC_SHA384_192 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:31
		_go_fuzz_dep_.CoverTab[85902]++
	Swtch:
		switch label[len(label)-1] {
		case 0x73:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:34
			_go_fuzz_dep_.CoverTab[85903]++

																kerblabel := []byte("kerberos")
																if len(label) != len(kerblabel) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:37
				_go_fuzz_dep_.CoverTab[85908]++
																	break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:38
				// _ = "end of CoverTab[85908]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:39
				_go_fuzz_dep_.CoverTab[85909]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:39
				// _ = "end of CoverTab[85909]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:39
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:39
			// _ = "end of CoverTab[85903]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:39
			_go_fuzz_dep_.CoverTab[85904]++
																for i, b := range label {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:40
				_go_fuzz_dep_.CoverTab[85910]++
																	if b != kerblabel[i] {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:41
					_go_fuzz_dep_.CoverTab[85911]++
																		kl = e.GetKeySeedBitLength()
																		break Swtch
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:43
					// _ = "end of CoverTab[85911]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:44
					_go_fuzz_dep_.CoverTab[85912]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:44
					// _ = "end of CoverTab[85912]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:44
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:44
				// _ = "end of CoverTab[85910]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:45
			// _ = "end of CoverTab[85904]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:45
			_go_fuzz_dep_.CoverTab[85905]++
																if kl == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:46
				_go_fuzz_dep_.CoverTab[85913]++

																	kl = 256
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:48
				// _ = "end of CoverTab[85913]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:49
				_go_fuzz_dep_.CoverTab[85914]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:49
				// _ = "end of CoverTab[85914]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:49
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:49
			// _ = "end of CoverTab[85905]"
		case 0xAA:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:50
			_go_fuzz_dep_.CoverTab[85906]++

																kl = 256
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:52
			// _ = "end of CoverTab[85906]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:52
		default:
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:52
			_go_fuzz_dep_.CoverTab[85907]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:52
			// _ = "end of CoverTab[85907]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:53
		// _ = "end of CoverTab[85902]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:54
		_go_fuzz_dep_.CoverTab[85915]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:54
		// _ = "end of CoverTab[85915]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:54
	// _ = "end of CoverTab[85899]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:54
	_go_fuzz_dep_.CoverTab[85900]++
														if kl == 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:55
		_go_fuzz_dep_.CoverTab[85916]++
															kl = e.GetKeySeedBitLength()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:56
		// _ = "end of CoverTab[85916]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:57
		_go_fuzz_dep_.CoverTab[85917]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:57
		// _ = "end of CoverTab[85917]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:57
	// _ = "end of CoverTab[85900]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:57
	_go_fuzz_dep_.CoverTab[85901]++
														return e.RandomToKey(KDF_HMAC_SHA2(protocolKey, label, context, kl, e))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:58
	// _ = "end of CoverTab[85901]"
}

// RandomToKey returns a key from the bytes provided according to the definition in RFC 8009.
func RandomToKey(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:62
	_go_fuzz_dep_.CoverTab[85918]++
														return b
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:63
	// _ = "end of CoverTab[85918]"
}

// StringToKey returns a key derived from the string provided according to the definition in RFC 8009.
func StringToKey(secret, salt, s2kparams string, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:67
	_go_fuzz_dep_.CoverTab[85919]++
														i, err := S2KparamsToItertions(s2kparams)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:69
		_go_fuzz_dep_.CoverTab[85921]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:70
		// _ = "end of CoverTab[85921]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:71
		_go_fuzz_dep_.CoverTab[85922]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:71
		// _ = "end of CoverTab[85922]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:71
	// _ = "end of CoverTab[85919]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:71
	_go_fuzz_dep_.CoverTab[85920]++
														return StringToKeyIter(secret, salt, i, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:72
	// _ = "end of CoverTab[85920]"
}

// StringToKeyIter returns a key derived from the string provided according to the definition in RFC 8009.
func StringToKeyIter(secret, salt string, iterations int, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:76
	_go_fuzz_dep_.CoverTab[85923]++
														tkey := e.RandomToKey(StringToPBKDF2(secret, salt, iterations, e))
														return e.DeriveKey(tkey, []byte("kerberos"))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:78
	// _ = "end of CoverTab[85923]"
}

// StringToPBKDF2 generates an encryption key from a pass phrase and salt string using the PBKDF2 function from PKCS #5 v2.0
func StringToPBKDF2(secret, salt string, iterations int, e etype.EType) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:82
	_go_fuzz_dep_.CoverTab[85924]++
														kl := e.GetKeyByteSize()
														if e.GetETypeID() == etypeID.AES256_CTS_HMAC_SHA384_192 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:84
		_go_fuzz_dep_.CoverTab[85926]++
															kl = 32
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:85
		// _ = "end of CoverTab[85926]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:86
		_go_fuzz_dep_.CoverTab[85927]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:86
		// _ = "end of CoverTab[85927]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:86
	// _ = "end of CoverTab[85924]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:86
	_go_fuzz_dep_.CoverTab[85925]++
														return pbkdf2.Key([]byte(secret), []byte(salt), iterations, kl, e.GetHashFunc())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:87
	// _ = "end of CoverTab[85925]"
}

// KDF_HMAC_SHA2 key derivation: https://tools.ietf.org/html/rfc8009#section-3
func KDF_HMAC_SHA2(protocolKey, label, context []byte, kl int, e etype.EType) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:91
	_go_fuzz_dep_.CoverTab[85928]++

														k := make([]byte, 4, 4)
														binary.BigEndian.PutUint32(k, uint32(kl))

														c := make([]byte, 4, 4)
														binary.BigEndian.PutUint32(c, uint32(1))
														c = append(c, label...)
														c = append(c, byte(0))
														if len(context) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:100
		_go_fuzz_dep_.CoverTab[85930]++
															c = append(c, context...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:101
		// _ = "end of CoverTab[85930]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:102
		_go_fuzz_dep_.CoverTab[85931]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:102
		// _ = "end of CoverTab[85931]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:102
	// _ = "end of CoverTab[85928]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:102
	_go_fuzz_dep_.CoverTab[85929]++
														c = append(c, k...)

														mac := hmac.New(e.GetHashFunc(), protocolKey)
														mac.Write(c)
														return mac.Sum(nil)[:(kl / 8)]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:107
	// _ = "end of CoverTab[85929]"
}

// GetSaltP returns the salt value based on the etype name: https://tools.ietf.org/html/rfc8009#section-4
func GetSaltP(salt, ename string) string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:111
	_go_fuzz_dep_.CoverTab[85932]++
														b := []byte(ename)
														b = append(b, byte(0))
														b = append(b, []byte(salt)...)
														return string(b)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:115
	// _ = "end of CoverTab[85932]"
}

// S2KparamsToItertions converts the string representation of iterations to an integer for RFC 8009.
func S2KparamsToItertions(s2kparams string) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:119
	_go_fuzz_dep_.CoverTab[85933]++
														var i uint32
														if len(s2kparams) != 8 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:121
		_go_fuzz_dep_.CoverTab[85937]++
															return s2kParamsZero, errors.New("Invalid s2kparams length")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:122
		// _ = "end of CoverTab[85937]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:123
		_go_fuzz_dep_.CoverTab[85938]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:123
		// _ = "end of CoverTab[85938]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:123
	// _ = "end of CoverTab[85933]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:123
	_go_fuzz_dep_.CoverTab[85934]++
														b, err := hex.DecodeString(s2kparams)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:125
		_go_fuzz_dep_.CoverTab[85939]++
															return s2kParamsZero, errors.New("Invalid s2kparams, cannot decode string to bytes")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:126
		// _ = "end of CoverTab[85939]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:127
		_go_fuzz_dep_.CoverTab[85940]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:127
		// _ = "end of CoverTab[85940]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:127
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:127
	// _ = "end of CoverTab[85934]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:127
	_go_fuzz_dep_.CoverTab[85935]++
														i = binary.BigEndian.Uint32(b)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:131
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:131
		_go_fuzz_dep_.CoverTab[85941]++
															return s2kParamsZero, errors.New("Invalid s2kparams, cannot convert to big endian int32")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:132
		// _ = "end of CoverTab[85941]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:133
		_go_fuzz_dep_.CoverTab[85942]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:133
		// _ = "end of CoverTab[85942]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:133
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:133
	// _ = "end of CoverTab[85935]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:133
	_go_fuzz_dep_.CoverTab[85936]++
														return int(i), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:134
	// _ = "end of CoverTab[85936]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:135
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/keyDerivation.go:135
var _ = _go_fuzz_dep_.CoverTab
