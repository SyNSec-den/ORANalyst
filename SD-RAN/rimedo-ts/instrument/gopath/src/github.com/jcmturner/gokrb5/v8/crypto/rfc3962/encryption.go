//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:1
// Package rfc3962 provides encryption and checksum methods as specified in RFC 3962
package rfc3962

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:2
)

import (
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/jcmturner/aescts/v2"
	"github.com/jcmturner/gokrb5/v8/crypto/common"
	"github.com/jcmturner/gokrb5/v8/crypto/etype"
)

// EncryptData encrypts the data provided using methods specific to the etype provided as defined in RFC 3962.
func EncryptData(key, data []byte, e etype.EType) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:15
	_go_fuzz_dep_.CoverTab[85709]++
														if len(key) != e.GetKeyByteSize() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:16
		_go_fuzz_dep_.CoverTab[85711]++
															return []byte{}, []byte{}, fmt.Errorf("incorrect keysize: expected: %v actual: %v", e.GetKeyByteSize(), len(key))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:17
		// _ = "end of CoverTab[85711]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:18
		_go_fuzz_dep_.CoverTab[85712]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:18
		// _ = "end of CoverTab[85712]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:18
	// _ = "end of CoverTab[85709]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:18
	_go_fuzz_dep_.CoverTab[85710]++
														ivz := make([]byte, e.GetCypherBlockBitLength()/8)
														return aescts.Encrypt(key, ivz, data)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:20
	// _ = "end of CoverTab[85710]"
}

