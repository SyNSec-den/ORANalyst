//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:1
package crypto

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:1
)

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"hash"
	"io"

	"github.com/jcmturner/gokrb5/v8/crypto/rfc3961"
	"github.com/jcmturner/gokrb5/v8/crypto/rfc4757"
	"github.com/jcmturner/gokrb5/v8/iana/chksumtype"
	"github.com/jcmturner/gokrb5/v8/iana/etypeID"
	"golang.org/x/crypto/md4"
)

// RC4HMAC implements Kerberos encryption type rc4-hmac
type RC4HMAC struct {
}

// GetETypeID returns the EType ID number.
func (e RC4HMAC) GetETypeID() int32 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:22
	_go_fuzz_dep_.CoverTab[86304]++
												return etypeID.RC4_HMAC
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:23
	// _ = "end of CoverTab[86304]"
}

// GetHashID returns the checksum type ID number.
func (e RC4HMAC) GetHashID() int32 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:27
	_go_fuzz_dep_.CoverTab[86305]++
												return chksumtype.KERB_CHECKSUM_HMAC_MD5
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:28
	// _ = "end of CoverTab[86305]"
}

// GetKeyByteSize returns the number of bytes for key of this etype.
func (e RC4HMAC) GetKeyByteSize() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:32
	_go_fuzz_dep_.CoverTab[86306]++
												return 16
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:33
	// _ = "end of CoverTab[86306]"
}

// GetKeySeedBitLength returns the number of bits for the seed for key generation.
func (e RC4HMAC) GetKeySeedBitLength() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:37
	_go_fuzz_dep_.CoverTab[86307]++
												return e.GetKeyByteSize() * 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:38
	// _ = "end of CoverTab[86307]"
}

// GetHashFunc returns the hash function for this etype.
func (e RC4HMAC) GetHashFunc() func() hash.Hash {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:42
	_go_fuzz_dep_.CoverTab[86308]++
												return md5.New
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:43
	// _ = "end of CoverTab[86308]"
}

// GetMessageBlockByteSize returns the block size for the etype's messages.
func (e RC4HMAC) GetMessageBlockByteSize() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:47
	_go_fuzz_dep_.CoverTab[86309]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:48
	// _ = "end of CoverTab[86309]"
}

// GetDefaultStringToKeyParams returns the default key derivation parameters in string form.
func (e RC4HMAC) GetDefaultStringToKeyParams() string {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:52
	_go_fuzz_dep_.CoverTab[86310]++
												return ""
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:53
	// _ = "end of CoverTab[86310]"
}

// GetConfounderByteSize returns the byte count for confounder to be used during cryptographic operations.
func (e RC4HMAC) GetConfounderByteSize() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:57
	_go_fuzz_dep_.CoverTab[86311]++
												return 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:58
	// _ = "end of CoverTab[86311]"
}

// GetHMACBitLength returns the bit count size of the integrity hash.
func (e RC4HMAC) GetHMACBitLength() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:62
	_go_fuzz_dep_.CoverTab[86312]++
												return md5.Size * 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:63
	// _ = "end of CoverTab[86312]"
}

// GetCypherBlockBitLength returns the bit count size of the cypher block.
func (e RC4HMAC) GetCypherBlockBitLength() int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:67
	_go_fuzz_dep_.CoverTab[86313]++
												return 8
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:68
	// _ = "end of CoverTab[86313]"
}

// StringToKey returns a key derived from the string provided.
func (e RC4HMAC) StringToKey(secret string, salt string, s2kparams string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:72
	_go_fuzz_dep_.CoverTab[86314]++
												return rfc4757.StringToKey(secret)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:73
	// _ = "end of CoverTab[86314]"
}

// RandomToKey returns a key from the bytes provided.
func (e RC4HMAC) RandomToKey(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:77
	_go_fuzz_dep_.CoverTab[86315]++
												r := bytes.NewReader(b)
												h := md4.New()
												io.Copy(h, r)
												return h.Sum(nil)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:81
	// _ = "end of CoverTab[86315]"
}

// EncryptData encrypts the data provided.
func (e RC4HMAC) EncryptData(key, data []byte) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:85
	_go_fuzz_dep_.CoverTab[86316]++
												b, err := rfc4757.EncryptData(key, data, e)
												return []byte{}, b, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:87
	// _ = "end of CoverTab[86316]"
}

// EncryptMessage encrypts the message provided and concatenates it with the integrity hash to create an encrypted message.
func (e RC4HMAC) EncryptMessage(key, message []byte, usage uint32) ([]byte, []byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:91
	_go_fuzz_dep_.CoverTab[86317]++
												b, err := rfc4757.EncryptMessage(key, message, usage, false, e)
												return []byte{}, b, err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:93
	// _ = "end of CoverTab[86317]"
}

// DecryptData decrypts the data provided.
func (e RC4HMAC) DecryptData(key, data []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:97
	_go_fuzz_dep_.CoverTab[86318]++
												return rfc4757.DecryptData(key, data, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:98
	// _ = "end of CoverTab[86318]"
}

// DecryptMessage decrypts the message provided and verifies the integrity of the message.
func (e RC4HMAC) DecryptMessage(key, ciphertext []byte, usage uint32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:102
	_go_fuzz_dep_.CoverTab[86319]++
												return rfc4757.DecryptMessage(key, ciphertext, usage, false, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:103
	// _ = "end of CoverTab[86319]"
}

// DeriveKey derives a key from the protocol key based on the usage value.
func (e RC4HMAC) DeriveKey(protocolKey, usage []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:107
	_go_fuzz_dep_.CoverTab[86320]++
												return rfc4757.HMAC(protocolKey, usage), nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:108
	// _ = "end of CoverTab[86320]"
}

// DeriveRandom generates data needed for key generation.
func (e RC4HMAC) DeriveRandom(protocolKey, usage []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:112
	_go_fuzz_dep_.CoverTab[86321]++
												return rfc3961.DeriveRandom(protocolKey, usage, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:113
	// _ = "end of CoverTab[86321]"
}

// VerifyIntegrity checks the integrity of the plaintext message.
func (e RC4HMAC) VerifyIntegrity(protocolKey, ct, pt []byte, usage uint32) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:117
	_go_fuzz_dep_.CoverTab[86322]++
												return rfc4757.VerifyIntegrity(protocolKey, pt, ct, e)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:118
	// _ = "end of CoverTab[86322]"
}

// GetChecksumHash returns a keyed checksum hash of the bytes provided.
func (e RC4HMAC) GetChecksumHash(protocolKey, data []byte, usage uint32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:122
	_go_fuzz_dep_.CoverTab[86323]++
												return rfc4757.Checksum(protocolKey, usage, data)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:123
	// _ = "end of CoverTab[86323]"
}

// VerifyChecksum compares the checksum of the message bytes is the same as the checksum provided.
func (e RC4HMAC) VerifyChecksum(protocolKey, data, chksum []byte, usage uint32) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:127
	_go_fuzz_dep_.CoverTab[86324]++
												checksum, err := rfc4757.Checksum(protocolKey, usage, data)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:129
		_go_fuzz_dep_.CoverTab[86326]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:130
		// _ = "end of CoverTab[86326]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:131
		_go_fuzz_dep_.CoverTab[86327]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:131
		// _ = "end of CoverTab[86327]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:131
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:131
	// _ = "end of CoverTab[86324]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:131
	_go_fuzz_dep_.CoverTab[86325]++
												return hmac.Equal(checksum, chksum)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:132
	// _ = "end of CoverTab[86325]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:133
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rc4-hmac.go:133
var _ = _go_fuzz_dep_.CoverTab
