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
	_go_fuzz_dep_.CoverTab[7535]++
						return pipeDeadline{cancel: make(chan struct{})}
//line /usr/local/go/src/net/pipe.go:22
	// _ = "end of CoverTab[7535]"
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
	_go_fuzz_dep_.CoverTab[7536]++
						d.mu.Lock()
						defer d.mu.Unlock()

						if d.timer != nil && func() bool {
//line /usr/local/go/src/net/pipe.go:35
		_go_fuzz_dep_.CoverTab[7540]++
//line /usr/local/go/src/net/pipe.go:35
		return !d.timer.Stop()
//line /usr/local/go/src/net/pipe.go:35
		// _ = "end of CoverTab[7540]"
//line /usr/local/go/src/net/pipe.go:35
	}() {
//line /usr/local/go/src/net/pipe.go:35
		_go_fuzz_dep_.CoverTab[7541]++
							<-d.cancel
//line /usr/local/go/src/net/pipe.go:36
		// _ = "end of CoverTab[7541]"
	} else {
//line /usr/local/go/src/net/pipe.go:37
		_go_fuzz_dep_.CoverTab[7542]++
//line /usr/local/go/src/net/pipe.go:37
		// _ = "end of CoverTab[7542]"
//line /usr/local/go/src/net/pipe.go:37
	}
//line /usr/local/go/src/net/pipe.go:37
	// _ = "end of CoverTab[7536]"
//line /usr/local/go/src/net/pipe.go:37
	_go_fuzz_dep_.CoverTab[7537]++
						d.timer = nil

//line /usr/local/go/src/net/pipe.go:41
	closed := isClosedChan(d.cancel)
	if t.IsZero() {
//line /usr/local/go/src/net/pipe.go:42
		_go_fuzz_dep_.CoverTab[7543]++
							if closed {
//line /usr/local/go/src/net/pipe.go:43
			_go_fuzz_dep_.CoverTab[7545]++
								d.cancel = make(chan struct{})
//line /usr/local/go/src/net/pipe.go:44
			// _ = "end of CoverTab[7545]"
		} else {
//line /usr/local/go/src/net/pipe.go:45
			_go_fuzz_dep_.CoverTab[7546]++
//line /usr/local/go/src/net/pipe.go:45
			// _ = "end of CoverTab[7546]"
//line /usr/local/go/src/net/pipe.go:45
		}
//line /usr/local/go/src/net/pipe.go:45
		// _ = "end of CoverTab[7543]"
//line /usr/local/go/src/net/pipe.go:45
		_go_fuzz_dep_.CoverTab[7544]++
							return
//line /usr/local/go/src/net/pipe.go:46
		// _ = "end of CoverTab[7544]"
	} else {
//line /usr/local/go/src/net/pipe.go:47
		_go_fuzz_dep_.CoverTab[7547]++
//line /usr/local/go/src/net/pipe.go:47
		// _ = "end of CoverTab[7547]"
//line /usr/local/go/src/net/pipe.go:47
	}
//line /usr/local/go/src/net/pipe.go:47
	// _ = "end of CoverTab[7537]"
//line /usr/local/go/src/net/pipe.go:47
	_go_fuzz_dep_.CoverTab[7538]++

