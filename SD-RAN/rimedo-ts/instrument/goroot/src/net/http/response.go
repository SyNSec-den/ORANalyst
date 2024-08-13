// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP Response reading and parsing.

//line /usr/local/go/src/net/http/response.go:7
package http

//line /usr/local/go/src/net/http/response.go:7
import (
//line /usr/local/go/src/net/http/response.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/response.go:7
)
//line /usr/local/go/src/net/http/response.go:7
import (
//line /usr/local/go/src/net/http/response.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/response.go:7
)

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/textproto"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/net/http/httpguts"
)

var respExcludeHeader = map[string]bool{
	"Content-Length":	true,
	"Transfer-Encoding":	true,
	"Trailer":		true,
}

// Response represents the response from an HTTP request.
//line /usr/local/go/src/net/http/response.go:30
//
//line /usr/local/go/src/net/http/response.go:30
// The Client and Transport return Responses from servers once
//line /usr/local/go/src/net/http/response.go:30
// the response headers have been received. The response body
//line /usr/local/go/src/net/http/response.go:30
// is streamed on demand as the Body field is read.
//line /usr/local/go/src/net/http/response.go:35
type Response struct {
	Status		string	// e.g. "200 OK"
	StatusCode	int	// e.g. 200
	Proto		string	// e.g. "HTTP/1.0"
	ProtoMajor	int	// e.g. 1
	ProtoMinor	int	// e.g. 0

	// Header maps header keys to values. If the response had multiple
	// headers with the same key, they may be concatenated, with comma
	// delimiters.  (RFC 7230, section 3.2.2 requires that multiple headers
	// be semantically equivalent to a comma-delimited sequence.) When
	// Header values are duplicated by other fields in this struct (e.g.,
	// ContentLength, TransferEncoding, Trailer), the field values are
	// authoritative.
	//
	// Keys in the map are canonicalized (see CanonicalHeaderKey).
	Header	Header

	// Body represents the response body.
	//
	// The response body is streamed on demand as the Body field
	// is read. If the network connection fails or the server
	// terminates the response, Body.Read calls return an error.
	//
	// The http Client and Transport guarantee that Body is always
	// non-nil, even on responses without a body or responses with
	// a zero-length body. It is the caller's responsibility to
	// close Body. The default HTTP client's Transport may not
	// reuse HTTP/1.x "keep-alive" TCP connections if the Body is
	// not read to completion and closed.
	//
	// The Body is automatically dechunked if the server replied
	// with a "chunked" Transfer-Encoding.
	//
	// As of Go 1.12, the Body will also implement io.Writer
	// on a successful "101 Switching Protocols" response,
	// as used by WebSockets and HTTP/2's "h2c" mode.
	Body	io.ReadCloser

	// ContentLength records the length of the associated content. The
	// value -1 indicates that the length is unknown. Unless Request.Method
	// is "HEAD", values >= 0 indicate that the given number of bytes may
	// be read from Body.
	ContentLength	int64

	// Contains transfer encodings from outer-most to inner-most. Value is
	// nil, means that "identity" encoding is used.
	TransferEncoding	[]string

	// Close records whether the header directed that the connection be
	// closed after reading Body. The value is advice for clients: neither
	// ReadResponse nor Response.Write ever closes a connection.
	Close	bool

	// Uncompressed reports whether the response was sent compressed but
	// was decompressed by the http package. When true, reading from
	// Body yields the uncompressed content instead of the compressed
	// content actually set from the server, ContentLength is set to -1,
	// and the "Content-Length" and "Content-Encoding" fields are deleted
	// from the responseHeader. To get the original response from
	// the server, set Transport.DisableCompression to true.
	Uncompressed	bool

	// Trailer maps trailer keys to values in the same
	// format as Header.
	//
	// The Trailer initially contains only nil values, one for
	// each key specified in the server's "Trailer" header
	// value. Those values are not added to Header.
	//
	// Trailer must not be accessed concurrently with Read calls
	// on the Body.
	//
	// After Body.Read has returned io.EOF, Trailer will contain
	// any trailer values sent by the server.
	Trailer	Header

	// Request is the request that was sent to obtain this Response.
	// Request's Body is nil (having already been consumed).
	// This is only populated for Client requests.
	Request	*Request

	// TLS contains information about the TLS connection on which the
	// response was received. It is nil for unencrypted responses.
	// The pointer is shared between responses and should not be
	// modified.
	TLS	*tls.ConnectionState
}

// Cookies parses and returns the cookies set in the Set-Cookie headers.
func (r *Response) Cookies() []*Cookie {
//line /usr/local/go/src/net/http/response.go:125
	_go_fuzz_dep_.CoverTab[41969]++
							return readSetCookies(r.Header)
//line /usr/local/go/src/net/http/response.go:126
	// _ = "end of CoverTab[41969]"
}

// ErrNoLocation is returned by Response's Location method
//line /usr/local/go/src/net/http/response.go:129
// when no Location header is present.
//line /usr/local/go/src/net/http/response.go:131
var ErrNoLocation = errors.New("http: no Location header in response")

