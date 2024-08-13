// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:5
package oauth2

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:5
)

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2/internal"
)

// expiryDelta determines how earlier a token should be considered
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:19
// expired than its actual expiration time. It is used to avoid late
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:19
// expirations due to client-server time mismatches.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:22
const expiryDelta = 10 * time.Second

// Token represents the credentials used to authorize
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:24
// the requests to access protected resources on the OAuth 2.0
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:24
// provider's backend.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:24
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:24
// Most users of this package should not access fields of Token
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:24
// directly. They're exported mostly for use by related packages
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:24
// implementing derivative OAuth2 flows.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:31
type Token struct {
	// AccessToken is the token that authorizes and authenticates
	// the requests.
	AccessToken	string	`json:"access_token"`

	// TokenType is the type of token.
	// The Type method returns either this or "Bearer", the default.
	TokenType	string	`json:"token_type,omitempty"`

	// RefreshToken is a token that's used by the application
	// (as opposed to the user) to refresh the access token
	// if it expires.
	RefreshToken	string	`json:"refresh_token,omitempty"`

	// Expiry is the optional expiration time of the access token.
	//
	// If zero, TokenSource implementations will reuse the same
	// token forever and RefreshToken or equivalent
	// mechanisms for that TokenSource will not be used.
	Expiry	time.Time	`json:"expiry,omitempty"`

	// raw optionally contains extra metadata from the server
	// when updating a token.
	raw	interface{}
}

