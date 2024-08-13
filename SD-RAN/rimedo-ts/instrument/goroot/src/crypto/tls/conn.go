// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TLS low level connection and record layer

//line /usr/local/go/src/crypto/tls/conn.go:7
package tls

//line /usr/local/go/src/crypto/tls/conn.go:7
import (
//line /usr/local/go/src/crypto/tls/conn.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/conn.go:7
)
//line /usr/local/go/src/crypto/tls/conn.go:7
import (
//line /usr/local/go/src/crypto/tls/conn.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/conn.go:7
)

import (
	"bytes"
	"context"
	"crypto/cipher"
	"crypto/subtle"
	"crypto/x509"
	"errors"
	"fmt"
	"hash"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

// A Conn represents a secured connection.
//line /usr/local/go/src/crypto/tls/conn.go:25
// It implements the net.Conn interface.
//line /usr/local/go/src/crypto/tls/conn.go:27
type Conn struct {
	// constant
	conn		net.Conn
	isClient	bool
	handshakeFn	func(context.Context) error	// (*Conn).clientHandshake or serverHandshake

	// isHandshakeComplete is true if the connection is currently transferring
	// application data (i.e. is not currently processing a handshake).
	// isHandshakeComplete is true implies handshakeErr == nil.
	isHandshakeComplete	atomic.Bool
	// constant after handshake; protected by handshakeMutex
	handshakeMutex	sync.Mutex
	handshakeErr	error	// error resulting from handshake
	vers		uint16	// TLS version
	haveVers	bool	// version has been negotiated
	config		*Config	// configuration passed to constructor
	// handshakes counts the number of handshakes performed on the
	// connection so far. If renegotiation is disabled then this is either
	// zero or one.
	handshakes		int
	didResume		bool	// whether this connection was a session resumption
	cipherSuite		uint16
	ocspResponse		[]byte		// stapled OCSP response
	scts			[][]byte	// signed certificate timestamps from server
	peerCertificates	[]*x509.Certificate
	// activeCertHandles contains the cache handles to certificates in
	// peerCertificates that are used to track active references.
	activeCertHandles	[]*activeCert
	// verifiedChains contains the certificate chains that we built, as
	// opposed to the ones presented by the server.
	verifiedChains	[][]*x509.Certificate
	// serverName contains the server name indicated by the client, if any.
	serverName	string
	// secureRenegotiation is true if the server echoed the secure
	// renegotiation extension. (This is meaningless as a server because
	// renegotiation is not supported in that case.)
	secureRenegotiation	bool
	// ekm is a closure for exporting keying material.
	ekm	func(label string, context []byte, length int) ([]byte, error)
	// resumptionSecret is the resumption_master_secret for handling
	// NewSessionTicket messages. nil if config.SessionTicketsDisabled.
	resumptionSecret	[]byte

	// ticketKeys is the set of active session ticket keys for this
	// connection. The first one is used to encrypt new tickets and
	// all are tried to decrypt tickets.
	ticketKeys	[]ticketKey

	// clientFinishedIsFirst is true if the client sent the first Finished
	// message during the most recent handshake. This is recorded because
	// the first transmitted Finished message is the tls-unique
	// channel-binding value.
	clientFinishedIsFirst	bool

	// closeNotifyErr is any error from sending the alertCloseNotify record.
	closeNotifyErr	error
	// closeNotifySent is true if the Conn attempted to send an
	// alertCloseNotify record.
	closeNotifySent	bool

	// clientFinished and serverFinished contain the Finished message sent
	// by the client or server in the most recent handshake. This is
	// retained to support the renegotiation extension and tls-unique
	// channel-binding.
	clientFinished	[12]byte
	serverFinished	[12]byte

	// clientProtocol is the negotiated ALPN protocol.
	clientProtocol	string

	// input/output
	in, out		halfConn
	rawInput	bytes.Buffer	// raw input, starting with a record header
	input		bytes.Reader	// application data waiting to be read, from rawInput.Next
	hand		bytes.Buffer	// handshake data waiting to be read
	buffering	bool		// whether records are buffered in sendBuf
	sendBuf		[]byte		// a buffer of records waiting to be sent

	// bytesSent counts the bytes of application data sent.
	// packetsSent counts packets.
	bytesSent	int64
	packetsSent	int64

	// retryCount counts the number of consecutive non-advancing records
	// received by Conn.readRecord. That is, records that neither advance the
	// handshake, nor deliver application data. Protected by in.Mutex.
	retryCount	int

	// activeCall indicates whether Close has been call in the low bit.
	// the rest of the bits are the number of goroutines in Conn.Write.
	activeCall	atomic.Int32

	tmp	[16]byte
}

//line /usr/local/go/src/crypto/tls/conn.go:126
// LocalAddr returns the local network address.
func (c *Conn) LocalAddr() net.Addr {
//line /usr/local/go/src/crypto/tls/conn.go:127
	_go_fuzz_dep_.CoverTab[21613]++
							return c.conn.LocalAddr()
//line /usr/local/go/src/crypto/tls/conn.go:128
	// _ = "end of CoverTab[21613]"
}

// RemoteAddr returns the remote network address.
func (c *Conn) RemoteAddr() net.Addr {
//line /usr/local/go/src/crypto/tls/conn.go:132
	_go_fuzz_dep_.CoverTab[21614]++
							return c.conn.RemoteAddr()
//line /usr/local/go/src/crypto/tls/conn.go:133
	// _ = "end of CoverTab[21614]"
}

// SetDeadline sets the read and write deadlines associated with the connection.
//line /usr/local/go/src/crypto/tls/conn.go:136
// A zero value for t means Read and Write will not time out.
//line /usr/local/go/src/crypto/tls/conn.go:136
// After a Write has timed out, the TLS state is corrupt and all future writes will return the same error.
//line /usr/local/go/src/crypto/tls/conn.go:139
func (c *Conn) SetDeadline(t time.Time) error {
//line /usr/local/go/src/crypto/tls/conn.go:139
	_go_fuzz_dep_.CoverTab[21615]++
							return c.conn.SetDeadline(t)
//line /usr/local/go/src/crypto/tls/conn.go:140
	// _ = "end of CoverTab[21615]"
}

// SetReadDeadline sets the read deadline on the underlying connection.
//line /usr/local/go/src/crypto/tls/conn.go:143
// A zero value for t means Read will not time out.
//line /usr/local/go/src/crypto/tls/conn.go:145
func (c *Conn) SetReadDeadline(t time.Time) error {
//line /usr/local/go/src/crypto/tls/conn.go:145
	_go_fuzz_dep_.CoverTab[21616]++
							return c.conn.SetReadDeadline(t)
//line /usr/local/go/src/crypto/tls/conn.go:146
	// _ = "end of CoverTab[21616]"
}

// SetWriteDeadline sets the write deadline on the underlying connection.
//line /usr/local/go/src/crypto/tls/conn.go:149
// A zero value for t means Write will not time out.
//line /usr/local/go/src/crypto/tls/conn.go:149
// After a Write has timed out, the TLS state is corrupt and all future writes will return the same error.
//line /usr/local/go/src/crypto/tls/conn.go:152
func (c *Conn) SetWriteDeadline(t time.Time) error {
//line /usr/local/go/src/crypto/tls/conn.go:152
	_go_fuzz_dep_.CoverTab[21617]++
							return c.conn.SetWriteDeadline(t)
//line /usr/local/go/src/crypto/tls/conn.go:153
	// _ = "end of CoverTab[21617]"
}

// NetConn returns the underlying connection that is wrapped by c.
//line /usr/local/go/src/crypto/tls/conn.go:156
// Note that writing to or reading from this connection directly will corrupt the
//line /usr/local/go/src/crypto/tls/conn.go:156
// TLS session.
//line /usr/local/go/src/crypto/tls/conn.go:159
func (c *Conn) NetConn() net.Conn {
//line /usr/local/go/src/crypto/tls/conn.go:159
	_go_fuzz_dep_.CoverTab[21618]++
							return c.conn
//line /usr/local/go/src/crypto/tls/conn.go:160
	// _ = "end of CoverTab[21618]"
}

// A halfConn represents one direction of the record layer
//line /usr/local/go/src/crypto/tls/conn.go:163
// connection, either sending or receiving.
//line /usr/local/go/src/crypto/tls/conn.go:165
type halfConn struct {
	sync.Mutex

	err	error	// first permanent error
	version	uint16	// protocol version
	cipher	any	// cipher algorithm
	mac	hash.Hash
	seq	[8]byte	// 64-bit sequence number

	scratchBuf	[13]byte	// to avoid allocs; interface method args escape

	nextCipher	any		// next encryption state
	nextMac		hash.Hash	// next MAC algorithm

	trafficSecret	[]byte	// current TLS 1.3 traffic secret
}

type permanentError struct {
	err net.Error
}

func (e *permanentError) Error() string {
//line /usr/local/go/src/crypto/tls/conn.go:186
	_go_fuzz_dep_.CoverTab[21619]++
//line /usr/local/go/src/crypto/tls/conn.go:186
	return e.err.Error()
//line /usr/local/go/src/crypto/tls/conn.go:186
	// _ = "end of CoverTab[21619]"
//line /usr/local/go/src/crypto/tls/conn.go:186
}
func (e *permanentError) Unwrap() error {
//line /usr/local/go/src/crypto/tls/conn.go:187
	_go_fuzz_dep_.CoverTab[21620]++
//line /usr/local/go/src/crypto/tls/conn.go:187
	return e.err
//line /usr/local/go/src/crypto/tls/conn.go:187
	// _ = "end of CoverTab[21620]"
//line /usr/local/go/src/crypto/tls/conn.go:187
}
func (e *permanentError) Timeout() bool {
//line /usr/local/go/src/crypto/tls/conn.go:188
	_go_fuzz_dep_.CoverTab[21621]++
//line /usr/local/go/src/crypto/tls/conn.go:188
	return e.err.Timeout()
//line /usr/local/go/src/crypto/tls/conn.go:188
	// _ = "end of CoverTab[21621]"
//line /usr/local/go/src/crypto/tls/conn.go:188
}
func (e *permanentError) Temporary() bool {
//line /usr/local/go/src/crypto/tls/conn.go:189
	_go_fuzz_dep_.CoverTab[21622]++
//line /usr/local/go/src/crypto/tls/conn.go:189
	return false
//line /usr/local/go/src/crypto/tls/conn.go:189
	// _ = "end of CoverTab[21622]"
//line /usr/local/go/src/crypto/tls/conn.go:189
}

func (hc *halfConn) setErrorLocked(err error) error {
//line /usr/local/go/src/crypto/tls/conn.go:191
	_go_fuzz_dep_.CoverTab[21623]++
							if e, ok := err.(net.Error); ok {
//line /usr/local/go/src/crypto/tls/conn.go:192
		_go_fuzz_dep_.CoverTab[21625]++
								hc.err = &permanentError{err: e}
//line /usr/local/go/src/crypto/tls/conn.go:193
		// _ = "end of CoverTab[21625]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:194
		_go_fuzz_dep_.CoverTab[21626]++
								hc.err = err
//line /usr/local/go/src/crypto/tls/conn.go:195
		// _ = "end of CoverTab[21626]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:196
	// _ = "end of CoverTab[21623]"
//line /usr/local/go/src/crypto/tls/conn.go:196
	_go_fuzz_dep_.CoverTab[21624]++
							return hc.err
//line /usr/local/go/src/crypto/tls/conn.go:197
	// _ = "end of CoverTab[21624]"
}

// prepareCipherSpec sets the encryption and MAC states
//line /usr/local/go/src/crypto/tls/conn.go:200
// that a subsequent changeCipherSpec will use.
//line /usr/local/go/src/crypto/tls/conn.go:202
func (hc *halfConn) prepareCipherSpec(version uint16, cipher any, mac hash.Hash) {
//line /usr/local/go/src/crypto/tls/conn.go:202
	_go_fuzz_dep_.CoverTab[21627]++
							hc.version = version
							hc.nextCipher = cipher
							hc.nextMac = mac
//line /usr/local/go/src/crypto/tls/conn.go:205
	// _ = "end of CoverTab[21627]"
}

// changeCipherSpec changes the encryption and MAC states
//line /usr/local/go/src/crypto/tls/conn.go:208
// to the ones previously passed to prepareCipherSpec.
//line /usr/local/go/src/crypto/tls/conn.go:210
func (hc *halfConn) changeCipherSpec() error {
//line /usr/local/go/src/crypto/tls/conn.go:210
	_go_fuzz_dep_.CoverTab[21628]++
							if hc.nextCipher == nil || func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:211
		_go_fuzz_dep_.CoverTab[21631]++
//line /usr/local/go/src/crypto/tls/conn.go:211
		return hc.version == VersionTLS13
//line /usr/local/go/src/crypto/tls/conn.go:211
		// _ = "end of CoverTab[21631]"
//line /usr/local/go/src/crypto/tls/conn.go:211
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:211
		_go_fuzz_dep_.CoverTab[21632]++
								return alertInternalError
//line /usr/local/go/src/crypto/tls/conn.go:212
		// _ = "end of CoverTab[21632]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:213
		_go_fuzz_dep_.CoverTab[21633]++
//line /usr/local/go/src/crypto/tls/conn.go:213
		// _ = "end of CoverTab[21633]"
//line /usr/local/go/src/crypto/tls/conn.go:213
	}
//line /usr/local/go/src/crypto/tls/conn.go:213
	// _ = "end of CoverTab[21628]"
//line /usr/local/go/src/crypto/tls/conn.go:213
	_go_fuzz_dep_.CoverTab[21629]++
							hc.cipher = hc.nextCipher
							hc.mac = hc.nextMac
							hc.nextCipher = nil
							hc.nextMac = nil
							for i := range hc.seq {
//line /usr/local/go/src/crypto/tls/conn.go:218
		_go_fuzz_dep_.CoverTab[21634]++
								hc.seq[i] = 0
//line /usr/local/go/src/crypto/tls/conn.go:219
		// _ = "end of CoverTab[21634]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:220
	// _ = "end of CoverTab[21629]"
//line /usr/local/go/src/crypto/tls/conn.go:220
	_go_fuzz_dep_.CoverTab[21630]++
							return nil
//line /usr/local/go/src/crypto/tls/conn.go:221
	// _ = "end of CoverTab[21630]"
}

func (hc *halfConn) setTrafficSecret(suite *cipherSuiteTLS13, secret []byte) {
//line /usr/local/go/src/crypto/tls/conn.go:224
	_go_fuzz_dep_.CoverTab[21635]++
							hc.trafficSecret = secret
							key, iv := suite.trafficKey(secret)
							hc.cipher = suite.aead(key, iv)
							for i := range hc.seq {
//line /usr/local/go/src/crypto/tls/conn.go:228
		_go_fuzz_dep_.CoverTab[21636]++
								hc.seq[i] = 0
//line /usr/local/go/src/crypto/tls/conn.go:229
		// _ = "end of CoverTab[21636]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:230
	// _ = "end of CoverTab[21635]"
}

// incSeq increments the sequence number.
func (hc *halfConn) incSeq() {
//line /usr/local/go/src/crypto/tls/conn.go:234
	_go_fuzz_dep_.CoverTab[21637]++
							for i := 7; i >= 0; i-- {
//line /usr/local/go/src/crypto/tls/conn.go:235
		_go_fuzz_dep_.CoverTab[21639]++
								hc.seq[i]++
								if hc.seq[i] != 0 {
//line /usr/local/go/src/crypto/tls/conn.go:237
			_go_fuzz_dep_.CoverTab[21640]++
									return
//line /usr/local/go/src/crypto/tls/conn.go:238
			// _ = "end of CoverTab[21640]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:239
			_go_fuzz_dep_.CoverTab[21641]++
//line /usr/local/go/src/crypto/tls/conn.go:239
			// _ = "end of CoverTab[21641]"
//line /usr/local/go/src/crypto/tls/conn.go:239
		}
//line /usr/local/go/src/crypto/tls/conn.go:239
		// _ = "end of CoverTab[21639]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:240
	// _ = "end of CoverTab[21637]"
//line /usr/local/go/src/crypto/tls/conn.go:240
	_go_fuzz_dep_.CoverTab[21638]++

//line /usr/local/go/src/crypto/tls/conn.go:245
	panic("TLS: sequence number wraparound")
//line /usr/local/go/src/crypto/tls/conn.go:245
	// _ = "end of CoverTab[21638]"
}

// explicitNonceLen returns the number of bytes of explicit nonce or IV included
//line /usr/local/go/src/crypto/tls/conn.go:248
// in each record. Explicit nonces are present only in CBC modes after TLS 1.0
//line /usr/local/go/src/crypto/tls/conn.go:248
// and in certain AEAD modes in TLS 1.2.
//line /usr/local/go/src/crypto/tls/conn.go:251
func (hc *halfConn) explicitNonceLen() int {
//line /usr/local/go/src/crypto/tls/conn.go:251
	_go_fuzz_dep_.CoverTab[21642]++
							if hc.cipher == nil {
//line /usr/local/go/src/crypto/tls/conn.go:252
		_go_fuzz_dep_.CoverTab[21644]++
								return 0
//line /usr/local/go/src/crypto/tls/conn.go:253
		// _ = "end of CoverTab[21644]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:254
		_go_fuzz_dep_.CoverTab[21645]++
//line /usr/local/go/src/crypto/tls/conn.go:254
		// _ = "end of CoverTab[21645]"
//line /usr/local/go/src/crypto/tls/conn.go:254
	}
//line /usr/local/go/src/crypto/tls/conn.go:254
	// _ = "end of CoverTab[21642]"
//line /usr/local/go/src/crypto/tls/conn.go:254
	_go_fuzz_dep_.CoverTab[21643]++

							switch c := hc.cipher.(type) {
	case cipher.Stream:
//line /usr/local/go/src/crypto/tls/conn.go:257
		_go_fuzz_dep_.CoverTab[21646]++
								return 0
//line /usr/local/go/src/crypto/tls/conn.go:258
		// _ = "end of CoverTab[21646]"
	case aead:
//line /usr/local/go/src/crypto/tls/conn.go:259
		_go_fuzz_dep_.CoverTab[21647]++
								return c.explicitNonceLen()
//line /usr/local/go/src/crypto/tls/conn.go:260
		// _ = "end of CoverTab[21647]"
	case cbcMode:
//line /usr/local/go/src/crypto/tls/conn.go:261
		_go_fuzz_dep_.CoverTab[21648]++

								if hc.version >= VersionTLS11 {
//line /usr/local/go/src/crypto/tls/conn.go:263
			_go_fuzz_dep_.CoverTab[21651]++
									return c.BlockSize()
//line /usr/local/go/src/crypto/tls/conn.go:264
			// _ = "end of CoverTab[21651]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:265
			_go_fuzz_dep_.CoverTab[21652]++
//line /usr/local/go/src/crypto/tls/conn.go:265
			// _ = "end of CoverTab[21652]"
//line /usr/local/go/src/crypto/tls/conn.go:265
		}
//line /usr/local/go/src/crypto/tls/conn.go:265
		// _ = "end of CoverTab[21648]"
//line /usr/local/go/src/crypto/tls/conn.go:265
		_go_fuzz_dep_.CoverTab[21649]++
								return 0
//line /usr/local/go/src/crypto/tls/conn.go:266
		// _ = "end of CoverTab[21649]"
	default:
//line /usr/local/go/src/crypto/tls/conn.go:267
		_go_fuzz_dep_.CoverTab[21650]++
								panic("unknown cipher type")
//line /usr/local/go/src/crypto/tls/conn.go:268
		// _ = "end of CoverTab[21650]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:269
	// _ = "end of CoverTab[21643]"
}

// extractPadding returns, in constant time, the length of the padding to remove
//line /usr/local/go/src/crypto/tls/conn.go:272
// from the end of payload. It also returns a byte which is equal to 255 if the
//line /usr/local/go/src/crypto/tls/conn.go:272
// padding was valid and 0 otherwise. See RFC 2246, Section 6.2.3.2.
//line /usr/local/go/src/crypto/tls/conn.go:275
func extractPadding(payload []byte) (toRemove int, good byte) {
//line /usr/local/go/src/crypto/tls/conn.go:275
	_go_fuzz_dep_.CoverTab[21653]++
							if len(payload) < 1 {
//line /usr/local/go/src/crypto/tls/conn.go:276
		_go_fuzz_dep_.CoverTab[21657]++
								return 0, 0
//line /usr/local/go/src/crypto/tls/conn.go:277
		// _ = "end of CoverTab[21657]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:278
		_go_fuzz_dep_.CoverTab[21658]++
//line /usr/local/go/src/crypto/tls/conn.go:278
		// _ = "end of CoverTab[21658]"
//line /usr/local/go/src/crypto/tls/conn.go:278
	}
//line /usr/local/go/src/crypto/tls/conn.go:278
	// _ = "end of CoverTab[21653]"
//line /usr/local/go/src/crypto/tls/conn.go:278
	_go_fuzz_dep_.CoverTab[21654]++

							paddingLen := payload[len(payload)-1]
							t := uint(len(payload)-1) - uint(paddingLen)

							good = byte(int32(^t) >> 31)

//line /usr/local/go/src/crypto/tls/conn.go:286
	toCheck := 256

	if toCheck > len(payload) {
//line /usr/local/go/src/crypto/tls/conn.go:288
		_go_fuzz_dep_.CoverTab[21659]++
								toCheck = len(payload)
//line /usr/local/go/src/crypto/tls/conn.go:289
		// _ = "end of CoverTab[21659]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:290
		_go_fuzz_dep_.CoverTab[21660]++
//line /usr/local/go/src/crypto/tls/conn.go:290
		// _ = "end of CoverTab[21660]"
//line /usr/local/go/src/crypto/tls/conn.go:290
	}
//line /usr/local/go/src/crypto/tls/conn.go:290
	// _ = "end of CoverTab[21654]"
//line /usr/local/go/src/crypto/tls/conn.go:290
	_go_fuzz_dep_.CoverTab[21655]++

							for i := 0; i < toCheck; i++ {
//line /usr/local/go/src/crypto/tls/conn.go:292
		_go_fuzz_dep_.CoverTab[21661]++
								t := uint(paddingLen) - uint(i)

								mask := byte(int32(^t) >> 31)
								b := payload[len(payload)-1-i]
								good &^= mask&paddingLen ^ mask&b
//line /usr/local/go/src/crypto/tls/conn.go:297
		// _ = "end of CoverTab[21661]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:298
	// _ = "end of CoverTab[21655]"
//line /usr/local/go/src/crypto/tls/conn.go:298
	_go_fuzz_dep_.CoverTab[21656]++

//line /usr/local/go/src/crypto/tls/conn.go:302
	good &= good << 4
							good &= good << 2
							good &= good << 1
							good = uint8(int8(good) >> 7)

//line /usr/local/go/src/crypto/tls/conn.go:316
	paddingLen &= good

							toRemove = int(paddingLen) + 1
							return
//line /usr/local/go/src/crypto/tls/conn.go:319
	// _ = "end of CoverTab[21656]"
}

func roundUp(a, b int) int {
//line /usr/local/go/src/crypto/tls/conn.go:322
	_go_fuzz_dep_.CoverTab[21662]++
							return a + (b-a%b)%b
//line /usr/local/go/src/crypto/tls/conn.go:323
	// _ = "end of CoverTab[21662]"
}

// cbcMode is an interface for block ciphers using cipher block chaining.
type cbcMode interface {
	cipher.BlockMode
	SetIV([]byte)
}

// decrypt authenticates and decrypts the record if protection is active at
//line /usr/local/go/src/crypto/tls/conn.go:332
// this stage. The returned plaintext might overlap with the input.
//line /usr/local/go/src/crypto/tls/conn.go:334
func (hc *halfConn) decrypt(record []byte) ([]byte, recordType, error) {
//line /usr/local/go/src/crypto/tls/conn.go:334
	_go_fuzz_dep_.CoverTab[21663]++
							var plaintext []byte
							typ := recordType(record[0])
							payload := record[recordHeaderLen:]

//line /usr/local/go/src/crypto/tls/conn.go:341
	if hc.version == VersionTLS13 && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:341
		_go_fuzz_dep_.CoverTab[21667]++
//line /usr/local/go/src/crypto/tls/conn.go:341
		return typ == recordTypeChangeCipherSpec
//line /usr/local/go/src/crypto/tls/conn.go:341
		// _ = "end of CoverTab[21667]"
//line /usr/local/go/src/crypto/tls/conn.go:341
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:341
		_go_fuzz_dep_.CoverTab[21668]++
								return payload, typ, nil
//line /usr/local/go/src/crypto/tls/conn.go:342
		// _ = "end of CoverTab[21668]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:343
		_go_fuzz_dep_.CoverTab[21669]++
//line /usr/local/go/src/crypto/tls/conn.go:343
		// _ = "end of CoverTab[21669]"
//line /usr/local/go/src/crypto/tls/conn.go:343
	}
//line /usr/local/go/src/crypto/tls/conn.go:343
	// _ = "end of CoverTab[21663]"
//line /usr/local/go/src/crypto/tls/conn.go:343
	_go_fuzz_dep_.CoverTab[21664]++

							paddingGood := byte(255)
							paddingLen := 0

							explicitNonceLen := hc.explicitNonceLen()

							if hc.cipher != nil {
//line /usr/local/go/src/crypto/tls/conn.go:350
		_go_fuzz_dep_.CoverTab[21670]++
								switch c := hc.cipher.(type) {
		case cipher.Stream:
//line /usr/local/go/src/crypto/tls/conn.go:352
			_go_fuzz_dep_.CoverTab[21672]++
									c.XORKeyStream(payload, payload)
//line /usr/local/go/src/crypto/tls/conn.go:353
			// _ = "end of CoverTab[21672]"
		case aead:
//line /usr/local/go/src/crypto/tls/conn.go:354
			_go_fuzz_dep_.CoverTab[21673]++
									if len(payload) < explicitNonceLen {
//line /usr/local/go/src/crypto/tls/conn.go:355
				_go_fuzz_dep_.CoverTab[21681]++
										return nil, 0, alertBadRecordMAC
//line /usr/local/go/src/crypto/tls/conn.go:356
				// _ = "end of CoverTab[21681]"
			} else {
//line /usr/local/go/src/crypto/tls/conn.go:357
				_go_fuzz_dep_.CoverTab[21682]++
//line /usr/local/go/src/crypto/tls/conn.go:357
				// _ = "end of CoverTab[21682]"
//line /usr/local/go/src/crypto/tls/conn.go:357
			}
//line /usr/local/go/src/crypto/tls/conn.go:357
			// _ = "end of CoverTab[21673]"
//line /usr/local/go/src/crypto/tls/conn.go:357
			_go_fuzz_dep_.CoverTab[21674]++
									nonce := payload[:explicitNonceLen]
									if len(nonce) == 0 {
//line /usr/local/go/src/crypto/tls/conn.go:359
				_go_fuzz_dep_.CoverTab[21683]++
										nonce = hc.seq[:]
//line /usr/local/go/src/crypto/tls/conn.go:360
				// _ = "end of CoverTab[21683]"
			} else {
//line /usr/local/go/src/crypto/tls/conn.go:361
				_go_fuzz_dep_.CoverTab[21684]++
//line /usr/local/go/src/crypto/tls/conn.go:361
				// _ = "end of CoverTab[21684]"
//line /usr/local/go/src/crypto/tls/conn.go:361
			}
//line /usr/local/go/src/crypto/tls/conn.go:361
			// _ = "end of CoverTab[21674]"
//line /usr/local/go/src/crypto/tls/conn.go:361
			_go_fuzz_dep_.CoverTab[21675]++
									payload = payload[explicitNonceLen:]

									var additionalData []byte
									if hc.version == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/conn.go:365
				_go_fuzz_dep_.CoverTab[21685]++
										additionalData = record[:recordHeaderLen]
//line /usr/local/go/src/crypto/tls/conn.go:366
				// _ = "end of CoverTab[21685]"
			} else {
//line /usr/local/go/src/crypto/tls/conn.go:367
				_go_fuzz_dep_.CoverTab[21686]++
										additionalData = append(hc.scratchBuf[:0], hc.seq[:]...)
										additionalData = append(additionalData, record[:3]...)
										n := len(payload) - c.Overhead()
										additionalData = append(additionalData, byte(n>>8), byte(n))
//line /usr/local/go/src/crypto/tls/conn.go:371
				// _ = "end of CoverTab[21686]"
			}
//line /usr/local/go/src/crypto/tls/conn.go:372
			// _ = "end of CoverTab[21675]"
//line /usr/local/go/src/crypto/tls/conn.go:372
			_go_fuzz_dep_.CoverTab[21676]++

									var err error
									plaintext, err = c.Open(payload[:0], nonce, payload, additionalData)
									if err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:376
				_go_fuzz_dep_.CoverTab[21687]++
										return nil, 0, alertBadRecordMAC
//line /usr/local/go/src/crypto/tls/conn.go:377
				// _ = "end of CoverTab[21687]"
			} else {
//line /usr/local/go/src/crypto/tls/conn.go:378
				_go_fuzz_dep_.CoverTab[21688]++
//line /usr/local/go/src/crypto/tls/conn.go:378
				// _ = "end of CoverTab[21688]"
//line /usr/local/go/src/crypto/tls/conn.go:378
			}
//line /usr/local/go/src/crypto/tls/conn.go:378
			// _ = "end of CoverTab[21676]"
		case cbcMode:
//line /usr/local/go/src/crypto/tls/conn.go:379
			_go_fuzz_dep_.CoverTab[21677]++
									blockSize := c.BlockSize()
									minPayload := explicitNonceLen + roundUp(hc.mac.Size()+1, blockSize)
									if len(payload)%blockSize != 0 || func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:382
				_go_fuzz_dep_.CoverTab[21689]++
//line /usr/local/go/src/crypto/tls/conn.go:382
				return len(payload) < minPayload
//line /usr/local/go/src/crypto/tls/conn.go:382
				// _ = "end of CoverTab[21689]"
//line /usr/local/go/src/crypto/tls/conn.go:382
			}() {
//line /usr/local/go/src/crypto/tls/conn.go:382
				_go_fuzz_dep_.CoverTab[21690]++
										return nil, 0, alertBadRecordMAC
//line /usr/local/go/src/crypto/tls/conn.go:383
				// _ = "end of CoverTab[21690]"
			} else {
//line /usr/local/go/src/crypto/tls/conn.go:384
				_go_fuzz_dep_.CoverTab[21691]++
//line /usr/local/go/src/crypto/tls/conn.go:384
				// _ = "end of CoverTab[21691]"
//line /usr/local/go/src/crypto/tls/conn.go:384
			}
//line /usr/local/go/src/crypto/tls/conn.go:384
			// _ = "end of CoverTab[21677]"
//line /usr/local/go/src/crypto/tls/conn.go:384
			_go_fuzz_dep_.CoverTab[21678]++

									if explicitNonceLen > 0 {
//line /usr/local/go/src/crypto/tls/conn.go:386
				_go_fuzz_dep_.CoverTab[21692]++
										c.SetIV(payload[:explicitNonceLen])
										payload = payload[explicitNonceLen:]
//line /usr/local/go/src/crypto/tls/conn.go:388
				// _ = "end of CoverTab[21692]"
			} else {
//line /usr/local/go/src/crypto/tls/conn.go:389
				_go_fuzz_dep_.CoverTab[21693]++
//line /usr/local/go/src/crypto/tls/conn.go:389
				// _ = "end of CoverTab[21693]"
//line /usr/local/go/src/crypto/tls/conn.go:389
			}
//line /usr/local/go/src/crypto/tls/conn.go:389
			// _ = "end of CoverTab[21678]"
//line /usr/local/go/src/crypto/tls/conn.go:389
			_go_fuzz_dep_.CoverTab[21679]++
									c.CryptBlocks(payload, payload)

//line /usr/local/go/src/crypto/tls/conn.go:398
			paddingLen, paddingGood = extractPadding(payload)
//line /usr/local/go/src/crypto/tls/conn.go:398
			// _ = "end of CoverTab[21679]"
		default:
//line /usr/local/go/src/crypto/tls/conn.go:399
			_go_fuzz_dep_.CoverTab[21680]++
									panic("unknown cipher type")
//line /usr/local/go/src/crypto/tls/conn.go:400
			// _ = "end of CoverTab[21680]"
		}
//line /usr/local/go/src/crypto/tls/conn.go:401
		// _ = "end of CoverTab[21670]"
//line /usr/local/go/src/crypto/tls/conn.go:401
		_go_fuzz_dep_.CoverTab[21671]++

								if hc.version == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/conn.go:403
			_go_fuzz_dep_.CoverTab[21694]++
									if typ != recordTypeApplicationData {
//line /usr/local/go/src/crypto/tls/conn.go:404
				_go_fuzz_dep_.CoverTab[21697]++
										return nil, 0, alertUnexpectedMessage
//line /usr/local/go/src/crypto/tls/conn.go:405
				// _ = "end of CoverTab[21697]"
			} else {
//line /usr/local/go/src/crypto/tls/conn.go:406
				_go_fuzz_dep_.CoverTab[21698]++
//line /usr/local/go/src/crypto/tls/conn.go:406
				// _ = "end of CoverTab[21698]"
//line /usr/local/go/src/crypto/tls/conn.go:406
			}
//line /usr/local/go/src/crypto/tls/conn.go:406
			// _ = "end of CoverTab[21694]"
//line /usr/local/go/src/crypto/tls/conn.go:406
			_go_fuzz_dep_.CoverTab[21695]++
									if len(plaintext) > maxPlaintext+1 {
//line /usr/local/go/src/crypto/tls/conn.go:407
				_go_fuzz_dep_.CoverTab[21699]++
										return nil, 0, alertRecordOverflow
//line /usr/local/go/src/crypto/tls/conn.go:408
				// _ = "end of CoverTab[21699]"
			} else {
//line /usr/local/go/src/crypto/tls/conn.go:409
				_go_fuzz_dep_.CoverTab[21700]++
//line /usr/local/go/src/crypto/tls/conn.go:409
				// _ = "end of CoverTab[21700]"
//line /usr/local/go/src/crypto/tls/conn.go:409
			}
//line /usr/local/go/src/crypto/tls/conn.go:409
			// _ = "end of CoverTab[21695]"
//line /usr/local/go/src/crypto/tls/conn.go:409
			_go_fuzz_dep_.CoverTab[21696]++

									for i := len(plaintext) - 1; i >= 0; i-- {
//line /usr/local/go/src/crypto/tls/conn.go:411
				_go_fuzz_dep_.CoverTab[21701]++
										if plaintext[i] != 0 {
//line /usr/local/go/src/crypto/tls/conn.go:412
					_go_fuzz_dep_.CoverTab[21703]++
											typ = recordType(plaintext[i])
											plaintext = plaintext[:i]
											break
//line /usr/local/go/src/crypto/tls/conn.go:415
					// _ = "end of CoverTab[21703]"
				} else {
//line /usr/local/go/src/crypto/tls/conn.go:416
					_go_fuzz_dep_.CoverTab[21704]++
//line /usr/local/go/src/crypto/tls/conn.go:416
					// _ = "end of CoverTab[21704]"
//line /usr/local/go/src/crypto/tls/conn.go:416
				}
//line /usr/local/go/src/crypto/tls/conn.go:416
				// _ = "end of CoverTab[21701]"
//line /usr/local/go/src/crypto/tls/conn.go:416
				_go_fuzz_dep_.CoverTab[21702]++
										if i == 0 {
//line /usr/local/go/src/crypto/tls/conn.go:417
					_go_fuzz_dep_.CoverTab[21705]++
											return nil, 0, alertUnexpectedMessage
//line /usr/local/go/src/crypto/tls/conn.go:418
					// _ = "end of CoverTab[21705]"
				} else {
//line /usr/local/go/src/crypto/tls/conn.go:419
					_go_fuzz_dep_.CoverTab[21706]++
//line /usr/local/go/src/crypto/tls/conn.go:419
					// _ = "end of CoverTab[21706]"
//line /usr/local/go/src/crypto/tls/conn.go:419
				}
//line /usr/local/go/src/crypto/tls/conn.go:419
				// _ = "end of CoverTab[21702]"
			}
//line /usr/local/go/src/crypto/tls/conn.go:420
			// _ = "end of CoverTab[21696]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:421
			_go_fuzz_dep_.CoverTab[21707]++
//line /usr/local/go/src/crypto/tls/conn.go:421
			// _ = "end of CoverTab[21707]"
//line /usr/local/go/src/crypto/tls/conn.go:421
		}
//line /usr/local/go/src/crypto/tls/conn.go:421
		// _ = "end of CoverTab[21671]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:422
		_go_fuzz_dep_.CoverTab[21708]++
								plaintext = payload
//line /usr/local/go/src/crypto/tls/conn.go:423
		// _ = "end of CoverTab[21708]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:424
	// _ = "end of CoverTab[21664]"
//line /usr/local/go/src/crypto/tls/conn.go:424
	_go_fuzz_dep_.CoverTab[21665]++

							if hc.mac != nil {
//line /usr/local/go/src/crypto/tls/conn.go:426
		_go_fuzz_dep_.CoverTab[21709]++
								macSize := hc.mac.Size()
								if len(payload) < macSize {
//line /usr/local/go/src/crypto/tls/conn.go:428
			_go_fuzz_dep_.CoverTab[21712]++
									return nil, 0, alertBadRecordMAC
//line /usr/local/go/src/crypto/tls/conn.go:429
			// _ = "end of CoverTab[21712]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:430
			_go_fuzz_dep_.CoverTab[21713]++
//line /usr/local/go/src/crypto/tls/conn.go:430
			// _ = "end of CoverTab[21713]"
//line /usr/local/go/src/crypto/tls/conn.go:430
		}
//line /usr/local/go/src/crypto/tls/conn.go:430
		// _ = "end of CoverTab[21709]"
//line /usr/local/go/src/crypto/tls/conn.go:430
		_go_fuzz_dep_.CoverTab[21710]++

								n := len(payload) - macSize - paddingLen
								n = subtle.ConstantTimeSelect(int(uint32(n)>>31), 0, n)
								record[3] = byte(n >> 8)
								record[4] = byte(n)
								remoteMAC := payload[n : n+macSize]
								localMAC := tls10MAC(hc.mac, hc.scratchBuf[:0], hc.seq[:], record[:recordHeaderLen], payload[:n], payload[n+macSize:])

//line /usr/local/go/src/crypto/tls/conn.go:446
		macAndPaddingGood := subtle.ConstantTimeCompare(localMAC, remoteMAC) & int(paddingGood)
		if macAndPaddingGood != 1 {
//line /usr/local/go/src/crypto/tls/conn.go:447
			_go_fuzz_dep_.CoverTab[21714]++
									return nil, 0, alertBadRecordMAC
//line /usr/local/go/src/crypto/tls/conn.go:448
			// _ = "end of CoverTab[21714]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:449
			_go_fuzz_dep_.CoverTab[21715]++
//line /usr/local/go/src/crypto/tls/conn.go:449
			// _ = "end of CoverTab[21715]"
//line /usr/local/go/src/crypto/tls/conn.go:449
		}
//line /usr/local/go/src/crypto/tls/conn.go:449
		// _ = "end of CoverTab[21710]"
//line /usr/local/go/src/crypto/tls/conn.go:449
		_go_fuzz_dep_.CoverTab[21711]++

								plaintext = payload[:n]
//line /usr/local/go/src/crypto/tls/conn.go:451
		// _ = "end of CoverTab[21711]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:452
		_go_fuzz_dep_.CoverTab[21716]++
//line /usr/local/go/src/crypto/tls/conn.go:452
		// _ = "end of CoverTab[21716]"
//line /usr/local/go/src/crypto/tls/conn.go:452
	}
//line /usr/local/go/src/crypto/tls/conn.go:452
	// _ = "end of CoverTab[21665]"
//line /usr/local/go/src/crypto/tls/conn.go:452
	_go_fuzz_dep_.CoverTab[21666]++

							hc.incSeq()
							return plaintext, typ, nil
//line /usr/local/go/src/crypto/tls/conn.go:455
	// _ = "end of CoverTab[21666]"
}

// sliceForAppend extends the input slice by n bytes. head is the full extended
//line /usr/local/go/src/crypto/tls/conn.go:458
// slice, while tail is the appended part. If the original slice has sufficient
//line /usr/local/go/src/crypto/tls/conn.go:458
// capacity no allocation is performed.
//line /usr/local/go/src/crypto/tls/conn.go:461
func sliceForAppend(in []byte, n int) (head, tail []byte) {
//line /usr/local/go/src/crypto/tls/conn.go:461
	_go_fuzz_dep_.CoverTab[21717]++
							if total := len(in) + n; cap(in) >= total {
//line /usr/local/go/src/crypto/tls/conn.go:462
		_go_fuzz_dep_.CoverTab[21719]++
								head = in[:total]
//line /usr/local/go/src/crypto/tls/conn.go:463
		// _ = "end of CoverTab[21719]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:464
		_go_fuzz_dep_.CoverTab[21720]++
								head = make([]byte, total)
								copy(head, in)
//line /usr/local/go/src/crypto/tls/conn.go:466
		// _ = "end of CoverTab[21720]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:467
	// _ = "end of CoverTab[21717]"
//line /usr/local/go/src/crypto/tls/conn.go:467
	_go_fuzz_dep_.CoverTab[21718]++
							tail = head[len(in):]
							return
//line /usr/local/go/src/crypto/tls/conn.go:469
	// _ = "end of CoverTab[21718]"
}

// encrypt encrypts payload, adding the appropriate nonce and/or MAC, and
//line /usr/local/go/src/crypto/tls/conn.go:472
// appends it to record, which must already contain the record header.
//line /usr/local/go/src/crypto/tls/conn.go:474
func (hc *halfConn) encrypt(record, payload []byte, rand io.Reader) ([]byte, error) {
//line /usr/local/go/src/crypto/tls/conn.go:474
	_go_fuzz_dep_.CoverTab[21721]++
							if hc.cipher == nil {
//line /usr/local/go/src/crypto/tls/conn.go:475
		_go_fuzz_dep_.CoverTab[21725]++
								return append(record, payload...), nil
//line /usr/local/go/src/crypto/tls/conn.go:476
		// _ = "end of CoverTab[21725]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:477
		_go_fuzz_dep_.CoverTab[21726]++
//line /usr/local/go/src/crypto/tls/conn.go:477
		// _ = "end of CoverTab[21726]"
//line /usr/local/go/src/crypto/tls/conn.go:477
	}
//line /usr/local/go/src/crypto/tls/conn.go:477
	// _ = "end of CoverTab[21721]"
//line /usr/local/go/src/crypto/tls/conn.go:477
	_go_fuzz_dep_.CoverTab[21722]++

							var explicitNonce []byte
							if explicitNonceLen := hc.explicitNonceLen(); explicitNonceLen > 0 {
//line /usr/local/go/src/crypto/tls/conn.go:480
		_go_fuzz_dep_.CoverTab[21727]++
								record, explicitNonce = sliceForAppend(record, explicitNonceLen)
								if _, isCBC := hc.cipher.(cbcMode); !isCBC && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:482
			_go_fuzz_dep_.CoverTab[21728]++
//line /usr/local/go/src/crypto/tls/conn.go:482
			return explicitNonceLen < 16
//line /usr/local/go/src/crypto/tls/conn.go:482
			// _ = "end of CoverTab[21728]"
//line /usr/local/go/src/crypto/tls/conn.go:482
		}() {
//line /usr/local/go/src/crypto/tls/conn.go:482
			_go_fuzz_dep_.CoverTab[21729]++

//line /usr/local/go/src/crypto/tls/conn.go:492
			copy(explicitNonce, hc.seq[:])
//line /usr/local/go/src/crypto/tls/conn.go:492
			// _ = "end of CoverTab[21729]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:493
			_go_fuzz_dep_.CoverTab[21730]++
									if _, err := io.ReadFull(rand, explicitNonce); err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:494
				_go_fuzz_dep_.CoverTab[21731]++
										return nil, err
//line /usr/local/go/src/crypto/tls/conn.go:495
				// _ = "end of CoverTab[21731]"
			} else {
//line /usr/local/go/src/crypto/tls/conn.go:496
				_go_fuzz_dep_.CoverTab[21732]++
//line /usr/local/go/src/crypto/tls/conn.go:496
				// _ = "end of CoverTab[21732]"
//line /usr/local/go/src/crypto/tls/conn.go:496
			}
//line /usr/local/go/src/crypto/tls/conn.go:496
			// _ = "end of CoverTab[21730]"
		}
//line /usr/local/go/src/crypto/tls/conn.go:497
		// _ = "end of CoverTab[21727]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:498
		_go_fuzz_dep_.CoverTab[21733]++
//line /usr/local/go/src/crypto/tls/conn.go:498
		// _ = "end of CoverTab[21733]"
//line /usr/local/go/src/crypto/tls/conn.go:498
	}
//line /usr/local/go/src/crypto/tls/conn.go:498
	// _ = "end of CoverTab[21722]"
//line /usr/local/go/src/crypto/tls/conn.go:498
	_go_fuzz_dep_.CoverTab[21723]++

							var dst []byte
							switch c := hc.cipher.(type) {
	case cipher.Stream:
//line /usr/local/go/src/crypto/tls/conn.go:502
		_go_fuzz_dep_.CoverTab[21734]++
								mac := tls10MAC(hc.mac, hc.scratchBuf[:0], hc.seq[:], record[:recordHeaderLen], payload, nil)
								record, dst = sliceForAppend(record, len(payload)+len(mac))
								c.XORKeyStream(dst[:len(payload)], payload)
								c.XORKeyStream(dst[len(payload):], mac)
//line /usr/local/go/src/crypto/tls/conn.go:506
		// _ = "end of CoverTab[21734]"
	case aead:
//line /usr/local/go/src/crypto/tls/conn.go:507
		_go_fuzz_dep_.CoverTab[21735]++
								nonce := explicitNonce
								if len(nonce) == 0 {
//line /usr/local/go/src/crypto/tls/conn.go:509
			_go_fuzz_dep_.CoverTab[21741]++
									nonce = hc.seq[:]
//line /usr/local/go/src/crypto/tls/conn.go:510
			// _ = "end of CoverTab[21741]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:511
			_go_fuzz_dep_.CoverTab[21742]++
//line /usr/local/go/src/crypto/tls/conn.go:511
			// _ = "end of CoverTab[21742]"
//line /usr/local/go/src/crypto/tls/conn.go:511
		}
//line /usr/local/go/src/crypto/tls/conn.go:511
		// _ = "end of CoverTab[21735]"
//line /usr/local/go/src/crypto/tls/conn.go:511
		_go_fuzz_dep_.CoverTab[21736]++

								if hc.version == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/conn.go:513
			_go_fuzz_dep_.CoverTab[21743]++
									record = append(record, payload...)

//line /usr/local/go/src/crypto/tls/conn.go:517
			record = append(record, record[0])
			record[0] = byte(recordTypeApplicationData)

			n := len(payload) + 1 + c.Overhead()
			record[3] = byte(n >> 8)
			record[4] = byte(n)

			record = c.Seal(record[:recordHeaderLen],
				nonce, record[recordHeaderLen:], record[:recordHeaderLen])
//line /usr/local/go/src/crypto/tls/conn.go:525
			// _ = "end of CoverTab[21743]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:526
			_go_fuzz_dep_.CoverTab[21744]++
									additionalData := append(hc.scratchBuf[:0], hc.seq[:]...)
									additionalData = append(additionalData, record[:recordHeaderLen]...)
									record = c.Seal(record, nonce, payload, additionalData)
//line /usr/local/go/src/crypto/tls/conn.go:529
			// _ = "end of CoverTab[21744]"
		}
//line /usr/local/go/src/crypto/tls/conn.go:530
		// _ = "end of CoverTab[21736]"
	case cbcMode:
//line /usr/local/go/src/crypto/tls/conn.go:531
		_go_fuzz_dep_.CoverTab[21737]++
								mac := tls10MAC(hc.mac, hc.scratchBuf[:0], hc.seq[:], record[:recordHeaderLen], payload, nil)
								blockSize := c.BlockSize()
								plaintextLen := len(payload) + len(mac)
								paddingLen := blockSize - plaintextLen%blockSize
								record, dst = sliceForAppend(record, plaintextLen+paddingLen)
								copy(dst, payload)
								copy(dst[len(payload):], mac)
								for i := plaintextLen; i < len(dst); i++ {
//line /usr/local/go/src/crypto/tls/conn.go:539
			_go_fuzz_dep_.CoverTab[21745]++
									dst[i] = byte(paddingLen - 1)
//line /usr/local/go/src/crypto/tls/conn.go:540
			// _ = "end of CoverTab[21745]"
		}
//line /usr/local/go/src/crypto/tls/conn.go:541
		// _ = "end of CoverTab[21737]"
//line /usr/local/go/src/crypto/tls/conn.go:541
		_go_fuzz_dep_.CoverTab[21738]++
								if len(explicitNonce) > 0 {
//line /usr/local/go/src/crypto/tls/conn.go:542
			_go_fuzz_dep_.CoverTab[21746]++
									c.SetIV(explicitNonce)
//line /usr/local/go/src/crypto/tls/conn.go:543
			// _ = "end of CoverTab[21746]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:544
			_go_fuzz_dep_.CoverTab[21747]++
//line /usr/local/go/src/crypto/tls/conn.go:544
			// _ = "end of CoverTab[21747]"
//line /usr/local/go/src/crypto/tls/conn.go:544
		}
//line /usr/local/go/src/crypto/tls/conn.go:544
		// _ = "end of CoverTab[21738]"
//line /usr/local/go/src/crypto/tls/conn.go:544
		_go_fuzz_dep_.CoverTab[21739]++
								c.CryptBlocks(dst, dst)
//line /usr/local/go/src/crypto/tls/conn.go:545
		// _ = "end of CoverTab[21739]"
	default:
//line /usr/local/go/src/crypto/tls/conn.go:546
		_go_fuzz_dep_.CoverTab[21740]++
								panic("unknown cipher type")
//line /usr/local/go/src/crypto/tls/conn.go:547
		// _ = "end of CoverTab[21740]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:548
	// _ = "end of CoverTab[21723]"
//line /usr/local/go/src/crypto/tls/conn.go:548
	_go_fuzz_dep_.CoverTab[21724]++

//line /usr/local/go/src/crypto/tls/conn.go:551
	n := len(record) - recordHeaderLen
							record[3] = byte(n >> 8)
							record[4] = byte(n)
							hc.incSeq()

							return record, nil
//line /usr/local/go/src/crypto/tls/conn.go:556
	// _ = "end of CoverTab[21724]"
}

// RecordHeaderError is returned when a TLS record header is invalid.
type RecordHeaderError struct {
	// Msg contains a human readable string that describes the error.
	Msg	string
	// RecordHeader contains the five bytes of TLS record header that
	// triggered the error.
	RecordHeader	[5]byte
	// Conn provides the underlying net.Conn in the case that a client
	// sent an initial handshake that didn't look like TLS.
	// It is nil if there's already been a handshake or a TLS alert has
	// been written to the connection.
	Conn	net.Conn
}

func (e RecordHeaderError) Error() string {
//line /usr/local/go/src/crypto/tls/conn.go:573
	_go_fuzz_dep_.CoverTab[21748]++
//line /usr/local/go/src/crypto/tls/conn.go:573
	return "tls: " + e.Msg
//line /usr/local/go/src/crypto/tls/conn.go:573
	// _ = "end of CoverTab[21748]"
//line /usr/local/go/src/crypto/tls/conn.go:573
}

func (c *Conn) newRecordHeaderError(conn net.Conn, msg string) (err RecordHeaderError) {
//line /usr/local/go/src/crypto/tls/conn.go:575
	_go_fuzz_dep_.CoverTab[21749]++
							err.Msg = msg
							err.Conn = conn
							copy(err.RecordHeader[:], c.rawInput.Bytes())
							return err
//line /usr/local/go/src/crypto/tls/conn.go:579
	// _ = "end of CoverTab[21749]"
}

func (c *Conn) readRecord() error {
//line /usr/local/go/src/crypto/tls/conn.go:582
	_go_fuzz_dep_.CoverTab[21750]++
							return c.readRecordOrCCS(false)
//line /usr/local/go/src/crypto/tls/conn.go:583
	// _ = "end of CoverTab[21750]"
}

func (c *Conn) readChangeCipherSpec() error {
//line /usr/local/go/src/crypto/tls/conn.go:586
	_go_fuzz_dep_.CoverTab[21751]++
							return c.readRecordOrCCS(true)
//line /usr/local/go/src/crypto/tls/conn.go:587
	// _ = "end of CoverTab[21751]"
}

// readRecordOrCCS reads one or more TLS records from the connection and
//line /usr/local/go/src/crypto/tls/conn.go:590
// updates the record layer state. Some invariants:
//line /usr/local/go/src/crypto/tls/conn.go:590
//   - c.in must be locked
//line /usr/local/go/src/crypto/tls/conn.go:590
//   - c.input must be empty
//line /usr/local/go/src/crypto/tls/conn.go:590
//
//line /usr/local/go/src/crypto/tls/conn.go:590
// During the handshake one and only one of the following will happen:
//line /usr/local/go/src/crypto/tls/conn.go:590
//   - c.hand grows
//line /usr/local/go/src/crypto/tls/conn.go:590
//   - c.in.changeCipherSpec is called
//line /usr/local/go/src/crypto/tls/conn.go:590
//   - an error is returned
//line /usr/local/go/src/crypto/tls/conn.go:590
//
//line /usr/local/go/src/crypto/tls/conn.go:590
// After the handshake one and only one of the following will happen:
//line /usr/local/go/src/crypto/tls/conn.go:590
//   - c.hand grows
//line /usr/local/go/src/crypto/tls/conn.go:590
//   - c.input is set
//line /usr/local/go/src/crypto/tls/conn.go:590
//   - an error is returned
//line /usr/local/go/src/crypto/tls/conn.go:604
func (c *Conn) readRecordOrCCS(expectChangeCipherSpec bool) error {
//line /usr/local/go/src/crypto/tls/conn.go:604
	_go_fuzz_dep_.CoverTab[21752]++
							if c.in.err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:605
		_go_fuzz_dep_.CoverTab[21767]++
								return c.in.err
//line /usr/local/go/src/crypto/tls/conn.go:606
		// _ = "end of CoverTab[21767]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:607
		_go_fuzz_dep_.CoverTab[21768]++
//line /usr/local/go/src/crypto/tls/conn.go:607
		// _ = "end of CoverTab[21768]"
//line /usr/local/go/src/crypto/tls/conn.go:607
	}
//line /usr/local/go/src/crypto/tls/conn.go:607
	// _ = "end of CoverTab[21752]"
//line /usr/local/go/src/crypto/tls/conn.go:607
	_go_fuzz_dep_.CoverTab[21753]++
							handshakeComplete := c.isHandshakeComplete.Load()

//line /usr/local/go/src/crypto/tls/conn.go:611
	if c.input.Len() != 0 {
//line /usr/local/go/src/crypto/tls/conn.go:611
		_go_fuzz_dep_.CoverTab[21769]++
								return c.in.setErrorLocked(errors.New("tls: internal error: attempted to read record with pending application data"))
//line /usr/local/go/src/crypto/tls/conn.go:612
		// _ = "end of CoverTab[21769]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:613
		_go_fuzz_dep_.CoverTab[21770]++
//line /usr/local/go/src/crypto/tls/conn.go:613
		// _ = "end of CoverTab[21770]"
//line /usr/local/go/src/crypto/tls/conn.go:613
	}
//line /usr/local/go/src/crypto/tls/conn.go:613
	// _ = "end of CoverTab[21753]"
//line /usr/local/go/src/crypto/tls/conn.go:613
	_go_fuzz_dep_.CoverTab[21754]++
							c.input.Reset(nil)

//line /usr/local/go/src/crypto/tls/conn.go:617
	if err := c.readFromUntil(c.conn, recordHeaderLen); err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:617
		_go_fuzz_dep_.CoverTab[21771]++

//line /usr/local/go/src/crypto/tls/conn.go:621
		if err == io.ErrUnexpectedEOF && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:621
			_go_fuzz_dep_.CoverTab[21774]++
//line /usr/local/go/src/crypto/tls/conn.go:621
			return c.rawInput.Len() == 0
//line /usr/local/go/src/crypto/tls/conn.go:621
			// _ = "end of CoverTab[21774]"
//line /usr/local/go/src/crypto/tls/conn.go:621
		}() {
//line /usr/local/go/src/crypto/tls/conn.go:621
			_go_fuzz_dep_.CoverTab[21775]++
									err = io.EOF
//line /usr/local/go/src/crypto/tls/conn.go:622
			// _ = "end of CoverTab[21775]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:623
			_go_fuzz_dep_.CoverTab[21776]++
//line /usr/local/go/src/crypto/tls/conn.go:623
			// _ = "end of CoverTab[21776]"
//line /usr/local/go/src/crypto/tls/conn.go:623
		}
