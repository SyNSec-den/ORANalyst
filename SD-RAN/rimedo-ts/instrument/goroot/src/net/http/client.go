// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP client. See RFC 7230 through 7235.
//
// This is the high-level Client interface.
// The low-level implementation is in transport.go.

//line /usr/local/go/src/net/http/client.go:10
package http

//line /usr/local/go/src/net/http/client.go:10
import (
//line /usr/local/go/src/net/http/client.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/client.go:10
)
//line /usr/local/go/src/net/http/client.go:10
import (
//line /usr/local/go/src/net/http/client.go:10
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/client.go:10
)

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http/internal/ascii"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// A Client is an HTTP client. Its zero value (DefaultClient) is a
//line /usr/local/go/src/net/http/client.go:30
// usable client that uses DefaultTransport.
//line /usr/local/go/src/net/http/client.go:30
//
//line /usr/local/go/src/net/http/client.go:30
// The Client's Transport typically has internal state (cached TCP
//line /usr/local/go/src/net/http/client.go:30
// connections), so Clients should be reused instead of created as
//line /usr/local/go/src/net/http/client.go:30
// needed. Clients are safe for concurrent use by multiple goroutines.
//line /usr/local/go/src/net/http/client.go:30
//
//line /usr/local/go/src/net/http/client.go:30
// A Client is higher-level than a RoundTripper (such as Transport)
//line /usr/local/go/src/net/http/client.go:30
// and additionally handles HTTP details such as cookies and
//line /usr/local/go/src/net/http/client.go:30
// redirects.
//line /usr/local/go/src/net/http/client.go:30
//
//line /usr/local/go/src/net/http/client.go:30
// When following redirects, the Client will forward all headers set on the
//line /usr/local/go/src/net/http/client.go:30
// initial Request except:
//line /usr/local/go/src/net/http/client.go:30
//
//line /usr/local/go/src/net/http/client.go:30
// • when forwarding sensitive headers like "Authorization",
//line /usr/local/go/src/net/http/client.go:30
// "WWW-Authenticate", and "Cookie" to untrusted targets.
//line /usr/local/go/src/net/http/client.go:30
// These headers will be ignored when following a redirect to a domain
//line /usr/local/go/src/net/http/client.go:30
// that is not a subdomain match or exact match of the initial domain.
//line /usr/local/go/src/net/http/client.go:30
// For example, a redirect from "foo.com" to either "foo.com" or "sub.foo.com"
//line /usr/local/go/src/net/http/client.go:30
// will forward the sensitive headers, but a redirect to "bar.com" will not.
//line /usr/local/go/src/net/http/client.go:30
//
//line /usr/local/go/src/net/http/client.go:30
// • when forwarding the "Cookie" header with a non-nil cookie Jar.
//line /usr/local/go/src/net/http/client.go:30
// Since each redirect may mutate the state of the cookie jar,
//line /usr/local/go/src/net/http/client.go:30
// a redirect may possibly alter a cookie set in the initial request.
//line /usr/local/go/src/net/http/client.go:30
// When forwarding the "Cookie" header, any mutated cookies will be omitted,
//line /usr/local/go/src/net/http/client.go:30
// with the expectation that the Jar will insert those mutated cookies
//line /usr/local/go/src/net/http/client.go:30
// with the updated values (assuming the origin matches).
//line /usr/local/go/src/net/http/client.go:30
// If Jar is nil, the initial cookies are forwarded without change.
//line /usr/local/go/src/net/http/client.go:58
type Client struct {
	// Transport specifies the mechanism by which individual
	// HTTP requests are made.
	// If nil, DefaultTransport is used.
	Transport	RoundTripper

	// CheckRedirect specifies the policy for handling redirects.
	// If CheckRedirect is not nil, the client calls it before
	// following an HTTP redirect. The arguments req and via are
	// the upcoming request and the requests made already, oldest
	// first. If CheckRedirect returns an error, the Client's Get
	// method returns both the previous Response (with its Body
	// closed) and CheckRedirect's error (wrapped in a url.Error)
	// instead of issuing the Request req.
	// As a special case, if CheckRedirect returns ErrUseLastResponse,
	// then the most recent response is returned with its body
	// unclosed, along with a nil error.
	//
	// If CheckRedirect is nil, the Client uses its default policy,
	// which is to stop after 10 consecutive requests.
	CheckRedirect	func(req *Request, via []*Request) error

	// Jar specifies the cookie jar.
	//
	// The Jar is used to insert relevant cookies into every
	// outbound Request and is updated with the cookie values
	// of every inbound Response. The Jar is consulted for every
	// redirect that the Client follows.
	//
	// If Jar is nil, cookies are only sent if they are explicitly
	// set on the Request.
	Jar	CookieJar

	// Timeout specifies a time limit for requests made by this
	// Client. The timeout includes connection time, any
	// redirects, and reading the response body. The timer remains
	// running after Get, Head, Post, or Do return and will
	// interrupt reading of the Response.Body.
	//
	// A Timeout of zero means no timeout.
	//
	// The Client cancels requests to the underlying Transport
	// as if the Request's Context ended.
	//
	// For compatibility, the Client will also use the deprecated
	// CancelRequest method on Transport if found. New
	// RoundTripper implementations should use the Request's Context
	// for cancellation instead of implementing CancelRequest.
	Timeout	time.Duration
}

// DefaultClient is the default Client and is used by Get, Head, and Post.
var DefaultClient = &Client{}

// RoundTripper is an interface representing the ability to execute a
//line /usr/local/go/src/net/http/client.go:112
// single HTTP transaction, obtaining the Response for a given Request.
//line /usr/local/go/src/net/http/client.go:112
//
//line /usr/local/go/src/net/http/client.go:112
// A RoundTripper must be safe for concurrent use by multiple
//line /usr/local/go/src/net/http/client.go:112
// goroutines.
//line /usr/local/go/src/net/http/client.go:117
type RoundTripper interface {
	// RoundTrip executes a single HTTP transaction, returning
	// a Response for the provided Request.
	//
	// RoundTrip should not attempt to interpret the response. In
	// particular, RoundTrip must return err == nil if it obtained
	// a response, regardless of the response's HTTP status code.
	// A non-nil err should be reserved for failure to obtain a
	// response. Similarly, RoundTrip should not attempt to
	// handle higher-level protocol details such as redirects,
	// authentication, or cookies.
	//
	// RoundTrip should not modify the request, except for
	// consuming and closing the Request's Body. RoundTrip may
	// read fields of the request in a separate goroutine. Callers
	// should not mutate or reuse the request until the Response's
	// Body has been closed.
	//
	// RoundTrip must always close the body, including on errors,
	// but depending on the implementation may do so in a separate
	// goroutine even after RoundTrip returns. This means that
	// callers wanting to reuse the body for subsequent requests
	// must arrange to wait for the Close call before doing so.
	//
	// The Request's URL and Header fields must be initialized.
	RoundTrip(*Request) (*Response, error)
}

// refererForURL returns a referer without any authentication info or
//line /usr/local/go/src/net/http/client.go:145
// an empty string if lastReq scheme is https and newReq scheme is http.
//line /usr/local/go/src/net/http/client.go:147
func refererForURL(lastReq, newReq *url.URL) string {
//line /usr/local/go/src/net/http/client.go:147
	_go_fuzz_dep_.CoverTab[36599]++

//line /usr/local/go/src/net/http/client.go:152
	if lastReq.Scheme == "https" && func() bool {
//line /usr/local/go/src/net/http/client.go:152
		_go_fuzz_dep_.CoverTab[36602]++
//line /usr/local/go/src/net/http/client.go:152
		return newReq.Scheme == "http"
//line /usr/local/go/src/net/http/client.go:152
		// _ = "end of CoverTab[36602]"
//line /usr/local/go/src/net/http/client.go:152
	}() {
//line /usr/local/go/src/net/http/client.go:152
		_go_fuzz_dep_.CoverTab[36603]++
								return ""
//line /usr/local/go/src/net/http/client.go:153
		// _ = "end of CoverTab[36603]"
	} else {
//line /usr/local/go/src/net/http/client.go:154
		_go_fuzz_dep_.CoverTab[36604]++
//line /usr/local/go/src/net/http/client.go:154
		// _ = "end of CoverTab[36604]"
//line /usr/local/go/src/net/http/client.go:154
	}
//line /usr/local/go/src/net/http/client.go:154
	// _ = "end of CoverTab[36599]"
//line /usr/local/go/src/net/http/client.go:154
	_go_fuzz_dep_.CoverTab[36600]++
							referer := lastReq.String()
							if lastReq.User != nil {
//line /usr/local/go/src/net/http/client.go:156
		_go_fuzz_dep_.CoverTab[36605]++

//line /usr/local/go/src/net/http/client.go:163
		auth := lastReq.User.String() + "@"
								referer = strings.Replace(referer, auth, "", 1)
//line /usr/local/go/src/net/http/client.go:164
		// _ = "end of CoverTab[36605]"
	} else {
//line /usr/local/go/src/net/http/client.go:165
		_go_fuzz_dep_.CoverTab[36606]++
//line /usr/local/go/src/net/http/client.go:165
		// _ = "end of CoverTab[36606]"
//line /usr/local/go/src/net/http/client.go:165
	}
//line /usr/local/go/src/net/http/client.go:165
	// _ = "end of CoverTab[36600]"
//line /usr/local/go/src/net/http/client.go:165
	_go_fuzz_dep_.CoverTab[36601]++
							return referer
//line /usr/local/go/src/net/http/client.go:166
	// _ = "end of CoverTab[36601]"
}

// didTimeout is non-nil only if err != nil.
func (c *Client) send(req *Request, deadline time.Time) (resp *Response, didTimeout func() bool, err error) {
//line /usr/local/go/src/net/http/client.go:170
	_go_fuzz_dep_.CoverTab[36607]++
							if c.Jar != nil {
//line /usr/local/go/src/net/http/client.go:171
		_go_fuzz_dep_.CoverTab[36611]++
								for _, cookie := range c.Jar.Cookies(req.URL) {
//line /usr/local/go/src/net/http/client.go:172
			_go_fuzz_dep_.CoverTab[36612]++
									req.AddCookie(cookie)
//line /usr/local/go/src/net/http/client.go:173
			// _ = "end of CoverTab[36612]"
		}
//line /usr/local/go/src/net/http/client.go:174
		// _ = "end of CoverTab[36611]"
	} else {
//line /usr/local/go/src/net/http/client.go:175
		_go_fuzz_dep_.CoverTab[36613]++
//line /usr/local/go/src/net/http/client.go:175
		// _ = "end of CoverTab[36613]"
//line /usr/local/go/src/net/http/client.go:175
	}
//line /usr/local/go/src/net/http/client.go:175
	// _ = "end of CoverTab[36607]"
//line /usr/local/go/src/net/http/client.go:175
	_go_fuzz_dep_.CoverTab[36608]++
							resp, didTimeout, err = send(req, c.transport(), deadline)
							if err != nil {
//line /usr/local/go/src/net/http/client.go:177
		_go_fuzz_dep_.CoverTab[36614]++
								return nil, didTimeout, err
//line /usr/local/go/src/net/http/client.go:178
		// _ = "end of CoverTab[36614]"
	} else {
//line /usr/local/go/src/net/http/client.go:179
		_go_fuzz_dep_.CoverTab[36615]++
//line /usr/local/go/src/net/http/client.go:179
		// _ = "end of CoverTab[36615]"
//line /usr/local/go/src/net/http/client.go:179
	}
//line /usr/local/go/src/net/http/client.go:179
	// _ = "end of CoverTab[36608]"
//line /usr/local/go/src/net/http/client.go:179
	_go_fuzz_dep_.CoverTab[36609]++
							if c.Jar != nil {
//line /usr/local/go/src/net/http/client.go:180
		_go_fuzz_dep_.CoverTab[36616]++
								if rc := resp.Cookies(); len(rc) > 0 {
//line /usr/local/go/src/net/http/client.go:181
			_go_fuzz_dep_.CoverTab[36617]++
									c.Jar.SetCookies(req.URL, rc)
//line /usr/local/go/src/net/http/client.go:182
			// _ = "end of CoverTab[36617]"
		} else {
//line /usr/local/go/src/net/http/client.go:183
			_go_fuzz_dep_.CoverTab[36618]++
//line /usr/local/go/src/net/http/client.go:183
			// _ = "end of CoverTab[36618]"
//line /usr/local/go/src/net/http/client.go:183
		}
//line /usr/local/go/src/net/http/client.go:183
		// _ = "end of CoverTab[36616]"
	} else {
//line /usr/local/go/src/net/http/client.go:184
		_go_fuzz_dep_.CoverTab[36619]++
//line /usr/local/go/src/net/http/client.go:184
		// _ = "end of CoverTab[36619]"
//line /usr/local/go/src/net/http/client.go:184
	}
//line /usr/local/go/src/net/http/client.go:184
	// _ = "end of CoverTab[36609]"
//line /usr/local/go/src/net/http/client.go:184
	_go_fuzz_dep_.CoverTab[36610]++
							return resp, nil, nil
//line /usr/local/go/src/net/http/client.go:185
	// _ = "end of CoverTab[36610]"
}

func (c *Client) deadline() time.Time {
//line /usr/local/go/src/net/http/client.go:188
	_go_fuzz_dep_.CoverTab[36620]++
							if c.Timeout > 0 {
//line /usr/local/go/src/net/http/client.go:189
		_go_fuzz_dep_.CoverTab[36622]++
								return time.Now().Add(c.Timeout)
//line /usr/local/go/src/net/http/client.go:190
		// _ = "end of CoverTab[36622]"
	} else {
//line /usr/local/go/src/net/http/client.go:191
		_go_fuzz_dep_.CoverTab[36623]++
//line /usr/local/go/src/net/http/client.go:191
		// _ = "end of CoverTab[36623]"
//line /usr/local/go/src/net/http/client.go:191
	}
//line /usr/local/go/src/net/http/client.go:191
	// _ = "end of CoverTab[36620]"
//line /usr/local/go/src/net/http/client.go:191
	_go_fuzz_dep_.CoverTab[36621]++
							return time.Time{}
//line /usr/local/go/src/net/http/client.go:192
	// _ = "end of CoverTab[36621]"
}

