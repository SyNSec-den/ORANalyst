// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Flow control

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:7
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:7
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:7
)

// inflowMinRefresh is the minimum number of bytes we'll send for a
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:9
// flow control window update.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:11
const inflowMinRefresh = 4 << 10

// inflow accounts for an inbound flow control window.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:13
// It tracks both the latest window sent to the peer (used for enforcement)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:13
// and the accumulated unsent window.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:16
type inflow struct {
	avail	int32
	unsent	int32
}

// init sets the initial window.
func (f *inflow) init(n int32) {
	f.avail = n
}

// add adds n bytes to the window, with a maximum window size of max,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:26
// indicating that the peer can now send us more data.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:26
// For example, the user read from a {Request,Response} body and consumed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:26
// some of the buffered data, so the peer can now send more.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:26
// It returns the number of bytes to send in a WINDOW_UPDATE frame to the peer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:26
// Window updates are accumulated and sent when the unsent capacity
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:26
// is at least inflowMinRefresh or will at least double the peer's available window.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:33
func (f *inflow) add(n int) (connAdd int32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:33
	_go_fuzz_dep_.CoverTab[72401]++
										if n < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:34
		_go_fuzz_dep_.CoverTab[72405]++
											panic("negative update")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:35
		// _ = "end of CoverTab[72405]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:36
		_go_fuzz_dep_.CoverTab[72406]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:36
		// _ = "end of CoverTab[72406]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:36
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:36
	// _ = "end of CoverTab[72401]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:36
	_go_fuzz_dep_.CoverTab[72402]++
										unsent := int64(f.unsent) + int64(n)
	// "A sender MUST NOT allow a flow-control window to exceed 2^31-1 octets."
	// RFC 7540 Section 6.9.1.
	const maxWindow = 1<<31 - 1
	if unsent+int64(f.avail) > maxWindow {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:41
		_go_fuzz_dep_.CoverTab[72407]++
											panic("flow control update exceeds maximum window size")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:42
		// _ = "end of CoverTab[72407]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:43
		_go_fuzz_dep_.CoverTab[72408]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:43
		// _ = "end of CoverTab[72408]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:43
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:43
	// _ = "end of CoverTab[72402]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:43
	_go_fuzz_dep_.CoverTab[72403]++
										f.unsent = int32(unsent)
										if f.unsent < inflowMinRefresh && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:45
		_go_fuzz_dep_.CoverTab[72409]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:45
		return f.unsent < f.avail
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:45
		// _ = "end of CoverTab[72409]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:45
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:45
		_go_fuzz_dep_.CoverTab[72410]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:48
		return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:48
		// _ = "end of CoverTab[72410]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:49
		_go_fuzz_dep_.CoverTab[72411]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:49
		// _ = "end of CoverTab[72411]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:49
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:49
	// _ = "end of CoverTab[72403]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:49
	_go_fuzz_dep_.CoverTab[72404]++
										f.avail += f.unsent
										f.unsent = 0
										return int32(unsent)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:52
	// _ = "end of CoverTab[72404]"
}

// take attempts to take n bytes from the peer's flow control window.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:55
// It reports whether the window has available capacity.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:57
func (f *inflow) take(n uint32) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:57
	_go_fuzz_dep_.CoverTab[72412]++
										if n > uint32(f.avail) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:58
		_go_fuzz_dep_.CoverTab[72414]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:59
		// _ = "end of CoverTab[72414]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:60
		_go_fuzz_dep_.CoverTab[72415]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:60
		// _ = "end of CoverTab[72415]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:60
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:60
	// _ = "end of CoverTab[72412]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:60
	_go_fuzz_dep_.CoverTab[72413]++
										f.avail -= int32(n)
										return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:62
	// _ = "end of CoverTab[72413]"
}

