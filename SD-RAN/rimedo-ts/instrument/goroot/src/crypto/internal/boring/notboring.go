// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !(boringcrypto && linux && (amd64 || arm64) && !android && !cmd_go_bootstrap && !msan && cgo)

//line /usr/local/go/src/crypto/internal/boring/notboring.go:7
package boring

//line /usr/local/go/src/crypto/internal/boring/notboring.go:7
import (
//line /usr/local/go/src/crypto/internal/boring/notboring.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:7
)
//line /usr/local/go/src/crypto/internal/boring/notboring.go:7
import (
//line /usr/local/go/src/crypto/internal/boring/notboring.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:7
)

import (
	"crypto"
	"crypto/cipher"
	"crypto/internal/boring/sig"
	"hash"
)

const available = false

// Unreachable marks code that should be unreachable
//line /usr/local/go/src/crypto/internal/boring/notboring.go:18
// when BoringCrypto is in use. It is a no-op without BoringCrypto.
//line /usr/local/go/src/crypto/internal/boring/notboring.go:20
func Unreachable() {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:20
	_go_fuzz_dep_.CoverTab[1706]++

//line /usr/local/go/src/crypto/internal/boring/notboring.go:24
	sig.StandardCrypto()
//line /usr/local/go/src/crypto/internal/boring/notboring.go:24
	// _ = "end of CoverTab[1706]"
}

