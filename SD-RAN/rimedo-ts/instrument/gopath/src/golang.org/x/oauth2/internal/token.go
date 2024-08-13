// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:5
package internal

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:5
)

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"mime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context/ctxhttp"
)

// Token represents the credentials used to authorize
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:26
// the requests to access protected resources on the OAuth 2.0
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:26
// provider's backend.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:26
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:26
// This type is a mirror of oauth2.Token and exists to break
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:26
// an otherwise-circular dependency. Other internal packages
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:26
// should convert this Token into an oauth2.Token before use.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:33
type Token struct {
	// AccessToken is the token that authorizes and authenticates
	// the requests.
	AccessToken	string

	// TokenType is the type of token.
	// The Type method returns either this or "Bearer", the default.
	TokenType	string

	// RefreshToken is a token that's used by the application
	// (as opposed to the user) to refresh the access token
	// if it expires.
	RefreshToken	string

	// Expiry is the optional expiration time of the access token.
	//
	// If zero, TokenSource implementations will reuse the same
	// token forever and RefreshToken or equivalent
	// mechanisms for that TokenSource will not be used.
	Expiry	time.Time

	// Raw optionally contains extra metadata from the server
	// when updating a token.
	Raw	interface{}
}

// tokenJSON is the struct representing the HTTP response from OAuth2
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:59
// providers returning a token in JSON form.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:61
type tokenJSON struct {
	AccessToken	string		`json:"access_token"`
	TokenType	string		`json:"token_type"`
	RefreshToken	string		`json:"refresh_token"`
	ExpiresIn	expirationTime	`json:"expires_in"`	// at least PayPal returns string, while most return number
}

func (e *tokenJSON) expiry() (t time.Time) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:68
	_go_fuzz_dep_.CoverTab[184060]++
											if v := e.ExpiresIn; v != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:69
		_go_fuzz_dep_.CoverTab[184062]++
												return time.Now().Add(time.Duration(v) * time.Second)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:70
		// _ = "end of CoverTab[184062]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:71
		_go_fuzz_dep_.CoverTab[184063]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:71
		// _ = "end of CoverTab[184063]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:71
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:71
	// _ = "end of CoverTab[184060]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:71
	_go_fuzz_dep_.CoverTab[184061]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:72
	// _ = "end of CoverTab[184061]"
}

type expirationTime int32

func (e *expirationTime) UnmarshalJSON(b []byte) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:77
	_go_fuzz_dep_.CoverTab[184064]++
											if len(b) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:78
		_go_fuzz_dep_.CoverTab[184069]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:78
		return string(b) == "null"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:78
		// _ = "end of CoverTab[184069]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:78
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:78
		_go_fuzz_dep_.CoverTab[184070]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:79
		// _ = "end of CoverTab[184070]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:80
		_go_fuzz_dep_.CoverTab[184071]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:80
		// _ = "end of CoverTab[184071]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:80
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:80
	// _ = "end of CoverTab[184064]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:80
	_go_fuzz_dep_.CoverTab[184065]++
											var n json.Number
											err := json.Unmarshal(b, &n)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:83
		_go_fuzz_dep_.CoverTab[184072]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:84
		// _ = "end of CoverTab[184072]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:85
		_go_fuzz_dep_.CoverTab[184073]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:85
		// _ = "end of CoverTab[184073]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:85
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:85
	// _ = "end of CoverTab[184065]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:85
	_go_fuzz_dep_.CoverTab[184066]++
											i, err := n.Int64()
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:87
		_go_fuzz_dep_.CoverTab[184074]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:88
		// _ = "end of CoverTab[184074]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:89
		_go_fuzz_dep_.CoverTab[184075]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:89
		// _ = "end of CoverTab[184075]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:89
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:89
	// _ = "end of CoverTab[184066]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:89
	_go_fuzz_dep_.CoverTab[184067]++
											if i > math.MaxInt32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:90
		_go_fuzz_dep_.CoverTab[184076]++
												i = math.MaxInt32
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:91
		// _ = "end of CoverTab[184076]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:92
		_go_fuzz_dep_.CoverTab[184077]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:92
		// _ = "end of CoverTab[184077]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:92
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:92
	// _ = "end of CoverTab[184067]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:92
	_go_fuzz_dep_.CoverTab[184068]++
											*e = expirationTime(i)
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:94
	// _ = "end of CoverTab[184068]"
}

