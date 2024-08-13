// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:5
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:5
)

import (
	"net/http"
	"sync"
)

var (
	commonBuildOnce		sync.Once
	commonLowerHeader	map[string]string	// Go-Canonical-Case -> lower-case
	commonCanonHeader	map[string]string	// lower-case -> Go-Canonical-Case
)

func buildCommonHeaderMapsOnce() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:18
	_go_fuzz_dep_.CoverTab[73046]++
											commonBuildOnce.Do(buildCommonHeaderMaps)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:19
	// _ = "end of CoverTab[73046]"
}

func buildCommonHeaderMaps() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:22
	_go_fuzz_dep_.CoverTab[73047]++
											common := []string{
		"accept",
		"accept-charset",
		"accept-encoding",
		"accept-language",
		"accept-ranges",
		"age",
		"access-control-allow-credentials",
		"access-control-allow-headers",
		"access-control-allow-methods",
		"access-control-allow-origin",
		"access-control-expose-headers",
		"access-control-max-age",
		"access-control-request-headers",
		"access-control-request-method",
		"allow",
		"authorization",
		"cache-control",
		"content-disposition",
		"content-encoding",
		"content-language",
		"content-length",
		"content-location",
		"content-range",
		"content-type",
		"cookie",
		"date",
		"etag",
		"expect",
		"expires",
		"from",
		"host",
		"if-match",
		"if-modified-since",
		"if-none-match",
		"if-unmodified-since",
		"last-modified",
		"link",
		"location",
		"max-forwards",
		"origin",
		"proxy-authenticate",
		"proxy-authorization",
		"range",
		"referer",
		"refresh",
		"retry-after",
		"server",
		"set-cookie",
		"strict-transport-security",
		"trailer",
		"transfer-encoding",
		"user-agent",
		"vary",
		"via",
		"www-authenticate",
		"x-forwarded-for",
		"x-forwarded-proto",
	}
	commonLowerHeader = make(map[string]string, len(common))
	commonCanonHeader = make(map[string]string, len(common))
	for _, v := range common {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:84
		_go_fuzz_dep_.CoverTab[73048]++
												chk := http.CanonicalHeaderKey(v)
												commonLowerHeader[chk] = v
												commonCanonHeader[v] = chk
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:87
		// _ = "end of CoverTab[73048]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:88
	// _ = "end of CoverTab[73047]"
}

func lowerHeader(v string) (lower string, ascii bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:91
	_go_fuzz_dep_.CoverTab[73049]++
											buildCommonHeaderMapsOnce()
											if s, ok := commonLowerHeader[v]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:93
		_go_fuzz_dep_.CoverTab[73051]++
												return s, true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:94
		// _ = "end of CoverTab[73051]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:95
		_go_fuzz_dep_.CoverTab[73052]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:95
		// _ = "end of CoverTab[73052]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:95
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:95
	// _ = "end of CoverTab[73049]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:95
	_go_fuzz_dep_.CoverTab[73050]++
											return asciiToLower(v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:96
	// _ = "end of CoverTab[73050]"
}

func canonicalHeader(v string) string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:99
	_go_fuzz_dep_.CoverTab[73053]++
											buildCommonHeaderMapsOnce()
											if s, ok := commonCanonHeader[v]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:101
		_go_fuzz_dep_.CoverTab[73055]++
												return s
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:102
		// _ = "end of CoverTab[73055]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:103
		_go_fuzz_dep_.CoverTab[73056]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:103
		// _ = "end of CoverTab[73056]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:103
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:103
	// _ = "end of CoverTab[73053]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:103
	_go_fuzz_dep_.CoverTab[73054]++
											return http.CanonicalHeaderKey(v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:104
	// _ = "end of CoverTab[73054]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:105
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/headermap.go:105
var _ = _go_fuzz_dep_.CoverTab
