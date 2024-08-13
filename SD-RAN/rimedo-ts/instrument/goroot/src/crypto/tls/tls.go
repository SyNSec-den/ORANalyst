// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/crypto/tls/tls.go:5
// Package tls partially implements TLS 1.2, as specified in RFC 5246,
//line /usr/local/go/src/crypto/tls/tls.go:5
// and TLS 1.3, as specified in RFC 8446.
//line /usr/local/go/src/crypto/tls/tls.go:7
package tls

//line /usr/local/go/src/crypto/tls/tls.go:7
import (
//line /usr/local/go/src/crypto/tls/tls.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/tls/tls.go:7
)
//line /usr/local/go/src/crypto/tls/tls.go:7
import (
//line /usr/local/go/src/crypto/tls/tls.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/tls/tls.go:7
)

//line /usr/local/go/src/crypto/tls/tls.go:14
import (
	"bytes"
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
)

// Server returns a new TLS server side connection
//line /usr/local/go/src/crypto/tls/tls.go:30
// using conn as the underlying transport.
//line /usr/local/go/src/crypto/tls/tls.go:30
// The configuration config must be non-nil and must include
//line /usr/local/go/src/crypto/tls/tls.go:30
// at least one certificate or else set GetCertificate.
//line /usr/local/go/src/crypto/tls/tls.go:34
func Server(conn net.Conn, config *Config) *Conn {
//line /usr/local/go/src/crypto/tls/tls.go:34
	_go_fuzz_dep_.CoverTab[25059]++
						c := &Conn{
		conn:	conn,
		config:	config,
	}
						c.handshakeFn = c.serverHandshake
						return c
//line /usr/local/go/src/crypto/tls/tls.go:40
	// _ = "end of CoverTab[25059]"
}

// Client returns a new TLS client side connection
//line /usr/local/go/src/crypto/tls/tls.go:43
// using conn as the underlying transport.
//line /usr/local/go/src/crypto/tls/tls.go:43
// The config cannot be nil: users must set either ServerName or
//line /usr/local/go/src/crypto/tls/tls.go:43
// InsecureSkipVerify in the config.
//line /usr/local/go/src/crypto/tls/tls.go:47
func Client(conn net.Conn, config *Config) *Conn {
//line /usr/local/go/src/crypto/tls/tls.go:47
	_go_fuzz_dep_.CoverTab[25060]++
						c := &Conn{
		conn:		conn,
		config:		config,
		isClient:	true,
	}
						c.handshakeFn = c.clientHandshake
						return c
//line /usr/local/go/src/crypto/tls/tls.go:54
	// _ = "end of CoverTab[25060]"
}

// A listener implements a network listener (net.Listener) for TLS connections.
type listener struct {
	net.Listener
	config	*Config
}

// Accept waits for and returns the next incoming TLS connection.
//line /usr/local/go/src/crypto/tls/tls.go:63
// The returned connection is of type *Conn.
//line /usr/local/go/src/crypto/tls/tls.go:65
func (l *listener) Accept() (net.Conn, error) {
//line /usr/local/go/src/crypto/tls/tls.go:65
	_go_fuzz_dep_.CoverTab[25061]++
						c, err := l.Listener.Accept()
						if err != nil {
//line /usr/local/go/src/crypto/tls/tls.go:67
		_go_fuzz_dep_.CoverTab[25063]++
							return nil, err
//line /usr/local/go/src/crypto/tls/tls.go:68
		// _ = "end of CoverTab[25063]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:69
		_go_fuzz_dep_.CoverTab[25064]++
//line /usr/local/go/src/crypto/tls/tls.go:69
		// _ = "end of CoverTab[25064]"
//line /usr/local/go/src/crypto/tls/tls.go:69
	}
//line /usr/local/go/src/crypto/tls/tls.go:69
	// _ = "end of CoverTab[25061]"
//line /usr/local/go/src/crypto/tls/tls.go:69
	_go_fuzz_dep_.CoverTab[25062]++
						return Server(c, l.config), nil
//line /usr/local/go/src/crypto/tls/tls.go:70
	// _ = "end of CoverTab[25062]"
}

// NewListener creates a Listener which accepts connections from an inner
//line /usr/local/go/src/crypto/tls/tls.go:73
// Listener and wraps each connection with Server.
//line /usr/local/go/src/crypto/tls/tls.go:73
// The configuration config must be non-nil and must include
//line /usr/local/go/src/crypto/tls/tls.go:73
// at least one certificate or else set GetCertificate.
//line /usr/local/go/src/crypto/tls/tls.go:77
func NewListener(inner net.Listener, config *Config) net.Listener {
//line /usr/local/go/src/crypto/tls/tls.go:77
	_go_fuzz_dep_.CoverTab[25065]++
						l := new(listener)
						l.Listener = inner
						l.config = config
						return l
//line /usr/local/go/src/crypto/tls/tls.go:81
	// _ = "end of CoverTab[25065]"
}

