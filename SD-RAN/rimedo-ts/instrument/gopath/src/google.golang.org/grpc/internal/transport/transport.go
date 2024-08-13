//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:19
// Package transport defines and implements message oriented communication
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:19
// channel to complete various transactions (e.g., an RPC).  It is meant for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:19
// grpc-internal usage and is not intended to be imported directly by users.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:22
package transport

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:22
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:22
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:22
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:22
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:22
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:22
)

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/tap"
)

// ErrNoHeaders is used as a signal that a trailers only response was received,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:46
// and is not a real error.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:48
var ErrNoHeaders = errors.New("stream has no headers")

const logLevel = 2

type bufferPool struct {
	pool sync.Pool
}

func newBufferPool() *bufferPool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:56
	_go_fuzz_dep_.CoverTab[78598]++
													return &bufferPool{
		pool: sync.Pool{
			New: func() interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:59
				_go_fuzz_dep_.CoverTab[78599]++
																return new(bytes.Buffer)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:60
				// _ = "end of CoverTab[78599]"
			},
		},
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:63
	// _ = "end of CoverTab[78598]"
}

func (p *bufferPool) get() *bytes.Buffer {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:66
	_go_fuzz_dep_.CoverTab[78600]++
													return p.pool.Get().(*bytes.Buffer)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:67
	// _ = "end of CoverTab[78600]"
}

func (p *bufferPool) put(b *bytes.Buffer) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:70
	_go_fuzz_dep_.CoverTab[78601]++
													p.pool.Put(b)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:71
	// _ = "end of CoverTab[78601]"
}

// recvMsg represents the received msg from the transport. All transport
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:74
// protocol specific info has been removed.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:76
type recvMsg struct {
	buffer	*bytes.Buffer
	// nil: received some data
	// io.EOF: stream is completed. data is nil.
	// other non-nil error: transport failure. data is nil.
	err	error
}

// recvBuffer is an unbounded channel of recvMsg structs.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:84
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:84
// Note: recvBuffer differs from buffer.Unbounded only in the fact that it
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:84
// holds a channel of recvMsg structs instead of objects implementing "item"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:84
// interface. recvBuffer is written to much more often and using strict recvMsg
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:84
// structs helps avoid allocation in "recvBuffer.put"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:90
type recvBuffer struct {
	c	chan recvMsg
	mu	sync.Mutex
	backlog	[]recvMsg
	err	error
}

func newRecvBuffer() *recvBuffer {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:97
	_go_fuzz_dep_.CoverTab[78602]++
													b := &recvBuffer{
		c: make(chan recvMsg, 1),
	}
													return b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:101
	// _ = "end of CoverTab[78602]"
}

func (b *recvBuffer) put(r recvMsg) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:104
	_go_fuzz_dep_.CoverTab[78603]++
													b.mu.Lock()
													if b.err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:106
		_go_fuzz_dep_.CoverTab[78606]++
														b.mu.Unlock()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:110
		return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:110
		// _ = "end of CoverTab[78606]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:111
		_go_fuzz_dep_.CoverTab[78607]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:111
		// _ = "end of CoverTab[78607]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:111
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:111
	// _ = "end of CoverTab[78603]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:111
	_go_fuzz_dep_.CoverTab[78604]++
													b.err = r.err
													if len(b.backlog) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:113
		_go_fuzz_dep_.CoverTab[78608]++
														select {
		case b.c <- r:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:115
			_go_fuzz_dep_.CoverTab[78609]++
															b.mu.Unlock()
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:117
			// _ = "end of CoverTab[78609]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:118
			_go_fuzz_dep_.CoverTab[78610]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:118
			// _ = "end of CoverTab[78610]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:119
		// _ = "end of CoverTab[78608]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:120
		_go_fuzz_dep_.CoverTab[78611]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:120
		// _ = "end of CoverTab[78611]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:120
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:120
	// _ = "end of CoverTab[78604]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:120
	_go_fuzz_dep_.CoverTab[78605]++
													b.backlog = append(b.backlog, r)
													b.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:122
	// _ = "end of CoverTab[78605]"
}