//line /usr/local/go/src/crypto/tls/conn.go:623
		// _ = "end of CoverTab[21771]"
//line /usr/local/go/src/crypto/tls/conn.go:623
		_go_fuzz_dep_.CoverTab[21772]++
								if e, ok := err.(net.Error); !ok || func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:624
			_go_fuzz_dep_.CoverTab[21777]++
//line /usr/local/go/src/crypto/tls/conn.go:624
			return !e.Temporary()
//line /usr/local/go/src/crypto/tls/conn.go:624
			// _ = "end of CoverTab[21777]"
//line /usr/local/go/src/crypto/tls/conn.go:624
		}() {
//line /usr/local/go/src/crypto/tls/conn.go:624
			_go_fuzz_dep_.CoverTab[21778]++
									c.in.setErrorLocked(err)
//line /usr/local/go/src/crypto/tls/conn.go:625
			// _ = "end of CoverTab[21778]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:626
			_go_fuzz_dep_.CoverTab[21779]++
//line /usr/local/go/src/crypto/tls/conn.go:626
			// _ = "end of CoverTab[21779]"
//line /usr/local/go/src/crypto/tls/conn.go:626
		}
//line /usr/local/go/src/crypto/tls/conn.go:626
		// _ = "end of CoverTab[21772]"
