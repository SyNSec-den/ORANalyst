// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/pipe.go:5
package net

//line /snap/go/10455/src/net/pipe.go:5
import (
//line /snap/go/10455/src/net/pipe.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/pipe.go:5
)
//line /snap/go/10455/src/net/pipe.go:5
import (
//line /snap/go/10455/src/net/pipe.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/pipe.go:5
)

import (
	"io"
	"os"
	"sync"
	"time"
)

// pipeDeadline is an abstraction for handling timeouts.
type pipeDeadline struct {
	mu	sync.Mutex	// Guards timer and cancel
	timer	*time.Timer
	cancel	chan struct{}	// Must be non-nil
}

func makePipeDeadline() pipeDeadline {
//line /snap/go/10455/src/net/pipe.go:21
	_go_fuzz_dep_.CoverTab[7815]++
						return pipeDeadline{cancel: make(chan struct{})}
//line /snap/go/10455/src/net/pipe.go:22
	// _ = "end of CoverTab[7815]"
}

// set sets the point in time when the deadline will time out.
//line /snap/go/10455/src/net/pipe.go:25
// A timeout event is signaled by closing the channel returned by waiter.
//line /snap/go/10455/src/net/pipe.go:25
// Once a timeout has occurred, the deadline can be refreshed by specifying a
//line /snap/go/10455/src/net/pipe.go:25
// t value in the future.
//line /snap/go/10455/src/net/pipe.go:25
//
//line /snap/go/10455/src/net/pipe.go:25
// A zero value for t prevents timeout.
//line /snap/go/10455/src/net/pipe.go:31
func (d *pipeDeadline) set(t time.Time) {
//line /snap/go/10455/src/net/pipe.go:31
	_go_fuzz_dep_.CoverTab[7816]++
						d.mu.Lock()
						defer d.mu.Unlock()

						if d.timer != nil && func() bool {
//line /snap/go/10455/src/net/pipe.go:35
		_go_fuzz_dep_.CoverTab[7820]++
//line /snap/go/10455/src/net/pipe.go:35
		return !d.timer.Stop()
//line /snap/go/10455/src/net/pipe.go:35
		// _ = "end of CoverTab[7820]"
//line /snap/go/10455/src/net/pipe.go:35
	}() {
//line /snap/go/10455/src/net/pipe.go:35
		_go_fuzz_dep_.CoverTab[529506]++
//line /snap/go/10455/src/net/pipe.go:35
		_go_fuzz_dep_.CoverTab[7821]++
							<-d.cancel
//line /snap/go/10455/src/net/pipe.go:36
		// _ = "end of CoverTab[7821]"
	} else {
//line /snap/go/10455/src/net/pipe.go:37
		_go_fuzz_dep_.CoverTab[529507]++
//line /snap/go/10455/src/net/pipe.go:37
		_go_fuzz_dep_.CoverTab[7822]++
//line /snap/go/10455/src/net/pipe.go:37
		// _ = "end of CoverTab[7822]"
//line /snap/go/10455/src/net/pipe.go:37
	}
//line /snap/go/10455/src/net/pipe.go:37
	// _ = "end of CoverTab[7816]"
//line /snap/go/10455/src/net/pipe.go:37
	_go_fuzz_dep_.CoverTab[7817]++
						d.timer = nil

//line /snap/go/10455/src/net/pipe.go:41
	closed := isClosedChan(d.cancel)
	if t.IsZero() {
//line /snap/go/10455/src/net/pipe.go:42
		_go_fuzz_dep_.CoverTab[529508]++
//line /snap/go/10455/src/net/pipe.go:42
		_go_fuzz_dep_.CoverTab[7823]++
							if closed {
//line /snap/go/10455/src/net/pipe.go:43
			_go_fuzz_dep_.CoverTab[529510]++
//line /snap/go/10455/src/net/pipe.go:43
			_go_fuzz_dep_.CoverTab[7825]++
								d.cancel = make(chan struct{})
//line /snap/go/10455/src/net/pipe.go:44
			// _ = "end of CoverTab[7825]"
		} else {
//line /snap/go/10455/src/net/pipe.go:45
			_go_fuzz_dep_.CoverTab[529511]++
//line /snap/go/10455/src/net/pipe.go:45
			_go_fuzz_dep_.CoverTab[7826]++
//line /snap/go/10455/src/net/pipe.go:45
			// _ = "end of CoverTab[7826]"
//line /snap/go/10455/src/net/pipe.go:45
		}
//line /snap/go/10455/src/net/pipe.go:45
		// _ = "end of CoverTab[7823]"
//line /snap/go/10455/src/net/pipe.go:45
		_go_fuzz_dep_.CoverTab[7824]++
							return
//line /snap/go/10455/src/net/pipe.go:46
		// _ = "end of CoverTab[7824]"
	} else {
//line /snap/go/10455/src/net/pipe.go:47
		_go_fuzz_dep_.CoverTab[529509]++
//line /snap/go/10455/src/net/pipe.go:47
		_go_fuzz_dep_.CoverTab[7827]++
//line /snap/go/10455/src/net/pipe.go:47
		// _ = "end of CoverTab[7827]"
//line /snap/go/10455/src/net/pipe.go:47
	}
//line /snap/go/10455/src/net/pipe.go:47
	// _ = "end of CoverTab[7817]"
//line /snap/go/10455/src/net/pipe.go:47
	_go_fuzz_dep_.CoverTab[7818]++

//line /snap/go/10455/src/net/pipe.go:50
	if dur := time.Until(t); dur > 0 {
//line /snap/go/10455/src/net/pipe.go:50
		_go_fuzz_dep_.CoverTab[529512]++
//line /snap/go/10455/src/net/pipe.go:50
		_go_fuzz_dep_.CoverTab[7828]++
							if closed {
//line /snap/go/10455/src/net/pipe.go:51
			_go_fuzz_dep_.CoverTab[529514]++
//line /snap/go/10455/src/net/pipe.go:51
			_go_fuzz_dep_.CoverTab[7831]++
								d.cancel = make(chan struct{})
//line /snap/go/10455/src/net/pipe.go:52
			// _ = "end of CoverTab[7831]"
		} else {
//line /snap/go/10455/src/net/pipe.go:53
			_go_fuzz_dep_.CoverTab[529515]++
//line /snap/go/10455/src/net/pipe.go:53
			_go_fuzz_dep_.CoverTab[7832]++
//line /snap/go/10455/src/net/pipe.go:53
			// _ = "end of CoverTab[7832]"
//line /snap/go/10455/src/net/pipe.go:53
		}
//line /snap/go/10455/src/net/pipe.go:53
		// _ = "end of CoverTab[7828]"
//line /snap/go/10455/src/net/pipe.go:53
		_go_fuzz_dep_.CoverTab[7829]++
							d.timer = time.AfterFunc(dur, func() {
//line /snap/go/10455/src/net/pipe.go:54
			_go_fuzz_dep_.CoverTab[7833]++
								close(d.cancel)
//line /snap/go/10455/src/net/pipe.go:55
			// _ = "end of CoverTab[7833]"
		})
//line /snap/go/10455/src/net/pipe.go:56
		// _ = "end of CoverTab[7829]"
//line /snap/go/10455/src/net/pipe.go:56
		_go_fuzz_dep_.CoverTab[7830]++
							return
//line /snap/go/10455/src/net/pipe.go:57
		// _ = "end of CoverTab[7830]"
	} else {
//line /snap/go/10455/src/net/pipe.go:58
		_go_fuzz_dep_.CoverTab[529513]++
//line /snap/go/10455/src/net/pipe.go:58
		_go_fuzz_dep_.CoverTab[7834]++
//line /snap/go/10455/src/net/pipe.go:58
		// _ = "end of CoverTab[7834]"
//line /snap/go/10455/src/net/pipe.go:58
	}
//line /snap/go/10455/src/net/pipe.go:58
	// _ = "end of CoverTab[7818]"
//line /snap/go/10455/src/net/pipe.go:58
	_go_fuzz_dep_.CoverTab[7819]++

//line /snap/go/10455/src/net/pipe.go:61
	if !closed {
//line /snap/go/10455/src/net/pipe.go:61
		_go_fuzz_dep_.CoverTab[529516]++
//line /snap/go/10455/src/net/pipe.go:61
		_go_fuzz_dep_.CoverTab[7835]++
							close(d.cancel)
//line /snap/go/10455/src/net/pipe.go:62
		// _ = "end of CoverTab[7835]"
	} else {
//line /snap/go/10455/src/net/pipe.go:63
		_go_fuzz_dep_.CoverTab[529517]++
//line /snap/go/10455/src/net/pipe.go:63
		_go_fuzz_dep_.CoverTab[7836]++
//line /snap/go/10455/src/net/pipe.go:63
		// _ = "end of CoverTab[7836]"
//line /snap/go/10455/src/net/pipe.go:63
	}
//line /snap/go/10455/src/net/pipe.go:63
	// _ = "end of CoverTab[7819]"
}

