//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:1
package crypto

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:1
)

import (
	"crypto/aes"
	"crypto/hmac"
	"crypto/sha256"
	"hash"

	"github.com/jcmturner/gokrb5/v8/crypto/common"
	"github.com/jcmturner/gokrb5/v8/crypto/rfc8009"
	"github.com/jcmturner/gokrb5/v8/iana/chksumtype"
	"github.com/jcmturner/gokrb5/v8/iana/etypeID"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:17
// Aes128CtsHmacSha256128 implements Kerberos encryption type aes128-cts-hmac-sha256-128
type Aes128CtsHmacSha256128 struct {
}

// GetETypeID returns the EType ID number.
func (e Aes128CtsHmacSha256128) GetETypeID() int32 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:22
	_go_fuzz_dep_.CoverTab[86127]++
															return etypeID.AES128_CTS_HMAC_SHA256_128
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:23
	// _ = "end of CoverTab[86127]"
}

// GetHashID returns the checksum type ID number.
func (e Aes128CtsHmacSha256128) GetHashID() int32 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:27
	_go_fuzz_dep_.CoverTab[86128]++
															return chksumtype.HMAC_SHA256_128_AES128
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:28
	// _ = "end of CoverTab[86128]"
}

// GetKeyByteSize returns the number of bytes for key of this etype.
func (e Aes128CtsHmacSha256128) GetKeyByteSize() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:32
	_go_fuzz_dep_.CoverTab[86129]++
															return 128 / 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:33
	// _ = "end of CoverTab[86129]"
}

// GetKeySeedBitLength returns the number of bits for the seed for key generation.
func (e Aes128CtsHmacSha256128) GetKeySeedBitLength() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:37
	_go_fuzz_dep_.CoverTab[86130]++
															return e.GetKeyByteSize() * 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:38
	// _ = "end of CoverTab[86130]"
}

// GetHashFunc returns the hash function for this etype.
func (e Aes128CtsHmacSha256128) GetHashFunc() func() hash.Hash {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:42
	_go_fuzz_dep_.CoverTab[86131]++
															return sha256.New
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:43
	// _ = "end of CoverTab[86131]"
}

// GetMessageBlockByteSize returns the block size for the etype's messages.
func (e Aes128CtsHmacSha256128) GetMessageBlockByteSize() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:47
	_go_fuzz_dep_.CoverTab[86132]++
															return 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:48
	// _ = "end of CoverTab[86132]"
}

// GetDefaultStringToKeyParams returns the default key derivation parameters in string form.
func (e Aes128CtsHmacSha256128) GetDefaultStringToKeyParams() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:52
	_go_fuzz_dep_.CoverTab[86133]++
															return "00008000"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:53
	// _ = "end of CoverTab[86133]"
}

// GetConfounderByteSize returns the byte count for confounder to be used during cryptographic operations.
func (e Aes128CtsHmacSha256128) GetConfounderByteSize() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:57
	_go_fuzz_dep_.CoverTab[86134]++
															return aes.BlockSize
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:58
	// _ = "end of CoverTab[86134]"
}

// GetHMACBitLength returns the bit count size of the integrity hash.
func (e Aes128CtsHmacSha256128) GetHMACBitLength() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:62
	_go_fuzz_dep_.CoverTab[86135]++
															return 128
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:63
	// _ = "end of CoverTab[86135]"
}

// GetCypherBlockBitLength returns the bit count size of the cypher block.
func (e Aes128CtsHmacSha256128) GetCypherBlockBitLength() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:67
	_go_fuzz_dep_.CoverTab[86136]++
															return aes.BlockSize * 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:68
	// _ = "end of CoverTab[86136]"
}