// Location returns the URL of the response's "Location" header,
//line /usr/local/go/src/net/http/response.go:133
// if present. Relative redirects are resolved relative to
//line /usr/local/go/src/net/http/response.go:133
// the Response's Request. ErrNoLocation is returned if no
//line /usr/local/go/src/net/http/response.go:133
// Location header is present.
//line /usr/local/go/src/net/http/response.go:137
func (r *Response) Location() (*url.URL, error) {
//line /usr/local/go/src/net/http/response.go:137
	_go_fuzz_dep_.CoverTab[41970]++
							lv := r.Header.Get("Location")
							if lv == "" {
//line /usr/local/go/src/net/http/response.go:139
		_go_fuzz_dep_.CoverTab[41973]++
								return nil, ErrNoLocation
//line /usr/local/go/src/net/http/response.go:140
		// _ = "end of CoverTab[41973]"
	} else {
//line /usr/local/go/src/net/http/response.go:141
		_go_fuzz_dep_.CoverTab[41974]++
//line /usr/local/go/src/net/http/response.go:141
		// _ = "end of CoverTab[41974]"
//line /usr/local/go/src/net/http/response.go:141
	}
//line /usr/local/go/src/net/http/response.go:141
	// _ = "end of CoverTab[41970]"
//line /usr/local/go/src/net/http/response.go:141
	_go_fuzz_dep_.CoverTab[41971]++
							if r.Request != nil && func() bool {
//line /usr/local/go/src/net/http/response.go:142
		_go_fuzz_dep_.CoverTab[41975]++
//line /usr/local/go/src/net/http/response.go:142
		return r.Request.URL != nil
//line /usr/local/go/src/net/http/response.go:142
		// _ = "end of CoverTab[41975]"
//line /usr/local/go/src/net/http/response.go:142
	}() {
//line /usr/local/go/src/net/http/response.go:142
		_go_fuzz_dep_.CoverTab[41976]++
								return r.Request.URL.Parse(lv)
//line /usr/local/go/src/net/http/response.go:143
		// _ = "end of CoverTab[41976]"
	} else {
//line /usr/local/go/src/net/http/response.go:144
		_go_fuzz_dep_.CoverTab[41977]++
//line /usr/local/go/src/net/http/response.go:144
		// _ = "end of CoverTab[41977]"
//line /usr/local/go/src/net/http/response.go:144
	}
//line /usr/local/go/src/net/http/response.go:144
	// _ = "end of CoverTab[41971]"
//line /usr/local/go/src/net/http/response.go:144
	_go_fuzz_dep_.CoverTab[41972]++
							return url.Parse(lv)
//line /usr/local/go/src/net/http/response.go:145
	// _ = "end of CoverTab[41972]"
}

