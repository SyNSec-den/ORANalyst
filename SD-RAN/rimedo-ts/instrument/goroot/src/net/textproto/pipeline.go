// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/textproto/pipeline.go:5
package textproto

//line /usr/local/go/src/net/textproto/pipeline.go:5
import (
//line /usr/local/go/src/net/textproto/pipeline.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/textproto/pipeline.go:5
)
//line /usr/local/go/src/net/textproto/pipeline.go:5
import (
//line /usr/local/go/src/net/textproto/pipeline.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/textproto/pipeline.go:5
)

import (
	"sync"
)

// A Pipeline manages a pipelined in-order request/response sequence.
//line /usr/local/go/src/net/textproto/pipeline.go:11
//
//line /usr/local/go/src/net/textproto/pipeline.go:11
// To use a Pipeline p to manage multiple clients on a connection,
//line /usr/local/go/src/net/textproto/pipeline.go:11
// each client should run:
//line /usr/local/go/src/net/textproto/pipeline.go:11
//
//line /usr/local/go/src/net/textproto/pipeline.go:11
//	id := p.Next()	// take a number
//line /usr/local/go/src/net/textproto/pipeline.go:11
//
//line /usr/local/go/src/net/textproto/pipeline.go:11
//	p.StartRequest(id)	// wait for turn to send request
//line /usr/local/go/src/net/textproto/pipeline.go:11
//	«send request»
//line /usr/local/go/src/net/textproto/pipeline.go:11
//	p.EndRequest(id)	// notify Pipeline that request is sent
//line /usr/local/go/src/net/textproto/pipeline.go:11
//
//line /usr/local/go/src/net/textproto/pipeline.go:11
//	p.StartResponse(id)	// wait for turn to read response
//line /usr/local/go/src/net/textproto/pipeline.go:11
//	«read response»
//line /usr/local/go/src/net/textproto/pipeline.go:11
//	p.EndResponse(id)	// notify Pipeline that response is read
//line /usr/local/go/src/net/textproto/pipeline.go:11
//
//line /usr/local/go/src/net/textproto/pipeline.go:11
// A pipelined server can use the same calls to ensure that
//line /usr/local/go/src/net/textproto/pipeline.go:11
// responses computed in parallel are written in the correct order.
//line /usr/local/go/src/net/textproto/pipeline.go:28
type Pipeline struct {
	mu		sync.Mutex
	id		uint
	request		sequencer
	response	sequencer
}

// Next returns the next id for a request/response pair.
func (p *Pipeline) Next() uint {
//line /usr/local/go/src/net/textproto/pipeline.go:36
	_go_fuzz_dep_.CoverTab[34517]++
							p.mu.Lock()
							id := p.id
							p.id++
							p.mu.Unlock()
							return id
//line /usr/local/go/src/net/textproto/pipeline.go:41
	// _ = "end of CoverTab[34517]"
}

// StartRequest blocks until it is time to send (or, if this is a server, receive)
//line /usr/local/go/src/net/textproto/pipeline.go:44
// the request with the given id.
//line /usr/local/go/src/net/textproto/pipeline.go:46
func (p *Pipeline) StartRequest(id uint) {
//line /usr/local/go/src/net/textproto/pipeline.go:46
	_go_fuzz_dep_.CoverTab[34518]++
							p.request.Start(id)
//line /usr/local/go/src/net/textproto/pipeline.go:47
	// _ = "end of CoverTab[34518]"
}

// EndRequest notifies p that the request with the given id has been sent
//line /usr/local/go/src/net/textproto/pipeline.go:50
// (or, if this is a server, received).
//line /usr/local/go/src/net/textproto/pipeline.go:52
func (p *Pipeline) EndRequest(id uint) {
//line /usr/local/go/src/net/textproto/pipeline.go:52
	_go_fuzz_dep_.CoverTab[34519]++
							p.request.End(id)
//line /usr/local/go/src/net/textproto/pipeline.go:53
	// _ = "end of CoverTab[34519]"
}

