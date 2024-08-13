// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.11
// +build go1.11

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:8
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:8
)

import (
	"net/http/httptrace"
	"net/textproto"
)

func traceHasWroteHeaderField(trace *httptrace.ClientTrace) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:15
	_go_fuzz_dep_.CoverTab[72965]++
										return trace != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:16
		_go_fuzz_dep_.CoverTab[72966]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:16
		return trace.WroteHeaderField != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:16
		// _ = "end of CoverTab[72966]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:16
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:16
	// _ = "end of CoverTab[72965]"
}

func traceWroteHeaderField(trace *httptrace.ClientTrace, k, v string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:19
	_go_fuzz_dep_.CoverTab[72967]++
										if trace != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:20
		_go_fuzz_dep_.CoverTab[72968]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:20
		return trace.WroteHeaderField != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:20
		// _ = "end of CoverTab[72968]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:20
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:20
		_go_fuzz_dep_.CoverTab[72969]++
											trace.WroteHeaderField(k, []string{v})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:21
		// _ = "end of CoverTab[72969]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:22
		_go_fuzz_dep_.CoverTab[72970]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:22
		// _ = "end of CoverTab[72970]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:22
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:22
	// _ = "end of CoverTab[72967]"
}

func traceGot1xxResponseFunc(trace *httptrace.ClientTrace) func(int, textproto.MIMEHeader) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:25
	_go_fuzz_dep_.CoverTab[72971]++
										if trace != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:26
		_go_fuzz_dep_.CoverTab[72973]++
											return trace.Got1xxResponse
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:27
		// _ = "end of CoverTab[72973]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:28
		_go_fuzz_dep_.CoverTab[72974]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:28
		// _ = "end of CoverTab[72974]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:28
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:28
	// _ = "end of CoverTab[72971]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:28
	_go_fuzz_dep_.CoverTab[72972]++
										return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:29
	// _ = "end of CoverTab[72972]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:30
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/go111.go:30
var _ = _go_fuzz_dep_.CoverTab
