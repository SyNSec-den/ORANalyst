// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/tls/common.go:5
package tls

//line /usr/local/go/src/crypto/tls/common.go:5
import (
//line /usr/local/go/src/crypto/tls/common.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/common.go:5
)
//line /usr/local/go/src/crypto/tls/common.go:5
import (
//line /usr/local/go/src/crypto/tls/common.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/common.go:5
)

import (
	"bytes"
	"container/list"
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"
)

const (
							VersionTLS10	= 0x0301
							VersionTLS11	= 0x0302
							VersionTLS12	= 0x0303
							VersionTLS13	= 0x0304

//line /usr/local/go/src/crypto/tls/common.go:36
	VersionSSL30	= 0x0300
)

const (
	maxPlaintext		= 16384
	maxCiphertext		= 16384 + 2048
	maxCiphertextTLS13	= 16384 + 256
	recordHeaderLen		= 5
	maxHandshake		= 65536
	maxUselessRecords	= 16
)

//line /usr/local/go/src/crypto/tls/common.go:49
type recordType uint8

const (
	recordTypeChangeCipherSpec	recordType	= 20
	recordTypeAlert			recordType	= 21
	recordTypeHandshake		recordType	= 22
	recordTypeApplicationData	recordType	= 23
)

//line /usr/local/go/src/crypto/tls/common.go:59
const (
	typeHelloRequest	uint8	= 0
	typeClientHello		uint8	= 1
	typeServerHello		uint8	= 2
	typeNewSessionTicket	uint8	= 4
	typeEndOfEarlyData	uint8	= 5
	typeEncryptedExtensions	uint8	= 8
	typeCertificate		uint8	= 11
	typeServerKeyExchange	uint8	= 12
	typeCertificateRequest	uint8	= 13
	typeServerHelloDone	uint8	= 14
	typeCertificateVerify	uint8	= 15
	typeClientKeyExchange	uint8	= 16
	typeFinished		uint8	= 20
	typeCertificateStatus	uint8	= 22
	typeKeyUpdate		uint8	= 24
	typeNextProtocol	uint8	= 67
	typeMessageHash		uint8	= 254
)

//line /usr/local/go/src/crypto/tls/common.go:80
const (
	compressionNone uint8 = 0
)

//line /usr/local/go/src/crypto/tls/common.go:85
const (
	extensionServerName			uint16	= 0
	extensionStatusRequest			uint16	= 5
	extensionSupportedCurves		uint16	= 10
	extensionSupportedPoints		uint16	= 11
	extensionSignatureAlgorithms		uint16	= 13
	extensionALPN				uint16	= 16
	extensionSCT				uint16	= 18
	extensionSessionTicket			uint16	= 35
	extensionPreSharedKey			uint16	= 41
	extensionEarlyData			uint16	= 42
	extensionSupportedVersions		uint16	= 43
	extensionCookie				uint16	= 44
	extensionPSKModes			uint16	= 45
	extensionCertificateAuthorities		uint16	= 47
	extensionSignatureAlgorithmsCert	uint16	= 50
	extensionKeyShare			uint16	= 51
	extensionRenegotiationInfo		uint16	= 0xff01
)

//line /usr/local/go/src/crypto/tls/common.go:106
const (
	scsvRenegotiation uint16 = 0x00ff
)

//line /usr/local/go/src/crypto/tls/common.go:115
type CurveID uint16

const (
	CurveP256	CurveID	= 23
	CurveP384	CurveID	= 24
	CurveP521	CurveID	= 25
	X25519		CurveID	= 29
)

//line /usr/local/go/src/crypto/tls/common.go:125
type keyShare struct {
	group	CurveID
	data	[]byte
}

//line /usr/local/go/src/crypto/tls/common.go:131
const (
	pskModePlain	uint8	= 0
	pskModeDHE	uint8	= 1
)

//line /usr/local/go/src/crypto/tls/common.go:138
type pskIdentity struct {
	label			[]byte
	obfuscatedTicketAge	uint32
}

//line /usr/local/go/src/crypto/tls/common.go:145
const (
	pointFormatUncompressed uint8 = 0
)

//line /usr/local/go/src/crypto/tls/common.go:150
const (
	statusTypeOCSP uint8 = 1
)

//line /usr/local/go/src/crypto/tls/common.go:155
const (
	certTypeRSASign		= 1
	certTypeECDSASign	= 64
)

//line /usr/local/go/src/crypto/tls/common.go:162
const (
	signaturePKCS1v15	uint8	= iota + 225
	signatureRSAPSS
	signatureECDSA
	signatureEd25519
)

//line /usr/local/go/src/crypto/tls/common.go:172
var directSigning crypto.Hash = 0

//line /usr/local/go/src/crypto/tls/common.go:178
var defaultSupportedSignatureAlgorithms = []SignatureScheme{
	PSSWithSHA256,
	ECDSAWithP256AndSHA256,
	Ed25519,
	PSSWithSHA384,
	PSSWithSHA512,
	PKCS1WithSHA256,
	PKCS1WithSHA384,
	PKCS1WithSHA512,
	ECDSAWithP384AndSHA384,
	ECDSAWithP521AndSHA512,
	PKCS1WithSHA1,
	ECDSAWithSHA1,
}

//line /usr/local/go/src/crypto/tls/common.go:195
var helloRetryRequestRandom = []byte{
	0xCF, 0x21, 0xAD, 0x74, 0xE5, 0x9A, 0x61, 0x11,
	0xBE, 0x1D, 0x8C, 0x02, 0x1E, 0x65, 0xB8, 0x91,
	0xC2, 0xA2, 0x11, 0x16, 0x7A, 0xBB, 0x8C, 0x5E,
	0x07, 0x9E, 0x09, 0xE2, 0xC8, 0xA8, 0x33, 0x9C,
}

const (
//line /usr/local/go/src/crypto/tls/common.go:206
	downgradeCanaryTLS12	= "DOWNGRD\x01"
	downgradeCanaryTLS11	= "DOWNGRD\x00"
)

//line /usr/local/go/src/crypto/tls/common.go:212
var testingOnlyForceDowngradeCanary bool

//line /usr/local/go/src/crypto/tls/common.go:215
type ConnectionState struct {
//line /usr/local/go/src/crypto/tls/common.go:217
	Version	uint16

//line /usr/local/go/src/crypto/tls/common.go:220
	HandshakeComplete	bool

//line /usr/local/go/src/crypto/tls/common.go:224
	DidResume	bool

//line /usr/local/go/src/crypto/tls/common.go:228
	CipherSuite	uint16

//line /usr/local/go/src/crypto/tls/common.go:231
	NegotiatedProtocol	string

//line /usr/local/go/src/crypto/tls/common.go:236
	NegotiatedProtocolIsMutual	bool

//line /usr/local/go/src/crypto/tls/common.go:240
	ServerName	string

//line /usr/local/go/src/crypto/tls/common.go:251
	PeerCertificates	[]*x509.Certificate

//line /usr/local/go/src/crypto/tls/common.go:262
	VerifiedChains	[][]*x509.Certificate

//line /usr/local/go/src/crypto/tls/common.go:266
	SignedCertificateTimestamps	[][]byte

//line /usr/local/go/src/crypto/tls/common.go:270
	OCSPResponse	[]byte

//line /usr/local/go/src/crypto/tls/common.go:279
	TLSUnique	[]byte

//line /usr/local/go/src/crypto/tls/common.go:282
	ekm	func(label string, context []byte, length int) ([]byte, error)
}

//line /usr/local/go/src/crypto/tls/common.go:289
func (cs *ConnectionState) ExportKeyingMaterial(label string, context []byte, length int) ([]byte, error) {
//line /usr/local/go/src/crypto/tls/common.go:289
	_go_fuzz_dep_.CoverTab[21275]++
							return cs.ekm(label, context, length)
//line /usr/local/go/src/crypto/tls/common.go:290
	// _ = "end of CoverTab[21275]"
}

//line /usr/local/go/src/crypto/tls/common.go:295
type ClientAuthType int

const (
//line /usr/local/go/src/crypto/tls/common.go:301
	NoClientCert	ClientAuthType	= iota

//line /usr/local/go/src/crypto/tls/common.go:305
	RequestClientCert

//line /usr/local/go/src/crypto/tls/common.go:309
	RequireAnyClientCert

//line /usr/local/go/src/crypto/tls/common.go:314
	VerifyClientCertIfGiven

//line /usr/local/go/src/crypto/tls/common.go:318
	RequireAndVerifyClientCert
)

//line /usr/local/go/src/crypto/tls/common.go:323
func requiresClientCert(c ClientAuthType) bool {
//line /usr/local/go/src/crypto/tls/common.go:323
	_go_fuzz_dep_.CoverTab[21276]++
							switch c {
	case RequireAnyClientCert, RequireAndVerifyClientCert:
//line /usr/local/go/src/crypto/tls/common.go:325
		_go_fuzz_dep_.CoverTab[21277]++
								return true
//line /usr/local/go/src/crypto/tls/common.go:326
		// _ = "end of CoverTab[21277]"
	default:
//line /usr/local/go/src/crypto/tls/common.go:327
		_go_fuzz_dep_.CoverTab[21278]++
								return false
//line /usr/local/go/src/crypto/tls/common.go:328
		// _ = "end of CoverTab[21278]"
	}
//line /usr/local/go/src/crypto/tls/common.go:329
	// _ = "end of CoverTab[21276]"
}

//line /usr/local/go/src/crypto/tls/common.go:334
type ClientSessionState struct {
							sessionTicket		[]uint8
							vers			uint16
							cipherSuite		uint16
							masterSecret		[]byte
							serverCertificates	[]*x509.Certificate
							verifiedChains		[][]*x509.Certificate
							receivedAt		time.Time
							ocspResponse		[]byte
							scts			[][]byte

//line /usr/local/go/src/crypto/tls/common.go:346
	nonce	[]byte
	useBy	time.Time
	ageAdd	uint32
}

//line /usr/local/go/src/crypto/tls/common.go:357
type ClientSessionCache interface {
//line /usr/local/go/src/crypto/tls/common.go:360
	Get(sessionKey string) (session *ClientSessionState, ok bool)

//line /usr/local/go/src/crypto/tls/common.go:366
	Put(sessionKey string, cs *ClientSessionState)
}

//go:generate stringer -type=SignatureScheme,CurveID,ClientAuthType -output=common_string.go

//line /usr/local/go/src/crypto/tls/common.go:373
type SignatureScheme uint16

const (
//line /usr/local/go/src/crypto/tls/common.go:377
	PKCS1WithSHA256	SignatureScheme	= 0x0401
							PKCS1WithSHA384	SignatureScheme	= 0x0501
							PKCS1WithSHA512	SignatureScheme	= 0x0601

//line /usr/local/go/src/crypto/tls/common.go:382
	PSSWithSHA256	SignatureScheme	= 0x0804
							PSSWithSHA384	SignatureScheme	= 0x0805
							PSSWithSHA512	SignatureScheme	= 0x0806

//line /usr/local/go/src/crypto/tls/common.go:387
	ECDSAWithP256AndSHA256	SignatureScheme	= 0x0403
							ECDSAWithP384AndSHA384	SignatureScheme	= 0x0503
							ECDSAWithP521AndSHA512	SignatureScheme	= 0x0603

//line /usr/local/go/src/crypto/tls/common.go:392
	Ed25519	SignatureScheme	= 0x0807

//line /usr/local/go/src/crypto/tls/common.go:395
	PKCS1WithSHA1	SignatureScheme	= 0x0201
	ECDSAWithSHA1	SignatureScheme	= 0x0203
)

//line /usr/local/go/src/crypto/tls/common.go:401
type ClientHelloInfo struct {
//line /usr/local/go/src/crypto/tls/common.go:404
	CipherSuites	[]uint16

//line /usr/local/go/src/crypto/tls/common.go:409
	ServerName	string

//line /usr/local/go/src/crypto/tls/common.go:414
	SupportedCurves	[]CurveID

//line /usr/local/go/src/crypto/tls/common.go:419
	SupportedPoints	[]uint8

//line /usr/local/go/src/crypto/tls/common.go:424
	SignatureSchemes	[]SignatureScheme

//line /usr/local/go/src/crypto/tls/common.go:432
	SupportedProtos	[]string

//line /usr/local/go/src/crypto/tls/common.go:438
	SupportedVersions	[]uint16

//line /usr/local/go/src/crypto/tls/common.go:443
	Conn	net.Conn

//line /usr/local/go/src/crypto/tls/common.go:447
	config	*Config

//line /usr/local/go/src/crypto/tls/common.go:450
	ctx	context.Context
}

//line /usr/local/go/src/crypto/tls/common.go:456
func (c *ClientHelloInfo) Context() context.Context {
//line /usr/local/go/src/crypto/tls/common.go:456
	_go_fuzz_dep_.CoverTab[21279]++
							return c.ctx
//line /usr/local/go/src/crypto/tls/common.go:457
	// _ = "end of CoverTab[21279]"
}

//line /usr/local/go/src/crypto/tls/common.go:463
type CertificateRequestInfo struct {
//line /usr/local/go/src/crypto/tls/common.go:468
	AcceptableCAs	[][]byte

//line /usr/local/go/src/crypto/tls/common.go:472
	SignatureSchemes	[]SignatureScheme

//line /usr/local/go/src/crypto/tls/common.go:475
	Version	uint16

//line /usr/local/go/src/crypto/tls/common.go:478
	ctx	context.Context
}

//line /usr/local/go/src/crypto/tls/common.go:484
func (c *CertificateRequestInfo) Context() context.Context {
//line /usr/local/go/src/crypto/tls/common.go:484
	_go_fuzz_dep_.CoverTab[21280]++
							return c.ctx
//line /usr/local/go/src/crypto/tls/common.go:485
	// _ = "end of CoverTab[21280]"
}

//line /usr/local/go/src/crypto/tls/common.go:502
type RenegotiationSupport int

const (
//line /usr/local/go/src/crypto/tls/common.go:506
	RenegotiateNever	RenegotiationSupport	= iota

//line /usr/local/go/src/crypto/tls/common.go:510
	RenegotiateOnceAsClient

//line /usr/local/go/src/crypto/tls/common.go:514
	RenegotiateFreelyAsClient
)

