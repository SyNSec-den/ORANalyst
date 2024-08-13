//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:1
// Package rfc4757 provides encryption and checksum methods as specified in RFC 4757
package rfc4757

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:2
)

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/rc4"
	"errors"
	"fmt"

	"github.com/jcmturner/gokrb5/v8/crypto/etype"
)

// EncryptData encrypts the data provided using methods specific to the etype provided as defined in RFC 4757.
func EncryptData(key, data []byte, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:15
	_go_fuzz_dep_.CoverTab[85804]++
														if len(key) != e.GetKeyByteSize() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:16
		_go_fuzz_dep_.CoverTab[85807]++
															return []byte{}, fmt.Errorf("incorrect keysize: expected: %v actual: %v", e.GetKeyByteSize(), len(key))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:17
		// _ = "end of CoverTab[85807]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:18
		_go_fuzz_dep_.CoverTab[85808]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:18
		// _ = "end of CoverTab[85808]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:18
	// _ = "end of CoverTab[85804]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:18
	_go_fuzz_dep_.CoverTab[85805]++
														rc4Cipher, err := rc4.NewCipher(key)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:20
		_go_fuzz_dep_.CoverTab[85809]++
															return []byte{}, fmt.Errorf("error creating RC4 cipher: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:21
		// _ = "end of CoverTab[85809]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:22
		_go_fuzz_dep_.CoverTab[85810]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:22
		// _ = "end of CoverTab[85810]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:22
	// _ = "end of CoverTab[85805]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:22
	_go_fuzz_dep_.CoverTab[85806]++
														ed := make([]byte, len(data))
														copy(ed, data)
														rc4Cipher.XORKeyStream(ed, ed)
														rc4Cipher.Reset()
														return ed, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:27
	// _ = "end of CoverTab[85806]"
}

// DecryptData decrypts the data provided using the methods specific to the etype provided as defined in RFC 4757.
func DecryptData(key, data []byte, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:31
	_go_fuzz_dep_.CoverTab[85811]++
														return EncryptData(key, data, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:32
	// _ = "end of CoverTab[85811]"
}

// EncryptMessage encrypts the message provided using the methods specific to the etype provided as defined in RFC 4757.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:35
// The encrypted data is concatenated with its RC4 header containing integrity checksum and confounder to create an encrypted message.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:37
func EncryptMessage(key, data []byte, usage uint32, export bool, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:37
	_go_fuzz_dep_.CoverTab[85812]++
														confounder := make([]byte, e.GetConfounderByteSize())
														_, err := rand.Read(confounder)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:40
		_go_fuzz_dep_.CoverTab[85815]++
															return []byte{}, fmt.Errorf("error generating confounder: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:41
		// _ = "end of CoverTab[85815]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:42
		_go_fuzz_dep_.CoverTab[85816]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:42
		// _ = "end of CoverTab[85816]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:42
	// _ = "end of CoverTab[85812]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:42
	_go_fuzz_dep_.CoverTab[85813]++
														k1 := key
														k2 := HMAC(k1, UsageToMSMsgType(usage))
														toenc := append(confounder, data...)
														chksum := HMAC(k2, toenc)
														k3 := HMAC(k2, chksum)

														ed, err := EncryptData(k3, toenc, e)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:50
		_go_fuzz_dep_.CoverTab[85817]++
															return []byte{}, fmt.Errorf("error encrypting data: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:51
		// _ = "end of CoverTab[85817]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:52
		_go_fuzz_dep_.CoverTab[85818]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:52
		// _ = "end of CoverTab[85818]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:52
	// _ = "end of CoverTab[85813]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:52
	_go_fuzz_dep_.CoverTab[85814]++

														msg := append(chksum, ed...)
														return msg, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:55
	// _ = "end of CoverTab[85814]"
}

// DecryptMessage decrypts the message provided using the methods specific to the etype provided as defined in RFC 4757.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:58
// The integrity of the message is also verified.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:60
func DecryptMessage(key, data []byte, usage uint32, export bool, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:60
	_go_fuzz_dep_.CoverTab[85819]++
														checksum := data[:e.GetHMACBitLength()/8]
														ct := data[e.GetHMACBitLength()/8:]
														_, k2, k3 := deriveKeys(key, checksum, usage, export)

														pt, err := DecryptData(k3, ct, e)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:66
		_go_fuzz_dep_.CoverTab[85822]++
															return []byte{}, fmt.Errorf("error decrypting data: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:67
		// _ = "end of CoverTab[85822]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:68
		_go_fuzz_dep_.CoverTab[85823]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:68
		// _ = "end of CoverTab[85823]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:68
	// _ = "end of CoverTab[85819]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:68
	_go_fuzz_dep_.CoverTab[85820]++

														if !VerifyIntegrity(k2, pt, data, e) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:70
		_go_fuzz_dep_.CoverTab[85824]++
															return []byte{}, errors.New("integrity checksum incorrect")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:71
		// _ = "end of CoverTab[85824]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:72
		_go_fuzz_dep_.CoverTab[85825]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:72
		// _ = "end of CoverTab[85825]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:72
	// _ = "end of CoverTab[85820]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:72
	_go_fuzz_dep_.CoverTab[85821]++
														return pt[e.GetConfounderByteSize():], nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:73
	// _ = "end of CoverTab[85821]"
}

// VerifyIntegrity checks the integrity checksum of the data matches that calculated from the decrypted data.
func VerifyIntegrity(key, pt, data []byte, e etype.EType) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:77
	_go_fuzz_dep_.CoverTab[85826]++
														chksum := HMAC(key, pt)
														return hmac.Equal(chksum, data[:e.GetHMACBitLength()/8])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:79
	// _ = "end of CoverTab[85826]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:80
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc4757/encryption.go:80
var _ = _go_fuzz_dep_.CoverTab
