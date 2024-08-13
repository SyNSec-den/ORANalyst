// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/tls/handshake_messages.go:5
package tls

//line /usr/local/go/src/crypto/tls/handshake_messages.go:5
import (
//line /usr/local/go/src/crypto/tls/handshake_messages.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:5
)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:5
import (
//line /usr/local/go/src/crypto/tls/handshake_messages.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:5
)

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/cryptobyte"
)

// The marshalingFunction type is an adapter to allow the use of ordinary
//line /usr/local/go/src/crypto/tls/handshake_messages.go:15
// functions as cryptobyte.MarshalingValue.
//line /usr/local/go/src/crypto/tls/handshake_messages.go:17
type marshalingFunction func(b *cryptobyte.Builder) error

func (f marshalingFunction) Marshal(b *cryptobyte.Builder) error {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:19
	_go_fuzz_dep_.CoverTab[23003]++
								return f(b)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:20
	// _ = "end of CoverTab[23003]"
}

// addBytesWithLength appends a sequence of bytes to the cryptobyte.Builder. If
//line /usr/local/go/src/crypto/tls/handshake_messages.go:23
// the length of the sequence is not the value specified, it produces an error.
//line /usr/local/go/src/crypto/tls/handshake_messages.go:25
func addBytesWithLength(b *cryptobyte.Builder, v []byte, n int) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:25
	_go_fuzz_dep_.CoverTab[23004]++
								b.AddValue(marshalingFunction(func(b *cryptobyte.Builder) error {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:26
		_go_fuzz_dep_.CoverTab[23005]++
									if len(v) != n {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:27
			_go_fuzz_dep_.CoverTab[23007]++
										return fmt.Errorf("invalid value length: expected %d, got %d", n, len(v))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:28
			// _ = "end of CoverTab[23007]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:29
			_go_fuzz_dep_.CoverTab[23008]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:29
			// _ = "end of CoverTab[23008]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:29
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:29
		// _ = "end of CoverTab[23005]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:29
		_go_fuzz_dep_.CoverTab[23006]++
									b.AddBytes(v)
									return nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:31
		// _ = "end of CoverTab[23006]"
	}))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:32
	// _ = "end of CoverTab[23004]"
}

// addUint64 appends a big-endian, 64-bit value to the cryptobyte.Builder.
func addUint64(b *cryptobyte.Builder, v uint64) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:36
	_go_fuzz_dep_.CoverTab[23009]++
								b.AddUint32(uint32(v >> 32))
								b.AddUint32(uint32(v))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:38
	// _ = "end of CoverTab[23009]"
}

// readUint64 decodes a big-endian, 64-bit value into out and advances over it.
//line /usr/local/go/src/crypto/tls/handshake_messages.go:41
// It reports whether the read was successful.
//line /usr/local/go/src/crypto/tls/handshake_messages.go:43
func readUint64(s *cryptobyte.String, out *uint64) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:43
	_go_fuzz_dep_.CoverTab[23010]++
								var hi, lo uint32
								if !s.ReadUint32(&hi) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:45
		_go_fuzz_dep_.CoverTab[23012]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:45
		return !s.ReadUint32(&lo)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:45
		// _ = "end of CoverTab[23012]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:45
	}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:45
		_go_fuzz_dep_.CoverTab[23013]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:46
		// _ = "end of CoverTab[23013]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:47
		_go_fuzz_dep_.CoverTab[23014]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:47
		// _ = "end of CoverTab[23014]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:47
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:47
	// _ = "end of CoverTab[23010]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:47
	_go_fuzz_dep_.CoverTab[23011]++
								*out = uint64(hi)<<32 | uint64(lo)
								return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:49
	// _ = "end of CoverTab[23011]"
}

// readUint8LengthPrefixed acts like s.ReadUint8LengthPrefixed, but targets a
//line /usr/local/go/src/crypto/tls/handshake_messages.go:52
// []byte instead of a cryptobyte.String.
//line /usr/local/go/src/crypto/tls/handshake_messages.go:54
func readUint8LengthPrefixed(s *cryptobyte.String, out *[]byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:54
	_go_fuzz_dep_.CoverTab[23015]++
								return s.ReadUint8LengthPrefixed((*cryptobyte.String)(out))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:55
	// _ = "end of CoverTab[23015]"
}

// readUint16LengthPrefixed acts like s.ReadUint16LengthPrefixed, but targets a
//line /usr/local/go/src/crypto/tls/handshake_messages.go:58
// []byte instead of a cryptobyte.String.
//line /usr/local/go/src/crypto/tls/handshake_messages.go:60
func readUint16LengthPrefixed(s *cryptobyte.String, out *[]byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:60
	_go_fuzz_dep_.CoverTab[23016]++
								return s.ReadUint16LengthPrefixed((*cryptobyte.String)(out))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:61
	// _ = "end of CoverTab[23016]"
}

// readUint24LengthPrefixed acts like s.ReadUint24LengthPrefixed, but targets a
//line /usr/local/go/src/crypto/tls/handshake_messages.go:64
// []byte instead of a cryptobyte.String.
//line /usr/local/go/src/crypto/tls/handshake_messages.go:66
func readUint24LengthPrefixed(s *cryptobyte.String, out *[]byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:66
	_go_fuzz_dep_.CoverTab[23017]++
								return s.ReadUint24LengthPrefixed((*cryptobyte.String)(out))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:67
	// _ = "end of CoverTab[23017]"
}

type clientHelloMsg struct {
	raw					[]byte
	vers					uint16
	random					[]byte
	sessionId				[]byte
	cipherSuites				[]uint16
	compressionMethods			[]uint8
	serverName				string
	ocspStapling				bool
	supportedCurves				[]CurveID
	supportedPoints				[]uint8
	ticketSupported				bool
	sessionTicket				[]uint8
	supportedSignatureAlgorithms		[]SignatureScheme
	supportedSignatureAlgorithmsCert	[]SignatureScheme
	secureRenegotiationSupported		bool
	secureRenegotiation			[]byte
	alpnProtocols				[]string
	scts					bool
	supportedVersions			[]uint16
	cookie					[]byte
	keyShares				[]keyShare
	earlyData				bool
	pskModes				[]uint8
	pskIdentities				[]pskIdentity
	pskBinders				[][]byte
}