//line /usr/local/go/src/crypto/tls/common.go:521
type Config struct {
//line /usr/local/go/src/crypto/tls/common.go:526
	Rand	io.Reader

//line /usr/local/go/src/crypto/tls/common.go:530
	Time	func() time.Time

//line /usr/local/go/src/crypto/tls/common.go:543
	Certificates	[]Certificate

//line /usr/local/go/src/crypto/tls/common.go:552
	NameToCertificate	map[string]*Certificate

//line /usr/local/go/src/crypto/tls/common.go:563
	GetCertificate	func(*ClientHelloInfo) (*Certificate, error)

//line /usr/local/go/src/crypto/tls/common.go:580
	GetClientCertificate	func(*CertificateRequestInfo) (*Certificate, error)

//line /usr/local/go/src/crypto/tls/common.go:595
	GetConfigForClient	func(*ClientHelloInfo) (*Config, error)

//line /usr/local/go/src/crypto/tls/common.go:610
	VerifyPeerCertificate	func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error

//line /usr/local/go/src/crypto/tls/common.go:620
	VerifyConnection	func(ConnectionState) error

//line /usr/local/go/src/crypto/tls/common.go:625
	RootCAs	*x509.CertPool

//line /usr/local/go/src/crypto/tls/common.go:633
	NextProtos	[]string

//line /usr/local/go/src/crypto/tls/common.go:639
	ServerName	string

//line /usr/local/go/src/crypto/tls/common.go:643
	ClientAuth	ClientAuthType

//line /usr/local/go/src/crypto/tls/common.go:648
	ClientCAs	*x509.CertPool

//line /usr/local/go/src/crypto/tls/common.go:656
	InsecureSkipVerify	bool

//line /usr/local/go/src/crypto/tls/common.go:663
	CipherSuites	[]uint16

//line /usr/local/go/src/crypto/tls/common.go:673
	PreferServerCipherSuites	bool

//line /usr/local/go/src/crypto/tls/common.go:678
	SessionTicketsDisabled	bool

//line /usr/local/go/src/crypto/tls/common.go:688
	SessionTicketKey	[32]byte

//line /usr/local/go/src/crypto/tls/common.go:692
	ClientSessionCache	ClientSessionCache

//line /usr/local/go/src/crypto/tls/common.go:704
	MinVersion	uint16

//line /usr/local/go/src/crypto/tls/common.go:710
	MaxVersion	uint16

//line /usr/local/go/src/crypto/tls/common.go:716
	CurvePreferences	[]CurveID

//line /usr/local/go/src/crypto/tls/common.go:722
	DynamicRecordSizingDisabled	bool

//line /usr/local/go/src/crypto/tls/common.go:726
	Renegotiation	RenegotiationSupport

//line /usr/local/go/src/crypto/tls/common.go:734
	KeyLogWriter	io.Writer

//line /usr/local/go/src/crypto/tls/common.go:737
	mutex	sync.RWMutex

//line /usr/local/go/src/crypto/tls/common.go:743
	sessionTicketKeys	[]ticketKey

//line /usr/local/go/src/crypto/tls/common.go:746
	autoSessionTicketKeys	[]ticketKey
}

const (
//line /usr/local/go/src/crypto/tls/common.go:752
	ticketKeyNameLen	= 16

//line /usr/local/go/src/crypto/tls/common.go:756
	ticketKeyLifetime	= 7 * 24 * time.Hour

//line /usr/local/go/src/crypto/tls/common.go:760
	ticketKeyRotation	= 24 * time.Hour
)

//line /usr/local/go/src/crypto/tls/common.go:764
type ticketKey struct {
//line /usr/local/go/src/crypto/tls/common.go:767
	keyName	[ticketKeyNameLen]byte
	aesKey	[16]byte
	hmacKey	[16]byte

	created	time.Time
}

//line /usr/local/go/src/crypto/tls/common.go:777
func (c *Config) ticketKeyFromBytes(b [32]byte) (key ticketKey) {
//line /usr/local/go/src/crypto/tls/common.go:777
	_go_fuzz_dep_.CoverTab[21281]++
							hashed := sha512.Sum512(b[:])
							copy(key.keyName[:], hashed[:ticketKeyNameLen])
							copy(key.aesKey[:], hashed[ticketKeyNameLen:ticketKeyNameLen+16])
							copy(key.hmacKey[:], hashed[ticketKeyNameLen+16:ticketKeyNameLen+32])
							key.created = c.time()
							return key
//line /usr/local/go/src/crypto/tls/common.go:783
	// _ = "end of CoverTab[21281]"
}

//line /usr/local/go/src/crypto/tls/common.go:788
const maxSessionTicketLifetime = 7 * 24 * time.Hour

//line /usr/local/go/src/crypto/tls/common.go:792
func (c *Config) Clone() *Config {
//line /usr/local/go/src/crypto/tls/common.go:792
	_go_fuzz_dep_.CoverTab[21282]++
							if c == nil {
//line /usr/local/go/src/crypto/tls/common.go:793
		_go_fuzz_dep_.CoverTab[21284]++
								return nil
//line /usr/local/go/src/crypto/tls/common.go:794
		// _ = "end of CoverTab[21284]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:795
		_go_fuzz_dep_.CoverTab[21285]++
//line /usr/local/go/src/crypto/tls/common.go:795
		// _ = "end of CoverTab[21285]"
//line /usr/local/go/src/crypto/tls/common.go:795
	}
//line /usr/local/go/src/crypto/tls/common.go:795
	// _ = "end of CoverTab[21282]"
//line /usr/local/go/src/crypto/tls/common.go:795
	_go_fuzz_dep_.CoverTab[21283]++
							c.mutex.RLock()
							defer c.mutex.RUnlock()
							return &Config{
		Rand:				c.Rand,
		Time:				c.Time,
		Certificates:			c.Certificates,
		NameToCertificate:		c.NameToCertificate,
		GetCertificate:			c.GetCertificate,
		GetClientCertificate:		c.GetClientCertificate,
		GetConfigForClient:		c.GetConfigForClient,
		VerifyPeerCertificate:		c.VerifyPeerCertificate,
		VerifyConnection:		c.VerifyConnection,
		RootCAs:			c.RootCAs,
		NextProtos:			c.NextProtos,
		ServerName:			c.ServerName,
		ClientAuth:			c.ClientAuth,
		ClientCAs:			c.ClientCAs,
		InsecureSkipVerify:		c.InsecureSkipVerify,
		CipherSuites:			c.CipherSuites,
		PreferServerCipherSuites:	c.PreferServerCipherSuites,
		SessionTicketsDisabled:		c.SessionTicketsDisabled,
		SessionTicketKey:		c.SessionTicketKey,
		ClientSessionCache:		c.ClientSessionCache,
		MinVersion:			c.MinVersion,
		MaxVersion:			c.MaxVersion,
		CurvePreferences:		c.CurvePreferences,
		DynamicRecordSizingDisabled:	c.DynamicRecordSizingDisabled,
		Renegotiation:			c.Renegotiation,
		KeyLogWriter:			c.KeyLogWriter,
		sessionTicketKeys:		c.sessionTicketKeys,
		autoSessionTicketKeys:		c.autoSessionTicketKeys,
	}
//line /usr/local/go/src/crypto/tls/common.go:827
	// _ = "end of CoverTab[21283]"
}

//line /usr/local/go/src/crypto/tls/common.go:832
var deprecatedSessionTicketKey = []byte("DEPRECATED")

//line /usr/local/go/src/crypto/tls/common.go:836
func (c *Config) initLegacySessionTicketKeyRLocked() {
//line /usr/local/go/src/crypto/tls/common.go:836
	_go_fuzz_dep_.CoverTab[21286]++

//line /usr/local/go/src/crypto/tls/common.go:839
	if c.SessionTicketKey != [32]byte{} && func() bool {
//line /usr/local/go/src/crypto/tls/common.go:839
		_go_fuzz_dep_.CoverTab[21288]++
//line /usr/local/go/src/crypto/tls/common.go:839
		return (bytes.HasPrefix(c.SessionTicketKey[:], deprecatedSessionTicketKey) || func() bool {
									_go_fuzz_dep_.CoverTab[21289]++
//line /usr/local/go/src/crypto/tls/common.go:840
			return len(c.sessionTicketKeys) > 0
//line /usr/local/go/src/crypto/tls/common.go:840
			// _ = "end of CoverTab[21289]"
//line /usr/local/go/src/crypto/tls/common.go:840
		}())
//line /usr/local/go/src/crypto/tls/common.go:840
		// _ = "end of CoverTab[21288]"
//line /usr/local/go/src/crypto/tls/common.go:840
	}() {
//line /usr/local/go/src/crypto/tls/common.go:840
		_go_fuzz_dep_.CoverTab[21290]++
								return
//line /usr/local/go/src/crypto/tls/common.go:841
		// _ = "end of CoverTab[21290]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:842
		_go_fuzz_dep_.CoverTab[21291]++
//line /usr/local/go/src/crypto/tls/common.go:842
		// _ = "end of CoverTab[21291]"
//line /usr/local/go/src/crypto/tls/common.go:842
	}
//line /usr/local/go/src/crypto/tls/common.go:842
	// _ = "end of CoverTab[21286]"
//line /usr/local/go/src/crypto/tls/common.go:842
	_go_fuzz_dep_.CoverTab[21287]++

//line /usr/local/go/src/crypto/tls/common.go:845
	c.mutex.RUnlock()
	defer c.mutex.RLock()
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.SessionTicketKey == [32]byte{} {
//line /usr/local/go/src/crypto/tls/common.go:849
		_go_fuzz_dep_.CoverTab[21292]++
								if _, err := io.ReadFull(c.rand(), c.SessionTicketKey[:]); err != nil {
//line /usr/local/go/src/crypto/tls/common.go:850
			_go_fuzz_dep_.CoverTab[21294]++
									panic(fmt.Sprintf("tls: unable to generate random session ticket key: %v", err))
//line /usr/local/go/src/crypto/tls/common.go:851
			// _ = "end of CoverTab[21294]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:852
			_go_fuzz_dep_.CoverTab[21295]++
//line /usr/local/go/src/crypto/tls/common.go:852
			// _ = "end of CoverTab[21295]"
//line /usr/local/go/src/crypto/tls/common.go:852
		}
//line /usr/local/go/src/crypto/tls/common.go:852
		// _ = "end of CoverTab[21292]"
//line /usr/local/go/src/crypto/tls/common.go:852
		_go_fuzz_dep_.CoverTab[21293]++

//line /usr/local/go/src/crypto/tls/common.go:857
		copy(c.SessionTicketKey[:], deprecatedSessionTicketKey)
//line /usr/local/go/src/crypto/tls/common.go:857
		// _ = "end of CoverTab[21293]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:858
		_go_fuzz_dep_.CoverTab[21296]++
//line /usr/local/go/src/crypto/tls/common.go:858
		if !bytes.HasPrefix(c.SessionTicketKey[:], deprecatedSessionTicketKey) && func() bool {
//line /usr/local/go/src/crypto/tls/common.go:858
			_go_fuzz_dep_.CoverTab[21297]++
//line /usr/local/go/src/crypto/tls/common.go:858
			return len(c.sessionTicketKeys) == 0
//line /usr/local/go/src/crypto/tls/common.go:858
			// _ = "end of CoverTab[21297]"
//line /usr/local/go/src/crypto/tls/common.go:858
		}() {
//line /usr/local/go/src/crypto/tls/common.go:858
			_go_fuzz_dep_.CoverTab[21298]++
									c.sessionTicketKeys = []ticketKey{c.ticketKeyFromBytes(c.SessionTicketKey)}
//line /usr/local/go/src/crypto/tls/common.go:859
			// _ = "end of CoverTab[21298]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:860
			_go_fuzz_dep_.CoverTab[21299]++
//line /usr/local/go/src/crypto/tls/common.go:860
			// _ = "end of CoverTab[21299]"
//line /usr/local/go/src/crypto/tls/common.go:860
		}
//line /usr/local/go/src/crypto/tls/common.go:860
		// _ = "end of CoverTab[21296]"
//line /usr/local/go/src/crypto/tls/common.go:860
	}
//line /usr/local/go/src/crypto/tls/common.go:860
	// _ = "end of CoverTab[21287]"

}

