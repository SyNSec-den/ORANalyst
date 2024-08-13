// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:5
package cryptobyte

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:5
)

import (
	encoding_asn1 "encoding/asn1"
	"fmt"
	"math/big"
	"reflect"
	"time"

	"golang.org/x/crypto/cryptobyte/asn1"
)

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:21
// AddASN1Int64 appends a DER-encoded ASN.1 INTEGER.
func (b *Builder) AddASN1Int64(v int64) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:22
	_go_fuzz_dep_.CoverTab[8303]++
										b.addASN1Signed(asn1.INTEGER, v)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:23
	// _ = "end of CoverTab[8303]"
}

// AddASN1Int64WithTag appends a DER-encoded ASN.1 INTEGER with the
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:26
// given tag.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:28
func (b *Builder) AddASN1Int64WithTag(v int64, tag asn1.Tag) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:28
	_go_fuzz_dep_.CoverTab[8304]++
										b.addASN1Signed(tag, v)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:29
	// _ = "end of CoverTab[8304]"
}

// AddASN1Enum appends a DER-encoded ASN.1 ENUMERATION.
func (b *Builder) AddASN1Enum(v int64) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:33
	_go_fuzz_dep_.CoverTab[8305]++
										b.addASN1Signed(asn1.ENUM, v)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:34
	// _ = "end of CoverTab[8305]"
}

func (b *Builder) addASN1Signed(tag asn1.Tag, v int64) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:37
	_go_fuzz_dep_.CoverTab[8306]++
										b.AddASN1(tag, func(c *Builder) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:38
		_go_fuzz_dep_.CoverTab[8307]++
											length := 1
											for i := v; i >= 0x80 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:40
			_go_fuzz_dep_.CoverTab[8309]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:40
			return i < -0x80
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:40
			// _ = "end of CoverTab[8309]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:40
		}(); i >>= 8 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:40
			_go_fuzz_dep_.CoverTab[8310]++
												length++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:41
			// _ = "end of CoverTab[8310]"
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:42
		// _ = "end of CoverTab[8307]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:42
		_go_fuzz_dep_.CoverTab[8308]++

											for ; length > 0; length-- {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:44
			_go_fuzz_dep_.CoverTab[8311]++
												i := v >> uint((length-1)*8) & 0xff
												c.AddUint8(uint8(i))
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:46
			// _ = "end of CoverTab[8311]"
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:47
		// _ = "end of CoverTab[8308]"
	})
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:48
	// _ = "end of CoverTab[8306]"
}

// AddASN1Uint64 appends a DER-encoded ASN.1 INTEGER.
func (b *Builder) AddASN1Uint64(v uint64) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:52
	_go_fuzz_dep_.CoverTab[8312]++
										b.AddASN1(asn1.INTEGER, func(c *Builder) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:53
		_go_fuzz_dep_.CoverTab[8313]++
											length := 1
											for i := v; i >= 0x80; i >>= 8 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:55
			_go_fuzz_dep_.CoverTab[8315]++
												length++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:56
			// _ = "end of CoverTab[8315]"
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:57
		// _ = "end of CoverTab[8313]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:57
		_go_fuzz_dep_.CoverTab[8314]++

											for ; length > 0; length-- {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:59
			_go_fuzz_dep_.CoverTab[8316]++
												i := v >> uint((length-1)*8) & 0xff
												c.AddUint8(uint8(i))
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:61
			// _ = "end of CoverTab[8316]"
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:62
		// _ = "end of CoverTab[8314]"
	})
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:63
	// _ = "end of CoverTab[8312]"
}

// AddASN1BigInt appends a DER-encoded ASN.1 INTEGER.
func (b *Builder) AddASN1BigInt(n *big.Int) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:67
	_go_fuzz_dep_.CoverTab[8317]++
										if b.err != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:68
		_go_fuzz_dep_.CoverTab[8319]++
											return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:69
		// _ = "end of CoverTab[8319]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:70
		_go_fuzz_dep_.CoverTab[8320]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:70
		// _ = "end of CoverTab[8320]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:70
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:70
	// _ = "end of CoverTab[8317]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:70
	_go_fuzz_dep_.CoverTab[8318]++

										b.AddASN1(asn1.INTEGER, func(c *Builder) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:72
		_go_fuzz_dep_.CoverTab[8321]++
											if n.Sign() < 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:73
			_go_fuzz_dep_.CoverTab[8322]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:78
			nMinus1 := new(big.Int).Neg(n)
			nMinus1.Sub(nMinus1, bigOne)
			bytes := nMinus1.Bytes()
			for i := range bytes {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:81
				_go_fuzz_dep_.CoverTab[8325]++
													bytes[i] ^= 0xff
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:82
				// _ = "end of CoverTab[8325]"
			}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:83
			// _ = "end of CoverTab[8322]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:83
			_go_fuzz_dep_.CoverTab[8323]++
												if len(bytes) == 0 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:84
				_go_fuzz_dep_.CoverTab[8326]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:84
				return bytes[0]&0x80 == 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:84
				// _ = "end of CoverTab[8326]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:84
			}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:84
				_go_fuzz_dep_.CoverTab[8327]++
													c.add(0xff)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:85
				// _ = "end of CoverTab[8327]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:86
				_go_fuzz_dep_.CoverTab[8328]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:86
				// _ = "end of CoverTab[8328]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:86
			}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:86
			// _ = "end of CoverTab[8323]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:86
			_go_fuzz_dep_.CoverTab[8324]++
												c.add(bytes...)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:87
			// _ = "end of CoverTab[8324]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:88
			_go_fuzz_dep_.CoverTab[8329]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:88
			if n.Sign() == 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:88
				_go_fuzz_dep_.CoverTab[8330]++
													c.add(0)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:89
				// _ = "end of CoverTab[8330]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:90
				_go_fuzz_dep_.CoverTab[8331]++
													bytes := n.Bytes()
													if bytes[0]&0x80 != 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:92
					_go_fuzz_dep_.CoverTab[8333]++
														c.add(0)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:93
					// _ = "end of CoverTab[8333]"
				} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:94
					_go_fuzz_dep_.CoverTab[8334]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:94
					// _ = "end of CoverTab[8334]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:94
				}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:94
				// _ = "end of CoverTab[8331]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:94
				_go_fuzz_dep_.CoverTab[8332]++
													c.add(bytes...)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:95
				// _ = "end of CoverTab[8332]"
			}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:96
			// _ = "end of CoverTab[8329]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:96
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:96
		// _ = "end of CoverTab[8321]"
	})
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:97
	// _ = "end of CoverTab[8318]"
}

// AddASN1OctetString appends a DER-encoded ASN.1 OCTET STRING.
func (b *Builder) AddASN1OctetString(bytes []byte) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:101
	_go_fuzz_dep_.CoverTab[8335]++
										b.AddASN1(asn1.OCTET_STRING, func(c *Builder) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:102
		_go_fuzz_dep_.CoverTab[8336]++
											c.AddBytes(bytes)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:103
		// _ = "end of CoverTab[8336]"
	})
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:104
	// _ = "end of CoverTab[8335]"
}

const generalizedTimeFormatStr = "20060102150405Z0700"

// AddASN1GeneralizedTime appends a DER-encoded ASN.1 GENERALIZEDTIME.
func (b *Builder) AddASN1GeneralizedTime(t time.Time) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:110
	_go_fuzz_dep_.CoverTab[8337]++
										if t.Year() < 0 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:111
		_go_fuzz_dep_.CoverTab[8339]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:111
		return t.Year() > 9999
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:111
		// _ = "end of CoverTab[8339]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:111
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:111
		_go_fuzz_dep_.CoverTab[8340]++
											b.err = fmt.Errorf("cryptobyte: cannot represent %v as a GeneralizedTime", t)
											return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:113
		// _ = "end of CoverTab[8340]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:114
		_go_fuzz_dep_.CoverTab[8341]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:114
		// _ = "end of CoverTab[8341]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:114
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:114
	// _ = "end of CoverTab[8337]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:114
	_go_fuzz_dep_.CoverTab[8338]++
										b.AddASN1(asn1.GeneralizedTime, func(c *Builder) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:115
		_go_fuzz_dep_.CoverTab[8342]++
											c.AddBytes([]byte(t.Format(generalizedTimeFormatStr)))
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:116
		// _ = "end of CoverTab[8342]"
	})
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:117
	// _ = "end of CoverTab[8338]"
}