func (b *recvBuffer) load() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:125
	_go_fuzz_dep_.CoverTab[78612]++
													b.mu.Lock()
													if len(b.backlog) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:127
		_go_fuzz_dep_.CoverTab[78614]++
														select {
		case b.c <- b.backlog[0]:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:129
			_go_fuzz_dep_.CoverTab[78615]++
															b.backlog[0] = recvMsg{}
															b.backlog = b.backlog[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:131
			// _ = "end of CoverTab[78615]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:132
			_go_fuzz_dep_.CoverTab[78616]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:132
			// _ = "end of CoverTab[78616]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:133
		// _ = "end of CoverTab[78614]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:134
		_go_fuzz_dep_.CoverTab[78617]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:134
		// _ = "end of CoverTab[78617]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:134
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:134
	// _ = "end of CoverTab[78612]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:134
	_go_fuzz_dep_.CoverTab[78613]++
													b.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:135
	// _ = "end of CoverTab[78613]"
}

// get returns the channel that receives a recvMsg in the buffer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:138
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:138
// Upon receipt of a recvMsg, the caller should call load to send another
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:138
// recvMsg onto the channel if there is any.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:142
func (b *recvBuffer) get() <-chan recvMsg {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:142
	_go_fuzz_dep_.CoverTab[78618]++
													return b.c
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:143
	// _ = "end of CoverTab[78618]"
}

// recvBufferReader implements io.Reader interface to read the data from
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:146
// recvBuffer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:148
type recvBufferReader struct {
	closeStream	func(error)	// Closes the client transport stream with the given error and nil trailer metadata.
	ctx		context.Context
	ctxDone		<-chan struct{}	// cache of ctx.Done() (for performance).
	recv		*recvBuffer
	last		*bytes.Buffer	// Stores the remaining data in the previous calls.
	err		error
	freeBuffer	func(*bytes.Buffer)
}

// Read reads the next len(p) bytes from last. If last is drained, it tries to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:158
// read additional data from recv. It blocks if there no additional data available
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:158
// in recv. If Read returns any non-nil error, it will continue to return that error.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:161
func (r *recvBufferReader) Read(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:161
	_go_fuzz_dep_.CoverTab[78619]++
													if r.err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:162
		_go_fuzz_dep_.CoverTab[78623]++
														return 0, r.err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:163
		// _ = "end of CoverTab[78623]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:164
		_go_fuzz_dep_.CoverTab[78624]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:164
		// _ = "end of CoverTab[78624]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:164
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:164
	// _ = "end of CoverTab[78619]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:164
	_go_fuzz_dep_.CoverTab[78620]++
													if r.last != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:165
		_go_fuzz_dep_.CoverTab[78625]++

														copied, _ := r.last.Read(p)
														if r.last.Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:168
			_go_fuzz_dep_.CoverTab[78627]++
															r.freeBuffer(r.last)
															r.last = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:170
			// _ = "end of CoverTab[78627]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:171
			_go_fuzz_dep_.CoverTab[78628]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:171
			// _ = "end of CoverTab[78628]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:171
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:171
		// _ = "end of CoverTab[78625]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:171
		_go_fuzz_dep_.CoverTab[78626]++
														return copied, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:172
		// _ = "end of CoverTab[78626]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:173
		_go_fuzz_dep_.CoverTab[78629]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:173
		// _ = "end of CoverTab[78629]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:173
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:173
	// _ = "end of CoverTab[78620]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:173
	_go_fuzz_dep_.CoverTab[78621]++
													if r.closeStream != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:174
		_go_fuzz_dep_.CoverTab[78630]++
														n, r.err = r.readClient(p)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:175
		// _ = "end of CoverTab[78630]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:176
		_go_fuzz_dep_.CoverTab[78631]++
														n, r.err = r.read(p)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:177
		// _ = "end of CoverTab[78631]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:178
	// _ = "end of CoverTab[78621]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:178
	_go_fuzz_dep_.CoverTab[78622]++
													return n, r.err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:179
	// _ = "end of CoverTab[78622]"
}

func (r *recvBufferReader) read(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:182
	_go_fuzz_dep_.CoverTab[78632]++
													select {
	case <-r.ctxDone:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:184
		_go_fuzz_dep_.CoverTab[78633]++
														return 0, ContextErr(r.ctx.Err())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:185
		// _ = "end of CoverTab[78633]"
	case m := <-r.recv.get():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:186
		_go_fuzz_dep_.CoverTab[78634]++
														return r.readAdditional(m, p)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:187
		// _ = "end of CoverTab[78634]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:188
	// _ = "end of CoverTab[78632]"
}

func (r *recvBufferReader) readClient(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:191
	_go_fuzz_dep_.CoverTab[78635]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:195
	select {
	case <-r.ctxDone:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:196
		_go_fuzz_dep_.CoverTab[78636]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:210
		r.closeStream(ContextErr(r.ctx.Err()))
														m := <-r.recv.get()
														return r.readAdditional(m, p)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:212
		// _ = "end of CoverTab[78636]"
	case m := <-r.recv.get():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:213
		_go_fuzz_dep_.CoverTab[78637]++
														return r.readAdditional(m, p)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:214
		// _ = "end of CoverTab[78637]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:215
	// _ = "end of CoverTab[78635]"
}

func (r *recvBufferReader) readAdditional(m recvMsg, p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:218
	_go_fuzz_dep_.CoverTab[78638]++
													r.recv.load()
													if m.err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:220
		_go_fuzz_dep_.CoverTab[78641]++
														return 0, m.err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:221
		// _ = "end of CoverTab[78641]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:222
		_go_fuzz_dep_.CoverTab[78642]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:222
		// _ = "end of CoverTab[78642]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:222
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:222
	// _ = "end of CoverTab[78638]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:222
	_go_fuzz_dep_.CoverTab[78639]++
													copied, _ := m.buffer.Read(p)
													if m.buffer.Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:224
		_go_fuzz_dep_.CoverTab[78643]++
														r.freeBuffer(m.buffer)
														r.last = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:226
		// _ = "end of CoverTab[78643]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:227
		_go_fuzz_dep_.CoverTab[78644]++
														r.last = m.buffer
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:228
		// _ = "end of CoverTab[78644]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:229
	// _ = "end of CoverTab[78639]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:229
	_go_fuzz_dep_.CoverTab[78640]++
													return copied, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:230
	// _ = "end of CoverTab[78640]"
}

type streamState uint32

const (
	streamActive	streamState	= iota
	streamWriteDone			// EndStream sent
	streamReadDone			// EndStream received
	streamDone			// the entire stream is finished.
)

// Stream represents an RPC in the transport layer.
type Stream struct {
	id		uint32
	st		ServerTransport		// nil for client side Stream
	ct		*http2Client		// nil for server side Stream
	ctx		context.Context		// the associated context of the stream
	cancel		context.CancelFunc	// always nil for client side Stream
	done		chan struct{}		// closed at the end of stream to unblock writers. On the client side.
	doneFunc	func()			// invoked at the end of stream on client side.
	ctxDone		<-chan struct{}		// same as done chan but for server side. Cache of ctx.Done() (for performance)
	method		string			// the associated RPC method of the stream
	recvCompress	string
	sendCompress	string
	buf		*recvBuffer
	trReader	io.Reader
	fc		*inFlow
	wq		*writeQuota

	// Holds compressor names passed in grpc-accept-encoding metadata from the
	// client. This is empty for the client side stream.
	clientAdvertisedCompressors	string
	// Callback to state application's intentions to read data. This
	// is used to adjust flow control, if needed.
	requestRead	func(int)

	headerChan		chan struct{}	// closed to indicate the end of header metadata.
	headerChanClosed	uint32		// set when headerChan is closed. Used to avoid closing headerChan multiple times.
	// headerValid indicates whether a valid header was received.  Only
	// meaningful after headerChan is closed (always call waitOnHeader() before
	// reading its value).  Not valid on server side.
	headerValid	bool

	// hdrMu protects header and trailer metadata on the server-side.
	hdrMu	sync.Mutex
	// On client side, header keeps the received header metadata.
	//
	// On server side, header keeps the header set by SetHeader(). The complete
	// header will merged into this after t.WriteHeader() is called.
	header	metadata.MD
	trailer	metadata.MD	// the key-value map of trailer metadata.

	noHeaders	bool	// set if the client never received headers (set only after the stream is done).

	// On the server-side, headerSent is atomically set to 1 when the headers are sent out.
	headerSent	uint32

	state	streamState

	// On client-side it is the status error received from the server.
	// On server-side it is unused.
	status	*status.Status

	bytesReceived	uint32	// indicates whether any bytes have been received on this stream
	unprocessed	uint32	// set if the server sends a refused stream or GOAWAY including this stream

	// contentSubtype is the content-subtype for requests.
	// this must be lowercase or the behavior is undefined.
	contentSubtype	string
}

// isHeaderSent is only valid on the server-side.
func (s *Stream) isHeaderSent() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:303
	_go_fuzz_dep_.CoverTab[78645]++
													return atomic.LoadUint32(&s.headerSent) == 1
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:304
	// _ = "end of CoverTab[78645]"
}

// updateHeaderSent updates headerSent and returns true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:307
// if it was alreay set. It is valid only on server-side.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:309
func (s *Stream) updateHeaderSent() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:309
	_go_fuzz_dep_.CoverTab[78646]++
													return atomic.SwapUint32(&s.headerSent, 1) == 1
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:310
	// _ = "end of CoverTab[78646]"
}

func (s *Stream) swapState(st streamState) streamState {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:313
	_go_fuzz_dep_.CoverTab[78647]++
													return streamState(atomic.SwapUint32((*uint32)(&s.state), uint32(st)))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:314
	// _ = "end of CoverTab[78647]"
}

func (s *Stream) compareAndSwapState(oldState, newState streamState) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:317
	_go_fuzz_dep_.CoverTab[78648]++
													return atomic.CompareAndSwapUint32((*uint32)(&s.state), uint32(oldState), uint32(newState))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:318
	// _ = "end of CoverTab[78648]"
}

