// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/http/cookie.go:5
package http

//line /usr/local/go/src/net/http/cookie.go:5
import (
//line /usr/local/go/src/net/http/cookie.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/cookie.go:5
)
//line /usr/local/go/src/net/http/cookie.go:5
import (
//line /usr/local/go/src/net/http/cookie.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/cookie.go:5
)

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http/internal/ascii"
	"net/textproto"
	"strconv"
	"strings"
	"time"
)

// A Cookie represents an HTTP cookie as sent in the Set-Cookie header of an
//line /usr/local/go/src/net/http/cookie.go:19
// HTTP response or the Cookie header of an HTTP request.
//line /usr/local/go/src/net/http/cookie.go:19
//
//line /usr/local/go/src/net/http/cookie.go:19
// See https://tools.ietf.org/html/rfc6265 for details.
//line /usr/local/go/src/net/http/cookie.go:23
type Cookie struct {
	Name	string
	Value	string

	Path		string		// optional
	Domain		string		// optional
	Expires		time.Time	// optional
	RawExpires	string		// for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge		int
	Secure		bool
	HttpOnly	bool
	SameSite	SameSite
	Raw		string
	Unparsed	[]string	// Raw text of unparsed attribute-value pairs
}

// SameSite allows a server to define a cookie attribute making it impossible for
//line /usr/local/go/src/net/http/cookie.go:43
// the browser to send this cookie along with cross-site requests. The main
//line /usr/local/go/src/net/http/cookie.go:43
// goal is to mitigate the risk of cross-origin information leakage, and provide
//line /usr/local/go/src/net/http/cookie.go:43
// some protection against cross-site request forgery attacks.
//line /usr/local/go/src/net/http/cookie.go:43
//
//line /usr/local/go/src/net/http/cookie.go:43
// See https://tools.ietf.org/html/draft-ietf-httpbis-cookie-same-site-00 for details.
//line /usr/local/go/src/net/http/cookie.go:49
type SameSite int

const (
	SameSiteDefaultMode	SameSite	= iota + 1
	SameSiteLaxMode
	SameSiteStrictMode
	SameSiteNoneMode
)