//line /usr/local/go/src/crypto/tls/conn.go:626
		_go_fuzz_dep_.CoverTab[21773]++
								return err
//line /usr/local/go/src/crypto/tls/conn.go:627
		// _ = "end of CoverTab[21773]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:628
		_go_fuzz_dep_.CoverTab[21780]++
//line /usr/local/go/src/crypto/tls/conn.go:628
		// _ = "end of CoverTab[21780]"
//line /usr/local/go/src/crypto/tls/conn.go:628
	}
//line /usr/local/go/src/crypto/tls/conn.go:628
	// _ = "end of CoverTab[21754]"
//line /usr/local/go/src/crypto/tls/conn.go:628
	_go_fuzz_dep_.CoverTab[21755]++
							hdr := c.rawInput.Bytes()[:recordHeaderLen]
							typ := recordType(hdr[0])

//line /usr/local/go/src/crypto/tls/conn.go:636
	if !handshakeComplete && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:636
		_go_fuzz_dep_.CoverTab[21781]++
//line /usr/local/go/src/crypto/tls/conn.go:636
		return typ == 0x80
//line /usr/local/go/src/crypto/tls/conn.go:636
		// _ = "end of CoverTab[21781]"
//line /usr/local/go/src/crypto/tls/conn.go:636
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:636
		_go_fuzz_dep_.CoverTab[21782]++
								c.sendAlert(alertProtocolVersion)
								return c.in.setErrorLocked(c.newRecordHeaderError(nil, "unsupported SSLv2 handshake received"))
//line /usr/local/go/src/crypto/tls/conn.go:638
		// _ = "end of CoverTab[21782]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:639
		_go_fuzz_dep_.CoverTab[21783]++
//line /usr/local/go/src/crypto/tls/conn.go:639
		// _ = "end of CoverTab[21783]"
//line /usr/local/go/src/crypto/tls/conn.go:639
	}
//line /usr/local/go/src/crypto/tls/conn.go:639
	// _ = "end of CoverTab[21755]"
//line /usr/local/go/src/crypto/tls/conn.go:639
	_go_fuzz_dep_.CoverTab[21756]++

							vers := uint16(hdr[1])<<8 | uint16(hdr[2])
							n := int(hdr[3])<<8 | int(hdr[4])
							if c.haveVers && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:643
		_go_fuzz_dep_.CoverTab[21784]++
//line /usr/local/go/src/crypto/tls/conn.go:643
		return c.vers != VersionTLS13
//line /usr/local/go/src/crypto/tls/conn.go:643
		// _ = "end of CoverTab[21784]"
//line /usr/local/go/src/crypto/tls/conn.go:643
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:643
		_go_fuzz_dep_.CoverTab[21785]++
//line /usr/local/go/src/crypto/tls/conn.go:643
		return vers != c.vers
//line /usr/local/go/src/crypto/tls/conn.go:643
		// _ = "end of CoverTab[21785]"
//line /usr/local/go/src/crypto/tls/conn.go:643
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:643
		_go_fuzz_dep_.CoverTab[21786]++
								c.sendAlert(alertProtocolVersion)
								msg := fmt.Sprintf("received record with version %x when expecting version %x", vers, c.vers)
								return c.in.setErrorLocked(c.newRecordHeaderError(nil, msg))
//line /usr/local/go/src/crypto/tls/conn.go:646
		// _ = "end of CoverTab[21786]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:647
		_go_fuzz_dep_.CoverTab[21787]++
//line /usr/local/go/src/crypto/tls/conn.go:647
		// _ = "end of CoverTab[21787]"
//line /usr/local/go/src/crypto/tls/conn.go:647
	}
//line /usr/local/go/src/crypto/tls/conn.go:647
	// _ = "end of CoverTab[21756]"
//line /usr/local/go/src/crypto/tls/conn.go:647
	_go_fuzz_dep_.CoverTab[21757]++
							if !c.haveVers {
//line /usr/local/go/src/crypto/tls/conn.go:648
		_go_fuzz_dep_.CoverTab[21788]++

//line /usr/local/go/src/crypto/tls/conn.go:653
		if (typ != recordTypeAlert && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:653
			_go_fuzz_dep_.CoverTab[21789]++
//line /usr/local/go/src/crypto/tls/conn.go:653
			return typ != recordTypeHandshake
//line /usr/local/go/src/crypto/tls/conn.go:653
			// _ = "end of CoverTab[21789]"
//line /usr/local/go/src/crypto/tls/conn.go:653
		}()) || func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:653
			_go_fuzz_dep_.CoverTab[21790]++
//line /usr/local/go/src/crypto/tls/conn.go:653
			return vers >= 0x1000
//line /usr/local/go/src/crypto/tls/conn.go:653
			// _ = "end of CoverTab[21790]"
//line /usr/local/go/src/crypto/tls/conn.go:653
		}() {
//line /usr/local/go/src/crypto/tls/conn.go:653
			_go_fuzz_dep_.CoverTab[21791]++
									return c.in.setErrorLocked(c.newRecordHeaderError(c.conn, "first record does not look like a TLS handshake"))
//line /usr/local/go/src/crypto/tls/conn.go:654
			// _ = "end of CoverTab[21791]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:655
			_go_fuzz_dep_.CoverTab[21792]++
//line /usr/local/go/src/crypto/tls/conn.go:655
			// _ = "end of CoverTab[21792]"
//line /usr/local/go/src/crypto/tls/conn.go:655
		}
//line /usr/local/go/src/crypto/tls/conn.go:655
		// _ = "end of CoverTab[21788]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:656
		_go_fuzz_dep_.CoverTab[21793]++
//line /usr/local/go/src/crypto/tls/conn.go:656
		// _ = "end of CoverTab[21793]"
//line /usr/local/go/src/crypto/tls/conn.go:656
	}
//line /usr/local/go/src/crypto/tls/conn.go:656
	// _ = "end of CoverTab[21757]"
//line /usr/local/go/src/crypto/tls/conn.go:656
	_go_fuzz_dep_.CoverTab[21758]++
							if c.vers == VersionTLS13 && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:657
		_go_fuzz_dep_.CoverTab[21794]++
//line /usr/local/go/src/crypto/tls/conn.go:657
		return n > maxCiphertextTLS13
//line /usr/local/go/src/crypto/tls/conn.go:657
		// _ = "end of CoverTab[21794]"
//line /usr/local/go/src/crypto/tls/conn.go:657
	}() || func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:657
		_go_fuzz_dep_.CoverTab[21795]++
//line /usr/local/go/src/crypto/tls/conn.go:657
		return n > maxCiphertext
//line /usr/local/go/src/crypto/tls/conn.go:657
		// _ = "end of CoverTab[21795]"
//line /usr/local/go/src/crypto/tls/conn.go:657
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:657
		_go_fuzz_dep_.CoverTab[21796]++
								c.sendAlert(alertRecordOverflow)
								msg := fmt.Sprintf("oversized record received with length %d", n)
								return c.in.setErrorLocked(c.newRecordHeaderError(nil, msg))
//line /usr/local/go/src/crypto/tls/conn.go:660
		// _ = "end of CoverTab[21796]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:661
		_go_fuzz_dep_.CoverTab[21797]++
//line /usr/local/go/src/crypto/tls/conn.go:661
		// _ = "end of CoverTab[21797]"
//line /usr/local/go/src/crypto/tls/conn.go:661
	}
//line /usr/local/go/src/crypto/tls/conn.go:661
	// _ = "end of CoverTab[21758]"
//line /usr/local/go/src/crypto/tls/conn.go:661
	_go_fuzz_dep_.CoverTab[21759]++
							if err := c.readFromUntil(c.conn, recordHeaderLen+n); err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:662
		_go_fuzz_dep_.CoverTab[21798]++
								if e, ok := err.(net.Error); !ok || func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:663
			_go_fuzz_dep_.CoverTab[21800]++
//line /usr/local/go/src/crypto/tls/conn.go:663
			return !e.Temporary()
//line /usr/local/go/src/crypto/tls/conn.go:663
			// _ = "end of CoverTab[21800]"
//line /usr/local/go/src/crypto/tls/conn.go:663
		}() {
//line /usr/local/go/src/crypto/tls/conn.go:663
			_go_fuzz_dep_.CoverTab[21801]++
									c.in.setErrorLocked(err)
//line /usr/local/go/src/crypto/tls/conn.go:664
			// _ = "end of CoverTab[21801]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:665
			_go_fuzz_dep_.CoverTab[21802]++
//line /usr/local/go/src/crypto/tls/conn.go:665
			// _ = "end of CoverTab[21802]"
//line /usr/local/go/src/crypto/tls/conn.go:665
		}
//line /usr/local/go/src/crypto/tls/conn.go:665
		// _ = "end of CoverTab[21798]"
//line /usr/local/go/src/crypto/tls/conn.go:665
		_go_fuzz_dep_.CoverTab[21799]++
								return err
//line /usr/local/go/src/crypto/tls/conn.go:666
		// _ = "end of CoverTab[21799]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:667
		_go_fuzz_dep_.CoverTab[21803]++
//line /usr/local/go/src/crypto/tls/conn.go:667
		// _ = "end of CoverTab[21803]"
//line /usr/local/go/src/crypto/tls/conn.go:667
	}
//line /usr/local/go/src/crypto/tls/conn.go:667
	// _ = "end of CoverTab[21759]"
//line /usr/local/go/src/crypto/tls/conn.go:667
	_go_fuzz_dep_.CoverTab[21760]++

//line /usr/local/go/src/crypto/tls/conn.go:670
	record := c.rawInput.Next(recordHeaderLen + n)
	data, typ, err := c.in.decrypt(record)
	if err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:672
		_go_fuzz_dep_.CoverTab[21804]++
								return c.in.setErrorLocked(c.sendAlert(err.(alert)))
//line /usr/local/go/src/crypto/tls/conn.go:673
		// _ = "end of CoverTab[21804]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:674
		_go_fuzz_dep_.CoverTab[21805]++
//line /usr/local/go/src/crypto/tls/conn.go:674
		// _ = "end of CoverTab[21805]"
//line /usr/local/go/src/crypto/tls/conn.go:674
	}
//line /usr/local/go/src/crypto/tls/conn.go:674
	// _ = "end of CoverTab[21760]"
//line /usr/local/go/src/crypto/tls/conn.go:674
	_go_fuzz_dep_.CoverTab[21761]++
							if len(data) > maxPlaintext {
//line /usr/local/go/src/crypto/tls/conn.go:675
		_go_fuzz_dep_.CoverTab[21806]++
								return c.in.setErrorLocked(c.sendAlert(alertRecordOverflow))
//line /usr/local/go/src/crypto/tls/conn.go:676
		// _ = "end of CoverTab[21806]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:677
		_go_fuzz_dep_.CoverTab[21807]++
//line /usr/local/go/src/crypto/tls/conn.go:677
		// _ = "end of CoverTab[21807]"
//line /usr/local/go/src/crypto/tls/conn.go:677
	}
//line /usr/local/go/src/crypto/tls/conn.go:677
	// _ = "end of CoverTab[21761]"
//line /usr/local/go/src/crypto/tls/conn.go:677
	_go_fuzz_dep_.CoverTab[21762]++

//line /usr/local/go/src/crypto/tls/conn.go:680
	if c.in.cipher == nil && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:680
		_go_fuzz_dep_.CoverTab[21808]++
//line /usr/local/go/src/crypto/tls/conn.go:680
		return typ == recordTypeApplicationData
//line /usr/local/go/src/crypto/tls/conn.go:680
		// _ = "end of CoverTab[21808]"
//line /usr/local/go/src/crypto/tls/conn.go:680
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:680
		_go_fuzz_dep_.CoverTab[21809]++
								return c.in.setErrorLocked(c.sendAlert(alertUnexpectedMessage))
//line /usr/local/go/src/crypto/tls/conn.go:681
		// _ = "end of CoverTab[21809]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:682
		_go_fuzz_dep_.CoverTab[21810]++
//line /usr/local/go/src/crypto/tls/conn.go:682
		// _ = "end of CoverTab[21810]"
//line /usr/local/go/src/crypto/tls/conn.go:682
	}
//line /usr/local/go/src/crypto/tls/conn.go:682
	// _ = "end of CoverTab[21762]"
