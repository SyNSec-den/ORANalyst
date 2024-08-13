// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/tls/handshake_client.go:5
package tls

//line /usr/local/go/src/crypto/tls/handshake_client.go:5
import (
//line /usr/local/go/src/crypto/tls/handshake_client.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/handshake_client.go:5
)
//line /usr/local/go/src/crypto/tls/handshake_client.go:5
import (
//line /usr/local/go/src/crypto/tls/handshake_client.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/handshake_client.go:5
)

import (
	"bytes"
	"context"
	"crypto"
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/subtle"
	"crypto/x509"
	"errors"
	"fmt"
	"hash"
	"io"
	"net"
	"strings"
	"time"
)

type clientHandshakeState struct {
	c		*Conn
	ctx		context.Context
	serverHello	*serverHelloMsg
	hello		*clientHelloMsg
	suite		*cipherSuite
	finishedHash	finishedHash
	masterSecret	[]byte
	session		*ClientSessionState
}

var testingOnlyForceClientHelloSignatureAlgorithms []SignatureScheme

func (c *Conn) makeClientHello() (*clientHelloMsg, *ecdh.PrivateKey, error) {
//line /usr/local/go/src/crypto/tls/handshake_client.go:39
	_go_fuzz_dep_.CoverTab[22181]++
								config := c.config
								if len(config.ServerName) == 0 && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:41
		_go_fuzz_dep_.CoverTab[22195]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:41
		return !config.InsecureSkipVerify
//line /usr/local/go/src/crypto/tls/handshake_client.go:41
		// _ = "end of CoverTab[22195]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:41
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:41
		_go_fuzz_dep_.CoverTab[22196]++
									return nil, nil, errors.New("tls: either ServerName or InsecureSkipVerify must be specified in the tls.Config")
//line /usr/local/go/src/crypto/tls/handshake_client.go:42
		// _ = "end of CoverTab[22196]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:43
		_go_fuzz_dep_.CoverTab[22197]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:43
		// _ = "end of CoverTab[22197]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:43
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:43
	// _ = "end of CoverTab[22181]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:43
	_go_fuzz_dep_.CoverTab[22182]++

								nextProtosLength := 0
								for _, proto := range config.NextProtos {
//line /usr/local/go/src/crypto/tls/handshake_client.go:46
		_go_fuzz_dep_.CoverTab[22198]++
									if l := len(proto); l == 0 || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:47
			_go_fuzz_dep_.CoverTab[22199]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:47
			return l > 255
//line /usr/local/go/src/crypto/tls/handshake_client.go:47
			// _ = "end of CoverTab[22199]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:47
		}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:47
			_go_fuzz_dep_.CoverTab[22200]++
										return nil, nil, errors.New("tls: invalid NextProtos value")
//line /usr/local/go/src/crypto/tls/handshake_client.go:48
			// _ = "end of CoverTab[22200]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:49
			_go_fuzz_dep_.CoverTab[22201]++
										nextProtosLength += 1 + l
//line /usr/local/go/src/crypto/tls/handshake_client.go:50
			// _ = "end of CoverTab[22201]"
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:51
		// _ = "end of CoverTab[22198]"
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:52
	// _ = "end of CoverTab[22182]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:52
	_go_fuzz_dep_.CoverTab[22183]++
								if nextProtosLength > 0xffff {
//line /usr/local/go/src/crypto/tls/handshake_client.go:53
		_go_fuzz_dep_.CoverTab[22202]++
									return nil, nil, errors.New("tls: NextProtos values too large")
//line /usr/local/go/src/crypto/tls/handshake_client.go:54
		// _ = "end of CoverTab[22202]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:55
		_go_fuzz_dep_.CoverTab[22203]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:55
		// _ = "end of CoverTab[22203]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:55
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:55
	// _ = "end of CoverTab[22183]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:55
	_go_fuzz_dep_.CoverTab[22184]++

								supportedVersions := config.supportedVersions(roleClient)
								if len(supportedVersions) == 0 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:58
		_go_fuzz_dep_.CoverTab[22204]++
									return nil, nil, errors.New("tls: no supported versions satisfy MinVersion and MaxVersion")
//line /usr/local/go/src/crypto/tls/handshake_client.go:59
		// _ = "end of CoverTab[22204]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:60
		_go_fuzz_dep_.CoverTab[22205]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:60
		// _ = "end of CoverTab[22205]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:60
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:60
	// _ = "end of CoverTab[22184]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:60
	_go_fuzz_dep_.CoverTab[22185]++

								clientHelloVersion := config.maxSupportedVersion(roleClient)

//line /usr/local/go/src/crypto/tls/handshake_client.go:66
	if clientHelloVersion > VersionTLS12 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:66
		_go_fuzz_dep_.CoverTab[22206]++
									clientHelloVersion = VersionTLS12
//line /usr/local/go/src/crypto/tls/handshake_client.go:67
		// _ = "end of CoverTab[22206]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:68
		_go_fuzz_dep_.CoverTab[22207]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:68
		// _ = "end of CoverTab[22207]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:68
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:68
	// _ = "end of CoverTab[22185]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:68
	_go_fuzz_dep_.CoverTab[22186]++

								hello := &clientHelloMsg{
		vers:				clientHelloVersion,
		compressionMethods:		[]uint8{compressionNone},
		random:				make([]byte, 32),
		sessionId:			make([]byte, 32),
		ocspStapling:			true,
		scts:				true,
		serverName:			hostnameInSNI(config.ServerName),
		supportedCurves:		config.curvePreferences(),
		supportedPoints:		[]uint8{pointFormatUncompressed},
		secureRenegotiationSupported:	true,
		alpnProtocols:			config.NextProtos,
		supportedVersions:		supportedVersions,
	}

	if c.handshakes > 0 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:85
		_go_fuzz_dep_.CoverTab[22208]++
									hello.secureRenegotiation = c.clientFinished[:]
//line /usr/local/go/src/crypto/tls/handshake_client.go:86
		// _ = "end of CoverTab[22208]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:87
		_go_fuzz_dep_.CoverTab[22209]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:87
		// _ = "end of CoverTab[22209]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:87
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:87
	// _ = "end of CoverTab[22186]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:87
	_go_fuzz_dep_.CoverTab[22187]++

								preferenceOrder := cipherSuitesPreferenceOrder
								if !hasAESGCMHardwareSupport {
//line /usr/local/go/src/crypto/tls/handshake_client.go:90
		_go_fuzz_dep_.CoverTab[22210]++
									preferenceOrder = cipherSuitesPreferenceOrderNoAES
//line /usr/local/go/src/crypto/tls/handshake_client.go:91
		// _ = "end of CoverTab[22210]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:92
		_go_fuzz_dep_.CoverTab[22211]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:92
		// _ = "end of CoverTab[22211]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:92
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:92
	// _ = "end of CoverTab[22187]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:92
	_go_fuzz_dep_.CoverTab[22188]++
								configCipherSuites := config.cipherSuites()
								hello.cipherSuites = make([]uint16, 0, len(configCipherSuites))

								for _, suiteId := range preferenceOrder {
//line /usr/local/go/src/crypto/tls/handshake_client.go:96
		_go_fuzz_dep_.CoverTab[22212]++
									suite := mutualCipherSuite(configCipherSuites, suiteId)
									if suite == nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:98
			_go_fuzz_dep_.CoverTab[22215]++
										continue
//line /usr/local/go/src/crypto/tls/handshake_client.go:99
			// _ = "end of CoverTab[22215]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:100
			_go_fuzz_dep_.CoverTab[22216]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:100
			// _ = "end of CoverTab[22216]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:100
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:100
		// _ = "end of CoverTab[22212]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:100
		_go_fuzz_dep_.CoverTab[22213]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:103
		if hello.vers < VersionTLS12 && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:103
			_go_fuzz_dep_.CoverTab[22217]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:103
			return suite.flags&suiteTLS12 != 0
//line /usr/local/go/src/crypto/tls/handshake_client.go:103
			// _ = "end of CoverTab[22217]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:103
		}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:103
			_go_fuzz_dep_.CoverTab[22218]++
										continue
//line /usr/local/go/src/crypto/tls/handshake_client.go:104
			// _ = "end of CoverTab[22218]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:105
			_go_fuzz_dep_.CoverTab[22219]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:105
			// _ = "end of CoverTab[22219]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:105
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:105
		// _ = "end of CoverTab[22213]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:105
		_go_fuzz_dep_.CoverTab[22214]++
									hello.cipherSuites = append(hello.cipherSuites, suiteId)
//line /usr/local/go/src/crypto/tls/handshake_client.go:106
		// _ = "end of CoverTab[22214]"
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:107
	// _ = "end of CoverTab[22188]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:107
	_go_fuzz_dep_.CoverTab[22189]++

								_, err := io.ReadFull(config.rand(), hello.random)
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:110
		_go_fuzz_dep_.CoverTab[22220]++
									return nil, nil, errors.New("tls: short read from Rand: " + err.Error())
//line /usr/local/go/src/crypto/tls/handshake_client.go:111
		// _ = "end of CoverTab[22220]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:112
		_go_fuzz_dep_.CoverTab[22221]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:112
		// _ = "end of CoverTab[22221]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:112
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:112
	// _ = "end of CoverTab[22189]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:112
	_go_fuzz_dep_.CoverTab[22190]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:117
	if _, err := io.ReadFull(config.rand(), hello.sessionId); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:117
		_go_fuzz_dep_.CoverTab[22222]++
									return nil, nil, errors.New("tls: short read from Rand: " + err.Error())
//line /usr/local/go/src/crypto/tls/handshake_client.go:118
		// _ = "end of CoverTab[22222]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:119
		_go_fuzz_dep_.CoverTab[22223]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:119
		// _ = "end of CoverTab[22223]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:119
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:119
	// _ = "end of CoverTab[22190]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:119
	_go_fuzz_dep_.CoverTab[22191]++

								if hello.vers >= VersionTLS12 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:121
		_go_fuzz_dep_.CoverTab[22224]++
									hello.supportedSignatureAlgorithms = supportedSignatureAlgorithms()
//line /usr/local/go/src/crypto/tls/handshake_client.go:122
		// _ = "end of CoverTab[22224]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:123
		_go_fuzz_dep_.CoverTab[22225]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:123
		// _ = "end of CoverTab[22225]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:123
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:123
	// _ = "end of CoverTab[22191]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:123
	_go_fuzz_dep_.CoverTab[22192]++
								if testingOnlyForceClientHelloSignatureAlgorithms != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:124
		_go_fuzz_dep_.CoverTab[22226]++
									hello.supportedSignatureAlgorithms = testingOnlyForceClientHelloSignatureAlgorithms
//line /usr/local/go/src/crypto/tls/handshake_client.go:125
		// _ = "end of CoverTab[22226]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:126
		_go_fuzz_dep_.CoverTab[22227]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:126
		// _ = "end of CoverTab[22227]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:126
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:126
	// _ = "end of CoverTab[22192]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:126
	_go_fuzz_dep_.CoverTab[22193]++

								var key *ecdh.PrivateKey
								if hello.supportedVersions[0] == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:129
		_go_fuzz_dep_.CoverTab[22228]++
									if hasAESGCMHardwareSupport {
//line /usr/local/go/src/crypto/tls/handshake_client.go:130
			_go_fuzz_dep_.CoverTab[22232]++
										hello.cipherSuites = append(hello.cipherSuites, defaultCipherSuitesTLS13...)
//line /usr/local/go/src/crypto/tls/handshake_client.go:131
			// _ = "end of CoverTab[22232]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:132
			_go_fuzz_dep_.CoverTab[22233]++
										hello.cipherSuites = append(hello.cipherSuites, defaultCipherSuitesTLS13NoAES...)
//line /usr/local/go/src/crypto/tls/handshake_client.go:133
			// _ = "end of CoverTab[22233]"
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:134
		// _ = "end of CoverTab[22228]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:134
		_go_fuzz_dep_.CoverTab[22229]++

									curveID := config.curvePreferences()[0]
									if _, ok := curveForCurveID(curveID); !ok {
//line /usr/local/go/src/crypto/tls/handshake_client.go:137
			_go_fuzz_dep_.CoverTab[22234]++
										return nil, nil, errors.New("tls: CurvePreferences includes unsupported curve")
//line /usr/local/go/src/crypto/tls/handshake_client.go:138
			// _ = "end of CoverTab[22234]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:139
			_go_fuzz_dep_.CoverTab[22235]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:139
			// _ = "end of CoverTab[22235]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:139
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:139
		// _ = "end of CoverTab[22229]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:139
		_go_fuzz_dep_.CoverTab[22230]++
									key, err = generateECDHEKey(config.rand(), curveID)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:141
			_go_fuzz_dep_.CoverTab[22236]++
										return nil, nil, err
//line /usr/local/go/src/crypto/tls/handshake_client.go:142
			// _ = "end of CoverTab[22236]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:143
			_go_fuzz_dep_.CoverTab[22237]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:143
			// _ = "end of CoverTab[22237]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:143
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:143
		// _ = "end of CoverTab[22230]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:143
		_go_fuzz_dep_.CoverTab[22231]++
									hello.keyShares = []keyShare{{group: curveID, data: key.PublicKey().Bytes()}}
//line /usr/local/go/src/crypto/tls/handshake_client.go:144
		// _ = "end of CoverTab[22231]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:145
		_go_fuzz_dep_.CoverTab[22238]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:145
		// _ = "end of CoverTab[22238]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:145
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:145
	// _ = "end of CoverTab[22193]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:145
	_go_fuzz_dep_.CoverTab[22194]++

								return hello, key, nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:147
	// _ = "end of CoverTab[22194]"
}

func (c *Conn) clientHandshake(ctx context.Context) (err error) {
//line /usr/local/go/src/crypto/tls/handshake_client.go:150
	_go_fuzz_dep_.CoverTab[22239]++
								if c.config == nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:151
		_go_fuzz_dep_.CoverTab[22252]++
									c.config = defaultConfig()
//line /usr/local/go/src/crypto/tls/handshake_client.go:152
		// _ = "end of CoverTab[22252]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:153
		_go_fuzz_dep_.CoverTab[22253]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:153
		// _ = "end of CoverTab[22253]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:153
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:153
	// _ = "end of CoverTab[22239]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:153
	_go_fuzz_dep_.CoverTab[22240]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:157
	c.didResume = false

	hello, ecdheKey, err := c.makeClientHello()
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:160
		_go_fuzz_dep_.CoverTab[22254]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:161
		// _ = "end of CoverTab[22254]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:162
		_go_fuzz_dep_.CoverTab[22255]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:162
		// _ = "end of CoverTab[22255]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:162
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:162
	// _ = "end of CoverTab[22240]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:162
	_go_fuzz_dep_.CoverTab[22241]++
								c.serverName = hello.serverName

								cacheKey, session, earlySecret, binderKey, err := c.loadSession(hello)
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:166
		_go_fuzz_dep_.CoverTab[22256]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:167
		// _ = "end of CoverTab[22256]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:168
		_go_fuzz_dep_.CoverTab[22257]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:168
		// _ = "end of CoverTab[22257]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:168
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:168
	// _ = "end of CoverTab[22241]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:168
	_go_fuzz_dep_.CoverTab[22242]++
								if cacheKey != "" && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:169
		_go_fuzz_dep_.CoverTab[22258]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:169
		return session != nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:169
		// _ = "end of CoverTab[22258]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:169
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:169
		_go_fuzz_dep_.CoverTab[22259]++
									defer func() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:170
			_go_fuzz_dep_.CoverTab[22260]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:177
			if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:177
				_go_fuzz_dep_.CoverTab[22261]++
											c.config.ClientSessionCache.Put(cacheKey, nil)
//line /usr/local/go/src/crypto/tls/handshake_client.go:178
				// _ = "end of CoverTab[22261]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:179
				_go_fuzz_dep_.CoverTab[22262]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:179
				// _ = "end of CoverTab[22262]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:179
			}
//line /usr/local/go/src/crypto/tls/handshake_client.go:179
			// _ = "end of CoverTab[22260]"
		}()
//line /usr/local/go/src/crypto/tls/handshake_client.go:180
		// _ = "end of CoverTab[22259]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:181
		_go_fuzz_dep_.CoverTab[22263]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:181
		// _ = "end of CoverTab[22263]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:181
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:181
	// _ = "end of CoverTab[22242]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:181
	_go_fuzz_dep_.CoverTab[22243]++

								if _, err := c.writeHandshakeRecord(hello, nil); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:183
		_go_fuzz_dep_.CoverTab[22264]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:184
		// _ = "end of CoverTab[22264]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:185
		_go_fuzz_dep_.CoverTab[22265]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:185
		// _ = "end of CoverTab[22265]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:185
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:185
	// _ = "end of CoverTab[22243]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:185
	_go_fuzz_dep_.CoverTab[22244]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:188
	msg, err := c.readHandshake(nil)
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:189
		_go_fuzz_dep_.CoverTab[22266]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:190
		// _ = "end of CoverTab[22266]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:191
		_go_fuzz_dep_.CoverTab[22267]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:191
		// _ = "end of CoverTab[22267]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:191
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:191
	// _ = "end of CoverTab[22244]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:191
	_go_fuzz_dep_.CoverTab[22245]++

								serverHello, ok := msg.(*serverHelloMsg)
								if !ok {
//line /usr/local/go/src/crypto/tls/handshake_client.go:194
		_go_fuzz_dep_.CoverTab[22268]++
									c.sendAlert(alertUnexpectedMessage)
									return unexpectedMessageError(serverHello, msg)
//line /usr/local/go/src/crypto/tls/handshake_client.go:196
		// _ = "end of CoverTab[22268]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:197
		_go_fuzz_dep_.CoverTab[22269]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:197
		// _ = "end of CoverTab[22269]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:197
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:197
	// _ = "end of CoverTab[22245]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:197
	_go_fuzz_dep_.CoverTab[22246]++

								if err := c.pickTLSVersion(serverHello); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:199
		_go_fuzz_dep_.CoverTab[22270]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:200
		// _ = "end of CoverTab[22270]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:201
		_go_fuzz_dep_.CoverTab[22271]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:201
		// _ = "end of CoverTab[22271]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:201
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:201
	// _ = "end of CoverTab[22246]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:201
	_go_fuzz_dep_.CoverTab[22247]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:206
	maxVers := c.config.maxSupportedVersion(roleClient)
	tls12Downgrade := string(serverHello.random[24:]) == downgradeCanaryTLS12
	tls11Downgrade := string(serverHello.random[24:]) == downgradeCanaryTLS11
	if maxVers == VersionTLS13 && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:209
		_go_fuzz_dep_.CoverTab[22272]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:209
		return c.vers <= VersionTLS12
//line /usr/local/go/src/crypto/tls/handshake_client.go:209
		// _ = "end of CoverTab[22272]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:209
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:209
		_go_fuzz_dep_.CoverTab[22273]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:209
		return (tls12Downgrade || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:209
			_go_fuzz_dep_.CoverTab[22274]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:209
			return tls11Downgrade
//line /usr/local/go/src/crypto/tls/handshake_client.go:209
			// _ = "end of CoverTab[22274]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:209
		}())
//line /usr/local/go/src/crypto/tls/handshake_client.go:209
		// _ = "end of CoverTab[22273]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:209
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:209
		_go_fuzz_dep_.CoverTab[22275]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:209
		return maxVers == VersionTLS12 && func() bool {
										_go_fuzz_dep_.CoverTab[22276]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:210
			return c.vers <= VersionTLS11
//line /usr/local/go/src/crypto/tls/handshake_client.go:210
			// _ = "end of CoverTab[22276]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:210
		}() && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:210
			_go_fuzz_dep_.CoverTab[22277]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:210
			return tls11Downgrade
//line /usr/local/go/src/crypto/tls/handshake_client.go:210
			// _ = "end of CoverTab[22277]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:210
		}()
//line /usr/local/go/src/crypto/tls/handshake_client.go:210
		// _ = "end of CoverTab[22275]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:210
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:210
		_go_fuzz_dep_.CoverTab[22278]++
									c.sendAlert(alertIllegalParameter)
									return errors.New("tls: downgrade attempt detected, possibly due to a MitM attack or a broken middlebox")
//line /usr/local/go/src/crypto/tls/handshake_client.go:212
		// _ = "end of CoverTab[22278]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:213
		_go_fuzz_dep_.CoverTab[22279]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:213
		// _ = "end of CoverTab[22279]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:213
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:213
	// _ = "end of CoverTab[22247]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:213
	_go_fuzz_dep_.CoverTab[22248]++

								if c.vers == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:215
		_go_fuzz_dep_.CoverTab[22280]++
									hs := &clientHandshakeStateTLS13{
			c:		c,
			ctx:		ctx,
			serverHello:	serverHello,
			hello:		hello,
			ecdheKey:	ecdheKey,
			session:	session,
			earlySecret:	earlySecret,
			binderKey:	binderKey,
		}

//line /usr/local/go/src/crypto/tls/handshake_client.go:228
		return hs.handshake()
//line /usr/local/go/src/crypto/tls/handshake_client.go:228
		// _ = "end of CoverTab[22280]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:229
		_go_fuzz_dep_.CoverTab[22281]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:229
		// _ = "end of CoverTab[22281]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:229
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:229
	// _ = "end of CoverTab[22248]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:229
	_go_fuzz_dep_.CoverTab[22249]++

								hs := &clientHandshakeState{
		c:		c,
		ctx:		ctx,
		serverHello:	serverHello,
		hello:		hello,
		session:	session,
	}

	if err := hs.handshake(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:239
		_go_fuzz_dep_.CoverTab[22282]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:240
		// _ = "end of CoverTab[22282]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:241
		_go_fuzz_dep_.CoverTab[22283]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:241
		// _ = "end of CoverTab[22283]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:241
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:241
	// _ = "end of CoverTab[22249]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:241
	_go_fuzz_dep_.CoverTab[22250]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:245
	if cacheKey != "" && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:245
		_go_fuzz_dep_.CoverTab[22284]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:245
		return hs.session != nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:245
		// _ = "end of CoverTab[22284]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:245
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:245
		_go_fuzz_dep_.CoverTab[22285]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:245
		return session != hs.session
//line /usr/local/go/src/crypto/tls/handshake_client.go:245
		// _ = "end of CoverTab[22285]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:245
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:245
		_go_fuzz_dep_.CoverTab[22286]++
									c.config.ClientSessionCache.Put(cacheKey, hs.session)
//line /usr/local/go/src/crypto/tls/handshake_client.go:246
		// _ = "end of CoverTab[22286]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:247
		_go_fuzz_dep_.CoverTab[22287]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:247
		// _ = "end of CoverTab[22287]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:247
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:247
	// _ = "end of CoverTab[22250]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:247
	_go_fuzz_dep_.CoverTab[22251]++

								return nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:249
	// _ = "end of CoverTab[22251]"
}

func (c *Conn) loadSession(hello *clientHelloMsg) (cacheKey string,
	session *ClientSessionState, earlySecret, binderKey []byte, err error) {
//line /usr/local/go/src/crypto/tls/handshake_client.go:253
	_go_fuzz_dep_.CoverTab[22288]++
								if c.config.SessionTicketsDisabled || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:254
		_go_fuzz_dep_.CoverTab[22303]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:254
		return c.config.ClientSessionCache == nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:254
		// _ = "end of CoverTab[22303]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:254
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:254
		_go_fuzz_dep_.CoverTab[22304]++
									return "", nil, nil, nil, nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:255
		// _ = "end of CoverTab[22304]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:256
		_go_fuzz_dep_.CoverTab[22305]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:256
		// _ = "end of CoverTab[22305]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:256
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:256
	// _ = "end of CoverTab[22288]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:256
	_go_fuzz_dep_.CoverTab[22289]++

								hello.ticketSupported = true

								if hello.supportedVersions[0] == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:260
		_go_fuzz_dep_.CoverTab[22306]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:263
		hello.pskModes = []uint8{pskModeDHE}
//line /usr/local/go/src/crypto/tls/handshake_client.go:263
		// _ = "end of CoverTab[22306]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:264
		_go_fuzz_dep_.CoverTab[22307]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:264
		// _ = "end of CoverTab[22307]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:264
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:264
	// _ = "end of CoverTab[22289]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:264
	_go_fuzz_dep_.CoverTab[22290]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:269
	if c.handshakes != 0 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:269
		_go_fuzz_dep_.CoverTab[22308]++
									return "", nil, nil, nil, nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:270
		// _ = "end of CoverTab[22308]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:271
		_go_fuzz_dep_.CoverTab[22309]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:271
		// _ = "end of CoverTab[22309]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:271
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:271
	// _ = "end of CoverTab[22290]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:271
	_go_fuzz_dep_.CoverTab[22291]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:274
	cacheKey = clientSessionCacheKey(c.conn.RemoteAddr(), c.config)
	session, ok := c.config.ClientSessionCache.Get(cacheKey)
	if !ok || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:276
		_go_fuzz_dep_.CoverTab[22310]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:276
		return session == nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:276
		// _ = "end of CoverTab[22310]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:276
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:276
		_go_fuzz_dep_.CoverTab[22311]++
									return cacheKey, nil, nil, nil, nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:277
		// _ = "end of CoverTab[22311]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:278
		_go_fuzz_dep_.CoverTab[22312]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:278
		// _ = "end of CoverTab[22312]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:278
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:278
	// _ = "end of CoverTab[22291]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:278
	_go_fuzz_dep_.CoverTab[22292]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:281
	versOk := false
	for _, v := range hello.supportedVersions {
//line /usr/local/go/src/crypto/tls/handshake_client.go:282
		_go_fuzz_dep_.CoverTab[22313]++
									if v == session.vers {
//line /usr/local/go/src/crypto/tls/handshake_client.go:283
			_go_fuzz_dep_.CoverTab[22314]++
										versOk = true
										break
//line /usr/local/go/src/crypto/tls/handshake_client.go:285
			// _ = "end of CoverTab[22314]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:286
			_go_fuzz_dep_.CoverTab[22315]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:286
			// _ = "end of CoverTab[22315]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:286
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:286
		// _ = "end of CoverTab[22313]"
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:287
	// _ = "end of CoverTab[22292]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:287
	_go_fuzz_dep_.CoverTab[22293]++
								if !versOk {
//line /usr/local/go/src/crypto/tls/handshake_client.go:288
		_go_fuzz_dep_.CoverTab[22316]++
									return cacheKey, nil, nil, nil, nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:289
		// _ = "end of CoverTab[22316]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:290
		_go_fuzz_dep_.CoverTab[22317]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:290
		// _ = "end of CoverTab[22317]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:290
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:290
	// _ = "end of CoverTab[22293]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:290
	_go_fuzz_dep_.CoverTab[22294]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:295
	if !c.config.InsecureSkipVerify {
//line /usr/local/go/src/crypto/tls/handshake_client.go:295
		_go_fuzz_dep_.CoverTab[22318]++
									if len(session.verifiedChains) == 0 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:296
			_go_fuzz_dep_.CoverTab[22321]++

										return cacheKey, nil, nil, nil, nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:298
			// _ = "end of CoverTab[22321]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:299
			_go_fuzz_dep_.CoverTab[22322]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:299
			// _ = "end of CoverTab[22322]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:299
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:299
		// _ = "end of CoverTab[22318]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:299
		_go_fuzz_dep_.CoverTab[22319]++
									serverCert := session.serverCertificates[0]
									if c.config.time().After(serverCert.NotAfter) {
//line /usr/local/go/src/crypto/tls/handshake_client.go:301
			_go_fuzz_dep_.CoverTab[22323]++

										c.config.ClientSessionCache.Put(cacheKey, nil)
										return cacheKey, nil, nil, nil, nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:304
			// _ = "end of CoverTab[22323]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:305
			_go_fuzz_dep_.CoverTab[22324]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:305
			// _ = "end of CoverTab[22324]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:305
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:305
		// _ = "end of CoverTab[22319]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:305
		_go_fuzz_dep_.CoverTab[22320]++
									if err := serverCert.VerifyHostname(c.config.ServerName); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:306
			_go_fuzz_dep_.CoverTab[22325]++
										return cacheKey, nil, nil, nil, nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:307
			// _ = "end of CoverTab[22325]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:308
			_go_fuzz_dep_.CoverTab[22326]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:308
			// _ = "end of CoverTab[22326]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:308
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:308
		// _ = "end of CoverTab[22320]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:309
		_go_fuzz_dep_.CoverTab[22327]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:309
		// _ = "end of CoverTab[22327]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:309
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:309
	// _ = "end of CoverTab[22294]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:309
	_go_fuzz_dep_.CoverTab[22295]++

								if session.vers != VersionTLS13 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:311
		_go_fuzz_dep_.CoverTab[22328]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:314
		if mutualCipherSuite(hello.cipherSuites, session.cipherSuite) == nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:314
			_go_fuzz_dep_.CoverTab[22330]++
										return cacheKey, nil, nil, nil, nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:315
			// _ = "end of CoverTab[22330]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:316
			_go_fuzz_dep_.CoverTab[22331]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:316
			// _ = "end of CoverTab[22331]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:316
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:316
		// _ = "end of CoverTab[22328]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:316
		_go_fuzz_dep_.CoverTab[22329]++

									hello.sessionTicket = session.sessionTicket
									return
//line /usr/local/go/src/crypto/tls/handshake_client.go:319
		// _ = "end of CoverTab[22329]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:320
		_go_fuzz_dep_.CoverTab[22332]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:320
		// _ = "end of CoverTab[22332]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:320
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:320
	// _ = "end of CoverTab[22295]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:320
	_go_fuzz_dep_.CoverTab[22296]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:323
	if c.config.time().After(session.useBy) {
//line /usr/local/go/src/crypto/tls/handshake_client.go:323
		_go_fuzz_dep_.CoverTab[22333]++
									c.config.ClientSessionCache.Put(cacheKey, nil)
									return cacheKey, nil, nil, nil, nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:325
		// _ = "end of CoverTab[22333]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:326
		_go_fuzz_dep_.CoverTab[22334]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:326
		// _ = "end of CoverTab[22334]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:326
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:326
	// _ = "end of CoverTab[22296]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:326
	_go_fuzz_dep_.CoverTab[22297]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:330
	cipherSuite := cipherSuiteTLS13ByID(session.cipherSuite)
	if cipherSuite == nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:331
		_go_fuzz_dep_.CoverTab[22335]++
									return cacheKey, nil, nil, nil, nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:332
		// _ = "end of CoverTab[22335]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:333
		_go_fuzz_dep_.CoverTab[22336]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:333
		// _ = "end of CoverTab[22336]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:333
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:333
	// _ = "end of CoverTab[22297]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:333
	_go_fuzz_dep_.CoverTab[22298]++
								cipherSuiteOk := false
								for _, offeredID := range hello.cipherSuites {
//line /usr/local/go/src/crypto/tls/handshake_client.go:335
		_go_fuzz_dep_.CoverTab[22337]++
									offeredSuite := cipherSuiteTLS13ByID(offeredID)
									if offeredSuite != nil && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:337
			_go_fuzz_dep_.CoverTab[22338]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:337
			return offeredSuite.hash == cipherSuite.hash
//line /usr/local/go/src/crypto/tls/handshake_client.go:337
			// _ = "end of CoverTab[22338]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:337
		}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:337
			_go_fuzz_dep_.CoverTab[22339]++
										cipherSuiteOk = true
										break
//line /usr/local/go/src/crypto/tls/handshake_client.go:339
			// _ = "end of CoverTab[22339]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:340
			_go_fuzz_dep_.CoverTab[22340]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:340
			// _ = "end of CoverTab[22340]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:340
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:340
		// _ = "end of CoverTab[22337]"
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:341
	// _ = "end of CoverTab[22298]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:341
	_go_fuzz_dep_.CoverTab[22299]++
								if !cipherSuiteOk {
//line /usr/local/go/src/crypto/tls/handshake_client.go:342
		_go_fuzz_dep_.CoverTab[22341]++
									return cacheKey, nil, nil, nil, nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:343
		// _ = "end of CoverTab[22341]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:344
		_go_fuzz_dep_.CoverTab[22342]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:344
		// _ = "end of CoverTab[22342]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:344
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:344
	// _ = "end of CoverTab[22299]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:344
	_go_fuzz_dep_.CoverTab[22300]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:347
	ticketAge := uint32(c.config.time().Sub(session.receivedAt) / time.Millisecond)
	identity := pskIdentity{
		label:			session.sessionTicket,
		obfuscatedTicketAge:	ticketAge + session.ageAdd,
	}
								hello.pskIdentities = []pskIdentity{identity}
								hello.pskBinders = [][]byte{make([]byte, cipherSuite.hash.Size())}

//line /usr/local/go/src/crypto/tls/handshake_client.go:356
	psk := cipherSuite.expandLabel(session.masterSecret, "resumption",
		session.nonce, cipherSuite.hash.Size())
	earlySecret = cipherSuite.extract(psk, nil)
	binderKey = cipherSuite.deriveSecret(earlySecret, resumptionBinderLabel, nil)
	transcript := cipherSuite.hash.New()
	helloBytes, err := hello.marshalWithoutBinders()
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:362
		_go_fuzz_dep_.CoverTab[22343]++
									return "", nil, nil, nil, err
//line /usr/local/go/src/crypto/tls/handshake_client.go:363
		// _ = "end of CoverTab[22343]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:364
		_go_fuzz_dep_.CoverTab[22344]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:364
		// _ = "end of CoverTab[22344]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:364
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:364
	// _ = "end of CoverTab[22300]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:364
	_go_fuzz_dep_.CoverTab[22301]++
								transcript.Write(helloBytes)
								pskBinders := [][]byte{cipherSuite.finishedHash(binderKey, transcript)}
								if err := hello.updateBinders(pskBinders); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:367
		_go_fuzz_dep_.CoverTab[22345]++
									return "", nil, nil, nil, err
//line /usr/local/go/src/crypto/tls/handshake_client.go:368
		// _ = "end of CoverTab[22345]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:369
		_go_fuzz_dep_.CoverTab[22346]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:369
		// _ = "end of CoverTab[22346]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:369
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:369
	// _ = "end of CoverTab[22301]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:369
	_go_fuzz_dep_.CoverTab[22302]++

								return
//line /usr/local/go/src/crypto/tls/handshake_client.go:371
	// _ = "end of CoverTab[22302]"
}

func (c *Conn) pickTLSVersion(serverHello *serverHelloMsg) error {
//line /usr/local/go/src/crypto/tls/handshake_client.go:374
	_go_fuzz_dep_.CoverTab[22347]++
								peerVersion := serverHello.vers
								if serverHello.supportedVersion != 0 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:376
		_go_fuzz_dep_.CoverTab[22350]++
									peerVersion = serverHello.supportedVersion
//line /usr/local/go/src/crypto/tls/handshake_client.go:377
		// _ = "end of CoverTab[22350]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:378
		_go_fuzz_dep_.CoverTab[22351]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:378
		// _ = "end of CoverTab[22351]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:378
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:378
	// _ = "end of CoverTab[22347]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:378
	_go_fuzz_dep_.CoverTab[22348]++

								vers, ok := c.config.mutualVersion(roleClient, []uint16{peerVersion})
								if !ok {
//line /usr/local/go/src/crypto/tls/handshake_client.go:381
		_go_fuzz_dep_.CoverTab[22352]++
									c.sendAlert(alertProtocolVersion)
									return fmt.Errorf("tls: server selected unsupported protocol version %x", peerVersion)
//line /usr/local/go/src/crypto/tls/handshake_client.go:383
		// _ = "end of CoverTab[22352]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:384
		_go_fuzz_dep_.CoverTab[22353]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:384
		// _ = "end of CoverTab[22353]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:384
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:384
	// _ = "end of CoverTab[22348]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:384
	_go_fuzz_dep_.CoverTab[22349]++

								c.vers = vers
								c.haveVers = true
								c.in.version = vers
								c.out.version = vers

								return nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:391
	// _ = "end of CoverTab[22349]"
}

// Does the handshake, either a full one or resumes old session. Requires hs.c,
//line /usr/local/go/src/crypto/tls/handshake_client.go:394
// hs.hello, hs.serverHello, and, optionally, hs.session to be set.
//line /usr/local/go/src/crypto/tls/handshake_client.go:396
func (hs *clientHandshakeState) handshake() error {
//line /usr/local/go/src/crypto/tls/handshake_client.go:396
	_go_fuzz_dep_.CoverTab[22354]++
								c := hs.c

								isResume, err := hs.processServerHello()
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:400
		_go_fuzz_dep_.CoverTab[22360]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:401
		// _ = "end of CoverTab[22360]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:402
		_go_fuzz_dep_.CoverTab[22361]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:402
		// _ = "end of CoverTab[22361]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:402
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:402
	// _ = "end of CoverTab[22354]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:402
	_go_fuzz_dep_.CoverTab[22355]++

								hs.finishedHash = newFinishedHash(c.vers, hs.suite)

//line /usr/local/go/src/crypto/tls/handshake_client.go:410
	if isResume || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:410
		_go_fuzz_dep_.CoverTab[22362]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:410
		return (len(c.config.Certificates) == 0 && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:410
			_go_fuzz_dep_.CoverTab[22363]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:410
			return c.config.GetClientCertificate == nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:410
			// _ = "end of CoverTab[22363]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:410
		}())
//line /usr/local/go/src/crypto/tls/handshake_client.go:410
		// _ = "end of CoverTab[22362]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:410
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:410
		_go_fuzz_dep_.CoverTab[22364]++
									hs.finishedHash.discardHandshakeBuffer()
//line /usr/local/go/src/crypto/tls/handshake_client.go:411
		// _ = "end of CoverTab[22364]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:412
		_go_fuzz_dep_.CoverTab[22365]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:412
		// _ = "end of CoverTab[22365]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:412
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:412
	// _ = "end of CoverTab[22355]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:412
	_go_fuzz_dep_.CoverTab[22356]++

								if err := transcriptMsg(hs.hello, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:414
		_go_fuzz_dep_.CoverTab[22366]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:415
		// _ = "end of CoverTab[22366]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:416
		_go_fuzz_dep_.CoverTab[22367]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:416
		// _ = "end of CoverTab[22367]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:416
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:416
	// _ = "end of CoverTab[22356]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:416
	_go_fuzz_dep_.CoverTab[22357]++
								if err := transcriptMsg(hs.serverHello, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:417
		_go_fuzz_dep_.CoverTab[22368]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:418
		// _ = "end of CoverTab[22368]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:419
		_go_fuzz_dep_.CoverTab[22369]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:419
		// _ = "end of CoverTab[22369]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:419
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:419
	// _ = "end of CoverTab[22357]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:419
	_go_fuzz_dep_.CoverTab[22358]++

								c.buffering = true
								c.didResume = isResume
								if isResume {
//line /usr/local/go/src/crypto/tls/handshake_client.go:423
		_go_fuzz_dep_.CoverTab[22370]++
									if err := hs.establishKeys(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:424
			_go_fuzz_dep_.CoverTab[22376]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:425
			// _ = "end of CoverTab[22376]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:426
			_go_fuzz_dep_.CoverTab[22377]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:426
			// _ = "end of CoverTab[22377]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:426
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:426
		// _ = "end of CoverTab[22370]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:426
		_go_fuzz_dep_.CoverTab[22371]++
									if err := hs.readSessionTicket(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:427
			_go_fuzz_dep_.CoverTab[22378]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:428
			// _ = "end of CoverTab[22378]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:429
			_go_fuzz_dep_.CoverTab[22379]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:429
			// _ = "end of CoverTab[22379]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:429
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:429
		// _ = "end of CoverTab[22371]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:429
		_go_fuzz_dep_.CoverTab[22372]++
									if err := hs.readFinished(c.serverFinished[:]); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:430
			_go_fuzz_dep_.CoverTab[22380]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:431
			// _ = "end of CoverTab[22380]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:432
			_go_fuzz_dep_.CoverTab[22381]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:432
			// _ = "end of CoverTab[22381]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:432
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:432
		// _ = "end of CoverTab[22372]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:432
		_go_fuzz_dep_.CoverTab[22373]++
									c.clientFinishedIsFirst = false

//line /usr/local/go/src/crypto/tls/handshake_client.go:437
		if c.config.VerifyConnection != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:437
			_go_fuzz_dep_.CoverTab[22382]++
										if err := c.config.VerifyConnection(c.connectionStateLocked()); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:438
				_go_fuzz_dep_.CoverTab[22383]++
											c.sendAlert(alertBadCertificate)
											return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:440
				// _ = "end of CoverTab[22383]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:441
				_go_fuzz_dep_.CoverTab[22384]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:441
				// _ = "end of CoverTab[22384]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:441
			}
//line /usr/local/go/src/crypto/tls/handshake_client.go:441
			// _ = "end of CoverTab[22382]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:442
			_go_fuzz_dep_.CoverTab[22385]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:442
			// _ = "end of CoverTab[22385]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:442
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:442
		// _ = "end of CoverTab[22373]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:442
		_go_fuzz_dep_.CoverTab[22374]++
									if err := hs.sendFinished(c.clientFinished[:]); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:443
			_go_fuzz_dep_.CoverTab[22386]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:444
			// _ = "end of CoverTab[22386]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:445
			_go_fuzz_dep_.CoverTab[22387]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:445
			// _ = "end of CoverTab[22387]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:445
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:445
		// _ = "end of CoverTab[22374]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:445
		_go_fuzz_dep_.CoverTab[22375]++
									if _, err := c.flush(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:446
			_go_fuzz_dep_.CoverTab[22388]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:447
			// _ = "end of CoverTab[22388]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:448
			_go_fuzz_dep_.CoverTab[22389]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:448
			// _ = "end of CoverTab[22389]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:448
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:448
		// _ = "end of CoverTab[22375]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:449
		_go_fuzz_dep_.CoverTab[22390]++
									if err := hs.doFullHandshake(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:450
			_go_fuzz_dep_.CoverTab[22396]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:451
			// _ = "end of CoverTab[22396]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:452
			_go_fuzz_dep_.CoverTab[22397]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:452
			// _ = "end of CoverTab[22397]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:452
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:452
		// _ = "end of CoverTab[22390]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:452
		_go_fuzz_dep_.CoverTab[22391]++
									if err := hs.establishKeys(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:453
			_go_fuzz_dep_.CoverTab[22398]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:454
			// _ = "end of CoverTab[22398]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:455
			_go_fuzz_dep_.CoverTab[22399]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:455
			// _ = "end of CoverTab[22399]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:455
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:455
		// _ = "end of CoverTab[22391]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:455
		_go_fuzz_dep_.CoverTab[22392]++
									if err := hs.sendFinished(c.clientFinished[:]); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:456
			_go_fuzz_dep_.CoverTab[22400]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:457
			// _ = "end of CoverTab[22400]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:458
			_go_fuzz_dep_.CoverTab[22401]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:458
			// _ = "end of CoverTab[22401]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:458
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:458
		// _ = "end of CoverTab[22392]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:458
		_go_fuzz_dep_.CoverTab[22393]++
									if _, err := c.flush(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:459
			_go_fuzz_dep_.CoverTab[22402]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:460
			// _ = "end of CoverTab[22402]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:461
			_go_fuzz_dep_.CoverTab[22403]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:461
			// _ = "end of CoverTab[22403]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:461
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:461
		// _ = "end of CoverTab[22393]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:461
		_go_fuzz_dep_.CoverTab[22394]++
									c.clientFinishedIsFirst = true
									if err := hs.readSessionTicket(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:463
			_go_fuzz_dep_.CoverTab[22404]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:464
			// _ = "end of CoverTab[22404]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:465
			_go_fuzz_dep_.CoverTab[22405]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:465
			// _ = "end of CoverTab[22405]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:465
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:465
		// _ = "end of CoverTab[22394]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:465
		_go_fuzz_dep_.CoverTab[22395]++
									if err := hs.readFinished(c.serverFinished[:]); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:466
			_go_fuzz_dep_.CoverTab[22406]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:467
			// _ = "end of CoverTab[22406]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:468
			_go_fuzz_dep_.CoverTab[22407]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:468
			// _ = "end of CoverTab[22407]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:468
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:468
		// _ = "end of CoverTab[22395]"
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:469
	// _ = "end of CoverTab[22358]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:469
	_go_fuzz_dep_.CoverTab[22359]++

								c.ekm = ekmFromMasterSecret(c.vers, hs.suite, hs.masterSecret, hs.hello.random, hs.serverHello.random)
								c.isHandshakeComplete.Store(true)

								return nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:474
	// _ = "end of CoverTab[22359]"
}

func (hs *clientHandshakeState) pickCipherSuite() error {
//line /usr/local/go/src/crypto/tls/handshake_client.go:477
	_go_fuzz_dep_.CoverTab[22408]++
								if hs.suite = mutualCipherSuite(hs.hello.cipherSuites, hs.serverHello.cipherSuite); hs.suite == nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:478
		_go_fuzz_dep_.CoverTab[22410]++
									hs.c.sendAlert(alertHandshakeFailure)
									return errors.New("tls: server chose an unconfigured cipher suite")
//line /usr/local/go/src/crypto/tls/handshake_client.go:480
		// _ = "end of CoverTab[22410]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:481
		_go_fuzz_dep_.CoverTab[22411]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:481
		// _ = "end of CoverTab[22411]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:481
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:481
	// _ = "end of CoverTab[22408]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:481
	_go_fuzz_dep_.CoverTab[22409]++

								hs.c.cipherSuite = hs.suite.id
								return nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:484
	// _ = "end of CoverTab[22409]"
}

func (hs *clientHandshakeState) doFullHandshake() error {
//line /usr/local/go/src/crypto/tls/handshake_client.go:487
	_go_fuzz_dep_.CoverTab[22412]++
								c := hs.c

								msg, err := c.readHandshake(&hs.finishedHash)
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:491
		_go_fuzz_dep_.CoverTab[22426]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:492
		// _ = "end of CoverTab[22426]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:493
		_go_fuzz_dep_.CoverTab[22427]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:493
		// _ = "end of CoverTab[22427]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:493
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:493
	// _ = "end of CoverTab[22412]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:493
	_go_fuzz_dep_.CoverTab[22413]++
								certMsg, ok := msg.(*certificateMsg)
								if !ok || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:495
		_go_fuzz_dep_.CoverTab[22428]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:495
		return len(certMsg.certificates) == 0
//line /usr/local/go/src/crypto/tls/handshake_client.go:495
		// _ = "end of CoverTab[22428]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:495
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:495
		_go_fuzz_dep_.CoverTab[22429]++
									c.sendAlert(alertUnexpectedMessage)
									return unexpectedMessageError(certMsg, msg)
//line /usr/local/go/src/crypto/tls/handshake_client.go:497
		// _ = "end of CoverTab[22429]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:498
		_go_fuzz_dep_.CoverTab[22430]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:498
		// _ = "end of CoverTab[22430]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:498
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:498
	// _ = "end of CoverTab[22413]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:498
	_go_fuzz_dep_.CoverTab[22414]++

								msg, err = c.readHandshake(&hs.finishedHash)
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:501
		_go_fuzz_dep_.CoverTab[22431]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:502
		// _ = "end of CoverTab[22431]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:503
		_go_fuzz_dep_.CoverTab[22432]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:503
		// _ = "end of CoverTab[22432]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:503
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:503
	// _ = "end of CoverTab[22414]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:503
	_go_fuzz_dep_.CoverTab[22415]++

								cs, ok := msg.(*certificateStatusMsg)
								if ok {
//line /usr/local/go/src/crypto/tls/handshake_client.go:506
		_go_fuzz_dep_.CoverTab[22433]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:510
		if !hs.serverHello.ocspStapling {
//line /usr/local/go/src/crypto/tls/handshake_client.go:510
			_go_fuzz_dep_.CoverTab[22435]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:515
			c.sendAlert(alertUnexpectedMessage)
										return errors.New("tls: received unexpected CertificateStatus message")
//line /usr/local/go/src/crypto/tls/handshake_client.go:516
			// _ = "end of CoverTab[22435]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:517
			_go_fuzz_dep_.CoverTab[22436]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:517
			// _ = "end of CoverTab[22436]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:517
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:517
		// _ = "end of CoverTab[22433]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:517
		_go_fuzz_dep_.CoverTab[22434]++

									c.ocspResponse = cs.response

									msg, err = c.readHandshake(&hs.finishedHash)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:522
			_go_fuzz_dep_.CoverTab[22437]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:523
			// _ = "end of CoverTab[22437]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:524
			_go_fuzz_dep_.CoverTab[22438]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:524
			// _ = "end of CoverTab[22438]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:524
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:524
		// _ = "end of CoverTab[22434]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:525
		_go_fuzz_dep_.CoverTab[22439]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:525
		// _ = "end of CoverTab[22439]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:525
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:525
	// _ = "end of CoverTab[22415]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:525
	_go_fuzz_dep_.CoverTab[22416]++

								if c.handshakes == 0 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:527
		_go_fuzz_dep_.CoverTab[22440]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:530
		if err := c.verifyServerCertificate(certMsg.certificates); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:530
			_go_fuzz_dep_.CoverTab[22441]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:531
			// _ = "end of CoverTab[22441]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:532
			_go_fuzz_dep_.CoverTab[22442]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:532
			// _ = "end of CoverTab[22442]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:532
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:532
		// _ = "end of CoverTab[22440]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:533
		_go_fuzz_dep_.CoverTab[22443]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:540
		if !bytes.Equal(c.peerCertificates[0].Raw, certMsg.certificates[0]) {
//line /usr/local/go/src/crypto/tls/handshake_client.go:540
			_go_fuzz_dep_.CoverTab[22444]++
										c.sendAlert(alertBadCertificate)
										return errors.New("tls: server's identity changed during renegotiation")
//line /usr/local/go/src/crypto/tls/handshake_client.go:542
			// _ = "end of CoverTab[22444]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:543
			_go_fuzz_dep_.CoverTab[22445]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:543
			// _ = "end of CoverTab[22445]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:543
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:543
		// _ = "end of CoverTab[22443]"
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:544
	// _ = "end of CoverTab[22416]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:544
	_go_fuzz_dep_.CoverTab[22417]++

								keyAgreement := hs.suite.ka(c.vers)

								skx, ok := msg.(*serverKeyExchangeMsg)
								if ok {
//line /usr/local/go/src/crypto/tls/handshake_client.go:549
		_go_fuzz_dep_.CoverTab[22446]++
									err = keyAgreement.processServerKeyExchange(c.config, hs.hello, hs.serverHello, c.peerCertificates[0], skx)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:551
			_go_fuzz_dep_.CoverTab[22448]++
										c.sendAlert(alertUnexpectedMessage)
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:553
			// _ = "end of CoverTab[22448]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:554
			_go_fuzz_dep_.CoverTab[22449]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:554
			// _ = "end of CoverTab[22449]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:554
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:554
		// _ = "end of CoverTab[22446]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:554
		_go_fuzz_dep_.CoverTab[22447]++

									msg, err = c.readHandshake(&hs.finishedHash)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:557
			_go_fuzz_dep_.CoverTab[22450]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:558
			// _ = "end of CoverTab[22450]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:559
			_go_fuzz_dep_.CoverTab[22451]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:559
			// _ = "end of CoverTab[22451]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:559
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:559
		// _ = "end of CoverTab[22447]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:560
		_go_fuzz_dep_.CoverTab[22452]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:560
		// _ = "end of CoverTab[22452]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:560
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:560
	// _ = "end of CoverTab[22417]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:560
	_go_fuzz_dep_.CoverTab[22418]++

								var chainToSend *Certificate
								var certRequested bool
								certReq, ok := msg.(*certificateRequestMsg)
								if ok {
//line /usr/local/go/src/crypto/tls/handshake_client.go:565
		_go_fuzz_dep_.CoverTab[22453]++
									certRequested = true

									cri := certificateRequestInfoFromMsg(hs.ctx, c.vers, certReq)
									if chainToSend, err = c.getClientCertificate(cri); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:569
			_go_fuzz_dep_.CoverTab[22455]++
										c.sendAlert(alertInternalError)
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:571
			// _ = "end of CoverTab[22455]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:572
			_go_fuzz_dep_.CoverTab[22456]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:572
			// _ = "end of CoverTab[22456]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:572
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:572
		// _ = "end of CoverTab[22453]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:572
		_go_fuzz_dep_.CoverTab[22454]++

									msg, err = c.readHandshake(&hs.finishedHash)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:575
			_go_fuzz_dep_.CoverTab[22457]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:576
			// _ = "end of CoverTab[22457]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:577
			_go_fuzz_dep_.CoverTab[22458]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:577
			// _ = "end of CoverTab[22458]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:577
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:577
		// _ = "end of CoverTab[22454]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:578
		_go_fuzz_dep_.CoverTab[22459]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:578
		// _ = "end of CoverTab[22459]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:578
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:578
	// _ = "end of CoverTab[22418]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:578
	_go_fuzz_dep_.CoverTab[22419]++

								shd, ok := msg.(*serverHelloDoneMsg)
								if !ok {
//line /usr/local/go/src/crypto/tls/handshake_client.go:581
		_go_fuzz_dep_.CoverTab[22460]++
									c.sendAlert(alertUnexpectedMessage)
									return unexpectedMessageError(shd, msg)
//line /usr/local/go/src/crypto/tls/handshake_client.go:583
		// _ = "end of CoverTab[22460]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:584
		_go_fuzz_dep_.CoverTab[22461]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:584
		// _ = "end of CoverTab[22461]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:584
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:584
	// _ = "end of CoverTab[22419]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:584
	_go_fuzz_dep_.CoverTab[22420]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:589
	if certRequested {
//line /usr/local/go/src/crypto/tls/handshake_client.go:589
		_go_fuzz_dep_.CoverTab[22462]++
									certMsg = new(certificateMsg)
									certMsg.certificates = chainToSend.Certificate
									if _, err := hs.c.writeHandshakeRecord(certMsg, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:592
			_go_fuzz_dep_.CoverTab[22463]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:593
			// _ = "end of CoverTab[22463]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:594
			_go_fuzz_dep_.CoverTab[22464]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:594
			// _ = "end of CoverTab[22464]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:594
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:594
		// _ = "end of CoverTab[22462]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:595
		_go_fuzz_dep_.CoverTab[22465]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:595
		// _ = "end of CoverTab[22465]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:595
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:595
	// _ = "end of CoverTab[22420]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:595
	_go_fuzz_dep_.CoverTab[22421]++

								preMasterSecret, ckx, err := keyAgreement.generateClientKeyExchange(c.config, hs.hello, c.peerCertificates[0])
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:598
		_go_fuzz_dep_.CoverTab[22466]++
									c.sendAlert(alertInternalError)
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:600
		// _ = "end of CoverTab[22466]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:601
		_go_fuzz_dep_.CoverTab[22467]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:601
		// _ = "end of CoverTab[22467]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:601
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:601
	// _ = "end of CoverTab[22421]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:601
	_go_fuzz_dep_.CoverTab[22422]++
								if ckx != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:602
		_go_fuzz_dep_.CoverTab[22468]++
									if _, err := hs.c.writeHandshakeRecord(ckx, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:603
			_go_fuzz_dep_.CoverTab[22469]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:604
			// _ = "end of CoverTab[22469]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:605
			_go_fuzz_dep_.CoverTab[22470]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:605
			// _ = "end of CoverTab[22470]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:605
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:605
		// _ = "end of CoverTab[22468]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:606
		_go_fuzz_dep_.CoverTab[22471]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:606
		// _ = "end of CoverTab[22471]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:606
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:606
	// _ = "end of CoverTab[22422]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:606
	_go_fuzz_dep_.CoverTab[22423]++

								if chainToSend != nil && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:608
		_go_fuzz_dep_.CoverTab[22472]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:608
		return len(chainToSend.Certificate) > 0
//line /usr/local/go/src/crypto/tls/handshake_client.go:608
		// _ = "end of CoverTab[22472]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:608
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:608
		_go_fuzz_dep_.CoverTab[22473]++
									certVerify := &certificateVerifyMsg{}

									key, ok := chainToSend.PrivateKey.(crypto.Signer)
									if !ok {
//line /usr/local/go/src/crypto/tls/handshake_client.go:612
			_go_fuzz_dep_.CoverTab[22478]++
										c.sendAlert(alertInternalError)
										return fmt.Errorf("tls: client certificate private key of type %T does not implement crypto.Signer", chainToSend.PrivateKey)
//line /usr/local/go/src/crypto/tls/handshake_client.go:614
			// _ = "end of CoverTab[22478]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:615
			_go_fuzz_dep_.CoverTab[22479]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:615
			// _ = "end of CoverTab[22479]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:615
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:615
		// _ = "end of CoverTab[22473]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:615
		_go_fuzz_dep_.CoverTab[22474]++

									var sigType uint8
									var sigHash crypto.Hash
									if c.vers >= VersionTLS12 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:619
			_go_fuzz_dep_.CoverTab[22480]++
										signatureAlgorithm, err := selectSignatureScheme(c.vers, chainToSend, certReq.supportedSignatureAlgorithms)
										if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:621
				_go_fuzz_dep_.CoverTab[22483]++
											c.sendAlert(alertIllegalParameter)
											return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:623
				// _ = "end of CoverTab[22483]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:624
				_go_fuzz_dep_.CoverTab[22484]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:624
				// _ = "end of CoverTab[22484]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:624
			}
//line /usr/local/go/src/crypto/tls/handshake_client.go:624
			// _ = "end of CoverTab[22480]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:624
			_go_fuzz_dep_.CoverTab[22481]++
										sigType, sigHash, err = typeAndHashFromSignatureScheme(signatureAlgorithm)
										if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:626
				_go_fuzz_dep_.CoverTab[22485]++
											return c.sendAlert(alertInternalError)
//line /usr/local/go/src/crypto/tls/handshake_client.go:627
				// _ = "end of CoverTab[22485]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:628
				_go_fuzz_dep_.CoverTab[22486]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:628
				// _ = "end of CoverTab[22486]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:628
			}
//line /usr/local/go/src/crypto/tls/handshake_client.go:628
			// _ = "end of CoverTab[22481]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:628
			_go_fuzz_dep_.CoverTab[22482]++
										certVerify.hasSignatureAlgorithm = true
										certVerify.signatureAlgorithm = signatureAlgorithm
//line /usr/local/go/src/crypto/tls/handshake_client.go:630
			// _ = "end of CoverTab[22482]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:631
			_go_fuzz_dep_.CoverTab[22487]++
										sigType, sigHash, err = legacyTypeAndHashFromPublicKey(key.Public())
										if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:633
				_go_fuzz_dep_.CoverTab[22488]++
											c.sendAlert(alertIllegalParameter)
											return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:635
				// _ = "end of CoverTab[22488]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:636
				_go_fuzz_dep_.CoverTab[22489]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:636
				// _ = "end of CoverTab[22489]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:636
			}
//line /usr/local/go/src/crypto/tls/handshake_client.go:636
			// _ = "end of CoverTab[22487]"
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:637
		// _ = "end of CoverTab[22474]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:637
		_go_fuzz_dep_.CoverTab[22475]++

									signed := hs.finishedHash.hashForClientCertificate(sigType, sigHash)
									signOpts := crypto.SignerOpts(sigHash)
									if sigType == signatureRSAPSS {
//line /usr/local/go/src/crypto/tls/handshake_client.go:641
			_go_fuzz_dep_.CoverTab[22490]++
										signOpts = &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash, Hash: sigHash}
//line /usr/local/go/src/crypto/tls/handshake_client.go:642
			// _ = "end of CoverTab[22490]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:643
			_go_fuzz_dep_.CoverTab[22491]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:643
			// _ = "end of CoverTab[22491]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:643
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:643
		// _ = "end of CoverTab[22475]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:643
		_go_fuzz_dep_.CoverTab[22476]++
									certVerify.signature, err = key.Sign(c.config.rand(), signed, signOpts)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:645
			_go_fuzz_dep_.CoverTab[22492]++
										c.sendAlert(alertInternalError)
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:647
			// _ = "end of CoverTab[22492]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:648
			_go_fuzz_dep_.CoverTab[22493]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:648
			// _ = "end of CoverTab[22493]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:648
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:648
		// _ = "end of CoverTab[22476]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:648
		_go_fuzz_dep_.CoverTab[22477]++

									if _, err := hs.c.writeHandshakeRecord(certVerify, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:650
			_go_fuzz_dep_.CoverTab[22494]++
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:651
			// _ = "end of CoverTab[22494]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:652
			_go_fuzz_dep_.CoverTab[22495]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:652
			// _ = "end of CoverTab[22495]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:652
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:652
		// _ = "end of CoverTab[22477]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:653
		_go_fuzz_dep_.CoverTab[22496]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:653
		// _ = "end of CoverTab[22496]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:653
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:653
	// _ = "end of CoverTab[22423]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:653
	_go_fuzz_dep_.CoverTab[22424]++

								hs.masterSecret = masterFromPreMasterSecret(c.vers, hs.suite, preMasterSecret, hs.hello.random, hs.serverHello.random)
								if err := c.config.writeKeyLog(keyLogLabelTLS12, hs.hello.random, hs.masterSecret); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:656
		_go_fuzz_dep_.CoverTab[22497]++
									c.sendAlert(alertInternalError)
									return errors.New("tls: failed to write to key log: " + err.Error())
//line /usr/local/go/src/crypto/tls/handshake_client.go:658
		// _ = "end of CoverTab[22497]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:659
		_go_fuzz_dep_.CoverTab[22498]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:659
		// _ = "end of CoverTab[22498]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:659
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:659
	// _ = "end of CoverTab[22424]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:659
	_go_fuzz_dep_.CoverTab[22425]++

								hs.finishedHash.discardHandshakeBuffer()

								return nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:663
	// _ = "end of CoverTab[22425]"
}

func (hs *clientHandshakeState) establishKeys() error {
//line /usr/local/go/src/crypto/tls/handshake_client.go:666
	_go_fuzz_dep_.CoverTab[22499]++
								c := hs.c

								clientMAC, serverMAC, clientKey, serverKey, clientIV, serverIV :=
		keysFromMasterSecret(c.vers, hs.suite, hs.masterSecret, hs.hello.random, hs.serverHello.random, hs.suite.macLen, hs.suite.keyLen, hs.suite.ivLen)
	var clientCipher, serverCipher any
	var clientHash, serverHash hash.Hash
	if hs.suite.cipher != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:673
		_go_fuzz_dep_.CoverTab[22501]++
									clientCipher = hs.suite.cipher(clientKey, clientIV, false)
									clientHash = hs.suite.mac(clientMAC)
									serverCipher = hs.suite.cipher(serverKey, serverIV, true)
									serverHash = hs.suite.mac(serverMAC)
//line /usr/local/go/src/crypto/tls/handshake_client.go:677
		// _ = "end of CoverTab[22501]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:678
		_go_fuzz_dep_.CoverTab[22502]++
									clientCipher = hs.suite.aead(clientKey, clientIV)
									serverCipher = hs.suite.aead(serverKey, serverIV)
//line /usr/local/go/src/crypto/tls/handshake_client.go:680
		// _ = "end of CoverTab[22502]"
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:681
	// _ = "end of CoverTab[22499]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:681
	_go_fuzz_dep_.CoverTab[22500]++

								c.in.prepareCipherSpec(c.vers, serverCipher, serverHash)
								c.out.prepareCipherSpec(c.vers, clientCipher, clientHash)
								return nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:685
	// _ = "end of CoverTab[22500]"
}

func (hs *clientHandshakeState) serverResumedSession() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:688
	_go_fuzz_dep_.CoverTab[22503]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:691
	return hs.session != nil && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:691
		_go_fuzz_dep_.CoverTab[22504]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:691
		return hs.hello.sessionId != nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:691
		// _ = "end of CoverTab[22504]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:691
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:691
		_go_fuzz_dep_.CoverTab[22505]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:691
		return bytes.Equal(hs.serverHello.sessionId, hs.hello.sessionId)
									// _ = "end of CoverTab[22505]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:692
	}()
//line /usr/local/go/src/crypto/tls/handshake_client.go:692
	// _ = "end of CoverTab[22503]"
}

func (hs *clientHandshakeState) processServerHello() (bool, error) {
//line /usr/local/go/src/crypto/tls/handshake_client.go:695
	_go_fuzz_dep_.CoverTab[22506]++
								c := hs.c

								if err := hs.pickCipherSuite(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:698
		_go_fuzz_dep_.CoverTab[22516]++
									return false, err
//line /usr/local/go/src/crypto/tls/handshake_client.go:699
		// _ = "end of CoverTab[22516]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:700
		_go_fuzz_dep_.CoverTab[22517]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:700
		// _ = "end of CoverTab[22517]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:700
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:700
	// _ = "end of CoverTab[22506]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:700
	_go_fuzz_dep_.CoverTab[22507]++

								if hs.serverHello.compressionMethod != compressionNone {
//line /usr/local/go/src/crypto/tls/handshake_client.go:702
		_go_fuzz_dep_.CoverTab[22518]++
									c.sendAlert(alertUnexpectedMessage)
									return false, errors.New("tls: server selected unsupported compression format")
//line /usr/local/go/src/crypto/tls/handshake_client.go:704
		// _ = "end of CoverTab[22518]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:705
		_go_fuzz_dep_.CoverTab[22519]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:705
		// _ = "end of CoverTab[22519]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:705
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:705
	// _ = "end of CoverTab[22507]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:705
	_go_fuzz_dep_.CoverTab[22508]++

								if c.handshakes == 0 && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:707
		_go_fuzz_dep_.CoverTab[22520]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:707
		return hs.serverHello.secureRenegotiationSupported
//line /usr/local/go/src/crypto/tls/handshake_client.go:707
		// _ = "end of CoverTab[22520]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:707
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:707
		_go_fuzz_dep_.CoverTab[22521]++
									c.secureRenegotiation = true
									if len(hs.serverHello.secureRenegotiation) != 0 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:709
			_go_fuzz_dep_.CoverTab[22522]++
										c.sendAlert(alertHandshakeFailure)
										return false, errors.New("tls: initial handshake had non-empty renegotiation extension")
//line /usr/local/go/src/crypto/tls/handshake_client.go:711
			// _ = "end of CoverTab[22522]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:712
			_go_fuzz_dep_.CoverTab[22523]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:712
			// _ = "end of CoverTab[22523]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:712
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:712
		// _ = "end of CoverTab[22521]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:713
		_go_fuzz_dep_.CoverTab[22524]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:713
		// _ = "end of CoverTab[22524]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:713
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:713
	// _ = "end of CoverTab[22508]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:713
	_go_fuzz_dep_.CoverTab[22509]++

								if c.handshakes > 0 && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:715
		_go_fuzz_dep_.CoverTab[22525]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:715
		return c.secureRenegotiation
//line /usr/local/go/src/crypto/tls/handshake_client.go:715
		// _ = "end of CoverTab[22525]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:715
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:715
		_go_fuzz_dep_.CoverTab[22526]++
									var expectedSecureRenegotiation [24]byte
									copy(expectedSecureRenegotiation[:], c.clientFinished[:])
									copy(expectedSecureRenegotiation[12:], c.serverFinished[:])
									if !bytes.Equal(hs.serverHello.secureRenegotiation, expectedSecureRenegotiation[:]) {
//line /usr/local/go/src/crypto/tls/handshake_client.go:719
			_go_fuzz_dep_.CoverTab[22527]++
										c.sendAlert(alertHandshakeFailure)
										return false, errors.New("tls: incorrect renegotiation extension contents")
//line /usr/local/go/src/crypto/tls/handshake_client.go:721
			// _ = "end of CoverTab[22527]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:722
			_go_fuzz_dep_.CoverTab[22528]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:722
			// _ = "end of CoverTab[22528]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:722
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:722
		// _ = "end of CoverTab[22526]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:723
		_go_fuzz_dep_.CoverTab[22529]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:723
		// _ = "end of CoverTab[22529]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:723
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:723
	// _ = "end of CoverTab[22509]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:723
	_go_fuzz_dep_.CoverTab[22510]++

								if err := checkALPN(hs.hello.alpnProtocols, hs.serverHello.alpnProtocol); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:725
		_go_fuzz_dep_.CoverTab[22530]++
									c.sendAlert(alertUnsupportedExtension)
									return false, err
//line /usr/local/go/src/crypto/tls/handshake_client.go:727
		// _ = "end of CoverTab[22530]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:728
		_go_fuzz_dep_.CoverTab[22531]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:728
		// _ = "end of CoverTab[22531]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:728
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:728
	// _ = "end of CoverTab[22510]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:728
	_go_fuzz_dep_.CoverTab[22511]++
								c.clientProtocol = hs.serverHello.alpnProtocol

								c.scts = hs.serverHello.scts

								if !hs.serverResumedSession() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:733
		_go_fuzz_dep_.CoverTab[22532]++
									return false, nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:734
		// _ = "end of CoverTab[22532]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:735
		_go_fuzz_dep_.CoverTab[22533]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:735
		// _ = "end of CoverTab[22533]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:735
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:735
	// _ = "end of CoverTab[22511]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:735
	_go_fuzz_dep_.CoverTab[22512]++

								if hs.session.vers != c.vers {
//line /usr/local/go/src/crypto/tls/handshake_client.go:737
		_go_fuzz_dep_.CoverTab[22534]++
									c.sendAlert(alertHandshakeFailure)
									return false, errors.New("tls: server resumed a session with a different version")
//line /usr/local/go/src/crypto/tls/handshake_client.go:739
		// _ = "end of CoverTab[22534]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:740
		_go_fuzz_dep_.CoverTab[22535]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:740
		// _ = "end of CoverTab[22535]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:740
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:740
	// _ = "end of CoverTab[22512]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:740
	_go_fuzz_dep_.CoverTab[22513]++

								if hs.session.cipherSuite != hs.suite.id {
//line /usr/local/go/src/crypto/tls/handshake_client.go:742
		_go_fuzz_dep_.CoverTab[22536]++
									c.sendAlert(alertHandshakeFailure)
									return false, errors.New("tls: server resumed a session with a different cipher suite")
//line /usr/local/go/src/crypto/tls/handshake_client.go:744
		// _ = "end of CoverTab[22536]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:745
		_go_fuzz_dep_.CoverTab[22537]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:745
		// _ = "end of CoverTab[22537]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:745
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:745
	// _ = "end of CoverTab[22513]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:745
	_go_fuzz_dep_.CoverTab[22514]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:748
	hs.masterSecret = hs.session.masterSecret
								c.peerCertificates = hs.session.serverCertificates
								c.verifiedChains = hs.session.verifiedChains
								c.ocspResponse = hs.session.ocspResponse

//line /usr/local/go/src/crypto/tls/handshake_client.go:754
	if len(c.scts) == 0 && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:754
		_go_fuzz_dep_.CoverTab[22538]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:754
		return len(hs.session.scts) != 0
//line /usr/local/go/src/crypto/tls/handshake_client.go:754
		// _ = "end of CoverTab[22538]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:754
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:754
		_go_fuzz_dep_.CoverTab[22539]++
									c.scts = hs.session.scts
//line /usr/local/go/src/crypto/tls/handshake_client.go:755
		// _ = "end of CoverTab[22539]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:756
		_go_fuzz_dep_.CoverTab[22540]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:756
		// _ = "end of CoverTab[22540]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:756
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:756
	// _ = "end of CoverTab[22514]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:756
	_go_fuzz_dep_.CoverTab[22515]++

								return true, nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:758
	// _ = "end of CoverTab[22515]"
}

// checkALPN ensure that the server's choice of ALPN protocol is compatible with
//line /usr/local/go/src/crypto/tls/handshake_client.go:761
// the protocols that we advertised in the Client Hello.
//line /usr/local/go/src/crypto/tls/handshake_client.go:763
func checkALPN(clientProtos []string, serverProto string) error {
//line /usr/local/go/src/crypto/tls/handshake_client.go:763
	_go_fuzz_dep_.CoverTab[22541]++
								if serverProto == "" {
//line /usr/local/go/src/crypto/tls/handshake_client.go:764
		_go_fuzz_dep_.CoverTab[22545]++
									return nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:765
		// _ = "end of CoverTab[22545]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:766
		_go_fuzz_dep_.CoverTab[22546]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:766
		// _ = "end of CoverTab[22546]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:766
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:766
	// _ = "end of CoverTab[22541]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:766
	_go_fuzz_dep_.CoverTab[22542]++
								if len(clientProtos) == 0 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:767
		_go_fuzz_dep_.CoverTab[22547]++
									return errors.New("tls: server advertised unrequested ALPN extension")
//line /usr/local/go/src/crypto/tls/handshake_client.go:768
		// _ = "end of CoverTab[22547]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:769
		_go_fuzz_dep_.CoverTab[22548]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:769
		// _ = "end of CoverTab[22548]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:769
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:769
	// _ = "end of CoverTab[22542]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:769
	_go_fuzz_dep_.CoverTab[22543]++
								for _, proto := range clientProtos {
//line /usr/local/go/src/crypto/tls/handshake_client.go:770
		_go_fuzz_dep_.CoverTab[22549]++
									if proto == serverProto {
//line /usr/local/go/src/crypto/tls/handshake_client.go:771
			_go_fuzz_dep_.CoverTab[22550]++
										return nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:772
			// _ = "end of CoverTab[22550]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:773
			_go_fuzz_dep_.CoverTab[22551]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:773
			// _ = "end of CoverTab[22551]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:773
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:773
		// _ = "end of CoverTab[22549]"
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:774
	// _ = "end of CoverTab[22543]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:774
	_go_fuzz_dep_.CoverTab[22544]++
								return errors.New("tls: server selected unadvertised ALPN protocol")
//line /usr/local/go/src/crypto/tls/handshake_client.go:775
	// _ = "end of CoverTab[22544]"
}

func (hs *clientHandshakeState) readFinished(out []byte) error {
//line /usr/local/go/src/crypto/tls/handshake_client.go:778
	_go_fuzz_dep_.CoverTab[22552]++
								c := hs.c

								if err := c.readChangeCipherSpec(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:781
		_go_fuzz_dep_.CoverTab[22558]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:782
		// _ = "end of CoverTab[22558]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:783
		_go_fuzz_dep_.CoverTab[22559]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:783
		// _ = "end of CoverTab[22559]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:783
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:783
	// _ = "end of CoverTab[22552]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:783
	_go_fuzz_dep_.CoverTab[22553]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:788
	msg, err := c.readHandshake(nil)
	if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:789
		_go_fuzz_dep_.CoverTab[22560]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:790
		// _ = "end of CoverTab[22560]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:791
		_go_fuzz_dep_.CoverTab[22561]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:791
		// _ = "end of CoverTab[22561]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:791
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:791
	// _ = "end of CoverTab[22553]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:791
	_go_fuzz_dep_.CoverTab[22554]++
								serverFinished, ok := msg.(*finishedMsg)
								if !ok {
//line /usr/local/go/src/crypto/tls/handshake_client.go:793
		_go_fuzz_dep_.CoverTab[22562]++
									c.sendAlert(alertUnexpectedMessage)
									return unexpectedMessageError(serverFinished, msg)
//line /usr/local/go/src/crypto/tls/handshake_client.go:795
		// _ = "end of CoverTab[22562]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:796
		_go_fuzz_dep_.CoverTab[22563]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:796
		// _ = "end of CoverTab[22563]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:796
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:796
	// _ = "end of CoverTab[22554]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:796
	_go_fuzz_dep_.CoverTab[22555]++

								verify := hs.finishedHash.serverSum(hs.masterSecret)
								if len(verify) != len(serverFinished.verifyData) || func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:799
		_go_fuzz_dep_.CoverTab[22564]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:799
		return subtle.ConstantTimeCompare(verify, serverFinished.verifyData) != 1
									// _ = "end of CoverTab[22564]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:800
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:800
		_go_fuzz_dep_.CoverTab[22565]++
									c.sendAlert(alertHandshakeFailure)
									return errors.New("tls: server's Finished message was incorrect")
//line /usr/local/go/src/crypto/tls/handshake_client.go:802
		// _ = "end of CoverTab[22565]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:803
		_go_fuzz_dep_.CoverTab[22566]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:803
		// _ = "end of CoverTab[22566]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:803
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:803
	// _ = "end of CoverTab[22555]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:803
	_go_fuzz_dep_.CoverTab[22556]++

								if err := transcriptMsg(serverFinished, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:805
		_go_fuzz_dep_.CoverTab[22567]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:806
		// _ = "end of CoverTab[22567]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:807
		_go_fuzz_dep_.CoverTab[22568]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:807
		// _ = "end of CoverTab[22568]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:807
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:807
	// _ = "end of CoverTab[22556]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:807
	_go_fuzz_dep_.CoverTab[22557]++

								copy(out, verify)
								return nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:810
	// _ = "end of CoverTab[22557]"
}

func (hs *clientHandshakeState) readSessionTicket() error {
//line /usr/local/go/src/crypto/tls/handshake_client.go:813
	_go_fuzz_dep_.CoverTab[22569]++
								if !hs.serverHello.ticketSupported {
//line /usr/local/go/src/crypto/tls/handshake_client.go:814
		_go_fuzz_dep_.CoverTab[22573]++
									return nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:815
		// _ = "end of CoverTab[22573]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:816
		_go_fuzz_dep_.CoverTab[22574]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:816
		// _ = "end of CoverTab[22574]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:816
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:816
	// _ = "end of CoverTab[22569]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:816
	_go_fuzz_dep_.CoverTab[22570]++

								c := hs.c
								msg, err := c.readHandshake(&hs.finishedHash)
								if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:820
		_go_fuzz_dep_.CoverTab[22575]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:821
		// _ = "end of CoverTab[22575]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:822
		_go_fuzz_dep_.CoverTab[22576]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:822
		// _ = "end of CoverTab[22576]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:822
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:822
	// _ = "end of CoverTab[22570]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:822
	_go_fuzz_dep_.CoverTab[22571]++
								sessionTicketMsg, ok := msg.(*newSessionTicketMsg)
								if !ok {
//line /usr/local/go/src/crypto/tls/handshake_client.go:824
		_go_fuzz_dep_.CoverTab[22577]++
									c.sendAlert(alertUnexpectedMessage)
									return unexpectedMessageError(sessionTicketMsg, msg)
//line /usr/local/go/src/crypto/tls/handshake_client.go:826
		// _ = "end of CoverTab[22577]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:827
		_go_fuzz_dep_.CoverTab[22578]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:827
		// _ = "end of CoverTab[22578]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:827
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:827
	// _ = "end of CoverTab[22571]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:827
	_go_fuzz_dep_.CoverTab[22572]++

								hs.session = &ClientSessionState{
		sessionTicket:		sessionTicketMsg.ticket,
		vers:			c.vers,
		cipherSuite:		hs.suite.id,
		masterSecret:		hs.masterSecret,
		serverCertificates:	c.peerCertificates,
		verifiedChains:		c.verifiedChains,
		receivedAt:		c.config.time(),
		ocspResponse:		c.ocspResponse,
		scts:			c.scts,
	}

								return nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:841
	// _ = "end of CoverTab[22572]"
}

func (hs *clientHandshakeState) sendFinished(out []byte) error {
//line /usr/local/go/src/crypto/tls/handshake_client.go:844
	_go_fuzz_dep_.CoverTab[22579]++
								c := hs.c

								if err := c.writeChangeCipherRecord(); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:847
		_go_fuzz_dep_.CoverTab[22582]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:848
		// _ = "end of CoverTab[22582]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:849
		_go_fuzz_dep_.CoverTab[22583]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:849
		// _ = "end of CoverTab[22583]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:849
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:849
	// _ = "end of CoverTab[22579]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:849
	_go_fuzz_dep_.CoverTab[22580]++

								finished := new(finishedMsg)
								finished.verifyData = hs.finishedHash.clientSum(hs.masterSecret)
								if _, err := hs.c.writeHandshakeRecord(finished, &hs.finishedHash); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:853
		_go_fuzz_dep_.CoverTab[22584]++
									return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:854
		// _ = "end of CoverTab[22584]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:855
		_go_fuzz_dep_.CoverTab[22585]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:855
		// _ = "end of CoverTab[22585]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:855
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:855
	// _ = "end of CoverTab[22580]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:855
	_go_fuzz_dep_.CoverTab[22581]++
								copy(out, finished.verifyData)
								return nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:857
	// _ = "end of CoverTab[22581]"
}

// maxRSAKeySize is the maximum RSA key size in bits that we are willing
//line /usr/local/go/src/crypto/tls/handshake_client.go:860
// to verify the signatures of during a TLS handshake.
//line /usr/local/go/src/crypto/tls/handshake_client.go:862
const maxRSAKeySize = 8192

// verifyServerCertificate parses and verifies the provided chain, setting
//line /usr/local/go/src/crypto/tls/handshake_client.go:864
// c.verifiedChains and c.peerCertificates or sending the appropriate alert.
//line /usr/local/go/src/crypto/tls/handshake_client.go:866
func (c *Conn) verifyServerCertificate(certificates [][]byte) error {
//line /usr/local/go/src/crypto/tls/handshake_client.go:866
	_go_fuzz_dep_.CoverTab[22586]++
								activeHandles := make([]*activeCert, len(certificates))
								certs := make([]*x509.Certificate, len(certificates))
								for i, asn1Data := range certificates {
//line /usr/local/go/src/crypto/tls/handshake_client.go:869
		_go_fuzz_dep_.CoverTab[22592]++
									cert, err := clientCertCache.newCert(asn1Data)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:871
			_go_fuzz_dep_.CoverTab[22595]++
										c.sendAlert(alertBadCertificate)
										return errors.New("tls: failed to parse certificate from server: " + err.Error())
//line /usr/local/go/src/crypto/tls/handshake_client.go:873
			// _ = "end of CoverTab[22595]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:874
			_go_fuzz_dep_.CoverTab[22596]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:874
			// _ = "end of CoverTab[22596]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:874
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:874
		// _ = "end of CoverTab[22592]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:874
		_go_fuzz_dep_.CoverTab[22593]++
									if cert.cert.PublicKeyAlgorithm == x509.RSA && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:875
			_go_fuzz_dep_.CoverTab[22597]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:875
			return cert.cert.PublicKey.(*rsa.PublicKey).N.BitLen() > maxRSAKeySize
//line /usr/local/go/src/crypto/tls/handshake_client.go:875
			// _ = "end of CoverTab[22597]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:875
		}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:875
			_go_fuzz_dep_.CoverTab[22598]++
										c.sendAlert(alertBadCertificate)
										return fmt.Errorf("tls: server sent certificate containing RSA key larger than %d bits", maxRSAKeySize)
//line /usr/local/go/src/crypto/tls/handshake_client.go:877
			// _ = "end of CoverTab[22598]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:878
			_go_fuzz_dep_.CoverTab[22599]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:878
			// _ = "end of CoverTab[22599]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:878
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:878
		// _ = "end of CoverTab[22593]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:878
		_go_fuzz_dep_.CoverTab[22594]++
									activeHandles[i] = cert
									certs[i] = cert.cert
//line /usr/local/go/src/crypto/tls/handshake_client.go:880
		// _ = "end of CoverTab[22594]"
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:881
	// _ = "end of CoverTab[22586]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:881
	_go_fuzz_dep_.CoverTab[22587]++

								if !c.config.InsecureSkipVerify {
//line /usr/local/go/src/crypto/tls/handshake_client.go:883
		_go_fuzz_dep_.CoverTab[22600]++
									opts := x509.VerifyOptions{
			Roots:		c.config.RootCAs,
			CurrentTime:	c.config.time(),
			DNSName:	c.config.ServerName,
			Intermediates:	x509.NewCertPool(),
		}

		for _, cert := range certs[1:] {
//line /usr/local/go/src/crypto/tls/handshake_client.go:891
			_go_fuzz_dep_.CoverTab[22602]++
										opts.Intermediates.AddCert(cert)
//line /usr/local/go/src/crypto/tls/handshake_client.go:892
			// _ = "end of CoverTab[22602]"
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:893
		// _ = "end of CoverTab[22600]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:893
		_go_fuzz_dep_.CoverTab[22601]++
									var err error
									c.verifiedChains, err = certs[0].Verify(opts)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:896
			_go_fuzz_dep_.CoverTab[22603]++
										c.sendAlert(alertBadCertificate)
										return &CertificateVerificationError{UnverifiedCertificates: certs, Err: err}
//line /usr/local/go/src/crypto/tls/handshake_client.go:898
			// _ = "end of CoverTab[22603]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:899
			_go_fuzz_dep_.CoverTab[22604]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:899
			// _ = "end of CoverTab[22604]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:899
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:899
		// _ = "end of CoverTab[22601]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:900
		_go_fuzz_dep_.CoverTab[22605]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:900
		// _ = "end of CoverTab[22605]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:900
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:900
	// _ = "end of CoverTab[22587]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:900
	_go_fuzz_dep_.CoverTab[22588]++

								switch certs[0].PublicKey.(type) {
	case *rsa.PublicKey, *ecdsa.PublicKey, ed25519.PublicKey:
//line /usr/local/go/src/crypto/tls/handshake_client.go:903
		_go_fuzz_dep_.CoverTab[22606]++
									break
//line /usr/local/go/src/crypto/tls/handshake_client.go:904
		// _ = "end of CoverTab[22606]"
	default:
//line /usr/local/go/src/crypto/tls/handshake_client.go:905
		_go_fuzz_dep_.CoverTab[22607]++
									c.sendAlert(alertUnsupportedCertificate)
									return fmt.Errorf("tls: server's certificate contains an unsupported type of public key: %T", certs[0].PublicKey)
//line /usr/local/go/src/crypto/tls/handshake_client.go:907
		// _ = "end of CoverTab[22607]"
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:908
	// _ = "end of CoverTab[22588]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:908
	_go_fuzz_dep_.CoverTab[22589]++

								c.activeCertHandles = activeHandles
								c.peerCertificates = certs

								if c.config.VerifyPeerCertificate != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:913
		_go_fuzz_dep_.CoverTab[22608]++
									if err := c.config.VerifyPeerCertificate(certificates, c.verifiedChains); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:914
			_go_fuzz_dep_.CoverTab[22609]++
										c.sendAlert(alertBadCertificate)
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:916
			// _ = "end of CoverTab[22609]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:917
			_go_fuzz_dep_.CoverTab[22610]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:917
			// _ = "end of CoverTab[22610]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:917
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:917
		// _ = "end of CoverTab[22608]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:918
		_go_fuzz_dep_.CoverTab[22611]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:918
		// _ = "end of CoverTab[22611]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:918
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:918
	// _ = "end of CoverTab[22589]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:918
	_go_fuzz_dep_.CoverTab[22590]++

								if c.config.VerifyConnection != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:920
		_go_fuzz_dep_.CoverTab[22612]++
									if err := c.config.VerifyConnection(c.connectionStateLocked()); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:921
			_go_fuzz_dep_.CoverTab[22613]++
										c.sendAlert(alertBadCertificate)
										return err
//line /usr/local/go/src/crypto/tls/handshake_client.go:923
			// _ = "end of CoverTab[22613]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:924
			_go_fuzz_dep_.CoverTab[22614]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:924
			// _ = "end of CoverTab[22614]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:924
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:924
		// _ = "end of CoverTab[22612]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:925
		_go_fuzz_dep_.CoverTab[22615]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:925
		// _ = "end of CoverTab[22615]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:925
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:925
	// _ = "end of CoverTab[22590]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:925
	_go_fuzz_dep_.CoverTab[22591]++

								return nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:927
	// _ = "end of CoverTab[22591]"
}

// certificateRequestInfoFromMsg generates a CertificateRequestInfo from a TLS
//line /usr/local/go/src/crypto/tls/handshake_client.go:930
// <= 1.2 CertificateRequest, making an effort to fill in missing information.
//line /usr/local/go/src/crypto/tls/handshake_client.go:932
func certificateRequestInfoFromMsg(ctx context.Context, vers uint16, certReq *certificateRequestMsg) *CertificateRequestInfo {
//line /usr/local/go/src/crypto/tls/handshake_client.go:932
	_go_fuzz_dep_.CoverTab[22616]++
								cri := &CertificateRequestInfo{
		AcceptableCAs:	certReq.certificateAuthorities,
		Version:	vers,
		ctx:		ctx,
	}

	var rsaAvail, ecAvail bool
	for _, certType := range certReq.certificateTypes {
//line /usr/local/go/src/crypto/tls/handshake_client.go:940
		_go_fuzz_dep_.CoverTab[22620]++
									switch certType {
		case certTypeRSASign:
//line /usr/local/go/src/crypto/tls/handshake_client.go:942
			_go_fuzz_dep_.CoverTab[22621]++
										rsaAvail = true
//line /usr/local/go/src/crypto/tls/handshake_client.go:943
			// _ = "end of CoverTab[22621]"
		case certTypeECDSASign:
//line /usr/local/go/src/crypto/tls/handshake_client.go:944
			_go_fuzz_dep_.CoverTab[22622]++
										ecAvail = true
//line /usr/local/go/src/crypto/tls/handshake_client.go:945
			// _ = "end of CoverTab[22622]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:945
		default:
//line /usr/local/go/src/crypto/tls/handshake_client.go:945
			_go_fuzz_dep_.CoverTab[22623]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:945
			// _ = "end of CoverTab[22623]"
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:946
		// _ = "end of CoverTab[22620]"
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:947
	// _ = "end of CoverTab[22616]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:947
	_go_fuzz_dep_.CoverTab[22617]++

								if !certReq.hasSignatureAlgorithm {
//line /usr/local/go/src/crypto/tls/handshake_client.go:949
		_go_fuzz_dep_.CoverTab[22624]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:955
		switch {
		case rsaAvail && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:956
			_go_fuzz_dep_.CoverTab[22630]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:956
			return ecAvail
//line /usr/local/go/src/crypto/tls/handshake_client.go:956
			// _ = "end of CoverTab[22630]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:956
		}():
//line /usr/local/go/src/crypto/tls/handshake_client.go:956
			_go_fuzz_dep_.CoverTab[22626]++
										cri.SignatureSchemes = []SignatureScheme{
				ECDSAWithP256AndSHA256, ECDSAWithP384AndSHA384, ECDSAWithP521AndSHA512,
				PKCS1WithSHA256, PKCS1WithSHA384, PKCS1WithSHA512, PKCS1WithSHA1,
			}
//line /usr/local/go/src/crypto/tls/handshake_client.go:960
			// _ = "end of CoverTab[22626]"
		case rsaAvail:
//line /usr/local/go/src/crypto/tls/handshake_client.go:961
			_go_fuzz_dep_.CoverTab[22627]++
										cri.SignatureSchemes = []SignatureScheme{
				PKCS1WithSHA256, PKCS1WithSHA384, PKCS1WithSHA512, PKCS1WithSHA1,
			}
//line /usr/local/go/src/crypto/tls/handshake_client.go:964
			// _ = "end of CoverTab[22627]"
		case ecAvail:
//line /usr/local/go/src/crypto/tls/handshake_client.go:965
			_go_fuzz_dep_.CoverTab[22628]++
										cri.SignatureSchemes = []SignatureScheme{
				ECDSAWithP256AndSHA256, ECDSAWithP384AndSHA384, ECDSAWithP521AndSHA512,
			}
//line /usr/local/go/src/crypto/tls/handshake_client.go:968
			// _ = "end of CoverTab[22628]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:968
		default:
//line /usr/local/go/src/crypto/tls/handshake_client.go:968
			_go_fuzz_dep_.CoverTab[22629]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:968
			// _ = "end of CoverTab[22629]"
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:969
		// _ = "end of CoverTab[22624]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:969
		_go_fuzz_dep_.CoverTab[22625]++
									return cri
//line /usr/local/go/src/crypto/tls/handshake_client.go:970
		// _ = "end of CoverTab[22625]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:971
		_go_fuzz_dep_.CoverTab[22631]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:971
		// _ = "end of CoverTab[22631]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:971
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:971
	// _ = "end of CoverTab[22617]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:971
	_go_fuzz_dep_.CoverTab[22618]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:975
	cri.SignatureSchemes = make([]SignatureScheme, 0, len(certReq.supportedSignatureAlgorithms))
	for _, sigScheme := range certReq.supportedSignatureAlgorithms {
//line /usr/local/go/src/crypto/tls/handshake_client.go:976
		_go_fuzz_dep_.CoverTab[22632]++
									sigType, _, err := typeAndHashFromSignatureScheme(sigScheme)
									if err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:978
			_go_fuzz_dep_.CoverTab[22634]++
										continue
//line /usr/local/go/src/crypto/tls/handshake_client.go:979
			// _ = "end of CoverTab[22634]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:980
			_go_fuzz_dep_.CoverTab[22635]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:980
			// _ = "end of CoverTab[22635]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:980
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:980
		// _ = "end of CoverTab[22632]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:980
		_go_fuzz_dep_.CoverTab[22633]++
									switch sigType {
		case signatureECDSA, signatureEd25519:
//line /usr/local/go/src/crypto/tls/handshake_client.go:982
			_go_fuzz_dep_.CoverTab[22636]++
										if ecAvail {
//line /usr/local/go/src/crypto/tls/handshake_client.go:983
				_go_fuzz_dep_.CoverTab[22639]++
											cri.SignatureSchemes = append(cri.SignatureSchemes, sigScheme)
//line /usr/local/go/src/crypto/tls/handshake_client.go:984
				// _ = "end of CoverTab[22639]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:985
				_go_fuzz_dep_.CoverTab[22640]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:985
				// _ = "end of CoverTab[22640]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:985
			}
//line /usr/local/go/src/crypto/tls/handshake_client.go:985
			// _ = "end of CoverTab[22636]"
		case signatureRSAPSS, signaturePKCS1v15:
//line /usr/local/go/src/crypto/tls/handshake_client.go:986
			_go_fuzz_dep_.CoverTab[22637]++
										if rsaAvail {
//line /usr/local/go/src/crypto/tls/handshake_client.go:987
				_go_fuzz_dep_.CoverTab[22641]++
											cri.SignatureSchemes = append(cri.SignatureSchemes, sigScheme)
//line /usr/local/go/src/crypto/tls/handshake_client.go:988
				// _ = "end of CoverTab[22641]"
			} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:989
				_go_fuzz_dep_.CoverTab[22642]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:989
				// _ = "end of CoverTab[22642]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:989
			}
//line /usr/local/go/src/crypto/tls/handshake_client.go:989
			// _ = "end of CoverTab[22637]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:989
		default:
//line /usr/local/go/src/crypto/tls/handshake_client.go:989
			_go_fuzz_dep_.CoverTab[22638]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:989
			// _ = "end of CoverTab[22638]"
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:990
		// _ = "end of CoverTab[22633]"
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:991
	// _ = "end of CoverTab[22618]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:991
	_go_fuzz_dep_.CoverTab[22619]++

								return cri
//line /usr/local/go/src/crypto/tls/handshake_client.go:993
	// _ = "end of CoverTab[22619]"
}

func (c *Conn) getClientCertificate(cri *CertificateRequestInfo) (*Certificate, error) {
//line /usr/local/go/src/crypto/tls/handshake_client.go:996
	_go_fuzz_dep_.CoverTab[22643]++
								if c.config.GetClientCertificate != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:997
		_go_fuzz_dep_.CoverTab[22646]++
									return c.config.GetClientCertificate(cri)
//line /usr/local/go/src/crypto/tls/handshake_client.go:998
		// _ = "end of CoverTab[22646]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:999
		_go_fuzz_dep_.CoverTab[22647]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:999
		// _ = "end of CoverTab[22647]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:999
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:999
	// _ = "end of CoverTab[22643]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:999
	_go_fuzz_dep_.CoverTab[22644]++

								for _, chain := range c.config.Certificates {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1001
		_go_fuzz_dep_.CoverTab[22648]++
									if err := cri.SupportsCertificate(&chain); err != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1002
			_go_fuzz_dep_.CoverTab[22650]++
										continue
//line /usr/local/go/src/crypto/tls/handshake_client.go:1003
			// _ = "end of CoverTab[22650]"
		} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1004
			_go_fuzz_dep_.CoverTab[22651]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:1004
			// _ = "end of CoverTab[22651]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:1004
		}
//line /usr/local/go/src/crypto/tls/handshake_client.go:1004
		// _ = "end of CoverTab[22648]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:1004
		_go_fuzz_dep_.CoverTab[22649]++
									return &chain, nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:1005
		// _ = "end of CoverTab[22649]"
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:1006
	// _ = "end of CoverTab[22644]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:1006
	_go_fuzz_dep_.CoverTab[22645]++

//line /usr/local/go/src/crypto/tls/handshake_client.go:1009
	return new(Certificate), nil
//line /usr/local/go/src/crypto/tls/handshake_client.go:1009
	// _ = "end of CoverTab[22645]"
}

// clientSessionCacheKey returns a key used to cache sessionTickets that could
//line /usr/local/go/src/crypto/tls/handshake_client.go:1012
// be used to resume previously negotiated TLS sessions with a server.
//line /usr/local/go/src/crypto/tls/handshake_client.go:1014
func clientSessionCacheKey(serverAddr net.Addr, config *Config) string {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1014
	_go_fuzz_dep_.CoverTab[22652]++
								if len(config.ServerName) > 0 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1015
		_go_fuzz_dep_.CoverTab[22654]++
									return config.ServerName
//line /usr/local/go/src/crypto/tls/handshake_client.go:1016
		// _ = "end of CoverTab[22654]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1017
		_go_fuzz_dep_.CoverTab[22655]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:1017
		// _ = "end of CoverTab[22655]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:1017
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:1017
	// _ = "end of CoverTab[22652]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:1017
	_go_fuzz_dep_.CoverTab[22653]++
								return serverAddr.String()
//line /usr/local/go/src/crypto/tls/handshake_client.go:1018
	// _ = "end of CoverTab[22653]"
}

// hostnameInSNI converts name into an appropriate hostname for SNI.
//line /usr/local/go/src/crypto/tls/handshake_client.go:1021
// Literal IP addresses and absolute FQDNs are not permitted as SNI values.
//line /usr/local/go/src/crypto/tls/handshake_client.go:1021
// See RFC 6066, Section 3.
//line /usr/local/go/src/crypto/tls/handshake_client.go:1024
func hostnameInSNI(name string) string {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1024
	_go_fuzz_dep_.CoverTab[22656]++
								host := name
								if len(host) > 0 && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1026
		_go_fuzz_dep_.CoverTab[22661]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:1026
		return host[0] == '['
//line /usr/local/go/src/crypto/tls/handshake_client.go:1026
		// _ = "end of CoverTab[22661]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:1026
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1026
		_go_fuzz_dep_.CoverTab[22662]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:1026
		return host[len(host)-1] == ']'
//line /usr/local/go/src/crypto/tls/handshake_client.go:1026
		// _ = "end of CoverTab[22662]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:1026
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1026
		_go_fuzz_dep_.CoverTab[22663]++
									host = host[1 : len(host)-1]
//line /usr/local/go/src/crypto/tls/handshake_client.go:1027
		// _ = "end of CoverTab[22663]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1028
		_go_fuzz_dep_.CoverTab[22664]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:1028
		// _ = "end of CoverTab[22664]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:1028
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:1028
	// _ = "end of CoverTab[22656]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:1028
	_go_fuzz_dep_.CoverTab[22657]++
								if i := strings.LastIndex(host, "%"); i > 0 {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1029
		_go_fuzz_dep_.CoverTab[22665]++
									host = host[:i]
//line /usr/local/go/src/crypto/tls/handshake_client.go:1030
		// _ = "end of CoverTab[22665]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1031
		_go_fuzz_dep_.CoverTab[22666]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:1031
		// _ = "end of CoverTab[22666]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:1031
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:1031
	// _ = "end of CoverTab[22657]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:1031
	_go_fuzz_dep_.CoverTab[22658]++
								if net.ParseIP(host) != nil {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1032
		_go_fuzz_dep_.CoverTab[22667]++
									return ""
//line /usr/local/go/src/crypto/tls/handshake_client.go:1033
		// _ = "end of CoverTab[22667]"
	} else {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1034
		_go_fuzz_dep_.CoverTab[22668]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:1034
		// _ = "end of CoverTab[22668]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:1034
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:1034
	// _ = "end of CoverTab[22658]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:1034
	_go_fuzz_dep_.CoverTab[22659]++
								for len(name) > 0 && func() bool {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1035
		_go_fuzz_dep_.CoverTab[22669]++
//line /usr/local/go/src/crypto/tls/handshake_client.go:1035
		return name[len(name)-1] == '.'
//line /usr/local/go/src/crypto/tls/handshake_client.go:1035
		// _ = "end of CoverTab[22669]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:1035
	}() {
//line /usr/local/go/src/crypto/tls/handshake_client.go:1035
		_go_fuzz_dep_.CoverTab[22670]++
									name = name[:len(name)-1]
//line /usr/local/go/src/crypto/tls/handshake_client.go:1036
		// _ = "end of CoverTab[22670]"
	}
//line /usr/local/go/src/crypto/tls/handshake_client.go:1037
	// _ = "end of CoverTab[22659]"
//line /usr/local/go/src/crypto/tls/handshake_client.go:1037
	_go_fuzz_dep_.CoverTab[22660]++
								return name
//line /usr/local/go/src/crypto/tls/handshake_client.go:1038
	// _ = "end of CoverTab[22660]"
}

//line /usr/local/go/src/crypto/tls/handshake_client.go:1039
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/handshake_client.go:1039
var _ = _go_fuzz_dep_.CoverTab