// AddASN1UTCTime appends a DER-encoded ASN.1 UTCTime.
func (b *Builder) AddASN1UTCTime(t time.Time) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:121
	_go_fuzz_dep_.CoverTab[8343]++
										b.AddASN1(asn1.UTCTime, func(c *Builder) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:122
		_go_fuzz_dep_.CoverTab[8344]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:125
		if t.Year() < 1950 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:125
			_go_fuzz_dep_.CoverTab[8346]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:125
			return t.Year() >= 2050
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:125
			// _ = "end of CoverTab[8346]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:125
		}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:125
			_go_fuzz_dep_.CoverTab[8347]++
												b.err = fmt.Errorf("cryptobyte: cannot represent %v as a UTCTime", t)
												return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:127
			// _ = "end of CoverTab[8347]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:128
			_go_fuzz_dep_.CoverTab[8348]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:128
			// _ = "end of CoverTab[8348]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:128
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:128
		// _ = "end of CoverTab[8344]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:128
		_go_fuzz_dep_.CoverTab[8345]++
											c.AddBytes([]byte(t.Format(defaultUTCTimeFormatStr)))
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:129
		// _ = "end of CoverTab[8345]"
	})
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:130
	// _ = "end of CoverTab[8343]"
}

// AddASN1BitString appends a DER-encoded ASN.1 BIT STRING. This does not
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:133
// support BIT STRINGs that are not a whole number of bytes.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:135
func (b *Builder) AddASN1BitString(data []byte) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:135
	_go_fuzz_dep_.CoverTab[8349]++
										b.AddASN1(asn1.BIT_STRING, func(b *Builder) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:136
		_go_fuzz_dep_.CoverTab[8350]++
											b.AddUint8(0)
											b.AddBytes(data)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:138
		// _ = "end of CoverTab[8350]"
	})
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:139
	// _ = "end of CoverTab[8349]"
}

func (b *Builder) addBase128Int(n int64) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:142
	_go_fuzz_dep_.CoverTab[8351]++
										var length int
										if n == 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:144
		_go_fuzz_dep_.CoverTab[8353]++
											length = 1
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:145
		// _ = "end of CoverTab[8353]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:146
		_go_fuzz_dep_.CoverTab[8354]++
											for i := n; i > 0; i >>= 7 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:147
			_go_fuzz_dep_.CoverTab[8355]++
												length++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:148
			// _ = "end of CoverTab[8355]"
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:149
		// _ = "end of CoverTab[8354]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:150
	// _ = "end of CoverTab[8351]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:150
	_go_fuzz_dep_.CoverTab[8352]++

										for i := length - 1; i >= 0; i-- {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:152
		_go_fuzz_dep_.CoverTab[8356]++
											o := byte(n >> uint(i*7))
											o &= 0x7f
											if i != 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:155
			_go_fuzz_dep_.CoverTab[8358]++
												o |= 0x80
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:156
			// _ = "end of CoverTab[8358]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:157
			_go_fuzz_dep_.CoverTab[8359]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:157
			// _ = "end of CoverTab[8359]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:157
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:157
		// _ = "end of CoverTab[8356]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:157
		_go_fuzz_dep_.CoverTab[8357]++

											b.add(o)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:159
		// _ = "end of CoverTab[8357]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:160
	// _ = "end of CoverTab[8352]"
}

func isValidOID(oid encoding_asn1.ObjectIdentifier) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:163
	_go_fuzz_dep_.CoverTab[8360]++
										if len(oid) < 2 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:164
		_go_fuzz_dep_.CoverTab[8364]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:165
		// _ = "end of CoverTab[8364]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:166
		_go_fuzz_dep_.CoverTab[8365]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:166
		// _ = "end of CoverTab[8365]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:166
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:166
	// _ = "end of CoverTab[8360]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:166
	_go_fuzz_dep_.CoverTab[8361]++

										if oid[0] > 2 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:168
		_go_fuzz_dep_.CoverTab[8366]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:168
		return (oid[0] <= 1 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:168
			_go_fuzz_dep_.CoverTab[8367]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:168
			return oid[1] >= 40
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:168
			// _ = "end of CoverTab[8367]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:168
		}())
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:168
		// _ = "end of CoverTab[8366]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:168
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:168
		_go_fuzz_dep_.CoverTab[8368]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:169
		// _ = "end of CoverTab[8368]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:170
		_go_fuzz_dep_.CoverTab[8369]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:170
		// _ = "end of CoverTab[8369]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:170
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:170
	// _ = "end of CoverTab[8361]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:170
	_go_fuzz_dep_.CoverTab[8362]++

										for _, v := range oid {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:172
		_go_fuzz_dep_.CoverTab[8370]++
											if v < 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:173
			_go_fuzz_dep_.CoverTab[8371]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:174
			// _ = "end of CoverTab[8371]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:175
			_go_fuzz_dep_.CoverTab[8372]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:175
			// _ = "end of CoverTab[8372]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:175
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:175
		// _ = "end of CoverTab[8370]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:176
	// _ = "end of CoverTab[8362]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:176
	_go_fuzz_dep_.CoverTab[8363]++

										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:178
	// _ = "end of CoverTab[8363]"
}

func (b *Builder) AddASN1ObjectIdentifier(oid encoding_asn1.ObjectIdentifier) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:181
	_go_fuzz_dep_.CoverTab[8373]++
										b.AddASN1(asn1.OBJECT_IDENTIFIER, func(b *Builder) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:182
		_go_fuzz_dep_.CoverTab[8374]++
											if !isValidOID(oid) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:183
			_go_fuzz_dep_.CoverTab[8376]++
												b.err = fmt.Errorf("cryptobyte: invalid OID: %v", oid)
												return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:185
			// _ = "end of CoverTab[8376]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:186
			_go_fuzz_dep_.CoverTab[8377]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:186
			// _ = "end of CoverTab[8377]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:186
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:186
		// _ = "end of CoverTab[8374]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:186
		_go_fuzz_dep_.CoverTab[8375]++

											b.addBase128Int(int64(oid[0])*40 + int64(oid[1]))
											for _, v := range oid[2:] {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:189
			_go_fuzz_dep_.CoverTab[8378]++
												b.addBase128Int(int64(v))
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:190
			// _ = "end of CoverTab[8378]"
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:191
		// _ = "end of CoverTab[8375]"
	})
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:192
	// _ = "end of CoverTab[8373]"
}

func (b *Builder) AddASN1Boolean(v bool) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:195
	_go_fuzz_dep_.CoverTab[8379]++
										b.AddASN1(asn1.BOOLEAN, func(b *Builder) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:196
		_go_fuzz_dep_.CoverTab[8380]++
											if v {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:197
			_go_fuzz_dep_.CoverTab[8381]++
												b.AddUint8(0xff)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:198
			// _ = "end of CoverTab[8381]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:199
			_go_fuzz_dep_.CoverTab[8382]++
												b.AddUint8(0)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:200
			// _ = "end of CoverTab[8382]"
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:201
		// _ = "end of CoverTab[8380]"
	})
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:202
	// _ = "end of CoverTab[8379]"
}

func (b *Builder) AddASN1NULL() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:205
	_go_fuzz_dep_.CoverTab[8383]++
										b.add(uint8(asn1.NULL), 0)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:206
	// _ = "end of CoverTab[8383]"
}

// MarshalASN1 calls encoding_asn1.Marshal on its input and appends the result if
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:209
// successful or records an error if one occurred.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:211
func (b *Builder) MarshalASN1(v interface{}) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:211
	_go_fuzz_dep_.CoverTab[8384]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:215
	if b.err != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:215
		_go_fuzz_dep_.CoverTab[8387]++
											return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:216
		// _ = "end of CoverTab[8387]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:217
		_go_fuzz_dep_.CoverTab[8388]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:217
		// _ = "end of CoverTab[8388]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:217
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:217
	// _ = "end of CoverTab[8384]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:217
	_go_fuzz_dep_.CoverTab[8385]++
										bytes, err := encoding_asn1.Marshal(v)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:219
		_go_fuzz_dep_.CoverTab[8389]++
											b.err = err
											return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:221
		// _ = "end of CoverTab[8389]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:222
		_go_fuzz_dep_.CoverTab[8390]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:222
		// _ = "end of CoverTab[8390]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:222
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:222
	// _ = "end of CoverTab[8385]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:222
	_go_fuzz_dep_.CoverTab[8386]++
										b.AddBytes(bytes)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:223
	// _ = "end of CoverTab[8386]"
}

