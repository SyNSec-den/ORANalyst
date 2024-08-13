// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package broker

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
	ch := make(chan e2api.Indication)
	return &bufferedStream{
		bufferedIO: &bufferedIO{
			streamID:  streamID,
			channelID: channelID,
			subSepc:   subSpec,
			node:      node,
			subName:   subName,
		},
		bufferedReader: newBufferedReader(ch),
		bufferedWriter: newBufferedWriter(ch),
	}
}

type bufferedIO struct {
	subSepc   e2api.SubscriptionSpec
	streamID  StreamID
	channelID e2api.ChannelID
	node      e2client.Node
	subName   string
}

func (s *bufferedIO) SubscriptionName() string {
	return s.subName
}

func (s *bufferedIO) Subscription() e2api.SubscriptionSpec {
	return s.subSepc
}

func (s *bufferedIO) ChannelID() e2api.ChannelID {
	return s.channelID
}

func (s *bufferedIO) Node() e2client.Node {
	return s.node
}

func (s *bufferedIO) StreamID() StreamID {
	return s.streamID
}

type bufferedStream struct {
	*bufferedIO
	*bufferedReader
	*bufferedWriter
}

var _ Stream = &bufferedStream{}

func newBufferedReader(ch <-chan e2api.Indication) *bufferedReader {
	return &bufferedReader{
		ch: ch,
	}
}

type bufferedReader struct {
	ch <-chan e2api.Indication
}

func (s *bufferedReader) Recv(ctx context.Context) (e2api.Indication, error) {
	select {
	case ind, ok := <-s.ch:
		if !ok {
			return e2api.Indication{}, io.EOF
		}
		return ind, nil
	case <-ctx.Done():
		return e2api.Indication{}, ctx.Err()
	}
}

func newBufferedWriter(ch chan<- e2api.Indication) *bufferedWriter {
	writer := &bufferedWriter{
		ch:     ch,
		buffer: list.New(),
		cond:   sync.NewCond(&sync.Mutex{}),
	}
	writer.open()
	return writer
}

type bufferedWriter struct {
	ch     chan<- e2api.Indication
	buffer *list.List
	cond   *sync.Cond
	closed bool
}

// open starts the goroutine propagating indications from the writer to the reader
func (s *bufferedWriter) open() {
	go s.drain()
}

// drain dequeues indications and writes them to the read channel
func (s *bufferedWriter) drain() {
	for {
		ind, ok := s.next()
		if !ok {
			close(s.ch)
			break
		}
		s.ch <- ind
	}
}

// next reads the next indication from the buffer or blocks until one becomes available
func (s *bufferedWriter) next() (e2api.Indication, bool) {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()
	for s.buffer.Len() == 0 {
		if s.closed {
			return e2api.Indication{}, false
		}
		s.cond.Wait()
	}
	result := s.buffer.Front().Value.(e2api.Indication)
	s.buffer.Remove(s.buffer.Front())
	return result, true
}

// Send appends the indication to the buffer and notifies the reader
func (s *bufferedWriter) Send(ind e2api.Indication) error {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()
	if s.closed {
		return io.EOF
	}
	if s.buffer.Len() == bufferMaxSize {
		return errors.NewUnavailable("cannot append indication to stream: maximum buffer size has been reached")
	}
	s.buffer.PushBack(ind)
	s.cond.Signal()
	return nil
}

func (s *bufferedWriter) Close() error {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()
	s.closed = true
	return nil
}
