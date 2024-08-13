// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:19
// Package uri parses URIs and implements query escaping.
package uri

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:20
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:20
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:20
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:20
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:20
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:20
)

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:25
import (
	"fmt"
	"sort"
	"strings"

	"github.com/onosproject/onos-lib-go/pkg/errors"
)

// URI represents a parsed URI
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
// The general form represented is:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
//	[scheme:][//[userinfo@]host][/]path[?query][#fragment]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
// URIs that do not start with a slash after the scheme are interpreted as:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
//	scheme:opaque[?query][#fragment]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
// Note that the Path field is stored in decoded form: /%47%6f%2f becomes /Go/.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
// A consequence is that it is impossible to tell which slashes in the Path were
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
// slashes in the raw URI and which were %2f. This distinction is rarely important,
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
// but when it is, the code should use RawPath, an optional field which only gets
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
// set if the default encoding is different from Path.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
// URI's String method uses the EscapedPath method to obtain the path. See the
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:33
// EscapedPath method for more details.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:51
type URI struct {
	Scheme		string
	Opaque		string		// encoded opaque data
	User		*Userinfo	// username and password information
	Host		string		// host or host:port
	Path		string		// path (relative paths may omit leading slash)
	RawPath		string		// encoded path hint (see EscapedPath method)
	ForceQuery	bool		// append a query ('?') even if RawQuery is empty
	RawQuery	string		// encoded query values, without '?'
	Fragment	string		// fragment for references, without '#'
	RawFragment	string		// encoded fragment hint (see EscapedFragment method)
}

// QueryUnescape does the inverse transformation of QueryEscape,
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:64
// converting each 3-byte encoded substring of the form "%AB" into the
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:64
// hex-decoded byte 0xAB.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:64
// It returns an error if any % is not followed by two hexadecimal
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:64
// digits.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:69
func QueryUnescape(s string) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:69
	_go_fuzz_dep_.CoverTab[183079]++
												return unescape(s, encodeQueryComponent)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:70
	// _ = "end of CoverTab[183079]"
}

// PathUnescape does the inverse transformation of PathEscape,
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:73
// converting each 3-byte encoded substring of the form "%AB" into the
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:73
// hex-decoded byte 0xAB. It returns an error if any % is not followed
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:73
// by two hexadecimal digits.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:73
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:73
// PathUnescape is identical to QueryUnescape except that it does not
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:73
// unescape '+' to ' ' (space).
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:80
func PathUnescape(s string) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:80
	_go_fuzz_dep_.CoverTab[183080]++
												return unescape(s, encodePathSegment)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:81
	// _ = "end of CoverTab[183080]"
}

// QueryEscape escapes the string so it can be safely placed
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:84
// inside a URI query.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:86
func QueryEscape(s string) string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:86
	_go_fuzz_dep_.CoverTab[183081]++
												return escape(s, encodeQueryComponent)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:87
	// _ = "end of CoverTab[183081]"
}

// PathEscape escapes the string so it can be safely placed inside a URI path segment,
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:90
// replacing special characters (including /) with %XX sequences as needed.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:92
func PathEscape(s string) string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:92
	_go_fuzz_dep_.CoverTab[183082]++
												return escape(s, encodePathSegment)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:93
	// _ = "end of CoverTab[183082]"
}

// User returns a Userinfo containing the provided username
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:96
// and no password set.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:98
func User(username string) *Userinfo {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:98
	_go_fuzz_dep_.CoverTab[183083]++
												return &Userinfo{username, "", false}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:99
	// _ = "end of CoverTab[183083]"
}

// UserPassword returns a Userinfo containing the provided username
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:102
// and password.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:102
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:102
// This functionality should only be used with legacy web sites.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:102
// RFC 2396 warns that interpreting Userinfo this way
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:102
// “is NOT RECOMMENDED, because the passing of authentication
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:102
// information in clear text (such as URI) has proven to be a
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:102
// security risk in almost every case where it has been used.”
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:110
func UserPassword(username, password string) *Userinfo {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:110
	_go_fuzz_dep_.CoverTab[183084]++
													return &Userinfo{username, password, true}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:111
	// _ = "end of CoverTab[183084]"
}

// The Userinfo type is an immutable encapsulation of username and
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:114
// password details for a URI. An existing Userinfo value is guaranteed
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:114
// to have a username set (potentially empty, as allowed by RFC 2396),
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:114
// and optionally a password.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:118
type Userinfo struct {
	username	string
	password	string
	passwordSet	bool
}

// Username returns the username.
func (u *Userinfo) Username() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:125
	_go_fuzz_dep_.CoverTab[183085]++
													if u == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:126
		_go_fuzz_dep_.CoverTab[183087]++
														return ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:127
		// _ = "end of CoverTab[183087]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:128
		_go_fuzz_dep_.CoverTab[183088]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:128
		// _ = "end of CoverTab[183088]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:128
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:128
	// _ = "end of CoverTab[183085]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:128
	_go_fuzz_dep_.CoverTab[183086]++
													return u.username
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:129
	// _ = "end of CoverTab[183086]"
}

// Password returns the password in case it is set, and whether it is set.
func (u *Userinfo) Password() (string, bool) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:133
	_go_fuzz_dep_.CoverTab[183089]++
													if u == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:134
		_go_fuzz_dep_.CoverTab[183091]++
														return "", false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:135
		// _ = "end of CoverTab[183091]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:136
		_go_fuzz_dep_.CoverTab[183092]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:136
		// _ = "end of CoverTab[183092]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:136
	// _ = "end of CoverTab[183089]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:136
	_go_fuzz_dep_.CoverTab[183090]++
													return u.password, u.passwordSet
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:137
	// _ = "end of CoverTab[183090]"
}

// String returns the encoded userinfo information in the standard form
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:140
// of "username[:password]".
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:142
func (u *Userinfo) String() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:142
	_go_fuzz_dep_.CoverTab[183093]++
													if u == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:143
		_go_fuzz_dep_.CoverTab[183096]++
														return ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:144
		// _ = "end of CoverTab[183096]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:145
		_go_fuzz_dep_.CoverTab[183097]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:145
		// _ = "end of CoverTab[183097]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:145
	// _ = "end of CoverTab[183093]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:145
	_go_fuzz_dep_.CoverTab[183094]++
													s := escape(u.username, encodeUserPassword)
													if u.passwordSet {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:147
		_go_fuzz_dep_.CoverTab[183098]++
														s += ":" + escape(u.password, encodeUserPassword)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:148
		// _ = "end of CoverTab[183098]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:149
		_go_fuzz_dep_.CoverTab[183099]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:149
		// _ = "end of CoverTab[183099]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:149
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:149
	// _ = "end of CoverTab[183094]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:149
	_go_fuzz_dep_.CoverTab[183095]++
													return s
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:150
	// _ = "end of CoverTab[183095]"
}

// Maybe rawURI is of the form scheme:path.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:153
// (Scheme must be [a-zA-Z][a-zA-Z0-9+-.]*)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:153
// If so, return scheme, path; else return "", rawURI.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:156
func getScheme(rawURI string) (scheme, path string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:156
	_go_fuzz_dep_.CoverTab[183100]++
													for i := 0; i < len(rawURI); i++ {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:157
		_go_fuzz_dep_.CoverTab[183102]++
														c := rawURI[i]
														switch {
		case 'a' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:160
			_go_fuzz_dep_.CoverTab[183108]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:160
			return c <= 'z'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:160
			// _ = "end of CoverTab[183108]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:160
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:160
			_go_fuzz_dep_.CoverTab[183109]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:160
			return 'A' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:160
				_go_fuzz_dep_.CoverTab[183110]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:160
				return c <= 'Z'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:160
				// _ = "end of CoverTab[183110]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:160
			}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:160
			// _ = "end of CoverTab[183109]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:160
		}():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:160
			_go_fuzz_dep_.CoverTab[183103]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:160
			// _ = "end of CoverTab[183103]"

		case '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
			_go_fuzz_dep_.CoverTab[183111]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
			return c <= '9'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
			// _ = "end of CoverTab[183111]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
			_go_fuzz_dep_.CoverTab[183112]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
			return c == '+'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
			// _ = "end of CoverTab[183112]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
			_go_fuzz_dep_.CoverTab[183113]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
			return c == '-'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
			// _ = "end of CoverTab[183113]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
			_go_fuzz_dep_.CoverTab[183114]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
			return c == '.'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
			// _ = "end of CoverTab[183114]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
		}():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:162
			_go_fuzz_dep_.CoverTab[183104]++
															if i == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:163
				_go_fuzz_dep_.CoverTab[183115]++
																return "", rawURI, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:164
				// _ = "end of CoverTab[183115]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:165
				_go_fuzz_dep_.CoverTab[183116]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:165
				// _ = "end of CoverTab[183116]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:165
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:165
			// _ = "end of CoverTab[183104]"
		case c == ':':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:166
			_go_fuzz_dep_.CoverTab[183105]++
															if i == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:167
				_go_fuzz_dep_.CoverTab[183117]++
																return "", "", errors.NewInvalid("missing protocol scheme")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:168
				// _ = "end of CoverTab[183117]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:169
				_go_fuzz_dep_.CoverTab[183118]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:169
				// _ = "end of CoverTab[183118]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:169
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:169
			// _ = "end of CoverTab[183105]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:169
			_go_fuzz_dep_.CoverTab[183106]++
															return rawURI[:i], rawURI[i+1:], nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:170
			// _ = "end of CoverTab[183106]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:171
			_go_fuzz_dep_.CoverTab[183107]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:174
			return "", rawURI, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:174
			// _ = "end of CoverTab[183107]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:175
		// _ = "end of CoverTab[183102]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:176
	// _ = "end of CoverTab[183100]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:176
	_go_fuzz_dep_.CoverTab[183101]++
													return "", rawURI, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:177
	// _ = "end of CoverTab[183101]"
}

