//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:1
package oidc

//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:1
)

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var (
	// ErrTokenExpired indicates that a token parsed by a verifier has expired.
	ErrTokenExpired	= errors.New("oidc: ID Token expired")
	// ErrNotSupported indicates that the requested optional OpenID Connect endpoint is not supported by the provider.
	ErrNotSupported	= errors.New("oidc: endpoint not supported")
)

const (
	// ScopeOpenID is the mandatory scope for all OpenID Connect OAuth2 requests.
	ScopeOpenID	= "openid"

	// ScopeOfflineAccess is an optional scope defined by OpenID Connect for requesting
	// OAuth2 refresh tokens.
	//
	// Support for this scope differs between OpenID Connect providers. For instance
	// Google rejects it, favoring appending "access_type=offline" as part of the
	// authorization request instead.
	//
	// See: https://openid.net/specs/openid-connect-core-1_0.html#OfflineAccess
	ScopeOfflineAccess	= "offline_access"
)

// Provider contains the subset of the OpenID Connect provider metadata needed to request
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:38
// and verify ID Tokens.
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:40
type Provider struct {
	Issuer		string	`json:"issuer"`
	AuthURL		string	`json:"authorization_endpoint"`
	TokenURL	string	`json:"token_endpoint"`
	JWKSURL		string	`json:"jwks_uri"`
	UserInfoURL	string	`json:"userinfo_endpoint"`

	// Raw claims returned by the server.
	rawClaims	[]byte
}

