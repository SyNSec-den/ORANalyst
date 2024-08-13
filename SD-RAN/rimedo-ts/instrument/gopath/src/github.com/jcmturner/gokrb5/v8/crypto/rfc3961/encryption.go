//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:1
// Package rfc3961 provides encryption and checksum methods as specified in RFC 3961
package rfc3961

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:2
)

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/jcmturner/gokrb5/v8/crypto/common"
	"github.com/jcmturner/gokrb5/v8/crypto/etype"
)

// DES3EncryptData encrypts the data provided using DES3 and methods specific to the etype provided.
func DES3EncryptData(key, data []byte, e etype.EType) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:17
	_go_fuzz_dep_.CoverTab[85524]++
														if len(key) != e.GetKeyByteSize() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:18
		_go_fuzz_dep_.CoverTab[85527]++
															return nil, nil, fmt.Errorf("incorrect keysize: expected: %v actual: %v", e.GetKeyByteSize(), len(key))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:19
		// _ = "end of CoverTab[85527]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:20
		_go_fuzz_dep_.CoverTab[85528]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:20
		// _ = "end of CoverTab[85528]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:20
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:20
	// _ = "end of CoverTab[85524]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:20
	_go_fuzz_dep_.CoverTab[85525]++
														data, _ = common.ZeroPad(data, e.GetMessageBlockByteSize())

														block, err := des.NewTripleDESCipher(key)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:24
		_go_fuzz_dep_.CoverTab[85529]++
															return nil, nil, fmt.Errorf("error creating cipher: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:25
		// _ = "end of CoverTab[85529]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:26
		_go_fuzz_dep_.CoverTab[85530]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:26
		// _ = "end of CoverTab[85530]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:26
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:26
	// _ = "end of CoverTab[85525]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:26
	_go_fuzz_dep_.CoverTab[85526]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:29
	ivz := make([]byte, des.BlockSize)

														ct := make([]byte, len(data))
														mode := cipher.NewCBCEncrypter(block, ivz)
														mode.CryptBlocks(ct, data)
														return ct[len(ct)-e.GetMessageBlockByteSize():], ct, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:34
	// _ = "end of CoverTab[85526]"
}