//line /usr/local/go/src/crypto/tls/conn.go:682
	_go_fuzz_dep_.CoverTab[21763]++

							if typ != recordTypeAlert && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:684
		_go_fuzz_dep_.CoverTab[21811]++
//line /usr/local/go/src/crypto/tls/conn.go:684
		return typ != recordTypeChangeCipherSpec
//line /usr/local/go/src/crypto/tls/conn.go:684
		// _ = "end of CoverTab[21811]"
//line /usr/local/go/src/crypto/tls/conn.go:684
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:684
		_go_fuzz_dep_.CoverTab[21812]++
//line /usr/local/go/src/crypto/tls/conn.go:684
		return len(data) > 0
//line /usr/local/go/src/crypto/tls/conn.go:684
		// _ = "end of CoverTab[21812]"
//line /usr/local/go/src/crypto/tls/conn.go:684
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:684
		_go_fuzz_dep_.CoverTab[21813]++

								c.retryCount = 0
//line /usr/local/go/src/crypto/tls/conn.go:686
		// _ = "end of CoverTab[21813]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:687
		_go_fuzz_dep_.CoverTab[21814]++
//line /usr/local/go/src/crypto/tls/conn.go:687
		// _ = "end of CoverTab[21814]"
//line /usr/local/go/src/crypto/tls/conn.go:687
	}
//line /usr/local/go/src/crypto/tls/conn.go:687
	// _ = "end of CoverTab[21763]"
//line /usr/local/go/src/crypto/tls/conn.go:687
	_go_fuzz_dep_.CoverTab[21764]++

//line /usr/local/go/src/crypto/tls/conn.go:690
	if c.vers == VersionTLS13 && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:690
		_go_fuzz_dep_.CoverTab[21815]++
//line /usr/local/go/src/crypto/tls/conn.go:690
		return typ != recordTypeHandshake
//line /usr/local/go/src/crypto/tls/conn.go:690
		// _ = "end of CoverTab[21815]"
//line /usr/local/go/src/crypto/tls/conn.go:690
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:690
		_go_fuzz_dep_.CoverTab[21816]++
//line /usr/local/go/src/crypto/tls/conn.go:690
		return c.hand.Len() > 0
//line /usr/local/go/src/crypto/tls/conn.go:690
		// _ = "end of CoverTab[21816]"
//line /usr/local/go/src/crypto/tls/conn.go:690
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:690
		_go_fuzz_dep_.CoverTab[21817]++
								return c.in.setErrorLocked(c.sendAlert(alertUnexpectedMessage))
//line /usr/local/go/src/crypto/tls/conn.go:691
		// _ = "end of CoverTab[21817]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:692
		_go_fuzz_dep_.CoverTab[21818]++
//line /usr/local/go/src/crypto/tls/conn.go:692
		// _ = "end of CoverTab[21818]"
//line /usr/local/go/src/crypto/tls/conn.go:692
	}
//line /usr/local/go/src/crypto/tls/conn.go:692
	// _ = "end of CoverTab[21764]"
//line /usr/local/go/src/crypto/tls/conn.go:692
	_go_fuzz_dep_.CoverTab[21765]++

							switch typ {
	default:
//line /usr/local/go/src/crypto/tls/conn.go:695
		_go_fuzz_dep_.CoverTab[21819]++
								return c.in.setErrorLocked(c.sendAlert(alertUnexpectedMessage))
//line /usr/local/go/src/crypto/tls/conn.go:696
		// _ = "end of CoverTab[21819]"

	case recordTypeAlert:
//line /usr/local/go/src/crypto/tls/conn.go:698
		_go_fuzz_dep_.CoverTab[21820]++
								if len(data) != 2 {
//line /usr/local/go/src/crypto/tls/conn.go:699
			_go_fuzz_dep_.CoverTab[21834]++
									return c.in.setErrorLocked(c.sendAlert(alertUnexpectedMessage))
//line /usr/local/go/src/crypto/tls/conn.go:700
			// _ = "end of CoverTab[21834]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:701
			_go_fuzz_dep_.CoverTab[21835]++
//line /usr/local/go/src/crypto/tls/conn.go:701
			// _ = "end of CoverTab[21835]"
//line /usr/local/go/src/crypto/tls/conn.go:701
		}
//line /usr/local/go/src/crypto/tls/conn.go:701
		// _ = "end of CoverTab[21820]"
//line /usr/local/go/src/crypto/tls/conn.go:701
		_go_fuzz_dep_.CoverTab[21821]++
								if alert(data[1]) == alertCloseNotify {
//line /usr/local/go/src/crypto/tls/conn.go:702
			_go_fuzz_dep_.CoverTab[21836]++
									return c.in.setErrorLocked(io.EOF)
//line /usr/local/go/src/crypto/tls/conn.go:703
			// _ = "end of CoverTab[21836]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:704
			_go_fuzz_dep_.CoverTab[21837]++
//line /usr/local/go/src/crypto/tls/conn.go:704
			// _ = "end of CoverTab[21837]"
//line /usr/local/go/src/crypto/tls/conn.go:704
		}
//line /usr/local/go/src/crypto/tls/conn.go:704
		// _ = "end of CoverTab[21821]"
//line /usr/local/go/src/crypto/tls/conn.go:704
		_go_fuzz_dep_.CoverTab[21822]++
								if c.vers == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/conn.go:705
			_go_fuzz_dep_.CoverTab[21838]++
									return c.in.setErrorLocked(&net.OpError{Op: "remote error", Err: alert(data[1])})
//line /usr/local/go/src/crypto/tls/conn.go:706
			// _ = "end of CoverTab[21838]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:707
			_go_fuzz_dep_.CoverTab[21839]++
//line /usr/local/go/src/crypto/tls/conn.go:707
			// _ = "end of CoverTab[21839]"
//line /usr/local/go/src/crypto/tls/conn.go:707
		}
//line /usr/local/go/src/crypto/tls/conn.go:707
		// _ = "end of CoverTab[21822]"
//line /usr/local/go/src/crypto/tls/conn.go:707
		_go_fuzz_dep_.CoverTab[21823]++
								switch data[0] {
		case alertLevelWarning:
//line /usr/local/go/src/crypto/tls/conn.go:709
			_go_fuzz_dep_.CoverTab[21840]++

									return c.retryReadRecord(expectChangeCipherSpec)
//line /usr/local/go/src/crypto/tls/conn.go:711
			// _ = "end of CoverTab[21840]"
		case alertLevelError:
//line /usr/local/go/src/crypto/tls/conn.go:712
			_go_fuzz_dep_.CoverTab[21841]++
									return c.in.setErrorLocked(&net.OpError{Op: "remote error", Err: alert(data[1])})
//line /usr/local/go/src/crypto/tls/conn.go:713
			// _ = "end of CoverTab[21841]"
		default:
//line /usr/local/go/src/crypto/tls/conn.go:714
			_go_fuzz_dep_.CoverTab[21842]++
									return c.in.setErrorLocked(c.sendAlert(alertUnexpectedMessage))
//line /usr/local/go/src/crypto/tls/conn.go:715
			// _ = "end of CoverTab[21842]"
		}
//line /usr/local/go/src/crypto/tls/conn.go:716
		// _ = "end of CoverTab[21823]"

	case recordTypeChangeCipherSpec:
//line /usr/local/go/src/crypto/tls/conn.go:718
		_go_fuzz_dep_.CoverTab[21824]++
								if len(data) != 1 || func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:719
			_go_fuzz_dep_.CoverTab[21843]++
//line /usr/local/go/src/crypto/tls/conn.go:719
			return data[0] != 1
//line /usr/local/go/src/crypto/tls/conn.go:719
			// _ = "end of CoverTab[21843]"
//line /usr/local/go/src/crypto/tls/conn.go:719
		}() {
//line /usr/local/go/src/crypto/tls/conn.go:719
			_go_fuzz_dep_.CoverTab[21844]++
									return c.in.setErrorLocked(c.sendAlert(alertDecodeError))
//line /usr/local/go/src/crypto/tls/conn.go:720
			// _ = "end of CoverTab[21844]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:721
			_go_fuzz_dep_.CoverTab[21845]++
//line /usr/local/go/src/crypto/tls/conn.go:721
			// _ = "end of CoverTab[21845]"
//line /usr/local/go/src/crypto/tls/conn.go:721
		}
//line /usr/local/go/src/crypto/tls/conn.go:721
		// _ = "end of CoverTab[21824]"
//line /usr/local/go/src/crypto/tls/conn.go:721
		_go_fuzz_dep_.CoverTab[21825]++

								if c.hand.Len() > 0 {
//line /usr/local/go/src/crypto/tls/conn.go:723
			_go_fuzz_dep_.CoverTab[21846]++
									return c.in.setErrorLocked(c.sendAlert(alertUnexpectedMessage))
//line /usr/local/go/src/crypto/tls/conn.go:724
			// _ = "end of CoverTab[21846]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:725
			_go_fuzz_dep_.CoverTab[21847]++
//line /usr/local/go/src/crypto/tls/conn.go:725
			// _ = "end of CoverTab[21847]"
//line /usr/local/go/src/crypto/tls/conn.go:725
		}
//line /usr/local/go/src/crypto/tls/conn.go:725
		// _ = "end of CoverTab[21825]"
//line /usr/local/go/src/crypto/tls/conn.go:725
		_go_fuzz_dep_.CoverTab[21826]++

//line /usr/local/go/src/crypto/tls/conn.go:731
		if c.vers == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/conn.go:731
			_go_fuzz_dep_.CoverTab[21848]++
									return c.retryReadRecord(expectChangeCipherSpec)
//line /usr/local/go/src/crypto/tls/conn.go:732
			// _ = "end of CoverTab[21848]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:733
			_go_fuzz_dep_.CoverTab[21849]++
//line /usr/local/go/src/crypto/tls/conn.go:733
			// _ = "end of CoverTab[21849]"
//line /usr/local/go/src/crypto/tls/conn.go:733
		}
//line /usr/local/go/src/crypto/tls/conn.go:733
		// _ = "end of CoverTab[21826]"
//line /usr/local/go/src/crypto/tls/conn.go:733
		_go_fuzz_dep_.CoverTab[21827]++
								if !expectChangeCipherSpec {
//line /usr/local/go/src/crypto/tls/conn.go:734
			_go_fuzz_dep_.CoverTab[21850]++
									return c.in.setErrorLocked(c.sendAlert(alertUnexpectedMessage))
//line /usr/local/go/src/crypto/tls/conn.go:735
			// _ = "end of CoverTab[21850]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:736
			_go_fuzz_dep_.CoverTab[21851]++
//line /usr/local/go/src/crypto/tls/conn.go:736
			// _ = "end of CoverTab[21851]"
//line /usr/local/go/src/crypto/tls/conn.go:736
		}
//line /usr/local/go/src/crypto/tls/conn.go:736
		// _ = "end of CoverTab[21827]"
//line /usr/local/go/src/crypto/tls/conn.go:736
		_go_fuzz_dep_.CoverTab[21828]++
								if err := c.in.changeCipherSpec(); err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:737
			_go_fuzz_dep_.CoverTab[21852]++
									return c.in.setErrorLocked(c.sendAlert(err.(alert)))
//line /usr/local/go/src/crypto/tls/conn.go:738
			// _ = "end of CoverTab[21852]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:739
			_go_fuzz_dep_.CoverTab[21853]++
//line /usr/local/go/src/crypto/tls/conn.go:739
			// _ = "end of CoverTab[21853]"
//line /usr/local/go/src/crypto/tls/conn.go:739
		}
//line /usr/local/go/src/crypto/tls/conn.go:739
		// _ = "end of CoverTab[21828]"

	case recordTypeApplicationData:
//line /usr/local/go/src/crypto/tls/conn.go:741
		_go_fuzz_dep_.CoverTab[21829]++
								if !handshakeComplete || func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:742
			_go_fuzz_dep_.CoverTab[21854]++
//line /usr/local/go/src/crypto/tls/conn.go:742
			return expectChangeCipherSpec
//line /usr/local/go/src/crypto/tls/conn.go:742
			// _ = "end of CoverTab[21854]"
//line /usr/local/go/src/crypto/tls/conn.go:742
		}() {
//line /usr/local/go/src/crypto/tls/conn.go:742
			_go_fuzz_dep_.CoverTab[21855]++
									return c.in.setErrorLocked(c.sendAlert(alertUnexpectedMessage))
//line /usr/local/go/src/crypto/tls/conn.go:743
			// _ = "end of CoverTab[21855]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:744
			_go_fuzz_dep_.CoverTab[21856]++
//line /usr/local/go/src/crypto/tls/conn.go:744
			// _ = "end of CoverTab[21856]"
//line /usr/local/go/src/crypto/tls/conn.go:744
		}
//line /usr/local/go/src/crypto/tls/conn.go:744
		// _ = "end of CoverTab[21829]"
//line /usr/local/go/src/crypto/tls/conn.go:744
		_go_fuzz_dep_.CoverTab[21830]++

//line /usr/local/go/src/crypto/tls/conn.go:747
		if len(data) == 0 {
//line /usr/local/go/src/crypto/tls/conn.go:747
			_go_fuzz_dep_.CoverTab[21857]++
									return c.retryReadRecord(expectChangeCipherSpec)
//line /usr/local/go/src/crypto/tls/conn.go:748
			// _ = "end of CoverTab[21857]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:749
			_go_fuzz_dep_.CoverTab[21858]++
//line /usr/local/go/src/crypto/tls/conn.go:749
			// _ = "end of CoverTab[21858]"
//line /usr/local/go/src/crypto/tls/conn.go:749
		}
//line /usr/local/go/src/crypto/tls/conn.go:749
		// _ = "end of CoverTab[21830]"
//line /usr/local/go/src/crypto/tls/conn.go:749
		_go_fuzz_dep_.CoverTab[21831]++

//line /usr/local/go/src/crypto/tls/conn.go:753
		c.input.Reset(data)
//line /usr/local/go/src/crypto/tls/conn.go:753
		// _ = "end of CoverTab[21831]"

	case recordTypeHandshake:
//line /usr/local/go/src/crypto/tls/conn.go:755
		_go_fuzz_dep_.CoverTab[21832]++
								if len(data) == 0 || func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:756
			_go_fuzz_dep_.CoverTab[21859]++
//line /usr/local/go/src/crypto/tls/conn.go:756
			return expectChangeCipherSpec
//line /usr/local/go/src/crypto/tls/conn.go:756
			// _ = "end of CoverTab[21859]"
//line /usr/local/go/src/crypto/tls/conn.go:756
		}() {
//line /usr/local/go/src/crypto/tls/conn.go:756
			_go_fuzz_dep_.CoverTab[21860]++
									return c.in.setErrorLocked(c.sendAlert(alertUnexpectedMessage))
//line /usr/local/go/src/crypto/tls/conn.go:757
			// _ = "end of CoverTab[21860]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:758
			_go_fuzz_dep_.CoverTab[21861]++
//line /usr/local/go/src/crypto/tls/conn.go:758
			// _ = "end of CoverTab[21861]"
//line /usr/local/go/src/crypto/tls/conn.go:758
		}
//line /usr/local/go/src/crypto/tls/conn.go:758
		// _ = "end of CoverTab[21832]"
//line /usr/local/go/src/crypto/tls/conn.go:758
		_go_fuzz_dep_.CoverTab[21833]++
								c.hand.Write(data)
//line /usr/local/go/src/crypto/tls/conn.go:759
		// _ = "end of CoverTab[21833]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:760
	// _ = "end of CoverTab[21765]"
//line /usr/local/go/src/crypto/tls/conn.go:760
	_go_fuzz_dep_.CoverTab[21766]++

							return nil
//line /usr/local/go/src/crypto/tls/conn.go:762
	// _ = "end of CoverTab[21766]"
}

// retryReadRecord recurs into readRecordOrCCS to drop a non-advancing record, like
//line /usr/local/go/src/crypto/tls/conn.go:765
// a warning alert, empty application_data, or a change_cipher_spec in TLS 1.3.
//line /usr/local/go/src/crypto/tls/conn.go:767
func (c *Conn) retryReadRecord(expectChangeCipherSpec bool) error {
//line /usr/local/go/src/crypto/tls/conn.go:767
	_go_fuzz_dep_.CoverTab[21862]++
							c.retryCount++
							if c.retryCount > maxUselessRecords {
//line /usr/local/go/src/crypto/tls/conn.go:769
		_go_fuzz_dep_.CoverTab[21864]++
								c.sendAlert(alertUnexpectedMessage)
								return c.in.setErrorLocked(errors.New("tls: too many ignored records"))
//line /usr/local/go/src/crypto/tls/conn.go:771
		// _ = "end of CoverTab[21864]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:772
		_go_fuzz_dep_.CoverTab[21865]++
//line /usr/local/go/src/crypto/tls/conn.go:772
		// _ = "end of CoverTab[21865]"
//line /usr/local/go/src/crypto/tls/conn.go:772
	}
//line /usr/local/go/src/crypto/tls/conn.go:772
	// _ = "end of CoverTab[21862]"
//line /usr/local/go/src/crypto/tls/conn.go:772
	_go_fuzz_dep_.CoverTab[21863]++
							return c.readRecordOrCCS(expectChangeCipherSpec)
//line /usr/local/go/src/crypto/tls/conn.go:773
	// _ = "end of CoverTab[21863]"
}

// atLeastReader reads from R, stopping with EOF once at least N bytes have been
//line /usr/local/go/src/crypto/tls/conn.go:776
// read. It is different from an io.LimitedReader in that it doesn't cut short
//line /usr/local/go/src/crypto/tls/conn.go:776
// the last Read call, and in that it considers an early EOF an error.
//line /usr/local/go/src/crypto/tls/conn.go:779
type atLeastReader struct {
	R	io.Reader
	N	int64
}

func (r *atLeastReader) Read(p []byte) (int, error) {
//line /usr/local/go/src/crypto/tls/conn.go:784
	_go_fuzz_dep_.CoverTab[21866]++
							if r.N <= 0 {
//line /usr/local/go/src/crypto/tls/conn.go:785
		_go_fuzz_dep_.CoverTab[21870]++
								return 0, io.EOF
//line /usr/local/go/src/crypto/tls/conn.go:786
		// _ = "end of CoverTab[21870]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:787
		_go_fuzz_dep_.CoverTab[21871]++
//line /usr/local/go/src/crypto/tls/conn.go:787
		// _ = "end of CoverTab[21871]"
//line /usr/local/go/src/crypto/tls/conn.go:787
	}
//line /usr/local/go/src/crypto/tls/conn.go:787
	// _ = "end of CoverTab[21866]"
//line /usr/local/go/src/crypto/tls/conn.go:787
	_go_fuzz_dep_.CoverTab[21867]++
							n, err := r.R.Read(p)
							r.N -= int64(n)
							if r.N > 0 && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:790
		_go_fuzz_dep_.CoverTab[21872]++
//line /usr/local/go/src/crypto/tls/conn.go:790
		return err == io.EOF
//line /usr/local/go/src/crypto/tls/conn.go:790
		// _ = "end of CoverTab[21872]"
//line /usr/local/go/src/crypto/tls/conn.go:790
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:790
		_go_fuzz_dep_.CoverTab[21873]++
								return n, io.ErrUnexpectedEOF
//line /usr/local/go/src/crypto/tls/conn.go:791
		// _ = "end of CoverTab[21873]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:792
		_go_fuzz_dep_.CoverTab[21874]++
//line /usr/local/go/src/crypto/tls/conn.go:792
		// _ = "end of CoverTab[21874]"
//line /usr/local/go/src/crypto/tls/conn.go:792
	}
//line /usr/local/go/src/crypto/tls/conn.go:792
	// _ = "end of CoverTab[21867]"
//line /usr/local/go/src/crypto/tls/conn.go:792
	_go_fuzz_dep_.CoverTab[21868]++
							if r.N <= 0 && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:793
		_go_fuzz_dep_.CoverTab[21875]++
//line /usr/local/go/src/crypto/tls/conn.go:793
		return err == nil
//line /usr/local/go/src/crypto/tls/conn.go:793
		// _ = "end of CoverTab[21875]"
//line /usr/local/go/src/crypto/tls/conn.go:793
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:793
		_go_fuzz_dep_.CoverTab[21876]++
								return n, io.EOF
//line /usr/local/go/src/crypto/tls/conn.go:794
		// _ = "end of CoverTab[21876]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:795
		_go_fuzz_dep_.CoverTab[21877]++
//line /usr/local/go/src/crypto/tls/conn.go:795
		// _ = "end of CoverTab[21877]"
//line /usr/local/go/src/crypto/tls/conn.go:795
	}
//line /usr/local/go/src/crypto/tls/conn.go:795
	// _ = "end of CoverTab[21868]"
//line /usr/local/go/src/crypto/tls/conn.go:795
	_go_fuzz_dep_.CoverTab[21869]++
							return n, err
//line /usr/local/go/src/crypto/tls/conn.go:796
	// _ = "end of CoverTab[21869]"
}

// readFromUntil reads from r into c.rawInput until c.rawInput contains
//line /usr/local/go/src/crypto/tls/conn.go:799
// at least n bytes or else returns an error.
//line /usr/local/go/src/crypto/tls/conn.go:801
func (c *Conn) readFromUntil(r io.Reader, n int) error {
//line /usr/local/go/src/crypto/tls/conn.go:801
	_go_fuzz_dep_.CoverTab[21878]++
							if c.rawInput.Len() >= n {
//line /usr/local/go/src/crypto/tls/conn.go:802
		_go_fuzz_dep_.CoverTab[21880]++
								return nil
//line /usr/local/go/src/crypto/tls/conn.go:803
		// _ = "end of CoverTab[21880]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:804
		_go_fuzz_dep_.CoverTab[21881]++
//line /usr/local/go/src/crypto/tls/conn.go:804
		// _ = "end of CoverTab[21881]"
//line /usr/local/go/src/crypto/tls/conn.go:804
	}
//line /usr/local/go/src/crypto/tls/conn.go:804
	// _ = "end of CoverTab[21878]"
//line /usr/local/go/src/crypto/tls/conn.go:804
	_go_fuzz_dep_.CoverTab[21879]++
							needs := n - c.rawInput.Len()

//line /usr/local/go/src/crypto/tls/conn.go:809
	c.rawInput.Grow(needs + bytes.MinRead)
							_, err := c.rawInput.ReadFrom(&atLeastReader{r, int64(needs)})
							return err
//line /usr/local/go/src/crypto/tls/conn.go:811
	// _ = "end of CoverTab[21879]"
}