// takeInflows attempts to take n bytes from two inflows,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:65
// typically connection-level and stream-level flows.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:65
// It reports whether both windows have available capacity.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:68
func takeInflows(f1, f2 *inflow, n uint32) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:68
	_go_fuzz_dep_.CoverTab[72416]++
										if n > uint32(f1.avail) || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:69
		_go_fuzz_dep_.CoverTab[72418]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:69
		return n > uint32(f2.avail)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:69
		// _ = "end of CoverTab[72418]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:69
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:69
		_go_fuzz_dep_.CoverTab[72419]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:70
		// _ = "end of CoverTab[72419]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:71
		_go_fuzz_dep_.CoverTab[72420]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:71
		// _ = "end of CoverTab[72420]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:71
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:71
	// _ = "end of CoverTab[72416]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:71
	_go_fuzz_dep_.CoverTab[72417]++
										f1.avail -= int32(n)
										f2.avail -= int32(n)
										return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:74
	// _ = "end of CoverTab[72417]"
}

// outflow is the outbound flow control window's size.
type outflow struct {
	_	incomparable

	// n is the number of DATA bytes we're allowed to send.
	// An outflow is kept both on a conn and a per-stream.
	n	int32

	// conn points to the shared connection-level outflow that is
	// shared by all streams on that conn. It is nil for the outflow
	// that's on the conn directly.
	conn	*outflow
}

func (f *outflow) setConnFlow(cf *outflow) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:91
	_go_fuzz_dep_.CoverTab[72421]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:91
	f.conn = cf
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:91
	// _ = "end of CoverTab[72421]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:91
}

func (f *outflow) available() int32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:93
	_go_fuzz_dep_.CoverTab[72422]++
										n := f.n
										if f.conn != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:95
		_go_fuzz_dep_.CoverTab[72424]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:95
		return f.conn.n < n
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:95
		// _ = "end of CoverTab[72424]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:95
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:95
		_go_fuzz_dep_.CoverTab[72425]++
											n = f.conn.n
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:96
		// _ = "end of CoverTab[72425]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:97
		_go_fuzz_dep_.CoverTab[72426]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:97
		// _ = "end of CoverTab[72426]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:97
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:97
	// _ = "end of CoverTab[72422]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:97
	_go_fuzz_dep_.CoverTab[72423]++
										return n
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:98
	// _ = "end of CoverTab[72423]"
}

func (f *outflow) take(n int32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:101
	_go_fuzz_dep_.CoverTab[72427]++
										if n > f.available() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:102
		_go_fuzz_dep_.CoverTab[72429]++
											panic("internal error: took too much")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:103
		// _ = "end of CoverTab[72429]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:104
		_go_fuzz_dep_.CoverTab[72430]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:104
		// _ = "end of CoverTab[72430]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:104
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:104
	// _ = "end of CoverTab[72427]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:104
	_go_fuzz_dep_.CoverTab[72428]++
										f.n -= n
										if f.conn != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:106
		_go_fuzz_dep_.CoverTab[72431]++
											f.conn.n -= n
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:107
		// _ = "end of CoverTab[72431]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:108
		_go_fuzz_dep_.CoverTab[72432]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:108
		// _ = "end of CoverTab[72432]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:108
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:108
	// _ = "end of CoverTab[72428]"
}

// add adds n bytes (positive or negative) to the flow control window.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:111
// It returns false if the sum would exceed 2^31-1.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:113
func (f *outflow) add(n int32) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:113
	_go_fuzz_dep_.CoverTab[72433]++
										sum := f.n + n
										if (sum > n) == (f.n > 0) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:115
		_go_fuzz_dep_.CoverTab[72435]++
											f.n = sum
											return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:117
		// _ = "end of CoverTab[72435]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:118
		_go_fuzz_dep_.CoverTab[72436]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:118
		// _ = "end of CoverTab[72436]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:118
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:118
	// _ = "end of CoverTab[72433]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:118
	_go_fuzz_dep_.CoverTab[72434]++
										return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:119
	// _ = "end of CoverTab[72434]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:120
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/flow.go:120
var _ = _go_fuzz_dep_.CoverTab