// RegisterBrokenAuthHeaderProvider previously did something. It is now a no-op.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:97
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:97
// Deprecated: this function no longer does anything. Caller code that
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:97
// wants to avoid potential extra HTTP requests made during
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:97
// auto-probing of the provider's auth style should set
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:97
// Endpoint.AuthStyle.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:103
func RegisterBrokenAuthHeaderProvider(tokenURL string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:103
	_go_fuzz_dep_.CoverTab[184078]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:103
	// _ = "end of CoverTab[184078]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:103
}

// AuthStyle is a copy of the golang.org/x/oauth2 package's AuthStyle type.
type AuthStyle int

const (
	AuthStyleUnknown	AuthStyle	= 0
	AuthStyleInParams	AuthStyle	= 1
	AuthStyleInHeader	AuthStyle	= 2
)

// authStyleCache is the set of tokenURLs we've successfully used via
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:114
// RetrieveToken and which style auth we ended up using.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:114
// It's called a cache, but it doesn't (yet?) shrink. It's expected that
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:114
// the set of OAuth2 servers a program contacts over time is fixed and
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:114
// small.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:119
var authStyleCache struct {
	sync.Mutex
	m	map[string]AuthStyle	// keyed by tokenURL
}

// ResetAuthCache resets the global authentication style cache used
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:124
// for AuthStyleUnknown token requests.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:126
func ResetAuthCache() {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:126
	_go_fuzz_dep_.CoverTab[184079]++
											authStyleCache.Lock()
											defer authStyleCache.Unlock()
											authStyleCache.m = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:129
	// _ = "end of CoverTab[184079]"
}

// lookupAuthStyle reports which auth style we last used with tokenURL
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:132
// when calling RetrieveToken and whether we have ever done so.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:134
func lookupAuthStyle(tokenURL string) (style AuthStyle, ok bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:134
	_go_fuzz_dep_.CoverTab[184080]++
											authStyleCache.Lock()
											defer authStyleCache.Unlock()
											style, ok = authStyleCache.m[tokenURL]
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:138
	// _ = "end of CoverTab[184080]"
}

// setAuthStyle adds an entry to authStyleCache, documented above.
func setAuthStyle(tokenURL string, v AuthStyle) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:142
	_go_fuzz_dep_.CoverTab[184081]++
											authStyleCache.Lock()
											defer authStyleCache.Unlock()
											if authStyleCache.m == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:145
		_go_fuzz_dep_.CoverTab[184083]++
												authStyleCache.m = make(map[string]AuthStyle)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:146
		// _ = "end of CoverTab[184083]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:147
		_go_fuzz_dep_.CoverTab[184084]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:147
		// _ = "end of CoverTab[184084]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:147
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:147
	// _ = "end of CoverTab[184081]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:147
	_go_fuzz_dep_.CoverTab[184082]++
											authStyleCache.m[tokenURL] = v
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:148
	// _ = "end of CoverTab[184082]"
}

