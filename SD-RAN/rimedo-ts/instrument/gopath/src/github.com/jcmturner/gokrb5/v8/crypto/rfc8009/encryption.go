//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:1
// Package rfc8009 provides encryption and checksum methods as specified in RFC 8009
package rfc8009

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:2
)

import (
	"crypto/aes"
	"crypto/hmac"
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/jcmturner/aescts/v2"
	"github.com/jcmturner/gokrb5/v8/crypto/common"
	"github.com/jcmturner/gokrb5/v8/crypto/etype"
	"github.com/jcmturner/gokrb5/v8/iana/etypeID"
)

// EncryptData encrypts the data provided using methods specific to the etype provided as defined in RFC 8009.
func EncryptData(key, data []byte, e etype.EType) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:18
	_go_fuzz_dep_.CoverTab[85848]++
														kl := e.GetKeyByteSize()
														if e.GetETypeID() == etypeID.AES256_CTS_HMAC_SHA384_192 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:20
		_go_fuzz_dep_.CoverTab[85851]++
															kl = 32
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:21
		// _ = "end of CoverTab[85851]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:22
		_go_fuzz_dep_.CoverTab[85852]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:22
		// _ = "end of CoverTab[85852]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:22
	// _ = "end of CoverTab[85848]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:22
	_go_fuzz_dep_.CoverTab[85849]++
														if len(key) != kl {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:23
		_go_fuzz_dep_.CoverTab[85853]++
															return []byte{}, []byte{}, fmt.Errorf("incorrect keysize: expected: %v actual: %v", e.GetKeyByteSize(), len(key))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:24
		// _ = "end of CoverTab[85853]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:25
		_go_fuzz_dep_.CoverTab[85854]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:25
		// _ = "end of CoverTab[85854]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:25
	// _ = "end of CoverTab[85849]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:25
	_go_fuzz_dep_.CoverTab[85850]++
														ivz := make([]byte, aes.BlockSize)
														return aescts.Encrypt(key, ivz, data)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:27
	// _ = "end of CoverTab[85850]"
}