//line /usr/local/go/src/net/pipe.go:50
	if dur := time.Until(t); dur > 0 {
//line /usr/local/go/src/net/pipe.go:50
		_go_fuzz_dep_.CoverTab[7548]++
							if closed {
//line /usr/local/go/src/net/pipe.go:51
			_go_fuzz_dep_.CoverTab[7551]++
								d.cancel = make(chan struct{})
//line /usr/local/go/src/net/pipe.go:52
			// _ = "end of CoverTab[7551]"
		} else {
//line /usr/local/go/src/net/pipe.go:53
			_go_fuzz_dep_.CoverTab[7552]++
//line /usr/local/go/src/net/pipe.go:53
			// _ = "end of CoverTab[7552]"
//line /usr/local/go/src/net/pipe.go:53
		}
//line /usr/local/go/src/net/pipe.go:53
		// _ = "end of CoverTab[7548]"
//line /usr/local/go/src/net/pipe.go:53
		_go_fuzz_dep_.CoverTab[7549]++
							d.timer = time.AfterFunc(dur, func() {
//line /usr/local/go/src/net/pipe.go:54
			_go_fuzz_dep_.CoverTab[7553]++
								close(d.cancel)
//line /usr/local/go/src/net/pipe.go:55
			// _ = "end of CoverTab[7553]"
		})
//line /usr/local/go/src/net/pipe.go:56
		// _ = "end of CoverTab[7549]"
//line /usr/local/go/src/net/pipe.go:56
		_go_fuzz_dep_.CoverTab[7550]++
							return
//line /usr/local/go/src/net/pipe.go:57
		// _ = "end of CoverTab[7550]"
	} else {
//line /usr/local/go/src/net/pipe.go:58
		_go_fuzz_dep_.CoverTab[7554]++
//line /usr/local/go/src/net/pipe.go:58
		// _ = "end of CoverTab[7554]"
//line /usr/local/go/src/net/pipe.go:58
	}
//line /usr/local/go/src/net/pipe.go:58
	// _ = "end of CoverTab[7538]"
//line /usr/local/go/src/net/pipe.go:58
	_go_fuzz_dep_.CoverTab[7539]++

//line /usr/local/go/src/net/pipe.go:61
	if !closed {
//line /usr/local/go/src/net/pipe.go:61
		_go_fuzz_dep_.CoverTab[7555]++
							close(d.cancel)
//line /usr/local/go/src/net/pipe.go:62
		// _ = "end of CoverTab[7555]"
	} else {
//line /usr/local/go/src/net/pipe.go:63
		_go_fuzz_dep_.CoverTab[7556]++
//line /usr/local/go/src/net/pipe.go:63
		// _ = "end of CoverTab[7556]"
//line /usr/local/go/src/net/pipe.go:63
	}
//line /usr/local/go/src/net/pipe.go:63
	// _ = "end of CoverTab[7539]"
}

// wait returns a channel that is closed when the deadline is exceeded.
func (d *pipeDeadline) wait() chan struct{} {
//line /usr/local/go/src/net/pipe.go:67
	_go_fuzz_dep_.CoverTab[7557]++
						d.mu.Lock()
						defer d.mu.Unlock()
						return d.cancel
//line /usr/local/go/src/net/pipe.go:70
	// _ = "end of CoverTab[7557]"
}

func isClosedChan(c <-chan struct{}) bool {
//line /usr/local/go/src/net/pipe.go:73
	_go_fuzz_dep_.CoverTab[7558]++
						select {
	case <-c:
//line /usr/local/go/src/net/pipe.go:75
		_go_fuzz_dep_.CoverTab[7559]++
							return true
//line /usr/local/go/src/net/pipe.go:76
		// _ = "end of CoverTab[7559]"
	default:
//line /usr/local/go/src/net/pipe.go:77
		_go_fuzz_dep_.CoverTab[7560]++
							return false
//line /usr/local/go/src/net/pipe.go:78
		// _ = "end of CoverTab[7560]"
	}
//line /usr/local/go/src/net/pipe.go:79
	// _ = "end of CoverTab[7558]"
}

type pipeAddr struct{}

func (pipeAddr) Network() string {
//line /usr/local/go/src/net/pipe.go:84
	_go_fuzz_dep_.CoverTab[7561]++
//line /usr/local/go/src/net/pipe.go:84
	return "pipe"
//line /usr/local/go/src/net/pipe.go:84
	// _ = "end of CoverTab[7561]"
//line /usr/local/go/src/net/pipe.go:84
}
func (pipeAddr) String() string {
//line /usr/local/go/src/net/pipe.go:85
	_go_fuzz_dep_.CoverTab[7562]++
//line /usr/local/go/src/net/pipe.go:85
	return "pipe"
//line /usr/local/go/src/net/pipe.go:85
	// _ = "end of CoverTab[7562]"
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
	_go_fuzz_dep_.CoverTab[7563]++
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
	// _ = "end of CoverTab[7563]"
}

