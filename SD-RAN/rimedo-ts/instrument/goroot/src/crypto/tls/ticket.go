// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/tls/ticket.go:5
package tls

//line /usr/local/go/src/crypto/tls/ticket.go:5
import (
//line /usr/local/go/src/crypto/tls/ticket.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/ticket.go:5
)
//line /usr/local/go/src/crypto/tls/ticket.go:5
import (
//line /usr/local/go/src/crypto/tls/ticket.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/ticket.go:5
)

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"io"

	"golang.org/x/crypto/cryptobyte"
)

// sessionState contains the information that is serialized into a session
//line /usr/local/go/src/crypto/tls/ticket.go:20
// ticket in order to later resume a connection.
//line /usr/local/go/src/crypto/tls/ticket.go:22
type sessionState struct {
	vers		uint16
	cipherSuite	uint16
	createdAt	uint64
	masterSecret	[]byte	// opaque master_secret<1..2^16-1>;
	// struct { opaque certificate<1..2^24-1> } Certificate;
	certificates	[][]byte	// Certificate certificate_list<0..2^24-1>;

	// usedOldKey is true if the ticket from which this session came from
	// was encrypted with an older key and thus should be refreshed.
	usedOldKey	bool
}

func (m *sessionState) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/ticket.go:35
	_go_fuzz_dep_.CoverTab[24996]++
							var b cryptobyte.Builder
							b.AddUint16(m.vers)
							b.AddUint16(m.cipherSuite)
							addUint64(&b, m.createdAt)
							b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/ticket.go:40
		_go_fuzz_dep_.CoverTab[24999]++
								b.AddBytes(m.masterSecret)
//line /usr/local/go/src/crypto/tls/ticket.go:41
		// _ = "end of CoverTab[24999]"
	})
//line /usr/local/go/src/crypto/tls/ticket.go:42
	// _ = "end of CoverTab[24996]"
//line /usr/local/go/src/crypto/tls/ticket.go:42
	_go_fuzz_dep_.CoverTab[24997]++
							b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/ticket.go:43
		_go_fuzz_dep_.CoverTab[25000]++
								for _, cert := range m.certificates {
//line /usr/local/go/src/crypto/tls/ticket.go:44
			_go_fuzz_dep_.CoverTab[25001]++
									b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/ticket.go:45
				_go_fuzz_dep_.CoverTab[25002]++
										b.AddBytes(cert)
//line /usr/local/go/src/crypto/tls/ticket.go:46
				// _ = "end of CoverTab[25002]"
			})
//line /usr/local/go/src/crypto/tls/ticket.go:47
			// _ = "end of CoverTab[25001]"
		}
//line /usr/local/go/src/crypto/tls/ticket.go:48
		// _ = "end of CoverTab[25000]"
	})
//line /usr/local/go/src/crypto/tls/ticket.go:49
	// _ = "end of CoverTab[24997]"
//line /usr/local/go/src/crypto/tls/ticket.go:49
	_go_fuzz_dep_.CoverTab[24998]++
							return b.Bytes()
//line /usr/local/go/src/crypto/tls/ticket.go:50
	// _ = "end of CoverTab[24998]"
}

