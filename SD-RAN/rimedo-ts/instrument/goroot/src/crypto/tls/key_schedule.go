// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/tls/key_schedule.go:5
package tls

//line /usr/local/go/src/crypto/tls/key_schedule.go:5
import (
//line /usr/local/go/src/crypto/tls/key_schedule.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/key_schedule.go:5
)
//line /usr/local/go/src/crypto/tls/key_schedule.go:5
import (
//line /usr/local/go/src/crypto/tls/key_schedule.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/key_schedule.go:5
)

import (
	"crypto/ecdh"
	"crypto/hmac"
	"errors"
	"fmt"
	"hash"
	"io"

	"golang.org/x/crypto/cryptobyte"
	"golang.org/x/crypto/hkdf"
)

//line /usr/local/go/src/crypto/tls/key_schedule.go:22
const (
	resumptionBinderLabel		= "res binder"
	clientHandshakeTrafficLabel	= "c hs traffic"
	serverHandshakeTrafficLabel	= "s hs traffic"
	clientApplicationTrafficLabel	= "c ap traffic"
	serverApplicationTrafficLabel	= "s ap traffic"
	exporterLabel			= "exp master"
	resumptionLabel			= "res master"
	trafficUpdateLabel		= "traffic upd"
)

// expandLabel implements HKDF-Expand-Label from RFC 8446, Section 7.1.
func (c *cipherSuiteTLS13) expandLabel(secret []byte, label string, context []byte, length int) []byte {
//line /usr/local/go/src/crypto/tls/key_schedule.go:34
	_go_fuzz_dep_.CoverTab[24881]++
							var hkdfLabel cryptobyte.Builder
							hkdfLabel.AddUint16(uint16(length))
							hkdfLabel.AddUint8LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/key_schedule.go:37
		_go_fuzz_dep_.CoverTab[24886]++
								b.AddBytes([]byte("tls13 "))
								b.AddBytes([]byte(label))
//line /usr/local/go/src/crypto/tls/key_schedule.go:39
		// _ = "end of CoverTab[24886]"
	})
//line /usr/local/go/src/crypto/tls/key_schedule.go:40
	// _ = "end of CoverTab[24881]"
//line /usr/local/go/src/crypto/tls/key_schedule.go:40
	_go_fuzz_dep_.CoverTab[24882]++
							hkdfLabel.AddUint8LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/key_schedule.go:41
		_go_fuzz_dep_.CoverTab[24887]++
								b.AddBytes(context)
//line /usr/local/go/src/crypto/tls/key_schedule.go:42
		// _ = "end of CoverTab[24887]"
	})
//line /usr/local/go/src/crypto/tls/key_schedule.go:43
	// _ = "end of CoverTab[24882]"
//line /usr/local/go/src/crypto/tls/key_schedule.go:43
	_go_fuzz_dep_.CoverTab[24883]++
							hkdfLabelBytes, err := hkdfLabel.Bytes()
							if err != nil {
//line /usr/local/go/src/crypto/tls/key_schedule.go:45
		_go_fuzz_dep_.CoverTab[24888]++

//line /usr/local/go/src/crypto/tls/key_schedule.go:58
		panic(fmt.Errorf("failed to construct HKDF label: %s", err))
//line /usr/local/go/src/crypto/tls/key_schedule.go:58
		// _ = "end of CoverTab[24888]"
	} else {
//line /usr/local/go/src/crypto/tls/key_schedule.go:59
		_go_fuzz_dep_.CoverTab[24889]++
//line /usr/local/go/src/crypto/tls/key_schedule.go:59
		// _ = "end of CoverTab[24889]"
//line /usr/local/go/src/crypto/tls/key_schedule.go:59
	}
//line /usr/local/go/src/crypto/tls/key_schedule.go:59
	// _ = "end of CoverTab[24883]"
