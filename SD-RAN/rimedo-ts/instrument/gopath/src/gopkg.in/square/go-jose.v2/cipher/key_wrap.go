//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:17
package josecipher

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:17
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:17
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:17
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
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:29
	_go_fuzz_dep_.CoverTab[187406]++
												if len(cek)%8 != 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:30
		_go_fuzz_dep_.CoverTab[187411]++
													return nil, errors.New("square/go-jose: key wrap input must be 8 byte blocks")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:31
		// _ = "end of CoverTab[187411]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:32
		_go_fuzz_dep_.CoverTab[187412]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:32
		// _ = "end of CoverTab[187412]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:32
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:32
	// _ = "end of CoverTab[187406]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:32
	_go_fuzz_dep_.CoverTab[187407]++

												n := len(cek) / 8
												r := make([][]byte, n)

												for i := range r {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:37
		_go_fuzz_dep_.CoverTab[187413]++
													r[i] = make([]byte, 8)
													copy(r[i], cek[i*8:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:39
		// _ = "end of CoverTab[187413]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:40
	// _ = "end of CoverTab[187407]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:40
	_go_fuzz_dep_.CoverTab[187408]++

												buffer := make([]byte, 16)
												tBytes := make([]byte, 8)
												copy(buffer, defaultIV)

												for t := 0; t < 6*n; t++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:46
		_go_fuzz_dep_.CoverTab[187414]++
													copy(buffer[8:], r[t%n])

													block.Encrypt(buffer, buffer)

													binary.BigEndian.PutUint64(tBytes, uint64(t+1))

													for i := 0; i < 8; i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:53
			_go_fuzz_dep_.CoverTab[187416]++
														buffer[i] = buffer[i] ^ tBytes[i]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:54
			// _ = "end of CoverTab[187416]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:55
		// _ = "end of CoverTab[187414]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:55
		_go_fuzz_dep_.CoverTab[187415]++
													copy(r[t%n], buffer[8:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:56
		// _ = "end of CoverTab[187415]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:57
	// _ = "end of CoverTab[187408]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:57
	_go_fuzz_dep_.CoverTab[187409]++

												out := make([]byte, (n+1)*8)
												copy(out, buffer[:8])
												for i := range r {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:61
		_go_fuzz_dep_.CoverTab[187417]++
													copy(out[(i+1)*8:], r[i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:62
		// _ = "end of CoverTab[187417]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:63
	// _ = "end of CoverTab[187409]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:63
	_go_fuzz_dep_.CoverTab[187410]++

												return out, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:65
	// _ = "end of CoverTab[187410]"
}

// KeyUnwrap implements NIST key unwrapping; it unwraps a content encryption key (cek) with the given block cipher.
func KeyUnwrap(block cipher.Block, ciphertext []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:69
	_go_fuzz_dep_.CoverTab[187418]++
												if len(ciphertext)%8 != 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:70
		_go_fuzz_dep_.CoverTab[187424]++
													return nil, errors.New("square/go-jose: key wrap input must be 8 byte blocks")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:71
		// _ = "end of CoverTab[187424]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:72
		_go_fuzz_dep_.CoverTab[187425]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:72
		// _ = "end of CoverTab[187425]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:72
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:72
	// _ = "end of CoverTab[187418]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:72
	_go_fuzz_dep_.CoverTab[187419]++

												n := (len(ciphertext) / 8) - 1
												r := make([][]byte, n)

												for i := range r {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:77
		_go_fuzz_dep_.CoverTab[187426]++
													r[i] = make([]byte, 8)
													copy(r[i], ciphertext[(i+1)*8:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:79
		// _ = "end of CoverTab[187426]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:80
	// _ = "end of CoverTab[187419]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:80
	_go_fuzz_dep_.CoverTab[187420]++

												buffer := make([]byte, 16)
												tBytes := make([]byte, 8)
												copy(buffer[:8], ciphertext[:8])

												for t := 6*n - 1; t >= 0; t-- {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:86
		_go_fuzz_dep_.CoverTab[187427]++
													binary.BigEndian.PutUint64(tBytes, uint64(t+1))

													for i := 0; i < 8; i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:89
			_go_fuzz_dep_.CoverTab[187429]++
														buffer[i] = buffer[i] ^ tBytes[i]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:90
			// _ = "end of CoverTab[187429]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:91
		// _ = "end of CoverTab[187427]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:91
		_go_fuzz_dep_.CoverTab[187428]++
													copy(buffer[8:], r[t%n])

													block.Decrypt(buffer, buffer)

													copy(r[t%n], buffer[8:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:96
		// _ = "end of CoverTab[187428]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:97
	// _ = "end of CoverTab[187420]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:97
	_go_fuzz_dep_.CoverTab[187421]++

												if subtle.ConstantTimeCompare(buffer[:8], defaultIV) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:99
		_go_fuzz_dep_.CoverTab[187430]++
													return nil, errors.New("square/go-jose: failed to unwrap key")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:100
		// _ = "end of CoverTab[187430]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:101
		_go_fuzz_dep_.CoverTab[187431]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:101
		// _ = "end of CoverTab[187431]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:101
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:101
	// _ = "end of CoverTab[187421]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:101
	_go_fuzz_dep_.CoverTab[187422]++

												out := make([]byte, n*8)
												for i := range r {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:104
		_go_fuzz_dep_.CoverTab[187432]++
													copy(out[i*8:], r[i])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:105
		// _ = "end of CoverTab[187432]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:106
	// _ = "end of CoverTab[187422]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:106
	_go_fuzz_dep_.CoverTab[187423]++

												return out, nil
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:108
	// _ = "end of CoverTab[187423]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:109
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v2@v2.5.1/cipher/key_wrap.go:109
var _ = _go_fuzz_dep_.CoverTab
