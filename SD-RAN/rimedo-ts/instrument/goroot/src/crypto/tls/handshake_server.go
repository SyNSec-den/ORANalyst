// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/tls/handshake_server.go:5
package tls

//line /usr/local/go/src/crypto/tls/handshake_server.go:5
import (
//line /usr/local/go/src/crypto/tls/handshake_server.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/handshake_server.go:5
)
//line /usr/local/go/src/crypto/tls/handshake_server.go:5
import (
//line /usr/local/go/src/crypto/tls/handshake_server.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/handshake_server.go:5
)

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/subtle"
	"crypto/x509"
	"errors"
	"fmt"
	"hash"
	"io"
	"time"
)

// serverHandshakeState contains details of a server handshake in progress.
//line /usr/local/go/src/crypto/tls/handshake_server.go:22
// It's discarded once the handshake has completed.
//line /usr/local/go/src/crypto/tls/handshake_server.go:24
type serverHandshakeState struct {
	c		*Conn
	ctx		context.Context
	clientHello	*clientHelloMsg
	hello		*serverHelloMsg
	suite		*cipherSuite
	ecdheOk		bool
	ecSignOk	bool
	rsaDecryptOk	bool
	rsaSignOk	bool
	sessionState	*sessionState
	finishedHash	finishedHash
	masterSecret	[]byte
	cert		*Certificate
}

// serverHandshake performs a TLS handshake as a server.
func (c *Conn) serverHandshake(ctx context.Context) error {
//line /usr/local/go/src/crypto/tls/handshake_server.go:41
	_go_fuzz_dep_.CoverTab[23879]++
								clientHello, err := c.readClientHello(ctx)
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:43
		_go_fuzz_dep_.CoverTab[23882]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:44
		// _ = "end of CoverTab[23882]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:45
		_go_fuzz_dep_.CoverTab[23883]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:45
		// _ = "end of CoverTab[23883]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:45
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:45
	// _ = "end of CoverTab[23879]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:45
	_go_fuzz_dep_.CoverTab[23880]++

								if c.vers == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/handshake_server.go:47
		_go_fuzz_dep_.CoverTab[23884]++
									hs := serverHandshakeStateTLS13{
			c:		c,
			ctx:		ctx,
			clientHello:	clientHello,
		}
									return hs.handshake()
//line /usr/local/go/src/crypto/tls/handshake_server.go:53
		// _ = "end of CoverTab[23884]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:54
		_go_fuzz_dep_.CoverTab[23885]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:54
		// _ = "end of CoverTab[23885]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:54
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:54
	// _ = "end of CoverTab[23880]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:54
	_go_fuzz_dep_.CoverTab[23881]++

								hs := serverHandshakeState{
		c:		c,
		ctx:		ctx,
		clientHello:	clientHello,
	}
								return hs.handshake()
//line /usr/local/go/src/crypto/tls/handshake_server.go:61
	// _ = "end of CoverTab[23881]"
}