// Listen creates a TLS listener accepting connections on the
//line /usr/local/go/src/crypto/tls/tls.go:84
// given network address using net.Listen.
//line /usr/local/go/src/crypto/tls/tls.go:84
// The configuration config must be non-nil and must include
//line /usr/local/go/src/crypto/tls/tls.go:84
// at least one certificate or else set GetCertificate.
//line /usr/local/go/src/crypto/tls/tls.go:88
func Listen(network, laddr string, config *Config) (net.Listener, error) {
//line /usr/local/go/src/crypto/tls/tls.go:88
	_go_fuzz_dep_.CoverTab[25066]++
						if config == nil || func() bool {
//line /usr/local/go/src/crypto/tls/tls.go:89
		_go_fuzz_dep_.CoverTab[25069]++
//line /usr/local/go/src/crypto/tls/tls.go:89
		return len(config.Certificates) == 0 && func() bool {
//line /usr/local/go/src/crypto/tls/tls.go:89
			_go_fuzz_dep_.CoverTab[25070]++
//line /usr/local/go/src/crypto/tls/tls.go:89
			return config.GetCertificate == nil
								// _ = "end of CoverTab[25070]"
//line /usr/local/go/src/crypto/tls/tls.go:90
		}() && func() bool {
//line /usr/local/go/src/crypto/tls/tls.go:90
			_go_fuzz_dep_.CoverTab[25071]++
//line /usr/local/go/src/crypto/tls/tls.go:90
			return config.GetConfigForClient == nil
//line /usr/local/go/src/crypto/tls/tls.go:90
			// _ = "end of CoverTab[25071]"
//line /usr/local/go/src/crypto/tls/tls.go:90
		}()
//line /usr/local/go/src/crypto/tls/tls.go:90
		// _ = "end of CoverTab[25069]"
//line /usr/local/go/src/crypto/tls/tls.go:90
	}() {
//line /usr/local/go/src/crypto/tls/tls.go:90
		_go_fuzz_dep_.CoverTab[25072]++
							return nil, errors.New("tls: neither Certificates, GetCertificate, nor GetConfigForClient set in Config")
//line /usr/local/go/src/crypto/tls/tls.go:91
		// _ = "end of CoverTab[25072]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:92
		_go_fuzz_dep_.CoverTab[25073]++
//line /usr/local/go/src/crypto/tls/tls.go:92
		// _ = "end of CoverTab[25073]"
//line /usr/local/go/src/crypto/tls/tls.go:92
	}
//line /usr/local/go/src/crypto/tls/tls.go:92
	// _ = "end of CoverTab[25066]"
//line /usr/local/go/src/crypto/tls/tls.go:92
	_go_fuzz_dep_.CoverTab[25067]++
						l, err := net.Listen(network, laddr)
						if err != nil {
//line /usr/local/go/src/crypto/tls/tls.go:94
		_go_fuzz_dep_.CoverTab[25074]++
							return nil, err
//line /usr/local/go/src/crypto/tls/tls.go:95
		// _ = "end of CoverTab[25074]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:96
		_go_fuzz_dep_.CoverTab[25075]++
//line /usr/local/go/src/crypto/tls/tls.go:96
		// _ = "end of CoverTab[25075]"
//line /usr/local/go/src/crypto/tls/tls.go:96
	}
//line /usr/local/go/src/crypto/tls/tls.go:96
	// _ = "end of CoverTab[25067]"
//line /usr/local/go/src/crypto/tls/tls.go:96
	_go_fuzz_dep_.CoverTab[25068]++
						return NewListener(l, config), nil
//line /usr/local/go/src/crypto/tls/tls.go:97
	// _ = "end of CoverTab[25068]"
}

type timeoutError struct{}

func (timeoutError) Error() string {
//line /usr/local/go/src/crypto/tls/tls.go:102
	_go_fuzz_dep_.CoverTab[25076]++
//line /usr/local/go/src/crypto/tls/tls.go:102
	return "tls: DialWithDialer timed out"
//line /usr/local/go/src/crypto/tls/tls.go:102
	// _ = "end of CoverTab[25076]"
//line /usr/local/go/src/crypto/tls/tls.go:102
}
func (timeoutError) Timeout() bool {
//line /usr/local/go/src/crypto/tls/tls.go:103
	_go_fuzz_dep_.CoverTab[25077]++
//line /usr/local/go/src/crypto/tls/tls.go:103
	return true
//line /usr/local/go/src/crypto/tls/tls.go:103
	// _ = "end of CoverTab[25077]"
//line /usr/local/go/src/crypto/tls/tls.go:103
}
func (timeoutError) Temporary() bool {
//line /usr/local/go/src/crypto/tls/tls.go:104
	_go_fuzz_dep_.CoverTab[25078]++
//line /usr/local/go/src/crypto/tls/tls.go:104
	return true
//line /usr/local/go/src/crypto/tls/tls.go:104
	// _ = "end of CoverTab[25078]"
//line /usr/local/go/src/crypto/tls/tls.go:104
}

// DialWithDialer connects to the given network address using dialer.Dial and
//line /usr/local/go/src/crypto/tls/tls.go:106
// then initiates a TLS handshake, returning the resulting TLS connection. Any
//line /usr/local/go/src/crypto/tls/tls.go:106
// timeout or deadline given in the dialer apply to connection and TLS
//line /usr/local/go/src/crypto/tls/tls.go:106
// handshake as a whole.
//line /usr/local/go/src/crypto/tls/tls.go:106
//
//line /usr/local/go/src/crypto/tls/tls.go:106
// DialWithDialer interprets a nil configuration as equivalent to the zero
//line /usr/local/go/src/crypto/tls/tls.go:106
// configuration; see the documentation of Config for the defaults.
//line /usr/local/go/src/crypto/tls/tls.go:106
//
//line /usr/local/go/src/crypto/tls/tls.go:106
// DialWithDialer uses context.Background internally; to specify the context,
//line /usr/local/go/src/crypto/tls/tls.go:106
// use Dialer.DialContext with NetDialer set to the desired dialer.
//line /usr/local/go/src/crypto/tls/tls.go:116
func DialWithDialer(dialer *net.Dialer, network, addr string, config *Config) (*Conn, error) {
//line /usr/local/go/src/crypto/tls/tls.go:116
	_go_fuzz_dep_.CoverTab[25079]++
						return dial(context.Background(), dialer, network, addr, config)
//line /usr/local/go/src/crypto/tls/tls.go:117
	// _ = "end of CoverTab[25079]"
}