func (m *sessionState) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/ticket.go:53
	_go_fuzz_dep_.CoverTab[25003]++
							*m = sessionState{usedOldKey: m.usedOldKey}
							s := cryptobyte.String(data)
							if ok := s.ReadUint16(&m.vers) && func() bool {
//line /usr/local/go/src/crypto/tls/ticket.go:56
		_go_fuzz_dep_.CoverTab[25007]++
//line /usr/local/go/src/crypto/tls/ticket.go:56
		return s.ReadUint16(&m.cipherSuite)
								// _ = "end of CoverTab[25007]"
//line /usr/local/go/src/crypto/tls/ticket.go:57
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/ticket.go:57
		_go_fuzz_dep_.CoverTab[25008]++
//line /usr/local/go/src/crypto/tls/ticket.go:57
		return readUint64(&s, &m.createdAt)
								// _ = "end of CoverTab[25008]"
//line /usr/local/go/src/crypto/tls/ticket.go:58
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/ticket.go:58
		_go_fuzz_dep_.CoverTab[25009]++
//line /usr/local/go/src/crypto/tls/ticket.go:58
		return readUint16LengthPrefixed(&s, &m.masterSecret)
								// _ = "end of CoverTab[25009]"
//line /usr/local/go/src/crypto/tls/ticket.go:59
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/ticket.go:59
		_go_fuzz_dep_.CoverTab[25010]++
//line /usr/local/go/src/crypto/tls/ticket.go:59
		return len(m.masterSecret) != 0
								// _ = "end of CoverTab[25010]"
//line /usr/local/go/src/crypto/tls/ticket.go:60
	}(); !ok {
//line /usr/local/go/src/crypto/tls/ticket.go:60
		_go_fuzz_dep_.CoverTab[25011]++
								return false
//line /usr/local/go/src/crypto/tls/ticket.go:61
		// _ = "end of CoverTab[25011]"
	} else {
//line /usr/local/go/src/crypto/tls/ticket.go:62
		_go_fuzz_dep_.CoverTab[25012]++
//line /usr/local/go/src/crypto/tls/ticket.go:62
		// _ = "end of CoverTab[25012]"
//line /usr/local/go/src/crypto/tls/ticket.go:62
	}
//line /usr/local/go/src/crypto/tls/ticket.go:62
	// _ = "end of CoverTab[25003]"
//line /usr/local/go/src/crypto/tls/ticket.go:62
	_go_fuzz_dep_.CoverTab[25004]++
							var certList cryptobyte.String
							if !s.ReadUint24LengthPrefixed(&certList) {
//line /usr/local/go/src/crypto/tls/ticket.go:64
		_go_fuzz_dep_.CoverTab[25013]++
								return false
//line /usr/local/go/src/crypto/tls/ticket.go:65
		// _ = "end of CoverTab[25013]"
	} else {
//line /usr/local/go/src/crypto/tls/ticket.go:66
		_go_fuzz_dep_.CoverTab[25014]++
//line /usr/local/go/src/crypto/tls/ticket.go:66
		// _ = "end of CoverTab[25014]"
//line /usr/local/go/src/crypto/tls/ticket.go:66
	}
//line /usr/local/go/src/crypto/tls/ticket.go:66
	// _ = "end of CoverTab[25004]"
//line /usr/local/go/src/crypto/tls/ticket.go:66
	_go_fuzz_dep_.CoverTab[25005]++
							for !certList.Empty() {
//line /usr/local/go/src/crypto/tls/ticket.go:67
		_go_fuzz_dep_.CoverTab[25015]++
								var cert []byte
								if !readUint24LengthPrefixed(&certList, &cert) {
//line /usr/local/go/src/crypto/tls/ticket.go:69
			_go_fuzz_dep_.CoverTab[25017]++
									return false
//line /usr/local/go/src/crypto/tls/ticket.go:70
			// _ = "end of CoverTab[25017]"
		} else {
//line /usr/local/go/src/crypto/tls/ticket.go:71
			_go_fuzz_dep_.CoverTab[25018]++
//line /usr/local/go/src/crypto/tls/ticket.go:71
			// _ = "end of CoverTab[25018]"
//line /usr/local/go/src/crypto/tls/ticket.go:71
		}
//line /usr/local/go/src/crypto/tls/ticket.go:71
		// _ = "end of CoverTab[25015]"
//line /usr/local/go/src/crypto/tls/ticket.go:71
		_go_fuzz_dep_.CoverTab[25016]++
								m.certificates = append(m.certificates, cert)
//line /usr/local/go/src/crypto/tls/ticket.go:72
		// _ = "end of CoverTab[25016]"
	}