// readSetCookies parses all "Set-Cookie" values from
//line /usr/local/go/src/net/http/cookie.go:58
// the header h and returns the successfully parsed Cookies.
//line /usr/local/go/src/net/http/cookie.go:60
func readSetCookies(h Header) []*Cookie {
//line /usr/local/go/src/net/http/cookie.go:60
	_go_fuzz_dep_.CoverTab[36912]++
						cookieCount := len(h["Set-Cookie"])
						if cookieCount == 0 {
//line /usr/local/go/src/net/http/cookie.go:62
		_go_fuzz_dep_.CoverTab[36915]++
							return []*Cookie{}
//line /usr/local/go/src/net/http/cookie.go:63
		// _ = "end of CoverTab[36915]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:64
		_go_fuzz_dep_.CoverTab[36916]++
//line /usr/local/go/src/net/http/cookie.go:64
		// _ = "end of CoverTab[36916]"
//line /usr/local/go/src/net/http/cookie.go:64
	}
//line /usr/local/go/src/net/http/cookie.go:64
	// _ = "end of CoverTab[36912]"
//line /usr/local/go/src/net/http/cookie.go:64
	_go_fuzz_dep_.CoverTab[36913]++
						cookies := make([]*Cookie, 0, cookieCount)
						for _, line := range h["Set-Cookie"] {
//line /usr/local/go/src/net/http/cookie.go:66
		_go_fuzz_dep_.CoverTab[36917]++
							parts := strings.Split(textproto.TrimString(line), ";")
							if len(parts) == 1 && func() bool {
//line /usr/local/go/src/net/http/cookie.go:68
			_go_fuzz_dep_.CoverTab[36923]++
//line /usr/local/go/src/net/http/cookie.go:68
			return parts[0] == ""
//line /usr/local/go/src/net/http/cookie.go:68
			// _ = "end of CoverTab[36923]"
//line /usr/local/go/src/net/http/cookie.go:68
		}() {
//line /usr/local/go/src/net/http/cookie.go:68
			_go_fuzz_dep_.CoverTab[36924]++
								continue
//line /usr/local/go/src/net/http/cookie.go:69
			// _ = "end of CoverTab[36924]"
		} else {
//line /usr/local/go/src/net/http/cookie.go:70
			_go_fuzz_dep_.CoverTab[36925]++
//line /usr/local/go/src/net/http/cookie.go:70
			// _ = "end of CoverTab[36925]"
//line /usr/local/go/src/net/http/cookie.go:70
		}
//line /usr/local/go/src/net/http/cookie.go:70
		// _ = "end of CoverTab[36917]"
//line /usr/local/go/src/net/http/cookie.go:70
		_go_fuzz_dep_.CoverTab[36918]++
							parts[0] = textproto.TrimString(parts[0])
							name, value, ok := strings.Cut(parts[0], "=")
							if !ok {
//line /usr/local/go/src/net/http/cookie.go:73
			_go_fuzz_dep_.CoverTab[36926]++
								continue
//line /usr/local/go/src/net/http/cookie.go:74
			// _ = "end of CoverTab[36926]"
		} else {
//line /usr/local/go/src/net/http/cookie.go:75
			_go_fuzz_dep_.CoverTab[36927]++
//line /usr/local/go/src/net/http/cookie.go:75
			// _ = "end of CoverTab[36927]"
//line /usr/local/go/src/net/http/cookie.go:75
		}
//line /usr/local/go/src/net/http/cookie.go:75
		// _ = "end of CoverTab[36918]"
//line /usr/local/go/src/net/http/cookie.go:75
		_go_fuzz_dep_.CoverTab[36919]++
							name = textproto.TrimString(name)
							if !isCookieNameValid(name) {
//line /usr/local/go/src/net/http/cookie.go:77
			_go_fuzz_dep_.CoverTab[36928]++
								continue
//line /usr/local/go/src/net/http/cookie.go:78
			// _ = "end of CoverTab[36928]"
		} else {
//line /usr/local/go/src/net/http/cookie.go:79
			_go_fuzz_dep_.CoverTab[36929]++
//line /usr/local/go/src/net/http/cookie.go:79
			// _ = "end of CoverTab[36929]"
//line /usr/local/go/src/net/http/cookie.go:79
		}
//line /usr/local/go/src/net/http/cookie.go:79
		// _ = "end of CoverTab[36919]"
//line /usr/local/go/src/net/http/cookie.go:79
		_go_fuzz_dep_.CoverTab[36920]++
							value, ok = parseCookieValue(value, true)
							if !ok {
//line /usr/local/go/src/net/http/cookie.go:81
			_go_fuzz_dep_.CoverTab[36930]++
								continue
//line /usr/local/go/src/net/http/cookie.go:82
			// _ = "end of CoverTab[36930]"
		} else {
//line /usr/local/go/src/net/http/cookie.go:83
			_go_fuzz_dep_.CoverTab[36931]++
//line /usr/local/go/src/net/http/cookie.go:83
			// _ = "end of CoverTab[36931]"
//line /usr/local/go/src/net/http/cookie.go:83
		}
//line /usr/local/go/src/net/http/cookie.go:83
		// _ = "end of CoverTab[36920]"
//line /usr/local/go/src/net/http/cookie.go:83
		_go_fuzz_dep_.CoverTab[36921]++
							c := &Cookie{
			Name:	name,
			Value:	value,
			Raw:	line,
		}
		for i := 1; i < len(parts); i++ {
//line /usr/local/go/src/net/http/cookie.go:89
			_go_fuzz_dep_.CoverTab[36932]++
								parts[i] = textproto.TrimString(parts[i])
								if len(parts[i]) == 0 {
//line /usr/local/go/src/net/http/cookie.go:91
				_go_fuzz_dep_.CoverTab[36937]++
									continue
//line /usr/local/go/src/net/http/cookie.go:92
				// _ = "end of CoverTab[36937]"
			} else {
//line /usr/local/go/src/net/http/cookie.go:93
				_go_fuzz_dep_.CoverTab[36938]++
//line /usr/local/go/src/net/http/cookie.go:93
				// _ = "end of CoverTab[36938]"
//line /usr/local/go/src/net/http/cookie.go:93
			}
//line /usr/local/go/src/net/http/cookie.go:93
			// _ = "end of CoverTab[36932]"
//line /usr/local/go/src/net/http/cookie.go:93
			_go_fuzz_dep_.CoverTab[36933]++

								attr, val, _ := strings.Cut(parts[i], "=")
								lowerAttr, isASCII := ascii.ToLower(attr)
								if !isASCII {
//line /usr/local/go/src/net/http/cookie.go:97
				_go_fuzz_dep_.CoverTab[36939]++
									continue
//line /usr/local/go/src/net/http/cookie.go:98
				// _ = "end of CoverTab[36939]"
			} else {
//line /usr/local/go/src/net/http/cookie.go:99
				_go_fuzz_dep_.CoverTab[36940]++
//line /usr/local/go/src/net/http/cookie.go:99
				// _ = "end of CoverTab[36940]"
//line /usr/local/go/src/net/http/cookie.go:99
			}
//line /usr/local/go/src/net/http/cookie.go:99
			// _ = "end of CoverTab[36933]"
//line /usr/local/go/src/net/http/cookie.go:99
			_go_fuzz_dep_.CoverTab[36934]++
								val, ok = parseCookieValue(val, false)
								if !ok {
//line /usr/local/go/src/net/http/cookie.go:101
				_go_fuzz_dep_.CoverTab[36941]++
										c.Unparsed = append(c.Unparsed, parts[i])
										continue
//line /usr/local/go/src/net/http/cookie.go:103
				// _ = "end of CoverTab[36941]"
			} else {
//line /usr/local/go/src/net/http/cookie.go:104
				_go_fuzz_dep_.CoverTab[36942]++
//line /usr/local/go/src/net/http/cookie.go:104
				// _ = "end of CoverTab[36942]"
//line /usr/local/go/src/net/http/cookie.go:104
			}
//line /usr/local/go/src/net/http/cookie.go:104
			// _ = "end of CoverTab[36934]"
//line /usr/local/go/src/net/http/cookie.go:104
			_go_fuzz_dep_.CoverTab[36935]++

									switch lowerAttr {
			case "samesite":
//line /usr/local/go/src/net/http/cookie.go:107
				_go_fuzz_dep_.CoverTab[36943]++
										lowerVal, ascii := ascii.ToLower(val)
										if !ascii {
//line /usr/local/go/src/net/http/cookie.go:109
					_go_fuzz_dep_.CoverTab[36956]++
											c.SameSite = SameSiteDefaultMode
											continue
//line /usr/local/go/src/net/http/cookie.go:111
					// _ = "end of CoverTab[36956]"
				} else {
//line /usr/local/go/src/net/http/cookie.go:112
					_go_fuzz_dep_.CoverTab[36957]++
//line /usr/local/go/src/net/http/cookie.go:112
					// _ = "end of CoverTab[36957]"
//line /usr/local/go/src/net/http/cookie.go:112
				}
//line /usr/local/go/src/net/http/cookie.go:112
				// _ = "end of CoverTab[36943]"
//line /usr/local/go/src/net/http/cookie.go:112
				_go_fuzz_dep_.CoverTab[36944]++
										switch lowerVal {
				case "lax":
//line /usr/local/go/src/net/http/cookie.go:114
					_go_fuzz_dep_.CoverTab[36958]++
											c.SameSite = SameSiteLaxMode
//line /usr/local/go/src/net/http/cookie.go:115
					// _ = "end of CoverTab[36958]"
				case "strict":
//line /usr/local/go/src/net/http/cookie.go:116
					_go_fuzz_dep_.CoverTab[36959]++
											c.SameSite = SameSiteStrictMode
//line /usr/local/go/src/net/http/cookie.go:117
					// _ = "end of CoverTab[36959]"
				case "none":
//line /usr/local/go/src/net/http/cookie.go:118
					_go_fuzz_dep_.CoverTab[36960]++
											c.SameSite = SameSiteNoneMode
//line /usr/local/go/src/net/http/cookie.go:119
					// _ = "end of CoverTab[36960]"
				default:
//line /usr/local/go/src/net/http/cookie.go:120
					_go_fuzz_dep_.CoverTab[36961]++
											c.SameSite = SameSiteDefaultMode
//line /usr/local/go/src/net/http/cookie.go:121
					// _ = "end of CoverTab[36961]"
				}
//line /usr/local/go/src/net/http/cookie.go:122
				// _ = "end of CoverTab[36944]"
//line /usr/local/go/src/net/http/cookie.go:122
				_go_fuzz_dep_.CoverTab[36945]++
										continue
//line /usr/local/go/src/net/http/cookie.go:123
				// _ = "end of CoverTab[36945]"
			case "secure":
//line /usr/local/go/src/net/http/cookie.go:124
				_go_fuzz_dep_.CoverTab[36946]++
										c.Secure = true
										continue
//line /usr/local/go/src/net/http/cookie.go:126
				// _ = "end of CoverTab[36946]"
			case "httponly":
//line /usr/local/go/src/net/http/cookie.go:127
				_go_fuzz_dep_.CoverTab[36947]++
										c.HttpOnly = true
										continue
//line /usr/local/go/src/net/http/cookie.go:129
				// _ = "end of CoverTab[36947]"
			case "domain":
//line /usr/local/go/src/net/http/cookie.go:130
				_go_fuzz_dep_.CoverTab[36948]++
										c.Domain = val
										continue
//line /usr/local/go/src/net/http/cookie.go:132
				// _ = "end of CoverTab[36948]"
			case "max-age":
//line /usr/local/go/src/net/http/cookie.go:133
				_go_fuzz_dep_.CoverTab[36949]++
										secs, err := strconv.Atoi(val)
										if err != nil || func() bool {
//line /usr/local/go/src/net/http/cookie.go:135
					_go_fuzz_dep_.CoverTab[36962]++
//line /usr/local/go/src/net/http/cookie.go:135
					return secs != 0 && func() bool {
//line /usr/local/go/src/net/http/cookie.go:135
						_go_fuzz_dep_.CoverTab[36963]++
//line /usr/local/go/src/net/http/cookie.go:135
						return val[0] == '0'
//line /usr/local/go/src/net/http/cookie.go:135
						// _ = "end of CoverTab[36963]"
//line /usr/local/go/src/net/http/cookie.go:135
					}()
//line /usr/local/go/src/net/http/cookie.go:135
					// _ = "end of CoverTab[36962]"
//line /usr/local/go/src/net/http/cookie.go:135
				}() {
//line /usr/local/go/src/net/http/cookie.go:135
					_go_fuzz_dep_.CoverTab[36964]++
											break
//line /usr/local/go/src/net/http/cookie.go:136
					// _ = "end of CoverTab[36964]"
				} else {
//line /usr/local/go/src/net/http/cookie.go:137
					_go_fuzz_dep_.CoverTab[36965]++
//line /usr/local/go/src/net/http/cookie.go:137
					// _ = "end of CoverTab[36965]"
//line /usr/local/go/src/net/http/cookie.go:137
				}
//line /usr/local/go/src/net/http/cookie.go:137
				// _ = "end of CoverTab[36949]"
//line /usr/local/go/src/net/http/cookie.go:137
				_go_fuzz_dep_.CoverTab[36950]++
										if secs <= 0 {
//line /usr/local/go/src/net/http/cookie.go:138
					_go_fuzz_dep_.CoverTab[36966]++
											secs = -1
//line /usr/local/go/src/net/http/cookie.go:139
					// _ = "end of CoverTab[36966]"
				} else {
//line /usr/local/go/src/net/http/cookie.go:140
					_go_fuzz_dep_.CoverTab[36967]++
//line /usr/local/go/src/net/http/cookie.go:140
					// _ = "end of CoverTab[36967]"
//line /usr/local/go/src/net/http/cookie.go:140
				}
//line /usr/local/go/src/net/http/cookie.go:140
				// _ = "end of CoverTab[36950]"
//line /usr/local/go/src/net/http/cookie.go:140
				_go_fuzz_dep_.CoverTab[36951]++
										c.MaxAge = secs
										continue
//line /usr/local/go/src/net/http/cookie.go:142
				// _ = "end of CoverTab[36951]"
			case "expires":
//line /usr/local/go/src/net/http/cookie.go:143
				_go_fuzz_dep_.CoverTab[36952]++
										c.RawExpires = val
										exptime, err := time.Parse(time.RFC1123, val)
										if err != nil {
//line /usr/local/go/src/net/http/cookie.go:146
					_go_fuzz_dep_.CoverTab[36968]++
											exptime, err = time.Parse("Mon, 02-Jan-2006 15:04:05 MST", val)
											if err != nil {
//line /usr/local/go/src/net/http/cookie.go:148
						_go_fuzz_dep_.CoverTab[36969]++
												c.Expires = time.Time{}
												break
//line /usr/local/go/src/net/http/cookie.go:150
						// _ = "end of CoverTab[36969]"
					} else {
//line /usr/local/go/src/net/http/cookie.go:151
						_go_fuzz_dep_.CoverTab[36970]++
//line /usr/local/go/src/net/http/cookie.go:151
						// _ = "end of CoverTab[36970]"
//line /usr/local/go/src/net/http/cookie.go:151
					}
//line /usr/local/go/src/net/http/cookie.go:151
					// _ = "end of CoverTab[36968]"
				} else {
//line /usr/local/go/src/net/http/cookie.go:152
					_go_fuzz_dep_.CoverTab[36971]++
//line /usr/local/go/src/net/http/cookie.go:152
					// _ = "end of CoverTab[36971]"
//line /usr/local/go/src/net/http/cookie.go:152
				}
//line /usr/local/go/src/net/http/cookie.go:152
				// _ = "end of CoverTab[36952]"
//line /usr/local/go/src/net/http/cookie.go:152
				_go_fuzz_dep_.CoverTab[36953]++
										c.Expires = exptime.UTC()
										continue
//line /usr/local/go/src/net/http/cookie.go:154
				// _ = "end of CoverTab[36953]"
			case "path":
//line /usr/local/go/src/net/http/cookie.go:155
				_go_fuzz_dep_.CoverTab[36954]++
										c.Path = val
										continue
//line /usr/local/go/src/net/http/cookie.go:157
				// _ = "end of CoverTab[36954]"
//line /usr/local/go/src/net/http/cookie.go:157
			default:
//line /usr/local/go/src/net/http/cookie.go:157
				_go_fuzz_dep_.CoverTab[36955]++
//line /usr/local/go/src/net/http/cookie.go:157
				// _ = "end of CoverTab[36955]"
			}
//line /usr/local/go/src/net/http/cookie.go:158
			// _ = "end of CoverTab[36935]"
//line /usr/local/go/src/net/http/cookie.go:158
			_go_fuzz_dep_.CoverTab[36936]++
									c.Unparsed = append(c.Unparsed, parts[i])
//line /usr/local/go/src/net/http/cookie.go:159
			// _ = "end of CoverTab[36936]"
		}
//line /usr/local/go/src/net/http/cookie.go:160
		// _ = "end of CoverTab[36921]"
//line /usr/local/go/src/net/http/cookie.go:160
		_go_fuzz_dep_.CoverTab[36922]++
								cookies = append(cookies, c)
//line /usr/local/go/src/net/http/cookie.go:161
		// _ = "end of CoverTab[36922]"
	}
//line /usr/local/go/src/net/http/cookie.go:162
	// _ = "end of CoverTab[36913]"
//line /usr/local/go/src/net/http/cookie.go:162
	_go_fuzz_dep_.CoverTab[36914]++
							return cookies
//line /usr/local/go/src/net/http/cookie.go:163
	// _ = "end of CoverTab[36914]"
}