func (c *Client) transport() RoundTripper {
//line /usr/local/go/src/net/http/client.go:195
	_go_fuzz_dep_.CoverTab[36624]++
							if c.Transport != nil {
//line /usr/local/go/src/net/http/client.go:196
		_go_fuzz_dep_.CoverTab[36626]++
								return c.Transport
//line /usr/local/go/src/net/http/client.go:197
		// _ = "end of CoverTab[36626]"
	} else {
//line /usr/local/go/src/net/http/client.go:198
		_go_fuzz_dep_.CoverTab[36627]++
//line /usr/local/go/src/net/http/client.go:198
		// _ = "end of CoverTab[36627]"
//line /usr/local/go/src/net/http/client.go:198
	}
//line /usr/local/go/src/net/http/client.go:198
	// _ = "end of CoverTab[36624]"
//line /usr/local/go/src/net/http/client.go:198
	_go_fuzz_dep_.CoverTab[36625]++
							return DefaultTransport
//line /usr/local/go/src/net/http/client.go:199
	// _ = "end of CoverTab[36625]"
}

// send issues an HTTP request.
//line /usr/local/go/src/net/http/client.go:202
// Caller should close resp.Body when done reading from it.
//line /usr/local/go/src/net/http/client.go:204
func send(ireq *Request, rt RoundTripper, deadline time.Time) (resp *Response, didTimeout func() bool, err error) {
//line /usr/local/go/src/net/http/client.go:204
	_go_fuzz_dep_.CoverTab[36628]++
							req := ireq

							if rt == nil {
//line /usr/local/go/src/net/http/client.go:207
		_go_fuzz_dep_.CoverTab[36640]++
								req.closeBody()
								return nil, alwaysFalse, errors.New("http: no Client.Transport or DefaultTransport")
//line /usr/local/go/src/net/http/client.go:209
		// _ = "end of CoverTab[36640]"
	} else {
//line /usr/local/go/src/net/http/client.go:210
		_go_fuzz_dep_.CoverTab[36641]++
//line /usr/local/go/src/net/http/client.go:210
		// _ = "end of CoverTab[36641]"
//line /usr/local/go/src/net/http/client.go:210
	}
//line /usr/local/go/src/net/http/client.go:210
	// _ = "end of CoverTab[36628]"
//line /usr/local/go/src/net/http/client.go:210
	_go_fuzz_dep_.CoverTab[36629]++

							if req.URL == nil {
//line /usr/local/go/src/net/http/client.go:212
		_go_fuzz_dep_.CoverTab[36642]++
								req.closeBody()
								return nil, alwaysFalse, errors.New("http: nil Request.URL")
//line /usr/local/go/src/net/http/client.go:214
		// _ = "end of CoverTab[36642]"
	} else {
//line /usr/local/go/src/net/http/client.go:215
		_go_fuzz_dep_.CoverTab[36643]++
//line /usr/local/go/src/net/http/client.go:215
		// _ = "end of CoverTab[36643]"
//line /usr/local/go/src/net/http/client.go:215
	}
//line /usr/local/go/src/net/http/client.go:215
	// _ = "end of CoverTab[36629]"
//line /usr/local/go/src/net/http/client.go:215
	_go_fuzz_dep_.CoverTab[36630]++

							if req.RequestURI != "" {
//line /usr/local/go/src/net/http/client.go:217
		_go_fuzz_dep_.CoverTab[36644]++
								req.closeBody()
								return nil, alwaysFalse, errors.New("http: Request.RequestURI can't be set in client requests")
//line /usr/local/go/src/net/http/client.go:219
		// _ = "end of CoverTab[36644]"
	} else {
//line /usr/local/go/src/net/http/client.go:220
		_go_fuzz_dep_.CoverTab[36645]++
//line /usr/local/go/src/net/http/client.go:220
		// _ = "end of CoverTab[36645]"
//line /usr/local/go/src/net/http/client.go:220
	}
//line /usr/local/go/src/net/http/client.go:220
	// _ = "end of CoverTab[36630]"
//line /usr/local/go/src/net/http/client.go:220
	_go_fuzz_dep_.CoverTab[36631]++

//line /usr/local/go/src/net/http/client.go:224
	forkReq := func() {
//line /usr/local/go/src/net/http/client.go:224
		_go_fuzz_dep_.CoverTab[36646]++
								if ireq == req {
//line /usr/local/go/src/net/http/client.go:225
			_go_fuzz_dep_.CoverTab[36647]++
									req = new(Request)
									*req = *ireq
//line /usr/local/go/src/net/http/client.go:227
			// _ = "end of CoverTab[36647]"
		} else {
//line /usr/local/go/src/net/http/client.go:228
			_go_fuzz_dep_.CoverTab[36648]++
//line /usr/local/go/src/net/http/client.go:228
			// _ = "end of CoverTab[36648]"
//line /usr/local/go/src/net/http/client.go:228
		}
//line /usr/local/go/src/net/http/client.go:228
		// _ = "end of CoverTab[36646]"
	}
//line /usr/local/go/src/net/http/client.go:229
	// _ = "end of CoverTab[36631]"
//line /usr/local/go/src/net/http/client.go:229
	_go_fuzz_dep_.CoverTab[36632]++

//line /usr/local/go/src/net/http/client.go:234
	if req.Header == nil {
//line /usr/local/go/src/net/http/client.go:234
		_go_fuzz_dep_.CoverTab[36649]++
								forkReq()
								req.Header = make(Header)
//line /usr/local/go/src/net/http/client.go:236
		// _ = "end of CoverTab[36649]"
	} else {
//line /usr/local/go/src/net/http/client.go:237
		_go_fuzz_dep_.CoverTab[36650]++
//line /usr/local/go/src/net/http/client.go:237
		// _ = "end of CoverTab[36650]"
//line /usr/local/go/src/net/http/client.go:237
	}
//line /usr/local/go/src/net/http/client.go:237
	// _ = "end of CoverTab[36632]"
//line /usr/local/go/src/net/http/client.go:237
	_go_fuzz_dep_.CoverTab[36633]++

							if u := req.URL.User; u != nil && func() bool {
//line /usr/local/go/src/net/http/client.go:239
		_go_fuzz_dep_.CoverTab[36651]++
//line /usr/local/go/src/net/http/client.go:239
		return req.Header.Get("Authorization") == ""
//line /usr/local/go/src/net/http/client.go:239
		// _ = "end of CoverTab[36651]"
//line /usr/local/go/src/net/http/client.go:239
	}() {
//line /usr/local/go/src/net/http/client.go:239
		_go_fuzz_dep_.CoverTab[36652]++
								username := u.Username()
								password, _ := u.Password()
								forkReq()
								req.Header = cloneOrMakeHeader(ireq.Header)
								req.Header.Set("Authorization", "Basic "+basicAuth(username, password))
//line /usr/local/go/src/net/http/client.go:244
		// _ = "end of CoverTab[36652]"
	} else {
//line /usr/local/go/src/net/http/client.go:245
		_go_fuzz_dep_.CoverTab[36653]++
//line /usr/local/go/src/net/http/client.go:245
		// _ = "end of CoverTab[36653]"
//line /usr/local/go/src/net/http/client.go:245
	}
//line /usr/local/go/src/net/http/client.go:245
	// _ = "end of CoverTab[36633]"
//line /usr/local/go/src/net/http/client.go:245
	_go_fuzz_dep_.CoverTab[36634]++

							if !deadline.IsZero() {
//line /usr/local/go/src/net/http/client.go:247
		_go_fuzz_dep_.CoverTab[36654]++
								forkReq()
//line /usr/local/go/src/net/http/client.go:248
		// _ = "end of CoverTab[36654]"
	} else {
//line /usr/local/go/src/net/http/client.go:249
		_go_fuzz_dep_.CoverTab[36655]++
//line /usr/local/go/src/net/http/client.go:249
		// _ = "end of CoverTab[36655]"
//line /usr/local/go/src/net/http/client.go:249
	}
//line /usr/local/go/src/net/http/client.go:249
	// _ = "end of CoverTab[36634]"
//line /usr/local/go/src/net/http/client.go:249
	_go_fuzz_dep_.CoverTab[36635]++
							stopTimer, didTimeout := setRequestCancel(req, rt, deadline)

							resp, err = rt.RoundTrip(req)
							if err != nil {
//line /usr/local/go/src/net/http/client.go:253
		_go_fuzz_dep_.CoverTab[36656]++
								stopTimer()
								if resp != nil {
//line /usr/local/go/src/net/http/client.go:255
			_go_fuzz_dep_.CoverTab[36659]++
									log.Printf("RoundTripper returned a response & error; ignoring response")
//line /usr/local/go/src/net/http/client.go:256
			// _ = "end of CoverTab[36659]"
		} else {
//line /usr/local/go/src/net/http/client.go:257
			_go_fuzz_dep_.CoverTab[36660]++
//line /usr/local/go/src/net/http/client.go:257
			// _ = "end of CoverTab[36660]"
//line /usr/local/go/src/net/http/client.go:257
		}
//line /usr/local/go/src/net/http/client.go:257
		// _ = "end of CoverTab[36656]"
//line /usr/local/go/src/net/http/client.go:257
		_go_fuzz_dep_.CoverTab[36657]++
								if tlsErr, ok := err.(tls.RecordHeaderError); ok {
//line /usr/local/go/src/net/http/client.go:258
			_go_fuzz_dep_.CoverTab[36661]++

//line /usr/local/go/src/net/http/client.go:262
			if string(tlsErr.RecordHeader[:]) == "HTTP/" {
//line /usr/local/go/src/net/http/client.go:262
				_go_fuzz_dep_.CoverTab[36662]++
										err = errors.New("http: server gave HTTP response to HTTPS client")
//line /usr/local/go/src/net/http/client.go:263
				// _ = "end of CoverTab[36662]"
			} else {
//line /usr/local/go/src/net/http/client.go:264
				_go_fuzz_dep_.CoverTab[36663]++
//line /usr/local/go/src/net/http/client.go:264
				// _ = "end of CoverTab[36663]"
//line /usr/local/go/src/net/http/client.go:264
			}
//line /usr/local/go/src/net/http/client.go:264
			// _ = "end of CoverTab[36661]"
		} else {
//line /usr/local/go/src/net/http/client.go:265
			_go_fuzz_dep_.CoverTab[36664]++
//line /usr/local/go/src/net/http/client.go:265
			// _ = "end of CoverTab[36664]"
//line /usr/local/go/src/net/http/client.go:265
		}
//line /usr/local/go/src/net/http/client.go:265
		// _ = "end of CoverTab[36657]"
//line /usr/local/go/src/net/http/client.go:265
		_go_fuzz_dep_.CoverTab[36658]++
								return nil, didTimeout, err
//line /usr/local/go/src/net/http/client.go:266
		// _ = "end of CoverTab[36658]"
	} else {
//line /usr/local/go/src/net/http/client.go:267
		_go_fuzz_dep_.CoverTab[36665]++
//line /usr/local/go/src/net/http/client.go:267
		// _ = "end of CoverTab[36665]"
//line /usr/local/go/src/net/http/client.go:267
	}
//line /usr/local/go/src/net/http/client.go:267
	// _ = "end of CoverTab[36635]"
//line /usr/local/go/src/net/http/client.go:267
	_go_fuzz_dep_.CoverTab[36636]++
							if resp == nil {
//line /usr/local/go/src/net/http/client.go:268
		_go_fuzz_dep_.CoverTab[36666]++
								return nil, didTimeout, fmt.Errorf("http: RoundTripper implementation (%T) returned a nil *Response with a nil error", rt)
//line /usr/local/go/src/net/http/client.go:269
		// _ = "end of CoverTab[36666]"
	} else {
//line /usr/local/go/src/net/http/client.go:270
		_go_fuzz_dep_.CoverTab[36667]++
//line /usr/local/go/src/net/http/client.go:270
		// _ = "end of CoverTab[36667]"
//line /usr/local/go/src/net/http/client.go:270
	}
//line /usr/local/go/src/net/http/client.go:270
	// _ = "end of CoverTab[36636]"
//line /usr/local/go/src/net/http/client.go:270
	_go_fuzz_dep_.CoverTab[36637]++
							if resp.Body == nil {
//line /usr/local/go/src/net/http/client.go:271
		_go_fuzz_dep_.CoverTab[36668]++

//line /usr/local/go/src/net/http/client.go:282
		if resp.ContentLength > 0 && func() bool {
//line /usr/local/go/src/net/http/client.go:282
			_go_fuzz_dep_.CoverTab[36670]++
//line /usr/local/go/src/net/http/client.go:282
			return req.Method != "HEAD"
//line /usr/local/go/src/net/http/client.go:282
			// _ = "end of CoverTab[36670]"
//line /usr/local/go/src/net/http/client.go:282
		}() {
//line /usr/local/go/src/net/http/client.go:282
			_go_fuzz_dep_.CoverTab[36671]++
									return nil, didTimeout, fmt.Errorf("http: RoundTripper implementation (%T) returned a *Response with content length %d but a nil Body", rt, resp.ContentLength)
//line /usr/local/go/src/net/http/client.go:283
			// _ = "end of CoverTab[36671]"
		} else {
//line /usr/local/go/src/net/http/client.go:284
			_go_fuzz_dep_.CoverTab[36672]++
//line /usr/local/go/src/net/http/client.go:284
			// _ = "end of CoverTab[36672]"
//line /usr/local/go/src/net/http/client.go:284
		}
//line /usr/local/go/src/net/http/client.go:284
		// _ = "end of CoverTab[36668]"
//line /usr/local/go/src/net/http/client.go:284
		_go_fuzz_dep_.CoverTab[36669]++
								resp.Body = io.NopCloser(strings.NewReader(""))
//line /usr/local/go/src/net/http/client.go:285
		// _ = "end of CoverTab[36669]"
	} else {
//line /usr/local/go/src/net/http/client.go:286
		_go_fuzz_dep_.CoverTab[36673]++
//line /usr/local/go/src/net/http/client.go:286
		// _ = "end of CoverTab[36673]"
//line /usr/local/go/src/net/http/client.go:286
	}
//line /usr/local/go/src/net/http/client.go:286
	// _ = "end of CoverTab[36637]"
//line /usr/local/go/src/net/http/client.go:286
	_go_fuzz_dep_.CoverTab[36638]++
							if !deadline.IsZero() {
//line /usr/local/go/src/net/http/client.go:287
		_go_fuzz_dep_.CoverTab[36674]++
								resp.Body = &cancelTimerBody{
			stop:		stopTimer,
			rc:		resp.Body,
			reqDidTimeout:	didTimeout,
		}
//line /usr/local/go/src/net/http/client.go:292
		// _ = "end of CoverTab[36674]"
	} else {
//line /usr/local/go/src/net/http/client.go:293
		_go_fuzz_dep_.CoverTab[36675]++
//line /usr/local/go/src/net/http/client.go:293
		// _ = "end of CoverTab[36675]"
//line /usr/local/go/src/net/http/client.go:293
	}
//line /usr/local/go/src/net/http/client.go:293
	// _ = "end of CoverTab[36638]"
//line /usr/local/go/src/net/http/client.go:293
	_go_fuzz_dep_.CoverTab[36639]++
							return resp, nil, nil
//line /usr/local/go/src/net/http/client.go:294
	// _ = "end of CoverTab[36639]"
}