// wait returns a channel that is closed when the deadline is exceeded.
func (d *pipeDeadline) wait() chan struct{} {
//line /snap/go/10455/src/net/pipe.go:67
	_go_fuzz_dep_.CoverTab[7837]++
						d.mu.Lock()
						defer d.mu.Unlock()
						return d.cancel
//line /snap/go/10455/src/net/pipe.go:70
	// _ = "end of CoverTab[7837]"
}

func isClosedChan(c <-chan struct{}) bool {
//line /snap/go/10455/src/net/pipe.go:73
	_go_fuzz_dep_.CoverTab[7838]++
						select {
	case <-c:
//line /snap/go/10455/src/net/pipe.go:75
		_go_fuzz_dep_.CoverTab[7839]++
							return true
//line /snap/go/10455/src/net/pipe.go:76
		// _ = "end of CoverTab[7839]"
	default:
//line /snap/go/10455/src/net/pipe.go:77
		_go_fuzz_dep_.CoverTab[7840]++
							return false
//line /snap/go/10455/src/net/pipe.go:78
		// _ = "end of CoverTab[7840]"
	}
//line /snap/go/10455/src/net/pipe.go:79
	// _ = "end of CoverTab[7838]"
}

type pipeAddr struct{}

func (pipeAddr) Network() string {
//line /snap/go/10455/src/net/pipe.go:84
	_go_fuzz_dep_.CoverTab[7841]++
//line /snap/go/10455/src/net/pipe.go:84
	return "pipe"
//line /snap/go/10455/src/net/pipe.go:84
	// _ = "end of CoverTab[7841]"
//line /snap/go/10455/src/net/pipe.go:84
}
func (pipeAddr) String() string {
//line /snap/go/10455/src/net/pipe.go:85
	_go_fuzz_dep_.CoverTab[7842]++
//line /snap/go/10455/src/net/pipe.go:85
	return "pipe"
//line /snap/go/10455/src/net/pipe.go:85
	// _ = "end of CoverTab[7842]"
//line /snap/go/10455/src/net/pipe.go:85
}

