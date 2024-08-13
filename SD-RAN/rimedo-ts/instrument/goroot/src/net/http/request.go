// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP Request reading and parsing.

//line /usr/local/go/src/net/http/request.go:7
package http

//line /usr/local/go/src/net/http/request.go:7
import (
//line /usr/local/go/src/net/http/request.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/request.go:7
)
//line /usr/local/go/src/net/http/request.go:7
import (
//line /usr/local/go/src/net/http/request.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/request.go:7
)

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http/httptrace"
	"net/http/internal/ascii"
	"net/textproto"
	"net/url"
	urlpkg "net/url"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/net/http/httpguts"
	"golang.org/x/net/idna"
)

const (
	defaultMaxMemory = 32 << 20	// 32 MB
)

// ErrMissingFile is returned by FormFile when the provided file field name
//line /usr/local/go/src/net/http/request.go:37
// is either not present in the request or not a file field.
//line /usr/local/go/src/net/http/request.go:39
var ErrMissingFile = errors.New("http: no such file")

// ProtocolError represents an HTTP protocol error.
//line /usr/local/go/src/net/http/request.go:41
//
//line /usr/local/go/src/net/http/request.go:41
// Deprecated: Not all errors in the http package related to protocol errors
//line /usr/local/go/src/net/http/request.go:41
// are of type ProtocolError.
//line /usr/local/go/src/net/http/request.go:45
type ProtocolError struct {
	ErrorString string
}

func (pe *ProtocolError) Error() string {
//line /usr/local/go/src/net/http/request.go:49
	_go_fuzz_dep_.CoverTab[41477]++
//line /usr/local/go/src/net/http/request.go:49
	return pe.ErrorString
//line /usr/local/go/src/net/http/request.go:49
	// _ = "end of CoverTab[41477]"
//line /usr/local/go/src/net/http/request.go:49
}

var (
	// ErrNotSupported indicates that a feature is not supported.
	//
	// It is returned by ResponseController methods to indicate that
	// the handler does not support the method, and by the Push method
	// of Pusher implementations to indicate that HTTP/2 Push support
	// is not available.
	ErrNotSupported	= &ProtocolError{"feature not supported"}

	// Deprecated: ErrUnexpectedTrailer is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	ErrUnexpectedTrailer	= &ProtocolError{"trailer header without chunked transfer encoding"}

	// ErrMissingBoundary is returned by Request.MultipartReader when the
	// request's Content-Type does not include a "boundary" parameter.
	ErrMissingBoundary	= &ProtocolError{"no multipart boundary param in Content-Type"}

	// ErrNotMultipart is returned by Request.MultipartReader when the
	// request's Content-Type is not multipart/form-data.
	ErrNotMultipart	= &ProtocolError{"request Content-Type isn't multipart/form-data"}

	// Deprecated: ErrHeaderTooLong is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	ErrHeaderTooLong	= &ProtocolError{"header too long"}

	// Deprecated: ErrShortBody is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	ErrShortBody	= &ProtocolError{"entity body too short"}

	// Deprecated: ErrMissingContentLength is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	ErrMissingContentLength	= &ProtocolError{"missing ContentLength in HEAD response"}
)

func badStringError(what, val string) error {
//line /usr/local/go/src/net/http/request.go:89
	_go_fuzz_dep_.CoverTab[41478]++
//line /usr/local/go/src/net/http/request.go:89
	return fmt.Errorf("%s %q", what, val)
//line /usr/local/go/src/net/http/request.go:89
	// _ = "end of CoverTab[41478]"
//line /usr/local/go/src/net/http/request.go:89
}

// Headers that Request.Write handles itself and should be skipped.
var reqWriteExcludeHeader = map[string]bool{
	"Host":			true,
	"User-Agent":		true,
	"Content-Length":	true,
	"Transfer-Encoding":	true,
	"Trailer":		true,
}

// A Request represents an HTTP request received by a server
//line /usr/local/go/src/net/http/request.go:100
// or to be sent by a client.
//line /usr/local/go/src/net/http/request.go:100
//
//line /usr/local/go/src/net/http/request.go:100
// The field semantics differ slightly between client and server
//line /usr/local/go/src/net/http/request.go:100
// usage. In addition to the notes on the fields below, see the
//line /usr/local/go/src/net/http/request.go:100
// documentation for Request.Write and RoundTripper.
//line /usr/local/go/src/net/http/request.go:106
type Request struct {
	// Method specifies the HTTP method (GET, POST, PUT, etc.).
	// For client requests, an empty string means GET.
	//
	// Go's HTTP client does not support sending a request with
	// the CONNECT method. See the documentation on Transport for
	// details.
	Method	string

	// URL specifies either the URI being requested (for server
	// requests) or the URL to access (for client requests).
	//
	// For server requests, the URL is parsed from the URI
	// supplied on the Request-Line as stored in RequestURI.  For
	// most requests, fields other than Path and RawQuery will be
	// empty. (See RFC 7230, Section 5.3)
	//
	// For client requests, the URL's Host specifies the server to
	// connect to, while the Request's Host field optionally
	// specifies the Host header value to send in the HTTP
	// request.
	URL	*url.URL

	// The protocol version for incoming server requests.
	//
	// For client requests, these fields are ignored. The HTTP
	// client code always uses either HTTP/1.1 or HTTP/2.
	// See the docs on Transport for details.
	Proto		string	// "HTTP/1.0"
	ProtoMajor	int	// 1
	ProtoMinor	int	// 0

	// Header contains the request header fields either received
	// by the server or to be sent by the client.
	//
	// If a server received a request with header lines,
	//
	//	Host: example.com
	//	accept-encoding: gzip, deflate
	//	Accept-Language: en-us
	//	fOO: Bar
	//	foo: two
	//
	// then
	//
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}
	//
	// For incoming requests, the Host header is promoted to the
	// Request.Host field and removed from the Header map.
	//
	// HTTP defines that header names are case-insensitive. The
	// request parser implements this by using CanonicalHeaderKey,
	// making the first character and any characters following a
	// hyphen uppercase and the rest lowercase.
	//
	// For client requests, certain headers such as Content-Length
	// and Connection are automatically written when needed and
	// values in Header may be ignored. See the documentation
	// for the Request.Write method.
	Header	Header

	// Body is the request's body.
	//
	// For client requests, a nil body means the request has no
	// body, such as a GET request. The HTTP Client's Transport
	// is responsible for calling the Close method.
	//
	// For server requests, the Request Body is always non-nil
	// but will return EOF immediately when no body is present.
	// The Server will close the request body. The ServeHTTP
	// Handler does not need to.
	//
	// Body must allow Read to be called concurrently with Close.
	// In particular, calling Close should unblock a Read waiting
	// for input.
	Body	io.ReadCloser

	// GetBody defines an optional func to return a new copy of
	// Body. It is used for client requests when a redirect requires
	// reading the body more than once. Use of GetBody still
	// requires setting Body.
	//
	// For server requests, it is unused.
	GetBody	func() (io.ReadCloser, error)

	// ContentLength records the length of the associated content.
	// The value -1 indicates that the length is unknown.
	// Values >= 0 indicate that the given number of bytes may
	// be read from Body.
	//
	// For client requests, a value of 0 with a non-nil Body is
	// also treated as unknown.
	ContentLength	int64

	// TransferEncoding lists the transfer encodings from outermost to
	// innermost. An empty list denotes the "identity" encoding.
	// TransferEncoding can usually be ignored; chunked encoding is
	// automatically added and removed as necessary when sending and
	// receiving requests.
	TransferEncoding	[]string

	// Close indicates whether to close the connection after
	// replying to this request (for servers) or after sending this
	// request and reading its response (for clients).
	//
	// For server requests, the HTTP server handles this automatically
	// and this field is not needed by Handlers.
	//
	// For client requests, setting this field prevents re-use of
	// TCP connections between requests to the same hosts, as if
	// Transport.DisableKeepAlives were set.
	Close	bool

	// For server requests, Host specifies the host on which the
	// URL is sought. For HTTP/1 (per RFC 7230, section 5.4), this
	// is either the value of the "Host" header or the host name
	// given in the URL itself. For HTTP/2, it is the value of the
	// ":authority" pseudo-header field.
	// It may be of the form "host:port". For international domain
	// names, Host may be in Punycode or Unicode form. Use
	// golang.org/x/net/idna to convert it to either format if
	// needed.
	// To prevent DNS rebinding attacks, server Handlers should
	// validate that the Host header has a value for which the
	// Handler considers itself authoritative. The included
	// ServeMux supports patterns registered to particular host
	// names and thus protects its registered Handlers.
	//
	// For client requests, Host optionally overrides the Host
	// header to send. If empty, the Request.Write method uses
	// the value of URL.Host. Host may contain an international
	// domain name.
	Host	string

	// Form contains the parsed form data, including both the URL
	// field's query parameters and the PATCH, POST, or PUT form data.
	// This field is only available after ParseForm is called.
	// The HTTP client ignores Form and uses Body instead.
	Form	url.Values

	// PostForm contains the parsed form data from PATCH, POST
	// or PUT body parameters.
	//
	// This field is only available after ParseForm is called.
	// The HTTP client ignores PostForm and uses Body instead.
	PostForm	url.Values

	// MultipartForm is the parsed multipart form, including file uploads.
	// This field is only available after ParseMultipartForm is called.
	// The HTTP client ignores MultipartForm and uses Body instead.
	MultipartForm	*multipart.Form

	// Trailer specifies additional headers that are sent after the request
	// body.
	//
	// For server requests, the Trailer map initially contains only the
	// trailer keys, with nil values. (The client declares which trailers it
	// will later send.)  While the handler is reading from Body, it must
	// not reference Trailer. After reading from Body returns EOF, Trailer
	// can be read again and will contain non-nil values, if they were sent
	// by the client.
	//
	// For client requests, Trailer must be initialized to a map containing
	// the trailer keys to later send. The values may be nil or their final
	// values. The ContentLength must be 0 or -1, to send a chunked request.
	// After the HTTP request is sent the map values can be updated while
	// the request body is read. Once the body returns EOF, the caller must
	// not mutate Trailer.
	//
	// Few HTTP clients, servers, or proxies support HTTP trailers.
	Trailer	Header

	// RemoteAddr allows HTTP servers and other software to record
	// the network address that sent the request, usually for
	// logging. This field is not filled in by ReadRequest and
	// has no defined format. The HTTP server in this package
	// sets RemoteAddr to an "IP:port" address before invoking a
	// handler.
	// This field is ignored by the HTTP client.
	RemoteAddr	string

	// RequestURI is the unmodified request-target of the
	// Request-Line (RFC 7230, Section 3.1.1) as sent by the client
	// to a server. Usually the URL field should be used instead.
	// It is an error to set this field in an HTTP client request.
	RequestURI	string

	// TLS allows HTTP servers and other software to record
	// information about the TLS connection on which the request
	// was received. This field is not filled in by ReadRequest.
	// The HTTP server in this package sets the field for
	// TLS-enabled connections before invoking a handler;
	// otherwise it leaves the field nil.
	// This field is ignored by the HTTP client.
	TLS	*tls.ConnectionState

	// Cancel is an optional channel whose closure indicates that the client
	// request should be regarded as canceled. Not all implementations of
	// RoundTripper may support Cancel.
	//
	// For server requests, this field is not applicable.
	//
	// Deprecated: Set the Request's context with NewRequestWithContext
	// instead. If a Request's Cancel field and context are both
	// set, it is undefined whether Cancel is respected.
	Cancel	<-chan struct{}

	// Response is the redirect response which caused this request
	// to be created. This field is only populated during client
	// redirects.
	Response	*Response

	// ctx is either the client or server context. It should only
	// be modified via copying the whole Request using Clone or WithContext.
	// It is unexported to prevent people from using Context wrong
	// and mutating the contexts held by callers of the same request.
	ctx	context.Context
}

// Context returns the request's context. To change the context, use
//line /usr/local/go/src/net/http/request.go:329
// Clone or WithContext.
//line /usr/local/go/src/net/http/request.go:329
//
//line /usr/local/go/src/net/http/request.go:329
// The returned context is always non-nil; it defaults to the
//line /usr/local/go/src/net/http/request.go:329
// background context.
//line /usr/local/go/src/net/http/request.go:329
//
//line /usr/local/go/src/net/http/request.go:329
// For outgoing client requests, the context controls cancellation.
//line /usr/local/go/src/net/http/request.go:329
//
//line /usr/local/go/src/net/http/request.go:329
// For incoming server requests, the context is canceled when the
//line /usr/local/go/src/net/http/request.go:329
// client's connection closes, the request is canceled (with HTTP/2),
//line /usr/local/go/src/net/http/request.go:329
// or when the ServeHTTP method returns.
//line /usr/local/go/src/net/http/request.go:340
func (r *Request) Context() context.Context {
//line /usr/local/go/src/net/http/request.go:340
	_go_fuzz_dep_.CoverTab[41479]++
							if r.ctx != nil {
//line /usr/local/go/src/net/http/request.go:341
		_go_fuzz_dep_.CoverTab[41481]++
								return r.ctx
//line /usr/local/go/src/net/http/request.go:342
		// _ = "end of CoverTab[41481]"
	} else {
//line /usr/local/go/src/net/http/request.go:343
		_go_fuzz_dep_.CoverTab[41482]++
//line /usr/local/go/src/net/http/request.go:343
		// _ = "end of CoverTab[41482]"
//line /usr/local/go/src/net/http/request.go:343
	}
//line /usr/local/go/src/net/http/request.go:343
	// _ = "end of CoverTab[41479]"
//line /usr/local/go/src/net/http/request.go:343
	_go_fuzz_dep_.CoverTab[41480]++
							return context.Background()
//line /usr/local/go/src/net/http/request.go:344
	// _ = "end of CoverTab[41480]"
}

// WithContext returns a shallow copy of r with its context changed
//line /usr/local/go/src/net/http/request.go:347
// to ctx. The provided ctx must be non-nil.
//line /usr/local/go/src/net/http/request.go:347
//
//line /usr/local/go/src/net/http/request.go:347
// For outgoing client request, the context controls the entire
//line /usr/local/go/src/net/http/request.go:347
// lifetime of a request and its response: obtaining a connection,
//line /usr/local/go/src/net/http/request.go:347
// sending the request, and reading the response headers and body.
//line /usr/local/go/src/net/http/request.go:347
//
//line /usr/local/go/src/net/http/request.go:347
// To create a new request with a context, use NewRequestWithContext.
//line /usr/local/go/src/net/http/request.go:347
// To make a deep copy of a request with a new context, use Request.Clone.
//line /usr/local/go/src/net/http/request.go:356
func (r *Request) WithContext(ctx context.Context) *Request {
//line /usr/local/go/src/net/http/request.go:356
	_go_fuzz_dep_.CoverTab[41483]++
							if ctx == nil {
//line /usr/local/go/src/net/http/request.go:357
		_go_fuzz_dep_.CoverTab[41485]++
								panic("nil context")
//line /usr/local/go/src/net/http/request.go:358
		// _ = "end of CoverTab[41485]"
	} else {
//line /usr/local/go/src/net/http/request.go:359
		_go_fuzz_dep_.CoverTab[41486]++
//line /usr/local/go/src/net/http/request.go:359
		// _ = "end of CoverTab[41486]"
//line /usr/local/go/src/net/http/request.go:359
	}
//line /usr/local/go/src/net/http/request.go:359
	// _ = "end of CoverTab[41483]"
//line /usr/local/go/src/net/http/request.go:359
	_go_fuzz_dep_.CoverTab[41484]++
							r2 := new(Request)
							*r2 = *r
							r2.ctx = ctx
							return r2
//line /usr/local/go/src/net/http/request.go:363
	// _ = "end of CoverTab[41484]"
}

// Clone returns a deep copy of r with its context changed to ctx.
//line /usr/local/go/src/net/http/request.go:366
// The provided ctx must be non-nil.
//line /usr/local/go/src/net/http/request.go:366
//
//line /usr/local/go/src/net/http/request.go:366
// For an outgoing client request, the context controls the entire
//line /usr/local/go/src/net/http/request.go:366
// lifetime of a request and its response: obtaining a connection,
//line /usr/local/go/src/net/http/request.go:366
// sending the request, and reading the response headers and body.
//line /usr/local/go/src/net/http/request.go:372
func (r *Request) Clone(ctx context.Context) *Request {
//line /usr/local/go/src/net/http/request.go:372
	_go_fuzz_dep_.CoverTab[41487]++
							if ctx == nil {
//line /usr/local/go/src/net/http/request.go:373
		_go_fuzz_dep_.CoverTab[41492]++
								panic("nil context")
//line /usr/local/go/src/net/http/request.go:374
		// _ = "end of CoverTab[41492]"
	} else {
//line /usr/local/go/src/net/http/request.go:375
		_go_fuzz_dep_.CoverTab[41493]++
//line /usr/local/go/src/net/http/request.go:375
		// _ = "end of CoverTab[41493]"
//line /usr/local/go/src/net/http/request.go:375
	}
//line /usr/local/go/src/net/http/request.go:375
	// _ = "end of CoverTab[41487]"
//line /usr/local/go/src/net/http/request.go:375
	_go_fuzz_dep_.CoverTab[41488]++
							r2 := new(Request)
							*r2 = *r
							r2.ctx = ctx
							r2.URL = cloneURL(r.URL)
							if r.Header != nil {
//line /usr/local/go/src/net/http/request.go:380
		_go_fuzz_dep_.CoverTab[41494]++
								r2.Header = r.Header.Clone()
//line /usr/local/go/src/net/http/request.go:381
		// _ = "end of CoverTab[41494]"
	} else {
//line /usr/local/go/src/net/http/request.go:382
		_go_fuzz_dep_.CoverTab[41495]++
//line /usr/local/go/src/net/http/request.go:382
		// _ = "end of CoverTab[41495]"
//line /usr/local/go/src/net/http/request.go:382
	}
//line /usr/local/go/src/net/http/request.go:382
	// _ = "end of CoverTab[41488]"
//line /usr/local/go/src/net/http/request.go:382
	_go_fuzz_dep_.CoverTab[41489]++
							if r.Trailer != nil {
//line /usr/local/go/src/net/http/request.go:383
		_go_fuzz_dep_.CoverTab[41496]++
								r2.Trailer = r.Trailer.Clone()
//line /usr/local/go/src/net/http/request.go:384
		// _ = "end of CoverTab[41496]"
	} else {
//line /usr/local/go/src/net/http/request.go:385
		_go_fuzz_dep_.CoverTab[41497]++
//line /usr/local/go/src/net/http/request.go:385
		// _ = "end of CoverTab[41497]"
//line /usr/local/go/src/net/http/request.go:385
	}
//line /usr/local/go/src/net/http/request.go:385
	// _ = "end of CoverTab[41489]"
//line /usr/local/go/src/net/http/request.go:385
	_go_fuzz_dep_.CoverTab[41490]++
							if s := r.TransferEncoding; s != nil {
//line /usr/local/go/src/net/http/request.go:386
		_go_fuzz_dep_.CoverTab[41498]++
								s2 := make([]string, len(s))
								copy(s2, s)
								r2.TransferEncoding = s2
//line /usr/local/go/src/net/http/request.go:389
		// _ = "end of CoverTab[41498]"
	} else {
//line /usr/local/go/src/net/http/request.go:390
		_go_fuzz_dep_.CoverTab[41499]++
//line /usr/local/go/src/net/http/request.go:390
		// _ = "end of CoverTab[41499]"
//line /usr/local/go/src/net/http/request.go:390
	}
//line /usr/local/go/src/net/http/request.go:390
	// _ = "end of CoverTab[41490]"
//line /usr/local/go/src/net/http/request.go:390
	_go_fuzz_dep_.CoverTab[41491]++
							r2.Form = cloneURLValues(r.Form)
							r2.PostForm = cloneURLValues(r.PostForm)
							r2.MultipartForm = cloneMultipartForm(r.MultipartForm)
							return r2
//line /usr/local/go/src/net/http/request.go:394
	// _ = "end of CoverTab[41491]"
}

// ProtoAtLeast reports whether the HTTP protocol used
//line /usr/local/go/src/net/http/request.go:397
// in the request is at least major.minor.
//line /usr/local/go/src/net/http/request.go:399
func (r *Request) ProtoAtLeast(major, minor int) bool {
//line /usr/local/go/src/net/http/request.go:399
	_go_fuzz_dep_.CoverTab[41500]++
							return r.ProtoMajor > major || func() bool {
//line /usr/local/go/src/net/http/request.go:400
		_go_fuzz_dep_.CoverTab[41501]++
//line /usr/local/go/src/net/http/request.go:400
		return r.ProtoMajor == major && func() bool {
									_go_fuzz_dep_.CoverTab[41502]++
//line /usr/local/go/src/net/http/request.go:401
			return r.ProtoMinor >= minor
//line /usr/local/go/src/net/http/request.go:401
			// _ = "end of CoverTab[41502]"
//line /usr/local/go/src/net/http/request.go:401
		}()
//line /usr/local/go/src/net/http/request.go:401
		// _ = "end of CoverTab[41501]"
//line /usr/local/go/src/net/http/request.go:401
	}()
//line /usr/local/go/src/net/http/request.go:401
	// _ = "end of CoverTab[41500]"
}

// UserAgent returns the client's User-Agent, if sent in the request.
func (r *Request) UserAgent() string {
//line /usr/local/go/src/net/http/request.go:405
	_go_fuzz_dep_.CoverTab[41503]++
							return r.Header.Get("User-Agent")
//line /usr/local/go/src/net/http/request.go:406
	// _ = "end of CoverTab[41503]"
}