//line /usr/local/go/src/crypto/tls/common.go:873
func (c *Config) ticketKeys(configForClient *Config) []ticketKey {
//line /usr/local/go/src/crypto/tls/common.go:873
	_go_fuzz_dep_.CoverTab[21300]++

//line /usr/local/go/src/crypto/tls/common.go:876
	if configForClient != nil {
//line /usr/local/go/src/crypto/tls/common.go:876
		_go_fuzz_dep_.CoverTab[21306]++
								configForClient.mutex.RLock()
								if configForClient.SessionTicketsDisabled {
//line /usr/local/go/src/crypto/tls/common.go:878
			_go_fuzz_dep_.CoverTab[21309]++
									return nil
//line /usr/local/go/src/crypto/tls/common.go:879
			// _ = "end of CoverTab[21309]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:880
			_go_fuzz_dep_.CoverTab[21310]++
//line /usr/local/go/src/crypto/tls/common.go:880
			// _ = "end of CoverTab[21310]"
//line /usr/local/go/src/crypto/tls/common.go:880
		}
//line /usr/local/go/src/crypto/tls/common.go:880
		// _ = "end of CoverTab[21306]"
//line /usr/local/go/src/crypto/tls/common.go:880
		_go_fuzz_dep_.CoverTab[21307]++
								configForClient.initLegacySessionTicketKeyRLocked()
								if len(configForClient.sessionTicketKeys) != 0 {
//line /usr/local/go/src/crypto/tls/common.go:882
			_go_fuzz_dep_.CoverTab[21311]++
									ret := configForClient.sessionTicketKeys
									configForClient.mutex.RUnlock()
									return ret
//line /usr/local/go/src/crypto/tls/common.go:885
			// _ = "end of CoverTab[21311]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:886
			_go_fuzz_dep_.CoverTab[21312]++
//line /usr/local/go/src/crypto/tls/common.go:886
			// _ = "end of CoverTab[21312]"
//line /usr/local/go/src/crypto/tls/common.go:886
		}
//line /usr/local/go/src/crypto/tls/common.go:886
		// _ = "end of CoverTab[21307]"
//line /usr/local/go/src/crypto/tls/common.go:886
		_go_fuzz_dep_.CoverTab[21308]++
								configForClient.mutex.RUnlock()
//line /usr/local/go/src/crypto/tls/common.go:887
		// _ = "end of CoverTab[21308]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:888
		_go_fuzz_dep_.CoverTab[21313]++
//line /usr/local/go/src/crypto/tls/common.go:888
		// _ = "end of CoverTab[21313]"
//line /usr/local/go/src/crypto/tls/common.go:888
	}
//line /usr/local/go/src/crypto/tls/common.go:888
	// _ = "end of CoverTab[21300]"
//line /usr/local/go/src/crypto/tls/common.go:888
	_go_fuzz_dep_.CoverTab[21301]++

							c.mutex.RLock()
							defer c.mutex.RUnlock()
							if c.SessionTicketsDisabled {
//line /usr/local/go/src/crypto/tls/common.go:892
		_go_fuzz_dep_.CoverTab[21314]++
								return nil
//line /usr/local/go/src/crypto/tls/common.go:893
		// _ = "end of CoverTab[21314]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:894
		_go_fuzz_dep_.CoverTab[21315]++
//line /usr/local/go/src/crypto/tls/common.go:894
		// _ = "end of CoverTab[21315]"
//line /usr/local/go/src/crypto/tls/common.go:894
	}
//line /usr/local/go/src/crypto/tls/common.go:894
	// _ = "end of CoverTab[21301]"
//line /usr/local/go/src/crypto/tls/common.go:894
	_go_fuzz_dep_.CoverTab[21302]++
							c.initLegacySessionTicketKeyRLocked()
							if len(c.sessionTicketKeys) != 0 {
//line /usr/local/go/src/crypto/tls/common.go:896
		_go_fuzz_dep_.CoverTab[21316]++
								return c.sessionTicketKeys
//line /usr/local/go/src/crypto/tls/common.go:897
		// _ = "end of CoverTab[21316]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:898
		_go_fuzz_dep_.CoverTab[21317]++
//line /usr/local/go/src/crypto/tls/common.go:898
		// _ = "end of CoverTab[21317]"
//line /usr/local/go/src/crypto/tls/common.go:898
	}
//line /usr/local/go/src/crypto/tls/common.go:898
	// _ = "end of CoverTab[21302]"
//line /usr/local/go/src/crypto/tls/common.go:898
	_go_fuzz_dep_.CoverTab[21303]++

							if len(c.autoSessionTicketKeys) > 0 && func() bool {
//line /usr/local/go/src/crypto/tls/common.go:900
		_go_fuzz_dep_.CoverTab[21318]++
//line /usr/local/go/src/crypto/tls/common.go:900
		return c.time().Sub(c.autoSessionTicketKeys[0].created) < ticketKeyRotation
//line /usr/local/go/src/crypto/tls/common.go:900
		// _ = "end of CoverTab[21318]"
//line /usr/local/go/src/crypto/tls/common.go:900
	}() {
//line /usr/local/go/src/crypto/tls/common.go:900
		_go_fuzz_dep_.CoverTab[21319]++
								return c.autoSessionTicketKeys
//line /usr/local/go/src/crypto/tls/common.go:901
		// _ = "end of CoverTab[21319]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:902
		_go_fuzz_dep_.CoverTab[21320]++
//line /usr/local/go/src/crypto/tls/common.go:902
		// _ = "end of CoverTab[21320]"
//line /usr/local/go/src/crypto/tls/common.go:902
	}
//line /usr/local/go/src/crypto/tls/common.go:902
	// _ = "end of CoverTab[21303]"
//line /usr/local/go/src/crypto/tls/common.go:902
	_go_fuzz_dep_.CoverTab[21304]++

//line /usr/local/go/src/crypto/tls/common.go:905
	c.mutex.RUnlock()
	defer c.mutex.RLock()
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.autoSessionTicketKeys) == 0 || func() bool {
//line /usr/local/go/src/crypto/tls/common.go:910
		_go_fuzz_dep_.CoverTab[21321]++
//line /usr/local/go/src/crypto/tls/common.go:910
		return c.time().Sub(c.autoSessionTicketKeys[0].created) >= ticketKeyRotation
//line /usr/local/go/src/crypto/tls/common.go:910
		// _ = "end of CoverTab[21321]"
//line /usr/local/go/src/crypto/tls/common.go:910
	}() {
//line /usr/local/go/src/crypto/tls/common.go:910
		_go_fuzz_dep_.CoverTab[21322]++
								var newKey [32]byte
								if _, err := io.ReadFull(c.rand(), newKey[:]); err != nil {
//line /usr/local/go/src/crypto/tls/common.go:912
			_go_fuzz_dep_.CoverTab[21325]++
									panic(fmt.Sprintf("unable to generate random session ticket key: %v", err))
//line /usr/local/go/src/crypto/tls/common.go:913
			// _ = "end of CoverTab[21325]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:914
			_go_fuzz_dep_.CoverTab[21326]++
//line /usr/local/go/src/crypto/tls/common.go:914
			// _ = "end of CoverTab[21326]"
//line /usr/local/go/src/crypto/tls/common.go:914
		}
//line /usr/local/go/src/crypto/tls/common.go:914
		// _ = "end of CoverTab[21322]"
//line /usr/local/go/src/crypto/tls/common.go:914
		_go_fuzz_dep_.CoverTab[21323]++
								valid := make([]ticketKey, 0, len(c.autoSessionTicketKeys)+1)
								valid = append(valid, c.ticketKeyFromBytes(newKey))
								for _, k := range c.autoSessionTicketKeys {
//line /usr/local/go/src/crypto/tls/common.go:917
			_go_fuzz_dep_.CoverTab[21327]++

									if c.time().Sub(k.created) < ticketKeyLifetime {
//line /usr/local/go/src/crypto/tls/common.go:919
				_go_fuzz_dep_.CoverTab[21328]++
										valid = append(valid, k)
//line /usr/local/go/src/crypto/tls/common.go:920
				// _ = "end of CoverTab[21328]"
			} else {
//line /usr/local/go/src/crypto/tls/common.go:921
				_go_fuzz_dep_.CoverTab[21329]++
//line /usr/local/go/src/crypto/tls/common.go:921
				// _ = "end of CoverTab[21329]"
//line /usr/local/go/src/crypto/tls/common.go:921
			}
//line /usr/local/go/src/crypto/tls/common.go:921
			// _ = "end of CoverTab[21327]"
		}
//line /usr/local/go/src/crypto/tls/common.go:922
		// _ = "end of CoverTab[21323]"
//line /usr/local/go/src/crypto/tls/common.go:922
		_go_fuzz_dep_.CoverTab[21324]++
								c.autoSessionTicketKeys = valid
//line /usr/local/go/src/crypto/tls/common.go:923
		// _ = "end of CoverTab[21324]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:924
		_go_fuzz_dep_.CoverTab[21330]++
//line /usr/local/go/src/crypto/tls/common.go:924
		// _ = "end of CoverTab[21330]"
//line /usr/local/go/src/crypto/tls/common.go:924
	}
//line /usr/local/go/src/crypto/tls/common.go:924
	// _ = "end of CoverTab[21304]"
//line /usr/local/go/src/crypto/tls/common.go:924
	_go_fuzz_dep_.CoverTab[21305]++
							return c.autoSessionTicketKeys
//line /usr/local/go/src/crypto/tls/common.go:925
	// _ = "end of CoverTab[21305]"
}

//line /usr/local/go/src/crypto/tls/common.go:941
func (c *Config) SetSessionTicketKeys(keys [][32]byte) {
//line /usr/local/go/src/crypto/tls/common.go:941
	_go_fuzz_dep_.CoverTab[21331]++
							if len(keys) == 0 {
//line /usr/local/go/src/crypto/tls/common.go:942
		_go_fuzz_dep_.CoverTab[21334]++
								panic("tls: keys must have at least one key")
//line /usr/local/go/src/crypto/tls/common.go:943
		// _ = "end of CoverTab[21334]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:944
		_go_fuzz_dep_.CoverTab[21335]++
//line /usr/local/go/src/crypto/tls/common.go:944
		// _ = "end of CoverTab[21335]"
//line /usr/local/go/src/crypto/tls/common.go:944
	}
//line /usr/local/go/src/crypto/tls/common.go:944
	// _ = "end of CoverTab[21331]"
//line /usr/local/go/src/crypto/tls/common.go:944
	_go_fuzz_dep_.CoverTab[21332]++

							newKeys := make([]ticketKey, len(keys))
							for i, bytes := range keys {
//line /usr/local/go/src/crypto/tls/common.go:947
		_go_fuzz_dep_.CoverTab[21336]++
								newKeys[i] = c.ticketKeyFromBytes(bytes)
//line /usr/local/go/src/crypto/tls/common.go:948
		// _ = "end of CoverTab[21336]"
	}
//line /usr/local/go/src/crypto/tls/common.go:949
	// _ = "end of CoverTab[21332]"
//line /usr/local/go/src/crypto/tls/common.go:949
	_go_fuzz_dep_.CoverTab[21333]++

							c.mutex.Lock()
							c.sessionTicketKeys = newKeys
							c.mutex.Unlock()
//line /usr/local/go/src/crypto/tls/common.go:953
	// _ = "end of CoverTab[21333]"
}

func (c *Config) rand() io.Reader {
//line /usr/local/go/src/crypto/tls/common.go:956
	_go_fuzz_dep_.CoverTab[21337]++
							r := c.Rand
							if r == nil {
//line /usr/local/go/src/crypto/tls/common.go:958
		_go_fuzz_dep_.CoverTab[21339]++
								return rand.Reader
//line /usr/local/go/src/crypto/tls/common.go:959
		// _ = "end of CoverTab[21339]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:960
		_go_fuzz_dep_.CoverTab[21340]++
//line /usr/local/go/src/crypto/tls/common.go:960
		// _ = "end of CoverTab[21340]"
//line /usr/local/go/src/crypto/tls/common.go:960
	}
//line /usr/local/go/src/crypto/tls/common.go:960
	// _ = "end of CoverTab[21337]"
//line /usr/local/go/src/crypto/tls/common.go:960
	_go_fuzz_dep_.CoverTab[21338]++
							return r
//line /usr/local/go/src/crypto/tls/common.go:961
	// _ = "end of CoverTab[21338]"
}

func (c *Config) time() time.Time {
//line /usr/local/go/src/crypto/tls/common.go:964
	_go_fuzz_dep_.CoverTab[21341]++
							t := c.Time
							if t == nil {
//line /usr/local/go/src/crypto/tls/common.go:966
		_go_fuzz_dep_.CoverTab[21343]++
								t = time.Now
//line /usr/local/go/src/crypto/tls/common.go:967
		// _ = "end of CoverTab[21343]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:968
		_go_fuzz_dep_.CoverTab[21344]++
//line /usr/local/go/src/crypto/tls/common.go:968
		// _ = "end of CoverTab[21344]"
//line /usr/local/go/src/crypto/tls/common.go:968
	}
//line /usr/local/go/src/crypto/tls/common.go:968
	// _ = "end of CoverTab[21341]"
//line /usr/local/go/src/crypto/tls/common.go:968
	_go_fuzz_dep_.CoverTab[21342]++
							return t()
//line /usr/local/go/src/crypto/tls/common.go:969
	// _ = "end of CoverTab[21342]"
}

func (c *Config) cipherSuites() []uint16 {
//line /usr/local/go/src/crypto/tls/common.go:972
	_go_fuzz_dep_.CoverTab[21345]++
							if needFIPS() {
//line /usr/local/go/src/crypto/tls/common.go:973
		_go_fuzz_dep_.CoverTab[21348]++
								return fipsCipherSuites(c)
//line /usr/local/go/src/crypto/tls/common.go:974
		// _ = "end of CoverTab[21348]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:975
		_go_fuzz_dep_.CoverTab[21349]++
//line /usr/local/go/src/crypto/tls/common.go:975
		// _ = "end of CoverTab[21349]"
//line /usr/local/go/src/crypto/tls/common.go:975
	}
//line /usr/local/go/src/crypto/tls/common.go:975
	// _ = "end of CoverTab[21345]"
//line /usr/local/go/src/crypto/tls/common.go:975
	_go_fuzz_dep_.CoverTab[21346]++
							if c.CipherSuites != nil {
//line /usr/local/go/src/crypto/tls/common.go:976
		_go_fuzz_dep_.CoverTab[21350]++
								return c.CipherSuites
//line /usr/local/go/src/crypto/tls/common.go:977
		// _ = "end of CoverTab[21350]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:978
		_go_fuzz_dep_.CoverTab[21351]++
//line /usr/local/go/src/crypto/tls/common.go:978
		// _ = "end of CoverTab[21351]"
//line /usr/local/go/src/crypto/tls/common.go:978
	}
//line /usr/local/go/src/crypto/tls/common.go:978
	// _ = "end of CoverTab[21346]"
//line /usr/local/go/src/crypto/tls/common.go:978
	_go_fuzz_dep_.CoverTab[21347]++
							return defaultCipherSuites
//line /usr/local/go/src/crypto/tls/common.go:979
	// _ = "end of CoverTab[21347]"
}

var supportedVersions = []uint16{
	VersionTLS13,
	VersionTLS12,
	VersionTLS11,
	VersionTLS10,
}

//line /usr/local/go/src/crypto/tls/common.go:991
const roleClient = true
const roleServer = false