// StartResponse blocks until it is time to receive (or, if this is a server, send)
//line /usr/local/go/src/net/textproto/pipeline.go:56
// the request with the given id.
//line /usr/local/go/src/net/textproto/pipeline.go:58
func (p *Pipeline) StartResponse(id uint) {
//line /usr/local/go/src/net/textproto/pipeline.go:58
	_go_fuzz_dep_.CoverTab[34520]++
							p.response.Start(id)
//line /usr/local/go/src/net/textproto/pipeline.go:59
	// _ = "end of CoverTab[34520]"
}

// EndResponse notifies p that the response with the given id has been received
//line /usr/local/go/src/net/textproto/pipeline.go:62
// (or, if this is a server, sent).
//line /usr/local/go/src/net/textproto/pipeline.go:64
func (p *Pipeline) EndResponse(id uint) {
//line /usr/local/go/src/net/textproto/pipeline.go:64
	_go_fuzz_dep_.CoverTab[34521]++
							p.response.End(id)
//line /usr/local/go/src/net/textproto/pipeline.go:65
	// _ = "end of CoverTab[34521]"
}

// A sequencer schedules a sequence of numbered events that must
//line /usr/local/go/src/net/textproto/pipeline.go:68
// happen in order, one after the other. The event numbering must start
//line /usr/local/go/src/net/textproto/pipeline.go:68
// at 0 and increment without skipping. The event number wraps around
//line /usr/local/go/src/net/textproto/pipeline.go:68
// safely as long as there are not 2^32 simultaneous events pending.
//line /usr/local/go/src/net/textproto/pipeline.go:72
type sequencer struct {
	mu	sync.Mutex
	id	uint
	wait	map[uint]chan struct{}
}

// Start waits until it is time for the event numbered id to begin.
//line /usr/local/go/src/net/textproto/pipeline.go:78
// That is, except for the first event, it waits until End(id-1) has
//line /usr/local/go/src/net/textproto/pipeline.go:78
// been called.
//line /usr/local/go/src/net/textproto/pipeline.go:81
func (s *sequencer) Start(id uint) {
//line /usr/local/go/src/net/textproto/pipeline.go:81
	_go_fuzz_dep_.CoverTab[34522]++
							s.mu.Lock()
							if s.id == id {
//line /usr/local/go/src/net/textproto/pipeline.go:83
		_go_fuzz_dep_.CoverTab[34525]++
								s.mu.Unlock()
								return
//line /usr/local/go/src/net/textproto/pipeline.go:85
		// _ = "end of CoverTab[34525]"
	} else {
//line /usr/local/go/src/net/textproto/pipeline.go:86
		_go_fuzz_dep_.CoverTab[34526]++
//line /usr/local/go/src/net/textproto/pipeline.go:86
		// _ = "end of CoverTab[34526]"
//line /usr/local/go/src/net/textproto/pipeline.go:86
	}
//line /usr/local/go/src/net/textproto/pipeline.go:86
	// _ = "end of CoverTab[34522]"
//line /usr/local/go/src/net/textproto/pipeline.go:86
	_go_fuzz_dep_.CoverTab[34523]++
							c := make(chan struct{})
							if s.wait == nil {
//line /usr/local/go/src/net/textproto/pipeline.go:88
		_go_fuzz_dep_.CoverTab[34527]++
								s.wait = make(map[uint]chan struct{})
//line /usr/local/go/src/net/textproto/pipeline.go:89
		// _ = "end of CoverTab[34527]"
	} else {
//line /usr/local/go/src/net/textproto/pipeline.go:90
		_go_fuzz_dep_.CoverTab[34528]++
//line /usr/local/go/src/net/textproto/pipeline.go:90
		// _ = "end of CoverTab[34528]"
//line /usr/local/go/src/net/textproto/pipeline.go:90
	}
//line /usr/local/go/src/net/textproto/pipeline.go:90
	// _ = "end of CoverTab[34523]"
//line /usr/local/go/src/net/textproto/pipeline.go:90
	_go_fuzz_dep_.CoverTab[34524]++
							s.wait[id] = c
							s.mu.Unlock()
							<-c
//line /usr/local/go/src/net/textproto/pipeline.go:93
	// _ = "end of CoverTab[34524]"
}