func dial(ctx context.Context, netDialer *net.Dialer, network, addr string, config *Config) (*Conn, error) {
//line /usr/local/go/src/crypto/tls/tls.go:120
	_go_fuzz_dep_.CoverTab[25080]++
						if netDialer.Timeout != 0 {
//line /usr/local/go/src/crypto/tls/tls.go:121
		_go_fuzz_dep_.CoverTab[25088]++
							var cancel context.CancelFunc
							ctx, cancel = context.WithTimeout(ctx, netDialer.Timeout)
							defer cancel()
//line /usr/local/go/src/crypto/tls/tls.go:124
		// _ = "end of CoverTab[25088]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:125
		_go_fuzz_dep_.CoverTab[25089]++
//line /usr/local/go/src/crypto/tls/tls.go:125
		// _ = "end of CoverTab[25089]"
//line /usr/local/go/src/crypto/tls/tls.go:125
	}
//line /usr/local/go/src/crypto/tls/tls.go:125
	// _ = "end of CoverTab[25080]"
//line /usr/local/go/src/crypto/tls/tls.go:125
	_go_fuzz_dep_.CoverTab[25081]++

						if !netDialer.Deadline.IsZero() {
//line /usr/local/go/src/crypto/tls/tls.go:127
		_go_fuzz_dep_.CoverTab[25090]++
							var cancel context.CancelFunc
							ctx, cancel = context.WithDeadline(ctx, netDialer.Deadline)
							defer cancel()
//line /usr/local/go/src/crypto/tls/tls.go:130
		// _ = "end of CoverTab[25090]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:131
		_go_fuzz_dep_.CoverTab[25091]++
//line /usr/local/go/src/crypto/tls/tls.go:131
		// _ = "end of CoverTab[25091]"
//line /usr/local/go/src/crypto/tls/tls.go:131
	}
//line /usr/local/go/src/crypto/tls/tls.go:131
	// _ = "end of CoverTab[25081]"
//line /usr/local/go/src/crypto/tls/tls.go:131
	_go_fuzz_dep_.CoverTab[25082]++

						rawConn, err := netDialer.DialContext(ctx, network, addr)
						if err != nil {
//line /usr/local/go/src/crypto/tls/tls.go:134
		_go_fuzz_dep_.CoverTab[25092]++
							return nil, err
//line /usr/local/go/src/crypto/tls/tls.go:135
		// _ = "end of CoverTab[25092]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:136
		_go_fuzz_dep_.CoverTab[25093]++
//line /usr/local/go/src/crypto/tls/tls.go:136
		// _ = "end of CoverTab[25093]"
//line /usr/local/go/src/crypto/tls/tls.go:136
	}
//line /usr/local/go/src/crypto/tls/tls.go:136
	// _ = "end of CoverTab[25082]"
//line /usr/local/go/src/crypto/tls/tls.go:136
	_go_fuzz_dep_.CoverTab[25083]++

						colonPos := strings.LastIndex(addr, ":")
						if colonPos == -1 {
//line /usr/local/go/src/crypto/tls/tls.go:139
		_go_fuzz_dep_.CoverTab[25094]++
							colonPos = len(addr)
//line /usr/local/go/src/crypto/tls/tls.go:140
		// _ = "end of CoverTab[25094]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:141
		_go_fuzz_dep_.CoverTab[25095]++
//line /usr/local/go/src/crypto/tls/tls.go:141
		// _ = "end of CoverTab[25095]"
//line /usr/local/go/src/crypto/tls/tls.go:141
	}
//line /usr/local/go/src/crypto/tls/tls.go:141
	// _ = "end of CoverTab[25083]"
//line /usr/local/go/src/crypto/tls/tls.go:141
	_go_fuzz_dep_.CoverTab[25084]++
						hostname := addr[:colonPos]

						if config == nil {
//line /usr/local/go/src/crypto/tls/tls.go:144
		_go_fuzz_dep_.CoverTab[25096]++
							config = defaultConfig()
//line /usr/local/go/src/crypto/tls/tls.go:145
		// _ = "end of CoverTab[25096]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:146
		_go_fuzz_dep_.CoverTab[25097]++
//line /usr/local/go/src/crypto/tls/tls.go:146
		// _ = "end of CoverTab[25097]"
//line /usr/local/go/src/crypto/tls/tls.go:146
	}
//line /usr/local/go/src/crypto/tls/tls.go:146
	// _ = "end of CoverTab[25084]"
//line /usr/local/go/src/crypto/tls/tls.go:146
	_go_fuzz_dep_.CoverTab[25085]++

//line /usr/local/go/src/crypto/tls/tls.go:149
	if config.ServerName == "" {
//line /usr/local/go/src/crypto/tls/tls.go:149
		_go_fuzz_dep_.CoverTab[25098]++

							c := config.Clone()
							c.ServerName = hostname
							config = c
//line /usr/local/go/src/crypto/tls/tls.go:153
		// _ = "end of CoverTab[25098]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:154
		_go_fuzz_dep_.CoverTab[25099]++
//line /usr/local/go/src/crypto/tls/tls.go:154
		// _ = "end of CoverTab[25099]"
//line /usr/local/go/src/crypto/tls/tls.go:154
	}
//line /usr/local/go/src/crypto/tls/tls.go:154
	// _ = "end of CoverTab[25085]"
//line /usr/local/go/src/crypto/tls/tls.go:154
	_go_fuzz_dep_.CoverTab[25086]++

						conn := Client(rawConn, config)
						if err := conn.HandshakeContext(ctx); err != nil {
//line /usr/local/go/src/crypto/tls/tls.go:157
		_go_fuzz_dep_.CoverTab[25100]++
							rawConn.Close()
							return nil, err
//line /usr/local/go/src/crypto/tls/tls.go:159
		// _ = "end of CoverTab[25100]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:160
		_go_fuzz_dep_.CoverTab[25101]++
//line /usr/local/go/src/crypto/tls/tls.go:160
		// _ = "end of CoverTab[25101]"
//line /usr/local/go/src/crypto/tls/tls.go:160
	}
//line /usr/local/go/src/crypto/tls/tls.go:160
	// _ = "end of CoverTab[25086]"
//line /usr/local/go/src/crypto/tls/tls.go:160
	_go_fuzz_dep_.CoverTab[25087]++
						return conn, nil
//line /usr/local/go/src/crypto/tls/tls.go:161
	// _ = "end of CoverTab[25087]"
}

// Dial connects to the given network address using net.Dial
//line /usr/local/go/src/crypto/tls/tls.go:164
// and then initiates a TLS handshake, returning the resulting
//line /usr/local/go/src/crypto/tls/tls.go:164
// TLS connection.
//line /usr/local/go/src/crypto/tls/tls.go:164
// Dial interprets a nil configuration as equivalent to
//line /usr/local/go/src/crypto/tls/tls.go:164
// the zero configuration; see the documentation of Config
//line /usr/local/go/src/crypto/tls/tls.go:164
// for the defaults.
//line /usr/local/go/src/crypto/tls/tls.go:170
func Dial(network, addr string, config *Config) (*Conn, error) {
//line /usr/local/go/src/crypto/tls/tls.go:170
	_go_fuzz_dep_.CoverTab[25102]++
						return DialWithDialer(new(net.Dialer), network, addr, config)
//line /usr/local/go/src/crypto/tls/tls.go:171
	// _ = "end of CoverTab[25102]"
}

// Dialer dials TLS connections given a configuration and a Dialer for the
//line /usr/local/go/src/crypto/tls/tls.go:174
// underlying connection.
//line /usr/local/go/src/crypto/tls/tls.go:176
type Dialer struct {
	// NetDialer is the optional dialer to use for the TLS connections'
	// underlying TCP connections.
	// A nil NetDialer is equivalent to the net.Dialer zero value.
	NetDialer	*net.Dialer

	// Config is the TLS configuration to use for new connections.
	// A nil configuration is equivalent to the zero
	// configuration; see the documentation of Config for the
	// defaults.
	Config	*Config
}