// Parse parses a raw uri into a URI structure.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:180
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:180
// The uri may be relative (a path, without a host) or absolute
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:180
// (starting with a scheme). Trying to parse a hostname and path
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:180
// without a scheme is invalid but may not necessarily return an
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:180
// error, due to parsing ambiguities.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:186
func Parse(rawURI string) (*URI, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:186
	_go_fuzz_dep_.CoverTab[183119]++

													u, frag := split(rawURI, '#', true)
													uri, err := parse(u, false)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:190
		_go_fuzz_dep_.CoverTab[183123]++
														return nil, errors.NewInvalid("parse error", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:191
		// _ = "end of CoverTab[183123]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:192
		_go_fuzz_dep_.CoverTab[183124]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:192
		// _ = "end of CoverTab[183124]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:192
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:192
	// _ = "end of CoverTab[183119]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:192
	_go_fuzz_dep_.CoverTab[183120]++
													if frag == "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:193
		_go_fuzz_dep_.CoverTab[183125]++
														return uri, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:194
		// _ = "end of CoverTab[183125]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:195
		_go_fuzz_dep_.CoverTab[183126]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:195
		// _ = "end of CoverTab[183126]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:195
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:195
	// _ = "end of CoverTab[183120]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:195
	_go_fuzz_dep_.CoverTab[183121]++
													if err = uri.setFragment(frag); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:196
		_go_fuzz_dep_.CoverTab[183127]++
														return nil, errors.NewInvalid("parse error", err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:197
		// _ = "end of CoverTab[183127]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:198
		_go_fuzz_dep_.CoverTab[183128]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:198
		// _ = "end of CoverTab[183128]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:198
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:198
	// _ = "end of CoverTab[183121]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:198
	_go_fuzz_dep_.CoverTab[183122]++
													return uri, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:199
	// _ = "end of CoverTab[183122]"
}

// parse parses a URI from a string in one of two contexts. If
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:202
// viaRequest is true, the URI is assumed to have arrived via an HTTP request,
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:202
// in which case only absolute URIs or path-absolute relative URIs are allowed.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:202
// If viaRequest is false, all forms of relative URIs are allowed.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:206
func parse(rawURI string, viaRequest bool) (*URI, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:206
	_go_fuzz_dep_.CoverTab[183129]++
													var rest string
													var err error

													if stringContainsCTLByte(rawURI) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:210
		_go_fuzz_dep_.CoverTab[183138]++
														return nil, errors.NewInvalid("invalid control character in URI")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:211
		// _ = "end of CoverTab[183138]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:212
		_go_fuzz_dep_.CoverTab[183139]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:212
		// _ = "end of CoverTab[183139]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:212
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:212
	// _ = "end of CoverTab[183129]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:212
	_go_fuzz_dep_.CoverTab[183130]++

													if rawURI == "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:214
		_go_fuzz_dep_.CoverTab[183140]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:214
		return viaRequest
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:214
		// _ = "end of CoverTab[183140]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:214
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:214
		_go_fuzz_dep_.CoverTab[183141]++
														return nil, errors.NewInvalid("empty uri")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:215
		// _ = "end of CoverTab[183141]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:216
		_go_fuzz_dep_.CoverTab[183142]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:216
		// _ = "end of CoverTab[183142]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:216
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:216
	// _ = "end of CoverTab[183130]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:216
	_go_fuzz_dep_.CoverTab[183131]++
													uri := new(URI)

													if rawURI == "*" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:219
		_go_fuzz_dep_.CoverTab[183143]++
														uri.Path = "*"
														return uri, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:221
		// _ = "end of CoverTab[183143]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:222
		_go_fuzz_dep_.CoverTab[183144]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:222
		// _ = "end of CoverTab[183144]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:222
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:222
	// _ = "end of CoverTab[183131]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:222
	_go_fuzz_dep_.CoverTab[183132]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:226
	if uri.Scheme, rest, err = getScheme(rawURI); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:226
		_go_fuzz_dep_.CoverTab[183145]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:227
		// _ = "end of CoverTab[183145]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:228
		_go_fuzz_dep_.CoverTab[183146]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:228
		// _ = "end of CoverTab[183146]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:228
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:228
	// _ = "end of CoverTab[183132]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:228
	_go_fuzz_dep_.CoverTab[183133]++
													uri.Scheme = strings.ToLower(uri.Scheme)

													if strings.HasSuffix(rest, "?") && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:231
		_go_fuzz_dep_.CoverTab[183147]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:231
		return strings.Count(rest, "?") == 1
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:231
		// _ = "end of CoverTab[183147]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:231
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:231
		_go_fuzz_dep_.CoverTab[183148]++
														uri.ForceQuery = true
														rest = rest[:len(rest)-1]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:233
		// _ = "end of CoverTab[183148]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:234
		_go_fuzz_dep_.CoverTab[183149]++
														rest, uri.RawQuery = split(rest, '?', true)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:235
		// _ = "end of CoverTab[183149]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:236
	// _ = "end of CoverTab[183133]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:236
	_go_fuzz_dep_.CoverTab[183134]++

													if !strings.HasPrefix(rest, "/") {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:238
		_go_fuzz_dep_.CoverTab[183150]++
														if uri.Scheme != "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:239
			_go_fuzz_dep_.CoverTab[183153]++

															uri.Opaque = rest
															return uri, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:242
			// _ = "end of CoverTab[183153]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:243
			_go_fuzz_dep_.CoverTab[183154]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:243
			// _ = "end of CoverTab[183154]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:243
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:243
		// _ = "end of CoverTab[183150]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:243
		_go_fuzz_dep_.CoverTab[183151]++
														if viaRequest {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:244
			_go_fuzz_dep_.CoverTab[183155]++
															return nil, errors.NewInvalid("invalid URI for request")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:245
			// _ = "end of CoverTab[183155]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:246
			_go_fuzz_dep_.CoverTab[183156]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:246
			// _ = "end of CoverTab[183156]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:246
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:246
		// _ = "end of CoverTab[183151]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:246
		_go_fuzz_dep_.CoverTab[183152]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:254
		colon := strings.Index(rest, ":")
		slash := strings.Index(rest, "/")
		if colon >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:256
			_go_fuzz_dep_.CoverTab[183157]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:256
			return (slash < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:256
				_go_fuzz_dep_.CoverTab[183158]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:256
				return colon < slash
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:256
				// _ = "end of CoverTab[183158]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:256
			}())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:256
			// _ = "end of CoverTab[183157]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:256
		}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:256
			_go_fuzz_dep_.CoverTab[183159]++

															return nil, errors.NewInvalid("first path segment in URI cannot contain colon")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:258
			// _ = "end of CoverTab[183159]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:259
			_go_fuzz_dep_.CoverTab[183160]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:259
			// _ = "end of CoverTab[183160]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:259
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:259
		// _ = "end of CoverTab[183152]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:260
		_go_fuzz_dep_.CoverTab[183161]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:260
		// _ = "end of CoverTab[183161]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:260
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:260
	// _ = "end of CoverTab[183134]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:260
	_go_fuzz_dep_.CoverTab[183135]++

													if (uri.Scheme != "" || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:262
		_go_fuzz_dep_.CoverTab[183162]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:262
		return !viaRequest && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:262
			_go_fuzz_dep_.CoverTab[183163]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:262
			return !strings.HasPrefix(rest, "///")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:262
			// _ = "end of CoverTab[183163]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:262
		}()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:262
		// _ = "end of CoverTab[183162]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:262
	}()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:262
		_go_fuzz_dep_.CoverTab[183164]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:262
		return strings.HasPrefix(rest, "//")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:262
		// _ = "end of CoverTab[183164]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:262
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:262
		_go_fuzz_dep_.CoverTab[183165]++
														var authority string
														authority, rest = split(rest[2:], '/', false)
														uri.User, uri.Host, err = parseAuthority(authority)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:266
			_go_fuzz_dep_.CoverTab[183166]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:267
			// _ = "end of CoverTab[183166]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:268
			_go_fuzz_dep_.CoverTab[183167]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:268
			// _ = "end of CoverTab[183167]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:268
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:268
		// _ = "end of CoverTab[183165]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:269
		_go_fuzz_dep_.CoverTab[183168]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:269
		// _ = "end of CoverTab[183168]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:269
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:269
	// _ = "end of CoverTab[183135]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:269
	_go_fuzz_dep_.CoverTab[183136]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:274
	if err := uri.setPath(rest); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:274
		_go_fuzz_dep_.CoverTab[183169]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:275
		// _ = "end of CoverTab[183169]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:276
		_go_fuzz_dep_.CoverTab[183170]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:276
		// _ = "end of CoverTab[183170]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:276
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:276
	// _ = "end of CoverTab[183136]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:276
	_go_fuzz_dep_.CoverTab[183137]++
													return uri, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:277
	// _ = "end of CoverTab[183137]"
}

func parseAuthority(authority string) (user *Userinfo, host string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:280
	_go_fuzz_dep_.CoverTab[183171]++
													i := strings.LastIndex(authority, "@")
													if i < 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:282
		_go_fuzz_dep_.CoverTab[183177]++
														host, err = parseHost(authority)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:283
		// _ = "end of CoverTab[183177]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:284
		_go_fuzz_dep_.CoverTab[183178]++
														host, err = parseHost(authority[i+1:])
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:285
		// _ = "end of CoverTab[183178]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:286
	// _ = "end of CoverTab[183171]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:286
	_go_fuzz_dep_.CoverTab[183172]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:287
		_go_fuzz_dep_.CoverTab[183179]++
														return nil, "", err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:288
		// _ = "end of CoverTab[183179]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:289
		_go_fuzz_dep_.CoverTab[183180]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:289
		// _ = "end of CoverTab[183180]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:289
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:289
	// _ = "end of CoverTab[183172]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:289
	_go_fuzz_dep_.CoverTab[183173]++
													if i < 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:290
		_go_fuzz_dep_.CoverTab[183181]++
														return nil, host, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:291
		// _ = "end of CoverTab[183181]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:292
		_go_fuzz_dep_.CoverTab[183182]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:292
		// _ = "end of CoverTab[183182]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:292
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:292
	// _ = "end of CoverTab[183173]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:292
	_go_fuzz_dep_.CoverTab[183174]++
													userinfo := authority[:i]
													if !validUserinfo(userinfo) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:294
		_go_fuzz_dep_.CoverTab[183183]++
														return nil, "", errors.NewInvalid("invalid userinfo")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:295
		// _ = "end of CoverTab[183183]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:296
		_go_fuzz_dep_.CoverTab[183184]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:296
		// _ = "end of CoverTab[183184]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:296
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:296
	// _ = "end of CoverTab[183174]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:296
	_go_fuzz_dep_.CoverTab[183175]++
													if !strings.Contains(userinfo, ":") {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:297
		_go_fuzz_dep_.CoverTab[183185]++
														if userinfo, err = unescape(userinfo, encodeUserPassword); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:298
			_go_fuzz_dep_.CoverTab[183187]++
															return nil, "", err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:299
			// _ = "end of CoverTab[183187]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:300
			_go_fuzz_dep_.CoverTab[183188]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:300
			// _ = "end of CoverTab[183188]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:300
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:300
		// _ = "end of CoverTab[183185]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:300
		_go_fuzz_dep_.CoverTab[183186]++
														user = User(userinfo)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:301
		// _ = "end of CoverTab[183186]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:302
		_go_fuzz_dep_.CoverTab[183189]++
														username, password := split(userinfo, ':', true)
														if username, err = unescape(username, encodeUserPassword); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:304
			_go_fuzz_dep_.CoverTab[183192]++
															return nil, "", err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:305
			// _ = "end of CoverTab[183192]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:306
			_go_fuzz_dep_.CoverTab[183193]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:306
			// _ = "end of CoverTab[183193]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:306
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:306
		// _ = "end of CoverTab[183189]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:306
		_go_fuzz_dep_.CoverTab[183190]++
														if password, err = unescape(password, encodeUserPassword); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:307
			_go_fuzz_dep_.CoverTab[183194]++
															return nil, "", err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:308
			// _ = "end of CoverTab[183194]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:309
			_go_fuzz_dep_.CoverTab[183195]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:309
			// _ = "end of CoverTab[183195]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:309
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:309
		// _ = "end of CoverTab[183190]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:309
		_go_fuzz_dep_.CoverTab[183191]++
														user = UserPassword(username, password)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:310
		// _ = "end of CoverTab[183191]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:311
	// _ = "end of CoverTab[183175]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:311
	_go_fuzz_dep_.CoverTab[183176]++
													return user, host, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:312
	// _ = "end of CoverTab[183176]"
}

// parseHost parses host as an authority without user
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:315
// information. That is, as host[:port].
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:317
func parseHost(host string) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:317
	_go_fuzz_dep_.CoverTab[183196]++
													if strings.HasPrefix(host, "[") {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:318
		_go_fuzz_dep_.CoverTab[183199]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:321
		i := strings.LastIndex(host, "]")
		if i < 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:322
			_go_fuzz_dep_.CoverTab[183202]++
															return "", errors.NewInvalid("missing ']' in host")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:323
			// _ = "end of CoverTab[183202]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:324
			_go_fuzz_dep_.CoverTab[183203]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:324
			// _ = "end of CoverTab[183203]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:324
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:324
		// _ = "end of CoverTab[183199]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:324
		_go_fuzz_dep_.CoverTab[183200]++
														colonPort := host[i+1:]
														if !validOptionalPort(colonPort) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:326
			_go_fuzz_dep_.CoverTab[183204]++
															return "", errors.NewInvalid("invalid port %q after host", colonPort)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:327
			// _ = "end of CoverTab[183204]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:328
			_go_fuzz_dep_.CoverTab[183205]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:328
			// _ = "end of CoverTab[183205]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:328
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:328
		// _ = "end of CoverTab[183200]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:328
		_go_fuzz_dep_.CoverTab[183201]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:336
		zone := strings.Index(host[:i], "%25")
		if zone >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:337
			_go_fuzz_dep_.CoverTab[183206]++
															host1, err := unescape(host[:zone], encodeHost)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:339
				_go_fuzz_dep_.CoverTab[183210]++
																return "", err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:340
				// _ = "end of CoverTab[183210]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:341
				_go_fuzz_dep_.CoverTab[183211]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:341
				// _ = "end of CoverTab[183211]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:341
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:341
			// _ = "end of CoverTab[183206]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:341
			_go_fuzz_dep_.CoverTab[183207]++
															host2, err := unescape(host[zone:i], encodeZone)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:343
				_go_fuzz_dep_.CoverTab[183212]++
																return "", err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:344
				// _ = "end of CoverTab[183212]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:345
				_go_fuzz_dep_.CoverTab[183213]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:345
				// _ = "end of CoverTab[183213]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:345
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:345
			// _ = "end of CoverTab[183207]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:345
			_go_fuzz_dep_.CoverTab[183208]++
															host3, err := unescape(host[i:], encodeHost)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:347
				_go_fuzz_dep_.CoverTab[183214]++
																return "", err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:348
				// _ = "end of CoverTab[183214]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:349
				_go_fuzz_dep_.CoverTab[183215]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:349
				// _ = "end of CoverTab[183215]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:349
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:349
			// _ = "end of CoverTab[183208]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:349
			_go_fuzz_dep_.CoverTab[183209]++
															return host1 + host2 + host3, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:350
			// _ = "end of CoverTab[183209]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:351
			_go_fuzz_dep_.CoverTab[183216]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:351
			// _ = "end of CoverTab[183216]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:351
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:351
		// _ = "end of CoverTab[183201]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:352
		_go_fuzz_dep_.CoverTab[183217]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:352
		if i := strings.LastIndex(host, ":"); i != -1 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:352
			_go_fuzz_dep_.CoverTab[183218]++
															colonPort := host[i:]
															if !validOptionalPort(colonPort) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:354
				_go_fuzz_dep_.CoverTab[183219]++
																return "", errors.NewInvalid("invalid port %q after host", colonPort)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:355
				// _ = "end of CoverTab[183219]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:356
				_go_fuzz_dep_.CoverTab[183220]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:356
				// _ = "end of CoverTab[183220]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:356
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:356
			// _ = "end of CoverTab[183218]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:357
			_go_fuzz_dep_.CoverTab[183221]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:357
			// _ = "end of CoverTab[183221]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:357
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:357
		// _ = "end of CoverTab[183217]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:357
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:357
	// _ = "end of CoverTab[183196]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:357
	_go_fuzz_dep_.CoverTab[183197]++

													var err error
													if host, err = unescape(host, encodeHost); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:360
		_go_fuzz_dep_.CoverTab[183222]++
														return "", err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:361
		// _ = "end of CoverTab[183222]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:362
		_go_fuzz_dep_.CoverTab[183223]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:362
		// _ = "end of CoverTab[183223]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:362
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:362
	// _ = "end of CoverTab[183197]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:362
	_go_fuzz_dep_.CoverTab[183198]++
													return host, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:363
	// _ = "end of CoverTab[183198]"
}

// setPath sets the Path and RawPath fields of the URI based on the provided
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:366
// escaped path p. It maintains the invariant that RawPath is only specified
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:366
// when it differs from the default encoding of the path.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:366
// For example:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:366
// - setPath("/foo/bar")   will set Path="/foo/bar" and RawPath=""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:366
// - setPath("/foo%2fbar") will set Path="/foo/bar" and RawPath="/foo%2fbar"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:366
// setPath will return an error only if the provided path contains an invalid
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:366
// escaping.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:374
func (u *URI) setPath(p string) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:374
	_go_fuzz_dep_.CoverTab[183224]++
													path, err := unescape(p, encodePath)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:376
		_go_fuzz_dep_.CoverTab[183227]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:377
		// _ = "end of CoverTab[183227]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:378
		_go_fuzz_dep_.CoverTab[183228]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:378
		// _ = "end of CoverTab[183228]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:378
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:378
	// _ = "end of CoverTab[183224]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:378
	_go_fuzz_dep_.CoverTab[183225]++
													u.Path = path
													if escp := escape(path, encodePath); p == escp {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:380
		_go_fuzz_dep_.CoverTab[183229]++

														u.RawPath = ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:382
		// _ = "end of CoverTab[183229]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:383
		_go_fuzz_dep_.CoverTab[183230]++
														u.RawPath = p
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:384
		// _ = "end of CoverTab[183230]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:385
	// _ = "end of CoverTab[183225]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:385
	_go_fuzz_dep_.CoverTab[183226]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:386
	// _ = "end of CoverTab[183226]"
}

// EscapedPath returns the escaped form of u.Path.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:389
// In general there are multiple possible escaped forms of any path.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:389
// EscapedPath returns u.RawPath when it is a valid escaping of u.Path.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:389
// Otherwise EscapedPath ignores u.RawPath and computes an escaped
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:389
// form on its own.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:389
// The String and RequestURI methods use EscapedPath to construct
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:389
// their results.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:389
// In general, code should call EscapedPath instead of
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:389
// reading u.RawPath directly.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:398
func (u *URI) EscapedPath() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:398
	_go_fuzz_dep_.CoverTab[183231]++
													if u.RawPath != "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:399
		_go_fuzz_dep_.CoverTab[183234]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:399
		return validEncoded(u.RawPath, encodePath)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:399
		// _ = "end of CoverTab[183234]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:399
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:399
		_go_fuzz_dep_.CoverTab[183235]++
														p, err := unescape(u.RawPath, encodePath)
														if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:401
			_go_fuzz_dep_.CoverTab[183236]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:401
			return p == u.Path
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:401
			// _ = "end of CoverTab[183236]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:401
		}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:401
			_go_fuzz_dep_.CoverTab[183237]++
															return u.RawPath
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:402
			// _ = "end of CoverTab[183237]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:403
			_go_fuzz_dep_.CoverTab[183238]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:403
			// _ = "end of CoverTab[183238]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:403
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:403
		// _ = "end of CoverTab[183235]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:404
		_go_fuzz_dep_.CoverTab[183239]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:404
		// _ = "end of CoverTab[183239]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:404
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:404
	// _ = "end of CoverTab[183231]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:404
	_go_fuzz_dep_.CoverTab[183232]++
													if u.Path == "*" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:405
		_go_fuzz_dep_.CoverTab[183240]++
														return "*"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:406
		// _ = "end of CoverTab[183240]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:407
		_go_fuzz_dep_.CoverTab[183241]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:407
		// _ = "end of CoverTab[183241]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:407
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:407
	// _ = "end of CoverTab[183232]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:407
	_go_fuzz_dep_.CoverTab[183233]++
													return escape(u.Path, encodePath)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:408
	// _ = "end of CoverTab[183233]"
}