// SetCookie adds a Set-Cookie header to the provided ResponseWriter's headers.
//line /usr/local/go/src/net/http/cookie.go:166
// The provided cookie must have a valid Name. Invalid cookies may be
//line /usr/local/go/src/net/http/cookie.go:166
// silently dropped.
//line /usr/local/go/src/net/http/cookie.go:169
func SetCookie(w ResponseWriter, cookie *Cookie) {
//line /usr/local/go/src/net/http/cookie.go:169
	_go_fuzz_dep_.CoverTab[36972]++
							if v := cookie.String(); v != "" {
//line /usr/local/go/src/net/http/cookie.go:170
		_go_fuzz_dep_.CoverTab[36973]++
								w.Header().Add("Set-Cookie", v)
//line /usr/local/go/src/net/http/cookie.go:171
		// _ = "end of CoverTab[36973]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:172
		_go_fuzz_dep_.CoverTab[36974]++
//line /usr/local/go/src/net/http/cookie.go:172
		// _ = "end of CoverTab[36974]"
//line /usr/local/go/src/net/http/cookie.go:172
	}
//line /usr/local/go/src/net/http/cookie.go:172
	// _ = "end of CoverTab[36972]"
}

// String returns the serialization of the cookie for use in a Cookie
//line /usr/local/go/src/net/http/cookie.go:175
// header (if only Name and Value are set) or a Set-Cookie response
//line /usr/local/go/src/net/http/cookie.go:175
// header (if other fields are set).
//line /usr/local/go/src/net/http/cookie.go:175
// If c is nil or c.Name is invalid, the empty string is returned.
//line /usr/local/go/src/net/http/cookie.go:179
func (c *Cookie) String() string {
//line /usr/local/go/src/net/http/cookie.go:179
	_go_fuzz_dep_.CoverTab[36975]++
							if c == nil || func() bool {
//line /usr/local/go/src/net/http/cookie.go:180
		_go_fuzz_dep_.CoverTab[36984]++
//line /usr/local/go/src/net/http/cookie.go:180
		return !isCookieNameValid(c.Name)
//line /usr/local/go/src/net/http/cookie.go:180
		// _ = "end of CoverTab[36984]"
//line /usr/local/go/src/net/http/cookie.go:180
	}() {
//line /usr/local/go/src/net/http/cookie.go:180
		_go_fuzz_dep_.CoverTab[36985]++
								return ""
//line /usr/local/go/src/net/http/cookie.go:181
		// _ = "end of CoverTab[36985]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:182
		_go_fuzz_dep_.CoverTab[36986]++
//line /usr/local/go/src/net/http/cookie.go:182
		// _ = "end of CoverTab[36986]"
//line /usr/local/go/src/net/http/cookie.go:182
	}
//line /usr/local/go/src/net/http/cookie.go:182
	// _ = "end of CoverTab[36975]"
//line /usr/local/go/src/net/http/cookie.go:182
	_go_fuzz_dep_.CoverTab[36976]++
	// extraCookieLength derived from typical length of cookie attributes
	// see RFC 6265 Sec 4.1.
	const extraCookieLength = 110
	var b strings.Builder
	b.Grow(len(c.Name) + len(c.Value) + len(c.Domain) + len(c.Path) + extraCookieLength)
	b.WriteString(c.Name)
	b.WriteRune('=')
	b.WriteString(sanitizeCookieValue(c.Value))

	if len(c.Path) > 0 {
//line /usr/local/go/src/net/http/cookie.go:192
		_go_fuzz_dep_.CoverTab[36987]++
								b.WriteString("; Path=")
								b.WriteString(sanitizeCookiePath(c.Path))
//line /usr/local/go/src/net/http/cookie.go:194
		// _ = "end of CoverTab[36987]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:195
		_go_fuzz_dep_.CoverTab[36988]++
//line /usr/local/go/src/net/http/cookie.go:195
		// _ = "end of CoverTab[36988]"
//line /usr/local/go/src/net/http/cookie.go:195
	}
//line /usr/local/go/src/net/http/cookie.go:195
	// _ = "end of CoverTab[36976]"
//line /usr/local/go/src/net/http/cookie.go:195
	_go_fuzz_dep_.CoverTab[36977]++
							if len(c.Domain) > 0 {
//line /usr/local/go/src/net/http/cookie.go:196
		_go_fuzz_dep_.CoverTab[36989]++
								if validCookieDomain(c.Domain) {
//line /usr/local/go/src/net/http/cookie.go:197
			_go_fuzz_dep_.CoverTab[36990]++

//line /usr/local/go/src/net/http/cookie.go:202
			d := c.Domain
			if d[0] == '.' {
//line /usr/local/go/src/net/http/cookie.go:203
				_go_fuzz_dep_.CoverTab[36992]++
										d = d[1:]
//line /usr/local/go/src/net/http/cookie.go:204
				// _ = "end of CoverTab[36992]"
			} else {
//line /usr/local/go/src/net/http/cookie.go:205
				_go_fuzz_dep_.CoverTab[36993]++
//line /usr/local/go/src/net/http/cookie.go:205
				// _ = "end of CoverTab[36993]"
//line /usr/local/go/src/net/http/cookie.go:205
			}
//line /usr/local/go/src/net/http/cookie.go:205
			// _ = "end of CoverTab[36990]"
//line /usr/local/go/src/net/http/cookie.go:205
			_go_fuzz_dep_.CoverTab[36991]++
									b.WriteString("; Domain=")
									b.WriteString(d)
//line /usr/local/go/src/net/http/cookie.go:207
			// _ = "end of CoverTab[36991]"
		} else {
//line /usr/local/go/src/net/http/cookie.go:208
			_go_fuzz_dep_.CoverTab[36994]++
									log.Printf("net/http: invalid Cookie.Domain %q; dropping domain attribute", c.Domain)
//line /usr/local/go/src/net/http/cookie.go:209
			// _ = "end of CoverTab[36994]"
		}
//line /usr/local/go/src/net/http/cookie.go:210
		// _ = "end of CoverTab[36989]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:211
		_go_fuzz_dep_.CoverTab[36995]++
//line /usr/local/go/src/net/http/cookie.go:211
		// _ = "end of CoverTab[36995]"
//line /usr/local/go/src/net/http/cookie.go:211
	}
//line /usr/local/go/src/net/http/cookie.go:211
	// _ = "end of CoverTab[36977]"
//line /usr/local/go/src/net/http/cookie.go:211
	_go_fuzz_dep_.CoverTab[36978]++
							var buf [len(TimeFormat)]byte
							if validCookieExpires(c.Expires) {
//line /usr/local/go/src/net/http/cookie.go:213
		_go_fuzz_dep_.CoverTab[36996]++
								b.WriteString("; Expires=")
								b.Write(c.Expires.UTC().AppendFormat(buf[:0], TimeFormat))
//line /usr/local/go/src/net/http/cookie.go:215
		// _ = "end of CoverTab[36996]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:216
		_go_fuzz_dep_.CoverTab[36997]++
//line /usr/local/go/src/net/http/cookie.go:216
		// _ = "end of CoverTab[36997]"
//line /usr/local/go/src/net/http/cookie.go:216
	}
//line /usr/local/go/src/net/http/cookie.go:216
	// _ = "end of CoverTab[36978]"
//line /usr/local/go/src/net/http/cookie.go:216
	_go_fuzz_dep_.CoverTab[36979]++
							if c.MaxAge > 0 {
//line /usr/local/go/src/net/http/cookie.go:217
		_go_fuzz_dep_.CoverTab[36998]++
								b.WriteString("; Max-Age=")
								b.Write(strconv.AppendInt(buf[:0], int64(c.MaxAge), 10))
//line /usr/local/go/src/net/http/cookie.go:219
		// _ = "end of CoverTab[36998]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:220
		_go_fuzz_dep_.CoverTab[36999]++
//line /usr/local/go/src/net/http/cookie.go:220
		if c.MaxAge < 0 {
//line /usr/local/go/src/net/http/cookie.go:220
			_go_fuzz_dep_.CoverTab[37000]++
									b.WriteString("; Max-Age=0")
//line /usr/local/go/src/net/http/cookie.go:221
			// _ = "end of CoverTab[37000]"
		} else {
//line /usr/local/go/src/net/http/cookie.go:222
			_go_fuzz_dep_.CoverTab[37001]++
//line /usr/local/go/src/net/http/cookie.go:222
			// _ = "end of CoverTab[37001]"
//line /usr/local/go/src/net/http/cookie.go:222
		}
//line /usr/local/go/src/net/http/cookie.go:222
		// _ = "end of CoverTab[36999]"
//line /usr/local/go/src/net/http/cookie.go:222
	}
//line /usr/local/go/src/net/http/cookie.go:222
	// _ = "end of CoverTab[36979]"
//line /usr/local/go/src/net/http/cookie.go:222
	_go_fuzz_dep_.CoverTab[36980]++
							if c.HttpOnly {
//line /usr/local/go/src/net/http/cookie.go:223
		_go_fuzz_dep_.CoverTab[37002]++
								b.WriteString("; HttpOnly")
//line /usr/local/go/src/net/http/cookie.go:224
		// _ = "end of CoverTab[37002]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:225
		_go_fuzz_dep_.CoverTab[37003]++
//line /usr/local/go/src/net/http/cookie.go:225
		// _ = "end of CoverTab[37003]"
//line /usr/local/go/src/net/http/cookie.go:225
	}
//line /usr/local/go/src/net/http/cookie.go:225
	// _ = "end of CoverTab[36980]"
//line /usr/local/go/src/net/http/cookie.go:225
	_go_fuzz_dep_.CoverTab[36981]++
							if c.Secure {
//line /usr/local/go/src/net/http/cookie.go:226
		_go_fuzz_dep_.CoverTab[37004]++
								b.WriteString("; Secure")
//line /usr/local/go/src/net/http/cookie.go:227
		// _ = "end of CoverTab[37004]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:228
		_go_fuzz_dep_.CoverTab[37005]++
//line /usr/local/go/src/net/http/cookie.go:228
		// _ = "end of CoverTab[37005]"
//line /usr/local/go/src/net/http/cookie.go:228
	}
//line /usr/local/go/src/net/http/cookie.go:228
	// _ = "end of CoverTab[36981]"
//line /usr/local/go/src/net/http/cookie.go:228
	_go_fuzz_dep_.CoverTab[36982]++
							switch c.SameSite {
	case SameSiteDefaultMode:
//line /usr/local/go/src/net/http/cookie.go:230
		_go_fuzz_dep_.CoverTab[37006]++
//line /usr/local/go/src/net/http/cookie.go:230
		// _ = "end of CoverTab[37006]"

	case SameSiteNoneMode:
//line /usr/local/go/src/net/http/cookie.go:232
		_go_fuzz_dep_.CoverTab[37007]++
								b.WriteString("; SameSite=None")
//line /usr/local/go/src/net/http/cookie.go:233
		// _ = "end of CoverTab[37007]"
	case SameSiteLaxMode:
//line /usr/local/go/src/net/http/cookie.go:234
		_go_fuzz_dep_.CoverTab[37008]++
								b.WriteString("; SameSite=Lax")
//line /usr/local/go/src/net/http/cookie.go:235
		// _ = "end of CoverTab[37008]"
	case SameSiteStrictMode:
//line /usr/local/go/src/net/http/cookie.go:236
		_go_fuzz_dep_.CoverTab[37009]++
								b.WriteString("; SameSite=Strict")
//line /usr/local/go/src/net/http/cookie.go:237
		// _ = "end of CoverTab[37009]"
//line /usr/local/go/src/net/http/cookie.go:237
	default:
//line /usr/local/go/src/net/http/cookie.go:237
		_go_fuzz_dep_.CoverTab[37010]++
//line /usr/local/go/src/net/http/cookie.go:237
		// _ = "end of CoverTab[37010]"
	}
//line /usr/local/go/src/net/http/cookie.go:238
	// _ = "end of CoverTab[36982]"
//line /usr/local/go/src/net/http/cookie.go:238
	_go_fuzz_dep_.CoverTab[36983]++
							return b.String()
//line /usr/local/go/src/net/http/cookie.go:239
	// _ = "end of CoverTab[36983]"
}