func (m *clientHelloMsg) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:98
	_go_fuzz_dep_.CoverTab[23018]++
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:99
		_go_fuzz_dep_.CoverTab[23038]++
									return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:100
		// _ = "end of CoverTab[23038]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:101
		_go_fuzz_dep_.CoverTab[23039]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:101
		// _ = "end of CoverTab[23039]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:101
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:101
	// _ = "end of CoverTab[23018]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:101
	_go_fuzz_dep_.CoverTab[23019]++

								var exts cryptobyte.Builder
								if len(m.serverName) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:104
		_go_fuzz_dep_.CoverTab[23040]++

									exts.AddUint16(extensionServerName)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:107
			_go_fuzz_dep_.CoverTab[23041]++
										exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:108
				_go_fuzz_dep_.CoverTab[23042]++
											exts.AddUint8(0)
											exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:110
					_go_fuzz_dep_.CoverTab[23043]++
												exts.AddBytes([]byte(m.serverName))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:111
					// _ = "end of CoverTab[23043]"
				})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:112
				// _ = "end of CoverTab[23042]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:113
			// _ = "end of CoverTab[23041]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:114
		// _ = "end of CoverTab[23040]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:115
		_go_fuzz_dep_.CoverTab[23044]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:115
		// _ = "end of CoverTab[23044]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:115
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:115
	// _ = "end of CoverTab[23019]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:115
	_go_fuzz_dep_.CoverTab[23020]++
								if m.ocspStapling {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:116
		_go_fuzz_dep_.CoverTab[23045]++

									exts.AddUint16(extensionStatusRequest)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:119
			_go_fuzz_dep_.CoverTab[23046]++
										exts.AddUint8(1)
										exts.AddUint16(0)
										exts.AddUint16(0)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:122
			// _ = "end of CoverTab[23046]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:123
		// _ = "end of CoverTab[23045]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:124
		_go_fuzz_dep_.CoverTab[23047]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:124
		// _ = "end of CoverTab[23047]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:124
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:124
	// _ = "end of CoverTab[23020]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:124
	_go_fuzz_dep_.CoverTab[23021]++
								if len(m.supportedCurves) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:125
		_go_fuzz_dep_.CoverTab[23048]++

									exts.AddUint16(extensionSupportedCurves)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:128
			_go_fuzz_dep_.CoverTab[23049]++
										exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:129
				_go_fuzz_dep_.CoverTab[23050]++
											for _, curve := range m.supportedCurves {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:130
					_go_fuzz_dep_.CoverTab[23051]++
												exts.AddUint16(uint16(curve))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:131
					// _ = "end of CoverTab[23051]"
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:132
				// _ = "end of CoverTab[23050]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:133
			// _ = "end of CoverTab[23049]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:134
		// _ = "end of CoverTab[23048]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:135
		_go_fuzz_dep_.CoverTab[23052]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:135
		// _ = "end of CoverTab[23052]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:135
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:135
	// _ = "end of CoverTab[23021]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:135
	_go_fuzz_dep_.CoverTab[23022]++
								if len(m.supportedPoints) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:136
		_go_fuzz_dep_.CoverTab[23053]++

									exts.AddUint16(extensionSupportedPoints)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:139
			_go_fuzz_dep_.CoverTab[23054]++
										exts.AddUint8LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:140
				_go_fuzz_dep_.CoverTab[23055]++
											exts.AddBytes(m.supportedPoints)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:141
				// _ = "end of CoverTab[23055]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:142
			// _ = "end of CoverTab[23054]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:143
		// _ = "end of CoverTab[23053]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:144
		_go_fuzz_dep_.CoverTab[23056]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:144
		// _ = "end of CoverTab[23056]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:144
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:144
	// _ = "end of CoverTab[23022]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:144
	_go_fuzz_dep_.CoverTab[23023]++
								if m.ticketSupported {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:145
		_go_fuzz_dep_.CoverTab[23057]++

									exts.AddUint16(extensionSessionTicket)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:148
			_go_fuzz_dep_.CoverTab[23058]++
										exts.AddBytes(m.sessionTicket)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:149
			// _ = "end of CoverTab[23058]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:150
		// _ = "end of CoverTab[23057]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:151
		_go_fuzz_dep_.CoverTab[23059]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:151
		// _ = "end of CoverTab[23059]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:151
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:151
	// _ = "end of CoverTab[23023]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:151
	_go_fuzz_dep_.CoverTab[23024]++
								if len(m.supportedSignatureAlgorithms) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:152
		_go_fuzz_dep_.CoverTab[23060]++

									exts.AddUint16(extensionSignatureAlgorithms)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:155
			_go_fuzz_dep_.CoverTab[23061]++
										exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:156
				_go_fuzz_dep_.CoverTab[23062]++
											for _, sigAlgo := range m.supportedSignatureAlgorithms {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:157
					_go_fuzz_dep_.CoverTab[23063]++
												exts.AddUint16(uint16(sigAlgo))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:158
					// _ = "end of CoverTab[23063]"
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:159
				// _ = "end of CoverTab[23062]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:160
			// _ = "end of CoverTab[23061]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:161
		// _ = "end of CoverTab[23060]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:162
		_go_fuzz_dep_.CoverTab[23064]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:162
		// _ = "end of CoverTab[23064]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:162
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:162
	// _ = "end of CoverTab[23024]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:162
	_go_fuzz_dep_.CoverTab[23025]++
								if len(m.supportedSignatureAlgorithmsCert) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:163
		_go_fuzz_dep_.CoverTab[23065]++

									exts.AddUint16(extensionSignatureAlgorithmsCert)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:166
			_go_fuzz_dep_.CoverTab[23066]++
										exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:167
				_go_fuzz_dep_.CoverTab[23067]++
											for _, sigAlgo := range m.supportedSignatureAlgorithmsCert {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:168
					_go_fuzz_dep_.CoverTab[23068]++
												exts.AddUint16(uint16(sigAlgo))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:169
					// _ = "end of CoverTab[23068]"
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:170
				// _ = "end of CoverTab[23067]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:171
			// _ = "end of CoverTab[23066]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:172
		// _ = "end of CoverTab[23065]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:173
		_go_fuzz_dep_.CoverTab[23069]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:173
		// _ = "end of CoverTab[23069]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:173
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:173
	// _ = "end of CoverTab[23025]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:173
	_go_fuzz_dep_.CoverTab[23026]++
								if m.secureRenegotiationSupported {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:174
		_go_fuzz_dep_.CoverTab[23070]++

									exts.AddUint16(extensionRenegotiationInfo)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:177
			_go_fuzz_dep_.CoverTab[23071]++
										exts.AddUint8LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:178
				_go_fuzz_dep_.CoverTab[23072]++
											exts.AddBytes(m.secureRenegotiation)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:179
				// _ = "end of CoverTab[23072]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:180
			// _ = "end of CoverTab[23071]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:181
		// _ = "end of CoverTab[23070]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:182
		_go_fuzz_dep_.CoverTab[23073]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:182
		// _ = "end of CoverTab[23073]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:182
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:182
	// _ = "end of CoverTab[23026]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:182
	_go_fuzz_dep_.CoverTab[23027]++
								if len(m.alpnProtocols) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:183
		_go_fuzz_dep_.CoverTab[23074]++

									exts.AddUint16(extensionALPN)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:186
			_go_fuzz_dep_.CoverTab[23075]++
										exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:187
				_go_fuzz_dep_.CoverTab[23076]++
											for _, proto := range m.alpnProtocols {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:188
					_go_fuzz_dep_.CoverTab[23077]++
												exts.AddUint8LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:189
						_go_fuzz_dep_.CoverTab[23078]++
													exts.AddBytes([]byte(proto))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:190
						// _ = "end of CoverTab[23078]"
					})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:191
					// _ = "end of CoverTab[23077]"
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:192
				// _ = "end of CoverTab[23076]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:193
			// _ = "end of CoverTab[23075]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:194
		// _ = "end of CoverTab[23074]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:195
		_go_fuzz_dep_.CoverTab[23079]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:195
		// _ = "end of CoverTab[23079]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:195
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:195
	// _ = "end of CoverTab[23027]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:195
	_go_fuzz_dep_.CoverTab[23028]++
								if m.scts {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:196
		_go_fuzz_dep_.CoverTab[23080]++

									exts.AddUint16(extensionSCT)
									exts.AddUint16(0)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:199
		// _ = "end of CoverTab[23080]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:200
		_go_fuzz_dep_.CoverTab[23081]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:200
		// _ = "end of CoverTab[23081]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:200
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:200
	// _ = "end of CoverTab[23028]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:200
	_go_fuzz_dep_.CoverTab[23029]++
								if len(m.supportedVersions) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:201
		_go_fuzz_dep_.CoverTab[23082]++

									exts.AddUint16(extensionSupportedVersions)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:204
			_go_fuzz_dep_.CoverTab[23083]++
										exts.AddUint8LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:205
				_go_fuzz_dep_.CoverTab[23084]++
											for _, vers := range m.supportedVersions {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:206
					_go_fuzz_dep_.CoverTab[23085]++
												exts.AddUint16(vers)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:207
					// _ = "end of CoverTab[23085]"
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:208
				// _ = "end of CoverTab[23084]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:209
			// _ = "end of CoverTab[23083]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:210
		// _ = "end of CoverTab[23082]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:211
		_go_fuzz_dep_.CoverTab[23086]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:211
		// _ = "end of CoverTab[23086]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:211
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:211
	// _ = "end of CoverTab[23029]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:211
	_go_fuzz_dep_.CoverTab[23030]++
								if len(m.cookie) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:212
		_go_fuzz_dep_.CoverTab[23087]++

									exts.AddUint16(extensionCookie)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:215
			_go_fuzz_dep_.CoverTab[23088]++
										exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:216
				_go_fuzz_dep_.CoverTab[23089]++
											exts.AddBytes(m.cookie)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:217
				// _ = "end of CoverTab[23089]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:218
			// _ = "end of CoverTab[23088]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:219
		// _ = "end of CoverTab[23087]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:220
		_go_fuzz_dep_.CoverTab[23090]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:220
		// _ = "end of CoverTab[23090]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:220
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:220
	// _ = "end of CoverTab[23030]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:220
	_go_fuzz_dep_.CoverTab[23031]++
								if len(m.keyShares) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:221
		_go_fuzz_dep_.CoverTab[23091]++

									exts.AddUint16(extensionKeyShare)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:224
			_go_fuzz_dep_.CoverTab[23092]++
										exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:225
				_go_fuzz_dep_.CoverTab[23093]++
											for _, ks := range m.keyShares {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:226
					_go_fuzz_dep_.CoverTab[23094]++
												exts.AddUint16(uint16(ks.group))
												exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:228
						_go_fuzz_dep_.CoverTab[23095]++
													exts.AddBytes(ks.data)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:229
						// _ = "end of CoverTab[23095]"
					})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:230
					// _ = "end of CoverTab[23094]"
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:231
				// _ = "end of CoverTab[23093]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:232
			// _ = "end of CoverTab[23092]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:233
		// _ = "end of CoverTab[23091]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:234
		_go_fuzz_dep_.CoverTab[23096]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:234
		// _ = "end of CoverTab[23096]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:234
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:234
	// _ = "end of CoverTab[23031]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:234
	_go_fuzz_dep_.CoverTab[23032]++
								if m.earlyData {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:235
		_go_fuzz_dep_.CoverTab[23097]++

									exts.AddUint16(extensionEarlyData)
									exts.AddUint16(0)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:238
		// _ = "end of CoverTab[23097]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:239
		_go_fuzz_dep_.CoverTab[23098]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:239
		// _ = "end of CoverTab[23098]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:239
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:239
	// _ = "end of CoverTab[23032]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:239
	_go_fuzz_dep_.CoverTab[23033]++
								if len(m.pskModes) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:240
		_go_fuzz_dep_.CoverTab[23099]++

									exts.AddUint16(extensionPSKModes)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:243
			_go_fuzz_dep_.CoverTab[23100]++
										exts.AddUint8LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:244
				_go_fuzz_dep_.CoverTab[23101]++
											exts.AddBytes(m.pskModes)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:245
				// _ = "end of CoverTab[23101]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:246
			// _ = "end of CoverTab[23100]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:247
		// _ = "end of CoverTab[23099]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:248
		_go_fuzz_dep_.CoverTab[23102]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:248
		// _ = "end of CoverTab[23102]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:248
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:248
	// _ = "end of CoverTab[23033]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:248
	_go_fuzz_dep_.CoverTab[23034]++
								if len(m.pskIdentities) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:249
		_go_fuzz_dep_.CoverTab[23103]++

									exts.AddUint16(extensionPreSharedKey)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:252
			_go_fuzz_dep_.CoverTab[23104]++
										exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:253
				_go_fuzz_dep_.CoverTab[23106]++
											for _, psk := range m.pskIdentities {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:254
					_go_fuzz_dep_.CoverTab[23107]++
												exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:255
						_go_fuzz_dep_.CoverTab[23109]++
													exts.AddBytes(psk.label)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:256
						// _ = "end of CoverTab[23109]"
					})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:257
					// _ = "end of CoverTab[23107]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:257
					_go_fuzz_dep_.CoverTab[23108]++
												exts.AddUint32(psk.obfuscatedTicketAge)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:258
					// _ = "end of CoverTab[23108]"
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:259
				// _ = "end of CoverTab[23106]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:260
			// _ = "end of CoverTab[23104]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:260
			_go_fuzz_dep_.CoverTab[23105]++
										exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:261
				_go_fuzz_dep_.CoverTab[23110]++
											for _, binder := range m.pskBinders {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:262
					_go_fuzz_dep_.CoverTab[23111]++
												exts.AddUint8LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:263
						_go_fuzz_dep_.CoverTab[23112]++
													exts.AddBytes(binder)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:264
						// _ = "end of CoverTab[23112]"
					})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:265
					// _ = "end of CoverTab[23111]"
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:266
				// _ = "end of CoverTab[23110]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:267
			// _ = "end of CoverTab[23105]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:268
		// _ = "end of CoverTab[23103]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:269
		_go_fuzz_dep_.CoverTab[23113]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:269
		// _ = "end of CoverTab[23113]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:269
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:269
	// _ = "end of CoverTab[23034]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:269
	_go_fuzz_dep_.CoverTab[23035]++
								extBytes, err := exts.Bytes()
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:271
		_go_fuzz_dep_.CoverTab[23114]++
									return nil, err
//line /usr/local/go/src/crypto/tls/handshake_messages.go:272
		// _ = "end of CoverTab[23114]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:273
		_go_fuzz_dep_.CoverTab[23115]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:273
		// _ = "end of CoverTab[23115]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:273
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:273
	// _ = "end of CoverTab[23035]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:273
	_go_fuzz_dep_.CoverTab[23036]++

								var b cryptobyte.Builder
								b.AddUint8(typeClientHello)
								b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:277
		_go_fuzz_dep_.CoverTab[23116]++
									b.AddUint16(m.vers)
									addBytesWithLength(b, m.random, 32)
									b.AddUint8LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:280
			_go_fuzz_dep_.CoverTab[23120]++
										b.AddBytes(m.sessionId)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:281
			// _ = "end of CoverTab[23120]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:282
		// _ = "end of CoverTab[23116]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:282
		_go_fuzz_dep_.CoverTab[23117]++
									b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:283
			_go_fuzz_dep_.CoverTab[23121]++
										for _, suite := range m.cipherSuites {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:284
				_go_fuzz_dep_.CoverTab[23122]++
											b.AddUint16(suite)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:285
				// _ = "end of CoverTab[23122]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:286
			// _ = "end of CoverTab[23121]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:287
		// _ = "end of CoverTab[23117]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:287
		_go_fuzz_dep_.CoverTab[23118]++
									b.AddUint8LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:288
			_go_fuzz_dep_.CoverTab[23123]++
										b.AddBytes(m.compressionMethods)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:289
			// _ = "end of CoverTab[23123]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:290
		// _ = "end of CoverTab[23118]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:290
		_go_fuzz_dep_.CoverTab[23119]++

									if len(extBytes) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:292
			_go_fuzz_dep_.CoverTab[23124]++
										b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:293
				_go_fuzz_dep_.CoverTab[23125]++
											b.AddBytes(extBytes)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:294
				// _ = "end of CoverTab[23125]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:295
			// _ = "end of CoverTab[23124]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:296
			_go_fuzz_dep_.CoverTab[23126]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:296
			// _ = "end of CoverTab[23126]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:296
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:296
		// _ = "end of CoverTab[23119]"
	})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:297
	// _ = "end of CoverTab[23036]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:297
	_go_fuzz_dep_.CoverTab[23037]++

								m.raw, err = b.Bytes()
								return m.raw, err
//line /usr/local/go/src/crypto/tls/handshake_messages.go:300
	// _ = "end of CoverTab[23037]"
}

// marshalWithoutBinders returns the ClientHello through the
//line /usr/local/go/src/crypto/tls/handshake_messages.go:303
// PreSharedKeyExtension.identities field, according to RFC 8446, Section
//line /usr/local/go/src/crypto/tls/handshake_messages.go:303
// 4.2.11.2. Note that m.pskBinders must be set to slices of the correct length.
//line /usr/local/go/src/crypto/tls/handshake_messages.go:306
func (m *clientHelloMsg) marshalWithoutBinders() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:306
	_go_fuzz_dep_.CoverTab[23127]++
								bindersLen := 2
								for _, binder := range m.pskBinders {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:308
		_go_fuzz_dep_.CoverTab[23130]++
									bindersLen += 1
									bindersLen += len(binder)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:310
		// _ = "end of CoverTab[23130]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:311
	// _ = "end of CoverTab[23127]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:311
	_go_fuzz_dep_.CoverTab[23128]++

								fullMessage, err := m.marshal()
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:314
		_go_fuzz_dep_.CoverTab[23131]++
									return nil, err
//line /usr/local/go/src/crypto/tls/handshake_messages.go:315
		// _ = "end of CoverTab[23131]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:316
		_go_fuzz_dep_.CoverTab[23132]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:316
		// _ = "end of CoverTab[23132]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:316
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:316
	// _ = "end of CoverTab[23128]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:316
	_go_fuzz_dep_.CoverTab[23129]++
								return fullMessage[:len(fullMessage)-bindersLen], nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:317
	// _ = "end of CoverTab[23129]"
}

// updateBinders updates the m.pskBinders field, if necessary updating the
//line /usr/local/go/src/crypto/tls/handshake_messages.go:320
// cached marshaled representation. The supplied binders must have the same
//line /usr/local/go/src/crypto/tls/handshake_messages.go:320
// length as the current m.pskBinders.
//line /usr/local/go/src/crypto/tls/handshake_messages.go:323
func (m *clientHelloMsg) updateBinders(pskBinders [][]byte) error {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:323
	_go_fuzz_dep_.CoverTab[23133]++
								if len(pskBinders) != len(m.pskBinders) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:324
		_go_fuzz_dep_.CoverTab[23137]++
									return errors.New("tls: internal error: pskBinders length mismatch")
//line /usr/local/go/src/crypto/tls/handshake_messages.go:325
		// _ = "end of CoverTab[23137]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:326
		_go_fuzz_dep_.CoverTab[23138]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:326
		// _ = "end of CoverTab[23138]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:326
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:326
	// _ = "end of CoverTab[23133]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:326
	_go_fuzz_dep_.CoverTab[23134]++
								for i := range m.pskBinders {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:327
		_go_fuzz_dep_.CoverTab[23139]++
									if len(pskBinders[i]) != len(m.pskBinders[i]) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:328
			_go_fuzz_dep_.CoverTab[23140]++
										return errors.New("tls: internal error: pskBinders length mismatch")
//line /usr/local/go/src/crypto/tls/handshake_messages.go:329
			// _ = "end of CoverTab[23140]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:330
			_go_fuzz_dep_.CoverTab[23141]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:330
			// _ = "end of CoverTab[23141]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:330
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:330
		// _ = "end of CoverTab[23139]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:331
	// _ = "end of CoverTab[23134]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:331
	_go_fuzz_dep_.CoverTab[23135]++
								m.pskBinders = pskBinders
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:333
		_go_fuzz_dep_.CoverTab[23142]++
									helloBytes, err := m.marshalWithoutBinders()
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:335
			_go_fuzz_dep_.CoverTab[23145]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_messages.go:336
			// _ = "end of CoverTab[23145]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:337
			_go_fuzz_dep_.CoverTab[23146]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:337
			// _ = "end of CoverTab[23146]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:337
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:337
		// _ = "end of CoverTab[23142]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:337
		_go_fuzz_dep_.CoverTab[23143]++
									lenWithoutBinders := len(helloBytes)
									b := cryptobyte.NewFixedBuilder(m.raw[:lenWithoutBinders])
									b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:340
			_go_fuzz_dep_.CoverTab[23147]++
										for _, binder := range m.pskBinders {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:341
				_go_fuzz_dep_.CoverTab[23148]++
											b.AddUint8LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:342
					_go_fuzz_dep_.CoverTab[23149]++
												b.AddBytes(binder)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:343
					// _ = "end of CoverTab[23149]"
				})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:344
				// _ = "end of CoverTab[23148]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:345
			// _ = "end of CoverTab[23147]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:346
		// _ = "end of CoverTab[23143]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:346
		_go_fuzz_dep_.CoverTab[23144]++
									if out, err := b.Bytes(); err != nil || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:347
			_go_fuzz_dep_.CoverTab[23150]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:347
			return len(out) != len(m.raw)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:347
			// _ = "end of CoverTab[23150]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:347
		}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:347
			_go_fuzz_dep_.CoverTab[23151]++
										return errors.New("tls: internal error: failed to update binders")
//line /usr/local/go/src/crypto/tls/handshake_messages.go:348
			// _ = "end of CoverTab[23151]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:349
			_go_fuzz_dep_.CoverTab[23152]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:349
			// _ = "end of CoverTab[23152]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:349
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:349
		// _ = "end of CoverTab[23144]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:350
		_go_fuzz_dep_.CoverTab[23153]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:350
		// _ = "end of CoverTab[23153]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:350
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:350
	// _ = "end of CoverTab[23135]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:350
	_go_fuzz_dep_.CoverTab[23136]++

								return nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:352
	// _ = "end of CoverTab[23136]"
}

func (m *clientHelloMsg) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:355
	_go_fuzz_dep_.CoverTab[23154]++
								*m = clientHelloMsg{raw: data}
								s := cryptobyte.String(data)

								if !s.Skip(4) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:359
		_go_fuzz_dep_.CoverTab[23162]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:359
		return !s.ReadUint16(&m.vers)
									// _ = "end of CoverTab[23162]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:360
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:360
		_go_fuzz_dep_.CoverTab[23163]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:360
		return !s.ReadBytes(&m.random, 32)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:360
		// _ = "end of CoverTab[23163]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:360
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:360
		_go_fuzz_dep_.CoverTab[23164]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:360
		return !readUint8LengthPrefixed(&s, &m.sessionId)
									// _ = "end of CoverTab[23164]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:361
	}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:361
		_go_fuzz_dep_.CoverTab[23165]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:362
		// _ = "end of CoverTab[23165]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:363
		_go_fuzz_dep_.CoverTab[23166]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:363
		// _ = "end of CoverTab[23166]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:363
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:363
	// _ = "end of CoverTab[23154]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:363
	_go_fuzz_dep_.CoverTab[23155]++

								var cipherSuites cryptobyte.String
								if !s.ReadUint16LengthPrefixed(&cipherSuites) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:366
		_go_fuzz_dep_.CoverTab[23167]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:367
		// _ = "end of CoverTab[23167]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:368
		_go_fuzz_dep_.CoverTab[23168]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:368
		// _ = "end of CoverTab[23168]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:368
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:368
	// _ = "end of CoverTab[23155]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:368
	_go_fuzz_dep_.CoverTab[23156]++
								m.cipherSuites = []uint16{}
								m.secureRenegotiationSupported = false
								for !cipherSuites.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:371
		_go_fuzz_dep_.CoverTab[23169]++
									var suite uint16
									if !cipherSuites.ReadUint16(&suite) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:373
			_go_fuzz_dep_.CoverTab[23172]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:374
			// _ = "end of CoverTab[23172]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:375
			_go_fuzz_dep_.CoverTab[23173]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:375
			// _ = "end of CoverTab[23173]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:375
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:375
		// _ = "end of CoverTab[23169]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:375
		_go_fuzz_dep_.CoverTab[23170]++
									if suite == scsvRenegotiation {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:376
			_go_fuzz_dep_.CoverTab[23174]++
										m.secureRenegotiationSupported = true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:377
			// _ = "end of CoverTab[23174]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:378
			_go_fuzz_dep_.CoverTab[23175]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:378
			// _ = "end of CoverTab[23175]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:378
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:378
		// _ = "end of CoverTab[23170]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:378
		_go_fuzz_dep_.CoverTab[23171]++
									m.cipherSuites = append(m.cipherSuites, suite)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:379
		// _ = "end of CoverTab[23171]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:380
	// _ = "end of CoverTab[23156]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:380
	_go_fuzz_dep_.CoverTab[23157]++

								if !readUint8LengthPrefixed(&s, &m.compressionMethods) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:382
		_go_fuzz_dep_.CoverTab[23176]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:383
		// _ = "end of CoverTab[23176]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:384
		_go_fuzz_dep_.CoverTab[23177]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:384
		// _ = "end of CoverTab[23177]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:384
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:384
	// _ = "end of CoverTab[23157]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:384
	_go_fuzz_dep_.CoverTab[23158]++

								if s.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:386
		_go_fuzz_dep_.CoverTab[23178]++

									return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:388
		// _ = "end of CoverTab[23178]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:389
		_go_fuzz_dep_.CoverTab[23179]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:389
		// _ = "end of CoverTab[23179]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:389
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:389
	// _ = "end of CoverTab[23158]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:389
	_go_fuzz_dep_.CoverTab[23159]++

								var extensions cryptobyte.String
								if !s.ReadUint16LengthPrefixed(&extensions) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:392
		_go_fuzz_dep_.CoverTab[23180]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:392
		return !s.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:392
		// _ = "end of CoverTab[23180]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:392
	}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:392
		_go_fuzz_dep_.CoverTab[23181]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:393
		// _ = "end of CoverTab[23181]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:394
		_go_fuzz_dep_.CoverTab[23182]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:394
		// _ = "end of CoverTab[23182]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:394
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:394
	// _ = "end of CoverTab[23159]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:394
	_go_fuzz_dep_.CoverTab[23160]++

								seenExts := make(map[uint16]bool)
								for !extensions.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:397
		_go_fuzz_dep_.CoverTab[23183]++
									var extension uint16
									var extData cryptobyte.String
									if !extensions.ReadUint16(&extension) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:400
			_go_fuzz_dep_.CoverTab[23187]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:400
			return !extensions.ReadUint16LengthPrefixed(&extData)
										// _ = "end of CoverTab[23187]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:401
		}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:401
			_go_fuzz_dep_.CoverTab[23188]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:402
			// _ = "end of CoverTab[23188]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:403
			_go_fuzz_dep_.CoverTab[23189]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:403
			// _ = "end of CoverTab[23189]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:403
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:403
		// _ = "end of CoverTab[23183]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:403
		_go_fuzz_dep_.CoverTab[23184]++

									if seenExts[extension] {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:405
			_go_fuzz_dep_.CoverTab[23190]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:406
			// _ = "end of CoverTab[23190]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:407
			_go_fuzz_dep_.CoverTab[23191]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:407
			// _ = "end of CoverTab[23191]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:407
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:407
		// _ = "end of CoverTab[23184]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:407
		_go_fuzz_dep_.CoverTab[23185]++
									seenExts[extension] = true

									switch extension {
		case extensionServerName:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:411
			_go_fuzz_dep_.CoverTab[23192]++
			// RFC 6066, Section 3
			var nameList cryptobyte.String
			if !extData.ReadUint16LengthPrefixed(&nameList) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:414
				_go_fuzz_dep_.CoverTab[23222]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:414
				return nameList.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:414
				// _ = "end of CoverTab[23222]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:414
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:414
				_go_fuzz_dep_.CoverTab[23223]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:415
				// _ = "end of CoverTab[23223]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:416
				_go_fuzz_dep_.CoverTab[23224]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:416
				// _ = "end of CoverTab[23224]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:416
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:416
			// _ = "end of CoverTab[23192]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:416
			_go_fuzz_dep_.CoverTab[23193]++
										for !nameList.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:417
				_go_fuzz_dep_.CoverTab[23225]++
											var nameType uint8
											var serverName cryptobyte.String
											if !nameList.ReadUint8(&nameType) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:420
					_go_fuzz_dep_.CoverTab[23229]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:420
					return !nameList.ReadUint16LengthPrefixed(&serverName)
												// _ = "end of CoverTab[23229]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:421
				}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:421
					_go_fuzz_dep_.CoverTab[23230]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:421
					return serverName.Empty()
												// _ = "end of CoverTab[23230]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:422
				}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:422
					_go_fuzz_dep_.CoverTab[23231]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:423
					// _ = "end of CoverTab[23231]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:424
					_go_fuzz_dep_.CoverTab[23232]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:424
					// _ = "end of CoverTab[23232]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:424
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:424
				// _ = "end of CoverTab[23225]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:424
				_go_fuzz_dep_.CoverTab[23226]++
											if nameType != 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:425
					_go_fuzz_dep_.CoverTab[23233]++
												continue
//line /usr/local/go/src/crypto/tls/handshake_messages.go:426
					// _ = "end of CoverTab[23233]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:427
					_go_fuzz_dep_.CoverTab[23234]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:427
					// _ = "end of CoverTab[23234]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:427
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:427
				// _ = "end of CoverTab[23226]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:427
				_go_fuzz_dep_.CoverTab[23227]++
											if len(m.serverName) != 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:428
					_go_fuzz_dep_.CoverTab[23235]++

												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:430
					// _ = "end of CoverTab[23235]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:431
					_go_fuzz_dep_.CoverTab[23236]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:431
					// _ = "end of CoverTab[23236]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:431
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:431
				// _ = "end of CoverTab[23227]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:431
				_go_fuzz_dep_.CoverTab[23228]++
											m.serverName = string(serverName)

											if strings.HasSuffix(m.serverName, ".") {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:434
					_go_fuzz_dep_.CoverTab[23237]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:435
					// _ = "end of CoverTab[23237]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:436
					_go_fuzz_dep_.CoverTab[23238]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:436
					// _ = "end of CoverTab[23238]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:436
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:436
				// _ = "end of CoverTab[23228]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:437
			// _ = "end of CoverTab[23193]"
		case extensionStatusRequest:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:438
			_go_fuzz_dep_.CoverTab[23194]++
			// RFC 4366, Section 3.6
			var statusType uint8
			var ignored cryptobyte.String
			if !extData.ReadUint8(&statusType) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:442
				_go_fuzz_dep_.CoverTab[23239]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:442
				return !extData.ReadUint16LengthPrefixed(&ignored)
											// _ = "end of CoverTab[23239]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:443
			}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:443
				_go_fuzz_dep_.CoverTab[23240]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:443
				return !extData.ReadUint16LengthPrefixed(&ignored)
											// _ = "end of CoverTab[23240]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:444
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:444
				_go_fuzz_dep_.CoverTab[23241]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:445
				// _ = "end of CoverTab[23241]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:446
				_go_fuzz_dep_.CoverTab[23242]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:446
				// _ = "end of CoverTab[23242]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:446
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:446
			// _ = "end of CoverTab[23194]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:446
			_go_fuzz_dep_.CoverTab[23195]++
										m.ocspStapling = statusType == statusTypeOCSP