func (*pipe) LocalAddr() Addr {
//line /usr/local/go/src/net/pipe.go:138
	_go_fuzz_dep_.CoverTab[7564]++
//line /usr/local/go/src/net/pipe.go:138
	return pipeAddr{}
//line /usr/local/go/src/net/pipe.go:138
	// _ = "end of CoverTab[7564]"
//line /usr/local/go/src/net/pipe.go:138
}
func (*pipe) RemoteAddr() Addr {
//line /usr/local/go/src/net/pipe.go:139
	_go_fuzz_dep_.CoverTab[7565]++
//line /usr/local/go/src/net/pipe.go:139
	return pipeAddr{}
//line /usr/local/go/src/net/pipe.go:139
	// _ = "end of CoverTab[7565]"
//line /usr/local/go/src/net/pipe.go:139
}

func (p *pipe) Read(b []byte) (int, error) {
//line /usr/local/go/src/net/pipe.go:141
	_go_fuzz_dep_.CoverTab[7566]++
						n, err := p.read(b)
						if err != nil && func() bool {
//line /usr/local/go/src/net/pipe.go:143
		_go_fuzz_dep_.CoverTab[7568]++
//line /usr/local/go/src/net/pipe.go:143
		return err != io.EOF
//line /usr/local/go/src/net/pipe.go:143
		// _ = "end of CoverTab[7568]"
//line /usr/local/go/src/net/pipe.go:143
	}() && func() bool {
//line /usr/local/go/src/net/pipe.go:143
		_go_fuzz_dep_.CoverTab[7569]++
//line /usr/local/go/src/net/pipe.go:143
		return err != io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:143
		// _ = "end of CoverTab[7569]"
//line /usr/local/go/src/net/pipe.go:143
	}() {
//line /usr/local/go/src/net/pipe.go:143
		_go_fuzz_dep_.CoverTab[7570]++
							err = &OpError{Op: "read", Net: "pipe", Err: err}
//line /usr/local/go/src/net/pipe.go:144
		// _ = "end of CoverTab[7570]"
	} else {
//line /usr/local/go/src/net/pipe.go:145
		_go_fuzz_dep_.CoverTab[7571]++
//line /usr/local/go/src/net/pipe.go:145
		// _ = "end of CoverTab[7571]"
//line /usr/local/go/src/net/pipe.go:145
	}
//line /usr/local/go/src/net/pipe.go:145
	// _ = "end of CoverTab[7566]"
//line /usr/local/go/src/net/pipe.go:145
	_go_fuzz_dep_.CoverTab[7567]++
						return n, err
//line /usr/local/go/src/net/pipe.go:146
	// _ = "end of CoverTab[7567]"
}