func (s *Stream) getState() streamState {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:321
	_go_fuzz_dep_.CoverTab[78649]++
													return streamState(atomic.LoadUint32((*uint32)(&s.state)))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:322
	// _ = "end of CoverTab[78649]"
}

func (s *Stream) waitOnHeader() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:325
	_go_fuzz_dep_.CoverTab[78650]++
													if s.headerChan == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:326
		_go_fuzz_dep_.CoverTab[78652]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:329
		return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:329
		// _ = "end of CoverTab[78652]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:330
		_go_fuzz_dep_.CoverTab[78653]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:330
		// _ = "end of CoverTab[78653]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:330
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:330
	// _ = "end of CoverTab[78650]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:330
	_go_fuzz_dep_.CoverTab[78651]++
													select {
	case <-s.ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:332
		_go_fuzz_dep_.CoverTab[78654]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:335
		s.ct.CloseStream(s, ContextErr(s.ctx.Err()))

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:338
		<-s.headerChan
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:338
		// _ = "end of CoverTab[78654]"
	case <-s.headerChan:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:339
		_go_fuzz_dep_.CoverTab[78655]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:339
		// _ = "end of CoverTab[78655]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:340
	// _ = "end of CoverTab[78651]"
}

// RecvCompress returns the compression algorithm applied to the inbound
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:343
// message. It is empty string if there is no compression applied.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:345
func (s *Stream) RecvCompress() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:345
	_go_fuzz_dep_.CoverTab[78656]++
													s.waitOnHeader()
													return s.recvCompress
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:347
	// _ = "end of CoverTab[78656]"
}

