// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:5
// Package ctxhttp provides helper functions for performing context-aware HTTP requests.
package ctxhttp

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:6
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:6
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:6
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:6
)

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Do sends an HTTP request with the provided http.Client and returns
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:16
// an HTTP response.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:16
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:16
// If the client is nil, http.DefaultClient is used.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:16
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:16
// The provided ctx must be non-nil. If it is canceled or times out,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:16
// ctx.Err() will be returned.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:23
func Do(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:23
	_go_fuzz_dep_.CoverTab[184026]++
												if client == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:24
		_go_fuzz_dep_.CoverTab[184029]++
													client = http.DefaultClient
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:25
		// _ = "end of CoverTab[184029]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:26
		_go_fuzz_dep_.CoverTab[184030]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:26
		// _ = "end of CoverTab[184030]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:26
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:26
	// _ = "end of CoverTab[184026]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:26
	_go_fuzz_dep_.CoverTab[184027]++
												resp, err := client.Do(req.WithContext(ctx))

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:30
	if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:30
		_go_fuzz_dep_.CoverTab[184031]++
													select {
		case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:32
			_go_fuzz_dep_.CoverTab[184032]++
														err = ctx.Err()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:33
			// _ = "end of CoverTab[184032]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:34
			_go_fuzz_dep_.CoverTab[184033]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:34
			// _ = "end of CoverTab[184033]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:35
		// _ = "end of CoverTab[184031]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:36
		_go_fuzz_dep_.CoverTab[184034]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:36
		// _ = "end of CoverTab[184034]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:36
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:36
	// _ = "end of CoverTab[184027]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:36
	_go_fuzz_dep_.CoverTab[184028]++
												return resp, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:37
	// _ = "end of CoverTab[184028]"
}

// Get issues a GET request via the Do function.
func Get(ctx context.Context, client *http.Client, url string) (*http.Response, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:41
	_go_fuzz_dep_.CoverTab[184035]++
												req, err := http.NewRequest("GET", url, nil)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:43
		_go_fuzz_dep_.CoverTab[184037]++
													return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:44
		// _ = "end of CoverTab[184037]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:45
		_go_fuzz_dep_.CoverTab[184038]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:45
		// _ = "end of CoverTab[184038]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:45
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:45
	// _ = "end of CoverTab[184035]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:45
	_go_fuzz_dep_.CoverTab[184036]++
												return Do(ctx, client, req)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:46
	// _ = "end of CoverTab[184036]"
}

// Head issues a HEAD request via the Do function.
func Head(ctx context.Context, client *http.Client, url string) (*http.Response, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:50
	_go_fuzz_dep_.CoverTab[184039]++
												req, err := http.NewRequest("HEAD", url, nil)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:52
		_go_fuzz_dep_.CoverTab[184041]++
													return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:53
		// _ = "end of CoverTab[184041]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:54
		_go_fuzz_dep_.CoverTab[184042]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:54
		// _ = "end of CoverTab[184042]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:54
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:54
	// _ = "end of CoverTab[184039]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:54
	_go_fuzz_dep_.CoverTab[184040]++
												return Do(ctx, client, req)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:55
	// _ = "end of CoverTab[184040]"
}

// Post issues a POST request via the Do function.
func Post(ctx context.Context, client *http.Client, url string, bodyType string, body io.Reader) (*http.Response, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:59
	_go_fuzz_dep_.CoverTab[184043]++
												req, err := http.NewRequest("POST", url, body)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:61
		_go_fuzz_dep_.CoverTab[184045]++
													return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:62
		// _ = "end of CoverTab[184045]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:63
		_go_fuzz_dep_.CoverTab[184046]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:63
		// _ = "end of CoverTab[184046]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:63
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:63
	// _ = "end of CoverTab[184043]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:63
	_go_fuzz_dep_.CoverTab[184044]++
												req.Header.Set("Content-Type", bodyType)
												return Do(ctx, client, req)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:65
	// _ = "end of CoverTab[184044]"
}

// PostForm issues a POST request via the Do function.
func PostForm(ctx context.Context, client *http.Client, url string, data url.Values) (*http.Response, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:69
	_go_fuzz_dep_.CoverTab[184047]++
												return Post(ctx, client, url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:70
	// _ = "end of CoverTab[184047]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:71
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/context/ctxhttp/ctxhttp.go:71
var _ = _go_fuzz_dep_.CoverTab