// UnreachableExceptTests marks code that should be unreachable
//line /usr/local/go/src/crypto/internal/boring/notboring.go:27
// when BoringCrypto is in use. It is a no-op without BoringCrypto.
//line /usr/local/go/src/crypto/internal/boring/notboring.go:29
func UnreachableExceptTests()	{ _go_fuzz_dep_.CoverTab[1707]++; // _ = "end of CoverTab[1707]" }

type randReader int

func (randReader) Read(b []byte) (int, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:33
	_go_fuzz_dep_.CoverTab[1708]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:33
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:33
	// _ = "end of CoverTab[1708]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:33
}

const RandReader = randReader(0)

func NewSHA1() hash.Hash {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:37
	_go_fuzz_dep_.CoverTab[1709]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:37
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:37
	// _ = "end of CoverTab[1709]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:37
}
func NewSHA224() hash.Hash {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:38
	_go_fuzz_dep_.CoverTab[1710]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:38
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:38
	// _ = "end of CoverTab[1710]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:38
}
func NewSHA256() hash.Hash {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:39
	_go_fuzz_dep_.CoverTab[1711]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:39
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:39
	// _ = "end of CoverTab[1711]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:39
}
func NewSHA384() hash.Hash {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:40
	_go_fuzz_dep_.CoverTab[1712]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:40
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:40
	// _ = "end of CoverTab[1712]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:40
}
func NewSHA512() hash.Hash {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:41
	_go_fuzz_dep_.CoverTab[1713]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:41
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:41
	// _ = "end of CoverTab[1713]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:41
}

func SHA1([]byte) [20]byte {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:43
	_go_fuzz_dep_.CoverTab[1714]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:43
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:43
	// _ = "end of CoverTab[1714]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:43
}
func SHA224([]byte) [28]byte {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:44
	_go_fuzz_dep_.CoverTab[1715]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:44
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:44
	// _ = "end of CoverTab[1715]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:44
}
func SHA256([]byte) [32]byte {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:45
	_go_fuzz_dep_.CoverTab[1716]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:45
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:45
	// _ = "end of CoverTab[1716]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:45
}
func SHA384([]byte) [48]byte {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:46
	_go_fuzz_dep_.CoverTab[1717]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:46
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:46
	// _ = "end of CoverTab[1717]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:46
}
func SHA512([]byte) [64]byte {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:47
	_go_fuzz_dep_.CoverTab[1718]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:47
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:47
	// _ = "end of CoverTab[1718]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:47
}

func NewHMAC(h func() hash.Hash, key []byte) hash.Hash {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:49
	_go_fuzz_dep_.CoverTab[1719]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:49
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:49
	// _ = "end of CoverTab[1719]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:49
}

func NewAESCipher(key []byte) (cipher.Block, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:51
	_go_fuzz_dep_.CoverTab[1720]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:51
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:51
	// _ = "end of CoverTab[1720]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:51
}
func NewGCMTLS(cipher.Block) (cipher.AEAD, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:52
	_go_fuzz_dep_.CoverTab[1721]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:52
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:52
	// _ = "end of CoverTab[1721]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:52
}

type PublicKeyECDSA struct{ _ int }
type PrivateKeyECDSA struct{ _ int }

func GenerateKeyECDSA(curve string) (X, Y, D BigInt, err error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:57
	_go_fuzz_dep_.CoverTab[1722]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:58
	// _ = "end of CoverTab[1722]"
}
func NewPrivateKeyECDSA(curve string, X, Y, D BigInt) (*PrivateKeyECDSA, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:60
	_go_fuzz_dep_.CoverTab[1723]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:61
	// _ = "end of CoverTab[1723]"
}
func NewPublicKeyECDSA(curve string, X, Y BigInt) (*PublicKeyECDSA, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:63
	_go_fuzz_dep_.CoverTab[1724]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:64
	// _ = "end of CoverTab[1724]"
}
func SignMarshalECDSA(priv *PrivateKeyECDSA, hash []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:66
	_go_fuzz_dep_.CoverTab[1725]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:67
	// _ = "end of CoverTab[1725]"
}
func VerifyECDSA(pub *PublicKeyECDSA, hash []byte, sig []byte) bool {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:69
	_go_fuzz_dep_.CoverTab[1726]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:70
	// _ = "end of CoverTab[1726]"
}

type PublicKeyRSA struct{ _ int }
type PrivateKeyRSA struct{ _ int }

func DecryptRSAOAEP(h, mgfHash hash.Hash, priv *PrivateKeyRSA, ciphertext, label []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:76
	_go_fuzz_dep_.CoverTab[1727]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:77
	// _ = "end of CoverTab[1727]"
}
func DecryptRSAPKCS1(priv *PrivateKeyRSA, ciphertext []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:79
	_go_fuzz_dep_.CoverTab[1728]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:80
	// _ = "end of CoverTab[1728]"
}
func DecryptRSANoPadding(priv *PrivateKeyRSA, ciphertext []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:82
	_go_fuzz_dep_.CoverTab[1729]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:83
	// _ = "end of CoverTab[1729]"
}
func EncryptRSAOAEP(h, mgfHash hash.Hash, pub *PublicKeyRSA, msg, label []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:85
	_go_fuzz_dep_.CoverTab[1730]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:86
	// _ = "end of CoverTab[1730]"
}
func EncryptRSAPKCS1(pub *PublicKeyRSA, msg []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:88
	_go_fuzz_dep_.CoverTab[1731]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:89
	// _ = "end of CoverTab[1731]"
}
func EncryptRSANoPadding(pub *PublicKeyRSA, msg []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:91
	_go_fuzz_dep_.CoverTab[1732]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:92
	// _ = "end of CoverTab[1732]"
}
func GenerateKeyRSA(bits int) (N, E, D, P, Q, Dp, Dq, Qinv BigInt, err error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:94
	_go_fuzz_dep_.CoverTab[1733]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:95
	// _ = "end of CoverTab[1733]"
}
func NewPrivateKeyRSA(N, E, D, P, Q, Dp, Dq, Qinv BigInt) (*PrivateKeyRSA, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:97
	_go_fuzz_dep_.CoverTab[1734]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:98
	// _ = "end of CoverTab[1734]"
}
func NewPublicKeyRSA(N, E BigInt) (*PublicKeyRSA, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:100
	_go_fuzz_dep_.CoverTab[1735]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:100
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:100
	// _ = "end of CoverTab[1735]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:100
}
func SignRSAPKCS1v15(priv *PrivateKeyRSA, h crypto.Hash, hashed []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:101
	_go_fuzz_dep_.CoverTab[1736]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:102
	// _ = "end of CoverTab[1736]"
}
func SignRSAPSS(priv *PrivateKeyRSA, h crypto.Hash, hashed []byte, saltLen int) ([]byte, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:104
	_go_fuzz_dep_.CoverTab[1737]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:105
	// _ = "end of CoverTab[1737]"
}
func VerifyRSAPKCS1v15(pub *PublicKeyRSA, h crypto.Hash, hashed, sig []byte) error {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:107
	_go_fuzz_dep_.CoverTab[1738]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:108
	// _ = "end of CoverTab[1738]"
}
func VerifyRSAPSS(pub *PublicKeyRSA, h crypto.Hash, hashed, sig []byte, saltLen int) error {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:110
	_go_fuzz_dep_.CoverTab[1739]++
									panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:111
	// _ = "end of CoverTab[1739]"
}

type PublicKeyECDH struct{}
type PrivateKeyECDH struct{}

func ECDH(*PrivateKeyECDH, *PublicKeyECDH) ([]byte, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:117
	_go_fuzz_dep_.CoverTab[1740]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:117
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:117
	// _ = "end of CoverTab[1740]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:117
}
func GenerateKeyECDH(string) (*PrivateKeyECDH, []byte, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:118
	_go_fuzz_dep_.CoverTab[1741]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:118
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:118
	// _ = "end of CoverTab[1741]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:118
}
func NewPrivateKeyECDH(string, []byte) (*PrivateKeyECDH, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:119
	_go_fuzz_dep_.CoverTab[1742]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:119
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:119
	// _ = "end of CoverTab[1742]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:119
}
func NewPublicKeyECDH(string, []byte) (*PublicKeyECDH, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:120
	_go_fuzz_dep_.CoverTab[1743]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:120
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:120
	// _ = "end of CoverTab[1743]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:120
}
func (*PublicKeyECDH) Bytes() []byte {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:121
	_go_fuzz_dep_.CoverTab[1744]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:121
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:121
	// _ = "end of CoverTab[1744]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:121
}
func (*PrivateKeyECDH) PublicKey() (*PublicKeyECDH, error) {
//line /usr/local/go/src/crypto/internal/boring/notboring.go:122
	_go_fuzz_dep_.CoverTab[1745]++
//line /usr/local/go/src/crypto/internal/boring/notboring.go:122
	panic("boringcrypto: not available")
//line /usr/local/go/src/crypto/internal/boring/notboring.go:122
	// _ = "end of CoverTab[1745]"
//line /usr/local/go/src/crypto/internal/boring/notboring.go:122
}

//line /usr/local/go/src/crypto/internal/boring/notboring.go:122
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/internal/boring/notboring.go:122
var _ = _go_fuzz_dep_.CoverTab