type pipe struct {
	wrMu	sync.Mutex	// Serialize Write operations

	// Used by local Read to interact with remote Write.
	// Successful receive on rdRx is always followed by send on rdTx.
	rdRx	<-chan []byte
	rdTx	chan<- int

	// Used by local Write to interact with remote Read.
	// Successful send on wrTx is always followed by receive on wrRx.
	wrTx	chan<- []byte
	wrRx	<-chan int

	once		sync.Once	// Protects closing localDone
	localDone	chan struct{}
	remoteDone	<-chan struct{}

	readDeadline	pipeDeadline
	writeDeadline	pipeDeadline
}

// Pipe creates a synchronous, in-memory, full duplex
//line /snap/go/10455/src/net/pipe.go:108
// network connection; both ends implement the Conn interface.
//line /snap/go/10455/src/net/pipe.go:108
// Reads on one end are matched with writes on the other,
//line /snap/go/10455/src/net/pipe.go:108
// copying data directly between the two; there is no internal
//line /snap/go/10455/src/net/pipe.go:108
// buffering.
//line /snap/go/10455/src/net/pipe.go:113
func Pipe() (Conn, Conn) {
//line /snap/go/10455/src/net/pipe.go:113
	_go_fuzz_dep_.CoverTab[7843]++
						cb1 := make(chan []byte)
						cb2 := make(chan []byte)
						cn1 := make(chan int)
						cn2 := make(chan int)
						done1 := make(chan struct{})
						done2 := make(chan struct{})

						p1 := &pipe{
		rdRx:	cb1, rdTx: cn1,
		wrTx:	cb2, wrRx: cn2,
		localDone:	done1, remoteDone: done2,
		readDeadline:	makePipeDeadline(),
		writeDeadline:	makePipeDeadline(),
	}
	p2 := &pipe{
		rdRx:	cb2, rdTx: cn2,
		wrTx:	cb1, wrRx: cn1,
		localDone:	done2, remoteDone: done1,
		readDeadline:	makePipeDeadline(),
		writeDeadline:	makePipeDeadline(),
	}
						return p1, p2
//line /snap/go/10455/src/net/pipe.go:135
	// _ = "end of CoverTab[7843]"
}