// ReadResponse reads and returns an HTTP response from r.
//line /usr/local/go/src/net/http/response.go:148
// The req parameter optionally specifies the Request that corresponds
//line /usr/local/go/src/net/http/response.go:148
// to this Response. If nil, a GET request is assumed.
//line /usr/local/go/src/net/http/response.go:148
// Clients must call resp.Body.Close when finished reading resp.Body.
//line /usr/local/go/src/net/http/response.go:148
// After that call, clients can inspect resp.Trailer to find key/value
//line /usr/local/go/src/net/http/response.go:148
// pairs included in the response trailer.
//line /usr/local/go/src/net/http/response.go:154
func ReadResponse(r *bufio.Reader, req *Request) (*Response, error) {
//line /usr/local/go/src/net/http/response.go:154
	_go_fuzz_dep_.CoverTab[41978]++
							tp := textproto.NewReader(r)
							resp := &Response{
		Request: req,
	}

//line /usr/local/go/src/net/http/response.go:161
	line, err := tp.ReadLine()
	if err != nil {
//line /usr/local/go/src/net/http/response.go:162
		_go_fuzz_dep_.CoverTab[41986]++
								if err == io.EOF {
//line /usr/local/go/src/net/http/response.go:163
			_go_fuzz_dep_.CoverTab[41988]++
									err = io.ErrUnexpectedEOF
//line /usr/local/go/src/net/http/response.go:164
			// _ = "end of CoverTab[41988]"
		} else {
//line /usr/local/go/src/net/http/response.go:165
			_go_fuzz_dep_.CoverTab[41989]++
//line /usr/local/go/src/net/http/response.go:165
			// _ = "end of CoverTab[41989]"
//line /usr/local/go/src/net/http/response.go:165
		}
//line /usr/local/go/src/net/http/response.go:165
		// _ = "end of CoverTab[41986]"
//line /usr/local/go/src/net/http/response.go:165
		_go_fuzz_dep_.CoverTab[41987]++
								return nil, err
//line /usr/local/go/src/net/http/response.go:166
		// _ = "end of CoverTab[41987]"
	} else {
//line /usr/local/go/src/net/http/response.go:167
		_go_fuzz_dep_.CoverTab[41990]++
//line /usr/local/go/src/net/http/response.go:167
		// _ = "end of CoverTab[41990]"
//line /usr/local/go/src/net/http/response.go:167
	}
//line /usr/local/go/src/net/http/response.go:167
	// _ = "end of CoverTab[41978]"
//line /usr/local/go/src/net/http/response.go:167
	_go_fuzz_dep_.CoverTab[41979]++
							proto, status, ok := strings.Cut(line, " ")
							if !ok {
//line /usr/local/go/src/net/http/response.go:169
		_go_fuzz_dep_.CoverTab[41991]++
								return nil, badStringError("malformed HTTP response", line)
//line /usr/local/go/src/net/http/response.go:170
		// _ = "end of CoverTab[41991]"
	} else {
//line /usr/local/go/src/net/http/response.go:171
		_go_fuzz_dep_.CoverTab[41992]++
//line /usr/local/go/src/net/http/response.go:171
		// _ = "end of CoverTab[41992]"
//line /usr/local/go/src/net/http/response.go:171
	}
//line /usr/local/go/src/net/http/response.go:171
	// _ = "end of CoverTab[41979]"
//line /usr/local/go/src/net/http/response.go:171
	_go_fuzz_dep_.CoverTab[41980]++
							resp.Proto = proto
							resp.Status = strings.TrimLeft(status, " ")

							statusCode, _, _ := strings.Cut(resp.Status, " ")
							if len(statusCode) != 3 {
//line /usr/local/go/src/net/http/response.go:176
		_go_fuzz_dep_.CoverTab[41993]++
								return nil, badStringError("malformed HTTP status code", statusCode)
//line /usr/local/go/src/net/http/response.go:177
		// _ = "end of CoverTab[41993]"
	} else {
//line /usr/local/go/src/net/http/response.go:178
		_go_fuzz_dep_.CoverTab[41994]++
//line /usr/local/go/src/net/http/response.go:178
		// _ = "end of CoverTab[41994]"
//line /usr/local/go/src/net/http/response.go:178
	}
//line /usr/local/go/src/net/http/response.go:178
	// _ = "end of CoverTab[41980]"
//line /usr/local/go/src/net/http/response.go:178
	_go_fuzz_dep_.CoverTab[41981]++
							resp.StatusCode, err = strconv.Atoi(statusCode)
							if err != nil || func() bool {
//line /usr/local/go/src/net/http/response.go:180
		_go_fuzz_dep_.CoverTab[41995]++
//line /usr/local/go/src/net/http/response.go:180
		return resp.StatusCode < 0
//line /usr/local/go/src/net/http/response.go:180
		// _ = "end of CoverTab[41995]"
//line /usr/local/go/src/net/http/response.go:180
	}() {
//line /usr/local/go/src/net/http/response.go:180
		_go_fuzz_dep_.CoverTab[41996]++
								return nil, badStringError("malformed HTTP status code", statusCode)
//line /usr/local/go/src/net/http/response.go:181
		// _ = "end of CoverTab[41996]"
	} else {
//line /usr/local/go/src/net/http/response.go:182
		_go_fuzz_dep_.CoverTab[41997]++
//line /usr/local/go/src/net/http/response.go:182
		// _ = "end of CoverTab[41997]"
//line /usr/local/go/src/net/http/response.go:182
	}
//line /usr/local/go/src/net/http/response.go:182
	// _ = "end of CoverTab[41981]"
//line /usr/local/go/src/net/http/response.go:182
	_go_fuzz_dep_.CoverTab[41982]++
							if resp.ProtoMajor, resp.ProtoMinor, ok = ParseHTTPVersion(resp.Proto); !ok {
//line /usr/local/go/src/net/http/response.go:183
		_go_fuzz_dep_.CoverTab[41998]++
								return nil, badStringError("malformed HTTP version", resp.Proto)
//line /usr/local/go/src/net/http/response.go:184
		// _ = "end of CoverTab[41998]"
	} else {
//line /usr/local/go/src/net/http/response.go:185
		_go_fuzz_dep_.CoverTab[41999]++
//line /usr/local/go/src/net/http/response.go:185
		// _ = "end of CoverTab[41999]"
//line /usr/local/go/src/net/http/response.go:185
	}
//line /usr/local/go/src/net/http/response.go:185
	// _ = "end of CoverTab[41982]"
//line /usr/local/go/src/net/http/response.go:185
	_go_fuzz_dep_.CoverTab[41983]++

//line /usr/local/go/src/net/http/response.go:188
	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
//line /usr/local/go/src/net/http/response.go:189
		_go_fuzz_dep_.CoverTab[42000]++
								if err == io.EOF {
//line /usr/local/go/src/net/http/response.go:190
			_go_fuzz_dep_.CoverTab[42002]++
									err = io.ErrUnexpectedEOF
//line /usr/local/go/src/net/http/response.go:191
			// _ = "end of CoverTab[42002]"
		} else {
//line /usr/local/go/src/net/http/response.go:192
			_go_fuzz_dep_.CoverTab[42003]++
//line /usr/local/go/src/net/http/response.go:192
			// _ = "end of CoverTab[42003]"
//line /usr/local/go/src/net/http/response.go:192
		}
//line /usr/local/go/src/net/http/response.go:192
		// _ = "end of CoverTab[42000]"
//line /usr/local/go/src/net/http/response.go:192
		_go_fuzz_dep_.CoverTab[42001]++
								return nil, err
//line /usr/local/go/src/net/http/response.go:193
		// _ = "end of CoverTab[42001]"
	} else {
//line /usr/local/go/src/net/http/response.go:194
		_go_fuzz_dep_.CoverTab[42004]++
//line /usr/local/go/src/net/http/response.go:194
		// _ = "end of CoverTab[42004]"
//line /usr/local/go/src/net/http/response.go:194
	}