// StringToKey returns a key derived from the string provided.
func (e Aes128CtsHmacSha256128) StringToKey(secret string, salt string, s2kparams string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:72
	_go_fuzz_dep_.CoverTab[86137]++
															saltp := rfc8009.GetSaltP(salt, "aes128-cts-hmac-sha256-128")
															return rfc8009.StringToKey(secret, saltp, s2kparams, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:74
	// _ = "end of CoverTab[86137]"
}

// RandomToKey returns a key from the bytes provided.
func (e Aes128CtsHmacSha256128) RandomToKey(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:78
	_go_fuzz_dep_.CoverTab[86138]++
															return rfc8009.RandomToKey(b)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:79
	// _ = "end of CoverTab[86138]"
}

// EncryptData encrypts the data provided.
func (e Aes128CtsHmacSha256128) EncryptData(key, data []byte) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:83
	_go_fuzz_dep_.CoverTab[86139]++
															return rfc8009.EncryptData(key, data, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:84
	// _ = "end of CoverTab[86139]"
}

// EncryptMessage encrypts the message provided and concatenates it with the integrity hash to create an encrypted message.
func (e Aes128CtsHmacSha256128) EncryptMessage(key, message []byte, usage uint32) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:88
	_go_fuzz_dep_.CoverTab[86140]++
															return rfc8009.EncryptMessage(key, message, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:89
	// _ = "end of CoverTab[86140]"
}

// DecryptData decrypts the data provided.
func (e Aes128CtsHmacSha256128) DecryptData(key, data []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:93
	_go_fuzz_dep_.CoverTab[86141]++
															return rfc8009.DecryptData(key, data, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:94
	// _ = "end of CoverTab[86141]"
}

// DecryptMessage decrypts the message provided and verifies the integrity of the message.
func (e Aes128CtsHmacSha256128) DecryptMessage(key, ciphertext []byte, usage uint32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:98
	_go_fuzz_dep_.CoverTab[86142]++
															return rfc8009.DecryptMessage(key, ciphertext, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:99
	// _ = "end of CoverTab[86142]"
}

// DeriveKey derives a key from the protocol key based on the usage value.
func (e Aes128CtsHmacSha256128) DeriveKey(protocolKey, usage []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:103
	_go_fuzz_dep_.CoverTab[86143]++
															return rfc8009.DeriveKey(protocolKey, usage, e), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:104
	// _ = "end of CoverTab[86143]"
}

// DeriveRandom generates data needed for key generation.
func (e Aes128CtsHmacSha256128) DeriveRandom(protocolKey, usage []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:108
	_go_fuzz_dep_.CoverTab[86144]++
															return rfc8009.DeriveRandom(protocolKey, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:109
	// _ = "end of CoverTab[86144]"
}

// VerifyIntegrity checks the integrity of the ciphertext message.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:112
// As the hash is calculated over the iv concatenated with the AES cipher output not the plaintext the pt value to this
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:112
// interface method is not use. Pass any []byte.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:115
func (e Aes128CtsHmacSha256128) VerifyIntegrity(protocolKey, ct, pt []byte, usage uint32) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:115
	_go_fuzz_dep_.CoverTab[86145]++

															return rfc8009.VerifyIntegrity(protocolKey, ct, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:117
	// _ = "end of CoverTab[86145]"
}

// GetChecksumHash returns a keyed checksum hash of the bytes provided.
func (e Aes128CtsHmacSha256128) GetChecksumHash(protocolKey, data []byte, usage uint32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:121
	_go_fuzz_dep_.CoverTab[86146]++
															return common.GetHash(data, protocolKey, common.GetUsageKc(usage), e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:122
	// _ = "end of CoverTab[86146]"
}

// VerifyChecksum compares the checksum of the message bytes is the same as the checksum provided.
func (e Aes128CtsHmacSha256128) VerifyChecksum(protocolKey, data, chksum []byte, usage uint32) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:126
	_go_fuzz_dep_.CoverTab[86147]++
															c, err := e.GetChecksumHash(protocolKey, data, usage)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:128
		_go_fuzz_dep_.CoverTab[86149]++
																return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:129
		// _ = "end of CoverTab[86149]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:130
		_go_fuzz_dep_.CoverTab[86150]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:130
		// _ = "end of CoverTab[86150]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:130
	// _ = "end of CoverTab[86147]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:130
	_go_fuzz_dep_.CoverTab[86148]++
															return hmac.Equal(chksum, c)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:131
	// _ = "end of CoverTab[86148]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:132
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha256-128.go:132
var _ = _go_fuzz_dep_.CoverTab