// validEncoded reports whether s is a valid encoded path or fragment,
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:411
// according to mode.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:411
// It must not contain any bytes that require escaping during encoding.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:414
func validEncoded(s string, mode encoding) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:414
	_go_fuzz_dep_.CoverTab[183242]++
													for i := 0; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:415
		_go_fuzz_dep_.CoverTab[183244]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:421
		switch s[i] {
		case '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=', ':', '@':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:422
			_go_fuzz_dep_.CoverTab[183245]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:422
			// _ = "end of CoverTab[183245]"

		case '[', ']':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:424
			_go_fuzz_dep_.CoverTab[183246]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:424
			// _ = "end of CoverTab[183246]"

		case '%':
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:426
			_go_fuzz_dep_.CoverTab[183247]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:426
			// _ = "end of CoverTab[183247]"

		default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:428
			_go_fuzz_dep_.CoverTab[183248]++
															if shouldEscape(s[i], mode) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:429
				_go_fuzz_dep_.CoverTab[183249]++
																return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:430
				// _ = "end of CoverTab[183249]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:431
				_go_fuzz_dep_.CoverTab[183250]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:431
				// _ = "end of CoverTab[183250]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:431
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:431
			// _ = "end of CoverTab[183248]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:432
		// _ = "end of CoverTab[183244]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:433
	// _ = "end of CoverTab[183242]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:433
	_go_fuzz_dep_.CoverTab[183243]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:434
	// _ = "end of CoverTab[183243]"
}