//line /usr/local/go/src/crypto/tls/handshake_messages.go:447
			// _ = "end of CoverTab[23195]"
		case extensionSupportedCurves:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:448
			_go_fuzz_dep_.CoverTab[23196]++
			// RFC 4492, sections 5.1.1 and RFC 8446, Section 4.2.7
			var curves cryptobyte.String
			if !extData.ReadUint16LengthPrefixed(&curves) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:451
				_go_fuzz_dep_.CoverTab[23243]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:451
				return curves.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:451
				// _ = "end of CoverTab[23243]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:451
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:451
				_go_fuzz_dep_.CoverTab[23244]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:452
				// _ = "end of CoverTab[23244]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:453
				_go_fuzz_dep_.CoverTab[23245]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:453
				// _ = "end of CoverTab[23245]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:453
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:453
			// _ = "end of CoverTab[23196]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:453
			_go_fuzz_dep_.CoverTab[23197]++
										for !curves.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:454
				_go_fuzz_dep_.CoverTab[23246]++
											var curve uint16
											if !curves.ReadUint16(&curve) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:456
					_go_fuzz_dep_.CoverTab[23248]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:457
					// _ = "end of CoverTab[23248]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:458
					_go_fuzz_dep_.CoverTab[23249]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:458
					// _ = "end of CoverTab[23249]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:458
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:458
				// _ = "end of CoverTab[23246]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:458
				_go_fuzz_dep_.CoverTab[23247]++
											m.supportedCurves = append(m.supportedCurves, CurveID(curve))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:459
				// _ = "end of CoverTab[23247]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:460
			// _ = "end of CoverTab[23197]"
		case extensionSupportedPoints:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:461
			_go_fuzz_dep_.CoverTab[23198]++

										if !readUint8LengthPrefixed(&extData, &m.supportedPoints) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:463
				_go_fuzz_dep_.CoverTab[23250]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:463
				return len(m.supportedPoints) == 0
											// _ = "end of CoverTab[23250]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:464
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:464
				_go_fuzz_dep_.CoverTab[23251]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:465
				// _ = "end of CoverTab[23251]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:466
				_go_fuzz_dep_.CoverTab[23252]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:466
				// _ = "end of CoverTab[23252]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:466
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:466
			// _ = "end of CoverTab[23198]"
		case extensionSessionTicket:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:467
			_go_fuzz_dep_.CoverTab[23199]++

										m.ticketSupported = true
										extData.ReadBytes(&m.sessionTicket, len(extData))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:470
			// _ = "end of CoverTab[23199]"
		case extensionSignatureAlgorithms:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:471
			_go_fuzz_dep_.CoverTab[23200]++
			// RFC 5246, Section 7.4.1.4.1
			var sigAndAlgs cryptobyte.String
			if !extData.ReadUint16LengthPrefixed(&sigAndAlgs) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:474
				_go_fuzz_dep_.CoverTab[23253]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:474
				return sigAndAlgs.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:474
				// _ = "end of CoverTab[23253]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:474
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:474
				_go_fuzz_dep_.CoverTab[23254]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:475
				// _ = "end of CoverTab[23254]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:476
				_go_fuzz_dep_.CoverTab[23255]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:476
				// _ = "end of CoverTab[23255]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:476
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:476
			// _ = "end of CoverTab[23200]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:476
			_go_fuzz_dep_.CoverTab[23201]++
										for !sigAndAlgs.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:477
				_go_fuzz_dep_.CoverTab[23256]++
											var sigAndAlg uint16
											if !sigAndAlgs.ReadUint16(&sigAndAlg) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:479
					_go_fuzz_dep_.CoverTab[23258]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:480
					// _ = "end of CoverTab[23258]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:481
					_go_fuzz_dep_.CoverTab[23259]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:481
					// _ = "end of CoverTab[23259]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:481
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:481
				// _ = "end of CoverTab[23256]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:481
				_go_fuzz_dep_.CoverTab[23257]++
											m.supportedSignatureAlgorithms = append(
					m.supportedSignatureAlgorithms, SignatureScheme(sigAndAlg))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:483
				// _ = "end of CoverTab[23257]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:484
			// _ = "end of CoverTab[23201]"
		case extensionSignatureAlgorithmsCert:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:485
			_go_fuzz_dep_.CoverTab[23202]++
			// RFC 8446, Section 4.2.3
			var sigAndAlgs cryptobyte.String
			if !extData.ReadUint16LengthPrefixed(&sigAndAlgs) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:488
				_go_fuzz_dep_.CoverTab[23260]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:488
				return sigAndAlgs.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:488
				// _ = "end of CoverTab[23260]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:488
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:488
				_go_fuzz_dep_.CoverTab[23261]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:489
				// _ = "end of CoverTab[23261]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:490
				_go_fuzz_dep_.CoverTab[23262]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:490
				// _ = "end of CoverTab[23262]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:490
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:490
			// _ = "end of CoverTab[23202]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:490
			_go_fuzz_dep_.CoverTab[23203]++
										for !sigAndAlgs.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:491
				_go_fuzz_dep_.CoverTab[23263]++
											var sigAndAlg uint16
											if !sigAndAlgs.ReadUint16(&sigAndAlg) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:493
					_go_fuzz_dep_.CoverTab[23265]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:494
					// _ = "end of CoverTab[23265]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:495
					_go_fuzz_dep_.CoverTab[23266]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:495
					// _ = "end of CoverTab[23266]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:495
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:495
				// _ = "end of CoverTab[23263]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:495
				_go_fuzz_dep_.CoverTab[23264]++
											m.supportedSignatureAlgorithmsCert = append(
					m.supportedSignatureAlgorithmsCert, SignatureScheme(sigAndAlg))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:497
				// _ = "end of CoverTab[23264]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:498
			// _ = "end of CoverTab[23203]"
		case extensionRenegotiationInfo:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:499
			_go_fuzz_dep_.CoverTab[23204]++

										if !readUint8LengthPrefixed(&extData, &m.secureRenegotiation) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:501
				_go_fuzz_dep_.CoverTab[23267]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:502
				// _ = "end of CoverTab[23267]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:503
				_go_fuzz_dep_.CoverTab[23268]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:503
				// _ = "end of CoverTab[23268]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:503
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:503
			// _ = "end of CoverTab[23204]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:503
			_go_fuzz_dep_.CoverTab[23205]++
										m.secureRenegotiationSupported = true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:504
			// _ = "end of CoverTab[23205]"
		case extensionALPN:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:505
			_go_fuzz_dep_.CoverTab[23206]++
			// RFC 7301, Section 3.1
			var protoList cryptobyte.String
			if !extData.ReadUint16LengthPrefixed(&protoList) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:508
				_go_fuzz_dep_.CoverTab[23269]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:508
				return protoList.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:508
				// _ = "end of CoverTab[23269]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:508
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:508
				_go_fuzz_dep_.CoverTab[23270]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:509
				// _ = "end of CoverTab[23270]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:510
				_go_fuzz_dep_.CoverTab[23271]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:510
				// _ = "end of CoverTab[23271]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:510
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:510
			// _ = "end of CoverTab[23206]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:510
			_go_fuzz_dep_.CoverTab[23207]++
										for !protoList.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:511
				_go_fuzz_dep_.CoverTab[23272]++
											var proto cryptobyte.String
											if !protoList.ReadUint8LengthPrefixed(&proto) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:513
					_go_fuzz_dep_.CoverTab[23274]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:513
					return proto.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:513
					// _ = "end of CoverTab[23274]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:513
				}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:513
					_go_fuzz_dep_.CoverTab[23275]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:514
					// _ = "end of CoverTab[23275]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:515
					_go_fuzz_dep_.CoverTab[23276]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:515
					// _ = "end of CoverTab[23276]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:515
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:515
				// _ = "end of CoverTab[23272]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:515
				_go_fuzz_dep_.CoverTab[23273]++
											m.alpnProtocols = append(m.alpnProtocols, string(proto))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:516
				// _ = "end of CoverTab[23273]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:517
			// _ = "end of CoverTab[23207]"
		case extensionSCT:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:518
			_go_fuzz_dep_.CoverTab[23208]++

										m.scts = true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:520
			// _ = "end of CoverTab[23208]"
		case extensionSupportedVersions:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:521
			_go_fuzz_dep_.CoverTab[23209]++
			// RFC 8446, Section 4.2.1
			var versList cryptobyte.String
			if !extData.ReadUint8LengthPrefixed(&versList) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:524
				_go_fuzz_dep_.CoverTab[23277]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:524
				return versList.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:524
				// _ = "end of CoverTab[23277]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:524
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:524
				_go_fuzz_dep_.CoverTab[23278]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:525
				// _ = "end of CoverTab[23278]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:526
				_go_fuzz_dep_.CoverTab[23279]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:526
				// _ = "end of CoverTab[23279]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:526
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:526
			// _ = "end of CoverTab[23209]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:526
			_go_fuzz_dep_.CoverTab[23210]++
										for !versList.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:527
				_go_fuzz_dep_.CoverTab[23280]++
											var vers uint16
											if !versList.ReadUint16(&vers) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:529
					_go_fuzz_dep_.CoverTab[23282]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:530
					// _ = "end of CoverTab[23282]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:531
					_go_fuzz_dep_.CoverTab[23283]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:531
					// _ = "end of CoverTab[23283]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:531
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:531
				// _ = "end of CoverTab[23280]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:531
				_go_fuzz_dep_.CoverTab[23281]++
											m.supportedVersions = append(m.supportedVersions, vers)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:532
				// _ = "end of CoverTab[23281]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:533
			// _ = "end of CoverTab[23210]"
		case extensionCookie:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:534
			_go_fuzz_dep_.CoverTab[23211]++

										if !readUint16LengthPrefixed(&extData, &m.cookie) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:536
				_go_fuzz_dep_.CoverTab[23284]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:536
				return len(m.cookie) == 0
											// _ = "end of CoverTab[23284]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:537
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:537
				_go_fuzz_dep_.CoverTab[23285]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:538
				// _ = "end of CoverTab[23285]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:539
				_go_fuzz_dep_.CoverTab[23286]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:539
				// _ = "end of CoverTab[23286]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:539
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:539
			// _ = "end of CoverTab[23211]"
		case extensionKeyShare:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:540
			_go_fuzz_dep_.CoverTab[23212]++
			// RFC 8446, Section 4.2.8
			var clientShares cryptobyte.String
			if !extData.ReadUint16LengthPrefixed(&clientShares) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:543
				_go_fuzz_dep_.CoverTab[23287]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:544
				// _ = "end of CoverTab[23287]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:545
				_go_fuzz_dep_.CoverTab[23288]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:545
				// _ = "end of CoverTab[23288]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:545
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:545
			// _ = "end of CoverTab[23212]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:545
			_go_fuzz_dep_.CoverTab[23213]++
										for !clientShares.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:546
				_go_fuzz_dep_.CoverTab[23289]++
											var ks keyShare
											if !clientShares.ReadUint16((*uint16)(&ks.group)) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:548
					_go_fuzz_dep_.CoverTab[23291]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:548
					return !readUint16LengthPrefixed(&clientShares, &ks.data)
												// _ = "end of CoverTab[23291]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:549
				}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:549
					_go_fuzz_dep_.CoverTab[23292]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:549
					return len(ks.data) == 0
												// _ = "end of CoverTab[23292]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:550
				}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:550
					_go_fuzz_dep_.CoverTab[23293]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:551
					// _ = "end of CoverTab[23293]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:552
					_go_fuzz_dep_.CoverTab[23294]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:552
					// _ = "end of CoverTab[23294]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:552
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:552
				// _ = "end of CoverTab[23289]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:552
				_go_fuzz_dep_.CoverTab[23290]++
											m.keyShares = append(m.keyShares, ks)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:553
				// _ = "end of CoverTab[23290]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:554
			// _ = "end of CoverTab[23213]"
		case extensionEarlyData:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:555
			_go_fuzz_dep_.CoverTab[23214]++

										m.earlyData = true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:557
			// _ = "end of CoverTab[23214]"
		case extensionPSKModes:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:558
			_go_fuzz_dep_.CoverTab[23215]++

										if !readUint8LengthPrefixed(&extData, &m.pskModes) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:560
				_go_fuzz_dep_.CoverTab[23295]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:561
				// _ = "end of CoverTab[23295]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:562
				_go_fuzz_dep_.CoverTab[23296]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:562
				// _ = "end of CoverTab[23296]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:562
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:562
			// _ = "end of CoverTab[23215]"
		case extensionPreSharedKey:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:563
			_go_fuzz_dep_.CoverTab[23216]++

										if !extensions.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:565
				_go_fuzz_dep_.CoverTab[23297]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:566
				// _ = "end of CoverTab[23297]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:567
				_go_fuzz_dep_.CoverTab[23298]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:567
				// _ = "end of CoverTab[23298]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:567
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:567
			// _ = "end of CoverTab[23216]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:567
			_go_fuzz_dep_.CoverTab[23217]++
										var identities cryptobyte.String
										if !extData.ReadUint16LengthPrefixed(&identities) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:569
				_go_fuzz_dep_.CoverTab[23299]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:569
				return identities.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:569
				// _ = "end of CoverTab[23299]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:569
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:569
				_go_fuzz_dep_.CoverTab[23300]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:570
				// _ = "end of CoverTab[23300]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:571
				_go_fuzz_dep_.CoverTab[23301]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:571
				// _ = "end of CoverTab[23301]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:571
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:571
			// _ = "end of CoverTab[23217]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:571
			_go_fuzz_dep_.CoverTab[23218]++
										for !identities.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:572
				_go_fuzz_dep_.CoverTab[23302]++
											var psk pskIdentity
											if !readUint16LengthPrefixed(&identities, &psk.label) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:574
					_go_fuzz_dep_.CoverTab[23304]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:574
					return !identities.ReadUint32(&psk.obfuscatedTicketAge)
												// _ = "end of CoverTab[23304]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:575
				}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:575
					_go_fuzz_dep_.CoverTab[23305]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:575
					return len(psk.label) == 0
												// _ = "end of CoverTab[23305]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:576
				}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:576
					_go_fuzz_dep_.CoverTab[23306]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:577
					// _ = "end of CoverTab[23306]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:578
					_go_fuzz_dep_.CoverTab[23307]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:578
					// _ = "end of CoverTab[23307]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:578
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:578
				// _ = "end of CoverTab[23302]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:578
				_go_fuzz_dep_.CoverTab[23303]++
											m.pskIdentities = append(m.pskIdentities, psk)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:579
				// _ = "end of CoverTab[23303]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:580
			// _ = "end of CoverTab[23218]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:580
			_go_fuzz_dep_.CoverTab[23219]++
										var binders cryptobyte.String
										if !extData.ReadUint16LengthPrefixed(&binders) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:582
				_go_fuzz_dep_.CoverTab[23308]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:582
				return binders.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:582
				// _ = "end of CoverTab[23308]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:582
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:582
				_go_fuzz_dep_.CoverTab[23309]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:583
				// _ = "end of CoverTab[23309]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:584
				_go_fuzz_dep_.CoverTab[23310]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:584
				// _ = "end of CoverTab[23310]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:584
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:584
			// _ = "end of CoverTab[23219]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:584
			_go_fuzz_dep_.CoverTab[23220]++
										for !binders.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:585
				_go_fuzz_dep_.CoverTab[23311]++
											var binder []byte
											if !readUint8LengthPrefixed(&binders, &binder) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:587
					_go_fuzz_dep_.CoverTab[23313]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:587
					return len(binder) == 0
												// _ = "end of CoverTab[23313]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:588
				}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:588
					_go_fuzz_dep_.CoverTab[23314]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:589
					// _ = "end of CoverTab[23314]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:590
					_go_fuzz_dep_.CoverTab[23315]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:590
					// _ = "end of CoverTab[23315]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:590
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:590
				// _ = "end of CoverTab[23311]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:590
				_go_fuzz_dep_.CoverTab[23312]++
											m.pskBinders = append(m.pskBinders, binder)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:591
				// _ = "end of CoverTab[23312]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:592
			// _ = "end of CoverTab[23220]"
		default:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:593
			_go_fuzz_dep_.CoverTab[23221]++

										continue
//line /usr/local/go/src/crypto/tls/handshake_messages.go:595
			// _ = "end of CoverTab[23221]"
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:596
		// _ = "end of CoverTab[23185]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:596
		_go_fuzz_dep_.CoverTab[23186]++

									if !extData.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:598
			_go_fuzz_dep_.CoverTab[23316]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:599
			// _ = "end of CoverTab[23316]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:600
			_go_fuzz_dep_.CoverTab[23317]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:600
			// _ = "end of CoverTab[23317]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:600
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:600
		// _ = "end of CoverTab[23186]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:601
	// _ = "end of CoverTab[23160]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:601
	_go_fuzz_dep_.CoverTab[23161]++

								return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:603
	// _ = "end of CoverTab[23161]"
}

type serverHelloMsg struct {
	raw				[]byte
	vers				uint16
	random				[]byte
	sessionId			[]byte
	cipherSuite			uint16
	compressionMethod		uint8
	ocspStapling			bool
	ticketSupported			bool
	secureRenegotiationSupported	bool
	secureRenegotiation		[]byte
	alpnProtocol			string
	scts				[][]byte
	supportedVersion		uint16
	serverShare			keyShare
	selectedIdentityPresent		bool
	selectedIdentity		uint16
	supportedPoints			[]uint8

	// HelloRetryRequest extensions
	cookie		[]byte
	selectedGroup	CurveID
}