func (*pipe) LocalAddr() Addr {
//line /snap/go/10455/src/net/pipe.go:138
	_go_fuzz_dep_.CoverTab[7844]++
//line /snap/go/10455/src/net/pipe.go:138
	return pipeAddr{}
//line /snap/go/10455/src/net/pipe.go:138
	// _ = "end of CoverTab[7844]"
//line /snap/go/10455/src/net/pipe.go:138
}
func (*pipe) RemoteAddr() Addr {
//line /snap/go/10455/src/net/pipe.go:139
	_go_fuzz_dep_.CoverTab[7845]++
//line /snap/go/10455/src/net/pipe.go:139
	return pipeAddr{}
//line /snap/go/10455/src/net/pipe.go:139
	// _ = "end of CoverTab[7845]"
//line /snap/go/10455/src/net/pipe.go:139
}

func (p *pipe) Read(b []byte) (int, error) {
//line /snap/go/10455/src/net/pipe.go:141
	_go_fuzz_dep_.CoverTab[7846]++
						n, err := p.read(b)
						if err != nil && func() bool {
//line /snap/go/10455/src/net/pipe.go:143
		_go_fuzz_dep_.CoverTab[7848]++
//line /snap/go/10455/src/net/pipe.go:143
		return err != io.EOF
//line /snap/go/10455/src/net/pipe.go:143
		// _ = "end of CoverTab[7848]"
//line /snap/go/10455/src/net/pipe.go:143
	}() && func() bool {
//line /snap/go/10455/src/net/pipe.go:143
		_go_fuzz_dep_.CoverTab[7849]++
//line /snap/go/10455/src/net/pipe.go:143
		return err != io.ErrClosedPipe
//line /snap/go/10455/src/net/pipe.go:143
		// _ = "end of CoverTab[7849]"
//line /snap/go/10455/src/net/pipe.go:143
	}() {
//line /snap/go/10455/src/net/pipe.go:143
		_go_fuzz_dep_.CoverTab[529518]++
//line /snap/go/10455/src/net/pipe.go:143
		_go_fuzz_dep_.CoverTab[7850]++
							err = &OpError{Op: "read", Net: "pipe", Err: err}
//line /snap/go/10455/src/net/pipe.go:144
		// _ = "end of CoverTab[7850]"
	} else {
//line /snap/go/10455/src/net/pipe.go:145
		_go_fuzz_dep_.CoverTab[529519]++
//line /snap/go/10455/src/net/pipe.go:145
		_go_fuzz_dep_.CoverTab[7851]++
//line /snap/go/10455/src/net/pipe.go:145
		// _ = "end of CoverTab[7851]"
//line /snap/go/10455/src/net/pipe.go:145
	}
//line /snap/go/10455/src/net/pipe.go:145
	// _ = "end of CoverTab[7846]"
//line /snap/go/10455/src/net/pipe.go:145
	_go_fuzz_dep_.CoverTab[7847]++
						return n, err
//line /snap/go/10455/src/net/pipe.go:146
	// _ = "end of CoverTab[7847]"
}

