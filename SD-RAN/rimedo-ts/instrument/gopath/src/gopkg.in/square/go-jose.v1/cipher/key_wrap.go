//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:17
package josecipher

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:17
)

import (
	"crypto/cipher"
	"crypto/subtle"
	"encoding/binary"
	"errors"
)

var defaultIV = []byte{0xA6, 0xA6, 0xA6, 0xA6, 0xA6, 0xA6, 0xA6, 0xA6}

// KeyWrap implements NIST key wrapping; it wraps a content encryption key (cek) with the given block cipher.
func KeyWrap(block cipher.Block, cek []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:29
	_go_fuzz_dep_.CoverTab[184348]++
												if len(cek)%8 != 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:30
		_go_fuzz_dep_.CoverTab[184353]++
													return nil, errors.New("square/go-jose: key wrap input must be 8 byte blocks")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:31
		// _ = "end of CoverTab[184353]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:32
		_go_fuzz_dep_.CoverTab[184354]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:32
		// _ = "end of CoverTab[184354]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:32
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:32
	// _ = "end of CoverTab[184348]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:32
	_go_fuzz_dep_.CoverTab[184349]++

												n := len(cek) / 8
												r := make([][]byte, n)

												for i := range r {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:37
		_go_fuzz_dep_.CoverTab[184355]++
													r[i] = make([]byte, 8)
													copy(r[i], cek[i*8:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:39
		// _ = "end of CoverTab[184355]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:40
	// _ = "end of CoverTab[184349]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:40
	_go_fuzz_dep_.CoverTab[184350]++

												buffer := make([]byte, 16)
												tBytes := make([]byte, 8)
												copy(buffer, defaultIV)

												for t := 0; t < 6*n; t++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:46
		_go_fuzz_dep_.CoverTab[184356]++
													copy(buffer[8:], r[t%n])

													block.Encrypt(buffer, buffer)

													binary.BigEndian.PutUint64(tBytes, uint64(t+1))

													for i := 0; i < 8; i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:53
			_go_fuzz_dep_.CoverTab[184358]++
														buffer[i] = buffer[i] ^ tBytes[i]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:54
			// _ = "end of CoverTab[184358]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:55
		// _ = "end of CoverTab[184356]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:55
		_go_fuzz_dep_.CoverTab[184357]++
													copy(r[t%n], buffer[8:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:56
		// _ = "end of CoverTab[184357]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:57
	// _ = "end of CoverTab[184350]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:57
	_go_fuzz_dep_.CoverTab[184351]++

												out := make([]byte, (n+1)*8)
												copy(out, buffer[:8])
												for i := range r {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:61
		_go_fuzz_dep_.CoverTab[184359]++
													copy(out[(i+1)*8:], r[i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:62
		// _ = "end of CoverTab[184359]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:63
	// _ = "end of CoverTab[184351]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:63
	_go_fuzz_dep_.CoverTab[184352]++

												return out, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:65
	// _ = "end of CoverTab[184352]"
}

// KeyUnwrap implements NIST key unwrapping; it unwraps a content encryption key (cek) with the given block cipher.
func KeyUnwrap(block cipher.Block, ciphertext []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:69
	_go_fuzz_dep_.CoverTab[184360]++
												if len(ciphertext)%8 != 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:70
		_go_fuzz_dep_.CoverTab[184366]++
													return nil, errors.New("square/go-jose: key wrap input must be 8 byte blocks")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:71
		// _ = "end of CoverTab[184366]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:72
		_go_fuzz_dep_.CoverTab[184367]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:72
		// _ = "end of CoverTab[184367]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:72
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:72
	// _ = "end of CoverTab[184360]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:72
	_go_fuzz_dep_.CoverTab[184361]++

												n := (len(ciphertext) / 8) - 1
												r := make([][]byte, n)

												for i := range r {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:77
		_go_fuzz_dep_.CoverTab[184368]++
													r[i] = make([]byte, 8)
													copy(r[i], ciphertext[(i+1)*8:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:79
		// _ = "end of CoverTab[184368]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:80
	// _ = "end of CoverTab[184361]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:80
	_go_fuzz_dep_.CoverTab[184362]++

												buffer := make([]byte, 16)
												tBytes := make([]byte, 8)
												copy(buffer[:8], ciphertext[:8])

												for t := 6*n - 1; t >= 0; t-- {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:86
		_go_fuzz_dep_.CoverTab[184369]++
													binary.BigEndian.PutUint64(tBytes, uint64(t+1))

													for i := 0; i < 8; i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:89
			_go_fuzz_dep_.CoverTab[184371]++
														buffer[i] = buffer[i] ^ tBytes[i]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:90
			// _ = "end of CoverTab[184371]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:91
		// _ = "end of CoverTab[184369]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:91
		_go_fuzz_dep_.CoverTab[184370]++
													copy(buffer[8:], r[t%n])

													block.Decrypt(buffer, buffer)

													copy(r[t%n], buffer[8:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:96
		// _ = "end of CoverTab[184370]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:97
	// _ = "end of CoverTab[184362]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:97
	_go_fuzz_dep_.CoverTab[184363]++

												if subtle.ConstantTimeCompare(buffer[:8], defaultIV) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:99
		_go_fuzz_dep_.CoverTab[184372]++
													return nil, errors.New("square/go-jose: failed to unwrap key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:100
		// _ = "end of CoverTab[184372]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:101
		_go_fuzz_dep_.CoverTab[184373]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:101
		// _ = "end of CoverTab[184373]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:101
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:101
	// _ = "end of CoverTab[184363]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:101
	_go_fuzz_dep_.CoverTab[184364]++

												out := make([]byte, n*8)
												for i := range r {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:104
		_go_fuzz_dep_.CoverTab[184374]++
													copy(out[i*8:], r[i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:105
		// _ = "end of CoverTab[184374]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:106
	// _ = "end of CoverTab[184364]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:106
	_go_fuzz_dep_.CoverTab[184365]++

												return out, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:108
	// _ = "end of CoverTab[184365]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:109
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/cipher/key_wrap.go:109
var _ = _go_fuzz_dep_.CoverTab