func (m *serverHelloMsg) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:630
	_go_fuzz_dep_.CoverTab[23318]++
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:631
		_go_fuzz_dep_.CoverTab[23333]++
									return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:632
		// _ = "end of CoverTab[23333]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:633
		_go_fuzz_dep_.CoverTab[23334]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:633
		// _ = "end of CoverTab[23334]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:633
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:633
	// _ = "end of CoverTab[23318]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:633
	_go_fuzz_dep_.CoverTab[23319]++

								var exts cryptobyte.Builder
								if m.ocspStapling {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:636
		_go_fuzz_dep_.CoverTab[23335]++
									exts.AddUint16(extensionStatusRequest)
									exts.AddUint16(0)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:638
		// _ = "end of CoverTab[23335]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:639
		_go_fuzz_dep_.CoverTab[23336]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:639
		// _ = "end of CoverTab[23336]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:639
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:639
	// _ = "end of CoverTab[23319]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:639
	_go_fuzz_dep_.CoverTab[23320]++
								if m.ticketSupported {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:640
		_go_fuzz_dep_.CoverTab[23337]++
									exts.AddUint16(extensionSessionTicket)
									exts.AddUint16(0)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:642
		// _ = "end of CoverTab[23337]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:643
		_go_fuzz_dep_.CoverTab[23338]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:643
		// _ = "end of CoverTab[23338]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:643
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:643
	// _ = "end of CoverTab[23320]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:643
	_go_fuzz_dep_.CoverTab[23321]++
								if m.secureRenegotiationSupported {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:644
		_go_fuzz_dep_.CoverTab[23339]++
									exts.AddUint16(extensionRenegotiationInfo)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:646
			_go_fuzz_dep_.CoverTab[23340]++
										exts.AddUint8LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:647
				_go_fuzz_dep_.CoverTab[23341]++
											exts.AddBytes(m.secureRenegotiation)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:648
				// _ = "end of CoverTab[23341]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:649
			// _ = "end of CoverTab[23340]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:650
		// _ = "end of CoverTab[23339]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:651
		_go_fuzz_dep_.CoverTab[23342]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:651
		// _ = "end of CoverTab[23342]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:651
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:651
	// _ = "end of CoverTab[23321]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:651
	_go_fuzz_dep_.CoverTab[23322]++
								if len(m.alpnProtocol) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:652
		_go_fuzz_dep_.CoverTab[23343]++
									exts.AddUint16(extensionALPN)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:654
			_go_fuzz_dep_.CoverTab[23344]++
										exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:655
				_go_fuzz_dep_.CoverTab[23345]++
											exts.AddUint8LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:656
					_go_fuzz_dep_.CoverTab[23346]++
												exts.AddBytes([]byte(m.alpnProtocol))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:657
					// _ = "end of CoverTab[23346]"
				})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:658
				// _ = "end of CoverTab[23345]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:659
			// _ = "end of CoverTab[23344]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:660
		// _ = "end of CoverTab[23343]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:661
		_go_fuzz_dep_.CoverTab[23347]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:661
		// _ = "end of CoverTab[23347]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:661
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:661
	// _ = "end of CoverTab[23322]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:661
	_go_fuzz_dep_.CoverTab[23323]++
								if len(m.scts) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:662
		_go_fuzz_dep_.CoverTab[23348]++
									exts.AddUint16(extensionSCT)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:664
			_go_fuzz_dep_.CoverTab[23349]++
										exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:665
				_go_fuzz_dep_.CoverTab[23350]++
											for _, sct := range m.scts {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:666
					_go_fuzz_dep_.CoverTab[23351]++
												exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:667
						_go_fuzz_dep_.CoverTab[23352]++
													exts.AddBytes(sct)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:668
						// _ = "end of CoverTab[23352]"
					})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:669
					// _ = "end of CoverTab[23351]"
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:670
				// _ = "end of CoverTab[23350]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:671
			// _ = "end of CoverTab[23349]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:672
		// _ = "end of CoverTab[23348]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:673
		_go_fuzz_dep_.CoverTab[23353]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:673
		// _ = "end of CoverTab[23353]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:673
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:673
	// _ = "end of CoverTab[23323]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:673
	_go_fuzz_dep_.CoverTab[23324]++
								if m.supportedVersion != 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:674
		_go_fuzz_dep_.CoverTab[23354]++
									exts.AddUint16(extensionSupportedVersions)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:676
			_go_fuzz_dep_.CoverTab[23355]++
										exts.AddUint16(m.supportedVersion)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:677
			// _ = "end of CoverTab[23355]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:678
		// _ = "end of CoverTab[23354]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:679
		_go_fuzz_dep_.CoverTab[23356]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:679
		// _ = "end of CoverTab[23356]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:679
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:679
	// _ = "end of CoverTab[23324]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:679
	_go_fuzz_dep_.CoverTab[23325]++
								if m.serverShare.group != 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:680
		_go_fuzz_dep_.CoverTab[23357]++
									exts.AddUint16(extensionKeyShare)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:682
			_go_fuzz_dep_.CoverTab[23358]++
										exts.AddUint16(uint16(m.serverShare.group))
										exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:684
				_go_fuzz_dep_.CoverTab[23359]++
											exts.AddBytes(m.serverShare.data)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:685
				// _ = "end of CoverTab[23359]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:686
			// _ = "end of CoverTab[23358]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:687
		// _ = "end of CoverTab[23357]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:688
		_go_fuzz_dep_.CoverTab[23360]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:688
		// _ = "end of CoverTab[23360]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:688
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:688
	// _ = "end of CoverTab[23325]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:688
	_go_fuzz_dep_.CoverTab[23326]++
								if m.selectedIdentityPresent {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:689
		_go_fuzz_dep_.CoverTab[23361]++
									exts.AddUint16(extensionPreSharedKey)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:691
			_go_fuzz_dep_.CoverTab[23362]++
										exts.AddUint16(m.selectedIdentity)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:692
			// _ = "end of CoverTab[23362]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:693
		// _ = "end of CoverTab[23361]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:694
		_go_fuzz_dep_.CoverTab[23363]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:694
		// _ = "end of CoverTab[23363]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:694
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:694
	// _ = "end of CoverTab[23326]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:694
	_go_fuzz_dep_.CoverTab[23327]++

								if len(m.cookie) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:696
		_go_fuzz_dep_.CoverTab[23364]++
									exts.AddUint16(extensionCookie)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:698
			_go_fuzz_dep_.CoverTab[23365]++
										exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:699
				_go_fuzz_dep_.CoverTab[23366]++
											exts.AddBytes(m.cookie)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:700
				// _ = "end of CoverTab[23366]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:701
			// _ = "end of CoverTab[23365]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:702
		// _ = "end of CoverTab[23364]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:703
		_go_fuzz_dep_.CoverTab[23367]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:703
		// _ = "end of CoverTab[23367]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:703
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:703
	// _ = "end of CoverTab[23327]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:703
	_go_fuzz_dep_.CoverTab[23328]++
								if m.selectedGroup != 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:704
		_go_fuzz_dep_.CoverTab[23368]++
									exts.AddUint16(extensionKeyShare)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:706
			_go_fuzz_dep_.CoverTab[23369]++
										exts.AddUint16(uint16(m.selectedGroup))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:707
			// _ = "end of CoverTab[23369]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:708
		// _ = "end of CoverTab[23368]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:709
		_go_fuzz_dep_.CoverTab[23370]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:709
		// _ = "end of CoverTab[23370]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:709
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:709
	// _ = "end of CoverTab[23328]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:709
	_go_fuzz_dep_.CoverTab[23329]++
								if len(m.supportedPoints) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:710
		_go_fuzz_dep_.CoverTab[23371]++
									exts.AddUint16(extensionSupportedPoints)
									exts.AddUint16LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:712
			_go_fuzz_dep_.CoverTab[23372]++
										exts.AddUint8LengthPrefixed(func(exts *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:713
				_go_fuzz_dep_.CoverTab[23373]++
											exts.AddBytes(m.supportedPoints)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:714
				// _ = "end of CoverTab[23373]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:715
			// _ = "end of CoverTab[23372]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:716
		// _ = "end of CoverTab[23371]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:717
		_go_fuzz_dep_.CoverTab[23374]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:717
		// _ = "end of CoverTab[23374]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:717
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:717
	// _ = "end of CoverTab[23329]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:717
	_go_fuzz_dep_.CoverTab[23330]++

								extBytes, err := exts.Bytes()
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:720
		_go_fuzz_dep_.CoverTab[23375]++
									return nil, err
//line /usr/local/go/src/crypto/tls/handshake_messages.go:721
		// _ = "end of CoverTab[23375]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:722
		_go_fuzz_dep_.CoverTab[23376]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:722
		// _ = "end of CoverTab[23376]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:722
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:722
	// _ = "end of CoverTab[23330]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:722
	_go_fuzz_dep_.CoverTab[23331]++

								var b cryptobyte.Builder
								b.AddUint8(typeServerHello)
								b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:726
		_go_fuzz_dep_.CoverTab[23377]++
									b.AddUint16(m.vers)
									addBytesWithLength(b, m.random, 32)
									b.AddUint8LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:729
			_go_fuzz_dep_.CoverTab[23379]++
										b.AddBytes(m.sessionId)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:730
			// _ = "end of CoverTab[23379]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:731
		// _ = "end of CoverTab[23377]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:731
		_go_fuzz_dep_.CoverTab[23378]++
									b.AddUint16(m.cipherSuite)
									b.AddUint8(m.compressionMethod)

									if len(extBytes) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:735
			_go_fuzz_dep_.CoverTab[23380]++
										b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:736
				_go_fuzz_dep_.CoverTab[23381]++
											b.AddBytes(extBytes)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:737
				// _ = "end of CoverTab[23381]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:738
			// _ = "end of CoverTab[23380]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:739
			_go_fuzz_dep_.CoverTab[23382]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:739
			// _ = "end of CoverTab[23382]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:739
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:739
		// _ = "end of CoverTab[23378]"
	})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:740
	// _ = "end of CoverTab[23331]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:740
	_go_fuzz_dep_.CoverTab[23332]++

								m.raw, err = b.Bytes()
								return m.raw, err
//line /usr/local/go/src/crypto/tls/handshake_messages.go:743
	// _ = "end of CoverTab[23332]"
}

func (m *serverHelloMsg) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:746
	_go_fuzz_dep_.CoverTab[23383]++
								*m = serverHelloMsg{raw: data}
								s := cryptobyte.String(data)

								if !s.Skip(4) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:750
		_go_fuzz_dep_.CoverTab[23388]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:750
		return !s.ReadUint16(&m.vers)
									// _ = "end of CoverTab[23388]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:751
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:751
		_go_fuzz_dep_.CoverTab[23389]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:751
		return !s.ReadBytes(&m.random, 32)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:751
		// _ = "end of CoverTab[23389]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:751
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:751
		_go_fuzz_dep_.CoverTab[23390]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:751
		return !readUint8LengthPrefixed(&s, &m.sessionId)
									// _ = "end of CoverTab[23390]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:752
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:752
		_go_fuzz_dep_.CoverTab[23391]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:752
		return !s.ReadUint16(&m.cipherSuite)
									// _ = "end of CoverTab[23391]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:753
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:753
		_go_fuzz_dep_.CoverTab[23392]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:753
		return !s.ReadUint8(&m.compressionMethod)
									// _ = "end of CoverTab[23392]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:754
	}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:754
		_go_fuzz_dep_.CoverTab[23393]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:755
		// _ = "end of CoverTab[23393]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:756
		_go_fuzz_dep_.CoverTab[23394]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:756
		// _ = "end of CoverTab[23394]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:756
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:756
	// _ = "end of CoverTab[23383]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:756
	_go_fuzz_dep_.CoverTab[23384]++

								if s.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:758
		_go_fuzz_dep_.CoverTab[23395]++

									return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:760
		// _ = "end of CoverTab[23395]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:761
		_go_fuzz_dep_.CoverTab[23396]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:761
		// _ = "end of CoverTab[23396]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:761
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:761
	// _ = "end of CoverTab[23384]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:761
	_go_fuzz_dep_.CoverTab[23385]++

								var extensions cryptobyte.String
								if !s.ReadUint16LengthPrefixed(&extensions) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:764
		_go_fuzz_dep_.CoverTab[23397]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:764
		return !s.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:764
		// _ = "end of CoverTab[23397]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:764
	}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:764
		_go_fuzz_dep_.CoverTab[23398]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:765
		// _ = "end of CoverTab[23398]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:766
		_go_fuzz_dep_.CoverTab[23399]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:766
		// _ = "end of CoverTab[23399]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:766
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:766
	// _ = "end of CoverTab[23385]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:766
	_go_fuzz_dep_.CoverTab[23386]++

								seenExts := make(map[uint16]bool)
								for !extensions.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:769
		_go_fuzz_dep_.CoverTab[23400]++
									var extension uint16
									var extData cryptobyte.String
									if !extensions.ReadUint16(&extension) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:772
			_go_fuzz_dep_.CoverTab[23404]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:772
			return !extensions.ReadUint16LengthPrefixed(&extData)
										// _ = "end of CoverTab[23404]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:773
		}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:773
			_go_fuzz_dep_.CoverTab[23405]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:774
			// _ = "end of CoverTab[23405]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:775
			_go_fuzz_dep_.CoverTab[23406]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:775
			// _ = "end of CoverTab[23406]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:775
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:775
		// _ = "end of CoverTab[23400]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:775
		_go_fuzz_dep_.CoverTab[23401]++

									if seenExts[extension] {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:777
			_go_fuzz_dep_.CoverTab[23407]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:778
			// _ = "end of CoverTab[23407]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:779
			_go_fuzz_dep_.CoverTab[23408]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:779
			// _ = "end of CoverTab[23408]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:779
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:779
		// _ = "end of CoverTab[23401]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:779
		_go_fuzz_dep_.CoverTab[23402]++
									seenExts[extension] = true

									switch extension {
		case extensionStatusRequest:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:783
			_go_fuzz_dep_.CoverTab[23409]++
										m.ocspStapling = true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:784
			// _ = "end of CoverTab[23409]"
		case extensionSessionTicket:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:785
			_go_fuzz_dep_.CoverTab[23410]++
										m.ticketSupported = true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:786
			// _ = "end of CoverTab[23410]"
		case extensionRenegotiationInfo:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:787
			_go_fuzz_dep_.CoverTab[23411]++
										if !readUint8LengthPrefixed(&extData, &m.secureRenegotiation) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:788
				_go_fuzz_dep_.CoverTab[23424]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:789
				// _ = "end of CoverTab[23424]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:790
				_go_fuzz_dep_.CoverTab[23425]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:790
				// _ = "end of CoverTab[23425]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:790
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:790
			// _ = "end of CoverTab[23411]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:790
			_go_fuzz_dep_.CoverTab[23412]++
										m.secureRenegotiationSupported = true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:791
			// _ = "end of CoverTab[23412]"
		case extensionALPN:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:792
			_go_fuzz_dep_.CoverTab[23413]++
										var protoList cryptobyte.String
										if !extData.ReadUint16LengthPrefixed(&protoList) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:794
				_go_fuzz_dep_.CoverTab[23426]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:794
				return protoList.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:794
				// _ = "end of CoverTab[23426]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:794
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:794
				_go_fuzz_dep_.CoverTab[23427]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:795
				// _ = "end of CoverTab[23427]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:796
				_go_fuzz_dep_.CoverTab[23428]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:796
				// _ = "end of CoverTab[23428]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:796
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:796
			// _ = "end of CoverTab[23413]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:796
			_go_fuzz_dep_.CoverTab[23414]++
										var proto cryptobyte.String
										if !protoList.ReadUint8LengthPrefixed(&proto) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:798
				_go_fuzz_dep_.CoverTab[23429]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:798
				return proto.Empty()
											// _ = "end of CoverTab[23429]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:799
			}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:799
				_go_fuzz_dep_.CoverTab[23430]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:799
				return !protoList.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:799
				// _ = "end of CoverTab[23430]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:799
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:799
				_go_fuzz_dep_.CoverTab[23431]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:800
				// _ = "end of CoverTab[23431]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:801
				_go_fuzz_dep_.CoverTab[23432]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:801
				// _ = "end of CoverTab[23432]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:801
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:801
			// _ = "end of CoverTab[23414]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:801
			_go_fuzz_dep_.CoverTab[23415]++
										m.alpnProtocol = string(proto)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:802
			// _ = "end of CoverTab[23415]"
		case extensionSCT:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:803
			_go_fuzz_dep_.CoverTab[23416]++
										var sctList cryptobyte.String
										if !extData.ReadUint16LengthPrefixed(&sctList) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:805
				_go_fuzz_dep_.CoverTab[23433]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:805
				return sctList.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:805
				// _ = "end of CoverTab[23433]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:805
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:805
				_go_fuzz_dep_.CoverTab[23434]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:806
				// _ = "end of CoverTab[23434]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:807
				_go_fuzz_dep_.CoverTab[23435]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:807
				// _ = "end of CoverTab[23435]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:807
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:807
			// _ = "end of CoverTab[23416]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:807
			_go_fuzz_dep_.CoverTab[23417]++
										for !sctList.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:808
				_go_fuzz_dep_.CoverTab[23436]++
											var sct []byte
											if !readUint16LengthPrefixed(&sctList, &sct) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:810
					_go_fuzz_dep_.CoverTab[23438]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:810
					return len(sct) == 0
												// _ = "end of CoverTab[23438]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:811
				}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:811
					_go_fuzz_dep_.CoverTab[23439]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:812
					// _ = "end of CoverTab[23439]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:813
					_go_fuzz_dep_.CoverTab[23440]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:813
					// _ = "end of CoverTab[23440]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:813
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:813
				// _ = "end of CoverTab[23436]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:813
				_go_fuzz_dep_.CoverTab[23437]++
											m.scts = append(m.scts, sct)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:814
				// _ = "end of CoverTab[23437]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:815
			// _ = "end of CoverTab[23417]"
		case extensionSupportedVersions:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:816
			_go_fuzz_dep_.CoverTab[23418]++
										if !extData.ReadUint16(&m.supportedVersion) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:817
				_go_fuzz_dep_.CoverTab[23441]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:818
				// _ = "end of CoverTab[23441]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:819
				_go_fuzz_dep_.CoverTab[23442]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:819
				// _ = "end of CoverTab[23442]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:819
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:819
			// _ = "end of CoverTab[23418]"
		case extensionCookie:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:820
			_go_fuzz_dep_.CoverTab[23419]++
										if !readUint16LengthPrefixed(&extData, &m.cookie) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:821
				_go_fuzz_dep_.CoverTab[23443]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:821
				return len(m.cookie) == 0
											// _ = "end of CoverTab[23443]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:822
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:822
				_go_fuzz_dep_.CoverTab[23444]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:823
				// _ = "end of CoverTab[23444]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:824
				_go_fuzz_dep_.CoverTab[23445]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:824
				// _ = "end of CoverTab[23445]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:824
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:824
			// _ = "end of CoverTab[23419]"
		case extensionKeyShare:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:825
			_go_fuzz_dep_.CoverTab[23420]++