// timeBeforeContextDeadline reports whether the non-zero Time t is
//line /usr/local/go/src/net/http/client.go:297
// before ctx's deadline, if any. If ctx does not have a deadline, it
//line /usr/local/go/src/net/http/client.go:297
// always reports true (the deadline is considered infinite).
//line /usr/local/go/src/net/http/client.go:300
func timeBeforeContextDeadline(t time.Time, ctx context.Context) bool {
//line /usr/local/go/src/net/http/client.go:300
	_go_fuzz_dep_.CoverTab[36676]++
							d, ok := ctx.Deadline()
							if !ok {
//line /usr/local/go/src/net/http/client.go:302
		_go_fuzz_dep_.CoverTab[36678]++
								return true
//line /usr/local/go/src/net/http/client.go:303
		// _ = "end of CoverTab[36678]"
	} else {
//line /usr/local/go/src/net/http/client.go:304
		_go_fuzz_dep_.CoverTab[36679]++
//line /usr/local/go/src/net/http/client.go:304
		// _ = "end of CoverTab[36679]"
//line /usr/local/go/src/net/http/client.go:304
	}
//line /usr/local/go/src/net/http/client.go:304
	// _ = "end of CoverTab[36676]"
//line /usr/local/go/src/net/http/client.go:304
	_go_fuzz_dep_.CoverTab[36677]++
							return t.Before(d)
//line /usr/local/go/src/net/http/client.go:305
	// _ = "end of CoverTab[36677]"
}

// knownRoundTripperImpl reports whether rt is a RoundTripper that's
//line /usr/local/go/src/net/http/client.go:308
// maintained by the Go team and known to implement the latest
//line /usr/local/go/src/net/http/client.go:308
// optional semantics (notably contexts). The Request is used
//line /usr/local/go/src/net/http/client.go:308
// to check whether this particular request is using an alternate protocol,
//line /usr/local/go/src/net/http/client.go:308
// in which case we need to check the RoundTripper for that protocol.
//line /usr/local/go/src/net/http/client.go:313
func knownRoundTripperImpl(rt RoundTripper, req *Request) bool {
//line /usr/local/go/src/net/http/client.go:313
	_go_fuzz_dep_.CoverTab[36680]++
							switch t := rt.(type) {
	case *Transport:
//line /usr/local/go/src/net/http/client.go:315
		_go_fuzz_dep_.CoverTab[36683]++
								if altRT := t.alternateRoundTripper(req); altRT != nil {
//line /usr/local/go/src/net/http/client.go:316
			_go_fuzz_dep_.CoverTab[36686]++
									return knownRoundTripperImpl(altRT, req)
//line /usr/local/go/src/net/http/client.go:317
			// _ = "end of CoverTab[36686]"
		} else {
//line /usr/local/go/src/net/http/client.go:318
			_go_fuzz_dep_.CoverTab[36687]++
//line /usr/local/go/src/net/http/client.go:318
			// _ = "end of CoverTab[36687]"
//line /usr/local/go/src/net/http/client.go:318
		}
//line /usr/local/go/src/net/http/client.go:318
		// _ = "end of CoverTab[36683]"
//line /usr/local/go/src/net/http/client.go:318
		_go_fuzz_dep_.CoverTab[36684]++
								return true
//line /usr/local/go/src/net/http/client.go:319
		// _ = "end of CoverTab[36684]"
	case *http2Transport, http2noDialH2RoundTripper:
//line /usr/local/go/src/net/http/client.go:320
		_go_fuzz_dep_.CoverTab[36685]++
								return true
//line /usr/local/go/src/net/http/client.go:321
		// _ = "end of CoverTab[36685]"
	}
//line /usr/local/go/src/net/http/client.go:322
	// _ = "end of CoverTab[36680]"
//line /usr/local/go/src/net/http/client.go:322
	_go_fuzz_dep_.CoverTab[36681]++

//line /usr/local/go/src/net/http/client.go:329
	if reflect.TypeOf(rt).String() == "*http2.Transport" {
//line /usr/local/go/src/net/http/client.go:329
		_go_fuzz_dep_.CoverTab[36688]++
								return true
//line /usr/local/go/src/net/http/client.go:330
		// _ = "end of CoverTab[36688]"
	} else {
//line /usr/local/go/src/net/http/client.go:331
		_go_fuzz_dep_.CoverTab[36689]++
//line /usr/local/go/src/net/http/client.go:331
		// _ = "end of CoverTab[36689]"
//line /usr/local/go/src/net/http/client.go:331
	}
//line /usr/local/go/src/net/http/client.go:331
	// _ = "end of CoverTab[36681]"
//line /usr/local/go/src/net/http/client.go:331
	_go_fuzz_dep_.CoverTab[36682]++
							return false
//line /usr/local/go/src/net/http/client.go:332
	// _ = "end of CoverTab[36682]"
}

// setRequestCancel sets req.Cancel and adds a deadline context to req
//line /usr/local/go/src/net/http/client.go:335
// if deadline is non-zero. The RoundTripper's type is used to
//line /usr/local/go/src/net/http/client.go:335
// determine whether the legacy CancelRequest behavior should be used.
//line /usr/local/go/src/net/http/client.go:335
//
//line /usr/local/go/src/net/http/client.go:335
// As background, there are three ways to cancel a request:
//line /usr/local/go/src/net/http/client.go:335
// First was Transport.CancelRequest. (deprecated)
//line /usr/local/go/src/net/http/client.go:335
// Second was Request.Cancel.
//line /usr/local/go/src/net/http/client.go:335
// Third was Request.Context.
//line /usr/local/go/src/net/http/client.go:335
// This function populates the second and third, and uses the first if it really needs to.
//line /usr/local/go/src/net/http/client.go:344
func setRequestCancel(req *Request, rt RoundTripper, deadline time.Time) (stopTimer func(), didTimeout func() bool) {
//line /usr/local/go/src/net/http/client.go:344
	_go_fuzz_dep_.CoverTab[36690]++
							if deadline.IsZero() {
//line /usr/local/go/src/net/http/client.go:345
		_go_fuzz_dep_.CoverTab[36697]++
								return nop, alwaysFalse
//line /usr/local/go/src/net/http/client.go:346
		// _ = "end of CoverTab[36697]"
	} else {
//line /usr/local/go/src/net/http/client.go:347
		_go_fuzz_dep_.CoverTab[36698]++
//line /usr/local/go/src/net/http/client.go:347
		// _ = "end of CoverTab[36698]"
//line /usr/local/go/src/net/http/client.go:347
	}
//line /usr/local/go/src/net/http/client.go:347
	// _ = "end of CoverTab[36690]"
//line /usr/local/go/src/net/http/client.go:347
	_go_fuzz_dep_.CoverTab[36691]++
							knownTransport := knownRoundTripperImpl(rt, req)
							oldCtx := req.Context()

							if req.Cancel == nil && func() bool {
//line /usr/local/go/src/net/http/client.go:351
		_go_fuzz_dep_.CoverTab[36699]++
//line /usr/local/go/src/net/http/client.go:351
		return knownTransport
//line /usr/local/go/src/net/http/client.go:351
		// _ = "end of CoverTab[36699]"
//line /usr/local/go/src/net/http/client.go:351
	}() {
//line /usr/local/go/src/net/http/client.go:351
		_go_fuzz_dep_.CoverTab[36700]++

//line /usr/local/go/src/net/http/client.go:354
		if !timeBeforeContextDeadline(deadline, oldCtx) {
//line /usr/local/go/src/net/http/client.go:354
			_go_fuzz_dep_.CoverTab[36702]++
									return nop, alwaysFalse
//line /usr/local/go/src/net/http/client.go:355
			// _ = "end of CoverTab[36702]"
		} else {
//line /usr/local/go/src/net/http/client.go:356
			_go_fuzz_dep_.CoverTab[36703]++
//line /usr/local/go/src/net/http/client.go:356
			// _ = "end of CoverTab[36703]"
//line /usr/local/go/src/net/http/client.go:356
		}
//line /usr/local/go/src/net/http/client.go:356
		// _ = "end of CoverTab[36700]"
//line /usr/local/go/src/net/http/client.go:356
		_go_fuzz_dep_.CoverTab[36701]++

								var cancelCtx func()
								req.ctx, cancelCtx = context.WithDeadline(oldCtx, deadline)
								return cancelCtx, func() bool {
//line /usr/local/go/src/net/http/client.go:360
			_go_fuzz_dep_.CoverTab[36704]++
//line /usr/local/go/src/net/http/client.go:360
			return time.Now().After(deadline)
//line /usr/local/go/src/net/http/client.go:360
			// _ = "end of CoverTab[36704]"
//line /usr/local/go/src/net/http/client.go:360
		}
//line /usr/local/go/src/net/http/client.go:360
		// _ = "end of CoverTab[36701]"
	} else {
//line /usr/local/go/src/net/http/client.go:361
		_go_fuzz_dep_.CoverTab[36705]++
//line /usr/local/go/src/net/http/client.go:361
		// _ = "end of CoverTab[36705]"
//line /usr/local/go/src/net/http/client.go:361
	}
//line /usr/local/go/src/net/http/client.go:361
	// _ = "end of CoverTab[36691]"
//line /usr/local/go/src/net/http/client.go:361
	_go_fuzz_dep_.CoverTab[36692]++
							initialReqCancel := req.Cancel

							var cancelCtx func()
							if timeBeforeContextDeadline(deadline, oldCtx) {
//line /usr/local/go/src/net/http/client.go:365
		_go_fuzz_dep_.CoverTab[36706]++
								req.ctx, cancelCtx = context.WithDeadline(oldCtx, deadline)
//line /usr/local/go/src/net/http/client.go:366
		// _ = "end of CoverTab[36706]"
	} else {
//line /usr/local/go/src/net/http/client.go:367
		_go_fuzz_dep_.CoverTab[36707]++
//line /usr/local/go/src/net/http/client.go:367
		// _ = "end of CoverTab[36707]"
//line /usr/local/go/src/net/http/client.go:367
	}
//line /usr/local/go/src/net/http/client.go:367
	// _ = "end of CoverTab[36692]"
//line /usr/local/go/src/net/http/client.go:367
	_go_fuzz_dep_.CoverTab[36693]++

							cancel := make(chan struct{})
							req.Cancel = cancel

							doCancel := func() {
//line /usr/local/go/src/net/http/client.go:372
		_go_fuzz_dep_.CoverTab[36708]++

								close(cancel)
		// The first way, used only for RoundTripper
		// implementations written before Go 1.5 or Go 1.6.
		type canceler interface{ CancelRequest(*Request) }
		if v, ok := rt.(canceler); ok {
//line /usr/local/go/src/net/http/client.go:378
			_go_fuzz_dep_.CoverTab[36709]++
									v.CancelRequest(req)
//line /usr/local/go/src/net/http/client.go:379
			// _ = "end of CoverTab[36709]"
		} else {
//line /usr/local/go/src/net/http/client.go:380
			_go_fuzz_dep_.CoverTab[36710]++
//line /usr/local/go/src/net/http/client.go:380
			// _ = "end of CoverTab[36710]"
//line /usr/local/go/src/net/http/client.go:380
		}
//line /usr/local/go/src/net/http/client.go:380
		// _ = "end of CoverTab[36708]"
	}
//line /usr/local/go/src/net/http/client.go:381
	// _ = "end of CoverTab[36693]"
//line /usr/local/go/src/net/http/client.go:381
	_go_fuzz_dep_.CoverTab[36694]++

							stopTimerCh := make(chan struct{})
							var once sync.Once
							stopTimer = func() {
//line /usr/local/go/src/net/http/client.go:385
		_go_fuzz_dep_.CoverTab[36711]++
								once.Do(func() {
//line /usr/local/go/src/net/http/client.go:386
			_go_fuzz_dep_.CoverTab[36712]++
									close(stopTimerCh)
									if cancelCtx != nil {
//line /usr/local/go/src/net/http/client.go:388
				_go_fuzz_dep_.CoverTab[36713]++
										cancelCtx()
//line /usr/local/go/src/net/http/client.go:389
				// _ = "end of CoverTab[36713]"
			} else {
//line /usr/local/go/src/net/http/client.go:390
				_go_fuzz_dep_.CoverTab[36714]++
//line /usr/local/go/src/net/http/client.go:390
				// _ = "end of CoverTab[36714]"
//line /usr/local/go/src/net/http/client.go:390
			}
//line /usr/local/go/src/net/http/client.go:390
			// _ = "end of CoverTab[36712]"
		})
//line /usr/local/go/src/net/http/client.go:391
		// _ = "end of CoverTab[36711]"
	}
//line /usr/local/go/src/net/http/client.go:392
	// _ = "end of CoverTab[36694]"
//line /usr/local/go/src/net/http/client.go:392
	_go_fuzz_dep_.CoverTab[36695]++

							timer := time.NewTimer(time.Until(deadline))
							var timedOut atomic.Bool
//line /usr/local/go/src/net/http/client.go:395
	_curRoutineNum13_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/client.go:395
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum13_)

							go func() {
//line /usr/local/go/src/net/http/client.go:397
		_go_fuzz_dep_.CoverTab[36715]++
//line /usr/local/go/src/net/http/client.go:397
		defer func() {
//line /usr/local/go/src/net/http/client.go:397
			_go_fuzz_dep_.CoverTab[36716]++
//line /usr/local/go/src/net/http/client.go:397
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum13_)
//line /usr/local/go/src/net/http/client.go:397
			// _ = "end of CoverTab[36716]"
//line /usr/local/go/src/net/http/client.go:397
		}()
								select {
		case <-initialReqCancel:
//line /usr/local/go/src/net/http/client.go:399
			_go_fuzz_dep_.CoverTab[36717]++
									doCancel()
									timer.Stop()
//line /usr/local/go/src/net/http/client.go:401
			// _ = "end of CoverTab[36717]"
		case <-timer.C:
//line /usr/local/go/src/net/http/client.go:402
			_go_fuzz_dep_.CoverTab[36718]++
									timedOut.Store(true)
									doCancel()
//line /usr/local/go/src/net/http/client.go:404
			// _ = "end of CoverTab[36718]"
		case <-stopTimerCh:
//line /usr/local/go/src/net/http/client.go:405
			_go_fuzz_dep_.CoverTab[36719]++
									timer.Stop()
//line /usr/local/go/src/net/http/client.go:406
			// _ = "end of CoverTab[36719]"
		}
