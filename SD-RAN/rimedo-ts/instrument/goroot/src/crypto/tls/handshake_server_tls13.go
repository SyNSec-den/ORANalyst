// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:5
package tls

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:5
import (
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:5
)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:5
import (
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:5
)

import (
	"bytes"
	"context"
	"crypto"
	"crypto/hmac"
	"crypto/rsa"
	"encoding/binary"
	"errors"
	"hash"
	"io"
	"time"
)

// maxClientPSKIdentities is the number of client PSK identities the server will
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:20
// attempt to validate. It will ignore the rest not to let cheap ClientHello
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:20
// messages cause too much work in session ticket decryption attempts.
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:23
const maxClientPSKIdentities = 5

type serverHandshakeStateTLS13 struct {
	c		*Conn
	ctx		context.Context
	clientHello	*clientHelloMsg
	hello		*serverHelloMsg
	sentDummyCCS	bool
	usingPSK	bool
	suite		*cipherSuiteTLS13
	cert		*Certificate
	sigAlg		SignatureScheme
	earlySecret	[]byte
	sharedKey	[]byte
	handshakeSecret	[]byte
	masterSecret	[]byte
	trafficSecret	[]byte	// client_application_traffic_secret_0
	transcript	hash.Hash
	clientFinished	[]byte
}