// setFragment is like setPath but for Fragment/RawFragment.
func (u *URI) setFragment(f string) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:438
	_go_fuzz_dep_.CoverTab[183251]++
													frag, err := unescape(f, encodeFragment)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:440
		_go_fuzz_dep_.CoverTab[183254]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:441
		// _ = "end of CoverTab[183254]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:442
		_go_fuzz_dep_.CoverTab[183255]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:442
		// _ = "end of CoverTab[183255]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:442
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:442
	// _ = "end of CoverTab[183251]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:442
	_go_fuzz_dep_.CoverTab[183252]++
													u.Fragment = frag
													if escf := escape(frag, encodeFragment); f == escf {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:444
		_go_fuzz_dep_.CoverTab[183256]++

														u.RawFragment = ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:446
		// _ = "end of CoverTab[183256]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:447
		_go_fuzz_dep_.CoverTab[183257]++
														u.RawFragment = f
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:448
		// _ = "end of CoverTab[183257]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:449
	// _ = "end of CoverTab[183252]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:449
	_go_fuzz_dep_.CoverTab[183253]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:450
	// _ = "end of CoverTab[183253]"
}

// EscapedFragment returns the escaped form of u.Fragment.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:453
// In general there are multiple possible escaped forms of any fragment.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:453
// EscapedFragment returns u.RawFragment when it is a valid escaping of u.Fragment.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:453
// Otherwise EscapedFragment ignores u.RawFragment and computes an escaped
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:453
// form on its own.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:453
// The String method uses EscapedFragment to construct its result.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:453
// In general, code should call EscapedFragment instead of
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:453
// reading u.RawFragment directly.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:461
func (u *URI) EscapedFragment() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:461
	_go_fuzz_dep_.CoverTab[183258]++
													if u.RawFragment != "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:462
		_go_fuzz_dep_.CoverTab[183260]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:462
		return validEncoded(u.RawFragment, encodeFragment)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:462
		// _ = "end of CoverTab[183260]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:462
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:462
		_go_fuzz_dep_.CoverTab[183261]++
														f, err := unescape(u.RawFragment, encodeFragment)
														if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:464
			_go_fuzz_dep_.CoverTab[183262]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:464
			return f == u.Fragment
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:464
			// _ = "end of CoverTab[183262]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:464
		}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:464
			_go_fuzz_dep_.CoverTab[183263]++
															return u.RawFragment
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:465
			// _ = "end of CoverTab[183263]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:466
			_go_fuzz_dep_.CoverTab[183264]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:466
			// _ = "end of CoverTab[183264]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:466
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:466
		// _ = "end of CoverTab[183261]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:467
		_go_fuzz_dep_.CoverTab[183265]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:467
		// _ = "end of CoverTab[183265]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:467
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:467
	// _ = "end of CoverTab[183258]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:467
	_go_fuzz_dep_.CoverTab[183259]++
													return escape(u.Fragment, encodeFragment)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:468
	// _ = "end of CoverTab[183259]"
}