//line /usr/local/go/src/net/http/client.go:407
		// _ = "end of CoverTab[36715]"
	}()
//line /usr/local/go/src/net/http/client.go:408
	// _ = "end of CoverTab[36695]"
//line /usr/local/go/src/net/http/client.go:408
	_go_fuzz_dep_.CoverTab[36696]++

							return stopTimer, timedOut.Load
//line /usr/local/go/src/net/http/client.go:410
	// _ = "end of CoverTab[36696]"
}

// See 2 (end of page 4) https://www.ietf.org/rfc/rfc2617.txt
//line /usr/local/go/src/net/http/client.go:413
// "To receive authorization, the client sends the userid and password,
//line /usr/local/go/src/net/http/client.go:413
// separated by a single colon (":") character, within a base64
//line /usr/local/go/src/net/http/client.go:413
// encoded string in the credentials."
//line /usr/local/go/src/net/http/client.go:413
// It is not meant to be urlencoded.
//line /usr/local/go/src/net/http/client.go:418
func basicAuth(username, password string) string {
//line /usr/local/go/src/net/http/client.go:418
	_go_fuzz_dep_.CoverTab[36720]++
							auth := username + ":" + password
							return base64.StdEncoding.EncodeToString([]byte(auth))
//line /usr/local/go/src/net/http/client.go:420
	// _ = "end of CoverTab[36720]"
}

// Get issues a GET to the specified URL. If the response is one of
//line /usr/local/go/src/net/http/client.go:423
// the following redirect codes, Get follows the redirect, up to a
//line /usr/local/go/src/net/http/client.go:423
// maximum of 10 redirects:
//line /usr/local/go/src/net/http/client.go:423
//
//line /usr/local/go/src/net/http/client.go:423
//	301 (Moved Permanently)
//line /usr/local/go/src/net/http/client.go:423
//	302 (Found)
//line /usr/local/go/src/net/http/client.go:423
//	303 (See Other)
//line /usr/local/go/src/net/http/client.go:423
//	307 (Temporary Redirect)
//line /usr/local/go/src/net/http/client.go:423
//	308 (Permanent Redirect)
//line /usr/local/go/src/net/http/client.go:423
//
//line /usr/local/go/src/net/http/client.go:423
// An error is returned if there were too many redirects or if there
//line /usr/local/go/src/net/http/client.go:423
// was an HTTP protocol error. A non-2xx response doesn't cause an
//line /usr/local/go/src/net/http/client.go:423
// error. Any returned error will be of type *url.Error. The url.Error
//line /usr/local/go/src/net/http/client.go:423
// value's Timeout method will report true if the request timed out.
//line /usr/local/go/src/net/http/client.go:423
//
//line /usr/local/go/src/net/http/client.go:423
// When err is nil, resp always contains a non-nil resp.Body.
//line /usr/local/go/src/net/http/client.go:423
// Caller should close resp.Body when done reading from it.
//line /usr/local/go/src/net/http/client.go:423
//
//line /usr/local/go/src/net/http/client.go:423
// Get is a wrapper around DefaultClient.Get.
//line /usr/local/go/src/net/http/client.go:423
//
//line /usr/local/go/src/net/http/client.go:423
// To make a request with custom headers, use NewRequest and
//line /usr/local/go/src/net/http/client.go:423
// DefaultClient.Do.
//line /usr/local/go/src/net/http/client.go:423
//
//line /usr/local/go/src/net/http/client.go:423
// To make a request with a specified context.Context, use NewRequestWithContext
//line /usr/local/go/src/net/http/client.go:423
// and DefaultClient.Do.
//line /usr/local/go/src/net/http/client.go:448
func Get(url string) (resp *Response, err error) {
//line /usr/local/go/src/net/http/client.go:448
	_go_fuzz_dep_.CoverTab[36721]++
							return DefaultClient.Get(url)
//line /usr/local/go/src/net/http/client.go:449
	// _ = "end of CoverTab[36721]"
}

// Get issues a GET to the specified URL. If the response is one of the
//line /usr/local/go/src/net/http/client.go:452
// following redirect codes, Get follows the redirect after calling the
//line /usr/local/go/src/net/http/client.go:452
// Client's CheckRedirect function:
//line /usr/local/go/src/net/http/client.go:452
//
//line /usr/local/go/src/net/http/client.go:452
//	301 (Moved Permanently)
//line /usr/local/go/src/net/http/client.go:452
//	302 (Found)
//line /usr/local/go/src/net/http/client.go:452
//	303 (See Other)
//line /usr/local/go/src/net/http/client.go:452
//	307 (Temporary Redirect)
//line /usr/local/go/src/net/http/client.go:452
//	308 (Permanent Redirect)
//line /usr/local/go/src/net/http/client.go:452
//
//line /usr/local/go/src/net/http/client.go:452
// An error is returned if the Client's CheckRedirect function fails
//line /usr/local/go/src/net/http/client.go:452
// or if there was an HTTP protocol error. A non-2xx response doesn't
//line /usr/local/go/src/net/http/client.go:452
// cause an error. Any returned error will be of type *url.Error. The
//line /usr/local/go/src/net/http/client.go:452
// url.Error value's Timeout method will report true if the request
//line /usr/local/go/src/net/http/client.go:452
// timed out.
//line /usr/local/go/src/net/http/client.go:452
//
//line /usr/local/go/src/net/http/client.go:452
// When err is nil, resp always contains a non-nil resp.Body.
//line /usr/local/go/src/net/http/client.go:452
// Caller should close resp.Body when done reading from it.
//line /usr/local/go/src/net/http/client.go:452
//
//line /usr/local/go/src/net/http/client.go:452
// To make a request with custom headers, use NewRequest and Client.Do.
//line /usr/local/go/src/net/http/client.go:452
//
//line /usr/local/go/src/net/http/client.go:452
// To make a request with a specified context.Context, use NewRequestWithContext
//line /usr/local/go/src/net/http/client.go:452
// and Client.Do.
//line /usr/local/go/src/net/http/client.go:475
func (c *Client) Get(url string) (resp *Response, err error) {
//line /usr/local/go/src/net/http/client.go:475
	_go_fuzz_dep_.CoverTab[36722]++
							req, err := NewRequest("GET", url, nil)
							if err != nil {
//line /usr/local/go/src/net/http/client.go:477
		_go_fuzz_dep_.CoverTab[36724]++
								return nil, err
//line /usr/local/go/src/net/http/client.go:478
		// _ = "end of CoverTab[36724]"
	} else {
//line /usr/local/go/src/net/http/client.go:479
		_go_fuzz_dep_.CoverTab[36725]++
//line /usr/local/go/src/net/http/client.go:479
		// _ = "end of CoverTab[36725]"
//line /usr/local/go/src/net/http/client.go:479
	}
//line /usr/local/go/src/net/http/client.go:479
	// _ = "end of CoverTab[36722]"
//line /usr/local/go/src/net/http/client.go:479
	_go_fuzz_dep_.CoverTab[36723]++
							return c.Do(req)
//line /usr/local/go/src/net/http/client.go:480
	// _ = "end of CoverTab[36723]"
}

func alwaysFalse() bool {
//line /usr/local/go/src/net/http/client.go:483
	_go_fuzz_dep_.CoverTab[36726]++
//line /usr/local/go/src/net/http/client.go:483
	return false
//line /usr/local/go/src/net/http/client.go:483
	// _ = "end of CoverTab[36726]"
//line /usr/local/go/src/net/http/client.go:483
}

// ErrUseLastResponse can be returned by Client.CheckRedirect hooks to
//line /usr/local/go/src/net/http/client.go:485
// control how redirects are processed. If returned, the next request
//line /usr/local/go/src/net/http/client.go:485
// is not sent and the most recent response is returned with its body
//line /usr/local/go/src/net/http/client.go:485
// unclosed.
//line /usr/local/go/src/net/http/client.go:489
var ErrUseLastResponse = errors.New("net/http: use last response")

// checkRedirect calls either the user's configured CheckRedirect
//line /usr/local/go/src/net/http/client.go:491
// function, or the default.
//line /usr/local/go/src/net/http/client.go:493
func (c *Client) checkRedirect(req *Request, via []*Request) error {
//line /usr/local/go/src/net/http/client.go:493
	_go_fuzz_dep_.CoverTab[36727]++
							fn := c.CheckRedirect
							if fn == nil {
//line /usr/local/go/src/net/http/client.go:495
		_go_fuzz_dep_.CoverTab[36729]++
								fn = defaultCheckRedirect
//line /usr/local/go/src/net/http/client.go:496
		// _ = "end of CoverTab[36729]"
	} else {
//line /usr/local/go/src/net/http/client.go:497
		_go_fuzz_dep_.CoverTab[36730]++
//line /usr/local/go/src/net/http/client.go:497
		// _ = "end of CoverTab[36730]"
//line /usr/local/go/src/net/http/client.go:497
	}
//line /usr/local/go/src/net/http/client.go:497
	// _ = "end of CoverTab[36727]"
//line /usr/local/go/src/net/http/client.go:497
	_go_fuzz_dep_.CoverTab[36728]++
							return fn(req, via)
//line /usr/local/go/src/net/http/client.go:498
	// _ = "end of CoverTab[36728]"
}

// redirectBehavior describes what should happen when the
//line /usr/local/go/src/net/http/client.go:501
// client encounters a 3xx status code from the server.
//line /usr/local/go/src/net/http/client.go:503
func redirectBehavior(reqMethod string, resp *Response, ireq *Request) (redirectMethod string, shouldRedirect, includeBody bool) {
//line /usr/local/go/src/net/http/client.go:503
	_go_fuzz_dep_.CoverTab[36731]++
							switch resp.StatusCode {
	case 301, 302, 303:
//line /usr/local/go/src/net/http/client.go:505
		_go_fuzz_dep_.CoverTab[36733]++
								redirectMethod = reqMethod
								shouldRedirect = true
								includeBody = false

//line /usr/local/go/src/net/http/client.go:514
		if reqMethod != "GET" && func() bool {
//line /usr/local/go/src/net/http/client.go:514
			_go_fuzz_dep_.CoverTab[36736]++
//line /usr/local/go/src/net/http/client.go:514
			return reqMethod != "HEAD"
//line /usr/local/go/src/net/http/client.go:514
			// _ = "end of CoverTab[36736]"
//line /usr/local/go/src/net/http/client.go:514
		}() {
//line /usr/local/go/src/net/http/client.go:514
			_go_fuzz_dep_.CoverTab[36737]++
									redirectMethod = "GET"
//line /usr/local/go/src/net/http/client.go:515
			// _ = "end of CoverTab[36737]"
		} else {
//line /usr/local/go/src/net/http/client.go:516
			_go_fuzz_dep_.CoverTab[36738]++
//line /usr/local/go/src/net/http/client.go:516
			// _ = "end of CoverTab[36738]"
//line /usr/local/go/src/net/http/client.go:516
		}
//line /usr/local/go/src/net/http/client.go:516
		// _ = "end of CoverTab[36733]"
	case 307, 308:
//line /usr/local/go/src/net/http/client.go:517
		_go_fuzz_dep_.CoverTab[36734]++
								redirectMethod = reqMethod
								shouldRedirect = true
								includeBody = true

								if ireq.GetBody == nil && func() bool {
//line /usr/local/go/src/net/http/client.go:522
			_go_fuzz_dep_.CoverTab[36739]++
//line /usr/local/go/src/net/http/client.go:522
			return ireq.outgoingLength() != 0
//line /usr/local/go/src/net/http/client.go:522
			// _ = "end of CoverTab[36739]"
//line /usr/local/go/src/net/http/client.go:522
		}() {
//line /usr/local/go/src/net/http/client.go:522
			_go_fuzz_dep_.CoverTab[36740]++

//line /usr/local/go/src/net/http/client.go:527
			shouldRedirect = false
//line /usr/local/go/src/net/http/client.go:527
			// _ = "end of CoverTab[36740]"
		} else {
//line /usr/local/go/src/net/http/client.go:528
			_go_fuzz_dep_.CoverTab[36741]++
//line /usr/local/go/src/net/http/client.go:528
			// _ = "end of CoverTab[36741]"
//line /usr/local/go/src/net/http/client.go:528
		}
//line /usr/local/go/src/net/http/client.go:528
		// _ = "end of CoverTab[36734]"
//line /usr/local/go/src/net/http/client.go:528
	default:
//line /usr/local/go/src/net/http/client.go:528
		_go_fuzz_dep_.CoverTab[36735]++
//line /usr/local/go/src/net/http/client.go:528
		// _ = "end of CoverTab[36735]"
	}
//line /usr/local/go/src/net/http/client.go:529
	// _ = "end of CoverTab[36731]"
//line /usr/local/go/src/net/http/client.go:529
	_go_fuzz_dep_.CoverTab[36732]++
							return redirectMethod, shouldRedirect, includeBody
//line /usr/local/go/src/net/http/client.go:530
	// _ = "end of CoverTab[36732]"
}