func (p *pipe) read(b []byte) (n int, err error) {
//line /usr/local/go/src/net/pipe.go:149
	_go_fuzz_dep_.CoverTab[7572]++
						switch {
	case isClosedChan(p.localDone):
//line /usr/local/go/src/net/pipe.go:151
		_go_fuzz_dep_.CoverTab[7574]++
							return 0, io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:152
		// _ = "end of CoverTab[7574]"
	case isClosedChan(p.remoteDone):
//line /usr/local/go/src/net/pipe.go:153
		_go_fuzz_dep_.CoverTab[7575]++
							return 0, io.EOF
//line /usr/local/go/src/net/pipe.go:154
		// _ = "end of CoverTab[7575]"
	case isClosedChan(p.readDeadline.wait()):
//line /usr/local/go/src/net/pipe.go:155
		_go_fuzz_dep_.CoverTab[7576]++
							return 0, os.ErrDeadlineExceeded
//line /usr/local/go/src/net/pipe.go:156
		// _ = "end of CoverTab[7576]"
//line /usr/local/go/src/net/pipe.go:156
	default:
//line /usr/local/go/src/net/pipe.go:156
		_go_fuzz_dep_.CoverTab[7577]++
//line /usr/local/go/src/net/pipe.go:156
		// _ = "end of CoverTab[7577]"
	}
//line /usr/local/go/src/net/pipe.go:157
	// _ = "end of CoverTab[7572]"
//line /usr/local/go/src/net/pipe.go:157
	_go_fuzz_dep_.CoverTab[7573]++

						select {
	case bw := <-p.rdRx:
//line /usr/local/go/src/net/pipe.go:160
		_go_fuzz_dep_.CoverTab[7578]++
							nr := copy(b, bw)
							p.rdTx <- nr
							return nr, nil
//line /usr/local/go/src/net/pipe.go:163
		// _ = "end of CoverTab[7578]"
	case <-p.localDone:
//line /usr/local/go/src/net/pipe.go:164
		_go_fuzz_dep_.CoverTab[7579]++
							return 0, io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:165
		// _ = "end of CoverTab[7579]"
	case <-p.remoteDone:
//line /usr/local/go/src/net/pipe.go:166
		_go_fuzz_dep_.CoverTab[7580]++
							return 0, io.EOF
//line /usr/local/go/src/net/pipe.go:167
		// _ = "end of CoverTab[7580]"
	case <-p.readDeadline.wait():
//line /usr/local/go/src/net/pipe.go:168
		_go_fuzz_dep_.CoverTab[7581]++
							return 0, os.ErrDeadlineExceeded
//line /usr/local/go/src/net/pipe.go:169
		// _ = "end of CoverTab[7581]"
	}
//line /usr/local/go/src/net/pipe.go:170
	// _ = "end of CoverTab[7573]"
}

func (p *pipe) Write(b []byte) (int, error) {
//line /usr/local/go/src/net/pipe.go:173
	_go_fuzz_dep_.CoverTab[7582]++
						n, err := p.write(b)
						if err != nil && func() bool {
//line /usr/local/go/src/net/pipe.go:175
		_go_fuzz_dep_.CoverTab[7584]++
//line /usr/local/go/src/net/pipe.go:175
		return err != io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:175
		// _ = "end of CoverTab[7584]"
//line /usr/local/go/src/net/pipe.go:175
	}() {
//line /usr/local/go/src/net/pipe.go:175
		_go_fuzz_dep_.CoverTab[7585]++
							err = &OpError{Op: "write", Net: "pipe", Err: err}
//line /usr/local/go/src/net/pipe.go:176
		// _ = "end of CoverTab[7585]"
	} else {
//line /usr/local/go/src/net/pipe.go:177
		_go_fuzz_dep_.CoverTab[7586]++
//line /usr/local/go/src/net/pipe.go:177
		// _ = "end of CoverTab[7586]"
//line /usr/local/go/src/net/pipe.go:177
	}
//line /usr/local/go/src/net/pipe.go:177
	// _ = "end of CoverTab[7582]"
//line /usr/local/go/src/net/pipe.go:177
	_go_fuzz_dep_.CoverTab[7583]++
						return n, err
//line /usr/local/go/src/net/pipe.go:178
	// _ = "end of CoverTab[7583]"
}

