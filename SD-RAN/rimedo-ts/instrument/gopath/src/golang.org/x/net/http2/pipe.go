// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:5
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:5
)

import (
	"errors"
	"io"
	"sync"
)

// pipe is a goroutine-safe io.Reader/io.Writer pair. It's like
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:13
// io.Pipe except there are no PipeReader/PipeWriter halves, and the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:13
// underlying buffer is an interface. (io.Pipe is always unbuffered)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:16
type pipe struct {
	mu		sync.Mutex
	c		sync.Cond	// c.L lazily initialized to &p.mu
	b		pipeBuffer	// nil when done reading
	unread		int		// bytes unread when done
	err		error		// read error once empty. non-nil means closed.
	breakErr	error		// immediate read error (caller doesn't see rest of b)
	donec		chan struct{}	// closed on error
	readFn		func()		// optional code to run in Read before error
}

type pipeBuffer interface {
	Len() int
	io.Writer
	io.Reader
}

// setBuffer initializes the pipe buffer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:33
// It has no effect if the pipe is already closed.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:35
func (p *pipe) setBuffer(b pipeBuffer) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:35
	_go_fuzz_dep_.CoverTab[73139]++
										p.mu.Lock()
										defer p.mu.Unlock()
										if p.err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:38
		_go_fuzz_dep_.CoverTab[73141]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:38
		return p.breakErr != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:38
		// _ = "end of CoverTab[73141]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:38
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:38
		_go_fuzz_dep_.CoverTab[73142]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:39
		// _ = "end of CoverTab[73142]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:40
		_go_fuzz_dep_.CoverTab[73143]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:40
		// _ = "end of CoverTab[73143]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:40
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:40
	// _ = "end of CoverTab[73139]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:40
	_go_fuzz_dep_.CoverTab[73140]++
										p.b = b
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:41
	// _ = "end of CoverTab[73140]"
}

func (p *pipe) Len() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:44
	_go_fuzz_dep_.CoverTab[73144]++
										p.mu.Lock()
										defer p.mu.Unlock()
										if p.b == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:47
		_go_fuzz_dep_.CoverTab[73146]++
											return p.unread
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:48
		// _ = "end of CoverTab[73146]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:49
		_go_fuzz_dep_.CoverTab[73147]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:49
		// _ = "end of CoverTab[73147]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:49
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:49
	// _ = "end of CoverTab[73144]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:49
	_go_fuzz_dep_.CoverTab[73145]++
										return p.b.Len()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:50
	// _ = "end of CoverTab[73145]"
}