// End notifies the sequencer that the event numbered id has completed,
//line /usr/local/go/src/net/textproto/pipeline.go:96
// allowing it to schedule the event numbered id+1.  It is a run-time error
//line /usr/local/go/src/net/textproto/pipeline.go:96
// to call End with an id that is not the number of the active event.
//line /usr/local/go/src/net/textproto/pipeline.go:99
func (s *sequencer) End(id uint) {
//line /usr/local/go/src/net/textproto/pipeline.go:99
	_go_fuzz_dep_.CoverTab[34529]++
							s.mu.Lock()
							if s.id != id {
//line /usr/local/go/src/net/textproto/pipeline.go:101
		_go_fuzz_dep_.CoverTab[34533]++
								s.mu.Unlock()
								panic("out of sync")
//line /usr/local/go/src/net/textproto/pipeline.go:103
		// _ = "end of CoverTab[34533]"
	} else {
//line /usr/local/go/src/net/textproto/pipeline.go:104
		_go_fuzz_dep_.CoverTab[34534]++
//line /usr/local/go/src/net/textproto/pipeline.go:104
		// _ = "end of CoverTab[34534]"
//line /usr/local/go/src/net/textproto/pipeline.go:104
	}
//line /usr/local/go/src/net/textproto/pipeline.go:104
	// _ = "end of CoverTab[34529]"
//line /usr/local/go/src/net/textproto/pipeline.go:104
	_go_fuzz_dep_.CoverTab[34530]++
							id++
							s.id = id
							if s.wait == nil {
//line /usr/local/go/src/net/textproto/pipeline.go:107
		_go_fuzz_dep_.CoverTab[34535]++
								s.wait = make(map[uint]chan struct{})
//line /usr/local/go/src/net/textproto/pipeline.go:108
		// _ = "end of CoverTab[34535]"
	} else {
//line /usr/local/go/src/net/textproto/pipeline.go:109
		_go_fuzz_dep_.CoverTab[34536]++
//line /usr/local/go/src/net/textproto/pipeline.go:109
		// _ = "end of CoverTab[34536]"
//line /usr/local/go/src/net/textproto/pipeline.go:109
	}
//line /usr/local/go/src/net/textproto/pipeline.go:109
	// _ = "end of CoverTab[34530]"
//line /usr/local/go/src/net/textproto/pipeline.go:109
	_go_fuzz_dep_.CoverTab[34531]++
							c, ok := s.wait[id]
							if ok {
//line /usr/local/go/src/net/textproto/pipeline.go:111
		_go_fuzz_dep_.CoverTab[34537]++
								delete(s.wait, id)
//line /usr/local/go/src/net/textproto/pipeline.go:112
		// _ = "end of CoverTab[34537]"
	} else {
//line /usr/local/go/src/net/textproto/pipeline.go:113
		_go_fuzz_dep_.CoverTab[34538]++
//line /usr/local/go/src/net/textproto/pipeline.go:113
		// _ = "end of CoverTab[34538]"
//line /usr/local/go/src/net/textproto/pipeline.go:113
	}
//line /usr/local/go/src/net/textproto/pipeline.go:113
	// _ = "end of CoverTab[34531]"
//line /usr/local/go/src/net/textproto/pipeline.go:113
	_go_fuzz_dep_.CoverTab[34532]++
							s.mu.Unlock()
							if ok {
//line /usr/local/go/src/net/textproto/pipeline.go:115
		_go_fuzz_dep_.CoverTab[34539]++
								close(c)
//line /usr/local/go/src/net/textproto/pipeline.go:116
		// _ = "end of CoverTab[34539]"
	} else {
//line /usr/local/go/src/net/textproto/pipeline.go:117
		_go_fuzz_dep_.CoverTab[34540]++
//line /usr/local/go/src/net/textproto/pipeline.go:117
		// _ = "end of CoverTab[34540]"
//line /usr/local/go/src/net/textproto/pipeline.go:117
	}
//line /usr/local/go/src/net/textproto/pipeline.go:117
	// _ = "end of CoverTab[34532]"
}

//line /usr/local/go/src/net/textproto/pipeline.go:118
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/textproto/pipeline.go:118
var _ = _go_fuzz_dep_.CoverTab
