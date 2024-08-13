// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/hmac/hmac.go:5
/*
Package hmac implements the Keyed-Hash Message Authentication Code (HMAC) as
defined in U.S. Federal Information Processing Standards Publication 198.
An HMAC is a cryptographic hash that uses a key to sign a message.
The receiver verifies the hash by recomputing it using the same key.

Receivers should be careful to use Equal to compare MACs in order to avoid
timing side-channels:

	// ValidMAC reports whether messageMAC is a valid HMAC tag for message.
	func ValidMAC(message, messageMAC, key []byte) bool {
		mac := hmac.New(sha256.New, key)
		mac.Write(message)
		expectedMAC := mac.Sum(nil)
		return hmac.Equal(messageMAC, expectedMAC)
	}
*/
package hmac

//line /usr/local/go/src/crypto/hmac/hmac.go:22
import (
//line /usr/local/go/src/crypto/hmac/hmac.go:22
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/hmac/hmac.go:22
)
//line /usr/local/go/src/crypto/hmac/hmac.go:22
import (
//line /usr/local/go/src/crypto/hmac/hmac.go:22
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/hmac/hmac.go:22
)

import (
	"crypto/internal/boring"
	"crypto/subtle"
	"hash"
)

//line /usr/local/go/src/crypto/hmac/hmac.go:38
// Marshalable is the combination of encoding.BinaryMarshaler and
//line /usr/local/go/src/crypto/hmac/hmac.go:38
// encoding.BinaryUnmarshaler. Their method definitions are repeated here to
//line /usr/local/go/src/crypto/hmac/hmac.go:38
// avoid a dependency on the encoding package.
//line /usr/local/go/src/crypto/hmac/hmac.go:41
type marshalable interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
}

type hmac struct {
	opad, ipad	[]byte
	outer, inner	hash.Hash

	// If marshaled is true, then opad and ipad do not contain a padded
	// copy of the key, but rather the marshaled state of outer/inner after
	// opad/ipad has been fed into it.
	marshaled	bool
}

func (h *hmac) Sum(in []byte) []byte {
//line /usr/local/go/src/crypto/hmac/hmac.go:56
	_go_fuzz_dep_.CoverTab[9501]++
							origLen := len(in)
							in = h.inner.Sum(in)

							if h.marshaled {
//line /usr/local/go/src/crypto/hmac/hmac.go:60
		_go_fuzz_dep_.CoverTab[9503]++
								if err := h.outer.(marshalable).UnmarshalBinary(h.opad); err != nil {
//line /usr/local/go/src/crypto/hmac/hmac.go:61
			_go_fuzz_dep_.CoverTab[9504]++
									panic(err)
//line /usr/local/go/src/crypto/hmac/hmac.go:62
			// _ = "end of CoverTab[9504]"
		} else {
//line /usr/local/go/src/crypto/hmac/hmac.go:63
			_go_fuzz_dep_.CoverTab[9505]++
//line /usr/local/go/src/crypto/hmac/hmac.go:63
			// _ = "end of CoverTab[9505]"
//line /usr/local/go/src/crypto/hmac/hmac.go:63
		}
//line /usr/local/go/src/crypto/hmac/hmac.go:63
		// _ = "end of CoverTab[9503]"
	} else {
//line /usr/local/go/src/crypto/hmac/hmac.go:64
		_go_fuzz_dep_.CoverTab[9506]++
								h.outer.Reset()
								h.outer.Write(h.opad)
//line /usr/local/go/src/crypto/hmac/hmac.go:66
		// _ = "end of CoverTab[9506]"
	}
//line /usr/local/go/src/crypto/hmac/hmac.go:67
	// _ = "end of CoverTab[9501]"
//line /usr/local/go/src/crypto/hmac/hmac.go:67
	_go_fuzz_dep_.CoverTab[9502]++
							h.outer.Write(in[origLen:])
							return h.outer.Sum(in[:origLen])
//line /usr/local/go/src/crypto/hmac/hmac.go:69
	// _ = "end of CoverTab[9502]"
}

