// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:5
package trace

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:5
)

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"text/tabwriter"
	"time"
)

const maxEventsPerLog = 100

type bucket struct {
	MaxErrAge	time.Duration
	String		string
}

var buckets = []bucket{
	{0, "total"},
	{10 * time.Second, "errs<10s"},
	{1 * time.Minute, "errs<1m"},
	{10 * time.Minute, "errs<10m"},
	{1 * time.Hour, "errs<1h"},
	{10 * time.Hour, "errs<10h"},
	{24000 * time.Hour, "errors"},
}

// RenderEvents renders the HTML page typically served at /debug/events.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:41
// It does not do any auth checking. The request may be nil.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:41
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:41
// Most users will use the Events handler.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:45
func RenderEvents(w http.ResponseWriter, req *http.Request, sensitive bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:45
	_go_fuzz_dep_.CoverTab[45240]++
										now := time.Now()
										data := &struct {
		Families	[]string	// family names
		Buckets		[]bucket
		Counts		[][]int	// eventLog count per family/bucket

		// Set when a bucket has been selected.
		Family		string
		Bucket		int
		EventLogs	eventLogs
		Expanded	bool
	}{
		Buckets: buckets,
	}

	data.Families = make([]string, 0, len(families))
	famMu.RLock()
	for name := range families {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:63
		_go_fuzz_dep_.CoverTab[45244]++
											data.Families = append(data.Families, name)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:64
		// _ = "end of CoverTab[45244]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:65
	// _ = "end of CoverTab[45240]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:65
	_go_fuzz_dep_.CoverTab[45241]++
										famMu.RUnlock()
										sort.Strings(data.Families)

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:70
	data.Counts = make([][]int, len(data.Families))
	for i, name := range data.Families {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:71
		_go_fuzz_dep_.CoverTab[45245]++

											f := getEventFamily(name)
											data.Counts[i] = make([]int, len(data.Buckets))
											for j, b := range data.Buckets {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:75
			_go_fuzz_dep_.CoverTab[45246]++
												data.Counts[i][j] = f.Count(now, b.MaxErrAge)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:76
			// _ = "end of CoverTab[45246]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:77
		// _ = "end of CoverTab[45245]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:78
	// _ = "end of CoverTab[45241]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:78
	_go_fuzz_dep_.CoverTab[45242]++

										if req != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:80
		_go_fuzz_dep_.CoverTab[45247]++
											var ok bool
											data.Family, data.Bucket, ok = parseEventsArgs(req)
											if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:83
			_go_fuzz_dep_.CoverTab[45250]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:83
			// _ = "end of CoverTab[45250]"

		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:85
			_go_fuzz_dep_.CoverTab[45251]++
												data.EventLogs = getEventFamily(data.Family).Copy(now, buckets[data.Bucket].MaxErrAge)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:86
			// _ = "end of CoverTab[45251]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:87
		// _ = "end of CoverTab[45247]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:87
		_go_fuzz_dep_.CoverTab[45248]++
											if data.EventLogs != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:88
			_go_fuzz_dep_.CoverTab[45252]++
												defer data.EventLogs.Free()
												sort.Sort(data.EventLogs)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:90
			// _ = "end of CoverTab[45252]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:91
			_go_fuzz_dep_.CoverTab[45253]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:91
			// _ = "end of CoverTab[45253]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:91
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:91
		// _ = "end of CoverTab[45248]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:91
		_go_fuzz_dep_.CoverTab[45249]++
											if exp, err := strconv.ParseBool(req.FormValue("exp")); err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:92
			_go_fuzz_dep_.CoverTab[45254]++
												data.Expanded = exp
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:93
			// _ = "end of CoverTab[45254]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:94
			_go_fuzz_dep_.CoverTab[45255]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:94
			// _ = "end of CoverTab[45255]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:94
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:94
		// _ = "end of CoverTab[45249]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:95
		_go_fuzz_dep_.CoverTab[45256]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:95
		// _ = "end of CoverTab[45256]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:95
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:95
	// _ = "end of CoverTab[45242]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:95
	_go_fuzz_dep_.CoverTab[45243]++

										famMu.RLock()
										defer famMu.RUnlock()
										if err := eventsTmpl().Execute(w, data); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:99
		_go_fuzz_dep_.CoverTab[45257]++
											log.Printf("net/trace: Failed executing template: %v", err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:100
		// _ = "end of CoverTab[45257]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:101
		_go_fuzz_dep_.CoverTab[45258]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:101
		// _ = "end of CoverTab[45258]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:101
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:101
	// _ = "end of CoverTab[45243]"
}

func parseEventsArgs(req *http.Request) (fam string, b int, ok bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:104
	_go_fuzz_dep_.CoverTab[45259]++
										fam, bStr := req.FormValue("fam"), req.FormValue("b")
										if fam == "" || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:106
		_go_fuzz_dep_.CoverTab[45262]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:106
		return bStr == ""
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:106
		// _ = "end of CoverTab[45262]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:106
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:106
		_go_fuzz_dep_.CoverTab[45263]++
											return "", 0, false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:107
		// _ = "end of CoverTab[45263]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:108
		_go_fuzz_dep_.CoverTab[45264]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:108
		// _ = "end of CoverTab[45264]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:108
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:108
	// _ = "end of CoverTab[45259]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:108
	_go_fuzz_dep_.CoverTab[45260]++
										b, err := strconv.Atoi(bStr)
										if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:110
		_go_fuzz_dep_.CoverTab[45265]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:110
		return b < 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:110
		// _ = "end of CoverTab[45265]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:110
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:110
		_go_fuzz_dep_.CoverTab[45266]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:110
		return b >= len(buckets)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:110
		// _ = "end of CoverTab[45266]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:110
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:110
		_go_fuzz_dep_.CoverTab[45267]++
											return "", 0, false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:111
		// _ = "end of CoverTab[45267]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:112
		_go_fuzz_dep_.CoverTab[45268]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:112
		// _ = "end of CoverTab[45268]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:112
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:112
	// _ = "end of CoverTab[45260]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:112
	_go_fuzz_dep_.CoverTab[45261]++
										return fam, b, true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:113
	// _ = "end of CoverTab[45261]"
}

// An EventLog provides a log of events associated with a specific object.
type EventLog interface {
	// Printf formats its arguments with fmt.Sprintf and adds the
	// result to the event log.
	Printf(format string, a ...interface{})

	// Errorf is like Printf, but it marks this event as an error.
	Errorf(format string, a ...interface{})

	// Finish declares that this event log is complete.
	// The event log should not be used after calling this method.
	Finish()
}

// NewEventLog returns a new EventLog with the specified family name
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:130
// and title.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:132
func NewEventLog(family, title string) EventLog {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:132
	_go_fuzz_dep_.CoverTab[45269]++
										el := newEventLog()
										el.ref()
										el.Family, el.Title = family, title
										el.Start = time.Now()
										el.events = make([]logEntry, 0, maxEventsPerLog)
										el.stack = make([]uintptr, 32)
										n := runtime.Callers(2, el.stack)
										el.stack = el.stack[:n]

										getEventFamily(family).add(el)
										return el
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:143
	// _ = "end of CoverTab[45269]"
}

func (el *eventLog) Finish() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:146
	_go_fuzz_dep_.CoverTab[45270]++
										getEventFamily(el.Family).remove(el)
										el.unref()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:148
	// _ = "end of CoverTab[45270]"
}

var (
	famMu		sync.RWMutex
	families	= make(map[string]*eventFamily)	// family name => family
)

func getEventFamily(fam string) *eventFamily {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:156
	_go_fuzz_dep_.CoverTab[45271]++
										famMu.Lock()
										defer famMu.Unlock()
										f := families[fam]
										if f == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:160
		_go_fuzz_dep_.CoverTab[45273]++
											f = &eventFamily{}
											families[fam] = f
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:162
		// _ = "end of CoverTab[45273]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:163
		_go_fuzz_dep_.CoverTab[45274]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:163
		// _ = "end of CoverTab[45274]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:163
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:163
	// _ = "end of CoverTab[45271]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:163
	_go_fuzz_dep_.CoverTab[45272]++
										return f
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:164
	// _ = "end of CoverTab[45272]"
}

type eventFamily struct {
	mu		sync.RWMutex
	eventLogs	eventLogs
}

func (f *eventFamily) add(el *eventLog) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:172
	_go_fuzz_dep_.CoverTab[45275]++
										f.mu.Lock()
										f.eventLogs = append(f.eventLogs, el)
										f.mu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:175
	// _ = "end of CoverTab[45275]"
}

func (f *eventFamily) remove(el *eventLog) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:178
	_go_fuzz_dep_.CoverTab[45276]++
										f.mu.Lock()
										defer f.mu.Unlock()
										for i, el0 := range f.eventLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:181
		_go_fuzz_dep_.CoverTab[45277]++
											if el == el0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:182
			_go_fuzz_dep_.CoverTab[45278]++
												copy(f.eventLogs[i:], f.eventLogs[i+1:])
												f.eventLogs = f.eventLogs[:len(f.eventLogs)-1]
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:185
			// _ = "end of CoverTab[45278]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:186
			_go_fuzz_dep_.CoverTab[45279]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:186
			// _ = "end of CoverTab[45279]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:186
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:186
		// _ = "end of CoverTab[45277]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:187
	// _ = "end of CoverTab[45276]"
}

func (f *eventFamily) Count(now time.Time, maxErrAge time.Duration) (n int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:190
	_go_fuzz_dep_.CoverTab[45280]++
										f.mu.RLock()
										defer f.mu.RUnlock()
										for _, el := range f.eventLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:193
		_go_fuzz_dep_.CoverTab[45282]++
											if el.hasRecentError(now, maxErrAge) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:194
			_go_fuzz_dep_.CoverTab[45283]++
												n++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:195
			// _ = "end of CoverTab[45283]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:196
			_go_fuzz_dep_.CoverTab[45284]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:196
			// _ = "end of CoverTab[45284]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:196
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:196
		// _ = "end of CoverTab[45282]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:197
	// _ = "end of CoverTab[45280]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:197
	_go_fuzz_dep_.CoverTab[45281]++
										return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:198
	// _ = "end of CoverTab[45281]"
}

func (f *eventFamily) Copy(now time.Time, maxErrAge time.Duration) (els eventLogs) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:201
	_go_fuzz_dep_.CoverTab[45285]++
										f.mu.RLock()
										defer f.mu.RUnlock()
										els = make(eventLogs, 0, len(f.eventLogs))
										for _, el := range f.eventLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:205
		_go_fuzz_dep_.CoverTab[45287]++
											if el.hasRecentError(now, maxErrAge) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:206
			_go_fuzz_dep_.CoverTab[45288]++
												el.ref()
												els = append(els, el)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:208
			// _ = "end of CoverTab[45288]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:209
			_go_fuzz_dep_.CoverTab[45289]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:209
			// _ = "end of CoverTab[45289]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:209
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:209
		// _ = "end of CoverTab[45287]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:210
	// _ = "end of CoverTab[45285]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:210
	_go_fuzz_dep_.CoverTab[45286]++
										return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:211
	// _ = "end of CoverTab[45286]"
}