func (c *Config) supportedVersions(isClient bool) []uint16 {
//line /usr/local/go/src/crypto/tls/common.go:994
	_go_fuzz_dep_.CoverTab[21352]++
							versions := make([]uint16, 0, len(supportedVersions))
							for _, v := range supportedVersions {
//line /usr/local/go/src/crypto/tls/common.go:996
		_go_fuzz_dep_.CoverTab[21354]++
								if needFIPS() && func() bool {
//line /usr/local/go/src/crypto/tls/common.go:997
			_go_fuzz_dep_.CoverTab[21359]++
//line /usr/local/go/src/crypto/tls/common.go:997
			return (v < fipsMinVersion(c) || func() bool {
//line /usr/local/go/src/crypto/tls/common.go:997
				_go_fuzz_dep_.CoverTab[21360]++
//line /usr/local/go/src/crypto/tls/common.go:997
				return v > fipsMaxVersion(c)
//line /usr/local/go/src/crypto/tls/common.go:997
				// _ = "end of CoverTab[21360]"
//line /usr/local/go/src/crypto/tls/common.go:997
			}())
//line /usr/local/go/src/crypto/tls/common.go:997
			// _ = "end of CoverTab[21359]"
//line /usr/local/go/src/crypto/tls/common.go:997
		}() {
//line /usr/local/go/src/crypto/tls/common.go:997
			_go_fuzz_dep_.CoverTab[21361]++
									continue
//line /usr/local/go/src/crypto/tls/common.go:998
			// _ = "end of CoverTab[21361]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:999
			_go_fuzz_dep_.CoverTab[21362]++
//line /usr/local/go/src/crypto/tls/common.go:999
			// _ = "end of CoverTab[21362]"
//line /usr/local/go/src/crypto/tls/common.go:999
		}
//line /usr/local/go/src/crypto/tls/common.go:999
		// _ = "end of CoverTab[21354]"
//line /usr/local/go/src/crypto/tls/common.go:999
		_go_fuzz_dep_.CoverTab[21355]++
								if (c == nil || func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1000
			_go_fuzz_dep_.CoverTab[21363]++
//line /usr/local/go/src/crypto/tls/common.go:1000
			return c.MinVersion == 0
//line /usr/local/go/src/crypto/tls/common.go:1000
			// _ = "end of CoverTab[21363]"
//line /usr/local/go/src/crypto/tls/common.go:1000
		}()) && func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1000
			_go_fuzz_dep_.CoverTab[21364]++
//line /usr/local/go/src/crypto/tls/common.go:1000
			return isClient
									// _ = "end of CoverTab[21364]"
//line /usr/local/go/src/crypto/tls/common.go:1001
		}() && func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1001
			_go_fuzz_dep_.CoverTab[21365]++
//line /usr/local/go/src/crypto/tls/common.go:1001
			return v < VersionTLS12
//line /usr/local/go/src/crypto/tls/common.go:1001
			// _ = "end of CoverTab[21365]"
//line /usr/local/go/src/crypto/tls/common.go:1001
		}() {
//line /usr/local/go/src/crypto/tls/common.go:1001
			_go_fuzz_dep_.CoverTab[21366]++
									continue
//line /usr/local/go/src/crypto/tls/common.go:1002
			// _ = "end of CoverTab[21366]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1003
			_go_fuzz_dep_.CoverTab[21367]++
//line /usr/local/go/src/crypto/tls/common.go:1003
			// _ = "end of CoverTab[21367]"
//line /usr/local/go/src/crypto/tls/common.go:1003
		}
//line /usr/local/go/src/crypto/tls/common.go:1003
		// _ = "end of CoverTab[21355]"
//line /usr/local/go/src/crypto/tls/common.go:1003
		_go_fuzz_dep_.CoverTab[21356]++
								if c != nil && func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1004
			_go_fuzz_dep_.CoverTab[21368]++
//line /usr/local/go/src/crypto/tls/common.go:1004
			return c.MinVersion != 0
//line /usr/local/go/src/crypto/tls/common.go:1004
			// _ = "end of CoverTab[21368]"
//line /usr/local/go/src/crypto/tls/common.go:1004
		}() && func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1004
			_go_fuzz_dep_.CoverTab[21369]++
//line /usr/local/go/src/crypto/tls/common.go:1004
			return v < c.MinVersion
//line /usr/local/go/src/crypto/tls/common.go:1004
			// _ = "end of CoverTab[21369]"
//line /usr/local/go/src/crypto/tls/common.go:1004
		}() {
//line /usr/local/go/src/crypto/tls/common.go:1004
			_go_fuzz_dep_.CoverTab[21370]++
									continue
//line /usr/local/go/src/crypto/tls/common.go:1005
			// _ = "end of CoverTab[21370]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1006
			_go_fuzz_dep_.CoverTab[21371]++
//line /usr/local/go/src/crypto/tls/common.go:1006
			// _ = "end of CoverTab[21371]"
//line /usr/local/go/src/crypto/tls/common.go:1006
		}
//line /usr/local/go/src/crypto/tls/common.go:1006
		// _ = "end of CoverTab[21356]"
//line /usr/local/go/src/crypto/tls/common.go:1006
		_go_fuzz_dep_.CoverTab[21357]++
								if c != nil && func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1007
			_go_fuzz_dep_.CoverTab[21372]++
//line /usr/local/go/src/crypto/tls/common.go:1007
			return c.MaxVersion != 0
//line /usr/local/go/src/crypto/tls/common.go:1007
			// _ = "end of CoverTab[21372]"
//line /usr/local/go/src/crypto/tls/common.go:1007
		}() && func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1007
			_go_fuzz_dep_.CoverTab[21373]++
//line /usr/local/go/src/crypto/tls/common.go:1007
			return v > c.MaxVersion
//line /usr/local/go/src/crypto/tls/common.go:1007
			// _ = "end of CoverTab[21373]"
//line /usr/local/go/src/crypto/tls/common.go:1007
		}() {
//line /usr/local/go/src/crypto/tls/common.go:1007
			_go_fuzz_dep_.CoverTab[21374]++
									continue
//line /usr/local/go/src/crypto/tls/common.go:1008
			// _ = "end of CoverTab[21374]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1009
			_go_fuzz_dep_.CoverTab[21375]++
//line /usr/local/go/src/crypto/tls/common.go:1009
			// _ = "end of CoverTab[21375]"
//line /usr/local/go/src/crypto/tls/common.go:1009
		}
//line /usr/local/go/src/crypto/tls/common.go:1009
		// _ = "end of CoverTab[21357]"
//line /usr/local/go/src/crypto/tls/common.go:1009
		_go_fuzz_dep_.CoverTab[21358]++
								versions = append(versions, v)
//line /usr/local/go/src/crypto/tls/common.go:1010
		// _ = "end of CoverTab[21358]"
	}
//line /usr/local/go/src/crypto/tls/common.go:1011
	// _ = "end of CoverTab[21352]"
//line /usr/local/go/src/crypto/tls/common.go:1011
	_go_fuzz_dep_.CoverTab[21353]++
							return versions
//line /usr/local/go/src/crypto/tls/common.go:1012
	// _ = "end of CoverTab[21353]"
}

func (c *Config) maxSupportedVersion(isClient bool) uint16 {
//line /usr/local/go/src/crypto/tls/common.go:1015
	_go_fuzz_dep_.CoverTab[21376]++
							supportedVersions := c.supportedVersions(isClient)
							if len(supportedVersions) == 0 {
//line /usr/local/go/src/crypto/tls/common.go:1017
		_go_fuzz_dep_.CoverTab[21378]++
								return 0
//line /usr/local/go/src/crypto/tls/common.go:1018
		// _ = "end of CoverTab[21378]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1019
		_go_fuzz_dep_.CoverTab[21379]++
//line /usr/local/go/src/crypto/tls/common.go:1019
		// _ = "end of CoverTab[21379]"
//line /usr/local/go/src/crypto/tls/common.go:1019
	}
//line /usr/local/go/src/crypto/tls/common.go:1019
	// _ = "end of CoverTab[21376]"
//line /usr/local/go/src/crypto/tls/common.go:1019
	_go_fuzz_dep_.CoverTab[21377]++
							return supportedVersions[0]
//line /usr/local/go/src/crypto/tls/common.go:1020
	// _ = "end of CoverTab[21377]"
}

//line /usr/local/go/src/crypto/tls/common.go:1026
func supportedVersionsFromMax(maxVersion uint16) []uint16 {
//line /usr/local/go/src/crypto/tls/common.go:1026
	_go_fuzz_dep_.CoverTab[21380]++
							versions := make([]uint16, 0, len(supportedVersions))
							for _, v := range supportedVersions {
//line /usr/local/go/src/crypto/tls/common.go:1028
		_go_fuzz_dep_.CoverTab[21382]++
								if v > maxVersion {
//line /usr/local/go/src/crypto/tls/common.go:1029
			_go_fuzz_dep_.CoverTab[21384]++
									continue
//line /usr/local/go/src/crypto/tls/common.go:1030
			// _ = "end of CoverTab[21384]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1031
			_go_fuzz_dep_.CoverTab[21385]++
//line /usr/local/go/src/crypto/tls/common.go:1031
			// _ = "end of CoverTab[21385]"
//line /usr/local/go/src/crypto/tls/common.go:1031
		}
//line /usr/local/go/src/crypto/tls/common.go:1031
		// _ = "end of CoverTab[21382]"
//line /usr/local/go/src/crypto/tls/common.go:1031
		_go_fuzz_dep_.CoverTab[21383]++
								versions = append(versions, v)
//line /usr/local/go/src/crypto/tls/common.go:1032
		// _ = "end of CoverTab[21383]"
	}
//line /usr/local/go/src/crypto/tls/common.go:1033
	// _ = "end of CoverTab[21380]"
//line /usr/local/go/src/crypto/tls/common.go:1033
	_go_fuzz_dep_.CoverTab[21381]++
							return versions
//line /usr/local/go/src/crypto/tls/common.go:1034
	// _ = "end of CoverTab[21381]"
}

var defaultCurvePreferences = []CurveID{X25519, CurveP256, CurveP384, CurveP521}

func (c *Config) curvePreferences() []CurveID {
//line /usr/local/go/src/crypto/tls/common.go:1039
	_go_fuzz_dep_.CoverTab[21386]++
							if needFIPS() {
//line /usr/local/go/src/crypto/tls/common.go:1040
		_go_fuzz_dep_.CoverTab[21389]++
								return fipsCurvePreferences(c)
//line /usr/local/go/src/crypto/tls/common.go:1041
		// _ = "end of CoverTab[21389]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1042
		_go_fuzz_dep_.CoverTab[21390]++
//line /usr/local/go/src/crypto/tls/common.go:1042
		// _ = "end of CoverTab[21390]"
//line /usr/local/go/src/crypto/tls/common.go:1042
	}
//line /usr/local/go/src/crypto/tls/common.go:1042
	// _ = "end of CoverTab[21386]"
//line /usr/local/go/src/crypto/tls/common.go:1042
	_go_fuzz_dep_.CoverTab[21387]++
							if c == nil || func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1043
		_go_fuzz_dep_.CoverTab[21391]++
//line /usr/local/go/src/crypto/tls/common.go:1043
		return len(c.CurvePreferences) == 0
//line /usr/local/go/src/crypto/tls/common.go:1043
		// _ = "end of CoverTab[21391]"
//line /usr/local/go/src/crypto/tls/common.go:1043
	}() {
//line /usr/local/go/src/crypto/tls/common.go:1043
		_go_fuzz_dep_.CoverTab[21392]++
								return defaultCurvePreferences
//line /usr/local/go/src/crypto/tls/common.go:1044
		// _ = "end of CoverTab[21392]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1045
		_go_fuzz_dep_.CoverTab[21393]++
//line /usr/local/go/src/crypto/tls/common.go:1045
		// _ = "end of CoverTab[21393]"
//line /usr/local/go/src/crypto/tls/common.go:1045
	}
//line /usr/local/go/src/crypto/tls/common.go:1045
	// _ = "end of CoverTab[21387]"
//line /usr/local/go/src/crypto/tls/common.go:1045
	_go_fuzz_dep_.CoverTab[21388]++
							return c.CurvePreferences
//line /usr/local/go/src/crypto/tls/common.go:1046
	// _ = "end of CoverTab[21388]"
}

func (c *Config) supportsCurve(curve CurveID) bool {
//line /usr/local/go/src/crypto/tls/common.go:1049
	_go_fuzz_dep_.CoverTab[21394]++
							for _, cc := range c.curvePreferences() {
//line /usr/local/go/src/crypto/tls/common.go:1050
		_go_fuzz_dep_.CoverTab[21396]++
								if cc == curve {
//line /usr/local/go/src/crypto/tls/common.go:1051
			_go_fuzz_dep_.CoverTab[21397]++
									return true
//line /usr/local/go/src/crypto/tls/common.go:1052
			// _ = "end of CoverTab[21397]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1053
			_go_fuzz_dep_.CoverTab[21398]++
//line /usr/local/go/src/crypto/tls/common.go:1053
			// _ = "end of CoverTab[21398]"
//line /usr/local/go/src/crypto/tls/common.go:1053
		}
//line /usr/local/go/src/crypto/tls/common.go:1053
		// _ = "end of CoverTab[21396]"
	}
//line /usr/local/go/src/crypto/tls/common.go:1054
	// _ = "end of CoverTab[21394]"
//line /usr/local/go/src/crypto/tls/common.go:1054
	_go_fuzz_dep_.CoverTab[21395]++
							return false
//line /usr/local/go/src/crypto/tls/common.go:1055
	// _ = "end of CoverTab[21395]"
}

//line /usr/local/go/src/crypto/tls/common.go:1060
func (c *Config) mutualVersion(isClient bool, peerVersions []uint16) (uint16, bool) {
//line /usr/local/go/src/crypto/tls/common.go:1060
	_go_fuzz_dep_.CoverTab[21399]++
							supportedVersions := c.supportedVersions(isClient)
							for _, peerVersion := range peerVersions {
//line /usr/local/go/src/crypto/tls/common.go:1062
		_go_fuzz_dep_.CoverTab[21401]++
								for _, v := range supportedVersions {
//line /usr/local/go/src/crypto/tls/common.go:1063
			_go_fuzz_dep_.CoverTab[21402]++
									if v == peerVersion {
//line /usr/local/go/src/crypto/tls/common.go:1064
				_go_fuzz_dep_.CoverTab[21403]++
										return v, true
//line /usr/local/go/src/crypto/tls/common.go:1065
				// _ = "end of CoverTab[21403]"
			} else {
//line /usr/local/go/src/crypto/tls/common.go:1066
				_go_fuzz_dep_.CoverTab[21404]++
//line /usr/local/go/src/crypto/tls/common.go:1066
				// _ = "end of CoverTab[21404]"
//line /usr/local/go/src/crypto/tls/common.go:1066
			}
//line /usr/local/go/src/crypto/tls/common.go:1066
			// _ = "end of CoverTab[21402]"
		}
//line /usr/local/go/src/crypto/tls/common.go:1067
		// _ = "end of CoverTab[21401]"
	}
//line /usr/local/go/src/crypto/tls/common.go:1068
	// _ = "end of CoverTab[21399]"
//line /usr/local/go/src/crypto/tls/common.go:1068
	_go_fuzz_dep_.CoverTab[21400]++
							return 0, false
//line /usr/local/go/src/crypto/tls/common.go:1069
	// _ = "end of CoverTab[21400]"
}

