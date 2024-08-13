// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:5
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:5
)

import "fmt"

// WriteScheduler is the interface implemented by HTTP/2 write schedulers.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:9
// Methods are never called concurrently.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:11
type WriteScheduler interface {
	// OpenStream opens a new stream in the write scheduler.
	// It is illegal to call this with streamID=0 or with a streamID that is
	// already open -- the call may panic.
	OpenStream(streamID uint32, options OpenStreamOptions)

	// CloseStream closes a stream in the write scheduler. Any frames queued on
	// this stream should be discarded. It is illegal to call this on a stream
	// that is not open -- the call may panic.
	CloseStream(streamID uint32)

	// AdjustStream adjusts the priority of the given stream. This may be called
	// on a stream that has not yet been opened or has been closed. Note that
	// RFC 7540 allows PRIORITY frames to be sent on streams in any state. See:
	// https://tools.ietf.org/html/rfc7540#section-5.1
	AdjustStream(streamID uint32, priority PriorityParam)

	// Push queues a frame in the scheduler. In most cases, this will not be
	// called with wr.StreamID()!=0 unless that stream is currently open. The one
	// exception is RST_STREAM frames, which may be sent on idle or closed streams.
	Push(wr FrameWriteRequest)

	// Pop dequeues the next frame to write. Returns false if no frames can
	// be written. Frames with a given wr.StreamID() are Pop'd in the same
	// order they are Push'd, except RST_STREAM frames. No frames should be
	// discarded except by CloseStream.
	Pop() (wr FrameWriteRequest, ok bool)
}

// OpenStreamOptions specifies extra options for WriteScheduler.OpenStream.
type OpenStreamOptions struct {
	// PusherID is zero if the stream was initiated by the client. Otherwise,
	// PusherID names the stream that pushed the newly opened stream.
	PusherID uint32
}

// FrameWriteRequest is a request to write a frame.
type FrameWriteRequest struct {
	// write is the interface value that does the writing, once the
	// WriteScheduler has selected this frame to write. The write
	// functions are all defined in write.go.
	write	writeFramer

	// stream is the stream on which this frame will be written.
	// nil for non-stream frames like PING and SETTINGS.
	// nil for RST_STREAM streams, which use the StreamError.StreamID field instead.
	stream	*stream

	// done, if non-nil, must be a buffered channel with space for
	// 1 message and is sent the return value from write (or an
	// earlier error) when the frame has been written.
	done	chan error
}

// StreamID returns the id of the stream this frame will be written to.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:65
// 0 is used for non-stream frames such as PING and SETTINGS.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:67
func (wr FrameWriteRequest) StreamID() uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:67
	_go_fuzz_dep_.CoverTab[75744]++
											if wr.stream == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:68
		_go_fuzz_dep_.CoverTab[75746]++
												if se, ok := wr.write.(StreamError); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:69
			_go_fuzz_dep_.CoverTab[75748]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:74
			return se.StreamID
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:74
			// _ = "end of CoverTab[75748]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:75
			_go_fuzz_dep_.CoverTab[75749]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:75
			// _ = "end of CoverTab[75749]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:75
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:75
		// _ = "end of CoverTab[75746]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:75
		_go_fuzz_dep_.CoverTab[75747]++
												return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:76
		// _ = "end of CoverTab[75747]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:77
		_go_fuzz_dep_.CoverTab[75750]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:77
		// _ = "end of CoverTab[75750]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:77
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:77
	// _ = "end of CoverTab[75744]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:77
	_go_fuzz_dep_.CoverTab[75745]++
											return wr.stream.id
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:78
	// _ = "end of CoverTab[75745]"
}

// isControl reports whether wr is a control frame for MaxQueuedControlFrames
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:81
// purposes. That includes non-stream frames and RST_STREAM frames.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:83
func (wr FrameWriteRequest) isControl() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:83
	_go_fuzz_dep_.CoverTab[75751]++
											return wr.stream == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:84
	// _ = "end of CoverTab[75751]"
}

// DataSize returns the number of flow control bytes that must be consumed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:87
// to write this entire frame. This is 0 for non-DATA frames.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:89
func (wr FrameWriteRequest) DataSize() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:89
	_go_fuzz_dep_.CoverTab[75752]++
											if wd, ok := wr.write.(*writeData); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:90
		_go_fuzz_dep_.CoverTab[75754]++
												return len(wd.p)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:91
		// _ = "end of CoverTab[75754]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:92
		_go_fuzz_dep_.CoverTab[75755]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:92
		// _ = "end of CoverTab[75755]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:92
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:92
	// _ = "end of CoverTab[75752]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:92
	_go_fuzz_dep_.CoverTab[75753]++
											return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:93
	// _ = "end of CoverTab[75753]"
}

