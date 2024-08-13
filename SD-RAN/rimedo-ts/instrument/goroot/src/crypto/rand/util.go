// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/rand/util.go:5
package rand

//line /usr/local/go/src/crypto/rand/util.go:5
import (
//line /usr/local/go/src/crypto/rand/util.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/rand/util.go:5
)
//line /usr/local/go/src/crypto/rand/util.go:5
import (
//line /usr/local/go/src/crypto/rand/util.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/rand/util.go:5
)

import (
	"crypto/internal/randutil"
	"errors"
	"io"
	"math/big"
)

// Prime returns a number of the given bit length that is prime with high probability.
//line /usr/local/go/src/crypto/rand/util.go:14
// Prime will return error for any error returned by rand.Read or if bits < 2.
//line /usr/local/go/src/crypto/rand/util.go:16
func Prime(rand io.Reader, bits int) (*big.Int, error) {
//line /usr/local/go/src/crypto/rand/util.go:16
	_go_fuzz_dep_.CoverTab[9361]++
							if bits < 2 {
//line /usr/local/go/src/crypto/rand/util.go:17
		_go_fuzz_dep_.CoverTab[9364]++
								return nil, errors.New("crypto/rand: prime size must be at least 2-bit")
//line /usr/local/go/src/crypto/rand/util.go:18
		// _ = "end of CoverTab[9364]"
	} else {
//line /usr/local/go/src/crypto/rand/util.go:19
		_go_fuzz_dep_.CoverTab[9365]++
//line /usr/local/go/src/crypto/rand/util.go:19
		// _ = "end of CoverTab[9365]"
//line /usr/local/go/src/crypto/rand/util.go:19
	}
//line /usr/local/go/src/crypto/rand/util.go:19
	// _ = "end of CoverTab[9361]"
//line /usr/local/go/src/crypto/rand/util.go:19
	_go_fuzz_dep_.CoverTab[9362]++

							randutil.MaybeReadByte(rand)

							b := uint(bits % 8)
							if b == 0 {
//line /usr/local/go/src/crypto/rand/util.go:24
		_go_fuzz_dep_.CoverTab[9366]++
								b = 8
//line /usr/local/go/src/crypto/rand/util.go:25
		// _ = "end of CoverTab[9366]"
	} else {
//line /usr/local/go/src/crypto/rand/util.go:26
		_go_fuzz_dep_.CoverTab[9367]++
//line /usr/local/go/src/crypto/rand/util.go:26
		// _ = "end of CoverTab[9367]"
//line /usr/local/go/src/crypto/rand/util.go:26
	}
//line /usr/local/go/src/crypto/rand/util.go:26
	// _ = "end of CoverTab[9362]"
//line /usr/local/go/src/crypto/rand/util.go:26
	_go_fuzz_dep_.CoverTab[9363]++

							bytes := make([]byte, (bits+7)/8)
							p := new(big.Int)

							for {
//line /usr/local/go/src/crypto/rand/util.go:31
		_go_fuzz_dep_.CoverTab[9368]++
								if _, err := io.ReadFull(rand, bytes); err != nil {
//line /usr/local/go/src/crypto/rand/util.go:32
			_go_fuzz_dep_.CoverTab[9371]++
									return nil, err
//line /usr/local/go/src/crypto/rand/util.go:33
			// _ = "end of CoverTab[9371]"
		} else {
//line /usr/local/go/src/crypto/rand/util.go:34
			_go_fuzz_dep_.CoverTab[9372]++
//line /usr/local/go/src/crypto/rand/util.go:34
			// _ = "end of CoverTab[9372]"
//line /usr/local/go/src/crypto/rand/util.go:34
		}
//line /usr/local/go/src/crypto/rand/util.go:34
		// _ = "end of CoverTab[9368]"
//line /usr/local/go/src/crypto/rand/util.go:34
		_go_fuzz_dep_.CoverTab[9369]++

//line /usr/local/go/src/crypto/rand/util.go:37
		bytes[0] &= uint8(int(1<<b) - 1)

//line /usr/local/go/src/crypto/rand/util.go:42
		if b >= 2 {
//line /usr/local/go/src/crypto/rand/util.go:42
			_go_fuzz_dep_.CoverTab[9373]++
									bytes[0] |= 3 << (b - 2)
//line /usr/local/go/src/crypto/rand/util.go:43
			// _ = "end of CoverTab[9373]"
		} else {
//line /usr/local/go/src/crypto/rand/util.go:44
			_go_fuzz_dep_.CoverTab[9374]++

									bytes[0] |= 1
									if len(bytes) > 1 {
//line /usr/local/go/src/crypto/rand/util.go:47
				_go_fuzz_dep_.CoverTab[9375]++
										bytes[1] |= 0x80
//line /usr/local/go/src/crypto/rand/util.go:48
				// _ = "end of CoverTab[9375]"
			} else {
//line /usr/local/go/src/crypto/rand/util.go:49
				_go_fuzz_dep_.CoverTab[9376]++
//line /usr/local/go/src/crypto/rand/util.go:49
				// _ = "end of CoverTab[9376]"
//line /usr/local/go/src/crypto/rand/util.go:49
			}
//line /usr/local/go/src/crypto/rand/util.go:49
			// _ = "end of CoverTab[9374]"
		}
//line /usr/local/go/src/crypto/rand/util.go:50
		// _ = "end of CoverTab[9369]"
//line /usr/local/go/src/crypto/rand/util.go:50
		_go_fuzz_dep_.CoverTab[9370]++

								bytes[len(bytes)-1] |= 1

								p.SetBytes(bytes)
								if p.ProbablyPrime(20) {
//line /usr/local/go/src/crypto/rand/util.go:55
			_go_fuzz_dep_.CoverTab[9377]++
									return p, nil
//line /usr/local/go/src/crypto/rand/util.go:56
			// _ = "end of CoverTab[9377]"
		} else {
//line /usr/local/go/src/crypto/rand/util.go:57
			_go_fuzz_dep_.CoverTab[9378]++
//line /usr/local/go/src/crypto/rand/util.go:57
			// _ = "end of CoverTab[9378]"
//line /usr/local/go/src/crypto/rand/util.go:57
		}
//line /usr/local/go/src/crypto/rand/util.go:57
		// _ = "end of CoverTab[9370]"
	}
//line /usr/local/go/src/crypto/rand/util.go:58
	// _ = "end of CoverTab[9363]"
}