func (hs *serverHandshakeStateTLS13) handshake() error {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:44
	_go_fuzz_dep_.CoverTab[24306]++
									c := hs.c

									if needFIPS() {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:47
		_go_fuzz_dep_.CoverTab[24317]++
										return errors.New("tls: internal error: TLS 1.3 reached in FIPS mode")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:48
		// _ = "end of CoverTab[24317]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:49
		_go_fuzz_dep_.CoverTab[24318]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:49
		// _ = "end of CoverTab[24318]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:49
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:49
	// _ = "end of CoverTab[24306]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:49
	_go_fuzz_dep_.CoverTab[24307]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:52
	if err := hs.processClientHello(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:52
		_go_fuzz_dep_.CoverTab[24319]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:53
		// _ = "end of CoverTab[24319]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:54
		_go_fuzz_dep_.CoverTab[24320]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:54
		// _ = "end of CoverTab[24320]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:54
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:54
	// _ = "end of CoverTab[24307]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:54
	_go_fuzz_dep_.CoverTab[24308]++
									if err := hs.checkForResumption(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:55
		_go_fuzz_dep_.CoverTab[24321]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:56
		// _ = "end of CoverTab[24321]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:57
		_go_fuzz_dep_.CoverTab[24322]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:57
		// _ = "end of CoverTab[24322]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:57
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:57
	// _ = "end of CoverTab[24308]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:57
	_go_fuzz_dep_.CoverTab[24309]++
									if err := hs.pickCertificate(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:58
		_go_fuzz_dep_.CoverTab[24323]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:59
		// _ = "end of CoverTab[24323]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:60
		_go_fuzz_dep_.CoverTab[24324]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:60
		// _ = "end of CoverTab[24324]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:60
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:60
	// _ = "end of CoverTab[24309]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:60
	_go_fuzz_dep_.CoverTab[24310]++
									c.buffering = true
									if err := hs.sendServerParameters(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:62
		_go_fuzz_dep_.CoverTab[24325]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:63
		// _ = "end of CoverTab[24325]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:64
		_go_fuzz_dep_.CoverTab[24326]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:64
		// _ = "end of CoverTab[24326]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:64
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:64
	// _ = "end of CoverTab[24310]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:64
	_go_fuzz_dep_.CoverTab[24311]++
									if err := hs.sendServerCertificate(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:65
		_go_fuzz_dep_.CoverTab[24327]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:66
		// _ = "end of CoverTab[24327]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:67
		_go_fuzz_dep_.CoverTab[24328]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:67
		// _ = "end of CoverTab[24328]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:67
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:67
	// _ = "end of CoverTab[24311]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:67
	_go_fuzz_dep_.CoverTab[24312]++
									if err := hs.sendServerFinished(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:68
		_go_fuzz_dep_.CoverTab[24329]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:69
		// _ = "end of CoverTab[24329]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:70
		_go_fuzz_dep_.CoverTab[24330]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:70
		// _ = "end of CoverTab[24330]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:70
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:70
	// _ = "end of CoverTab[24312]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:70
	_go_fuzz_dep_.CoverTab[24313]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:74
	if _, err := c.flush(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:74
		_go_fuzz_dep_.CoverTab[24331]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:75
		// _ = "end of CoverTab[24331]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:76
		_go_fuzz_dep_.CoverTab[24332]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:76
		// _ = "end of CoverTab[24332]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:76
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:76
	// _ = "end of CoverTab[24313]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:76
	_go_fuzz_dep_.CoverTab[24314]++
									if err := hs.readClientCertificate(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:77
		_go_fuzz_dep_.CoverTab[24333]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:78
		// _ = "end of CoverTab[24333]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:79
		_go_fuzz_dep_.CoverTab[24334]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:79
		// _ = "end of CoverTab[24334]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:79
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:79
	// _ = "end of CoverTab[24314]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:79
	_go_fuzz_dep_.CoverTab[24315]++
									if err := hs.readClientFinished(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:80
		_go_fuzz_dep_.CoverTab[24335]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:81
		// _ = "end of CoverTab[24335]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:82
		_go_fuzz_dep_.CoverTab[24336]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:82
		// _ = "end of CoverTab[24336]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:82
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:82
	// _ = "end of CoverTab[24315]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:82
	_go_fuzz_dep_.CoverTab[24316]++

									c.isHandshakeComplete.Store(true)

									return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:86
	// _ = "end of CoverTab[24316]"
}

func (hs *serverHandshakeStateTLS13) processClientHello() error {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:89
	_go_fuzz_dep_.CoverTab[24337]++
									c := hs.c

									hs.hello = new(serverHelloMsg)

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:96
	hs.hello.vers = VersionTLS12
	hs.hello.supportedVersion = c.vers

	if len(hs.clientHello.supportedVersions) == 0 {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:99
		_go_fuzz_dep_.CoverTab[24354]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: client used the legacy version field to negotiate TLS 1.3")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:101
		// _ = "end of CoverTab[24354]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:102
		_go_fuzz_dep_.CoverTab[24355]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:102
		// _ = "end of CoverTab[24355]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:102
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:102
	// _ = "end of CoverTab[24337]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:102
	_go_fuzz_dep_.CoverTab[24338]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:113
	for _, id := range hs.clientHello.cipherSuites {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:113
		_go_fuzz_dep_.CoverTab[24356]++
										if id == TLS_FALLBACK_SCSV {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:114
			_go_fuzz_dep_.CoverTab[24357]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:117
			if c.vers < c.config.maxSupportedVersion(roleServer) {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:117
				_go_fuzz_dep_.CoverTab[24359]++
												c.sendAlert(alertInappropriateFallback)
												return errors.New("tls: client using inappropriate protocol fallback")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:119
				// _ = "end of CoverTab[24359]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:120
				_go_fuzz_dep_.CoverTab[24360]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:120
				// _ = "end of CoverTab[24360]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:120
			}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:120
			// _ = "end of CoverTab[24357]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:120
			_go_fuzz_dep_.CoverTab[24358]++
											break
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:121
			// _ = "end of CoverTab[24358]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:122
			_go_fuzz_dep_.CoverTab[24361]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:122
			// _ = "end of CoverTab[24361]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:122
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:122
		// _ = "end of CoverTab[24356]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:123
	// _ = "end of CoverTab[24338]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:123
	_go_fuzz_dep_.CoverTab[24339]++

									if len(hs.clientHello.compressionMethods) != 1 || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:125
		_go_fuzz_dep_.CoverTab[24362]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:125
		return hs.clientHello.compressionMethods[0] != compressionNone
										// _ = "end of CoverTab[24362]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:126
	}() {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:126
		_go_fuzz_dep_.CoverTab[24363]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: TLS 1.3 client supports illegal compression methods")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:128
		// _ = "end of CoverTab[24363]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:129
		_go_fuzz_dep_.CoverTab[24364]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:129
		// _ = "end of CoverTab[24364]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:129
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:129
	// _ = "end of CoverTab[24339]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:129
	_go_fuzz_dep_.CoverTab[24340]++

									hs.hello.random = make([]byte, 32)
									if _, err := io.ReadFull(c.config.rand(), hs.hello.random); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:132
		_go_fuzz_dep_.CoverTab[24365]++
										c.sendAlert(alertInternalError)
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:134
		// _ = "end of CoverTab[24365]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:135
		_go_fuzz_dep_.CoverTab[24366]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:135
		// _ = "end of CoverTab[24366]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:135
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:135
	// _ = "end of CoverTab[24340]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:135
	_go_fuzz_dep_.CoverTab[24341]++

									if len(hs.clientHello.secureRenegotiation) != 0 {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:137
		_go_fuzz_dep_.CoverTab[24367]++
										c.sendAlert(alertHandshakeFailure)
										return errors.New("tls: initial handshake had non-empty renegotiation extension")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:139
		// _ = "end of CoverTab[24367]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:140
		_go_fuzz_dep_.CoverTab[24368]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:140
		// _ = "end of CoverTab[24368]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:140
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:140
	// _ = "end of CoverTab[24341]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:140
	_go_fuzz_dep_.CoverTab[24342]++

									if hs.clientHello.earlyData {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:142
		_go_fuzz_dep_.CoverTab[24369]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:149
		c.sendAlert(alertUnsupportedExtension)
										return errors.New("tls: client sent unexpected early data")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:150
		// _ = "end of CoverTab[24369]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:151
		_go_fuzz_dep_.CoverTab[24370]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:151
		// _ = "end of CoverTab[24370]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:151
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:151
	// _ = "end of CoverTab[24342]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:151
	_go_fuzz_dep_.CoverTab[24343]++

									hs.hello.sessionId = hs.clientHello.sessionId
									hs.hello.compressionMethod = compressionNone

									preferenceList := defaultCipherSuitesTLS13
									if !hasAESGCMHardwareSupport || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:157
		_go_fuzz_dep_.CoverTab[24371]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:157
		return !aesgcmPreferred(hs.clientHello.cipherSuites)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:157
		// _ = "end of CoverTab[24371]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:157
	}() {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:157
		_go_fuzz_dep_.CoverTab[24372]++
										preferenceList = defaultCipherSuitesTLS13NoAES
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:158
		// _ = "end of CoverTab[24372]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:159
		_go_fuzz_dep_.CoverTab[24373]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:159
		// _ = "end of CoverTab[24373]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:159
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:159
	// _ = "end of CoverTab[24343]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:159
	_go_fuzz_dep_.CoverTab[24344]++
									for _, suiteID := range preferenceList {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:160
		_go_fuzz_dep_.CoverTab[24374]++
										hs.suite = mutualCipherSuiteTLS13(hs.clientHello.cipherSuites, suiteID)
										if hs.suite != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:162
			_go_fuzz_dep_.CoverTab[24375]++
											break
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:163
			// _ = "end of CoverTab[24375]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:164
			_go_fuzz_dep_.CoverTab[24376]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:164
			// _ = "end of CoverTab[24376]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:164
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:164
		// _ = "end of CoverTab[24374]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:165
	// _ = "end of CoverTab[24344]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:165
	_go_fuzz_dep_.CoverTab[24345]++
									if hs.suite == nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:166
		_go_fuzz_dep_.CoverTab[24377]++
										c.sendAlert(alertHandshakeFailure)
										return errors.New("tls: no cipher suite supported by both client and server")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:168
		// _ = "end of CoverTab[24377]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:169
		_go_fuzz_dep_.CoverTab[24378]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:169
		// _ = "end of CoverTab[24378]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:169
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:169
	// _ = "end of CoverTab[24345]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:169
	_go_fuzz_dep_.CoverTab[24346]++
									c.cipherSuite = hs.suite.id
									hs.hello.cipherSuite = hs.suite.id
									hs.transcript = hs.suite.hash.New()

	// Pick the ECDHE group in server preference order, but give priority to
	// groups with a key share, to avoid a HelloRetryRequest round-trip.
	var selectedGroup CurveID
	var clientKeyShare *keyShare
GroupSelection:
	for _, preferredGroup := range c.config.curvePreferences() {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:179
		_go_fuzz_dep_.CoverTab[24379]++
										for _, ks := range hs.clientHello.keyShares {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:180
			_go_fuzz_dep_.CoverTab[24382]++
											if ks.group == preferredGroup {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:181
				_go_fuzz_dep_.CoverTab[24383]++
												selectedGroup = ks.group
												clientKeyShare = &ks
												break GroupSelection
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:184
				// _ = "end of CoverTab[24383]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:185
				_go_fuzz_dep_.CoverTab[24384]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:185
				// _ = "end of CoverTab[24384]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:185
			}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:185
			// _ = "end of CoverTab[24382]"
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:186
		// _ = "end of CoverTab[24379]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:186
		_go_fuzz_dep_.CoverTab[24380]++
										if selectedGroup != 0 {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:187
			_go_fuzz_dep_.CoverTab[24385]++
											continue
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:188
			// _ = "end of CoverTab[24385]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:189
			_go_fuzz_dep_.CoverTab[24386]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:189
			// _ = "end of CoverTab[24386]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:189
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:189
		// _ = "end of CoverTab[24380]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:189
		_go_fuzz_dep_.CoverTab[24381]++
										for _, group := range hs.clientHello.supportedCurves {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:190
			_go_fuzz_dep_.CoverTab[24387]++
											if group == preferredGroup {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:191
				_go_fuzz_dep_.CoverTab[24388]++
												selectedGroup = group
												break
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:193
				// _ = "end of CoverTab[24388]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:194
				_go_fuzz_dep_.CoverTab[24389]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:194
				// _ = "end of CoverTab[24389]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:194
			}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:194
			// _ = "end of CoverTab[24387]"
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:195
		// _ = "end of CoverTab[24381]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:196
	// _ = "end of CoverTab[24346]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:196
	_go_fuzz_dep_.CoverTab[24347]++
									if selectedGroup == 0 {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:197
		_go_fuzz_dep_.CoverTab[24390]++
										c.sendAlert(alertHandshakeFailure)
										return errors.New("tls: no ECDHE curve supported by both client and server")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:199
		// _ = "end of CoverTab[24390]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:200
		_go_fuzz_dep_.CoverTab[24391]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:200
		// _ = "end of CoverTab[24391]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:200
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:200
	// _ = "end of CoverTab[24347]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:200
	_go_fuzz_dep_.CoverTab[24348]++
									if clientKeyShare == nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:201
		_go_fuzz_dep_.CoverTab[24392]++
										if err := hs.doHelloRetryRequest(selectedGroup); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:202
			_go_fuzz_dep_.CoverTab[24394]++
											return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:203
			// _ = "end of CoverTab[24394]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:204
			_go_fuzz_dep_.CoverTab[24395]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:204
			// _ = "end of CoverTab[24395]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:204
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:204
		// _ = "end of CoverTab[24392]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:204
		_go_fuzz_dep_.CoverTab[24393]++
										clientKeyShare = &hs.clientHello.keyShares[0]
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:205
		// _ = "end of CoverTab[24393]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:206
		_go_fuzz_dep_.CoverTab[24396]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:206
		// _ = "end of CoverTab[24396]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:206
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:206
	// _ = "end of CoverTab[24348]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:206
	_go_fuzz_dep_.CoverTab[24349]++

									if _, ok := curveForCurveID(selectedGroup); !ok {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:208
		_go_fuzz_dep_.CoverTab[24397]++
										c.sendAlert(alertInternalError)
										return errors.New("tls: CurvePreferences includes unsupported curve")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:210
		// _ = "end of CoverTab[24397]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:211
		_go_fuzz_dep_.CoverTab[24398]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:211
		// _ = "end of CoverTab[24398]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:211
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:211
	// _ = "end of CoverTab[24349]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:211
	_go_fuzz_dep_.CoverTab[24350]++
									key, err := generateECDHEKey(c.config.rand(), selectedGroup)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:213
		_go_fuzz_dep_.CoverTab[24399]++
										c.sendAlert(alertInternalError)
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:215
		// _ = "end of CoverTab[24399]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:216
		_go_fuzz_dep_.CoverTab[24400]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:216
		// _ = "end of CoverTab[24400]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:216
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:216
	// _ = "end of CoverTab[24350]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:216
	_go_fuzz_dep_.CoverTab[24351]++
									hs.hello.serverShare = keyShare{group: selectedGroup, data: key.PublicKey().Bytes()}
									peerKey, err := key.Curve().NewPublicKey(clientKeyShare.data)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:219
		_go_fuzz_dep_.CoverTab[24401]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: invalid client key share")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:221
		// _ = "end of CoverTab[24401]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:222
		_go_fuzz_dep_.CoverTab[24402]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:222
		// _ = "end of CoverTab[24402]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:222
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:222
	// _ = "end of CoverTab[24351]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:222
	_go_fuzz_dep_.CoverTab[24352]++
									hs.sharedKey, err = key.ECDH(peerKey)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:224
		_go_fuzz_dep_.CoverTab[24403]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: invalid client key share")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:226
		// _ = "end of CoverTab[24403]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:227
		_go_fuzz_dep_.CoverTab[24404]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:227
		// _ = "end of CoverTab[24404]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:227
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:227
	// _ = "end of CoverTab[24352]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:227
	_go_fuzz_dep_.CoverTab[24353]++

									c.serverName = hs.clientHello.serverName
									return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:230
	// _ = "end of CoverTab[24353]"
}

func (hs *serverHandshakeStateTLS13) checkForResumption() error {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:233
	_go_fuzz_dep_.CoverTab[24405]++
									c := hs.c

									if c.config.SessionTicketsDisabled {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:236
		_go_fuzz_dep_.CoverTab[24412]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:237
		// _ = "end of CoverTab[24412]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:238
		_go_fuzz_dep_.CoverTab[24413]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:238
		// _ = "end of CoverTab[24413]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:238
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:238
	// _ = "end of CoverTab[24405]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:238
	_go_fuzz_dep_.CoverTab[24406]++

									modeOK := false
									for _, mode := range hs.clientHello.pskModes {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:241
		_go_fuzz_dep_.CoverTab[24414]++
										if mode == pskModeDHE {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:242
			_go_fuzz_dep_.CoverTab[24415]++
											modeOK = true
											break
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:244
			// _ = "end of CoverTab[24415]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:245
			_go_fuzz_dep_.CoverTab[24416]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:245
			// _ = "end of CoverTab[24416]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:245
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:245
		// _ = "end of CoverTab[24414]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:246
	// _ = "end of CoverTab[24406]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:246
	_go_fuzz_dep_.CoverTab[24407]++
									if !modeOK {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:247
		_go_fuzz_dep_.CoverTab[24417]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:248
		// _ = "end of CoverTab[24417]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:249
		_go_fuzz_dep_.CoverTab[24418]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:249
		// _ = "end of CoverTab[24418]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:249
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:249
	// _ = "end of CoverTab[24407]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:249
	_go_fuzz_dep_.CoverTab[24408]++

									if len(hs.clientHello.pskIdentities) != len(hs.clientHello.pskBinders) {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:251
		_go_fuzz_dep_.CoverTab[24419]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: invalid or missing PSK binders")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:253
		// _ = "end of CoverTab[24419]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:254
		_go_fuzz_dep_.CoverTab[24420]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:254
		// _ = "end of CoverTab[24420]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:254
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:254
	// _ = "end of CoverTab[24408]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:254
	_go_fuzz_dep_.CoverTab[24409]++
									if len(hs.clientHello.pskIdentities) == 0 {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:255
		_go_fuzz_dep_.CoverTab[24421]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:256
		// _ = "end of CoverTab[24421]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:257
		_go_fuzz_dep_.CoverTab[24422]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:257
		// _ = "end of CoverTab[24422]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:257
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:257
	// _ = "end of CoverTab[24409]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:257
	_go_fuzz_dep_.CoverTab[24410]++

									for i, identity := range hs.clientHello.pskIdentities {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:259
		_go_fuzz_dep_.CoverTab[24423]++
										if i >= maxClientPSKIdentities {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:260
			_go_fuzz_dep_.CoverTab[24435]++
											break
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:261
			// _ = "end of CoverTab[24435]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:262
			_go_fuzz_dep_.CoverTab[24436]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:262
			// _ = "end of CoverTab[24436]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:262
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:262
		// _ = "end of CoverTab[24423]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:262
		_go_fuzz_dep_.CoverTab[24424]++

										plaintext, _ := c.decryptTicket(identity.label)
										if plaintext == nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:265
			_go_fuzz_dep_.CoverTab[24437]++
											continue
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:266
			// _ = "end of CoverTab[24437]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:267
			_go_fuzz_dep_.CoverTab[24438]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:267
			// _ = "end of CoverTab[24438]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:267
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:267
		// _ = "end of CoverTab[24424]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:267
		_go_fuzz_dep_.CoverTab[24425]++
										sessionState := new(sessionStateTLS13)
										if ok := sessionState.unmarshal(plaintext); !ok {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:269
			_go_fuzz_dep_.CoverTab[24439]++
											continue
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:270
			// _ = "end of CoverTab[24439]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:271
			_go_fuzz_dep_.CoverTab[24440]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:271
			// _ = "end of CoverTab[24440]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:271
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:271
		// _ = "end of CoverTab[24425]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:271
		_go_fuzz_dep_.CoverTab[24426]++

										createdAt := time.Unix(int64(sessionState.createdAt), 0)
										if c.config.time().Sub(createdAt) > maxSessionTicketLifetime {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:274
			_go_fuzz_dep_.CoverTab[24441]++
											continue
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:275
			// _ = "end of CoverTab[24441]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:276
			_go_fuzz_dep_.CoverTab[24442]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:276
			// _ = "end of CoverTab[24442]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:276
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:276
		// _ = "end of CoverTab[24426]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:276
		_go_fuzz_dep_.CoverTab[24427]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:282
		pskSuite := cipherSuiteTLS13ByID(sessionState.cipherSuite)
		if pskSuite == nil || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:283
			_go_fuzz_dep_.CoverTab[24443]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:283
			return pskSuite.hash != hs.suite.hash
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:283
			// _ = "end of CoverTab[24443]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:283
		}() {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:283
			_go_fuzz_dep_.CoverTab[24444]++
											continue
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:284
			// _ = "end of CoverTab[24444]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:285
			_go_fuzz_dep_.CoverTab[24445]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:285
			// _ = "end of CoverTab[24445]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:285
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:285
		// _ = "end of CoverTab[24427]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:285
		_go_fuzz_dep_.CoverTab[24428]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:290
		sessionHasClientCerts := len(sessionState.certificate.Certificate) != 0
		needClientCerts := requiresClientCert(c.config.ClientAuth)
		if needClientCerts && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:292
			_go_fuzz_dep_.CoverTab[24446]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:292
			return !sessionHasClientCerts
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:292
			// _ = "end of CoverTab[24446]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:292
		}() {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:292
			_go_fuzz_dep_.CoverTab[24447]++
											continue
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:293
			// _ = "end of CoverTab[24447]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:294
			_go_fuzz_dep_.CoverTab[24448]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:294
			// _ = "end of CoverTab[24448]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:294
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:294
		// _ = "end of CoverTab[24428]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:294
		_go_fuzz_dep_.CoverTab[24429]++
										if sessionHasClientCerts && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:295
			_go_fuzz_dep_.CoverTab[24449]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:295
			return c.config.ClientAuth == NoClientCert
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:295
			// _ = "end of CoverTab[24449]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:295
		}() {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:295
			_go_fuzz_dep_.CoverTab[24450]++
											continue
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:296
			// _ = "end of CoverTab[24450]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:297
			_go_fuzz_dep_.CoverTab[24451]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:297
			// _ = "end of CoverTab[24451]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:297
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:297
		// _ = "end of CoverTab[24429]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:297
		_go_fuzz_dep_.CoverTab[24430]++

										psk := hs.suite.expandLabel(sessionState.resumptionSecret, "resumption",
			nil, hs.suite.hash.Size())
		hs.earlySecret = hs.suite.extract(psk, nil)
		binderKey := hs.suite.deriveSecret(hs.earlySecret, resumptionBinderLabel, nil)

		transcript := cloneHash(hs.transcript, hs.suite.hash)
		if transcript == nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:305
			_go_fuzz_dep_.CoverTab[24452]++
											c.sendAlert(alertInternalError)
											return errors.New("tls: internal error: failed to clone hash")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:307
			// _ = "end of CoverTab[24452]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:308
			_go_fuzz_dep_.CoverTab[24453]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:308
			// _ = "end of CoverTab[24453]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:308
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:308
		// _ = "end of CoverTab[24430]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:308
		_go_fuzz_dep_.CoverTab[24431]++
										clientHelloBytes, err := hs.clientHello.marshalWithoutBinders()
										if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:310
			_go_fuzz_dep_.CoverTab[24454]++
											c.sendAlert(alertInternalError)
											return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:312
			// _ = "end of CoverTab[24454]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:313
			_go_fuzz_dep_.CoverTab[24455]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:313
			// _ = "end of CoverTab[24455]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:313
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:313
		// _ = "end of CoverTab[24431]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:313
		_go_fuzz_dep_.CoverTab[24432]++
										transcript.Write(clientHelloBytes)
										pskBinder := hs.suite.finishedHash(binderKey, transcript)
										if !hmac.Equal(hs.clientHello.pskBinders[i], pskBinder) {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:316
			_go_fuzz_dep_.CoverTab[24456]++
											c.sendAlert(alertDecryptError)
											return errors.New("tls: invalid PSK binder")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:318
			// _ = "end of CoverTab[24456]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:319
			_go_fuzz_dep_.CoverTab[24457]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:319
			// _ = "end of CoverTab[24457]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:319
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:319
		// _ = "end of CoverTab[24432]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:319
		_go_fuzz_dep_.CoverTab[24433]++

										c.didResume = true
										if err := c.processCertsFromClient(sessionState.certificate); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:322
			_go_fuzz_dep_.CoverTab[24458]++
											return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:323
			// _ = "end of CoverTab[24458]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:324
			_go_fuzz_dep_.CoverTab[24459]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:324
			// _ = "end of CoverTab[24459]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:324
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:324
		// _ = "end of CoverTab[24433]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:324
		_go_fuzz_dep_.CoverTab[24434]++

										hs.hello.selectedIdentityPresent = true
										hs.hello.selectedIdentity = uint16(i)
										hs.usingPSK = true
										return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:329
		// _ = "end of CoverTab[24434]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:330
	// _ = "end of CoverTab[24410]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:330
	_go_fuzz_dep_.CoverTab[24411]++

									return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:332
	// _ = "end of CoverTab[24411]"
}

// cloneHash uses the encoding.BinaryMarshaler and encoding.BinaryUnmarshaler
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:335
// interfaces implemented by standard library hashes to clone the state of in
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:335
// to a new instance of h. It returns nil if the operation fails.
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:338
func cloneHash(in hash.Hash, h crypto.Hash) hash.Hash {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:338
	_go_fuzz_dep_.CoverTab[24460]++
	// Recreate the interface to avoid importing encoding.
	type binaryMarshaler interface {
		MarshalBinary() (data []byte, err error)
		UnmarshalBinary(data []byte) error
	}
	marshaler, ok := in.(binaryMarshaler)
	if !ok {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:345
		_go_fuzz_dep_.CoverTab[24465]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:346
		// _ = "end of CoverTab[24465]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:347
		_go_fuzz_dep_.CoverTab[24466]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:347
		// _ = "end of CoverTab[24466]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:347
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:347
	// _ = "end of CoverTab[24460]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:347
	_go_fuzz_dep_.CoverTab[24461]++
									state, err := marshaler.MarshalBinary()
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:349
		_go_fuzz_dep_.CoverTab[24467]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:350
		// _ = "end of CoverTab[24467]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:351
		_go_fuzz_dep_.CoverTab[24468]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:351
		// _ = "end of CoverTab[24468]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:351
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:351
	// _ = "end of CoverTab[24461]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:351
	_go_fuzz_dep_.CoverTab[24462]++
									out := h.New()
									unmarshaler, ok := out.(binaryMarshaler)
									if !ok {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:354
		_go_fuzz_dep_.CoverTab[24469]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:355
		// _ = "end of CoverTab[24469]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:356
		_go_fuzz_dep_.CoverTab[24470]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:356
		// _ = "end of CoverTab[24470]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:356
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:356
	// _ = "end of CoverTab[24462]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:356
	_go_fuzz_dep_.CoverTab[24463]++
									if err := unmarshaler.UnmarshalBinary(state); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:357
		_go_fuzz_dep_.CoverTab[24471]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:358
		// _ = "end of CoverTab[24471]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:359
		_go_fuzz_dep_.CoverTab[24472]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:359
		// _ = "end of CoverTab[24472]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:359
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:359
	// _ = "end of CoverTab[24463]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:359
	_go_fuzz_dep_.CoverTab[24464]++
									return out
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:360
	// _ = "end of CoverTab[24464]"
}

func (hs *serverHandshakeStateTLS13) pickCertificate() error {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:363
	_go_fuzz_dep_.CoverTab[24473]++
									c := hs.c

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:367
	if hs.usingPSK {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:367
		_go_fuzz_dep_.CoverTab[24478]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:368
		// _ = "end of CoverTab[24478]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:369
		_go_fuzz_dep_.CoverTab[24479]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:369
		// _ = "end of CoverTab[24479]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:369
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:369
	// _ = "end of CoverTab[24473]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:369
	_go_fuzz_dep_.CoverTab[24474]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:372
	if len(hs.clientHello.supportedSignatureAlgorithms) == 0 {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:372
		_go_fuzz_dep_.CoverTab[24480]++
										return c.sendAlert(alertMissingExtension)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:373
		// _ = "end of CoverTab[24480]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:374
		_go_fuzz_dep_.CoverTab[24481]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:374
		// _ = "end of CoverTab[24481]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:374
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:374
	// _ = "end of CoverTab[24474]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:374
	_go_fuzz_dep_.CoverTab[24475]++

									certificate, err := c.config.getCertificate(clientHelloInfo(hs.ctx, c, hs.clientHello))
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:377
		_go_fuzz_dep_.CoverTab[24482]++
										if err == errNoCertificates {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:378
			_go_fuzz_dep_.CoverTab[24484]++
											c.sendAlert(alertUnrecognizedName)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:379
			// _ = "end of CoverTab[24484]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:380
			_go_fuzz_dep_.CoverTab[24485]++
											c.sendAlert(alertInternalError)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:381
			// _ = "end of CoverTab[24485]"
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:382
		// _ = "end of CoverTab[24482]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:382
		_go_fuzz_dep_.CoverTab[24483]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:383
		// _ = "end of CoverTab[24483]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:384
		_go_fuzz_dep_.CoverTab[24486]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:384
		// _ = "end of CoverTab[24486]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:384
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:384
	// _ = "end of CoverTab[24475]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:384
	_go_fuzz_dep_.CoverTab[24476]++
									hs.sigAlg, err = selectSignatureScheme(c.vers, certificate, hs.clientHello.supportedSignatureAlgorithms)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:386
		_go_fuzz_dep_.CoverTab[24487]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:389
		c.sendAlert(alertHandshakeFailure)
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:390
		// _ = "end of CoverTab[24487]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:391
		_go_fuzz_dep_.CoverTab[24488]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:391
		// _ = "end of CoverTab[24488]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:391
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:391
	// _ = "end of CoverTab[24476]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:391
	_go_fuzz_dep_.CoverTab[24477]++
									hs.cert = certificate

									return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:394
	// _ = "end of CoverTab[24477]"
}

// sendDummyChangeCipherSpec sends a ChangeCipherSpec record for compatibility
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:397
// with middleboxes that didn't implement TLS correctly. See RFC 8446, Appendix D.4.
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:399
func (hs *serverHandshakeStateTLS13) sendDummyChangeCipherSpec() error {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:399
	_go_fuzz_dep_.CoverTab[24489]++
									if hs.sentDummyCCS {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:400
		_go_fuzz_dep_.CoverTab[24491]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:401
		// _ = "end of CoverTab[24491]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:402
		_go_fuzz_dep_.CoverTab[24492]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:402
		// _ = "end of CoverTab[24492]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:402
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:402
	// _ = "end of CoverTab[24489]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:402
	_go_fuzz_dep_.CoverTab[24490]++
									hs.sentDummyCCS = true

									return hs.c.writeChangeCipherRecord()
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:405
	// _ = "end of CoverTab[24490]"
}

func (hs *serverHandshakeStateTLS13) doHelloRetryRequest(selectedGroup CurveID) error {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:408
	_go_fuzz_dep_.CoverTab[24493]++
									c := hs.c

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:413
	if err := transcriptMsg(hs.clientHello, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:413
		_go_fuzz_dep_.CoverTab[24502]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:414
		// _ = "end of CoverTab[24502]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:415
		_go_fuzz_dep_.CoverTab[24503]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:415
		// _ = "end of CoverTab[24503]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:415
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:415
	// _ = "end of CoverTab[24493]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:415
	_go_fuzz_dep_.CoverTab[24494]++
									chHash := hs.transcript.Sum(nil)
									hs.transcript.Reset()
									hs.transcript.Write([]byte{typeMessageHash, 0, 0, uint8(len(chHash))})
									hs.transcript.Write(chHash)

									helloRetryRequest := &serverHelloMsg{
		vers:			hs.hello.vers,
		random:			helloRetryRequestRandom,
		sessionId:		hs.hello.sessionId,
		cipherSuite:		hs.hello.cipherSuite,
		compressionMethod:	hs.hello.compressionMethod,
		supportedVersion:	hs.hello.supportedVersion,
		selectedGroup:		selectedGroup,
	}

	if _, err := hs.c.writeHandshakeRecord(helloRetryRequest, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:431
		_go_fuzz_dep_.CoverTab[24504]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:432
		// _ = "end of CoverTab[24504]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:433
		_go_fuzz_dep_.CoverTab[24505]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:433
		// _ = "end of CoverTab[24505]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:433
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:433
	// _ = "end of CoverTab[24494]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:433
	_go_fuzz_dep_.CoverTab[24495]++

									if err := hs.sendDummyChangeCipherSpec(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:435
		_go_fuzz_dep_.CoverTab[24506]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:436
		// _ = "end of CoverTab[24506]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:437
		_go_fuzz_dep_.CoverTab[24507]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:437
		// _ = "end of CoverTab[24507]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:437
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:437
	// _ = "end of CoverTab[24495]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:437
	_go_fuzz_dep_.CoverTab[24496]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:440
	msg, err := c.readHandshake(nil)
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:441
		_go_fuzz_dep_.CoverTab[24508]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:442
		// _ = "end of CoverTab[24508]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:443
		_go_fuzz_dep_.CoverTab[24509]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:443
		// _ = "end of CoverTab[24509]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:443
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:443
	// _ = "end of CoverTab[24496]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:443
	_go_fuzz_dep_.CoverTab[24497]++

									clientHello, ok := msg.(*clientHelloMsg)
									if !ok {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:446
		_go_fuzz_dep_.CoverTab[24510]++
										c.sendAlert(alertUnexpectedMessage)
										return unexpectedMessageError(clientHello, msg)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:448
		// _ = "end of CoverTab[24510]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:449
		_go_fuzz_dep_.CoverTab[24511]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:449
		// _ = "end of CoverTab[24511]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:449
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:449
	// _ = "end of CoverTab[24497]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:449
	_go_fuzz_dep_.CoverTab[24498]++

									if len(clientHello.keyShares) != 1 || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:451
		_go_fuzz_dep_.CoverTab[24512]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:451
		return clientHello.keyShares[0].group != selectedGroup
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:451
		// _ = "end of CoverTab[24512]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:451
	}() {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:451
		_go_fuzz_dep_.CoverTab[24513]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: client sent invalid key share in second ClientHello")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:453
		// _ = "end of CoverTab[24513]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:454
		_go_fuzz_dep_.CoverTab[24514]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:454
		// _ = "end of CoverTab[24514]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:454
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:454
	// _ = "end of CoverTab[24498]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:454
	_go_fuzz_dep_.CoverTab[24499]++

									if clientHello.earlyData {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:456
		_go_fuzz_dep_.CoverTab[24515]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: client indicated early data in second ClientHello")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:458
		// _ = "end of CoverTab[24515]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:459
		_go_fuzz_dep_.CoverTab[24516]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:459
		// _ = "end of CoverTab[24516]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:459
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:459
	// _ = "end of CoverTab[24499]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:459
	_go_fuzz_dep_.CoverTab[24500]++

									if illegalClientHelloChange(clientHello, hs.clientHello) {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:461
		_go_fuzz_dep_.CoverTab[24517]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: client illegally modified second ClientHello")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:463
		// _ = "end of CoverTab[24517]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:464
		_go_fuzz_dep_.CoverTab[24518]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:464
		// _ = "end of CoverTab[24518]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:464
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:464
	// _ = "end of CoverTab[24500]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:464
	_go_fuzz_dep_.CoverTab[24501]++

									hs.clientHello = clientHello
									return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:467
	// _ = "end of CoverTab[24501]"
}

// illegalClientHelloChange reports whether the two ClientHello messages are
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:470
// different, with the exception of the changes allowed before and after a
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:470
// HelloRetryRequest. See RFC 8446, Section 4.1.2.
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:473
func illegalClientHelloChange(ch, ch1 *clientHelloMsg) bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:473
	_go_fuzz_dep_.CoverTab[24519]++
									if len(ch.supportedVersions) != len(ch1.supportedVersions) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:474
		_go_fuzz_dep_.CoverTab[24527]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:474
		return len(ch.cipherSuites) != len(ch1.cipherSuites)
										// _ = "end of CoverTab[24527]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:475
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:475
		_go_fuzz_dep_.CoverTab[24528]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:475
		return len(ch.supportedCurves) != len(ch1.supportedCurves)
										// _ = "end of CoverTab[24528]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:476
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:476
		_go_fuzz_dep_.CoverTab[24529]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:476
		return len(ch.supportedSignatureAlgorithms) != len(ch1.supportedSignatureAlgorithms)
										// _ = "end of CoverTab[24529]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:477
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:477
		_go_fuzz_dep_.CoverTab[24530]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:477
		return len(ch.supportedSignatureAlgorithmsCert) != len(ch1.supportedSignatureAlgorithmsCert)
										// _ = "end of CoverTab[24530]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:478
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:478
		_go_fuzz_dep_.CoverTab[24531]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:478
		return len(ch.alpnProtocols) != len(ch1.alpnProtocols)
										// _ = "end of CoverTab[24531]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:479
	}() {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:479
		_go_fuzz_dep_.CoverTab[24532]++
										return true
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:480
		// _ = "end of CoverTab[24532]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:481
		_go_fuzz_dep_.CoverTab[24533]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:481
		// _ = "end of CoverTab[24533]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:481
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:481
	// _ = "end of CoverTab[24519]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:481
	_go_fuzz_dep_.CoverTab[24520]++
									for i := range ch.supportedVersions {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:482
		_go_fuzz_dep_.CoverTab[24534]++
										if ch.supportedVersions[i] != ch1.supportedVersions[i] {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:483
			_go_fuzz_dep_.CoverTab[24535]++
											return true
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:484
			// _ = "end of CoverTab[24535]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:485
			_go_fuzz_dep_.CoverTab[24536]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:485
			// _ = "end of CoverTab[24536]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:485
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:485
		// _ = "end of CoverTab[24534]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:486
	// _ = "end of CoverTab[24520]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:486
	_go_fuzz_dep_.CoverTab[24521]++
									for i := range ch.cipherSuites {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:487
		_go_fuzz_dep_.CoverTab[24537]++
										if ch.cipherSuites[i] != ch1.cipherSuites[i] {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:488
			_go_fuzz_dep_.CoverTab[24538]++
											return true
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:489
			// _ = "end of CoverTab[24538]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:490
			_go_fuzz_dep_.CoverTab[24539]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:490
			// _ = "end of CoverTab[24539]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:490
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:490
		// _ = "end of CoverTab[24537]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:491
	// _ = "end of CoverTab[24521]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:491
	_go_fuzz_dep_.CoverTab[24522]++
									for i := range ch.supportedCurves {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:492
		_go_fuzz_dep_.CoverTab[24540]++
										if ch.supportedCurves[i] != ch1.supportedCurves[i] {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:493
			_go_fuzz_dep_.CoverTab[24541]++
											return true
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:494
			// _ = "end of CoverTab[24541]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:495
			_go_fuzz_dep_.CoverTab[24542]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:495
			// _ = "end of CoverTab[24542]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:495
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:495
		// _ = "end of CoverTab[24540]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:496
	// _ = "end of CoverTab[24522]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:496
	_go_fuzz_dep_.CoverTab[24523]++
									for i := range ch.supportedSignatureAlgorithms {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:497
		_go_fuzz_dep_.CoverTab[24543]++
										if ch.supportedSignatureAlgorithms[i] != ch1.supportedSignatureAlgorithms[i] {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:498
			_go_fuzz_dep_.CoverTab[24544]++
											return true
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:499
			// _ = "end of CoverTab[24544]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:500
			_go_fuzz_dep_.CoverTab[24545]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:500
			// _ = "end of CoverTab[24545]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:500
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:500
		// _ = "end of CoverTab[24543]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:501
	// _ = "end of CoverTab[24523]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:501
	_go_fuzz_dep_.CoverTab[24524]++
									for i := range ch.supportedSignatureAlgorithmsCert {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:502
		_go_fuzz_dep_.CoverTab[24546]++
										if ch.supportedSignatureAlgorithmsCert[i] != ch1.supportedSignatureAlgorithmsCert[i] {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:503
			_go_fuzz_dep_.CoverTab[24547]++
											return true
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:504
			// _ = "end of CoverTab[24547]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:505
			_go_fuzz_dep_.CoverTab[24548]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:505
			// _ = "end of CoverTab[24548]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:505
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:505
		// _ = "end of CoverTab[24546]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:506
	// _ = "end of CoverTab[24524]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:506
	_go_fuzz_dep_.CoverTab[24525]++
									for i := range ch.alpnProtocols {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:507
		_go_fuzz_dep_.CoverTab[24549]++
										if ch.alpnProtocols[i] != ch1.alpnProtocols[i] {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:508
			_go_fuzz_dep_.CoverTab[24550]++
											return true
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:509
			// _ = "end of CoverTab[24550]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:510
			_go_fuzz_dep_.CoverTab[24551]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:510
			// _ = "end of CoverTab[24551]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:510
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:510
		// _ = "end of CoverTab[24549]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:511
	// _ = "end of CoverTab[24525]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:511
	_go_fuzz_dep_.CoverTab[24526]++
									return ch.vers != ch1.vers || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:512
		_go_fuzz_dep_.CoverTab[24552]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:512
		return !bytes.Equal(ch.random, ch1.random)
										// _ = "end of CoverTab[24552]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:513
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:513
		_go_fuzz_dep_.CoverTab[24553]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:513
		return !bytes.Equal(ch.sessionId, ch1.sessionId)
										// _ = "end of CoverTab[24553]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:514
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:514
		_go_fuzz_dep_.CoverTab[24554]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:514
		return !bytes.Equal(ch.compressionMethods, ch1.compressionMethods)
										// _ = "end of CoverTab[24554]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:515
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:515
		_go_fuzz_dep_.CoverTab[24555]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:515
		return ch.serverName != ch1.serverName
										// _ = "end of CoverTab[24555]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:516
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:516
		_go_fuzz_dep_.CoverTab[24556]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:516
		return ch.ocspStapling != ch1.ocspStapling
										// _ = "end of CoverTab[24556]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:517
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:517
		_go_fuzz_dep_.CoverTab[24557]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:517
		return !bytes.Equal(ch.supportedPoints, ch1.supportedPoints)
										// _ = "end of CoverTab[24557]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:518
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:518
		_go_fuzz_dep_.CoverTab[24558]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:518
		return ch.ticketSupported != ch1.ticketSupported
										// _ = "end of CoverTab[24558]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:519
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:519
		_go_fuzz_dep_.CoverTab[24559]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:519
		return !bytes.Equal(ch.sessionTicket, ch1.sessionTicket)
										// _ = "end of CoverTab[24559]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:520
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:520
		_go_fuzz_dep_.CoverTab[24560]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:520
		return ch.secureRenegotiationSupported != ch1.secureRenegotiationSupported
										// _ = "end of CoverTab[24560]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:521
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:521
		_go_fuzz_dep_.CoverTab[24561]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:521
		return !bytes.Equal(ch.secureRenegotiation, ch1.secureRenegotiation)
										// _ = "end of CoverTab[24561]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:522
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:522
		_go_fuzz_dep_.CoverTab[24562]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:522
		return ch.scts != ch1.scts
										// _ = "end of CoverTab[24562]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:523
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:523
		_go_fuzz_dep_.CoverTab[24563]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:523
		return !bytes.Equal(ch.cookie, ch1.cookie)
										// _ = "end of CoverTab[24563]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:524
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:524
		_go_fuzz_dep_.CoverTab[24564]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:524
		return !bytes.Equal(ch.pskModes, ch1.pskModes)
										// _ = "end of CoverTab[24564]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:525
	}()
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:525
	// _ = "end of CoverTab[24526]"
}

func (hs *serverHandshakeStateTLS13) sendServerParameters() error {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:528
	_go_fuzz_dep_.CoverTab[24565]++
									c := hs.c

									if err := transcriptMsg(hs.clientHello, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:531
		_go_fuzz_dep_.CoverTab[24574]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:532
		// _ = "end of CoverTab[24574]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:533
		_go_fuzz_dep_.CoverTab[24575]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:533
		// _ = "end of CoverTab[24575]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:533
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:533
	// _ = "end of CoverTab[24565]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:533
	_go_fuzz_dep_.CoverTab[24566]++
									if _, err := hs.c.writeHandshakeRecord(hs.hello, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:534
		_go_fuzz_dep_.CoverTab[24576]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:535
		// _ = "end of CoverTab[24576]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:536
		_go_fuzz_dep_.CoverTab[24577]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:536
		// _ = "end of CoverTab[24577]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:536
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:536
	// _ = "end of CoverTab[24566]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:536
	_go_fuzz_dep_.CoverTab[24567]++

									if err := hs.sendDummyChangeCipherSpec(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:538
		_go_fuzz_dep_.CoverTab[24578]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:539
		// _ = "end of CoverTab[24578]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:540
		_go_fuzz_dep_.CoverTab[24579]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:540
		// _ = "end of CoverTab[24579]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:540
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:540
	// _ = "end of CoverTab[24567]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:540
	_go_fuzz_dep_.CoverTab[24568]++

									earlySecret := hs.earlySecret
									if earlySecret == nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:543
		_go_fuzz_dep_.CoverTab[24580]++
										earlySecret = hs.suite.extract(nil, nil)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:544
		// _ = "end of CoverTab[24580]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:545
		_go_fuzz_dep_.CoverTab[24581]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:545
		// _ = "end of CoverTab[24581]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:545
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:545
	// _ = "end of CoverTab[24568]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:545
	_go_fuzz_dep_.CoverTab[24569]++
									hs.handshakeSecret = hs.suite.extract(hs.sharedKey,
		hs.suite.deriveSecret(earlySecret, "derived", nil))

	clientSecret := hs.suite.deriveSecret(hs.handshakeSecret,
		clientHandshakeTrafficLabel, hs.transcript)
	c.in.setTrafficSecret(hs.suite, clientSecret)
	serverSecret := hs.suite.deriveSecret(hs.handshakeSecret,
		serverHandshakeTrafficLabel, hs.transcript)
	c.out.setTrafficSecret(hs.suite, serverSecret)

	err := c.config.writeKeyLog(keyLogLabelClientHandshake, hs.clientHello.random, clientSecret)
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:557
		_go_fuzz_dep_.CoverTab[24582]++
										c.sendAlert(alertInternalError)
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:559
		// _ = "end of CoverTab[24582]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:560
		_go_fuzz_dep_.CoverTab[24583]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:560
		// _ = "end of CoverTab[24583]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:560
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:560
	// _ = "end of CoverTab[24569]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:560
	_go_fuzz_dep_.CoverTab[24570]++
									err = c.config.writeKeyLog(keyLogLabelServerHandshake, hs.clientHello.random, serverSecret)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:562
		_go_fuzz_dep_.CoverTab[24584]++
										c.sendAlert(alertInternalError)
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:564
		// _ = "end of CoverTab[24584]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:565
		_go_fuzz_dep_.CoverTab[24585]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:565
		// _ = "end of CoverTab[24585]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:565
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:565
	// _ = "end of CoverTab[24570]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:565
	_go_fuzz_dep_.CoverTab[24571]++

									encryptedExtensions := new(encryptedExtensionsMsg)

									selectedProto, err := negotiateALPN(c.config.NextProtos, hs.clientHello.alpnProtocols)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:570
		_go_fuzz_dep_.CoverTab[24586]++
										c.sendAlert(alertNoApplicationProtocol)
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:572
		// _ = "end of CoverTab[24586]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:573
		_go_fuzz_dep_.CoverTab[24587]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:573
		// _ = "end of CoverTab[24587]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:573
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:573
	// _ = "end of CoverTab[24571]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:573
	_go_fuzz_dep_.CoverTab[24572]++
									encryptedExtensions.alpnProtocol = selectedProto
									c.clientProtocol = selectedProto

									if _, err := hs.c.writeHandshakeRecord(encryptedExtensions, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:577
		_go_fuzz_dep_.CoverTab[24588]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:578
		// _ = "end of CoverTab[24588]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:579
		_go_fuzz_dep_.CoverTab[24589]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:579
		// _ = "end of CoverTab[24589]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:579
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:579
	// _ = "end of CoverTab[24572]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:579
	_go_fuzz_dep_.CoverTab[24573]++

									return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:581
	// _ = "end of CoverTab[24573]"
}

func (hs *serverHandshakeStateTLS13) requestClientCert() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:584
	_go_fuzz_dep_.CoverTab[24590]++
									return hs.c.config.ClientAuth >= RequestClientCert && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:585
		_go_fuzz_dep_.CoverTab[24591]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:585
		return !hs.usingPSK
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:585
		// _ = "end of CoverTab[24591]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:585
	}()
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:585
	// _ = "end of CoverTab[24590]"
}

func (hs *serverHandshakeStateTLS13) sendServerCertificate() error {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:588
	_go_fuzz_dep_.CoverTab[24592]++
									c := hs.c

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:592
	if hs.usingPSK {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:592
		_go_fuzz_dep_.CoverTab[24600]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:593
		// _ = "end of CoverTab[24600]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:594
		_go_fuzz_dep_.CoverTab[24601]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:594
		// _ = "end of CoverTab[24601]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:594
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:594
	// _ = "end of CoverTab[24592]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:594
	_go_fuzz_dep_.CoverTab[24593]++

									if hs.requestClientCert() {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:596
		_go_fuzz_dep_.CoverTab[24602]++

										certReq := new(certificateRequestMsgTLS13)
										certReq.ocspStapling = true
										certReq.scts = true
										certReq.supportedSignatureAlgorithms = supportedSignatureAlgorithms()
										if c.config.ClientCAs != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:602
			_go_fuzz_dep_.CoverTab[24604]++
											certReq.certificateAuthorities = c.config.ClientCAs.Subjects()
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:603
			// _ = "end of CoverTab[24604]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:604
			_go_fuzz_dep_.CoverTab[24605]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:604
			// _ = "end of CoverTab[24605]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:604
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:604
		// _ = "end of CoverTab[24602]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:604
		_go_fuzz_dep_.CoverTab[24603]++

										if _, err := hs.c.writeHandshakeRecord(certReq, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:606
			_go_fuzz_dep_.CoverTab[24606]++
											return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:607
			// _ = "end of CoverTab[24606]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:608
			_go_fuzz_dep_.CoverTab[24607]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:608
			// _ = "end of CoverTab[24607]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:608
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:608
		// _ = "end of CoverTab[24603]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:609
		_go_fuzz_dep_.CoverTab[24608]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:609
		// _ = "end of CoverTab[24608]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:609
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:609
	// _ = "end of CoverTab[24593]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:609
	_go_fuzz_dep_.CoverTab[24594]++

									certMsg := new(certificateMsgTLS13)

									certMsg.certificate = *hs.cert
									certMsg.scts = hs.clientHello.scts && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:614
		_go_fuzz_dep_.CoverTab[24609]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:614
		return len(hs.cert.SignedCertificateTimestamps) > 0
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:614
		// _ = "end of CoverTab[24609]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:614
	}()
									certMsg.ocspStapling = hs.clientHello.ocspStapling && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:615
		_go_fuzz_dep_.CoverTab[24610]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:615
		return len(hs.cert.OCSPStaple) > 0
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:615
		// _ = "end of CoverTab[24610]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:615
	}()

									if _, err := hs.c.writeHandshakeRecord(certMsg, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:617
		_go_fuzz_dep_.CoverTab[24611]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:618
		// _ = "end of CoverTab[24611]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:619
		_go_fuzz_dep_.CoverTab[24612]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:619
		// _ = "end of CoverTab[24612]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:619
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:619
	// _ = "end of CoverTab[24594]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:619
	_go_fuzz_dep_.CoverTab[24595]++

									certVerifyMsg := new(certificateVerifyMsg)
									certVerifyMsg.hasSignatureAlgorithm = true
									certVerifyMsg.signatureAlgorithm = hs.sigAlg

									sigType, sigHash, err := typeAndHashFromSignatureScheme(hs.sigAlg)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:626
		_go_fuzz_dep_.CoverTab[24613]++
										return c.sendAlert(alertInternalError)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:627
		// _ = "end of CoverTab[24613]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:628
		_go_fuzz_dep_.CoverTab[24614]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:628
		// _ = "end of CoverTab[24614]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:628
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:628
	// _ = "end of CoverTab[24595]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:628
	_go_fuzz_dep_.CoverTab[24596]++

									signed := signedMessage(sigHash, serverSignatureContext, hs.transcript)
									signOpts := crypto.SignerOpts(sigHash)
									if sigType == signatureRSAPSS {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:632
		_go_fuzz_dep_.CoverTab[24615]++
										signOpts = &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash, Hash: sigHash}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:633
		// _ = "end of CoverTab[24615]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:634
		_go_fuzz_dep_.CoverTab[24616]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:634
		// _ = "end of CoverTab[24616]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:634
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:634
	// _ = "end of CoverTab[24596]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:634
	_go_fuzz_dep_.CoverTab[24597]++
									sig, err := hs.cert.PrivateKey.(crypto.Signer).Sign(c.config.rand(), signed, signOpts)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:636
		_go_fuzz_dep_.CoverTab[24617]++
										public := hs.cert.PrivateKey.(crypto.Signer).Public()
										if rsaKey, ok := public.(*rsa.PublicKey); ok && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:638
			_go_fuzz_dep_.CoverTab[24619]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:638
			return sigType == signatureRSAPSS
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:638
			// _ = "end of CoverTab[24619]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:638
		}() && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:638
			_go_fuzz_dep_.CoverTab[24620]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:638
			return rsaKey.N.BitLen()/8 < sigHash.Size()*2+2
											// _ = "end of CoverTab[24620]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:639
		}() {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:639
			_go_fuzz_dep_.CoverTab[24621]++
											c.sendAlert(alertHandshakeFailure)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:640
			// _ = "end of CoverTab[24621]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:641
			_go_fuzz_dep_.CoverTab[24622]++
											c.sendAlert(alertInternalError)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:642
			// _ = "end of CoverTab[24622]"
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:643
		// _ = "end of CoverTab[24617]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:643
		_go_fuzz_dep_.CoverTab[24618]++
										return errors.New("tls: failed to sign handshake: " + err.Error())
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:644
		// _ = "end of CoverTab[24618]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:645
		_go_fuzz_dep_.CoverTab[24623]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:645
		// _ = "end of CoverTab[24623]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:645
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:645
	// _ = "end of CoverTab[24597]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:645
	_go_fuzz_dep_.CoverTab[24598]++
									certVerifyMsg.signature = sig

									if _, err := hs.c.writeHandshakeRecord(certVerifyMsg, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:648
		_go_fuzz_dep_.CoverTab[24624]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:649
		// _ = "end of CoverTab[24624]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:650
		_go_fuzz_dep_.CoverTab[24625]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:650
		// _ = "end of CoverTab[24625]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:650
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:650
	// _ = "end of CoverTab[24598]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:650
	_go_fuzz_dep_.CoverTab[24599]++

									return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:652
	// _ = "end of CoverTab[24599]"
}

func (hs *serverHandshakeStateTLS13) sendServerFinished() error {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:655
	_go_fuzz_dep_.CoverTab[24626]++
									c := hs.c

									finished := &finishedMsg{
		verifyData: hs.suite.finishedHash(c.out.trafficSecret, hs.transcript),
	}

	if _, err := hs.c.writeHandshakeRecord(finished, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:662
		_go_fuzz_dep_.CoverTab[24631]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:663
		// _ = "end of CoverTab[24631]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:664
		_go_fuzz_dep_.CoverTab[24632]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:664
		// _ = "end of CoverTab[24632]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:664
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:664
	// _ = "end of CoverTab[24626]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:664
	_go_fuzz_dep_.CoverTab[24627]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:668
	hs.masterSecret = hs.suite.extract(nil,
		hs.suite.deriveSecret(hs.handshakeSecret, "derived", nil))

	hs.trafficSecret = hs.suite.deriveSecret(hs.masterSecret,
		clientApplicationTrafficLabel, hs.transcript)
	serverSecret := hs.suite.deriveSecret(hs.masterSecret,
		serverApplicationTrafficLabel, hs.transcript)
	c.out.setTrafficSecret(hs.suite, serverSecret)

	err := c.config.writeKeyLog(keyLogLabelClientTraffic, hs.clientHello.random, hs.trafficSecret)
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:678
		_go_fuzz_dep_.CoverTab[24633]++
										c.sendAlert(alertInternalError)
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:680
		// _ = "end of CoverTab[24633]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:681
		_go_fuzz_dep_.CoverTab[24634]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:681
		// _ = "end of CoverTab[24634]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:681
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:681
	// _ = "end of CoverTab[24627]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:681
	_go_fuzz_dep_.CoverTab[24628]++
									err = c.config.writeKeyLog(keyLogLabelServerTraffic, hs.clientHello.random, serverSecret)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:683
		_go_fuzz_dep_.CoverTab[24635]++
										c.sendAlert(alertInternalError)
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:685
		// _ = "end of CoverTab[24635]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:686
		_go_fuzz_dep_.CoverTab[24636]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:686
		// _ = "end of CoverTab[24636]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:686
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:686
	// _ = "end of CoverTab[24628]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:686
	_go_fuzz_dep_.CoverTab[24629]++

									c.ekm = hs.suite.exportKeyingMaterial(hs.masterSecret, hs.transcript)

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:693
	if !hs.requestClientCert() {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:693
		_go_fuzz_dep_.CoverTab[24637]++
										if err := hs.sendSessionTickets(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:694
			_go_fuzz_dep_.CoverTab[24638]++
											return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:695
			// _ = "end of CoverTab[24638]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:696
			_go_fuzz_dep_.CoverTab[24639]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:696
			// _ = "end of CoverTab[24639]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:696
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:696
		// _ = "end of CoverTab[24637]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:697
		_go_fuzz_dep_.CoverTab[24640]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:697
		// _ = "end of CoverTab[24640]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:697
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:697
	// _ = "end of CoverTab[24629]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:697
	_go_fuzz_dep_.CoverTab[24630]++

									return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:699
	// _ = "end of CoverTab[24630]"
}

func (hs *serverHandshakeStateTLS13) shouldSendSessionTickets() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:702
	_go_fuzz_dep_.CoverTab[24641]++
									if hs.c.config.SessionTicketsDisabled {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:703
		_go_fuzz_dep_.CoverTab[24644]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:704
		// _ = "end of CoverTab[24644]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:705
		_go_fuzz_dep_.CoverTab[24645]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:705
		// _ = "end of CoverTab[24645]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:705
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:705
	// _ = "end of CoverTab[24641]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:705
	_go_fuzz_dep_.CoverTab[24642]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:708
	for _, pskMode := range hs.clientHello.pskModes {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:708
		_go_fuzz_dep_.CoverTab[24646]++
										if pskMode == pskModeDHE {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:709
			_go_fuzz_dep_.CoverTab[24647]++
											return true
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:710
			// _ = "end of CoverTab[24647]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:711
			_go_fuzz_dep_.CoverTab[24648]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:711
			// _ = "end of CoverTab[24648]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:711
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:711
		// _ = "end of CoverTab[24646]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:712
	// _ = "end of CoverTab[24642]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:712
	_go_fuzz_dep_.CoverTab[24643]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:713
	// _ = "end of CoverTab[24643]"
}

func (hs *serverHandshakeStateTLS13) sendSessionTickets() error {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:716
	_go_fuzz_dep_.CoverTab[24649]++
									c := hs.c

									hs.clientFinished = hs.suite.finishedHash(c.in.trafficSecret, hs.transcript)
									finishedMsg := &finishedMsg{
		verifyData: hs.clientFinished,
	}
	if err := transcriptMsg(finishedMsg, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:723
		_go_fuzz_dep_.CoverTab[24657]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:724
		// _ = "end of CoverTab[24657]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:725
		_go_fuzz_dep_.CoverTab[24658]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:725
		// _ = "end of CoverTab[24658]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:725
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:725
	// _ = "end of CoverTab[24649]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:725
	_go_fuzz_dep_.CoverTab[24650]++

									if !hs.shouldSendSessionTickets() {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:727
		_go_fuzz_dep_.CoverTab[24659]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:728
		// _ = "end of CoverTab[24659]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:729
		_go_fuzz_dep_.CoverTab[24660]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:729
		// _ = "end of CoverTab[24660]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:729
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:729
	// _ = "end of CoverTab[24650]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:729
	_go_fuzz_dep_.CoverTab[24651]++

									resumptionSecret := hs.suite.deriveSecret(hs.masterSecret,
		resumptionLabel, hs.transcript)

	m := new(newSessionTicketMsgTLS13)

	var certsFromClient [][]byte
	for _, cert := range c.peerCertificates {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:737
		_go_fuzz_dep_.CoverTab[24661]++
										certsFromClient = append(certsFromClient, cert.Raw)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:738
		// _ = "end of CoverTab[24661]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:739
	// _ = "end of CoverTab[24651]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:739
	_go_fuzz_dep_.CoverTab[24652]++
									state := sessionStateTLS13{
		cipherSuite:		hs.suite.id,
		createdAt:		uint64(c.config.time().Unix()),
		resumptionSecret:	resumptionSecret,
		certificate: Certificate{
			Certificate:			certsFromClient,
			OCSPStaple:			c.ocspResponse,
			SignedCertificateTimestamps:	c.scts,
		},
	}
	stateBytes, err := state.marshal()
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:751
		_go_fuzz_dep_.CoverTab[24662]++
										c.sendAlert(alertInternalError)
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:753
		// _ = "end of CoverTab[24662]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:754
		_go_fuzz_dep_.CoverTab[24663]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:754
		// _ = "end of CoverTab[24663]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:754
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:754
	// _ = "end of CoverTab[24652]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:754
	_go_fuzz_dep_.CoverTab[24653]++
									m.label, err = c.encryptTicket(stateBytes)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:756
		_go_fuzz_dep_.CoverTab[24664]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:757
		// _ = "end of CoverTab[24664]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:758
		_go_fuzz_dep_.CoverTab[24665]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:758
		// _ = "end of CoverTab[24665]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:758
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:758
	// _ = "end of CoverTab[24653]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:758
	_go_fuzz_dep_.CoverTab[24654]++
									m.lifetime = uint32(maxSessionTicketLifetime / time.Second)

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:764
	ageAdd := make([]byte, 4)
	_, err = hs.c.config.rand().Read(ageAdd)
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:766
		_go_fuzz_dep_.CoverTab[24666]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:767
		// _ = "end of CoverTab[24666]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:768
		_go_fuzz_dep_.CoverTab[24667]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:768
		// _ = "end of CoverTab[24667]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:768
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:768
	// _ = "end of CoverTab[24654]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:768
	_go_fuzz_dep_.CoverTab[24655]++
									m.ageAdd = binary.LittleEndian.Uint32(ageAdd)

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:774
	if _, err := c.writeHandshakeRecord(m, nil); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:774
		_go_fuzz_dep_.CoverTab[24668]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:775
		// _ = "end of CoverTab[24668]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:776
		_go_fuzz_dep_.CoverTab[24669]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:776
		// _ = "end of CoverTab[24669]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:776
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:776
	// _ = "end of CoverTab[24655]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:776
	_go_fuzz_dep_.CoverTab[24656]++

									return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:778
	// _ = "end of CoverTab[24656]"
}

func (hs *serverHandshakeStateTLS13) readClientCertificate() error {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:781
	_go_fuzz_dep_.CoverTab[24670]++
									c := hs.c

									if !hs.requestClientCert() {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:784
		_go_fuzz_dep_.CoverTab[24678]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:787
		if c.config.VerifyConnection != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:787
			_go_fuzz_dep_.CoverTab[24680]++
											if err := c.config.VerifyConnection(c.connectionStateLocked()); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:788
				_go_fuzz_dep_.CoverTab[24681]++
												c.sendAlert(alertBadCertificate)
												return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:790
				// _ = "end of CoverTab[24681]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:791
				_go_fuzz_dep_.CoverTab[24682]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:791
				// _ = "end of CoverTab[24682]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:791
			}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:791
			// _ = "end of CoverTab[24680]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:792
			_go_fuzz_dep_.CoverTab[24683]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:792
			// _ = "end of CoverTab[24683]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:792
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:792
		// _ = "end of CoverTab[24678]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:792
		_go_fuzz_dep_.CoverTab[24679]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:793
		// _ = "end of CoverTab[24679]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:794
		_go_fuzz_dep_.CoverTab[24684]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:794
		// _ = "end of CoverTab[24684]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:794
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:794
	// _ = "end of CoverTab[24670]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:794
	_go_fuzz_dep_.CoverTab[24671]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:799
	msg, err := c.readHandshake(hs.transcript)
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:800
		_go_fuzz_dep_.CoverTab[24685]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:801
		// _ = "end of CoverTab[24685]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:802
		_go_fuzz_dep_.CoverTab[24686]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:802
		// _ = "end of CoverTab[24686]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:802
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:802
	// _ = "end of CoverTab[24671]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:802
	_go_fuzz_dep_.CoverTab[24672]++

									certMsg, ok := msg.(*certificateMsgTLS13)
									if !ok {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:805
		_go_fuzz_dep_.CoverTab[24687]++
										c.sendAlert(alertUnexpectedMessage)
										return unexpectedMessageError(certMsg, msg)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:807
		// _ = "end of CoverTab[24687]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:808
		_go_fuzz_dep_.CoverTab[24688]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:808
		// _ = "end of CoverTab[24688]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:808
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:808
	// _ = "end of CoverTab[24672]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:808
	_go_fuzz_dep_.CoverTab[24673]++

									if err := c.processCertsFromClient(certMsg.certificate); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:810
		_go_fuzz_dep_.CoverTab[24689]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:811
		// _ = "end of CoverTab[24689]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:812
		_go_fuzz_dep_.CoverTab[24690]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:812
		// _ = "end of CoverTab[24690]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:812
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:812
	// _ = "end of CoverTab[24673]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:812
	_go_fuzz_dep_.CoverTab[24674]++

									if c.config.VerifyConnection != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:814
		_go_fuzz_dep_.CoverTab[24691]++
										if err := c.config.VerifyConnection(c.connectionStateLocked()); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:815
			_go_fuzz_dep_.CoverTab[24692]++
											c.sendAlert(alertBadCertificate)
											return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:817
			// _ = "end of CoverTab[24692]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:818
			_go_fuzz_dep_.CoverTab[24693]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:818
			// _ = "end of CoverTab[24693]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:818
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:818
		// _ = "end of CoverTab[24691]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:819
		_go_fuzz_dep_.CoverTab[24694]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:819
		// _ = "end of CoverTab[24694]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:819
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:819
	// _ = "end of CoverTab[24674]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:819
	_go_fuzz_dep_.CoverTab[24675]++

									if len(certMsg.certificate.Certificate) != 0 {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:821
		_go_fuzz_dep_.CoverTab[24695]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:825
		msg, err = c.readHandshake(nil)
		if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:826
			_go_fuzz_dep_.CoverTab[24702]++
											return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:827
			// _ = "end of CoverTab[24702]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:828
			_go_fuzz_dep_.CoverTab[24703]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:828
			// _ = "end of CoverTab[24703]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:828
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:828
		// _ = "end of CoverTab[24695]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:828
		_go_fuzz_dep_.CoverTab[24696]++

										certVerify, ok := msg.(*certificateVerifyMsg)
										if !ok {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:831
			_go_fuzz_dep_.CoverTab[24704]++
											c.sendAlert(alertUnexpectedMessage)
											return unexpectedMessageError(certVerify, msg)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:833
			// _ = "end of CoverTab[24704]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:834
			_go_fuzz_dep_.CoverTab[24705]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:834
			// _ = "end of CoverTab[24705]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:834
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:834
		// _ = "end of CoverTab[24696]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:834
		_go_fuzz_dep_.CoverTab[24697]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:837
		if !isSupportedSignatureAlgorithm(certVerify.signatureAlgorithm, supportedSignatureAlgorithms()) {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:837
			_go_fuzz_dep_.CoverTab[24706]++
											c.sendAlert(alertIllegalParameter)
											return errors.New("tls: client certificate used with invalid signature algorithm")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:839
			// _ = "end of CoverTab[24706]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:840
			_go_fuzz_dep_.CoverTab[24707]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:840
			// _ = "end of CoverTab[24707]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:840
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:840
		// _ = "end of CoverTab[24697]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:840
		_go_fuzz_dep_.CoverTab[24698]++
										sigType, sigHash, err := typeAndHashFromSignatureScheme(certVerify.signatureAlgorithm)
										if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:842
			_go_fuzz_dep_.CoverTab[24708]++
											return c.sendAlert(alertInternalError)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:843
			// _ = "end of CoverTab[24708]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:844
			_go_fuzz_dep_.CoverTab[24709]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:844
			// _ = "end of CoverTab[24709]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:844
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:844
		// _ = "end of CoverTab[24698]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:844
		_go_fuzz_dep_.CoverTab[24699]++
										if sigType == signaturePKCS1v15 || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:845
			_go_fuzz_dep_.CoverTab[24710]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:845
			return sigHash == crypto.SHA1
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:845
			// _ = "end of CoverTab[24710]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:845
		}() {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:845
			_go_fuzz_dep_.CoverTab[24711]++
											c.sendAlert(alertIllegalParameter)
											return errors.New("tls: client certificate used with invalid signature algorithm")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:847
			// _ = "end of CoverTab[24711]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:848
			_go_fuzz_dep_.CoverTab[24712]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:848
			// _ = "end of CoverTab[24712]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:848
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:848
		// _ = "end of CoverTab[24699]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:848
		_go_fuzz_dep_.CoverTab[24700]++
										signed := signedMessage(sigHash, clientSignatureContext, hs.transcript)
										if err := verifyHandshakeSignature(sigType, c.peerCertificates[0].PublicKey,
			sigHash, signed, certVerify.signature); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:851
			_go_fuzz_dep_.CoverTab[24713]++
											c.sendAlert(alertDecryptError)
											return errors.New("tls: invalid signature by the client certificate: " + err.Error())
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:853
			// _ = "end of CoverTab[24713]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:854
			_go_fuzz_dep_.CoverTab[24714]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:854
			// _ = "end of CoverTab[24714]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:854
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:854
		// _ = "end of CoverTab[24700]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:854
		_go_fuzz_dep_.CoverTab[24701]++

										if err := transcriptMsg(certVerify, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:856
			_go_fuzz_dep_.CoverTab[24715]++
											return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:857
			// _ = "end of CoverTab[24715]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:858
			_go_fuzz_dep_.CoverTab[24716]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:858
			// _ = "end of CoverTab[24716]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:858
		}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:858
		// _ = "end of CoverTab[24701]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:859
		_go_fuzz_dep_.CoverTab[24717]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:859
		// _ = "end of CoverTab[24717]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:859
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:859
	// _ = "end of CoverTab[24675]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:859
	_go_fuzz_dep_.CoverTab[24676]++

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:863
	if err := hs.sendSessionTickets(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:863
		_go_fuzz_dep_.CoverTab[24718]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:864
		// _ = "end of CoverTab[24718]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:865
		_go_fuzz_dep_.CoverTab[24719]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:865
		// _ = "end of CoverTab[24719]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:865
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:865
	// _ = "end of CoverTab[24676]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:865
	_go_fuzz_dep_.CoverTab[24677]++

									return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:867
	// _ = "end of CoverTab[24677]"
}

func (hs *serverHandshakeStateTLS13) readClientFinished() error {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:870
	_go_fuzz_dep_.CoverTab[24720]++
									c := hs.c

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:874
	msg, err := c.readHandshake(nil)
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:875
		_go_fuzz_dep_.CoverTab[24724]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:876
		// _ = "end of CoverTab[24724]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:877
		_go_fuzz_dep_.CoverTab[24725]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:877
		// _ = "end of CoverTab[24725]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:877
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:877
	// _ = "end of CoverTab[24720]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:877
	_go_fuzz_dep_.CoverTab[24721]++

									finished, ok := msg.(*finishedMsg)
									if !ok {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:880
		_go_fuzz_dep_.CoverTab[24726]++
										c.sendAlert(alertUnexpectedMessage)
										return unexpectedMessageError(finished, msg)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:882
		// _ = "end of CoverTab[24726]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:883
		_go_fuzz_dep_.CoverTab[24727]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:883
		// _ = "end of CoverTab[24727]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:883
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:883
	// _ = "end of CoverTab[24721]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:883
	_go_fuzz_dep_.CoverTab[24722]++

									if !hmac.Equal(hs.clientFinished, finished.verifyData) {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:885
		_go_fuzz_dep_.CoverTab[24728]++
										c.sendAlert(alertDecryptError)
										return errors.New("tls: invalid client finished hash")
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:887
		// _ = "end of CoverTab[24728]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:888
		_go_fuzz_dep_.CoverTab[24729]++
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:888
		// _ = "end of CoverTab[24729]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:888
	}
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:888
	// _ = "end of CoverTab[24722]"
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:888
	_go_fuzz_dep_.CoverTab[24723]++

									c.in.setTrafficSecret(hs.suite, hs.trafficSecret)

									return nil
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:892
	// _ = "end of CoverTab[24723]"
}

//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:893
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/handshake_server_tls13.go:893
var _ = _go_fuzz_dep_.CoverTab
