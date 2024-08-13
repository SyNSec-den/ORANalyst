// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:21
)

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
	"sync"

	"go.uber.org/zap/zapcore"
)

const schemeFile = "file"

var (
	_sinkMutex	sync.RWMutex
	_sinkFactories	map[string]func(*url.URL) (Sink, error)	// keyed by scheme
)

func init() {
	resetSinkRegistry()
}

func resetSinkRegistry() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:46
	_go_fuzz_dep_.CoverTab[131780]++
									_sinkMutex.Lock()
									defer _sinkMutex.Unlock()

									_sinkFactories = map[string]func(*url.URL) (Sink, error){
		schemeFile: newFileSink,
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:52
	// _ = "end of CoverTab[131780]"
}

// Sink defines the interface to write to and close logger destinations.
type Sink interface {
	zapcore.WriteSyncer
	io.Closer
}

type nopCloserSink struct{ zapcore.WriteSyncer }

func (nopCloserSink) Close() error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:63
	_go_fuzz_dep_.CoverTab[131781]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:63
	return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:63
	// _ = "end of CoverTab[131781]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:63
}

type errSinkNotFound struct {
	scheme string
}

func (e *errSinkNotFound) Error() string {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:69
	_go_fuzz_dep_.CoverTab[131782]++
									return fmt.Sprintf("no sink found for scheme %q", e.scheme)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:70
	// _ = "end of CoverTab[131782]"
}

// RegisterSink registers a user-supplied factory for all sinks with a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:73
// particular scheme.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:73
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:73
// All schemes must be ASCII, valid under section 3.1 of RFC 3986
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:73
// (https://tools.ietf.org/html/rfc3986#section-3.1), and must not already
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:73
// have a factory registered. Zap automatically registers a factory for the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:73
// "file" scheme.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:80
func RegisterSink(scheme string, factory func(*url.URL) (Sink, error)) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:80
	_go_fuzz_dep_.CoverTab[131783]++
									_sinkMutex.Lock()
									defer _sinkMutex.Unlock()

									if scheme == "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:84
		_go_fuzz_dep_.CoverTab[131787]++
										return errors.New("can't register a sink factory for empty string")
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:85
		// _ = "end of CoverTab[131787]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:86
		_go_fuzz_dep_.CoverTab[131788]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:86
		// _ = "end of CoverTab[131788]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:86
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:86
	// _ = "end of CoverTab[131783]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:86
	_go_fuzz_dep_.CoverTab[131784]++
									normalized, err := normalizeScheme(scheme)
									if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:88
		_go_fuzz_dep_.CoverTab[131789]++
										return fmt.Errorf("%q is not a valid scheme: %v", scheme, err)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:89
		// _ = "end of CoverTab[131789]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:90
		_go_fuzz_dep_.CoverTab[131790]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:90
		// _ = "end of CoverTab[131790]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:90
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:90
	// _ = "end of CoverTab[131784]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:90
	_go_fuzz_dep_.CoverTab[131785]++
									if _, ok := _sinkFactories[normalized]; ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:91
		_go_fuzz_dep_.CoverTab[131791]++
										return fmt.Errorf("sink factory already registered for scheme %q", normalized)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:92
		// _ = "end of CoverTab[131791]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:93
		_go_fuzz_dep_.CoverTab[131792]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:93
		// _ = "end of CoverTab[131792]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:93
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:93
	// _ = "end of CoverTab[131785]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:93
	_go_fuzz_dep_.CoverTab[131786]++
									_sinkFactories[normalized] = factory
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:95
	// _ = "end of CoverTab[131786]"
}

func newSink(rawURL string) (Sink, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:98
	_go_fuzz_dep_.CoverTab[131793]++
									u, err := url.Parse(rawURL)
									if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:100
		_go_fuzz_dep_.CoverTab[131797]++
										return nil, fmt.Errorf("can't parse %q as a URL: %v", rawURL, err)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:101
		// _ = "end of CoverTab[131797]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:102
		_go_fuzz_dep_.CoverTab[131798]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:102
		// _ = "end of CoverTab[131798]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:102
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:102
	// _ = "end of CoverTab[131793]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:102
	_go_fuzz_dep_.CoverTab[131794]++
									if u.Scheme == "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:103
		_go_fuzz_dep_.CoverTab[131799]++
										u.Scheme = schemeFile
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:104
		// _ = "end of CoverTab[131799]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:105
		_go_fuzz_dep_.CoverTab[131800]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:105
		// _ = "end of CoverTab[131800]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:105
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:105
	// _ = "end of CoverTab[131794]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:105
	_go_fuzz_dep_.CoverTab[131795]++

									_sinkMutex.RLock()
									factory, ok := _sinkFactories[u.Scheme]
									_sinkMutex.RUnlock()
									if !ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:110
		_go_fuzz_dep_.CoverTab[131801]++
										return nil, &errSinkNotFound{u.Scheme}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:111
		// _ = "end of CoverTab[131801]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:112
		_go_fuzz_dep_.CoverTab[131802]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:112
		// _ = "end of CoverTab[131802]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:112
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:112
	// _ = "end of CoverTab[131795]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:112
	_go_fuzz_dep_.CoverTab[131796]++
									return factory(u)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:113
	// _ = "end of CoverTab[131796]"
}