func (p *pipe) read(b []byte) (n int, err error) {
//line /snap/go/10455/src/net/pipe.go:149
	_go_fuzz_dep_.CoverTab[7852]++
						switch {
	case isClosedChan(p.localDone):
//line /snap/go/10455/src/net/pipe.go:151
		_go_fuzz_dep_.CoverTab[529520]++
//line /snap/go/10455/src/net/pipe.go:151
		_go_fuzz_dep_.CoverTab[7854]++
							return 0, io.ErrClosedPipe
//line /snap/go/10455/src/net/pipe.go:152
		// _ = "end of CoverTab[7854]"
	case isClosedChan(p.remoteDone):
//line /snap/go/10455/src/net/pipe.go:153
		_go_fuzz_dep_.CoverTab[529521]++
//line /snap/go/10455/src/net/pipe.go:153
		_go_fuzz_dep_.CoverTab[7855]++
							return 0, io.EOF
//line /snap/go/10455/src/net/pipe.go:154
		// _ = "end of CoverTab[7855]"
	case isClosedChan(p.readDeadline.wait()):
//line /snap/go/10455/src/net/pipe.go:155
		_go_fuzz_dep_.CoverTab[529522]++
//line /snap/go/10455/src/net/pipe.go:155
		_go_fuzz_dep_.CoverTab[7856]++
							return 0, os.ErrDeadlineExceeded
//line /snap/go/10455/src/net/pipe.go:156
		// _ = "end of CoverTab[7856]"
//line /snap/go/10455/src/net/pipe.go:156
	default:
//line /snap/go/10455/src/net/pipe.go:156
		_go_fuzz_dep_.CoverTab[529523]++
//line /snap/go/10455/src/net/pipe.go:156
		_go_fuzz_dep_.CoverTab[7857]++
//line /snap/go/10455/src/net/pipe.go:156
		// _ = "end of CoverTab[7857]"
	}
//line /snap/go/10455/src/net/pipe.go:157
	// _ = "end of CoverTab[7852]"
//line /snap/go/10455/src/net/pipe.go:157
	_go_fuzz_dep_.CoverTab[7853]++

						select {
	case bw := <-p.rdRx:
//line /snap/go/10455/src/net/pipe.go:160
		_go_fuzz_dep_.CoverTab[7858]++
							nr := copy(b, bw)
							p.rdTx <- nr
							return nr, nil
//line /snap/go/10455/src/net/pipe.go:163
		// _ = "end of CoverTab[7858]"
	case <-p.localDone:
//line /snap/go/10455/src/net/pipe.go:164
		_go_fuzz_dep_.CoverTab[7859]++
							return 0, io.ErrClosedPipe
//line /snap/go/10455/src/net/pipe.go:165
		// _ = "end of CoverTab[7859]"
	case <-p.remoteDone:
//line /snap/go/10455/src/net/pipe.go:166
		_go_fuzz_dep_.CoverTab[7860]++
							return 0, io.EOF
//line /snap/go/10455/src/net/pipe.go:167
		// _ = "end of CoverTab[7860]"
	case <-p.readDeadline.wait():
//line /snap/go/10455/src/net/pipe.go:168
		_go_fuzz_dep_.CoverTab[7861]++
							return 0, os.ErrDeadlineExceeded
//line /snap/go/10455/src/net/pipe.go:169
		// _ = "end of CoverTab[7861]"
	}
//line /snap/go/10455/src/net/pipe.go:170
	// _ = "end of CoverTab[7853]"
}

func (p *pipe) Write(b []byte) (int, error) {
//line /snap/go/10455/src/net/pipe.go:173
	_go_fuzz_dep_.CoverTab[7862]++
						n, err := p.write(b)
						if err != nil && func() bool {
//line /snap/go/10455/src/net/pipe.go:175
		_go_fuzz_dep_.CoverTab[7864]++
//line /snap/go/10455/src/net/pipe.go:175
		return err != io.ErrClosedPipe
//line /snap/go/10455/src/net/pipe.go:175
		// _ = "end of CoverTab[7864]"
//line /snap/go/10455/src/net/pipe.go:175
	}() {
//line /snap/go/10455/src/net/pipe.go:175
		_go_fuzz_dep_.CoverTab[529524]++
//line /snap/go/10455/src/net/pipe.go:175
		_go_fuzz_dep_.CoverTab[7865]++
							err = &OpError{Op: "write", Net: "pipe", Err: err}
//line /snap/go/10455/src/net/pipe.go:176
		// _ = "end of CoverTab[7865]"
	} else {
//line /snap/go/10455/src/net/pipe.go:177
		_go_fuzz_dep_.CoverTab[529525]++
//line /snap/go/10455/src/net/pipe.go:177
		_go_fuzz_dep_.CoverTab[7866]++
//line /snap/go/10455/src/net/pipe.go:177
		// _ = "end of CoverTab[7866]"
//line /snap/go/10455/src/net/pipe.go:177
	}
//line /snap/go/10455/src/net/pipe.go:177
	// _ = "end of CoverTab[7862]"
//line /snap/go/10455/src/net/pipe.go:177
	_go_fuzz_dep_.CoverTab[7863]++
						return n, err
//line /snap/go/10455/src/net/pipe.go:178
	// _ = "end of CoverTab[7863]"
}