// Cookies parses and returns the HTTP cookies sent with the request.
func (r *Request) Cookies() []*Cookie {
//line /usr/local/go/src/net/http/request.go:410
	_go_fuzz_dep_.CoverTab[41504]++
							return readCookies(r.Header, "")
//line /usr/local/go/src/net/http/request.go:411
	// _ = "end of CoverTab[41504]"
}

// ErrNoCookie is returned by Request's Cookie method when a cookie is not found.
var ErrNoCookie = errors.New("http: named cookie not present")

// Cookie returns the named cookie provided in the request or
//line /usr/local/go/src/net/http/request.go:417
// ErrNoCookie if not found.
//line /usr/local/go/src/net/http/request.go:417
// If multiple cookies match the given name, only one cookie will
//line /usr/local/go/src/net/http/request.go:417
// be returned.
//line /usr/local/go/src/net/http/request.go:421
func (r *Request) Cookie(name string) (*Cookie, error) {
//line /usr/local/go/src/net/http/request.go:421
	_go_fuzz_dep_.CoverTab[41505]++
							if name == "" {
//line /usr/local/go/src/net/http/request.go:422
		_go_fuzz_dep_.CoverTab[41508]++
								return nil, ErrNoCookie
//line /usr/local/go/src/net/http/request.go:423
		// _ = "end of CoverTab[41508]"
	} else {
//line /usr/local/go/src/net/http/request.go:424
		_go_fuzz_dep_.CoverTab[41509]++
//line /usr/local/go/src/net/http/request.go:424
		// _ = "end of CoverTab[41509]"
//line /usr/local/go/src/net/http/request.go:424
	}
//line /usr/local/go/src/net/http/request.go:424
	// _ = "end of CoverTab[41505]"
//line /usr/local/go/src/net/http/request.go:424
	_go_fuzz_dep_.CoverTab[41506]++
							for _, c := range readCookies(r.Header, name) {
//line /usr/local/go/src/net/http/request.go:425
		_go_fuzz_dep_.CoverTab[41510]++
								return c, nil
//line /usr/local/go/src/net/http/request.go:426
		// _ = "end of CoverTab[41510]"
	}
//line /usr/local/go/src/net/http/request.go:427
	// _ = "end of CoverTab[41506]"
//line /usr/local/go/src/net/http/request.go:427
	_go_fuzz_dep_.CoverTab[41507]++
							return nil, ErrNoCookie
//line /usr/local/go/src/net/http/request.go:428
	// _ = "end of CoverTab[41507]"
}

// AddCookie adds a cookie to the request. Per RFC 6265 section 5.4,
//line /usr/local/go/src/net/http/request.go:431
// AddCookie does not attach more than one Cookie header field. That
//line /usr/local/go/src/net/http/request.go:431
// means all cookies, if any, are written into the same line,
//line /usr/local/go/src/net/http/request.go:431
// separated by semicolon.
//line /usr/local/go/src/net/http/request.go:431
// AddCookie only sanitizes c's name and value, and does not sanitize
//line /usr/local/go/src/net/http/request.go:431
// a Cookie header already present in the request.
//line /usr/local/go/src/net/http/request.go:437
func (r *Request) AddCookie(c *Cookie) {
//line /usr/local/go/src/net/http/request.go:437
	_go_fuzz_dep_.CoverTab[41511]++
							s := fmt.Sprintf("%s=%s", sanitizeCookieName(c.Name), sanitizeCookieValue(c.Value))
							if c := r.Header.Get("Cookie"); c != "" {
//line /usr/local/go/src/net/http/request.go:439
		_go_fuzz_dep_.CoverTab[41512]++
								r.Header.Set("Cookie", c+"; "+s)
//line /usr/local/go/src/net/http/request.go:440
		// _ = "end of CoverTab[41512]"
	} else {
//line /usr/local/go/src/net/http/request.go:441
		_go_fuzz_dep_.CoverTab[41513]++
								r.Header.Set("Cookie", s)
//line /usr/local/go/src/net/http/request.go:442
		// _ = "end of CoverTab[41513]"
	}
//line /usr/local/go/src/net/http/request.go:443
	// _ = "end of CoverTab[41511]"
}

// Referer returns the referring URL, if sent in the request.
//line /usr/local/go/src/net/http/request.go:446
//
//line /usr/local/go/src/net/http/request.go:446
// Referer is misspelled as in the request itself, a mistake from the
//line /usr/local/go/src/net/http/request.go:446
// earliest days of HTTP.  This value can also be fetched from the
//line /usr/local/go/src/net/http/request.go:446
// Header map as Header["Referer"]; the benefit of making it available
//line /usr/local/go/src/net/http/request.go:446
// as a method is that the compiler can diagnose programs that use the
//line /usr/local/go/src/net/http/request.go:446
// alternate (correct English) spelling req.Referrer() but cannot
//line /usr/local/go/src/net/http/request.go:446
// diagnose programs that use Header["Referrer"].
//line /usr/local/go/src/net/http/request.go:454
func (r *Request) Referer() string {
//line /usr/local/go/src/net/http/request.go:454
	_go_fuzz_dep_.CoverTab[41514]++
							return r.Header.Get("Referer")
//line /usr/local/go/src/net/http/request.go:455
	// _ = "end of CoverTab[41514]"
}

// multipartByReader is a sentinel value.
//line /usr/local/go/src/net/http/request.go:458
// Its presence in Request.MultipartForm indicates that parsing of the request
//line /usr/local/go/src/net/http/request.go:458
// body has been handed off to a MultipartReader instead of ParseMultipartForm.
//line /usr/local/go/src/net/http/request.go:461
var multipartByReader = &multipart.Form{
	Value:	make(map[string][]string),
	File:	make(map[string][]*multipart.FileHeader),
}

// MultipartReader returns a MIME multipart reader if this is a
//line /usr/local/go/src/net/http/request.go:466
// multipart/form-data or a multipart/mixed POST request, else returns nil and an error.
//line /usr/local/go/src/net/http/request.go:466
// Use this function instead of ParseMultipartForm to
//line /usr/local/go/src/net/http/request.go:466
// process the request body as a stream.
//line /usr/local/go/src/net/http/request.go:470
func (r *Request) MultipartReader() (*multipart.Reader, error) {
//line /usr/local/go/src/net/http/request.go:470
	_go_fuzz_dep_.CoverTab[41515]++
							if r.MultipartForm == multipartByReader {
//line /usr/local/go/src/net/http/request.go:471
		_go_fuzz_dep_.CoverTab[41518]++
								return nil, errors.New("http: MultipartReader called twice")
//line /usr/local/go/src/net/http/request.go:472
		// _ = "end of CoverTab[41518]"
	} else {
//line /usr/local/go/src/net/http/request.go:473
		_go_fuzz_dep_.CoverTab[41519]++
//line /usr/local/go/src/net/http/request.go:473
		// _ = "end of CoverTab[41519]"
//line /usr/local/go/src/net/http/request.go:473
	}
//line /usr/local/go/src/net/http/request.go:473
	// _ = "end of CoverTab[41515]"
//line /usr/local/go/src/net/http/request.go:473
	_go_fuzz_dep_.CoverTab[41516]++
							if r.MultipartForm != nil {
//line /usr/local/go/src/net/http/request.go:474
		_go_fuzz_dep_.CoverTab[41520]++
								return nil, errors.New("http: multipart handled by ParseMultipartForm")
//line /usr/local/go/src/net/http/request.go:475
		// _ = "end of CoverTab[41520]"
	} else {
//line /usr/local/go/src/net/http/request.go:476
		_go_fuzz_dep_.CoverTab[41521]++
//line /usr/local/go/src/net/http/request.go:476
		// _ = "end of CoverTab[41521]"
//line /usr/local/go/src/net/http/request.go:476
	}
//line /usr/local/go/src/net/http/request.go:476
	// _ = "end of CoverTab[41516]"
//line /usr/local/go/src/net/http/request.go:476
	_go_fuzz_dep_.CoverTab[41517]++
							r.MultipartForm = multipartByReader
							return r.multipartReader(true)
//line /usr/local/go/src/net/http/request.go:478
	// _ = "end of CoverTab[41517]"
}

func (r *Request) multipartReader(allowMixed bool) (*multipart.Reader, error) {
//line /usr/local/go/src/net/http/request.go:481
	_go_fuzz_dep_.CoverTab[41522]++
							v := r.Header.Get("Content-Type")
							if v == "" {
//line /usr/local/go/src/net/http/request.go:483
		_go_fuzz_dep_.CoverTab[41527]++
								return nil, ErrNotMultipart
//line /usr/local/go/src/net/http/request.go:484
		// _ = "end of CoverTab[41527]"
	} else {
//line /usr/local/go/src/net/http/request.go:485
		_go_fuzz_dep_.CoverTab[41528]++
//line /usr/local/go/src/net/http/request.go:485
		// _ = "end of CoverTab[41528]"
//line /usr/local/go/src/net/http/request.go:485
	}
//line /usr/local/go/src/net/http/request.go:485
	// _ = "end of CoverTab[41522]"
//line /usr/local/go/src/net/http/request.go:485
	_go_fuzz_dep_.CoverTab[41523]++
							if r.Body == nil {
//line /usr/local/go/src/net/http/request.go:486
		_go_fuzz_dep_.CoverTab[41529]++
								return nil, errors.New("missing form body")
//line /usr/local/go/src/net/http/request.go:487
		// _ = "end of CoverTab[41529]"
	} else {
//line /usr/local/go/src/net/http/request.go:488
		_go_fuzz_dep_.CoverTab[41530]++
//line /usr/local/go/src/net/http/request.go:488
		// _ = "end of CoverTab[41530]"
//line /usr/local/go/src/net/http/request.go:488
	}
//line /usr/local/go/src/net/http/request.go:488
	// _ = "end of CoverTab[41523]"
//line /usr/local/go/src/net/http/request.go:488
	_go_fuzz_dep_.CoverTab[41524]++
							d, params, err := mime.ParseMediaType(v)
							if err != nil || func() bool {
//line /usr/local/go/src/net/http/request.go:490
		_go_fuzz_dep_.CoverTab[41531]++
//line /usr/local/go/src/net/http/request.go:490
		return !(d == "multipart/form-data" || func() bool {
//line /usr/local/go/src/net/http/request.go:490
			_go_fuzz_dep_.CoverTab[41532]++
//line /usr/local/go/src/net/http/request.go:490
			return allowMixed && func() bool {
//line /usr/local/go/src/net/http/request.go:490
				_go_fuzz_dep_.CoverTab[41533]++
//line /usr/local/go/src/net/http/request.go:490
				return d == "multipart/mixed"
//line /usr/local/go/src/net/http/request.go:490
				// _ = "end of CoverTab[41533]"
//line /usr/local/go/src/net/http/request.go:490
			}()
//line /usr/local/go/src/net/http/request.go:490
			// _ = "end of CoverTab[41532]"
//line /usr/local/go/src/net/http/request.go:490
		}())
//line /usr/local/go/src/net/http/request.go:490
		// _ = "end of CoverTab[41531]"
//line /usr/local/go/src/net/http/request.go:490
	}() {
//line /usr/local/go/src/net/http/request.go:490
		_go_fuzz_dep_.CoverTab[41534]++
								return nil, ErrNotMultipart
//line /usr/local/go/src/net/http/request.go:491
		// _ = "end of CoverTab[41534]"
	} else {
//line /usr/local/go/src/net/http/request.go:492
		_go_fuzz_dep_.CoverTab[41535]++
//line /usr/local/go/src/net/http/request.go:492
		// _ = "end of CoverTab[41535]"
//line /usr/local/go/src/net/http/request.go:492
	}
//line /usr/local/go/src/net/http/request.go:492
	// _ = "end of CoverTab[41524]"
//line /usr/local/go/src/net/http/request.go:492
	_go_fuzz_dep_.CoverTab[41525]++
							boundary, ok := params["boundary"]
							if !ok {
//line /usr/local/go/src/net/http/request.go:494
		_go_fuzz_dep_.CoverTab[41536]++
								return nil, ErrMissingBoundary
//line /usr/local/go/src/net/http/request.go:495
		// _ = "end of CoverTab[41536]"
	} else {
//line /usr/local/go/src/net/http/request.go:496
		_go_fuzz_dep_.CoverTab[41537]++
//line /usr/local/go/src/net/http/request.go:496
		// _ = "end of CoverTab[41537]"
//line /usr/local/go/src/net/http/request.go:496
	}
//line /usr/local/go/src/net/http/request.go:496
	// _ = "end of CoverTab[41525]"
//line /usr/local/go/src/net/http/request.go:496
	_go_fuzz_dep_.CoverTab[41526]++
							return multipart.NewReader(r.Body, boundary), nil
//line /usr/local/go/src/net/http/request.go:497
	// _ = "end of CoverTab[41526]"
}

// isH2Upgrade reports whether r represents the http2 "client preface"
//line /usr/local/go/src/net/http/request.go:500
// magic string.
//line /usr/local/go/src/net/http/request.go:502
func (r *Request) isH2Upgrade() bool {
//line /usr/local/go/src/net/http/request.go:502
	_go_fuzz_dep_.CoverTab[41538]++
							return r.Method == "PRI" && func() bool {
//line /usr/local/go/src/net/http/request.go:503
		_go_fuzz_dep_.CoverTab[41539]++
//line /usr/local/go/src/net/http/request.go:503
		return len(r.Header) == 0
//line /usr/local/go/src/net/http/request.go:503
		// _ = "end of CoverTab[41539]"
//line /usr/local/go/src/net/http/request.go:503
	}() && func() bool {
//line /usr/local/go/src/net/http/request.go:503
		_go_fuzz_dep_.CoverTab[41540]++
//line /usr/local/go/src/net/http/request.go:503
		return r.URL.Path == "*"
//line /usr/local/go/src/net/http/request.go:503
		// _ = "end of CoverTab[41540]"
//line /usr/local/go/src/net/http/request.go:503
	}() && func() bool {
//line /usr/local/go/src/net/http/request.go:503
		_go_fuzz_dep_.CoverTab[41541]++
//line /usr/local/go/src/net/http/request.go:503
		return r.Proto == "HTTP/2.0"
//line /usr/local/go/src/net/http/request.go:503
		// _ = "end of CoverTab[41541]"
//line /usr/local/go/src/net/http/request.go:503
	}()
//line /usr/local/go/src/net/http/request.go:503
	// _ = "end of CoverTab[41538]"
}

// Return value if nonempty, def otherwise.
func valueOrDefault(value, def string) string {
//line /usr/local/go/src/net/http/request.go:507
	_go_fuzz_dep_.CoverTab[41542]++
							if value != "" {
//line /usr/local/go/src/net/http/request.go:508
		_go_fuzz_dep_.CoverTab[41544]++
								return value
//line /usr/local/go/src/net/http/request.go:509
		// _ = "end of CoverTab[41544]"
	} else {
//line /usr/local/go/src/net/http/request.go:510
		_go_fuzz_dep_.CoverTab[41545]++
//line /usr/local/go/src/net/http/request.go:510
		// _ = "end of CoverTab[41545]"
//line /usr/local/go/src/net/http/request.go:510
	}
//line /usr/local/go/src/net/http/request.go:510
	// _ = "end of CoverTab[41542]"
//line /usr/local/go/src/net/http/request.go:510
	_go_fuzz_dep_.CoverTab[41543]++
							return def
//line /usr/local/go/src/net/http/request.go:511
	// _ = "end of CoverTab[41543]"
}

// NOTE: This is not intended to reflect the actual Go version being used.
//line /usr/local/go/src/net/http/request.go:514
// It was changed at the time of Go 1.1 release because the former User-Agent
//line /usr/local/go/src/net/http/request.go:514
// had ended up blocked by some intrusion detection systems.
//line /usr/local/go/src/net/http/request.go:514
// See https://codereview.appspot.com/7532043.
//line /usr/local/go/src/net/http/request.go:518
const defaultUserAgent = "Go-http-client/1.1"

// Write writes an HTTP/1.1 request, which is the header and body, in wire format.
//line /usr/local/go/src/net/http/request.go:520
// This method consults the following fields of the request:
//line /usr/local/go/src/net/http/request.go:520
//
//line /usr/local/go/src/net/http/request.go:520
//	Host
//line /usr/local/go/src/net/http/request.go:520
//	URL
//line /usr/local/go/src/net/http/request.go:520
//	Method (defaults to "GET")
//line /usr/local/go/src/net/http/request.go:520
//	Header
//line /usr/local/go/src/net/http/request.go:520
//	ContentLength
//line /usr/local/go/src/net/http/request.go:520
//	TransferEncoding
//line /usr/local/go/src/net/http/request.go:520
//	Body
//line /usr/local/go/src/net/http/request.go:520
//
//line /usr/local/go/src/net/http/request.go:520
// If Body is present, Content-Length is <= 0 and TransferEncoding
//line /usr/local/go/src/net/http/request.go:520
// hasn't been set to "identity", Write adds "Transfer-Encoding:
//line /usr/local/go/src/net/http/request.go:520
// chunked" to the header. Body is closed after it is sent.
//line /usr/local/go/src/net/http/request.go:534
func (r *Request) Write(w io.Writer) error {
//line /usr/local/go/src/net/http/request.go:534
	_go_fuzz_dep_.CoverTab[41546]++
							return r.write(w, false, nil, nil)
//line /usr/local/go/src/net/http/request.go:535
	// _ = "end of CoverTab[41546]"
}

// WriteProxy is like Write but writes the request in the form
//line /usr/local/go/src/net/http/request.go:538
// expected by an HTTP proxy. In particular, WriteProxy writes the
//line /usr/local/go/src/net/http/request.go:538
// initial Request-URI line of the request with an absolute URI, per
//line /usr/local/go/src/net/http/request.go:538
// section 5.3 of RFC 7230, including the scheme and host.
//line /usr/local/go/src/net/http/request.go:538
// In either case, WriteProxy also writes a Host header, using
//line /usr/local/go/src/net/http/request.go:538
// either r.Host or r.URL.Host.
//line /usr/local/go/src/net/http/request.go:544
func (r *Request) WriteProxy(w io.Writer) error {
//line /usr/local/go/src/net/http/request.go:544
	_go_fuzz_dep_.CoverTab[41547]++
							return r.write(w, true, nil, nil)
//line /usr/local/go/src/net/http/request.go:545
	// _ = "end of CoverTab[41547]"
}

// errMissingHost is returned by Write when there is no Host or URL present in
//line /usr/local/go/src/net/http/request.go:548
// the Request.
//line /usr/local/go/src/net/http/request.go:550
var errMissingHost = errors.New("http: Request.Write on Request with no Host or URL set")