func newFileSink(u *url.URL) (Sink, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:116
	_go_fuzz_dep_.CoverTab[131803]++
									if u.User != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:117
		_go_fuzz_dep_.CoverTab[131810]++
										return nil, fmt.Errorf("user and password not allowed with file URLs: got %v", u)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:118
		// _ = "end of CoverTab[131810]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:119
		_go_fuzz_dep_.CoverTab[131811]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:119
		// _ = "end of CoverTab[131811]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:119
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:119
	// _ = "end of CoverTab[131803]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:119
	_go_fuzz_dep_.CoverTab[131804]++
									if u.Fragment != "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:120
		_go_fuzz_dep_.CoverTab[131812]++
										return nil, fmt.Errorf("fragments not allowed with file URLs: got %v", u)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:121
		// _ = "end of CoverTab[131812]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:122
		_go_fuzz_dep_.CoverTab[131813]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:122
		// _ = "end of CoverTab[131813]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:122
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:122
	// _ = "end of CoverTab[131804]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:122
	_go_fuzz_dep_.CoverTab[131805]++
									if u.RawQuery != "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:123
		_go_fuzz_dep_.CoverTab[131814]++
										return nil, fmt.Errorf("query parameters not allowed with file URLs: got %v", u)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:124
		// _ = "end of CoverTab[131814]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:125
		_go_fuzz_dep_.CoverTab[131815]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:125
		// _ = "end of CoverTab[131815]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:125
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:125
	// _ = "end of CoverTab[131805]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:125
	_go_fuzz_dep_.CoverTab[131806]++

									if u.Port() != "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:127
		_go_fuzz_dep_.CoverTab[131816]++
										return nil, fmt.Errorf("ports not allowed with file URLs: got %v", u)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:128
		// _ = "end of CoverTab[131816]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:129
		_go_fuzz_dep_.CoverTab[131817]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:129
		// _ = "end of CoverTab[131817]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:129
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:129
	// _ = "end of CoverTab[131806]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:129
	_go_fuzz_dep_.CoverTab[131807]++
									if hn := u.Hostname(); hn != "" && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:130
		_go_fuzz_dep_.CoverTab[131818]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:130
		return hn != "localhost"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:130
		// _ = "end of CoverTab[131818]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:130
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:130
		_go_fuzz_dep_.CoverTab[131819]++
										return nil, fmt.Errorf("file URLs must leave host empty or use localhost: got %v", u)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:131
		// _ = "end of CoverTab[131819]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:132
		_go_fuzz_dep_.CoverTab[131820]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:132
		// _ = "end of CoverTab[131820]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:132
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:132
	// _ = "end of CoverTab[131807]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:132
	_go_fuzz_dep_.CoverTab[131808]++
									switch u.Path {
	case "stdout":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:134
		_go_fuzz_dep_.CoverTab[131821]++
										return nopCloserSink{os.Stdout}, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:135
		// _ = "end of CoverTab[131821]"
	case "stderr":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:136
		_go_fuzz_dep_.CoverTab[131822]++
										return nopCloserSink{os.Stderr}, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:137
		// _ = "end of CoverTab[131822]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:137
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:137
		_go_fuzz_dep_.CoverTab[131823]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:137
		// _ = "end of CoverTab[131823]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:138
	// _ = "end of CoverTab[131808]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:138
	_go_fuzz_dep_.CoverTab[131809]++
									return os.OpenFile(u.Path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:139
	// _ = "end of CoverTab[131809]"
}

func normalizeScheme(s string) (string, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:142
	_go_fuzz_dep_.CoverTab[131824]++

									s = strings.ToLower(s)
									if first := s[0]; 'a' > first || func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:145
		_go_fuzz_dep_.CoverTab[131827]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:145
		return 'z' < first
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:145
		// _ = "end of CoverTab[131827]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:145
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:145
		_go_fuzz_dep_.CoverTab[131828]++
										return "", errors.New("must start with a letter")
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:146
		// _ = "end of CoverTab[131828]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:147
		_go_fuzz_dep_.CoverTab[131829]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:147
		// _ = "end of CoverTab[131829]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:147
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:147
	// _ = "end of CoverTab[131824]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:147
	_go_fuzz_dep_.CoverTab[131825]++
									for i := 1; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:148
		_go_fuzz_dep_.CoverTab[131830]++
										c := s[i]
										switch {
		case 'a' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:151
			_go_fuzz_dep_.CoverTab[131836]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:151
			return c <= 'z'
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:151
			// _ = "end of CoverTab[131836]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:151
		}():
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:151
			_go_fuzz_dep_.CoverTab[131832]++
											continue
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:152
			// _ = "end of CoverTab[131832]"
		case '0' <= c && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:153
			_go_fuzz_dep_.CoverTab[131837]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:153
			return c <= '9'
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:153
			// _ = "end of CoverTab[131837]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:153
		}():
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:153
			_go_fuzz_dep_.CoverTab[131833]++
											continue
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:154
			// _ = "end of CoverTab[131833]"
		case c == '.' || func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:155
			_go_fuzz_dep_.CoverTab[131838]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:155
			return c == '+'
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:155
			// _ = "end of CoverTab[131838]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:155
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:155
			_go_fuzz_dep_.CoverTab[131839]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:155
			return c == '-'
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:155
			// _ = "end of CoverTab[131839]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:155
		}():
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:155
			_go_fuzz_dep_.CoverTab[131834]++
											continue
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:156
			// _ = "end of CoverTab[131834]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:156
		default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:156
			_go_fuzz_dep_.CoverTab[131835]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:156
			// _ = "end of CoverTab[131835]"
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:157
		// _ = "end of CoverTab[131830]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:157
		_go_fuzz_dep_.CoverTab[131831]++
										return "", fmt.Errorf("may not contain %q", c)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:158
		// _ = "end of CoverTab[131831]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:159
	// _ = "end of CoverTab[131825]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:159
	_go_fuzz_dep_.CoverTab[131826]++
									return s, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:160
	// _ = "end of CoverTab[131826]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:161
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sink.go:161
var _ = _go_fuzz_dep_.CoverTab