// Read waits until data is available and copies bytes
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:53
// from the buffer into p.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:55
func (p *pipe) Read(d []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:55
	_go_fuzz_dep_.CoverTab[73148]++
										p.mu.Lock()
										defer p.mu.Unlock()
										if p.c.L == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:58
		_go_fuzz_dep_.CoverTab[73150]++
											p.c.L = &p.mu
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:59
		// _ = "end of CoverTab[73150]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:60
		_go_fuzz_dep_.CoverTab[73151]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:60
		// _ = "end of CoverTab[73151]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:60
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:60
	// _ = "end of CoverTab[73148]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:60
	_go_fuzz_dep_.CoverTab[73149]++
										for {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:61
		_go_fuzz_dep_.CoverTab[73152]++
											if p.breakErr != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:62
			_go_fuzz_dep_.CoverTab[73156]++
												return 0, p.breakErr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:63
			// _ = "end of CoverTab[73156]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:64
			_go_fuzz_dep_.CoverTab[73157]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:64
			// _ = "end of CoverTab[73157]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:64
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:64
		// _ = "end of CoverTab[73152]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:64
		_go_fuzz_dep_.CoverTab[73153]++
											if p.b != nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:65
			_go_fuzz_dep_.CoverTab[73158]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:65
			return p.b.Len() > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:65
			// _ = "end of CoverTab[73158]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:65
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:65
			_go_fuzz_dep_.CoverTab[73159]++
												return p.b.Read(d)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:66
			// _ = "end of CoverTab[73159]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:67
			_go_fuzz_dep_.CoverTab[73160]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:67
			// _ = "end of CoverTab[73160]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:67
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:67
		// _ = "end of CoverTab[73153]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:67
		_go_fuzz_dep_.CoverTab[73154]++
											if p.err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:68
			_go_fuzz_dep_.CoverTab[73161]++
												if p.readFn != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:69
				_go_fuzz_dep_.CoverTab[73163]++
													p.readFn()
													p.readFn = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:71
				// _ = "end of CoverTab[73163]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:72
				_go_fuzz_dep_.CoverTab[73164]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:72
				// _ = "end of CoverTab[73164]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:72
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:72
			// _ = "end of CoverTab[73161]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:72
			_go_fuzz_dep_.CoverTab[73162]++
												p.b = nil
												return 0, p.err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:74
			// _ = "end of CoverTab[73162]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:75
			_go_fuzz_dep_.CoverTab[73165]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:75
			// _ = "end of CoverTab[73165]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:75
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:75
		// _ = "end of CoverTab[73154]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:75
		_go_fuzz_dep_.CoverTab[73155]++
											p.c.Wait()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:76
		// _ = "end of CoverTab[73155]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:77
	// _ = "end of CoverTab[73149]"
}

var errClosedPipeWrite = errors.New("write on closed buffer")

// Write copies bytes from p into the buffer and wakes a reader.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:82
// It is an error to write more data than the buffer can hold.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:84
func (p *pipe) Write(d []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:84
	_go_fuzz_dep_.CoverTab[73166]++
										p.mu.Lock()
										defer p.mu.Unlock()
										if p.c.L == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:87
		_go_fuzz_dep_.CoverTab[73169]++
											p.c.L = &p.mu
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:88
		// _ = "end of CoverTab[73169]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:89
		_go_fuzz_dep_.CoverTab[73170]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:89
		// _ = "end of CoverTab[73170]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:89
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:89
	// _ = "end of CoverTab[73166]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:89
	_go_fuzz_dep_.CoverTab[73167]++
										defer p.c.Signal()
										if p.err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:91
		_go_fuzz_dep_.CoverTab[73171]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:91
		return p.breakErr != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:91
		// _ = "end of CoverTab[73171]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:91
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:91
		_go_fuzz_dep_.CoverTab[73172]++
											return 0, errClosedPipeWrite
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:92
		// _ = "end of CoverTab[73172]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:93
		_go_fuzz_dep_.CoverTab[73173]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:93
		// _ = "end of CoverTab[73173]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:93
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:93
	// _ = "end of CoverTab[73167]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:93
	_go_fuzz_dep_.CoverTab[73168]++
										return p.b.Write(d)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:94
	// _ = "end of CoverTab[73168]"
}

// CloseWithError causes the next Read (waking up a current blocked
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:97
// Read if needed) to return the provided err after all data has been
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:97
// read.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:97
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:97
// The error must be non-nil.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:102
func (p *pipe) CloseWithError(err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:102
	_go_fuzz_dep_.CoverTab[73174]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:102
	p.closeWithError(&p.err, err, nil)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:102
	// _ = "end of CoverTab[73174]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:102
}

// BreakWithError causes the next Read (waking up a current blocked
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:104
// Read if needed) to return the provided err immediately, without
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:104
// waiting for unread data.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:107
func (p *pipe) BreakWithError(err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:107
	_go_fuzz_dep_.CoverTab[73175]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:107
	p.closeWithError(&p.breakErr, err, nil)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:107
	// _ = "end of CoverTab[73175]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:107
}

// closeWithErrorAndCode is like CloseWithError but also sets some code to run
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:109
// in the caller's goroutine before returning the error.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:111
func (p *pipe) closeWithErrorAndCode(err error, fn func()) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:111
	_go_fuzz_dep_.CoverTab[73176]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:111
	p.closeWithError(&p.err, err, fn)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:111
	// _ = "end of CoverTab[73176]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:111
}

