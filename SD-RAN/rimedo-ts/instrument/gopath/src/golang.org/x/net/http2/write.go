// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:5
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:5
)

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/http/httpguts"
	"golang.org/x/net/http2/hpack"
)

// writeFramer is implemented by any type that is used to write frames.
type writeFramer interface {
	writeFrame(writeContext) error

	// staysWithinBuffer reports whether this writer promises that
	// it will only write less than or equal to size bytes, and it
	// won't Flush the write context.
	staysWithinBuffer(size int) bool
}

// writeContext is the interface needed by the various frame writer
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:28
// types below. All the writeFrame methods below are scheduled via the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:28
// frame writing scheduler (see writeScheduler in writesched.go).
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:28
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:28
// This interface is implemented by *serverConn.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:28
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:28
// TODO: decide whether to a) use this in the client code (which didn't
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:28
// end up using this yet, because it has a simpler design, not
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:28
// currently implementing priorities), or b) delete this and
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:28
// make the server code a bit more concrete.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:38
type writeContext interface {
	Framer() *Framer
	Flush() error
	CloseConn() error
	// HeaderEncoder returns an HPACK encoder that writes to the
	// returned buffer.
	HeaderEncoder() (*hpack.Encoder, *bytes.Buffer)
}

// writeEndsStream reports whether w writes a frame that will transition
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:47
// the stream to a half-closed local state. This returns false for RST_STREAM,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:47
// which closes the entire stream (not just the local half).
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:50
func writeEndsStream(w writeFramer) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:50
	_go_fuzz_dep_.CoverTab[75657]++
										switch v := w.(type) {
	case *writeData:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:52
		_go_fuzz_dep_.CoverTab[75659]++
											return v.endStream
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:53
		// _ = "end of CoverTab[75659]"
	case *writeResHeaders:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:54
		_go_fuzz_dep_.CoverTab[75660]++
											return v.endStream
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:55
		// _ = "end of CoverTab[75660]"
	case nil:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:56
		_go_fuzz_dep_.CoverTab[75661]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:60
		panic("writeEndsStream called on nil writeFramer")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:60
		// _ = "end of CoverTab[75661]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:61
	// _ = "end of CoverTab[75657]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:61
	_go_fuzz_dep_.CoverTab[75658]++
										return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:62
	// _ = "end of CoverTab[75658]"
}

type flushFrameWriter struct{}

func (flushFrameWriter) writeFrame(ctx writeContext) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:67
	_go_fuzz_dep_.CoverTab[75662]++
										return ctx.Flush()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:68
	// _ = "end of CoverTab[75662]"
}

func (flushFrameWriter) staysWithinBuffer(max int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:71
	_go_fuzz_dep_.CoverTab[75663]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:71
	return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:71
	// _ = "end of CoverTab[75663]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:71
}

type writeSettings []Setting

func (s writeSettings) staysWithinBuffer(max int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:75
	_go_fuzz_dep_.CoverTab[75664]++
										const settingSize = 6	// uint16 + uint32
										return frameHeaderLen+settingSize*len(s) <= max
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:77
	// _ = "end of CoverTab[75664]"

}

func (s writeSettings) writeFrame(ctx writeContext) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:81
	_go_fuzz_dep_.CoverTab[75665]++
										return ctx.Framer().WriteSettings([]Setting(s)...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:82
	// _ = "end of CoverTab[75665]"
}

type writeGoAway struct {
	maxStreamID	uint32
	code		ErrCode
}

func (p *writeGoAway) writeFrame(ctx writeContext) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:90
	_go_fuzz_dep_.CoverTab[75666]++
										err := ctx.Framer().WriteGoAway(p.maxStreamID, p.code, nil)
										ctx.Flush()
										return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:93
	// _ = "end of CoverTab[75666]"
}

func (*writeGoAway) staysWithinBuffer(max int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:96
	_go_fuzz_dep_.CoverTab[75667]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:96
	return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:96
	// _ = "end of CoverTab[75667]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:96
}

type writeData struct {
	streamID	uint32
	p		[]byte
	endStream	bool
}

func (w *writeData) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:104
	_go_fuzz_dep_.CoverTab[75668]++
										return fmt.Sprintf("writeData(stream=%d, p=%d, endStream=%v)", w.streamID, len(w.p), w.endStream)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:105
	// _ = "end of CoverTab[75668]"
}

