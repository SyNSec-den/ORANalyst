//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:1
package crypto

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:1
)

import (
	"crypto/aes"
	"crypto/hmac"
	"crypto/sha1"
	"hash"

	"github.com/jcmturner/gokrb5/v8/crypto/common"
	"github.com/jcmturner/gokrb5/v8/crypto/rfc3961"
	"github.com/jcmturner/gokrb5/v8/crypto/rfc3962"
	"github.com/jcmturner/gokrb5/v8/iana/chksumtype"
	"github.com/jcmturner/gokrb5/v8/iana/etypeID"
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:18
// Aes128CtsHmacSha96 implements Kerberos encryption type aes128-cts-hmac-sha1-96
type Aes128CtsHmacSha96 struct {
}

// GetETypeID returns the EType ID number.
func (e Aes128CtsHmacSha96) GetETypeID() int32 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:23
	_go_fuzz_dep_.CoverTab[86103]++
														return etypeID.AES128_CTS_HMAC_SHA1_96
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:24
	// _ = "end of CoverTab[86103]"
}

// GetHashID returns the checksum type ID number.
func (e Aes128CtsHmacSha96) GetHashID() int32 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:28
	_go_fuzz_dep_.CoverTab[86104]++
														return chksumtype.HMAC_SHA1_96_AES128
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:29
	// _ = "end of CoverTab[86104]"
}

// GetKeyByteSize returns the number of bytes for key of this etype.
func (e Aes128CtsHmacSha96) GetKeyByteSize() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:33
	_go_fuzz_dep_.CoverTab[86105]++
														return 128 / 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:34
	// _ = "end of CoverTab[86105]"
}

// GetKeySeedBitLength returns the number of bits for the seed for key generation.
func (e Aes128CtsHmacSha96) GetKeySeedBitLength() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:38
	_go_fuzz_dep_.CoverTab[86106]++
														return e.GetKeyByteSize() * 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:39
	// _ = "end of CoverTab[86106]"
}

// GetHashFunc returns the hash function for this etype.
func (e Aes128CtsHmacSha96) GetHashFunc() func() hash.Hash {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:43
	_go_fuzz_dep_.CoverTab[86107]++
														return sha1.New
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:44
	// _ = "end of CoverTab[86107]"
}

// GetMessageBlockByteSize returns the block size for the etype's messages.
func (e Aes128CtsHmacSha96) GetMessageBlockByteSize() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:48
	_go_fuzz_dep_.CoverTab[86108]++
														return 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:49
	// _ = "end of CoverTab[86108]"
}

// GetDefaultStringToKeyParams returns the default key derivation parameters in string form.
func (e Aes128CtsHmacSha96) GetDefaultStringToKeyParams() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:53
	_go_fuzz_dep_.CoverTab[86109]++
														return "00001000"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:54
	// _ = "end of CoverTab[86109]"
}

// GetConfounderByteSize returns the byte count for confounder to be used during cryptographic operations.
func (e Aes128CtsHmacSha96) GetConfounderByteSize() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:58
	_go_fuzz_dep_.CoverTab[86110]++
														return aes.BlockSize
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:59
	// _ = "end of CoverTab[86110]"
}

// GetHMACBitLength returns the bit count size of the integrity hash.
func (e Aes128CtsHmacSha96) GetHMACBitLength() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:63
	_go_fuzz_dep_.CoverTab[86111]++
														return 96
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:64
	// _ = "end of CoverTab[86111]"
}

// GetCypherBlockBitLength returns the bit count size of the cypher block.
func (e Aes128CtsHmacSha96) GetCypherBlockBitLength() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:68
	_go_fuzz_dep_.CoverTab[86112]++
														return aes.BlockSize * 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:69
	// _ = "end of CoverTab[86112]"
}

