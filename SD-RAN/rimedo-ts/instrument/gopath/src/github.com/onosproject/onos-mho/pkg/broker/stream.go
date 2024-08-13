// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:5
package broker

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:5
)

import (
	"container/list"
	"context"
	"io"
	"sync"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
)

const bufferMaxSize = 10000

// StreamReader defines methods for reading indications from a Stream
type StreamReader interface {
	StreamIO

	// Recv reads an indication from the stream
	// This method is thread-safe. If multiple goroutines are receiving from the stream, indications will be
	// distributed randomly between them. If no indications are available, the goroutine will be blocked until
	// an indication is received or the Context is canceled. If the context is canceled, a context.Canceled error
	// will be returned. If the stream has been closed, an io.EOF error will be returned.
	Recv(context.Context) (e2api.Indication, error)
}

// StreamWriter is a write stream
type StreamWriter interface {
	StreamIO

	// Send sends an indication on the stream
	// The Send method is non-blocking. If no StreamReader is immediately available to consume the indication
	// it will be placed in a bounded memory buffer. If the buffer is full, an Unavailable error will be returned.
	// This method is thread-safe.
	Send(indication e2api.Indication) error
}

// StreamID is a stream identifier
type StreamID int

// StreamIO is a base interface for Stream information
type StreamIO interface {
	io.Closer
	ChannelID() e2api.ChannelID
	StreamID() StreamID
	SubscriptionName() string
	Subscription() e2api.SubscriptionSpec
	Node() e2client.Node
}

// Stream is a read/write stream
type Stream interface {
	StreamIO
	StreamReader
	StreamWriter
}

func newBufferedStream(node e2client.Node, subName string, streamID StreamID, channelID e2api.ChannelID, subSpec e2api.SubscriptionSpec) Stream {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:63
	_go_fuzz_dep_.CoverTab[196653]++
													ch := make(chan e2api.Indication)
													return &bufferedStream{
		bufferedIO: &bufferedIO{
			streamID:	streamID,
			channelID:	channelID,
			subSepc:	subSpec,
			node:		node,
			subName:	subName,
		},
		bufferedReader:	newBufferedReader(ch),
		bufferedWriter:	newBufferedWriter(ch),
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:75
	// _ = "end of CoverTab[196653]"
}

type bufferedIO struct {
	subSepc		e2api.SubscriptionSpec
	streamID	StreamID
	channelID	e2api.ChannelID
	node		e2client.Node
	subName		string
}

func (s *bufferedIO) SubscriptionName() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:86
	_go_fuzz_dep_.CoverTab[196654]++
													return s.subName
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:87
	// _ = "end of CoverTab[196654]"
}

func (s *bufferedIO) Subscription() e2api.SubscriptionSpec {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:90
	_go_fuzz_dep_.CoverTab[196655]++
													return s.subSepc
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:91
	// _ = "end of CoverTab[196655]"
}

func (s *bufferedIO) ChannelID() e2api.ChannelID {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:94
	_go_fuzz_dep_.CoverTab[196656]++
													return s.channelID
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:95
	// _ = "end of CoverTab[196656]"
}

func (s *bufferedIO) Node() e2client.Node {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:98
	_go_fuzz_dep_.CoverTab[196657]++
													return s.node
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:99
	// _ = "end of CoverTab[196657]"
}

func (s *bufferedIO) StreamID() StreamID {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:102
	_go_fuzz_dep_.CoverTab[196658]++
													return s.streamID
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:103
	// _ = "end of CoverTab[196658]"
}

type bufferedStream struct {
	*bufferedIO
	*bufferedReader
	*bufferedWriter
}

var _ Stream = &bufferedStream{}

func newBufferedReader(ch <-chan e2api.Indication) *bufferedReader {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:114
	_go_fuzz_dep_.CoverTab[196659]++
													return &bufferedReader{
		ch: ch,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:117
	// _ = "end of CoverTab[196659]"
}

type bufferedReader struct {
	ch <-chan e2api.Indication
}

func (s *bufferedReader) Recv(ctx context.Context) (e2api.Indication, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:124
	_go_fuzz_dep_.CoverTab[196660]++
													select {
	case ind, ok := <-s.ch:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:126
		_go_fuzz_dep_.CoverTab[196661]++
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:127
			_go_fuzz_dep_.CoverTab[196664]++
															return e2api.Indication{}, io.EOF
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:128
			// _ = "end of CoverTab[196664]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:129
			_go_fuzz_dep_.CoverTab[196665]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:129
			// _ = "end of CoverTab[196665]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:129
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:129
		// _ = "end of CoverTab[196661]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:129
		_go_fuzz_dep_.CoverTab[196662]++
														return ind, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:130
		// _ = "end of CoverTab[196662]"
	case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:131
		_go_fuzz_dep_.CoverTab[196663]++
														return e2api.Indication{}, ctx.Err()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:132
		// _ = "end of CoverTab[196663]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:133
	// _ = "end of CoverTab[196660]"
}

func newBufferedWriter(ch chan<- e2api.Indication) *bufferedWriter {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:136
	_go_fuzz_dep_.CoverTab[196666]++
													writer := &bufferedWriter{
		ch:	ch,
		buffer:	list.New(),
		cond:	sync.NewCond(&sync.Mutex{}),
	}
													writer.open()
													return writer
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:143
	// _ = "end of CoverTab[196666]"
}

type bufferedWriter struct {
	ch	chan<- e2api.Indication
	buffer	*list.List
	cond	*sync.Cond
	closed	bool
}

// open starts the goroutine propagating indications from the writer to the reader
func (s *bufferedWriter) open() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:154
	_go_fuzz_dep_.CoverTab[196667]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:154
	_curRoutineNum187_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:154
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum187_)
													go s.drain()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:155
	// _ = "end of CoverTab[196667]"
}