func (p *pipe) write(b []byte) (n int, err error) {
//line /snap/go/10455/src/net/pipe.go:181
	_go_fuzz_dep_.CoverTab[7867]++
						switch {
	case isClosedChan(p.localDone):
//line /snap/go/10455/src/net/pipe.go:183
		_go_fuzz_dep_.CoverTab[529526]++
//line /snap/go/10455/src/net/pipe.go:183
		_go_fuzz_dep_.CoverTab[7870]++
							return 0, io.ErrClosedPipe
//line /snap/go/10455/src/net/pipe.go:184
		// _ = "end of CoverTab[7870]"
	case isClosedChan(p.remoteDone):
//line /snap/go/10455/src/net/pipe.go:185
		_go_fuzz_dep_.CoverTab[529527]++
//line /snap/go/10455/src/net/pipe.go:185
		_go_fuzz_dep_.CoverTab[7871]++
							return 0, io.ErrClosedPipe
//line /snap/go/10455/src/net/pipe.go:186
		// _ = "end of CoverTab[7871]"
	case isClosedChan(p.writeDeadline.wait()):
//line /snap/go/10455/src/net/pipe.go:187
		_go_fuzz_dep_.CoverTab[529528]++
//line /snap/go/10455/src/net/pipe.go:187
		_go_fuzz_dep_.CoverTab[7872]++
							return 0, os.ErrDeadlineExceeded
//line /snap/go/10455/src/net/pipe.go:188
		// _ = "end of CoverTab[7872]"
//line /snap/go/10455/src/net/pipe.go:188
	default:
//line /snap/go/10455/src/net/pipe.go:188
		_go_fuzz_dep_.CoverTab[529529]++
//line /snap/go/10455/src/net/pipe.go:188
		_go_fuzz_dep_.CoverTab[7873]++
//line /snap/go/10455/src/net/pipe.go:188
		// _ = "end of CoverTab[7873]"
	}
//line /snap/go/10455/src/net/pipe.go:189
	// _ = "end of CoverTab[7867]"
//line /snap/go/10455/src/net/pipe.go:189
	_go_fuzz_dep_.CoverTab[7868]++

						p.wrMu.Lock()
						defer p.wrMu.Unlock()
//line /snap/go/10455/src/net/pipe.go:192
	_go_fuzz_dep_.CoverTab[786740] = 0
						for once := true; once || func() bool {
//line /snap/go/10455/src/net/pipe.go:193
		_go_fuzz_dep_.CoverTab[7874]++
//line /snap/go/10455/src/net/pipe.go:193
		return len(b) > 0
//line /snap/go/10455/src/net/pipe.go:193
		// _ = "end of CoverTab[7874]"
//line /snap/go/10455/src/net/pipe.go:193
	}(); once = false {
//line /snap/go/10455/src/net/pipe.go:193
		if _go_fuzz_dep_.CoverTab[786740] == 0 {
//line /snap/go/10455/src/net/pipe.go:193
			_go_fuzz_dep_.CoverTab[529536]++
//line /snap/go/10455/src/net/pipe.go:193
		} else {
//line /snap/go/10455/src/net/pipe.go:193
			_go_fuzz_dep_.CoverTab[529537]++
//line /snap/go/10455/src/net/pipe.go:193
		}
//line /snap/go/10455/src/net/pipe.go:193
		_go_fuzz_dep_.CoverTab[786740] = 1
//line /snap/go/10455/src/net/pipe.go:193
		_go_fuzz_dep_.CoverTab[7875]++
							select {
		case p.wrTx <- b:
//line /snap/go/10455/src/net/pipe.go:195
			_go_fuzz_dep_.CoverTab[7876]++
								nw := <-p.wrRx
								b = b[nw:]
								n += nw
//line /snap/go/10455/src/net/pipe.go:198
			// _ = "end of CoverTab[7876]"
		case <-p.localDone:
//line /snap/go/10455/src/net/pipe.go:199
			_go_fuzz_dep_.CoverTab[7877]++
								return n, io.ErrClosedPipe
//line /snap/go/10455/src/net/pipe.go:200
			// _ = "end of CoverTab[7877]"
		case <-p.remoteDone:
//line /snap/go/10455/src/net/pipe.go:201
			_go_fuzz_dep_.CoverTab[7878]++
								return n, io.ErrClosedPipe
//line /snap/go/10455/src/net/pipe.go:202
			// _ = "end of CoverTab[7878]"
		case <-p.writeDeadline.wait():
//line /snap/go/10455/src/net/pipe.go:203
			_go_fuzz_dep_.CoverTab[7879]++
								return n, os.ErrDeadlineExceeded
//line /snap/go/10455/src/net/pipe.go:204
			// _ = "end of CoverTab[7879]"
		}
//line /snap/go/10455/src/net/pipe.go:205
		// _ = "end of CoverTab[7875]"
	}
//line /snap/go/10455/src/net/pipe.go:206
	if _go_fuzz_dep_.CoverTab[786740] == 0 {
//line /snap/go/10455/src/net/pipe.go:206
		_go_fuzz_dep_.CoverTab[529538]++
//line /snap/go/10455/src/net/pipe.go:206
	} else {
//line /snap/go/10455/src/net/pipe.go:206
		_go_fuzz_dep_.CoverTab[529539]++
//line /snap/go/10455/src/net/pipe.go:206
	}
//line /snap/go/10455/src/net/pipe.go:206
	// _ = "end of CoverTab[7868]"
//line /snap/go/10455/src/net/pipe.go:206
	_go_fuzz_dep_.CoverTab[7869]++
						return n, nil
//line /snap/go/10455/src/net/pipe.go:207
	// _ = "end of CoverTab[7869]"
}