//line /usr/local/go/src/crypto/tls/key_schedule.go:59
	_go_fuzz_dep_.CoverTab[24884]++
							out := make([]byte, length)
							n, err := hkdf.Expand(c.hash.New, secret, hkdfLabelBytes).Read(out)
							if err != nil || func() bool {
//line /usr/local/go/src/crypto/tls/key_schedule.go:62
		_go_fuzz_dep_.CoverTab[24890]++
//line /usr/local/go/src/crypto/tls/key_schedule.go:62
		return n != length
//line /usr/local/go/src/crypto/tls/key_schedule.go:62
		// _ = "end of CoverTab[24890]"
//line /usr/local/go/src/crypto/tls/key_schedule.go:62
	}() {
//line /usr/local/go/src/crypto/tls/key_schedule.go:62
		_go_fuzz_dep_.CoverTab[24891]++
								panic("tls: HKDF-Expand-Label invocation failed unexpectedly")
//line /usr/local/go/src/crypto/tls/key_schedule.go:63
		// _ = "end of CoverTab[24891]"
	} else {
//line /usr/local/go/src/crypto/tls/key_schedule.go:64
		_go_fuzz_dep_.CoverTab[24892]++
//line /usr/local/go/src/crypto/tls/key_schedule.go:64
		// _ = "end of CoverTab[24892]"
//line /usr/local/go/src/crypto/tls/key_schedule.go:64
	}
//line /usr/local/go/src/crypto/tls/key_schedule.go:64
	// _ = "end of CoverTab[24884]"
//line /usr/local/go/src/crypto/tls/key_schedule.go:64
	_go_fuzz_dep_.CoverTab[24885]++
							return out
//line /usr/local/go/src/crypto/tls/key_schedule.go:65
	// _ = "end of CoverTab[24885]"
}

// deriveSecret implements Derive-Secret from RFC 8446, Section 7.1.
func (c *cipherSuiteTLS13) deriveSecret(secret []byte, label string, transcript hash.Hash) []byte {
//line /usr/local/go/src/crypto/tls/key_schedule.go:69
	_go_fuzz_dep_.CoverTab[24893]++
							if transcript == nil {
//line /usr/local/go/src/crypto/tls/key_schedule.go:70
		_go_fuzz_dep_.CoverTab[24895]++
								transcript = c.hash.New()
//line /usr/local/go/src/crypto/tls/key_schedule.go:71
		// _ = "end of CoverTab[24895]"
	} else {
//line /usr/local/go/src/crypto/tls/key_schedule.go:72
		_go_fuzz_dep_.CoverTab[24896]++
//line /usr/local/go/src/crypto/tls/key_schedule.go:72
		// _ = "end of CoverTab[24896]"
//line /usr/local/go/src/crypto/tls/key_schedule.go:72
	}
//line /usr/local/go/src/crypto/tls/key_schedule.go:72
	// _ = "end of CoverTab[24893]"
//line /usr/local/go/src/crypto/tls/key_schedule.go:72
	_go_fuzz_dep_.CoverTab[24894]++
							return c.expandLabel(secret, label, transcript.Sum(nil), c.hash.Size())
//line /usr/local/go/src/crypto/tls/key_schedule.go:73
	// _ = "end of CoverTab[24894]"
}

// extract implements HKDF-Extract with the cipher suite hash.
func (c *cipherSuiteTLS13) extract(newSecret, currentSecret []byte) []byte {
//line /usr/local/go/src/crypto/tls/key_schedule.go:77
	_go_fuzz_dep_.CoverTab[24897]++
							if newSecret == nil {
//line /usr/local/go/src/crypto/tls/key_schedule.go:78
		_go_fuzz_dep_.CoverTab[24899]++
								newSecret = make([]byte, c.hash.Size())
//line /usr/local/go/src/crypto/tls/key_schedule.go:79
		// _ = "end of CoverTab[24899]"
	} else {
//line /usr/local/go/src/crypto/tls/key_schedule.go:80
		_go_fuzz_dep_.CoverTab[24900]++
//line /usr/local/go/src/crypto/tls/key_schedule.go:80
		// _ = "end of CoverTab[24900]"
//line /usr/local/go/src/crypto/tls/key_schedule.go:80
	}
//line /usr/local/go/src/crypto/tls/key_schedule.go:80
	// _ = "end of CoverTab[24897]"
//line /usr/local/go/src/crypto/tls/key_schedule.go:80
	_go_fuzz_dep_.CoverTab[24898]++
							return hkdf.Extract(c.hash.New, newSecret, currentSecret)
//line /usr/local/go/src/crypto/tls/key_schedule.go:81
	// _ = "end of CoverTab[24898]"
}

// nextTrafficSecret generates the next traffic secret, given the current one,
//line /usr/local/go/src/crypto/tls/key_schedule.go:84
// according to RFC 8446, Section 7.2.
//line /usr/local/go/src/crypto/tls/key_schedule.go:86
func (c *cipherSuiteTLS13) nextTrafficSecret(trafficSecret []byte) []byte {
//line /usr/local/go/src/crypto/tls/key_schedule.go:86
	_go_fuzz_dep_.CoverTab[24901]++
							return c.expandLabel(trafficSecret, trafficUpdateLabel, nil, c.hash.Size())
//line /usr/local/go/src/crypto/tls/key_schedule.go:87
	// _ = "end of CoverTab[24901]"
}