//line /usr/local/go/src/crypto/tls/handshake_messages.go:828
			if len(extData) == 2 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:828
				_go_fuzz_dep_.CoverTab[23446]++
											if !extData.ReadUint16((*uint16)(&m.selectedGroup)) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:829
					_go_fuzz_dep_.CoverTab[23447]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:830
					// _ = "end of CoverTab[23447]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:831
					_go_fuzz_dep_.CoverTab[23448]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:831
					// _ = "end of CoverTab[23448]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:831
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:831
				// _ = "end of CoverTab[23446]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:832
				_go_fuzz_dep_.CoverTab[23449]++
											if !extData.ReadUint16((*uint16)(&m.serverShare.group)) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:833
					_go_fuzz_dep_.CoverTab[23450]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:833
					return !readUint16LengthPrefixed(&extData, &m.serverShare.data)
												// _ = "end of CoverTab[23450]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:834
				}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:834
					_go_fuzz_dep_.CoverTab[23451]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:835
					// _ = "end of CoverTab[23451]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:836
					_go_fuzz_dep_.CoverTab[23452]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:836
					// _ = "end of CoverTab[23452]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:836
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:836
				// _ = "end of CoverTab[23449]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:837
			// _ = "end of CoverTab[23420]"
		case extensionPreSharedKey:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:838
			_go_fuzz_dep_.CoverTab[23421]++
										m.selectedIdentityPresent = true
										if !extData.ReadUint16(&m.selectedIdentity) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:840
				_go_fuzz_dep_.CoverTab[23453]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:841
				// _ = "end of CoverTab[23453]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:842
				_go_fuzz_dep_.CoverTab[23454]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:842
				// _ = "end of CoverTab[23454]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:842
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:842
			// _ = "end of CoverTab[23421]"
		case extensionSupportedPoints:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:843
			_go_fuzz_dep_.CoverTab[23422]++

										if !readUint8LengthPrefixed(&extData, &m.supportedPoints) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:845
				_go_fuzz_dep_.CoverTab[23455]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:845
				return len(m.supportedPoints) == 0
											// _ = "end of CoverTab[23455]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:846
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:846
				_go_fuzz_dep_.CoverTab[23456]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:847
				// _ = "end of CoverTab[23456]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:848
				_go_fuzz_dep_.CoverTab[23457]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:848
				// _ = "end of CoverTab[23457]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:848
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:848
			// _ = "end of CoverTab[23422]"
		default:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:849
			_go_fuzz_dep_.CoverTab[23423]++

										continue
//line /usr/local/go/src/crypto/tls/handshake_messages.go:851
			// _ = "end of CoverTab[23423]"
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:852
		// _ = "end of CoverTab[23402]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:852
		_go_fuzz_dep_.CoverTab[23403]++

									if !extData.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:854
			_go_fuzz_dep_.CoverTab[23458]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:855
			// _ = "end of CoverTab[23458]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:856
			_go_fuzz_dep_.CoverTab[23459]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:856
			// _ = "end of CoverTab[23459]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:856
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:856
		// _ = "end of CoverTab[23403]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:857
	// _ = "end of CoverTab[23386]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:857
	_go_fuzz_dep_.CoverTab[23387]++

								return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:859
	// _ = "end of CoverTab[23387]"
}

type encryptedExtensionsMsg struct {
	raw		[]byte
	alpnProtocol	string
}

func (m *encryptedExtensionsMsg) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:867
	_go_fuzz_dep_.CoverTab[23460]++
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:868
		_go_fuzz_dep_.CoverTab[23463]++
									return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:869
		// _ = "end of CoverTab[23463]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:870
		_go_fuzz_dep_.CoverTab[23464]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:870
		// _ = "end of CoverTab[23464]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:870
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:870
	// _ = "end of CoverTab[23460]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:870
	_go_fuzz_dep_.CoverTab[23461]++

								var b cryptobyte.Builder
								b.AddUint8(typeEncryptedExtensions)
								b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:874
		_go_fuzz_dep_.CoverTab[23465]++
									b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:875
			_go_fuzz_dep_.CoverTab[23466]++
										if len(m.alpnProtocol) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:876
				_go_fuzz_dep_.CoverTab[23467]++
											b.AddUint16(extensionALPN)
											b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:878
					_go_fuzz_dep_.CoverTab[23468]++
												b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:879
						_go_fuzz_dep_.CoverTab[23469]++
													b.AddUint8LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:880
							_go_fuzz_dep_.CoverTab[23470]++
														b.AddBytes([]byte(m.alpnProtocol))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:881
							// _ = "end of CoverTab[23470]"
						})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:882
						// _ = "end of CoverTab[23469]"
					})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:883
					// _ = "end of CoverTab[23468]"
				})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:884
				// _ = "end of CoverTab[23467]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:885
				_go_fuzz_dep_.CoverTab[23471]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:885
				// _ = "end of CoverTab[23471]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:885
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:885
			// _ = "end of CoverTab[23466]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:886
		// _ = "end of CoverTab[23465]"
	})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:887
	// _ = "end of CoverTab[23461]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:887
	_go_fuzz_dep_.CoverTab[23462]++

								var err error
								m.raw, err = b.Bytes()
								return m.raw, err
//line /usr/local/go/src/crypto/tls/handshake_messages.go:891
	// _ = "end of CoverTab[23462]"
}

func (m *encryptedExtensionsMsg) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:894
	_go_fuzz_dep_.CoverTab[23472]++
								*m = encryptedExtensionsMsg{raw: data}
								s := cryptobyte.String(data)

								var extensions cryptobyte.String
								if !s.Skip(4) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:899
		_go_fuzz_dep_.CoverTab[23475]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:899
		return !s.ReadUint16LengthPrefixed(&extensions)
									// _ = "end of CoverTab[23475]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:900
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:900
		_go_fuzz_dep_.CoverTab[23476]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:900
		return !s.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:900
		// _ = "end of CoverTab[23476]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:900
	}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:900
		_go_fuzz_dep_.CoverTab[23477]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:901
		// _ = "end of CoverTab[23477]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:902
		_go_fuzz_dep_.CoverTab[23478]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:902
		// _ = "end of CoverTab[23478]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:902
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:902
	// _ = "end of CoverTab[23472]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:902
	_go_fuzz_dep_.CoverTab[23473]++

								for !extensions.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:904
		_go_fuzz_dep_.CoverTab[23479]++
									var extension uint16
									var extData cryptobyte.String
									if !extensions.ReadUint16(&extension) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:907
			_go_fuzz_dep_.CoverTab[23482]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:907
			return !extensions.ReadUint16LengthPrefixed(&extData)
										// _ = "end of CoverTab[23482]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:908
		}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:908
			_go_fuzz_dep_.CoverTab[23483]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:909
			// _ = "end of CoverTab[23483]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:910
			_go_fuzz_dep_.CoverTab[23484]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:910
			// _ = "end of CoverTab[23484]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:910
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:910
		// _ = "end of CoverTab[23479]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:910
		_go_fuzz_dep_.CoverTab[23480]++

									switch extension {
		case extensionALPN:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:913
			_go_fuzz_dep_.CoverTab[23485]++
										var protoList cryptobyte.String
										if !extData.ReadUint16LengthPrefixed(&protoList) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:915
				_go_fuzz_dep_.CoverTab[23489]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:915
				return protoList.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:915
				// _ = "end of CoverTab[23489]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:915
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:915
				_go_fuzz_dep_.CoverTab[23490]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:916
				// _ = "end of CoverTab[23490]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:917
				_go_fuzz_dep_.CoverTab[23491]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:917
				// _ = "end of CoverTab[23491]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:917
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:917
			// _ = "end of CoverTab[23485]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:917
			_go_fuzz_dep_.CoverTab[23486]++
										var proto cryptobyte.String
										if !protoList.ReadUint8LengthPrefixed(&proto) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:919
				_go_fuzz_dep_.CoverTab[23492]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:919
				return proto.Empty()
											// _ = "end of CoverTab[23492]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:920
			}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:920
				_go_fuzz_dep_.CoverTab[23493]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:920
				return !protoList.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:920
				// _ = "end of CoverTab[23493]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:920
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:920
				_go_fuzz_dep_.CoverTab[23494]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:921
				// _ = "end of CoverTab[23494]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:922
				_go_fuzz_dep_.CoverTab[23495]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:922
				// _ = "end of CoverTab[23495]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:922
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:922
			// _ = "end of CoverTab[23486]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:922
			_go_fuzz_dep_.CoverTab[23487]++
										m.alpnProtocol = string(proto)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:923
			// _ = "end of CoverTab[23487]"
		default:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:924
			_go_fuzz_dep_.CoverTab[23488]++

										continue
//line /usr/local/go/src/crypto/tls/handshake_messages.go:926
			// _ = "end of CoverTab[23488]"
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:927
		// _ = "end of CoverTab[23480]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:927
		_go_fuzz_dep_.CoverTab[23481]++

									if !extData.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:929
			_go_fuzz_dep_.CoverTab[23496]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:930
			// _ = "end of CoverTab[23496]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:931
			_go_fuzz_dep_.CoverTab[23497]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:931
			// _ = "end of CoverTab[23497]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:931
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:931
		// _ = "end of CoverTab[23481]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:932
	// _ = "end of CoverTab[23473]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:932
	_go_fuzz_dep_.CoverTab[23474]++

								return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:934
	// _ = "end of CoverTab[23474]"
}

type endOfEarlyDataMsg struct{}

func (m *endOfEarlyDataMsg) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:939
	_go_fuzz_dep_.CoverTab[23498]++
								x := make([]byte, 4)
								x[0] = typeEndOfEarlyData
								return x, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:942
	// _ = "end of CoverTab[23498]"
}

func (m *endOfEarlyDataMsg) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:945
	_go_fuzz_dep_.CoverTab[23499]++
								return len(data) == 4
//line /usr/local/go/src/crypto/tls/handshake_messages.go:946
	// _ = "end of CoverTab[23499]"
}

type keyUpdateMsg struct {
	raw		[]byte
	updateRequested	bool
}

func (m *keyUpdateMsg) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:954
	_go_fuzz_dep_.CoverTab[23500]++
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:955
		_go_fuzz_dep_.CoverTab[23503]++
									return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:956
		// _ = "end of CoverTab[23503]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:957
		_go_fuzz_dep_.CoverTab[23504]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:957
		// _ = "end of CoverTab[23504]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:957
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:957
	// _ = "end of CoverTab[23500]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:957
	_go_fuzz_dep_.CoverTab[23501]++

								var b cryptobyte.Builder
								b.AddUint8(typeKeyUpdate)
								b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:961
		_go_fuzz_dep_.CoverTab[23505]++
									if m.updateRequested {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:962
			_go_fuzz_dep_.CoverTab[23506]++
										b.AddUint8(1)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:963
			// _ = "end of CoverTab[23506]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:964
			_go_fuzz_dep_.CoverTab[23507]++
										b.AddUint8(0)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:965
			// _ = "end of CoverTab[23507]"
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:966
		// _ = "end of CoverTab[23505]"
	})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:967
	// _ = "end of CoverTab[23501]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:967
	_go_fuzz_dep_.CoverTab[23502]++

								var err error
								m.raw, err = b.Bytes()
								return m.raw, err
//line /usr/local/go/src/crypto/tls/handshake_messages.go:971
	// _ = "end of CoverTab[23502]"
}

func (m *keyUpdateMsg) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:974
	_go_fuzz_dep_.CoverTab[23508]++
								m.raw = data
								s := cryptobyte.String(data)

								var updateRequested uint8
								if !s.Skip(4) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:979
		_go_fuzz_dep_.CoverTab[23511]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:979
		return !s.ReadUint8(&updateRequested)
									// _ = "end of CoverTab[23511]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:980
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:980
		_go_fuzz_dep_.CoverTab[23512]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:980
		return !s.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:980
		// _ = "end of CoverTab[23512]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:980
	}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:980
		_go_fuzz_dep_.CoverTab[23513]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:981
		// _ = "end of CoverTab[23513]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:982
		_go_fuzz_dep_.CoverTab[23514]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:982
		// _ = "end of CoverTab[23514]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:982
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:982
	// _ = "end of CoverTab[23508]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:982
	_go_fuzz_dep_.CoverTab[23509]++
								switch updateRequested {
	case 0:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:984
		_go_fuzz_dep_.CoverTab[23515]++
									m.updateRequested = false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:985
		// _ = "end of CoverTab[23515]"
	case 1:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:986
		_go_fuzz_dep_.CoverTab[23516]++
									m.updateRequested = true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:987
		// _ = "end of CoverTab[23516]"
	default:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:988
		_go_fuzz_dep_.CoverTab[23517]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:989
		// _ = "end of CoverTab[23517]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:990
	// _ = "end of CoverTab[23509]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:990
	_go_fuzz_dep_.CoverTab[23510]++
								return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:991
	// _ = "end of CoverTab[23510]"
}

type newSessionTicketMsgTLS13 struct {
	raw		[]byte
	lifetime	uint32
	ageAdd		uint32
	nonce		[]byte
	label		[]byte
	maxEarlyData	uint32
}

func (m *newSessionTicketMsgTLS13) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1003
	_go_fuzz_dep_.CoverTab[23518]++
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1004
		_go_fuzz_dep_.CoverTab[23521]++
									return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1005
		// _ = "end of CoverTab[23521]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1006
		_go_fuzz_dep_.CoverTab[23522]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1006
		// _ = "end of CoverTab[23522]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1006
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1006
	// _ = "end of CoverTab[23518]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1006
	_go_fuzz_dep_.CoverTab[23519]++

								var b cryptobyte.Builder
								b.AddUint8(typeNewSessionTicket)
								b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1010
		_go_fuzz_dep_.CoverTab[23523]++
									b.AddUint32(m.lifetime)
									b.AddUint32(m.ageAdd)
									b.AddUint8LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1013
			_go_fuzz_dep_.CoverTab[23526]++
										b.AddBytes(m.nonce)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1014
			// _ = "end of CoverTab[23526]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1015
		// _ = "end of CoverTab[23523]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1015
		_go_fuzz_dep_.CoverTab[23524]++
									b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1016
			_go_fuzz_dep_.CoverTab[23527]++
										b.AddBytes(m.label)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1017
			// _ = "end of CoverTab[23527]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1018
		// _ = "end of CoverTab[23524]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1018
		_go_fuzz_dep_.CoverTab[23525]++

									b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1020
			_go_fuzz_dep_.CoverTab[23528]++
										if m.maxEarlyData > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1021
				_go_fuzz_dep_.CoverTab[23529]++
											b.AddUint16(extensionEarlyData)
											b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1023
					_go_fuzz_dep_.CoverTab[23530]++
												b.AddUint32(m.maxEarlyData)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1024
					// _ = "end of CoverTab[23530]"
				})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1025
				// _ = "end of CoverTab[23529]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1026
				_go_fuzz_dep_.CoverTab[23531]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1026
				// _ = "end of CoverTab[23531]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1026
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1026
			// _ = "end of CoverTab[23528]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1027
		// _ = "end of CoverTab[23525]"
	})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1028
	// _ = "end of CoverTab[23519]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1028
	_go_fuzz_dep_.CoverTab[23520]++

								var err error
								m.raw, err = b.Bytes()
								return m.raw, err
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1032
	// _ = "end of CoverTab[23520]"
}

func (m *newSessionTicketMsgTLS13) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1035
	_go_fuzz_dep_.CoverTab[23532]++
								*m = newSessionTicketMsgTLS13{raw: data}
								s := cryptobyte.String(data)

								var extensions cryptobyte.String
								if !s.Skip(4) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1040
		_go_fuzz_dep_.CoverTab[23535]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1040
		return !s.ReadUint32(&m.lifetime)
									// _ = "end of CoverTab[23535]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1041
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1041
		_go_fuzz_dep_.CoverTab[23536]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1041
		return !s.ReadUint32(&m.ageAdd)
									// _ = "end of CoverTab[23536]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1042
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1042
		_go_fuzz_dep_.CoverTab[23537]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1042
		return !readUint8LengthPrefixed(&s, &m.nonce)
									// _ = "end of CoverTab[23537]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1043
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1043
		_go_fuzz_dep_.CoverTab[23538]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1043
		return !readUint16LengthPrefixed(&s, &m.label)
									// _ = "end of CoverTab[23538]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1044
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1044
		_go_fuzz_dep_.CoverTab[23539]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1044
		return !s.ReadUint16LengthPrefixed(&extensions)
									// _ = "end of CoverTab[23539]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1045
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1045
		_go_fuzz_dep_.CoverTab[23540]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1045
		return !s.Empty()
									// _ = "end of CoverTab[23540]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1046
	}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1046
		_go_fuzz_dep_.CoverTab[23541]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1047
		// _ = "end of CoverTab[23541]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1048
		_go_fuzz_dep_.CoverTab[23542]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1048
		// _ = "end of CoverTab[23542]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1048
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1048
	// _ = "end of CoverTab[23532]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1048
	_go_fuzz_dep_.CoverTab[23533]++

								for !extensions.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1050
		_go_fuzz_dep_.CoverTab[23543]++
									var extension uint16
									var extData cryptobyte.String
									if !extensions.ReadUint16(&extension) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1053
			_go_fuzz_dep_.CoverTab[23546]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1053
			return !extensions.ReadUint16LengthPrefixed(&extData)
										// _ = "end of CoverTab[23546]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1054
		}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1054
			_go_fuzz_dep_.CoverTab[23547]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1055
			// _ = "end of CoverTab[23547]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1056
			_go_fuzz_dep_.CoverTab[23548]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1056
			// _ = "end of CoverTab[23548]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1056
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1056
		// _ = "end of CoverTab[23543]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1056
		_go_fuzz_dep_.CoverTab[23544]++

									switch extension {
		case extensionEarlyData:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1059
			_go_fuzz_dep_.CoverTab[23549]++
										if !extData.ReadUint32(&m.maxEarlyData) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1060
				_go_fuzz_dep_.CoverTab[23551]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1061
				// _ = "end of CoverTab[23551]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1062
				_go_fuzz_dep_.CoverTab[23552]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1062
				// _ = "end of CoverTab[23552]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1062
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1062
			// _ = "end of CoverTab[23549]"
		default:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1063
			_go_fuzz_dep_.CoverTab[23550]++

										continue
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1065
			// _ = "end of CoverTab[23550]"
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1066
		// _ = "end of CoverTab[23544]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1066
		_go_fuzz_dep_.CoverTab[23545]++

									if !extData.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1068
			_go_fuzz_dep_.CoverTab[23553]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1069
			// _ = "end of CoverTab[23553]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1070
			_go_fuzz_dep_.CoverTab[23554]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1070
			// _ = "end of CoverTab[23554]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1070
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1070
		// _ = "end of CoverTab[23545]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1071
	// _ = "end of CoverTab[23533]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1071
	_go_fuzz_dep_.CoverTab[23534]++

								return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1073
	// _ = "end of CoverTab[23534]"
}

type certificateRequestMsgTLS13 struct {
	raw					[]byte
	ocspStapling				bool
	scts					bool
	supportedSignatureAlgorithms		[]SignatureScheme
	supportedSignatureAlgorithmsCert	[]SignatureScheme
	certificateAuthorities			[][]byte
}