type eventLogs []*eventLog

// Free calls unref on each element of the list.
func (els eventLogs) Free() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:217
	_go_fuzz_dep_.CoverTab[45290]++
										for _, el := range els {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:218
		_go_fuzz_dep_.CoverTab[45291]++
											el.unref()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:219
		// _ = "end of CoverTab[45291]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:220
	// _ = "end of CoverTab[45290]"
}

// eventLogs may be sorted in reverse chronological order.
func (els eventLogs) Len() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:224
	_go_fuzz_dep_.CoverTab[45292]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:224
	return len(els)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:224
	// _ = "end of CoverTab[45292]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:224
}
func (els eventLogs) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:225
	_go_fuzz_dep_.CoverTab[45293]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:225
	return els[i].Start.After(els[j].Start)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:225
	// _ = "end of CoverTab[45293]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:225
}
func (els eventLogs) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:226
	_go_fuzz_dep_.CoverTab[45294]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:226
	els[i], els[j] = els[j], els[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:226
	// _ = "end of CoverTab[45294]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:226
}

// A logEntry is a timestamped log entry in an event log.
type logEntry struct {
	When	time.Time
	Elapsed	time.Duration	// since previous event in log
	NewDay	bool		// whether this event is on a different day to the previous event
	What	string
	IsErr	bool
}

