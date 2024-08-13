// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:5
/*
Package trace implements tracing of requests and long-lived objects.
It exports HTTP interfaces on /debug/requests and /debug/events.

A trace.Trace provides tracing for short-lived objects, usually requests.
A request handler might be implemented like this:

	func fooHandler(w http.ResponseWriter, req *http.Request) {
		tr := trace.New("mypkg.Foo", req.URL.Path)
		defer tr.Finish()
		...
		tr.LazyPrintf("some event %q happened", str)
		...
		if err := somethingImportant(); err != nil {
			tr.LazyPrintf("somethingImportant failed: %v", err)
			tr.SetError()
		}
	}

The /debug/requests HTTP endpoint organizes the traces by family,
errors, and duration.  It also provides histogram of request duration
for each family.

A trace.EventLog provides tracing for long-lived objects, such as RPC
connections.

	// A Fetcher fetches URL paths for a single domain.
	type Fetcher struct {
		domain string
		events trace.EventLog
	}

	func NewFetcher(domain string) *Fetcher {
		return &Fetcher{
			domain,
			trace.NewEventLog("mypkg.Fetcher", domain),
		}
	}

	func (f *Fetcher) Fetch(path string) (string, error) {
		resp, err := http.Get("http://" + f.domain + "/" + path)
		if err != nil {
			f.events.Errorf("Get(%q) = %v", path, err)
			return "", err
		}
		f.events.Printf("Get(%q) = %s", path, resp.Status)
		...
	}

	func (f *Fetcher) Close() error {
		f.events.Finish()
		return nil
	}

The /debug/events HTTP endpoint organizes the event logs by family and
by time since the last error.  The expanded view displays recent log
entries and the log's call stack.
*/
package trace

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:63
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:63
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:63
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:63
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:63
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:63
)

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/internal/timeseries"
)

// DebugUseAfterFinish controls whether to debug uses of Trace values after finishing.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:85
// FOR DEBUGGING ONLY. This will slow down the program.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:87
var DebugUseAfterFinish = false

// HTTP ServeMux paths.
const (
	debugRequestsPath	= "/debug/requests"
	debugEventsPath		= "/debug/events"
)