// SetSendCompress sets the compression algorithm to the stream.
func (s *Stream) SetSendCompress(name string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:351
	_go_fuzz_dep_.CoverTab[78657]++
													if s.isHeaderSent() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:352
		_go_fuzz_dep_.CoverTab[78659]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:352
		return s.getState() == streamDone
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:352
		// _ = "end of CoverTab[78659]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:352
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:352
		_go_fuzz_dep_.CoverTab[78660]++
														return errors.New("transport: set send compressor called after headers sent or stream done")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:353
		// _ = "end of CoverTab[78660]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:354
		_go_fuzz_dep_.CoverTab[78661]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:354
		// _ = "end of CoverTab[78661]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:354
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:354
	// _ = "end of CoverTab[78657]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:354
	_go_fuzz_dep_.CoverTab[78658]++

													s.sendCompress = name
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:357
	// _ = "end of CoverTab[78658]"
}

// SendCompress returns the send compressor name.
func (s *Stream) SendCompress() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:361
	_go_fuzz_dep_.CoverTab[78662]++
													return s.sendCompress
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:362
	// _ = "end of CoverTab[78662]"
}

// ClientAdvertisedCompressors returns the compressor names advertised by the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:365
// client via grpc-accept-encoding header.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:367
func (s *Stream) ClientAdvertisedCompressors() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:367
	_go_fuzz_dep_.CoverTab[78663]++
													return s.clientAdvertisedCompressors
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:368
	// _ = "end of CoverTab[78663]"
}

// Done returns a channel which is closed when it receives the final status
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:371
// from the server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:373
func (s *Stream) Done() <-chan struct{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:373
	_go_fuzz_dep_.CoverTab[78664]++
													return s.done
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:374
	// _ = "end of CoverTab[78664]"
}

// Header returns the header metadata of the stream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:377
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:377
// On client side, it acquires the key-value pairs of header metadata once it is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:377
// available. It blocks until i) the metadata is ready or ii) there is no header
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:377
// metadata or iii) the stream is canceled/expired.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:377
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:377
// On server side, it returns the out header after t.WriteHeader is called.  It
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:377
// does not block and must not be called until after WriteHeader.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:385
func (s *Stream) Header() (metadata.MD, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:385
	_go_fuzz_dep_.CoverTab[78665]++
													if s.headerChan == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:386
		_go_fuzz_dep_.CoverTab[78669]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:389
		return s.header.Copy(), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:389
		// _ = "end of CoverTab[78669]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:390
		_go_fuzz_dep_.CoverTab[78670]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:390
		// _ = "end of CoverTab[78670]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:390
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:390
	// _ = "end of CoverTab[78665]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:390
	_go_fuzz_dep_.CoverTab[78666]++
													s.waitOnHeader()

													if !s.headerValid {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:393
		_go_fuzz_dep_.CoverTab[78671]++
														return nil, s.status.Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:394
		// _ = "end of CoverTab[78671]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:395
		_go_fuzz_dep_.CoverTab[78672]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:395
		// _ = "end of CoverTab[78672]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:395
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:395
	// _ = "end of CoverTab[78666]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:395
	_go_fuzz_dep_.CoverTab[78667]++

													if s.noHeaders {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:397
		_go_fuzz_dep_.CoverTab[78673]++
														return nil, ErrNoHeaders
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:398
		// _ = "end of CoverTab[78673]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:399
		_go_fuzz_dep_.CoverTab[78674]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:399
		// _ = "end of CoverTab[78674]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:399
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:399
	// _ = "end of CoverTab[78667]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:399
	_go_fuzz_dep_.CoverTab[78668]++

													return s.header.Copy(), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:401
	// _ = "end of CoverTab[78668]"
}

// TrailersOnly blocks until a header or trailers-only frame is received and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:404
// then returns true if the stream was trailers-only.  If the stream ends
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:404
// before headers are received, returns true, nil.  Client-side only.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:407
func (s *Stream) TrailersOnly() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:407
	_go_fuzz_dep_.CoverTab[78675]++
													s.waitOnHeader()
													return s.noHeaders
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:409
	// _ = "end of CoverTab[78675]"
}

// Trailer returns the cached trailer metedata. Note that if it is not called
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:412
// after the entire stream is done, it could return an empty MD. Client
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:412
// side only.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:412
// It can be safely read only after stream has ended that is either read
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:412
// or write have returned io.EOF.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:417
func (s *Stream) Trailer() metadata.MD {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:417
	_go_fuzz_dep_.CoverTab[78676]++
													c := s.trailer.Copy()
													return c
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:419
	// _ = "end of CoverTab[78676]"
}

// ContentSubtype returns the content-subtype for a request. For example, a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:422
// content-subtype of "proto" will result in a content-type of
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:422
// "application/grpc+proto". This will always be lowercase.  See
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:422
// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#requests for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:422
// more details.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:427
func (s *Stream) ContentSubtype() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:427
	_go_fuzz_dep_.CoverTab[78677]++
													return s.contentSubtype
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:428
	// _ = "end of CoverTab[78677]"
}

