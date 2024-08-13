// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Transport code.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:7
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:7
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:7
)

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/rand"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math"
	mathrand "math/rand"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/textproto"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/http/httpguts"
	"golang.org/x/net/http2/hpack"
	"golang.org/x/net/idna"
)

const (
	// transportDefaultConnFlow is how many connection-level flow control
	// tokens we give the server at start-up, past the default 64k.
	transportDefaultConnFlow	= 1 << 30

	// transportDefaultStreamFlow is how many stream-level flow
	// control tokens we announce to the peer, and how many bytes
	// we buffer per stream.
	transportDefaultStreamFlow	= 4 << 20

	defaultUserAgent	= "Go-http-client/2.0"

	// initialMaxConcurrentStreams is a connections maxConcurrentStreams until
	// it's received servers initial SETTINGS frame, which corresponds with the
	// spec's minimum recommended value.
	initialMaxConcurrentStreams	= 100

	// defaultMaxConcurrentStreams is a connections default maxConcurrentStreams
	// if the server doesn't include one in its initial SETTINGS frame.
	defaultMaxConcurrentStreams	= 1000
)

// Transport is an HTTP/2 Transport.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:62
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:62
// A Transport internally caches connections to servers. It is safe
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:62
// for concurrent use by multiple goroutines.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:66
type Transport struct {
	// DialTLSContext specifies an optional dial function with context for
	// creating TLS connections for requests.
	//
	// If DialTLSContext and DialTLS is nil, tls.Dial is used.
	//
	// If the returned net.Conn has a ConnectionState method like tls.Conn,
	// it will be used to set http.Response.TLS.
	DialTLSContext	func(ctx context.Context, network, addr string, cfg *tls.Config) (net.Conn, error)

	// DialTLS specifies an optional dial function for creating
	// TLS connections for requests.
	//
	// If DialTLSContext and DialTLS is nil, tls.Dial is used.
	//
	// Deprecated: Use DialTLSContext instead, which allows the transport
	// to cancel dials as soon as they are no longer needed.
	// If both are set, DialTLSContext takes priority.
	DialTLS	func(network, addr string, cfg *tls.Config) (net.Conn, error)

	// TLSClientConfig specifies the TLS configuration to use with
	// tls.Client. If nil, the default configuration is used.
	TLSClientConfig	*tls.Config

	// ConnPool optionally specifies an alternate connection pool to use.
	// If nil, the default is used.
	ConnPool	ClientConnPool

	// DisableCompression, if true, prevents the Transport from
	// requesting compression with an "Accept-Encoding: gzip"
	// request header when the Request contains no existing
	// Accept-Encoding value. If the Transport requests gzip on
	// its own and gets a gzipped response, it's transparently
	// decoded in the Response.Body. However, if the user
	// explicitly requested gzip it is not automatically
	// uncompressed.
	DisableCompression	bool

	// AllowHTTP, if true, permits HTTP/2 requests using the insecure,
	// plain-text "http" scheme. Note that this does not enable h2c support.
	AllowHTTP	bool

	// MaxHeaderListSize is the http2 SETTINGS_MAX_HEADER_LIST_SIZE to
	// send in the initial settings frame. It is how many bytes
	// of response headers are allowed. Unlike the http2 spec, zero here
	// means to use a default limit (currently 10MB). If you actually
	// want to advertise an unlimited value to the peer, Transport
	// interprets the highest possible value here (0xffffffff or 1<<32-1)
	// to mean no limit.
	MaxHeaderListSize	uint32

	// MaxReadFrameSize is the http2 SETTINGS_MAX_FRAME_SIZE to send in the
	// initial settings frame. It is the size in bytes of the largest frame
	// payload that the sender is willing to receive. If 0, no setting is
	// sent, and the value is provided by the peer, which should be 16384
	// according to the spec:
	// https://datatracker.ietf.org/doc/html/rfc7540#section-6.5.2.
	// Values are bounded in the range 16k to 16M.
	MaxReadFrameSize	uint32

	// MaxDecoderHeaderTableSize optionally specifies the http2
	// SETTINGS_HEADER_TABLE_SIZE to send in the initial settings frame. It
	// informs the remote endpoint of the maximum size of the header compression
	// table used to decode header blocks, in octets. If zero, the default value
	// of 4096 is used.
	MaxDecoderHeaderTableSize	uint32

	// MaxEncoderHeaderTableSize optionally specifies an upper limit for the
	// header compression table used for encoding request headers. Received
	// SETTINGS_HEADER_TABLE_SIZE settings are capped at this limit. If zero,
	// the default value of 4096 is used.
	MaxEncoderHeaderTableSize	uint32

	// StrictMaxConcurrentStreams controls whether the server's
	// SETTINGS_MAX_CONCURRENT_STREAMS should be respected
	// globally. If false, new TCP connections are created to the
	// server as needed to keep each under the per-connection
	// SETTINGS_MAX_CONCURRENT_STREAMS limit. If true, the
	// server's SETTINGS_MAX_CONCURRENT_STREAMS is interpreted as
	// a global limit and callers of RoundTrip block when needed,
	// waiting for their turn.
	StrictMaxConcurrentStreams	bool

	// ReadIdleTimeout is the timeout after which a health check using ping
	// frame will be carried out if no frame is received on the connection.
	// Note that a ping response will is considered a received frame, so if
	// there is no other traffic on the connection, the health check will
	// be performed every ReadIdleTimeout interval.
	// If zero, no health check is performed.
	ReadIdleTimeout	time.Duration

	// PingTimeout is the timeout after which the connection will be closed
	// if a response to Ping is not received.
	// Defaults to 15s.
	PingTimeout	time.Duration

	// WriteByteTimeout is the timeout after which the connection will be
	// closed no data can be written to it. The timeout begins when data is
	// available to write, and is extended whenever any bytes are written.
	WriteByteTimeout	time.Duration

	// CountError, if non-nil, is called on HTTP/2 transport errors.
	// It's intended to increment a metric for monitoring, such
	// as an expvar or Prometheus metric.
	// The errType consists of only ASCII word characters.
	CountError	func(errType string)

	// t1, if non-nil, is the standard library Transport using
	// this transport. Its settings are used (but not its
	// RoundTrip method, etc).
	t1	*http.Transport

	connPoolOnce	sync.Once
	connPoolOrDef	ClientConnPool	// non-nil version of ConnPool
}

func (t *Transport) maxHeaderListSize() uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:182
	_go_fuzz_dep_.CoverTab[74394]++
											if t.MaxHeaderListSize == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:183
		_go_fuzz_dep_.CoverTab[74397]++
												return 10 << 20
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:184
		// _ = "end of CoverTab[74397]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:185
		_go_fuzz_dep_.CoverTab[74398]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:185
		// _ = "end of CoverTab[74398]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:185
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:185
	// _ = "end of CoverTab[74394]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:185
	_go_fuzz_dep_.CoverTab[74395]++
											if t.MaxHeaderListSize == 0xffffffff {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:186
		_go_fuzz_dep_.CoverTab[74399]++
												return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:187
		// _ = "end of CoverTab[74399]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:188
		_go_fuzz_dep_.CoverTab[74400]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:188
		// _ = "end of CoverTab[74400]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:188
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:188
	// _ = "end of CoverTab[74395]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:188
	_go_fuzz_dep_.CoverTab[74396]++
											return t.MaxHeaderListSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:189
	// _ = "end of CoverTab[74396]"
}

func (t *Transport) maxFrameReadSize() uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:192
	_go_fuzz_dep_.CoverTab[74401]++
											if t.MaxReadFrameSize == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:193
		_go_fuzz_dep_.CoverTab[74405]++
												return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:194
		// _ = "end of CoverTab[74405]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:195
		_go_fuzz_dep_.CoverTab[74406]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:195
		// _ = "end of CoverTab[74406]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:195
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:195
	// _ = "end of CoverTab[74401]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:195
	_go_fuzz_dep_.CoverTab[74402]++
											if t.MaxReadFrameSize < minMaxFrameSize {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:196
		_go_fuzz_dep_.CoverTab[74407]++
												return minMaxFrameSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:197
		// _ = "end of CoverTab[74407]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:198
		_go_fuzz_dep_.CoverTab[74408]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:198
		// _ = "end of CoverTab[74408]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:198
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:198
	// _ = "end of CoverTab[74402]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:198
	_go_fuzz_dep_.CoverTab[74403]++
											if t.MaxReadFrameSize > maxFrameSize {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:199
		_go_fuzz_dep_.CoverTab[74409]++
												return maxFrameSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:200
		// _ = "end of CoverTab[74409]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:201
		_go_fuzz_dep_.CoverTab[74410]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:201
		// _ = "end of CoverTab[74410]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:201
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:201
	// _ = "end of CoverTab[74403]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:201
	_go_fuzz_dep_.CoverTab[74404]++
											return t.MaxReadFrameSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:202
	// _ = "end of CoverTab[74404]"
}

func (t *Transport) disableCompression() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:205
	_go_fuzz_dep_.CoverTab[74411]++
											return t.DisableCompression || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:206
		_go_fuzz_dep_.CoverTab[74412]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:206
		return (t.t1 != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:206
			_go_fuzz_dep_.CoverTab[74413]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:206
			return t.t1.DisableCompression
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:206
			// _ = "end of CoverTab[74413]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:206
		}())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:206
		// _ = "end of CoverTab[74412]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:206
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:206
	// _ = "end of CoverTab[74411]"
}

func (t *Transport) pingTimeout() time.Duration {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:209
	_go_fuzz_dep_.CoverTab[74414]++
											if t.PingTimeout == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:210
		_go_fuzz_dep_.CoverTab[74416]++
												return 15 * time.Second
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:211
		// _ = "end of CoverTab[74416]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:212
		_go_fuzz_dep_.CoverTab[74417]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:212
		// _ = "end of CoverTab[74417]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:212
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:212
	// _ = "end of CoverTab[74414]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:212
	_go_fuzz_dep_.CoverTab[74415]++
											return t.PingTimeout
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:213
	// _ = "end of CoverTab[74415]"

}

// ConfigureTransport configures a net/http HTTP/1 Transport to use HTTP/2.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:217
// It returns an error if t1 has already been HTTP/2-enabled.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:217
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:217
// Use ConfigureTransports instead to configure the HTTP/2 Transport.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:221
func ConfigureTransport(t1 *http.Transport) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:221
	_go_fuzz_dep_.CoverTab[74418]++
											_, err := ConfigureTransports(t1)
											return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:223
	// _ = "end of CoverTab[74418]"
}

// ConfigureTransports configures a net/http HTTP/1 Transport to use HTTP/2.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:226
// It returns a new HTTP/2 Transport for further configuration.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:226
// It returns an error if t1 has already been HTTP/2-enabled.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:229
func ConfigureTransports(t1 *http.Transport) (*Transport, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:229
	_go_fuzz_dep_.CoverTab[74419]++
											return configureTransports(t1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:230
	// _ = "end of CoverTab[74419]"
}

func configureTransports(t1 *http.Transport) (*Transport, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:233
	_go_fuzz_dep_.CoverTab[74420]++
											connPool := new(clientConnPool)
											t2 := &Transport{
		ConnPool:	noDialClientConnPool{connPool},
		t1:		t1,
	}
	connPool.t = t2
	if err := registerHTTPSProtocol(t1, noDialH2RoundTripper{t2}); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:240
		_go_fuzz_dep_.CoverTab[74427]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:241
		// _ = "end of CoverTab[74427]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:242
		_go_fuzz_dep_.CoverTab[74428]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:242
		// _ = "end of CoverTab[74428]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:242
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:242
	// _ = "end of CoverTab[74420]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:242
	_go_fuzz_dep_.CoverTab[74421]++
											if t1.TLSClientConfig == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:243
		_go_fuzz_dep_.CoverTab[74429]++
												t1.TLSClientConfig = new(tls.Config)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:244
		// _ = "end of CoverTab[74429]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:245
		_go_fuzz_dep_.CoverTab[74430]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:245
		// _ = "end of CoverTab[74430]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:245
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:245
	// _ = "end of CoverTab[74421]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:245
	_go_fuzz_dep_.CoverTab[74422]++
											if !strSliceContains(t1.TLSClientConfig.NextProtos, "h2") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:246
		_go_fuzz_dep_.CoverTab[74431]++
												t1.TLSClientConfig.NextProtos = append([]string{"h2"}, t1.TLSClientConfig.NextProtos...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:247
		// _ = "end of CoverTab[74431]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:248
		_go_fuzz_dep_.CoverTab[74432]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:248
		// _ = "end of CoverTab[74432]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:248
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:248
	// _ = "end of CoverTab[74422]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:248
	_go_fuzz_dep_.CoverTab[74423]++
											if !strSliceContains(t1.TLSClientConfig.NextProtos, "http/1.1") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:249
		_go_fuzz_dep_.CoverTab[74433]++
												t1.TLSClientConfig.NextProtos = append(t1.TLSClientConfig.NextProtos, "http/1.1")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:250
		// _ = "end of CoverTab[74433]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:251
		_go_fuzz_dep_.CoverTab[74434]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:251
		// _ = "end of CoverTab[74434]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:251
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:251
	// _ = "end of CoverTab[74423]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:251
	_go_fuzz_dep_.CoverTab[74424]++
											upgradeFn := func(authority string, c *tls.Conn) http.RoundTripper {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:252
		_go_fuzz_dep_.CoverTab[74435]++
												addr := authorityAddr("https", authority)
												if used, err := connPool.addConnIfNeeded(addr, t2, c); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:254
			_go_fuzz_dep_.CoverTab[74437]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:254
			_curRoutineNum64_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:254
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum64_)
													go c.Close()
													return erringRoundTripper{err}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:256
			// _ = "end of CoverTab[74437]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:257
			_go_fuzz_dep_.CoverTab[74438]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:257
			if !used {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:257
				_go_fuzz_dep_.CoverTab[74439]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:257
				_curRoutineNum65_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:257
				_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum65_)

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:262
				go c.Close()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:262
				// _ = "end of CoverTab[74439]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:263
				_go_fuzz_dep_.CoverTab[74440]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:263
				// _ = "end of CoverTab[74440]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:263
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:263
			// _ = "end of CoverTab[74438]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:263
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:263
		// _ = "end of CoverTab[74435]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:263
		_go_fuzz_dep_.CoverTab[74436]++
												return t2
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:264
		// _ = "end of CoverTab[74436]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:265
	// _ = "end of CoverTab[74424]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:265
	_go_fuzz_dep_.CoverTab[74425]++
											if m := t1.TLSNextProto; len(m) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:266
		_go_fuzz_dep_.CoverTab[74441]++
												t1.TLSNextProto = map[string]func(string, *tls.Conn) http.RoundTripper{
			"h2": upgradeFn,
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:269
		// _ = "end of CoverTab[74441]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:270
		_go_fuzz_dep_.CoverTab[74442]++
												m["h2"] = upgradeFn
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:271
		// _ = "end of CoverTab[74442]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:272
	// _ = "end of CoverTab[74425]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:272
	_go_fuzz_dep_.CoverTab[74426]++
											return t2, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:273
	// _ = "end of CoverTab[74426]"
}

func (t *Transport) connPool() ClientConnPool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:276
	_go_fuzz_dep_.CoverTab[74443]++
											t.connPoolOnce.Do(t.initConnPool)
											return t.connPoolOrDef
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:278
	// _ = "end of CoverTab[74443]"
}

func (t *Transport) initConnPool() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:281
	_go_fuzz_dep_.CoverTab[74444]++
											if t.ConnPool != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:282
		_go_fuzz_dep_.CoverTab[74445]++
												t.connPoolOrDef = t.ConnPool
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:283
		// _ = "end of CoverTab[74445]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:284
		_go_fuzz_dep_.CoverTab[74446]++
												t.connPoolOrDef = &clientConnPool{t: t}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:285
		// _ = "end of CoverTab[74446]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:286
	// _ = "end of CoverTab[74444]"
}

// ClientConn is the state of a single HTTP/2 client connection to an
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:289
// HTTP/2 server.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:291
type ClientConn struct {
	t		*Transport
	tconn		net.Conn	// usually *tls.Conn, except specialized impls
	tconnClosed	bool
	tlsState	*tls.ConnectionState	// nil only for specialized impls
	reused		uint32			// whether conn is being reused; atomic
	singleUse	bool			// whether being used for a single http.Request
	getConnCalled	bool			// used by clientConnPool

	// readLoop goroutine fields:
	readerDone	chan struct{}	// closed on error
	readerErr	error		// set before readerDone is closed

	idleTimeout	time.Duration	// or 0 for never
	idleTimer	*time.Timer

	mu		sync.Mutex	// guards following
	cond		*sync.Cond	// hold mu; broadcast on flow/closed changes
	flow		outflow		// our conn-level flow control quota (cs.outflow is per stream)
	inflow		inflow		// peer's conn-level flow control
	doNotReuse	bool		// whether conn is marked to not be reused for any future requests
	closing		bool
	closed		bool
	seenSettings	bool				// true if we've seen a settings frame, false otherwise
	wantSettingsAck	bool				// we sent a SETTINGS frame and haven't heard back
	goAway		*GoAwayFrame			// if non-nil, the GoAwayFrame we received
	goAwayDebug	string				// goAway frame's debug data, retained as a string
	streams		map[uint32]*clientStream	// client-initiated
	streamsReserved	int				// incr by ReserveNewRequest; decr on RoundTrip
	nextStreamID	uint32
	pendingRequests	int				// requests blocked and waiting to be sent because len(streams) == maxConcurrentStreams
	pings		map[[8]byte]chan struct{}	// in flight ping data to notification channel
	br		*bufio.Reader
	lastActive	time.Time
	lastIdle	time.Time	// time last idle
	// Settings from peer: (also guarded by wmu)
	maxFrameSize		uint32
	maxConcurrentStreams	uint32
	peerMaxHeaderListSize	uint64
	peerMaxHeaderTableSize	uint32
	initialWindowSize	uint32

	// reqHeaderMu is a 1-element semaphore channel controlling access to sending new requests.
	// Write to reqHeaderMu to lock it, read from it to unlock.
	// Lock reqmu BEFORE mu or wmu.
	reqHeaderMu	chan struct{}

	// wmu is held while writing.
	// Acquire BEFORE mu when holding both, to avoid blocking mu on network writes.
	// Only acquire both at the same time when changing peer settings.
	wmu	sync.Mutex
	bw	*bufio.Writer
	fr	*Framer
	werr	error		// first write error that has occurred
	hbuf	bytes.Buffer	// HPACK encoder writes into this
	henc	*hpack.Encoder
}

// clientStream is the state for a single HTTP/2 stream. One of these
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:349
// is created for each Transport.RoundTrip call.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:351
type clientStream struct {
	cc	*ClientConn

	// Fields of Request that we may access even after the response body is closed.
	ctx		context.Context
	reqCancel	<-chan struct{}

	trace		*httptrace.ClientTrace	// or nil
	ID		uint32
	bufPipe		pipe	// buffered pipe with the flow-controlled response payload
	requestedGzip	bool
	isHead		bool

	abortOnce	sync.Once
	abort		chan struct{}	// closed to signal stream should end immediately
	abortErr	error		// set if abort is closed

	peerClosed	chan struct{}	// closed when the peer sends an END_STREAM flag
	donec		chan struct{}	// closed after the stream is in the closed state
	on100		chan struct{}	// buffered; written to if a 100 is received

	respHeaderRecv	chan struct{}	// closed when headers are received
	res		*http.Response	// set if respHeaderRecv is closed

	flow		outflow	// guarded by cc.mu
	inflow		inflow	// guarded by cc.mu
	bytesRemain	int64	// -1 means unknown; owned by transportResponseBody.Read
	readErr		error	// sticky read error; owned by transportResponseBody.Read

	reqBody			io.ReadCloser
	reqBodyContentLength	int64		// -1 means unknown
	reqBodyClosed		chan struct{}	// guarded by cc.mu; non-nil on Close, closed when done

	// owned by writeRequest:
	sentEndStream	bool	// sent an END_STREAM flag to the peer
	sentHeaders	bool

	// owned by clientConnReadLoop:
	firstByte	bool	// got the first response byte
	pastHeaders	bool	// got first MetaHeadersFrame (actual headers)
	pastTrailers	bool	// got optional second MetaHeadersFrame (trailers)
	num1xx		uint8	// number of 1xx responses seen
	readClosed	bool	// peer sent an END_STREAM flag
	readAborted	bool	// read loop reset the stream

	trailer		http.Header	// accumulated trailers
	resTrailer	*http.Header	// client's Response.Trailer
}

var got1xxFuncForTests func(int, textproto.MIMEHeader) error

// get1xxTraceFunc returns the value of request's httptrace.ClientTrace.Got1xxResponse func,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:402
// if any. It returns nil if not set or if the Go version is too old.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:404
func (cs *clientStream) get1xxTraceFunc() func(int, textproto.MIMEHeader) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:404
	_go_fuzz_dep_.CoverTab[74447]++
											if fn := got1xxFuncForTests; fn != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:405
		_go_fuzz_dep_.CoverTab[74449]++
												return fn
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:406
		// _ = "end of CoverTab[74449]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:407
		_go_fuzz_dep_.CoverTab[74450]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:407
		// _ = "end of CoverTab[74450]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:407
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:407
	// _ = "end of CoverTab[74447]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:407
	_go_fuzz_dep_.CoverTab[74448]++
											return traceGot1xxResponseFunc(cs.trace)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:408
	// _ = "end of CoverTab[74448]"
}

func (cs *clientStream) abortStream(err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:411
	_go_fuzz_dep_.CoverTab[74451]++
											cs.cc.mu.Lock()
											defer cs.cc.mu.Unlock()
											cs.abortStreamLocked(err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:414
	// _ = "end of CoverTab[74451]"
}

func (cs *clientStream) abortStreamLocked(err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:417
	_go_fuzz_dep_.CoverTab[74452]++
											cs.abortOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:418
		_go_fuzz_dep_.CoverTab[74455]++
												cs.abortErr = err
												close(cs.abort)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:420
		// _ = "end of CoverTab[74455]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:421
	// _ = "end of CoverTab[74452]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:421
	_go_fuzz_dep_.CoverTab[74453]++
											if cs.reqBody != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:422
		_go_fuzz_dep_.CoverTab[74456]++
												cs.closeReqBodyLocked()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:423
		// _ = "end of CoverTab[74456]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:424
		_go_fuzz_dep_.CoverTab[74457]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:424
		// _ = "end of CoverTab[74457]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:424
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:424
	// _ = "end of CoverTab[74453]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:424
	_go_fuzz_dep_.CoverTab[74454]++

											if cs.cc.cond != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:426
		_go_fuzz_dep_.CoverTab[74458]++

												cs.cc.cond.Broadcast()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:428
		// _ = "end of CoverTab[74458]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:429
		_go_fuzz_dep_.CoverTab[74459]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:429
		// _ = "end of CoverTab[74459]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:429
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:429
	// _ = "end of CoverTab[74454]"
}

func (cs *clientStream) abortRequestBodyWrite() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:432
	_go_fuzz_dep_.CoverTab[74460]++
											cc := cs.cc
											cc.mu.Lock()
											defer cc.mu.Unlock()
											if cs.reqBody != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:436
		_go_fuzz_dep_.CoverTab[74461]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:436
		return cs.reqBodyClosed == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:436
		// _ = "end of CoverTab[74461]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:436
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:436
		_go_fuzz_dep_.CoverTab[74462]++
												cs.closeReqBodyLocked()
												cc.cond.Broadcast()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:438
		// _ = "end of CoverTab[74462]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:439
		_go_fuzz_dep_.CoverTab[74463]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:439
		// _ = "end of CoverTab[74463]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:439
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:439
	// _ = "end of CoverTab[74460]"
}

func (cs *clientStream) closeReqBodyLocked() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:442
	_go_fuzz_dep_.CoverTab[74464]++
											if cs.reqBodyClosed != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:443
		_go_fuzz_dep_.CoverTab[74466]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:444
		// _ = "end of CoverTab[74466]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:445
		_go_fuzz_dep_.CoverTab[74467]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:445
		// _ = "end of CoverTab[74467]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:445
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:445
	// _ = "end of CoverTab[74464]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:445
	_go_fuzz_dep_.CoverTab[74465]++
											cs.reqBodyClosed = make(chan struct{})
											reqBodyClosed := cs.reqBodyClosed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:447
	_curRoutineNum66_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:447
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum66_)
											go func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:448
		_go_fuzz_dep_.CoverTab[74468]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:448
		defer func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:448
			_go_fuzz_dep_.CoverTab[74469]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:448
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum66_)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:448
			// _ = "end of CoverTab[74469]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:448
		}()
												cs.reqBody.Close()
												close(reqBodyClosed)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:450
		// _ = "end of CoverTab[74468]"
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:451
	// _ = "end of CoverTab[74465]"
}

type stickyErrWriter struct {
	conn	net.Conn
	timeout	time.Duration
	err	*error
}

func (sew stickyErrWriter) Write(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:460
	_go_fuzz_dep_.CoverTab[74470]++
											if *sew.err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:461
		_go_fuzz_dep_.CoverTab[74472]++
												return 0, *sew.err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:462
		// _ = "end of CoverTab[74472]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:463
		_go_fuzz_dep_.CoverTab[74473]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:463
		// _ = "end of CoverTab[74473]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:463
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:463
	// _ = "end of CoverTab[74470]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:463
	_go_fuzz_dep_.CoverTab[74471]++
											for {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:464
		_go_fuzz_dep_.CoverTab[74474]++
												if sew.timeout != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:465
			_go_fuzz_dep_.CoverTab[74478]++
													sew.conn.SetWriteDeadline(time.Now().Add(sew.timeout))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:466
			// _ = "end of CoverTab[74478]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:467
			_go_fuzz_dep_.CoverTab[74479]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:467
			// _ = "end of CoverTab[74479]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:467
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:467
		// _ = "end of CoverTab[74474]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:467
		_go_fuzz_dep_.CoverTab[74475]++
												nn, err := sew.conn.Write(p[n:])
												n += nn
												if n < len(p) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:470
			_go_fuzz_dep_.CoverTab[74480]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:470
			return nn > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:470
			// _ = "end of CoverTab[74480]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:470
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:470
			_go_fuzz_dep_.CoverTab[74481]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:470
			return errors.Is(err, os.ErrDeadlineExceeded)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:470
			// _ = "end of CoverTab[74481]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:470
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:470
			_go_fuzz_dep_.CoverTab[74482]++

													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:472
			// _ = "end of CoverTab[74482]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:473
			_go_fuzz_dep_.CoverTab[74483]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:473
			// _ = "end of CoverTab[74483]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:473
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:473
		// _ = "end of CoverTab[74475]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:473
		_go_fuzz_dep_.CoverTab[74476]++
												if sew.timeout != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:474
			_go_fuzz_dep_.CoverTab[74484]++
													sew.conn.SetWriteDeadline(time.Time{})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:475
			// _ = "end of CoverTab[74484]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:476
			_go_fuzz_dep_.CoverTab[74485]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:476
			// _ = "end of CoverTab[74485]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:476
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:476
		// _ = "end of CoverTab[74476]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:476
		_go_fuzz_dep_.CoverTab[74477]++
												*sew.err = err
												return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:478
		// _ = "end of CoverTab[74477]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:479
	// _ = "end of CoverTab[74471]"
}

// noCachedConnError is the concrete type of ErrNoCachedConn, which
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:482
// needs to be detected by net/http regardless of whether it's its
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:482
// bundled version (in h2_bundle.go with a rewritten type name) or
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:482
// from a user's x/net/http2. As such, as it has a unique method name
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:482
// (IsHTTP2NoCachedConnError) that net/http sniffs for via func
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:482
// isNoCachedConnError.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:488
type noCachedConnError struct{}

func (noCachedConnError) IsHTTP2NoCachedConnError() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:490
	_go_fuzz_dep_.CoverTab[74486]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:490
	// _ = "end of CoverTab[74486]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:490
}
func (noCachedConnError) Error() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:491
	_go_fuzz_dep_.CoverTab[74487]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:491
	return "http2: no cached connection was available"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:491
	// _ = "end of CoverTab[74487]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:491
}

// isNoCachedConnError reports whether err is of type noCachedConnError
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:493
// or its equivalent renamed type in net/http2's h2_bundle.go. Both types
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:493
// may coexist in the same running program.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:496
func isNoCachedConnError(err error) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:496
	_go_fuzz_dep_.CoverTab[74488]++
											_, ok := err.(interface{ IsHTTP2NoCachedConnError() })
											return ok
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:498
	// _ = "end of CoverTab[74488]"
}

var ErrNoCachedConn error = noCachedConnError{}

// RoundTripOpt are options for the Transport.RoundTripOpt method.
type RoundTripOpt struct {
	// OnlyCachedConn controls whether RoundTripOpt may
	// create a new TCP connection. If set true and
	// no cached connection is available, RoundTripOpt
	// will return ErrNoCachedConn.
	OnlyCachedConn bool
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:512
	_go_fuzz_dep_.CoverTab[74489]++
											return t.RoundTripOpt(req, RoundTripOpt{})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:513
	// _ = "end of CoverTab[74489]"
}