// Consume consumes min(n, available) bytes from this frame, where available
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:96
// is the number of flow control bytes available on the stream. Consume returns
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:96
// 0, 1, or 2 frames, where the integer return value gives the number of frames
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:96
// returned.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:96
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:96
// If flow control prevents consuming any bytes, this returns (_, _, 0). If
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:96
// the entire frame was consumed, this returns (wr, _, 1). Otherwise, this
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:96
// returns (consumed, rest, 2), where 'consumed' contains the consumed bytes and
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:96
// 'rest' contains the remaining bytes. The consumed bytes are deducted from the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:96
// underlying stream's flow control budget.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:106
func (wr FrameWriteRequest) Consume(n int32) (FrameWriteRequest, FrameWriteRequest, int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:106
	_go_fuzz_dep_.CoverTab[75756]++
											var empty FrameWriteRequest

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:110
	wd, ok := wr.write.(*writeData)
	if !ok || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:111
		_go_fuzz_dep_.CoverTab[75762]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:111
		return len(wd.p) == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:111
		// _ = "end of CoverTab[75762]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:111
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:111
		_go_fuzz_dep_.CoverTab[75763]++
												return wr, empty, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:112
		// _ = "end of CoverTab[75763]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:113
		_go_fuzz_dep_.CoverTab[75764]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:113
		// _ = "end of CoverTab[75764]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:113
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:113
	// _ = "end of CoverTab[75756]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:113
	_go_fuzz_dep_.CoverTab[75757]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:116
	allowed := wr.stream.flow.available()
	if n < allowed {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:117
		_go_fuzz_dep_.CoverTab[75765]++
												allowed = n
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:118
		// _ = "end of CoverTab[75765]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:119
		_go_fuzz_dep_.CoverTab[75766]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:119
		// _ = "end of CoverTab[75766]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:119
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:119
	// _ = "end of CoverTab[75757]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:119
	_go_fuzz_dep_.CoverTab[75758]++
											if wr.stream.sc.maxFrameSize < allowed {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:120
		_go_fuzz_dep_.CoverTab[75767]++
												allowed = wr.stream.sc.maxFrameSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:121
		// _ = "end of CoverTab[75767]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:122
		_go_fuzz_dep_.CoverTab[75768]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:122
		// _ = "end of CoverTab[75768]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:122
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:122
	// _ = "end of CoverTab[75758]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:122
	_go_fuzz_dep_.CoverTab[75759]++
											if allowed <= 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:123
		_go_fuzz_dep_.CoverTab[75769]++
												return empty, empty, 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:124
		// _ = "end of CoverTab[75769]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:125
		_go_fuzz_dep_.CoverTab[75770]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:125
		// _ = "end of CoverTab[75770]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:125
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:125
	// _ = "end of CoverTab[75759]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:125
	_go_fuzz_dep_.CoverTab[75760]++
											if len(wd.p) > int(allowed) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:126
		_go_fuzz_dep_.CoverTab[75771]++
												wr.stream.flow.take(allowed)
												consumed := FrameWriteRequest{
			stream:	wr.stream,
			write: &writeData{
														streamID:	wd.streamID,
														p:		wd.p[:allowed],

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:136
				endStream:	false,
			},

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:140
			done:	nil,
		}
		rest := FrameWriteRequest{
			stream:	wr.stream,
			write: &writeData{
				streamID:	wd.streamID,
				p:		wd.p[allowed:],
				endStream:	wd.endStream,
			},
			done:	wr.done,
		}
												return consumed, rest, 2
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:151
		// _ = "end of CoverTab[75771]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:152
		_go_fuzz_dep_.CoverTab[75772]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:152
		// _ = "end of CoverTab[75772]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:152
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:152
	// _ = "end of CoverTab[75760]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:152
	_go_fuzz_dep_.CoverTab[75761]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:156
	wr.stream.flow.take(int32(len(wd.p)))
											return wr, empty, 1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:157
	// _ = "end of CoverTab[75761]"
}

// String is for debugging only.
func (wr FrameWriteRequest) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:161
	_go_fuzz_dep_.CoverTab[75773]++
											var des string
											if s, ok := wr.write.(fmt.Stringer); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:163
		_go_fuzz_dep_.CoverTab[75775]++
												des = s.String()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:164
		// _ = "end of CoverTab[75775]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:165
		_go_fuzz_dep_.CoverTab[75776]++
												des = fmt.Sprintf("%T", wr.write)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:166
		// _ = "end of CoverTab[75776]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:167
	// _ = "end of CoverTab[75773]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:167
	_go_fuzz_dep_.CoverTab[75774]++
											return fmt.Sprintf("[FrameWriteRequest stream=%d, ch=%v, writer=%v]", wr.StreamID(), wr.done != nil, des)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:168
	// _ = "end of CoverTab[75774]"
}

// replyToWriter sends err to wr.done and panics if the send must block
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:171
// This does nothing if wr.done is nil.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:173
func (wr *FrameWriteRequest) replyToWriter(err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:173
	_go_fuzz_dep_.CoverTab[75777]++
											if wr.done == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:174
		_go_fuzz_dep_.CoverTab[75780]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:175
		// _ = "end of CoverTab[75780]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:176
		_go_fuzz_dep_.CoverTab[75781]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:176
		// _ = "end of CoverTab[75781]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:176
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:176
	// _ = "end of CoverTab[75777]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:176
	_go_fuzz_dep_.CoverTab[75778]++
											select {
	case wr.done <- err:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:178
		_go_fuzz_dep_.CoverTab[75782]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:178
		// _ = "end of CoverTab[75782]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:179
		_go_fuzz_dep_.CoverTab[75783]++
												panic(fmt.Sprintf("unbuffered done channel passed in for type %T", wr.write))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:180
		// _ = "end of CoverTab[75783]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:181
	// _ = "end of CoverTab[75778]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:181
	_go_fuzz_dep_.CoverTab[75779]++
											wr.write = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:182
	// _ = "end of CoverTab[75779]"
}

// writeQueue is used by implementations of WriteScheduler.
type writeQueue struct {
	s []FrameWriteRequest
}

func (q *writeQueue) empty() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:190
	_go_fuzz_dep_.CoverTab[75784]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:190
	return len(q.s) == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:190
	// _ = "end of CoverTab[75784]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:190
}

func (q *writeQueue) push(wr FrameWriteRequest) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:192
	_go_fuzz_dep_.CoverTab[75785]++
											q.s = append(q.s, wr)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:193
	// _ = "end of CoverTab[75785]"
}