// Context returns the context of the stream.
func (s *Stream) Context() context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:432
	_go_fuzz_dep_.CoverTab[78678]++
													return s.ctx
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:433
	// _ = "end of CoverTab[78678]"
}

// Method returns the method for the stream.
func (s *Stream) Method() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:437
	_go_fuzz_dep_.CoverTab[78679]++
													return s.method
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:438
	// _ = "end of CoverTab[78679]"
}

// Status returns the status received from the server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:441
// Status can be read safely only after the stream has ended,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:441
// that is, after Done() is closed.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:444
func (s *Stream) Status() *status.Status {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:444
	_go_fuzz_dep_.CoverTab[78680]++
													return s.status
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:445
	// _ = "end of CoverTab[78680]"
}

// SetHeader sets the header metadata. This can be called multiple times.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:448
// Server side only.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:448
// This should not be called in parallel to other data writes.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:451
func (s *Stream) SetHeader(md metadata.MD) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:451
	_go_fuzz_dep_.CoverTab[78681]++
													if md.Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:452
		_go_fuzz_dep_.CoverTab[78684]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:453
		// _ = "end of CoverTab[78684]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:454
		_go_fuzz_dep_.CoverTab[78685]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:454
		// _ = "end of CoverTab[78685]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:454
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:454
	// _ = "end of CoverTab[78681]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:454
	_go_fuzz_dep_.CoverTab[78682]++
													if s.isHeaderSent() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:455
		_go_fuzz_dep_.CoverTab[78686]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:455
		return s.getState() == streamDone
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:455
		// _ = "end of CoverTab[78686]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:455
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:455
		_go_fuzz_dep_.CoverTab[78687]++
														return ErrIllegalHeaderWrite
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:456
		// _ = "end of CoverTab[78687]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:457
		_go_fuzz_dep_.CoverTab[78688]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:457
		// _ = "end of CoverTab[78688]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:457
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:457
	// _ = "end of CoverTab[78682]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:457
	_go_fuzz_dep_.CoverTab[78683]++
													s.hdrMu.Lock()
													s.header = metadata.Join(s.header, md)
													s.hdrMu.Unlock()
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:461
	// _ = "end of CoverTab[78683]"
}

// SendHeader sends the given header metadata. The given metadata is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:464
// combined with any metadata set by previous calls to SetHeader and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:464
// then written to the transport stream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:467
func (s *Stream) SendHeader(md metadata.MD) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:467
	_go_fuzz_dep_.CoverTab[78689]++
													return s.st.WriteHeader(s, md)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:468
	// _ = "end of CoverTab[78689]"
}

// SetTrailer sets the trailer metadata which will be sent with the RPC status
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:471
// by the server. This can be called multiple times. Server side only.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:471
// This should not be called parallel to other data writes.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:474
func (s *Stream) SetTrailer(md metadata.MD) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:474
	_go_fuzz_dep_.CoverTab[78690]++
													if md.Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:475
		_go_fuzz_dep_.CoverTab[78693]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:476
		// _ = "end of CoverTab[78693]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:477
		_go_fuzz_dep_.CoverTab[78694]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:477
		// _ = "end of CoverTab[78694]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:477
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:477
	// _ = "end of CoverTab[78690]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:477
	_go_fuzz_dep_.CoverTab[78691]++
													if s.getState() == streamDone {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:478
		_go_fuzz_dep_.CoverTab[78695]++
														return ErrIllegalHeaderWrite
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:479
		// _ = "end of CoverTab[78695]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:480
		_go_fuzz_dep_.CoverTab[78696]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:480
		// _ = "end of CoverTab[78696]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:480
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:480
	// _ = "end of CoverTab[78691]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:480
	_go_fuzz_dep_.CoverTab[78692]++
													s.hdrMu.Lock()
													s.trailer = metadata.Join(s.trailer, md)
													s.hdrMu.Unlock()
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:484
	// _ = "end of CoverTab[78692]"
}

func (s *Stream) write(m recvMsg) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:487
	_go_fuzz_dep_.CoverTab[78697]++
													s.buf.put(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:488
	// _ = "end of CoverTab[78697]"
}

// Read reads all p bytes from the wire for this stream.
func (s *Stream) Read(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:492
	_go_fuzz_dep_.CoverTab[78698]++

													if er := s.trReader.(*transportReader).er; er != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:494
		_go_fuzz_dep_.CoverTab[78700]++
														return 0, er
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:495
		// _ = "end of CoverTab[78700]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:496
		_go_fuzz_dep_.CoverTab[78701]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:496
		// _ = "end of CoverTab[78701]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:496
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:496
	// _ = "end of CoverTab[78698]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:496
	_go_fuzz_dep_.CoverTab[78699]++
													s.requestRead(len(p))
													return io.ReadFull(s.trReader, p)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:498
	// _ = "end of CoverTab[78699]"
}