// Dial connects to the given network address and initiates a TLS
//line /usr/local/go/src/crypto/tls/tls.go:189
// handshake, returning the resulting TLS connection.
//line /usr/local/go/src/crypto/tls/tls.go:189
//
//line /usr/local/go/src/crypto/tls/tls.go:189
// The returned Conn, if any, will always be of type *Conn.
//line /usr/local/go/src/crypto/tls/tls.go:189
//
//line /usr/local/go/src/crypto/tls/tls.go:189
// Dial uses context.Background internally; to specify the context,
//line /usr/local/go/src/crypto/tls/tls.go:189
// use DialContext.
//line /usr/local/go/src/crypto/tls/tls.go:196
func (d *Dialer) Dial(network, addr string) (net.Conn, error) {
//line /usr/local/go/src/crypto/tls/tls.go:196
	_go_fuzz_dep_.CoverTab[25103]++
						return d.DialContext(context.Background(), network, addr)
//line /usr/local/go/src/crypto/tls/tls.go:197
	// _ = "end of CoverTab[25103]"
}

func (d *Dialer) netDialer() *net.Dialer {
//line /usr/local/go/src/crypto/tls/tls.go:200
	_go_fuzz_dep_.CoverTab[25104]++
						if d.NetDialer != nil {
//line /usr/local/go/src/crypto/tls/tls.go:201
		_go_fuzz_dep_.CoverTab[25106]++
							return d.NetDialer
//line /usr/local/go/src/crypto/tls/tls.go:202
		// _ = "end of CoverTab[25106]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:203
		_go_fuzz_dep_.CoverTab[25107]++
//line /usr/local/go/src/crypto/tls/tls.go:203
		// _ = "end of CoverTab[25107]"
//line /usr/local/go/src/crypto/tls/tls.go:203
	}
//line /usr/local/go/src/crypto/tls/tls.go:203
	// _ = "end of CoverTab[25104]"
//line /usr/local/go/src/crypto/tls/tls.go:203
	_go_fuzz_dep_.CoverTab[25105]++
						return new(net.Dialer)
//line /usr/local/go/src/crypto/tls/tls.go:204
	// _ = "end of CoverTab[25105]"
}

// DialContext connects to the given network address and initiates a TLS
//line /usr/local/go/src/crypto/tls/tls.go:207
// handshake, returning the resulting TLS connection.
//line /usr/local/go/src/crypto/tls/tls.go:207
//
//line /usr/local/go/src/crypto/tls/tls.go:207
// The provided Context must be non-nil. If the context expires before
//line /usr/local/go/src/crypto/tls/tls.go:207
// the connection is complete, an error is returned. Once successfully
//line /usr/local/go/src/crypto/tls/tls.go:207
// connected, any expiration of the context will not affect the
//line /usr/local/go/src/crypto/tls/tls.go:207
// connection.
//line /usr/local/go/src/crypto/tls/tls.go:207
//
//line /usr/local/go/src/crypto/tls/tls.go:207
// The returned Conn, if any, will always be of type *Conn.
//line /usr/local/go/src/crypto/tls/tls.go:216
func (d *Dialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
//line /usr/local/go/src/crypto/tls/tls.go:216
	_go_fuzz_dep_.CoverTab[25108]++
						c, err := dial(ctx, d.netDialer(), network, addr, d.Config)
						if err != nil {
//line /usr/local/go/src/crypto/tls/tls.go:218
		_go_fuzz_dep_.CoverTab[25110]++

							return nil, err
//line /usr/local/go/src/crypto/tls/tls.go:220
		// _ = "end of CoverTab[25110]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:221
		_go_fuzz_dep_.CoverTab[25111]++
//line /usr/local/go/src/crypto/tls/tls.go:221
		// _ = "end of CoverTab[25111]"
//line /usr/local/go/src/crypto/tls/tls.go:221
	}
//line /usr/local/go/src/crypto/tls/tls.go:221
	// _ = "end of CoverTab[25108]"
//line /usr/local/go/src/crypto/tls/tls.go:221
	_go_fuzz_dep_.CoverTab[25109]++
						return c, nil
//line /usr/local/go/src/crypto/tls/tls.go:222
	// _ = "end of CoverTab[25109]"
}

// LoadX509KeyPair reads and parses a public/private key pair from a pair
//line /usr/local/go/src/crypto/tls/tls.go:225
// of files. The files must contain PEM encoded data. The certificate file
//line /usr/local/go/src/crypto/tls/tls.go:225
// may contain intermediate certificates following the leaf certificate to
//line /usr/local/go/src/crypto/tls/tls.go:225
// form a certificate chain. On successful return, Certificate.Leaf will
//line /usr/local/go/src/crypto/tls/tls.go:225
// be nil because the parsed form of the certificate is not retained.
//line /usr/local/go/src/crypto/tls/tls.go:230
func LoadX509KeyPair(certFile, keyFile string) (Certificate, error) {
//line /usr/local/go/src/crypto/tls/tls.go:230
	_go_fuzz_dep_.CoverTab[25112]++
						certPEMBlock, err := os.ReadFile(certFile)
						if err != nil {
//line /usr/local/go/src/crypto/tls/tls.go:232
		_go_fuzz_dep_.CoverTab[25115]++
							return Certificate{}, err
//line /usr/local/go/src/crypto/tls/tls.go:233
		// _ = "end of CoverTab[25115]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:234
		_go_fuzz_dep_.CoverTab[25116]++
//line /usr/local/go/src/crypto/tls/tls.go:234
		// _ = "end of CoverTab[25116]"
//line /usr/local/go/src/crypto/tls/tls.go:234
	}
//line /usr/local/go/src/crypto/tls/tls.go:234
	// _ = "end of CoverTab[25112]"
//line /usr/local/go/src/crypto/tls/tls.go:234
	_go_fuzz_dep_.CoverTab[25113]++
						keyPEMBlock, err := os.ReadFile(keyFile)
						if err != nil {
//line /usr/local/go/src/crypto/tls/tls.go:236
		_go_fuzz_dep_.CoverTab[25117]++
							return Certificate{}, err
//line /usr/local/go/src/crypto/tls/tls.go:237
		// _ = "end of CoverTab[25117]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:238
		_go_fuzz_dep_.CoverTab[25118]++
//line /usr/local/go/src/crypto/tls/tls.go:238
		// _ = "end of CoverTab[25118]"
//line /usr/local/go/src/crypto/tls/tls.go:238
	}
//line /usr/local/go/src/crypto/tls/tls.go:238
	// _ = "end of CoverTab[25113]"
//line /usr/local/go/src/crypto/tls/tls.go:238
	_go_fuzz_dep_.CoverTab[25114]++
						return X509KeyPair(certPEMBlock, keyPEMBlock)
//line /usr/local/go/src/crypto/tls/tls.go:239
	// _ = "end of CoverTab[25114]"
}