//line /usr/local/go/src/crypto/tls/ticket.go:73
	// _ = "end of CoverTab[25005]"
//line /usr/local/go/src/crypto/tls/ticket.go:73
	_go_fuzz_dep_.CoverTab[25006]++
							return s.Empty()
//line /usr/local/go/src/crypto/tls/ticket.go:74
	// _ = "end of CoverTab[25006]"
}

// sessionStateTLS13 is the content of a TLS 1.3 session ticket. Its first
//line /usr/local/go/src/crypto/tls/ticket.go:77
// version (revision = 0) doesn't carry any of the information needed for 0-RTT
//line /usr/local/go/src/crypto/tls/ticket.go:77
// validation and the nonce is always empty.
//line /usr/local/go/src/crypto/tls/ticket.go:80
type sessionStateTLS13 struct {
	// uint8 version  = 0x0304;
	// uint8 revision = 0;
	cipherSuite		uint16
	createdAt		uint64
	resumptionSecret	[]byte		// opaque resumption_master_secret<1..2^8-1>;
	certificate		Certificate	// CertificateEntry certificate_list<0..2^24-1>;
}

func (m *sessionStateTLS13) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/ticket.go:89
	_go_fuzz_dep_.CoverTab[25019]++
							var b cryptobyte.Builder
							b.AddUint16(VersionTLS13)
							b.AddUint8(0)
							b.AddUint16(m.cipherSuite)
							addUint64(&b, m.createdAt)
							b.AddUint8LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/ticket.go:95
		_go_fuzz_dep_.CoverTab[25021]++
								b.AddBytes(m.resumptionSecret)
//line /usr/local/go/src/crypto/tls/ticket.go:96
		// _ = "end of CoverTab[25021]"
	})
//line /usr/local/go/src/crypto/tls/ticket.go:97
	// _ = "end of CoverTab[25019]"
//line /usr/local/go/src/crypto/tls/ticket.go:97
	_go_fuzz_dep_.CoverTab[25020]++
							marshalCertificate(&b, m.certificate)
							return b.Bytes()
//line /usr/local/go/src/crypto/tls/ticket.go:99
	// _ = "end of CoverTab[25020]"
}

func (m *sessionStateTLS13) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/ticket.go:102
	_go_fuzz_dep_.CoverTab[25022]++
							*m = sessionStateTLS13{}
							s := cryptobyte.String(data)
							var version uint16
							var revision uint8
							return s.ReadUint16(&version) && func() bool {
//line /usr/local/go/src/crypto/tls/ticket.go:107
		_go_fuzz_dep_.CoverTab[25023]++
//line /usr/local/go/src/crypto/tls/ticket.go:107
		return version == VersionTLS13
								// _ = "end of CoverTab[25023]"
//line /usr/local/go/src/crypto/tls/ticket.go:108
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/ticket.go:108
		_go_fuzz_dep_.CoverTab[25024]++
//line /usr/local/go/src/crypto/tls/ticket.go:108
		return s.ReadUint8(&revision)
								// _ = "end of CoverTab[25024]"
//line /usr/local/go/src/crypto/tls/ticket.go:109
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/ticket.go:109
		_go_fuzz_dep_.CoverTab[25025]++
//line /usr/local/go/src/crypto/tls/ticket.go:109
		return revision == 0
								// _ = "end of CoverTab[25025]"
//line /usr/local/go/src/crypto/tls/ticket.go:110
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/ticket.go:110
		_go_fuzz_dep_.CoverTab[25026]++
//line /usr/local/go/src/crypto/tls/ticket.go:110
		return s.ReadUint16(&m.cipherSuite)
								// _ = "end of CoverTab[25026]"
//line /usr/local/go/src/crypto/tls/ticket.go:111
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/ticket.go:111
		_go_fuzz_dep_.CoverTab[25027]++
//line /usr/local/go/src/crypto/tls/ticket.go:111
		return readUint64(&s, &m.createdAt)
								// _ = "end of CoverTab[25027]"
//line /usr/local/go/src/crypto/tls/ticket.go:112
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/ticket.go:112
		_go_fuzz_dep_.CoverTab[25028]++
//line /usr/local/go/src/crypto/tls/ticket.go:112
		return readUint8LengthPrefixed(&s, &m.resumptionSecret)
								// _ = "end of CoverTab[25028]"
//line /usr/local/go/src/crypto/tls/ticket.go:113
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/ticket.go:113
		_go_fuzz_dep_.CoverTab[25029]++
//line /usr/local/go/src/crypto/tls/ticket.go:113
		return len(m.resumptionSecret) != 0
								// _ = "end of CoverTab[25029]"
//line /usr/local/go/src/crypto/tls/ticket.go:114
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/ticket.go:114
		_go_fuzz_dep_.CoverTab[25030]++
//line /usr/local/go/src/crypto/tls/ticket.go:114
		return unmarshalCertificate(&s, &m.certificate)
								// _ = "end of CoverTab[25030]"
//line /usr/local/go/src/crypto/tls/ticket.go:115
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/ticket.go:115
		_go_fuzz_dep_.CoverTab[25031]++
//line /usr/local/go/src/crypto/tls/ticket.go:115
		return s.Empty()
								// _ = "end of CoverTab[25031]"
//line /usr/local/go/src/crypto/tls/ticket.go:116
	}()