func (p *pipe) SetDeadline(t time.Time) error {
//line /snap/go/10455/src/net/pipe.go:210
	_go_fuzz_dep_.CoverTab[7880]++
						if isClosedChan(p.localDone) || func() bool {
//line /snap/go/10455/src/net/pipe.go:211
		_go_fuzz_dep_.CoverTab[7882]++
//line /snap/go/10455/src/net/pipe.go:211
		return isClosedChan(p.remoteDone)
//line /snap/go/10455/src/net/pipe.go:211
		// _ = "end of CoverTab[7882]"
//line /snap/go/10455/src/net/pipe.go:211
	}() {
//line /snap/go/10455/src/net/pipe.go:211
		_go_fuzz_dep_.CoverTab[529530]++
//line /snap/go/10455/src/net/pipe.go:211
		_go_fuzz_dep_.CoverTab[7883]++
							return io.ErrClosedPipe
//line /snap/go/10455/src/net/pipe.go:212
		// _ = "end of CoverTab[7883]"
	} else {
//line /snap/go/10455/src/net/pipe.go:213
		_go_fuzz_dep_.CoverTab[529531]++
//line /snap/go/10455/src/net/pipe.go:213
		_go_fuzz_dep_.CoverTab[7884]++
//line /snap/go/10455/src/net/pipe.go:213
		// _ = "end of CoverTab[7884]"
//line /snap/go/10455/src/net/pipe.go:213
	}
//line /snap/go/10455/src/net/pipe.go:213
	// _ = "end of CoverTab[7880]"
//line /snap/go/10455/src/net/pipe.go:213
	_go_fuzz_dep_.CoverTab[7881]++
						p.readDeadline.set(t)
						p.writeDeadline.set(t)
						return nil
//line /snap/go/10455/src/net/pipe.go:216
	// _ = "end of CoverTab[7881]"
}