// extraHeaders may be nil
//line /usr/local/go/src/net/http/request.go:552
// waitForContinue may be nil
//line /usr/local/go/src/net/http/request.go:552
// always closes body
//line /usr/local/go/src/net/http/request.go:555
func (r *Request) write(w io.Writer, usingProxy bool, extraHeaders Header, waitForContinue func() bool) (err error) {
//line /usr/local/go/src/net/http/request.go:555
	_go_fuzz_dep_.CoverTab[41548]++
							trace := httptrace.ContextClientTrace(r.Context())
							if trace != nil && func() bool {
//line /usr/local/go/src/net/http/request.go:557
		_go_fuzz_dep_.CoverTab[41572]++
//line /usr/local/go/src/net/http/request.go:557
		return trace.WroteRequest != nil
//line /usr/local/go/src/net/http/request.go:557
		// _ = "end of CoverTab[41572]"
//line /usr/local/go/src/net/http/request.go:557
	}() {
//line /usr/local/go/src/net/http/request.go:557
		_go_fuzz_dep_.CoverTab[41573]++
								defer func() {
//line /usr/local/go/src/net/http/request.go:558
			_go_fuzz_dep_.CoverTab[41574]++
									trace.WroteRequest(httptrace.WroteRequestInfo{
				Err: err,
			})
//line /usr/local/go/src/net/http/request.go:561
			// _ = "end of CoverTab[41574]"
		}()
//line /usr/local/go/src/net/http/request.go:562
		// _ = "end of CoverTab[41573]"
	} else {
//line /usr/local/go/src/net/http/request.go:563
		_go_fuzz_dep_.CoverTab[41575]++
//line /usr/local/go/src/net/http/request.go:563
		// _ = "end of CoverTab[41575]"
//line /usr/local/go/src/net/http/request.go:563
	}
//line /usr/local/go/src/net/http/request.go:563
	// _ = "end of CoverTab[41548]"
//line /usr/local/go/src/net/http/request.go:563
	_go_fuzz_dep_.CoverTab[41549]++
							closed := false
							defer func() {
//line /usr/local/go/src/net/http/request.go:565
		_go_fuzz_dep_.CoverTab[41576]++
								if closed {
//line /usr/local/go/src/net/http/request.go:566
			_go_fuzz_dep_.CoverTab[41578]++
									return
//line /usr/local/go/src/net/http/request.go:567
			// _ = "end of CoverTab[41578]"
		} else {
//line /usr/local/go/src/net/http/request.go:568
			_go_fuzz_dep_.CoverTab[41579]++
//line /usr/local/go/src/net/http/request.go:568
			// _ = "end of CoverTab[41579]"
//line /usr/local/go/src/net/http/request.go:568
		}
//line /usr/local/go/src/net/http/request.go:568
		// _ = "end of CoverTab[41576]"
//line /usr/local/go/src/net/http/request.go:568
		_go_fuzz_dep_.CoverTab[41577]++
								if closeErr := r.closeBody(); closeErr != nil && func() bool {
//line /usr/local/go/src/net/http/request.go:569
			_go_fuzz_dep_.CoverTab[41580]++
//line /usr/local/go/src/net/http/request.go:569
			return err == nil
//line /usr/local/go/src/net/http/request.go:569
			// _ = "end of CoverTab[41580]"
//line /usr/local/go/src/net/http/request.go:569
		}() {
//line /usr/local/go/src/net/http/request.go:569
			_go_fuzz_dep_.CoverTab[41581]++
									err = closeErr
//line /usr/local/go/src/net/http/request.go:570
			// _ = "end of CoverTab[41581]"
		} else {
//line /usr/local/go/src/net/http/request.go:571
			_go_fuzz_dep_.CoverTab[41582]++
//line /usr/local/go/src/net/http/request.go:571
			// _ = "end of CoverTab[41582]"
//line /usr/local/go/src/net/http/request.go:571
		}
//line /usr/local/go/src/net/http/request.go:571
		// _ = "end of CoverTab[41577]"
	}()
//line /usr/local/go/src/net/http/request.go:572
	// _ = "end of CoverTab[41549]"
//line /usr/local/go/src/net/http/request.go:572
	_go_fuzz_dep_.CoverTab[41550]++

//line /usr/local/go/src/net/http/request.go:578
	host := r.Host
	if host == "" {
//line /usr/local/go/src/net/http/request.go:579
		_go_fuzz_dep_.CoverTab[41583]++
								if r.URL == nil {
//line /usr/local/go/src/net/http/request.go:580
			_go_fuzz_dep_.CoverTab[41585]++
									return errMissingHost
//line /usr/local/go/src/net/http/request.go:581
			// _ = "end of CoverTab[41585]"
		} else {
//line /usr/local/go/src/net/http/request.go:582
			_go_fuzz_dep_.CoverTab[41586]++
//line /usr/local/go/src/net/http/request.go:582
			// _ = "end of CoverTab[41586]"
//line /usr/local/go/src/net/http/request.go:582
		}
//line /usr/local/go/src/net/http/request.go:582
		// _ = "end of CoverTab[41583]"
//line /usr/local/go/src/net/http/request.go:582
		_go_fuzz_dep_.CoverTab[41584]++
								host = r.URL.Host
//line /usr/local/go/src/net/http/request.go:583
		// _ = "end of CoverTab[41584]"
	} else {
//line /usr/local/go/src/net/http/request.go:584
		_go_fuzz_dep_.CoverTab[41587]++
//line /usr/local/go/src/net/http/request.go:584
		// _ = "end of CoverTab[41587]"
//line /usr/local/go/src/net/http/request.go:584
	}
//line /usr/local/go/src/net/http/request.go:584
	// _ = "end of CoverTab[41550]"
//line /usr/local/go/src/net/http/request.go:584
	_go_fuzz_dep_.CoverTab[41551]++
							host, err = httpguts.PunycodeHostPort(host)
							if err != nil {
//line /usr/local/go/src/net/http/request.go:586
		_go_fuzz_dep_.CoverTab[41588]++
								return err
//line /usr/local/go/src/net/http/request.go:587
		// _ = "end of CoverTab[41588]"
	} else {
//line /usr/local/go/src/net/http/request.go:588
		_go_fuzz_dep_.CoverTab[41589]++
//line /usr/local/go/src/net/http/request.go:588
		// _ = "end of CoverTab[41589]"
//line /usr/local/go/src/net/http/request.go:588
	}
//line /usr/local/go/src/net/http/request.go:588
	// _ = "end of CoverTab[41551]"
//line /usr/local/go/src/net/http/request.go:588
	_go_fuzz_dep_.CoverTab[41552]++
							if !httpguts.ValidHostHeader(host) {
//line /usr/local/go/src/net/http/request.go:589
		_go_fuzz_dep_.CoverTab[41590]++
								return errors.New("http: invalid Host header")
//line /usr/local/go/src/net/http/request.go:590
		// _ = "end of CoverTab[41590]"
	} else {
//line /usr/local/go/src/net/http/request.go:591
		_go_fuzz_dep_.CoverTab[41591]++
//line /usr/local/go/src/net/http/request.go:591
		// _ = "end of CoverTab[41591]"
//line /usr/local/go/src/net/http/request.go:591
	}
//line /usr/local/go/src/net/http/request.go:591
	// _ = "end of CoverTab[41552]"
//line /usr/local/go/src/net/http/request.go:591
	_go_fuzz_dep_.CoverTab[41553]++

//line /usr/local/go/src/net/http/request.go:596
	host = removeZone(host)

	ruri := r.URL.RequestURI()
	if usingProxy && func() bool {
//line /usr/local/go/src/net/http/request.go:599
		_go_fuzz_dep_.CoverTab[41592]++
//line /usr/local/go/src/net/http/request.go:599
		return r.URL.Scheme != ""
//line /usr/local/go/src/net/http/request.go:599
		// _ = "end of CoverTab[41592]"
//line /usr/local/go/src/net/http/request.go:599
	}() && func() bool {
//line /usr/local/go/src/net/http/request.go:599
		_go_fuzz_dep_.CoverTab[41593]++
//line /usr/local/go/src/net/http/request.go:599
		return r.URL.Opaque == ""
//line /usr/local/go/src/net/http/request.go:599
		// _ = "end of CoverTab[41593]"
//line /usr/local/go/src/net/http/request.go:599
	}() {
//line /usr/local/go/src/net/http/request.go:599
		_go_fuzz_dep_.CoverTab[41594]++
								ruri = r.URL.Scheme + "://" + host + ruri
//line /usr/local/go/src/net/http/request.go:600
		// _ = "end of CoverTab[41594]"
	} else {
//line /usr/local/go/src/net/http/request.go:601
		_go_fuzz_dep_.CoverTab[41595]++
//line /usr/local/go/src/net/http/request.go:601
		if r.Method == "CONNECT" && func() bool {
//line /usr/local/go/src/net/http/request.go:601
			_go_fuzz_dep_.CoverTab[41596]++
//line /usr/local/go/src/net/http/request.go:601
			return r.URL.Path == ""
//line /usr/local/go/src/net/http/request.go:601
			// _ = "end of CoverTab[41596]"
//line /usr/local/go/src/net/http/request.go:601
		}() {
//line /usr/local/go/src/net/http/request.go:601
			_go_fuzz_dep_.CoverTab[41597]++

									ruri = host
									if r.URL.Opaque != "" {
//line /usr/local/go/src/net/http/request.go:604
				_go_fuzz_dep_.CoverTab[41598]++
										ruri = r.URL.Opaque
//line /usr/local/go/src/net/http/request.go:605
				// _ = "end of CoverTab[41598]"
			} else {
//line /usr/local/go/src/net/http/request.go:606
				_go_fuzz_dep_.CoverTab[41599]++
//line /usr/local/go/src/net/http/request.go:606
				// _ = "end of CoverTab[41599]"
//line /usr/local/go/src/net/http/request.go:606
			}
//line /usr/local/go/src/net/http/request.go:606
			// _ = "end of CoverTab[41597]"
		} else {
//line /usr/local/go/src/net/http/request.go:607
			_go_fuzz_dep_.CoverTab[41600]++
//line /usr/local/go/src/net/http/request.go:607
			// _ = "end of CoverTab[41600]"
//line /usr/local/go/src/net/http/request.go:607
		}
//line /usr/local/go/src/net/http/request.go:607
		// _ = "end of CoverTab[41595]"
//line /usr/local/go/src/net/http/request.go:607
	}
//line /usr/local/go/src/net/http/request.go:607
	// _ = "end of CoverTab[41553]"
//line /usr/local/go/src/net/http/request.go:607
	_go_fuzz_dep_.CoverTab[41554]++
							if stringContainsCTLByte(ruri) {
//line /usr/local/go/src/net/http/request.go:608
		_go_fuzz_dep_.CoverTab[41601]++
								return errors.New("net/http: can't write control character in Request.URL")
//line /usr/local/go/src/net/http/request.go:609
		// _ = "end of CoverTab[41601]"
	} else {
//line /usr/local/go/src/net/http/request.go:610
		_go_fuzz_dep_.CoverTab[41602]++
//line /usr/local/go/src/net/http/request.go:610
		// _ = "end of CoverTab[41602]"
//line /usr/local/go/src/net/http/request.go:610
	}
//line /usr/local/go/src/net/http/request.go:610
	// _ = "end of CoverTab[41554]"
//line /usr/local/go/src/net/http/request.go:610
	_go_fuzz_dep_.CoverTab[41555]++

//line /usr/local/go/src/net/http/request.go:615
	// Wrap the writer in a bufio Writer if it's not already buffered.
	// Don't always call NewWriter, as that forces a bytes.Buffer
	// and other small bufio Writers to have a minimum 4k buffer
	// size.
	var bw *bufio.Writer
	if _, ok := w.(io.ByteWriter); !ok {
//line /usr/local/go/src/net/http/request.go:620
		_go_fuzz_dep_.CoverTab[41603]++
								bw = bufio.NewWriter(w)
								w = bw
//line /usr/local/go/src/net/http/request.go:622
		// _ = "end of CoverTab[41603]"
	} else {
//line /usr/local/go/src/net/http/request.go:623
		_go_fuzz_dep_.CoverTab[41604]++
//line /usr/local/go/src/net/http/request.go:623
		// _ = "end of CoverTab[41604]"
//line /usr/local/go/src/net/http/request.go:623
	}
//line /usr/local/go/src/net/http/request.go:623
	// _ = "end of CoverTab[41555]"
//line /usr/local/go/src/net/http/request.go:623
	_go_fuzz_dep_.CoverTab[41556]++

							_, err = fmt.Fprintf(w, "%s %s HTTP/1.1\r\n", valueOrDefault(r.Method, "GET"), ruri)
							if err != nil {
//line /usr/local/go/src/net/http/request.go:626
		_go_fuzz_dep_.CoverTab[41605]++
								return err
//line /usr/local/go/src/net/http/request.go:627
		// _ = "end of CoverTab[41605]"
	} else {
//line /usr/local/go/src/net/http/request.go:628
		_go_fuzz_dep_.CoverTab[41606]++
//line /usr/local/go/src/net/http/request.go:628
		// _ = "end of CoverTab[41606]"
//line /usr/local/go/src/net/http/request.go:628
	}
//line /usr/local/go/src/net/http/request.go:628
	// _ = "end of CoverTab[41556]"
//line /usr/local/go/src/net/http/request.go:628
	_go_fuzz_dep_.CoverTab[41557]++

//line /usr/local/go/src/net/http/request.go:631
	_, err = fmt.Fprintf(w, "Host: %s\r\n", host)
	if err != nil {
//line /usr/local/go/src/net/http/request.go:632
		_go_fuzz_dep_.CoverTab[41607]++
								return err
//line /usr/local/go/src/net/http/request.go:633
		// _ = "end of CoverTab[41607]"
	} else {
//line /usr/local/go/src/net/http/request.go:634
		_go_fuzz_dep_.CoverTab[41608]++
//line /usr/local/go/src/net/http/request.go:634
		// _ = "end of CoverTab[41608]"
//line /usr/local/go/src/net/http/request.go:634
	}
//line /usr/local/go/src/net/http/request.go:634
	// _ = "end of CoverTab[41557]"
//line /usr/local/go/src/net/http/request.go:634
	_go_fuzz_dep_.CoverTab[41558]++
							if trace != nil && func() bool {
//line /usr/local/go/src/net/http/request.go:635
		_go_fuzz_dep_.CoverTab[41609]++
//line /usr/local/go/src/net/http/request.go:635
		return trace.WroteHeaderField != nil
//line /usr/local/go/src/net/http/request.go:635
		// _ = "end of CoverTab[41609]"
//line /usr/local/go/src/net/http/request.go:635
	}() {
//line /usr/local/go/src/net/http/request.go:635
		_go_fuzz_dep_.CoverTab[41610]++
								trace.WroteHeaderField("Host", []string{host})
//line /usr/local/go/src/net/http/request.go:636
		// _ = "end of CoverTab[41610]"
	} else {
//line /usr/local/go/src/net/http/request.go:637
		_go_fuzz_dep_.CoverTab[41611]++
//line /usr/local/go/src/net/http/request.go:637
		// _ = "end of CoverTab[41611]"
//line /usr/local/go/src/net/http/request.go:637
	}
//line /usr/local/go/src/net/http/request.go:637
	// _ = "end of CoverTab[41558]"
//line /usr/local/go/src/net/http/request.go:637
	_go_fuzz_dep_.CoverTab[41559]++

//line /usr/local/go/src/net/http/request.go:641
	userAgent := defaultUserAgent
	if r.Header.has("User-Agent") {
//line /usr/local/go/src/net/http/request.go:642
		_go_fuzz_dep_.CoverTab[41612]++
								userAgent = r.Header.Get("User-Agent")
//line /usr/local/go/src/net/http/request.go:643
		// _ = "end of CoverTab[41612]"
	} else {
//line /usr/local/go/src/net/http/request.go:644
		_go_fuzz_dep_.CoverTab[41613]++
//line /usr/local/go/src/net/http/request.go:644
		// _ = "end of CoverTab[41613]"
//line /usr/local/go/src/net/http/request.go:644
	}
//line /usr/local/go/src/net/http/request.go:644
	// _ = "end of CoverTab[41559]"
//line /usr/local/go/src/net/http/request.go:644
	_go_fuzz_dep_.CoverTab[41560]++
							if userAgent != "" {
//line /usr/local/go/src/net/http/request.go:645
		_go_fuzz_dep_.CoverTab[41614]++
								_, err = fmt.Fprintf(w, "User-Agent: %s\r\n", userAgent)
								if err != nil {
//line /usr/local/go/src/net/http/request.go:647
			_go_fuzz_dep_.CoverTab[41616]++
									return err
//line /usr/local/go/src/net/http/request.go:648
			// _ = "end of CoverTab[41616]"
		} else {
//line /usr/local/go/src/net/http/request.go:649
			_go_fuzz_dep_.CoverTab[41617]++
//line /usr/local/go/src/net/http/request.go:649
			// _ = "end of CoverTab[41617]"
//line /usr/local/go/src/net/http/request.go:649
		}
//line /usr/local/go/src/net/http/request.go:649
		// _ = "end of CoverTab[41614]"
//line /usr/local/go/src/net/http/request.go:649
		_go_fuzz_dep_.CoverTab[41615]++
								if trace != nil && func() bool {
//line /usr/local/go/src/net/http/request.go:650
			_go_fuzz_dep_.CoverTab[41618]++
//line /usr/local/go/src/net/http/request.go:650
			return trace.WroteHeaderField != nil
//line /usr/local/go/src/net/http/request.go:650
			// _ = "end of CoverTab[41618]"
//line /usr/local/go/src/net/http/request.go:650
		}() {
//line /usr/local/go/src/net/http/request.go:650
			_go_fuzz_dep_.CoverTab[41619]++
									trace.WroteHeaderField("User-Agent", []string{userAgent})
//line /usr/local/go/src/net/http/request.go:651
			// _ = "end of CoverTab[41619]"
		} else {
//line /usr/local/go/src/net/http/request.go:652
			_go_fuzz_dep_.CoverTab[41620]++
//line /usr/local/go/src/net/http/request.go:652
			// _ = "end of CoverTab[41620]"
//line /usr/local/go/src/net/http/request.go:652
		}
//line /usr/local/go/src/net/http/request.go:652
		// _ = "end of CoverTab[41615]"
	} else {
//line /usr/local/go/src/net/http/request.go:653
		_go_fuzz_dep_.CoverTab[41621]++
//line /usr/local/go/src/net/http/request.go:653
		// _ = "end of CoverTab[41621]"
//line /usr/local/go/src/net/http/request.go:653
	}
//line /usr/local/go/src/net/http/request.go:653
	// _ = "end of CoverTab[41560]"
//line /usr/local/go/src/net/http/request.go:653
	_go_fuzz_dep_.CoverTab[41561]++

//line /usr/local/go/src/net/http/request.go:656
	tw, err := newTransferWriter(r)
	if err != nil {
//line /usr/local/go/src/net/http/request.go:657
		_go_fuzz_dep_.CoverTab[41622]++
								return err
//line /usr/local/go/src/net/http/request.go:658
		// _ = "end of CoverTab[41622]"
	} else {
//line /usr/local/go/src/net/http/request.go:659
		_go_fuzz_dep_.CoverTab[41623]++
//line /usr/local/go/src/net/http/request.go:659
		// _ = "end of CoverTab[41623]"
//line /usr/local/go/src/net/http/request.go:659
	}
//line /usr/local/go/src/net/http/request.go:659
	// _ = "end of CoverTab[41561]"
//line /usr/local/go/src/net/http/request.go:659
	_go_fuzz_dep_.CoverTab[41562]++
							err = tw.writeHeader(w, trace)
							if err != nil {
//line /usr/local/go/src/net/http/request.go:661
		_go_fuzz_dep_.CoverTab[41624]++
								return err
//line /usr/local/go/src/net/http/request.go:662
		// _ = "end of CoverTab[41624]"
	} else {
//line /usr/local/go/src/net/http/request.go:663
		_go_fuzz_dep_.CoverTab[41625]++
//line /usr/local/go/src/net/http/request.go:663
		// _ = "end of CoverTab[41625]"
//line /usr/local/go/src/net/http/request.go:663
	}
//line /usr/local/go/src/net/http/request.go:663
	// _ = "end of CoverTab[41562]"
//line /usr/local/go/src/net/http/request.go:663
	_go_fuzz_dep_.CoverTab[41563]++

							err = r.Header.writeSubset(w, reqWriteExcludeHeader, trace)
							if err != nil {
//line /usr/local/go/src/net/http/request.go:666
		_go_fuzz_dep_.CoverTab[41626]++
								return err
//line /usr/local/go/src/net/http/request.go:667
		// _ = "end of CoverTab[41626]"
	} else {
//line /usr/local/go/src/net/http/request.go:668
		_go_fuzz_dep_.CoverTab[41627]++
//line /usr/local/go/src/net/http/request.go:668
		// _ = "end of CoverTab[41627]"
//line /usr/local/go/src/net/http/request.go:668
	}
//line /usr/local/go/src/net/http/request.go:668
	// _ = "end of CoverTab[41563]"
//line /usr/local/go/src/net/http/request.go:668
	_go_fuzz_dep_.CoverTab[41564]++

							if extraHeaders != nil {
//line /usr/local/go/src/net/http/request.go:670
		_go_fuzz_dep_.CoverTab[41628]++
								err = extraHeaders.write(w, trace)
								if err != nil {
//line /usr/local/go/src/net/http/request.go:672
			_go_fuzz_dep_.CoverTab[41629]++
									return err
//line /usr/local/go/src/net/http/request.go:673
			// _ = "end of CoverTab[41629]"
		} else {
//line /usr/local/go/src/net/http/request.go:674
			_go_fuzz_dep_.CoverTab[41630]++
//line /usr/local/go/src/net/http/request.go:674
			// _ = "end of CoverTab[41630]"
//line /usr/local/go/src/net/http/request.go:674
		}
//line /usr/local/go/src/net/http/request.go:674
		// _ = "end of CoverTab[41628]"
	} else {
//line /usr/local/go/src/net/http/request.go:675
		_go_fuzz_dep_.CoverTab[41631]++
//line /usr/local/go/src/net/http/request.go:675
		// _ = "end of CoverTab[41631]"
//line /usr/local/go/src/net/http/request.go:675
	}
//line /usr/local/go/src/net/http/request.go:675
	// _ = "end of CoverTab[41564]"
//line /usr/local/go/src/net/http/request.go:675
	_go_fuzz_dep_.CoverTab[41565]++

							_, err = io.WriteString(w, "\r\n")
							if err != nil {
//line /usr/local/go/src/net/http/request.go:678
		_go_fuzz_dep_.CoverTab[41632]++
								return err
//line /usr/local/go/src/net/http/request.go:679
		// _ = "end of CoverTab[41632]"
	} else {
//line /usr/local/go/src/net/http/request.go:680
		_go_fuzz_dep_.CoverTab[41633]++
//line /usr/local/go/src/net/http/request.go:680
		// _ = "end of CoverTab[41633]"
//line /usr/local/go/src/net/http/request.go:680
	}
//line /usr/local/go/src/net/http/request.go:680
	// _ = "end of CoverTab[41565]"
//line /usr/local/go/src/net/http/request.go:680
	_go_fuzz_dep_.CoverTab[41566]++

							if trace != nil && func() bool {
//line /usr/local/go/src/net/http/request.go:682
		_go_fuzz_dep_.CoverTab[41634]++
//line /usr/local/go/src/net/http/request.go:682
		return trace.WroteHeaders != nil
//line /usr/local/go/src/net/http/request.go:682
		// _ = "end of CoverTab[41634]"
//line /usr/local/go/src/net/http/request.go:682
	}() {
//line /usr/local/go/src/net/http/request.go:682
		_go_fuzz_dep_.CoverTab[41635]++
								trace.WroteHeaders()
//line /usr/local/go/src/net/http/request.go:683
		// _ = "end of CoverTab[41635]"
	} else {
//line /usr/local/go/src/net/http/request.go:684
		_go_fuzz_dep_.CoverTab[41636]++
//line /usr/local/go/src/net/http/request.go:684
		// _ = "end of CoverTab[41636]"
//line /usr/local/go/src/net/http/request.go:684
	}
//line /usr/local/go/src/net/http/request.go:684
	// _ = "end of CoverTab[41566]"
//line /usr/local/go/src/net/http/request.go:684
	_go_fuzz_dep_.CoverTab[41567]++

//line /usr/local/go/src/net/http/request.go:687
	if waitForContinue != nil {
//line /usr/local/go/src/net/http/request.go:687
		_go_fuzz_dep_.CoverTab[41637]++
								if bw, ok := w.(*bufio.Writer); ok {
//line /usr/local/go/src/net/http/request.go:688
			_go_fuzz_dep_.CoverTab[41640]++
									err = bw.Flush()
									if err != nil {
//line /usr/local/go/src/net/http/request.go:690
				_go_fuzz_dep_.CoverTab[41641]++
										return err
//line /usr/local/go/src/net/http/request.go:691
				// _ = "end of CoverTab[41641]"
			} else {
//line /usr/local/go/src/net/http/request.go:692
				_go_fuzz_dep_.CoverTab[41642]++
//line /usr/local/go/src/net/http/request.go:692
				// _ = "end of CoverTab[41642]"
//line /usr/local/go/src/net/http/request.go:692
			}
//line /usr/local/go/src/net/http/request.go:692
			// _ = "end of CoverTab[41640]"
		} else {
//line /usr/local/go/src/net/http/request.go:693
			_go_fuzz_dep_.CoverTab[41643]++
//line /usr/local/go/src/net/http/request.go:693
			// _ = "end of CoverTab[41643]"
//line /usr/local/go/src/net/http/request.go:693
		}
//line /usr/local/go/src/net/http/request.go:693
		// _ = "end of CoverTab[41637]"
//line /usr/local/go/src/net/http/request.go:693
		_go_fuzz_dep_.CoverTab[41638]++
								if trace != nil && func() bool {
//line /usr/local/go/src/net/http/request.go:694
			_go_fuzz_dep_.CoverTab[41644]++
//line /usr/local/go/src/net/http/request.go:694
			return trace.Wait100Continue != nil
//line /usr/local/go/src/net/http/request.go:694
			// _ = "end of CoverTab[41644]"
//line /usr/local/go/src/net/http/request.go:694
		}() {
//line /usr/local/go/src/net/http/request.go:694
			_go_fuzz_dep_.CoverTab[41645]++
									trace.Wait100Continue()
//line /usr/local/go/src/net/http/request.go:695
			// _ = "end of CoverTab[41645]"
		} else {
//line /usr/local/go/src/net/http/request.go:696
			_go_fuzz_dep_.CoverTab[41646]++
//line /usr/local/go/src/net/http/request.go:696
			// _ = "end of CoverTab[41646]"
//line /usr/local/go/src/net/http/request.go:696
		}
//line /usr/local/go/src/net/http/request.go:696
		// _ = "end of CoverTab[41638]"
//line /usr/local/go/src/net/http/request.go:696
		_go_fuzz_dep_.CoverTab[41639]++
								if !waitForContinue() {
//line /usr/local/go/src/net/http/request.go:697
			_go_fuzz_dep_.CoverTab[41647]++
									closed = true
									r.closeBody()
									return nil
//line /usr/local/go/src/net/http/request.go:700
			// _ = "end of CoverTab[41647]"
		} else {
//line /usr/local/go/src/net/http/request.go:701
			_go_fuzz_dep_.CoverTab[41648]++
//line /usr/local/go/src/net/http/request.go:701
			// _ = "end of CoverTab[41648]"
//line /usr/local/go/src/net/http/request.go:701
		}
//line /usr/local/go/src/net/http/request.go:701
		// _ = "end of CoverTab[41639]"
	} else {
//line /usr/local/go/src/net/http/request.go:702
		_go_fuzz_dep_.CoverTab[41649]++
//line /usr/local/go/src/net/http/request.go:702
		// _ = "end of CoverTab[41649]"
//line /usr/local/go/src/net/http/request.go:702
	}
//line /usr/local/go/src/net/http/request.go:702
	// _ = "end of CoverTab[41567]"
//line /usr/local/go/src/net/http/request.go:702
	_go_fuzz_dep_.CoverTab[41568]++

							if bw, ok := w.(*bufio.Writer); ok && func() bool {
//line /usr/local/go/src/net/http/request.go:704
		_go_fuzz_dep_.CoverTab[41650]++
//line /usr/local/go/src/net/http/request.go:704
		return tw.FlushHeaders
//line /usr/local/go/src/net/http/request.go:704
		// _ = "end of CoverTab[41650]"
//line /usr/local/go/src/net/http/request.go:704
	}() {
//line /usr/local/go/src/net/http/request.go:704
		_go_fuzz_dep_.CoverTab[41651]++
								if err := bw.Flush(); err != nil {
//line /usr/local/go/src/net/http/request.go:705
			_go_fuzz_dep_.CoverTab[41652]++
									return err
//line /usr/local/go/src/net/http/request.go:706
			// _ = "end of CoverTab[41652]"
		} else {
//line /usr/local/go/src/net/http/request.go:707
			_go_fuzz_dep_.CoverTab[41653]++
//line /usr/local/go/src/net/http/request.go:707
			// _ = "end of CoverTab[41653]"
//line /usr/local/go/src/net/http/request.go:707
		}
//line /usr/local/go/src/net/http/request.go:707
		// _ = "end of CoverTab[41651]"
	} else {
//line /usr/local/go/src/net/http/request.go:708
		_go_fuzz_dep_.CoverTab[41654]++
//line /usr/local/go/src/net/http/request.go:708
		// _ = "end of CoverTab[41654]"
//line /usr/local/go/src/net/http/request.go:708
	}
//line /usr/local/go/src/net/http/request.go:708
	// _ = "end of CoverTab[41568]"
//line /usr/local/go/src/net/http/request.go:708
	_go_fuzz_dep_.CoverTab[41569]++

//line /usr/local/go/src/net/http/request.go:711
	closed = true
	err = tw.writeBody(w)
	if err != nil {
//line /usr/local/go/src/net/http/request.go:713
		_go_fuzz_dep_.CoverTab[41655]++
								if tw.bodyReadError == err {
//line /usr/local/go/src/net/http/request.go:714
			_go_fuzz_dep_.CoverTab[41657]++
									err = requestBodyReadError{err}
//line /usr/local/go/src/net/http/request.go:715
			// _ = "end of CoverTab[41657]"
		} else {
//line /usr/local/go/src/net/http/request.go:716
			_go_fuzz_dep_.CoverTab[41658]++
//line /usr/local/go/src/net/http/request.go:716
			// _ = "end of CoverTab[41658]"
//line /usr/local/go/src/net/http/request.go:716
		}
//line /usr/local/go/src/net/http/request.go:716
		// _ = "end of CoverTab[41655]"
//line /usr/local/go/src/net/http/request.go:716
		_go_fuzz_dep_.CoverTab[41656]++
								return err
//line /usr/local/go/src/net/http/request.go:717
		// _ = "end of CoverTab[41656]"
	} else {
//line /usr/local/go/src/net/http/request.go:718
		_go_fuzz_dep_.CoverTab[41659]++
//line /usr/local/go/src/net/http/request.go:718
		// _ = "end of CoverTab[41659]"
//line /usr/local/go/src/net/http/request.go:718
	}
//line /usr/local/go/src/net/http/request.go:718
	// _ = "end of CoverTab[41569]"
//line /usr/local/go/src/net/http/request.go:718
	_go_fuzz_dep_.CoverTab[41570]++

							if bw != nil {
//line /usr/local/go/src/net/http/request.go:720
		_go_fuzz_dep_.CoverTab[41660]++
								return bw.Flush()
//line /usr/local/go/src/net/http/request.go:721
		// _ = "end of CoverTab[41660]"
	} else {
//line /usr/local/go/src/net/http/request.go:722
		_go_fuzz_dep_.CoverTab[41661]++
//line /usr/local/go/src/net/http/request.go:722
		// _ = "end of CoverTab[41661]"
//line /usr/local/go/src/net/http/request.go:722
	}
//line /usr/local/go/src/net/http/request.go:722
	// _ = "end of CoverTab[41570]"
//line /usr/local/go/src/net/http/request.go:722
	_go_fuzz_dep_.CoverTab[41571]++
							return nil
//line /usr/local/go/src/net/http/request.go:723
	// _ = "end of CoverTab[41571]"
}