// authorityAddr returns a given authority (a host/IP, or host:port / ip:port)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:516
// and returns a host:port. The port 443 is added if needed.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:518
func authorityAddr(scheme string, authority string) (addr string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:518
	_go_fuzz_dep_.CoverTab[74490]++
											host, port, err := net.SplitHostPort(authority)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:520
		_go_fuzz_dep_.CoverTab[74494]++
												port = "443"
												if scheme == "http" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:522
			_go_fuzz_dep_.CoverTab[74496]++
													port = "80"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:523
			// _ = "end of CoverTab[74496]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:524
			_go_fuzz_dep_.CoverTab[74497]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:524
			// _ = "end of CoverTab[74497]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:524
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:524
		// _ = "end of CoverTab[74494]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:524
		_go_fuzz_dep_.CoverTab[74495]++
												host = authority
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:525
		// _ = "end of CoverTab[74495]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:526
		_go_fuzz_dep_.CoverTab[74498]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:526
		// _ = "end of CoverTab[74498]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:526
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:526
	// _ = "end of CoverTab[74490]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:526
	_go_fuzz_dep_.CoverTab[74491]++
											if a, err := idna.ToASCII(host); err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:527
		_go_fuzz_dep_.CoverTab[74499]++
												host = a
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:528
		// _ = "end of CoverTab[74499]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:529
		_go_fuzz_dep_.CoverTab[74500]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:529
		// _ = "end of CoverTab[74500]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:529
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:529
	// _ = "end of CoverTab[74491]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:529
	_go_fuzz_dep_.CoverTab[74492]++

											if strings.HasPrefix(host, "[") && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:531
		_go_fuzz_dep_.CoverTab[74501]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:531
		return strings.HasSuffix(host, "]")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:531
		// _ = "end of CoverTab[74501]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:531
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:531
		_go_fuzz_dep_.CoverTab[74502]++
												return host + ":" + port
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:532
		// _ = "end of CoverTab[74502]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:533
		_go_fuzz_dep_.CoverTab[74503]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:533
		// _ = "end of CoverTab[74503]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:533
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:533
	// _ = "end of CoverTab[74492]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:533
	_go_fuzz_dep_.CoverTab[74493]++
											return net.JoinHostPort(host, port)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:534
	// _ = "end of CoverTab[74493]"
}

var retryBackoffHook func(time.Duration) *time.Timer

func backoffNewTimer(d time.Duration) *time.Timer {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:539
	_go_fuzz_dep_.CoverTab[74504]++
											if retryBackoffHook != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:540
		_go_fuzz_dep_.CoverTab[74506]++
												return retryBackoffHook(d)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:541
		// _ = "end of CoverTab[74506]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:542
		_go_fuzz_dep_.CoverTab[74507]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:542
		// _ = "end of CoverTab[74507]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:542
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:542
	// _ = "end of CoverTab[74504]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:542
	_go_fuzz_dep_.CoverTab[74505]++
											return time.NewTimer(d)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:543
	// _ = "end of CoverTab[74505]"
}

// RoundTripOpt is like RoundTrip, but takes options.
func (t *Transport) RoundTripOpt(req *http.Request, opt RoundTripOpt) (*http.Response, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:547
	_go_fuzz_dep_.CoverTab[74508]++
											if !(req.URL.Scheme == "https" || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:548
		_go_fuzz_dep_.CoverTab[74510]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:548
		return (req.URL.Scheme == "http" && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:548
			_go_fuzz_dep_.CoverTab[74511]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:548
			return t.AllowHTTP
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:548
			// _ = "end of CoverTab[74511]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:548
		}())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:548
		// _ = "end of CoverTab[74510]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:548
	}()) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:548
		_go_fuzz_dep_.CoverTab[74512]++
												return nil, errors.New("http2: unsupported scheme")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:549
		// _ = "end of CoverTab[74512]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:550
		_go_fuzz_dep_.CoverTab[74513]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:550
		// _ = "end of CoverTab[74513]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:550
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:550
	// _ = "end of CoverTab[74508]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:550
	_go_fuzz_dep_.CoverTab[74509]++

											addr := authorityAddr(req.URL.Scheme, req.URL.Host)
											for retry := 0; ; retry++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:553
		_go_fuzz_dep_.CoverTab[74514]++
												cc, err := t.connPool().GetClientConn(req, addr)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:555
			_go_fuzz_dep_.CoverTab[74518]++
													t.vlogf("http2: Transport failed to get client conn for %s: %v", addr, err)
													return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:557
			// _ = "end of CoverTab[74518]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:558
			_go_fuzz_dep_.CoverTab[74519]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:558
			// _ = "end of CoverTab[74519]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:558
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:558
		// _ = "end of CoverTab[74514]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:558
		_go_fuzz_dep_.CoverTab[74515]++
												reused := !atomic.CompareAndSwapUint32(&cc.reused, 0, 1)
												traceGotConn(req, cc, reused)
												res, err := cc.RoundTrip(req)
												if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:562
			_go_fuzz_dep_.CoverTab[74520]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:562
			return retry <= 6
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:562
			// _ = "end of CoverTab[74520]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:562
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:562
			_go_fuzz_dep_.CoverTab[74521]++
													roundTripErr := err
													if req, err = shouldRetryRequest(req, err); err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:564
				_go_fuzz_dep_.CoverTab[74522]++

														if retry == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:566
					_go_fuzz_dep_.CoverTab[74524]++
															t.vlogf("RoundTrip retrying after failure: %v", roundTripErr)
															continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:568
					// _ = "end of CoverTab[74524]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:569
					_go_fuzz_dep_.CoverTab[74525]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:569
					// _ = "end of CoverTab[74525]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:569
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:569
				// _ = "end of CoverTab[74522]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:569
				_go_fuzz_dep_.CoverTab[74523]++
														backoff := float64(uint(1) << (uint(retry) - 1))
														backoff += backoff * (0.1 * mathrand.Float64())
														d := time.Second * time.Duration(backoff)
														timer := backoffNewTimer(d)
														select {
				case <-timer.C:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:575
					_go_fuzz_dep_.CoverTab[74526]++
															t.vlogf("RoundTrip retrying after failure: %v", roundTripErr)
															continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:577
					// _ = "end of CoverTab[74526]"
				case <-req.Context().Done():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:578
					_go_fuzz_dep_.CoverTab[74527]++
															timer.Stop()
															err = req.Context().Err()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:580
					// _ = "end of CoverTab[74527]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:581
				// _ = "end of CoverTab[74523]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:582
				_go_fuzz_dep_.CoverTab[74528]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:582
				// _ = "end of CoverTab[74528]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:582
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:582
			// _ = "end of CoverTab[74521]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:583
			_go_fuzz_dep_.CoverTab[74529]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:583
			// _ = "end of CoverTab[74529]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:583
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:583
		// _ = "end of CoverTab[74515]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:583
		_go_fuzz_dep_.CoverTab[74516]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:584
			_go_fuzz_dep_.CoverTab[74530]++
													t.vlogf("RoundTrip failure: %v", err)
													return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:586
			// _ = "end of CoverTab[74530]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:587
			_go_fuzz_dep_.CoverTab[74531]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:587
			// _ = "end of CoverTab[74531]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:587
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:587
		// _ = "end of CoverTab[74516]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:587
		_go_fuzz_dep_.CoverTab[74517]++
												return res, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:588
		// _ = "end of CoverTab[74517]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:589
	// _ = "end of CoverTab[74509]"
}

// CloseIdleConnections closes any connections which were previously
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:592
// connected from previous requests but are now sitting idle.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:592
// It does not interrupt any connections currently in use.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:595
func (t *Transport) CloseIdleConnections() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:595
	_go_fuzz_dep_.CoverTab[74532]++
											if cp, ok := t.connPool().(clientConnPoolIdleCloser); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:596
		_go_fuzz_dep_.CoverTab[74533]++
												cp.closeIdleConnections()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:597
		// _ = "end of CoverTab[74533]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:598
		_go_fuzz_dep_.CoverTab[74534]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:598
		// _ = "end of CoverTab[74534]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:598
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:598
	// _ = "end of CoverTab[74532]"
}

var (
	errClientConnClosed	= errors.New("http2: client conn is closed")
	errClientConnUnusable	= errors.New("http2: client conn not usable")
	errClientConnGotGoAway	= errors.New("http2: Transport received Server's graceful shutdown GOAWAY")
)

// shouldRetryRequest is called by RoundTrip when a request fails to get
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:607
// response headers. It is always called with a non-nil error.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:607
// It returns either a request to retry (either the same request, or a
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:607
// modified clone), or an error if the request can't be replayed.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:611
func shouldRetryRequest(req *http.Request, err error) (*http.Request, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:611
	_go_fuzz_dep_.CoverTab[74535]++
											if !canRetryError(err) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:612
		_go_fuzz_dep_.CoverTab[74540]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:613
		// _ = "end of CoverTab[74540]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:614
		_go_fuzz_dep_.CoverTab[74541]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:614
		// _ = "end of CoverTab[74541]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:614
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:614
	// _ = "end of CoverTab[74535]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:614
	_go_fuzz_dep_.CoverTab[74536]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:617
	if req.Body == nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:617
		_go_fuzz_dep_.CoverTab[74542]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:617
		return req.Body == http.NoBody
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:617
		// _ = "end of CoverTab[74542]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:617
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:617
		_go_fuzz_dep_.CoverTab[74543]++
												return req, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:618
		// _ = "end of CoverTab[74543]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:619
		_go_fuzz_dep_.CoverTab[74544]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:619
		// _ = "end of CoverTab[74544]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:619
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:619
	// _ = "end of CoverTab[74536]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:619
	_go_fuzz_dep_.CoverTab[74537]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:623
	if req.GetBody != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:623
		_go_fuzz_dep_.CoverTab[74545]++
												body, err := req.GetBody()
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:625
			_go_fuzz_dep_.CoverTab[74547]++
													return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:626
			// _ = "end of CoverTab[74547]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:627
			_go_fuzz_dep_.CoverTab[74548]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:627
			// _ = "end of CoverTab[74548]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:627
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:627
		// _ = "end of CoverTab[74545]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:627
		_go_fuzz_dep_.CoverTab[74546]++
												newReq := *req
												newReq.Body = body
												return &newReq, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:630
		// _ = "end of CoverTab[74546]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:631
		_go_fuzz_dep_.CoverTab[74549]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:631
		// _ = "end of CoverTab[74549]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:631
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:631
	// _ = "end of CoverTab[74537]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:631
	_go_fuzz_dep_.CoverTab[74538]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:636
	if err == errClientConnUnusable {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:636
		_go_fuzz_dep_.CoverTab[74550]++
												return req, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:637
		// _ = "end of CoverTab[74550]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:638
		_go_fuzz_dep_.CoverTab[74551]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:638
		// _ = "end of CoverTab[74551]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:638
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:638
	// _ = "end of CoverTab[74538]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:638
	_go_fuzz_dep_.CoverTab[74539]++

											return nil, fmt.Errorf("http2: Transport: cannot retry err [%v] after Request.Body was written; define Request.GetBody to avoid this error", err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:640
	// _ = "end of CoverTab[74539]"
}

func canRetryError(err error) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:643
	_go_fuzz_dep_.CoverTab[74552]++
											if err == errClientConnUnusable || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:644
		_go_fuzz_dep_.CoverTab[74555]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:644
		return err == errClientConnGotGoAway
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:644
		// _ = "end of CoverTab[74555]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:644
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:644
		_go_fuzz_dep_.CoverTab[74556]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:645
		// _ = "end of CoverTab[74556]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:646
		_go_fuzz_dep_.CoverTab[74557]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:646
		// _ = "end of CoverTab[74557]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:646
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:646
	// _ = "end of CoverTab[74552]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:646
	_go_fuzz_dep_.CoverTab[74553]++
											if se, ok := err.(StreamError); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:647
		_go_fuzz_dep_.CoverTab[74558]++
												if se.Code == ErrCodeProtocol && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:648
			_go_fuzz_dep_.CoverTab[74560]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:648
			return se.Cause == errFromPeer
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:648
			// _ = "end of CoverTab[74560]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:648
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:648
			_go_fuzz_dep_.CoverTab[74561]++

													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:650
			// _ = "end of CoverTab[74561]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:651
			_go_fuzz_dep_.CoverTab[74562]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:651
			// _ = "end of CoverTab[74562]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:651
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:651
		// _ = "end of CoverTab[74558]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:651
		_go_fuzz_dep_.CoverTab[74559]++
												return se.Code == ErrCodeRefusedStream
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:652
		// _ = "end of CoverTab[74559]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:653
		_go_fuzz_dep_.CoverTab[74563]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:653
		// _ = "end of CoverTab[74563]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:653
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:653
	// _ = "end of CoverTab[74553]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:653
	_go_fuzz_dep_.CoverTab[74554]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:654
	// _ = "end of CoverTab[74554]"
}

func (t *Transport) dialClientConn(ctx context.Context, addr string, singleUse bool) (*ClientConn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:657
	_go_fuzz_dep_.CoverTab[74564]++
											host, _, err := net.SplitHostPort(addr)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:659
		_go_fuzz_dep_.CoverTab[74567]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:660
		// _ = "end of CoverTab[74567]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:661
		_go_fuzz_dep_.CoverTab[74568]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:661
		// _ = "end of CoverTab[74568]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:661
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:661
	// _ = "end of CoverTab[74564]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:661
	_go_fuzz_dep_.CoverTab[74565]++
											tconn, err := t.dialTLS(ctx, "tcp", addr, t.newTLSConfig(host))
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:663
		_go_fuzz_dep_.CoverTab[74569]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:664
		// _ = "end of CoverTab[74569]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:665
		_go_fuzz_dep_.CoverTab[74570]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:665
		// _ = "end of CoverTab[74570]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:665
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:665
	// _ = "end of CoverTab[74565]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:665
	_go_fuzz_dep_.CoverTab[74566]++
											return t.newClientConn(tconn, singleUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:666
	// _ = "end of CoverTab[74566]"
}

func (t *Transport) newTLSConfig(host string) *tls.Config {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:669
	_go_fuzz_dep_.CoverTab[74571]++
											cfg := new(tls.Config)
											if t.TLSClientConfig != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:671
		_go_fuzz_dep_.CoverTab[74575]++
												*cfg = *t.TLSClientConfig.Clone()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:672
		// _ = "end of CoverTab[74575]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:673
		_go_fuzz_dep_.CoverTab[74576]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:673
		// _ = "end of CoverTab[74576]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:673
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:673
	// _ = "end of CoverTab[74571]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:673
	_go_fuzz_dep_.CoverTab[74572]++
											if !strSliceContains(cfg.NextProtos, NextProtoTLS) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:674
		_go_fuzz_dep_.CoverTab[74577]++
												cfg.NextProtos = append([]string{NextProtoTLS}, cfg.NextProtos...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:675
		// _ = "end of CoverTab[74577]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:676
		_go_fuzz_dep_.CoverTab[74578]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:676
		// _ = "end of CoverTab[74578]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:676
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:676
	// _ = "end of CoverTab[74572]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:676
	_go_fuzz_dep_.CoverTab[74573]++
											if cfg.ServerName == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:677
		_go_fuzz_dep_.CoverTab[74579]++
												cfg.ServerName = host
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:678
		// _ = "end of CoverTab[74579]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:679
		_go_fuzz_dep_.CoverTab[74580]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:679
		// _ = "end of CoverTab[74580]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:679
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:679
	// _ = "end of CoverTab[74573]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:679
	_go_fuzz_dep_.CoverTab[74574]++
											return cfg
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:680
	// _ = "end of CoverTab[74574]"
}

func (t *Transport) dialTLS(ctx context.Context, network, addr string, tlsCfg *tls.Config) (net.Conn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:683
	_go_fuzz_dep_.CoverTab[74581]++
											if t.DialTLSContext != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:684
		_go_fuzz_dep_.CoverTab[74586]++
												return t.DialTLSContext(ctx, network, addr, tlsCfg)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:685
		// _ = "end of CoverTab[74586]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:686
		_go_fuzz_dep_.CoverTab[74587]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:686
		if t.DialTLS != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:686
			_go_fuzz_dep_.CoverTab[74588]++
													return t.DialTLS(network, addr, tlsCfg)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:687
			// _ = "end of CoverTab[74588]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:688
			_go_fuzz_dep_.CoverTab[74589]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:688
			// _ = "end of CoverTab[74589]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:688
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:688
		// _ = "end of CoverTab[74587]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:688
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:688
	// _ = "end of CoverTab[74581]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:688
	_go_fuzz_dep_.CoverTab[74582]++

											tlsCn, err := t.dialTLSWithContext(ctx, network, addr, tlsCfg)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:691
		_go_fuzz_dep_.CoverTab[74590]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:692
		// _ = "end of CoverTab[74590]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:693
		_go_fuzz_dep_.CoverTab[74591]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:693
		// _ = "end of CoverTab[74591]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:693
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:693
	// _ = "end of CoverTab[74582]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:693
	_go_fuzz_dep_.CoverTab[74583]++
											state := tlsCn.ConnectionState()
											if p := state.NegotiatedProtocol; p != NextProtoTLS {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:695
		_go_fuzz_dep_.CoverTab[74592]++
												return nil, fmt.Errorf("http2: unexpected ALPN protocol %q; want %q", p, NextProtoTLS)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:696
		// _ = "end of CoverTab[74592]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:697
		_go_fuzz_dep_.CoverTab[74593]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:697
		// _ = "end of CoverTab[74593]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:697
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:697
	// _ = "end of CoverTab[74583]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:697
	_go_fuzz_dep_.CoverTab[74584]++
											if !state.NegotiatedProtocolIsMutual {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:698
		_go_fuzz_dep_.CoverTab[74594]++
												return nil, errors.New("http2: could not negotiate protocol mutually")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:699
		// _ = "end of CoverTab[74594]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:700
		_go_fuzz_dep_.CoverTab[74595]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:700
		// _ = "end of CoverTab[74595]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:700
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:700
	// _ = "end of CoverTab[74584]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:700
	_go_fuzz_dep_.CoverTab[74585]++
											return tlsCn, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:701
	// _ = "end of CoverTab[74585]"
}

// disableKeepAlives reports whether connections should be closed as
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:704
// soon as possible after handling the first request.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:706
func (t *Transport) disableKeepAlives() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:706
	_go_fuzz_dep_.CoverTab[74596]++
											return t.t1 != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:707
		_go_fuzz_dep_.CoverTab[74597]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:707
		return t.t1.DisableKeepAlives
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:707
		// _ = "end of CoverTab[74597]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:707
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:707
	// _ = "end of CoverTab[74596]"
}

func (t *Transport) expectContinueTimeout() time.Duration {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:710
	_go_fuzz_dep_.CoverTab[74598]++
											if t.t1 == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:711
		_go_fuzz_dep_.CoverTab[74600]++
												return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:712
		// _ = "end of CoverTab[74600]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:713
		_go_fuzz_dep_.CoverTab[74601]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:713
		// _ = "end of CoverTab[74601]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:713
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:713
	// _ = "end of CoverTab[74598]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:713
	_go_fuzz_dep_.CoverTab[74599]++
											return t.t1.ExpectContinueTimeout
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:714
	// _ = "end of CoverTab[74599]"
}

func (t *Transport) maxDecoderHeaderTableSize() uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:717
	_go_fuzz_dep_.CoverTab[74602]++
											if v := t.MaxDecoderHeaderTableSize; v > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:718
		_go_fuzz_dep_.CoverTab[74604]++
												return v
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:719
		// _ = "end of CoverTab[74604]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:720
		_go_fuzz_dep_.CoverTab[74605]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:720
		// _ = "end of CoverTab[74605]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:720
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:720
	// _ = "end of CoverTab[74602]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:720
	_go_fuzz_dep_.CoverTab[74603]++
											return initialHeaderTableSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:721
	// _ = "end of CoverTab[74603]"
}

func (t *Transport) maxEncoderHeaderTableSize() uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:724
	_go_fuzz_dep_.CoverTab[74606]++
											if v := t.MaxEncoderHeaderTableSize; v > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:725
		_go_fuzz_dep_.CoverTab[74608]++
												return v
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:726
		// _ = "end of CoverTab[74608]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:727
		_go_fuzz_dep_.CoverTab[74609]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:727
		// _ = "end of CoverTab[74609]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:727
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:727
	// _ = "end of CoverTab[74606]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:727
	_go_fuzz_dep_.CoverTab[74607]++
											return initialHeaderTableSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:728
	// _ = "end of CoverTab[74607]"
}

func (t *Transport) NewClientConn(c net.Conn) (*ClientConn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:731
	_go_fuzz_dep_.CoverTab[74610]++
											return t.newClientConn(c, t.disableKeepAlives())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:732
	// _ = "end of CoverTab[74610]"
}

func (t *Transport) newClientConn(c net.Conn, singleUse bool) (*ClientConn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:735
	_go_fuzz_dep_.CoverTab[74611]++
											cc := &ClientConn{
		t:			t,
		tconn:			c,
		readerDone:		make(chan struct{}),
		nextStreamID:		1,
		maxFrameSize:		16 << 10,
		initialWindowSize:	65535,
		maxConcurrentStreams:	initialMaxConcurrentStreams,
		peerMaxHeaderListSize:	0xffffffffffffffff,
		streams:		make(map[uint32]*clientStream),
		singleUse:		singleUse,
		wantSettingsAck:	true,
		pings:			make(map[[8]byte]chan struct{}),
		reqHeaderMu:		make(chan struct{}, 1),
	}
	if d := t.idleConnTimeout(); d != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:751
		_go_fuzz_dep_.CoverTab[74622]++
												cc.idleTimeout = d
												cc.idleTimer = time.AfterFunc(d, cc.onIdleTimeout)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:753
		// _ = "end of CoverTab[74622]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:754
		_go_fuzz_dep_.CoverTab[74623]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:754
		// _ = "end of CoverTab[74623]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:754
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:754
	// _ = "end of CoverTab[74611]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:754
	_go_fuzz_dep_.CoverTab[74612]++
											if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:755
		_go_fuzz_dep_.CoverTab[74624]++
												t.vlogf("http2: Transport creating client conn %p to %v", cc, c.RemoteAddr())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:756
		// _ = "end of CoverTab[74624]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:757
		_go_fuzz_dep_.CoverTab[74625]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:757
		// _ = "end of CoverTab[74625]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:757
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:757
	// _ = "end of CoverTab[74612]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:757
	_go_fuzz_dep_.CoverTab[74613]++

											cc.cond = sync.NewCond(&cc.mu)
											cc.flow.add(int32(initialWindowSize))

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:764
	cc.bw = bufio.NewWriter(stickyErrWriter{
		conn:		c,
		timeout:	t.WriteByteTimeout,
		err:		&cc.werr,
	})
	cc.br = bufio.NewReader(c)
	cc.fr = NewFramer(cc.bw, cc.br)
	if t.maxFrameReadSize() != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:771
		_go_fuzz_dep_.CoverTab[74626]++
												cc.fr.SetMaxReadFrameSize(t.maxFrameReadSize())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:772
		// _ = "end of CoverTab[74626]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:773
		_go_fuzz_dep_.CoverTab[74627]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:773
		// _ = "end of CoverTab[74627]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:773
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:773
	// _ = "end of CoverTab[74613]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:773
	_go_fuzz_dep_.CoverTab[74614]++
											if t.CountError != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:774
		_go_fuzz_dep_.CoverTab[74628]++
												cc.fr.countError = t.CountError
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:775
		// _ = "end of CoverTab[74628]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:776
		_go_fuzz_dep_.CoverTab[74629]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:776
		// _ = "end of CoverTab[74629]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:776
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:776
	// _ = "end of CoverTab[74614]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:776
	_go_fuzz_dep_.CoverTab[74615]++
											maxHeaderTableSize := t.maxDecoderHeaderTableSize()
											cc.fr.ReadMetaHeaders = hpack.NewDecoder(maxHeaderTableSize, nil)
											cc.fr.MaxHeaderListSize = t.maxHeaderListSize()

											cc.henc = hpack.NewEncoder(&cc.hbuf)
											cc.henc.SetMaxDynamicTableSizeLimit(t.maxEncoderHeaderTableSize())
											cc.peerMaxHeaderTableSize = initialHeaderTableSize

											if t.AllowHTTP {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:785
		_go_fuzz_dep_.CoverTab[74630]++
												cc.nextStreamID = 3
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:786
		// _ = "end of CoverTab[74630]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:787
		_go_fuzz_dep_.CoverTab[74631]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:787
		// _ = "end of CoverTab[74631]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:787
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:787
	// _ = "end of CoverTab[74615]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:787
	_go_fuzz_dep_.CoverTab[74616]++

											if cs, ok := c.(connectionStater); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:789
		_go_fuzz_dep_.CoverTab[74632]++
												state := cs.ConnectionState()
												cc.tlsState = &state
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:791
		// _ = "end of CoverTab[74632]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:792
		_go_fuzz_dep_.CoverTab[74633]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:792
		// _ = "end of CoverTab[74633]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:792
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:792
	// _ = "end of CoverTab[74616]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:792
	_go_fuzz_dep_.CoverTab[74617]++

											initialSettings := []Setting{
		{ID: SettingEnablePush, Val: 0},
		{ID: SettingInitialWindowSize, Val: transportDefaultStreamFlow},
	}
	if max := t.maxFrameReadSize(); max != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:798
		_go_fuzz_dep_.CoverTab[74634]++
												initialSettings = append(initialSettings, Setting{ID: SettingMaxFrameSize, Val: max})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:799
		// _ = "end of CoverTab[74634]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:800
		_go_fuzz_dep_.CoverTab[74635]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:800
		// _ = "end of CoverTab[74635]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:800
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:800
	// _ = "end of CoverTab[74617]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:800
	_go_fuzz_dep_.CoverTab[74618]++
											if max := t.maxHeaderListSize(); max != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:801
		_go_fuzz_dep_.CoverTab[74636]++
												initialSettings = append(initialSettings, Setting{ID: SettingMaxHeaderListSize, Val: max})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:802
		// _ = "end of CoverTab[74636]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:803
		_go_fuzz_dep_.CoverTab[74637]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:803
		// _ = "end of CoverTab[74637]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:803
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:803
	// _ = "end of CoverTab[74618]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:803
	_go_fuzz_dep_.CoverTab[74619]++
											if maxHeaderTableSize != initialHeaderTableSize {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:804
		_go_fuzz_dep_.CoverTab[74638]++
												initialSettings = append(initialSettings, Setting{ID: SettingHeaderTableSize, Val: maxHeaderTableSize})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:805
		// _ = "end of CoverTab[74638]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:806
		_go_fuzz_dep_.CoverTab[74639]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:806
		// _ = "end of CoverTab[74639]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:806
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:806
	// _ = "end of CoverTab[74619]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:806
	_go_fuzz_dep_.CoverTab[74620]++

											cc.bw.Write(clientPreface)
											cc.fr.WriteSettings(initialSettings...)
											cc.fr.WriteWindowUpdate(0, transportDefaultConnFlow)
											cc.inflow.init(transportDefaultConnFlow + initialWindowSize)
											cc.bw.Flush()
											if cc.werr != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:813
		_go_fuzz_dep_.CoverTab[74640]++
												cc.Close()
												return nil, cc.werr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:815
		// _ = "end of CoverTab[74640]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:816
		_go_fuzz_dep_.CoverTab[74641]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:816
		// _ = "end of CoverTab[74641]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:816
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:816
	// _ = "end of CoverTab[74620]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:816
	_go_fuzz_dep_.CoverTab[74621]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:816
	_curRoutineNum67_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:816
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum67_)

											go cc.readLoop()
											return cc, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:819
	// _ = "end of CoverTab[74621]"
}

func (cc *ClientConn) healthCheck() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:822
	_go_fuzz_dep_.CoverTab[74642]++
											pingTimeout := cc.t.pingTimeout()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:826
	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()
	cc.vlogf("http2: Transport sending health check")
	err := cc.Ping(ctx)
	if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:830
		_go_fuzz_dep_.CoverTab[74643]++
												cc.vlogf("http2: Transport health check failure: %v", err)
												cc.closeForLostPing()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:832
		// _ = "end of CoverTab[74643]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:833
		_go_fuzz_dep_.CoverTab[74644]++
												cc.vlogf("http2: Transport health check success")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:834
		// _ = "end of CoverTab[74644]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:835
	// _ = "end of CoverTab[74642]"
}

// SetDoNotReuse marks cc as not reusable for future HTTP requests.
func (cc *ClientConn) SetDoNotReuse() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:839
	_go_fuzz_dep_.CoverTab[74645]++
											cc.mu.Lock()
											defer cc.mu.Unlock()
											cc.doNotReuse = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:842
	// _ = "end of CoverTab[74645]"
}

func (cc *ClientConn) setGoAway(f *GoAwayFrame) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:845
	_go_fuzz_dep_.CoverTab[74646]++
											cc.mu.Lock()
											defer cc.mu.Unlock()

											old := cc.goAway
											cc.goAway = f

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:853
	if cc.goAwayDebug == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:853
		_go_fuzz_dep_.CoverTab[74649]++
												cc.goAwayDebug = string(f.DebugData())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:854
		// _ = "end of CoverTab[74649]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:855
		_go_fuzz_dep_.CoverTab[74650]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:855
		// _ = "end of CoverTab[74650]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:855
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:855
	// _ = "end of CoverTab[74646]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:855
	_go_fuzz_dep_.CoverTab[74647]++
											if old != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:856
		_go_fuzz_dep_.CoverTab[74651]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:856
		return old.ErrCode != ErrCodeNo
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:856
		// _ = "end of CoverTab[74651]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:856
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:856
		_go_fuzz_dep_.CoverTab[74652]++
												cc.goAway.ErrCode = old.ErrCode
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:857
		// _ = "end of CoverTab[74652]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:858
		_go_fuzz_dep_.CoverTab[74653]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:858
		// _ = "end of CoverTab[74653]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:858
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:858
	// _ = "end of CoverTab[74647]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:858
	_go_fuzz_dep_.CoverTab[74648]++
											last := f.LastStreamID
											for streamID, cs := range cc.streams {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:860
		_go_fuzz_dep_.CoverTab[74654]++
												if streamID > last {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:861
			_go_fuzz_dep_.CoverTab[74655]++
													cs.abortStreamLocked(errClientConnGotGoAway)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:862
			// _ = "end of CoverTab[74655]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:863
			_go_fuzz_dep_.CoverTab[74656]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:863
			// _ = "end of CoverTab[74656]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:863
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:863
		// _ = "end of CoverTab[74654]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:864
	// _ = "end of CoverTab[74648]"
}

// CanTakeNewRequest reports whether the connection can take a new request,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:867
// meaning it has not been closed or received or sent a GOAWAY.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:867
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:867
// If the caller is going to immediately make a new request on this
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:867
// connection, use ReserveNewRequest instead.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:872
func (cc *ClientConn) CanTakeNewRequest() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:872
	_go_fuzz_dep_.CoverTab[74657]++
											cc.mu.Lock()
											defer cc.mu.Unlock()
											return cc.canTakeNewRequestLocked()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:875
	// _ = "end of CoverTab[74657]"
}

// ReserveNewRequest is like CanTakeNewRequest but also reserves a
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:878
// concurrent stream in cc. The reservation is decremented on the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:878
// next call to RoundTrip.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:881
func (cc *ClientConn) ReserveNewRequest() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:881
	_go_fuzz_dep_.CoverTab[74658]++
											cc.mu.Lock()
											defer cc.mu.Unlock()
											if st := cc.idleStateLocked(); !st.canTakeNewRequest {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:884
		_go_fuzz_dep_.CoverTab[74660]++
												return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:885
		// _ = "end of CoverTab[74660]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:886
		_go_fuzz_dep_.CoverTab[74661]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:886
		// _ = "end of CoverTab[74661]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:886
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:886
	// _ = "end of CoverTab[74658]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:886
	_go_fuzz_dep_.CoverTab[74659]++
											cc.streamsReserved++
											return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:888
	// _ = "end of CoverTab[74659]"
}

// ClientConnState describes the state of a ClientConn.
type ClientConnState struct {
	// Closed is whether the connection is closed.
	Closed	bool

	// Closing is whether the connection is in the process of
	// closing. It may be closing due to shutdown, being a
	// single-use connection, being marked as DoNotReuse, or
	// having received a GOAWAY frame.
	Closing	bool

	// StreamsActive is how many streams are active.
	StreamsActive	int

	// StreamsReserved is how many streams have been reserved via
	// ClientConn.ReserveNewRequest.
	StreamsReserved	int

	// StreamsPending is how many requests have been sent in excess
	// of the peer's advertised MaxConcurrentStreams setting and
	// are waiting for other streams to complete.
	StreamsPending	int

	// MaxConcurrentStreams is how many concurrent streams the
	// peer advertised as acceptable. Zero means no SETTINGS
	// frame has been received yet.
	MaxConcurrentStreams	uint32

	// LastIdle, if non-zero, is when the connection last
	// transitioned to idle state.
	LastIdle	time.Time
}

// State returns a snapshot of cc's state.
func (cc *ClientConn) State() ClientConnState {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:925
	_go_fuzz_dep_.CoverTab[74662]++
											cc.wmu.Lock()
											maxConcurrent := cc.maxConcurrentStreams
											if !cc.seenSettings {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:928
		_go_fuzz_dep_.CoverTab[74664]++
												maxConcurrent = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:929
		// _ = "end of CoverTab[74664]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:930
		_go_fuzz_dep_.CoverTab[74665]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:930
		// _ = "end of CoverTab[74665]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:930
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:930
	// _ = "end of CoverTab[74662]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:930
	_go_fuzz_dep_.CoverTab[74663]++
											cc.wmu.Unlock()

											cc.mu.Lock()
											defer cc.mu.Unlock()
											return ClientConnState{
		Closed:	cc.closed,
		Closing: cc.closing || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:937
			_go_fuzz_dep_.CoverTab[74666]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:937
			return cc.singleUse
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:937
			// _ = "end of CoverTab[74666]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:937
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:937
			_go_fuzz_dep_.CoverTab[74667]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:937
			return cc.doNotReuse
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:937
			// _ = "end of CoverTab[74667]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:937
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:937
			_go_fuzz_dep_.CoverTab[74668]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:937
			return cc.goAway != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:937
			// _ = "end of CoverTab[74668]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:937
		}(),
		StreamsActive:		len(cc.streams),
		StreamsReserved:	cc.streamsReserved,
		StreamsPending:		cc.pendingRequests,
		LastIdle:		cc.lastIdle,
		MaxConcurrentStreams:	maxConcurrent,
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:943
	// _ = "end of CoverTab[74663]"
}

// clientConnIdleState describes the suitability of a client
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:946
// connection to initiate a new RoundTrip request.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:948
type clientConnIdleState struct {
	canTakeNewRequest bool
}

func (cc *ClientConn) idleState() clientConnIdleState {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:952
	_go_fuzz_dep_.CoverTab[74669]++
											cc.mu.Lock()
											defer cc.mu.Unlock()
											return cc.idleStateLocked()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:955
	// _ = "end of CoverTab[74669]"
}

func (cc *ClientConn) idleStateLocked() (st clientConnIdleState) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:958
	_go_fuzz_dep_.CoverTab[74670]++
											if cc.singleUse && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:959
		_go_fuzz_dep_.CoverTab[74673]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:959
		return cc.nextStreamID > 1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:959
		// _ = "end of CoverTab[74673]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:959
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:959
		_go_fuzz_dep_.CoverTab[74674]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:960
		// _ = "end of CoverTab[74674]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:961
		_go_fuzz_dep_.CoverTab[74675]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:961
		// _ = "end of CoverTab[74675]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:961
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:961
	// _ = "end of CoverTab[74670]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:961
	_go_fuzz_dep_.CoverTab[74671]++
											var maxConcurrentOkay bool
											if cc.t.StrictMaxConcurrentStreams {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:963
		_go_fuzz_dep_.CoverTab[74676]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:968
		maxConcurrentOkay = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:968
		// _ = "end of CoverTab[74676]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:969
		_go_fuzz_dep_.CoverTab[74677]++
												maxConcurrentOkay = int64(len(cc.streams)+cc.streamsReserved+1) <= int64(cc.maxConcurrentStreams)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:970
		// _ = "end of CoverTab[74677]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:971
	// _ = "end of CoverTab[74671]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:971
	_go_fuzz_dep_.CoverTab[74672]++

											st.canTakeNewRequest = cc.goAway == nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:973
		_go_fuzz_dep_.CoverTab[74678]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:973
		return !cc.closed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:973
		// _ = "end of CoverTab[74678]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:973
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:973
		_go_fuzz_dep_.CoverTab[74679]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:973
		return !cc.closing
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:973
		// _ = "end of CoverTab[74679]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:973
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:973
		_go_fuzz_dep_.CoverTab[74680]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:973
		return maxConcurrentOkay
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:973
		// _ = "end of CoverTab[74680]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:973
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:973
		_go_fuzz_dep_.CoverTab[74681]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:973
		return !cc.doNotReuse
												// _ = "end of CoverTab[74681]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:974
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:974
		_go_fuzz_dep_.CoverTab[74682]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:974
		return int64(cc.nextStreamID)+2*int64(cc.pendingRequests) < math.MaxInt32
												// _ = "end of CoverTab[74682]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:975
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:975
		_go_fuzz_dep_.CoverTab[74683]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:975
		return !cc.tooIdleLocked()
												// _ = "end of CoverTab[74683]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:976
	}()
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:977
	// _ = "end of CoverTab[74672]"
}

func (cc *ClientConn) canTakeNewRequestLocked() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:980
	_go_fuzz_dep_.CoverTab[74684]++
											st := cc.idleStateLocked()
											return st.canTakeNewRequest
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:982
	// _ = "end of CoverTab[74684]"
}

