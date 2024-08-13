// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/http/jar.go:5
package http

//line /usr/local/go/src/net/http/jar.go:5
import (
//line /usr/local/go/src/net/http/jar.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/jar.go:5
)
//line /usr/local/go/src/net/http/jar.go:5
import (
//line /usr/local/go/src/net/http/jar.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/jar.go:5
)

import (
	"net/url"
)

// A CookieJar manages storage and use of cookies in HTTP requests.
//line /usr/local/go/src/net/http/jar.go:11
//
//line /usr/local/go/src/net/http/jar.go:11
// Implementations of CookieJar must be safe for concurrent use by multiple
//line /usr/local/go/src/net/http/jar.go:11
// goroutines.
//line /usr/local/go/src/net/http/jar.go:11
//
//line /usr/local/go/src/net/http/jar.go:11
// The net/http/cookiejar package provides a CookieJar implementation.
//line /usr/local/go/src/net/http/jar.go:17
type CookieJar interface {
	// SetCookies handles the receipt of the cookies in a reply for the
	// given URL.  It may or may not choose to save the cookies, depending
	// on the jar's policy and implementation.
	SetCookies(u *url.URL, cookies []*Cookie)

	// Cookies returns the cookies to send in a request for the given URL.
	// It is up to the implementation to honor the standard cookie use
	// restrictions such as in RFC 6265.
	Cookies(u *url.URL) []*Cookie
}

//line /usr/local/go/src/net/http/jar.go:27
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/jar.go:27
var _ = _go_fuzz_dep_.CoverTab