// requestBodyReadError wraps an error from (*Request).write to indicate
//line /usr/local/go/src/net/http/request.go:726
// that the error came from a Read call on the Request.Body.
//line /usr/local/go/src/net/http/request.go:726
// This error type should not escape the net/http package to users.
//line /usr/local/go/src/net/http/request.go:729
type requestBodyReadError struct{ error }

func idnaASCII(v string) (string, error) {
//line /usr/local/go/src/net/http/request.go:731
	_go_fuzz_dep_.CoverTab[41662]++

//line /usr/local/go/src/net/http/request.go:741
	if ascii.Is(v) {
//line /usr/local/go/src/net/http/request.go:741
		_go_fuzz_dep_.CoverTab[41664]++
								return v, nil
//line /usr/local/go/src/net/http/request.go:742
		// _ = "end of CoverTab[41664]"
	} else {
//line /usr/local/go/src/net/http/request.go:743
		_go_fuzz_dep_.CoverTab[41665]++
//line /usr/local/go/src/net/http/request.go:743
		// _ = "end of CoverTab[41665]"
//line /usr/local/go/src/net/http/request.go:743
	}
//line /usr/local/go/src/net/http/request.go:743
	// _ = "end of CoverTab[41662]"
//line /usr/local/go/src/net/http/request.go:743
	_go_fuzz_dep_.CoverTab[41663]++
							return idna.Lookup.ToASCII(v)
//line /usr/local/go/src/net/http/request.go:744
	// _ = "end of CoverTab[41663]"
}

// removeZone removes IPv6 zone identifier from host.
//line /usr/local/go/src/net/http/request.go:747
// E.g., "[fe80::1%en0]:8080" to "[fe80::1]:8080"
//line /usr/local/go/src/net/http/request.go:749
func removeZone(host string) string {
//line /usr/local/go/src/net/http/request.go:749
	_go_fuzz_dep_.CoverTab[41666]++
							if !strings.HasPrefix(host, "[") {
//line /usr/local/go/src/net/http/request.go:750
		_go_fuzz_dep_.CoverTab[41670]++
								return host
//line /usr/local/go/src/net/http/request.go:751
		// _ = "end of CoverTab[41670]"
	} else {
//line /usr/local/go/src/net/http/request.go:752
		_go_fuzz_dep_.CoverTab[41671]++
//line /usr/local/go/src/net/http/request.go:752
		// _ = "end of CoverTab[41671]"
//line /usr/local/go/src/net/http/request.go:752
	}
//line /usr/local/go/src/net/http/request.go:752
	// _ = "end of CoverTab[41666]"
//line /usr/local/go/src/net/http/request.go:752
	_go_fuzz_dep_.CoverTab[41667]++
							i := strings.LastIndex(host, "]")
							if i < 0 {
//line /usr/local/go/src/net/http/request.go:754
		_go_fuzz_dep_.CoverTab[41672]++
								return host
//line /usr/local/go/src/net/http/request.go:755
		// _ = "end of CoverTab[41672]"
	} else {
//line /usr/local/go/src/net/http/request.go:756
		_go_fuzz_dep_.CoverTab[41673]++
//line /usr/local/go/src/net/http/request.go:756
		// _ = "end of CoverTab[41673]"
//line /usr/local/go/src/net/http/request.go:756
	}
//line /usr/local/go/src/net/http/request.go:756
	// _ = "end of CoverTab[41667]"
//line /usr/local/go/src/net/http/request.go:756
	_go_fuzz_dep_.CoverTab[41668]++
							j := strings.LastIndex(host[:i], "%")
							if j < 0 {
//line /usr/local/go/src/net/http/request.go:758
		_go_fuzz_dep_.CoverTab[41674]++
								return host
//line /usr/local/go/src/net/http/request.go:759
		// _ = "end of CoverTab[41674]"
	} else {
//line /usr/local/go/src/net/http/request.go:760
		_go_fuzz_dep_.CoverTab[41675]++
//line /usr/local/go/src/net/http/request.go:760
		// _ = "end of CoverTab[41675]"
//line /usr/local/go/src/net/http/request.go:760
	}
//line /usr/local/go/src/net/http/request.go:760
	// _ = "end of CoverTab[41668]"
//line /usr/local/go/src/net/http/request.go:760
	_go_fuzz_dep_.CoverTab[41669]++
							return host[:j] + host[i:]
//line /usr/local/go/src/net/http/request.go:761
	// _ = "end of CoverTab[41669]"
}

// ParseHTTPVersion parses an HTTP version string according to RFC 7230, section 2.6.
//line /usr/local/go/src/net/http/request.go:764
// "HTTP/1.0" returns (1, 0, true). Note that strings without
//line /usr/local/go/src/net/http/request.go:764
// a minor version, such as "HTTP/2", are not valid.
//line /usr/local/go/src/net/http/request.go:767
func ParseHTTPVersion(vers string) (major, minor int, ok bool) {
//line /usr/local/go/src/net/http/request.go:767
	_go_fuzz_dep_.CoverTab[41676]++
							switch vers {
	case "HTTP/1.1":
//line /usr/local/go/src/net/http/request.go:769
		_go_fuzz_dep_.CoverTab[41683]++
								return 1, 1, true
//line /usr/local/go/src/net/http/request.go:770
		// _ = "end of CoverTab[41683]"
	case "HTTP/1.0":
//line /usr/local/go/src/net/http/request.go:771
		_go_fuzz_dep_.CoverTab[41684]++
								return 1, 0, true
//line /usr/local/go/src/net/http/request.go:772
		// _ = "end of CoverTab[41684]"
//line /usr/local/go/src/net/http/request.go:772
	default:
//line /usr/local/go/src/net/http/request.go:772
		_go_fuzz_dep_.CoverTab[41685]++
//line /usr/local/go/src/net/http/request.go:772
		// _ = "end of CoverTab[41685]"
	}
//line /usr/local/go/src/net/http/request.go:773
	// _ = "end of CoverTab[41676]"
//line /usr/local/go/src/net/http/request.go:773
	_go_fuzz_dep_.CoverTab[41677]++
							if !strings.HasPrefix(vers, "HTTP/") {
//line /usr/local/go/src/net/http/request.go:774
		_go_fuzz_dep_.CoverTab[41686]++
								return 0, 0, false
//line /usr/local/go/src/net/http/request.go:775
		// _ = "end of CoverTab[41686]"
	} else {
//line /usr/local/go/src/net/http/request.go:776
		_go_fuzz_dep_.CoverTab[41687]++
//line /usr/local/go/src/net/http/request.go:776
		// _ = "end of CoverTab[41687]"
//line /usr/local/go/src/net/http/request.go:776
	}
//line /usr/local/go/src/net/http/request.go:776
	// _ = "end of CoverTab[41677]"
//line /usr/local/go/src/net/http/request.go:776
	_go_fuzz_dep_.CoverTab[41678]++
							if len(vers) != len("HTTP/X.Y") {
//line /usr/local/go/src/net/http/request.go:777
		_go_fuzz_dep_.CoverTab[41688]++
								return 0, 0, false
//line /usr/local/go/src/net/http/request.go:778
		// _ = "end of CoverTab[41688]"
	} else {
//line /usr/local/go/src/net/http/request.go:779
		_go_fuzz_dep_.CoverTab[41689]++
//line /usr/local/go/src/net/http/request.go:779
		// _ = "end of CoverTab[41689]"
//line /usr/local/go/src/net/http/request.go:779
	}
//line /usr/local/go/src/net/http/request.go:779
	// _ = "end of CoverTab[41678]"
//line /usr/local/go/src/net/http/request.go:779
	_go_fuzz_dep_.CoverTab[41679]++
							if vers[6] != '.' {
//line /usr/local/go/src/net/http/request.go:780
		_go_fuzz_dep_.CoverTab[41690]++
								return 0, 0, false
//line /usr/local/go/src/net/http/request.go:781
		// _ = "end of CoverTab[41690]"
	} else {
//line /usr/local/go/src/net/http/request.go:782
		_go_fuzz_dep_.CoverTab[41691]++
//line /usr/local/go/src/net/http/request.go:782
		// _ = "end of CoverTab[41691]"
//line /usr/local/go/src/net/http/request.go:782
	}
//line /usr/local/go/src/net/http/request.go:782
	// _ = "end of CoverTab[41679]"
//line /usr/local/go/src/net/http/request.go:782
	_go_fuzz_dep_.CoverTab[41680]++
							maj, err := strconv.ParseUint(vers[5:6], 10, 0)
							if err != nil {
//line /usr/local/go/src/net/http/request.go:784
		_go_fuzz_dep_.CoverTab[41692]++
								return 0, 0, false
//line /usr/local/go/src/net/http/request.go:785
		// _ = "end of CoverTab[41692]"
	} else {
//line /usr/local/go/src/net/http/request.go:786
		_go_fuzz_dep_.CoverTab[41693]++
//line /usr/local/go/src/net/http/request.go:786
		// _ = "end of CoverTab[41693]"
//line /usr/local/go/src/net/http/request.go:786
	}
//line /usr/local/go/src/net/http/request.go:786
	// _ = "end of CoverTab[41680]"
//line /usr/local/go/src/net/http/request.go:786
	_go_fuzz_dep_.CoverTab[41681]++
							min, err := strconv.ParseUint(vers[7:8], 10, 0)
							if err != nil {
//line /usr/local/go/src/net/http/request.go:788
		_go_fuzz_dep_.CoverTab[41694]++
								return 0, 0, false
//line /usr/local/go/src/net/http/request.go:789
		// _ = "end of CoverTab[41694]"
	} else {
//line /usr/local/go/src/net/http/request.go:790
		_go_fuzz_dep_.CoverTab[41695]++
//line /usr/local/go/src/net/http/request.go:790
		// _ = "end of CoverTab[41695]"
//line /usr/local/go/src/net/http/request.go:790
	}
//line /usr/local/go/src/net/http/request.go:790
	// _ = "end of CoverTab[41681]"
//line /usr/local/go/src/net/http/request.go:790
	_go_fuzz_dep_.CoverTab[41682]++
							return int(maj), int(min), true
//line /usr/local/go/src/net/http/request.go:791
	// _ = "end of CoverTab[41682]"
}

func validMethod(method string) bool {
//line /usr/local/go/src/net/http/request.go:794
	_go_fuzz_dep_.CoverTab[41696]++

//line /usr/local/go/src/net/http/request.go:808
	return len(method) > 0 && func() bool {
//line /usr/local/go/src/net/http/request.go:808
		_go_fuzz_dep_.CoverTab[41697]++
//line /usr/local/go/src/net/http/request.go:808
		return strings.IndexFunc(method, isNotToken) == -1
//line /usr/local/go/src/net/http/request.go:808
		// _ = "end of CoverTab[41697]"
//line /usr/local/go/src/net/http/request.go:808
	}()
//line /usr/local/go/src/net/http/request.go:808
	// _ = "end of CoverTab[41696]"
}

// NewRequest wraps NewRequestWithContext using context.Background.
func NewRequest(method, url string, body io.Reader) (*Request, error) {
//line /usr/local/go/src/net/http/request.go:812
	_go_fuzz_dep_.CoverTab[41698]++
							return NewRequestWithContext(context.Background(), method, url, body)
//line /usr/local/go/src/net/http/request.go:813
	// _ = "end of CoverTab[41698]"
}

