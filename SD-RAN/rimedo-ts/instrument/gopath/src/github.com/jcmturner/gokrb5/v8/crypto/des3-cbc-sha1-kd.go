//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:1
package crypto

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:1
)

import (
	"crypto/des"
	"crypto/hmac"
	"crypto/sha1"
	"errors"
	"hash"

	"github.com/jcmturner/gokrb5/v8/crypto/common"
	"github.com/jcmturner/gokrb5/v8/crypto/rfc3961"
	"github.com/jcmturner/gokrb5/v8/iana/chksumtype"
	"github.com/jcmturner/gokrb5/v8/iana/etypeID"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:18
// Des3CbcSha1Kd implements Kerberos encryption type des3-cbc-hmac-sha1-kd
type Des3CbcSha1Kd struct {
}

// GetETypeID returns the EType ID number.
func (e Des3CbcSha1Kd) GetETypeID() int32 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:23
	_go_fuzz_dep_.CoverTab[86274]++
													return etypeID.DES3_CBC_SHA1_KD
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:24
	// _ = "end of CoverTab[86274]"
}

// GetHashID returns the checksum type ID number.
func (e Des3CbcSha1Kd) GetHashID() int32 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:28
	_go_fuzz_dep_.CoverTab[86275]++
													return chksumtype.HMAC_SHA1_DES3_KD
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:29
	// _ = "end of CoverTab[86275]"
}

// GetKeyByteSize returns the number of bytes for key of this etype.
func (e Des3CbcSha1Kd) GetKeyByteSize() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:33
	_go_fuzz_dep_.CoverTab[86276]++
													return 24
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:34
	// _ = "end of CoverTab[86276]"
}

// GetKeySeedBitLength returns the number of bits for the seed for key generation.
func (e Des3CbcSha1Kd) GetKeySeedBitLength() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:38
	_go_fuzz_dep_.CoverTab[86277]++
													return 21 * 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:39
	// _ = "end of CoverTab[86277]"
}

// GetHashFunc returns the hash function for this etype.
func (e Des3CbcSha1Kd) GetHashFunc() func() hash.Hash {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:43
	_go_fuzz_dep_.CoverTab[86278]++
													return sha1.New
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:44
	// _ = "end of CoverTab[86278]"
}

// GetMessageBlockByteSize returns the block size for the etype's messages.
func (e Des3CbcSha1Kd) GetMessageBlockByteSize() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:48
	_go_fuzz_dep_.CoverTab[86279]++

													return des.BlockSize
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:50
	// _ = "end of CoverTab[86279]"
}

// GetDefaultStringToKeyParams returns the default key derivation parameters in string form.
func (e Des3CbcSha1Kd) GetDefaultStringToKeyParams() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:54
	_go_fuzz_dep_.CoverTab[86280]++
													var s string
													return s
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:56
	// _ = "end of CoverTab[86280]"
}

// GetConfounderByteSize returns the byte count for confounder to be used during cryptographic operations.
func (e Des3CbcSha1Kd) GetConfounderByteSize() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:60
	_go_fuzz_dep_.CoverTab[86281]++
													return des.BlockSize
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:61
	// _ = "end of CoverTab[86281]"
}

// GetHMACBitLength returns the bit count size of the integrity hash.
func (e Des3CbcSha1Kd) GetHMACBitLength() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:65
	_go_fuzz_dep_.CoverTab[86282]++
													return e.GetHashFunc()().Size() * 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:66
	// _ = "end of CoverTab[86282]"
}

// GetCypherBlockBitLength returns the bit count size of the cypher block.
func (e Des3CbcSha1Kd) GetCypherBlockBitLength() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:70
	_go_fuzz_dep_.CoverTab[86283]++
													return des.BlockSize * 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:71
	// _ = "end of CoverTab[86283]"
}