// StringToKey returns a key derived from the string provided.
func (e Aes128CtsHmacSha96) StringToKey(secret string, salt string, s2kparams string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:73
	_go_fuzz_dep_.CoverTab[86113]++
														return rfc3962.StringToKey(secret, salt, s2kparams, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:74
	// _ = "end of CoverTab[86113]"
}

// RandomToKey returns a key from the bytes provided.
func (e Aes128CtsHmacSha96) RandomToKey(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:78
	_go_fuzz_dep_.CoverTab[86114]++
														return rfc3961.RandomToKey(b)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:79
	// _ = "end of CoverTab[86114]"
}

// EncryptData encrypts the data provided.
func (e Aes128CtsHmacSha96) EncryptData(key, data []byte) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:83
	_go_fuzz_dep_.CoverTab[86115]++
														return rfc3962.EncryptData(key, data, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:84
	// _ = "end of CoverTab[86115]"
}

// EncryptMessage encrypts the message provided and concatenates it with the integrity hash to create an encrypted message.
func (e Aes128CtsHmacSha96) EncryptMessage(key, message []byte, usage uint32) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:88
	_go_fuzz_dep_.CoverTab[86116]++
														return rfc3962.EncryptMessage(key, message, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:89
	// _ = "end of CoverTab[86116]"
}

// DecryptData decrypts the data provided.
func (e Aes128CtsHmacSha96) DecryptData(key, data []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:93
	_go_fuzz_dep_.CoverTab[86117]++
														return rfc3962.DecryptData(key, data, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:94
	// _ = "end of CoverTab[86117]"
}

// DecryptMessage decrypts the message provided and verifies the integrity of the message.
func (e Aes128CtsHmacSha96) DecryptMessage(key, ciphertext []byte, usage uint32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:98
	_go_fuzz_dep_.CoverTab[86118]++
														return rfc3962.DecryptMessage(key, ciphertext, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:99
	// _ = "end of CoverTab[86118]"
}

// DeriveKey derives a key from the protocol key based on the usage value.
func (e Aes128CtsHmacSha96) DeriveKey(protocolKey, usage []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:103
	_go_fuzz_dep_.CoverTab[86119]++
														return rfc3961.DeriveKey(protocolKey, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:104
	// _ = "end of CoverTab[86119]"
}

// DeriveRandom generates data needed for key generation.
func (e Aes128CtsHmacSha96) DeriveRandom(protocolKey, usage []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:108
	_go_fuzz_dep_.CoverTab[86120]++
														return rfc3961.DeriveRandom(protocolKey, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:109
	// _ = "end of CoverTab[86120]"
}

// VerifyIntegrity checks the integrity of the plaintext message.
func (e Aes128CtsHmacSha96) VerifyIntegrity(protocolKey, ct, pt []byte, usage uint32) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:113
	_go_fuzz_dep_.CoverTab[86121]++
														return rfc3961.VerifyIntegrity(protocolKey, ct, pt, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:114
	// _ = "end of CoverTab[86121]"
}

// GetChecksumHash returns a keyed checksum hash of the bytes provided.
func (e Aes128CtsHmacSha96) GetChecksumHash(protocolKey, data []byte, usage uint32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:118
	_go_fuzz_dep_.CoverTab[86122]++
														return common.GetHash(data, protocolKey, common.GetUsageKc(usage), e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:119
	// _ = "end of CoverTab[86122]"
}

// VerifyChecksum compares the checksum of the message bytes is the same as the checksum provided.
func (e Aes128CtsHmacSha96) VerifyChecksum(protocolKey, data, chksum []byte, usage uint32) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:123
	_go_fuzz_dep_.CoverTab[86123]++
														c, err := e.GetChecksumHash(protocolKey, data, usage)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:125
		_go_fuzz_dep_.CoverTab[86125]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:126
		// _ = "end of CoverTab[86125]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:127
		_go_fuzz_dep_.CoverTab[86126]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:127
		// _ = "end of CoverTab[86126]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:127
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:127
	// _ = "end of CoverTab[86123]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:127
	_go_fuzz_dep_.CoverTab[86124]++
														return hmac.Equal(chksum, c)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:128
	// _ = "end of CoverTab[86124]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:129
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/aes128-cts-hmac-sha1-96.go:129
var _ = _go_fuzz_dep_.CoverTab