// Type returns t.TokenType if non-empty, else "Bearer".
func (t *Token) Type() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:58
	_go_fuzz_dep_.CoverTab[184222]++
										if strings.EqualFold(t.TokenType, "bearer") {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:59
		_go_fuzz_dep_.CoverTab[184227]++
											return "Bearer"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:60
		// _ = "end of CoverTab[184227]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:61
		_go_fuzz_dep_.CoverTab[184228]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:61
		// _ = "end of CoverTab[184228]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:61
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:61
	// _ = "end of CoverTab[184222]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:61
	_go_fuzz_dep_.CoverTab[184223]++
										if strings.EqualFold(t.TokenType, "mac") {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:62
		_go_fuzz_dep_.CoverTab[184229]++
											return "MAC"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:63
		// _ = "end of CoverTab[184229]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:64
		_go_fuzz_dep_.CoverTab[184230]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:64
		// _ = "end of CoverTab[184230]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:64
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:64
	// _ = "end of CoverTab[184223]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:64
	_go_fuzz_dep_.CoverTab[184224]++
										if strings.EqualFold(t.TokenType, "basic") {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:65
		_go_fuzz_dep_.CoverTab[184231]++
											return "Basic"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:66
		// _ = "end of CoverTab[184231]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:67
		_go_fuzz_dep_.CoverTab[184232]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:67
		// _ = "end of CoverTab[184232]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:67
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:67
	// _ = "end of CoverTab[184224]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:67
	_go_fuzz_dep_.CoverTab[184225]++
										if t.TokenType != "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:68
		_go_fuzz_dep_.CoverTab[184233]++
											return t.TokenType
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:69
		// _ = "end of CoverTab[184233]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:70
		_go_fuzz_dep_.CoverTab[184234]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:70
		// _ = "end of CoverTab[184234]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:70
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:70
	// _ = "end of CoverTab[184225]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:70
	_go_fuzz_dep_.CoverTab[184226]++
										return "Bearer"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:71
	// _ = "end of CoverTab[184226]"
}

// SetAuthHeader sets the Authorization header to r using the access
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:74
// token in t.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:74
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:74
// This method is unnecessary when using Transport or an HTTP Client
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:74
// returned by this package.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:79
func (t *Token) SetAuthHeader(r *http.Request) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:79
	_go_fuzz_dep_.CoverTab[184235]++
										r.Header.Set("Authorization", t.Type()+" "+t.AccessToken)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:80
	// _ = "end of CoverTab[184235]"
}

// WithExtra returns a new Token that's a clone of t, but using the
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:83
// provided raw extra map. This is only intended for use by packages
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:83
// implementing derivative OAuth2 flows.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:86
func (t *Token) WithExtra(extra interface{}) *Token {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:86
	_go_fuzz_dep_.CoverTab[184236]++
										t2 := new(Token)
										*t2 = *t
										t2.raw = extra
										return t2
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:90
	// _ = "end of CoverTab[184236]"
}

// Extra returns an extra field.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:93
// Extra fields are key-value pairs returned by the server as a
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:93
// part of the token retrieval response.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:96
func (t *Token) Extra(key string) interface{} {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:96
	_go_fuzz_dep_.CoverTab[184237]++
										if raw, ok := t.raw.(map[string]interface{}); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:97
		_go_fuzz_dep_.CoverTab[184241]++
											return raw[key]
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:98
		// _ = "end of CoverTab[184241]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:99
		_go_fuzz_dep_.CoverTab[184242]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:99
		// _ = "end of CoverTab[184242]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:99
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:99
	// _ = "end of CoverTab[184237]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:99
	_go_fuzz_dep_.CoverTab[184238]++

										vals, ok := t.raw.(url.Values)
										if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:102
		_go_fuzz_dep_.CoverTab[184243]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:103
		// _ = "end of CoverTab[184243]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:104
		_go_fuzz_dep_.CoverTab[184244]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:104
		// _ = "end of CoverTab[184244]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:104
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:104
	// _ = "end of CoverTab[184238]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:104
	_go_fuzz_dep_.CoverTab[184239]++

										v := vals.Get(key)
										switch s := strings.TrimSpace(v); strings.Count(s, ".") {
	case 0:
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:108
		_go_fuzz_dep_.CoverTab[184245]++
											if i, err := strconv.ParseInt(s, 10, 64); err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:109
			_go_fuzz_dep_.CoverTab[184248]++
												return i
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:110
			// _ = "end of CoverTab[184248]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:111
			_go_fuzz_dep_.CoverTab[184249]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:111
			// _ = "end of CoverTab[184249]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:111
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:111
		// _ = "end of CoverTab[184245]"
	case 1:
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:112
		_go_fuzz_dep_.CoverTab[184246]++
											if f, err := strconv.ParseFloat(s, 64); err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:113
			_go_fuzz_dep_.CoverTab[184250]++
												return f
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:114
			// _ = "end of CoverTab[184250]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:115
			_go_fuzz_dep_.CoverTab[184251]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:115
			// _ = "end of CoverTab[184251]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:115
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:115
		// _ = "end of CoverTab[184246]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:115
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:115
		_go_fuzz_dep_.CoverTab[184247]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:115
		// _ = "end of CoverTab[184247]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:116
	// _ = "end of CoverTab[184239]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:116
	_go_fuzz_dep_.CoverTab[184240]++

										return v
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:118
	// _ = "end of CoverTab[184240]"
}

// timeNow is time.Now but pulled out as a variable for tests.
var timeNow = time.Now

// expired reports whether the token is expired.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:124
// t must be non-nil.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:126
func (t *Token) expired() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:126
	_go_fuzz_dep_.CoverTab[184252]++
										if t.Expiry.IsZero() {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:127
		_go_fuzz_dep_.CoverTab[184254]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:128
		// _ = "end of CoverTab[184254]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:129
		_go_fuzz_dep_.CoverTab[184255]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:129
		// _ = "end of CoverTab[184255]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:129
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:129
	// _ = "end of CoverTab[184252]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:129
	_go_fuzz_dep_.CoverTab[184253]++
										return t.Expiry.Round(0).Add(-expiryDelta).Before(timeNow())
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:130
	// _ = "end of CoverTab[184253]"
}

// Valid reports whether t is non-nil, has an AccessToken, and is not expired.
func (t *Token) Valid() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:134
	_go_fuzz_dep_.CoverTab[184256]++
										return t != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:135
		_go_fuzz_dep_.CoverTab[184257]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:135
		return t.AccessToken != ""
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:135
		// _ = "end of CoverTab[184257]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:135
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:135
		_go_fuzz_dep_.CoverTab[184258]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:135
		return !t.expired()
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:135
		// _ = "end of CoverTab[184258]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:135
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:135
	// _ = "end of CoverTab[184256]"
}

// tokenFromInternal maps an *internal.Token struct into
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:138
// a *Token struct.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:140
func tokenFromInternal(t *internal.Token) *Token {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:140
	_go_fuzz_dep_.CoverTab[184259]++
										if t == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:141
		_go_fuzz_dep_.CoverTab[184261]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:142
		// _ = "end of CoverTab[184261]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:143
		_go_fuzz_dep_.CoverTab[184262]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:143
		// _ = "end of CoverTab[184262]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:143
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:143
	// _ = "end of CoverTab[184259]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:143
	_go_fuzz_dep_.CoverTab[184260]++
										return &Token{
		AccessToken:	t.AccessToken,
		TokenType:	t.TokenType,
		RefreshToken:	t.RefreshToken,
		Expiry:		t.Expiry,
		raw:		t.Raw,
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:150
	// _ = "end of CoverTab[184260]"
}

// retrieveToken takes a *Config and uses that to retrieve an *internal.Token.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:153
// This token is then mapped from *internal.Token into an *oauth2.Token which is returned along
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:153
// with an error..
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:156
func retrieveToken(ctx context.Context, c *Config, v url.Values) (*Token, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:156
	_go_fuzz_dep_.CoverTab[184263]++
										tk, err := internal.RetrieveToken(ctx, c.ClientID, c.ClientSecret, c.Endpoint.TokenURL, v, internal.AuthStyle(c.Endpoint.AuthStyle))
										if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:158
		_go_fuzz_dep_.CoverTab[184265]++
											if rErr, ok := err.(*internal.RetrieveError); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:159
			_go_fuzz_dep_.CoverTab[184267]++
												return nil, (*RetrieveError)(rErr)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:160
			// _ = "end of CoverTab[184267]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:161
			_go_fuzz_dep_.CoverTab[184268]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:161
			// _ = "end of CoverTab[184268]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:161
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:161
		// _ = "end of CoverTab[184265]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:161
		_go_fuzz_dep_.CoverTab[184266]++
											return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:162
		// _ = "end of CoverTab[184266]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:163
		_go_fuzz_dep_.CoverTab[184269]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:163
		// _ = "end of CoverTab[184269]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:163
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:163
	// _ = "end of CoverTab[184263]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:163
	_go_fuzz_dep_.CoverTab[184264]++
										return tokenFromInternal(tk), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:164
	// _ = "end of CoverTab[184264]"
}

// RetrieveError is the error returned when the token endpoint returns a
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:167
// non-2XX HTTP status code.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:169
type RetrieveError struct {
	Response	*http.Response
	// Body is the body that was consumed by reading Response.Body.
	// It may be truncated.
	Body	[]byte
}

func (r *RetrieveError) Error() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:176
	_go_fuzz_dep_.CoverTab[184270]++
										return fmt.Sprintf("oauth2: cannot fetch token: %v\nResponse: %s", r.Response.Status, r.Body)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:177
	// _ = "end of CoverTab[184270]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:178
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/token.go:178
var _ = _go_fuzz_dep_.CoverTab