// sendAlert sends a TLS alert message.
func (c *Conn) sendAlertLocked(err alert) error {
//line /usr/local/go/src/crypto/tls/conn.go:815
	_go_fuzz_dep_.CoverTab[21882]++
							switch err {
	case alertNoRenegotiation, alertCloseNotify:
//line /usr/local/go/src/crypto/tls/conn.go:817
		_go_fuzz_dep_.CoverTab[21885]++
								c.tmp[0] = alertLevelWarning
//line /usr/local/go/src/crypto/tls/conn.go:818
		// _ = "end of CoverTab[21885]"
	default:
//line /usr/local/go/src/crypto/tls/conn.go:819
		_go_fuzz_dep_.CoverTab[21886]++
								c.tmp[0] = alertLevelError
//line /usr/local/go/src/crypto/tls/conn.go:820
		// _ = "end of CoverTab[21886]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:821
	// _ = "end of CoverTab[21882]"
//line /usr/local/go/src/crypto/tls/conn.go:821
	_go_fuzz_dep_.CoverTab[21883]++
							c.tmp[1] = byte(err)

							_, writeErr := c.writeRecordLocked(recordTypeAlert, c.tmp[0:2])
							if err == alertCloseNotify {
//line /usr/local/go/src/crypto/tls/conn.go:825
		_go_fuzz_dep_.CoverTab[21887]++

								return writeErr
//line /usr/local/go/src/crypto/tls/conn.go:827
		// _ = "end of CoverTab[21887]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:828
		_go_fuzz_dep_.CoverTab[21888]++
//line /usr/local/go/src/crypto/tls/conn.go:828
		// _ = "end of CoverTab[21888]"
//line /usr/local/go/src/crypto/tls/conn.go:828
	}
//line /usr/local/go/src/crypto/tls/conn.go:828
	// _ = "end of CoverTab[21883]"
//line /usr/local/go/src/crypto/tls/conn.go:828
	_go_fuzz_dep_.CoverTab[21884]++

							return c.out.setErrorLocked(&net.OpError{Op: "local error", Err: err})
//line /usr/local/go/src/crypto/tls/conn.go:830
	// _ = "end of CoverTab[21884]"
}

// sendAlert sends a TLS alert message.
func (c *Conn) sendAlert(err alert) error {
//line /usr/local/go/src/crypto/tls/conn.go:834
	_go_fuzz_dep_.CoverTab[21889]++
							c.out.Lock()
							defer c.out.Unlock()
							return c.sendAlertLocked(err)
//line /usr/local/go/src/crypto/tls/conn.go:837
	// _ = "end of CoverTab[21889]"
}

const (
	// tcpMSSEstimate is a conservative estimate of the TCP maximum segment
	// size (MSS). A constant is used, rather than querying the kernel for
	// the actual MSS, to avoid complexity. The value here is the IPv6
	// minimum MTU (1280 bytes) minus the overhead of an IPv6 header (40
	// bytes) and a TCP header with timestamps (32 bytes).
	tcpMSSEstimate	= 1208

	// recordSizeBoostThreshold is the number of bytes of application data
	// sent after which the TLS record size will be increased to the
	// maximum.
	recordSizeBoostThreshold	= 128 * 1024
)

// maxPayloadSizeForWrite returns the maximum TLS payload size to use for the
//line /usr/local/go/src/crypto/tls/conn.go:854
// next application data record. There is the following trade-off:
//line /usr/local/go/src/crypto/tls/conn.go:854
//
//line /usr/local/go/src/crypto/tls/conn.go:854
//   - For latency-sensitive applications, such as web browsing, each TLS
//line /usr/local/go/src/crypto/tls/conn.go:854
//     record should fit in one TCP segment.
//line /usr/local/go/src/crypto/tls/conn.go:854
//   - For throughput-sensitive applications, such as large file transfers,
//line /usr/local/go/src/crypto/tls/conn.go:854
//     larger TLS records better amortize framing and encryption overheads.
//line /usr/local/go/src/crypto/tls/conn.go:854
//
//line /usr/local/go/src/crypto/tls/conn.go:854
// A simple heuristic that works well in practice is to use small records for
//line /usr/local/go/src/crypto/tls/conn.go:854
// the first 1MB of data, then use larger records for subsequent data, and
//line /usr/local/go/src/crypto/tls/conn.go:854
// reset back to smaller records after the connection becomes idle. See "High
//line /usr/local/go/src/crypto/tls/conn.go:854
// Performance Web Networking", Chapter 4, or:
//line /usr/local/go/src/crypto/tls/conn.go:854
// https://www.igvita.com/2013/10/24/optimizing-tls-record-size-and-buffering-latency/
//line /usr/local/go/src/crypto/tls/conn.go:854
//
//line /usr/local/go/src/crypto/tls/conn.go:854
// In the interests of simplicity and determinism, this code does not attempt
//line /usr/local/go/src/crypto/tls/conn.go:854
// to reset the record size once the connection is idle, however.
//line /usr/local/go/src/crypto/tls/conn.go:870
func (c *Conn) maxPayloadSizeForWrite(typ recordType) int {
//line /usr/local/go/src/crypto/tls/conn.go:870
	_go_fuzz_dep_.CoverTab[21890]++
							if c.config.DynamicRecordSizingDisabled || func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:871
		_go_fuzz_dep_.CoverTab[21897]++
//line /usr/local/go/src/crypto/tls/conn.go:871
		return typ != recordTypeApplicationData
//line /usr/local/go/src/crypto/tls/conn.go:871
		// _ = "end of CoverTab[21897]"
//line /usr/local/go/src/crypto/tls/conn.go:871
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:871
		_go_fuzz_dep_.CoverTab[21898]++
								return maxPlaintext
//line /usr/local/go/src/crypto/tls/conn.go:872
		// _ = "end of CoverTab[21898]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:873
		_go_fuzz_dep_.CoverTab[21899]++
//line /usr/local/go/src/crypto/tls/conn.go:873
		// _ = "end of CoverTab[21899]"
//line /usr/local/go/src/crypto/tls/conn.go:873
	}
//line /usr/local/go/src/crypto/tls/conn.go:873
	// _ = "end of CoverTab[21890]"
//line /usr/local/go/src/crypto/tls/conn.go:873
	_go_fuzz_dep_.CoverTab[21891]++

							if c.bytesSent >= recordSizeBoostThreshold {
//line /usr/local/go/src/crypto/tls/conn.go:875
		_go_fuzz_dep_.CoverTab[21900]++
								return maxPlaintext
//line /usr/local/go/src/crypto/tls/conn.go:876
		// _ = "end of CoverTab[21900]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:877
		_go_fuzz_dep_.CoverTab[21901]++
//line /usr/local/go/src/crypto/tls/conn.go:877
		// _ = "end of CoverTab[21901]"
//line /usr/local/go/src/crypto/tls/conn.go:877
	}
//line /usr/local/go/src/crypto/tls/conn.go:877
	// _ = "end of CoverTab[21891]"
//line /usr/local/go/src/crypto/tls/conn.go:877
	_go_fuzz_dep_.CoverTab[21892]++

//line /usr/local/go/src/crypto/tls/conn.go:880
	payloadBytes := tcpMSSEstimate - recordHeaderLen - c.out.explicitNonceLen()
	if c.out.cipher != nil {
//line /usr/local/go/src/crypto/tls/conn.go:881
		_go_fuzz_dep_.CoverTab[21902]++
								switch ciph := c.out.cipher.(type) {
		case cipher.Stream:
//line /usr/local/go/src/crypto/tls/conn.go:883
			_go_fuzz_dep_.CoverTab[21903]++
									payloadBytes -= c.out.mac.Size()
//line /usr/local/go/src/crypto/tls/conn.go:884
			// _ = "end of CoverTab[21903]"
		case cipher.AEAD:
//line /usr/local/go/src/crypto/tls/conn.go:885
			_go_fuzz_dep_.CoverTab[21904]++
									payloadBytes -= ciph.Overhead()
//line /usr/local/go/src/crypto/tls/conn.go:886
			// _ = "end of CoverTab[21904]"
		case cbcMode:
//line /usr/local/go/src/crypto/tls/conn.go:887
			_go_fuzz_dep_.CoverTab[21905]++
									blockSize := ciph.BlockSize()

//line /usr/local/go/src/crypto/tls/conn.go:891
			payloadBytes = (payloadBytes & ^(blockSize - 1)) - 1

//line /usr/local/go/src/crypto/tls/conn.go:894
			payloadBytes -= c.out.mac.Size()
//line /usr/local/go/src/crypto/tls/conn.go:894
			// _ = "end of CoverTab[21905]"
		default:
//line /usr/local/go/src/crypto/tls/conn.go:895
			_go_fuzz_dep_.CoverTab[21906]++
									panic("unknown cipher type")
//line /usr/local/go/src/crypto/tls/conn.go:896
			// _ = "end of CoverTab[21906]"
		}
//line /usr/local/go/src/crypto/tls/conn.go:897
		// _ = "end of CoverTab[21902]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:898
		_go_fuzz_dep_.CoverTab[21907]++
//line /usr/local/go/src/crypto/tls/conn.go:898
		// _ = "end of CoverTab[21907]"
//line /usr/local/go/src/crypto/tls/conn.go:898
	}
//line /usr/local/go/src/crypto/tls/conn.go:898
	// _ = "end of CoverTab[21892]"
//line /usr/local/go/src/crypto/tls/conn.go:898
	_go_fuzz_dep_.CoverTab[21893]++
							if c.vers == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/conn.go:899
		_go_fuzz_dep_.CoverTab[21908]++
								payloadBytes--
//line /usr/local/go/src/crypto/tls/conn.go:900
		// _ = "end of CoverTab[21908]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:901
		_go_fuzz_dep_.CoverTab[21909]++
//line /usr/local/go/src/crypto/tls/conn.go:901
		// _ = "end of CoverTab[21909]"
//line /usr/local/go/src/crypto/tls/conn.go:901
	}
//line /usr/local/go/src/crypto/tls/conn.go:901
	// _ = "end of CoverTab[21893]"
//line /usr/local/go/src/crypto/tls/conn.go:901
	_go_fuzz_dep_.CoverTab[21894]++

//line /usr/local/go/src/crypto/tls/conn.go:904
	pkt := c.packetsSent
	c.packetsSent++
	if pkt > 1000 {
//line /usr/local/go/src/crypto/tls/conn.go:906
		_go_fuzz_dep_.CoverTab[21910]++
								return maxPlaintext
//line /usr/local/go/src/crypto/tls/conn.go:907
		// _ = "end of CoverTab[21910]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:908
		_go_fuzz_dep_.CoverTab[21911]++
//line /usr/local/go/src/crypto/tls/conn.go:908
		// _ = "end of CoverTab[21911]"
//line /usr/local/go/src/crypto/tls/conn.go:908
	}
//line /usr/local/go/src/crypto/tls/conn.go:908
	// _ = "end of CoverTab[21894]"
//line /usr/local/go/src/crypto/tls/conn.go:908
	_go_fuzz_dep_.CoverTab[21895]++

							n := payloadBytes * int(pkt+1)
							if n > maxPlaintext {
//line /usr/local/go/src/crypto/tls/conn.go:911
		_go_fuzz_dep_.CoverTab[21912]++
								n = maxPlaintext
//line /usr/local/go/src/crypto/tls/conn.go:912
		// _ = "end of CoverTab[21912]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:913
		_go_fuzz_dep_.CoverTab[21913]++
//line /usr/local/go/src/crypto/tls/conn.go:913
		// _ = "end of CoverTab[21913]"
//line /usr/local/go/src/crypto/tls/conn.go:913
	}
//line /usr/local/go/src/crypto/tls/conn.go:913
	// _ = "end of CoverTab[21895]"
//line /usr/local/go/src/crypto/tls/conn.go:913
	_go_fuzz_dep_.CoverTab[21896]++
							return n
//line /usr/local/go/src/crypto/tls/conn.go:914
	// _ = "end of CoverTab[21896]"
}

func (c *Conn) write(data []byte) (int, error) {
//line /usr/local/go/src/crypto/tls/conn.go:917
	_go_fuzz_dep_.CoverTab[21914]++
							if c.buffering {
//line /usr/local/go/src/crypto/tls/conn.go:918
		_go_fuzz_dep_.CoverTab[21916]++
								c.sendBuf = append(c.sendBuf, data...)
								return len(data), nil
//line /usr/local/go/src/crypto/tls/conn.go:920
		// _ = "end of CoverTab[21916]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:921
		_go_fuzz_dep_.CoverTab[21917]++
//line /usr/local/go/src/crypto/tls/conn.go:921
		// _ = "end of CoverTab[21917]"
//line /usr/local/go/src/crypto/tls/conn.go:921
	}
//line /usr/local/go/src/crypto/tls/conn.go:921
	// _ = "end of CoverTab[21914]"
//line /usr/local/go/src/crypto/tls/conn.go:921
	_go_fuzz_dep_.CoverTab[21915]++

							n, err := c.conn.Write(data)
							c.bytesSent += int64(n)
							return n, err
//line /usr/local/go/src/crypto/tls/conn.go:925
	// _ = "end of CoverTab[21915]"
}

func (c *Conn) flush() (int, error) {
//line /usr/local/go/src/crypto/tls/conn.go:928
	_go_fuzz_dep_.CoverTab[21918]++
							if len(c.sendBuf) == 0 {
//line /usr/local/go/src/crypto/tls/conn.go:929
		_go_fuzz_dep_.CoverTab[21920]++
								return 0, nil
//line /usr/local/go/src/crypto/tls/conn.go:930
		// _ = "end of CoverTab[21920]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:931
		_go_fuzz_dep_.CoverTab[21921]++
//line /usr/local/go/src/crypto/tls/conn.go:931
		// _ = "end of CoverTab[21921]"
//line /usr/local/go/src/crypto/tls/conn.go:931
	}
//line /usr/local/go/src/crypto/tls/conn.go:931
	// _ = "end of CoverTab[21918]"
//line /usr/local/go/src/crypto/tls/conn.go:931
	_go_fuzz_dep_.CoverTab[21919]++

							n, err := c.conn.Write(c.sendBuf)
							c.bytesSent += int64(n)
							c.sendBuf = nil
							c.buffering = false
							return n, err
//line /usr/local/go/src/crypto/tls/conn.go:937
	// _ = "end of CoverTab[21919]"
}

// outBufPool pools the record-sized scratch buffers used by writeRecordLocked.
var outBufPool = sync.Pool{
	New: func() any {
//line /usr/local/go/src/crypto/tls/conn.go:942
		_go_fuzz_dep_.CoverTab[21922]++
								return new([]byte)
//line /usr/local/go/src/crypto/tls/conn.go:943
		// _ = "end of CoverTab[21922]"
	},
}

// writeRecordLocked writes a TLS record with the given type and payload to the
//line /usr/local/go/src/crypto/tls/conn.go:947
// connection and updates the record layer state.
//line /usr/local/go/src/crypto/tls/conn.go:949
func (c *Conn) writeRecordLocked(typ recordType, data []byte) (int, error) {
//line /usr/local/go/src/crypto/tls/conn.go:949
	_go_fuzz_dep_.CoverTab[21923]++
							outBufPtr := outBufPool.Get().(*[]byte)
							outBuf := *outBufPtr
							defer func() {
//line /usr/local/go/src/crypto/tls/conn.go:952
		_go_fuzz_dep_.CoverTab[21927]++

//line /usr/local/go/src/crypto/tls/conn.go:958
		*outBufPtr = outBuf
								outBufPool.Put(outBufPtr)
//line /usr/local/go/src/crypto/tls/conn.go:959
		// _ = "end of CoverTab[21927]"
	}()
//line /usr/local/go/src/crypto/tls/conn.go:960
	// _ = "end of CoverTab[21923]"
//line /usr/local/go/src/crypto/tls/conn.go:960
	_go_fuzz_dep_.CoverTab[21924]++

							var n int
							for len(data) > 0 {
//line /usr/local/go/src/crypto/tls/conn.go:963
		_go_fuzz_dep_.CoverTab[21928]++
								m := len(data)
								if maxPayload := c.maxPayloadSizeForWrite(typ); m > maxPayload {
//line /usr/local/go/src/crypto/tls/conn.go:965
			_go_fuzz_dep_.CoverTab[21933]++
									m = maxPayload
//line /usr/local/go/src/crypto/tls/conn.go:966
			// _ = "end of CoverTab[21933]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:967
			_go_fuzz_dep_.CoverTab[21934]++
//line /usr/local/go/src/crypto/tls/conn.go:967
			// _ = "end of CoverTab[21934]"
//line /usr/local/go/src/crypto/tls/conn.go:967
		}
//line /usr/local/go/src/crypto/tls/conn.go:967
		// _ = "end of CoverTab[21928]"
//line /usr/local/go/src/crypto/tls/conn.go:967
		_go_fuzz_dep_.CoverTab[21929]++

								_, outBuf = sliceForAppend(outBuf[:0], recordHeaderLen)
								outBuf[0] = byte(typ)
								vers := c.vers
								if vers == 0 {
//line /usr/local/go/src/crypto/tls/conn.go:972
			_go_fuzz_dep_.CoverTab[21935]++

//line /usr/local/go/src/crypto/tls/conn.go:975
			vers = VersionTLS10
//line /usr/local/go/src/crypto/tls/conn.go:975
			// _ = "end of CoverTab[21935]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:976
			_go_fuzz_dep_.CoverTab[21936]++
//line /usr/local/go/src/crypto/tls/conn.go:976
			if vers == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/conn.go:976
				_go_fuzz_dep_.CoverTab[21937]++

//line /usr/local/go/src/crypto/tls/conn.go:979
				vers = VersionTLS12
//line /usr/local/go/src/crypto/tls/conn.go:979
				// _ = "end of CoverTab[21937]"
			} else {
//line /usr/local/go/src/crypto/tls/conn.go:980
				_go_fuzz_dep_.CoverTab[21938]++
//line /usr/local/go/src/crypto/tls/conn.go:980
				// _ = "end of CoverTab[21938]"
//line /usr/local/go/src/crypto/tls/conn.go:980
			}
//line /usr/local/go/src/crypto/tls/conn.go:980
			// _ = "end of CoverTab[21936]"
//line /usr/local/go/src/crypto/tls/conn.go:980
		}
//line /usr/local/go/src/crypto/tls/conn.go:980
		// _ = "end of CoverTab[21929]"
//line /usr/local/go/src/crypto/tls/conn.go:980
		_go_fuzz_dep_.CoverTab[21930]++
								outBuf[1] = byte(vers >> 8)
								outBuf[2] = byte(vers)
								outBuf[3] = byte(m >> 8)
								outBuf[4] = byte(m)

								var err error
								outBuf, err = c.out.encrypt(outBuf, data[:m], c.config.rand())
								if err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:988
			_go_fuzz_dep_.CoverTab[21939]++
									return n, err
//line /usr/local/go/src/crypto/tls/conn.go:989
			// _ = "end of CoverTab[21939]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:990
			_go_fuzz_dep_.CoverTab[21940]++
//line /usr/local/go/src/crypto/tls/conn.go:990
			// _ = "end of CoverTab[21940]"
//line /usr/local/go/src/crypto/tls/conn.go:990
		}
//line /usr/local/go/src/crypto/tls/conn.go:990
		// _ = "end of CoverTab[21930]"
//line /usr/local/go/src/crypto/tls/conn.go:990
		_go_fuzz_dep_.CoverTab[21931]++
								if _, err := c.write(outBuf); err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:991
			_go_fuzz_dep_.CoverTab[21941]++
									return n, err
//line /usr/local/go/src/crypto/tls/conn.go:992
			// _ = "end of CoverTab[21941]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:993
			_go_fuzz_dep_.CoverTab[21942]++
//line /usr/local/go/src/crypto/tls/conn.go:993
			// _ = "end of CoverTab[21942]"
//line /usr/local/go/src/crypto/tls/conn.go:993
		}
//line /usr/local/go/src/crypto/tls/conn.go:993
		// _ = "end of CoverTab[21931]"
//line /usr/local/go/src/crypto/tls/conn.go:993
		_go_fuzz_dep_.CoverTab[21932]++
								n += m
								data = data[m:]
//line /usr/local/go/src/crypto/tls/conn.go:995
		// _ = "end of CoverTab[21932]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:996
	// _ = "end of CoverTab[21924]"
//line /usr/local/go/src/crypto/tls/conn.go:996
	_go_fuzz_dep_.CoverTab[21925]++

							if typ == recordTypeChangeCipherSpec && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:998
		_go_fuzz_dep_.CoverTab[21943]++
//line /usr/local/go/src/crypto/tls/conn.go:998
		return c.vers != VersionTLS13
//line /usr/local/go/src/crypto/tls/conn.go:998
		// _ = "end of CoverTab[21943]"
//line /usr/local/go/src/crypto/tls/conn.go:998
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:998
		_go_fuzz_dep_.CoverTab[21944]++
								if err := c.out.changeCipherSpec(); err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:999
			_go_fuzz_dep_.CoverTab[21945]++
									return n, c.sendAlertLocked(err.(alert))
//line /usr/local/go/src/crypto/tls/conn.go:1000
			// _ = "end of CoverTab[21945]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1001
			_go_fuzz_dep_.CoverTab[21946]++
//line /usr/local/go/src/crypto/tls/conn.go:1001
			// _ = "end of CoverTab[21946]"
//line /usr/local/go/src/crypto/tls/conn.go:1001
		}
//line /usr/local/go/src/crypto/tls/conn.go:1001
		// _ = "end of CoverTab[21944]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1002
		_go_fuzz_dep_.CoverTab[21947]++
//line /usr/local/go/src/crypto/tls/conn.go:1002
		// _ = "end of CoverTab[21947]"
//line /usr/local/go/src/crypto/tls/conn.go:1002
	}
//line /usr/local/go/src/crypto/tls/conn.go:1002
	// _ = "end of CoverTab[21925]"
//line /usr/local/go/src/crypto/tls/conn.go:1002
	_go_fuzz_dep_.CoverTab[21926]++

							return n, nil
//line /usr/local/go/src/crypto/tls/conn.go:1004
	// _ = "end of CoverTab[21926]"
}

// writeHandshakeRecord writes a handshake message to the connection and updates
//line /usr/local/go/src/crypto/tls/conn.go:1007
// the record layer state. If transcript is non-nil the marshalled message is
//line /usr/local/go/src/crypto/tls/conn.go:1007
// written to it.
//line /usr/local/go/src/crypto/tls/conn.go:1010
func (c *Conn) writeHandshakeRecord(msg handshakeMessage, transcript transcriptHash) (int, error) {
//line /usr/local/go/src/crypto/tls/conn.go:1010
	_go_fuzz_dep_.CoverTab[21948]++
							c.out.Lock()
							defer c.out.Unlock()

							data, err := msg.marshal()
							if err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1015
		_go_fuzz_dep_.CoverTab[21951]++
								return 0, err
//line /usr/local/go/src/crypto/tls/conn.go:1016
		// _ = "end of CoverTab[21951]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1017
		_go_fuzz_dep_.CoverTab[21952]++
//line /usr/local/go/src/crypto/tls/conn.go:1017
		// _ = "end of CoverTab[21952]"
//line /usr/local/go/src/crypto/tls/conn.go:1017
	}
//line /usr/local/go/src/crypto/tls/conn.go:1017
	// _ = "end of CoverTab[21948]"
//line /usr/local/go/src/crypto/tls/conn.go:1017
	_go_fuzz_dep_.CoverTab[21949]++
							if transcript != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1018
		_go_fuzz_dep_.CoverTab[21953]++
								transcript.Write(data)
//line /usr/local/go/src/crypto/tls/conn.go:1019
		// _ = "end of CoverTab[21953]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1020
		_go_fuzz_dep_.CoverTab[21954]++
//line /usr/local/go/src/crypto/tls/conn.go:1020
		// _ = "end of CoverTab[21954]"
//line /usr/local/go/src/crypto/tls/conn.go:1020
	}
//line /usr/local/go/src/crypto/tls/conn.go:1020
	// _ = "end of CoverTab[21949]"
//line /usr/local/go/src/crypto/tls/conn.go:1020
	_go_fuzz_dep_.CoverTab[21950]++

							return c.writeRecordLocked(recordTypeHandshake, data)
//line /usr/local/go/src/crypto/tls/conn.go:1022
	// _ = "end of CoverTab[21950]"
}