var errNoCertificates = errors.New("tls: no certificates configured")

//line /usr/local/go/src/crypto/tls/common.go:1076
func (c *Config) getCertificate(clientHello *ClientHelloInfo) (*Certificate, error) {
//line /usr/local/go/src/crypto/tls/common.go:1076
	_go_fuzz_dep_.CoverTab[21405]++
							if c.GetCertificate != nil && func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1077
		_go_fuzz_dep_.CoverTab[21411]++
//line /usr/local/go/src/crypto/tls/common.go:1077
		return (len(c.Certificates) == 0 || func() bool {
									_go_fuzz_dep_.CoverTab[21412]++
//line /usr/local/go/src/crypto/tls/common.go:1078
			return len(clientHello.ServerName) > 0
//line /usr/local/go/src/crypto/tls/common.go:1078
			// _ = "end of CoverTab[21412]"
//line /usr/local/go/src/crypto/tls/common.go:1078
		}())
//line /usr/local/go/src/crypto/tls/common.go:1078
		// _ = "end of CoverTab[21411]"
//line /usr/local/go/src/crypto/tls/common.go:1078
	}() {
//line /usr/local/go/src/crypto/tls/common.go:1078
		_go_fuzz_dep_.CoverTab[21413]++
								cert, err := c.GetCertificate(clientHello)
								if cert != nil || func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1080
			_go_fuzz_dep_.CoverTab[21414]++
//line /usr/local/go/src/crypto/tls/common.go:1080
			return err != nil
//line /usr/local/go/src/crypto/tls/common.go:1080
			// _ = "end of CoverTab[21414]"
//line /usr/local/go/src/crypto/tls/common.go:1080
		}() {
//line /usr/local/go/src/crypto/tls/common.go:1080
			_go_fuzz_dep_.CoverTab[21415]++
									return cert, err
//line /usr/local/go/src/crypto/tls/common.go:1081
			// _ = "end of CoverTab[21415]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1082
			_go_fuzz_dep_.CoverTab[21416]++
//line /usr/local/go/src/crypto/tls/common.go:1082
			// _ = "end of CoverTab[21416]"
//line /usr/local/go/src/crypto/tls/common.go:1082
		}
//line /usr/local/go/src/crypto/tls/common.go:1082
		// _ = "end of CoverTab[21413]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1083
		_go_fuzz_dep_.CoverTab[21417]++
//line /usr/local/go/src/crypto/tls/common.go:1083
		// _ = "end of CoverTab[21417]"
//line /usr/local/go/src/crypto/tls/common.go:1083
	}
//line /usr/local/go/src/crypto/tls/common.go:1083
	// _ = "end of CoverTab[21405]"
//line /usr/local/go/src/crypto/tls/common.go:1083
	_go_fuzz_dep_.CoverTab[21406]++

							if len(c.Certificates) == 0 {
//line /usr/local/go/src/crypto/tls/common.go:1085
		_go_fuzz_dep_.CoverTab[21418]++
								return nil, errNoCertificates
//line /usr/local/go/src/crypto/tls/common.go:1086
		// _ = "end of CoverTab[21418]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1087
		_go_fuzz_dep_.CoverTab[21419]++
//line /usr/local/go/src/crypto/tls/common.go:1087
		// _ = "end of CoverTab[21419]"
//line /usr/local/go/src/crypto/tls/common.go:1087
	}
//line /usr/local/go/src/crypto/tls/common.go:1087
	// _ = "end of CoverTab[21406]"
//line /usr/local/go/src/crypto/tls/common.go:1087
	_go_fuzz_dep_.CoverTab[21407]++

							if len(c.Certificates) == 1 {
//line /usr/local/go/src/crypto/tls/common.go:1089
		_go_fuzz_dep_.CoverTab[21420]++

								return &c.Certificates[0], nil
//line /usr/local/go/src/crypto/tls/common.go:1091
		// _ = "end of CoverTab[21420]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1092
		_go_fuzz_dep_.CoverTab[21421]++
//line /usr/local/go/src/crypto/tls/common.go:1092
		// _ = "end of CoverTab[21421]"
//line /usr/local/go/src/crypto/tls/common.go:1092
	}
//line /usr/local/go/src/crypto/tls/common.go:1092
	// _ = "end of CoverTab[21407]"
//line /usr/local/go/src/crypto/tls/common.go:1092
	_go_fuzz_dep_.CoverTab[21408]++

							if c.NameToCertificate != nil {
//line /usr/local/go/src/crypto/tls/common.go:1094
		_go_fuzz_dep_.CoverTab[21422]++
								name := strings.ToLower(clientHello.ServerName)
								if cert, ok := c.NameToCertificate[name]; ok {
//line /usr/local/go/src/crypto/tls/common.go:1096
			_go_fuzz_dep_.CoverTab[21424]++
									return cert, nil
//line /usr/local/go/src/crypto/tls/common.go:1097
			// _ = "end of CoverTab[21424]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1098
			_go_fuzz_dep_.CoverTab[21425]++
//line /usr/local/go/src/crypto/tls/common.go:1098
			// _ = "end of CoverTab[21425]"
//line /usr/local/go/src/crypto/tls/common.go:1098
		}
//line /usr/local/go/src/crypto/tls/common.go:1098
		// _ = "end of CoverTab[21422]"
//line /usr/local/go/src/crypto/tls/common.go:1098
		_go_fuzz_dep_.CoverTab[21423]++
								if len(name) > 0 {
//line /usr/local/go/src/crypto/tls/common.go:1099
			_go_fuzz_dep_.CoverTab[21426]++
									labels := strings.Split(name, ".")
									labels[0] = "*"
									wildcardName := strings.Join(labels, ".")
									if cert, ok := c.NameToCertificate[wildcardName]; ok {
//line /usr/local/go/src/crypto/tls/common.go:1103
				_go_fuzz_dep_.CoverTab[21427]++
										return cert, nil
//line /usr/local/go/src/crypto/tls/common.go:1104
				// _ = "end of CoverTab[21427]"
			} else {
//line /usr/local/go/src/crypto/tls/common.go:1105
				_go_fuzz_dep_.CoverTab[21428]++
//line /usr/local/go/src/crypto/tls/common.go:1105
				// _ = "end of CoverTab[21428]"
//line /usr/local/go/src/crypto/tls/common.go:1105
			}
//line /usr/local/go/src/crypto/tls/common.go:1105
			// _ = "end of CoverTab[21426]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1106
			_go_fuzz_dep_.CoverTab[21429]++
//line /usr/local/go/src/crypto/tls/common.go:1106
			// _ = "end of CoverTab[21429]"
//line /usr/local/go/src/crypto/tls/common.go:1106
		}
//line /usr/local/go/src/crypto/tls/common.go:1106
		// _ = "end of CoverTab[21423]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1107
		_go_fuzz_dep_.CoverTab[21430]++
//line /usr/local/go/src/crypto/tls/common.go:1107
		// _ = "end of CoverTab[21430]"
//line /usr/local/go/src/crypto/tls/common.go:1107
	}
//line /usr/local/go/src/crypto/tls/common.go:1107
	// _ = "end of CoverTab[21408]"
//line /usr/local/go/src/crypto/tls/common.go:1107
	_go_fuzz_dep_.CoverTab[21409]++

							for _, cert := range c.Certificates {
//line /usr/local/go/src/crypto/tls/common.go:1109
		_go_fuzz_dep_.CoverTab[21431]++
								if err := clientHello.SupportsCertificate(&cert); err == nil {
//line /usr/local/go/src/crypto/tls/common.go:1110
			_go_fuzz_dep_.CoverTab[21432]++
									return &cert, nil
//line /usr/local/go/src/crypto/tls/common.go:1111
			// _ = "end of CoverTab[21432]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1112
			_go_fuzz_dep_.CoverTab[21433]++
//line /usr/local/go/src/crypto/tls/common.go:1112
			// _ = "end of CoverTab[21433]"
//line /usr/local/go/src/crypto/tls/common.go:1112
		}
//line /usr/local/go/src/crypto/tls/common.go:1112
		// _ = "end of CoverTab[21431]"
	}
//line /usr/local/go/src/crypto/tls/common.go:1113
	// _ = "end of CoverTab[21409]"
//line /usr/local/go/src/crypto/tls/common.go:1113
	_go_fuzz_dep_.CoverTab[21410]++

//line /usr/local/go/src/crypto/tls/common.go:1116
	return &c.Certificates[0], nil
//line /usr/local/go/src/crypto/tls/common.go:1116
	// _ = "end of CoverTab[21410]"
}

//line /usr/local/go/src/crypto/tls/common.go:1130
func (chi *ClientHelloInfo) SupportsCertificate(c *Certificate) error {
//line /usr/local/go/src/crypto/tls/common.go:1130
	_go_fuzz_dep_.CoverTab[21434]++

//line /usr/local/go/src/crypto/tls/common.go:1136
	config := chi.config
	if config == nil {
//line /usr/local/go/src/crypto/tls/common.go:1137
		_go_fuzz_dep_.CoverTab[21445]++
								config = &Config{}
//line /usr/local/go/src/crypto/tls/common.go:1138
		// _ = "end of CoverTab[21445]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1139
		_go_fuzz_dep_.CoverTab[21446]++
//line /usr/local/go/src/crypto/tls/common.go:1139
		// _ = "end of CoverTab[21446]"
//line /usr/local/go/src/crypto/tls/common.go:1139
	}
//line /usr/local/go/src/crypto/tls/common.go:1139
	// _ = "end of CoverTab[21434]"
//line /usr/local/go/src/crypto/tls/common.go:1139
	_go_fuzz_dep_.CoverTab[21435]++
							vers, ok := config.mutualVersion(roleServer, chi.SupportedVersions)
							if !ok {
//line /usr/local/go/src/crypto/tls/common.go:1141
		_go_fuzz_dep_.CoverTab[21447]++
								return errors.New("no mutually supported protocol versions")
//line /usr/local/go/src/crypto/tls/common.go:1142
		// _ = "end of CoverTab[21447]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1143
		_go_fuzz_dep_.CoverTab[21448]++
//line /usr/local/go/src/crypto/tls/common.go:1143
		// _ = "end of CoverTab[21448]"
//line /usr/local/go/src/crypto/tls/common.go:1143
	}
//line /usr/local/go/src/crypto/tls/common.go:1143
	// _ = "end of CoverTab[21435]"
//line /usr/local/go/src/crypto/tls/common.go:1143
	_go_fuzz_dep_.CoverTab[21436]++

//line /usr/local/go/src/crypto/tls/common.go:1147
	if chi.ServerName != "" {
//line /usr/local/go/src/crypto/tls/common.go:1147
		_go_fuzz_dep_.CoverTab[21449]++
								x509Cert, err := c.leaf()
								if err != nil {
//line /usr/local/go/src/crypto/tls/common.go:1149
			_go_fuzz_dep_.CoverTab[21451]++
									return fmt.Errorf("failed to parse certificate: %w", err)
//line /usr/local/go/src/crypto/tls/common.go:1150
			// _ = "end of CoverTab[21451]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1151
			_go_fuzz_dep_.CoverTab[21452]++
//line /usr/local/go/src/crypto/tls/common.go:1151
			// _ = "end of CoverTab[21452]"
//line /usr/local/go/src/crypto/tls/common.go:1151
		}
//line /usr/local/go/src/crypto/tls/common.go:1151
		// _ = "end of CoverTab[21449]"
//line /usr/local/go/src/crypto/tls/common.go:1151
		_go_fuzz_dep_.CoverTab[21450]++
								if err := x509Cert.VerifyHostname(chi.ServerName); err != nil {
//line /usr/local/go/src/crypto/tls/common.go:1152
			_go_fuzz_dep_.CoverTab[21453]++
									return fmt.Errorf("certificate is not valid for requested server name: %w", err)
//line /usr/local/go/src/crypto/tls/common.go:1153
			// _ = "end of CoverTab[21453]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1154
			_go_fuzz_dep_.CoverTab[21454]++
//line /usr/local/go/src/crypto/tls/common.go:1154
			// _ = "end of CoverTab[21454]"
//line /usr/local/go/src/crypto/tls/common.go:1154
		}
//line /usr/local/go/src/crypto/tls/common.go:1154
		// _ = "end of CoverTab[21450]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1155
		_go_fuzz_dep_.CoverTab[21455]++
//line /usr/local/go/src/crypto/tls/common.go:1155
		// _ = "end of CoverTab[21455]"
//line /usr/local/go/src/crypto/tls/common.go:1155
	}
//line /usr/local/go/src/crypto/tls/common.go:1155
	// _ = "end of CoverTab[21436]"
//line /usr/local/go/src/crypto/tls/common.go:1155
	_go_fuzz_dep_.CoverTab[21437]++

//line /usr/local/go/src/crypto/tls/common.go:1161
	supportsRSAFallback := func(unsupported error) error {
//line /usr/local/go/src/crypto/tls/common.go:1161
		_go_fuzz_dep_.CoverTab[21456]++

								if vers == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/common.go:1163
			_go_fuzz_dep_.CoverTab[21461]++
									return unsupported
//line /usr/local/go/src/crypto/tls/common.go:1164
			// _ = "end of CoverTab[21461]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1165
			_go_fuzz_dep_.CoverTab[21462]++
//line /usr/local/go/src/crypto/tls/common.go:1165
			// _ = "end of CoverTab[21462]"
//line /usr/local/go/src/crypto/tls/common.go:1165
		}
//line /usr/local/go/src/crypto/tls/common.go:1165
		// _ = "end of CoverTab[21456]"
//line /usr/local/go/src/crypto/tls/common.go:1165
		_go_fuzz_dep_.CoverTab[21457]++

//line /usr/local/go/src/crypto/tls/common.go:1169
		if priv, ok := c.PrivateKey.(crypto.Decrypter); ok {
//line /usr/local/go/src/crypto/tls/common.go:1169
			_go_fuzz_dep_.CoverTab[21463]++
									if _, ok := priv.Public().(*rsa.PublicKey); !ok {
//line /usr/local/go/src/crypto/tls/common.go:1170
				_go_fuzz_dep_.CoverTab[21464]++
										return unsupported
//line /usr/local/go/src/crypto/tls/common.go:1171
				// _ = "end of CoverTab[21464]"
			} else {
//line /usr/local/go/src/crypto/tls/common.go:1172
				_go_fuzz_dep_.CoverTab[21465]++
//line /usr/local/go/src/crypto/tls/common.go:1172
				// _ = "end of CoverTab[21465]"
//line /usr/local/go/src/crypto/tls/common.go:1172
			}
//line /usr/local/go/src/crypto/tls/common.go:1172
			// _ = "end of CoverTab[21463]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1173
			_go_fuzz_dep_.CoverTab[21466]++
									return unsupported
//line /usr/local/go/src/crypto/tls/common.go:1174
			// _ = "end of CoverTab[21466]"
		}
//line /usr/local/go/src/crypto/tls/common.go:1175
		// _ = "end of CoverTab[21457]"
//line /usr/local/go/src/crypto/tls/common.go:1175
		_go_fuzz_dep_.CoverTab[21458]++

//line /usr/local/go/src/crypto/tls/common.go:1178
		rsaCipherSuite := selectCipherSuite(chi.CipherSuites, config.cipherSuites(), func(c *cipherSuite) bool {
//line /usr/local/go/src/crypto/tls/common.go:1178
			_go_fuzz_dep_.CoverTab[21467]++
									if c.flags&suiteECDHE != 0 {
//line /usr/local/go/src/crypto/tls/common.go:1179
				_go_fuzz_dep_.CoverTab[21470]++
										return false
//line /usr/local/go/src/crypto/tls/common.go:1180
				// _ = "end of CoverTab[21470]"
			} else {
//line /usr/local/go/src/crypto/tls/common.go:1181
				_go_fuzz_dep_.CoverTab[21471]++
//line /usr/local/go/src/crypto/tls/common.go:1181
				// _ = "end of CoverTab[21471]"
//line /usr/local/go/src/crypto/tls/common.go:1181
			}
//line /usr/local/go/src/crypto/tls/common.go:1181
			// _ = "end of CoverTab[21467]"
//line /usr/local/go/src/crypto/tls/common.go:1181
			_go_fuzz_dep_.CoverTab[21468]++
									if vers < VersionTLS12 && func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1182
				_go_fuzz_dep_.CoverTab[21472]++
//line /usr/local/go/src/crypto/tls/common.go:1182
				return c.flags&suiteTLS12 != 0
//line /usr/local/go/src/crypto/tls/common.go:1182
				// _ = "end of CoverTab[21472]"
//line /usr/local/go/src/crypto/tls/common.go:1182
			}() {
//line /usr/local/go/src/crypto/tls/common.go:1182
				_go_fuzz_dep_.CoverTab[21473]++
										return false
//line /usr/local/go/src/crypto/tls/common.go:1183
				// _ = "end of CoverTab[21473]"
			} else {
//line /usr/local/go/src/crypto/tls/common.go:1184
				_go_fuzz_dep_.CoverTab[21474]++
//line /usr/local/go/src/crypto/tls/common.go:1184
				// _ = "end of CoverTab[21474]"
//line /usr/local/go/src/crypto/tls/common.go:1184
			}
//line /usr/local/go/src/crypto/tls/common.go:1184
			// _ = "end of CoverTab[21468]"
//line /usr/local/go/src/crypto/tls/common.go:1184
			_go_fuzz_dep_.CoverTab[21469]++
									return true
//line /usr/local/go/src/crypto/tls/common.go:1185
			// _ = "end of CoverTab[21469]"
		})
