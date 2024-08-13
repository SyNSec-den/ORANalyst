// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:5
// Package oauth2 provides support for making
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:5
// OAuth2 authorized and authenticated HTTP requests,
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:5
// as specified in RFC 6749.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:5
// It can additionally grant authorization with Bearer JWT.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:9
package oauth2

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:9
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:9
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:9
)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:9
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:9
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:9
)

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/oauth2/internal"
)

// NoContext is the default context you should supply if not using
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:23
// your own context.Context (see https://golang.org/x/net/context).
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:23
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:23
// Deprecated: Use context.Background() or context.TODO() instead.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:27
var NoContext = context.TODO()

// RegisterBrokenAuthHeaderProvider previously did something. It is now a no-op.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:29
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:29
// Deprecated: this function no longer does anything. Caller code that
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:29
// wants to avoid potential extra HTTP requests made during
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:29
// auto-probing of the provider's auth style should set
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:29
// Endpoint.AuthStyle.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:35
func RegisterBrokenAuthHeaderProvider(tokenURL string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:35
	_go_fuzz_dep_.CoverTab[184159]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:35
	// _ = "end of CoverTab[184159]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:35
}

// Config describes a typical 3-legged OAuth2 flow, with both the
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:37
// client application information and the server's endpoint URLs.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:37
// For the client credentials 2-legged OAuth2 flow, see the clientcredentials
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:37
// package (https://golang.org/x/oauth2/clientcredentials).
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:41
type Config struct {
	// ClientID is the application's ID.
	ClientID	string

	// ClientSecret is the application's secret.
	ClientSecret	string

	// Endpoint contains the resource server's token endpoint
	// URLs. These are constants specific to each server and are
	// often available via site-specific packages, such as
	// google.Endpoint or github.Endpoint.
	Endpoint	Endpoint

	// RedirectURL is the URL to redirect users going through
	// the OAuth flow, after the resource owner's URLs.
	RedirectURL	string

	// Scope specifies optional requested permissions.
	Scopes	[]string
}

// A TokenSource is anything that can return a token.
type TokenSource interface {
	// Token returns a token or an error.
	// Token must be safe for concurrent use by multiple goroutines.
	// The returned Token must not be modified.
	Token() (*Token, error)
}

// Endpoint represents an OAuth 2.0 provider's authorization and token
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:70
// endpoint URLs.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:72
type Endpoint struct {
	AuthURL		string
	TokenURL	string

	// AuthStyle optionally specifies how the endpoint wants the
	// client ID & client secret sent. The zero value means to
	// auto-detect.
	AuthStyle	AuthStyle
}

// AuthStyle represents how requests for tokens are authenticated
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:82
// to the server.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:84
type AuthStyle int

const (
	// AuthStyleAutoDetect means to auto-detect which authentication
	// style the provider wants by trying both ways and caching
	// the successful way for the future.
	AuthStyleAutoDetect	AuthStyle	= 0

	// AuthStyleInParams sends the "client_id" and "client_secret"
	// in the POST body as application/x-www-form-urlencoded parameters.
	AuthStyleInParams	AuthStyle	= 1

	// AuthStyleInHeader sends the client_id and client_password
	// using HTTP Basic Authorization. This is an optional style
	// described in the OAuth2 RFC 6749 section 2.3.1.
	AuthStyleInHeader	AuthStyle	= 2
)

var (
	// AccessTypeOnline and AccessTypeOffline are options passed
	// to the Options.AuthCodeURL method. They modify the
	// "access_type" field that gets sent in the URL returned by
	// AuthCodeURL.
	//
	// Online is the default if neither is specified. If your
	// application needs to refresh access tokens when the user
	// is not present at the browser, then use offline. This will
	// result in your application obtaining a refresh token the
	// first time your application exchanges an authorization
	// code for a user.
	AccessTypeOnline	AuthCodeOption	= SetAuthURLParam("access_type", "online")
	AccessTypeOffline	AuthCodeOption	= SetAuthURLParam("access_type", "offline")

	// ApprovalForce forces the users to view the consent dialog
	// and confirm the permissions request at the URL returned
	// from AuthCodeURL, even if they've already done so.
	ApprovalForce	AuthCodeOption	= SetAuthURLParam("prompt", "consent")
)

