//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:17
package josecipher

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:17
)

import (
	"crypto"
	"crypto/ecdsa"
	"encoding/binary"
)

// DeriveECDHES derives a shared encryption key using ECDH/ConcatKDF as described in JWE/JWA.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:25
// It is an error to call this function with a private/public key that are not on the same
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:25
// curve. Callers must ensure that the keys are valid before calling this function. Output
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:25
// size may be at most 1<<16 bytes (64 KiB).
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:29
func DeriveECDHES(alg string, apuData, apvData []byte, priv *ecdsa.PrivateKey, pub *ecdsa.PublicKey, size int) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:29
	_go_fuzz_dep_.CoverTab[184340]++
												if size > 1<<16 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:30
		_go_fuzz_dep_.CoverTab[184343]++
													panic("ECDH-ES output size too large, must be less than or equal to 1<<16")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:31
		// _ = "end of CoverTab[184343]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:32
		_go_fuzz_dep_.CoverTab[184344]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:32
		// _ = "end of CoverTab[184344]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:32
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:32
	// _ = "end of CoverTab[184340]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:32
	_go_fuzz_dep_.CoverTab[184341]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:35
	algID := lengthPrefixed([]byte(alg))
												ptyUInfo := lengthPrefixed(apuData)
												ptyVInfo := lengthPrefixed(apvData)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:40
	supPubInfo := make([]byte, 4)
	binary.BigEndian.PutUint32(supPubInfo, uint32(size)*8)

	if !priv.PublicKey.Curve.IsOnCurve(pub.X, pub.Y) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:43
		_go_fuzz_dep_.CoverTab[184345]++
													panic("public key not on same curve as private key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:44
		// _ = "end of CoverTab[184345]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:45
		_go_fuzz_dep_.CoverTab[184346]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:45
		// _ = "end of CoverTab[184346]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:45
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:45
	// _ = "end of CoverTab[184341]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:45
	_go_fuzz_dep_.CoverTab[184342]++

												z, _ := priv.PublicKey.Curve.ScalarMult(pub.X, pub.Y, priv.D.Bytes())
												reader := NewConcatKDF(crypto.SHA256, z.Bytes(), algID, ptyUInfo, ptyVInfo, supPubInfo, []byte{})

												key := make([]byte, size)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:53
	_, _ = reader.Read(key)
												return key
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:54
	// _ = "end of CoverTab[184342]"
}

func lengthPrefixed(data []byte) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:57
	_go_fuzz_dep_.CoverTab[184347]++
												out := make([]byte, len(data)+4)
												binary.BigEndian.PutUint32(out, uint32(len(data)))
												copy(out[4:], data)
												return out
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:61
	// _ = "end of CoverTab[184347]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:62
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/ecdh_es.go:62
var _ = _go_fuzz_dep_.CoverTab