func (m *certificateRequestMsgTLS13) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1085
	_go_fuzz_dep_.CoverTab[23555]++
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1086
		_go_fuzz_dep_.CoverTab[23558]++
									return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1087
		// _ = "end of CoverTab[23558]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1088
		_go_fuzz_dep_.CoverTab[23559]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1088
		// _ = "end of CoverTab[23559]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1088
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1088
	// _ = "end of CoverTab[23555]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1088
	_go_fuzz_dep_.CoverTab[23556]++

								var b cryptobyte.Builder
								b.AddUint8(typeCertificateRequest)
								b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1092
		_go_fuzz_dep_.CoverTab[23560]++

//line /usr/local/go/src/crypto/tls/handshake_messages.go:1095
		b.AddUint8(0)

		b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1097
			_go_fuzz_dep_.CoverTab[23561]++
										if m.ocspStapling {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1098
				_go_fuzz_dep_.CoverTab[23566]++
											b.AddUint16(extensionStatusRequest)
											b.AddUint16(0)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1100
				// _ = "end of CoverTab[23566]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1101
				_go_fuzz_dep_.CoverTab[23567]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1101
				// _ = "end of CoverTab[23567]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1101
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1101
			// _ = "end of CoverTab[23561]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1101
			_go_fuzz_dep_.CoverTab[23562]++
										if m.scts {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1102
				_go_fuzz_dep_.CoverTab[23568]++

//line /usr/local/go/src/crypto/tls/handshake_messages.go:1108
				b.AddUint16(extensionSCT)
											b.AddUint16(0)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1109
				// _ = "end of CoverTab[23568]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1110
				_go_fuzz_dep_.CoverTab[23569]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1110
				// _ = "end of CoverTab[23569]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1110
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1110
			// _ = "end of CoverTab[23562]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1110
			_go_fuzz_dep_.CoverTab[23563]++
										if len(m.supportedSignatureAlgorithms) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1111
				_go_fuzz_dep_.CoverTab[23570]++
											b.AddUint16(extensionSignatureAlgorithms)
											b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1113
					_go_fuzz_dep_.CoverTab[23571]++
												b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1114
						_go_fuzz_dep_.CoverTab[23572]++
													for _, sigAlgo := range m.supportedSignatureAlgorithms {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1115
							_go_fuzz_dep_.CoverTab[23573]++
														b.AddUint16(uint16(sigAlgo))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1116
							// _ = "end of CoverTab[23573]"
						}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1117
						// _ = "end of CoverTab[23572]"
					})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1118
					// _ = "end of CoverTab[23571]"
				})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1119
				// _ = "end of CoverTab[23570]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1120
				_go_fuzz_dep_.CoverTab[23574]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1120
				// _ = "end of CoverTab[23574]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1120
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1120
			// _ = "end of CoverTab[23563]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1120
			_go_fuzz_dep_.CoverTab[23564]++
										if len(m.supportedSignatureAlgorithmsCert) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1121
				_go_fuzz_dep_.CoverTab[23575]++
											b.AddUint16(extensionSignatureAlgorithmsCert)
											b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1123
					_go_fuzz_dep_.CoverTab[23576]++
												b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1124
						_go_fuzz_dep_.CoverTab[23577]++
													for _, sigAlgo := range m.supportedSignatureAlgorithmsCert {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1125
							_go_fuzz_dep_.CoverTab[23578]++
														b.AddUint16(uint16(sigAlgo))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1126
							// _ = "end of CoverTab[23578]"
						}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1127
						// _ = "end of CoverTab[23577]"
					})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1128
					// _ = "end of CoverTab[23576]"
				})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1129
				// _ = "end of CoverTab[23575]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1130
				_go_fuzz_dep_.CoverTab[23579]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1130
				// _ = "end of CoverTab[23579]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1130
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1130
			// _ = "end of CoverTab[23564]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1130
			_go_fuzz_dep_.CoverTab[23565]++
										if len(m.certificateAuthorities) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1131
				_go_fuzz_dep_.CoverTab[23580]++
											b.AddUint16(extensionCertificateAuthorities)
											b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1133
					_go_fuzz_dep_.CoverTab[23581]++
												b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1134
						_go_fuzz_dep_.CoverTab[23582]++
													for _, ca := range m.certificateAuthorities {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1135
							_go_fuzz_dep_.CoverTab[23583]++
														b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1136
								_go_fuzz_dep_.CoverTab[23584]++
															b.AddBytes(ca)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1137
								// _ = "end of CoverTab[23584]"
							})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1138
							// _ = "end of CoverTab[23583]"
						}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1139
						// _ = "end of CoverTab[23582]"
					})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1140
					// _ = "end of CoverTab[23581]"
				})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1141
				// _ = "end of CoverTab[23580]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1142
				_go_fuzz_dep_.CoverTab[23585]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1142
				// _ = "end of CoverTab[23585]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1142
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1142
			// _ = "end of CoverTab[23565]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1143
		// _ = "end of CoverTab[23560]"
	})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1144
	// _ = "end of CoverTab[23556]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1144
	_go_fuzz_dep_.CoverTab[23557]++

								var err error
								m.raw, err = b.Bytes()
								return m.raw, err
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1148
	// _ = "end of CoverTab[23557]"
}

func (m *certificateRequestMsgTLS13) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1151
	_go_fuzz_dep_.CoverTab[23586]++
								*m = certificateRequestMsgTLS13{raw: data}
								s := cryptobyte.String(data)

								var context, extensions cryptobyte.String
								if !s.Skip(4) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1156
		_go_fuzz_dep_.CoverTab[23589]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1156
		return !s.ReadUint8LengthPrefixed(&context)
									// _ = "end of CoverTab[23589]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1157
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1157
		_go_fuzz_dep_.CoverTab[23590]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1157
		return !context.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1157
		// _ = "end of CoverTab[23590]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1157
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1157
		_go_fuzz_dep_.CoverTab[23591]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1157
		return !s.ReadUint16LengthPrefixed(&extensions)
									// _ = "end of CoverTab[23591]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1158
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1158
		_go_fuzz_dep_.CoverTab[23592]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1158
		return !s.Empty()
									// _ = "end of CoverTab[23592]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1159
	}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1159
		_go_fuzz_dep_.CoverTab[23593]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1160
		// _ = "end of CoverTab[23593]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1161
		_go_fuzz_dep_.CoverTab[23594]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1161
		// _ = "end of CoverTab[23594]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1161
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1161
	// _ = "end of CoverTab[23586]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1161
	_go_fuzz_dep_.CoverTab[23587]++

								for !extensions.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1163
		_go_fuzz_dep_.CoverTab[23595]++
									var extension uint16
									var extData cryptobyte.String
									if !extensions.ReadUint16(&extension) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1166
			_go_fuzz_dep_.CoverTab[23598]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1166
			return !extensions.ReadUint16LengthPrefixed(&extData)
										// _ = "end of CoverTab[23598]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1167
		}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1167
			_go_fuzz_dep_.CoverTab[23599]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1168
			// _ = "end of CoverTab[23599]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1169
			_go_fuzz_dep_.CoverTab[23600]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1169
			// _ = "end of CoverTab[23600]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1169
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1169
		// _ = "end of CoverTab[23595]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1169
		_go_fuzz_dep_.CoverTab[23596]++

									switch extension {
		case extensionStatusRequest:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1172
			_go_fuzz_dep_.CoverTab[23601]++
										m.ocspStapling = true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1173
			// _ = "end of CoverTab[23601]"
		case extensionSCT:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1174
			_go_fuzz_dep_.CoverTab[23602]++
										m.scts = true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1175
			// _ = "end of CoverTab[23602]"
		case extensionSignatureAlgorithms:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1176
			_go_fuzz_dep_.CoverTab[23603]++
										var sigAndAlgs cryptobyte.String
										if !extData.ReadUint16LengthPrefixed(&sigAndAlgs) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1178
				_go_fuzz_dep_.CoverTab[23610]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1178
				return sigAndAlgs.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1178
				// _ = "end of CoverTab[23610]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1178
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1178
				_go_fuzz_dep_.CoverTab[23611]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1179
				// _ = "end of CoverTab[23611]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1180
				_go_fuzz_dep_.CoverTab[23612]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1180
				// _ = "end of CoverTab[23612]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1180
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1180
			// _ = "end of CoverTab[23603]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1180
			_go_fuzz_dep_.CoverTab[23604]++
										for !sigAndAlgs.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1181
				_go_fuzz_dep_.CoverTab[23613]++
											var sigAndAlg uint16
											if !sigAndAlgs.ReadUint16(&sigAndAlg) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1183
					_go_fuzz_dep_.CoverTab[23615]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1184
					// _ = "end of CoverTab[23615]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1185
					_go_fuzz_dep_.CoverTab[23616]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1185
					// _ = "end of CoverTab[23616]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1185
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1185
				// _ = "end of CoverTab[23613]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1185
				_go_fuzz_dep_.CoverTab[23614]++
											m.supportedSignatureAlgorithms = append(
					m.supportedSignatureAlgorithms, SignatureScheme(sigAndAlg))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1187
				// _ = "end of CoverTab[23614]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1188
			// _ = "end of CoverTab[23604]"
		case extensionSignatureAlgorithmsCert:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1189
			_go_fuzz_dep_.CoverTab[23605]++
										var sigAndAlgs cryptobyte.String
										if !extData.ReadUint16LengthPrefixed(&sigAndAlgs) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1191
				_go_fuzz_dep_.CoverTab[23617]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1191
				return sigAndAlgs.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1191
				// _ = "end of CoverTab[23617]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1191
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1191
				_go_fuzz_dep_.CoverTab[23618]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1192
				// _ = "end of CoverTab[23618]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1193
				_go_fuzz_dep_.CoverTab[23619]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1193
				// _ = "end of CoverTab[23619]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1193
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1193
			// _ = "end of CoverTab[23605]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1193
			_go_fuzz_dep_.CoverTab[23606]++
										for !sigAndAlgs.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1194
				_go_fuzz_dep_.CoverTab[23620]++
											var sigAndAlg uint16
											if !sigAndAlgs.ReadUint16(&sigAndAlg) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1196
					_go_fuzz_dep_.CoverTab[23622]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1197
					// _ = "end of CoverTab[23622]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1198
					_go_fuzz_dep_.CoverTab[23623]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1198
					// _ = "end of CoverTab[23623]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1198
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1198
				// _ = "end of CoverTab[23620]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1198
				_go_fuzz_dep_.CoverTab[23621]++
											m.supportedSignatureAlgorithmsCert = append(
					m.supportedSignatureAlgorithmsCert, SignatureScheme(sigAndAlg))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1200
				// _ = "end of CoverTab[23621]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1201
			// _ = "end of CoverTab[23606]"
		case extensionCertificateAuthorities:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1202
			_go_fuzz_dep_.CoverTab[23607]++
										var auths cryptobyte.String
										if !extData.ReadUint16LengthPrefixed(&auths) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1204
				_go_fuzz_dep_.CoverTab[23624]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1204
				return auths.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1204
				// _ = "end of CoverTab[23624]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1204
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1204
				_go_fuzz_dep_.CoverTab[23625]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1205
				// _ = "end of CoverTab[23625]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1206
				_go_fuzz_dep_.CoverTab[23626]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1206
				// _ = "end of CoverTab[23626]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1206
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1206
			// _ = "end of CoverTab[23607]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1206
			_go_fuzz_dep_.CoverTab[23608]++
										for !auths.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1207
				_go_fuzz_dep_.CoverTab[23627]++
											var ca []byte
											if !readUint16LengthPrefixed(&auths, &ca) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1209
					_go_fuzz_dep_.CoverTab[23629]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1209
					return len(ca) == 0
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1209
					// _ = "end of CoverTab[23629]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1209
				}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1209
					_go_fuzz_dep_.CoverTab[23630]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1210
					// _ = "end of CoverTab[23630]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1211
					_go_fuzz_dep_.CoverTab[23631]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1211
					// _ = "end of CoverTab[23631]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1211
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1211
				// _ = "end of CoverTab[23627]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1211
				_go_fuzz_dep_.CoverTab[23628]++
											m.certificateAuthorities = append(m.certificateAuthorities, ca)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1212
				// _ = "end of CoverTab[23628]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1213
			// _ = "end of CoverTab[23608]"
		default:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1214
			_go_fuzz_dep_.CoverTab[23609]++

										continue
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1216
			// _ = "end of CoverTab[23609]"
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1217
		// _ = "end of CoverTab[23596]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1217
		_go_fuzz_dep_.CoverTab[23597]++

									if !extData.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1219
			_go_fuzz_dep_.CoverTab[23632]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1220
			// _ = "end of CoverTab[23632]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1221
			_go_fuzz_dep_.CoverTab[23633]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1221
			// _ = "end of CoverTab[23633]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1221
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1221
		// _ = "end of CoverTab[23597]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1222
	// _ = "end of CoverTab[23587]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1222
	_go_fuzz_dep_.CoverTab[23588]++

								return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1224
	// _ = "end of CoverTab[23588]"
}

type certificateMsg struct {
	raw		[]byte
	certificates	[][]byte
}

func (m *certificateMsg) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1232
	_go_fuzz_dep_.CoverTab[23634]++
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1233
		_go_fuzz_dep_.CoverTab[23638]++
									return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1234
		// _ = "end of CoverTab[23638]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1235
		_go_fuzz_dep_.CoverTab[23639]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1235
		// _ = "end of CoverTab[23639]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1235
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1235
	// _ = "end of CoverTab[23634]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1235
	_go_fuzz_dep_.CoverTab[23635]++

								var i int
								for _, slice := range m.certificates {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1238
		_go_fuzz_dep_.CoverTab[23640]++
									i += len(slice)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1239
		// _ = "end of CoverTab[23640]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1240
	// _ = "end of CoverTab[23635]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1240
	_go_fuzz_dep_.CoverTab[23636]++

								length := 3 + 3*len(m.certificates) + i
								x := make([]byte, 4+length)
								x[0] = typeCertificate
								x[1] = uint8(length >> 16)
								x[2] = uint8(length >> 8)
								x[3] = uint8(length)

								certificateOctets := length - 3
								x[4] = uint8(certificateOctets >> 16)
								x[5] = uint8(certificateOctets >> 8)
								x[6] = uint8(certificateOctets)

								y := x[7:]
								for _, slice := range m.certificates {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1255
		_go_fuzz_dep_.CoverTab[23641]++
									y[0] = uint8(len(slice) >> 16)
									y[1] = uint8(len(slice) >> 8)
									y[2] = uint8(len(slice))
									copy(y[3:], slice)
									y = y[3+len(slice):]
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1260
		// _ = "end of CoverTab[23641]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1261
	// _ = "end of CoverTab[23636]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1261
	_go_fuzz_dep_.CoverTab[23637]++

								m.raw = x
								return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1264
	// _ = "end of CoverTab[23637]"
}

func (m *certificateMsg) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1267
	_go_fuzz_dep_.CoverTab[23642]++
								if len(data) < 7 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1268
		_go_fuzz_dep_.CoverTab[23647]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1269
		// _ = "end of CoverTab[23647]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1270
		_go_fuzz_dep_.CoverTab[23648]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1270
		// _ = "end of CoverTab[23648]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1270
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1270
	// _ = "end of CoverTab[23642]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1270
	_go_fuzz_dep_.CoverTab[23643]++

								m.raw = data
								certsLen := uint32(data[4])<<16 | uint32(data[5])<<8 | uint32(data[6])
								if uint32(len(data)) != certsLen+7 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1274
		_go_fuzz_dep_.CoverTab[23649]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1275
		// _ = "end of CoverTab[23649]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1276
		_go_fuzz_dep_.CoverTab[23650]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1276
		// _ = "end of CoverTab[23650]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1276
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1276
	// _ = "end of CoverTab[23643]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1276
	_go_fuzz_dep_.CoverTab[23644]++

								numCerts := 0
								d := data[7:]
								for certsLen > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1280
		_go_fuzz_dep_.CoverTab[23651]++
									if len(d) < 4 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1281
			_go_fuzz_dep_.CoverTab[23654]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1282
			// _ = "end of CoverTab[23654]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1283
			_go_fuzz_dep_.CoverTab[23655]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1283
			// _ = "end of CoverTab[23655]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1283
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1283
		// _ = "end of CoverTab[23651]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1283
		_go_fuzz_dep_.CoverTab[23652]++
									certLen := uint32(d[0])<<16 | uint32(d[1])<<8 | uint32(d[2])
									if uint32(len(d)) < 3+certLen {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1285
			_go_fuzz_dep_.CoverTab[23656]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1286
			// _ = "end of CoverTab[23656]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1287
			_go_fuzz_dep_.CoverTab[23657]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1287
			// _ = "end of CoverTab[23657]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1287
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1287
		// _ = "end of CoverTab[23652]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1287
		_go_fuzz_dep_.CoverTab[23653]++
									d = d[3+certLen:]
									certsLen -= 3 + certLen
									numCerts++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1290
		// _ = "end of CoverTab[23653]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1291
	// _ = "end of CoverTab[23644]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1291
	_go_fuzz_dep_.CoverTab[23645]++

								m.certificates = make([][]byte, numCerts)
								d = data[7:]
								for i := 0; i < numCerts; i++ {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1295
		_go_fuzz_dep_.CoverTab[23658]++
									certLen := uint32(d[0])<<16 | uint32(d[1])<<8 | uint32(d[2])
									m.certificates[i] = d[3 : 3+certLen]
									d = d[3+certLen:]
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1298
		// _ = "end of CoverTab[23658]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1299
	// _ = "end of CoverTab[23645]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1299
	_go_fuzz_dep_.CoverTab[23646]++

								return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1301
	// _ = "end of CoverTab[23646]"
}

type certificateMsgTLS13 struct {
	raw		[]byte
	certificate	Certificate
	ocspStapling	bool
	scts		bool
}

func (m *certificateMsgTLS13) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1311
	_go_fuzz_dep_.CoverTab[23659]++
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1312
		_go_fuzz_dep_.CoverTab[23662]++
									return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1313
		// _ = "end of CoverTab[23662]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1314
		_go_fuzz_dep_.CoverTab[23663]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1314
		// _ = "end of CoverTab[23663]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1314
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1314
	// _ = "end of CoverTab[23659]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1314
	_go_fuzz_dep_.CoverTab[23660]++

								var b cryptobyte.Builder
								b.AddUint8(typeCertificate)
								b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1318
		_go_fuzz_dep_.CoverTab[23664]++
									b.AddUint8(0)

									certificate := m.certificate
									if !m.ocspStapling {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1322
			_go_fuzz_dep_.CoverTab[23667]++
										certificate.OCSPStaple = nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1323
			// _ = "end of CoverTab[23667]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1324
			_go_fuzz_dep_.CoverTab[23668]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1324
			// _ = "end of CoverTab[23668]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1324
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1324
		// _ = "end of CoverTab[23664]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1324
		_go_fuzz_dep_.CoverTab[23665]++
									if !m.scts {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1325
			_go_fuzz_dep_.CoverTab[23669]++
										certificate.SignedCertificateTimestamps = nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1326
			// _ = "end of CoverTab[23669]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1327
			_go_fuzz_dep_.CoverTab[23670]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1327
			// _ = "end of CoverTab[23670]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1327
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1327
		// _ = "end of CoverTab[23665]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1327
		_go_fuzz_dep_.CoverTab[23666]++
									marshalCertificate(b, certificate)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1328
		// _ = "end of CoverTab[23666]"
	})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1329
	// _ = "end of CoverTab[23660]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1329
	_go_fuzz_dep_.CoverTab[23661]++

								var err error
								m.raw, err = b.Bytes()
								return m.raw, err
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1333
	// _ = "end of CoverTab[23661]"
}

func marshalCertificate(b *cryptobyte.Builder, certificate Certificate) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1336
	_go_fuzz_dep_.CoverTab[23671]++
								b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1337
		_go_fuzz_dep_.CoverTab[23672]++
									for i, cert := range certificate.Certificate {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1338
			_go_fuzz_dep_.CoverTab[23673]++
										b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1339
				_go_fuzz_dep_.CoverTab[23675]++
											b.AddBytes(cert)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1340
				// _ = "end of CoverTab[23675]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1341
			// _ = "end of CoverTab[23673]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1341
			_go_fuzz_dep_.CoverTab[23674]++
										b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1342
				_go_fuzz_dep_.CoverTab[23676]++
											if i > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1343
					_go_fuzz_dep_.CoverTab[23679]++

												return
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1345
					// _ = "end of CoverTab[23679]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1346
					_go_fuzz_dep_.CoverTab[23680]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1346
					// _ = "end of CoverTab[23680]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1346
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1346
				// _ = "end of CoverTab[23676]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1346
				_go_fuzz_dep_.CoverTab[23677]++
											if certificate.OCSPStaple != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1347
					_go_fuzz_dep_.CoverTab[23681]++
												b.AddUint16(extensionStatusRequest)
												b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1349
						_go_fuzz_dep_.CoverTab[23682]++
													b.AddUint8(statusTypeOCSP)
													b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1351
							_go_fuzz_dep_.CoverTab[23683]++
														b.AddBytes(certificate.OCSPStaple)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1352
							// _ = "end of CoverTab[23683]"
						})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1353
						// _ = "end of CoverTab[23682]"
					})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1354
					// _ = "end of CoverTab[23681]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1355
					_go_fuzz_dep_.CoverTab[23684]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1355
					// _ = "end of CoverTab[23684]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1355
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1355
				// _ = "end of CoverTab[23677]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1355
				_go_fuzz_dep_.CoverTab[23678]++
											if certificate.SignedCertificateTimestamps != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1356
					_go_fuzz_dep_.CoverTab[23685]++
												b.AddUint16(extensionSCT)
												b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1358
						_go_fuzz_dep_.CoverTab[23686]++
													b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1359
							_go_fuzz_dep_.CoverTab[23687]++
														for _, sct := range certificate.SignedCertificateTimestamps {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1360
								_go_fuzz_dep_.CoverTab[23688]++
															b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1361
									_go_fuzz_dep_.CoverTab[23689]++
																b.AddBytes(sct)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1362
									// _ = "end of CoverTab[23689]"
								})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1363
								// _ = "end of CoverTab[23688]"
							}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1364
							// _ = "end of CoverTab[23687]"
						})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1365
						// _ = "end of CoverTab[23686]"
					})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1366
					// _ = "end of CoverTab[23685]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1367
					_go_fuzz_dep_.CoverTab[23690]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1367
					// _ = "end of CoverTab[23690]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1367
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1367
				// _ = "end of CoverTab[23678]"
			})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1368
			// _ = "end of CoverTab[23674]"
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1369
		// _ = "end of CoverTab[23672]"
	})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1370
	// _ = "end of CoverTab[23671]"
}