// AuthRequest determines whether a specific request is permitted to load the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:95
// /debug/requests or /debug/events pages.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:95
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:95
// It returns two bools; the first indicates whether the page may be viewed at all,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:95
// and the second indicates whether sensitive events will be shown.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:95
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:95
// AuthRequest may be replaced by a program to customize its authorization requirements.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:95
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:95
// The default AuthRequest function returns (true, true) if and only if the request
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:95
// comes from localhost/127.0.0.1/[::1].
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:105
var AuthRequest = func(req *http.Request) (any, sensitive bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:105
	_go_fuzz_dep_.CoverTab[45448]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:108
	host, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:109
		_go_fuzz_dep_.CoverTab[45450]++
											host = req.RemoteAddr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:110
		// _ = "end of CoverTab[45450]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:111
		_go_fuzz_dep_.CoverTab[45451]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:111
		// _ = "end of CoverTab[45451]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:111
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:111
	// _ = "end of CoverTab[45448]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:111
	_go_fuzz_dep_.CoverTab[45449]++
										switch host {
	case "localhost", "127.0.0.1", "::1":
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:113
		_go_fuzz_dep_.CoverTab[45452]++
											return true, true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:114
		// _ = "end of CoverTab[45452]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:115
		_go_fuzz_dep_.CoverTab[45453]++
											return false, false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:116
		// _ = "end of CoverTab[45453]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:117
	// _ = "end of CoverTab[45449]"
}

func init() {
	_, pat := http.DefaultServeMux.Handler(&http.Request{URL: &url.URL{Path: debugRequestsPath}})
	if pat == debugRequestsPath {
		panic("/debug/requests is already registered. You may have two independent copies of " +
			"golang.org/x/net/trace in your binary, trying to maintain separate state. This may " +
			"involve a vendored copy of golang.org/x/net/trace.")
	}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:130
	http.HandleFunc(debugRequestsPath, Traces)
	http.HandleFunc(debugEventsPath, Events)
}

// NewContext returns a copy of the parent context
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:134
// and associates it with a Trace.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:136
func NewContext(ctx context.Context, tr Trace) context.Context {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:136
	_go_fuzz_dep_.CoverTab[45454]++
										return context.WithValue(ctx, contextKey, tr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:137
	// _ = "end of CoverTab[45454]"
}

// FromContext returns the Trace bound to the context, if any.
func FromContext(ctx context.Context) (tr Trace, ok bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:141
	_go_fuzz_dep_.CoverTab[45455]++
										tr, ok = ctx.Value(contextKey).(Trace)
										return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:143
	// _ = "end of CoverTab[45455]"
}

// Traces responds with traces from the program.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:146
// The package initialization registers it in http.DefaultServeMux
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:146
// at /debug/requests.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:146
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:146
// It performs authorization by running AuthRequest.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:151
func Traces(w http.ResponseWriter, req *http.Request) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:151
	_go_fuzz_dep_.CoverTab[45456]++
										any, sensitive := AuthRequest(req)
										if !any {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:153
		_go_fuzz_dep_.CoverTab[45458]++
											http.Error(w, "not allowed", http.StatusUnauthorized)
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:155
		// _ = "end of CoverTab[45458]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:156
		_go_fuzz_dep_.CoverTab[45459]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:156
		// _ = "end of CoverTab[45459]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:156
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:156
	// _ = "end of CoverTab[45456]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:156
	_go_fuzz_dep_.CoverTab[45457]++
										w.Header().Set("Content-Type", "text/html; charset=utf-8")
										Render(w, req, sensitive)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:158
	// _ = "end of CoverTab[45457]"
}

// Events responds with a page of events collected by EventLogs.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:161
// The package initialization registers it in http.DefaultServeMux
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:161
// at /debug/events.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:161
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:161
// It performs authorization by running AuthRequest.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:166
func Events(w http.ResponseWriter, req *http.Request) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:166
	_go_fuzz_dep_.CoverTab[45460]++
										any, sensitive := AuthRequest(req)
										if !any {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:168
		_go_fuzz_dep_.CoverTab[45462]++
											http.Error(w, "not allowed", http.StatusUnauthorized)
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:170
		// _ = "end of CoverTab[45462]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:171
		_go_fuzz_dep_.CoverTab[45463]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:171
		// _ = "end of CoverTab[45463]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:171
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:171
	// _ = "end of CoverTab[45460]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:171
	_go_fuzz_dep_.CoverTab[45461]++
										w.Header().Set("Content-Type", "text/html; charset=utf-8")
										RenderEvents(w, req, sensitive)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:173
	// _ = "end of CoverTab[45461]"
}

// Render renders the HTML page typically served at /debug/requests.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:176
// It does not do any auth checking. The request may be nil.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:176
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:176
// Most users will use the Traces handler.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:180
func Render(w io.Writer, req *http.Request, sensitive bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:180
	_go_fuzz_dep_.CoverTab[45464]++
										data := &struct {
		Families		[]string
		ActiveTraceCount	map[string]int
		CompletedTraces		map[string]*family

		// Set when a bucket has been selected.
		Traces		traceList
		Family		string
		Bucket		int
		Expanded	bool
		Traced		bool
		Active		bool
		ShowSensitive	bool	// whether to show sensitive events

		Histogram	template.HTML
		HistogramWindow	string	// e.g. "last minute", "last hour", "all time"

		// If non-zero, the set of traces is a partial set,
		// and this is the total number.
		Total	int
	}{
		CompletedTraces: completedTraces,
	}

	data.ShowSensitive = sensitive
	if req != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:206
		_go_fuzz_dep_.CoverTab[45470]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:209
		if req.FormValue("show_sensitive") == "0" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:209
			_go_fuzz_dep_.CoverTab[45473]++
												data.ShowSensitive = false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:210
			// _ = "end of CoverTab[45473]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:211
			_go_fuzz_dep_.CoverTab[45474]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:211
			// _ = "end of CoverTab[45474]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:211
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:211
		// _ = "end of CoverTab[45470]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:211
		_go_fuzz_dep_.CoverTab[45471]++

											if exp, err := strconv.ParseBool(req.FormValue("exp")); err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:213
			_go_fuzz_dep_.CoverTab[45475]++
												data.Expanded = exp
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:214
			// _ = "end of CoverTab[45475]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:215
			_go_fuzz_dep_.CoverTab[45476]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:215
			// _ = "end of CoverTab[45476]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:215
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:215
		// _ = "end of CoverTab[45471]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:215
		_go_fuzz_dep_.CoverTab[45472]++
											if exp, err := strconv.ParseBool(req.FormValue("rtraced")); err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:216
			_go_fuzz_dep_.CoverTab[45477]++
												data.Traced = exp
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:217
			// _ = "end of CoverTab[45477]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:218
			_go_fuzz_dep_.CoverTab[45478]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:218
			// _ = "end of CoverTab[45478]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:218
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:218
		// _ = "end of CoverTab[45472]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:219
		_go_fuzz_dep_.CoverTab[45479]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:219
		// _ = "end of CoverTab[45479]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:219
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:219
	// _ = "end of CoverTab[45464]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:219
	_go_fuzz_dep_.CoverTab[45465]++

										completedMu.RLock()
										data.Families = make([]string, 0, len(completedTraces))
										for fam := range completedTraces {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:223
		_go_fuzz_dep_.CoverTab[45480]++
											data.Families = append(data.Families, fam)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:224
		// _ = "end of CoverTab[45480]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:225
	// _ = "end of CoverTab[45465]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:225
	_go_fuzz_dep_.CoverTab[45466]++
										completedMu.RUnlock()
										sort.Strings(data.Families)

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:231
	data.ActiveTraceCount = make(map[string]int, len(data.Families))
	activeMu.RLock()
	for fam, s := range activeTraces {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:233
		_go_fuzz_dep_.CoverTab[45481]++
											data.ActiveTraceCount[fam] = s.Len()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:234
		// _ = "end of CoverTab[45481]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:235
	// _ = "end of CoverTab[45466]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:235
	_go_fuzz_dep_.CoverTab[45467]++
										activeMu.RUnlock()

										var ok bool
										data.Family, data.Bucket, ok = parseArgs(req)
										switch {
	case !ok:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:241
		_go_fuzz_dep_.CoverTab[45482]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:241
		// _ = "end of CoverTab[45482]"

	case data.Bucket == -1:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:243
		_go_fuzz_dep_.CoverTab[45483]++
											data.Active = true
											n := data.ActiveTraceCount[data.Family]
											data.Traces = getActiveTraces(data.Family)
											if len(data.Traces) < n {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:247
			_go_fuzz_dep_.CoverTab[45486]++
												data.Total = n
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:248
			// _ = "end of CoverTab[45486]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:249
			_go_fuzz_dep_.CoverTab[45487]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:249
			// _ = "end of CoverTab[45487]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:249
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:249
		// _ = "end of CoverTab[45483]"
	case data.Bucket < bucketsPerFamily:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:250
		_go_fuzz_dep_.CoverTab[45484]++
											if b := lookupBucket(data.Family, data.Bucket); b != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:251
			_go_fuzz_dep_.CoverTab[45488]++
												data.Traces = b.Copy(data.Traced)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:252
			// _ = "end of CoverTab[45488]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:253
			_go_fuzz_dep_.CoverTab[45489]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:253
			// _ = "end of CoverTab[45489]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:253
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:253
		// _ = "end of CoverTab[45484]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:254
		_go_fuzz_dep_.CoverTab[45485]++
											if f := getFamily(data.Family, false); f != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:255
			_go_fuzz_dep_.CoverTab[45490]++
												var obs timeseries.Observable
												f.LatencyMu.RLock()
												switch o := data.Bucket - bucketsPerFamily; o {
			case 0:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:259
				_go_fuzz_dep_.CoverTab[45492]++
													obs = f.Latency.Minute()
													data.HistogramWindow = "last minute"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:261
				// _ = "end of CoverTab[45492]"
			case 1:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:262
				_go_fuzz_dep_.CoverTab[45493]++
													obs = f.Latency.Hour()
													data.HistogramWindow = "last hour"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:264
				// _ = "end of CoverTab[45493]"
			case 2:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:265
				_go_fuzz_dep_.CoverTab[45494]++
													obs = f.Latency.Total()
													data.HistogramWindow = "all time"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:267
				// _ = "end of CoverTab[45494]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:267
			default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:267
				_go_fuzz_dep_.CoverTab[45495]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:267
				// _ = "end of CoverTab[45495]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:268
			// _ = "end of CoverTab[45490]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:268
			_go_fuzz_dep_.CoverTab[45491]++
												f.LatencyMu.RUnlock()
												if obs != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:270
				_go_fuzz_dep_.CoverTab[45496]++
													data.Histogram = obs.(*histogram).html()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:271
				// _ = "end of CoverTab[45496]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:272
				_go_fuzz_dep_.CoverTab[45497]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:272
				// _ = "end of CoverTab[45497]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:272
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:272
			// _ = "end of CoverTab[45491]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:273
			_go_fuzz_dep_.CoverTab[45498]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:273
			// _ = "end of CoverTab[45498]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:273
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:273
		// _ = "end of CoverTab[45485]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:274
	// _ = "end of CoverTab[45467]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:274
	_go_fuzz_dep_.CoverTab[45468]++

										if data.Traces != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:276
		_go_fuzz_dep_.CoverTab[45499]++
											defer data.Traces.Free()
											sort.Sort(data.Traces)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:278
		// _ = "end of CoverTab[45499]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:279
		_go_fuzz_dep_.CoverTab[45500]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:279
		// _ = "end of CoverTab[45500]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:279
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:279
	// _ = "end of CoverTab[45468]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:279
	_go_fuzz_dep_.CoverTab[45469]++

										completedMu.RLock()
										defer completedMu.RUnlock()
										if err := pageTmpl().ExecuteTemplate(w, "Page", data); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:283
		_go_fuzz_dep_.CoverTab[45501]++
											log.Printf("net/trace: Failed executing template: %v", err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:284
		// _ = "end of CoverTab[45501]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:285
		_go_fuzz_dep_.CoverTab[45502]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:285
		// _ = "end of CoverTab[45502]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:285
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:285
	// _ = "end of CoverTab[45469]"
}

func parseArgs(req *http.Request) (fam string, b int, ok bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:288
	_go_fuzz_dep_.CoverTab[45503]++
										if req == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:289
		_go_fuzz_dep_.CoverTab[45507]++
											return "", 0, false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:290
		// _ = "end of CoverTab[45507]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:291
		_go_fuzz_dep_.CoverTab[45508]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:291
		// _ = "end of CoverTab[45508]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:291
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:291
	// _ = "end of CoverTab[45503]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:291
	_go_fuzz_dep_.CoverTab[45504]++
										fam, bStr := req.FormValue("fam"), req.FormValue("b")
										if fam == "" || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:293
		_go_fuzz_dep_.CoverTab[45509]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:293
		return bStr == ""
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:293
		// _ = "end of CoverTab[45509]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:293
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:293
		_go_fuzz_dep_.CoverTab[45510]++
											return "", 0, false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:294
		// _ = "end of CoverTab[45510]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:295
		_go_fuzz_dep_.CoverTab[45511]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:295
		// _ = "end of CoverTab[45511]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:295
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:295
	// _ = "end of CoverTab[45504]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:295
	_go_fuzz_dep_.CoverTab[45505]++
										b, err := strconv.Atoi(bStr)
										if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:297
		_go_fuzz_dep_.CoverTab[45512]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:297
		return b < -1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:297
		// _ = "end of CoverTab[45512]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:297
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:297
		_go_fuzz_dep_.CoverTab[45513]++
											return "", 0, false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:298
		// _ = "end of CoverTab[45513]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:299
		_go_fuzz_dep_.CoverTab[45514]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:299
		// _ = "end of CoverTab[45514]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:299
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:299
	// _ = "end of CoverTab[45505]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:299
	_go_fuzz_dep_.CoverTab[45506]++

										return fam, b, true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:301
	// _ = "end of CoverTab[45506]"
}

func lookupBucket(fam string, b int) *traceBucket {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:304
	_go_fuzz_dep_.CoverTab[45515]++
										f := getFamily(fam, false)
										if f == nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:306
		_go_fuzz_dep_.CoverTab[45517]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:306
		return b < 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:306
		// _ = "end of CoverTab[45517]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:306
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:306
		_go_fuzz_dep_.CoverTab[45518]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:306
		return b >= len(f.Buckets)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:306
		// _ = "end of CoverTab[45518]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:306
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:306
		_go_fuzz_dep_.CoverTab[45519]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:307
		// _ = "end of CoverTab[45519]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:308
		_go_fuzz_dep_.CoverTab[45520]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:308
		// _ = "end of CoverTab[45520]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:308
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:308
	// _ = "end of CoverTab[45515]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:308
	_go_fuzz_dep_.CoverTab[45516]++
										return f.Buckets[b]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:309
	// _ = "end of CoverTab[45516]"
}

type contextKeyT string

var contextKey = contextKeyT("golang.org/x/net/trace.Trace")

// Trace represents an active request.
type Trace interface {
	// LazyLog adds x to the event log. It will be evaluated each time the
	// /debug/requests page is rendered. Any memory referenced by x will be
	// pinned until the trace is finished and later discarded.
	LazyLog(x fmt.Stringer, sensitive bool)

	// LazyPrintf evaluates its arguments with fmt.Sprintf each time the
	// /debug/requests page is rendered. Any memory referenced by a will be
	// pinned until the trace is finished and later discarded.
	LazyPrintf(format string, a ...interface{})

	// SetError declares that this trace resulted in an error.
	SetError()

	// SetRecycler sets a recycler for the trace.
	// f will be called for each event passed to LazyLog at a time when
	// it is no longer required, whether while the trace is still active
	// and the event is discarded, or when a completed trace is discarded.
	SetRecycler(f func(interface{}))

	// SetTraceInfo sets the trace info for the trace.
	// This is currently unused.
	SetTraceInfo(traceID, spanID uint64)

	// SetMaxEvents sets the maximum number of events that will be stored
	// in the trace. This has no effect if any events have already been
	// added to the trace.
	SetMaxEvents(m int)

	// Finish declares that this trace is complete.
	// The trace should not be used after calling this method.
	Finish()
}

type lazySprintf struct {
	format	string
	a	[]interface{}
}

func (l *lazySprintf) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:356
	_go_fuzz_dep_.CoverTab[45521]++
										return fmt.Sprintf(l.format, l.a...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:357
	// _ = "end of CoverTab[45521]"
}

// New returns a new Trace with the specified family and title.
func New(family, title string) Trace {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:361
	_go_fuzz_dep_.CoverTab[45522]++
										tr := newTrace()
										tr.ref()
										tr.Family, tr.Title = family, title
										tr.Start = time.Now()
										tr.maxEvents = maxEventsPerTrace
										tr.events = tr.eventsBuf[:0]

										activeMu.RLock()
										s := activeTraces[tr.Family]
										activeMu.RUnlock()
										if s == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:372
		_go_fuzz_dep_.CoverTab[45525]++
											activeMu.Lock()
											s = activeTraces[tr.Family]
											if s == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:375
			_go_fuzz_dep_.CoverTab[45527]++
												s = new(traceSet)
												activeTraces[tr.Family] = s
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:377
			// _ = "end of CoverTab[45527]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:378
			_go_fuzz_dep_.CoverTab[45528]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:378
			// _ = "end of CoverTab[45528]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:378
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:378
		// _ = "end of CoverTab[45525]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:378
		_go_fuzz_dep_.CoverTab[45526]++
											activeMu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:379
		// _ = "end of CoverTab[45526]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:380
		_go_fuzz_dep_.CoverTab[45529]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:380
		// _ = "end of CoverTab[45529]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:380
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:380
	// _ = "end of CoverTab[45522]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:380
	_go_fuzz_dep_.CoverTab[45523]++
										s.Add(tr)

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:388
	completedMu.RLock()
	if _, ok := completedTraces[tr.Family]; !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:389
		_go_fuzz_dep_.CoverTab[45530]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:389
		_curRoutineNum45_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:389
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum45_)
											go func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:390
			_go_fuzz_dep_.CoverTab[45531]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:390
			defer func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:390
				_go_fuzz_dep_.CoverTab[45532]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:390
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum45_)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:390
				// _ = "end of CoverTab[45532]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:390
			}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:390
			allocFamily(tr.Family)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:390
			// _ = "end of CoverTab[45531]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:390
		}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:390
		// _ = "end of CoverTab[45530]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:391
		_go_fuzz_dep_.CoverTab[45533]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:391
		// _ = "end of CoverTab[45533]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:391
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:391
	// _ = "end of CoverTab[45523]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:391
	_go_fuzz_dep_.CoverTab[45524]++
										completedMu.RUnlock()

										return tr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:394
	// _ = "end of CoverTab[45524]"
}

func (tr *trace) Finish() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:397
	_go_fuzz_dep_.CoverTab[45534]++
										elapsed := time.Since(tr.Start)
										tr.mu.Lock()
										tr.Elapsed = elapsed
										tr.mu.Unlock()

										if DebugUseAfterFinish {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:403
		_go_fuzz_dep_.CoverTab[45537]++
											buf := make([]byte, 4<<10)
											n := runtime.Stack(buf, false)
											tr.finishStack = buf[:n]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:406
		// _ = "end of CoverTab[45537]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:407
		_go_fuzz_dep_.CoverTab[45538]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:407
		// _ = "end of CoverTab[45538]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:407
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:407
	// _ = "end of CoverTab[45534]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:407
	_go_fuzz_dep_.CoverTab[45535]++

										activeMu.RLock()
										m := activeTraces[tr.Family]
										activeMu.RUnlock()
										m.Remove(tr)

										f := getFamily(tr.Family, true)
										tr.mu.RLock()
										for _, b := range f.Buckets {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:416
		_go_fuzz_dep_.CoverTab[45539]++
											if b.Cond.match(tr) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:417
			_go_fuzz_dep_.CoverTab[45540]++
												b.Add(tr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:418
			// _ = "end of CoverTab[45540]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:419
			_go_fuzz_dep_.CoverTab[45541]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:419
			// _ = "end of CoverTab[45541]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:419
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:419
		// _ = "end of CoverTab[45539]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:420
	// _ = "end of CoverTab[45535]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:420
	_go_fuzz_dep_.CoverTab[45536]++
										tr.mu.RUnlock()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:424
	h := new(histogram)
										h.addMeasurement(elapsed.Nanoseconds() / 1e3)
										f.LatencyMu.Lock()
										f.Latency.Add(h)
										f.LatencyMu.Unlock()

										tr.unref()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:430
	// _ = "end of CoverTab[45536]"
}

const (
	bucketsPerFamily	= 9
	tracesPerBucket		= 10
	maxActiveTraces		= 20	// Maximum number of active traces to show.
	maxEventsPerTrace	= 10
	numHistogramBuckets	= 38
)

var (
	// The active traces.
	activeMu	sync.RWMutex
	activeTraces	= make(map[string]*traceSet)	// family -> traces

	// Families of completed traces.
	completedMu	sync.RWMutex
	completedTraces	= make(map[string]*family)	// family -> traces
)

type traceSet struct {
	mu	sync.RWMutex
	m	map[*trace]bool
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:460
}

func (ts *traceSet) Len() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:462
	_go_fuzz_dep_.CoverTab[45542]++
										ts.mu.RLock()
										defer ts.mu.RUnlock()
										return len(ts.m)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:465
	// _ = "end of CoverTab[45542]"
}

func (ts *traceSet) Add(tr *trace) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:468
	_go_fuzz_dep_.CoverTab[45543]++
										ts.mu.Lock()
										if ts.m == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:470
		_go_fuzz_dep_.CoverTab[45545]++
											ts.m = make(map[*trace]bool)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:471
		// _ = "end of CoverTab[45545]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:472
		_go_fuzz_dep_.CoverTab[45546]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:472
		// _ = "end of CoverTab[45546]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:472
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:472
	// _ = "end of CoverTab[45543]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:472
	_go_fuzz_dep_.CoverTab[45544]++
										ts.m[tr] = true
										ts.mu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:474
	// _ = "end of CoverTab[45544]"
}

func (ts *traceSet) Remove(tr *trace) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:477
	_go_fuzz_dep_.CoverTab[45547]++
										ts.mu.Lock()
										delete(ts.m, tr)
										ts.mu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:480
	// _ = "end of CoverTab[45547]"
}

// FirstN returns the first n traces ordered by time.
func (ts *traceSet) FirstN(n int) traceList {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:484
	_go_fuzz_dep_.CoverTab[45548]++
										ts.mu.RLock()
										defer ts.mu.RUnlock()

										if n > len(ts.m) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:488
		_go_fuzz_dep_.CoverTab[45552]++
											n = len(ts.m)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:489
		// _ = "end of CoverTab[45552]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:490
		_go_fuzz_dep_.CoverTab[45553]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:490
		// _ = "end of CoverTab[45553]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:490
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:490
	// _ = "end of CoverTab[45548]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:490
	_go_fuzz_dep_.CoverTab[45549]++
										trl := make(traceList, 0, n)

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:494
	if n == len(ts.m) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:494
		_go_fuzz_dep_.CoverTab[45554]++
											for tr := range ts.m {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:495
			_go_fuzz_dep_.CoverTab[45556]++
												tr.ref()
												trl = append(trl, tr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:497
			// _ = "end of CoverTab[45556]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:498
		// _ = "end of CoverTab[45554]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:498
		_go_fuzz_dep_.CoverTab[45555]++
											sort.Sort(trl)
											return trl
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:500
		// _ = "end of CoverTab[45555]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:501
		_go_fuzz_dep_.CoverTab[45557]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:501
		// _ = "end of CoverTab[45557]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:501
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:501
	// _ = "end of CoverTab[45549]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:501
	_go_fuzz_dep_.CoverTab[45550]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:505
	for tr := range ts.m {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:505
		_go_fuzz_dep_.CoverTab[45558]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:508
		if len(trl) < n {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:508
			_go_fuzz_dep_.CoverTab[45562]++
												tr.ref()
												trl = append(trl, tr)
												if len(trl) == n {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:511
				_go_fuzz_dep_.CoverTab[45564]++

													sort.Sort(trl)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:513
				// _ = "end of CoverTab[45564]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:514
				_go_fuzz_dep_.CoverTab[45565]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:514
				// _ = "end of CoverTab[45565]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:514
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:514
			// _ = "end of CoverTab[45562]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:514
			_go_fuzz_dep_.CoverTab[45563]++
												continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:515
			// _ = "end of CoverTab[45563]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:516
			_go_fuzz_dep_.CoverTab[45566]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:516
			// _ = "end of CoverTab[45566]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:516
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:516
		// _ = "end of CoverTab[45558]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:516
		_go_fuzz_dep_.CoverTab[45559]++
											if tr.Start.After(trl[n-1].Start) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:517
			_go_fuzz_dep_.CoverTab[45567]++
												continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:518
			// _ = "end of CoverTab[45567]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:519
			_go_fuzz_dep_.CoverTab[45568]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:519
			// _ = "end of CoverTab[45568]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:519
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:519
		// _ = "end of CoverTab[45559]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:519
		_go_fuzz_dep_.CoverTab[45560]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:522
		tr.ref()
		i := sort.Search(n, func(i int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:523
			_go_fuzz_dep_.CoverTab[45569]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:523
			return trl[i].Start.After(tr.Start)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:523
			// _ = "end of CoverTab[45569]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:523
		})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:523
		// _ = "end of CoverTab[45560]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:523
		_go_fuzz_dep_.CoverTab[45561]++
											trl[n-1].unref()
											copy(trl[i+1:], trl[i:])
											trl[i] = tr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:526
		// _ = "end of CoverTab[45561]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:527
	// _ = "end of CoverTab[45550]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:527
	_go_fuzz_dep_.CoverTab[45551]++

										return trl
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:529
	// _ = "end of CoverTab[45551]"
}

func getActiveTraces(fam string) traceList {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:532
	_go_fuzz_dep_.CoverTab[45570]++
										activeMu.RLock()
										s := activeTraces[fam]
										activeMu.RUnlock()
										if s == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:536
		_go_fuzz_dep_.CoverTab[45572]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:537
		// _ = "end of CoverTab[45572]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:538
		_go_fuzz_dep_.CoverTab[45573]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:538
		// _ = "end of CoverTab[45573]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:538
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:538
	// _ = "end of CoverTab[45570]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:538
	_go_fuzz_dep_.CoverTab[45571]++
										return s.FirstN(maxActiveTraces)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:539
	// _ = "end of CoverTab[45571]"
}

func getFamily(fam string, allocNew bool) *family {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:542
	_go_fuzz_dep_.CoverTab[45574]++
										completedMu.RLock()
										f := completedTraces[fam]
										completedMu.RUnlock()
										if f == nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:546
		_go_fuzz_dep_.CoverTab[45576]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:546
		return allocNew
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:546
		// _ = "end of CoverTab[45576]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:546
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:546
		_go_fuzz_dep_.CoverTab[45577]++
											f = allocFamily(fam)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:547
		// _ = "end of CoverTab[45577]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:548
		_go_fuzz_dep_.CoverTab[45578]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:548
		// _ = "end of CoverTab[45578]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:548
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:548
	// _ = "end of CoverTab[45574]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:548
	_go_fuzz_dep_.CoverTab[45575]++
										return f
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:549
	// _ = "end of CoverTab[45575]"
}

func allocFamily(fam string) *family {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:552
	_go_fuzz_dep_.CoverTab[45579]++
										completedMu.Lock()
										defer completedMu.Unlock()
										f := completedTraces[fam]
										if f == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:556
		_go_fuzz_dep_.CoverTab[45581]++
											f = newFamily()
											completedTraces[fam] = f
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:558
		// _ = "end of CoverTab[45581]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:559
		_go_fuzz_dep_.CoverTab[45582]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:559
		// _ = "end of CoverTab[45582]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:559
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:559
	// _ = "end of CoverTab[45579]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:559
	_go_fuzz_dep_.CoverTab[45580]++
										return f
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:560
	// _ = "end of CoverTab[45580]"
}

// family represents a set of trace buckets and associated latency information.
type family struct {
	// traces may occur in multiple buckets.
	Buckets	[bucketsPerFamily]*traceBucket

	// latency time series
	LatencyMu	sync.RWMutex
	Latency		*timeseries.MinuteHourSeries
}

func newFamily() *family {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:573
	_go_fuzz_dep_.CoverTab[45583]++
										return &family{
		Buckets: [bucketsPerFamily]*traceBucket{
			{Cond: minCond(0)},
			{Cond: minCond(50 * time.Millisecond)},
			{Cond: minCond(100 * time.Millisecond)},
			{Cond: minCond(200 * time.Millisecond)},
			{Cond: minCond(500 * time.Millisecond)},
			{Cond: minCond(1 * time.Second)},
			{Cond: minCond(10 * time.Second)},
			{Cond: minCond(100 * time.Second)},
			{Cond: errorCond{}},
		},
		Latency: timeseries.NewMinuteHourSeries(func() timeseries.Observable {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:586
			_go_fuzz_dep_.CoverTab[45584]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:586
			return new(histogram)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:586
			// _ = "end of CoverTab[45584]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:586
		}),
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:587
	// _ = "end of CoverTab[45583]"
}

// traceBucket represents a size-capped bucket of historic traces,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:590
// along with a condition for a trace to belong to the bucket.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:592
type traceBucket struct {
	Cond	cond

	// Ring buffer implementation of a fixed-size FIFO queue.
	mu	sync.RWMutex
	buf	[tracesPerBucket]*trace
	start	int	// < tracesPerBucket
	length	int	// <= tracesPerBucket
}

func (b *traceBucket) Add(tr *trace) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:602
	_go_fuzz_dep_.CoverTab[45585]++
										b.mu.Lock()
										defer b.mu.Unlock()

										i := b.start + b.length
										if i >= tracesPerBucket {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:607
		_go_fuzz_dep_.CoverTab[45589]++
											i -= tracesPerBucket
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:608
		// _ = "end of CoverTab[45589]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:609
		_go_fuzz_dep_.CoverTab[45590]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:609
		// _ = "end of CoverTab[45590]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:609
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:609
	// _ = "end of CoverTab[45585]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:609
	_go_fuzz_dep_.CoverTab[45586]++
										if b.length == tracesPerBucket {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:610
		_go_fuzz_dep_.CoverTab[45591]++

											b.buf[i].unref()
											b.start++
											if b.start == tracesPerBucket {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:614
			_go_fuzz_dep_.CoverTab[45592]++
												b.start = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:615
			// _ = "end of CoverTab[45592]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:616
			_go_fuzz_dep_.CoverTab[45593]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:616
			// _ = "end of CoverTab[45593]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:616
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:616
		// _ = "end of CoverTab[45591]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:617
		_go_fuzz_dep_.CoverTab[45594]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:617
		// _ = "end of CoverTab[45594]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:617
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:617
	// _ = "end of CoverTab[45586]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:617
	_go_fuzz_dep_.CoverTab[45587]++
										b.buf[i] = tr
										if b.length < tracesPerBucket {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:619
		_go_fuzz_dep_.CoverTab[45595]++
											b.length++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:620
		// _ = "end of CoverTab[45595]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:621
		_go_fuzz_dep_.CoverTab[45596]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:621
		// _ = "end of CoverTab[45596]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:621
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:621
	// _ = "end of CoverTab[45587]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:621
	_go_fuzz_dep_.CoverTab[45588]++
										tr.ref()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:622
	// _ = "end of CoverTab[45588]"
}

// Copy returns a copy of the traces in the bucket.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:625
// If tracedOnly is true, only the traces with trace information will be returned.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:625
// The logs will be ref'd before returning; the caller should call
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:625
// the Free method when it is done with them.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:625
// TODO(dsymonds): keep track of traced requests in separate buckets.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:630
func (b *traceBucket) Copy(tracedOnly bool) traceList {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:630
	_go_fuzz_dep_.CoverTab[45597]++
										b.mu.RLock()
										defer b.mu.RUnlock()

										trl := make(traceList, 0, b.length)
										for i, x := 0, b.start; i < b.length; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:635
		_go_fuzz_dep_.CoverTab[45599]++
											tr := b.buf[x]
											if !tracedOnly || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:637
			_go_fuzz_dep_.CoverTab[45601]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:637
			return tr.spanID != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:637
			// _ = "end of CoverTab[45601]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:637
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:637
			_go_fuzz_dep_.CoverTab[45602]++
												tr.ref()
												trl = append(trl, tr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:639
			// _ = "end of CoverTab[45602]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:640
			_go_fuzz_dep_.CoverTab[45603]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:640
			// _ = "end of CoverTab[45603]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:640
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:640
		// _ = "end of CoverTab[45599]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:640
		_go_fuzz_dep_.CoverTab[45600]++
											x++
											if x == b.length {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:642
			_go_fuzz_dep_.CoverTab[45604]++
												x = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:643
			// _ = "end of CoverTab[45604]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:644
			_go_fuzz_dep_.CoverTab[45605]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:644
			// _ = "end of CoverTab[45605]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:644
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:644
		// _ = "end of CoverTab[45600]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:645
	// _ = "end of CoverTab[45597]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:645
	_go_fuzz_dep_.CoverTab[45598]++
										return trl
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:646
	// _ = "end of CoverTab[45598]"
}

func (b *traceBucket) Empty() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:649
	_go_fuzz_dep_.CoverTab[45606]++
										b.mu.RLock()
										defer b.mu.RUnlock()
										return b.length == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:652
	// _ = "end of CoverTab[45606]"
}

// cond represents a condition on a trace.
type cond interface {
	match(t *trace) bool
	String() string
}

type minCond time.Duration

func (m minCond) match(t *trace) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:663
	_go_fuzz_dep_.CoverTab[45607]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:663
	return t.Elapsed >= time.Duration(m)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:663
	// _ = "end of CoverTab[45607]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:663
}
func (m minCond) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:664
	_go_fuzz_dep_.CoverTab[45608]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:664
	return fmt.Sprintf("≥%gs", time.Duration(m).Seconds())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:664
	// _ = "end of CoverTab[45608]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:664
}

type errorCond struct{}

func (e errorCond) match(t *trace) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:668
	_go_fuzz_dep_.CoverTab[45609]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:668
	return t.IsError
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:668
	// _ = "end of CoverTab[45609]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:668
}
func (e errorCond) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:669
	_go_fuzz_dep_.CoverTab[45610]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:669
	return "errors"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:669
	// _ = "end of CoverTab[45610]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:669
}

type traceList []*trace

// Free calls unref on each element of the list.
func (trl traceList) Free() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:674
	_go_fuzz_dep_.CoverTab[45611]++
										for _, t := range trl {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:675
		_go_fuzz_dep_.CoverTab[45612]++
											t.unref()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:676
		// _ = "end of CoverTab[45612]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:677
	// _ = "end of CoverTab[45611]"
}

// traceList may be sorted in reverse chronological order.
func (trl traceList) Len() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:681
	_go_fuzz_dep_.CoverTab[45613]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:681
	return len(trl)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:681
	// _ = "end of CoverTab[45613]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:681
}
func (trl traceList) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:682
	_go_fuzz_dep_.CoverTab[45614]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:682
	return trl[i].Start.After(trl[j].Start)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:682
	// _ = "end of CoverTab[45614]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:682
}
func (trl traceList) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:683
	_go_fuzz_dep_.CoverTab[45615]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:683
	trl[i], trl[j] = trl[j], trl[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:683
	// _ = "end of CoverTab[45615]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:683
}

// An event is a timestamped log entry in a trace.
type event struct {
	When		time.Time
	Elapsed		time.Duration	// since previous event in trace
	NewDay		bool		// whether this event is on a different day to the previous event
	Recyclable	bool		// whether this event was passed via LazyLog
	Sensitive	bool		// whether this event contains sensitive information
	What		interface{}	// string or fmt.Stringer
}

// WhenString returns a string representation of the elapsed time of the event.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:695
// It will include the date if midnight was crossed.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:697
func (e event) WhenString() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:697
	_go_fuzz_dep_.CoverTab[45616]++
										if e.NewDay {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:698
		_go_fuzz_dep_.CoverTab[45618]++
											return e.When.Format("2006/01/02 15:04:05.000000")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:699
		// _ = "end of CoverTab[45618]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:700
		_go_fuzz_dep_.CoverTab[45619]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:700
		// _ = "end of CoverTab[45619]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:700
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:700
	// _ = "end of CoverTab[45616]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:700
	_go_fuzz_dep_.CoverTab[45617]++
										return e.When.Format("15:04:05.000000")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:701
	// _ = "end of CoverTab[45617]"
}

// discarded represents a number of discarded events.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:704
// It is stored as *discarded to make it easier to update in-place.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:706
type discarded int

func (d *discarded) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:708
	_go_fuzz_dep_.CoverTab[45620]++
										return fmt.Sprintf("(%d events discarded)", int(*d))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:709
	// _ = "end of CoverTab[45620]"
}

// trace represents an active or complete request,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:712
// either sent or received by this program.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:714
type trace struct {
	// Family is the top-level grouping of traces to which this belongs.
	Family	string

	// Title is the title of this trace.
	Title	string

	// Start time of the this trace.
	Start	time.Time

	mu		sync.RWMutex
	events		[]event	// Append-only sequence of events (modulo discards).
	maxEvents	int
	recycler	func(interface{})
	IsError		bool		// Whether this trace resulted in an error.
	Elapsed		time.Duration	// Elapsed time for this trace, zero while active.
	traceID		uint64		// Trace information if non-zero.
	spanID		uint64

	refs	int32		// how many buckets this is in
	disc	discarded	// scratch space to avoid allocation

	finishStack	[]byte	// where finish was called, if DebugUseAfterFinish is set

	eventsBuf	[4]event	// preallocated buffer in case we only log a few events
}

func (tr *trace) reset() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:741
	_go_fuzz_dep_.CoverTab[45621]++

										tr.Family = ""
										tr.Title = ""
										tr.Start = time.Time{}

										tr.mu.Lock()
										tr.Elapsed = 0
										tr.traceID = 0
										tr.spanID = 0
										tr.IsError = false
										tr.maxEvents = 0
										tr.events = nil
										tr.recycler = nil
										tr.mu.Unlock()

										tr.refs = 0
										tr.disc = 0
										tr.finishStack = nil
										for i := range tr.eventsBuf {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:760
		_go_fuzz_dep_.CoverTab[45622]++
											tr.eventsBuf[i] = event{}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:761
		// _ = "end of CoverTab[45622]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:762
	// _ = "end of CoverTab[45621]"
}

// delta returns the elapsed time since the last event or the trace start,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:765
// and whether it spans midnight.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:765
// L >= tr.mu
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:768
func (tr *trace) delta(t time.Time) (time.Duration, bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:768
	_go_fuzz_dep_.CoverTab[45623]++
										if len(tr.events) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:769
		_go_fuzz_dep_.CoverTab[45625]++
											return t.Sub(tr.Start), false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:770
		// _ = "end of CoverTab[45625]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:771
		_go_fuzz_dep_.CoverTab[45626]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:771
		// _ = "end of CoverTab[45626]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:771
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:771
	// _ = "end of CoverTab[45623]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:771
	_go_fuzz_dep_.CoverTab[45624]++
										prev := tr.events[len(tr.events)-1].When
										return t.Sub(prev), prev.Day() != t.Day()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:773
	// _ = "end of CoverTab[45624]"
}

func (tr *trace) addEvent(x interface{}, recyclable, sensitive bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:776
	_go_fuzz_dep_.CoverTab[45627]++
										if DebugUseAfterFinish && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:777
		_go_fuzz_dep_.CoverTab[45630]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:777
		return tr.finishStack != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:777
		// _ = "end of CoverTab[45630]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:777
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:777
		_go_fuzz_dep_.CoverTab[45631]++
											buf := make([]byte, 4<<10)
											n := runtime.Stack(buf, false)
											log.Printf("net/trace: trace used after finish:\nFinished at:\n%s\nUsed at:\n%s", tr.finishStack, buf[:n])
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:780
		// _ = "end of CoverTab[45631]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:781
		_go_fuzz_dep_.CoverTab[45632]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:781
		// _ = "end of CoverTab[45632]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:781
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:781
	// _ = "end of CoverTab[45627]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:781
	_go_fuzz_dep_.CoverTab[45628]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:797
	e := event{When: time.Now(), What: x, Recyclable: recyclable, Sensitive: sensitive}
	tr.mu.Lock()
	e.Elapsed, e.NewDay = tr.delta(e.When)
	if len(tr.events) < tr.maxEvents {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:800
		_go_fuzz_dep_.CoverTab[45633]++
											tr.events = append(tr.events, e)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:801
		// _ = "end of CoverTab[45633]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:802
		_go_fuzz_dep_.CoverTab[45634]++

											di := int((tr.maxEvents - 1) / 2)
											if d, ok := tr.events[di].What.(*discarded); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:805
			_go_fuzz_dep_.CoverTab[45637]++
												(*d)++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:806
			// _ = "end of CoverTab[45637]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:807
			_go_fuzz_dep_.CoverTab[45638]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:810
			tr.disc = 2
			if tr.recycler != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:811
				_go_fuzz_dep_.CoverTab[45640]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:811
				return tr.events[di].Recyclable
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:811
				// _ = "end of CoverTab[45640]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:811
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:811
				_go_fuzz_dep_.CoverTab[45641]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:811
				_curRoutineNum46_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:811
				_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum46_)
													go tr.recycler(tr.events[di].What)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:812
				// _ = "end of CoverTab[45641]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:813
				_go_fuzz_dep_.CoverTab[45642]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:813
				// _ = "end of CoverTab[45642]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:813
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:813
			// _ = "end of CoverTab[45638]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:813
			_go_fuzz_dep_.CoverTab[45639]++
												tr.events[di].What = &tr.disc
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:814
			// _ = "end of CoverTab[45639]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:815
		// _ = "end of CoverTab[45634]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:815
		_go_fuzz_dep_.CoverTab[45635]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:818
		tr.events[di].When = tr.events[di+1].When

		if tr.recycler != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:820
			_go_fuzz_dep_.CoverTab[45643]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:820
			return tr.events[di+1].Recyclable
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:820
			// _ = "end of CoverTab[45643]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:820
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:820
			_go_fuzz_dep_.CoverTab[45644]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:820
			_curRoutineNum47_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:820
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum47_)
												go tr.recycler(tr.events[di+1].What)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:821
			// _ = "end of CoverTab[45644]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:822
			_go_fuzz_dep_.CoverTab[45645]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:822
			// _ = "end of CoverTab[45645]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:822
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:822
		// _ = "end of CoverTab[45635]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:822
		_go_fuzz_dep_.CoverTab[45636]++
											copy(tr.events[di+1:], tr.events[di+2:])
											tr.events[tr.maxEvents-1] = e
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:824
		// _ = "end of CoverTab[45636]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:825
	// _ = "end of CoverTab[45628]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:825
	_go_fuzz_dep_.CoverTab[45629]++
										tr.mu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:826
	// _ = "end of CoverTab[45629]"
}

func (tr *trace) LazyLog(x fmt.Stringer, sensitive bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:829
	_go_fuzz_dep_.CoverTab[45646]++
										tr.addEvent(x, true, sensitive)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:830
	// _ = "end of CoverTab[45646]"
}

func (tr *trace) LazyPrintf(format string, a ...interface{}) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:833
	_go_fuzz_dep_.CoverTab[45647]++
										tr.addEvent(&lazySprintf{format, a}, false, false)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:834
	// _ = "end of CoverTab[45647]"
}

func (tr *trace) SetError() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:837
	_go_fuzz_dep_.CoverTab[45648]++
										tr.mu.Lock()
										tr.IsError = true
										tr.mu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:840
	// _ = "end of CoverTab[45648]"
}

func (tr *trace) SetRecycler(f func(interface{})) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:843
	_go_fuzz_dep_.CoverTab[45649]++
										tr.mu.Lock()
										tr.recycler = f
										tr.mu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:846
	// _ = "end of CoverTab[45649]"
}

func (tr *trace) SetTraceInfo(traceID, spanID uint64) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:849
	_go_fuzz_dep_.CoverTab[45650]++
										tr.mu.Lock()
										tr.traceID, tr.spanID = traceID, spanID
										tr.mu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:852
	// _ = "end of CoverTab[45650]"
}

func (tr *trace) SetMaxEvents(m int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:855
	_go_fuzz_dep_.CoverTab[45651]++
										tr.mu.Lock()

										if len(tr.events) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:858
		_go_fuzz_dep_.CoverTab[45653]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:858
		return m > 3
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:858
		// _ = "end of CoverTab[45653]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:858
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:858
		_go_fuzz_dep_.CoverTab[45654]++
											tr.maxEvents = m
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:859
		// _ = "end of CoverTab[45654]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:860
		_go_fuzz_dep_.CoverTab[45655]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:860
		// _ = "end of CoverTab[45655]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:860
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:860
	// _ = "end of CoverTab[45651]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:860
	_go_fuzz_dep_.CoverTab[45652]++
										tr.mu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:861
	// _ = "end of CoverTab[45652]"
}

func (tr *trace) ref() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:864
	_go_fuzz_dep_.CoverTab[45656]++
										atomic.AddInt32(&tr.refs, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:865
	// _ = "end of CoverTab[45656]"
}

func (tr *trace) unref() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:868
	_go_fuzz_dep_.CoverTab[45657]++
										if atomic.AddInt32(&tr.refs, -1) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:869
		_go_fuzz_dep_.CoverTab[45658]++
											tr.mu.RLock()
											if tr.recycler != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:871
			_go_fuzz_dep_.CoverTab[45660]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:871
			_curRoutineNum48_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:871
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum48_)

												go func(f func(interface{}), es []event) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:873
				_go_fuzz_dep_.CoverTab[45661]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:873
				defer func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:873
					_go_fuzz_dep_.CoverTab[45662]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:873
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum48_)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:873
					// _ = "end of CoverTab[45662]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:873
				}()
													for _, e := range es {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:874
					_go_fuzz_dep_.CoverTab[45663]++
														if e.Recyclable {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:875
						_go_fuzz_dep_.CoverTab[45664]++
															f(e.What)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:876
						// _ = "end of CoverTab[45664]"
					} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:877
						_go_fuzz_dep_.CoverTab[45665]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:877
						// _ = "end of CoverTab[45665]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:877
					}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:877
					// _ = "end of CoverTab[45663]"
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:878
				// _ = "end of CoverTab[45661]"
			}(tr.recycler, tr.events)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:879
			// _ = "end of CoverTab[45660]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:880
			_go_fuzz_dep_.CoverTab[45666]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:880
			// _ = "end of CoverTab[45666]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:880
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:880
		// _ = "end of CoverTab[45658]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:880
		_go_fuzz_dep_.CoverTab[45659]++
											tr.mu.RUnlock()

											freeTrace(tr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:883
		// _ = "end of CoverTab[45659]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:884
		_go_fuzz_dep_.CoverTab[45667]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:884
		// _ = "end of CoverTab[45667]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:884
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:884
	// _ = "end of CoverTab[45657]"
}

