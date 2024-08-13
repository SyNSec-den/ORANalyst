//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:1
package crypto

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:1
)

import (
	"crypto/aes"
	"crypto/hmac"
	"crypto/sha512"
	"hash"

	"github.com/jcmturner/gokrb5/v8/crypto/common"
	"github.com/jcmturner/gokrb5/v8/crypto/rfc8009"
	"github.com/jcmturner/gokrb5/v8/iana/chksumtype"
	"github.com/jcmturner/gokrb5/v8/iana/etypeID"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:17
// Aes256CtsHmacSha384192 implements Kerberos encryption type aes256-cts-hmac-sha384-192
type Aes256CtsHmacSha384192 struct {
}

// GetETypeID returns the EType ID number.
func (e Aes256CtsHmacSha384192) GetETypeID() int32 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:22
	_go_fuzz_dep_.CoverTab[86175]++
															return etypeID.AES256_CTS_HMAC_SHA384_192
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:23
	// _ = "end of CoverTab[86175]"
}

// GetHashID returns the checksum type ID number.
func (e Aes256CtsHmacSha384192) GetHashID() int32 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:27
	_go_fuzz_dep_.CoverTab[86176]++
															return chksumtype.HMAC_SHA384_192_AES256
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:28
	// _ = "end of CoverTab[86176]"
}

// GetKeyByteSize returns the number of bytes for key of this etype.
func (e Aes256CtsHmacSha384192) GetKeyByteSize() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:32
	_go_fuzz_dep_.CoverTab[86177]++
															return 192 / 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:33
	// _ = "end of CoverTab[86177]"
}

// GetKeySeedBitLength returns the number of bits for the seed for key generation.
func (e Aes256CtsHmacSha384192) GetKeySeedBitLength() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:37
	_go_fuzz_dep_.CoverTab[86178]++
															return e.GetKeyByteSize() * 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:38
	// _ = "end of CoverTab[86178]"
}

// GetHashFunc returns the hash function for this etype.
func (e Aes256CtsHmacSha384192) GetHashFunc() func() hash.Hash {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:42
	_go_fuzz_dep_.CoverTab[86179]++
															return sha512.New384
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:43
	// _ = "end of CoverTab[86179]"
}

// GetMessageBlockByteSize returns the block size for the etype's messages.
func (e Aes256CtsHmacSha384192) GetMessageBlockByteSize() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:47
	_go_fuzz_dep_.CoverTab[86180]++
															return 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:48
	// _ = "end of CoverTab[86180]"
}

// GetDefaultStringToKeyParams returns the default key derivation parameters in string form.
func (e Aes256CtsHmacSha384192) GetDefaultStringToKeyParams() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:52
	_go_fuzz_dep_.CoverTab[86181]++
															return "00008000"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:53
	// _ = "end of CoverTab[86181]"
}

// GetConfounderByteSize returns the byte count for confounder to be used during cryptographic operations.
func (e Aes256CtsHmacSha384192) GetConfounderByteSize() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:57
	_go_fuzz_dep_.CoverTab[86182]++
															return aes.BlockSize
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:58
	// _ = "end of CoverTab[86182]"
}

// GetHMACBitLength returns the bit count size of the integrity hash.
func (e Aes256CtsHmacSha384192) GetHMACBitLength() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:62
	_go_fuzz_dep_.CoverTab[86183]++
															return 192
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:63
	// _ = "end of CoverTab[86183]"
}

// GetCypherBlockBitLength returns the bit count size of the cypher block.
func (e Aes256CtsHmacSha384192) GetCypherBlockBitLength() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:67
	_go_fuzz_dep_.CoverTab[86184]++
															return aes.BlockSize * 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:68
	// _ = "end of CoverTab[86184]"
}