// WhenString returns a string representation of the elapsed time of the event.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:237
// It will include the date if midnight was crossed.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:239
func (e logEntry) WhenString() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:239
	_go_fuzz_dep_.CoverTab[45295]++
										if e.NewDay {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:240
		_go_fuzz_dep_.CoverTab[45297]++
											return e.When.Format("2006/01/02 15:04:05.000000")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:241
		// _ = "end of CoverTab[45297]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:242
		_go_fuzz_dep_.CoverTab[45298]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:242
		// _ = "end of CoverTab[45298]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:242
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:242
	// _ = "end of CoverTab[45295]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:242
	_go_fuzz_dep_.CoverTab[45296]++
										return e.When.Format("15:04:05.000000")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:243
	// _ = "end of CoverTab[45296]"
}

// An eventLog represents an active event log.
type eventLog struct {
	// Family is the top-level grouping of event logs to which this belongs.
	Family	string

	// Title is the title of this event log.
	Title	string

	// Timing information.
	Start	time.Time

	// Call stack where this event log was created.
	stack	[]uintptr

	// Append-only sequence of events.
	//
	// TODO(sameer): change this to a ring buffer to avoid the array copy
	// when we hit maxEventsPerLog.
	mu		sync.RWMutex
	events		[]logEntry
	LastErrorTime	time.Time
	discarded	int

	refs	int32	// how many buckets this is in
}