func (h *hmac) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/crypto/hmac/hmac.go:72
	_go_fuzz_dep_.CoverTab[9507]++
							return h.inner.Write(p)
//line /usr/local/go/src/crypto/hmac/hmac.go:73
	// _ = "end of CoverTab[9507]"
}

func (h *hmac) Size() int {
//line /usr/local/go/src/crypto/hmac/hmac.go:76
	_go_fuzz_dep_.CoverTab[9508]++
//line /usr/local/go/src/crypto/hmac/hmac.go:76
	return h.outer.Size()
//line /usr/local/go/src/crypto/hmac/hmac.go:76
	// _ = "end of CoverTab[9508]"
//line /usr/local/go/src/crypto/hmac/hmac.go:76
}
func (h *hmac) BlockSize() int {
//line /usr/local/go/src/crypto/hmac/hmac.go:77
	_go_fuzz_dep_.CoverTab[9509]++
//line /usr/local/go/src/crypto/hmac/hmac.go:77
	return h.inner.BlockSize()
//line /usr/local/go/src/crypto/hmac/hmac.go:77
	// _ = "end of CoverTab[9509]"
//line /usr/local/go/src/crypto/hmac/hmac.go:77
}

func (h *hmac) Reset() {
//line /usr/local/go/src/crypto/hmac/hmac.go:79
	_go_fuzz_dep_.CoverTab[9510]++
							if h.marshaled {
//line /usr/local/go/src/crypto/hmac/hmac.go:80
		_go_fuzz_dep_.CoverTab[9516]++
								if err := h.inner.(marshalable).UnmarshalBinary(h.ipad); err != nil {
//line /usr/local/go/src/crypto/hmac/hmac.go:81
			_go_fuzz_dep_.CoverTab[9518]++
									panic(err)
//line /usr/local/go/src/crypto/hmac/hmac.go:82
			// _ = "end of CoverTab[9518]"
		} else {
//line /usr/local/go/src/crypto/hmac/hmac.go:83
			_go_fuzz_dep_.CoverTab[9519]++
//line /usr/local/go/src/crypto/hmac/hmac.go:83
			// _ = "end of CoverTab[9519]"
//line /usr/local/go/src/crypto/hmac/hmac.go:83
		}
//line /usr/local/go/src/crypto/hmac/hmac.go:83
		// _ = "end of CoverTab[9516]"
//line /usr/local/go/src/crypto/hmac/hmac.go:83
		_go_fuzz_dep_.CoverTab[9517]++
								return
//line /usr/local/go/src/crypto/hmac/hmac.go:84
		// _ = "end of CoverTab[9517]"
	} else {
//line /usr/local/go/src/crypto/hmac/hmac.go:85
		_go_fuzz_dep_.CoverTab[9520]++
//line /usr/local/go/src/crypto/hmac/hmac.go:85
		// _ = "end of CoverTab[9520]"
//line /usr/local/go/src/crypto/hmac/hmac.go:85
	}
//line /usr/local/go/src/crypto/hmac/hmac.go:85
	// _ = "end of CoverTab[9510]"
//line /usr/local/go/src/crypto/hmac/hmac.go:85
	_go_fuzz_dep_.CoverTab[9511]++

							h.inner.Reset()
							h.inner.Write(h.ipad)

//line /usr/local/go/src/crypto/hmac/hmac.go:96
	marshalableInner, innerOK := h.inner.(marshalable)
	if !innerOK {
//line /usr/local/go/src/crypto/hmac/hmac.go:97
		_go_fuzz_dep_.CoverTab[9521]++
								return
//line /usr/local/go/src/crypto/hmac/hmac.go:98
		// _ = "end of CoverTab[9521]"
	} else {
//line /usr/local/go/src/crypto/hmac/hmac.go:99
		_go_fuzz_dep_.CoverTab[9522]++
//line /usr/local/go/src/crypto/hmac/hmac.go:99
		// _ = "end of CoverTab[9522]"
//line /usr/local/go/src/crypto/hmac/hmac.go:99
	}
//line /usr/local/go/src/crypto/hmac/hmac.go:99
	// _ = "end of CoverTab[9511]"
//line /usr/local/go/src/crypto/hmac/hmac.go:99
	_go_fuzz_dep_.CoverTab[9512]++
							marshalableOuter, outerOK := h.outer.(marshalable)
							if !outerOK {
//line /usr/local/go/src/crypto/hmac/hmac.go:101
		_go_fuzz_dep_.CoverTab[9523]++
								return
//line /usr/local/go/src/crypto/hmac/hmac.go:102
		// _ = "end of CoverTab[9523]"
	} else {
//line /usr/local/go/src/crypto/hmac/hmac.go:103
		_go_fuzz_dep_.CoverTab[9524]++
//line /usr/local/go/src/crypto/hmac/hmac.go:103
		// _ = "end of CoverTab[9524]"
//line /usr/local/go/src/crypto/hmac/hmac.go:103
	}
//line /usr/local/go/src/crypto/hmac/hmac.go:103
	// _ = "end of CoverTab[9512]"
//line /usr/local/go/src/crypto/hmac/hmac.go:103
	_go_fuzz_dep_.CoverTab[9513]++

							imarshal, err := marshalableInner.MarshalBinary()
							if err != nil {
//line /usr/local/go/src/crypto/hmac/hmac.go:106
		_go_fuzz_dep_.CoverTab[9525]++
								return
//line /usr/local/go/src/crypto/hmac/hmac.go:107
		// _ = "end of CoverTab[9525]"
	} else {
//line /usr/local/go/src/crypto/hmac/hmac.go:108
		_go_fuzz_dep_.CoverTab[9526]++
//line /usr/local/go/src/crypto/hmac/hmac.go:108
		// _ = "end of CoverTab[9526]"
//line /usr/local/go/src/crypto/hmac/hmac.go:108
	}
//line /usr/local/go/src/crypto/hmac/hmac.go:108
	// _ = "end of CoverTab[9513]"
//line /usr/local/go/src/crypto/hmac/hmac.go:108
	_go_fuzz_dep_.CoverTab[9514]++

							h.outer.Reset()
							h.outer.Write(h.opad)
							omarshal, err := marshalableOuter.MarshalBinary()
							if err != nil {
//line /usr/local/go/src/crypto/hmac/hmac.go:113
		_go_fuzz_dep_.CoverTab[9527]++
								return
//line /usr/local/go/src/crypto/hmac/hmac.go:114
		// _ = "end of CoverTab[9527]"
	} else {
//line /usr/local/go/src/crypto/hmac/hmac.go:115
		_go_fuzz_dep_.CoverTab[9528]++
//line /usr/local/go/src/crypto/hmac/hmac.go:115
		// _ = "end of CoverTab[9528]"
//line /usr/local/go/src/crypto/hmac/hmac.go:115
	}
//line /usr/local/go/src/crypto/hmac/hmac.go:115
	// _ = "end of CoverTab[9514]"
//line /usr/local/go/src/crypto/hmac/hmac.go:115
	_go_fuzz_dep_.CoverTab[9515]++

//line /usr/local/go/src/crypto/hmac/hmac.go:118
	h.ipad = imarshal
							h.opad = omarshal
							h.marshaled = true
//line /usr/local/go/src/crypto/hmac/hmac.go:120
	// _ = "end of CoverTab[9515]"
}