// StringToKey returns a key derived from the string provided.
func (e Aes256CtsHmacSha384192) StringToKey(secret string, salt string, s2kparams string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:72
	_go_fuzz_dep_.CoverTab[86185]++
															saltp := rfc8009.GetSaltP(salt, "aes256-cts-hmac-sha384-192")
															return rfc8009.StringToKey(secret, saltp, s2kparams, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:74
	// _ = "end of CoverTab[86185]"
}

// RandomToKey returns a key from the bytes provided.
func (e Aes256CtsHmacSha384192) RandomToKey(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:78
	_go_fuzz_dep_.CoverTab[86186]++
															return rfc8009.RandomToKey(b)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:79
	// _ = "end of CoverTab[86186]"
}

// EncryptData encrypts the data provided.
func (e Aes256CtsHmacSha384192) EncryptData(key, data []byte) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:83
	_go_fuzz_dep_.CoverTab[86187]++
															return rfc8009.EncryptData(key, data, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:84
	// _ = "end of CoverTab[86187]"
}

// EncryptMessage encrypts the message provided and concatenates it with the integrity hash to create an encrypted message.
func (e Aes256CtsHmacSha384192) EncryptMessage(key, message []byte, usage uint32) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:88
	_go_fuzz_dep_.CoverTab[86188]++
															return rfc8009.EncryptMessage(key, message, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:89
	// _ = "end of CoverTab[86188]"
}

// DecryptData decrypts the data provided.
func (e Aes256CtsHmacSha384192) DecryptData(key, data []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:93
	_go_fuzz_dep_.CoverTab[86189]++
															return rfc8009.DecryptData(key, data, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:94
	// _ = "end of CoverTab[86189]"
}

// DecryptMessage decrypts the message provided and verifies the integrity of the message.
func (e Aes256CtsHmacSha384192) DecryptMessage(key, ciphertext []byte, usage uint32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:98
	_go_fuzz_dep_.CoverTab[86190]++
															return rfc8009.DecryptMessage(key, ciphertext, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:99
	// _ = "end of CoverTab[86190]"
}

// DeriveKey derives a key from the protocol key based on the usage value.
func (e Aes256CtsHmacSha384192) DeriveKey(protocolKey, usage []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:103
	_go_fuzz_dep_.CoverTab[86191]++
															return rfc8009.DeriveKey(protocolKey, usage, e), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:104
	// _ = "end of CoverTab[86191]"
}

// DeriveRandom generates data needed for key generation.
func (e Aes256CtsHmacSha384192) DeriveRandom(protocolKey, usage []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:108
	_go_fuzz_dep_.CoverTab[86192]++
															return rfc8009.DeriveRandom(protocolKey, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:109
	// _ = "end of CoverTab[86192]"
}

// VerifyIntegrity checks the integrity of the ciphertext message.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:112
// As the hash is calculated over the iv concatenated with the AES cipher output not the plaintext the pt value to this
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:112
// interface method is not use. Pass any []byte.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:115
func (e Aes256CtsHmacSha384192) VerifyIntegrity(protocolKey, ct, pt []byte, usage uint32) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:115
	_go_fuzz_dep_.CoverTab[86193]++

															return rfc8009.VerifyIntegrity(protocolKey, ct, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:117
	// _ = "end of CoverTab[86193]"
}

// GetChecksumHash returns a keyed checksum hash of the bytes provided.
func (e Aes256CtsHmacSha384192) GetChecksumHash(protocolKey, data []byte, usage uint32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:121
	_go_fuzz_dep_.CoverTab[86194]++
															return common.GetHash(data, protocolKey, common.GetUsageKc(usage), e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:122
	// _ = "end of CoverTab[86194]"
}

// VerifyChecksum compares the checksum of the message bytes is the same as the checksum provided.
func (e Aes256CtsHmacSha384192) VerifyChecksum(protocolKey, data, chksum []byte, usage uint32) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:126
	_go_fuzz_dep_.CoverTab[86195]++
															c, err := e.GetChecksumHash(protocolKey, data, usage)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:128
		_go_fuzz_dep_.CoverTab[86197]++
																return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:129
		// _ = "end of CoverTab[86197]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:130
		_go_fuzz_dep_.CoverTab[86198]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:130
		// _ = "end of CoverTab[86198]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:130
	// _ = "end of CoverTab[86195]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:130
	_go_fuzz_dep_.CoverTab[86196]++
															return hmac.Equal(chksum, c)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:131
	// _ = "end of CoverTab[86196]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:132
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes256-cts-hmac-sha384-192.go:132
var _ = _go_fuzz_dep_.CoverTab