func (el *eventLog) reset() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:272
	_go_fuzz_dep_.CoverTab[45299]++

										el.Family = ""
										el.Title = ""
										el.Start = time.Time{}
										el.stack = nil
										el.events = nil
										el.LastErrorTime = time.Time{}
										el.discarded = 0
										el.refs = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:281
	// _ = "end of CoverTab[45299]"
}

func (el *eventLog) hasRecentError(now time.Time, maxErrAge time.Duration) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:284
	_go_fuzz_dep_.CoverTab[45300]++
										if maxErrAge == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:285
		_go_fuzz_dep_.CoverTab[45302]++
											return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:286
		// _ = "end of CoverTab[45302]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:287
		_go_fuzz_dep_.CoverTab[45303]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:287
		// _ = "end of CoverTab[45303]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:287
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:287
	// _ = "end of CoverTab[45300]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:287
	_go_fuzz_dep_.CoverTab[45301]++
										el.mu.RLock()
										defer el.mu.RUnlock()
										return now.Sub(el.LastErrorTime) < maxErrAge
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:290
	// _ = "end of CoverTab[45301]"
}

// delta returns the elapsed time since the last event or the log start,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:293
// and whether it spans midnight.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:293
// L >= el.mu
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:296
func (el *eventLog) delta(t time.Time) (time.Duration, bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:296
	_go_fuzz_dep_.CoverTab[45304]++
										if len(el.events) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:297
		_go_fuzz_dep_.CoverTab[45306]++
											return t.Sub(el.Start), false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:298
		// _ = "end of CoverTab[45306]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:299
		_go_fuzz_dep_.CoverTab[45307]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:299
		// _ = "end of CoverTab[45307]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:299
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:299
	// _ = "end of CoverTab[45304]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:299
	_go_fuzz_dep_.CoverTab[45305]++
										prev := el.events[len(el.events)-1].When
										return t.Sub(prev), prev.Day() != t.Day()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:301
	// _ = "end of CoverTab[45305]"

}

func (el *eventLog) Printf(format string, a ...interface{}) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:305
	_go_fuzz_dep_.CoverTab[45308]++
										el.printf(false, format, a...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:306
	// _ = "end of CoverTab[45308]"
}

func (el *eventLog) Errorf(format string, a ...interface{}) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:309
	_go_fuzz_dep_.CoverTab[45309]++
										el.printf(true, format, a...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:310
	// _ = "end of CoverTab[45309]"
}