func (tr *trace) When() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:887
	_go_fuzz_dep_.CoverTab[45668]++
										return tr.Start.Format("2006/01/02 15:04:05.000000")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:888
	// _ = "end of CoverTab[45668]"
}

func (tr *trace) ElapsedTime() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:891
	_go_fuzz_dep_.CoverTab[45669]++
										tr.mu.RLock()
										t := tr.Elapsed
										tr.mu.RUnlock()

										if t == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:896
		_go_fuzz_dep_.CoverTab[45671]++

											t = time.Since(tr.Start)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:898
		// _ = "end of CoverTab[45671]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:899
		_go_fuzz_dep_.CoverTab[45672]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:899
		// _ = "end of CoverTab[45672]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:899
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:899
	// _ = "end of CoverTab[45669]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:899
	_go_fuzz_dep_.CoverTab[45670]++
										return fmt.Sprintf("%.6f", t.Seconds())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:900
	// _ = "end of CoverTab[45670]"
}

func (tr *trace) Events() []event {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:903
	_go_fuzz_dep_.CoverTab[45673]++
										tr.mu.RLock()
										defer tr.mu.RUnlock()
										return tr.events
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:906
	// _ = "end of CoverTab[45673]"
}

var traceFreeList = make(chan *trace, 1000)	// TODO(dsymonds): Use sync.Pool?