func (w *writeData) writeFrame(ctx writeContext) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:108
	_go_fuzz_dep_.CoverTab[75669]++
										return ctx.Framer().WriteData(w.streamID, w.endStream, w.p)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:109
	// _ = "end of CoverTab[75669]"
}

func (w *writeData) staysWithinBuffer(max int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:112
	_go_fuzz_dep_.CoverTab[75670]++
										return frameHeaderLen+len(w.p) <= max
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:113
	// _ = "end of CoverTab[75670]"
}

// handlerPanicRST is the message sent from handler goroutines when
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:116
// the handler panics.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:118
type handlerPanicRST struct {
	StreamID uint32
}

func (hp handlerPanicRST) writeFrame(ctx writeContext) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:122
	_go_fuzz_dep_.CoverTab[75671]++
										return ctx.Framer().WriteRSTStream(hp.StreamID, ErrCodeInternal)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:123
	// _ = "end of CoverTab[75671]"
}

func (hp handlerPanicRST) staysWithinBuffer(max int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:126
	_go_fuzz_dep_.CoverTab[75672]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:126
	return frameHeaderLen+4 <= max
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:126
	// _ = "end of CoverTab[75672]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:126
}

func (se StreamError) writeFrame(ctx writeContext) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:128
	_go_fuzz_dep_.CoverTab[75673]++
										return ctx.Framer().WriteRSTStream(se.StreamID, se.Code)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:129
	// _ = "end of CoverTab[75673]"
}

func (se StreamError) staysWithinBuffer(max int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:132
	_go_fuzz_dep_.CoverTab[75674]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:132
	return frameHeaderLen+4 <= max
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:132
	// _ = "end of CoverTab[75674]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:132
}

type writePingAck struct{ pf *PingFrame }

func (w writePingAck) writeFrame(ctx writeContext) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:136
	_go_fuzz_dep_.CoverTab[75675]++
										return ctx.Framer().WritePing(true, w.pf.Data)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:137
	// _ = "end of CoverTab[75675]"
}

func (w writePingAck) staysWithinBuffer(max int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:140
	_go_fuzz_dep_.CoverTab[75676]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:140
	return frameHeaderLen+len(w.pf.Data) <= max
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:140
	// _ = "end of CoverTab[75676]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:140
}

type writeSettingsAck struct{}

func (writeSettingsAck) writeFrame(ctx writeContext) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:144
	_go_fuzz_dep_.CoverTab[75677]++
										return ctx.Framer().WriteSettingsAck()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:145
	// _ = "end of CoverTab[75677]"
}

func (writeSettingsAck) staysWithinBuffer(max int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:148
	_go_fuzz_dep_.CoverTab[75678]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:148
	return frameHeaderLen <= max
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:148
	// _ = "end of CoverTab[75678]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:148
}

// splitHeaderBlock splits headerBlock into fragments so that each fragment fits
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:150
// in a single frame, then calls fn for each fragment. firstFrag/lastFrag are true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:150
// for the first/last fragment, respectively.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:153
func splitHeaderBlock(ctx writeContext, headerBlock []byte, fn func(ctx writeContext, frag []byte, firstFrag, lastFrag bool) error) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:153
	_go_fuzz_dep_.CoverTab[75679]++
	// For now we're lazy and just pick the minimum MAX_FRAME_SIZE
	// that all peers must support (16KB). Later we could care
	// more and send larger frames if the peer advertised it, but
	// there's little point. Most headers are small anyway (so we
	// generally won't have CONTINUATION frames), and extra frames
	// only waste 9 bytes anyway.
	const maxFrameSize = 16384

	first := true
	for len(headerBlock) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:163
		_go_fuzz_dep_.CoverTab[75681]++
											frag := headerBlock
											if len(frag) > maxFrameSize {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:165
			_go_fuzz_dep_.CoverTab[75684]++
												frag = frag[:maxFrameSize]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:166
			// _ = "end of CoverTab[75684]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:167
			_go_fuzz_dep_.CoverTab[75685]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:167
			// _ = "end of CoverTab[75685]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:167
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:167
		// _ = "end of CoverTab[75681]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:167
		_go_fuzz_dep_.CoverTab[75682]++
											headerBlock = headerBlock[len(frag):]
											if err := fn(ctx, frag, first, len(headerBlock) == 0); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:169
			_go_fuzz_dep_.CoverTab[75686]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:170
			// _ = "end of CoverTab[75686]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:171
			_go_fuzz_dep_.CoverTab[75687]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:171
			// _ = "end of CoverTab[75687]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:171
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:171
		// _ = "end of CoverTab[75682]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:171
		_go_fuzz_dep_.CoverTab[75683]++
											first = false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:172
		// _ = "end of CoverTab[75683]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:173
	// _ = "end of CoverTab[75679]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:173
	_go_fuzz_dep_.CoverTab[75680]++
										return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:174
	// _ = "end of CoverTab[75680]"
}