func (p *pipe) SetReadDeadline(t time.Time) error {
//line /snap/go/10455/src/net/pipe.go:219
	_go_fuzz_dep_.CoverTab[7885]++
						if isClosedChan(p.localDone) || func() bool {
//line /snap/go/10455/src/net/pipe.go:220
		_go_fuzz_dep_.CoverTab[7887]++
//line /snap/go/10455/src/net/pipe.go:220
		return isClosedChan(p.remoteDone)
//line /snap/go/10455/src/net/pipe.go:220
		// _ = "end of CoverTab[7887]"
//line /snap/go/10455/src/net/pipe.go:220
	}() {
//line /snap/go/10455/src/net/pipe.go:220
		_go_fuzz_dep_.CoverTab[529532]++
//line /snap/go/10455/src/net/pipe.go:220
		_go_fuzz_dep_.CoverTab[7888]++
							return io.ErrClosedPipe
//line /snap/go/10455/src/net/pipe.go:221
		// _ = "end of CoverTab[7888]"
	} else {
//line /snap/go/10455/src/net/pipe.go:222
		_go_fuzz_dep_.CoverTab[529533]++
//line /snap/go/10455/src/net/pipe.go:222
		_go_fuzz_dep_.CoverTab[7889]++
//line /snap/go/10455/src/net/pipe.go:222
		// _ = "end of CoverTab[7889]"
//line /snap/go/10455/src/net/pipe.go:222
	}
//line /snap/go/10455/src/net/pipe.go:222
	// _ = "end of CoverTab[7885]"
//line /snap/go/10455/src/net/pipe.go:222
	_go_fuzz_dep_.CoverTab[7886]++
						p.readDeadline.set(t)
						return nil
//line /snap/go/10455/src/net/pipe.go:224
	// _ = "end of CoverTab[7886]"
}

func (p *pipe) SetWriteDeadline(t time.Time) error {
//line /snap/go/10455/src/net/pipe.go:227
	_go_fuzz_dep_.CoverTab[7890]++
						if isClosedChan(p.localDone) || func() bool {
//line /snap/go/10455/src/net/pipe.go:228
		_go_fuzz_dep_.CoverTab[7892]++
//line /snap/go/10455/src/net/pipe.go:228
		return isClosedChan(p.remoteDone)
//line /snap/go/10455/src/net/pipe.go:228
		// _ = "end of CoverTab[7892]"
//line /snap/go/10455/src/net/pipe.go:228
	}() {
//line /snap/go/10455/src/net/pipe.go:228
		_go_fuzz_dep_.CoverTab[529534]++
//line /snap/go/10455/src/net/pipe.go:228
		_go_fuzz_dep_.CoverTab[7893]++
							return io.ErrClosedPipe
//line /snap/go/10455/src/net/pipe.go:229
		// _ = "end of CoverTab[7893]"
	} else {
//line /snap/go/10455/src/net/pipe.go:230
		_go_fuzz_dep_.CoverTab[529535]++
//line /snap/go/10455/src/net/pipe.go:230
		_go_fuzz_dep_.CoverTab[7894]++
//line /snap/go/10455/src/net/pipe.go:230
		// _ = "end of CoverTab[7894]"
//line /snap/go/10455/src/net/pipe.go:230
	}
//line /snap/go/10455/src/net/pipe.go:230
	// _ = "end of CoverTab[7890]"
//line /snap/go/10455/src/net/pipe.go:230
	_go_fuzz_dep_.CoverTab[7891]++
						p.writeDeadline.set(t)
						return nil
//line /snap/go/10455/src/net/pipe.go:232
	// _ = "end of CoverTab[7891]"
}

func (p *pipe) Close() error {
//line /snap/go/10455/src/net/pipe.go:235
	_go_fuzz_dep_.CoverTab[7895]++
						p.once.Do(func() { _go_fuzz_dep_.CoverTab[7897]++; close(p.localDone); // _ = "end of CoverTab[7897]" })
//line /snap/go/10455/src/net/pipe.go:236
	// _ = "end of CoverTab[7895]"
//line /snap/go/10455/src/net/pipe.go:236
	_go_fuzz_dep_.CoverTab[7896]++
						return nil
//line /snap/go/10455/src/net/pipe.go:237
	// _ = "end of CoverTab[7896]"
}

//line /snap/go/10455/src/net/pipe.go:238
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/pipe.go:238
var _ = _go_fuzz_dep_.CoverTab
