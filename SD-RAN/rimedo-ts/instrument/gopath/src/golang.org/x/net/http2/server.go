// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO: turn off the serve goroutine when idle, so
// an idle conn only has the readFrames goroutine active. (which could
// also be optimized probably to pin less memory in crypto/tls). This
// would involve tracking when the serve goroutine is active (atomic
// int32 read/CAS probably?) and starting it up when frames arrive,
// and shutting it down when all handlers exit. the occasional PING
// packets could use time.AfterFunc to call sc.wakeStartServeLoop()
// (which is a no-op if already running) and then queue the PING write
// as normal. The serve loop would then exit in most cases (if no
// Handlers running) and not be woken up again until the PING packet
// returns.

// TODO (maybe): add a mechanism for Handlers to going into
// half-closed-local mode (rw.(io.Closer) test?) but not exit their
// handler, and continue to be able to read from the
// Request.Body. This would be a somewhat semantic change from HTTP/1
// (or at least what we expose in net/http), so I'd probably want to
// add it there too. For now, this package says that returning from
// the Handler ServeHTTP function means you're both done reading and
// done writing, without a way to stop just one or the other.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:26
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:26
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:26
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:26
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:26
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:26
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:26
)

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/http/httpguts"
	"golang.org/x/net/http2/hpack"
)

const (
	prefaceTimeout		= 10 * time.Second
	firstSettingsTimeout	= 2 * time.Second	// should be in-flight with preface anyway
	handlerChunkWriteSize	= 4 << 10
	defaultMaxStreams	= 250	// TODO: make this 100 as the GFE seems to?
	maxQueuedControlFrames	= 10000
)

var (
	errClientDisconnected	= errors.New("client disconnected")
	errClosedBody		= errors.New("body closed by handler")
	errHandlerComplete	= errors.New("http2: request body closed due to handler exiting")
	errStreamClosed		= errors.New("http2: stream closed")
)

var responseWriterStatePool = sync.Pool{
	New: func() interface{} {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:70
		_go_fuzz_dep_.CoverTab[73210]++
											rws := &responseWriterState{}
											rws.bw = bufio.NewWriterSize(chunkWriter{rws}, handlerChunkWriteSize)
											return rws
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:73
		// _ = "end of CoverTab[73210]"
	},
}

// Test hooks.
var (
	testHookOnConn		func()
	testHookGetServerConn	func(*serverConn)
	testHookOnPanicMu	*sync.Mutex	// nil except in tests
	testHookOnPanic		func(sc *serverConn, panicVal interface{}) (rePanic bool)
)

// Server is an HTTP/2 server.
type Server struct {
	// MaxHandlers limits the number of http.Handler ServeHTTP goroutines
	// which may run at a time over all connections.
	// Negative or zero no limit.
	// TODO: implement
	MaxHandlers	int

	// MaxConcurrentStreams optionally specifies the number of
	// concurrent streams that each client may have open at a
	// time. This is unrelated to the number of http.Handler goroutines
	// which may be active globally, which is MaxHandlers.
	// If zero, MaxConcurrentStreams defaults to at least 100, per
	// the HTTP/2 spec's recommendations.
	MaxConcurrentStreams	uint32

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

	// MaxReadFrameSize optionally specifies the largest frame
	// this server is willing to read. A valid value is between
	// 16k and 16M, inclusive. If zero or otherwise invalid, a
	// default value is used.
	MaxReadFrameSize	uint32

	// PermitProhibitedCipherSuites, if true, permits the use of
	// cipher suites prohibited by the HTTP/2 spec.
	PermitProhibitedCipherSuites	bool

	// IdleTimeout specifies how long until idle clients should be
	// closed with a GOAWAY frame. PING frames are not considered
	// activity for the purposes of IdleTimeout.
	IdleTimeout	time.Duration

	// MaxUploadBufferPerConnection is the size of the initial flow
	// control window for each connections. The HTTP/2 spec does not
	// allow this to be smaller than 65535 or larger than 2^32-1.
	// If the value is outside this range, a default value will be
	// used instead.
	MaxUploadBufferPerConnection	int32

	// MaxUploadBufferPerStream is the size of the initial flow control
	// window for each stream. The HTTP/2 spec does not allow this to
	// be larger than 2^32-1. If the value is zero or larger than the
	// maximum, a default value will be used instead.
	MaxUploadBufferPerStream	int32

	// NewWriteScheduler constructs a write scheduler for a connection.
	// If nil, a default scheduler is chosen.
	NewWriteScheduler	func() WriteScheduler

	// CountError, if non-nil, is called on HTTP/2 server errors.
	// It's intended to increment a metric for monitoring, such
	// as an expvar or Prometheus metric.
	// The errType consists of only ASCII word characters.
	CountError	func(errType string)

	// Internal state. This is a pointer (rather than embedded directly)
	// so that we don't embed a Mutex in this struct, which will make the
	// struct non-copyable, which might break some callers.
	state	*serverInternalState
}

func (s *Server) initialConnRecvWindowSize() int32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:158
	_go_fuzz_dep_.CoverTab[73211]++
										if s.MaxUploadBufferPerConnection >= initialWindowSize {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:159
		_go_fuzz_dep_.CoverTab[73213]++
											return s.MaxUploadBufferPerConnection
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:160
		// _ = "end of CoverTab[73213]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:161
		_go_fuzz_dep_.CoverTab[73214]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:161
		// _ = "end of CoverTab[73214]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:161
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:161
	// _ = "end of CoverTab[73211]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:161
	_go_fuzz_dep_.CoverTab[73212]++
										return 1 << 20
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:162
	// _ = "end of CoverTab[73212]"
}

func (s *Server) initialStreamRecvWindowSize() int32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:165
	_go_fuzz_dep_.CoverTab[73215]++
										if s.MaxUploadBufferPerStream > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:166
		_go_fuzz_dep_.CoverTab[73217]++
											return s.MaxUploadBufferPerStream
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:167
		// _ = "end of CoverTab[73217]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:168
		_go_fuzz_dep_.CoverTab[73218]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:168
		// _ = "end of CoverTab[73218]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:168
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:168
	// _ = "end of CoverTab[73215]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:168
	_go_fuzz_dep_.CoverTab[73216]++
										return 1 << 20
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:169
	// _ = "end of CoverTab[73216]"
}

func (s *Server) maxReadFrameSize() uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:172
	_go_fuzz_dep_.CoverTab[73219]++
										if v := s.MaxReadFrameSize; v >= minMaxFrameSize && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:173
		_go_fuzz_dep_.CoverTab[73221]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:173
		return v <= maxFrameSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:173
		// _ = "end of CoverTab[73221]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:173
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:173
		_go_fuzz_dep_.CoverTab[73222]++
											return v
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:174
		// _ = "end of CoverTab[73222]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:175
		_go_fuzz_dep_.CoverTab[73223]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:175
		// _ = "end of CoverTab[73223]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:175
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:175
	// _ = "end of CoverTab[73219]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:175
	_go_fuzz_dep_.CoverTab[73220]++
										return defaultMaxReadFrameSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:176
	// _ = "end of CoverTab[73220]"
}

func (s *Server) maxConcurrentStreams() uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:179
	_go_fuzz_dep_.CoverTab[73224]++
										if v := s.MaxConcurrentStreams; v > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:180
		_go_fuzz_dep_.CoverTab[73226]++
											return v
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:181
		// _ = "end of CoverTab[73226]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:182
		_go_fuzz_dep_.CoverTab[73227]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:182
		// _ = "end of CoverTab[73227]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:182
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:182
	// _ = "end of CoverTab[73224]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:182
	_go_fuzz_dep_.CoverTab[73225]++
										return defaultMaxStreams
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:183
	// _ = "end of CoverTab[73225]"
}

func (s *Server) maxDecoderHeaderTableSize() uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:186
	_go_fuzz_dep_.CoverTab[73228]++
										if v := s.MaxDecoderHeaderTableSize; v > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:187
		_go_fuzz_dep_.CoverTab[73230]++
											return v
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:188
		// _ = "end of CoverTab[73230]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:189
		_go_fuzz_dep_.CoverTab[73231]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:189
		// _ = "end of CoverTab[73231]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:189
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:189
	// _ = "end of CoverTab[73228]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:189
	_go_fuzz_dep_.CoverTab[73229]++
										return initialHeaderTableSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:190
	// _ = "end of CoverTab[73229]"
}

func (s *Server) maxEncoderHeaderTableSize() uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:193
	_go_fuzz_dep_.CoverTab[73232]++
										if v := s.MaxEncoderHeaderTableSize; v > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:194
		_go_fuzz_dep_.CoverTab[73234]++
											return v
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:195
		// _ = "end of CoverTab[73234]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:196
		_go_fuzz_dep_.CoverTab[73235]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:196
		// _ = "end of CoverTab[73235]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:196
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:196
	// _ = "end of CoverTab[73232]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:196
	_go_fuzz_dep_.CoverTab[73233]++
										return initialHeaderTableSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:197
	// _ = "end of CoverTab[73233]"
}

// maxQueuedControlFrames is the maximum number of control frames like
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:200
// SETTINGS, PING and RST_STREAM that will be queued for writing before
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:200
// the connection is closed to prevent memory exhaustion attacks.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:203
func (s *Server) maxQueuedControlFrames() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:203
	_go_fuzz_dep_.CoverTab[73236]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:206
	return maxQueuedControlFrames
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:206
	// _ = "end of CoverTab[73236]"
}

type serverInternalState struct {
	mu		sync.Mutex
	activeConns	map[*serverConn]struct{}
}

func (s *serverInternalState) registerConn(sc *serverConn) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:214
	_go_fuzz_dep_.CoverTab[73237]++
										if s == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:215
		_go_fuzz_dep_.CoverTab[73239]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:216
		// _ = "end of CoverTab[73239]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:217
		_go_fuzz_dep_.CoverTab[73240]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:217
		// _ = "end of CoverTab[73240]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:217
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:217
	// _ = "end of CoverTab[73237]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:217
	_go_fuzz_dep_.CoverTab[73238]++
										s.mu.Lock()
										s.activeConns[sc] = struct{}{}
										s.mu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:220
	// _ = "end of CoverTab[73238]"
}

func (s *serverInternalState) unregisterConn(sc *serverConn) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:223
	_go_fuzz_dep_.CoverTab[73241]++
										if s == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:224
		_go_fuzz_dep_.CoverTab[73243]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:225
		// _ = "end of CoverTab[73243]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:226
		_go_fuzz_dep_.CoverTab[73244]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:226
		// _ = "end of CoverTab[73244]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:226
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:226
	// _ = "end of CoverTab[73241]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:226
	_go_fuzz_dep_.CoverTab[73242]++
										s.mu.Lock()
										delete(s.activeConns, sc)
										s.mu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:229
	// _ = "end of CoverTab[73242]"
}

func (s *serverInternalState) startGracefulShutdown() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:232
	_go_fuzz_dep_.CoverTab[73245]++
										if s == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:233
		_go_fuzz_dep_.CoverTab[73248]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:234
		// _ = "end of CoverTab[73248]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:235
		_go_fuzz_dep_.CoverTab[73249]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:235
		// _ = "end of CoverTab[73249]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:235
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:235
	// _ = "end of CoverTab[73245]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:235
	_go_fuzz_dep_.CoverTab[73246]++
										s.mu.Lock()
										for sc := range s.activeConns {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:237
		_go_fuzz_dep_.CoverTab[73250]++
											sc.startGracefulShutdown()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:238
		// _ = "end of CoverTab[73250]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:239
	// _ = "end of CoverTab[73246]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:239
	_go_fuzz_dep_.CoverTab[73247]++
										s.mu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:240
	// _ = "end of CoverTab[73247]"
}

// ConfigureServer adds HTTP/2 support to a net/http Server.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:243
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:243
// The configuration conf may be nil.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:243
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:243
// ConfigureServer must be called before s begins serving.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:248
func ConfigureServer(s *http.Server, conf *Server) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:248
	_go_fuzz_dep_.CoverTab[73251]++
										if s == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:249
		_go_fuzz_dep_.CoverTab[73260]++
											panic("nil *http.Server")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:250
		// _ = "end of CoverTab[73260]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:251
		_go_fuzz_dep_.CoverTab[73261]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:251
		// _ = "end of CoverTab[73261]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:251
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:251
	// _ = "end of CoverTab[73251]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:251
	_go_fuzz_dep_.CoverTab[73252]++
										if conf == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:252
		_go_fuzz_dep_.CoverTab[73262]++
											conf = new(Server)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:253
		// _ = "end of CoverTab[73262]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:254
		_go_fuzz_dep_.CoverTab[73263]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:254
		// _ = "end of CoverTab[73263]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:254
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:254
	// _ = "end of CoverTab[73252]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:254
	_go_fuzz_dep_.CoverTab[73253]++
										conf.state = &serverInternalState{activeConns: make(map[*serverConn]struct{})}
										if h1, h2 := s, conf; h2.IdleTimeout == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:256
		_go_fuzz_dep_.CoverTab[73264]++
											if h1.IdleTimeout != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:257
			_go_fuzz_dep_.CoverTab[73265]++
												h2.IdleTimeout = h1.IdleTimeout
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:258
			// _ = "end of CoverTab[73265]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:259
			_go_fuzz_dep_.CoverTab[73266]++
												h2.IdleTimeout = h1.ReadTimeout
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:260
			// _ = "end of CoverTab[73266]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:261
		// _ = "end of CoverTab[73264]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:262
		_go_fuzz_dep_.CoverTab[73267]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:262
		// _ = "end of CoverTab[73267]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:262
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:262
	// _ = "end of CoverTab[73253]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:262
	_go_fuzz_dep_.CoverTab[73254]++
										s.RegisterOnShutdown(conf.state.startGracefulShutdown)

										if s.TLSConfig == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:265
		_go_fuzz_dep_.CoverTab[73268]++
											s.TLSConfig = new(tls.Config)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:266
		// _ = "end of CoverTab[73268]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:267
		_go_fuzz_dep_.CoverTab[73269]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:267
		if s.TLSConfig.CipherSuites != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:267
			_go_fuzz_dep_.CoverTab[73270]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:267
			return s.TLSConfig.MinVersion < tls.VersionTLS13
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:267
			// _ = "end of CoverTab[73270]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:267
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:267
			_go_fuzz_dep_.CoverTab[73271]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:271
			haveRequired := false
			for _, cs := range s.TLSConfig.CipherSuites {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:272
				_go_fuzz_dep_.CoverTab[73273]++
													switch cs {
				case tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:277
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:277
					_go_fuzz_dep_.CoverTab[73274]++
														haveRequired = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:278
					// _ = "end of CoverTab[73274]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:278
				default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:278
					_go_fuzz_dep_.CoverTab[73275]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:278
					// _ = "end of CoverTab[73275]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:279
				// _ = "end of CoverTab[73273]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:280
			// _ = "end of CoverTab[73271]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:280
			_go_fuzz_dep_.CoverTab[73272]++
												if !haveRequired {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:281
				_go_fuzz_dep_.CoverTab[73276]++
													return fmt.Errorf("http2: TLSConfig.CipherSuites is missing an HTTP/2-required AES_128_GCM_SHA256 cipher (need at least one of TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 or TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256)")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:282
				// _ = "end of CoverTab[73276]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:283
				_go_fuzz_dep_.CoverTab[73277]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:283
				// _ = "end of CoverTab[73277]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:283
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:283
			// _ = "end of CoverTab[73272]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:284
			_go_fuzz_dep_.CoverTab[73278]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:284
			// _ = "end of CoverTab[73278]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:284
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:284
		// _ = "end of CoverTab[73269]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:284
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:284
	// _ = "end of CoverTab[73254]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:284
	_go_fuzz_dep_.CoverTab[73255]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:293
	s.TLSConfig.PreferServerCipherSuites = true

	if !strSliceContains(s.TLSConfig.NextProtos, NextProtoTLS) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:295
		_go_fuzz_dep_.CoverTab[73279]++
											s.TLSConfig.NextProtos = append(s.TLSConfig.NextProtos, NextProtoTLS)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:296
		// _ = "end of CoverTab[73279]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:297
		_go_fuzz_dep_.CoverTab[73280]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:297
		// _ = "end of CoverTab[73280]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:297
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:297
	// _ = "end of CoverTab[73255]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:297
	_go_fuzz_dep_.CoverTab[73256]++
										if !strSliceContains(s.TLSConfig.NextProtos, "http/1.1") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:298
		_go_fuzz_dep_.CoverTab[73281]++
											s.TLSConfig.NextProtos = append(s.TLSConfig.NextProtos, "http/1.1")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:299
		// _ = "end of CoverTab[73281]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:300
		_go_fuzz_dep_.CoverTab[73282]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:300
		// _ = "end of CoverTab[73282]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:300
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:300
	// _ = "end of CoverTab[73256]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:300
	_go_fuzz_dep_.CoverTab[73257]++

										if s.TLSNextProto == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:302
		_go_fuzz_dep_.CoverTab[73283]++
											s.TLSNextProto = map[string]func(*http.Server, *tls.Conn, http.Handler){}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:303
		// _ = "end of CoverTab[73283]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:304
		_go_fuzz_dep_.CoverTab[73284]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:304
		// _ = "end of CoverTab[73284]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:304
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:304
	// _ = "end of CoverTab[73257]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:304
	_go_fuzz_dep_.CoverTab[73258]++
										protoHandler := func(hs *http.Server, c *tls.Conn, h http.Handler) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:305
		_go_fuzz_dep_.CoverTab[73285]++
											if testHookOnConn != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:306
			_go_fuzz_dep_.CoverTab[73288]++
												testHookOnConn()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:307
			// _ = "end of CoverTab[73288]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:308
			_go_fuzz_dep_.CoverTab[73289]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:308
			// _ = "end of CoverTab[73289]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:308
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:308
		// _ = "end of CoverTab[73285]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:308
		_go_fuzz_dep_.CoverTab[73286]++
		// The TLSNextProto interface predates contexts, so
		// the net/http package passes down its per-connection
		// base context via an exported but unadvertised
		// method on the Handler. This is for internal
		// net/http<=>http2 use only.
		var ctx context.Context
		type baseContexter interface {
			BaseContext() context.Context
		}
		if bc, ok := h.(baseContexter); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:318
			_go_fuzz_dep_.CoverTab[73290]++
												ctx = bc.BaseContext()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:319
			// _ = "end of CoverTab[73290]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:320
			_go_fuzz_dep_.CoverTab[73291]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:320
			// _ = "end of CoverTab[73291]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:320
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:320
		// _ = "end of CoverTab[73286]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:320
		_go_fuzz_dep_.CoverTab[73287]++
											conf.ServeConn(c, &ServeConnOpts{
			Context:	ctx,
			Handler:	h,
			BaseConfig:	hs,
		})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:325
		// _ = "end of CoverTab[73287]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:326
	// _ = "end of CoverTab[73258]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:326
	_go_fuzz_dep_.CoverTab[73259]++
										s.TLSNextProto[NextProtoTLS] = protoHandler
										return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:328
	// _ = "end of CoverTab[73259]"
}

// ServeConnOpts are options for the Server.ServeConn method.
type ServeConnOpts struct {
	// Context is the base context to use.
	// If nil, context.Background is used.
	Context	context.Context

	// BaseConfig optionally sets the base configuration
	// for values. If nil, defaults are used.
	BaseConfig	*http.Server

	// Handler specifies which handler to use for processing
	// requests. If nil, BaseConfig.Handler is used. If BaseConfig
	// or BaseConfig.Handler is nil, http.DefaultServeMux is used.
	Handler	http.Handler

	// UpgradeRequest is an initial request received on a connection
	// undergoing an h2c upgrade. The request body must have been
	// completely read from the connection before calling ServeConn,
	// and the 101 Switching Protocols response written.
	UpgradeRequest	*http.Request

	// Settings is the decoded contents of the HTTP2-Settings header
	// in an h2c upgrade request.
	Settings	[]byte

	// SawClientPreface is set if the HTTP/2 connection preface
	// has already been read from the connection.
	SawClientPreface	bool
}

func (o *ServeConnOpts) context() context.Context {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:361
	_go_fuzz_dep_.CoverTab[73292]++
										if o != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:362
		_go_fuzz_dep_.CoverTab[73294]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:362
		return o.Context != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:362
		// _ = "end of CoverTab[73294]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:362
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:362
		_go_fuzz_dep_.CoverTab[73295]++
											return o.Context
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:363
		// _ = "end of CoverTab[73295]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:364
		_go_fuzz_dep_.CoverTab[73296]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:364
		// _ = "end of CoverTab[73296]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:364
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:364
	// _ = "end of CoverTab[73292]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:364
	_go_fuzz_dep_.CoverTab[73293]++
										return context.Background()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:365
	// _ = "end of CoverTab[73293]"
}

func (o *ServeConnOpts) baseConfig() *http.Server {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:368
	_go_fuzz_dep_.CoverTab[73297]++
										if o != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:369
		_go_fuzz_dep_.CoverTab[73299]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:369
		return o.BaseConfig != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:369
		// _ = "end of CoverTab[73299]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:369
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:369
		_go_fuzz_dep_.CoverTab[73300]++
											return o.BaseConfig
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:370
		// _ = "end of CoverTab[73300]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:371
		_go_fuzz_dep_.CoverTab[73301]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:371
		// _ = "end of CoverTab[73301]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:371
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:371
	// _ = "end of CoverTab[73297]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:371
	_go_fuzz_dep_.CoverTab[73298]++
										return new(http.Server)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:372
	// _ = "end of CoverTab[73298]"
}

func (o *ServeConnOpts) handler() http.Handler {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:375
	_go_fuzz_dep_.CoverTab[73302]++
										if o != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:376
		_go_fuzz_dep_.CoverTab[73304]++
											if o.Handler != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:377
			_go_fuzz_dep_.CoverTab[73306]++
												return o.Handler
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:378
			// _ = "end of CoverTab[73306]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:379
			_go_fuzz_dep_.CoverTab[73307]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:379
			// _ = "end of CoverTab[73307]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:379
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:379
		// _ = "end of CoverTab[73304]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:379
		_go_fuzz_dep_.CoverTab[73305]++
											if o.BaseConfig != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:380
			_go_fuzz_dep_.CoverTab[73308]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:380
			return o.BaseConfig.Handler != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:380
			// _ = "end of CoverTab[73308]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:380
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:380
			_go_fuzz_dep_.CoverTab[73309]++
												return o.BaseConfig.Handler
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:381
			// _ = "end of CoverTab[73309]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:382
			_go_fuzz_dep_.CoverTab[73310]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:382
			// _ = "end of CoverTab[73310]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:382
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:382
		// _ = "end of CoverTab[73305]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:383
		_go_fuzz_dep_.CoverTab[73311]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:383
		// _ = "end of CoverTab[73311]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:383
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:383
	// _ = "end of CoverTab[73302]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:383
	_go_fuzz_dep_.CoverTab[73303]++
										return http.DefaultServeMux
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:384
	// _ = "end of CoverTab[73303]"
}

// ServeConn serves HTTP/2 requests on the provided connection and
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:387
// blocks until the connection is no longer readable.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:387
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:387
// ServeConn starts speaking HTTP/2 assuming that c has not had any
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:387
// reads or writes. It writes its initial settings frame and expects
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:387
// to be able to read the preface and settings frame from the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:387
// client. If c has a ConnectionState method like a *tls.Conn, the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:387
// ConnectionState is used to verify the TLS ciphersuite and to set
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:387
// the Request.TLS field in Handlers.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:387
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:387
// ServeConn does not support h2c by itself. Any h2c support must be
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:387
// implemented in terms of providing a suitably-behaving net.Conn.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:387
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:387
// The opts parameter is optional. If nil, default values are used.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:401
func (s *Server) ServeConn(c net.Conn, opts *ServeConnOpts) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:401
	_go_fuzz_dep_.CoverTab[73312]++
										baseCtx, cancel := serverConnBaseContext(c, opts)
										defer cancel()

										sc := &serverConn{
		srv:				s,
		hs:				opts.baseConfig(),
		conn:				c,
		baseCtx:			baseCtx,
		remoteAddrStr:			c.RemoteAddr().String(),
		bw:				newBufferedWriter(c),
		handler:			opts.handler(),
		streams:			make(map[uint32]*stream),
		readFrameCh:			make(chan readFrameResult),
		wantWriteFrameCh:		make(chan FrameWriteRequest, 8),
		serveMsgCh:			make(chan interface{}, 8),
		wroteFrameCh:			make(chan frameWriteResult, 1),
		bodyReadCh:			make(chan bodyReadMsg),
		doneServing:			make(chan struct{}),
		clientMaxStreams:		math.MaxUint32,
		advMaxStreams:			s.maxConcurrentStreams(),
		initialStreamSendWindowSize:	initialWindowSize,
		maxFrameSize:			initialMaxFrameSize,
		serveG:				newGoroutineLock(),
		pushEnabled:			true,
		sawClientPreface:		opts.SawClientPreface,
	}

										s.state.registerConn(sc)
										defer s.state.unregisterConn(sc)

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:437
	if sc.hs.WriteTimeout != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:437
		_go_fuzz_dep_.CoverTab[73320]++
											sc.conn.SetWriteDeadline(time.Time{})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:438
		// _ = "end of CoverTab[73320]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:439
		_go_fuzz_dep_.CoverTab[73321]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:439
		// _ = "end of CoverTab[73321]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:439
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:439
	// _ = "end of CoverTab[73312]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:439
	_go_fuzz_dep_.CoverTab[73313]++

										if s.NewWriteScheduler != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:441
		_go_fuzz_dep_.CoverTab[73322]++
											sc.writeSched = s.NewWriteScheduler()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:442
		// _ = "end of CoverTab[73322]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:443
		_go_fuzz_dep_.CoverTab[73323]++
											sc.writeSched = NewPriorityWriteScheduler(nil)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:444
		// _ = "end of CoverTab[73323]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:445
	// _ = "end of CoverTab[73313]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:445
	_go_fuzz_dep_.CoverTab[73314]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:450
	sc.flow.add(initialWindowSize)
	sc.inflow.init(initialWindowSize)
	sc.hpackEncoder = hpack.NewEncoder(&sc.headerWriteBuf)
	sc.hpackEncoder.SetMaxDynamicTableSizeLimit(s.maxEncoderHeaderTableSize())

	fr := NewFramer(sc.bw, c)
	if s.CountError != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:456
		_go_fuzz_dep_.CoverTab[73324]++
											fr.countError = s.CountError
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:457
		// _ = "end of CoverTab[73324]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:458
		_go_fuzz_dep_.CoverTab[73325]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:458
		// _ = "end of CoverTab[73325]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:458
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:458
	// _ = "end of CoverTab[73314]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:458
	_go_fuzz_dep_.CoverTab[73315]++
										fr.ReadMetaHeaders = hpack.NewDecoder(s.maxDecoderHeaderTableSize(), nil)
										fr.MaxHeaderListSize = sc.maxHeaderListSize()
										fr.SetMaxReadFrameSize(s.maxReadFrameSize())
										sc.framer = fr

										if tc, ok := c.(connectionStater); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:464
		_go_fuzz_dep_.CoverTab[73326]++
											sc.tlsState = new(tls.ConnectionState)
											*sc.tlsState = tc.ConnectionState()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:477
		if sc.tlsState.Version < tls.VersionTLS12 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:477
			_go_fuzz_dep_.CoverTab[73329]++
												sc.rejectConn(ErrCodeInadequateSecurity, "TLS version too low")
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:479
			// _ = "end of CoverTab[73329]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:480
			_go_fuzz_dep_.CoverTab[73330]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:480
			// _ = "end of CoverTab[73330]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:480
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:480
		// _ = "end of CoverTab[73326]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:480
		_go_fuzz_dep_.CoverTab[73327]++

											if sc.tlsState.ServerName == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:482
			_go_fuzz_dep_.CoverTab[73331]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:482
			// _ = "end of CoverTab[73331]"

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:492
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:492
			_go_fuzz_dep_.CoverTab[73332]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:492
			// _ = "end of CoverTab[73332]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:492
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:492
		// _ = "end of CoverTab[73327]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:492
		_go_fuzz_dep_.CoverTab[73328]++

											if !s.PermitProhibitedCipherSuites && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:494
			_go_fuzz_dep_.CoverTab[73333]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:494
			return isBadCipher(sc.tlsState.CipherSuite)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:494
			// _ = "end of CoverTab[73333]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:494
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:494
			_go_fuzz_dep_.CoverTab[73334]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:505
			sc.rejectConn(ErrCodeInadequateSecurity, fmt.Sprintf("Prohibited TLS 1.2 Cipher Suite: %x", sc.tlsState.CipherSuite))
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:506
			// _ = "end of CoverTab[73334]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:507
			_go_fuzz_dep_.CoverTab[73335]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:507
			// _ = "end of CoverTab[73335]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:507
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:507
		// _ = "end of CoverTab[73328]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:508
		_go_fuzz_dep_.CoverTab[73336]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:508
		// _ = "end of CoverTab[73336]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:508
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:508
	// _ = "end of CoverTab[73315]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:508
	_go_fuzz_dep_.CoverTab[73316]++

										if opts.Settings != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:510
		_go_fuzz_dep_.CoverTab[73337]++
											fr := &SettingsFrame{
			FrameHeader:	FrameHeader{valid: true},
			p:		opts.Settings,
		}
		if err := fr.ForeachSetting(sc.processSetting); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:515
			_go_fuzz_dep_.CoverTab[73339]++
												sc.rejectConn(ErrCodeProtocol, "invalid settings")
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:517
			// _ = "end of CoverTab[73339]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:518
			_go_fuzz_dep_.CoverTab[73340]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:518
			// _ = "end of CoverTab[73340]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:518
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:518
		// _ = "end of CoverTab[73337]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:518
		_go_fuzz_dep_.CoverTab[73338]++
											opts.Settings = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:519
		// _ = "end of CoverTab[73338]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:520
		_go_fuzz_dep_.CoverTab[73341]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:520
		// _ = "end of CoverTab[73341]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:520
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:520
	// _ = "end of CoverTab[73316]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:520
	_go_fuzz_dep_.CoverTab[73317]++

										if hook := testHookGetServerConn; hook != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:522
		_go_fuzz_dep_.CoverTab[73342]++
											hook(sc)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:523
		// _ = "end of CoverTab[73342]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:524
		_go_fuzz_dep_.CoverTab[73343]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:524
		// _ = "end of CoverTab[73343]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:524
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:524
	// _ = "end of CoverTab[73317]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:524
	_go_fuzz_dep_.CoverTab[73318]++

										if opts.UpgradeRequest != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:526
		_go_fuzz_dep_.CoverTab[73344]++
											sc.upgradeRequest(opts.UpgradeRequest)
											opts.UpgradeRequest = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:528
		// _ = "end of CoverTab[73344]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:529
		_go_fuzz_dep_.CoverTab[73345]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:529
		// _ = "end of CoverTab[73345]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:529
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:529
	// _ = "end of CoverTab[73318]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:529
	_go_fuzz_dep_.CoverTab[73319]++

										sc.serve()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:531
	// _ = "end of CoverTab[73319]"
}

func serverConnBaseContext(c net.Conn, opts *ServeConnOpts) (ctx context.Context, cancel func()) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:534
	_go_fuzz_dep_.CoverTab[73346]++
										ctx, cancel = context.WithCancel(opts.context())
										ctx = context.WithValue(ctx, http.LocalAddrContextKey, c.LocalAddr())
										if hs := opts.baseConfig(); hs != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:537
		_go_fuzz_dep_.CoverTab[73348]++
											ctx = context.WithValue(ctx, http.ServerContextKey, hs)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:538
		// _ = "end of CoverTab[73348]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:539
		_go_fuzz_dep_.CoverTab[73349]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:539
		// _ = "end of CoverTab[73349]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:539
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:539
	// _ = "end of CoverTab[73346]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:539
	_go_fuzz_dep_.CoverTab[73347]++
										return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:540
	// _ = "end of CoverTab[73347]"
}