// trafficKey generates traffic keys according to RFC 8446, Section 7.3.
func (c *cipherSuiteTLS13) trafficKey(trafficSecret []byte) (key, iv []byte) {
//line /usr/local/go/src/crypto/tls/key_schedule.go:91
	_go_fuzz_dep_.CoverTab[24902]++
							key = c.expandLabel(trafficSecret, "key", nil, c.keyLen)
							iv = c.expandLabel(trafficSecret, "iv", nil, aeadNonceLength)
							return
//line /usr/local/go/src/crypto/tls/key_schedule.go:94
	// _ = "end of CoverTab[24902]"
}

// finishedHash generates the Finished verify_data or PskBinderEntry according
//line /usr/local/go/src/crypto/tls/key_schedule.go:97
// to RFC 8446, Section 4.4.4. See sections 4.4 and 4.2.11.2 for the baseKey
//line /usr/local/go/src/crypto/tls/key_schedule.go:97
// selection.
//line /usr/local/go/src/crypto/tls/key_schedule.go:100
func (c *cipherSuiteTLS13) finishedHash(baseKey []byte, transcript hash.Hash) []byte {
//line /usr/local/go/src/crypto/tls/key_schedule.go:100
	_go_fuzz_dep_.CoverTab[24903]++
								finishedKey := c.expandLabel(baseKey, "finished", nil, c.hash.Size())
								verifyData := hmac.New(c.hash.New, finishedKey)
								verifyData.Write(transcript.Sum(nil))
								return verifyData.Sum(nil)
//line /usr/local/go/src/crypto/tls/key_schedule.go:104
	// _ = "end of CoverTab[24903]"
}

// exportKeyingMaterial implements RFC5705 exporters for TLS 1.3 according to
//line /usr/local/go/src/crypto/tls/key_schedule.go:107
// RFC 8446, Section 7.5.
//line /usr/local/go/src/crypto/tls/key_schedule.go:109
func (c *cipherSuiteTLS13) exportKeyingMaterial(masterSecret []byte, transcript hash.Hash) func(string, []byte, int) ([]byte, error) {
//line /usr/local/go/src/crypto/tls/key_schedule.go:109
	_go_fuzz_dep_.CoverTab[24904]++
								expMasterSecret := c.deriveSecret(masterSecret, exporterLabel, transcript)
								return func(label string, context []byte, length int) ([]byte, error) {
//line /usr/local/go/src/crypto/tls/key_schedule.go:111
		_go_fuzz_dep_.CoverTab[24905]++
									secret := c.deriveSecret(expMasterSecret, label, nil)
									h := c.hash.New()
									h.Write(context)
									return c.expandLabel(secret, "exporter", h.Sum(nil), length), nil
//line /usr/local/go/src/crypto/tls/key_schedule.go:115
		// _ = "end of CoverTab[24905]"
	}
//line /usr/local/go/src/crypto/tls/key_schedule.go:116
	// _ = "end of CoverTab[24904]"
}

// generateECDHEKey returns a PrivateKey that implements Diffie-Hellman
//line /usr/local/go/src/crypto/tls/key_schedule.go:119
// according to RFC 8446, Section 4.2.8.2.
//line /usr/local/go/src/crypto/tls/key_schedule.go:121
func generateECDHEKey(rand io.Reader, curveID CurveID) (*ecdh.PrivateKey, error) {
//line /usr/local/go/src/crypto/tls/key_schedule.go:121
	_go_fuzz_dep_.CoverTab[24906]++
								curve, ok := curveForCurveID(curveID)
								if !ok {
//line /usr/local/go/src/crypto/tls/key_schedule.go:123
		_go_fuzz_dep_.CoverTab[24908]++
									return nil, errors.New("tls: internal error: unsupported curve")
//line /usr/local/go/src/crypto/tls/key_schedule.go:124
		// _ = "end of CoverTab[24908]"
	} else {
//line /usr/local/go/src/crypto/tls/key_schedule.go:125
		_go_fuzz_dep_.CoverTab[24909]++
//line /usr/local/go/src/crypto/tls/key_schedule.go:125
		// _ = "end of CoverTab[24909]"
//line /usr/local/go/src/crypto/tls/key_schedule.go:125
	}
//line /usr/local/go/src/crypto/tls/key_schedule.go:125
	// _ = "end of CoverTab[24906]"
//line /usr/local/go/src/crypto/tls/key_schedule.go:125
	_go_fuzz_dep_.CoverTab[24907]++

								return curve.GenerateKey(rand)
//line /usr/local/go/src/crypto/tls/key_schedule.go:127
	// _ = "end of CoverTab[24907]"
}