func (p *pipe) write(b []byte) (n int, err error) {
//line /usr/local/go/src/net/pipe.go:181
	_go_fuzz_dep_.CoverTab[7587]++
						switch {
	case isClosedChan(p.localDone):
//line /usr/local/go/src/net/pipe.go:183
		_go_fuzz_dep_.CoverTab[7590]++
							return 0, io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:184
		// _ = "end of CoverTab[7590]"
	case isClosedChan(p.remoteDone):
//line /usr/local/go/src/net/pipe.go:185
		_go_fuzz_dep_.CoverTab[7591]++
							return 0, io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:186
		// _ = "end of CoverTab[7591]"
	case isClosedChan(p.writeDeadline.wait()):
//line /usr/local/go/src/net/pipe.go:187
		_go_fuzz_dep_.CoverTab[7592]++
							return 0, os.ErrDeadlineExceeded
//line /usr/local/go/src/net/pipe.go:188
		// _ = "end of CoverTab[7592]"
//line /usr/local/go/src/net/pipe.go:188
	default:
//line /usr/local/go/src/net/pipe.go:188
		_go_fuzz_dep_.CoverTab[7593]++
//line /usr/local/go/src/net/pipe.go:188
		// _ = "end of CoverTab[7593]"
	}
//line /usr/local/go/src/net/pipe.go:189
	// _ = "end of CoverTab[7587]"
//line /usr/local/go/src/net/pipe.go:189
	_go_fuzz_dep_.CoverTab[7588]++

						p.wrMu.Lock()
						defer p.wrMu.Unlock()
						for once := true; once || func() bool {
//line /usr/local/go/src/net/pipe.go:193
		_go_fuzz_dep_.CoverTab[7594]++
//line /usr/local/go/src/net/pipe.go:193
		return len(b) > 0
//line /usr/local/go/src/net/pipe.go:193
		// _ = "end of CoverTab[7594]"
//line /usr/local/go/src/net/pipe.go:193
	}(); once = false {
//line /usr/local/go/src/net/pipe.go:193
		_go_fuzz_dep_.CoverTab[7595]++
							select {
		case p.wrTx <- b:
//line /usr/local/go/src/net/pipe.go:195
			_go_fuzz_dep_.CoverTab[7596]++
								nw := <-p.wrRx
								b = b[nw:]
								n += nw
//line /usr/local/go/src/net/pipe.go:198
			// _ = "end of CoverTab[7596]"
		case <-p.localDone:
//line /usr/local/go/src/net/pipe.go:199
			_go_fuzz_dep_.CoverTab[7597]++
								return n, io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:200
			// _ = "end of CoverTab[7597]"
		case <-p.remoteDone:
//line /usr/local/go/src/net/pipe.go:201
			_go_fuzz_dep_.CoverTab[7598]++
								return n, io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:202
			// _ = "end of CoverTab[7598]"
		case <-p.writeDeadline.wait():
//line /usr/local/go/src/net/pipe.go:203
			_go_fuzz_dep_.CoverTab[7599]++
								return n, os.ErrDeadlineExceeded
//line /usr/local/go/src/net/pipe.go:204
			// _ = "end of CoverTab[7599]"
		}
//line /usr/local/go/src/net/pipe.go:205
		// _ = "end of CoverTab[7595]"
	}
//line /usr/local/go/src/net/pipe.go:206
	// _ = "end of CoverTab[7588]"
//line /usr/local/go/src/net/pipe.go:206
	_go_fuzz_dep_.CoverTab[7589]++
						return n, nil
//line /usr/local/go/src/net/pipe.go:207
	// _ = "end of CoverTab[7589]"
}

func (p *pipe) SetDeadline(t time.Time) error {
//line /usr/local/go/src/net/pipe.go:210
	_go_fuzz_dep_.CoverTab[7600]++
						if isClosedChan(p.localDone) || func() bool {
//line /usr/local/go/src/net/pipe.go:211
		_go_fuzz_dep_.CoverTab[7602]++
//line /usr/local/go/src/net/pipe.go:211
		return isClosedChan(p.remoteDone)
//line /usr/local/go/src/net/pipe.go:211
		// _ = "end of CoverTab[7602]"
//line /usr/local/go/src/net/pipe.go:211
	}() {
//line /usr/local/go/src/net/pipe.go:211
		_go_fuzz_dep_.CoverTab[7603]++
							return io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:212
		// _ = "end of CoverTab[7603]"
	} else {
//line /usr/local/go/src/net/pipe.go:213
		_go_fuzz_dep_.CoverTab[7604]++
//line /usr/local/go/src/net/pipe.go:213
		// _ = "end of CoverTab[7604]"
//line /usr/local/go/src/net/pipe.go:213
	}
//line /usr/local/go/src/net/pipe.go:213
	// _ = "end of CoverTab[7600]"
//line /usr/local/go/src/net/pipe.go:213
	_go_fuzz_dep_.CoverTab[7601]++
						p.readDeadline.set(t)
						p.writeDeadline.set(t)
						return nil
//line /usr/local/go/src/net/pipe.go:216
	// _ = "end of CoverTab[7601]"
}