// urlErrorOp returns the (*url.Error).Op value to use for the
//line /usr/local/go/src/net/http/client.go:533
// provided (*Request).Method value.
//line /usr/local/go/src/net/http/client.go:535
func urlErrorOp(method string) string {
//line /usr/local/go/src/net/http/client.go:535
	_go_fuzz_dep_.CoverTab[36742]++
							if method == "" {
//line /usr/local/go/src/net/http/client.go:536
		_go_fuzz_dep_.CoverTab[36745]++
								return "Get"
//line /usr/local/go/src/net/http/client.go:537
		// _ = "end of CoverTab[36745]"
	} else {
//line /usr/local/go/src/net/http/client.go:538
		_go_fuzz_dep_.CoverTab[36746]++
//line /usr/local/go/src/net/http/client.go:538
		// _ = "end of CoverTab[36746]"
//line /usr/local/go/src/net/http/client.go:538
	}
//line /usr/local/go/src/net/http/client.go:538
	// _ = "end of CoverTab[36742]"
//line /usr/local/go/src/net/http/client.go:538
	_go_fuzz_dep_.CoverTab[36743]++
							if lowerMethod, ok := ascii.ToLower(method); ok {
//line /usr/local/go/src/net/http/client.go:539
		_go_fuzz_dep_.CoverTab[36747]++
								return method[:1] + lowerMethod[1:]
//line /usr/local/go/src/net/http/client.go:540
		// _ = "end of CoverTab[36747]"
	} else {
//line /usr/local/go/src/net/http/client.go:541
		_go_fuzz_dep_.CoverTab[36748]++
//line /usr/local/go/src/net/http/client.go:541
		// _ = "end of CoverTab[36748]"
//line /usr/local/go/src/net/http/client.go:541
	}
//line /usr/local/go/src/net/http/client.go:541
	// _ = "end of CoverTab[36743]"
//line /usr/local/go/src/net/http/client.go:541
	_go_fuzz_dep_.CoverTab[36744]++
							return method
//line /usr/local/go/src/net/http/client.go:542
	// _ = "end of CoverTab[36744]"
}

// Do sends an HTTP request and returns an HTTP response, following
//line /usr/local/go/src/net/http/client.go:545
// policy (such as redirects, cookies, auth) as configured on the
//line /usr/local/go/src/net/http/client.go:545
// client.
//line /usr/local/go/src/net/http/client.go:545
//
//line /usr/local/go/src/net/http/client.go:545
// An error is returned if caused by client policy (such as
//line /usr/local/go/src/net/http/client.go:545
// CheckRedirect), or failure to speak HTTP (such as a network
//line /usr/local/go/src/net/http/client.go:545
// connectivity problem). A non-2xx status code doesn't cause an
//line /usr/local/go/src/net/http/client.go:545
// error.
//line /usr/local/go/src/net/http/client.go:545
//
//line /usr/local/go/src/net/http/client.go:545
// If the returned error is nil, the Response will contain a non-nil
//line /usr/local/go/src/net/http/client.go:545
// Body which the user is expected to close. If the Body is not both
//line /usr/local/go/src/net/http/client.go:545
// read to EOF and closed, the Client's underlying RoundTripper
//line /usr/local/go/src/net/http/client.go:545
// (typically Transport) may not be able to re-use a persistent TCP
//line /usr/local/go/src/net/http/client.go:545
// connection to the server for a subsequent "keep-alive" request.
//line /usr/local/go/src/net/http/client.go:545
//
//line /usr/local/go/src/net/http/client.go:545
// The request Body, if non-nil, will be closed by the underlying
//line /usr/local/go/src/net/http/client.go:545
// Transport, even on errors.
//line /usr/local/go/src/net/http/client.go:545
//
//line /usr/local/go/src/net/http/client.go:545
// On error, any Response can be ignored. A non-nil Response with a
//line /usr/local/go/src/net/http/client.go:545
// non-nil error only occurs when CheckRedirect fails, and even then
//line /usr/local/go/src/net/http/client.go:545
// the returned Response.Body is already closed.
//line /usr/local/go/src/net/http/client.go:545
//
//line /usr/local/go/src/net/http/client.go:545
// Generally Get, Post, or PostForm will be used instead of Do.
//line /usr/local/go/src/net/http/client.go:545
//
//line /usr/local/go/src/net/http/client.go:545
// If the server replies with a redirect, the Client first uses the
//line /usr/local/go/src/net/http/client.go:545
// CheckRedirect function to determine whether the redirect should be
//line /usr/local/go/src/net/http/client.go:545
// followed. If permitted, a 301, 302, or 303 redirect causes
//line /usr/local/go/src/net/http/client.go:545
// subsequent requests to use HTTP method GET
//line /usr/local/go/src/net/http/client.go:545
// (or HEAD if the original request was HEAD), with no body.
//line /usr/local/go/src/net/http/client.go:545
// A 307 or 308 redirect preserves the original HTTP method and body,
//line /usr/local/go/src/net/http/client.go:545
// provided that the Request.GetBody function is defined.
//line /usr/local/go/src/net/http/client.go:545
// The NewRequest function automatically sets GetBody for common
//line /usr/local/go/src/net/http/client.go:545
// standard library body types.
//line /usr/local/go/src/net/http/client.go:545
//
//line /usr/local/go/src/net/http/client.go:545
// Any returned error will be of type *url.Error. The url.Error
//line /usr/local/go/src/net/http/client.go:545
// value's Timeout method will report true if the request timed out.
//line /usr/local/go/src/net/http/client.go:581
func (c *Client) Do(req *Request) (*Response, error) {
//line /usr/local/go/src/net/http/client.go:581
	_go_fuzz_dep_.CoverTab[36749]++
							return c.do(req)
//line /usr/local/go/src/net/http/client.go:582
	// _ = "end of CoverTab[36749]"
}

var testHookClientDoResult func(retres *Response, reterr error)