// validOptionalPort reports whether port is either an empty string
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:471
// or matches /^:\d*$/
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:473
func validOptionalPort(port string) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:473
	_go_fuzz_dep_.CoverTab[183266]++
													if port == "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:474
		_go_fuzz_dep_.CoverTab[183270]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:475
		// _ = "end of CoverTab[183270]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:476
		_go_fuzz_dep_.CoverTab[183271]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:476
		// _ = "end of CoverTab[183271]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:476
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:476
	// _ = "end of CoverTab[183266]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:476
	_go_fuzz_dep_.CoverTab[183267]++
													if port[0] != ':' {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:477
		_go_fuzz_dep_.CoverTab[183272]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:478
		// _ = "end of CoverTab[183272]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:479
		_go_fuzz_dep_.CoverTab[183273]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:479
		// _ = "end of CoverTab[183273]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:479
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:479
	// _ = "end of CoverTab[183267]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:479
	_go_fuzz_dep_.CoverTab[183268]++
													for _, b := range port[1:] {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:480
		_go_fuzz_dep_.CoverTab[183274]++
														if b < '0' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:481
			_go_fuzz_dep_.CoverTab[183275]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:481
			return b > '9'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:481
			// _ = "end of CoverTab[183275]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:481
		}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:481
			_go_fuzz_dep_.CoverTab[183276]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:482
			// _ = "end of CoverTab[183276]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:483
			_go_fuzz_dep_.CoverTab[183277]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:483
			// _ = "end of CoverTab[183277]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:483
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:483
		// _ = "end of CoverTab[183274]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:484
	// _ = "end of CoverTab[183268]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:484
	_go_fuzz_dep_.CoverTab[183269]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:485
	// _ = "end of CoverTab[183269]"
}

// String reassembles the URI into a valid URI string.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
// The general form of the result is one of:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
//	scheme:opaque?query#fragment
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
//	scheme:userinfo@host/path?query#fragment
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
// If u.Opaque is non-empty, String uses the first form;
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
// otherwise it uses the second form.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
// Any non-ASCII characters in host are escaped.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
// To obtain the path, String uses u.EscapedPath().
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
// In the second form, the following rules apply:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
//   - if u.Scheme is empty, scheme: is omitted.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
//   - if u.User is nil, userinfo@ is omitted.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
//   - if u.Host is empty, host/ is omitted.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
//   - if u.Scheme and u.Host are empty and u.User is nil,
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
//     the entire scheme://userinfo@host/ is omitted.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
//   - if u.Host is non-empty and u.Path begins with a /,
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
//     the form host/path does not add its own /.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
//   - if u.RawQuery is empty, ?query is omitted.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:488
//   - if u.Fragment is empty, #fragment is omitted.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:509
func (u *URI) String() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:509
	_go_fuzz_dep_.CoverTab[183278]++
													var buf strings.Builder
													if u.Scheme != "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:511
		_go_fuzz_dep_.CoverTab[183283]++
														buf.WriteString(u.Scheme)
														buf.WriteByte(':')
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:513
		// _ = "end of CoverTab[183283]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:514
		_go_fuzz_dep_.CoverTab[183284]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:514
		// _ = "end of CoverTab[183284]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:514
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:514
	// _ = "end of CoverTab[183278]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:514
	_go_fuzz_dep_.CoverTab[183279]++
													if u.Opaque != "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:515
		_go_fuzz_dep_.CoverTab[183285]++
														buf.WriteString(u.Opaque)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:516
		// _ = "end of CoverTab[183285]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:517
		_go_fuzz_dep_.CoverTab[183286]++
														if u.Scheme != "" || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:518
			_go_fuzz_dep_.CoverTab[183290]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:518
			return u.Host != ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:518
			// _ = "end of CoverTab[183290]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:518
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:518
			_go_fuzz_dep_.CoverTab[183291]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:518
			return u.User != nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:518
			// _ = "end of CoverTab[183291]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:518
		}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:518
			_go_fuzz_dep_.CoverTab[183292]++
															if u.Host != "" || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:519
				_go_fuzz_dep_.CoverTab[183295]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:519
				return u.Path != ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:519
				// _ = "end of CoverTab[183295]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:519
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:519
				_go_fuzz_dep_.CoverTab[183296]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:519
				return u.User != nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:519
				// _ = "end of CoverTab[183296]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:519
			}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:519
				_go_fuzz_dep_.CoverTab[183297]++
																buf.WriteString("//")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:520
				// _ = "end of CoverTab[183297]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:521
				_go_fuzz_dep_.CoverTab[183298]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:521
				// _ = "end of CoverTab[183298]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:521
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:521
			// _ = "end of CoverTab[183292]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:521
			_go_fuzz_dep_.CoverTab[183293]++
															if ui := u.User; ui != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:522
				_go_fuzz_dep_.CoverTab[183299]++
																buf.WriteString(ui.String())
																buf.WriteByte('@')
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:524
				// _ = "end of CoverTab[183299]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:525
				_go_fuzz_dep_.CoverTab[183300]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:525
				// _ = "end of CoverTab[183300]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:525
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:525
			// _ = "end of CoverTab[183293]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:525
			_go_fuzz_dep_.CoverTab[183294]++
															if h := u.Host; h != "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:526
				_go_fuzz_dep_.CoverTab[183301]++
																buf.WriteString(escape(h, encodeHost))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:527
				// _ = "end of CoverTab[183301]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:528
				_go_fuzz_dep_.CoverTab[183302]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:528
				// _ = "end of CoverTab[183302]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:528
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:528
			// _ = "end of CoverTab[183294]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:529
			_go_fuzz_dep_.CoverTab[183303]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:529
			// _ = "end of CoverTab[183303]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:529
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:529
		// _ = "end of CoverTab[183286]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:529
		_go_fuzz_dep_.CoverTab[183287]++
														path := u.EscapedPath()
														if path != "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:531
			_go_fuzz_dep_.CoverTab[183304]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:531
			return path[0] != '/'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:531
			// _ = "end of CoverTab[183304]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:531
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:531
			_go_fuzz_dep_.CoverTab[183305]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:531
			return u.Host != ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:531
			// _ = "end of CoverTab[183305]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:531
		}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:531
			_go_fuzz_dep_.CoverTab[183306]++
															buf.WriteByte('/')
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:532
			// _ = "end of CoverTab[183306]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:533
			_go_fuzz_dep_.CoverTab[183307]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:533
			// _ = "end of CoverTab[183307]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:533
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:533
		// _ = "end of CoverTab[183287]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:533
		_go_fuzz_dep_.CoverTab[183288]++
														if buf.Len() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:534
			_go_fuzz_dep_.CoverTab[183308]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:541
			if i := strings.IndexByte(path, ':'); i > -1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:541
				_go_fuzz_dep_.CoverTab[183309]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:541
				return strings.IndexByte(path[:i], '/') == -1
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:541
				// _ = "end of CoverTab[183309]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:541
			}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:541
				_go_fuzz_dep_.CoverTab[183310]++
																buf.WriteString("./")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:542
				// _ = "end of CoverTab[183310]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:543
				_go_fuzz_dep_.CoverTab[183311]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:543
				// _ = "end of CoverTab[183311]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:543
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:543
			// _ = "end of CoverTab[183308]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:544
			_go_fuzz_dep_.CoverTab[183312]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:544
			// _ = "end of CoverTab[183312]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:544
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:544
		// _ = "end of CoverTab[183288]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:544
		_go_fuzz_dep_.CoverTab[183289]++
														buf.WriteString(path)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:545
		// _ = "end of CoverTab[183289]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:546
	// _ = "end of CoverTab[183279]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:546
	_go_fuzz_dep_.CoverTab[183280]++
													if u.ForceQuery || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:547
		_go_fuzz_dep_.CoverTab[183313]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:547
		return u.RawQuery != ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:547
		// _ = "end of CoverTab[183313]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:547
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:547
		_go_fuzz_dep_.CoverTab[183314]++
														buf.WriteByte('?')
														buf.WriteString(u.RawQuery)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:549
		// _ = "end of CoverTab[183314]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:550
		_go_fuzz_dep_.CoverTab[183315]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:550
		// _ = "end of CoverTab[183315]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:550
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:550
	// _ = "end of CoverTab[183280]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:550
	_go_fuzz_dep_.CoverTab[183281]++
													if u.Fragment != "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:551
		_go_fuzz_dep_.CoverTab[183316]++
														buf.WriteByte('#')
														buf.WriteString(u.EscapedFragment())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:553
		// _ = "end of CoverTab[183316]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:554
		_go_fuzz_dep_.CoverTab[183317]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:554
		// _ = "end of CoverTab[183317]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:554
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:554
	// _ = "end of CoverTab[183281]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:554
	_go_fuzz_dep_.CoverTab[183282]++
													return buf.String()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:555
	// _ = "end of CoverTab[183282]"
}