//line /usr/local/go/src/net/http/response.go:194
	// _ = "end of CoverTab[41983]"
//line /usr/local/go/src/net/http/response.go:194
	_go_fuzz_dep_.CoverTab[41984]++
							resp.Header = Header(mimeHeader)

							fixPragmaCacheControl(resp.Header)

							err = readTransfer(resp, r)
							if err != nil {
//line /usr/local/go/src/net/http/response.go:200
		_go_fuzz_dep_.CoverTab[42005]++
								return nil, err
//line /usr/local/go/src/net/http/response.go:201
		// _ = "end of CoverTab[42005]"
	} else {
//line /usr/local/go/src/net/http/response.go:202
		_go_fuzz_dep_.CoverTab[42006]++
//line /usr/local/go/src/net/http/response.go:202
		// _ = "end of CoverTab[42006]"
//line /usr/local/go/src/net/http/response.go:202
	}
//line /usr/local/go/src/net/http/response.go:202
	// _ = "end of CoverTab[41984]"
//line /usr/local/go/src/net/http/response.go:202
	_go_fuzz_dep_.CoverTab[41985]++

							return resp, nil
//line /usr/local/go/src/net/http/response.go:204
	// _ = "end of CoverTab[41985]"
}

// RFC 7234, section 5.4: Should treat
//line /usr/local/go/src/net/http/response.go:207
//
//line /usr/local/go/src/net/http/response.go:207
//	Pragma: no-cache
//line /usr/local/go/src/net/http/response.go:207
//
//line /usr/local/go/src/net/http/response.go:207
// like
//line /usr/local/go/src/net/http/response.go:207
//
//line /usr/local/go/src/net/http/response.go:207
//	Cache-Control: no-cache
//line /usr/local/go/src/net/http/response.go:214
func fixPragmaCacheControl(header Header) {
//line /usr/local/go/src/net/http/response.go:214
	_go_fuzz_dep_.CoverTab[42007]++
							if hp, ok := header["Pragma"]; ok && func() bool {
//line /usr/local/go/src/net/http/response.go:215
		_go_fuzz_dep_.CoverTab[42008]++
//line /usr/local/go/src/net/http/response.go:215
		return len(hp) > 0
//line /usr/local/go/src/net/http/response.go:215
		// _ = "end of CoverTab[42008]"
//line /usr/local/go/src/net/http/response.go:215
	}() && func() bool {
//line /usr/local/go/src/net/http/response.go:215
		_go_fuzz_dep_.CoverTab[42009]++
//line /usr/local/go/src/net/http/response.go:215
		return hp[0] == "no-cache"
//line /usr/local/go/src/net/http/response.go:215
		// _ = "end of CoverTab[42009]"
//line /usr/local/go/src/net/http/response.go:215
	}() {
//line /usr/local/go/src/net/http/response.go:215
		_go_fuzz_dep_.CoverTab[42010]++
								if _, presentcc := header["Cache-Control"]; !presentcc {
//line /usr/local/go/src/net/http/response.go:216
			_go_fuzz_dep_.CoverTab[42011]++
									header["Cache-Control"] = []string{"no-cache"}
//line /usr/local/go/src/net/http/response.go:217
			// _ = "end of CoverTab[42011]"
		} else {
//line /usr/local/go/src/net/http/response.go:218
			_go_fuzz_dep_.CoverTab[42012]++
//line /usr/local/go/src/net/http/response.go:218
			// _ = "end of CoverTab[42012]"
//line /usr/local/go/src/net/http/response.go:218
		}
//line /usr/local/go/src/net/http/response.go:218
		// _ = "end of CoverTab[42010]"
	} else {
//line /usr/local/go/src/net/http/response.go:219
		_go_fuzz_dep_.CoverTab[42013]++
//line /usr/local/go/src/net/http/response.go:219
		// _ = "end of CoverTab[42013]"
//line /usr/local/go/src/net/http/response.go:219
	}
//line /usr/local/go/src/net/http/response.go:219
	// _ = "end of CoverTab[42007]"
}

// ProtoAtLeast reports whether the HTTP protocol used
//line /usr/local/go/src/net/http/response.go:222
// in the response is at least major.minor.
//line /usr/local/go/src/net/http/response.go:224
func (r *Response) ProtoAtLeast(major, minor int) bool {
//line /usr/local/go/src/net/http/response.go:224
	_go_fuzz_dep_.CoverTab[42014]++
							return r.ProtoMajor > major || func() bool {
//line /usr/local/go/src/net/http/response.go:225
		_go_fuzz_dep_.CoverTab[42015]++
//line /usr/local/go/src/net/http/response.go:225
		return r.ProtoMajor == major && func() bool {
									_go_fuzz_dep_.CoverTab[42016]++
//line /usr/local/go/src/net/http/response.go:226
			return r.ProtoMinor >= minor
//line /usr/local/go/src/net/http/response.go:226
			// _ = "end of CoverTab[42016]"
//line /usr/local/go/src/net/http/response.go:226
		}()
//line /usr/local/go/src/net/http/response.go:226
		// _ = "end of CoverTab[42015]"
//line /usr/local/go/src/net/http/response.go:226
	}()
//line /usr/local/go/src/net/http/response.go:226
	// _ = "end of CoverTab[42014]"
}

