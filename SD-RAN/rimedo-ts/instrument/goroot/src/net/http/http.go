// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate bundle -o=h2_bundle.go -prefix=http2 -tags=!nethttpomithttp2 golang.org/x/net/http2

//line /usr/local/go/src/net/http/http.go:5
//go:generate bundle -o=h2_bundle.go -prefix=http2 -tags=!nethttpomithttp2 golang.org/x/net/http2

package http

//line /usr/local/go/src/net/http/http.go:7
import (
//line /usr/local/go/src/net/http/http.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/http.go:7
)
//line /usr/local/go/src/net/http/http.go:7
import (
//line /usr/local/go/src/net/http/http.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/http.go:7
)

import (
	"io"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/net/http/httpguts"
)

//line /usr/local/go/src/net/http/http.go:22
type incomparable [0]func()

//line /usr/local/go/src/net/http/http.go:26
const maxInt64 = 1<<63 - 1

//line /usr/local/go/src/net/http/http.go:30
var aLongTimeAgo = time.Unix(1, 0)

//line /usr/local/go/src/net/http/http.go:35
var omitBundledHTTP2 bool

//line /usr/local/go/src/net/http/http.go:42
type contextKey struct {
	name string
}

func (k *contextKey) String() string {
//line /usr/local/go/src/net/http/http.go:46
	_go_fuzz_dep_.CoverTab[41449]++
//line /usr/local/go/src/net/http/http.go:46
	return "net/http context value " + k.name
//line /usr/local/go/src/net/http/http.go:46
	// _ = "end of CoverTab[41449]"
//line /usr/local/go/src/net/http/http.go:46
}

//line /usr/local/go/src/net/http/http.go:50
func hasPort(s string) bool {
//line /usr/local/go/src/net/http/http.go:50
	_go_fuzz_dep_.CoverTab[41450]++
//line /usr/local/go/src/net/http/http.go:50
	return strings.LastIndex(s, ":") > strings.LastIndex(s, "]")
//line /usr/local/go/src/net/http/http.go:50
	// _ = "end of CoverTab[41450]"
//line /usr/local/go/src/net/http/http.go:50
}

//line /usr/local/go/src/net/http/http.go:54
func removeEmptyPort(host string) string {
//line /usr/local/go/src/net/http/http.go:54
	_go_fuzz_dep_.CoverTab[41451]++
						if hasPort(host) {
//line /usr/local/go/src/net/http/http.go:55
		_go_fuzz_dep_.CoverTab[41453]++
							return strings.TrimSuffix(host, ":")
//line /usr/local/go/src/net/http/http.go:56
		// _ = "end of CoverTab[41453]"
	} else {
//line /usr/local/go/src/net/http/http.go:57
		_go_fuzz_dep_.CoverTab[41454]++
//line /usr/local/go/src/net/http/http.go:57
		// _ = "end of CoverTab[41454]"
//line /usr/local/go/src/net/http/http.go:57
	}
//line /usr/local/go/src/net/http/http.go:57
	// _ = "end of CoverTab[41451]"
//line /usr/local/go/src/net/http/http.go:57
	_go_fuzz_dep_.CoverTab[41452]++
						return host
//line /usr/local/go/src/net/http/http.go:58
	// _ = "end of CoverTab[41452]"
}

func isNotToken(r rune) bool {
//line /usr/local/go/src/net/http/http.go:61
	_go_fuzz_dep_.CoverTab[41455]++
						return !httpguts.IsTokenRune(r)
//line /usr/local/go/src/net/http/http.go:62
	// _ = "end of CoverTab[41455]"
}

//line /usr/local/go/src/net/http/http.go:66
func stringContainsCTLByte(s string) bool {
//line /usr/local/go/src/net/http/http.go:66
	_go_fuzz_dep_.CoverTab[41456]++
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/http/http.go:67
		_go_fuzz_dep_.CoverTab[41458]++
							b := s[i]
							if b < ' ' || func() bool {
//line /usr/local/go/src/net/http/http.go:69
			_go_fuzz_dep_.CoverTab[41459]++
//line /usr/local/go/src/net/http/http.go:69
			return b == 0x7f
//line /usr/local/go/src/net/http/http.go:69
			// _ = "end of CoverTab[41459]"
//line /usr/local/go/src/net/http/http.go:69
		}() {
//line /usr/local/go/src/net/http/http.go:69
			_go_fuzz_dep_.CoverTab[41460]++
								return true
//line /usr/local/go/src/net/http/http.go:70
			// _ = "end of CoverTab[41460]"
		} else {
//line /usr/local/go/src/net/http/http.go:71
			_go_fuzz_dep_.CoverTab[41461]++
//line /usr/local/go/src/net/http/http.go:71
			// _ = "end of CoverTab[41461]"
//line /usr/local/go/src/net/http/http.go:71
		}
//line /usr/local/go/src/net/http/http.go:71
		// _ = "end of CoverTab[41458]"
	}
//line /usr/local/go/src/net/http/http.go:72
	// _ = "end of CoverTab[41456]"
//line /usr/local/go/src/net/http/http.go:72
	_go_fuzz_dep_.CoverTab[41457]++
						return false