// tranportReader reads all the data available for this Stream from the transport and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:501
// passes them into the decoder, which converts them into a gRPC message stream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:501
// The error is io.EOF when the stream is done or another non-nil error if
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:501
// the stream broke.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:505
type transportReader struct {
	reader	io.Reader
	// The handler to control the window update procedure for both this
	// particular stream and the associated transport.
	windowHandler	func(int)
	er		error
}

func (t *transportReader) Read(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:513
	_go_fuzz_dep_.CoverTab[78702]++
													n, err = t.reader.Read(p)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:515
		_go_fuzz_dep_.CoverTab[78704]++
														t.er = err
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:517
		// _ = "end of CoverTab[78704]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:518
		_go_fuzz_dep_.CoverTab[78705]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:518
		// _ = "end of CoverTab[78705]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:518
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:518
	// _ = "end of CoverTab[78702]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:518
	_go_fuzz_dep_.CoverTab[78703]++
													t.windowHandler(n)
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:520
	// _ = "end of CoverTab[78703]"
}

// BytesReceived indicates whether any bytes have been received on this stream.
func (s *Stream) BytesReceived() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:524
	_go_fuzz_dep_.CoverTab[78706]++
													return atomic.LoadUint32(&s.bytesReceived) == 1
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:525
	// _ = "end of CoverTab[78706]"
}

// Unprocessed indicates whether the server did not process this stream --
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:528
// i.e. it sent a refused stream or GOAWAY including this stream ID.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:530
func (s *Stream) Unprocessed() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:530
	_go_fuzz_dep_.CoverTab[78707]++
													return atomic.LoadUint32(&s.unprocessed) == 1
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:531
	// _ = "end of CoverTab[78707]"
}

// GoString is implemented by Stream so context.String() won't
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:534
// race when printing %#v.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:536
func (s *Stream) GoString() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:536
	_go_fuzz_dep_.CoverTab[78708]++
													return fmt.Sprintf("<stream: %p, %v>", s, s.method)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:537
	// _ = "end of CoverTab[78708]"
}

// state of transport
type transportState int

const (
	reachable	transportState	= iota
	closing
	draining
)

// ServerConfig consists of all the configurations to establish a server transport.
type ServerConfig struct {
	MaxStreams		uint32
	ConnectionTimeout	time.Duration
	Credentials		credentials.TransportCredentials
	InTapHandle		tap.ServerInHandle
	StatsHandlers		[]stats.Handler
	KeepaliveParams		keepalive.ServerParameters
	KeepalivePolicy		keepalive.EnforcementPolicy
	InitialWindowSize	int32
	InitialConnWindowSize	int32
	WriteBufferSize		int
	ReadBufferSize		int
	ChannelzParentID	*channelz.Identifier
	MaxHeaderListSize	*uint32
	HeaderTableSize		*uint32
}

// ConnectOptions covers all relevant options for communicating with the server.
type ConnectOptions struct {
	// UserAgent is the application user agent.
	UserAgent	string
	// Dialer specifies how to dial a network address.
	Dialer	func(context.Context, string) (net.Conn, error)
	// FailOnNonTempDialError specifies if gRPC fails on non-temporary dial errors.
	FailOnNonTempDialError	bool
	// PerRPCCredentials stores the PerRPCCredentials required to issue RPCs.
	PerRPCCredentials	[]credentials.PerRPCCredentials
	// TransportCredentials stores the Authenticator required to setup a client
	// connection. Only one of TransportCredentials and CredsBundle is non-nil.
	TransportCredentials	credentials.TransportCredentials
	// CredsBundle is the credentials bundle to be used. Only one of
	// TransportCredentials and CredsBundle is non-nil.
	CredsBundle	credentials.Bundle
	// KeepaliveParams stores the keepalive parameters.
	KeepaliveParams	keepalive.ClientParameters
	// StatsHandlers stores the handler for stats.
	StatsHandlers	[]stats.Handler
	// InitialWindowSize sets the initial window size for a stream.
	InitialWindowSize	int32
	// InitialConnWindowSize sets the initial window size for a connection.
	InitialConnWindowSize	int32
	// WriteBufferSize sets the size of write buffer which in turn determines how much data can be batched before it's written on the wire.
	WriteBufferSize	int
	// ReadBufferSize sets the size of read buffer, which in turn determines how much data can be read at most for one read syscall.
	ReadBufferSize	int
	// ChannelzParentID sets the addrConn id which initiate the creation of this client transport.
	ChannelzParentID	*channelz.Identifier
	// MaxHeaderListSize sets the max (uncompressed) size of header list that is prepared to be received.
	MaxHeaderListSize	*uint32
	// UseProxy specifies if a proxy should be used.
	UseProxy	bool
}

