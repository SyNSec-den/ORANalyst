// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:5
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:5
)

import "math"

// NewRandomWriteScheduler constructs a WriteScheduler that ignores HTTP/2
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:9
// priorities. Control frames like SETTINGS and PING are written before DATA
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:9
// frames, but if no control frames are queued and multiple streams have queued
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:9
// HEADERS or DATA frames, Pop selects a ready stream arbitrarily.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:13
func NewRandomWriteScheduler() WriteScheduler {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:13
	_go_fuzz_dep_.CoverTab[75963]++
												return &randomWriteScheduler{sq: make(map[uint32]*writeQueue)}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:14
	// _ = "end of CoverTab[75963]"
}

type randomWriteScheduler struct {
	// zero are frames not associated with a specific stream.
	zero	writeQueue

	// sq contains the stream-specific queues, keyed by stream ID.
	// When a stream is idle, closed, or emptied, it's deleted
	// from the map.
	sq	map[uint32]*writeQueue

	// pool of empty queues for reuse.
	queuePool	writeQueuePool
}

func (ws *randomWriteScheduler) OpenStream(streamID uint32, options OpenStreamOptions) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:30
	_go_fuzz_dep_.CoverTab[75964]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:30
	// _ = "end of CoverTab[75964]"

}

func (ws *randomWriteScheduler) CloseStream(streamID uint32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:34
	_go_fuzz_dep_.CoverTab[75965]++
												q, ok := ws.sq[streamID]
												if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:36
		_go_fuzz_dep_.CoverTab[75967]++
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:37
		// _ = "end of CoverTab[75967]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:38
		_go_fuzz_dep_.CoverTab[75968]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:38
		// _ = "end of CoverTab[75968]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:38
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:38
	// _ = "end of CoverTab[75965]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:38
	_go_fuzz_dep_.CoverTab[75966]++
												delete(ws.sq, streamID)
												ws.queuePool.put(q)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:40
	// _ = "end of CoverTab[75966]"
}

func (ws *randomWriteScheduler) AdjustStream(streamID uint32, priority PriorityParam) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:43
	_go_fuzz_dep_.CoverTab[75969]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:43
	// _ = "end of CoverTab[75969]"

}

func (ws *randomWriteScheduler) Push(wr FrameWriteRequest) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:47
	_go_fuzz_dep_.CoverTab[75970]++
												if wr.isControl() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:48
		_go_fuzz_dep_.CoverTab[75973]++
													ws.zero.push(wr)
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:50
		// _ = "end of CoverTab[75973]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:51
		_go_fuzz_dep_.CoverTab[75974]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:51
		// _ = "end of CoverTab[75974]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:51
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:51
	// _ = "end of CoverTab[75970]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:51
	_go_fuzz_dep_.CoverTab[75971]++
												id := wr.StreamID()
												q, ok := ws.sq[id]
												if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:54
		_go_fuzz_dep_.CoverTab[75975]++
													q = ws.queuePool.get()
													ws.sq[id] = q
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:56
		// _ = "end of CoverTab[75975]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:57
		_go_fuzz_dep_.CoverTab[75976]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:57
		// _ = "end of CoverTab[75976]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:57
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:57
	// _ = "end of CoverTab[75971]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:57
	_go_fuzz_dep_.CoverTab[75972]++
												q.push(wr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:58
	// _ = "end of CoverTab[75972]"
}

func (ws *randomWriteScheduler) Pop() (FrameWriteRequest, bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:61
	_go_fuzz_dep_.CoverTab[75977]++

												if !ws.zero.empty() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:63
		_go_fuzz_dep_.CoverTab[75980]++
													return ws.zero.shift(), true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:64
		// _ = "end of CoverTab[75980]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:65
		_go_fuzz_dep_.CoverTab[75981]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:65
		// _ = "end of CoverTab[75981]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:65
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:65
	// _ = "end of CoverTab[75977]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:65
	_go_fuzz_dep_.CoverTab[75978]++

												for streamID, q := range ws.sq {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:67
		_go_fuzz_dep_.CoverTab[75982]++
													if wr, ok := q.consume(math.MaxInt32); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:68
			_go_fuzz_dep_.CoverTab[75983]++
														if q.empty() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:69
				_go_fuzz_dep_.CoverTab[75985]++
															delete(ws.sq, streamID)
															ws.queuePool.put(q)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:71
				// _ = "end of CoverTab[75985]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:72
				_go_fuzz_dep_.CoverTab[75986]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:72
				// _ = "end of CoverTab[75986]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:72
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:72
			// _ = "end of CoverTab[75983]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:72
			_go_fuzz_dep_.CoverTab[75984]++
														return wr, true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:73
			// _ = "end of CoverTab[75984]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:74
			_go_fuzz_dep_.CoverTab[75987]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:74
			// _ = "end of CoverTab[75987]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:74
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:74
		// _ = "end of CoverTab[75982]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:75
	// _ = "end of CoverTab[75978]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:75
	_go_fuzz_dep_.CoverTab[75979]++
												return FrameWriteRequest{}, false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:76
	// _ = "end of CoverTab[75979]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:77
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched_random.go:77
var _ = _go_fuzz_dep_.CoverTab