func (sc *serverConn) rejectConn(err ErrCode, debug string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:543
	_go_fuzz_dep_.CoverTab[73350]++
										sc.vlogf("http2: server rejecting conn: %v, %s", err, debug)

										sc.framer.WriteGoAway(0, err, []byte(debug))
										sc.bw.Flush()
										sc.conn.Close()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:548
	// _ = "end of CoverTab[73350]"
}

type serverConn struct {
	// Immutable:
	srv			*Server
	hs			*http.Server
	conn			net.Conn
	bw			*bufferedWriter	// writing to conn
	handler			http.Handler
	baseCtx			context.Context
	framer			*Framer
	doneServing		chan struct{}		// closed when serverConn.serve ends
	readFrameCh		chan readFrameResult	// written by serverConn.readFrames
	wantWriteFrameCh	chan FrameWriteRequest	// from handlers -> serve
	wroteFrameCh		chan frameWriteResult	// from writeFrameAsync -> serve, tickles more frame writes
	bodyReadCh		chan bodyReadMsg	// from handlers -> serve
	serveMsgCh		chan interface{}	// misc messages & code to send to / run on the serve loop
	flow			outflow			// conn-wide (not stream-specific) outbound flow control
	inflow			inflow			// conn-wide inbound flow control
	tlsState		*tls.ConnectionState	// shared by all handlers, like net/http
	remoteAddrStr		string
	writeSched		WriteScheduler

	// Everything following is owned by the serve loop; use serveG.check():
	serveG				goroutineLock	// used to verify funcs are on serve()
	pushEnabled			bool
	sawClientPreface		bool	// preface has already been read, used in h2c upgrade
	sawFirstSettings		bool	// got the initial SETTINGS frame after the preface
	needToSendSettingsAck		bool
	unackedSettings			int	// how many SETTINGS have we sent without ACKs?
	queuedControlFrames		int	// control frames in the writeSched queue
	clientMaxStreams		uint32	// SETTINGS_MAX_CONCURRENT_STREAMS from client (our PUSH_PROMISE limit)
	advMaxStreams			uint32	// our SETTINGS_MAX_CONCURRENT_STREAMS advertised the client
	curClientStreams		uint32	// number of open streams initiated by the client
	curPushedStreams		uint32	// number of open streams initiated by server push
	maxClientStreamID		uint32	// max ever seen from client (odd), or 0 if there have been no client requests
	maxPushPromiseID		uint32	// ID of the last push promise (even), or 0 if there have been no pushes
	streams				map[uint32]*stream
	initialStreamSendWindowSize	int32
	maxFrameSize			int32
	peerMaxHeaderListSize		uint32			// zero means unknown (default)
	canonHeader			map[string]string	// http2-lower-case -> Go-Canonical-Case
	canonHeaderKeysSize		int			// canonHeader keys size in bytes
	writingFrame			bool			// started writing a frame (on serve goroutine or separate)
	writingFrameAsync		bool			// started a frame on its own goroutine but haven't heard back on wroteFrameCh
	needsFrameFlush			bool			// last frame write wasn't a flush
	inGoAway			bool			// we've started to or sent GOAWAY
	inFrameScheduleLoop		bool			// whether we're in the scheduleFrameWrite loop
	needToSendGoAway		bool			// we need to schedule a GOAWAY frame write
	goAwayCode			ErrCode
	shutdownTimer			*time.Timer	// nil until used
	idleTimer			*time.Timer	// nil if unused

	// Owned by the writeFrameAsync goroutine:
	headerWriteBuf	bytes.Buffer
	hpackEncoder	*hpack.Encoder

	// Used by startGracefulShutdown.
	shutdownOnce	sync.Once
}

func (sc *serverConn) maxHeaderListSize() uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:610
	_go_fuzz_dep_.CoverTab[73351]++
										n := sc.hs.MaxHeaderBytes
										if n <= 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:612
		_go_fuzz_dep_.CoverTab[73353]++
											n = http.DefaultMaxHeaderBytes
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:613
		// _ = "end of CoverTab[73353]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:614
		_go_fuzz_dep_.CoverTab[73354]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:614
		// _ = "end of CoverTab[73354]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:614
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:614
	// _ = "end of CoverTab[73351]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:614
	_go_fuzz_dep_.CoverTab[73352]++
	// http2's count is in a slightly different unit and includes 32 bytes per pair.
										// So, take the net/http.Server value and pad it up a bit, assuming 10 headers.
										const perFieldOverhead = 32	// per http2 spec
										const typicalHeaders = 10	// conservative
										return uint32(n + typicalHeaders*perFieldOverhead)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:619
	// _ = "end of CoverTab[73352]"
}

func (sc *serverConn) curOpenStreams() uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:622
	_go_fuzz_dep_.CoverTab[73355]++
										sc.serveG.check()
										return sc.curClientStreams + sc.curPushedStreams
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:624
	// _ = "end of CoverTab[73355]"
}

// stream represents a stream. This is the minimal metadata needed by
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:627
// the serve goroutine. Most of the actual stream state is owned by
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:627
// the http.Handler's goroutine in the responseWriter. Because the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:627
// responseWriter's responseWriterState is recycled at the end of a
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:627
// handler, this struct intentionally has no pointer to the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:627
// *responseWriter{,State} itself, as the Handler ending nils out the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:627
// responseWriter's state field.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:634
type stream struct {
	// immutable:
	sc		*serverConn
	id		uint32
	body		*pipe		// non-nil if expecting DATA frames
	cw		closeWaiter	// closed wait stream transitions to closed state
	ctx		context.Context
	cancelCtx	func()

	// owned by serverConn's serve loop:
	bodyBytes		int64	// body bytes seen so far
	declBodyBytes		int64	// or -1 if undeclared
	flow			outflow	// limits writing from Handler to client
	inflow			inflow	// what the client is allowed to POST/etc to us
	state			streamState
	resetQueued		bool		// RST_STREAM queued for write; set by sc.resetStream
	gotTrailerHeader	bool		// HEADER frame for trailers was seen
	wroteHeaders		bool		// whether we wrote headers (not status 100)
	readDeadline		*time.Timer	// nil if unused
	writeDeadline		*time.Timer	// nil if unused
	closeErr		error		// set before cw is closed

	trailer		http.Header	// accumulated trailers
	reqTrailer	http.Header	// handler's Request.Trailer
}

func (sc *serverConn) Framer() *Framer {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:660
	_go_fuzz_dep_.CoverTab[73356]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:660
	return sc.framer
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:660
	// _ = "end of CoverTab[73356]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:660
}
func (sc *serverConn) CloseConn() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:661
	_go_fuzz_dep_.CoverTab[73357]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:661
	return sc.conn.Close()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:661
	// _ = "end of CoverTab[73357]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:661
}
func (sc *serverConn) Flush() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:662
	_go_fuzz_dep_.CoverTab[73358]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:662
	return sc.bw.Flush()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:662
	// _ = "end of CoverTab[73358]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:662
}
func (sc *serverConn) HeaderEncoder() (*hpack.Encoder, *bytes.Buffer) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:663
	_go_fuzz_dep_.CoverTab[73359]++
										return sc.hpackEncoder, &sc.headerWriteBuf
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:664
	// _ = "end of CoverTab[73359]"
}

func (sc *serverConn) state(streamID uint32) (streamState, *stream) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:667
	_go_fuzz_dep_.CoverTab[73360]++
										sc.serveG.check()

										if st, ok := sc.streams[streamID]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:670
		_go_fuzz_dep_.CoverTab[73363]++
											return st.state, st
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:671
		// _ = "end of CoverTab[73363]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:672
		_go_fuzz_dep_.CoverTab[73364]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:672
		// _ = "end of CoverTab[73364]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:672
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:672
	// _ = "end of CoverTab[73360]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:672
	_go_fuzz_dep_.CoverTab[73361]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:679
	if streamID%2 == 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:679
		_go_fuzz_dep_.CoverTab[73365]++
											if streamID <= sc.maxClientStreamID {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:680
			_go_fuzz_dep_.CoverTab[73366]++
												return stateClosed, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:681
			// _ = "end of CoverTab[73366]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:682
			_go_fuzz_dep_.CoverTab[73367]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:682
			// _ = "end of CoverTab[73367]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:682
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:682
		// _ = "end of CoverTab[73365]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:683
		_go_fuzz_dep_.CoverTab[73368]++
											if streamID <= sc.maxPushPromiseID {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:684
			_go_fuzz_dep_.CoverTab[73369]++
												return stateClosed, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:685
			// _ = "end of CoverTab[73369]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:686
			_go_fuzz_dep_.CoverTab[73370]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:686
			// _ = "end of CoverTab[73370]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:686
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:686
		// _ = "end of CoverTab[73368]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:687
	// _ = "end of CoverTab[73361]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:687
	_go_fuzz_dep_.CoverTab[73362]++
										return stateIdle, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:688
	// _ = "end of CoverTab[73362]"
}

// setConnState calls the net/http ConnState hook for this connection, if configured.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:691
// Note that the net/http package does StateNew and StateClosed for us.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:691
// There is currently no plan for StateHijacked or hijacking HTTP/2 connections.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:694
func (sc *serverConn) setConnState(state http.ConnState) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:694
	_go_fuzz_dep_.CoverTab[73371]++
										if sc.hs.ConnState != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:695
		_go_fuzz_dep_.CoverTab[73372]++
											sc.hs.ConnState(sc.conn, state)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:696
		// _ = "end of CoverTab[73372]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:697
		_go_fuzz_dep_.CoverTab[73373]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:697
		// _ = "end of CoverTab[73373]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:697
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:697
	// _ = "end of CoverTab[73371]"
}

func (sc *serverConn) vlogf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:700
	_go_fuzz_dep_.CoverTab[73374]++
										if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:701
		_go_fuzz_dep_.CoverTab[73375]++
											sc.logf(format, args...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:702
		// _ = "end of CoverTab[73375]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:703
		_go_fuzz_dep_.CoverTab[73376]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:703
		// _ = "end of CoverTab[73376]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:703
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:703
	// _ = "end of CoverTab[73374]"
}

func (sc *serverConn) logf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:706
	_go_fuzz_dep_.CoverTab[73377]++
										if lg := sc.hs.ErrorLog; lg != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:707
		_go_fuzz_dep_.CoverTab[73378]++
											lg.Printf(format, args...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:708
		// _ = "end of CoverTab[73378]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:709
		_go_fuzz_dep_.CoverTab[73379]++
											log.Printf(format, args...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:710
		// _ = "end of CoverTab[73379]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:711
	// _ = "end of CoverTab[73377]"
}

// errno returns v's underlying uintptr, else 0.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:714
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:714
// TODO: remove this helper function once http2 can use build
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:714
// tags. See comment in isClosedConnError.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:718
func errno(v error) uintptr {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:718
	_go_fuzz_dep_.CoverTab[73380]++
										if rv := reflect.ValueOf(v); rv.Kind() == reflect.Uintptr {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:719
		_go_fuzz_dep_.CoverTab[73382]++
											return uintptr(rv.Uint())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:720
		// _ = "end of CoverTab[73382]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:721
		_go_fuzz_dep_.CoverTab[73383]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:721
		// _ = "end of CoverTab[73383]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:721
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:721
	// _ = "end of CoverTab[73380]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:721
	_go_fuzz_dep_.CoverTab[73381]++
										return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:722
	// _ = "end of CoverTab[73381]"
}

// isClosedConnError reports whether err is an error from use of a closed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:725
// network connection.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:727
func isClosedConnError(err error) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:727
	_go_fuzz_dep_.CoverTab[73384]++
										if err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:728
		_go_fuzz_dep_.CoverTab[73388]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:729
		// _ = "end of CoverTab[73388]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:730
		_go_fuzz_dep_.CoverTab[73389]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:730
		// _ = "end of CoverTab[73389]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:730
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:730
	// _ = "end of CoverTab[73384]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:730
	_go_fuzz_dep_.CoverTab[73385]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:735
	str := err.Error()
	if strings.Contains(str, "use of closed network connection") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:736
		_go_fuzz_dep_.CoverTab[73390]++
											return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:737
		// _ = "end of CoverTab[73390]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:738
		_go_fuzz_dep_.CoverTab[73391]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:738
		// _ = "end of CoverTab[73391]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:738
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:738
	// _ = "end of CoverTab[73385]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:738
	_go_fuzz_dep_.CoverTab[73386]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:744
	if runtime.GOOS == "windows" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:744
		_go_fuzz_dep_.CoverTab[73392]++
											if oe, ok := err.(*net.OpError); ok && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:745
			_go_fuzz_dep_.CoverTab[73393]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:745
			return oe.Op == "read"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:745
			// _ = "end of CoverTab[73393]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:745
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:745
			_go_fuzz_dep_.CoverTab[73394]++
												if se, ok := oe.Err.(*os.SyscallError); ok && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:746
				_go_fuzz_dep_.CoverTab[73395]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:746
				return se.Syscall == "wsarecv"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:746
				// _ = "end of CoverTab[73395]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:746
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:746
				_go_fuzz_dep_.CoverTab[73396]++
													const WSAECONNABORTED = 10053
													const WSAECONNRESET = 10054
													if n := errno(se.Err); n == WSAECONNRESET || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:749
					_go_fuzz_dep_.CoverTab[73397]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:749
					return n == WSAECONNABORTED
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:749
					// _ = "end of CoverTab[73397]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:749
				}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:749
					_go_fuzz_dep_.CoverTab[73398]++
														return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:750
					// _ = "end of CoverTab[73398]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:751
					_go_fuzz_dep_.CoverTab[73399]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:751
					// _ = "end of CoverTab[73399]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:751
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:751
				// _ = "end of CoverTab[73396]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:752
				_go_fuzz_dep_.CoverTab[73400]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:752
				// _ = "end of CoverTab[73400]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:752
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:752
			// _ = "end of CoverTab[73394]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:753
			_go_fuzz_dep_.CoverTab[73401]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:753
			// _ = "end of CoverTab[73401]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:753
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:753
		// _ = "end of CoverTab[73392]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:754
		_go_fuzz_dep_.CoverTab[73402]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:754
		// _ = "end of CoverTab[73402]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:754
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:754
	// _ = "end of CoverTab[73386]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:754
	_go_fuzz_dep_.CoverTab[73387]++
										return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:755
	// _ = "end of CoverTab[73387]"
}