// Redacted is like String but replaces any password with "xxxxx".
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:558
// Only the password in u.URI is redacted.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:560
func (u *URI) Redacted() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:560
	_go_fuzz_dep_.CoverTab[183318]++
													if u == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:561
		_go_fuzz_dep_.CoverTab[183321]++
														return ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:562
		// _ = "end of CoverTab[183321]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:563
		_go_fuzz_dep_.CoverTab[183322]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:563
		// _ = "end of CoverTab[183322]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:563
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:563
	// _ = "end of CoverTab[183318]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:563
	_go_fuzz_dep_.CoverTab[183319]++

													ru := *u
													if _, has := ru.User.Password(); has {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:566
		_go_fuzz_dep_.CoverTab[183323]++
														ru.User = UserPassword(ru.User.Username(), "xxxxx")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:567
		// _ = "end of CoverTab[183323]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:568
		_go_fuzz_dep_.CoverTab[183324]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:568
		// _ = "end of CoverTab[183324]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:568
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:568
	// _ = "end of CoverTab[183319]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:568
	_go_fuzz_dep_.CoverTab[183320]++
													return ru.String()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:569
	// _ = "end of CoverTab[183320]"
}

// Values maps a string key to a list of values.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:572
// It is typically used for query parameters and form values.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:572
// Unlike in the http.Header map, the keys in a Values map
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:572
// are case-sensitive.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:576
type Values map[string][]string

// Get gets the first value associated with the given key.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:578
// If there are no values associated with the key, Get returns
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:578
// the empty string. To access multiple values, use the map
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:578
// directly.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:582
func (v Values) Get(key string) string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:582
	_go_fuzz_dep_.CoverTab[183325]++
													if v == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:583
		_go_fuzz_dep_.CoverTab[183328]++
														return ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:584
		// _ = "end of CoverTab[183328]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:585
		_go_fuzz_dep_.CoverTab[183329]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:585
		// _ = "end of CoverTab[183329]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:585
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:585
	// _ = "end of CoverTab[183325]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:585
	_go_fuzz_dep_.CoverTab[183326]++
													vs := v[key]
													if len(vs) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:587
		_go_fuzz_dep_.CoverTab[183330]++
														return ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:588
		// _ = "end of CoverTab[183330]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:589
		_go_fuzz_dep_.CoverTab[183331]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:589
		// _ = "end of CoverTab[183331]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:589
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:589
	// _ = "end of CoverTab[183326]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:589
	_go_fuzz_dep_.CoverTab[183327]++
													return vs[0]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:590
	// _ = "end of CoverTab[183327]"
}

// Set sets the key to value. It replaces any existing
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:593
// values.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:595
func (v Values) Set(key, value string) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:595
	_go_fuzz_dep_.CoverTab[183332]++
													v[key] = []string{value}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:596
	// _ = "end of CoverTab[183332]"
}

// Add adds the value to key. It appends to any existing
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:599
// values associated with key.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:601
func (v Values) Add(key, value string) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:601
	_go_fuzz_dep_.CoverTab[183333]++
													v[key] = append(v[key], value)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:602
	// _ = "end of CoverTab[183333]"
}

// Del deletes the values associated with key.
func (v Values) Del(key string) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:606
	_go_fuzz_dep_.CoverTab[183334]++
													delete(v, key)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:607
	// _ = "end of CoverTab[183334]"
}

// Has checks whether a given key is set.
func (v Values) Has(key string) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:611
	_go_fuzz_dep_.CoverTab[183335]++
													_, ok := v[key]
													return ok
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:613
	// _ = "end of CoverTab[183335]"
}

// ParseQuery parses the URI-encoded query string and returns
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:616
// a map listing the values specified for each key.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:616
// ParseQuery always returns a non-nil map containing all the
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:616
// valid query parameters found; err describes the first decoding error
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:616
// encountered, if any.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:616
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:616
// Query is expected to be a list of key=value settings separated by ampersands.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:616
// A setting without an equals sign is interpreted as a key set to an empty
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:616
// value.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:616
// Settings containing a non-URI-encoded semicolon are considered invalid.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:626
func ParseQuery(query string) (Values, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:626
	_go_fuzz_dep_.CoverTab[183336]++
													m := make(Values)
													err := parseQuery(m, query)
													return m, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:629
	// _ = "end of CoverTab[183336]"
}

func parseQuery(m Values, query string) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:632
	_go_fuzz_dep_.CoverTab[183337]++
													for query != "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:633
		_go_fuzz_dep_.CoverTab[183339]++
														key := query
														if i := strings.IndexAny(key, "&"); i >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:635
			_go_fuzz_dep_.CoverTab[183346]++
															key, query = key[:i], key[i+1:]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:636
			// _ = "end of CoverTab[183346]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:637
			_go_fuzz_dep_.CoverTab[183347]++
															query = ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:638
			// _ = "end of CoverTab[183347]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:639
		// _ = "end of CoverTab[183339]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:639
		_go_fuzz_dep_.CoverTab[183340]++
														if strings.Contains(key, ";") {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:640
			_go_fuzz_dep_.CoverTab[183348]++
															err = fmt.Errorf("invalid semicolon separator in query")
															continue
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:642
			// _ = "end of CoverTab[183348]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:643
			_go_fuzz_dep_.CoverTab[183349]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:643
			// _ = "end of CoverTab[183349]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:643
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:643
		// _ = "end of CoverTab[183340]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:643
		_go_fuzz_dep_.CoverTab[183341]++
														if key == "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:644
			_go_fuzz_dep_.CoverTab[183350]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:645
			// _ = "end of CoverTab[183350]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:646
			_go_fuzz_dep_.CoverTab[183351]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:646
			// _ = "end of CoverTab[183351]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:646
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:646
		// _ = "end of CoverTab[183341]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:646
		_go_fuzz_dep_.CoverTab[183342]++
														value := ""
														if i := strings.Index(key, "="); i >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:648
			_go_fuzz_dep_.CoverTab[183352]++
															key, value = key[:i], key[i+1:]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:649
			// _ = "end of CoverTab[183352]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:650
			_go_fuzz_dep_.CoverTab[183353]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:650
			// _ = "end of CoverTab[183353]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:650
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:650
		// _ = "end of CoverTab[183342]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:650
		_go_fuzz_dep_.CoverTab[183343]++
														key, err1 := QueryUnescape(key)
														if err1 != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:652
			_go_fuzz_dep_.CoverTab[183354]++
															if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:653
				_go_fuzz_dep_.CoverTab[183356]++
																err = err1
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:654
				// _ = "end of CoverTab[183356]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:655
				_go_fuzz_dep_.CoverTab[183357]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:655
				// _ = "end of CoverTab[183357]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:655
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:655
			// _ = "end of CoverTab[183354]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:655
			_go_fuzz_dep_.CoverTab[183355]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:656
			// _ = "end of CoverTab[183355]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:657
			_go_fuzz_dep_.CoverTab[183358]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:657
			// _ = "end of CoverTab[183358]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:657
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:657
		// _ = "end of CoverTab[183343]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:657
		_go_fuzz_dep_.CoverTab[183344]++
														value, err1 = QueryUnescape(value)
														if err1 != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:659
			_go_fuzz_dep_.CoverTab[183359]++
															if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:660
				_go_fuzz_dep_.CoverTab[183361]++
																err = err1
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:661
				// _ = "end of CoverTab[183361]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:662
				_go_fuzz_dep_.CoverTab[183362]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:662
				// _ = "end of CoverTab[183362]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:662
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:662
			// _ = "end of CoverTab[183359]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:662
			_go_fuzz_dep_.CoverTab[183360]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:663
			// _ = "end of CoverTab[183360]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:664
			_go_fuzz_dep_.CoverTab[183363]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:664
			// _ = "end of CoverTab[183363]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:664
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:664
		// _ = "end of CoverTab[183344]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:664
		_go_fuzz_dep_.CoverTab[183345]++
														m[key] = append(m[key], value)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:665
		// _ = "end of CoverTab[183345]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:666
	// _ = "end of CoverTab[183337]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:666
	_go_fuzz_dep_.CoverTab[183338]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:667
	// _ = "end of CoverTab[183338]"
}