// writeResHeaders is a request to write a HEADERS and 0+ CONTINUATION frames
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:177
// for HTTP response headers or trailers from a server handler.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:179
type writeResHeaders struct {
	streamID	uint32
	httpResCode	int		// 0 means no ":status" line
	h		http.Header	// may be nil
	trailers	[]string	// if non-nil, which keys of h to write. nil means all.
	endStream	bool

	date		string
	contentType	string
	contentLength	string
}

func encKV(enc *hpack.Encoder, k, v string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:191
	_go_fuzz_dep_.CoverTab[75688]++
										if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:192
		_go_fuzz_dep_.CoverTab[75690]++
											log.Printf("http2: server encoding header %q = %q", k, v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:193
		// _ = "end of CoverTab[75690]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:194
		_go_fuzz_dep_.CoverTab[75691]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:194
		// _ = "end of CoverTab[75691]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:194
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:194
	// _ = "end of CoverTab[75688]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:194
	_go_fuzz_dep_.CoverTab[75689]++
										enc.WriteField(hpack.HeaderField{Name: k, Value: v})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:195
	// _ = "end of CoverTab[75689]"
}

func (w *writeResHeaders) staysWithinBuffer(max int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:198
	_go_fuzz_dep_.CoverTab[75692]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:206
	return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:206
	// _ = "end of CoverTab[75692]"
}

func (w *writeResHeaders) writeFrame(ctx writeContext) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:209
	_go_fuzz_dep_.CoverTab[75693]++
										enc, buf := ctx.HeaderEncoder()
										buf.Reset()

										if w.httpResCode != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:213
		_go_fuzz_dep_.CoverTab[75699]++
											encKV(enc, ":status", httpCodeString(w.httpResCode))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:214
		// _ = "end of CoverTab[75699]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:215
		_go_fuzz_dep_.CoverTab[75700]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:215
		// _ = "end of CoverTab[75700]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:215
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:215
	// _ = "end of CoverTab[75693]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:215
	_go_fuzz_dep_.CoverTab[75694]++

										encodeHeaders(enc, w.h, w.trailers)

										if w.contentType != "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:219
		_go_fuzz_dep_.CoverTab[75701]++
											encKV(enc, "content-type", w.contentType)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:220
		// _ = "end of CoverTab[75701]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:221
		_go_fuzz_dep_.CoverTab[75702]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:221
		// _ = "end of CoverTab[75702]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:221
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:221
	// _ = "end of CoverTab[75694]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:221
	_go_fuzz_dep_.CoverTab[75695]++
										if w.contentLength != "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:222
		_go_fuzz_dep_.CoverTab[75703]++
											encKV(enc, "content-length", w.contentLength)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:223
		// _ = "end of CoverTab[75703]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:224
		_go_fuzz_dep_.CoverTab[75704]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:224
		// _ = "end of CoverTab[75704]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:224
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:224
	// _ = "end of CoverTab[75695]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:224
	_go_fuzz_dep_.CoverTab[75696]++
										if w.date != "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:225
		_go_fuzz_dep_.CoverTab[75705]++
											encKV(enc, "date", w.date)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:226
		// _ = "end of CoverTab[75705]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:227
		_go_fuzz_dep_.CoverTab[75706]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:227
		// _ = "end of CoverTab[75706]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:227
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:227
	// _ = "end of CoverTab[75696]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:227
	_go_fuzz_dep_.CoverTab[75697]++

										headerBlock := buf.Bytes()
										if len(headerBlock) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:230
		_go_fuzz_dep_.CoverTab[75707]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:230
		return w.trailers == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:230
		// _ = "end of CoverTab[75707]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:230
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:230
		_go_fuzz_dep_.CoverTab[75708]++
											panic("unexpected empty hpack")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:231
		// _ = "end of CoverTab[75708]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:232
		_go_fuzz_dep_.CoverTab[75709]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:232
		// _ = "end of CoverTab[75709]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:232
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:232
	// _ = "end of CoverTab[75697]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:232
	_go_fuzz_dep_.CoverTab[75698]++

										return splitHeaderBlock(ctx, headerBlock, w.writeHeaderBlock)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:234
	// _ = "end of CoverTab[75698]"
}