// writeChangeCipherRecord writes a ChangeCipherSpec message to the connection and
//line /usr/local/go/src/crypto/tls/conn.go:1025
// updates the record layer state.
//line /usr/local/go/src/crypto/tls/conn.go:1027
func (c *Conn) writeChangeCipherRecord() error {
//line /usr/local/go/src/crypto/tls/conn.go:1027
	_go_fuzz_dep_.CoverTab[21955]++
							c.out.Lock()
							defer c.out.Unlock()
							_, err := c.writeRecordLocked(recordTypeChangeCipherSpec, []byte{1})
							return err
//line /usr/local/go/src/crypto/tls/conn.go:1031
	// _ = "end of CoverTab[21955]"
}

// readHandshake reads the next handshake message from
//line /usr/local/go/src/crypto/tls/conn.go:1034
// the record layer. If transcript is non-nil, the message
//line /usr/local/go/src/crypto/tls/conn.go:1034
// is written to the passed transcriptHash.
//line /usr/local/go/src/crypto/tls/conn.go:1037
func (c *Conn) readHandshake(transcript transcriptHash) (any, error) {
//line /usr/local/go/src/crypto/tls/conn.go:1037
	_go_fuzz_dep_.CoverTab[21956]++
							for c.hand.Len() < 4 {
//line /usr/local/go/src/crypto/tls/conn.go:1038
		_go_fuzz_dep_.CoverTab[21963]++
								if err := c.readRecord(); err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1039
			_go_fuzz_dep_.CoverTab[21964]++
									return nil, err
//line /usr/local/go/src/crypto/tls/conn.go:1040
			// _ = "end of CoverTab[21964]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1041
			_go_fuzz_dep_.CoverTab[21965]++
//line /usr/local/go/src/crypto/tls/conn.go:1041
			// _ = "end of CoverTab[21965]"
//line /usr/local/go/src/crypto/tls/conn.go:1041
		}
//line /usr/local/go/src/crypto/tls/conn.go:1041
		// _ = "end of CoverTab[21963]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:1042
	// _ = "end of CoverTab[21956]"
//line /usr/local/go/src/crypto/tls/conn.go:1042
	_go_fuzz_dep_.CoverTab[21957]++

							data := c.hand.Bytes()
							n := int(data[1])<<16 | int(data[2])<<8 | int(data[3])
							if n > maxHandshake {
//line /usr/local/go/src/crypto/tls/conn.go:1046
		_go_fuzz_dep_.CoverTab[21966]++
								c.sendAlertLocked(alertInternalError)
								return nil, c.in.setErrorLocked(fmt.Errorf("tls: handshake message of length %d bytes exceeds maximum of %d bytes", n, maxHandshake))
//line /usr/local/go/src/crypto/tls/conn.go:1048
		// _ = "end of CoverTab[21966]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1049
		_go_fuzz_dep_.CoverTab[21967]++
//line /usr/local/go/src/crypto/tls/conn.go:1049
		// _ = "end of CoverTab[21967]"
//line /usr/local/go/src/crypto/tls/conn.go:1049
	}
//line /usr/local/go/src/crypto/tls/conn.go:1049
	// _ = "end of CoverTab[21957]"
//line /usr/local/go/src/crypto/tls/conn.go:1049
	_go_fuzz_dep_.CoverTab[21958]++
							for c.hand.Len() < 4+n {
//line /usr/local/go/src/crypto/tls/conn.go:1050
		_go_fuzz_dep_.CoverTab[21968]++
								if err := c.readRecord(); err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1051
			_go_fuzz_dep_.CoverTab[21969]++
									return nil, err
//line /usr/local/go/src/crypto/tls/conn.go:1052
			// _ = "end of CoverTab[21969]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1053
			_go_fuzz_dep_.CoverTab[21970]++
//line /usr/local/go/src/crypto/tls/conn.go:1053
			// _ = "end of CoverTab[21970]"
//line /usr/local/go/src/crypto/tls/conn.go:1053
		}
//line /usr/local/go/src/crypto/tls/conn.go:1053
		// _ = "end of CoverTab[21968]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:1054
	// _ = "end of CoverTab[21958]"
//line /usr/local/go/src/crypto/tls/conn.go:1054
	_go_fuzz_dep_.CoverTab[21959]++
							data = c.hand.Next(4 + n)
							var m handshakeMessage
							switch data[0] {
	case typeHelloRequest:
//line /usr/local/go/src/crypto/tls/conn.go:1058
		_go_fuzz_dep_.CoverTab[21971]++
								m = new(helloRequestMsg)
//line /usr/local/go/src/crypto/tls/conn.go:1059
		// _ = "end of CoverTab[21971]"
	case typeClientHello:
//line /usr/local/go/src/crypto/tls/conn.go:1060
		_go_fuzz_dep_.CoverTab[21972]++
								m = new(clientHelloMsg)
//line /usr/local/go/src/crypto/tls/conn.go:1061
		// _ = "end of CoverTab[21972]"
	case typeServerHello:
//line /usr/local/go/src/crypto/tls/conn.go:1062
		_go_fuzz_dep_.CoverTab[21973]++
								m = new(serverHelloMsg)
//line /usr/local/go/src/crypto/tls/conn.go:1063
		// _ = "end of CoverTab[21973]"
	case typeNewSessionTicket:
//line /usr/local/go/src/crypto/tls/conn.go:1064
		_go_fuzz_dep_.CoverTab[21974]++
								if c.vers == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/conn.go:1065
			_go_fuzz_dep_.CoverTab[21987]++
									m = new(newSessionTicketMsgTLS13)
//line /usr/local/go/src/crypto/tls/conn.go:1066
			// _ = "end of CoverTab[21987]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1067
			_go_fuzz_dep_.CoverTab[21988]++
									m = new(newSessionTicketMsg)
//line /usr/local/go/src/crypto/tls/conn.go:1068
			// _ = "end of CoverTab[21988]"
		}
//line /usr/local/go/src/crypto/tls/conn.go:1069
		// _ = "end of CoverTab[21974]"
	case typeCertificate:
//line /usr/local/go/src/crypto/tls/conn.go:1070
		_go_fuzz_dep_.CoverTab[21975]++
								if c.vers == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/conn.go:1071
			_go_fuzz_dep_.CoverTab[21989]++
									m = new(certificateMsgTLS13)
//line /usr/local/go/src/crypto/tls/conn.go:1072
			// _ = "end of CoverTab[21989]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1073
			_go_fuzz_dep_.CoverTab[21990]++
									m = new(certificateMsg)
//line /usr/local/go/src/crypto/tls/conn.go:1074
			// _ = "end of CoverTab[21990]"
		}
//line /usr/local/go/src/crypto/tls/conn.go:1075
		// _ = "end of CoverTab[21975]"
	case typeCertificateRequest:
//line /usr/local/go/src/crypto/tls/conn.go:1076
		_go_fuzz_dep_.CoverTab[21976]++
								if c.vers == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/conn.go:1077
			_go_fuzz_dep_.CoverTab[21991]++
									m = new(certificateRequestMsgTLS13)
//line /usr/local/go/src/crypto/tls/conn.go:1078
			// _ = "end of CoverTab[21991]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1079
			_go_fuzz_dep_.CoverTab[21992]++
									m = &certificateRequestMsg{
				hasSignatureAlgorithm: c.vers >= VersionTLS12,
			}
//line /usr/local/go/src/crypto/tls/conn.go:1082
			// _ = "end of CoverTab[21992]"
		}
//line /usr/local/go/src/crypto/tls/conn.go:1083
		// _ = "end of CoverTab[21976]"
	case typeCertificateStatus:
//line /usr/local/go/src/crypto/tls/conn.go:1084
		_go_fuzz_dep_.CoverTab[21977]++
								m = new(certificateStatusMsg)
//line /usr/local/go/src/crypto/tls/conn.go:1085
		// _ = "end of CoverTab[21977]"
	case typeServerKeyExchange:
//line /usr/local/go/src/crypto/tls/conn.go:1086
		_go_fuzz_dep_.CoverTab[21978]++
								m = new(serverKeyExchangeMsg)
//line /usr/local/go/src/crypto/tls/conn.go:1087
		// _ = "end of CoverTab[21978]"
	case typeServerHelloDone:
//line /usr/local/go/src/crypto/tls/conn.go:1088
		_go_fuzz_dep_.CoverTab[21979]++
								m = new(serverHelloDoneMsg)
//line /usr/local/go/src/crypto/tls/conn.go:1089
		// _ = "end of CoverTab[21979]"
	case typeClientKeyExchange:
//line /usr/local/go/src/crypto/tls/conn.go:1090
		_go_fuzz_dep_.CoverTab[21980]++
								m = new(clientKeyExchangeMsg)
//line /usr/local/go/src/crypto/tls/conn.go:1091
		// _ = "end of CoverTab[21980]"
	case typeCertificateVerify:
//line /usr/local/go/src/crypto/tls/conn.go:1092
		_go_fuzz_dep_.CoverTab[21981]++
								m = &certificateVerifyMsg{
			hasSignatureAlgorithm: c.vers >= VersionTLS12,
		}
//line /usr/local/go/src/crypto/tls/conn.go:1095
		// _ = "end of CoverTab[21981]"
	case typeFinished:
//line /usr/local/go/src/crypto/tls/conn.go:1096
		_go_fuzz_dep_.CoverTab[21982]++
								m = new(finishedMsg)
//line /usr/local/go/src/crypto/tls/conn.go:1097
		// _ = "end of CoverTab[21982]"
	case typeEncryptedExtensions:
//line /usr/local/go/src/crypto/tls/conn.go:1098
		_go_fuzz_dep_.CoverTab[21983]++
								m = new(encryptedExtensionsMsg)
//line /usr/local/go/src/crypto/tls/conn.go:1099
		// _ = "end of CoverTab[21983]"
	case typeEndOfEarlyData:
//line /usr/local/go/src/crypto/tls/conn.go:1100
		_go_fuzz_dep_.CoverTab[21984]++
								m = new(endOfEarlyDataMsg)
//line /usr/local/go/src/crypto/tls/conn.go:1101
		// _ = "end of CoverTab[21984]"
	case typeKeyUpdate:
//line /usr/local/go/src/crypto/tls/conn.go:1102
		_go_fuzz_dep_.CoverTab[21985]++
								m = new(keyUpdateMsg)
//line /usr/local/go/src/crypto/tls/conn.go:1103
		// _ = "end of CoverTab[21985]"
	default:
//line /usr/local/go/src/crypto/tls/conn.go:1104
		_go_fuzz_dep_.CoverTab[21986]++
								return nil, c.in.setErrorLocked(c.sendAlert(alertUnexpectedMessage))
//line /usr/local/go/src/crypto/tls/conn.go:1105
		// _ = "end of CoverTab[21986]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:1106
	// _ = "end of CoverTab[21959]"
//line /usr/local/go/src/crypto/tls/conn.go:1106
	_go_fuzz_dep_.CoverTab[21960]++

//line /usr/local/go/src/crypto/tls/conn.go:1111
	data = append([]byte(nil), data...)

	if !m.unmarshal(data) {
//line /usr/local/go/src/crypto/tls/conn.go:1113
		_go_fuzz_dep_.CoverTab[21993]++
								return nil, c.in.setErrorLocked(c.sendAlert(alertUnexpectedMessage))
//line /usr/local/go/src/crypto/tls/conn.go:1114
		// _ = "end of CoverTab[21993]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1115
		_go_fuzz_dep_.CoverTab[21994]++
//line /usr/local/go/src/crypto/tls/conn.go:1115
		// _ = "end of CoverTab[21994]"
//line /usr/local/go/src/crypto/tls/conn.go:1115
	}
//line /usr/local/go/src/crypto/tls/conn.go:1115
	// _ = "end of CoverTab[21960]"
//line /usr/local/go/src/crypto/tls/conn.go:1115
	_go_fuzz_dep_.CoverTab[21961]++

							if transcript != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1117
		_go_fuzz_dep_.CoverTab[21995]++
								transcript.Write(data)
//line /usr/local/go/src/crypto/tls/conn.go:1118
		// _ = "end of CoverTab[21995]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1119
		_go_fuzz_dep_.CoverTab[21996]++
//line /usr/local/go/src/crypto/tls/conn.go:1119
		// _ = "end of CoverTab[21996]"
//line /usr/local/go/src/crypto/tls/conn.go:1119
	}
//line /usr/local/go/src/crypto/tls/conn.go:1119
	// _ = "end of CoverTab[21961]"
//line /usr/local/go/src/crypto/tls/conn.go:1119
	_go_fuzz_dep_.CoverTab[21962]++

							return m, nil
//line /usr/local/go/src/crypto/tls/conn.go:1121
	// _ = "end of CoverTab[21962]"
}

var (
	errShutdown = errors.New("tls: protocol is shutdown")
)

// Write writes data to the connection.
//line /usr/local/go/src/crypto/tls/conn.go:1128
//
//line /usr/local/go/src/crypto/tls/conn.go:1128
// As Write calls Handshake, in order to prevent indefinite blocking a deadline
//line /usr/local/go/src/crypto/tls/conn.go:1128
// must be set for both Read and Write before Write is called when the handshake
//line /usr/local/go/src/crypto/tls/conn.go:1128
// has not yet completed. See SetDeadline, SetReadDeadline, and
//line /usr/local/go/src/crypto/tls/conn.go:1128
// SetWriteDeadline.
//line /usr/local/go/src/crypto/tls/conn.go:1134
func (c *Conn) Write(b []byte) (int, error) {
//line /usr/local/go/src/crypto/tls/conn.go:1134
	_go_fuzz_dep_.CoverTab[21997]++

							for {
//line /usr/local/go/src/crypto/tls/conn.go:1136
		_go_fuzz_dep_.CoverTab[22004]++
								x := c.activeCall.Load()
								if x&1 != 0 {
//line /usr/local/go/src/crypto/tls/conn.go:1138
			_go_fuzz_dep_.CoverTab[22006]++
									return 0, net.ErrClosed
//line /usr/local/go/src/crypto/tls/conn.go:1139
			// _ = "end of CoverTab[22006]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1140
			_go_fuzz_dep_.CoverTab[22007]++
//line /usr/local/go/src/crypto/tls/conn.go:1140
			// _ = "end of CoverTab[22007]"
//line /usr/local/go/src/crypto/tls/conn.go:1140
		}
//line /usr/local/go/src/crypto/tls/conn.go:1140
		// _ = "end of CoverTab[22004]"
//line /usr/local/go/src/crypto/tls/conn.go:1140
		_go_fuzz_dep_.CoverTab[22005]++
								if c.activeCall.CompareAndSwap(x, x+2) {
//line /usr/local/go/src/crypto/tls/conn.go:1141
			_go_fuzz_dep_.CoverTab[22008]++
									break
//line /usr/local/go/src/crypto/tls/conn.go:1142
			// _ = "end of CoverTab[22008]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1143
			_go_fuzz_dep_.CoverTab[22009]++
//line /usr/local/go/src/crypto/tls/conn.go:1143
			// _ = "end of CoverTab[22009]"
//line /usr/local/go/src/crypto/tls/conn.go:1143
		}
//line /usr/local/go/src/crypto/tls/conn.go:1143
		// _ = "end of CoverTab[22005]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:1144
	// _ = "end of CoverTab[21997]"
//line /usr/local/go/src/crypto/tls/conn.go:1144
	_go_fuzz_dep_.CoverTab[21998]++
							defer c.activeCall.Add(-2)

							if err := c.Handshake(); err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1147
		_go_fuzz_dep_.CoverTab[22010]++
								return 0, err
//line /usr/local/go/src/crypto/tls/conn.go:1148
		// _ = "end of CoverTab[22010]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1149
		_go_fuzz_dep_.CoverTab[22011]++
//line /usr/local/go/src/crypto/tls/conn.go:1149
		// _ = "end of CoverTab[22011]"
//line /usr/local/go/src/crypto/tls/conn.go:1149
	}
//line /usr/local/go/src/crypto/tls/conn.go:1149
	// _ = "end of CoverTab[21998]"
//line /usr/local/go/src/crypto/tls/conn.go:1149
	_go_fuzz_dep_.CoverTab[21999]++

							c.out.Lock()
							defer c.out.Unlock()

							if err := c.out.err; err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1154
		_go_fuzz_dep_.CoverTab[22012]++
								return 0, err
//line /usr/local/go/src/crypto/tls/conn.go:1155
		// _ = "end of CoverTab[22012]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1156
		_go_fuzz_dep_.CoverTab[22013]++
//line /usr/local/go/src/crypto/tls/conn.go:1156
		// _ = "end of CoverTab[22013]"
//line /usr/local/go/src/crypto/tls/conn.go:1156
	}
//line /usr/local/go/src/crypto/tls/conn.go:1156
	// _ = "end of CoverTab[21999]"
//line /usr/local/go/src/crypto/tls/conn.go:1156
	_go_fuzz_dep_.CoverTab[22000]++

							if !c.isHandshakeComplete.Load() {
//line /usr/local/go/src/crypto/tls/conn.go:1158
		_go_fuzz_dep_.CoverTab[22014]++
								return 0, alertInternalError
//line /usr/local/go/src/crypto/tls/conn.go:1159
		// _ = "end of CoverTab[22014]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1160
		_go_fuzz_dep_.CoverTab[22015]++
//line /usr/local/go/src/crypto/tls/conn.go:1160
		// _ = "end of CoverTab[22015]"
//line /usr/local/go/src/crypto/tls/conn.go:1160
	}
//line /usr/local/go/src/crypto/tls/conn.go:1160
	// _ = "end of CoverTab[22000]"
//line /usr/local/go/src/crypto/tls/conn.go:1160
	_go_fuzz_dep_.CoverTab[22001]++

							if c.closeNotifySent {
//line /usr/local/go/src/crypto/tls/conn.go:1162
		_go_fuzz_dep_.CoverTab[22016]++
								return 0, errShutdown
//line /usr/local/go/src/crypto/tls/conn.go:1163
		// _ = "end of CoverTab[22016]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1164
		_go_fuzz_dep_.CoverTab[22017]++
//line /usr/local/go/src/crypto/tls/conn.go:1164
		// _ = "end of CoverTab[22017]"
//line /usr/local/go/src/crypto/tls/conn.go:1164
	}
//line /usr/local/go/src/crypto/tls/conn.go:1164
	// _ = "end of CoverTab[22001]"
//line /usr/local/go/src/crypto/tls/conn.go:1164
	_go_fuzz_dep_.CoverTab[22002]++

//line /usr/local/go/src/crypto/tls/conn.go:1175
	var m int
	if len(b) > 1 && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:1176
		_go_fuzz_dep_.CoverTab[22018]++
//line /usr/local/go/src/crypto/tls/conn.go:1176
		return c.vers == VersionTLS10
//line /usr/local/go/src/crypto/tls/conn.go:1176
		// _ = "end of CoverTab[22018]"
//line /usr/local/go/src/crypto/tls/conn.go:1176
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:1176
		_go_fuzz_dep_.CoverTab[22019]++
								if _, ok := c.out.cipher.(cipher.BlockMode); ok {
//line /usr/local/go/src/crypto/tls/conn.go:1177
			_go_fuzz_dep_.CoverTab[22020]++
									n, err := c.writeRecordLocked(recordTypeApplicationData, b[:1])
									if err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1179
				_go_fuzz_dep_.CoverTab[22022]++
										return n, c.out.setErrorLocked(err)
//line /usr/local/go/src/crypto/tls/conn.go:1180
				// _ = "end of CoverTab[22022]"
			} else {
//line /usr/local/go/src/crypto/tls/conn.go:1181
				_go_fuzz_dep_.CoverTab[22023]++
//line /usr/local/go/src/crypto/tls/conn.go:1181
				// _ = "end of CoverTab[22023]"
//line /usr/local/go/src/crypto/tls/conn.go:1181
			}
//line /usr/local/go/src/crypto/tls/conn.go:1181
			// _ = "end of CoverTab[22020]"
//line /usr/local/go/src/crypto/tls/conn.go:1181
			_go_fuzz_dep_.CoverTab[22021]++
									m, b = 1, b[1:]
//line /usr/local/go/src/crypto/tls/conn.go:1182
			// _ = "end of CoverTab[22021]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1183
			_go_fuzz_dep_.CoverTab[22024]++
//line /usr/local/go/src/crypto/tls/conn.go:1183
			// _ = "end of CoverTab[22024]"
//line /usr/local/go/src/crypto/tls/conn.go:1183
		}
//line /usr/local/go/src/crypto/tls/conn.go:1183
		// _ = "end of CoverTab[22019]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1184
		_go_fuzz_dep_.CoverTab[22025]++
//line /usr/local/go/src/crypto/tls/conn.go:1184
		// _ = "end of CoverTab[22025]"
//line /usr/local/go/src/crypto/tls/conn.go:1184
	}
//line /usr/local/go/src/crypto/tls/conn.go:1184
	// _ = "end of CoverTab[22002]"
//line /usr/local/go/src/crypto/tls/conn.go:1184
	_go_fuzz_dep_.CoverTab[22003]++

							n, err := c.writeRecordLocked(recordTypeApplicationData, b)
							return n + m, c.out.setErrorLocked(err)
//line /usr/local/go/src/crypto/tls/conn.go:1187
	// _ = "end of CoverTab[22003]"
}

// handleRenegotiation processes a HelloRequest handshake message.
func (c *Conn) handleRenegotiation() error {
//line /usr/local/go/src/crypto/tls/conn.go:1191
	_go_fuzz_dep_.CoverTab[22026]++
							if c.vers == VersionTLS13 {
//line /usr/local/go/src/crypto/tls/conn.go:1192
		_go_fuzz_dep_.CoverTab[22033]++
								return errors.New("tls: internal error: unexpected renegotiation")
//line /usr/local/go/src/crypto/tls/conn.go:1193
		// _ = "end of CoverTab[22033]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1194
		_go_fuzz_dep_.CoverTab[22034]++
//line /usr/local/go/src/crypto/tls/conn.go:1194
		// _ = "end of CoverTab[22034]"
//line /usr/local/go/src/crypto/tls/conn.go:1194
	}
//line /usr/local/go/src/crypto/tls/conn.go:1194
	// _ = "end of CoverTab[22026]"
//line /usr/local/go/src/crypto/tls/conn.go:1194
	_go_fuzz_dep_.CoverTab[22027]++

							msg, err := c.readHandshake(nil)
							if err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1197
		_go_fuzz_dep_.CoverTab[22035]++
								return err
//line /usr/local/go/src/crypto/tls/conn.go:1198
		// _ = "end of CoverTab[22035]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1199
		_go_fuzz_dep_.CoverTab[22036]++
//line /usr/local/go/src/crypto/tls/conn.go:1199
		// _ = "end of CoverTab[22036]"
//line /usr/local/go/src/crypto/tls/conn.go:1199
	}
//line /usr/local/go/src/crypto/tls/conn.go:1199
	// _ = "end of CoverTab[22027]"
//line /usr/local/go/src/crypto/tls/conn.go:1199
	_go_fuzz_dep_.CoverTab[22028]++

							helloReq, ok := msg.(*helloRequestMsg)
							if !ok {
//line /usr/local/go/src/crypto/tls/conn.go:1202
		_go_fuzz_dep_.CoverTab[22037]++
								c.sendAlert(alertUnexpectedMessage)
								return unexpectedMessageError(helloReq, msg)
//line /usr/local/go/src/crypto/tls/conn.go:1204
		// _ = "end of CoverTab[22037]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1205
		_go_fuzz_dep_.CoverTab[22038]++
//line /usr/local/go/src/crypto/tls/conn.go:1205
		// _ = "end of CoverTab[22038]"
//line /usr/local/go/src/crypto/tls/conn.go:1205
	}