// NewClientTransport establishes the transport with the required ConnectOptions
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:603
// and returns it to the caller.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:605
func NewClientTransport(connectCtx, ctx context.Context, addr resolver.Address, opts ConnectOptions, onClose func(GoAwayReason)) (ClientTransport, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:605
	_go_fuzz_dep_.CoverTab[78709]++
													return newHTTP2Client(connectCtx, ctx, addr, opts, onClose)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:606
	// _ = "end of CoverTab[78709]"
}

// Options provides additional hints and information for message
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:609
// transmission.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:611
type Options struct {
	// Last indicates whether this write is the last piece for
	// this stream.
	Last bool
}

// CallHdr carries the information of a particular RPC.
type CallHdr struct {
	// Host specifies the peer's host.
	Host	string

	// Method specifies the operation to perform.
	Method	string

	// SendCompress specifies the compression algorithm applied on
	// outbound message.
	SendCompress	string

	// Creds specifies credentials.PerRPCCredentials for a call.
	Creds	credentials.PerRPCCredentials

	// ContentSubtype specifies the content-subtype for a request. For example, a
	// content-subtype of "proto" will result in a content-type of
	// "application/grpc+proto". The value of ContentSubtype must be all
	// lowercase, otherwise the behavior is undefined. See
	// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#requests
	// for more details.
	ContentSubtype	string

	PreviousAttempts	int	// value of grpc-previous-rpc-attempts header to set

	DoneFunc	func()	// called when the stream is finished
}

// ClientTransport is the common interface for all gRPC client-side transport
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:645
// implementations.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:647
type ClientTransport interface {
	// Close tears down this transport. Once it returns, the transport
	// should not be accessed any more. The caller must make sure this
	// is called only once.
	Close(err error)

	// GracefulClose starts to tear down the transport: the transport will stop
	// accepting new RPCs and NewStream will return error. Once all streams are
	// finished, the transport will close.
	//
	// It does not block.
	GracefulClose()

	// Write sends the data for the given stream. A nil stream indicates
	// the write is to be performed on the transport as a whole.
	Write(s *Stream, hdr []byte, data []byte, opts *Options) error

	// NewStream creates a Stream for an RPC.
	NewStream(ctx context.Context, callHdr *CallHdr) (*Stream, error)

	// CloseStream clears the footprint of a stream when the stream is
	// not needed any more. The err indicates the error incurred when
	// CloseStream is called. Must be called when a stream is finished
	// unless the associated transport is closing.
	CloseStream(stream *Stream, err error)

	// Error returns a channel that is closed when some I/O error
	// happens. Typically the caller should have a goroutine to monitor
	// this in order to take action (e.g., close the current transport
	// and create a new one) in error case. It should not return nil
	// once the transport is initiated.
	Error() <-chan struct{}

	// GoAway returns a channel that is closed when ClientTransport
	// receives the draining signal from the server (e.g., GOAWAY frame in
	// HTTP/2).
	GoAway() <-chan struct{}

	// GetGoAwayReason returns the reason why GoAway frame was received, along
	// with a human readable string with debug info.
	GetGoAwayReason() (GoAwayReason, string)

	// RemoteAddr returns the remote network address.
	RemoteAddr() net.Addr

	// IncrMsgSent increments the number of message sent through this transport.
	IncrMsgSent()

	// IncrMsgRecv increments the number of message received through this transport.
	IncrMsgRecv()
}

// ServerTransport is the common interface for all gRPC server-side transport
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:699
// implementations.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:699
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:699
// Methods may be called concurrently from multiple goroutines, but
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:699
// Write methods for a given Stream will be called serially.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:704
type ServerTransport interface {
	// HandleStreams receives incoming streams using the given handler.
	HandleStreams(func(*Stream), func(context.Context, string) context.Context)

	// WriteHeader sends the header metadata for the given stream.
	// WriteHeader may not be called on all streams.
	WriteHeader(s *Stream, md metadata.MD) error

	// Write sends the data for the given stream.
	// Write may not be called on all streams.
	Write(s *Stream, hdr []byte, data []byte, opts *Options) error

	// WriteStatus sends the status of a stream to the client.  WriteStatus is
	// the final call made on a stream and always occurs.
	WriteStatus(s *Stream, st *status.Status) error

	// Close tears down the transport. Once it is called, the transport
	// should not be accessed any more. All the pending streams and their
	// handlers will be terminated asynchronously.
	Close(err error)

	// RemoteAddr returns the remote network address.
	RemoteAddr() net.Addr

	// Drain notifies the client this ServerTransport stops accepting new RPCs.
	Drain()

	// IncrMsgSent increments the number of message sent through this transport.
	IncrMsgSent()

	// IncrMsgRecv increments the number of message received through this transport.
	IncrMsgRecv()
}