// Write writes r to w in the HTTP/1.x server response format,
//line /usr/local/go/src/net/http/response.go:229
// including the status line, headers, body, and optional trailer.
//line /usr/local/go/src/net/http/response.go:229
//
//line /usr/local/go/src/net/http/response.go:229
// This method consults the following fields of the response r:
//line /usr/local/go/src/net/http/response.go:229
//
//line /usr/local/go/src/net/http/response.go:229
//	StatusCode
//line /usr/local/go/src/net/http/response.go:229
//	ProtoMajor
//line /usr/local/go/src/net/http/response.go:229
//	ProtoMinor
//line /usr/local/go/src/net/http/response.go:229
//	Request.Method
//line /usr/local/go/src/net/http/response.go:229
//	TransferEncoding
//line /usr/local/go/src/net/http/response.go:229
//	Trailer
//line /usr/local/go/src/net/http/response.go:229
//	Body
//line /usr/local/go/src/net/http/response.go:229
//	ContentLength
//line /usr/local/go/src/net/http/response.go:229
//	Header, values for non-canonical keys will have unpredictable behavior
//line /usr/local/go/src/net/http/response.go:229
//
//line /usr/local/go/src/net/http/response.go:229
// The Response Body is closed after it is sent.
//line /usr/local/go/src/net/http/response.go:245
func (r *Response) Write(w io.Writer) error {
//line /usr/local/go/src/net/http/response.go:245
	_go_fuzz_dep_.CoverTab[42017]++

							text := r.Status
							if text == "" {
//line /usr/local/go/src/net/http/response.go:248
		_go_fuzz_dep_.CoverTab[42028]++
								text = StatusText(r.StatusCode)
								if text == "" {
//line /usr/local/go/src/net/http/response.go:250
			_go_fuzz_dep_.CoverTab[42029]++
									text = "status code " + strconv.Itoa(r.StatusCode)
//line /usr/local/go/src/net/http/response.go:251
			// _ = "end of CoverTab[42029]"
		} else {
//line /usr/local/go/src/net/http/response.go:252
			_go_fuzz_dep_.CoverTab[42030]++
//line /usr/local/go/src/net/http/response.go:252
			// _ = "end of CoverTab[42030]"
//line /usr/local/go/src/net/http/response.go:252
		}
//line /usr/local/go/src/net/http/response.go:252
		// _ = "end of CoverTab[42028]"
	} else {
//line /usr/local/go/src/net/http/response.go:253
		_go_fuzz_dep_.CoverTab[42031]++

//line /usr/local/go/src/net/http/response.go:256
		text = strings.TrimPrefix(text, strconv.Itoa(r.StatusCode)+" ")
//line /usr/local/go/src/net/http/response.go:256
		// _ = "end of CoverTab[42031]"
	}
//line /usr/local/go/src/net/http/response.go:257
	// _ = "end of CoverTab[42017]"
//line /usr/local/go/src/net/http/response.go:257
	_go_fuzz_dep_.CoverTab[42018]++

							if _, err := fmt.Fprintf(w, "HTTP/%d.%d %03d %s\r\n", r.ProtoMajor, r.ProtoMinor, r.StatusCode, text); err != nil {
//line /usr/local/go/src/net/http/response.go:259
		_go_fuzz_dep_.CoverTab[42032]++
								return err
//line /usr/local/go/src/net/http/response.go:260
		// _ = "end of CoverTab[42032]"
	} else {
//line /usr/local/go/src/net/http/response.go:261
		_go_fuzz_dep_.CoverTab[42033]++
//line /usr/local/go/src/net/http/response.go:261
		// _ = "end of CoverTab[42033]"
//line /usr/local/go/src/net/http/response.go:261
	}
//line /usr/local/go/src/net/http/response.go:261
	// _ = "end of CoverTab[42018]"
//line /usr/local/go/src/net/http/response.go:261
	_go_fuzz_dep_.CoverTab[42019]++

//line /usr/local/go/src/net/http/response.go:264
	r1 := new(Response)
	*r1 = *r
	if r1.ContentLength == 0 && func() bool {
//line /usr/local/go/src/net/http/response.go:266
		_go_fuzz_dep_.CoverTab[42034]++
//line /usr/local/go/src/net/http/response.go:266
		return r1.Body != nil
//line /usr/local/go/src/net/http/response.go:266
		// _ = "end of CoverTab[42034]"
//line /usr/local/go/src/net/http/response.go:266
	}() {
//line /usr/local/go/src/net/http/response.go:266
		_go_fuzz_dep_.CoverTab[42035]++
		// Is it actually 0 length? Or just unknown?
		var buf [1]byte
		n, err := r1.Body.Read(buf[:])
		if err != nil && func() bool {
//line /usr/local/go/src/net/http/response.go:270
			_go_fuzz_dep_.CoverTab[42037]++
//line /usr/local/go/src/net/http/response.go:270
			return err != io.EOF
//line /usr/local/go/src/net/http/response.go:270
			// _ = "end of CoverTab[42037]"
//line /usr/local/go/src/net/http/response.go:270
		}() {
//line /usr/local/go/src/net/http/response.go:270
			_go_fuzz_dep_.CoverTab[42038]++
									return err
//line /usr/local/go/src/net/http/response.go:271
			// _ = "end of CoverTab[42038]"
		} else {
//line /usr/local/go/src/net/http/response.go:272
			_go_fuzz_dep_.CoverTab[42039]++
//line /usr/local/go/src/net/http/response.go:272
			// _ = "end of CoverTab[42039]"
//line /usr/local/go/src/net/http/response.go:272
		}
//line /usr/local/go/src/net/http/response.go:272
		// _ = "end of CoverTab[42035]"
//line /usr/local/go/src/net/http/response.go:272
		_go_fuzz_dep_.CoverTab[42036]++
								if n == 0 {
//line /usr/local/go/src/net/http/response.go:273
			_go_fuzz_dep_.CoverTab[42040]++

//line /usr/local/go/src/net/http/response.go:276
			r1.Body = NoBody
//line /usr/local/go/src/net/http/response.go:276
			// _ = "end of CoverTab[42040]"
		} else {
//line /usr/local/go/src/net/http/response.go:277
			_go_fuzz_dep_.CoverTab[42041]++
									r1.ContentLength = -1
									r1.Body = struct {
				io.Reader
				io.Closer
			}{
				io.MultiReader(bytes.NewReader(buf[:1]), r.Body),
				r.Body,
			}
//line /usr/local/go/src/net/http/response.go:285
			// _ = "end of CoverTab[42041]"
		}
//line /usr/local/go/src/net/http/response.go:286
		// _ = "end of CoverTab[42036]"
	} else {
//line /usr/local/go/src/net/http/response.go:287
		_go_fuzz_dep_.CoverTab[42042]++
//line /usr/local/go/src/net/http/response.go:287
		// _ = "end of CoverTab[42042]"
//line /usr/local/go/src/net/http/response.go:287
	}
//line /usr/local/go/src/net/http/response.go:287
	// _ = "end of CoverTab[42019]"
//line /usr/local/go/src/net/http/response.go:287
	_go_fuzz_dep_.CoverTab[42020]++

//line /usr/local/go/src/net/http/response.go:292
	if r1.ContentLength == -1 && func() bool {
//line /usr/local/go/src/net/http/response.go:292
		_go_fuzz_dep_.CoverTab[42043]++
//line /usr/local/go/src/net/http/response.go:292
		return !r1.Close
//line /usr/local/go/src/net/http/response.go:292
		// _ = "end of CoverTab[42043]"
//line /usr/local/go/src/net/http/response.go:292
	}() && func() bool {
//line /usr/local/go/src/net/http/response.go:292
		_go_fuzz_dep_.CoverTab[42044]++
//line /usr/local/go/src/net/http/response.go:292
		return r1.ProtoAtLeast(1, 1)
//line /usr/local/go/src/net/http/response.go:292
		// _ = "end of CoverTab[42044]"
//line /usr/local/go/src/net/http/response.go:292
	}() && func() bool {
//line /usr/local/go/src/net/http/response.go:292
		_go_fuzz_dep_.CoverTab[42045]++
//line /usr/local/go/src/net/http/response.go:292
		return !chunked(r1.TransferEncoding)
//line /usr/local/go/src/net/http/response.go:292
		// _ = "end of CoverTab[42045]"
//line /usr/local/go/src/net/http/response.go:292
	}() && func() bool {
//line /usr/local/go/src/net/http/response.go:292
		_go_fuzz_dep_.CoverTab[42046]++
//line /usr/local/go/src/net/http/response.go:292
		return !r1.Uncompressed
//line /usr/local/go/src/net/http/response.go:292
		// _ = "end of CoverTab[42046]"
//line /usr/local/go/src/net/http/response.go:292
	}() {
//line /usr/local/go/src/net/http/response.go:292
		_go_fuzz_dep_.CoverTab[42047]++
								r1.Close = true
//line /usr/local/go/src/net/http/response.go:293
		// _ = "end of CoverTab[42047]"
	} else {
//line /usr/local/go/src/net/http/response.go:294
		_go_fuzz_dep_.CoverTab[42048]++
//line /usr/local/go/src/net/http/response.go:294
		// _ = "end of CoverTab[42048]"
//line /usr/local/go/src/net/http/response.go:294
	}
//line /usr/local/go/src/net/http/response.go:294
	// _ = "end of CoverTab[42020]"
//line /usr/local/go/src/net/http/response.go:294
	_go_fuzz_dep_.CoverTab[42021]++

//line /usr/local/go/src/net/http/response.go:297
	tw, err := newTransferWriter(r1)
	if err != nil {
//line /usr/local/go/src/net/http/response.go:298
		_go_fuzz_dep_.CoverTab[42049]++
								return err
//line /usr/local/go/src/net/http/response.go:299
		// _ = "end of CoverTab[42049]"
	} else {
//line /usr/local/go/src/net/http/response.go:300
		_go_fuzz_dep_.CoverTab[42050]++
//line /usr/local/go/src/net/http/response.go:300
		// _ = "end of CoverTab[42050]"
//line /usr/local/go/src/net/http/response.go:300
	}
//line /usr/local/go/src/net/http/response.go:300
	// _ = "end of CoverTab[42021]"
//line /usr/local/go/src/net/http/response.go:300
	_go_fuzz_dep_.CoverTab[42022]++
							err = tw.writeHeader(w, nil)
							if err != nil {
//line /usr/local/go/src/net/http/response.go:302
		_go_fuzz_dep_.CoverTab[42051]++
								return err
//line /usr/local/go/src/net/http/response.go:303
		// _ = "end of CoverTab[42051]"
	} else {
//line /usr/local/go/src/net/http/response.go:304
		_go_fuzz_dep_.CoverTab[42052]++
//line /usr/local/go/src/net/http/response.go:304
		// _ = "end of CoverTab[42052]"
//line /usr/local/go/src/net/http/response.go:304
	}
//line /usr/local/go/src/net/http/response.go:304
	// _ = "end of CoverTab[42022]"
//line /usr/local/go/src/net/http/response.go:304
	_go_fuzz_dep_.CoverTab[42023]++

//line /usr/local/go/src/net/http/response.go:307
	err = r.Header.WriteSubset(w, respExcludeHeader)
	if err != nil {
//line /usr/local/go/src/net/http/response.go:308
		_go_fuzz_dep_.CoverTab[42053]++
								return err
//line /usr/local/go/src/net/http/response.go:309
		// _ = "end of CoverTab[42053]"
	} else {
//line /usr/local/go/src/net/http/response.go:310
		_go_fuzz_dep_.CoverTab[42054]++
//line /usr/local/go/src/net/http/response.go:310
		// _ = "end of CoverTab[42054]"
//line /usr/local/go/src/net/http/response.go:310
	}
//line /usr/local/go/src/net/http/response.go:310
	// _ = "end of CoverTab[42023]"
//line /usr/local/go/src/net/http/response.go:310
	_go_fuzz_dep_.CoverTab[42024]++

//line /usr/local/go/src/net/http/response.go:314
	contentLengthAlreadySent := tw.shouldSendContentLength()
	if r1.ContentLength == 0 && func() bool {
//line /usr/local/go/src/net/http/response.go:315
		_go_fuzz_dep_.CoverTab[42055]++
//line /usr/local/go/src/net/http/response.go:315
		return !chunked(r1.TransferEncoding)
//line /usr/local/go/src/net/http/response.go:315
		// _ = "end of CoverTab[42055]"
//line /usr/local/go/src/net/http/response.go:315
	}() && func() bool {
//line /usr/local/go/src/net/http/response.go:315
		_go_fuzz_dep_.CoverTab[42056]++
//line /usr/local/go/src/net/http/response.go:315
		return !contentLengthAlreadySent
//line /usr/local/go/src/net/http/response.go:315
		// _ = "end of CoverTab[42056]"
//line /usr/local/go/src/net/http/response.go:315
	}() && func() bool {
//line /usr/local/go/src/net/http/response.go:315
		_go_fuzz_dep_.CoverTab[42057]++
//line /usr/local/go/src/net/http/response.go:315
		return bodyAllowedForStatus(r.StatusCode)
//line /usr/local/go/src/net/http/response.go:315
		// _ = "end of CoverTab[42057]"
//line /usr/local/go/src/net/http/response.go:315
	}() {
//line /usr/local/go/src/net/http/response.go:315
		_go_fuzz_dep_.CoverTab[42058]++
								if _, err := io.WriteString(w, "Content-Length: 0\r\n"); err != nil {
//line /usr/local/go/src/net/http/response.go:316
			_go_fuzz_dep_.CoverTab[42059]++
									return err
//line /usr/local/go/src/net/http/response.go:317
			// _ = "end of CoverTab[42059]"
		} else {
//line /usr/local/go/src/net/http/response.go:318
			_go_fuzz_dep_.CoverTab[42060]++
//line /usr/local/go/src/net/http/response.go:318
			// _ = "end of CoverTab[42060]"
//line /usr/local/go/src/net/http/response.go:318
		}
//line /usr/local/go/src/net/http/response.go:318
		// _ = "end of CoverTab[42058]"
	} else {
//line /usr/local/go/src/net/http/response.go:319
		_go_fuzz_dep_.CoverTab[42061]++
//line /usr/local/go/src/net/http/response.go:319
		// _ = "end of CoverTab[42061]"
//line /usr/local/go/src/net/http/response.go:319
	}
//line /usr/local/go/src/net/http/response.go:319
	// _ = "end of CoverTab[42024]"
//line /usr/local/go/src/net/http/response.go:319
	_go_fuzz_dep_.CoverTab[42025]++

//line /usr/local/go/src/net/http/response.go:322
	if _, err := io.WriteString(w, "\r\n"); err != nil {
//line /usr/local/go/src/net/http/response.go:322
		_go_fuzz_dep_.CoverTab[42062]++
								return err
//line /usr/local/go/src/net/http/response.go:323
		// _ = "end of CoverTab[42062]"
	} else {
//line /usr/local/go/src/net/http/response.go:324
		_go_fuzz_dep_.CoverTab[42063]++
//line /usr/local/go/src/net/http/response.go:324
		// _ = "end of CoverTab[42063]"
//line /usr/local/go/src/net/http/response.go:324
	}
//line /usr/local/go/src/net/http/response.go:324
	// _ = "end of CoverTab[42025]"
//line /usr/local/go/src/net/http/response.go:324
	_go_fuzz_dep_.CoverTab[42026]++

//line /usr/local/go/src/net/http/response.go:327
	err = tw.writeBody(w)
	if err != nil {
//line /usr/local/go/src/net/http/response.go:328
		_go_fuzz_dep_.CoverTab[42064]++
								return err
//line /usr/local/go/src/net/http/response.go:329
		// _ = "end of CoverTab[42064]"
	} else {
//line /usr/local/go/src/net/http/response.go:330
		_go_fuzz_dep_.CoverTab[42065]++
//line /usr/local/go/src/net/http/response.go:330
		// _ = "end of CoverTab[42065]"
//line /usr/local/go/src/net/http/response.go:330
	}
//line /usr/local/go/src/net/http/response.go:330
	// _ = "end of CoverTab[42026]"
//line /usr/local/go/src/net/http/response.go:330
	_go_fuzz_dep_.CoverTab[42027]++

//line /usr/local/go/src/net/http/response.go:333
	return nil
//line /usr/local/go/src/net/http/response.go:333
	// _ = "end of CoverTab[42027]"
}