func (el *eventLog) printf(isErr bool, format string, a ...interface{}) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:313
	_go_fuzz_dep_.CoverTab[45310]++
										e := logEntry{When: time.Now(), IsErr: isErr, What: fmt.Sprintf(format, a...)}
										el.mu.Lock()
										e.Elapsed, e.NewDay = el.delta(e.When)
										if len(el.events) < maxEventsPerLog {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:317
		_go_fuzz_dep_.CoverTab[45313]++
											el.events = append(el.events, e)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:318
		// _ = "end of CoverTab[45313]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:319
		_go_fuzz_dep_.CoverTab[45314]++

											if el.discarded == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:321
			_go_fuzz_dep_.CoverTab[45316]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:325
			el.discarded = 2
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:325
			// _ = "end of CoverTab[45316]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:326
			_go_fuzz_dep_.CoverTab[45317]++
												el.discarded++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:327
			// _ = "end of CoverTab[45317]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:328
		// _ = "end of CoverTab[45314]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:328
		_go_fuzz_dep_.CoverTab[45315]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:331
		el.events[0].What = fmt.Sprintf("(%d events discarded)", el.discarded)

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:334
		el.events[0].When = el.events[1].When
											copy(el.events[1:], el.events[2:])
											el.events[maxEventsPerLog-1] = e
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:336
		// _ = "end of CoverTab[45315]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:337
	// _ = "end of CoverTab[45310]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:337
	_go_fuzz_dep_.CoverTab[45311]++
										if e.IsErr {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:338
		_go_fuzz_dep_.CoverTab[45318]++
											el.LastErrorTime = e.When
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:339
		// _ = "end of CoverTab[45318]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:340
		_go_fuzz_dep_.CoverTab[45319]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:340
		// _ = "end of CoverTab[45319]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:340
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:340
	// _ = "end of CoverTab[45311]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:340
	_go_fuzz_dep_.CoverTab[45312]++
										el.mu.Unlock()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:341
	// _ = "end of CoverTab[45312]"
}

func (el *eventLog) ref() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:344
	_go_fuzz_dep_.CoverTab[45320]++
										atomic.AddInt32(&el.refs, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:345
	// _ = "end of CoverTab[45320]"
}

func (el *eventLog) unref() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:348
	_go_fuzz_dep_.CoverTab[45321]++
										if atomic.AddInt32(&el.refs, -1) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:349
		_go_fuzz_dep_.CoverTab[45322]++
											freeEventLog(el)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:350
		// _ = "end of CoverTab[45322]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:351
		_go_fuzz_dep_.CoverTab[45323]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:351
		// _ = "end of CoverTab[45323]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:351
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:351
	// _ = "end of CoverTab[45321]"
}

func (el *eventLog) When() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:354
	_go_fuzz_dep_.CoverTab[45324]++
										return el.Start.Format("2006/01/02 15:04:05.000000")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:355
	// _ = "end of CoverTab[45324]"
}

func (el *eventLog) ElapsedTime() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:358
	_go_fuzz_dep_.CoverTab[45325]++
										elapsed := time.Since(el.Start)
										return fmt.Sprintf("%.6f", elapsed.Seconds())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:360
	// _ = "end of CoverTab[45325]"
}

func (el *eventLog) Stack() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:363
	_go_fuzz_dep_.CoverTab[45326]++
										buf := new(bytes.Buffer)
										tw := tabwriter.NewWriter(buf, 1, 8, 1, '\t', 0)
										printStackRecord(tw, el.stack)
										tw.Flush()
										return buf.String()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:368
	// _ = "end of CoverTab[45326]"
}

