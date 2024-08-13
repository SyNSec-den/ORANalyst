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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:21
)

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap/zapcore"
)

// ServeHTTP is a simple JSON endpoint that can report on or change the current
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
// logging level.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
// # GET
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
// The GET request returns a JSON description of the current logging level like:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//	{"level":"info"}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
// # PUT
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
// The PUT request changes the logging level. It is perfectly safe to change the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
// logging level while a program is running. Two content types are supported:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//	Content-Type: application/x-www-form-urlencoded
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
// With this content type, the level can be provided through the request body or
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
// a query parameter. The log level is URL encoded like:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//	level=debug
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
// The request body takes precedence over the query parameter, if both are
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
// specified.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
// This content type is the default for a curl PUT request. Following are two
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
// example curl requests that both set the logging level to debug.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//	curl -X PUT localhost:8080/log/level?level=debug
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//	curl -X PUT localhost:8080/log/level -d level=debug
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
// For any other content type, the payload is expected to be JSON encoded and
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
// look like:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//	{"level":"info"}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
// An example curl request could look like this:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:32
//	curl -X PUT localhost:8080/log/level -H "Content-Type: application/json" -d '{"level":"debug"}'
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:70
func (lvl AtomicLevel) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:70
	_go_fuzz_dep_.CoverTab[131637]++
										type errorResponse struct {
		Error string `json:"error"`
	}
	type payload struct {
		Level zapcore.Level `json:"level"`
	}

	enc := json.NewEncoder(w)

	switch r.Method {
	case http.MethodGet:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:81
		_go_fuzz_dep_.CoverTab[131638]++
											enc.Encode(payload{Level: lvl.Level()})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:82
		// _ = "end of CoverTab[131638]"
	case http.MethodPut:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:83
		_go_fuzz_dep_.CoverTab[131639]++
											requestedLvl, err := decodePutRequest(r.Header.Get("Content-Type"), r)
											if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:85
			_go_fuzz_dep_.CoverTab[131642]++
												w.WriteHeader(http.StatusBadRequest)
												enc.Encode(errorResponse{Error: err.Error()})
												return
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:88
			// _ = "end of CoverTab[131642]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:89
			_go_fuzz_dep_.CoverTab[131643]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:89
			// _ = "end of CoverTab[131643]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:89
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:89
		// _ = "end of CoverTab[131639]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:89
		_go_fuzz_dep_.CoverTab[131640]++
											lvl.SetLevel(requestedLvl)
											enc.Encode(payload{Level: lvl.Level()})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:91
		// _ = "end of CoverTab[131640]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:92
		_go_fuzz_dep_.CoverTab[131641]++
											w.WriteHeader(http.StatusMethodNotAllowed)
											enc.Encode(errorResponse{
			Error: "Only GET and PUT are supported.",
		})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:96
		// _ = "end of CoverTab[131641]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:97
	// _ = "end of CoverTab[131637]"
}

// Decodes incoming PUT requests and returns the requested logging level.
func decodePutRequest(contentType string, r *http.Request) (zapcore.Level, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:101
	_go_fuzz_dep_.CoverTab[131644]++
										if contentType == "application/x-www-form-urlencoded" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:102
		_go_fuzz_dep_.CoverTab[131646]++
											return decodePutURL(r)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:103
		// _ = "end of CoverTab[131646]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:104
		_go_fuzz_dep_.CoverTab[131647]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:104
		// _ = "end of CoverTab[131647]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:104
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:104
	// _ = "end of CoverTab[131644]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:104
	_go_fuzz_dep_.CoverTab[131645]++
										return decodePutJSON(r.Body)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:105
	// _ = "end of CoverTab[131645]"
}

func decodePutURL(r *http.Request) (zapcore.Level, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:108
	_go_fuzz_dep_.CoverTab[131648]++
										lvl := r.FormValue("level")
										if lvl == "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:110
		_go_fuzz_dep_.CoverTab[131651]++
											return 0, fmt.Errorf("must specify logging level")
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:111
		// _ = "end of CoverTab[131651]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:112
		_go_fuzz_dep_.CoverTab[131652]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:112
		// _ = "end of CoverTab[131652]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:112
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:112
	// _ = "end of CoverTab[131648]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:112
	_go_fuzz_dep_.CoverTab[131649]++
										var l zapcore.Level
										if err := l.UnmarshalText([]byte(lvl)); err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:114
		_go_fuzz_dep_.CoverTab[131653]++
											return 0, err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:115
		// _ = "end of CoverTab[131653]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:116
		_go_fuzz_dep_.CoverTab[131654]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:116
		// _ = "end of CoverTab[131654]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:116
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:116
	// _ = "end of CoverTab[131649]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:116
	_go_fuzz_dep_.CoverTab[131650]++
										return l, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:117
	// _ = "end of CoverTab[131650]"
}

func decodePutJSON(body io.Reader) (zapcore.Level, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:120
	_go_fuzz_dep_.CoverTab[131655]++
										var pld struct {
		Level *zapcore.Level `json:"level"`
	}
	if err := json.NewDecoder(body).Decode(&pld); err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:124
		_go_fuzz_dep_.CoverTab[131658]++
											return 0, fmt.Errorf("malformed request body: %v", err)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:125
		// _ = "end of CoverTab[131658]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:126
		_go_fuzz_dep_.CoverTab[131659]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:126
		// _ = "end of CoverTab[131659]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:126
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:126
	// _ = "end of CoverTab[131655]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:126
	_go_fuzz_dep_.CoverTab[131656]++
										if pld.Level == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:127
		_go_fuzz_dep_.CoverTab[131660]++
											return 0, fmt.Errorf("must specify logging level")
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:128
		// _ = "end of CoverTab[131660]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:129
		_go_fuzz_dep_.CoverTab[131661]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:129
		// _ = "end of CoverTab[131661]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:129
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:129
	// _ = "end of CoverTab[131656]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:129
	_go_fuzz_dep_.CoverTab[131657]++
										return *pld.Level, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:130
	// _ = "end of CoverTab[131657]"

}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:132
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/http_handler.go:132
var _ = _go_fuzz_dep_.CoverTab