func (c *Client) do(req *Request) (retres *Response, reterr error) {
//line /usr/local/go/src/net/http/client.go:587
	_go_fuzz_dep_.CoverTab[36750]++
							if testHookClientDoResult != nil {
//line /usr/local/go/src/net/http/client.go:588
		_go_fuzz_dep_.CoverTab[36754]++
								defer func() {
//line /usr/local/go/src/net/http/client.go:589
			_go_fuzz_dep_.CoverTab[36755]++
//line /usr/local/go/src/net/http/client.go:589
			testHookClientDoResult(retres, reterr)
//line /usr/local/go/src/net/http/client.go:589
			// _ = "end of CoverTab[36755]"
//line /usr/local/go/src/net/http/client.go:589
		}()
//line /usr/local/go/src/net/http/client.go:589
		// _ = "end of CoverTab[36754]"
	} else {
//line /usr/local/go/src/net/http/client.go:590
		_go_fuzz_dep_.CoverTab[36756]++
//line /usr/local/go/src/net/http/client.go:590
		// _ = "end of CoverTab[36756]"
//line /usr/local/go/src/net/http/client.go:590
	}
//line /usr/local/go/src/net/http/client.go:590
	// _ = "end of CoverTab[36750]"
//line /usr/local/go/src/net/http/client.go:590
	_go_fuzz_dep_.CoverTab[36751]++
							if req.URL == nil {
//line /usr/local/go/src/net/http/client.go:591
		_go_fuzz_dep_.CoverTab[36757]++
								req.closeBody()
								return nil, &url.Error{
			Op:	urlErrorOp(req.Method),
			Err:	errors.New("http: nil Request.URL"),
		}
//line /usr/local/go/src/net/http/client.go:596
		// _ = "end of CoverTab[36757]"
	} else {
//line /usr/local/go/src/net/http/client.go:597
		_go_fuzz_dep_.CoverTab[36758]++
//line /usr/local/go/src/net/http/client.go:597
		// _ = "end of CoverTab[36758]"
//line /usr/local/go/src/net/http/client.go:597
	}
//line /usr/local/go/src/net/http/client.go:597
	// _ = "end of CoverTab[36751]"
//line /usr/local/go/src/net/http/client.go:597
	_go_fuzz_dep_.CoverTab[36752]++

							var (
		deadline	= c.deadline()
		reqs		[]*Request
		resp		*Response
		copyHeaders	= c.makeHeadersCopier(req)
		reqBodyClosed	= false	// have we closed the current req.Body?

		// Redirect behavior:
		redirectMethod	string
		includeBody	bool
	)
	uerr := func(err error) error {
//line /usr/local/go/src/net/http/client.go:610
		_go_fuzz_dep_.CoverTab[36759]++

								if !reqBodyClosed {
//line /usr/local/go/src/net/http/client.go:612
			_go_fuzz_dep_.CoverTab[36762]++
									req.closeBody()
//line /usr/local/go/src/net/http/client.go:613
			// _ = "end of CoverTab[36762]"
		} else {
//line /usr/local/go/src/net/http/client.go:614
			_go_fuzz_dep_.CoverTab[36763]++
//line /usr/local/go/src/net/http/client.go:614
			// _ = "end of CoverTab[36763]"
//line /usr/local/go/src/net/http/client.go:614
		}
//line /usr/local/go/src/net/http/client.go:614
		// _ = "end of CoverTab[36759]"
//line /usr/local/go/src/net/http/client.go:614
		_go_fuzz_dep_.CoverTab[36760]++
								var urlStr string
								if resp != nil && func() bool {
//line /usr/local/go/src/net/http/client.go:616
			_go_fuzz_dep_.CoverTab[36764]++
//line /usr/local/go/src/net/http/client.go:616
			return resp.Request != nil
//line /usr/local/go/src/net/http/client.go:616
			// _ = "end of CoverTab[36764]"
//line /usr/local/go/src/net/http/client.go:616
		}() {
//line /usr/local/go/src/net/http/client.go:616
			_go_fuzz_dep_.CoverTab[36765]++
									urlStr = stripPassword(resp.Request.URL)
//line /usr/local/go/src/net/http/client.go:617
			// _ = "end of CoverTab[36765]"
		} else {
//line /usr/local/go/src/net/http/client.go:618
			_go_fuzz_dep_.CoverTab[36766]++
									urlStr = stripPassword(req.URL)
//line /usr/local/go/src/net/http/client.go:619
			// _ = "end of CoverTab[36766]"
		}
//line /usr/local/go/src/net/http/client.go:620
		// _ = "end of CoverTab[36760]"
//line /usr/local/go/src/net/http/client.go:620
		_go_fuzz_dep_.CoverTab[36761]++
								return &url.Error{
			Op:	urlErrorOp(reqs[0].Method),
			URL:	urlStr,
			Err:	err,
		}
//line /usr/local/go/src/net/http/client.go:625
		// _ = "end of CoverTab[36761]"
	}
//line /usr/local/go/src/net/http/client.go:626
	// _ = "end of CoverTab[36752]"
//line /usr/local/go/src/net/http/client.go:626
	_go_fuzz_dep_.CoverTab[36753]++
							for {
//line /usr/local/go/src/net/http/client.go:627
		_go_fuzz_dep_.CoverTab[36767]++

//line /usr/local/go/src/net/http/client.go:630
		if len(reqs) > 0 {
//line /usr/local/go/src/net/http/client.go:630
			_go_fuzz_dep_.CoverTab[36771]++
									loc := resp.Header.Get("Location")
									if loc == "" {
//line /usr/local/go/src/net/http/client.go:632
				_go_fuzz_dep_.CoverTab[36779]++

//line /usr/local/go/src/net/http/client.go:636
				return resp, nil
//line /usr/local/go/src/net/http/client.go:636
				// _ = "end of CoverTab[36779]"
			} else {
//line /usr/local/go/src/net/http/client.go:637
				_go_fuzz_dep_.CoverTab[36780]++
//line /usr/local/go/src/net/http/client.go:637
				// _ = "end of CoverTab[36780]"
//line /usr/local/go/src/net/http/client.go:637
			}
//line /usr/local/go/src/net/http/client.go:637
			// _ = "end of CoverTab[36771]"
//line /usr/local/go/src/net/http/client.go:637
			_go_fuzz_dep_.CoverTab[36772]++
									u, err := req.URL.Parse(loc)
									if err != nil {
//line /usr/local/go/src/net/http/client.go:639
				_go_fuzz_dep_.CoverTab[36781]++
										resp.closeBody()
										return nil, uerr(fmt.Errorf("failed to parse Location header %q: %v", loc, err))
//line /usr/local/go/src/net/http/client.go:641
				// _ = "end of CoverTab[36781]"
			} else {
//line /usr/local/go/src/net/http/client.go:642
				_go_fuzz_dep_.CoverTab[36782]++
//line /usr/local/go/src/net/http/client.go:642
				// _ = "end of CoverTab[36782]"
//line /usr/local/go/src/net/http/client.go:642
			}
//line /usr/local/go/src/net/http/client.go:642
			// _ = "end of CoverTab[36772]"
//line /usr/local/go/src/net/http/client.go:642
			_go_fuzz_dep_.CoverTab[36773]++
									host := ""
									if req.Host != "" && func() bool {
//line /usr/local/go/src/net/http/client.go:644
				_go_fuzz_dep_.CoverTab[36783]++
//line /usr/local/go/src/net/http/client.go:644
				return req.Host != req.URL.Host
//line /usr/local/go/src/net/http/client.go:644
				// _ = "end of CoverTab[36783]"
//line /usr/local/go/src/net/http/client.go:644
			}() {
//line /usr/local/go/src/net/http/client.go:644
				_go_fuzz_dep_.CoverTab[36784]++

//line /usr/local/go/src/net/http/client.go:648
				if u, _ := url.Parse(loc); u != nil && func() bool {
//line /usr/local/go/src/net/http/client.go:648
					_go_fuzz_dep_.CoverTab[36785]++
//line /usr/local/go/src/net/http/client.go:648
					return !u.IsAbs()
//line /usr/local/go/src/net/http/client.go:648
					// _ = "end of CoverTab[36785]"
//line /usr/local/go/src/net/http/client.go:648
				}() {
//line /usr/local/go/src/net/http/client.go:648
					_go_fuzz_dep_.CoverTab[36786]++
											host = req.Host
//line /usr/local/go/src/net/http/client.go:649
					// _ = "end of CoverTab[36786]"
				} else {
//line /usr/local/go/src/net/http/client.go:650
					_go_fuzz_dep_.CoverTab[36787]++
//line /usr/local/go/src/net/http/client.go:650
					// _ = "end of CoverTab[36787]"
//line /usr/local/go/src/net/http/client.go:650
				}
//line /usr/local/go/src/net/http/client.go:650
				// _ = "end of CoverTab[36784]"
			} else {
//line /usr/local/go/src/net/http/client.go:651
				_go_fuzz_dep_.CoverTab[36788]++
//line /usr/local/go/src/net/http/client.go:651
				// _ = "end of CoverTab[36788]"
//line /usr/local/go/src/net/http/client.go:651
			}
//line /usr/local/go/src/net/http/client.go:651
			// _ = "end of CoverTab[36773]"
//line /usr/local/go/src/net/http/client.go:651
			_go_fuzz_dep_.CoverTab[36774]++
									ireq := reqs[0]
									req = &Request{
				Method:		redirectMethod,
				Response:	resp,
				URL:		u,
				Header:		make(Header),
				Host:		host,
				Cancel:		ireq.Cancel,
				ctx:		ireq.ctx,
			}
			if includeBody && func() bool {
//line /usr/local/go/src/net/http/client.go:662
				_go_fuzz_dep_.CoverTab[36789]++
//line /usr/local/go/src/net/http/client.go:662
				return ireq.GetBody != nil
//line /usr/local/go/src/net/http/client.go:662
				// _ = "end of CoverTab[36789]"
//line /usr/local/go/src/net/http/client.go:662
			}() {
//line /usr/local/go/src/net/http/client.go:662
				_go_fuzz_dep_.CoverTab[36790]++
										req.Body, err = ireq.GetBody()
										if err != nil {
//line /usr/local/go/src/net/http/client.go:664
					_go_fuzz_dep_.CoverTab[36792]++
											resp.closeBody()
											return nil, uerr(err)
//line /usr/local/go/src/net/http/client.go:666
					// _ = "end of CoverTab[36792]"
				} else {
//line /usr/local/go/src/net/http/client.go:667
					_go_fuzz_dep_.CoverTab[36793]++
//line /usr/local/go/src/net/http/client.go:667
					// _ = "end of CoverTab[36793]"
//line /usr/local/go/src/net/http/client.go:667
				}
//line /usr/local/go/src/net/http/client.go:667
				// _ = "end of CoverTab[36790]"
//line /usr/local/go/src/net/http/client.go:667
				_go_fuzz_dep_.CoverTab[36791]++
										req.ContentLength = ireq.ContentLength
//line /usr/local/go/src/net/http/client.go:668
				// _ = "end of CoverTab[36791]"
			} else {
//line /usr/local/go/src/net/http/client.go:669
				_go_fuzz_dep_.CoverTab[36794]++
//line /usr/local/go/src/net/http/client.go:669
				// _ = "end of CoverTab[36794]"
//line /usr/local/go/src/net/http/client.go:669
			}
//line /usr/local/go/src/net/http/client.go:669
			// _ = "end of CoverTab[36774]"
//line /usr/local/go/src/net/http/client.go:669
			_go_fuzz_dep_.CoverTab[36775]++

//line /usr/local/go/src/net/http/client.go:675
			copyHeaders(req)

//line /usr/local/go/src/net/http/client.go:679
			if ref := refererForURL(reqs[len(reqs)-1].URL, req.URL); ref != "" {
//line /usr/local/go/src/net/http/client.go:679
				_go_fuzz_dep_.CoverTab[36795]++
										req.Header.Set("Referer", ref)
//line /usr/local/go/src/net/http/client.go:680
				// _ = "end of CoverTab[36795]"
			} else {
//line /usr/local/go/src/net/http/client.go:681
				_go_fuzz_dep_.CoverTab[36796]++
//line /usr/local/go/src/net/http/client.go:681
				// _ = "end of CoverTab[36796]"
//line /usr/local/go/src/net/http/client.go:681
			}
//line /usr/local/go/src/net/http/client.go:681
			// _ = "end of CoverTab[36775]"
//line /usr/local/go/src/net/http/client.go:681
			_go_fuzz_dep_.CoverTab[36776]++
									err = c.checkRedirect(req, reqs)

//line /usr/local/go/src/net/http/client.go:687
			if err == ErrUseLastResponse {
//line /usr/local/go/src/net/http/client.go:687
				_go_fuzz_dep_.CoverTab[36797]++
										return resp, nil
//line /usr/local/go/src/net/http/client.go:688
				// _ = "end of CoverTab[36797]"
			} else {
//line /usr/local/go/src/net/http/client.go:689
				_go_fuzz_dep_.CoverTab[36798]++
//line /usr/local/go/src/net/http/client.go:689
				// _ = "end of CoverTab[36798]"
//line /usr/local/go/src/net/http/client.go:689
			}
//line /usr/local/go/src/net/http/client.go:689
			// _ = "end of CoverTab[36776]"
//line /usr/local/go/src/net/http/client.go:689
			_go_fuzz_dep_.CoverTab[36777]++

			// Close the previous response's body. But
			// read at least some of the body so if it's
			// small the underlying TCP connection will be
			// re-used. No need to check for errors: if it
			// fails, the Transport won't reuse it anyway.
			const maxBodySlurpSize = 2 << 10
			if resp.ContentLength == -1 || func() bool {
//line /usr/local/go/src/net/http/client.go:697
				_go_fuzz_dep_.CoverTab[36799]++
//line /usr/local/go/src/net/http/client.go:697
				return resp.ContentLength <= maxBodySlurpSize
//line /usr/local/go/src/net/http/client.go:697
				// _ = "end of CoverTab[36799]"
//line /usr/local/go/src/net/http/client.go:697
			}() {
//line /usr/local/go/src/net/http/client.go:697
				_go_fuzz_dep_.CoverTab[36800]++
										io.CopyN(io.Discard, resp.Body, maxBodySlurpSize)
//line /usr/local/go/src/net/http/client.go:698
				// _ = "end of CoverTab[36800]"
			} else {
//line /usr/local/go/src/net/http/client.go:699
				_go_fuzz_dep_.CoverTab[36801]++
//line /usr/local/go/src/net/http/client.go:699
				// _ = "end of CoverTab[36801]"
//line /usr/local/go/src/net/http/client.go:699
			}
//line /usr/local/go/src/net/http/client.go:699
			// _ = "end of CoverTab[36777]"
//line /usr/local/go/src/net/http/client.go:699
			_go_fuzz_dep_.CoverTab[36778]++
									resp.Body.Close()

									if err != nil {
//line /usr/local/go/src/net/http/client.go:702
				_go_fuzz_dep_.CoverTab[36802]++

//line /usr/local/go/src/net/http/client.go:707
				ue := uerr(err)
										ue.(*url.Error).URL = loc
										return resp, ue
//line /usr/local/go/src/net/http/client.go:709
				// _ = "end of CoverTab[36802]"
			} else {
//line /usr/local/go/src/net/http/client.go:710
				_go_fuzz_dep_.CoverTab[36803]++
//line /usr/local/go/src/net/http/client.go:710
				// _ = "end of CoverTab[36803]"
//line /usr/local/go/src/net/http/client.go:710
			}
//line /usr/local/go/src/net/http/client.go:710
			// _ = "end of CoverTab[36778]"
		} else {
//line /usr/local/go/src/net/http/client.go:711
			_go_fuzz_dep_.CoverTab[36804]++
//line /usr/local/go/src/net/http/client.go:711
			// _ = "end of CoverTab[36804]"
//line /usr/local/go/src/net/http/client.go:711
		}
//line /usr/local/go/src/net/http/client.go:711
		// _ = "end of CoverTab[36767]"
//line /usr/local/go/src/net/http/client.go:711
		_go_fuzz_dep_.CoverTab[36768]++

								reqs = append(reqs, req)
								var err error
								var didTimeout func() bool
								if resp, didTimeout, err = c.send(req, deadline); err != nil {
//line /usr/local/go/src/net/http/client.go:716
			_go_fuzz_dep_.CoverTab[36805]++

									reqBodyClosed = true
									if !deadline.IsZero() && func() bool {
//line /usr/local/go/src/net/http/client.go:719
				_go_fuzz_dep_.CoverTab[36807]++
//line /usr/local/go/src/net/http/client.go:719
				return didTimeout()
//line /usr/local/go/src/net/http/client.go:719
				// _ = "end of CoverTab[36807]"
//line /usr/local/go/src/net/http/client.go:719
			}() {
//line /usr/local/go/src/net/http/client.go:719
				_go_fuzz_dep_.CoverTab[36808]++
										err = &httpError{
					err:		err.Error() + " (Client.Timeout exceeded while awaiting headers)",
					timeout:	true,
				}
//line /usr/local/go/src/net/http/client.go:723
				// _ = "end of CoverTab[36808]"
			} else {
//line /usr/local/go/src/net/http/client.go:724
				_go_fuzz_dep_.CoverTab[36809]++
//line /usr/local/go/src/net/http/client.go:724
				// _ = "end of CoverTab[36809]"
//line /usr/local/go/src/net/http/client.go:724
			}
//line /usr/local/go/src/net/http/client.go:724
			// _ = "end of CoverTab[36805]"
//line /usr/local/go/src/net/http/client.go:724
			_go_fuzz_dep_.CoverTab[36806]++
									return nil, uerr(err)
//line /usr/local/go/src/net/http/client.go:725
			// _ = "end of CoverTab[36806]"
		} else {
//line /usr/local/go/src/net/http/client.go:726
			_go_fuzz_dep_.CoverTab[36810]++
//line /usr/local/go/src/net/http/client.go:726
			// _ = "end of CoverTab[36810]"
//line /usr/local/go/src/net/http/client.go:726
		}
//line /usr/local/go/src/net/http/client.go:726
		// _ = "end of CoverTab[36768]"
//line /usr/local/go/src/net/http/client.go:726
		_go_fuzz_dep_.CoverTab[36769]++

								var shouldRedirect bool
								redirectMethod, shouldRedirect, includeBody = redirectBehavior(req.Method, resp, reqs[0])
								if !shouldRedirect {
//line /usr/local/go/src/net/http/client.go:730
			_go_fuzz_dep_.CoverTab[36811]++
									return resp, nil
//line /usr/local/go/src/net/http/client.go:731
			// _ = "end of CoverTab[36811]"
		} else {
//line /usr/local/go/src/net/http/client.go:732
			_go_fuzz_dep_.CoverTab[36812]++
//line /usr/local/go/src/net/http/client.go:732
			// _ = "end of CoverTab[36812]"
//line /usr/local/go/src/net/http/client.go:732
		}
//line /usr/local/go/src/net/http/client.go:732
		// _ = "end of CoverTab[36769]"
//line /usr/local/go/src/net/http/client.go:732
		_go_fuzz_dep_.CoverTab[36770]++

								req.closeBody()
//line /usr/local/go/src/net/http/client.go:734
		// _ = "end of CoverTab[36770]"
	}
//line /usr/local/go/src/net/http/client.go:735
	// _ = "end of CoverTab[36753]"
}

