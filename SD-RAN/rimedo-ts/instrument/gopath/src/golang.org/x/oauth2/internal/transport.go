// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:5
package internal

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:5
)

import (
	"context"
	"net/http"
)

// HTTPClient is the context key to use with golang.org/x/net/context's
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:12
// WithValue function to associate an *http.Client value with a context.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:14
var HTTPClient ContextKey

// ContextKey is just an empty struct. It exists so HTTPClient can be
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:16
// an immutable public variable with a unique type. It's immutable
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:16
// because nobody else can create a ContextKey, being unexported.
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:19
type ContextKey struct{}

var appengineClientHook func(context.Context) *http.Client

func ContextClient(ctx context.Context) *http.Client {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:23
	_go_fuzz_dep_.CoverTab[184150]++
											if ctx != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:24
		_go_fuzz_dep_.CoverTab[184153]++
												if hc, ok := ctx.Value(HTTPClient).(*http.Client); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:25
			_go_fuzz_dep_.CoverTab[184154]++
													return hc
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:26
			// _ = "end of CoverTab[184154]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:27
			_go_fuzz_dep_.CoverTab[184155]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:27
			// _ = "end of CoverTab[184155]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:27
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:27
		// _ = "end of CoverTab[184153]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:28
		_go_fuzz_dep_.CoverTab[184156]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:28
		// _ = "end of CoverTab[184156]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:28
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:28
	// _ = "end of CoverTab[184150]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:28
	_go_fuzz_dep_.CoverTab[184151]++
											if appengineClientHook != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:29
		_go_fuzz_dep_.CoverTab[184157]++
												return appengineClientHook(ctx)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:30
		// _ = "end of CoverTab[184157]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:31
		_go_fuzz_dep_.CoverTab[184158]++
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:31
		// _ = "end of CoverTab[184158]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:31
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:31
	// _ = "end of CoverTab[184151]"
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:31
	_go_fuzz_dep_.CoverTab[184152]++
											return http.DefaultClient
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:32
	// _ = "end of CoverTab[184152]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:33
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/oauth2@v0.4.0/internal/transport.go:33
var _ = _go_fuzz_dep_.CoverTab
