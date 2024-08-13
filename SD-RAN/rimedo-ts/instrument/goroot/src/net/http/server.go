// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP server. See RFC 7230 through 7235.

//line /usr/local/go/src/net/http/server.go:7
package http

//line /usr/local/go/src/net/http/server.go:7
import (
//line /usr/local/go/src/net/http/server.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/server.go:7
)
//line /usr/local/go/src/net/http/server.go:7
import (
//line /usr/local/go/src/net/http/server.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/server.go:7
)

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"internal/godebug"
	"io"
	"log"
	"math/rand"
	"net"
	"net/textproto"
	"net/url"
	urlpkg "net/url"
	"path"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/http/httpguts"
)

// Errors used by the HTTP server.
var (
	// ErrBodyNotAllowed is returned by ResponseWriter.Write calls
	// when the HTTP method or response code does not permit a
	// body.
	ErrBodyNotAllowed	= errors.New("http: request method or response status code does not allow body")

	// ErrHijacked is returned by ResponseWriter.Write calls when
	// the underlying connection has been hijacked using the
	// Hijacker interface. A zero-byte write on a hijacked
	// connection will return ErrHijacked without any other side
	// effects.
	ErrHijacked	= errors.New("http: connection has been hijacked")

	// ErrContentLength is returned by ResponseWriter.Write calls
	// when a Handler set a Content-Length response header with a
	// declared size and then attempted to write more bytes than
	// declared.
	ErrContentLength	= errors.New("http: wrote more than the declared Content-Length")

	// Deprecated: ErrWriteAfterFlush is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	ErrWriteAfterFlush	= errors.New("unused")
)

// A Handler responds to an HTTP request.
//line /usr/local/go/src/net/http/server.go:62
//
//line /usr/local/go/src/net/http/server.go:62
// ServeHTTP should write reply headers and data to the ResponseWriter
//line /usr/local/go/src/net/http/server.go:62
// and then return. Returning signals that the request is finished; it
//line /usr/local/go/src/net/http/server.go:62
// is not valid to use the ResponseWriter or read from the
//line /usr/local/go/src/net/http/server.go:62
// Request.Body after or concurrently with the completion of the
//line /usr/local/go/src/net/http/server.go:62
// ServeHTTP call.
//line /usr/local/go/src/net/http/server.go:62
//
//line /usr/local/go/src/net/http/server.go:62
// Depending on the HTTP client software, HTTP protocol version, and
//line /usr/local/go/src/net/http/server.go:62
// any intermediaries between the client and the Go server, it may not
//line /usr/local/go/src/net/http/server.go:62
// be possible to read from the Request.Body after writing to the
//line /usr/local/go/src/net/http/server.go:62
// ResponseWriter. Cautious handlers should read the Request.Body
//line /usr/local/go/src/net/http/server.go:62
// first, and then reply.
//line /usr/local/go/src/net/http/server.go:62
//
//line /usr/local/go/src/net/http/server.go:62
// Except for reading the body, handlers should not modify the
//line /usr/local/go/src/net/http/server.go:62
// provided Request.
//line /usr/local/go/src/net/http/server.go:62
//
//line /usr/local/go/src/net/http/server.go:62
// If ServeHTTP panics, the server (the caller of ServeHTTP) assumes
//line /usr/local/go/src/net/http/server.go:62
// that the effect of the panic was isolated to the active request.
//line /usr/local/go/src/net/http/server.go:62
// It recovers the panic, logs a stack trace to the server error log,
//line /usr/local/go/src/net/http/server.go:62
// and either closes the network connection or sends an HTTP/2
//line /usr/local/go/src/net/http/server.go:62
// RST_STREAM, depending on the HTTP protocol. To abort a handler so
//line /usr/local/go/src/net/http/server.go:62
// the client sees an interrupted response but the server doesn't log
//line /usr/local/go/src/net/http/server.go:62
// an error, panic with the value ErrAbortHandler.
//line /usr/local/go/src/net/http/server.go:86
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

// A ResponseWriter interface is used by an HTTP handler to
//line /usr/local/go/src/net/http/server.go:90
// construct an HTTP response.
//line /usr/local/go/src/net/http/server.go:90
//
//line /usr/local/go/src/net/http/server.go:90
// A ResponseWriter may not be used after the Handler.ServeHTTP method
//line /usr/local/go/src/net/http/server.go:90
// has returned.
//line /usr/local/go/src/net/http/server.go:95
type ResponseWriter interface {
	// Header returns the header map that will be sent by
	// WriteHeader. The Header map also is the mechanism with which
	// Handlers can set HTTP trailers.
	//
	// Changing the header map after a call to WriteHeader (or
	// Write) has no effect unless the HTTP status code was of the
	// 1xx class or the modified headers are trailers.
	//
	// There are two ways to set Trailers. The preferred way is to
	// predeclare in the headers which trailers you will later
	// send by setting the "Trailer" header to the names of the
	// trailer keys which will come later. In this case, those
	// keys of the Header map are treated as if they were
	// trailers. See the example. The second way, for trailer
	// keys not known to the Handler until after the first Write,
	// is to prefix the Header map keys with the TrailerPrefix
	// constant value. See TrailerPrefix.
	//
	// To suppress automatic response headers (such as "Date"), set
	// their value to nil.
	Header() Header

	// Write writes the data to the connection as part of an HTTP reply.
	//
	// If WriteHeader has not yet been called, Write calls
	// WriteHeader(http.StatusOK) before writing the data. If the Header
	// does not contain a Content-Type line, Write adds a Content-Type set
	// to the result of passing the initial 512 bytes of written data to
	// DetectContentType. Additionally, if the total size of all written
	// data is under a few KB and there are no Flush calls, the
	// Content-Length header is added automatically.
	//
	// Depending on the HTTP protocol version and the client, calling
	// Write or WriteHeader may prevent future reads on the
	// Request.Body. For HTTP/1.x requests, handlers should read any
	// needed request body data before writing the response. Once the
	// headers have been flushed (due to either an explicit Flusher.Flush
	// call or writing enough data to trigger a flush), the request body
	// may be unavailable. For HTTP/2 requests, the Go HTTP server permits
	// handlers to continue to read the request body while concurrently
	// writing the response. However, such behavior may not be supported
	// by all HTTP/2 clients. Handlers should read before writing if
	// possible to maximize compatibility.
	Write([]byte) (int, error)

	// WriteHeader sends an HTTP response header with the provided
	// status code.
	//
	// If WriteHeader is not called explicitly, the first call to Write
	// will trigger an implicit WriteHeader(http.StatusOK).
	// Thus explicit calls to WriteHeader are mainly used to
	// send error codes or 1xx informational responses.
	//
	// The provided code must be a valid HTTP 1xx-5xx status code.
	// Any number of 1xx headers may be written, followed by at most
	// one 2xx-5xx header. 1xx headers are sent immediately, but 2xx-5xx
	// headers may be buffered. Use the Flusher interface to send
	// buffered data. The header map is cleared when 2xx-5xx headers are
	// sent, but not with 1xx headers.
	//
	// The server will automatically send a 100 (Continue) header
	// on the first read from the request body if the request has
	// an "Expect: 100-continue" header.
	WriteHeader(statusCode int)
}

// The Flusher interface is implemented by ResponseWriters that allow
//line /usr/local/go/src/net/http/server.go:162
// an HTTP handler to flush buffered data to the client.
//line /usr/local/go/src/net/http/server.go:162
//
//line /usr/local/go/src/net/http/server.go:162
// The default HTTP/1.x and HTTP/2 ResponseWriter implementations
//line /usr/local/go/src/net/http/server.go:162
// support Flusher, but ResponseWriter wrappers may not. Handlers
//line /usr/local/go/src/net/http/server.go:162
// should always test for this ability at runtime.
//line /usr/local/go/src/net/http/server.go:162
//
//line /usr/local/go/src/net/http/server.go:162
// Note that even for ResponseWriters that support Flush,
//line /usr/local/go/src/net/http/server.go:162
// if the client is connected through an HTTP proxy,
//line /usr/local/go/src/net/http/server.go:162
// the buffered data may not reach the client until the response
//line /usr/local/go/src/net/http/server.go:162
// completes.
//line /usr/local/go/src/net/http/server.go:173
type Flusher interface {
	// Flush sends any buffered data to the client.
	Flush()
}

// The Hijacker interface is implemented by ResponseWriters that allow
//line /usr/local/go/src/net/http/server.go:178
// an HTTP handler to take over the connection.
//line /usr/local/go/src/net/http/server.go:178
//
//line /usr/local/go/src/net/http/server.go:178
// The default ResponseWriter for HTTP/1.x connections supports
//line /usr/local/go/src/net/http/server.go:178
// Hijacker, but HTTP/2 connections intentionally do not.
//line /usr/local/go/src/net/http/server.go:178
// ResponseWriter wrappers may also not support Hijacker. Handlers
//line /usr/local/go/src/net/http/server.go:178
// should always test for this ability at runtime.
//line /usr/local/go/src/net/http/server.go:185
type Hijacker interface {
	// Hijack lets the caller take over the connection.
	// After a call to Hijack the HTTP server library
	// will not do anything else with the connection.
	//
	// It becomes the caller's responsibility to manage
	// and close the connection.
	//
	// The returned net.Conn may have read or write deadlines
	// already set, depending on the configuration of the
	// Server. It is the caller's responsibility to set
	// or clear those deadlines as needed.
	//
	// The returned bufio.Reader may contain unprocessed buffered
	// data from the client.
	//
	// After a call to Hijack, the original Request.Body must not
	// be used. The original Request's Context remains valid and
	// is not canceled until the Request's ServeHTTP method
	// returns.
	Hijack() (net.Conn, *bufio.ReadWriter, error)
}

// The CloseNotifier interface is implemented by ResponseWriters which
//line /usr/local/go/src/net/http/server.go:208
// allow detecting when the underlying connection has gone away.
//line /usr/local/go/src/net/http/server.go:208
//
//line /usr/local/go/src/net/http/server.go:208
// This mechanism can be used to cancel long operations on the server
//line /usr/local/go/src/net/http/server.go:208
// if the client has disconnected before the response is ready.
//line /usr/local/go/src/net/http/server.go:208
//
//line /usr/local/go/src/net/http/server.go:208
// Deprecated: the CloseNotifier interface predates Go's context package.
//line /usr/local/go/src/net/http/server.go:208
// New code should use Request.Context instead.
//line /usr/local/go/src/net/http/server.go:216
type CloseNotifier interface {
	// CloseNotify returns a channel that receives at most a
	// single value (true) when the client connection has gone
	// away.
	//
	// CloseNotify may wait to notify until Request.Body has been
	// fully read.
	//
	// After the Handler has returned, there is no guarantee
	// that the channel receives a value.
	//
	// If the protocol is HTTP/1.1 and CloseNotify is called while
	// processing an idempotent request (such a GET) while
	// HTTP/1.1 pipelining is in use, the arrival of a subsequent
	// pipelined request may cause a value to be sent on the
	// returned channel. In practice HTTP/1.1 pipelining is not
	// enabled in browsers and not seen often in the wild. If this
	// is a problem, use HTTP/2 or only use CloseNotify on methods
	// such as POST.
	CloseNotify() <-chan bool
}

var (
	// ServerContextKey is a context key. It can be used in HTTP
	// handlers with Context.Value to access the server that
	// started the handler. The associated value will be of
	// type *Server.
	ServerContextKey	= &contextKey{"http-server"}

	// LocalAddrContextKey is a context key. It can be used in
	// HTTP handlers with Context.Value to access the local
	// address the connection arrived on.
	// The associated value will be of type net.Addr.
	LocalAddrContextKey	= &contextKey{"local-addr"}
)

// A conn represents the server side of an HTTP connection.
type conn struct {
	// server is the server on which the connection arrived.
	// Immutable; never nil.
	server	*Server

	// cancelCtx cancels the connection-level context.
	cancelCtx	context.CancelFunc

	// rwc is the underlying network connection.
	// This is never wrapped by other types and is the value given out
	// to CloseNotifier callers. It is usually of type *net.TCPConn or
	// *tls.Conn.
	rwc	net.Conn

	// remoteAddr is rwc.RemoteAddr().String(). It is not populated synchronously
	// inside the Listener's Accept goroutine, as some implementations block.
	// It is populated immediately inside the (*conn).serve goroutine.
	// This is the value of a Handler's (*Request).RemoteAddr.
	remoteAddr	string

	// tlsState is the TLS connection state when using TLS.
	// nil means not TLS.
	tlsState	*tls.ConnectionState

	// werr is set to the first write error to rwc.
	// It is set via checkConnErrorWriter{w}, where bufw writes.
	werr	error

	// r is bufr's read source. It's a wrapper around rwc that provides
	// io.LimitedReader-style limiting (while reading request headers)
	// and functionality to support CloseNotifier. See *connReader docs.
	r	*connReader

	// bufr reads from r.
	bufr	*bufio.Reader

	// bufw writes to checkConnErrorWriter{c}, which populates werr on error.
	bufw	*bufio.Writer

	// lastMethod is the method of the most recent request
	// on this connection, if any.
	lastMethod	string

	curReq	atomic.Pointer[response]	// (which has a Request in it)

	curState	atomic.Uint64	// packed (unixtime<<8|uint8(ConnState))

	// mu guards hijackedv
	mu	sync.Mutex

	// hijackedv is whether this connection has been hijacked
	// by a Handler with the Hijacker interface.
	// It is guarded by mu.
	hijackedv	bool
}

func (c *conn) hijacked() bool {
//line /usr/local/go/src/net/http/server.go:309
	_go_fuzz_dep_.CoverTab[42099]++
							c.mu.Lock()
							defer c.mu.Unlock()
							return c.hijackedv
//line /usr/local/go/src/net/http/server.go:312
	// _ = "end of CoverTab[42099]"
}

// c.mu must be held.
func (c *conn) hijackLocked() (rwc net.Conn, buf *bufio.ReadWriter, err error) {
//line /usr/local/go/src/net/http/server.go:316
	_go_fuzz_dep_.CoverTab[42100]++
							if c.hijackedv {
//line /usr/local/go/src/net/http/server.go:317
		_go_fuzz_dep_.CoverTab[42103]++
								return nil, nil, ErrHijacked
//line /usr/local/go/src/net/http/server.go:318
		// _ = "end of CoverTab[42103]"
	} else {
//line /usr/local/go/src/net/http/server.go:319
		_go_fuzz_dep_.CoverTab[42104]++
//line /usr/local/go/src/net/http/server.go:319
		// _ = "end of CoverTab[42104]"
//line /usr/local/go/src/net/http/server.go:319
	}
//line /usr/local/go/src/net/http/server.go:319
	// _ = "end of CoverTab[42100]"
//line /usr/local/go/src/net/http/server.go:319
	_go_fuzz_dep_.CoverTab[42101]++
							c.r.abortPendingRead()

							c.hijackedv = true
							rwc = c.rwc
							rwc.SetDeadline(time.Time{})

							buf = bufio.NewReadWriter(c.bufr, bufio.NewWriter(rwc))
							if c.r.hasByte {
//line /usr/local/go/src/net/http/server.go:327
		_go_fuzz_dep_.CoverTab[42105]++
								if _, err := c.bufr.Peek(c.bufr.Buffered() + 1); err != nil {
//line /usr/local/go/src/net/http/server.go:328
			_go_fuzz_dep_.CoverTab[42106]++
									return nil, nil, fmt.Errorf("unexpected Peek failure reading buffered byte: %v", err)
//line /usr/local/go/src/net/http/server.go:329
			// _ = "end of CoverTab[42106]"
		} else {
//line /usr/local/go/src/net/http/server.go:330
			_go_fuzz_dep_.CoverTab[42107]++
//line /usr/local/go/src/net/http/server.go:330
			// _ = "end of CoverTab[42107]"
//line /usr/local/go/src/net/http/server.go:330
		}
//line /usr/local/go/src/net/http/server.go:330
		// _ = "end of CoverTab[42105]"
	} else {
//line /usr/local/go/src/net/http/server.go:331
		_go_fuzz_dep_.CoverTab[42108]++
//line /usr/local/go/src/net/http/server.go:331
		// _ = "end of CoverTab[42108]"
//line /usr/local/go/src/net/http/server.go:331
	}
//line /usr/local/go/src/net/http/server.go:331
	// _ = "end of CoverTab[42101]"
//line /usr/local/go/src/net/http/server.go:331
	_go_fuzz_dep_.CoverTab[42102]++
							c.setState(rwc, StateHijacked, runHooks)
							return
//line /usr/local/go/src/net/http/server.go:333
	// _ = "end of CoverTab[42102]"
}

// This should be >= 512 bytes for DetectContentType,
//line /usr/local/go/src/net/http/server.go:336
// but otherwise it's somewhat arbitrary.
//line /usr/local/go/src/net/http/server.go:338
const bufferBeforeChunkingSize = 2048

// chunkWriter writes to a response's conn buffer, and is the writer
//line /usr/local/go/src/net/http/server.go:340
// wrapped by the response.w buffered writer.
//line /usr/local/go/src/net/http/server.go:340
//
//line /usr/local/go/src/net/http/server.go:340
// chunkWriter also is responsible for finalizing the Header, including
//line /usr/local/go/src/net/http/server.go:340
// conditionally setting the Content-Type and setting a Content-Length
//line /usr/local/go/src/net/http/server.go:340
// in cases where the handler's final output is smaller than the buffer
//line /usr/local/go/src/net/http/server.go:340
// size. It also conditionally adds chunk headers, when in chunking mode.
//line /usr/local/go/src/net/http/server.go:340
//
//line /usr/local/go/src/net/http/server.go:340
// See the comment above (*response).Write for the entire write flow.
//line /usr/local/go/src/net/http/server.go:349
type chunkWriter struct {
	res	*response

	// header is either nil or a deep clone of res.handlerHeader
	// at the time of res.writeHeader, if res.writeHeader is
	// called and extra buffering is being done to calculate
	// Content-Type and/or Content-Length.
	header	Header

	// wroteHeader tells whether the header's been written to "the
	// wire" (or rather: w.conn.buf). this is unlike
	// (*response).wroteHeader, which tells only whether it was
	// logically written.
	wroteHeader	bool

	// set by the writeHeader method:
	chunking	bool	// using chunked transfer encoding for reply body
}

var (
	crlf		= []byte("\r\n")
	colonSpace	= []byte(": ")
)

func (cw *chunkWriter) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/server.go:373
	_go_fuzz_dep_.CoverTab[42109]++
							if !cw.wroteHeader {
//line /usr/local/go/src/net/http/server.go:374
		_go_fuzz_dep_.CoverTab[42115]++
								cw.writeHeader(p)
//line /usr/local/go/src/net/http/server.go:375
		// _ = "end of CoverTab[42115]"
	} else {
//line /usr/local/go/src/net/http/server.go:376
		_go_fuzz_dep_.CoverTab[42116]++
//line /usr/local/go/src/net/http/server.go:376
		// _ = "end of CoverTab[42116]"
//line /usr/local/go/src/net/http/server.go:376
	}
//line /usr/local/go/src/net/http/server.go:376
	// _ = "end of CoverTab[42109]"
//line /usr/local/go/src/net/http/server.go:376
	_go_fuzz_dep_.CoverTab[42110]++
							if cw.res.req.Method == "HEAD" {
//line /usr/local/go/src/net/http/server.go:377
		_go_fuzz_dep_.CoverTab[42117]++

								return len(p), nil
//line /usr/local/go/src/net/http/server.go:379
		// _ = "end of CoverTab[42117]"
	} else {
//line /usr/local/go/src/net/http/server.go:380
		_go_fuzz_dep_.CoverTab[42118]++
//line /usr/local/go/src/net/http/server.go:380
		// _ = "end of CoverTab[42118]"
//line /usr/local/go/src/net/http/server.go:380
	}
//line /usr/local/go/src/net/http/server.go:380
	// _ = "end of CoverTab[42110]"
//line /usr/local/go/src/net/http/server.go:380
	_go_fuzz_dep_.CoverTab[42111]++
							if cw.chunking {
//line /usr/local/go/src/net/http/server.go:381
		_go_fuzz_dep_.CoverTab[42119]++
								_, err = fmt.Fprintf(cw.res.conn.bufw, "%x\r\n", len(p))
								if err != nil {
//line /usr/local/go/src/net/http/server.go:383
			_go_fuzz_dep_.CoverTab[42120]++
									cw.res.conn.rwc.Close()
									return
//line /usr/local/go/src/net/http/server.go:385
			// _ = "end of CoverTab[42120]"
		} else {
//line /usr/local/go/src/net/http/server.go:386
			_go_fuzz_dep_.CoverTab[42121]++
//line /usr/local/go/src/net/http/server.go:386
			// _ = "end of CoverTab[42121]"
//line /usr/local/go/src/net/http/server.go:386
		}
//line /usr/local/go/src/net/http/server.go:386
		// _ = "end of CoverTab[42119]"
	} else {
//line /usr/local/go/src/net/http/server.go:387
		_go_fuzz_dep_.CoverTab[42122]++
//line /usr/local/go/src/net/http/server.go:387
		// _ = "end of CoverTab[42122]"
//line /usr/local/go/src/net/http/server.go:387
	}
//line /usr/local/go/src/net/http/server.go:387
	// _ = "end of CoverTab[42111]"
//line /usr/local/go/src/net/http/server.go:387
	_go_fuzz_dep_.CoverTab[42112]++
							n, err = cw.res.conn.bufw.Write(p)
							if cw.chunking && func() bool {
//line /usr/local/go/src/net/http/server.go:389
		_go_fuzz_dep_.CoverTab[42123]++
//line /usr/local/go/src/net/http/server.go:389
		return err == nil
//line /usr/local/go/src/net/http/server.go:389
		// _ = "end of CoverTab[42123]"
//line /usr/local/go/src/net/http/server.go:389
	}() {
//line /usr/local/go/src/net/http/server.go:389
		_go_fuzz_dep_.CoverTab[42124]++
								_, err = cw.res.conn.bufw.Write(crlf)
//line /usr/local/go/src/net/http/server.go:390
		// _ = "end of CoverTab[42124]"
	} else {
//line /usr/local/go/src/net/http/server.go:391
		_go_fuzz_dep_.CoverTab[42125]++
//line /usr/local/go/src/net/http/server.go:391
		// _ = "end of CoverTab[42125]"
//line /usr/local/go/src/net/http/server.go:391
	}
//line /usr/local/go/src/net/http/server.go:391
	// _ = "end of CoverTab[42112]"
//line /usr/local/go/src/net/http/server.go:391
	_go_fuzz_dep_.CoverTab[42113]++
							if err != nil {
//line /usr/local/go/src/net/http/server.go:392
		_go_fuzz_dep_.CoverTab[42126]++
								cw.res.conn.rwc.Close()
//line /usr/local/go/src/net/http/server.go:393
		// _ = "end of CoverTab[42126]"
	} else {
//line /usr/local/go/src/net/http/server.go:394
		_go_fuzz_dep_.CoverTab[42127]++
//line /usr/local/go/src/net/http/server.go:394
		// _ = "end of CoverTab[42127]"
//line /usr/local/go/src/net/http/server.go:394
	}
//line /usr/local/go/src/net/http/server.go:394
	// _ = "end of CoverTab[42113]"
//line /usr/local/go/src/net/http/server.go:394
	_go_fuzz_dep_.CoverTab[42114]++
							return
//line /usr/local/go/src/net/http/server.go:395
	// _ = "end of CoverTab[42114]"
}

func (cw *chunkWriter) flush() error {
//line /usr/local/go/src/net/http/server.go:398
	_go_fuzz_dep_.CoverTab[42128]++
							if !cw.wroteHeader {
//line /usr/local/go/src/net/http/server.go:399
		_go_fuzz_dep_.CoverTab[42130]++
								cw.writeHeader(nil)
//line /usr/local/go/src/net/http/server.go:400
		// _ = "end of CoverTab[42130]"
	} else {
//line /usr/local/go/src/net/http/server.go:401
		_go_fuzz_dep_.CoverTab[42131]++
//line /usr/local/go/src/net/http/server.go:401
		// _ = "end of CoverTab[42131]"
//line /usr/local/go/src/net/http/server.go:401
	}
//line /usr/local/go/src/net/http/server.go:401
	// _ = "end of CoverTab[42128]"
//line /usr/local/go/src/net/http/server.go:401
	_go_fuzz_dep_.CoverTab[42129]++
							return cw.res.conn.bufw.Flush()
//line /usr/local/go/src/net/http/server.go:402
	// _ = "end of CoverTab[42129]"
}

func (cw *chunkWriter) close() {
//line /usr/local/go/src/net/http/server.go:405
	_go_fuzz_dep_.CoverTab[42132]++
							if !cw.wroteHeader {
//line /usr/local/go/src/net/http/server.go:406
		_go_fuzz_dep_.CoverTab[42134]++
								cw.writeHeader(nil)
//line /usr/local/go/src/net/http/server.go:407
		// _ = "end of CoverTab[42134]"
	} else {
//line /usr/local/go/src/net/http/server.go:408
		_go_fuzz_dep_.CoverTab[42135]++
//line /usr/local/go/src/net/http/server.go:408
		// _ = "end of CoverTab[42135]"
//line /usr/local/go/src/net/http/server.go:408
	}
//line /usr/local/go/src/net/http/server.go:408
	// _ = "end of CoverTab[42132]"
//line /usr/local/go/src/net/http/server.go:408
	_go_fuzz_dep_.CoverTab[42133]++
							if cw.chunking {
//line /usr/local/go/src/net/http/server.go:409
		_go_fuzz_dep_.CoverTab[42136]++
								bw := cw.res.conn.bufw

								bw.WriteString("0\r\n")
								if trailers := cw.res.finalTrailers(); trailers != nil {
//line /usr/local/go/src/net/http/server.go:413
			_go_fuzz_dep_.CoverTab[42138]++
									trailers.Write(bw)
//line /usr/local/go/src/net/http/server.go:414
			// _ = "end of CoverTab[42138]"
		} else {
//line /usr/local/go/src/net/http/server.go:415
			_go_fuzz_dep_.CoverTab[42139]++
//line /usr/local/go/src/net/http/server.go:415
			// _ = "end of CoverTab[42139]"
//line /usr/local/go/src/net/http/server.go:415
		}
//line /usr/local/go/src/net/http/server.go:415
		// _ = "end of CoverTab[42136]"
//line /usr/local/go/src/net/http/server.go:415
		_go_fuzz_dep_.CoverTab[42137]++

//line /usr/local/go/src/net/http/server.go:418
		bw.WriteString("\r\n")
//line /usr/local/go/src/net/http/server.go:418
		// _ = "end of CoverTab[42137]"
	} else {
//line /usr/local/go/src/net/http/server.go:419
		_go_fuzz_dep_.CoverTab[42140]++
//line /usr/local/go/src/net/http/server.go:419
		// _ = "end of CoverTab[42140]"
//line /usr/local/go/src/net/http/server.go:419
	}
//line /usr/local/go/src/net/http/server.go:419
	// _ = "end of CoverTab[42133]"
}

// A response represents the server side of an HTTP response.
type response struct {
	conn			*conn
	req			*Request	// request for this response
	reqBody			io.ReadCloser
	cancelCtx		context.CancelFunc	// when ServeHTTP exits
	wroteHeader		bool			// a non-1xx header has been (logically) written
	wroteContinue		bool			// 100 Continue response was written
	wants10KeepAlive	bool			// HTTP/1.0 w/ Connection "keep-alive"
	wantsClose		bool			// HTTP request has Connection "close"

	// canWriteContinue is an atomic boolean that says whether or
	// not a 100 Continue header can be written to the
	// connection.
	// writeContinueMu must be held while writing the header.
	// These two fields together synchronize the body reader (the
	// expectContinueReader, which wants to write 100 Continue)
	// against the main writer.
	canWriteContinue	atomic.Bool
	writeContinueMu		sync.Mutex

	w	*bufio.Writer	// buffers output in chunks to chunkWriter
	cw	chunkWriter

	// handlerHeader is the Header that Handlers get access to,
	// which may be retained and mutated even after WriteHeader.
	// handlerHeader is copied into cw.header at WriteHeader
	// time, and privately mutated thereafter.
	handlerHeader	Header
	calledHeader	bool	// handler accessed handlerHeader via Header

	written		int64	// number of bytes written in body
	contentLength	int64	// explicitly-declared Content-Length; or -1
	status		int	// status code passed to WriteHeader

	// close connection after this reply.  set on request and
	// updated after response from handler if there's a
	// "Connection: keep-alive" response header and a
	// Content-Length.
	closeAfterReply	bool

	// requestBodyLimitHit is set by requestTooLarge when
	// maxBytesReader hits its max size. It is checked in
	// WriteHeader, to make sure we don't consume the
	// remaining request body to try to advance to the next HTTP
	// request. Instead, when this is set, we stop reading
	// subsequent requests on this connection and stop reading
	// input from it.
	requestBodyLimitHit	bool

	// trailers are the headers to be sent after the handler
	// finishes writing the body. This field is initialized from
	// the Trailer response header when the response header is
	// written.
	trailers	[]string

	handlerDone	atomic.Bool	// set true when the handler exits

	// Buffers for Date, Content-Length, and status code
	dateBuf		[len(TimeFormat)]byte
	clenBuf		[10]byte
	statusBuf	[3]byte

	// closeNotifyCh is the channel returned by CloseNotify.
	// TODO(bradfitz): this is currently (for Go 1.8) always
	// non-nil. Make this lazily-created again as it used to be?
	closeNotifyCh	chan bool
	didCloseNotify	atomic.Bool	// atomic (only false->true winner should send)
}

func (c *response) SetReadDeadline(deadline time.Time) error {
//line /usr/local/go/src/net/http/server.go:492
	_go_fuzz_dep_.CoverTab[42141]++
							return c.conn.rwc.SetReadDeadline(deadline)
//line /usr/local/go/src/net/http/server.go:493
	// _ = "end of CoverTab[42141]"
}

func (c *response) SetWriteDeadline(deadline time.Time) error {
//line /usr/local/go/src/net/http/server.go:496
	_go_fuzz_dep_.CoverTab[42142]++
							return c.conn.rwc.SetWriteDeadline(deadline)
//line /usr/local/go/src/net/http/server.go:497
	// _ = "end of CoverTab[42142]"
}

// TrailerPrefix is a magic prefix for ResponseWriter.Header map keys
//line /usr/local/go/src/net/http/server.go:500
// that, if present, signals that the map entry is actually for
//line /usr/local/go/src/net/http/server.go:500
// the response trailers, and not the response headers. The prefix
//line /usr/local/go/src/net/http/server.go:500
// is stripped after the ServeHTTP call finishes and the values are
//line /usr/local/go/src/net/http/server.go:500
// sent in the trailers.
//line /usr/local/go/src/net/http/server.go:500
//
//line /usr/local/go/src/net/http/server.go:500
// This mechanism is intended only for trailers that are not known
//line /usr/local/go/src/net/http/server.go:500
// prior to the headers being written. If the set of trailers is fixed
//line /usr/local/go/src/net/http/server.go:500
// or known before the header is written, the normal Go trailers mechanism
//line /usr/local/go/src/net/http/server.go:500
// is preferred:
//line /usr/local/go/src/net/http/server.go:500
//
//line /usr/local/go/src/net/http/server.go:500
//	https://pkg.go.dev/net/http#ResponseWriter
//line /usr/local/go/src/net/http/server.go:500
//	https://pkg.go.dev/net/http#example-ResponseWriter-Trailers
//line /usr/local/go/src/net/http/server.go:513
const TrailerPrefix = "Trailer:"

// finalTrailers is called after the Handler exits and returns a non-nil
//line /usr/local/go/src/net/http/server.go:515
// value if the Handler set any trailers.
//line /usr/local/go/src/net/http/server.go:517
func (w *response) finalTrailers() Header {
//line /usr/local/go/src/net/http/server.go:517
	_go_fuzz_dep_.CoverTab[42143]++
							var t Header
							for k, vv := range w.handlerHeader {
//line /usr/local/go/src/net/http/server.go:519
		_go_fuzz_dep_.CoverTab[42146]++
								if kk, found := strings.CutPrefix(k, TrailerPrefix); found {
//line /usr/local/go/src/net/http/server.go:520
			_go_fuzz_dep_.CoverTab[42147]++
									if t == nil {
//line /usr/local/go/src/net/http/server.go:521
				_go_fuzz_dep_.CoverTab[42149]++
										t = make(Header)
//line /usr/local/go/src/net/http/server.go:522
				// _ = "end of CoverTab[42149]"
			} else {
//line /usr/local/go/src/net/http/server.go:523
				_go_fuzz_dep_.CoverTab[42150]++
//line /usr/local/go/src/net/http/server.go:523
				// _ = "end of CoverTab[42150]"
//line /usr/local/go/src/net/http/server.go:523
			}
//line /usr/local/go/src/net/http/server.go:523
			// _ = "end of CoverTab[42147]"
//line /usr/local/go/src/net/http/server.go:523
			_go_fuzz_dep_.CoverTab[42148]++
									t[kk] = vv
//line /usr/local/go/src/net/http/server.go:524
			// _ = "end of CoverTab[42148]"
		} else {
//line /usr/local/go/src/net/http/server.go:525
			_go_fuzz_dep_.CoverTab[42151]++
//line /usr/local/go/src/net/http/server.go:525
			// _ = "end of CoverTab[42151]"
//line /usr/local/go/src/net/http/server.go:525
		}
//line /usr/local/go/src/net/http/server.go:525
		// _ = "end of CoverTab[42146]"
	}
//line /usr/local/go/src/net/http/server.go:526
	// _ = "end of CoverTab[42143]"
//line /usr/local/go/src/net/http/server.go:526
	_go_fuzz_dep_.CoverTab[42144]++
							for _, k := range w.trailers {
//line /usr/local/go/src/net/http/server.go:527
		_go_fuzz_dep_.CoverTab[42152]++
								if t == nil {
//line /usr/local/go/src/net/http/server.go:528
			_go_fuzz_dep_.CoverTab[42154]++
									t = make(Header)
//line /usr/local/go/src/net/http/server.go:529
			// _ = "end of CoverTab[42154]"
		} else {
//line /usr/local/go/src/net/http/server.go:530
			_go_fuzz_dep_.CoverTab[42155]++
//line /usr/local/go/src/net/http/server.go:530
			// _ = "end of CoverTab[42155]"
//line /usr/local/go/src/net/http/server.go:530
		}
//line /usr/local/go/src/net/http/server.go:530
		// _ = "end of CoverTab[42152]"
//line /usr/local/go/src/net/http/server.go:530
		_go_fuzz_dep_.CoverTab[42153]++
								for _, v := range w.handlerHeader[k] {
//line /usr/local/go/src/net/http/server.go:531
			_go_fuzz_dep_.CoverTab[42156]++
									t.Add(k, v)
//line /usr/local/go/src/net/http/server.go:532
			// _ = "end of CoverTab[42156]"
		}
//line /usr/local/go/src/net/http/server.go:533
		// _ = "end of CoverTab[42153]"
	}
//line /usr/local/go/src/net/http/server.go:534
	// _ = "end of CoverTab[42144]"
//line /usr/local/go/src/net/http/server.go:534
	_go_fuzz_dep_.CoverTab[42145]++
							return t
//line /usr/local/go/src/net/http/server.go:535
	// _ = "end of CoverTab[42145]"
}

// declareTrailer is called for each Trailer header when the
//line /usr/local/go/src/net/http/server.go:538
// response header is written. It notes that a header will need to be
//line /usr/local/go/src/net/http/server.go:538
// written in the trailers at the end of the response.
//line /usr/local/go/src/net/http/server.go:541
func (w *response) declareTrailer(k string) {
//line /usr/local/go/src/net/http/server.go:541
	_go_fuzz_dep_.CoverTab[42157]++
							k = CanonicalHeaderKey(k)
							if !httpguts.ValidTrailerHeader(k) {
//line /usr/local/go/src/net/http/server.go:543
		_go_fuzz_dep_.CoverTab[42159]++

								return
//line /usr/local/go/src/net/http/server.go:545
		// _ = "end of CoverTab[42159]"
	} else {
//line /usr/local/go/src/net/http/server.go:546
		_go_fuzz_dep_.CoverTab[42160]++
//line /usr/local/go/src/net/http/server.go:546
		// _ = "end of CoverTab[42160]"
//line /usr/local/go/src/net/http/server.go:546
	}
//line /usr/local/go/src/net/http/server.go:546
	// _ = "end of CoverTab[42157]"
//line /usr/local/go/src/net/http/server.go:546
	_go_fuzz_dep_.CoverTab[42158]++
							w.trailers = append(w.trailers, k)
//line /usr/local/go/src/net/http/server.go:547
	// _ = "end of CoverTab[42158]"
}

// requestTooLarge is called by maxBytesReader when too much input has
//line /usr/local/go/src/net/http/server.go:550
// been read from the client.
//line /usr/local/go/src/net/http/server.go:552
func (w *response) requestTooLarge() {
//line /usr/local/go/src/net/http/server.go:552
	_go_fuzz_dep_.CoverTab[42161]++
							w.closeAfterReply = true
							w.requestBodyLimitHit = true
							if !w.wroteHeader {
//line /usr/local/go/src/net/http/server.go:555
		_go_fuzz_dep_.CoverTab[42162]++
								w.Header().Set("Connection", "close")
//line /usr/local/go/src/net/http/server.go:556
		// _ = "end of CoverTab[42162]"
	} else {
//line /usr/local/go/src/net/http/server.go:557
		_go_fuzz_dep_.CoverTab[42163]++
//line /usr/local/go/src/net/http/server.go:557
		// _ = "end of CoverTab[42163]"
//line /usr/local/go/src/net/http/server.go:557
	}
//line /usr/local/go/src/net/http/server.go:557
	// _ = "end of CoverTab[42161]"
}

// writerOnly hides an io.Writer value's optional ReadFrom method
//line /usr/local/go/src/net/http/server.go:560
// from io.Copy.
//line /usr/local/go/src/net/http/server.go:562
type writerOnly struct {
	io.Writer
}

// ReadFrom is here to optimize copying from an *os.File regular file
//line /usr/local/go/src/net/http/server.go:566
// to a *net.TCPConn with sendfile, or from a supported src type such
//line /usr/local/go/src/net/http/server.go:566
// as a *net.TCPConn on Linux with splice.
//line /usr/local/go/src/net/http/server.go:569
func (w *response) ReadFrom(src io.Reader) (n int64, err error) {
//line /usr/local/go/src/net/http/server.go:569
	_go_fuzz_dep_.CoverTab[42164]++
							bufp := copyBufPool.Get().(*[]byte)
							buf := *bufp
							defer copyBufPool.Put(bufp)

//line /usr/local/go/src/net/http/server.go:577
	rf, ok := w.conn.rwc.(io.ReaderFrom)
	if !ok {
//line /usr/local/go/src/net/http/server.go:578
		_go_fuzz_dep_.CoverTab[42168]++
								return io.CopyBuffer(writerOnly{w}, src, buf)
//line /usr/local/go/src/net/http/server.go:579
		// _ = "end of CoverTab[42168]"
	} else {
//line /usr/local/go/src/net/http/server.go:580
		_go_fuzz_dep_.CoverTab[42169]++
//line /usr/local/go/src/net/http/server.go:580
		// _ = "end of CoverTab[42169]"
//line /usr/local/go/src/net/http/server.go:580
	}
//line /usr/local/go/src/net/http/server.go:580
	// _ = "end of CoverTab[42164]"
//line /usr/local/go/src/net/http/server.go:580
	_go_fuzz_dep_.CoverTab[42165]++

//line /usr/local/go/src/net/http/server.go:586
	if !w.cw.wroteHeader {
//line /usr/local/go/src/net/http/server.go:586
		_go_fuzz_dep_.CoverTab[42170]++
								n0, err := io.CopyBuffer(writerOnly{w}, io.LimitReader(src, sniffLen), buf)
								n += n0
								if err != nil || func() bool {
//line /usr/local/go/src/net/http/server.go:589
			_go_fuzz_dep_.CoverTab[42171]++
//line /usr/local/go/src/net/http/server.go:589
			return n0 < sniffLen
//line /usr/local/go/src/net/http/server.go:589
			// _ = "end of CoverTab[42171]"
//line /usr/local/go/src/net/http/server.go:589
		}() {
//line /usr/local/go/src/net/http/server.go:589
			_go_fuzz_dep_.CoverTab[42172]++
									return n, err
//line /usr/local/go/src/net/http/server.go:590
			// _ = "end of CoverTab[42172]"
		} else {
//line /usr/local/go/src/net/http/server.go:591
			_go_fuzz_dep_.CoverTab[42173]++
//line /usr/local/go/src/net/http/server.go:591
			// _ = "end of CoverTab[42173]"
//line /usr/local/go/src/net/http/server.go:591
		}
//line /usr/local/go/src/net/http/server.go:591
		// _ = "end of CoverTab[42170]"
	} else {
//line /usr/local/go/src/net/http/server.go:592
		_go_fuzz_dep_.CoverTab[42174]++
//line /usr/local/go/src/net/http/server.go:592
		// _ = "end of CoverTab[42174]"
//line /usr/local/go/src/net/http/server.go:592
	}
//line /usr/local/go/src/net/http/server.go:592
	// _ = "end of CoverTab[42165]"
//line /usr/local/go/src/net/http/server.go:592
	_go_fuzz_dep_.CoverTab[42166]++

							w.w.Flush()
							w.cw.flush()

//line /usr/local/go/src/net/http/server.go:598
	if !w.cw.chunking && func() bool {
//line /usr/local/go/src/net/http/server.go:598
		_go_fuzz_dep_.CoverTab[42175]++
//line /usr/local/go/src/net/http/server.go:598
		return w.bodyAllowed()
//line /usr/local/go/src/net/http/server.go:598
		// _ = "end of CoverTab[42175]"
//line /usr/local/go/src/net/http/server.go:598
	}() {
//line /usr/local/go/src/net/http/server.go:598
		_go_fuzz_dep_.CoverTab[42176]++
								n0, err := rf.ReadFrom(src)
								n += n0
								w.written += n0
								return n, err
//line /usr/local/go/src/net/http/server.go:602
		// _ = "end of CoverTab[42176]"
	} else {
//line /usr/local/go/src/net/http/server.go:603
		_go_fuzz_dep_.CoverTab[42177]++
//line /usr/local/go/src/net/http/server.go:603
		// _ = "end of CoverTab[42177]"
//line /usr/local/go/src/net/http/server.go:603
	}
//line /usr/local/go/src/net/http/server.go:603
	// _ = "end of CoverTab[42166]"
//line /usr/local/go/src/net/http/server.go:603
	_go_fuzz_dep_.CoverTab[42167]++

							n0, err := io.CopyBuffer(writerOnly{w}, src, buf)
							n += n0
							return n, err
//line /usr/local/go/src/net/http/server.go:607
	// _ = "end of CoverTab[42167]"
}

// debugServerConnections controls whether all server connections are wrapped
//line /usr/local/go/src/net/http/server.go:610
// with a verbose logging wrapper.
//line /usr/local/go/src/net/http/server.go:612
const debugServerConnections = false

// Create new connection from rwc.
func (srv *Server) newConn(rwc net.Conn) *conn {
//line /usr/local/go/src/net/http/server.go:615
	_go_fuzz_dep_.CoverTab[42178]++
							c := &conn{
		server:	srv,
		rwc:	rwc,
	}
	if debugServerConnections {
//line /usr/local/go/src/net/http/server.go:620
		_go_fuzz_dep_.CoverTab[42180]++
								c.rwc = newLoggingConn("server", c.rwc)
//line /usr/local/go/src/net/http/server.go:621
		// _ = "end of CoverTab[42180]"
	} else {
//line /usr/local/go/src/net/http/server.go:622
		_go_fuzz_dep_.CoverTab[42181]++
//line /usr/local/go/src/net/http/server.go:622
		// _ = "end of CoverTab[42181]"
//line /usr/local/go/src/net/http/server.go:622
	}
//line /usr/local/go/src/net/http/server.go:622
	// _ = "end of CoverTab[42178]"
//line /usr/local/go/src/net/http/server.go:622
	_go_fuzz_dep_.CoverTab[42179]++
							return c
//line /usr/local/go/src/net/http/server.go:623
	// _ = "end of CoverTab[42179]"
}

type readResult struct {
	_	incomparable
	n	int
	err	error
	b	byte	// byte read, if n == 1
}

// connReader is the io.Reader wrapper used by *conn. It combines a
//line /usr/local/go/src/net/http/server.go:633
// selectively-activated io.LimitedReader (to bound request header
//line /usr/local/go/src/net/http/server.go:633
// read sizes) with support for selectively keeping an io.Reader.Read
//line /usr/local/go/src/net/http/server.go:633
// call blocked in a background goroutine to wait for activity and
//line /usr/local/go/src/net/http/server.go:633
// trigger a CloseNotifier channel.
//line /usr/local/go/src/net/http/server.go:638
type connReader struct {
	conn	*conn

	mu	sync.Mutex	// guards following
	hasByte	bool
	byteBuf	[1]byte
	cond	*sync.Cond
	inRead	bool
	aborted	bool	// set true before conn.rwc deadline is set to past
	remain	int64	// bytes remaining
}

func (cr *connReader) lock() {
//line /usr/local/go/src/net/http/server.go:650
	_go_fuzz_dep_.CoverTab[42182]++
							cr.mu.Lock()
							if cr.cond == nil {
//line /usr/local/go/src/net/http/server.go:652
		_go_fuzz_dep_.CoverTab[42183]++
								cr.cond = sync.NewCond(&cr.mu)
//line /usr/local/go/src/net/http/server.go:653
		// _ = "end of CoverTab[42183]"
	} else {
//line /usr/local/go/src/net/http/server.go:654
		_go_fuzz_dep_.CoverTab[42184]++
//line /usr/local/go/src/net/http/server.go:654
		// _ = "end of CoverTab[42184]"
//line /usr/local/go/src/net/http/server.go:654
	}
//line /usr/local/go/src/net/http/server.go:654
	// _ = "end of CoverTab[42182]"
}

func (cr *connReader) unlock() {
//line /usr/local/go/src/net/http/server.go:657
	_go_fuzz_dep_.CoverTab[42185]++
//line /usr/local/go/src/net/http/server.go:657
	cr.mu.Unlock()
//line /usr/local/go/src/net/http/server.go:657
	// _ = "end of CoverTab[42185]"
//line /usr/local/go/src/net/http/server.go:657
}

func (cr *connReader) startBackgroundRead() {
//line /usr/local/go/src/net/http/server.go:659
	_go_fuzz_dep_.CoverTab[42186]++
							cr.lock()
							defer cr.unlock()
							if cr.inRead {
//line /usr/local/go/src/net/http/server.go:662
		_go_fuzz_dep_.CoverTab[42189]++
								panic("invalid concurrent Body.Read call")
//line /usr/local/go/src/net/http/server.go:663
		// _ = "end of CoverTab[42189]"
	} else {
//line /usr/local/go/src/net/http/server.go:664
		_go_fuzz_dep_.CoverTab[42190]++
//line /usr/local/go/src/net/http/server.go:664
		// _ = "end of CoverTab[42190]"
//line /usr/local/go/src/net/http/server.go:664
	}
//line /usr/local/go/src/net/http/server.go:664
	// _ = "end of CoverTab[42186]"
//line /usr/local/go/src/net/http/server.go:664
	_go_fuzz_dep_.CoverTab[42187]++
							if cr.hasByte {
//line /usr/local/go/src/net/http/server.go:665
		_go_fuzz_dep_.CoverTab[42191]++
								return
//line /usr/local/go/src/net/http/server.go:666
		// _ = "end of CoverTab[42191]"
	} else {
//line /usr/local/go/src/net/http/server.go:667
		_go_fuzz_dep_.CoverTab[42192]++
//line /usr/local/go/src/net/http/server.go:667
		// _ = "end of CoverTab[42192]"
//line /usr/local/go/src/net/http/server.go:667
	}
//line /usr/local/go/src/net/http/server.go:667
	// _ = "end of CoverTab[42187]"
//line /usr/local/go/src/net/http/server.go:667
	_go_fuzz_dep_.CoverTab[42188]++
							cr.inRead = true
							cr.conn.rwc.SetReadDeadline(time.Time{})
//line /usr/local/go/src/net/http/server.go:669
	_curRoutineNum31_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/server.go:669
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum31_)
							go cr.backgroundRead()
//line /usr/local/go/src/net/http/server.go:670
	// _ = "end of CoverTab[42188]"
}

func (cr *connReader) backgroundRead() {
//line /usr/local/go/src/net/http/server.go:673
	_go_fuzz_dep_.CoverTab[42193]++
							n, err := cr.conn.rwc.Read(cr.byteBuf[:])
							cr.lock()
							if n == 1 {
//line /usr/local/go/src/net/http/server.go:676
		_go_fuzz_dep_.CoverTab[42196]++
								cr.hasByte = true
//line /usr/local/go/src/net/http/server.go:677
		// _ = "end of CoverTab[42196]"

//line /usr/local/go/src/net/http/server.go:700
	} else {
//line /usr/local/go/src/net/http/server.go:700
		_go_fuzz_dep_.CoverTab[42197]++
//line /usr/local/go/src/net/http/server.go:700
		// _ = "end of CoverTab[42197]"
//line /usr/local/go/src/net/http/server.go:700
	}
//line /usr/local/go/src/net/http/server.go:700
	// _ = "end of CoverTab[42193]"
//line /usr/local/go/src/net/http/server.go:700
	_go_fuzz_dep_.CoverTab[42194]++
							if ne, ok := err.(net.Error); ok && func() bool {
//line /usr/local/go/src/net/http/server.go:701
		_go_fuzz_dep_.CoverTab[42198]++
//line /usr/local/go/src/net/http/server.go:701
		return cr.aborted
//line /usr/local/go/src/net/http/server.go:701
		// _ = "end of CoverTab[42198]"
//line /usr/local/go/src/net/http/server.go:701
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:701
		_go_fuzz_dep_.CoverTab[42199]++
//line /usr/local/go/src/net/http/server.go:701
		return ne.Timeout()
//line /usr/local/go/src/net/http/server.go:701
		// _ = "end of CoverTab[42199]"
//line /usr/local/go/src/net/http/server.go:701
	}() {
//line /usr/local/go/src/net/http/server.go:701
		_go_fuzz_dep_.CoverTab[42200]++
//line /usr/local/go/src/net/http/server.go:701
		// _ = "end of CoverTab[42200]"

//line /usr/local/go/src/net/http/server.go:704
	} else {
//line /usr/local/go/src/net/http/server.go:704
		_go_fuzz_dep_.CoverTab[42201]++
//line /usr/local/go/src/net/http/server.go:704
		if err != nil {
//line /usr/local/go/src/net/http/server.go:704
			_go_fuzz_dep_.CoverTab[42202]++
									cr.handleReadError(err)
//line /usr/local/go/src/net/http/server.go:705
			// _ = "end of CoverTab[42202]"
		} else {
//line /usr/local/go/src/net/http/server.go:706
			_go_fuzz_dep_.CoverTab[42203]++
//line /usr/local/go/src/net/http/server.go:706
			// _ = "end of CoverTab[42203]"
//line /usr/local/go/src/net/http/server.go:706
		}
//line /usr/local/go/src/net/http/server.go:706
		// _ = "end of CoverTab[42201]"
//line /usr/local/go/src/net/http/server.go:706
	}
//line /usr/local/go/src/net/http/server.go:706
	// _ = "end of CoverTab[42194]"
//line /usr/local/go/src/net/http/server.go:706
	_go_fuzz_dep_.CoverTab[42195]++
							cr.aborted = false
							cr.inRead = false
							cr.unlock()
							cr.cond.Broadcast()
//line /usr/local/go/src/net/http/server.go:710
	// _ = "end of CoverTab[42195]"
}

func (cr *connReader) abortPendingRead() {
//line /usr/local/go/src/net/http/server.go:713
	_go_fuzz_dep_.CoverTab[42204]++
							cr.lock()
							defer cr.unlock()
							if !cr.inRead {
//line /usr/local/go/src/net/http/server.go:716
		_go_fuzz_dep_.CoverTab[42207]++
								return
//line /usr/local/go/src/net/http/server.go:717
		// _ = "end of CoverTab[42207]"
	} else {
//line /usr/local/go/src/net/http/server.go:718
		_go_fuzz_dep_.CoverTab[42208]++
//line /usr/local/go/src/net/http/server.go:718
		// _ = "end of CoverTab[42208]"
//line /usr/local/go/src/net/http/server.go:718
	}
//line /usr/local/go/src/net/http/server.go:718
	// _ = "end of CoverTab[42204]"
//line /usr/local/go/src/net/http/server.go:718
	_go_fuzz_dep_.CoverTab[42205]++
							cr.aborted = true
							cr.conn.rwc.SetReadDeadline(aLongTimeAgo)
							for cr.inRead {
//line /usr/local/go/src/net/http/server.go:721
		_go_fuzz_dep_.CoverTab[42209]++
								cr.cond.Wait()
//line /usr/local/go/src/net/http/server.go:722
		// _ = "end of CoverTab[42209]"
	}
//line /usr/local/go/src/net/http/server.go:723
	// _ = "end of CoverTab[42205]"
//line /usr/local/go/src/net/http/server.go:723
	_go_fuzz_dep_.CoverTab[42206]++
							cr.conn.rwc.SetReadDeadline(time.Time{})
//line /usr/local/go/src/net/http/server.go:724
	// _ = "end of CoverTab[42206]"
}

func (cr *connReader) setReadLimit(remain int64) {
//line /usr/local/go/src/net/http/server.go:727
	_go_fuzz_dep_.CoverTab[42210]++
//line /usr/local/go/src/net/http/server.go:727
	cr.remain = remain
//line /usr/local/go/src/net/http/server.go:727
	// _ = "end of CoverTab[42210]"
//line /usr/local/go/src/net/http/server.go:727
}
func (cr *connReader) setInfiniteReadLimit() {
//line /usr/local/go/src/net/http/server.go:728
	_go_fuzz_dep_.CoverTab[42211]++
//line /usr/local/go/src/net/http/server.go:728
	cr.remain = maxInt64
//line /usr/local/go/src/net/http/server.go:728
	// _ = "end of CoverTab[42211]"
//line /usr/local/go/src/net/http/server.go:728
}
func (cr *connReader) hitReadLimit() bool {
//line /usr/local/go/src/net/http/server.go:729
	_go_fuzz_dep_.CoverTab[42212]++
//line /usr/local/go/src/net/http/server.go:729
	return cr.remain <= 0
//line /usr/local/go/src/net/http/server.go:729
	// _ = "end of CoverTab[42212]"
//line /usr/local/go/src/net/http/server.go:729
}

// handleReadError is called whenever a Read from the client returns a
//line /usr/local/go/src/net/http/server.go:731
// non-nil error.
//line /usr/local/go/src/net/http/server.go:731
//
//line /usr/local/go/src/net/http/server.go:731
// The provided non-nil err is almost always io.EOF or a "use of
//line /usr/local/go/src/net/http/server.go:731
// closed network connection". In any case, the error is not
//line /usr/local/go/src/net/http/server.go:731
// particularly interesting, except perhaps for debugging during
//line /usr/local/go/src/net/http/server.go:731
// development. Any error means the connection is dead and we should
//line /usr/local/go/src/net/http/server.go:731
// down its context.
//line /usr/local/go/src/net/http/server.go:731
//
//line /usr/local/go/src/net/http/server.go:731
// It may be called from multiple goroutines.
//line /usr/local/go/src/net/http/server.go:741
func (cr *connReader) handleReadError(_ error) {
//line /usr/local/go/src/net/http/server.go:741
	_go_fuzz_dep_.CoverTab[42213]++
							cr.conn.cancelCtx()
							cr.closeNotify()
//line /usr/local/go/src/net/http/server.go:743
	// _ = "end of CoverTab[42213]"
}

// may be called from multiple goroutines.
func (cr *connReader) closeNotify() {
//line /usr/local/go/src/net/http/server.go:747
	_go_fuzz_dep_.CoverTab[42214]++
							res := cr.conn.curReq.Load()
							if res != nil && func() bool {
//line /usr/local/go/src/net/http/server.go:749
		_go_fuzz_dep_.CoverTab[42215]++
//line /usr/local/go/src/net/http/server.go:749
		return !res.didCloseNotify.Swap(true)
//line /usr/local/go/src/net/http/server.go:749
		// _ = "end of CoverTab[42215]"
//line /usr/local/go/src/net/http/server.go:749
	}() {
//line /usr/local/go/src/net/http/server.go:749
		_go_fuzz_dep_.CoverTab[42216]++
								res.closeNotifyCh <- true
//line /usr/local/go/src/net/http/server.go:750
		// _ = "end of CoverTab[42216]"
	} else {
//line /usr/local/go/src/net/http/server.go:751
		_go_fuzz_dep_.CoverTab[42217]++
//line /usr/local/go/src/net/http/server.go:751
		// _ = "end of CoverTab[42217]"
//line /usr/local/go/src/net/http/server.go:751
	}
//line /usr/local/go/src/net/http/server.go:751
	// _ = "end of CoverTab[42214]"
}

func (cr *connReader) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/server.go:754
	_go_fuzz_dep_.CoverTab[42218]++
							cr.lock()
							if cr.inRead {
//line /usr/local/go/src/net/http/server.go:756
		_go_fuzz_dep_.CoverTab[42225]++
								cr.unlock()
								if cr.conn.hijacked() {
//line /usr/local/go/src/net/http/server.go:758
			_go_fuzz_dep_.CoverTab[42227]++
									panic("invalid Body.Read call. After hijacked, the original Request must not be used")
//line /usr/local/go/src/net/http/server.go:759
			// _ = "end of CoverTab[42227]"
		} else {
//line /usr/local/go/src/net/http/server.go:760
			_go_fuzz_dep_.CoverTab[42228]++
//line /usr/local/go/src/net/http/server.go:760
			// _ = "end of CoverTab[42228]"
//line /usr/local/go/src/net/http/server.go:760
		}
//line /usr/local/go/src/net/http/server.go:760
		// _ = "end of CoverTab[42225]"
//line /usr/local/go/src/net/http/server.go:760
		_go_fuzz_dep_.CoverTab[42226]++
								panic("invalid concurrent Body.Read call")
//line /usr/local/go/src/net/http/server.go:761
		// _ = "end of CoverTab[42226]"
	} else {
//line /usr/local/go/src/net/http/server.go:762
		_go_fuzz_dep_.CoverTab[42229]++
//line /usr/local/go/src/net/http/server.go:762
		// _ = "end of CoverTab[42229]"
//line /usr/local/go/src/net/http/server.go:762
	}
//line /usr/local/go/src/net/http/server.go:762
	// _ = "end of CoverTab[42218]"
//line /usr/local/go/src/net/http/server.go:762
	_go_fuzz_dep_.CoverTab[42219]++
							if cr.hitReadLimit() {
//line /usr/local/go/src/net/http/server.go:763
		_go_fuzz_dep_.CoverTab[42230]++
								cr.unlock()
								return 0, io.EOF
//line /usr/local/go/src/net/http/server.go:765
		// _ = "end of CoverTab[42230]"
	} else {
//line /usr/local/go/src/net/http/server.go:766
		_go_fuzz_dep_.CoverTab[42231]++
//line /usr/local/go/src/net/http/server.go:766
		// _ = "end of CoverTab[42231]"
//line /usr/local/go/src/net/http/server.go:766
	}
//line /usr/local/go/src/net/http/server.go:766
	// _ = "end of CoverTab[42219]"
//line /usr/local/go/src/net/http/server.go:766
	_go_fuzz_dep_.CoverTab[42220]++
							if len(p) == 0 {
//line /usr/local/go/src/net/http/server.go:767
		_go_fuzz_dep_.CoverTab[42232]++
								cr.unlock()
								return 0, nil
//line /usr/local/go/src/net/http/server.go:769
		// _ = "end of CoverTab[42232]"
	} else {
//line /usr/local/go/src/net/http/server.go:770
		_go_fuzz_dep_.CoverTab[42233]++
//line /usr/local/go/src/net/http/server.go:770
		// _ = "end of CoverTab[42233]"
//line /usr/local/go/src/net/http/server.go:770
	}
//line /usr/local/go/src/net/http/server.go:770
	// _ = "end of CoverTab[42220]"
//line /usr/local/go/src/net/http/server.go:770
	_go_fuzz_dep_.CoverTab[42221]++
							if int64(len(p)) > cr.remain {
//line /usr/local/go/src/net/http/server.go:771
		_go_fuzz_dep_.CoverTab[42234]++
								p = p[:cr.remain]
//line /usr/local/go/src/net/http/server.go:772
		// _ = "end of CoverTab[42234]"
	} else {
//line /usr/local/go/src/net/http/server.go:773
		_go_fuzz_dep_.CoverTab[42235]++
//line /usr/local/go/src/net/http/server.go:773
		// _ = "end of CoverTab[42235]"
//line /usr/local/go/src/net/http/server.go:773
	}
//line /usr/local/go/src/net/http/server.go:773
	// _ = "end of CoverTab[42221]"
//line /usr/local/go/src/net/http/server.go:773
	_go_fuzz_dep_.CoverTab[42222]++
							if cr.hasByte {
//line /usr/local/go/src/net/http/server.go:774
		_go_fuzz_dep_.CoverTab[42236]++
								p[0] = cr.byteBuf[0]
								cr.hasByte = false
								cr.unlock()
								return 1, nil
//line /usr/local/go/src/net/http/server.go:778
		// _ = "end of CoverTab[42236]"
	} else {
//line /usr/local/go/src/net/http/server.go:779
		_go_fuzz_dep_.CoverTab[42237]++
//line /usr/local/go/src/net/http/server.go:779
		// _ = "end of CoverTab[42237]"
//line /usr/local/go/src/net/http/server.go:779
	}
//line /usr/local/go/src/net/http/server.go:779
	// _ = "end of CoverTab[42222]"
//line /usr/local/go/src/net/http/server.go:779
	_go_fuzz_dep_.CoverTab[42223]++
							cr.inRead = true
							cr.unlock()
							n, err = cr.conn.rwc.Read(p)

							cr.lock()
							cr.inRead = false
							if err != nil {
//line /usr/local/go/src/net/http/server.go:786
		_go_fuzz_dep_.CoverTab[42238]++
								cr.handleReadError(err)
//line /usr/local/go/src/net/http/server.go:787
		// _ = "end of CoverTab[42238]"
	} else {
//line /usr/local/go/src/net/http/server.go:788
		_go_fuzz_dep_.CoverTab[42239]++
//line /usr/local/go/src/net/http/server.go:788
		// _ = "end of CoverTab[42239]"
//line /usr/local/go/src/net/http/server.go:788
	}
//line /usr/local/go/src/net/http/server.go:788
	// _ = "end of CoverTab[42223]"
//line /usr/local/go/src/net/http/server.go:788
	_go_fuzz_dep_.CoverTab[42224]++
							cr.remain -= int64(n)
							cr.unlock()

							cr.cond.Broadcast()
							return n, err
//line /usr/local/go/src/net/http/server.go:793
	// _ = "end of CoverTab[42224]"
}

var (
	bufioReaderPool		sync.Pool
	bufioWriter2kPool	sync.Pool
	bufioWriter4kPool	sync.Pool
)

var copyBufPool = sync.Pool{
	New: func() any {
//line /usr/local/go/src/net/http/server.go:803
		_go_fuzz_dep_.CoverTab[42240]++
								b := make([]byte, 32*1024)
								return &b
//line /usr/local/go/src/net/http/server.go:805
		// _ = "end of CoverTab[42240]"
	},
}

func bufioWriterPool(size int) *sync.Pool {
//line /usr/local/go/src/net/http/server.go:809
	_go_fuzz_dep_.CoverTab[42241]++
							switch size {
	case 2 << 10:
//line /usr/local/go/src/net/http/server.go:811
		_go_fuzz_dep_.CoverTab[42243]++
								return &bufioWriter2kPool
//line /usr/local/go/src/net/http/server.go:812
		// _ = "end of CoverTab[42243]"
	case 4 << 10:
//line /usr/local/go/src/net/http/server.go:813
		_go_fuzz_dep_.CoverTab[42244]++
								return &bufioWriter4kPool
//line /usr/local/go/src/net/http/server.go:814
		// _ = "end of CoverTab[42244]"
//line /usr/local/go/src/net/http/server.go:814
	default:
//line /usr/local/go/src/net/http/server.go:814
		_go_fuzz_dep_.CoverTab[42245]++
//line /usr/local/go/src/net/http/server.go:814
		// _ = "end of CoverTab[42245]"
	}
//line /usr/local/go/src/net/http/server.go:815
	// _ = "end of CoverTab[42241]"
//line /usr/local/go/src/net/http/server.go:815
	_go_fuzz_dep_.CoverTab[42242]++
							return nil
//line /usr/local/go/src/net/http/server.go:816
	// _ = "end of CoverTab[42242]"
}

func newBufioReader(r io.Reader) *bufio.Reader {
//line /usr/local/go/src/net/http/server.go:819
	_go_fuzz_dep_.CoverTab[42246]++
							if v := bufioReaderPool.Get(); v != nil {
//line /usr/local/go/src/net/http/server.go:820
		_go_fuzz_dep_.CoverTab[42248]++
								br := v.(*bufio.Reader)
								br.Reset(r)
								return br
//line /usr/local/go/src/net/http/server.go:823
		// _ = "end of CoverTab[42248]"
	} else {
//line /usr/local/go/src/net/http/server.go:824
		_go_fuzz_dep_.CoverTab[42249]++
//line /usr/local/go/src/net/http/server.go:824
		// _ = "end of CoverTab[42249]"
//line /usr/local/go/src/net/http/server.go:824
	}
//line /usr/local/go/src/net/http/server.go:824
	// _ = "end of CoverTab[42246]"
//line /usr/local/go/src/net/http/server.go:824
	_go_fuzz_dep_.CoverTab[42247]++

//line /usr/local/go/src/net/http/server.go:827
	return bufio.NewReader(r)
//line /usr/local/go/src/net/http/server.go:827
	// _ = "end of CoverTab[42247]"
}

func putBufioReader(br *bufio.Reader) {
//line /usr/local/go/src/net/http/server.go:830
	_go_fuzz_dep_.CoverTab[42250]++
							br.Reset(nil)
							bufioReaderPool.Put(br)
//line /usr/local/go/src/net/http/server.go:832
	// _ = "end of CoverTab[42250]"
}

func newBufioWriterSize(w io.Writer, size int) *bufio.Writer {
//line /usr/local/go/src/net/http/server.go:835
	_go_fuzz_dep_.CoverTab[42251]++
							pool := bufioWriterPool(size)
							if pool != nil {
//line /usr/local/go/src/net/http/server.go:837
		_go_fuzz_dep_.CoverTab[42253]++
								if v := pool.Get(); v != nil {
//line /usr/local/go/src/net/http/server.go:838
			_go_fuzz_dep_.CoverTab[42254]++
									bw := v.(*bufio.Writer)
									bw.Reset(w)
									return bw
//line /usr/local/go/src/net/http/server.go:841
			// _ = "end of CoverTab[42254]"
		} else {
//line /usr/local/go/src/net/http/server.go:842
			_go_fuzz_dep_.CoverTab[42255]++
//line /usr/local/go/src/net/http/server.go:842
			// _ = "end of CoverTab[42255]"
//line /usr/local/go/src/net/http/server.go:842
		}
//line /usr/local/go/src/net/http/server.go:842
		// _ = "end of CoverTab[42253]"
	} else {
//line /usr/local/go/src/net/http/server.go:843
		_go_fuzz_dep_.CoverTab[42256]++
//line /usr/local/go/src/net/http/server.go:843
		// _ = "end of CoverTab[42256]"
//line /usr/local/go/src/net/http/server.go:843
	}
//line /usr/local/go/src/net/http/server.go:843
	// _ = "end of CoverTab[42251]"
//line /usr/local/go/src/net/http/server.go:843
	_go_fuzz_dep_.CoverTab[42252]++
							return bufio.NewWriterSize(w, size)
//line /usr/local/go/src/net/http/server.go:844
	// _ = "end of CoverTab[42252]"
}

func putBufioWriter(bw *bufio.Writer) {
//line /usr/local/go/src/net/http/server.go:847
	_go_fuzz_dep_.CoverTab[42257]++
							bw.Reset(nil)
							if pool := bufioWriterPool(bw.Available()); pool != nil {
//line /usr/local/go/src/net/http/server.go:849
		_go_fuzz_dep_.CoverTab[42258]++
								pool.Put(bw)
//line /usr/local/go/src/net/http/server.go:850
		// _ = "end of CoverTab[42258]"
	} else {
//line /usr/local/go/src/net/http/server.go:851
		_go_fuzz_dep_.CoverTab[42259]++
//line /usr/local/go/src/net/http/server.go:851
		// _ = "end of CoverTab[42259]"
//line /usr/local/go/src/net/http/server.go:851
	}
//line /usr/local/go/src/net/http/server.go:851
	// _ = "end of CoverTab[42257]"
}

// DefaultMaxHeaderBytes is the maximum permitted size of the headers
//line /usr/local/go/src/net/http/server.go:854
// in an HTTP request.
//line /usr/local/go/src/net/http/server.go:854
// This can be overridden by setting Server.MaxHeaderBytes.
//line /usr/local/go/src/net/http/server.go:857
const DefaultMaxHeaderBytes = 1 << 20	// 1 MB

func (srv *Server) maxHeaderBytes() int {
//line /usr/local/go/src/net/http/server.go:859
	_go_fuzz_dep_.CoverTab[42260]++
							if srv.MaxHeaderBytes > 0 {
//line /usr/local/go/src/net/http/server.go:860
		_go_fuzz_dep_.CoverTab[42262]++
								return srv.MaxHeaderBytes
//line /usr/local/go/src/net/http/server.go:861
		// _ = "end of CoverTab[42262]"
	} else {
//line /usr/local/go/src/net/http/server.go:862
		_go_fuzz_dep_.CoverTab[42263]++
//line /usr/local/go/src/net/http/server.go:862
		// _ = "end of CoverTab[42263]"
//line /usr/local/go/src/net/http/server.go:862
	}
//line /usr/local/go/src/net/http/server.go:862
	// _ = "end of CoverTab[42260]"
//line /usr/local/go/src/net/http/server.go:862
	_go_fuzz_dep_.CoverTab[42261]++
							return DefaultMaxHeaderBytes
//line /usr/local/go/src/net/http/server.go:863
	// _ = "end of CoverTab[42261]"
}

func (srv *Server) initialReadLimitSize() int64 {
//line /usr/local/go/src/net/http/server.go:866
	_go_fuzz_dep_.CoverTab[42264]++
							return int64(srv.maxHeaderBytes()) + 4096
//line /usr/local/go/src/net/http/server.go:867
	// _ = "end of CoverTab[42264]"
}

// tlsHandshakeTimeout returns the time limit permitted for the TLS
//line /usr/local/go/src/net/http/server.go:870
// handshake, or zero for unlimited.
//line /usr/local/go/src/net/http/server.go:870
//
//line /usr/local/go/src/net/http/server.go:870
// It returns the minimum of any positive ReadHeaderTimeout,
//line /usr/local/go/src/net/http/server.go:870
// ReadTimeout, or WriteTimeout.
//line /usr/local/go/src/net/http/server.go:875
func (srv *Server) tlsHandshakeTimeout() time.Duration {
//line /usr/local/go/src/net/http/server.go:875
	_go_fuzz_dep_.CoverTab[42265]++
							var ret time.Duration
							for _, v := range [...]time.Duration{
		srv.ReadHeaderTimeout,
		srv.ReadTimeout,
		srv.WriteTimeout,
	} {
//line /usr/local/go/src/net/http/server.go:881
		_go_fuzz_dep_.CoverTab[42267]++
								if v <= 0 {
//line /usr/local/go/src/net/http/server.go:882
			_go_fuzz_dep_.CoverTab[42269]++
									continue
//line /usr/local/go/src/net/http/server.go:883
			// _ = "end of CoverTab[42269]"
		} else {
//line /usr/local/go/src/net/http/server.go:884
			_go_fuzz_dep_.CoverTab[42270]++
//line /usr/local/go/src/net/http/server.go:884
			// _ = "end of CoverTab[42270]"
//line /usr/local/go/src/net/http/server.go:884
		}
//line /usr/local/go/src/net/http/server.go:884
		// _ = "end of CoverTab[42267]"
//line /usr/local/go/src/net/http/server.go:884
		_go_fuzz_dep_.CoverTab[42268]++
								if ret == 0 || func() bool {
//line /usr/local/go/src/net/http/server.go:885
			_go_fuzz_dep_.CoverTab[42271]++
//line /usr/local/go/src/net/http/server.go:885
			return v < ret
//line /usr/local/go/src/net/http/server.go:885
			// _ = "end of CoverTab[42271]"
//line /usr/local/go/src/net/http/server.go:885
		}() {
//line /usr/local/go/src/net/http/server.go:885
			_go_fuzz_dep_.CoverTab[42272]++
									ret = v
//line /usr/local/go/src/net/http/server.go:886
			// _ = "end of CoverTab[42272]"
		} else {
//line /usr/local/go/src/net/http/server.go:887
			_go_fuzz_dep_.CoverTab[42273]++
//line /usr/local/go/src/net/http/server.go:887
			// _ = "end of CoverTab[42273]"
//line /usr/local/go/src/net/http/server.go:887
		}
//line /usr/local/go/src/net/http/server.go:887
		// _ = "end of CoverTab[42268]"
	}
//line /usr/local/go/src/net/http/server.go:888
	// _ = "end of CoverTab[42265]"
//line /usr/local/go/src/net/http/server.go:888
	_go_fuzz_dep_.CoverTab[42266]++
							return ret
//line /usr/local/go/src/net/http/server.go:889
	// _ = "end of CoverTab[42266]"
}

// wrapper around io.ReadCloser which on first read, sends an
//line /usr/local/go/src/net/http/server.go:892
// HTTP/1.1 100 Continue header
//line /usr/local/go/src/net/http/server.go:894
type expectContinueReader struct {
	resp		*response
	readCloser	io.ReadCloser
	closed		atomic.Bool
	sawEOF		atomic.Bool
}

func (ecr *expectContinueReader) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/server.go:901
	_go_fuzz_dep_.CoverTab[42274]++
							if ecr.closed.Load() {
//line /usr/local/go/src/net/http/server.go:902
		_go_fuzz_dep_.CoverTab[42278]++
								return 0, ErrBodyReadAfterClose
//line /usr/local/go/src/net/http/server.go:903
		// _ = "end of CoverTab[42278]"
	} else {
//line /usr/local/go/src/net/http/server.go:904
		_go_fuzz_dep_.CoverTab[42279]++
//line /usr/local/go/src/net/http/server.go:904
		// _ = "end of CoverTab[42279]"
//line /usr/local/go/src/net/http/server.go:904
	}
//line /usr/local/go/src/net/http/server.go:904
	// _ = "end of CoverTab[42274]"
//line /usr/local/go/src/net/http/server.go:904
	_go_fuzz_dep_.CoverTab[42275]++
							w := ecr.resp
							if !w.wroteContinue && func() bool {
//line /usr/local/go/src/net/http/server.go:906
		_go_fuzz_dep_.CoverTab[42280]++
//line /usr/local/go/src/net/http/server.go:906
		return w.canWriteContinue.Load()
//line /usr/local/go/src/net/http/server.go:906
		// _ = "end of CoverTab[42280]"
//line /usr/local/go/src/net/http/server.go:906
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:906
		_go_fuzz_dep_.CoverTab[42281]++
//line /usr/local/go/src/net/http/server.go:906
		return !w.conn.hijacked()
//line /usr/local/go/src/net/http/server.go:906
		// _ = "end of CoverTab[42281]"
//line /usr/local/go/src/net/http/server.go:906
	}() {
//line /usr/local/go/src/net/http/server.go:906
		_go_fuzz_dep_.CoverTab[42282]++
								w.wroteContinue = true
								w.writeContinueMu.Lock()
								if w.canWriteContinue.Load() {
//line /usr/local/go/src/net/http/server.go:909
			_go_fuzz_dep_.CoverTab[42284]++
									w.conn.bufw.WriteString("HTTP/1.1 100 Continue\r\n\r\n")
									w.conn.bufw.Flush()
									w.canWriteContinue.Store(false)
//line /usr/local/go/src/net/http/server.go:912
			// _ = "end of CoverTab[42284]"
		} else {
//line /usr/local/go/src/net/http/server.go:913
			_go_fuzz_dep_.CoverTab[42285]++
//line /usr/local/go/src/net/http/server.go:913
			// _ = "end of CoverTab[42285]"
//line /usr/local/go/src/net/http/server.go:913
		}
//line /usr/local/go/src/net/http/server.go:913
		// _ = "end of CoverTab[42282]"
//line /usr/local/go/src/net/http/server.go:913
		_go_fuzz_dep_.CoverTab[42283]++
								w.writeContinueMu.Unlock()
//line /usr/local/go/src/net/http/server.go:914
		// _ = "end of CoverTab[42283]"
	} else {
//line /usr/local/go/src/net/http/server.go:915
		_go_fuzz_dep_.CoverTab[42286]++
//line /usr/local/go/src/net/http/server.go:915
		// _ = "end of CoverTab[42286]"
//line /usr/local/go/src/net/http/server.go:915
	}
//line /usr/local/go/src/net/http/server.go:915
	// _ = "end of CoverTab[42275]"
//line /usr/local/go/src/net/http/server.go:915
	_go_fuzz_dep_.CoverTab[42276]++
							n, err = ecr.readCloser.Read(p)
							if err == io.EOF {
//line /usr/local/go/src/net/http/server.go:917
		_go_fuzz_dep_.CoverTab[42287]++
								ecr.sawEOF.Store(true)
//line /usr/local/go/src/net/http/server.go:918
		// _ = "end of CoverTab[42287]"
	} else {
//line /usr/local/go/src/net/http/server.go:919
		_go_fuzz_dep_.CoverTab[42288]++
//line /usr/local/go/src/net/http/server.go:919
		// _ = "end of CoverTab[42288]"
//line /usr/local/go/src/net/http/server.go:919
	}
//line /usr/local/go/src/net/http/server.go:919
	// _ = "end of CoverTab[42276]"
//line /usr/local/go/src/net/http/server.go:919
	_go_fuzz_dep_.CoverTab[42277]++
							return
//line /usr/local/go/src/net/http/server.go:920
	// _ = "end of CoverTab[42277]"
}

func (ecr *expectContinueReader) Close() error {
//line /usr/local/go/src/net/http/server.go:923
	_go_fuzz_dep_.CoverTab[42289]++
							ecr.closed.Store(true)
							return ecr.readCloser.Close()
//line /usr/local/go/src/net/http/server.go:925
	// _ = "end of CoverTab[42289]"
}

// TimeFormat is the time format to use when generating times in HTTP
//line /usr/local/go/src/net/http/server.go:928
// headers. It is like time.RFC1123 but hard-codes GMT as the time
//line /usr/local/go/src/net/http/server.go:928
// zone. The time being formatted must be in UTC for Format to
//line /usr/local/go/src/net/http/server.go:928
// generate the correct format.
//line /usr/local/go/src/net/http/server.go:928
//
//line /usr/local/go/src/net/http/server.go:928
// For parsing this time format, see ParseTime.
//line /usr/local/go/src/net/http/server.go:934
const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

// appendTime is a non-allocating version of []byte(t.UTC().Format(TimeFormat))
func appendTime(b []byte, t time.Time) []byte {
//line /usr/local/go/src/net/http/server.go:937
	_go_fuzz_dep_.CoverTab[42290]++
							const days = "SunMonTueWedThuFriSat"
							const months = "JanFebMarAprMayJunJulAugSepOctNovDec"

							t = t.UTC()
							yy, mm, dd := t.Date()
							hh, mn, ss := t.Clock()
							day := days[3*t.Weekday():]
							mon := months[3*(mm-1):]

							return append(b,
		day[0], day[1], day[2], ',', ' ',
		byte('0'+dd/10), byte('0'+dd%10), ' ',
		mon[0], mon[1], mon[2], ' ',
		byte('0'+yy/1000), byte('0'+(yy/100)%10), byte('0'+(yy/10)%10), byte('0'+yy%10), ' ',
		byte('0'+hh/10), byte('0'+hh%10), ':',
		byte('0'+mn/10), byte('0'+mn%10), ':',
		byte('0'+ss/10), byte('0'+ss%10), ' ',
		'G', 'M', 'T')
//line /usr/local/go/src/net/http/server.go:955
	// _ = "end of CoverTab[42290]"
}

var errTooLarge = errors.New("http: request too large")

// Read next request from connection.
func (c *conn) readRequest(ctx context.Context) (w *response, err error) {
//line /usr/local/go/src/net/http/server.go:961
	_go_fuzz_dep_.CoverTab[42291]++
							if c.hijacked() {
//line /usr/local/go/src/net/http/server.go:962
		_go_fuzz_dep_.CoverTab[42305]++
								return nil, ErrHijacked
//line /usr/local/go/src/net/http/server.go:963
		// _ = "end of CoverTab[42305]"
	} else {
//line /usr/local/go/src/net/http/server.go:964
		_go_fuzz_dep_.CoverTab[42306]++
//line /usr/local/go/src/net/http/server.go:964
		// _ = "end of CoverTab[42306]"
//line /usr/local/go/src/net/http/server.go:964
	}
//line /usr/local/go/src/net/http/server.go:964
	// _ = "end of CoverTab[42291]"
//line /usr/local/go/src/net/http/server.go:964
	_go_fuzz_dep_.CoverTab[42292]++

							var (
		wholeReqDeadline	time.Time	// or zero if none
		hdrDeadline		time.Time	// or zero if none
	)
	t0 := time.Now()
	if d := c.server.readHeaderTimeout(); d > 0 {
//line /usr/local/go/src/net/http/server.go:971
		_go_fuzz_dep_.CoverTab[42307]++
								hdrDeadline = t0.Add(d)
//line /usr/local/go/src/net/http/server.go:972
		// _ = "end of CoverTab[42307]"
	} else {
//line /usr/local/go/src/net/http/server.go:973
		_go_fuzz_dep_.CoverTab[42308]++
//line /usr/local/go/src/net/http/server.go:973
		// _ = "end of CoverTab[42308]"
//line /usr/local/go/src/net/http/server.go:973
	}
//line /usr/local/go/src/net/http/server.go:973
	// _ = "end of CoverTab[42292]"
//line /usr/local/go/src/net/http/server.go:973
	_go_fuzz_dep_.CoverTab[42293]++
							if d := c.server.ReadTimeout; d > 0 {
//line /usr/local/go/src/net/http/server.go:974
		_go_fuzz_dep_.CoverTab[42309]++
								wholeReqDeadline = t0.Add(d)
//line /usr/local/go/src/net/http/server.go:975
		// _ = "end of CoverTab[42309]"
	} else {
//line /usr/local/go/src/net/http/server.go:976
		_go_fuzz_dep_.CoverTab[42310]++
//line /usr/local/go/src/net/http/server.go:976
		// _ = "end of CoverTab[42310]"
//line /usr/local/go/src/net/http/server.go:976
	}
//line /usr/local/go/src/net/http/server.go:976
	// _ = "end of CoverTab[42293]"
//line /usr/local/go/src/net/http/server.go:976
	_go_fuzz_dep_.CoverTab[42294]++
							c.rwc.SetReadDeadline(hdrDeadline)
							if d := c.server.WriteTimeout; d > 0 {
//line /usr/local/go/src/net/http/server.go:978
		_go_fuzz_dep_.CoverTab[42311]++
								defer func() {
//line /usr/local/go/src/net/http/server.go:979
			_go_fuzz_dep_.CoverTab[42312]++
									c.rwc.SetWriteDeadline(time.Now().Add(d))
//line /usr/local/go/src/net/http/server.go:980
			// _ = "end of CoverTab[42312]"
		}()
//line /usr/local/go/src/net/http/server.go:981
		// _ = "end of CoverTab[42311]"
	} else {
//line /usr/local/go/src/net/http/server.go:982
		_go_fuzz_dep_.CoverTab[42313]++
//line /usr/local/go/src/net/http/server.go:982
		// _ = "end of CoverTab[42313]"
//line /usr/local/go/src/net/http/server.go:982
	}
//line /usr/local/go/src/net/http/server.go:982
	// _ = "end of CoverTab[42294]"
//line /usr/local/go/src/net/http/server.go:982
	_go_fuzz_dep_.CoverTab[42295]++

							c.r.setReadLimit(c.server.initialReadLimitSize())
							if c.lastMethod == "POST" {
//line /usr/local/go/src/net/http/server.go:985
		_go_fuzz_dep_.CoverTab[42314]++

								peek, _ := c.bufr.Peek(4)
								c.bufr.Discard(numLeadingCRorLF(peek))
//line /usr/local/go/src/net/http/server.go:988
		// _ = "end of CoverTab[42314]"
	} else {
//line /usr/local/go/src/net/http/server.go:989
		_go_fuzz_dep_.CoverTab[42315]++
//line /usr/local/go/src/net/http/server.go:989
		// _ = "end of CoverTab[42315]"
//line /usr/local/go/src/net/http/server.go:989
	}
//line /usr/local/go/src/net/http/server.go:989
	// _ = "end of CoverTab[42295]"
//line /usr/local/go/src/net/http/server.go:989
	_go_fuzz_dep_.CoverTab[42296]++
							req, err := readRequest(c.bufr)
							if err != nil {
//line /usr/local/go/src/net/http/server.go:991
		_go_fuzz_dep_.CoverTab[42316]++
								if c.r.hitReadLimit() {
//line /usr/local/go/src/net/http/server.go:992
			_go_fuzz_dep_.CoverTab[42318]++
									return nil, errTooLarge
//line /usr/local/go/src/net/http/server.go:993
			// _ = "end of CoverTab[42318]"
		} else {
//line /usr/local/go/src/net/http/server.go:994
			_go_fuzz_dep_.CoverTab[42319]++
//line /usr/local/go/src/net/http/server.go:994
			// _ = "end of CoverTab[42319]"
//line /usr/local/go/src/net/http/server.go:994
		}
//line /usr/local/go/src/net/http/server.go:994
		// _ = "end of CoverTab[42316]"
//line /usr/local/go/src/net/http/server.go:994
		_go_fuzz_dep_.CoverTab[42317]++
								return nil, err
//line /usr/local/go/src/net/http/server.go:995
		// _ = "end of CoverTab[42317]"
	} else {
//line /usr/local/go/src/net/http/server.go:996
		_go_fuzz_dep_.CoverTab[42320]++
//line /usr/local/go/src/net/http/server.go:996
		// _ = "end of CoverTab[42320]"
//line /usr/local/go/src/net/http/server.go:996
	}
//line /usr/local/go/src/net/http/server.go:996
	// _ = "end of CoverTab[42296]"
//line /usr/local/go/src/net/http/server.go:996
	_go_fuzz_dep_.CoverTab[42297]++

							if !http1ServerSupportsRequest(req) {
//line /usr/local/go/src/net/http/server.go:998
		_go_fuzz_dep_.CoverTab[42321]++
								return nil, statusError{StatusHTTPVersionNotSupported, "unsupported protocol version"}
//line /usr/local/go/src/net/http/server.go:999
		// _ = "end of CoverTab[42321]"
	} else {
//line /usr/local/go/src/net/http/server.go:1000
		_go_fuzz_dep_.CoverTab[42322]++
//line /usr/local/go/src/net/http/server.go:1000
		// _ = "end of CoverTab[42322]"
//line /usr/local/go/src/net/http/server.go:1000
	}
//line /usr/local/go/src/net/http/server.go:1000
	// _ = "end of CoverTab[42297]"
//line /usr/local/go/src/net/http/server.go:1000
	_go_fuzz_dep_.CoverTab[42298]++

							c.lastMethod = req.Method
							c.r.setInfiniteReadLimit()

							hosts, haveHost := req.Header["Host"]
							isH2Upgrade := req.isH2Upgrade()
							if req.ProtoAtLeast(1, 1) && func() bool {
//line /usr/local/go/src/net/http/server.go:1007
		_go_fuzz_dep_.CoverTab[42323]++
//line /usr/local/go/src/net/http/server.go:1007
		return (!haveHost || func() bool {
//line /usr/local/go/src/net/http/server.go:1007
			_go_fuzz_dep_.CoverTab[42324]++
//line /usr/local/go/src/net/http/server.go:1007
			return len(hosts) == 0
//line /usr/local/go/src/net/http/server.go:1007
			// _ = "end of CoverTab[42324]"
//line /usr/local/go/src/net/http/server.go:1007
		}())
//line /usr/local/go/src/net/http/server.go:1007
		// _ = "end of CoverTab[42323]"
//line /usr/local/go/src/net/http/server.go:1007
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1007
		_go_fuzz_dep_.CoverTab[42325]++
//line /usr/local/go/src/net/http/server.go:1007
		return !isH2Upgrade
//line /usr/local/go/src/net/http/server.go:1007
		// _ = "end of CoverTab[42325]"
//line /usr/local/go/src/net/http/server.go:1007
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1007
		_go_fuzz_dep_.CoverTab[42326]++
//line /usr/local/go/src/net/http/server.go:1007
		return req.Method != "CONNECT"
//line /usr/local/go/src/net/http/server.go:1007
		// _ = "end of CoverTab[42326]"
//line /usr/local/go/src/net/http/server.go:1007
	}() {
//line /usr/local/go/src/net/http/server.go:1007
		_go_fuzz_dep_.CoverTab[42327]++
								return nil, badRequestError("missing required Host header")
//line /usr/local/go/src/net/http/server.go:1008
		// _ = "end of CoverTab[42327]"
	} else {
//line /usr/local/go/src/net/http/server.go:1009
		_go_fuzz_dep_.CoverTab[42328]++
//line /usr/local/go/src/net/http/server.go:1009
		// _ = "end of CoverTab[42328]"
//line /usr/local/go/src/net/http/server.go:1009
	}
//line /usr/local/go/src/net/http/server.go:1009
	// _ = "end of CoverTab[42298]"
//line /usr/local/go/src/net/http/server.go:1009
	_go_fuzz_dep_.CoverTab[42299]++
							if len(hosts) == 1 && func() bool {
//line /usr/local/go/src/net/http/server.go:1010
		_go_fuzz_dep_.CoverTab[42329]++
//line /usr/local/go/src/net/http/server.go:1010
		return !httpguts.ValidHostHeader(hosts[0])
//line /usr/local/go/src/net/http/server.go:1010
		// _ = "end of CoverTab[42329]"
//line /usr/local/go/src/net/http/server.go:1010
	}() {
//line /usr/local/go/src/net/http/server.go:1010
		_go_fuzz_dep_.CoverTab[42330]++
								return nil, badRequestError("malformed Host header")
//line /usr/local/go/src/net/http/server.go:1011
		// _ = "end of CoverTab[42330]"
	} else {
//line /usr/local/go/src/net/http/server.go:1012
		_go_fuzz_dep_.CoverTab[42331]++
//line /usr/local/go/src/net/http/server.go:1012
		// _ = "end of CoverTab[42331]"
//line /usr/local/go/src/net/http/server.go:1012
	}
//line /usr/local/go/src/net/http/server.go:1012
	// _ = "end of CoverTab[42299]"
//line /usr/local/go/src/net/http/server.go:1012
	_go_fuzz_dep_.CoverTab[42300]++
							for k, vv := range req.Header {
//line /usr/local/go/src/net/http/server.go:1013
		_go_fuzz_dep_.CoverTab[42332]++
								if !httpguts.ValidHeaderFieldName(k) {
//line /usr/local/go/src/net/http/server.go:1014
			_go_fuzz_dep_.CoverTab[42334]++
									return nil, badRequestError("invalid header name")
//line /usr/local/go/src/net/http/server.go:1015
			// _ = "end of CoverTab[42334]"
		} else {
//line /usr/local/go/src/net/http/server.go:1016
			_go_fuzz_dep_.CoverTab[42335]++
//line /usr/local/go/src/net/http/server.go:1016
			// _ = "end of CoverTab[42335]"
//line /usr/local/go/src/net/http/server.go:1016
		}
//line /usr/local/go/src/net/http/server.go:1016
		// _ = "end of CoverTab[42332]"
//line /usr/local/go/src/net/http/server.go:1016
		_go_fuzz_dep_.CoverTab[42333]++
								for _, v := range vv {
//line /usr/local/go/src/net/http/server.go:1017
			_go_fuzz_dep_.CoverTab[42336]++
									if !httpguts.ValidHeaderFieldValue(v) {
//line /usr/local/go/src/net/http/server.go:1018
				_go_fuzz_dep_.CoverTab[42337]++
										return nil, badRequestError("invalid header value")
//line /usr/local/go/src/net/http/server.go:1019
				// _ = "end of CoverTab[42337]"
			} else {
//line /usr/local/go/src/net/http/server.go:1020
				_go_fuzz_dep_.CoverTab[42338]++
//line /usr/local/go/src/net/http/server.go:1020
				// _ = "end of CoverTab[42338]"
//line /usr/local/go/src/net/http/server.go:1020
			}
//line /usr/local/go/src/net/http/server.go:1020
			// _ = "end of CoverTab[42336]"
		}
//line /usr/local/go/src/net/http/server.go:1021
		// _ = "end of CoverTab[42333]"
	}
//line /usr/local/go/src/net/http/server.go:1022
	// _ = "end of CoverTab[42300]"
//line /usr/local/go/src/net/http/server.go:1022
	_go_fuzz_dep_.CoverTab[42301]++
							delete(req.Header, "Host")

							ctx, cancelCtx := context.WithCancel(ctx)
							req.ctx = ctx
							req.RemoteAddr = c.remoteAddr
							req.TLS = c.tlsState
							if body, ok := req.Body.(*body); ok {
//line /usr/local/go/src/net/http/server.go:1029
		_go_fuzz_dep_.CoverTab[42339]++
								body.doEarlyClose = true
//line /usr/local/go/src/net/http/server.go:1030
		// _ = "end of CoverTab[42339]"
	} else {
//line /usr/local/go/src/net/http/server.go:1031
		_go_fuzz_dep_.CoverTab[42340]++
//line /usr/local/go/src/net/http/server.go:1031
		// _ = "end of CoverTab[42340]"
//line /usr/local/go/src/net/http/server.go:1031
	}
//line /usr/local/go/src/net/http/server.go:1031
	// _ = "end of CoverTab[42301]"
//line /usr/local/go/src/net/http/server.go:1031
	_go_fuzz_dep_.CoverTab[42302]++

//line /usr/local/go/src/net/http/server.go:1034
	if !hdrDeadline.Equal(wholeReqDeadline) {
//line /usr/local/go/src/net/http/server.go:1034
		_go_fuzz_dep_.CoverTab[42341]++
								c.rwc.SetReadDeadline(wholeReqDeadline)
//line /usr/local/go/src/net/http/server.go:1035
		// _ = "end of CoverTab[42341]"
	} else {
//line /usr/local/go/src/net/http/server.go:1036
		_go_fuzz_dep_.CoverTab[42342]++
//line /usr/local/go/src/net/http/server.go:1036
		// _ = "end of CoverTab[42342]"
//line /usr/local/go/src/net/http/server.go:1036
	}
//line /usr/local/go/src/net/http/server.go:1036
	// _ = "end of CoverTab[42302]"
//line /usr/local/go/src/net/http/server.go:1036
	_go_fuzz_dep_.CoverTab[42303]++

							w = &response{
								conn:		c,
								cancelCtx:	cancelCtx,
								req:		req,
								reqBody:	req.Body,
								handlerHeader:	make(Header),
								contentLength:	-1,
								closeNotifyCh:	make(chan bool, 1),

//line /usr/local/go/src/net/http/server.go:1050
		wants10KeepAlive:	req.wantsHttp10KeepAlive(),
								wantsClose:		req.wantsClose(),
	}
	if isH2Upgrade {
//line /usr/local/go/src/net/http/server.go:1053
		_go_fuzz_dep_.CoverTab[42343]++
								w.closeAfterReply = true
//line /usr/local/go/src/net/http/server.go:1054
		// _ = "end of CoverTab[42343]"
	} else {
//line /usr/local/go/src/net/http/server.go:1055
		_go_fuzz_dep_.CoverTab[42344]++
//line /usr/local/go/src/net/http/server.go:1055
		// _ = "end of CoverTab[42344]"
//line /usr/local/go/src/net/http/server.go:1055
	}
//line /usr/local/go/src/net/http/server.go:1055
	// _ = "end of CoverTab[42303]"
//line /usr/local/go/src/net/http/server.go:1055
	_go_fuzz_dep_.CoverTab[42304]++
							w.cw.res = w
							w.w = newBufioWriterSize(&w.cw, bufferBeforeChunkingSize)
							return w, nil
//line /usr/local/go/src/net/http/server.go:1058
	// _ = "end of CoverTab[42304]"
}

// http1ServerSupportsRequest reports whether Go's HTTP/1.x server
//line /usr/local/go/src/net/http/server.go:1061
// supports the given request.
//line /usr/local/go/src/net/http/server.go:1063
func http1ServerSupportsRequest(req *Request) bool {
//line /usr/local/go/src/net/http/server.go:1063
	_go_fuzz_dep_.CoverTab[42345]++
							if req.ProtoMajor == 1 {
//line /usr/local/go/src/net/http/server.go:1064
		_go_fuzz_dep_.CoverTab[42348]++
								return true
//line /usr/local/go/src/net/http/server.go:1065
		// _ = "end of CoverTab[42348]"
	} else {
//line /usr/local/go/src/net/http/server.go:1066
		_go_fuzz_dep_.CoverTab[42349]++
//line /usr/local/go/src/net/http/server.go:1066
		// _ = "end of CoverTab[42349]"
//line /usr/local/go/src/net/http/server.go:1066
	}
//line /usr/local/go/src/net/http/server.go:1066
	// _ = "end of CoverTab[42345]"
//line /usr/local/go/src/net/http/server.go:1066
	_go_fuzz_dep_.CoverTab[42346]++

//line /usr/local/go/src/net/http/server.go:1069
	if req.ProtoMajor == 2 && func() bool {
//line /usr/local/go/src/net/http/server.go:1069
		_go_fuzz_dep_.CoverTab[42350]++
//line /usr/local/go/src/net/http/server.go:1069
		return req.ProtoMinor == 0
//line /usr/local/go/src/net/http/server.go:1069
		// _ = "end of CoverTab[42350]"
//line /usr/local/go/src/net/http/server.go:1069
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1069
		_go_fuzz_dep_.CoverTab[42351]++
//line /usr/local/go/src/net/http/server.go:1069
		return req.Method == "PRI"
								// _ = "end of CoverTab[42351]"
//line /usr/local/go/src/net/http/server.go:1070
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1070
		_go_fuzz_dep_.CoverTab[42352]++
//line /usr/local/go/src/net/http/server.go:1070
		return req.RequestURI == "*"
//line /usr/local/go/src/net/http/server.go:1070
		// _ = "end of CoverTab[42352]"
//line /usr/local/go/src/net/http/server.go:1070
	}() {
//line /usr/local/go/src/net/http/server.go:1070
		_go_fuzz_dep_.CoverTab[42353]++
								return true
//line /usr/local/go/src/net/http/server.go:1071
		// _ = "end of CoverTab[42353]"
	} else {
//line /usr/local/go/src/net/http/server.go:1072
		_go_fuzz_dep_.CoverTab[42354]++
//line /usr/local/go/src/net/http/server.go:1072
		// _ = "end of CoverTab[42354]"
//line /usr/local/go/src/net/http/server.go:1072
	}
//line /usr/local/go/src/net/http/server.go:1072
	// _ = "end of CoverTab[42346]"
//line /usr/local/go/src/net/http/server.go:1072
	_go_fuzz_dep_.CoverTab[42347]++

//line /usr/local/go/src/net/http/server.go:1075
	return false
//line /usr/local/go/src/net/http/server.go:1075
	// _ = "end of CoverTab[42347]"
}

func (w *response) Header() Header {
//line /usr/local/go/src/net/http/server.go:1078
	_go_fuzz_dep_.CoverTab[42355]++
							if w.cw.header == nil && func() bool {
//line /usr/local/go/src/net/http/server.go:1079
		_go_fuzz_dep_.CoverTab[42357]++
//line /usr/local/go/src/net/http/server.go:1079
		return w.wroteHeader
//line /usr/local/go/src/net/http/server.go:1079
		// _ = "end of CoverTab[42357]"
//line /usr/local/go/src/net/http/server.go:1079
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1079
		_go_fuzz_dep_.CoverTab[42358]++
//line /usr/local/go/src/net/http/server.go:1079
		return !w.cw.wroteHeader
//line /usr/local/go/src/net/http/server.go:1079
		// _ = "end of CoverTab[42358]"
//line /usr/local/go/src/net/http/server.go:1079
	}() {
//line /usr/local/go/src/net/http/server.go:1079
		_go_fuzz_dep_.CoverTab[42359]++

//line /usr/local/go/src/net/http/server.go:1083
		w.cw.header = w.handlerHeader.Clone()
//line /usr/local/go/src/net/http/server.go:1083
		// _ = "end of CoverTab[42359]"
	} else {
//line /usr/local/go/src/net/http/server.go:1084
		_go_fuzz_dep_.CoverTab[42360]++
//line /usr/local/go/src/net/http/server.go:1084
		// _ = "end of CoverTab[42360]"
//line /usr/local/go/src/net/http/server.go:1084
	}
//line /usr/local/go/src/net/http/server.go:1084
	// _ = "end of CoverTab[42355]"
//line /usr/local/go/src/net/http/server.go:1084
	_go_fuzz_dep_.CoverTab[42356]++
							w.calledHeader = true
							return w.handlerHeader
//line /usr/local/go/src/net/http/server.go:1086
	// _ = "end of CoverTab[42356]"
}

// maxPostHandlerReadBytes is the max number of Request.Body bytes not
//line /usr/local/go/src/net/http/server.go:1089
// consumed by a handler that the server will read from the client
//line /usr/local/go/src/net/http/server.go:1089
// in order to keep a connection alive. If there are more bytes than
//line /usr/local/go/src/net/http/server.go:1089
// this then the server to be paranoid instead sends a "Connection:
//line /usr/local/go/src/net/http/server.go:1089
// close" response.
//line /usr/local/go/src/net/http/server.go:1089
//
//line /usr/local/go/src/net/http/server.go:1089
// This number is approximately what a typical machine's TCP buffer
//line /usr/local/go/src/net/http/server.go:1089
// size is anyway.  (if we have the bytes on the machine, we might as
//line /usr/local/go/src/net/http/server.go:1089
// well read them)
//line /usr/local/go/src/net/http/server.go:1098
const maxPostHandlerReadBytes = 256 << 10

func checkWriteHeaderCode(code int) {
//line /usr/local/go/src/net/http/server.go:1100
	_go_fuzz_dep_.CoverTab[42361]++

//line /usr/local/go/src/net/http/server.go:1111
	if code < 100 || func() bool {
//line /usr/local/go/src/net/http/server.go:1111
		_go_fuzz_dep_.CoverTab[42362]++
//line /usr/local/go/src/net/http/server.go:1111
		return code > 999
//line /usr/local/go/src/net/http/server.go:1111
		// _ = "end of CoverTab[42362]"
//line /usr/local/go/src/net/http/server.go:1111
	}() {
//line /usr/local/go/src/net/http/server.go:1111
		_go_fuzz_dep_.CoverTab[42363]++
								panic(fmt.Sprintf("invalid WriteHeader code %v", code))
//line /usr/local/go/src/net/http/server.go:1112
		// _ = "end of CoverTab[42363]"
	} else {
//line /usr/local/go/src/net/http/server.go:1113
		_go_fuzz_dep_.CoverTab[42364]++
//line /usr/local/go/src/net/http/server.go:1113
		// _ = "end of CoverTab[42364]"
//line /usr/local/go/src/net/http/server.go:1113
	}
//line /usr/local/go/src/net/http/server.go:1113
	// _ = "end of CoverTab[42361]"
}

// relevantCaller searches the call stack for the first function outside of net/http.
//line /usr/local/go/src/net/http/server.go:1116
// The purpose of this function is to provide more helpful error messages.
//line /usr/local/go/src/net/http/server.go:1118
func relevantCaller() runtime.Frame {
//line /usr/local/go/src/net/http/server.go:1118
	_go_fuzz_dep_.CoverTab[42365]++
							pc := make([]uintptr, 16)
							n := runtime.Callers(1, pc)
							frames := runtime.CallersFrames(pc[:n])
							var frame runtime.Frame
							for {
//line /usr/local/go/src/net/http/server.go:1123
		_go_fuzz_dep_.CoverTab[42367]++
								frame, more := frames.Next()
								if !strings.HasPrefix(frame.Function, "net/http.") {
//line /usr/local/go/src/net/http/server.go:1125
			_go_fuzz_dep_.CoverTab[42369]++
									return frame
//line /usr/local/go/src/net/http/server.go:1126
			// _ = "end of CoverTab[42369]"
		} else {
//line /usr/local/go/src/net/http/server.go:1127
			_go_fuzz_dep_.CoverTab[42370]++
//line /usr/local/go/src/net/http/server.go:1127
			// _ = "end of CoverTab[42370]"
//line /usr/local/go/src/net/http/server.go:1127
		}
//line /usr/local/go/src/net/http/server.go:1127
		// _ = "end of CoverTab[42367]"
//line /usr/local/go/src/net/http/server.go:1127
		_go_fuzz_dep_.CoverTab[42368]++
								if !more {
//line /usr/local/go/src/net/http/server.go:1128
			_go_fuzz_dep_.CoverTab[42371]++
									break
//line /usr/local/go/src/net/http/server.go:1129
			// _ = "end of CoverTab[42371]"
		} else {
//line /usr/local/go/src/net/http/server.go:1130
			_go_fuzz_dep_.CoverTab[42372]++
//line /usr/local/go/src/net/http/server.go:1130
			// _ = "end of CoverTab[42372]"
//line /usr/local/go/src/net/http/server.go:1130
		}
//line /usr/local/go/src/net/http/server.go:1130
		// _ = "end of CoverTab[42368]"
	}
//line /usr/local/go/src/net/http/server.go:1131
	// _ = "end of CoverTab[42365]"
//line /usr/local/go/src/net/http/server.go:1131
	_go_fuzz_dep_.CoverTab[42366]++
							return frame
//line /usr/local/go/src/net/http/server.go:1132
	// _ = "end of CoverTab[42366]"
}

func (w *response) WriteHeader(code int) {
//line /usr/local/go/src/net/http/server.go:1135
	_go_fuzz_dep_.CoverTab[42373]++
							if w.conn.hijacked() {
//line /usr/local/go/src/net/http/server.go:1136
		_go_fuzz_dep_.CoverTab[42378]++
								caller := relevantCaller()
								w.conn.server.logf("http: response.WriteHeader on hijacked connection from %s (%s:%d)", caller.Function, path.Base(caller.File), caller.Line)
								return
//line /usr/local/go/src/net/http/server.go:1139
		// _ = "end of CoverTab[42378]"
	} else {
//line /usr/local/go/src/net/http/server.go:1140
		_go_fuzz_dep_.CoverTab[42379]++
//line /usr/local/go/src/net/http/server.go:1140
		// _ = "end of CoverTab[42379]"
//line /usr/local/go/src/net/http/server.go:1140
	}
//line /usr/local/go/src/net/http/server.go:1140
	// _ = "end of CoverTab[42373]"
//line /usr/local/go/src/net/http/server.go:1140
	_go_fuzz_dep_.CoverTab[42374]++
							if w.wroteHeader {
//line /usr/local/go/src/net/http/server.go:1141
		_go_fuzz_dep_.CoverTab[42380]++
								caller := relevantCaller()
								w.conn.server.logf("http: superfluous response.WriteHeader call from %s (%s:%d)", caller.Function, path.Base(caller.File), caller.Line)
								return
//line /usr/local/go/src/net/http/server.go:1144
		// _ = "end of CoverTab[42380]"
	} else {
//line /usr/local/go/src/net/http/server.go:1145
		_go_fuzz_dep_.CoverTab[42381]++
//line /usr/local/go/src/net/http/server.go:1145
		// _ = "end of CoverTab[42381]"
//line /usr/local/go/src/net/http/server.go:1145
	}
//line /usr/local/go/src/net/http/server.go:1145
	// _ = "end of CoverTab[42374]"
//line /usr/local/go/src/net/http/server.go:1145
	_go_fuzz_dep_.CoverTab[42375]++
							checkWriteHeaderCode(code)

//line /usr/local/go/src/net/http/server.go:1149
	if code >= 100 && func() bool {
//line /usr/local/go/src/net/http/server.go:1149
		_go_fuzz_dep_.CoverTab[42382]++
//line /usr/local/go/src/net/http/server.go:1149
		return code <= 199
//line /usr/local/go/src/net/http/server.go:1149
		// _ = "end of CoverTab[42382]"
//line /usr/local/go/src/net/http/server.go:1149
	}() {
//line /usr/local/go/src/net/http/server.go:1149
		_go_fuzz_dep_.CoverTab[42383]++

								if code == 100 && func() bool {
//line /usr/local/go/src/net/http/server.go:1151
			_go_fuzz_dep_.CoverTab[42385]++
//line /usr/local/go/src/net/http/server.go:1151
			return w.canWriteContinue.Load()
//line /usr/local/go/src/net/http/server.go:1151
			// _ = "end of CoverTab[42385]"
//line /usr/local/go/src/net/http/server.go:1151
		}() {
//line /usr/local/go/src/net/http/server.go:1151
			_go_fuzz_dep_.CoverTab[42386]++
									w.writeContinueMu.Lock()
									w.canWriteContinue.Store(false)
									w.writeContinueMu.Unlock()
//line /usr/local/go/src/net/http/server.go:1154
			// _ = "end of CoverTab[42386]"
		} else {
//line /usr/local/go/src/net/http/server.go:1155
			_go_fuzz_dep_.CoverTab[42387]++
//line /usr/local/go/src/net/http/server.go:1155
			// _ = "end of CoverTab[42387]"
//line /usr/local/go/src/net/http/server.go:1155
		}
//line /usr/local/go/src/net/http/server.go:1155
		// _ = "end of CoverTab[42383]"
//line /usr/local/go/src/net/http/server.go:1155
		_go_fuzz_dep_.CoverTab[42384]++

								writeStatusLine(w.conn.bufw, w.req.ProtoAtLeast(1, 1), code, w.statusBuf[:])

//line /usr/local/go/src/net/http/server.go:1160
		w.handlerHeader.WriteSubset(w.conn.bufw, excludedHeadersNoBody)
								w.conn.bufw.Write(crlf)
								w.conn.bufw.Flush()

								return
//line /usr/local/go/src/net/http/server.go:1164
		// _ = "end of CoverTab[42384]"
	} else {
//line /usr/local/go/src/net/http/server.go:1165
		_go_fuzz_dep_.CoverTab[42388]++
//line /usr/local/go/src/net/http/server.go:1165
		// _ = "end of CoverTab[42388]"
//line /usr/local/go/src/net/http/server.go:1165
	}
//line /usr/local/go/src/net/http/server.go:1165
	// _ = "end of CoverTab[42375]"
//line /usr/local/go/src/net/http/server.go:1165
	_go_fuzz_dep_.CoverTab[42376]++

							w.wroteHeader = true
							w.status = code

							if w.calledHeader && func() bool {
//line /usr/local/go/src/net/http/server.go:1170
		_go_fuzz_dep_.CoverTab[42389]++
//line /usr/local/go/src/net/http/server.go:1170
		return w.cw.header == nil
//line /usr/local/go/src/net/http/server.go:1170
		// _ = "end of CoverTab[42389]"
//line /usr/local/go/src/net/http/server.go:1170
	}() {
//line /usr/local/go/src/net/http/server.go:1170
		_go_fuzz_dep_.CoverTab[42390]++
								w.cw.header = w.handlerHeader.Clone()
//line /usr/local/go/src/net/http/server.go:1171
		// _ = "end of CoverTab[42390]"
	} else {
//line /usr/local/go/src/net/http/server.go:1172
		_go_fuzz_dep_.CoverTab[42391]++
//line /usr/local/go/src/net/http/server.go:1172
		// _ = "end of CoverTab[42391]"
//line /usr/local/go/src/net/http/server.go:1172
	}
//line /usr/local/go/src/net/http/server.go:1172
	// _ = "end of CoverTab[42376]"
//line /usr/local/go/src/net/http/server.go:1172
	_go_fuzz_dep_.CoverTab[42377]++

							if cl := w.handlerHeader.get("Content-Length"); cl != "" {
//line /usr/local/go/src/net/http/server.go:1174
		_go_fuzz_dep_.CoverTab[42392]++
								v, err := strconv.ParseInt(cl, 10, 64)
								if err == nil && func() bool {
//line /usr/local/go/src/net/http/server.go:1176
			_go_fuzz_dep_.CoverTab[42393]++
//line /usr/local/go/src/net/http/server.go:1176
			return v >= 0
//line /usr/local/go/src/net/http/server.go:1176
			// _ = "end of CoverTab[42393]"
//line /usr/local/go/src/net/http/server.go:1176
		}() {
//line /usr/local/go/src/net/http/server.go:1176
			_go_fuzz_dep_.CoverTab[42394]++
									w.contentLength = v
//line /usr/local/go/src/net/http/server.go:1177
			// _ = "end of CoverTab[42394]"
		} else {
//line /usr/local/go/src/net/http/server.go:1178
			_go_fuzz_dep_.CoverTab[42395]++
									w.conn.server.logf("http: invalid Content-Length of %q", cl)
									w.handlerHeader.Del("Content-Length")
//line /usr/local/go/src/net/http/server.go:1180
			// _ = "end of CoverTab[42395]"
		}
//line /usr/local/go/src/net/http/server.go:1181
		// _ = "end of CoverTab[42392]"
	} else {
//line /usr/local/go/src/net/http/server.go:1182
		_go_fuzz_dep_.CoverTab[42396]++
//line /usr/local/go/src/net/http/server.go:1182
		// _ = "end of CoverTab[42396]"
//line /usr/local/go/src/net/http/server.go:1182
	}
//line /usr/local/go/src/net/http/server.go:1182
	// _ = "end of CoverTab[42377]"
}

// extraHeader is the set of headers sometimes added by chunkWriter.writeHeader.
//line /usr/local/go/src/net/http/server.go:1185
// This type is used to avoid extra allocations from cloning and/or populating
//line /usr/local/go/src/net/http/server.go:1185
// the response Header map and all its 1-element slices.
//line /usr/local/go/src/net/http/server.go:1188
type extraHeader struct {
	contentType		string
	connection		string
	transferEncoding	string
	date			[]byte	// written if not nil
	contentLength		[]byte	// written if not nil
}

// Sorted the same as extraHeader.Write's loop.
var extraHeaderKeys = [][]byte{
	[]byte("Content-Type"),
	[]byte("Connection"),
	[]byte("Transfer-Encoding"),
}

var (
	headerContentLength	= []byte("Content-Length: ")
	headerDate		= []byte("Date: ")
)

// Write writes the headers described in h to w.
//line /usr/local/go/src/net/http/server.go:1208
//
//line /usr/local/go/src/net/http/server.go:1208
// This method has a value receiver, despite the somewhat large size
//line /usr/local/go/src/net/http/server.go:1208
// of h, because it prevents an allocation. The escape analysis isn't
//line /usr/local/go/src/net/http/server.go:1208
// smart enough to realize this function doesn't mutate h.
//line /usr/local/go/src/net/http/server.go:1213
func (h extraHeader) Write(w *bufio.Writer) {
//line /usr/local/go/src/net/http/server.go:1213
	_go_fuzz_dep_.CoverTab[42397]++
							if h.date != nil {
//line /usr/local/go/src/net/http/server.go:1214
		_go_fuzz_dep_.CoverTab[42400]++
								w.Write(headerDate)
								w.Write(h.date)
								w.Write(crlf)
//line /usr/local/go/src/net/http/server.go:1217
		// _ = "end of CoverTab[42400]"
	} else {
//line /usr/local/go/src/net/http/server.go:1218
		_go_fuzz_dep_.CoverTab[42401]++
//line /usr/local/go/src/net/http/server.go:1218
		// _ = "end of CoverTab[42401]"
//line /usr/local/go/src/net/http/server.go:1218
	}
//line /usr/local/go/src/net/http/server.go:1218
	// _ = "end of CoverTab[42397]"
//line /usr/local/go/src/net/http/server.go:1218
	_go_fuzz_dep_.CoverTab[42398]++
							if h.contentLength != nil {
//line /usr/local/go/src/net/http/server.go:1219
		_go_fuzz_dep_.CoverTab[42402]++
								w.Write(headerContentLength)
								w.Write(h.contentLength)
								w.Write(crlf)
//line /usr/local/go/src/net/http/server.go:1222
		// _ = "end of CoverTab[42402]"
	} else {
//line /usr/local/go/src/net/http/server.go:1223
		_go_fuzz_dep_.CoverTab[42403]++
//line /usr/local/go/src/net/http/server.go:1223
		// _ = "end of CoverTab[42403]"
//line /usr/local/go/src/net/http/server.go:1223
	}
//line /usr/local/go/src/net/http/server.go:1223
	// _ = "end of CoverTab[42398]"
//line /usr/local/go/src/net/http/server.go:1223
	_go_fuzz_dep_.CoverTab[42399]++
							for i, v := range []string{h.contentType, h.connection, h.transferEncoding} {
//line /usr/local/go/src/net/http/server.go:1224
		_go_fuzz_dep_.CoverTab[42404]++
								if v != "" {
//line /usr/local/go/src/net/http/server.go:1225
			_go_fuzz_dep_.CoverTab[42405]++
									w.Write(extraHeaderKeys[i])
									w.Write(colonSpace)
									w.WriteString(v)
									w.Write(crlf)
//line /usr/local/go/src/net/http/server.go:1229
			// _ = "end of CoverTab[42405]"
		} else {
//line /usr/local/go/src/net/http/server.go:1230
			_go_fuzz_dep_.CoverTab[42406]++
//line /usr/local/go/src/net/http/server.go:1230
			// _ = "end of CoverTab[42406]"
//line /usr/local/go/src/net/http/server.go:1230
		}
//line /usr/local/go/src/net/http/server.go:1230
		// _ = "end of CoverTab[42404]"
	}
//line /usr/local/go/src/net/http/server.go:1231
	// _ = "end of CoverTab[42399]"
}

// writeHeader finalizes the header sent to the client and writes it
//line /usr/local/go/src/net/http/server.go:1234
// to cw.res.conn.bufw.
//line /usr/local/go/src/net/http/server.go:1234
//
//line /usr/local/go/src/net/http/server.go:1234
// p is not written by writeHeader, but is the first chunk of the body
//line /usr/local/go/src/net/http/server.go:1234
// that will be written. It is sniffed for a Content-Type if none is
//line /usr/local/go/src/net/http/server.go:1234
// set explicitly. It's also used to set the Content-Length, if the
//line /usr/local/go/src/net/http/server.go:1234
// total body size was small and the handler has already finished
//line /usr/local/go/src/net/http/server.go:1234
// running.
//line /usr/local/go/src/net/http/server.go:1242
func (cw *chunkWriter) writeHeader(p []byte) {
//line /usr/local/go/src/net/http/server.go:1242
	_go_fuzz_dep_.CoverTab[42407]++
							if cw.wroteHeader {
//line /usr/local/go/src/net/http/server.go:1243
		_go_fuzz_dep_.CoverTab[42426]++
								return
//line /usr/local/go/src/net/http/server.go:1244
		// _ = "end of CoverTab[42426]"
	} else {
//line /usr/local/go/src/net/http/server.go:1245
		_go_fuzz_dep_.CoverTab[42427]++
//line /usr/local/go/src/net/http/server.go:1245
		// _ = "end of CoverTab[42427]"
//line /usr/local/go/src/net/http/server.go:1245
	}
//line /usr/local/go/src/net/http/server.go:1245
	// _ = "end of CoverTab[42407]"
//line /usr/local/go/src/net/http/server.go:1245
	_go_fuzz_dep_.CoverTab[42408]++
							cw.wroteHeader = true

							w := cw.res
							keepAlivesEnabled := w.conn.server.doKeepAlives()
							isHEAD := w.req.Method == "HEAD"

//line /usr/local/go/src/net/http/server.go:1257
	header := cw.header
	owned := header != nil
	if !owned {
//line /usr/local/go/src/net/http/server.go:1259
		_go_fuzz_dep_.CoverTab[42428]++
								header = w.handlerHeader
//line /usr/local/go/src/net/http/server.go:1260
		// _ = "end of CoverTab[42428]"
	} else {
//line /usr/local/go/src/net/http/server.go:1261
		_go_fuzz_dep_.CoverTab[42429]++
//line /usr/local/go/src/net/http/server.go:1261
		// _ = "end of CoverTab[42429]"
//line /usr/local/go/src/net/http/server.go:1261
	}
//line /usr/local/go/src/net/http/server.go:1261
	// _ = "end of CoverTab[42408]"
//line /usr/local/go/src/net/http/server.go:1261
	_go_fuzz_dep_.CoverTab[42409]++
							var excludeHeader map[string]bool
							delHeader := func(key string) {
//line /usr/local/go/src/net/http/server.go:1263
		_go_fuzz_dep_.CoverTab[42430]++
								if owned {
//line /usr/local/go/src/net/http/server.go:1264
			_go_fuzz_dep_.CoverTab[42434]++
									header.Del(key)
									return
//line /usr/local/go/src/net/http/server.go:1266
			// _ = "end of CoverTab[42434]"
		} else {
//line /usr/local/go/src/net/http/server.go:1267
			_go_fuzz_dep_.CoverTab[42435]++
//line /usr/local/go/src/net/http/server.go:1267
			// _ = "end of CoverTab[42435]"
//line /usr/local/go/src/net/http/server.go:1267
		}
//line /usr/local/go/src/net/http/server.go:1267
		// _ = "end of CoverTab[42430]"
//line /usr/local/go/src/net/http/server.go:1267
		_go_fuzz_dep_.CoverTab[42431]++
								if _, ok := header[key]; !ok {
//line /usr/local/go/src/net/http/server.go:1268
			_go_fuzz_dep_.CoverTab[42436]++
									return
//line /usr/local/go/src/net/http/server.go:1269
			// _ = "end of CoverTab[42436]"
		} else {
//line /usr/local/go/src/net/http/server.go:1270
			_go_fuzz_dep_.CoverTab[42437]++
//line /usr/local/go/src/net/http/server.go:1270
			// _ = "end of CoverTab[42437]"
//line /usr/local/go/src/net/http/server.go:1270
		}
//line /usr/local/go/src/net/http/server.go:1270
		// _ = "end of CoverTab[42431]"
//line /usr/local/go/src/net/http/server.go:1270
		_go_fuzz_dep_.CoverTab[42432]++
								if excludeHeader == nil {
//line /usr/local/go/src/net/http/server.go:1271
			_go_fuzz_dep_.CoverTab[42438]++
									excludeHeader = make(map[string]bool)
//line /usr/local/go/src/net/http/server.go:1272
			// _ = "end of CoverTab[42438]"
		} else {
//line /usr/local/go/src/net/http/server.go:1273
			_go_fuzz_dep_.CoverTab[42439]++
//line /usr/local/go/src/net/http/server.go:1273
			// _ = "end of CoverTab[42439]"
//line /usr/local/go/src/net/http/server.go:1273
		}
//line /usr/local/go/src/net/http/server.go:1273
		// _ = "end of CoverTab[42432]"
//line /usr/local/go/src/net/http/server.go:1273
		_go_fuzz_dep_.CoverTab[42433]++
								excludeHeader[key] = true
//line /usr/local/go/src/net/http/server.go:1274
		// _ = "end of CoverTab[42433]"
	}
//line /usr/local/go/src/net/http/server.go:1275
	// _ = "end of CoverTab[42409]"
//line /usr/local/go/src/net/http/server.go:1275
	_go_fuzz_dep_.CoverTab[42410]++
							var setHeader extraHeader

//line /usr/local/go/src/net/http/server.go:1279
	trailers := false
	for k := range cw.header {
//line /usr/local/go/src/net/http/server.go:1280
		_go_fuzz_dep_.CoverTab[42440]++
								if strings.HasPrefix(k, TrailerPrefix) {
//line /usr/local/go/src/net/http/server.go:1281
			_go_fuzz_dep_.CoverTab[42441]++
									if excludeHeader == nil {
//line /usr/local/go/src/net/http/server.go:1282
				_go_fuzz_dep_.CoverTab[42443]++
										excludeHeader = make(map[string]bool)
//line /usr/local/go/src/net/http/server.go:1283
				// _ = "end of CoverTab[42443]"
			} else {
//line /usr/local/go/src/net/http/server.go:1284
				_go_fuzz_dep_.CoverTab[42444]++
//line /usr/local/go/src/net/http/server.go:1284
				// _ = "end of CoverTab[42444]"
//line /usr/local/go/src/net/http/server.go:1284
			}
//line /usr/local/go/src/net/http/server.go:1284
			// _ = "end of CoverTab[42441]"
//line /usr/local/go/src/net/http/server.go:1284
			_go_fuzz_dep_.CoverTab[42442]++
									excludeHeader[k] = true
									trailers = true
//line /usr/local/go/src/net/http/server.go:1286
			// _ = "end of CoverTab[42442]"
		} else {
//line /usr/local/go/src/net/http/server.go:1287
			_go_fuzz_dep_.CoverTab[42445]++
//line /usr/local/go/src/net/http/server.go:1287
			// _ = "end of CoverTab[42445]"
//line /usr/local/go/src/net/http/server.go:1287
		}
//line /usr/local/go/src/net/http/server.go:1287
		// _ = "end of CoverTab[42440]"
	}
//line /usr/local/go/src/net/http/server.go:1288
	// _ = "end of CoverTab[42410]"
//line /usr/local/go/src/net/http/server.go:1288
	_go_fuzz_dep_.CoverTab[42411]++
							for _, v := range cw.header["Trailer"] {
//line /usr/local/go/src/net/http/server.go:1289
		_go_fuzz_dep_.CoverTab[42446]++
								trailers = true
								foreachHeaderElement(v, cw.res.declareTrailer)
//line /usr/local/go/src/net/http/server.go:1291
		// _ = "end of CoverTab[42446]"
	}
//line /usr/local/go/src/net/http/server.go:1292
	// _ = "end of CoverTab[42411]"
//line /usr/local/go/src/net/http/server.go:1292
	_go_fuzz_dep_.CoverTab[42412]++

							te := header.get("Transfer-Encoding")
							hasTE := te != ""

//line /usr/local/go/src/net/http/server.go:1311
	if w.handlerDone.Load() && func() bool {
//line /usr/local/go/src/net/http/server.go:1311
		_go_fuzz_dep_.CoverTab[42447]++
//line /usr/local/go/src/net/http/server.go:1311
		return !trailers
//line /usr/local/go/src/net/http/server.go:1311
		// _ = "end of CoverTab[42447]"
//line /usr/local/go/src/net/http/server.go:1311
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1311
		_go_fuzz_dep_.CoverTab[42448]++
//line /usr/local/go/src/net/http/server.go:1311
		return !hasTE
//line /usr/local/go/src/net/http/server.go:1311
		// _ = "end of CoverTab[42448]"
//line /usr/local/go/src/net/http/server.go:1311
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1311
		_go_fuzz_dep_.CoverTab[42449]++
//line /usr/local/go/src/net/http/server.go:1311
		return bodyAllowedForStatus(w.status)
//line /usr/local/go/src/net/http/server.go:1311
		// _ = "end of CoverTab[42449]"
//line /usr/local/go/src/net/http/server.go:1311
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1311
		_go_fuzz_dep_.CoverTab[42450]++
//line /usr/local/go/src/net/http/server.go:1311
		return header.get("Content-Length") == ""
//line /usr/local/go/src/net/http/server.go:1311
		// _ = "end of CoverTab[42450]"
//line /usr/local/go/src/net/http/server.go:1311
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1311
		_go_fuzz_dep_.CoverTab[42451]++
//line /usr/local/go/src/net/http/server.go:1311
		return (!isHEAD || func() bool {
//line /usr/local/go/src/net/http/server.go:1311
			_go_fuzz_dep_.CoverTab[42452]++
//line /usr/local/go/src/net/http/server.go:1311
			return len(p) > 0
//line /usr/local/go/src/net/http/server.go:1311
			// _ = "end of CoverTab[42452]"
//line /usr/local/go/src/net/http/server.go:1311
		}())
//line /usr/local/go/src/net/http/server.go:1311
		// _ = "end of CoverTab[42451]"
//line /usr/local/go/src/net/http/server.go:1311
	}() {
//line /usr/local/go/src/net/http/server.go:1311
		_go_fuzz_dep_.CoverTab[42453]++
								w.contentLength = int64(len(p))
								setHeader.contentLength = strconv.AppendInt(cw.res.clenBuf[:0], int64(len(p)), 10)
//line /usr/local/go/src/net/http/server.go:1313
		// _ = "end of CoverTab[42453]"
	} else {
//line /usr/local/go/src/net/http/server.go:1314
		_go_fuzz_dep_.CoverTab[42454]++
//line /usr/local/go/src/net/http/server.go:1314
		// _ = "end of CoverTab[42454]"
//line /usr/local/go/src/net/http/server.go:1314
	}
//line /usr/local/go/src/net/http/server.go:1314
	// _ = "end of CoverTab[42412]"
//line /usr/local/go/src/net/http/server.go:1314
	_go_fuzz_dep_.CoverTab[42413]++

//line /usr/local/go/src/net/http/server.go:1318
	if w.wants10KeepAlive && func() bool {
//line /usr/local/go/src/net/http/server.go:1318
		_go_fuzz_dep_.CoverTab[42455]++
//line /usr/local/go/src/net/http/server.go:1318
		return keepAlivesEnabled
//line /usr/local/go/src/net/http/server.go:1318
		// _ = "end of CoverTab[42455]"
//line /usr/local/go/src/net/http/server.go:1318
	}() {
//line /usr/local/go/src/net/http/server.go:1318
		_go_fuzz_dep_.CoverTab[42456]++
								sentLength := header.get("Content-Length") != ""
								if sentLength && func() bool {
//line /usr/local/go/src/net/http/server.go:1320
			_go_fuzz_dep_.CoverTab[42457]++
//line /usr/local/go/src/net/http/server.go:1320
			return header.get("Connection") == "keep-alive"
//line /usr/local/go/src/net/http/server.go:1320
			// _ = "end of CoverTab[42457]"
//line /usr/local/go/src/net/http/server.go:1320
		}() {
//line /usr/local/go/src/net/http/server.go:1320
			_go_fuzz_dep_.CoverTab[42458]++
									w.closeAfterReply = false
//line /usr/local/go/src/net/http/server.go:1321
			// _ = "end of CoverTab[42458]"
		} else {
//line /usr/local/go/src/net/http/server.go:1322
			_go_fuzz_dep_.CoverTab[42459]++
//line /usr/local/go/src/net/http/server.go:1322
			// _ = "end of CoverTab[42459]"
//line /usr/local/go/src/net/http/server.go:1322
		}
//line /usr/local/go/src/net/http/server.go:1322
		// _ = "end of CoverTab[42456]"
	} else {
//line /usr/local/go/src/net/http/server.go:1323
		_go_fuzz_dep_.CoverTab[42460]++
//line /usr/local/go/src/net/http/server.go:1323
		// _ = "end of CoverTab[42460]"
//line /usr/local/go/src/net/http/server.go:1323
	}
//line /usr/local/go/src/net/http/server.go:1323
	// _ = "end of CoverTab[42413]"
//line /usr/local/go/src/net/http/server.go:1323
	_go_fuzz_dep_.CoverTab[42414]++

//line /usr/local/go/src/net/http/server.go:1326
	hasCL := w.contentLength != -1

	if w.wants10KeepAlive && func() bool {
//line /usr/local/go/src/net/http/server.go:1328
		_go_fuzz_dep_.CoverTab[42461]++
//line /usr/local/go/src/net/http/server.go:1328
		return (isHEAD || func() bool {
//line /usr/local/go/src/net/http/server.go:1328
			_go_fuzz_dep_.CoverTab[42462]++
//line /usr/local/go/src/net/http/server.go:1328
			return hasCL
//line /usr/local/go/src/net/http/server.go:1328
			// _ = "end of CoverTab[42462]"
//line /usr/local/go/src/net/http/server.go:1328
		}() || func() bool {
//line /usr/local/go/src/net/http/server.go:1328
			_go_fuzz_dep_.CoverTab[42463]++
//line /usr/local/go/src/net/http/server.go:1328
			return !bodyAllowedForStatus(w.status)
//line /usr/local/go/src/net/http/server.go:1328
			// _ = "end of CoverTab[42463]"
//line /usr/local/go/src/net/http/server.go:1328
		}())
//line /usr/local/go/src/net/http/server.go:1328
		// _ = "end of CoverTab[42461]"
//line /usr/local/go/src/net/http/server.go:1328
	}() {
//line /usr/local/go/src/net/http/server.go:1328
		_go_fuzz_dep_.CoverTab[42464]++
								_, connectionHeaderSet := header["Connection"]
								if !connectionHeaderSet {
//line /usr/local/go/src/net/http/server.go:1330
			_go_fuzz_dep_.CoverTab[42465]++
									setHeader.connection = "keep-alive"
//line /usr/local/go/src/net/http/server.go:1331
			// _ = "end of CoverTab[42465]"
		} else {
//line /usr/local/go/src/net/http/server.go:1332
			_go_fuzz_dep_.CoverTab[42466]++
//line /usr/local/go/src/net/http/server.go:1332
			// _ = "end of CoverTab[42466]"
//line /usr/local/go/src/net/http/server.go:1332
		}
//line /usr/local/go/src/net/http/server.go:1332
		// _ = "end of CoverTab[42464]"
	} else {
//line /usr/local/go/src/net/http/server.go:1333
		_go_fuzz_dep_.CoverTab[42467]++
//line /usr/local/go/src/net/http/server.go:1333
		if !w.req.ProtoAtLeast(1, 1) || func() bool {
//line /usr/local/go/src/net/http/server.go:1333
			_go_fuzz_dep_.CoverTab[42468]++
//line /usr/local/go/src/net/http/server.go:1333
			return w.wantsClose
//line /usr/local/go/src/net/http/server.go:1333
			// _ = "end of CoverTab[42468]"
//line /usr/local/go/src/net/http/server.go:1333
		}() {
//line /usr/local/go/src/net/http/server.go:1333
			_go_fuzz_dep_.CoverTab[42469]++
									w.closeAfterReply = true
//line /usr/local/go/src/net/http/server.go:1334
			// _ = "end of CoverTab[42469]"
		} else {
//line /usr/local/go/src/net/http/server.go:1335
			_go_fuzz_dep_.CoverTab[42470]++
//line /usr/local/go/src/net/http/server.go:1335
			// _ = "end of CoverTab[42470]"
//line /usr/local/go/src/net/http/server.go:1335
		}
//line /usr/local/go/src/net/http/server.go:1335
		// _ = "end of CoverTab[42467]"
//line /usr/local/go/src/net/http/server.go:1335
	}
//line /usr/local/go/src/net/http/server.go:1335
	// _ = "end of CoverTab[42414]"
//line /usr/local/go/src/net/http/server.go:1335
	_go_fuzz_dep_.CoverTab[42415]++

							if header.get("Connection") == "close" || func() bool {
//line /usr/local/go/src/net/http/server.go:1337
		_go_fuzz_dep_.CoverTab[42471]++
//line /usr/local/go/src/net/http/server.go:1337
		return !keepAlivesEnabled
//line /usr/local/go/src/net/http/server.go:1337
		// _ = "end of CoverTab[42471]"
//line /usr/local/go/src/net/http/server.go:1337
	}() {
//line /usr/local/go/src/net/http/server.go:1337
		_go_fuzz_dep_.CoverTab[42472]++
								w.closeAfterReply = true
//line /usr/local/go/src/net/http/server.go:1338
		// _ = "end of CoverTab[42472]"
	} else {
//line /usr/local/go/src/net/http/server.go:1339
		_go_fuzz_dep_.CoverTab[42473]++
//line /usr/local/go/src/net/http/server.go:1339
		// _ = "end of CoverTab[42473]"
//line /usr/local/go/src/net/http/server.go:1339
	}
//line /usr/local/go/src/net/http/server.go:1339
	// _ = "end of CoverTab[42415]"
//line /usr/local/go/src/net/http/server.go:1339
	_go_fuzz_dep_.CoverTab[42416]++

//line /usr/local/go/src/net/http/server.go:1353
	if ecr, ok := w.req.Body.(*expectContinueReader); ok && func() bool {
//line /usr/local/go/src/net/http/server.go:1353
		_go_fuzz_dep_.CoverTab[42474]++
//line /usr/local/go/src/net/http/server.go:1353
		return !ecr.sawEOF.Load()
//line /usr/local/go/src/net/http/server.go:1353
		// _ = "end of CoverTab[42474]"
//line /usr/local/go/src/net/http/server.go:1353
	}() {
//line /usr/local/go/src/net/http/server.go:1353
		_go_fuzz_dep_.CoverTab[42475]++
								w.closeAfterReply = true
//line /usr/local/go/src/net/http/server.go:1354
		// _ = "end of CoverTab[42475]"
	} else {
//line /usr/local/go/src/net/http/server.go:1355
		_go_fuzz_dep_.CoverTab[42476]++
//line /usr/local/go/src/net/http/server.go:1355
		// _ = "end of CoverTab[42476]"
//line /usr/local/go/src/net/http/server.go:1355
	}
//line /usr/local/go/src/net/http/server.go:1355
	// _ = "end of CoverTab[42416]"
//line /usr/local/go/src/net/http/server.go:1355
	_go_fuzz_dep_.CoverTab[42417]++

//line /usr/local/go/src/net/http/server.go:1364
	if w.req.ContentLength != 0 && func() bool {
//line /usr/local/go/src/net/http/server.go:1364
		_go_fuzz_dep_.CoverTab[42477]++
//line /usr/local/go/src/net/http/server.go:1364
		return !w.closeAfterReply
//line /usr/local/go/src/net/http/server.go:1364
		// _ = "end of CoverTab[42477]"
//line /usr/local/go/src/net/http/server.go:1364
	}() {
//line /usr/local/go/src/net/http/server.go:1364
		_go_fuzz_dep_.CoverTab[42478]++
								var discard, tooBig bool

								switch bdy := w.req.Body.(type) {
		case *expectContinueReader:
//line /usr/local/go/src/net/http/server.go:1368
			_go_fuzz_dep_.CoverTab[42481]++
									if bdy.resp.wroteContinue {
//line /usr/local/go/src/net/http/server.go:1369
				_go_fuzz_dep_.CoverTab[42485]++
										discard = true
//line /usr/local/go/src/net/http/server.go:1370
				// _ = "end of CoverTab[42485]"
			} else {
//line /usr/local/go/src/net/http/server.go:1371
				_go_fuzz_dep_.CoverTab[42486]++
//line /usr/local/go/src/net/http/server.go:1371
				// _ = "end of CoverTab[42486]"
//line /usr/local/go/src/net/http/server.go:1371
			}
//line /usr/local/go/src/net/http/server.go:1371
			// _ = "end of CoverTab[42481]"
		case *body:
//line /usr/local/go/src/net/http/server.go:1372
			_go_fuzz_dep_.CoverTab[42482]++
									bdy.mu.Lock()
									switch {
			case bdy.closed:
//line /usr/local/go/src/net/http/server.go:1375
				_go_fuzz_dep_.CoverTab[42487]++
										if !bdy.sawEOF {
//line /usr/local/go/src/net/http/server.go:1376
					_go_fuzz_dep_.CoverTab[42490]++

											w.closeAfterReply = true
//line /usr/local/go/src/net/http/server.go:1378
					// _ = "end of CoverTab[42490]"
				} else {
//line /usr/local/go/src/net/http/server.go:1379
					_go_fuzz_dep_.CoverTab[42491]++
//line /usr/local/go/src/net/http/server.go:1379
					// _ = "end of CoverTab[42491]"
//line /usr/local/go/src/net/http/server.go:1379
				}
//line /usr/local/go/src/net/http/server.go:1379
				// _ = "end of CoverTab[42487]"
			case bdy.unreadDataSizeLocked() >= maxPostHandlerReadBytes:
//line /usr/local/go/src/net/http/server.go:1380
				_go_fuzz_dep_.CoverTab[42488]++
										tooBig = true
//line /usr/local/go/src/net/http/server.go:1381
				// _ = "end of CoverTab[42488]"
			default:
//line /usr/local/go/src/net/http/server.go:1382
				_go_fuzz_dep_.CoverTab[42489]++
										discard = true
//line /usr/local/go/src/net/http/server.go:1383
				// _ = "end of CoverTab[42489]"
			}
//line /usr/local/go/src/net/http/server.go:1384
			// _ = "end of CoverTab[42482]"
//line /usr/local/go/src/net/http/server.go:1384
			_go_fuzz_dep_.CoverTab[42483]++
									bdy.mu.Unlock()
//line /usr/local/go/src/net/http/server.go:1385
			// _ = "end of CoverTab[42483]"
		default:
//line /usr/local/go/src/net/http/server.go:1386
			_go_fuzz_dep_.CoverTab[42484]++
									discard = true
//line /usr/local/go/src/net/http/server.go:1387
			// _ = "end of CoverTab[42484]"
		}
//line /usr/local/go/src/net/http/server.go:1388
		// _ = "end of CoverTab[42478]"
//line /usr/local/go/src/net/http/server.go:1388
		_go_fuzz_dep_.CoverTab[42479]++

								if discard {
//line /usr/local/go/src/net/http/server.go:1390
			_go_fuzz_dep_.CoverTab[42492]++
									_, err := io.CopyN(io.Discard, w.reqBody, maxPostHandlerReadBytes+1)
									switch err {
			case nil:
//line /usr/local/go/src/net/http/server.go:1393
				_go_fuzz_dep_.CoverTab[42493]++

										tooBig = true
//line /usr/local/go/src/net/http/server.go:1395
				// _ = "end of CoverTab[42493]"
			case ErrBodyReadAfterClose:
//line /usr/local/go/src/net/http/server.go:1396
				_go_fuzz_dep_.CoverTab[42494]++
//line /usr/local/go/src/net/http/server.go:1396
				// _ = "end of CoverTab[42494]"

			case io.EOF:
//line /usr/local/go/src/net/http/server.go:1398
				_go_fuzz_dep_.CoverTab[42495]++

										err = w.reqBody.Close()
										if err != nil {
//line /usr/local/go/src/net/http/server.go:1401
					_go_fuzz_dep_.CoverTab[42497]++
											w.closeAfterReply = true
//line /usr/local/go/src/net/http/server.go:1402
					// _ = "end of CoverTab[42497]"
				} else {
//line /usr/local/go/src/net/http/server.go:1403
					_go_fuzz_dep_.CoverTab[42498]++
//line /usr/local/go/src/net/http/server.go:1403
					// _ = "end of CoverTab[42498]"
//line /usr/local/go/src/net/http/server.go:1403
				}
//line /usr/local/go/src/net/http/server.go:1403
				// _ = "end of CoverTab[42495]"
			default:
//line /usr/local/go/src/net/http/server.go:1404
				_go_fuzz_dep_.CoverTab[42496]++

//line /usr/local/go/src/net/http/server.go:1408
				w.closeAfterReply = true
//line /usr/local/go/src/net/http/server.go:1408
				// _ = "end of CoverTab[42496]"
			}
//line /usr/local/go/src/net/http/server.go:1409
			// _ = "end of CoverTab[42492]"
		} else {
//line /usr/local/go/src/net/http/server.go:1410
			_go_fuzz_dep_.CoverTab[42499]++
//line /usr/local/go/src/net/http/server.go:1410
			// _ = "end of CoverTab[42499]"
//line /usr/local/go/src/net/http/server.go:1410
		}
//line /usr/local/go/src/net/http/server.go:1410
		// _ = "end of CoverTab[42479]"
//line /usr/local/go/src/net/http/server.go:1410
		_go_fuzz_dep_.CoverTab[42480]++

								if tooBig {
//line /usr/local/go/src/net/http/server.go:1412
			_go_fuzz_dep_.CoverTab[42500]++
									w.requestTooLarge()
									delHeader("Connection")
									setHeader.connection = "close"
//line /usr/local/go/src/net/http/server.go:1415
			// _ = "end of CoverTab[42500]"
		} else {
//line /usr/local/go/src/net/http/server.go:1416
			_go_fuzz_dep_.CoverTab[42501]++
//line /usr/local/go/src/net/http/server.go:1416
			// _ = "end of CoverTab[42501]"
//line /usr/local/go/src/net/http/server.go:1416
		}
//line /usr/local/go/src/net/http/server.go:1416
		// _ = "end of CoverTab[42480]"
	} else {
//line /usr/local/go/src/net/http/server.go:1417
		_go_fuzz_dep_.CoverTab[42502]++
//line /usr/local/go/src/net/http/server.go:1417
		// _ = "end of CoverTab[42502]"
//line /usr/local/go/src/net/http/server.go:1417
	}
//line /usr/local/go/src/net/http/server.go:1417
	// _ = "end of CoverTab[42417]"
//line /usr/local/go/src/net/http/server.go:1417
	_go_fuzz_dep_.CoverTab[42418]++

							code := w.status
							if bodyAllowedForStatus(code) {
//line /usr/local/go/src/net/http/server.go:1420
		_go_fuzz_dep_.CoverTab[42503]++

								_, haveType := header["Content-Type"]

//line /usr/local/go/src/net/http/server.go:1426
		ce := header.Get("Content-Encoding")
		hasCE := len(ce) > 0
		if !hasCE && func() bool {
//line /usr/local/go/src/net/http/server.go:1428
			_go_fuzz_dep_.CoverTab[42504]++
//line /usr/local/go/src/net/http/server.go:1428
			return !haveType
//line /usr/local/go/src/net/http/server.go:1428
			// _ = "end of CoverTab[42504]"
//line /usr/local/go/src/net/http/server.go:1428
		}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1428
			_go_fuzz_dep_.CoverTab[42505]++
//line /usr/local/go/src/net/http/server.go:1428
			return !hasTE
//line /usr/local/go/src/net/http/server.go:1428
			// _ = "end of CoverTab[42505]"
//line /usr/local/go/src/net/http/server.go:1428
		}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1428
			_go_fuzz_dep_.CoverTab[42506]++
//line /usr/local/go/src/net/http/server.go:1428
			return len(p) > 0
//line /usr/local/go/src/net/http/server.go:1428
			// _ = "end of CoverTab[42506]"
//line /usr/local/go/src/net/http/server.go:1428
		}() {
//line /usr/local/go/src/net/http/server.go:1428
			_go_fuzz_dep_.CoverTab[42507]++
									setHeader.contentType = DetectContentType(p)
//line /usr/local/go/src/net/http/server.go:1429
			// _ = "end of CoverTab[42507]"
		} else {
//line /usr/local/go/src/net/http/server.go:1430
			_go_fuzz_dep_.CoverTab[42508]++
//line /usr/local/go/src/net/http/server.go:1430
			// _ = "end of CoverTab[42508]"
//line /usr/local/go/src/net/http/server.go:1430
		}
//line /usr/local/go/src/net/http/server.go:1430
		// _ = "end of CoverTab[42503]"
	} else {
//line /usr/local/go/src/net/http/server.go:1431
		_go_fuzz_dep_.CoverTab[42509]++
								for _, k := range suppressedHeaders(code) {
//line /usr/local/go/src/net/http/server.go:1432
			_go_fuzz_dep_.CoverTab[42510]++
									delHeader(k)
//line /usr/local/go/src/net/http/server.go:1433
			// _ = "end of CoverTab[42510]"
		}
//line /usr/local/go/src/net/http/server.go:1434
		// _ = "end of CoverTab[42509]"
	}
//line /usr/local/go/src/net/http/server.go:1435
	// _ = "end of CoverTab[42418]"
//line /usr/local/go/src/net/http/server.go:1435
	_go_fuzz_dep_.CoverTab[42419]++

							if !header.has("Date") {
//line /usr/local/go/src/net/http/server.go:1437
		_go_fuzz_dep_.CoverTab[42511]++
								setHeader.date = appendTime(cw.res.dateBuf[:0], time.Now())
//line /usr/local/go/src/net/http/server.go:1438
		// _ = "end of CoverTab[42511]"
	} else {
//line /usr/local/go/src/net/http/server.go:1439
		_go_fuzz_dep_.CoverTab[42512]++
//line /usr/local/go/src/net/http/server.go:1439
		// _ = "end of CoverTab[42512]"
//line /usr/local/go/src/net/http/server.go:1439
	}
//line /usr/local/go/src/net/http/server.go:1439
	// _ = "end of CoverTab[42419]"
//line /usr/local/go/src/net/http/server.go:1439
	_go_fuzz_dep_.CoverTab[42420]++

							if hasCL && func() bool {
//line /usr/local/go/src/net/http/server.go:1441
		_go_fuzz_dep_.CoverTab[42513]++
//line /usr/local/go/src/net/http/server.go:1441
		return hasTE
//line /usr/local/go/src/net/http/server.go:1441
		// _ = "end of CoverTab[42513]"
//line /usr/local/go/src/net/http/server.go:1441
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1441
		_go_fuzz_dep_.CoverTab[42514]++
//line /usr/local/go/src/net/http/server.go:1441
		return te != "identity"
//line /usr/local/go/src/net/http/server.go:1441
		// _ = "end of CoverTab[42514]"
//line /usr/local/go/src/net/http/server.go:1441
	}() {
//line /usr/local/go/src/net/http/server.go:1441
		_go_fuzz_dep_.CoverTab[42515]++

//line /usr/local/go/src/net/http/server.go:1444
		w.conn.server.logf("http: WriteHeader called with both Transfer-Encoding of %q and a Content-Length of %d",
			te, w.contentLength)
								delHeader("Content-Length")
								hasCL = false
//line /usr/local/go/src/net/http/server.go:1447
		// _ = "end of CoverTab[42515]"
	} else {
//line /usr/local/go/src/net/http/server.go:1448
		_go_fuzz_dep_.CoverTab[42516]++
//line /usr/local/go/src/net/http/server.go:1448
		// _ = "end of CoverTab[42516]"
//line /usr/local/go/src/net/http/server.go:1448
	}
//line /usr/local/go/src/net/http/server.go:1448
	// _ = "end of CoverTab[42420]"
//line /usr/local/go/src/net/http/server.go:1448
	_go_fuzz_dep_.CoverTab[42421]++

							if w.req.Method == "HEAD" || func() bool {
//line /usr/local/go/src/net/http/server.go:1450
		_go_fuzz_dep_.CoverTab[42517]++
//line /usr/local/go/src/net/http/server.go:1450
		return !bodyAllowedForStatus(code)
//line /usr/local/go/src/net/http/server.go:1450
		// _ = "end of CoverTab[42517]"
//line /usr/local/go/src/net/http/server.go:1450
	}() || func() bool {
//line /usr/local/go/src/net/http/server.go:1450
		_go_fuzz_dep_.CoverTab[42518]++
//line /usr/local/go/src/net/http/server.go:1450
		return code == StatusNoContent
//line /usr/local/go/src/net/http/server.go:1450
		// _ = "end of CoverTab[42518]"
//line /usr/local/go/src/net/http/server.go:1450
	}() {
//line /usr/local/go/src/net/http/server.go:1450
		_go_fuzz_dep_.CoverTab[42519]++

								delHeader("Transfer-Encoding")
//line /usr/local/go/src/net/http/server.go:1452
		// _ = "end of CoverTab[42519]"
	} else {
//line /usr/local/go/src/net/http/server.go:1453
		_go_fuzz_dep_.CoverTab[42520]++
//line /usr/local/go/src/net/http/server.go:1453
		if hasCL {
//line /usr/local/go/src/net/http/server.go:1453
			_go_fuzz_dep_.CoverTab[42521]++

									delHeader("Transfer-Encoding")
//line /usr/local/go/src/net/http/server.go:1455
			// _ = "end of CoverTab[42521]"
		} else {
//line /usr/local/go/src/net/http/server.go:1456
			_go_fuzz_dep_.CoverTab[42522]++
//line /usr/local/go/src/net/http/server.go:1456
			if w.req.ProtoAtLeast(1, 1) {
//line /usr/local/go/src/net/http/server.go:1456
				_go_fuzz_dep_.CoverTab[42523]++

//line /usr/local/go/src/net/http/server.go:1462
				if hasTE && func() bool {
//line /usr/local/go/src/net/http/server.go:1462
					_go_fuzz_dep_.CoverTab[42524]++
//line /usr/local/go/src/net/http/server.go:1462
					return te == "identity"
//line /usr/local/go/src/net/http/server.go:1462
					// _ = "end of CoverTab[42524]"
//line /usr/local/go/src/net/http/server.go:1462
				}() {
//line /usr/local/go/src/net/http/server.go:1462
					_go_fuzz_dep_.CoverTab[42525]++
											cw.chunking = false
											w.closeAfterReply = true
											delHeader("Transfer-Encoding")
//line /usr/local/go/src/net/http/server.go:1465
					// _ = "end of CoverTab[42525]"
				} else {
//line /usr/local/go/src/net/http/server.go:1466
					_go_fuzz_dep_.CoverTab[42526]++

//line /usr/local/go/src/net/http/server.go:1469
					cw.chunking = true
					setHeader.transferEncoding = "chunked"
					if hasTE && func() bool {
//line /usr/local/go/src/net/http/server.go:1471
						_go_fuzz_dep_.CoverTab[42527]++
//line /usr/local/go/src/net/http/server.go:1471
						return te == "chunked"
//line /usr/local/go/src/net/http/server.go:1471
						// _ = "end of CoverTab[42527]"
//line /usr/local/go/src/net/http/server.go:1471
					}() {
//line /usr/local/go/src/net/http/server.go:1471
						_go_fuzz_dep_.CoverTab[42528]++

												delHeader("Transfer-Encoding")
//line /usr/local/go/src/net/http/server.go:1473
						// _ = "end of CoverTab[42528]"
					} else {
//line /usr/local/go/src/net/http/server.go:1474
						_go_fuzz_dep_.CoverTab[42529]++
//line /usr/local/go/src/net/http/server.go:1474
						// _ = "end of CoverTab[42529]"
//line /usr/local/go/src/net/http/server.go:1474
					}
//line /usr/local/go/src/net/http/server.go:1474
					// _ = "end of CoverTab[42526]"
				}
//line /usr/local/go/src/net/http/server.go:1475
				// _ = "end of CoverTab[42523]"
			} else {
//line /usr/local/go/src/net/http/server.go:1476
				_go_fuzz_dep_.CoverTab[42530]++

//line /usr/local/go/src/net/http/server.go:1480
				w.closeAfterReply = true
										delHeader("Transfer-Encoding")
//line /usr/local/go/src/net/http/server.go:1481
				// _ = "end of CoverTab[42530]"
			}
//line /usr/local/go/src/net/http/server.go:1482
			// _ = "end of CoverTab[42522]"
//line /usr/local/go/src/net/http/server.go:1482
		}
//line /usr/local/go/src/net/http/server.go:1482
		// _ = "end of CoverTab[42520]"
//line /usr/local/go/src/net/http/server.go:1482
	}
//line /usr/local/go/src/net/http/server.go:1482
	// _ = "end of CoverTab[42421]"
//line /usr/local/go/src/net/http/server.go:1482
	_go_fuzz_dep_.CoverTab[42422]++

//line /usr/local/go/src/net/http/server.go:1485
	if cw.chunking {
//line /usr/local/go/src/net/http/server.go:1485
		_go_fuzz_dep_.CoverTab[42531]++
								delHeader("Content-Length")
//line /usr/local/go/src/net/http/server.go:1486
		// _ = "end of CoverTab[42531]"
	} else {
//line /usr/local/go/src/net/http/server.go:1487
		_go_fuzz_dep_.CoverTab[42532]++
//line /usr/local/go/src/net/http/server.go:1487
		// _ = "end of CoverTab[42532]"
//line /usr/local/go/src/net/http/server.go:1487
	}
//line /usr/local/go/src/net/http/server.go:1487
	// _ = "end of CoverTab[42422]"
//line /usr/local/go/src/net/http/server.go:1487
	_go_fuzz_dep_.CoverTab[42423]++
							if !w.req.ProtoAtLeast(1, 0) {
//line /usr/local/go/src/net/http/server.go:1488
		_go_fuzz_dep_.CoverTab[42533]++
								return
//line /usr/local/go/src/net/http/server.go:1489
		// _ = "end of CoverTab[42533]"
	} else {
//line /usr/local/go/src/net/http/server.go:1490
		_go_fuzz_dep_.CoverTab[42534]++
//line /usr/local/go/src/net/http/server.go:1490
		// _ = "end of CoverTab[42534]"
//line /usr/local/go/src/net/http/server.go:1490
	}
//line /usr/local/go/src/net/http/server.go:1490
	// _ = "end of CoverTab[42423]"
//line /usr/local/go/src/net/http/server.go:1490
	_go_fuzz_dep_.CoverTab[42424]++

//line /usr/local/go/src/net/http/server.go:1495
	delConnectionHeader := w.closeAfterReply && func() bool {
//line /usr/local/go/src/net/http/server.go:1495
		_go_fuzz_dep_.CoverTab[42535]++
//line /usr/local/go/src/net/http/server.go:1495
		return (!keepAlivesEnabled || func() bool {
									_go_fuzz_dep_.CoverTab[42536]++
//line /usr/local/go/src/net/http/server.go:1496
			return !hasToken(cw.header.get("Connection"), "close")
//line /usr/local/go/src/net/http/server.go:1496
			// _ = "end of CoverTab[42536]"
//line /usr/local/go/src/net/http/server.go:1496
		}())
//line /usr/local/go/src/net/http/server.go:1496
		// _ = "end of CoverTab[42535]"
//line /usr/local/go/src/net/http/server.go:1496
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1496
		_go_fuzz_dep_.CoverTab[42537]++
//line /usr/local/go/src/net/http/server.go:1496
		return !isProtocolSwitchResponse(w.status, header)
								// _ = "end of CoverTab[42537]"
//line /usr/local/go/src/net/http/server.go:1497
	}()
	if delConnectionHeader {
//line /usr/local/go/src/net/http/server.go:1498
		_go_fuzz_dep_.CoverTab[42538]++
								delHeader("Connection")
								if w.req.ProtoAtLeast(1, 1) {
//line /usr/local/go/src/net/http/server.go:1500
			_go_fuzz_dep_.CoverTab[42539]++
									setHeader.connection = "close"
//line /usr/local/go/src/net/http/server.go:1501
			// _ = "end of CoverTab[42539]"
		} else {
//line /usr/local/go/src/net/http/server.go:1502
			_go_fuzz_dep_.CoverTab[42540]++
//line /usr/local/go/src/net/http/server.go:1502
			// _ = "end of CoverTab[42540]"
//line /usr/local/go/src/net/http/server.go:1502
		}
//line /usr/local/go/src/net/http/server.go:1502
		// _ = "end of CoverTab[42538]"
	} else {
//line /usr/local/go/src/net/http/server.go:1503
		_go_fuzz_dep_.CoverTab[42541]++
//line /usr/local/go/src/net/http/server.go:1503
		// _ = "end of CoverTab[42541]"
//line /usr/local/go/src/net/http/server.go:1503
	}
//line /usr/local/go/src/net/http/server.go:1503
	// _ = "end of CoverTab[42424]"
//line /usr/local/go/src/net/http/server.go:1503
	_go_fuzz_dep_.CoverTab[42425]++

							writeStatusLine(w.conn.bufw, w.req.ProtoAtLeast(1, 1), code, w.statusBuf[:])
							cw.header.WriteSubset(w.conn.bufw, excludeHeader)
							setHeader.Write(w.conn.bufw)
							w.conn.bufw.Write(crlf)
//line /usr/local/go/src/net/http/server.go:1508
	// _ = "end of CoverTab[42425]"
}

// foreachHeaderElement splits v according to the "#rule" construction
//line /usr/local/go/src/net/http/server.go:1511
// in RFC 7230 section 7 and calls fn for each non-empty element.
//line /usr/local/go/src/net/http/server.go:1513
func foreachHeaderElement(v string, fn func(string)) {
//line /usr/local/go/src/net/http/server.go:1513
	_go_fuzz_dep_.CoverTab[42542]++
							v = textproto.TrimString(v)
							if v == "" {
//line /usr/local/go/src/net/http/server.go:1515
		_go_fuzz_dep_.CoverTab[42545]++
								return
//line /usr/local/go/src/net/http/server.go:1516
		// _ = "end of CoverTab[42545]"
	} else {
//line /usr/local/go/src/net/http/server.go:1517
		_go_fuzz_dep_.CoverTab[42546]++
//line /usr/local/go/src/net/http/server.go:1517
		// _ = "end of CoverTab[42546]"
//line /usr/local/go/src/net/http/server.go:1517
	}
//line /usr/local/go/src/net/http/server.go:1517
	// _ = "end of CoverTab[42542]"
//line /usr/local/go/src/net/http/server.go:1517
	_go_fuzz_dep_.CoverTab[42543]++
							if !strings.Contains(v, ",") {
//line /usr/local/go/src/net/http/server.go:1518
		_go_fuzz_dep_.CoverTab[42547]++
								fn(v)
								return
//line /usr/local/go/src/net/http/server.go:1520
		// _ = "end of CoverTab[42547]"
	} else {
//line /usr/local/go/src/net/http/server.go:1521
		_go_fuzz_dep_.CoverTab[42548]++
//line /usr/local/go/src/net/http/server.go:1521
		// _ = "end of CoverTab[42548]"
//line /usr/local/go/src/net/http/server.go:1521
	}
//line /usr/local/go/src/net/http/server.go:1521
	// _ = "end of CoverTab[42543]"
//line /usr/local/go/src/net/http/server.go:1521
	_go_fuzz_dep_.CoverTab[42544]++
							for _, f := range strings.Split(v, ",") {
//line /usr/local/go/src/net/http/server.go:1522
		_go_fuzz_dep_.CoverTab[42549]++
								if f = textproto.TrimString(f); f != "" {
//line /usr/local/go/src/net/http/server.go:1523
			_go_fuzz_dep_.CoverTab[42550]++
									fn(f)
//line /usr/local/go/src/net/http/server.go:1524
			// _ = "end of CoverTab[42550]"
		} else {
//line /usr/local/go/src/net/http/server.go:1525
			_go_fuzz_dep_.CoverTab[42551]++
//line /usr/local/go/src/net/http/server.go:1525
			// _ = "end of CoverTab[42551]"
//line /usr/local/go/src/net/http/server.go:1525
		}
//line /usr/local/go/src/net/http/server.go:1525
		// _ = "end of CoverTab[42549]"
	}
//line /usr/local/go/src/net/http/server.go:1526
	// _ = "end of CoverTab[42544]"
}

// writeStatusLine writes an HTTP/1.x Status-Line (RFC 7230 Section 3.1.2)
//line /usr/local/go/src/net/http/server.go:1529
// to bw. is11 is whether the HTTP request is HTTP/1.1. false means HTTP/1.0.
//line /usr/local/go/src/net/http/server.go:1529
// code is the response status code.
//line /usr/local/go/src/net/http/server.go:1529
// scratch is an optional scratch buffer. If it has at least capacity 3, it's used.
//line /usr/local/go/src/net/http/server.go:1533
func writeStatusLine(bw *bufio.Writer, is11 bool, code int, scratch []byte) {
//line /usr/local/go/src/net/http/server.go:1533
	_go_fuzz_dep_.CoverTab[42552]++
							if is11 {
//line /usr/local/go/src/net/http/server.go:1534
		_go_fuzz_dep_.CoverTab[42554]++
								bw.WriteString("HTTP/1.1 ")
//line /usr/local/go/src/net/http/server.go:1535
		// _ = "end of CoverTab[42554]"
	} else {
//line /usr/local/go/src/net/http/server.go:1536
		_go_fuzz_dep_.CoverTab[42555]++
								bw.WriteString("HTTP/1.0 ")
//line /usr/local/go/src/net/http/server.go:1537
		// _ = "end of CoverTab[42555]"
	}
//line /usr/local/go/src/net/http/server.go:1538
	// _ = "end of CoverTab[42552]"
//line /usr/local/go/src/net/http/server.go:1538
	_go_fuzz_dep_.CoverTab[42553]++
							if text := StatusText(code); text != "" {
//line /usr/local/go/src/net/http/server.go:1539
		_go_fuzz_dep_.CoverTab[42556]++
								bw.Write(strconv.AppendInt(scratch[:0], int64(code), 10))
								bw.WriteByte(' ')
								bw.WriteString(text)
								bw.WriteString("\r\n")
//line /usr/local/go/src/net/http/server.go:1543
		// _ = "end of CoverTab[42556]"
	} else {
//line /usr/local/go/src/net/http/server.go:1544
		_go_fuzz_dep_.CoverTab[42557]++

								fmt.Fprintf(bw, "%03d status code %d\r\n", code, code)
//line /usr/local/go/src/net/http/server.go:1546
		// _ = "end of CoverTab[42557]"
	}
//line /usr/local/go/src/net/http/server.go:1547
	// _ = "end of CoverTab[42553]"
}

// bodyAllowed reports whether a Write is allowed for this response type.
//line /usr/local/go/src/net/http/server.go:1550
// It's illegal to call this before the header has been flushed.
//line /usr/local/go/src/net/http/server.go:1552
func (w *response) bodyAllowed() bool {
//line /usr/local/go/src/net/http/server.go:1552
	_go_fuzz_dep_.CoverTab[42558]++
							if !w.wroteHeader {
//line /usr/local/go/src/net/http/server.go:1553
		_go_fuzz_dep_.CoverTab[42560]++
								panic("")
//line /usr/local/go/src/net/http/server.go:1554
		// _ = "end of CoverTab[42560]"
	} else {
//line /usr/local/go/src/net/http/server.go:1555
		_go_fuzz_dep_.CoverTab[42561]++
//line /usr/local/go/src/net/http/server.go:1555
		// _ = "end of CoverTab[42561]"
//line /usr/local/go/src/net/http/server.go:1555
	}
//line /usr/local/go/src/net/http/server.go:1555
	// _ = "end of CoverTab[42558]"
//line /usr/local/go/src/net/http/server.go:1555
	_go_fuzz_dep_.CoverTab[42559]++
							return bodyAllowedForStatus(w.status)
//line /usr/local/go/src/net/http/server.go:1556
	// _ = "end of CoverTab[42559]"
}

// The Life Of A Write is like this:
//line /usr/local/go/src/net/http/server.go:1559
//
//line /usr/local/go/src/net/http/server.go:1559
// Handler starts. No header has been sent. The handler can either
//line /usr/local/go/src/net/http/server.go:1559
// write a header, or just start writing. Writing before sending a header
//line /usr/local/go/src/net/http/server.go:1559
// sends an implicitly empty 200 OK header.
//line /usr/local/go/src/net/http/server.go:1559
//
//line /usr/local/go/src/net/http/server.go:1559
// If the handler didn't declare a Content-Length up front, we either
//line /usr/local/go/src/net/http/server.go:1559
// go into chunking mode or, if the handler finishes running before
//line /usr/local/go/src/net/http/server.go:1559
// the chunking buffer size, we compute a Content-Length and send that
//line /usr/local/go/src/net/http/server.go:1559
// in the header instead.
//line /usr/local/go/src/net/http/server.go:1559
//
//line /usr/local/go/src/net/http/server.go:1559
// Likewise, if the handler didn't set a Content-Type, we sniff that
//line /usr/local/go/src/net/http/server.go:1559
// from the initial chunk of output.
//line /usr/local/go/src/net/http/server.go:1559
//
//line /usr/local/go/src/net/http/server.go:1559
// The Writers are wired together like:
//line /usr/local/go/src/net/http/server.go:1559
//
//line /usr/local/go/src/net/http/server.go:1559
//  1. *response (the ResponseWriter) ->
//line /usr/local/go/src/net/http/server.go:1559
//  2. (*response).w, a *bufio.Writer of bufferBeforeChunkingSize bytes ->
//line /usr/local/go/src/net/http/server.go:1559
//  3. chunkWriter.Writer (whose writeHeader finalizes Content-Length/Type)
//line /usr/local/go/src/net/http/server.go:1559
//     and which writes the chunk headers, if needed ->
//line /usr/local/go/src/net/http/server.go:1559
//  4. conn.bufw, a *bufio.Writer of default (4kB) bytes, writing to ->
//line /usr/local/go/src/net/http/server.go:1559
//  5. checkConnErrorWriter{c}, which notes any non-nil error on Write
//line /usr/local/go/src/net/http/server.go:1559
//     and populates c.werr with it if so, but otherwise writes to ->
//line /usr/local/go/src/net/http/server.go:1559
//  6. the rwc, the net.Conn.
//line /usr/local/go/src/net/http/server.go:1559
//
//line /usr/local/go/src/net/http/server.go:1559
// TODO(bradfitz): short-circuit some of the buffering when the
//line /usr/local/go/src/net/http/server.go:1559
// initial header contains both a Content-Type and Content-Length.
//line /usr/local/go/src/net/http/server.go:1559
// Also short-circuit in (1) when the header's been sent and not in
//line /usr/local/go/src/net/http/server.go:1559
// chunking mode, writing directly to (4) instead, if (2) has no
//line /usr/local/go/src/net/http/server.go:1559
// buffered data. More generally, we could short-circuit from (1) to
//line /usr/local/go/src/net/http/server.go:1559
// (3) even in chunking mode if the write size from (1) is over some
//line /usr/local/go/src/net/http/server.go:1559
// threshold and nothing is in (2).  The answer might be mostly making
//line /usr/local/go/src/net/http/server.go:1559
// bufferBeforeChunkingSize smaller and having bufio's fast-paths deal
//line /usr/local/go/src/net/http/server.go:1559
// with this instead.
//line /usr/local/go/src/net/http/server.go:1593
func (w *response) Write(data []byte) (n int, err error) {
//line /usr/local/go/src/net/http/server.go:1593
	_go_fuzz_dep_.CoverTab[42562]++
							return w.write(len(data), data, "")
//line /usr/local/go/src/net/http/server.go:1594
	// _ = "end of CoverTab[42562]"
}

func (w *response) WriteString(data string) (n int, err error) {
//line /usr/local/go/src/net/http/server.go:1597
	_go_fuzz_dep_.CoverTab[42563]++
							return w.write(len(data), nil, data)
//line /usr/local/go/src/net/http/server.go:1598
	// _ = "end of CoverTab[42563]"
}

// either dataB or dataS is non-zero.
func (w *response) write(lenData int, dataB []byte, dataS string) (n int, err error) {
//line /usr/local/go/src/net/http/server.go:1602
	_go_fuzz_dep_.CoverTab[42564]++
							if w.conn.hijacked() {
//line /usr/local/go/src/net/http/server.go:1603
		_go_fuzz_dep_.CoverTab[42571]++
								if lenData > 0 {
//line /usr/local/go/src/net/http/server.go:1604
			_go_fuzz_dep_.CoverTab[42573]++
									caller := relevantCaller()
									w.conn.server.logf("http: response.Write on hijacked connection from %s (%s:%d)", caller.Function, path.Base(caller.File), caller.Line)
//line /usr/local/go/src/net/http/server.go:1606
			// _ = "end of CoverTab[42573]"
		} else {
//line /usr/local/go/src/net/http/server.go:1607
			_go_fuzz_dep_.CoverTab[42574]++
//line /usr/local/go/src/net/http/server.go:1607
			// _ = "end of CoverTab[42574]"
//line /usr/local/go/src/net/http/server.go:1607
		}
//line /usr/local/go/src/net/http/server.go:1607
		// _ = "end of CoverTab[42571]"
//line /usr/local/go/src/net/http/server.go:1607
		_go_fuzz_dep_.CoverTab[42572]++
								return 0, ErrHijacked
//line /usr/local/go/src/net/http/server.go:1608
		// _ = "end of CoverTab[42572]"
	} else {
//line /usr/local/go/src/net/http/server.go:1609
		_go_fuzz_dep_.CoverTab[42575]++
//line /usr/local/go/src/net/http/server.go:1609
		// _ = "end of CoverTab[42575]"
//line /usr/local/go/src/net/http/server.go:1609
	}
//line /usr/local/go/src/net/http/server.go:1609
	// _ = "end of CoverTab[42564]"
//line /usr/local/go/src/net/http/server.go:1609
	_go_fuzz_dep_.CoverTab[42565]++

							if w.canWriteContinue.Load() {
//line /usr/local/go/src/net/http/server.go:1611
		_go_fuzz_dep_.CoverTab[42576]++

//line /usr/local/go/src/net/http/server.go:1616
		w.writeContinueMu.Lock()
								w.canWriteContinue.Store(false)
								w.writeContinueMu.Unlock()
//line /usr/local/go/src/net/http/server.go:1618
		// _ = "end of CoverTab[42576]"
	} else {
//line /usr/local/go/src/net/http/server.go:1619
		_go_fuzz_dep_.CoverTab[42577]++
//line /usr/local/go/src/net/http/server.go:1619
		// _ = "end of CoverTab[42577]"
//line /usr/local/go/src/net/http/server.go:1619
	}
//line /usr/local/go/src/net/http/server.go:1619
	// _ = "end of CoverTab[42565]"
//line /usr/local/go/src/net/http/server.go:1619
	_go_fuzz_dep_.CoverTab[42566]++

							if !w.wroteHeader {
//line /usr/local/go/src/net/http/server.go:1621
		_go_fuzz_dep_.CoverTab[42578]++
								w.WriteHeader(StatusOK)
//line /usr/local/go/src/net/http/server.go:1622
		// _ = "end of CoverTab[42578]"
	} else {
//line /usr/local/go/src/net/http/server.go:1623
		_go_fuzz_dep_.CoverTab[42579]++
//line /usr/local/go/src/net/http/server.go:1623
		// _ = "end of CoverTab[42579]"
//line /usr/local/go/src/net/http/server.go:1623
	}
//line /usr/local/go/src/net/http/server.go:1623
	// _ = "end of CoverTab[42566]"
//line /usr/local/go/src/net/http/server.go:1623
	_go_fuzz_dep_.CoverTab[42567]++
							if lenData == 0 {
//line /usr/local/go/src/net/http/server.go:1624
		_go_fuzz_dep_.CoverTab[42580]++
								return 0, nil
//line /usr/local/go/src/net/http/server.go:1625
		// _ = "end of CoverTab[42580]"
	} else {
//line /usr/local/go/src/net/http/server.go:1626
		_go_fuzz_dep_.CoverTab[42581]++
//line /usr/local/go/src/net/http/server.go:1626
		// _ = "end of CoverTab[42581]"
//line /usr/local/go/src/net/http/server.go:1626
	}
//line /usr/local/go/src/net/http/server.go:1626
	// _ = "end of CoverTab[42567]"
//line /usr/local/go/src/net/http/server.go:1626
	_go_fuzz_dep_.CoverTab[42568]++
							if !w.bodyAllowed() {
//line /usr/local/go/src/net/http/server.go:1627
		_go_fuzz_dep_.CoverTab[42582]++
								return 0, ErrBodyNotAllowed
//line /usr/local/go/src/net/http/server.go:1628
		// _ = "end of CoverTab[42582]"
	} else {
//line /usr/local/go/src/net/http/server.go:1629
		_go_fuzz_dep_.CoverTab[42583]++
//line /usr/local/go/src/net/http/server.go:1629
		// _ = "end of CoverTab[42583]"
//line /usr/local/go/src/net/http/server.go:1629
	}
//line /usr/local/go/src/net/http/server.go:1629
	// _ = "end of CoverTab[42568]"
//line /usr/local/go/src/net/http/server.go:1629
	_go_fuzz_dep_.CoverTab[42569]++

							w.written += int64(lenData)
							if w.contentLength != -1 && func() bool {
//line /usr/local/go/src/net/http/server.go:1632
		_go_fuzz_dep_.CoverTab[42584]++
//line /usr/local/go/src/net/http/server.go:1632
		return w.written > w.contentLength
//line /usr/local/go/src/net/http/server.go:1632
		// _ = "end of CoverTab[42584]"
//line /usr/local/go/src/net/http/server.go:1632
	}() {
//line /usr/local/go/src/net/http/server.go:1632
		_go_fuzz_dep_.CoverTab[42585]++
								return 0, ErrContentLength
//line /usr/local/go/src/net/http/server.go:1633
		// _ = "end of CoverTab[42585]"
	} else {
//line /usr/local/go/src/net/http/server.go:1634
		_go_fuzz_dep_.CoverTab[42586]++
//line /usr/local/go/src/net/http/server.go:1634
		// _ = "end of CoverTab[42586]"
//line /usr/local/go/src/net/http/server.go:1634
	}
//line /usr/local/go/src/net/http/server.go:1634
	// _ = "end of CoverTab[42569]"
//line /usr/local/go/src/net/http/server.go:1634
	_go_fuzz_dep_.CoverTab[42570]++
							if dataB != nil {
//line /usr/local/go/src/net/http/server.go:1635
		_go_fuzz_dep_.CoverTab[42587]++
								return w.w.Write(dataB)
//line /usr/local/go/src/net/http/server.go:1636
		// _ = "end of CoverTab[42587]"
	} else {
//line /usr/local/go/src/net/http/server.go:1637
		_go_fuzz_dep_.CoverTab[42588]++
								return w.w.WriteString(dataS)
//line /usr/local/go/src/net/http/server.go:1638
		// _ = "end of CoverTab[42588]"
	}
//line /usr/local/go/src/net/http/server.go:1639
	// _ = "end of CoverTab[42570]"
}

func (w *response) finishRequest() {
//line /usr/local/go/src/net/http/server.go:1642
	_go_fuzz_dep_.CoverTab[42589]++
							w.handlerDone.Store(true)

							if !w.wroteHeader {
//line /usr/local/go/src/net/http/server.go:1645
		_go_fuzz_dep_.CoverTab[42591]++
								w.WriteHeader(StatusOK)
//line /usr/local/go/src/net/http/server.go:1646
		// _ = "end of CoverTab[42591]"
	} else {
//line /usr/local/go/src/net/http/server.go:1647
		_go_fuzz_dep_.CoverTab[42592]++
//line /usr/local/go/src/net/http/server.go:1647
		// _ = "end of CoverTab[42592]"
//line /usr/local/go/src/net/http/server.go:1647
	}
//line /usr/local/go/src/net/http/server.go:1647
	// _ = "end of CoverTab[42589]"
//line /usr/local/go/src/net/http/server.go:1647
	_go_fuzz_dep_.CoverTab[42590]++

							w.w.Flush()
							putBufioWriter(w.w)
							w.cw.close()
							w.conn.bufw.Flush()

							w.conn.r.abortPendingRead()

//line /usr/local/go/src/net/http/server.go:1658
	w.reqBody.Close()

	if w.req.MultipartForm != nil {
//line /usr/local/go/src/net/http/server.go:1660
		_go_fuzz_dep_.CoverTab[42593]++
								w.req.MultipartForm.RemoveAll()
//line /usr/local/go/src/net/http/server.go:1661
		// _ = "end of CoverTab[42593]"
	} else {
//line /usr/local/go/src/net/http/server.go:1662
		_go_fuzz_dep_.CoverTab[42594]++
//line /usr/local/go/src/net/http/server.go:1662
		// _ = "end of CoverTab[42594]"
//line /usr/local/go/src/net/http/server.go:1662
	}
//line /usr/local/go/src/net/http/server.go:1662
	// _ = "end of CoverTab[42590]"
}

// shouldReuseConnection reports whether the underlying TCP connection can be reused.
//line /usr/local/go/src/net/http/server.go:1665
// It must only be called after the handler is done executing.
//line /usr/local/go/src/net/http/server.go:1667
func (w *response) shouldReuseConnection() bool {
//line /usr/local/go/src/net/http/server.go:1667
	_go_fuzz_dep_.CoverTab[42595]++
							if w.closeAfterReply {
//line /usr/local/go/src/net/http/server.go:1668
		_go_fuzz_dep_.CoverTab[42600]++

//line /usr/local/go/src/net/http/server.go:1672
		return false
//line /usr/local/go/src/net/http/server.go:1672
		// _ = "end of CoverTab[42600]"
	} else {
//line /usr/local/go/src/net/http/server.go:1673
		_go_fuzz_dep_.CoverTab[42601]++
//line /usr/local/go/src/net/http/server.go:1673
		// _ = "end of CoverTab[42601]"
//line /usr/local/go/src/net/http/server.go:1673
	}
//line /usr/local/go/src/net/http/server.go:1673
	// _ = "end of CoverTab[42595]"
//line /usr/local/go/src/net/http/server.go:1673
	_go_fuzz_dep_.CoverTab[42596]++

							if w.req.Method != "HEAD" && func() bool {
//line /usr/local/go/src/net/http/server.go:1675
		_go_fuzz_dep_.CoverTab[42602]++
//line /usr/local/go/src/net/http/server.go:1675
		return w.contentLength != -1
//line /usr/local/go/src/net/http/server.go:1675
		// _ = "end of CoverTab[42602]"
//line /usr/local/go/src/net/http/server.go:1675
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1675
		_go_fuzz_dep_.CoverTab[42603]++
//line /usr/local/go/src/net/http/server.go:1675
		return w.bodyAllowed()
//line /usr/local/go/src/net/http/server.go:1675
		// _ = "end of CoverTab[42603]"
//line /usr/local/go/src/net/http/server.go:1675
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1675
		_go_fuzz_dep_.CoverTab[42604]++
//line /usr/local/go/src/net/http/server.go:1675
		return w.contentLength != w.written
//line /usr/local/go/src/net/http/server.go:1675
		// _ = "end of CoverTab[42604]"
//line /usr/local/go/src/net/http/server.go:1675
	}() {
//line /usr/local/go/src/net/http/server.go:1675
		_go_fuzz_dep_.CoverTab[42605]++

								return false
//line /usr/local/go/src/net/http/server.go:1677
		// _ = "end of CoverTab[42605]"
	} else {
//line /usr/local/go/src/net/http/server.go:1678
		_go_fuzz_dep_.CoverTab[42606]++
//line /usr/local/go/src/net/http/server.go:1678
		// _ = "end of CoverTab[42606]"
//line /usr/local/go/src/net/http/server.go:1678
	}
//line /usr/local/go/src/net/http/server.go:1678
	// _ = "end of CoverTab[42596]"
//line /usr/local/go/src/net/http/server.go:1678
	_go_fuzz_dep_.CoverTab[42597]++

//line /usr/local/go/src/net/http/server.go:1682
	if w.conn.werr != nil {
//line /usr/local/go/src/net/http/server.go:1682
		_go_fuzz_dep_.CoverTab[42607]++
								return false
//line /usr/local/go/src/net/http/server.go:1683
		// _ = "end of CoverTab[42607]"
	} else {
//line /usr/local/go/src/net/http/server.go:1684
		_go_fuzz_dep_.CoverTab[42608]++
//line /usr/local/go/src/net/http/server.go:1684
		// _ = "end of CoverTab[42608]"
//line /usr/local/go/src/net/http/server.go:1684
	}
//line /usr/local/go/src/net/http/server.go:1684
	// _ = "end of CoverTab[42597]"
//line /usr/local/go/src/net/http/server.go:1684
	_go_fuzz_dep_.CoverTab[42598]++

							if w.closedRequestBodyEarly() {
//line /usr/local/go/src/net/http/server.go:1686
		_go_fuzz_dep_.CoverTab[42609]++
								return false
//line /usr/local/go/src/net/http/server.go:1687
		// _ = "end of CoverTab[42609]"
	} else {
//line /usr/local/go/src/net/http/server.go:1688
		_go_fuzz_dep_.CoverTab[42610]++
//line /usr/local/go/src/net/http/server.go:1688
		// _ = "end of CoverTab[42610]"
//line /usr/local/go/src/net/http/server.go:1688
	}
//line /usr/local/go/src/net/http/server.go:1688
	// _ = "end of CoverTab[42598]"
//line /usr/local/go/src/net/http/server.go:1688
	_go_fuzz_dep_.CoverTab[42599]++

							return true
//line /usr/local/go/src/net/http/server.go:1690
	// _ = "end of CoverTab[42599]"
}

func (w *response) closedRequestBodyEarly() bool {
//line /usr/local/go/src/net/http/server.go:1693
	_go_fuzz_dep_.CoverTab[42611]++
							body, ok := w.req.Body.(*body)
							return ok && func() bool {
//line /usr/local/go/src/net/http/server.go:1695
		_go_fuzz_dep_.CoverTab[42612]++
//line /usr/local/go/src/net/http/server.go:1695
		return body.didEarlyClose()
//line /usr/local/go/src/net/http/server.go:1695
		// _ = "end of CoverTab[42612]"
//line /usr/local/go/src/net/http/server.go:1695
	}()
//line /usr/local/go/src/net/http/server.go:1695
	// _ = "end of CoverTab[42611]"
}

func (w *response) Flush() {
//line /usr/local/go/src/net/http/server.go:1698
	_go_fuzz_dep_.CoverTab[42613]++
							w.FlushError()
//line /usr/local/go/src/net/http/server.go:1699
	// _ = "end of CoverTab[42613]"
}

func (w *response) FlushError() error {
//line /usr/local/go/src/net/http/server.go:1702
	_go_fuzz_dep_.CoverTab[42614]++
							if !w.wroteHeader {
//line /usr/local/go/src/net/http/server.go:1703
		_go_fuzz_dep_.CoverTab[42617]++
								w.WriteHeader(StatusOK)
//line /usr/local/go/src/net/http/server.go:1704
		// _ = "end of CoverTab[42617]"
	} else {
//line /usr/local/go/src/net/http/server.go:1705
		_go_fuzz_dep_.CoverTab[42618]++
//line /usr/local/go/src/net/http/server.go:1705
		// _ = "end of CoverTab[42618]"
//line /usr/local/go/src/net/http/server.go:1705
	}
//line /usr/local/go/src/net/http/server.go:1705
	// _ = "end of CoverTab[42614]"
//line /usr/local/go/src/net/http/server.go:1705
	_go_fuzz_dep_.CoverTab[42615]++
							err := w.w.Flush()
							e2 := w.cw.flush()
							if err == nil {
//line /usr/local/go/src/net/http/server.go:1708
		_go_fuzz_dep_.CoverTab[42619]++
								err = e2
//line /usr/local/go/src/net/http/server.go:1709
		// _ = "end of CoverTab[42619]"
	} else {
//line /usr/local/go/src/net/http/server.go:1710
		_go_fuzz_dep_.CoverTab[42620]++
//line /usr/local/go/src/net/http/server.go:1710
		// _ = "end of CoverTab[42620]"
//line /usr/local/go/src/net/http/server.go:1710
	}
//line /usr/local/go/src/net/http/server.go:1710
	// _ = "end of CoverTab[42615]"
//line /usr/local/go/src/net/http/server.go:1710
	_go_fuzz_dep_.CoverTab[42616]++
							return err
//line /usr/local/go/src/net/http/server.go:1711
	// _ = "end of CoverTab[42616]"
}

func (c *conn) finalFlush() {
//line /usr/local/go/src/net/http/server.go:1714
	_go_fuzz_dep_.CoverTab[42621]++
							if c.bufr != nil {
//line /usr/local/go/src/net/http/server.go:1715
		_go_fuzz_dep_.CoverTab[42623]++

//line /usr/local/go/src/net/http/server.go:1718
		putBufioReader(c.bufr)
								c.bufr = nil
//line /usr/local/go/src/net/http/server.go:1719
		// _ = "end of CoverTab[42623]"
	} else {
//line /usr/local/go/src/net/http/server.go:1720
		_go_fuzz_dep_.CoverTab[42624]++
//line /usr/local/go/src/net/http/server.go:1720
		// _ = "end of CoverTab[42624]"
//line /usr/local/go/src/net/http/server.go:1720
	}
//line /usr/local/go/src/net/http/server.go:1720
	// _ = "end of CoverTab[42621]"
//line /usr/local/go/src/net/http/server.go:1720
	_go_fuzz_dep_.CoverTab[42622]++

							if c.bufw != nil {
//line /usr/local/go/src/net/http/server.go:1722
		_go_fuzz_dep_.CoverTab[42625]++
								c.bufw.Flush()

//line /usr/local/go/src/net/http/server.go:1726
		putBufioWriter(c.bufw)
								c.bufw = nil
//line /usr/local/go/src/net/http/server.go:1727
		// _ = "end of CoverTab[42625]"
	} else {
//line /usr/local/go/src/net/http/server.go:1728
		_go_fuzz_dep_.CoverTab[42626]++
//line /usr/local/go/src/net/http/server.go:1728
		// _ = "end of CoverTab[42626]"
//line /usr/local/go/src/net/http/server.go:1728
	}
//line /usr/local/go/src/net/http/server.go:1728
	// _ = "end of CoverTab[42622]"
}

// Close the connection.
func (c *conn) close() {
//line /usr/local/go/src/net/http/server.go:1732
	_go_fuzz_dep_.CoverTab[42627]++
							c.finalFlush()
							c.rwc.Close()
//line /usr/local/go/src/net/http/server.go:1734
	// _ = "end of CoverTab[42627]"
}

// rstAvoidanceDelay is the amount of time we sleep after closing the
//line /usr/local/go/src/net/http/server.go:1737
// write side of a TCP connection before closing the entire socket.
//line /usr/local/go/src/net/http/server.go:1737
// By sleeping, we increase the chances that the client sees our FIN
//line /usr/local/go/src/net/http/server.go:1737
// and processes its final data before they process the subsequent RST
//line /usr/local/go/src/net/http/server.go:1737
// from closing a connection with known unread data.
//line /usr/local/go/src/net/http/server.go:1737
// This RST seems to occur mostly on BSD systems. (And Windows?)
//line /usr/local/go/src/net/http/server.go:1737
// This timeout is somewhat arbitrary (~latency around the planet).
//line /usr/local/go/src/net/http/server.go:1744
const rstAvoidanceDelay = 500 * time.Millisecond

type closeWriter interface {
	CloseWrite() error
}

var _ closeWriter = (*net.TCPConn)(nil)

// closeWrite flushes any outstanding data and sends a FIN packet (if
//line /usr/local/go/src/net/http/server.go:1752
// client is connected via TCP), signaling that we're done. We then
//line /usr/local/go/src/net/http/server.go:1752
// pause for a bit, hoping the client processes it before any
//line /usr/local/go/src/net/http/server.go:1752
// subsequent RST.
//line /usr/local/go/src/net/http/server.go:1752
//
//line /usr/local/go/src/net/http/server.go:1752
// See https://golang.org/issue/3595
//line /usr/local/go/src/net/http/server.go:1758
func (c *conn) closeWriteAndWait() {
//line /usr/local/go/src/net/http/server.go:1758
	_go_fuzz_dep_.CoverTab[42628]++
							c.finalFlush()
							if tcp, ok := c.rwc.(closeWriter); ok {
//line /usr/local/go/src/net/http/server.go:1760
		_go_fuzz_dep_.CoverTab[42630]++
								tcp.CloseWrite()
//line /usr/local/go/src/net/http/server.go:1761
		// _ = "end of CoverTab[42630]"
	} else {
//line /usr/local/go/src/net/http/server.go:1762
		_go_fuzz_dep_.CoverTab[42631]++
//line /usr/local/go/src/net/http/server.go:1762
		// _ = "end of CoverTab[42631]"
//line /usr/local/go/src/net/http/server.go:1762
	}
//line /usr/local/go/src/net/http/server.go:1762
	// _ = "end of CoverTab[42628]"
//line /usr/local/go/src/net/http/server.go:1762
	_go_fuzz_dep_.CoverTab[42629]++
							time.Sleep(rstAvoidanceDelay)
//line /usr/local/go/src/net/http/server.go:1763
	// _ = "end of CoverTab[42629]"
}

// validNextProto reports whether the proto is a valid ALPN protocol name.
//line /usr/local/go/src/net/http/server.go:1766
// Everything is valid except the empty string and built-in protocol types,
//line /usr/local/go/src/net/http/server.go:1766
// so that those can't be overridden with alternate implementations.
//line /usr/local/go/src/net/http/server.go:1769
func validNextProto(proto string) bool {
//line /usr/local/go/src/net/http/server.go:1769
	_go_fuzz_dep_.CoverTab[42632]++
							switch proto {
	case "", "http/1.1", "http/1.0":
//line /usr/local/go/src/net/http/server.go:1771
		_go_fuzz_dep_.CoverTab[42634]++
								return false
//line /usr/local/go/src/net/http/server.go:1772
		// _ = "end of CoverTab[42634]"
//line /usr/local/go/src/net/http/server.go:1772
	default:
//line /usr/local/go/src/net/http/server.go:1772
		_go_fuzz_dep_.CoverTab[42635]++
//line /usr/local/go/src/net/http/server.go:1772
		// _ = "end of CoverTab[42635]"
	}
//line /usr/local/go/src/net/http/server.go:1773
	// _ = "end of CoverTab[42632]"
//line /usr/local/go/src/net/http/server.go:1773
	_go_fuzz_dep_.CoverTab[42633]++
							return true
//line /usr/local/go/src/net/http/server.go:1774
	// _ = "end of CoverTab[42633]"
}

const (
	runHooks	= true
	skipHooks	= false
)

func (c *conn) setState(nc net.Conn, state ConnState, runHook bool) {
//line /usr/local/go/src/net/http/server.go:1782
	_go_fuzz_dep_.CoverTab[42636]++
							srv := c.server
							switch state {
	case StateNew:
//line /usr/local/go/src/net/http/server.go:1785
		_go_fuzz_dep_.CoverTab[42640]++
								srv.trackConn(c, true)
//line /usr/local/go/src/net/http/server.go:1786
		// _ = "end of CoverTab[42640]"
	case StateHijacked, StateClosed:
//line /usr/local/go/src/net/http/server.go:1787
		_go_fuzz_dep_.CoverTab[42641]++
								srv.trackConn(c, false)
//line /usr/local/go/src/net/http/server.go:1788
		// _ = "end of CoverTab[42641]"
//line /usr/local/go/src/net/http/server.go:1788
	default:
//line /usr/local/go/src/net/http/server.go:1788
		_go_fuzz_dep_.CoverTab[42642]++
//line /usr/local/go/src/net/http/server.go:1788
		// _ = "end of CoverTab[42642]"
	}
//line /usr/local/go/src/net/http/server.go:1789
	// _ = "end of CoverTab[42636]"
//line /usr/local/go/src/net/http/server.go:1789
	_go_fuzz_dep_.CoverTab[42637]++
							if state > 0xff || func() bool {
//line /usr/local/go/src/net/http/server.go:1790
		_go_fuzz_dep_.CoverTab[42643]++
//line /usr/local/go/src/net/http/server.go:1790
		return state < 0
//line /usr/local/go/src/net/http/server.go:1790
		// _ = "end of CoverTab[42643]"
//line /usr/local/go/src/net/http/server.go:1790
	}() {
//line /usr/local/go/src/net/http/server.go:1790
		_go_fuzz_dep_.CoverTab[42644]++
								panic("internal error")
//line /usr/local/go/src/net/http/server.go:1791
		// _ = "end of CoverTab[42644]"
	} else {
//line /usr/local/go/src/net/http/server.go:1792
		_go_fuzz_dep_.CoverTab[42645]++
//line /usr/local/go/src/net/http/server.go:1792
		// _ = "end of CoverTab[42645]"
//line /usr/local/go/src/net/http/server.go:1792
	}
//line /usr/local/go/src/net/http/server.go:1792
	// _ = "end of CoverTab[42637]"
//line /usr/local/go/src/net/http/server.go:1792
	_go_fuzz_dep_.CoverTab[42638]++
							packedState := uint64(time.Now().Unix()<<8) | uint64(state)
							c.curState.Store(packedState)
							if !runHook {
//line /usr/local/go/src/net/http/server.go:1795
		_go_fuzz_dep_.CoverTab[42646]++
								return
//line /usr/local/go/src/net/http/server.go:1796
		// _ = "end of CoverTab[42646]"
	} else {
//line /usr/local/go/src/net/http/server.go:1797
		_go_fuzz_dep_.CoverTab[42647]++
//line /usr/local/go/src/net/http/server.go:1797
		// _ = "end of CoverTab[42647]"
//line /usr/local/go/src/net/http/server.go:1797
	}
//line /usr/local/go/src/net/http/server.go:1797
	// _ = "end of CoverTab[42638]"
//line /usr/local/go/src/net/http/server.go:1797
	_go_fuzz_dep_.CoverTab[42639]++
							if hook := srv.ConnState; hook != nil {
//line /usr/local/go/src/net/http/server.go:1798
		_go_fuzz_dep_.CoverTab[42648]++
								hook(nc, state)
//line /usr/local/go/src/net/http/server.go:1799
		// _ = "end of CoverTab[42648]"
	} else {
//line /usr/local/go/src/net/http/server.go:1800
		_go_fuzz_dep_.CoverTab[42649]++
//line /usr/local/go/src/net/http/server.go:1800
		// _ = "end of CoverTab[42649]"
//line /usr/local/go/src/net/http/server.go:1800
	}
//line /usr/local/go/src/net/http/server.go:1800
	// _ = "end of CoverTab[42639]"
}

func (c *conn) getState() (state ConnState, unixSec int64) {
//line /usr/local/go/src/net/http/server.go:1803
	_go_fuzz_dep_.CoverTab[42650]++
							packedState := c.curState.Load()
							return ConnState(packedState & 0xff), int64(packedState >> 8)
//line /usr/local/go/src/net/http/server.go:1805
	// _ = "end of CoverTab[42650]"
}

// badRequestError is a literal string (used by in the server in HTML,
//line /usr/local/go/src/net/http/server.go:1808
// unescaped) to tell the user why their request was bad. It should
//line /usr/local/go/src/net/http/server.go:1808
// be plain text without user info or other embedded errors.
//line /usr/local/go/src/net/http/server.go:1811
func badRequestError(e string) error {
//line /usr/local/go/src/net/http/server.go:1811
	_go_fuzz_dep_.CoverTab[42651]++
//line /usr/local/go/src/net/http/server.go:1811
	return statusError{StatusBadRequest, e}
//line /usr/local/go/src/net/http/server.go:1811
	// _ = "end of CoverTab[42651]"
//line /usr/local/go/src/net/http/server.go:1811
}

// statusError is an error used to respond to a request with an HTTP status.
//line /usr/local/go/src/net/http/server.go:1813
// The text should be plain text without user info or other embedded errors.
//line /usr/local/go/src/net/http/server.go:1815
type statusError struct {
	code	int
	text	string
}

func (e statusError) Error() string {
//line /usr/local/go/src/net/http/server.go:1820
	_go_fuzz_dep_.CoverTab[42652]++
//line /usr/local/go/src/net/http/server.go:1820
	return StatusText(e.code) + ": " + e.text
//line /usr/local/go/src/net/http/server.go:1820
	// _ = "end of CoverTab[42652]"
//line /usr/local/go/src/net/http/server.go:1820
}

// ErrAbortHandler is a sentinel panic value to abort a handler.
//line /usr/local/go/src/net/http/server.go:1822
// While any panic from ServeHTTP aborts the response to the client,
//line /usr/local/go/src/net/http/server.go:1822
// panicking with ErrAbortHandler also suppresses logging of a stack
//line /usr/local/go/src/net/http/server.go:1822
// trace to the server's error log.
//line /usr/local/go/src/net/http/server.go:1826
var ErrAbortHandler = errors.New("net/http: abort Handler")

// isCommonNetReadError reports whether err is a common error
//line /usr/local/go/src/net/http/server.go:1828
// encountered during reading a request off the network when the
//line /usr/local/go/src/net/http/server.go:1828
// client has gone away or had its read fail somehow. This is used to
//line /usr/local/go/src/net/http/server.go:1828
// determine which logs are interesting enough to log about.
//line /usr/local/go/src/net/http/server.go:1832
func isCommonNetReadError(err error) bool {
//line /usr/local/go/src/net/http/server.go:1832
	_go_fuzz_dep_.CoverTab[42653]++
							if err == io.EOF {
//line /usr/local/go/src/net/http/server.go:1833
		_go_fuzz_dep_.CoverTab[42657]++
								return true
//line /usr/local/go/src/net/http/server.go:1834
		// _ = "end of CoverTab[42657]"
	} else {
//line /usr/local/go/src/net/http/server.go:1835
		_go_fuzz_dep_.CoverTab[42658]++
//line /usr/local/go/src/net/http/server.go:1835
		// _ = "end of CoverTab[42658]"
//line /usr/local/go/src/net/http/server.go:1835
	}
//line /usr/local/go/src/net/http/server.go:1835
	// _ = "end of CoverTab[42653]"
//line /usr/local/go/src/net/http/server.go:1835
	_go_fuzz_dep_.CoverTab[42654]++
							if neterr, ok := err.(net.Error); ok && func() bool {
//line /usr/local/go/src/net/http/server.go:1836
		_go_fuzz_dep_.CoverTab[42659]++
//line /usr/local/go/src/net/http/server.go:1836
		return neterr.Timeout()
//line /usr/local/go/src/net/http/server.go:1836
		// _ = "end of CoverTab[42659]"
//line /usr/local/go/src/net/http/server.go:1836
	}() {
//line /usr/local/go/src/net/http/server.go:1836
		_go_fuzz_dep_.CoverTab[42660]++
								return true
//line /usr/local/go/src/net/http/server.go:1837
		// _ = "end of CoverTab[42660]"
	} else {
//line /usr/local/go/src/net/http/server.go:1838
		_go_fuzz_dep_.CoverTab[42661]++
//line /usr/local/go/src/net/http/server.go:1838
		// _ = "end of CoverTab[42661]"
//line /usr/local/go/src/net/http/server.go:1838
	}
//line /usr/local/go/src/net/http/server.go:1838
	// _ = "end of CoverTab[42654]"
//line /usr/local/go/src/net/http/server.go:1838
	_go_fuzz_dep_.CoverTab[42655]++
							if oe, ok := err.(*net.OpError); ok && func() bool {
//line /usr/local/go/src/net/http/server.go:1839
		_go_fuzz_dep_.CoverTab[42662]++
//line /usr/local/go/src/net/http/server.go:1839
		return oe.Op == "read"
//line /usr/local/go/src/net/http/server.go:1839
		// _ = "end of CoverTab[42662]"
//line /usr/local/go/src/net/http/server.go:1839
	}() {
//line /usr/local/go/src/net/http/server.go:1839
		_go_fuzz_dep_.CoverTab[42663]++
								return true
//line /usr/local/go/src/net/http/server.go:1840
		// _ = "end of CoverTab[42663]"
	} else {
//line /usr/local/go/src/net/http/server.go:1841
		_go_fuzz_dep_.CoverTab[42664]++
//line /usr/local/go/src/net/http/server.go:1841
		// _ = "end of CoverTab[42664]"
//line /usr/local/go/src/net/http/server.go:1841
	}
//line /usr/local/go/src/net/http/server.go:1841
	// _ = "end of CoverTab[42655]"
//line /usr/local/go/src/net/http/server.go:1841
	_go_fuzz_dep_.CoverTab[42656]++
							return false
//line /usr/local/go/src/net/http/server.go:1842
	// _ = "end of CoverTab[42656]"
}

// Serve a new connection.
func (c *conn) serve(ctx context.Context) {
//line /usr/local/go/src/net/http/server.go:1846
	_go_fuzz_dep_.CoverTab[42665]++
							c.remoteAddr = c.rwc.RemoteAddr().String()
							ctx = context.WithValue(ctx, LocalAddrContextKey, c.rwc.LocalAddr())
							var inFlightResponse *response
							defer func() {
//line /usr/local/go/src/net/http/server.go:1850
		_go_fuzz_dep_.CoverTab[42668]++
								if err := recover(); err != nil && func() bool {
//line /usr/local/go/src/net/http/server.go:1851
			_go_fuzz_dep_.CoverTab[42671]++
//line /usr/local/go/src/net/http/server.go:1851
			return err != ErrAbortHandler
//line /usr/local/go/src/net/http/server.go:1851
			// _ = "end of CoverTab[42671]"
//line /usr/local/go/src/net/http/server.go:1851
		}() {
//line /usr/local/go/src/net/http/server.go:1851
			_go_fuzz_dep_.CoverTab[42672]++
									const size = 64 << 10
									buf := make([]byte, size)
									buf = buf[:runtime.Stack(buf, false)]
									c.server.logf("http: panic serving %v: %v\n%s", c.remoteAddr, err, buf)
//line /usr/local/go/src/net/http/server.go:1855
			// _ = "end of CoverTab[42672]"
		} else {
//line /usr/local/go/src/net/http/server.go:1856
			_go_fuzz_dep_.CoverTab[42673]++
//line /usr/local/go/src/net/http/server.go:1856
			// _ = "end of CoverTab[42673]"
//line /usr/local/go/src/net/http/server.go:1856
		}
//line /usr/local/go/src/net/http/server.go:1856
		// _ = "end of CoverTab[42668]"
//line /usr/local/go/src/net/http/server.go:1856
		_go_fuzz_dep_.CoverTab[42669]++
								if inFlightResponse != nil {
//line /usr/local/go/src/net/http/server.go:1857
			_go_fuzz_dep_.CoverTab[42674]++
									inFlightResponse.cancelCtx()
//line /usr/local/go/src/net/http/server.go:1858
			// _ = "end of CoverTab[42674]"
		} else {
//line /usr/local/go/src/net/http/server.go:1859
			_go_fuzz_dep_.CoverTab[42675]++
//line /usr/local/go/src/net/http/server.go:1859
			// _ = "end of CoverTab[42675]"
//line /usr/local/go/src/net/http/server.go:1859
		}
//line /usr/local/go/src/net/http/server.go:1859
		// _ = "end of CoverTab[42669]"
//line /usr/local/go/src/net/http/server.go:1859
		_go_fuzz_dep_.CoverTab[42670]++
								if !c.hijacked() {
//line /usr/local/go/src/net/http/server.go:1860
			_go_fuzz_dep_.CoverTab[42676]++
									if inFlightResponse != nil {
//line /usr/local/go/src/net/http/server.go:1861
				_go_fuzz_dep_.CoverTab[42678]++
										inFlightResponse.conn.r.abortPendingRead()
										inFlightResponse.reqBody.Close()
//line /usr/local/go/src/net/http/server.go:1863
				// _ = "end of CoverTab[42678]"
			} else {
//line /usr/local/go/src/net/http/server.go:1864
				_go_fuzz_dep_.CoverTab[42679]++
//line /usr/local/go/src/net/http/server.go:1864
				// _ = "end of CoverTab[42679]"
//line /usr/local/go/src/net/http/server.go:1864
			}
//line /usr/local/go/src/net/http/server.go:1864
			// _ = "end of CoverTab[42676]"
//line /usr/local/go/src/net/http/server.go:1864
			_go_fuzz_dep_.CoverTab[42677]++
									c.close()
									c.setState(c.rwc, StateClosed, runHooks)
//line /usr/local/go/src/net/http/server.go:1866
			// _ = "end of CoverTab[42677]"
		} else {
//line /usr/local/go/src/net/http/server.go:1867
			_go_fuzz_dep_.CoverTab[42680]++
//line /usr/local/go/src/net/http/server.go:1867
			// _ = "end of CoverTab[42680]"
//line /usr/local/go/src/net/http/server.go:1867
		}
//line /usr/local/go/src/net/http/server.go:1867
		// _ = "end of CoverTab[42670]"
	}()
//line /usr/local/go/src/net/http/server.go:1868
	// _ = "end of CoverTab[42665]"
//line /usr/local/go/src/net/http/server.go:1868
	_go_fuzz_dep_.CoverTab[42666]++

							if tlsConn, ok := c.rwc.(*tls.Conn); ok {
//line /usr/local/go/src/net/http/server.go:1870
		_go_fuzz_dep_.CoverTab[42681]++
								tlsTO := c.server.tlsHandshakeTimeout()
								if tlsTO > 0 {
//line /usr/local/go/src/net/http/server.go:1872
			_go_fuzz_dep_.CoverTab[42685]++
									dl := time.Now().Add(tlsTO)
									c.rwc.SetReadDeadline(dl)
									c.rwc.SetWriteDeadline(dl)
//line /usr/local/go/src/net/http/server.go:1875
			// _ = "end of CoverTab[42685]"
		} else {
//line /usr/local/go/src/net/http/server.go:1876
			_go_fuzz_dep_.CoverTab[42686]++
//line /usr/local/go/src/net/http/server.go:1876
			// _ = "end of CoverTab[42686]"
//line /usr/local/go/src/net/http/server.go:1876
		}
//line /usr/local/go/src/net/http/server.go:1876
		// _ = "end of CoverTab[42681]"
//line /usr/local/go/src/net/http/server.go:1876
		_go_fuzz_dep_.CoverTab[42682]++
								if err := tlsConn.HandshakeContext(ctx); err != nil {
//line /usr/local/go/src/net/http/server.go:1877
			_go_fuzz_dep_.CoverTab[42687]++

//line /usr/local/go/src/net/http/server.go:1881
			if re, ok := err.(tls.RecordHeaderError); ok && func() bool {
//line /usr/local/go/src/net/http/server.go:1881
				_go_fuzz_dep_.CoverTab[42689]++
//line /usr/local/go/src/net/http/server.go:1881
				return re.Conn != nil
//line /usr/local/go/src/net/http/server.go:1881
				// _ = "end of CoverTab[42689]"
//line /usr/local/go/src/net/http/server.go:1881
			}() && func() bool {
//line /usr/local/go/src/net/http/server.go:1881
				_go_fuzz_dep_.CoverTab[42690]++
//line /usr/local/go/src/net/http/server.go:1881
				return tlsRecordHeaderLooksLikeHTTP(re.RecordHeader)
//line /usr/local/go/src/net/http/server.go:1881
				// _ = "end of CoverTab[42690]"
//line /usr/local/go/src/net/http/server.go:1881
			}() {
//line /usr/local/go/src/net/http/server.go:1881
				_go_fuzz_dep_.CoverTab[42691]++
										io.WriteString(re.Conn, "HTTP/1.0 400 Bad Request\r\n\r\nClient sent an HTTP request to an HTTPS server.\n")
										re.Conn.Close()
										return
//line /usr/local/go/src/net/http/server.go:1884
				// _ = "end of CoverTab[42691]"
			} else {
//line /usr/local/go/src/net/http/server.go:1885
				_go_fuzz_dep_.CoverTab[42692]++
//line /usr/local/go/src/net/http/server.go:1885
				// _ = "end of CoverTab[42692]"
//line /usr/local/go/src/net/http/server.go:1885
			}
//line /usr/local/go/src/net/http/server.go:1885
			// _ = "end of CoverTab[42687]"
//line /usr/local/go/src/net/http/server.go:1885
			_go_fuzz_dep_.CoverTab[42688]++
									c.server.logf("http: TLS handshake error from %s: %v", c.rwc.RemoteAddr(), err)
									return
//line /usr/local/go/src/net/http/server.go:1887
			// _ = "end of CoverTab[42688]"
		} else {
//line /usr/local/go/src/net/http/server.go:1888
			_go_fuzz_dep_.CoverTab[42693]++
//line /usr/local/go/src/net/http/server.go:1888
			// _ = "end of CoverTab[42693]"
//line /usr/local/go/src/net/http/server.go:1888
		}
//line /usr/local/go/src/net/http/server.go:1888
		// _ = "end of CoverTab[42682]"
//line /usr/local/go/src/net/http/server.go:1888
		_go_fuzz_dep_.CoverTab[42683]++

								if tlsTO > 0 {
//line /usr/local/go/src/net/http/server.go:1890
			_go_fuzz_dep_.CoverTab[42694]++
									c.rwc.SetReadDeadline(time.Time{})
									c.rwc.SetWriteDeadline(time.Time{})
//line /usr/local/go/src/net/http/server.go:1892
			// _ = "end of CoverTab[42694]"
		} else {
//line /usr/local/go/src/net/http/server.go:1893
			_go_fuzz_dep_.CoverTab[42695]++
//line /usr/local/go/src/net/http/server.go:1893
			// _ = "end of CoverTab[42695]"
//line /usr/local/go/src/net/http/server.go:1893
		}
//line /usr/local/go/src/net/http/server.go:1893
		// _ = "end of CoverTab[42683]"
//line /usr/local/go/src/net/http/server.go:1893
		_go_fuzz_dep_.CoverTab[42684]++
								c.tlsState = new(tls.ConnectionState)
								*c.tlsState = tlsConn.ConnectionState()
								if proto := c.tlsState.NegotiatedProtocol; validNextProto(proto) {
//line /usr/local/go/src/net/http/server.go:1896
			_go_fuzz_dep_.CoverTab[42696]++
									if fn := c.server.TLSNextProto[proto]; fn != nil {
//line /usr/local/go/src/net/http/server.go:1897
				_go_fuzz_dep_.CoverTab[42698]++
										h := initALPNRequest{ctx, tlsConn, serverHandler{c.server}}

//line /usr/local/go/src/net/http/server.go:1902
				c.setState(c.rwc, StateActive, skipHooks)
										fn(c.server, tlsConn, h)
//line /usr/local/go/src/net/http/server.go:1903
				// _ = "end of CoverTab[42698]"
			} else {
//line /usr/local/go/src/net/http/server.go:1904
				_go_fuzz_dep_.CoverTab[42699]++
//line /usr/local/go/src/net/http/server.go:1904
				// _ = "end of CoverTab[42699]"
//line /usr/local/go/src/net/http/server.go:1904
			}
//line /usr/local/go/src/net/http/server.go:1904
			// _ = "end of CoverTab[42696]"
//line /usr/local/go/src/net/http/server.go:1904
			_go_fuzz_dep_.CoverTab[42697]++
									return
//line /usr/local/go/src/net/http/server.go:1905
			// _ = "end of CoverTab[42697]"
		} else {
//line /usr/local/go/src/net/http/server.go:1906
			_go_fuzz_dep_.CoverTab[42700]++
//line /usr/local/go/src/net/http/server.go:1906
			// _ = "end of CoverTab[42700]"
//line /usr/local/go/src/net/http/server.go:1906
		}
//line /usr/local/go/src/net/http/server.go:1906
		// _ = "end of CoverTab[42684]"
	} else {
//line /usr/local/go/src/net/http/server.go:1907
		_go_fuzz_dep_.CoverTab[42701]++
//line /usr/local/go/src/net/http/server.go:1907
		// _ = "end of CoverTab[42701]"
//line /usr/local/go/src/net/http/server.go:1907
	}
//line /usr/local/go/src/net/http/server.go:1907
	// _ = "end of CoverTab[42666]"
//line /usr/local/go/src/net/http/server.go:1907
	_go_fuzz_dep_.CoverTab[42667]++

//line /usr/local/go/src/net/http/server.go:1911
	ctx, cancelCtx := context.WithCancel(ctx)
	c.cancelCtx = cancelCtx
	defer cancelCtx()

	c.r = &connReader{conn: c}
	c.bufr = newBufioReader(c.r)
	c.bufw = newBufioWriterSize(checkConnErrorWriter{c}, 4<<10)

	for {
//line /usr/local/go/src/net/http/server.go:1919
		_go_fuzz_dep_.CoverTab[42702]++
								w, err := c.readRequest(ctx)
								if c.r.remain != c.server.initialReadLimitSize() {
//line /usr/local/go/src/net/http/server.go:1921
			_go_fuzz_dep_.CoverTab[42712]++

									c.setState(c.rwc, StateActive, runHooks)
//line /usr/local/go/src/net/http/server.go:1923
			// _ = "end of CoverTab[42712]"
		} else {
//line /usr/local/go/src/net/http/server.go:1924
			_go_fuzz_dep_.CoverTab[42713]++
//line /usr/local/go/src/net/http/server.go:1924
			// _ = "end of CoverTab[42713]"
//line /usr/local/go/src/net/http/server.go:1924
		}
//line /usr/local/go/src/net/http/server.go:1924
		// _ = "end of CoverTab[42702]"
//line /usr/local/go/src/net/http/server.go:1924
		_go_fuzz_dep_.CoverTab[42703]++
								if err != nil {
//line /usr/local/go/src/net/http/server.go:1925
			_go_fuzz_dep_.CoverTab[42714]++
									const errorHeaders = "\r\nContent-Type: text/plain; charset=utf-8\r\nConnection: close\r\n\r\n"

									switch {
			case err == errTooLarge:
//line /usr/local/go/src/net/http/server.go:1929
				_go_fuzz_dep_.CoverTab[42715]++
				// Their HTTP client may or may not be
				// able to read this if we're
				// responding to them and hanging up
				// while they're still writing their
										// request. Undefined behavior.
										const publicErr = "431 Request Header Fields Too Large"
										fmt.Fprintf(c.rwc, "HTTP/1.1 "+publicErr+errorHeaders+publicErr)
										c.closeWriteAndWait()
										return
//line /usr/local/go/src/net/http/server.go:1938
				// _ = "end of CoverTab[42715]"

			case isUnsupportedTEError(err):
//line /usr/local/go/src/net/http/server.go:1940
				_go_fuzz_dep_.CoverTab[42716]++

//line /usr/local/go/src/net/http/server.go:1945
				code := StatusNotImplemented

//line /usr/local/go/src/net/http/server.go:1949
				fmt.Fprintf(c.rwc, "HTTP/1.1 %d %s%sUnsupported transfer encoding", code, StatusText(code), errorHeaders)
										return
//line /usr/local/go/src/net/http/server.go:1950
				// _ = "end of CoverTab[42716]"

			case isCommonNetReadError(err):
//line /usr/local/go/src/net/http/server.go:1952
				_go_fuzz_dep_.CoverTab[42717]++
										return
//line /usr/local/go/src/net/http/server.go:1953
				// _ = "end of CoverTab[42717]"

			default:
//line /usr/local/go/src/net/http/server.go:1955
				_go_fuzz_dep_.CoverTab[42718]++
										if v, ok := err.(statusError); ok {
//line /usr/local/go/src/net/http/server.go:1956
					_go_fuzz_dep_.CoverTab[42720]++
											fmt.Fprintf(c.rwc, "HTTP/1.1 %d %s: %s%s%d %s: %s", v.code, StatusText(v.code), v.text, errorHeaders, v.code, StatusText(v.code), v.text)
											return
//line /usr/local/go/src/net/http/server.go:1958
					// _ = "end of CoverTab[42720]"
				} else {
//line /usr/local/go/src/net/http/server.go:1959
					_go_fuzz_dep_.CoverTab[42721]++
//line /usr/local/go/src/net/http/server.go:1959
					// _ = "end of CoverTab[42721]"
//line /usr/local/go/src/net/http/server.go:1959
				}
//line /usr/local/go/src/net/http/server.go:1959
				// _ = "end of CoverTab[42718]"
//line /usr/local/go/src/net/http/server.go:1959
				_go_fuzz_dep_.CoverTab[42719]++
										publicErr := "400 Bad Request"
										fmt.Fprintf(c.rwc, "HTTP/1.1 "+publicErr+errorHeaders+publicErr)
										return
//line /usr/local/go/src/net/http/server.go:1962
				// _ = "end of CoverTab[42719]"
			}
//line /usr/local/go/src/net/http/server.go:1963
			// _ = "end of CoverTab[42714]"
		} else {
//line /usr/local/go/src/net/http/server.go:1964
			_go_fuzz_dep_.CoverTab[42722]++
//line /usr/local/go/src/net/http/server.go:1964
			// _ = "end of CoverTab[42722]"
//line /usr/local/go/src/net/http/server.go:1964
		}
//line /usr/local/go/src/net/http/server.go:1964
		// _ = "end of CoverTab[42703]"
//line /usr/local/go/src/net/http/server.go:1964
		_go_fuzz_dep_.CoverTab[42704]++

//line /usr/local/go/src/net/http/server.go:1967
		req := w.req
		if req.expectsContinue() {
//line /usr/local/go/src/net/http/server.go:1968
			_go_fuzz_dep_.CoverTab[42723]++
									if req.ProtoAtLeast(1, 1) && func() bool {
//line /usr/local/go/src/net/http/server.go:1969
				_go_fuzz_dep_.CoverTab[42724]++
//line /usr/local/go/src/net/http/server.go:1969
				return req.ContentLength != 0
//line /usr/local/go/src/net/http/server.go:1969
				// _ = "end of CoverTab[42724]"
//line /usr/local/go/src/net/http/server.go:1969
			}() {
//line /usr/local/go/src/net/http/server.go:1969
				_go_fuzz_dep_.CoverTab[42725]++

										req.Body = &expectContinueReader{readCloser: req.Body, resp: w}
										w.canWriteContinue.Store(true)
//line /usr/local/go/src/net/http/server.go:1972
				// _ = "end of CoverTab[42725]"
			} else {
//line /usr/local/go/src/net/http/server.go:1973
				_go_fuzz_dep_.CoverTab[42726]++
//line /usr/local/go/src/net/http/server.go:1973
				// _ = "end of CoverTab[42726]"
//line /usr/local/go/src/net/http/server.go:1973
			}
//line /usr/local/go/src/net/http/server.go:1973
			// _ = "end of CoverTab[42723]"
		} else {
//line /usr/local/go/src/net/http/server.go:1974
			_go_fuzz_dep_.CoverTab[42727]++
//line /usr/local/go/src/net/http/server.go:1974
			if req.Header.get("Expect") != "" {
//line /usr/local/go/src/net/http/server.go:1974
				_go_fuzz_dep_.CoverTab[42728]++
										w.sendExpectationFailed()
										return
//line /usr/local/go/src/net/http/server.go:1976
				// _ = "end of CoverTab[42728]"
			} else {
//line /usr/local/go/src/net/http/server.go:1977
				_go_fuzz_dep_.CoverTab[42729]++
//line /usr/local/go/src/net/http/server.go:1977
				// _ = "end of CoverTab[42729]"
//line /usr/local/go/src/net/http/server.go:1977
			}
//line /usr/local/go/src/net/http/server.go:1977
			// _ = "end of CoverTab[42727]"
//line /usr/local/go/src/net/http/server.go:1977
		}
//line /usr/local/go/src/net/http/server.go:1977
		// _ = "end of CoverTab[42704]"
//line /usr/local/go/src/net/http/server.go:1977
		_go_fuzz_dep_.CoverTab[42705]++

								c.curReq.Store(w)

								if requestBodyRemains(req.Body) {
//line /usr/local/go/src/net/http/server.go:1981
			_go_fuzz_dep_.CoverTab[42730]++
									registerOnHitEOF(req.Body, w.conn.r.startBackgroundRead)
//line /usr/local/go/src/net/http/server.go:1982
			// _ = "end of CoverTab[42730]"
		} else {
//line /usr/local/go/src/net/http/server.go:1983
			_go_fuzz_dep_.CoverTab[42731]++
									w.conn.r.startBackgroundRead()
//line /usr/local/go/src/net/http/server.go:1984
			// _ = "end of CoverTab[42731]"
		}
//line /usr/local/go/src/net/http/server.go:1985
		// _ = "end of CoverTab[42705]"
//line /usr/local/go/src/net/http/server.go:1985
		_go_fuzz_dep_.CoverTab[42706]++

//line /usr/local/go/src/net/http/server.go:1994
		inFlightResponse = w
		serverHandler{c.server}.ServeHTTP(w, w.req)
		inFlightResponse = nil
		w.cancelCtx()
		if c.hijacked() {
//line /usr/local/go/src/net/http/server.go:1998
			_go_fuzz_dep_.CoverTab[42732]++
									return
//line /usr/local/go/src/net/http/server.go:1999
			// _ = "end of CoverTab[42732]"
		} else {
//line /usr/local/go/src/net/http/server.go:2000
			_go_fuzz_dep_.CoverTab[42733]++
//line /usr/local/go/src/net/http/server.go:2000
			// _ = "end of CoverTab[42733]"
//line /usr/local/go/src/net/http/server.go:2000
		}
//line /usr/local/go/src/net/http/server.go:2000
		// _ = "end of CoverTab[42706]"
//line /usr/local/go/src/net/http/server.go:2000
		_go_fuzz_dep_.CoverTab[42707]++
								w.finishRequest()
								c.rwc.SetWriteDeadline(time.Time{})
								if !w.shouldReuseConnection() {
//line /usr/local/go/src/net/http/server.go:2003
			_go_fuzz_dep_.CoverTab[42734]++
									if w.requestBodyLimitHit || func() bool {
//line /usr/local/go/src/net/http/server.go:2004
				_go_fuzz_dep_.CoverTab[42736]++
//line /usr/local/go/src/net/http/server.go:2004
				return w.closedRequestBodyEarly()
//line /usr/local/go/src/net/http/server.go:2004
				// _ = "end of CoverTab[42736]"
//line /usr/local/go/src/net/http/server.go:2004
			}() {
//line /usr/local/go/src/net/http/server.go:2004
				_go_fuzz_dep_.CoverTab[42737]++
										c.closeWriteAndWait()
//line /usr/local/go/src/net/http/server.go:2005
				// _ = "end of CoverTab[42737]"
			} else {
//line /usr/local/go/src/net/http/server.go:2006
				_go_fuzz_dep_.CoverTab[42738]++
//line /usr/local/go/src/net/http/server.go:2006
				// _ = "end of CoverTab[42738]"
//line /usr/local/go/src/net/http/server.go:2006
			}
//line /usr/local/go/src/net/http/server.go:2006
			// _ = "end of CoverTab[42734]"
//line /usr/local/go/src/net/http/server.go:2006
			_go_fuzz_dep_.CoverTab[42735]++
									return
//line /usr/local/go/src/net/http/server.go:2007
			// _ = "end of CoverTab[42735]"
		} else {
//line /usr/local/go/src/net/http/server.go:2008
			_go_fuzz_dep_.CoverTab[42739]++
//line /usr/local/go/src/net/http/server.go:2008
			// _ = "end of CoverTab[42739]"
//line /usr/local/go/src/net/http/server.go:2008
		}
//line /usr/local/go/src/net/http/server.go:2008
		// _ = "end of CoverTab[42707]"
//line /usr/local/go/src/net/http/server.go:2008
		_go_fuzz_dep_.CoverTab[42708]++
								c.setState(c.rwc, StateIdle, runHooks)
								c.curReq.Store(nil)

								if !w.conn.server.doKeepAlives() {
//line /usr/local/go/src/net/http/server.go:2012
			_go_fuzz_dep_.CoverTab[42740]++

//line /usr/local/go/src/net/http/server.go:2017
			return
//line /usr/local/go/src/net/http/server.go:2017
			// _ = "end of CoverTab[42740]"
		} else {
//line /usr/local/go/src/net/http/server.go:2018
			_go_fuzz_dep_.CoverTab[42741]++
//line /usr/local/go/src/net/http/server.go:2018
			// _ = "end of CoverTab[42741]"
//line /usr/local/go/src/net/http/server.go:2018
		}
//line /usr/local/go/src/net/http/server.go:2018
		// _ = "end of CoverTab[42708]"
//line /usr/local/go/src/net/http/server.go:2018
		_go_fuzz_dep_.CoverTab[42709]++

								if d := c.server.idleTimeout(); d != 0 {
//line /usr/local/go/src/net/http/server.go:2020
			_go_fuzz_dep_.CoverTab[42742]++
									c.rwc.SetReadDeadline(time.Now().Add(d))
//line /usr/local/go/src/net/http/server.go:2021
			// _ = "end of CoverTab[42742]"
		} else {
//line /usr/local/go/src/net/http/server.go:2022
			_go_fuzz_dep_.CoverTab[42743]++
									c.rwc.SetReadDeadline(time.Time{})
//line /usr/local/go/src/net/http/server.go:2023
			// _ = "end of CoverTab[42743]"
		}
//line /usr/local/go/src/net/http/server.go:2024
		// _ = "end of CoverTab[42709]"
//line /usr/local/go/src/net/http/server.go:2024
		_go_fuzz_dep_.CoverTab[42710]++

//line /usr/local/go/src/net/http/server.go:2030
		if _, err := c.bufr.Peek(4); err != nil {
//line /usr/local/go/src/net/http/server.go:2030
			_go_fuzz_dep_.CoverTab[42744]++
									return
//line /usr/local/go/src/net/http/server.go:2031
			// _ = "end of CoverTab[42744]"
		} else {
//line /usr/local/go/src/net/http/server.go:2032
			_go_fuzz_dep_.CoverTab[42745]++
//line /usr/local/go/src/net/http/server.go:2032
			// _ = "end of CoverTab[42745]"
//line /usr/local/go/src/net/http/server.go:2032
		}
//line /usr/local/go/src/net/http/server.go:2032
		// _ = "end of CoverTab[42710]"
//line /usr/local/go/src/net/http/server.go:2032
		_go_fuzz_dep_.CoverTab[42711]++

								c.rwc.SetReadDeadline(time.Time{})
//line /usr/local/go/src/net/http/server.go:2034
		// _ = "end of CoverTab[42711]"
	}
//line /usr/local/go/src/net/http/server.go:2035
	// _ = "end of CoverTab[42667]"
}

func (w *response) sendExpectationFailed() {
//line /usr/local/go/src/net/http/server.go:2038
	_go_fuzz_dep_.CoverTab[42746]++

//line /usr/local/go/src/net/http/server.go:2051
	w.Header().Set("Connection", "close")
							w.WriteHeader(StatusExpectationFailed)
							w.finishRequest()
//line /usr/local/go/src/net/http/server.go:2053
	// _ = "end of CoverTab[42746]"
}

// Hijack implements the Hijacker.Hijack method. Our response is both a ResponseWriter
//line /usr/local/go/src/net/http/server.go:2056
// and a Hijacker.
//line /usr/local/go/src/net/http/server.go:2058
func (w *response) Hijack() (rwc net.Conn, buf *bufio.ReadWriter, err error) {
//line /usr/local/go/src/net/http/server.go:2058
	_go_fuzz_dep_.CoverTab[42747]++
							if w.handlerDone.Load() {
//line /usr/local/go/src/net/http/server.go:2059
		_go_fuzz_dep_.CoverTab[42751]++
								panic("net/http: Hijack called after ServeHTTP finished")
//line /usr/local/go/src/net/http/server.go:2060
		// _ = "end of CoverTab[42751]"
	} else {
//line /usr/local/go/src/net/http/server.go:2061
		_go_fuzz_dep_.CoverTab[42752]++
//line /usr/local/go/src/net/http/server.go:2061
		// _ = "end of CoverTab[42752]"
//line /usr/local/go/src/net/http/server.go:2061
	}
//line /usr/local/go/src/net/http/server.go:2061
	// _ = "end of CoverTab[42747]"
//line /usr/local/go/src/net/http/server.go:2061
	_go_fuzz_dep_.CoverTab[42748]++
							if w.wroteHeader {
//line /usr/local/go/src/net/http/server.go:2062
		_go_fuzz_dep_.CoverTab[42753]++
								w.cw.flush()
//line /usr/local/go/src/net/http/server.go:2063
		// _ = "end of CoverTab[42753]"
	} else {
//line /usr/local/go/src/net/http/server.go:2064
		_go_fuzz_dep_.CoverTab[42754]++
//line /usr/local/go/src/net/http/server.go:2064
		// _ = "end of CoverTab[42754]"
//line /usr/local/go/src/net/http/server.go:2064
	}
//line /usr/local/go/src/net/http/server.go:2064
	// _ = "end of CoverTab[42748]"
//line /usr/local/go/src/net/http/server.go:2064
	_go_fuzz_dep_.CoverTab[42749]++

							c := w.conn
							c.mu.Lock()
							defer c.mu.Unlock()

//line /usr/local/go/src/net/http/server.go:2072
	rwc, buf, err = c.hijackLocked()
	if err == nil {
//line /usr/local/go/src/net/http/server.go:2073
		_go_fuzz_dep_.CoverTab[42755]++
								putBufioWriter(w.w)
								w.w = nil
//line /usr/local/go/src/net/http/server.go:2075
		// _ = "end of CoverTab[42755]"
	} else {
//line /usr/local/go/src/net/http/server.go:2076
		_go_fuzz_dep_.CoverTab[42756]++
//line /usr/local/go/src/net/http/server.go:2076
		// _ = "end of CoverTab[42756]"
//line /usr/local/go/src/net/http/server.go:2076
	}
//line /usr/local/go/src/net/http/server.go:2076
	// _ = "end of CoverTab[42749]"
//line /usr/local/go/src/net/http/server.go:2076
	_go_fuzz_dep_.CoverTab[42750]++
							return rwc, buf, err
//line /usr/local/go/src/net/http/server.go:2077
	// _ = "end of CoverTab[42750]"
}

func (w *response) CloseNotify() <-chan bool {
//line /usr/local/go/src/net/http/server.go:2080
	_go_fuzz_dep_.CoverTab[42757]++
							if w.handlerDone.Load() {
//line /usr/local/go/src/net/http/server.go:2081
		_go_fuzz_dep_.CoverTab[42759]++
								panic("net/http: CloseNotify called after ServeHTTP finished")
//line /usr/local/go/src/net/http/server.go:2082
		// _ = "end of CoverTab[42759]"
	} else {
//line /usr/local/go/src/net/http/server.go:2083
		_go_fuzz_dep_.CoverTab[42760]++
//line /usr/local/go/src/net/http/server.go:2083
		// _ = "end of CoverTab[42760]"
//line /usr/local/go/src/net/http/server.go:2083
	}
//line /usr/local/go/src/net/http/server.go:2083
	// _ = "end of CoverTab[42757]"
//line /usr/local/go/src/net/http/server.go:2083
	_go_fuzz_dep_.CoverTab[42758]++
							return w.closeNotifyCh
//line /usr/local/go/src/net/http/server.go:2084
	// _ = "end of CoverTab[42758]"
}

func registerOnHitEOF(rc io.ReadCloser, fn func()) {
//line /usr/local/go/src/net/http/server.go:2087
	_go_fuzz_dep_.CoverTab[42761]++
							switch v := rc.(type) {
	case *expectContinueReader:
//line /usr/local/go/src/net/http/server.go:2089
		_go_fuzz_dep_.CoverTab[42762]++
								registerOnHitEOF(v.readCloser, fn)
//line /usr/local/go/src/net/http/server.go:2090
		// _ = "end of CoverTab[42762]"
	case *body:
//line /usr/local/go/src/net/http/server.go:2091
		_go_fuzz_dep_.CoverTab[42763]++
								v.registerOnHitEOF(fn)
//line /usr/local/go/src/net/http/server.go:2092
		// _ = "end of CoverTab[42763]"
	default:
//line /usr/local/go/src/net/http/server.go:2093
		_go_fuzz_dep_.CoverTab[42764]++
								panic("unexpected type " + fmt.Sprintf("%T", rc))
//line /usr/local/go/src/net/http/server.go:2094
		// _ = "end of CoverTab[42764]"
	}
//line /usr/local/go/src/net/http/server.go:2095
	// _ = "end of CoverTab[42761]"
}

// requestBodyRemains reports whether future calls to Read
//line /usr/local/go/src/net/http/server.go:2098
// on rc might yield more data.
//line /usr/local/go/src/net/http/server.go:2100
func requestBodyRemains(rc io.ReadCloser) bool {
//line /usr/local/go/src/net/http/server.go:2100
	_go_fuzz_dep_.CoverTab[42765]++
							if rc == NoBody {
//line /usr/local/go/src/net/http/server.go:2101
		_go_fuzz_dep_.CoverTab[42767]++
								return false
//line /usr/local/go/src/net/http/server.go:2102
		// _ = "end of CoverTab[42767]"
	} else {
//line /usr/local/go/src/net/http/server.go:2103
		_go_fuzz_dep_.CoverTab[42768]++
//line /usr/local/go/src/net/http/server.go:2103
		// _ = "end of CoverTab[42768]"
//line /usr/local/go/src/net/http/server.go:2103
	}
//line /usr/local/go/src/net/http/server.go:2103
	// _ = "end of CoverTab[42765]"
//line /usr/local/go/src/net/http/server.go:2103
	_go_fuzz_dep_.CoverTab[42766]++
							switch v := rc.(type) {
	case *expectContinueReader:
//line /usr/local/go/src/net/http/server.go:2105
		_go_fuzz_dep_.CoverTab[42769]++
								return requestBodyRemains(v.readCloser)
//line /usr/local/go/src/net/http/server.go:2106
		// _ = "end of CoverTab[42769]"
	case *body:
//line /usr/local/go/src/net/http/server.go:2107
		_go_fuzz_dep_.CoverTab[42770]++
								return v.bodyRemains()
//line /usr/local/go/src/net/http/server.go:2108
		// _ = "end of CoverTab[42770]"
	default:
//line /usr/local/go/src/net/http/server.go:2109
		_go_fuzz_dep_.CoverTab[42771]++
								panic("unexpected type " + fmt.Sprintf("%T", rc))
//line /usr/local/go/src/net/http/server.go:2110
		// _ = "end of CoverTab[42771]"
	}
//line /usr/local/go/src/net/http/server.go:2111
	// _ = "end of CoverTab[42766]"
}

// The HandlerFunc type is an adapter to allow the use of
//line /usr/local/go/src/net/http/server.go:2114
// ordinary functions as HTTP handlers. If f is a function
//line /usr/local/go/src/net/http/server.go:2114
// with the appropriate signature, HandlerFunc(f) is a
//line /usr/local/go/src/net/http/server.go:2114
// Handler that calls f.
//line /usr/local/go/src/net/http/server.go:2118
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
//line /usr/local/go/src/net/http/server.go:2121
	_go_fuzz_dep_.CoverTab[42772]++
							f(w, r)
//line /usr/local/go/src/net/http/server.go:2122
	// _ = "end of CoverTab[42772]"
}

//line /usr/local/go/src/net/http/server.go:2127
// Error replies to the request with the specified error message and HTTP code.
//line /usr/local/go/src/net/http/server.go:2127
// It does not otherwise end the request; the caller should ensure no further
//line /usr/local/go/src/net/http/server.go:2127
// writes are done to w.
//line /usr/local/go/src/net/http/server.go:2127
// The error message should be plain text.
//line /usr/local/go/src/net/http/server.go:2131
func Error(w ResponseWriter, error string, code int) {
//line /usr/local/go/src/net/http/server.go:2131
	_go_fuzz_dep_.CoverTab[42773]++
							w.Header().Set("Content-Type", "text/plain; charset=utf-8")
							w.Header().Set("X-Content-Type-Options", "nosniff")
							w.WriteHeader(code)
							fmt.Fprintln(w, error)
//line /usr/local/go/src/net/http/server.go:2135
	// _ = "end of CoverTab[42773]"
}

// NotFound replies to the request with an HTTP 404 not found error.
func NotFound(w ResponseWriter, r *Request) {
//line /usr/local/go/src/net/http/server.go:2139
	_go_fuzz_dep_.CoverTab[42774]++
//line /usr/local/go/src/net/http/server.go:2139
	Error(w, "404 page not found", StatusNotFound)
//line /usr/local/go/src/net/http/server.go:2139
	// _ = "end of CoverTab[42774]"
//line /usr/local/go/src/net/http/server.go:2139
}

// NotFoundHandler returns a simple request handler
//line /usr/local/go/src/net/http/server.go:2141
// that replies to each request with a “404 page not found” reply.
//line /usr/local/go/src/net/http/server.go:2143
func NotFoundHandler() Handler {
//line /usr/local/go/src/net/http/server.go:2143
	_go_fuzz_dep_.CoverTab[42775]++
//line /usr/local/go/src/net/http/server.go:2143
	return HandlerFunc(NotFound)
//line /usr/local/go/src/net/http/server.go:2143
	// _ = "end of CoverTab[42775]"
//line /usr/local/go/src/net/http/server.go:2143
}

// StripPrefix returns a handler that serves HTTP requests by removing the
//line /usr/local/go/src/net/http/server.go:2145
// given prefix from the request URL's Path (and RawPath if set) and invoking
//line /usr/local/go/src/net/http/server.go:2145
// the handler h. StripPrefix handles a request for a path that doesn't begin
//line /usr/local/go/src/net/http/server.go:2145
// with prefix by replying with an HTTP 404 not found error. The prefix must
//line /usr/local/go/src/net/http/server.go:2145
// match exactly: if the prefix in the request contains escaped characters
//line /usr/local/go/src/net/http/server.go:2145
// the reply is also an HTTP 404 not found error.
//line /usr/local/go/src/net/http/server.go:2151
func StripPrefix(prefix string, h Handler) Handler {
//line /usr/local/go/src/net/http/server.go:2151
	_go_fuzz_dep_.CoverTab[42776]++
							if prefix == "" {
//line /usr/local/go/src/net/http/server.go:2152
		_go_fuzz_dep_.CoverTab[42778]++
								return h
//line /usr/local/go/src/net/http/server.go:2153
		// _ = "end of CoverTab[42778]"
	} else {
//line /usr/local/go/src/net/http/server.go:2154
		_go_fuzz_dep_.CoverTab[42779]++
//line /usr/local/go/src/net/http/server.go:2154
		// _ = "end of CoverTab[42779]"
//line /usr/local/go/src/net/http/server.go:2154
	}
//line /usr/local/go/src/net/http/server.go:2154
	// _ = "end of CoverTab[42776]"
//line /usr/local/go/src/net/http/server.go:2154
	_go_fuzz_dep_.CoverTab[42777]++
							return HandlerFunc(func(w ResponseWriter, r *Request) {
//line /usr/local/go/src/net/http/server.go:2155
		_go_fuzz_dep_.CoverTab[42780]++
								p := strings.TrimPrefix(r.URL.Path, prefix)
								rp := strings.TrimPrefix(r.URL.RawPath, prefix)
								if len(p) < len(r.URL.Path) && func() bool {
//line /usr/local/go/src/net/http/server.go:2158
			_go_fuzz_dep_.CoverTab[42781]++
//line /usr/local/go/src/net/http/server.go:2158
			return (r.URL.RawPath == "" || func() bool {
//line /usr/local/go/src/net/http/server.go:2158
				_go_fuzz_dep_.CoverTab[42782]++
//line /usr/local/go/src/net/http/server.go:2158
				return len(rp) < len(r.URL.RawPath)
//line /usr/local/go/src/net/http/server.go:2158
				// _ = "end of CoverTab[42782]"
//line /usr/local/go/src/net/http/server.go:2158
			}())
//line /usr/local/go/src/net/http/server.go:2158
			// _ = "end of CoverTab[42781]"
//line /usr/local/go/src/net/http/server.go:2158
		}() {
//line /usr/local/go/src/net/http/server.go:2158
			_go_fuzz_dep_.CoverTab[42783]++
									r2 := new(Request)
									*r2 = *r
									r2.URL = new(url.URL)
									*r2.URL = *r.URL
									r2.URL.Path = p
									r2.URL.RawPath = rp
									h.ServeHTTP(w, r2)
//line /usr/local/go/src/net/http/server.go:2165
			// _ = "end of CoverTab[42783]"
		} else {
//line /usr/local/go/src/net/http/server.go:2166
			_go_fuzz_dep_.CoverTab[42784]++
									NotFound(w, r)
//line /usr/local/go/src/net/http/server.go:2167
			// _ = "end of CoverTab[42784]"
		}
//line /usr/local/go/src/net/http/server.go:2168
		// _ = "end of CoverTab[42780]"
	})
//line /usr/local/go/src/net/http/server.go:2169
	// _ = "end of CoverTab[42777]"
}

// Redirect replies to the request with a redirect to url,
//line /usr/local/go/src/net/http/server.go:2172
// which may be a path relative to the request path.
//line /usr/local/go/src/net/http/server.go:2172
//
//line /usr/local/go/src/net/http/server.go:2172
// The provided code should be in the 3xx range and is usually
//line /usr/local/go/src/net/http/server.go:2172
// StatusMovedPermanently, StatusFound or StatusSeeOther.
//line /usr/local/go/src/net/http/server.go:2172
//
//line /usr/local/go/src/net/http/server.go:2172
// If the Content-Type header has not been set, Redirect sets it
//line /usr/local/go/src/net/http/server.go:2172
// to "text/html; charset=utf-8" and writes a small HTML body.
//line /usr/local/go/src/net/http/server.go:2172
// Setting the Content-Type header to any value, including nil,
//line /usr/local/go/src/net/http/server.go:2172
// disables that behavior.
//line /usr/local/go/src/net/http/server.go:2182
func Redirect(w ResponseWriter, r *Request, url string, code int) {
//line /usr/local/go/src/net/http/server.go:2182
	_go_fuzz_dep_.CoverTab[42785]++
							if u, err := urlpkg.Parse(url); err == nil {
//line /usr/local/go/src/net/http/server.go:2183
		_go_fuzz_dep_.CoverTab[42788]++

//line /usr/local/go/src/net/http/server.go:2189
		if u.Scheme == "" && func() bool {
//line /usr/local/go/src/net/http/server.go:2189
			_go_fuzz_dep_.CoverTab[42789]++
//line /usr/local/go/src/net/http/server.go:2189
			return u.Host == ""
//line /usr/local/go/src/net/http/server.go:2189
			// _ = "end of CoverTab[42789]"
//line /usr/local/go/src/net/http/server.go:2189
		}() {
//line /usr/local/go/src/net/http/server.go:2189
			_go_fuzz_dep_.CoverTab[42790]++
									oldpath := r.URL.Path
									if oldpath == "" {
//line /usr/local/go/src/net/http/server.go:2191
				_go_fuzz_dep_.CoverTab[42795]++
										oldpath = "/"
//line /usr/local/go/src/net/http/server.go:2192
				// _ = "end of CoverTab[42795]"
			} else {
//line /usr/local/go/src/net/http/server.go:2193
				_go_fuzz_dep_.CoverTab[42796]++
//line /usr/local/go/src/net/http/server.go:2193
				// _ = "end of CoverTab[42796]"
//line /usr/local/go/src/net/http/server.go:2193
			}
//line /usr/local/go/src/net/http/server.go:2193
			// _ = "end of CoverTab[42790]"
//line /usr/local/go/src/net/http/server.go:2193
			_go_fuzz_dep_.CoverTab[42791]++

//line /usr/local/go/src/net/http/server.go:2196
			if url == "" || func() bool {
//line /usr/local/go/src/net/http/server.go:2196
				_go_fuzz_dep_.CoverTab[42797]++
//line /usr/local/go/src/net/http/server.go:2196
				return url[0] != '/'
//line /usr/local/go/src/net/http/server.go:2196
				// _ = "end of CoverTab[42797]"
//line /usr/local/go/src/net/http/server.go:2196
			}() {
//line /usr/local/go/src/net/http/server.go:2196
				_go_fuzz_dep_.CoverTab[42798]++

										olddir, _ := path.Split(oldpath)
										url = olddir + url
//line /usr/local/go/src/net/http/server.go:2199
				// _ = "end of CoverTab[42798]"
			} else {
//line /usr/local/go/src/net/http/server.go:2200
				_go_fuzz_dep_.CoverTab[42799]++
//line /usr/local/go/src/net/http/server.go:2200
				// _ = "end of CoverTab[42799]"
//line /usr/local/go/src/net/http/server.go:2200
			}
//line /usr/local/go/src/net/http/server.go:2200
			// _ = "end of CoverTab[42791]"
//line /usr/local/go/src/net/http/server.go:2200
			_go_fuzz_dep_.CoverTab[42792]++

									var query string
									if i := strings.Index(url, "?"); i != -1 {
//line /usr/local/go/src/net/http/server.go:2203
				_go_fuzz_dep_.CoverTab[42800]++
										url, query = url[:i], url[i:]
//line /usr/local/go/src/net/http/server.go:2204
				// _ = "end of CoverTab[42800]"
			} else {
//line /usr/local/go/src/net/http/server.go:2205
				_go_fuzz_dep_.CoverTab[42801]++
//line /usr/local/go/src/net/http/server.go:2205
				// _ = "end of CoverTab[42801]"
//line /usr/local/go/src/net/http/server.go:2205
			}
//line /usr/local/go/src/net/http/server.go:2205
			// _ = "end of CoverTab[42792]"
//line /usr/local/go/src/net/http/server.go:2205
			_go_fuzz_dep_.CoverTab[42793]++

//line /usr/local/go/src/net/http/server.go:2208
			trailing := strings.HasSuffix(url, "/")
			url = path.Clean(url)
			if trailing && func() bool {
//line /usr/local/go/src/net/http/server.go:2210
				_go_fuzz_dep_.CoverTab[42802]++
//line /usr/local/go/src/net/http/server.go:2210
				return !strings.HasSuffix(url, "/")
//line /usr/local/go/src/net/http/server.go:2210
				// _ = "end of CoverTab[42802]"
//line /usr/local/go/src/net/http/server.go:2210
			}() {
//line /usr/local/go/src/net/http/server.go:2210
				_go_fuzz_dep_.CoverTab[42803]++
										url += "/"
//line /usr/local/go/src/net/http/server.go:2211
				// _ = "end of CoverTab[42803]"
			} else {
//line /usr/local/go/src/net/http/server.go:2212
				_go_fuzz_dep_.CoverTab[42804]++
//line /usr/local/go/src/net/http/server.go:2212
				// _ = "end of CoverTab[42804]"
//line /usr/local/go/src/net/http/server.go:2212
			}
//line /usr/local/go/src/net/http/server.go:2212
			// _ = "end of CoverTab[42793]"
//line /usr/local/go/src/net/http/server.go:2212
			_go_fuzz_dep_.CoverTab[42794]++
									url += query
//line /usr/local/go/src/net/http/server.go:2213
			// _ = "end of CoverTab[42794]"
		} else {
//line /usr/local/go/src/net/http/server.go:2214
			_go_fuzz_dep_.CoverTab[42805]++
//line /usr/local/go/src/net/http/server.go:2214
			// _ = "end of CoverTab[42805]"
//line /usr/local/go/src/net/http/server.go:2214
		}
//line /usr/local/go/src/net/http/server.go:2214
		// _ = "end of CoverTab[42788]"
	} else {
//line /usr/local/go/src/net/http/server.go:2215
		_go_fuzz_dep_.CoverTab[42806]++
//line /usr/local/go/src/net/http/server.go:2215
		// _ = "end of CoverTab[42806]"
//line /usr/local/go/src/net/http/server.go:2215
	}
//line /usr/local/go/src/net/http/server.go:2215
	// _ = "end of CoverTab[42785]"
//line /usr/local/go/src/net/http/server.go:2215
	_go_fuzz_dep_.CoverTab[42786]++

							h := w.Header()

//line /usr/local/go/src/net/http/server.go:2222
	_, hadCT := h["Content-Type"]

	h.Set("Location", hexEscapeNonASCII(url))
	if !hadCT && func() bool {
//line /usr/local/go/src/net/http/server.go:2225
		_go_fuzz_dep_.CoverTab[42807]++
//line /usr/local/go/src/net/http/server.go:2225
		return (r.Method == "GET" || func() bool {
//line /usr/local/go/src/net/http/server.go:2225
			_go_fuzz_dep_.CoverTab[42808]++
//line /usr/local/go/src/net/http/server.go:2225
			return r.Method == "HEAD"
//line /usr/local/go/src/net/http/server.go:2225
			// _ = "end of CoverTab[42808]"
//line /usr/local/go/src/net/http/server.go:2225
		}())
//line /usr/local/go/src/net/http/server.go:2225
		// _ = "end of CoverTab[42807]"
//line /usr/local/go/src/net/http/server.go:2225
	}() {
//line /usr/local/go/src/net/http/server.go:2225
		_go_fuzz_dep_.CoverTab[42809]++
								h.Set("Content-Type", "text/html; charset=utf-8")
//line /usr/local/go/src/net/http/server.go:2226
		// _ = "end of CoverTab[42809]"
	} else {
//line /usr/local/go/src/net/http/server.go:2227
		_go_fuzz_dep_.CoverTab[42810]++
//line /usr/local/go/src/net/http/server.go:2227
		// _ = "end of CoverTab[42810]"
//line /usr/local/go/src/net/http/server.go:2227
	}
//line /usr/local/go/src/net/http/server.go:2227
	// _ = "end of CoverTab[42786]"
//line /usr/local/go/src/net/http/server.go:2227
	_go_fuzz_dep_.CoverTab[42787]++
							w.WriteHeader(code)

//line /usr/local/go/src/net/http/server.go:2231
	if !hadCT && func() bool {
//line /usr/local/go/src/net/http/server.go:2231
		_go_fuzz_dep_.CoverTab[42811]++
//line /usr/local/go/src/net/http/server.go:2231
		return r.Method == "GET"
//line /usr/local/go/src/net/http/server.go:2231
		// _ = "end of CoverTab[42811]"
//line /usr/local/go/src/net/http/server.go:2231
	}() {
//line /usr/local/go/src/net/http/server.go:2231
		_go_fuzz_dep_.CoverTab[42812]++
								body := "<a href=\"" + htmlEscape(url) + "\">" + StatusText(code) + "</a>.\n"
								fmt.Fprintln(w, body)
//line /usr/local/go/src/net/http/server.go:2233
		// _ = "end of CoverTab[42812]"
	} else {
//line /usr/local/go/src/net/http/server.go:2234
		_go_fuzz_dep_.CoverTab[42813]++
//line /usr/local/go/src/net/http/server.go:2234
		// _ = "end of CoverTab[42813]"
//line /usr/local/go/src/net/http/server.go:2234
	}
//line /usr/local/go/src/net/http/server.go:2234
	// _ = "end of CoverTab[42787]"
}

var htmlReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",

	`"`, "&#34;",

	"'", "&#39;",
)

func htmlEscape(s string) string {
//line /usr/local/go/src/net/http/server.go:2247
	_go_fuzz_dep_.CoverTab[42814]++
							return htmlReplacer.Replace(s)
//line /usr/local/go/src/net/http/server.go:2248
	// _ = "end of CoverTab[42814]"
}

// Redirect to a fixed URL
type redirectHandler struct {
	url	string
	code	int
}

func (rh *redirectHandler) ServeHTTP(w ResponseWriter, r *Request) {
//line /usr/local/go/src/net/http/server.go:2257
	_go_fuzz_dep_.CoverTab[42815]++
							Redirect(w, r, rh.url, rh.code)
//line /usr/local/go/src/net/http/server.go:2258
	// _ = "end of CoverTab[42815]"
}

// RedirectHandler returns a request handler that redirects
//line /usr/local/go/src/net/http/server.go:2261
// each request it receives to the given url using the given
//line /usr/local/go/src/net/http/server.go:2261
// status code.
//line /usr/local/go/src/net/http/server.go:2261
//
//line /usr/local/go/src/net/http/server.go:2261
// The provided code should be in the 3xx range and is usually
//line /usr/local/go/src/net/http/server.go:2261
// StatusMovedPermanently, StatusFound or StatusSeeOther.
//line /usr/local/go/src/net/http/server.go:2267
func RedirectHandler(url string, code int) Handler {
//line /usr/local/go/src/net/http/server.go:2267
	_go_fuzz_dep_.CoverTab[42816]++
							return &redirectHandler{url, code}
//line /usr/local/go/src/net/http/server.go:2268
	// _ = "end of CoverTab[42816]"
}

// ServeMux is an HTTP request multiplexer.
//line /usr/local/go/src/net/http/server.go:2271
// It matches the URL of each incoming request against a list of registered
//line /usr/local/go/src/net/http/server.go:2271
// patterns and calls the handler for the pattern that
//line /usr/local/go/src/net/http/server.go:2271
// most closely matches the URL.
//line /usr/local/go/src/net/http/server.go:2271
//
//line /usr/local/go/src/net/http/server.go:2271
// Patterns name fixed, rooted paths, like "/favicon.ico",
//line /usr/local/go/src/net/http/server.go:2271
// or rooted subtrees, like "/images/" (note the trailing slash).
//line /usr/local/go/src/net/http/server.go:2271
// Longer patterns take precedence over shorter ones, so that
//line /usr/local/go/src/net/http/server.go:2271
// if there are handlers registered for both "/images/"
//line /usr/local/go/src/net/http/server.go:2271
// and "/images/thumbnails/", the latter handler will be
//line /usr/local/go/src/net/http/server.go:2271
// called for paths beginning "/images/thumbnails/" and the
//line /usr/local/go/src/net/http/server.go:2271
// former will receive requests for any other paths in the
//line /usr/local/go/src/net/http/server.go:2271
// "/images/" subtree.
//line /usr/local/go/src/net/http/server.go:2271
//
//line /usr/local/go/src/net/http/server.go:2271
// Note that since a pattern ending in a slash names a rooted subtree,
//line /usr/local/go/src/net/http/server.go:2271
// the pattern "/" matches all paths not matched by other registered
//line /usr/local/go/src/net/http/server.go:2271
// patterns, not just the URL with Path == "/".
//line /usr/local/go/src/net/http/server.go:2271
//
//line /usr/local/go/src/net/http/server.go:2271
// If a subtree has been registered and a request is received naming the
//line /usr/local/go/src/net/http/server.go:2271
// subtree root without its trailing slash, ServeMux redirects that
//line /usr/local/go/src/net/http/server.go:2271
// request to the subtree root (adding the trailing slash). This behavior can
//line /usr/local/go/src/net/http/server.go:2271
// be overridden with a separate registration for the path without
//line /usr/local/go/src/net/http/server.go:2271
// the trailing slash. For example, registering "/images/" causes ServeMux
//line /usr/local/go/src/net/http/server.go:2271
// to redirect a request for "/images" to "/images/", unless "/images" has
//line /usr/local/go/src/net/http/server.go:2271
// been registered separately.
//line /usr/local/go/src/net/http/server.go:2271
//
//line /usr/local/go/src/net/http/server.go:2271
// Patterns may optionally begin with a host name, restricting matches to
//line /usr/local/go/src/net/http/server.go:2271
// URLs on that host only. Host-specific patterns take precedence over
//line /usr/local/go/src/net/http/server.go:2271
// general patterns, so that a handler might register for the two patterns
//line /usr/local/go/src/net/http/server.go:2271
// "/codesearch" and "codesearch.google.com/" without also taking over
//line /usr/local/go/src/net/http/server.go:2271
// requests for "http://www.google.com/".
//line /usr/local/go/src/net/http/server.go:2271
//
//line /usr/local/go/src/net/http/server.go:2271
// ServeMux also takes care of sanitizing the URL request path and the Host
//line /usr/local/go/src/net/http/server.go:2271
// header, stripping the port number and redirecting any request containing . or
//line /usr/local/go/src/net/http/server.go:2271
// .. elements or repeated slashes to an equivalent, cleaner URL.
//line /usr/local/go/src/net/http/server.go:2306
type ServeMux struct {
	mu	sync.RWMutex
	m	map[string]muxEntry
	es	[]muxEntry	// slice of entries sorted from longest to shortest.
	hosts	bool		// whether any patterns contain hostnames
}

type muxEntry struct {
	h	Handler
	pattern	string
}

// NewServeMux allocates and returns a new ServeMux.
func NewServeMux() *ServeMux {
//line /usr/local/go/src/net/http/server.go:2319
	_go_fuzz_dep_.CoverTab[42817]++
//line /usr/local/go/src/net/http/server.go:2319
	return new(ServeMux)
//line /usr/local/go/src/net/http/server.go:2319
	// _ = "end of CoverTab[42817]"
//line /usr/local/go/src/net/http/server.go:2319
}

// DefaultServeMux is the default ServeMux used by Serve.
var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux

// cleanPath returns the canonical path for p, eliminating . and .. elements.
func cleanPath(p string) string {
//line /usr/local/go/src/net/http/server.go:2327
	_go_fuzz_dep_.CoverTab[42818]++
							if p == "" {
//line /usr/local/go/src/net/http/server.go:2328
		_go_fuzz_dep_.CoverTab[42822]++
								return "/"
//line /usr/local/go/src/net/http/server.go:2329
		// _ = "end of CoverTab[42822]"
	} else {
//line /usr/local/go/src/net/http/server.go:2330
		_go_fuzz_dep_.CoverTab[42823]++
//line /usr/local/go/src/net/http/server.go:2330
		// _ = "end of CoverTab[42823]"
//line /usr/local/go/src/net/http/server.go:2330
	}
//line /usr/local/go/src/net/http/server.go:2330
	// _ = "end of CoverTab[42818]"
//line /usr/local/go/src/net/http/server.go:2330
	_go_fuzz_dep_.CoverTab[42819]++
							if p[0] != '/' {
//line /usr/local/go/src/net/http/server.go:2331
		_go_fuzz_dep_.CoverTab[42824]++
								p = "/" + p
//line /usr/local/go/src/net/http/server.go:2332
		// _ = "end of CoverTab[42824]"
	} else {
//line /usr/local/go/src/net/http/server.go:2333
		_go_fuzz_dep_.CoverTab[42825]++
//line /usr/local/go/src/net/http/server.go:2333
		// _ = "end of CoverTab[42825]"
//line /usr/local/go/src/net/http/server.go:2333
	}
//line /usr/local/go/src/net/http/server.go:2333
	// _ = "end of CoverTab[42819]"
//line /usr/local/go/src/net/http/server.go:2333
	_go_fuzz_dep_.CoverTab[42820]++
							np := path.Clean(p)

//line /usr/local/go/src/net/http/server.go:2337
	if p[len(p)-1] == '/' && func() bool {
//line /usr/local/go/src/net/http/server.go:2337
		_go_fuzz_dep_.CoverTab[42826]++
//line /usr/local/go/src/net/http/server.go:2337
		return np != "/"
//line /usr/local/go/src/net/http/server.go:2337
		// _ = "end of CoverTab[42826]"
//line /usr/local/go/src/net/http/server.go:2337
	}() {
//line /usr/local/go/src/net/http/server.go:2337
		_go_fuzz_dep_.CoverTab[42827]++

								if len(p) == len(np)+1 && func() bool {
//line /usr/local/go/src/net/http/server.go:2339
			_go_fuzz_dep_.CoverTab[42828]++
//line /usr/local/go/src/net/http/server.go:2339
			return strings.HasPrefix(p, np)
//line /usr/local/go/src/net/http/server.go:2339
			// _ = "end of CoverTab[42828]"
//line /usr/local/go/src/net/http/server.go:2339
		}() {
//line /usr/local/go/src/net/http/server.go:2339
			_go_fuzz_dep_.CoverTab[42829]++
									np = p
//line /usr/local/go/src/net/http/server.go:2340
			// _ = "end of CoverTab[42829]"
		} else {
//line /usr/local/go/src/net/http/server.go:2341
			_go_fuzz_dep_.CoverTab[42830]++
									np += "/"
//line /usr/local/go/src/net/http/server.go:2342
			// _ = "end of CoverTab[42830]"
		}
//line /usr/local/go/src/net/http/server.go:2343
		// _ = "end of CoverTab[42827]"
	} else {
//line /usr/local/go/src/net/http/server.go:2344
		_go_fuzz_dep_.CoverTab[42831]++
//line /usr/local/go/src/net/http/server.go:2344
		// _ = "end of CoverTab[42831]"
//line /usr/local/go/src/net/http/server.go:2344
	}
//line /usr/local/go/src/net/http/server.go:2344
	// _ = "end of CoverTab[42820]"
//line /usr/local/go/src/net/http/server.go:2344
	_go_fuzz_dep_.CoverTab[42821]++
							return np
//line /usr/local/go/src/net/http/server.go:2345
	// _ = "end of CoverTab[42821]"
}

// stripHostPort returns h without any trailing ":<port>".
func stripHostPort(h string) string {
//line /usr/local/go/src/net/http/server.go:2349
	_go_fuzz_dep_.CoverTab[42832]++

							if !strings.Contains(h, ":") {
//line /usr/local/go/src/net/http/server.go:2351
		_go_fuzz_dep_.CoverTab[42835]++
								return h
//line /usr/local/go/src/net/http/server.go:2352
		// _ = "end of CoverTab[42835]"
	} else {
//line /usr/local/go/src/net/http/server.go:2353
		_go_fuzz_dep_.CoverTab[42836]++
//line /usr/local/go/src/net/http/server.go:2353
		// _ = "end of CoverTab[42836]"
//line /usr/local/go/src/net/http/server.go:2353
	}
//line /usr/local/go/src/net/http/server.go:2353
	// _ = "end of CoverTab[42832]"
//line /usr/local/go/src/net/http/server.go:2353
	_go_fuzz_dep_.CoverTab[42833]++
							host, _, err := net.SplitHostPort(h)
							if err != nil {
//line /usr/local/go/src/net/http/server.go:2355
		_go_fuzz_dep_.CoverTab[42837]++
								return h
//line /usr/local/go/src/net/http/server.go:2356
		// _ = "end of CoverTab[42837]"
	} else {
//line /usr/local/go/src/net/http/server.go:2357
		_go_fuzz_dep_.CoverTab[42838]++
//line /usr/local/go/src/net/http/server.go:2357
		// _ = "end of CoverTab[42838]"
//line /usr/local/go/src/net/http/server.go:2357
	}
//line /usr/local/go/src/net/http/server.go:2357
	// _ = "end of CoverTab[42833]"
//line /usr/local/go/src/net/http/server.go:2357
	_go_fuzz_dep_.CoverTab[42834]++
							return host
//line /usr/local/go/src/net/http/server.go:2358
	// _ = "end of CoverTab[42834]"
}

// Find a handler on a handler map given a path string.
//line /usr/local/go/src/net/http/server.go:2361
// Most-specific (longest) pattern wins.
//line /usr/local/go/src/net/http/server.go:2363
func (mux *ServeMux) match(path string) (h Handler, pattern string) {
//line /usr/local/go/src/net/http/server.go:2363
	_go_fuzz_dep_.CoverTab[42839]++

							v, ok := mux.m[path]
							if ok {
//line /usr/local/go/src/net/http/server.go:2366
		_go_fuzz_dep_.CoverTab[42842]++
								return v.h, v.pattern
//line /usr/local/go/src/net/http/server.go:2367
		// _ = "end of CoverTab[42842]"
	} else {
//line /usr/local/go/src/net/http/server.go:2368
		_go_fuzz_dep_.CoverTab[42843]++
//line /usr/local/go/src/net/http/server.go:2368
		// _ = "end of CoverTab[42843]"
//line /usr/local/go/src/net/http/server.go:2368
	}
//line /usr/local/go/src/net/http/server.go:2368
	// _ = "end of CoverTab[42839]"
//line /usr/local/go/src/net/http/server.go:2368
	_go_fuzz_dep_.CoverTab[42840]++

//line /usr/local/go/src/net/http/server.go:2372
	for _, e := range mux.es {
//line /usr/local/go/src/net/http/server.go:2372
		_go_fuzz_dep_.CoverTab[42844]++
								if strings.HasPrefix(path, e.pattern) {
//line /usr/local/go/src/net/http/server.go:2373
			_go_fuzz_dep_.CoverTab[42845]++
									return e.h, e.pattern
//line /usr/local/go/src/net/http/server.go:2374
			// _ = "end of CoverTab[42845]"
		} else {
//line /usr/local/go/src/net/http/server.go:2375
			_go_fuzz_dep_.CoverTab[42846]++
//line /usr/local/go/src/net/http/server.go:2375
			// _ = "end of CoverTab[42846]"
//line /usr/local/go/src/net/http/server.go:2375
		}
//line /usr/local/go/src/net/http/server.go:2375
		// _ = "end of CoverTab[42844]"
	}
//line /usr/local/go/src/net/http/server.go:2376
	// _ = "end of CoverTab[42840]"
//line /usr/local/go/src/net/http/server.go:2376
	_go_fuzz_dep_.CoverTab[42841]++
							return nil, ""
//line /usr/local/go/src/net/http/server.go:2377
	// _ = "end of CoverTab[42841]"
}

// redirectToPathSlash determines if the given path needs appending "/" to it.
//line /usr/local/go/src/net/http/server.go:2380
// This occurs when a handler for path + "/" was already registered, but
//line /usr/local/go/src/net/http/server.go:2380
// not for path itself. If the path needs appending to, it creates a new
//line /usr/local/go/src/net/http/server.go:2380
// URL, setting the path to u.Path + "/" and returning true to indicate so.
//line /usr/local/go/src/net/http/server.go:2384
func (mux *ServeMux) redirectToPathSlash(host, path string, u *url.URL) (*url.URL, bool) {
//line /usr/local/go/src/net/http/server.go:2384
	_go_fuzz_dep_.CoverTab[42847]++
							mux.mu.RLock()
							shouldRedirect := mux.shouldRedirectRLocked(host, path)
							mux.mu.RUnlock()
							if !shouldRedirect {
//line /usr/local/go/src/net/http/server.go:2388
		_go_fuzz_dep_.CoverTab[42849]++
								return u, false
//line /usr/local/go/src/net/http/server.go:2389
		// _ = "end of CoverTab[42849]"
	} else {
//line /usr/local/go/src/net/http/server.go:2390
		_go_fuzz_dep_.CoverTab[42850]++
//line /usr/local/go/src/net/http/server.go:2390
		// _ = "end of CoverTab[42850]"
//line /usr/local/go/src/net/http/server.go:2390
	}
//line /usr/local/go/src/net/http/server.go:2390
	// _ = "end of CoverTab[42847]"
//line /usr/local/go/src/net/http/server.go:2390
	_go_fuzz_dep_.CoverTab[42848]++
							path = path + "/"
							u = &url.URL{Path: path, RawQuery: u.RawQuery}
							return u, true
//line /usr/local/go/src/net/http/server.go:2393
	// _ = "end of CoverTab[42848]"
}

// shouldRedirectRLocked reports whether the given path and host should be redirected to
//line /usr/local/go/src/net/http/server.go:2396
// path+"/". This should happen if a handler is registered for path+"/" but
//line /usr/local/go/src/net/http/server.go:2396
// not path -- see comments at ServeMux.
//line /usr/local/go/src/net/http/server.go:2399
func (mux *ServeMux) shouldRedirectRLocked(host, path string) bool {
//line /usr/local/go/src/net/http/server.go:2399
	_go_fuzz_dep_.CoverTab[42851]++
							p := []string{path, host + path}

							for _, c := range p {
//line /usr/local/go/src/net/http/server.go:2402
		_go_fuzz_dep_.CoverTab[42855]++
								if _, exist := mux.m[c]; exist {
//line /usr/local/go/src/net/http/server.go:2403
			_go_fuzz_dep_.CoverTab[42856]++
									return false
//line /usr/local/go/src/net/http/server.go:2404
			// _ = "end of CoverTab[42856]"
		} else {
//line /usr/local/go/src/net/http/server.go:2405
			_go_fuzz_dep_.CoverTab[42857]++
//line /usr/local/go/src/net/http/server.go:2405
			// _ = "end of CoverTab[42857]"
//line /usr/local/go/src/net/http/server.go:2405
		}
//line /usr/local/go/src/net/http/server.go:2405
		// _ = "end of CoverTab[42855]"
	}
//line /usr/local/go/src/net/http/server.go:2406
	// _ = "end of CoverTab[42851]"
//line /usr/local/go/src/net/http/server.go:2406
	_go_fuzz_dep_.CoverTab[42852]++

							n := len(path)
							if n == 0 {
//line /usr/local/go/src/net/http/server.go:2409
		_go_fuzz_dep_.CoverTab[42858]++
								return false
//line /usr/local/go/src/net/http/server.go:2410
		// _ = "end of CoverTab[42858]"
	} else {
//line /usr/local/go/src/net/http/server.go:2411
		_go_fuzz_dep_.CoverTab[42859]++
//line /usr/local/go/src/net/http/server.go:2411
		// _ = "end of CoverTab[42859]"
//line /usr/local/go/src/net/http/server.go:2411
	}
//line /usr/local/go/src/net/http/server.go:2411
	// _ = "end of CoverTab[42852]"
//line /usr/local/go/src/net/http/server.go:2411
	_go_fuzz_dep_.CoverTab[42853]++
							for _, c := range p {
//line /usr/local/go/src/net/http/server.go:2412
		_go_fuzz_dep_.CoverTab[42860]++
								if _, exist := mux.m[c+"/"]; exist {
//line /usr/local/go/src/net/http/server.go:2413
			_go_fuzz_dep_.CoverTab[42861]++
									return path[n-1] != '/'
//line /usr/local/go/src/net/http/server.go:2414
			// _ = "end of CoverTab[42861]"
		} else {
//line /usr/local/go/src/net/http/server.go:2415
			_go_fuzz_dep_.CoverTab[42862]++
//line /usr/local/go/src/net/http/server.go:2415
			// _ = "end of CoverTab[42862]"
//line /usr/local/go/src/net/http/server.go:2415
		}
//line /usr/local/go/src/net/http/server.go:2415
		// _ = "end of CoverTab[42860]"
	}
//line /usr/local/go/src/net/http/server.go:2416
	// _ = "end of CoverTab[42853]"
//line /usr/local/go/src/net/http/server.go:2416
	_go_fuzz_dep_.CoverTab[42854]++

							return false
//line /usr/local/go/src/net/http/server.go:2418
	// _ = "end of CoverTab[42854]"
}

// Handler returns the handler to use for the given request,
//line /usr/local/go/src/net/http/server.go:2421
// consulting r.Method, r.Host, and r.URL.Path. It always returns
//line /usr/local/go/src/net/http/server.go:2421
// a non-nil handler. If the path is not in its canonical form, the
//line /usr/local/go/src/net/http/server.go:2421
// handler will be an internally-generated handler that redirects
//line /usr/local/go/src/net/http/server.go:2421
// to the canonical path. If the host contains a port, it is ignored
//line /usr/local/go/src/net/http/server.go:2421
// when matching handlers.
//line /usr/local/go/src/net/http/server.go:2421
//
//line /usr/local/go/src/net/http/server.go:2421
// The path and host are used unchanged for CONNECT requests.
//line /usr/local/go/src/net/http/server.go:2421
//
//line /usr/local/go/src/net/http/server.go:2421
// Handler also returns the registered pattern that matches the
//line /usr/local/go/src/net/http/server.go:2421
// request or, in the case of internally-generated redirects,
//line /usr/local/go/src/net/http/server.go:2421
// the pattern that will match after following the redirect.
//line /usr/local/go/src/net/http/server.go:2421
//
//line /usr/local/go/src/net/http/server.go:2421
// If there is no registered handler that applies to the request,
//line /usr/local/go/src/net/http/server.go:2421
// Handler returns a “page not found” handler and an empty pattern.
//line /usr/local/go/src/net/http/server.go:2436
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) {
//line /usr/local/go/src/net/http/server.go:2436
	_go_fuzz_dep_.CoverTab[42863]++

//line /usr/local/go/src/net/http/server.go:2439
	if r.Method == "CONNECT" {
//line /usr/local/go/src/net/http/server.go:2439
		_go_fuzz_dep_.CoverTab[42867]++

//line /usr/local/go/src/net/http/server.go:2443
		if u, ok := mux.redirectToPathSlash(r.URL.Host, r.URL.Path, r.URL); ok {
//line /usr/local/go/src/net/http/server.go:2443
			_go_fuzz_dep_.CoverTab[42869]++
									return RedirectHandler(u.String(), StatusMovedPermanently), u.Path
//line /usr/local/go/src/net/http/server.go:2444
			// _ = "end of CoverTab[42869]"
		} else {
//line /usr/local/go/src/net/http/server.go:2445
			_go_fuzz_dep_.CoverTab[42870]++
//line /usr/local/go/src/net/http/server.go:2445
			// _ = "end of CoverTab[42870]"
//line /usr/local/go/src/net/http/server.go:2445
		}
//line /usr/local/go/src/net/http/server.go:2445
		// _ = "end of CoverTab[42867]"
//line /usr/local/go/src/net/http/server.go:2445
		_go_fuzz_dep_.CoverTab[42868]++

								return mux.handler(r.Host, r.URL.Path)
//line /usr/local/go/src/net/http/server.go:2447
		// _ = "end of CoverTab[42868]"
	} else {
//line /usr/local/go/src/net/http/server.go:2448
		_go_fuzz_dep_.CoverTab[42871]++
//line /usr/local/go/src/net/http/server.go:2448
		// _ = "end of CoverTab[42871]"
//line /usr/local/go/src/net/http/server.go:2448
	}
//line /usr/local/go/src/net/http/server.go:2448
	// _ = "end of CoverTab[42863]"
//line /usr/local/go/src/net/http/server.go:2448
	_go_fuzz_dep_.CoverTab[42864]++

//line /usr/local/go/src/net/http/server.go:2452
	host := stripHostPort(r.Host)
							path := cleanPath(r.URL.Path)

//line /usr/local/go/src/net/http/server.go:2457
	if u, ok := mux.redirectToPathSlash(host, path, r.URL); ok {
//line /usr/local/go/src/net/http/server.go:2457
		_go_fuzz_dep_.CoverTab[42872]++
								return RedirectHandler(u.String(), StatusMovedPermanently), u.Path
//line /usr/local/go/src/net/http/server.go:2458
		// _ = "end of CoverTab[42872]"
	} else {
//line /usr/local/go/src/net/http/server.go:2459
		_go_fuzz_dep_.CoverTab[42873]++
//line /usr/local/go/src/net/http/server.go:2459
		// _ = "end of CoverTab[42873]"
//line /usr/local/go/src/net/http/server.go:2459
	}
//line /usr/local/go/src/net/http/server.go:2459
	// _ = "end of CoverTab[42864]"
//line /usr/local/go/src/net/http/server.go:2459
	_go_fuzz_dep_.CoverTab[42865]++

							if path != r.URL.Path {
//line /usr/local/go/src/net/http/server.go:2461
		_go_fuzz_dep_.CoverTab[42874]++
								_, pattern = mux.handler(host, path)
								u := &url.URL{Path: path, RawQuery: r.URL.RawQuery}
								return RedirectHandler(u.String(), StatusMovedPermanently), pattern
//line /usr/local/go/src/net/http/server.go:2464
		// _ = "end of CoverTab[42874]"
	} else {
//line /usr/local/go/src/net/http/server.go:2465
		_go_fuzz_dep_.CoverTab[42875]++
//line /usr/local/go/src/net/http/server.go:2465
		// _ = "end of CoverTab[42875]"
//line /usr/local/go/src/net/http/server.go:2465
	}
//line /usr/local/go/src/net/http/server.go:2465
	// _ = "end of CoverTab[42865]"
//line /usr/local/go/src/net/http/server.go:2465
	_go_fuzz_dep_.CoverTab[42866]++

							return mux.handler(host, r.URL.Path)
//line /usr/local/go/src/net/http/server.go:2467
	// _ = "end of CoverTab[42866]"
}

// handler is the main implementation of Handler.
//line /usr/local/go/src/net/http/server.go:2470
// The path is known to be in canonical form, except for CONNECT methods.
//line /usr/local/go/src/net/http/server.go:2472
func (mux *ServeMux) handler(host, path string) (h Handler, pattern string) {
//line /usr/local/go/src/net/http/server.go:2472
	_go_fuzz_dep_.CoverTab[42876]++
							mux.mu.RLock()
							defer mux.mu.RUnlock()

//line /usr/local/go/src/net/http/server.go:2477
	if mux.hosts {
//line /usr/local/go/src/net/http/server.go:2477
		_go_fuzz_dep_.CoverTab[42880]++
								h, pattern = mux.match(host + path)
//line /usr/local/go/src/net/http/server.go:2478
		// _ = "end of CoverTab[42880]"
	} else {
//line /usr/local/go/src/net/http/server.go:2479
		_go_fuzz_dep_.CoverTab[42881]++
//line /usr/local/go/src/net/http/server.go:2479
		// _ = "end of CoverTab[42881]"
//line /usr/local/go/src/net/http/server.go:2479
	}
//line /usr/local/go/src/net/http/server.go:2479
	// _ = "end of CoverTab[42876]"
//line /usr/local/go/src/net/http/server.go:2479
	_go_fuzz_dep_.CoverTab[42877]++
							if h == nil {
//line /usr/local/go/src/net/http/server.go:2480
		_go_fuzz_dep_.CoverTab[42882]++
								h, pattern = mux.match(path)
//line /usr/local/go/src/net/http/server.go:2481
		// _ = "end of CoverTab[42882]"
	} else {
//line /usr/local/go/src/net/http/server.go:2482
		_go_fuzz_dep_.CoverTab[42883]++
//line /usr/local/go/src/net/http/server.go:2482
		// _ = "end of CoverTab[42883]"
//line /usr/local/go/src/net/http/server.go:2482
	}
//line /usr/local/go/src/net/http/server.go:2482
	// _ = "end of CoverTab[42877]"
//line /usr/local/go/src/net/http/server.go:2482
	_go_fuzz_dep_.CoverTab[42878]++
							if h == nil {
//line /usr/local/go/src/net/http/server.go:2483
		_go_fuzz_dep_.CoverTab[42884]++
								h, pattern = NotFoundHandler(), ""
//line /usr/local/go/src/net/http/server.go:2484
		// _ = "end of CoverTab[42884]"
	} else {
//line /usr/local/go/src/net/http/server.go:2485
		_go_fuzz_dep_.CoverTab[42885]++
//line /usr/local/go/src/net/http/server.go:2485
		// _ = "end of CoverTab[42885]"
//line /usr/local/go/src/net/http/server.go:2485
	}
//line /usr/local/go/src/net/http/server.go:2485
	// _ = "end of CoverTab[42878]"
//line /usr/local/go/src/net/http/server.go:2485
	_go_fuzz_dep_.CoverTab[42879]++
							return
//line /usr/local/go/src/net/http/server.go:2486
	// _ = "end of CoverTab[42879]"
}

// ServeHTTP dispatches the request to the handler whose
//line /usr/local/go/src/net/http/server.go:2489
// pattern most closely matches the request URL.
//line /usr/local/go/src/net/http/server.go:2491
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
//line /usr/local/go/src/net/http/server.go:2491
	_go_fuzz_dep_.CoverTab[42886]++
							if r.RequestURI == "*" {
//line /usr/local/go/src/net/http/server.go:2492
		_go_fuzz_dep_.CoverTab[42888]++
								if r.ProtoAtLeast(1, 1) {
//line /usr/local/go/src/net/http/server.go:2493
			_go_fuzz_dep_.CoverTab[42890]++
									w.Header().Set("Connection", "close")
//line /usr/local/go/src/net/http/server.go:2494
			// _ = "end of CoverTab[42890]"
		} else {
//line /usr/local/go/src/net/http/server.go:2495
			_go_fuzz_dep_.CoverTab[42891]++
//line /usr/local/go/src/net/http/server.go:2495
			// _ = "end of CoverTab[42891]"
//line /usr/local/go/src/net/http/server.go:2495
		}
//line /usr/local/go/src/net/http/server.go:2495
		// _ = "end of CoverTab[42888]"
//line /usr/local/go/src/net/http/server.go:2495
		_go_fuzz_dep_.CoverTab[42889]++
								w.WriteHeader(StatusBadRequest)
								return
//line /usr/local/go/src/net/http/server.go:2497
		// _ = "end of CoverTab[42889]"
	} else {
//line /usr/local/go/src/net/http/server.go:2498
		_go_fuzz_dep_.CoverTab[42892]++
//line /usr/local/go/src/net/http/server.go:2498
		// _ = "end of CoverTab[42892]"
//line /usr/local/go/src/net/http/server.go:2498
	}
//line /usr/local/go/src/net/http/server.go:2498
	// _ = "end of CoverTab[42886]"
//line /usr/local/go/src/net/http/server.go:2498
	_go_fuzz_dep_.CoverTab[42887]++
							h, _ := mux.Handler(r)
							h.ServeHTTP(w, r)
//line /usr/local/go/src/net/http/server.go:2500
	// _ = "end of CoverTab[42887]"
}

// Handle registers the handler for the given pattern.
//line /usr/local/go/src/net/http/server.go:2503
// If a handler already exists for pattern, Handle panics.
//line /usr/local/go/src/net/http/server.go:2505
func (mux *ServeMux) Handle(pattern string, handler Handler) {
//line /usr/local/go/src/net/http/server.go:2505
	_go_fuzz_dep_.CoverTab[42893]++
							mux.mu.Lock()
							defer mux.mu.Unlock()

							if pattern == "" {
//line /usr/local/go/src/net/http/server.go:2509
		_go_fuzz_dep_.CoverTab[42899]++
								panic("http: invalid pattern")
//line /usr/local/go/src/net/http/server.go:2510
		// _ = "end of CoverTab[42899]"
	} else {
//line /usr/local/go/src/net/http/server.go:2511
		_go_fuzz_dep_.CoverTab[42900]++
//line /usr/local/go/src/net/http/server.go:2511
		// _ = "end of CoverTab[42900]"
//line /usr/local/go/src/net/http/server.go:2511
	}
//line /usr/local/go/src/net/http/server.go:2511
	// _ = "end of CoverTab[42893]"
//line /usr/local/go/src/net/http/server.go:2511
	_go_fuzz_dep_.CoverTab[42894]++
							if handler == nil {
//line /usr/local/go/src/net/http/server.go:2512
		_go_fuzz_dep_.CoverTab[42901]++
								panic("http: nil handler")
//line /usr/local/go/src/net/http/server.go:2513
		// _ = "end of CoverTab[42901]"
	} else {
//line /usr/local/go/src/net/http/server.go:2514
		_go_fuzz_dep_.CoverTab[42902]++
//line /usr/local/go/src/net/http/server.go:2514
		// _ = "end of CoverTab[42902]"
//line /usr/local/go/src/net/http/server.go:2514
	}
//line /usr/local/go/src/net/http/server.go:2514
	// _ = "end of CoverTab[42894]"
//line /usr/local/go/src/net/http/server.go:2514
	_go_fuzz_dep_.CoverTab[42895]++
							if _, exist := mux.m[pattern]; exist {
//line /usr/local/go/src/net/http/server.go:2515
		_go_fuzz_dep_.CoverTab[42903]++
								panic("http: multiple registrations for " + pattern)
//line /usr/local/go/src/net/http/server.go:2516
		// _ = "end of CoverTab[42903]"
	} else {
//line /usr/local/go/src/net/http/server.go:2517
		_go_fuzz_dep_.CoverTab[42904]++
//line /usr/local/go/src/net/http/server.go:2517
		// _ = "end of CoverTab[42904]"
//line /usr/local/go/src/net/http/server.go:2517
	}
//line /usr/local/go/src/net/http/server.go:2517
	// _ = "end of CoverTab[42895]"
//line /usr/local/go/src/net/http/server.go:2517
	_go_fuzz_dep_.CoverTab[42896]++

							if mux.m == nil {
//line /usr/local/go/src/net/http/server.go:2519
		_go_fuzz_dep_.CoverTab[42905]++
								mux.m = make(map[string]muxEntry)
//line /usr/local/go/src/net/http/server.go:2520
		// _ = "end of CoverTab[42905]"
	} else {
//line /usr/local/go/src/net/http/server.go:2521
		_go_fuzz_dep_.CoverTab[42906]++
//line /usr/local/go/src/net/http/server.go:2521
		// _ = "end of CoverTab[42906]"
//line /usr/local/go/src/net/http/server.go:2521
	}
//line /usr/local/go/src/net/http/server.go:2521
	// _ = "end of CoverTab[42896]"
//line /usr/local/go/src/net/http/server.go:2521
	_go_fuzz_dep_.CoverTab[42897]++
							e := muxEntry{h: handler, pattern: pattern}
							mux.m[pattern] = e
							if pattern[len(pattern)-1] == '/' {
//line /usr/local/go/src/net/http/server.go:2524
		_go_fuzz_dep_.CoverTab[42907]++
								mux.es = appendSorted(mux.es, e)
//line /usr/local/go/src/net/http/server.go:2525
		// _ = "end of CoverTab[42907]"
	} else {
//line /usr/local/go/src/net/http/server.go:2526
		_go_fuzz_dep_.CoverTab[42908]++
//line /usr/local/go/src/net/http/server.go:2526
		// _ = "end of CoverTab[42908]"
//line /usr/local/go/src/net/http/server.go:2526
	}
//line /usr/local/go/src/net/http/server.go:2526
	// _ = "end of CoverTab[42897]"
//line /usr/local/go/src/net/http/server.go:2526
	_go_fuzz_dep_.CoverTab[42898]++

							if pattern[0] != '/' {
//line /usr/local/go/src/net/http/server.go:2528
		_go_fuzz_dep_.CoverTab[42909]++
								mux.hosts = true
//line /usr/local/go/src/net/http/server.go:2529
		// _ = "end of CoverTab[42909]"
	} else {
//line /usr/local/go/src/net/http/server.go:2530
		_go_fuzz_dep_.CoverTab[42910]++
//line /usr/local/go/src/net/http/server.go:2530
		// _ = "end of CoverTab[42910]"
//line /usr/local/go/src/net/http/server.go:2530
	}
//line /usr/local/go/src/net/http/server.go:2530
	// _ = "end of CoverTab[42898]"
}

func appendSorted(es []muxEntry, e muxEntry) []muxEntry {
//line /usr/local/go/src/net/http/server.go:2533
	_go_fuzz_dep_.CoverTab[42911]++
							n := len(es)
							i := sort.Search(n, func(i int) bool {
//line /usr/local/go/src/net/http/server.go:2535
		_go_fuzz_dep_.CoverTab[42914]++
								return len(es[i].pattern) < len(e.pattern)
//line /usr/local/go/src/net/http/server.go:2536
		// _ = "end of CoverTab[42914]"
	})
//line /usr/local/go/src/net/http/server.go:2537
	// _ = "end of CoverTab[42911]"
//line /usr/local/go/src/net/http/server.go:2537
	_go_fuzz_dep_.CoverTab[42912]++
							if i == n {
//line /usr/local/go/src/net/http/server.go:2538
		_go_fuzz_dep_.CoverTab[42915]++
								return append(es, e)
//line /usr/local/go/src/net/http/server.go:2539
		// _ = "end of CoverTab[42915]"
	} else {
//line /usr/local/go/src/net/http/server.go:2540
		_go_fuzz_dep_.CoverTab[42916]++
//line /usr/local/go/src/net/http/server.go:2540
		// _ = "end of CoverTab[42916]"
//line /usr/local/go/src/net/http/server.go:2540
	}
//line /usr/local/go/src/net/http/server.go:2540
	// _ = "end of CoverTab[42912]"
//line /usr/local/go/src/net/http/server.go:2540
	_go_fuzz_dep_.CoverTab[42913]++

							es = append(es, muxEntry{})
							copy(es[i+1:], es[i:])
							es[i] = e
							return es
//line /usr/local/go/src/net/http/server.go:2545
	// _ = "end of CoverTab[42913]"
}

// HandleFunc registers the handler function for the given pattern.
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
//line /usr/local/go/src/net/http/server.go:2549
	_go_fuzz_dep_.CoverTab[42917]++
							if handler == nil {
//line /usr/local/go/src/net/http/server.go:2550
		_go_fuzz_dep_.CoverTab[42919]++
								panic("http: nil handler")
//line /usr/local/go/src/net/http/server.go:2551
		// _ = "end of CoverTab[42919]"
	} else {
//line /usr/local/go/src/net/http/server.go:2552
		_go_fuzz_dep_.CoverTab[42920]++
//line /usr/local/go/src/net/http/server.go:2552
		// _ = "end of CoverTab[42920]"
//line /usr/local/go/src/net/http/server.go:2552
	}
//line /usr/local/go/src/net/http/server.go:2552
	// _ = "end of CoverTab[42917]"
//line /usr/local/go/src/net/http/server.go:2552
	_go_fuzz_dep_.CoverTab[42918]++
							mux.Handle(pattern, HandlerFunc(handler))
//line /usr/local/go/src/net/http/server.go:2553
	// _ = "end of CoverTab[42918]"
}

// Handle registers the handler for the given pattern
//line /usr/local/go/src/net/http/server.go:2556
// in the DefaultServeMux.
//line /usr/local/go/src/net/http/server.go:2556
// The documentation for ServeMux explains how patterns are matched.
//line /usr/local/go/src/net/http/server.go:2559
func Handle(pattern string, handler Handler) {
//line /usr/local/go/src/net/http/server.go:2559
	_go_fuzz_dep_.CoverTab[42921]++
//line /usr/local/go/src/net/http/server.go:2559
	DefaultServeMux.Handle(pattern, handler)
//line /usr/local/go/src/net/http/server.go:2559
	// _ = "end of CoverTab[42921]"
//line /usr/local/go/src/net/http/server.go:2559
}

// HandleFunc registers the handler function for the given pattern
//line /usr/local/go/src/net/http/server.go:2561
// in the DefaultServeMux.
//line /usr/local/go/src/net/http/server.go:2561
// The documentation for ServeMux explains how patterns are matched.
//line /usr/local/go/src/net/http/server.go:2564
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
//line /usr/local/go/src/net/http/server.go:2564
	_go_fuzz_dep_.CoverTab[42922]++
							DefaultServeMux.HandleFunc(pattern, handler)
//line /usr/local/go/src/net/http/server.go:2565
	// _ = "end of CoverTab[42922]"
}

// Serve accepts incoming HTTP connections on the listener l,
//line /usr/local/go/src/net/http/server.go:2568
// creating a new service goroutine for each. The service goroutines
//line /usr/local/go/src/net/http/server.go:2568
// read requests and then call handler to reply to them.
//line /usr/local/go/src/net/http/server.go:2568
//
//line /usr/local/go/src/net/http/server.go:2568
// The handler is typically nil, in which case the DefaultServeMux is used.
//line /usr/local/go/src/net/http/server.go:2568
//
//line /usr/local/go/src/net/http/server.go:2568
// HTTP/2 support is only enabled if the Listener returns *tls.Conn
//line /usr/local/go/src/net/http/server.go:2568
// connections and they were configured with "h2" in the TLS
//line /usr/local/go/src/net/http/server.go:2568
// Config.NextProtos.
//line /usr/local/go/src/net/http/server.go:2568
//
//line /usr/local/go/src/net/http/server.go:2568
// Serve always returns a non-nil error.
//line /usr/local/go/src/net/http/server.go:2579
func Serve(l net.Listener, handler Handler) error {
//line /usr/local/go/src/net/http/server.go:2579
	_go_fuzz_dep_.CoverTab[42923]++
							srv := &Server{Handler: handler}
							return srv.Serve(l)
//line /usr/local/go/src/net/http/server.go:2581
	// _ = "end of CoverTab[42923]"
}

// ServeTLS accepts incoming HTTPS connections on the listener l,
//line /usr/local/go/src/net/http/server.go:2584
// creating a new service goroutine for each. The service goroutines
//line /usr/local/go/src/net/http/server.go:2584
// read requests and then call handler to reply to them.
//line /usr/local/go/src/net/http/server.go:2584
//
//line /usr/local/go/src/net/http/server.go:2584
// The handler is typically nil, in which case the DefaultServeMux is used.
//line /usr/local/go/src/net/http/server.go:2584
//
//line /usr/local/go/src/net/http/server.go:2584
// Additionally, files containing a certificate and matching private key
//line /usr/local/go/src/net/http/server.go:2584
// for the server must be provided. If the certificate is signed by a
//line /usr/local/go/src/net/http/server.go:2584
// certificate authority, the certFile should be the concatenation
//line /usr/local/go/src/net/http/server.go:2584
// of the server's certificate, any intermediates, and the CA's certificate.
//line /usr/local/go/src/net/http/server.go:2584
//
//line /usr/local/go/src/net/http/server.go:2584
// ServeTLS always returns a non-nil error.
//line /usr/local/go/src/net/http/server.go:2596
func ServeTLS(l net.Listener, handler Handler, certFile, keyFile string) error {
//line /usr/local/go/src/net/http/server.go:2596
	_go_fuzz_dep_.CoverTab[42924]++
							srv := &Server{Handler: handler}
							return srv.ServeTLS(l, certFile, keyFile)
//line /usr/local/go/src/net/http/server.go:2598
	// _ = "end of CoverTab[42924]"
}

// A Server defines parameters for running an HTTP server.
//line /usr/local/go/src/net/http/server.go:2601
// The zero value for Server is a valid configuration.
//line /usr/local/go/src/net/http/server.go:2603
type Server struct {
	// Addr optionally specifies the TCP address for the server to listen on,
	// in the form "host:port". If empty, ":http" (port 80) is used.
	// The service names are defined in RFC 6335 and assigned by IANA.
	// See net.Dial for details of the address format.
	Addr	string

	Handler	Handler	// handler to invoke, http.DefaultServeMux if nil

	// DisableGeneralOptionsHandler, if true, passes "OPTIONS *" requests to the Handler,
	// otherwise responds with 200 OK and Content-Length: 0.
	DisableGeneralOptionsHandler	bool

	// TLSConfig optionally provides a TLS configuration for use
	// by ServeTLS and ListenAndServeTLS. Note that this value is
	// cloned by ServeTLS and ListenAndServeTLS, so it's not
	// possible to modify the configuration with methods like
	// tls.Config.SetSessionTicketKeys. To use
	// SetSessionTicketKeys, use Server.Serve with a TLS Listener
	// instead.
	TLSConfig	*tls.Config

	// ReadTimeout is the maximum duration for reading the entire
	// request, including the body. A zero or negative value means
	// there will be no timeout.
	//
	// Because ReadTimeout does not let Handlers make per-request
	// decisions on each request body's acceptable deadline or
	// upload rate, most users will prefer to use
	// ReadHeaderTimeout. It is valid to use them both.
	ReadTimeout	time.Duration

	// ReadHeaderTimeout is the amount of time allowed to read
	// request headers. The connection's read deadline is reset
	// after reading the headers and the Handler can decide what
	// is considered too slow for the body. If ReadHeaderTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	ReadHeaderTimeout	time.Duration

	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
	// A zero or negative value means there will be no timeout.
	WriteTimeout	time.Duration

	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alives are enabled. If IdleTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	IdleTimeout	time.Duration

	// MaxHeaderBytes controls the maximum number of bytes the
	// server will read parsing the request header's keys and
	// values, including the request line. It does not limit the
	// size of the request body.
	// If zero, DefaultMaxHeaderBytes is used.
	MaxHeaderBytes	int

	// TLSNextProto optionally specifies a function to take over
	// ownership of the provided TLS connection when an ALPN
	// protocol upgrade has occurred. The map key is the protocol
	// name negotiated. The Handler argument should be used to
	// handle HTTP requests and will initialize the Request's TLS
	// and RemoteAddr if not already set. The connection is
	// automatically closed when the function returns.
	// If TLSNextProto is not nil, HTTP/2 support is not enabled
	// automatically.
	TLSNextProto	map[string]func(*Server, *tls.Conn, Handler)

	// ConnState specifies an optional callback function that is
	// called when a client connection changes state. See the
	// ConnState type and associated constants for details.
	ConnState	func(net.Conn, ConnState)

	// ErrorLog specifies an optional logger for errors accepting
	// connections, unexpected behavior from handlers, and
	// underlying FileSystem errors.
	// If nil, logging is done via the log package's standard logger.
	ErrorLog	*log.Logger

	// BaseContext optionally specifies a function that returns
	// the base context for incoming requests on this server.
	// The provided Listener is the specific Listener that's
	// about to start accepting requests.
	// If BaseContext is nil, the default is context.Background().
	// If non-nil, it must return a non-nil context.
	BaseContext	func(net.Listener) context.Context

	// ConnContext optionally specifies a function that modifies
	// the context used for a new connection c. The provided ctx
	// is derived from the base context and has a ServerContextKey
	// value.
	ConnContext	func(ctx context.Context, c net.Conn) context.Context

	inShutdown	atomic.Bool	// true when server is in shutdown

	disableKeepAlives	atomic.Bool
	nextProtoOnce		sync.Once	// guards setupHTTP2_* init
	nextProtoErr		error		// result of http2.ConfigureServer if used

	mu		sync.Mutex
	listeners	map[*net.Listener]struct{}
	activeConn	map[*conn]struct{}
	onShutdown	[]func()

	listenerGroup	sync.WaitGroup
}

// Close immediately closes all active net.Listeners and any
//line /usr/local/go/src/net/http/server.go:2713
// connections in state StateNew, StateActive, or StateIdle. For a
//line /usr/local/go/src/net/http/server.go:2713
// graceful shutdown, use Shutdown.
//line /usr/local/go/src/net/http/server.go:2713
//
//line /usr/local/go/src/net/http/server.go:2713
// Close does not attempt to close (and does not even know about)
//line /usr/local/go/src/net/http/server.go:2713
// any hijacked connections, such as WebSockets.
//line /usr/local/go/src/net/http/server.go:2713
//
//line /usr/local/go/src/net/http/server.go:2713
// Close returns any error returned from closing the Server's
//line /usr/local/go/src/net/http/server.go:2713
// underlying Listener(s).
//line /usr/local/go/src/net/http/server.go:2722
func (srv *Server) Close() error {
//line /usr/local/go/src/net/http/server.go:2722
	_go_fuzz_dep_.CoverTab[42925]++
							srv.inShutdown.Store(true)
							srv.mu.Lock()
							defer srv.mu.Unlock()
							err := srv.closeListenersLocked()

//line /usr/local/go/src/net/http/server.go:2732
	srv.mu.Unlock()
	srv.listenerGroup.Wait()
	srv.mu.Lock()

	for c := range srv.activeConn {
//line /usr/local/go/src/net/http/server.go:2736
		_go_fuzz_dep_.CoverTab[42927]++
								c.rwc.Close()
								delete(srv.activeConn, c)
//line /usr/local/go/src/net/http/server.go:2738
		// _ = "end of CoverTab[42927]"
	}
//line /usr/local/go/src/net/http/server.go:2739
	// _ = "end of CoverTab[42925]"
//line /usr/local/go/src/net/http/server.go:2739
	_go_fuzz_dep_.CoverTab[42926]++
							return err
//line /usr/local/go/src/net/http/server.go:2740
	// _ = "end of CoverTab[42926]"
}

// shutdownPollIntervalMax is the max polling interval when checking
//line /usr/local/go/src/net/http/server.go:2743
// quiescence during Server.Shutdown. Polling starts with a small
//line /usr/local/go/src/net/http/server.go:2743
// interval and backs off to the max.
//line /usr/local/go/src/net/http/server.go:2743
// Ideally we could find a solution that doesn't involve polling,
//line /usr/local/go/src/net/http/server.go:2743
// but which also doesn't have a high runtime cost (and doesn't
//line /usr/local/go/src/net/http/server.go:2743
// involve any contentious mutexes), but that is left as an
//line /usr/local/go/src/net/http/server.go:2743
// exercise for the reader.
//line /usr/local/go/src/net/http/server.go:2750
const shutdownPollIntervalMax = 500 * time.Millisecond

// Shutdown gracefully shuts down the server without interrupting any
//line /usr/local/go/src/net/http/server.go:2752
// active connections. Shutdown works by first closing all open
//line /usr/local/go/src/net/http/server.go:2752
// listeners, then closing all idle connections, and then waiting
//line /usr/local/go/src/net/http/server.go:2752
// indefinitely for connections to return to idle and then shut down.
//line /usr/local/go/src/net/http/server.go:2752
// If the provided context expires before the shutdown is complete,
//line /usr/local/go/src/net/http/server.go:2752
// Shutdown returns the context's error, otherwise it returns any
//line /usr/local/go/src/net/http/server.go:2752
// error returned from closing the Server's underlying Listener(s).
//line /usr/local/go/src/net/http/server.go:2752
//
//line /usr/local/go/src/net/http/server.go:2752
// When Shutdown is called, Serve, ListenAndServe, and
//line /usr/local/go/src/net/http/server.go:2752
// ListenAndServeTLS immediately return ErrServerClosed. Make sure the
//line /usr/local/go/src/net/http/server.go:2752
// program doesn't exit and waits instead for Shutdown to return.
//line /usr/local/go/src/net/http/server.go:2752
//
//line /usr/local/go/src/net/http/server.go:2752
// Shutdown does not attempt to close nor wait for hijacked
//line /usr/local/go/src/net/http/server.go:2752
// connections such as WebSockets. The caller of Shutdown should
//line /usr/local/go/src/net/http/server.go:2752
// separately notify such long-lived connections of shutdown and wait
//line /usr/local/go/src/net/http/server.go:2752
// for them to close, if desired. See RegisterOnShutdown for a way to
//line /usr/local/go/src/net/http/server.go:2752
// register shutdown notification functions.
//line /usr/local/go/src/net/http/server.go:2752
//
//line /usr/local/go/src/net/http/server.go:2752
// Once Shutdown has been called on a server, it may not be reused;
//line /usr/local/go/src/net/http/server.go:2752
// future calls to methods such as Serve will return ErrServerClosed.
//line /usr/local/go/src/net/http/server.go:2772
func (srv *Server) Shutdown(ctx context.Context) error {
//line /usr/local/go/src/net/http/server.go:2772
	_go_fuzz_dep_.CoverTab[42928]++
							srv.inShutdown.Store(true)

							srv.mu.Lock()
							lnerr := srv.closeListenersLocked()
							for _, f := range srv.onShutdown {
//line /usr/local/go/src/net/http/server.go:2777
		_go_fuzz_dep_.CoverTab[42931]++
//line /usr/local/go/src/net/http/server.go:2777
		_curRoutineNum32_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/server.go:2777
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum32_)
								go func() {
//line /usr/local/go/src/net/http/server.go:2778
			_go_fuzz_dep_.CoverTab[42932]++
//line /usr/local/go/src/net/http/server.go:2778
			defer func() {
//line /usr/local/go/src/net/http/server.go:2778
				_go_fuzz_dep_.CoverTab[42933]++
//line /usr/local/go/src/net/http/server.go:2778
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum32_)
//line /usr/local/go/src/net/http/server.go:2778
				// _ = "end of CoverTab[42933]"
//line /usr/local/go/src/net/http/server.go:2778
			}()
//line /usr/local/go/src/net/http/server.go:2778
			f()
//line /usr/local/go/src/net/http/server.go:2778
			// _ = "end of CoverTab[42932]"
//line /usr/local/go/src/net/http/server.go:2778
		}()
//line /usr/local/go/src/net/http/server.go:2778
		// _ = "end of CoverTab[42931]"
	}
//line /usr/local/go/src/net/http/server.go:2779
	// _ = "end of CoverTab[42928]"
//line /usr/local/go/src/net/http/server.go:2779
	_go_fuzz_dep_.CoverTab[42929]++
							srv.mu.Unlock()
							srv.listenerGroup.Wait()

							pollIntervalBase := time.Millisecond
							nextPollInterval := func() time.Duration {
//line /usr/local/go/src/net/http/server.go:2784
		_go_fuzz_dep_.CoverTab[42934]++

								interval := pollIntervalBase + time.Duration(rand.Intn(int(pollIntervalBase/10)))

								pollIntervalBase *= 2
								if pollIntervalBase > shutdownPollIntervalMax {
//line /usr/local/go/src/net/http/server.go:2789
			_go_fuzz_dep_.CoverTab[42936]++
									pollIntervalBase = shutdownPollIntervalMax
//line /usr/local/go/src/net/http/server.go:2790
			// _ = "end of CoverTab[42936]"
		} else {
//line /usr/local/go/src/net/http/server.go:2791
			_go_fuzz_dep_.CoverTab[42937]++
//line /usr/local/go/src/net/http/server.go:2791
			// _ = "end of CoverTab[42937]"
//line /usr/local/go/src/net/http/server.go:2791
		}
//line /usr/local/go/src/net/http/server.go:2791
		// _ = "end of CoverTab[42934]"
//line /usr/local/go/src/net/http/server.go:2791
		_go_fuzz_dep_.CoverTab[42935]++
								return interval
//line /usr/local/go/src/net/http/server.go:2792
		// _ = "end of CoverTab[42935]"
	}
//line /usr/local/go/src/net/http/server.go:2793
	// _ = "end of CoverTab[42929]"
//line /usr/local/go/src/net/http/server.go:2793
	_go_fuzz_dep_.CoverTab[42930]++

							timer := time.NewTimer(nextPollInterval())
							defer timer.Stop()
							for {
//line /usr/local/go/src/net/http/server.go:2797
		_go_fuzz_dep_.CoverTab[42938]++
								if srv.closeIdleConns() {
//line /usr/local/go/src/net/http/server.go:2798
			_go_fuzz_dep_.CoverTab[42940]++
									return lnerr
//line /usr/local/go/src/net/http/server.go:2799
			// _ = "end of CoverTab[42940]"
		} else {
//line /usr/local/go/src/net/http/server.go:2800
			_go_fuzz_dep_.CoverTab[42941]++
//line /usr/local/go/src/net/http/server.go:2800
			// _ = "end of CoverTab[42941]"
//line /usr/local/go/src/net/http/server.go:2800
		}
//line /usr/local/go/src/net/http/server.go:2800
		// _ = "end of CoverTab[42938]"
//line /usr/local/go/src/net/http/server.go:2800
		_go_fuzz_dep_.CoverTab[42939]++
								select {
		case <-ctx.Done():
//line /usr/local/go/src/net/http/server.go:2802
			_go_fuzz_dep_.CoverTab[42942]++
									return ctx.Err()
//line /usr/local/go/src/net/http/server.go:2803
			// _ = "end of CoverTab[42942]"
		case <-timer.C:
//line /usr/local/go/src/net/http/server.go:2804
			_go_fuzz_dep_.CoverTab[42943]++
									timer.Reset(nextPollInterval())
//line /usr/local/go/src/net/http/server.go:2805
			// _ = "end of CoverTab[42943]"
		}
//line /usr/local/go/src/net/http/server.go:2806
		// _ = "end of CoverTab[42939]"
	}
//line /usr/local/go/src/net/http/server.go:2807
	// _ = "end of CoverTab[42930]"
}

// RegisterOnShutdown registers a function to call on Shutdown.
//line /usr/local/go/src/net/http/server.go:2810
// This can be used to gracefully shutdown connections that have
//line /usr/local/go/src/net/http/server.go:2810
// undergone ALPN protocol upgrade or that have been hijacked.
//line /usr/local/go/src/net/http/server.go:2810
// This function should start protocol-specific graceful shutdown,
//line /usr/local/go/src/net/http/server.go:2810
// but should not wait for shutdown to complete.
//line /usr/local/go/src/net/http/server.go:2815
func (srv *Server) RegisterOnShutdown(f func()) {
//line /usr/local/go/src/net/http/server.go:2815
	_go_fuzz_dep_.CoverTab[42944]++
							srv.mu.Lock()
							srv.onShutdown = append(srv.onShutdown, f)
							srv.mu.Unlock()
//line /usr/local/go/src/net/http/server.go:2818
	// _ = "end of CoverTab[42944]"
}

// closeIdleConns closes all idle connections and reports whether the
//line /usr/local/go/src/net/http/server.go:2821
// server is quiescent.
//line /usr/local/go/src/net/http/server.go:2823
func (s *Server) closeIdleConns() bool {
//line /usr/local/go/src/net/http/server.go:2823
	_go_fuzz_dep_.CoverTab[42945]++
							s.mu.Lock()
							defer s.mu.Unlock()
							quiescent := true
							for c := range s.activeConn {
//line /usr/local/go/src/net/http/server.go:2827
		_go_fuzz_dep_.CoverTab[42947]++
								st, unixSec := c.getState()

//line /usr/local/go/src/net/http/server.go:2832
		if st == StateNew && func() bool {
//line /usr/local/go/src/net/http/server.go:2832
			_go_fuzz_dep_.CoverTab[42950]++
//line /usr/local/go/src/net/http/server.go:2832
			return unixSec < time.Now().Unix()-5
//line /usr/local/go/src/net/http/server.go:2832
			// _ = "end of CoverTab[42950]"
//line /usr/local/go/src/net/http/server.go:2832
		}() {
//line /usr/local/go/src/net/http/server.go:2832
			_go_fuzz_dep_.CoverTab[42951]++
									st = StateIdle
//line /usr/local/go/src/net/http/server.go:2833
			// _ = "end of CoverTab[42951]"
		} else {
//line /usr/local/go/src/net/http/server.go:2834
			_go_fuzz_dep_.CoverTab[42952]++
//line /usr/local/go/src/net/http/server.go:2834
			// _ = "end of CoverTab[42952]"
//line /usr/local/go/src/net/http/server.go:2834
		}
//line /usr/local/go/src/net/http/server.go:2834
		// _ = "end of CoverTab[42947]"
//line /usr/local/go/src/net/http/server.go:2834
		_go_fuzz_dep_.CoverTab[42948]++
								if st != StateIdle || func() bool {
//line /usr/local/go/src/net/http/server.go:2835
			_go_fuzz_dep_.CoverTab[42953]++
//line /usr/local/go/src/net/http/server.go:2835
			return unixSec == 0
//line /usr/local/go/src/net/http/server.go:2835
			// _ = "end of CoverTab[42953]"
//line /usr/local/go/src/net/http/server.go:2835
		}() {
//line /usr/local/go/src/net/http/server.go:2835
			_go_fuzz_dep_.CoverTab[42954]++

//line /usr/local/go/src/net/http/server.go:2838
			quiescent = false
									continue
//line /usr/local/go/src/net/http/server.go:2839
			// _ = "end of CoverTab[42954]"
		} else {
//line /usr/local/go/src/net/http/server.go:2840
			_go_fuzz_dep_.CoverTab[42955]++
//line /usr/local/go/src/net/http/server.go:2840
			// _ = "end of CoverTab[42955]"
//line /usr/local/go/src/net/http/server.go:2840
		}
//line /usr/local/go/src/net/http/server.go:2840
		// _ = "end of CoverTab[42948]"
//line /usr/local/go/src/net/http/server.go:2840
		_go_fuzz_dep_.CoverTab[42949]++
								c.rwc.Close()
								delete(s.activeConn, c)
//line /usr/local/go/src/net/http/server.go:2842
		// _ = "end of CoverTab[42949]"
	}
//line /usr/local/go/src/net/http/server.go:2843
	// _ = "end of CoverTab[42945]"
//line /usr/local/go/src/net/http/server.go:2843
	_go_fuzz_dep_.CoverTab[42946]++
							return quiescent
//line /usr/local/go/src/net/http/server.go:2844
	// _ = "end of CoverTab[42946]"
}

func (s *Server) closeListenersLocked() error {
//line /usr/local/go/src/net/http/server.go:2847
	_go_fuzz_dep_.CoverTab[42956]++
							var err error
							for ln := range s.listeners {
//line /usr/local/go/src/net/http/server.go:2849
		_go_fuzz_dep_.CoverTab[42958]++
								if cerr := (*ln).Close(); cerr != nil && func() bool {
//line /usr/local/go/src/net/http/server.go:2850
			_go_fuzz_dep_.CoverTab[42959]++
//line /usr/local/go/src/net/http/server.go:2850
			return err == nil
//line /usr/local/go/src/net/http/server.go:2850
			// _ = "end of CoverTab[42959]"
//line /usr/local/go/src/net/http/server.go:2850
		}() {
//line /usr/local/go/src/net/http/server.go:2850
			_go_fuzz_dep_.CoverTab[42960]++
									err = cerr
//line /usr/local/go/src/net/http/server.go:2851
			// _ = "end of CoverTab[42960]"
		} else {
//line /usr/local/go/src/net/http/server.go:2852
			_go_fuzz_dep_.CoverTab[42961]++
//line /usr/local/go/src/net/http/server.go:2852
			// _ = "end of CoverTab[42961]"
//line /usr/local/go/src/net/http/server.go:2852
		}
//line /usr/local/go/src/net/http/server.go:2852
		// _ = "end of CoverTab[42958]"
	}
//line /usr/local/go/src/net/http/server.go:2853
	// _ = "end of CoverTab[42956]"
//line /usr/local/go/src/net/http/server.go:2853
	_go_fuzz_dep_.CoverTab[42957]++
							return err
//line /usr/local/go/src/net/http/server.go:2854
	// _ = "end of CoverTab[42957]"
}

// A ConnState represents the state of a client connection to a server.
//line /usr/local/go/src/net/http/server.go:2857
// It's used by the optional Server.ConnState hook.
//line /usr/local/go/src/net/http/server.go:2859
type ConnState int

const (
	// StateNew represents a new connection that is expected to
	// send a request immediately. Connections begin at this
	// state and then transition to either StateActive or
	// StateClosed.
	StateNew	ConnState	= iota

	// StateActive represents a connection that has read 1 or more
	// bytes of a request. The Server.ConnState hook for
	// StateActive fires before the request has entered a handler
	// and doesn't fire again until the request has been
	// handled. After the request is handled, the state
	// transitions to StateClosed, StateHijacked, or StateIdle.
	// For HTTP/2, StateActive fires on the transition from zero
	// to one active request, and only transitions away once all
	// active requests are complete. That means that ConnState
	// cannot be used to do per-request work; ConnState only notes
	// the overall state of the connection.
	StateActive

	// StateIdle represents a connection that has finished
	// handling a request and is in the keep-alive state, waiting
	// for a new request. Connections transition from StateIdle
	// to either StateActive or StateClosed.
	StateIdle

	// StateHijacked represents a hijacked connection.
	// This is a terminal state. It does not transition to StateClosed.
	StateHijacked

	// StateClosed represents a closed connection.
	// This is a terminal state. Hijacked connections do not
	// transition to StateClosed.
	StateClosed
)

var stateName = map[ConnState]string{
	StateNew:	"new",
	StateActive:	"active",
	StateIdle:	"idle",
	StateHijacked:	"hijacked",
	StateClosed:	"closed",
}

func (c ConnState) String() string {
//line /usr/local/go/src/net/http/server.go:2905
	_go_fuzz_dep_.CoverTab[42962]++
							return stateName[c]
//line /usr/local/go/src/net/http/server.go:2906
	// _ = "end of CoverTab[42962]"
}

// serverHandler delegates to either the server's Handler or
//line /usr/local/go/src/net/http/server.go:2909
// DefaultServeMux and also handles "OPTIONS *" requests.
//line /usr/local/go/src/net/http/server.go:2911
type serverHandler struct {
	srv *Server
}

func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
//line /usr/local/go/src/net/http/server.go:2915
	_go_fuzz_dep_.CoverTab[42963]++
							handler := sh.srv.Handler
							if handler == nil {
//line /usr/local/go/src/net/http/server.go:2917
		_go_fuzz_dep_.CoverTab[42967]++
								handler = DefaultServeMux
//line /usr/local/go/src/net/http/server.go:2918
		// _ = "end of CoverTab[42967]"
	} else {
//line /usr/local/go/src/net/http/server.go:2919
		_go_fuzz_dep_.CoverTab[42968]++
//line /usr/local/go/src/net/http/server.go:2919
		// _ = "end of CoverTab[42968]"
//line /usr/local/go/src/net/http/server.go:2919
	}
//line /usr/local/go/src/net/http/server.go:2919
	// _ = "end of CoverTab[42963]"
//line /usr/local/go/src/net/http/server.go:2919
	_go_fuzz_dep_.CoverTab[42964]++
							if !sh.srv.DisableGeneralOptionsHandler && func() bool {
//line /usr/local/go/src/net/http/server.go:2920
		_go_fuzz_dep_.CoverTab[42969]++
//line /usr/local/go/src/net/http/server.go:2920
		return req.RequestURI == "*"
//line /usr/local/go/src/net/http/server.go:2920
		// _ = "end of CoverTab[42969]"
//line /usr/local/go/src/net/http/server.go:2920
	}() && func() bool {
//line /usr/local/go/src/net/http/server.go:2920
		_go_fuzz_dep_.CoverTab[42970]++
//line /usr/local/go/src/net/http/server.go:2920
		return req.Method == "OPTIONS"
//line /usr/local/go/src/net/http/server.go:2920
		// _ = "end of CoverTab[42970]"
//line /usr/local/go/src/net/http/server.go:2920
	}() {
//line /usr/local/go/src/net/http/server.go:2920
		_go_fuzz_dep_.CoverTab[42971]++
								handler = globalOptionsHandler{}
//line /usr/local/go/src/net/http/server.go:2921
		// _ = "end of CoverTab[42971]"
	} else {
//line /usr/local/go/src/net/http/server.go:2922
		_go_fuzz_dep_.CoverTab[42972]++
//line /usr/local/go/src/net/http/server.go:2922
		// _ = "end of CoverTab[42972]"
//line /usr/local/go/src/net/http/server.go:2922
	}
//line /usr/local/go/src/net/http/server.go:2922
	// _ = "end of CoverTab[42964]"
//line /usr/local/go/src/net/http/server.go:2922
	_go_fuzz_dep_.CoverTab[42965]++

							if req.URL != nil && func() bool {
//line /usr/local/go/src/net/http/server.go:2924
		_go_fuzz_dep_.CoverTab[42973]++
//line /usr/local/go/src/net/http/server.go:2924
		return strings.Contains(req.URL.RawQuery, ";")
//line /usr/local/go/src/net/http/server.go:2924
		// _ = "end of CoverTab[42973]"
//line /usr/local/go/src/net/http/server.go:2924
	}() {
//line /usr/local/go/src/net/http/server.go:2924
		_go_fuzz_dep_.CoverTab[42974]++
								var allowQuerySemicolonsInUse atomic.Bool
								req = req.WithContext(context.WithValue(req.Context(), silenceSemWarnContextKey, func() {
//line /usr/local/go/src/net/http/server.go:2926
			_go_fuzz_dep_.CoverTab[42976]++
									allowQuerySemicolonsInUse.Store(true)
//line /usr/local/go/src/net/http/server.go:2927
			// _ = "end of CoverTab[42976]"
		}))
//line /usr/local/go/src/net/http/server.go:2928
		// _ = "end of CoverTab[42974]"
//line /usr/local/go/src/net/http/server.go:2928
		_go_fuzz_dep_.CoverTab[42975]++
								defer func() {
//line /usr/local/go/src/net/http/server.go:2929
			_go_fuzz_dep_.CoverTab[42977]++
									if !allowQuerySemicolonsInUse.Load() {
//line /usr/local/go/src/net/http/server.go:2930
				_go_fuzz_dep_.CoverTab[42978]++
										sh.srv.logf("http: URL query contains semicolon, which is no longer a supported separator; parts of the query may be stripped when parsed; see golang.org/issue/25192")
//line /usr/local/go/src/net/http/server.go:2931
				// _ = "end of CoverTab[42978]"
			} else {
//line /usr/local/go/src/net/http/server.go:2932
				_go_fuzz_dep_.CoverTab[42979]++
//line /usr/local/go/src/net/http/server.go:2932
				// _ = "end of CoverTab[42979]"
//line /usr/local/go/src/net/http/server.go:2932
			}
//line /usr/local/go/src/net/http/server.go:2932
			// _ = "end of CoverTab[42977]"
		}()
//line /usr/local/go/src/net/http/server.go:2933
		// _ = "end of CoverTab[42975]"
	} else {
//line /usr/local/go/src/net/http/server.go:2934
		_go_fuzz_dep_.CoverTab[42980]++
//line /usr/local/go/src/net/http/server.go:2934
		// _ = "end of CoverTab[42980]"
//line /usr/local/go/src/net/http/server.go:2934
	}
//line /usr/local/go/src/net/http/server.go:2934
	// _ = "end of CoverTab[42965]"
//line /usr/local/go/src/net/http/server.go:2934
	_go_fuzz_dep_.CoverTab[42966]++

							handler.ServeHTTP(rw, req)
//line /usr/local/go/src/net/http/server.go:2936
	// _ = "end of CoverTab[42966]"
}

var silenceSemWarnContextKey = &contextKey{"silence-semicolons"}

// AllowQuerySemicolons returns a handler that serves requests by converting any
//line /usr/local/go/src/net/http/server.go:2941
// unescaped semicolons in the URL query to ampersands, and invoking the handler h.
//line /usr/local/go/src/net/http/server.go:2941
//
//line /usr/local/go/src/net/http/server.go:2941
// This restores the pre-Go 1.17 behavior of splitting query parameters on both
//line /usr/local/go/src/net/http/server.go:2941
// semicolons and ampersands. (See golang.org/issue/25192). Note that this
//line /usr/local/go/src/net/http/server.go:2941
// behavior doesn't match that of many proxies, and the mismatch can lead to
//line /usr/local/go/src/net/http/server.go:2941
// security issues.
//line /usr/local/go/src/net/http/server.go:2941
//
//line /usr/local/go/src/net/http/server.go:2941
// AllowQuerySemicolons should be invoked before Request.ParseForm is called.
//line /usr/local/go/src/net/http/server.go:2950
func AllowQuerySemicolons(h Handler) Handler {
//line /usr/local/go/src/net/http/server.go:2950
	_go_fuzz_dep_.CoverTab[42981]++
							return HandlerFunc(func(w ResponseWriter, r *Request) {
//line /usr/local/go/src/net/http/server.go:2951
		_go_fuzz_dep_.CoverTab[42982]++
								if silenceSemicolonsWarning, ok := r.Context().Value(silenceSemWarnContextKey).(func()); ok {
//line /usr/local/go/src/net/http/server.go:2952
			_go_fuzz_dep_.CoverTab[42984]++
									silenceSemicolonsWarning()
//line /usr/local/go/src/net/http/server.go:2953
			// _ = "end of CoverTab[42984]"
		} else {
//line /usr/local/go/src/net/http/server.go:2954
			_go_fuzz_dep_.CoverTab[42985]++
//line /usr/local/go/src/net/http/server.go:2954
			// _ = "end of CoverTab[42985]"
//line /usr/local/go/src/net/http/server.go:2954
		}
//line /usr/local/go/src/net/http/server.go:2954
		// _ = "end of CoverTab[42982]"
//line /usr/local/go/src/net/http/server.go:2954
		_go_fuzz_dep_.CoverTab[42983]++
								if strings.Contains(r.URL.RawQuery, ";") {
//line /usr/local/go/src/net/http/server.go:2955
			_go_fuzz_dep_.CoverTab[42986]++
									r2 := new(Request)
									*r2 = *r
									r2.URL = new(url.URL)
									*r2.URL = *r.URL
									r2.URL.RawQuery = strings.ReplaceAll(r.URL.RawQuery, ";", "&")
									h.ServeHTTP(w, r2)
//line /usr/local/go/src/net/http/server.go:2961
			// _ = "end of CoverTab[42986]"
		} else {
//line /usr/local/go/src/net/http/server.go:2962
			_go_fuzz_dep_.CoverTab[42987]++
									h.ServeHTTP(w, r)
//line /usr/local/go/src/net/http/server.go:2963
			// _ = "end of CoverTab[42987]"
		}
//line /usr/local/go/src/net/http/server.go:2964
		// _ = "end of CoverTab[42983]"
	})
//line /usr/local/go/src/net/http/server.go:2965
	// _ = "end of CoverTab[42981]"
}

// ListenAndServe listens on the TCP network address srv.Addr and then
//line /usr/local/go/src/net/http/server.go:2968
// calls Serve to handle requests on incoming connections.
//line /usr/local/go/src/net/http/server.go:2968
// Accepted connections are configured to enable TCP keep-alives.
//line /usr/local/go/src/net/http/server.go:2968
//
//line /usr/local/go/src/net/http/server.go:2968
// If srv.Addr is blank, ":http" is used.
//line /usr/local/go/src/net/http/server.go:2968
//
//line /usr/local/go/src/net/http/server.go:2968
// ListenAndServe always returns a non-nil error. After Shutdown or Close,
//line /usr/local/go/src/net/http/server.go:2968
// the returned error is ErrServerClosed.
//line /usr/local/go/src/net/http/server.go:2976
func (srv *Server) ListenAndServe() error {
//line /usr/local/go/src/net/http/server.go:2976
	_go_fuzz_dep_.CoverTab[42988]++
							if srv.shuttingDown() {
//line /usr/local/go/src/net/http/server.go:2977
		_go_fuzz_dep_.CoverTab[42992]++
								return ErrServerClosed
//line /usr/local/go/src/net/http/server.go:2978
		// _ = "end of CoverTab[42992]"
	} else {
//line /usr/local/go/src/net/http/server.go:2979
		_go_fuzz_dep_.CoverTab[42993]++
//line /usr/local/go/src/net/http/server.go:2979
		// _ = "end of CoverTab[42993]"
//line /usr/local/go/src/net/http/server.go:2979
	}
//line /usr/local/go/src/net/http/server.go:2979
	// _ = "end of CoverTab[42988]"
//line /usr/local/go/src/net/http/server.go:2979
	_go_fuzz_dep_.CoverTab[42989]++
							addr := srv.Addr
							if addr == "" {
//line /usr/local/go/src/net/http/server.go:2981
		_go_fuzz_dep_.CoverTab[42994]++
								addr = ":http"
//line /usr/local/go/src/net/http/server.go:2982
		// _ = "end of CoverTab[42994]"
	} else {
//line /usr/local/go/src/net/http/server.go:2983
		_go_fuzz_dep_.CoverTab[42995]++
//line /usr/local/go/src/net/http/server.go:2983
		// _ = "end of CoverTab[42995]"
//line /usr/local/go/src/net/http/server.go:2983
	}
//line /usr/local/go/src/net/http/server.go:2983
	// _ = "end of CoverTab[42989]"
//line /usr/local/go/src/net/http/server.go:2983
	_go_fuzz_dep_.CoverTab[42990]++
							ln, err := net.Listen("tcp", addr)
							if err != nil {
//line /usr/local/go/src/net/http/server.go:2985
		_go_fuzz_dep_.CoverTab[42996]++
								return err
//line /usr/local/go/src/net/http/server.go:2986
		// _ = "end of CoverTab[42996]"
	} else {
//line /usr/local/go/src/net/http/server.go:2987
		_go_fuzz_dep_.CoverTab[42997]++
//line /usr/local/go/src/net/http/server.go:2987
		// _ = "end of CoverTab[42997]"
//line /usr/local/go/src/net/http/server.go:2987
	}
//line /usr/local/go/src/net/http/server.go:2987
	// _ = "end of CoverTab[42990]"
//line /usr/local/go/src/net/http/server.go:2987
	_go_fuzz_dep_.CoverTab[42991]++
							return srv.Serve(ln)
//line /usr/local/go/src/net/http/server.go:2988
	// _ = "end of CoverTab[42991]"
}

var testHookServerServe func(*Server, net.Listener)	// used if non-nil

// shouldDoServeHTTP2 reports whether Server.Serve should configure
//line /usr/local/go/src/net/http/server.go:2993
// automatic HTTP/2. (which sets up the srv.TLSNextProto map)
//line /usr/local/go/src/net/http/server.go:2995
func (srv *Server) shouldConfigureHTTP2ForServe() bool {
//line /usr/local/go/src/net/http/server.go:2995
	_go_fuzz_dep_.CoverTab[42998]++
							if srv.TLSConfig == nil {
//line /usr/local/go/src/net/http/server.go:2996
		_go_fuzz_dep_.CoverTab[43000]++

//line /usr/local/go/src/net/http/server.go:3003
		return true
//line /usr/local/go/src/net/http/server.go:3003
		// _ = "end of CoverTab[43000]"
	} else {
//line /usr/local/go/src/net/http/server.go:3004
		_go_fuzz_dep_.CoverTab[43001]++
//line /usr/local/go/src/net/http/server.go:3004
		// _ = "end of CoverTab[43001]"
//line /usr/local/go/src/net/http/server.go:3004
	}
//line /usr/local/go/src/net/http/server.go:3004
	// _ = "end of CoverTab[42998]"
//line /usr/local/go/src/net/http/server.go:3004
	_go_fuzz_dep_.CoverTab[42999]++

//line /usr/local/go/src/net/http/server.go:3012
	return strSliceContains(srv.TLSConfig.NextProtos, http2NextProtoTLS)
//line /usr/local/go/src/net/http/server.go:3012
	// _ = "end of CoverTab[42999]"
}

// ErrServerClosed is returned by the Server's Serve, ServeTLS, ListenAndServe,
//line /usr/local/go/src/net/http/server.go:3015
// and ListenAndServeTLS methods after a call to Shutdown or Close.
//line /usr/local/go/src/net/http/server.go:3017
var ErrServerClosed = errors.New("http: Server closed")

// Serve accepts incoming connections on the Listener l, creating a
//line /usr/local/go/src/net/http/server.go:3019
// new service goroutine for each. The service goroutines read requests and
//line /usr/local/go/src/net/http/server.go:3019
// then call srv.Handler to reply to them.
//line /usr/local/go/src/net/http/server.go:3019
//
//line /usr/local/go/src/net/http/server.go:3019
// HTTP/2 support is only enabled if the Listener returns *tls.Conn
//line /usr/local/go/src/net/http/server.go:3019
// connections and they were configured with "h2" in the TLS
//line /usr/local/go/src/net/http/server.go:3019
// Config.NextProtos.
//line /usr/local/go/src/net/http/server.go:3019
//
//line /usr/local/go/src/net/http/server.go:3019
// Serve always returns a non-nil error and closes l.
//line /usr/local/go/src/net/http/server.go:3019
// After Shutdown or Close, the returned error is ErrServerClosed.
//line /usr/local/go/src/net/http/server.go:3029
func (srv *Server) Serve(l net.Listener) error {
//line /usr/local/go/src/net/http/server.go:3029
	_go_fuzz_dep_.CoverTab[43002]++
							if fn := testHookServerServe; fn != nil {
//line /usr/local/go/src/net/http/server.go:3030
		_go_fuzz_dep_.CoverTab[43007]++
								fn(srv, l)
//line /usr/local/go/src/net/http/server.go:3031
		// _ = "end of CoverTab[43007]"
	} else {
//line /usr/local/go/src/net/http/server.go:3032
		_go_fuzz_dep_.CoverTab[43008]++
//line /usr/local/go/src/net/http/server.go:3032
		// _ = "end of CoverTab[43008]"
//line /usr/local/go/src/net/http/server.go:3032
	}
//line /usr/local/go/src/net/http/server.go:3032
	// _ = "end of CoverTab[43002]"
//line /usr/local/go/src/net/http/server.go:3032
	_go_fuzz_dep_.CoverTab[43003]++

							origListener := l
							l = &onceCloseListener{Listener: l}
							defer l.Close()

							if err := srv.setupHTTP2_Serve(); err != nil {
//line /usr/local/go/src/net/http/server.go:3038
		_go_fuzz_dep_.CoverTab[43009]++
								return err
//line /usr/local/go/src/net/http/server.go:3039
		// _ = "end of CoverTab[43009]"
	} else {
//line /usr/local/go/src/net/http/server.go:3040
		_go_fuzz_dep_.CoverTab[43010]++
//line /usr/local/go/src/net/http/server.go:3040
		// _ = "end of CoverTab[43010]"
//line /usr/local/go/src/net/http/server.go:3040
	}
//line /usr/local/go/src/net/http/server.go:3040
	// _ = "end of CoverTab[43003]"
//line /usr/local/go/src/net/http/server.go:3040
	_go_fuzz_dep_.CoverTab[43004]++

							if !srv.trackListener(&l, true) {
//line /usr/local/go/src/net/http/server.go:3042
		_go_fuzz_dep_.CoverTab[43011]++
								return ErrServerClosed
//line /usr/local/go/src/net/http/server.go:3043
		// _ = "end of CoverTab[43011]"
	} else {
//line /usr/local/go/src/net/http/server.go:3044
		_go_fuzz_dep_.CoverTab[43012]++
//line /usr/local/go/src/net/http/server.go:3044
		// _ = "end of CoverTab[43012]"
//line /usr/local/go/src/net/http/server.go:3044
	}
//line /usr/local/go/src/net/http/server.go:3044
	// _ = "end of CoverTab[43004]"
//line /usr/local/go/src/net/http/server.go:3044
	_go_fuzz_dep_.CoverTab[43005]++
							defer srv.trackListener(&l, false)

							baseCtx := context.Background()
							if srv.BaseContext != nil {
//line /usr/local/go/src/net/http/server.go:3048
		_go_fuzz_dep_.CoverTab[43013]++
								baseCtx = srv.BaseContext(origListener)
								if baseCtx == nil {
//line /usr/local/go/src/net/http/server.go:3050
			_go_fuzz_dep_.CoverTab[43014]++
									panic("BaseContext returned a nil context")
//line /usr/local/go/src/net/http/server.go:3051
			// _ = "end of CoverTab[43014]"
		} else {
//line /usr/local/go/src/net/http/server.go:3052
			_go_fuzz_dep_.CoverTab[43015]++
//line /usr/local/go/src/net/http/server.go:3052
			// _ = "end of CoverTab[43015]"
//line /usr/local/go/src/net/http/server.go:3052
		}
//line /usr/local/go/src/net/http/server.go:3052
		// _ = "end of CoverTab[43013]"
	} else {
//line /usr/local/go/src/net/http/server.go:3053
		_go_fuzz_dep_.CoverTab[43016]++
//line /usr/local/go/src/net/http/server.go:3053
		// _ = "end of CoverTab[43016]"
//line /usr/local/go/src/net/http/server.go:3053
	}
//line /usr/local/go/src/net/http/server.go:3053
	// _ = "end of CoverTab[43005]"
//line /usr/local/go/src/net/http/server.go:3053
	_go_fuzz_dep_.CoverTab[43006]++

							var tempDelay time.Duration	// how long to sleep on accept failure

							ctx := context.WithValue(baseCtx, ServerContextKey, srv)
							for {
//line /usr/local/go/src/net/http/server.go:3058
		_go_fuzz_dep_.CoverTab[43017]++
								rw, err := l.Accept()
								if err != nil {
//line /usr/local/go/src/net/http/server.go:3060
			_go_fuzz_dep_.CoverTab[43020]++
									if srv.shuttingDown() {
//line /usr/local/go/src/net/http/server.go:3061
				_go_fuzz_dep_.CoverTab[43023]++
										return ErrServerClosed
//line /usr/local/go/src/net/http/server.go:3062
				// _ = "end of CoverTab[43023]"
			} else {
//line /usr/local/go/src/net/http/server.go:3063
				_go_fuzz_dep_.CoverTab[43024]++
//line /usr/local/go/src/net/http/server.go:3063
				// _ = "end of CoverTab[43024]"
//line /usr/local/go/src/net/http/server.go:3063
			}
//line /usr/local/go/src/net/http/server.go:3063
			// _ = "end of CoverTab[43020]"
//line /usr/local/go/src/net/http/server.go:3063
			_go_fuzz_dep_.CoverTab[43021]++
									if ne, ok := err.(net.Error); ok && func() bool {
//line /usr/local/go/src/net/http/server.go:3064
				_go_fuzz_dep_.CoverTab[43025]++
//line /usr/local/go/src/net/http/server.go:3064
				return ne.Temporary()
//line /usr/local/go/src/net/http/server.go:3064
				// _ = "end of CoverTab[43025]"
//line /usr/local/go/src/net/http/server.go:3064
			}() {
//line /usr/local/go/src/net/http/server.go:3064
				_go_fuzz_dep_.CoverTab[43026]++
										if tempDelay == 0 {
//line /usr/local/go/src/net/http/server.go:3065
					_go_fuzz_dep_.CoverTab[43029]++
											tempDelay = 5 * time.Millisecond
//line /usr/local/go/src/net/http/server.go:3066
					// _ = "end of CoverTab[43029]"
				} else {
//line /usr/local/go/src/net/http/server.go:3067
					_go_fuzz_dep_.CoverTab[43030]++
											tempDelay *= 2
//line /usr/local/go/src/net/http/server.go:3068
					// _ = "end of CoverTab[43030]"
				}
//line /usr/local/go/src/net/http/server.go:3069
				// _ = "end of CoverTab[43026]"
//line /usr/local/go/src/net/http/server.go:3069
				_go_fuzz_dep_.CoverTab[43027]++
										if max := 1 * time.Second; tempDelay > max {
//line /usr/local/go/src/net/http/server.go:3070
					_go_fuzz_dep_.CoverTab[43031]++
											tempDelay = max
//line /usr/local/go/src/net/http/server.go:3071
					// _ = "end of CoverTab[43031]"
				} else {
//line /usr/local/go/src/net/http/server.go:3072
					_go_fuzz_dep_.CoverTab[43032]++
//line /usr/local/go/src/net/http/server.go:3072
					// _ = "end of CoverTab[43032]"
//line /usr/local/go/src/net/http/server.go:3072
				}
//line /usr/local/go/src/net/http/server.go:3072
				// _ = "end of CoverTab[43027]"
//line /usr/local/go/src/net/http/server.go:3072
				_go_fuzz_dep_.CoverTab[43028]++
										srv.logf("http: Accept error: %v; retrying in %v", err, tempDelay)
										time.Sleep(tempDelay)
										continue
//line /usr/local/go/src/net/http/server.go:3075
				// _ = "end of CoverTab[43028]"
			} else {
//line /usr/local/go/src/net/http/server.go:3076
				_go_fuzz_dep_.CoverTab[43033]++
//line /usr/local/go/src/net/http/server.go:3076
				// _ = "end of CoverTab[43033]"
//line /usr/local/go/src/net/http/server.go:3076
			}
//line /usr/local/go/src/net/http/server.go:3076
			// _ = "end of CoverTab[43021]"
//line /usr/local/go/src/net/http/server.go:3076
			_go_fuzz_dep_.CoverTab[43022]++
									return err
//line /usr/local/go/src/net/http/server.go:3077
			// _ = "end of CoverTab[43022]"
		} else {
//line /usr/local/go/src/net/http/server.go:3078
			_go_fuzz_dep_.CoverTab[43034]++
//line /usr/local/go/src/net/http/server.go:3078
			// _ = "end of CoverTab[43034]"
//line /usr/local/go/src/net/http/server.go:3078
		}
//line /usr/local/go/src/net/http/server.go:3078
		// _ = "end of CoverTab[43017]"
//line /usr/local/go/src/net/http/server.go:3078
		_go_fuzz_dep_.CoverTab[43018]++
								connCtx := ctx
								if cc := srv.ConnContext; cc != nil {
//line /usr/local/go/src/net/http/server.go:3080
			_go_fuzz_dep_.CoverTab[43035]++
									connCtx = cc(connCtx, rw)
									if connCtx == nil {
//line /usr/local/go/src/net/http/server.go:3082
				_go_fuzz_dep_.CoverTab[43036]++
										panic("ConnContext returned nil")
//line /usr/local/go/src/net/http/server.go:3083
				// _ = "end of CoverTab[43036]"
			} else {
//line /usr/local/go/src/net/http/server.go:3084
				_go_fuzz_dep_.CoverTab[43037]++
//line /usr/local/go/src/net/http/server.go:3084
				// _ = "end of CoverTab[43037]"
//line /usr/local/go/src/net/http/server.go:3084
			}
//line /usr/local/go/src/net/http/server.go:3084
			// _ = "end of CoverTab[43035]"
		} else {
//line /usr/local/go/src/net/http/server.go:3085
			_go_fuzz_dep_.CoverTab[43038]++
//line /usr/local/go/src/net/http/server.go:3085
			// _ = "end of CoverTab[43038]"
//line /usr/local/go/src/net/http/server.go:3085
		}
//line /usr/local/go/src/net/http/server.go:3085
		// _ = "end of CoverTab[43018]"
//line /usr/local/go/src/net/http/server.go:3085
		_go_fuzz_dep_.CoverTab[43019]++
								tempDelay = 0
								c := srv.newConn(rw)
								c.setState(c.rwc, StateNew, runHooks)
//line /usr/local/go/src/net/http/server.go:3088
		_curRoutineNum33_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/server.go:3088
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum33_)
								go c.serve(connCtx)
//line /usr/local/go/src/net/http/server.go:3089
		// _ = "end of CoverTab[43019]"
	}
//line /usr/local/go/src/net/http/server.go:3090
	// _ = "end of CoverTab[43006]"
}

// ServeTLS accepts incoming connections on the Listener l, creating a
//line /usr/local/go/src/net/http/server.go:3093
// new service goroutine for each. The service goroutines perform TLS
//line /usr/local/go/src/net/http/server.go:3093
// setup and then read requests, calling srv.Handler to reply to them.
//line /usr/local/go/src/net/http/server.go:3093
//
//line /usr/local/go/src/net/http/server.go:3093
// Files containing a certificate and matching private key for the
//line /usr/local/go/src/net/http/server.go:3093
// server must be provided if neither the Server's
//line /usr/local/go/src/net/http/server.go:3093
// TLSConfig.Certificates nor TLSConfig.GetCertificate are populated.
//line /usr/local/go/src/net/http/server.go:3093
// If the certificate is signed by a certificate authority, the
//line /usr/local/go/src/net/http/server.go:3093
// certFile should be the concatenation of the server's certificate,
//line /usr/local/go/src/net/http/server.go:3093
// any intermediates, and the CA's certificate.
//line /usr/local/go/src/net/http/server.go:3093
//
//line /usr/local/go/src/net/http/server.go:3093
// ServeTLS always returns a non-nil error. After Shutdown or Close, the
//line /usr/local/go/src/net/http/server.go:3093
// returned error is ErrServerClosed.
//line /usr/local/go/src/net/http/server.go:3106
func (srv *Server) ServeTLS(l net.Listener, certFile, keyFile string) error {
//line /usr/local/go/src/net/http/server.go:3106
	_go_fuzz_dep_.CoverTab[43039]++

//line /usr/local/go/src/net/http/server.go:3109
	if err := srv.setupHTTP2_ServeTLS(); err != nil {
//line /usr/local/go/src/net/http/server.go:3109
		_go_fuzz_dep_.CoverTab[43043]++
								return err
//line /usr/local/go/src/net/http/server.go:3110
		// _ = "end of CoverTab[43043]"
	} else {
//line /usr/local/go/src/net/http/server.go:3111
		_go_fuzz_dep_.CoverTab[43044]++
//line /usr/local/go/src/net/http/server.go:3111
		// _ = "end of CoverTab[43044]"
//line /usr/local/go/src/net/http/server.go:3111
	}
//line /usr/local/go/src/net/http/server.go:3111
	// _ = "end of CoverTab[43039]"
//line /usr/local/go/src/net/http/server.go:3111
	_go_fuzz_dep_.CoverTab[43040]++

							config := cloneTLSConfig(srv.TLSConfig)
							if !strSliceContains(config.NextProtos, "http/1.1") {
//line /usr/local/go/src/net/http/server.go:3114
		_go_fuzz_dep_.CoverTab[43045]++
								config.NextProtos = append(config.NextProtos, "http/1.1")
//line /usr/local/go/src/net/http/server.go:3115
		// _ = "end of CoverTab[43045]"
	} else {
//line /usr/local/go/src/net/http/server.go:3116
		_go_fuzz_dep_.CoverTab[43046]++
//line /usr/local/go/src/net/http/server.go:3116
		// _ = "end of CoverTab[43046]"
//line /usr/local/go/src/net/http/server.go:3116
	}
//line /usr/local/go/src/net/http/server.go:3116
	// _ = "end of CoverTab[43040]"
//line /usr/local/go/src/net/http/server.go:3116
	_go_fuzz_dep_.CoverTab[43041]++

							configHasCert := len(config.Certificates) > 0 || func() bool {
//line /usr/local/go/src/net/http/server.go:3118
		_go_fuzz_dep_.CoverTab[43047]++
//line /usr/local/go/src/net/http/server.go:3118
		return config.GetCertificate != nil
//line /usr/local/go/src/net/http/server.go:3118
		// _ = "end of CoverTab[43047]"
//line /usr/local/go/src/net/http/server.go:3118
	}()
							if !configHasCert || func() bool {
//line /usr/local/go/src/net/http/server.go:3119
		_go_fuzz_dep_.CoverTab[43048]++
//line /usr/local/go/src/net/http/server.go:3119
		return certFile != ""
//line /usr/local/go/src/net/http/server.go:3119
		// _ = "end of CoverTab[43048]"
//line /usr/local/go/src/net/http/server.go:3119
	}() || func() bool {
//line /usr/local/go/src/net/http/server.go:3119
		_go_fuzz_dep_.CoverTab[43049]++
//line /usr/local/go/src/net/http/server.go:3119
		return keyFile != ""
//line /usr/local/go/src/net/http/server.go:3119
		// _ = "end of CoverTab[43049]"
//line /usr/local/go/src/net/http/server.go:3119
	}() {
//line /usr/local/go/src/net/http/server.go:3119
		_go_fuzz_dep_.CoverTab[43050]++
								var err error
								config.Certificates = make([]tls.Certificate, 1)
								config.Certificates[0], err = tls.LoadX509KeyPair(certFile, keyFile)
								if err != nil {
//line /usr/local/go/src/net/http/server.go:3123
			_go_fuzz_dep_.CoverTab[43051]++
									return err
//line /usr/local/go/src/net/http/server.go:3124
			// _ = "end of CoverTab[43051]"
		} else {
//line /usr/local/go/src/net/http/server.go:3125
			_go_fuzz_dep_.CoverTab[43052]++
//line /usr/local/go/src/net/http/server.go:3125
			// _ = "end of CoverTab[43052]"
//line /usr/local/go/src/net/http/server.go:3125
		}
//line /usr/local/go/src/net/http/server.go:3125
		// _ = "end of CoverTab[43050]"
	} else {
//line /usr/local/go/src/net/http/server.go:3126
		_go_fuzz_dep_.CoverTab[43053]++
//line /usr/local/go/src/net/http/server.go:3126
		// _ = "end of CoverTab[43053]"
//line /usr/local/go/src/net/http/server.go:3126
	}
//line /usr/local/go/src/net/http/server.go:3126
	// _ = "end of CoverTab[43041]"
//line /usr/local/go/src/net/http/server.go:3126
	_go_fuzz_dep_.CoverTab[43042]++

							tlsListener := tls.NewListener(l, config)
							return srv.Serve(tlsListener)
//line /usr/local/go/src/net/http/server.go:3129
	// _ = "end of CoverTab[43042]"
}

// trackListener adds or removes a net.Listener to the set of tracked
//line /usr/local/go/src/net/http/server.go:3132
// listeners.
//line /usr/local/go/src/net/http/server.go:3132
//
//line /usr/local/go/src/net/http/server.go:3132
// We store a pointer to interface in the map set, in case the
//line /usr/local/go/src/net/http/server.go:3132
// net.Listener is not comparable. This is safe because we only call
//line /usr/local/go/src/net/http/server.go:3132
// trackListener via Serve and can track+defer untrack the same
//line /usr/local/go/src/net/http/server.go:3132
// pointer to local variable there. We never need to compare a
//line /usr/local/go/src/net/http/server.go:3132
// Listener from another caller.
//line /usr/local/go/src/net/http/server.go:3132
//
//line /usr/local/go/src/net/http/server.go:3132
// It reports whether the server is still up (not Shutdown or Closed).
//line /usr/local/go/src/net/http/server.go:3142
func (s *Server) trackListener(ln *net.Listener, add bool) bool {
//line /usr/local/go/src/net/http/server.go:3142
	_go_fuzz_dep_.CoverTab[43054]++
							s.mu.Lock()
							defer s.mu.Unlock()
							if s.listeners == nil {
//line /usr/local/go/src/net/http/server.go:3145
		_go_fuzz_dep_.CoverTab[43057]++
								s.listeners = make(map[*net.Listener]struct{})
//line /usr/local/go/src/net/http/server.go:3146
		// _ = "end of CoverTab[43057]"
	} else {
//line /usr/local/go/src/net/http/server.go:3147
		_go_fuzz_dep_.CoverTab[43058]++
//line /usr/local/go/src/net/http/server.go:3147
		// _ = "end of CoverTab[43058]"
//line /usr/local/go/src/net/http/server.go:3147
	}
//line /usr/local/go/src/net/http/server.go:3147
	// _ = "end of CoverTab[43054]"
//line /usr/local/go/src/net/http/server.go:3147
	_go_fuzz_dep_.CoverTab[43055]++
							if add {
//line /usr/local/go/src/net/http/server.go:3148
		_go_fuzz_dep_.CoverTab[43059]++
								if s.shuttingDown() {
//line /usr/local/go/src/net/http/server.go:3149
			_go_fuzz_dep_.CoverTab[43061]++
									return false
//line /usr/local/go/src/net/http/server.go:3150
			// _ = "end of CoverTab[43061]"
		} else {
//line /usr/local/go/src/net/http/server.go:3151
			_go_fuzz_dep_.CoverTab[43062]++
//line /usr/local/go/src/net/http/server.go:3151
			// _ = "end of CoverTab[43062]"
//line /usr/local/go/src/net/http/server.go:3151
		}
//line /usr/local/go/src/net/http/server.go:3151
		// _ = "end of CoverTab[43059]"
//line /usr/local/go/src/net/http/server.go:3151
		_go_fuzz_dep_.CoverTab[43060]++
								s.listeners[ln] = struct{}{}
								s.listenerGroup.Add(1)
//line /usr/local/go/src/net/http/server.go:3153
		// _ = "end of CoverTab[43060]"
	} else {
//line /usr/local/go/src/net/http/server.go:3154
		_go_fuzz_dep_.CoverTab[43063]++
								delete(s.listeners, ln)
								s.listenerGroup.Done()
//line /usr/local/go/src/net/http/server.go:3156
		// _ = "end of CoverTab[43063]"
	}
//line /usr/local/go/src/net/http/server.go:3157
	// _ = "end of CoverTab[43055]"
//line /usr/local/go/src/net/http/server.go:3157
	_go_fuzz_dep_.CoverTab[43056]++
							return true
//line /usr/local/go/src/net/http/server.go:3158
	// _ = "end of CoverTab[43056]"
}

func (s *Server) trackConn(c *conn, add bool) {
//line /usr/local/go/src/net/http/server.go:3161
	_go_fuzz_dep_.CoverTab[43064]++
							s.mu.Lock()
							defer s.mu.Unlock()
							if s.activeConn == nil {
//line /usr/local/go/src/net/http/server.go:3164
		_go_fuzz_dep_.CoverTab[43066]++
								s.activeConn = make(map[*conn]struct{})
//line /usr/local/go/src/net/http/server.go:3165
		// _ = "end of CoverTab[43066]"
	} else {
//line /usr/local/go/src/net/http/server.go:3166
		_go_fuzz_dep_.CoverTab[43067]++
//line /usr/local/go/src/net/http/server.go:3166
		// _ = "end of CoverTab[43067]"
//line /usr/local/go/src/net/http/server.go:3166
	}
//line /usr/local/go/src/net/http/server.go:3166
	// _ = "end of CoverTab[43064]"
//line /usr/local/go/src/net/http/server.go:3166
	_go_fuzz_dep_.CoverTab[43065]++
							if add {
//line /usr/local/go/src/net/http/server.go:3167
		_go_fuzz_dep_.CoverTab[43068]++
								s.activeConn[c] = struct{}{}
//line /usr/local/go/src/net/http/server.go:3168
		// _ = "end of CoverTab[43068]"
	} else {
//line /usr/local/go/src/net/http/server.go:3169
		_go_fuzz_dep_.CoverTab[43069]++
								delete(s.activeConn, c)
//line /usr/local/go/src/net/http/server.go:3170
		// _ = "end of CoverTab[43069]"
	}
//line /usr/local/go/src/net/http/server.go:3171
	// _ = "end of CoverTab[43065]"
}

func (s *Server) idleTimeout() time.Duration {
//line /usr/local/go/src/net/http/server.go:3174
	_go_fuzz_dep_.CoverTab[43070]++
							if s.IdleTimeout != 0 {
//line /usr/local/go/src/net/http/server.go:3175
		_go_fuzz_dep_.CoverTab[43072]++
								return s.IdleTimeout
//line /usr/local/go/src/net/http/server.go:3176
		// _ = "end of CoverTab[43072]"
	} else {
//line /usr/local/go/src/net/http/server.go:3177
		_go_fuzz_dep_.CoverTab[43073]++
//line /usr/local/go/src/net/http/server.go:3177
		// _ = "end of CoverTab[43073]"
//line /usr/local/go/src/net/http/server.go:3177
	}
//line /usr/local/go/src/net/http/server.go:3177
	// _ = "end of CoverTab[43070]"
//line /usr/local/go/src/net/http/server.go:3177
	_go_fuzz_dep_.CoverTab[43071]++
							return s.ReadTimeout
//line /usr/local/go/src/net/http/server.go:3178
	// _ = "end of CoverTab[43071]"
}

func (s *Server) readHeaderTimeout() time.Duration {
//line /usr/local/go/src/net/http/server.go:3181
	_go_fuzz_dep_.CoverTab[43074]++
							if s.ReadHeaderTimeout != 0 {
//line /usr/local/go/src/net/http/server.go:3182
		_go_fuzz_dep_.CoverTab[43076]++
								return s.ReadHeaderTimeout
//line /usr/local/go/src/net/http/server.go:3183
		// _ = "end of CoverTab[43076]"
	} else {
//line /usr/local/go/src/net/http/server.go:3184
		_go_fuzz_dep_.CoverTab[43077]++
//line /usr/local/go/src/net/http/server.go:3184
		// _ = "end of CoverTab[43077]"
//line /usr/local/go/src/net/http/server.go:3184
	}
//line /usr/local/go/src/net/http/server.go:3184
	// _ = "end of CoverTab[43074]"
//line /usr/local/go/src/net/http/server.go:3184
	_go_fuzz_dep_.CoverTab[43075]++
							return s.ReadTimeout
//line /usr/local/go/src/net/http/server.go:3185
	// _ = "end of CoverTab[43075]"
}

func (s *Server) doKeepAlives() bool {
//line /usr/local/go/src/net/http/server.go:3188
	_go_fuzz_dep_.CoverTab[43078]++
							return !s.disableKeepAlives.Load() && func() bool {
//line /usr/local/go/src/net/http/server.go:3189
		_go_fuzz_dep_.CoverTab[43079]++
//line /usr/local/go/src/net/http/server.go:3189
		return !s.shuttingDown()
//line /usr/local/go/src/net/http/server.go:3189
		// _ = "end of CoverTab[43079]"
//line /usr/local/go/src/net/http/server.go:3189
	}()
//line /usr/local/go/src/net/http/server.go:3189
	// _ = "end of CoverTab[43078]"
}

func (s *Server) shuttingDown() bool {
//line /usr/local/go/src/net/http/server.go:3192
	_go_fuzz_dep_.CoverTab[43080]++
							return s.inShutdown.Load()
//line /usr/local/go/src/net/http/server.go:3193
	// _ = "end of CoverTab[43080]"
}

// SetKeepAlivesEnabled controls whether HTTP keep-alives are enabled.
//line /usr/local/go/src/net/http/server.go:3196
// By default, keep-alives are always enabled. Only very
//line /usr/local/go/src/net/http/server.go:3196
// resource-constrained environments or servers in the process of
//line /usr/local/go/src/net/http/server.go:3196
// shutting down should disable them.
//line /usr/local/go/src/net/http/server.go:3200
func (srv *Server) SetKeepAlivesEnabled(v bool) {
//line /usr/local/go/src/net/http/server.go:3200
	_go_fuzz_dep_.CoverTab[43081]++
							if v {
//line /usr/local/go/src/net/http/server.go:3201
		_go_fuzz_dep_.CoverTab[43083]++
								srv.disableKeepAlives.Store(false)
								return
//line /usr/local/go/src/net/http/server.go:3203
		// _ = "end of CoverTab[43083]"
	} else {
//line /usr/local/go/src/net/http/server.go:3204
		_go_fuzz_dep_.CoverTab[43084]++
//line /usr/local/go/src/net/http/server.go:3204
		// _ = "end of CoverTab[43084]"
//line /usr/local/go/src/net/http/server.go:3204
	}
//line /usr/local/go/src/net/http/server.go:3204
	// _ = "end of CoverTab[43081]"
//line /usr/local/go/src/net/http/server.go:3204
	_go_fuzz_dep_.CoverTab[43082]++
							srv.disableKeepAlives.Store(true)

//line /usr/local/go/src/net/http/server.go:3208
	srv.closeIdleConns()
//line /usr/local/go/src/net/http/server.go:3208
	// _ = "end of CoverTab[43082]"

//line /usr/local/go/src/net/http/server.go:3211
}

func (s *Server) logf(format string, args ...any) {
//line /usr/local/go/src/net/http/server.go:3213
	_go_fuzz_dep_.CoverTab[43085]++
							if s.ErrorLog != nil {
//line /usr/local/go/src/net/http/server.go:3214
		_go_fuzz_dep_.CoverTab[43086]++
								s.ErrorLog.Printf(format, args...)
//line /usr/local/go/src/net/http/server.go:3215
		// _ = "end of CoverTab[43086]"
	} else {
//line /usr/local/go/src/net/http/server.go:3216
		_go_fuzz_dep_.CoverTab[43087]++
								log.Printf(format, args...)
//line /usr/local/go/src/net/http/server.go:3217
		// _ = "end of CoverTab[43087]"
	}
//line /usr/local/go/src/net/http/server.go:3218
	// _ = "end of CoverTab[43085]"
}

// logf prints to the ErrorLog of the *Server associated with request r
//line /usr/local/go/src/net/http/server.go:3221
// via ServerContextKey. If there's no associated server, or if ErrorLog
//line /usr/local/go/src/net/http/server.go:3221
// is nil, logging is done via the log package's standard logger.
//line /usr/local/go/src/net/http/server.go:3224
func logf(r *Request, format string, args ...any) {
//line /usr/local/go/src/net/http/server.go:3224
	_go_fuzz_dep_.CoverTab[43088]++
							s, _ := r.Context().Value(ServerContextKey).(*Server)
							if s != nil && func() bool {
//line /usr/local/go/src/net/http/server.go:3226
		_go_fuzz_dep_.CoverTab[43089]++
//line /usr/local/go/src/net/http/server.go:3226
		return s.ErrorLog != nil
//line /usr/local/go/src/net/http/server.go:3226
		// _ = "end of CoverTab[43089]"
//line /usr/local/go/src/net/http/server.go:3226
	}() {
//line /usr/local/go/src/net/http/server.go:3226
		_go_fuzz_dep_.CoverTab[43090]++
								s.ErrorLog.Printf(format, args...)
//line /usr/local/go/src/net/http/server.go:3227
		// _ = "end of CoverTab[43090]"
	} else {
//line /usr/local/go/src/net/http/server.go:3228
		_go_fuzz_dep_.CoverTab[43091]++
								log.Printf(format, args...)
//line /usr/local/go/src/net/http/server.go:3229
		// _ = "end of CoverTab[43091]"
	}
//line /usr/local/go/src/net/http/server.go:3230
	// _ = "end of CoverTab[43088]"
}

// ListenAndServe listens on the TCP network address addr and then calls
//line /usr/local/go/src/net/http/server.go:3233
// Serve with handler to handle requests on incoming connections.
//line /usr/local/go/src/net/http/server.go:3233
// Accepted connections are configured to enable TCP keep-alives.
//line /usr/local/go/src/net/http/server.go:3233
//
//line /usr/local/go/src/net/http/server.go:3233
// The handler is typically nil, in which case the DefaultServeMux is used.
//line /usr/local/go/src/net/http/server.go:3233
//
//line /usr/local/go/src/net/http/server.go:3233
// ListenAndServe always returns a non-nil error.
//line /usr/local/go/src/net/http/server.go:3240
func ListenAndServe(addr string, handler Handler) error {
//line /usr/local/go/src/net/http/server.go:3240
	_go_fuzz_dep_.CoverTab[43092]++
							server := &Server{Addr: addr, Handler: handler}
							return server.ListenAndServe()
//line /usr/local/go/src/net/http/server.go:3242
	// _ = "end of CoverTab[43092]"
}

// ListenAndServeTLS acts identically to ListenAndServe, except that it
//line /usr/local/go/src/net/http/server.go:3245
// expects HTTPS connections. Additionally, files containing a certificate and
//line /usr/local/go/src/net/http/server.go:3245
// matching private key for the server must be provided. If the certificate
//line /usr/local/go/src/net/http/server.go:3245
// is signed by a certificate authority, the certFile should be the concatenation
//line /usr/local/go/src/net/http/server.go:3245
// of the server's certificate, any intermediates, and the CA's certificate.
//line /usr/local/go/src/net/http/server.go:3250
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error {
//line /usr/local/go/src/net/http/server.go:3250
	_go_fuzz_dep_.CoverTab[43093]++
							server := &Server{Addr: addr, Handler: handler}
							return server.ListenAndServeTLS(certFile, keyFile)
//line /usr/local/go/src/net/http/server.go:3252
	// _ = "end of CoverTab[43093]"
}

// ListenAndServeTLS listens on the TCP network address srv.Addr and
//line /usr/local/go/src/net/http/server.go:3255
// then calls ServeTLS to handle requests on incoming TLS connections.
//line /usr/local/go/src/net/http/server.go:3255
// Accepted connections are configured to enable TCP keep-alives.
//line /usr/local/go/src/net/http/server.go:3255
//
//line /usr/local/go/src/net/http/server.go:3255
// Filenames containing a certificate and matching private key for the
//line /usr/local/go/src/net/http/server.go:3255
// server must be provided if neither the Server's TLSConfig.Certificates
//line /usr/local/go/src/net/http/server.go:3255
// nor TLSConfig.GetCertificate are populated. If the certificate is
//line /usr/local/go/src/net/http/server.go:3255
// signed by a certificate authority, the certFile should be the
//line /usr/local/go/src/net/http/server.go:3255
// concatenation of the server's certificate, any intermediates, and
//line /usr/local/go/src/net/http/server.go:3255
// the CA's certificate.
//line /usr/local/go/src/net/http/server.go:3255
//
//line /usr/local/go/src/net/http/server.go:3255
// If srv.Addr is blank, ":https" is used.
//line /usr/local/go/src/net/http/server.go:3255
//
//line /usr/local/go/src/net/http/server.go:3255
// ListenAndServeTLS always returns a non-nil error. After Shutdown or
//line /usr/local/go/src/net/http/server.go:3255
// Close, the returned error is ErrServerClosed.
//line /usr/local/go/src/net/http/server.go:3270
func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error {
//line /usr/local/go/src/net/http/server.go:3270
	_go_fuzz_dep_.CoverTab[43094]++
							if srv.shuttingDown() {
//line /usr/local/go/src/net/http/server.go:3271
		_go_fuzz_dep_.CoverTab[43098]++
								return ErrServerClosed
//line /usr/local/go/src/net/http/server.go:3272
		// _ = "end of CoverTab[43098]"
	} else {
//line /usr/local/go/src/net/http/server.go:3273
		_go_fuzz_dep_.CoverTab[43099]++
//line /usr/local/go/src/net/http/server.go:3273
		// _ = "end of CoverTab[43099]"
//line /usr/local/go/src/net/http/server.go:3273
	}
//line /usr/local/go/src/net/http/server.go:3273
	// _ = "end of CoverTab[43094]"
//line /usr/local/go/src/net/http/server.go:3273
	_go_fuzz_dep_.CoverTab[43095]++
							addr := srv.Addr
							if addr == "" {
//line /usr/local/go/src/net/http/server.go:3275
		_go_fuzz_dep_.CoverTab[43100]++
								addr = ":https"
//line /usr/local/go/src/net/http/server.go:3276
		// _ = "end of CoverTab[43100]"
	} else {
//line /usr/local/go/src/net/http/server.go:3277
		_go_fuzz_dep_.CoverTab[43101]++
//line /usr/local/go/src/net/http/server.go:3277
		// _ = "end of CoverTab[43101]"
//line /usr/local/go/src/net/http/server.go:3277
	}
//line /usr/local/go/src/net/http/server.go:3277
	// _ = "end of CoverTab[43095]"
//line /usr/local/go/src/net/http/server.go:3277
	_go_fuzz_dep_.CoverTab[43096]++

							ln, err := net.Listen("tcp", addr)
							if err != nil {
//line /usr/local/go/src/net/http/server.go:3280
		_go_fuzz_dep_.CoverTab[43102]++
								return err
//line /usr/local/go/src/net/http/server.go:3281
		// _ = "end of CoverTab[43102]"
	} else {
//line /usr/local/go/src/net/http/server.go:3282
		_go_fuzz_dep_.CoverTab[43103]++
//line /usr/local/go/src/net/http/server.go:3282
		// _ = "end of CoverTab[43103]"
//line /usr/local/go/src/net/http/server.go:3282
	}
//line /usr/local/go/src/net/http/server.go:3282
	// _ = "end of CoverTab[43096]"
//line /usr/local/go/src/net/http/server.go:3282
	_go_fuzz_dep_.CoverTab[43097]++

							defer ln.Close()

							return srv.ServeTLS(ln, certFile, keyFile)
//line /usr/local/go/src/net/http/server.go:3286
	// _ = "end of CoverTab[43097]"
}

// setupHTTP2_ServeTLS conditionally configures HTTP/2 on
//line /usr/local/go/src/net/http/server.go:3289
// srv and reports whether there was an error setting it up. If it is
//line /usr/local/go/src/net/http/server.go:3289
// not configured for policy reasons, nil is returned.
//line /usr/local/go/src/net/http/server.go:3292
func (srv *Server) setupHTTP2_ServeTLS() error {
//line /usr/local/go/src/net/http/server.go:3292
	_go_fuzz_dep_.CoverTab[43104]++
							srv.nextProtoOnce.Do(srv.onceSetNextProtoDefaults)
							return srv.nextProtoErr
//line /usr/local/go/src/net/http/server.go:3294
	// _ = "end of CoverTab[43104]"
}

// setupHTTP2_Serve is called from (*Server).Serve and conditionally
//line /usr/local/go/src/net/http/server.go:3297
// configures HTTP/2 on srv using a more conservative policy than
//line /usr/local/go/src/net/http/server.go:3297
// setupHTTP2_ServeTLS because Serve is called after tls.Listen,
//line /usr/local/go/src/net/http/server.go:3297
// and may be called concurrently. See shouldConfigureHTTP2ForServe.
//line /usr/local/go/src/net/http/server.go:3297
//
//line /usr/local/go/src/net/http/server.go:3297
// The tests named TestTransportAutomaticHTTP2* and
//line /usr/local/go/src/net/http/server.go:3297
// TestConcurrentServerServe in server_test.go demonstrate some
//line /usr/local/go/src/net/http/server.go:3297
// of the supported use cases and motivations.
//line /usr/local/go/src/net/http/server.go:3305
func (srv *Server) setupHTTP2_Serve() error {
//line /usr/local/go/src/net/http/server.go:3305
	_go_fuzz_dep_.CoverTab[43105]++
							srv.nextProtoOnce.Do(srv.onceSetNextProtoDefaults_Serve)
							return srv.nextProtoErr
//line /usr/local/go/src/net/http/server.go:3307
	// _ = "end of CoverTab[43105]"
}

func (srv *Server) onceSetNextProtoDefaults_Serve() {
//line /usr/local/go/src/net/http/server.go:3310
	_go_fuzz_dep_.CoverTab[43106]++
							if srv.shouldConfigureHTTP2ForServe() {
//line /usr/local/go/src/net/http/server.go:3311
		_go_fuzz_dep_.CoverTab[43107]++
								srv.onceSetNextProtoDefaults()
//line /usr/local/go/src/net/http/server.go:3312
		// _ = "end of CoverTab[43107]"
	} else {
//line /usr/local/go/src/net/http/server.go:3313
		_go_fuzz_dep_.CoverTab[43108]++
//line /usr/local/go/src/net/http/server.go:3313
		// _ = "end of CoverTab[43108]"
//line /usr/local/go/src/net/http/server.go:3313
	}
//line /usr/local/go/src/net/http/server.go:3313
	// _ = "end of CoverTab[43106]"
}

var http2server = godebug.New("http2server")

// onceSetNextProtoDefaults configures HTTP/2, if the user hasn't
//line /usr/local/go/src/net/http/server.go:3318
// configured otherwise. (by setting srv.TLSNextProto non-nil)
//line /usr/local/go/src/net/http/server.go:3318
// It must only be called via srv.nextProtoOnce (use srv.setupHTTP2_*).
//line /usr/local/go/src/net/http/server.go:3321
func (srv *Server) onceSetNextProtoDefaults() {
//line /usr/local/go/src/net/http/server.go:3321
	_go_fuzz_dep_.CoverTab[43109]++
							if omitBundledHTTP2 || func() bool {
//line /usr/local/go/src/net/http/server.go:3322
		_go_fuzz_dep_.CoverTab[43111]++
//line /usr/local/go/src/net/http/server.go:3322
		return http2server.Value() == "0"
//line /usr/local/go/src/net/http/server.go:3322
		// _ = "end of CoverTab[43111]"
//line /usr/local/go/src/net/http/server.go:3322
	}() {
//line /usr/local/go/src/net/http/server.go:3322
		_go_fuzz_dep_.CoverTab[43112]++
								return
//line /usr/local/go/src/net/http/server.go:3323
		// _ = "end of CoverTab[43112]"
	} else {
//line /usr/local/go/src/net/http/server.go:3324
		_go_fuzz_dep_.CoverTab[43113]++
//line /usr/local/go/src/net/http/server.go:3324
		// _ = "end of CoverTab[43113]"
//line /usr/local/go/src/net/http/server.go:3324
	}
//line /usr/local/go/src/net/http/server.go:3324
	// _ = "end of CoverTab[43109]"
//line /usr/local/go/src/net/http/server.go:3324
	_go_fuzz_dep_.CoverTab[43110]++

//line /usr/local/go/src/net/http/server.go:3327
	if srv.TLSNextProto == nil {
//line /usr/local/go/src/net/http/server.go:3327
		_go_fuzz_dep_.CoverTab[43114]++
								conf := &http2Server{
			NewWriteScheduler: func() http2WriteScheduler {
//line /usr/local/go/src/net/http/server.go:3329
				_go_fuzz_dep_.CoverTab[43116]++
//line /usr/local/go/src/net/http/server.go:3329
				return http2NewPriorityWriteScheduler(nil)
//line /usr/local/go/src/net/http/server.go:3329
				// _ = "end of CoverTab[43116]"
//line /usr/local/go/src/net/http/server.go:3329
			},
		}
//line /usr/local/go/src/net/http/server.go:3330
		// _ = "end of CoverTab[43114]"
//line /usr/local/go/src/net/http/server.go:3330
		_go_fuzz_dep_.CoverTab[43115]++
								srv.nextProtoErr = http2ConfigureServer(srv, conf)
//line /usr/local/go/src/net/http/server.go:3331
		// _ = "end of CoverTab[43115]"
	} else {
//line /usr/local/go/src/net/http/server.go:3332
		_go_fuzz_dep_.CoverTab[43117]++
//line /usr/local/go/src/net/http/server.go:3332
		// _ = "end of CoverTab[43117]"
//line /usr/local/go/src/net/http/server.go:3332
	}
//line /usr/local/go/src/net/http/server.go:3332
	// _ = "end of CoverTab[43110]"
}

// TimeoutHandler returns a Handler that runs h with the given time limit.
//line /usr/local/go/src/net/http/server.go:3335
//
//line /usr/local/go/src/net/http/server.go:3335
// The new Handler calls h.ServeHTTP to handle each request, but if a
//line /usr/local/go/src/net/http/server.go:3335
// call runs for longer than its time limit, the handler responds with
//line /usr/local/go/src/net/http/server.go:3335
// a 503 Service Unavailable error and the given message in its body.
//line /usr/local/go/src/net/http/server.go:3335
// (If msg is empty, a suitable default message will be sent.)
//line /usr/local/go/src/net/http/server.go:3335
// After such a timeout, writes by h to its ResponseWriter will return
//line /usr/local/go/src/net/http/server.go:3335
// ErrHandlerTimeout.
//line /usr/local/go/src/net/http/server.go:3335
//
//line /usr/local/go/src/net/http/server.go:3335
// TimeoutHandler supports the Pusher interface but does not support
//line /usr/local/go/src/net/http/server.go:3335
// the Hijacker or Flusher interfaces.
//line /usr/local/go/src/net/http/server.go:3346
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler {
//line /usr/local/go/src/net/http/server.go:3346
	_go_fuzz_dep_.CoverTab[43118]++
							return &timeoutHandler{
		handler:	h,
		body:		msg,
		dt:		dt,
	}
//line /usr/local/go/src/net/http/server.go:3351
	// _ = "end of CoverTab[43118]"
}

// ErrHandlerTimeout is returned on ResponseWriter Write calls
//line /usr/local/go/src/net/http/server.go:3354
// in handlers which have timed out.
//line /usr/local/go/src/net/http/server.go:3356
var ErrHandlerTimeout = errors.New("http: Handler timeout")

type timeoutHandler struct {
	handler	Handler
	body	string
	dt	time.Duration

	// When set, no context will be created and this context will
	// be used instead.
	testContext	context.Context
}

func (h *timeoutHandler) errorBody() string {
//line /usr/local/go/src/net/http/server.go:3368
	_go_fuzz_dep_.CoverTab[43119]++
							if h.body != "" {
//line /usr/local/go/src/net/http/server.go:3369
		_go_fuzz_dep_.CoverTab[43121]++
								return h.body
//line /usr/local/go/src/net/http/server.go:3370
		// _ = "end of CoverTab[43121]"
	} else {
//line /usr/local/go/src/net/http/server.go:3371
		_go_fuzz_dep_.CoverTab[43122]++
//line /usr/local/go/src/net/http/server.go:3371
		// _ = "end of CoverTab[43122]"
//line /usr/local/go/src/net/http/server.go:3371
	}
//line /usr/local/go/src/net/http/server.go:3371
	// _ = "end of CoverTab[43119]"
//line /usr/local/go/src/net/http/server.go:3371
	_go_fuzz_dep_.CoverTab[43120]++
							return "<html><head><title>Timeout</title></head><body><h1>Timeout</h1></body></html>"
//line /usr/local/go/src/net/http/server.go:3372
	// _ = "end of CoverTab[43120]"
}

func (h *timeoutHandler) ServeHTTP(w ResponseWriter, r *Request) {
//line /usr/local/go/src/net/http/server.go:3375
	_go_fuzz_dep_.CoverTab[43123]++
							ctx := h.testContext
							if ctx == nil {
//line /usr/local/go/src/net/http/server.go:3377
		_go_fuzz_dep_.CoverTab[43126]++
								var cancelCtx context.CancelFunc
								ctx, cancelCtx = context.WithTimeout(r.Context(), h.dt)
								defer cancelCtx()
//line /usr/local/go/src/net/http/server.go:3380
		// _ = "end of CoverTab[43126]"
	} else {
//line /usr/local/go/src/net/http/server.go:3381
		_go_fuzz_dep_.CoverTab[43127]++
//line /usr/local/go/src/net/http/server.go:3381
		// _ = "end of CoverTab[43127]"
//line /usr/local/go/src/net/http/server.go:3381
	}
//line /usr/local/go/src/net/http/server.go:3381
	// _ = "end of CoverTab[43123]"
//line /usr/local/go/src/net/http/server.go:3381
	_go_fuzz_dep_.CoverTab[43124]++
							r = r.WithContext(ctx)
							done := make(chan struct{})
							tw := &timeoutWriter{
		w:	w,
		h:	make(Header),
		req:	r,
	}
							panicChan := make(chan any, 1)
//line /usr/local/go/src/net/http/server.go:3389
	_curRoutineNum34_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/server.go:3389
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum34_)
							go func() {
//line /usr/local/go/src/net/http/server.go:3390
		_go_fuzz_dep_.CoverTab[43128]++
//line /usr/local/go/src/net/http/server.go:3390
		defer func() {
//line /usr/local/go/src/net/http/server.go:3390
			_go_fuzz_dep_.CoverTab[43130]++
//line /usr/local/go/src/net/http/server.go:3390
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum34_)
//line /usr/local/go/src/net/http/server.go:3390
			// _ = "end of CoverTab[43130]"
//line /usr/local/go/src/net/http/server.go:3390
		}()
								defer func() {
//line /usr/local/go/src/net/http/server.go:3391
			_go_fuzz_dep_.CoverTab[43131]++
									if p := recover(); p != nil {
//line /usr/local/go/src/net/http/server.go:3392
				_go_fuzz_dep_.CoverTab[43132]++
										panicChan <- p
//line /usr/local/go/src/net/http/server.go:3393
				// _ = "end of CoverTab[43132]"
			} else {
//line /usr/local/go/src/net/http/server.go:3394
				_go_fuzz_dep_.CoverTab[43133]++
//line /usr/local/go/src/net/http/server.go:3394
				// _ = "end of CoverTab[43133]"
//line /usr/local/go/src/net/http/server.go:3394
			}
//line /usr/local/go/src/net/http/server.go:3394
			// _ = "end of CoverTab[43131]"
		}()
//line /usr/local/go/src/net/http/server.go:3395
		// _ = "end of CoverTab[43128]"
//line /usr/local/go/src/net/http/server.go:3395
		_go_fuzz_dep_.CoverTab[43129]++
								h.handler.ServeHTTP(tw, r)
								close(done)
//line /usr/local/go/src/net/http/server.go:3397
		// _ = "end of CoverTab[43129]"
	}()
//line /usr/local/go/src/net/http/server.go:3398
	// _ = "end of CoverTab[43124]"
//line /usr/local/go/src/net/http/server.go:3398
	_go_fuzz_dep_.CoverTab[43125]++
							select {
	case p := <-panicChan:
//line /usr/local/go/src/net/http/server.go:3400
		_go_fuzz_dep_.CoverTab[43134]++
								panic(p)
//line /usr/local/go/src/net/http/server.go:3401
		// _ = "end of CoverTab[43134]"
	case <-done:
//line /usr/local/go/src/net/http/server.go:3402
		_go_fuzz_dep_.CoverTab[43135]++
								tw.mu.Lock()
								defer tw.mu.Unlock()
								dst := w.Header()
								for k, vv := range tw.h {
//line /usr/local/go/src/net/http/server.go:3406
			_go_fuzz_dep_.CoverTab[43139]++
									dst[k] = vv
//line /usr/local/go/src/net/http/server.go:3407
			// _ = "end of CoverTab[43139]"
		}
//line /usr/local/go/src/net/http/server.go:3408
		// _ = "end of CoverTab[43135]"
//line /usr/local/go/src/net/http/server.go:3408
		_go_fuzz_dep_.CoverTab[43136]++
								if !tw.wroteHeader {
//line /usr/local/go/src/net/http/server.go:3409
			_go_fuzz_dep_.CoverTab[43140]++
									tw.code = StatusOK
//line /usr/local/go/src/net/http/server.go:3410
			// _ = "end of CoverTab[43140]"
		} else {
//line /usr/local/go/src/net/http/server.go:3411
			_go_fuzz_dep_.CoverTab[43141]++
//line /usr/local/go/src/net/http/server.go:3411
			// _ = "end of CoverTab[43141]"
//line /usr/local/go/src/net/http/server.go:3411
		}
//line /usr/local/go/src/net/http/server.go:3411
		// _ = "end of CoverTab[43136]"
//line /usr/local/go/src/net/http/server.go:3411
		_go_fuzz_dep_.CoverTab[43137]++
								w.WriteHeader(tw.code)
								w.Write(tw.wbuf.Bytes())
//line /usr/local/go/src/net/http/server.go:3413
		// _ = "end of CoverTab[43137]"
	case <-ctx.Done():
//line /usr/local/go/src/net/http/server.go:3414
		_go_fuzz_dep_.CoverTab[43138]++
								tw.mu.Lock()
								defer tw.mu.Unlock()
								switch err := ctx.Err(); err {
		case context.DeadlineExceeded:
//line /usr/local/go/src/net/http/server.go:3418
			_go_fuzz_dep_.CoverTab[43142]++
									w.WriteHeader(StatusServiceUnavailable)
									io.WriteString(w, h.errorBody())
									tw.err = ErrHandlerTimeout
//line /usr/local/go/src/net/http/server.go:3421
			// _ = "end of CoverTab[43142]"
		default:
//line /usr/local/go/src/net/http/server.go:3422
			_go_fuzz_dep_.CoverTab[43143]++
									w.WriteHeader(StatusServiceUnavailable)
									tw.err = err
//line /usr/local/go/src/net/http/server.go:3424
			// _ = "end of CoverTab[43143]"
		}
//line /usr/local/go/src/net/http/server.go:3425
		// _ = "end of CoverTab[43138]"
	}
//line /usr/local/go/src/net/http/server.go:3426
	// _ = "end of CoverTab[43125]"
}

type timeoutWriter struct {
	w	ResponseWriter
	h	Header
	wbuf	bytes.Buffer
	req	*Request

	mu		sync.Mutex
	err		error
	wroteHeader	bool
	code		int
}

var _ Pusher = (*timeoutWriter)(nil)

// Push implements the Pusher interface.
func (tw *timeoutWriter) Push(target string, opts *PushOptions) error {
//line /usr/local/go/src/net/http/server.go:3444
	_go_fuzz_dep_.CoverTab[43144]++
							if pusher, ok := tw.w.(Pusher); ok {
//line /usr/local/go/src/net/http/server.go:3445
		_go_fuzz_dep_.CoverTab[43146]++
								return pusher.Push(target, opts)
//line /usr/local/go/src/net/http/server.go:3446
		// _ = "end of CoverTab[43146]"
	} else {
//line /usr/local/go/src/net/http/server.go:3447
		_go_fuzz_dep_.CoverTab[43147]++
//line /usr/local/go/src/net/http/server.go:3447
		// _ = "end of CoverTab[43147]"
//line /usr/local/go/src/net/http/server.go:3447
	}
//line /usr/local/go/src/net/http/server.go:3447
	// _ = "end of CoverTab[43144]"
//line /usr/local/go/src/net/http/server.go:3447
	_go_fuzz_dep_.CoverTab[43145]++
							return ErrNotSupported
//line /usr/local/go/src/net/http/server.go:3448
	// _ = "end of CoverTab[43145]"
}

func (tw *timeoutWriter) Header() Header {
//line /usr/local/go/src/net/http/server.go:3451
	_go_fuzz_dep_.CoverTab[43148]++
//line /usr/local/go/src/net/http/server.go:3451
	return tw.h
//line /usr/local/go/src/net/http/server.go:3451
	// _ = "end of CoverTab[43148]"
//line /usr/local/go/src/net/http/server.go:3451
}

func (tw *timeoutWriter) Write(p []byte) (int, error) {
//line /usr/local/go/src/net/http/server.go:3453
	_go_fuzz_dep_.CoverTab[43149]++
							tw.mu.Lock()
							defer tw.mu.Unlock()
							if tw.err != nil {
//line /usr/local/go/src/net/http/server.go:3456
		_go_fuzz_dep_.CoverTab[43152]++
								return 0, tw.err
//line /usr/local/go/src/net/http/server.go:3457
		// _ = "end of CoverTab[43152]"
	} else {
//line /usr/local/go/src/net/http/server.go:3458
		_go_fuzz_dep_.CoverTab[43153]++
//line /usr/local/go/src/net/http/server.go:3458
		// _ = "end of CoverTab[43153]"
//line /usr/local/go/src/net/http/server.go:3458
	}
//line /usr/local/go/src/net/http/server.go:3458
	// _ = "end of CoverTab[43149]"
//line /usr/local/go/src/net/http/server.go:3458
	_go_fuzz_dep_.CoverTab[43150]++
							if !tw.wroteHeader {
//line /usr/local/go/src/net/http/server.go:3459
		_go_fuzz_dep_.CoverTab[43154]++
								tw.writeHeaderLocked(StatusOK)
//line /usr/local/go/src/net/http/server.go:3460
		// _ = "end of CoverTab[43154]"
	} else {
//line /usr/local/go/src/net/http/server.go:3461
		_go_fuzz_dep_.CoverTab[43155]++
//line /usr/local/go/src/net/http/server.go:3461
		// _ = "end of CoverTab[43155]"
//line /usr/local/go/src/net/http/server.go:3461
	}
//line /usr/local/go/src/net/http/server.go:3461
	// _ = "end of CoverTab[43150]"
//line /usr/local/go/src/net/http/server.go:3461
	_go_fuzz_dep_.CoverTab[43151]++
							return tw.wbuf.Write(p)
//line /usr/local/go/src/net/http/server.go:3462
	// _ = "end of CoverTab[43151]"
}

func (tw *timeoutWriter) writeHeaderLocked(code int) {
//line /usr/local/go/src/net/http/server.go:3465
	_go_fuzz_dep_.CoverTab[43156]++
							checkWriteHeaderCode(code)

							switch {
	case tw.err != nil:
//line /usr/local/go/src/net/http/server.go:3469
		_go_fuzz_dep_.CoverTab[43157]++
								return
//line /usr/local/go/src/net/http/server.go:3470
		// _ = "end of CoverTab[43157]"
	case tw.wroteHeader:
//line /usr/local/go/src/net/http/server.go:3471
		_go_fuzz_dep_.CoverTab[43158]++
								if tw.req != nil {
//line /usr/local/go/src/net/http/server.go:3472
			_go_fuzz_dep_.CoverTab[43160]++
									caller := relevantCaller()
									logf(tw.req, "http: superfluous response.WriteHeader call from %s (%s:%d)", caller.Function, path.Base(caller.File), caller.Line)
//line /usr/local/go/src/net/http/server.go:3474
			// _ = "end of CoverTab[43160]"
		} else {
//line /usr/local/go/src/net/http/server.go:3475
			_go_fuzz_dep_.CoverTab[43161]++
//line /usr/local/go/src/net/http/server.go:3475
			// _ = "end of CoverTab[43161]"
//line /usr/local/go/src/net/http/server.go:3475
		}
//line /usr/local/go/src/net/http/server.go:3475
		// _ = "end of CoverTab[43158]"
	default:
//line /usr/local/go/src/net/http/server.go:3476
		_go_fuzz_dep_.CoverTab[43159]++
								tw.wroteHeader = true
								tw.code = code
//line /usr/local/go/src/net/http/server.go:3478
		// _ = "end of CoverTab[43159]"
	}
//line /usr/local/go/src/net/http/server.go:3479
	// _ = "end of CoverTab[43156]"
}

func (tw *timeoutWriter) WriteHeader(code int) {
//line /usr/local/go/src/net/http/server.go:3482
	_go_fuzz_dep_.CoverTab[43162]++
							tw.mu.Lock()
							defer tw.mu.Unlock()
							tw.writeHeaderLocked(code)
//line /usr/local/go/src/net/http/server.go:3485
	// _ = "end of CoverTab[43162]"
}

// onceCloseListener wraps a net.Listener, protecting it from
//line /usr/local/go/src/net/http/server.go:3488
// multiple Close calls.
//line /usr/local/go/src/net/http/server.go:3490
type onceCloseListener struct {
	net.Listener
	once		sync.Once
	closeErr	error
}

func (oc *onceCloseListener) Close() error {
//line /usr/local/go/src/net/http/server.go:3496
	_go_fuzz_dep_.CoverTab[43163]++
							oc.once.Do(oc.close)
							return oc.closeErr
//line /usr/local/go/src/net/http/server.go:3498
	// _ = "end of CoverTab[43163]"
}

func (oc *onceCloseListener) close() {
//line /usr/local/go/src/net/http/server.go:3501
	_go_fuzz_dep_.CoverTab[43164]++
//line /usr/local/go/src/net/http/server.go:3501
	oc.closeErr = oc.Listener.Close()
//line /usr/local/go/src/net/http/server.go:3501
	// _ = "end of CoverTab[43164]"
//line /usr/local/go/src/net/http/server.go:3501
}

// globalOptionsHandler responds to "OPTIONS *" requests.
type globalOptionsHandler struct{}

func (globalOptionsHandler) ServeHTTP(w ResponseWriter, r *Request) {
//line /usr/local/go/src/net/http/server.go:3506
	_go_fuzz_dep_.CoverTab[43165]++
							w.Header().Set("Content-Length", "0")
							if r.ContentLength != 0 {
//line /usr/local/go/src/net/http/server.go:3508
		_go_fuzz_dep_.CoverTab[43166]++

//line /usr/local/go/src/net/http/server.go:3514
		mb := MaxBytesReader(w, r.Body, 4<<10)
								io.Copy(io.Discard, mb)
//line /usr/local/go/src/net/http/server.go:3515
		// _ = "end of CoverTab[43166]"
	} else {
//line /usr/local/go/src/net/http/server.go:3516
		_go_fuzz_dep_.CoverTab[43167]++
//line /usr/local/go/src/net/http/server.go:3516
		// _ = "end of CoverTab[43167]"
//line /usr/local/go/src/net/http/server.go:3516
	}
//line /usr/local/go/src/net/http/server.go:3516
	// _ = "end of CoverTab[43165]"
}

// initALPNRequest is an HTTP handler that initializes certain
//line /usr/local/go/src/net/http/server.go:3519
// uninitialized fields in its *Request. Such partially-initialized
//line /usr/local/go/src/net/http/server.go:3519
// Requests come from ALPN protocol handlers.
//line /usr/local/go/src/net/http/server.go:3522
type initALPNRequest struct {
	ctx	context.Context
	c	*tls.Conn
	h	serverHandler
}

// BaseContext is an exported but unadvertised http.Handler method
//line /usr/local/go/src/net/http/server.go:3528
// recognized by x/net/http2 to pass down a context; the TLSNextProto
//line /usr/local/go/src/net/http/server.go:3528
// API predates context support so we shoehorn through the only
//line /usr/local/go/src/net/http/server.go:3528
// interface we have available.
//line /usr/local/go/src/net/http/server.go:3532
func (h initALPNRequest) BaseContext() context.Context {
//line /usr/local/go/src/net/http/server.go:3532
	_go_fuzz_dep_.CoverTab[43168]++
//line /usr/local/go/src/net/http/server.go:3532
	return h.ctx
//line /usr/local/go/src/net/http/server.go:3532
	// _ = "end of CoverTab[43168]"
//line /usr/local/go/src/net/http/server.go:3532
}

func (h initALPNRequest) ServeHTTP(rw ResponseWriter, req *Request) {
//line /usr/local/go/src/net/http/server.go:3534
	_go_fuzz_dep_.CoverTab[43169]++
							if req.TLS == nil {
//line /usr/local/go/src/net/http/server.go:3535
		_go_fuzz_dep_.CoverTab[43173]++
								req.TLS = &tls.ConnectionState{}
								*req.TLS = h.c.ConnectionState()
//line /usr/local/go/src/net/http/server.go:3537
		// _ = "end of CoverTab[43173]"
	} else {
//line /usr/local/go/src/net/http/server.go:3538
		_go_fuzz_dep_.CoverTab[43174]++
//line /usr/local/go/src/net/http/server.go:3538
		// _ = "end of CoverTab[43174]"
//line /usr/local/go/src/net/http/server.go:3538
	}
//line /usr/local/go/src/net/http/server.go:3538
	// _ = "end of CoverTab[43169]"
//line /usr/local/go/src/net/http/server.go:3538
	_go_fuzz_dep_.CoverTab[43170]++
							if req.Body == nil {
//line /usr/local/go/src/net/http/server.go:3539
		_go_fuzz_dep_.CoverTab[43175]++
								req.Body = NoBody
//line /usr/local/go/src/net/http/server.go:3540
		// _ = "end of CoverTab[43175]"
	} else {
//line /usr/local/go/src/net/http/server.go:3541
		_go_fuzz_dep_.CoverTab[43176]++
//line /usr/local/go/src/net/http/server.go:3541
		// _ = "end of CoverTab[43176]"
//line /usr/local/go/src/net/http/server.go:3541
	}
//line /usr/local/go/src/net/http/server.go:3541
	// _ = "end of CoverTab[43170]"
//line /usr/local/go/src/net/http/server.go:3541
	_go_fuzz_dep_.CoverTab[43171]++
							if req.RemoteAddr == "" {
//line /usr/local/go/src/net/http/server.go:3542
		_go_fuzz_dep_.CoverTab[43177]++
								req.RemoteAddr = h.c.RemoteAddr().String()
//line /usr/local/go/src/net/http/server.go:3543
		// _ = "end of CoverTab[43177]"
	} else {
//line /usr/local/go/src/net/http/server.go:3544
		_go_fuzz_dep_.CoverTab[43178]++
//line /usr/local/go/src/net/http/server.go:3544
		// _ = "end of CoverTab[43178]"
//line /usr/local/go/src/net/http/server.go:3544
	}
//line /usr/local/go/src/net/http/server.go:3544
	// _ = "end of CoverTab[43171]"
//line /usr/local/go/src/net/http/server.go:3544
	_go_fuzz_dep_.CoverTab[43172]++
							h.h.ServeHTTP(rw, req)
//line /usr/local/go/src/net/http/server.go:3545
	// _ = "end of CoverTab[43172]"
}

// loggingConn is used for debugging.
type loggingConn struct {
	name	string
	net.Conn
}

var (
	uniqNameMu	sync.Mutex
	uniqNameNext	= make(map[string]int)
)

func newLoggingConn(baseName string, c net.Conn) net.Conn {
//line /usr/local/go/src/net/http/server.go:3559
	_go_fuzz_dep_.CoverTab[43179]++
							uniqNameMu.Lock()
							defer uniqNameMu.Unlock()
							uniqNameNext[baseName]++
							return &loggingConn{
		name:	fmt.Sprintf("%s-%d", baseName, uniqNameNext[baseName]),
		Conn:	c,
	}
//line /usr/local/go/src/net/http/server.go:3566
	// _ = "end of CoverTab[43179]"
}

func (c *loggingConn) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/server.go:3569
	_go_fuzz_dep_.CoverTab[43180]++
							log.Printf("%s.Write(%d) = ....", c.name, len(p))
							n, err = c.Conn.Write(p)
							log.Printf("%s.Write(%d) = %d, %v", c.name, len(p), n, err)
							return
//line /usr/local/go/src/net/http/server.go:3573
	// _ = "end of CoverTab[43180]"
}

func (c *loggingConn) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/server.go:3576
	_go_fuzz_dep_.CoverTab[43181]++
							log.Printf("%s.Read(%d) = ....", c.name, len(p))
							n, err = c.Conn.Read(p)
							log.Printf("%s.Read(%d) = %d, %v", c.name, len(p), n, err)
							return
//line /usr/local/go/src/net/http/server.go:3580
	// _ = "end of CoverTab[43181]"
}

func (c *loggingConn) Close() (err error) {
//line /usr/local/go/src/net/http/server.go:3583
	_go_fuzz_dep_.CoverTab[43182]++
							log.Printf("%s.Close() = ...", c.name)
							err = c.Conn.Close()
							log.Printf("%s.Close() = %v", c.name, err)
							return
//line /usr/local/go/src/net/http/server.go:3587
	// _ = "end of CoverTab[43182]"
}

// checkConnErrorWriter writes to c.rwc and records any write errors to c.werr.
//line /usr/local/go/src/net/http/server.go:3590
// It only contains one field (and a pointer field at that), so it
//line /usr/local/go/src/net/http/server.go:3590
// fits in an interface value without an extra allocation.
//line /usr/local/go/src/net/http/server.go:3593
type checkConnErrorWriter struct {
	c *conn
}

func (w checkConnErrorWriter) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/server.go:3597
	_go_fuzz_dep_.CoverTab[43183]++
							n, err = w.c.rwc.Write(p)
							if err != nil && func() bool {
//line /usr/local/go/src/net/http/server.go:3599
		_go_fuzz_dep_.CoverTab[43185]++
//line /usr/local/go/src/net/http/server.go:3599
		return w.c.werr == nil
//line /usr/local/go/src/net/http/server.go:3599
		// _ = "end of CoverTab[43185]"
//line /usr/local/go/src/net/http/server.go:3599
	}() {
//line /usr/local/go/src/net/http/server.go:3599
		_go_fuzz_dep_.CoverTab[43186]++
								w.c.werr = err
								w.c.cancelCtx()
//line /usr/local/go/src/net/http/server.go:3601
		// _ = "end of CoverTab[43186]"
	} else {
//line /usr/local/go/src/net/http/server.go:3602
		_go_fuzz_dep_.CoverTab[43187]++
//line /usr/local/go/src/net/http/server.go:3602
		// _ = "end of CoverTab[43187]"
//line /usr/local/go/src/net/http/server.go:3602
	}
//line /usr/local/go/src/net/http/server.go:3602
	// _ = "end of CoverTab[43183]"
//line /usr/local/go/src/net/http/server.go:3602
	_go_fuzz_dep_.CoverTab[43184]++
							return
//line /usr/local/go/src/net/http/server.go:3603
	// _ = "end of CoverTab[43184]"
}

func numLeadingCRorLF(v []byte) (n int) {
//line /usr/local/go/src/net/http/server.go:3606
	_go_fuzz_dep_.CoverTab[43188]++
							for _, b := range v {
//line /usr/local/go/src/net/http/server.go:3607
		_go_fuzz_dep_.CoverTab[43190]++
								if b == '\r' || func() bool {
//line /usr/local/go/src/net/http/server.go:3608
			_go_fuzz_dep_.CoverTab[43192]++
//line /usr/local/go/src/net/http/server.go:3608
			return b == '\n'
//line /usr/local/go/src/net/http/server.go:3608
			// _ = "end of CoverTab[43192]"
//line /usr/local/go/src/net/http/server.go:3608
		}() {
//line /usr/local/go/src/net/http/server.go:3608
			_go_fuzz_dep_.CoverTab[43193]++
									n++
									continue
//line /usr/local/go/src/net/http/server.go:3610
			// _ = "end of CoverTab[43193]"
		} else {
//line /usr/local/go/src/net/http/server.go:3611
			_go_fuzz_dep_.CoverTab[43194]++
//line /usr/local/go/src/net/http/server.go:3611
			// _ = "end of CoverTab[43194]"
//line /usr/local/go/src/net/http/server.go:3611
		}
//line /usr/local/go/src/net/http/server.go:3611
		// _ = "end of CoverTab[43190]"
//line /usr/local/go/src/net/http/server.go:3611
		_go_fuzz_dep_.CoverTab[43191]++
								break
//line /usr/local/go/src/net/http/server.go:3612
		// _ = "end of CoverTab[43191]"
	}
//line /usr/local/go/src/net/http/server.go:3613
	// _ = "end of CoverTab[43188]"
//line /usr/local/go/src/net/http/server.go:3613
	_go_fuzz_dep_.CoverTab[43189]++
							return
//line /usr/local/go/src/net/http/server.go:3614
	// _ = "end of CoverTab[43189]"

}

func strSliceContains(ss []string, s string) bool {
//line /usr/local/go/src/net/http/server.go:3618
	_go_fuzz_dep_.CoverTab[43195]++
							for _, v := range ss {
//line /usr/local/go/src/net/http/server.go:3619
		_go_fuzz_dep_.CoverTab[43197]++
								if v == s {
//line /usr/local/go/src/net/http/server.go:3620
			_go_fuzz_dep_.CoverTab[43198]++
									return true
//line /usr/local/go/src/net/http/server.go:3621
			// _ = "end of CoverTab[43198]"
		} else {
//line /usr/local/go/src/net/http/server.go:3622
			_go_fuzz_dep_.CoverTab[43199]++
//line /usr/local/go/src/net/http/server.go:3622
			// _ = "end of CoverTab[43199]"
//line /usr/local/go/src/net/http/server.go:3622
		}
//line /usr/local/go/src/net/http/server.go:3622
		// _ = "end of CoverTab[43197]"
	}
//line /usr/local/go/src/net/http/server.go:3623
	// _ = "end of CoverTab[43195]"
//line /usr/local/go/src/net/http/server.go:3623
	_go_fuzz_dep_.CoverTab[43196]++
							return false
//line /usr/local/go/src/net/http/server.go:3624
	// _ = "end of CoverTab[43196]"
}

// tlsRecordHeaderLooksLikeHTTP reports whether a TLS record header
//line /usr/local/go/src/net/http/server.go:3627
// looks like it might've been a misdirected plaintext HTTP request.
//line /usr/local/go/src/net/http/server.go:3629
func tlsRecordHeaderLooksLikeHTTP(hdr [5]byte) bool {
//line /usr/local/go/src/net/http/server.go:3629
	_go_fuzz_dep_.CoverTab[43200]++
							switch string(hdr[:]) {
	case "GET /", "HEAD ", "POST ", "PUT /", "OPTIO":
//line /usr/local/go/src/net/http/server.go:3631
		_go_fuzz_dep_.CoverTab[43202]++
								return true
//line /usr/local/go/src/net/http/server.go:3632
		// _ = "end of CoverTab[43202]"
//line /usr/local/go/src/net/http/server.go:3632
	default:
//line /usr/local/go/src/net/http/server.go:3632
		_go_fuzz_dep_.CoverTab[43203]++
//line /usr/local/go/src/net/http/server.go:3632
		// _ = "end of CoverTab[43203]"
	}
//line /usr/local/go/src/net/http/server.go:3633
	// _ = "end of CoverTab[43200]"
//line /usr/local/go/src/net/http/server.go:3633
	_go_fuzz_dep_.CoverTab[43201]++
							return false
//line /usr/local/go/src/net/http/server.go:3634
	// _ = "end of CoverTab[43201]"
}

// MaxBytesHandler returns a Handler that runs h with its ResponseWriter and Request.Body wrapped by a MaxBytesReader.
func MaxBytesHandler(h Handler, n int64) Handler {
//line /usr/local/go/src/net/http/server.go:3638
	_go_fuzz_dep_.CoverTab[43204]++
							return HandlerFunc(func(w ResponseWriter, r *Request) {
//line /usr/local/go/src/net/http/server.go:3639
		_go_fuzz_dep_.CoverTab[43205]++
								r2 := *r
								r2.Body = MaxBytesReader(w, r.Body, n)
								h.ServeHTTP(w, &r2)
//line /usr/local/go/src/net/http/server.go:3642
		// _ = "end of CoverTab[43205]"
	})
//line /usr/local/go/src/net/http/server.go:3643
	// _ = "end of CoverTab[43204]"
}

//line /usr/local/go/src/net/http/server.go:3644
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/server.go:3644
var _ = _go_fuzz_dep_.CoverTab