// AddASN1 appends an ASN.1 object. The object is prefixed with the given tag.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:226
// Tags greater than 30 are not supported and result in an error (i.e.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:226
// low-tag-number form only). The child builder passed to the
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:226
// BuilderContinuation can be used to build the content of the ASN.1 object.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:230
func (b *Builder) AddASN1(tag asn1.Tag, f BuilderContinuation) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:230
	_go_fuzz_dep_.CoverTab[8391]++
										if b.err != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:231
		_go_fuzz_dep_.CoverTab[8394]++
											return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:232
		// _ = "end of CoverTab[8394]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:233
		_go_fuzz_dep_.CoverTab[8395]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:233
		// _ = "end of CoverTab[8395]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:233
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:233
	// _ = "end of CoverTab[8391]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:233
	_go_fuzz_dep_.CoverTab[8392]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:236
	if tag&0x1f == 0x1f {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:236
		_go_fuzz_dep_.CoverTab[8396]++
											b.err = fmt.Errorf("cryptobyte: high-tag number identifier octects not supported: 0x%x", tag)
											return
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:238
		// _ = "end of CoverTab[8396]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:239
		_go_fuzz_dep_.CoverTab[8397]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:239
		// _ = "end of CoverTab[8397]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:239
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:239
	// _ = "end of CoverTab[8392]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:239
	_go_fuzz_dep_.CoverTab[8393]++
										b.AddUint8(uint8(tag))
										b.addLengthPrefixed(1, true, f)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:241
	// _ = "end of CoverTab[8393]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:246
// ReadASN1Boolean decodes an ASN.1 BOOLEAN and converts it to a boolean
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:246
// representation into out and advances. It reports whether the read
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:246
// was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:249
func (s *String) ReadASN1Boolean(out *bool) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:249
	_go_fuzz_dep_.CoverTab[8398]++
										var bytes String
										if !s.ReadASN1(&bytes, asn1.BOOLEAN) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:251
		_go_fuzz_dep_.CoverTab[8401]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:251
		return len(bytes) != 1
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:251
		// _ = "end of CoverTab[8401]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:251
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:251
		_go_fuzz_dep_.CoverTab[8402]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:252
		// _ = "end of CoverTab[8402]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:253
		_go_fuzz_dep_.CoverTab[8403]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:253
		// _ = "end of CoverTab[8403]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:253
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:253
	// _ = "end of CoverTab[8398]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:253
	_go_fuzz_dep_.CoverTab[8399]++

										switch bytes[0] {
	case 0:
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:256
		_go_fuzz_dep_.CoverTab[8404]++
											*out = false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:257
		// _ = "end of CoverTab[8404]"
	case 0xff:
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:258
		_go_fuzz_dep_.CoverTab[8405]++
											*out = true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:259
		// _ = "end of CoverTab[8405]"
	default:
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:260
		_go_fuzz_dep_.CoverTab[8406]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:261
		// _ = "end of CoverTab[8406]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:262
	// _ = "end of CoverTab[8399]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:262
	_go_fuzz_dep_.CoverTab[8400]++

										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:264
	// _ = "end of CoverTab[8400]"
}

// ReadASN1Integer decodes an ASN.1 INTEGER into out and advances. If out does
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:267
// not point to an integer, to a big.Int, or to a []byte it panics. Only
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:267
// positive and zero values can be decoded into []byte, and they are returned as
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:267
// big-endian binary values that share memory with s. Positive values will have
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:267
// no leading zeroes, and zero will be returned as a single zero byte.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:267
// ReadASN1Integer reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:273
func (s *String) ReadASN1Integer(out interface{}) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:273
	_go_fuzz_dep_.CoverTab[8407]++
										switch out := out.(type) {
	case *int, *int8, *int16, *int32, *int64:
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:275
		_go_fuzz_dep_.CoverTab[8408]++
											var i int64
											if !s.readASN1Int64(&i) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:277
			_go_fuzz_dep_.CoverTab[8415]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:277
			return reflect.ValueOf(out).Elem().OverflowInt(i)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:277
			// _ = "end of CoverTab[8415]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:277
		}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:277
			_go_fuzz_dep_.CoverTab[8416]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:278
			// _ = "end of CoverTab[8416]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:279
			_go_fuzz_dep_.CoverTab[8417]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:279
			// _ = "end of CoverTab[8417]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:279
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:279
		// _ = "end of CoverTab[8408]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:279
		_go_fuzz_dep_.CoverTab[8409]++
											reflect.ValueOf(out).Elem().SetInt(i)
											return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:281
		// _ = "end of CoverTab[8409]"
	case *uint, *uint8, *uint16, *uint32, *uint64:
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:282
		_go_fuzz_dep_.CoverTab[8410]++
											var u uint64
											if !s.readASN1Uint64(&u) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:284
			_go_fuzz_dep_.CoverTab[8418]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:284
			return reflect.ValueOf(out).Elem().OverflowUint(u)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:284
			// _ = "end of CoverTab[8418]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:284
		}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:284
			_go_fuzz_dep_.CoverTab[8419]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:285
			// _ = "end of CoverTab[8419]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:286
			_go_fuzz_dep_.CoverTab[8420]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:286
			// _ = "end of CoverTab[8420]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:286
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:286
		// _ = "end of CoverTab[8410]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:286
		_go_fuzz_dep_.CoverTab[8411]++
											reflect.ValueOf(out).Elem().SetUint(u)
											return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:288
		// _ = "end of CoverTab[8411]"
	case *big.Int:
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:289
		_go_fuzz_dep_.CoverTab[8412]++
											return s.readASN1BigInt(out)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:290
		// _ = "end of CoverTab[8412]"
	case *[]byte:
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:291
		_go_fuzz_dep_.CoverTab[8413]++
											return s.readASN1Bytes(out)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:292
		// _ = "end of CoverTab[8413]"
	default:
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:293
		_go_fuzz_dep_.CoverTab[8414]++
											panic("out does not point to an integer type")
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:294
		// _ = "end of CoverTab[8414]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:295
	// _ = "end of CoverTab[8407]"
}

func checkASN1Integer(bytes []byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:298
	_go_fuzz_dep_.CoverTab[8421]++
										if len(bytes) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:299
		_go_fuzz_dep_.CoverTab[8425]++

											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:301
		// _ = "end of CoverTab[8425]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:302
		_go_fuzz_dep_.CoverTab[8426]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:302
		// _ = "end of CoverTab[8426]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:302
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:302
	// _ = "end of CoverTab[8421]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:302
	_go_fuzz_dep_.CoverTab[8422]++
										if len(bytes) == 1 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:303
		_go_fuzz_dep_.CoverTab[8427]++
											return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:304
		// _ = "end of CoverTab[8427]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:305
		_go_fuzz_dep_.CoverTab[8428]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:305
		// _ = "end of CoverTab[8428]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:305
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:305
	// _ = "end of CoverTab[8422]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:305
	_go_fuzz_dep_.CoverTab[8423]++
										if bytes[0] == 0 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:306
		_go_fuzz_dep_.CoverTab[8429]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:306
		return bytes[1]&0x80 == 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:306
		// _ = "end of CoverTab[8429]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:306
	}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:306
		_go_fuzz_dep_.CoverTab[8430]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:306
		return bytes[0] == 0xff && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:306
			_go_fuzz_dep_.CoverTab[8431]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:306
			return bytes[1]&0x80 == 0x80
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:306
			// _ = "end of CoverTab[8431]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:306
		}()
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:306
		// _ = "end of CoverTab[8430]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:306
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:306
		_go_fuzz_dep_.CoverTab[8432]++

											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:308
		// _ = "end of CoverTab[8432]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:309
		_go_fuzz_dep_.CoverTab[8433]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:309
		// _ = "end of CoverTab[8433]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:309
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:309
	// _ = "end of CoverTab[8423]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:309
	_go_fuzz_dep_.CoverTab[8424]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:310
	// _ = "end of CoverTab[8424]"
}

var bigOne = big.NewInt(1)

