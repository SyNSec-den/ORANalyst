//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:17
package josecipher

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:17
)

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/binary"
)

// DeriveECDHES derives a shared encryption key using ECDH/ConcatKDF as described in JWE/JWA.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:27
// It is an error to call this function with a private/public key that are not on the same
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:27
// curve. Callers must ensure that the keys are valid before calling this function. Output
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:27
// size may be at most 1<<16 bytes (64 KiB).
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:31
func DeriveECDHES(alg string, apuData, apvData []byte, priv *ecdsa.PrivateKey, pub *ecdsa.PublicKey, size int) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:31
	_go_fuzz_dep_.CoverTab[187391]++
												if size > 1<<16 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:32
		_go_fuzz_dep_.CoverTab[187395]++
													panic("ECDH-ES output size too large, must be less than or equal to 1<<16")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:33
		// _ = "end of CoverTab[187395]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:34
		_go_fuzz_dep_.CoverTab[187396]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:34
		// _ = "end of CoverTab[187396]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:34
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:34
	// _ = "end of CoverTab[187391]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:34
	_go_fuzz_dep_.CoverTab[187392]++

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:37
	algID := lengthPrefixed([]byte(alg))
												ptyUInfo := lengthPrefixed(apuData)
												ptyVInfo := lengthPrefixed(apvData)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:42
	supPubInfo := make([]byte, 4)
	binary.BigEndian.PutUint32(supPubInfo, uint32(size)*8)

	if !priv.PublicKey.Curve.IsOnCurve(pub.X, pub.Y) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:45
		_go_fuzz_dep_.CoverTab[187397]++
													panic("public key not on same curve as private key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:46
		// _ = "end of CoverTab[187397]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:47
		_go_fuzz_dep_.CoverTab[187398]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:47
		// _ = "end of CoverTab[187398]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:47
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:47
	// _ = "end of CoverTab[187392]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:47
	_go_fuzz_dep_.CoverTab[187393]++

												z, _ := priv.Curve.ScalarMult(pub.X, pub.Y, priv.D.Bytes())
												zBytes := z.Bytes()

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:56
	octSize := dSize(priv.Curve)
	if len(zBytes) != octSize {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:57
		_go_fuzz_dep_.CoverTab[187399]++
													zBytes = append(bytes.Repeat([]byte{0}, octSize-len(zBytes)), zBytes...)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:58
		// _ = "end of CoverTab[187399]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:59
		_go_fuzz_dep_.CoverTab[187400]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:59
		// _ = "end of CoverTab[187400]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:59
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:59
	// _ = "end of CoverTab[187393]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:59
	_go_fuzz_dep_.CoverTab[187394]++

												reader := NewConcatKDF(crypto.SHA256, zBytes, algID, ptyUInfo, ptyVInfo, supPubInfo, []byte{})
												key := make([]byte, size)

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:65
	_, _ = reader.Read(key)

												return key
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:67
	// _ = "end of CoverTab[187394]"
}

// dSize returns the size in octets for a coordinate on a elliptic curve.
func dSize(curve elliptic.Curve) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:71
	_go_fuzz_dep_.CoverTab[187401]++
												order := curve.Params().P
												bitLen := order.BitLen()
												size := bitLen / 8
												if bitLen%8 != 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:75
		_go_fuzz_dep_.CoverTab[187403]++
													size++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:76
		// _ = "end of CoverTab[187403]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:77
		_go_fuzz_dep_.CoverTab[187404]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:77
		// _ = "end of CoverTab[187404]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:77
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:77
	// _ = "end of CoverTab[187401]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:77
	_go_fuzz_dep_.CoverTab[187402]++
												return size
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:78
	// _ = "end of CoverTab[187402]"
}

func lengthPrefixed(data []byte) []byte {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:81
	_go_fuzz_dep_.CoverTab[187405]++
												out := make([]byte, len(data)+4)
												binary.BigEndian.PutUint32(out, uint32(len(data)))
												copy(out[4:], data)
												return out
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:85
	// _ = "end of CoverTab[187405]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:86
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/ecdh_es.go:86
var _ = _go_fuzz_dep_.CoverTab