// X509KeyPair parses a public/private key pair from a pair of
//line /usr/local/go/src/crypto/tls/tls.go:242
// PEM encoded data. On successful return, Certificate.Leaf will be nil because
//line /usr/local/go/src/crypto/tls/tls.go:242
// the parsed form of the certificate is not retained.
//line /usr/local/go/src/crypto/tls/tls.go:245
func X509KeyPair(certPEMBlock, keyPEMBlock []byte) (Certificate, error) {
//line /usr/local/go/src/crypto/tls/tls.go:245
	_go_fuzz_dep_.CoverTab[25119]++
						fail := func(err error) (Certificate, error) {
//line /usr/local/go/src/crypto/tls/tls.go:246
		_go_fuzz_dep_.CoverTab[25127]++
//line /usr/local/go/src/crypto/tls/tls.go:246
		return Certificate{}, err
//line /usr/local/go/src/crypto/tls/tls.go:246
		// _ = "end of CoverTab[25127]"
//line /usr/local/go/src/crypto/tls/tls.go:246
	}
//line /usr/local/go/src/crypto/tls/tls.go:246
	// _ = "end of CoverTab[25119]"
//line /usr/local/go/src/crypto/tls/tls.go:246
	_go_fuzz_dep_.CoverTab[25120]++

						var cert Certificate
						var skippedBlockTypes []string
						for {
//line /usr/local/go/src/crypto/tls/tls.go:250
		_go_fuzz_dep_.CoverTab[25128]++
							var certDERBlock *pem.Block
							certDERBlock, certPEMBlock = pem.Decode(certPEMBlock)
							if certDERBlock == nil {
//line /usr/local/go/src/crypto/tls/tls.go:253
			_go_fuzz_dep_.CoverTab[25130]++
								break
//line /usr/local/go/src/crypto/tls/tls.go:254
			// _ = "end of CoverTab[25130]"
		} else {
//line /usr/local/go/src/crypto/tls/tls.go:255
			_go_fuzz_dep_.CoverTab[25131]++
//line /usr/local/go/src/crypto/tls/tls.go:255
			// _ = "end of CoverTab[25131]"
//line /usr/local/go/src/crypto/tls/tls.go:255
		}
//line /usr/local/go/src/crypto/tls/tls.go:255
		// _ = "end of CoverTab[25128]"
//line /usr/local/go/src/crypto/tls/tls.go:255
		_go_fuzz_dep_.CoverTab[25129]++
							if certDERBlock.Type == "CERTIFICATE" {
//line /usr/local/go/src/crypto/tls/tls.go:256
			_go_fuzz_dep_.CoverTab[25132]++
								cert.Certificate = append(cert.Certificate, certDERBlock.Bytes)
//line /usr/local/go/src/crypto/tls/tls.go:257
			// _ = "end of CoverTab[25132]"
		} else {
//line /usr/local/go/src/crypto/tls/tls.go:258
			_go_fuzz_dep_.CoverTab[25133]++
								skippedBlockTypes = append(skippedBlockTypes, certDERBlock.Type)
//line /usr/local/go/src/crypto/tls/tls.go:259
			// _ = "end of CoverTab[25133]"
		}
//line /usr/local/go/src/crypto/tls/tls.go:260
		// _ = "end of CoverTab[25129]"
	}
//line /usr/local/go/src/crypto/tls/tls.go:261
	// _ = "end of CoverTab[25120]"
//line /usr/local/go/src/crypto/tls/tls.go:261
	_go_fuzz_dep_.CoverTab[25121]++

						if len(cert.Certificate) == 0 {
//line /usr/local/go/src/crypto/tls/tls.go:263
		_go_fuzz_dep_.CoverTab[25134]++
							if len(skippedBlockTypes) == 0 {
//line /usr/local/go/src/crypto/tls/tls.go:264
			_go_fuzz_dep_.CoverTab[25137]++
								return fail(errors.New("tls: failed to find any PEM data in certificate input"))
//line /usr/local/go/src/crypto/tls/tls.go:265
			// _ = "end of CoverTab[25137]"
		} else {
//line /usr/local/go/src/crypto/tls/tls.go:266
			_go_fuzz_dep_.CoverTab[25138]++
//line /usr/local/go/src/crypto/tls/tls.go:266
			// _ = "end of CoverTab[25138]"
//line /usr/local/go/src/crypto/tls/tls.go:266
		}
//line /usr/local/go/src/crypto/tls/tls.go:266
		// _ = "end of CoverTab[25134]"
//line /usr/local/go/src/crypto/tls/tls.go:266
		_go_fuzz_dep_.CoverTab[25135]++
							if len(skippedBlockTypes) == 1 && func() bool {
//line /usr/local/go/src/crypto/tls/tls.go:267
			_go_fuzz_dep_.CoverTab[25139]++
//line /usr/local/go/src/crypto/tls/tls.go:267
			return strings.HasSuffix(skippedBlockTypes[0], "PRIVATE KEY")
//line /usr/local/go/src/crypto/tls/tls.go:267
			// _ = "end of CoverTab[25139]"
//line /usr/local/go/src/crypto/tls/tls.go:267
		}() {
//line /usr/local/go/src/crypto/tls/tls.go:267
			_go_fuzz_dep_.CoverTab[25140]++
								return fail(errors.New("tls: failed to find certificate PEM data in certificate input, but did find a private key; PEM inputs may have been switched"))
//line /usr/local/go/src/crypto/tls/tls.go:268
			// _ = "end of CoverTab[25140]"
		} else {
//line /usr/local/go/src/crypto/tls/tls.go:269
			_go_fuzz_dep_.CoverTab[25141]++
//line /usr/local/go/src/crypto/tls/tls.go:269
			// _ = "end of CoverTab[25141]"
//line /usr/local/go/src/crypto/tls/tls.go:269
		}
//line /usr/local/go/src/crypto/tls/tls.go:269
		// _ = "end of CoverTab[25135]"
//line /usr/local/go/src/crypto/tls/tls.go:269
		_go_fuzz_dep_.CoverTab[25136]++
							return fail(fmt.Errorf("tls: failed to find \"CERTIFICATE\" PEM block in certificate input after skipping PEM blocks of the following types: %v", skippedBlockTypes))
//line /usr/local/go/src/crypto/tls/tls.go:270
		// _ = "end of CoverTab[25136]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:271
		_go_fuzz_dep_.CoverTab[25142]++
//line /usr/local/go/src/crypto/tls/tls.go:271
		// _ = "end of CoverTab[25142]"
//line /usr/local/go/src/crypto/tls/tls.go:271
	}
//line /usr/local/go/src/crypto/tls/tls.go:271
	// _ = "end of CoverTab[25121]"
//line /usr/local/go/src/crypto/tls/tls.go:271
	_go_fuzz_dep_.CoverTab[25122]++

						skippedBlockTypes = skippedBlockTypes[:0]
						var keyDERBlock *pem.Block
						for {
//line /usr/local/go/src/crypto/tls/tls.go:275
		_go_fuzz_dep_.CoverTab[25143]++
							keyDERBlock, keyPEMBlock = pem.Decode(keyPEMBlock)
							if keyDERBlock == nil {
//line /usr/local/go/src/crypto/tls/tls.go:277
			_go_fuzz_dep_.CoverTab[25146]++
								if len(skippedBlockTypes) == 0 {
//line /usr/local/go/src/crypto/tls/tls.go:278
				_go_fuzz_dep_.CoverTab[25149]++
									return fail(errors.New("tls: failed to find any PEM data in key input"))
//line /usr/local/go/src/crypto/tls/tls.go:279
				// _ = "end of CoverTab[25149]"
			} else {
//line /usr/local/go/src/crypto/tls/tls.go:280
				_go_fuzz_dep_.CoverTab[25150]++
//line /usr/local/go/src/crypto/tls/tls.go:280
				// _ = "end of CoverTab[25150]"
//line /usr/local/go/src/crypto/tls/tls.go:280
			}
//line /usr/local/go/src/crypto/tls/tls.go:280
			// _ = "end of CoverTab[25146]"
//line /usr/local/go/src/crypto/tls/tls.go:280
			_go_fuzz_dep_.CoverTab[25147]++
								if len(skippedBlockTypes) == 1 && func() bool {
//line /usr/local/go/src/crypto/tls/tls.go:281
				_go_fuzz_dep_.CoverTab[25151]++
//line /usr/local/go/src/crypto/tls/tls.go:281
				return skippedBlockTypes[0] == "CERTIFICATE"
//line /usr/local/go/src/crypto/tls/tls.go:281
				// _ = "end of CoverTab[25151]"
//line /usr/local/go/src/crypto/tls/tls.go:281
			}() {
//line /usr/local/go/src/crypto/tls/tls.go:281
				_go_fuzz_dep_.CoverTab[25152]++
									return fail(errors.New("tls: found a certificate rather than a key in the PEM for the private key"))
//line /usr/local/go/src/crypto/tls/tls.go:282
				// _ = "end of CoverTab[25152]"
			} else {
//line /usr/local/go/src/crypto/tls/tls.go:283
				_go_fuzz_dep_.CoverTab[25153]++
//line /usr/local/go/src/crypto/tls/tls.go:283
				// _ = "end of CoverTab[25153]"
//line /usr/local/go/src/crypto/tls/tls.go:283
			}
//line /usr/local/go/src/crypto/tls/tls.go:283
			// _ = "end of CoverTab[25147]"
//line /usr/local/go/src/crypto/tls/tls.go:283
			_go_fuzz_dep_.CoverTab[25148]++
								return fail(fmt.Errorf("tls: failed to find PEM block with type ending in \"PRIVATE KEY\" in key input after skipping PEM blocks of the following types: %v", skippedBlockTypes))
//line /usr/local/go/src/crypto/tls/tls.go:284
			// _ = "end of CoverTab[25148]"
		} else {
//line /usr/local/go/src/crypto/tls/tls.go:285
			_go_fuzz_dep_.CoverTab[25154]++
//line /usr/local/go/src/crypto/tls/tls.go:285
			// _ = "end of CoverTab[25154]"
//line /usr/local/go/src/crypto/tls/tls.go:285
		}
//line /usr/local/go/src/crypto/tls/tls.go:285
		// _ = "end of CoverTab[25143]"
//line /usr/local/go/src/crypto/tls/tls.go:285
		_go_fuzz_dep_.CoverTab[25144]++
							if keyDERBlock.Type == "PRIVATE KEY" || func() bool {
//line /usr/local/go/src/crypto/tls/tls.go:286
			_go_fuzz_dep_.CoverTab[25155]++
//line /usr/local/go/src/crypto/tls/tls.go:286
			return strings.HasSuffix(keyDERBlock.Type, " PRIVATE KEY")
//line /usr/local/go/src/crypto/tls/tls.go:286
			// _ = "end of CoverTab[25155]"
//line /usr/local/go/src/crypto/tls/tls.go:286
		}() {
//line /usr/local/go/src/crypto/tls/tls.go:286
			_go_fuzz_dep_.CoverTab[25156]++
								break
//line /usr/local/go/src/crypto/tls/tls.go:287
			// _ = "end of CoverTab[25156]"
		} else {
//line /usr/local/go/src/crypto/tls/tls.go:288
			_go_fuzz_dep_.CoverTab[25157]++
//line /usr/local/go/src/crypto/tls/tls.go:288
			// _ = "end of CoverTab[25157]"
//line /usr/local/go/src/crypto/tls/tls.go:288
		}
//line /usr/local/go/src/crypto/tls/tls.go:288
		// _ = "end of CoverTab[25144]"
//line /usr/local/go/src/crypto/tls/tls.go:288
		_go_fuzz_dep_.CoverTab[25145]++
							skippedBlockTypes = append(skippedBlockTypes, keyDERBlock.Type)
//line /usr/local/go/src/crypto/tls/tls.go:289
		// _ = "end of CoverTab[25145]"
	}
//line /usr/local/go/src/crypto/tls/tls.go:290
	// _ = "end of CoverTab[25122]"
//line /usr/local/go/src/crypto/tls/tls.go:290
	_go_fuzz_dep_.CoverTab[25123]++

//line /usr/local/go/src/crypto/tls/tls.go:294
	x509Cert, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
//line /usr/local/go/src/crypto/tls/tls.go:295
		_go_fuzz_dep_.CoverTab[25158]++
							return fail(err)
//line /usr/local/go/src/crypto/tls/tls.go:296
		// _ = "end of CoverTab[25158]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:297
		_go_fuzz_dep_.CoverTab[25159]++
//line /usr/local/go/src/crypto/tls/tls.go:297
		// _ = "end of CoverTab[25159]"
//line /usr/local/go/src/crypto/tls/tls.go:297
	}