// New returns a new HMAC hash using the given hash.Hash type and key.
//line /usr/local/go/src/crypto/hmac/hmac.go:123
// New functions like sha256.New from crypto/sha256 can be used as h.
//line /usr/local/go/src/crypto/hmac/hmac.go:123
// h must return a new Hash every time it is called.
//line /usr/local/go/src/crypto/hmac/hmac.go:123
// Note that unlike other hash implementations in the standard library,
//line /usr/local/go/src/crypto/hmac/hmac.go:123
// the returned Hash does not implement encoding.BinaryMarshaler
//line /usr/local/go/src/crypto/hmac/hmac.go:123
// or encoding.BinaryUnmarshaler.
//line /usr/local/go/src/crypto/hmac/hmac.go:129
func New(h func() hash.Hash, key []byte) hash.Hash {
//line /usr/local/go/src/crypto/hmac/hmac.go:129
	_go_fuzz_dep_.CoverTab[9529]++
							if boring.Enabled {
//line /usr/local/go/src/crypto/hmac/hmac.go:130
		_go_fuzz_dep_.CoverTab[9536]++
								hm := boring.NewHMAC(h, key)
								if hm != nil {
//line /usr/local/go/src/crypto/hmac/hmac.go:132
			_go_fuzz_dep_.CoverTab[9537]++
									return hm
//line /usr/local/go/src/crypto/hmac/hmac.go:133
			// _ = "end of CoverTab[9537]"
		} else {
//line /usr/local/go/src/crypto/hmac/hmac.go:134
			_go_fuzz_dep_.CoverTab[9538]++
//line /usr/local/go/src/crypto/hmac/hmac.go:134
			// _ = "end of CoverTab[9538]"
//line /usr/local/go/src/crypto/hmac/hmac.go:134
		}
//line /usr/local/go/src/crypto/hmac/hmac.go:134
		// _ = "end of CoverTab[9536]"

	} else {
//line /usr/local/go/src/crypto/hmac/hmac.go:136
		_go_fuzz_dep_.CoverTab[9539]++
//line /usr/local/go/src/crypto/hmac/hmac.go:136
		// _ = "end of CoverTab[9539]"
//line /usr/local/go/src/crypto/hmac/hmac.go:136
	}
//line /usr/local/go/src/crypto/hmac/hmac.go:136
	// _ = "end of CoverTab[9529]"
//line /usr/local/go/src/crypto/hmac/hmac.go:136
	_go_fuzz_dep_.CoverTab[9530]++
							hm := new(hmac)
							hm.outer = h()
							hm.inner = h()
							unique := true
							func() {
//line /usr/local/go/src/crypto/hmac/hmac.go:141
		_go_fuzz_dep_.CoverTab[9540]++
								defer func() {
//line /usr/local/go/src/crypto/hmac/hmac.go:142
			_go_fuzz_dep_.CoverTab[9542]++

									_ = recover()
//line /usr/local/go/src/crypto/hmac/hmac.go:144
			// _ = "end of CoverTab[9542]"
		}()
//line /usr/local/go/src/crypto/hmac/hmac.go:145
		// _ = "end of CoverTab[9540]"
//line /usr/local/go/src/crypto/hmac/hmac.go:145
		_go_fuzz_dep_.CoverTab[9541]++
								if hm.outer == hm.inner {
//line /usr/local/go/src/crypto/hmac/hmac.go:146
			_go_fuzz_dep_.CoverTab[9543]++
									unique = false
//line /usr/local/go/src/crypto/hmac/hmac.go:147
			// _ = "end of CoverTab[9543]"
		} else {
//line /usr/local/go/src/crypto/hmac/hmac.go:148
			_go_fuzz_dep_.CoverTab[9544]++
//line /usr/local/go/src/crypto/hmac/hmac.go:148
			// _ = "end of CoverTab[9544]"
//line /usr/local/go/src/crypto/hmac/hmac.go:148
		}
//line /usr/local/go/src/crypto/hmac/hmac.go:148
		// _ = "end of CoverTab[9541]"
	}()
//line /usr/local/go/src/crypto/hmac/hmac.go:149
	// _ = "end of CoverTab[9530]"
//line /usr/local/go/src/crypto/hmac/hmac.go:149
	_go_fuzz_dep_.CoverTab[9531]++
							if !unique {
//line /usr/local/go/src/crypto/hmac/hmac.go:150
		_go_fuzz_dep_.CoverTab[9545]++
								panic("crypto/hmac: hash generation function does not produce unique values")
//line /usr/local/go/src/crypto/hmac/hmac.go:151
		// _ = "end of CoverTab[9545]"
	} else {
//line /usr/local/go/src/crypto/hmac/hmac.go:152
		_go_fuzz_dep_.CoverTab[9546]++
//line /usr/local/go/src/crypto/hmac/hmac.go:152
		// _ = "end of CoverTab[9546]"
//line /usr/local/go/src/crypto/hmac/hmac.go:152
	}
//line /usr/local/go/src/crypto/hmac/hmac.go:152
	// _ = "end of CoverTab[9531]"
//line /usr/local/go/src/crypto/hmac/hmac.go:152
	_go_fuzz_dep_.CoverTab[9532]++
							blocksize := hm.inner.BlockSize()
							hm.ipad = make([]byte, blocksize)
							hm.opad = make([]byte, blocksize)
							if len(key) > blocksize {
//line /usr/local/go/src/crypto/hmac/hmac.go:156
		_go_fuzz_dep_.CoverTab[9547]++

								hm.outer.Write(key)
								key = hm.outer.Sum(nil)
//line /usr/local/go/src/crypto/hmac/hmac.go:159
		// _ = "end of CoverTab[9547]"
	} else {
//line /usr/local/go/src/crypto/hmac/hmac.go:160
		_go_fuzz_dep_.CoverTab[9548]++
//line /usr/local/go/src/crypto/hmac/hmac.go:160
		// _ = "end of CoverTab[9548]"
//line /usr/local/go/src/crypto/hmac/hmac.go:160
	}
//line /usr/local/go/src/crypto/hmac/hmac.go:160
	// _ = "end of CoverTab[9532]"
//line /usr/local/go/src/crypto/hmac/hmac.go:160
	_go_fuzz_dep_.CoverTab[9533]++
							copy(hm.ipad, key)
							copy(hm.opad, key)
							for i := range hm.ipad {
//line /usr/local/go/src/crypto/hmac/hmac.go:163
		_go_fuzz_dep_.CoverTab[9549]++
								hm.ipad[i] ^= 0x36
//line /usr/local/go/src/crypto/hmac/hmac.go:164
		// _ = "end of CoverTab[9549]"
	}
//line /usr/local/go/src/crypto/hmac/hmac.go:165
	// _ = "end of CoverTab[9533]"
//line /usr/local/go/src/crypto/hmac/hmac.go:165
	_go_fuzz_dep_.CoverTab[9534]++
							for i := range hm.opad {
//line /usr/local/go/src/crypto/hmac/hmac.go:166
		_go_fuzz_dep_.CoverTab[9550]++
								hm.opad[i] ^= 0x5c
//line /usr/local/go/src/crypto/hmac/hmac.go:167
		// _ = "end of CoverTab[9550]"
	}
//line /usr/local/go/src/crypto/hmac/hmac.go:168
	// _ = "end of CoverTab[9534]"
//line /usr/local/go/src/crypto/hmac/hmac.go:168
	_go_fuzz_dep_.CoverTab[9535]++
							hm.inner.Write(hm.ipad)

							return hm
//line /usr/local/go/src/crypto/hmac/hmac.go:171
	// _ = "end of CoverTab[9535]"
}

// Equal compares two MACs for equality without leaking timing information.
func Equal(mac1, mac2 []byte) bool {
//line /usr/local/go/src/crypto/hmac/hmac.go:175
	_go_fuzz_dep_.CoverTab[9551]++

//line /usr/local/go/src/crypto/hmac/hmac.go:179
	return subtle.ConstantTimeCompare(mac1, mac2) == 1
//line /usr/local/go/src/crypto/hmac/hmac.go:179
	// _ = "end of CoverTab[9551]"
}

//line /usr/local/go/src/crypto/hmac/hmac.go:180
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/hmac/hmac.go:180
var _ = _go_fuzz_dep_.CoverTab