// connectionErrorf creates an ConnectionError with the specified error description.
func connectionErrorf(temp bool, e error, format string, a ...interface{}) ConnectionError {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:739
	_go_fuzz_dep_.CoverTab[78710]++
													return ConnectionError{
		Desc:	fmt.Sprintf(format, a...),
		temp:	temp,
		err:	e,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:744
	// _ = "end of CoverTab[78710]"
}

// ConnectionError is an error that results in the termination of the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:747
// entire connection and the retry of all the active streams.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:749
type ConnectionError struct {
	Desc	string
	temp	bool
	err	error
}

func (e ConnectionError) Error() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:755
	_go_fuzz_dep_.CoverTab[78711]++
													return fmt.Sprintf("connection error: desc = %q", e.Desc)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:756
	// _ = "end of CoverTab[78711]"
}

// Temporary indicates if this connection error is temporary or fatal.
func (e ConnectionError) Temporary() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:760
	_go_fuzz_dep_.CoverTab[78712]++
													return e.temp
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:761
	// _ = "end of CoverTab[78712]"
}

// Origin returns the original error of this connection error.
func (e ConnectionError) Origin() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:765
	_go_fuzz_dep_.CoverTab[78713]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:768
	if e.err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:768
		_go_fuzz_dep_.CoverTab[78715]++
														return e
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:769
		// _ = "end of CoverTab[78715]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:770
		_go_fuzz_dep_.CoverTab[78716]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:770
		// _ = "end of CoverTab[78716]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:770
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:770
	// _ = "end of CoverTab[78713]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:770
	_go_fuzz_dep_.CoverTab[78714]++
													return e.err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:771
	// _ = "end of CoverTab[78714]"
}

// Unwrap returns the original error of this connection error or nil when the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:774
// origin is nil.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:776
func (e ConnectionError) Unwrap() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:776
	_go_fuzz_dep_.CoverTab[78717]++
													return e.err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:777
	// _ = "end of CoverTab[78717]"
}

var (
	// ErrConnClosing indicates that the transport is closing.
	ErrConnClosing	= connectionErrorf(true, nil, "transport is closing")
	// errStreamDrain indicates that the stream is rejected because the
	// connection is draining. This could be caused by goaway or balancer
	// removing the address.
	errStreamDrain	= status.Error(codes.Unavailable, "the connection is draining")
	// errStreamDone is returned from write at the client side to indiacte application
	// layer of an error.
	errStreamDone	= errors.New("the stream is done")
	// StatusGoAway indicates that the server sent a GOAWAY that included this
	// stream's ID in unprocessed RPCs.
	statusGoAway	= status.New(codes.Unavailable, "the stream is rejected because server is draining the connection")
)

// GoAwayReason contains the reason for the GoAway frame received.
type GoAwayReason uint8

const (
	// GoAwayInvalid indicates that no GoAway frame is received.
	GoAwayInvalid	GoAwayReason	= 0
	// GoAwayNoReason is the default value when GoAway frame is received.
	GoAwayNoReason	GoAwayReason	= 1
	// GoAwayTooManyPings indicates that a GoAway frame with
	// ErrCodeEnhanceYourCalm was received and that the debug data said
	// "too_many_pings".
	GoAwayTooManyPings	GoAwayReason	= 2
)

// channelzData is used to store channelz related data for http2Client and http2Server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:809
// These fields cannot be embedded in the original structs (e.g. http2Client), since to do atomic
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:809
// operation on int64 variable on 32-bit machine, user is responsible to enforce memory alignment.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:809
// Here, by grouping those int64 fields inside a struct, we are enforcing the alignment.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:813
type channelzData struct {
	kpCount	int64
	// The number of streams that have started, including already finished ones.
	streamsStarted	int64
	// Client side: The number of streams that have ended successfully by receiving
	// EoS bit set frame from server.
	// Server side: The number of streams that have ended successfully by sending
	// frame with EoS bit set.
	streamsSucceeded	int64
	streamsFailed		int64
	// lastStreamCreatedTime stores the timestamp that the last stream gets created. It is of int64 type
	// instead of time.Time since it's more costly to atomically update time.Time variable than int64
	// variable. The same goes for lastMsgSentTime and lastMsgRecvTime.
	lastStreamCreatedTime	int64
	msgSent			int64
	msgRecv			int64
	lastMsgSentTime		int64
	lastMsgRecvTime		int64
}

// ContextErr converts the error from context package into a status error.
func ContextErr(err error) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:834
	_go_fuzz_dep_.CoverTab[78718]++
													switch err {
	case context.DeadlineExceeded:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:836
		_go_fuzz_dep_.CoverTab[78720]++
														return status.Error(codes.DeadlineExceeded, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:837
		// _ = "end of CoverTab[78720]"
	case context.Canceled:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:838
		_go_fuzz_dep_.CoverTab[78721]++
														return status.Error(codes.Canceled, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:839
		// _ = "end of CoverTab[78721]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:839
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:839
		_go_fuzz_dep_.CoverTab[78722]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:839
		// _ = "end of CoverTab[78722]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:840
	// _ = "end of CoverTab[78718]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:840
	_go_fuzz_dep_.CoverTab[78719]++
													return status.Errorf(codes.Internal, "Unexpected error from context packet: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:841
	// _ = "end of CoverTab[78719]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:842
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/transport.go:842
var _ = _go_fuzz_dep_.CoverTab