//line /usr/local/go/src/crypto/tls/tls.go:297
	// _ = "end of CoverTab[25123]"
//line /usr/local/go/src/crypto/tls/tls.go:297
	_go_fuzz_dep_.CoverTab[25124]++

						cert.PrivateKey, err = parsePrivateKey(keyDERBlock.Bytes)
						if err != nil {
//line /usr/local/go/src/crypto/tls/tls.go:300
		_go_fuzz_dep_.CoverTab[25160]++
							return fail(err)
//line /usr/local/go/src/crypto/tls/tls.go:301
		// _ = "end of CoverTab[25160]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:302
		_go_fuzz_dep_.CoverTab[25161]++
//line /usr/local/go/src/crypto/tls/tls.go:302
		// _ = "end of CoverTab[25161]"
//line /usr/local/go/src/crypto/tls/tls.go:302
	}
//line /usr/local/go/src/crypto/tls/tls.go:302
	// _ = "end of CoverTab[25124]"
//line /usr/local/go/src/crypto/tls/tls.go:302
	_go_fuzz_dep_.CoverTab[25125]++

						switch pub := x509Cert.PublicKey.(type) {
	case *rsa.PublicKey:
//line /usr/local/go/src/crypto/tls/tls.go:305
		_go_fuzz_dep_.CoverTab[25162]++
							priv, ok := cert.PrivateKey.(*rsa.PrivateKey)
							if !ok {
//line /usr/local/go/src/crypto/tls/tls.go:307
			_go_fuzz_dep_.CoverTab[25169]++
								return fail(errors.New("tls: private key type does not match public key type"))
//line /usr/local/go/src/crypto/tls/tls.go:308
			// _ = "end of CoverTab[25169]"
		} else {
//line /usr/local/go/src/crypto/tls/tls.go:309
			_go_fuzz_dep_.CoverTab[25170]++
//line /usr/local/go/src/crypto/tls/tls.go:309
			// _ = "end of CoverTab[25170]"
//line /usr/local/go/src/crypto/tls/tls.go:309
		}
//line /usr/local/go/src/crypto/tls/tls.go:309
		// _ = "end of CoverTab[25162]"
//line /usr/local/go/src/crypto/tls/tls.go:309
		_go_fuzz_dep_.CoverTab[25163]++
							if pub.N.Cmp(priv.N) != 0 {
//line /usr/local/go/src/crypto/tls/tls.go:310
			_go_fuzz_dep_.CoverTab[25171]++
								return fail(errors.New("tls: private key does not match public key"))
//line /usr/local/go/src/crypto/tls/tls.go:311
			// _ = "end of CoverTab[25171]"
		} else {
//line /usr/local/go/src/crypto/tls/tls.go:312
			_go_fuzz_dep_.CoverTab[25172]++
//line /usr/local/go/src/crypto/tls/tls.go:312
			// _ = "end of CoverTab[25172]"
//line /usr/local/go/src/crypto/tls/tls.go:312
		}
//line /usr/local/go/src/crypto/tls/tls.go:312
		// _ = "end of CoverTab[25163]"
	case *ecdsa.PublicKey:
//line /usr/local/go/src/crypto/tls/tls.go:313
		_go_fuzz_dep_.CoverTab[25164]++
							priv, ok := cert.PrivateKey.(*ecdsa.PrivateKey)
							if !ok {
//line /usr/local/go/src/crypto/tls/tls.go:315
			_go_fuzz_dep_.CoverTab[25173]++
								return fail(errors.New("tls: private key type does not match public key type"))
//line /usr/local/go/src/crypto/tls/tls.go:316
			// _ = "end of CoverTab[25173]"
		} else {
//line /usr/local/go/src/crypto/tls/tls.go:317
			_go_fuzz_dep_.CoverTab[25174]++
//line /usr/local/go/src/crypto/tls/tls.go:317
			// _ = "end of CoverTab[25174]"
//line /usr/local/go/src/crypto/tls/tls.go:317
		}
//line /usr/local/go/src/crypto/tls/tls.go:317
		// _ = "end of CoverTab[25164]"
//line /usr/local/go/src/crypto/tls/tls.go:317
		_go_fuzz_dep_.CoverTab[25165]++
							if pub.X.Cmp(priv.X) != 0 || func() bool {
//line /usr/local/go/src/crypto/tls/tls.go:318
			_go_fuzz_dep_.CoverTab[25175]++
//line /usr/local/go/src/crypto/tls/tls.go:318
			return pub.Y.Cmp(priv.Y) != 0
//line /usr/local/go/src/crypto/tls/tls.go:318
			// _ = "end of CoverTab[25175]"
//line /usr/local/go/src/crypto/tls/tls.go:318
		}() {
//line /usr/local/go/src/crypto/tls/tls.go:318
			_go_fuzz_dep_.CoverTab[25176]++
								return fail(errors.New("tls: private key does not match public key"))
//line /usr/local/go/src/crypto/tls/tls.go:319
			// _ = "end of CoverTab[25176]"
		} else {
//line /usr/local/go/src/crypto/tls/tls.go:320
			_go_fuzz_dep_.CoverTab[25177]++
//line /usr/local/go/src/crypto/tls/tls.go:320
			// _ = "end of CoverTab[25177]"
//line /usr/local/go/src/crypto/tls/tls.go:320
		}
//line /usr/local/go/src/crypto/tls/tls.go:320
		// _ = "end of CoverTab[25165]"
	case ed25519.PublicKey:
//line /usr/local/go/src/crypto/tls/tls.go:321
		_go_fuzz_dep_.CoverTab[25166]++
							priv, ok := cert.PrivateKey.(ed25519.PrivateKey)
							if !ok {
//line /usr/local/go/src/crypto/tls/tls.go:323
			_go_fuzz_dep_.CoverTab[25178]++
								return fail(errors.New("tls: private key type does not match public key type"))
//line /usr/local/go/src/crypto/tls/tls.go:324
			// _ = "end of CoverTab[25178]"
		} else {
//line /usr/local/go/src/crypto/tls/tls.go:325
			_go_fuzz_dep_.CoverTab[25179]++
//line /usr/local/go/src/crypto/tls/tls.go:325
			// _ = "end of CoverTab[25179]"
//line /usr/local/go/src/crypto/tls/tls.go:325
		}
//line /usr/local/go/src/crypto/tls/tls.go:325
		// _ = "end of CoverTab[25166]"
//line /usr/local/go/src/crypto/tls/tls.go:325
		_go_fuzz_dep_.CoverTab[25167]++
							if !bytes.Equal(priv.Public().(ed25519.PublicKey), pub) {
//line /usr/local/go/src/crypto/tls/tls.go:326
			_go_fuzz_dep_.CoverTab[25180]++
								return fail(errors.New("tls: private key does not match public key"))
//line /usr/local/go/src/crypto/tls/tls.go:327
			// _ = "end of CoverTab[25180]"
		} else {
//line /usr/local/go/src/crypto/tls/tls.go:328
			_go_fuzz_dep_.CoverTab[25181]++
//line /usr/local/go/src/crypto/tls/tls.go:328
			// _ = "end of CoverTab[25181]"
//line /usr/local/go/src/crypto/tls/tls.go:328
		}
//line /usr/local/go/src/crypto/tls/tls.go:328
		// _ = "end of CoverTab[25167]"
	default:
//line /usr/local/go/src/crypto/tls/tls.go:329
		_go_fuzz_dep_.CoverTab[25168]++
							return fail(errors.New("tls: unknown public key algorithm"))
//line /usr/local/go/src/crypto/tls/tls.go:330
		// _ = "end of CoverTab[25168]"
	}