//line /usr/local/go/src/crypto/tls/ticket.go:116
	// _ = "end of CoverTab[25022]"
}

func (c *Conn) encryptTicket(state []byte) ([]byte, error) {
//line /usr/local/go/src/crypto/tls/ticket.go:119
	_go_fuzz_dep_.CoverTab[25032]++
							if len(c.ticketKeys) == 0 {
//line /usr/local/go/src/crypto/tls/ticket.go:120
		_go_fuzz_dep_.CoverTab[25036]++
								return nil, errors.New("tls: internal error: session ticket keys unavailable")
//line /usr/local/go/src/crypto/tls/ticket.go:121
		// _ = "end of CoverTab[25036]"
	} else {
//line /usr/local/go/src/crypto/tls/ticket.go:122
		_go_fuzz_dep_.CoverTab[25037]++
//line /usr/local/go/src/crypto/tls/ticket.go:122
		// _ = "end of CoverTab[25037]"
//line /usr/local/go/src/crypto/tls/ticket.go:122
	}
//line /usr/local/go/src/crypto/tls/ticket.go:122
	// _ = "end of CoverTab[25032]"
//line /usr/local/go/src/crypto/tls/ticket.go:122
	_go_fuzz_dep_.CoverTab[25033]++

							encrypted := make([]byte, ticketKeyNameLen+aes.BlockSize+len(state)+sha256.Size)
							keyName := encrypted[:ticketKeyNameLen]
							iv := encrypted[ticketKeyNameLen : ticketKeyNameLen+aes.BlockSize]
							macBytes := encrypted[len(encrypted)-sha256.Size:]

							if _, err := io.ReadFull(c.config.rand(), iv); err != nil {
//line /usr/local/go/src/crypto/tls/ticket.go:129
		_go_fuzz_dep_.CoverTab[25038]++
								return nil, err
//line /usr/local/go/src/crypto/tls/ticket.go:130
		// _ = "end of CoverTab[25038]"
	} else {
//line /usr/local/go/src/crypto/tls/ticket.go:131
		_go_fuzz_dep_.CoverTab[25039]++
//line /usr/local/go/src/crypto/tls/ticket.go:131
		// _ = "end of CoverTab[25039]"
//line /usr/local/go/src/crypto/tls/ticket.go:131
	}
//line /usr/local/go/src/crypto/tls/ticket.go:131
	// _ = "end of CoverTab[25033]"
//line /usr/local/go/src/crypto/tls/ticket.go:131
	_go_fuzz_dep_.CoverTab[25034]++
							key := c.ticketKeys[0]
							copy(keyName, key.keyName[:])
							block, err := aes.NewCipher(key.aesKey[:])
							if err != nil {
//line /usr/local/go/src/crypto/tls/ticket.go:135
		_go_fuzz_dep_.CoverTab[25040]++
								return nil, errors.New("tls: failed to create cipher while encrypting ticket: " + err.Error())
//line /usr/local/go/src/crypto/tls/ticket.go:136
		// _ = "end of CoverTab[25040]"
	} else {
//line /usr/local/go/src/crypto/tls/ticket.go:137
		_go_fuzz_dep_.CoverTab[25041]++
//line /usr/local/go/src/crypto/tls/ticket.go:137
		// _ = "end of CoverTab[25041]"
//line /usr/local/go/src/crypto/tls/ticket.go:137
	}
//line /usr/local/go/src/crypto/tls/ticket.go:137
	// _ = "end of CoverTab[25034]"
//line /usr/local/go/src/crypto/tls/ticket.go:137
	_go_fuzz_dep_.CoverTab[25035]++
							cipher.NewCTR(block, iv).XORKeyStream(encrypted[ticketKeyNameLen+aes.BlockSize:], state)

							mac := hmac.New(sha256.New, key.hmacKey[:])
							mac.Write(encrypted[:len(encrypted)-sha256.Size])
							mac.Sum(macBytes[:0])

							return encrypted, nil
//line /usr/local/go/src/crypto/tls/ticket.go:144
	// _ = "end of CoverTab[25035]"
}