func (s *String) readASN1BigInt(out *big.Int) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:315
	_go_fuzz_dep_.CoverTab[8434]++
										var bytes String
										if !s.ReadASN1(&bytes, asn1.INTEGER) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:317
		_go_fuzz_dep_.CoverTab[8437]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:317
		return !checkASN1Integer(bytes)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:317
		// _ = "end of CoverTab[8437]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:317
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:317
		_go_fuzz_dep_.CoverTab[8438]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:318
		// _ = "end of CoverTab[8438]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:319
		_go_fuzz_dep_.CoverTab[8439]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:319
		// _ = "end of CoverTab[8439]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:319
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:319
	// _ = "end of CoverTab[8434]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:319
	_go_fuzz_dep_.CoverTab[8435]++
										if bytes[0]&0x80 == 0x80 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:320
		_go_fuzz_dep_.CoverTab[8440]++

											neg := make([]byte, len(bytes))
											for i, b := range bytes {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:323
			_go_fuzz_dep_.CoverTab[8442]++
												neg[i] = ^b
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:324
			// _ = "end of CoverTab[8442]"
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:325
		// _ = "end of CoverTab[8440]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:325
		_go_fuzz_dep_.CoverTab[8441]++
											out.SetBytes(neg)
											out.Add(out, bigOne)
											out.Neg(out)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:328
		// _ = "end of CoverTab[8441]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:329
		_go_fuzz_dep_.CoverTab[8443]++
											out.SetBytes(bytes)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:330
		// _ = "end of CoverTab[8443]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:331
	// _ = "end of CoverTab[8435]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:331
	_go_fuzz_dep_.CoverTab[8436]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:332
	// _ = "end of CoverTab[8436]"
}

func (s *String) readASN1Bytes(out *[]byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:335
	_go_fuzz_dep_.CoverTab[8444]++
										var bytes String
										if !s.ReadASN1(&bytes, asn1.INTEGER) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:337
		_go_fuzz_dep_.CoverTab[8448]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:337
		return !checkASN1Integer(bytes)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:337
		// _ = "end of CoverTab[8448]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:337
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:337
		_go_fuzz_dep_.CoverTab[8449]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:338
		// _ = "end of CoverTab[8449]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:339
		_go_fuzz_dep_.CoverTab[8450]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:339
		// _ = "end of CoverTab[8450]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:339
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:339
	// _ = "end of CoverTab[8444]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:339
	_go_fuzz_dep_.CoverTab[8445]++
										if bytes[0]&0x80 == 0x80 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:340
		_go_fuzz_dep_.CoverTab[8451]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:341
		// _ = "end of CoverTab[8451]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:342
		_go_fuzz_dep_.CoverTab[8452]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:342
		// _ = "end of CoverTab[8452]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:342
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:342
	// _ = "end of CoverTab[8445]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:342
	_go_fuzz_dep_.CoverTab[8446]++
										for len(bytes) > 1 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:343
		_go_fuzz_dep_.CoverTab[8453]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:343
		return bytes[0] == 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:343
		// _ = "end of CoverTab[8453]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:343
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:343
		_go_fuzz_dep_.CoverTab[8454]++
											bytes = bytes[1:]
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:344
		// _ = "end of CoverTab[8454]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:345
	// _ = "end of CoverTab[8446]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:345
	_go_fuzz_dep_.CoverTab[8447]++
										*out = bytes
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:347
	// _ = "end of CoverTab[8447]"
}

func (s *String) readASN1Int64(out *int64) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:350
	_go_fuzz_dep_.CoverTab[8455]++
										var bytes String
										if !s.ReadASN1(&bytes, asn1.INTEGER) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:352
		_go_fuzz_dep_.CoverTab[8457]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:352
		return !checkASN1Integer(bytes)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:352
		// _ = "end of CoverTab[8457]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:352
	}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:352
		_go_fuzz_dep_.CoverTab[8458]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:352
		return !asn1Signed(out, bytes)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:352
		// _ = "end of CoverTab[8458]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:352
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:352
		_go_fuzz_dep_.CoverTab[8459]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:353
		// _ = "end of CoverTab[8459]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:354
		_go_fuzz_dep_.CoverTab[8460]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:354
		// _ = "end of CoverTab[8460]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:354
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:354
	// _ = "end of CoverTab[8455]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:354
	_go_fuzz_dep_.CoverTab[8456]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:355
	// _ = "end of CoverTab[8456]"
}

func asn1Signed(out *int64, n []byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:358
	_go_fuzz_dep_.CoverTab[8461]++
										length := len(n)
										if length > 8 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:360
		_go_fuzz_dep_.CoverTab[8464]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:361
		// _ = "end of CoverTab[8464]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:362
		_go_fuzz_dep_.CoverTab[8465]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:362
		// _ = "end of CoverTab[8465]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:362
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:362
	// _ = "end of CoverTab[8461]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:362
	_go_fuzz_dep_.CoverTab[8462]++
										for i := 0; i < length; i++ {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:363
		_go_fuzz_dep_.CoverTab[8466]++
											*out <<= 8
											*out |= int64(n[i])
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:365
		// _ = "end of CoverTab[8466]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:366
	// _ = "end of CoverTab[8462]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:366
	_go_fuzz_dep_.CoverTab[8463]++

										*out <<= 64 - uint8(length)*8
										*out >>= 64 - uint8(length)*8
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:370
	// _ = "end of CoverTab[8463]"
}

func (s *String) readASN1Uint64(out *uint64) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:373
	_go_fuzz_dep_.CoverTab[8467]++
										var bytes String
										if !s.ReadASN1(&bytes, asn1.INTEGER) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:375
		_go_fuzz_dep_.CoverTab[8469]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:375
		return !checkASN1Integer(bytes)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:375
		// _ = "end of CoverTab[8469]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:375
	}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:375
		_go_fuzz_dep_.CoverTab[8470]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:375
		return !asn1Unsigned(out, bytes)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:375
		// _ = "end of CoverTab[8470]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:375
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:375
		_go_fuzz_dep_.CoverTab[8471]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:376
		// _ = "end of CoverTab[8471]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:377
		_go_fuzz_dep_.CoverTab[8472]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:377
		// _ = "end of CoverTab[8472]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:377
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:377
	// _ = "end of CoverTab[8467]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:377
	_go_fuzz_dep_.CoverTab[8468]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:378
	// _ = "end of CoverTab[8468]"
}

func asn1Unsigned(out *uint64, n []byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:381
	_go_fuzz_dep_.CoverTab[8473]++
										length := len(n)
										if length > 9 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:383
		_go_fuzz_dep_.CoverTab[8477]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:383
		return length == 9 && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:383
			_go_fuzz_dep_.CoverTab[8478]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:383
			return n[0] != 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:383
			// _ = "end of CoverTab[8478]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:383
		}()
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:383
		// _ = "end of CoverTab[8477]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:383
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:383
		_go_fuzz_dep_.CoverTab[8479]++

											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:385
		// _ = "end of CoverTab[8479]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:386
		_go_fuzz_dep_.CoverTab[8480]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:386
		// _ = "end of CoverTab[8480]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:386
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:386
	// _ = "end of CoverTab[8473]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:386
	_go_fuzz_dep_.CoverTab[8474]++
										if n[0]&0x80 != 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:387
		_go_fuzz_dep_.CoverTab[8481]++

											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:389
		// _ = "end of CoverTab[8481]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:390
		_go_fuzz_dep_.CoverTab[8482]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:390
		// _ = "end of CoverTab[8482]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:390
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:390
	// _ = "end of CoverTab[8474]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:390
	_go_fuzz_dep_.CoverTab[8475]++
										for i := 0; i < length; i++ {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:391
		_go_fuzz_dep_.CoverTab[8483]++
											*out <<= 8
											*out |= uint64(n[i])
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:393
		// _ = "end of CoverTab[8483]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:394
	// _ = "end of CoverTab[8475]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:394
	_go_fuzz_dep_.CoverTab[8476]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:395
	// _ = "end of CoverTab[8476]"
}

// ReadASN1Int64WithTag decodes an ASN.1 INTEGER with the given tag into out
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:398
// and advances. It reports whether the read was successful and resulted in a
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:398
// value that can be represented in an int64.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:401
func (s *String) ReadASN1Int64WithTag(out *int64, tag asn1.Tag) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:401
	_go_fuzz_dep_.CoverTab[8484]++
										var bytes String
										return s.ReadASN1(&bytes, tag) && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:403
		_go_fuzz_dep_.CoverTab[8485]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:403
		return checkASN1Integer(bytes)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:403
		// _ = "end of CoverTab[8485]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:403
	}() && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:403
		_go_fuzz_dep_.CoverTab[8486]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:403
		return asn1Signed(out, bytes)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:403
		// _ = "end of CoverTab[8486]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:403
	}()
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:403
	// _ = "end of CoverTab[8484]"
}