//line /usr/local/go/src/crypto/tls/conn.go:1205
	// _ = "end of CoverTab[22028]"
//line /usr/local/go/src/crypto/tls/conn.go:1205
	_go_fuzz_dep_.CoverTab[22029]++

							if !c.isClient {
//line /usr/local/go/src/crypto/tls/conn.go:1207
		_go_fuzz_dep_.CoverTab[22039]++
								return c.sendAlert(alertNoRenegotiation)
//line /usr/local/go/src/crypto/tls/conn.go:1208
		// _ = "end of CoverTab[22039]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1209
		_go_fuzz_dep_.CoverTab[22040]++
//line /usr/local/go/src/crypto/tls/conn.go:1209
		// _ = "end of CoverTab[22040]"
//line /usr/local/go/src/crypto/tls/conn.go:1209
	}
//line /usr/local/go/src/crypto/tls/conn.go:1209
	// _ = "end of CoverTab[22029]"
//line /usr/local/go/src/crypto/tls/conn.go:1209
	_go_fuzz_dep_.CoverTab[22030]++

							switch c.config.Renegotiation {
	case RenegotiateNever:
//line /usr/local/go/src/crypto/tls/conn.go:1212
		_go_fuzz_dep_.CoverTab[22041]++
								return c.sendAlert(alertNoRenegotiation)
//line /usr/local/go/src/crypto/tls/conn.go:1213
		// _ = "end of CoverTab[22041]"
	case RenegotiateOnceAsClient:
//line /usr/local/go/src/crypto/tls/conn.go:1214
		_go_fuzz_dep_.CoverTab[22042]++
								if c.handshakes > 1 {
//line /usr/local/go/src/crypto/tls/conn.go:1215
			_go_fuzz_dep_.CoverTab[22045]++
									return c.sendAlert(alertNoRenegotiation)
//line /usr/local/go/src/crypto/tls/conn.go:1216
			// _ = "end of CoverTab[22045]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1217
			_go_fuzz_dep_.CoverTab[22046]++
//line /usr/local/go/src/crypto/tls/conn.go:1217
			// _ = "end of CoverTab[22046]"
//line /usr/local/go/src/crypto/tls/conn.go:1217
		}
//line /usr/local/go/src/crypto/tls/conn.go:1217
		// _ = "end of CoverTab[22042]"
	case RenegotiateFreelyAsClient:
//line /usr/local/go/src/crypto/tls/conn.go:1218
		_go_fuzz_dep_.CoverTab[22043]++
//line /usr/local/go/src/crypto/tls/conn.go:1218
		// _ = "end of CoverTab[22043]"

	default:
//line /usr/local/go/src/crypto/tls/conn.go:1220
		_go_fuzz_dep_.CoverTab[22044]++
								c.sendAlert(alertInternalError)
								return errors.New("tls: unknown Renegotiation value")
//line /usr/local/go/src/crypto/tls/conn.go:1222
		// _ = "end of CoverTab[22044]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:1223
	// _ = "end of CoverTab[22030]"
//line /usr/local/go/src/crypto/tls/conn.go:1223
	_go_fuzz_dep_.CoverTab[22031]++

							c.handshakeMutex.Lock()
							defer c.handshakeMutex.Unlock()

							c.isHandshakeComplete.Store(false)
							if c.handshakeErr = c.clientHandshake(context.Background()); c.handshakeErr == nil {
//line /usr/local/go/src/crypto/tls/conn.go:1229
		_go_fuzz_dep_.CoverTab[22047]++
								c.handshakes++
//line /usr/local/go/src/crypto/tls/conn.go:1230
		// _ = "end of CoverTab[22047]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1231
		_go_fuzz_dep_.CoverTab[22048]++
//line /usr/local/go/src/crypto/tls/conn.go:1231
		// _ = "end of CoverTab[22048]"
//line /usr/local/go/src/crypto/tls/conn.go:1231
	}
//line /usr/local/go/src/crypto/tls/conn.go:1231
	// _ = "end of CoverTab[22031]"
//line /usr/local/go/src/crypto/tls/conn.go:1231
	_go_fuzz_dep_.CoverTab[22032]++
							return c.handshakeErr
//line /usr/local/go/src/crypto/tls/conn.go:1232
	// _ = "end of CoverTab[22032]"
}

// handlePostHandshakeMessage processes a handshake message arrived after the
//line /usr/local/go/src/crypto/tls/conn.go:1235
// handshake is complete. Up to TLS 1.2, it indicates the start of a renegotiation.
//line /usr/local/go/src/crypto/tls/conn.go:1237
func (c *Conn) handlePostHandshakeMessage() error {
//line /usr/local/go/src/crypto/tls/conn.go:1237
	_go_fuzz_dep_.CoverTab[22049]++
							if c.vers != VersionTLS13 {
//line /usr/local/go/src/crypto/tls/conn.go:1238
		_go_fuzz_dep_.CoverTab[22053]++
								return c.handleRenegotiation()
//line /usr/local/go/src/crypto/tls/conn.go:1239
		// _ = "end of CoverTab[22053]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1240
		_go_fuzz_dep_.CoverTab[22054]++
//line /usr/local/go/src/crypto/tls/conn.go:1240
		// _ = "end of CoverTab[22054]"
//line /usr/local/go/src/crypto/tls/conn.go:1240
	}
//line /usr/local/go/src/crypto/tls/conn.go:1240
	// _ = "end of CoverTab[22049]"
//line /usr/local/go/src/crypto/tls/conn.go:1240
	_go_fuzz_dep_.CoverTab[22050]++

							msg, err := c.readHandshake(nil)
							if err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1243
		_go_fuzz_dep_.CoverTab[22055]++
								return err
//line /usr/local/go/src/crypto/tls/conn.go:1244
		// _ = "end of CoverTab[22055]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1245
		_go_fuzz_dep_.CoverTab[22056]++
//line /usr/local/go/src/crypto/tls/conn.go:1245
		// _ = "end of CoverTab[22056]"
//line /usr/local/go/src/crypto/tls/conn.go:1245
	}
//line /usr/local/go/src/crypto/tls/conn.go:1245
	// _ = "end of CoverTab[22050]"
//line /usr/local/go/src/crypto/tls/conn.go:1245
	_go_fuzz_dep_.CoverTab[22051]++

							c.retryCount++
							if c.retryCount > maxUselessRecords {
//line /usr/local/go/src/crypto/tls/conn.go:1248
		_go_fuzz_dep_.CoverTab[22057]++
								c.sendAlert(alertUnexpectedMessage)
								return c.in.setErrorLocked(errors.New("tls: too many non-advancing records"))
//line /usr/local/go/src/crypto/tls/conn.go:1250
		// _ = "end of CoverTab[22057]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1251
		_go_fuzz_dep_.CoverTab[22058]++
//line /usr/local/go/src/crypto/tls/conn.go:1251
		// _ = "end of CoverTab[22058]"
//line /usr/local/go/src/crypto/tls/conn.go:1251
	}
//line /usr/local/go/src/crypto/tls/conn.go:1251
	// _ = "end of CoverTab[22051]"
//line /usr/local/go/src/crypto/tls/conn.go:1251
	_go_fuzz_dep_.CoverTab[22052]++

							switch msg := msg.(type) {
	case *newSessionTicketMsgTLS13:
//line /usr/local/go/src/crypto/tls/conn.go:1254
		_go_fuzz_dep_.CoverTab[22059]++
								return c.handleNewSessionTicket(msg)
//line /usr/local/go/src/crypto/tls/conn.go:1255
		// _ = "end of CoverTab[22059]"
	case *keyUpdateMsg:
//line /usr/local/go/src/crypto/tls/conn.go:1256
		_go_fuzz_dep_.CoverTab[22060]++
								return c.handleKeyUpdate(msg)
//line /usr/local/go/src/crypto/tls/conn.go:1257
		// _ = "end of CoverTab[22060]"
	default:
//line /usr/local/go/src/crypto/tls/conn.go:1258
		_go_fuzz_dep_.CoverTab[22061]++
								c.sendAlert(alertUnexpectedMessage)
								return fmt.Errorf("tls: received unexpected handshake message of type %T", msg)
//line /usr/local/go/src/crypto/tls/conn.go:1260
		// _ = "end of CoverTab[22061]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:1261
	// _ = "end of CoverTab[22052]"
}

func (c *Conn) handleKeyUpdate(keyUpdate *keyUpdateMsg) error {
//line /usr/local/go/src/crypto/tls/conn.go:1264
	_go_fuzz_dep_.CoverTab[22062]++
							cipherSuite := cipherSuiteTLS13ByID(c.cipherSuite)
							if cipherSuite == nil {
//line /usr/local/go/src/crypto/tls/conn.go:1266
		_go_fuzz_dep_.CoverTab[22065]++
								return c.in.setErrorLocked(c.sendAlert(alertInternalError))
//line /usr/local/go/src/crypto/tls/conn.go:1267
		// _ = "end of CoverTab[22065]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1268
		_go_fuzz_dep_.CoverTab[22066]++
//line /usr/local/go/src/crypto/tls/conn.go:1268
		// _ = "end of CoverTab[22066]"
//line /usr/local/go/src/crypto/tls/conn.go:1268
	}
//line /usr/local/go/src/crypto/tls/conn.go:1268
	// _ = "end of CoverTab[22062]"
//line /usr/local/go/src/crypto/tls/conn.go:1268
	_go_fuzz_dep_.CoverTab[22063]++

							newSecret := cipherSuite.nextTrafficSecret(c.in.trafficSecret)
							c.in.setTrafficSecret(cipherSuite, newSecret)

							if keyUpdate.updateRequested {
//line /usr/local/go/src/crypto/tls/conn.go:1273
		_go_fuzz_dep_.CoverTab[22067]++
								c.out.Lock()
								defer c.out.Unlock()

								msg := &keyUpdateMsg{}
								msgBytes, err := msg.marshal()
								if err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1279
			_go_fuzz_dep_.CoverTab[22070]++
									return err
//line /usr/local/go/src/crypto/tls/conn.go:1280
			// _ = "end of CoverTab[22070]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1281
			_go_fuzz_dep_.CoverTab[22071]++
//line /usr/local/go/src/crypto/tls/conn.go:1281
			// _ = "end of CoverTab[22071]"
//line /usr/local/go/src/crypto/tls/conn.go:1281
		}
//line /usr/local/go/src/crypto/tls/conn.go:1281
		// _ = "end of CoverTab[22067]"
//line /usr/local/go/src/crypto/tls/conn.go:1281
		_go_fuzz_dep_.CoverTab[22068]++
								_, err = c.writeRecordLocked(recordTypeHandshake, msgBytes)
								if err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1283
			_go_fuzz_dep_.CoverTab[22072]++

									c.out.setErrorLocked(err)
									return nil
//line /usr/local/go/src/crypto/tls/conn.go:1286
			// _ = "end of CoverTab[22072]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1287
			_go_fuzz_dep_.CoverTab[22073]++
//line /usr/local/go/src/crypto/tls/conn.go:1287
			// _ = "end of CoverTab[22073]"
//line /usr/local/go/src/crypto/tls/conn.go:1287
		}
//line /usr/local/go/src/crypto/tls/conn.go:1287
		// _ = "end of CoverTab[22068]"
//line /usr/local/go/src/crypto/tls/conn.go:1287
		_go_fuzz_dep_.CoverTab[22069]++

								newSecret := cipherSuite.nextTrafficSecret(c.out.trafficSecret)
								c.out.setTrafficSecret(cipherSuite, newSecret)
//line /usr/local/go/src/crypto/tls/conn.go:1290
		// _ = "end of CoverTab[22069]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1291
		_go_fuzz_dep_.CoverTab[22074]++
//line /usr/local/go/src/crypto/tls/conn.go:1291
		// _ = "end of CoverTab[22074]"
//line /usr/local/go/src/crypto/tls/conn.go:1291
	}
//line /usr/local/go/src/crypto/tls/conn.go:1291
	// _ = "end of CoverTab[22063]"
//line /usr/local/go/src/crypto/tls/conn.go:1291
	_go_fuzz_dep_.CoverTab[22064]++

							return nil
//line /usr/local/go/src/crypto/tls/conn.go:1293
	// _ = "end of CoverTab[22064]"
}

// Read reads data from the connection.
//line /usr/local/go/src/crypto/tls/conn.go:1296
//
//line /usr/local/go/src/crypto/tls/conn.go:1296
// As Read calls Handshake, in order to prevent indefinite blocking a deadline
//line /usr/local/go/src/crypto/tls/conn.go:1296
// must be set for both Read and Write before Read is called when the handshake
//line /usr/local/go/src/crypto/tls/conn.go:1296
// has not yet completed. See SetDeadline, SetReadDeadline, and
//line /usr/local/go/src/crypto/tls/conn.go:1296
// SetWriteDeadline.
//line /usr/local/go/src/crypto/tls/conn.go:1302
func (c *Conn) Read(b []byte) (int, error) {
//line /usr/local/go/src/crypto/tls/conn.go:1302
	_go_fuzz_dep_.CoverTab[22075]++
							if err := c.Handshake(); err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1303
		_go_fuzz_dep_.CoverTab[22080]++
								return 0, err
//line /usr/local/go/src/crypto/tls/conn.go:1304
		// _ = "end of CoverTab[22080]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1305
		_go_fuzz_dep_.CoverTab[22081]++
//line /usr/local/go/src/crypto/tls/conn.go:1305
		// _ = "end of CoverTab[22081]"
//line /usr/local/go/src/crypto/tls/conn.go:1305
	}
//line /usr/local/go/src/crypto/tls/conn.go:1305
	// _ = "end of CoverTab[22075]"
//line /usr/local/go/src/crypto/tls/conn.go:1305
	_go_fuzz_dep_.CoverTab[22076]++
							if len(b) == 0 {
//line /usr/local/go/src/crypto/tls/conn.go:1306
		_go_fuzz_dep_.CoverTab[22082]++

//line /usr/local/go/src/crypto/tls/conn.go:1309
		return 0, nil
//line /usr/local/go/src/crypto/tls/conn.go:1309
		// _ = "end of CoverTab[22082]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1310
		_go_fuzz_dep_.CoverTab[22083]++
//line /usr/local/go/src/crypto/tls/conn.go:1310
		// _ = "end of CoverTab[22083]"
//line /usr/local/go/src/crypto/tls/conn.go:1310
	}
//line /usr/local/go/src/crypto/tls/conn.go:1310
	// _ = "end of CoverTab[22076]"
//line /usr/local/go/src/crypto/tls/conn.go:1310
	_go_fuzz_dep_.CoverTab[22077]++

							c.in.Lock()
							defer c.in.Unlock()

							for c.input.Len() == 0 {
//line /usr/local/go/src/crypto/tls/conn.go:1315
		_go_fuzz_dep_.CoverTab[22084]++
								if err := c.readRecord(); err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1316
			_go_fuzz_dep_.CoverTab[22086]++
									return 0, err
//line /usr/local/go/src/crypto/tls/conn.go:1317
			// _ = "end of CoverTab[22086]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1318
			_go_fuzz_dep_.CoverTab[22087]++
//line /usr/local/go/src/crypto/tls/conn.go:1318
			// _ = "end of CoverTab[22087]"
//line /usr/local/go/src/crypto/tls/conn.go:1318
		}
//line /usr/local/go/src/crypto/tls/conn.go:1318
		// _ = "end of CoverTab[22084]"
//line /usr/local/go/src/crypto/tls/conn.go:1318
		_go_fuzz_dep_.CoverTab[22085]++
								for c.hand.Len() > 0 {
//line /usr/local/go/src/crypto/tls/conn.go:1319
			_go_fuzz_dep_.CoverTab[22088]++
									if err := c.handlePostHandshakeMessage(); err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1320
				_go_fuzz_dep_.CoverTab[22089]++
										return 0, err
//line /usr/local/go/src/crypto/tls/conn.go:1321
				// _ = "end of CoverTab[22089]"
			} else {
//line /usr/local/go/src/crypto/tls/conn.go:1322
				_go_fuzz_dep_.CoverTab[22090]++
//line /usr/local/go/src/crypto/tls/conn.go:1322
				// _ = "end of CoverTab[22090]"
//line /usr/local/go/src/crypto/tls/conn.go:1322
			}
//line /usr/local/go/src/crypto/tls/conn.go:1322
			// _ = "end of CoverTab[22088]"
		}
//line /usr/local/go/src/crypto/tls/conn.go:1323
		// _ = "end of CoverTab[22085]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:1324
	// _ = "end of CoverTab[22077]"
//line /usr/local/go/src/crypto/tls/conn.go:1324
	_go_fuzz_dep_.CoverTab[22078]++

							n, _ := c.input.Read(b)

//line /usr/local/go/src/crypto/tls/conn.go:1335
	if n != 0 && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:1335
		_go_fuzz_dep_.CoverTab[22091]++
//line /usr/local/go/src/crypto/tls/conn.go:1335
		return c.input.Len() == 0
//line /usr/local/go/src/crypto/tls/conn.go:1335
		// _ = "end of CoverTab[22091]"
//line /usr/local/go/src/crypto/tls/conn.go:1335
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:1335
		_go_fuzz_dep_.CoverTab[22092]++
//line /usr/local/go/src/crypto/tls/conn.go:1335
		return c.rawInput.Len() > 0
//line /usr/local/go/src/crypto/tls/conn.go:1335
		// _ = "end of CoverTab[22092]"
//line /usr/local/go/src/crypto/tls/conn.go:1335
	}() && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:1335
		_go_fuzz_dep_.CoverTab[22093]++
//line /usr/local/go/src/crypto/tls/conn.go:1335
		return recordType(c.rawInput.Bytes()[0]) == recordTypeAlert
								// _ = "end of CoverTab[22093]"
//line /usr/local/go/src/crypto/tls/conn.go:1336
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:1336
		_go_fuzz_dep_.CoverTab[22094]++
								if err := c.readRecord(); err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1337
			_go_fuzz_dep_.CoverTab[22095]++
									return n, err
//line /usr/local/go/src/crypto/tls/conn.go:1338
			// _ = "end of CoverTab[22095]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1339
			_go_fuzz_dep_.CoverTab[22096]++
//line /usr/local/go/src/crypto/tls/conn.go:1339
			// _ = "end of CoverTab[22096]"
//line /usr/local/go/src/crypto/tls/conn.go:1339
		}
//line /usr/local/go/src/crypto/tls/conn.go:1339
		// _ = "end of CoverTab[22094]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1340
		_go_fuzz_dep_.CoverTab[22097]++
//line /usr/local/go/src/crypto/tls/conn.go:1340
		// _ = "end of CoverTab[22097]"
//line /usr/local/go/src/crypto/tls/conn.go:1340
	}
//line /usr/local/go/src/crypto/tls/conn.go:1340
	// _ = "end of CoverTab[22078]"
//line /usr/local/go/src/crypto/tls/conn.go:1340
	_go_fuzz_dep_.CoverTab[22079]++

							return n, nil
//line /usr/local/go/src/crypto/tls/conn.go:1342
	// _ = "end of CoverTab[22079]"
}

// Close closes the connection.
func (c *Conn) Close() error {
//line /usr/local/go/src/crypto/tls/conn.go:1346
	_go_fuzz_dep_.CoverTab[22098]++
	// Interlock with Conn.Write above.
	var x int32
	for {
//line /usr/local/go/src/crypto/tls/conn.go:1349
		_go_fuzz_dep_.CoverTab[22103]++
								x = c.activeCall.Load()
								if x&1 != 0 {
//line /usr/local/go/src/crypto/tls/conn.go:1351
			_go_fuzz_dep_.CoverTab[22105]++
									return net.ErrClosed
//line /usr/local/go/src/crypto/tls/conn.go:1352
			// _ = "end of CoverTab[22105]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1353
			_go_fuzz_dep_.CoverTab[22106]++
//line /usr/local/go/src/crypto/tls/conn.go:1353
			// _ = "end of CoverTab[22106]"
//line /usr/local/go/src/crypto/tls/conn.go:1353
		}
//line /usr/local/go/src/crypto/tls/conn.go:1353
		// _ = "end of CoverTab[22103]"
//line /usr/local/go/src/crypto/tls/conn.go:1353
		_go_fuzz_dep_.CoverTab[22104]++
								if c.activeCall.CompareAndSwap(x, x|1) {
//line /usr/local/go/src/crypto/tls/conn.go:1354
			_go_fuzz_dep_.CoverTab[22107]++
									break
//line /usr/local/go/src/crypto/tls/conn.go:1355
			// _ = "end of CoverTab[22107]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1356
			_go_fuzz_dep_.CoverTab[22108]++
//line /usr/local/go/src/crypto/tls/conn.go:1356
			// _ = "end of CoverTab[22108]"
//line /usr/local/go/src/crypto/tls/conn.go:1356
		}
//line /usr/local/go/src/crypto/tls/conn.go:1356
		// _ = "end of CoverTab[22104]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:1357
	// _ = "end of CoverTab[22098]"
//line /usr/local/go/src/crypto/tls/conn.go:1357
	_go_fuzz_dep_.CoverTab[22099]++
							if x != 0 {
//line /usr/local/go/src/crypto/tls/conn.go:1358
		_go_fuzz_dep_.CoverTab[22109]++

//line /usr/local/go/src/crypto/tls/conn.go:1365
		return c.conn.Close()
//line /usr/local/go/src/crypto/tls/conn.go:1365
		// _ = "end of CoverTab[22109]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1366
		_go_fuzz_dep_.CoverTab[22110]++
//line /usr/local/go/src/crypto/tls/conn.go:1366
		// _ = "end of CoverTab[22110]"
//line /usr/local/go/src/crypto/tls/conn.go:1366
	}
//line /usr/local/go/src/crypto/tls/conn.go:1366
	// _ = "end of CoverTab[22099]"
//line /usr/local/go/src/crypto/tls/conn.go:1366
	_go_fuzz_dep_.CoverTab[22100]++

							var alertErr error
							if c.isHandshakeComplete.Load() {
//line /usr/local/go/src/crypto/tls/conn.go:1369
		_go_fuzz_dep_.CoverTab[22111]++
								if err := c.closeNotify(); err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1370
			_go_fuzz_dep_.CoverTab[22112]++
									alertErr = fmt.Errorf("tls: failed to send closeNotify alert (but connection was closed anyway): %w", err)
//line /usr/local/go/src/crypto/tls/conn.go:1371
			// _ = "end of CoverTab[22112]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1372
			_go_fuzz_dep_.CoverTab[22113]++
//line /usr/local/go/src/crypto/tls/conn.go:1372
			// _ = "end of CoverTab[22113]"
//line /usr/local/go/src/crypto/tls/conn.go:1372
		}
//line /usr/local/go/src/crypto/tls/conn.go:1372
		// _ = "end of CoverTab[22111]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1373
		_go_fuzz_dep_.CoverTab[22114]++
//line /usr/local/go/src/crypto/tls/conn.go:1373
		// _ = "end of CoverTab[22114]"
//line /usr/local/go/src/crypto/tls/conn.go:1373
	}
//line /usr/local/go/src/crypto/tls/conn.go:1373
	// _ = "end of CoverTab[22100]"
//line /usr/local/go/src/crypto/tls/conn.go:1373
	_go_fuzz_dep_.CoverTab[22101]++

							if err := c.conn.Close(); err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1375
		_go_fuzz_dep_.CoverTab[22115]++
								return err
//line /usr/local/go/src/crypto/tls/conn.go:1376
		// _ = "end of CoverTab[22115]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1377
		_go_fuzz_dep_.CoverTab[22116]++
//line /usr/local/go/src/crypto/tls/conn.go:1377
		// _ = "end of CoverTab[22116]"
//line /usr/local/go/src/crypto/tls/conn.go:1377
	}