// Int returns a uniform random value in [0, max). It panics if max <= 0.
func Int(rand io.Reader, max *big.Int) (n *big.Int, err error) {
//line /usr/local/go/src/crypto/rand/util.go:62
	_go_fuzz_dep_.CoverTab[9379]++
							if max.Sign() <= 0 {
//line /usr/local/go/src/crypto/rand/util.go:63
		_go_fuzz_dep_.CoverTab[9383]++
								panic("crypto/rand: argument to Int is <= 0")
//line /usr/local/go/src/crypto/rand/util.go:64
		// _ = "end of CoverTab[9383]"
	} else {
//line /usr/local/go/src/crypto/rand/util.go:65
		_go_fuzz_dep_.CoverTab[9384]++
//line /usr/local/go/src/crypto/rand/util.go:65
		// _ = "end of CoverTab[9384]"
//line /usr/local/go/src/crypto/rand/util.go:65
	}
//line /usr/local/go/src/crypto/rand/util.go:65
	// _ = "end of CoverTab[9379]"
//line /usr/local/go/src/crypto/rand/util.go:65
	_go_fuzz_dep_.CoverTab[9380]++
							n = new(big.Int)
							n.Sub(max, n.SetUint64(1))

							bitLen := n.BitLen()
							if bitLen == 0 {
//line /usr/local/go/src/crypto/rand/util.go:70
		_go_fuzz_dep_.CoverTab[9385]++

								return
//line /usr/local/go/src/crypto/rand/util.go:72
		// _ = "end of CoverTab[9385]"
	} else {
//line /usr/local/go/src/crypto/rand/util.go:73
		_go_fuzz_dep_.CoverTab[9386]++
//line /usr/local/go/src/crypto/rand/util.go:73
		// _ = "end of CoverTab[9386]"
//line /usr/local/go/src/crypto/rand/util.go:73
	}
//line /usr/local/go/src/crypto/rand/util.go:73
	// _ = "end of CoverTab[9380]"
//line /usr/local/go/src/crypto/rand/util.go:73
	_go_fuzz_dep_.CoverTab[9381]++

							k := (bitLen + 7) / 8

							b := uint(bitLen % 8)
							if b == 0 {
//line /usr/local/go/src/crypto/rand/util.go:78
		_go_fuzz_dep_.CoverTab[9387]++
								b = 8
//line /usr/local/go/src/crypto/rand/util.go:79
		// _ = "end of CoverTab[9387]"
	} else {
//line /usr/local/go/src/crypto/rand/util.go:80
		_go_fuzz_dep_.CoverTab[9388]++
//line /usr/local/go/src/crypto/rand/util.go:80
		// _ = "end of CoverTab[9388]"
//line /usr/local/go/src/crypto/rand/util.go:80
	}
//line /usr/local/go/src/crypto/rand/util.go:80
	// _ = "end of CoverTab[9381]"
//line /usr/local/go/src/crypto/rand/util.go:80
	_go_fuzz_dep_.CoverTab[9382]++

							bytes := make([]byte, k)

							for {
//line /usr/local/go/src/crypto/rand/util.go:84
		_go_fuzz_dep_.CoverTab[9389]++
								_, err = io.ReadFull(rand, bytes)
								if err != nil {
//line /usr/local/go/src/crypto/rand/util.go:86
			_go_fuzz_dep_.CoverTab[9391]++
									return nil, err
//line /usr/local/go/src/crypto/rand/util.go:87
			// _ = "end of CoverTab[9391]"
		} else {
//line /usr/local/go/src/crypto/rand/util.go:88
			_go_fuzz_dep_.CoverTab[9392]++
//line /usr/local/go/src/crypto/rand/util.go:88
			// _ = "end of CoverTab[9392]"
//line /usr/local/go/src/crypto/rand/util.go:88
		}
//line /usr/local/go/src/crypto/rand/util.go:88
		// _ = "end of CoverTab[9389]"
//line /usr/local/go/src/crypto/rand/util.go:88
		_go_fuzz_dep_.CoverTab[9390]++

//line /usr/local/go/src/crypto/rand/util.go:92
		bytes[0] &= uint8(int(1<<b) - 1)

		n.SetBytes(bytes)
		if n.Cmp(max) < 0 {
//line /usr/local/go/src/crypto/rand/util.go:95
			_go_fuzz_dep_.CoverTab[9393]++
									return
//line /usr/local/go/src/crypto/rand/util.go:96
			// _ = "end of CoverTab[9393]"
		} else {
//line /usr/local/go/src/crypto/rand/util.go:97
			_go_fuzz_dep_.CoverTab[9394]++
//line /usr/local/go/src/crypto/rand/util.go:97
			// _ = "end of CoverTab[9394]"
//line /usr/local/go/src/crypto/rand/util.go:97
		}
//line /usr/local/go/src/crypto/rand/util.go:97
		// _ = "end of CoverTab[9390]"
	}
//line /usr/local/go/src/crypto/rand/util.go:98
	// _ = "end of CoverTab[9382]"
}

//line /usr/local/go/src/crypto/rand/util.go:99
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/rand/util.go:99
var _ = _go_fuzz_dep_.CoverTab