// ReadASN1Enum decodes an ASN.1 ENUMERATION into out and advances. It reports
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:406
// whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:408
func (s *String) ReadASN1Enum(out *int) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:408
	_go_fuzz_dep_.CoverTab[8487]++
										var bytes String
										var i int64
										if !s.ReadASN1(&bytes, asn1.ENUM) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:411
		_go_fuzz_dep_.CoverTab[8490]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:411
		return !checkASN1Integer(bytes)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:411
		// _ = "end of CoverTab[8490]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:411
	}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:411
		_go_fuzz_dep_.CoverTab[8491]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:411
		return !asn1Signed(&i, bytes)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:411
		// _ = "end of CoverTab[8491]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:411
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:411
		_go_fuzz_dep_.CoverTab[8492]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:412
		// _ = "end of CoverTab[8492]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:413
		_go_fuzz_dep_.CoverTab[8493]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:413
		// _ = "end of CoverTab[8493]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:413
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:413
	// _ = "end of CoverTab[8487]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:413
	_go_fuzz_dep_.CoverTab[8488]++
										if int64(int(i)) != i {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:414
		_go_fuzz_dep_.CoverTab[8494]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:415
		// _ = "end of CoverTab[8494]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:416
		_go_fuzz_dep_.CoverTab[8495]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:416
		// _ = "end of CoverTab[8495]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:416
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:416
	// _ = "end of CoverTab[8488]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:416
	_go_fuzz_dep_.CoverTab[8489]++
										*out = int(i)
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:418
	// _ = "end of CoverTab[8489]"
}

func (s *String) readBase128Int(out *int) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:421
	_go_fuzz_dep_.CoverTab[8496]++
										ret := 0
										for i := 0; len(*s) > 0; i++ {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:423
		_go_fuzz_dep_.CoverTab[8498]++
											if i == 5 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:424
			_go_fuzz_dep_.CoverTab[8501]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:425
			// _ = "end of CoverTab[8501]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:426
			_go_fuzz_dep_.CoverTab[8502]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:426
			// _ = "end of CoverTab[8502]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:426
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:426
		// _ = "end of CoverTab[8498]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:426
		_go_fuzz_dep_.CoverTab[8499]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:429
		if ret >= 1<<(31-7) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:429
			_go_fuzz_dep_.CoverTab[8503]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:430
			// _ = "end of CoverTab[8503]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:431
			_go_fuzz_dep_.CoverTab[8504]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:431
			// _ = "end of CoverTab[8504]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:431
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:431
		// _ = "end of CoverTab[8499]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:431
		_go_fuzz_dep_.CoverTab[8500]++
											ret <<= 7
											b := s.read(1)[0]
											ret |= int(b & 0x7f)
											if b&0x80 == 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:435
			_go_fuzz_dep_.CoverTab[8505]++
												*out = ret
												return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:437
			// _ = "end of CoverTab[8505]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:438
			_go_fuzz_dep_.CoverTab[8506]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:438
			// _ = "end of CoverTab[8506]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:438
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:438
		// _ = "end of CoverTab[8500]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:439
	// _ = "end of CoverTab[8496]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:439
	_go_fuzz_dep_.CoverTab[8497]++
										return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:440
	// _ = "end of CoverTab[8497]"
}

// ReadASN1ObjectIdentifier decodes an ASN.1 OBJECT IDENTIFIER into out and
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:443
// advances. It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:445
func (s *String) ReadASN1ObjectIdentifier(out *encoding_asn1.ObjectIdentifier) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:445
	_go_fuzz_dep_.CoverTab[8507]++
										var bytes String
										if !s.ReadASN1(&bytes, asn1.OBJECT_IDENTIFIER) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:447
		_go_fuzz_dep_.CoverTab[8512]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:447
		return len(bytes) == 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:447
		// _ = "end of CoverTab[8512]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:447
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:447
		_go_fuzz_dep_.CoverTab[8513]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:448
		// _ = "end of CoverTab[8513]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:449
		_go_fuzz_dep_.CoverTab[8514]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:449
		// _ = "end of CoverTab[8514]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:449
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:449
	// _ = "end of CoverTab[8507]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:449
	_go_fuzz_dep_.CoverTab[8508]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:453
	components := make([]int, len(bytes)+1)

	// The first varint is 40*value1 + value2:
	// According to this packing, value1 can take the values 0, 1 and 2 only.
	// When value1 = 0 or value1 = 1, then value2 is <= 39. When value1 = 2,
	// then there are no restrictions on value2.
	var v int
	if !bytes.readBase128Int(&v) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:460
		_go_fuzz_dep_.CoverTab[8515]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:461
		// _ = "end of CoverTab[8515]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:462
		_go_fuzz_dep_.CoverTab[8516]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:462
		// _ = "end of CoverTab[8516]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:462
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:462
	// _ = "end of CoverTab[8508]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:462
	_go_fuzz_dep_.CoverTab[8509]++
										if v < 80 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:463
		_go_fuzz_dep_.CoverTab[8517]++
											components[0] = v / 40
											components[1] = v % 40
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:465
		// _ = "end of CoverTab[8517]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:466
		_go_fuzz_dep_.CoverTab[8518]++
											components[0] = 2
											components[1] = v - 80
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:468
		// _ = "end of CoverTab[8518]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:469
	// _ = "end of CoverTab[8509]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:469
	_go_fuzz_dep_.CoverTab[8510]++

										i := 2
										for ; len(bytes) > 0; i++ {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:472
		_go_fuzz_dep_.CoverTab[8519]++
											if !bytes.readBase128Int(&v) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:473
			_go_fuzz_dep_.CoverTab[8521]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:474
			// _ = "end of CoverTab[8521]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:475
			_go_fuzz_dep_.CoverTab[8522]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:475
			// _ = "end of CoverTab[8522]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:475
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:475
		// _ = "end of CoverTab[8519]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:475
		_go_fuzz_dep_.CoverTab[8520]++
											components[i] = v
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:476
		// _ = "end of CoverTab[8520]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:477
	// _ = "end of CoverTab[8510]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:477
	_go_fuzz_dep_.CoverTab[8511]++
										*out = components[:i]
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:479
	// _ = "end of CoverTab[8511]"
}

// ReadASN1GeneralizedTime decodes an ASN.1 GENERALIZEDTIME into out and
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:482
// advances. It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:484
func (s *String) ReadASN1GeneralizedTime(out *time.Time) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:484
	_go_fuzz_dep_.CoverTab[8523]++
										var bytes String
										if !s.ReadASN1(&bytes, asn1.GeneralizedTime) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:486
		_go_fuzz_dep_.CoverTab[8527]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:487
		// _ = "end of CoverTab[8527]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:488
		_go_fuzz_dep_.CoverTab[8528]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:488
		// _ = "end of CoverTab[8528]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:488
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:488
	// _ = "end of CoverTab[8523]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:488
	_go_fuzz_dep_.CoverTab[8524]++
										t := string(bytes)
										res, err := time.Parse(generalizedTimeFormatStr, t)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:491
		_go_fuzz_dep_.CoverTab[8529]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:492
		// _ = "end of CoverTab[8529]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:493
		_go_fuzz_dep_.CoverTab[8530]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:493
		// _ = "end of CoverTab[8530]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:493
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:493
	// _ = "end of CoverTab[8524]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:493
	_go_fuzz_dep_.CoverTab[8525]++
										if serialized := res.Format(generalizedTimeFormatStr); serialized != t {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:494
		_go_fuzz_dep_.CoverTab[8531]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:495
		// _ = "end of CoverTab[8531]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:496
		_go_fuzz_dep_.CoverTab[8532]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:496
		// _ = "end of CoverTab[8532]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:496
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:496
	// _ = "end of CoverTab[8525]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:496
	_go_fuzz_dep_.CoverTab[8526]++
										*out = res
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:498
	// _ = "end of CoverTab[8526]"
}

const defaultUTCTimeFormatStr = "060102150405Z0700"