//line /usr/local/go/src/crypto/tls/conn.go:1377
	// _ = "end of CoverTab[22101]"
//line /usr/local/go/src/crypto/tls/conn.go:1377
	_go_fuzz_dep_.CoverTab[22102]++
							return alertErr
//line /usr/local/go/src/crypto/tls/conn.go:1378
	// _ = "end of CoverTab[22102]"
}

var errEarlyCloseWrite = errors.New("tls: CloseWrite called before handshake complete")

// CloseWrite shuts down the writing side of the connection. It should only be
//line /usr/local/go/src/crypto/tls/conn.go:1383
// called once the handshake has completed and does not call CloseWrite on the
//line /usr/local/go/src/crypto/tls/conn.go:1383
// underlying connection. Most callers should just use Close.
//line /usr/local/go/src/crypto/tls/conn.go:1386
func (c *Conn) CloseWrite() error {
//line /usr/local/go/src/crypto/tls/conn.go:1386
	_go_fuzz_dep_.CoverTab[22117]++
							if !c.isHandshakeComplete.Load() {
//line /usr/local/go/src/crypto/tls/conn.go:1387
		_go_fuzz_dep_.CoverTab[22119]++
								return errEarlyCloseWrite
//line /usr/local/go/src/crypto/tls/conn.go:1388
		// _ = "end of CoverTab[22119]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1389
		_go_fuzz_dep_.CoverTab[22120]++
//line /usr/local/go/src/crypto/tls/conn.go:1389
		// _ = "end of CoverTab[22120]"
//line /usr/local/go/src/crypto/tls/conn.go:1389
	}
//line /usr/local/go/src/crypto/tls/conn.go:1389
	// _ = "end of CoverTab[22117]"
//line /usr/local/go/src/crypto/tls/conn.go:1389
	_go_fuzz_dep_.CoverTab[22118]++

							return c.closeNotify()
//line /usr/local/go/src/crypto/tls/conn.go:1391
	// _ = "end of CoverTab[22118]"
}

func (c *Conn) closeNotify() error {
//line /usr/local/go/src/crypto/tls/conn.go:1394
	_go_fuzz_dep_.CoverTab[22121]++
							c.out.Lock()
							defer c.out.Unlock()

							if !c.closeNotifySent {
//line /usr/local/go/src/crypto/tls/conn.go:1398
		_go_fuzz_dep_.CoverTab[22123]++

								c.SetWriteDeadline(time.Now().Add(time.Second * 5))
								c.closeNotifyErr = c.sendAlertLocked(alertCloseNotify)
								c.closeNotifySent = true

								c.SetWriteDeadline(time.Now())
//line /usr/local/go/src/crypto/tls/conn.go:1404
		// _ = "end of CoverTab[22123]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1405
		_go_fuzz_dep_.CoverTab[22124]++
//line /usr/local/go/src/crypto/tls/conn.go:1405
		// _ = "end of CoverTab[22124]"
//line /usr/local/go/src/crypto/tls/conn.go:1405
	}
//line /usr/local/go/src/crypto/tls/conn.go:1405
	// _ = "end of CoverTab[22121]"
//line /usr/local/go/src/crypto/tls/conn.go:1405
	_go_fuzz_dep_.CoverTab[22122]++
							return c.closeNotifyErr
//line /usr/local/go/src/crypto/tls/conn.go:1406
	// _ = "end of CoverTab[22122]"
}

// Handshake runs the client or server handshake
//line /usr/local/go/src/crypto/tls/conn.go:1409
// protocol if it has not yet been run.
//line /usr/local/go/src/crypto/tls/conn.go:1409
//
//line /usr/local/go/src/crypto/tls/conn.go:1409
// Most uses of this package need not call Handshake explicitly: the
//line /usr/local/go/src/crypto/tls/conn.go:1409
// first Read or Write will call it automatically.
//line /usr/local/go/src/crypto/tls/conn.go:1409
//
//line /usr/local/go/src/crypto/tls/conn.go:1409
// For control over canceling or setting a timeout on a handshake, use
//line /usr/local/go/src/crypto/tls/conn.go:1409
// HandshakeContext or the Dialer's DialContext method instead.
//line /usr/local/go/src/crypto/tls/conn.go:1417
func (c *Conn) Handshake() error {
//line /usr/local/go/src/crypto/tls/conn.go:1417
	_go_fuzz_dep_.CoverTab[22125]++
							return c.HandshakeContext(context.Background())
//line /usr/local/go/src/crypto/tls/conn.go:1418
	// _ = "end of CoverTab[22125]"
}

// HandshakeContext runs the client or server handshake
//line /usr/local/go/src/crypto/tls/conn.go:1421
// protocol if it has not yet been run.
//line /usr/local/go/src/crypto/tls/conn.go:1421
//
//line /usr/local/go/src/crypto/tls/conn.go:1421
// The provided Context must be non-nil. If the context is canceled before
//line /usr/local/go/src/crypto/tls/conn.go:1421
// the handshake is complete, the handshake is interrupted and an error is returned.
//line /usr/local/go/src/crypto/tls/conn.go:1421
// Once the handshake has completed, cancellation of the context will not affect the
//line /usr/local/go/src/crypto/tls/conn.go:1421
// connection.
//line /usr/local/go/src/crypto/tls/conn.go:1421
//
//line /usr/local/go/src/crypto/tls/conn.go:1421
// Most uses of this package need not call HandshakeContext explicitly: the
//line /usr/local/go/src/crypto/tls/conn.go:1421
// first Read or Write will call it automatically.
//line /usr/local/go/src/crypto/tls/conn.go:1431
func (c *Conn) HandshakeContext(ctx context.Context) error {
//line /usr/local/go/src/crypto/tls/conn.go:1431
	_go_fuzz_dep_.CoverTab[22126]++

//line /usr/local/go/src/crypto/tls/conn.go:1434
	return c.handshakeContext(ctx)
//line /usr/local/go/src/crypto/tls/conn.go:1434
	// _ = "end of CoverTab[22126]"
}

func (c *Conn) handshakeContext(ctx context.Context) (ret error) {
//line /usr/local/go/src/crypto/tls/conn.go:1437
	_go_fuzz_dep_.CoverTab[22127]++

//line /usr/local/go/src/crypto/tls/conn.go:1441
	if c.isHandshakeComplete.Load() {
//line /usr/local/go/src/crypto/tls/conn.go:1441
		_go_fuzz_dep_.CoverTab[22135]++
								return nil
//line /usr/local/go/src/crypto/tls/conn.go:1442
		// _ = "end of CoverTab[22135]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1443
		_go_fuzz_dep_.CoverTab[22136]++
//line /usr/local/go/src/crypto/tls/conn.go:1443
		// _ = "end of CoverTab[22136]"
//line /usr/local/go/src/crypto/tls/conn.go:1443
	}
//line /usr/local/go/src/crypto/tls/conn.go:1443
	// _ = "end of CoverTab[22127]"
//line /usr/local/go/src/crypto/tls/conn.go:1443
	_go_fuzz_dep_.CoverTab[22128]++

							handshakeCtx, cancel := context.WithCancel(ctx)

//line /usr/local/go/src/crypto/tls/conn.go:1449
	defer cancel()

//line /usr/local/go/src/crypto/tls/conn.go:1456
	if ctx.Done() != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1456
		_go_fuzz_dep_.CoverTab[22137]++
								done := make(chan struct{})
								interruptRes := make(chan error, 1)
								defer func() {
//line /usr/local/go/src/crypto/tls/conn.go:1459
			_go_fuzz_dep_.CoverTab[22139]++
									close(done)
									if ctxErr := <-interruptRes; ctxErr != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1461
				_go_fuzz_dep_.CoverTab[22140]++

										ret = ctxErr
//line /usr/local/go/src/crypto/tls/conn.go:1463
				// _ = "end of CoverTab[22140]"
			} else {
//line /usr/local/go/src/crypto/tls/conn.go:1464
				_go_fuzz_dep_.CoverTab[22141]++
//line /usr/local/go/src/crypto/tls/conn.go:1464
				// _ = "end of CoverTab[22141]"
//line /usr/local/go/src/crypto/tls/conn.go:1464
			}
//line /usr/local/go/src/crypto/tls/conn.go:1464
			// _ = "end of CoverTab[22139]"
		}()
//line /usr/local/go/src/crypto/tls/conn.go:1465
		// _ = "end of CoverTab[22137]"
//line /usr/local/go/src/crypto/tls/conn.go:1465
		_go_fuzz_dep_.CoverTab[22138]++
//line /usr/local/go/src/crypto/tls/conn.go:1465
		_curRoutineNum12_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/crypto/tls/conn.go:1465
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum12_)
								go func() {
//line /usr/local/go/src/crypto/tls/conn.go:1466
			_go_fuzz_dep_.CoverTab[22142]++
//line /usr/local/go/src/crypto/tls/conn.go:1466
			defer func() {
//line /usr/local/go/src/crypto/tls/conn.go:1466
				_go_fuzz_dep_.CoverTab[22143]++
//line /usr/local/go/src/crypto/tls/conn.go:1466
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum12_)
//line /usr/local/go/src/crypto/tls/conn.go:1466
				// _ = "end of CoverTab[22143]"
//line /usr/local/go/src/crypto/tls/conn.go:1466
			}()
									select {
			case <-handshakeCtx.Done():
//line /usr/local/go/src/crypto/tls/conn.go:1468
				_go_fuzz_dep_.CoverTab[22144]++

										_ = c.conn.Close()
										interruptRes <- handshakeCtx.Err()
//line /usr/local/go/src/crypto/tls/conn.go:1471
				// _ = "end of CoverTab[22144]"
			case <-done:
//line /usr/local/go/src/crypto/tls/conn.go:1472
				_go_fuzz_dep_.CoverTab[22145]++
										interruptRes <- nil
//line /usr/local/go/src/crypto/tls/conn.go:1473
				// _ = "end of CoverTab[22145]"
			}
//line /usr/local/go/src/crypto/tls/conn.go:1474
			// _ = "end of CoverTab[22142]"
		}()
//line /usr/local/go/src/crypto/tls/conn.go:1475
		// _ = "end of CoverTab[22138]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1476
		_go_fuzz_dep_.CoverTab[22146]++
//line /usr/local/go/src/crypto/tls/conn.go:1476
		// _ = "end of CoverTab[22146]"
//line /usr/local/go/src/crypto/tls/conn.go:1476
	}
//line /usr/local/go/src/crypto/tls/conn.go:1476
	// _ = "end of CoverTab[22128]"
//line /usr/local/go/src/crypto/tls/conn.go:1476
	_go_fuzz_dep_.CoverTab[22129]++

							c.handshakeMutex.Lock()
							defer c.handshakeMutex.Unlock()

							if err := c.handshakeErr; err != nil {
//line /usr/local/go/src/crypto/tls/conn.go:1481
		_go_fuzz_dep_.CoverTab[22147]++
								return err
//line /usr/local/go/src/crypto/tls/conn.go:1482
		// _ = "end of CoverTab[22147]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1483
		_go_fuzz_dep_.CoverTab[22148]++
//line /usr/local/go/src/crypto/tls/conn.go:1483
		// _ = "end of CoverTab[22148]"
//line /usr/local/go/src/crypto/tls/conn.go:1483
	}
//line /usr/local/go/src/crypto/tls/conn.go:1483
	// _ = "end of CoverTab[22129]"
//line /usr/local/go/src/crypto/tls/conn.go:1483
	_go_fuzz_dep_.CoverTab[22130]++
							if c.isHandshakeComplete.Load() {
//line /usr/local/go/src/crypto/tls/conn.go:1484
		_go_fuzz_dep_.CoverTab[22149]++
								return nil
//line /usr/local/go/src/crypto/tls/conn.go:1485
		// _ = "end of CoverTab[22149]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1486
		_go_fuzz_dep_.CoverTab[22150]++
//line /usr/local/go/src/crypto/tls/conn.go:1486
		// _ = "end of CoverTab[22150]"
//line /usr/local/go/src/crypto/tls/conn.go:1486
	}
//line /usr/local/go/src/crypto/tls/conn.go:1486
	// _ = "end of CoverTab[22130]"
//line /usr/local/go/src/crypto/tls/conn.go:1486
	_go_fuzz_dep_.CoverTab[22131]++

							c.in.Lock()
							defer c.in.Unlock()

							c.handshakeErr = c.handshakeFn(handshakeCtx)
							if c.handshakeErr == nil {
//line /usr/local/go/src/crypto/tls/conn.go:1492
		_go_fuzz_dep_.CoverTab[22151]++
								c.handshakes++
//line /usr/local/go/src/crypto/tls/conn.go:1493
		// _ = "end of CoverTab[22151]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1494
		_go_fuzz_dep_.CoverTab[22152]++

//line /usr/local/go/src/crypto/tls/conn.go:1497
		c.flush()
//line /usr/local/go/src/crypto/tls/conn.go:1497
		// _ = "end of CoverTab[22152]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:1498
	// _ = "end of CoverTab[22131]"
//line /usr/local/go/src/crypto/tls/conn.go:1498
	_go_fuzz_dep_.CoverTab[22132]++

							if c.handshakeErr == nil && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:1500
		_go_fuzz_dep_.CoverTab[22153]++
//line /usr/local/go/src/crypto/tls/conn.go:1500
		return !c.isHandshakeComplete.Load()
//line /usr/local/go/src/crypto/tls/conn.go:1500
		// _ = "end of CoverTab[22153]"
//line /usr/local/go/src/crypto/tls/conn.go:1500
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:1500
		_go_fuzz_dep_.CoverTab[22154]++
								c.handshakeErr = errors.New("tls: internal error: handshake should have had a result")
//line /usr/local/go/src/crypto/tls/conn.go:1501
		// _ = "end of CoverTab[22154]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1502
		_go_fuzz_dep_.CoverTab[22155]++
//line /usr/local/go/src/crypto/tls/conn.go:1502
		// _ = "end of CoverTab[22155]"
//line /usr/local/go/src/crypto/tls/conn.go:1502
	}
//line /usr/local/go/src/crypto/tls/conn.go:1502
	// _ = "end of CoverTab[22132]"
//line /usr/local/go/src/crypto/tls/conn.go:1502
	_go_fuzz_dep_.CoverTab[22133]++
							if c.handshakeErr != nil && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:1503
		_go_fuzz_dep_.CoverTab[22156]++
//line /usr/local/go/src/crypto/tls/conn.go:1503
		return c.isHandshakeComplete.Load()
//line /usr/local/go/src/crypto/tls/conn.go:1503
		// _ = "end of CoverTab[22156]"
//line /usr/local/go/src/crypto/tls/conn.go:1503
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:1503
		_go_fuzz_dep_.CoverTab[22157]++
								panic("tls: internal error: handshake returned an error but is marked successful")
//line /usr/local/go/src/crypto/tls/conn.go:1504
		// _ = "end of CoverTab[22157]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1505
		_go_fuzz_dep_.CoverTab[22158]++
//line /usr/local/go/src/crypto/tls/conn.go:1505
		// _ = "end of CoverTab[22158]"
//line /usr/local/go/src/crypto/tls/conn.go:1505
	}
//line /usr/local/go/src/crypto/tls/conn.go:1505
	// _ = "end of CoverTab[22133]"
//line /usr/local/go/src/crypto/tls/conn.go:1505
	_go_fuzz_dep_.CoverTab[22134]++

							return c.handshakeErr
//line /usr/local/go/src/crypto/tls/conn.go:1507
	// _ = "end of CoverTab[22134]"
}

// ConnectionState returns basic TLS details about the connection.
func (c *Conn) ConnectionState() ConnectionState {
//line /usr/local/go/src/crypto/tls/conn.go:1511
	_go_fuzz_dep_.CoverTab[22159]++
							c.handshakeMutex.Lock()
							defer c.handshakeMutex.Unlock()
							return c.connectionStateLocked()
//line /usr/local/go/src/crypto/tls/conn.go:1514
	// _ = "end of CoverTab[22159]"
}

func (c *Conn) connectionStateLocked() ConnectionState {
//line /usr/local/go/src/crypto/tls/conn.go:1517
	_go_fuzz_dep_.CoverTab[22160]++
							var state ConnectionState
							state.HandshakeComplete = c.isHandshakeComplete.Load()
							state.Version = c.vers
							state.NegotiatedProtocol = c.clientProtocol
							state.DidResume = c.didResume
							state.NegotiatedProtocolIsMutual = true
							state.ServerName = c.serverName
							state.CipherSuite = c.cipherSuite
							state.PeerCertificates = c.peerCertificates
							state.VerifiedChains = c.verifiedChains
							state.SignedCertificateTimestamps = c.scts
							state.OCSPResponse = c.ocspResponse
							if !c.didResume && func() bool {
//line /usr/local/go/src/crypto/tls/conn.go:1530
		_go_fuzz_dep_.CoverTab[22163]++
//line /usr/local/go/src/crypto/tls/conn.go:1530
		return c.vers != VersionTLS13
//line /usr/local/go/src/crypto/tls/conn.go:1530
		// _ = "end of CoverTab[22163]"
//line /usr/local/go/src/crypto/tls/conn.go:1530
	}() {
//line /usr/local/go/src/crypto/tls/conn.go:1530
		_go_fuzz_dep_.CoverTab[22164]++
								if c.clientFinishedIsFirst {
//line /usr/local/go/src/crypto/tls/conn.go:1531
			_go_fuzz_dep_.CoverTab[22165]++
									state.TLSUnique = c.clientFinished[:]
//line /usr/local/go/src/crypto/tls/conn.go:1532
			// _ = "end of CoverTab[22165]"
		} else {
//line /usr/local/go/src/crypto/tls/conn.go:1533
			_go_fuzz_dep_.CoverTab[22166]++
									state.TLSUnique = c.serverFinished[:]
//line /usr/local/go/src/crypto/tls/conn.go:1534
			// _ = "end of CoverTab[22166]"
		}
//line /usr/local/go/src/crypto/tls/conn.go:1535
		// _ = "end of CoverTab[22164]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1536
		_go_fuzz_dep_.CoverTab[22167]++
//line /usr/local/go/src/crypto/tls/conn.go:1536
		// _ = "end of CoverTab[22167]"
//line /usr/local/go/src/crypto/tls/conn.go:1536
	}
//line /usr/local/go/src/crypto/tls/conn.go:1536
	// _ = "end of CoverTab[22160]"
//line /usr/local/go/src/crypto/tls/conn.go:1536
	_go_fuzz_dep_.CoverTab[22161]++
							if c.config.Renegotiation != RenegotiateNever {
//line /usr/local/go/src/crypto/tls/conn.go:1537
		_go_fuzz_dep_.CoverTab[22168]++
								state.ekm = noExportedKeyingMaterial
//line /usr/local/go/src/crypto/tls/conn.go:1538
		// _ = "end of CoverTab[22168]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1539
		_go_fuzz_dep_.CoverTab[22169]++
								state.ekm = c.ekm
//line /usr/local/go/src/crypto/tls/conn.go:1540
		// _ = "end of CoverTab[22169]"
	}
//line /usr/local/go/src/crypto/tls/conn.go:1541
	// _ = "end of CoverTab[22161]"
//line /usr/local/go/src/crypto/tls/conn.go:1541
	_go_fuzz_dep_.CoverTab[22162]++
							return state
//line /usr/local/go/src/crypto/tls/conn.go:1542
	// _ = "end of CoverTab[22162]"
}

// OCSPResponse returns the stapled OCSP response from the TLS server, if
//line /usr/local/go/src/crypto/tls/conn.go:1545
// any. (Only valid for client connections.)
//line /usr/local/go/src/crypto/tls/conn.go:1547
func (c *Conn) OCSPResponse() []byte {
//line /usr/local/go/src/crypto/tls/conn.go:1547
	_go_fuzz_dep_.CoverTab[22170]++
							c.handshakeMutex.Lock()
							defer c.handshakeMutex.Unlock()

							return c.ocspResponse
//line /usr/local/go/src/crypto/tls/conn.go:1551
	// _ = "end of CoverTab[22170]"
}

// VerifyHostname checks that the peer certificate chain is valid for
//line /usr/local/go/src/crypto/tls/conn.go:1554
// connecting to host. If so, it returns nil; if not, it returns an error
//line /usr/local/go/src/crypto/tls/conn.go:1554
// describing the problem.
//line /usr/local/go/src/crypto/tls/conn.go:1557
func (c *Conn) VerifyHostname(host string) error {
//line /usr/local/go/src/crypto/tls/conn.go:1557
	_go_fuzz_dep_.CoverTab[22171]++
							c.handshakeMutex.Lock()
							defer c.handshakeMutex.Unlock()
							if !c.isClient {
//line /usr/local/go/src/crypto/tls/conn.go:1560
		_go_fuzz_dep_.CoverTab[22175]++
								return errors.New("tls: VerifyHostname called on TLS server connection")
//line /usr/local/go/src/crypto/tls/conn.go:1561
		// _ = "end of CoverTab[22175]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1562
		_go_fuzz_dep_.CoverTab[22176]++
//line /usr/local/go/src/crypto/tls/conn.go:1562
		// _ = "end of CoverTab[22176]"
//line /usr/local/go/src/crypto/tls/conn.go:1562
	}
//line /usr/local/go/src/crypto/tls/conn.go:1562
	// _ = "end of CoverTab[22171]"
//line /usr/local/go/src/crypto/tls/conn.go:1562
	_go_fuzz_dep_.CoverTab[22172]++
							if !c.isHandshakeComplete.Load() {
//line /usr/local/go/src/crypto/tls/conn.go:1563
		_go_fuzz_dep_.CoverTab[22177]++
								return errors.New("tls: handshake has not yet been performed")
//line /usr/local/go/src/crypto/tls/conn.go:1564
		// _ = "end of CoverTab[22177]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1565
		_go_fuzz_dep_.CoverTab[22178]++
//line /usr/local/go/src/crypto/tls/conn.go:1565
		// _ = "end of CoverTab[22178]"
//line /usr/local/go/src/crypto/tls/conn.go:1565
	}
//line /usr/local/go/src/crypto/tls/conn.go:1565
	// _ = "end of CoverTab[22172]"
//line /usr/local/go/src/crypto/tls/conn.go:1565
	_go_fuzz_dep_.CoverTab[22173]++
							if len(c.verifiedChains) == 0 {
//line /usr/local/go/src/crypto/tls/conn.go:1566
		_go_fuzz_dep_.CoverTab[22179]++
								return errors.New("tls: handshake did not verify certificate chain")
//line /usr/local/go/src/crypto/tls/conn.go:1567
		// _ = "end of CoverTab[22179]"
	} else {
//line /usr/local/go/src/crypto/tls/conn.go:1568
		_go_fuzz_dep_.CoverTab[22180]++
//line /usr/local/go/src/crypto/tls/conn.go:1568
		// _ = "end of CoverTab[22180]"
//line /usr/local/go/src/crypto/tls/conn.go:1568
	}
//line /usr/local/go/src/crypto/tls/conn.go:1568
	// _ = "end of CoverTab[22173]"
//line /usr/local/go/src/crypto/tls/conn.go:1568
	_go_fuzz_dep_.CoverTab[22174]++
							return c.peerCertificates[0].VerifyHostname(host)
//line /usr/local/go/src/crypto/tls/conn.go:1569
	// _ = "end of CoverTab[22174]"
}

//line /usr/local/go/src/crypto/tls/conn.go:1570
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/conn.go:1570
var _ = _go_fuzz_dep_.CoverTab