// StringToKey returns a key derived from the string provided.
func (e Des3CbcSha1Kd) StringToKey(secret string, salt string, s2kparams string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:75
	_go_fuzz_dep_.CoverTab[86284]++
													if s2kparams != "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:76
		_go_fuzz_dep_.CoverTab[86286]++
														return []byte{}, errors.New("s2kparams must be an empty string")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:77
		// _ = "end of CoverTab[86286]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:78
		_go_fuzz_dep_.CoverTab[86287]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:78
		// _ = "end of CoverTab[86287]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:78
	// _ = "end of CoverTab[86284]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:78
	_go_fuzz_dep_.CoverTab[86285]++
													return rfc3961.DES3StringToKey(secret, salt, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:79
	// _ = "end of CoverTab[86285]"
}

// RandomToKey returns a key from the bytes provided.
func (e Des3CbcSha1Kd) RandomToKey(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:83
	_go_fuzz_dep_.CoverTab[86288]++
													return rfc3961.DES3RandomToKey(b)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:84
	// _ = "end of CoverTab[86288]"
}

// DeriveRandom generates data needed for key generation.
func (e Des3CbcSha1Kd) DeriveRandom(protocolKey, usage []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:88
	_go_fuzz_dep_.CoverTab[86289]++
													r, err := rfc3961.DeriveRandom(protocolKey, usage, e)
													return r, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:90
	// _ = "end of CoverTab[86289]"
}

// DeriveKey derives a key from the protocol key based on the usage value.
func (e Des3CbcSha1Kd) DeriveKey(protocolKey, usage []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:94
	_go_fuzz_dep_.CoverTab[86290]++
													r, err := e.DeriveRandom(protocolKey, usage)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:96
		_go_fuzz_dep_.CoverTab[86292]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:97
		// _ = "end of CoverTab[86292]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:98
		_go_fuzz_dep_.CoverTab[86293]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:98
		// _ = "end of CoverTab[86293]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:98
	// _ = "end of CoverTab[86290]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:98
	_go_fuzz_dep_.CoverTab[86291]++
													return e.RandomToKey(r), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:99
	// _ = "end of CoverTab[86291]"
}

// EncryptData encrypts the data provided.
func (e Des3CbcSha1Kd) EncryptData(key, data []byte) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:103
	_go_fuzz_dep_.CoverTab[86294]++
													return rfc3961.DES3EncryptData(key, data, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:104
	// _ = "end of CoverTab[86294]"
}

// EncryptMessage encrypts the message provided and concatenates it with the integrity hash to create an encrypted message.
func (e Des3CbcSha1Kd) EncryptMessage(key, message []byte, usage uint32) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:108
	_go_fuzz_dep_.CoverTab[86295]++
													return rfc3961.DES3EncryptMessage(key, message, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:109
	// _ = "end of CoverTab[86295]"
}

// DecryptData decrypts the data provided.
func (e Des3CbcSha1Kd) DecryptData(key, data []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:113
	_go_fuzz_dep_.CoverTab[86296]++
													return rfc3961.DES3DecryptData(key, data, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:114
	// _ = "end of CoverTab[86296]"
}

// DecryptMessage decrypts the message provided and verifies the integrity of the message.
func (e Des3CbcSha1Kd) DecryptMessage(key, ciphertext []byte, usage uint32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:118
	_go_fuzz_dep_.CoverTab[86297]++
													return rfc3961.DES3DecryptMessage(key, ciphertext, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:119
	// _ = "end of CoverTab[86297]"
}

// VerifyIntegrity checks the integrity of the plaintext message.
func (e Des3CbcSha1Kd) VerifyIntegrity(protocolKey, ct, pt []byte, usage uint32) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:123
	_go_fuzz_dep_.CoverTab[86298]++
													return rfc3961.VerifyIntegrity(protocolKey, ct, pt, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:124
	// _ = "end of CoverTab[86298]"
}

// GetChecksumHash returns a keyed checksum hash of the bytes provided.
func (e Des3CbcSha1Kd) GetChecksumHash(protocolKey, data []byte, usage uint32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:128
	_go_fuzz_dep_.CoverTab[86299]++
													return common.GetHash(data, protocolKey, common.GetUsageKc(usage), e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:129
	// _ = "end of CoverTab[86299]"
}

// VerifyChecksum compares the checksum of the message bytes is the same as the checksum provided.
func (e Des3CbcSha1Kd) VerifyChecksum(protocolKey, data, chksum []byte, usage uint32) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:133
	_go_fuzz_dep_.CoverTab[86300]++
													c, err := e.GetChecksumHash(protocolKey, data, usage)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:135
		_go_fuzz_dep_.CoverTab[86302]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:136
		// _ = "end of CoverTab[86302]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:137
		_go_fuzz_dep_.CoverTab[86303]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:137
		// _ = "end of CoverTab[86303]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:137
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:137
	// _ = "end of CoverTab[86300]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:137
	_go_fuzz_dep_.CoverTab[86301]++
													return hmac.Equal(chksum, c)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:138
	// _ = "end of CoverTab[86301]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:139
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/des3-cbc-sha1-kd.go:139
var _ = _go_fuzz_dep_.CoverTab