// NewRequestWithContext returns a new Request given a method, URL, and
//line /usr/local/go/src/net/http/request.go:816
// optional body.
//line /usr/local/go/src/net/http/request.go:816
//
//line /usr/local/go/src/net/http/request.go:816
// If the provided body is also an io.Closer, the returned
//line /usr/local/go/src/net/http/request.go:816
// Request.Body is set to body and will be closed by the Client
//line /usr/local/go/src/net/http/request.go:816
// methods Do, Post, and PostForm, and Transport.RoundTrip.
//line /usr/local/go/src/net/http/request.go:816
//
//line /usr/local/go/src/net/http/request.go:816
// NewRequestWithContext returns a Request suitable for use with
//line /usr/local/go/src/net/http/request.go:816
// Client.Do or Transport.RoundTrip. To create a request for use with
//line /usr/local/go/src/net/http/request.go:816
// testing a Server Handler, either use the NewRequest function in the
//line /usr/local/go/src/net/http/request.go:816
// net/http/httptest package, use ReadRequest, or manually update the
//line /usr/local/go/src/net/http/request.go:816
// Request fields. For an outgoing client request, the context
//line /usr/local/go/src/net/http/request.go:816
// controls the entire lifetime of a request and its response:
//line /usr/local/go/src/net/http/request.go:816
// obtaining a connection, sending the request, and reading the
//line /usr/local/go/src/net/http/request.go:816
// response headers and body. See the Request type's documentation for
//line /usr/local/go/src/net/http/request.go:816
// the difference between inbound and outbound request fields.
//line /usr/local/go/src/net/http/request.go:816
//
//line /usr/local/go/src/net/http/request.go:816
// If body is of type *bytes.Buffer, *bytes.Reader, or
//line /usr/local/go/src/net/http/request.go:816
// *strings.Reader, the returned request's ContentLength is set to its
//line /usr/local/go/src/net/http/request.go:816
// exact value (instead of -1), GetBody is populated (so 307 and 308
//line /usr/local/go/src/net/http/request.go:816
// redirects can replay the body), and Body is set to NoBody if the
//line /usr/local/go/src/net/http/request.go:816
// ContentLength is 0.
//line /usr/local/go/src/net/http/request.go:838
func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error) {
//line /usr/local/go/src/net/http/request.go:838
	_go_fuzz_dep_.CoverTab[41699]++
							if method == "" {
//line /usr/local/go/src/net/http/request.go:839
		_go_fuzz_dep_.CoverTab[41706]++

//line /usr/local/go/src/net/http/request.go:843
		method = "GET"
//line /usr/local/go/src/net/http/request.go:843
		// _ = "end of CoverTab[41706]"
	} else {
//line /usr/local/go/src/net/http/request.go:844
		_go_fuzz_dep_.CoverTab[41707]++
//line /usr/local/go/src/net/http/request.go:844
		// _ = "end of CoverTab[41707]"
//line /usr/local/go/src/net/http/request.go:844
	}
//line /usr/local/go/src/net/http/request.go:844
	// _ = "end of CoverTab[41699]"
//line /usr/local/go/src/net/http/request.go:844
	_go_fuzz_dep_.CoverTab[41700]++
							if !validMethod(method) {
//line /usr/local/go/src/net/http/request.go:845
		_go_fuzz_dep_.CoverTab[41708]++
								return nil, fmt.Errorf("net/http: invalid method %q", method)
//line /usr/local/go/src/net/http/request.go:846
		// _ = "end of CoverTab[41708]"
	} else {
//line /usr/local/go/src/net/http/request.go:847
		_go_fuzz_dep_.CoverTab[41709]++
//line /usr/local/go/src/net/http/request.go:847
		// _ = "end of CoverTab[41709]"
//line /usr/local/go/src/net/http/request.go:847
	}
//line /usr/local/go/src/net/http/request.go:847
	// _ = "end of CoverTab[41700]"
//line /usr/local/go/src/net/http/request.go:847
	_go_fuzz_dep_.CoverTab[41701]++
							if ctx == nil {
//line /usr/local/go/src/net/http/request.go:848
		_go_fuzz_dep_.CoverTab[41710]++
								return nil, errors.New("net/http: nil Context")
//line /usr/local/go/src/net/http/request.go:849
		// _ = "end of CoverTab[41710]"
	} else {
//line /usr/local/go/src/net/http/request.go:850
		_go_fuzz_dep_.CoverTab[41711]++
//line /usr/local/go/src/net/http/request.go:850
		// _ = "end of CoverTab[41711]"
//line /usr/local/go/src/net/http/request.go:850
	}
//line /usr/local/go/src/net/http/request.go:850
	// _ = "end of CoverTab[41701]"
//line /usr/local/go/src/net/http/request.go:850
	_go_fuzz_dep_.CoverTab[41702]++
							u, err := urlpkg.Parse(url)
							if err != nil {
//line /usr/local/go/src/net/http/request.go:852
		_go_fuzz_dep_.CoverTab[41712]++
								return nil, err
//line /usr/local/go/src/net/http/request.go:853
		// _ = "end of CoverTab[41712]"
	} else {
//line /usr/local/go/src/net/http/request.go:854
		_go_fuzz_dep_.CoverTab[41713]++
//line /usr/local/go/src/net/http/request.go:854
		// _ = "end of CoverTab[41713]"
//line /usr/local/go/src/net/http/request.go:854
	}
//line /usr/local/go/src/net/http/request.go:854
	// _ = "end of CoverTab[41702]"
//line /usr/local/go/src/net/http/request.go:854
	_go_fuzz_dep_.CoverTab[41703]++
							rc, ok := body.(io.ReadCloser)
							if !ok && func() bool {
//line /usr/local/go/src/net/http/request.go:856
		_go_fuzz_dep_.CoverTab[41714]++
//line /usr/local/go/src/net/http/request.go:856
		return body != nil
//line /usr/local/go/src/net/http/request.go:856
		// _ = "end of CoverTab[41714]"
//line /usr/local/go/src/net/http/request.go:856
	}() {
//line /usr/local/go/src/net/http/request.go:856
		_go_fuzz_dep_.CoverTab[41715]++
								rc = io.NopCloser(body)
//line /usr/local/go/src/net/http/request.go:857
		// _ = "end of CoverTab[41715]"
	} else {
//line /usr/local/go/src/net/http/request.go:858
		_go_fuzz_dep_.CoverTab[41716]++
//line /usr/local/go/src/net/http/request.go:858
		// _ = "end of CoverTab[41716]"
//line /usr/local/go/src/net/http/request.go:858
	}
//line /usr/local/go/src/net/http/request.go:858
	// _ = "end of CoverTab[41703]"
//line /usr/local/go/src/net/http/request.go:858
	_go_fuzz_dep_.CoverTab[41704]++

							u.Host = removeEmptyPort(u.Host)
							req := &Request{
		ctx:		ctx,
		Method:		method,
		URL:		u,
		Proto:		"HTTP/1.1",
		ProtoMajor:	1,
		ProtoMinor:	1,
		Header:		make(Header),
		Body:		rc,
		Host:		u.Host,
	}
	if body != nil {
//line /usr/local/go/src/net/http/request.go:872
		_go_fuzz_dep_.CoverTab[41717]++
								switch v := body.(type) {
		case *bytes.Buffer:
//line /usr/local/go/src/net/http/request.go:874
			_go_fuzz_dep_.CoverTab[41719]++
									req.ContentLength = int64(v.Len())
									buf := v.Bytes()
									req.GetBody = func() (io.ReadCloser, error) {
//line /usr/local/go/src/net/http/request.go:877
				_go_fuzz_dep_.CoverTab[41723]++
										r := bytes.NewReader(buf)
										return io.NopCloser(r), nil
//line /usr/local/go/src/net/http/request.go:879
				// _ = "end of CoverTab[41723]"
			}
//line /usr/local/go/src/net/http/request.go:880
			// _ = "end of CoverTab[41719]"
		case *bytes.Reader:
//line /usr/local/go/src/net/http/request.go:881
			_go_fuzz_dep_.CoverTab[41720]++
									req.ContentLength = int64(v.Len())
									snapshot := *v
									req.GetBody = func() (io.ReadCloser, error) {
//line /usr/local/go/src/net/http/request.go:884
				_go_fuzz_dep_.CoverTab[41724]++
										r := snapshot
										return io.NopCloser(&r), nil
//line /usr/local/go/src/net/http/request.go:886
				// _ = "end of CoverTab[41724]"
			}
//line /usr/local/go/src/net/http/request.go:887
			// _ = "end of CoverTab[41720]"
		case *strings.Reader:
//line /usr/local/go/src/net/http/request.go:888
			_go_fuzz_dep_.CoverTab[41721]++
									req.ContentLength = int64(v.Len())
									snapshot := *v
									req.GetBody = func() (io.ReadCloser, error) {
//line /usr/local/go/src/net/http/request.go:891
				_go_fuzz_dep_.CoverTab[41725]++
										r := snapshot
										return io.NopCloser(&r), nil
//line /usr/local/go/src/net/http/request.go:893
				// _ = "end of CoverTab[41725]"
			}
//line /usr/local/go/src/net/http/request.go:894
			// _ = "end of CoverTab[41721]"
		default:
//line /usr/local/go/src/net/http/request.go:895
			_go_fuzz_dep_.CoverTab[41722]++
//line /usr/local/go/src/net/http/request.go:895
			// _ = "end of CoverTab[41722]"

//line /usr/local/go/src/net/http/request.go:901
		}
//line /usr/local/go/src/net/http/request.go:901
		// _ = "end of CoverTab[41717]"
//line /usr/local/go/src/net/http/request.go:901
		_go_fuzz_dep_.CoverTab[41718]++

//line /usr/local/go/src/net/http/request.go:910
		if req.GetBody != nil && func() bool {
//line /usr/local/go/src/net/http/request.go:910
			_go_fuzz_dep_.CoverTab[41726]++
//line /usr/local/go/src/net/http/request.go:910
			return req.ContentLength == 0
//line /usr/local/go/src/net/http/request.go:910
			// _ = "end of CoverTab[41726]"
//line /usr/local/go/src/net/http/request.go:910
		}() {
//line /usr/local/go/src/net/http/request.go:910
			_go_fuzz_dep_.CoverTab[41727]++
									req.Body = NoBody
									req.GetBody = func() (io.ReadCloser, error) {
//line /usr/local/go/src/net/http/request.go:912
				_go_fuzz_dep_.CoverTab[41728]++
//line /usr/local/go/src/net/http/request.go:912
				return NoBody, nil
//line /usr/local/go/src/net/http/request.go:912
				// _ = "end of CoverTab[41728]"
//line /usr/local/go/src/net/http/request.go:912
			}
//line /usr/local/go/src/net/http/request.go:912
			// _ = "end of CoverTab[41727]"
		} else {
//line /usr/local/go/src/net/http/request.go:913
			_go_fuzz_dep_.CoverTab[41729]++
//line /usr/local/go/src/net/http/request.go:913
			// _ = "end of CoverTab[41729]"
//line /usr/local/go/src/net/http/request.go:913
		}
//line /usr/local/go/src/net/http/request.go:913
		// _ = "end of CoverTab[41718]"
	} else {
//line /usr/local/go/src/net/http/request.go:914
		_go_fuzz_dep_.CoverTab[41730]++
//line /usr/local/go/src/net/http/request.go:914
		// _ = "end of CoverTab[41730]"
//line /usr/local/go/src/net/http/request.go:914
	}
//line /usr/local/go/src/net/http/request.go:914
	// _ = "end of CoverTab[41704]"
//line /usr/local/go/src/net/http/request.go:914
	_go_fuzz_dep_.CoverTab[41705]++

							return req, nil
//line /usr/local/go/src/net/http/request.go:916
	// _ = "end of CoverTab[41705]"
}

// BasicAuth returns the username and password provided in the request's
//line /usr/local/go/src/net/http/request.go:919
// Authorization header, if the request uses HTTP Basic Authentication.
//line /usr/local/go/src/net/http/request.go:919
// See RFC 2617, Section 2.
//line /usr/local/go/src/net/http/request.go:922
func (r *Request) BasicAuth() (username, password string, ok bool) {
//line /usr/local/go/src/net/http/request.go:922
	_go_fuzz_dep_.CoverTab[41731]++
							auth := r.Header.Get("Authorization")
							if auth == "" {
//line /usr/local/go/src/net/http/request.go:924
		_go_fuzz_dep_.CoverTab[41733]++
								return "", "", false
//line /usr/local/go/src/net/http/request.go:925
		// _ = "end of CoverTab[41733]"
	} else {
//line /usr/local/go/src/net/http/request.go:926
		_go_fuzz_dep_.CoverTab[41734]++
//line /usr/local/go/src/net/http/request.go:926
		// _ = "end of CoverTab[41734]"
//line /usr/local/go/src/net/http/request.go:926
	}
//line /usr/local/go/src/net/http/request.go:926
	// _ = "end of CoverTab[41731]"
//line /usr/local/go/src/net/http/request.go:926
	_go_fuzz_dep_.CoverTab[41732]++
							return parseBasicAuth(auth)
//line /usr/local/go/src/net/http/request.go:927
	// _ = "end of CoverTab[41732]"
}

// parseBasicAuth parses an HTTP Basic Authentication string.
//line /usr/local/go/src/net/http/request.go:930
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
//line /usr/local/go/src/net/http/request.go:932
func parseBasicAuth(auth string) (username, password string, ok bool) {
//line /usr/local/go/src/net/http/request.go:932
	_go_fuzz_dep_.CoverTab[41735]++
							const prefix = "Basic "

							if len(auth) < len(prefix) || func() bool {
//line /usr/local/go/src/net/http/request.go:935
		_go_fuzz_dep_.CoverTab[41739]++
//line /usr/local/go/src/net/http/request.go:935
		return !ascii.EqualFold(auth[:len(prefix)], prefix)
//line /usr/local/go/src/net/http/request.go:935
		// _ = "end of CoverTab[41739]"
//line /usr/local/go/src/net/http/request.go:935
	}() {
//line /usr/local/go/src/net/http/request.go:935
		_go_fuzz_dep_.CoverTab[41740]++
								return "", "", false
//line /usr/local/go/src/net/http/request.go:936
		// _ = "end of CoverTab[41740]"
	} else {
//line /usr/local/go/src/net/http/request.go:937
		_go_fuzz_dep_.CoverTab[41741]++
//line /usr/local/go/src/net/http/request.go:937
		// _ = "end of CoverTab[41741]"
//line /usr/local/go/src/net/http/request.go:937
	}
//line /usr/local/go/src/net/http/request.go:937
	// _ = "end of CoverTab[41735]"
//line /usr/local/go/src/net/http/request.go:937
	_go_fuzz_dep_.CoverTab[41736]++
							c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
							if err != nil {
//line /usr/local/go/src/net/http/request.go:939
		_go_fuzz_dep_.CoverTab[41742]++
								return "", "", false
//line /usr/local/go/src/net/http/request.go:940
		// _ = "end of CoverTab[41742]"
	} else {
//line /usr/local/go/src/net/http/request.go:941
		_go_fuzz_dep_.CoverTab[41743]++
//line /usr/local/go/src/net/http/request.go:941
		// _ = "end of CoverTab[41743]"
//line /usr/local/go/src/net/http/request.go:941
	}
//line /usr/local/go/src/net/http/request.go:941
	// _ = "end of CoverTab[41736]"
//line /usr/local/go/src/net/http/request.go:941
	_go_fuzz_dep_.CoverTab[41737]++
							cs := string(c)
							username, password, ok = strings.Cut(cs, ":")
							if !ok {
//line /usr/local/go/src/net/http/request.go:944
		_go_fuzz_dep_.CoverTab[41744]++
								return "", "", false
//line /usr/local/go/src/net/http/request.go:945
		// _ = "end of CoverTab[41744]"
	} else {
//line /usr/local/go/src/net/http/request.go:946
		_go_fuzz_dep_.CoverTab[41745]++
//line /usr/local/go/src/net/http/request.go:946
		// _ = "end of CoverTab[41745]"
//line /usr/local/go/src/net/http/request.go:946
	}
//line /usr/local/go/src/net/http/request.go:946
	// _ = "end of CoverTab[41737]"
//line /usr/local/go/src/net/http/request.go:946
	_go_fuzz_dep_.CoverTab[41738]++
							return username, password, true
//line /usr/local/go/src/net/http/request.go:947
	// _ = "end of CoverTab[41738]"
}

// SetBasicAuth sets the request's Authorization header to use HTTP
//line /usr/local/go/src/net/http/request.go:950
// Basic Authentication with the provided username and password.
//line /usr/local/go/src/net/http/request.go:950
//
//line /usr/local/go/src/net/http/request.go:950
// With HTTP Basic Authentication the provided username and password
//line /usr/local/go/src/net/http/request.go:950
// are not encrypted. It should generally only be used in an HTTPS
//line /usr/local/go/src/net/http/request.go:950
// request.
//line /usr/local/go/src/net/http/request.go:950
//
//line /usr/local/go/src/net/http/request.go:950
// The username may not contain a colon. Some protocols may impose
//line /usr/local/go/src/net/http/request.go:950
// additional requirements on pre-escaping the username and
//line /usr/local/go/src/net/http/request.go:950
// password. For instance, when used with OAuth2, both arguments must
//line /usr/local/go/src/net/http/request.go:950
// be URL encoded first with url.QueryEscape.
//line /usr/local/go/src/net/http/request.go:961
func (r *Request) SetBasicAuth(username, password string) {
//line /usr/local/go/src/net/http/request.go:961
	_go_fuzz_dep_.CoverTab[41746]++
							r.Header.Set("Authorization", "Basic "+basicAuth(username, password))
//line /usr/local/go/src/net/http/request.go:962
	// _ = "end of CoverTab[41746]"
}

// parseRequestLine parses "GET /foo HTTP/1.1" into its three parts.
func parseRequestLine(line string) (method, requestURI, proto string, ok bool) {
//line /usr/local/go/src/net/http/request.go:966
	_go_fuzz_dep_.CoverTab[41747]++
							method, rest, ok1 := strings.Cut(line, " ")
							requestURI, proto, ok2 := strings.Cut(rest, " ")
							if !ok1 || func() bool {
//line /usr/local/go/src/net/http/request.go:969
		_go_fuzz_dep_.CoverTab[41749]++
//line /usr/local/go/src/net/http/request.go:969
		return !ok2
//line /usr/local/go/src/net/http/request.go:969
		// _ = "end of CoverTab[41749]"
//line /usr/local/go/src/net/http/request.go:969
	}() {
//line /usr/local/go/src/net/http/request.go:969
		_go_fuzz_dep_.CoverTab[41750]++
								return "", "", "", false
//line /usr/local/go/src/net/http/request.go:970
		// _ = "end of CoverTab[41750]"
	} else {
//line /usr/local/go/src/net/http/request.go:971
		_go_fuzz_dep_.CoverTab[41751]++
//line /usr/local/go/src/net/http/request.go:971
		// _ = "end of CoverTab[41751]"
//line /usr/local/go/src/net/http/request.go:971
	}
//line /usr/local/go/src/net/http/request.go:971
	// _ = "end of CoverTab[41747]"
//line /usr/local/go/src/net/http/request.go:971
	_go_fuzz_dep_.CoverTab[41748]++
							return method, requestURI, proto, true
//line /usr/local/go/src/net/http/request.go:972
	// _ = "end of CoverTab[41748]"
}

var textprotoReaderPool sync.Pool

func newTextprotoReader(br *bufio.Reader) *textproto.Reader {
//line /usr/local/go/src/net/http/request.go:977
	_go_fuzz_dep_.CoverTab[41752]++
							if v := textprotoReaderPool.Get(); v != nil {
//line /usr/local/go/src/net/http/request.go:978
		_go_fuzz_dep_.CoverTab[41754]++
								tr := v.(*textproto.Reader)
								tr.R = br
								return tr
//line /usr/local/go/src/net/http/request.go:981
		// _ = "end of CoverTab[41754]"
	} else {
//line /usr/local/go/src/net/http/request.go:982
		_go_fuzz_dep_.CoverTab[41755]++
//line /usr/local/go/src/net/http/request.go:982
		// _ = "end of CoverTab[41755]"
//line /usr/local/go/src/net/http/request.go:982
	}
//line /usr/local/go/src/net/http/request.go:982
	// _ = "end of CoverTab[41752]"
//line /usr/local/go/src/net/http/request.go:982
	_go_fuzz_dep_.CoverTab[41753]++
							return textproto.NewReader(br)
//line /usr/local/go/src/net/http/request.go:983
	// _ = "end of CoverTab[41753]"
}

func putTextprotoReader(r *textproto.Reader) {
//line /usr/local/go/src/net/http/request.go:986
	_go_fuzz_dep_.CoverTab[41756]++
							r.R = nil
							textprotoReaderPool.Put(r)
//line /usr/local/go/src/net/http/request.go:988
	// _ = "end of CoverTab[41756]"
}

// ReadRequest reads and parses an incoming request from b.
//line /usr/local/go/src/net/http/request.go:991
//
//line /usr/local/go/src/net/http/request.go:991
// ReadRequest is a low-level function and should only be used for
//line /usr/local/go/src/net/http/request.go:991
// specialized applications; most code should use the Server to read
//line /usr/local/go/src/net/http/request.go:991
// requests and handle them via the Handler interface. ReadRequest
//line /usr/local/go/src/net/http/request.go:991
// only supports HTTP/1.x requests. For HTTP/2, use golang.org/x/net/http2.
//line /usr/local/go/src/net/http/request.go:997
func ReadRequest(b *bufio.Reader) (*Request, error) {
//line /usr/local/go/src/net/http/request.go:997
	_go_fuzz_dep_.CoverTab[41757]++
							req, err := readRequest(b)
							if err != nil {
//line /usr/local/go/src/net/http/request.go:999
		_go_fuzz_dep_.CoverTab[41759]++
								return nil, err
//line /usr/local/go/src/net/http/request.go:1000
		// _ = "end of CoverTab[41759]"
	} else {
//line /usr/local/go/src/net/http/request.go:1001
		_go_fuzz_dep_.CoverTab[41760]++
//line /usr/local/go/src/net/http/request.go:1001
		// _ = "end of CoverTab[41760]"
//line /usr/local/go/src/net/http/request.go:1001
	}
//line /usr/local/go/src/net/http/request.go:1001
	// _ = "end of CoverTab[41757]"
//line /usr/local/go/src/net/http/request.go:1001
	_go_fuzz_dep_.CoverTab[41758]++

							delete(req.Header, "Host")
							return req, err
//line /usr/local/go/src/net/http/request.go:1004
	// _ = "end of CoverTab[41758]"
}