func (r *Response) closeBody() {
//line /usr/local/go/src/net/http/response.go:336
	_go_fuzz_dep_.CoverTab[42066]++
							if r.Body != nil {
//line /usr/local/go/src/net/http/response.go:337
		_go_fuzz_dep_.CoverTab[42067]++
								r.Body.Close()
//line /usr/local/go/src/net/http/response.go:338
		// _ = "end of CoverTab[42067]"
	} else {
//line /usr/local/go/src/net/http/response.go:339
		_go_fuzz_dep_.CoverTab[42068]++
//line /usr/local/go/src/net/http/response.go:339
		// _ = "end of CoverTab[42068]"
//line /usr/local/go/src/net/http/response.go:339
	}
//line /usr/local/go/src/net/http/response.go:339
	// _ = "end of CoverTab[42066]"
}

// bodyIsWritable reports whether the Body supports writing. The
//line /usr/local/go/src/net/http/response.go:342
// Transport returns Writable bodies for 101 Switching Protocols
//line /usr/local/go/src/net/http/response.go:342
// responses.
//line /usr/local/go/src/net/http/response.go:342
// The Transport uses this method to determine whether a persistent
//line /usr/local/go/src/net/http/response.go:342
// connection is done being managed from its perspective. Once we
//line /usr/local/go/src/net/http/response.go:342
// return a writable response body to a user, the net/http package is
//line /usr/local/go/src/net/http/response.go:342
// done managing that connection.
//line /usr/local/go/src/net/http/response.go:349
func (r *Response) bodyIsWritable() bool {
//line /usr/local/go/src/net/http/response.go:349
	_go_fuzz_dep_.CoverTab[42069]++
							_, ok := r.Body.(io.Writer)
							return ok
//line /usr/local/go/src/net/http/response.go:351
	// _ = "end of CoverTab[42069]"
}