func (m *certificateMsgTLS13) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1373
	_go_fuzz_dep_.CoverTab[23691]++
								*m = certificateMsgTLS13{raw: data}
								s := cryptobyte.String(data)

								var context cryptobyte.String
								if !s.Skip(4) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1378
		_go_fuzz_dep_.CoverTab[23693]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1378
		return !s.ReadUint8LengthPrefixed(&context)
									// _ = "end of CoverTab[23693]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1379
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1379
		_go_fuzz_dep_.CoverTab[23694]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1379
		return !context.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1379
		// _ = "end of CoverTab[23694]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1379
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1379
		_go_fuzz_dep_.CoverTab[23695]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1379
		return !unmarshalCertificate(&s, &m.certificate)
									// _ = "end of CoverTab[23695]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1380
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1380
		_go_fuzz_dep_.CoverTab[23696]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1380
		return !s.Empty()
									// _ = "end of CoverTab[23696]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1381
	}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1381
		_go_fuzz_dep_.CoverTab[23697]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1382
		// _ = "end of CoverTab[23697]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1383
		_go_fuzz_dep_.CoverTab[23698]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1383
		// _ = "end of CoverTab[23698]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1383
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1383
	// _ = "end of CoverTab[23691]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1383
	_go_fuzz_dep_.CoverTab[23692]++

								m.scts = m.certificate.SignedCertificateTimestamps != nil
								m.ocspStapling = m.certificate.OCSPStaple != nil

								return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1388
	// _ = "end of CoverTab[23692]"
}

func unmarshalCertificate(s *cryptobyte.String, certificate *Certificate) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1391
	_go_fuzz_dep_.CoverTab[23699]++
								var certList cryptobyte.String
								if !s.ReadUint24LengthPrefixed(&certList) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1393
		_go_fuzz_dep_.CoverTab[23702]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1394
		// _ = "end of CoverTab[23702]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1395
		_go_fuzz_dep_.CoverTab[23703]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1395
		// _ = "end of CoverTab[23703]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1395
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1395
	// _ = "end of CoverTab[23699]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1395
	_go_fuzz_dep_.CoverTab[23700]++
								for !certList.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1396
		_go_fuzz_dep_.CoverTab[23704]++
									var cert []byte
									var extensions cryptobyte.String
									if !readUint24LengthPrefixed(&certList, &cert) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1399
			_go_fuzz_dep_.CoverTab[23706]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1399
			return !certList.ReadUint16LengthPrefixed(&extensions)
										// _ = "end of CoverTab[23706]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1400
		}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1400
			_go_fuzz_dep_.CoverTab[23707]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1401
			// _ = "end of CoverTab[23707]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1402
			_go_fuzz_dep_.CoverTab[23708]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1402
			// _ = "end of CoverTab[23708]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1402
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1402
		// _ = "end of CoverTab[23704]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1402
		_go_fuzz_dep_.CoverTab[23705]++
									certificate.Certificate = append(certificate.Certificate, cert)
									for !extensions.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1404
			_go_fuzz_dep_.CoverTab[23709]++
										var extension uint16
										var extData cryptobyte.String
										if !extensions.ReadUint16(&extension) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1407
				_go_fuzz_dep_.CoverTab[23713]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1407
				return !extensions.ReadUint16LengthPrefixed(&extData)
											// _ = "end of CoverTab[23713]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1408
			}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1408
				_go_fuzz_dep_.CoverTab[23714]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1409
				// _ = "end of CoverTab[23714]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1410
				_go_fuzz_dep_.CoverTab[23715]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1410
				// _ = "end of CoverTab[23715]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1410
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1410
			// _ = "end of CoverTab[23709]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1410
			_go_fuzz_dep_.CoverTab[23710]++
										if len(certificate.Certificate) > 1 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1411
				_go_fuzz_dep_.CoverTab[23716]++

											continue
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1413
				// _ = "end of CoverTab[23716]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1414
				_go_fuzz_dep_.CoverTab[23717]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1414
				// _ = "end of CoverTab[23717]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1414
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1414
			// _ = "end of CoverTab[23710]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1414
			_go_fuzz_dep_.CoverTab[23711]++

										switch extension {
			case extensionStatusRequest:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1417
				_go_fuzz_dep_.CoverTab[23718]++
											var statusType uint8
											if !extData.ReadUint8(&statusType) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1419
					_go_fuzz_dep_.CoverTab[23722]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1419
					return statusType != statusTypeOCSP
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1419
					// _ = "end of CoverTab[23722]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1419
				}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1419
					_go_fuzz_dep_.CoverTab[23723]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1419
					return !readUint24LengthPrefixed(&extData, &certificate.OCSPStaple)
												// _ = "end of CoverTab[23723]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1420
				}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1420
					_go_fuzz_dep_.CoverTab[23724]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1420
					return len(certificate.OCSPStaple) == 0
												// _ = "end of CoverTab[23724]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1421
				}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1421
					_go_fuzz_dep_.CoverTab[23725]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1422
					// _ = "end of CoverTab[23725]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1423
					_go_fuzz_dep_.CoverTab[23726]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1423
					// _ = "end of CoverTab[23726]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1423
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1423
				// _ = "end of CoverTab[23718]"
			case extensionSCT:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1424
				_go_fuzz_dep_.CoverTab[23719]++
											var sctList cryptobyte.String
											if !extData.ReadUint16LengthPrefixed(&sctList) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1426
					_go_fuzz_dep_.CoverTab[23727]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1426
					return sctList.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1426
					// _ = "end of CoverTab[23727]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1426
				}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1426
					_go_fuzz_dep_.CoverTab[23728]++
												return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1427
					// _ = "end of CoverTab[23728]"
				} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1428
					_go_fuzz_dep_.CoverTab[23729]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1428
					// _ = "end of CoverTab[23729]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1428
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1428
				// _ = "end of CoverTab[23719]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1428
				_go_fuzz_dep_.CoverTab[23720]++
											for !sctList.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1429
					_go_fuzz_dep_.CoverTab[23730]++
												var sct []byte
												if !readUint16LengthPrefixed(&sctList, &sct) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1431
						_go_fuzz_dep_.CoverTab[23732]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1431
						return len(sct) == 0
													// _ = "end of CoverTab[23732]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1432
					}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1432
						_go_fuzz_dep_.CoverTab[23733]++
													return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1433
						// _ = "end of CoverTab[23733]"
					} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1434
						_go_fuzz_dep_.CoverTab[23734]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1434
						// _ = "end of CoverTab[23734]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1434
					}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1434
					// _ = "end of CoverTab[23730]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1434
					_go_fuzz_dep_.CoverTab[23731]++
												certificate.SignedCertificateTimestamps = append(
						certificate.SignedCertificateTimestamps, sct)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1436
					// _ = "end of CoverTab[23731]"
				}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1437
				// _ = "end of CoverTab[23720]"
			default:
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1438
				_go_fuzz_dep_.CoverTab[23721]++

											continue
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1440
				// _ = "end of CoverTab[23721]"
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1441
			// _ = "end of CoverTab[23711]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1441
			_go_fuzz_dep_.CoverTab[23712]++

										if !extData.Empty() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1443
				_go_fuzz_dep_.CoverTab[23735]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1444
				// _ = "end of CoverTab[23735]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1445
				_go_fuzz_dep_.CoverTab[23736]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1445
				// _ = "end of CoverTab[23736]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1445
			}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1445
			// _ = "end of CoverTab[23712]"
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1446
		// _ = "end of CoverTab[23705]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1447
	// _ = "end of CoverTab[23700]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1447
	_go_fuzz_dep_.CoverTab[23701]++
								return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1448
	// _ = "end of CoverTab[23701]"
}

type serverKeyExchangeMsg struct {
	raw	[]byte
	key	[]byte
}

func (m *serverKeyExchangeMsg) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1456
	_go_fuzz_dep_.CoverTab[23737]++
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1457
		_go_fuzz_dep_.CoverTab[23739]++
									return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1458
		// _ = "end of CoverTab[23739]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1459
		_go_fuzz_dep_.CoverTab[23740]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1459
		// _ = "end of CoverTab[23740]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1459
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1459
	// _ = "end of CoverTab[23737]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1459
	_go_fuzz_dep_.CoverTab[23738]++
								length := len(m.key)
								x := make([]byte, length+4)
								x[0] = typeServerKeyExchange
								x[1] = uint8(length >> 16)
								x[2] = uint8(length >> 8)
								x[3] = uint8(length)
								copy(x[4:], m.key)

								m.raw = x
								return x, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1469
	// _ = "end of CoverTab[23738]"
}

func (m *serverKeyExchangeMsg) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1472
	_go_fuzz_dep_.CoverTab[23741]++
								m.raw = data
								if len(data) < 4 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1474
		_go_fuzz_dep_.CoverTab[23743]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1475
		// _ = "end of CoverTab[23743]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1476
		_go_fuzz_dep_.CoverTab[23744]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1476
		// _ = "end of CoverTab[23744]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1476
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1476
	// _ = "end of CoverTab[23741]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1476
	_go_fuzz_dep_.CoverTab[23742]++
								m.key = data[4:]
								return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1478
	// _ = "end of CoverTab[23742]"
}

type certificateStatusMsg struct {
	raw		[]byte
	response	[]byte
}

func (m *certificateStatusMsg) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1486
	_go_fuzz_dep_.CoverTab[23745]++
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1487
		_go_fuzz_dep_.CoverTab[23748]++
									return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1488
		// _ = "end of CoverTab[23748]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1489
		_go_fuzz_dep_.CoverTab[23749]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1489
		// _ = "end of CoverTab[23749]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1489
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1489
	// _ = "end of CoverTab[23745]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1489
	_go_fuzz_dep_.CoverTab[23746]++

								var b cryptobyte.Builder
								b.AddUint8(typeCertificateStatus)
								b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1493
		_go_fuzz_dep_.CoverTab[23750]++
									b.AddUint8(statusTypeOCSP)
									b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1495
			_go_fuzz_dep_.CoverTab[23751]++
										b.AddBytes(m.response)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1496
			// _ = "end of CoverTab[23751]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1497
		// _ = "end of CoverTab[23750]"
	})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1498
	// _ = "end of CoverTab[23746]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1498
	_go_fuzz_dep_.CoverTab[23747]++

								var err error
								m.raw, err = b.Bytes()
								return m.raw, err
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1502
	// _ = "end of CoverTab[23747]"
}

func (m *certificateStatusMsg) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1505
	_go_fuzz_dep_.CoverTab[23752]++
								m.raw = data
								s := cryptobyte.String(data)

								var statusType uint8
								if !s.Skip(4) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1510
		_go_fuzz_dep_.CoverTab[23754]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1510
		return !s.ReadUint8(&statusType)
									// _ = "end of CoverTab[23754]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1511
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1511
		_go_fuzz_dep_.CoverTab[23755]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1511
		return statusType != statusTypeOCSP
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1511
		// _ = "end of CoverTab[23755]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1511
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1511
		_go_fuzz_dep_.CoverTab[23756]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1511
		return !readUint24LengthPrefixed(&s, &m.response)
									// _ = "end of CoverTab[23756]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1512
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1512
		_go_fuzz_dep_.CoverTab[23757]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1512
		return len(m.response) == 0
									// _ = "end of CoverTab[23757]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1513
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1513
		_go_fuzz_dep_.CoverTab[23758]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1513
		return !s.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1513
		// _ = "end of CoverTab[23758]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1513
	}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1513
		_go_fuzz_dep_.CoverTab[23759]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1514
		// _ = "end of CoverTab[23759]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1515
		_go_fuzz_dep_.CoverTab[23760]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1515
		// _ = "end of CoverTab[23760]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1515
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1515
	// _ = "end of CoverTab[23752]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1515
	_go_fuzz_dep_.CoverTab[23753]++
								return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1516
	// _ = "end of CoverTab[23753]"
}

type serverHelloDoneMsg struct{}

func (m *serverHelloDoneMsg) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1521
	_go_fuzz_dep_.CoverTab[23761]++
								x := make([]byte, 4)
								x[0] = typeServerHelloDone
								return x, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1524
	// _ = "end of CoverTab[23761]"
}

func (m *serverHelloDoneMsg) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1527
	_go_fuzz_dep_.CoverTab[23762]++
								return len(data) == 4
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1528
	// _ = "end of CoverTab[23762]"
}

type clientKeyExchangeMsg struct {
	raw		[]byte
	ciphertext	[]byte
}

func (m *clientKeyExchangeMsg) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1536
	_go_fuzz_dep_.CoverTab[23763]++
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1537
		_go_fuzz_dep_.CoverTab[23765]++
									return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1538
		// _ = "end of CoverTab[23765]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1539
		_go_fuzz_dep_.CoverTab[23766]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1539
		// _ = "end of CoverTab[23766]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1539
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1539
	// _ = "end of CoverTab[23763]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1539
	_go_fuzz_dep_.CoverTab[23764]++
								length := len(m.ciphertext)
								x := make([]byte, length+4)
								x[0] = typeClientKeyExchange
								x[1] = uint8(length >> 16)
								x[2] = uint8(length >> 8)
								x[3] = uint8(length)
								copy(x[4:], m.ciphertext)

								m.raw = x
								return x, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1549
	// _ = "end of CoverTab[23764]"
}

func (m *clientKeyExchangeMsg) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1552
	_go_fuzz_dep_.CoverTab[23767]++
								m.raw = data
								if len(data) < 4 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1554
		_go_fuzz_dep_.CoverTab[23770]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1555
		// _ = "end of CoverTab[23770]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1556
		_go_fuzz_dep_.CoverTab[23771]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1556
		// _ = "end of CoverTab[23771]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1556
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1556
	// _ = "end of CoverTab[23767]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1556
	_go_fuzz_dep_.CoverTab[23768]++
								l := int(data[1])<<16 | int(data[2])<<8 | int(data[3])
								if l != len(data)-4 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1558
		_go_fuzz_dep_.CoverTab[23772]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1559
		// _ = "end of CoverTab[23772]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1560
		_go_fuzz_dep_.CoverTab[23773]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1560
		// _ = "end of CoverTab[23773]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1560
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1560
	// _ = "end of CoverTab[23768]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1560
	_go_fuzz_dep_.CoverTab[23769]++
								m.ciphertext = data[4:]
								return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1562
	// _ = "end of CoverTab[23769]"
}

type finishedMsg struct {
	raw		[]byte
	verifyData	[]byte
}

func (m *finishedMsg) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1570
	_go_fuzz_dep_.CoverTab[23774]++
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1571
		_go_fuzz_dep_.CoverTab[23777]++
									return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1572
		// _ = "end of CoverTab[23777]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1573
		_go_fuzz_dep_.CoverTab[23778]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1573
		// _ = "end of CoverTab[23778]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1573
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1573
	// _ = "end of CoverTab[23774]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1573
	_go_fuzz_dep_.CoverTab[23775]++

								var b cryptobyte.Builder
								b.AddUint8(typeFinished)
								b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1577
		_go_fuzz_dep_.CoverTab[23779]++
									b.AddBytes(m.verifyData)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1578
		// _ = "end of CoverTab[23779]"
	})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1579
	// _ = "end of CoverTab[23775]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1579
	_go_fuzz_dep_.CoverTab[23776]++

								var err error
								m.raw, err = b.Bytes()
								return m.raw, err
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1583
	// _ = "end of CoverTab[23776]"
}