// printStackRecord prints the function + source line information
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:371
// for a single stack trace.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:371
// Adapted from runtime/pprof/pprof.go.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:374
func printStackRecord(w io.Writer, stk []uintptr) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:374
	_go_fuzz_dep_.CoverTab[45327]++
										for _, pc := range stk {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:375
		_go_fuzz_dep_.CoverTab[45328]++
											f := runtime.FuncForPC(pc)
											if f == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:377
			_go_fuzz_dep_.CoverTab[45331]++
												continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:378
			// _ = "end of CoverTab[45331]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:379
			_go_fuzz_dep_.CoverTab[45332]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:379
			// _ = "end of CoverTab[45332]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:379
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:379
		// _ = "end of CoverTab[45328]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:379
		_go_fuzz_dep_.CoverTab[45329]++
											file, line := f.FileLine(pc)
											name := f.Name()

											if strings.HasPrefix(name, "runtime.") {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:383
			_go_fuzz_dep_.CoverTab[45333]++
												continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:384
			// _ = "end of CoverTab[45333]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:385
			_go_fuzz_dep_.CoverTab[45334]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:385
			// _ = "end of CoverTab[45334]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:385
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:385
		// _ = "end of CoverTab[45329]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:385
		_go_fuzz_dep_.CoverTab[45330]++
											fmt.Fprintf(w, "#   %s\t%s:%d\n", name, file, line)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:386
		// _ = "end of CoverTab[45330]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:387
	// _ = "end of CoverTab[45327]"
}

func (el *eventLog) Events() []logEntry {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:390
	_go_fuzz_dep_.CoverTab[45335]++
										el.mu.RLock()
										defer el.mu.RUnlock()
										return el.events
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:393
	// _ = "end of CoverTab[45335]"
}

// freeEventLogs is a freelist of *eventLog
var freeEventLogs = make(chan *eventLog, 1000)

// newEventLog returns a event log ready to use.
func newEventLog() *eventLog {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:400
	_go_fuzz_dep_.CoverTab[45336]++
										select {
	case el := <-freeEventLogs:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:402
		_go_fuzz_dep_.CoverTab[45337]++
											return el
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:403
		// _ = "end of CoverTab[45337]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:404
		_go_fuzz_dep_.CoverTab[45338]++
											return new(eventLog)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:405
		// _ = "end of CoverTab[45338]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:406
	// _ = "end of CoverTab[45336]"
}

// freeEventLog adds el to freeEventLogs if there's room.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:409
// This is non-blocking.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:411
func freeEventLog(el *eventLog) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:411
	_go_fuzz_dep_.CoverTab[45339]++
										el.reset()
										select {
	case freeEventLogs <- el:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:414
		_go_fuzz_dep_.CoverTab[45340]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:414
		// _ = "end of CoverTab[45340]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:415
		_go_fuzz_dep_.CoverTab[45341]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:415
		// _ = "end of CoverTab[45341]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:416
	// _ = "end of CoverTab[45339]"
}

var eventsTmplCache *template.Template
var eventsTmplOnce sync.Once

func eventsTmpl() *template.Template {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:422
	_go_fuzz_dep_.CoverTab[45342]++
										eventsTmplOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:423
		_go_fuzz_dep_.CoverTab[45344]++
											eventsTmplCache = template.Must(template.New("events").Funcs(template.FuncMap{
			"elapsed":	elapsed,
			"trimSpace":	strings.TrimSpace,
		}).Parse(eventsHTML))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:427
		// _ = "end of CoverTab[45344]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:428
	// _ = "end of CoverTab[45342]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:428
	_go_fuzz_dep_.CoverTab[45343]++
										return eventsTmplCache
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:429
	// _ = "end of CoverTab[45343]"
}

const eventsHTML = `
<html>
	<head>
		<title>events</title>
	</head>
	<style type="text/css">
		body {
			font-family: sans-serif;
		}
		table#req-status td.family {
			padding-right: 2em;
		}
		table#req-status td.active {
			padding-right: 1em;
		}
		table#req-status td.empty {
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
	<body>

<h1>/debug/events</h1>

<table id="req-status">
	{{range $i, $fam := .Families}}
	<tr>
		<td class="family">{{$fam}}</td>

	        {{range $j, $bucket := $.Buckets}}
	        {{$n := index $.Counts $i $j}}
		<td class="{{if not $bucket.MaxErrAge}}active{{end}}{{if not $n}}empty{{end}}">
	                {{if $n}}<a href="?fam={{$fam}}&b={{$j}}{{if $.Expanded}}&exp=1{{end}}">{{end}}
		        [{{$n}} {{$bucket.String}}]
			{{if $n}}</a>{{end}}
		</td>
                {{end}}

	</tr>{{end}}
</table>

{{if $.EventLogs}}
<hr />
<h3>Family: {{$.Family}}</h3>

{{if $.Expanded}}<a href="?fam={{$.Family}}&b={{$.Bucket}}">{{end}}
[Summary]{{if $.Expanded}}</a>{{end}}

{{if not $.Expanded}}<a href="?fam={{$.Family}}&b={{$.Bucket}}&exp=1">{{end}}
[Expanded]{{if not $.Expanded}}</a>{{end}}

<table id="reqs">
	<tr><th>When</th><th>Elapsed</th></tr>
	{{range $el := $.EventLogs}}
	<tr class="first">
		<td class="when">{{$el.When}}</td>
		<td class="elapsed">{{$el.ElapsedTime}}</td>
		<td>{{$el.Title}}
	</tr>
	{{if $.Expanded}}
	<tr>
		<td class="when"></td>
		<td class="elapsed"></td>
		<td><pre>{{$el.Stack|trimSpace}}</pre></td>
	</tr>
	{{range $el.Events}}
	<tr>
		<td class="when">{{.WhenString}}</td>
		<td class="elapsed">{{elapsed .Elapsed}}</td>
		<td>.{{if .IsErr}}E{{else}}.{{end}}. {{.What}}</td>
	</tr>
	{{end}}
	{{end}}
	{{end}}
</table>
{{end}}
	</body>
</html>
`

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:532
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/trace/events.go:532
var _ = _go_fuzz_dep_.CoverTab
