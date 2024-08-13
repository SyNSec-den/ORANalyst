// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:5
package oauth2

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:5
)

import (
	"errors"
	"log"
	"net/http"
	"sync"
)

// Transport is an http.RoundTripper that makes OAuth 2.0 HTTP requests,
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:14
// wrapping a base RoundTripper and adding an Authorization header
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:14
// with a token from the supplied Sources.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:14
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:14
// Transport is a low-level mechanism. Most code will use the
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:14
// higher-level Config.Client method instead.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:20
type Transport struct {
	// Source supplies the token to add to outgoing requests'
	// Authorization headers.
	Source	TokenSource

	// Base is the base RoundTripper used to make HTTP requests.
	// If nil, http.DefaultTransport is used.
	Base	http.RoundTripper
}

// RoundTrip authorizes and authenticates the request with an
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:30
// access token from Transport's Source.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:32
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:32
	_go_fuzz_dep_.CoverTab[184271]++
										reqBodyClosed := false
										if req.Body != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:34
		_go_fuzz_dep_.CoverTab[184275]++
											defer func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:35
			_go_fuzz_dep_.CoverTab[184276]++
												if !reqBodyClosed {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:36
				_go_fuzz_dep_.CoverTab[184277]++
													req.Body.Close()
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:37
				// _ = "end of CoverTab[184277]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:38
				_go_fuzz_dep_.CoverTab[184278]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:38
				// _ = "end of CoverTab[184278]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:38
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:38
			// _ = "end of CoverTab[184276]"
		}()
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:39
		// _ = "end of CoverTab[184275]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:40
		_go_fuzz_dep_.CoverTab[184279]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:40
		// _ = "end of CoverTab[184279]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:40
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:40
	// _ = "end of CoverTab[184271]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:40
	_go_fuzz_dep_.CoverTab[184272]++

										if t.Source == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:42
		_go_fuzz_dep_.CoverTab[184280]++
											return nil, errors.New("oauth2: Transport's Source is nil")
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:43
		// _ = "end of CoverTab[184280]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:44
		_go_fuzz_dep_.CoverTab[184281]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:44
		// _ = "end of CoverTab[184281]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:44
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:44
	// _ = "end of CoverTab[184272]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:44
	_go_fuzz_dep_.CoverTab[184273]++
										token, err := t.Source.Token()
										if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:46
		_go_fuzz_dep_.CoverTab[184282]++
											return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:47
		// _ = "end of CoverTab[184282]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:48
		_go_fuzz_dep_.CoverTab[184283]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:48
		// _ = "end of CoverTab[184283]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:48
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:48
	// _ = "end of CoverTab[184273]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:48
	_go_fuzz_dep_.CoverTab[184274]++

										req2 := cloneRequest(req)
										token.SetAuthHeader(req2)

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:54
	reqBodyClosed = true
										return t.base().RoundTrip(req2)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:55
	// _ = "end of CoverTab[184274]"
}

var cancelOnce sync.Once

// CancelRequest does nothing. It used to be a legacy cancellation mechanism
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:60
// but now only it only logs on first use to warn that it's deprecated.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:60
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:60
// Deprecated: use contexts for cancellation instead.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:64
func (t *Transport) CancelRequest(req *http.Request) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:64
	_go_fuzz_dep_.CoverTab[184284]++
										cancelOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:65
		_go_fuzz_dep_.CoverTab[184285]++
											log.Printf("deprecated: golang.org/x/oauth2: Transport.CancelRequest no longer does anything; use contexts")
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:66
		// _ = "end of CoverTab[184285]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:67
	// _ = "end of CoverTab[184284]"
}

func (t *Transport) base() http.RoundTripper {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:70
	_go_fuzz_dep_.CoverTab[184286]++
										if t.Base != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:71
		_go_fuzz_dep_.CoverTab[184288]++
											return t.Base
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:72
		// _ = "end of CoverTab[184288]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:73
		_go_fuzz_dep_.CoverTab[184289]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:73
		// _ = "end of CoverTab[184289]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:73
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:73
	// _ = "end of CoverTab[184286]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:73
	_go_fuzz_dep_.CoverTab[184287]++
										return http.DefaultTransport
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:74
	// _ = "end of CoverTab[184287]"
}

// cloneRequest returns a clone of the provided *http.Request.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:77
// The clone is a shallow copy of the struct and its Header map.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:79
func cloneRequest(r *http.Request) *http.Request {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:79
	_go_fuzz_dep_.CoverTab[184290]++

										r2 := new(http.Request)
										*r2 = *r

										r2.Header = make(http.Header, len(r.Header))
										for k, s := range r.Header {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:85
		_go_fuzz_dep_.CoverTab[184292]++
											r2.Header[k] = append([]string(nil), s...)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:86
		// _ = "end of CoverTab[184292]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:87
	// _ = "end of CoverTab[184290]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:87
	_go_fuzz_dep_.CoverTab[184291]++
										return r2
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:88
	// _ = "end of CoverTab[184291]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:89
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/transport.go:89
var _ = _go_fuzz_dep_.CoverTab