func (m *finishedMsg) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1586
	_go_fuzz_dep_.CoverTab[23780]++
								m.raw = data
								s := cryptobyte.String(data)
								return s.Skip(1) && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1589
		_go_fuzz_dep_.CoverTab[23781]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1589
		return readUint24LengthPrefixed(&s, &m.verifyData)
									// _ = "end of CoverTab[23781]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1590
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1590
		_go_fuzz_dep_.CoverTab[23782]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1590
		return s.Empty()
									// _ = "end of CoverTab[23782]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1591
	}()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1591
	// _ = "end of CoverTab[23780]"
}

type certificateRequestMsg struct {
	raw	[]byte
	// hasSignatureAlgorithm indicates whether this message includes a list of
	// supported signature algorithms. This change was introduced with TLS 1.2.
	hasSignatureAlgorithm	bool

	certificateTypes		[]byte
	supportedSignatureAlgorithms	[]SignatureScheme
	certificateAuthorities		[][]byte
}

func (m *certificateRequestMsg) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1605
	_go_fuzz_dep_.CoverTab[23783]++
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1606
		_go_fuzz_dep_.CoverTab[23789]++
									return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1607
		// _ = "end of CoverTab[23789]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1608
		_go_fuzz_dep_.CoverTab[23790]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1608
		// _ = "end of CoverTab[23790]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1608
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1608
	// _ = "end of CoverTab[23783]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1608
	_go_fuzz_dep_.CoverTab[23784]++

//line /usr/local/go/src/crypto/tls/handshake_messages.go:1611
	length := 1 + len(m.certificateTypes) + 2
	casLength := 0
	for _, ca := range m.certificateAuthorities {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1613
		_go_fuzz_dep_.CoverTab[23791]++
									casLength += 2 + len(ca)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1614
		// _ = "end of CoverTab[23791]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1615
	// _ = "end of CoverTab[23784]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1615
	_go_fuzz_dep_.CoverTab[23785]++
								length += casLength

								if m.hasSignatureAlgorithm {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1618
		_go_fuzz_dep_.CoverTab[23792]++
									length += 2 + 2*len(m.supportedSignatureAlgorithms)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1619
		// _ = "end of CoverTab[23792]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1620
		_go_fuzz_dep_.CoverTab[23793]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1620
		// _ = "end of CoverTab[23793]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1620
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1620
	// _ = "end of CoverTab[23785]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1620
	_go_fuzz_dep_.CoverTab[23786]++

								x := make([]byte, 4+length)
								x[0] = typeCertificateRequest
								x[1] = uint8(length >> 16)
								x[2] = uint8(length >> 8)
								x[3] = uint8(length)

								x[4] = uint8(len(m.certificateTypes))

								copy(x[5:], m.certificateTypes)
								y := x[5+len(m.certificateTypes):]

								if m.hasSignatureAlgorithm {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1633
		_go_fuzz_dep_.CoverTab[23794]++
									n := len(m.supportedSignatureAlgorithms) * 2
									y[0] = uint8(n >> 8)
									y[1] = uint8(n)
									y = y[2:]
									for _, sigAlgo := range m.supportedSignatureAlgorithms {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1638
			_go_fuzz_dep_.CoverTab[23795]++
										y[0] = uint8(sigAlgo >> 8)
										y[1] = uint8(sigAlgo)
										y = y[2:]
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1641
			// _ = "end of CoverTab[23795]"
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1642
		// _ = "end of CoverTab[23794]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1643
		_go_fuzz_dep_.CoverTab[23796]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1643
		// _ = "end of CoverTab[23796]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1643
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1643
	// _ = "end of CoverTab[23786]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1643
	_go_fuzz_dep_.CoverTab[23787]++

								y[0] = uint8(casLength >> 8)
								y[1] = uint8(casLength)
								y = y[2:]
								for _, ca := range m.certificateAuthorities {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1648
		_go_fuzz_dep_.CoverTab[23797]++
									y[0] = uint8(len(ca) >> 8)
									y[1] = uint8(len(ca))
									y = y[2:]
									copy(y, ca)
									y = y[len(ca):]
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1653
		// _ = "end of CoverTab[23797]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1654
	// _ = "end of CoverTab[23787]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1654
	_go_fuzz_dep_.CoverTab[23788]++

								m.raw = x
								return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1657
	// _ = "end of CoverTab[23788]"
}

func (m *certificateRequestMsg) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1660
	_go_fuzz_dep_.CoverTab[23798]++
								m.raw = data

								if len(data) < 5 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1663
		_go_fuzz_dep_.CoverTab[23807]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1664
		// _ = "end of CoverTab[23807]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1665
		_go_fuzz_dep_.CoverTab[23808]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1665
		// _ = "end of CoverTab[23808]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1665
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1665
	// _ = "end of CoverTab[23798]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1665
	_go_fuzz_dep_.CoverTab[23799]++

								length := uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])
								if uint32(len(data))-4 != length {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1668
		_go_fuzz_dep_.CoverTab[23809]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1669
		// _ = "end of CoverTab[23809]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1670
		_go_fuzz_dep_.CoverTab[23810]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1670
		// _ = "end of CoverTab[23810]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1670
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1670
	// _ = "end of CoverTab[23799]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1670
	_go_fuzz_dep_.CoverTab[23800]++

								numCertTypes := int(data[4])
								data = data[5:]
								if numCertTypes == 0 || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1674
		_go_fuzz_dep_.CoverTab[23811]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1674
		return len(data) <= numCertTypes
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1674
		// _ = "end of CoverTab[23811]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1674
	}() {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1674
		_go_fuzz_dep_.CoverTab[23812]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1675
		// _ = "end of CoverTab[23812]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1676
		_go_fuzz_dep_.CoverTab[23813]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1676
		// _ = "end of CoverTab[23813]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1676
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1676
	// _ = "end of CoverTab[23800]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1676
	_go_fuzz_dep_.CoverTab[23801]++

								m.certificateTypes = make([]byte, numCertTypes)
								if copy(m.certificateTypes, data) != numCertTypes {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1679
		_go_fuzz_dep_.CoverTab[23814]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1680
		// _ = "end of CoverTab[23814]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1681
		_go_fuzz_dep_.CoverTab[23815]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1681
		// _ = "end of CoverTab[23815]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1681
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1681
	// _ = "end of CoverTab[23801]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1681
	_go_fuzz_dep_.CoverTab[23802]++

								data = data[numCertTypes:]

								if m.hasSignatureAlgorithm {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1685
		_go_fuzz_dep_.CoverTab[23816]++
									if len(data) < 2 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1686
			_go_fuzz_dep_.CoverTab[23820]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1687
			// _ = "end of CoverTab[23820]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1688
			_go_fuzz_dep_.CoverTab[23821]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1688
			// _ = "end of CoverTab[23821]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1688
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1688
		// _ = "end of CoverTab[23816]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1688
		_go_fuzz_dep_.CoverTab[23817]++
									sigAndHashLen := uint16(data[0])<<8 | uint16(data[1])
									data = data[2:]
									if sigAndHashLen&1 != 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1691
			_go_fuzz_dep_.CoverTab[23822]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1692
			// _ = "end of CoverTab[23822]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1693
			_go_fuzz_dep_.CoverTab[23823]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1693
			// _ = "end of CoverTab[23823]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1693
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1693
		// _ = "end of CoverTab[23817]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1693
		_go_fuzz_dep_.CoverTab[23818]++
									if len(data) < int(sigAndHashLen) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1694
			_go_fuzz_dep_.CoverTab[23824]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1695
			// _ = "end of CoverTab[23824]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1696
			_go_fuzz_dep_.CoverTab[23825]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1696
			// _ = "end of CoverTab[23825]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1696
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1696
		// _ = "end of CoverTab[23818]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1696
		_go_fuzz_dep_.CoverTab[23819]++
									numSigAlgos := sigAndHashLen / 2
									m.supportedSignatureAlgorithms = make([]SignatureScheme, numSigAlgos)
									for i := range m.supportedSignatureAlgorithms {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1699
			_go_fuzz_dep_.CoverTab[23826]++
										m.supportedSignatureAlgorithms[i] = SignatureScheme(data[0])<<8 | SignatureScheme(data[1])
										data = data[2:]
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1701
			// _ = "end of CoverTab[23826]"
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1702
		// _ = "end of CoverTab[23819]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1703
		_go_fuzz_dep_.CoverTab[23827]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1703
		// _ = "end of CoverTab[23827]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1703
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1703
	// _ = "end of CoverTab[23802]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1703
	_go_fuzz_dep_.CoverTab[23803]++

								if len(data) < 2 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1705
		_go_fuzz_dep_.CoverTab[23828]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1706
		// _ = "end of CoverTab[23828]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1707
		_go_fuzz_dep_.CoverTab[23829]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1707
		// _ = "end of CoverTab[23829]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1707
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1707
	// _ = "end of CoverTab[23803]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1707
	_go_fuzz_dep_.CoverTab[23804]++
								casLength := uint16(data[0])<<8 | uint16(data[1])
								data = data[2:]
								if len(data) < int(casLength) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1710
		_go_fuzz_dep_.CoverTab[23830]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1711
		// _ = "end of CoverTab[23830]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1712
		_go_fuzz_dep_.CoverTab[23831]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1712
		// _ = "end of CoverTab[23831]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1712
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1712
	// _ = "end of CoverTab[23804]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1712
	_go_fuzz_dep_.CoverTab[23805]++
								cas := make([]byte, casLength)
								copy(cas, data)
								data = data[casLength:]

								m.certificateAuthorities = nil
								for len(cas) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1718
		_go_fuzz_dep_.CoverTab[23832]++
									if len(cas) < 2 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1719
			_go_fuzz_dep_.CoverTab[23835]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1720
			// _ = "end of CoverTab[23835]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1721
			_go_fuzz_dep_.CoverTab[23836]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1721
			// _ = "end of CoverTab[23836]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1721
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1721
		// _ = "end of CoverTab[23832]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1721
		_go_fuzz_dep_.CoverTab[23833]++
									caLen := uint16(cas[0])<<8 | uint16(cas[1])
									cas = cas[2:]

									if len(cas) < int(caLen) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1725
			_go_fuzz_dep_.CoverTab[23837]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1726
			// _ = "end of CoverTab[23837]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1727
			_go_fuzz_dep_.CoverTab[23838]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1727
			// _ = "end of CoverTab[23838]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1727
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1727
		// _ = "end of CoverTab[23833]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1727
		_go_fuzz_dep_.CoverTab[23834]++

									m.certificateAuthorities = append(m.certificateAuthorities, cas[:caLen])
									cas = cas[caLen:]
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1730
		// _ = "end of CoverTab[23834]"
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1731
	// _ = "end of CoverTab[23805]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1731
	_go_fuzz_dep_.CoverTab[23806]++

								return len(data) == 0
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1733
	// _ = "end of CoverTab[23806]"
}

type certificateVerifyMsg struct {
	raw			[]byte
	hasSignatureAlgorithm	bool	// format change introduced in TLS 1.2
	signatureAlgorithm	SignatureScheme
	signature		[]byte
}

func (m *certificateVerifyMsg) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1743
	_go_fuzz_dep_.CoverTab[23839]++
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1744
		_go_fuzz_dep_.CoverTab[23842]++
									return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1745
		// _ = "end of CoverTab[23842]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1746
		_go_fuzz_dep_.CoverTab[23843]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1746
		// _ = "end of CoverTab[23843]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1746
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1746
	// _ = "end of CoverTab[23839]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1746
	_go_fuzz_dep_.CoverTab[23840]++

								var b cryptobyte.Builder
								b.AddUint8(typeCertificateVerify)
								b.AddUint24LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1750
		_go_fuzz_dep_.CoverTab[23844]++
									if m.hasSignatureAlgorithm {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1751
			_go_fuzz_dep_.CoverTab[23846]++
										b.AddUint16(uint16(m.signatureAlgorithm))
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1752
			// _ = "end of CoverTab[23846]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1753
			_go_fuzz_dep_.CoverTab[23847]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1753
			// _ = "end of CoverTab[23847]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1753
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1753
		// _ = "end of CoverTab[23844]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1753
		_go_fuzz_dep_.CoverTab[23845]++
									b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1754
			_go_fuzz_dep_.CoverTab[23848]++
										b.AddBytes(m.signature)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1755
			// _ = "end of CoverTab[23848]"
		})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1756
		// _ = "end of CoverTab[23845]"
	})
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1757
	// _ = "end of CoverTab[23840]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1757
	_go_fuzz_dep_.CoverTab[23841]++

								var err error
								m.raw, err = b.Bytes()
								return m.raw, err
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1761
	// _ = "end of CoverTab[23841]"
}

func (m *certificateVerifyMsg) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1764
	_go_fuzz_dep_.CoverTab[23849]++
								m.raw = data
								s := cryptobyte.String(data)

								if !s.Skip(4) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1768
		_go_fuzz_dep_.CoverTab[23852]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1769
		// _ = "end of CoverTab[23852]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1770
		_go_fuzz_dep_.CoverTab[23853]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1770
		// _ = "end of CoverTab[23853]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1770
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1770
	// _ = "end of CoverTab[23849]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1770
	_go_fuzz_dep_.CoverTab[23850]++
								if m.hasSignatureAlgorithm {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1771
		_go_fuzz_dep_.CoverTab[23854]++
									if !s.ReadUint16((*uint16)(&m.signatureAlgorithm)) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1772
			_go_fuzz_dep_.CoverTab[23855]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1773
			// _ = "end of CoverTab[23855]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1774
			_go_fuzz_dep_.CoverTab[23856]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1774
			// _ = "end of CoverTab[23856]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1774
		}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1774
		// _ = "end of CoverTab[23854]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1775
		_go_fuzz_dep_.CoverTab[23857]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1775
		// _ = "end of CoverTab[23857]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1775
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1775
	// _ = "end of CoverTab[23850]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1775
	_go_fuzz_dep_.CoverTab[23851]++
								return readUint16LengthPrefixed(&s, &m.signature) && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1776
		_go_fuzz_dep_.CoverTab[23858]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1776
		return s.Empty()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1776
		// _ = "end of CoverTab[23858]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1776
	}()
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1776
	// _ = "end of CoverTab[23851]"
}

type newSessionTicketMsg struct {
	raw	[]byte
	ticket	[]byte
}

func (m *newSessionTicketMsg) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1784
	_go_fuzz_dep_.CoverTab[23859]++
								if m.raw != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1785
		_go_fuzz_dep_.CoverTab[23861]++
									return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1786
		// _ = "end of CoverTab[23861]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1787
		_go_fuzz_dep_.CoverTab[23862]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1787
		// _ = "end of CoverTab[23862]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1787
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1787
	// _ = "end of CoverTab[23859]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1787
	_go_fuzz_dep_.CoverTab[23860]++

//line /usr/local/go/src/crypto/tls/handshake_messages.go:1790
	ticketLen := len(m.ticket)
								length := 2 + 4 + ticketLen
								x := make([]byte, 4+length)
								x[0] = typeNewSessionTicket
								x[1] = uint8(length >> 16)
								x[2] = uint8(length >> 8)
								x[3] = uint8(length)
								x[8] = uint8(ticketLen >> 8)
								x[9] = uint8(ticketLen)
								copy(x[10:], m.ticket)

								m.raw = x

								return m.raw, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1803
	// _ = "end of CoverTab[23860]"
}

func (m *newSessionTicketMsg) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1806
	_go_fuzz_dep_.CoverTab[23863]++
								m.raw = data

								if len(data) < 10 {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1809
		_go_fuzz_dep_.CoverTab[23867]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1810
		// _ = "end of CoverTab[23867]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1811
		_go_fuzz_dep_.CoverTab[23868]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1811
		// _ = "end of CoverTab[23868]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1811
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1811
	// _ = "end of CoverTab[23863]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1811
	_go_fuzz_dep_.CoverTab[23864]++

								length := uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])
								if uint32(len(data))-4 != length {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1814
		_go_fuzz_dep_.CoverTab[23869]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1815
		// _ = "end of CoverTab[23869]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1816
		_go_fuzz_dep_.CoverTab[23870]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1816
		// _ = "end of CoverTab[23870]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1816
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1816
	// _ = "end of CoverTab[23864]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1816
	_go_fuzz_dep_.CoverTab[23865]++

								ticketLen := int(data[8])<<8 + int(data[9])
								if len(data)-10 != ticketLen {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1819
		_go_fuzz_dep_.CoverTab[23871]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1820
		// _ = "end of CoverTab[23871]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1821
		_go_fuzz_dep_.CoverTab[23872]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1821
		// _ = "end of CoverTab[23872]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1821
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1821
	// _ = "end of CoverTab[23865]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1821
	_go_fuzz_dep_.CoverTab[23866]++

								m.ticket = data[10:]

								return true
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1825
	// _ = "end of CoverTab[23866]"
}

type helloRequestMsg struct {
}

func (*helloRequestMsg) marshal() ([]byte, error) {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1831
	_go_fuzz_dep_.CoverTab[23873]++
								return []byte{typeHelloRequest, 0, 0, 0}, nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1832
	// _ = "end of CoverTab[23873]"
}

func (*helloRequestMsg) unmarshal(data []byte) bool {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1835
	_go_fuzz_dep_.CoverTab[23874]++
								return len(data) == 4
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1836
	// _ = "end of CoverTab[23874]"
}

type transcriptHash interface {
	Write([]byte) (int, error)
}

// transcriptMsg is a helper used to marshal and hash messages which typically
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1843
// are not written to the wire, and as such aren't hashed during Conn.writeRecord.
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1845
func transcriptMsg(msg handshakeMessage, h transcriptHash) error {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1845
	_go_fuzz_dep_.CoverTab[23875]++
								data, err := msg.marshal()
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1847
		_go_fuzz_dep_.CoverTab[23877]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1848
		// _ = "end of CoverTab[23877]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1849
		_go_fuzz_dep_.CoverTab[23878]++
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1849
		// _ = "end of CoverTab[23878]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1849
	}
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1849
	// _ = "end of CoverTab[23875]"
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1849
	_go_fuzz_dep_.CoverTab[23876]++
								h.Write(data)
								return nil
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1851
	// _ = "end of CoverTab[23876]"
}

//line /usr/local/go/src/crypto/tls/handshake_messages.go:1852
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/handshake_messages.go:1852
var _ = _go_fuzz_dep_.CoverTab