func readRequest(b *bufio.Reader) (req *Request, err error) {
//line /usr/local/go/src/net/http/request.go:1007
	_go_fuzz_dep_.CoverTab[41761]++
							tp := newTextprotoReader(b)
							defer putTextprotoReader(tp)

							req = new(Request)

	// First line: GET /index.html HTTP/1.0
	var s string
	if s, err = tp.ReadLine(); err != nil {
//line /usr/local/go/src/net/http/request.go:1015
		_go_fuzz_dep_.CoverTab[41775]++
								return nil, err
//line /usr/local/go/src/net/http/request.go:1016
		// _ = "end of CoverTab[41775]"
	} else {
//line /usr/local/go/src/net/http/request.go:1017
		_go_fuzz_dep_.CoverTab[41776]++
//line /usr/local/go/src/net/http/request.go:1017
		// _ = "end of CoverTab[41776]"
//line /usr/local/go/src/net/http/request.go:1017
	}
//line /usr/local/go/src/net/http/request.go:1017
	// _ = "end of CoverTab[41761]"
//line /usr/local/go/src/net/http/request.go:1017
	_go_fuzz_dep_.CoverTab[41762]++
							defer func() {
//line /usr/local/go/src/net/http/request.go:1018
		_go_fuzz_dep_.CoverTab[41777]++
								if err == io.EOF {
//line /usr/local/go/src/net/http/request.go:1019
			_go_fuzz_dep_.CoverTab[41778]++
									err = io.ErrUnexpectedEOF
//line /usr/local/go/src/net/http/request.go:1020
			// _ = "end of CoverTab[41778]"
		} else {
//line /usr/local/go/src/net/http/request.go:1021
			_go_fuzz_dep_.CoverTab[41779]++
//line /usr/local/go/src/net/http/request.go:1021
			// _ = "end of CoverTab[41779]"
//line /usr/local/go/src/net/http/request.go:1021
		}
//line /usr/local/go/src/net/http/request.go:1021
		// _ = "end of CoverTab[41777]"
	}()
//line /usr/local/go/src/net/http/request.go:1022
	// _ = "end of CoverTab[41762]"
//line /usr/local/go/src/net/http/request.go:1022
	_go_fuzz_dep_.CoverTab[41763]++

							var ok bool
							req.Method, req.RequestURI, req.Proto, ok = parseRequestLine(s)
							if !ok {
//line /usr/local/go/src/net/http/request.go:1026
		_go_fuzz_dep_.CoverTab[41780]++
								return nil, badStringError("malformed HTTP request", s)
//line /usr/local/go/src/net/http/request.go:1027
		// _ = "end of CoverTab[41780]"
	} else {
//line /usr/local/go/src/net/http/request.go:1028
		_go_fuzz_dep_.CoverTab[41781]++
//line /usr/local/go/src/net/http/request.go:1028
		// _ = "end of CoverTab[41781]"
//line /usr/local/go/src/net/http/request.go:1028
	}
//line /usr/local/go/src/net/http/request.go:1028
	// _ = "end of CoverTab[41763]"
//line /usr/local/go/src/net/http/request.go:1028
	_go_fuzz_dep_.CoverTab[41764]++
							if !validMethod(req.Method) {
//line /usr/local/go/src/net/http/request.go:1029
		_go_fuzz_dep_.CoverTab[41782]++
								return nil, badStringError("invalid method", req.Method)
//line /usr/local/go/src/net/http/request.go:1030
		// _ = "end of CoverTab[41782]"
	} else {
//line /usr/local/go/src/net/http/request.go:1031
		_go_fuzz_dep_.CoverTab[41783]++
//line /usr/local/go/src/net/http/request.go:1031
		// _ = "end of CoverTab[41783]"
//line /usr/local/go/src/net/http/request.go:1031
	}
//line /usr/local/go/src/net/http/request.go:1031
	// _ = "end of CoverTab[41764]"
//line /usr/local/go/src/net/http/request.go:1031
	_go_fuzz_dep_.CoverTab[41765]++
							rawurl := req.RequestURI
							if req.ProtoMajor, req.ProtoMinor, ok = ParseHTTPVersion(req.Proto); !ok {
//line /usr/local/go/src/net/http/request.go:1033
		_go_fuzz_dep_.CoverTab[41784]++
								return nil, badStringError("malformed HTTP version", req.Proto)
//line /usr/local/go/src/net/http/request.go:1034
		// _ = "end of CoverTab[41784]"
	} else {
//line /usr/local/go/src/net/http/request.go:1035
		_go_fuzz_dep_.CoverTab[41785]++
//line /usr/local/go/src/net/http/request.go:1035
		// _ = "end of CoverTab[41785]"
//line /usr/local/go/src/net/http/request.go:1035
	}
//line /usr/local/go/src/net/http/request.go:1035
	// _ = "end of CoverTab[41765]"
//line /usr/local/go/src/net/http/request.go:1035
	_go_fuzz_dep_.CoverTab[41766]++

//line /usr/local/go/src/net/http/request.go:1046
	justAuthority := req.Method == "CONNECT" && func() bool {
//line /usr/local/go/src/net/http/request.go:1046
		_go_fuzz_dep_.CoverTab[41786]++
//line /usr/local/go/src/net/http/request.go:1046
		return !strings.HasPrefix(rawurl, "/")
//line /usr/local/go/src/net/http/request.go:1046
		// _ = "end of CoverTab[41786]"
//line /usr/local/go/src/net/http/request.go:1046
	}()
	if justAuthority {
//line /usr/local/go/src/net/http/request.go:1047
		_go_fuzz_dep_.CoverTab[41787]++
								rawurl = "http://" + rawurl
//line /usr/local/go/src/net/http/request.go:1048
		// _ = "end of CoverTab[41787]"
	} else {
//line /usr/local/go/src/net/http/request.go:1049
		_go_fuzz_dep_.CoverTab[41788]++
//line /usr/local/go/src/net/http/request.go:1049
		// _ = "end of CoverTab[41788]"
//line /usr/local/go/src/net/http/request.go:1049
	}
//line /usr/local/go/src/net/http/request.go:1049
	// _ = "end of CoverTab[41766]"
//line /usr/local/go/src/net/http/request.go:1049
	_go_fuzz_dep_.CoverTab[41767]++

							if req.URL, err = url.ParseRequestURI(rawurl); err != nil {
//line /usr/local/go/src/net/http/request.go:1051
		_go_fuzz_dep_.CoverTab[41789]++
								return nil, err
//line /usr/local/go/src/net/http/request.go:1052
		// _ = "end of CoverTab[41789]"
	} else {
//line /usr/local/go/src/net/http/request.go:1053
		_go_fuzz_dep_.CoverTab[41790]++
//line /usr/local/go/src/net/http/request.go:1053
		// _ = "end of CoverTab[41790]"
//line /usr/local/go/src/net/http/request.go:1053
	}
//line /usr/local/go/src/net/http/request.go:1053
	// _ = "end of CoverTab[41767]"
//line /usr/local/go/src/net/http/request.go:1053
	_go_fuzz_dep_.CoverTab[41768]++

							if justAuthority {
//line /usr/local/go/src/net/http/request.go:1055
		_go_fuzz_dep_.CoverTab[41791]++

								req.URL.Scheme = ""
//line /usr/local/go/src/net/http/request.go:1057
		// _ = "end of CoverTab[41791]"
	} else {
//line /usr/local/go/src/net/http/request.go:1058
		_go_fuzz_dep_.CoverTab[41792]++
//line /usr/local/go/src/net/http/request.go:1058
		// _ = "end of CoverTab[41792]"
//line /usr/local/go/src/net/http/request.go:1058
	}
//line /usr/local/go/src/net/http/request.go:1058
	// _ = "end of CoverTab[41768]"
//line /usr/local/go/src/net/http/request.go:1058
	_go_fuzz_dep_.CoverTab[41769]++

//line /usr/local/go/src/net/http/request.go:1061
	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
//line /usr/local/go/src/net/http/request.go:1062
		_go_fuzz_dep_.CoverTab[41793]++
								return nil, err
//line /usr/local/go/src/net/http/request.go:1063
		// _ = "end of CoverTab[41793]"
	} else {
//line /usr/local/go/src/net/http/request.go:1064
		_go_fuzz_dep_.CoverTab[41794]++
//line /usr/local/go/src/net/http/request.go:1064
		// _ = "end of CoverTab[41794]"
//line /usr/local/go/src/net/http/request.go:1064
	}
//line /usr/local/go/src/net/http/request.go:1064
	// _ = "end of CoverTab[41769]"
//line /usr/local/go/src/net/http/request.go:1064
	_go_fuzz_dep_.CoverTab[41770]++
							req.Header = Header(mimeHeader)
							if len(req.Header["Host"]) > 1 {
//line /usr/local/go/src/net/http/request.go:1066
		_go_fuzz_dep_.CoverTab[41795]++
								return nil, fmt.Errorf("too many Host headers")
//line /usr/local/go/src/net/http/request.go:1067
		// _ = "end of CoverTab[41795]"
	} else {
//line /usr/local/go/src/net/http/request.go:1068
		_go_fuzz_dep_.CoverTab[41796]++
//line /usr/local/go/src/net/http/request.go:1068
		// _ = "end of CoverTab[41796]"
//line /usr/local/go/src/net/http/request.go:1068
	}
//line /usr/local/go/src/net/http/request.go:1068
	// _ = "end of CoverTab[41770]"
//line /usr/local/go/src/net/http/request.go:1068
	_go_fuzz_dep_.CoverTab[41771]++

//line /usr/local/go/src/net/http/request.go:1077
	req.Host = req.URL.Host
	if req.Host == "" {
//line /usr/local/go/src/net/http/request.go:1078
		_go_fuzz_dep_.CoverTab[41797]++
								req.Host = req.Header.get("Host")
//line /usr/local/go/src/net/http/request.go:1079
		// _ = "end of CoverTab[41797]"
	} else {
//line /usr/local/go/src/net/http/request.go:1080
		_go_fuzz_dep_.CoverTab[41798]++
//line /usr/local/go/src/net/http/request.go:1080
		// _ = "end of CoverTab[41798]"
//line /usr/local/go/src/net/http/request.go:1080
	}
//line /usr/local/go/src/net/http/request.go:1080
	// _ = "end of CoverTab[41771]"
//line /usr/local/go/src/net/http/request.go:1080
	_go_fuzz_dep_.CoverTab[41772]++

							fixPragmaCacheControl(req.Header)

							req.Close = shouldClose(req.ProtoMajor, req.ProtoMinor, req.Header, false)

							err = readTransfer(req, b)
							if err != nil {
//line /usr/local/go/src/net/http/request.go:1087
		_go_fuzz_dep_.CoverTab[41799]++
								return nil, err
//line /usr/local/go/src/net/http/request.go:1088
		// _ = "end of CoverTab[41799]"
	} else {
//line /usr/local/go/src/net/http/request.go:1089
		_go_fuzz_dep_.CoverTab[41800]++
//line /usr/local/go/src/net/http/request.go:1089
		// _ = "end of CoverTab[41800]"
//line /usr/local/go/src/net/http/request.go:1089
	}
//line /usr/local/go/src/net/http/request.go:1089
	// _ = "end of CoverTab[41772]"
//line /usr/local/go/src/net/http/request.go:1089
	_go_fuzz_dep_.CoverTab[41773]++

							if req.isH2Upgrade() {
//line /usr/local/go/src/net/http/request.go:1091
		_go_fuzz_dep_.CoverTab[41801]++

								req.ContentLength = -1

//line /usr/local/go/src/net/http/request.go:1099
		req.Close = true
//line /usr/local/go/src/net/http/request.go:1099
		// _ = "end of CoverTab[41801]"
	} else {
//line /usr/local/go/src/net/http/request.go:1100
		_go_fuzz_dep_.CoverTab[41802]++
//line /usr/local/go/src/net/http/request.go:1100
		// _ = "end of CoverTab[41802]"
//line /usr/local/go/src/net/http/request.go:1100
	}
//line /usr/local/go/src/net/http/request.go:1100
	// _ = "end of CoverTab[41773]"
//line /usr/local/go/src/net/http/request.go:1100
	_go_fuzz_dep_.CoverTab[41774]++
							return req, nil
//line /usr/local/go/src/net/http/request.go:1101
	// _ = "end of CoverTab[41774]"
}

// MaxBytesReader is similar to io.LimitReader but is intended for
//line /usr/local/go/src/net/http/request.go:1104
// limiting the size of incoming request bodies. In contrast to
//line /usr/local/go/src/net/http/request.go:1104
// io.LimitReader, MaxBytesReader's result is a ReadCloser, returns a
//line /usr/local/go/src/net/http/request.go:1104
// non-nil error of type *MaxBytesError for a Read beyond the limit,
//line /usr/local/go/src/net/http/request.go:1104
// and closes the underlying reader when its Close method is called.
//line /usr/local/go/src/net/http/request.go:1104
//
//line /usr/local/go/src/net/http/request.go:1104
// MaxBytesReader prevents clients from accidentally or maliciously
//line /usr/local/go/src/net/http/request.go:1104
// sending a large request and wasting server resources. If possible,
//line /usr/local/go/src/net/http/request.go:1104
// it tells the ResponseWriter to close the connection after the limit
//line /usr/local/go/src/net/http/request.go:1104
// has been reached.
//line /usr/local/go/src/net/http/request.go:1114
func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser {
//line /usr/local/go/src/net/http/request.go:1114
	_go_fuzz_dep_.CoverTab[41803]++
							if n < 0 {
//line /usr/local/go/src/net/http/request.go:1115
		_go_fuzz_dep_.CoverTab[41805]++
								n = 0
//line /usr/local/go/src/net/http/request.go:1116
		// _ = "end of CoverTab[41805]"
	} else {
//line /usr/local/go/src/net/http/request.go:1117
		_go_fuzz_dep_.CoverTab[41806]++
//line /usr/local/go/src/net/http/request.go:1117
		// _ = "end of CoverTab[41806]"
//line /usr/local/go/src/net/http/request.go:1117
	}
//line /usr/local/go/src/net/http/request.go:1117
	// _ = "end of CoverTab[41803]"
//line /usr/local/go/src/net/http/request.go:1117
	_go_fuzz_dep_.CoverTab[41804]++
							return &maxBytesReader{w: w, r: r, i: n, n: n}
//line /usr/local/go/src/net/http/request.go:1118
	// _ = "end of CoverTab[41804]"
}

// MaxBytesError is returned by MaxBytesReader when its read limit is exceeded.
type MaxBytesError struct {
	Limit int64
}

func (e *MaxBytesError) Error() string {
//line /usr/local/go/src/net/http/request.go:1126
	_go_fuzz_dep_.CoverTab[41807]++

							return "http: request body too large"
//line /usr/local/go/src/net/http/request.go:1128
	// _ = "end of CoverTab[41807]"
}

type maxBytesReader struct {
	w	ResponseWriter
	r	io.ReadCloser	// underlying reader
	i	int64		// max bytes initially, for MaxBytesError
	n	int64		// max bytes remaining
	err	error		// sticky error
}

func (l *maxBytesReader) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/request.go:1139
	_go_fuzz_dep_.CoverTab[41808]++
							if l.err != nil {
//line /usr/local/go/src/net/http/request.go:1140
		_go_fuzz_dep_.CoverTab[41814]++
								return 0, l.err
//line /usr/local/go/src/net/http/request.go:1141
		// _ = "end of CoverTab[41814]"
	} else {
//line /usr/local/go/src/net/http/request.go:1142
		_go_fuzz_dep_.CoverTab[41815]++
//line /usr/local/go/src/net/http/request.go:1142
		// _ = "end of CoverTab[41815]"
//line /usr/local/go/src/net/http/request.go:1142
	}
//line /usr/local/go/src/net/http/request.go:1142
	// _ = "end of CoverTab[41808]"
//line /usr/local/go/src/net/http/request.go:1142
	_go_fuzz_dep_.CoverTab[41809]++
							if len(p) == 0 {
//line /usr/local/go/src/net/http/request.go:1143
		_go_fuzz_dep_.CoverTab[41816]++
								return 0, nil
//line /usr/local/go/src/net/http/request.go:1144
		// _ = "end of CoverTab[41816]"
	} else {
//line /usr/local/go/src/net/http/request.go:1145
		_go_fuzz_dep_.CoverTab[41817]++
//line /usr/local/go/src/net/http/request.go:1145
		// _ = "end of CoverTab[41817]"
//line /usr/local/go/src/net/http/request.go:1145
	}
//line /usr/local/go/src/net/http/request.go:1145
	// _ = "end of CoverTab[41809]"
//line /usr/local/go/src/net/http/request.go:1145
	_go_fuzz_dep_.CoverTab[41810]++

//line /usr/local/go/src/net/http/request.go:1150
	if int64(len(p))-1 > l.n {
//line /usr/local/go/src/net/http/request.go:1150
		_go_fuzz_dep_.CoverTab[41818]++
								p = p[:l.n+1]
//line /usr/local/go/src/net/http/request.go:1151
		// _ = "end of CoverTab[41818]"
	} else {
//line /usr/local/go/src/net/http/request.go:1152
		_go_fuzz_dep_.CoverTab[41819]++
//line /usr/local/go/src/net/http/request.go:1152
		// _ = "end of CoverTab[41819]"
//line /usr/local/go/src/net/http/request.go:1152
	}
//line /usr/local/go/src/net/http/request.go:1152
	// _ = "end of CoverTab[41810]"
//line /usr/local/go/src/net/http/request.go:1152
	_go_fuzz_dep_.CoverTab[41811]++
							n, err = l.r.Read(p)

							if int64(n) <= l.n {
//line /usr/local/go/src/net/http/request.go:1155
		_go_fuzz_dep_.CoverTab[41820]++
								l.n -= int64(n)
								l.err = err
								return n, err
//line /usr/local/go/src/net/http/request.go:1158
		// _ = "end of CoverTab[41820]"
	} else {
//line /usr/local/go/src/net/http/request.go:1159
		_go_fuzz_dep_.CoverTab[41821]++
//line /usr/local/go/src/net/http/request.go:1159
		// _ = "end of CoverTab[41821]"
//line /usr/local/go/src/net/http/request.go:1159
	}
//line /usr/local/go/src/net/http/request.go:1159
	// _ = "end of CoverTab[41811]"
//line /usr/local/go/src/net/http/request.go:1159
	_go_fuzz_dep_.CoverTab[41812]++

							n = int(l.n)
							l.n = 0

	// The server code and client code both use
	// maxBytesReader. This "requestTooLarge" check is
	// only used by the server code. To prevent binaries
	// which only using the HTTP Client code (such as
	// cmd/go) from also linking in the HTTP server, don't
	// use a static type assertion to the server
	// "*response" type. Check this interface instead:
	type requestTooLarger interface {
		requestTooLarge()
	}
	if res, ok := l.w.(requestTooLarger); ok {
//line /usr/local/go/src/net/http/request.go:1174
		_go_fuzz_dep_.CoverTab[41822]++
								res.requestTooLarge()
//line /usr/local/go/src/net/http/request.go:1175
		// _ = "end of CoverTab[41822]"
	} else {
//line /usr/local/go/src/net/http/request.go:1176
		_go_fuzz_dep_.CoverTab[41823]++
//line /usr/local/go/src/net/http/request.go:1176
		// _ = "end of CoverTab[41823]"
//line /usr/local/go/src/net/http/request.go:1176
	}
//line /usr/local/go/src/net/http/request.go:1176
	// _ = "end of CoverTab[41812]"
//line /usr/local/go/src/net/http/request.go:1176
	_go_fuzz_dep_.CoverTab[41813]++
							l.err = &MaxBytesError{l.i}
							return n, l.err
//line /usr/local/go/src/net/http/request.go:1178
	// _ = "end of CoverTab[41813]"
}

func (l *maxBytesReader) Close() error {
//line /usr/local/go/src/net/http/request.go:1181
	_go_fuzz_dep_.CoverTab[41824]++
							return l.r.Close()
//line /usr/local/go/src/net/http/request.go:1182
	// _ = "end of CoverTab[41824]"
}

func copyValues(dst, src url.Values) {
//line /usr/local/go/src/net/http/request.go:1185
	_go_fuzz_dep_.CoverTab[41825]++
							for k, vs := range src {
//line /usr/local/go/src/net/http/request.go:1186
		_go_fuzz_dep_.CoverTab[41826]++
								dst[k] = append(dst[k], vs...)
//line /usr/local/go/src/net/http/request.go:1187
		// _ = "end of CoverTab[41826]"
	}
//line /usr/local/go/src/net/http/request.go:1188
	// _ = "end of CoverTab[41825]"
}