// Valid reports whether the cookie is valid.
func (c *Cookie) Valid() error {
//line /usr/local/go/src/net/http/cookie.go:243
	_go_fuzz_dep_.CoverTab[37011]++
							if c == nil {
//line /usr/local/go/src/net/http/cookie.go:244
		_go_fuzz_dep_.CoverTab[37018]++
								return errors.New("http: nil Cookie")
//line /usr/local/go/src/net/http/cookie.go:245
		// _ = "end of CoverTab[37018]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:246
		_go_fuzz_dep_.CoverTab[37019]++
//line /usr/local/go/src/net/http/cookie.go:246
		// _ = "end of CoverTab[37019]"
//line /usr/local/go/src/net/http/cookie.go:246
	}
//line /usr/local/go/src/net/http/cookie.go:246
	// _ = "end of CoverTab[37011]"
//line /usr/local/go/src/net/http/cookie.go:246
	_go_fuzz_dep_.CoverTab[37012]++
							if !isCookieNameValid(c.Name) {
//line /usr/local/go/src/net/http/cookie.go:247
		_go_fuzz_dep_.CoverTab[37020]++
								return errors.New("http: invalid Cookie.Name")
//line /usr/local/go/src/net/http/cookie.go:248
		// _ = "end of CoverTab[37020]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:249
		_go_fuzz_dep_.CoverTab[37021]++
//line /usr/local/go/src/net/http/cookie.go:249
		// _ = "end of CoverTab[37021]"
//line /usr/local/go/src/net/http/cookie.go:249
	}
//line /usr/local/go/src/net/http/cookie.go:249
	// _ = "end of CoverTab[37012]"
//line /usr/local/go/src/net/http/cookie.go:249
	_go_fuzz_dep_.CoverTab[37013]++
							if !c.Expires.IsZero() && func() bool {
//line /usr/local/go/src/net/http/cookie.go:250
		_go_fuzz_dep_.CoverTab[37022]++
//line /usr/local/go/src/net/http/cookie.go:250
		return !validCookieExpires(c.Expires)
//line /usr/local/go/src/net/http/cookie.go:250
		// _ = "end of CoverTab[37022]"
//line /usr/local/go/src/net/http/cookie.go:250
	}() {
//line /usr/local/go/src/net/http/cookie.go:250
		_go_fuzz_dep_.CoverTab[37023]++
								return errors.New("http: invalid Cookie.Expires")
//line /usr/local/go/src/net/http/cookie.go:251
		// _ = "end of CoverTab[37023]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:252
		_go_fuzz_dep_.CoverTab[37024]++
//line /usr/local/go/src/net/http/cookie.go:252
		// _ = "end of CoverTab[37024]"
//line /usr/local/go/src/net/http/cookie.go:252
	}
//line /usr/local/go/src/net/http/cookie.go:252
	// _ = "end of CoverTab[37013]"
//line /usr/local/go/src/net/http/cookie.go:252
	_go_fuzz_dep_.CoverTab[37014]++
							for i := 0; i < len(c.Value); i++ {
//line /usr/local/go/src/net/http/cookie.go:253
		_go_fuzz_dep_.CoverTab[37025]++
								if !validCookieValueByte(c.Value[i]) {
//line /usr/local/go/src/net/http/cookie.go:254
			_go_fuzz_dep_.CoverTab[37026]++
									return fmt.Errorf("http: invalid byte %q in Cookie.Value", c.Value[i])
//line /usr/local/go/src/net/http/cookie.go:255
			// _ = "end of CoverTab[37026]"
		} else {
//line /usr/local/go/src/net/http/cookie.go:256
			_go_fuzz_dep_.CoverTab[37027]++
//line /usr/local/go/src/net/http/cookie.go:256
			// _ = "end of CoverTab[37027]"
//line /usr/local/go/src/net/http/cookie.go:256
		}
//line /usr/local/go/src/net/http/cookie.go:256
		// _ = "end of CoverTab[37025]"
	}
//line /usr/local/go/src/net/http/cookie.go:257
	// _ = "end of CoverTab[37014]"
//line /usr/local/go/src/net/http/cookie.go:257
	_go_fuzz_dep_.CoverTab[37015]++
							if len(c.Path) > 0 {
//line /usr/local/go/src/net/http/cookie.go:258
		_go_fuzz_dep_.CoverTab[37028]++
								for i := 0; i < len(c.Path); i++ {
//line /usr/local/go/src/net/http/cookie.go:259
			_go_fuzz_dep_.CoverTab[37029]++
									if !validCookiePathByte(c.Path[i]) {
//line /usr/local/go/src/net/http/cookie.go:260
				_go_fuzz_dep_.CoverTab[37030]++
										return fmt.Errorf("http: invalid byte %q in Cookie.Path", c.Path[i])
//line /usr/local/go/src/net/http/cookie.go:261
				// _ = "end of CoverTab[37030]"
			} else {
//line /usr/local/go/src/net/http/cookie.go:262
				_go_fuzz_dep_.CoverTab[37031]++
//line /usr/local/go/src/net/http/cookie.go:262
				// _ = "end of CoverTab[37031]"
//line /usr/local/go/src/net/http/cookie.go:262
			}
//line /usr/local/go/src/net/http/cookie.go:262
			// _ = "end of CoverTab[37029]"
		}
//line /usr/local/go/src/net/http/cookie.go:263
		// _ = "end of CoverTab[37028]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:264
		_go_fuzz_dep_.CoverTab[37032]++
//line /usr/local/go/src/net/http/cookie.go:264
		// _ = "end of CoverTab[37032]"
//line /usr/local/go/src/net/http/cookie.go:264
	}
//line /usr/local/go/src/net/http/cookie.go:264
	// _ = "end of CoverTab[37015]"
//line /usr/local/go/src/net/http/cookie.go:264
	_go_fuzz_dep_.CoverTab[37016]++
							if len(c.Domain) > 0 {
//line /usr/local/go/src/net/http/cookie.go:265
		_go_fuzz_dep_.CoverTab[37033]++
								if !validCookieDomain(c.Domain) {
//line /usr/local/go/src/net/http/cookie.go:266
			_go_fuzz_dep_.CoverTab[37034]++
									return errors.New("http: invalid Cookie.Domain")
//line /usr/local/go/src/net/http/cookie.go:267
			// _ = "end of CoverTab[37034]"
		} else {
//line /usr/local/go/src/net/http/cookie.go:268
			_go_fuzz_dep_.CoverTab[37035]++
//line /usr/local/go/src/net/http/cookie.go:268
			// _ = "end of CoverTab[37035]"
//line /usr/local/go/src/net/http/cookie.go:268
		}
//line /usr/local/go/src/net/http/cookie.go:268
		// _ = "end of CoverTab[37033]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:269
		_go_fuzz_dep_.CoverTab[37036]++
//line /usr/local/go/src/net/http/cookie.go:269
		// _ = "end of CoverTab[37036]"
//line /usr/local/go/src/net/http/cookie.go:269
	}
//line /usr/local/go/src/net/http/cookie.go:269
	// _ = "end of CoverTab[37016]"
//line /usr/local/go/src/net/http/cookie.go:269
	_go_fuzz_dep_.CoverTab[37017]++
							return nil
//line /usr/local/go/src/net/http/cookie.go:270
	// _ = "end of CoverTab[37017]"
}