// ReadASN1UTCTime decodes an ASN.1 UTCTime into out and advances.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:503
// It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:505
func (s *String) ReadASN1UTCTime(out *time.Time) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:505
	_go_fuzz_dep_.CoverTab[8533]++
										var bytes String
										if !s.ReadASN1(&bytes, asn1.UTCTime) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:507
		_go_fuzz_dep_.CoverTab[8539]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:508
		// _ = "end of CoverTab[8539]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:509
		_go_fuzz_dep_.CoverTab[8540]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:509
		// _ = "end of CoverTab[8540]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:509
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:509
	// _ = "end of CoverTab[8533]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:509
	_go_fuzz_dep_.CoverTab[8534]++
										t := string(bytes)

										formatStr := defaultUTCTimeFormatStr
										var err error
										res, err := time.Parse(formatStr, t)
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:515
		_go_fuzz_dep_.CoverTab[8541]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:519
		formatStr = "0601021504Z0700"
											res, err = time.Parse(formatStr, t)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:520
		// _ = "end of CoverTab[8541]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:521
		_go_fuzz_dep_.CoverTab[8542]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:521
		// _ = "end of CoverTab[8542]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:521
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:521
	// _ = "end of CoverTab[8534]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:521
	_go_fuzz_dep_.CoverTab[8535]++
										if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:522
		_go_fuzz_dep_.CoverTab[8543]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:523
		// _ = "end of CoverTab[8543]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:524
		_go_fuzz_dep_.CoverTab[8544]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:524
		// _ = "end of CoverTab[8544]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:524
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:524
	// _ = "end of CoverTab[8535]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:524
	_go_fuzz_dep_.CoverTab[8536]++

										if serialized := res.Format(formatStr); serialized != t {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:526
		_go_fuzz_dep_.CoverTab[8545]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:527
		// _ = "end of CoverTab[8545]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:528
		_go_fuzz_dep_.CoverTab[8546]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:528
		// _ = "end of CoverTab[8546]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:528
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:528
	// _ = "end of CoverTab[8536]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:528
	_go_fuzz_dep_.CoverTab[8537]++

										if res.Year() >= 2050 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:530
		_go_fuzz_dep_.CoverTab[8547]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:534
		res = res.AddDate(-100, 0, 0)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:534
		// _ = "end of CoverTab[8547]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:535
		_go_fuzz_dep_.CoverTab[8548]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:535
		// _ = "end of CoverTab[8548]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:535
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:535
	// _ = "end of CoverTab[8537]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:535
	_go_fuzz_dep_.CoverTab[8538]++
										*out = res
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:537
	// _ = "end of CoverTab[8538]"
}

// ReadASN1BitString decodes an ASN.1 BIT STRING into out and advances.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:540
// It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:542
func (s *String) ReadASN1BitString(out *encoding_asn1.BitString) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:542
	_go_fuzz_dep_.CoverTab[8549]++
										var bytes String
										if !s.ReadASN1(&bytes, asn1.BIT_STRING) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:544
		_go_fuzz_dep_.CoverTab[8552]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:544
		return len(bytes) == 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:544
		// _ = "end of CoverTab[8552]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:544
	}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:544
		_go_fuzz_dep_.CoverTab[8553]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:544
		return len(bytes)*8/8 != len(bytes)
											// _ = "end of CoverTab[8553]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:545
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:545
		_go_fuzz_dep_.CoverTab[8554]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:546
		// _ = "end of CoverTab[8554]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:547
		_go_fuzz_dep_.CoverTab[8555]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:547
		// _ = "end of CoverTab[8555]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:547
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:547
	// _ = "end of CoverTab[8549]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:547
	_go_fuzz_dep_.CoverTab[8550]++

										paddingBits := bytes[0]
										bytes = bytes[1:]
										if paddingBits > 7 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:551
		_go_fuzz_dep_.CoverTab[8556]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:551
		return len(bytes) == 0 && func() bool {
												_go_fuzz_dep_.CoverTab[8557]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:552
			return paddingBits != 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:552
			// _ = "end of CoverTab[8557]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:552
		}()
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:552
		// _ = "end of CoverTab[8556]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:552
	}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:552
		_go_fuzz_dep_.CoverTab[8558]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:552
		return len(bytes) > 0 && func() bool {
												_go_fuzz_dep_.CoverTab[8559]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:553
			return bytes[len(bytes)-1]&(1<<paddingBits-1) != 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:553
			// _ = "end of CoverTab[8559]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:553
		}()
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:553
		// _ = "end of CoverTab[8558]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:553
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:553
		_go_fuzz_dep_.CoverTab[8560]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:554
		// _ = "end of CoverTab[8560]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:555
		_go_fuzz_dep_.CoverTab[8561]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:555
		// _ = "end of CoverTab[8561]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:555
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:555
	// _ = "end of CoverTab[8550]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:555
	_go_fuzz_dep_.CoverTab[8551]++

										out.BitLength = len(bytes)*8 - int(paddingBits)
										out.Bytes = bytes
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:559
	// _ = "end of CoverTab[8551]"
}

// ReadASN1BitString decodes an ASN.1 BIT STRING into out and advances. It is
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:562
// an error if the BIT STRING is not a whole number of bytes. It reports
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:562
// whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:565
func (s *String) ReadASN1BitStringAsBytes(out *[]byte) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:565
	_go_fuzz_dep_.CoverTab[8562]++
										var bytes String
										if !s.ReadASN1(&bytes, asn1.BIT_STRING) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:567
		_go_fuzz_dep_.CoverTab[8565]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:567
		return len(bytes) == 0
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:567
		// _ = "end of CoverTab[8565]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:567
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:567
		_go_fuzz_dep_.CoverTab[8566]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:568
		// _ = "end of CoverTab[8566]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:569
		_go_fuzz_dep_.CoverTab[8567]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:569
		// _ = "end of CoverTab[8567]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:569
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:569
	// _ = "end of CoverTab[8562]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:569
	_go_fuzz_dep_.CoverTab[8563]++

										paddingBits := bytes[0]
										if paddingBits != 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:572
		_go_fuzz_dep_.CoverTab[8568]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:573
		// _ = "end of CoverTab[8568]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:574
		_go_fuzz_dep_.CoverTab[8569]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:574
		// _ = "end of CoverTab[8569]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:574
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:574
	// _ = "end of CoverTab[8563]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:574
	_go_fuzz_dep_.CoverTab[8564]++
										*out = bytes[1:]
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:576
	// _ = "end of CoverTab[8564]"
}

// ReadASN1Bytes reads the contents of a DER-encoded ASN.1 element (not including
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:579
// tag and length bytes) into out, and advances. The element must match the
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:579
// given tag. It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:582
func (s *String) ReadASN1Bytes(out *[]byte, tag asn1.Tag) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:582
	_go_fuzz_dep_.CoverTab[8570]++
										return s.ReadASN1((*String)(out), tag)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:583
	// _ = "end of CoverTab[8570]"
}

// ReadASN1 reads the contents of a DER-encoded ASN.1 element (not including
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:586
// tag and length bytes) into out, and advances. The element must match the
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:586
// given tag. It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:586
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:586
// Tags greater than 30 are not supported (i.e. low-tag-number format only).
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:591
func (s *String) ReadASN1(out *String, tag asn1.Tag) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:591
	_go_fuzz_dep_.CoverTab[8571]++
										var t asn1.Tag
										if !s.ReadAnyASN1(out, &t) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:593
		_go_fuzz_dep_.CoverTab[8573]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:593
		return t != tag
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:593
		// _ = "end of CoverTab[8573]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:593
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:593
		_go_fuzz_dep_.CoverTab[8574]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:594
		// _ = "end of CoverTab[8574]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:595
		_go_fuzz_dep_.CoverTab[8575]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:595
		// _ = "end of CoverTab[8575]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:595
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:595
	// _ = "end of CoverTab[8571]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:595
	_go_fuzz_dep_.CoverTab[8572]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:596
	// _ = "end of CoverTab[8572]"
}

// ReadASN1Element reads the contents of a DER-encoded ASN.1 element (including
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:599
// tag and length bytes) into out, and advances. The element must match the
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:599
// given tag. It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:599
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:599
// Tags greater than 30 are not supported (i.e. low-tag-number format only).
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:604
func (s *String) ReadASN1Element(out *String, tag asn1.Tag) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:604
	_go_fuzz_dep_.CoverTab[8576]++
										var t asn1.Tag
										if !s.ReadAnyASN1Element(out, &t) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:606
		_go_fuzz_dep_.CoverTab[8578]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:606
		return t != tag
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:606
		// _ = "end of CoverTab[8578]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:606
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:606
		_go_fuzz_dep_.CoverTab[8579]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:607
		// _ = "end of CoverTab[8579]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:608
		_go_fuzz_dep_.CoverTab[8580]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:608
		// _ = "end of CoverTab[8580]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:608
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:608
	// _ = "end of CoverTab[8576]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:608
	_go_fuzz_dep_.CoverTab[8577]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:609
	// _ = "end of CoverTab[8577]"
}

// ReadAnyASN1 reads the contents of a DER-encoded ASN.1 element (not including
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:612
// tag and length bytes) into out, sets outTag to its tag, and advances.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:612
// It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:612
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:612
// Tags greater than 30 are not supported (i.e. low-tag-number format only).
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:617
func (s *String) ReadAnyASN1(out *String, outTag *asn1.Tag) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:617
	_go_fuzz_dep_.CoverTab[8581]++
										return s.readASN1(out, outTag, true)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:618
	// _ = "end of CoverTab[8581]"
}