func (w *writeResHeaders) writeHeaderBlock(ctx writeContext, frag []byte, firstFrag, lastFrag bool) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:237
	_go_fuzz_dep_.CoverTab[75710]++
										if firstFrag {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:238
		_go_fuzz_dep_.CoverTab[75711]++
											return ctx.Framer().WriteHeaders(HeadersFrameParam{
			StreamID:	w.streamID,
			BlockFragment:	frag,
			EndStream:	w.endStream,
			EndHeaders:	lastFrag,
		})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:244
		// _ = "end of CoverTab[75711]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:245
		_go_fuzz_dep_.CoverTab[75712]++
											return ctx.Framer().WriteContinuation(w.streamID, lastFrag, frag)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:246
		// _ = "end of CoverTab[75712]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:247
	// _ = "end of CoverTab[75710]"
}

// writePushPromise is a request to write a PUSH_PROMISE and 0+ CONTINUATION frames.
type writePushPromise struct {
	streamID	uint32		// pusher stream
	method		string		// for :method
	url		*url.URL	// for :scheme, :authority, :path
	h		http.Header

	// Creates an ID for a pushed stream. This runs on serveG just before
	// the frame is written. The returned ID is copied to promisedID.
	allocatePromisedID	func() (uint32, error)
	promisedID		uint32
}

func (w *writePushPromise) staysWithinBuffer(max int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:263
	_go_fuzz_dep_.CoverTab[75713]++

										return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:265
	// _ = "end of CoverTab[75713]"
}

func (w *writePushPromise) writeFrame(ctx writeContext) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:268
	_go_fuzz_dep_.CoverTab[75714]++
										enc, buf := ctx.HeaderEncoder()
										buf.Reset()

										encKV(enc, ":method", w.method)
										encKV(enc, ":scheme", w.url.Scheme)
										encKV(enc, ":authority", w.url.Host)
										encKV(enc, ":path", w.url.RequestURI())
										encodeHeaders(enc, w.h, nil)

										headerBlock := buf.Bytes()
										if len(headerBlock) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:279
		_go_fuzz_dep_.CoverTab[75716]++
											panic("unexpected empty hpack")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:280
		// _ = "end of CoverTab[75716]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:281
		_go_fuzz_dep_.CoverTab[75717]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:281
		// _ = "end of CoverTab[75717]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:281
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:281
	// _ = "end of CoverTab[75714]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:281
	_go_fuzz_dep_.CoverTab[75715]++

										return splitHeaderBlock(ctx, headerBlock, w.writeHeaderBlock)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:283
	// _ = "end of CoverTab[75715]"
}

func (w *writePushPromise) writeHeaderBlock(ctx writeContext, frag []byte, firstFrag, lastFrag bool) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:286
	_go_fuzz_dep_.CoverTab[75718]++
										if firstFrag {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:287
		_go_fuzz_dep_.CoverTab[75719]++
											return ctx.Framer().WritePushPromise(PushPromiseParam{
			StreamID:	w.streamID,
			PromiseID:	w.promisedID,
			BlockFragment:	frag,
			EndHeaders:	lastFrag,
		})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:293
		// _ = "end of CoverTab[75719]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:294
		_go_fuzz_dep_.CoverTab[75720]++
											return ctx.Framer().WriteContinuation(w.streamID, lastFrag, frag)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:295
		// _ = "end of CoverTab[75720]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:296
	// _ = "end of CoverTab[75718]"
}

type write100ContinueHeadersFrame struct {
	streamID uint32
}

