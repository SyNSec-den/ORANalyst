// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/pipe.go:5
package net

//line /usr/local/go/src/net/pipe.go:5
import (
//line /usr/local/go/src/net/pipe.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/pipe.go:5
)
//line /usr/local/go/src/net/pipe.go:5
import (
//line /usr/local/go/src/net/pipe.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/pipe.go:5
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
//line /usr/local/go/src/net/pipe.go:21
	_go_fuzz_dep_.CoverTab[15925]++
						return pipeDeadline{cancel: make(chan struct{})}
//line /usr/local/go/src/net/pipe.go:22
	// _ = "end of CoverTab[15925]"
}

// set sets the point in time when the deadline will time out.
//line /usr/local/go/src/net/pipe.go:25
// A timeout event is signaled by closing the channel returned by waiter.
//line /usr/local/go/src/net/pipe.go:25
// Once a timeout has occurred, the deadline can be refreshed by specifying a
//line /usr/local/go/src/net/pipe.go:25
// t value in the future.
//line /usr/local/go/src/net/pipe.go:25
//
//line /usr/local/go/src/net/pipe.go:25
// A zero value for t prevents timeout.
//line /usr/local/go/src/net/pipe.go:31
func (d *pipeDeadline) set(t time.Time) {
//line /usr/local/go/src/net/pipe.go:31
	_go_fuzz_dep_.CoverTab[15926]++
						d.mu.Lock()
						defer d.mu.Unlock()

						if d.timer != nil && func() bool {
//line /usr/local/go/src/net/pipe.go:35
		_go_fuzz_dep_.CoverTab[15930]++
//line /usr/local/go/src/net/pipe.go:35
		return !d.timer.Stop()
//line /usr/local/go/src/net/pipe.go:35
		// _ = "end of CoverTab[15930]"
//line /usr/local/go/src/net/pipe.go:35
	}() {
//line /usr/local/go/src/net/pipe.go:35
		_go_fuzz_dep_.CoverTab[15931]++
							<-d.cancel
//line /usr/local/go/src/net/pipe.go:36
		// _ = "end of CoverTab[15931]"
	} else {
//line /usr/local/go/src/net/pipe.go:37
		_go_fuzz_dep_.CoverTab[15932]++
//line /usr/local/go/src/net/pipe.go:37
		// _ = "end of CoverTab[15932]"
//line /usr/local/go/src/net/pipe.go:37
	}
//line /usr/local/go/src/net/pipe.go:37
	// _ = "end of CoverTab[15926]"
//line /usr/local/go/src/net/pipe.go:37
	_go_fuzz_dep_.CoverTab[15927]++
						d.timer = nil

//line /usr/local/go/src/net/pipe.go:41
	closed := isClosedChan(d.cancel)
	if t.IsZero() {
//line /usr/local/go/src/net/pipe.go:42
		_go_fuzz_dep_.CoverTab[15933]++
							if closed {
//line /usr/local/go/src/net/pipe.go:43
			_go_fuzz_dep_.CoverTab[15935]++
								d.cancel = make(chan struct{})
//line /usr/local/go/src/net/pipe.go:44
			// _ = "end of CoverTab[15935]"
		} else {
//line /usr/local/go/src/net/pipe.go:45
			_go_fuzz_dep_.CoverTab[15936]++
//line /usr/local/go/src/net/pipe.go:45
			// _ = "end of CoverTab[15936]"
//line /usr/local/go/src/net/pipe.go:45
		}
//line /usr/local/go/src/net/pipe.go:45
		// _ = "end of CoverTab[15933]"
//line /usr/local/go/src/net/pipe.go:45
		_go_fuzz_dep_.CoverTab[15934]++
							return
//line /usr/local/go/src/net/pipe.go:46
		// _ = "end of CoverTab[15934]"
	} else {
//line /usr/local/go/src/net/pipe.go:47
		_go_fuzz_dep_.CoverTab[15937]++
//line /usr/local/go/src/net/pipe.go:47
		// _ = "end of CoverTab[15937]"
//line /usr/local/go/src/net/pipe.go:47
	}
//line /usr/local/go/src/net/pipe.go:47
	// _ = "end of CoverTab[15927]"
//line /usr/local/go/src/net/pipe.go:47
	_go_fuzz_dep_.CoverTab[15928]++

//line /usr/local/go/src/net/pipe.go:50
	if dur := time.Until(t); dur > 0 {
//line /usr/local/go/src/net/pipe.go:50
		_go_fuzz_dep_.CoverTab[15938]++
							if closed {
//line /usr/local/go/src/net/pipe.go:51
			_go_fuzz_dep_.CoverTab[15941]++
								d.cancel = make(chan struct{})
//line /usr/local/go/src/net/pipe.go:52
			// _ = "end of CoverTab[15941]"
		} else {
//line /usr/local/go/src/net/pipe.go:53
			_go_fuzz_dep_.CoverTab[15942]++
//line /usr/local/go/src/net/pipe.go:53
			// _ = "end of CoverTab[15942]"
//line /usr/local/go/src/net/pipe.go:53
		}
//line /usr/local/go/src/net/pipe.go:53
		// _ = "end of CoverTab[15938]"
//line /usr/local/go/src/net/pipe.go:53
		_go_fuzz_dep_.CoverTab[15939]++
							d.timer = time.AfterFunc(dur, func() {
//line /usr/local/go/src/net/pipe.go:54
			_go_fuzz_dep_.CoverTab[15943]++
								close(d.cancel)
//line /usr/local/go/src/net/pipe.go:55
			// _ = "end of CoverTab[15943]"
		})
//line /usr/local/go/src/net/pipe.go:56
		// _ = "end of CoverTab[15939]"
//line /usr/local/go/src/net/pipe.go:56
		_go_fuzz_dep_.CoverTab[15940]++
							return
//line /usr/local/go/src/net/pipe.go:57
		// _ = "end of CoverTab[15940]"
	} else {
//line /usr/local/go/src/net/pipe.go:58
		_go_fuzz_dep_.CoverTab[15944]++
//line /usr/local/go/src/net/pipe.go:58
		// _ = "end of CoverTab[15944]"
//line /usr/local/go/src/net/pipe.go:58
	}
//line /usr/local/go/src/net/pipe.go:58
	// _ = "end of CoverTab[15928]"
//line /usr/local/go/src/net/pipe.go:58
	_go_fuzz_dep_.CoverTab[15929]++

//line /usr/local/go/src/net/pipe.go:61
	if !closed {
//line /usr/local/go/src/net/pipe.go:61
		_go_fuzz_dep_.CoverTab[15945]++
							close(d.cancel)
//line /usr/local/go/src/net/pipe.go:62
		// _ = "end of CoverTab[15945]"
	} else {
//line /usr/local/go/src/net/pipe.go:63
		_go_fuzz_dep_.CoverTab[15946]++
//line /usr/local/go/src/net/pipe.go:63
		// _ = "end of CoverTab[15946]"
//line /usr/local/go/src/net/pipe.go:63
	}
//line /usr/local/go/src/net/pipe.go:63
	// _ = "end of CoverTab[15929]"
}