func parsePostForm(r *Request) (vs url.Values, err error) {
//line /usr/local/go/src/net/http/request.go:1191
	_go_fuzz_dep_.CoverTab[41827]++
							if r.Body == nil {
//line /usr/local/go/src/net/http/request.go:1192
		_go_fuzz_dep_.CoverTab[41831]++
								err = errors.New("missing form body")
								return
//line /usr/local/go/src/net/http/request.go:1194
		// _ = "end of CoverTab[41831]"
	} else {
//line /usr/local/go/src/net/http/request.go:1195
		_go_fuzz_dep_.CoverTab[41832]++
//line /usr/local/go/src/net/http/request.go:1195
		// _ = "end of CoverTab[41832]"
//line /usr/local/go/src/net/http/request.go:1195
	}
//line /usr/local/go/src/net/http/request.go:1195
	// _ = "end of CoverTab[41827]"
//line /usr/local/go/src/net/http/request.go:1195
	_go_fuzz_dep_.CoverTab[41828]++
							ct := r.Header.Get("Content-Type")

//line /usr/local/go/src/net/http/request.go:1199
	if ct == "" {
//line /usr/local/go/src/net/http/request.go:1199
		_go_fuzz_dep_.CoverTab[41833]++
								ct = "application/octet-stream"
//line /usr/local/go/src/net/http/request.go:1200
		// _ = "end of CoverTab[41833]"
	} else {
//line /usr/local/go/src/net/http/request.go:1201
		_go_fuzz_dep_.CoverTab[41834]++
//line /usr/local/go/src/net/http/request.go:1201
		// _ = "end of CoverTab[41834]"
//line /usr/local/go/src/net/http/request.go:1201
	}
//line /usr/local/go/src/net/http/request.go:1201
	// _ = "end of CoverTab[41828]"
//line /usr/local/go/src/net/http/request.go:1201
	_go_fuzz_dep_.CoverTab[41829]++
							ct, _, err = mime.ParseMediaType(ct)
							switch {
	case ct == "application/x-www-form-urlencoded":
//line /usr/local/go/src/net/http/request.go:1204
		_go_fuzz_dep_.CoverTab[41835]++
								var reader io.Reader = r.Body
								maxFormSize := int64(1<<63 - 1)
								if _, ok := r.Body.(*maxBytesReader); !ok {
//line /usr/local/go/src/net/http/request.go:1207
			_go_fuzz_dep_.CoverTab[41841]++
									maxFormSize = int64(10 << 20)
									reader = io.LimitReader(r.Body, maxFormSize+1)
//line /usr/local/go/src/net/http/request.go:1209
			// _ = "end of CoverTab[41841]"
		} else {
//line /usr/local/go/src/net/http/request.go:1210
			_go_fuzz_dep_.CoverTab[41842]++
//line /usr/local/go/src/net/http/request.go:1210
			// _ = "end of CoverTab[41842]"
//line /usr/local/go/src/net/http/request.go:1210
		}
//line /usr/local/go/src/net/http/request.go:1210
		// _ = "end of CoverTab[41835]"
//line /usr/local/go/src/net/http/request.go:1210
		_go_fuzz_dep_.CoverTab[41836]++
								b, e := io.ReadAll(reader)
								if e != nil {
//line /usr/local/go/src/net/http/request.go:1212
			_go_fuzz_dep_.CoverTab[41843]++
									if err == nil {
//line /usr/local/go/src/net/http/request.go:1213
				_go_fuzz_dep_.CoverTab[41845]++
										err = e
//line /usr/local/go/src/net/http/request.go:1214
				// _ = "end of CoverTab[41845]"
			} else {
//line /usr/local/go/src/net/http/request.go:1215
				_go_fuzz_dep_.CoverTab[41846]++
//line /usr/local/go/src/net/http/request.go:1215
				// _ = "end of CoverTab[41846]"
//line /usr/local/go/src/net/http/request.go:1215
			}
//line /usr/local/go/src/net/http/request.go:1215
			// _ = "end of CoverTab[41843]"
//line /usr/local/go/src/net/http/request.go:1215
			_go_fuzz_dep_.CoverTab[41844]++
									break
//line /usr/local/go/src/net/http/request.go:1216
			// _ = "end of CoverTab[41844]"
		} else {
//line /usr/local/go/src/net/http/request.go:1217
			_go_fuzz_dep_.CoverTab[41847]++
//line /usr/local/go/src/net/http/request.go:1217
			// _ = "end of CoverTab[41847]"
//line /usr/local/go/src/net/http/request.go:1217
		}
//line /usr/local/go/src/net/http/request.go:1217
		// _ = "end of CoverTab[41836]"
//line /usr/local/go/src/net/http/request.go:1217
		_go_fuzz_dep_.CoverTab[41837]++
								if int64(len(b)) > maxFormSize {
//line /usr/local/go/src/net/http/request.go:1218
			_go_fuzz_dep_.CoverTab[41848]++
									err = errors.New("http: POST too large")
									return
//line /usr/local/go/src/net/http/request.go:1220
			// _ = "end of CoverTab[41848]"
		} else {
//line /usr/local/go/src/net/http/request.go:1221
			_go_fuzz_dep_.CoverTab[41849]++
//line /usr/local/go/src/net/http/request.go:1221
			// _ = "end of CoverTab[41849]"
//line /usr/local/go/src/net/http/request.go:1221
		}
//line /usr/local/go/src/net/http/request.go:1221
		// _ = "end of CoverTab[41837]"
//line /usr/local/go/src/net/http/request.go:1221
		_go_fuzz_dep_.CoverTab[41838]++
								vs, e = url.ParseQuery(string(b))
								if err == nil {
//line /usr/local/go/src/net/http/request.go:1223
			_go_fuzz_dep_.CoverTab[41850]++
									err = e
//line /usr/local/go/src/net/http/request.go:1224
			// _ = "end of CoverTab[41850]"
		} else {
//line /usr/local/go/src/net/http/request.go:1225
			_go_fuzz_dep_.CoverTab[41851]++
//line /usr/local/go/src/net/http/request.go:1225
			// _ = "end of CoverTab[41851]"
//line /usr/local/go/src/net/http/request.go:1225
		}
//line /usr/local/go/src/net/http/request.go:1225
		// _ = "end of CoverTab[41838]"
	case ct == "multipart/form-data":
//line /usr/local/go/src/net/http/request.go:1226
		_go_fuzz_dep_.CoverTab[41839]++
//line /usr/local/go/src/net/http/request.go:1226
		// _ = "end of CoverTab[41839]"
//line /usr/local/go/src/net/http/request.go:1226
	default:
//line /usr/local/go/src/net/http/request.go:1226
		_go_fuzz_dep_.CoverTab[41840]++
//line /usr/local/go/src/net/http/request.go:1226
		// _ = "end of CoverTab[41840]"

//line /usr/local/go/src/net/http/request.go:1233
	}
//line /usr/local/go/src/net/http/request.go:1233
	// _ = "end of CoverTab[41829]"
//line /usr/local/go/src/net/http/request.go:1233
	_go_fuzz_dep_.CoverTab[41830]++
							return
//line /usr/local/go/src/net/http/request.go:1234
	// _ = "end of CoverTab[41830]"
}

// ParseForm populates r.Form and r.PostForm.
//line /usr/local/go/src/net/http/request.go:1237
//
//line /usr/local/go/src/net/http/request.go:1237
// For all requests, ParseForm parses the raw query from the URL and updates
//line /usr/local/go/src/net/http/request.go:1237
// r.Form.
//line /usr/local/go/src/net/http/request.go:1237
//
//line /usr/local/go/src/net/http/request.go:1237
// For POST, PUT, and PATCH requests, it also reads the request body, parses it
//line /usr/local/go/src/net/http/request.go:1237
// as a form and puts the results into both r.PostForm and r.Form. Request body
//line /usr/local/go/src/net/http/request.go:1237
// parameters take precedence over URL query string values in r.Form.
//line /usr/local/go/src/net/http/request.go:1237
//
//line /usr/local/go/src/net/http/request.go:1237
// If the request Body's size has not already been limited by MaxBytesReader,
//line /usr/local/go/src/net/http/request.go:1237
// the size is capped at 10MB.
//line /usr/local/go/src/net/http/request.go:1237
//
//line /usr/local/go/src/net/http/request.go:1237
// For other HTTP methods, or when the Content-Type is not
//line /usr/local/go/src/net/http/request.go:1237
// application/x-www-form-urlencoded, the request Body is not read, and
//line /usr/local/go/src/net/http/request.go:1237
// r.PostForm is initialized to a non-nil, empty value.
//line /usr/local/go/src/net/http/request.go:1237
//
//line /usr/local/go/src/net/http/request.go:1237
// ParseMultipartForm calls ParseForm automatically.
//line /usr/local/go/src/net/http/request.go:1237
// ParseForm is idempotent.
//line /usr/local/go/src/net/http/request.go:1255
func (r *Request) ParseForm() error {
//line /usr/local/go/src/net/http/request.go:1255
	_go_fuzz_dep_.CoverTab[41852]++
							var err error
							if r.PostForm == nil {
//line /usr/local/go/src/net/http/request.go:1257
		_go_fuzz_dep_.CoverTab[41855]++
								if r.Method == "POST" || func() bool {
//line /usr/local/go/src/net/http/request.go:1258
			_go_fuzz_dep_.CoverTab[41857]++
//line /usr/local/go/src/net/http/request.go:1258
			return r.Method == "PUT"
//line /usr/local/go/src/net/http/request.go:1258
			// _ = "end of CoverTab[41857]"
//line /usr/local/go/src/net/http/request.go:1258
		}() || func() bool {
//line /usr/local/go/src/net/http/request.go:1258
			_go_fuzz_dep_.CoverTab[41858]++
//line /usr/local/go/src/net/http/request.go:1258
			return r.Method == "PATCH"
//line /usr/local/go/src/net/http/request.go:1258
			// _ = "end of CoverTab[41858]"
//line /usr/local/go/src/net/http/request.go:1258
		}() {
//line /usr/local/go/src/net/http/request.go:1258
			_go_fuzz_dep_.CoverTab[41859]++
									r.PostForm, err = parsePostForm(r)
//line /usr/local/go/src/net/http/request.go:1259
			// _ = "end of CoverTab[41859]"
		} else {
//line /usr/local/go/src/net/http/request.go:1260
			_go_fuzz_dep_.CoverTab[41860]++
//line /usr/local/go/src/net/http/request.go:1260
			// _ = "end of CoverTab[41860]"
//line /usr/local/go/src/net/http/request.go:1260
		}
//line /usr/local/go/src/net/http/request.go:1260
		// _ = "end of CoverTab[41855]"
//line /usr/local/go/src/net/http/request.go:1260
		_go_fuzz_dep_.CoverTab[41856]++
								if r.PostForm == nil {
//line /usr/local/go/src/net/http/request.go:1261
			_go_fuzz_dep_.CoverTab[41861]++
									r.PostForm = make(url.Values)
//line /usr/local/go/src/net/http/request.go:1262
			// _ = "end of CoverTab[41861]"
		} else {
//line /usr/local/go/src/net/http/request.go:1263
			_go_fuzz_dep_.CoverTab[41862]++
//line /usr/local/go/src/net/http/request.go:1263
			// _ = "end of CoverTab[41862]"
//line /usr/local/go/src/net/http/request.go:1263
		}
//line /usr/local/go/src/net/http/request.go:1263
		// _ = "end of CoverTab[41856]"
	} else {
//line /usr/local/go/src/net/http/request.go:1264
		_go_fuzz_dep_.CoverTab[41863]++
//line /usr/local/go/src/net/http/request.go:1264
		// _ = "end of CoverTab[41863]"
//line /usr/local/go/src/net/http/request.go:1264
	}
//line /usr/local/go/src/net/http/request.go:1264
	// _ = "end of CoverTab[41852]"
//line /usr/local/go/src/net/http/request.go:1264
	_go_fuzz_dep_.CoverTab[41853]++
							if r.Form == nil {
//line /usr/local/go/src/net/http/request.go:1265
		_go_fuzz_dep_.CoverTab[41864]++
								if len(r.PostForm) > 0 {
//line /usr/local/go/src/net/http/request.go:1266
			_go_fuzz_dep_.CoverTab[41868]++
									r.Form = make(url.Values)
									copyValues(r.Form, r.PostForm)
//line /usr/local/go/src/net/http/request.go:1268
			// _ = "end of CoverTab[41868]"
		} else {
//line /usr/local/go/src/net/http/request.go:1269
			_go_fuzz_dep_.CoverTab[41869]++
//line /usr/local/go/src/net/http/request.go:1269
			// _ = "end of CoverTab[41869]"
//line /usr/local/go/src/net/http/request.go:1269
		}
//line /usr/local/go/src/net/http/request.go:1269
		// _ = "end of CoverTab[41864]"
//line /usr/local/go/src/net/http/request.go:1269
		_go_fuzz_dep_.CoverTab[41865]++
								var newValues url.Values
								if r.URL != nil {
//line /usr/local/go/src/net/http/request.go:1271
			_go_fuzz_dep_.CoverTab[41870]++
									var e error
									newValues, e = url.ParseQuery(r.URL.RawQuery)
									if err == nil {
//line /usr/local/go/src/net/http/request.go:1274
				_go_fuzz_dep_.CoverTab[41871]++
										err = e
//line /usr/local/go/src/net/http/request.go:1275
				// _ = "end of CoverTab[41871]"
			} else {
//line /usr/local/go/src/net/http/request.go:1276
				_go_fuzz_dep_.CoverTab[41872]++
//line /usr/local/go/src/net/http/request.go:1276
				// _ = "end of CoverTab[41872]"
//line /usr/local/go/src/net/http/request.go:1276
			}
//line /usr/local/go/src/net/http/request.go:1276
			// _ = "end of CoverTab[41870]"
		} else {
//line /usr/local/go/src/net/http/request.go:1277
			_go_fuzz_dep_.CoverTab[41873]++
//line /usr/local/go/src/net/http/request.go:1277
			// _ = "end of CoverTab[41873]"
//line /usr/local/go/src/net/http/request.go:1277
		}
//line /usr/local/go/src/net/http/request.go:1277
		// _ = "end of CoverTab[41865]"
//line /usr/local/go/src/net/http/request.go:1277
		_go_fuzz_dep_.CoverTab[41866]++
								if newValues == nil {
//line /usr/local/go/src/net/http/request.go:1278
			_go_fuzz_dep_.CoverTab[41874]++
									newValues = make(url.Values)
//line /usr/local/go/src/net/http/request.go:1279
			// _ = "end of CoverTab[41874]"
		} else {
//line /usr/local/go/src/net/http/request.go:1280
			_go_fuzz_dep_.CoverTab[41875]++
//line /usr/local/go/src/net/http/request.go:1280
			// _ = "end of CoverTab[41875]"
//line /usr/local/go/src/net/http/request.go:1280
		}
//line /usr/local/go/src/net/http/request.go:1280
		// _ = "end of CoverTab[41866]"
//line /usr/local/go/src/net/http/request.go:1280
		_go_fuzz_dep_.CoverTab[41867]++
								if r.Form == nil {
//line /usr/local/go/src/net/http/request.go:1281
			_go_fuzz_dep_.CoverTab[41876]++
									r.Form = newValues
//line /usr/local/go/src/net/http/request.go:1282
			// _ = "end of CoverTab[41876]"
		} else {
//line /usr/local/go/src/net/http/request.go:1283
			_go_fuzz_dep_.CoverTab[41877]++
									copyValues(r.Form, newValues)
//line /usr/local/go/src/net/http/request.go:1284
			// _ = "end of CoverTab[41877]"
		}
//line /usr/local/go/src/net/http/request.go:1285
		// _ = "end of CoverTab[41867]"
	} else {
//line /usr/local/go/src/net/http/request.go:1286
		_go_fuzz_dep_.CoverTab[41878]++
//line /usr/local/go/src/net/http/request.go:1286
		// _ = "end of CoverTab[41878]"
//line /usr/local/go/src/net/http/request.go:1286
	}
//line /usr/local/go/src/net/http/request.go:1286
	// _ = "end of CoverTab[41853]"
//line /usr/local/go/src/net/http/request.go:1286
	_go_fuzz_dep_.CoverTab[41854]++
							return err
//line /usr/local/go/src/net/http/request.go:1287
	// _ = "end of CoverTab[41854]"
}

// ParseMultipartForm parses a request body as multipart/form-data.
//line /usr/local/go/src/net/http/request.go:1290
// The whole request body is parsed and up to a total of maxMemory bytes of
//line /usr/local/go/src/net/http/request.go:1290
// its file parts are stored in memory, with the remainder stored on
//line /usr/local/go/src/net/http/request.go:1290
// disk in temporary files.
//line /usr/local/go/src/net/http/request.go:1290
// ParseMultipartForm calls ParseForm if necessary.
//line /usr/local/go/src/net/http/request.go:1290
// If ParseForm returns an error, ParseMultipartForm returns it but also
//line /usr/local/go/src/net/http/request.go:1290
// continues parsing the request body.
//line /usr/local/go/src/net/http/request.go:1290
// After one call to ParseMultipartForm, subsequent calls have no effect.
//line /usr/local/go/src/net/http/request.go:1298
func (r *Request) ParseMultipartForm(maxMemory int64) error {
//line /usr/local/go/src/net/http/request.go:1298
	_go_fuzz_dep_.CoverTab[41879]++
							if r.MultipartForm == multipartByReader {
//line /usr/local/go/src/net/http/request.go:1299
		_go_fuzz_dep_.CoverTab[41887]++
								return errors.New("http: multipart handled by MultipartReader")
//line /usr/local/go/src/net/http/request.go:1300
		// _ = "end of CoverTab[41887]"
	} else {
//line /usr/local/go/src/net/http/request.go:1301
		_go_fuzz_dep_.CoverTab[41888]++
//line /usr/local/go/src/net/http/request.go:1301
		// _ = "end of CoverTab[41888]"
//line /usr/local/go/src/net/http/request.go:1301
	}
//line /usr/local/go/src/net/http/request.go:1301
	// _ = "end of CoverTab[41879]"
//line /usr/local/go/src/net/http/request.go:1301
	_go_fuzz_dep_.CoverTab[41880]++
							var parseFormErr error
							if r.Form == nil {
//line /usr/local/go/src/net/http/request.go:1303
		_go_fuzz_dep_.CoverTab[41889]++

//line /usr/local/go/src/net/http/request.go:1306
		parseFormErr = r.ParseForm()
//line /usr/local/go/src/net/http/request.go:1306
		// _ = "end of CoverTab[41889]"
	} else {
//line /usr/local/go/src/net/http/request.go:1307
		_go_fuzz_dep_.CoverTab[41890]++
//line /usr/local/go/src/net/http/request.go:1307
		// _ = "end of CoverTab[41890]"
//line /usr/local/go/src/net/http/request.go:1307
	}
//line /usr/local/go/src/net/http/request.go:1307
	// _ = "end of CoverTab[41880]"
//line /usr/local/go/src/net/http/request.go:1307
	_go_fuzz_dep_.CoverTab[41881]++
							if r.MultipartForm != nil {
//line /usr/local/go/src/net/http/request.go:1308
		_go_fuzz_dep_.CoverTab[41891]++
								return nil
//line /usr/local/go/src/net/http/request.go:1309
		// _ = "end of CoverTab[41891]"
	} else {
//line /usr/local/go/src/net/http/request.go:1310
		_go_fuzz_dep_.CoverTab[41892]++
//line /usr/local/go/src/net/http/request.go:1310
		// _ = "end of CoverTab[41892]"
//line /usr/local/go/src/net/http/request.go:1310
	}
//line /usr/local/go/src/net/http/request.go:1310
	// _ = "end of CoverTab[41881]"
//line /usr/local/go/src/net/http/request.go:1310
	_go_fuzz_dep_.CoverTab[41882]++

							mr, err := r.multipartReader(false)
							if err != nil {
//line /usr/local/go/src/net/http/request.go:1313
		_go_fuzz_dep_.CoverTab[41893]++
								return err
//line /usr/local/go/src/net/http/request.go:1314
		// _ = "end of CoverTab[41893]"
	} else {
//line /usr/local/go/src/net/http/request.go:1315
		_go_fuzz_dep_.CoverTab[41894]++
//line /usr/local/go/src/net/http/request.go:1315
		// _ = "end of CoverTab[41894]"
//line /usr/local/go/src/net/http/request.go:1315
	}
//line /usr/local/go/src/net/http/request.go:1315
	// _ = "end of CoverTab[41882]"
//line /usr/local/go/src/net/http/request.go:1315
	_go_fuzz_dep_.CoverTab[41883]++

							f, err := mr.ReadForm(maxMemory)
							if err != nil {
//line /usr/local/go/src/net/http/request.go:1318
		_go_fuzz_dep_.CoverTab[41895]++
								return err
//line /usr/local/go/src/net/http/request.go:1319
		// _ = "end of CoverTab[41895]"
	} else {
//line /usr/local/go/src/net/http/request.go:1320
		_go_fuzz_dep_.CoverTab[41896]++
//line /usr/local/go/src/net/http/request.go:1320
		// _ = "end of CoverTab[41896]"
//line /usr/local/go/src/net/http/request.go:1320
	}
//line /usr/local/go/src/net/http/request.go:1320
	// _ = "end of CoverTab[41883]"
//line /usr/local/go/src/net/http/request.go:1320
	_go_fuzz_dep_.CoverTab[41884]++

							if r.PostForm == nil {
//line /usr/local/go/src/net/http/request.go:1322
		_go_fuzz_dep_.CoverTab[41897]++
								r.PostForm = make(url.Values)
//line /usr/local/go/src/net/http/request.go:1323
		// _ = "end of CoverTab[41897]"
	} else {
//line /usr/local/go/src/net/http/request.go:1324
		_go_fuzz_dep_.CoverTab[41898]++
//line /usr/local/go/src/net/http/request.go:1324
		// _ = "end of CoverTab[41898]"
//line /usr/local/go/src/net/http/request.go:1324
	}
//line /usr/local/go/src/net/http/request.go:1324
	// _ = "end of CoverTab[41884]"
//line /usr/local/go/src/net/http/request.go:1324
	_go_fuzz_dep_.CoverTab[41885]++
							for k, v := range f.Value {
//line /usr/local/go/src/net/http/request.go:1325
		_go_fuzz_dep_.CoverTab[41899]++
								r.Form[k] = append(r.Form[k], v...)

								r.PostForm[k] = append(r.PostForm[k], v...)
//line /usr/local/go/src/net/http/request.go:1328
		// _ = "end of CoverTab[41899]"
	}
//line /usr/local/go/src/net/http/request.go:1329
	// _ = "end of CoverTab[41885]"
//line /usr/local/go/src/net/http/request.go:1329
	_go_fuzz_dep_.CoverTab[41886]++

							r.MultipartForm = f

							return parseFormErr
//line /usr/local/go/src/net/http/request.go:1333
	// _ = "end of CoverTab[41886]"
}