// readCookies parses all "Cookie" values from the header h and
//line /usr/local/go/src/net/http/cookie.go:273
// returns the successfully parsed Cookies.
//line /usr/local/go/src/net/http/cookie.go:273
//
//line /usr/local/go/src/net/http/cookie.go:273
// if filter isn't empty, only cookies of that name are returned.
//line /usr/local/go/src/net/http/cookie.go:277
func readCookies(h Header, filter string) []*Cookie {
//line /usr/local/go/src/net/http/cookie.go:277
	_go_fuzz_dep_.CoverTab[37037]++
							lines := h["Cookie"]
							if len(lines) == 0 {
//line /usr/local/go/src/net/http/cookie.go:279
		_go_fuzz_dep_.CoverTab[37040]++
								return []*Cookie{}
//line /usr/local/go/src/net/http/cookie.go:280
		// _ = "end of CoverTab[37040]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:281
		_go_fuzz_dep_.CoverTab[37041]++
//line /usr/local/go/src/net/http/cookie.go:281
		// _ = "end of CoverTab[37041]"
//line /usr/local/go/src/net/http/cookie.go:281
	}
//line /usr/local/go/src/net/http/cookie.go:281
	// _ = "end of CoverTab[37037]"
//line /usr/local/go/src/net/http/cookie.go:281
	_go_fuzz_dep_.CoverTab[37038]++

							cookies := make([]*Cookie, 0, len(lines)+strings.Count(lines[0], ";"))
							for _, line := range lines {
//line /usr/local/go/src/net/http/cookie.go:284
		_go_fuzz_dep_.CoverTab[37042]++
								line = textproto.TrimString(line)

								var part string
								for len(line) > 0 {
//line /usr/local/go/src/net/http/cookie.go:288
			_go_fuzz_dep_.CoverTab[37043]++
									part, line, _ = strings.Cut(line, ";")
									part = textproto.TrimString(part)
									if part == "" {
//line /usr/local/go/src/net/http/cookie.go:291
				_go_fuzz_dep_.CoverTab[37048]++
										continue
//line /usr/local/go/src/net/http/cookie.go:292
				// _ = "end of CoverTab[37048]"
			} else {
//line /usr/local/go/src/net/http/cookie.go:293
				_go_fuzz_dep_.CoverTab[37049]++
//line /usr/local/go/src/net/http/cookie.go:293
				// _ = "end of CoverTab[37049]"
//line /usr/local/go/src/net/http/cookie.go:293
			}
//line /usr/local/go/src/net/http/cookie.go:293
			// _ = "end of CoverTab[37043]"
//line /usr/local/go/src/net/http/cookie.go:293
			_go_fuzz_dep_.CoverTab[37044]++
									name, val, _ := strings.Cut(part, "=")
									name = textproto.TrimString(name)
									if !isCookieNameValid(name) {
//line /usr/local/go/src/net/http/cookie.go:296
				_go_fuzz_dep_.CoverTab[37050]++
										continue
//line /usr/local/go/src/net/http/cookie.go:297
				// _ = "end of CoverTab[37050]"
			} else {
//line /usr/local/go/src/net/http/cookie.go:298
				_go_fuzz_dep_.CoverTab[37051]++
//line /usr/local/go/src/net/http/cookie.go:298
				// _ = "end of CoverTab[37051]"
//line /usr/local/go/src/net/http/cookie.go:298
			}
//line /usr/local/go/src/net/http/cookie.go:298
			// _ = "end of CoverTab[37044]"
//line /usr/local/go/src/net/http/cookie.go:298
			_go_fuzz_dep_.CoverTab[37045]++
									if filter != "" && func() bool {
//line /usr/local/go/src/net/http/cookie.go:299
				_go_fuzz_dep_.CoverTab[37052]++
//line /usr/local/go/src/net/http/cookie.go:299
				return filter != name
//line /usr/local/go/src/net/http/cookie.go:299
				// _ = "end of CoverTab[37052]"
//line /usr/local/go/src/net/http/cookie.go:299
			}() {
//line /usr/local/go/src/net/http/cookie.go:299
				_go_fuzz_dep_.CoverTab[37053]++
										continue
//line /usr/local/go/src/net/http/cookie.go:300
				// _ = "end of CoverTab[37053]"
			} else {
//line /usr/local/go/src/net/http/cookie.go:301
				_go_fuzz_dep_.CoverTab[37054]++
//line /usr/local/go/src/net/http/cookie.go:301
				// _ = "end of CoverTab[37054]"
//line /usr/local/go/src/net/http/cookie.go:301
			}
//line /usr/local/go/src/net/http/cookie.go:301
			// _ = "end of CoverTab[37045]"
//line /usr/local/go/src/net/http/cookie.go:301
			_go_fuzz_dep_.CoverTab[37046]++
									val, ok := parseCookieValue(val, true)
									if !ok {
//line /usr/local/go/src/net/http/cookie.go:303
				_go_fuzz_dep_.CoverTab[37055]++
										continue
//line /usr/local/go/src/net/http/cookie.go:304
				// _ = "end of CoverTab[37055]"
			} else {
//line /usr/local/go/src/net/http/cookie.go:305
				_go_fuzz_dep_.CoverTab[37056]++
//line /usr/local/go/src/net/http/cookie.go:305
				// _ = "end of CoverTab[37056]"
//line /usr/local/go/src/net/http/cookie.go:305
			}
//line /usr/local/go/src/net/http/cookie.go:305
			// _ = "end of CoverTab[37046]"
//line /usr/local/go/src/net/http/cookie.go:305
			_go_fuzz_dep_.CoverTab[37047]++
									cookies = append(cookies, &Cookie{Name: name, Value: val})
//line /usr/local/go/src/net/http/cookie.go:306
			// _ = "end of CoverTab[37047]"
		}
//line /usr/local/go/src/net/http/cookie.go:307
		// _ = "end of CoverTab[37042]"
	}
//line /usr/local/go/src/net/http/cookie.go:308
	// _ = "end of CoverTab[37038]"
//line /usr/local/go/src/net/http/cookie.go:308
	_go_fuzz_dep_.CoverTab[37039]++
							return cookies
//line /usr/local/go/src/net/http/cookie.go:309
	// _ = "end of CoverTab[37039]"
}