func (c *Conn) decryptTicket(encrypted []byte) (plaintext []byte, usedOldKey bool) {
//line /usr/local/go/src/crypto/tls/ticket.go:147
	_go_fuzz_dep_.CoverTab[25042]++
							if len(encrypted) < ticketKeyNameLen+aes.BlockSize+sha256.Size {
//line /usr/local/go/src/crypto/tls/ticket.go:148
		_go_fuzz_dep_.CoverTab[25048]++
								return nil, false
//line /usr/local/go/src/crypto/tls/ticket.go:149
		// _ = "end of CoverTab[25048]"
	} else {
//line /usr/local/go/src/crypto/tls/ticket.go:150
		_go_fuzz_dep_.CoverTab[25049]++
//line /usr/local/go/src/crypto/tls/ticket.go:150
		// _ = "end of CoverTab[25049]"
//line /usr/local/go/src/crypto/tls/ticket.go:150
	}
//line /usr/local/go/src/crypto/tls/ticket.go:150
	// _ = "end of CoverTab[25042]"
//line /usr/local/go/src/crypto/tls/ticket.go:150
	_go_fuzz_dep_.CoverTab[25043]++

							keyName := encrypted[:ticketKeyNameLen]
							iv := encrypted[ticketKeyNameLen : ticketKeyNameLen+aes.BlockSize]
							macBytes := encrypted[len(encrypted)-sha256.Size:]
							ciphertext := encrypted[ticketKeyNameLen+aes.BlockSize : len(encrypted)-sha256.Size]

							keyIndex := -1
							for i, candidateKey := range c.ticketKeys {
//line /usr/local/go/src/crypto/tls/ticket.go:158
		_go_fuzz_dep_.CoverTab[25050]++
								if bytes.Equal(keyName, candidateKey.keyName[:]) {
//line /usr/local/go/src/crypto/tls/ticket.go:159
			_go_fuzz_dep_.CoverTab[25051]++
									keyIndex = i
									break
//line /usr/local/go/src/crypto/tls/ticket.go:161
			// _ = "end of CoverTab[25051]"
		} else {
//line /usr/local/go/src/crypto/tls/ticket.go:162
			_go_fuzz_dep_.CoverTab[25052]++
//line /usr/local/go/src/crypto/tls/ticket.go:162
			// _ = "end of CoverTab[25052]"
//line /usr/local/go/src/crypto/tls/ticket.go:162
		}
//line /usr/local/go/src/crypto/tls/ticket.go:162
		// _ = "end of CoverTab[25050]"
	}
//line /usr/local/go/src/crypto/tls/ticket.go:163
	// _ = "end of CoverTab[25043]"
//line /usr/local/go/src/crypto/tls/ticket.go:163
	_go_fuzz_dep_.CoverTab[25044]++
							if keyIndex == -1 {
//line /usr/local/go/src/crypto/tls/ticket.go:164
		_go_fuzz_dep_.CoverTab[25053]++
								return nil, false
//line /usr/local/go/src/crypto/tls/ticket.go:165
		// _ = "end of CoverTab[25053]"
	} else {
//line /usr/local/go/src/crypto/tls/ticket.go:166
		_go_fuzz_dep_.CoverTab[25054]++
//line /usr/local/go/src/crypto/tls/ticket.go:166
		// _ = "end of CoverTab[25054]"
//line /usr/local/go/src/crypto/tls/ticket.go:166
	}
//line /usr/local/go/src/crypto/tls/ticket.go:166
	// _ = "end of CoverTab[25044]"
//line /usr/local/go/src/crypto/tls/ticket.go:166
	_go_fuzz_dep_.CoverTab[25045]++
							key := &c.ticketKeys[keyIndex]

							mac := hmac.New(sha256.New, key.hmacKey[:])
							mac.Write(encrypted[:len(encrypted)-sha256.Size])
							expected := mac.Sum(nil)

							if subtle.ConstantTimeCompare(macBytes, expected) != 1 {
//line /usr/local/go/src/crypto/tls/ticket.go:173
		_go_fuzz_dep_.CoverTab[25055]++
								return nil, false
//line /usr/local/go/src/crypto/tls/ticket.go:174
		// _ = "end of CoverTab[25055]"
	} else {
//line /usr/local/go/src/crypto/tls/ticket.go:175
		_go_fuzz_dep_.CoverTab[25056]++
//line /usr/local/go/src/crypto/tls/ticket.go:175
		// _ = "end of CoverTab[25056]"
//line /usr/local/go/src/crypto/tls/ticket.go:175
	}
//line /usr/local/go/src/crypto/tls/ticket.go:175
	// _ = "end of CoverTab[25045]"
//line /usr/local/go/src/crypto/tls/ticket.go:175
	_go_fuzz_dep_.CoverTab[25046]++

							block, err := aes.NewCipher(key.aesKey[:])
							if err != nil {
//line /usr/local/go/src/crypto/tls/ticket.go:178
		_go_fuzz_dep_.CoverTab[25057]++
								return nil, false
//line /usr/local/go/src/crypto/tls/ticket.go:179
		// _ = "end of CoverTab[25057]"
	} else {
//line /usr/local/go/src/crypto/tls/ticket.go:180
		_go_fuzz_dep_.CoverTab[25058]++
//line /usr/local/go/src/crypto/tls/ticket.go:180
		// _ = "end of CoverTab[25058]"
//line /usr/local/go/src/crypto/tls/ticket.go:180
	}
//line /usr/local/go/src/crypto/tls/ticket.go:180
	// _ = "end of CoverTab[25046]"
//line /usr/local/go/src/crypto/tls/ticket.go:180
	_go_fuzz_dep_.CoverTab[25047]++
							plaintext = make([]byte, len(ciphertext))
							cipher.NewCTR(block, iv).XORKeyStream(plaintext, ciphertext)

							return plaintext, keyIndex > 0
//line /usr/local/go/src/crypto/tls/ticket.go:184
	// _ = "end of CoverTab[25047]"
}

//line /usr/local/go/src/crypto/tls/ticket.go:185
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/ticket.go:185
var _ = _go_fuzz_dep_.CoverTab