// FormValue returns the first value for the named component of the query.
//line /usr/local/go/src/net/http/request.go:1336
// POST and PUT body parameters take precedence over URL query string values.
//line /usr/local/go/src/net/http/request.go:1336
// FormValue calls ParseMultipartForm and ParseForm if necessary and ignores
//line /usr/local/go/src/net/http/request.go:1336
// any errors returned by these functions.
//line /usr/local/go/src/net/http/request.go:1336
// If key is not present, FormValue returns the empty string.
//line /usr/local/go/src/net/http/request.go:1336
// To access multiple values of the same key, call ParseForm and
//line /usr/local/go/src/net/http/request.go:1336
// then inspect Request.Form directly.
//line /usr/local/go/src/net/http/request.go:1343
func (r *Request) FormValue(key string) string {
//line /usr/local/go/src/net/http/request.go:1343
	_go_fuzz_dep_.CoverTab[41900]++
							if r.Form == nil {
//line /usr/local/go/src/net/http/request.go:1344
		_go_fuzz_dep_.CoverTab[41903]++
								r.ParseMultipartForm(defaultMaxMemory)
//line /usr/local/go/src/net/http/request.go:1345
		// _ = "end of CoverTab[41903]"
	} else {
//line /usr/local/go/src/net/http/request.go:1346
		_go_fuzz_dep_.CoverTab[41904]++
//line /usr/local/go/src/net/http/request.go:1346
		// _ = "end of CoverTab[41904]"
//line /usr/local/go/src/net/http/request.go:1346
	}
//line /usr/local/go/src/net/http/request.go:1346
	// _ = "end of CoverTab[41900]"
//line /usr/local/go/src/net/http/request.go:1346
	_go_fuzz_dep_.CoverTab[41901]++
							if vs := r.Form[key]; len(vs) > 0 {
//line /usr/local/go/src/net/http/request.go:1347
		_go_fuzz_dep_.CoverTab[41905]++
								return vs[0]
//line /usr/local/go/src/net/http/request.go:1348
		// _ = "end of CoverTab[41905]"
	} else {
//line /usr/local/go/src/net/http/request.go:1349
		_go_fuzz_dep_.CoverTab[41906]++
//line /usr/local/go/src/net/http/request.go:1349
		// _ = "end of CoverTab[41906]"
//line /usr/local/go/src/net/http/request.go:1349
	}
//line /usr/local/go/src/net/http/request.go:1349
	// _ = "end of CoverTab[41901]"
//line /usr/local/go/src/net/http/request.go:1349
	_go_fuzz_dep_.CoverTab[41902]++
							return ""
//line /usr/local/go/src/net/http/request.go:1350
	// _ = "end of CoverTab[41902]"
}

// PostFormValue returns the first value for the named component of the POST,
//line /usr/local/go/src/net/http/request.go:1353
// PATCH, or PUT request body. URL query parameters are ignored.
//line /usr/local/go/src/net/http/request.go:1353
// PostFormValue calls ParseMultipartForm and ParseForm if necessary and ignores
//line /usr/local/go/src/net/http/request.go:1353
// any errors returned by these functions.
//line /usr/local/go/src/net/http/request.go:1353
// If key is not present, PostFormValue returns the empty string.
//line /usr/local/go/src/net/http/request.go:1358
func (r *Request) PostFormValue(key string) string {
//line /usr/local/go/src/net/http/request.go:1358
	_go_fuzz_dep_.CoverTab[41907]++
							if r.PostForm == nil {
//line /usr/local/go/src/net/http/request.go:1359
		_go_fuzz_dep_.CoverTab[41910]++
								r.ParseMultipartForm(defaultMaxMemory)
//line /usr/local/go/src/net/http/request.go:1360
		// _ = "end of CoverTab[41910]"
	} else {
//line /usr/local/go/src/net/http/request.go:1361
		_go_fuzz_dep_.CoverTab[41911]++
//line /usr/local/go/src/net/http/request.go:1361
		// _ = "end of CoverTab[41911]"
//line /usr/local/go/src/net/http/request.go:1361
	}
//line /usr/local/go/src/net/http/request.go:1361
	// _ = "end of CoverTab[41907]"
//line /usr/local/go/src/net/http/request.go:1361
	_go_fuzz_dep_.CoverTab[41908]++
							if vs := r.PostForm[key]; len(vs) > 0 {
//line /usr/local/go/src/net/http/request.go:1362
		_go_fuzz_dep_.CoverTab[41912]++
								return vs[0]
//line /usr/local/go/src/net/http/request.go:1363
		// _ = "end of CoverTab[41912]"
	} else {
//line /usr/local/go/src/net/http/request.go:1364
		_go_fuzz_dep_.CoverTab[41913]++
//line /usr/local/go/src/net/http/request.go:1364
		// _ = "end of CoverTab[41913]"
//line /usr/local/go/src/net/http/request.go:1364
	}
//line /usr/local/go/src/net/http/request.go:1364
	// _ = "end of CoverTab[41908]"
//line /usr/local/go/src/net/http/request.go:1364
	_go_fuzz_dep_.CoverTab[41909]++
							return ""
//line /usr/local/go/src/net/http/request.go:1365
	// _ = "end of CoverTab[41909]"
}

// FormFile returns the first file for the provided form key.
//line /usr/local/go/src/net/http/request.go:1368
// FormFile calls ParseMultipartForm and ParseForm if necessary.
//line /usr/local/go/src/net/http/request.go:1370
func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
//line /usr/local/go/src/net/http/request.go:1370
	_go_fuzz_dep_.CoverTab[41914]++
							if r.MultipartForm == multipartByReader {
//line /usr/local/go/src/net/http/request.go:1371
		_go_fuzz_dep_.CoverTab[41918]++
								return nil, nil, errors.New("http: multipart handled by MultipartReader")
//line /usr/local/go/src/net/http/request.go:1372
		// _ = "end of CoverTab[41918]"
	} else {
//line /usr/local/go/src/net/http/request.go:1373
		_go_fuzz_dep_.CoverTab[41919]++
//line /usr/local/go/src/net/http/request.go:1373
		// _ = "end of CoverTab[41919]"
//line /usr/local/go/src/net/http/request.go:1373
	}
//line /usr/local/go/src/net/http/request.go:1373
	// _ = "end of CoverTab[41914]"
//line /usr/local/go/src/net/http/request.go:1373
	_go_fuzz_dep_.CoverTab[41915]++
							if r.MultipartForm == nil {
//line /usr/local/go/src/net/http/request.go:1374
		_go_fuzz_dep_.CoverTab[41920]++
								err := r.ParseMultipartForm(defaultMaxMemory)
								if err != nil {
//line /usr/local/go/src/net/http/request.go:1376
			_go_fuzz_dep_.CoverTab[41921]++
									return nil, nil, err
//line /usr/local/go/src/net/http/request.go:1377
			// _ = "end of CoverTab[41921]"
		} else {
//line /usr/local/go/src/net/http/request.go:1378
			_go_fuzz_dep_.CoverTab[41922]++
//line /usr/local/go/src/net/http/request.go:1378
			// _ = "end of CoverTab[41922]"
//line /usr/local/go/src/net/http/request.go:1378
		}
//line /usr/local/go/src/net/http/request.go:1378
		// _ = "end of CoverTab[41920]"
	} else {
//line /usr/local/go/src/net/http/request.go:1379
		_go_fuzz_dep_.CoverTab[41923]++
//line /usr/local/go/src/net/http/request.go:1379
		// _ = "end of CoverTab[41923]"
//line /usr/local/go/src/net/http/request.go:1379
	}
//line /usr/local/go/src/net/http/request.go:1379
	// _ = "end of CoverTab[41915]"
//line /usr/local/go/src/net/http/request.go:1379
	_go_fuzz_dep_.CoverTab[41916]++
							if r.MultipartForm != nil && func() bool {
//line /usr/local/go/src/net/http/request.go:1380
		_go_fuzz_dep_.CoverTab[41924]++
//line /usr/local/go/src/net/http/request.go:1380
		return r.MultipartForm.File != nil
//line /usr/local/go/src/net/http/request.go:1380
		// _ = "end of CoverTab[41924]"
//line /usr/local/go/src/net/http/request.go:1380
	}() {
//line /usr/local/go/src/net/http/request.go:1380
		_go_fuzz_dep_.CoverTab[41925]++
								if fhs := r.MultipartForm.File[key]; len(fhs) > 0 {
//line /usr/local/go/src/net/http/request.go:1381
			_go_fuzz_dep_.CoverTab[41926]++
									f, err := fhs[0].Open()
									return f, fhs[0], err
//line /usr/local/go/src/net/http/request.go:1383
			// _ = "end of CoverTab[41926]"
		} else {
//line /usr/local/go/src/net/http/request.go:1384
			_go_fuzz_dep_.CoverTab[41927]++
//line /usr/local/go/src/net/http/request.go:1384
			// _ = "end of CoverTab[41927]"
//line /usr/local/go/src/net/http/request.go:1384
		}
//line /usr/local/go/src/net/http/request.go:1384
		// _ = "end of CoverTab[41925]"
	} else {
//line /usr/local/go/src/net/http/request.go:1385
		_go_fuzz_dep_.CoverTab[41928]++
//line /usr/local/go/src/net/http/request.go:1385
		// _ = "end of CoverTab[41928]"
//line /usr/local/go/src/net/http/request.go:1385
	}
//line /usr/local/go/src/net/http/request.go:1385
	// _ = "end of CoverTab[41916]"
//line /usr/local/go/src/net/http/request.go:1385
	_go_fuzz_dep_.CoverTab[41917]++
							return nil, nil, ErrMissingFile
//line /usr/local/go/src/net/http/request.go:1386
	// _ = "end of CoverTab[41917]"
}

func (r *Request) expectsContinue() bool {
//line /usr/local/go/src/net/http/request.go:1389
	_go_fuzz_dep_.CoverTab[41929]++
							return hasToken(r.Header.get("Expect"), "100-continue")
//line /usr/local/go/src/net/http/request.go:1390
	// _ = "end of CoverTab[41929]"
}

func (r *Request) wantsHttp10KeepAlive() bool {
//line /usr/local/go/src/net/http/request.go:1393
	_go_fuzz_dep_.CoverTab[41930]++
							if r.ProtoMajor != 1 || func() bool {
//line /usr/local/go/src/net/http/request.go:1394
		_go_fuzz_dep_.CoverTab[41932]++
//line /usr/local/go/src/net/http/request.go:1394
		return r.ProtoMinor != 0
//line /usr/local/go/src/net/http/request.go:1394
		// _ = "end of CoverTab[41932]"
//line /usr/local/go/src/net/http/request.go:1394
	}() {
//line /usr/local/go/src/net/http/request.go:1394
		_go_fuzz_dep_.CoverTab[41933]++
								return false
//line /usr/local/go/src/net/http/request.go:1395
		// _ = "end of CoverTab[41933]"
	} else {
//line /usr/local/go/src/net/http/request.go:1396
		_go_fuzz_dep_.CoverTab[41934]++
//line /usr/local/go/src/net/http/request.go:1396
		// _ = "end of CoverTab[41934]"
//line /usr/local/go/src/net/http/request.go:1396
	}
//line /usr/local/go/src/net/http/request.go:1396
	// _ = "end of CoverTab[41930]"
//line /usr/local/go/src/net/http/request.go:1396
	_go_fuzz_dep_.CoverTab[41931]++
							return hasToken(r.Header.get("Connection"), "keep-alive")
//line /usr/local/go/src/net/http/request.go:1397
	// _ = "end of CoverTab[41931]"
}

func (r *Request) wantsClose() bool {
//line /usr/local/go/src/net/http/request.go:1400
	_go_fuzz_dep_.CoverTab[41935]++
							if r.Close {
//line /usr/local/go/src/net/http/request.go:1401
		_go_fuzz_dep_.CoverTab[41937]++
								return true
//line /usr/local/go/src/net/http/request.go:1402
		// _ = "end of CoverTab[41937]"
	} else {
//line /usr/local/go/src/net/http/request.go:1403
		_go_fuzz_dep_.CoverTab[41938]++
//line /usr/local/go/src/net/http/request.go:1403
		// _ = "end of CoverTab[41938]"
//line /usr/local/go/src/net/http/request.go:1403
	}
//line /usr/local/go/src/net/http/request.go:1403
	// _ = "end of CoverTab[41935]"
//line /usr/local/go/src/net/http/request.go:1403
	_go_fuzz_dep_.CoverTab[41936]++
							return hasToken(r.Header.get("Connection"), "close")
//line /usr/local/go/src/net/http/request.go:1404
	// _ = "end of CoverTab[41936]"
}

func (r *Request) closeBody() error {
//line /usr/local/go/src/net/http/request.go:1407
	_go_fuzz_dep_.CoverTab[41939]++
							if r.Body == nil {
//line /usr/local/go/src/net/http/request.go:1408
		_go_fuzz_dep_.CoverTab[41941]++
								return nil
//line /usr/local/go/src/net/http/request.go:1409
		// _ = "end of CoverTab[41941]"
	} else {
//line /usr/local/go/src/net/http/request.go:1410
		_go_fuzz_dep_.CoverTab[41942]++
//line /usr/local/go/src/net/http/request.go:1410
		// _ = "end of CoverTab[41942]"
//line /usr/local/go/src/net/http/request.go:1410
	}
//line /usr/local/go/src/net/http/request.go:1410
	// _ = "end of CoverTab[41939]"
//line /usr/local/go/src/net/http/request.go:1410
	_go_fuzz_dep_.CoverTab[41940]++
							return r.Body.Close()
//line /usr/local/go/src/net/http/request.go:1411
	// _ = "end of CoverTab[41940]"
}

func (r *Request) isReplayable() bool {
//line /usr/local/go/src/net/http/request.go:1414
	_go_fuzz_dep_.CoverTab[41943]++
							if r.Body == nil || func() bool {
//line /usr/local/go/src/net/http/request.go:1415
		_go_fuzz_dep_.CoverTab[41945]++
//line /usr/local/go/src/net/http/request.go:1415
		return r.Body == NoBody
//line /usr/local/go/src/net/http/request.go:1415
		// _ = "end of CoverTab[41945]"
//line /usr/local/go/src/net/http/request.go:1415
	}() || func() bool {
//line /usr/local/go/src/net/http/request.go:1415
		_go_fuzz_dep_.CoverTab[41946]++
//line /usr/local/go/src/net/http/request.go:1415
		return r.GetBody != nil
//line /usr/local/go/src/net/http/request.go:1415
		// _ = "end of CoverTab[41946]"
//line /usr/local/go/src/net/http/request.go:1415
	}() {
//line /usr/local/go/src/net/http/request.go:1415
		_go_fuzz_dep_.CoverTab[41947]++
								switch valueOrDefault(r.Method, "GET") {
		case "GET", "HEAD", "OPTIONS", "TRACE":
//line /usr/local/go/src/net/http/request.go:1417
			_go_fuzz_dep_.CoverTab[41949]++
									return true
//line /usr/local/go/src/net/http/request.go:1418
			// _ = "end of CoverTab[41949]"
//line /usr/local/go/src/net/http/request.go:1418
		default:
//line /usr/local/go/src/net/http/request.go:1418
			_go_fuzz_dep_.CoverTab[41950]++
//line /usr/local/go/src/net/http/request.go:1418
			// _ = "end of CoverTab[41950]"
		}
//line /usr/local/go/src/net/http/request.go:1419
		// _ = "end of CoverTab[41947]"
//line /usr/local/go/src/net/http/request.go:1419
		_go_fuzz_dep_.CoverTab[41948]++

//line /usr/local/go/src/net/http/request.go:1423
		if r.Header.has("Idempotency-Key") || func() bool {
//line /usr/local/go/src/net/http/request.go:1423
			_go_fuzz_dep_.CoverTab[41951]++
//line /usr/local/go/src/net/http/request.go:1423
			return r.Header.has("X-Idempotency-Key")
//line /usr/local/go/src/net/http/request.go:1423
			// _ = "end of CoverTab[41951]"
//line /usr/local/go/src/net/http/request.go:1423
		}() {
//line /usr/local/go/src/net/http/request.go:1423
			_go_fuzz_dep_.CoverTab[41952]++
									return true
//line /usr/local/go/src/net/http/request.go:1424
			// _ = "end of CoverTab[41952]"
		} else {
//line /usr/local/go/src/net/http/request.go:1425
			_go_fuzz_dep_.CoverTab[41953]++
//line /usr/local/go/src/net/http/request.go:1425
			// _ = "end of CoverTab[41953]"
//line /usr/local/go/src/net/http/request.go:1425
		}
//line /usr/local/go/src/net/http/request.go:1425
		// _ = "end of CoverTab[41948]"
	} else {
//line /usr/local/go/src/net/http/request.go:1426
		_go_fuzz_dep_.CoverTab[41954]++
//line /usr/local/go/src/net/http/request.go:1426
		// _ = "end of CoverTab[41954]"
//line /usr/local/go/src/net/http/request.go:1426
	}
//line /usr/local/go/src/net/http/request.go:1426
	// _ = "end of CoverTab[41943]"
//line /usr/local/go/src/net/http/request.go:1426
	_go_fuzz_dep_.CoverTab[41944]++
							return false
//line /usr/local/go/src/net/http/request.go:1427
	// _ = "end of CoverTab[41944]"
}

// outgoingLength reports the Content-Length of this outgoing (Client) request.
//line /usr/local/go/src/net/http/request.go:1430
// It maps 0 into -1 (unknown) when the Body is non-nil.
//line /usr/local/go/src/net/http/request.go:1432
func (r *Request) outgoingLength() int64 {
//line /usr/local/go/src/net/http/request.go:1432
	_go_fuzz_dep_.CoverTab[41955]++
							if r.Body == nil || func() bool {
//line /usr/local/go/src/net/http/request.go:1433
		_go_fuzz_dep_.CoverTab[41958]++
//line /usr/local/go/src/net/http/request.go:1433
		return r.Body == NoBody
//line /usr/local/go/src/net/http/request.go:1433
		// _ = "end of CoverTab[41958]"
//line /usr/local/go/src/net/http/request.go:1433
	}() {
//line /usr/local/go/src/net/http/request.go:1433
		_go_fuzz_dep_.CoverTab[41959]++
								return 0
//line /usr/local/go/src/net/http/request.go:1434
		// _ = "end of CoverTab[41959]"
	} else {
//line /usr/local/go/src/net/http/request.go:1435
		_go_fuzz_dep_.CoverTab[41960]++
//line /usr/local/go/src/net/http/request.go:1435
		// _ = "end of CoverTab[41960]"
//line /usr/local/go/src/net/http/request.go:1435
	}
//line /usr/local/go/src/net/http/request.go:1435
	// _ = "end of CoverTab[41955]"
//line /usr/local/go/src/net/http/request.go:1435
	_go_fuzz_dep_.CoverTab[41956]++
							if r.ContentLength != 0 {
//line /usr/local/go/src/net/http/request.go:1436
		_go_fuzz_dep_.CoverTab[41961]++
								return r.ContentLength
//line /usr/local/go/src/net/http/request.go:1437
		// _ = "end of CoverTab[41961]"
	} else {
//line /usr/local/go/src/net/http/request.go:1438
		_go_fuzz_dep_.CoverTab[41962]++
//line /usr/local/go/src/net/http/request.go:1438
		// _ = "end of CoverTab[41962]"
//line /usr/local/go/src/net/http/request.go:1438
	}
//line /usr/local/go/src/net/http/request.go:1438
	// _ = "end of CoverTab[41956]"
//line /usr/local/go/src/net/http/request.go:1438
	_go_fuzz_dep_.CoverTab[41957]++
							return -1
//line /usr/local/go/src/net/http/request.go:1439
	// _ = "end of CoverTab[41957]"
}

// requestMethodUsuallyLacksBody reports whether the given request
//line /usr/local/go/src/net/http/request.go:1442
// method is one that typically does not involve a request body.
//line /usr/local/go/src/net/http/request.go:1442
// This is used by the Transport (via
//line /usr/local/go/src/net/http/request.go:1442
// transferWriter.shouldSendChunkedRequestBody) to determine whether
//line /usr/local/go/src/net/http/request.go:1442
// we try to test-read a byte from a non-nil Request.Body when
//line /usr/local/go/src/net/http/request.go:1442
// Request.outgoingLength() returns -1. See the comments in
//line /usr/local/go/src/net/http/request.go:1442
// shouldSendChunkedRequestBody.
//line /usr/local/go/src/net/http/request.go:1449
func requestMethodUsuallyLacksBody(method string) bool {
//line /usr/local/go/src/net/http/request.go:1449
	_go_fuzz_dep_.CoverTab[41963]++
							switch method {
	case "GET", "HEAD", "DELETE", "OPTIONS", "PROPFIND", "SEARCH":
//line /usr/local/go/src/net/http/request.go:1451
		_go_fuzz_dep_.CoverTab[41965]++
								return true
//line /usr/local/go/src/net/http/request.go:1452
		// _ = "end of CoverTab[41965]"
//line /usr/local/go/src/net/http/request.go:1452
	default:
//line /usr/local/go/src/net/http/request.go:1452
		_go_fuzz_dep_.CoverTab[41966]++
//line /usr/local/go/src/net/http/request.go:1452
		// _ = "end of CoverTab[41966]"
	}
//line /usr/local/go/src/net/http/request.go:1453
	// _ = "end of CoverTab[41963]"
//line /usr/local/go/src/net/http/request.go:1453
	_go_fuzz_dep_.CoverTab[41964]++
							return false
//line /usr/local/go/src/net/http/request.go:1454
	// _ = "end of CoverTab[41964]"
}

// requiresHTTP1 reports whether this request requires being sent on
//line /usr/local/go/src/net/http/request.go:1457
// an HTTP/1 connection.
//line /usr/local/go/src/net/http/request.go:1459
func (r *Request) requiresHTTP1() bool {
//line /usr/local/go/src/net/http/request.go:1459
	_go_fuzz_dep_.CoverTab[41967]++
							return hasToken(r.Header.Get("Connection"), "upgrade") && func() bool {
//line /usr/local/go/src/net/http/request.go:1460
		_go_fuzz_dep_.CoverTab[41968]++
//line /usr/local/go/src/net/http/request.go:1460
		return ascii.EqualFold(r.Header.Get("Upgrade"), "websocket")
								// _ = "end of CoverTab[41968]"
//line /usr/local/go/src/net/http/request.go:1461
	}()
//line /usr/local/go/src/net/http/request.go:1461
	// _ = "end of CoverTab[41967]"
}

//line /usr/local/go/src/net/http/request.go:1462
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/request.go:1462
var _ = _go_fuzz_dep_.CoverTab