func (sc *serverConn) condlogf(err error, format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:758
	_go_fuzz_dep_.CoverTab[73403]++
										if err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:759
		_go_fuzz_dep_.CoverTab[73405]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:760
		// _ = "end of CoverTab[73405]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:761
		_go_fuzz_dep_.CoverTab[73406]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:761
		// _ = "end of CoverTab[73406]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:761
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:761
	// _ = "end of CoverTab[73403]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:761
	_go_fuzz_dep_.CoverTab[73404]++
										if err == io.EOF || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:762
		_go_fuzz_dep_.CoverTab[73407]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:762
		return err == io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:762
		// _ = "end of CoverTab[73407]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:762
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:762
		_go_fuzz_dep_.CoverTab[73408]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:762
		return isClosedConnError(err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:762
		// _ = "end of CoverTab[73408]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:762
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:762
		_go_fuzz_dep_.CoverTab[73409]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:762
		return err == errPrefaceTimeout
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:762
		// _ = "end of CoverTab[73409]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:762
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:762
		_go_fuzz_dep_.CoverTab[73410]++

											sc.vlogf(format, args...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:764
		// _ = "end of CoverTab[73410]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:765
		_go_fuzz_dep_.CoverTab[73411]++
											sc.logf(format, args...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:766
		// _ = "end of CoverTab[73411]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:767
	// _ = "end of CoverTab[73404]"
}

// maxCachedCanonicalHeadersKeysSize is an arbitrarily-chosen limit on the size
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:770
// of the entries in the canonHeader cache.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:770
// This should be larger than the size of unique, uncommon header keys likely to
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:770
// be sent by the peer, while not so high as to permit unreasonable memory usage
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:770
// if the peer sends an unbounded number of unique header keys.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:775
const maxCachedCanonicalHeadersKeysSize = 2048

func (sc *serverConn) canonicalHeader(v string) string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:777
	_go_fuzz_dep_.CoverTab[73412]++
										sc.serveG.check()
										buildCommonHeaderMapsOnce()
										cv, ok := commonCanonHeader[v]
										if ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:781
		_go_fuzz_dep_.CoverTab[73417]++
											return cv
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:782
		// _ = "end of CoverTab[73417]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:783
		_go_fuzz_dep_.CoverTab[73418]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:783
		// _ = "end of CoverTab[73418]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:783
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:783
	// _ = "end of CoverTab[73412]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:783
	_go_fuzz_dep_.CoverTab[73413]++
										cv, ok = sc.canonHeader[v]
										if ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:785
		_go_fuzz_dep_.CoverTab[73419]++
											return cv
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:786
		// _ = "end of CoverTab[73419]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:787
		_go_fuzz_dep_.CoverTab[73420]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:787
		// _ = "end of CoverTab[73420]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:787
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:787
	// _ = "end of CoverTab[73413]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:787
	_go_fuzz_dep_.CoverTab[73414]++
										if sc.canonHeader == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:788
		_go_fuzz_dep_.CoverTab[73421]++
											sc.canonHeader = make(map[string]string)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:789
		// _ = "end of CoverTab[73421]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:790
		_go_fuzz_dep_.CoverTab[73422]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:790
		// _ = "end of CoverTab[73422]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:790
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:790
	// _ = "end of CoverTab[73414]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:790
	_go_fuzz_dep_.CoverTab[73415]++
										cv = http.CanonicalHeaderKey(v)
										size := 100 + len(v)*2
										if sc.canonHeaderKeysSize+size <= maxCachedCanonicalHeadersKeysSize {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:793
		_go_fuzz_dep_.CoverTab[73423]++
											sc.canonHeader[v] = cv
											sc.canonHeaderKeysSize += size
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:795
		// _ = "end of CoverTab[73423]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:796
		_go_fuzz_dep_.CoverTab[73424]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:796
		// _ = "end of CoverTab[73424]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:796
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:796
	// _ = "end of CoverTab[73415]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:796
	_go_fuzz_dep_.CoverTab[73416]++
										return cv
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:797
	// _ = "end of CoverTab[73416]"
}

type readFrameResult struct {
	f	Frame	// valid until readMore is called
	err	error

	// readMore should be called once the consumer no longer needs or
	// retains f. After readMore, f is invalid and more frames can be
	// read.
	readMore	func()
}

// readFrames is the loop that reads incoming frames.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:810
// It takes care to only read one frame at a time, blocking until the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:810
// consumer is done with the frame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:810
// It's run on its own goroutine.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:814
func (sc *serverConn) readFrames() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:814
	_go_fuzz_dep_.CoverTab[73425]++
										gate := make(gate)
										gateDone := gate.Done
										for {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:817
		_go_fuzz_dep_.CoverTab[73426]++
											f, err := sc.framer.ReadFrame()
											select {
		case sc.readFrameCh <- readFrameResult{f, err, gateDone}:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:820
			_go_fuzz_dep_.CoverTab[73429]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:820
			// _ = "end of CoverTab[73429]"
		case <-sc.doneServing:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:821
			_go_fuzz_dep_.CoverTab[73430]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:822
			// _ = "end of CoverTab[73430]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:823
		// _ = "end of CoverTab[73426]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:823
		_go_fuzz_dep_.CoverTab[73427]++
											select {
		case <-gate:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:825
			_go_fuzz_dep_.CoverTab[73431]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:825
			// _ = "end of CoverTab[73431]"
		case <-sc.doneServing:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:826
			_go_fuzz_dep_.CoverTab[73432]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:827
			// _ = "end of CoverTab[73432]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:828
		// _ = "end of CoverTab[73427]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:828
		_go_fuzz_dep_.CoverTab[73428]++
											if terminalReadFrameError(err) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:829
			_go_fuzz_dep_.CoverTab[73433]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:830
			// _ = "end of CoverTab[73433]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:831
			_go_fuzz_dep_.CoverTab[73434]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:831
			// _ = "end of CoverTab[73434]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:831
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:831
		// _ = "end of CoverTab[73428]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:832
	// _ = "end of CoverTab[73425]"
}

// frameWriteResult is the message passed from writeFrameAsync to the serve goroutine.
type frameWriteResult struct {
	_	incomparable
	wr	FrameWriteRequest	// what was written (or attempted)
	err	error			// result of the writeFrame call
}

// writeFrameAsync runs in its own goroutine and writes a single frame
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:842
// and then reports when it's done.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:842
// At most one goroutine can be running writeFrameAsync at a time per
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:842
// serverConn.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:846
func (sc *serverConn) writeFrameAsync(wr FrameWriteRequest, wd *writeData) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:846
	_go_fuzz_dep_.CoverTab[73435]++
										var err error
										if wd == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:848
		_go_fuzz_dep_.CoverTab[73437]++
											err = wr.write.writeFrame(sc)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:849
		// _ = "end of CoverTab[73437]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:850
		_go_fuzz_dep_.CoverTab[73438]++
											err = sc.framer.endWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:851
		// _ = "end of CoverTab[73438]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:852
	// _ = "end of CoverTab[73435]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:852
	_go_fuzz_dep_.CoverTab[73436]++
										sc.wroteFrameCh <- frameWriteResult{wr: wr, err: err}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:853
	// _ = "end of CoverTab[73436]"
}

func (sc *serverConn) closeAllStreamsOnConnClose() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:856
	_go_fuzz_dep_.CoverTab[73439]++
										sc.serveG.check()
										for _, st := range sc.streams {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:858
		_go_fuzz_dep_.CoverTab[73440]++
											sc.closeStream(st, errClientDisconnected)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:859
		// _ = "end of CoverTab[73440]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:860
	// _ = "end of CoverTab[73439]"
}

func (sc *serverConn) stopShutdownTimer() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:863
	_go_fuzz_dep_.CoverTab[73441]++
										sc.serveG.check()
										if t := sc.shutdownTimer; t != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:865
		_go_fuzz_dep_.CoverTab[73442]++
											t.Stop()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:866
		// _ = "end of CoverTab[73442]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:867
		_go_fuzz_dep_.CoverTab[73443]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:867
		// _ = "end of CoverTab[73443]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:867
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:867
	// _ = "end of CoverTab[73441]"
}

func (sc *serverConn) notePanic() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:870
	_go_fuzz_dep_.CoverTab[73444]++

										if testHookOnPanicMu != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:872
		_go_fuzz_dep_.CoverTab[73446]++
											testHookOnPanicMu.Lock()
											defer testHookOnPanicMu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:874
		// _ = "end of CoverTab[73446]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:875
		_go_fuzz_dep_.CoverTab[73447]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:875
		// _ = "end of CoverTab[73447]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:875
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:875
	// _ = "end of CoverTab[73444]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:875
	_go_fuzz_dep_.CoverTab[73445]++
										if testHookOnPanic != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:876
		_go_fuzz_dep_.CoverTab[73448]++
											if e := recover(); e != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:877
			_go_fuzz_dep_.CoverTab[73449]++
												if testHookOnPanic(sc, e) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:878
				_go_fuzz_dep_.CoverTab[73450]++
													panic(e)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:879
				// _ = "end of CoverTab[73450]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:880
				_go_fuzz_dep_.CoverTab[73451]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:880
				// _ = "end of CoverTab[73451]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:880
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:880
			// _ = "end of CoverTab[73449]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:881
			_go_fuzz_dep_.CoverTab[73452]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:881
			// _ = "end of CoverTab[73452]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:881
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:881
		// _ = "end of CoverTab[73448]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:882
		_go_fuzz_dep_.CoverTab[73453]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:882
		// _ = "end of CoverTab[73453]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:882
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:882
	// _ = "end of CoverTab[73445]"
}

func (sc *serverConn) serve() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:885
	_go_fuzz_dep_.CoverTab[73454]++
										sc.serveG.check()
										defer sc.notePanic()
										defer sc.conn.Close()
										defer sc.closeAllStreamsOnConnClose()
										defer sc.stopShutdownTimer()
										defer close(sc.doneServing)

										if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:893
		_go_fuzz_dep_.CoverTab[73459]++
											sc.vlogf("http2: server connection from %v on %p", sc.conn.RemoteAddr(), sc.hs)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:894
		// _ = "end of CoverTab[73459]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:895
		_go_fuzz_dep_.CoverTab[73460]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:895
		// _ = "end of CoverTab[73460]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:895
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:895
	// _ = "end of CoverTab[73454]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:895
	_go_fuzz_dep_.CoverTab[73455]++

										sc.writeFrame(FrameWriteRequest{
		write: writeSettings{
			{SettingMaxFrameSize, sc.srv.maxReadFrameSize()},
			{SettingMaxConcurrentStreams, sc.advMaxStreams},
			{SettingMaxHeaderListSize, sc.maxHeaderListSize()},
			{SettingHeaderTableSize, sc.srv.maxDecoderHeaderTableSize()},
			{SettingInitialWindowSize, uint32(sc.srv.initialStreamRecvWindowSize())},
		},
	})
										sc.unackedSettings++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:910
	if diff := sc.srv.initialConnRecvWindowSize() - initialWindowSize; diff > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:910
		_go_fuzz_dep_.CoverTab[73461]++
											sc.sendWindowUpdate(nil, int(diff))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:911
		// _ = "end of CoverTab[73461]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:912
		_go_fuzz_dep_.CoverTab[73462]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:912
		// _ = "end of CoverTab[73462]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:912
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:912
	// _ = "end of CoverTab[73455]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:912
	_go_fuzz_dep_.CoverTab[73456]++

										if err := sc.readPreface(); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:914
		_go_fuzz_dep_.CoverTab[73463]++
											sc.condlogf(err, "http2: server: error reading preface from client %v: %v", sc.conn.RemoteAddr(), err)
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:916
		// _ = "end of CoverTab[73463]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:917
		_go_fuzz_dep_.CoverTab[73464]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:917
		// _ = "end of CoverTab[73464]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:917
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:917
	// _ = "end of CoverTab[73456]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:917
	_go_fuzz_dep_.CoverTab[73457]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:922
	sc.setConnState(http.StateActive)
	sc.setConnState(http.StateIdle)

	if sc.srv.IdleTimeout != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:925
		_go_fuzz_dep_.CoverTab[73465]++
											sc.idleTimer = time.AfterFunc(sc.srv.IdleTimeout, sc.onIdleTimer)
											defer sc.idleTimer.Stop()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:927
		// _ = "end of CoverTab[73465]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:928
		_go_fuzz_dep_.CoverTab[73466]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:928
		// _ = "end of CoverTab[73466]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:928
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:928
	// _ = "end of CoverTab[73457]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:928
	_go_fuzz_dep_.CoverTab[73458]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:928
	_curRoutineNum56_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:928
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum56_)

										go sc.readFrames()

										settingsTimer := time.AfterFunc(firstSettingsTimeout, sc.onSettingsTimer)
										defer settingsTimer.Stop()

										loopNum := 0
										for {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:936
		_go_fuzz_dep_.CoverTab[73467]++
											loopNum++
											select {
		case wr := <-sc.wantWriteFrameCh:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:939
			_go_fuzz_dep_.CoverTab[73470]++
												if se, ok := wr.write.(StreamError); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:940
				_go_fuzz_dep_.CoverTab[73478]++
													sc.resetStream(se)
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:942
				// _ = "end of CoverTab[73478]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:943
				_go_fuzz_dep_.CoverTab[73479]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:943
				// _ = "end of CoverTab[73479]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:943
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:943
			// _ = "end of CoverTab[73470]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:943
			_go_fuzz_dep_.CoverTab[73471]++
												sc.writeFrame(wr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:944
			// _ = "end of CoverTab[73471]"
		case res := <-sc.wroteFrameCh:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:945
			_go_fuzz_dep_.CoverTab[73472]++
												sc.wroteFrame(res)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:946
			// _ = "end of CoverTab[73472]"
		case res := <-sc.readFrameCh:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:947
			_go_fuzz_dep_.CoverTab[73473]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:950
			if sc.writingFrameAsync {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:950
				_go_fuzz_dep_.CoverTab[73480]++
													select {
				case wroteRes := <-sc.wroteFrameCh:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:952
					_go_fuzz_dep_.CoverTab[73481]++
														sc.wroteFrame(wroteRes)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:953
					// _ = "end of CoverTab[73481]"
				default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:954
					_go_fuzz_dep_.CoverTab[73482]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:954
					// _ = "end of CoverTab[73482]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:955
				// _ = "end of CoverTab[73480]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:956
				_go_fuzz_dep_.CoverTab[73483]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:956
				// _ = "end of CoverTab[73483]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:956
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:956
			// _ = "end of CoverTab[73473]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:956
			_go_fuzz_dep_.CoverTab[73474]++
												if !sc.processFrameFromReader(res) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:957
				_go_fuzz_dep_.CoverTab[73484]++
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:958
				// _ = "end of CoverTab[73484]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:959
				_go_fuzz_dep_.CoverTab[73485]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:959
				// _ = "end of CoverTab[73485]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:959
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:959
			// _ = "end of CoverTab[73474]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:959
			_go_fuzz_dep_.CoverTab[73475]++
												res.readMore()
												if settingsTimer != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:961
				_go_fuzz_dep_.CoverTab[73486]++
													settingsTimer.Stop()
													settingsTimer = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:963
				// _ = "end of CoverTab[73486]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:964
				_go_fuzz_dep_.CoverTab[73487]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:964
				// _ = "end of CoverTab[73487]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:964
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:964
			// _ = "end of CoverTab[73475]"
		case m := <-sc.bodyReadCh:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:965
			_go_fuzz_dep_.CoverTab[73476]++
												sc.noteBodyRead(m.st, m.n)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:966
			// _ = "end of CoverTab[73476]"
		case msg := <-sc.serveMsgCh:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:967
			_go_fuzz_dep_.CoverTab[73477]++
												switch v := msg.(type) {
			case func(int):
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:969
				_go_fuzz_dep_.CoverTab[73488]++
													v(loopNum)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:970
				// _ = "end of CoverTab[73488]"
			case *serverMessage:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:971
				_go_fuzz_dep_.CoverTab[73489]++
													switch v {
				case settingsTimerMsg:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:973
					_go_fuzz_dep_.CoverTab[73493]++
														sc.logf("timeout waiting for SETTINGS frames from %v", sc.conn.RemoteAddr())
														return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:975
					// _ = "end of CoverTab[73493]"
				case idleTimerMsg:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:976
					_go_fuzz_dep_.CoverTab[73494]++
														sc.vlogf("connection is idle")
														sc.goAway(ErrCodeNo)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:978
					// _ = "end of CoverTab[73494]"
				case shutdownTimerMsg:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:979
					_go_fuzz_dep_.CoverTab[73495]++
														sc.vlogf("GOAWAY close timer fired; closing conn from %v", sc.conn.RemoteAddr())
														return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:981
					// _ = "end of CoverTab[73495]"
				case gracefulShutdownMsg:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:982
					_go_fuzz_dep_.CoverTab[73496]++
														sc.startGracefulShutdownInternal()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:983
					// _ = "end of CoverTab[73496]"
				default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:984
					_go_fuzz_dep_.CoverTab[73497]++
														panic("unknown timer")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:985
					// _ = "end of CoverTab[73497]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:986
				// _ = "end of CoverTab[73489]"
			case *startPushRequest:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:987
				_go_fuzz_dep_.CoverTab[73490]++
													sc.startPush(v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:988
				// _ = "end of CoverTab[73490]"
			case func(*serverConn):
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:989
				_go_fuzz_dep_.CoverTab[73491]++
													v(sc)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:990
				// _ = "end of CoverTab[73491]"
			default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:991
				_go_fuzz_dep_.CoverTab[73492]++
													panic(fmt.Sprintf("unexpected type %T", v))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:992
				// _ = "end of CoverTab[73492]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:993
			// _ = "end of CoverTab[73477]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:994
		// _ = "end of CoverTab[73467]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:994
		_go_fuzz_dep_.CoverTab[73468]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:999
		if sc.queuedControlFrames > sc.srv.maxQueuedControlFrames() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:999
				_go_fuzz_dep_.CoverTab[73498]++
													sc.vlogf("http2: too many control frames in send queue, closing connection")
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1001
			// _ = "end of CoverTab[73498]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1002
			_go_fuzz_dep_.CoverTab[73499]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1002
			// _ = "end of CoverTab[73499]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1002
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1002
		// _ = "end of CoverTab[73468]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1002
		_go_fuzz_dep_.CoverTab[73469]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1007
		sentGoAway := sc.inGoAway && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1007
			_go_fuzz_dep_.CoverTab[73500]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1007
			return !sc.needToSendGoAway
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1007
			// _ = "end of CoverTab[73500]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1007
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1007
			_go_fuzz_dep_.CoverTab[73501]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1007
			return !sc.writingFrame
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1007
			// _ = "end of CoverTab[73501]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1007
		}()
		gracefulShutdownComplete := sc.goAwayCode == ErrCodeNo && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1008
			_go_fuzz_dep_.CoverTab[73502]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1008
			return sc.curOpenStreams() == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1008
			// _ = "end of CoverTab[73502]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1008
		}()
												if sentGoAway && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1009
			_go_fuzz_dep_.CoverTab[73503]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1009
			return sc.shutdownTimer == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1009
			// _ = "end of CoverTab[73503]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1009
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1009
			_go_fuzz_dep_.CoverTab[73504]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1009
			return (sc.goAwayCode != ErrCodeNo || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1009
				_go_fuzz_dep_.CoverTab[73505]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1009
				return gracefulShutdownComplete
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1009
				// _ = "end of CoverTab[73505]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1009
			}())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1009
			// _ = "end of CoverTab[73504]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1009
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1009
			_go_fuzz_dep_.CoverTab[73506]++
													sc.shutDownIn(goAwayTimeout)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1010
			// _ = "end of CoverTab[73506]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1011
			_go_fuzz_dep_.CoverTab[73507]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1011
			// _ = "end of CoverTab[73507]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1011
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1011
		// _ = "end of CoverTab[73469]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1012
	// _ = "end of CoverTab[73458]"
}

func (sc *serverConn) awaitGracefulShutdown(sharedCh <-chan struct{}, privateCh chan struct{}) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1015
	_go_fuzz_dep_.CoverTab[73508]++
											select {
	case <-sc.doneServing:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1017
		_go_fuzz_dep_.CoverTab[73509]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1017
		// _ = "end of CoverTab[73509]"
	case <-sharedCh:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1018
		_go_fuzz_dep_.CoverTab[73510]++
												close(privateCh)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1019
		// _ = "end of CoverTab[73510]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1020
	// _ = "end of CoverTab[73508]"
}

type serverMessage int

// Message values sent to serveMsgCh.
var (
	settingsTimerMsg	= new(serverMessage)
	idleTimerMsg		= new(serverMessage)
	shutdownTimerMsg	= new(serverMessage)
	gracefulShutdownMsg	= new(serverMessage)
)

func (sc *serverConn) onSettingsTimer() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1033
	_go_fuzz_dep_.CoverTab[73511]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1033
	sc.sendServeMsg(settingsTimerMsg)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1033
	// _ = "end of CoverTab[73511]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1033
}
func (sc *serverConn) onIdleTimer() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1034
	_go_fuzz_dep_.CoverTab[73512]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1034
	sc.sendServeMsg(idleTimerMsg)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1034
	// _ = "end of CoverTab[73512]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1034
}
func (sc *serverConn) onShutdownTimer() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1035
	_go_fuzz_dep_.CoverTab[73513]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1035
	sc.sendServeMsg(shutdownTimerMsg)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1035
	// _ = "end of CoverTab[73513]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1035
}

func (sc *serverConn) sendServeMsg(msg interface{}) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1037
	_go_fuzz_dep_.CoverTab[73514]++
											sc.serveG.checkNotOn()
											select {
	case sc.serveMsgCh <- msg:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1040
		_go_fuzz_dep_.CoverTab[73515]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1040
		// _ = "end of CoverTab[73515]"
	case <-sc.doneServing:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1041
		_go_fuzz_dep_.CoverTab[73516]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1041
		// _ = "end of CoverTab[73516]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1042
	// _ = "end of CoverTab[73514]"
}

var errPrefaceTimeout = errors.New("timeout waiting for client preface")

// readPreface reads the ClientPreface greeting from the peer or
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1047
// returns errPrefaceTimeout on timeout, or an error if the greeting
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1047
// is invalid.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1050
func (sc *serverConn) readPreface() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1050
	_go_fuzz_dep_.CoverTab[73517]++
											if sc.sawClientPreface {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1051
		_go_fuzz_dep_.CoverTab[73520]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1052
		// _ = "end of CoverTab[73520]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1053
		_go_fuzz_dep_.CoverTab[73521]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1053
		// _ = "end of CoverTab[73521]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1053
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1053
	// _ = "end of CoverTab[73517]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1053
	_go_fuzz_dep_.CoverTab[73518]++
											errc := make(chan error, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1054
	_curRoutineNum57_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1054
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum57_)
											go func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1055
		_go_fuzz_dep_.CoverTab[73522]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1055
		defer func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1055
			_go_fuzz_dep_.CoverTab[73523]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1055
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum57_)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1055
			// _ = "end of CoverTab[73523]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1055
		}()

												buf := make([]byte, len(ClientPreface))
												if _, err := io.ReadFull(sc.conn, buf); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1058
			_go_fuzz_dep_.CoverTab[73524]++
													errc <- err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1059
			// _ = "end of CoverTab[73524]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1060
			_go_fuzz_dep_.CoverTab[73525]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1060
			if !bytes.Equal(buf, clientPreface) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1060
				_go_fuzz_dep_.CoverTab[73526]++
														errc <- fmt.Errorf("bogus greeting %q", buf)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1061
				// _ = "end of CoverTab[73526]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1062
				_go_fuzz_dep_.CoverTab[73527]++
														errc <- nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1063
				// _ = "end of CoverTab[73527]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1064
			// _ = "end of CoverTab[73525]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1064
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1064
		// _ = "end of CoverTab[73522]"
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1065
	// _ = "end of CoverTab[73518]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1065
	_go_fuzz_dep_.CoverTab[73519]++
											timer := time.NewTimer(prefaceTimeout)
											defer timer.Stop()
											select {
	case <-timer.C:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1069
		_go_fuzz_dep_.CoverTab[73528]++
												return errPrefaceTimeout
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1070
		// _ = "end of CoverTab[73528]"
	case err := <-errc:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1071
		_go_fuzz_dep_.CoverTab[73529]++
												if err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1072
			_go_fuzz_dep_.CoverTab[73531]++
													if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1073
				_go_fuzz_dep_.CoverTab[73532]++
														sc.vlogf("http2: server: client %v said hello", sc.conn.RemoteAddr())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1074
				// _ = "end of CoverTab[73532]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1075
				_go_fuzz_dep_.CoverTab[73533]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1075
				// _ = "end of CoverTab[73533]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1075
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1075
			// _ = "end of CoverTab[73531]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1076
			_go_fuzz_dep_.CoverTab[73534]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1076
			// _ = "end of CoverTab[73534]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1076
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1076
		// _ = "end of CoverTab[73529]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1076
		_go_fuzz_dep_.CoverTab[73530]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1077
		// _ = "end of CoverTab[73530]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1078
	// _ = "end of CoverTab[73519]"
}

var errChanPool = sync.Pool{
	New: func() interface{} {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1082
		_go_fuzz_dep_.CoverTab[73535]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1082
		return make(chan error, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1082
		// _ = "end of CoverTab[73535]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1082
	},
}

var writeDataPool = sync.Pool{
	New: func() interface{} {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1086
		_go_fuzz_dep_.CoverTab[73536]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1086
		return new(writeData)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1086
		// _ = "end of CoverTab[73536]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1086
	},
}

// writeDataFromHandler writes DATA response frames from a handler on
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1089
// the given stream.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1091
func (sc *serverConn) writeDataFromHandler(stream *stream, data []byte, endStream bool) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1091
	_go_fuzz_dep_.CoverTab[73537]++
											ch := errChanPool.Get().(chan error)
											writeArg := writeDataPool.Get().(*writeData)
											*writeArg = writeData{stream.id, data, endStream}
											err := sc.writeFrameFromHandler(FrameWriteRequest{
		write:	writeArg,
		stream:	stream,
		done:	ch,
	})
	if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1100
		_go_fuzz_dep_.CoverTab[73541]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1101
		// _ = "end of CoverTab[73541]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1102
		_go_fuzz_dep_.CoverTab[73542]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1102
		// _ = "end of CoverTab[73542]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1102
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1102
	// _ = "end of CoverTab[73537]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1102
	_go_fuzz_dep_.CoverTab[73538]++
											var frameWriteDone bool	// the frame write is done (successfully or not)
											select {
	case err = <-ch:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1105
		_go_fuzz_dep_.CoverTab[73543]++
												frameWriteDone = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1106
		// _ = "end of CoverTab[73543]"
	case <-sc.doneServing:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1107
		_go_fuzz_dep_.CoverTab[73544]++
												return errClientDisconnected
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1108
		// _ = "end of CoverTab[73544]"
	case <-stream.cw:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1109
		_go_fuzz_dep_.CoverTab[73545]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1117
		select {
		case err = <-ch:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1118
			_go_fuzz_dep_.CoverTab[73546]++
													frameWriteDone = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1119
			// _ = "end of CoverTab[73546]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1120
			_go_fuzz_dep_.CoverTab[73547]++
													return errStreamClosed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1121
			// _ = "end of CoverTab[73547]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1122
		// _ = "end of CoverTab[73545]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1123
	// _ = "end of CoverTab[73538]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1123
	_go_fuzz_dep_.CoverTab[73539]++
											errChanPool.Put(ch)
											if frameWriteDone {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1125
		_go_fuzz_dep_.CoverTab[73548]++
												writeDataPool.Put(writeArg)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1126
		// _ = "end of CoverTab[73548]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1127
		_go_fuzz_dep_.CoverTab[73549]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1127
		// _ = "end of CoverTab[73549]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1127
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1127
	// _ = "end of CoverTab[73539]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1127
	_go_fuzz_dep_.CoverTab[73540]++
											return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1128
	// _ = "end of CoverTab[73540]"
}

// writeFrameFromHandler sends wr to sc.wantWriteFrameCh, but aborts
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1131
// if the connection has gone away.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1131
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1131
// This must not be run from the serve goroutine itself, else it might
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1131
// deadlock writing to sc.wantWriteFrameCh (which is only mildly
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1131
// buffered and is read by serve itself). If you're on the serve
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1131
// goroutine, call writeFrame instead.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1138
func (sc *serverConn) writeFrameFromHandler(wr FrameWriteRequest) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1138
	_go_fuzz_dep_.CoverTab[73550]++
											sc.serveG.checkNotOn()
											select {
	case sc.wantWriteFrameCh <- wr:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1141
		_go_fuzz_dep_.CoverTab[73551]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1142
		// _ = "end of CoverTab[73551]"
	case <-sc.doneServing:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1143
		_go_fuzz_dep_.CoverTab[73552]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1146
		return errClientDisconnected
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1146
		// _ = "end of CoverTab[73552]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1147
	// _ = "end of CoverTab[73550]"
}

// writeFrame schedules a frame to write and sends it if there's nothing
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1150
// already being written.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1150
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1150
// There is no pushback here (the serve goroutine never blocks). It's
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1150
// the http.Handlers that block, waiting for their previous frames to
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1150
// make it onto the wire
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1150
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1150
// If you're not on the serve goroutine, use writeFrameFromHandler instead.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1158
func (sc *serverConn) writeFrame(wr FrameWriteRequest) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1158
	_go_fuzz_dep_.CoverTab[73553]++
											sc.serveG.check()

											// If true, wr will not be written and wr.done will not be signaled.
											var ignoreWrite bool

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1182
	if wr.StreamID() != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1182
		_go_fuzz_dep_.CoverTab[73557]++
												_, isReset := wr.write.(StreamError)
												if state, _ := sc.state(wr.StreamID()); state == stateClosed && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1184
			_go_fuzz_dep_.CoverTab[73558]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1184
			return !isReset
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1184
			// _ = "end of CoverTab[73558]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1184
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1184
			_go_fuzz_dep_.CoverTab[73559]++
													ignoreWrite = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1185
			// _ = "end of CoverTab[73559]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1186
			_go_fuzz_dep_.CoverTab[73560]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1186
			// _ = "end of CoverTab[73560]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1186
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1186
		// _ = "end of CoverTab[73557]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1187
		_go_fuzz_dep_.CoverTab[73561]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1187
		// _ = "end of CoverTab[73561]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1187
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1187
	// _ = "end of CoverTab[73553]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1187
	_go_fuzz_dep_.CoverTab[73554]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1191
	switch wr.write.(type) {
	case *writeResHeaders:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1192
		_go_fuzz_dep_.CoverTab[73562]++
												wr.stream.wroteHeaders = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1193
		// _ = "end of CoverTab[73562]"
	case write100ContinueHeadersFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1194
		_go_fuzz_dep_.CoverTab[73563]++
												if wr.stream.wroteHeaders {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1195
			_go_fuzz_dep_.CoverTab[73564]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1198
			if wr.done != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1198
				_go_fuzz_dep_.CoverTab[73566]++
														panic("wr.done != nil for write100ContinueHeadersFrame")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1199
				// _ = "end of CoverTab[73566]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1200
				_go_fuzz_dep_.CoverTab[73567]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1200
				// _ = "end of CoverTab[73567]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1200
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1200
			// _ = "end of CoverTab[73564]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1200
			_go_fuzz_dep_.CoverTab[73565]++
													ignoreWrite = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1201
			// _ = "end of CoverTab[73565]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1202
			_go_fuzz_dep_.CoverTab[73568]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1202
			// _ = "end of CoverTab[73568]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1202
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1202
		// _ = "end of CoverTab[73563]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1203
	// _ = "end of CoverTab[73554]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1203
	_go_fuzz_dep_.CoverTab[73555]++

											if !ignoreWrite {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1205
		_go_fuzz_dep_.CoverTab[73569]++
												if wr.isControl() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1206
			_go_fuzz_dep_.CoverTab[73571]++
													sc.queuedControlFrames++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1210
			if sc.queuedControlFrames < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1210
				_go_fuzz_dep_.CoverTab[73572]++
														sc.conn.Close()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1211
				// _ = "end of CoverTab[73572]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1212
				_go_fuzz_dep_.CoverTab[73573]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1212
				// _ = "end of CoverTab[73573]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1212
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1212
			// _ = "end of CoverTab[73571]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1213
			_go_fuzz_dep_.CoverTab[73574]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1213
			// _ = "end of CoverTab[73574]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1213
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1213
		// _ = "end of CoverTab[73569]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1213
		_go_fuzz_dep_.CoverTab[73570]++
												sc.writeSched.Push(wr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1214
		// _ = "end of CoverTab[73570]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1215
		_go_fuzz_dep_.CoverTab[73575]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1215
		// _ = "end of CoverTab[73575]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1215
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1215
	// _ = "end of CoverTab[73555]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1215
	_go_fuzz_dep_.CoverTab[73556]++
											sc.scheduleFrameWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1216
	// _ = "end of CoverTab[73556]"
}

// startFrameWrite starts a goroutine to write wr (in a separate
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1219
// goroutine since that might block on the network), and updates the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1219
// serve goroutine's state about the world, updated from info in wr.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1222
func (sc *serverConn) startFrameWrite(wr FrameWriteRequest) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1222
	_go_fuzz_dep_.CoverTab[73576]++
											sc.serveG.check()
											if sc.writingFrame {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1224
		_go_fuzz_dep_.CoverTab[73580]++
												panic("internal error: can only be writing one frame at a time")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1225
		// _ = "end of CoverTab[73580]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1226
		_go_fuzz_dep_.CoverTab[73581]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1226
		// _ = "end of CoverTab[73581]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1226
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1226
	// _ = "end of CoverTab[73576]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1226
	_go_fuzz_dep_.CoverTab[73577]++

											st := wr.stream
											if st != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1229
		_go_fuzz_dep_.CoverTab[73582]++
												switch st.state {
		case stateHalfClosedLocal:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1231
			_go_fuzz_dep_.CoverTab[73583]++
													switch wr.write.(type) {
			case StreamError, handlerPanicRST, writeWindowUpdate:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1233
				_go_fuzz_dep_.CoverTab[73586]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1233
				// _ = "end of CoverTab[73586]"

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1236
			default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1236
				_go_fuzz_dep_.CoverTab[73587]++
														panic(fmt.Sprintf("internal error: attempt to send frame on a half-closed-local stream: %v", wr))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1237
				// _ = "end of CoverTab[73587]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1238
			// _ = "end of CoverTab[73583]"
		case stateClosed:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1239
			_go_fuzz_dep_.CoverTab[73584]++
													panic(fmt.Sprintf("internal error: attempt to send frame on a closed stream: %v", wr))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1240
			// _ = "end of CoverTab[73584]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1240
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1240
			_go_fuzz_dep_.CoverTab[73585]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1240
			// _ = "end of CoverTab[73585]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1241
		// _ = "end of CoverTab[73582]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1242
		_go_fuzz_dep_.CoverTab[73588]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1242
		// _ = "end of CoverTab[73588]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1242
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1242
	// _ = "end of CoverTab[73577]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1242
	_go_fuzz_dep_.CoverTab[73578]++
											if wpp, ok := wr.write.(*writePushPromise); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1243
		_go_fuzz_dep_.CoverTab[73589]++
												var err error
												wpp.promisedID, err = wpp.allocatePromisedID()
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1246
			_go_fuzz_dep_.CoverTab[73590]++
													sc.writingFrameAsync = false
													wr.replyToWriter(err)
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1249
			// _ = "end of CoverTab[73590]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1250
			_go_fuzz_dep_.CoverTab[73591]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1250
			// _ = "end of CoverTab[73591]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1250
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1250
		// _ = "end of CoverTab[73589]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1251
		_go_fuzz_dep_.CoverTab[73592]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1251
		// _ = "end of CoverTab[73592]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1251
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1251
	// _ = "end of CoverTab[73578]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1251
	_go_fuzz_dep_.CoverTab[73579]++

											sc.writingFrame = true
											sc.needsFrameFlush = true
											if wr.write.staysWithinBuffer(sc.bw.Available()) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1255
		_go_fuzz_dep_.CoverTab[73593]++
												sc.writingFrameAsync = false
												err := wr.write.writeFrame(sc)
												sc.wroteFrame(frameWriteResult{wr: wr, err: err})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1258
		// _ = "end of CoverTab[73593]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1259
		_go_fuzz_dep_.CoverTab[73594]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1259
		if wd, ok := wr.write.(*writeData); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1259
			_go_fuzz_dep_.CoverTab[73595]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1263
			sc.framer.startWriteDataPadded(wd.streamID, wd.endStream, wd.p, nil)
													sc.writingFrameAsync = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1264
			_curRoutineNum58_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1264
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum58_)
													go sc.writeFrameAsync(wr, wd)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1265
			// _ = "end of CoverTab[73595]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1266
			_go_fuzz_dep_.CoverTab[73596]++
													sc.writingFrameAsync = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1267
			_curRoutineNum59_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1267
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum59_)
													go sc.writeFrameAsync(wr, nil)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1268
			// _ = "end of CoverTab[73596]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1269
		// _ = "end of CoverTab[73594]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1269
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1269
	// _ = "end of CoverTab[73579]"
}

// errHandlerPanicked is the error given to any callers blocked in a read from
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1272
// Request.Body when the main goroutine panics. Since most handlers read in the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1272
// main ServeHTTP goroutine, this will show up rarely.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1275
var errHandlerPanicked = errors.New("http2: handler panicked")

// wroteFrame is called on the serve goroutine with the result of
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1277
// whatever happened on writeFrameAsync.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1279
func (sc *serverConn) wroteFrame(res frameWriteResult) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1279
	_go_fuzz_dep_.CoverTab[73597]++
											sc.serveG.check()
											if !sc.writingFrame {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1281
		_go_fuzz_dep_.CoverTab[73600]++
												panic("internal error: expected to be already writing a frame")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1282
		// _ = "end of CoverTab[73600]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1283
		_go_fuzz_dep_.CoverTab[73601]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1283
		// _ = "end of CoverTab[73601]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1283
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1283
	// _ = "end of CoverTab[73597]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1283
	_go_fuzz_dep_.CoverTab[73598]++
											sc.writingFrame = false
											sc.writingFrameAsync = false

											wr := res.wr

											if writeEndsStream(wr.write) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1289
		_go_fuzz_dep_.CoverTab[73602]++
												st := wr.stream
												if st == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1291
			_go_fuzz_dep_.CoverTab[73604]++
													panic("internal error: expecting non-nil stream")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1292
			// _ = "end of CoverTab[73604]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1293
			_go_fuzz_dep_.CoverTab[73605]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1293
			// _ = "end of CoverTab[73605]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1293
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1293
		// _ = "end of CoverTab[73602]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1293
		_go_fuzz_dep_.CoverTab[73603]++
												switch st.state {
		case stateOpen:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1295
			_go_fuzz_dep_.CoverTab[73606]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1306
			st.state = stateHalfClosedLocal

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1311
			sc.resetStream(streamError(st.id, ErrCodeNo))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1311
			// _ = "end of CoverTab[73606]"
		case stateHalfClosedRemote:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1312
			_go_fuzz_dep_.CoverTab[73607]++
													sc.closeStream(st, errHandlerComplete)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1313
			// _ = "end of CoverTab[73607]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1313
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1313
			_go_fuzz_dep_.CoverTab[73608]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1313
			// _ = "end of CoverTab[73608]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1314
		// _ = "end of CoverTab[73603]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1315
		_go_fuzz_dep_.CoverTab[73609]++
												switch v := wr.write.(type) {
		case StreamError:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1317
			_go_fuzz_dep_.CoverTab[73610]++

													if st, ok := sc.streams[v.StreamID]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1319
				_go_fuzz_dep_.CoverTab[73612]++
														sc.closeStream(st, v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1320
				// _ = "end of CoverTab[73612]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1321
				_go_fuzz_dep_.CoverTab[73613]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1321
				// _ = "end of CoverTab[73613]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1321
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1321
			// _ = "end of CoverTab[73610]"
		case handlerPanicRST:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1322
			_go_fuzz_dep_.CoverTab[73611]++
													sc.closeStream(wr.stream, errHandlerPanicked)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1323
			// _ = "end of CoverTab[73611]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1324
		// _ = "end of CoverTab[73609]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1325
	// _ = "end of CoverTab[73598]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1325
	_go_fuzz_dep_.CoverTab[73599]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1328
	wr.replyToWriter(res.err)

											sc.scheduleFrameWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1330
	// _ = "end of CoverTab[73599]"
}

// scheduleFrameWrite tickles the frame writing scheduler.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1333
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1333
// If a frame is already being written, nothing happens. This will be called again
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1333
// when the frame is done being written.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1333
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1333
// If a frame isn't being written and we need to send one, the best frame
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1333
// to send is selected by writeSched.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1333
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1333
// If a frame isn't being written and there's nothing else to send, we
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1333
// flush the write buffer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1343
func (sc *serverConn) scheduleFrameWrite() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1343
	_go_fuzz_dep_.CoverTab[73614]++
											sc.serveG.check()
											if sc.writingFrame || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1345
		_go_fuzz_dep_.CoverTab[73617]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1345
		return sc.inFrameScheduleLoop
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1345
		// _ = "end of CoverTab[73617]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1345
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1345
		_go_fuzz_dep_.CoverTab[73618]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1346
		// _ = "end of CoverTab[73618]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1347
		_go_fuzz_dep_.CoverTab[73619]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1347
		// _ = "end of CoverTab[73619]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1347
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1347
	// _ = "end of CoverTab[73614]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1347
	_go_fuzz_dep_.CoverTab[73615]++
											sc.inFrameScheduleLoop = true
											for !sc.writingFrameAsync {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1349
		_go_fuzz_dep_.CoverTab[73620]++
												if sc.needToSendGoAway {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1350
			_go_fuzz_dep_.CoverTab[73625]++
													sc.needToSendGoAway = false
													sc.startFrameWrite(FrameWriteRequest{
				write: &writeGoAway{
					maxStreamID:	sc.maxClientStreamID,
					code:		sc.goAwayCode,
				},
			})
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1358
			// _ = "end of CoverTab[73625]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1359
			_go_fuzz_dep_.CoverTab[73626]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1359
			// _ = "end of CoverTab[73626]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1359
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1359
		// _ = "end of CoverTab[73620]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1359
		_go_fuzz_dep_.CoverTab[73621]++
												if sc.needToSendSettingsAck {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1360
			_go_fuzz_dep_.CoverTab[73627]++
													sc.needToSendSettingsAck = false
													sc.startFrameWrite(FrameWriteRequest{write: writeSettingsAck{}})
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1363
			// _ = "end of CoverTab[73627]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1364
			_go_fuzz_dep_.CoverTab[73628]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1364
			// _ = "end of CoverTab[73628]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1364
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1364
		// _ = "end of CoverTab[73621]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1364
		_go_fuzz_dep_.CoverTab[73622]++
												if !sc.inGoAway || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1365
			_go_fuzz_dep_.CoverTab[73629]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1365
			return sc.goAwayCode == ErrCodeNo
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1365
			// _ = "end of CoverTab[73629]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1365
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1365
			_go_fuzz_dep_.CoverTab[73630]++
													if wr, ok := sc.writeSched.Pop(); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1366
				_go_fuzz_dep_.CoverTab[73631]++
														if wr.isControl() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1367
					_go_fuzz_dep_.CoverTab[73633]++
															sc.queuedControlFrames--
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1368
					// _ = "end of CoverTab[73633]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1369
					_go_fuzz_dep_.CoverTab[73634]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1369
					// _ = "end of CoverTab[73634]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1369
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1369
				// _ = "end of CoverTab[73631]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1369
				_go_fuzz_dep_.CoverTab[73632]++
														sc.startFrameWrite(wr)
														continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1371
				// _ = "end of CoverTab[73632]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1372
				_go_fuzz_dep_.CoverTab[73635]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1372
				// _ = "end of CoverTab[73635]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1372
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1372
			// _ = "end of CoverTab[73630]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1373
			_go_fuzz_dep_.CoverTab[73636]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1373
			// _ = "end of CoverTab[73636]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1373
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1373
		// _ = "end of CoverTab[73622]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1373
		_go_fuzz_dep_.CoverTab[73623]++
												if sc.needsFrameFlush {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1374
			_go_fuzz_dep_.CoverTab[73637]++
													sc.startFrameWrite(FrameWriteRequest{write: flushFrameWriter{}})
													sc.needsFrameFlush = false
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1377
			// _ = "end of CoverTab[73637]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1378
			_go_fuzz_dep_.CoverTab[73638]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1378
			// _ = "end of CoverTab[73638]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1378
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1378
		// _ = "end of CoverTab[73623]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1378
		_go_fuzz_dep_.CoverTab[73624]++
												break
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1379
		// _ = "end of CoverTab[73624]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1380
	// _ = "end of CoverTab[73615]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1380
	_go_fuzz_dep_.CoverTab[73616]++
											sc.inFrameScheduleLoop = false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1381
	// _ = "end of CoverTab[73616]"
}

// startGracefulShutdown gracefully shuts down a connection. This
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1384
// sends GOAWAY with ErrCodeNo to tell the client we're gracefully
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1384
// shutting down. The connection isn't closed until all current
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1384
// streams are done.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1384
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1384
// startGracefulShutdown returns immediately; it does not wait until
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1384
// the connection has shut down.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1391
func (sc *serverConn) startGracefulShutdown() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1391
	_go_fuzz_dep_.CoverTab[73639]++
											sc.serveG.checkNotOn()
											sc.shutdownOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1393
		_go_fuzz_dep_.CoverTab[73640]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1393
		sc.sendServeMsg(gracefulShutdownMsg)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1393
		// _ = "end of CoverTab[73640]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1393
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1393
	// _ = "end of CoverTab[73639]"
}

// After sending GOAWAY with an error code (non-graceful shutdown), the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1396
// connection will close after goAwayTimeout.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1396
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1396
// If we close the connection immediately after sending GOAWAY, there may
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1396
// be unsent data in our kernel receive buffer, which will cause the kernel
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1396
// to send a TCP RST on close() instead of a FIN. This RST will abort the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1396
// connection immediately, whether or not the client had received the GOAWAY.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1396
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1396
// Ideally we should delay for at least 1 RTT + epsilon so the client has
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1396
// a chance to read the GOAWAY and stop sending messages. Measuring RTT
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1396
// is hard, so we approximate with 1 second. See golang.org/issue/18701.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1396
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1396
// This is a var so it can be shorter in tests, where all requests uses the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1396
// loopback interface making the expected RTT very small.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1396
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1396
// TODO: configurable?
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1412
var goAwayTimeout = 1 * time.Second

func (sc *serverConn) startGracefulShutdownInternal() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1414
	_go_fuzz_dep_.CoverTab[73641]++
											sc.goAway(ErrCodeNo)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1415
	// _ = "end of CoverTab[73641]"
}

func (sc *serverConn) goAway(code ErrCode) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1418
	_go_fuzz_dep_.CoverTab[73642]++
											sc.serveG.check()
											if sc.inGoAway {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1420
		_go_fuzz_dep_.CoverTab[73644]++
												if sc.goAwayCode == ErrCodeNo {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1421
			_go_fuzz_dep_.CoverTab[73646]++
													sc.goAwayCode = code
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1422
			// _ = "end of CoverTab[73646]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1423
			_go_fuzz_dep_.CoverTab[73647]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1423
			// _ = "end of CoverTab[73647]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1423
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1423
		// _ = "end of CoverTab[73644]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1423
		_go_fuzz_dep_.CoverTab[73645]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1424
		// _ = "end of CoverTab[73645]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1425
		_go_fuzz_dep_.CoverTab[73648]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1425
		// _ = "end of CoverTab[73648]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1425
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1425
	// _ = "end of CoverTab[73642]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1425
	_go_fuzz_dep_.CoverTab[73643]++
											sc.inGoAway = true
											sc.needToSendGoAway = true
											sc.goAwayCode = code
											sc.scheduleFrameWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1429
	// _ = "end of CoverTab[73643]"
}

func (sc *serverConn) shutDownIn(d time.Duration) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1432
	_go_fuzz_dep_.CoverTab[73649]++
											sc.serveG.check()
											sc.shutdownTimer = time.AfterFunc(d, sc.onShutdownTimer)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1434
	// _ = "end of CoverTab[73649]"
}

func (sc *serverConn) resetStream(se StreamError) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1437
	_go_fuzz_dep_.CoverTab[73650]++
											sc.serveG.check()
											sc.writeFrame(FrameWriteRequest{write: se})
											if st, ok := sc.streams[se.StreamID]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1440
		_go_fuzz_dep_.CoverTab[73651]++
												st.resetQueued = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1441
		// _ = "end of CoverTab[73651]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1442
		_go_fuzz_dep_.CoverTab[73652]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1442
		// _ = "end of CoverTab[73652]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1442
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1442
	// _ = "end of CoverTab[73650]"
}

// processFrameFromReader processes the serve loop's read from readFrameCh from the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1445
// frame-reading goroutine.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1445
// processFrameFromReader returns whether the connection should be kept open.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1448
func (sc *serverConn) processFrameFromReader(res readFrameResult) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1448
	_go_fuzz_dep_.CoverTab[73653]++
											sc.serveG.check()
											err := res.err
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1451
		_go_fuzz_dep_.CoverTab[73655]++
												if err == ErrFrameTooLarge {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1452
			_go_fuzz_dep_.CoverTab[73657]++
													sc.goAway(ErrCodeFrameSize)
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1454
			// _ = "end of CoverTab[73657]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1455
			_go_fuzz_dep_.CoverTab[73658]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1455
			// _ = "end of CoverTab[73658]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1455
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1455
		// _ = "end of CoverTab[73655]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1455
		_go_fuzz_dep_.CoverTab[73656]++
												clientGone := err == io.EOF || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1456
			_go_fuzz_dep_.CoverTab[73659]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1456
			return err == io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1456
			// _ = "end of CoverTab[73659]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1456
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1456
			_go_fuzz_dep_.CoverTab[73660]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1456
			return isClosedConnError(err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1456
			// _ = "end of CoverTab[73660]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1456
		}()
												if clientGone {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1457
			_go_fuzz_dep_.CoverTab[73661]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1466
			return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1466
			// _ = "end of CoverTab[73661]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1467
			_go_fuzz_dep_.CoverTab[73662]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1467
			// _ = "end of CoverTab[73662]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1467
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1467
		// _ = "end of CoverTab[73656]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1468
		_go_fuzz_dep_.CoverTab[73663]++
												f := res.f
												if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1470
			_go_fuzz_dep_.CoverTab[73665]++
													sc.vlogf("http2: server read frame %v", summarizeFrame(f))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1471
			// _ = "end of CoverTab[73665]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1472
			_go_fuzz_dep_.CoverTab[73666]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1472
			// _ = "end of CoverTab[73666]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1472
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1472
		// _ = "end of CoverTab[73663]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1472
		_go_fuzz_dep_.CoverTab[73664]++
												err = sc.processFrame(f)
												if err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1474
			_go_fuzz_dep_.CoverTab[73667]++
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1475
			// _ = "end of CoverTab[73667]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1476
			_go_fuzz_dep_.CoverTab[73668]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1476
			// _ = "end of CoverTab[73668]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1476
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1476
		// _ = "end of CoverTab[73664]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1477
	// _ = "end of CoverTab[73653]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1477
	_go_fuzz_dep_.CoverTab[73654]++

											switch ev := err.(type) {
	case StreamError:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1480
		_go_fuzz_dep_.CoverTab[73669]++
												sc.resetStream(ev)
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1482
		// _ = "end of CoverTab[73669]"
	case goAwayFlowError:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1483
		_go_fuzz_dep_.CoverTab[73670]++
												sc.goAway(ErrCodeFlowControl)
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1485
		// _ = "end of CoverTab[73670]"
	case ConnectionError:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1486
		_go_fuzz_dep_.CoverTab[73671]++
												sc.logf("http2: server connection error from %v: %v", sc.conn.RemoteAddr(), ev)
												sc.goAway(ErrCode(ev))
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1489
		// _ = "end of CoverTab[73671]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1490
		_go_fuzz_dep_.CoverTab[73672]++
												if res.err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1491
			_go_fuzz_dep_.CoverTab[73674]++
													sc.vlogf("http2: server closing client connection; error reading frame from client %s: %v", sc.conn.RemoteAddr(), err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1492
			// _ = "end of CoverTab[73674]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1493
			_go_fuzz_dep_.CoverTab[73675]++
													sc.logf("http2: server closing client connection: %v", err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1494
			// _ = "end of CoverTab[73675]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1495
		// _ = "end of CoverTab[73672]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1495
		_go_fuzz_dep_.CoverTab[73673]++
												return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1496
		// _ = "end of CoverTab[73673]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1497
	// _ = "end of CoverTab[73654]"
}

func (sc *serverConn) processFrame(f Frame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1500
	_go_fuzz_dep_.CoverTab[73676]++
											sc.serveG.check()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1504
	if !sc.sawFirstSettings {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1504
		_go_fuzz_dep_.CoverTab[73679]++
												if _, ok := f.(*SettingsFrame); !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1505
			_go_fuzz_dep_.CoverTab[73681]++
													return sc.countError("first_settings", ConnectionError(ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1506
			// _ = "end of CoverTab[73681]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1507
			_go_fuzz_dep_.CoverTab[73682]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1507
			// _ = "end of CoverTab[73682]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1507
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1507
		// _ = "end of CoverTab[73679]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1507
		_go_fuzz_dep_.CoverTab[73680]++
												sc.sawFirstSettings = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1508
		// _ = "end of CoverTab[73680]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1509
		_go_fuzz_dep_.CoverTab[73683]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1509
		// _ = "end of CoverTab[73683]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1509
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1509
	// _ = "end of CoverTab[73676]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1509
	_go_fuzz_dep_.CoverTab[73677]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1515
	if sc.inGoAway && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1515
		_go_fuzz_dep_.CoverTab[73684]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1515
		return (sc.goAwayCode != ErrCodeNo || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1515
			_go_fuzz_dep_.CoverTab[73685]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1515
			return f.Header().StreamID > sc.maxClientStreamID
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1515
			// _ = "end of CoverTab[73685]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1515
		}())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1515
		// _ = "end of CoverTab[73684]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1515
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1515
		_go_fuzz_dep_.CoverTab[73686]++

												if f, ok := f.(*DataFrame); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1517
			_go_fuzz_dep_.CoverTab[73688]++
													if !sc.inflow.take(f.Length) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1518
				_go_fuzz_dep_.CoverTab[73690]++
														return sc.countError("data_flow", streamError(f.Header().StreamID, ErrCodeFlowControl))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1519
				// _ = "end of CoverTab[73690]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1520
				_go_fuzz_dep_.CoverTab[73691]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1520
				// _ = "end of CoverTab[73691]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1520
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1520
			// _ = "end of CoverTab[73688]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1520
			_go_fuzz_dep_.CoverTab[73689]++
													sc.sendWindowUpdate(nil, int(f.Length))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1521
			// _ = "end of CoverTab[73689]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1522
			_go_fuzz_dep_.CoverTab[73692]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1522
			// _ = "end of CoverTab[73692]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1522
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1522
		// _ = "end of CoverTab[73686]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1522
		_go_fuzz_dep_.CoverTab[73687]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1523
		// _ = "end of CoverTab[73687]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1524
		_go_fuzz_dep_.CoverTab[73693]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1524
		// _ = "end of CoverTab[73693]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1524
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1524
	// _ = "end of CoverTab[73677]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1524
	_go_fuzz_dep_.CoverTab[73678]++

											switch f := f.(type) {
	case *SettingsFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1527
		_go_fuzz_dep_.CoverTab[73694]++
												return sc.processSettings(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1528
		// _ = "end of CoverTab[73694]"
	case *MetaHeadersFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1529
		_go_fuzz_dep_.CoverTab[73695]++
												return sc.processHeaders(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1530
		// _ = "end of CoverTab[73695]"
	case *WindowUpdateFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1531
		_go_fuzz_dep_.CoverTab[73696]++
												return sc.processWindowUpdate(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1532
		// _ = "end of CoverTab[73696]"
	case *PingFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1533
		_go_fuzz_dep_.CoverTab[73697]++
												return sc.processPing(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1534
		// _ = "end of CoverTab[73697]"
	case *DataFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1535
		_go_fuzz_dep_.CoverTab[73698]++
												return sc.processData(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1536
		// _ = "end of CoverTab[73698]"
	case *RSTStreamFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1537
		_go_fuzz_dep_.CoverTab[73699]++
												return sc.processResetStream(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1538
		// _ = "end of CoverTab[73699]"
	case *PriorityFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1539
		_go_fuzz_dep_.CoverTab[73700]++
												return sc.processPriority(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1540
		// _ = "end of CoverTab[73700]"
	case *GoAwayFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1541
		_go_fuzz_dep_.CoverTab[73701]++
												return sc.processGoAway(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1542
		// _ = "end of CoverTab[73701]"
	case *PushPromiseFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1543
		_go_fuzz_dep_.CoverTab[73702]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1546
		return sc.countError("push_promise", ConnectionError(ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1546
		// _ = "end of CoverTab[73702]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1547
		_go_fuzz_dep_.CoverTab[73703]++
												sc.vlogf("http2: server ignoring frame: %v", f.Header())
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1549
		// _ = "end of CoverTab[73703]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1550
	// _ = "end of CoverTab[73678]"
}

func (sc *serverConn) processPing(f *PingFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1553
	_go_fuzz_dep_.CoverTab[73704]++
											sc.serveG.check()
											if f.IsAck() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1555
		_go_fuzz_dep_.CoverTab[73707]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1558
		return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1558
		// _ = "end of CoverTab[73707]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1559
		_go_fuzz_dep_.CoverTab[73708]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1559
		// _ = "end of CoverTab[73708]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1559
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1559
	// _ = "end of CoverTab[73704]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1559
	_go_fuzz_dep_.CoverTab[73705]++
											if f.StreamID != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1560
		_go_fuzz_dep_.CoverTab[73709]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1566
		return sc.countError("ping_on_stream", ConnectionError(ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1566
		// _ = "end of CoverTab[73709]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1567
		_go_fuzz_dep_.CoverTab[73710]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1567
		// _ = "end of CoverTab[73710]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1567
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1567
	// _ = "end of CoverTab[73705]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1567
	_go_fuzz_dep_.CoverTab[73706]++
											sc.writeFrame(FrameWriteRequest{write: writePingAck{f}})
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1569
	// _ = "end of CoverTab[73706]"
}

func (sc *serverConn) processWindowUpdate(f *WindowUpdateFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1572
	_go_fuzz_dep_.CoverTab[73711]++
											sc.serveG.check()
											switch {
	case f.StreamID != 0:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1575
		_go_fuzz_dep_.CoverTab[73713]++
												state, st := sc.state(f.StreamID)
												if state == stateIdle {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1577
			_go_fuzz_dep_.CoverTab[73717]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1582
			return sc.countError("stream_idle", ConnectionError(ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1582
			// _ = "end of CoverTab[73717]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1583
			_go_fuzz_dep_.CoverTab[73718]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1583
			// _ = "end of CoverTab[73718]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1583
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1583
		// _ = "end of CoverTab[73713]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1583
		_go_fuzz_dep_.CoverTab[73714]++
												if st == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1584
			_go_fuzz_dep_.CoverTab[73719]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1590
			return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1590
			// _ = "end of CoverTab[73719]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1591
			_go_fuzz_dep_.CoverTab[73720]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1591
			// _ = "end of CoverTab[73720]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1591
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1591
		// _ = "end of CoverTab[73714]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1591
		_go_fuzz_dep_.CoverTab[73715]++
												if !st.flow.add(int32(f.Increment)) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1592
			_go_fuzz_dep_.CoverTab[73721]++
													return sc.countError("bad_flow", streamError(f.StreamID, ErrCodeFlowControl))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1593
			// _ = "end of CoverTab[73721]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1594
			_go_fuzz_dep_.CoverTab[73722]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1594
			// _ = "end of CoverTab[73722]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1594
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1594
		// _ = "end of CoverTab[73715]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1595
		_go_fuzz_dep_.CoverTab[73716]++
												if !sc.flow.add(int32(f.Increment)) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1596
			_go_fuzz_dep_.CoverTab[73723]++
													return goAwayFlowError{}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1597
			// _ = "end of CoverTab[73723]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1598
			_go_fuzz_dep_.CoverTab[73724]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1598
			// _ = "end of CoverTab[73724]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1598
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1598
		// _ = "end of CoverTab[73716]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1599
	// _ = "end of CoverTab[73711]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1599
	_go_fuzz_dep_.CoverTab[73712]++
											sc.scheduleFrameWrite()
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1601
	// _ = "end of CoverTab[73712]"
}

func (sc *serverConn) processResetStream(f *RSTStreamFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1604
	_go_fuzz_dep_.CoverTab[73725]++
											sc.serveG.check()

											state, st := sc.state(f.StreamID)
											if state == stateIdle {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1608
		_go_fuzz_dep_.CoverTab[73728]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1614
		return sc.countError("reset_idle_stream", ConnectionError(ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1614
		// _ = "end of CoverTab[73728]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1615
		_go_fuzz_dep_.CoverTab[73729]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1615
		// _ = "end of CoverTab[73729]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1615
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1615
	// _ = "end of CoverTab[73725]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1615
	_go_fuzz_dep_.CoverTab[73726]++
											if st != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1616
		_go_fuzz_dep_.CoverTab[73730]++
												st.cancelCtx()
												sc.closeStream(st, streamError(f.StreamID, f.ErrCode))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1618
		// _ = "end of CoverTab[73730]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1619
		_go_fuzz_dep_.CoverTab[73731]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1619
		// _ = "end of CoverTab[73731]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1619
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1619
	// _ = "end of CoverTab[73726]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1619
	_go_fuzz_dep_.CoverTab[73727]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1620
	// _ = "end of CoverTab[73727]"
}

func (sc *serverConn) closeStream(st *stream, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1623
	_go_fuzz_dep_.CoverTab[73732]++
											sc.serveG.check()
											if st.state == stateIdle || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1625
		_go_fuzz_dep_.CoverTab[73740]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1625
		return st.state == stateClosed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1625
		// _ = "end of CoverTab[73740]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1625
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1625
		_go_fuzz_dep_.CoverTab[73741]++
												panic(fmt.Sprintf("invariant; can't close stream in state %v", st.state))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1626
		// _ = "end of CoverTab[73741]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1627
		_go_fuzz_dep_.CoverTab[73742]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1627
		// _ = "end of CoverTab[73742]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1627
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1627
	// _ = "end of CoverTab[73732]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1627
	_go_fuzz_dep_.CoverTab[73733]++
											st.state = stateClosed
											if st.readDeadline != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1629
		_go_fuzz_dep_.CoverTab[73743]++
												st.readDeadline.Stop()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1630
		// _ = "end of CoverTab[73743]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1631
		_go_fuzz_dep_.CoverTab[73744]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1631
		// _ = "end of CoverTab[73744]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1631
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1631
	// _ = "end of CoverTab[73733]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1631
	_go_fuzz_dep_.CoverTab[73734]++
											if st.writeDeadline != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1632
		_go_fuzz_dep_.CoverTab[73745]++
												st.writeDeadline.Stop()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1633
		// _ = "end of CoverTab[73745]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1634
		_go_fuzz_dep_.CoverTab[73746]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1634
		// _ = "end of CoverTab[73746]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1634
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1634
	// _ = "end of CoverTab[73734]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1634
	_go_fuzz_dep_.CoverTab[73735]++
											if st.isPushed() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1635
		_go_fuzz_dep_.CoverTab[73747]++
												sc.curPushedStreams--
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1636
		// _ = "end of CoverTab[73747]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1637
		_go_fuzz_dep_.CoverTab[73748]++
												sc.curClientStreams--
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1638
		// _ = "end of CoverTab[73748]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1639
	// _ = "end of CoverTab[73735]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1639
	_go_fuzz_dep_.CoverTab[73736]++
											delete(sc.streams, st.id)
											if len(sc.streams) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1641
		_go_fuzz_dep_.CoverTab[73749]++
												sc.setConnState(http.StateIdle)
												if sc.srv.IdleTimeout != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1643
			_go_fuzz_dep_.CoverTab[73751]++
													sc.idleTimer.Reset(sc.srv.IdleTimeout)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1644
			// _ = "end of CoverTab[73751]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1645
			_go_fuzz_dep_.CoverTab[73752]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1645
			// _ = "end of CoverTab[73752]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1645
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1645
		// _ = "end of CoverTab[73749]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1645
		_go_fuzz_dep_.CoverTab[73750]++
												if h1ServerKeepAlivesDisabled(sc.hs) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1646
			_go_fuzz_dep_.CoverTab[73753]++
													sc.startGracefulShutdownInternal()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1647
			// _ = "end of CoverTab[73753]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1648
			_go_fuzz_dep_.CoverTab[73754]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1648
			// _ = "end of CoverTab[73754]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1648
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1648
		// _ = "end of CoverTab[73750]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1649
		_go_fuzz_dep_.CoverTab[73755]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1649
		// _ = "end of CoverTab[73755]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1649
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1649
	// _ = "end of CoverTab[73736]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1649
	_go_fuzz_dep_.CoverTab[73737]++
											if p := st.body; p != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1650
		_go_fuzz_dep_.CoverTab[73756]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1653
		sc.sendWindowUpdate(nil, p.Len())

												p.CloseWithError(err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1655
		// _ = "end of CoverTab[73756]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1656
		_go_fuzz_dep_.CoverTab[73757]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1656
		// _ = "end of CoverTab[73757]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1656
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1656
	// _ = "end of CoverTab[73737]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1656
	_go_fuzz_dep_.CoverTab[73738]++
											if e, ok := err.(StreamError); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1657
		_go_fuzz_dep_.CoverTab[73758]++
												if e.Cause != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1658
			_go_fuzz_dep_.CoverTab[73759]++
													err = e.Cause
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1659
			// _ = "end of CoverTab[73759]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1660
			_go_fuzz_dep_.CoverTab[73760]++
													err = errStreamClosed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1661
			// _ = "end of CoverTab[73760]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1662
		// _ = "end of CoverTab[73758]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1663
		_go_fuzz_dep_.CoverTab[73761]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1663
		// _ = "end of CoverTab[73761]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1663
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1663
	// _ = "end of CoverTab[73738]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1663
	_go_fuzz_dep_.CoverTab[73739]++
											st.closeErr = err
											st.cw.Close()
											sc.writeSched.CloseStream(st.id)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1666
	// _ = "end of CoverTab[73739]"
}

func (sc *serverConn) processSettings(f *SettingsFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1669
	_go_fuzz_dep_.CoverTab[73762]++
											sc.serveG.check()
											if f.IsAck() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1671
		_go_fuzz_dep_.CoverTab[73766]++
												sc.unackedSettings--
												if sc.unackedSettings < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1673
			_go_fuzz_dep_.CoverTab[73768]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1677
			return sc.countError("ack_mystery", ConnectionError(ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1677
			// _ = "end of CoverTab[73768]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1678
			_go_fuzz_dep_.CoverTab[73769]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1678
			// _ = "end of CoverTab[73769]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1678
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1678
		// _ = "end of CoverTab[73766]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1678
		_go_fuzz_dep_.CoverTab[73767]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1679
		// _ = "end of CoverTab[73767]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1680
		_go_fuzz_dep_.CoverTab[73770]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1680
		// _ = "end of CoverTab[73770]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1680
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1680
	// _ = "end of CoverTab[73762]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1680
	_go_fuzz_dep_.CoverTab[73763]++
											if f.NumSettings() > 100 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1681
		_go_fuzz_dep_.CoverTab[73771]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1681
		return f.HasDuplicates()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1681
		// _ = "end of CoverTab[73771]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1681
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1681
		_go_fuzz_dep_.CoverTab[73772]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1685
		return sc.countError("settings_big_or_dups", ConnectionError(ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1685
		// _ = "end of CoverTab[73772]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1686
		_go_fuzz_dep_.CoverTab[73773]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1686
		// _ = "end of CoverTab[73773]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1686
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1686
	// _ = "end of CoverTab[73763]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1686
	_go_fuzz_dep_.CoverTab[73764]++
											if err := f.ForeachSetting(sc.processSetting); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1687
		_go_fuzz_dep_.CoverTab[73774]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1688
		// _ = "end of CoverTab[73774]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1689
		_go_fuzz_dep_.CoverTab[73775]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1689
		// _ = "end of CoverTab[73775]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1689
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1689
	// _ = "end of CoverTab[73764]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1689
	_go_fuzz_dep_.CoverTab[73765]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1692
	sc.needToSendSettingsAck = true
											sc.scheduleFrameWrite()
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1694
	// _ = "end of CoverTab[73765]"
}

func (sc *serverConn) processSetting(s Setting) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1697
	_go_fuzz_dep_.CoverTab[73776]++
											sc.serveG.check()
											if err := s.Valid(); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1699
		_go_fuzz_dep_.CoverTab[73780]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1700
		// _ = "end of CoverTab[73780]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1701
		_go_fuzz_dep_.CoverTab[73781]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1701
		// _ = "end of CoverTab[73781]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1701
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1701
	// _ = "end of CoverTab[73776]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1701
	_go_fuzz_dep_.CoverTab[73777]++
											if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1702
		_go_fuzz_dep_.CoverTab[73782]++
												sc.vlogf("http2: server processing setting %v", s)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1703
		// _ = "end of CoverTab[73782]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1704
		_go_fuzz_dep_.CoverTab[73783]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1704
		// _ = "end of CoverTab[73783]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1704
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1704
	// _ = "end of CoverTab[73777]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1704
	_go_fuzz_dep_.CoverTab[73778]++
											switch s.ID {
	case SettingHeaderTableSize:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1706
		_go_fuzz_dep_.CoverTab[73784]++
												sc.hpackEncoder.SetMaxDynamicTableSize(s.Val)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1707
		// _ = "end of CoverTab[73784]"
	case SettingEnablePush:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1708
		_go_fuzz_dep_.CoverTab[73785]++
												sc.pushEnabled = s.Val != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1709
		// _ = "end of CoverTab[73785]"
	case SettingMaxConcurrentStreams:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1710
		_go_fuzz_dep_.CoverTab[73786]++
												sc.clientMaxStreams = s.Val
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1711
		// _ = "end of CoverTab[73786]"
	case SettingInitialWindowSize:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1712
		_go_fuzz_dep_.CoverTab[73787]++
												return sc.processSettingInitialWindowSize(s.Val)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1713
		// _ = "end of CoverTab[73787]"
	case SettingMaxFrameSize:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1714
		_go_fuzz_dep_.CoverTab[73788]++
												sc.maxFrameSize = int32(s.Val)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1715
		// _ = "end of CoverTab[73788]"
	case SettingMaxHeaderListSize:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1716
		_go_fuzz_dep_.CoverTab[73789]++
												sc.peerMaxHeaderListSize = s.Val
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1717
		// _ = "end of CoverTab[73789]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1718
		_go_fuzz_dep_.CoverTab[73790]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1722
		if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1722
			_go_fuzz_dep_.CoverTab[73791]++
													sc.vlogf("http2: server ignoring unknown setting %v", s)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1723
			// _ = "end of CoverTab[73791]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1724
			_go_fuzz_dep_.CoverTab[73792]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1724
			// _ = "end of CoverTab[73792]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1724
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1724
		// _ = "end of CoverTab[73790]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1725
	// _ = "end of CoverTab[73778]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1725
	_go_fuzz_dep_.CoverTab[73779]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1726
	// _ = "end of CoverTab[73779]"
}

func (sc *serverConn) processSettingInitialWindowSize(val uint32) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1729
	_go_fuzz_dep_.CoverTab[73793]++
											sc.serveG.check()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1740
	old := sc.initialStreamSendWindowSize
	sc.initialStreamSendWindowSize = int32(val)
	growth := int32(val) - old
	for _, st := range sc.streams {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1743
		_go_fuzz_dep_.CoverTab[73795]++
												if !st.flow.add(growth) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1744
			_go_fuzz_dep_.CoverTab[73796]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1751
			return sc.countError("setting_win_size", ConnectionError(ErrCodeFlowControl))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1751
			// _ = "end of CoverTab[73796]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1752
			_go_fuzz_dep_.CoverTab[73797]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1752
			// _ = "end of CoverTab[73797]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1752
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1752
		// _ = "end of CoverTab[73795]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1753
	// _ = "end of CoverTab[73793]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1753
	_go_fuzz_dep_.CoverTab[73794]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1754
	// _ = "end of CoverTab[73794]"
}

func (sc *serverConn) processData(f *DataFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1757
	_go_fuzz_dep_.CoverTab[73798]++
											sc.serveG.check()
											id := f.Header().StreamID

											data := f.Data()
											state, st := sc.state(id)
											if id == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1763
		_go_fuzz_dep_.CoverTab[73805]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1763
		return state == stateIdle
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1763
		// _ = "end of CoverTab[73805]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1763
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1763
		_go_fuzz_dep_.CoverTab[73806]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1774
		return sc.countError("data_on_idle", ConnectionError(ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1774
		// _ = "end of CoverTab[73806]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1775
		_go_fuzz_dep_.CoverTab[73807]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1775
		// _ = "end of CoverTab[73807]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1775
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1775
	// _ = "end of CoverTab[73798]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1775
	_go_fuzz_dep_.CoverTab[73799]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1780
	if st == nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1780
		_go_fuzz_dep_.CoverTab[73808]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1780
		return state != stateOpen
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1780
		// _ = "end of CoverTab[73808]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1780
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1780
		_go_fuzz_dep_.CoverTab[73809]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1780
		return st.gotTrailerHeader
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1780
		// _ = "end of CoverTab[73809]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1780
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1780
		_go_fuzz_dep_.CoverTab[73810]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1780
		return st.resetQueued
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1780
		// _ = "end of CoverTab[73810]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1780
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1780
		_go_fuzz_dep_.CoverTab[73811]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1790
		if !sc.inflow.take(f.Length) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1790
			_go_fuzz_dep_.CoverTab[73814]++
													return sc.countError("data_flow", streamError(id, ErrCodeFlowControl))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1791
			// _ = "end of CoverTab[73814]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1792
			_go_fuzz_dep_.CoverTab[73815]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1792
			// _ = "end of CoverTab[73815]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1792
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1792
		// _ = "end of CoverTab[73811]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1792
		_go_fuzz_dep_.CoverTab[73812]++
												sc.sendWindowUpdate(nil, int(f.Length))

												if st != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1795
			_go_fuzz_dep_.CoverTab[73816]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1795
			return st.resetQueued
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1795
			// _ = "end of CoverTab[73816]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1795
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1795
			_go_fuzz_dep_.CoverTab[73817]++

													return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1797
			// _ = "end of CoverTab[73817]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1798
			_go_fuzz_dep_.CoverTab[73818]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1798
			// _ = "end of CoverTab[73818]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1798
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1798
		// _ = "end of CoverTab[73812]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1798
		_go_fuzz_dep_.CoverTab[73813]++
												return sc.countError("closed", streamError(id, ErrCodeStreamClosed))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1799
		// _ = "end of CoverTab[73813]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1800
		_go_fuzz_dep_.CoverTab[73819]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1800
		// _ = "end of CoverTab[73819]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1800
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1800
	// _ = "end of CoverTab[73799]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1800
	_go_fuzz_dep_.CoverTab[73800]++
											if st.body == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1801
		_go_fuzz_dep_.CoverTab[73820]++
												panic("internal error: should have a body in this state")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1802
		// _ = "end of CoverTab[73820]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1803
		_go_fuzz_dep_.CoverTab[73821]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1803
		// _ = "end of CoverTab[73821]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1803
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1803
	// _ = "end of CoverTab[73800]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1803
	_go_fuzz_dep_.CoverTab[73801]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1806
	if st.declBodyBytes != -1 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1806
		_go_fuzz_dep_.CoverTab[73822]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1806
		return st.bodyBytes+int64(len(data)) > st.declBodyBytes
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1806
		// _ = "end of CoverTab[73822]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1806
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1806
		_go_fuzz_dep_.CoverTab[73823]++
												if !sc.inflow.take(f.Length) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1807
			_go_fuzz_dep_.CoverTab[73825]++
													return sc.countError("data_flow", streamError(id, ErrCodeFlowControl))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1808
			// _ = "end of CoverTab[73825]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1809
			_go_fuzz_dep_.CoverTab[73826]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1809
			// _ = "end of CoverTab[73826]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1809
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1809
		// _ = "end of CoverTab[73823]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1809
		_go_fuzz_dep_.CoverTab[73824]++
												sc.sendWindowUpdate(nil, int(f.Length))

												st.body.CloseWithError(fmt.Errorf("sender tried to send more than declared Content-Length of %d bytes", st.declBodyBytes))

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1816
		return sc.countError("send_too_much", streamError(id, ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1816
		// _ = "end of CoverTab[73824]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1817
		_go_fuzz_dep_.CoverTab[73827]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1817
		// _ = "end of CoverTab[73827]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1817
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1817
	// _ = "end of CoverTab[73801]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1817
	_go_fuzz_dep_.CoverTab[73802]++
											if f.Length > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1818
		_go_fuzz_dep_.CoverTab[73828]++

												if !takeInflows(&sc.inflow, &st.inflow, f.Length) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1820
			_go_fuzz_dep_.CoverTab[73831]++
													return sc.countError("flow_on_data_length", streamError(id, ErrCodeFlowControl))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1821
			// _ = "end of CoverTab[73831]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1822
			_go_fuzz_dep_.CoverTab[73832]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1822
			// _ = "end of CoverTab[73832]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1822
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1822
		// _ = "end of CoverTab[73828]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1822
		_go_fuzz_dep_.CoverTab[73829]++

												if len(data) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1824
			_go_fuzz_dep_.CoverTab[73833]++
													st.bodyBytes += int64(len(data))
													wrote, err := st.body.Write(data)
													if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1827
				_go_fuzz_dep_.CoverTab[73835]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1831
				sc.sendWindowUpdate(nil, int(f.Length)-wrote)
														return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1832
				// _ = "end of CoverTab[73835]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1833
				_go_fuzz_dep_.CoverTab[73836]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1833
				// _ = "end of CoverTab[73836]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1833
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1833
			// _ = "end of CoverTab[73833]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1833
			_go_fuzz_dep_.CoverTab[73834]++
													if wrote != len(data) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1834
				_go_fuzz_dep_.CoverTab[73837]++
														panic("internal error: bad Writer")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1835
				// _ = "end of CoverTab[73837]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1836
				_go_fuzz_dep_.CoverTab[73838]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1836
				// _ = "end of CoverTab[73838]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1836
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1836
			// _ = "end of CoverTab[73834]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1837
			_go_fuzz_dep_.CoverTab[73839]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1837
			// _ = "end of CoverTab[73839]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1837
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1837
		// _ = "end of CoverTab[73829]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1837
		_go_fuzz_dep_.CoverTab[73830]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1844
		pad := int32(f.Length) - int32(len(data))
												sc.sendWindowUpdate32(nil, pad)
												sc.sendWindowUpdate32(st, pad)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1846
		// _ = "end of CoverTab[73830]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1847
		_go_fuzz_dep_.CoverTab[73840]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1847
		// _ = "end of CoverTab[73840]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1847
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1847
	// _ = "end of CoverTab[73802]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1847
	_go_fuzz_dep_.CoverTab[73803]++
											if f.StreamEnded() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1848
		_go_fuzz_dep_.CoverTab[73841]++
												st.endStream()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1849
		// _ = "end of CoverTab[73841]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1850
		_go_fuzz_dep_.CoverTab[73842]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1850
		// _ = "end of CoverTab[73842]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1850
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1850
	// _ = "end of CoverTab[73803]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1850
	_go_fuzz_dep_.CoverTab[73804]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1851
	// _ = "end of CoverTab[73804]"
}

func (sc *serverConn) processGoAway(f *GoAwayFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1854
	_go_fuzz_dep_.CoverTab[73843]++
											sc.serveG.check()
											if f.ErrCode != ErrCodeNo {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1856
		_go_fuzz_dep_.CoverTab[73845]++
												sc.logf("http2: received GOAWAY %+v, starting graceful shutdown", f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1857
		// _ = "end of CoverTab[73845]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1858
		_go_fuzz_dep_.CoverTab[73846]++
												sc.vlogf("http2: received GOAWAY %+v, starting graceful shutdown", f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1859
		// _ = "end of CoverTab[73846]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1860
	// _ = "end of CoverTab[73843]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1860
	_go_fuzz_dep_.CoverTab[73844]++
											sc.startGracefulShutdownInternal()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1864
	sc.pushEnabled = false
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1865
	// _ = "end of CoverTab[73844]"
}

// isPushed reports whether the stream is server-initiated.
func (st *stream) isPushed() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1869
	_go_fuzz_dep_.CoverTab[73847]++
											return st.id%2 == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1870
	// _ = "end of CoverTab[73847]"
}

// endStream closes a Request.Body's pipe. It is called when a DATA
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1873
// frame says a request body is over (or after trailers).
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1875
func (st *stream) endStream() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1875
	_go_fuzz_dep_.CoverTab[73848]++
											sc := st.sc
											sc.serveG.check()

											if st.declBodyBytes != -1 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1879
		_go_fuzz_dep_.CoverTab[73850]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1879
		return st.declBodyBytes != st.bodyBytes
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1879
		// _ = "end of CoverTab[73850]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1879
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1879
		_go_fuzz_dep_.CoverTab[73851]++
												st.body.CloseWithError(fmt.Errorf("request declared a Content-Length of %d but only wrote %d bytes",
			st.declBodyBytes, st.bodyBytes))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1881
		// _ = "end of CoverTab[73851]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1882
		_go_fuzz_dep_.CoverTab[73852]++
												st.body.closeWithErrorAndCode(io.EOF, st.copyTrailersToHandlerRequest)
												st.body.CloseWithError(io.EOF)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1884
		// _ = "end of CoverTab[73852]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1885
	// _ = "end of CoverTab[73848]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1885
	_go_fuzz_dep_.CoverTab[73849]++
											st.state = stateHalfClosedRemote
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1886
	// _ = "end of CoverTab[73849]"
}

// copyTrailersToHandlerRequest is run in the Handler's goroutine in
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1889
// its Request.Body.Read just before it gets io.EOF.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1891
func (st *stream) copyTrailersToHandlerRequest() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1891
	_go_fuzz_dep_.CoverTab[73853]++
											for k, vv := range st.trailer {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1892
		_go_fuzz_dep_.CoverTab[73854]++
												if _, ok := st.reqTrailer[k]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1893
			_go_fuzz_dep_.CoverTab[73855]++

													st.reqTrailer[k] = vv
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1895
			// _ = "end of CoverTab[73855]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1896
			_go_fuzz_dep_.CoverTab[73856]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1896
			// _ = "end of CoverTab[73856]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1896
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1896
		// _ = "end of CoverTab[73854]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1897
	// _ = "end of CoverTab[73853]"
}

// onReadTimeout is run on its own goroutine (from time.AfterFunc)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1900
// when the stream's ReadTimeout has fired.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1902
func (st *stream) onReadTimeout() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1902
	_go_fuzz_dep_.CoverTab[73857]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1905
	st.body.CloseWithError(fmt.Errorf("%w", os.ErrDeadlineExceeded))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1905
	// _ = "end of CoverTab[73857]"
}

// onWriteTimeout is run on its own goroutine (from time.AfterFunc)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1908
// when the stream's WriteTimeout has fired.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1910
func (st *stream) onWriteTimeout() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1910
	_go_fuzz_dep_.CoverTab[73858]++
											st.sc.writeFrameFromHandler(FrameWriteRequest{write: StreamError{
		StreamID:	st.id,
		Code:		ErrCodeInternal,
		Cause:		os.ErrDeadlineExceeded,
	}})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1915
	// _ = "end of CoverTab[73858]"
}

func (sc *serverConn) processHeaders(f *MetaHeadersFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1918
	_go_fuzz_dep_.CoverTab[73859]++
											sc.serveG.check()
											id := f.StreamID

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1926
	if id%2 != 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1926
		_go_fuzz_dep_.CoverTab[73871]++
												return sc.countError("headers_even", ConnectionError(ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1927
		// _ = "end of CoverTab[73871]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1928
		_go_fuzz_dep_.CoverTab[73872]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1928
		// _ = "end of CoverTab[73872]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1928
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1928
	// _ = "end of CoverTab[73859]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1928
	_go_fuzz_dep_.CoverTab[73860]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1933
	if st := sc.streams[f.StreamID]; st != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1933
		_go_fuzz_dep_.CoverTab[73873]++
												if st.resetQueued {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1934
			_go_fuzz_dep_.CoverTab[73876]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1937
			return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1937
			// _ = "end of CoverTab[73876]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1938
			_go_fuzz_dep_.CoverTab[73877]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1938
			// _ = "end of CoverTab[73877]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1938
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1938
		// _ = "end of CoverTab[73873]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1938
		_go_fuzz_dep_.CoverTab[73874]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1943
		if st.state == stateHalfClosedRemote {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1943
			_go_fuzz_dep_.CoverTab[73878]++
													return sc.countError("headers_half_closed", streamError(id, ErrCodeStreamClosed))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1944
			// _ = "end of CoverTab[73878]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1945
			_go_fuzz_dep_.CoverTab[73879]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1945
			// _ = "end of CoverTab[73879]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1945
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1945
		// _ = "end of CoverTab[73874]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1945
		_go_fuzz_dep_.CoverTab[73875]++
												return st.processTrailerHeaders(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1946
		// _ = "end of CoverTab[73875]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1947
		_go_fuzz_dep_.CoverTab[73880]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1947
		// _ = "end of CoverTab[73880]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1947
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1947
	// _ = "end of CoverTab[73860]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1947
	_go_fuzz_dep_.CoverTab[73861]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1954
	if id <= sc.maxClientStreamID {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1954
		_go_fuzz_dep_.CoverTab[73881]++
												return sc.countError("stream_went_down", ConnectionError(ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1955
		// _ = "end of CoverTab[73881]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1956
		_go_fuzz_dep_.CoverTab[73882]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1956
		// _ = "end of CoverTab[73882]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1956
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1956
	// _ = "end of CoverTab[73861]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1956
	_go_fuzz_dep_.CoverTab[73862]++
											sc.maxClientStreamID = id

											if sc.idleTimer != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1959
		_go_fuzz_dep_.CoverTab[73883]++
												sc.idleTimer.Stop()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1960
		// _ = "end of CoverTab[73883]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1961
		_go_fuzz_dep_.CoverTab[73884]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1961
		// _ = "end of CoverTab[73884]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1961
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1961
	// _ = "end of CoverTab[73862]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1961
	_go_fuzz_dep_.CoverTab[73863]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1969
	if sc.curClientStreams+1 > sc.advMaxStreams {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1969
		_go_fuzz_dep_.CoverTab[73885]++
												if sc.unackedSettings == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1970
			_go_fuzz_dep_.CoverTab[73887]++

													return sc.countError("over_max_streams", streamError(id, ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1972
			// _ = "end of CoverTab[73887]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1973
			_go_fuzz_dep_.CoverTab[73888]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1973
			// _ = "end of CoverTab[73888]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1973
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1973
		// _ = "end of CoverTab[73885]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1973
		_go_fuzz_dep_.CoverTab[73886]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1979
		return sc.countError("over_max_streams_race", streamError(id, ErrCodeRefusedStream))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1979
		// _ = "end of CoverTab[73886]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1980
		_go_fuzz_dep_.CoverTab[73889]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1980
		// _ = "end of CoverTab[73889]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1980
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1980
	// _ = "end of CoverTab[73863]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1980
	_go_fuzz_dep_.CoverTab[73864]++

											initialState := stateOpen
											if f.StreamEnded() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1983
		_go_fuzz_dep_.CoverTab[73890]++
												initialState = stateHalfClosedRemote
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1984
		// _ = "end of CoverTab[73890]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1985
		_go_fuzz_dep_.CoverTab[73891]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1985
		// _ = "end of CoverTab[73891]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1985
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1985
	// _ = "end of CoverTab[73864]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1985
	_go_fuzz_dep_.CoverTab[73865]++
											st := sc.newStream(id, 0, initialState)

											if f.HasPriority() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1988
		_go_fuzz_dep_.CoverTab[73892]++
												if err := sc.checkPriority(f.StreamID, f.Priority); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1989
			_go_fuzz_dep_.CoverTab[73894]++
													return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1990
			// _ = "end of CoverTab[73894]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1991
			_go_fuzz_dep_.CoverTab[73895]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1991
			// _ = "end of CoverTab[73895]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1991
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1991
		// _ = "end of CoverTab[73892]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1991
		_go_fuzz_dep_.CoverTab[73893]++
												sc.writeSched.AdjustStream(st.id, f.Priority)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1992
		// _ = "end of CoverTab[73893]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1993
		_go_fuzz_dep_.CoverTab[73896]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1993
		// _ = "end of CoverTab[73896]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1993
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1993
	// _ = "end of CoverTab[73865]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1993
	_go_fuzz_dep_.CoverTab[73866]++

											rw, req, err := sc.newWriterAndRequest(st, f)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1996
		_go_fuzz_dep_.CoverTab[73897]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1997
		// _ = "end of CoverTab[73897]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1998
		_go_fuzz_dep_.CoverTab[73898]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1998
		// _ = "end of CoverTab[73898]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1998
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1998
	// _ = "end of CoverTab[73866]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:1998
	_go_fuzz_dep_.CoverTab[73867]++
											st.reqTrailer = req.Trailer
											if st.reqTrailer != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2000
		_go_fuzz_dep_.CoverTab[73899]++
												st.trailer = make(http.Header)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2001
		// _ = "end of CoverTab[73899]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2002
		_go_fuzz_dep_.CoverTab[73900]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2002
		// _ = "end of CoverTab[73900]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2002
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2002
	// _ = "end of CoverTab[73867]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2002
	_go_fuzz_dep_.CoverTab[73868]++
											st.body = req.Body.(*requestBody).pipe
											st.declBodyBytes = req.ContentLength

											handler := sc.handler.ServeHTTP
											if f.Truncated {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2007
		_go_fuzz_dep_.CoverTab[73901]++

												handler = handleHeaderListTooLong
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2009
		// _ = "end of CoverTab[73901]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2010
		_go_fuzz_dep_.CoverTab[73902]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2010
		if err := checkValidHTTP2RequestHeaders(req.Header); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2010
			_go_fuzz_dep_.CoverTab[73903]++
													handler = new400Handler(err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2011
			// _ = "end of CoverTab[73903]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2012
			_go_fuzz_dep_.CoverTab[73904]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2012
			// _ = "end of CoverTab[73904]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2012
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2012
		// _ = "end of CoverTab[73902]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2012
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2012
	// _ = "end of CoverTab[73868]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2012
	_go_fuzz_dep_.CoverTab[73869]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2021
	if sc.hs.ReadTimeout != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2021
		_go_fuzz_dep_.CoverTab[73905]++
												sc.conn.SetReadDeadline(time.Time{})
												if st.body != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2023
			_go_fuzz_dep_.CoverTab[73906]++
													st.readDeadline = time.AfterFunc(sc.hs.ReadTimeout, st.onReadTimeout)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2024
			// _ = "end of CoverTab[73906]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2025
			_go_fuzz_dep_.CoverTab[73907]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2025
			// _ = "end of CoverTab[73907]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2025
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2025
		// _ = "end of CoverTab[73905]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2026
		_go_fuzz_dep_.CoverTab[73908]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2026
		// _ = "end of CoverTab[73908]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2026
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2026
	// _ = "end of CoverTab[73869]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2026
	_go_fuzz_dep_.CoverTab[73870]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2026
	_curRoutineNum60_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2026
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum60_)

											go sc.runHandler(rw, req, handler)
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2029
	// _ = "end of CoverTab[73870]"
}

func (sc *serverConn) upgradeRequest(req *http.Request) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2032
	_go_fuzz_dep_.CoverTab[73909]++
											sc.serveG.check()
											id := uint32(1)
											sc.maxClientStreamID = id
											st := sc.newStream(id, 0, stateHalfClosedRemote)
											st.reqTrailer = req.Trailer
											if st.reqTrailer != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2038
		_go_fuzz_dep_.CoverTab[73912]++
												st.trailer = make(http.Header)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2039
		// _ = "end of CoverTab[73912]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2040
		_go_fuzz_dep_.CoverTab[73913]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2040
		// _ = "end of CoverTab[73913]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2040
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2040
	// _ = "end of CoverTab[73909]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2040
	_go_fuzz_dep_.CoverTab[73910]++
											rw := sc.newResponseWriter(st, req)

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2045
	if sc.hs.ReadTimeout != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2045
		_go_fuzz_dep_.CoverTab[73914]++
												sc.conn.SetReadDeadline(time.Time{})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2046
		// _ = "end of CoverTab[73914]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2047
		_go_fuzz_dep_.CoverTab[73915]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2047
		// _ = "end of CoverTab[73915]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2047
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2047
	// _ = "end of CoverTab[73910]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2047
	_go_fuzz_dep_.CoverTab[73911]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2047
	_curRoutineNum61_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2047
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum61_)

											go sc.runHandler(rw, req, sc.handler.ServeHTTP)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2049
	// _ = "end of CoverTab[73911]"
}

func (st *stream) processTrailerHeaders(f *MetaHeadersFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2052
	_go_fuzz_dep_.CoverTab[73916]++
											sc := st.sc
											sc.serveG.check()
											if st.gotTrailerHeader {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2055
		_go_fuzz_dep_.CoverTab[73921]++
												return sc.countError("dup_trailers", ConnectionError(ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2056
		// _ = "end of CoverTab[73921]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2057
		_go_fuzz_dep_.CoverTab[73922]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2057
		// _ = "end of CoverTab[73922]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2057
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2057
	// _ = "end of CoverTab[73916]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2057
	_go_fuzz_dep_.CoverTab[73917]++
											st.gotTrailerHeader = true
											if !f.StreamEnded() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2059
		_go_fuzz_dep_.CoverTab[73923]++
												return sc.countError("trailers_not_ended", streamError(st.id, ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2060
		// _ = "end of CoverTab[73923]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2061
		_go_fuzz_dep_.CoverTab[73924]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2061
		// _ = "end of CoverTab[73924]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2061
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2061
	// _ = "end of CoverTab[73917]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2061
	_go_fuzz_dep_.CoverTab[73918]++

											if len(f.PseudoFields()) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2063
		_go_fuzz_dep_.CoverTab[73925]++
												return sc.countError("trailers_pseudo", streamError(st.id, ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2064
		// _ = "end of CoverTab[73925]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2065
		_go_fuzz_dep_.CoverTab[73926]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2065
		// _ = "end of CoverTab[73926]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2065
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2065
	// _ = "end of CoverTab[73918]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2065
	_go_fuzz_dep_.CoverTab[73919]++
											if st.trailer != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2066
		_go_fuzz_dep_.CoverTab[73927]++
												for _, hf := range f.RegularFields() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2067
			_go_fuzz_dep_.CoverTab[73928]++
													key := sc.canonicalHeader(hf.Name)
													if !httpguts.ValidTrailerHeader(key) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2069
				_go_fuzz_dep_.CoverTab[73930]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2073
				return sc.countError("trailers_bogus", streamError(st.id, ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2073
				// _ = "end of CoverTab[73930]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2074
				_go_fuzz_dep_.CoverTab[73931]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2074
				// _ = "end of CoverTab[73931]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2074
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2074
			// _ = "end of CoverTab[73928]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2074
			_go_fuzz_dep_.CoverTab[73929]++
													st.trailer[key] = append(st.trailer[key], hf.Value)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2075
			// _ = "end of CoverTab[73929]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2076
		// _ = "end of CoverTab[73927]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2077
		_go_fuzz_dep_.CoverTab[73932]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2077
		// _ = "end of CoverTab[73932]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2077
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2077
	// _ = "end of CoverTab[73919]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2077
	_go_fuzz_dep_.CoverTab[73920]++
											st.endStream()
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2079
	// _ = "end of CoverTab[73920]"
}

func (sc *serverConn) checkPriority(streamID uint32, p PriorityParam) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2082
	_go_fuzz_dep_.CoverTab[73933]++
											if streamID == p.StreamDep {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2083
		_go_fuzz_dep_.CoverTab[73935]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2088
		return sc.countError("priority", streamError(streamID, ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2088
		// _ = "end of CoverTab[73935]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2089
		_go_fuzz_dep_.CoverTab[73936]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2089
		// _ = "end of CoverTab[73936]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2089
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2089
	// _ = "end of CoverTab[73933]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2089
	_go_fuzz_dep_.CoverTab[73934]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2090
	// _ = "end of CoverTab[73934]"
}

func (sc *serverConn) processPriority(f *PriorityFrame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2093
	_go_fuzz_dep_.CoverTab[73937]++
											if err := sc.checkPriority(f.StreamID, f.PriorityParam); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2094
		_go_fuzz_dep_.CoverTab[73939]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2095
		// _ = "end of CoverTab[73939]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2096
		_go_fuzz_dep_.CoverTab[73940]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2096
		// _ = "end of CoverTab[73940]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2096
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2096
	// _ = "end of CoverTab[73937]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2096
	_go_fuzz_dep_.CoverTab[73938]++
											sc.writeSched.AdjustStream(f.StreamID, f.PriorityParam)
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2098
	// _ = "end of CoverTab[73938]"
}

func (sc *serverConn) newStream(id, pusherID uint32, state streamState) *stream {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2101
	_go_fuzz_dep_.CoverTab[73941]++
											sc.serveG.check()
											if id == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2103
		_go_fuzz_dep_.CoverTab[73946]++
												panic("internal error: cannot create stream with id 0")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2104
		// _ = "end of CoverTab[73946]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2105
		_go_fuzz_dep_.CoverTab[73947]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2105
		// _ = "end of CoverTab[73947]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2105
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2105
	// _ = "end of CoverTab[73941]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2105
	_go_fuzz_dep_.CoverTab[73942]++

											ctx, cancelCtx := context.WithCancel(sc.baseCtx)
											st := &stream{
		sc:		sc,
		id:		id,
		state:		state,
		ctx:		ctx,
		cancelCtx:	cancelCtx,
	}
	st.cw.Init()
	st.flow.conn = &sc.flow
	st.flow.add(sc.initialStreamSendWindowSize)
	st.inflow.init(sc.srv.initialStreamRecvWindowSize())
	if sc.hs.WriteTimeout != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2119
		_go_fuzz_dep_.CoverTab[73948]++
												st.writeDeadline = time.AfterFunc(sc.hs.WriteTimeout, st.onWriteTimeout)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2120
		// _ = "end of CoverTab[73948]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2121
		_go_fuzz_dep_.CoverTab[73949]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2121
		// _ = "end of CoverTab[73949]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2121
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2121
	// _ = "end of CoverTab[73942]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2121
	_go_fuzz_dep_.CoverTab[73943]++

											sc.streams[id] = st
											sc.writeSched.OpenStream(st.id, OpenStreamOptions{PusherID: pusherID})
											if st.isPushed() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2125
		_go_fuzz_dep_.CoverTab[73950]++
												sc.curPushedStreams++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2126
		// _ = "end of CoverTab[73950]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2127
		_go_fuzz_dep_.CoverTab[73951]++
												sc.curClientStreams++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2128
		// _ = "end of CoverTab[73951]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2129
	// _ = "end of CoverTab[73943]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2129
	_go_fuzz_dep_.CoverTab[73944]++
											if sc.curOpenStreams() == 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2130
		_go_fuzz_dep_.CoverTab[73952]++
												sc.setConnState(http.StateActive)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2131
		// _ = "end of CoverTab[73952]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2132
		_go_fuzz_dep_.CoverTab[73953]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2132
		// _ = "end of CoverTab[73953]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2132
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2132
	// _ = "end of CoverTab[73944]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2132
	_go_fuzz_dep_.CoverTab[73945]++

											return st
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2134
	// _ = "end of CoverTab[73945]"
}

func (sc *serverConn) newWriterAndRequest(st *stream, f *MetaHeadersFrame) (*responseWriter, *http.Request, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2137
	_go_fuzz_dep_.CoverTab[73954]++
											sc.serveG.check()

											rp := requestParam{
		method:		f.PseudoValue("method"),
		scheme:		f.PseudoValue("scheme"),
		authority:	f.PseudoValue("authority"),
		path:		f.PseudoValue("path"),
	}

	isConnect := rp.method == "CONNECT"
	if isConnect {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2148
		_go_fuzz_dep_.CoverTab[73960]++
												if rp.path != "" || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2149
			_go_fuzz_dep_.CoverTab[73961]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2149
			return rp.scheme != ""
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2149
			// _ = "end of CoverTab[73961]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2149
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2149
			_go_fuzz_dep_.CoverTab[73962]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2149
			return rp.authority == ""
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2149
			// _ = "end of CoverTab[73962]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2149
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2149
			_go_fuzz_dep_.CoverTab[73963]++
													return nil, nil, sc.countError("bad_connect", streamError(f.StreamID, ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2150
			// _ = "end of CoverTab[73963]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2151
			_go_fuzz_dep_.CoverTab[73964]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2151
			// _ = "end of CoverTab[73964]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2151
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2151
		// _ = "end of CoverTab[73960]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2152
		_go_fuzz_dep_.CoverTab[73965]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2152
		if rp.method == "" || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2152
			_go_fuzz_dep_.CoverTab[73966]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2152
			return rp.path == ""
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2152
			// _ = "end of CoverTab[73966]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2152
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2152
			_go_fuzz_dep_.CoverTab[73967]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2152
			return (rp.scheme != "https" && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2152
				_go_fuzz_dep_.CoverTab[73968]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2152
				return rp.scheme != "http"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2152
				// _ = "end of CoverTab[73968]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2152
			}())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2152
			// _ = "end of CoverTab[73967]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2152
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2152
			_go_fuzz_dep_.CoverTab[73969]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2163
			return nil, nil, sc.countError("bad_path_method", streamError(f.StreamID, ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2163
			// _ = "end of CoverTab[73969]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2164
			_go_fuzz_dep_.CoverTab[73970]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2164
			// _ = "end of CoverTab[73970]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2164
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2164
		// _ = "end of CoverTab[73965]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2164
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2164
	// _ = "end of CoverTab[73954]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2164
	_go_fuzz_dep_.CoverTab[73955]++

											rp.header = make(http.Header)
											for _, hf := range f.RegularFields() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2167
		_go_fuzz_dep_.CoverTab[73971]++
												rp.header.Add(sc.canonicalHeader(hf.Name), hf.Value)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2168
		// _ = "end of CoverTab[73971]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2169
	// _ = "end of CoverTab[73955]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2169
	_go_fuzz_dep_.CoverTab[73956]++
											if rp.authority == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2170
		_go_fuzz_dep_.CoverTab[73972]++
												rp.authority = rp.header.Get("Host")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2171
		// _ = "end of CoverTab[73972]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2172
		_go_fuzz_dep_.CoverTab[73973]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2172
		// _ = "end of CoverTab[73973]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2172
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2172
	// _ = "end of CoverTab[73956]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2172
	_go_fuzz_dep_.CoverTab[73957]++

											rw, req, err := sc.newWriterAndRequestNoBody(st, rp)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2175
		_go_fuzz_dep_.CoverTab[73974]++
												return nil, nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2176
		// _ = "end of CoverTab[73974]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2177
		_go_fuzz_dep_.CoverTab[73975]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2177
		// _ = "end of CoverTab[73975]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2177
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2177
	// _ = "end of CoverTab[73957]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2177
	_go_fuzz_dep_.CoverTab[73958]++
											bodyOpen := !f.StreamEnded()
											if bodyOpen {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2179
		_go_fuzz_dep_.CoverTab[73976]++
												if vv, ok := rp.header["Content-Length"]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2180
			_go_fuzz_dep_.CoverTab[73978]++
													if cl, err := strconv.ParseUint(vv[0], 10, 63); err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2181
				_go_fuzz_dep_.CoverTab[73979]++
														req.ContentLength = int64(cl)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2182
				// _ = "end of CoverTab[73979]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2183
				_go_fuzz_dep_.CoverTab[73980]++
														req.ContentLength = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2184
				// _ = "end of CoverTab[73980]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2185
			// _ = "end of CoverTab[73978]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2186
			_go_fuzz_dep_.CoverTab[73981]++
													req.ContentLength = -1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2187
			// _ = "end of CoverTab[73981]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2188
		// _ = "end of CoverTab[73976]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2188
		_go_fuzz_dep_.CoverTab[73977]++
												req.Body.(*requestBody).pipe = &pipe{
			b: &dataBuffer{expected: req.ContentLength},
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2191
		// _ = "end of CoverTab[73977]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2192
		_go_fuzz_dep_.CoverTab[73982]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2192
		// _ = "end of CoverTab[73982]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2192
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2192
	// _ = "end of CoverTab[73958]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2192
	_go_fuzz_dep_.CoverTab[73959]++
											return rw, req, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2193
	// _ = "end of CoverTab[73959]"
}

type requestParam struct {
	method			string
	scheme, authority, path	string
	header			http.Header
}

func (sc *serverConn) newWriterAndRequestNoBody(st *stream, rp requestParam) (*responseWriter, *http.Request, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2202
	_go_fuzz_dep_.CoverTab[73983]++
											sc.serveG.check()

											var tlsState *tls.ConnectionState	// nil if not scheme https
											if rp.scheme == "https" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2206
		_go_fuzz_dep_.CoverTab[73989]++
												tlsState = sc.tlsState
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2207
		// _ = "end of CoverTab[73989]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2208
		_go_fuzz_dep_.CoverTab[73990]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2208
		// _ = "end of CoverTab[73990]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2208
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2208
	// _ = "end of CoverTab[73983]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2208
	_go_fuzz_dep_.CoverTab[73984]++

											needsContinue := httpguts.HeaderValuesContainsToken(rp.header["Expect"], "100-continue")
											if needsContinue {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2211
		_go_fuzz_dep_.CoverTab[73991]++
												rp.header.Del("Expect")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2212
		// _ = "end of CoverTab[73991]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2213
		_go_fuzz_dep_.CoverTab[73992]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2213
		// _ = "end of CoverTab[73992]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2213
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2213
	// _ = "end of CoverTab[73984]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2213
	_go_fuzz_dep_.CoverTab[73985]++

											if cookies := rp.header["Cookie"]; len(cookies) > 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2215
		_go_fuzz_dep_.CoverTab[73993]++
												rp.header.Set("Cookie", strings.Join(cookies, "; "))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2216
		// _ = "end of CoverTab[73993]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2217
		_go_fuzz_dep_.CoverTab[73994]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2217
		// _ = "end of CoverTab[73994]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2217
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2217
	// _ = "end of CoverTab[73985]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2217
	_go_fuzz_dep_.CoverTab[73986]++

	// Setup Trailers
	var trailer http.Header
	for _, v := range rp.header["Trailer"] {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2221
		_go_fuzz_dep_.CoverTab[73995]++
												for _, key := range strings.Split(v, ",") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2222
			_go_fuzz_dep_.CoverTab[73996]++
													key = http.CanonicalHeaderKey(textproto.TrimString(key))
													switch key {
			case "Transfer-Encoding", "Trailer", "Content-Length":
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2225
				_go_fuzz_dep_.CoverTab[73997]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2225
				// _ = "end of CoverTab[73997]"

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2228
			default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2228
				_go_fuzz_dep_.CoverTab[73998]++
														if trailer == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2229
					_go_fuzz_dep_.CoverTab[74000]++
															trailer = make(http.Header)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2230
					// _ = "end of CoverTab[74000]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2231
					_go_fuzz_dep_.CoverTab[74001]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2231
					// _ = "end of CoverTab[74001]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2231
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2231
				// _ = "end of CoverTab[73998]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2231
				_go_fuzz_dep_.CoverTab[73999]++
														trailer[key] = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2232
				// _ = "end of CoverTab[73999]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2233
			// _ = "end of CoverTab[73996]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2234
		// _ = "end of CoverTab[73995]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2235
	// _ = "end of CoverTab[73986]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2235
	_go_fuzz_dep_.CoverTab[73987]++
											delete(rp.header, "Trailer")

											var url_ *url.URL
											var requestURI string
											if rp.method == "CONNECT" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2240
		_go_fuzz_dep_.CoverTab[74002]++
												url_ = &url.URL{Host: rp.authority}
												requestURI = rp.authority
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2242
		// _ = "end of CoverTab[74002]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2243
		_go_fuzz_dep_.CoverTab[74003]++
												var err error
												url_, err = url.ParseRequestURI(rp.path)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2246
			_go_fuzz_dep_.CoverTab[74005]++
													return nil, nil, sc.countError("bad_path", streamError(st.id, ErrCodeProtocol))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2247
			// _ = "end of CoverTab[74005]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2248
			_go_fuzz_dep_.CoverTab[74006]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2248
			// _ = "end of CoverTab[74006]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2248
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2248
		// _ = "end of CoverTab[74003]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2248
		_go_fuzz_dep_.CoverTab[74004]++
												requestURI = rp.path
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2249
		// _ = "end of CoverTab[74004]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2250
	// _ = "end of CoverTab[73987]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2250
	_go_fuzz_dep_.CoverTab[73988]++

											body := &requestBody{
		conn:		sc,
		stream:		st,
		needsContinue:	needsContinue,
	}
	req := &http.Request{
		Method:		rp.method,
		URL:		url_,
		RemoteAddr:	sc.remoteAddrStr,
		Header:		rp.header,
		RequestURI:	requestURI,
		Proto:		"HTTP/2.0",
		ProtoMajor:	2,
		ProtoMinor:	0,
		TLS:		tlsState,
		Host:		rp.authority,
		Body:		body,
		Trailer:	trailer,
	}
											req = req.WithContext(st.ctx)

											rw := sc.newResponseWriter(st, req)
											return rw, req, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2274
	// _ = "end of CoverTab[73988]"
}

func (sc *serverConn) newResponseWriter(st *stream, req *http.Request) *responseWriter {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2277
	_go_fuzz_dep_.CoverTab[74007]++
											rws := responseWriterStatePool.Get().(*responseWriterState)
											bwSave := rws.bw
											*rws = responseWriterState{}
											rws.conn = sc
											rws.bw = bwSave
											rws.bw.Reset(chunkWriter{rws})
											rws.stream = st
											rws.req = req
											return &responseWriter{rws: rws}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2286
	// _ = "end of CoverTab[74007]"
}

// Run on its own goroutine.
func (sc *serverConn) runHandler(rw *responseWriter, req *http.Request, handler func(http.ResponseWriter, *http.Request)) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2290
	_go_fuzz_dep_.CoverTab[74008]++
											didPanic := true
											defer func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2292
		_go_fuzz_dep_.CoverTab[74010]++
												rw.rws.stream.cancelCtx()
												if req.MultipartForm != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2294
			_go_fuzz_dep_.CoverTab[74013]++
													req.MultipartForm.RemoveAll()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2295
			// _ = "end of CoverTab[74013]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2296
			_go_fuzz_dep_.CoverTab[74014]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2296
			// _ = "end of CoverTab[74014]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2296
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2296
		// _ = "end of CoverTab[74010]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2296
		_go_fuzz_dep_.CoverTab[74011]++
												if didPanic {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2297
			_go_fuzz_dep_.CoverTab[74015]++
													e := recover()
													sc.writeFrameFromHandler(FrameWriteRequest{
				write:	handlerPanicRST{rw.rws.stream.id},
				stream:	rw.rws.stream,
			})

			if e != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2304
				_go_fuzz_dep_.CoverTab[74017]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2304
				return e != http.ErrAbortHandler
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2304
				// _ = "end of CoverTab[74017]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2304
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2304
				_go_fuzz_dep_.CoverTab[74018]++
														const size = 64 << 10
														buf := make([]byte, size)
														buf = buf[:runtime.Stack(buf, false)]
														sc.logf("http2: panic serving %v: %v\n%s", sc.conn.RemoteAddr(), e, buf)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2308
				// _ = "end of CoverTab[74018]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2309
				_go_fuzz_dep_.CoverTab[74019]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2309
				// _ = "end of CoverTab[74019]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2309
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2309
			// _ = "end of CoverTab[74015]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2309
			_go_fuzz_dep_.CoverTab[74016]++
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2310
			// _ = "end of CoverTab[74016]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2311
			_go_fuzz_dep_.CoverTab[74020]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2311
			// _ = "end of CoverTab[74020]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2311
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2311
		// _ = "end of CoverTab[74011]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2311
		_go_fuzz_dep_.CoverTab[74012]++
												rw.handlerDone()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2312
		// _ = "end of CoverTab[74012]"
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2313
	// _ = "end of CoverTab[74008]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2313
	_go_fuzz_dep_.CoverTab[74009]++
											handler(rw, req)
											didPanic = false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2315
	// _ = "end of CoverTab[74009]"
}

func handleHeaderListTooLong(w http.ResponseWriter, r *http.Request) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2318
	_go_fuzz_dep_.CoverTab[74021]++
	// 10.5.1 Limits on Header Block Size:
	// .. "A server that receives a larger header block than it is
	// willing to handle can send an HTTP 431 (Request Header Fields Too
											// Large) status code"
											const statusRequestHeaderFieldsTooLarge = 431	// only in Go 1.6+
											w.WriteHeader(statusRequestHeaderFieldsTooLarge)
											io.WriteString(w, "<h1>HTTP Error 431</h1><p>Request Header Field(s) Too Large</p>")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2325
	// _ = "end of CoverTab[74021]"
}

// called from handler goroutines.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2328
// h may be nil.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2330
func (sc *serverConn) writeHeaders(st *stream, headerData *writeResHeaders) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2330
	_go_fuzz_dep_.CoverTab[74022]++
											sc.serveG.checkNotOn()
											var errc chan error
											if headerData.h != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2333
		_go_fuzz_dep_.CoverTab[74026]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2338
		errc = errChanPool.Get().(chan error)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2338
		// _ = "end of CoverTab[74026]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2339
		_go_fuzz_dep_.CoverTab[74027]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2339
		// _ = "end of CoverTab[74027]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2339
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2339
	// _ = "end of CoverTab[74022]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2339
	_go_fuzz_dep_.CoverTab[74023]++
											if err := sc.writeFrameFromHandler(FrameWriteRequest{
		write:	headerData,
		stream:	st,
		done:	errc,
	}); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2344
		_go_fuzz_dep_.CoverTab[74028]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2345
		// _ = "end of CoverTab[74028]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2346
		_go_fuzz_dep_.CoverTab[74029]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2346
		// _ = "end of CoverTab[74029]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2346
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2346
	// _ = "end of CoverTab[74023]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2346
	_go_fuzz_dep_.CoverTab[74024]++
											if errc != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2347
		_go_fuzz_dep_.CoverTab[74030]++
												select {
		case err := <-errc:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2349
			_go_fuzz_dep_.CoverTab[74031]++
													errChanPool.Put(errc)
													return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2351
			// _ = "end of CoverTab[74031]"
		case <-sc.doneServing:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2352
			_go_fuzz_dep_.CoverTab[74032]++
													return errClientDisconnected
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2353
			// _ = "end of CoverTab[74032]"
		case <-st.cw:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2354
			_go_fuzz_dep_.CoverTab[74033]++
													return errStreamClosed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2355
			// _ = "end of CoverTab[74033]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2356
		// _ = "end of CoverTab[74030]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2357
		_go_fuzz_dep_.CoverTab[74034]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2357
		// _ = "end of CoverTab[74034]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2357
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2357
	// _ = "end of CoverTab[74024]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2357
	_go_fuzz_dep_.CoverTab[74025]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2358
	// _ = "end of CoverTab[74025]"
}

// called from handler goroutines.
func (sc *serverConn) write100ContinueHeaders(st *stream) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2362
	_go_fuzz_dep_.CoverTab[74035]++
											sc.writeFrameFromHandler(FrameWriteRequest{
		write:	write100ContinueHeadersFrame{st.id},
		stream:	st,
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2366
	// _ = "end of CoverTab[74035]"
}

// A bodyReadMsg tells the server loop that the http.Handler read n
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2369
// bytes of the DATA from the client on the given stream.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2371
type bodyReadMsg struct {
	st	*stream
	n	int
}

// called from handler goroutines.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2376
// Notes that the handler for the given stream ID read n bytes of its body
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2376
// and schedules flow control tokens to be sent.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2379
func (sc *serverConn) noteBodyReadFromHandler(st *stream, n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2379
	_go_fuzz_dep_.CoverTab[74036]++
											sc.serveG.checkNotOn()
											if n > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2381
		_go_fuzz_dep_.CoverTab[74037]++
												select {
		case sc.bodyReadCh <- bodyReadMsg{st, n}:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2383
			_go_fuzz_dep_.CoverTab[74038]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2383
			// _ = "end of CoverTab[74038]"
		case <-sc.doneServing:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2384
			_go_fuzz_dep_.CoverTab[74039]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2384
			// _ = "end of CoverTab[74039]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2385
		// _ = "end of CoverTab[74037]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2386
		_go_fuzz_dep_.CoverTab[74040]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2386
		// _ = "end of CoverTab[74040]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2386
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2386
	// _ = "end of CoverTab[74036]"
}

func (sc *serverConn) noteBodyRead(st *stream, n int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2389
	_go_fuzz_dep_.CoverTab[74041]++
											sc.serveG.check()
											sc.sendWindowUpdate(nil, n)
											if st.state != stateHalfClosedRemote && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2392
		_go_fuzz_dep_.CoverTab[74042]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2392
		return st.state != stateClosed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2392
		// _ = "end of CoverTab[74042]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2392
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2392
		_go_fuzz_dep_.CoverTab[74043]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2395
		sc.sendWindowUpdate(st, n)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2395
		// _ = "end of CoverTab[74043]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2396
		_go_fuzz_dep_.CoverTab[74044]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2396
		// _ = "end of CoverTab[74044]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2396
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2396
	// _ = "end of CoverTab[74041]"
}

// st may be nil for conn-level
func (sc *serverConn) sendWindowUpdate32(st *stream, n int32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2400
	_go_fuzz_dep_.CoverTab[74045]++
											sc.sendWindowUpdate(st, int(n))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2401
	// _ = "end of CoverTab[74045]"
}

// st may be nil for conn-level
func (sc *serverConn) sendWindowUpdate(st *stream, n int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2405
	_go_fuzz_dep_.CoverTab[74046]++
											sc.serveG.check()
											var streamID uint32
											var send int32
											if st == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2409
		_go_fuzz_dep_.CoverTab[74049]++
												send = sc.inflow.add(n)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2410
		// _ = "end of CoverTab[74049]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2411
		_go_fuzz_dep_.CoverTab[74050]++
												streamID = st.id
												send = st.inflow.add(n)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2413
		// _ = "end of CoverTab[74050]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2414
	// _ = "end of CoverTab[74046]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2414
	_go_fuzz_dep_.CoverTab[74047]++
											if send == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2415
		_go_fuzz_dep_.CoverTab[74051]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2416
		// _ = "end of CoverTab[74051]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2417
		_go_fuzz_dep_.CoverTab[74052]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2417
		// _ = "end of CoverTab[74052]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2417
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2417
	// _ = "end of CoverTab[74047]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2417
	_go_fuzz_dep_.CoverTab[74048]++
											sc.writeFrame(FrameWriteRequest{
		write:	writeWindowUpdate{streamID: streamID, n: uint32(send)},
		stream:	st,
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2421
	// _ = "end of CoverTab[74048]"
}

// requestBody is the Handler's Request.Body type.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2424
// Read and Close may be called concurrently.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2426
type requestBody struct {
	_		incomparable
	stream		*stream
	conn		*serverConn
	closeOnce	sync.Once	// for use by Close only
	sawEOF		bool		// for use by Read only
	pipe		*pipe		// non-nil if we have a HTTP entity message body
	needsContinue	bool		// need to send a 100-continue
}

func (b *requestBody) Close() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2436
	_go_fuzz_dep_.CoverTab[74053]++
											b.closeOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2437
		_go_fuzz_dep_.CoverTab[74055]++
												if b.pipe != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2438
			_go_fuzz_dep_.CoverTab[74056]++
													b.pipe.BreakWithError(errClosedBody)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2439
			// _ = "end of CoverTab[74056]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2440
			_go_fuzz_dep_.CoverTab[74057]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2440
			// _ = "end of CoverTab[74057]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2440
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2440
		// _ = "end of CoverTab[74055]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2441
	// _ = "end of CoverTab[74053]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2441
	_go_fuzz_dep_.CoverTab[74054]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2442
	// _ = "end of CoverTab[74054]"
}

func (b *requestBody) Read(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2445
	_go_fuzz_dep_.CoverTab[74058]++
											if b.needsContinue {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2446
		_go_fuzz_dep_.CoverTab[74063]++
												b.needsContinue = false
												b.conn.write100ContinueHeaders(b.stream)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2448
		// _ = "end of CoverTab[74063]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2449
		_go_fuzz_dep_.CoverTab[74064]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2449
		// _ = "end of CoverTab[74064]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2449
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2449
	// _ = "end of CoverTab[74058]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2449
	_go_fuzz_dep_.CoverTab[74059]++
											if b.pipe == nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2450
		_go_fuzz_dep_.CoverTab[74065]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2450
		return b.sawEOF
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2450
		// _ = "end of CoverTab[74065]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2450
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2450
		_go_fuzz_dep_.CoverTab[74066]++
												return 0, io.EOF
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2451
		// _ = "end of CoverTab[74066]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2452
		_go_fuzz_dep_.CoverTab[74067]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2452
		// _ = "end of CoverTab[74067]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2452
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2452
	// _ = "end of CoverTab[74059]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2452
	_go_fuzz_dep_.CoverTab[74060]++
											n, err = b.pipe.Read(p)
											if err == io.EOF {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2454
		_go_fuzz_dep_.CoverTab[74068]++
												b.sawEOF = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2455
		// _ = "end of CoverTab[74068]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2456
		_go_fuzz_dep_.CoverTab[74069]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2456
		// _ = "end of CoverTab[74069]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2456
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2456
	// _ = "end of CoverTab[74060]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2456
	_go_fuzz_dep_.CoverTab[74061]++
											if b.conn == nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2457
		_go_fuzz_dep_.CoverTab[74070]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2457
		return inTests
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2457
		// _ = "end of CoverTab[74070]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2457
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2457
		_go_fuzz_dep_.CoverTab[74071]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2458
		// _ = "end of CoverTab[74071]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2459
		_go_fuzz_dep_.CoverTab[74072]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2459
		// _ = "end of CoverTab[74072]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2459
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2459
	// _ = "end of CoverTab[74061]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2459
	_go_fuzz_dep_.CoverTab[74062]++
											b.conn.noteBodyReadFromHandler(b.stream, n, err)
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2461
	// _ = "end of CoverTab[74062]"
}

// responseWriter is the http.ResponseWriter implementation. It's
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2464
// intentionally small (1 pointer wide) to minimize garbage. The
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2464
// responseWriterState pointer inside is zeroed at the end of a
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2464
// request (in handlerDone) and calls on the responseWriter thereafter
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2464
// simply crash (caller's mistake), but the much larger responseWriterState
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2464
// and buffers are reused between multiple requests.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2470
type responseWriter struct {
	rws *responseWriterState
}

// Optional http.ResponseWriter interfaces implemented.
var (
	_	http.CloseNotifier	= (*responseWriter)(nil)
	_	http.Flusher		= (*responseWriter)(nil)
	_	stringWriter		= (*responseWriter)(nil)
)

type responseWriterState struct {
	// immutable within a request:
	stream	*stream
	req	*http.Request
	conn	*serverConn

	// TODO: adjust buffer writing sizes based on server config, frame size updates from peer, etc
	bw	*bufio.Writer	// writing to a chunkWriter{this *responseWriterState}

	// mutated by http.Handler goroutine:
	handlerHeader	http.Header	// nil until called
	snapHeader	http.Header	// snapshot of handlerHeader at WriteHeader time
	trailers	[]string	// set in writeChunk
	status		int		// status code passed to WriteHeader
	wroteHeader	bool		// WriteHeader called (explicitly or implicitly). Not necessarily sent to user yet.
	sentHeader	bool		// have we sent the header frame?
	handlerDone	bool		// handler has finished
	dirty		bool		// a Write failed; don't reuse this responseWriterState

	sentContentLen	int64	// non-zero if handler set a Content-Length header
	wroteBytes	int64

	closeNotifierMu	sync.Mutex	// guards closeNotifierCh
	closeNotifierCh	chan bool	// nil until first used
}

type chunkWriter struct{ rws *responseWriterState }

func (cw chunkWriter) Write(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2509
	_go_fuzz_dep_.CoverTab[74073]++
											n, err = cw.rws.writeChunk(p)
											if err == errStreamClosed {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2511
		_go_fuzz_dep_.CoverTab[74075]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2514
		err = cw.rws.stream.closeErr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2514
		// _ = "end of CoverTab[74075]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2515
		_go_fuzz_dep_.CoverTab[74076]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2515
		// _ = "end of CoverTab[74076]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2515
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2515
	// _ = "end of CoverTab[74073]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2515
	_go_fuzz_dep_.CoverTab[74074]++
											return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2516
	// _ = "end of CoverTab[74074]"
}

func (rws *responseWriterState) hasTrailers() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2519
	_go_fuzz_dep_.CoverTab[74077]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2519
	return len(rws.trailers) > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2519
	// _ = "end of CoverTab[74077]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2519
}

func (rws *responseWriterState) hasNonemptyTrailers() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2521
	_go_fuzz_dep_.CoverTab[74078]++
											for _, trailer := range rws.trailers {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2522
		_go_fuzz_dep_.CoverTab[74080]++
												if _, ok := rws.handlerHeader[trailer]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2523
			_go_fuzz_dep_.CoverTab[74081]++
													return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2524
			// _ = "end of CoverTab[74081]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2525
			_go_fuzz_dep_.CoverTab[74082]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2525
			// _ = "end of CoverTab[74082]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2525
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2525
		// _ = "end of CoverTab[74080]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2526
	// _ = "end of CoverTab[74078]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2526
	_go_fuzz_dep_.CoverTab[74079]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2527
	// _ = "end of CoverTab[74079]"
}

// declareTrailer is called for each Trailer header when the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2530
// response header is written. It notes that a header will need to be
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2530
// written in the trailers at the end of the response.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2533
func (rws *responseWriterState) declareTrailer(k string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2533
	_go_fuzz_dep_.CoverTab[74083]++
											k = http.CanonicalHeaderKey(k)
											if !httpguts.ValidTrailerHeader(k) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2535
		_go_fuzz_dep_.CoverTab[74085]++

												rws.conn.logf("ignoring invalid trailer %q", k)
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2538
		// _ = "end of CoverTab[74085]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2539
		_go_fuzz_dep_.CoverTab[74086]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2539
		// _ = "end of CoverTab[74086]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2539
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2539
	// _ = "end of CoverTab[74083]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2539
	_go_fuzz_dep_.CoverTab[74084]++
											if !strSliceContains(rws.trailers, k) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2540
		_go_fuzz_dep_.CoverTab[74087]++
												rws.trailers = append(rws.trailers, k)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2541
		// _ = "end of CoverTab[74087]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2542
		_go_fuzz_dep_.CoverTab[74088]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2542
		// _ = "end of CoverTab[74088]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2542
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2542
	// _ = "end of CoverTab[74084]"
}

// writeChunk writes chunks from the bufio.Writer. But because
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2545
// bufio.Writer may bypass its chunking, sometimes p may be
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2545
// arbitrarily large.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2545
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2545
// writeChunk is also responsible (on the first chunk) for sending the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2545
// HEADER response.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2551
func (rws *responseWriterState) writeChunk(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2551
	_go_fuzz_dep_.CoverTab[74089]++
											if !rws.wroteHeader {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2552
		_go_fuzz_dep_.CoverTab[74097]++
												rws.writeHeader(200)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2553
		// _ = "end of CoverTab[74097]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2554
		_go_fuzz_dep_.CoverTab[74098]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2554
		// _ = "end of CoverTab[74098]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2554
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2554
	// _ = "end of CoverTab[74089]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2554
	_go_fuzz_dep_.CoverTab[74090]++

											if rws.handlerDone {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2556
		_go_fuzz_dep_.CoverTab[74099]++
												rws.promoteUndeclaredTrailers()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2557
		// _ = "end of CoverTab[74099]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2558
		_go_fuzz_dep_.CoverTab[74100]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2558
		// _ = "end of CoverTab[74100]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2558
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2558
	// _ = "end of CoverTab[74090]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2558
	_go_fuzz_dep_.CoverTab[74091]++

											isHeadResp := rws.req.Method == "HEAD"
											if !rws.sentHeader {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2561
		_go_fuzz_dep_.CoverTab[74101]++
												rws.sentHeader = true
												var ctype, clen string
												if clen = rws.snapHeader.Get("Content-Length"); clen != "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2564
			_go_fuzz_dep_.CoverTab[74109]++
													rws.snapHeader.Del("Content-Length")
													if cl, err := strconv.ParseUint(clen, 10, 63); err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2566
				_go_fuzz_dep_.CoverTab[74110]++
														rws.sentContentLen = int64(cl)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2567
				// _ = "end of CoverTab[74110]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2568
				_go_fuzz_dep_.CoverTab[74111]++
														clen = ""
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2569
				// _ = "end of CoverTab[74111]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2570
			// _ = "end of CoverTab[74109]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2571
			_go_fuzz_dep_.CoverTab[74112]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2571
			// _ = "end of CoverTab[74112]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2571
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2571
		// _ = "end of CoverTab[74101]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2571
		_go_fuzz_dep_.CoverTab[74102]++
												if clen == "" && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
			_go_fuzz_dep_.CoverTab[74113]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
			return rws.handlerDone
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
			// _ = "end of CoverTab[74113]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
			_go_fuzz_dep_.CoverTab[74114]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
			return bodyAllowedForStatus(rws.status)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
			// _ = "end of CoverTab[74114]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
			_go_fuzz_dep_.CoverTab[74115]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
			return (len(p) > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
				_go_fuzz_dep_.CoverTab[74116]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
				return !isHeadResp
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
				// _ = "end of CoverTab[74116]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
			}())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
			// _ = "end of CoverTab[74115]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2572
			_go_fuzz_dep_.CoverTab[74117]++
													clen = strconv.Itoa(len(p))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2573
			// _ = "end of CoverTab[74117]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2574
			_go_fuzz_dep_.CoverTab[74118]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2574
			// _ = "end of CoverTab[74118]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2574
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2574
		// _ = "end of CoverTab[74102]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2574
		_go_fuzz_dep_.CoverTab[74103]++
												_, hasContentType := rws.snapHeader["Content-Type"]

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2578
		ce := rws.snapHeader.Get("Content-Encoding")
		hasCE := len(ce) > 0
		if !hasCE && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2580
			_go_fuzz_dep_.CoverTab[74119]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2580
			return !hasContentType
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2580
			// _ = "end of CoverTab[74119]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2580
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2580
			_go_fuzz_dep_.CoverTab[74120]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2580
			return bodyAllowedForStatus(rws.status)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2580
			// _ = "end of CoverTab[74120]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2580
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2580
			_go_fuzz_dep_.CoverTab[74121]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2580
			return len(p) > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2580
			// _ = "end of CoverTab[74121]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2580
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2580
			_go_fuzz_dep_.CoverTab[74122]++
													ctype = http.DetectContentType(p)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2581
			// _ = "end of CoverTab[74122]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2582
			_go_fuzz_dep_.CoverTab[74123]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2582
			// _ = "end of CoverTab[74123]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2582
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2582
		// _ = "end of CoverTab[74103]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2582
		_go_fuzz_dep_.CoverTab[74104]++
												var date string
												if _, ok := rws.snapHeader["Date"]; !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2584
			_go_fuzz_dep_.CoverTab[74124]++

													date = time.Now().UTC().Format(http.TimeFormat)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2586
			// _ = "end of CoverTab[74124]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2587
			_go_fuzz_dep_.CoverTab[74125]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2587
			// _ = "end of CoverTab[74125]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2587
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2587
		// _ = "end of CoverTab[74104]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2587
		_go_fuzz_dep_.CoverTab[74105]++

												for _, v := range rws.snapHeader["Trailer"] {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2589
			_go_fuzz_dep_.CoverTab[74126]++
													foreachHeaderElement(v, rws.declareTrailer)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2590
			// _ = "end of CoverTab[74126]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2591
		// _ = "end of CoverTab[74105]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2591
		_go_fuzz_dep_.CoverTab[74106]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2598
		if _, ok := rws.snapHeader["Connection"]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2598
			_go_fuzz_dep_.CoverTab[74127]++
													v := rws.snapHeader.Get("Connection")
													delete(rws.snapHeader, "Connection")
													if v == "close" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2601
				_go_fuzz_dep_.CoverTab[74128]++
														rws.conn.startGracefulShutdown()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2602
				// _ = "end of CoverTab[74128]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2603
				_go_fuzz_dep_.CoverTab[74129]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2603
				// _ = "end of CoverTab[74129]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2603
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2603
			// _ = "end of CoverTab[74127]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2604
			_go_fuzz_dep_.CoverTab[74130]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2604
			// _ = "end of CoverTab[74130]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2604
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2604
		// _ = "end of CoverTab[74106]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2604
		_go_fuzz_dep_.CoverTab[74107]++

												endStream := (rws.handlerDone && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2606
			_go_fuzz_dep_.CoverTab[74131]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2606
			return !rws.hasTrailers()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2606
			// _ = "end of CoverTab[74131]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2606
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2606
			_go_fuzz_dep_.CoverTab[74132]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2606
			return len(p) == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2606
			// _ = "end of CoverTab[74132]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2606
		}()) || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2606
			_go_fuzz_dep_.CoverTab[74133]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2606
			return isHeadResp
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2606
			// _ = "end of CoverTab[74133]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2606
		}()
												err = rws.conn.writeHeaders(rws.stream, &writeResHeaders{
			streamID:	rws.stream.id,
			httpResCode:	rws.status,
			h:		rws.snapHeader,
			endStream:	endStream,
			contentType:	ctype,
			contentLength:	clen,
			date:		date,
		})
		if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2616
			_go_fuzz_dep_.CoverTab[74134]++
													rws.dirty = true
													return 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2618
			// _ = "end of CoverTab[74134]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2619
			_go_fuzz_dep_.CoverTab[74135]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2619
			// _ = "end of CoverTab[74135]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2619
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2619
		// _ = "end of CoverTab[74107]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2619
		_go_fuzz_dep_.CoverTab[74108]++
												if endStream {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2620
			_go_fuzz_dep_.CoverTab[74136]++
													return 0, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2621
			// _ = "end of CoverTab[74136]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2622
			_go_fuzz_dep_.CoverTab[74137]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2622
			// _ = "end of CoverTab[74137]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2622
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2622
		// _ = "end of CoverTab[74108]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2623
		_go_fuzz_dep_.CoverTab[74138]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2623
		// _ = "end of CoverTab[74138]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2623
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2623
	// _ = "end of CoverTab[74091]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2623
	_go_fuzz_dep_.CoverTab[74092]++
											if isHeadResp {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2624
		_go_fuzz_dep_.CoverTab[74139]++
												return len(p), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2625
		// _ = "end of CoverTab[74139]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2626
		_go_fuzz_dep_.CoverTab[74140]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2626
		// _ = "end of CoverTab[74140]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2626
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2626
	// _ = "end of CoverTab[74092]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2626
	_go_fuzz_dep_.CoverTab[74093]++
											if len(p) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2627
		_go_fuzz_dep_.CoverTab[74141]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2627
		return !rws.handlerDone
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2627
		// _ = "end of CoverTab[74141]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2627
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2627
		_go_fuzz_dep_.CoverTab[74142]++
												return 0, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2628
		// _ = "end of CoverTab[74142]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2629
		_go_fuzz_dep_.CoverTab[74143]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2629
		// _ = "end of CoverTab[74143]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2629
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2629
	// _ = "end of CoverTab[74093]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2629
	_go_fuzz_dep_.CoverTab[74094]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2633
	hasNonemptyTrailers := rws.hasNonemptyTrailers()
	endStream := rws.handlerDone && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2634
		_go_fuzz_dep_.CoverTab[74144]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2634
		return !hasNonemptyTrailers
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2634
		// _ = "end of CoverTab[74144]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2634
	}()
											if len(p) > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2635
		_go_fuzz_dep_.CoverTab[74145]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2635
		return endStream
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2635
		// _ = "end of CoverTab[74145]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2635
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2635
		_go_fuzz_dep_.CoverTab[74146]++

												if err := rws.conn.writeDataFromHandler(rws.stream, p, endStream); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2637
			_go_fuzz_dep_.CoverTab[74147]++
													rws.dirty = true
													return 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2639
			// _ = "end of CoverTab[74147]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2640
			_go_fuzz_dep_.CoverTab[74148]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2640
			// _ = "end of CoverTab[74148]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2640
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2640
		// _ = "end of CoverTab[74146]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2641
		_go_fuzz_dep_.CoverTab[74149]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2641
		// _ = "end of CoverTab[74149]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2641
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2641
	// _ = "end of CoverTab[74094]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2641
	_go_fuzz_dep_.CoverTab[74095]++

											if rws.handlerDone && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2643
		_go_fuzz_dep_.CoverTab[74150]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2643
		return hasNonemptyTrailers
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2643
		// _ = "end of CoverTab[74150]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2643
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2643
		_go_fuzz_dep_.CoverTab[74151]++
												err = rws.conn.writeHeaders(rws.stream, &writeResHeaders{
			streamID:	rws.stream.id,
			h:		rws.handlerHeader,
			trailers:	rws.trailers,
			endStream:	true,
		})
		if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2650
			_go_fuzz_dep_.CoverTab[74153]++
													rws.dirty = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2651
			// _ = "end of CoverTab[74153]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2652
			_go_fuzz_dep_.CoverTab[74154]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2652
			// _ = "end of CoverTab[74154]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2652
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2652
		// _ = "end of CoverTab[74151]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2652
		_go_fuzz_dep_.CoverTab[74152]++
												return len(p), err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2653
		// _ = "end of CoverTab[74152]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2654
		_go_fuzz_dep_.CoverTab[74155]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2654
		// _ = "end of CoverTab[74155]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2654
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2654
	// _ = "end of CoverTab[74095]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2654
	_go_fuzz_dep_.CoverTab[74096]++
											return len(p), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2655
	// _ = "end of CoverTab[74096]"
}

// TrailerPrefix is a magic prefix for ResponseWriter.Header map keys
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2658
// that, if present, signals that the map entry is actually for
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2658
// the response trailers, and not the response headers. The prefix
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2658
// is stripped after the ServeHTTP call finishes and the values are
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2658
// sent in the trailers.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2658
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2658
// This mechanism is intended only for trailers that are not known
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2658
// prior to the headers being written. If the set of trailers is fixed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2658
// or known before the header is written, the normal Go trailers mechanism
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2658
// is preferred:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2658
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2658
//	https://golang.org/pkg/net/http/#ResponseWriter
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2658
//	https://golang.org/pkg/net/http/#example_ResponseWriter_trailers
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2671
const TrailerPrefix = "Trailer:"

// promoteUndeclaredTrailers permits http.Handlers to set trailers
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// after the header has already been flushed. Because the Go
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// ResponseWriter interface has no way to set Trailers (only the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// Header), and because we didn't want to expand the ResponseWriter
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// interface, and because nobody used trailers, and because RFC 7230
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// says you SHOULD (but not must) predeclare any trailers in the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// header, the official ResponseWriter rules said trailers in Go must
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// be predeclared, and then we reuse the same ResponseWriter.Header()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// map to mean both Headers and Trailers. When it's time to write the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// Trailers, we pick out the fields of Headers that were declared as
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// trailers. That worked for a while, until we found the first major
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// user of Trailers in the wild: gRPC (using them only over http2),
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// and gRPC libraries permit setting trailers mid-stream without
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// predeclaring them. So: change of plans. We still permit the old
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// way, but we also permit this hack: if a Header() key begins with
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// "Trailer:", the suffix of that key is a Trailer. Because ':' is an
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// invalid token byte anyway, there is no ambiguity. (And it's already
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// filtered out) It's mildly hacky, but not terrible.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// This method runs after the Handler is done and promotes any Header
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2673
// fields to be trailers.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2694
func (rws *responseWriterState) promoteUndeclaredTrailers() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2694
	_go_fuzz_dep_.CoverTab[74156]++
											for k, vv := range rws.handlerHeader {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2695
		_go_fuzz_dep_.CoverTab[74158]++
												if !strings.HasPrefix(k, TrailerPrefix) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2696
			_go_fuzz_dep_.CoverTab[74160]++
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2697
			// _ = "end of CoverTab[74160]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2698
			_go_fuzz_dep_.CoverTab[74161]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2698
			// _ = "end of CoverTab[74161]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2698
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2698
		// _ = "end of CoverTab[74158]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2698
		_go_fuzz_dep_.CoverTab[74159]++
												trailerKey := strings.TrimPrefix(k, TrailerPrefix)
												rws.declareTrailer(trailerKey)
												rws.handlerHeader[http.CanonicalHeaderKey(trailerKey)] = vv
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2701
		// _ = "end of CoverTab[74159]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2702
	// _ = "end of CoverTab[74156]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2702
	_go_fuzz_dep_.CoverTab[74157]++

											if len(rws.trailers) > 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2704
		_go_fuzz_dep_.CoverTab[74162]++
												sorter := sorterPool.Get().(*sorter)
												sorter.SortStrings(rws.trailers)
												sorterPool.Put(sorter)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2707
		// _ = "end of CoverTab[74162]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2708
		_go_fuzz_dep_.CoverTab[74163]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2708
		// _ = "end of CoverTab[74163]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2708
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2708
	// _ = "end of CoverTab[74157]"
}

func (w *responseWriter) SetReadDeadline(deadline time.Time) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2711
	_go_fuzz_dep_.CoverTab[74164]++
											st := w.rws.stream
											if !deadline.IsZero() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2713
		_go_fuzz_dep_.CoverTab[74167]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2713
		return deadline.Before(time.Now())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2713
		// _ = "end of CoverTab[74167]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2713
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2713
		_go_fuzz_dep_.CoverTab[74168]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2716
		st.onReadTimeout()
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2717
		// _ = "end of CoverTab[74168]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2718
		_go_fuzz_dep_.CoverTab[74169]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2718
		// _ = "end of CoverTab[74169]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2718
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2718
	// _ = "end of CoverTab[74164]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2718
	_go_fuzz_dep_.CoverTab[74165]++
											w.rws.conn.sendServeMsg(func(sc *serverConn) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2719
		_go_fuzz_dep_.CoverTab[74170]++
												if st.readDeadline != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2720
			_go_fuzz_dep_.CoverTab[74172]++
													if !st.readDeadline.Stop() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2721
				_go_fuzz_dep_.CoverTab[74173]++

														return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2723
				// _ = "end of CoverTab[74173]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2724
				_go_fuzz_dep_.CoverTab[74174]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2724
				// _ = "end of CoverTab[74174]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2724
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2724
			// _ = "end of CoverTab[74172]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2725
			_go_fuzz_dep_.CoverTab[74175]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2725
			// _ = "end of CoverTab[74175]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2725
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2725
		// _ = "end of CoverTab[74170]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2725
		_go_fuzz_dep_.CoverTab[74171]++
												if deadline.IsZero() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2726
			_go_fuzz_dep_.CoverTab[74176]++
													st.readDeadline = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2727
			// _ = "end of CoverTab[74176]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2728
			_go_fuzz_dep_.CoverTab[74177]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2728
			if st.readDeadline == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2728
				_go_fuzz_dep_.CoverTab[74178]++
														st.readDeadline = time.AfterFunc(deadline.Sub(time.Now()), st.onReadTimeout)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2729
				// _ = "end of CoverTab[74178]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2730
				_go_fuzz_dep_.CoverTab[74179]++
														st.readDeadline.Reset(deadline.Sub(time.Now()))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2731
				// _ = "end of CoverTab[74179]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2732
			// _ = "end of CoverTab[74177]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2732
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2732
		// _ = "end of CoverTab[74171]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2733
	// _ = "end of CoverTab[74165]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2733
	_go_fuzz_dep_.CoverTab[74166]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2734
	// _ = "end of CoverTab[74166]"
}

func (w *responseWriter) SetWriteDeadline(deadline time.Time) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2737
	_go_fuzz_dep_.CoverTab[74180]++
											st := w.rws.stream
											if !deadline.IsZero() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2739
		_go_fuzz_dep_.CoverTab[74183]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2739
		return deadline.Before(time.Now())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2739
		// _ = "end of CoverTab[74183]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2739
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2739
		_go_fuzz_dep_.CoverTab[74184]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2742
		st.onWriteTimeout()
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2743
		// _ = "end of CoverTab[74184]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2744
		_go_fuzz_dep_.CoverTab[74185]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2744
		// _ = "end of CoverTab[74185]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2744
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2744
	// _ = "end of CoverTab[74180]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2744
	_go_fuzz_dep_.CoverTab[74181]++
											w.rws.conn.sendServeMsg(func(sc *serverConn) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2745
		_go_fuzz_dep_.CoverTab[74186]++
												if st.writeDeadline != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2746
			_go_fuzz_dep_.CoverTab[74188]++
													if !st.writeDeadline.Stop() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2747
				_go_fuzz_dep_.CoverTab[74189]++

														return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2749
				// _ = "end of CoverTab[74189]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2750
				_go_fuzz_dep_.CoverTab[74190]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2750
				// _ = "end of CoverTab[74190]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2750
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2750
			// _ = "end of CoverTab[74188]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2751
			_go_fuzz_dep_.CoverTab[74191]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2751
			// _ = "end of CoverTab[74191]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2751
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2751
		// _ = "end of CoverTab[74186]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2751
		_go_fuzz_dep_.CoverTab[74187]++
												if deadline.IsZero() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2752
			_go_fuzz_dep_.CoverTab[74192]++
													st.writeDeadline = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2753
			// _ = "end of CoverTab[74192]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2754
			_go_fuzz_dep_.CoverTab[74193]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2754
			if st.writeDeadline == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2754
				_go_fuzz_dep_.CoverTab[74194]++
														st.writeDeadline = time.AfterFunc(deadline.Sub(time.Now()), st.onWriteTimeout)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2755
				// _ = "end of CoverTab[74194]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2756
				_go_fuzz_dep_.CoverTab[74195]++
														st.writeDeadline.Reset(deadline.Sub(time.Now()))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2757
				// _ = "end of CoverTab[74195]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2758
			// _ = "end of CoverTab[74193]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2758
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2758
		// _ = "end of CoverTab[74187]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2759
	// _ = "end of CoverTab[74181]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2759
	_go_fuzz_dep_.CoverTab[74182]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2760
	// _ = "end of CoverTab[74182]"
}

func (w *responseWriter) Flush() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2763
	_go_fuzz_dep_.CoverTab[74196]++
											w.FlushError()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2764
	// _ = "end of CoverTab[74196]"
}

func (w *responseWriter) FlushError() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2767
	_go_fuzz_dep_.CoverTab[74197]++
											rws := w.rws
											if rws == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2769
		_go_fuzz_dep_.CoverTab[74200]++
												panic("Header called after Handler finished")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2770
		// _ = "end of CoverTab[74200]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2771
		_go_fuzz_dep_.CoverTab[74201]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2771
		// _ = "end of CoverTab[74201]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2771
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2771
	// _ = "end of CoverTab[74197]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2771
	_go_fuzz_dep_.CoverTab[74198]++
											var err error
											if rws.bw.Buffered() > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2773
		_go_fuzz_dep_.CoverTab[74202]++
												err = rws.bw.Flush()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2774
		// _ = "end of CoverTab[74202]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2775
		_go_fuzz_dep_.CoverTab[74203]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2780
		_, err = chunkWriter{rws}.Write(nil)
		if err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2781
			_go_fuzz_dep_.CoverTab[74204]++
													select {
			case <-rws.stream.cw:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2783
				_go_fuzz_dep_.CoverTab[74205]++
														err = rws.stream.closeErr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2784
				// _ = "end of CoverTab[74205]"
			default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2785
				_go_fuzz_dep_.CoverTab[74206]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2785
				// _ = "end of CoverTab[74206]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2786
			// _ = "end of CoverTab[74204]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2787
			_go_fuzz_dep_.CoverTab[74207]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2787
			// _ = "end of CoverTab[74207]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2787
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2787
		// _ = "end of CoverTab[74203]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2788
	// _ = "end of CoverTab[74198]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2788
	_go_fuzz_dep_.CoverTab[74199]++
											return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2789
	// _ = "end of CoverTab[74199]"
}

func (w *responseWriter) CloseNotify() <-chan bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2792
	_go_fuzz_dep_.CoverTab[74208]++
											rws := w.rws
											if rws == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2794
		_go_fuzz_dep_.CoverTab[74211]++
												panic("CloseNotify called after Handler finished")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2795
		// _ = "end of CoverTab[74211]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2796
		_go_fuzz_dep_.CoverTab[74212]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2796
		// _ = "end of CoverTab[74212]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2796
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2796
	// _ = "end of CoverTab[74208]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2796
	_go_fuzz_dep_.CoverTab[74209]++
											rws.closeNotifierMu.Lock()
											ch := rws.closeNotifierCh
											if ch == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2799
		_go_fuzz_dep_.CoverTab[74213]++
												ch = make(chan bool, 1)
												rws.closeNotifierCh = ch
												cw := rws.stream.cw
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2802
		_curRoutineNum62_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2802
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum62_)
												go func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2803
			_go_fuzz_dep_.CoverTab[74214]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2803
			defer func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2803
				_go_fuzz_dep_.CoverTab[74215]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2803
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum62_)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2803
				// _ = "end of CoverTab[74215]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2803
			}()
													cw.Wait()
													ch <- true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2805
			// _ = "end of CoverTab[74214]"
		}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2806
		// _ = "end of CoverTab[74213]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2807
		_go_fuzz_dep_.CoverTab[74216]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2807
		// _ = "end of CoverTab[74216]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2807
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2807
	// _ = "end of CoverTab[74209]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2807
	_go_fuzz_dep_.CoverTab[74210]++
											rws.closeNotifierMu.Unlock()
											return ch
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2809
	// _ = "end of CoverTab[74210]"
}

func (w *responseWriter) Header() http.Header {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2812
	_go_fuzz_dep_.CoverTab[74217]++
											rws := w.rws
											if rws == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2814
		_go_fuzz_dep_.CoverTab[74220]++
												panic("Header called after Handler finished")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2815
		// _ = "end of CoverTab[74220]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2816
		_go_fuzz_dep_.CoverTab[74221]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2816
		// _ = "end of CoverTab[74221]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2816
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2816
	// _ = "end of CoverTab[74217]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2816
	_go_fuzz_dep_.CoverTab[74218]++
											if rws.handlerHeader == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2817
		_go_fuzz_dep_.CoverTab[74222]++
												rws.handlerHeader = make(http.Header)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2818
		// _ = "end of CoverTab[74222]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2819
		_go_fuzz_dep_.CoverTab[74223]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2819
		// _ = "end of CoverTab[74223]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2819
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2819
	// _ = "end of CoverTab[74218]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2819
	_go_fuzz_dep_.CoverTab[74219]++
											return rws.handlerHeader
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2820
	// _ = "end of CoverTab[74219]"
}

// checkWriteHeaderCode is a copy of net/http's checkWriteHeaderCode.
func checkWriteHeaderCode(code int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2824
	_go_fuzz_dep_.CoverTab[74224]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2835
	if code < 100 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2835
		_go_fuzz_dep_.CoverTab[74225]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2835
		return code > 999
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2835
		// _ = "end of CoverTab[74225]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2835
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2835
		_go_fuzz_dep_.CoverTab[74226]++
												panic(fmt.Sprintf("invalid WriteHeader code %v", code))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2836
		// _ = "end of CoverTab[74226]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2837
		_go_fuzz_dep_.CoverTab[74227]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2837
		// _ = "end of CoverTab[74227]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2837
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2837
	// _ = "end of CoverTab[74224]"
}

func (w *responseWriter) WriteHeader(code int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2840
	_go_fuzz_dep_.CoverTab[74228]++
											rws := w.rws
											if rws == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2842
		_go_fuzz_dep_.CoverTab[74230]++
												panic("WriteHeader called after Handler finished")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2843
		// _ = "end of CoverTab[74230]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2844
		_go_fuzz_dep_.CoverTab[74231]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2844
		// _ = "end of CoverTab[74231]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2844
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2844
	// _ = "end of CoverTab[74228]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2844
	_go_fuzz_dep_.CoverTab[74229]++
											rws.writeHeader(code)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2845
	// _ = "end of CoverTab[74229]"
}

func (rws *responseWriterState) writeHeader(code int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2848
	_go_fuzz_dep_.CoverTab[74232]++
											if rws.wroteHeader {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2849
		_go_fuzz_dep_.CoverTab[74235]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2850
		// _ = "end of CoverTab[74235]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2851
		_go_fuzz_dep_.CoverTab[74236]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2851
		// _ = "end of CoverTab[74236]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2851
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2851
	// _ = "end of CoverTab[74232]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2851
	_go_fuzz_dep_.CoverTab[74233]++

											checkWriteHeaderCode(code)

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2856
	if code >= 100 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2856
		_go_fuzz_dep_.CoverTab[74237]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2856
		return code <= 199
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2856
		// _ = "end of CoverTab[74237]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2856
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2856
		_go_fuzz_dep_.CoverTab[74238]++

												h := rws.handlerHeader

												_, cl := h["Content-Length"]
												_, te := h["Transfer-Encoding"]
												if cl || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2862
			_go_fuzz_dep_.CoverTab[74241]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2862
			return te
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2862
			// _ = "end of CoverTab[74241]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2862
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2862
			_go_fuzz_dep_.CoverTab[74242]++
													h = h.Clone()
													h.Del("Content-Length")
													h.Del("Transfer-Encoding")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2865
			// _ = "end of CoverTab[74242]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2866
			_go_fuzz_dep_.CoverTab[74243]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2866
			// _ = "end of CoverTab[74243]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2866
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2866
		// _ = "end of CoverTab[74238]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2866
		_go_fuzz_dep_.CoverTab[74239]++

												if rws.conn.writeHeaders(rws.stream, &writeResHeaders{
			streamID:	rws.stream.id,
			httpResCode:	code,
			h:		h,
			endStream: rws.handlerDone && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2872
				_go_fuzz_dep_.CoverTab[74244]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2872
				return !rws.hasTrailers()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2872
				// _ = "end of CoverTab[74244]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2872
			}(),
		}) != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2873
			_go_fuzz_dep_.CoverTab[74245]++
													rws.dirty = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2874
			// _ = "end of CoverTab[74245]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2875
			_go_fuzz_dep_.CoverTab[74246]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2875
			// _ = "end of CoverTab[74246]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2875
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2875
		// _ = "end of CoverTab[74239]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2875
		_go_fuzz_dep_.CoverTab[74240]++

												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2877
		// _ = "end of CoverTab[74240]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2878
		_go_fuzz_dep_.CoverTab[74247]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2878
		// _ = "end of CoverTab[74247]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2878
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2878
	// _ = "end of CoverTab[74233]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2878
	_go_fuzz_dep_.CoverTab[74234]++

											rws.wroteHeader = true
											rws.status = code
											if len(rws.handlerHeader) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2882
		_go_fuzz_dep_.CoverTab[74248]++
												rws.snapHeader = cloneHeader(rws.handlerHeader)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2883
		// _ = "end of CoverTab[74248]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2884
		_go_fuzz_dep_.CoverTab[74249]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2884
		// _ = "end of CoverTab[74249]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2884
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2884
	// _ = "end of CoverTab[74234]"
}

func cloneHeader(h http.Header) http.Header {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2887
	_go_fuzz_dep_.CoverTab[74250]++
											h2 := make(http.Header, len(h))
											for k, vv := range h {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2889
		_go_fuzz_dep_.CoverTab[74252]++
												vv2 := make([]string, len(vv))
												copy(vv2, vv)
												h2[k] = vv2
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2892
		// _ = "end of CoverTab[74252]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2893
	// _ = "end of CoverTab[74250]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2893
	_go_fuzz_dep_.CoverTab[74251]++
											return h2
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2894
	// _ = "end of CoverTab[74251]"
}

// The Life Of A Write is like this:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2897
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2897
// * Handler calls w.Write or w.WriteString ->
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2897
// * -> rws.bw (*bufio.Writer) ->
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2897
// * (Handler might call Flush)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2897
// * -> chunkWriter{rws}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2897
// * -> responseWriterState.writeChunk(p []byte)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2897
// * -> responseWriterState.writeChunk (most of the magic; see comment there)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2905
func (w *responseWriter) Write(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2905
	_go_fuzz_dep_.CoverTab[74253]++
											return w.write(len(p), p, "")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2906
	// _ = "end of CoverTab[74253]"
}

func (w *responseWriter) WriteString(s string) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2909
	_go_fuzz_dep_.CoverTab[74254]++
											return w.write(len(s), nil, s)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2910
	// _ = "end of CoverTab[74254]"
}

// either dataB or dataS is non-zero.
func (w *responseWriter) write(lenData int, dataB []byte, dataS string) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2914
	_go_fuzz_dep_.CoverTab[74255]++
											rws := w.rws
											if rws == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2916
		_go_fuzz_dep_.CoverTab[74260]++
												panic("Write called after Handler finished")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2917
		// _ = "end of CoverTab[74260]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2918
		_go_fuzz_dep_.CoverTab[74261]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2918
		// _ = "end of CoverTab[74261]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2918
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2918
	// _ = "end of CoverTab[74255]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2918
	_go_fuzz_dep_.CoverTab[74256]++
											if !rws.wroteHeader {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2919
		_go_fuzz_dep_.CoverTab[74262]++
												w.WriteHeader(200)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2920
		// _ = "end of CoverTab[74262]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2921
		_go_fuzz_dep_.CoverTab[74263]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2921
		// _ = "end of CoverTab[74263]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2921
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2921
	// _ = "end of CoverTab[74256]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2921
	_go_fuzz_dep_.CoverTab[74257]++
											if !bodyAllowedForStatus(rws.status) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2922
		_go_fuzz_dep_.CoverTab[74264]++
												return 0, http.ErrBodyNotAllowed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2923
		// _ = "end of CoverTab[74264]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2924
		_go_fuzz_dep_.CoverTab[74265]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2924
		// _ = "end of CoverTab[74265]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2924
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2924
	// _ = "end of CoverTab[74257]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2924
	_go_fuzz_dep_.CoverTab[74258]++
											rws.wroteBytes += int64(len(dataB)) + int64(len(dataS))
											if rws.sentContentLen != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2926
		_go_fuzz_dep_.CoverTab[74266]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2926
		return rws.wroteBytes > rws.sentContentLen
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2926
		// _ = "end of CoverTab[74266]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2926
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2926
		_go_fuzz_dep_.CoverTab[74267]++

												return 0, errors.New("http2: handler wrote more than declared Content-Length")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2928
		// _ = "end of CoverTab[74267]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2929
		_go_fuzz_dep_.CoverTab[74268]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2929
		// _ = "end of CoverTab[74268]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2929
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2929
	// _ = "end of CoverTab[74258]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2929
	_go_fuzz_dep_.CoverTab[74259]++

											if dataB != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2931
		_go_fuzz_dep_.CoverTab[74269]++
												return rws.bw.Write(dataB)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2932
		// _ = "end of CoverTab[74269]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2933
		_go_fuzz_dep_.CoverTab[74270]++
												return rws.bw.WriteString(dataS)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2934
		// _ = "end of CoverTab[74270]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2935
	// _ = "end of CoverTab[74259]"
}

func (w *responseWriter) handlerDone() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2938
	_go_fuzz_dep_.CoverTab[74271]++
											rws := w.rws
											dirty := rws.dirty
											rws.handlerDone = true
											w.Flush()
											w.rws = nil
											if !dirty {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2944
		_go_fuzz_dep_.CoverTab[74272]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2951
		responseWriterStatePool.Put(rws)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2951
		// _ = "end of CoverTab[74272]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2952
		_go_fuzz_dep_.CoverTab[74273]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2952
		// _ = "end of CoverTab[74273]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2952
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2952
	// _ = "end of CoverTab[74271]"
}

// Push errors.
var (
	ErrRecursivePush	= errors.New("http2: recursive push not allowed")
	ErrPushLimitReached	= errors.New("http2: push would exceed peer's SETTINGS_MAX_CONCURRENT_STREAMS")
)

var _ http.Pusher = (*responseWriter)(nil)

func (w *responseWriter) Push(target string, opts *http.PushOptions) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2963
	_go_fuzz_dep_.CoverTab[74274]++
											st := w.rws.stream
											sc := st.sc
											sc.serveG.checkNotOn()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2970
	if st.isPushed() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2970
		_go_fuzz_dep_.CoverTab[74286]++
												return ErrRecursivePush
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2971
		// _ = "end of CoverTab[74286]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2972
		_go_fuzz_dep_.CoverTab[74287]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2972
		// _ = "end of CoverTab[74287]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2972
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2972
	// _ = "end of CoverTab[74274]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2972
	_go_fuzz_dep_.CoverTab[74275]++

											if opts == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2974
		_go_fuzz_dep_.CoverTab[74288]++
												opts = new(http.PushOptions)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2975
		// _ = "end of CoverTab[74288]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2976
		_go_fuzz_dep_.CoverTab[74289]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2976
		// _ = "end of CoverTab[74289]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2976
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2976
	// _ = "end of CoverTab[74275]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2976
	_go_fuzz_dep_.CoverTab[74276]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2979
	if opts.Method == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2979
		_go_fuzz_dep_.CoverTab[74290]++
												opts.Method = "GET"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2980
		// _ = "end of CoverTab[74290]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2981
		_go_fuzz_dep_.CoverTab[74291]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2981
		// _ = "end of CoverTab[74291]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2981
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2981
	// _ = "end of CoverTab[74276]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2981
	_go_fuzz_dep_.CoverTab[74277]++
											if opts.Header == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2982
		_go_fuzz_dep_.CoverTab[74292]++
												opts.Header = http.Header{}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2983
		// _ = "end of CoverTab[74292]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2984
		_go_fuzz_dep_.CoverTab[74293]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2984
		// _ = "end of CoverTab[74293]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2984
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2984
	// _ = "end of CoverTab[74277]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2984
	_go_fuzz_dep_.CoverTab[74278]++
											wantScheme := "http"
											if w.rws.req.TLS != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2986
		_go_fuzz_dep_.CoverTab[74294]++
												wantScheme = "https"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2987
		// _ = "end of CoverTab[74294]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2988
		_go_fuzz_dep_.CoverTab[74295]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2988
		// _ = "end of CoverTab[74295]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2988
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2988
	// _ = "end of CoverTab[74278]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2988
	_go_fuzz_dep_.CoverTab[74279]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2991
	u, err := url.Parse(target)
	if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2992
		_go_fuzz_dep_.CoverTab[74296]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2993
		// _ = "end of CoverTab[74296]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2994
		_go_fuzz_dep_.CoverTab[74297]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2994
		// _ = "end of CoverTab[74297]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2994
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2994
	// _ = "end of CoverTab[74279]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2994
	_go_fuzz_dep_.CoverTab[74280]++
											if u.Scheme == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2995
		_go_fuzz_dep_.CoverTab[74298]++
												if !strings.HasPrefix(target, "/") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2996
			_go_fuzz_dep_.CoverTab[74300]++
													return fmt.Errorf("target must be an absolute URL or an absolute path: %q", target)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2997
			// _ = "end of CoverTab[74300]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2998
			_go_fuzz_dep_.CoverTab[74301]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2998
			// _ = "end of CoverTab[74301]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2998
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2998
		// _ = "end of CoverTab[74298]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:2998
		_go_fuzz_dep_.CoverTab[74299]++
												u.Scheme = wantScheme
												u.Host = w.rws.req.Host
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3000
		// _ = "end of CoverTab[74299]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3001
		_go_fuzz_dep_.CoverTab[74302]++
												if u.Scheme != wantScheme {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3002
			_go_fuzz_dep_.CoverTab[74304]++
													return fmt.Errorf("cannot push URL with scheme %q from request with scheme %q", u.Scheme, wantScheme)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3003
			// _ = "end of CoverTab[74304]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3004
			_go_fuzz_dep_.CoverTab[74305]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3004
			// _ = "end of CoverTab[74305]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3004
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3004
		// _ = "end of CoverTab[74302]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3004
		_go_fuzz_dep_.CoverTab[74303]++
												if u.Host == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3005
			_go_fuzz_dep_.CoverTab[74306]++
													return errors.New("URL must have a host")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3006
			// _ = "end of CoverTab[74306]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3007
			_go_fuzz_dep_.CoverTab[74307]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3007
			// _ = "end of CoverTab[74307]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3007
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3007
		// _ = "end of CoverTab[74303]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3008
	// _ = "end of CoverTab[74280]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3008
	_go_fuzz_dep_.CoverTab[74281]++
											for k := range opts.Header {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3009
		_go_fuzz_dep_.CoverTab[74308]++
												if strings.HasPrefix(k, ":") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3010
			_go_fuzz_dep_.CoverTab[74310]++
													return fmt.Errorf("promised request headers cannot include pseudo header %q", k)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3011
			// _ = "end of CoverTab[74310]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3012
			_go_fuzz_dep_.CoverTab[74311]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3012
			// _ = "end of CoverTab[74311]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3012
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3012
		// _ = "end of CoverTab[74308]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3012
		_go_fuzz_dep_.CoverTab[74309]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3017
		if asciiEqualFold(k, "content-length") || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3017
			_go_fuzz_dep_.CoverTab[74312]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3017
			return asciiEqualFold(k, "content-encoding")
													// _ = "end of CoverTab[74312]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3018
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3018
			_go_fuzz_dep_.CoverTab[74313]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3018
			return asciiEqualFold(k, "trailer")
													// _ = "end of CoverTab[74313]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3019
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3019
			_go_fuzz_dep_.CoverTab[74314]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3019
			return asciiEqualFold(k, "te")
													// _ = "end of CoverTab[74314]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3020
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3020
			_go_fuzz_dep_.CoverTab[74315]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3020
			return asciiEqualFold(k, "expect")
													// _ = "end of CoverTab[74315]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3021
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3021
			_go_fuzz_dep_.CoverTab[74316]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3021
			return asciiEqualFold(k, "host")
													// _ = "end of CoverTab[74316]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3022
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3022
			_go_fuzz_dep_.CoverTab[74317]++
													return fmt.Errorf("promised request headers cannot include %q", k)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3023
			// _ = "end of CoverTab[74317]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3024
			_go_fuzz_dep_.CoverTab[74318]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3024
			// _ = "end of CoverTab[74318]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3024
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3024
		// _ = "end of CoverTab[74309]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3025
	// _ = "end of CoverTab[74281]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3025
	_go_fuzz_dep_.CoverTab[74282]++
											if err := checkValidHTTP2RequestHeaders(opts.Header); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3026
		_go_fuzz_dep_.CoverTab[74319]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3027
		// _ = "end of CoverTab[74319]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3028
		_go_fuzz_dep_.CoverTab[74320]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3028
		// _ = "end of CoverTab[74320]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3028
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3028
	// _ = "end of CoverTab[74282]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3028
	_go_fuzz_dep_.CoverTab[74283]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3033
	if opts.Method != "GET" && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3033
		_go_fuzz_dep_.CoverTab[74321]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3033
		return opts.Method != "HEAD"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3033
		// _ = "end of CoverTab[74321]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3033
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3033
		_go_fuzz_dep_.CoverTab[74322]++
												return fmt.Errorf("method %q must be GET or HEAD", opts.Method)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3034
		// _ = "end of CoverTab[74322]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3035
		_go_fuzz_dep_.CoverTab[74323]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3035
		// _ = "end of CoverTab[74323]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3035
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3035
	// _ = "end of CoverTab[74283]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3035
	_go_fuzz_dep_.CoverTab[74284]++

											msg := &startPushRequest{
		parent:	st,
		method:	opts.Method,
		url:	u,
		header:	cloneHeader(opts.Header),
		done:	errChanPool.Get().(chan error),
	}

	select {
	case <-sc.doneServing:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3046
		_go_fuzz_dep_.CoverTab[74324]++
												return errClientDisconnected
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3047
		// _ = "end of CoverTab[74324]"
	case <-st.cw:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3048
		_go_fuzz_dep_.CoverTab[74325]++
												return errStreamClosed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3049
		// _ = "end of CoverTab[74325]"
	case sc.serveMsgCh <- msg:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3050
		_go_fuzz_dep_.CoverTab[74326]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3050
		// _ = "end of CoverTab[74326]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3051
	// _ = "end of CoverTab[74284]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3051
	_go_fuzz_dep_.CoverTab[74285]++

											select {
	case <-sc.doneServing:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3054
		_go_fuzz_dep_.CoverTab[74327]++
												return errClientDisconnected
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3055
		// _ = "end of CoverTab[74327]"
	case <-st.cw:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3056
		_go_fuzz_dep_.CoverTab[74328]++
												return errStreamClosed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3057
		// _ = "end of CoverTab[74328]"
	case err := <-msg.done:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3058
		_go_fuzz_dep_.CoverTab[74329]++
												errChanPool.Put(msg.done)
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3060
		// _ = "end of CoverTab[74329]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3061
	// _ = "end of CoverTab[74285]"
}

type startPushRequest struct {
	parent	*stream
	method	string
	url	*url.URL
	header	http.Header
	done	chan error
}

func (sc *serverConn) startPush(msg *startPushRequest) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3072
	_go_fuzz_dep_.CoverTab[74330]++
											sc.serveG.check()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3078
	if msg.parent.state != stateOpen && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3078
		_go_fuzz_dep_.CoverTab[74334]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3078
		return msg.parent.state != stateHalfClosedRemote
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3078
		// _ = "end of CoverTab[74334]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3078
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3078
		_go_fuzz_dep_.CoverTab[74335]++

												msg.done <- errStreamClosed
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3081
		// _ = "end of CoverTab[74335]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3082
		_go_fuzz_dep_.CoverTab[74336]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3082
		// _ = "end of CoverTab[74336]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3082
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3082
	// _ = "end of CoverTab[74330]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3082
	_go_fuzz_dep_.CoverTab[74331]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3085
	if !sc.pushEnabled {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3085
		_go_fuzz_dep_.CoverTab[74337]++
												msg.done <- http.ErrNotSupported
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3087
		// _ = "end of CoverTab[74337]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3088
		_go_fuzz_dep_.CoverTab[74338]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3088
		// _ = "end of CoverTab[74338]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3088
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3088
	// _ = "end of CoverTab[74331]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3088
	_go_fuzz_dep_.CoverTab[74332]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3093
	allocatePromisedID := func() (uint32, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3093
		_go_fuzz_dep_.CoverTab[74339]++
												sc.serveG.check()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3098
		if !sc.pushEnabled {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3098
			_go_fuzz_dep_.CoverTab[74344]++
													return 0, http.ErrNotSupported
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3099
			// _ = "end of CoverTab[74344]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3100
			_go_fuzz_dep_.CoverTab[74345]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3100
			// _ = "end of CoverTab[74345]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3100
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3100
		// _ = "end of CoverTab[74339]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3100
		_go_fuzz_dep_.CoverTab[74340]++

												if sc.curPushedStreams+1 > sc.clientMaxStreams {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3102
			_go_fuzz_dep_.CoverTab[74346]++
													return 0, ErrPushLimitReached
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3103
			// _ = "end of CoverTab[74346]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3104
			_go_fuzz_dep_.CoverTab[74347]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3104
			// _ = "end of CoverTab[74347]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3104
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3104
		// _ = "end of CoverTab[74340]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3104
		_go_fuzz_dep_.CoverTab[74341]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3110
		if sc.maxPushPromiseID+2 >= 1<<31 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3110
			_go_fuzz_dep_.CoverTab[74348]++
													sc.startGracefulShutdownInternal()
													return 0, ErrPushLimitReached
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3112
			// _ = "end of CoverTab[74348]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3113
			_go_fuzz_dep_.CoverTab[74349]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3113
			// _ = "end of CoverTab[74349]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3113
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3113
		// _ = "end of CoverTab[74341]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3113
		_go_fuzz_dep_.CoverTab[74342]++
												sc.maxPushPromiseID += 2
												promisedID := sc.maxPushPromiseID

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3122
		promised := sc.newStream(promisedID, msg.parent.id, stateHalfClosedRemote)
		rw, req, err := sc.newWriterAndRequestNoBody(promised, requestParam{
			method:		msg.method,
			scheme:		msg.url.Scheme,
			authority:	msg.url.Host,
			path:		msg.url.RequestURI(),
			header:		cloneHeader(msg.header),
		})
		if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3130
			_go_fuzz_dep_.CoverTab[74350]++

													panic(fmt.Sprintf("newWriterAndRequestNoBody(%+v): %v", msg.url, err))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3132
			// _ = "end of CoverTab[74350]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3133
			_go_fuzz_dep_.CoverTab[74351]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3133
			// _ = "end of CoverTab[74351]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3133
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3133
		// _ = "end of CoverTab[74342]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3133
		_go_fuzz_dep_.CoverTab[74343]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3133
		_curRoutineNum63_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3133
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum63_)

												go sc.runHandler(rw, req, sc.handler.ServeHTTP)
												return promisedID, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3136
		// _ = "end of CoverTab[74343]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3137
	// _ = "end of CoverTab[74332]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3137
	_go_fuzz_dep_.CoverTab[74333]++

											sc.writeFrame(FrameWriteRequest{
		write: &writePushPromise{
			streamID:		msg.parent.id,
			method:			msg.method,
			url:			msg.url,
			h:			msg.header,
			allocatePromisedID:	allocatePromisedID,
		},
		stream:	msg.parent,
		done:	msg.done,
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3149
	// _ = "end of CoverTab[74333]"
}

// foreachHeaderElement splits v according to the "#rule" construction
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3152
// in RFC 7230 section 7 and calls fn for each non-empty element.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3154
func foreachHeaderElement(v string, fn func(string)) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3154
	_go_fuzz_dep_.CoverTab[74352]++
											v = textproto.TrimString(v)
											if v == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3156
		_go_fuzz_dep_.CoverTab[74355]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3157
		// _ = "end of CoverTab[74355]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3158
		_go_fuzz_dep_.CoverTab[74356]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3158
		// _ = "end of CoverTab[74356]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3158
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3158
	// _ = "end of CoverTab[74352]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3158
	_go_fuzz_dep_.CoverTab[74353]++
											if !strings.Contains(v, ",") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3159
		_go_fuzz_dep_.CoverTab[74357]++
												fn(v)
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3161
		// _ = "end of CoverTab[74357]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3162
		_go_fuzz_dep_.CoverTab[74358]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3162
		// _ = "end of CoverTab[74358]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3162
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3162
	// _ = "end of CoverTab[74353]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3162
	_go_fuzz_dep_.CoverTab[74354]++
											for _, f := range strings.Split(v, ",") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3163
		_go_fuzz_dep_.CoverTab[74359]++
												if f = textproto.TrimString(f); f != "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3164
			_go_fuzz_dep_.CoverTab[74360]++
													fn(f)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3165
			// _ = "end of CoverTab[74360]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3166
			_go_fuzz_dep_.CoverTab[74361]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3166
			// _ = "end of CoverTab[74361]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3166
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3166
		// _ = "end of CoverTab[74359]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3167
	// _ = "end of CoverTab[74354]"
}

// From http://httpwg.org/specs/rfc7540.html#rfc.section.8.1.2.2
var connHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Proxy-Connection",
	"Transfer-Encoding",
	"Upgrade",
}

// checkValidHTTP2RequestHeaders checks whether h is a valid HTTP/2 request,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3179
// per RFC 7540 Section 8.1.2.2.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3179
// The returned error is reported to users.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3182
func checkValidHTTP2RequestHeaders(h http.Header) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3182
	_go_fuzz_dep_.CoverTab[74362]++
											for _, k := range connHeaders {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3183
		_go_fuzz_dep_.CoverTab[74365]++
												if _, ok := h[k]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3184
			_go_fuzz_dep_.CoverTab[74366]++
													return fmt.Errorf("request header %q is not valid in HTTP/2", k)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3185
			// _ = "end of CoverTab[74366]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3186
			_go_fuzz_dep_.CoverTab[74367]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3186
			// _ = "end of CoverTab[74367]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3186
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3186
		// _ = "end of CoverTab[74365]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3187
	// _ = "end of CoverTab[74362]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3187
	_go_fuzz_dep_.CoverTab[74363]++
											te := h["Te"]
											if len(te) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3189
		_go_fuzz_dep_.CoverTab[74368]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3189
		return (len(te) > 1 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3189
			_go_fuzz_dep_.CoverTab[74369]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3189
			return (te[0] != "trailers" && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3189
				_go_fuzz_dep_.CoverTab[74370]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3189
				return te[0] != ""
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3189
				// _ = "end of CoverTab[74370]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3189
			}())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3189
			// _ = "end of CoverTab[74369]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3189
		}())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3189
		// _ = "end of CoverTab[74368]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3189
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3189
		_go_fuzz_dep_.CoverTab[74371]++
												return errors.New(`request header "TE" may only be "trailers" in HTTP/2`)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3190
		// _ = "end of CoverTab[74371]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3191
		_go_fuzz_dep_.CoverTab[74372]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3191
		// _ = "end of CoverTab[74372]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3191
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3191
	// _ = "end of CoverTab[74363]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3191
	_go_fuzz_dep_.CoverTab[74364]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3192
	// _ = "end of CoverTab[74364]"
}

func new400Handler(err error) http.HandlerFunc {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3195
	_go_fuzz_dep_.CoverTab[74373]++
											return func(w http.ResponseWriter, r *http.Request) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3196
		_go_fuzz_dep_.CoverTab[74374]++
												http.Error(w, err.Error(), http.StatusBadRequest)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3197
		// _ = "end of CoverTab[74374]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3198
	// _ = "end of CoverTab[74373]"
}

// h1ServerKeepAlivesDisabled reports whether hs has its keep-alives
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3201
// disabled. See comments on h1ServerShutdownChan above for why
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3201
// the code is written this way.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3204
func h1ServerKeepAlivesDisabled(hs *http.Server) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3204
	_go_fuzz_dep_.CoverTab[74375]++
											var x interface{} = hs
											type I interface {
		doKeepAlives() bool
	}
	if hs, ok := x.(I); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3209
		_go_fuzz_dep_.CoverTab[74377]++
												return !hs.doKeepAlives()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3210
		// _ = "end of CoverTab[74377]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3211
		_go_fuzz_dep_.CoverTab[74378]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3211
		// _ = "end of CoverTab[74378]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3211
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3211
	// _ = "end of CoverTab[74375]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3211
	_go_fuzz_dep_.CoverTab[74376]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3212
	// _ = "end of CoverTab[74376]"
}

func (sc *serverConn) countError(name string, err error) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3215
	_go_fuzz_dep_.CoverTab[74379]++
											if sc == nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3216
		_go_fuzz_dep_.CoverTab[74384]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3216
		return sc.srv == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3216
		// _ = "end of CoverTab[74384]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3216
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3216
		_go_fuzz_dep_.CoverTab[74385]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3217
		// _ = "end of CoverTab[74385]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3218
		_go_fuzz_dep_.CoverTab[74386]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3218
		// _ = "end of CoverTab[74386]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3218
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3218
	// _ = "end of CoverTab[74379]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3218
	_go_fuzz_dep_.CoverTab[74380]++
											f := sc.srv.CountError
											if f == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3220
		_go_fuzz_dep_.CoverTab[74387]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3221
		// _ = "end of CoverTab[74387]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3222
		_go_fuzz_dep_.CoverTab[74388]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3222
		// _ = "end of CoverTab[74388]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3222
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3222
	// _ = "end of CoverTab[74380]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3222
	_go_fuzz_dep_.CoverTab[74381]++
											var typ string
											var code ErrCode
											switch e := err.(type) {
	case ConnectionError:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3226
		_go_fuzz_dep_.CoverTab[74389]++
												typ = "conn"
												code = ErrCode(e)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3228
		// _ = "end of CoverTab[74389]"
	case StreamError:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3229
		_go_fuzz_dep_.CoverTab[74390]++
												typ = "stream"
												code = ErrCode(e.Code)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3231
		// _ = "end of CoverTab[74390]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3232
		_go_fuzz_dep_.CoverTab[74391]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3233
		// _ = "end of CoverTab[74391]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3234
	// _ = "end of CoverTab[74381]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3234
	_go_fuzz_dep_.CoverTab[74382]++
											codeStr := errCodeName[code]
											if codeStr == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3236
		_go_fuzz_dep_.CoverTab[74392]++
												codeStr = strconv.Itoa(int(code))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3237
		// _ = "end of CoverTab[74392]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3238
		_go_fuzz_dep_.CoverTab[74393]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3238
		// _ = "end of CoverTab[74393]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3238
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3238
	// _ = "end of CoverTab[74382]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3238
	_go_fuzz_dep_.CoverTab[74383]++
											f(fmt.Sprintf("%s_%s_%s", typ, codeStr, name))
											return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3240
	// _ = "end of CoverTab[74383]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3241
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/server.go:3241
var _ = _go_fuzz_dep_.CoverTab