// tooIdleLocked reports whether this connection has been been sitting idle
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:985
// for too much wall time.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:987
func (cc *ClientConn) tooIdleLocked() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:987
	_go_fuzz_dep_.CoverTab[74685]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:992
	return cc.idleTimeout != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:992
		_go_fuzz_dep_.CoverTab[74686]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:992
		return !cc.lastIdle.IsZero()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:992
		// _ = "end of CoverTab[74686]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:992
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:992
		_go_fuzz_dep_.CoverTab[74687]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:992
		return time.Since(cc.lastIdle.Round(0)) > cc.idleTimeout
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:992
		// _ = "end of CoverTab[74687]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:992
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:992
	// _ = "end of CoverTab[74685]"
}

// onIdleTimeout is called from a time.AfterFunc goroutine. It will
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:995
// only be called when we're idle, but because we're coming from a new
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:995
// goroutine, there could be a new request coming in at the same time,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:995
// so this simply calls the synchronized closeIfIdle to shut down this
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:995
// connection. The timer could just call closeIfIdle, but this is more
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:995
// clear.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1001
func (cc *ClientConn) onIdleTimeout() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1001
	_go_fuzz_dep_.CoverTab[74688]++
											cc.closeIfIdle()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1002
	// _ = "end of CoverTab[74688]"
}

func (cc *ClientConn) closeConn() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1005
	_go_fuzz_dep_.CoverTab[74689]++
											t := time.AfterFunc(250*time.Millisecond, cc.forceCloseConn)
											defer t.Stop()
											cc.tconn.Close()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1008
	// _ = "end of CoverTab[74689]"
}

// A tls.Conn.Close can hang for a long time if the peer is unresponsive.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1011
// Try to shut it down more aggressively.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1013
func (cc *ClientConn) forceCloseConn() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1013
	_go_fuzz_dep_.CoverTab[74690]++
											tc, ok := cc.tconn.(*tls.Conn)
											if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1015
		_go_fuzz_dep_.CoverTab[74692]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1016
		// _ = "end of CoverTab[74692]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1017
		_go_fuzz_dep_.CoverTab[74693]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1017
		// _ = "end of CoverTab[74693]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1017
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1017
	// _ = "end of CoverTab[74690]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1017
	_go_fuzz_dep_.CoverTab[74691]++
											if nc := tlsUnderlyingConn(tc); nc != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1018
		_go_fuzz_dep_.CoverTab[74694]++
												nc.Close()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1019
		// _ = "end of CoverTab[74694]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1020
		_go_fuzz_dep_.CoverTab[74695]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1020
		// _ = "end of CoverTab[74695]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1020
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1020
	// _ = "end of CoverTab[74691]"
}

func (cc *ClientConn) closeIfIdle() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1023
	_go_fuzz_dep_.CoverTab[74696]++
											cc.mu.Lock()
											if len(cc.streams) > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1025
		_go_fuzz_dep_.CoverTab[74699]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1025
		return cc.streamsReserved > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1025
		// _ = "end of CoverTab[74699]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1025
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1025
		_go_fuzz_dep_.CoverTab[74700]++
												cc.mu.Unlock()
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1027
		// _ = "end of CoverTab[74700]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1028
		_go_fuzz_dep_.CoverTab[74701]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1028
		// _ = "end of CoverTab[74701]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1028
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1028
	// _ = "end of CoverTab[74696]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1028
	_go_fuzz_dep_.CoverTab[74697]++
											cc.closed = true
											nextID := cc.nextStreamID

											cc.mu.Unlock()

											if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1034
		_go_fuzz_dep_.CoverTab[74702]++
												cc.vlogf("http2: Transport closing idle conn %p (forSingleUse=%v, maxStream=%v)", cc, cc.singleUse, nextID-2)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1035
		// _ = "end of CoverTab[74702]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1036
		_go_fuzz_dep_.CoverTab[74703]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1036
		// _ = "end of CoverTab[74703]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1036
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1036
	// _ = "end of CoverTab[74697]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1036
	_go_fuzz_dep_.CoverTab[74698]++
											cc.closeConn()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1037
	// _ = "end of CoverTab[74698]"
}

func (cc *ClientConn) isDoNotReuseAndIdle() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1040
	_go_fuzz_dep_.CoverTab[74704]++
											cc.mu.Lock()
											defer cc.mu.Unlock()
											return cc.doNotReuse && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1043
		_go_fuzz_dep_.CoverTab[74705]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1043
		return len(cc.streams) == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1043
		// _ = "end of CoverTab[74705]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1043
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1043
	// _ = "end of CoverTab[74704]"
}

var shutdownEnterWaitStateHook = func() { _go_fuzz_dep_.CoverTab[74706]++; // _ = "end of CoverTab[74706]" }

// Shutdown gracefully closes the client connection, waiting for running streams to complete.
func (cc *ClientConn) Shutdown(ctx context.Context) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1049
	_go_fuzz_dep_.CoverTab[74707]++
											if err := cc.sendGoAway(); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1050
		_go_fuzz_dep_.CoverTab[74710]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1051
		// _ = "end of CoverTab[74710]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1052
		_go_fuzz_dep_.CoverTab[74711]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1052
		// _ = "end of CoverTab[74711]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1052
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1052
	// _ = "end of CoverTab[74707]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1052
	_go_fuzz_dep_.CoverTab[74708]++

											done := make(chan struct{})
											cancelled := false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1055
	_curRoutineNum68_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1055
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum68_)
											go func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1056
		_go_fuzz_dep_.CoverTab[74712]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1056
		defer func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1056
			_go_fuzz_dep_.CoverTab[74713]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1056
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum68_)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1056
			// _ = "end of CoverTab[74713]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1056
		}()
												cc.mu.Lock()
												defer cc.mu.Unlock()
												for {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1059
			_go_fuzz_dep_.CoverTab[74714]++
													if len(cc.streams) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1060
				_go_fuzz_dep_.CoverTab[74717]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1060
				return cc.closed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1060
				// _ = "end of CoverTab[74717]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1060
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1060
				_go_fuzz_dep_.CoverTab[74718]++
														cc.closed = true
														close(done)
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1063
				// _ = "end of CoverTab[74718]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1064
				_go_fuzz_dep_.CoverTab[74719]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1064
				// _ = "end of CoverTab[74719]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1064
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1064
			// _ = "end of CoverTab[74714]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1064
			_go_fuzz_dep_.CoverTab[74715]++
													if cancelled {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1065
				_go_fuzz_dep_.CoverTab[74720]++
														break
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1066
				// _ = "end of CoverTab[74720]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1067
				_go_fuzz_dep_.CoverTab[74721]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1067
				// _ = "end of CoverTab[74721]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1067
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1067
			// _ = "end of CoverTab[74715]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1067
			_go_fuzz_dep_.CoverTab[74716]++
													cc.cond.Wait()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1068
			// _ = "end of CoverTab[74716]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1069
		// _ = "end of CoverTab[74712]"
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1070
	// _ = "end of CoverTab[74708]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1070
	_go_fuzz_dep_.CoverTab[74709]++
											shutdownEnterWaitStateHook()
											select {
	case <-done:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1073
		_go_fuzz_dep_.CoverTab[74722]++
												cc.closeConn()
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1075
		// _ = "end of CoverTab[74722]"
	case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1076
		_go_fuzz_dep_.CoverTab[74723]++
												cc.mu.Lock()

												cancelled = true
												cc.cond.Broadcast()
												cc.mu.Unlock()
												return ctx.Err()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1082
		// _ = "end of CoverTab[74723]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1083
	// _ = "end of CoverTab[74709]"
}

func (cc *ClientConn) sendGoAway() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1086
	_go_fuzz_dep_.CoverTab[74724]++
											cc.mu.Lock()
											closing := cc.closing
											cc.closing = true
											maxStreamID := cc.nextStreamID
											cc.mu.Unlock()
											if closing {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1092
		_go_fuzz_dep_.CoverTab[74728]++

												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1094
		// _ = "end of CoverTab[74728]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1095
		_go_fuzz_dep_.CoverTab[74729]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1095
		// _ = "end of CoverTab[74729]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1095
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1095
	// _ = "end of CoverTab[74724]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1095
	_go_fuzz_dep_.CoverTab[74725]++

											cc.wmu.Lock()
											defer cc.wmu.Unlock()

											if err := cc.fr.WriteGoAway(maxStreamID, ErrCodeNo, nil); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1100
		_go_fuzz_dep_.CoverTab[74730]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1101
		// _ = "end of CoverTab[74730]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1102
		_go_fuzz_dep_.CoverTab[74731]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1102
		// _ = "end of CoverTab[74731]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1102
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1102
	// _ = "end of CoverTab[74725]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1102
	_go_fuzz_dep_.CoverTab[74726]++
											if err := cc.bw.Flush(); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1103
		_go_fuzz_dep_.CoverTab[74732]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1104
		// _ = "end of CoverTab[74732]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1105
		_go_fuzz_dep_.CoverTab[74733]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1105
		// _ = "end of CoverTab[74733]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1105
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1105
	// _ = "end of CoverTab[74726]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1105
	_go_fuzz_dep_.CoverTab[74727]++

											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1107
	// _ = "end of CoverTab[74727]"
}

// closes the client connection immediately. In-flight requests are interrupted.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1110
// err is sent to streams.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1112
func (cc *ClientConn) closeForError(err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1112
	_go_fuzz_dep_.CoverTab[74734]++
											cc.mu.Lock()
											cc.closed = true
											for _, cs := range cc.streams {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1115
		_go_fuzz_dep_.CoverTab[74736]++
												cs.abortStreamLocked(err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1116
		// _ = "end of CoverTab[74736]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1117
	// _ = "end of CoverTab[74734]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1117
	_go_fuzz_dep_.CoverTab[74735]++
											cc.cond.Broadcast()
											cc.mu.Unlock()
											cc.closeConn()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1120
	// _ = "end of CoverTab[74735]"
}

// Close closes the client connection immediately.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1123
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1123
// In-flight requests are interrupted. For a graceful shutdown, use Shutdown instead.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1126
func (cc *ClientConn) Close() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1126
	_go_fuzz_dep_.CoverTab[74737]++
											err := errors.New("http2: client connection force closed via ClientConn.Close")
											cc.closeForError(err)
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1129
	// _ = "end of CoverTab[74737]"
}

// closes the client connection immediately. In-flight requests are interrupted.
func (cc *ClientConn) closeForLostPing() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1133
	_go_fuzz_dep_.CoverTab[74738]++
											err := errors.New("http2: client connection lost")
											if f := cc.t.CountError; f != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1135
		_go_fuzz_dep_.CoverTab[74740]++
												f("conn_close_lost_ping")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1136
		// _ = "end of CoverTab[74740]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1137
		_go_fuzz_dep_.CoverTab[74741]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1137
		// _ = "end of CoverTab[74741]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1137
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1137
	// _ = "end of CoverTab[74738]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1137
	_go_fuzz_dep_.CoverTab[74739]++
											cc.closeForError(err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1138
	// _ = "end of CoverTab[74739]"
}

// errRequestCanceled is a copy of net/http's errRequestCanceled because it's not
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1141
// exported. At least they'll be DeepEqual for h1-vs-h2 comparisons tests.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1143
var errRequestCanceled = errors.New("net/http: request canceled")

func commaSeparatedTrailers(req *http.Request) (string, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1145
	_go_fuzz_dep_.CoverTab[74742]++
											keys := make([]string, 0, len(req.Trailer))
											for k := range req.Trailer {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1147
		_go_fuzz_dep_.CoverTab[74745]++
												k = canonicalHeader(k)
												switch k {
		case "Transfer-Encoding", "Trailer", "Content-Length":
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1150
			_go_fuzz_dep_.CoverTab[74747]++
													return "", fmt.Errorf("invalid Trailer key %q", k)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1151
			// _ = "end of CoverTab[74747]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1151
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1151
			_go_fuzz_dep_.CoverTab[74748]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1151
			// _ = "end of CoverTab[74748]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1152
		// _ = "end of CoverTab[74745]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1152
		_go_fuzz_dep_.CoverTab[74746]++
												keys = append(keys, k)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1153
		// _ = "end of CoverTab[74746]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1154
	// _ = "end of CoverTab[74742]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1154
	_go_fuzz_dep_.CoverTab[74743]++
											if len(keys) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1155
		_go_fuzz_dep_.CoverTab[74749]++
												sort.Strings(keys)
												return strings.Join(keys, ","), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1157
		// _ = "end of CoverTab[74749]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1158
		_go_fuzz_dep_.CoverTab[74750]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1158
		// _ = "end of CoverTab[74750]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1158
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1158
	// _ = "end of CoverTab[74743]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1158
	_go_fuzz_dep_.CoverTab[74744]++
											return "", nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1159
	// _ = "end of CoverTab[74744]"
}

func (cc *ClientConn) responseHeaderTimeout() time.Duration {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1162
	_go_fuzz_dep_.CoverTab[74751]++
											if cc.t.t1 != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1163
		_go_fuzz_dep_.CoverTab[74753]++
												return cc.t.t1.ResponseHeaderTimeout
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1164
		// _ = "end of CoverTab[74753]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1165
		_go_fuzz_dep_.CoverTab[74754]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1165
		// _ = "end of CoverTab[74754]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1165
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1165
	// _ = "end of CoverTab[74751]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1165
	_go_fuzz_dep_.CoverTab[74752]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1170
	return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1170
	// _ = "end of CoverTab[74752]"
}

// checkConnHeaders checks whether req has any invalid connection-level headers.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1173
// per RFC 7540 section 8.1.2.2: Connection-Specific Header Fields.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1173
// Certain headers are special-cased as okay but not transmitted later.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1176
func checkConnHeaders(req *http.Request) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1176
	_go_fuzz_dep_.CoverTab[74755]++
											if v := req.Header.Get("Upgrade"); v != "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1177
		_go_fuzz_dep_.CoverTab[74759]++
												return fmt.Errorf("http2: invalid Upgrade request header: %q", req.Header["Upgrade"])
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1178
		// _ = "end of CoverTab[74759]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1179
		_go_fuzz_dep_.CoverTab[74760]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1179
		// _ = "end of CoverTab[74760]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1179
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1179
	// _ = "end of CoverTab[74755]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1179
	_go_fuzz_dep_.CoverTab[74756]++
											if vv := req.Header["Transfer-Encoding"]; len(vv) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1180
		_go_fuzz_dep_.CoverTab[74761]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1180
		return (len(vv) > 1 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1180
			_go_fuzz_dep_.CoverTab[74762]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1180
			return vv[0] != "" && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1180
				_go_fuzz_dep_.CoverTab[74763]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1180
				return vv[0] != "chunked"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1180
				// _ = "end of CoverTab[74763]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1180
			}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1180
			// _ = "end of CoverTab[74762]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1180
		}())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1180
		// _ = "end of CoverTab[74761]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1180
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1180
		_go_fuzz_dep_.CoverTab[74764]++
												return fmt.Errorf("http2: invalid Transfer-Encoding request header: %q", vv)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1181
		// _ = "end of CoverTab[74764]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1182
		_go_fuzz_dep_.CoverTab[74765]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1182
		// _ = "end of CoverTab[74765]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1182
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1182
	// _ = "end of CoverTab[74756]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1182
	_go_fuzz_dep_.CoverTab[74757]++
											if vv := req.Header["Connection"]; len(vv) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
		_go_fuzz_dep_.CoverTab[74766]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
		return (len(vv) > 1 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
			_go_fuzz_dep_.CoverTab[74767]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
			return vv[0] != "" && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
				_go_fuzz_dep_.CoverTab[74768]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
				return !asciiEqualFold(vv[0], "close")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
				// _ = "end of CoverTab[74768]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
				_go_fuzz_dep_.CoverTab[74769]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
				return !asciiEqualFold(vv[0], "keep-alive")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
				// _ = "end of CoverTab[74769]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
			}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
			// _ = "end of CoverTab[74767]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
		}())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
		// _ = "end of CoverTab[74766]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1183
		_go_fuzz_dep_.CoverTab[74770]++
												return fmt.Errorf("http2: invalid Connection request header: %q", vv)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1184
		// _ = "end of CoverTab[74770]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1185
		_go_fuzz_dep_.CoverTab[74771]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1185
		// _ = "end of CoverTab[74771]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1185
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1185
	// _ = "end of CoverTab[74757]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1185
	_go_fuzz_dep_.CoverTab[74758]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1186
	// _ = "end of CoverTab[74758]"
}

// actualContentLength returns a sanitized version of
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1189
// req.ContentLength, where 0 actually means zero (not unknown) and -1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1189
// means unknown.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1192
func actualContentLength(req *http.Request) int64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1192
	_go_fuzz_dep_.CoverTab[74772]++
											if req.Body == nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1193
		_go_fuzz_dep_.CoverTab[74775]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1193
		return req.Body == http.NoBody
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1193
		// _ = "end of CoverTab[74775]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1193
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1193
		_go_fuzz_dep_.CoverTab[74776]++
												return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1194
		// _ = "end of CoverTab[74776]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1195
		_go_fuzz_dep_.CoverTab[74777]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1195
		// _ = "end of CoverTab[74777]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1195
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1195
	// _ = "end of CoverTab[74772]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1195
	_go_fuzz_dep_.CoverTab[74773]++
											if req.ContentLength != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1196
		_go_fuzz_dep_.CoverTab[74778]++
												return req.ContentLength
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1197
		// _ = "end of CoverTab[74778]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1198
		_go_fuzz_dep_.CoverTab[74779]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1198
		// _ = "end of CoverTab[74779]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1198
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1198
	// _ = "end of CoverTab[74773]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1198
	_go_fuzz_dep_.CoverTab[74774]++
											return -1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1199
	// _ = "end of CoverTab[74774]"
}

func (cc *ClientConn) decrStreamReservations() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1202
	_go_fuzz_dep_.CoverTab[74780]++
											cc.mu.Lock()
											defer cc.mu.Unlock()
											cc.decrStreamReservationsLocked()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1205
	// _ = "end of CoverTab[74780]"
}

func (cc *ClientConn) decrStreamReservationsLocked() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1208
	_go_fuzz_dep_.CoverTab[74781]++
											if cc.streamsReserved > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1209
		_go_fuzz_dep_.CoverTab[74782]++
												cc.streamsReserved--
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1210
		// _ = "end of CoverTab[74782]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1211
		_go_fuzz_dep_.CoverTab[74783]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1211
		// _ = "end of CoverTab[74783]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1211
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1211
	// _ = "end of CoverTab[74781]"
}