//line /usr/local/go/src/crypto/tls/tls.go:331
	// _ = "end of CoverTab[25125]"
//line /usr/local/go/src/crypto/tls/tls.go:331
	_go_fuzz_dep_.CoverTab[25126]++

						return cert, nil
//line /usr/local/go/src/crypto/tls/tls.go:333
	// _ = "end of CoverTab[25126]"
}

// Attempt to parse the given private key DER block. OpenSSL 0.9.8 generates
//line /usr/local/go/src/crypto/tls/tls.go:336
// PKCS #1 private keys by default, while OpenSSL 1.0.0 generates PKCS #8 keys.
//line /usr/local/go/src/crypto/tls/tls.go:336
// OpenSSL ecparam generates SEC1 EC private keys for ECDSA. We try all three.
//line /usr/local/go/src/crypto/tls/tls.go:339
func parsePrivateKey(der []byte) (crypto.PrivateKey, error) {
//line /usr/local/go/src/crypto/tls/tls.go:339
	_go_fuzz_dep_.CoverTab[25182]++
						if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
//line /usr/local/go/src/crypto/tls/tls.go:340
		_go_fuzz_dep_.CoverTab[25186]++
							return key, nil
//line /usr/local/go/src/crypto/tls/tls.go:341
		// _ = "end of CoverTab[25186]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:342
		_go_fuzz_dep_.CoverTab[25187]++
//line /usr/local/go/src/crypto/tls/tls.go:342
		// _ = "end of CoverTab[25187]"
//line /usr/local/go/src/crypto/tls/tls.go:342
	}