// NewProvider uses the OpenID Connect disovery mechanism to construct a Provider.
func NewProvider(ctx context.Context, issuer string) (*Provider, error) {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:52
	_go_fuzz_dep_.CoverTab[186842]++
														wellKnown := strings.TrimSuffix(issuer, "/") + "/.well-known/openid-configuration"
														resp, err := contextClient(ctx).Get(wellKnown)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:55
		_go_fuzz_dep_.CoverTab[186848]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:56
		// _ = "end of CoverTab[186848]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:57
		_go_fuzz_dep_.CoverTab[186849]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:57
		// _ = "end of CoverTab[186849]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:57
	// _ = "end of CoverTab[186842]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:57
	_go_fuzz_dep_.CoverTab[186843]++
														body, err := ioutil.ReadAll(resp.Body)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:59
		_go_fuzz_dep_.CoverTab[186850]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:60
		// _ = "end of CoverTab[186850]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:61
		_go_fuzz_dep_.CoverTab[186851]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:61
		// _ = "end of CoverTab[186851]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:61
	// _ = "end of CoverTab[186843]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:61
	_go_fuzz_dep_.CoverTab[186844]++
														if resp.StatusCode != http.StatusOK {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:62
		_go_fuzz_dep_.CoverTab[186852]++
															return nil, fmt.Errorf("%s: %s", resp.Status, body)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:63
		// _ = "end of CoverTab[186852]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:64
		_go_fuzz_dep_.CoverTab[186853]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:64
		// _ = "end of CoverTab[186853]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:64
	// _ = "end of CoverTab[186844]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:64
	_go_fuzz_dep_.CoverTab[186845]++
														defer resp.Body.Close()
														var p Provider
														if err := json.Unmarshal(body, &p); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:67
		_go_fuzz_dep_.CoverTab[186854]++
															return nil, fmt.Errorf("oidc: failed to decode provider discovery object: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:68
		// _ = "end of CoverTab[186854]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:69
		_go_fuzz_dep_.CoverTab[186855]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:69
		// _ = "end of CoverTab[186855]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:69
	// _ = "end of CoverTab[186845]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:69
	_go_fuzz_dep_.CoverTab[186846]++
														p.rawClaims = body
														if p.Issuer != issuer {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:71
		_go_fuzz_dep_.CoverTab[186856]++
															return nil, fmt.Errorf("oidc: issuer did not match the issuer returned by provider, expected %q got %q", issuer, p.Issuer)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:72
		// _ = "end of CoverTab[186856]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:73
		_go_fuzz_dep_.CoverTab[186857]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:73
		// _ = "end of CoverTab[186857]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:73
	// _ = "end of CoverTab[186846]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:73
	_go_fuzz_dep_.CoverTab[186847]++
														return &p, nil
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:74
	// _ = "end of CoverTab[186847]"
}

// Claims returns additional fields returned by the server during discovery.
func (p *Provider) Claims(v interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:78
	_go_fuzz_dep_.CoverTab[186858]++
														if p.rawClaims == nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:79
		_go_fuzz_dep_.CoverTab[186860]++
															return errors.New("oidc: claims not set")
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:80
		// _ = "end of CoverTab[186860]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:81
		_go_fuzz_dep_.CoverTab[186861]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:81
		// _ = "end of CoverTab[186861]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:81
	// _ = "end of CoverTab[186858]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:81
	_go_fuzz_dep_.CoverTab[186859]++
														return json.Unmarshal(p.rawClaims, v)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:82
	// _ = "end of CoverTab[186859]"
}

// Endpoint returns the OAuth2 auth and token endpoints for the given provider.
func (p *Provider) Endpoint() oauth2.Endpoint {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:86
	_go_fuzz_dep_.CoverTab[186862]++
														return oauth2.Endpoint{AuthURL: p.AuthURL, TokenURL: p.TokenURL}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:87
	// _ = "end of CoverTab[186862]"
}

// UserInfo represents the OpenID Connect userinfo claims.
type UserInfo struct {
	Subject		string	`json:"sub"`
	Profile		string	`json:"profile"`
	Email		string	`json:"email"`
	EmailVerified	bool	`json:"email_verified"`

	claims	[]byte
}

// Claims unmarshals the raw JSON object claims into the provided object.
func (u *UserInfo) Claims(v interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:101
	_go_fuzz_dep_.CoverTab[186863]++
														if u.claims == nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:102
		_go_fuzz_dep_.CoverTab[186865]++
															return errors.New("oidc: claims not set")
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:103
		// _ = "end of CoverTab[186865]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:104
		_go_fuzz_dep_.CoverTab[186866]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:104
		// _ = "end of CoverTab[186866]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:104
	// _ = "end of CoverTab[186863]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:104
	_go_fuzz_dep_.CoverTab[186864]++
														return json.Unmarshal(u.claims, v)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:105
	// _ = "end of CoverTab[186864]"
}

// UserInfo uses the token source to query the provider's user info endpoint.
func (p *Provider) UserInfo(ctx context.Context, tokenSource oauth2.TokenSource) (*UserInfo, error) {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:109
	_go_fuzz_dep_.CoverTab[186867]++
														if p.UserInfoURL == "" {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:110
		_go_fuzz_dep_.CoverTab[186873]++
															return nil, ErrNotSupported
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:111
		// _ = "end of CoverTab[186873]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:112
		_go_fuzz_dep_.CoverTab[186874]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:112
		// _ = "end of CoverTab[186874]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:112
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:112
	// _ = "end of CoverTab[186867]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:112
	_go_fuzz_dep_.CoverTab[186868]++
														cli := oauth2.NewClient(ctx, tokenSource)
														resp, err := cli.Get(p.UserInfoURL)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:115
		_go_fuzz_dep_.CoverTab[186875]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:116
		// _ = "end of CoverTab[186875]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:117
		_go_fuzz_dep_.CoverTab[186876]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:117
		// _ = "end of CoverTab[186876]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:117
	// _ = "end of CoverTab[186868]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:117
	_go_fuzz_dep_.CoverTab[186869]++
														defer resp.Body.Close()
														body, err := ioutil.ReadAll(resp.Body)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:120
		_go_fuzz_dep_.CoverTab[186877]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:121
		// _ = "end of CoverTab[186877]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:122
		_go_fuzz_dep_.CoverTab[186878]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:122
		// _ = "end of CoverTab[186878]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:122
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:122
	// _ = "end of CoverTab[186869]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:122
	_go_fuzz_dep_.CoverTab[186870]++
														if resp.StatusCode != http.StatusOK {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:123
		_go_fuzz_dep_.CoverTab[186879]++
															return nil, fmt.Errorf("%s: %s", resp.Status, body)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:124
		// _ = "end of CoverTab[186879]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:125
		_go_fuzz_dep_.CoverTab[186880]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:125
		// _ = "end of CoverTab[186880]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:125
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:125
	// _ = "end of CoverTab[186870]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:125
	_go_fuzz_dep_.CoverTab[186871]++

														var userInfo UserInfo
														if err := json.Unmarshal(body, &userInfo); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:128
		_go_fuzz_dep_.CoverTab[186881]++
															return nil, fmt.Errorf("oidc: failed to decode userinfo: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:129
		// _ = "end of CoverTab[186881]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:130
		_go_fuzz_dep_.CoverTab[186882]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:130
		// _ = "end of CoverTab[186882]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:130
	// _ = "end of CoverTab[186871]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:130
	_go_fuzz_dep_.CoverTab[186872]++
														userInfo.claims = body
														return &userInfo, nil
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:132
	// _ = "end of CoverTab[186872]"
}

// IDToken is an OpenID Connect extension that provides a predictable representation
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
// of an authorization event.
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
//
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
// The ID Token only holds fields OpenID Connect requires. To access additional
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
// claims returned by the server, use the Claims method.
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
//
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
//	idToken, err := idTokenVerifier.Verify(rawIDToken)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
//	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
//		// handle error
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
//	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
//	var claims struct {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
//		Email         string `json:"email"`
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
//		EmailVerified bool   `json:"email_verified"`
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
//	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
//	if err := idToken.Claims(&claims); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
//		// handle error
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:135
//	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:153
type IDToken struct {
	// The URL of the server which issued this token. This will always be the same
	// as the URL used for initial discovery.
	Issuer	string

	// The client, or set of clients, that this token is issued for.
	Audience	[]string

	// A unique string which identifies the end user.
	Subject	string

	IssuedAt	time.Time
	Expiry		time.Time
	Nonce		string

	claims	[]byte
}

// Claims unmarshals the raw JSON payload of the ID Token into a provided struct.
func (i *IDToken) Claims(v interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:172
	_go_fuzz_dep_.CoverTab[186883]++
														if i.claims == nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:173
		_go_fuzz_dep_.CoverTab[186885]++
															return errors.New("oidc: claims not set")
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:174
		// _ = "end of CoverTab[186885]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:175
		_go_fuzz_dep_.CoverTab[186886]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:175
		// _ = "end of CoverTab[186886]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:175
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:175
	// _ = "end of CoverTab[186883]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:175
	_go_fuzz_dep_.CoverTab[186884]++
														return json.Unmarshal(i.claims, v)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:176
	// _ = "end of CoverTab[186884]"
}

type audience []string

func (a *audience) UnmarshalJSON(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:181
	_go_fuzz_dep_.CoverTab[186887]++
														var s string
														if json.Unmarshal(b, &s) == nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:183
		_go_fuzz_dep_.CoverTab[186890]++
															*a = audience{s}
															return nil
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:185
		// _ = "end of CoverTab[186890]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:186
		_go_fuzz_dep_.CoverTab[186891]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:186
		// _ = "end of CoverTab[186891]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:186
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:186
	// _ = "end of CoverTab[186887]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:186
	_go_fuzz_dep_.CoverTab[186888]++
														var auds []string
														if err := json.Unmarshal(b, &auds); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:188
		_go_fuzz_dep_.CoverTab[186892]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:189
		// _ = "end of CoverTab[186892]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:190
		_go_fuzz_dep_.CoverTab[186893]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:190
		// _ = "end of CoverTab[186893]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:190
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:190
	// _ = "end of CoverTab[186888]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:190
	_go_fuzz_dep_.CoverTab[186889]++
														*a = audience(auds)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:192
	// _ = "end of CoverTab[186889]"
}

type jsonTime time.Time

func (j *jsonTime) UnmarshalJSON(b []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:197
	_go_fuzz_dep_.CoverTab[186894]++
														var n json.Number
														if err := json.Unmarshal(b, &n); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:199
		_go_fuzz_dep_.CoverTab[186897]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:200
		// _ = "end of CoverTab[186897]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:201
		_go_fuzz_dep_.CoverTab[186898]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:201
		// _ = "end of CoverTab[186898]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:201
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:201
	// _ = "end of CoverTab[186894]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:201
	_go_fuzz_dep_.CoverTab[186895]++
														var unix int64

														if t, err := n.Int64(); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:204
		_go_fuzz_dep_.CoverTab[186899]++
															unix = t
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:205
		// _ = "end of CoverTab[186899]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:206
		_go_fuzz_dep_.CoverTab[186900]++
															f, err := n.Float64()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:208
			_go_fuzz_dep_.CoverTab[186902]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:209
			// _ = "end of CoverTab[186902]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:210
			_go_fuzz_dep_.CoverTab[186903]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:210
			// _ = "end of CoverTab[186903]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:210
		}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:210
		// _ = "end of CoverTab[186900]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:210
		_go_fuzz_dep_.CoverTab[186901]++
															unix = int64(f)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:211
		// _ = "end of CoverTab[186901]"
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:212
	// _ = "end of CoverTab[186895]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:212
	_go_fuzz_dep_.CoverTab[186896]++
														*j = jsonTime(time.Unix(unix, 0))
														return nil
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:214
	// _ = "end of CoverTab[186896]"
}

type idToken struct {
	Issuer		string		`json:"iss"`
	Subject		string		`json:"sub"`
	Audience	audience	`json:"aud"`
	Expiry		jsonTime	`json:"exp"`
	IssuedAt	jsonTime	`json:"iat"`
	Nonce		string		`json:"nonce"`
}

// IDTokenVerifier provides verification for ID Tokens.
type IDTokenVerifier struct {
	issuer	string
	keySet	*remoteKeySet
	options	[]VerificationOption
}

// Verify parse the raw ID Token, verifies it's been signed by the provider, preforms
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:233
// additional verification, and returns the claims.
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:235
func (v *IDTokenVerifier) Verify(rawIDToken string) (*IDToken, error) {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:235
	_go_fuzz_dep_.CoverTab[186904]++
														payload, err := v.keySet.verifyJWT(rawIDToken)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:237
		_go_fuzz_dep_.CoverTab[186909]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:238
		// _ = "end of CoverTab[186909]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:239
		_go_fuzz_dep_.CoverTab[186910]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:239
		// _ = "end of CoverTab[186910]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:239
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:239
	// _ = "end of CoverTab[186904]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:239
	_go_fuzz_dep_.CoverTab[186905]++
														var token idToken
														if err := json.Unmarshal(payload, &token); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:241
		_go_fuzz_dep_.CoverTab[186911]++
															return nil, fmt.Errorf("oidc: failed to unmarshal claims: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:242
		// _ = "end of CoverTab[186911]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:243
		_go_fuzz_dep_.CoverTab[186912]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:243
		// _ = "end of CoverTab[186912]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:243
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:243
	// _ = "end of CoverTab[186905]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:243
	_go_fuzz_dep_.CoverTab[186906]++
														if v.issuer != token.Issuer {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:244
		_go_fuzz_dep_.CoverTab[186913]++
															return nil, fmt.Errorf("oidc: iss field did not match provider issuer")
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:245
		// _ = "end of CoverTab[186913]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:246
		_go_fuzz_dep_.CoverTab[186914]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:246
		// _ = "end of CoverTab[186914]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:246
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:246
	// _ = "end of CoverTab[186906]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:246
	_go_fuzz_dep_.CoverTab[186907]++
														t := &IDToken{
		Issuer:		token.Issuer,
		Subject:	token.Subject,
		Audience:	[]string(token.Audience),
		Expiry:		time.Time(token.Expiry),
		IssuedAt:	time.Time(token.IssuedAt),
		Nonce:		token.Nonce,
		claims:		payload,
	}
	for _, option := range v.options {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:256
		_go_fuzz_dep_.CoverTab[186915]++
															if err := option.verifyIDToken(t); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:257
			_go_fuzz_dep_.CoverTab[186916]++
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:258
			// _ = "end of CoverTab[186916]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:259
			_go_fuzz_dep_.CoverTab[186917]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:259
			// _ = "end of CoverTab[186917]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:259
		}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:259
		// _ = "end of CoverTab[186915]"
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:260
	// _ = "end of CoverTab[186907]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:260
	_go_fuzz_dep_.CoverTab[186908]++
														return t, nil
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:261
	// _ = "end of CoverTab[186908]"
}

// NewVerifier returns an IDTokenVerifier that uses the provider's key set to verify JWTs.
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:264
//
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:264
// The verifier queries the provider to update keys when a signature cannot be verified by the
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:264
// set of keys cached from the previous request.
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:268
func (p *Provider) NewVerifier(ctx context.Context, options ...VerificationOption) *IDTokenVerifier {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:268
	_go_fuzz_dep_.CoverTab[186918]++
														return &IDTokenVerifier{
		issuer:		p.Issuer,
		keySet:		newRemoteKeySet(ctx, p.JWKSURL),
		options:	options,
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:273
	// _ = "end of CoverTab[186918]"
}

// VerificationOption is an option provided to Provider.NewVerifier.
type VerificationOption interface {
	verifyIDToken(token *IDToken) error
}

// VerifyAudience ensures that an ID Token was issued for the specific client.
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:281
//
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:281
// Note that a verified token may be valid for other clients, as OpenID Connect allows a token to have
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:281
// multiple audiences.
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:285
func VerifyAudience(clientID string) VerificationOption {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:285
	_go_fuzz_dep_.CoverTab[186919]++
														return clientVerifier{clientID}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:286
	// _ = "end of CoverTab[186919]"
}

type clientVerifier struct {
	clientID string
}

func (c clientVerifier) verifyIDToken(token *IDToken) error {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:293
	_go_fuzz_dep_.CoverTab[186920]++
														for _, aud := range token.Audience {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:294
		_go_fuzz_dep_.CoverTab[186922]++
															if aud == c.clientID {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:295
			_go_fuzz_dep_.CoverTab[186923]++
																return nil
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:296
			// _ = "end of CoverTab[186923]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:297
			_go_fuzz_dep_.CoverTab[186924]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:297
			// _ = "end of CoverTab[186924]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:297
		}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:297
		// _ = "end of CoverTab[186922]"
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:298
	// _ = "end of CoverTab[186920]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:298
	_go_fuzz_dep_.CoverTab[186921]++
														return errors.New("oidc: id token aud field did not match client_id")
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:299
	// _ = "end of CoverTab[186921]"
}

// VerifyExpiry ensures that an ID Token has not expired.
func VerifyExpiry() VerificationOption {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:303
	_go_fuzz_dep_.CoverTab[186925]++
														return expiryVerifier{time.Now}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:304
	// _ = "end of CoverTab[186925]"
}

type expiryVerifier struct {
	now func() time.Time
}

func (e expiryVerifier) verifyIDToken(token *IDToken) error {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:311
	_go_fuzz_dep_.CoverTab[186926]++
														if e.now().After(token.Expiry) {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:312
		_go_fuzz_dep_.CoverTab[186928]++
															return ErrTokenExpired
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:313
		// _ = "end of CoverTab[186928]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:314
		_go_fuzz_dep_.CoverTab[186929]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:314
		// _ = "end of CoverTab[186929]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:314
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:314
	// _ = "end of CoverTab[186926]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:314
	_go_fuzz_dep_.CoverTab[186927]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:315
	// _ = "end of CoverTab[186927]"
}

// This method is internal to golang.org/x/oauth2. Just copy it.
func contextClient(ctx context.Context) *http.Client {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:319
	_go_fuzz_dep_.CoverTab[186930]++
														if ctx != nil {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:320
		_go_fuzz_dep_.CoverTab[186932]++
															if hc, ok := ctx.Value(oauth2.HTTPClient).(*http.Client); ok {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:321
			_go_fuzz_dep_.CoverTab[186933]++
																return hc
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:322
			// _ = "end of CoverTab[186933]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:323
			_go_fuzz_dep_.CoverTab[186934]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:323
			// _ = "end of CoverTab[186934]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:323
		}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:323
		// _ = "end of CoverTab[186932]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:324
		_go_fuzz_dep_.CoverTab[186935]++
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:324
		// _ = "end of CoverTab[186935]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:324
	}
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:324
	// _ = "end of CoverTab[186930]"
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:324
	_go_fuzz_dep_.CoverTab[186931]++
														return http.DefaultClient
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:325
	// _ = "end of CoverTab[186931]"
}

//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:326
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/ericchiang/oidc@v0.0.0-20160908143337-11f62933e071/oidc.go:326
var _ = _go_fuzz_dep_.CoverTab