// isProtocolSwitch reports whether the response code and header
//line /usr/local/go/src/net/http/response.go:354
// indicate a successful protocol upgrade response.
//line /usr/local/go/src/net/http/response.go:356
func (r *Response) isProtocolSwitch() bool {
//line /usr/local/go/src/net/http/response.go:356
	_go_fuzz_dep_.CoverTab[42070]++
							return isProtocolSwitchResponse(r.StatusCode, r.Header)
//line /usr/local/go/src/net/http/response.go:357
	// _ = "end of CoverTab[42070]"
}

// isProtocolSwitchResponse reports whether the response code and
//line /usr/local/go/src/net/http/response.go:360
// response header indicate a successful protocol upgrade response.
//line /usr/local/go/src/net/http/response.go:362
func isProtocolSwitchResponse(code int, h Header) bool {
//line /usr/local/go/src/net/http/response.go:362
	_go_fuzz_dep_.CoverTab[42071]++
							return code == StatusSwitchingProtocols && func() bool {
//line /usr/local/go/src/net/http/response.go:363
		_go_fuzz_dep_.CoverTab[42072]++
//line /usr/local/go/src/net/http/response.go:363
		return isProtocolSwitchHeader(h)
//line /usr/local/go/src/net/http/response.go:363
		// _ = "end of CoverTab[42072]"
//line /usr/local/go/src/net/http/response.go:363
	}()
//line /usr/local/go/src/net/http/response.go:363
	// _ = "end of CoverTab[42071]"
}

// isProtocolSwitchHeader reports whether the request or response header
//line /usr/local/go/src/net/http/response.go:366
// is for a protocol switch.
//line /usr/local/go/src/net/http/response.go:368
func isProtocolSwitchHeader(h Header) bool {
//line /usr/local/go/src/net/http/response.go:368
	_go_fuzz_dep_.CoverTab[42073]++
							return h.Get("Upgrade") != "" && func() bool {
//line /usr/local/go/src/net/http/response.go:369
		_go_fuzz_dep_.CoverTab[42074]++
//line /usr/local/go/src/net/http/response.go:369
		return httpguts.HeaderValuesContainsToken(h["Connection"], "Upgrade")
								// _ = "end of CoverTab[42074]"
//line /usr/local/go/src/net/http/response.go:370
	}()
//line /usr/local/go/src/net/http/response.go:370
	// _ = "end of CoverTab[42073]"
}

//line /usr/local/go/src/net/http/response.go:371
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/response.go:371
var _ = _go_fuzz_dep_.CoverTab