// validCookieDomain reports whether v is a valid cookie domain-value.
func validCookieDomain(v string) bool {
//line /usr/local/go/src/net/http/cookie.go:313
	_go_fuzz_dep_.CoverTab[37057]++
							if isCookieDomainName(v) {
//line /usr/local/go/src/net/http/cookie.go:314
		_go_fuzz_dep_.CoverTab[37060]++
								return true
//line /usr/local/go/src/net/http/cookie.go:315
		// _ = "end of CoverTab[37060]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:316
		_go_fuzz_dep_.CoverTab[37061]++
//line /usr/local/go/src/net/http/cookie.go:316
		// _ = "end of CoverTab[37061]"
//line /usr/local/go/src/net/http/cookie.go:316
	}
//line /usr/local/go/src/net/http/cookie.go:316
	// _ = "end of CoverTab[37057]"
//line /usr/local/go/src/net/http/cookie.go:316
	_go_fuzz_dep_.CoverTab[37058]++
							if net.ParseIP(v) != nil && func() bool {
//line /usr/local/go/src/net/http/cookie.go:317
		_go_fuzz_dep_.CoverTab[37062]++
//line /usr/local/go/src/net/http/cookie.go:317
		return !strings.Contains(v, ":")
//line /usr/local/go/src/net/http/cookie.go:317
		// _ = "end of CoverTab[37062]"
//line /usr/local/go/src/net/http/cookie.go:317
	}() {
//line /usr/local/go/src/net/http/cookie.go:317
		_go_fuzz_dep_.CoverTab[37063]++
								return true
//line /usr/local/go/src/net/http/cookie.go:318
		// _ = "end of CoverTab[37063]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:319
		_go_fuzz_dep_.CoverTab[37064]++
//line /usr/local/go/src/net/http/cookie.go:319
		// _ = "end of CoverTab[37064]"
//line /usr/local/go/src/net/http/cookie.go:319
	}
//line /usr/local/go/src/net/http/cookie.go:319
	// _ = "end of CoverTab[37058]"
//line /usr/local/go/src/net/http/cookie.go:319
	_go_fuzz_dep_.CoverTab[37059]++
							return false
//line /usr/local/go/src/net/http/cookie.go:320
	// _ = "end of CoverTab[37059]"
}

// validCookieExpires reports whether v is a valid cookie expires-value.
func validCookieExpires(t time.Time) bool {
//line /usr/local/go/src/net/http/cookie.go:324
	_go_fuzz_dep_.CoverTab[37065]++

							return t.Year() >= 1601
//line /usr/local/go/src/net/http/cookie.go:326
	// _ = "end of CoverTab[37065]"
}

// isCookieDomainName reports whether s is a valid domain name or a valid
//line /usr/local/go/src/net/http/cookie.go:329
// domain name with a leading dot '.'.  It is almost a direct copy of
//line /usr/local/go/src/net/http/cookie.go:329
// package net's isDomainName.
//line /usr/local/go/src/net/http/cookie.go:332
func isCookieDomainName(s string) bool {
//line /usr/local/go/src/net/http/cookie.go:332
	_go_fuzz_dep_.CoverTab[37066]++
							if len(s) == 0 {
//line /usr/local/go/src/net/http/cookie.go:333
		_go_fuzz_dep_.CoverTab[37072]++
								return false
//line /usr/local/go/src/net/http/cookie.go:334
		// _ = "end of CoverTab[37072]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:335
		_go_fuzz_dep_.CoverTab[37073]++
//line /usr/local/go/src/net/http/cookie.go:335
		// _ = "end of CoverTab[37073]"
//line /usr/local/go/src/net/http/cookie.go:335
	}
//line /usr/local/go/src/net/http/cookie.go:335
	// _ = "end of CoverTab[37066]"
//line /usr/local/go/src/net/http/cookie.go:335
	_go_fuzz_dep_.CoverTab[37067]++
							if len(s) > 255 {
//line /usr/local/go/src/net/http/cookie.go:336
		_go_fuzz_dep_.CoverTab[37074]++
								return false
//line /usr/local/go/src/net/http/cookie.go:337
		// _ = "end of CoverTab[37074]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:338
		_go_fuzz_dep_.CoverTab[37075]++
//line /usr/local/go/src/net/http/cookie.go:338
		// _ = "end of CoverTab[37075]"
//line /usr/local/go/src/net/http/cookie.go:338
	}
//line /usr/local/go/src/net/http/cookie.go:338
	// _ = "end of CoverTab[37067]"
//line /usr/local/go/src/net/http/cookie.go:338
	_go_fuzz_dep_.CoverTab[37068]++

							if s[0] == '.' {
//line /usr/local/go/src/net/http/cookie.go:340
		_go_fuzz_dep_.CoverTab[37076]++

								s = s[1:]
//line /usr/local/go/src/net/http/cookie.go:342
		// _ = "end of CoverTab[37076]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:343
		_go_fuzz_dep_.CoverTab[37077]++
//line /usr/local/go/src/net/http/cookie.go:343
		// _ = "end of CoverTab[37077]"
//line /usr/local/go/src/net/http/cookie.go:343
	}
//line /usr/local/go/src/net/http/cookie.go:343
	// _ = "end of CoverTab[37068]"
//line /usr/local/go/src/net/http/cookie.go:343
	_go_fuzz_dep_.CoverTab[37069]++
							last := byte('.')
							ok := false
							partlen := 0
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/http/cookie.go:347
		_go_fuzz_dep_.CoverTab[37078]++
								c := s[i]
								switch {
		default:
//line /usr/local/go/src/net/http/cookie.go:350
			_go_fuzz_dep_.CoverTab[37080]++
									return false
//line /usr/local/go/src/net/http/cookie.go:351
			// _ = "end of CoverTab[37080]"
		case 'a' <= c && func() bool {
//line /usr/local/go/src/net/http/cookie.go:352
			_go_fuzz_dep_.CoverTab[37088]++
//line /usr/local/go/src/net/http/cookie.go:352
			return c <= 'z'
//line /usr/local/go/src/net/http/cookie.go:352
			// _ = "end of CoverTab[37088]"
//line /usr/local/go/src/net/http/cookie.go:352
		}() || func() bool {
//line /usr/local/go/src/net/http/cookie.go:352
			_go_fuzz_dep_.CoverTab[37089]++
//line /usr/local/go/src/net/http/cookie.go:352
			return 'A' <= c && func() bool {
//line /usr/local/go/src/net/http/cookie.go:352
				_go_fuzz_dep_.CoverTab[37090]++
//line /usr/local/go/src/net/http/cookie.go:352
				return c <= 'Z'
//line /usr/local/go/src/net/http/cookie.go:352
				// _ = "end of CoverTab[37090]"
//line /usr/local/go/src/net/http/cookie.go:352
			}()
//line /usr/local/go/src/net/http/cookie.go:352
			// _ = "end of CoverTab[37089]"
//line /usr/local/go/src/net/http/cookie.go:352
		}():
//line /usr/local/go/src/net/http/cookie.go:352
			_go_fuzz_dep_.CoverTab[37081]++

									ok = true
									partlen++
//line /usr/local/go/src/net/http/cookie.go:355
			// _ = "end of CoverTab[37081]"
		case '0' <= c && func() bool {
//line /usr/local/go/src/net/http/cookie.go:356
			_go_fuzz_dep_.CoverTab[37091]++
//line /usr/local/go/src/net/http/cookie.go:356
			return c <= '9'
//line /usr/local/go/src/net/http/cookie.go:356
			// _ = "end of CoverTab[37091]"
//line /usr/local/go/src/net/http/cookie.go:356
		}():
//line /usr/local/go/src/net/http/cookie.go:356
			_go_fuzz_dep_.CoverTab[37082]++

									partlen++
//line /usr/local/go/src/net/http/cookie.go:358
			// _ = "end of CoverTab[37082]"
		case c == '-':
//line /usr/local/go/src/net/http/cookie.go:359
			_go_fuzz_dep_.CoverTab[37083]++

									if last == '.' {
//line /usr/local/go/src/net/http/cookie.go:361
				_go_fuzz_dep_.CoverTab[37092]++
										return false
//line /usr/local/go/src/net/http/cookie.go:362
				// _ = "end of CoverTab[37092]"
			} else {
//line /usr/local/go/src/net/http/cookie.go:363
				_go_fuzz_dep_.CoverTab[37093]++
//line /usr/local/go/src/net/http/cookie.go:363
				// _ = "end of CoverTab[37093]"
//line /usr/local/go/src/net/http/cookie.go:363
			}
//line /usr/local/go/src/net/http/cookie.go:363
			// _ = "end of CoverTab[37083]"
//line /usr/local/go/src/net/http/cookie.go:363
			_go_fuzz_dep_.CoverTab[37084]++
									partlen++
//line /usr/local/go/src/net/http/cookie.go:364
			// _ = "end of CoverTab[37084]"
		case c == '.':
//line /usr/local/go/src/net/http/cookie.go:365
			_go_fuzz_dep_.CoverTab[37085]++

									if last == '.' || func() bool {
//line /usr/local/go/src/net/http/cookie.go:367
				_go_fuzz_dep_.CoverTab[37094]++
//line /usr/local/go/src/net/http/cookie.go:367
				return last == '-'
//line /usr/local/go/src/net/http/cookie.go:367
				// _ = "end of CoverTab[37094]"
//line /usr/local/go/src/net/http/cookie.go:367
			}() {
//line /usr/local/go/src/net/http/cookie.go:367
				_go_fuzz_dep_.CoverTab[37095]++
										return false
//line /usr/local/go/src/net/http/cookie.go:368
				// _ = "end of CoverTab[37095]"
			} else {
//line /usr/local/go/src/net/http/cookie.go:369
				_go_fuzz_dep_.CoverTab[37096]++
//line /usr/local/go/src/net/http/cookie.go:369
				// _ = "end of CoverTab[37096]"
//line /usr/local/go/src/net/http/cookie.go:369
			}
//line /usr/local/go/src/net/http/cookie.go:369
			// _ = "end of CoverTab[37085]"
//line /usr/local/go/src/net/http/cookie.go:369
			_go_fuzz_dep_.CoverTab[37086]++
									if partlen > 63 || func() bool {
//line /usr/local/go/src/net/http/cookie.go:370
				_go_fuzz_dep_.CoverTab[37097]++
//line /usr/local/go/src/net/http/cookie.go:370
				return partlen == 0
//line /usr/local/go/src/net/http/cookie.go:370
				// _ = "end of CoverTab[37097]"
//line /usr/local/go/src/net/http/cookie.go:370
			}() {
//line /usr/local/go/src/net/http/cookie.go:370
				_go_fuzz_dep_.CoverTab[37098]++
										return false
//line /usr/local/go/src/net/http/cookie.go:371
				// _ = "end of CoverTab[37098]"
			} else {
//line /usr/local/go/src/net/http/cookie.go:372
				_go_fuzz_dep_.CoverTab[37099]++
//line /usr/local/go/src/net/http/cookie.go:372
				// _ = "end of CoverTab[37099]"
//line /usr/local/go/src/net/http/cookie.go:372
			}
//line /usr/local/go/src/net/http/cookie.go:372
			// _ = "end of CoverTab[37086]"
//line /usr/local/go/src/net/http/cookie.go:372
			_go_fuzz_dep_.CoverTab[37087]++
									partlen = 0
//line /usr/local/go/src/net/http/cookie.go:373
			// _ = "end of CoverTab[37087]"
		}
//line /usr/local/go/src/net/http/cookie.go:374
		// _ = "end of CoverTab[37078]"
//line /usr/local/go/src/net/http/cookie.go:374
		_go_fuzz_dep_.CoverTab[37079]++
								last = c
//line /usr/local/go/src/net/http/cookie.go:375
		// _ = "end of CoverTab[37079]"
	}