// DES3EncryptMessage encrypts the message provided using DES3 and methods specific to the etype provided.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:37
// The encrypted data is concatenated with its integrity hash to create an encrypted message.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:39
func DES3EncryptMessage(key, message []byte, usage uint32, e etype.EType) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:39
	_go_fuzz_dep_.CoverTab[85531]++

														c := make([]byte, e.GetConfounderByteSize())
														_, err := rand.Read(c)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:43
		_go_fuzz_dep_.CoverTab[85536]++
															return []byte{}, []byte{}, fmt.Errorf("could not generate random confounder: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:44
		// _ = "end of CoverTab[85536]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:45
		_go_fuzz_dep_.CoverTab[85537]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:45
		// _ = "end of CoverTab[85537]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:45
	// _ = "end of CoverTab[85531]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:45
	_go_fuzz_dep_.CoverTab[85532]++
														plainBytes := append(c, message...)
														plainBytes, _ = common.ZeroPad(plainBytes, e.GetMessageBlockByteSize())

	// Derive key for encryption from usage
	var k []byte
	if usage != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:51
		_go_fuzz_dep_.CoverTab[85538]++
															k, err = e.DeriveKey(key, common.GetUsageKe(usage))
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:53
			_go_fuzz_dep_.CoverTab[85539]++
																return []byte{}, []byte{}, fmt.Errorf("error deriving key for encryption: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:54
			// _ = "end of CoverTab[85539]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:55
			_go_fuzz_dep_.CoverTab[85540]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:55
			// _ = "end of CoverTab[85540]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:55
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:55
		// _ = "end of CoverTab[85538]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:56
		_go_fuzz_dep_.CoverTab[85541]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:56
		// _ = "end of CoverTab[85541]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:56
	// _ = "end of CoverTab[85532]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:56
	_go_fuzz_dep_.CoverTab[85533]++

														iv, b, err := e.EncryptData(k, plainBytes)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:59
		_go_fuzz_dep_.CoverTab[85542]++
															return iv, b, fmt.Errorf("error encrypting data: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:60
		// _ = "end of CoverTab[85542]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:61
		_go_fuzz_dep_.CoverTab[85543]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:61
		// _ = "end of CoverTab[85543]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:61
	// _ = "end of CoverTab[85533]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:61
	_go_fuzz_dep_.CoverTab[85534]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:64
	ih, err := common.GetIntegrityHash(plainBytes, key, usage, e)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:65
		_go_fuzz_dep_.CoverTab[85544]++
															return iv, b, fmt.Errorf("error encrypting data: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:66
		// _ = "end of CoverTab[85544]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:67
		_go_fuzz_dep_.CoverTab[85545]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:67
		// _ = "end of CoverTab[85545]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:67
	// _ = "end of CoverTab[85534]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:67
	_go_fuzz_dep_.CoverTab[85535]++
														b = append(b, ih...)
														return iv, b, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:69
	// _ = "end of CoverTab[85535]"
}

// DES3DecryptData decrypts the data provided using DES3 and methods specific to the etype provided.
func DES3DecryptData(key, data []byte, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:73
	_go_fuzz_dep_.CoverTab[85546]++
														if len(key) != e.GetKeyByteSize() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:74
		_go_fuzz_dep_.CoverTab[85550]++
															return []byte{}, fmt.Errorf("incorrect keysize: expected: %v actual: %v", e.GetKeyByteSize(), len(key))
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:75
		// _ = "end of CoverTab[85550]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:76
		_go_fuzz_dep_.CoverTab[85551]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:76
		// _ = "end of CoverTab[85551]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:76
	// _ = "end of CoverTab[85546]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:76
	_go_fuzz_dep_.CoverTab[85547]++

														if len(data) < des.BlockSize || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:78
		_go_fuzz_dep_.CoverTab[85552]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:78
		return len(data)%des.BlockSize != 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:78
		// _ = "end of CoverTab[85552]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:78
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:78
		_go_fuzz_dep_.CoverTab[85553]++
															return []byte{}, errors.New("ciphertext is not a multiple of the block size")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:79
		// _ = "end of CoverTab[85553]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:80
		_go_fuzz_dep_.CoverTab[85554]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:80
		// _ = "end of CoverTab[85554]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:80
	// _ = "end of CoverTab[85547]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:80
	_go_fuzz_dep_.CoverTab[85548]++
														block, err := des.NewTripleDESCipher(key)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:82
		_go_fuzz_dep_.CoverTab[85555]++
															return []byte{}, fmt.Errorf("error creating cipher: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:83
		// _ = "end of CoverTab[85555]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:84
		_go_fuzz_dep_.CoverTab[85556]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:84
		// _ = "end of CoverTab[85556]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:84
	// _ = "end of CoverTab[85548]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:84
	_go_fuzz_dep_.CoverTab[85549]++
														pt := make([]byte, len(data))
														ivz := make([]byte, des.BlockSize)
														mode := cipher.NewCBCDecrypter(block, ivz)
														mode.CryptBlocks(pt, data)
														return pt, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:89
	// _ = "end of CoverTab[85549]"
}

// DES3DecryptMessage decrypts the message provided using DES3 and methods specific to the etype provided.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:92
// The integrity of the message is also verified.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:94
func DES3DecryptMessage(key, ciphertext []byte, usage uint32, e etype.EType) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:94
	_go_fuzz_dep_.CoverTab[85557]++

														k, err := e.DeriveKey(key, common.GetUsageKe(usage))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:97
		_go_fuzz_dep_.CoverTab[85561]++
															return nil, fmt.Errorf("error deriving key: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:98
		// _ = "end of CoverTab[85561]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:99
		_go_fuzz_dep_.CoverTab[85562]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:99
		// _ = "end of CoverTab[85562]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:99
	// _ = "end of CoverTab[85557]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:99
	_go_fuzz_dep_.CoverTab[85558]++

														b, err := e.DecryptData(k, ciphertext[:len(ciphertext)-e.GetHMACBitLength()/8])
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:102
		_go_fuzz_dep_.CoverTab[85563]++
															return nil, fmt.Errorf("error decrypting: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:103
		// _ = "end of CoverTab[85563]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:104
		_go_fuzz_dep_.CoverTab[85564]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:104
		// _ = "end of CoverTab[85564]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:104
	// _ = "end of CoverTab[85558]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:104
	_go_fuzz_dep_.CoverTab[85559]++

														if !e.VerifyIntegrity(key, ciphertext, b, usage) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:106
		_go_fuzz_dep_.CoverTab[85565]++
															return nil, errors.New("error decrypting: integrity verification failed")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:107
		// _ = "end of CoverTab[85565]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:108
		_go_fuzz_dep_.CoverTab[85566]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:108
		// _ = "end of CoverTab[85566]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:108
	// _ = "end of CoverTab[85559]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:108
	_go_fuzz_dep_.CoverTab[85560]++

														return b[e.GetConfounderByteSize():], nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:110
	// _ = "end of CoverTab[85560]"
}

// VerifyIntegrity verifies the integrity of cipertext bytes ct.
func VerifyIntegrity(key, ct, pt []byte, usage uint32, etype etype.EType) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:114
	_go_fuzz_dep_.CoverTab[85567]++
														h := make([]byte, etype.GetHMACBitLength()/8)
														copy(h, ct[len(ct)-etype.GetHMACBitLength()/8:])
														expectedMAC, _ := common.GetIntegrityHash(pt, key, usage, etype)
														return hmac.Equal(h, expectedMAC)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:118
	// _ = "end of CoverTab[85567]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:119
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/encryption.go:119
var _ = _go_fuzz_dep_.CoverTab