//line /usr/local/go/src/crypto/tls/tls.go:342
	// _ = "end of CoverTab[25182]"
//line /usr/local/go/src/crypto/tls/tls.go:342
	_go_fuzz_dep_.CoverTab[25183]++
						if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
//line /usr/local/go/src/crypto/tls/tls.go:343
		_go_fuzz_dep_.CoverTab[25188]++
							switch key := key.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey, ed25519.PrivateKey:
//line /usr/local/go/src/crypto/tls/tls.go:345
			_go_fuzz_dep_.CoverTab[25189]++
								return key, nil
//line /usr/local/go/src/crypto/tls/tls.go:346
			// _ = "end of CoverTab[25189]"
		default:
//line /usr/local/go/src/crypto/tls/tls.go:347
			_go_fuzz_dep_.CoverTab[25190]++
								return nil, errors.New("tls: found unknown private key type in PKCS#8 wrapping")
//line /usr/local/go/src/crypto/tls/tls.go:348
			// _ = "end of CoverTab[25190]"
		}
//line /usr/local/go/src/crypto/tls/tls.go:349
		// _ = "end of CoverTab[25188]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:350
		_go_fuzz_dep_.CoverTab[25191]++
//line /usr/local/go/src/crypto/tls/tls.go:350
		// _ = "end of CoverTab[25191]"
//line /usr/local/go/src/crypto/tls/tls.go:350
	}
//line /usr/local/go/src/crypto/tls/tls.go:350
	// _ = "end of CoverTab[25183]"
//line /usr/local/go/src/crypto/tls/tls.go:350
	_go_fuzz_dep_.CoverTab[25184]++
						if key, err := x509.ParseECPrivateKey(der); err == nil {
//line /usr/local/go/src/crypto/tls/tls.go:351
		_go_fuzz_dep_.CoverTab[25192]++
							return key, nil
//line /usr/local/go/src/crypto/tls/tls.go:352
		// _ = "end of CoverTab[25192]"
	} else {
//line /usr/local/go/src/crypto/tls/tls.go:353
		_go_fuzz_dep_.CoverTab[25193]++
//line /usr/local/go/src/crypto/tls/tls.go:353
		// _ = "end of CoverTab[25193]"
//line /usr/local/go/src/crypto/tls/tls.go:353
	}
//line /usr/local/go/src/crypto/tls/tls.go:353
	// _ = "end of CoverTab[25184]"
//line /usr/local/go/src/crypto/tls/tls.go:353
	_go_fuzz_dep_.CoverTab[25185]++

						return nil, errors.New("tls: failed to parse private key")
//line /usr/local/go/src/crypto/tls/tls.go:355
	// _ = "end of CoverTab[25185]"
}

//line /usr/local/go/src/crypto/tls/tls.go:356
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/tls/tls.go:356
var _ = _go_fuzz_dep_.CoverTab