//line /usr/local/go/src/net/http/cookie.go:376
	// _ = "end of CoverTab[37069]"
//line /usr/local/go/src/net/http/cookie.go:376
	_go_fuzz_dep_.CoverTab[37070]++
							if last == '-' || func() bool {
//line /usr/local/go/src/net/http/cookie.go:377
		_go_fuzz_dep_.CoverTab[37100]++
//line /usr/local/go/src/net/http/cookie.go:377
		return partlen > 63
//line /usr/local/go/src/net/http/cookie.go:377
		// _ = "end of CoverTab[37100]"
//line /usr/local/go/src/net/http/cookie.go:377
	}() {
//line /usr/local/go/src/net/http/cookie.go:377
		_go_fuzz_dep_.CoverTab[37101]++
								return false
//line /usr/local/go/src/net/http/cookie.go:378
		// _ = "end of CoverTab[37101]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:379
		_go_fuzz_dep_.CoverTab[37102]++
//line /usr/local/go/src/net/http/cookie.go:379
		// _ = "end of CoverTab[37102]"
//line /usr/local/go/src/net/http/cookie.go:379
	}
//line /usr/local/go/src/net/http/cookie.go:379
	// _ = "end of CoverTab[37070]"
//line /usr/local/go/src/net/http/cookie.go:379
	_go_fuzz_dep_.CoverTab[37071]++

							return ok
//line /usr/local/go/src/net/http/cookie.go:381
	// _ = "end of CoverTab[37071]"
}

var cookieNameSanitizer = strings.NewReplacer("\n", "-", "\r", "-")

func sanitizeCookieName(n string) string {
//line /usr/local/go/src/net/http/cookie.go:386
	_go_fuzz_dep_.CoverTab[37103]++
							return cookieNameSanitizer.Replace(n)
//line /usr/local/go/src/net/http/cookie.go:387
	// _ = "end of CoverTab[37103]"
}

// sanitizeCookieValue produces a suitable cookie-value from v.
//line /usr/local/go/src/net/http/cookie.go:390
// https://tools.ietf.org/html/rfc6265#section-4.1.1
//line /usr/local/go/src/net/http/cookie.go:390
//
//line /usr/local/go/src/net/http/cookie.go:390
//	cookie-value      = *cookie-octet / ( DQUOTE *cookie-octet DQUOTE )
//line /usr/local/go/src/net/http/cookie.go:390
//	cookie-octet      = %x21 / %x23-2B / %x2D-3A / %x3C-5B / %x5D-7E
//line /usr/local/go/src/net/http/cookie.go:390
//	          ; US-ASCII characters excluding CTLs,
//line /usr/local/go/src/net/http/cookie.go:390
//	          ; whitespace DQUOTE, comma, semicolon,
//line /usr/local/go/src/net/http/cookie.go:390
//	          ; and backslash
//line /usr/local/go/src/net/http/cookie.go:390
//
//line /usr/local/go/src/net/http/cookie.go:390
// We loosen this as spaces and commas are common in cookie values
//line /usr/local/go/src/net/http/cookie.go:390
// but we produce a quoted cookie-value if and only if v contains
//line /usr/local/go/src/net/http/cookie.go:390
// commas or spaces.
//line /usr/local/go/src/net/http/cookie.go:390
// See https://golang.org/issue/7243 for the discussion.
//line /usr/local/go/src/net/http/cookie.go:403
func sanitizeCookieValue(v string) string {
//line /usr/local/go/src/net/http/cookie.go:403
	_go_fuzz_dep_.CoverTab[37104]++
							v = sanitizeOrWarn("Cookie.Value", validCookieValueByte, v)
							if len(v) == 0 {
//line /usr/local/go/src/net/http/cookie.go:405
		_go_fuzz_dep_.CoverTab[37107]++
								return v
//line /usr/local/go/src/net/http/cookie.go:406
		// _ = "end of CoverTab[37107]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:407
		_go_fuzz_dep_.CoverTab[37108]++
//line /usr/local/go/src/net/http/cookie.go:407
		// _ = "end of CoverTab[37108]"
//line /usr/local/go/src/net/http/cookie.go:407
	}
//line /usr/local/go/src/net/http/cookie.go:407
	// _ = "end of CoverTab[37104]"
//line /usr/local/go/src/net/http/cookie.go:407
	_go_fuzz_dep_.CoverTab[37105]++
							if strings.ContainsAny(v, " ,") {
//line /usr/local/go/src/net/http/cookie.go:408
		_go_fuzz_dep_.CoverTab[37109]++
								return `"` + v + `"`
//line /usr/local/go/src/net/http/cookie.go:409
		// _ = "end of CoverTab[37109]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:410
		_go_fuzz_dep_.CoverTab[37110]++
//line /usr/local/go/src/net/http/cookie.go:410
		// _ = "end of CoverTab[37110]"
//line /usr/local/go/src/net/http/cookie.go:410
	}
//line /usr/local/go/src/net/http/cookie.go:410
	// _ = "end of CoverTab[37105]"
//line /usr/local/go/src/net/http/cookie.go:410
	_go_fuzz_dep_.CoverTab[37106]++
							return v
//line /usr/local/go/src/net/http/cookie.go:411
	// _ = "end of CoverTab[37106]"
}

func validCookieValueByte(b byte) bool {
//line /usr/local/go/src/net/http/cookie.go:414
	_go_fuzz_dep_.CoverTab[37111]++
							return 0x20 <= b && func() bool {
//line /usr/local/go/src/net/http/cookie.go:415
		_go_fuzz_dep_.CoverTab[37112]++
//line /usr/local/go/src/net/http/cookie.go:415
		return b < 0x7f
//line /usr/local/go/src/net/http/cookie.go:415
		// _ = "end of CoverTab[37112]"
//line /usr/local/go/src/net/http/cookie.go:415
	}() && func() bool {
//line /usr/local/go/src/net/http/cookie.go:415
		_go_fuzz_dep_.CoverTab[37113]++
//line /usr/local/go/src/net/http/cookie.go:415
		return b != '"'
//line /usr/local/go/src/net/http/cookie.go:415
		// _ = "end of CoverTab[37113]"
//line /usr/local/go/src/net/http/cookie.go:415
	}() && func() bool {
//line /usr/local/go/src/net/http/cookie.go:415
		_go_fuzz_dep_.CoverTab[37114]++
//line /usr/local/go/src/net/http/cookie.go:415
		return b != ';'
//line /usr/local/go/src/net/http/cookie.go:415
		// _ = "end of CoverTab[37114]"
//line /usr/local/go/src/net/http/cookie.go:415
	}() && func() bool {
//line /usr/local/go/src/net/http/cookie.go:415
		_go_fuzz_dep_.CoverTab[37115]++
//line /usr/local/go/src/net/http/cookie.go:415
		return b != '\\'
//line /usr/local/go/src/net/http/cookie.go:415
		// _ = "end of CoverTab[37115]"
//line /usr/local/go/src/net/http/cookie.go:415
	}()
//line /usr/local/go/src/net/http/cookie.go:415
	// _ = "end of CoverTab[37111]"
}

// path-av           = "Path=" path-value
//line /usr/local/go/src/net/http/cookie.go:418
// path-value        = <any CHAR except CTLs or ";">
//line /usr/local/go/src/net/http/cookie.go:420
func sanitizeCookiePath(v string) string {
//line /usr/local/go/src/net/http/cookie.go:420
	_go_fuzz_dep_.CoverTab[37116]++
							return sanitizeOrWarn("Cookie.Path", validCookiePathByte, v)
//line /usr/local/go/src/net/http/cookie.go:421
	// _ = "end of CoverTab[37116]"
}