// wait returns a channel that is closed when the deadline is exceeded.
func (d *pipeDeadline) wait() chan struct{} {
//line /usr/local/go/src/net/pipe.go:67
	_go_fuzz_dep_.CoverTab[15947]++
						d.mu.Lock()
						defer d.mu.Unlock()
						return d.cancel
//line /usr/local/go/src/net/pipe.go:70
	// _ = "end of CoverTab[15947]"
}

func isClosedChan(c <-chan struct{}) bool {
//line /usr/local/go/src/net/pipe.go:73
	_go_fuzz_dep_.CoverTab[15948]++
						select {
	case <-c:
//line /usr/local/go/src/net/pipe.go:75
		_go_fuzz_dep_.CoverTab[15949]++
							return true
//line /usr/local/go/src/net/pipe.go:76
		// _ = "end of CoverTab[15949]"
	default:
//line /usr/local/go/src/net/pipe.go:77
		_go_fuzz_dep_.CoverTab[15950]++
							return false
//line /usr/local/go/src/net/pipe.go:78
		// _ = "end of CoverTab[15950]"
	}
//line /usr/local/go/src/net/pipe.go:79
	// _ = "end of CoverTab[15948]"
}

type pipeAddr struct{}

func (pipeAddr) Network() string {
//line /usr/local/go/src/net/pipe.go:84
	_go_fuzz_dep_.CoverTab[15951]++
//line /usr/local/go/src/net/pipe.go:84
	return "pipe"
//line /usr/local/go/src/net/pipe.go:84
	// _ = "end of CoverTab[15951]"
//line /usr/local/go/src/net/pipe.go:84
}
func (pipeAddr) String() string {
//line /usr/local/go/src/net/pipe.go:85
	_go_fuzz_dep_.CoverTab[15952]++
//line /usr/local/go/src/net/pipe.go:85
	return "pipe"
//line /usr/local/go/src/net/pipe.go:85
	// _ = "end of CoverTab[15952]"
//line /usr/local/go/src/net/pipe.go:85
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
//line /usr/local/go/src/net/pipe.go:108
// network connection; both ends implement the Conn interface.
//line /usr/local/go/src/net/pipe.go:108
// Reads on one end are matched with writes on the other,
//line /usr/local/go/src/net/pipe.go:108
// copying data directly between the two; there is no internal
//line /usr/local/go/src/net/pipe.go:108
// buffering.
//line /usr/local/go/src/net/pipe.go:113
func Pipe() (Conn, Conn) {
//line /usr/local/go/src/net/pipe.go:113
	_go_fuzz_dep_.CoverTab[15953]++
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
//line /usr/local/go/src/net/pipe.go:135
	// _ = "end of CoverTab[15953]"
}

func (*pipe) LocalAddr() Addr {
//line /usr/local/go/src/net/pipe.go:138
	_go_fuzz_dep_.CoverTab[15954]++
//line /usr/local/go/src/net/pipe.go:138
	return pipeAddr{}
//line /usr/local/go/src/net/pipe.go:138
	// _ = "end of CoverTab[15954]"
//line /usr/local/go/src/net/pipe.go:138
}
func (*pipe) RemoteAddr() Addr {
//line /usr/local/go/src/net/pipe.go:139
	_go_fuzz_dep_.CoverTab[15955]++
//line /usr/local/go/src/net/pipe.go:139
	return pipeAddr{}
//line /usr/local/go/src/net/pipe.go:139
	// _ = "end of CoverTab[15955]"
//line /usr/local/go/src/net/pipe.go:139
}

func (p *pipe) Read(b []byte) (int, error) {
//line /usr/local/go/src/net/pipe.go:141
	_go_fuzz_dep_.CoverTab[15956]++
						n, err := p.read(b)
						if err != nil && func() bool {
//line /usr/local/go/src/net/pipe.go:143
		_go_fuzz_dep_.CoverTab[15958]++
//line /usr/local/go/src/net/pipe.go:143
		return err != io.EOF
//line /usr/local/go/src/net/pipe.go:143
		// _ = "end of CoverTab[15958]"
//line /usr/local/go/src/net/pipe.go:143
	}() && func() bool {
//line /usr/local/go/src/net/pipe.go:143
		_go_fuzz_dep_.CoverTab[15959]++
//line /usr/local/go/src/net/pipe.go:143
		return err != io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:143
		// _ = "end of CoverTab[15959]"
//line /usr/local/go/src/net/pipe.go:143
	}() {
//line /usr/local/go/src/net/pipe.go:143
		_go_fuzz_dep_.CoverTab[15960]++
							err = &OpError{Op: "read", Net: "pipe", Err: err}
//line /usr/local/go/src/net/pipe.go:144
		// _ = "end of CoverTab[15960]"
	} else {
//line /usr/local/go/src/net/pipe.go:145
		_go_fuzz_dep_.CoverTab[15961]++
//line /usr/local/go/src/net/pipe.go:145
		// _ = "end of CoverTab[15961]"
//line /usr/local/go/src/net/pipe.go:145
	}
//line /usr/local/go/src/net/pipe.go:145
	// _ = "end of CoverTab[15956]"
//line /usr/local/go/src/net/pipe.go:145
	_go_fuzz_dep_.CoverTab[15957]++
						return n, err
//line /usr/local/go/src/net/pipe.go:146
	// _ = "end of CoverTab[15957]"
}

func (p *pipe) read(b []byte) (n int, err error) {
//line /usr/local/go/src/net/pipe.go:149
	_go_fuzz_dep_.CoverTab[15962]++
						switch {
	case isClosedChan(p.localDone):
//line /usr/local/go/src/net/pipe.go:151
		_go_fuzz_dep_.CoverTab[15964]++
							return 0, io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:152
		// _ = "end of CoverTab[15964]"
	case isClosedChan(p.remoteDone):
//line /usr/local/go/src/net/pipe.go:153
		_go_fuzz_dep_.CoverTab[15965]++
							return 0, io.EOF
//line /usr/local/go/src/net/pipe.go:154
		// _ = "end of CoverTab[15965]"
	case isClosedChan(p.readDeadline.wait()):
//line /usr/local/go/src/net/pipe.go:155
		_go_fuzz_dep_.CoverTab[15966]++
							return 0, os.ErrDeadlineExceeded
//line /usr/local/go/src/net/pipe.go:156
		// _ = "end of CoverTab[15966]"
//line /usr/local/go/src/net/pipe.go:156
	default:
//line /usr/local/go/src/net/pipe.go:156
		_go_fuzz_dep_.CoverTab[15967]++
//line /usr/local/go/src/net/pipe.go:156
		// _ = "end of CoverTab[15967]"
	}
//line /usr/local/go/src/net/pipe.go:157
	// _ = "end of CoverTab[15962]"
//line /usr/local/go/src/net/pipe.go:157
	_go_fuzz_dep_.CoverTab[15963]++

						select {
	case bw := <-p.rdRx:
//line /usr/local/go/src/net/pipe.go:160
		_go_fuzz_dep_.CoverTab[15968]++
							nr := copy(b, bw)
							p.rdTx <- nr
							return nr, nil
//line /usr/local/go/src/net/pipe.go:163
		// _ = "end of CoverTab[15968]"
	case <-p.localDone:
//line /usr/local/go/src/net/pipe.go:164
		_go_fuzz_dep_.CoverTab[15969]++
							return 0, io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:165
		// _ = "end of CoverTab[15969]"
	case <-p.remoteDone:
//line /usr/local/go/src/net/pipe.go:166
		_go_fuzz_dep_.CoverTab[15970]++
							return 0, io.EOF
//line /usr/local/go/src/net/pipe.go:167
		// _ = "end of CoverTab[15970]"
	case <-p.readDeadline.wait():
//line /usr/local/go/src/net/pipe.go:168
		_go_fuzz_dep_.CoverTab[15971]++
							return 0, os.ErrDeadlineExceeded
//line /usr/local/go/src/net/pipe.go:169
		// _ = "end of CoverTab[15971]"
	}
//line /usr/local/go/src/net/pipe.go:170
	// _ = "end of CoverTab[15963]"
}

func (p *pipe) Write(b []byte) (int, error) {
//line /usr/local/go/src/net/pipe.go:173
	_go_fuzz_dep_.CoverTab[15972]++
						n, err := p.write(b)
						if err != nil && func() bool {
//line /usr/local/go/src/net/pipe.go:175
		_go_fuzz_dep_.CoverTab[15974]++
//line /usr/local/go/src/net/pipe.go:175
		return err != io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:175
		// _ = "end of CoverTab[15974]"
//line /usr/local/go/src/net/pipe.go:175
	}() {
//line /usr/local/go/src/net/pipe.go:175
		_go_fuzz_dep_.CoverTab[15975]++
							err = &OpError{Op: "write", Net: "pipe", Err: err}
//line /usr/local/go/src/net/pipe.go:176
		// _ = "end of CoverTab[15975]"
	} else {
//line /usr/local/go/src/net/pipe.go:177
		_go_fuzz_dep_.CoverTab[15976]++
//line /usr/local/go/src/net/pipe.go:177
		// _ = "end of CoverTab[15976]"
//line /usr/local/go/src/net/pipe.go:177
	}
//line /usr/local/go/src/net/pipe.go:177
	// _ = "end of CoverTab[15972]"
//line /usr/local/go/src/net/pipe.go:177
	_go_fuzz_dep_.CoverTab[15973]++
						return n, err
//line /usr/local/go/src/net/pipe.go:178
	// _ = "end of CoverTab[15973]"
}

func (p *pipe) write(b []byte) (n int, err error) {
//line /usr/local/go/src/net/pipe.go:181
	_go_fuzz_dep_.CoverTab[15977]++
						switch {
	case isClosedChan(p.localDone):
//line /usr/local/go/src/net/pipe.go:183
		_go_fuzz_dep_.CoverTab[15980]++
							return 0, io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:184
		// _ = "end of CoverTab[15980]"
	case isClosedChan(p.remoteDone):
//line /usr/local/go/src/net/pipe.go:185
		_go_fuzz_dep_.CoverTab[15981]++
							return 0, io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:186
		// _ = "end of CoverTab[15981]"
	case isClosedChan(p.writeDeadline.wait()):
//line /usr/local/go/src/net/pipe.go:187
		_go_fuzz_dep_.CoverTab[15982]++
							return 0, os.ErrDeadlineExceeded
//line /usr/local/go/src/net/pipe.go:188
		// _ = "end of CoverTab[15982]"
//line /usr/local/go/src/net/pipe.go:188
	default:
//line /usr/local/go/src/net/pipe.go:188
		_go_fuzz_dep_.CoverTab[15983]++
//line /usr/local/go/src/net/pipe.go:188
		// _ = "end of CoverTab[15983]"
	}
//line /usr/local/go/src/net/pipe.go:189
	// _ = "end of CoverTab[15977]"
//line /usr/local/go/src/net/pipe.go:189
	_go_fuzz_dep_.CoverTab[15978]++

						p.wrMu.Lock()
						defer p.wrMu.Unlock()
						for once := true; once || func() bool {
//line /usr/local/go/src/net/pipe.go:193
		_go_fuzz_dep_.CoverTab[15984]++
//line /usr/local/go/src/net/pipe.go:193
		return len(b) > 0
//line /usr/local/go/src/net/pipe.go:193
		// _ = "end of CoverTab[15984]"
//line /usr/local/go/src/net/pipe.go:193
	}(); once = false {
//line /usr/local/go/src/net/pipe.go:193
		_go_fuzz_dep_.CoverTab[15985]++
							select {
		case p.wrTx <- b:
//line /usr/local/go/src/net/pipe.go:195
			_go_fuzz_dep_.CoverTab[15986]++
								nw := <-p.wrRx
								b = b[nw:]
								n += nw
//line /usr/local/go/src/net/pipe.go:198
			// _ = "end of CoverTab[15986]"
		case <-p.localDone:
//line /usr/local/go/src/net/pipe.go:199
			_go_fuzz_dep_.CoverTab[15987]++
								return n, io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:200
			// _ = "end of CoverTab[15987]"
		case <-p.remoteDone:
//line /usr/local/go/src/net/pipe.go:201
			_go_fuzz_dep_.CoverTab[15988]++
								return n, io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:202
			// _ = "end of CoverTab[15988]"
		case <-p.writeDeadline.wait():
//line /usr/local/go/src/net/pipe.go:203
			_go_fuzz_dep_.CoverTab[15989]++
								return n, os.ErrDeadlineExceeded
//line /usr/local/go/src/net/pipe.go:204
			// _ = "end of CoverTab[15989]"
		}
//line /usr/local/go/src/net/pipe.go:205
		// _ = "end of CoverTab[15985]"
	}
//line /usr/local/go/src/net/pipe.go:206
	// _ = "end of CoverTab[15978]"
//line /usr/local/go/src/net/pipe.go:206
	_go_fuzz_dep_.CoverTab[15979]++
						return n, nil
//line /usr/local/go/src/net/pipe.go:207
	// _ = "end of CoverTab[15979]"
}

func (p *pipe) SetDeadline(t time.Time) error {
//line /usr/local/go/src/net/pipe.go:210
	_go_fuzz_dep_.CoverTab[15990]++
						if isClosedChan(p.localDone) || func() bool {
//line /usr/local/go/src/net/pipe.go:211
		_go_fuzz_dep_.CoverTab[15992]++
//line /usr/local/go/src/net/pipe.go:211
		return isClosedChan(p.remoteDone)
//line /usr/local/go/src/net/pipe.go:211
		// _ = "end of CoverTab[15992]"
//line /usr/local/go/src/net/pipe.go:211
	}() {
//line /usr/local/go/src/net/pipe.go:211
		_go_fuzz_dep_.CoverTab[15993]++
							return io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:212
		// _ = "end of CoverTab[15993]"
	} else {
//line /usr/local/go/src/net/pipe.go:213
		_go_fuzz_dep_.CoverTab[15994]++
//line /usr/local/go/src/net/pipe.go:213
		// _ = "end of CoverTab[15994]"
//line /usr/local/go/src/net/pipe.go:213
	}
//line /usr/local/go/src/net/pipe.go:213
	// _ = "end of CoverTab[15990]"
//line /usr/local/go/src/net/pipe.go:213
	_go_fuzz_dep_.CoverTab[15991]++
						p.readDeadline.set(t)
						p.writeDeadline.set(t)
						return nil
//line /usr/local/go/src/net/pipe.go:216
	// _ = "end of CoverTab[15991]"
}

func (p *pipe) SetReadDeadline(t time.Time) error {
//line /usr/local/go/src/net/pipe.go:219
	_go_fuzz_dep_.CoverTab[15995]++
						if isClosedChan(p.localDone) || func() bool {
//line /usr/local/go/src/net/pipe.go:220
		_go_fuzz_dep_.CoverTab[15997]++
//line /usr/local/go/src/net/pipe.go:220
		return isClosedChan(p.remoteDone)
//line /usr/local/go/src/net/pipe.go:220
		// _ = "end of CoverTab[15997]"
//line /usr/local/go/src/net/pipe.go:220
	}() {
//line /usr/local/go/src/net/pipe.go:220
		_go_fuzz_dep_.CoverTab[15998]++
							return io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:221
		// _ = "end of CoverTab[15998]"
	} else {
//line /usr/local/go/src/net/pipe.go:222
		_go_fuzz_dep_.CoverTab[15999]++
//line /usr/local/go/src/net/pipe.go:222
		// _ = "end of CoverTab[15999]"
//line /usr/local/go/src/net/pipe.go:222
	}
//line /usr/local/go/src/net/pipe.go:222
	// _ = "end of CoverTab[15995]"
//line /usr/local/go/src/net/pipe.go:222
	_go_fuzz_dep_.CoverTab[15996]++
						p.readDeadline.set(t)
						return nil
//line /usr/local/go/src/net/pipe.go:224
	// _ = "end of CoverTab[15996]"
}

func (p *pipe) SetWriteDeadline(t time.Time) error {
//line /usr/local/go/src/net/pipe.go:227
	_go_fuzz_dep_.CoverTab[16000]++
						if isClosedChan(p.localDone) || func() bool {
//line /usr/local/go/src/net/pipe.go:228
		_go_fuzz_dep_.CoverTab[16002]++
//line /usr/local/go/src/net/pipe.go:228
		return isClosedChan(p.remoteDone)
//line /usr/local/go/src/net/pipe.go:228
		// _ = "end of CoverTab[16002]"
//line /usr/local/go/src/net/pipe.go:228
	}() {
//line /usr/local/go/src/net/pipe.go:228
		_go_fuzz_dep_.CoverTab[16003]++
							return io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:229
		// _ = "end of CoverTab[16003]"
	} else {
//line /usr/local/go/src/net/pipe.go:230
		_go_fuzz_dep_.CoverTab[16004]++
//line /usr/local/go/src/net/pipe.go:230
		// _ = "end of CoverTab[16004]"
//line /usr/local/go/src/net/pipe.go:230
	}
//line /usr/local/go/src/net/pipe.go:230
	// _ = "end of CoverTab[16000]"
//line /usr/local/go/src/net/pipe.go:230
	_go_fuzz_dep_.CoverTab[16001]++
						p.writeDeadline.set(t)
						return nil
//line /usr/local/go/src/net/pipe.go:232
	// _ = "end of CoverTab[16001]"
}

func (p *pipe) Close() error {
//line /usr/local/go/src/net/pipe.go:235
	_go_fuzz_dep_.CoverTab[16005]++
						p.once.Do(func() { _go_fuzz_dep_.CoverTab[16007]++; close(p.localDone); // _ = "end of CoverTab[16007]" })
//line /usr/local/go/src/net/pipe.go:236
	// _ = "end of CoverTab[16005]"
//line /usr/local/go/src/net/pipe.go:236
	_go_fuzz_dep_.CoverTab[16006]++
						return nil
//line /usr/local/go/src/net/pipe.go:237
	// _ = "end of CoverTab[16006]"
}

//line /usr/local/go/src/net/pipe.go:238
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/pipe.go:238
var _ = _go_fuzz_dep_.CoverTab