// EncryptMessage encrypts the message provided using the methods specific to the etype provided as defined in RFC 3962.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:23
// The encrypted data is concatenated with its integrity hash to create an encrypted message.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:25
func EncryptMessage(key, message []byte, usage uint32, e etype.EType) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:25
	_go_fuzz_dep_.CoverTab[85713]++
														if len(key) != e.GetKeyByteSize() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:26
		_go_fuzz_dep_.CoverTab[85719]++
															return []byte{}, []byte{}, fmt.Errorf("incorrect keysize: expected: %v actual: %v", e.GetKeyByteSize(), len(key))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:27
		// _ = "end of CoverTab[85719]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:28
		_go_fuzz_dep_.CoverTab[85720]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:28
		// _ = "end of CoverTab[85720]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:28
	// _ = "end of CoverTab[85713]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:28
	_go_fuzz_dep_.CoverTab[85714]++

														c := make([]byte, e.GetConfounderByteSize())
														_, err := rand.Read(c)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:32
		_go_fuzz_dep_.CoverTab[85721]++
															return []byte{}, []byte{}, fmt.Errorf("could not generate random confounder: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:33
		// _ = "end of CoverTab[85721]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:34
		_go_fuzz_dep_.CoverTab[85722]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:34
		// _ = "end of CoverTab[85722]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:34
	// _ = "end of CoverTab[85714]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:34
	_go_fuzz_dep_.CoverTab[85715]++
														plainBytes := append(c, message...)

	// Derive key for encryption from usage
	var k []byte
	if usage != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:39
		_go_fuzz_dep_.CoverTab[85723]++
															k, err = e.DeriveKey(key, common.GetUsageKe(usage))
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:41
			_go_fuzz_dep_.CoverTab[85724]++
																return []byte{}, []byte{}, fmt.Errorf("error deriving key for encryption: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:42
			// _ = "end of CoverTab[85724]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:43
			_go_fuzz_dep_.CoverTab[85725]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:43
			// _ = "end of CoverTab[85725]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:43
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:43
		// _ = "end of CoverTab[85723]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:44
		_go_fuzz_dep_.CoverTab[85726]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:44
		// _ = "end of CoverTab[85726]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:44
	// _ = "end of CoverTab[85715]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:44
	_go_fuzz_dep_.CoverTab[85716]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:47
	iv, b, err := e.EncryptData(k, plainBytes)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:48
		_go_fuzz_dep_.CoverTab[85727]++
															return iv, b, fmt.Errorf("error encrypting data: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:49
		// _ = "end of CoverTab[85727]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:50
		_go_fuzz_dep_.CoverTab[85728]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:50
		// _ = "end of CoverTab[85728]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:50
	// _ = "end of CoverTab[85716]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:50
	_go_fuzz_dep_.CoverTab[85717]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:53
	ih, err := common.GetIntegrityHash(plainBytes, key, usage, e)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:54
		_go_fuzz_dep_.CoverTab[85729]++
															return iv, b, fmt.Errorf("error encrypting data: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:55
		// _ = "end of CoverTab[85729]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:56
		_go_fuzz_dep_.CoverTab[85730]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:56
		// _ = "end of CoverTab[85730]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:56
	// _ = "end of CoverTab[85717]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:56
	_go_fuzz_dep_.CoverTab[85718]++
														b = append(b, ih...)
														return iv, b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:58
	// _ = "end of CoverTab[85718]"
}

// DecryptData decrypts the data provided using the methods specific to the etype provided as defined in RFC 3962.
func DecryptData(key, data []byte, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:62
	_go_fuzz_dep_.CoverTab[85731]++
														if len(key) != e.GetKeyByteSize() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:63
		_go_fuzz_dep_.CoverTab[85733]++
															return []byte{}, fmt.Errorf("incorrect keysize: expected: %v actual: %v", e.GetKeyByteSize(), len(key))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:64
		// _ = "end of CoverTab[85733]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:65
		_go_fuzz_dep_.CoverTab[85734]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:65
		// _ = "end of CoverTab[85734]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:65
	// _ = "end of CoverTab[85731]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:65
	_go_fuzz_dep_.CoverTab[85732]++
														ivz := make([]byte, e.GetCypherBlockBitLength()/8)
														return aescts.Decrypt(key, ivz, data)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:67
	// _ = "end of CoverTab[85732]"
}

// DecryptMessage decrypts the message provided using the methods specific to the etype provided as defined in RFC 3962.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:70
// The integrity of the message is also verified.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:72
func DecryptMessage(key, ciphertext []byte, usage uint32, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:72
	_go_fuzz_dep_.CoverTab[85735]++

														k, err := e.DeriveKey(key, common.GetUsageKe(usage))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:75
		_go_fuzz_dep_.CoverTab[85739]++
															return nil, fmt.Errorf("error deriving key: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:76
		// _ = "end of CoverTab[85739]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:77
		_go_fuzz_dep_.CoverTab[85740]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:77
		// _ = "end of CoverTab[85740]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:77
	// _ = "end of CoverTab[85735]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:77
	_go_fuzz_dep_.CoverTab[85736]++

														b, err := e.DecryptData(k, ciphertext[:len(ciphertext)-e.GetHMACBitLength()/8])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:80
		_go_fuzz_dep_.CoverTab[85741]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:81
		// _ = "end of CoverTab[85741]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:82
		_go_fuzz_dep_.CoverTab[85742]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:82
		// _ = "end of CoverTab[85742]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:82
	// _ = "end of CoverTab[85736]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:82
	_go_fuzz_dep_.CoverTab[85737]++

														if !e.VerifyIntegrity(key, ciphertext, b, usage) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:84
		_go_fuzz_dep_.CoverTab[85743]++
															return nil, errors.New("integrity verification failed")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:85
		// _ = "end of CoverTab[85743]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:86
		_go_fuzz_dep_.CoverTab[85744]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:86
		// _ = "end of CoverTab[85744]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:86
	// _ = "end of CoverTab[85737]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:86
	_go_fuzz_dep_.CoverTab[85738]++

														return b[e.GetConfounderByteSize():], nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:88
	// _ = "end of CoverTab[85738]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:89
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3962/encryption.go:89
var _ = _go_fuzz_dep_.CoverTab