func curveForCurveID(id CurveID) (ecdh.Curve, bool) {
//line /usr/local/go/src/crypto/tls/key_schedule.go:130
	_go_fuzz_dep_.CoverTab[24910]++
								switch id {
	case X25519:
//line /usr/local/go/src/crypto/tls/key_schedule.go:132
		_go_fuzz_dep_.CoverTab[24911]++
									return ecdh.X25519(), true
//line /usr/local/go/src/crypto/tls/key_schedule.go:133
		// _ = "end of CoverTab[24911]"
	case CurveP256:
//line /usr/local/go/src/crypto/tls/key_schedule.go:134
		_go_fuzz_dep_.CoverTab[24912]++
									return ecdh.P256(), true
//line /usr/local/go/src/crypto/tls/key_schedule.go:135
		// _ = "end of CoverTab[24912]"
	case CurveP384:
//line /usr/local/go/src/crypto/tls/key_schedule.go:136
		_go_fuzz_dep_.CoverTab[24913]++
									return ecdh.P384(), true
//line /usr/local/go/src/crypto/tls/key_schedule.go:137
		// _ = "end of CoverTab[24913]"
	case CurveP521:
//line /usr/local/go/src/crypto/tls/key_schedule.go:138
		_go_fuzz_dep_.CoverTab[24914]++
									return ecdh.P521(), true
//line /usr/local/go/src/crypto/tls/key_schedule.go:139
		// _ = "end of CoverTab[24914]"
	default:
//line /usr/local/go/src/crypto/tls/key_schedule.go:140
		_go_fuzz_dep_.CoverTab[24915]++
									return nil, false
//line /usr/local/go/src/crypto/tls/key_schedule.go:141
		// _ = "end of CoverTab[24915]"
	}
//line /usr/local/go/src/crypto/tls/key_schedule.go:142
	// _ = "end of CoverTab[24910]"
}

func curveIDForCurve(curve ecdh.Curve) (CurveID, bool) {
//line /usr/local/go/src/crypto/tls/key_schedule.go:145
	_go_fuzz_dep_.CoverTab[24916]++
								switch curve {
	case ecdh.X25519():
//line /usr/local/go/src/crypto/tls/key_schedule.go:147
		_go_fuzz_dep_.CoverTab[24917]++
									return X25519, true
//line /usr/local/go/src/crypto/tls/key_schedule.go:148
		// _ = "end of CoverTab[24917]"
	case ecdh.P256():
//line /usr/local/go/src/crypto/tls/key_schedule.go:149
		_go_fuzz_dep_.CoverTab[24918]++
									return CurveP256, true
//line /usr/local/go/src/crypto/tls/key_schedule.go:150
		// _ = "end of CoverTab[24918]"
	case ecdh.P384():
//line /usr/local/go/src/crypto/tls/key_schedule.go:151
		_go_fuzz_dep_.CoverTab[24919]++
									return CurveP384, true
//line /usr/local/go/src/crypto/tls/key_schedule.go:152
		// _ = "end of CoverTab[24919]"
	case ecdh.P521():
//line /usr/local/go/src/crypto/tls/key_schedule.go:153
		_go_fuzz_dep_.CoverTab[24920]++
									return CurveP521, true
//line /usr/local/go/src/crypto/tls/key_schedule.go:154
		// _ = "end of CoverTab[24920]"
	default:
//line /usr/local/go/src/crypto/tls/key_schedule.go:155
		_go_fuzz_dep_.CoverTab[24921]++
									return 0, false
//line /usr/local/go/src/crypto/tls/key_schedule.go:156
		// _ = "end of CoverTab[24921]"
	}
//line /usr/local/go/src/crypto/tls/key_schedule.go:157
	// _ = "end of CoverTab[24916]"
}

//line /usr/local/go/src/crypto/tls/key_schedule.go:158
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/key_schedule.go:158
var _ = _go_fuzz_dep_.CoverTab