// An AuthCodeOption is passed to Config.AuthCodeURL.
type AuthCodeOption interface {
	setValue(url.Values)
}

type setParam struct{ k, v string }

func (p setParam) setValue(m url.Values) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:130
	_go_fuzz_dep_.CoverTab[184160]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:130
	m.Set(p.k, p.v)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:130
	// _ = "end of CoverTab[184160]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:130
}

// SetAuthURLParam builds an AuthCodeOption which passes key/value parameters
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:132
// to a provider's authorization endpoint.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:134
func SetAuthURLParam(key, value string) AuthCodeOption {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:134
	_go_fuzz_dep_.CoverTab[184161]++
										return setParam{key, value}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:135
	// _ = "end of CoverTab[184161]"
}

// AuthCodeURL returns a URL to OAuth 2.0 provider's consent page
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:138
// that asks for permissions for the required scopes explicitly.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:138
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:138
// State is a token to protect the user from CSRF attacks. You must
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:138
// always provide a non-empty string and validate that it matches the
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:138
// the state query parameter on your redirect callback.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:138
// See http://tools.ietf.org/html/rfc6749#section-10.12 for more info.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:138
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:138
// Opts may include AccessTypeOnline or AccessTypeOffline, as well
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:138
// as ApprovalForce.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:138
// It can also be used to pass the PKCE challenge.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:138
// See https://www.oauth.com/oauth2-servers/pkce/ for more info.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:150
func (c *Config) AuthCodeURL(state string, opts ...AuthCodeOption) string {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:150
	_go_fuzz_dep_.CoverTab[184162]++
										var buf bytes.Buffer
										buf.WriteString(c.Endpoint.AuthURL)
										v := url.Values{
		"response_type":	{"code"},
		"client_id":		{c.ClientID},
	}
	if c.RedirectURL != "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:157
		_go_fuzz_dep_.CoverTab[184168]++
											v.Set("redirect_uri", c.RedirectURL)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:158
		// _ = "end of CoverTab[184168]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:159
		_go_fuzz_dep_.CoverTab[184169]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:159
		// _ = "end of CoverTab[184169]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:159
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:159
	// _ = "end of CoverTab[184162]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:159
	_go_fuzz_dep_.CoverTab[184163]++
										if len(c.Scopes) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:160
		_go_fuzz_dep_.CoverTab[184170]++
											v.Set("scope", strings.Join(c.Scopes, " "))
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:161
		// _ = "end of CoverTab[184170]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:162
		_go_fuzz_dep_.CoverTab[184171]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:162
		// _ = "end of CoverTab[184171]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:162
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:162
	// _ = "end of CoverTab[184163]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:162
	_go_fuzz_dep_.CoverTab[184164]++
										if state != "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:163
		_go_fuzz_dep_.CoverTab[184172]++

											v.Set("state", state)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:165
		// _ = "end of CoverTab[184172]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:166
		_go_fuzz_dep_.CoverTab[184173]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:166
		// _ = "end of CoverTab[184173]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:166
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:166
	// _ = "end of CoverTab[184164]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:166
	_go_fuzz_dep_.CoverTab[184165]++
										for _, opt := range opts {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:167
		_go_fuzz_dep_.CoverTab[184174]++
											opt.setValue(v)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:168
		// _ = "end of CoverTab[184174]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:169
	// _ = "end of CoverTab[184165]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:169
	_go_fuzz_dep_.CoverTab[184166]++
										if strings.Contains(c.Endpoint.AuthURL, "?") {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:170
		_go_fuzz_dep_.CoverTab[184175]++
											buf.WriteByte('&')
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:171
		// _ = "end of CoverTab[184175]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:172
		_go_fuzz_dep_.CoverTab[184176]++
											buf.WriteByte('?')
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:173
		// _ = "end of CoverTab[184176]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:174
	// _ = "end of CoverTab[184166]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:174
	_go_fuzz_dep_.CoverTab[184167]++
										buf.WriteString(v.Encode())
										return buf.String()
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:176
	// _ = "end of CoverTab[184167]"
}

// PasswordCredentialsToken converts a resource owner username and password
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:179
// pair into a token.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:179
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:179
// Per the RFC, this grant type should only be used "when there is a high
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:179
// degree of trust between the resource owner and the client (e.g., the client
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:179
// is part of the device operating system or a highly privileged application),
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:179
// and when other authorization grant types are not available."
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:179
// See https://tools.ietf.org/html/rfc6749#section-4.3 for more info.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:179
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:179
// The provided context optionally controls which HTTP client is used. See the HTTPClient variable.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:189
func (c *Config) PasswordCredentialsToken(ctx context.Context, username, password string) (*Token, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:189
	_go_fuzz_dep_.CoverTab[184177]++
										v := url.Values{
		"grant_type":	{"password"},
		"username":	{username},
		"password":	{password},
	}
	if len(c.Scopes) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:195
		_go_fuzz_dep_.CoverTab[184179]++
											v.Set("scope", strings.Join(c.Scopes, " "))
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:196
		// _ = "end of CoverTab[184179]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:197
		_go_fuzz_dep_.CoverTab[184180]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:197
		// _ = "end of CoverTab[184180]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:197
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:197
	// _ = "end of CoverTab[184177]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:197
	_go_fuzz_dep_.CoverTab[184178]++
										return retrieveToken(ctx, c, v)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:198
	// _ = "end of CoverTab[184178]"
}

// Exchange converts an authorization code into a token.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:201
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:201
// It is used after a resource provider redirects the user back
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:201
// to the Redirect URI (the URL obtained from AuthCodeURL).
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:201
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:201
// The provided context optionally controls which HTTP client is used. See the HTTPClient variable.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:201
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:201
// The code will be in the *http.Request.FormValue("code"). Before
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:201
// calling Exchange, be sure to validate FormValue("state").
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:201
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:201
// Opts may include the PKCE verifier code if previously used in AuthCodeURL.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:201
// See https://www.oauth.com/oauth2-servers/pkce/ for more info.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:213
func (c *Config) Exchange(ctx context.Context, code string, opts ...AuthCodeOption) (*Token, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:213
	_go_fuzz_dep_.CoverTab[184181]++
										v := url.Values{
		"grant_type":	{"authorization_code"},
		"code":		{code},
	}
	if c.RedirectURL != "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:218
		_go_fuzz_dep_.CoverTab[184184]++
											v.Set("redirect_uri", c.RedirectURL)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:219
		// _ = "end of CoverTab[184184]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:220
		_go_fuzz_dep_.CoverTab[184185]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:220
		// _ = "end of CoverTab[184185]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:220
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:220
	// _ = "end of CoverTab[184181]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:220
	_go_fuzz_dep_.CoverTab[184182]++
										for _, opt := range opts {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:221
		_go_fuzz_dep_.CoverTab[184186]++
											opt.setValue(v)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:222
		// _ = "end of CoverTab[184186]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:223
	// _ = "end of CoverTab[184182]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:223
	_go_fuzz_dep_.CoverTab[184183]++
										return retrieveToken(ctx, c, v)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:224
	// _ = "end of CoverTab[184183]"
}

// Client returns an HTTP client using the provided token.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:227
// The token will auto-refresh as necessary. The underlying
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:227
// HTTP transport will be obtained using the provided context.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:227
// The returned client and its Transport should not be modified.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:231
func (c *Config) Client(ctx context.Context, t *Token) *http.Client {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:231
	_go_fuzz_dep_.CoverTab[184187]++
										return NewClient(ctx, c.TokenSource(ctx, t))
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:232
	// _ = "end of CoverTab[184187]"
}

// TokenSource returns a TokenSource that returns t until t expires,
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:235
// automatically refreshing it as necessary using the provided context.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:235
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:235
// Most users will use Config.Client instead.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:239
func (c *Config) TokenSource(ctx context.Context, t *Token) TokenSource {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:239
	_go_fuzz_dep_.CoverTab[184188]++
										tkr := &tokenRefresher{
		ctx:	ctx,
		conf:	c,
	}
	if t != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:244
		_go_fuzz_dep_.CoverTab[184190]++
											tkr.refreshToken = t.RefreshToken
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:245
		// _ = "end of CoverTab[184190]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:246
		_go_fuzz_dep_.CoverTab[184191]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:246
		// _ = "end of CoverTab[184191]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:246
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:246
	// _ = "end of CoverTab[184188]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:246
	_go_fuzz_dep_.CoverTab[184189]++
										return &reuseTokenSource{
		t:	t,
		new:	tkr,
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:250
	// _ = "end of CoverTab[184189]"
}

// tokenRefresher is a TokenSource that makes "grant_type"=="refresh_token"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:253
// HTTP requests to renew a token using a RefreshToken.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:255
type tokenRefresher struct {
	ctx		context.Context	// used to get HTTP requests
	conf		*Config
	refreshToken	string
}

// WARNING: Token is not safe for concurrent access, as it
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:261
// updates the tokenRefresher's refreshToken field.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:261
// Within this package, it is used by reuseTokenSource which
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:261
// synchronizes calls to this method with its own mutex.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:265
func (tf *tokenRefresher) Token() (*Token, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:265
	_go_fuzz_dep_.CoverTab[184192]++
										if tf.refreshToken == "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:266
		_go_fuzz_dep_.CoverTab[184196]++
											return nil, errors.New("oauth2: token expired and refresh token is not set")
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:267
		// _ = "end of CoverTab[184196]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:268
		_go_fuzz_dep_.CoverTab[184197]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:268
		// _ = "end of CoverTab[184197]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:268
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:268
	// _ = "end of CoverTab[184192]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:268
	_go_fuzz_dep_.CoverTab[184193]++

										tk, err := retrieveToken(tf.ctx, tf.conf, url.Values{
		"grant_type":		{"refresh_token"},
		"refresh_token":	{tf.refreshToken},
	})

	if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:275
		_go_fuzz_dep_.CoverTab[184198]++
											return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:276
		// _ = "end of CoverTab[184198]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:277
		_go_fuzz_dep_.CoverTab[184199]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:277
		// _ = "end of CoverTab[184199]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:277
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:277
	// _ = "end of CoverTab[184193]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:277
	_go_fuzz_dep_.CoverTab[184194]++
										if tf.refreshToken != tk.RefreshToken {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:278
		_go_fuzz_dep_.CoverTab[184200]++
											tf.refreshToken = tk.RefreshToken
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:279
		// _ = "end of CoverTab[184200]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:280
		_go_fuzz_dep_.CoverTab[184201]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:280
		// _ = "end of CoverTab[184201]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:280
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:280
	// _ = "end of CoverTab[184194]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:280
	_go_fuzz_dep_.CoverTab[184195]++
										return tk, err
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:281
	// _ = "end of CoverTab[184195]"
}

// reuseTokenSource is a TokenSource that holds a single token in memory
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:284
// and validates its expiry before each call to retrieve it with
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:284
// Token. If it's expired, it will be auto-refreshed using the
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:284
// new TokenSource.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:288
type reuseTokenSource struct {
	new	TokenSource	// called when t is expired.

	mu	sync.Mutex	// guards t
	t	*Token
}

// Token returns the current token if it's still valid, else will
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:295
// refresh the current token (using r.Context for HTTP client
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:295
// information) and return the new one.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:298
func (s *reuseTokenSource) Token() (*Token, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:298
	_go_fuzz_dep_.CoverTab[184202]++
										s.mu.Lock()
										defer s.mu.Unlock()
										if s.t.Valid() {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:301
		_go_fuzz_dep_.CoverTab[184205]++
											return s.t, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:302
		// _ = "end of CoverTab[184205]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:303
		_go_fuzz_dep_.CoverTab[184206]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:303
		// _ = "end of CoverTab[184206]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:303
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:303
	// _ = "end of CoverTab[184202]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:303
	_go_fuzz_dep_.CoverTab[184203]++
										t, err := s.new.Token()
										if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:305
		_go_fuzz_dep_.CoverTab[184207]++
											return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:306
		// _ = "end of CoverTab[184207]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:307
		_go_fuzz_dep_.CoverTab[184208]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:307
		// _ = "end of CoverTab[184208]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:307
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:307
	// _ = "end of CoverTab[184203]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:307
	_go_fuzz_dep_.CoverTab[184204]++
										s.t = t
										return t, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:309
	// _ = "end of CoverTab[184204]"
}

// StaticTokenSource returns a TokenSource that always returns the same token.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:312
// Because the provided token t is never refreshed, StaticTokenSource is only
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:312
// useful for tokens that never expire.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:315
func StaticTokenSource(t *Token) TokenSource {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:315
	_go_fuzz_dep_.CoverTab[184209]++
										return staticTokenSource{t}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:316
	// _ = "end of CoverTab[184209]"
}

// staticTokenSource is a TokenSource that always returns the same Token.
type staticTokenSource struct {
	t *Token
}

func (s staticTokenSource) Token() (*Token, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:324
	_go_fuzz_dep_.CoverTab[184210]++
										return s.t, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:325
	// _ = "end of CoverTab[184210]"
}

// HTTPClient is the context key to use with golang.org/x/net/context's
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:328
// WithValue function to associate an *http.Client value with a context.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:330
var HTTPClient internal.ContextKey

// NewClient creates an *http.Client from a Context and TokenSource.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:332
// The returned client is not valid beyond the lifetime of the context.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:332
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:332
// Note that if a custom *http.Client is provided via the Context it
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:332
// is used only for token acquisition and is not used to configure the
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:332
// *http.Client returned from NewClient.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:332
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:332
// As a special case, if src is nil, a non-OAuth2 client is returned
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:332
// using the provided context. This exists to support related OAuth2
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:332
// packages.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:342
func NewClient(ctx context.Context, src TokenSource) *http.Client {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:342
	_go_fuzz_dep_.CoverTab[184211]++
										if src == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:343
		_go_fuzz_dep_.CoverTab[184213]++
											return internal.ContextClient(ctx)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:344
		// _ = "end of CoverTab[184213]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:345
		_go_fuzz_dep_.CoverTab[184214]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:345
		// _ = "end of CoverTab[184214]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:345
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:345
	// _ = "end of CoverTab[184211]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:345
	_go_fuzz_dep_.CoverTab[184212]++
										return &http.Client{
		Transport: &Transport{
			Base:	internal.ContextClient(ctx).Transport,
			Source:	ReuseTokenSource(nil, src),
		},
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:351
	// _ = "end of CoverTab[184212]"
}

// ReuseTokenSource returns a TokenSource which repeatedly returns the
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:354
// same token as long as it's valid, starting with t.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:354
// When its cached token is invalid, a new token is obtained from src.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:354
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:354
// ReuseTokenSource is typically used to reuse tokens from a cache
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:354
// (such as a file on disk) between runs of a program, rather than
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:354
// obtaining new tokens unnecessarily.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:354
//
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:354
// The initial token t may be nil, in which case the TokenSource is
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:354
// wrapped in a caching version if it isn't one already. This also
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:354
// means it's always safe to wrap ReuseTokenSource around any other
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:354
// TokenSource without adverse effects.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:366
func ReuseTokenSource(t *Token, src TokenSource) TokenSource {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:366
	_go_fuzz_dep_.CoverTab[184215]++

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:370
	if rt, ok := src.(*reuseTokenSource); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:370
		_go_fuzz_dep_.CoverTab[184217]++
											if t == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:371
			_go_fuzz_dep_.CoverTab[184219]++

												return rt
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:373
			// _ = "end of CoverTab[184219]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:374
			_go_fuzz_dep_.CoverTab[184220]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:374
			// _ = "end of CoverTab[184220]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:374
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:374
		// _ = "end of CoverTab[184217]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:374
		_go_fuzz_dep_.CoverTab[184218]++
											src = rt.new
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:375
		// _ = "end of CoverTab[184218]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:376
		_go_fuzz_dep_.CoverTab[184221]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:376
		// _ = "end of CoverTab[184221]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:376
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:376
	// _ = "end of CoverTab[184215]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:376
	_go_fuzz_dep_.CoverTab[184216]++
										return &reuseTokenSource{
		t:	t,
		new:	src,
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:380
	// _ = "end of CoverTab[184216]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:381
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/oauth2.go:381
var _ = _go_fuzz_dep_.CoverTab