func validCookiePathByte(b byte) bool {
//line /usr/local/go/src/net/http/cookie.go:424
	_go_fuzz_dep_.CoverTab[37117]++
							return 0x20 <= b && func() bool {
//line /usr/local/go/src/net/http/cookie.go:425
		_go_fuzz_dep_.CoverTab[37118]++
//line /usr/local/go/src/net/http/cookie.go:425
		return b < 0x7f
//line /usr/local/go/src/net/http/cookie.go:425
		// _ = "end of CoverTab[37118]"
//line /usr/local/go/src/net/http/cookie.go:425
	}() && func() bool {
//line /usr/local/go/src/net/http/cookie.go:425
		_go_fuzz_dep_.CoverTab[37119]++
//line /usr/local/go/src/net/http/cookie.go:425
		return b != ';'
//line /usr/local/go/src/net/http/cookie.go:425
		// _ = "end of CoverTab[37119]"
//line /usr/local/go/src/net/http/cookie.go:425
	}()
//line /usr/local/go/src/net/http/cookie.go:425
	// _ = "end of CoverTab[37117]"
}

func sanitizeOrWarn(fieldName string, valid func(byte) bool, v string) string {
//line /usr/local/go/src/net/http/cookie.go:428
	_go_fuzz_dep_.CoverTab[37120]++
							ok := true
							for i := 0; i < len(v); i++ {
//line /usr/local/go/src/net/http/cookie.go:430
		_go_fuzz_dep_.CoverTab[37124]++
								if valid(v[i]) {
//line /usr/local/go/src/net/http/cookie.go:431
			_go_fuzz_dep_.CoverTab[37126]++
									continue
//line /usr/local/go/src/net/http/cookie.go:432
			// _ = "end of CoverTab[37126]"
		} else {
//line /usr/local/go/src/net/http/cookie.go:433
			_go_fuzz_dep_.CoverTab[37127]++
//line /usr/local/go/src/net/http/cookie.go:433
			// _ = "end of CoverTab[37127]"
//line /usr/local/go/src/net/http/cookie.go:433
		}
//line /usr/local/go/src/net/http/cookie.go:433
		// _ = "end of CoverTab[37124]"
//line /usr/local/go/src/net/http/cookie.go:433
		_go_fuzz_dep_.CoverTab[37125]++
								log.Printf("net/http: invalid byte %q in %s; dropping invalid bytes", v[i], fieldName)
								ok = false
								break
//line /usr/local/go/src/net/http/cookie.go:436
		// _ = "end of CoverTab[37125]"
	}
//line /usr/local/go/src/net/http/cookie.go:437
	// _ = "end of CoverTab[37120]"
//line /usr/local/go/src/net/http/cookie.go:437
	_go_fuzz_dep_.CoverTab[37121]++
							if ok {
//line /usr/local/go/src/net/http/cookie.go:438
		_go_fuzz_dep_.CoverTab[37128]++
								return v
//line /usr/local/go/src/net/http/cookie.go:439
		// _ = "end of CoverTab[37128]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:440
		_go_fuzz_dep_.CoverTab[37129]++
//line /usr/local/go/src/net/http/cookie.go:440
		// _ = "end of CoverTab[37129]"
//line /usr/local/go/src/net/http/cookie.go:440
	}
//line /usr/local/go/src/net/http/cookie.go:440
	// _ = "end of CoverTab[37121]"
//line /usr/local/go/src/net/http/cookie.go:440
	_go_fuzz_dep_.CoverTab[37122]++
							buf := make([]byte, 0, len(v))
							for i := 0; i < len(v); i++ {
//line /usr/local/go/src/net/http/cookie.go:442
		_go_fuzz_dep_.CoverTab[37130]++
								if b := v[i]; valid(b) {
//line /usr/local/go/src/net/http/cookie.go:443
			_go_fuzz_dep_.CoverTab[37131]++
									buf = append(buf, b)
//line /usr/local/go/src/net/http/cookie.go:444
			// _ = "end of CoverTab[37131]"
		} else {
//line /usr/local/go/src/net/http/cookie.go:445
			_go_fuzz_dep_.CoverTab[37132]++
//line /usr/local/go/src/net/http/cookie.go:445
			// _ = "end of CoverTab[37132]"
//line /usr/local/go/src/net/http/cookie.go:445
		}
//line /usr/local/go/src/net/http/cookie.go:445
		// _ = "end of CoverTab[37130]"
	}
//line /usr/local/go/src/net/http/cookie.go:446
	// _ = "end of CoverTab[37122]"
//line /usr/local/go/src/net/http/cookie.go:446
	_go_fuzz_dep_.CoverTab[37123]++
							return string(buf)
//line /usr/local/go/src/net/http/cookie.go:447
	// _ = "end of CoverTab[37123]"
}

func parseCookieValue(raw string, allowDoubleQuote bool) (string, bool) {
//line /usr/local/go/src/net/http/cookie.go:450
	_go_fuzz_dep_.CoverTab[37133]++

							if allowDoubleQuote && func() bool {
//line /usr/local/go/src/net/http/cookie.go:452
		_go_fuzz_dep_.CoverTab[37136]++
//line /usr/local/go/src/net/http/cookie.go:452
		return len(raw) > 1
//line /usr/local/go/src/net/http/cookie.go:452
		// _ = "end of CoverTab[37136]"
//line /usr/local/go/src/net/http/cookie.go:452
	}() && func() bool {
//line /usr/local/go/src/net/http/cookie.go:452
		_go_fuzz_dep_.CoverTab[37137]++
//line /usr/local/go/src/net/http/cookie.go:452
		return raw[0] == '"'
//line /usr/local/go/src/net/http/cookie.go:452
		// _ = "end of CoverTab[37137]"
//line /usr/local/go/src/net/http/cookie.go:452
	}() && func() bool {
//line /usr/local/go/src/net/http/cookie.go:452
		_go_fuzz_dep_.CoverTab[37138]++
//line /usr/local/go/src/net/http/cookie.go:452
		return raw[len(raw)-1] == '"'
//line /usr/local/go/src/net/http/cookie.go:452
		// _ = "end of CoverTab[37138]"
//line /usr/local/go/src/net/http/cookie.go:452
	}() {
//line /usr/local/go/src/net/http/cookie.go:452
		_go_fuzz_dep_.CoverTab[37139]++
								raw = raw[1 : len(raw)-1]
//line /usr/local/go/src/net/http/cookie.go:453
		// _ = "end of CoverTab[37139]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:454
		_go_fuzz_dep_.CoverTab[37140]++
//line /usr/local/go/src/net/http/cookie.go:454
		// _ = "end of CoverTab[37140]"
//line /usr/local/go/src/net/http/cookie.go:454
	}
//line /usr/local/go/src/net/http/cookie.go:454
	// _ = "end of CoverTab[37133]"
//line /usr/local/go/src/net/http/cookie.go:454
	_go_fuzz_dep_.CoverTab[37134]++
							for i := 0; i < len(raw); i++ {
//line /usr/local/go/src/net/http/cookie.go:455
		_go_fuzz_dep_.CoverTab[37141]++
								if !validCookieValueByte(raw[i]) {
//line /usr/local/go/src/net/http/cookie.go:456
			_go_fuzz_dep_.CoverTab[37142]++
									return "", false
//line /usr/local/go/src/net/http/cookie.go:457
			// _ = "end of CoverTab[37142]"
		} else {
//line /usr/local/go/src/net/http/cookie.go:458
			_go_fuzz_dep_.CoverTab[37143]++
//line /usr/local/go/src/net/http/cookie.go:458
			// _ = "end of CoverTab[37143]"
//line /usr/local/go/src/net/http/cookie.go:458
		}
//line /usr/local/go/src/net/http/cookie.go:458
		// _ = "end of CoverTab[37141]"
	}
//line /usr/local/go/src/net/http/cookie.go:459
	// _ = "end of CoverTab[37134]"
//line /usr/local/go/src/net/http/cookie.go:459
	_go_fuzz_dep_.CoverTab[37135]++
							return raw, true
//line /usr/local/go/src/net/http/cookie.go:460
	// _ = "end of CoverTab[37135]"
}

func isCookieNameValid(raw string) bool {
//line /usr/local/go/src/net/http/cookie.go:463
	_go_fuzz_dep_.CoverTab[37144]++
							if raw == "" {
//line /usr/local/go/src/net/http/cookie.go:464
		_go_fuzz_dep_.CoverTab[37146]++
								return false
//line /usr/local/go/src/net/http/cookie.go:465
		// _ = "end of CoverTab[37146]"
	} else {
//line /usr/local/go/src/net/http/cookie.go:466
		_go_fuzz_dep_.CoverTab[37147]++
//line /usr/local/go/src/net/http/cookie.go:466
		// _ = "end of CoverTab[37147]"
//line /usr/local/go/src/net/http/cookie.go:466
	}
//line /usr/local/go/src/net/http/cookie.go:466
	// _ = "end of CoverTab[37144]"
//line /usr/local/go/src/net/http/cookie.go:466
	_go_fuzz_dep_.CoverTab[37145]++
							return strings.IndexFunc(raw, isNotToken) < 0
//line /usr/local/go/src/net/http/cookie.go:467
	// _ = "end of CoverTab[37145]"
}

//line /usr/local/go/src/net/http/cookie.go:468
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/cookie.go:468
var _ = _go_fuzz_dep_.CoverTab
