// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:5
package tls

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:5
import (
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:5
)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:5
import (
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:5
)

import (
	"bytes"
	"context"
	"crypto"
	"crypto/ecdh"
	"crypto/hmac"
	"crypto/rsa"
	"errors"
	"hash"
	"time"
)

type clientHandshakeStateTLS13 struct {
	c		*Conn
	ctx		context.Context
	serverHello	*serverHelloMsg
	hello		*clientHelloMsg
	ecdheKey	*ecdh.PrivateKey

	session		*ClientSessionState
	earlySecret	[]byte
	binderKey	[]byte

	certReq		*certificateRequestMsgTLS13
	usingPSK	bool
	sentDummyCCS	bool
	suite		*cipherSuiteTLS13
	transcript	hash.Hash
	masterSecret	[]byte
	trafficSecret	[]byte	// client_application_traffic_secret_0
}

// handshake requires hs.c, hs.hello, hs.serverHello, hs.ecdheKey, and,
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:39
// optionally, hs.session, hs.earlySecret and hs.binderKey to be set.
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:41
func (hs *clientHandshakeStateTLS13) handshake() error {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:41
	_go_fuzz_dep_.CoverTab[22671]++
									c := hs.c

									if needFIPS() {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:44
		_go_fuzz_dep_.CoverTab[22688]++
										return errors.New("tls: internal error: TLS 1.3 reached in FIPS mode")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:45
		// _ = "end of CoverTab[22688]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:46
		_go_fuzz_dep_.CoverTab[22689]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:46
		// _ = "end of CoverTab[22689]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:46
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:46
	// _ = "end of CoverTab[22671]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:46
	_go_fuzz_dep_.CoverTab[22672]++

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:50
	if c.handshakes > 0 {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:50
		_go_fuzz_dep_.CoverTab[22690]++
										c.sendAlert(alertProtocolVersion)
										return errors.New("tls: server selected TLS 1.3 in a renegotiation")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:52
		// _ = "end of CoverTab[22690]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:53
		_go_fuzz_dep_.CoverTab[22691]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:53
		// _ = "end of CoverTab[22691]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:53
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:53
	// _ = "end of CoverTab[22672]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:53
	_go_fuzz_dep_.CoverTab[22673]++

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:56
	if hs.ecdheKey == nil || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:56
		_go_fuzz_dep_.CoverTab[22692]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:56
		return len(hs.hello.keyShares) != 1
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:56
		// _ = "end of CoverTab[22692]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:56
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:56
		_go_fuzz_dep_.CoverTab[22693]++
										return c.sendAlert(alertInternalError)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:57
		// _ = "end of CoverTab[22693]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:58
		_go_fuzz_dep_.CoverTab[22694]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:58
		// _ = "end of CoverTab[22694]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:58
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:58
	// _ = "end of CoverTab[22673]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:58
	_go_fuzz_dep_.CoverTab[22674]++

									if err := hs.checkServerHelloOrHRR(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:60
		_go_fuzz_dep_.CoverTab[22695]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:61
		// _ = "end of CoverTab[22695]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:62
		_go_fuzz_dep_.CoverTab[22696]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:62
		// _ = "end of CoverTab[22696]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:62
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:62
	// _ = "end of CoverTab[22674]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:62
	_go_fuzz_dep_.CoverTab[22675]++

									hs.transcript = hs.suite.hash.New()

									if err := transcriptMsg(hs.hello, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:66
		_go_fuzz_dep_.CoverTab[22697]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:67
		// _ = "end of CoverTab[22697]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:68
		_go_fuzz_dep_.CoverTab[22698]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:68
		// _ = "end of CoverTab[22698]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:68
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:68
	// _ = "end of CoverTab[22675]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:68
	_go_fuzz_dep_.CoverTab[22676]++

									if bytes.Equal(hs.serverHello.random, helloRetryRequestRandom) {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:70
		_go_fuzz_dep_.CoverTab[22699]++
										if err := hs.sendDummyChangeCipherSpec(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:71
			_go_fuzz_dep_.CoverTab[22701]++
											return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:72
			// _ = "end of CoverTab[22701]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:73
			_go_fuzz_dep_.CoverTab[22702]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:73
			// _ = "end of CoverTab[22702]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:73
		}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:73
		// _ = "end of CoverTab[22699]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:73
		_go_fuzz_dep_.CoverTab[22700]++
										if err := hs.processHelloRetryRequest(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:74
			_go_fuzz_dep_.CoverTab[22703]++
											return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:75
			// _ = "end of CoverTab[22703]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:76
			_go_fuzz_dep_.CoverTab[22704]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:76
			// _ = "end of CoverTab[22704]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:76
		}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:76
		// _ = "end of CoverTab[22700]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:77
		_go_fuzz_dep_.CoverTab[22705]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:77
		// _ = "end of CoverTab[22705]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:77
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:77
	// _ = "end of CoverTab[22676]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:77
	_go_fuzz_dep_.CoverTab[22677]++

									if err := transcriptMsg(hs.serverHello, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:79
		_go_fuzz_dep_.CoverTab[22706]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:80
		// _ = "end of CoverTab[22706]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:81
		_go_fuzz_dep_.CoverTab[22707]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:81
		// _ = "end of CoverTab[22707]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:81
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:81
	// _ = "end of CoverTab[22677]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:81
	_go_fuzz_dep_.CoverTab[22678]++

									c.buffering = true
									if err := hs.processServerHello(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:84
		_go_fuzz_dep_.CoverTab[22708]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:85
		// _ = "end of CoverTab[22708]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:86
		_go_fuzz_dep_.CoverTab[22709]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:86
		// _ = "end of CoverTab[22709]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:86
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:86
	// _ = "end of CoverTab[22678]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:86
	_go_fuzz_dep_.CoverTab[22679]++
									if err := hs.sendDummyChangeCipherSpec(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:87
		_go_fuzz_dep_.CoverTab[22710]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:88
		// _ = "end of CoverTab[22710]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:89
		_go_fuzz_dep_.CoverTab[22711]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:89
		// _ = "end of CoverTab[22711]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:89
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:89
	// _ = "end of CoverTab[22679]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:89
	_go_fuzz_dep_.CoverTab[22680]++
									if err := hs.establishHandshakeKeys(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:90
		_go_fuzz_dep_.CoverTab[22712]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:91
		// _ = "end of CoverTab[22712]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:92
		_go_fuzz_dep_.CoverTab[22713]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:92
		// _ = "end of CoverTab[22713]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:92
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:92
	// _ = "end of CoverTab[22680]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:92
	_go_fuzz_dep_.CoverTab[22681]++
									if err := hs.readServerParameters(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:93
		_go_fuzz_dep_.CoverTab[22714]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:94
		// _ = "end of CoverTab[22714]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:95
		_go_fuzz_dep_.CoverTab[22715]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:95
		// _ = "end of CoverTab[22715]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:95
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:95
	// _ = "end of CoverTab[22681]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:95
	_go_fuzz_dep_.CoverTab[22682]++
									if err := hs.readServerCertificate(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:96
		_go_fuzz_dep_.CoverTab[22716]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:97
		// _ = "end of CoverTab[22716]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:98
		_go_fuzz_dep_.CoverTab[22717]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:98
		// _ = "end of CoverTab[22717]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:98
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:98
	// _ = "end of CoverTab[22682]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:98
	_go_fuzz_dep_.CoverTab[22683]++
									if err := hs.readServerFinished(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:99
		_go_fuzz_dep_.CoverTab[22718]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:100
		// _ = "end of CoverTab[22718]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:101
		_go_fuzz_dep_.CoverTab[22719]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:101
		// _ = "end of CoverTab[22719]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:101
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:101
	// _ = "end of CoverTab[22683]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:101
	_go_fuzz_dep_.CoverTab[22684]++
									if err := hs.sendClientCertificate(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:102
		_go_fuzz_dep_.CoverTab[22720]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:103
		// _ = "end of CoverTab[22720]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:104
		_go_fuzz_dep_.CoverTab[22721]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:104
		// _ = "end of CoverTab[22721]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:104
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:104
	// _ = "end of CoverTab[22684]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:104
	_go_fuzz_dep_.CoverTab[22685]++
									if err := hs.sendClientFinished(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:105
		_go_fuzz_dep_.CoverTab[22722]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:106
		// _ = "end of CoverTab[22722]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:107
		_go_fuzz_dep_.CoverTab[22723]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:107
		// _ = "end of CoverTab[22723]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:107
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:107
	// _ = "end of CoverTab[22685]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:107
	_go_fuzz_dep_.CoverTab[22686]++
									if _, err := c.flush(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:108
		_go_fuzz_dep_.CoverTab[22724]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:109
		// _ = "end of CoverTab[22724]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:110
		_go_fuzz_dep_.CoverTab[22725]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:110
		// _ = "end of CoverTab[22725]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:110
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:110
	// _ = "end of CoverTab[22686]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:110
	_go_fuzz_dep_.CoverTab[22687]++

									c.isHandshakeComplete.Store(true)

									return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:114
	// _ = "end of CoverTab[22687]"
}

// checkServerHelloOrHRR does validity checks that apply to both ServerHello and
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:117
// HelloRetryRequest messages. It sets hs.suite.
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:119
func (hs *clientHandshakeStateTLS13) checkServerHelloOrHRR() error {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:119
	_go_fuzz_dep_.CoverTab[22726]++
									c := hs.c

									if hs.serverHello.supportedVersion == 0 {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:122
		_go_fuzz_dep_.CoverTab[22735]++
										c.sendAlert(alertMissingExtension)
										return errors.New("tls: server selected TLS 1.3 using the legacy version field")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:124
		// _ = "end of CoverTab[22735]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:125
		_go_fuzz_dep_.CoverTab[22736]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:125
		// _ = "end of CoverTab[22736]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:125
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:125
	// _ = "end of CoverTab[22726]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:125
	_go_fuzz_dep_.CoverTab[22727]++

									if hs.serverHello.supportedVersion != VersionTLS13 {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:127
		_go_fuzz_dep_.CoverTab[22737]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: server selected an invalid version after a HelloRetryRequest")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:129
		// _ = "end of CoverTab[22737]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:130
		_go_fuzz_dep_.CoverTab[22738]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:130
		// _ = "end of CoverTab[22738]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:130
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:130
	// _ = "end of CoverTab[22727]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:130
	_go_fuzz_dep_.CoverTab[22728]++

									if hs.serverHello.vers != VersionTLS12 {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:132
		_go_fuzz_dep_.CoverTab[22739]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: server sent an incorrect legacy version")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:134
		// _ = "end of CoverTab[22739]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:135
		_go_fuzz_dep_.CoverTab[22740]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:135
		// _ = "end of CoverTab[22740]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:135
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:135
	// _ = "end of CoverTab[22728]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:135
	_go_fuzz_dep_.CoverTab[22729]++

									if hs.serverHello.ocspStapling || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:137
		_go_fuzz_dep_.CoverTab[22741]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:137
		return hs.serverHello.ticketSupported
										// _ = "end of CoverTab[22741]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:138
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:138
		_go_fuzz_dep_.CoverTab[22742]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:138
		return hs.serverHello.secureRenegotiationSupported
										// _ = "end of CoverTab[22742]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:139
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:139
		_go_fuzz_dep_.CoverTab[22743]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:139
		return len(hs.serverHello.secureRenegotiation) != 0
										// _ = "end of CoverTab[22743]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:140
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:140
		_go_fuzz_dep_.CoverTab[22744]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:140
		return len(hs.serverHello.alpnProtocol) != 0
										// _ = "end of CoverTab[22744]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:141
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:141
		_go_fuzz_dep_.CoverTab[22745]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:141
		return len(hs.serverHello.scts) != 0
										// _ = "end of CoverTab[22745]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:142
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:142
		_go_fuzz_dep_.CoverTab[22746]++
										c.sendAlert(alertUnsupportedExtension)
										return errors.New("tls: server sent a ServerHello extension forbidden in TLS 1.3")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:144
		// _ = "end of CoverTab[22746]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:145
		_go_fuzz_dep_.CoverTab[22747]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:145
		// _ = "end of CoverTab[22747]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:145
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:145
	// _ = "end of CoverTab[22729]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:145
	_go_fuzz_dep_.CoverTab[22730]++

									if !bytes.Equal(hs.hello.sessionId, hs.serverHello.sessionId) {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:147
		_go_fuzz_dep_.CoverTab[22748]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: server did not echo the legacy session ID")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:149
		// _ = "end of CoverTab[22748]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:150
		_go_fuzz_dep_.CoverTab[22749]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:150
		// _ = "end of CoverTab[22749]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:150
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:150
	// _ = "end of CoverTab[22730]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:150
	_go_fuzz_dep_.CoverTab[22731]++

									if hs.serverHello.compressionMethod != compressionNone {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:152
		_go_fuzz_dep_.CoverTab[22750]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: server selected unsupported compression format")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:154
		// _ = "end of CoverTab[22750]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:155
		_go_fuzz_dep_.CoverTab[22751]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:155
		// _ = "end of CoverTab[22751]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:155
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:155
	// _ = "end of CoverTab[22731]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:155
	_go_fuzz_dep_.CoverTab[22732]++

									selectedSuite := mutualCipherSuiteTLS13(hs.hello.cipherSuites, hs.serverHello.cipherSuite)
									if hs.suite != nil && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:158
		_go_fuzz_dep_.CoverTab[22752]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:158
		return selectedSuite != hs.suite
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:158
		// _ = "end of CoverTab[22752]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:158
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:158
		_go_fuzz_dep_.CoverTab[22753]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: server changed cipher suite after a HelloRetryRequest")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:160
		// _ = "end of CoverTab[22753]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:161
		_go_fuzz_dep_.CoverTab[22754]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:161
		// _ = "end of CoverTab[22754]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:161
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:161
	// _ = "end of CoverTab[22732]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:161
	_go_fuzz_dep_.CoverTab[22733]++
									if selectedSuite == nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:162
		_go_fuzz_dep_.CoverTab[22755]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: server chose an unconfigured cipher suite")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:164
		// _ = "end of CoverTab[22755]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:165
		_go_fuzz_dep_.CoverTab[22756]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:165
		// _ = "end of CoverTab[22756]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:165
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:165
	// _ = "end of CoverTab[22733]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:165
	_go_fuzz_dep_.CoverTab[22734]++
									hs.suite = selectedSuite
									c.cipherSuite = hs.suite.id

									return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:169
	// _ = "end of CoverTab[22734]"
}

// sendDummyChangeCipherSpec sends a ChangeCipherSpec record for compatibility
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:172
// with middleboxes that didn't implement TLS correctly. See RFC 8446, Appendix D.4.
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:174
func (hs *clientHandshakeStateTLS13) sendDummyChangeCipherSpec() error {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:174
	_go_fuzz_dep_.CoverTab[22757]++
									if hs.sentDummyCCS {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:175
		_go_fuzz_dep_.CoverTab[22759]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:176
		// _ = "end of CoverTab[22759]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:177
		_go_fuzz_dep_.CoverTab[22760]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:177
		// _ = "end of CoverTab[22760]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:177
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:177
	// _ = "end of CoverTab[22757]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:177
	_go_fuzz_dep_.CoverTab[22758]++
									hs.sentDummyCCS = true

									return hs.c.writeChangeCipherRecord()
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:180
	// _ = "end of CoverTab[22758]"
}

// processHelloRetryRequest handles the HRR in hs.serverHello, modifies and
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:183
// resends hs.hello, and reads the new ServerHello into hs.serverHello.
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:185
func (hs *clientHandshakeStateTLS13) processHelloRetryRequest() error {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:185
	_go_fuzz_dep_.CoverTab[22761]++
									c := hs.c

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:191
	chHash := hs.transcript.Sum(nil)
	hs.transcript.Reset()
	hs.transcript.Write([]byte{typeMessageHash, 0, 0, uint8(len(chHash))})
	hs.transcript.Write(chHash)
	if err := transcriptMsg(hs.serverHello, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:195
		_go_fuzz_dep_.CoverTab[22772]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:196
		// _ = "end of CoverTab[22772]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:197
		_go_fuzz_dep_.CoverTab[22773]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:197
		// _ = "end of CoverTab[22773]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:197
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:197
	// _ = "end of CoverTab[22761]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:197
	_go_fuzz_dep_.CoverTab[22762]++

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:202
	if hs.serverHello.selectedGroup == 0 && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:202
		_go_fuzz_dep_.CoverTab[22774]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:202
		return hs.serverHello.cookie == nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:202
		// _ = "end of CoverTab[22774]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:202
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:202
		_go_fuzz_dep_.CoverTab[22775]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: server sent an unnecessary HelloRetryRequest message")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:204
		// _ = "end of CoverTab[22775]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:205
		_go_fuzz_dep_.CoverTab[22776]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:205
		// _ = "end of CoverTab[22776]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:205
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:205
	// _ = "end of CoverTab[22762]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:205
	_go_fuzz_dep_.CoverTab[22763]++

									if hs.serverHello.cookie != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:207
		_go_fuzz_dep_.CoverTab[22777]++
										hs.hello.cookie = hs.serverHello.cookie
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:208
		// _ = "end of CoverTab[22777]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:209
		_go_fuzz_dep_.CoverTab[22778]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:209
		// _ = "end of CoverTab[22778]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:209
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:209
	// _ = "end of CoverTab[22763]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:209
	_go_fuzz_dep_.CoverTab[22764]++

									if hs.serverHello.serverShare.group != 0 {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:211
		_go_fuzz_dep_.CoverTab[22779]++
										c.sendAlert(alertDecodeError)
										return errors.New("tls: received malformed key_share extension")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:213
		// _ = "end of CoverTab[22779]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:214
		_go_fuzz_dep_.CoverTab[22780]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:214
		// _ = "end of CoverTab[22780]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:214
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:214
	// _ = "end of CoverTab[22764]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:214
	_go_fuzz_dep_.CoverTab[22765]++

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:219
	if curveID := hs.serverHello.selectedGroup; curveID != 0 {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:219
		_go_fuzz_dep_.CoverTab[22781]++
										curveOK := false
										for _, id := range hs.hello.supportedCurves {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:221
			_go_fuzz_dep_.CoverTab[22787]++
											if id == curveID {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:222
				_go_fuzz_dep_.CoverTab[22788]++
												curveOK = true
												break
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:224
				// _ = "end of CoverTab[22788]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:225
				_go_fuzz_dep_.CoverTab[22789]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:225
				// _ = "end of CoverTab[22789]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:225
			}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:225
			// _ = "end of CoverTab[22787]"
		}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:226
		// _ = "end of CoverTab[22781]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:226
		_go_fuzz_dep_.CoverTab[22782]++
										if !curveOK {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:227
			_go_fuzz_dep_.CoverTab[22790]++
											c.sendAlert(alertIllegalParameter)
											return errors.New("tls: server selected unsupported group")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:229
			// _ = "end of CoverTab[22790]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:230
			_go_fuzz_dep_.CoverTab[22791]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:230
			// _ = "end of CoverTab[22791]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:230
		}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:230
		// _ = "end of CoverTab[22782]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:230
		_go_fuzz_dep_.CoverTab[22783]++
										if sentID, _ := curveIDForCurve(hs.ecdheKey.Curve()); sentID == curveID {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:231
			_go_fuzz_dep_.CoverTab[22792]++
											c.sendAlert(alertIllegalParameter)
											return errors.New("tls: server sent an unnecessary HelloRetryRequest key_share")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:233
			// _ = "end of CoverTab[22792]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:234
			_go_fuzz_dep_.CoverTab[22793]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:234
			// _ = "end of CoverTab[22793]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:234
		}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:234
		// _ = "end of CoverTab[22783]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:234
		_go_fuzz_dep_.CoverTab[22784]++
										if _, ok := curveForCurveID(curveID); !ok {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:235
			_go_fuzz_dep_.CoverTab[22794]++
											c.sendAlert(alertInternalError)
											return errors.New("tls: CurvePreferences includes unsupported curve")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:237
			// _ = "end of CoverTab[22794]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:238
			_go_fuzz_dep_.CoverTab[22795]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:238
			// _ = "end of CoverTab[22795]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:238
		}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:238
		// _ = "end of CoverTab[22784]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:238
		_go_fuzz_dep_.CoverTab[22785]++
										key, err := generateECDHEKey(c.config.rand(), curveID)
										if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:240
			_go_fuzz_dep_.CoverTab[22796]++
											c.sendAlert(alertInternalError)
											return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:242
			// _ = "end of CoverTab[22796]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:243
			_go_fuzz_dep_.CoverTab[22797]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:243
			// _ = "end of CoverTab[22797]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:243
		}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:243
		// _ = "end of CoverTab[22785]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:243
		_go_fuzz_dep_.CoverTab[22786]++
										hs.ecdheKey = key
										hs.hello.keyShares = []keyShare{{group: curveID, data: key.PublicKey().Bytes()}}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:245
		// _ = "end of CoverTab[22786]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:246
		_go_fuzz_dep_.CoverTab[22798]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:246
		// _ = "end of CoverTab[22798]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:246
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:246
	// _ = "end of CoverTab[22765]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:246
	_go_fuzz_dep_.CoverTab[22766]++

									hs.hello.raw = nil
									if len(hs.hello.pskIdentities) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:249
		_go_fuzz_dep_.CoverTab[22799]++
										pskSuite := cipherSuiteTLS13ByID(hs.session.cipherSuite)
										if pskSuite == nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:251
			_go_fuzz_dep_.CoverTab[22801]++
											return c.sendAlert(alertInternalError)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:252
			// _ = "end of CoverTab[22801]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:253
			_go_fuzz_dep_.CoverTab[22802]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:253
			// _ = "end of CoverTab[22802]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:253
		}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:253
		// _ = "end of CoverTab[22799]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:253
		_go_fuzz_dep_.CoverTab[22800]++
										if pskSuite.hash == hs.suite.hash {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:254
			_go_fuzz_dep_.CoverTab[22803]++

											ticketAge := uint32(c.config.time().Sub(hs.session.receivedAt) / time.Millisecond)
											hs.hello.pskIdentities[0].obfuscatedTicketAge = ticketAge + hs.session.ageAdd

											transcript := hs.suite.hash.New()
											transcript.Write([]byte{typeMessageHash, 0, 0, uint8(len(chHash))})
											transcript.Write(chHash)
											if err := transcriptMsg(hs.serverHello, transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:262
				_go_fuzz_dep_.CoverTab[22806]++
												return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:263
				// _ = "end of CoverTab[22806]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:264
				_go_fuzz_dep_.CoverTab[22807]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:264
				// _ = "end of CoverTab[22807]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:264
			}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:264
			// _ = "end of CoverTab[22803]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:264
			_go_fuzz_dep_.CoverTab[22804]++
											helloBytes, err := hs.hello.marshalWithoutBinders()
											if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:266
				_go_fuzz_dep_.CoverTab[22808]++
												return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:267
				// _ = "end of CoverTab[22808]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:268
				_go_fuzz_dep_.CoverTab[22809]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:268
				// _ = "end of CoverTab[22809]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:268
			}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:268
			// _ = "end of CoverTab[22804]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:268
			_go_fuzz_dep_.CoverTab[22805]++
											transcript.Write(helloBytes)
											pskBinders := [][]byte{hs.suite.finishedHash(hs.binderKey, transcript)}
											if err := hs.hello.updateBinders(pskBinders); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:271
				_go_fuzz_dep_.CoverTab[22810]++
												return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:272
				// _ = "end of CoverTab[22810]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:273
				_go_fuzz_dep_.CoverTab[22811]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:273
				// _ = "end of CoverTab[22811]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:273
			}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:273
			// _ = "end of CoverTab[22805]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:274
			_go_fuzz_dep_.CoverTab[22812]++

											hs.hello.pskIdentities = nil
											hs.hello.pskBinders = nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:277
			// _ = "end of CoverTab[22812]"
		}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:278
		// _ = "end of CoverTab[22800]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:279
		_go_fuzz_dep_.CoverTab[22813]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:279
		// _ = "end of CoverTab[22813]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:279
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:279
	// _ = "end of CoverTab[22766]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:279
	_go_fuzz_dep_.CoverTab[22767]++

									if _, err := hs.c.writeHandshakeRecord(hs.hello, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:281
		_go_fuzz_dep_.CoverTab[22814]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:282
		// _ = "end of CoverTab[22814]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:283
		_go_fuzz_dep_.CoverTab[22815]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:283
		// _ = "end of CoverTab[22815]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:283
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:283
	// _ = "end of CoverTab[22767]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:283
	_go_fuzz_dep_.CoverTab[22768]++

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:286
	msg, err := c.readHandshake(nil)
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:287
		_go_fuzz_dep_.CoverTab[22816]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:288
		// _ = "end of CoverTab[22816]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:289
		_go_fuzz_dep_.CoverTab[22817]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:289
		// _ = "end of CoverTab[22817]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:289
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:289
	// _ = "end of CoverTab[22768]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:289
	_go_fuzz_dep_.CoverTab[22769]++

									serverHello, ok := msg.(*serverHelloMsg)
									if !ok {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:292
		_go_fuzz_dep_.CoverTab[22818]++
										c.sendAlert(alertUnexpectedMessage)
										return unexpectedMessageError(serverHello, msg)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:294
		// _ = "end of CoverTab[22818]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:295
		_go_fuzz_dep_.CoverTab[22819]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:295
		// _ = "end of CoverTab[22819]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:295
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:295
	// _ = "end of CoverTab[22769]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:295
	_go_fuzz_dep_.CoverTab[22770]++
									hs.serverHello = serverHello

									if err := hs.checkServerHelloOrHRR(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:298
		_go_fuzz_dep_.CoverTab[22820]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:299
		// _ = "end of CoverTab[22820]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:300
		_go_fuzz_dep_.CoverTab[22821]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:300
		// _ = "end of CoverTab[22821]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:300
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:300
	// _ = "end of CoverTab[22770]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:300
	_go_fuzz_dep_.CoverTab[22771]++

									return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:302
	// _ = "end of CoverTab[22771]"
}

func (hs *clientHandshakeStateTLS13) processServerHello() error {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:305
	_go_fuzz_dep_.CoverTab[22822]++
									c := hs.c

									if bytes.Equal(hs.serverHello.random, helloRetryRequestRandom) {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:308
		_go_fuzz_dep_.CoverTab[22833]++
										c.sendAlert(alertUnexpectedMessage)
										return errors.New("tls: server sent two HelloRetryRequest messages")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:310
		// _ = "end of CoverTab[22833]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:311
		_go_fuzz_dep_.CoverTab[22834]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:311
		// _ = "end of CoverTab[22834]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:311
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:311
	// _ = "end of CoverTab[22822]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:311
	_go_fuzz_dep_.CoverTab[22823]++

									if len(hs.serverHello.cookie) != 0 {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:313
		_go_fuzz_dep_.CoverTab[22835]++
										c.sendAlert(alertUnsupportedExtension)
										return errors.New("tls: server sent a cookie in a normal ServerHello")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:315
		// _ = "end of CoverTab[22835]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:316
		_go_fuzz_dep_.CoverTab[22836]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:316
		// _ = "end of CoverTab[22836]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:316
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:316
	// _ = "end of CoverTab[22823]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:316
	_go_fuzz_dep_.CoverTab[22824]++

									if hs.serverHello.selectedGroup != 0 {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:318
		_go_fuzz_dep_.CoverTab[22837]++
										c.sendAlert(alertDecodeError)
										return errors.New("tls: malformed key_share extension")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:320
		// _ = "end of CoverTab[22837]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:321
		_go_fuzz_dep_.CoverTab[22838]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:321
		// _ = "end of CoverTab[22838]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:321
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:321
	// _ = "end of CoverTab[22824]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:321
	_go_fuzz_dep_.CoverTab[22825]++

									if hs.serverHello.serverShare.group == 0 {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:323
		_go_fuzz_dep_.CoverTab[22839]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: server did not send a key share")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:325
		// _ = "end of CoverTab[22839]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:326
		_go_fuzz_dep_.CoverTab[22840]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:326
		// _ = "end of CoverTab[22840]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:326
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:326
	// _ = "end of CoverTab[22825]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:326
	_go_fuzz_dep_.CoverTab[22826]++
									if sentID, _ := curveIDForCurve(hs.ecdheKey.Curve()); hs.serverHello.serverShare.group != sentID {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:327
		_go_fuzz_dep_.CoverTab[22841]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: server selected unsupported group")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:329
		// _ = "end of CoverTab[22841]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:330
		_go_fuzz_dep_.CoverTab[22842]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:330
		// _ = "end of CoverTab[22842]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:330
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:330
	// _ = "end of CoverTab[22826]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:330
	_go_fuzz_dep_.CoverTab[22827]++

									if !hs.serverHello.selectedIdentityPresent {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:332
		_go_fuzz_dep_.CoverTab[22843]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:333
		// _ = "end of CoverTab[22843]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:334
		_go_fuzz_dep_.CoverTab[22844]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:334
		// _ = "end of CoverTab[22844]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:334
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:334
	// _ = "end of CoverTab[22827]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:334
	_go_fuzz_dep_.CoverTab[22828]++

									if int(hs.serverHello.selectedIdentity) >= len(hs.hello.pskIdentities) {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:336
		_go_fuzz_dep_.CoverTab[22845]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: server selected an invalid PSK")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:338
		// _ = "end of CoverTab[22845]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:339
		_go_fuzz_dep_.CoverTab[22846]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:339
		// _ = "end of CoverTab[22846]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:339
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:339
	// _ = "end of CoverTab[22828]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:339
	_go_fuzz_dep_.CoverTab[22829]++

									if len(hs.hello.pskIdentities) != 1 || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:341
		_go_fuzz_dep_.CoverTab[22847]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:341
		return hs.session == nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:341
		// _ = "end of CoverTab[22847]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:341
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:341
		_go_fuzz_dep_.CoverTab[22848]++
										return c.sendAlert(alertInternalError)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:342
		// _ = "end of CoverTab[22848]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:343
		_go_fuzz_dep_.CoverTab[22849]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:343
		// _ = "end of CoverTab[22849]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:343
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:343
	// _ = "end of CoverTab[22829]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:343
	_go_fuzz_dep_.CoverTab[22830]++
									pskSuite := cipherSuiteTLS13ByID(hs.session.cipherSuite)
									if pskSuite == nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:345
		_go_fuzz_dep_.CoverTab[22850]++
										return c.sendAlert(alertInternalError)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:346
		// _ = "end of CoverTab[22850]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:347
		_go_fuzz_dep_.CoverTab[22851]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:347
		// _ = "end of CoverTab[22851]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:347
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:347
	// _ = "end of CoverTab[22830]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:347
	_go_fuzz_dep_.CoverTab[22831]++
									if pskSuite.hash != hs.suite.hash {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:348
		_go_fuzz_dep_.CoverTab[22852]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: server selected an invalid PSK and cipher suite pair")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:350
		// _ = "end of CoverTab[22852]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:351
		_go_fuzz_dep_.CoverTab[22853]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:351
		// _ = "end of CoverTab[22853]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:351
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:351
	// _ = "end of CoverTab[22831]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:351
	_go_fuzz_dep_.CoverTab[22832]++

									hs.usingPSK = true
									c.didResume = true
									c.peerCertificates = hs.session.serverCertificates
									c.verifiedChains = hs.session.verifiedChains
									c.ocspResponse = hs.session.ocspResponse
									c.scts = hs.session.scts
									return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:359
	// _ = "end of CoverTab[22832]"
}

func (hs *clientHandshakeStateTLS13) establishHandshakeKeys() error {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:362
	_go_fuzz_dep_.CoverTab[22854]++
									c := hs.c

									peerKey, err := hs.ecdheKey.Curve().NewPublicKey(hs.serverHello.serverShare.data)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:366
		_go_fuzz_dep_.CoverTab[22860]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: invalid server key share")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:368
		// _ = "end of CoverTab[22860]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:369
		_go_fuzz_dep_.CoverTab[22861]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:369
		// _ = "end of CoverTab[22861]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:369
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:369
	// _ = "end of CoverTab[22854]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:369
	_go_fuzz_dep_.CoverTab[22855]++
									sharedKey, err := hs.ecdheKey.ECDH(peerKey)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:371
		_go_fuzz_dep_.CoverTab[22862]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: invalid server key share")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:373
		// _ = "end of CoverTab[22862]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:374
		_go_fuzz_dep_.CoverTab[22863]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:374
		// _ = "end of CoverTab[22863]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:374
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:374
	// _ = "end of CoverTab[22855]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:374
	_go_fuzz_dep_.CoverTab[22856]++

									earlySecret := hs.earlySecret
									if !hs.usingPSK {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:377
		_go_fuzz_dep_.CoverTab[22864]++
										earlySecret = hs.suite.extract(nil, nil)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:378
		// _ = "end of CoverTab[22864]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:379
		_go_fuzz_dep_.CoverTab[22865]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:379
		// _ = "end of CoverTab[22865]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:379
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:379
	// _ = "end of CoverTab[22856]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:379
	_go_fuzz_dep_.CoverTab[22857]++

									handshakeSecret := hs.suite.extract(sharedKey,
		hs.suite.deriveSecret(earlySecret, "derived", nil))

	clientSecret := hs.suite.deriveSecret(handshakeSecret,
		clientHandshakeTrafficLabel, hs.transcript)
	c.out.setTrafficSecret(hs.suite, clientSecret)
	serverSecret := hs.suite.deriveSecret(handshakeSecret,
		serverHandshakeTrafficLabel, hs.transcript)
	c.in.setTrafficSecret(hs.suite, serverSecret)

	err = c.config.writeKeyLog(keyLogLabelClientHandshake, hs.hello.random, clientSecret)
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:392
		_go_fuzz_dep_.CoverTab[22866]++
										c.sendAlert(alertInternalError)
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:394
		// _ = "end of CoverTab[22866]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:395
		_go_fuzz_dep_.CoverTab[22867]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:395
		// _ = "end of CoverTab[22867]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:395
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:395
	// _ = "end of CoverTab[22857]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:395
	_go_fuzz_dep_.CoverTab[22858]++
									err = c.config.writeKeyLog(keyLogLabelServerHandshake, hs.hello.random, serverSecret)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:397
		_go_fuzz_dep_.CoverTab[22868]++
										c.sendAlert(alertInternalError)
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:399
		// _ = "end of CoverTab[22868]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:400
		_go_fuzz_dep_.CoverTab[22869]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:400
		// _ = "end of CoverTab[22869]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:400
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:400
	// _ = "end of CoverTab[22858]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:400
	_go_fuzz_dep_.CoverTab[22859]++

									hs.masterSecret = hs.suite.extract(nil,
		hs.suite.deriveSecret(handshakeSecret, "derived", nil))

									return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:405
	// _ = "end of CoverTab[22859]"
}

func (hs *clientHandshakeStateTLS13) readServerParameters() error {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:408
	_go_fuzz_dep_.CoverTab[22870]++
									c := hs.c

									msg, err := c.readHandshake(hs.transcript)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:412
		_go_fuzz_dep_.CoverTab[22874]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:413
		// _ = "end of CoverTab[22874]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:414
		_go_fuzz_dep_.CoverTab[22875]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:414
		// _ = "end of CoverTab[22875]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:414
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:414
	// _ = "end of CoverTab[22870]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:414
	_go_fuzz_dep_.CoverTab[22871]++

									encryptedExtensions, ok := msg.(*encryptedExtensionsMsg)
									if !ok {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:417
		_go_fuzz_dep_.CoverTab[22876]++
										c.sendAlert(alertUnexpectedMessage)
										return unexpectedMessageError(encryptedExtensions, msg)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:419
		// _ = "end of CoverTab[22876]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:420
		_go_fuzz_dep_.CoverTab[22877]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:420
		// _ = "end of CoverTab[22877]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:420
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:420
	// _ = "end of CoverTab[22871]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:420
	_go_fuzz_dep_.CoverTab[22872]++

									if err := checkALPN(hs.hello.alpnProtocols, encryptedExtensions.alpnProtocol); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:422
		_go_fuzz_dep_.CoverTab[22878]++
										c.sendAlert(alertUnsupportedExtension)
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:424
		// _ = "end of CoverTab[22878]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:425
		_go_fuzz_dep_.CoverTab[22879]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:425
		// _ = "end of CoverTab[22879]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:425
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:425
	// _ = "end of CoverTab[22872]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:425
	_go_fuzz_dep_.CoverTab[22873]++
									c.clientProtocol = encryptedExtensions.alpnProtocol

									return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:428
	// _ = "end of CoverTab[22873]"
}

func (hs *clientHandshakeStateTLS13) readServerCertificate() error {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:431
	_go_fuzz_dep_.CoverTab[22880]++
									c := hs.c

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:436
	if hs.usingPSK {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:436
		_go_fuzz_dep_.CoverTab[22894]++

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:440
		if c.config.VerifyConnection != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:440
			_go_fuzz_dep_.CoverTab[22896]++
											if err := c.config.VerifyConnection(c.connectionStateLocked()); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:441
				_go_fuzz_dep_.CoverTab[22897]++
												c.sendAlert(alertBadCertificate)
												return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:443
				// _ = "end of CoverTab[22897]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:444
				_go_fuzz_dep_.CoverTab[22898]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:444
				// _ = "end of CoverTab[22898]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:444
			}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:444
			// _ = "end of CoverTab[22896]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:445
			_go_fuzz_dep_.CoverTab[22899]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:445
			// _ = "end of CoverTab[22899]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:445
		}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:445
		// _ = "end of CoverTab[22894]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:445
		_go_fuzz_dep_.CoverTab[22895]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:446
		// _ = "end of CoverTab[22895]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:447
		_go_fuzz_dep_.CoverTab[22900]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:447
		// _ = "end of CoverTab[22900]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:447
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:447
	// _ = "end of CoverTab[22880]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:447
	_go_fuzz_dep_.CoverTab[22881]++

									msg, err := c.readHandshake(hs.transcript)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:450
		_go_fuzz_dep_.CoverTab[22901]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:451
		// _ = "end of CoverTab[22901]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:452
		_go_fuzz_dep_.CoverTab[22902]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:452
		// _ = "end of CoverTab[22902]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:452
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:452
	// _ = "end of CoverTab[22881]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:452
	_go_fuzz_dep_.CoverTab[22882]++

									certReq, ok := msg.(*certificateRequestMsgTLS13)
									if ok {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:455
		_go_fuzz_dep_.CoverTab[22903]++
										hs.certReq = certReq

										msg, err = c.readHandshake(hs.transcript)
										if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:459
			_go_fuzz_dep_.CoverTab[22904]++
											return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:460
			// _ = "end of CoverTab[22904]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:461
			_go_fuzz_dep_.CoverTab[22905]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:461
			// _ = "end of CoverTab[22905]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:461
		}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:461
		// _ = "end of CoverTab[22903]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:462
		_go_fuzz_dep_.CoverTab[22906]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:462
		// _ = "end of CoverTab[22906]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:462
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:462
	// _ = "end of CoverTab[22882]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:462
	_go_fuzz_dep_.CoverTab[22883]++

									certMsg, ok := msg.(*certificateMsgTLS13)
									if !ok {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:465
		_go_fuzz_dep_.CoverTab[22907]++
										c.sendAlert(alertUnexpectedMessage)
										return unexpectedMessageError(certMsg, msg)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:467
		// _ = "end of CoverTab[22907]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:468
		_go_fuzz_dep_.CoverTab[22908]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:468
		// _ = "end of CoverTab[22908]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:468
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:468
	// _ = "end of CoverTab[22883]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:468
	_go_fuzz_dep_.CoverTab[22884]++
									if len(certMsg.certificate.Certificate) == 0 {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:469
		_go_fuzz_dep_.CoverTab[22909]++
										c.sendAlert(alertDecodeError)
										return errors.New("tls: received empty certificates message")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:471
		// _ = "end of CoverTab[22909]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:472
		_go_fuzz_dep_.CoverTab[22910]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:472
		// _ = "end of CoverTab[22910]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:472
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:472
	// _ = "end of CoverTab[22884]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:472
	_go_fuzz_dep_.CoverTab[22885]++

									c.scts = certMsg.certificate.SignedCertificateTimestamps
									c.ocspResponse = certMsg.certificate.OCSPStaple

									if err := c.verifyServerCertificate(certMsg.certificate.Certificate); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:477
		_go_fuzz_dep_.CoverTab[22911]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:478
		// _ = "end of CoverTab[22911]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:479
		_go_fuzz_dep_.CoverTab[22912]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:479
		// _ = "end of CoverTab[22912]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:479
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:479
	// _ = "end of CoverTab[22885]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:479
	_go_fuzz_dep_.CoverTab[22886]++

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:484
	msg, err = c.readHandshake(nil)
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:485
		_go_fuzz_dep_.CoverTab[22913]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:486
		// _ = "end of CoverTab[22913]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:487
		_go_fuzz_dep_.CoverTab[22914]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:487
		// _ = "end of CoverTab[22914]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:487
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:487
	// _ = "end of CoverTab[22886]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:487
	_go_fuzz_dep_.CoverTab[22887]++

									certVerify, ok := msg.(*certificateVerifyMsg)
									if !ok {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:490
		_go_fuzz_dep_.CoverTab[22915]++
										c.sendAlert(alertUnexpectedMessage)
										return unexpectedMessageError(certVerify, msg)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:492
		// _ = "end of CoverTab[22915]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:493
		_go_fuzz_dep_.CoverTab[22916]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:493
		// _ = "end of CoverTab[22916]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:493
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:493
	// _ = "end of CoverTab[22887]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:493
	_go_fuzz_dep_.CoverTab[22888]++

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:496
	if !isSupportedSignatureAlgorithm(certVerify.signatureAlgorithm, supportedSignatureAlgorithms()) {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:496
		_go_fuzz_dep_.CoverTab[22917]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: certificate used with invalid signature algorithm")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:498
		// _ = "end of CoverTab[22917]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:499
		_go_fuzz_dep_.CoverTab[22918]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:499
		// _ = "end of CoverTab[22918]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:499
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:499
	// _ = "end of CoverTab[22888]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:499
	_go_fuzz_dep_.CoverTab[22889]++
									sigType, sigHash, err := typeAndHashFromSignatureScheme(certVerify.signatureAlgorithm)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:501
		_go_fuzz_dep_.CoverTab[22919]++
										return c.sendAlert(alertInternalError)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:502
		// _ = "end of CoverTab[22919]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:503
		_go_fuzz_dep_.CoverTab[22920]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:503
		// _ = "end of CoverTab[22920]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:503
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:503
	// _ = "end of CoverTab[22889]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:503
	_go_fuzz_dep_.CoverTab[22890]++
									if sigType == signaturePKCS1v15 || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:504
		_go_fuzz_dep_.CoverTab[22921]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:504
		return sigHash == crypto.SHA1
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:504
		// _ = "end of CoverTab[22921]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:504
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:504
		_go_fuzz_dep_.CoverTab[22922]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: certificate used with invalid signature algorithm")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:506
		// _ = "end of CoverTab[22922]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:507
		_go_fuzz_dep_.CoverTab[22923]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:507
		// _ = "end of CoverTab[22923]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:507
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:507
	// _ = "end of CoverTab[22890]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:507
	_go_fuzz_dep_.CoverTab[22891]++
									signed := signedMessage(sigHash, serverSignatureContext, hs.transcript)
									if err := verifyHandshakeSignature(sigType, c.peerCertificates[0].PublicKey,
		sigHash, signed, certVerify.signature); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:510
		_go_fuzz_dep_.CoverTab[22924]++
										c.sendAlert(alertDecryptError)
										return errors.New("tls: invalid signature by the server certificate: " + err.Error())
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:512
		// _ = "end of CoverTab[22924]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:513
		_go_fuzz_dep_.CoverTab[22925]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:513
		// _ = "end of CoverTab[22925]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:513
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:513
	// _ = "end of CoverTab[22891]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:513
	_go_fuzz_dep_.CoverTab[22892]++

									if err := transcriptMsg(certVerify, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:515
		_go_fuzz_dep_.CoverTab[22926]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:516
		// _ = "end of CoverTab[22926]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:517
		_go_fuzz_dep_.CoverTab[22927]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:517
		// _ = "end of CoverTab[22927]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:517
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:517
	// _ = "end of CoverTab[22892]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:517
	_go_fuzz_dep_.CoverTab[22893]++

									return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:519
	// _ = "end of CoverTab[22893]"
}

func (hs *clientHandshakeStateTLS13) readServerFinished() error {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:522
	_go_fuzz_dep_.CoverTab[22928]++
									c := hs.c

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:528
	msg, err := c.readHandshake(nil)
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:529
		_go_fuzz_dep_.CoverTab[22935]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:530
		// _ = "end of CoverTab[22935]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:531
		_go_fuzz_dep_.CoverTab[22936]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:531
		// _ = "end of CoverTab[22936]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:531
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:531
	// _ = "end of CoverTab[22928]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:531
	_go_fuzz_dep_.CoverTab[22929]++

									finished, ok := msg.(*finishedMsg)
									if !ok {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:534
		_go_fuzz_dep_.CoverTab[22937]++
										c.sendAlert(alertUnexpectedMessage)
										return unexpectedMessageError(finished, msg)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:536
		// _ = "end of CoverTab[22937]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:537
		_go_fuzz_dep_.CoverTab[22938]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:537
		// _ = "end of CoverTab[22938]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:537
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:537
	// _ = "end of CoverTab[22929]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:537
	_go_fuzz_dep_.CoverTab[22930]++

									expectedMAC := hs.suite.finishedHash(c.in.trafficSecret, hs.transcript)
									if !hmac.Equal(expectedMAC, finished.verifyData) {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:540
		_go_fuzz_dep_.CoverTab[22939]++
										c.sendAlert(alertDecryptError)
										return errors.New("tls: invalid server finished hash")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:542
		// _ = "end of CoverTab[22939]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:543
		_go_fuzz_dep_.CoverTab[22940]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:543
		// _ = "end of CoverTab[22940]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:543
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:543
	// _ = "end of CoverTab[22930]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:543
	_go_fuzz_dep_.CoverTab[22931]++

									if err := transcriptMsg(finished, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:545
		_go_fuzz_dep_.CoverTab[22941]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:546
		// _ = "end of CoverTab[22941]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:547
		_go_fuzz_dep_.CoverTab[22942]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:547
		// _ = "end of CoverTab[22942]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:547
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:547
	// _ = "end of CoverTab[22931]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:547
	_go_fuzz_dep_.CoverTab[22932]++

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:551
	hs.trafficSecret = hs.suite.deriveSecret(hs.masterSecret,
		clientApplicationTrafficLabel, hs.transcript)
	serverSecret := hs.suite.deriveSecret(hs.masterSecret,
		serverApplicationTrafficLabel, hs.transcript)
	c.in.setTrafficSecret(hs.suite, serverSecret)

	err = c.config.writeKeyLog(keyLogLabelClientTraffic, hs.hello.random, hs.trafficSecret)
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:558
		_go_fuzz_dep_.CoverTab[22943]++
										c.sendAlert(alertInternalError)
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:560
		// _ = "end of CoverTab[22943]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:561
		_go_fuzz_dep_.CoverTab[22944]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:561
		// _ = "end of CoverTab[22944]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:561
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:561
	// _ = "end of CoverTab[22932]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:561
	_go_fuzz_dep_.CoverTab[22933]++
									err = c.config.writeKeyLog(keyLogLabelServerTraffic, hs.hello.random, serverSecret)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:563
		_go_fuzz_dep_.CoverTab[22945]++
										c.sendAlert(alertInternalError)
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:565
		// _ = "end of CoverTab[22945]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:566
		_go_fuzz_dep_.CoverTab[22946]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:566
		// _ = "end of CoverTab[22946]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:566
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:566
	// _ = "end of CoverTab[22933]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:566
	_go_fuzz_dep_.CoverTab[22934]++

									c.ekm = hs.suite.exportKeyingMaterial(hs.masterSecret, hs.transcript)

									return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:570
	// _ = "end of CoverTab[22934]"
}

func (hs *clientHandshakeStateTLS13) sendClientCertificate() error {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:573
	_go_fuzz_dep_.CoverTab[22947]++
									c := hs.c

									if hs.certReq == nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:576
		_go_fuzz_dep_.CoverTab[22957]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:577
		// _ = "end of CoverTab[22957]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:578
		_go_fuzz_dep_.CoverTab[22958]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:578
		// _ = "end of CoverTab[22958]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:578
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:578
	// _ = "end of CoverTab[22947]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:578
	_go_fuzz_dep_.CoverTab[22948]++

									cert, err := c.getClientCertificate(&CertificateRequestInfo{
		AcceptableCAs:		hs.certReq.certificateAuthorities,
		SignatureSchemes:	hs.certReq.supportedSignatureAlgorithms,
		Version:		c.vers,
		ctx:			hs.ctx,
	})
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:586
		_go_fuzz_dep_.CoverTab[22959]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:587
		// _ = "end of CoverTab[22959]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:588
		_go_fuzz_dep_.CoverTab[22960]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:588
		// _ = "end of CoverTab[22960]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:588
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:588
	// _ = "end of CoverTab[22948]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:588
	_go_fuzz_dep_.CoverTab[22949]++

									certMsg := new(certificateMsgTLS13)

									certMsg.certificate = *cert
									certMsg.scts = hs.certReq.scts && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:593
		_go_fuzz_dep_.CoverTab[22961]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:593
		return len(cert.SignedCertificateTimestamps) > 0
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:593
		// _ = "end of CoverTab[22961]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:593
	}()
									certMsg.ocspStapling = hs.certReq.ocspStapling && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:594
		_go_fuzz_dep_.CoverTab[22962]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:594
		return len(cert.OCSPStaple) > 0
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:594
		// _ = "end of CoverTab[22962]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:594
	}()

									if _, err := hs.c.writeHandshakeRecord(certMsg, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:596
		_go_fuzz_dep_.CoverTab[22963]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:597
		// _ = "end of CoverTab[22963]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:598
		_go_fuzz_dep_.CoverTab[22964]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:598
		// _ = "end of CoverTab[22964]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:598
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:598
	// _ = "end of CoverTab[22949]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:598
	_go_fuzz_dep_.CoverTab[22950]++

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:601
	if len(cert.Certificate) == 0 {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:601
		_go_fuzz_dep_.CoverTab[22965]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:602
		// _ = "end of CoverTab[22965]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:603
		_go_fuzz_dep_.CoverTab[22966]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:603
		// _ = "end of CoverTab[22966]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:603
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:603
	// _ = "end of CoverTab[22950]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:603
	_go_fuzz_dep_.CoverTab[22951]++

									certVerifyMsg := new(certificateVerifyMsg)
									certVerifyMsg.hasSignatureAlgorithm = true

									certVerifyMsg.signatureAlgorithm, err = selectSignatureScheme(c.vers, cert, hs.certReq.supportedSignatureAlgorithms)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:609
		_go_fuzz_dep_.CoverTab[22967]++

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:612
		c.sendAlert(alertHandshakeFailure)
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:613
		// _ = "end of CoverTab[22967]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:614
		_go_fuzz_dep_.CoverTab[22968]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:614
		// _ = "end of CoverTab[22968]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:614
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:614
	// _ = "end of CoverTab[22951]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:614
	_go_fuzz_dep_.CoverTab[22952]++

									sigType, sigHash, err := typeAndHashFromSignatureScheme(certVerifyMsg.signatureAlgorithm)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:617
		_go_fuzz_dep_.CoverTab[22969]++
										return c.sendAlert(alertInternalError)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:618
		// _ = "end of CoverTab[22969]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:619
		_go_fuzz_dep_.CoverTab[22970]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:619
		// _ = "end of CoverTab[22970]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:619
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:619
	// _ = "end of CoverTab[22952]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:619
	_go_fuzz_dep_.CoverTab[22953]++

									signed := signedMessage(sigHash, clientSignatureContext, hs.transcript)
									signOpts := crypto.SignerOpts(sigHash)
									if sigType == signatureRSAPSS {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:623
		_go_fuzz_dep_.CoverTab[22971]++
										signOpts = &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash, Hash: sigHash}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:624
		// _ = "end of CoverTab[22971]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:625
		_go_fuzz_dep_.CoverTab[22972]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:625
		// _ = "end of CoverTab[22972]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:625
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:625
	// _ = "end of CoverTab[22953]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:625
	_go_fuzz_dep_.CoverTab[22954]++
									sig, err := cert.PrivateKey.(crypto.Signer).Sign(c.config.rand(), signed, signOpts)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:627
		_go_fuzz_dep_.CoverTab[22973]++
										c.sendAlert(alertInternalError)
										return errors.New("tls: failed to sign handshake: " + err.Error())
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:629
		// _ = "end of CoverTab[22973]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:630
		_go_fuzz_dep_.CoverTab[22974]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:630
		// _ = "end of CoverTab[22974]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:630
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:630
	// _ = "end of CoverTab[22954]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:630
	_go_fuzz_dep_.CoverTab[22955]++
									certVerifyMsg.signature = sig

									if _, err := hs.c.writeHandshakeRecord(certVerifyMsg, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:633
		_go_fuzz_dep_.CoverTab[22975]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:634
		// _ = "end of CoverTab[22975]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:635
		_go_fuzz_dep_.CoverTab[22976]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:635
		// _ = "end of CoverTab[22976]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:635
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:635
	// _ = "end of CoverTab[22955]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:635
	_go_fuzz_dep_.CoverTab[22956]++

									return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:637
	// _ = "end of CoverTab[22956]"
}

func (hs *clientHandshakeStateTLS13) sendClientFinished() error {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:640
	_go_fuzz_dep_.CoverTab[22977]++
									c := hs.c

									finished := &finishedMsg{
		verifyData: hs.suite.finishedHash(c.out.trafficSecret, hs.transcript),
	}

	if _, err := hs.c.writeHandshakeRecord(finished, hs.transcript); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:647
		_go_fuzz_dep_.CoverTab[22980]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:648
		// _ = "end of CoverTab[22980]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:649
		_go_fuzz_dep_.CoverTab[22981]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:649
		// _ = "end of CoverTab[22981]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:649
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:649
	// _ = "end of CoverTab[22977]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:649
	_go_fuzz_dep_.CoverTab[22978]++

									c.out.setTrafficSecret(hs.suite, hs.trafficSecret)

									if !c.config.SessionTicketsDisabled && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:653
		_go_fuzz_dep_.CoverTab[22982]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:653
		return c.config.ClientSessionCache != nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:653
		// _ = "end of CoverTab[22982]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:653
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:653
		_go_fuzz_dep_.CoverTab[22983]++
										c.resumptionSecret = hs.suite.deriveSecret(hs.masterSecret,
			resumptionLabel, hs.transcript)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:655
		// _ = "end of CoverTab[22983]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:656
		_go_fuzz_dep_.CoverTab[22984]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:656
		// _ = "end of CoverTab[22984]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:656
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:656
	// _ = "end of CoverTab[22978]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:656
	_go_fuzz_dep_.CoverTab[22979]++

									return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:658
	// _ = "end of CoverTab[22979]"
}

func (c *Conn) handleNewSessionTicket(msg *newSessionTicketMsgTLS13) error {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:661
	_go_fuzz_dep_.CoverTab[22985]++
									if !c.isClient {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:662
		_go_fuzz_dep_.CoverTab[22991]++
										c.sendAlert(alertUnexpectedMessage)
										return errors.New("tls: received new session ticket from a client")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:664
		// _ = "end of CoverTab[22991]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:665
		_go_fuzz_dep_.CoverTab[22992]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:665
		// _ = "end of CoverTab[22992]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:665
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:665
	// _ = "end of CoverTab[22985]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:665
	_go_fuzz_dep_.CoverTab[22986]++

									if c.config.SessionTicketsDisabled || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:667
		_go_fuzz_dep_.CoverTab[22993]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:667
		return c.config.ClientSessionCache == nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:667
		// _ = "end of CoverTab[22993]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:667
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:667
		_go_fuzz_dep_.CoverTab[22994]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:668
		// _ = "end of CoverTab[22994]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:669
		_go_fuzz_dep_.CoverTab[22995]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:669
		// _ = "end of CoverTab[22995]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:669
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:669
	// _ = "end of CoverTab[22986]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:669
	_go_fuzz_dep_.CoverTab[22987]++

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:672
	if msg.lifetime == 0 {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:672
		_go_fuzz_dep_.CoverTab[22996]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:673
		// _ = "end of CoverTab[22996]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:674
		_go_fuzz_dep_.CoverTab[22997]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:674
		// _ = "end of CoverTab[22997]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:674
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:674
	// _ = "end of CoverTab[22987]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:674
	_go_fuzz_dep_.CoverTab[22988]++
									lifetime := time.Duration(msg.lifetime) * time.Second
									if lifetime > maxSessionTicketLifetime {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:676
		_go_fuzz_dep_.CoverTab[22998]++
										c.sendAlert(alertIllegalParameter)
										return errors.New("tls: received a session ticket with invalid lifetime")
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:678
		// _ = "end of CoverTab[22998]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:679
		_go_fuzz_dep_.CoverTab[22999]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:679
		// _ = "end of CoverTab[22999]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:679
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:679
	// _ = "end of CoverTab[22988]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:679
	_go_fuzz_dep_.CoverTab[22989]++

									cipherSuite := cipherSuiteTLS13ByID(c.cipherSuite)
									if cipherSuite == nil || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:682
		_go_fuzz_dep_.CoverTab[23000]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:682
		return c.resumptionSecret == nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:682
		// _ = "end of CoverTab[23000]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:682
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:682
		_go_fuzz_dep_.CoverTab[23001]++
										return c.sendAlert(alertInternalError)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:683
		// _ = "end of CoverTab[23001]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:684
		_go_fuzz_dep_.CoverTab[23002]++
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:684
		// _ = "end of CoverTab[23002]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:684
	}
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:684
	// _ = "end of CoverTab[22989]"
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:684
	_go_fuzz_dep_.CoverTab[22990]++

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:690
	session := &ClientSessionState{
		sessionTicket:		msg.label,
		vers:			c.vers,
		cipherSuite:		c.cipherSuite,
		masterSecret:		c.resumptionSecret,
		serverCertificates:	c.peerCertificates,
		verifiedChains:		c.verifiedChains,
		receivedAt:		c.config.time(),
		nonce:			msg.nonce,
		useBy:			c.config.time().Add(lifetime),
		ageAdd:			msg.ageAdd,
		ocspResponse:		c.ocspResponse,
		scts:			c.scts,
	}

									cacheKey := clientSessionCacheKey(c.conn.RemoteAddr(), c.config)
									c.config.ClientSessionCache.Put(cacheKey, session)

									return nil
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:708
	// _ = "end of CoverTab[22990]"
}

//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:709
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/handshake_client_tls13.go:709
var _ = _go_fuzz_dep_.CoverTab