func (q *writeQueue) shift() FrameWriteRequest {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:196
	_go_fuzz_dep_.CoverTab[75786]++
											if len(q.s) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:197
		_go_fuzz_dep_.CoverTab[75788]++
												panic("invalid use of queue")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:198
		// _ = "end of CoverTab[75788]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:199
		_go_fuzz_dep_.CoverTab[75789]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:199
		// _ = "end of CoverTab[75789]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:199
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:199
	// _ = "end of CoverTab[75786]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:199
	_go_fuzz_dep_.CoverTab[75787]++
											wr := q.s[0]

											copy(q.s, q.s[1:])
											q.s[len(q.s)-1] = FrameWriteRequest{}
											q.s = q.s[:len(q.s)-1]
											return wr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:205
	// _ = "end of CoverTab[75787]"
}

// consume consumes up to n bytes from q.s[0]. If the frame is
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:208
// entirely consumed, it is removed from the queue. If the frame
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:208
// is partially consumed, the frame is kept with the consumed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:208
// bytes removed. Returns true iff any bytes were consumed.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:212
func (q *writeQueue) consume(n int32) (FrameWriteRequest, bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:212
	_go_fuzz_dep_.CoverTab[75790]++
											if len(q.s) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:213
		_go_fuzz_dep_.CoverTab[75793]++
												return FrameWriteRequest{}, false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:214
		// _ = "end of CoverTab[75793]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:215
		_go_fuzz_dep_.CoverTab[75794]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:215
		// _ = "end of CoverTab[75794]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:215
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:215
	// _ = "end of CoverTab[75790]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:215
	_go_fuzz_dep_.CoverTab[75791]++
											consumed, rest, numresult := q.s[0].Consume(n)
											switch numresult {
	case 0:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:218
		_go_fuzz_dep_.CoverTab[75795]++
												return FrameWriteRequest{}, false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:219
		// _ = "end of CoverTab[75795]"
	case 1:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:220
		_go_fuzz_dep_.CoverTab[75796]++
												q.shift()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:221
		// _ = "end of CoverTab[75796]"
	case 2:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:222
		_go_fuzz_dep_.CoverTab[75797]++
												q.s[0] = rest
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:223
		// _ = "end of CoverTab[75797]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:223
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:223
		_go_fuzz_dep_.CoverTab[75798]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:223
		// _ = "end of CoverTab[75798]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:224
	// _ = "end of CoverTab[75791]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:224
	_go_fuzz_dep_.CoverTab[75792]++
											return consumed, true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:225
	// _ = "end of CoverTab[75792]"
}

type writeQueuePool []*writeQueue

// put inserts an unused writeQueue into the pool.
func (p *writeQueuePool) put(q *writeQueue) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:231
	_go_fuzz_dep_.CoverTab[75799]++
											for i := range q.s {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:232
		_go_fuzz_dep_.CoverTab[75801]++
												q.s[i] = FrameWriteRequest{}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:233
		// _ = "end of CoverTab[75801]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:234
	// _ = "end of CoverTab[75799]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:234
	_go_fuzz_dep_.CoverTab[75800]++
											q.s = q.s[:0]
											*p = append(*p, q)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:236
	// _ = "end of CoverTab[75800]"
}

// get returns an empty writeQueue.
func (p *writeQueuePool) get() *writeQueue {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:240
	_go_fuzz_dep_.CoverTab[75802]++
											ln := len(*p)
											if ln == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:242
		_go_fuzz_dep_.CoverTab[75804]++
												return new(writeQueue)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:243
		// _ = "end of CoverTab[75804]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:244
		_go_fuzz_dep_.CoverTab[75805]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:244
		// _ = "end of CoverTab[75805]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:244
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:244
	// _ = "end of CoverTab[75802]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:244
	_go_fuzz_dep_.CoverTab[75803]++
											x := ln - 1
											q := (*p)[x]
											(*p)[x] = nil
											*p = (*p)[:x]
											return q
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:249
	// _ = "end of CoverTab[75803]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:250
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/writesched.go:250
var _ = _go_fuzz_dep_.CoverTab