func (cc *ClientConn) RoundTrip(req *http.Request) (*http.Response, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1214
	_go_fuzz_dep_.CoverTab[74784]++
											ctx := req.Context()
											cs := &clientStream{
		cc:			cc,
		ctx:			ctx,
		reqCancel:		req.Cancel,
		isHead:			req.Method == "HEAD",
		reqBody:		req.Body,
		reqBodyContentLength:	actualContentLength(req),
		trace:			httptrace.ContextClientTrace(ctx),
		peerClosed:		make(chan struct{}),
		abort:			make(chan struct{}),
		respHeaderRecv:		make(chan struct{}),
		donec:			make(chan struct{}),
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1228
	_curRoutineNum69_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1228
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum69_)
											go cs.doRequest(req)

											waitDone := func() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1231
		_go_fuzz_dep_.CoverTab[74788]++
												select {
		case <-cs.donec:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1233
			_go_fuzz_dep_.CoverTab[74789]++
													return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1234
			// _ = "end of CoverTab[74789]"
		case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1235
			_go_fuzz_dep_.CoverTab[74790]++
													return ctx.Err()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1236
			// _ = "end of CoverTab[74790]"
		case <-cs.reqCancel:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1237
			_go_fuzz_dep_.CoverTab[74791]++
													return errRequestCanceled
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1238
			// _ = "end of CoverTab[74791]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1239
		// _ = "end of CoverTab[74788]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1240
	// _ = "end of CoverTab[74784]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1240
	_go_fuzz_dep_.CoverTab[74785]++

											handleResponseHeaders := func() (*http.Response, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1242
		_go_fuzz_dep_.CoverTab[74792]++
												res := cs.res
												if res.StatusCode > 299 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1244
			_go_fuzz_dep_.CoverTab[74795]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1254
			cs.abortRequestBodyWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1254
			// _ = "end of CoverTab[74795]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1255
			_go_fuzz_dep_.CoverTab[74796]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1255
			// _ = "end of CoverTab[74796]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1255
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1255
		// _ = "end of CoverTab[74792]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1255
		_go_fuzz_dep_.CoverTab[74793]++
												res.Request = req
												res.TLS = cc.tlsState
												if res.Body == noBody && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1258
			_go_fuzz_dep_.CoverTab[74797]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1258
			return actualContentLength(req) == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1258
			// _ = "end of CoverTab[74797]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1258
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1258
			_go_fuzz_dep_.CoverTab[74798]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1262
			if err := waitDone(); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1262
				_go_fuzz_dep_.CoverTab[74799]++
														return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1263
				// _ = "end of CoverTab[74799]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1264
				_go_fuzz_dep_.CoverTab[74800]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1264
				// _ = "end of CoverTab[74800]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1264
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1264
			// _ = "end of CoverTab[74798]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1265
			_go_fuzz_dep_.CoverTab[74801]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1265
			// _ = "end of CoverTab[74801]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1265
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1265
		// _ = "end of CoverTab[74793]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1265
		_go_fuzz_dep_.CoverTab[74794]++
												return res, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1266
		// _ = "end of CoverTab[74794]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1267
	// _ = "end of CoverTab[74785]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1267
	_go_fuzz_dep_.CoverTab[74786]++

											cancelRequest := func(cs *clientStream, err error) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1269
		_go_fuzz_dep_.CoverTab[74802]++
												cs.cc.mu.Lock()
												defer cs.cc.mu.Unlock()
												cs.abortStreamLocked(err)
												if cs.ID != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1273
			_go_fuzz_dep_.CoverTab[74804]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1285
			cs.cc.doNotReuse = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1285
			// _ = "end of CoverTab[74804]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1286
			_go_fuzz_dep_.CoverTab[74805]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1286
			// _ = "end of CoverTab[74805]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1286
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1286
		// _ = "end of CoverTab[74802]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1286
		_go_fuzz_dep_.CoverTab[74803]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1287
		// _ = "end of CoverTab[74803]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1288
	// _ = "end of CoverTab[74786]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1288
	_go_fuzz_dep_.CoverTab[74787]++

											for {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1290
		_go_fuzz_dep_.CoverTab[74806]++
												select {
		case <-cs.respHeaderRecv:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1292
			_go_fuzz_dep_.CoverTab[74807]++
													return handleResponseHeaders()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1293
			// _ = "end of CoverTab[74807]"
		case <-cs.abort:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1294
			_go_fuzz_dep_.CoverTab[74808]++
													select {
			case <-cs.respHeaderRecv:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1296
				_go_fuzz_dep_.CoverTab[74811]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1301
				return handleResponseHeaders()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1301
				// _ = "end of CoverTab[74811]"
			default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1302
				_go_fuzz_dep_.CoverTab[74812]++
														waitDone()
														return nil, cancelRequest(cs, cs.abortErr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1304
				// _ = "end of CoverTab[74812]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1305
			// _ = "end of CoverTab[74808]"
		case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1306
			_go_fuzz_dep_.CoverTab[74809]++
													return nil, cancelRequest(cs, ctx.Err())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1307
			// _ = "end of CoverTab[74809]"
		case <-cs.reqCancel:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1308
			_go_fuzz_dep_.CoverTab[74810]++
													return nil, cancelRequest(cs, errRequestCanceled)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1309
			// _ = "end of CoverTab[74810]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1310
		// _ = "end of CoverTab[74806]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1311
	// _ = "end of CoverTab[74787]"
}

// doRequest runs for the duration of the request lifetime.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1314
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1314
// It sends the request and performs post-request cleanup (closing Request.Body, etc.).
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1317
func (cs *clientStream) doRequest(req *http.Request) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1317
	_go_fuzz_dep_.CoverTab[74813]++
											err := cs.writeRequest(req)
											cs.cleanupWriteRequest(err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1319
	// _ = "end of CoverTab[74813]"
}

// writeRequest sends a request.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1322
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1322
// It returns nil after the request is written, the response read,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1322
// and the request stream is half-closed by the peer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1322
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1322
// It returns non-nil if the request ends otherwise.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1322
// If the returned error is StreamError, the error Code may be used in resetting the stream.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1329
func (cs *clientStream) writeRequest(req *http.Request) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1329
	_go_fuzz_dep_.CoverTab[74814]++
											cc := cs.cc
											ctx := cs.ctx

											if err := checkConnHeaders(req); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1333
		_go_fuzz_dep_.CoverTab[74826]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1334
		// _ = "end of CoverTab[74826]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1335
		_go_fuzz_dep_.CoverTab[74827]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1335
		// _ = "end of CoverTab[74827]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1335
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1335
	// _ = "end of CoverTab[74814]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1335
	_go_fuzz_dep_.CoverTab[74815]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1340
	if cc.reqHeaderMu == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1340
		_go_fuzz_dep_.CoverTab[74828]++
												panic("RoundTrip on uninitialized ClientConn")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1341
		// _ = "end of CoverTab[74828]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1342
		_go_fuzz_dep_.CoverTab[74829]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1342
		// _ = "end of CoverTab[74829]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1342
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1342
	// _ = "end of CoverTab[74815]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1342
	_go_fuzz_dep_.CoverTab[74816]++
											select {
	case cc.reqHeaderMu <- struct{}{}:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1344
		_go_fuzz_dep_.CoverTab[74830]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1344
		// _ = "end of CoverTab[74830]"
	case <-cs.reqCancel:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1345
		_go_fuzz_dep_.CoverTab[74831]++
												return errRequestCanceled
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1346
		// _ = "end of CoverTab[74831]"
	case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1347
		_go_fuzz_dep_.CoverTab[74832]++
												return ctx.Err()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1348
		// _ = "end of CoverTab[74832]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1349
	// _ = "end of CoverTab[74816]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1349
	_go_fuzz_dep_.CoverTab[74817]++

											cc.mu.Lock()
											if cc.idleTimer != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1352
		_go_fuzz_dep_.CoverTab[74833]++
												cc.idleTimer.Stop()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1353
		// _ = "end of CoverTab[74833]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1354
		_go_fuzz_dep_.CoverTab[74834]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1354
		// _ = "end of CoverTab[74834]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1354
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1354
	// _ = "end of CoverTab[74817]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1354
	_go_fuzz_dep_.CoverTab[74818]++
											cc.decrStreamReservationsLocked()
											if err := cc.awaitOpenSlotForStreamLocked(cs); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1356
		_go_fuzz_dep_.CoverTab[74835]++
												cc.mu.Unlock()
												<-cc.reqHeaderMu
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1359
		// _ = "end of CoverTab[74835]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1360
		_go_fuzz_dep_.CoverTab[74836]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1360
		// _ = "end of CoverTab[74836]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1360
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1360
	// _ = "end of CoverTab[74818]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1360
	_go_fuzz_dep_.CoverTab[74819]++
											cc.addStreamLocked(cs)
											if isConnectionCloseRequest(req) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1362
		_go_fuzz_dep_.CoverTab[74837]++
												cc.doNotReuse = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1363
		// _ = "end of CoverTab[74837]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1364
		_go_fuzz_dep_.CoverTab[74838]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1364
		// _ = "end of CoverTab[74838]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1364
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1364
	// _ = "end of CoverTab[74819]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1364
	_go_fuzz_dep_.CoverTab[74820]++
											cc.mu.Unlock()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1368
	if !cc.t.disableCompression() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1368
		_go_fuzz_dep_.CoverTab[74839]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1368
		return req.Header.Get("Accept-Encoding") == ""
												// _ = "end of CoverTab[74839]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1369
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1369
		_go_fuzz_dep_.CoverTab[74840]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1369
		return req.Header.Get("Range") == ""
												// _ = "end of CoverTab[74840]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1370
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1370
		_go_fuzz_dep_.CoverTab[74841]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1370
		return !cs.isHead
												// _ = "end of CoverTab[74841]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1371
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1371
		_go_fuzz_dep_.CoverTab[74842]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1384
		cs.requestedGzip = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1384
		// _ = "end of CoverTab[74842]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1385
		_go_fuzz_dep_.CoverTab[74843]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1385
		// _ = "end of CoverTab[74843]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1385
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1385
	// _ = "end of CoverTab[74820]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1385
	_go_fuzz_dep_.CoverTab[74821]++

											continueTimeout := cc.t.expectContinueTimeout()
											if continueTimeout != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1388
		_go_fuzz_dep_.CoverTab[74844]++
												if !httpguts.HeaderValuesContainsToken(req.Header["Expect"], "100-continue") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1389
			_go_fuzz_dep_.CoverTab[74845]++
													continueTimeout = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1390
			// _ = "end of CoverTab[74845]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1391
			_go_fuzz_dep_.CoverTab[74846]++
													cs.on100 = make(chan struct{}, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1392
			// _ = "end of CoverTab[74846]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1393
		// _ = "end of CoverTab[74844]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1394
		_go_fuzz_dep_.CoverTab[74847]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1394
		// _ = "end of CoverTab[74847]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1394
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1394
	// _ = "end of CoverTab[74821]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1394
	_go_fuzz_dep_.CoverTab[74822]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1400
	err = cs.encodeAndWriteHeaders(req)
	<-cc.reqHeaderMu
	if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1402
		_go_fuzz_dep_.CoverTab[74848]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1403
		// _ = "end of CoverTab[74848]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1404
		_go_fuzz_dep_.CoverTab[74849]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1404
		// _ = "end of CoverTab[74849]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1404
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1404
	// _ = "end of CoverTab[74822]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1404
	_go_fuzz_dep_.CoverTab[74823]++

											hasBody := cs.reqBodyContentLength != 0
											if !hasBody {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1407
		_go_fuzz_dep_.CoverTab[74850]++
												cs.sentEndStream = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1408
		// _ = "end of CoverTab[74850]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1409
		_go_fuzz_dep_.CoverTab[74851]++
												if continueTimeout != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1410
			_go_fuzz_dep_.CoverTab[74853]++
													traceWait100Continue(cs.trace)
													timer := time.NewTimer(continueTimeout)
													select {
			case <-timer.C:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1414
				_go_fuzz_dep_.CoverTab[74855]++
														err = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1415
				// _ = "end of CoverTab[74855]"
			case <-cs.on100:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1416
				_go_fuzz_dep_.CoverTab[74856]++
														err = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1417
				// _ = "end of CoverTab[74856]"
			case <-cs.abort:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1418
				_go_fuzz_dep_.CoverTab[74857]++
														err = cs.abortErr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1419
				// _ = "end of CoverTab[74857]"
			case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1420
				_go_fuzz_dep_.CoverTab[74858]++
														err = ctx.Err()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1421
				// _ = "end of CoverTab[74858]"
			case <-cs.reqCancel:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1422
				_go_fuzz_dep_.CoverTab[74859]++
														err = errRequestCanceled
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1423
				// _ = "end of CoverTab[74859]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1424
			// _ = "end of CoverTab[74853]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1424
			_go_fuzz_dep_.CoverTab[74854]++
													timer.Stop()
													if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1426
				_go_fuzz_dep_.CoverTab[74860]++
														traceWroteRequest(cs.trace, err)
														return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1428
				// _ = "end of CoverTab[74860]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1429
				_go_fuzz_dep_.CoverTab[74861]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1429
				// _ = "end of CoverTab[74861]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1429
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1429
			// _ = "end of CoverTab[74854]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1430
			_go_fuzz_dep_.CoverTab[74862]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1430
			// _ = "end of CoverTab[74862]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1430
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1430
		// _ = "end of CoverTab[74851]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1430
		_go_fuzz_dep_.CoverTab[74852]++

												if err = cs.writeRequestBody(req); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1432
			_go_fuzz_dep_.CoverTab[74863]++
													if err != errStopReqBodyWrite {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1433
				_go_fuzz_dep_.CoverTab[74864]++
														traceWroteRequest(cs.trace, err)
														return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1435
				// _ = "end of CoverTab[74864]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1436
				_go_fuzz_dep_.CoverTab[74865]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1436
				// _ = "end of CoverTab[74865]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1436
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1436
			// _ = "end of CoverTab[74863]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1437
			_go_fuzz_dep_.CoverTab[74866]++
													cs.sentEndStream = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1438
			// _ = "end of CoverTab[74866]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1439
		// _ = "end of CoverTab[74852]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1440
	// _ = "end of CoverTab[74823]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1440
	_go_fuzz_dep_.CoverTab[74824]++

											traceWroteRequest(cs.trace, err)

											var respHeaderTimer <-chan time.Time
											var respHeaderRecv chan struct{}
											if d := cc.responseHeaderTimeout(); d != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1446
		_go_fuzz_dep_.CoverTab[74867]++
												timer := time.NewTimer(d)
												defer timer.Stop()
												respHeaderTimer = timer.C
												respHeaderRecv = cs.respHeaderRecv
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1450
		// _ = "end of CoverTab[74867]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1451
		_go_fuzz_dep_.CoverTab[74868]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1451
		// _ = "end of CoverTab[74868]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1451
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1451
	// _ = "end of CoverTab[74824]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1451
	_go_fuzz_dep_.CoverTab[74825]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1455
	for {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1455
		_go_fuzz_dep_.CoverTab[74869]++
												select {
		case <-cs.peerClosed:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1457
			_go_fuzz_dep_.CoverTab[74870]++
													return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1458
			// _ = "end of CoverTab[74870]"
		case <-respHeaderTimer:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1459
			_go_fuzz_dep_.CoverTab[74871]++
													return errTimeout
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1460
			// _ = "end of CoverTab[74871]"
		case <-respHeaderRecv:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1461
			_go_fuzz_dep_.CoverTab[74872]++
													respHeaderRecv = nil
													respHeaderTimer = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1463
			// _ = "end of CoverTab[74872]"
		case <-cs.abort:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1464
			_go_fuzz_dep_.CoverTab[74873]++
													return cs.abortErr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1465
			// _ = "end of CoverTab[74873]"
		case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1466
			_go_fuzz_dep_.CoverTab[74874]++
													return ctx.Err()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1467
			// _ = "end of CoverTab[74874]"
		case <-cs.reqCancel:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1468
			_go_fuzz_dep_.CoverTab[74875]++
													return errRequestCanceled
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1469
			// _ = "end of CoverTab[74875]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1470
		// _ = "end of CoverTab[74869]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1471
	// _ = "end of CoverTab[74825]"
}

func (cs *clientStream) encodeAndWriteHeaders(req *http.Request) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1474
	_go_fuzz_dep_.CoverTab[74876]++
											cc := cs.cc
											ctx := cs.ctx

											cc.wmu.Lock()
											defer cc.wmu.Unlock()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1482
	select {
	case <-cs.abort:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1483
		_go_fuzz_dep_.CoverTab[74880]++
												return cs.abortErr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1484
		// _ = "end of CoverTab[74880]"
	case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1485
		_go_fuzz_dep_.CoverTab[74881]++
												return ctx.Err()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1486
		// _ = "end of CoverTab[74881]"
	case <-cs.reqCancel:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1487
		_go_fuzz_dep_.CoverTab[74882]++
												return errRequestCanceled
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1488
		// _ = "end of CoverTab[74882]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1489
		_go_fuzz_dep_.CoverTab[74883]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1489
		// _ = "end of CoverTab[74883]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1490
	// _ = "end of CoverTab[74876]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1490
	_go_fuzz_dep_.CoverTab[74877]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1497
	trailers, err := commaSeparatedTrailers(req)
	if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1498
		_go_fuzz_dep_.CoverTab[74884]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1499
		// _ = "end of CoverTab[74884]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1500
		_go_fuzz_dep_.CoverTab[74885]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1500
		// _ = "end of CoverTab[74885]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1500
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1500
	// _ = "end of CoverTab[74877]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1500
	_go_fuzz_dep_.CoverTab[74878]++
											hasTrailers := trailers != ""
											contentLen := actualContentLength(req)
											hasBody := contentLen != 0
											hdrs, err := cc.encodeHeaders(req, cs.requestedGzip, trailers, contentLen)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1505
		_go_fuzz_dep_.CoverTab[74886]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1506
		// _ = "end of CoverTab[74886]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1507
		_go_fuzz_dep_.CoverTab[74887]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1507
		// _ = "end of CoverTab[74887]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1507
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1507
	// _ = "end of CoverTab[74878]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1507
	_go_fuzz_dep_.CoverTab[74879]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1510
	endStream := !hasBody && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1510
		_go_fuzz_dep_.CoverTab[74888]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1510
		return !hasTrailers
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1510
		// _ = "end of CoverTab[74888]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1510
	}()
											cs.sentHeaders = true
											err = cc.writeHeaders(cs.ID, endStream, int(cc.maxFrameSize), hdrs)
											traceWroteHeaders(cs.trace)
											return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1514
	// _ = "end of CoverTab[74879]"
}

// cleanupWriteRequest performs post-request tasks.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1517
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1517
// If err (the result of writeRequest) is non-nil and the stream is not closed,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1517
// cleanupWriteRequest will send a reset to the peer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1521
func (cs *clientStream) cleanupWriteRequest(err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1521
	_go_fuzz_dep_.CoverTab[74889]++
											cc := cs.cc

											if cs.ID == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1524
		_go_fuzz_dep_.CoverTab[74898]++

												cc.decrStreamReservations()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1526
		// _ = "end of CoverTab[74898]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1527
		_go_fuzz_dep_.CoverTab[74899]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1527
		// _ = "end of CoverTab[74899]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1527
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1527
	// _ = "end of CoverTab[74889]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1527
	_go_fuzz_dep_.CoverTab[74890]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1533
	cc.mu.Lock()
	mustCloseBody := false
	if cs.reqBody != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1535
		_go_fuzz_dep_.CoverTab[74900]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1535
		return cs.reqBodyClosed == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1535
		// _ = "end of CoverTab[74900]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1535
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1535
		_go_fuzz_dep_.CoverTab[74901]++
												mustCloseBody = true
												cs.reqBodyClosed = make(chan struct{})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1537
		// _ = "end of CoverTab[74901]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1538
		_go_fuzz_dep_.CoverTab[74902]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1538
		// _ = "end of CoverTab[74902]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1538
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1538
	// _ = "end of CoverTab[74890]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1538
	_go_fuzz_dep_.CoverTab[74891]++
											bodyClosed := cs.reqBodyClosed
											cc.mu.Unlock()
											if mustCloseBody {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1541
		_go_fuzz_dep_.CoverTab[74903]++
												cs.reqBody.Close()
												close(bodyClosed)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1543
		// _ = "end of CoverTab[74903]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1544
		_go_fuzz_dep_.CoverTab[74904]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1544
		// _ = "end of CoverTab[74904]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1544
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1544
	// _ = "end of CoverTab[74891]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1544
	_go_fuzz_dep_.CoverTab[74892]++
											if bodyClosed != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1545
		_go_fuzz_dep_.CoverTab[74905]++
												<-bodyClosed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1546
		// _ = "end of CoverTab[74905]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1547
		_go_fuzz_dep_.CoverTab[74906]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1547
		// _ = "end of CoverTab[74906]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1547
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1547
	// _ = "end of CoverTab[74892]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1547
	_go_fuzz_dep_.CoverTab[74893]++

											if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1549
		_go_fuzz_dep_.CoverTab[74907]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1549
		return cs.sentEndStream
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1549
		// _ = "end of CoverTab[74907]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1549
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1549
		_go_fuzz_dep_.CoverTab[74908]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1553
		select {
		case <-cs.peerClosed:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1554
			_go_fuzz_dep_.CoverTab[74909]++
													err = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1555
			// _ = "end of CoverTab[74909]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1556
			_go_fuzz_dep_.CoverTab[74910]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1556
			// _ = "end of CoverTab[74910]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1557
		// _ = "end of CoverTab[74908]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1558
		_go_fuzz_dep_.CoverTab[74911]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1558
		// _ = "end of CoverTab[74911]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1558
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1558
	// _ = "end of CoverTab[74893]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1558
	_go_fuzz_dep_.CoverTab[74894]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1559
		_go_fuzz_dep_.CoverTab[74912]++
												cs.abortStream(err)
												if cs.sentHeaders {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1561
			_go_fuzz_dep_.CoverTab[74914]++
													if se, ok := err.(StreamError); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1562
				_go_fuzz_dep_.CoverTab[74915]++
														if se.Cause != errFromPeer {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1563
					_go_fuzz_dep_.CoverTab[74916]++
															cc.writeStreamReset(cs.ID, se.Code, err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1564
					// _ = "end of CoverTab[74916]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1565
					_go_fuzz_dep_.CoverTab[74917]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1565
					// _ = "end of CoverTab[74917]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1565
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1565
				// _ = "end of CoverTab[74915]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1566
				_go_fuzz_dep_.CoverTab[74918]++
														cc.writeStreamReset(cs.ID, ErrCodeCancel, err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1567
				// _ = "end of CoverTab[74918]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1568
			// _ = "end of CoverTab[74914]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1569
			_go_fuzz_dep_.CoverTab[74919]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1569
			// _ = "end of CoverTab[74919]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1569
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1569
		// _ = "end of CoverTab[74912]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1569
		_go_fuzz_dep_.CoverTab[74913]++
												cs.bufPipe.CloseWithError(err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1570
		// _ = "end of CoverTab[74913]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1571
		_go_fuzz_dep_.CoverTab[74920]++
												if cs.sentHeaders && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1572
			_go_fuzz_dep_.CoverTab[74922]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1572
			return !cs.sentEndStream
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1572
			// _ = "end of CoverTab[74922]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1572
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1572
			_go_fuzz_dep_.CoverTab[74923]++
													cc.writeStreamReset(cs.ID, ErrCodeNo, nil)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1573
			// _ = "end of CoverTab[74923]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1574
			_go_fuzz_dep_.CoverTab[74924]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1574
			// _ = "end of CoverTab[74924]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1574
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1574
		// _ = "end of CoverTab[74920]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1574
		_go_fuzz_dep_.CoverTab[74921]++
												cs.bufPipe.CloseWithError(errRequestCanceled)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1575
		// _ = "end of CoverTab[74921]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1576
	// _ = "end of CoverTab[74894]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1576
	_go_fuzz_dep_.CoverTab[74895]++
											if cs.ID != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1577
		_go_fuzz_dep_.CoverTab[74925]++
												cc.forgetStreamID(cs.ID)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1578
		// _ = "end of CoverTab[74925]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1579
		_go_fuzz_dep_.CoverTab[74926]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1579
		// _ = "end of CoverTab[74926]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1579
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1579
	// _ = "end of CoverTab[74895]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1579
	_go_fuzz_dep_.CoverTab[74896]++

											cc.wmu.Lock()
											werr := cc.werr
											cc.wmu.Unlock()
											if werr != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1584
		_go_fuzz_dep_.CoverTab[74927]++
												cc.Close()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1585
		// _ = "end of CoverTab[74927]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1586
		_go_fuzz_dep_.CoverTab[74928]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1586
		// _ = "end of CoverTab[74928]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1586
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1586
	// _ = "end of CoverTab[74896]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1586
	_go_fuzz_dep_.CoverTab[74897]++

											close(cs.donec)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1588
	// _ = "end of CoverTab[74897]"
}

// awaitOpenSlotForStreamLocked waits until len(streams) < maxConcurrentStreams.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1591
// Must hold cc.mu.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1593
func (cc *ClientConn) awaitOpenSlotForStreamLocked(cs *clientStream) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1593
	_go_fuzz_dep_.CoverTab[74929]++
											for {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1594
		_go_fuzz_dep_.CoverTab[74930]++
												cc.lastActive = time.Now()
												if cc.closed || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1596
			_go_fuzz_dep_.CoverTab[74933]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1596
			return !cc.canTakeNewRequestLocked()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1596
			// _ = "end of CoverTab[74933]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1596
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1596
			_go_fuzz_dep_.CoverTab[74934]++
													return errClientConnUnusable
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1597
			// _ = "end of CoverTab[74934]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1598
			_go_fuzz_dep_.CoverTab[74935]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1598
			// _ = "end of CoverTab[74935]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1598
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1598
		// _ = "end of CoverTab[74930]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1598
		_go_fuzz_dep_.CoverTab[74931]++
												cc.lastIdle = time.Time{}
												if int64(len(cc.streams)) < int64(cc.maxConcurrentStreams) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1600
			_go_fuzz_dep_.CoverTab[74936]++
													return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1601
			// _ = "end of CoverTab[74936]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1602
			_go_fuzz_dep_.CoverTab[74937]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1602
			// _ = "end of CoverTab[74937]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1602
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1602
		// _ = "end of CoverTab[74931]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1602
		_go_fuzz_dep_.CoverTab[74932]++
												cc.pendingRequests++
												cc.cond.Wait()
												cc.pendingRequests--
												select {
		case <-cs.abort:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1607
			_go_fuzz_dep_.CoverTab[74938]++
													return cs.abortErr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1608
			// _ = "end of CoverTab[74938]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1609
			_go_fuzz_dep_.CoverTab[74939]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1609
			// _ = "end of CoverTab[74939]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1610
		// _ = "end of CoverTab[74932]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1611
	// _ = "end of CoverTab[74929]"
}

// requires cc.wmu be held
func (cc *ClientConn) writeHeaders(streamID uint32, endStream bool, maxFrameSize int, hdrs []byte) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1615
	_go_fuzz_dep_.CoverTab[74940]++
											first := true
											for len(hdrs) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1617
		_go_fuzz_dep_.CoverTab[74942]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1617
		return cc.werr == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1617
		// _ = "end of CoverTab[74942]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1617
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1617
		_go_fuzz_dep_.CoverTab[74943]++
												chunk := hdrs
												if len(chunk) > maxFrameSize {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1619
			_go_fuzz_dep_.CoverTab[74945]++
													chunk = chunk[:maxFrameSize]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1620
			// _ = "end of CoverTab[74945]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1621
			_go_fuzz_dep_.CoverTab[74946]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1621
			// _ = "end of CoverTab[74946]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1621
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1621
		// _ = "end of CoverTab[74943]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1621
		_go_fuzz_dep_.CoverTab[74944]++
												hdrs = hdrs[len(chunk):]
												endHeaders := len(hdrs) == 0
												if first {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1624
			_go_fuzz_dep_.CoverTab[74947]++
													cc.fr.WriteHeaders(HeadersFrameParam{
				StreamID:	streamID,
				BlockFragment:	chunk,
				EndStream:	endStream,
				EndHeaders:	endHeaders,
			})
													first = false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1631
			// _ = "end of CoverTab[74947]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1632
			_go_fuzz_dep_.CoverTab[74948]++
													cc.fr.WriteContinuation(streamID, endHeaders, chunk)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1633
			// _ = "end of CoverTab[74948]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1634
		// _ = "end of CoverTab[74944]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1635
	// _ = "end of CoverTab[74940]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1635
	_go_fuzz_dep_.CoverTab[74941]++
											cc.bw.Flush()
											return cc.werr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1637
	// _ = "end of CoverTab[74941]"
}

// internal error values; they don't escape to callers
var (
	// abort request body write; don't send cancel
	errStopReqBodyWrite	= errors.New("http2: aborting request body write")

	// abort request body write, but send stream reset of cancel.
	errStopReqBodyWriteAndCancel	= errors.New("http2: canceling request")

	errReqBodyTooLong	= errors.New("http2: request body larger than specified content length")
)

// frameScratchBufferLen returns the length of a buffer to use for
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1651
// outgoing request bodies to read/write to/from.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1651
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1651
// It returns max(1, min(peer's advertised max frame size,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1651
// Request.ContentLength+1, 512KB)).
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1656
func (cs *clientStream) frameScratchBufferLen(maxFrameSize int) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1656
	_go_fuzz_dep_.CoverTab[74949]++
											const max = 512 << 10
											n := int64(maxFrameSize)
											if n > max {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1659
		_go_fuzz_dep_.CoverTab[74953]++
												n = max
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1660
		// _ = "end of CoverTab[74953]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1661
		_go_fuzz_dep_.CoverTab[74954]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1661
		// _ = "end of CoverTab[74954]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1661
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1661
	// _ = "end of CoverTab[74949]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1661
	_go_fuzz_dep_.CoverTab[74950]++
											if cl := cs.reqBodyContentLength; cl != -1 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1662
		_go_fuzz_dep_.CoverTab[74955]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1662
		return cl+1 < n
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1662
		// _ = "end of CoverTab[74955]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1662
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1662
		_go_fuzz_dep_.CoverTab[74956]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1667
		n = cl + 1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1667
		// _ = "end of CoverTab[74956]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1668
		_go_fuzz_dep_.CoverTab[74957]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1668
		// _ = "end of CoverTab[74957]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1668
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1668
	// _ = "end of CoverTab[74950]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1668
	_go_fuzz_dep_.CoverTab[74951]++
											if n < 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1669
		_go_fuzz_dep_.CoverTab[74958]++
												return 1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1670
		// _ = "end of CoverTab[74958]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1671
		_go_fuzz_dep_.CoverTab[74959]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1671
		// _ = "end of CoverTab[74959]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1671
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1671
	// _ = "end of CoverTab[74951]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1671
	_go_fuzz_dep_.CoverTab[74952]++
											return int(n)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1672
	// _ = "end of CoverTab[74952]"
}

var bufPool sync.Pool	// of *[]byte

func (cs *clientStream) writeRequestBody(req *http.Request) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1677
	_go_fuzz_dep_.CoverTab[74960]++
											cc := cs.cc
											body := cs.reqBody
											sentEnd := false

											hasTrailers := req.Trailer != nil
											remainLen := cs.reqBodyContentLength
											hasContentLen := remainLen != -1

											cc.mu.Lock()
											maxFrameSize := int(cc.maxFrameSize)
											cc.mu.Unlock()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1691
	scratchLen := cs.frameScratchBufferLen(maxFrameSize)
	var buf []byte
	if bp, ok := bufPool.Get().(*[]byte); ok && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1693
		_go_fuzz_dep_.CoverTab[74968]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1693
		return len(*bp) >= scratchLen
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1693
		// _ = "end of CoverTab[74968]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1693
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1693
		_go_fuzz_dep_.CoverTab[74969]++
												defer bufPool.Put(bp)
												buf = *bp
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1695
		// _ = "end of CoverTab[74969]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1696
		_go_fuzz_dep_.CoverTab[74970]++
												buf = make([]byte, scratchLen)
												defer bufPool.Put(&buf)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1698
		// _ = "end of CoverTab[74970]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1699
	// _ = "end of CoverTab[74960]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1699
	_go_fuzz_dep_.CoverTab[74961]++

											var sawEOF bool
											for !sawEOF {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1702
		_go_fuzz_dep_.CoverTab[74971]++
												n, err := body.Read(buf)
												if hasContentLen {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1704
			_go_fuzz_dep_.CoverTab[74975]++
													remainLen -= int64(n)
													if remainLen == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1706
				_go_fuzz_dep_.CoverTab[74977]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1706
				return err == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1706
				// _ = "end of CoverTab[74977]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1706
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1706
				_go_fuzz_dep_.CoverTab[74978]++
				// The request body's Content-Length was predeclared and
				// we just finished reading it all, but the underlying io.Reader
				// returned the final chunk with a nil error (which is one of
				// the two valid things a Reader can do at EOF). Because we'd prefer
				// to send the END_STREAM bit early, double-check that we're actually
				// at EOF. Subsequent reads should return (0, EOF) at this point.
														// If either value is different, we return an error in one of two ways below.
														var scratch [1]byte
														var n1 int
														n1, err = body.Read(scratch[:])
														remainLen -= int64(n1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1717
				// _ = "end of CoverTab[74978]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1718
				_go_fuzz_dep_.CoverTab[74979]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1718
				// _ = "end of CoverTab[74979]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1718
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1718
			// _ = "end of CoverTab[74975]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1718
			_go_fuzz_dep_.CoverTab[74976]++
													if remainLen < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1719
				_go_fuzz_dep_.CoverTab[74980]++
														err = errReqBodyTooLong
														return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1721
				// _ = "end of CoverTab[74980]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1722
				_go_fuzz_dep_.CoverTab[74981]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1722
				// _ = "end of CoverTab[74981]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1722
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1722
			// _ = "end of CoverTab[74976]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1723
			_go_fuzz_dep_.CoverTab[74982]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1723
			// _ = "end of CoverTab[74982]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1723
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1723
		// _ = "end of CoverTab[74971]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1723
		_go_fuzz_dep_.CoverTab[74972]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1724
			_go_fuzz_dep_.CoverTab[74983]++
													cc.mu.Lock()
													bodyClosed := cs.reqBodyClosed != nil
													cc.mu.Unlock()
													switch {
			case bodyClosed:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1729
				_go_fuzz_dep_.CoverTab[74984]++
														return errStopReqBodyWrite
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1730
				// _ = "end of CoverTab[74984]"
			case err == io.EOF:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1731
				_go_fuzz_dep_.CoverTab[74985]++
														sawEOF = true
														err = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1733
				// _ = "end of CoverTab[74985]"
			default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1734
				_go_fuzz_dep_.CoverTab[74986]++
														return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1735
				// _ = "end of CoverTab[74986]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1736
			// _ = "end of CoverTab[74983]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1737
			_go_fuzz_dep_.CoverTab[74987]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1737
			// _ = "end of CoverTab[74987]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1737
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1737
		// _ = "end of CoverTab[74972]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1737
		_go_fuzz_dep_.CoverTab[74973]++

												remain := buf[:n]
												for len(remain) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1740
			_go_fuzz_dep_.CoverTab[74988]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1740
			return err == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1740
			// _ = "end of CoverTab[74988]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1740
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1740
			_go_fuzz_dep_.CoverTab[74989]++
													var allowed int32
													allowed, err = cs.awaitFlowControl(len(remain))
													if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1743
				_go_fuzz_dep_.CoverTab[74992]++
														return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1744
				// _ = "end of CoverTab[74992]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1745
				_go_fuzz_dep_.CoverTab[74993]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1745
				// _ = "end of CoverTab[74993]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1745
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1745
			// _ = "end of CoverTab[74989]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1745
			_go_fuzz_dep_.CoverTab[74990]++
													cc.wmu.Lock()
													data := remain[:allowed]
													remain = remain[allowed:]
													sentEnd = sawEOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1749
				_go_fuzz_dep_.CoverTab[74994]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1749
				return len(remain) == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1749
				// _ = "end of CoverTab[74994]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1749
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1749
				_go_fuzz_dep_.CoverTab[74995]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1749
				return !hasTrailers
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1749
				// _ = "end of CoverTab[74995]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1749
			}()
													err = cc.fr.WriteData(cs.ID, sentEnd, data)
													if err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1751
				_go_fuzz_dep_.CoverTab[74996]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1758
				err = cc.bw.Flush()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1758
				// _ = "end of CoverTab[74996]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1759
				_go_fuzz_dep_.CoverTab[74997]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1759
				// _ = "end of CoverTab[74997]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1759
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1759
			// _ = "end of CoverTab[74990]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1759
			_go_fuzz_dep_.CoverTab[74991]++
													cc.wmu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1760
			// _ = "end of CoverTab[74991]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1761
		// _ = "end of CoverTab[74973]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1761
		_go_fuzz_dep_.CoverTab[74974]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1762
			_go_fuzz_dep_.CoverTab[74998]++
													return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1763
			// _ = "end of CoverTab[74998]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1764
			_go_fuzz_dep_.CoverTab[74999]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1764
			// _ = "end of CoverTab[74999]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1764
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1764
		// _ = "end of CoverTab[74974]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1765
	// _ = "end of CoverTab[74961]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1765
	_go_fuzz_dep_.CoverTab[74962]++

											if sentEnd {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1767
		_go_fuzz_dep_.CoverTab[75000]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1771
		return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1771
		// _ = "end of CoverTab[75000]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1772
		_go_fuzz_dep_.CoverTab[75001]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1772
		// _ = "end of CoverTab[75001]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1772
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1772
	// _ = "end of CoverTab[74962]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1772
	_go_fuzz_dep_.CoverTab[74963]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1777
	cc.mu.Lock()
	trailer := req.Trailer
	err = cs.abortErr
	cc.mu.Unlock()
	if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1781
		_go_fuzz_dep_.CoverTab[75002]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1782
		// _ = "end of CoverTab[75002]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1783
		_go_fuzz_dep_.CoverTab[75003]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1783
		// _ = "end of CoverTab[75003]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1783
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1783
	// _ = "end of CoverTab[74963]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1783
	_go_fuzz_dep_.CoverTab[74964]++

											cc.wmu.Lock()
											defer cc.wmu.Unlock()
											var trls []byte
											if len(trailer) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1788
		_go_fuzz_dep_.CoverTab[75004]++
												trls, err = cc.encodeTrailers(trailer)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1790
			_go_fuzz_dep_.CoverTab[75005]++
													return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1791
			// _ = "end of CoverTab[75005]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1792
			_go_fuzz_dep_.CoverTab[75006]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1792
			// _ = "end of CoverTab[75006]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1792
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1792
		// _ = "end of CoverTab[75004]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1793
		_go_fuzz_dep_.CoverTab[75007]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1793
		// _ = "end of CoverTab[75007]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1793
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1793
	// _ = "end of CoverTab[74964]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1793
	_go_fuzz_dep_.CoverTab[74965]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1797
	if len(trls) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1797
		_go_fuzz_dep_.CoverTab[75008]++
												err = cc.writeHeaders(cs.ID, true, maxFrameSize, trls)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1798
		// _ = "end of CoverTab[75008]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1799
		_go_fuzz_dep_.CoverTab[75009]++
												err = cc.fr.WriteData(cs.ID, true, nil)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1800
		// _ = "end of CoverTab[75009]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1801
	// _ = "end of CoverTab[74965]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1801
	_go_fuzz_dep_.CoverTab[74966]++
											if ferr := cc.bw.Flush(); ferr != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1802
		_go_fuzz_dep_.CoverTab[75010]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1802
		return err == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1802
		// _ = "end of CoverTab[75010]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1802
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1802
		_go_fuzz_dep_.CoverTab[75011]++
												err = ferr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1803
		// _ = "end of CoverTab[75011]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1804
		_go_fuzz_dep_.CoverTab[75012]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1804
		// _ = "end of CoverTab[75012]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1804
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1804
	// _ = "end of CoverTab[74966]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1804
	_go_fuzz_dep_.CoverTab[74967]++
											return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1805
	// _ = "end of CoverTab[74967]"
}

// awaitFlowControl waits for [1, min(maxBytes, cc.cs.maxFrameSize)] flow
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1808
// control tokens from the server.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1808
// It returns either the non-zero number of tokens taken or an error
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1808
// if the stream is dead.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1812
func (cs *clientStream) awaitFlowControl(maxBytes int) (taken int32, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1812
	_go_fuzz_dep_.CoverTab[75013]++
											cc := cs.cc
											ctx := cs.ctx
											cc.mu.Lock()
											defer cc.mu.Unlock()
											for {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1817
		_go_fuzz_dep_.CoverTab[75014]++
												if cc.closed {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1818
			_go_fuzz_dep_.CoverTab[75019]++
													return 0, errClientConnClosed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1819
			// _ = "end of CoverTab[75019]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1820
			_go_fuzz_dep_.CoverTab[75020]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1820
			// _ = "end of CoverTab[75020]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1820
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1820
		// _ = "end of CoverTab[75014]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1820
		_go_fuzz_dep_.CoverTab[75015]++
												if cs.reqBodyClosed != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1821
			_go_fuzz_dep_.CoverTab[75021]++
													return 0, errStopReqBodyWrite
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1822
			// _ = "end of CoverTab[75021]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1823
			_go_fuzz_dep_.CoverTab[75022]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1823
			// _ = "end of CoverTab[75022]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1823
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1823
		// _ = "end of CoverTab[75015]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1823
		_go_fuzz_dep_.CoverTab[75016]++
												select {
		case <-cs.abort:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1825
			_go_fuzz_dep_.CoverTab[75023]++
													return 0, cs.abortErr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1826
			// _ = "end of CoverTab[75023]"
		case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1827
			_go_fuzz_dep_.CoverTab[75024]++
													return 0, ctx.Err()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1828
			// _ = "end of CoverTab[75024]"
		case <-cs.reqCancel:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1829
			_go_fuzz_dep_.CoverTab[75025]++
													return 0, errRequestCanceled
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1830
			// _ = "end of CoverTab[75025]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1831
			_go_fuzz_dep_.CoverTab[75026]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1831
			// _ = "end of CoverTab[75026]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1832
		// _ = "end of CoverTab[75016]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1832
		_go_fuzz_dep_.CoverTab[75017]++
												if a := cs.flow.available(); a > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1833
			_go_fuzz_dep_.CoverTab[75027]++
													take := a
													if int(take) > maxBytes {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1835
				_go_fuzz_dep_.CoverTab[75030]++

														take = int32(maxBytes)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1837
				// _ = "end of CoverTab[75030]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1838
				_go_fuzz_dep_.CoverTab[75031]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1838
				// _ = "end of CoverTab[75031]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1838
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1838
			// _ = "end of CoverTab[75027]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1838
			_go_fuzz_dep_.CoverTab[75028]++
													if take > int32(cc.maxFrameSize) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1839
				_go_fuzz_dep_.CoverTab[75032]++
														take = int32(cc.maxFrameSize)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1840
				// _ = "end of CoverTab[75032]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1841
				_go_fuzz_dep_.CoverTab[75033]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1841
				// _ = "end of CoverTab[75033]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1841
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1841
			// _ = "end of CoverTab[75028]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1841
			_go_fuzz_dep_.CoverTab[75029]++
													cs.flow.take(take)
													return take, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1843
			// _ = "end of CoverTab[75029]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1844
			_go_fuzz_dep_.CoverTab[75034]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1844
			// _ = "end of CoverTab[75034]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1844
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1844
		// _ = "end of CoverTab[75017]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1844
		_go_fuzz_dep_.CoverTab[75018]++
												cc.cond.Wait()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1845
		// _ = "end of CoverTab[75018]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1846
	// _ = "end of CoverTab[75013]"
}

var errNilRequestURL = errors.New("http2: Request.URI is nil")

// requires cc.wmu be held.
func (cc *ClientConn) encodeHeaders(req *http.Request, addGzipHeader bool, trailers string, contentLength int64) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1852
	_go_fuzz_dep_.CoverTab[75035]++
											cc.hbuf.Reset()
											if req.URL == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1854
		_go_fuzz_dep_.CoverTab[75045]++
												return nil, errNilRequestURL
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1855
		// _ = "end of CoverTab[75045]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1856
		_go_fuzz_dep_.CoverTab[75046]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1856
		// _ = "end of CoverTab[75046]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1856
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1856
	// _ = "end of CoverTab[75035]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1856
	_go_fuzz_dep_.CoverTab[75036]++

											host := req.Host
											if host == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1859
		_go_fuzz_dep_.CoverTab[75047]++
												host = req.URL.Host
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1860
		// _ = "end of CoverTab[75047]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1861
		_go_fuzz_dep_.CoverTab[75048]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1861
		// _ = "end of CoverTab[75048]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1861
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1861
	// _ = "end of CoverTab[75036]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1861
	_go_fuzz_dep_.CoverTab[75037]++
											host, err := httpguts.PunycodeHostPort(host)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1863
		_go_fuzz_dep_.CoverTab[75049]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1864
		// _ = "end of CoverTab[75049]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1865
		_go_fuzz_dep_.CoverTab[75050]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1865
		// _ = "end of CoverTab[75050]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1865
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1865
	// _ = "end of CoverTab[75037]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1865
	_go_fuzz_dep_.CoverTab[75038]++

											var path string
											if req.Method != "CONNECT" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1868
		_go_fuzz_dep_.CoverTab[75051]++
												path = req.URL.RequestURI()
												if !validPseudoPath(path) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1870
			_go_fuzz_dep_.CoverTab[75052]++
													orig := path
													path = strings.TrimPrefix(path, req.URL.Scheme+"://"+host)
													if !validPseudoPath(path) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1873
				_go_fuzz_dep_.CoverTab[75053]++
														if req.URL.Opaque != "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1874
					_go_fuzz_dep_.CoverTab[75054]++
															return nil, fmt.Errorf("invalid request :path %q from URL.Opaque = %q", orig, req.URL.Opaque)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1875
					// _ = "end of CoverTab[75054]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1876
					_go_fuzz_dep_.CoverTab[75055]++
															return nil, fmt.Errorf("invalid request :path %q", orig)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1877
					// _ = "end of CoverTab[75055]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1878
				// _ = "end of CoverTab[75053]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1879
				_go_fuzz_dep_.CoverTab[75056]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1879
				// _ = "end of CoverTab[75056]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1879
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1879
			// _ = "end of CoverTab[75052]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1880
			_go_fuzz_dep_.CoverTab[75057]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1880
			// _ = "end of CoverTab[75057]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1880
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1880
		// _ = "end of CoverTab[75051]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1881
		_go_fuzz_dep_.CoverTab[75058]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1881
		// _ = "end of CoverTab[75058]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1881
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1881
	// _ = "end of CoverTab[75038]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1881
	_go_fuzz_dep_.CoverTab[75039]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1886
	for k, vv := range req.Header {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1886
		_go_fuzz_dep_.CoverTab[75059]++
												if !httpguts.ValidHeaderFieldName(k) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1887
			_go_fuzz_dep_.CoverTab[75061]++
													return nil, fmt.Errorf("invalid HTTP header name %q", k)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1888
			// _ = "end of CoverTab[75061]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1889
			_go_fuzz_dep_.CoverTab[75062]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1889
			// _ = "end of CoverTab[75062]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1889
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1889
		// _ = "end of CoverTab[75059]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1889
		_go_fuzz_dep_.CoverTab[75060]++
												for _, v := range vv {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1890
			_go_fuzz_dep_.CoverTab[75063]++
													if !httpguts.ValidHeaderFieldValue(v) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1891
				_go_fuzz_dep_.CoverTab[75064]++

														return nil, fmt.Errorf("invalid HTTP header value for header %q", k)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1893
				// _ = "end of CoverTab[75064]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1894
				_go_fuzz_dep_.CoverTab[75065]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1894
				// _ = "end of CoverTab[75065]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1894
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1894
			// _ = "end of CoverTab[75063]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1895
		// _ = "end of CoverTab[75060]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1896
	// _ = "end of CoverTab[75039]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1896
	_go_fuzz_dep_.CoverTab[75040]++

											enumerateHeaders := func(f func(name, value string)) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1898
		_go_fuzz_dep_.CoverTab[75066]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1904
		f(":authority", host)
		m := req.Method
		if m == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1906
			_go_fuzz_dep_.CoverTab[75073]++
													m = http.MethodGet
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1907
			// _ = "end of CoverTab[75073]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1908
			_go_fuzz_dep_.CoverTab[75074]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1908
			// _ = "end of CoverTab[75074]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1908
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1908
		// _ = "end of CoverTab[75066]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1908
		_go_fuzz_dep_.CoverTab[75067]++
												f(":method", m)
												if req.Method != "CONNECT" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1910
			_go_fuzz_dep_.CoverTab[75075]++
													f(":path", path)
													f(":scheme", req.URL.Scheme)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1912
			// _ = "end of CoverTab[75075]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1913
			_go_fuzz_dep_.CoverTab[75076]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1913
			// _ = "end of CoverTab[75076]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1913
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1913
		// _ = "end of CoverTab[75067]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1913
		_go_fuzz_dep_.CoverTab[75068]++
												if trailers != "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1914
			_go_fuzz_dep_.CoverTab[75077]++
													f("trailer", trailers)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1915
			// _ = "end of CoverTab[75077]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1916
			_go_fuzz_dep_.CoverTab[75078]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1916
			// _ = "end of CoverTab[75078]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1916
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1916
		// _ = "end of CoverTab[75068]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1916
		_go_fuzz_dep_.CoverTab[75069]++

												var didUA bool
												for k, vv := range req.Header {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1919
			_go_fuzz_dep_.CoverTab[75079]++
													if asciiEqualFold(k, "host") || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1920
				_go_fuzz_dep_.CoverTab[75081]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1920
				return asciiEqualFold(k, "content-length")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1920
				// _ = "end of CoverTab[75081]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1920
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1920
				_go_fuzz_dep_.CoverTab[75082]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1923
				continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1923
				// _ = "end of CoverTab[75082]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1924
				_go_fuzz_dep_.CoverTab[75083]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1924
				if asciiEqualFold(k, "connection") || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1924
					_go_fuzz_dep_.CoverTab[75084]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1924
					return asciiEqualFold(k, "proxy-connection")
															// _ = "end of CoverTab[75084]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1925
				}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1925
					_go_fuzz_dep_.CoverTab[75085]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1925
					return asciiEqualFold(k, "transfer-encoding")
															// _ = "end of CoverTab[75085]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1926
				}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1926
					_go_fuzz_dep_.CoverTab[75086]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1926
					return asciiEqualFold(k, "upgrade")
															// _ = "end of CoverTab[75086]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1927
				}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1927
					_go_fuzz_dep_.CoverTab[75087]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1927
					return asciiEqualFold(k, "keep-alive")
															// _ = "end of CoverTab[75087]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1928
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1928
					_go_fuzz_dep_.CoverTab[75088]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1933
					continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1933
					// _ = "end of CoverTab[75088]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1934
					_go_fuzz_dep_.CoverTab[75089]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1934
					if asciiEqualFold(k, "user-agent") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1934
						_go_fuzz_dep_.CoverTab[75090]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1939
						didUA = true
						if len(vv) < 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1940
							_go_fuzz_dep_.CoverTab[75092]++
																	continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1941
							// _ = "end of CoverTab[75092]"
						} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1942
							_go_fuzz_dep_.CoverTab[75093]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1942
							// _ = "end of CoverTab[75093]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1942
						}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1942
						// _ = "end of CoverTab[75090]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1942
						_go_fuzz_dep_.CoverTab[75091]++
																vv = vv[:1]
																if vv[0] == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1944
							_go_fuzz_dep_.CoverTab[75094]++
																	continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1945
							// _ = "end of CoverTab[75094]"
						} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1946
							_go_fuzz_dep_.CoverTab[75095]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1946
							// _ = "end of CoverTab[75095]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1946
						}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1946
						// _ = "end of CoverTab[75091]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1947
						_go_fuzz_dep_.CoverTab[75096]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1947
						if asciiEqualFold(k, "cookie") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1947
							_go_fuzz_dep_.CoverTab[75097]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1951
							for _, v := range vv {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1951
								_go_fuzz_dep_.CoverTab[75099]++
																		for {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1952
									_go_fuzz_dep_.CoverTab[75101]++
																			p := strings.IndexByte(v, ';')
																			if p < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1954
										_go_fuzz_dep_.CoverTab[75104]++
																				break
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1955
										// _ = "end of CoverTab[75104]"
									} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1956
										_go_fuzz_dep_.CoverTab[75105]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1956
										// _ = "end of CoverTab[75105]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1956
									}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1956
									// _ = "end of CoverTab[75101]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1956
									_go_fuzz_dep_.CoverTab[75102]++
																			f("cookie", v[:p])
																			p++

																			for p+1 <= len(v) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1960
										_go_fuzz_dep_.CoverTab[75106]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1960
										return v[p] == ' '
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1960
										// _ = "end of CoverTab[75106]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1960
									}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1960
										_go_fuzz_dep_.CoverTab[75107]++
																				p++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1961
										// _ = "end of CoverTab[75107]"
									}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1962
									// _ = "end of CoverTab[75102]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1962
									_go_fuzz_dep_.CoverTab[75103]++
																			v = v[p:]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1963
									// _ = "end of CoverTab[75103]"
								}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1964
								// _ = "end of CoverTab[75099]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1964
								_go_fuzz_dep_.CoverTab[75100]++
																		if len(v) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1965
									_go_fuzz_dep_.CoverTab[75108]++
																			f("cookie", v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1966
									// _ = "end of CoverTab[75108]"
								} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1967
									_go_fuzz_dep_.CoverTab[75109]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1967
									// _ = "end of CoverTab[75109]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1967
								}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1967
								// _ = "end of CoverTab[75100]"
							}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1968
							// _ = "end of CoverTab[75097]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1968
							_go_fuzz_dep_.CoverTab[75098]++
																	continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1969
							// _ = "end of CoverTab[75098]"
						} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1970
							_go_fuzz_dep_.CoverTab[75110]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1970
							// _ = "end of CoverTab[75110]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1970
						}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1970
						// _ = "end of CoverTab[75096]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1970
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1970
					// _ = "end of CoverTab[75089]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1970
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1970
				// _ = "end of CoverTab[75083]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1970
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1970
			// _ = "end of CoverTab[75079]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1970
			_go_fuzz_dep_.CoverTab[75080]++

													for _, v := range vv {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1972
				_go_fuzz_dep_.CoverTab[75111]++
														f(k, v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1973
				// _ = "end of CoverTab[75111]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1974
			// _ = "end of CoverTab[75080]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1975
		// _ = "end of CoverTab[75069]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1975
		_go_fuzz_dep_.CoverTab[75070]++
												if shouldSendReqContentLength(req.Method, contentLength) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1976
			_go_fuzz_dep_.CoverTab[75112]++
													f("content-length", strconv.FormatInt(contentLength, 10))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1977
			// _ = "end of CoverTab[75112]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1978
			_go_fuzz_dep_.CoverTab[75113]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1978
			// _ = "end of CoverTab[75113]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1978
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1978
		// _ = "end of CoverTab[75070]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1978
		_go_fuzz_dep_.CoverTab[75071]++
												if addGzipHeader {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1979
			_go_fuzz_dep_.CoverTab[75114]++
													f("accept-encoding", "gzip")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1980
			// _ = "end of CoverTab[75114]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1981
			_go_fuzz_dep_.CoverTab[75115]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1981
			// _ = "end of CoverTab[75115]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1981
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1981
		// _ = "end of CoverTab[75071]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1981
		_go_fuzz_dep_.CoverTab[75072]++
												if !didUA {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1982
			_go_fuzz_dep_.CoverTab[75116]++
													f("user-agent", defaultUserAgent)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1983
			// _ = "end of CoverTab[75116]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1984
			_go_fuzz_dep_.CoverTab[75117]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1984
			// _ = "end of CoverTab[75117]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1984
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1984
		// _ = "end of CoverTab[75072]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1985
	// _ = "end of CoverTab[75040]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1985
	_go_fuzz_dep_.CoverTab[75041]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1991
	hlSize := uint64(0)
	enumerateHeaders(func(name, value string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1992
		_go_fuzz_dep_.CoverTab[75118]++
												hf := hpack.HeaderField{Name: name, Value: value}
												hlSize += uint64(hf.Size())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1994
		// _ = "end of CoverTab[75118]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1995
	// _ = "end of CoverTab[75041]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1995
	_go_fuzz_dep_.CoverTab[75042]++

											if hlSize > cc.peerMaxHeaderListSize {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1997
		_go_fuzz_dep_.CoverTab[75119]++
												return nil, errRequestHeaderListSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1998
		// _ = "end of CoverTab[75119]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1999
		_go_fuzz_dep_.CoverTab[75120]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1999
		// _ = "end of CoverTab[75120]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1999
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1999
	// _ = "end of CoverTab[75042]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:1999
	_go_fuzz_dep_.CoverTab[75043]++

											trace := httptrace.ContextClientTrace(req.Context())
											traceHeaders := traceHasWroteHeaderField(trace)

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2005
	enumerateHeaders(func(name, value string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2005
		_go_fuzz_dep_.CoverTab[75121]++
												name, ascii := lowerHeader(name)
												if !ascii {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2007
			_go_fuzz_dep_.CoverTab[75123]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2010
			return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2010
			// _ = "end of CoverTab[75123]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2011
			_go_fuzz_dep_.CoverTab[75124]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2011
			// _ = "end of CoverTab[75124]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2011
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2011
		// _ = "end of CoverTab[75121]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2011
		_go_fuzz_dep_.CoverTab[75122]++
												cc.writeHeader(name, value)
												if traceHeaders {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2013
			_go_fuzz_dep_.CoverTab[75125]++
													traceWroteHeaderField(trace, name, value)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2014
			// _ = "end of CoverTab[75125]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2015
			_go_fuzz_dep_.CoverTab[75126]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2015
			// _ = "end of CoverTab[75126]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2015
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2015
		// _ = "end of CoverTab[75122]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2016
	// _ = "end of CoverTab[75043]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2016
	_go_fuzz_dep_.CoverTab[75044]++

											return cc.hbuf.Bytes(), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2018
	// _ = "end of CoverTab[75044]"
}

// shouldSendReqContentLength reports whether the http2.Transport should send
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2021
// a "content-length" request header. This logic is basically a copy of the net/http
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2021
// transferWriter.shouldSendContentLength.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2021
// The contentLength is the corrected contentLength (so 0 means actually 0, not unknown).
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2021
// -1 means unknown.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2026
func shouldSendReqContentLength(method string, contentLength int64) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2026
	_go_fuzz_dep_.CoverTab[75127]++
											if contentLength > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2027
		_go_fuzz_dep_.CoverTab[75130]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2028
		// _ = "end of CoverTab[75130]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2029
		_go_fuzz_dep_.CoverTab[75131]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2029
		// _ = "end of CoverTab[75131]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2029
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2029
	// _ = "end of CoverTab[75127]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2029
	_go_fuzz_dep_.CoverTab[75128]++
											if contentLength < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2030
		_go_fuzz_dep_.CoverTab[75132]++
												return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2031
		// _ = "end of CoverTab[75132]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2032
		_go_fuzz_dep_.CoverTab[75133]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2032
		// _ = "end of CoverTab[75133]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2032
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2032
	// _ = "end of CoverTab[75128]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2032
	_go_fuzz_dep_.CoverTab[75129]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2035
	switch method {
	case "POST", "PUT", "PATCH":
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2036
		_go_fuzz_dep_.CoverTab[75134]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2037
		// _ = "end of CoverTab[75134]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2038
		_go_fuzz_dep_.CoverTab[75135]++
												return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2039
		// _ = "end of CoverTab[75135]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2040
	// _ = "end of CoverTab[75129]"
}

// requires cc.wmu be held.
func (cc *ClientConn) encodeTrailers(trailer http.Header) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2044
	_go_fuzz_dep_.CoverTab[75136]++
											cc.hbuf.Reset()

											hlSize := uint64(0)
											for k, vv := range trailer {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2048
		_go_fuzz_dep_.CoverTab[75140]++
												for _, v := range vv {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2049
			_go_fuzz_dep_.CoverTab[75141]++
													hf := hpack.HeaderField{Name: k, Value: v}
													hlSize += uint64(hf.Size())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2051
			// _ = "end of CoverTab[75141]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2052
		// _ = "end of CoverTab[75140]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2053
	// _ = "end of CoverTab[75136]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2053
	_go_fuzz_dep_.CoverTab[75137]++
											if hlSize > cc.peerMaxHeaderListSize {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2054
		_go_fuzz_dep_.CoverTab[75142]++
												return nil, errRequestHeaderListSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2055
		// _ = "end of CoverTab[75142]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2056
		_go_fuzz_dep_.CoverTab[75143]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2056
		// _ = "end of CoverTab[75143]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2056
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2056
	// _ = "end of CoverTab[75137]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2056
	_go_fuzz_dep_.CoverTab[75138]++

											for k, vv := range trailer {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2058
		_go_fuzz_dep_.CoverTab[75144]++
												lowKey, ascii := lowerHeader(k)
												if !ascii {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2060
			_go_fuzz_dep_.CoverTab[75146]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2063
			continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2063
			// _ = "end of CoverTab[75146]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2064
			_go_fuzz_dep_.CoverTab[75147]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2064
			// _ = "end of CoverTab[75147]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2064
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2064
		// _ = "end of CoverTab[75144]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2064
		_go_fuzz_dep_.CoverTab[75145]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2067
		for _, v := range vv {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2067
			_go_fuzz_dep_.CoverTab[75148]++
													cc.writeHeader(lowKey, v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2068
			// _ = "end of CoverTab[75148]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2069
		// _ = "end of CoverTab[75145]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2070
	// _ = "end of CoverTab[75138]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2070
	_go_fuzz_dep_.CoverTab[75139]++
											return cc.hbuf.Bytes(), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2071
	// _ = "end of CoverTab[75139]"
}

func (cc *ClientConn) writeHeader(name, value string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2074
	_go_fuzz_dep_.CoverTab[75149]++
											if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2075
		_go_fuzz_dep_.CoverTab[75151]++
												log.Printf("http2: Transport encoding header %q = %q", name, value)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2076
		// _ = "end of CoverTab[75151]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2077
		_go_fuzz_dep_.CoverTab[75152]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2077
		// _ = "end of CoverTab[75152]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2077
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2077
	// _ = "end of CoverTab[75149]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2077
	_go_fuzz_dep_.CoverTab[75150]++
											cc.henc.WriteField(hpack.HeaderField{Name: name, Value: value})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2078
	// _ = "end of CoverTab[75150]"
}

type resAndError struct {
	_	incomparable
	res	*http.Response
	err	error
}

// requires cc.mu be held.
func (cc *ClientConn) addStreamLocked(cs *clientStream) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2088
	_go_fuzz_dep_.CoverTab[75153]++
											cs.flow.add(int32(cc.initialWindowSize))
											cs.flow.setConnFlow(&cc.flow)
											cs.inflow.init(transportDefaultStreamFlow)
											cs.ID = cc.nextStreamID
											cc.nextStreamID += 2
											cc.streams[cs.ID] = cs
											if cs.ID == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2095
		_go_fuzz_dep_.CoverTab[75154]++
												panic("assigned stream ID 0")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2096
		// _ = "end of CoverTab[75154]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2097
		_go_fuzz_dep_.CoverTab[75155]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2097
		// _ = "end of CoverTab[75155]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2097
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2097
	// _ = "end of CoverTab[75153]"
}

func (cc *ClientConn) forgetStreamID(id uint32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2100
	_go_fuzz_dep_.CoverTab[75156]++
											cc.mu.Lock()
											slen := len(cc.streams)
											delete(cc.streams, id)
											if len(cc.streams) != slen-1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2104
		_go_fuzz_dep_.CoverTab[75160]++
												panic("forgetting unknown stream id")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2105
		// _ = "end of CoverTab[75160]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2106
		_go_fuzz_dep_.CoverTab[75161]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2106
		// _ = "end of CoverTab[75161]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2106
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2106
	// _ = "end of CoverTab[75156]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2106
	_go_fuzz_dep_.CoverTab[75157]++
											cc.lastActive = time.Now()
											if len(cc.streams) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2108
		_go_fuzz_dep_.CoverTab[75162]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2108
		return cc.idleTimer != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2108
		// _ = "end of CoverTab[75162]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2108
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2108
		_go_fuzz_dep_.CoverTab[75163]++
												cc.idleTimer.Reset(cc.idleTimeout)
												cc.lastIdle = time.Now()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2110
		// _ = "end of CoverTab[75163]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2111
		_go_fuzz_dep_.CoverTab[75164]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2111
		// _ = "end of CoverTab[75164]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2111
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2111
	// _ = "end of CoverTab[75157]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2111
	_go_fuzz_dep_.CoverTab[75158]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2114
	cc.cond.Broadcast()

	closeOnIdle := cc.singleUse || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2116
		_go_fuzz_dep_.CoverTab[75165]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2116
		return cc.doNotReuse
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2116
		// _ = "end of CoverTab[75165]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2116
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2116
		_go_fuzz_dep_.CoverTab[75166]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2116
		return cc.t.disableKeepAlives()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2116
		// _ = "end of CoverTab[75166]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2116
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2116
		_go_fuzz_dep_.CoverTab[75167]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2116
		return cc.goAway != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2116
		// _ = "end of CoverTab[75167]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2116
	}()
											if closeOnIdle && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2117
		_go_fuzz_dep_.CoverTab[75168]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2117
		return cc.streamsReserved == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2117
		// _ = "end of CoverTab[75168]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2117
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2117
		_go_fuzz_dep_.CoverTab[75169]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2117
		return len(cc.streams) == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2117
		// _ = "end of CoverTab[75169]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2117
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2117
		_go_fuzz_dep_.CoverTab[75170]++
												if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2118
			_go_fuzz_dep_.CoverTab[75172]++
													cc.vlogf("http2: Transport closing idle conn %p (forSingleUse=%v, maxStream=%v)", cc, cc.singleUse, cc.nextStreamID-2)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2119
			// _ = "end of CoverTab[75172]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2120
			_go_fuzz_dep_.CoverTab[75173]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2120
			// _ = "end of CoverTab[75173]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2120
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2120
		// _ = "end of CoverTab[75170]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2120
		_go_fuzz_dep_.CoverTab[75171]++
												cc.closed = true
												defer cc.closeConn()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2122
		// _ = "end of CoverTab[75171]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2123
		_go_fuzz_dep_.CoverTab[75174]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2123
		// _ = "end of CoverTab[75174]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2123
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2123
	// _ = "end of CoverTab[75158]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2123
	_go_fuzz_dep_.CoverTab[75159]++

											cc.mu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2125
	// _ = "end of CoverTab[75159]"
}

// clientConnReadLoop is the state owned by the clientConn's frame-reading readLoop.
type clientConnReadLoop struct {
	_	incomparable
	cc	*ClientConn
}

// readLoop runs in its own goroutine and reads and dispatches frames.
func (cc *ClientConn) readLoop() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2135
	_go_fuzz_dep_.CoverTab[75175]++
											rl := &clientConnReadLoop{cc: cc}
											defer rl.cleanup()
											cc.readerErr = rl.run()
											if ce, ok := cc.readerErr.(ConnectionError); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2139
		_go_fuzz_dep_.CoverTab[75176]++
												cc.wmu.Lock()
												cc.fr.WriteGoAway(0, ErrCode(ce), nil)
												cc.wmu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2142
		// _ = "end of CoverTab[75176]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2143
		_go_fuzz_dep_.CoverTab[75177]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2143
		// _ = "end of CoverTab[75177]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2143
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2143
	// _ = "end of CoverTab[75175]"
}

// GoAwayError is returned by the Transport when the server closes the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2146
// TCP connection after sending a GOAWAY frame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2148
type GoAwayError struct {
	LastStreamID	uint32
	ErrCode		ErrCode
	DebugData	string
}

func (e GoAwayError) Error() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2154
	_go_fuzz_dep_.CoverTab[75178]++
											return fmt.Sprintf("http2: server sent GOAWAY and closed the connection; LastStreamID=%v, ErrCode=%v, debug=%q",
		e.LastStreamID, e.ErrCode, e.DebugData)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2156
	// _ = "end of CoverTab[75178]"
}

func isEOFOrNetReadError(err error) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2159
	_go_fuzz_dep_.CoverTab[75179]++
											if err == io.EOF {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2160
		_go_fuzz_dep_.CoverTab[75181]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2161
		// _ = "end of CoverTab[75181]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2162
		_go_fuzz_dep_.CoverTab[75182]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2162
		// _ = "end of CoverTab[75182]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2162
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2162
	// _ = "end of CoverTab[75179]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2162
	_go_fuzz_dep_.CoverTab[75180]++
											ne, ok := err.(*net.OpError)
											return ok && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2164
		_go_fuzz_dep_.CoverTab[75183]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2164
		return ne.Op == "read"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2164
		// _ = "end of CoverTab[75183]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2164
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2164
	// _ = "end of CoverTab[75180]"
}

func (rl *clientConnReadLoop) cleanup() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2167
	_go_fuzz_dep_.CoverTab[75184]++
											cc := rl.cc
											cc.t.connPool().MarkDead(cc)
											defer cc.closeConn()
											defer close(cc.readerDone)

											if cc.idleTimer != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2173
		_go_fuzz_dep_.CoverTab[75188]++
												cc.idleTimer.Stop()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2174
		// _ = "end of CoverTab[75188]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2175
		_go_fuzz_dep_.CoverTab[75189]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2175
		// _ = "end of CoverTab[75189]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2175
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2175
	// _ = "end of CoverTab[75184]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2175
	_go_fuzz_dep_.CoverTab[75185]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2180
	err := cc.readerErr
	cc.mu.Lock()
	if cc.goAway != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2182
		_go_fuzz_dep_.CoverTab[75190]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2182
		return isEOFOrNetReadError(err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2182
		// _ = "end of CoverTab[75190]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2182
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2182
		_go_fuzz_dep_.CoverTab[75191]++
												err = GoAwayError{
			LastStreamID:	cc.goAway.LastStreamID,
			ErrCode:	cc.goAway.ErrCode,
			DebugData:	cc.goAwayDebug,
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2187
		// _ = "end of CoverTab[75191]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2188
		_go_fuzz_dep_.CoverTab[75192]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2188
		if err == io.EOF {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2188
			_go_fuzz_dep_.CoverTab[75193]++
													err = io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2189
			// _ = "end of CoverTab[75193]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2190
			_go_fuzz_dep_.CoverTab[75194]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2190
			// _ = "end of CoverTab[75194]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2190
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2190
		// _ = "end of CoverTab[75192]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2190
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2190
	// _ = "end of CoverTab[75185]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2190
	_go_fuzz_dep_.CoverTab[75186]++
											cc.closed = true

											for _, cs := range cc.streams {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2193
		_go_fuzz_dep_.CoverTab[75195]++
												select {
		case <-cs.peerClosed:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2195
			_go_fuzz_dep_.CoverTab[75196]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2195
			// _ = "end of CoverTab[75196]"

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2198
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2198
			_go_fuzz_dep_.CoverTab[75197]++
													cs.abortStreamLocked(err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2199
			// _ = "end of CoverTab[75197]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2200
		// _ = "end of CoverTab[75195]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2201
	// _ = "end of CoverTab[75186]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2201
	_go_fuzz_dep_.CoverTab[75187]++
											cc.cond.Broadcast()
											cc.mu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2203
	// _ = "end of CoverTab[75187]"
}

// countReadFrameError calls Transport.CountError with a string
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2206
// representing err.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2208
func (cc *ClientConn) countReadFrameError(err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2208
	_go_fuzz_dep_.CoverTab[75198]++
											f := cc.t.CountError
											if f == nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2210
		_go_fuzz_dep_.CoverTab[75204]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2210
		return err == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2210
		// _ = "end of CoverTab[75204]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2210
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2210
		_go_fuzz_dep_.CoverTab[75205]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2211
		// _ = "end of CoverTab[75205]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2212
		_go_fuzz_dep_.CoverTab[75206]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2212
		// _ = "end of CoverTab[75206]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2212
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2212
	// _ = "end of CoverTab[75198]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2212
	_go_fuzz_dep_.CoverTab[75199]++
											if ce, ok := err.(ConnectionError); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2213
		_go_fuzz_dep_.CoverTab[75207]++
												errCode := ErrCode(ce)
												f(fmt.Sprintf("read_frame_conn_error_%s", errCode.stringToken()))
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2216
		// _ = "end of CoverTab[75207]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2217
		_go_fuzz_dep_.CoverTab[75208]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2217
		// _ = "end of CoverTab[75208]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2217
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2217
	// _ = "end of CoverTab[75199]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2217
	_go_fuzz_dep_.CoverTab[75200]++
											if errors.Is(err, io.EOF) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2218
		_go_fuzz_dep_.CoverTab[75209]++
												f("read_frame_eof")
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2220
		// _ = "end of CoverTab[75209]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2221
		_go_fuzz_dep_.CoverTab[75210]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2221
		// _ = "end of CoverTab[75210]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2221
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2221
	// _ = "end of CoverTab[75200]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2221
	_go_fuzz_dep_.CoverTab[75201]++
											if errors.Is(err, io.ErrUnexpectedEOF) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2222
		_go_fuzz_dep_.CoverTab[75211]++
												f("read_frame_unexpected_eof")
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2224
		// _ = "end of CoverTab[75211]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2225
		_go_fuzz_dep_.CoverTab[75212]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2225
		// _ = "end of CoverTab[75212]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2225
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2225
	// _ = "end of CoverTab[75201]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2225
	_go_fuzz_dep_.CoverTab[75202]++
											if errors.Is(err, ErrFrameTooLarge) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2226
		_go_fuzz_dep_.CoverTab[75213]++
												f("read_frame_too_large")
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2228
		// _ = "end of CoverTab[75213]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2229
		_go_fuzz_dep_.CoverTab[75214]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2229
		// _ = "end of CoverTab[75214]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2229
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2229
	// _ = "end of CoverTab[75202]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2229
	_go_fuzz_dep_.CoverTab[75203]++
											f("read_frame_other")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2230
	// _ = "end of CoverTab[75203]"
}

func (rl *clientConnReadLoop) run() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2233
	_go_fuzz_dep_.CoverTab[75215]++
											cc := rl.cc
											gotSettings := false
											readIdleTimeout := cc.t.ReadIdleTimeout
											var t *time.Timer
											if readIdleTimeout != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2238
		_go_fuzz_dep_.CoverTab[75217]++
												t = time.AfterFunc(readIdleTimeout, cc.healthCheck)
												defer t.Stop()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2240
		// _ = "end of CoverTab[75217]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2241
		_go_fuzz_dep_.CoverTab[75218]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2241
		// _ = "end of CoverTab[75218]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2241
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2241
	// _ = "end of CoverTab[75215]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2241
	_go_fuzz_dep_.CoverTab[75216]++
											for {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2242
		_go_fuzz_dep_.CoverTab[75219]++
												f, err := cc.fr.ReadFrame()
												if t != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2244
			_go_fuzz_dep_.CoverTab[75226]++
													t.Reset(readIdleTimeout)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2245
			// _ = "end of CoverTab[75226]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2246
			_go_fuzz_dep_.CoverTab[75227]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2246
			// _ = "end of CoverTab[75227]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2246
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2246
		// _ = "end of CoverTab[75219]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2246
		_go_fuzz_dep_.CoverTab[75220]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2247
			_go_fuzz_dep_.CoverTab[75228]++
													cc.vlogf("http2: Transport readFrame error on conn %p: (%T) %v", cc, err, err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2248
			// _ = "end of CoverTab[75228]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2249
			_go_fuzz_dep_.CoverTab[75229]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2249
			// _ = "end of CoverTab[75229]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2249
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2249
		// _ = "end of CoverTab[75220]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2249
		_go_fuzz_dep_.CoverTab[75221]++
												if se, ok := err.(StreamError); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2250
			_go_fuzz_dep_.CoverTab[75230]++
													if cs := rl.streamByID(se.StreamID); cs != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2251
				_go_fuzz_dep_.CoverTab[75232]++
														if se.Cause == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2252
					_go_fuzz_dep_.CoverTab[75234]++
															se.Cause = cc.fr.errDetail
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2253
					// _ = "end of CoverTab[75234]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2254
					_go_fuzz_dep_.CoverTab[75235]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2254
					// _ = "end of CoverTab[75235]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2254
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2254
				// _ = "end of CoverTab[75232]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2254
				_go_fuzz_dep_.CoverTab[75233]++
														rl.endStreamError(cs, se)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2255
				// _ = "end of CoverTab[75233]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2256
				_go_fuzz_dep_.CoverTab[75236]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2256
				// _ = "end of CoverTab[75236]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2256
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2256
			// _ = "end of CoverTab[75230]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2256
			_go_fuzz_dep_.CoverTab[75231]++
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2257
			// _ = "end of CoverTab[75231]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2258
			_go_fuzz_dep_.CoverTab[75237]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2258
			if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2258
				_go_fuzz_dep_.CoverTab[75238]++
														cc.countReadFrameError(err)
														return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2260
				// _ = "end of CoverTab[75238]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2261
				_go_fuzz_dep_.CoverTab[75239]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2261
				// _ = "end of CoverTab[75239]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2261
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2261
			// _ = "end of CoverTab[75237]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2261
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2261
		// _ = "end of CoverTab[75221]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2261
		_go_fuzz_dep_.CoverTab[75222]++
												if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2262
			_go_fuzz_dep_.CoverTab[75240]++
													cc.vlogf("http2: Transport received %s", summarizeFrame(f))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2263
			// _ = "end of CoverTab[75240]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2264
			_go_fuzz_dep_.CoverTab[75241]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2264
			// _ = "end of CoverTab[75241]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2264
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2264
		// _ = "end of CoverTab[75222]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2264
		_go_fuzz_dep_.CoverTab[75223]++
												if !gotSettings {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2265
			_go_fuzz_dep_.CoverTab[75242]++
													if _, ok := f.(*SettingsFrame); !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2266
				_go_fuzz_dep_.CoverTab[75244]++
														cc.logf("protocol error: received %T before a SETTINGS frame", f)
														return ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2268
				// _ = "end of CoverTab[75244]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2269
				_go_fuzz_dep_.CoverTab[75245]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2269
				// _ = "end of CoverTab[75245]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2269
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2269
			// _ = "end of CoverTab[75242]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2269
			_go_fuzz_dep_.CoverTab[75243]++
													gotSettings = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2270
			// _ = "end of CoverTab[75243]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2271
			_go_fuzz_dep_.CoverTab[75246]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2271
			// _ = "end of CoverTab[75246]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2271
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2271
		// _ = "end of CoverTab[75223]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2271
		_go_fuzz_dep_.CoverTab[75224]++

												switch f := f.(type) {
		case *MetaHeadersFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2274
			_go_fuzz_dep_.CoverTab[75247]++
													err = rl.processHeaders(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2275
			// _ = "end of CoverTab[75247]"
		case *DataFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2276
			_go_fuzz_dep_.CoverTab[75248]++
													err = rl.processData(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2277
			// _ = "end of CoverTab[75248]"
		case *GoAwayFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2278
			_go_fuzz_dep_.CoverTab[75249]++
													err = rl.processGoAway(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2279
			// _ = "end of CoverTab[75249]"
		case *RSTStreamFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2280
			_go_fuzz_dep_.CoverTab[75250]++
													err = rl.processResetStream(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2281
			// _ = "end of CoverTab[75250]"
		case *SettingsFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2282
			_go_fuzz_dep_.CoverTab[75251]++
													err = rl.processSettings(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2283
			// _ = "end of CoverTab[75251]"
		case *PushPromiseFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2284
			_go_fuzz_dep_.CoverTab[75252]++
													err = rl.processPushPromise(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2285
			// _ = "end of CoverTab[75252]"
		case *WindowUpdateFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2286
			_go_fuzz_dep_.CoverTab[75253]++
													err = rl.processWindowUpdate(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2287
			// _ = "end of CoverTab[75253]"
		case *PingFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2288
			_go_fuzz_dep_.CoverTab[75254]++
													err = rl.processPing(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2289
			// _ = "end of CoverTab[75254]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2290
			_go_fuzz_dep_.CoverTab[75255]++
													cc.logf("Transport: unhandled response frame type %T", f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2291
			// _ = "end of CoverTab[75255]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2292
		// _ = "end of CoverTab[75224]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2292
		_go_fuzz_dep_.CoverTab[75225]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2293
			_go_fuzz_dep_.CoverTab[75256]++
													if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2294
				_go_fuzz_dep_.CoverTab[75258]++
														cc.vlogf("http2: Transport conn %p received error from processing frame %v: %v", cc, summarizeFrame(f), err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2295
				// _ = "end of CoverTab[75258]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2296
				_go_fuzz_dep_.CoverTab[75259]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2296
				// _ = "end of CoverTab[75259]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2296
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2296
			// _ = "end of CoverTab[75256]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2296
			_go_fuzz_dep_.CoverTab[75257]++
													return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2297
			// _ = "end of CoverTab[75257]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2298
			_go_fuzz_dep_.CoverTab[75260]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2298
			// _ = "end of CoverTab[75260]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2298
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2298
		// _ = "end of CoverTab[75225]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2299
	// _ = "end of CoverTab[75216]"
}

func (rl *clientConnReadLoop) processHeaders(f *MetaHeadersFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2302
	_go_fuzz_dep_.CoverTab[75261]++
											cs := rl.streamByID(f.StreamID)
											if cs == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2304
		_go_fuzz_dep_.CoverTab[75269]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2308
		return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2308
		// _ = "end of CoverTab[75269]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2309
		_go_fuzz_dep_.CoverTab[75270]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2309
		// _ = "end of CoverTab[75270]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2309
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2309
	// _ = "end of CoverTab[75261]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2309
	_go_fuzz_dep_.CoverTab[75262]++
											if cs.readClosed {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2310
		_go_fuzz_dep_.CoverTab[75271]++
												rl.endStreamError(cs, StreamError{
			StreamID:	f.StreamID,
			Code:		ErrCodeProtocol,
			Cause:		errors.New("protocol error: headers after END_STREAM"),
		})
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2316
		// _ = "end of CoverTab[75271]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2317
		_go_fuzz_dep_.CoverTab[75272]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2317
		// _ = "end of CoverTab[75272]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2317
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2317
	// _ = "end of CoverTab[75262]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2317
	_go_fuzz_dep_.CoverTab[75263]++
											if !cs.firstByte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2318
		_go_fuzz_dep_.CoverTab[75273]++
												if cs.trace != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2319
			_go_fuzz_dep_.CoverTab[75275]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2324
			traceFirstResponseByte(cs.trace)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2324
			// _ = "end of CoverTab[75275]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2325
			_go_fuzz_dep_.CoverTab[75276]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2325
			// _ = "end of CoverTab[75276]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2325
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2325
		// _ = "end of CoverTab[75273]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2325
		_go_fuzz_dep_.CoverTab[75274]++
												cs.firstByte = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2326
		// _ = "end of CoverTab[75274]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2327
		_go_fuzz_dep_.CoverTab[75277]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2327
		// _ = "end of CoverTab[75277]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2327
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2327
	// _ = "end of CoverTab[75263]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2327
	_go_fuzz_dep_.CoverTab[75264]++
											if !cs.pastHeaders {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2328
		_go_fuzz_dep_.CoverTab[75278]++
												cs.pastHeaders = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2329
		// _ = "end of CoverTab[75278]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2330
		_go_fuzz_dep_.CoverTab[75279]++
												return rl.processTrailers(cs, f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2331
		// _ = "end of CoverTab[75279]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2332
	// _ = "end of CoverTab[75264]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2332
	_go_fuzz_dep_.CoverTab[75265]++

											res, err := rl.handleResponse(cs, f)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2335
		_go_fuzz_dep_.CoverTab[75280]++
												if _, ok := err.(ConnectionError); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2336
			_go_fuzz_dep_.CoverTab[75282]++
													return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2337
			// _ = "end of CoverTab[75282]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2338
			_go_fuzz_dep_.CoverTab[75283]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2338
			// _ = "end of CoverTab[75283]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2338
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2338
		// _ = "end of CoverTab[75280]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2338
		_go_fuzz_dep_.CoverTab[75281]++

												rl.endStreamError(cs, StreamError{
			StreamID:	f.StreamID,
			Code:		ErrCodeProtocol,
			Cause:		err,
		})
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2345
		// _ = "end of CoverTab[75281]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2346
		_go_fuzz_dep_.CoverTab[75284]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2346
		// _ = "end of CoverTab[75284]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2346
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2346
	// _ = "end of CoverTab[75265]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2346
	_go_fuzz_dep_.CoverTab[75266]++
											if res == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2347
		_go_fuzz_dep_.CoverTab[75285]++

												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2349
		// _ = "end of CoverTab[75285]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2350
		_go_fuzz_dep_.CoverTab[75286]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2350
		// _ = "end of CoverTab[75286]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2350
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2350
	// _ = "end of CoverTab[75266]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2350
	_go_fuzz_dep_.CoverTab[75267]++
											cs.resTrailer = &res.Trailer
											cs.res = res
											close(cs.respHeaderRecv)
											if f.StreamEnded() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2354
		_go_fuzz_dep_.CoverTab[75287]++
												rl.endStream(cs)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2355
		// _ = "end of CoverTab[75287]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2356
		_go_fuzz_dep_.CoverTab[75288]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2356
		// _ = "end of CoverTab[75288]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2356
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2356
	// _ = "end of CoverTab[75267]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2356
	_go_fuzz_dep_.CoverTab[75268]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2357
	// _ = "end of CoverTab[75268]"
}

// may return error types nil, or ConnectionError. Any other error value
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2360
// is a StreamError of type ErrCodeProtocol. The returned error in that case
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2360
// is the detail.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2360
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2360
// As a special case, handleResponse may return (nil, nil) to skip the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2360
// frame (currently only used for 1xx responses).
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2366
func (rl *clientConnReadLoop) handleResponse(cs *clientStream, f *MetaHeadersFrame) (*http.Response, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2366
	_go_fuzz_dep_.CoverTab[75289]++
											if f.Truncated {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2367
		_go_fuzz_dep_.CoverTab[75299]++
												return nil, errResponseHeaderListSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2368
		// _ = "end of CoverTab[75299]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2369
		_go_fuzz_dep_.CoverTab[75300]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2369
		// _ = "end of CoverTab[75300]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2369
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2369
	// _ = "end of CoverTab[75289]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2369
	_go_fuzz_dep_.CoverTab[75290]++

											status := f.PseudoValue("status")
											if status == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2372
		_go_fuzz_dep_.CoverTab[75301]++
												return nil, errors.New("malformed response from server: missing status pseudo header")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2373
		// _ = "end of CoverTab[75301]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2374
		_go_fuzz_dep_.CoverTab[75302]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2374
		// _ = "end of CoverTab[75302]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2374
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2374
	// _ = "end of CoverTab[75290]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2374
	_go_fuzz_dep_.CoverTab[75291]++
											statusCode, err := strconv.Atoi(status)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2376
		_go_fuzz_dep_.CoverTab[75303]++
												return nil, errors.New("malformed response from server: malformed non-numeric status pseudo header")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2377
		// _ = "end of CoverTab[75303]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2378
		_go_fuzz_dep_.CoverTab[75304]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2378
		// _ = "end of CoverTab[75304]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2378
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2378
	// _ = "end of CoverTab[75291]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2378
	_go_fuzz_dep_.CoverTab[75292]++

											regularFields := f.RegularFields()
											strs := make([]string, len(regularFields))
											header := make(http.Header, len(regularFields))
											res := &http.Response{
		Proto:		"HTTP/2.0",
		ProtoMajor:	2,
		Header:		header,
		StatusCode:	statusCode,
		Status:		status + " " + http.StatusText(statusCode),
	}
	for _, hf := range regularFields {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2390
		_go_fuzz_dep_.CoverTab[75305]++
												key := canonicalHeader(hf.Name)
												if key == "Trailer" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2392
			_go_fuzz_dep_.CoverTab[75306]++
													t := res.Trailer
													if t == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2394
				_go_fuzz_dep_.CoverTab[75308]++
														t = make(http.Header)
														res.Trailer = t
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2396
				// _ = "end of CoverTab[75308]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2397
				_go_fuzz_dep_.CoverTab[75309]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2397
				// _ = "end of CoverTab[75309]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2397
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2397
			// _ = "end of CoverTab[75306]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2397
			_go_fuzz_dep_.CoverTab[75307]++
													foreachHeaderElement(hf.Value, func(v string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2398
				_go_fuzz_dep_.CoverTab[75310]++
														t[canonicalHeader(v)] = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2399
				// _ = "end of CoverTab[75310]"
			})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2400
			// _ = "end of CoverTab[75307]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2401
			_go_fuzz_dep_.CoverTab[75311]++
													vv := header[key]
													if vv == nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2403
				_go_fuzz_dep_.CoverTab[75312]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2403
				return len(strs) > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2403
				// _ = "end of CoverTab[75312]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2403
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2403
				_go_fuzz_dep_.CoverTab[75313]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2408
				vv, strs = strs[:1:1], strs[1:]
														vv[0] = hf.Value
														header[key] = vv
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2410
				// _ = "end of CoverTab[75313]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2411
				_go_fuzz_dep_.CoverTab[75314]++
														header[key] = append(vv, hf.Value)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2412
				// _ = "end of CoverTab[75314]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2413
			// _ = "end of CoverTab[75311]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2414
		// _ = "end of CoverTab[75305]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2415
	// _ = "end of CoverTab[75292]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2415
	_go_fuzz_dep_.CoverTab[75293]++

											if statusCode >= 100 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2417
		_go_fuzz_dep_.CoverTab[75315]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2417
		return statusCode <= 199
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2417
		// _ = "end of CoverTab[75315]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2417
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2417
		_go_fuzz_dep_.CoverTab[75316]++
												if f.StreamEnded() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2418
			_go_fuzz_dep_.CoverTab[75321]++
													return nil, errors.New("1xx informational response with END_STREAM flag")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2419
			// _ = "end of CoverTab[75321]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2420
			_go_fuzz_dep_.CoverTab[75322]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2420
			// _ = "end of CoverTab[75322]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2420
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2420
		// _ = "end of CoverTab[75316]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2420
		_go_fuzz_dep_.CoverTab[75317]++
												cs.num1xx++
												const max1xxResponses = 5	// arbitrary bound on number of informational responses, same as net/http
												if cs.num1xx > max1xxResponses {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2423
			_go_fuzz_dep_.CoverTab[75323]++
													return nil, errors.New("http2: too many 1xx informational responses")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2424
			// _ = "end of CoverTab[75323]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2425
			_go_fuzz_dep_.CoverTab[75324]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2425
			// _ = "end of CoverTab[75324]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2425
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2425
		// _ = "end of CoverTab[75317]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2425
		_go_fuzz_dep_.CoverTab[75318]++
												if fn := cs.get1xxTraceFunc(); fn != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2426
			_go_fuzz_dep_.CoverTab[75325]++
													if err := fn(statusCode, textproto.MIMEHeader(header)); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2427
				_go_fuzz_dep_.CoverTab[75326]++
														return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2428
				// _ = "end of CoverTab[75326]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2429
				_go_fuzz_dep_.CoverTab[75327]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2429
				// _ = "end of CoverTab[75327]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2429
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2429
			// _ = "end of CoverTab[75325]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2430
			_go_fuzz_dep_.CoverTab[75328]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2430
			// _ = "end of CoverTab[75328]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2430
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2430
		// _ = "end of CoverTab[75318]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2430
		_go_fuzz_dep_.CoverTab[75319]++
												if statusCode == 100 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2431
			_go_fuzz_dep_.CoverTab[75329]++
													traceGot100Continue(cs.trace)
													select {
			case cs.on100 <- struct{}{}:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2434
				_go_fuzz_dep_.CoverTab[75330]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2434
				// _ = "end of CoverTab[75330]"
			default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2435
				_go_fuzz_dep_.CoverTab[75331]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2435
				// _ = "end of CoverTab[75331]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2436
			// _ = "end of CoverTab[75329]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2437
			_go_fuzz_dep_.CoverTab[75332]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2437
			// _ = "end of CoverTab[75332]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2437
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2437
		// _ = "end of CoverTab[75319]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2437
		_go_fuzz_dep_.CoverTab[75320]++
												cs.pastHeaders = false
												return nil, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2439
		// _ = "end of CoverTab[75320]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2440
		_go_fuzz_dep_.CoverTab[75333]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2440
		// _ = "end of CoverTab[75333]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2440
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2440
	// _ = "end of CoverTab[75293]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2440
	_go_fuzz_dep_.CoverTab[75294]++

											res.ContentLength = -1
											if clens := res.Header["Content-Length"]; len(clens) == 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2443
		_go_fuzz_dep_.CoverTab[75334]++
												if cl, err := strconv.ParseUint(clens[0], 10, 63); err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2444
			_go_fuzz_dep_.CoverTab[75335]++
													res.ContentLength = int64(cl)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2445
			// _ = "end of CoverTab[75335]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2446
			_go_fuzz_dep_.CoverTab[75336]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2446
			// _ = "end of CoverTab[75336]"

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2449
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2449
		// _ = "end of CoverTab[75334]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2450
		_go_fuzz_dep_.CoverTab[75337]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2450
		if len(clens) > 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2450
			_go_fuzz_dep_.CoverTab[75338]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2450
			// _ = "end of CoverTab[75338]"

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2453
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2453
			_go_fuzz_dep_.CoverTab[75339]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2453
			if f.StreamEnded() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2453
				_go_fuzz_dep_.CoverTab[75340]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2453
				return !cs.isHead
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2453
				// _ = "end of CoverTab[75340]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2453
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2453
				_go_fuzz_dep_.CoverTab[75341]++
														res.ContentLength = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2454
				// _ = "end of CoverTab[75341]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2455
				_go_fuzz_dep_.CoverTab[75342]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2455
				// _ = "end of CoverTab[75342]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2455
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2455
			// _ = "end of CoverTab[75339]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2455
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2455
		// _ = "end of CoverTab[75337]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2455
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2455
	// _ = "end of CoverTab[75294]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2455
	_go_fuzz_dep_.CoverTab[75295]++

											if cs.isHead {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2457
		_go_fuzz_dep_.CoverTab[75343]++
												res.Body = noBody
												return res, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2459
		// _ = "end of CoverTab[75343]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2460
		_go_fuzz_dep_.CoverTab[75344]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2460
		// _ = "end of CoverTab[75344]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2460
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2460
	// _ = "end of CoverTab[75295]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2460
	_go_fuzz_dep_.CoverTab[75296]++

											if f.StreamEnded() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2462
		_go_fuzz_dep_.CoverTab[75345]++
												if res.ContentLength > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2463
			_go_fuzz_dep_.CoverTab[75347]++
													res.Body = missingBody{}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2464
			// _ = "end of CoverTab[75347]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2465
			_go_fuzz_dep_.CoverTab[75348]++
													res.Body = noBody
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2466
			// _ = "end of CoverTab[75348]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2467
		// _ = "end of CoverTab[75345]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2467
		_go_fuzz_dep_.CoverTab[75346]++
												return res, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2468
		// _ = "end of CoverTab[75346]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2469
		_go_fuzz_dep_.CoverTab[75349]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2469
		// _ = "end of CoverTab[75349]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2469
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2469
	// _ = "end of CoverTab[75296]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2469
	_go_fuzz_dep_.CoverTab[75297]++

											cs.bufPipe.setBuffer(&dataBuffer{expected: res.ContentLength})
											cs.bytesRemain = res.ContentLength
											res.Body = transportResponseBody{cs}

											if cs.requestedGzip && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2475
		_go_fuzz_dep_.CoverTab[75350]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2475
		return asciiEqualFold(res.Header.Get("Content-Encoding"), "gzip")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2475
		// _ = "end of CoverTab[75350]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2475
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2475
		_go_fuzz_dep_.CoverTab[75351]++
												res.Header.Del("Content-Encoding")
												res.Header.Del("Content-Length")
												res.ContentLength = -1
												res.Body = &gzipReader{body: res.Body}
												res.Uncompressed = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2480
		// _ = "end of CoverTab[75351]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2481
		_go_fuzz_dep_.CoverTab[75352]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2481
		// _ = "end of CoverTab[75352]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2481
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2481
	// _ = "end of CoverTab[75297]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2481
	_go_fuzz_dep_.CoverTab[75298]++
											return res, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2482
	// _ = "end of CoverTab[75298]"
}

func (rl *clientConnReadLoop) processTrailers(cs *clientStream, f *MetaHeadersFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2485
	_go_fuzz_dep_.CoverTab[75353]++
											if cs.pastTrailers {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2486
		_go_fuzz_dep_.CoverTab[75358]++

												return ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2488
		// _ = "end of CoverTab[75358]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2489
		_go_fuzz_dep_.CoverTab[75359]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2489
		// _ = "end of CoverTab[75359]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2489
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2489
	// _ = "end of CoverTab[75353]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2489
	_go_fuzz_dep_.CoverTab[75354]++
											cs.pastTrailers = true
											if !f.StreamEnded() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2491
		_go_fuzz_dep_.CoverTab[75360]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2494
		return ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2494
		// _ = "end of CoverTab[75360]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2495
		_go_fuzz_dep_.CoverTab[75361]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2495
		// _ = "end of CoverTab[75361]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2495
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2495
	// _ = "end of CoverTab[75354]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2495
	_go_fuzz_dep_.CoverTab[75355]++
											if len(f.PseudoFields()) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2496
		_go_fuzz_dep_.CoverTab[75362]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2499
		return ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2499
		// _ = "end of CoverTab[75362]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2500
		_go_fuzz_dep_.CoverTab[75363]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2500
		// _ = "end of CoverTab[75363]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2500
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2500
	// _ = "end of CoverTab[75355]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2500
	_go_fuzz_dep_.CoverTab[75356]++

											trailer := make(http.Header)
											for _, hf := range f.RegularFields() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2503
		_go_fuzz_dep_.CoverTab[75364]++
												key := canonicalHeader(hf.Name)
												trailer[key] = append(trailer[key], hf.Value)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2505
		// _ = "end of CoverTab[75364]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2506
	// _ = "end of CoverTab[75356]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2506
	_go_fuzz_dep_.CoverTab[75357]++
											cs.trailer = trailer

											rl.endStream(cs)
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2510
	// _ = "end of CoverTab[75357]"
}

// transportResponseBody is the concrete type of Transport.RoundTrip's
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2513
// Response.Body. It is an io.ReadCloser.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2515
type transportResponseBody struct {
	cs *clientStream
}

func (b transportResponseBody) Read(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2519
	_go_fuzz_dep_.CoverTab[75365]++
											cs := b.cs
											cc := cs.cc

											if cs.readErr != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2523
		_go_fuzz_dep_.CoverTab[75371]++
												return 0, cs.readErr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2524
		// _ = "end of CoverTab[75371]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2525
		_go_fuzz_dep_.CoverTab[75372]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2525
		// _ = "end of CoverTab[75372]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2525
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2525
	// _ = "end of CoverTab[75365]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2525
	_go_fuzz_dep_.CoverTab[75366]++
											n, err = b.cs.bufPipe.Read(p)
											if cs.bytesRemain != -1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2527
		_go_fuzz_dep_.CoverTab[75373]++
												if int64(n) > cs.bytesRemain {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2528
			_go_fuzz_dep_.CoverTab[75375]++
													n = int(cs.bytesRemain)
													if err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2530
				_go_fuzz_dep_.CoverTab[75377]++
														err = errors.New("net/http: server replied with more than declared Content-Length; truncated")
														cs.abortStream(err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2532
				// _ = "end of CoverTab[75377]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2533
				_go_fuzz_dep_.CoverTab[75378]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2533
				// _ = "end of CoverTab[75378]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2533
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2533
			// _ = "end of CoverTab[75375]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2533
			_go_fuzz_dep_.CoverTab[75376]++
													cs.readErr = err
													return int(cs.bytesRemain), err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2535
			// _ = "end of CoverTab[75376]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2536
			_go_fuzz_dep_.CoverTab[75379]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2536
			// _ = "end of CoverTab[75379]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2536
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2536
		// _ = "end of CoverTab[75373]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2536
		_go_fuzz_dep_.CoverTab[75374]++
												cs.bytesRemain -= int64(n)
												if err == io.EOF && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2538
			_go_fuzz_dep_.CoverTab[75380]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2538
			return cs.bytesRemain > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2538
			// _ = "end of CoverTab[75380]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2538
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2538
			_go_fuzz_dep_.CoverTab[75381]++
													err = io.ErrUnexpectedEOF
													cs.readErr = err
													return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2541
			// _ = "end of CoverTab[75381]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2542
			_go_fuzz_dep_.CoverTab[75382]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2542
			// _ = "end of CoverTab[75382]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2542
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2542
		// _ = "end of CoverTab[75374]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2543
		_go_fuzz_dep_.CoverTab[75383]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2543
		// _ = "end of CoverTab[75383]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2543
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2543
	// _ = "end of CoverTab[75366]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2543
	_go_fuzz_dep_.CoverTab[75367]++
											if n == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2544
		_go_fuzz_dep_.CoverTab[75384]++

												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2546
		// _ = "end of CoverTab[75384]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2547
		_go_fuzz_dep_.CoverTab[75385]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2547
		// _ = "end of CoverTab[75385]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2547
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2547
	// _ = "end of CoverTab[75367]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2547
	_go_fuzz_dep_.CoverTab[75368]++

											cc.mu.Lock()
											connAdd := cc.inflow.add(n)
											var streamAdd int32
											if err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2552
		_go_fuzz_dep_.CoverTab[75386]++
												streamAdd = cs.inflow.add(n)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2553
		// _ = "end of CoverTab[75386]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2554
		_go_fuzz_dep_.CoverTab[75387]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2554
		// _ = "end of CoverTab[75387]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2554
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2554
	// _ = "end of CoverTab[75368]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2554
	_go_fuzz_dep_.CoverTab[75369]++
											cc.mu.Unlock()

											if connAdd != 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2557
		_go_fuzz_dep_.CoverTab[75388]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2557
		return streamAdd != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2557
		// _ = "end of CoverTab[75388]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2557
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2557
		_go_fuzz_dep_.CoverTab[75389]++
												cc.wmu.Lock()
												defer cc.wmu.Unlock()
												if connAdd != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2560
			_go_fuzz_dep_.CoverTab[75392]++
													cc.fr.WriteWindowUpdate(0, mustUint31(connAdd))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2561
			// _ = "end of CoverTab[75392]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2562
			_go_fuzz_dep_.CoverTab[75393]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2562
			// _ = "end of CoverTab[75393]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2562
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2562
		// _ = "end of CoverTab[75389]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2562
		_go_fuzz_dep_.CoverTab[75390]++
												if streamAdd != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2563
			_go_fuzz_dep_.CoverTab[75394]++
													cc.fr.WriteWindowUpdate(cs.ID, mustUint31(streamAdd))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2564
			// _ = "end of CoverTab[75394]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2565
			_go_fuzz_dep_.CoverTab[75395]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2565
			// _ = "end of CoverTab[75395]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2565
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2565
		// _ = "end of CoverTab[75390]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2565
		_go_fuzz_dep_.CoverTab[75391]++
												cc.bw.Flush()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2566
		// _ = "end of CoverTab[75391]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2567
		_go_fuzz_dep_.CoverTab[75396]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2567
		// _ = "end of CoverTab[75396]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2567
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2567
	// _ = "end of CoverTab[75369]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2567
	_go_fuzz_dep_.CoverTab[75370]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2568
	// _ = "end of CoverTab[75370]"
}

var errClosedResponseBody = errors.New("http2: response body closed")

func (b transportResponseBody) Close() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2573
	_go_fuzz_dep_.CoverTab[75397]++
											cs := b.cs
											cc := cs.cc

											cs.bufPipe.BreakWithError(errClosedResponseBody)
											cs.abortStream(errClosedResponseBody)

											unread := cs.bufPipe.Len()
											if unread > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2581
		_go_fuzz_dep_.CoverTab[75400]++
												cc.mu.Lock()

												connAdd := cc.inflow.add(unread)
												cc.mu.Unlock()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2589
		cc.wmu.Lock()

		if connAdd > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2591
			_go_fuzz_dep_.CoverTab[75402]++
													cc.fr.WriteWindowUpdate(0, uint32(connAdd))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2592
			// _ = "end of CoverTab[75402]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2593
			_go_fuzz_dep_.CoverTab[75403]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2593
			// _ = "end of CoverTab[75403]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2593
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2593
		// _ = "end of CoverTab[75400]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2593
		_go_fuzz_dep_.CoverTab[75401]++
												cc.bw.Flush()
												cc.wmu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2595
		// _ = "end of CoverTab[75401]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2596
		_go_fuzz_dep_.CoverTab[75404]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2596
		// _ = "end of CoverTab[75404]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2596
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2596
	// _ = "end of CoverTab[75397]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2596
	_go_fuzz_dep_.CoverTab[75398]++

											select {
	case <-cs.donec:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2599
		_go_fuzz_dep_.CoverTab[75405]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2599
		// _ = "end of CoverTab[75405]"
	case <-cs.ctx.Done():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2600
		_go_fuzz_dep_.CoverTab[75406]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2604
		return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2604
		// _ = "end of CoverTab[75406]"
	case <-cs.reqCancel:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2605
		_go_fuzz_dep_.CoverTab[75407]++
												return errRequestCanceled
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2606
		// _ = "end of CoverTab[75407]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2607
	// _ = "end of CoverTab[75398]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2607
	_go_fuzz_dep_.CoverTab[75399]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2608
	// _ = "end of CoverTab[75399]"
}

func (rl *clientConnReadLoop) processData(f *DataFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2611
	_go_fuzz_dep_.CoverTab[75408]++
											cc := rl.cc
											cs := rl.streamByID(f.StreamID)
											data := f.Data()
											if cs == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2615
		_go_fuzz_dep_.CoverTab[75414]++
												cc.mu.Lock()
												neverSent := cc.nextStreamID
												cc.mu.Unlock()
												if f.StreamID >= neverSent {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2619
			_go_fuzz_dep_.CoverTab[75417]++

													cc.logf("http2: Transport received unsolicited DATA frame; closing connection")
													return ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2622
			// _ = "end of CoverTab[75417]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2623
			_go_fuzz_dep_.CoverTab[75418]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2623
			// _ = "end of CoverTab[75418]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2623
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2623
		// _ = "end of CoverTab[75414]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2623
		_go_fuzz_dep_.CoverTab[75415]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2630
		if f.Length > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2630
			_go_fuzz_dep_.CoverTab[75419]++
													cc.mu.Lock()
													ok := cc.inflow.take(f.Length)
													connAdd := cc.inflow.add(int(f.Length))
													cc.mu.Unlock()
													if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2635
				_go_fuzz_dep_.CoverTab[75421]++
														return ConnectionError(ErrCodeFlowControl)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2636
				// _ = "end of CoverTab[75421]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2637
				_go_fuzz_dep_.CoverTab[75422]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2637
				// _ = "end of CoverTab[75422]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2637
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2637
			// _ = "end of CoverTab[75419]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2637
			_go_fuzz_dep_.CoverTab[75420]++
													if connAdd > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2638
				_go_fuzz_dep_.CoverTab[75423]++
														cc.wmu.Lock()
														cc.fr.WriteWindowUpdate(0, uint32(connAdd))
														cc.bw.Flush()
														cc.wmu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2642
				// _ = "end of CoverTab[75423]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2643
				_go_fuzz_dep_.CoverTab[75424]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2643
				// _ = "end of CoverTab[75424]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2643
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2643
			// _ = "end of CoverTab[75420]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2644
			_go_fuzz_dep_.CoverTab[75425]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2644
			// _ = "end of CoverTab[75425]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2644
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2644
		// _ = "end of CoverTab[75415]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2644
		_go_fuzz_dep_.CoverTab[75416]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2645
		// _ = "end of CoverTab[75416]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2646
		_go_fuzz_dep_.CoverTab[75426]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2646
		// _ = "end of CoverTab[75426]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2646
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2646
	// _ = "end of CoverTab[75408]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2646
	_go_fuzz_dep_.CoverTab[75409]++
											if cs.readClosed {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2647
		_go_fuzz_dep_.CoverTab[75427]++
												cc.logf("protocol error: received DATA after END_STREAM")
												rl.endStreamError(cs, StreamError{
			StreamID:	f.StreamID,
			Code:		ErrCodeProtocol,
		})
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2653
		// _ = "end of CoverTab[75427]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2654
		_go_fuzz_dep_.CoverTab[75428]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2654
		// _ = "end of CoverTab[75428]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2654
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2654
	// _ = "end of CoverTab[75409]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2654
	_go_fuzz_dep_.CoverTab[75410]++
											if !cs.firstByte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2655
		_go_fuzz_dep_.CoverTab[75429]++
												cc.logf("protocol error: received DATA before a HEADERS frame")
												rl.endStreamError(cs, StreamError{
			StreamID:	f.StreamID,
			Code:		ErrCodeProtocol,
		})
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2661
		// _ = "end of CoverTab[75429]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2662
		_go_fuzz_dep_.CoverTab[75430]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2662
		// _ = "end of CoverTab[75430]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2662
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2662
	// _ = "end of CoverTab[75410]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2662
	_go_fuzz_dep_.CoverTab[75411]++
											if f.Length > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2663
		_go_fuzz_dep_.CoverTab[75431]++
												if cs.isHead && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2664
			_go_fuzz_dep_.CoverTab[75438]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2664
			return len(data) > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2664
			// _ = "end of CoverTab[75438]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2664
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2664
			_go_fuzz_dep_.CoverTab[75439]++
													cc.logf("protocol error: received DATA on a HEAD request")
													rl.endStreamError(cs, StreamError{
				StreamID:	f.StreamID,
				Code:		ErrCodeProtocol,
			})
													return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2670
			// _ = "end of CoverTab[75439]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2671
			_go_fuzz_dep_.CoverTab[75440]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2671
			// _ = "end of CoverTab[75440]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2671
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2671
		// _ = "end of CoverTab[75431]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2671
		_go_fuzz_dep_.CoverTab[75432]++

												cc.mu.Lock()
												if !takeInflows(&cc.inflow, &cs.inflow, f.Length) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2674
			_go_fuzz_dep_.CoverTab[75441]++
													cc.mu.Unlock()
													return ConnectionError(ErrCodeFlowControl)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2676
			// _ = "end of CoverTab[75441]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2677
			_go_fuzz_dep_.CoverTab[75442]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2677
			// _ = "end of CoverTab[75442]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2677
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2677
		// _ = "end of CoverTab[75432]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2677
		_go_fuzz_dep_.CoverTab[75433]++
		// Return any padded flow control now, since we won't
		// refund it later on body reads.
		var refund int
		if pad := int(f.Length) - len(data); pad > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2681
			_go_fuzz_dep_.CoverTab[75443]++
													refund += pad
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2682
			// _ = "end of CoverTab[75443]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2683
			_go_fuzz_dep_.CoverTab[75444]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2683
			// _ = "end of CoverTab[75444]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2683
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2683
		// _ = "end of CoverTab[75433]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2683
		_go_fuzz_dep_.CoverTab[75434]++

												didReset := false
												var err error
												if len(data) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2687
			_go_fuzz_dep_.CoverTab[75445]++
													if _, err = cs.bufPipe.Write(data); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2688
				_go_fuzz_dep_.CoverTab[75446]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2691
				didReset = true
														refund += len(data)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2692
				// _ = "end of CoverTab[75446]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2693
				_go_fuzz_dep_.CoverTab[75447]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2693
				// _ = "end of CoverTab[75447]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2693
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2693
			// _ = "end of CoverTab[75445]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2694
			_go_fuzz_dep_.CoverTab[75448]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2694
			// _ = "end of CoverTab[75448]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2694
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2694
		// _ = "end of CoverTab[75434]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2694
		_go_fuzz_dep_.CoverTab[75435]++

												sendConn := cc.inflow.add(refund)
												var sendStream int32
												if !didReset {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2698
			_go_fuzz_dep_.CoverTab[75449]++
													sendStream = cs.inflow.add(refund)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2699
			// _ = "end of CoverTab[75449]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2700
			_go_fuzz_dep_.CoverTab[75450]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2700
			// _ = "end of CoverTab[75450]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2700
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2700
		// _ = "end of CoverTab[75435]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2700
		_go_fuzz_dep_.CoverTab[75436]++
												cc.mu.Unlock()

												if sendConn > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2703
			_go_fuzz_dep_.CoverTab[75451]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2703
			return sendStream > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2703
			// _ = "end of CoverTab[75451]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2703
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2703
			_go_fuzz_dep_.CoverTab[75452]++
													cc.wmu.Lock()
													if sendConn > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2705
				_go_fuzz_dep_.CoverTab[75455]++
														cc.fr.WriteWindowUpdate(0, uint32(sendConn))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2706
				// _ = "end of CoverTab[75455]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2707
				_go_fuzz_dep_.CoverTab[75456]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2707
				// _ = "end of CoverTab[75456]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2707
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2707
			// _ = "end of CoverTab[75452]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2707
			_go_fuzz_dep_.CoverTab[75453]++
													if sendStream > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2708
				_go_fuzz_dep_.CoverTab[75457]++
														cc.fr.WriteWindowUpdate(cs.ID, uint32(sendStream))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2709
				// _ = "end of CoverTab[75457]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2710
				_go_fuzz_dep_.CoverTab[75458]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2710
				// _ = "end of CoverTab[75458]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2710
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2710
			// _ = "end of CoverTab[75453]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2710
			_go_fuzz_dep_.CoverTab[75454]++
													cc.bw.Flush()
													cc.wmu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2712
			// _ = "end of CoverTab[75454]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2713
			_go_fuzz_dep_.CoverTab[75459]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2713
			// _ = "end of CoverTab[75459]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2713
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2713
		// _ = "end of CoverTab[75436]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2713
		_go_fuzz_dep_.CoverTab[75437]++

												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2715
			_go_fuzz_dep_.CoverTab[75460]++
													rl.endStreamError(cs, err)
													return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2717
			// _ = "end of CoverTab[75460]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2718
			_go_fuzz_dep_.CoverTab[75461]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2718
			// _ = "end of CoverTab[75461]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2718
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2718
		// _ = "end of CoverTab[75437]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2719
		_go_fuzz_dep_.CoverTab[75462]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2719
		// _ = "end of CoverTab[75462]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2719
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2719
	// _ = "end of CoverTab[75411]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2719
	_go_fuzz_dep_.CoverTab[75412]++

											if f.StreamEnded() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2721
		_go_fuzz_dep_.CoverTab[75463]++
												rl.endStream(cs)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2722
		// _ = "end of CoverTab[75463]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2723
		_go_fuzz_dep_.CoverTab[75464]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2723
		// _ = "end of CoverTab[75464]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2723
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2723
	// _ = "end of CoverTab[75412]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2723
	_go_fuzz_dep_.CoverTab[75413]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2724
	// _ = "end of CoverTab[75413]"
}

func (rl *clientConnReadLoop) endStream(cs *clientStream) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2727
	_go_fuzz_dep_.CoverTab[75465]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2730
	if !cs.readClosed {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2730
		_go_fuzz_dep_.CoverTab[75466]++
												cs.readClosed = true

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2736
		rl.cc.mu.Lock()
												defer rl.cc.mu.Unlock()
												cs.bufPipe.closeWithErrorAndCode(io.EOF, cs.copyTrailers)
												close(cs.peerClosed)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2739
		// _ = "end of CoverTab[75466]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2740
		_go_fuzz_dep_.CoverTab[75467]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2740
		// _ = "end of CoverTab[75467]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2740
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2740
	// _ = "end of CoverTab[75465]"
}

func (rl *clientConnReadLoop) endStreamError(cs *clientStream, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2743
	_go_fuzz_dep_.CoverTab[75468]++
											cs.readAborted = true
											cs.abortStream(err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2745
	// _ = "end of CoverTab[75468]"
}

func (rl *clientConnReadLoop) streamByID(id uint32) *clientStream {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2748
	_go_fuzz_dep_.CoverTab[75469]++
											rl.cc.mu.Lock()
											defer rl.cc.mu.Unlock()
											cs := rl.cc.streams[id]
											if cs != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2752
		_go_fuzz_dep_.CoverTab[75471]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2752
		return !cs.readAborted
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2752
		// _ = "end of CoverTab[75471]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2752
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2752
		_go_fuzz_dep_.CoverTab[75472]++
												return cs
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2753
		// _ = "end of CoverTab[75472]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2754
		_go_fuzz_dep_.CoverTab[75473]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2754
		// _ = "end of CoverTab[75473]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2754
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2754
	// _ = "end of CoverTab[75469]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2754
	_go_fuzz_dep_.CoverTab[75470]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2755
	// _ = "end of CoverTab[75470]"
}

func (cs *clientStream) copyTrailers() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2758
	_go_fuzz_dep_.CoverTab[75474]++
											for k, vv := range cs.trailer {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2759
		_go_fuzz_dep_.CoverTab[75475]++
												t := cs.resTrailer
												if *t == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2761
			_go_fuzz_dep_.CoverTab[75477]++
													*t = make(http.Header)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2762
			// _ = "end of CoverTab[75477]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2763
			_go_fuzz_dep_.CoverTab[75478]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2763
			// _ = "end of CoverTab[75478]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2763
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2763
		// _ = "end of CoverTab[75475]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2763
		_go_fuzz_dep_.CoverTab[75476]++
												(*t)[k] = vv
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2764
		// _ = "end of CoverTab[75476]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2765
	// _ = "end of CoverTab[75474]"
}

func (rl *clientConnReadLoop) processGoAway(f *GoAwayFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2768
	_go_fuzz_dep_.CoverTab[75479]++
											cc := rl.cc
											cc.t.connPool().MarkDead(cc)
											if f.ErrCode != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2771
		_go_fuzz_dep_.CoverTab[75481]++

												cc.vlogf("transport got GOAWAY with error code = %v", f.ErrCode)
												if fn := cc.t.CountError; fn != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2774
			_go_fuzz_dep_.CoverTab[75482]++
													fn("recv_goaway_" + f.ErrCode.stringToken())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2775
			// _ = "end of CoverTab[75482]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2776
			_go_fuzz_dep_.CoverTab[75483]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2776
			// _ = "end of CoverTab[75483]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2776
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2776
		// _ = "end of CoverTab[75481]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2777
		_go_fuzz_dep_.CoverTab[75484]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2777
		// _ = "end of CoverTab[75484]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2777
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2777
	// _ = "end of CoverTab[75479]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2777
	_go_fuzz_dep_.CoverTab[75480]++
											cc.setGoAway(f)
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2779
	// _ = "end of CoverTab[75480]"
}

func (rl *clientConnReadLoop) processSettings(f *SettingsFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2782
	_go_fuzz_dep_.CoverTab[75485]++
											cc := rl.cc

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2786
	cc.wmu.Lock()
	defer cc.wmu.Unlock()

	if err := rl.processSettingsNoWrite(f); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2789
		_go_fuzz_dep_.CoverTab[75488]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2790
		// _ = "end of CoverTab[75488]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2791
		_go_fuzz_dep_.CoverTab[75489]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2791
		// _ = "end of CoverTab[75489]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2791
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2791
	// _ = "end of CoverTab[75485]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2791
	_go_fuzz_dep_.CoverTab[75486]++
											if !f.IsAck() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2792
		_go_fuzz_dep_.CoverTab[75490]++
												cc.fr.WriteSettingsAck()
												cc.bw.Flush()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2794
		// _ = "end of CoverTab[75490]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2795
		_go_fuzz_dep_.CoverTab[75491]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2795
		// _ = "end of CoverTab[75491]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2795
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2795
	// _ = "end of CoverTab[75486]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2795
	_go_fuzz_dep_.CoverTab[75487]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2796
	// _ = "end of CoverTab[75487]"
}

func (rl *clientConnReadLoop) processSettingsNoWrite(f *SettingsFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2799
	_go_fuzz_dep_.CoverTab[75492]++
											cc := rl.cc
											cc.mu.Lock()
											defer cc.mu.Unlock()

											if f.IsAck() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2804
		_go_fuzz_dep_.CoverTab[75497]++
												if cc.wantSettingsAck {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2805
			_go_fuzz_dep_.CoverTab[75499]++
													cc.wantSettingsAck = false
													return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2807
			// _ = "end of CoverTab[75499]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2808
			_go_fuzz_dep_.CoverTab[75500]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2808
			// _ = "end of CoverTab[75500]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2808
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2808
		// _ = "end of CoverTab[75497]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2808
		_go_fuzz_dep_.CoverTab[75498]++
												return ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2809
		// _ = "end of CoverTab[75498]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2810
		_go_fuzz_dep_.CoverTab[75501]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2810
		// _ = "end of CoverTab[75501]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2810
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2810
	// _ = "end of CoverTab[75492]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2810
	_go_fuzz_dep_.CoverTab[75493]++

											var seenMaxConcurrentStreams bool
											err := f.ForeachSetting(func(s Setting) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2813
		_go_fuzz_dep_.CoverTab[75502]++
												switch s.ID {
		case SettingMaxFrameSize:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2815
			_go_fuzz_dep_.CoverTab[75504]++
													cc.maxFrameSize = s.Val
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2816
			// _ = "end of CoverTab[75504]"
		case SettingMaxConcurrentStreams:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2817
			_go_fuzz_dep_.CoverTab[75505]++
													cc.maxConcurrentStreams = s.Val
													seenMaxConcurrentStreams = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2819
			// _ = "end of CoverTab[75505]"
		case SettingMaxHeaderListSize:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2820
			_go_fuzz_dep_.CoverTab[75506]++
													cc.peerMaxHeaderListSize = uint64(s.Val)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2821
			// _ = "end of CoverTab[75506]"
		case SettingInitialWindowSize:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2822
			_go_fuzz_dep_.CoverTab[75507]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2827
			if s.Val > math.MaxInt32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2827
				_go_fuzz_dep_.CoverTab[75512]++
														return ConnectionError(ErrCodeFlowControl)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2828
				// _ = "end of CoverTab[75512]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2829
				_go_fuzz_dep_.CoverTab[75513]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2829
				// _ = "end of CoverTab[75513]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2829
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2829
			// _ = "end of CoverTab[75507]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2829
			_go_fuzz_dep_.CoverTab[75508]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2834
			delta := int32(s.Val) - int32(cc.initialWindowSize)
			for _, cs := range cc.streams {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2835
				_go_fuzz_dep_.CoverTab[75514]++
														cs.flow.add(delta)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2836
				// _ = "end of CoverTab[75514]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2837
			// _ = "end of CoverTab[75508]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2837
			_go_fuzz_dep_.CoverTab[75509]++
													cc.cond.Broadcast()

													cc.initialWindowSize = s.Val
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2840
			// _ = "end of CoverTab[75509]"
		case SettingHeaderTableSize:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2841
			_go_fuzz_dep_.CoverTab[75510]++
													cc.henc.SetMaxDynamicTableSize(s.Val)
													cc.peerMaxHeaderTableSize = s.Val
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2843
			// _ = "end of CoverTab[75510]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2844
			_go_fuzz_dep_.CoverTab[75511]++
													cc.vlogf("Unhandled Setting: %v", s)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2845
			// _ = "end of CoverTab[75511]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2846
		// _ = "end of CoverTab[75502]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2846
		_go_fuzz_dep_.CoverTab[75503]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2847
		// _ = "end of CoverTab[75503]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2848
	// _ = "end of CoverTab[75493]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2848
	_go_fuzz_dep_.CoverTab[75494]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2849
		_go_fuzz_dep_.CoverTab[75515]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2850
		// _ = "end of CoverTab[75515]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2851
		_go_fuzz_dep_.CoverTab[75516]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2851
		// _ = "end of CoverTab[75516]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2851
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2851
	// _ = "end of CoverTab[75494]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2851
	_go_fuzz_dep_.CoverTab[75495]++

											if !cc.seenSettings {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2853
		_go_fuzz_dep_.CoverTab[75517]++
												if !seenMaxConcurrentStreams {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2854
			_go_fuzz_dep_.CoverTab[75519]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2859
			cc.maxConcurrentStreams = defaultMaxConcurrentStreams
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2859
			// _ = "end of CoverTab[75519]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2860
			_go_fuzz_dep_.CoverTab[75520]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2860
			// _ = "end of CoverTab[75520]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2860
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2860
		// _ = "end of CoverTab[75517]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2860
		_go_fuzz_dep_.CoverTab[75518]++
												cc.seenSettings = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2861
		// _ = "end of CoverTab[75518]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2862
		_go_fuzz_dep_.CoverTab[75521]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2862
		// _ = "end of CoverTab[75521]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2862
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2862
	// _ = "end of CoverTab[75495]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2862
	_go_fuzz_dep_.CoverTab[75496]++

											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2864
	// _ = "end of CoverTab[75496]"
}

func (rl *clientConnReadLoop) processWindowUpdate(f *WindowUpdateFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2867
	_go_fuzz_dep_.CoverTab[75522]++
											cc := rl.cc
											cs := rl.streamByID(f.StreamID)
											if f.StreamID != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2870
		_go_fuzz_dep_.CoverTab[75526]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2870
		return cs == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2870
		// _ = "end of CoverTab[75526]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2870
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2870
		_go_fuzz_dep_.CoverTab[75527]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2871
		// _ = "end of CoverTab[75527]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2872
		_go_fuzz_dep_.CoverTab[75528]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2872
		// _ = "end of CoverTab[75528]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2872
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2872
	// _ = "end of CoverTab[75522]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2872
	_go_fuzz_dep_.CoverTab[75523]++

											cc.mu.Lock()
											defer cc.mu.Unlock()

											fl := &cc.flow
											if cs != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2878
		_go_fuzz_dep_.CoverTab[75529]++
												fl = &cs.flow
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2879
		// _ = "end of CoverTab[75529]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2880
		_go_fuzz_dep_.CoverTab[75530]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2880
		// _ = "end of CoverTab[75530]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2880
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2880
	// _ = "end of CoverTab[75523]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2880
	_go_fuzz_dep_.CoverTab[75524]++
											if !fl.add(int32(f.Increment)) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2881
		_go_fuzz_dep_.CoverTab[75531]++
												return ConnectionError(ErrCodeFlowControl)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2882
		// _ = "end of CoverTab[75531]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2883
		_go_fuzz_dep_.CoverTab[75532]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2883
		// _ = "end of CoverTab[75532]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2883
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2883
	// _ = "end of CoverTab[75524]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2883
	_go_fuzz_dep_.CoverTab[75525]++
											cc.cond.Broadcast()
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2885
	// _ = "end of CoverTab[75525]"
}

func (rl *clientConnReadLoop) processResetStream(f *RSTStreamFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2888
	_go_fuzz_dep_.CoverTab[75533]++
											cs := rl.streamByID(f.StreamID)
											if cs == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2890
		_go_fuzz_dep_.CoverTab[75537]++

												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2892
		// _ = "end of CoverTab[75537]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2893
		_go_fuzz_dep_.CoverTab[75538]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2893
		// _ = "end of CoverTab[75538]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2893
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2893
	// _ = "end of CoverTab[75533]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2893
	_go_fuzz_dep_.CoverTab[75534]++
											serr := streamError(cs.ID, f.ErrCode)
											serr.Cause = errFromPeer
											if f.ErrCode == ErrCodeProtocol {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2896
		_go_fuzz_dep_.CoverTab[75539]++
												rl.cc.SetDoNotReuse()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2897
		// _ = "end of CoverTab[75539]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2898
		_go_fuzz_dep_.CoverTab[75540]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2898
		// _ = "end of CoverTab[75540]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2898
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2898
	// _ = "end of CoverTab[75534]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2898
	_go_fuzz_dep_.CoverTab[75535]++
											if fn := cs.cc.t.CountError; fn != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2899
		_go_fuzz_dep_.CoverTab[75541]++
												fn("recv_rststream_" + f.ErrCode.stringToken())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2900
		// _ = "end of CoverTab[75541]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2901
		_go_fuzz_dep_.CoverTab[75542]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2901
		// _ = "end of CoverTab[75542]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2901
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2901
	// _ = "end of CoverTab[75535]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2901
	_go_fuzz_dep_.CoverTab[75536]++
											cs.abortStream(serr)

											cs.bufPipe.CloseWithError(serr)
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2905
	// _ = "end of CoverTab[75536]"
}

// Ping sends a PING frame to the server and waits for the ack.
func (cc *ClientConn) Ping(ctx context.Context) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2909
	_go_fuzz_dep_.CoverTab[75543]++
											c := make(chan struct{})
	// Generate a random payload
	var p [8]byte
	for {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2913
		_go_fuzz_dep_.CoverTab[75546]++
												if _, err := rand.Read(p[:]); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2914
			_go_fuzz_dep_.CoverTab[75549]++
													return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2915
			// _ = "end of CoverTab[75549]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2916
			_go_fuzz_dep_.CoverTab[75550]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2916
			// _ = "end of CoverTab[75550]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2916
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2916
		// _ = "end of CoverTab[75546]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2916
		_go_fuzz_dep_.CoverTab[75547]++
												cc.mu.Lock()

												if _, found := cc.pings[p]; !found {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2919
			_go_fuzz_dep_.CoverTab[75551]++
													cc.pings[p] = c
													cc.mu.Unlock()
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2922
			// _ = "end of CoverTab[75551]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2923
			_go_fuzz_dep_.CoverTab[75552]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2923
			// _ = "end of CoverTab[75552]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2923
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2923
		// _ = "end of CoverTab[75547]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2923
		_go_fuzz_dep_.CoverTab[75548]++
												cc.mu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2924
		// _ = "end of CoverTab[75548]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2925
	// _ = "end of CoverTab[75543]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2925
	_go_fuzz_dep_.CoverTab[75544]++
											errc := make(chan error, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2926
	_curRoutineNum70_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2926
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum70_)
											go func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2927
		_go_fuzz_dep_.CoverTab[75553]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2927
		defer func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2927
			_go_fuzz_dep_.CoverTab[75555]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2927
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum70_)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2927
			// _ = "end of CoverTab[75555]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2927
		}()
												cc.wmu.Lock()
												defer cc.wmu.Unlock()
												if err := cc.fr.WritePing(false, p); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2930
			_go_fuzz_dep_.CoverTab[75556]++
													errc <- err
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2932
			// _ = "end of CoverTab[75556]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2933
			_go_fuzz_dep_.CoverTab[75557]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2933
			// _ = "end of CoverTab[75557]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2933
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2933
		// _ = "end of CoverTab[75553]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2933
		_go_fuzz_dep_.CoverTab[75554]++
												if err := cc.bw.Flush(); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2934
			_go_fuzz_dep_.CoverTab[75558]++
													errc <- err
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2936
			// _ = "end of CoverTab[75558]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2937
			_go_fuzz_dep_.CoverTab[75559]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2937
			// _ = "end of CoverTab[75559]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2937
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2937
		// _ = "end of CoverTab[75554]"
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2938
	// _ = "end of CoverTab[75544]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2938
	_go_fuzz_dep_.CoverTab[75545]++
											select {
	case <-c:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2940
		_go_fuzz_dep_.CoverTab[75560]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2941
		// _ = "end of CoverTab[75560]"
	case err := <-errc:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2942
		_go_fuzz_dep_.CoverTab[75561]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2943
		// _ = "end of CoverTab[75561]"
	case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2944
		_go_fuzz_dep_.CoverTab[75562]++
												return ctx.Err()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2945
		// _ = "end of CoverTab[75562]"
	case <-cc.readerDone:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2946
		_go_fuzz_dep_.CoverTab[75563]++

												return cc.readerErr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2948
		// _ = "end of CoverTab[75563]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2949
	// _ = "end of CoverTab[75545]"
}

func (rl *clientConnReadLoop) processPing(f *PingFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2952
	_go_fuzz_dep_.CoverTab[75564]++
											if f.IsAck() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2953
		_go_fuzz_dep_.CoverTab[75567]++
												cc := rl.cc
												cc.mu.Lock()
												defer cc.mu.Unlock()

												if c, ok := cc.pings[f.Data]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2958
			_go_fuzz_dep_.CoverTab[75569]++
													close(c)
													delete(cc.pings, f.Data)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2960
			// _ = "end of CoverTab[75569]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2961
			_go_fuzz_dep_.CoverTab[75570]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2961
			// _ = "end of CoverTab[75570]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2961
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2961
		// _ = "end of CoverTab[75567]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2961
		_go_fuzz_dep_.CoverTab[75568]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2962
		// _ = "end of CoverTab[75568]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2963
		_go_fuzz_dep_.CoverTab[75571]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2963
		// _ = "end of CoverTab[75571]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2963
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2963
	// _ = "end of CoverTab[75564]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2963
	_go_fuzz_dep_.CoverTab[75565]++
											cc := rl.cc
											cc.wmu.Lock()
											defer cc.wmu.Unlock()
											if err := cc.fr.WritePing(true, f.Data); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2967
		_go_fuzz_dep_.CoverTab[75572]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2968
		// _ = "end of CoverTab[75572]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2969
		_go_fuzz_dep_.CoverTab[75573]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2969
		// _ = "end of CoverTab[75573]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2969
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2969
	// _ = "end of CoverTab[75565]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2969
	_go_fuzz_dep_.CoverTab[75566]++
											return cc.bw.Flush()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2970
	// _ = "end of CoverTab[75566]"
}

func (rl *clientConnReadLoop) processPushPromise(f *PushPromiseFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2973
	_go_fuzz_dep_.CoverTab[75574]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2981
	return ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2981
	// _ = "end of CoverTab[75574]"
}

func (cc *ClientConn) writeStreamReset(streamID uint32, code ErrCode, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2984
	_go_fuzz_dep_.CoverTab[75575]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2989
	cc.wmu.Lock()
											cc.fr.WriteRSTStream(streamID, code)
											cc.bw.Flush()
											cc.wmu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:2992
	// _ = "end of CoverTab[75575]"
}

var (
	errResponseHeaderListSize	= errors.New("http2: response header list larger than advertised limit")
	errRequestHeaderListSize	= errors.New("http2: request header list larger than peer's advertised limit")
)

func (cc *ClientConn) logf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3000
	_go_fuzz_dep_.CoverTab[75576]++
											cc.t.logf(format, args...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3001
	// _ = "end of CoverTab[75576]"
}

func (cc *ClientConn) vlogf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3004
	_go_fuzz_dep_.CoverTab[75577]++
											cc.t.vlogf(format, args...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3005
	// _ = "end of CoverTab[75577]"
}

func (t *Transport) vlogf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3008
	_go_fuzz_dep_.CoverTab[75578]++
											if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3009
		_go_fuzz_dep_.CoverTab[75579]++
												t.logf(format, args...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3010
		// _ = "end of CoverTab[75579]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3011
		_go_fuzz_dep_.CoverTab[75580]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3011
		// _ = "end of CoverTab[75580]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3011
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3011
	// _ = "end of CoverTab[75578]"
}

func (t *Transport) logf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3014
	_go_fuzz_dep_.CoverTab[75581]++
											log.Printf(format, args...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3015
	// _ = "end of CoverTab[75581]"
}

var noBody io.ReadCloser = noBodyReader{}

type noBodyReader struct{}

func (noBodyReader) Close() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3022
	_go_fuzz_dep_.CoverTab[75582]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3022
	return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3022
	// _ = "end of CoverTab[75582]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3022
}
func (noBodyReader) Read([]byte) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3023
	_go_fuzz_dep_.CoverTab[75583]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3023
	return 0, io.EOF
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3023
	// _ = "end of CoverTab[75583]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3023
}

type missingBody struct{}

func (missingBody) Close() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3027
	_go_fuzz_dep_.CoverTab[75584]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3027
	return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3027
	// _ = "end of CoverTab[75584]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3027
}
func (missingBody) Read([]byte) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3028
	_go_fuzz_dep_.CoverTab[75585]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3028
	return 0, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3028
	// _ = "end of CoverTab[75585]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3028
}

func strSliceContains(ss []string, s string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3030
	_go_fuzz_dep_.CoverTab[75586]++
											for _, v := range ss {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3031
		_go_fuzz_dep_.CoverTab[75588]++
												if v == s {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3032
			_go_fuzz_dep_.CoverTab[75589]++
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3033
			// _ = "end of CoverTab[75589]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3034
			_go_fuzz_dep_.CoverTab[75590]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3034
			// _ = "end of CoverTab[75590]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3034
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3034
		// _ = "end of CoverTab[75588]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3035
	// _ = "end of CoverTab[75586]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3035
	_go_fuzz_dep_.CoverTab[75587]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3036
	// _ = "end of CoverTab[75587]"
}

type erringRoundTripper struct{ err error }

func (rt erringRoundTripper) RoundTripErr() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3041
	_go_fuzz_dep_.CoverTab[75591]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3041
	return rt.err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3041
	// _ = "end of CoverTab[75591]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3041
}
func (rt erringRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3042
	_go_fuzz_dep_.CoverTab[75592]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3042
	return nil, rt.err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3042
	// _ = "end of CoverTab[75592]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3042
}

// gzipReader wraps a response body so it can lazily
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3044
// call gzip.NewReader on the first call to Read
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3046
type gzipReader struct {
	_	incomparable
	body	io.ReadCloser	// underlying Response.Body
	zr	*gzip.Reader	// lazily-initialized gzip reader
	zerr	error		// sticky error
}

func (gz *gzipReader) Read(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3053
	_go_fuzz_dep_.CoverTab[75593]++
											if gz.zerr != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3054
		_go_fuzz_dep_.CoverTab[75596]++
												return 0, gz.zerr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3055
		// _ = "end of CoverTab[75596]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3056
		_go_fuzz_dep_.CoverTab[75597]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3056
		// _ = "end of CoverTab[75597]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3056
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3056
	// _ = "end of CoverTab[75593]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3056
	_go_fuzz_dep_.CoverTab[75594]++
											if gz.zr == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3057
		_go_fuzz_dep_.CoverTab[75598]++
												gz.zr, err = gzip.NewReader(gz.body)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3059
			_go_fuzz_dep_.CoverTab[75599]++
													gz.zerr = err
													return 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3061
			// _ = "end of CoverTab[75599]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3062
			_go_fuzz_dep_.CoverTab[75600]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3062
			// _ = "end of CoverTab[75600]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3062
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3062
		// _ = "end of CoverTab[75598]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3063
		_go_fuzz_dep_.CoverTab[75601]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3063
		// _ = "end of CoverTab[75601]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3063
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3063
	// _ = "end of CoverTab[75594]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3063
	_go_fuzz_dep_.CoverTab[75595]++
											return gz.zr.Read(p)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3064
	// _ = "end of CoverTab[75595]"
}

func (gz *gzipReader) Close() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3067
	_go_fuzz_dep_.CoverTab[75602]++
											if err := gz.body.Close(); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3068
		_go_fuzz_dep_.CoverTab[75604]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3069
		// _ = "end of CoverTab[75604]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3070
		_go_fuzz_dep_.CoverTab[75605]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3070
		// _ = "end of CoverTab[75605]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3070
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3070
	// _ = "end of CoverTab[75602]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3070
	_go_fuzz_dep_.CoverTab[75603]++
											gz.zerr = fs.ErrClosed
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3072
	// _ = "end of CoverTab[75603]"
}

type errorReader struct{ err error }

func (r errorReader) Read(p []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3077
	_go_fuzz_dep_.CoverTab[75606]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3077
	return 0, r.err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3077
	// _ = "end of CoverTab[75606]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3077
}

// isConnectionCloseRequest reports whether req should use its own
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3079
// connection for a single request and then close the connection.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3081
func isConnectionCloseRequest(req *http.Request) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3081
	_go_fuzz_dep_.CoverTab[75607]++
											return req.Close || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3082
		_go_fuzz_dep_.CoverTab[75608]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3082
		return httpguts.HeaderValuesContainsToken(req.Header["Connection"], "close")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3082
		// _ = "end of CoverTab[75608]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3082
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3082
	// _ = "end of CoverTab[75607]"
}

// registerHTTPSProtocol calls Transport.RegisterProtocol but
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3085
// converting panics into errors.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3087
func registerHTTPSProtocol(t *http.Transport, rt noDialH2RoundTripper) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3087
	_go_fuzz_dep_.CoverTab[75609]++
											defer func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3088
		_go_fuzz_dep_.CoverTab[75611]++
												if e := recover(); e != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3089
			_go_fuzz_dep_.CoverTab[75612]++
													err = fmt.Errorf("%v", e)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3090
			// _ = "end of CoverTab[75612]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3091
			_go_fuzz_dep_.CoverTab[75613]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3091
			// _ = "end of CoverTab[75613]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3091
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3091
		// _ = "end of CoverTab[75611]"
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3092
	// _ = "end of CoverTab[75609]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3092
	_go_fuzz_dep_.CoverTab[75610]++
											t.RegisterProtocol("https", rt)
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3094
	// _ = "end of CoverTab[75610]"
}

// noDialH2RoundTripper is a RoundTripper which only tries to complete the request
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3097
// if there's already has a cached connection to the host.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3097
// (The field is exported so it can be accessed via reflect from net/http; tested
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3097
// by TestNoDialH2RoundTripperType)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3101
type noDialH2RoundTripper struct{ *Transport }

func (rt noDialH2RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3103
	_go_fuzz_dep_.CoverTab[75614]++
											res, err := rt.Transport.RoundTrip(req)
											if isNoCachedConnError(err) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3105
		_go_fuzz_dep_.CoverTab[75616]++
												return nil, http.ErrSkipAltProtocol
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3106
		// _ = "end of CoverTab[75616]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3107
		_go_fuzz_dep_.CoverTab[75617]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3107
		// _ = "end of CoverTab[75617]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3107
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3107
	// _ = "end of CoverTab[75614]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3107
	_go_fuzz_dep_.CoverTab[75615]++
											return res, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3108
	// _ = "end of CoverTab[75615]"
}

func (t *Transport) idleConnTimeout() time.Duration {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3111
	_go_fuzz_dep_.CoverTab[75618]++
											if t.t1 != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3112
		_go_fuzz_dep_.CoverTab[75620]++
												return t.t1.IdleConnTimeout
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3113
		// _ = "end of CoverTab[75620]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3114
		_go_fuzz_dep_.CoverTab[75621]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3114
		// _ = "end of CoverTab[75621]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3114
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3114
	// _ = "end of CoverTab[75618]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3114
	_go_fuzz_dep_.CoverTab[75619]++
											return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3115
	// _ = "end of CoverTab[75619]"
}

func traceGetConn(req *http.Request, hostPort string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3118
	_go_fuzz_dep_.CoverTab[75622]++
											trace := httptrace.ContextClientTrace(req.Context())
											if trace == nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3120
		_go_fuzz_dep_.CoverTab[75624]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3120
		return trace.GetConn == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3120
		// _ = "end of CoverTab[75624]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3120
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3120
		_go_fuzz_dep_.CoverTab[75625]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3121
		// _ = "end of CoverTab[75625]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3122
		_go_fuzz_dep_.CoverTab[75626]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3122
		// _ = "end of CoverTab[75626]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3122
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3122
	// _ = "end of CoverTab[75622]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3122
	_go_fuzz_dep_.CoverTab[75623]++
											trace.GetConn(hostPort)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3123
	// _ = "end of CoverTab[75623]"
}

func traceGotConn(req *http.Request, cc *ClientConn, reused bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3126
	_go_fuzz_dep_.CoverTab[75627]++
											trace := httptrace.ContextClientTrace(req.Context())
											if trace == nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3128
		_go_fuzz_dep_.CoverTab[75630]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3128
		return trace.GotConn == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3128
		// _ = "end of CoverTab[75630]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3128
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3128
		_go_fuzz_dep_.CoverTab[75631]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3129
		// _ = "end of CoverTab[75631]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3130
		_go_fuzz_dep_.CoverTab[75632]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3130
		// _ = "end of CoverTab[75632]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3130
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3130
	// _ = "end of CoverTab[75627]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3130
	_go_fuzz_dep_.CoverTab[75628]++
											ci := httptrace.GotConnInfo{Conn: cc.tconn}
											ci.Reused = reused
											cc.mu.Lock()
											ci.WasIdle = len(cc.streams) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3134
		_go_fuzz_dep_.CoverTab[75633]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3134
		return reused
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3134
		// _ = "end of CoverTab[75633]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3134
	}()
											if ci.WasIdle && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3135
		_go_fuzz_dep_.CoverTab[75634]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3135
		return !cc.lastActive.IsZero()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3135
		// _ = "end of CoverTab[75634]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3135
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3135
		_go_fuzz_dep_.CoverTab[75635]++
												ci.IdleTime = time.Since(cc.lastActive)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3136
		// _ = "end of CoverTab[75635]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3137
		_go_fuzz_dep_.CoverTab[75636]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3137
		// _ = "end of CoverTab[75636]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3137
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3137
	// _ = "end of CoverTab[75628]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3137
	_go_fuzz_dep_.CoverTab[75629]++
											cc.mu.Unlock()

											trace.GotConn(ci)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3140
	// _ = "end of CoverTab[75629]"
}

func traceWroteHeaders(trace *httptrace.ClientTrace) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3143
	_go_fuzz_dep_.CoverTab[75637]++
											if trace != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3144
		_go_fuzz_dep_.CoverTab[75638]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3144
		return trace.WroteHeaders != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3144
		// _ = "end of CoverTab[75638]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3144
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3144
		_go_fuzz_dep_.CoverTab[75639]++
												trace.WroteHeaders()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3145
		// _ = "end of CoverTab[75639]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3146
		_go_fuzz_dep_.CoverTab[75640]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3146
		// _ = "end of CoverTab[75640]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3146
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3146
	// _ = "end of CoverTab[75637]"
}

func traceGot100Continue(trace *httptrace.ClientTrace) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3149
	_go_fuzz_dep_.CoverTab[75641]++
											if trace != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3150
		_go_fuzz_dep_.CoverTab[75642]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3150
		return trace.Got100Continue != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3150
		// _ = "end of CoverTab[75642]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3150
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3150
		_go_fuzz_dep_.CoverTab[75643]++
												trace.Got100Continue()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3151
		// _ = "end of CoverTab[75643]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3152
		_go_fuzz_dep_.CoverTab[75644]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3152
		// _ = "end of CoverTab[75644]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3152
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3152
	// _ = "end of CoverTab[75641]"
}

func traceWait100Continue(trace *httptrace.ClientTrace) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3155
	_go_fuzz_dep_.CoverTab[75645]++
											if trace != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3156
		_go_fuzz_dep_.CoverTab[75646]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3156
		return trace.Wait100Continue != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3156
		// _ = "end of CoverTab[75646]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3156
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3156
		_go_fuzz_dep_.CoverTab[75647]++
												trace.Wait100Continue()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3157
		// _ = "end of CoverTab[75647]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3158
		_go_fuzz_dep_.CoverTab[75648]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3158
		// _ = "end of CoverTab[75648]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3158
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3158
	// _ = "end of CoverTab[75645]"
}

func traceWroteRequest(trace *httptrace.ClientTrace, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3161
	_go_fuzz_dep_.CoverTab[75649]++
											if trace != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3162
		_go_fuzz_dep_.CoverTab[75650]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3162
		return trace.WroteRequest != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3162
		// _ = "end of CoverTab[75650]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3162
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3162
		_go_fuzz_dep_.CoverTab[75651]++
												trace.WroteRequest(httptrace.WroteRequestInfo{Err: err})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3163
		// _ = "end of CoverTab[75651]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3164
		_go_fuzz_dep_.CoverTab[75652]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3164
		// _ = "end of CoverTab[75652]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3164
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3164
	// _ = "end of CoverTab[75649]"
}

func traceFirstResponseByte(trace *httptrace.ClientTrace) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3167
	_go_fuzz_dep_.CoverTab[75653]++
											if trace != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3168
		_go_fuzz_dep_.CoverTab[75654]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3168
		return trace.GotFirstResponseByte != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3168
		// _ = "end of CoverTab[75654]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3168
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3168
		_go_fuzz_dep_.CoverTab[75655]++
												trace.GotFirstResponseByte()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3169
		// _ = "end of CoverTab[75655]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3170
		_go_fuzz_dep_.CoverTab[75656]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3170
		// _ = "end of CoverTab[75656]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3170
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3170
	// _ = "end of CoverTab[75653]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3171
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/transport.go:3171
var _ = _go_fuzz_dep_.CoverTab