//line /usr/local/go/src/crypto/tls/common.go:1186
		// _ = "end of CoverTab[21458]"
//line /usr/local/go/src/crypto/tls/common.go:1186
		_go_fuzz_dep_.CoverTab[21459]++
								if rsaCipherSuite == nil {
//line /usr/local/go/src/crypto/tls/common.go:1187
			_go_fuzz_dep_.CoverTab[21475]++
									return unsupported
//line /usr/local/go/src/crypto/tls/common.go:1188
			// _ = "end of CoverTab[21475]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1189
			_go_fuzz_dep_.CoverTab[21476]++
//line /usr/local/go/src/crypto/tls/common.go:1189
			// _ = "end of CoverTab[21476]"
//line /usr/local/go/src/crypto/tls/common.go:1189
		}
//line /usr/local/go/src/crypto/tls/common.go:1189
		// _ = "end of CoverTab[21459]"
//line /usr/local/go/src/crypto/tls/common.go:1189
		_go_fuzz_dep_.CoverTab[21460]++
								return nil
//line /usr/local/go/src/crypto/tls/common.go:1190
		// _ = "end of CoverTab[21460]"
	}
//line /usr/local/go/src/crypto/tls/common.go:1191
	// _ = "end of CoverTab[21437]"
//line /usr/local/go/src/crypto/tls/common.go:1191
	_go_fuzz_dep_.CoverTab[21438]++

//line /usr/local/go/src/crypto/tls/common.go:1195
	if len(chi.SignatureSchemes) > 0 {
//line /usr/local/go/src/crypto/tls/common.go:1195
		_go_fuzz_dep_.CoverTab[21477]++
								if _, err := selectSignatureScheme(vers, c, chi.SignatureSchemes); err != nil {
//line /usr/local/go/src/crypto/tls/common.go:1196
			_go_fuzz_dep_.CoverTab[21478]++
									return supportsRSAFallback(err)
//line /usr/local/go/src/crypto/tls/common.go:1197
			// _ = "end of CoverTab[21478]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1198
			_go_fuzz_dep_.CoverTab[21479]++
//line /usr/local/go/src/crypto/tls/common.go:1198
			// _ = "end of CoverTab[21479]"
//line /usr/local/go/src/crypto/tls/common.go:1198
		}
//line /usr/local/go/src/crypto/tls/common.go:1198
		// _ = "end of CoverTab[21477]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1199
		_go_fuzz_dep_.CoverTab[21480]++
//line /usr/local/go/src/crypto/tls/common.go:1199
		// _ = "end of CoverTab[21480]"
//line /usr/local/go/src/crypto/tls/common.go:1199
	}
//line /usr/local/go/src/crypto/tls/common.go:1199
	// _ = "end of CoverTab[21438]"
//line /usr/local/go/src/crypto/tls/common.go:1199
	_go_fuzz_dep_.CoverTab[21439]++

//line /usr/local/go/src/crypto/tls/common.go:1204
	if vers == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/common.go:1204
		_go_fuzz_dep_.CoverTab[21481]++
								return nil
//line /usr/local/go/src/crypto/tls/common.go:1205
		// _ = "end of CoverTab[21481]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1206
		_go_fuzz_dep_.CoverTab[21482]++
//line /usr/local/go/src/crypto/tls/common.go:1206
		// _ = "end of CoverTab[21482]"
//line /usr/local/go/src/crypto/tls/common.go:1206
	}
//line /usr/local/go/src/crypto/tls/common.go:1206
	// _ = "end of CoverTab[21439]"
//line /usr/local/go/src/crypto/tls/common.go:1206
	_go_fuzz_dep_.CoverTab[21440]++

//line /usr/local/go/src/crypto/tls/common.go:1209
	if !supportsECDHE(config, chi.SupportedCurves, chi.SupportedPoints) {
//line /usr/local/go/src/crypto/tls/common.go:1209
		_go_fuzz_dep_.CoverTab[21483]++
								return supportsRSAFallback(errors.New("client doesn't support ECDHE, can only use legacy RSA key exchange"))
//line /usr/local/go/src/crypto/tls/common.go:1210
		// _ = "end of CoverTab[21483]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1211
		_go_fuzz_dep_.CoverTab[21484]++
//line /usr/local/go/src/crypto/tls/common.go:1211
		// _ = "end of CoverTab[21484]"
//line /usr/local/go/src/crypto/tls/common.go:1211
	}
//line /usr/local/go/src/crypto/tls/common.go:1211
	// _ = "end of CoverTab[21440]"
//line /usr/local/go/src/crypto/tls/common.go:1211
	_go_fuzz_dep_.CoverTab[21441]++

							var ecdsaCipherSuite bool
							if priv, ok := c.PrivateKey.(crypto.Signer); ok {
//line /usr/local/go/src/crypto/tls/common.go:1214
		_go_fuzz_dep_.CoverTab[21485]++
								switch pub := priv.Public().(type) {
		case *ecdsa.PublicKey:
//line /usr/local/go/src/crypto/tls/common.go:1216
			_go_fuzz_dep_.CoverTab[21486]++
									var curve CurveID
									switch pub.Curve {
			case elliptic.P256():
//line /usr/local/go/src/crypto/tls/common.go:1219
				_go_fuzz_dep_.CoverTab[21494]++
										curve = CurveP256
//line /usr/local/go/src/crypto/tls/common.go:1220
				// _ = "end of CoverTab[21494]"
			case elliptic.P384():
//line /usr/local/go/src/crypto/tls/common.go:1221
				_go_fuzz_dep_.CoverTab[21495]++
										curve = CurveP384
//line /usr/local/go/src/crypto/tls/common.go:1222
				// _ = "end of CoverTab[21495]"
			case elliptic.P521():
//line /usr/local/go/src/crypto/tls/common.go:1223
				_go_fuzz_dep_.CoverTab[21496]++
										curve = CurveP521
//line /usr/local/go/src/crypto/tls/common.go:1224
				// _ = "end of CoverTab[21496]"
			default:
//line /usr/local/go/src/crypto/tls/common.go:1225
				_go_fuzz_dep_.CoverTab[21497]++
										return supportsRSAFallback(unsupportedCertificateError(c))
//line /usr/local/go/src/crypto/tls/common.go:1226
				// _ = "end of CoverTab[21497]"
			}
//line /usr/local/go/src/crypto/tls/common.go:1227
			// _ = "end of CoverTab[21486]"
//line /usr/local/go/src/crypto/tls/common.go:1227
			_go_fuzz_dep_.CoverTab[21487]++
									var curveOk bool
									for _, c := range chi.SupportedCurves {
//line /usr/local/go/src/crypto/tls/common.go:1229
				_go_fuzz_dep_.CoverTab[21498]++
										if c == curve && func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1230
					_go_fuzz_dep_.CoverTab[21499]++
//line /usr/local/go/src/crypto/tls/common.go:1230
					return config.supportsCurve(c)
//line /usr/local/go/src/crypto/tls/common.go:1230
					// _ = "end of CoverTab[21499]"
//line /usr/local/go/src/crypto/tls/common.go:1230
				}() {
//line /usr/local/go/src/crypto/tls/common.go:1230
					_go_fuzz_dep_.CoverTab[21500]++
											curveOk = true
											break
//line /usr/local/go/src/crypto/tls/common.go:1232
					// _ = "end of CoverTab[21500]"
				} else {
//line /usr/local/go/src/crypto/tls/common.go:1233
					_go_fuzz_dep_.CoverTab[21501]++
//line /usr/local/go/src/crypto/tls/common.go:1233
					// _ = "end of CoverTab[21501]"
//line /usr/local/go/src/crypto/tls/common.go:1233
				}
//line /usr/local/go/src/crypto/tls/common.go:1233
				// _ = "end of CoverTab[21498]"
			}
//line /usr/local/go/src/crypto/tls/common.go:1234
			// _ = "end of CoverTab[21487]"
//line /usr/local/go/src/crypto/tls/common.go:1234
			_go_fuzz_dep_.CoverTab[21488]++
									if !curveOk {
//line /usr/local/go/src/crypto/tls/common.go:1235
				_go_fuzz_dep_.CoverTab[21502]++
										return errors.New("client doesn't support certificate curve")
//line /usr/local/go/src/crypto/tls/common.go:1236
				// _ = "end of CoverTab[21502]"
			} else {
//line /usr/local/go/src/crypto/tls/common.go:1237
				_go_fuzz_dep_.CoverTab[21503]++
//line /usr/local/go/src/crypto/tls/common.go:1237
				// _ = "end of CoverTab[21503]"
//line /usr/local/go/src/crypto/tls/common.go:1237
			}
//line /usr/local/go/src/crypto/tls/common.go:1237
			// _ = "end of CoverTab[21488]"
//line /usr/local/go/src/crypto/tls/common.go:1237
			_go_fuzz_dep_.CoverTab[21489]++
									ecdsaCipherSuite = true
//line /usr/local/go/src/crypto/tls/common.go:1238
			// _ = "end of CoverTab[21489]"
		case ed25519.PublicKey:
//line /usr/local/go/src/crypto/tls/common.go:1239
			_go_fuzz_dep_.CoverTab[21490]++
									if vers < VersionTLS12 || func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1240
				_go_fuzz_dep_.CoverTab[21504]++
//line /usr/local/go/src/crypto/tls/common.go:1240
				return len(chi.SignatureSchemes) == 0
//line /usr/local/go/src/crypto/tls/common.go:1240
				// _ = "end of CoverTab[21504]"
//line /usr/local/go/src/crypto/tls/common.go:1240
			}() {
//line /usr/local/go/src/crypto/tls/common.go:1240
				_go_fuzz_dep_.CoverTab[21505]++
										return errors.New("connection doesn't support Ed25519")
//line /usr/local/go/src/crypto/tls/common.go:1241
				// _ = "end of CoverTab[21505]"
			} else {
//line /usr/local/go/src/crypto/tls/common.go:1242
				_go_fuzz_dep_.CoverTab[21506]++
//line /usr/local/go/src/crypto/tls/common.go:1242
				// _ = "end of CoverTab[21506]"
//line /usr/local/go/src/crypto/tls/common.go:1242
			}
//line /usr/local/go/src/crypto/tls/common.go:1242
			// _ = "end of CoverTab[21490]"
//line /usr/local/go/src/crypto/tls/common.go:1242
			_go_fuzz_dep_.CoverTab[21491]++
									ecdsaCipherSuite = true
//line /usr/local/go/src/crypto/tls/common.go:1243
			// _ = "end of CoverTab[21491]"
		case *rsa.PublicKey:
//line /usr/local/go/src/crypto/tls/common.go:1244
			_go_fuzz_dep_.CoverTab[21492]++
//line /usr/local/go/src/crypto/tls/common.go:1244
			// _ = "end of CoverTab[21492]"
		default:
//line /usr/local/go/src/crypto/tls/common.go:1245
			_go_fuzz_dep_.CoverTab[21493]++
									return supportsRSAFallback(unsupportedCertificateError(c))
//line /usr/local/go/src/crypto/tls/common.go:1246
			// _ = "end of CoverTab[21493]"
		}
//line /usr/local/go/src/crypto/tls/common.go:1247
		// _ = "end of CoverTab[21485]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1248
		_go_fuzz_dep_.CoverTab[21507]++
								return supportsRSAFallback(unsupportedCertificateError(c))
//line /usr/local/go/src/crypto/tls/common.go:1249
		// _ = "end of CoverTab[21507]"
	}
//line /usr/local/go/src/crypto/tls/common.go:1250
	// _ = "end of CoverTab[21441]"
//line /usr/local/go/src/crypto/tls/common.go:1250
	_go_fuzz_dep_.CoverTab[21442]++

//line /usr/local/go/src/crypto/tls/common.go:1255
	cipherSuite := selectCipherSuite(chi.CipherSuites, config.cipherSuites(), func(c *cipherSuite) bool {
//line /usr/local/go/src/crypto/tls/common.go:1255
		_go_fuzz_dep_.CoverTab[21508]++
								if c.flags&suiteECDHE == 0 {
//line /usr/local/go/src/crypto/tls/common.go:1256
			_go_fuzz_dep_.CoverTab[21512]++
									return false
//line /usr/local/go/src/crypto/tls/common.go:1257
			// _ = "end of CoverTab[21512]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1258
			_go_fuzz_dep_.CoverTab[21513]++
//line /usr/local/go/src/crypto/tls/common.go:1258
			// _ = "end of CoverTab[21513]"
//line /usr/local/go/src/crypto/tls/common.go:1258
		}
//line /usr/local/go/src/crypto/tls/common.go:1258
		// _ = "end of CoverTab[21508]"
//line /usr/local/go/src/crypto/tls/common.go:1258
		_go_fuzz_dep_.CoverTab[21509]++
								if c.flags&suiteECSign != 0 {
//line /usr/local/go/src/crypto/tls/common.go:1259
			_go_fuzz_dep_.CoverTab[21514]++
									if !ecdsaCipherSuite {
//line /usr/local/go/src/crypto/tls/common.go:1260
				_go_fuzz_dep_.CoverTab[21515]++
										return false
//line /usr/local/go/src/crypto/tls/common.go:1261
				// _ = "end of CoverTab[21515]"
			} else {
//line /usr/local/go/src/crypto/tls/common.go:1262
				_go_fuzz_dep_.CoverTab[21516]++
//line /usr/local/go/src/crypto/tls/common.go:1262
				// _ = "end of CoverTab[21516]"
//line /usr/local/go/src/crypto/tls/common.go:1262
			}
//line /usr/local/go/src/crypto/tls/common.go:1262
			// _ = "end of CoverTab[21514]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1263
			_go_fuzz_dep_.CoverTab[21517]++
									if ecdsaCipherSuite {
//line /usr/local/go/src/crypto/tls/common.go:1264
				_go_fuzz_dep_.CoverTab[21518]++
										return false
//line /usr/local/go/src/crypto/tls/common.go:1265
				// _ = "end of CoverTab[21518]"
			} else {
//line /usr/local/go/src/crypto/tls/common.go:1266
				_go_fuzz_dep_.CoverTab[21519]++
//line /usr/local/go/src/crypto/tls/common.go:1266
				// _ = "end of CoverTab[21519]"
//line /usr/local/go/src/crypto/tls/common.go:1266
			}
//line /usr/local/go/src/crypto/tls/common.go:1266
			// _ = "end of CoverTab[21517]"
		}
//line /usr/local/go/src/crypto/tls/common.go:1267
		// _ = "end of CoverTab[21509]"
//line /usr/local/go/src/crypto/tls/common.go:1267
		_go_fuzz_dep_.CoverTab[21510]++
								if vers < VersionTLS12 && func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1268
			_go_fuzz_dep_.CoverTab[21520]++
//line /usr/local/go/src/crypto/tls/common.go:1268
			return c.flags&suiteTLS12 != 0
//line /usr/local/go/src/crypto/tls/common.go:1268
			// _ = "end of CoverTab[21520]"
//line /usr/local/go/src/crypto/tls/common.go:1268
		}() {
//line /usr/local/go/src/crypto/tls/common.go:1268
			_go_fuzz_dep_.CoverTab[21521]++
									return false
//line /usr/local/go/src/crypto/tls/common.go:1269
			// _ = "end of CoverTab[21521]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1270
			_go_fuzz_dep_.CoverTab[21522]++
//line /usr/local/go/src/crypto/tls/common.go:1270
			// _ = "end of CoverTab[21522]"
//line /usr/local/go/src/crypto/tls/common.go:1270
		}
//line /usr/local/go/src/crypto/tls/common.go:1270
		// _ = "end of CoverTab[21510]"
//line /usr/local/go/src/crypto/tls/common.go:1270
		_go_fuzz_dep_.CoverTab[21511]++
								return true
//line /usr/local/go/src/crypto/tls/common.go:1271
		// _ = "end of CoverTab[21511]"
	})
//line /usr/local/go/src/crypto/tls/common.go:1272
	// _ = "end of CoverTab[21442]"
//line /usr/local/go/src/crypto/tls/common.go:1272
	_go_fuzz_dep_.CoverTab[21443]++
							if cipherSuite == nil {
//line /usr/local/go/src/crypto/tls/common.go:1273
		_go_fuzz_dep_.CoverTab[21523]++
								return supportsRSAFallback(errors.New("client doesn't support any cipher suites compatible with the certificate"))
//line /usr/local/go/src/crypto/tls/common.go:1274
		// _ = "end of CoverTab[21523]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1275
		_go_fuzz_dep_.CoverTab[21524]++
//line /usr/local/go/src/crypto/tls/common.go:1275
		// _ = "end of CoverTab[21524]"
//line /usr/local/go/src/crypto/tls/common.go:1275
	}
//line /usr/local/go/src/crypto/tls/common.go:1275
	// _ = "end of CoverTab[21443]"
//line /usr/local/go/src/crypto/tls/common.go:1275
	_go_fuzz_dep_.CoverTab[21444]++

							return nil
//line /usr/local/go/src/crypto/tls/common.go:1277
	// _ = "end of CoverTab[21444]"
}

//line /usr/local/go/src/crypto/tls/common.go:1283
func (cri *CertificateRequestInfo) SupportsCertificate(c *Certificate) error {
//line /usr/local/go/src/crypto/tls/common.go:1283
	_go_fuzz_dep_.CoverTab[21525]++
							if _, err := selectSignatureScheme(cri.Version, c, cri.SignatureSchemes); err != nil {
//line /usr/local/go/src/crypto/tls/common.go:1284
		_go_fuzz_dep_.CoverTab[21529]++
								return err
//line /usr/local/go/src/crypto/tls/common.go:1285
		// _ = "end of CoverTab[21529]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1286
		_go_fuzz_dep_.CoverTab[21530]++
//line /usr/local/go/src/crypto/tls/common.go:1286
		// _ = "end of CoverTab[21530]"
//line /usr/local/go/src/crypto/tls/common.go:1286
	}
//line /usr/local/go/src/crypto/tls/common.go:1286
	// _ = "end of CoverTab[21525]"
//line /usr/local/go/src/crypto/tls/common.go:1286
	_go_fuzz_dep_.CoverTab[21526]++

							if len(cri.AcceptableCAs) == 0 {
//line /usr/local/go/src/crypto/tls/common.go:1288
		_go_fuzz_dep_.CoverTab[21531]++
								return nil
//line /usr/local/go/src/crypto/tls/common.go:1289
		// _ = "end of CoverTab[21531]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1290
		_go_fuzz_dep_.CoverTab[21532]++
//line /usr/local/go/src/crypto/tls/common.go:1290
		// _ = "end of CoverTab[21532]"
//line /usr/local/go/src/crypto/tls/common.go:1290
	}
//line /usr/local/go/src/crypto/tls/common.go:1290
	// _ = "end of CoverTab[21526]"
//line /usr/local/go/src/crypto/tls/common.go:1290
	_go_fuzz_dep_.CoverTab[21527]++

							for j, cert := range c.Certificate {
//line /usr/local/go/src/crypto/tls/common.go:1292
		_go_fuzz_dep_.CoverTab[21533]++
								x509Cert := c.Leaf

//line /usr/local/go/src/crypto/tls/common.go:1296
		if j != 0 || func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1296
			_go_fuzz_dep_.CoverTab[21535]++
//line /usr/local/go/src/crypto/tls/common.go:1296
			return x509Cert == nil
//line /usr/local/go/src/crypto/tls/common.go:1296
			// _ = "end of CoverTab[21535]"
//line /usr/local/go/src/crypto/tls/common.go:1296
		}() {
//line /usr/local/go/src/crypto/tls/common.go:1296
			_go_fuzz_dep_.CoverTab[21536]++
									var err error
									if x509Cert, err = x509.ParseCertificate(cert); err != nil {
//line /usr/local/go/src/crypto/tls/common.go:1298
				_go_fuzz_dep_.CoverTab[21537]++
										return fmt.Errorf("failed to parse certificate #%d in the chain: %w", j, err)
//line /usr/local/go/src/crypto/tls/common.go:1299
				// _ = "end of CoverTab[21537]"
			} else {
//line /usr/local/go/src/crypto/tls/common.go:1300
				_go_fuzz_dep_.CoverTab[21538]++
//line /usr/local/go/src/crypto/tls/common.go:1300
				// _ = "end of CoverTab[21538]"
//line /usr/local/go/src/crypto/tls/common.go:1300
			}
//line /usr/local/go/src/crypto/tls/common.go:1300
			// _ = "end of CoverTab[21536]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1301
			_go_fuzz_dep_.CoverTab[21539]++
//line /usr/local/go/src/crypto/tls/common.go:1301
			// _ = "end of CoverTab[21539]"
//line /usr/local/go/src/crypto/tls/common.go:1301
		}
//line /usr/local/go/src/crypto/tls/common.go:1301
		// _ = "end of CoverTab[21533]"
//line /usr/local/go/src/crypto/tls/common.go:1301
		_go_fuzz_dep_.CoverTab[21534]++

								for _, ca := range cri.AcceptableCAs {
//line /usr/local/go/src/crypto/tls/common.go:1303
			_go_fuzz_dep_.CoverTab[21540]++
									if bytes.Equal(x509Cert.RawIssuer, ca) {
//line /usr/local/go/src/crypto/tls/common.go:1304
				_go_fuzz_dep_.CoverTab[21541]++
										return nil
//line /usr/local/go/src/crypto/tls/common.go:1305
				// _ = "end of CoverTab[21541]"
			} else {
//line /usr/local/go/src/crypto/tls/common.go:1306
				_go_fuzz_dep_.CoverTab[21542]++
//line /usr/local/go/src/crypto/tls/common.go:1306
				// _ = "end of CoverTab[21542]"
//line /usr/local/go/src/crypto/tls/common.go:1306
			}
//line /usr/local/go/src/crypto/tls/common.go:1306
			// _ = "end of CoverTab[21540]"
		}
//line /usr/local/go/src/crypto/tls/common.go:1307
		// _ = "end of CoverTab[21534]"
	}
//line /usr/local/go/src/crypto/tls/common.go:1308
	// _ = "end of CoverTab[21527]"
//line /usr/local/go/src/crypto/tls/common.go:1308
	_go_fuzz_dep_.CoverTab[21528]++
							return errors.New("chain is not signed by an acceptable CA")
//line /usr/local/go/src/crypto/tls/common.go:1309
	// _ = "end of CoverTab[21528]"
}

//line /usr/local/go/src/crypto/tls/common.go:1319
func (c *Config) BuildNameToCertificate() {
//line /usr/local/go/src/crypto/tls/common.go:1319
	_go_fuzz_dep_.CoverTab[21543]++
							c.NameToCertificate = make(map[string]*Certificate)
							for i := range c.Certificates {
//line /usr/local/go/src/crypto/tls/common.go:1321
		_go_fuzz_dep_.CoverTab[21544]++
								cert := &c.Certificates[i]
								x509Cert, err := cert.leaf()
								if err != nil {
//line /usr/local/go/src/crypto/tls/common.go:1324
			_go_fuzz_dep_.CoverTab[21547]++
									continue
//line /usr/local/go/src/crypto/tls/common.go:1325
			// _ = "end of CoverTab[21547]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1326
			_go_fuzz_dep_.CoverTab[21548]++
//line /usr/local/go/src/crypto/tls/common.go:1326
			// _ = "end of CoverTab[21548]"
//line /usr/local/go/src/crypto/tls/common.go:1326
		}
//line /usr/local/go/src/crypto/tls/common.go:1326
		// _ = "end of CoverTab[21544]"
//line /usr/local/go/src/crypto/tls/common.go:1326
		_go_fuzz_dep_.CoverTab[21545]++

//line /usr/local/go/src/crypto/tls/common.go:1329
		if x509Cert.Subject.CommonName != "" && func() bool {
//line /usr/local/go/src/crypto/tls/common.go:1329
			_go_fuzz_dep_.CoverTab[21549]++
//line /usr/local/go/src/crypto/tls/common.go:1329
			return len(x509Cert.DNSNames) == 0
//line /usr/local/go/src/crypto/tls/common.go:1329
			// _ = "end of CoverTab[21549]"
//line /usr/local/go/src/crypto/tls/common.go:1329
		}() {
//line /usr/local/go/src/crypto/tls/common.go:1329
			_go_fuzz_dep_.CoverTab[21550]++
									c.NameToCertificate[x509Cert.Subject.CommonName] = cert
//line /usr/local/go/src/crypto/tls/common.go:1330
			// _ = "end of CoverTab[21550]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1331
			_go_fuzz_dep_.CoverTab[21551]++
//line /usr/local/go/src/crypto/tls/common.go:1331
			// _ = "end of CoverTab[21551]"
//line /usr/local/go/src/crypto/tls/common.go:1331
		}
//line /usr/local/go/src/crypto/tls/common.go:1331
		// _ = "end of CoverTab[21545]"
//line /usr/local/go/src/crypto/tls/common.go:1331
		_go_fuzz_dep_.CoverTab[21546]++
								for _, san := range x509Cert.DNSNames {
//line /usr/local/go/src/crypto/tls/common.go:1332
			_go_fuzz_dep_.CoverTab[21552]++
									c.NameToCertificate[san] = cert
//line /usr/local/go/src/crypto/tls/common.go:1333
			// _ = "end of CoverTab[21552]"
		}
//line /usr/local/go/src/crypto/tls/common.go:1334
		// _ = "end of CoverTab[21546]"
	}
//line /usr/local/go/src/crypto/tls/common.go:1335
	// _ = "end of CoverTab[21543]"
}