// newTrace returns a trace ready to use.
func newTrace() *trace {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:912
	_go_fuzz_dep_.CoverTab[45674]++
										select {
	case tr := <-traceFreeList:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:914
		_go_fuzz_dep_.CoverTab[45675]++
											return tr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:915
		// _ = "end of CoverTab[45675]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:916
		_go_fuzz_dep_.CoverTab[45676]++
											return new(trace)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:917
		// _ = "end of CoverTab[45676]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:918
	// _ = "end of CoverTab[45674]"
}

// freeTrace adds tr to traceFreeList if there's room.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:921
// This is non-blocking.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:923
func freeTrace(tr *trace) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:923
	_go_fuzz_dep_.CoverTab[45677]++
										if DebugUseAfterFinish {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:924
		_go_fuzz_dep_.CoverTab[45679]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:925
		// _ = "end of CoverTab[45679]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:926
		_go_fuzz_dep_.CoverTab[45680]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:926
		// _ = "end of CoverTab[45680]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:926
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:926
	// _ = "end of CoverTab[45677]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:926
	_go_fuzz_dep_.CoverTab[45678]++
										tr.reset()
										select {
	case traceFreeList <- tr:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:929
		_go_fuzz_dep_.CoverTab[45681]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:929
		// _ = "end of CoverTab[45681]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:930
		_go_fuzz_dep_.CoverTab[45682]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:930
		// _ = "end of CoverTab[45682]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:931
	// _ = "end of CoverTab[45678]"
}

func elapsed(d time.Duration) string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:934
	_go_fuzz_dep_.CoverTab[45683]++
										b := []byte(fmt.Sprintf("%.6f", d.Seconds()))

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:939
	if d < time.Second {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:939
		_go_fuzz_dep_.CoverTab[45685]++
											dot := bytes.IndexByte(b, '.')
											for i := 0; i < dot; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:941
			_go_fuzz_dep_.CoverTab[45687]++
												b[i] = ' '
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:942
			// _ = "end of CoverTab[45687]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:943
		// _ = "end of CoverTab[45685]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:943
		_go_fuzz_dep_.CoverTab[45686]++
											for i := dot + 1; i < len(b); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:944
			_go_fuzz_dep_.CoverTab[45688]++
												if b[i] == '0' {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:945
				_go_fuzz_dep_.CoverTab[45689]++
													b[i] = ' '
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:946
				// _ = "end of CoverTab[45689]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:947
				_go_fuzz_dep_.CoverTab[45690]++
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:948
				// _ = "end of CoverTab[45690]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:949
			// _ = "end of CoverTab[45688]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:950
		// _ = "end of CoverTab[45686]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:951
		_go_fuzz_dep_.CoverTab[45691]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:951
		// _ = "end of CoverTab[45691]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:951
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:951
	// _ = "end of CoverTab[45683]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:951
	_go_fuzz_dep_.CoverTab[45684]++

										return string(b)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:953
	// _ = "end of CoverTab[45684]"
}

var pageTmplCache *template.Template
var pageTmplOnce sync.Once

func pageTmpl() *template.Template {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:959
	_go_fuzz_dep_.CoverTab[45692]++
										pageTmplOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:960
		_go_fuzz_dep_.CoverTab[45694]++
											pageTmplCache = template.Must(template.New("Page").Funcs(template.FuncMap{
			"elapsed":	elapsed,
			"add":		func(a, b int) int { _go_fuzz_dep_.CoverTab[45695]++; return a + b; // _ = "end of CoverTab[45695]" },
		}).Parse(pageHTML))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:964
		// _ = "end of CoverTab[45694]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:965
	// _ = "end of CoverTab[45692]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:965
	_go_fuzz_dep_.CoverTab[45693]++
										return pageTmplCache
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:966
	// _ = "end of CoverTab[45693]"
}

const pageHTML = `
{{template "Prolog" .}}
{{template "StatusTable" .}}
{{template "Epilog" .}}

{{define "Prolog"}}
<html>
	<head>
	<title>/debug/requests</title>
	<style type="text/css">
		body {
			font-family: sans-serif;
		}
		table#tr-status td.family {
			padding-right: 2em;
		}
		table#tr-status td.active {
			padding-right: 1em;
		}
		table#tr-status td.latency-first {
			padding-left: 1em;
		}
		table#tr-status td.empty {
			color: #aaa;
		}
		table#reqs {
			margin-top: 1em;
		}
		table#reqs tr.first {
			{{if $.Expanded}}font-weight: bold;{{end}}
		}
		table#reqs td {
			font-family: monospace;
		}
		table#reqs td.when {
			text-align: right;
			white-space: nowrap;
		}
		table#reqs td.elapsed {
			padding: 0 0.5em;
			text-align: right;
			white-space: pre;
			width: 10em;
		}
		address {
			font-size: smaller;
			margin-top: 5em;
		}
	</style>
	</head>
	<body>

<h1>/debug/requests</h1>
{{end}} {{/* end of Prolog */}}

{{define "StatusTable"}}
<table id="tr-status">
	{{range $fam := .Families}}
	<tr>
		<td class="family">{{$fam}}</td>

		{{$n := index $.ActiveTraceCount $fam}}
		<td class="active {{if not $n}}empty{{end}}">
			{{if $n}}<a href="?fam={{$fam}}&b=-1{{if $.Expanded}}&exp=1{{end}}">{{end}}
			[{{$n}} active]
			{{if $n}}</a>{{end}}
		</td>

		{{$f := index $.CompletedTraces $fam}}
		{{range $i, $b := $f.Buckets}}
		{{$empty := $b.Empty}}
		<td {{if $empty}}class="empty"{{end}}>
		{{if not $empty}}<a href="?fam={{$fam}}&b={{$i}}{{if $.Expanded}}&exp=1{{end}}">{{end}}
		[{{.Cond}}]
		{{if not $empty}}</a>{{end}}
		</td>
		{{end}}

		{{$nb := len $f.Buckets}}
		<td class="latency-first">
		<a href="?fam={{$fam}}&b={{$nb}}">[minute]</a>
		</td>
		<td>
		<a href="?fam={{$fam}}&b={{add $nb 1}}">[hour]</a>
		</td>
		<td>
		<a href="?fam={{$fam}}&b={{add $nb 2}}">[total]</a>
		</td>

	</tr>
	{{end}}
</table>
{{end}} {{/* end of StatusTable */}}

{{define "Epilog"}}
{{if $.Traces}}
<hr />
<h3>Family: {{$.Family}}</h3>

{{if or $.Expanded $.Traced}}
  <a href="?fam={{$.Family}}&b={{$.Bucket}}">[Normal/Summary]</a>
{{else}}
  [Normal/Summary]
{{end}}

{{if or (not $.Expanded) $.Traced}}
  <a href="?fam={{$.Family}}&b={{$.Bucket}}&exp=1">[Normal/Expanded]</a>
{{else}}
  [Normal/Expanded]
{{end}}

{{if not $.Active}}
	{{if or $.Expanded (not $.Traced)}}
	<a href="?fam={{$.Family}}&b={{$.Bucket}}&rtraced=1">[Traced/Summary]</a>
	{{else}}
	[Traced/Summary]
	{{end}}
	{{if or (not $.Expanded) (not $.Traced)}}
	<a href="?fam={{$.Family}}&b={{$.Bucket}}&exp=1&rtraced=1">[Traced/Expanded]</a>
        {{else}}
	[Traced/Expanded]
	{{end}}
{{end}}

{{if $.Total}}
<p><em>Showing <b>{{len $.Traces}}</b> of <b>{{$.Total}}</b> traces.</em></p>
{{end}}

<table id="reqs">
	<caption>
		{{if $.Active}}Active{{else}}Completed{{end}} Requests
	</caption>
	<tr><th>When</th><th>Elapsed&nbsp;(s)</th></tr>
	{{range $tr := $.Traces}}
	<tr class="first">
		<td class="when">{{$tr.When}}</td>
		<td class="elapsed">{{$tr.ElapsedTime}}</td>
		<td>{{$tr.Title}}</td>
		{{/* TODO: include traceID/spanID */}}
	</tr>
	{{if $.Expanded}}
	{{range $tr.Events}}
	<tr>
		<td class="when">{{.WhenString}}</td>
		<td class="elapsed">{{elapsed .Elapsed}}</td>
		<td>{{if or $.ShowSensitive (not .Sensitive)}}... {{.What}}{{else}}<em>[redacted]</em>{{end}}</td>
	</tr>
	{{end}}
	{{end}}
	{{end}}
</table>
{{end}} {{/* if $.Traces */}}

{{if $.Histogram}}
<h4>Latency (&micro;s) of {{$.Family}} over {{$.HistogramWindow}}</h4>
{{$.Histogram}}
{{end}} {{/* if $.Histogram */}}

	</body>
</html>
{{end}} {{/* end of Epilog */}}
`

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:1130
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/trace.go:1130
var _ = _go_fuzz_dep_.CoverTab