// drain dequeues indications and writes them to the read channel
func (s *bufferedWriter) drain() {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:159
	_go_fuzz_dep_.CoverTab[196668]++
													for {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:160
		_go_fuzz_dep_.CoverTab[196669]++
														ind, ok := s.next()
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:162
			_go_fuzz_dep_.CoverTab[196671]++
															close(s.ch)
															break
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:164
			// _ = "end of CoverTab[196671]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:165
			_go_fuzz_dep_.CoverTab[196672]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:165
			// _ = "end of CoverTab[196672]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:165
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:165
		// _ = "end of CoverTab[196669]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:165
		_go_fuzz_dep_.CoverTab[196670]++
														s.ch <- ind
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:166
		// _ = "end of CoverTab[196670]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:167
	// _ = "end of CoverTab[196668]"
}

// next reads the next indication from the buffer or blocks until one becomes available
func (s *bufferedWriter) next() (e2api.Indication, bool) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:171
	_go_fuzz_dep_.CoverTab[196673]++
													s.cond.L.Lock()
													defer s.cond.L.Unlock()
													for s.buffer.Len() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:174
		_go_fuzz_dep_.CoverTab[196675]++
														if s.closed {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:175
			_go_fuzz_dep_.CoverTab[196677]++
															return e2api.Indication{}, false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:176
			// _ = "end of CoverTab[196677]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:177
			_go_fuzz_dep_.CoverTab[196678]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:177
			// _ = "end of CoverTab[196678]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:177
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:177
		// _ = "end of CoverTab[196675]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:177
		_go_fuzz_dep_.CoverTab[196676]++
														s.cond.Wait()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:178
		// _ = "end of CoverTab[196676]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:179
	// _ = "end of CoverTab[196673]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:179
	_go_fuzz_dep_.CoverTab[196674]++
													result := s.buffer.Front().Value.(e2api.Indication)
													s.buffer.Remove(s.buffer.Front())
													return result, true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:182
	// _ = "end of CoverTab[196674]"
}

// Send appends the indication to the buffer and notifies the reader
func (s *bufferedWriter) Send(ind e2api.Indication) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:186
	_go_fuzz_dep_.CoverTab[196679]++
													s.cond.L.Lock()
													defer s.cond.L.Unlock()
													if s.closed {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:189
		_go_fuzz_dep_.CoverTab[196682]++
														return io.EOF
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:190
		// _ = "end of CoverTab[196682]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:191
		_go_fuzz_dep_.CoverTab[196683]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:191
		// _ = "end of CoverTab[196683]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:191
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:191
	// _ = "end of CoverTab[196679]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:191
	_go_fuzz_dep_.CoverTab[196680]++
													if s.buffer.Len() == bufferMaxSize {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:192
		_go_fuzz_dep_.CoverTab[196684]++
														return errors.NewUnavailable("cannot append indication to stream: maximum buffer size has been reached")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:193
		// _ = "end of CoverTab[196684]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:194
		_go_fuzz_dep_.CoverTab[196685]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:194
		// _ = "end of CoverTab[196685]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:194
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:194
	// _ = "end of CoverTab[196680]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:194
	_go_fuzz_dep_.CoverTab[196681]++
													s.buffer.PushBack(ind)
													s.cond.Signal()
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:197
	// _ = "end of CoverTab[196681]"
}

func (s *bufferedWriter) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:200
	_go_fuzz_dep_.CoverTab[196686]++
													s.cond.L.Lock()
													defer s.cond.L.Unlock()
													s.closed = true
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:204
	// _ = "end of CoverTab[196686]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:205
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/broker/stream.go:205
var _ = _go_fuzz_dep_.CoverTab