const (
	keyLogLabelTLS12		= "CLIENT_RANDOM"
	keyLogLabelClientHandshake	= "CLIENT_HANDSHAKE_TRAFFIC_SECRET"
	keyLogLabelServerHandshake	= "SERVER_HANDSHAKE_TRAFFIC_SECRET"
	keyLogLabelClientTraffic	= "CLIENT_TRAFFIC_SECRET_0"
	keyLogLabelServerTraffic	= "SERVER_TRAFFIC_SECRET_0"
)

func (c *Config) writeKeyLog(label string, clientRandom, secret []byte) error {
//line /usr/local/go/src/crypto/tls/common.go:1346
	_go_fuzz_dep_.CoverTab[21553]++
							if c.KeyLogWriter == nil {
//line /usr/local/go/src/crypto/tls/common.go:1347
		_go_fuzz_dep_.CoverTab[21555]++
								return nil
//line /usr/local/go/src/crypto/tls/common.go:1348
		// _ = "end of CoverTab[21555]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1349
		_go_fuzz_dep_.CoverTab[21556]++
//line /usr/local/go/src/crypto/tls/common.go:1349
		// _ = "end of CoverTab[21556]"
//line /usr/local/go/src/crypto/tls/common.go:1349
	}
//line /usr/local/go/src/crypto/tls/common.go:1349
	// _ = "end of CoverTab[21553]"
//line /usr/local/go/src/crypto/tls/common.go:1349
	_go_fuzz_dep_.CoverTab[21554]++

							logLine := fmt.Appendf(nil, "%s %x %x\n", label, clientRandom, secret)

							writerMutex.Lock()
							_, err := c.KeyLogWriter.Write(logLine)
							writerMutex.Unlock()

							return err
//line /usr/local/go/src/crypto/tls/common.go:1357
	// _ = "end of CoverTab[21554]"
}

//line /usr/local/go/src/crypto/tls/common.go:1362
var writerMutex sync.Mutex

//line /usr/local/go/src/crypto/tls/common.go:1365
type Certificate struct {
							Certificate	[][]byte

//line /usr/local/go/src/crypto/tls/common.go:1371
	PrivateKey	crypto.PrivateKey

//line /usr/local/go/src/crypto/tls/common.go:1374
	SupportedSignatureAlgorithms	[]SignatureScheme

//line /usr/local/go/src/crypto/tls/common.go:1377
	OCSPStaple	[]byte

//line /usr/local/go/src/crypto/tls/common.go:1380
	SignedCertificateTimestamps	[][]byte

//line /usr/local/go/src/crypto/tls/common.go:1384
	Leaf	*x509.Certificate
}

//line /usr/local/go/src/crypto/tls/common.go:1389
func (c *Certificate) leaf() (*x509.Certificate, error) {
//line /usr/local/go/src/crypto/tls/common.go:1389
	_go_fuzz_dep_.CoverTab[21557]++
							if c.Leaf != nil {
//line /usr/local/go/src/crypto/tls/common.go:1390
		_go_fuzz_dep_.CoverTab[21559]++
								return c.Leaf, nil
//line /usr/local/go/src/crypto/tls/common.go:1391
		// _ = "end of CoverTab[21559]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1392
		_go_fuzz_dep_.CoverTab[21560]++
//line /usr/local/go/src/crypto/tls/common.go:1392
		// _ = "end of CoverTab[21560]"
//line /usr/local/go/src/crypto/tls/common.go:1392
	}
//line /usr/local/go/src/crypto/tls/common.go:1392
	// _ = "end of CoverTab[21557]"
//line /usr/local/go/src/crypto/tls/common.go:1392
	_go_fuzz_dep_.CoverTab[21558]++
							return x509.ParseCertificate(c.Certificate[0])
//line /usr/local/go/src/crypto/tls/common.go:1393
	// _ = "end of CoverTab[21558]"
}

type handshakeMessage interface {
	marshal() ([]byte, error)
	unmarshal([]byte) bool
}

//line /usr/local/go/src/crypto/tls/common.go:1403
type lruSessionCache struct {
	sync.Mutex

	m		map[string]*list.Element
	q		*list.List
	capacity	int
}

type lruSessionCacheEntry struct {
	sessionKey	string
	state		*ClientSessionState
}

//line /usr/local/go/src/crypto/tls/common.go:1419
func NewLRUClientSessionCache(capacity int) ClientSessionCache {
//line /usr/local/go/src/crypto/tls/common.go:1419
	_go_fuzz_dep_.CoverTab[21561]++
							const defaultSessionCacheCapacity = 64

							if capacity < 1 {
//line /usr/local/go/src/crypto/tls/common.go:1422
		_go_fuzz_dep_.CoverTab[21563]++
								capacity = defaultSessionCacheCapacity
//line /usr/local/go/src/crypto/tls/common.go:1423
		// _ = "end of CoverTab[21563]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1424
		_go_fuzz_dep_.CoverTab[21564]++
//line /usr/local/go/src/crypto/tls/common.go:1424
		// _ = "end of CoverTab[21564]"
//line /usr/local/go/src/crypto/tls/common.go:1424
	}
//line /usr/local/go/src/crypto/tls/common.go:1424
	// _ = "end of CoverTab[21561]"
//line /usr/local/go/src/crypto/tls/common.go:1424
	_go_fuzz_dep_.CoverTab[21562]++
							return &lruSessionCache{
		m:		make(map[string]*list.Element),
		q:		list.New(),
		capacity:	capacity,
	}
//line /usr/local/go/src/crypto/tls/common.go:1429
	// _ = "end of CoverTab[21562]"
}

//line /usr/local/go/src/crypto/tls/common.go:1434
func (c *lruSessionCache) Put(sessionKey string, cs *ClientSessionState) {
//line /usr/local/go/src/crypto/tls/common.go:1434
	_go_fuzz_dep_.CoverTab[21565]++
							c.Lock()
							defer c.Unlock()

							if elem, ok := c.m[sessionKey]; ok {
//line /usr/local/go/src/crypto/tls/common.go:1438
		_go_fuzz_dep_.CoverTab[21568]++
								if cs == nil {
//line /usr/local/go/src/crypto/tls/common.go:1439
			_go_fuzz_dep_.CoverTab[21570]++
									c.q.Remove(elem)
									delete(c.m, sessionKey)
//line /usr/local/go/src/crypto/tls/common.go:1441
			// _ = "end of CoverTab[21570]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1442
			_go_fuzz_dep_.CoverTab[21571]++
									entry := elem.Value.(*lruSessionCacheEntry)
									entry.state = cs
									c.q.MoveToFront(elem)
//line /usr/local/go/src/crypto/tls/common.go:1445
			// _ = "end of CoverTab[21571]"
		}
//line /usr/local/go/src/crypto/tls/common.go:1446
		// _ = "end of CoverTab[21568]"
//line /usr/local/go/src/crypto/tls/common.go:1446
		_go_fuzz_dep_.CoverTab[21569]++
								return
//line /usr/local/go/src/crypto/tls/common.go:1447
		// _ = "end of CoverTab[21569]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1448
		_go_fuzz_dep_.CoverTab[21572]++
//line /usr/local/go/src/crypto/tls/common.go:1448
		// _ = "end of CoverTab[21572]"
//line /usr/local/go/src/crypto/tls/common.go:1448
	}
//line /usr/local/go/src/crypto/tls/common.go:1448
	// _ = "end of CoverTab[21565]"
//line /usr/local/go/src/crypto/tls/common.go:1448
	_go_fuzz_dep_.CoverTab[21566]++

							if c.q.Len() < c.capacity {
//line /usr/local/go/src/crypto/tls/common.go:1450
		_go_fuzz_dep_.CoverTab[21573]++
								entry := &lruSessionCacheEntry{sessionKey, cs}
								c.m[sessionKey] = c.q.PushFront(entry)
								return
//line /usr/local/go/src/crypto/tls/common.go:1453
		// _ = "end of CoverTab[21573]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1454
		_go_fuzz_dep_.CoverTab[21574]++
//line /usr/local/go/src/crypto/tls/common.go:1454
		// _ = "end of CoverTab[21574]"
//line /usr/local/go/src/crypto/tls/common.go:1454
	}
//line /usr/local/go/src/crypto/tls/common.go:1454
	// _ = "end of CoverTab[21566]"
//line /usr/local/go/src/crypto/tls/common.go:1454
	_go_fuzz_dep_.CoverTab[21567]++

							elem := c.q.Back()
							entry := elem.Value.(*lruSessionCacheEntry)
							delete(c.m, entry.sessionKey)
							entry.sessionKey = sessionKey
							entry.state = cs
							c.q.MoveToFront(elem)
							c.m[sessionKey] = elem
//line /usr/local/go/src/crypto/tls/common.go:1462
	// _ = "end of CoverTab[21567]"
}

//line /usr/local/go/src/crypto/tls/common.go:1467
func (c *lruSessionCache) Get(sessionKey string) (*ClientSessionState, bool) {
//line /usr/local/go/src/crypto/tls/common.go:1467
	_go_fuzz_dep_.CoverTab[21575]++
							c.Lock()
							defer c.Unlock()

							if elem, ok := c.m[sessionKey]; ok {
//line /usr/local/go/src/crypto/tls/common.go:1471
		_go_fuzz_dep_.CoverTab[21577]++
								c.q.MoveToFront(elem)
								return elem.Value.(*lruSessionCacheEntry).state, true
//line /usr/local/go/src/crypto/tls/common.go:1473
		// _ = "end of CoverTab[21577]"
	} else {
//line /usr/local/go/src/crypto/tls/common.go:1474
		_go_fuzz_dep_.CoverTab[21578]++
//line /usr/local/go/src/crypto/tls/common.go:1474
		// _ = "end of CoverTab[21578]"
//line /usr/local/go/src/crypto/tls/common.go:1474
	}
//line /usr/local/go/src/crypto/tls/common.go:1474
	// _ = "end of CoverTab[21575]"
//line /usr/local/go/src/crypto/tls/common.go:1474
	_go_fuzz_dep_.CoverTab[21576]++
							return nil, false
//line /usr/local/go/src/crypto/tls/common.go:1475
	// _ = "end of CoverTab[21576]"
}

var emptyConfig Config

func defaultConfig() *Config {
//line /usr/local/go/src/crypto/tls/common.go:1480
	_go_fuzz_dep_.CoverTab[21579]++
							return &emptyConfig
//line /usr/local/go/src/crypto/tls/common.go:1481
	// _ = "end of CoverTab[21579]"
}

func unexpectedMessageError(wanted, got any) error {
//line /usr/local/go/src/crypto/tls/common.go:1484
	_go_fuzz_dep_.CoverTab[21580]++
							return fmt.Errorf("tls: received unexpected handshake message of type %T when waiting for %T", got, wanted)
//line /usr/local/go/src/crypto/tls/common.go:1485
	// _ = "end of CoverTab[21580]"
}

func isSupportedSignatureAlgorithm(sigAlg SignatureScheme, supportedSignatureAlgorithms []SignatureScheme) bool {
//line /usr/local/go/src/crypto/tls/common.go:1488
	_go_fuzz_dep_.CoverTab[21581]++
							for _, s := range supportedSignatureAlgorithms {
//line /usr/local/go/src/crypto/tls/common.go:1489
		_go_fuzz_dep_.CoverTab[21583]++
								if s == sigAlg {
//line /usr/local/go/src/crypto/tls/common.go:1490
			_go_fuzz_dep_.CoverTab[21584]++
									return true
//line /usr/local/go/src/crypto/tls/common.go:1491
			// _ = "end of CoverTab[21584]"
		} else {
//line /usr/local/go/src/crypto/tls/common.go:1492
			_go_fuzz_dep_.CoverTab[21585]++
//line /usr/local/go/src/crypto/tls/common.go:1492
			// _ = "end of CoverTab[21585]"
//line /usr/local/go/src/crypto/tls/common.go:1492
		}
//line /usr/local/go/src/crypto/tls/common.go:1492
		// _ = "end of CoverTab[21583]"
	}
//line /usr/local/go/src/crypto/tls/common.go:1493
	// _ = "end of CoverTab[21581]"
//line /usr/local/go/src/crypto/tls/common.go:1493
	_go_fuzz_dep_.CoverTab[21582]++
							return false
//line /usr/local/go/src/crypto/tls/common.go:1494
	// _ = "end of CoverTab[21582]"
}

//line /usr/local/go/src/crypto/tls/common.go:1498
type CertificateVerificationError struct {
//line /usr/local/go/src/crypto/tls/common.go:1500
	UnverifiedCertificates	[]*x509.Certificate
	Err	error
}

func (e *CertificateVerificationError) Error() string {
//line /usr/local/go/src/crypto/tls/common.go:1504
	_go_fuzz_dep_.CoverTab[21586]++
							return fmt.Sprintf("tls: failed to verify certificate: %s", e.Err)
//line /usr/local/go/src/crypto/tls/common.go:1505
	// _ = "end of CoverTab[21586]"
}

func (e *CertificateVerificationError) Unwrap() error {
//line /usr/local/go/src/crypto/tls/common.go:1508
	_go_fuzz_dep_.CoverTab[21587]++
							return e.Err
//line /usr/local/go/src/crypto/tls/common.go:1509
	// _ = "end of CoverTab[21587]"
}

//line /usr/local/go/src/crypto/tls/common.go:1510
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/common.go:1510
var _ = _go_fuzz_dep_.CoverTab