// ReadAnyASN1Element reads the contents of a DER-encoded ASN.1 element
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:621
// (including tag and length bytes) into out, sets outTag to is tag, and
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:621
// advances. It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:621
//
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:621
// Tags greater than 30 are not supported (i.e. low-tag-number format only).
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:626
func (s *String) ReadAnyASN1Element(out *String, outTag *asn1.Tag) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:626
	_go_fuzz_dep_.CoverTab[8582]++
										return s.readASN1(out, outTag, false)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:627
	// _ = "end of CoverTab[8582]"
}

// PeekASN1Tag reports whether the next ASN.1 value on the string starts with
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:630
// the given tag.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:632
func (s String) PeekASN1Tag(tag asn1.Tag) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:632
	_go_fuzz_dep_.CoverTab[8583]++
										if len(s) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:633
		_go_fuzz_dep_.CoverTab[8585]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:634
		// _ = "end of CoverTab[8585]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:635
		_go_fuzz_dep_.CoverTab[8586]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:635
		// _ = "end of CoverTab[8586]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:635
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:635
	// _ = "end of CoverTab[8583]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:635
	_go_fuzz_dep_.CoverTab[8584]++
										return asn1.Tag(s[0]) == tag
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:636
	// _ = "end of CoverTab[8584]"
}

// SkipASN1 reads and discards an ASN.1 element with the given tag. It
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:639
// reports whether the operation was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:641
func (s *String) SkipASN1(tag asn1.Tag) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:641
	_go_fuzz_dep_.CoverTab[8587]++
										var unused String
										return s.ReadASN1(&unused, tag)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:643
	// _ = "end of CoverTab[8587]"
}

// ReadOptionalASN1 attempts to read the contents of a DER-encoded ASN.1
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:646
// element (not including tag and length bytes) tagged with the given tag into
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:646
// out. It stores whether an element with the tag was found in outPresent,
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:646
// unless outPresent is nil. It reports whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:650
func (s *String) ReadOptionalASN1(out *String, outPresent *bool, tag asn1.Tag) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:650
	_go_fuzz_dep_.CoverTab[8588]++
										present := s.PeekASN1Tag(tag)
										if outPresent != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:652
		_go_fuzz_dep_.CoverTab[8591]++
											*outPresent = present
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:653
		// _ = "end of CoverTab[8591]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:654
		_go_fuzz_dep_.CoverTab[8592]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:654
		// _ = "end of CoverTab[8592]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:654
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:654
	// _ = "end of CoverTab[8588]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:654
	_go_fuzz_dep_.CoverTab[8589]++
										if present && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:655
		_go_fuzz_dep_.CoverTab[8593]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:655
		return !s.ReadASN1(out, tag)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:655
		// _ = "end of CoverTab[8593]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:655
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:655
		_go_fuzz_dep_.CoverTab[8594]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:656
		// _ = "end of CoverTab[8594]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:657
		_go_fuzz_dep_.CoverTab[8595]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:657
		// _ = "end of CoverTab[8595]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:657
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:657
	// _ = "end of CoverTab[8589]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:657
	_go_fuzz_dep_.CoverTab[8590]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:658
	// _ = "end of CoverTab[8590]"
}

// SkipOptionalASN1 advances s over an ASN.1 element with the given tag, or
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:661
// else leaves s unchanged. It reports whether the operation was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:663
func (s *String) SkipOptionalASN1(tag asn1.Tag) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:663
	_go_fuzz_dep_.CoverTab[8596]++
										if !s.PeekASN1Tag(tag) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:664
		_go_fuzz_dep_.CoverTab[8598]++
											return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:665
		// _ = "end of CoverTab[8598]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:666
		_go_fuzz_dep_.CoverTab[8599]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:666
		// _ = "end of CoverTab[8599]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:666
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:666
	// _ = "end of CoverTab[8596]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:666
	_go_fuzz_dep_.CoverTab[8597]++
										var unused String
										return s.ReadASN1(&unused, tag)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:668
	// _ = "end of CoverTab[8597]"
}

// ReadOptionalASN1Integer attempts to read an optional ASN.1 INTEGER explicitly
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:671
// tagged with tag into out and advances. If no element with a matching tag is
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:671
// present, it writes defaultValue into out instead. Otherwise, it behaves like
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:671
// ReadASN1Integer.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:675
func (s *String) ReadOptionalASN1Integer(out interface{}, tag asn1.Tag, defaultValue interface{}) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:675
	_go_fuzz_dep_.CoverTab[8600]++
										var present bool
										var i String
										if !s.ReadOptionalASN1(&i, &present, tag) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:678
		_go_fuzz_dep_.CoverTab[8604]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:679
		// _ = "end of CoverTab[8604]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:680
		_go_fuzz_dep_.CoverTab[8605]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:680
		// _ = "end of CoverTab[8605]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:680
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:680
	// _ = "end of CoverTab[8600]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:680
	_go_fuzz_dep_.CoverTab[8601]++
										if !present {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:681
		_go_fuzz_dep_.CoverTab[8606]++
											switch out.(type) {
		case *int, *int8, *int16, *int32, *int64,
			*uint, *uint8, *uint16, *uint32, *uint64, *[]byte:
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:684
			_go_fuzz_dep_.CoverTab[8608]++
												reflect.ValueOf(out).Elem().Set(reflect.ValueOf(defaultValue))
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:685
			// _ = "end of CoverTab[8608]"
		case *big.Int:
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:686
			_go_fuzz_dep_.CoverTab[8609]++
												if defaultValue, ok := defaultValue.(*big.Int); ok {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:687
				_go_fuzz_dep_.CoverTab[8611]++
													out.(*big.Int).Set(defaultValue)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:688
				// _ = "end of CoverTab[8611]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:689
				_go_fuzz_dep_.CoverTab[8612]++
													panic("out points to big.Int, but defaultValue does not")
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:690
				// _ = "end of CoverTab[8612]"
			}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:691
			// _ = "end of CoverTab[8609]"
		default:
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:692
			_go_fuzz_dep_.CoverTab[8610]++
												panic("invalid integer type")
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:693
			// _ = "end of CoverTab[8610]"
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:694
		// _ = "end of CoverTab[8606]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:694
		_go_fuzz_dep_.CoverTab[8607]++
											return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:695
		// _ = "end of CoverTab[8607]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:696
		_go_fuzz_dep_.CoverTab[8613]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:696
		// _ = "end of CoverTab[8613]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:696
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:696
	// _ = "end of CoverTab[8601]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:696
	_go_fuzz_dep_.CoverTab[8602]++
										if !i.ReadASN1Integer(out) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:697
		_go_fuzz_dep_.CoverTab[8614]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:697
		return !i.Empty()
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:697
		// _ = "end of CoverTab[8614]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:697
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:697
		_go_fuzz_dep_.CoverTab[8615]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:698
		// _ = "end of CoverTab[8615]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:699
		_go_fuzz_dep_.CoverTab[8616]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:699
		// _ = "end of CoverTab[8616]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:699
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:699
	// _ = "end of CoverTab[8602]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:699
	_go_fuzz_dep_.CoverTab[8603]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:700
	// _ = "end of CoverTab[8603]"
}

// ReadOptionalASN1OctetString attempts to read an optional ASN.1 OCTET STRING
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:703
// explicitly tagged with tag into out and advances. If no element with a
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:703
// matching tag is present, it sets "out" to nil instead. It reports
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:703
// whether the read was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:707
func (s *String) ReadOptionalASN1OctetString(out *[]byte, outPresent *bool, tag asn1.Tag) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:707
	_go_fuzz_dep_.CoverTab[8617]++
										var present bool
										var child String
										if !s.ReadOptionalASN1(&child, &present, tag) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:710
		_go_fuzz_dep_.CoverTab[8621]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:711
		// _ = "end of CoverTab[8621]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:712
		_go_fuzz_dep_.CoverTab[8622]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:712
		// _ = "end of CoverTab[8622]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:712
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:712
	// _ = "end of CoverTab[8617]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:712
	_go_fuzz_dep_.CoverTab[8618]++
										if outPresent != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:713
		_go_fuzz_dep_.CoverTab[8623]++
											*outPresent = present
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:714
		// _ = "end of CoverTab[8623]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:715
		_go_fuzz_dep_.CoverTab[8624]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:715
		// _ = "end of CoverTab[8624]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:715
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:715
	// _ = "end of CoverTab[8618]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:715
	_go_fuzz_dep_.CoverTab[8619]++
										if present {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:716
		_go_fuzz_dep_.CoverTab[8625]++
											var oct String
											if !child.ReadASN1(&oct, asn1.OCTET_STRING) || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:718
			_go_fuzz_dep_.CoverTab[8627]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:718
			return !child.Empty()
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:718
			// _ = "end of CoverTab[8627]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:718
		}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:718
			_go_fuzz_dep_.CoverTab[8628]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:719
			// _ = "end of CoverTab[8628]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:720
			_go_fuzz_dep_.CoverTab[8629]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:720
			// _ = "end of CoverTab[8629]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:720
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:720
		// _ = "end of CoverTab[8625]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:720
		_go_fuzz_dep_.CoverTab[8626]++
											*out = oct
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:721
		// _ = "end of CoverTab[8626]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:722
		_go_fuzz_dep_.CoverTab[8630]++
											*out = nil
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:723
		// _ = "end of CoverTab[8630]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:724
	// _ = "end of CoverTab[8619]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:724
	_go_fuzz_dep_.CoverTab[8620]++
										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:725
	// _ = "end of CoverTab[8620]"
}