func (p *pipe) SetReadDeadline(t time.Time) error {
//line /usr/local/go/src/net/pipe.go:219
	_go_fuzz_dep_.CoverTab[7605]++
						if isClosedChan(p.localDone) || func() bool {
//line /usr/local/go/src/net/pipe.go:220
		_go_fuzz_dep_.CoverTab[7607]++
//line /usr/local/go/src/net/pipe.go:220
		return isClosedChan(p.remoteDone)
//line /usr/local/go/src/net/pipe.go:220
		// _ = "end of CoverTab[7607]"
//line /usr/local/go/src/net/pipe.go:220
	}() {
//line /usr/local/go/src/net/pipe.go:220
		_go_fuzz_dep_.CoverTab[7608]++
							return io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:221
		// _ = "end of CoverTab[7608]"
	} else {
//line /usr/local/go/src/net/pipe.go:222
		_go_fuzz_dep_.CoverTab[7609]++
//line /usr/local/go/src/net/pipe.go:222
		// _ = "end of CoverTab[7609]"
//line /usr/local/go/src/net/pipe.go:222
	}
//line /usr/local/go/src/net/pipe.go:222
	// _ = "end of CoverTab[7605]"
//line /usr/local/go/src/net/pipe.go:222
	_go_fuzz_dep_.CoverTab[7606]++
						p.readDeadline.set(t)
						return nil
//line /usr/local/go/src/net/pipe.go:224
	// _ = "end of CoverTab[7606]"
}

func (p *pipe) SetWriteDeadline(t time.Time) error {
//line /usr/local/go/src/net/pipe.go:227
	_go_fuzz_dep_.CoverTab[7610]++
						if isClosedChan(p.localDone) || func() bool {
//line /usr/local/go/src/net/pipe.go:228
		_go_fuzz_dep_.CoverTab[7612]++
//line /usr/local/go/src/net/pipe.go:228
		return isClosedChan(p.remoteDone)
//line /usr/local/go/src/net/pipe.go:228
		// _ = "end of CoverTab[7612]"
//line /usr/local/go/src/net/pipe.go:228
	}() {
//line /usr/local/go/src/net/pipe.go:228
		_go_fuzz_dep_.CoverTab[7613]++
							return io.ErrClosedPipe
//line /usr/local/go/src/net/pipe.go:229
		// _ = "end of CoverTab[7613]"
	} else {
//line /usr/local/go/src/net/pipe.go:230
		_go_fuzz_dep_.CoverTab[7614]++
//line /usr/local/go/src/net/pipe.go:230
		// _ = "end of CoverTab[7614]"
//line /usr/local/go/src/net/pipe.go:230
	}
//line /usr/local/go/src/net/pipe.go:230
	// _ = "end of CoverTab[7610]"
//line /usr/local/go/src/net/pipe.go:230
	_go_fuzz_dep_.CoverTab[7611]++
						p.writeDeadline.set(t)
						return nil
//line /usr/local/go/src/net/pipe.go:232
	// _ = "end of CoverTab[7611]"
}

func (p *pipe) Close() error {
//line /usr/local/go/src/net/pipe.go:235
	_go_fuzz_dep_.CoverTab[7615]++
						p.once.Do(func() { _go_fuzz_dep_.CoverTab[7617]++; close(p.localDone); // _ = "end of CoverTab[7617]" })
//line /usr/local/go/src/net/pipe.go:236
	// _ = "end of CoverTab[7615]"
//line /usr/local/go/src/net/pipe.go:236
	_go_fuzz_dep_.CoverTab[7616]++
						return nil
//line /usr/local/go/src/net/pipe.go:237
	// _ = "end of CoverTab[7616]"
}

//line /usr/local/go/src/net/pipe.go:238
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/pipe.go:238
var _ = _go_fuzz_dep_.CoverTab