func (p *pipe) closeWithError(dst *error, err error, fn func()) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:113
	_go_fuzz_dep_.CoverTab[73177]++
										if err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:114
		_go_fuzz_dep_.CoverTab[73182]++
											panic("err must be non-nil")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:115
		// _ = "end of CoverTab[73182]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:116
		_go_fuzz_dep_.CoverTab[73183]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:116
		// _ = "end of CoverTab[73183]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:116
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:116
	// _ = "end of CoverTab[73177]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:116
	_go_fuzz_dep_.CoverTab[73178]++
										p.mu.Lock()
										defer p.mu.Unlock()
										if p.c.L == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:119
		_go_fuzz_dep_.CoverTab[73184]++
											p.c.L = &p.mu
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:120
		// _ = "end of CoverTab[73184]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:121
		_go_fuzz_dep_.CoverTab[73185]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:121
		// _ = "end of CoverTab[73185]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:121
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:121
	// _ = "end of CoverTab[73178]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:121
	_go_fuzz_dep_.CoverTab[73179]++
										defer p.c.Signal()
										if *dst != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:123
		_go_fuzz_dep_.CoverTab[73186]++

											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:125
		// _ = "end of CoverTab[73186]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:126
		_go_fuzz_dep_.CoverTab[73187]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:126
		// _ = "end of CoverTab[73187]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:126
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:126
	// _ = "end of CoverTab[73179]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:126
	_go_fuzz_dep_.CoverTab[73180]++
										p.readFn = fn
										if dst == &p.breakErr {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:128
		_go_fuzz_dep_.CoverTab[73188]++
											if p.b != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:129
			_go_fuzz_dep_.CoverTab[73190]++
												p.unread += p.b.Len()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:130
			// _ = "end of CoverTab[73190]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:131
			_go_fuzz_dep_.CoverTab[73191]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:131
			// _ = "end of CoverTab[73191]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:131
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:131
		// _ = "end of CoverTab[73188]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:131
		_go_fuzz_dep_.CoverTab[73189]++
											p.b = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:132
		// _ = "end of CoverTab[73189]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:133
		_go_fuzz_dep_.CoverTab[73192]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:133
		// _ = "end of CoverTab[73192]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:133
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:133
	// _ = "end of CoverTab[73180]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:133
	_go_fuzz_dep_.CoverTab[73181]++
										*dst = err
										p.closeDoneLocked()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:135
	// _ = "end of CoverTab[73181]"
}

// requires p.mu be held.
func (p *pipe) closeDoneLocked() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:139
	_go_fuzz_dep_.CoverTab[73193]++
										if p.donec == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:140
		_go_fuzz_dep_.CoverTab[73195]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:141
		// _ = "end of CoverTab[73195]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:142
		_go_fuzz_dep_.CoverTab[73196]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:142
		// _ = "end of CoverTab[73196]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:142
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:142
	// _ = "end of CoverTab[73193]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:142
	_go_fuzz_dep_.CoverTab[73194]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:145
	select {
	case <-p.donec:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:146
		_go_fuzz_dep_.CoverTab[73197]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:146
		// _ = "end of CoverTab[73197]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:147
		_go_fuzz_dep_.CoverTab[73198]++
											close(p.donec)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:148
		// _ = "end of CoverTab[73198]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:149
	// _ = "end of CoverTab[73194]"
}

// Err returns the error (if any) first set by BreakWithError or CloseWithError.
func (p *pipe) Err() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:153
	_go_fuzz_dep_.CoverTab[73199]++
										p.mu.Lock()
										defer p.mu.Unlock()
										if p.breakErr != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:156
		_go_fuzz_dep_.CoverTab[73201]++
											return p.breakErr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:157
		// _ = "end of CoverTab[73201]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:158
		_go_fuzz_dep_.CoverTab[73202]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:158
		// _ = "end of CoverTab[73202]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:158
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:158
	// _ = "end of CoverTab[73199]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:158
	_go_fuzz_dep_.CoverTab[73200]++
										return p.err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:159
	// _ = "end of CoverTab[73200]"
}

// Done returns a channel which is closed if and when this pipe is closed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:162
// with CloseWithError.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:164
func (p *pipe) Done() <-chan struct{} {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:164
	_go_fuzz_dep_.CoverTab[73203]++
										p.mu.Lock()
										defer p.mu.Unlock()
										if p.donec == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:167
		_go_fuzz_dep_.CoverTab[73205]++
											p.donec = make(chan struct{})
											if p.err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:169
			_go_fuzz_dep_.CoverTab[73206]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:169
			return p.breakErr != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:169
			// _ = "end of CoverTab[73206]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:169
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:169
			_go_fuzz_dep_.CoverTab[73207]++

												p.closeDoneLocked()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:171
			// _ = "end of CoverTab[73207]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:172
			_go_fuzz_dep_.CoverTab[73208]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:172
			// _ = "end of CoverTab[73208]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:172
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:172
		// _ = "end of CoverTab[73205]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:173
		_go_fuzz_dep_.CoverTab[73209]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:173
		// _ = "end of CoverTab[73209]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:173
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:173
	// _ = "end of CoverTab[73203]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:173
	_go_fuzz_dep_.CoverTab[73204]++
										return p.donec
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:174
	// _ = "end of CoverTab[73204]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:175
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/pipe.go:175
var _ = _go_fuzz_dep_.CoverTab