// ReadOptionalASN1Boolean sets *out to the value of the next ASN.1 BOOLEAN or,
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:728
// if the next bytes are not an ASN.1 BOOLEAN, to the value of defaultValue.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:728
// It reports whether the operation was successful.
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:731
func (s *String) ReadOptionalASN1Boolean(out *bool, defaultValue bool) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:731
	_go_fuzz_dep_.CoverTab[8631]++
										var present bool
										var child String
										if !s.ReadOptionalASN1(&child, &present, asn1.BOOLEAN) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:734
		_go_fuzz_dep_.CoverTab[8634]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:735
		// _ = "end of CoverTab[8634]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:736
		_go_fuzz_dep_.CoverTab[8635]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:736
		// _ = "end of CoverTab[8635]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:736
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:736
	// _ = "end of CoverTab[8631]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:736
	_go_fuzz_dep_.CoverTab[8632]++

										if !present {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:738
		_go_fuzz_dep_.CoverTab[8636]++
											*out = defaultValue
											return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:740
		// _ = "end of CoverTab[8636]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:741
		_go_fuzz_dep_.CoverTab[8637]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:741
		// _ = "end of CoverTab[8637]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:741
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:741
	// _ = "end of CoverTab[8632]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:741
	_go_fuzz_dep_.CoverTab[8633]++

										return s.ReadASN1Boolean(out)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:743
	// _ = "end of CoverTab[8633]"
}

func (s *String) readASN1(out *String, outTag *asn1.Tag, skipHeader bool) bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:746
	_go_fuzz_dep_.CoverTab[8638]++
										if len(*s) < 2 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:747
		_go_fuzz_dep_.CoverTab[8645]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:748
		// _ = "end of CoverTab[8645]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:749
		_go_fuzz_dep_.CoverTab[8646]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:749
		// _ = "end of CoverTab[8646]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:749
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:749
	// _ = "end of CoverTab[8638]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:749
	_go_fuzz_dep_.CoverTab[8639]++
										tag, lenByte := (*s)[0], (*s)[1]

										if tag&0x1f == 0x1f {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:752
		_go_fuzz_dep_.CoverTab[8647]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:758
		return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:758
		// _ = "end of CoverTab[8647]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:759
		_go_fuzz_dep_.CoverTab[8648]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:759
		// _ = "end of CoverTab[8648]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:759
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:759
	// _ = "end of CoverTab[8639]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:759
	_go_fuzz_dep_.CoverTab[8640]++

										if outTag != nil {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:761
		_go_fuzz_dep_.CoverTab[8649]++
											*outTag = asn1.Tag(tag)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:762
		// _ = "end of CoverTab[8649]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:763
		_go_fuzz_dep_.CoverTab[8650]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:763
		// _ = "end of CoverTab[8650]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:763
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:763
	// _ = "end of CoverTab[8640]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:763
	_go_fuzz_dep_.CoverTab[8641]++

	// ITU-T X.690 section 8.1.3
	//
	// Bit 8 of the first length byte indicates whether the length is short- or
	// long-form.
	var length, headerLen uint32	// length includes headerLen
	if lenByte&0x80 == 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:770
		_go_fuzz_dep_.CoverTab[8651]++

											length = uint32(lenByte) + 2
											headerLen = 2
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:773
		// _ = "end of CoverTab[8651]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:774
		_go_fuzz_dep_.CoverTab[8652]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:777
		lenLen := lenByte & 0x7f
		var len32 uint32

		if lenLen == 0 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:780
			_go_fuzz_dep_.CoverTab[8658]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:780
			return lenLen > 4
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:780
			// _ = "end of CoverTab[8658]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:780
		}() || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:780
			_go_fuzz_dep_.CoverTab[8659]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:780
			return len(*s) < int(2+lenLen)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:780
			// _ = "end of CoverTab[8659]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:780
		}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:780
			_go_fuzz_dep_.CoverTab[8660]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:781
			// _ = "end of CoverTab[8660]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:782
			_go_fuzz_dep_.CoverTab[8661]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:782
			// _ = "end of CoverTab[8661]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:782
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:782
		// _ = "end of CoverTab[8652]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:782
		_go_fuzz_dep_.CoverTab[8653]++

											lenBytes := String((*s)[2 : 2+lenLen])
											if !lenBytes.readUnsigned(&len32, int(lenLen)) {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:785
			_go_fuzz_dep_.CoverTab[8662]++
												return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:786
			// _ = "end of CoverTab[8662]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:787
			_go_fuzz_dep_.CoverTab[8663]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:787
			// _ = "end of CoverTab[8663]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:787
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:787
		// _ = "end of CoverTab[8653]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:787
		_go_fuzz_dep_.CoverTab[8654]++

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:791
		if len32 < 128 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:791
			_go_fuzz_dep_.CoverTab[8664]++

												return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:793
			// _ = "end of CoverTab[8664]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:794
			_go_fuzz_dep_.CoverTab[8665]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:794
			// _ = "end of CoverTab[8665]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:794
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:794
		// _ = "end of CoverTab[8654]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:794
		_go_fuzz_dep_.CoverTab[8655]++
											if len32>>((lenLen-1)*8) == 0 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:795
			_go_fuzz_dep_.CoverTab[8666]++

												return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:797
			// _ = "end of CoverTab[8666]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:798
			_go_fuzz_dep_.CoverTab[8667]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:798
			// _ = "end of CoverTab[8667]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:798
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:798
		// _ = "end of CoverTab[8655]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:798
		_go_fuzz_dep_.CoverTab[8656]++

											headerLen = 2 + uint32(lenLen)
											if headerLen+len32 < len32 {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:801
			_go_fuzz_dep_.CoverTab[8668]++

												return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:803
			// _ = "end of CoverTab[8668]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:804
			_go_fuzz_dep_.CoverTab[8669]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:804
			// _ = "end of CoverTab[8669]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:804
		}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:804
		// _ = "end of CoverTab[8656]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:804
		_go_fuzz_dep_.CoverTab[8657]++
											length = headerLen + len32
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:805
		// _ = "end of CoverTab[8657]"
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:806
	// _ = "end of CoverTab[8641]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:806
	_go_fuzz_dep_.CoverTab[8642]++

										if int(length) < 0 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:808
		_go_fuzz_dep_.CoverTab[8670]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:808
		return !s.ReadBytes((*[]byte)(out), int(length))
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:808
		// _ = "end of CoverTab[8670]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:808
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:808
		_go_fuzz_dep_.CoverTab[8671]++
											return false
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:809
		// _ = "end of CoverTab[8671]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:810
		_go_fuzz_dep_.CoverTab[8672]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:810
		// _ = "end of CoverTab[8672]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:810
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:810
	// _ = "end of CoverTab[8642]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:810
	_go_fuzz_dep_.CoverTab[8643]++
										if skipHeader && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:811
		_go_fuzz_dep_.CoverTab[8673]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:811
		return !out.Skip(int(headerLen))
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:811
		// _ = "end of CoverTab[8673]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:811
	}() {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:811
		_go_fuzz_dep_.CoverTab[8674]++
											panic("cryptobyte: internal error")
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:812
		// _ = "end of CoverTab[8674]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:813
		_go_fuzz_dep_.CoverTab[8675]++
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:813
		// _ = "end of CoverTab[8675]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:813
	}
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:813
	// _ = "end of CoverTab[8643]"
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:813
	_go_fuzz_dep_.CoverTab[8644]++

										return true
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:815
	// _ = "end of CoverTab[8644]"
}

//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:816
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go:816
var _ = _go_fuzz_dep_.CoverTab