//line /usr/local/go/src/net/http/http.go:73
	// _ = "end of CoverTab[41457]"
}

func hexEscapeNonASCII(s string) string {
//line /usr/local/go/src/net/http/http.go:76
	_go_fuzz_dep_.CoverTab[41462]++
						newLen := 0
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/http/http.go:78
		_go_fuzz_dep_.CoverTab[41466]++
							if s[i] >= utf8.RuneSelf {
//line /usr/local/go/src/net/http/http.go:79
			_go_fuzz_dep_.CoverTab[41467]++
								newLen += 3
//line /usr/local/go/src/net/http/http.go:80
			// _ = "end of CoverTab[41467]"
		} else {
//line /usr/local/go/src/net/http/http.go:81
			_go_fuzz_dep_.CoverTab[41468]++
								newLen++
//line /usr/local/go/src/net/http/http.go:82
			// _ = "end of CoverTab[41468]"
		}
//line /usr/local/go/src/net/http/http.go:83
		// _ = "end of CoverTab[41466]"
	}
//line /usr/local/go/src/net/http/http.go:84
	// _ = "end of CoverTab[41462]"
//line /usr/local/go/src/net/http/http.go:84
	_go_fuzz_dep_.CoverTab[41463]++
						if newLen == len(s) {
//line /usr/local/go/src/net/http/http.go:85
		_go_fuzz_dep_.CoverTab[41469]++
							return s
//line /usr/local/go/src/net/http/http.go:86
		// _ = "end of CoverTab[41469]"
	} else {
//line /usr/local/go/src/net/http/http.go:87
		_go_fuzz_dep_.CoverTab[41470]++
//line /usr/local/go/src/net/http/http.go:87
		// _ = "end of CoverTab[41470]"
//line /usr/local/go/src/net/http/http.go:87
	}
//line /usr/local/go/src/net/http/http.go:87
	// _ = "end of CoverTab[41463]"
//line /usr/local/go/src/net/http/http.go:87
	_go_fuzz_dep_.CoverTab[41464]++
						b := make([]byte, 0, newLen)
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/http/http.go:89
		_go_fuzz_dep_.CoverTab[41471]++
							if s[i] >= utf8.RuneSelf {
//line /usr/local/go/src/net/http/http.go:90
			_go_fuzz_dep_.CoverTab[41472]++
								b = append(b, '%')
								b = strconv.AppendInt(b, int64(s[i]), 16)
//line /usr/local/go/src/net/http/http.go:92
			// _ = "end of CoverTab[41472]"
		} else {
//line /usr/local/go/src/net/http/http.go:93
			_go_fuzz_dep_.CoverTab[41473]++
								b = append(b, s[i])
//line /usr/local/go/src/net/http/http.go:94
			// _ = "end of CoverTab[41473]"
		}
//line /usr/local/go/src/net/http/http.go:95
		// _ = "end of CoverTab[41471]"
	}
//line /usr/local/go/src/net/http/http.go:96
	// _ = "end of CoverTab[41464]"
//line /usr/local/go/src/net/http/http.go:96
	_go_fuzz_dep_.CoverTab[41465]++
						return string(b)
//line /usr/local/go/src/net/http/http.go:97
	// _ = "end of CoverTab[41465]"
}

//line /usr/local/go/src/net/http/http.go:104
var NoBody = noBody{}

type noBody struct{}

func (noBody) Read([]byte) (int, error) {
//line /usr/local/go/src/net/http/http.go:108
	_go_fuzz_dep_.CoverTab[41474]++
//line /usr/local/go/src/net/http/http.go:108
	return 0, io.EOF
//line /usr/local/go/src/net/http/http.go:108
	// _ = "end of CoverTab[41474]"
//line /usr/local/go/src/net/http/http.go:108
}
func (noBody) Close() error {
//line /usr/local/go/src/net/http/http.go:109
	_go_fuzz_dep_.CoverTab[41475]++
//line /usr/local/go/src/net/http/http.go:109
	return nil
//line /usr/local/go/src/net/http/http.go:109
	// _ = "end of CoverTab[41475]"
//line /usr/local/go/src/net/http/http.go:109
}
func (noBody) WriteTo(io.Writer) (int64, error) {
//line /usr/local/go/src/net/http/http.go:110
	_go_fuzz_dep_.CoverTab[41476]++
//line /usr/local/go/src/net/http/http.go:110
	return 0, nil
//line /usr/local/go/src/net/http/http.go:110
	// _ = "end of CoverTab[41476]"
//line /usr/local/go/src/net/http/http.go:110
}

var (
//line /usr/local/go/src/net/http/http.go:114
	_	io.WriterTo	= NoBody
	_	io.ReadCloser	= NoBody
)

//line /usr/local/go/src/net/http/http.go:119
type PushOptions struct {
//line /usr/local/go/src/net/http/http.go:122
	Method	string

//line /usr/local/go/src/net/http/http.go:127
	Header	Header
}

//line /usr/local/go/src/net/http/http.go:133
type Pusher interface {
//line /usr/local/go/src/net/http/http.go:158
	Push(target string, opts *PushOptions) error
}

//line /usr/local/go/src/net/http/http.go:159
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/http.go:159
var _ = _go_fuzz_dep_.CoverTab