// makeHeadersCopier makes a function that copies headers from the
//line /usr/local/go/src/net/http/client.go:738
// initial Request, ireq. For every redirect, this function must be called
//line /usr/local/go/src/net/http/client.go:738
// so that it can copy headers into the upcoming Request.
//line /usr/local/go/src/net/http/client.go:741
func (c *Client) makeHeadersCopier(ireq *Request) func(*Request) {
//line /usr/local/go/src/net/http/client.go:741
	_go_fuzz_dep_.CoverTab[36813]++
	// The headers to copy are from the very initial request.
	// We use a closured callback to keep a reference to these original headers.
	var (
		ireqhdr		= cloneOrMakeHeader(ireq.Header)
		icookies	map[string][]*Cookie
	)
	if c.Jar != nil && func() bool {
//line /usr/local/go/src/net/http/client.go:748
		_go_fuzz_dep_.CoverTab[36815]++
//line /usr/local/go/src/net/http/client.go:748
		return ireq.Header.Get("Cookie") != ""
//line /usr/local/go/src/net/http/client.go:748
		// _ = "end of CoverTab[36815]"
//line /usr/local/go/src/net/http/client.go:748
	}() {
//line /usr/local/go/src/net/http/client.go:748
		_go_fuzz_dep_.CoverTab[36816]++
								icookies = make(map[string][]*Cookie)
								for _, c := range ireq.Cookies() {
//line /usr/local/go/src/net/http/client.go:750
			_go_fuzz_dep_.CoverTab[36817]++
									icookies[c.Name] = append(icookies[c.Name], c)
//line /usr/local/go/src/net/http/client.go:751
			// _ = "end of CoverTab[36817]"
		}
//line /usr/local/go/src/net/http/client.go:752
		// _ = "end of CoverTab[36816]"
	} else {
//line /usr/local/go/src/net/http/client.go:753
		_go_fuzz_dep_.CoverTab[36818]++
//line /usr/local/go/src/net/http/client.go:753
		// _ = "end of CoverTab[36818]"
//line /usr/local/go/src/net/http/client.go:753
	}
//line /usr/local/go/src/net/http/client.go:753
	// _ = "end of CoverTab[36813]"
//line /usr/local/go/src/net/http/client.go:753
	_go_fuzz_dep_.CoverTab[36814]++

							preq := ireq
							return func(req *Request) {
//line /usr/local/go/src/net/http/client.go:756
		_go_fuzz_dep_.CoverTab[36819]++

//line /usr/local/go/src/net/http/client.go:768
		if c.Jar != nil && func() bool {
//line /usr/local/go/src/net/http/client.go:768
			_go_fuzz_dep_.CoverTab[36822]++
//line /usr/local/go/src/net/http/client.go:768
			return icookies != nil
//line /usr/local/go/src/net/http/client.go:768
			// _ = "end of CoverTab[36822]"
//line /usr/local/go/src/net/http/client.go:768
		}() {
//line /usr/local/go/src/net/http/client.go:768
			_go_fuzz_dep_.CoverTab[36823]++
									var changed bool
									resp := req.Response
									for _, c := range resp.Cookies() {
//line /usr/local/go/src/net/http/client.go:771
				_go_fuzz_dep_.CoverTab[36825]++
										if _, ok := icookies[c.Name]; ok {
//line /usr/local/go/src/net/http/client.go:772
					_go_fuzz_dep_.CoverTab[36826]++
											delete(icookies, c.Name)
											changed = true
//line /usr/local/go/src/net/http/client.go:774
					// _ = "end of CoverTab[36826]"
				} else {
//line /usr/local/go/src/net/http/client.go:775
					_go_fuzz_dep_.CoverTab[36827]++
//line /usr/local/go/src/net/http/client.go:775
					// _ = "end of CoverTab[36827]"
//line /usr/local/go/src/net/http/client.go:775
				}
//line /usr/local/go/src/net/http/client.go:775
				// _ = "end of CoverTab[36825]"
			}
//line /usr/local/go/src/net/http/client.go:776
			// _ = "end of CoverTab[36823]"
//line /usr/local/go/src/net/http/client.go:776
			_go_fuzz_dep_.CoverTab[36824]++
									if changed {
//line /usr/local/go/src/net/http/client.go:777
				_go_fuzz_dep_.CoverTab[36828]++
										ireqhdr.Del("Cookie")
										var ss []string
										for _, cs := range icookies {
//line /usr/local/go/src/net/http/client.go:780
					_go_fuzz_dep_.CoverTab[36830]++
											for _, c := range cs {
//line /usr/local/go/src/net/http/client.go:781
						_go_fuzz_dep_.CoverTab[36831]++
												ss = append(ss, c.Name+"="+c.Value)
//line /usr/local/go/src/net/http/client.go:782
						// _ = "end of CoverTab[36831]"
					}
//line /usr/local/go/src/net/http/client.go:783
					// _ = "end of CoverTab[36830]"
				}
//line /usr/local/go/src/net/http/client.go:784
				// _ = "end of CoverTab[36828]"
//line /usr/local/go/src/net/http/client.go:784
				_go_fuzz_dep_.CoverTab[36829]++
										sort.Strings(ss)
										ireqhdr.Set("Cookie", strings.Join(ss, "; "))
//line /usr/local/go/src/net/http/client.go:786
				// _ = "end of CoverTab[36829]"
			} else {
//line /usr/local/go/src/net/http/client.go:787
				_go_fuzz_dep_.CoverTab[36832]++
//line /usr/local/go/src/net/http/client.go:787
				// _ = "end of CoverTab[36832]"
//line /usr/local/go/src/net/http/client.go:787
			}
//line /usr/local/go/src/net/http/client.go:787
			// _ = "end of CoverTab[36824]"
		} else {
//line /usr/local/go/src/net/http/client.go:788
			_go_fuzz_dep_.CoverTab[36833]++
//line /usr/local/go/src/net/http/client.go:788
			// _ = "end of CoverTab[36833]"
//line /usr/local/go/src/net/http/client.go:788
		}
//line /usr/local/go/src/net/http/client.go:788
		// _ = "end of CoverTab[36819]"
//line /usr/local/go/src/net/http/client.go:788
		_go_fuzz_dep_.CoverTab[36820]++

//line /usr/local/go/src/net/http/client.go:792
		for k, vv := range ireqhdr {
//line /usr/local/go/src/net/http/client.go:792
			_go_fuzz_dep_.CoverTab[36834]++
									if shouldCopyHeaderOnRedirect(k, preq.URL, req.URL) {
//line /usr/local/go/src/net/http/client.go:793
				_go_fuzz_dep_.CoverTab[36835]++
										req.Header[k] = vv
//line /usr/local/go/src/net/http/client.go:794
				// _ = "end of CoverTab[36835]"
			} else {
//line /usr/local/go/src/net/http/client.go:795
				_go_fuzz_dep_.CoverTab[36836]++
//line /usr/local/go/src/net/http/client.go:795
				// _ = "end of CoverTab[36836]"
//line /usr/local/go/src/net/http/client.go:795
			}
//line /usr/local/go/src/net/http/client.go:795
			// _ = "end of CoverTab[36834]"
		}
//line /usr/local/go/src/net/http/client.go:796
		// _ = "end of CoverTab[36820]"
//line /usr/local/go/src/net/http/client.go:796
		_go_fuzz_dep_.CoverTab[36821]++

								preq = req
//line /usr/local/go/src/net/http/client.go:798
		// _ = "end of CoverTab[36821]"
	}
//line /usr/local/go/src/net/http/client.go:799
	// _ = "end of CoverTab[36814]"
}

func defaultCheckRedirect(req *Request, via []*Request) error {
//line /usr/local/go/src/net/http/client.go:802
	_go_fuzz_dep_.CoverTab[36837]++
							if len(via) >= 10 {
//line /usr/local/go/src/net/http/client.go:803
		_go_fuzz_dep_.CoverTab[36839]++
								return errors.New("stopped after 10 redirects")
//line /usr/local/go/src/net/http/client.go:804
		// _ = "end of CoverTab[36839]"
	} else {
//line /usr/local/go/src/net/http/client.go:805
		_go_fuzz_dep_.CoverTab[36840]++
//line /usr/local/go/src/net/http/client.go:805
		// _ = "end of CoverTab[36840]"
//line /usr/local/go/src/net/http/client.go:805
	}
//line /usr/local/go/src/net/http/client.go:805
	// _ = "end of CoverTab[36837]"
//line /usr/local/go/src/net/http/client.go:805
	_go_fuzz_dep_.CoverTab[36838]++
							return nil
//line /usr/local/go/src/net/http/client.go:806
	// _ = "end of CoverTab[36838]"
}

// Post issues a POST to the specified URL.
//line /usr/local/go/src/net/http/client.go:809
//
//line /usr/local/go/src/net/http/client.go:809
// Caller should close resp.Body when done reading from it.
//line /usr/local/go/src/net/http/client.go:809
//
//line /usr/local/go/src/net/http/client.go:809
// If the provided body is an io.Closer, it is closed after the
//line /usr/local/go/src/net/http/client.go:809
// request.
//line /usr/local/go/src/net/http/client.go:809
//
//line /usr/local/go/src/net/http/client.go:809
// Post is a wrapper around DefaultClient.Post.
//line /usr/local/go/src/net/http/client.go:809
//
//line /usr/local/go/src/net/http/client.go:809
// To set custom headers, use NewRequest and DefaultClient.Do.
//line /usr/local/go/src/net/http/client.go:809
//
//line /usr/local/go/src/net/http/client.go:809
// See the Client.Do method documentation for details on how redirects
//line /usr/local/go/src/net/http/client.go:809
// are handled.
//line /usr/local/go/src/net/http/client.go:809
//
//line /usr/local/go/src/net/http/client.go:809
// To make a request with a specified context.Context, use NewRequestWithContext
//line /usr/local/go/src/net/http/client.go:809
// and DefaultClient.Do.
//line /usr/local/go/src/net/http/client.go:825
func Post(url, contentType string, body io.Reader) (resp *Response, err error) {
//line /usr/local/go/src/net/http/client.go:825
	_go_fuzz_dep_.CoverTab[36841]++
							return DefaultClient.Post(url, contentType, body)
//line /usr/local/go/src/net/http/client.go:826
	// _ = "end of CoverTab[36841]"
}

// Post issues a POST to the specified URL.
//line /usr/local/go/src/net/http/client.go:829
//
//line /usr/local/go/src/net/http/client.go:829
// Caller should close resp.Body when done reading from it.
//line /usr/local/go/src/net/http/client.go:829
//
//line /usr/local/go/src/net/http/client.go:829
// If the provided body is an io.Closer, it is closed after the
//line /usr/local/go/src/net/http/client.go:829
// request.
//line /usr/local/go/src/net/http/client.go:829
//
//line /usr/local/go/src/net/http/client.go:829
// To set custom headers, use NewRequest and Client.Do.
//line /usr/local/go/src/net/http/client.go:829
//
//line /usr/local/go/src/net/http/client.go:829
// To make a request with a specified context.Context, use NewRequestWithContext
//line /usr/local/go/src/net/http/client.go:829
// and Client.Do.
//line /usr/local/go/src/net/http/client.go:829
//
//line /usr/local/go/src/net/http/client.go:829
// See the Client.Do method documentation for details on how redirects
//line /usr/local/go/src/net/http/client.go:829
// are handled.
//line /usr/local/go/src/net/http/client.go:843
func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error) {
//line /usr/local/go/src/net/http/client.go:843
	_go_fuzz_dep_.CoverTab[36842]++
							req, err := NewRequest("POST", url, body)
							if err != nil {
//line /usr/local/go/src/net/http/client.go:845
		_go_fuzz_dep_.CoverTab[36844]++
								return nil, err
//line /usr/local/go/src/net/http/client.go:846
		// _ = "end of CoverTab[36844]"
	} else {
//line /usr/local/go/src/net/http/client.go:847
		_go_fuzz_dep_.CoverTab[36845]++
//line /usr/local/go/src/net/http/client.go:847
		// _ = "end of CoverTab[36845]"
//line /usr/local/go/src/net/http/client.go:847
	}
//line /usr/local/go/src/net/http/client.go:847
	// _ = "end of CoverTab[36842]"
//line /usr/local/go/src/net/http/client.go:847
	_go_fuzz_dep_.CoverTab[36843]++
							req.Header.Set("Content-Type", contentType)
							return c.Do(req)
//line /usr/local/go/src/net/http/client.go:849
	// _ = "end of CoverTab[36843]"
}

// PostForm issues a POST to the specified URL, with data's keys and
//line /usr/local/go/src/net/http/client.go:852
// values URL-encoded as the request body.
//line /usr/local/go/src/net/http/client.go:852
//
//line /usr/local/go/src/net/http/client.go:852
// The Content-Type header is set to application/x-www-form-urlencoded.
//line /usr/local/go/src/net/http/client.go:852
// To set other headers, use NewRequest and DefaultClient.Do.
//line /usr/local/go/src/net/http/client.go:852
//
//line /usr/local/go/src/net/http/client.go:852
// When err is nil, resp always contains a non-nil resp.Body.
//line /usr/local/go/src/net/http/client.go:852
// Caller should close resp.Body when done reading from it.
//line /usr/local/go/src/net/http/client.go:852
//
//line /usr/local/go/src/net/http/client.go:852
// PostForm is a wrapper around DefaultClient.PostForm.
//line /usr/local/go/src/net/http/client.go:852
//
//line /usr/local/go/src/net/http/client.go:852
// See the Client.Do method documentation for details on how redirects
//line /usr/local/go/src/net/http/client.go:852
// are handled.
//line /usr/local/go/src/net/http/client.go:852
//
//line /usr/local/go/src/net/http/client.go:852
// To make a request with a specified context.Context, use NewRequestWithContext
//line /usr/local/go/src/net/http/client.go:852
// and DefaultClient.Do.
//line /usr/local/go/src/net/http/client.go:868
func PostForm(url string, data url.Values) (resp *Response, err error) {
//line /usr/local/go/src/net/http/client.go:868
	_go_fuzz_dep_.CoverTab[36846]++
							return DefaultClient.PostForm(url, data)
//line /usr/local/go/src/net/http/client.go:869
	// _ = "end of CoverTab[36846]"
}

// PostForm issues a POST to the specified URL,
//line /usr/local/go/src/net/http/client.go:872
// with data's keys and values URL-encoded as the request body.
//line /usr/local/go/src/net/http/client.go:872
//
//line /usr/local/go/src/net/http/client.go:872
// The Content-Type header is set to application/x-www-form-urlencoded.
//line /usr/local/go/src/net/http/client.go:872
// To set other headers, use NewRequest and Client.Do.
//line /usr/local/go/src/net/http/client.go:872
//
//line /usr/local/go/src/net/http/client.go:872
// When err is nil, resp always contains a non-nil resp.Body.
//line /usr/local/go/src/net/http/client.go:872
// Caller should close resp.Body when done reading from it.
//line /usr/local/go/src/net/http/client.go:872
//
//line /usr/local/go/src/net/http/client.go:872
// See the Client.Do method documentation for details on how redirects
//line /usr/local/go/src/net/http/client.go:872
// are handled.
//line /usr/local/go/src/net/http/client.go:872
//
//line /usr/local/go/src/net/http/client.go:872
// To make a request with a specified context.Context, use NewRequestWithContext
//line /usr/local/go/src/net/http/client.go:872
// and Client.Do.
//line /usr/local/go/src/net/http/client.go:886
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error) {
//line /usr/local/go/src/net/http/client.go:886
	_go_fuzz_dep_.CoverTab[36847]++
							return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
//line /usr/local/go/src/net/http/client.go:887
	// _ = "end of CoverTab[36847]"
}