func (hs *serverHandshakeState) handshake() error {
//line /usr/local/go/src/crypto/tls/handshake_server.go:64
	_go_fuzz_dep_.CoverTab[23886]++
								c := hs.c

								if err := hs.processClientHello(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:67
		_go_fuzz_dep_.CoverTab[23889]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:68
		// _ = "end of CoverTab[23889]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:69
		_go_fuzz_dep_.CoverTab[23890]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:69
		// _ = "end of CoverTab[23890]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:69
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:69
	// _ = "end of CoverTab[23886]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:69
	_go_fuzz_dep_.CoverTab[23887]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:72
	c.buffering = true
	if hs.checkForResumption() {
//line /usr/local/go/src/crypto/tls/handshake_server.go:73
		_go_fuzz_dep_.CoverTab[23891]++

									c.didResume = true
									if err := hs.doResumeHandshake(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:76
			_go_fuzz_dep_.CoverTab[23897]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:77
			// _ = "end of CoverTab[23897]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:78
			_go_fuzz_dep_.CoverTab[23898]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:78
			// _ = "end of CoverTab[23898]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:78
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:78
		// _ = "end of CoverTab[23891]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:78
		_go_fuzz_dep_.CoverTab[23892]++
									if err := hs.establishKeys(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:79
			_go_fuzz_dep_.CoverTab[23899]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:80
			// _ = "end of CoverTab[23899]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:81
			_go_fuzz_dep_.CoverTab[23900]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:81
			// _ = "end of CoverTab[23900]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:81
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:81
		// _ = "end of CoverTab[23892]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:81
		_go_fuzz_dep_.CoverTab[23893]++
									if err := hs.sendSessionTicket(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:82
			_go_fuzz_dep_.CoverTab[23901]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:83
			// _ = "end of CoverTab[23901]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:84
			_go_fuzz_dep_.CoverTab[23902]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:84
			// _ = "end of CoverTab[23902]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:84
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:84
		// _ = "end of CoverTab[23893]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:84
		_go_fuzz_dep_.CoverTab[23894]++
									if err := hs.sendFinished(c.serverFinished[:]); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:85
			_go_fuzz_dep_.CoverTab[23903]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:86
			// _ = "end of CoverTab[23903]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:87
			_go_fuzz_dep_.CoverTab[23904]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:87
			// _ = "end of CoverTab[23904]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:87
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:87
		// _ = "end of CoverTab[23894]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:87
		_go_fuzz_dep_.CoverTab[23895]++
									if _, err := c.flush(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:88
			_go_fuzz_dep_.CoverTab[23905]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:89
			// _ = "end of CoverTab[23905]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:90
			_go_fuzz_dep_.CoverTab[23906]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:90
			// _ = "end of CoverTab[23906]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:90
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:90
		// _ = "end of CoverTab[23895]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:90
		_go_fuzz_dep_.CoverTab[23896]++
									c.clientFinishedIsFirst = false
									if err := hs.readFinished(nil); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:92
			_go_fuzz_dep_.CoverTab[23907]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:93
			// _ = "end of CoverTab[23907]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:94
			_go_fuzz_dep_.CoverTab[23908]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:94
			// _ = "end of CoverTab[23908]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:94
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:94
		// _ = "end of CoverTab[23896]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:95
		_go_fuzz_dep_.CoverTab[23909]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:98
		if err := hs.pickCipherSuite(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:98
			_go_fuzz_dep_.CoverTab[23916]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:99
			// _ = "end of CoverTab[23916]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:100
			_go_fuzz_dep_.CoverTab[23917]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:100
			// _ = "end of CoverTab[23917]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:100
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:100
		// _ = "end of CoverTab[23909]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:100
		_go_fuzz_dep_.CoverTab[23910]++
									if err := hs.doFullHandshake(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:101
			_go_fuzz_dep_.CoverTab[23918]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:102
			// _ = "end of CoverTab[23918]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:103
			_go_fuzz_dep_.CoverTab[23919]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:103
			// _ = "end of CoverTab[23919]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:103
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:103
		// _ = "end of CoverTab[23910]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:103
		_go_fuzz_dep_.CoverTab[23911]++
									if err := hs.establishKeys(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:104
			_go_fuzz_dep_.CoverTab[23920]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:105
			// _ = "end of CoverTab[23920]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:106
			_go_fuzz_dep_.CoverTab[23921]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:106
			// _ = "end of CoverTab[23921]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:106
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:106
		// _ = "end of CoverTab[23911]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:106
		_go_fuzz_dep_.CoverTab[23912]++
									if err := hs.readFinished(c.clientFinished[:]); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:107
			_go_fuzz_dep_.CoverTab[23922]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:108
			// _ = "end of CoverTab[23922]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:109
			_go_fuzz_dep_.CoverTab[23923]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:109
			// _ = "end of CoverTab[23923]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:109
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:109
		// _ = "end of CoverTab[23912]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:109
		_go_fuzz_dep_.CoverTab[23913]++
									c.clientFinishedIsFirst = true
									c.buffering = true
									if err := hs.sendSessionTicket(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:112
			_go_fuzz_dep_.CoverTab[23924]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:113
			// _ = "end of CoverTab[23924]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:114
			_go_fuzz_dep_.CoverTab[23925]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:114
			// _ = "end of CoverTab[23925]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:114
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:114
		// _ = "end of CoverTab[23913]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:114
		_go_fuzz_dep_.CoverTab[23914]++
									if err := hs.sendFinished(nil); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:115
			_go_fuzz_dep_.CoverTab[23926]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:116
			// _ = "end of CoverTab[23926]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:117
			_go_fuzz_dep_.CoverTab[23927]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:117
			// _ = "end of CoverTab[23927]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:117
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:117
		// _ = "end of CoverTab[23914]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:117
		_go_fuzz_dep_.CoverTab[23915]++
									if _, err := c.flush(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:118
			_go_fuzz_dep_.CoverTab[23928]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:119
			// _ = "end of CoverTab[23928]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:120
			_go_fuzz_dep_.CoverTab[23929]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:120
			// _ = "end of CoverTab[23929]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:120
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:120
		// _ = "end of CoverTab[23915]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:121
	// _ = "end of CoverTab[23887]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:121
	_go_fuzz_dep_.CoverTab[23888]++

								c.ekm = ekmFromMasterSecret(c.vers, hs.suite, hs.masterSecret, hs.clientHello.random, hs.hello.random)
								c.isHandshakeComplete.Store(true)

								return nil
//line /usr/local/go/src/crypto/tls/handshake_server.go:126
	// _ = "end of CoverTab[23888]"
}

// readClientHello reads a ClientHello message and selects the protocol version.
func (c *Conn) readClientHello(ctx context.Context) (*clientHelloMsg, error) {
//line /usr/local/go/src/crypto/tls/handshake_server.go:130
	_go_fuzz_dep_.CoverTab[23930]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:133
	msg, err := c.readHandshake(nil)
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:134
		_go_fuzz_dep_.CoverTab[23936]++
									return nil, err
//line /usr/local/go/src/crypto/tls/handshake_server.go:135
		// _ = "end of CoverTab[23936]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:136
		_go_fuzz_dep_.CoverTab[23937]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:136
		// _ = "end of CoverTab[23937]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:136
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:136
	// _ = "end of CoverTab[23930]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:136
	_go_fuzz_dep_.CoverTab[23931]++
								clientHello, ok := msg.(*clientHelloMsg)
								if !ok {
//line /usr/local/go/src/crypto/tls/handshake_server.go:138
		_go_fuzz_dep_.CoverTab[23938]++
									c.sendAlert(alertUnexpectedMessage)
									return nil, unexpectedMessageError(clientHello, msg)
//line /usr/local/go/src/crypto/tls/handshake_server.go:140
		// _ = "end of CoverTab[23938]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:141
		_go_fuzz_dep_.CoverTab[23939]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:141
		// _ = "end of CoverTab[23939]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:141
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:141
	// _ = "end of CoverTab[23931]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:141
	_go_fuzz_dep_.CoverTab[23932]++

								var configForClient *Config
								originalConfig := c.config
								if c.config.GetConfigForClient != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:145
		_go_fuzz_dep_.CoverTab[23940]++
									chi := clientHelloInfo(ctx, c, clientHello)
									if configForClient, err = c.config.GetConfigForClient(chi); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:147
			_go_fuzz_dep_.CoverTab[23941]++
										c.sendAlert(alertInternalError)
										return nil, err
//line /usr/local/go/src/crypto/tls/handshake_server.go:149
			// _ = "end of CoverTab[23941]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:150
			_go_fuzz_dep_.CoverTab[23942]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:150
			if configForClient != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:150
				_go_fuzz_dep_.CoverTab[23943]++
											c.config = configForClient
//line /usr/local/go/src/crypto/tls/handshake_server.go:151
				// _ = "end of CoverTab[23943]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:152
				_go_fuzz_dep_.CoverTab[23944]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:152
				// _ = "end of CoverTab[23944]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:152
			}
//line /usr/local/go/src/crypto/tls/handshake_server.go:152
			// _ = "end of CoverTab[23942]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:152
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:152
		// _ = "end of CoverTab[23940]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:153
		_go_fuzz_dep_.CoverTab[23945]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:153
		// _ = "end of CoverTab[23945]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:153
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:153
	// _ = "end of CoverTab[23932]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:153
	_go_fuzz_dep_.CoverTab[23933]++
								c.ticketKeys = originalConfig.ticketKeys(configForClient)

								clientVersions := clientHello.supportedVersions
								if len(clientHello.supportedVersions) == 0 {
//line /usr/local/go/src/crypto/tls/handshake_server.go:157
		_go_fuzz_dep_.CoverTab[23946]++
									clientVersions = supportedVersionsFromMax(clientHello.vers)
//line /usr/local/go/src/crypto/tls/handshake_server.go:158
		// _ = "end of CoverTab[23946]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:159
		_go_fuzz_dep_.CoverTab[23947]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:159
		// _ = "end of CoverTab[23947]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:159
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:159
	// _ = "end of CoverTab[23933]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:159
	_go_fuzz_dep_.CoverTab[23934]++
								c.vers, ok = c.config.mutualVersion(roleServer, clientVersions)
								if !ok {
//line /usr/local/go/src/crypto/tls/handshake_server.go:161
		_go_fuzz_dep_.CoverTab[23948]++
									c.sendAlert(alertProtocolVersion)
									return nil, fmt.Errorf("tls: client offered only unsupported versions: %x", clientVersions)
//line /usr/local/go/src/crypto/tls/handshake_server.go:163
		// _ = "end of CoverTab[23948]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:164
		_go_fuzz_dep_.CoverTab[23949]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:164
		// _ = "end of CoverTab[23949]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:164
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:164
	// _ = "end of CoverTab[23934]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:164
	_go_fuzz_dep_.CoverTab[23935]++
								c.haveVers = true
								c.in.version = c.vers
								c.out.version = c.vers

								return clientHello, nil
//line /usr/local/go/src/crypto/tls/handshake_server.go:169
	// _ = "end of CoverTab[23935]"
}

func (hs *serverHandshakeState) processClientHello() error {
//line /usr/local/go/src/crypto/tls/handshake_server.go:172
	_go_fuzz_dep_.CoverTab[23950]++
								c := hs.c

								hs.hello = new(serverHelloMsg)
								hs.hello.vers = c.vers

								foundCompression := false

								for _, compression := range hs.clientHello.compressionMethods {
//line /usr/local/go/src/crypto/tls/handshake_server.go:180
		_go_fuzz_dep_.CoverTab[23963]++
									if compression == compressionNone {
//line /usr/local/go/src/crypto/tls/handshake_server.go:181
			_go_fuzz_dep_.CoverTab[23964]++
										foundCompression = true
										break
//line /usr/local/go/src/crypto/tls/handshake_server.go:183
			// _ = "end of CoverTab[23964]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:184
			_go_fuzz_dep_.CoverTab[23965]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:184
			// _ = "end of CoverTab[23965]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:184
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:184
		// _ = "end of CoverTab[23963]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:185
	// _ = "end of CoverTab[23950]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:185
	_go_fuzz_dep_.CoverTab[23951]++

								if !foundCompression {
//line /usr/local/go/src/crypto/tls/handshake_server.go:187
		_go_fuzz_dep_.CoverTab[23966]++
									c.sendAlert(alertHandshakeFailure)
									return errors.New("tls: client does not support uncompressed connections")
//line /usr/local/go/src/crypto/tls/handshake_server.go:189
		// _ = "end of CoverTab[23966]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:190
		_go_fuzz_dep_.CoverTab[23967]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:190
		// _ = "end of CoverTab[23967]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:190
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:190
	// _ = "end of CoverTab[23951]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:190
	_go_fuzz_dep_.CoverTab[23952]++

								hs.hello.random = make([]byte, 32)
								serverRandom := hs.hello.random

								maxVers := c.config.maxSupportedVersion(roleServer)
								if maxVers >= VersionTLS12 && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:196
		_go_fuzz_dep_.CoverTab[23968]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:196
		return c.vers < maxVers
//line /usr/local/go/src/crypto/tls/handshake_server.go:196
		// _ = "end of CoverTab[23968]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:196
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:196
		_go_fuzz_dep_.CoverTab[23969]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:196
		return testingOnlyForceDowngradeCanary
//line /usr/local/go/src/crypto/tls/handshake_server.go:196
		// _ = "end of CoverTab[23969]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:196
	}() {
//line /usr/local/go/src/crypto/tls/handshake_server.go:196
		_go_fuzz_dep_.CoverTab[23970]++
									if c.vers == VersionTLS12 {
//line /usr/local/go/src/crypto/tls/handshake_server.go:197
			_go_fuzz_dep_.CoverTab[23972]++
										copy(serverRandom[24:], downgradeCanaryTLS12)
//line /usr/local/go/src/crypto/tls/handshake_server.go:198
			// _ = "end of CoverTab[23972]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:199
			_go_fuzz_dep_.CoverTab[23973]++
										copy(serverRandom[24:], downgradeCanaryTLS11)
//line /usr/local/go/src/crypto/tls/handshake_server.go:200
			// _ = "end of CoverTab[23973]"
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:201
		// _ = "end of CoverTab[23970]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:201
		_go_fuzz_dep_.CoverTab[23971]++
									serverRandom = serverRandom[:24]
//line /usr/local/go/src/crypto/tls/handshake_server.go:202
		// _ = "end of CoverTab[23971]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:203
		_go_fuzz_dep_.CoverTab[23974]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:203
		// _ = "end of CoverTab[23974]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:203
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:203
	// _ = "end of CoverTab[23952]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:203
	_go_fuzz_dep_.CoverTab[23953]++
								_, err := io.ReadFull(c.config.rand(), serverRandom)
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:205
		_go_fuzz_dep_.CoverTab[23975]++
									c.sendAlert(alertInternalError)
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:207
		// _ = "end of CoverTab[23975]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:208
		_go_fuzz_dep_.CoverTab[23976]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:208
		// _ = "end of CoverTab[23976]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:208
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:208
	// _ = "end of CoverTab[23953]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:208
	_go_fuzz_dep_.CoverTab[23954]++

								if len(hs.clientHello.secureRenegotiation) != 0 {
//line /usr/local/go/src/crypto/tls/handshake_server.go:210
		_go_fuzz_dep_.CoverTab[23977]++
									c.sendAlert(alertHandshakeFailure)
									return errors.New("tls: initial handshake had non-empty renegotiation extension")
//line /usr/local/go/src/crypto/tls/handshake_server.go:212
		// _ = "end of CoverTab[23977]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:213
		_go_fuzz_dep_.CoverTab[23978]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:213
		// _ = "end of CoverTab[23978]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:213
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:213
	// _ = "end of CoverTab[23954]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:213
	_go_fuzz_dep_.CoverTab[23955]++

								hs.hello.secureRenegotiationSupported = hs.clientHello.secureRenegotiationSupported
								hs.hello.compressionMethod = compressionNone
								if len(hs.clientHello.serverName) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_server.go:217
		_go_fuzz_dep_.CoverTab[23979]++
									c.serverName = hs.clientHello.serverName
//line /usr/local/go/src/crypto/tls/handshake_server.go:218
		// _ = "end of CoverTab[23979]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:219
		_go_fuzz_dep_.CoverTab[23980]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:219
		// _ = "end of CoverTab[23980]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:219
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:219
	// _ = "end of CoverTab[23955]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:219
	_go_fuzz_dep_.CoverTab[23956]++

								selectedProto, err := negotiateALPN(c.config.NextProtos, hs.clientHello.alpnProtocols)
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:222
		_go_fuzz_dep_.CoverTab[23981]++
									c.sendAlert(alertNoApplicationProtocol)
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:224
		// _ = "end of CoverTab[23981]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:225
		_go_fuzz_dep_.CoverTab[23982]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:225
		// _ = "end of CoverTab[23982]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:225
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:225
	// _ = "end of CoverTab[23956]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:225
	_go_fuzz_dep_.CoverTab[23957]++
								hs.hello.alpnProtocol = selectedProto
								c.clientProtocol = selectedProto

								hs.cert, err = c.config.getCertificate(clientHelloInfo(hs.ctx, c, hs.clientHello))
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:230
		_go_fuzz_dep_.CoverTab[23983]++
									if err == errNoCertificates {
//line /usr/local/go/src/crypto/tls/handshake_server.go:231
			_go_fuzz_dep_.CoverTab[23985]++
										c.sendAlert(alertUnrecognizedName)
//line /usr/local/go/src/crypto/tls/handshake_server.go:232
			// _ = "end of CoverTab[23985]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:233
			_go_fuzz_dep_.CoverTab[23986]++
										c.sendAlert(alertInternalError)
//line /usr/local/go/src/crypto/tls/handshake_server.go:234
			// _ = "end of CoverTab[23986]"
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:235
		// _ = "end of CoverTab[23983]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:235
		_go_fuzz_dep_.CoverTab[23984]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:236
		// _ = "end of CoverTab[23984]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:237
		_go_fuzz_dep_.CoverTab[23987]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:237
		// _ = "end of CoverTab[23987]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:237
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:237
	// _ = "end of CoverTab[23957]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:237
	_go_fuzz_dep_.CoverTab[23958]++
								if hs.clientHello.scts {
//line /usr/local/go/src/crypto/tls/handshake_server.go:238
		_go_fuzz_dep_.CoverTab[23988]++
									hs.hello.scts = hs.cert.SignedCertificateTimestamps
//line /usr/local/go/src/crypto/tls/handshake_server.go:239
		// _ = "end of CoverTab[23988]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:240
		_go_fuzz_dep_.CoverTab[23989]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:240
		// _ = "end of CoverTab[23989]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:240
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:240
	// _ = "end of CoverTab[23958]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:240
	_go_fuzz_dep_.CoverTab[23959]++

								hs.ecdheOk = supportsECDHE(c.config, hs.clientHello.supportedCurves, hs.clientHello.supportedPoints)

								if hs.ecdheOk && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:244
		_go_fuzz_dep_.CoverTab[23990]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:244
		return len(hs.clientHello.supportedPoints) > 0
//line /usr/local/go/src/crypto/tls/handshake_server.go:244
		// _ = "end of CoverTab[23990]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:244
	}() {
//line /usr/local/go/src/crypto/tls/handshake_server.go:244
		_go_fuzz_dep_.CoverTab[23991]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:250
		hs.hello.supportedPoints = []uint8{pointFormatUncompressed}
//line /usr/local/go/src/crypto/tls/handshake_server.go:250
		// _ = "end of CoverTab[23991]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:251
		_go_fuzz_dep_.CoverTab[23992]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:251
		// _ = "end of CoverTab[23992]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:251
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:251
	// _ = "end of CoverTab[23959]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:251
	_go_fuzz_dep_.CoverTab[23960]++

								if priv, ok := hs.cert.PrivateKey.(crypto.Signer); ok {
//line /usr/local/go/src/crypto/tls/handshake_server.go:253
		_go_fuzz_dep_.CoverTab[23993]++
									switch priv.Public().(type) {
		case *ecdsa.PublicKey:
//line /usr/local/go/src/crypto/tls/handshake_server.go:255
			_go_fuzz_dep_.CoverTab[23994]++
										hs.ecSignOk = true
//line /usr/local/go/src/crypto/tls/handshake_server.go:256
			// _ = "end of CoverTab[23994]"
		case ed25519.PublicKey:
//line /usr/local/go/src/crypto/tls/handshake_server.go:257
			_go_fuzz_dep_.CoverTab[23995]++
										hs.ecSignOk = true
//line /usr/local/go/src/crypto/tls/handshake_server.go:258
			// _ = "end of CoverTab[23995]"
		case *rsa.PublicKey:
//line /usr/local/go/src/crypto/tls/handshake_server.go:259
			_go_fuzz_dep_.CoverTab[23996]++
										hs.rsaSignOk = true
//line /usr/local/go/src/crypto/tls/handshake_server.go:260
			// _ = "end of CoverTab[23996]"
		default:
//line /usr/local/go/src/crypto/tls/handshake_server.go:261
			_go_fuzz_dep_.CoverTab[23997]++
										c.sendAlert(alertInternalError)
										return fmt.Errorf("tls: unsupported signing key type (%T)", priv.Public())
//line /usr/local/go/src/crypto/tls/handshake_server.go:263
			// _ = "end of CoverTab[23997]"
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:264
		// _ = "end of CoverTab[23993]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:265
		_go_fuzz_dep_.CoverTab[23998]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:265
		// _ = "end of CoverTab[23998]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:265
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:265
	// _ = "end of CoverTab[23960]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:265
	_go_fuzz_dep_.CoverTab[23961]++
								if priv, ok := hs.cert.PrivateKey.(crypto.Decrypter); ok {
//line /usr/local/go/src/crypto/tls/handshake_server.go:266
		_go_fuzz_dep_.CoverTab[23999]++
									switch priv.Public().(type) {
		case *rsa.PublicKey:
//line /usr/local/go/src/crypto/tls/handshake_server.go:268
			_go_fuzz_dep_.CoverTab[24000]++
										hs.rsaDecryptOk = true
//line /usr/local/go/src/crypto/tls/handshake_server.go:269
			// _ = "end of CoverTab[24000]"
		default:
//line /usr/local/go/src/crypto/tls/handshake_server.go:270
			_go_fuzz_dep_.CoverTab[24001]++
										c.sendAlert(alertInternalError)
										return fmt.Errorf("tls: unsupported decryption key type (%T)", priv.Public())
//line /usr/local/go/src/crypto/tls/handshake_server.go:272
			// _ = "end of CoverTab[24001]"
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:273
		// _ = "end of CoverTab[23999]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:274
		_go_fuzz_dep_.CoverTab[24002]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:274
		// _ = "end of CoverTab[24002]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:274
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:274
	// _ = "end of CoverTab[23961]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:274
	_go_fuzz_dep_.CoverTab[23962]++

								return nil
//line /usr/local/go/src/crypto/tls/handshake_server.go:276
	// _ = "end of CoverTab[23962]"
}

// negotiateALPN picks a shared ALPN protocol that both sides support in server
//line /usr/local/go/src/crypto/tls/handshake_server.go:279
// preference order. If ALPN is not configured or the peer doesn't support it,
//line /usr/local/go/src/crypto/tls/handshake_server.go:279
// it returns "" and no error.
//line /usr/local/go/src/crypto/tls/handshake_server.go:282
func negotiateALPN(serverProtos, clientProtos []string) (string, error) {
//line /usr/local/go/src/crypto/tls/handshake_server.go:282
	_go_fuzz_dep_.CoverTab[24003]++
								if len(serverProtos) == 0 || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:283
		_go_fuzz_dep_.CoverTab[24007]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:283
		return len(clientProtos) == 0
//line /usr/local/go/src/crypto/tls/handshake_server.go:283
		// _ = "end of CoverTab[24007]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:283
	}() {
//line /usr/local/go/src/crypto/tls/handshake_server.go:283
		_go_fuzz_dep_.CoverTab[24008]++
									return "", nil
//line /usr/local/go/src/crypto/tls/handshake_server.go:284
		// _ = "end of CoverTab[24008]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:285
		_go_fuzz_dep_.CoverTab[24009]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:285
		// _ = "end of CoverTab[24009]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:285
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:285
	// _ = "end of CoverTab[24003]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:285
	_go_fuzz_dep_.CoverTab[24004]++
								var http11fallback bool
								for _, s := range serverProtos {
//line /usr/local/go/src/crypto/tls/handshake_server.go:287
		_go_fuzz_dep_.CoverTab[24010]++
									for _, c := range clientProtos {
//line /usr/local/go/src/crypto/tls/handshake_server.go:288
			_go_fuzz_dep_.CoverTab[24011]++
										if s == c {
//line /usr/local/go/src/crypto/tls/handshake_server.go:289
				_go_fuzz_dep_.CoverTab[24013]++
											return s, nil
//line /usr/local/go/src/crypto/tls/handshake_server.go:290
				// _ = "end of CoverTab[24013]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:291
				_go_fuzz_dep_.CoverTab[24014]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:291
				// _ = "end of CoverTab[24014]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:291
			}
//line /usr/local/go/src/crypto/tls/handshake_server.go:291
			// _ = "end of CoverTab[24011]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:291
			_go_fuzz_dep_.CoverTab[24012]++
										if s == "h2" && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:292
				_go_fuzz_dep_.CoverTab[24015]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:292
				return c == "http/1.1"
//line /usr/local/go/src/crypto/tls/handshake_server.go:292
				// _ = "end of CoverTab[24015]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:292
			}() {
//line /usr/local/go/src/crypto/tls/handshake_server.go:292
				_go_fuzz_dep_.CoverTab[24016]++
											http11fallback = true
//line /usr/local/go/src/crypto/tls/handshake_server.go:293
				// _ = "end of CoverTab[24016]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:294
				_go_fuzz_dep_.CoverTab[24017]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:294
				// _ = "end of CoverTab[24017]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:294
			}
//line /usr/local/go/src/crypto/tls/handshake_server.go:294
			// _ = "end of CoverTab[24012]"
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:295
		// _ = "end of CoverTab[24010]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:296
	// _ = "end of CoverTab[24004]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:296
	_go_fuzz_dep_.CoverTab[24005]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:301
	if http11fallback {
//line /usr/local/go/src/crypto/tls/handshake_server.go:301
		_go_fuzz_dep_.CoverTab[24018]++
									return "", nil
//line /usr/local/go/src/crypto/tls/handshake_server.go:302
		// _ = "end of CoverTab[24018]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:303
		_go_fuzz_dep_.CoverTab[24019]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:303
		// _ = "end of CoverTab[24019]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:303
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:303
	// _ = "end of CoverTab[24005]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:303
	_go_fuzz_dep_.CoverTab[24006]++
								return "", fmt.Errorf("tls: client requested unsupported application protocols (%s)", clientProtos)
//line /usr/local/go/src/crypto/tls/handshake_server.go:304
	// _ = "end of CoverTab[24006]"
}

// supportsECDHE returns whether ECDHE key exchanges can be used with this
//line /usr/local/go/src/crypto/tls/handshake_server.go:307
// pre-TLS 1.3 client.
//line /usr/local/go/src/crypto/tls/handshake_server.go:309
func supportsECDHE(c *Config, supportedCurves []CurveID, supportedPoints []uint8) bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:309
	_go_fuzz_dep_.CoverTab[24020]++
								supportsCurve := false
								for _, curve := range supportedCurves {
//line /usr/local/go/src/crypto/tls/handshake_server.go:311
		_go_fuzz_dep_.CoverTab[24024]++
									if c.supportsCurve(curve) {
//line /usr/local/go/src/crypto/tls/handshake_server.go:312
			_go_fuzz_dep_.CoverTab[24025]++
										supportsCurve = true
										break
//line /usr/local/go/src/crypto/tls/handshake_server.go:314
			// _ = "end of CoverTab[24025]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:315
			_go_fuzz_dep_.CoverTab[24026]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:315
			// _ = "end of CoverTab[24026]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:315
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:315
		// _ = "end of CoverTab[24024]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:316
	// _ = "end of CoverTab[24020]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:316
	_go_fuzz_dep_.CoverTab[24021]++

								supportsPointFormat := false
								for _, pointFormat := range supportedPoints {
//line /usr/local/go/src/crypto/tls/handshake_server.go:319
		_go_fuzz_dep_.CoverTab[24027]++
									if pointFormat == pointFormatUncompressed {
//line /usr/local/go/src/crypto/tls/handshake_server.go:320
			_go_fuzz_dep_.CoverTab[24028]++
										supportsPointFormat = true
										break
//line /usr/local/go/src/crypto/tls/handshake_server.go:322
			// _ = "end of CoverTab[24028]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:323
			_go_fuzz_dep_.CoverTab[24029]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:323
			// _ = "end of CoverTab[24029]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:323
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:323
		// _ = "end of CoverTab[24027]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:324
	// _ = "end of CoverTab[24021]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:324
	_go_fuzz_dep_.CoverTab[24022]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:329
	if len(supportedPoints) == 0 {
//line /usr/local/go/src/crypto/tls/handshake_server.go:329
		_go_fuzz_dep_.CoverTab[24030]++
									supportsPointFormat = true
//line /usr/local/go/src/crypto/tls/handshake_server.go:330
		// _ = "end of CoverTab[24030]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:331
		_go_fuzz_dep_.CoverTab[24031]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:331
		// _ = "end of CoverTab[24031]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:331
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:331
	// _ = "end of CoverTab[24022]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:331
	_go_fuzz_dep_.CoverTab[24023]++

								return supportsCurve && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:333
		_go_fuzz_dep_.CoverTab[24032]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:333
		return supportsPointFormat
//line /usr/local/go/src/crypto/tls/handshake_server.go:333
		// _ = "end of CoverTab[24032]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:333
	}()
//line /usr/local/go/src/crypto/tls/handshake_server.go:333
	// _ = "end of CoverTab[24023]"
}

func (hs *serverHandshakeState) pickCipherSuite() error {
//line /usr/local/go/src/crypto/tls/handshake_server.go:336
	_go_fuzz_dep_.CoverTab[24033]++
								c := hs.c

								preferenceOrder := cipherSuitesPreferenceOrder
								if !hasAESGCMHardwareSupport || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:340
		_go_fuzz_dep_.CoverTab[24038]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:340
		return !aesgcmPreferred(hs.clientHello.cipherSuites)
//line /usr/local/go/src/crypto/tls/handshake_server.go:340
		// _ = "end of CoverTab[24038]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:340
	}() {
//line /usr/local/go/src/crypto/tls/handshake_server.go:340
		_go_fuzz_dep_.CoverTab[24039]++
									preferenceOrder = cipherSuitesPreferenceOrderNoAES
//line /usr/local/go/src/crypto/tls/handshake_server.go:341
		// _ = "end of CoverTab[24039]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:342
		_go_fuzz_dep_.CoverTab[24040]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:342
		// _ = "end of CoverTab[24040]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:342
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:342
	// _ = "end of CoverTab[24033]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:342
	_go_fuzz_dep_.CoverTab[24034]++

								configCipherSuites := c.config.cipherSuites()
								preferenceList := make([]uint16, 0, len(configCipherSuites))
								for _, suiteID := range preferenceOrder {
//line /usr/local/go/src/crypto/tls/handshake_server.go:346
		_go_fuzz_dep_.CoverTab[24041]++
									for _, id := range configCipherSuites {
//line /usr/local/go/src/crypto/tls/handshake_server.go:347
			_go_fuzz_dep_.CoverTab[24042]++
										if id == suiteID {
//line /usr/local/go/src/crypto/tls/handshake_server.go:348
				_go_fuzz_dep_.CoverTab[24043]++
											preferenceList = append(preferenceList, id)
											break
//line /usr/local/go/src/crypto/tls/handshake_server.go:350
				// _ = "end of CoverTab[24043]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:351
				_go_fuzz_dep_.CoverTab[24044]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:351
				// _ = "end of CoverTab[24044]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:351
			}
//line /usr/local/go/src/crypto/tls/handshake_server.go:351
			// _ = "end of CoverTab[24042]"
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:352
		// _ = "end of CoverTab[24041]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:353
	// _ = "end of CoverTab[24034]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:353
	_go_fuzz_dep_.CoverTab[24035]++

								hs.suite = selectCipherSuite(preferenceList, hs.clientHello.cipherSuites, hs.cipherSuiteOk)
								if hs.suite == nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:356
		_go_fuzz_dep_.CoverTab[24045]++
									c.sendAlert(alertHandshakeFailure)
									return errors.New("tls: no cipher suite supported by both client and server")
//line /usr/local/go/src/crypto/tls/handshake_server.go:358
		// _ = "end of CoverTab[24045]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:359
		_go_fuzz_dep_.CoverTab[24046]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:359
		// _ = "end of CoverTab[24046]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:359
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:359
	// _ = "end of CoverTab[24035]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:359
	_go_fuzz_dep_.CoverTab[24036]++
								c.cipherSuite = hs.suite.id

								for _, id := range hs.clientHello.cipherSuites {
//line /usr/local/go/src/crypto/tls/handshake_server.go:362
		_go_fuzz_dep_.CoverTab[24047]++
									if id == TLS_FALLBACK_SCSV {
//line /usr/local/go/src/crypto/tls/handshake_server.go:363
			_go_fuzz_dep_.CoverTab[24048]++

										if hs.clientHello.vers < c.config.maxSupportedVersion(roleServer) {
//line /usr/local/go/src/crypto/tls/handshake_server.go:365
				_go_fuzz_dep_.CoverTab[24050]++
											c.sendAlert(alertInappropriateFallback)
											return errors.New("tls: client using inappropriate protocol fallback")
//line /usr/local/go/src/crypto/tls/handshake_server.go:367
				// _ = "end of CoverTab[24050]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:368
				_go_fuzz_dep_.CoverTab[24051]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:368
				// _ = "end of CoverTab[24051]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:368
			}
//line /usr/local/go/src/crypto/tls/handshake_server.go:368
			// _ = "end of CoverTab[24048]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:368
			_go_fuzz_dep_.CoverTab[24049]++
										break
//line /usr/local/go/src/crypto/tls/handshake_server.go:369
			// _ = "end of CoverTab[24049]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:370
			_go_fuzz_dep_.CoverTab[24052]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:370
			// _ = "end of CoverTab[24052]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:370
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:370
		// _ = "end of CoverTab[24047]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:371
	// _ = "end of CoverTab[24036]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:371
	_go_fuzz_dep_.CoverTab[24037]++

								return nil
//line /usr/local/go/src/crypto/tls/handshake_server.go:373
	// _ = "end of CoverTab[24037]"
}

func (hs *serverHandshakeState) cipherSuiteOk(c *cipherSuite) bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:376
	_go_fuzz_dep_.CoverTab[24053]++
								if c.flags&suiteECDHE != 0 {
//line /usr/local/go/src/crypto/tls/handshake_server.go:377
		_go_fuzz_dep_.CoverTab[24056]++
									if !hs.ecdheOk {
//line /usr/local/go/src/crypto/tls/handshake_server.go:378
			_go_fuzz_dep_.CoverTab[24058]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_server.go:379
			// _ = "end of CoverTab[24058]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:380
			_go_fuzz_dep_.CoverTab[24059]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:380
			// _ = "end of CoverTab[24059]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:380
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:380
		// _ = "end of CoverTab[24056]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:380
		_go_fuzz_dep_.CoverTab[24057]++
									if c.flags&suiteECSign != 0 {
//line /usr/local/go/src/crypto/tls/handshake_server.go:381
			_go_fuzz_dep_.CoverTab[24060]++
										if !hs.ecSignOk {
//line /usr/local/go/src/crypto/tls/handshake_server.go:382
				_go_fuzz_dep_.CoverTab[24061]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_server.go:383
				// _ = "end of CoverTab[24061]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:384
				_go_fuzz_dep_.CoverTab[24062]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:384
				// _ = "end of CoverTab[24062]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:384
			}
//line /usr/local/go/src/crypto/tls/handshake_server.go:384
			// _ = "end of CoverTab[24060]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:385
			_go_fuzz_dep_.CoverTab[24063]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:385
			if !hs.rsaSignOk {
//line /usr/local/go/src/crypto/tls/handshake_server.go:385
				_go_fuzz_dep_.CoverTab[24064]++
											return false
//line /usr/local/go/src/crypto/tls/handshake_server.go:386
				// _ = "end of CoverTab[24064]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:387
				_go_fuzz_dep_.CoverTab[24065]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:387
				// _ = "end of CoverTab[24065]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:387
			}
//line /usr/local/go/src/crypto/tls/handshake_server.go:387
			// _ = "end of CoverTab[24063]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:387
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:387
		// _ = "end of CoverTab[24057]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:388
		_go_fuzz_dep_.CoverTab[24066]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:388
		if !hs.rsaDecryptOk {
//line /usr/local/go/src/crypto/tls/handshake_server.go:388
			_go_fuzz_dep_.CoverTab[24067]++
										return false
//line /usr/local/go/src/crypto/tls/handshake_server.go:389
			// _ = "end of CoverTab[24067]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:390
			_go_fuzz_dep_.CoverTab[24068]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:390
			// _ = "end of CoverTab[24068]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:390
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:390
		// _ = "end of CoverTab[24066]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:390
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:390
	// _ = "end of CoverTab[24053]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:390
	_go_fuzz_dep_.CoverTab[24054]++
								if hs.c.vers < VersionTLS12 && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:391
		_go_fuzz_dep_.CoverTab[24069]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:391
		return c.flags&suiteTLS12 != 0
//line /usr/local/go/src/crypto/tls/handshake_server.go:391
		// _ = "end of CoverTab[24069]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:391
	}() {
//line /usr/local/go/src/crypto/tls/handshake_server.go:391
		_go_fuzz_dep_.CoverTab[24070]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_server.go:392
		// _ = "end of CoverTab[24070]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:393
		_go_fuzz_dep_.CoverTab[24071]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:393
		// _ = "end of CoverTab[24071]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:393
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:393
	// _ = "end of CoverTab[24054]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:393
	_go_fuzz_dep_.CoverTab[24055]++
								return true
//line /usr/local/go/src/crypto/tls/handshake_server.go:394
	// _ = "end of CoverTab[24055]"
}

// checkForResumption reports whether we should perform resumption on this connection.
func (hs *serverHandshakeState) checkForResumption() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:398
	_go_fuzz_dep_.CoverTab[24072]++
								c := hs.c

								if c.config.SessionTicketsDisabled {
//line /usr/local/go/src/crypto/tls/handshake_server.go:401
		_go_fuzz_dep_.CoverTab[24083]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_server.go:402
		// _ = "end of CoverTab[24083]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:403
		_go_fuzz_dep_.CoverTab[24084]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:403
		// _ = "end of CoverTab[24084]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:403
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:403
	// _ = "end of CoverTab[24072]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:403
	_go_fuzz_dep_.CoverTab[24073]++

								plaintext, usedOldKey := c.decryptTicket(hs.clientHello.sessionTicket)
								if plaintext == nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:406
		_go_fuzz_dep_.CoverTab[24085]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_server.go:407
		// _ = "end of CoverTab[24085]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:408
		_go_fuzz_dep_.CoverTab[24086]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:408
		// _ = "end of CoverTab[24086]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:408
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:408
	// _ = "end of CoverTab[24073]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:408
	_go_fuzz_dep_.CoverTab[24074]++
								hs.sessionState = &sessionState{usedOldKey: usedOldKey}
								ok := hs.sessionState.unmarshal(plaintext)
								if !ok {
//line /usr/local/go/src/crypto/tls/handshake_server.go:411
		_go_fuzz_dep_.CoverTab[24087]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_server.go:412
		// _ = "end of CoverTab[24087]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:413
		_go_fuzz_dep_.CoverTab[24088]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:413
		// _ = "end of CoverTab[24088]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:413
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:413
	// _ = "end of CoverTab[24074]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:413
	_go_fuzz_dep_.CoverTab[24075]++

								createdAt := time.Unix(int64(hs.sessionState.createdAt), 0)
								if c.config.time().Sub(createdAt) > maxSessionTicketLifetime {
//line /usr/local/go/src/crypto/tls/handshake_server.go:416
		_go_fuzz_dep_.CoverTab[24089]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_server.go:417
		// _ = "end of CoverTab[24089]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:418
		_go_fuzz_dep_.CoverTab[24090]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:418
		// _ = "end of CoverTab[24090]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:418
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:418
	// _ = "end of CoverTab[24075]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:418
	_go_fuzz_dep_.CoverTab[24076]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:421
	if c.vers != hs.sessionState.vers {
//line /usr/local/go/src/crypto/tls/handshake_server.go:421
		_go_fuzz_dep_.CoverTab[24091]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_server.go:422
		// _ = "end of CoverTab[24091]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:423
		_go_fuzz_dep_.CoverTab[24092]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:423
		// _ = "end of CoverTab[24092]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:423
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:423
	// _ = "end of CoverTab[24076]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:423
	_go_fuzz_dep_.CoverTab[24077]++

								cipherSuiteOk := false

								for _, id := range hs.clientHello.cipherSuites {
//line /usr/local/go/src/crypto/tls/handshake_server.go:427
		_go_fuzz_dep_.CoverTab[24093]++
									if id == hs.sessionState.cipherSuite {
//line /usr/local/go/src/crypto/tls/handshake_server.go:428
			_go_fuzz_dep_.CoverTab[24094]++
										cipherSuiteOk = true
										break
//line /usr/local/go/src/crypto/tls/handshake_server.go:430
			// _ = "end of CoverTab[24094]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:431
			_go_fuzz_dep_.CoverTab[24095]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:431
			// _ = "end of CoverTab[24095]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:431
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:431
		// _ = "end of CoverTab[24093]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:432
	// _ = "end of CoverTab[24077]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:432
	_go_fuzz_dep_.CoverTab[24078]++
								if !cipherSuiteOk {
//line /usr/local/go/src/crypto/tls/handshake_server.go:433
		_go_fuzz_dep_.CoverTab[24096]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_server.go:434
		// _ = "end of CoverTab[24096]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:435
		_go_fuzz_dep_.CoverTab[24097]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:435
		// _ = "end of CoverTab[24097]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:435
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:435
	// _ = "end of CoverTab[24078]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:435
	_go_fuzz_dep_.CoverTab[24079]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:438
	hs.suite = selectCipherSuite([]uint16{hs.sessionState.cipherSuite},
		c.config.cipherSuites(), hs.cipherSuiteOk)
	if hs.suite == nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:440
		_go_fuzz_dep_.CoverTab[24098]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_server.go:441
		// _ = "end of CoverTab[24098]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:442
		_go_fuzz_dep_.CoverTab[24099]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:442
		// _ = "end of CoverTab[24099]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:442
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:442
	// _ = "end of CoverTab[24079]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:442
	_go_fuzz_dep_.CoverTab[24080]++

								sessionHasClientCerts := len(hs.sessionState.certificates) != 0
								needClientCerts := requiresClientCert(c.config.ClientAuth)
								if needClientCerts && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:446
		_go_fuzz_dep_.CoverTab[24100]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:446
		return !sessionHasClientCerts
//line /usr/local/go/src/crypto/tls/handshake_server.go:446
		// _ = "end of CoverTab[24100]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:446
	}() {
//line /usr/local/go/src/crypto/tls/handshake_server.go:446
		_go_fuzz_dep_.CoverTab[24101]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_server.go:447
		// _ = "end of CoverTab[24101]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:448
		_go_fuzz_dep_.CoverTab[24102]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:448
		// _ = "end of CoverTab[24102]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:448
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:448
	// _ = "end of CoverTab[24080]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:448
	_go_fuzz_dep_.CoverTab[24081]++
								if sessionHasClientCerts && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:449
		_go_fuzz_dep_.CoverTab[24103]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:449
		return c.config.ClientAuth == NoClientCert
//line /usr/local/go/src/crypto/tls/handshake_server.go:449
		// _ = "end of CoverTab[24103]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:449
	}() {
//line /usr/local/go/src/crypto/tls/handshake_server.go:449
		_go_fuzz_dep_.CoverTab[24104]++
									return false
//line /usr/local/go/src/crypto/tls/handshake_server.go:450
		// _ = "end of CoverTab[24104]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:451
		_go_fuzz_dep_.CoverTab[24105]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:451
		// _ = "end of CoverTab[24105]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:451
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:451
	// _ = "end of CoverTab[24081]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:451
	_go_fuzz_dep_.CoverTab[24082]++

								return true
//line /usr/local/go/src/crypto/tls/handshake_server.go:453
	// _ = "end of CoverTab[24082]"
}

func (hs *serverHandshakeState) doResumeHandshake() error {
//line /usr/local/go/src/crypto/tls/handshake_server.go:456
	_go_fuzz_dep_.CoverTab[24106]++
								c := hs.c

								hs.hello.cipherSuite = hs.suite.id
								c.cipherSuite = hs.suite.id

//line /usr/local/go/src/crypto/tls/handshake_server.go:463
	hs.hello.sessionId = hs.clientHello.sessionId
	hs.hello.ticketSupported = hs.sessionState.usedOldKey
	hs.finishedHash = newFinishedHash(c.vers, hs.suite)
	hs.finishedHash.discardHandshakeBuffer()
	if err := transcriptMsg(hs.clientHello, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:467
		_go_fuzz_dep_.CoverTab[24111]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:468
		// _ = "end of CoverTab[24111]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:469
		_go_fuzz_dep_.CoverTab[24112]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:469
		// _ = "end of CoverTab[24112]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:469
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:469
	// _ = "end of CoverTab[24106]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:469
	_go_fuzz_dep_.CoverTab[24107]++
								if _, err := hs.c.writeHandshakeRecord(hs.hello, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:470
		_go_fuzz_dep_.CoverTab[24113]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:471
		// _ = "end of CoverTab[24113]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:472
		_go_fuzz_dep_.CoverTab[24114]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:472
		// _ = "end of CoverTab[24114]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:472
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:472
	// _ = "end of CoverTab[24107]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:472
	_go_fuzz_dep_.CoverTab[24108]++

								if err := c.processCertsFromClient(Certificate{
		Certificate: hs.sessionState.certificates,
	}); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:476
		_go_fuzz_dep_.CoverTab[24115]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:477
		// _ = "end of CoverTab[24115]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:478
		_go_fuzz_dep_.CoverTab[24116]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:478
		// _ = "end of CoverTab[24116]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:478
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:478
	// _ = "end of CoverTab[24108]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:478
	_go_fuzz_dep_.CoverTab[24109]++

								if c.config.VerifyConnection != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:480
		_go_fuzz_dep_.CoverTab[24117]++
									if err := c.config.VerifyConnection(c.connectionStateLocked()); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:481
			_go_fuzz_dep_.CoverTab[24118]++
										c.sendAlert(alertBadCertificate)
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:483
			// _ = "end of CoverTab[24118]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:484
			_go_fuzz_dep_.CoverTab[24119]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:484
			// _ = "end of CoverTab[24119]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:484
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:484
		// _ = "end of CoverTab[24117]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:485
		_go_fuzz_dep_.CoverTab[24120]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:485
		// _ = "end of CoverTab[24120]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:485
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:485
	// _ = "end of CoverTab[24109]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:485
	_go_fuzz_dep_.CoverTab[24110]++

								hs.masterSecret = hs.sessionState.masterSecret

								return nil
//line /usr/local/go/src/crypto/tls/handshake_server.go:489
	// _ = "end of CoverTab[24110]"
}

func (hs *serverHandshakeState) doFullHandshake() error {
//line /usr/local/go/src/crypto/tls/handshake_server.go:492
	_go_fuzz_dep_.CoverTab[24121]++
								c := hs.c

								if hs.clientHello.ocspStapling && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:495
		_go_fuzz_dep_.CoverTab[24140]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:495
		return len(hs.cert.OCSPStaple) > 0
//line /usr/local/go/src/crypto/tls/handshake_server.go:495
		// _ = "end of CoverTab[24140]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:495
	}() {
//line /usr/local/go/src/crypto/tls/handshake_server.go:495
		_go_fuzz_dep_.CoverTab[24141]++
									hs.hello.ocspStapling = true
//line /usr/local/go/src/crypto/tls/handshake_server.go:496
		// _ = "end of CoverTab[24141]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:497
		_go_fuzz_dep_.CoverTab[24142]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:497
		// _ = "end of CoverTab[24142]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:497
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:497
	// _ = "end of CoverTab[24121]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:497
	_go_fuzz_dep_.CoverTab[24122]++

								hs.hello.ticketSupported = hs.clientHello.ticketSupported && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:499
		_go_fuzz_dep_.CoverTab[24143]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:499
		return !c.config.SessionTicketsDisabled
//line /usr/local/go/src/crypto/tls/handshake_server.go:499
		// _ = "end of CoverTab[24143]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:499
	}()
								hs.hello.cipherSuite = hs.suite.id

								hs.finishedHash = newFinishedHash(hs.c.vers, hs.suite)
								if c.config.ClientAuth == NoClientCert {
//line /usr/local/go/src/crypto/tls/handshake_server.go:503
		_go_fuzz_dep_.CoverTab[24144]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:506
		hs.finishedHash.discardHandshakeBuffer()
//line /usr/local/go/src/crypto/tls/handshake_server.go:506
		// _ = "end of CoverTab[24144]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:507
		_go_fuzz_dep_.CoverTab[24145]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:507
		// _ = "end of CoverTab[24145]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:507
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:507
	// _ = "end of CoverTab[24122]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:507
	_go_fuzz_dep_.CoverTab[24123]++
								if err := transcriptMsg(hs.clientHello, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:508
		_go_fuzz_dep_.CoverTab[24146]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:509
		// _ = "end of CoverTab[24146]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:510
		_go_fuzz_dep_.CoverTab[24147]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:510
		// _ = "end of CoverTab[24147]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:510
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:510
	// _ = "end of CoverTab[24123]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:510
	_go_fuzz_dep_.CoverTab[24124]++
								if _, err := hs.c.writeHandshakeRecord(hs.hello, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:511
		_go_fuzz_dep_.CoverTab[24148]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:512
		// _ = "end of CoverTab[24148]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:513
		_go_fuzz_dep_.CoverTab[24149]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:513
		// _ = "end of CoverTab[24149]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:513
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:513
	// _ = "end of CoverTab[24124]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:513
	_go_fuzz_dep_.CoverTab[24125]++

								certMsg := new(certificateMsg)
								certMsg.certificates = hs.cert.Certificate
								if _, err := hs.c.writeHandshakeRecord(certMsg, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:517
		_go_fuzz_dep_.CoverTab[24150]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:518
		// _ = "end of CoverTab[24150]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:519
		_go_fuzz_dep_.CoverTab[24151]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:519
		// _ = "end of CoverTab[24151]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:519
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:519
	// _ = "end of CoverTab[24125]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:519
	_go_fuzz_dep_.CoverTab[24126]++

								if hs.hello.ocspStapling {
//line /usr/local/go/src/crypto/tls/handshake_server.go:521
		_go_fuzz_dep_.CoverTab[24152]++
									certStatus := new(certificateStatusMsg)
									certStatus.response = hs.cert.OCSPStaple
									if _, err := hs.c.writeHandshakeRecord(certStatus, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:524
			_go_fuzz_dep_.CoverTab[24153]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:525
			// _ = "end of CoverTab[24153]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:526
			_go_fuzz_dep_.CoverTab[24154]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:526
			// _ = "end of CoverTab[24154]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:526
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:526
		// _ = "end of CoverTab[24152]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:527
		_go_fuzz_dep_.CoverTab[24155]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:527
		// _ = "end of CoverTab[24155]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:527
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:527
	// _ = "end of CoverTab[24126]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:527
	_go_fuzz_dep_.CoverTab[24127]++

								keyAgreement := hs.suite.ka(c.vers)
								skx, err := keyAgreement.generateServerKeyExchange(c.config, hs.cert, hs.clientHello, hs.hello)
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:531
		_go_fuzz_dep_.CoverTab[24156]++
									c.sendAlert(alertHandshakeFailure)
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:533
		// _ = "end of CoverTab[24156]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:534
		_go_fuzz_dep_.CoverTab[24157]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:534
		// _ = "end of CoverTab[24157]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:534
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:534
	// _ = "end of CoverTab[24127]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:534
	_go_fuzz_dep_.CoverTab[24128]++
								if skx != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:535
		_go_fuzz_dep_.CoverTab[24158]++
									if _, err := hs.c.writeHandshakeRecord(skx, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:536
			_go_fuzz_dep_.CoverTab[24159]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:537
			// _ = "end of CoverTab[24159]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:538
			_go_fuzz_dep_.CoverTab[24160]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:538
			// _ = "end of CoverTab[24160]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:538
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:538
		// _ = "end of CoverTab[24158]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:539
		_go_fuzz_dep_.CoverTab[24161]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:539
		// _ = "end of CoverTab[24161]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:539
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:539
	// _ = "end of CoverTab[24128]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:539
	_go_fuzz_dep_.CoverTab[24129]++

								var certReq *certificateRequestMsg
								if c.config.ClientAuth >= RequestClientCert {
//line /usr/local/go/src/crypto/tls/handshake_server.go:542
		_go_fuzz_dep_.CoverTab[24162]++

									certReq = new(certificateRequestMsg)
									certReq.certificateTypes = []byte{
			byte(certTypeRSASign),
			byte(certTypeECDSASign),
		}
		if c.vers >= VersionTLS12 {
//line /usr/local/go/src/crypto/tls/handshake_server.go:549
			_go_fuzz_dep_.CoverTab[24165]++
										certReq.hasSignatureAlgorithm = true
										certReq.supportedSignatureAlgorithms = supportedSignatureAlgorithms()
//line /usr/local/go/src/crypto/tls/handshake_server.go:551
			// _ = "end of CoverTab[24165]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:552
			_go_fuzz_dep_.CoverTab[24166]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:552
			// _ = "end of CoverTab[24166]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:552
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:552
		// _ = "end of CoverTab[24162]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:552
		_go_fuzz_dep_.CoverTab[24163]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:559
		if c.config.ClientCAs != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:559
			_go_fuzz_dep_.CoverTab[24167]++
										certReq.certificateAuthorities = c.config.ClientCAs.Subjects()
//line /usr/local/go/src/crypto/tls/handshake_server.go:560
			// _ = "end of CoverTab[24167]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:561
			_go_fuzz_dep_.CoverTab[24168]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:561
			// _ = "end of CoverTab[24168]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:561
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:561
		// _ = "end of CoverTab[24163]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:561
		_go_fuzz_dep_.CoverTab[24164]++
									if _, err := hs.c.writeHandshakeRecord(certReq, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:562
			_go_fuzz_dep_.CoverTab[24169]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:563
			// _ = "end of CoverTab[24169]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:564
			_go_fuzz_dep_.CoverTab[24170]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:564
			// _ = "end of CoverTab[24170]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:564
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:564
		// _ = "end of CoverTab[24164]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:565
		_go_fuzz_dep_.CoverTab[24171]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:565
		// _ = "end of CoverTab[24171]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:565
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:565
	// _ = "end of CoverTab[24129]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:565
	_go_fuzz_dep_.CoverTab[24130]++

								helloDone := new(serverHelloDoneMsg)
								if _, err := hs.c.writeHandshakeRecord(helloDone, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:568
		_go_fuzz_dep_.CoverTab[24172]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:569
		// _ = "end of CoverTab[24172]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:570
		_go_fuzz_dep_.CoverTab[24173]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:570
		// _ = "end of CoverTab[24173]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:570
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:570
	// _ = "end of CoverTab[24130]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:570
	_go_fuzz_dep_.CoverTab[24131]++

								if _, err := c.flush(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:572
		_go_fuzz_dep_.CoverTab[24174]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:573
		// _ = "end of CoverTab[24174]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:574
		_go_fuzz_dep_.CoverTab[24175]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:574
		// _ = "end of CoverTab[24175]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:574
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:574
	// _ = "end of CoverTab[24131]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:574
	_go_fuzz_dep_.CoverTab[24132]++

								var pub crypto.PublicKey	// public key for client auth, if any

								msg, err := c.readHandshake(&hs.finishedHash)
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:579
		_go_fuzz_dep_.CoverTab[24176]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:580
		// _ = "end of CoverTab[24176]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:581
		_go_fuzz_dep_.CoverTab[24177]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:581
		// _ = "end of CoverTab[24177]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:581
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:581
	// _ = "end of CoverTab[24132]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:581
	_go_fuzz_dep_.CoverTab[24133]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:585
	if c.config.ClientAuth >= RequestClientCert {
//line /usr/local/go/src/crypto/tls/handshake_server.go:585
		_go_fuzz_dep_.CoverTab[24178]++
									certMsg, ok := msg.(*certificateMsg)
									if !ok {
//line /usr/local/go/src/crypto/tls/handshake_server.go:587
			_go_fuzz_dep_.CoverTab[24182]++
										c.sendAlert(alertUnexpectedMessage)
										return unexpectedMessageError(certMsg, msg)
//line /usr/local/go/src/crypto/tls/handshake_server.go:589
			// _ = "end of CoverTab[24182]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:590
			_go_fuzz_dep_.CoverTab[24183]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:590
			// _ = "end of CoverTab[24183]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:590
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:590
		// _ = "end of CoverTab[24178]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:590
		_go_fuzz_dep_.CoverTab[24179]++

									if err := c.processCertsFromClient(Certificate{
			Certificate: certMsg.certificates,
		}); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:594
			_go_fuzz_dep_.CoverTab[24184]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:595
			// _ = "end of CoverTab[24184]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:596
			_go_fuzz_dep_.CoverTab[24185]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:596
			// _ = "end of CoverTab[24185]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:596
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:596
		// _ = "end of CoverTab[24179]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:596
		_go_fuzz_dep_.CoverTab[24180]++
									if len(certMsg.certificates) != 0 {
//line /usr/local/go/src/crypto/tls/handshake_server.go:597
			_go_fuzz_dep_.CoverTab[24186]++
										pub = c.peerCertificates[0].PublicKey
//line /usr/local/go/src/crypto/tls/handshake_server.go:598
			// _ = "end of CoverTab[24186]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:599
			_go_fuzz_dep_.CoverTab[24187]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:599
			// _ = "end of CoverTab[24187]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:599
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:599
		// _ = "end of CoverTab[24180]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:599
		_go_fuzz_dep_.CoverTab[24181]++

									msg, err = c.readHandshake(&hs.finishedHash)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:602
			_go_fuzz_dep_.CoverTab[24188]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:603
			// _ = "end of CoverTab[24188]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:604
			_go_fuzz_dep_.CoverTab[24189]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:604
			// _ = "end of CoverTab[24189]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:604
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:604
		// _ = "end of CoverTab[24181]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:605
		_go_fuzz_dep_.CoverTab[24190]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:605
		// _ = "end of CoverTab[24190]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:605
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:605
	// _ = "end of CoverTab[24133]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:605
	_go_fuzz_dep_.CoverTab[24134]++
								if c.config.VerifyConnection != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:606
		_go_fuzz_dep_.CoverTab[24191]++
									if err := c.config.VerifyConnection(c.connectionStateLocked()); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:607
			_go_fuzz_dep_.CoverTab[24192]++
										c.sendAlert(alertBadCertificate)
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:609
			// _ = "end of CoverTab[24192]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:610
			_go_fuzz_dep_.CoverTab[24193]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:610
			// _ = "end of CoverTab[24193]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:610
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:610
		// _ = "end of CoverTab[24191]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:611
		_go_fuzz_dep_.CoverTab[24194]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:611
		// _ = "end of CoverTab[24194]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:611
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:611
	// _ = "end of CoverTab[24134]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:611
	_go_fuzz_dep_.CoverTab[24135]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:614
	ckx, ok := msg.(*clientKeyExchangeMsg)
	if !ok {
//line /usr/local/go/src/crypto/tls/handshake_server.go:615
		_go_fuzz_dep_.CoverTab[24195]++
									c.sendAlert(alertUnexpectedMessage)
									return unexpectedMessageError(ckx, msg)
//line /usr/local/go/src/crypto/tls/handshake_server.go:617
		// _ = "end of CoverTab[24195]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:618
		_go_fuzz_dep_.CoverTab[24196]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:618
		// _ = "end of CoverTab[24196]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:618
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:618
	// _ = "end of CoverTab[24135]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:618
	_go_fuzz_dep_.CoverTab[24136]++

								preMasterSecret, err := keyAgreement.processClientKeyExchange(c.config, hs.cert, ckx, c.vers)
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:621
		_go_fuzz_dep_.CoverTab[24197]++
									c.sendAlert(alertHandshakeFailure)
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:623
		// _ = "end of CoverTab[24197]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:624
		_go_fuzz_dep_.CoverTab[24198]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:624
		// _ = "end of CoverTab[24198]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:624
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:624
	// _ = "end of CoverTab[24136]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:624
	_go_fuzz_dep_.CoverTab[24137]++
								hs.masterSecret = masterFromPreMasterSecret(c.vers, hs.suite, preMasterSecret, hs.clientHello.random, hs.hello.random)
								if err := c.config.writeKeyLog(keyLogLabelTLS12, hs.clientHello.random, hs.masterSecret); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:626
		_go_fuzz_dep_.CoverTab[24199]++
									c.sendAlert(alertInternalError)
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:628
		// _ = "end of CoverTab[24199]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:629
		_go_fuzz_dep_.CoverTab[24200]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:629
		// _ = "end of CoverTab[24200]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:629
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:629
	// _ = "end of CoverTab[24137]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:629
	_go_fuzz_dep_.CoverTab[24138]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:637
	if len(c.peerCertificates) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_server.go:637
		_go_fuzz_dep_.CoverTab[24201]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:641
		msg, err = c.readHandshake(nil)
		if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:642
			_go_fuzz_dep_.CoverTab[24206]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:643
			// _ = "end of CoverTab[24206]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:644
			_go_fuzz_dep_.CoverTab[24207]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:644
			// _ = "end of CoverTab[24207]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:644
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:644
		// _ = "end of CoverTab[24201]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:644
		_go_fuzz_dep_.CoverTab[24202]++
									certVerify, ok := msg.(*certificateVerifyMsg)
									if !ok {
//line /usr/local/go/src/crypto/tls/handshake_server.go:646
			_go_fuzz_dep_.CoverTab[24208]++
										c.sendAlert(alertUnexpectedMessage)
										return unexpectedMessageError(certVerify, msg)
//line /usr/local/go/src/crypto/tls/handshake_server.go:648
			// _ = "end of CoverTab[24208]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:649
			_go_fuzz_dep_.CoverTab[24209]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:649
			// _ = "end of CoverTab[24209]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:649
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:649
		// _ = "end of CoverTab[24202]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:649
		_go_fuzz_dep_.CoverTab[24203]++

									var sigType uint8
									var sigHash crypto.Hash
									if c.vers >= VersionTLS12 {
//line /usr/local/go/src/crypto/tls/handshake_server.go:653
			_go_fuzz_dep_.CoverTab[24210]++
										if !isSupportedSignatureAlgorithm(certVerify.signatureAlgorithm, certReq.supportedSignatureAlgorithms) {
//line /usr/local/go/src/crypto/tls/handshake_server.go:654
				_go_fuzz_dep_.CoverTab[24212]++
											c.sendAlert(alertIllegalParameter)
											return errors.New("tls: client certificate used with invalid signature algorithm")
//line /usr/local/go/src/crypto/tls/handshake_server.go:656
				// _ = "end of CoverTab[24212]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:657
				_go_fuzz_dep_.CoverTab[24213]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:657
				// _ = "end of CoverTab[24213]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:657
			}
//line /usr/local/go/src/crypto/tls/handshake_server.go:657
			// _ = "end of CoverTab[24210]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:657
			_go_fuzz_dep_.CoverTab[24211]++
										sigType, sigHash, err = typeAndHashFromSignatureScheme(certVerify.signatureAlgorithm)
										if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:659
				_go_fuzz_dep_.CoverTab[24214]++
											return c.sendAlert(alertInternalError)
//line /usr/local/go/src/crypto/tls/handshake_server.go:660
				// _ = "end of CoverTab[24214]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:661
				_go_fuzz_dep_.CoverTab[24215]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:661
				// _ = "end of CoverTab[24215]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:661
			}
//line /usr/local/go/src/crypto/tls/handshake_server.go:661
			// _ = "end of CoverTab[24211]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:662
			_go_fuzz_dep_.CoverTab[24216]++
										sigType, sigHash, err = legacyTypeAndHashFromPublicKey(pub)
										if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:664
				_go_fuzz_dep_.CoverTab[24217]++
											c.sendAlert(alertIllegalParameter)
											return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:666
				// _ = "end of CoverTab[24217]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:667
				_go_fuzz_dep_.CoverTab[24218]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:667
				// _ = "end of CoverTab[24218]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:667
			}
//line /usr/local/go/src/crypto/tls/handshake_server.go:667
			// _ = "end of CoverTab[24216]"
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:668
		// _ = "end of CoverTab[24203]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:668
		_go_fuzz_dep_.CoverTab[24204]++

									signed := hs.finishedHash.hashForClientCertificate(sigType, sigHash)
									if err := verifyHandshakeSignature(sigType, pub, sigHash, signed, certVerify.signature); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:671
			_go_fuzz_dep_.CoverTab[24219]++
										c.sendAlert(alertDecryptError)
										return errors.New("tls: invalid signature by the client certificate: " + err.Error())
//line /usr/local/go/src/crypto/tls/handshake_server.go:673
			// _ = "end of CoverTab[24219]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:674
			_go_fuzz_dep_.CoverTab[24220]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:674
			// _ = "end of CoverTab[24220]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:674
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:674
		// _ = "end of CoverTab[24204]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:674
		_go_fuzz_dep_.CoverTab[24205]++

									if err := transcriptMsg(certVerify, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:676
			_go_fuzz_dep_.CoverTab[24221]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:677
			// _ = "end of CoverTab[24221]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:678
			_go_fuzz_dep_.CoverTab[24222]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:678
			// _ = "end of CoverTab[24222]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:678
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:678
		// _ = "end of CoverTab[24205]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:679
		_go_fuzz_dep_.CoverTab[24223]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:679
		// _ = "end of CoverTab[24223]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:679
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:679
	// _ = "end of CoverTab[24138]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:679
	_go_fuzz_dep_.CoverTab[24139]++

								hs.finishedHash.discardHandshakeBuffer()

								return nil
//line /usr/local/go/src/crypto/tls/handshake_server.go:683
	// _ = "end of CoverTab[24139]"
}

func (hs *serverHandshakeState) establishKeys() error {
//line /usr/local/go/src/crypto/tls/handshake_server.go:686
	_go_fuzz_dep_.CoverTab[24224]++
								c := hs.c

								clientMAC, serverMAC, clientKey, serverKey, clientIV, serverIV :=
		keysFromMasterSecret(c.vers, hs.suite, hs.masterSecret, hs.clientHello.random, hs.hello.random, hs.suite.macLen, hs.suite.keyLen, hs.suite.ivLen)

	var clientCipher, serverCipher any
	var clientHash, serverHash hash.Hash

	if hs.suite.aead == nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:695
		_go_fuzz_dep_.CoverTab[24226]++
									clientCipher = hs.suite.cipher(clientKey, clientIV, true)
									clientHash = hs.suite.mac(clientMAC)
									serverCipher = hs.suite.cipher(serverKey, serverIV, false)
									serverHash = hs.suite.mac(serverMAC)
//line /usr/local/go/src/crypto/tls/handshake_server.go:699
		// _ = "end of CoverTab[24226]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:700
		_go_fuzz_dep_.CoverTab[24227]++
									clientCipher = hs.suite.aead(clientKey, clientIV)
									serverCipher = hs.suite.aead(serverKey, serverIV)
//line /usr/local/go/src/crypto/tls/handshake_server.go:702
		// _ = "end of CoverTab[24227]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:703
	// _ = "end of CoverTab[24224]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:703
	_go_fuzz_dep_.CoverTab[24225]++

								c.in.prepareCipherSpec(c.vers, clientCipher, clientHash)
								c.out.prepareCipherSpec(c.vers, serverCipher, serverHash)

								return nil
//line /usr/local/go/src/crypto/tls/handshake_server.go:708
	// _ = "end of CoverTab[24225]"
}

func (hs *serverHandshakeState) readFinished(out []byte) error {
//line /usr/local/go/src/crypto/tls/handshake_server.go:711
	_go_fuzz_dep_.CoverTab[24228]++
								c := hs.c

								if err := c.readChangeCipherSpec(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:714
		_go_fuzz_dep_.CoverTab[24234]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:715
		// _ = "end of CoverTab[24234]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:716
		_go_fuzz_dep_.CoverTab[24235]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:716
		// _ = "end of CoverTab[24235]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:716
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:716
	// _ = "end of CoverTab[24228]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:716
	_go_fuzz_dep_.CoverTab[24229]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:721
	msg, err := c.readHandshake(nil)
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:722
		_go_fuzz_dep_.CoverTab[24236]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:723
		// _ = "end of CoverTab[24236]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:724
		_go_fuzz_dep_.CoverTab[24237]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:724
		// _ = "end of CoverTab[24237]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:724
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:724
	// _ = "end of CoverTab[24229]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:724
	_go_fuzz_dep_.CoverTab[24230]++
								clientFinished, ok := msg.(*finishedMsg)
								if !ok {
//line /usr/local/go/src/crypto/tls/handshake_server.go:726
		_go_fuzz_dep_.CoverTab[24238]++
									c.sendAlert(alertUnexpectedMessage)
									return unexpectedMessageError(clientFinished, msg)
//line /usr/local/go/src/crypto/tls/handshake_server.go:728
		// _ = "end of CoverTab[24238]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:729
		_go_fuzz_dep_.CoverTab[24239]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:729
		// _ = "end of CoverTab[24239]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:729
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:729
	// _ = "end of CoverTab[24230]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:729
	_go_fuzz_dep_.CoverTab[24231]++

								verify := hs.finishedHash.clientSum(hs.masterSecret)
								if len(verify) != len(clientFinished.verifyData) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:732
		_go_fuzz_dep_.CoverTab[24240]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:732
		return subtle.ConstantTimeCompare(verify, clientFinished.verifyData) != 1
									// _ = "end of CoverTab[24240]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:733
	}() {
//line /usr/local/go/src/crypto/tls/handshake_server.go:733
		_go_fuzz_dep_.CoverTab[24241]++
									c.sendAlert(alertHandshakeFailure)
									return errors.New("tls: client's Finished message is incorrect")
//line /usr/local/go/src/crypto/tls/handshake_server.go:735
		// _ = "end of CoverTab[24241]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:736
		_go_fuzz_dep_.CoverTab[24242]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:736
		// _ = "end of CoverTab[24242]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:736
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:736
	// _ = "end of CoverTab[24231]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:736
	_go_fuzz_dep_.CoverTab[24232]++

								if err := transcriptMsg(clientFinished, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:738
		_go_fuzz_dep_.CoverTab[24243]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:739
		// _ = "end of CoverTab[24243]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:740
		_go_fuzz_dep_.CoverTab[24244]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:740
		// _ = "end of CoverTab[24244]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:740
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:740
	// _ = "end of CoverTab[24232]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:740
	_go_fuzz_dep_.CoverTab[24233]++

								copy(out, verify)
								return nil
//line /usr/local/go/src/crypto/tls/handshake_server.go:743
	// _ = "end of CoverTab[24233]"
}

func (hs *serverHandshakeState) sendSessionTicket() error {
//line /usr/local/go/src/crypto/tls/handshake_server.go:746
	_go_fuzz_dep_.CoverTab[24245]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:750
	if !hs.hello.ticketSupported {
//line /usr/local/go/src/crypto/tls/handshake_server.go:750
		_go_fuzz_dep_.CoverTab[24252]++
									return nil
//line /usr/local/go/src/crypto/tls/handshake_server.go:751
		// _ = "end of CoverTab[24252]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:752
		_go_fuzz_dep_.CoverTab[24253]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:752
		// _ = "end of CoverTab[24253]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:752
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:752
	// _ = "end of CoverTab[24245]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:752
	_go_fuzz_dep_.CoverTab[24246]++

								c := hs.c
								m := new(newSessionTicketMsg)

								createdAt := uint64(c.config.time().Unix())
								if hs.sessionState != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:758
		_go_fuzz_dep_.CoverTab[24254]++

//line /usr/local/go/src/crypto/tls/handshake_server.go:761
		createdAt = hs.sessionState.createdAt
//line /usr/local/go/src/crypto/tls/handshake_server.go:761
		// _ = "end of CoverTab[24254]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:762
		_go_fuzz_dep_.CoverTab[24255]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:762
		// _ = "end of CoverTab[24255]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:762
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:762
	// _ = "end of CoverTab[24246]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:762
	_go_fuzz_dep_.CoverTab[24247]++

								var certsFromClient [][]byte
								for _, cert := range c.peerCertificates {
//line /usr/local/go/src/crypto/tls/handshake_server.go:765
		_go_fuzz_dep_.CoverTab[24256]++
									certsFromClient = append(certsFromClient, cert.Raw)
//line /usr/local/go/src/crypto/tls/handshake_server.go:766
		// _ = "end of CoverTab[24256]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:767
	// _ = "end of CoverTab[24247]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:767
	_go_fuzz_dep_.CoverTab[24248]++
								state := sessionState{
		vers:		c.vers,
		cipherSuite:	hs.suite.id,
		createdAt:	createdAt,
		masterSecret:	hs.masterSecret,
		certificates:	certsFromClient,
	}
	stateBytes, err := state.marshal()
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:776
		_go_fuzz_dep_.CoverTab[24257]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:777
		// _ = "end of CoverTab[24257]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:778
		_go_fuzz_dep_.CoverTab[24258]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:778
		// _ = "end of CoverTab[24258]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:778
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:778
	// _ = "end of CoverTab[24248]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:778
	_go_fuzz_dep_.CoverTab[24249]++
								m.ticket, err = c.encryptTicket(stateBytes)
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:780
		_go_fuzz_dep_.CoverTab[24259]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:781
		// _ = "end of CoverTab[24259]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:782
		_go_fuzz_dep_.CoverTab[24260]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:782
		// _ = "end of CoverTab[24260]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:782
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:782
	// _ = "end of CoverTab[24249]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:782
	_go_fuzz_dep_.CoverTab[24250]++

								if _, err := hs.c.writeHandshakeRecord(m, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:784
		_go_fuzz_dep_.CoverTab[24261]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:785
		// _ = "end of CoverTab[24261]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:786
		_go_fuzz_dep_.CoverTab[24262]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:786
		// _ = "end of CoverTab[24262]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:786
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:786
	// _ = "end of CoverTab[24250]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:786
	_go_fuzz_dep_.CoverTab[24251]++

								return nil
//line /usr/local/go/src/crypto/tls/handshake_server.go:788
	// _ = "end of CoverTab[24251]"
}

func (hs *serverHandshakeState) sendFinished(out []byte) error {
//line /usr/local/go/src/crypto/tls/handshake_server.go:791
	_go_fuzz_dep_.CoverTab[24263]++
								c := hs.c

								if err := c.writeChangeCipherRecord(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:794
		_go_fuzz_dep_.CoverTab[24266]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:795
		// _ = "end of CoverTab[24266]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:796
		_go_fuzz_dep_.CoverTab[24267]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:796
		// _ = "end of CoverTab[24267]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:796
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:796
	// _ = "end of CoverTab[24263]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:796
	_go_fuzz_dep_.CoverTab[24264]++

								finished := new(finishedMsg)
								finished.verifyData = hs.finishedHash.serverSum(hs.masterSecret)
								if _, err := hs.c.writeHandshakeRecord(finished, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:800
		_go_fuzz_dep_.CoverTab[24268]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:801
		// _ = "end of CoverTab[24268]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:802
		_go_fuzz_dep_.CoverTab[24269]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:802
		// _ = "end of CoverTab[24269]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:802
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:802
	// _ = "end of CoverTab[24264]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:802
	_go_fuzz_dep_.CoverTab[24265]++

								copy(out, finished.verifyData)

								return nil
//line /usr/local/go/src/crypto/tls/handshake_server.go:806
	// _ = "end of CoverTab[24265]"
}

// processCertsFromClient takes a chain of client certificates either from a
//line /usr/local/go/src/crypto/tls/handshake_server.go:809
// Certificates message or from a sessionState and verifies them. It returns
//line /usr/local/go/src/crypto/tls/handshake_server.go:809
// the public key of the leaf certificate.
//line /usr/local/go/src/crypto/tls/handshake_server.go:812
func (c *Conn) processCertsFromClient(certificate Certificate) error {
//line /usr/local/go/src/crypto/tls/handshake_server.go:812
	_go_fuzz_dep_.CoverTab[24270]++
								certificates := certificate.Certificate
								certs := make([]*x509.Certificate, len(certificates))
								var err error
								for i, asn1Data := range certificates {
//line /usr/local/go/src/crypto/tls/handshake_server.go:816
		_go_fuzz_dep_.CoverTab[24276]++
									if certs[i], err = x509.ParseCertificate(asn1Data); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:817
			_go_fuzz_dep_.CoverTab[24278]++
										c.sendAlert(alertBadCertificate)
										return errors.New("tls: failed to parse client certificate: " + err.Error())
//line /usr/local/go/src/crypto/tls/handshake_server.go:819
			// _ = "end of CoverTab[24278]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:820
			_go_fuzz_dep_.CoverTab[24279]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:820
			// _ = "end of CoverTab[24279]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:820
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:820
		// _ = "end of CoverTab[24276]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:820
		_go_fuzz_dep_.CoverTab[24277]++
									if certs[i].PublicKeyAlgorithm == x509.RSA && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:821
			_go_fuzz_dep_.CoverTab[24280]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:821
			return certs[i].PublicKey.(*rsa.PublicKey).N.BitLen() > maxRSAKeySize
//line /usr/local/go/src/crypto/tls/handshake_server.go:821
			// _ = "end of CoverTab[24280]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:821
		}() {
//line /usr/local/go/src/crypto/tls/handshake_server.go:821
			_go_fuzz_dep_.CoverTab[24281]++
										c.sendAlert(alertBadCertificate)
										return fmt.Errorf("tls: client sent certificate containing RSA key larger than %d bits", maxRSAKeySize)
//line /usr/local/go/src/crypto/tls/handshake_server.go:823
			// _ = "end of CoverTab[24281]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:824
			_go_fuzz_dep_.CoverTab[24282]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:824
			// _ = "end of CoverTab[24282]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:824
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:824
		// _ = "end of CoverTab[24277]"
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:825
	// _ = "end of CoverTab[24270]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:825
	_go_fuzz_dep_.CoverTab[24271]++

								if len(certs) == 0 && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:827
		_go_fuzz_dep_.CoverTab[24283]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:827
		return requiresClientCert(c.config.ClientAuth)
//line /usr/local/go/src/crypto/tls/handshake_server.go:827
		// _ = "end of CoverTab[24283]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:827
	}() {
//line /usr/local/go/src/crypto/tls/handshake_server.go:827
		_go_fuzz_dep_.CoverTab[24284]++
									c.sendAlert(alertBadCertificate)
									return errors.New("tls: client didn't provide a certificate")
//line /usr/local/go/src/crypto/tls/handshake_server.go:829
		// _ = "end of CoverTab[24284]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:830
		_go_fuzz_dep_.CoverTab[24285]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:830
		// _ = "end of CoverTab[24285]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:830
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:830
	// _ = "end of CoverTab[24271]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:830
	_go_fuzz_dep_.CoverTab[24272]++

								if c.config.ClientAuth >= VerifyClientCertIfGiven && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_server.go:832
		_go_fuzz_dep_.CoverTab[24286]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:832
		return len(certs) > 0
//line /usr/local/go/src/crypto/tls/handshake_server.go:832
		// _ = "end of CoverTab[24286]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:832
	}() {
//line /usr/local/go/src/crypto/tls/handshake_server.go:832
		_go_fuzz_dep_.CoverTab[24287]++
									opts := x509.VerifyOptions{
			Roots:		c.config.ClientCAs,
			CurrentTime:	c.config.time(),
			Intermediates:	x509.NewCertPool(),
			KeyUsages:	[]x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		}

		for _, cert := range certs[1:] {
//line /usr/local/go/src/crypto/tls/handshake_server.go:840
			_go_fuzz_dep_.CoverTab[24290]++
										opts.Intermediates.AddCert(cert)
//line /usr/local/go/src/crypto/tls/handshake_server.go:841
			// _ = "end of CoverTab[24290]"
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:842
		// _ = "end of CoverTab[24287]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:842
		_go_fuzz_dep_.CoverTab[24288]++

									chains, err := certs[0].Verify(opts)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:845
			_go_fuzz_dep_.CoverTab[24291]++
										c.sendAlert(alertBadCertificate)
										return &CertificateVerificationError{UnverifiedCertificates: certs, Err: err}
//line /usr/local/go/src/crypto/tls/handshake_server.go:847
			// _ = "end of CoverTab[24291]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:848
			_go_fuzz_dep_.CoverTab[24292]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:848
			// _ = "end of CoverTab[24292]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:848
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:848
		// _ = "end of CoverTab[24288]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:848
		_go_fuzz_dep_.CoverTab[24289]++

									c.verifiedChains = chains
//line /usr/local/go/src/crypto/tls/handshake_server.go:850
		// _ = "end of CoverTab[24289]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:851
		_go_fuzz_dep_.CoverTab[24293]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:851
		// _ = "end of CoverTab[24293]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:851
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:851
	// _ = "end of CoverTab[24272]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:851
	_go_fuzz_dep_.CoverTab[24273]++

								c.peerCertificates = certs
								c.ocspResponse = certificate.OCSPStaple
								c.scts = certificate.SignedCertificateTimestamps

								if len(certs) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_server.go:857
		_go_fuzz_dep_.CoverTab[24294]++
									switch certs[0].PublicKey.(type) {
		case *ecdsa.PublicKey, *rsa.PublicKey, ed25519.PublicKey:
//line /usr/local/go/src/crypto/tls/handshake_server.go:859
			_go_fuzz_dep_.CoverTab[24295]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:859
			// _ = "end of CoverTab[24295]"
		default:
//line /usr/local/go/src/crypto/tls/handshake_server.go:860
			_go_fuzz_dep_.CoverTab[24296]++
										c.sendAlert(alertUnsupportedCertificate)
										return fmt.Errorf("tls: client certificate contains an unsupported public key of type %T", certs[0].PublicKey)
//line /usr/local/go/src/crypto/tls/handshake_server.go:862
			// _ = "end of CoverTab[24296]"
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:863
		// _ = "end of CoverTab[24294]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:864
		_go_fuzz_dep_.CoverTab[24297]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:864
		// _ = "end of CoverTab[24297]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:864
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:864
	// _ = "end of CoverTab[24273]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:864
	_go_fuzz_dep_.CoverTab[24274]++

								if c.config.VerifyPeerCertificate != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:866
		_go_fuzz_dep_.CoverTab[24298]++
									if err := c.config.VerifyPeerCertificate(certificates, c.verifiedChains); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_server.go:867
			_go_fuzz_dep_.CoverTab[24299]++
										c.sendAlert(alertBadCertificate)
										return err
//line /usr/local/go/src/crypto/tls/handshake_server.go:869
			// _ = "end of CoverTab[24299]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:870
			_go_fuzz_dep_.CoverTab[24300]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:870
			// _ = "end of CoverTab[24300]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:870
		}
//line /usr/local/go/src/crypto/tls/handshake_server.go:870
		// _ = "end of CoverTab[24298]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:871
		_go_fuzz_dep_.CoverTab[24301]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:871
		// _ = "end of CoverTab[24301]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:871
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:871
	// _ = "end of CoverTab[24274]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:871
	_go_fuzz_dep_.CoverTab[24275]++

								return nil
//line /usr/local/go/src/crypto/tls/handshake_server.go:873
	// _ = "end of CoverTab[24275]"
}

func clientHelloInfo(ctx context.Context, c *Conn, clientHello *clientHelloMsg) *ClientHelloInfo {
//line /usr/local/go/src/crypto/tls/handshake_server.go:876
	_go_fuzz_dep_.CoverTab[24302]++
								supportedVersions := clientHello.supportedVersions
								if len(clientHello.supportedVersions) == 0 {
//line /usr/local/go/src/crypto/tls/handshake_server.go:878
		_go_fuzz_dep_.CoverTab[24304]++
									supportedVersions = supportedVersionsFromMax(clientHello.vers)
//line /usr/local/go/src/crypto/tls/handshake_server.go:879
		// _ = "end of CoverTab[24304]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_server.go:880
		_go_fuzz_dep_.CoverTab[24305]++
//line /usr/local/go/src/crypto/tls/handshake_server.go:880
		// _ = "end of CoverTab[24305]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:880
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:880
	// _ = "end of CoverTab[24302]"
//line /usr/local/go/src/crypto/tls/handshake_server.go:880
	_go_fuzz_dep_.CoverTab[24303]++

								return &ClientHelloInfo{
		CipherSuites:		clientHello.cipherSuites,
		ServerName:		clientHello.serverName,
		SupportedCurves:	clientHello.supportedCurves,
		SupportedPoints:	clientHello.supportedPoints,
		SignatureSchemes:	clientHello.supportedSignatureAlgorithms,
		SupportedProtos:	clientHello.alpnProtocols,
		SupportedVersions:	supportedVersions,
		Conn:			c.conn,
		config:			c.config,
		ctx:			ctx,
	}
//line /usr/local/go/src/crypto/tls/handshake_server.go:893
	// _ = "end of CoverTab[24303]"
}

//line /usr/local/go/src/crypto/tls/handshake_server.go:894
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/handshake_server.go:894
var _ = _go_fuzz_dep_.CoverTab