// EncryptMessage encrypts the message provided using the methods specific to the etype provided as defined in RFC 8009.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:30
// The encrypted data is concatenated with its integrity hash to create an encrypted message.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:32
func EncryptMessage(key, message []byte, usage uint32, e etype.EType) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:32
	_go_fuzz_dep_.CoverTab[85855]++
														kl := e.GetKeyByteSize()
														if e.GetETypeID() == etypeID.AES256_CTS_HMAC_SHA384_192 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:34
		_go_fuzz_dep_.CoverTab[85863]++
															kl = 32
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:35
		// _ = "end of CoverTab[85863]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:36
		_go_fuzz_dep_.CoverTab[85864]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:36
		// _ = "end of CoverTab[85864]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:36
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:36
	// _ = "end of CoverTab[85855]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:36
	_go_fuzz_dep_.CoverTab[85856]++
														if len(key) != kl {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:37
		_go_fuzz_dep_.CoverTab[85865]++
															return []byte{}, []byte{}, fmt.Errorf("incorrect keysize: expected: %v actual: %v", kl, len(key))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:38
		// _ = "end of CoverTab[85865]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:39
		_go_fuzz_dep_.CoverTab[85866]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:39
		// _ = "end of CoverTab[85866]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:39
	// _ = "end of CoverTab[85856]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:39
	_go_fuzz_dep_.CoverTab[85857]++
														if len(key) != e.GetKeyByteSize() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:40
		_go_fuzz_dep_.CoverTab[85867]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:40
		// _ = "end of CoverTab[85867]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:41
		_go_fuzz_dep_.CoverTab[85868]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:41
		// _ = "end of CoverTab[85868]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:41
	// _ = "end of CoverTab[85857]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:41
	_go_fuzz_dep_.CoverTab[85858]++

														c := make([]byte, e.GetConfounderByteSize())
														_, err := rand.Read(c)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:45
		_go_fuzz_dep_.CoverTab[85869]++
															return []byte{}, []byte{}, fmt.Errorf("could not generate random confounder: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:46
		// _ = "end of CoverTab[85869]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:47
		_go_fuzz_dep_.CoverTab[85870]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:47
		// _ = "end of CoverTab[85870]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:47
	// _ = "end of CoverTab[85858]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:47
	_go_fuzz_dep_.CoverTab[85859]++
														plainBytes := append(c, message...)

	// Derive key for encryption from usage
	var k []byte
	if usage != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:52
		_go_fuzz_dep_.CoverTab[85871]++
															k, err = e.DeriveKey(key, common.GetUsageKe(usage))
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:54
			_go_fuzz_dep_.CoverTab[85872]++
																return []byte{}, []byte{}, fmt.Errorf("error deriving key for encryption: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:55
			// _ = "end of CoverTab[85872]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:56
			_go_fuzz_dep_.CoverTab[85873]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:56
			// _ = "end of CoverTab[85873]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:56
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:56
		// _ = "end of CoverTab[85871]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:57
		_go_fuzz_dep_.CoverTab[85874]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:57
		// _ = "end of CoverTab[85874]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:57
	// _ = "end of CoverTab[85859]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:57
	_go_fuzz_dep_.CoverTab[85860]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:60
	iv, b, err := e.EncryptData(k, plainBytes)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:61
		_go_fuzz_dep_.CoverTab[85875]++
															return iv, b, fmt.Errorf("error encrypting data: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:62
		// _ = "end of CoverTab[85875]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:63
		_go_fuzz_dep_.CoverTab[85876]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:63
		// _ = "end of CoverTab[85876]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:63
	// _ = "end of CoverTab[85860]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:63
	_go_fuzz_dep_.CoverTab[85861]++

														ivz := make([]byte, e.GetConfounderByteSize())
														ih, err := GetIntegityHash(ivz, b, key, usage, e)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:67
		_go_fuzz_dep_.CoverTab[85877]++
															return iv, b, fmt.Errorf("error encrypting data: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:68
		// _ = "end of CoverTab[85877]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:69
		_go_fuzz_dep_.CoverTab[85878]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:69
		// _ = "end of CoverTab[85878]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:69
	// _ = "end of CoverTab[85861]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:69
	_go_fuzz_dep_.CoverTab[85862]++
														b = append(b, ih...)
														return iv, b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:71
	// _ = "end of CoverTab[85862]"
}

// DecryptData decrypts the data provided using the methods specific to the etype provided as defined in RFC 8009.
func DecryptData(key, data []byte, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:75
	_go_fuzz_dep_.CoverTab[85879]++
														kl := e.GetKeyByteSize()
														if e.GetETypeID() == etypeID.AES256_CTS_HMAC_SHA384_192 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:77
		_go_fuzz_dep_.CoverTab[85882]++
															kl = 32
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:78
		// _ = "end of CoverTab[85882]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:79
		_go_fuzz_dep_.CoverTab[85883]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:79
		// _ = "end of CoverTab[85883]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:79
	// _ = "end of CoverTab[85879]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:79
	_go_fuzz_dep_.CoverTab[85880]++
														if len(key) != kl {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:80
		_go_fuzz_dep_.CoverTab[85884]++
															return []byte{}, fmt.Errorf("incorrect keysize: expected: %v actual: %v", kl, len(key))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:81
		// _ = "end of CoverTab[85884]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:82
		_go_fuzz_dep_.CoverTab[85885]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:82
		// _ = "end of CoverTab[85885]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:82
	// _ = "end of CoverTab[85880]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:82
	_go_fuzz_dep_.CoverTab[85881]++
														ivz := make([]byte, aes.BlockSize)
														return aescts.Decrypt(key, ivz, data)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:84
	// _ = "end of CoverTab[85881]"
}

// DecryptMessage decrypts the message provided using the methods specific to the etype provided as defined in RFC 8009.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:87
// The integrity of the message is also verified.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:89
func DecryptMessage(key, ciphertext []byte, usage uint32, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:89
	_go_fuzz_dep_.CoverTab[85886]++

														k, err := e.DeriveKey(key, common.GetUsageKe(usage))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:92
		_go_fuzz_dep_.CoverTab[85890]++
															return nil, fmt.Errorf("error deriving key: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:93
		// _ = "end of CoverTab[85890]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:94
		_go_fuzz_dep_.CoverTab[85891]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:94
		// _ = "end of CoverTab[85891]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:94
	// _ = "end of CoverTab[85886]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:94
	_go_fuzz_dep_.CoverTab[85887]++

														b, err := e.DecryptData(k, ciphertext[:len(ciphertext)-e.GetHMACBitLength()/8])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:97
		_go_fuzz_dep_.CoverTab[85892]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:98
		// _ = "end of CoverTab[85892]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:99
		_go_fuzz_dep_.CoverTab[85893]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:99
		// _ = "end of CoverTab[85893]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:99
	// _ = "end of CoverTab[85887]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:99
	_go_fuzz_dep_.CoverTab[85888]++

														if !e.VerifyIntegrity(key, ciphertext, b, usage) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:101
		_go_fuzz_dep_.CoverTab[85894]++
															return nil, errors.New("integrity verification failed")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:102
		// _ = "end of CoverTab[85894]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:103
		_go_fuzz_dep_.CoverTab[85895]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:103
		// _ = "end of CoverTab[85895]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:103
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:103
	// _ = "end of CoverTab[85888]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:103
	_go_fuzz_dep_.CoverTab[85889]++

														return b[e.GetConfounderByteSize():], nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:105
	// _ = "end of CoverTab[85889]"
}

// GetIntegityHash returns a keyed integrity hash of the bytes provided as defined in RFC 8009
func GetIntegityHash(iv, c, key []byte, usage uint32, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:109
	_go_fuzz_dep_.CoverTab[85896]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:113
	ib := append(iv, c...)
														return common.GetIntegrityHash(ib, key, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:114
	// _ = "end of CoverTab[85896]"
}

// VerifyIntegrity verifies the integrity of cipertext bytes ct.
func VerifyIntegrity(key, ct []byte, usage uint32, etype etype.EType) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:118
	_go_fuzz_dep_.CoverTab[85897]++
														h := make([]byte, etype.GetHMACBitLength()/8)
														copy(h, ct[len(ct)-etype.GetHMACBitLength()/8:])
														ivz := make([]byte, etype.GetConfounderByteSize())
														ib := append(ivz, ct[:len(ct)-(etype.GetHMACBitLength()/8)]...)
														expectedMAC, _ := common.GetIntegrityHash(ib, key, usage, etype)
														return hmac.Equal(h, expectedMAC)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:124
	// _ = "end of CoverTab[85897]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:125
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc8009/encryption.go:125
var _ = _go_fuzz_dep_.CoverTab