// Head issues a HEAD to the specified URL. If the response is one of
//line /usr/local/go/src/net/http/client.go:890
// the following redirect codes, Head follows the redirect, up to a
//line /usr/local/go/src/net/http/client.go:890
// maximum of 10 redirects:
//line /usr/local/go/src/net/http/client.go:890
//
//line /usr/local/go/src/net/http/client.go:890
//	301 (Moved Permanently)
//line /usr/local/go/src/net/http/client.go:890
//	302 (Found)
//line /usr/local/go/src/net/http/client.go:890
//	303 (See Other)
//line /usr/local/go/src/net/http/client.go:890
//	307 (Temporary Redirect)
//line /usr/local/go/src/net/http/client.go:890
//	308 (Permanent Redirect)
//line /usr/local/go/src/net/http/client.go:890
//
//line /usr/local/go/src/net/http/client.go:890
// Head is a wrapper around DefaultClient.Head.
//line /usr/local/go/src/net/http/client.go:890
//
//line /usr/local/go/src/net/http/client.go:890
// To make a request with a specified context.Context, use NewRequestWithContext
//line /usr/local/go/src/net/http/client.go:890
// and DefaultClient.Do.
//line /usr/local/go/src/net/http/client.go:904
func Head(url string) (resp *Response, err error) {
//line /usr/local/go/src/net/http/client.go:904
	_go_fuzz_dep_.CoverTab[36848]++
							return DefaultClient.Head(url)
//line /usr/local/go/src/net/http/client.go:905
	// _ = "end of CoverTab[36848]"
}

// Head issues a HEAD to the specified URL. If the response is one of the
//line /usr/local/go/src/net/http/client.go:908
// following redirect codes, Head follows the redirect after calling the
//line /usr/local/go/src/net/http/client.go:908
// Client's CheckRedirect function:
//line /usr/local/go/src/net/http/client.go:908
//
//line /usr/local/go/src/net/http/client.go:908
//	301 (Moved Permanently)
//line /usr/local/go/src/net/http/client.go:908
//	302 (Found)
//line /usr/local/go/src/net/http/client.go:908
//	303 (See Other)
//line /usr/local/go/src/net/http/client.go:908
//	307 (Temporary Redirect)
//line /usr/local/go/src/net/http/client.go:908
//	308 (Permanent Redirect)
//line /usr/local/go/src/net/http/client.go:908
//
//line /usr/local/go/src/net/http/client.go:908
// To make a request with a specified context.Context, use NewRequestWithContext
//line /usr/local/go/src/net/http/client.go:908
// and Client.Do.
//line /usr/local/go/src/net/http/client.go:920
func (c *Client) Head(url string) (resp *Response, err error) {
//line /usr/local/go/src/net/http/client.go:920
	_go_fuzz_dep_.CoverTab[36849]++
							req, err := NewRequest("HEAD", url, nil)
							if err != nil {
//line /usr/local/go/src/net/http/client.go:922
		_go_fuzz_dep_.CoverTab[36851]++
								return nil, err
//line /usr/local/go/src/net/http/client.go:923
		// _ = "end of CoverTab[36851]"
	} else {
//line /usr/local/go/src/net/http/client.go:924
		_go_fuzz_dep_.CoverTab[36852]++
//line /usr/local/go/src/net/http/client.go:924
		// _ = "end of CoverTab[36852]"
//line /usr/local/go/src/net/http/client.go:924
	}
//line /usr/local/go/src/net/http/client.go:924
	// _ = "end of CoverTab[36849]"
//line /usr/local/go/src/net/http/client.go:924
	_go_fuzz_dep_.CoverTab[36850]++
							return c.Do(req)
//line /usr/local/go/src/net/http/client.go:925
	// _ = "end of CoverTab[36850]"
}

// CloseIdleConnections closes any connections on its Transport which
//line /usr/local/go/src/net/http/client.go:928
// were previously connected from previous requests but are now
//line /usr/local/go/src/net/http/client.go:928
// sitting idle in a "keep-alive" state. It does not interrupt any
//line /usr/local/go/src/net/http/client.go:928
// connections currently in use.
//line /usr/local/go/src/net/http/client.go:928
//
//line /usr/local/go/src/net/http/client.go:928
// If the Client's Transport does not have a CloseIdleConnections method
//line /usr/local/go/src/net/http/client.go:928
// then this method does nothing.
//line /usr/local/go/src/net/http/client.go:935
func (c *Client) CloseIdleConnections() {
//line /usr/local/go/src/net/http/client.go:935
	_go_fuzz_dep_.CoverTab[36853]++
							type closeIdler interface {
		CloseIdleConnections()
	}
	if tr, ok := c.transport().(closeIdler); ok {
//line /usr/local/go/src/net/http/client.go:939
		_go_fuzz_dep_.CoverTab[36854]++
								tr.CloseIdleConnections()
//line /usr/local/go/src/net/http/client.go:940
		// _ = "end of CoverTab[36854]"
	} else {
//line /usr/local/go/src/net/http/client.go:941
		_go_fuzz_dep_.CoverTab[36855]++
//line /usr/local/go/src/net/http/client.go:941
		// _ = "end of CoverTab[36855]"
//line /usr/local/go/src/net/http/client.go:941
	}
//line /usr/local/go/src/net/http/client.go:941
	// _ = "end of CoverTab[36853]"
}

// cancelTimerBody is an io.ReadCloser that wraps rc with two features:
//line /usr/local/go/src/net/http/client.go:944
//  1. On Read error or close, the stop func is called.
//line /usr/local/go/src/net/http/client.go:944
//  2. On Read failure, if reqDidTimeout is true, the error is wrapped and
//line /usr/local/go/src/net/http/client.go:944
//     marked as net.Error that hit its timeout.
//line /usr/local/go/src/net/http/client.go:948
type cancelTimerBody struct {
	stop		func()	// stops the time.Timer waiting to cancel the request
	rc		io.ReadCloser
	reqDidTimeout	func() bool
}

func (b *cancelTimerBody) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/client.go:954
	_go_fuzz_dep_.CoverTab[36856]++
							n, err = b.rc.Read(p)
							if err == nil {
//line /usr/local/go/src/net/http/client.go:956
		_go_fuzz_dep_.CoverTab[36860]++
								return n, nil
//line /usr/local/go/src/net/http/client.go:957
		// _ = "end of CoverTab[36860]"
	} else {
//line /usr/local/go/src/net/http/client.go:958
		_go_fuzz_dep_.CoverTab[36861]++
//line /usr/local/go/src/net/http/client.go:958
		// _ = "end of CoverTab[36861]"
//line /usr/local/go/src/net/http/client.go:958
	}
//line /usr/local/go/src/net/http/client.go:958
	// _ = "end of CoverTab[36856]"
//line /usr/local/go/src/net/http/client.go:958
	_go_fuzz_dep_.CoverTab[36857]++
							if err == io.EOF {
//line /usr/local/go/src/net/http/client.go:959
		_go_fuzz_dep_.CoverTab[36862]++
								return n, err
//line /usr/local/go/src/net/http/client.go:960
		// _ = "end of CoverTab[36862]"
	} else {
//line /usr/local/go/src/net/http/client.go:961
		_go_fuzz_dep_.CoverTab[36863]++
//line /usr/local/go/src/net/http/client.go:961
		// _ = "end of CoverTab[36863]"
//line /usr/local/go/src/net/http/client.go:961
	}
//line /usr/local/go/src/net/http/client.go:961
	// _ = "end of CoverTab[36857]"
//line /usr/local/go/src/net/http/client.go:961
	_go_fuzz_dep_.CoverTab[36858]++
							if b.reqDidTimeout() {
//line /usr/local/go/src/net/http/client.go:962
		_go_fuzz_dep_.CoverTab[36864]++
								err = &httpError{
			err:		err.Error() + " (Client.Timeout or context cancellation while reading body)",
			timeout:	true,
		}
//line /usr/local/go/src/net/http/client.go:966
		// _ = "end of CoverTab[36864]"
	} else {
//line /usr/local/go/src/net/http/client.go:967
		_go_fuzz_dep_.CoverTab[36865]++
//line /usr/local/go/src/net/http/client.go:967
		// _ = "end of CoverTab[36865]"
//line /usr/local/go/src/net/http/client.go:967
	}
//line /usr/local/go/src/net/http/client.go:967
	// _ = "end of CoverTab[36858]"
//line /usr/local/go/src/net/http/client.go:967
	_go_fuzz_dep_.CoverTab[36859]++
							return n, err
//line /usr/local/go/src/net/http/client.go:968
	// _ = "end of CoverTab[36859]"
}

func (b *cancelTimerBody) Close() error {
//line /usr/local/go/src/net/http/client.go:971
	_go_fuzz_dep_.CoverTab[36866]++
							err := b.rc.Close()
							b.stop()
							return err
//line /usr/local/go/src/net/http/client.go:974
	// _ = "end of CoverTab[36866]"
}

func shouldCopyHeaderOnRedirect(headerKey string, initial, dest *url.URL) bool {
//line /usr/local/go/src/net/http/client.go:977
	_go_fuzz_dep_.CoverTab[36867]++
							switch CanonicalHeaderKey(headerKey) {
	case "Authorization", "Www-Authenticate", "Cookie", "Cookie2":
//line /usr/local/go/src/net/http/client.go:979
		_go_fuzz_dep_.CoverTab[36869]++

//line /usr/local/go/src/net/http/client.go:993
		ihost := canonicalAddr(initial)
								dhost := canonicalAddr(dest)
								return isDomainOrSubdomain(dhost, ihost)
//line /usr/local/go/src/net/http/client.go:995
		// _ = "end of CoverTab[36869]"
//line /usr/local/go/src/net/http/client.go:995
	default:
//line /usr/local/go/src/net/http/client.go:995
		_go_fuzz_dep_.CoverTab[36870]++
//line /usr/local/go/src/net/http/client.go:995
		// _ = "end of CoverTab[36870]"
	}
//line /usr/local/go/src/net/http/client.go:996
	// _ = "end of CoverTab[36867]"
//line /usr/local/go/src/net/http/client.go:996
	_go_fuzz_dep_.CoverTab[36868]++

							return true
//line /usr/local/go/src/net/http/client.go:998
	// _ = "end of CoverTab[36868]"
}

// isDomainOrSubdomain reports whether sub is a subdomain (or exact
//line /usr/local/go/src/net/http/client.go:1001
// match) of the parent domain.
//line /usr/local/go/src/net/http/client.go:1001
//
//line /usr/local/go/src/net/http/client.go:1001
// Both domains must already be in canonical form.
//line /usr/local/go/src/net/http/client.go:1005
func isDomainOrSubdomain(sub, parent string) bool {
//line /usr/local/go/src/net/http/client.go:1005
	_go_fuzz_dep_.CoverTab[36871]++
							if sub == parent {
//line /usr/local/go/src/net/http/client.go:1006
		_go_fuzz_dep_.CoverTab[36874]++
								return true
//line /usr/local/go/src/net/http/client.go:1007
		// _ = "end of CoverTab[36874]"
	} else {
//line /usr/local/go/src/net/http/client.go:1008
		_go_fuzz_dep_.CoverTab[36875]++
//line /usr/local/go/src/net/http/client.go:1008
		// _ = "end of CoverTab[36875]"
//line /usr/local/go/src/net/http/client.go:1008
	}
//line /usr/local/go/src/net/http/client.go:1008
	// _ = "end of CoverTab[36871]"
//line /usr/local/go/src/net/http/client.go:1008
	_go_fuzz_dep_.CoverTab[36872]++

//line /usr/local/go/src/net/http/client.go:1012
	if !strings.HasSuffix(sub, parent) {
//line /usr/local/go/src/net/http/client.go:1012
		_go_fuzz_dep_.CoverTab[36876]++
								return false
//line /usr/local/go/src/net/http/client.go:1013
		// _ = "end of CoverTab[36876]"
	} else {
//line /usr/local/go/src/net/http/client.go:1014
		_go_fuzz_dep_.CoverTab[36877]++
//line /usr/local/go/src/net/http/client.go:1014
		// _ = "end of CoverTab[36877]"
//line /usr/local/go/src/net/http/client.go:1014
	}
//line /usr/local/go/src/net/http/client.go:1014
	// _ = "end of CoverTab[36872]"
//line /usr/local/go/src/net/http/client.go:1014
	_go_fuzz_dep_.CoverTab[36873]++
							return sub[len(sub)-len(parent)-1] == '.'
//line /usr/local/go/src/net/http/client.go:1015
	// _ = "end of CoverTab[36873]"
}

func stripPassword(u *url.URL) string {
//line /usr/local/go/src/net/http/client.go:1018
	_go_fuzz_dep_.CoverTab[36878]++
							_, passSet := u.User.Password()
							if passSet {
//line /usr/local/go/src/net/http/client.go:1020
		_go_fuzz_dep_.CoverTab[36880]++
								return strings.Replace(u.String(), u.User.String()+"@", u.User.Username()+":***@", 1)
//line /usr/local/go/src/net/http/client.go:1021
		// _ = "end of CoverTab[36880]"
	} else {
//line /usr/local/go/src/net/http/client.go:1022
		_go_fuzz_dep_.CoverTab[36881]++
//line /usr/local/go/src/net/http/client.go:1022
		// _ = "end of CoverTab[36881]"
//line /usr/local/go/src/net/http/client.go:1022
	}
//line /usr/local/go/src/net/http/client.go:1022
	// _ = "end of CoverTab[36878]"
//line /usr/local/go/src/net/http/client.go:1022
	_go_fuzz_dep_.CoverTab[36879]++
							return u.String()
//line /usr/local/go/src/net/http/client.go:1023
	// _ = "end of CoverTab[36879]"
}

//line /usr/local/go/src/net/http/client.go:1024
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/client.go:1024
var _ = _go_fuzz_dep_.CoverTab