// Encode encodes the values into “URI encoded” form
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:670
// ("bar=baz&foo=quux") sorted by key.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:672
func (v Values) Encode() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:672
	_go_fuzz_dep_.CoverTab[183364]++
													if v == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:673
		_go_fuzz_dep_.CoverTab[183368]++
														return ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:674
		// _ = "end of CoverTab[183368]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:675
		_go_fuzz_dep_.CoverTab[183369]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:675
		// _ = "end of CoverTab[183369]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:675
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:675
	// _ = "end of CoverTab[183364]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:675
	_go_fuzz_dep_.CoverTab[183365]++
													var buf strings.Builder
													keys := make([]string, 0, len(v))
													for k := range v {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:678
		_go_fuzz_dep_.CoverTab[183370]++
														keys = append(keys, k)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:679
		// _ = "end of CoverTab[183370]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:680
	// _ = "end of CoverTab[183365]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:680
	_go_fuzz_dep_.CoverTab[183366]++
													sort.Strings(keys)
													for _, k := range keys {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:682
		_go_fuzz_dep_.CoverTab[183371]++
														vs := v[k]
														keyEscaped := QueryEscape(k)
														for _, v := range vs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:685
			_go_fuzz_dep_.CoverTab[183372]++
															if buf.Len() > 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:686
				_go_fuzz_dep_.CoverTab[183374]++
																buf.WriteByte('&')
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:687
				// _ = "end of CoverTab[183374]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:688
				_go_fuzz_dep_.CoverTab[183375]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:688
				// _ = "end of CoverTab[183375]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:688
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:688
			// _ = "end of CoverTab[183372]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:688
			_go_fuzz_dep_.CoverTab[183373]++
															buf.WriteString(keyEscaped)
															buf.WriteByte('=')
															buf.WriteString(QueryEscape(v))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:691
			// _ = "end of CoverTab[183373]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:692
		// _ = "end of CoverTab[183371]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:693
	// _ = "end of CoverTab[183366]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:693
	_go_fuzz_dep_.CoverTab[183367]++
													return buf.String()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:694
	// _ = "end of CoverTab[183367]"
}

// resolvePath applies special path segments from refs and applies
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:697
// them to base, per RFC 3986.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:699
func resolvePath(base, ref string) string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:699
	_go_fuzz_dep_.CoverTab[183376]++
													var full string
													if ref == "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:701
		_go_fuzz_dep_.CoverTab[183382]++
														full = base
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:702
		// _ = "end of CoverTab[183382]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:703
		_go_fuzz_dep_.CoverTab[183383]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:703
		if ref[0] != '/' {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:703
			_go_fuzz_dep_.CoverTab[183384]++
															i := strings.LastIndex(base, "/")
															full = base[:i+1] + ref
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:705
			// _ = "end of CoverTab[183384]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:706
			_go_fuzz_dep_.CoverTab[183385]++
															full = ref
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:707
			// _ = "end of CoverTab[183385]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:708
		// _ = "end of CoverTab[183383]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:708
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:708
	// _ = "end of CoverTab[183376]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:708
	_go_fuzz_dep_.CoverTab[183377]++
													if full == "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:709
		_go_fuzz_dep_.CoverTab[183386]++
														return ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:710
		// _ = "end of CoverTab[183386]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:711
		_go_fuzz_dep_.CoverTab[183387]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:711
		// _ = "end of CoverTab[183387]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:711
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:711
	// _ = "end of CoverTab[183377]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:711
	_go_fuzz_dep_.CoverTab[183378]++

													var (
		last	string
		elem	string
		i	int
		dst	strings.Builder
	)
	first := true
	remaining := full

	dst.WriteByte('/')
	for i >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:723
		_go_fuzz_dep_.CoverTab[183388]++
														i = strings.IndexByte(remaining, '/')
														if i < 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:725
			_go_fuzz_dep_.CoverTab[183391]++
															last, elem, remaining = remaining, remaining, ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:726
			// _ = "end of CoverTab[183391]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:727
			_go_fuzz_dep_.CoverTab[183392]++
															elem, remaining = remaining[:i], remaining[i+1:]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:728
			// _ = "end of CoverTab[183392]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:729
		// _ = "end of CoverTab[183388]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:729
		_go_fuzz_dep_.CoverTab[183389]++
														if elem == "." {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:730
			_go_fuzz_dep_.CoverTab[183393]++
															first = false

															continue
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:733
			// _ = "end of CoverTab[183393]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:734
			_go_fuzz_dep_.CoverTab[183394]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:734
			// _ = "end of CoverTab[183394]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:734
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:734
		// _ = "end of CoverTab[183389]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:734
		_go_fuzz_dep_.CoverTab[183390]++

														if elem == ".." {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:736
			_go_fuzz_dep_.CoverTab[183395]++

															str := dst.String()[1:]
															index := strings.LastIndexByte(str, '/')

															dst.Reset()
															dst.WriteByte('/')
															if index == -1 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:743
				_go_fuzz_dep_.CoverTab[183396]++
																first = true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:744
				// _ = "end of CoverTab[183396]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:745
				_go_fuzz_dep_.CoverTab[183397]++
																dst.WriteString(str[:index])
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:746
				// _ = "end of CoverTab[183397]"
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:747
			// _ = "end of CoverTab[183395]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:748
			_go_fuzz_dep_.CoverTab[183398]++
															if !first {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:749
				_go_fuzz_dep_.CoverTab[183400]++
																dst.WriteByte('/')
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:750
				// _ = "end of CoverTab[183400]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:751
				_go_fuzz_dep_.CoverTab[183401]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:751
				// _ = "end of CoverTab[183401]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:751
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:751
			// _ = "end of CoverTab[183398]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:751
			_go_fuzz_dep_.CoverTab[183399]++
															dst.WriteString(elem)
															first = false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:753
			// _ = "end of CoverTab[183399]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:754
		// _ = "end of CoverTab[183390]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:755
	// _ = "end of CoverTab[183378]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:755
	_go_fuzz_dep_.CoverTab[183379]++

													if last == "." || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:757
		_go_fuzz_dep_.CoverTab[183402]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:757
		return last == ".."
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:757
		// _ = "end of CoverTab[183402]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:757
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:757
		_go_fuzz_dep_.CoverTab[183403]++
														dst.WriteByte('/')
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:758
		// _ = "end of CoverTab[183403]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:759
		_go_fuzz_dep_.CoverTab[183404]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:759
		// _ = "end of CoverTab[183404]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:759
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:759
	// _ = "end of CoverTab[183379]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:759
	_go_fuzz_dep_.CoverTab[183380]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:762
	r := dst.String()
	if len(r) > 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:763
		_go_fuzz_dep_.CoverTab[183405]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:763
		return r[1] == '/'
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:763
		// _ = "end of CoverTab[183405]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:763
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:763
		_go_fuzz_dep_.CoverTab[183406]++
														r = r[1:]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:764
		// _ = "end of CoverTab[183406]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:765
		_go_fuzz_dep_.CoverTab[183407]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:765
		// _ = "end of CoverTab[183407]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:765
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:765
	// _ = "end of CoverTab[183380]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:765
	_go_fuzz_dep_.CoverTab[183381]++
													return r
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:766
	// _ = "end of CoverTab[183381]"
}

// IsAbs reports whether the URI is absolute.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:769
// Absolute means that it has a non-empty scheme.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:771
func (u *URI) IsAbs() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:771
	_go_fuzz_dep_.CoverTab[183408]++
													return u.Scheme != ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:772
	// _ = "end of CoverTab[183408]"
}

// Parse parses a URI in the context of the receiver. The provided URI
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:775
// may be relative or absolute. Parse returns nil, err on parse
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:775
// failure, otherwise its return value is the same as ResolveReference.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:778
func (u *URI) Parse(ref string) (*URI, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:778
	_go_fuzz_dep_.CoverTab[183409]++
													refURI, err := Parse(ref)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:780
		_go_fuzz_dep_.CoverTab[183411]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:781
		// _ = "end of CoverTab[183411]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:782
		_go_fuzz_dep_.CoverTab[183412]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:782
		// _ = "end of CoverTab[183412]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:782
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:782
	// _ = "end of CoverTab[183409]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:782
	_go_fuzz_dep_.CoverTab[183410]++
													return u.ResolveReference(refURI), nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:783
	// _ = "end of CoverTab[183410]"
}

// ResolveReference resolves a URI reference to an absolute URI from
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:786
// an absolute base URI u, per RFC 3986 Section 5.2. The URI reference
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:786
// may be relative or absolute. ResolveReference always returns a new
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:786
// URI instance, even if the returned URI is identical to either the
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:786
// base or reference. If ref is an absolute URI, then ResolveReference
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:786
// ignores base and returns a copy of ref.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:792
func (u *URI) ResolveReference(ref *URI) *URI {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:792
	_go_fuzz_dep_.CoverTab[183413]++
													url := *ref
													if ref.Scheme == "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:794
		_go_fuzz_dep_.CoverTab[183418]++
														url.Scheme = u.Scheme
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:795
		// _ = "end of CoverTab[183418]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:796
		_go_fuzz_dep_.CoverTab[183419]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:796
		// _ = "end of CoverTab[183419]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:796
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:796
	// _ = "end of CoverTab[183413]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:796
	_go_fuzz_dep_.CoverTab[183414]++
													if ref.Scheme != "" || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:797
		_go_fuzz_dep_.CoverTab[183420]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:797
		return ref.Host != ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:797
		// _ = "end of CoverTab[183420]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:797
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:797
		_go_fuzz_dep_.CoverTab[183421]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:797
		return ref.User != nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:797
		// _ = "end of CoverTab[183421]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:797
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:797
		_go_fuzz_dep_.CoverTab[183422]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:801
		_ = url.setPath(resolvePath(ref.EscapedPath(), ""))
														return &url
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:802
		// _ = "end of CoverTab[183422]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:803
		_go_fuzz_dep_.CoverTab[183423]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:803
		// _ = "end of CoverTab[183423]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:803
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:803
	// _ = "end of CoverTab[183414]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:803
	_go_fuzz_dep_.CoverTab[183415]++
													if ref.Opaque != "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:804
		_go_fuzz_dep_.CoverTab[183424]++
														url.User = nil
														url.Host = ""
														url.Path = ""
														return &url
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:808
		// _ = "end of CoverTab[183424]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:809
		_go_fuzz_dep_.CoverTab[183425]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:809
		// _ = "end of CoverTab[183425]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:809
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:809
	// _ = "end of CoverTab[183415]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:809
	_go_fuzz_dep_.CoverTab[183416]++
													if ref.Path == "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:810
		_go_fuzz_dep_.CoverTab[183426]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:810
		return ref.RawQuery == ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:810
		// _ = "end of CoverTab[183426]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:810
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:810
		_go_fuzz_dep_.CoverTab[183427]++
														url.RawQuery = u.RawQuery
														if ref.Fragment == "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:812
			_go_fuzz_dep_.CoverTab[183428]++
															url.Fragment = u.Fragment
															url.RawFragment = u.RawFragment
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:814
			// _ = "end of CoverTab[183428]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:815
			_go_fuzz_dep_.CoverTab[183429]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:815
			// _ = "end of CoverTab[183429]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:815
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:815
		// _ = "end of CoverTab[183427]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:816
		_go_fuzz_dep_.CoverTab[183430]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:816
		// _ = "end of CoverTab[183430]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:816
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:816
	// _ = "end of CoverTab[183416]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:816
	_go_fuzz_dep_.CoverTab[183417]++

													url.Host = u.Host
													url.User = u.User
													_ = url.setPath(resolvePath(u.EscapedPath(), ref.EscapedPath()))
													return &url
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:821
	// _ = "end of CoverTab[183417]"
}

// Query parses RawQuery and returns the corresponding values.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:824
// It silently discards malformed value pairs.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:824
// To check errors use ParseQuery.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:827
func (u *URI) Query() Values {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:827
	_go_fuzz_dep_.CoverTab[183431]++
													v, _ := ParseQuery(u.RawQuery)
													return v
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:829
	// _ = "end of CoverTab[183431]"
}

// RequestURI returns the encoded path?query or opaque?query
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:832
// string that would be used in an HTTP request for u.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:834
func (u *URI) RequestURI() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:834
	_go_fuzz_dep_.CoverTab[183432]++
													result := u.Opaque
													if result == "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:836
		_go_fuzz_dep_.CoverTab[183435]++
														result = u.EscapedPath()
														if result == "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:838
			_go_fuzz_dep_.CoverTab[183436]++
															result = "/"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:839
			// _ = "end of CoverTab[183436]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:840
			_go_fuzz_dep_.CoverTab[183437]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:840
			// _ = "end of CoverTab[183437]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:840
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:840
		// _ = "end of CoverTab[183435]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:841
		_go_fuzz_dep_.CoverTab[183438]++
														if strings.HasPrefix(result, "//") {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:842
			_go_fuzz_dep_.CoverTab[183439]++
															result = u.Scheme + ":" + result
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:843
			// _ = "end of CoverTab[183439]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:844
			_go_fuzz_dep_.CoverTab[183440]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:844
			// _ = "end of CoverTab[183440]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:844
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:844
		// _ = "end of CoverTab[183438]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:845
	// _ = "end of CoverTab[183432]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:845
	_go_fuzz_dep_.CoverTab[183433]++
													if u.ForceQuery || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:846
		_go_fuzz_dep_.CoverTab[183441]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:846
		return u.RawQuery != ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:846
		// _ = "end of CoverTab[183441]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:846
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:846
		_go_fuzz_dep_.CoverTab[183442]++
														result += "?" + u.RawQuery
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:847
		// _ = "end of CoverTab[183442]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:848
		_go_fuzz_dep_.CoverTab[183443]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:848
		// _ = "end of CoverTab[183443]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:848
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:848
	// _ = "end of CoverTab[183433]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:848
	_go_fuzz_dep_.CoverTab[183434]++
													return result
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:849
	// _ = "end of CoverTab[183434]"
}

// Hostname returns u.Host, stripping any valid port number if present.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:852
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:852
// If the result is enclosed in square brackets, as literal IPv6 addresses are,
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:852
// the square brackets are removed from the result.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:856
func (u *URI) Hostname() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:856
	_go_fuzz_dep_.CoverTab[183444]++
													host, _ := splitHostPort(u.Host)
													return host
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:858
	// _ = "end of CoverTab[183444]"
}

// Port returns the port part of u.Host, without the leading colon.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:861
//
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:861
// If u.Host doesn't contain a valid numeric port, Port returns an empty string.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:864
func (u *URI) Port() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:864
	_go_fuzz_dep_.CoverTab[183445]++
													_, port := splitHostPort(u.Host)
													return port
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:866
	// _ = "end of CoverTab[183445]"
}

// splitHostPort separates host and port. If the port is not valid, it returns
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:869
// the entire input as host, and it doesn't check the validity of the host.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:869
// Unlike net.SplitHostPort, but per RFC 3986, it requires ports to be numeric.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:872
func splitHostPort(hostPort string) (host, port string) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:872
	_go_fuzz_dep_.CoverTab[183446]++
													host = hostPort

													colon := strings.LastIndexByte(host, ':')
													if colon != -1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:876
		_go_fuzz_dep_.CoverTab[183449]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:876
		return validOptionalPort(host[colon:])
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:876
		// _ = "end of CoverTab[183449]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:876
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:876
		_go_fuzz_dep_.CoverTab[183450]++
														host, port = host[:colon], host[colon+1:]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:877
		// _ = "end of CoverTab[183450]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:878
		_go_fuzz_dep_.CoverTab[183451]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:878
		// _ = "end of CoverTab[183451]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:878
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:878
	// _ = "end of CoverTab[183446]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:878
	_go_fuzz_dep_.CoverTab[183447]++

													if strings.HasPrefix(host, "[") && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:880
		_go_fuzz_dep_.CoverTab[183452]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:880
		return strings.HasSuffix(host, "]")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:880
		// _ = "end of CoverTab[183452]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:880
	}() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:880
		_go_fuzz_dep_.CoverTab[183453]++
														host = host[1 : len(host)-1]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:881
		// _ = "end of CoverTab[183453]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:882
		_go_fuzz_dep_.CoverTab[183454]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:882
		// _ = "end of CoverTab[183454]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:882
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:882
	// _ = "end of CoverTab[183447]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:882
	_go_fuzz_dep_.CoverTab[183448]++

													return
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:884
	// _ = "end of CoverTab[183448]"
}

// MarshalBinary Marshaling interface implementations.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:887
// Would like to implement MarshalText/UnmarshalText but that will change the JSON representation of URIs.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:889
func (u *URI) MarshalBinary() (text []byte, err error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:889
	_go_fuzz_dep_.CoverTab[183455]++
													return []byte(u.String()), nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:890
	// _ = "end of CoverTab[183455]"
}

// UnmarshalBinary ...
func (u *URI) UnmarshalBinary(text []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:894
	_go_fuzz_dep_.CoverTab[183456]++
													u1, err := Parse(string(text))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:896
		_go_fuzz_dep_.CoverTab[183458]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:897
		// _ = "end of CoverTab[183458]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:898
		_go_fuzz_dep_.CoverTab[183459]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:898
		// _ = "end of CoverTab[183459]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:898
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:898
	// _ = "end of CoverTab[183456]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:898
	_go_fuzz_dep_.CoverTab[183457]++
													*u = *u1
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:900
	// _ = "end of CoverTab[183457]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:901
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/uri/uri.go:901
var _ = _go_fuzz_dep_.CoverTab