func (w write100ContinueHeadersFrame) writeFrame(ctx writeContext) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:303
	_go_fuzz_dep_.CoverTab[75721]++
										enc, buf := ctx.HeaderEncoder()
										buf.Reset()
										encKV(enc, ":status", "100")
										return ctx.Framer().WriteHeaders(HeadersFrameParam{
		StreamID:	w.streamID,
		BlockFragment:	buf.Bytes(),
		EndStream:	false,
		EndHeaders:	true,
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:312
	// _ = "end of CoverTab[75721]"
}

func (w write100ContinueHeadersFrame) staysWithinBuffer(max int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:315
	_go_fuzz_dep_.CoverTab[75722]++

										return 9+2*(len(":status")+len("100")) <= max
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:317
	// _ = "end of CoverTab[75722]"
}

type writeWindowUpdate struct {
	streamID	uint32	// or 0 for conn-level
	n		uint32
}

func (wu writeWindowUpdate) staysWithinBuffer(max int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:325
	_go_fuzz_dep_.CoverTab[75723]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:325
	return frameHeaderLen+4 <= max
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:325
	// _ = "end of CoverTab[75723]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:325
}

func (wu writeWindowUpdate) writeFrame(ctx writeContext) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:327
	_go_fuzz_dep_.CoverTab[75724]++
										return ctx.Framer().WriteWindowUpdate(wu.streamID, wu.n)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:328
	// _ = "end of CoverTab[75724]"
}

// encodeHeaders encodes an http.Header. If keys is not nil, then (k, h[k])
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:331
// is encoded only if k is in keys.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:333
func encodeHeaders(enc *hpack.Encoder, h http.Header, keys []string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:333
	_go_fuzz_dep_.CoverTab[75725]++
										if keys == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:334
		_go_fuzz_dep_.CoverTab[75727]++
											sorter := sorterPool.Get().(*sorter)

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:339
		defer sorterPool.Put(sorter)
											keys = sorter.Keys(h)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:340
		// _ = "end of CoverTab[75727]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:341
		_go_fuzz_dep_.CoverTab[75728]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:341
		// _ = "end of CoverTab[75728]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:341
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:341
	// _ = "end of CoverTab[75725]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:341
	_go_fuzz_dep_.CoverTab[75726]++
										for _, k := range keys {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:342
		_go_fuzz_dep_.CoverTab[75729]++
											vv := h[k]
											k, ascii := lowerHeader(k)
											if !ascii {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:345
			_go_fuzz_dep_.CoverTab[75732]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:348
			continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:348
			// _ = "end of CoverTab[75732]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:349
			_go_fuzz_dep_.CoverTab[75733]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:349
			// _ = "end of CoverTab[75733]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:349
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:349
		// _ = "end of CoverTab[75729]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:349
		_go_fuzz_dep_.CoverTab[75730]++
											if !validWireHeaderFieldName(k) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:350
			_go_fuzz_dep_.CoverTab[75734]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:354
			continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:354
			// _ = "end of CoverTab[75734]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:355
			_go_fuzz_dep_.CoverTab[75735]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:355
			// _ = "end of CoverTab[75735]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:355
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:355
		// _ = "end of CoverTab[75730]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:355
		_go_fuzz_dep_.CoverTab[75731]++
											isTE := k == "transfer-encoding"
											for _, v := range vv {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:357
			_go_fuzz_dep_.CoverTab[75736]++
												if !httpguts.ValidHeaderFieldValue(v) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:358
				_go_fuzz_dep_.CoverTab[75739]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:361
				continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:361
				// _ = "end of CoverTab[75739]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:362
				_go_fuzz_dep_.CoverTab[75740]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:362
				// _ = "end of CoverTab[75740]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:362
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:362
			// _ = "end of CoverTab[75736]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:362
			_go_fuzz_dep_.CoverTab[75737]++

												if isTE && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:364
				_go_fuzz_dep_.CoverTab[75741]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:364
				return v != "trailers"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:364
				// _ = "end of CoverTab[75741]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:364
			}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:364
				_go_fuzz_dep_.CoverTab[75742]++
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:365
				// _ = "end of CoverTab[75742]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:366
				_go_fuzz_dep_.CoverTab[75743]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:366
				// _ = "end of CoverTab[75743]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:366
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:366
			// _ = "end of CoverTab[75737]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:366
			_go_fuzz_dep_.CoverTab[75738]++
												encKV(enc, k, v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:367
			// _ = "end of CoverTab[75738]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:368
		// _ = "end of CoverTab[75731]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:369
	// _ = "end of CoverTab[75726]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:370
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/write.go:370
var _ = _go_fuzz_dep_.CoverTab