// newTokenRequest returns a new *http.Request to retrieve a new token
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:151
// from tokenURL using the provided clientID, clientSecret, and POST
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:151
// body parameters.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:151
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:151
// inParams is whether the clientID & clientSecret should be encoded
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:151
// as the POST body. An 'inParams' value of true means to send it in
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:151
// the POST body (along with any values in v); false means to send it
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:151
// in the Authorization header.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:159
func newTokenRequest(tokenURL, clientID, clientSecret string, v url.Values, authStyle AuthStyle) (*http.Request, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:159
	_go_fuzz_dep_.CoverTab[184085]++
											if authStyle == AuthStyleInParams {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:160
		_go_fuzz_dep_.CoverTab[184089]++
												v = cloneURLValues(v)
												if clientID != "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:162
			_go_fuzz_dep_.CoverTab[184091]++
													v.Set("client_id", clientID)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:163
			// _ = "end of CoverTab[184091]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:164
			_go_fuzz_dep_.CoverTab[184092]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:164
			// _ = "end of CoverTab[184092]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:164
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:164
		// _ = "end of CoverTab[184089]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:164
		_go_fuzz_dep_.CoverTab[184090]++
												if clientSecret != "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:165
			_go_fuzz_dep_.CoverTab[184093]++
													v.Set("client_secret", clientSecret)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:166
			// _ = "end of CoverTab[184093]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:167
			_go_fuzz_dep_.CoverTab[184094]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:167
			// _ = "end of CoverTab[184094]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:167
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:167
		// _ = "end of CoverTab[184090]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:168
		_go_fuzz_dep_.CoverTab[184095]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:168
		// _ = "end of CoverTab[184095]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:168
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:168
	// _ = "end of CoverTab[184085]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:168
	_go_fuzz_dep_.CoverTab[184086]++
											req, err := http.NewRequest("POST", tokenURL, strings.NewReader(v.Encode()))
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:170
		_go_fuzz_dep_.CoverTab[184096]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:171
		// _ = "end of CoverTab[184096]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:172
		_go_fuzz_dep_.CoverTab[184097]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:172
		// _ = "end of CoverTab[184097]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:172
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:172
	// _ = "end of CoverTab[184086]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:172
	_go_fuzz_dep_.CoverTab[184087]++
											req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
											if authStyle == AuthStyleInHeader {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:174
		_go_fuzz_dep_.CoverTab[184098]++
												req.SetBasicAuth(url.QueryEscape(clientID), url.QueryEscape(clientSecret))
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:175
		// _ = "end of CoverTab[184098]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:176
		_go_fuzz_dep_.CoverTab[184099]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:176
		// _ = "end of CoverTab[184099]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:176
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:176
	// _ = "end of CoverTab[184087]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:176
	_go_fuzz_dep_.CoverTab[184088]++
											return req, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:177
	// _ = "end of CoverTab[184088]"
}

func cloneURLValues(v url.Values) url.Values {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:180
	_go_fuzz_dep_.CoverTab[184100]++
											v2 := make(url.Values, len(v))
											for k, vv := range v {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:182
		_go_fuzz_dep_.CoverTab[184102]++
												v2[k] = append([]string(nil), vv...)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:183
		// _ = "end of CoverTab[184102]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:184
	// _ = "end of CoverTab[184100]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:184
	_go_fuzz_dep_.CoverTab[184101]++
											return v2
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:185
	// _ = "end of CoverTab[184101]"
}

func RetrieveToken(ctx context.Context, clientID, clientSecret, tokenURL string, v url.Values, authStyle AuthStyle) (*Token, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:188
	_go_fuzz_dep_.CoverTab[184103]++
											needsAuthStyleProbe := authStyle == 0
											if needsAuthStyleProbe {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:190
		_go_fuzz_dep_.CoverTab[184109]++
												if style, ok := lookupAuthStyle(tokenURL); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:191
			_go_fuzz_dep_.CoverTab[184110]++
													authStyle = style
													needsAuthStyleProbe = false
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:193
			// _ = "end of CoverTab[184110]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:194
			_go_fuzz_dep_.CoverTab[184111]++
													authStyle = AuthStyleInHeader
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:195
			// _ = "end of CoverTab[184111]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:196
		// _ = "end of CoverTab[184109]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:197
		_go_fuzz_dep_.CoverTab[184112]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:197
		// _ = "end of CoverTab[184112]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:197
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:197
	// _ = "end of CoverTab[184103]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:197
	_go_fuzz_dep_.CoverTab[184104]++
											req, err := newTokenRequest(tokenURL, clientID, clientSecret, v, authStyle)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:199
		_go_fuzz_dep_.CoverTab[184113]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:200
		// _ = "end of CoverTab[184113]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:201
		_go_fuzz_dep_.CoverTab[184114]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:201
		// _ = "end of CoverTab[184114]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:201
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:201
	// _ = "end of CoverTab[184104]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:201
	_go_fuzz_dep_.CoverTab[184105]++
											token, err := doTokenRoundTrip(ctx, req)
											if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:203
		_go_fuzz_dep_.CoverTab[184115]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:203
		return needsAuthStyleProbe
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:203
		// _ = "end of CoverTab[184115]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:203
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:203
		_go_fuzz_dep_.CoverTab[184116]++

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:216
		authStyle = AuthStyleInParams
												req, _ = newTokenRequest(tokenURL, clientID, clientSecret, v, authStyle)
												token, err = doTokenRoundTrip(ctx, req)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:218
		// _ = "end of CoverTab[184116]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:219
		_go_fuzz_dep_.CoverTab[184117]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:219
		// _ = "end of CoverTab[184117]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:219
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:219
	// _ = "end of CoverTab[184105]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:219
	_go_fuzz_dep_.CoverTab[184106]++
											if needsAuthStyleProbe && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:220
		_go_fuzz_dep_.CoverTab[184118]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:220
		return err == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:220
		// _ = "end of CoverTab[184118]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:220
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:220
		_go_fuzz_dep_.CoverTab[184119]++
												setAuthStyle(tokenURL, authStyle)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:221
		// _ = "end of CoverTab[184119]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:222
		_go_fuzz_dep_.CoverTab[184120]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:222
		// _ = "end of CoverTab[184120]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:222
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:222
	// _ = "end of CoverTab[184106]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:222
	_go_fuzz_dep_.CoverTab[184107]++

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:225
	if token != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:225
		_go_fuzz_dep_.CoverTab[184121]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:225
		return token.RefreshToken == ""
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:225
		// _ = "end of CoverTab[184121]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:225
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:225
		_go_fuzz_dep_.CoverTab[184122]++
												token.RefreshToken = v.Get("refresh_token")
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:226
		// _ = "end of CoverTab[184122]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:227
		_go_fuzz_dep_.CoverTab[184123]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:227
		// _ = "end of CoverTab[184123]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:227
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:227
	// _ = "end of CoverTab[184107]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:227
	_go_fuzz_dep_.CoverTab[184108]++
											return token, err
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:228
	// _ = "end of CoverTab[184108]"
}

func doTokenRoundTrip(ctx context.Context, req *http.Request) (*Token, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:231
	_go_fuzz_dep_.CoverTab[184124]++
											r, err := ctxhttp.Do(ctx, ContextClient(ctx), req)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:233
		_go_fuzz_dep_.CoverTab[184130]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:234
		// _ = "end of CoverTab[184130]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:235
		_go_fuzz_dep_.CoverTab[184131]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:235
		// _ = "end of CoverTab[184131]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:235
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:235
	// _ = "end of CoverTab[184124]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:235
	_go_fuzz_dep_.CoverTab[184125]++
											body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))
											r.Body.Close()
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:238
		_go_fuzz_dep_.CoverTab[184132]++
												return nil, fmt.Errorf("oauth2: cannot fetch token: %v", err)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:239
		// _ = "end of CoverTab[184132]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:240
		_go_fuzz_dep_.CoverTab[184133]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:240
		// _ = "end of CoverTab[184133]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:240
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:240
	// _ = "end of CoverTab[184125]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:240
	_go_fuzz_dep_.CoverTab[184126]++
											if code := r.StatusCode; code < 200 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:241
		_go_fuzz_dep_.CoverTab[184134]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:241
		return code > 299
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:241
		// _ = "end of CoverTab[184134]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:241
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:241
		_go_fuzz_dep_.CoverTab[184135]++
												return nil, &RetrieveError{
			Response:	r,
			Body:		body,
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:245
		// _ = "end of CoverTab[184135]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:246
		_go_fuzz_dep_.CoverTab[184136]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:246
		// _ = "end of CoverTab[184136]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:246
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:246
	// _ = "end of CoverTab[184126]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:246
	_go_fuzz_dep_.CoverTab[184127]++

											var token *Token
											content, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
											switch content {
	case "application/x-www-form-urlencoded", "text/plain":
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:251
		_go_fuzz_dep_.CoverTab[184137]++
												vals, err := url.ParseQuery(string(body))
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:253
			_go_fuzz_dep_.CoverTab[184141]++
													return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:254
			// _ = "end of CoverTab[184141]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:255
			_go_fuzz_dep_.CoverTab[184142]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:255
			// _ = "end of CoverTab[184142]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:255
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:255
		// _ = "end of CoverTab[184137]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:255
		_go_fuzz_dep_.CoverTab[184138]++
												token = &Token{
			AccessToken:	vals.Get("access_token"),
			TokenType:	vals.Get("token_type"),
			RefreshToken:	vals.Get("refresh_token"),
			Raw:		vals,
		}
		e := vals.Get("expires_in")
		expires, _ := strconv.Atoi(e)
		if expires != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:264
			_go_fuzz_dep_.CoverTab[184143]++
													token.Expiry = time.Now().Add(time.Duration(expires) * time.Second)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:265
			// _ = "end of CoverTab[184143]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:266
			_go_fuzz_dep_.CoverTab[184144]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:266
			// _ = "end of CoverTab[184144]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:266
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:266
		// _ = "end of CoverTab[184138]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:267
		_go_fuzz_dep_.CoverTab[184139]++
												var tj tokenJSON
												if err = json.Unmarshal(body, &tj); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:269
			_go_fuzz_dep_.CoverTab[184145]++
													return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:270
			// _ = "end of CoverTab[184145]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:271
			_go_fuzz_dep_.CoverTab[184146]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:271
			// _ = "end of CoverTab[184146]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:271
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:271
		// _ = "end of CoverTab[184139]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:271
		_go_fuzz_dep_.CoverTab[184140]++
												token = &Token{
			AccessToken:	tj.AccessToken,
			TokenType:	tj.TokenType,
			RefreshToken:	tj.RefreshToken,
			Expiry:		tj.expiry(),
			Raw:		make(map[string]interface{}),
		}
												json.Unmarshal(body, &token.Raw)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:279
		// _ = "end of CoverTab[184140]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:280
	// _ = "end of CoverTab[184127]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:280
	_go_fuzz_dep_.CoverTab[184128]++
											if token.AccessToken == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:281
		_go_fuzz_dep_.CoverTab[184147]++
												return nil, errors.New("oauth2: server response missing access_token")
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:282
		// _ = "end of CoverTab[184147]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:283
		_go_fuzz_dep_.CoverTab[184148]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:283
		// _ = "end of CoverTab[184148]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:283
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:283
	// _ = "end of CoverTab[184128]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:283
	_go_fuzz_dep_.CoverTab[184129]++
											return token, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:284
	// _ = "end of CoverTab[184129]"
}

type RetrieveError struct {
	Response	*http.Response
	Body		[]byte
}

func (r *RetrieveError) Error() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:292
	_go_fuzz_dep_.CoverTab[184149]++
											return fmt.Sprintf("oauth2: cannot fetch token: %v\nResponse: %s", r.Response.Status, r.Body)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:293
	// _ = "end of CoverTab[184149]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:294
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/token.go:294
var _ = _go_fuzz_dep_.CoverTab
