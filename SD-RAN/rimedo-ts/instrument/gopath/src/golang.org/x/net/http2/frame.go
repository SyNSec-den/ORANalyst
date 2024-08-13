// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:5
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:5
)

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"

	"golang.org/x/net/http/httpguts"
	"golang.org/x/net/http2/hpack"
)

const frameHeaderLen = 9

var padZeros = make([]byte, 255)	// zeros for padding

// A FrameType is a registered frame type as defined in
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:25
// https://httpwg.org/specs/rfc7540.html#rfc.section.11.2
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:27
type FrameType uint8

const (
	FrameData		FrameType	= 0x0
	FrameHeaders		FrameType	= 0x1
	FramePriority		FrameType	= 0x2
	FrameRSTStream		FrameType	= 0x3
	FrameSettings		FrameType	= 0x4
	FramePushPromise	FrameType	= 0x5
	FramePing		FrameType	= 0x6
	FrameGoAway		FrameType	= 0x7
	FrameWindowUpdate	FrameType	= 0x8
	FrameContinuation	FrameType	= 0x9
)

var frameName = map[FrameType]string{
	FrameData:		"DATA",
	FrameHeaders:		"HEADERS",
	FramePriority:		"PRIORITY",
	FrameRSTStream:		"RST_STREAM",
	FrameSettings:		"SETTINGS",
	FramePushPromise:	"PUSH_PROMISE",
	FramePing:		"PING",
	FrameGoAway:		"GOAWAY",
	FrameWindowUpdate:	"WINDOW_UPDATE",
	FrameContinuation:	"CONTINUATION",
}

func (t FrameType) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:55
	_go_fuzz_dep_.CoverTab[72437]++
										if s, ok := frameName[t]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:56
		_go_fuzz_dep_.CoverTab[72439]++
											return s
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:57
		// _ = "end of CoverTab[72439]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:58
		_go_fuzz_dep_.CoverTab[72440]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:58
		// _ = "end of CoverTab[72440]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:58
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:58
	// _ = "end of CoverTab[72437]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:58
	_go_fuzz_dep_.CoverTab[72438]++
										return fmt.Sprintf("UNKNOWN_FRAME_TYPE_%d", uint8(t))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:59
	// _ = "end of CoverTab[72438]"
}

// Flags is a bitmask of HTTP/2 flags.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:62
// The meaning of flags varies depending on the frame type.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:64
type Flags uint8

// Has reports whether f contains all (0 or more) flags in v.
func (f Flags) Has(v Flags) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:67
	_go_fuzz_dep_.CoverTab[72441]++
										return (f & v) == v
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:68
	// _ = "end of CoverTab[72441]"
}

// Frame-specific FrameHeader flag bits.
const (
	// Data Frame
	FlagDataEndStream	Flags	= 0x1
	FlagDataPadded		Flags	= 0x8

	// Headers Frame
	FlagHeadersEndStream	Flags	= 0x1
	FlagHeadersEndHeaders	Flags	= 0x4
	FlagHeadersPadded	Flags	= 0x8
	FlagHeadersPriority	Flags	= 0x20

	// Settings Frame
	FlagSettingsAck	Flags	= 0x1

	// Ping Frame
	FlagPingAck	Flags	= 0x1

	// Continuation Frame
	FlagContinuationEndHeaders	Flags	= 0x4

	FlagPushPromiseEndHeaders	Flags	= 0x4
	FlagPushPromisePadded		Flags	= 0x8
)

var flagName = map[FrameType]map[Flags]string{
	FrameData: {
		FlagDataEndStream:	"END_STREAM",
		FlagDataPadded:		"PADDED",
	},
	FrameHeaders: {
		FlagHeadersEndStream:	"END_STREAM",
		FlagHeadersEndHeaders:	"END_HEADERS",
		FlagHeadersPadded:	"PADDED",
		FlagHeadersPriority:	"PRIORITY",
	},
	FrameSettings: {
		FlagSettingsAck: "ACK",
	},
	FramePing: {
		FlagPingAck: "ACK",
	},
	FrameContinuation: {
		FlagContinuationEndHeaders: "END_HEADERS",
	},
	FramePushPromise: {
		FlagPushPromiseEndHeaders:	"END_HEADERS",
		FlagPushPromisePadded:		"PADDED",
	},
}

// a frameParser parses a frame given its FrameHeader and payload
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:122
// bytes. The length of payload will always equal fh.Length (which
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:122
// might be 0).
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:125
type frameParser func(fc *frameCache, fh FrameHeader, countError func(string), payload []byte) (Frame, error)

var frameParsers = map[FrameType]frameParser{
	FrameData:		parseDataFrame,
	FrameHeaders:		parseHeadersFrame,
	FramePriority:		parsePriorityFrame,
	FrameRSTStream:		parseRSTStreamFrame,
	FrameSettings:		parseSettingsFrame,
	FramePushPromise:	parsePushPromise,
	FramePing:		parsePingFrame,
	FrameGoAway:		parseGoAwayFrame,
	FrameWindowUpdate:	parseWindowUpdateFrame,
	FrameContinuation:	parseContinuationFrame,
}

func typeFrameParser(t FrameType) frameParser {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:140
	_go_fuzz_dep_.CoverTab[72442]++
										if f := frameParsers[t]; f != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:141
		_go_fuzz_dep_.CoverTab[72444]++
											return f
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:142
		// _ = "end of CoverTab[72444]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:143
		_go_fuzz_dep_.CoverTab[72445]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:143
		// _ = "end of CoverTab[72445]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:143
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:143
	// _ = "end of CoverTab[72442]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:143
	_go_fuzz_dep_.CoverTab[72443]++
										return parseUnknownFrame
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:144
	// _ = "end of CoverTab[72443]"
}

// A FrameHeader is the 9 byte header of all HTTP/2 frames.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:147
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:147
// See https://httpwg.org/specs/rfc7540.html#FrameHeader
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:150
type FrameHeader struct {
	valid	bool	// caller can access []byte fields in the Frame

	// Type is the 1 byte frame type. There are ten standard frame
	// types, but extension frame types may be written by WriteRawFrame
	// and will be returned by ReadFrame (as UnknownFrame).
	Type	FrameType

	// Flags are the 1 byte of 8 potential bit flags per frame.
	// They are specific to the frame type.
	Flags	Flags

	// Length is the length of the frame, not including the 9 byte header.
	// The maximum size is one byte less than 16MB (uint24), but only
	// frames up to 16KB are allowed without peer agreement.
	Length	uint32

	// StreamID is which stream this frame is for. Certain frames
	// are not stream-specific, in which case this field is 0.
	StreamID	uint32
}

// Header returns h. It exists so FrameHeaders can be embedded in other
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:172
// specific frame types and implement the Frame interface.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:174
func (h FrameHeader) Header() FrameHeader {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:174
	_go_fuzz_dep_.CoverTab[72446]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:174
	return h
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:174
	// _ = "end of CoverTab[72446]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:174
}

func (h FrameHeader) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:176
	_go_fuzz_dep_.CoverTab[72447]++
										var buf bytes.Buffer
										buf.WriteString("[FrameHeader ")
										h.writeDebug(&buf)
										buf.WriteByte(']')
										return buf.String()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:181
	// _ = "end of CoverTab[72447]"
}

func (h FrameHeader) writeDebug(buf *bytes.Buffer) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:184
	_go_fuzz_dep_.CoverTab[72448]++
										buf.WriteString(h.Type.String())
										if h.Flags != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:186
		_go_fuzz_dep_.CoverTab[72451]++
											buf.WriteString(" flags=")
											set := 0
											for i := uint8(0); i < 8; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:189
			_go_fuzz_dep_.CoverTab[72452]++
												if h.Flags&(1<<i) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:190
				_go_fuzz_dep_.CoverTab[72455]++
													continue
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:191
				// _ = "end of CoverTab[72455]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:192
				_go_fuzz_dep_.CoverTab[72456]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:192
				// _ = "end of CoverTab[72456]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:192
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:192
			// _ = "end of CoverTab[72452]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:192
			_go_fuzz_dep_.CoverTab[72453]++
												set++
												if set > 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:194
				_go_fuzz_dep_.CoverTab[72457]++
													buf.WriteByte('|')
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:195
				// _ = "end of CoverTab[72457]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:196
				_go_fuzz_dep_.CoverTab[72458]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:196
				// _ = "end of CoverTab[72458]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:196
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:196
			// _ = "end of CoverTab[72453]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:196
			_go_fuzz_dep_.CoverTab[72454]++
												name := flagName[h.Type][Flags(1<<i)]
												if name != "" {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:198
				_go_fuzz_dep_.CoverTab[72459]++
													buf.WriteString(name)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:199
				// _ = "end of CoverTab[72459]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:200
				_go_fuzz_dep_.CoverTab[72460]++
													fmt.Fprintf(buf, "0x%x", 1<<i)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:201
				// _ = "end of CoverTab[72460]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:202
			// _ = "end of CoverTab[72454]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:203
		// _ = "end of CoverTab[72451]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:204
		_go_fuzz_dep_.CoverTab[72461]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:204
		// _ = "end of CoverTab[72461]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:204
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:204
	// _ = "end of CoverTab[72448]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:204
	_go_fuzz_dep_.CoverTab[72449]++
										if h.StreamID != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:205
		_go_fuzz_dep_.CoverTab[72462]++
											fmt.Fprintf(buf, " stream=%d", h.StreamID)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:206
		// _ = "end of CoverTab[72462]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:207
		_go_fuzz_dep_.CoverTab[72463]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:207
		// _ = "end of CoverTab[72463]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:207
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:207
	// _ = "end of CoverTab[72449]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:207
	_go_fuzz_dep_.CoverTab[72450]++
										fmt.Fprintf(buf, " len=%d", h.Length)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:208
	// _ = "end of CoverTab[72450]"
}

func (h *FrameHeader) checkValid() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:211
	_go_fuzz_dep_.CoverTab[72464]++
										if !h.valid {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:212
		_go_fuzz_dep_.CoverTab[72465]++
											panic("Frame accessor called on non-owned Frame")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:213
		// _ = "end of CoverTab[72465]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:214
		_go_fuzz_dep_.CoverTab[72466]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:214
		// _ = "end of CoverTab[72466]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:214
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:214
	// _ = "end of CoverTab[72464]"
}

func (h *FrameHeader) invalidate() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:217
	_go_fuzz_dep_.CoverTab[72467]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:217
	h.valid = false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:217
	// _ = "end of CoverTab[72467]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:217
}

// frame header bytes.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:219
// Used only by ReadFrameHeader.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:221
var fhBytes = sync.Pool{
	New: func() interface{} {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:222
		_go_fuzz_dep_.CoverTab[72468]++
											buf := make([]byte, frameHeaderLen)
											return &buf
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:224
		// _ = "end of CoverTab[72468]"
	},
}

// ReadFrameHeader reads 9 bytes from r and returns a FrameHeader.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:228
// Most users should use Framer.ReadFrame instead.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:230
func ReadFrameHeader(r io.Reader) (FrameHeader, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:230
	_go_fuzz_dep_.CoverTab[72469]++
										bufp := fhBytes.Get().(*[]byte)
										defer fhBytes.Put(bufp)
										return readFrameHeader(*bufp, r)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:233
	// _ = "end of CoverTab[72469]"
}

func readFrameHeader(buf []byte, r io.Reader) (FrameHeader, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:236
	_go_fuzz_dep_.CoverTab[72470]++
										_, err := io.ReadFull(r, buf[:frameHeaderLen])
										if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:238
		_go_fuzz_dep_.CoverTab[72472]++
											return FrameHeader{}, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:239
		// _ = "end of CoverTab[72472]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:240
		_go_fuzz_dep_.CoverTab[72473]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:240
		// _ = "end of CoverTab[72473]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:240
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:240
	// _ = "end of CoverTab[72470]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:240
	_go_fuzz_dep_.CoverTab[72471]++
										return FrameHeader{
		Length:		(uint32(buf[0])<<16 | uint32(buf[1])<<8 | uint32(buf[2])),
		Type:		FrameType(buf[3]),
		Flags:		Flags(buf[4]),
		StreamID:	binary.BigEndian.Uint32(buf[5:]) & (1<<31 - 1),
		valid:		true,
	}, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:247
	// _ = "end of CoverTab[72471]"
}

// A Frame is the base interface implemented by all frame types.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:250
// Callers will generally type-assert the specific frame type:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:250
// *HeadersFrame, *SettingsFrame, *WindowUpdateFrame, etc.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:250
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:250
// Frames are only valid until the next call to Framer.ReadFrame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:255
type Frame interface {
	Header() FrameHeader

	// invalidate is called by Framer.ReadFrame to make this
	// frame's buffers as being invalid, since the subsequent
	// frame will reuse them.
	invalidate()
}

// A Framer reads and writes Frames.
type Framer struct {
	r		io.Reader
	lastFrame	Frame
	errDetail	error

	// countError is a non-nil func that's called on a frame parse
	// error with some unique error path token. It's initialized
	// from Transport.CountError or Server.CountError.
	countError	func(errToken string)

	// lastHeaderStream is non-zero if the last frame was an
	// unfinished HEADERS/CONTINUATION.
	lastHeaderStream	uint32

	maxReadSize	uint32
	headerBuf	[frameHeaderLen]byte

	// TODO: let getReadBuf be configurable, and use a less memory-pinning
	// allocator in server.go to minimize memory pinned for many idle conns.
	// Will probably also need to make frame invalidation have a hook too.
	getReadBuf	func(size uint32) []byte
	readBuf		[]byte	// cache for default getReadBuf

	maxWriteSize	uint32	// zero means unlimited; TODO: implement

	w	io.Writer
	wbuf	[]byte

	// AllowIllegalWrites permits the Framer's Write methods to
	// write frames that do not conform to the HTTP/2 spec. This
	// permits using the Framer to test other HTTP/2
	// implementations' conformance to the spec.
	// If false, the Write methods will prefer to return an error
	// rather than comply.
	AllowIllegalWrites	bool

	// AllowIllegalReads permits the Framer's ReadFrame method
	// to return non-compliant frames or frame orders.
	// This is for testing and permits using the Framer to test
	// other HTTP/2 implementations' conformance to the spec.
	// It is not compatible with ReadMetaHeaders.
	AllowIllegalReads	bool

	// ReadMetaHeaders if non-nil causes ReadFrame to merge
	// HEADERS and CONTINUATION frames together and return
	// MetaHeadersFrame instead.
	ReadMetaHeaders	*hpack.Decoder

	// MaxHeaderListSize is the http2 MAX_HEADER_LIST_SIZE.
	// It's used only if ReadMetaHeaders is set; 0 means a sane default
	// (currently 16MB)
										// If the limit is hit, MetaHeadersFrame.Truncated is set true.
										MaxHeaderListSize	uint32

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:325
	logReads, logWrites	bool

	debugFramer		*Framer	// only use for logging written writes
	debugFramerBuf		*bytes.Buffer
	debugReadLoggerf	func(string, ...interface{})
	debugWriteLoggerf	func(string, ...interface{})

	frameCache	*frameCache	// nil if frames aren't reused (default)
}

func (fr *Framer) maxHeaderListSize() uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:335
	_go_fuzz_dep_.CoverTab[72474]++
										if fr.MaxHeaderListSize == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:336
		_go_fuzz_dep_.CoverTab[72476]++
											return 16 << 20
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:337
		// _ = "end of CoverTab[72476]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:338
		_go_fuzz_dep_.CoverTab[72477]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:338
		// _ = "end of CoverTab[72477]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:338
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:338
	// _ = "end of CoverTab[72474]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:338
	_go_fuzz_dep_.CoverTab[72475]++
										return fr.MaxHeaderListSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:339
	// _ = "end of CoverTab[72475]"
}

func (f *Framer) startWrite(ftype FrameType, flags Flags, streamID uint32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:342
	_go_fuzz_dep_.CoverTab[72478]++

										f.wbuf = append(f.wbuf[:0],
		0,
		0,
		0,
		byte(ftype),
		byte(flags),
		byte(streamID>>24),
		byte(streamID>>16),
		byte(streamID>>8),
		byte(streamID))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:353
	// _ = "end of CoverTab[72478]"
}

func (f *Framer) endWrite() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:356
	_go_fuzz_dep_.CoverTab[72479]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:359
	length := len(f.wbuf) - frameHeaderLen
	if length >= (1 << 24) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:360
		_go_fuzz_dep_.CoverTab[72483]++
											return ErrFrameTooLarge
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:361
		// _ = "end of CoverTab[72483]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:362
		_go_fuzz_dep_.CoverTab[72484]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:362
		// _ = "end of CoverTab[72484]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:362
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:362
	// _ = "end of CoverTab[72479]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:362
	_go_fuzz_dep_.CoverTab[72480]++
										_ = append(f.wbuf[:0],
		byte(length>>16),
		byte(length>>8),
		byte(length))
	if f.logWrites {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:367
		_go_fuzz_dep_.CoverTab[72485]++
											f.logWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:368
		// _ = "end of CoverTab[72485]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:369
		_go_fuzz_dep_.CoverTab[72486]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:369
		// _ = "end of CoverTab[72486]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:369
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:369
	// _ = "end of CoverTab[72480]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:369
	_go_fuzz_dep_.CoverTab[72481]++

										n, err := f.w.Write(f.wbuf)
										if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:372
		_go_fuzz_dep_.CoverTab[72487]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:372
		return n != len(f.wbuf)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:372
		// _ = "end of CoverTab[72487]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:372
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:372
		_go_fuzz_dep_.CoverTab[72488]++
											err = io.ErrShortWrite
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:373
		// _ = "end of CoverTab[72488]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:374
		_go_fuzz_dep_.CoverTab[72489]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:374
		// _ = "end of CoverTab[72489]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:374
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:374
	// _ = "end of CoverTab[72481]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:374
	_go_fuzz_dep_.CoverTab[72482]++
										return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:375
	// _ = "end of CoverTab[72482]"
}

func (f *Framer) logWrite() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:378
	_go_fuzz_dep_.CoverTab[72490]++
										if f.debugFramer == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:379
		_go_fuzz_dep_.CoverTab[72493]++
											f.debugFramerBuf = new(bytes.Buffer)
											f.debugFramer = NewFramer(nil, f.debugFramerBuf)
											f.debugFramer.logReads = false

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:385
		f.debugFramer.AllowIllegalReads = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:385
		// _ = "end of CoverTab[72493]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:386
		_go_fuzz_dep_.CoverTab[72494]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:386
		// _ = "end of CoverTab[72494]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:386
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:386
	// _ = "end of CoverTab[72490]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:386
	_go_fuzz_dep_.CoverTab[72491]++
										f.debugFramerBuf.Write(f.wbuf)
										fr, err := f.debugFramer.ReadFrame()
										if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:389
		_go_fuzz_dep_.CoverTab[72495]++
											f.debugWriteLoggerf("http2: Framer %p: failed to decode just-written frame", f)
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:391
		// _ = "end of CoverTab[72495]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:392
		_go_fuzz_dep_.CoverTab[72496]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:392
		// _ = "end of CoverTab[72496]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:392
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:392
	// _ = "end of CoverTab[72491]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:392
	_go_fuzz_dep_.CoverTab[72492]++
										f.debugWriteLoggerf("http2: Framer %p: wrote %v", f, summarizeFrame(fr))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:393
	// _ = "end of CoverTab[72492]"
}

func (f *Framer) writeByte(v byte) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:396
	_go_fuzz_dep_.CoverTab[72497]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:396
	f.wbuf = append(f.wbuf, v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:396
	// _ = "end of CoverTab[72497]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:396
}
func (f *Framer) writeBytes(v []byte) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:397
	_go_fuzz_dep_.CoverTab[72498]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:397
	f.wbuf = append(f.wbuf, v...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:397
	// _ = "end of CoverTab[72498]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:397
}
func (f *Framer) writeUint16(v uint16) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:398
	_go_fuzz_dep_.CoverTab[72499]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:398
	f.wbuf = append(f.wbuf, byte(v>>8), byte(v))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:398
	// _ = "end of CoverTab[72499]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:398
}
func (f *Framer) writeUint32(v uint32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:399
	_go_fuzz_dep_.CoverTab[72500]++
										f.wbuf = append(f.wbuf, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:400
	// _ = "end of CoverTab[72500]"
}

const (
	minMaxFrameSize	= 1 << 14
	maxFrameSize	= 1<<24 - 1
)

// SetReuseFrames allows the Framer to reuse Frames.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:408
// If called on a Framer, Frames returned by calls to ReadFrame are only
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:408
// valid until the next call to ReadFrame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:411
func (fr *Framer) SetReuseFrames() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:411
	_go_fuzz_dep_.CoverTab[72501]++
										if fr.frameCache != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:412
		_go_fuzz_dep_.CoverTab[72503]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:413
		// _ = "end of CoverTab[72503]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:414
		_go_fuzz_dep_.CoverTab[72504]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:414
		// _ = "end of CoverTab[72504]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:414
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:414
	// _ = "end of CoverTab[72501]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:414
	_go_fuzz_dep_.CoverTab[72502]++
										fr.frameCache = &frameCache{}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:415
	// _ = "end of CoverTab[72502]"
}

type frameCache struct {
	dataFrame DataFrame
}

func (fc *frameCache) getDataFrame() *DataFrame {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:422
	_go_fuzz_dep_.CoverTab[72505]++
										if fc == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:423
		_go_fuzz_dep_.CoverTab[72507]++
											return &DataFrame{}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:424
		// _ = "end of CoverTab[72507]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:425
		_go_fuzz_dep_.CoverTab[72508]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:425
		// _ = "end of CoverTab[72508]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:425
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:425
	// _ = "end of CoverTab[72505]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:425
	_go_fuzz_dep_.CoverTab[72506]++
										return &fc.dataFrame
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:426
	// _ = "end of CoverTab[72506]"
}

// NewFramer returns a Framer that writes frames to w and reads them from r.
func NewFramer(w io.Writer, r io.Reader) *Framer {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:430
	_go_fuzz_dep_.CoverTab[72509]++
										fr := &Framer{
		w:			w,
		r:			r,
		countError:		func(string) { _go_fuzz_dep_.CoverTab[72512]++; // _ = "end of CoverTab[72512]" },
		logReads:		logFrameReads,
		logWrites:		logFrameWrites,
		debugReadLoggerf:	log.Printf,
		debugWriteLoggerf:	log.Printf,
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:439
	// _ = "end of CoverTab[72509]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:439
	_go_fuzz_dep_.CoverTab[72510]++
										fr.getReadBuf = func(size uint32) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:440
		_go_fuzz_dep_.CoverTab[72513]++
											if cap(fr.readBuf) >= int(size) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:441
			_go_fuzz_dep_.CoverTab[72515]++
												return fr.readBuf[:size]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:442
			// _ = "end of CoverTab[72515]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:443
			_go_fuzz_dep_.CoverTab[72516]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:443
			// _ = "end of CoverTab[72516]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:443
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:443
		// _ = "end of CoverTab[72513]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:443
		_go_fuzz_dep_.CoverTab[72514]++
											fr.readBuf = make([]byte, size)
											return fr.readBuf
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:445
		// _ = "end of CoverTab[72514]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:446
	// _ = "end of CoverTab[72510]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:446
	_go_fuzz_dep_.CoverTab[72511]++
										fr.SetMaxReadFrameSize(maxFrameSize)
										return fr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:448
	// _ = "end of CoverTab[72511]"
}

// SetMaxReadFrameSize sets the maximum size of a frame
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:451
// that will be read by a subsequent call to ReadFrame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:451
// It is the caller's responsibility to advertise this
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:451
// limit with a SETTINGS frame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:455
func (fr *Framer) SetMaxReadFrameSize(v uint32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:455
	_go_fuzz_dep_.CoverTab[72517]++
										if v > maxFrameSize {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:456
		_go_fuzz_dep_.CoverTab[72519]++
											v = maxFrameSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:457
		// _ = "end of CoverTab[72519]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:458
		_go_fuzz_dep_.CoverTab[72520]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:458
		// _ = "end of CoverTab[72520]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:458
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:458
	// _ = "end of CoverTab[72517]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:458
	_go_fuzz_dep_.CoverTab[72518]++
										fr.maxReadSize = v
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:459
	// _ = "end of CoverTab[72518]"
}

// ErrorDetail returns a more detailed error of the last error
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:462
// returned by Framer.ReadFrame. For instance, if ReadFrame
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:462
// returns a StreamError with code PROTOCOL_ERROR, ErrorDetail
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:462
// will say exactly what was invalid. ErrorDetail is not guaranteed
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:462
// to return a non-nil value and like the rest of the http2 package,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:462
// its return value is not protected by an API compatibility promise.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:462
// ErrorDetail is reset after the next call to ReadFrame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:469
func (fr *Framer) ErrorDetail() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:469
	_go_fuzz_dep_.CoverTab[72521]++
										return fr.errDetail
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:470
	// _ = "end of CoverTab[72521]"
}

// ErrFrameTooLarge is returned from Framer.ReadFrame when the peer
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:473
// sends a frame that is larger than declared with SetMaxReadFrameSize.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:475
var ErrFrameTooLarge = errors.New("http2: frame too large")

// terminalReadFrameError reports whether err is an unrecoverable
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:477
// error from ReadFrame and no other frames should be read.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:479
func terminalReadFrameError(err error) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:479
	_go_fuzz_dep_.CoverTab[72522]++
										if _, ok := err.(StreamError); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:480
		_go_fuzz_dep_.CoverTab[72524]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:481
		// _ = "end of CoverTab[72524]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:482
		_go_fuzz_dep_.CoverTab[72525]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:482
		// _ = "end of CoverTab[72525]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:482
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:482
	// _ = "end of CoverTab[72522]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:482
	_go_fuzz_dep_.CoverTab[72523]++
										return err != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:483
	// _ = "end of CoverTab[72523]"
}

// ReadFrame reads a single frame. The returned Frame is only valid
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:486
// until the next call to ReadFrame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:486
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:486
// If the frame is larger than previously set with SetMaxReadFrameSize, the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:486
// returned error is ErrFrameTooLarge. Other errors may be of type
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:486
// ConnectionError, StreamError, or anything else from the underlying
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:486
// reader.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:493
func (fr *Framer) ReadFrame() (Frame, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:493
	_go_fuzz_dep_.CoverTab[72526]++
										fr.errDetail = nil
										if fr.lastFrame != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:495
		_go_fuzz_dep_.CoverTab[72535]++
											fr.lastFrame.invalidate()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:496
		// _ = "end of CoverTab[72535]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:497
		_go_fuzz_dep_.CoverTab[72536]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:497
		// _ = "end of CoverTab[72536]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:497
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:497
	// _ = "end of CoverTab[72526]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:497
	_go_fuzz_dep_.CoverTab[72527]++
										fh, err := readFrameHeader(fr.headerBuf[:], fr.r)
										if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:499
		_go_fuzz_dep_.CoverTab[72537]++
											return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:500
		// _ = "end of CoverTab[72537]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:501
		_go_fuzz_dep_.CoverTab[72538]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:501
		// _ = "end of CoverTab[72538]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:501
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:501
	// _ = "end of CoverTab[72527]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:501
	_go_fuzz_dep_.CoverTab[72528]++
										if fh.Length > fr.maxReadSize {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:502
		_go_fuzz_dep_.CoverTab[72539]++
											return nil, ErrFrameTooLarge
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:503
		// _ = "end of CoverTab[72539]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:504
		_go_fuzz_dep_.CoverTab[72540]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:504
		// _ = "end of CoverTab[72540]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:504
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:504
	// _ = "end of CoverTab[72528]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:504
	_go_fuzz_dep_.CoverTab[72529]++
										payload := fr.getReadBuf(fh.Length)
										if _, err := io.ReadFull(fr.r, payload); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:506
		_go_fuzz_dep_.CoverTab[72541]++
											return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:507
		// _ = "end of CoverTab[72541]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:508
		_go_fuzz_dep_.CoverTab[72542]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:508
		// _ = "end of CoverTab[72542]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:508
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:508
	// _ = "end of CoverTab[72529]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:508
	_go_fuzz_dep_.CoverTab[72530]++
										f, err := typeFrameParser(fh.Type)(fr.frameCache, fh, fr.countError, payload)
										if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:510
		_go_fuzz_dep_.CoverTab[72543]++
											if ce, ok := err.(connError); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:511
			_go_fuzz_dep_.CoverTab[72545]++
												return nil, fr.connError(ce.Code, ce.Reason)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:512
			// _ = "end of CoverTab[72545]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:513
			_go_fuzz_dep_.CoverTab[72546]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:513
			// _ = "end of CoverTab[72546]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:513
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:513
		// _ = "end of CoverTab[72543]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:513
		_go_fuzz_dep_.CoverTab[72544]++
											return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:514
		// _ = "end of CoverTab[72544]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:515
		_go_fuzz_dep_.CoverTab[72547]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:515
		// _ = "end of CoverTab[72547]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:515
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:515
	// _ = "end of CoverTab[72530]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:515
	_go_fuzz_dep_.CoverTab[72531]++
										if err := fr.checkFrameOrder(f); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:516
		_go_fuzz_dep_.CoverTab[72548]++
											return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:517
		// _ = "end of CoverTab[72548]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:518
		_go_fuzz_dep_.CoverTab[72549]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:518
		// _ = "end of CoverTab[72549]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:518
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:518
	// _ = "end of CoverTab[72531]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:518
	_go_fuzz_dep_.CoverTab[72532]++
										if fr.logReads {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:519
		_go_fuzz_dep_.CoverTab[72550]++
											fr.debugReadLoggerf("http2: Framer %p: read %v", fr, summarizeFrame(f))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:520
		// _ = "end of CoverTab[72550]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:521
		_go_fuzz_dep_.CoverTab[72551]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:521
		// _ = "end of CoverTab[72551]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:521
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:521
	// _ = "end of CoverTab[72532]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:521
	_go_fuzz_dep_.CoverTab[72533]++
										if fh.Type == FrameHeaders && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:522
		_go_fuzz_dep_.CoverTab[72552]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:522
		return fr.ReadMetaHeaders != nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:522
		// _ = "end of CoverTab[72552]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:522
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:522
		_go_fuzz_dep_.CoverTab[72553]++
											return fr.readMetaFrame(f.(*HeadersFrame))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:523
		// _ = "end of CoverTab[72553]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:524
		_go_fuzz_dep_.CoverTab[72554]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:524
		// _ = "end of CoverTab[72554]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:524
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:524
	// _ = "end of CoverTab[72533]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:524
	_go_fuzz_dep_.CoverTab[72534]++
										return f, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:525
	// _ = "end of CoverTab[72534]"
}

// connError returns ConnectionError(code) but first
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:528
// stashes away a public reason to the caller can optionally relay it
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:528
// to the peer before hanging up on them. This might help others debug
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:528
// their implementations.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:532
func (fr *Framer) connError(code ErrCode, reason string) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:532
	_go_fuzz_dep_.CoverTab[72555]++
										fr.errDetail = errors.New(reason)
										return ConnectionError(code)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:534
	// _ = "end of CoverTab[72555]"
}

// checkFrameOrder reports an error if f is an invalid frame to return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:537
// next from ReadFrame. Mostly it checks whether HEADERS and
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:537
// CONTINUATION frames are contiguous.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:540
func (fr *Framer) checkFrameOrder(f Frame) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:540
	_go_fuzz_dep_.CoverTab[72556]++
										last := fr.lastFrame
										fr.lastFrame = f
										if fr.AllowIllegalReads {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:543
		_go_fuzz_dep_.CoverTab[72560]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:544
		// _ = "end of CoverTab[72560]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:545
		_go_fuzz_dep_.CoverTab[72561]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:545
		// _ = "end of CoverTab[72561]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:545
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:545
	// _ = "end of CoverTab[72556]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:545
	_go_fuzz_dep_.CoverTab[72557]++

										fh := f.Header()
										if fr.lastHeaderStream != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:548
		_go_fuzz_dep_.CoverTab[72562]++
											if fh.Type != FrameContinuation {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:549
			_go_fuzz_dep_.CoverTab[72564]++
												return fr.connError(ErrCodeProtocol,
				fmt.Sprintf("got %s for stream %d; expected CONTINUATION following %s for stream %d",
					fh.Type, fh.StreamID,
					last.Header().Type, fr.lastHeaderStream))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:553
			// _ = "end of CoverTab[72564]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:554
			_go_fuzz_dep_.CoverTab[72565]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:554
			// _ = "end of CoverTab[72565]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:554
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:554
		// _ = "end of CoverTab[72562]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:554
		_go_fuzz_dep_.CoverTab[72563]++
											if fh.StreamID != fr.lastHeaderStream {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:555
			_go_fuzz_dep_.CoverTab[72566]++
												return fr.connError(ErrCodeProtocol,
				fmt.Sprintf("got CONTINUATION for stream %d; expected stream %d",
					fh.StreamID, fr.lastHeaderStream))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:558
			// _ = "end of CoverTab[72566]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:559
			_go_fuzz_dep_.CoverTab[72567]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:559
			// _ = "end of CoverTab[72567]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:559
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:559
		// _ = "end of CoverTab[72563]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:560
		_go_fuzz_dep_.CoverTab[72568]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:560
		if fh.Type == FrameContinuation {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:560
			_go_fuzz_dep_.CoverTab[72569]++
												return fr.connError(ErrCodeProtocol, fmt.Sprintf("unexpected CONTINUATION for stream %d", fh.StreamID))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:561
			// _ = "end of CoverTab[72569]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:562
			_go_fuzz_dep_.CoverTab[72570]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:562
			// _ = "end of CoverTab[72570]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:562
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:562
		// _ = "end of CoverTab[72568]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:562
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:562
	// _ = "end of CoverTab[72557]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:562
	_go_fuzz_dep_.CoverTab[72558]++

										switch fh.Type {
	case FrameHeaders, FrameContinuation:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:565
		_go_fuzz_dep_.CoverTab[72571]++
											if fh.Flags.Has(FlagHeadersEndHeaders) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:566
			_go_fuzz_dep_.CoverTab[72573]++
												fr.lastHeaderStream = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:567
			// _ = "end of CoverTab[72573]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:568
			_go_fuzz_dep_.CoverTab[72574]++
												fr.lastHeaderStream = fh.StreamID
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:569
			// _ = "end of CoverTab[72574]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:570
		// _ = "end of CoverTab[72571]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:570
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:570
		_go_fuzz_dep_.CoverTab[72572]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:570
		// _ = "end of CoverTab[72572]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:571
	// _ = "end of CoverTab[72558]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:571
	_go_fuzz_dep_.CoverTab[72559]++

										return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:573
	// _ = "end of CoverTab[72559]"
}

// A DataFrame conveys arbitrary, variable-length sequences of octets
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:576
// associated with a stream.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:576
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:579
type DataFrame struct {
	FrameHeader
	data	[]byte
}

func (f *DataFrame) StreamEnded() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:584
	_go_fuzz_dep_.CoverTab[72575]++
										return f.FrameHeader.Flags.Has(FlagDataEndStream)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:585
	// _ = "end of CoverTab[72575]"
}

// Data returns the frame's data octets, not including any padding
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:588
// size byte or padding suffix bytes.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:588
// The caller must not retain the returned memory past the next
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:588
// call to ReadFrame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:592
func (f *DataFrame) Data() []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:592
	_go_fuzz_dep_.CoverTab[72576]++
										f.checkValid()
										return f.data
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:594
	// _ = "end of CoverTab[72576]"
}

func parseDataFrame(fc *frameCache, fh FrameHeader, countError func(string), payload []byte) (Frame, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:597
	_go_fuzz_dep_.CoverTab[72577]++
										if fh.StreamID == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:598
		_go_fuzz_dep_.CoverTab[72581]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:604
		countError("frame_data_stream_0")
											return nil, connError{ErrCodeProtocol, "DATA frame with stream ID 0"}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:605
		// _ = "end of CoverTab[72581]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:606
		_go_fuzz_dep_.CoverTab[72582]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:606
		// _ = "end of CoverTab[72582]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:606
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:606
	// _ = "end of CoverTab[72577]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:606
	_go_fuzz_dep_.CoverTab[72578]++
										f := fc.getDataFrame()
										f.FrameHeader = fh

										var padSize byte
										if fh.Flags.Has(FlagDataPadded) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:611
		_go_fuzz_dep_.CoverTab[72583]++
											var err error
											payload, padSize, err = readByte(payload)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:614
			_go_fuzz_dep_.CoverTab[72584]++
												countError("frame_data_pad_byte_short")
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:616
			// _ = "end of CoverTab[72584]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:617
			_go_fuzz_dep_.CoverTab[72585]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:617
			// _ = "end of CoverTab[72585]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:617
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:617
		// _ = "end of CoverTab[72583]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:618
		_go_fuzz_dep_.CoverTab[72586]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:618
		// _ = "end of CoverTab[72586]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:618
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:618
	// _ = "end of CoverTab[72578]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:618
	_go_fuzz_dep_.CoverTab[72579]++
										if int(padSize) > len(payload) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:619
		_go_fuzz_dep_.CoverTab[72587]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:624
		countError("frame_data_pad_too_big")
											return nil, connError{ErrCodeProtocol, "pad size larger than data payload"}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:625
		// _ = "end of CoverTab[72587]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:626
		_go_fuzz_dep_.CoverTab[72588]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:626
		// _ = "end of CoverTab[72588]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:626
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:626
	// _ = "end of CoverTab[72579]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:626
	_go_fuzz_dep_.CoverTab[72580]++
										f.data = payload[:len(payload)-int(padSize)]
										return f, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:628
	// _ = "end of CoverTab[72580]"
}

var (
	errStreamID	= errors.New("invalid stream ID")
	errDepStreamID	= errors.New("invalid dependent stream ID")
	errPadLength	= errors.New("pad length too large")
	errPadBytes	= errors.New("padding bytes must all be zeros unless AllowIllegalWrites is enabled")
)

func validStreamIDOrZero(streamID uint32) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:638
	_go_fuzz_dep_.CoverTab[72589]++
										return streamID&(1<<31) == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:639
	// _ = "end of CoverTab[72589]"
}

func validStreamID(streamID uint32) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:642
	_go_fuzz_dep_.CoverTab[72590]++
										return streamID != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:643
		_go_fuzz_dep_.CoverTab[72591]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:643
		return streamID&(1<<31) == 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:643
		// _ = "end of CoverTab[72591]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:643
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:643
	// _ = "end of CoverTab[72590]"
}

// WriteData writes a DATA frame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:646
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:646
// It will perform exactly one Write to the underlying Writer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:646
// It is the caller's responsibility not to violate the maximum frame size
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:646
// and to not call other Write methods concurrently.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:651
func (f *Framer) WriteData(streamID uint32, endStream bool, data []byte) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:651
	_go_fuzz_dep_.CoverTab[72592]++
										return f.WriteDataPadded(streamID, endStream, data, nil)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:652
	// _ = "end of CoverTab[72592]"
}

// WriteDataPadded writes a DATA frame with optional padding.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:655
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:655
// If pad is nil, the padding bit is not sent.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:655
// The length of pad must not exceed 255 bytes.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:655
// The bytes of pad must all be zero, unless f.AllowIllegalWrites is set.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:655
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:655
// It will perform exactly one Write to the underlying Writer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:655
// It is the caller's responsibility not to violate the maximum frame size
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:655
// and to not call other Write methods concurrently.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:664
func (f *Framer) WriteDataPadded(streamID uint32, endStream bool, data, pad []byte) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:664
	_go_fuzz_dep_.CoverTab[72593]++
										if err := f.startWriteDataPadded(streamID, endStream, data, pad); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:665
		_go_fuzz_dep_.CoverTab[72595]++
											return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:666
		// _ = "end of CoverTab[72595]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:667
		_go_fuzz_dep_.CoverTab[72596]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:667
		// _ = "end of CoverTab[72596]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:667
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:667
	// _ = "end of CoverTab[72593]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:667
	_go_fuzz_dep_.CoverTab[72594]++
										return f.endWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:668
	// _ = "end of CoverTab[72594]"
}

// startWriteDataPadded is WriteDataPadded, but only writes the frame to the Framer's internal buffer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:671
// The caller should call endWrite to flush the frame to the underlying writer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:673
func (f *Framer) startWriteDataPadded(streamID uint32, endStream bool, data, pad []byte) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:673
	_go_fuzz_dep_.CoverTab[72597]++
										if !validStreamID(streamID) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:674
		_go_fuzz_dep_.CoverTab[72603]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:674
		return !f.AllowIllegalWrites
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:674
		// _ = "end of CoverTab[72603]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:674
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:674
		_go_fuzz_dep_.CoverTab[72604]++
											return errStreamID
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:675
		// _ = "end of CoverTab[72604]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:676
		_go_fuzz_dep_.CoverTab[72605]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:676
		// _ = "end of CoverTab[72605]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:676
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:676
	// _ = "end of CoverTab[72597]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:676
	_go_fuzz_dep_.CoverTab[72598]++
										if len(pad) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:677
		_go_fuzz_dep_.CoverTab[72606]++
											if len(pad) > 255 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:678
			_go_fuzz_dep_.CoverTab[72608]++
												return errPadLength
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:679
			// _ = "end of CoverTab[72608]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:680
			_go_fuzz_dep_.CoverTab[72609]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:680
			// _ = "end of CoverTab[72609]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:680
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:680
		// _ = "end of CoverTab[72606]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:680
		_go_fuzz_dep_.CoverTab[72607]++
											if !f.AllowIllegalWrites {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:681
			_go_fuzz_dep_.CoverTab[72610]++
												for _, b := range pad {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:682
				_go_fuzz_dep_.CoverTab[72611]++
													if b != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:683
					_go_fuzz_dep_.CoverTab[72612]++

														return errPadBytes
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:685
					// _ = "end of CoverTab[72612]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:686
					_go_fuzz_dep_.CoverTab[72613]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:686
					// _ = "end of CoverTab[72613]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:686
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:686
				// _ = "end of CoverTab[72611]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:687
			// _ = "end of CoverTab[72610]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:688
			_go_fuzz_dep_.CoverTab[72614]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:688
			// _ = "end of CoverTab[72614]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:688
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:688
		// _ = "end of CoverTab[72607]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:689
		_go_fuzz_dep_.CoverTab[72615]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:689
		// _ = "end of CoverTab[72615]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:689
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:689
	// _ = "end of CoverTab[72598]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:689
	_go_fuzz_dep_.CoverTab[72599]++
										var flags Flags
										if endStream {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:691
		_go_fuzz_dep_.CoverTab[72616]++
											flags |= FlagDataEndStream
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:692
		// _ = "end of CoverTab[72616]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:693
		_go_fuzz_dep_.CoverTab[72617]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:693
		// _ = "end of CoverTab[72617]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:693
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:693
	// _ = "end of CoverTab[72599]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:693
	_go_fuzz_dep_.CoverTab[72600]++
										if pad != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:694
		_go_fuzz_dep_.CoverTab[72618]++
											flags |= FlagDataPadded
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:695
		// _ = "end of CoverTab[72618]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:696
		_go_fuzz_dep_.CoverTab[72619]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:696
		// _ = "end of CoverTab[72619]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:696
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:696
	// _ = "end of CoverTab[72600]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:696
	_go_fuzz_dep_.CoverTab[72601]++
										f.startWrite(FrameData, flags, streamID)
										if pad != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:698
		_go_fuzz_dep_.CoverTab[72620]++
											f.wbuf = append(f.wbuf, byte(len(pad)))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:699
		// _ = "end of CoverTab[72620]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:700
		_go_fuzz_dep_.CoverTab[72621]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:700
		// _ = "end of CoverTab[72621]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:700
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:700
	// _ = "end of CoverTab[72601]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:700
	_go_fuzz_dep_.CoverTab[72602]++
										f.wbuf = append(f.wbuf, data...)
										f.wbuf = append(f.wbuf, pad...)
										return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:703
	// _ = "end of CoverTab[72602]"
}

// A SettingsFrame conveys configuration parameters that affect how
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:706
// endpoints communicate, such as preferences and constraints on peer
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:706
// behavior.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:706
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:706
// See https://httpwg.org/specs/rfc7540.html#SETTINGS
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:711
type SettingsFrame struct {
	FrameHeader
	p	[]byte
}

func parseSettingsFrame(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (Frame, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:716
	_go_fuzz_dep_.CoverTab[72622]++
										if fh.Flags.Has(FlagSettingsAck) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:717
		_go_fuzz_dep_.CoverTab[72627]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:717
		return fh.Length > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:717
		// _ = "end of CoverTab[72627]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:717
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:717
		_go_fuzz_dep_.CoverTab[72628]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:724
		countError("frame_settings_ack_with_length")
											return nil, ConnectionError(ErrCodeFrameSize)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:725
		// _ = "end of CoverTab[72628]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:726
		_go_fuzz_dep_.CoverTab[72629]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:726
		// _ = "end of CoverTab[72629]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:726
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:726
	// _ = "end of CoverTab[72622]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:726
	_go_fuzz_dep_.CoverTab[72623]++
										if fh.StreamID != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:727
		_go_fuzz_dep_.CoverTab[72630]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:735
		countError("frame_settings_has_stream")
											return nil, ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:736
		// _ = "end of CoverTab[72630]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:737
		_go_fuzz_dep_.CoverTab[72631]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:737
		// _ = "end of CoverTab[72631]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:737
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:737
	// _ = "end of CoverTab[72623]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:737
	_go_fuzz_dep_.CoverTab[72624]++
										if len(p)%6 != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:738
		_go_fuzz_dep_.CoverTab[72632]++
											countError("frame_settings_mod_6")

											return nil, ConnectionError(ErrCodeFrameSize)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:741
		// _ = "end of CoverTab[72632]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:742
		_go_fuzz_dep_.CoverTab[72633]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:742
		// _ = "end of CoverTab[72633]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:742
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:742
	// _ = "end of CoverTab[72624]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:742
	_go_fuzz_dep_.CoverTab[72625]++
										f := &SettingsFrame{FrameHeader: fh, p: p}
										if v, ok := f.Value(SettingInitialWindowSize); ok && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:744
		_go_fuzz_dep_.CoverTab[72634]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:744
		return v > (1<<31)-1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:744
		// _ = "end of CoverTab[72634]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:744
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:744
		_go_fuzz_dep_.CoverTab[72635]++
											countError("frame_settings_window_size_too_big")

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:749
		return nil, ConnectionError(ErrCodeFlowControl)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:749
		// _ = "end of CoverTab[72635]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:750
		_go_fuzz_dep_.CoverTab[72636]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:750
		// _ = "end of CoverTab[72636]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:750
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:750
	// _ = "end of CoverTab[72625]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:750
	_go_fuzz_dep_.CoverTab[72626]++
										return f, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:751
	// _ = "end of CoverTab[72626]"
}

func (f *SettingsFrame) IsAck() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:754
	_go_fuzz_dep_.CoverTab[72637]++
										return f.FrameHeader.Flags.Has(FlagSettingsAck)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:755
	// _ = "end of CoverTab[72637]"
}

func (f *SettingsFrame) Value(id SettingID) (v uint32, ok bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:758
	_go_fuzz_dep_.CoverTab[72638]++
										f.checkValid()
										for i := 0; i < f.NumSettings(); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:760
		_go_fuzz_dep_.CoverTab[72640]++
											if s := f.Setting(i); s.ID == id {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:761
			_go_fuzz_dep_.CoverTab[72641]++
												return s.Val, true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:762
			// _ = "end of CoverTab[72641]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:763
			_go_fuzz_dep_.CoverTab[72642]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:763
			// _ = "end of CoverTab[72642]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:763
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:763
		// _ = "end of CoverTab[72640]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:764
	// _ = "end of CoverTab[72638]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:764
	_go_fuzz_dep_.CoverTab[72639]++
										return 0, false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:765
	// _ = "end of CoverTab[72639]"
}

// Setting returns the setting from the frame at the given 0-based index.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:768
// The index must be >= 0 and less than f.NumSettings().
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:770
func (f *SettingsFrame) Setting(i int) Setting {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:770
	_go_fuzz_dep_.CoverTab[72643]++
										buf := f.p
										return Setting{
		ID:	SettingID(binary.BigEndian.Uint16(buf[i*6 : i*6+2])),
		Val:	binary.BigEndian.Uint32(buf[i*6+2 : i*6+6]),
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:775
	// _ = "end of CoverTab[72643]"
}

func (f *SettingsFrame) NumSettings() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:778
	_go_fuzz_dep_.CoverTab[72644]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:778
	return len(f.p) / 6
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:778
	// _ = "end of CoverTab[72644]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:778
}

// HasDuplicates reports whether f contains any duplicate setting IDs.
func (f *SettingsFrame) HasDuplicates() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:781
	_go_fuzz_dep_.CoverTab[72645]++
										num := f.NumSettings()
										if num == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:783
		_go_fuzz_dep_.CoverTab[72649]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:784
		// _ = "end of CoverTab[72649]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:785
		_go_fuzz_dep_.CoverTab[72650]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:785
		// _ = "end of CoverTab[72650]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:785
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:785
	// _ = "end of CoverTab[72645]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:785
	_go_fuzz_dep_.CoverTab[72646]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:788
	if num < 10 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:788
		_go_fuzz_dep_.CoverTab[72651]++
											for i := 0; i < num; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:789
			_go_fuzz_dep_.CoverTab[72653]++
												idi := f.Setting(i).ID
												for j := i + 1; j < num; j++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:791
				_go_fuzz_dep_.CoverTab[72654]++
													idj := f.Setting(j).ID
													if idi == idj {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:793
					_go_fuzz_dep_.CoverTab[72655]++
														return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:794
					// _ = "end of CoverTab[72655]"
				} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:795
					_go_fuzz_dep_.CoverTab[72656]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:795
					// _ = "end of CoverTab[72656]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:795
				}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:795
				// _ = "end of CoverTab[72654]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:796
			// _ = "end of CoverTab[72653]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:797
		// _ = "end of CoverTab[72651]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:797
		_go_fuzz_dep_.CoverTab[72652]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:798
		// _ = "end of CoverTab[72652]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:799
		_go_fuzz_dep_.CoverTab[72657]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:799
		// _ = "end of CoverTab[72657]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:799
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:799
	// _ = "end of CoverTab[72646]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:799
	_go_fuzz_dep_.CoverTab[72647]++
										seen := map[SettingID]bool{}
										for i := 0; i < num; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:801
		_go_fuzz_dep_.CoverTab[72658]++
											id := f.Setting(i).ID
											if seen[id] {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:803
			_go_fuzz_dep_.CoverTab[72660]++
												return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:804
			// _ = "end of CoverTab[72660]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:805
			_go_fuzz_dep_.CoverTab[72661]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:805
			// _ = "end of CoverTab[72661]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:805
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:805
		// _ = "end of CoverTab[72658]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:805
		_go_fuzz_dep_.CoverTab[72659]++
											seen[id] = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:806
		// _ = "end of CoverTab[72659]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:807
	// _ = "end of CoverTab[72647]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:807
	_go_fuzz_dep_.CoverTab[72648]++
										return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:808
	// _ = "end of CoverTab[72648]"
}

// ForeachSetting runs fn for each setting.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:811
// It stops and returns the first error.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:813
func (f *SettingsFrame) ForeachSetting(fn func(Setting) error) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:813
	_go_fuzz_dep_.CoverTab[72662]++
										f.checkValid()
										for i := 0; i < f.NumSettings(); i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:815
		_go_fuzz_dep_.CoverTab[72664]++
											if err := fn(f.Setting(i)); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:816
			_go_fuzz_dep_.CoverTab[72665]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:817
			// _ = "end of CoverTab[72665]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:818
			_go_fuzz_dep_.CoverTab[72666]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:818
			// _ = "end of CoverTab[72666]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:818
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:818
		// _ = "end of CoverTab[72664]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:819
	// _ = "end of CoverTab[72662]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:819
	_go_fuzz_dep_.CoverTab[72663]++
										return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:820
	// _ = "end of CoverTab[72663]"
}

// WriteSettings writes a SETTINGS frame with zero or more settings
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:823
// specified and the ACK bit not set.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:823
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:823
// It will perform exactly one Write to the underlying Writer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:823
// It is the caller's responsibility to not call other Write methods concurrently.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:828
func (f *Framer) WriteSettings(settings ...Setting) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:828
	_go_fuzz_dep_.CoverTab[72667]++
										f.startWrite(FrameSettings, 0, 0)
										for _, s := range settings {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:830
		_go_fuzz_dep_.CoverTab[72669]++
											f.writeUint16(uint16(s.ID))
											f.writeUint32(s.Val)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:832
		// _ = "end of CoverTab[72669]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:833
	// _ = "end of CoverTab[72667]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:833
	_go_fuzz_dep_.CoverTab[72668]++
										return f.endWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:834
	// _ = "end of CoverTab[72668]"
}

// WriteSettingsAck writes an empty SETTINGS frame with the ACK bit set.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:837
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:837
// It will perform exactly one Write to the underlying Writer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:837
// It is the caller's responsibility to not call other Write methods concurrently.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:841
func (f *Framer) WriteSettingsAck() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:841
	_go_fuzz_dep_.CoverTab[72670]++
										f.startWrite(FrameSettings, FlagSettingsAck, 0)
										return f.endWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:843
	// _ = "end of CoverTab[72670]"
}

// A PingFrame is a mechanism for measuring a minimal round trip time
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:846
// from the sender, as well as determining whether an idle connection
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:846
// is still functional.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:846
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.7
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:850
type PingFrame struct {
	FrameHeader
	Data	[8]byte
}

func (f *PingFrame) IsAck() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:855
	_go_fuzz_dep_.CoverTab[72671]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:855
	return f.Flags.Has(FlagPingAck)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:855
	// _ = "end of CoverTab[72671]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:855
}

func parsePingFrame(_ *frameCache, fh FrameHeader, countError func(string), payload []byte) (Frame, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:857
	_go_fuzz_dep_.CoverTab[72672]++
										if len(payload) != 8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:858
		_go_fuzz_dep_.CoverTab[72675]++
											countError("frame_ping_length")
											return nil, ConnectionError(ErrCodeFrameSize)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:860
		// _ = "end of CoverTab[72675]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:861
		_go_fuzz_dep_.CoverTab[72676]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:861
		// _ = "end of CoverTab[72676]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:861
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:861
	// _ = "end of CoverTab[72672]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:861
	_go_fuzz_dep_.CoverTab[72673]++
										if fh.StreamID != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:862
		_go_fuzz_dep_.CoverTab[72677]++
											countError("frame_ping_has_stream")
											return nil, ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:864
		// _ = "end of CoverTab[72677]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:865
		_go_fuzz_dep_.CoverTab[72678]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:865
		// _ = "end of CoverTab[72678]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:865
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:865
	// _ = "end of CoverTab[72673]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:865
	_go_fuzz_dep_.CoverTab[72674]++
										f := &PingFrame{FrameHeader: fh}
										copy(f.Data[:], payload)
										return f, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:868
	// _ = "end of CoverTab[72674]"
}

func (f *Framer) WritePing(ack bool, data [8]byte) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:871
	_go_fuzz_dep_.CoverTab[72679]++
										var flags Flags
										if ack {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:873
		_go_fuzz_dep_.CoverTab[72681]++
											flags = FlagPingAck
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:874
		// _ = "end of CoverTab[72681]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:875
		_go_fuzz_dep_.CoverTab[72682]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:875
		// _ = "end of CoverTab[72682]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:875
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:875
	// _ = "end of CoverTab[72679]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:875
	_go_fuzz_dep_.CoverTab[72680]++
										f.startWrite(FramePing, flags, 0)
										f.writeBytes(data[:])
										return f.endWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:878
	// _ = "end of CoverTab[72680]"
}

// A GoAwayFrame informs the remote peer to stop creating streams on this connection.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:881
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.8
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:883
type GoAwayFrame struct {
	FrameHeader
	LastStreamID	uint32
	ErrCode		ErrCode
	debugData	[]byte
}

// DebugData returns any debug data in the GOAWAY frame. Its contents
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:890
// are not defined.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:890
// The caller must not retain the returned memory past the next
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:890
// call to ReadFrame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:894
func (f *GoAwayFrame) DebugData() []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:894
	_go_fuzz_dep_.CoverTab[72683]++
										f.checkValid()
										return f.debugData
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:896
	// _ = "end of CoverTab[72683]"
}

func parseGoAwayFrame(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (Frame, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:899
	_go_fuzz_dep_.CoverTab[72684]++
										if fh.StreamID != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:900
		_go_fuzz_dep_.CoverTab[72687]++
											countError("frame_goaway_has_stream")
											return nil, ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:902
		// _ = "end of CoverTab[72687]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:903
		_go_fuzz_dep_.CoverTab[72688]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:903
		// _ = "end of CoverTab[72688]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:903
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:903
	// _ = "end of CoverTab[72684]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:903
	_go_fuzz_dep_.CoverTab[72685]++
										if len(p) < 8 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:904
		_go_fuzz_dep_.CoverTab[72689]++
											countError("frame_goaway_short")
											return nil, ConnectionError(ErrCodeFrameSize)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:906
		// _ = "end of CoverTab[72689]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:907
		_go_fuzz_dep_.CoverTab[72690]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:907
		// _ = "end of CoverTab[72690]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:907
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:907
	// _ = "end of CoverTab[72685]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:907
	_go_fuzz_dep_.CoverTab[72686]++
										return &GoAwayFrame{
		FrameHeader:	fh,
		LastStreamID:	binary.BigEndian.Uint32(p[:4]) & (1<<31 - 1),
		ErrCode:	ErrCode(binary.BigEndian.Uint32(p[4:8])),
		debugData:	p[8:],
	}, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:913
	// _ = "end of CoverTab[72686]"
}

func (f *Framer) WriteGoAway(maxStreamID uint32, code ErrCode, debugData []byte) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:916
	_go_fuzz_dep_.CoverTab[72691]++
										f.startWrite(FrameGoAway, 0, 0)
										f.writeUint32(maxStreamID & (1<<31 - 1))
										f.writeUint32(uint32(code))
										f.writeBytes(debugData)
										return f.endWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:921
	// _ = "end of CoverTab[72691]"
}

// An UnknownFrame is the frame type returned when the frame type is unknown
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:924
// or no specific frame type parser exists.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:926
type UnknownFrame struct {
	FrameHeader
	p	[]byte
}

// Payload returns the frame's payload (after the header).  It is not
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:931
// valid to call this method after a subsequent call to
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:931
// Framer.ReadFrame, nor is it valid to retain the returned slice.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:931
// The memory is owned by the Framer and is invalidated when the next
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:931
// frame is read.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:936
func (f *UnknownFrame) Payload() []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:936
	_go_fuzz_dep_.CoverTab[72692]++
										f.checkValid()
										return f.p
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:938
	// _ = "end of CoverTab[72692]"
}

func parseUnknownFrame(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (Frame, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:941
	_go_fuzz_dep_.CoverTab[72693]++
										return &UnknownFrame{fh, p}, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:942
	// _ = "end of CoverTab[72693]"
}

// A WindowUpdateFrame is used to implement flow control.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:945
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.9
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:947
type WindowUpdateFrame struct {
	FrameHeader
	Increment	uint32	// never read with high bit set
}

func parseWindowUpdateFrame(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (Frame, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:952
	_go_fuzz_dep_.CoverTab[72694]++
										if len(p) != 4 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:953
		_go_fuzz_dep_.CoverTab[72697]++
											countError("frame_windowupdate_bad_len")
											return nil, ConnectionError(ErrCodeFrameSize)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:955
		// _ = "end of CoverTab[72697]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:956
		_go_fuzz_dep_.CoverTab[72698]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:956
		// _ = "end of CoverTab[72698]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:956
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:956
	// _ = "end of CoverTab[72694]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:956
	_go_fuzz_dep_.CoverTab[72695]++
										inc := binary.BigEndian.Uint32(p[:4]) & 0x7fffffff
										if inc == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:958
		_go_fuzz_dep_.CoverTab[72699]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:965
		if fh.StreamID == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:965
			_go_fuzz_dep_.CoverTab[72701]++
												countError("frame_windowupdate_zero_inc_conn")
												return nil, ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:967
			// _ = "end of CoverTab[72701]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:968
			_go_fuzz_dep_.CoverTab[72702]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:968
			// _ = "end of CoverTab[72702]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:968
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:968
		// _ = "end of CoverTab[72699]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:968
		_go_fuzz_dep_.CoverTab[72700]++
											countError("frame_windowupdate_zero_inc_stream")
											return nil, streamError(fh.StreamID, ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:970
		// _ = "end of CoverTab[72700]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:971
		_go_fuzz_dep_.CoverTab[72703]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:971
		// _ = "end of CoverTab[72703]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:971
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:971
	// _ = "end of CoverTab[72695]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:971
	_go_fuzz_dep_.CoverTab[72696]++
										return &WindowUpdateFrame{
		FrameHeader:	fh,
		Increment:	inc,
	}, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:975
	// _ = "end of CoverTab[72696]"
}

// WriteWindowUpdate writes a WINDOW_UPDATE frame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:978
// The increment value must be between 1 and 2,147,483,647, inclusive.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:978
// If the Stream ID is zero, the window update applies to the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:978
// connection as a whole.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:982
func (f *Framer) WriteWindowUpdate(streamID, incr uint32) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:982
	_go_fuzz_dep_.CoverTab[72704]++

										if (incr < 1 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:984
		_go_fuzz_dep_.CoverTab[72706]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:984
		return incr > 2147483647
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:984
		// _ = "end of CoverTab[72706]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:984
	}()) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:984
		_go_fuzz_dep_.CoverTab[72707]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:984
		return !f.AllowIllegalWrites
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:984
		// _ = "end of CoverTab[72707]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:984
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:984
		_go_fuzz_dep_.CoverTab[72708]++
											return errors.New("illegal window increment value")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:985
		// _ = "end of CoverTab[72708]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:986
		_go_fuzz_dep_.CoverTab[72709]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:986
		// _ = "end of CoverTab[72709]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:986
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:986
	// _ = "end of CoverTab[72704]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:986
	_go_fuzz_dep_.CoverTab[72705]++
										f.startWrite(FrameWindowUpdate, 0, streamID)
										f.writeUint32(incr)
										return f.endWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:989
	// _ = "end of CoverTab[72705]"
}

// A HeadersFrame is used to open a stream and additionally carries a
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:992
// header block fragment.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:994
type HeadersFrame struct {
	FrameHeader

	// Priority is set if FlagHeadersPriority is set in the FrameHeader.
	Priority	PriorityParam

	headerFragBuf	[]byte	// not owned
}

func (f *HeadersFrame) HeaderBlockFragment() []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1003
	_go_fuzz_dep_.CoverTab[72710]++
										f.checkValid()
										return f.headerFragBuf
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1005
	// _ = "end of CoverTab[72710]"
}

func (f *HeadersFrame) HeadersEnded() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1008
	_go_fuzz_dep_.CoverTab[72711]++
										return f.FrameHeader.Flags.Has(FlagHeadersEndHeaders)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1009
	// _ = "end of CoverTab[72711]"
}

func (f *HeadersFrame) StreamEnded() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1012
	_go_fuzz_dep_.CoverTab[72712]++
										return f.FrameHeader.Flags.Has(FlagHeadersEndStream)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1013
	// _ = "end of CoverTab[72712]"
}

func (f *HeadersFrame) HasPriority() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1016
	_go_fuzz_dep_.CoverTab[72713]++
										return f.FrameHeader.Flags.Has(FlagHeadersPriority)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1017
	// _ = "end of CoverTab[72713]"
}

func parseHeadersFrame(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (_ Frame, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1020
	_go_fuzz_dep_.CoverTab[72714]++
										hf := &HeadersFrame{
		FrameHeader: fh,
	}
	if fh.StreamID == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1024
		_go_fuzz_dep_.CoverTab[72719]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1029
		countError("frame_headers_zero_stream")
											return nil, connError{ErrCodeProtocol, "HEADERS frame with stream ID 0"}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1030
		// _ = "end of CoverTab[72719]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1031
		_go_fuzz_dep_.CoverTab[72720]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1031
		// _ = "end of CoverTab[72720]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1031
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1031
	// _ = "end of CoverTab[72714]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1031
	_go_fuzz_dep_.CoverTab[72715]++
										var padLength uint8
										if fh.Flags.Has(FlagHeadersPadded) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1033
		_go_fuzz_dep_.CoverTab[72721]++
											if p, padLength, err = readByte(p); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1034
			_go_fuzz_dep_.CoverTab[72722]++
												countError("frame_headers_pad_short")
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1036
			// _ = "end of CoverTab[72722]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1037
			_go_fuzz_dep_.CoverTab[72723]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1037
			// _ = "end of CoverTab[72723]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1037
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1037
		// _ = "end of CoverTab[72721]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1038
		_go_fuzz_dep_.CoverTab[72724]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1038
		// _ = "end of CoverTab[72724]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1038
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1038
	// _ = "end of CoverTab[72715]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1038
	_go_fuzz_dep_.CoverTab[72716]++
										if fh.Flags.Has(FlagHeadersPriority) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1039
		_go_fuzz_dep_.CoverTab[72725]++
											var v uint32
											p, v, err = readUint32(p)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1042
			_go_fuzz_dep_.CoverTab[72727]++
												countError("frame_headers_prio_short")
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1044
			// _ = "end of CoverTab[72727]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1045
			_go_fuzz_dep_.CoverTab[72728]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1045
			// _ = "end of CoverTab[72728]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1045
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1045
		// _ = "end of CoverTab[72725]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1045
		_go_fuzz_dep_.CoverTab[72726]++
											hf.Priority.StreamDep = v & 0x7fffffff
											hf.Priority.Exclusive = (v != hf.Priority.StreamDep)
											p, hf.Priority.Weight, err = readByte(p)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1049
			_go_fuzz_dep_.CoverTab[72729]++
												countError("frame_headers_prio_weight_short")
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1051
			// _ = "end of CoverTab[72729]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1052
			_go_fuzz_dep_.CoverTab[72730]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1052
			// _ = "end of CoverTab[72730]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1052
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1052
		// _ = "end of CoverTab[72726]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1053
		_go_fuzz_dep_.CoverTab[72731]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1053
		// _ = "end of CoverTab[72731]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1053
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1053
	// _ = "end of CoverTab[72716]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1053
	_go_fuzz_dep_.CoverTab[72717]++
										if len(p)-int(padLength) < 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1054
		_go_fuzz_dep_.CoverTab[72732]++
											countError("frame_headers_pad_too_big")
											return nil, streamError(fh.StreamID, ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1056
		// _ = "end of CoverTab[72732]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1057
		_go_fuzz_dep_.CoverTab[72733]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1057
		// _ = "end of CoverTab[72733]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1057
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1057
	// _ = "end of CoverTab[72717]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1057
	_go_fuzz_dep_.CoverTab[72718]++
										hf.headerFragBuf = p[:len(p)-int(padLength)]
										return hf, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1059
	// _ = "end of CoverTab[72718]"
}

// HeadersFrameParam are the parameters for writing a HEADERS frame.
type HeadersFrameParam struct {
	// StreamID is the required Stream ID to initiate.
	StreamID	uint32
	// BlockFragment is part (or all) of a Header Block.
	BlockFragment	[]byte

	// EndStream indicates that the header block is the last that
	// the endpoint will send for the identified stream. Setting
	// this flag causes the stream to enter one of "half closed"
	// states.
	EndStream	bool

	// EndHeaders indicates that this frame contains an entire
	// header block and is not followed by any
	// CONTINUATION frames.
	EndHeaders	bool

	// PadLength is the optional number of bytes of zeros to add
	// to this frame.
	PadLength	uint8

	// Priority, if non-zero, includes stream priority information
	// in the HEADER frame.
	Priority	PriorityParam
}

// WriteHeaders writes a single HEADERS frame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1089
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1089
// This is a low-level header writing method. Encoding headers and
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1089
// splitting them into any necessary CONTINUATION frames is handled
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1089
// elsewhere.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1089
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1089
// It will perform exactly one Write to the underlying Writer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1089
// It is the caller's responsibility to not call other Write methods concurrently.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1097
func (f *Framer) WriteHeaders(p HeadersFrameParam) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1097
	_go_fuzz_dep_.CoverTab[72734]++
										if !validStreamID(p.StreamID) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1098
		_go_fuzz_dep_.CoverTab[72742]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1098
		return !f.AllowIllegalWrites
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1098
		// _ = "end of CoverTab[72742]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1098
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1098
		_go_fuzz_dep_.CoverTab[72743]++
											return errStreamID
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1099
		// _ = "end of CoverTab[72743]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1100
		_go_fuzz_dep_.CoverTab[72744]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1100
		// _ = "end of CoverTab[72744]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1100
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1100
	// _ = "end of CoverTab[72734]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1100
	_go_fuzz_dep_.CoverTab[72735]++
										var flags Flags
										if p.PadLength != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1102
		_go_fuzz_dep_.CoverTab[72745]++
											flags |= FlagHeadersPadded
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1103
		// _ = "end of CoverTab[72745]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1104
		_go_fuzz_dep_.CoverTab[72746]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1104
		// _ = "end of CoverTab[72746]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1104
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1104
	// _ = "end of CoverTab[72735]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1104
	_go_fuzz_dep_.CoverTab[72736]++
										if p.EndStream {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1105
		_go_fuzz_dep_.CoverTab[72747]++
											flags |= FlagHeadersEndStream
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1106
		// _ = "end of CoverTab[72747]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1107
		_go_fuzz_dep_.CoverTab[72748]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1107
		// _ = "end of CoverTab[72748]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1107
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1107
	// _ = "end of CoverTab[72736]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1107
	_go_fuzz_dep_.CoverTab[72737]++
										if p.EndHeaders {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1108
		_go_fuzz_dep_.CoverTab[72749]++
											flags |= FlagHeadersEndHeaders
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1109
		// _ = "end of CoverTab[72749]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1110
		_go_fuzz_dep_.CoverTab[72750]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1110
		// _ = "end of CoverTab[72750]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1110
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1110
	// _ = "end of CoverTab[72737]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1110
	_go_fuzz_dep_.CoverTab[72738]++
										if !p.Priority.IsZero() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1111
		_go_fuzz_dep_.CoverTab[72751]++
											flags |= FlagHeadersPriority
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1112
		// _ = "end of CoverTab[72751]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1113
		_go_fuzz_dep_.CoverTab[72752]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1113
		// _ = "end of CoverTab[72752]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1113
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1113
	// _ = "end of CoverTab[72738]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1113
	_go_fuzz_dep_.CoverTab[72739]++
										f.startWrite(FrameHeaders, flags, p.StreamID)
										if p.PadLength != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1115
		_go_fuzz_dep_.CoverTab[72753]++
											f.writeByte(p.PadLength)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1116
		// _ = "end of CoverTab[72753]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1117
		_go_fuzz_dep_.CoverTab[72754]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1117
		// _ = "end of CoverTab[72754]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1117
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1117
	// _ = "end of CoverTab[72739]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1117
	_go_fuzz_dep_.CoverTab[72740]++
										if !p.Priority.IsZero() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1118
		_go_fuzz_dep_.CoverTab[72755]++
											v := p.Priority.StreamDep
											if !validStreamIDOrZero(v) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1120
			_go_fuzz_dep_.CoverTab[72758]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1120
			return !f.AllowIllegalWrites
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1120
			// _ = "end of CoverTab[72758]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1120
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1120
			_go_fuzz_dep_.CoverTab[72759]++
												return errDepStreamID
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1121
			// _ = "end of CoverTab[72759]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1122
			_go_fuzz_dep_.CoverTab[72760]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1122
			// _ = "end of CoverTab[72760]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1122
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1122
		// _ = "end of CoverTab[72755]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1122
		_go_fuzz_dep_.CoverTab[72756]++
											if p.Priority.Exclusive {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1123
			_go_fuzz_dep_.CoverTab[72761]++
												v |= 1 << 31
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1124
			// _ = "end of CoverTab[72761]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1125
			_go_fuzz_dep_.CoverTab[72762]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1125
			// _ = "end of CoverTab[72762]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1125
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1125
		// _ = "end of CoverTab[72756]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1125
		_go_fuzz_dep_.CoverTab[72757]++
											f.writeUint32(v)
											f.writeByte(p.Priority.Weight)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1127
		// _ = "end of CoverTab[72757]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1128
		_go_fuzz_dep_.CoverTab[72763]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1128
		// _ = "end of CoverTab[72763]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1128
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1128
	// _ = "end of CoverTab[72740]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1128
	_go_fuzz_dep_.CoverTab[72741]++
										f.wbuf = append(f.wbuf, p.BlockFragment...)
										f.wbuf = append(f.wbuf, padZeros[:p.PadLength]...)
										return f.endWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1131
	// _ = "end of CoverTab[72741]"
}

// A PriorityFrame specifies the sender-advised priority of a stream.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1134
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.3
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1136
type PriorityFrame struct {
	FrameHeader
	PriorityParam
}

// PriorityParam are the stream prioritzation parameters.
type PriorityParam struct {
	// StreamDep is a 31-bit stream identifier for the
	// stream that this stream depends on. Zero means no
	// dependency.
	StreamDep	uint32

	// Exclusive is whether the dependency is exclusive.
	Exclusive	bool

	// Weight is the stream's zero-indexed weight. It should be
	// set together with StreamDep, or neither should be set. Per
	// the spec, "Add one to the value to obtain a weight between
	// 1 and 256."
	Weight	uint8
}

func (p PriorityParam) IsZero() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1158
	_go_fuzz_dep_.CoverTab[72764]++
										return p == PriorityParam{}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1159
	// _ = "end of CoverTab[72764]"
}

func parsePriorityFrame(_ *frameCache, fh FrameHeader, countError func(string), payload []byte) (Frame, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1162
	_go_fuzz_dep_.CoverTab[72765]++
										if fh.StreamID == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1163
		_go_fuzz_dep_.CoverTab[72768]++
											countError("frame_priority_zero_stream")
											return nil, connError{ErrCodeProtocol, "PRIORITY frame with stream ID 0"}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1165
		// _ = "end of CoverTab[72768]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1166
		_go_fuzz_dep_.CoverTab[72769]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1166
		// _ = "end of CoverTab[72769]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1166
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1166
	// _ = "end of CoverTab[72765]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1166
	_go_fuzz_dep_.CoverTab[72766]++
										if len(payload) != 5 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1167
		_go_fuzz_dep_.CoverTab[72770]++
											countError("frame_priority_bad_length")
											return nil, connError{ErrCodeFrameSize, fmt.Sprintf("PRIORITY frame payload size was %d; want 5", len(payload))}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1169
		// _ = "end of CoverTab[72770]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1170
		_go_fuzz_dep_.CoverTab[72771]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1170
		// _ = "end of CoverTab[72771]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1170
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1170
	// _ = "end of CoverTab[72766]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1170
	_go_fuzz_dep_.CoverTab[72767]++
										v := binary.BigEndian.Uint32(payload[:4])
										streamID := v & 0x7fffffff
										return &PriorityFrame{
		FrameHeader:	fh,
		PriorityParam: PriorityParam{
			Weight:		payload[4],
			StreamDep:	streamID,
			Exclusive:	streamID != v,
		},
	}, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1180
	// _ = "end of CoverTab[72767]"
}

// WritePriority writes a PRIORITY frame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1183
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1183
// It will perform exactly one Write to the underlying Writer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1183
// It is the caller's responsibility to not call other Write methods concurrently.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1187
func (f *Framer) WritePriority(streamID uint32, p PriorityParam) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1187
	_go_fuzz_dep_.CoverTab[72772]++
										if !validStreamID(streamID) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1188
		_go_fuzz_dep_.CoverTab[72776]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1188
		return !f.AllowIllegalWrites
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1188
		// _ = "end of CoverTab[72776]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1188
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1188
		_go_fuzz_dep_.CoverTab[72777]++
											return errStreamID
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1189
		// _ = "end of CoverTab[72777]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1190
		_go_fuzz_dep_.CoverTab[72778]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1190
		// _ = "end of CoverTab[72778]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1190
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1190
	// _ = "end of CoverTab[72772]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1190
	_go_fuzz_dep_.CoverTab[72773]++
										if !validStreamIDOrZero(p.StreamDep) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1191
		_go_fuzz_dep_.CoverTab[72779]++
											return errDepStreamID
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1192
		// _ = "end of CoverTab[72779]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1193
		_go_fuzz_dep_.CoverTab[72780]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1193
		// _ = "end of CoverTab[72780]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1193
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1193
	// _ = "end of CoverTab[72773]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1193
	_go_fuzz_dep_.CoverTab[72774]++
										f.startWrite(FramePriority, 0, streamID)
										v := p.StreamDep
										if p.Exclusive {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1196
		_go_fuzz_dep_.CoverTab[72781]++
											v |= 1 << 31
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1197
		// _ = "end of CoverTab[72781]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1198
		_go_fuzz_dep_.CoverTab[72782]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1198
		// _ = "end of CoverTab[72782]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1198
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1198
	// _ = "end of CoverTab[72774]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1198
	_go_fuzz_dep_.CoverTab[72775]++
										f.writeUint32(v)
										f.writeByte(p.Weight)
										return f.endWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1201
	// _ = "end of CoverTab[72775]"
}

// A RSTStreamFrame allows for abnormal termination of a stream.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1204
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.4
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1206
type RSTStreamFrame struct {
	FrameHeader
	ErrCode	ErrCode
}

func parseRSTStreamFrame(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (Frame, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1211
	_go_fuzz_dep_.CoverTab[72783]++
										if len(p) != 4 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1212
		_go_fuzz_dep_.CoverTab[72786]++
											countError("frame_rststream_bad_len")
											return nil, ConnectionError(ErrCodeFrameSize)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1214
		// _ = "end of CoverTab[72786]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1215
		_go_fuzz_dep_.CoverTab[72787]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1215
		// _ = "end of CoverTab[72787]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1215
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1215
	// _ = "end of CoverTab[72783]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1215
	_go_fuzz_dep_.CoverTab[72784]++
										if fh.StreamID == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1216
		_go_fuzz_dep_.CoverTab[72788]++
											countError("frame_rststream_zero_stream")
											return nil, ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1218
		// _ = "end of CoverTab[72788]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1219
		_go_fuzz_dep_.CoverTab[72789]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1219
		// _ = "end of CoverTab[72789]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1219
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1219
	// _ = "end of CoverTab[72784]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1219
	_go_fuzz_dep_.CoverTab[72785]++
										return &RSTStreamFrame{fh, ErrCode(binary.BigEndian.Uint32(p[:4]))}, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1220
	// _ = "end of CoverTab[72785]"
}

// WriteRSTStream writes a RST_STREAM frame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1223
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1223
// It will perform exactly one Write to the underlying Writer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1223
// It is the caller's responsibility to not call other Write methods concurrently.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1227
func (f *Framer) WriteRSTStream(streamID uint32, code ErrCode) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1227
	_go_fuzz_dep_.CoverTab[72790]++
										if !validStreamID(streamID) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1228
		_go_fuzz_dep_.CoverTab[72792]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1228
		return !f.AllowIllegalWrites
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1228
		// _ = "end of CoverTab[72792]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1228
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1228
		_go_fuzz_dep_.CoverTab[72793]++
											return errStreamID
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1229
		// _ = "end of CoverTab[72793]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1230
		_go_fuzz_dep_.CoverTab[72794]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1230
		// _ = "end of CoverTab[72794]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1230
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1230
	// _ = "end of CoverTab[72790]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1230
	_go_fuzz_dep_.CoverTab[72791]++
										f.startWrite(FrameRSTStream, 0, streamID)
										f.writeUint32(uint32(code))
										return f.endWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1233
	// _ = "end of CoverTab[72791]"
}

// A ContinuationFrame is used to continue a sequence of header block fragments.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1236
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.10
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1238
type ContinuationFrame struct {
	FrameHeader
	headerFragBuf	[]byte
}

func parseContinuationFrame(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (Frame, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1243
	_go_fuzz_dep_.CoverTab[72795]++
										if fh.StreamID == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1244
		_go_fuzz_dep_.CoverTab[72797]++
											countError("frame_continuation_zero_stream")
											return nil, connError{ErrCodeProtocol, "CONTINUATION frame with stream ID 0"}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1246
		// _ = "end of CoverTab[72797]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1247
		_go_fuzz_dep_.CoverTab[72798]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1247
		// _ = "end of CoverTab[72798]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1247
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1247
	// _ = "end of CoverTab[72795]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1247
	_go_fuzz_dep_.CoverTab[72796]++
										return &ContinuationFrame{fh, p}, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1248
	// _ = "end of CoverTab[72796]"
}

func (f *ContinuationFrame) HeaderBlockFragment() []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1251
	_go_fuzz_dep_.CoverTab[72799]++
										f.checkValid()
										return f.headerFragBuf
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1253
	// _ = "end of CoverTab[72799]"
}

func (f *ContinuationFrame) HeadersEnded() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1256
	_go_fuzz_dep_.CoverTab[72800]++
										return f.FrameHeader.Flags.Has(FlagContinuationEndHeaders)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1257
	// _ = "end of CoverTab[72800]"
}

// WriteContinuation writes a CONTINUATION frame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1260
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1260
// It will perform exactly one Write to the underlying Writer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1260
// It is the caller's responsibility to not call other Write methods concurrently.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1264
func (f *Framer) WriteContinuation(streamID uint32, endHeaders bool, headerBlockFragment []byte) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1264
	_go_fuzz_dep_.CoverTab[72801]++
										if !validStreamID(streamID) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1265
		_go_fuzz_dep_.CoverTab[72804]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1265
		return !f.AllowIllegalWrites
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1265
		// _ = "end of CoverTab[72804]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1265
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1265
		_go_fuzz_dep_.CoverTab[72805]++
											return errStreamID
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1266
		// _ = "end of CoverTab[72805]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1267
		_go_fuzz_dep_.CoverTab[72806]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1267
		// _ = "end of CoverTab[72806]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1267
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1267
	// _ = "end of CoverTab[72801]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1267
	_go_fuzz_dep_.CoverTab[72802]++
										var flags Flags
										if endHeaders {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1269
		_go_fuzz_dep_.CoverTab[72807]++
											flags |= FlagContinuationEndHeaders
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1270
		// _ = "end of CoverTab[72807]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1271
		_go_fuzz_dep_.CoverTab[72808]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1271
		// _ = "end of CoverTab[72808]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1271
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1271
	// _ = "end of CoverTab[72802]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1271
	_go_fuzz_dep_.CoverTab[72803]++
										f.startWrite(FrameContinuation, flags, streamID)
										f.wbuf = append(f.wbuf, headerBlockFragment...)
										return f.endWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1274
	// _ = "end of CoverTab[72803]"
}

// A PushPromiseFrame is used to initiate a server stream.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1277
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.6
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1279
type PushPromiseFrame struct {
	FrameHeader
	PromiseID	uint32
	headerFragBuf	[]byte	// not owned
}

func (f *PushPromiseFrame) HeaderBlockFragment() []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1285
	_go_fuzz_dep_.CoverTab[72809]++
										f.checkValid()
										return f.headerFragBuf
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1287
	// _ = "end of CoverTab[72809]"
}

func (f *PushPromiseFrame) HeadersEnded() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1290
	_go_fuzz_dep_.CoverTab[72810]++
										return f.FrameHeader.Flags.Has(FlagPushPromiseEndHeaders)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1291
	// _ = "end of CoverTab[72810]"
}

func parsePushPromise(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (_ Frame, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1294
	_go_fuzz_dep_.CoverTab[72811]++
										pp := &PushPromiseFrame{
		FrameHeader: fh,
	}
	if pp.StreamID == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1298
		_go_fuzz_dep_.CoverTab[72816]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1305
		countError("frame_pushpromise_zero_stream")
											return nil, ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1306
		// _ = "end of CoverTab[72816]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1307
		_go_fuzz_dep_.CoverTab[72817]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1307
		// _ = "end of CoverTab[72817]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1307
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1307
	// _ = "end of CoverTab[72811]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1307
	_go_fuzz_dep_.CoverTab[72812]++
	// The PUSH_PROMISE frame includes optional padding.
	// Padding fields and flags are identical to those defined for DATA frames
	var padLength uint8
	if fh.Flags.Has(FlagPushPromisePadded) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1311
		_go_fuzz_dep_.CoverTab[72818]++
											if p, padLength, err = readByte(p); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1312
			_go_fuzz_dep_.CoverTab[72819]++
												countError("frame_pushpromise_pad_short")
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1314
			// _ = "end of CoverTab[72819]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1315
			_go_fuzz_dep_.CoverTab[72820]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1315
			// _ = "end of CoverTab[72820]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1315
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1315
		// _ = "end of CoverTab[72818]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1316
		_go_fuzz_dep_.CoverTab[72821]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1316
		// _ = "end of CoverTab[72821]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1316
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1316
	// _ = "end of CoverTab[72812]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1316
	_go_fuzz_dep_.CoverTab[72813]++

										p, pp.PromiseID, err = readUint32(p)
										if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1319
		_go_fuzz_dep_.CoverTab[72822]++
											countError("frame_pushpromise_promiseid_short")
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1321
		// _ = "end of CoverTab[72822]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1322
		_go_fuzz_dep_.CoverTab[72823]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1322
		// _ = "end of CoverTab[72823]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1322
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1322
	// _ = "end of CoverTab[72813]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1322
	_go_fuzz_dep_.CoverTab[72814]++
										pp.PromiseID = pp.PromiseID & (1<<31 - 1)

										if int(padLength) > len(p) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1325
		_go_fuzz_dep_.CoverTab[72824]++

											countError("frame_pushpromise_pad_too_big")
											return nil, ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1328
		// _ = "end of CoverTab[72824]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1329
		_go_fuzz_dep_.CoverTab[72825]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1329
		// _ = "end of CoverTab[72825]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1329
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1329
	// _ = "end of CoverTab[72814]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1329
	_go_fuzz_dep_.CoverTab[72815]++
										pp.headerFragBuf = p[:len(p)-int(padLength)]
										return pp, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1331
	// _ = "end of CoverTab[72815]"
}

// PushPromiseParam are the parameters for writing a PUSH_PROMISE frame.
type PushPromiseParam struct {
	// StreamID is the required Stream ID to initiate.
	StreamID	uint32

	// PromiseID is the required Stream ID which this
	// Push Promises
	PromiseID	uint32

	// BlockFragment is part (or all) of a Header Block.
	BlockFragment	[]byte

	// EndHeaders indicates that this frame contains an entire
	// header block and is not followed by any
	// CONTINUATION frames.
	EndHeaders	bool

	// PadLength is the optional number of bytes of zeros to add
	// to this frame.
	PadLength	uint8
}

// WritePushPromise writes a single PushPromise Frame.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1356
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1356
// As with Header Frames, This is the low level call for writing
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1356
// individual frames. Continuation frames are handled elsewhere.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1356
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1356
// It will perform exactly one Write to the underlying Writer.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1356
// It is the caller's responsibility to not call other Write methods concurrently.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1363
func (f *Framer) WritePushPromise(p PushPromiseParam) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1363
	_go_fuzz_dep_.CoverTab[72826]++
										if !validStreamID(p.StreamID) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1364
		_go_fuzz_dep_.CoverTab[72832]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1364
		return !f.AllowIllegalWrites
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1364
		// _ = "end of CoverTab[72832]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1364
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1364
		_go_fuzz_dep_.CoverTab[72833]++
											return errStreamID
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1365
		// _ = "end of CoverTab[72833]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1366
		_go_fuzz_dep_.CoverTab[72834]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1366
		// _ = "end of CoverTab[72834]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1366
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1366
	// _ = "end of CoverTab[72826]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1366
	_go_fuzz_dep_.CoverTab[72827]++
										var flags Flags
										if p.PadLength != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1368
		_go_fuzz_dep_.CoverTab[72835]++
											flags |= FlagPushPromisePadded
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1369
		// _ = "end of CoverTab[72835]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1370
		_go_fuzz_dep_.CoverTab[72836]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1370
		// _ = "end of CoverTab[72836]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1370
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1370
	// _ = "end of CoverTab[72827]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1370
	_go_fuzz_dep_.CoverTab[72828]++
										if p.EndHeaders {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1371
		_go_fuzz_dep_.CoverTab[72837]++
											flags |= FlagPushPromiseEndHeaders
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1372
		// _ = "end of CoverTab[72837]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1373
		_go_fuzz_dep_.CoverTab[72838]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1373
		// _ = "end of CoverTab[72838]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1373
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1373
	// _ = "end of CoverTab[72828]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1373
	_go_fuzz_dep_.CoverTab[72829]++
										f.startWrite(FramePushPromise, flags, p.StreamID)
										if p.PadLength != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1375
		_go_fuzz_dep_.CoverTab[72839]++
											f.writeByte(p.PadLength)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1376
		// _ = "end of CoverTab[72839]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1377
		_go_fuzz_dep_.CoverTab[72840]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1377
		// _ = "end of CoverTab[72840]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1377
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1377
	// _ = "end of CoverTab[72829]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1377
	_go_fuzz_dep_.CoverTab[72830]++
										if !validStreamID(p.PromiseID) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1378
		_go_fuzz_dep_.CoverTab[72841]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1378
		return !f.AllowIllegalWrites
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1378
		// _ = "end of CoverTab[72841]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1378
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1378
		_go_fuzz_dep_.CoverTab[72842]++
											return errStreamID
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1379
		// _ = "end of CoverTab[72842]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1380
		_go_fuzz_dep_.CoverTab[72843]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1380
		// _ = "end of CoverTab[72843]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1380
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1380
	// _ = "end of CoverTab[72830]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1380
	_go_fuzz_dep_.CoverTab[72831]++
										f.writeUint32(p.PromiseID)
										f.wbuf = append(f.wbuf, p.BlockFragment...)
										f.wbuf = append(f.wbuf, padZeros[:p.PadLength]...)
										return f.endWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1384
	// _ = "end of CoverTab[72831]"
}

// WriteRawFrame writes a raw frame. This can be used to write
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1387
// extension frames unknown to this package.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1389
func (f *Framer) WriteRawFrame(t FrameType, flags Flags, streamID uint32, payload []byte) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1389
	_go_fuzz_dep_.CoverTab[72844]++
										f.startWrite(t, flags, streamID)
										f.writeBytes(payload)
										return f.endWrite()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1392
	// _ = "end of CoverTab[72844]"
}

func readByte(p []byte) (remain []byte, b byte, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1395
	_go_fuzz_dep_.CoverTab[72845]++
										if len(p) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1396
		_go_fuzz_dep_.CoverTab[72847]++
											return nil, 0, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1397
		// _ = "end of CoverTab[72847]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1398
		_go_fuzz_dep_.CoverTab[72848]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1398
		// _ = "end of CoverTab[72848]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1398
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1398
	// _ = "end of CoverTab[72845]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1398
	_go_fuzz_dep_.CoverTab[72846]++
										return p[1:], p[0], nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1399
	// _ = "end of CoverTab[72846]"
}

func readUint32(p []byte) (remain []byte, v uint32, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1402
	_go_fuzz_dep_.CoverTab[72849]++
										if len(p) < 4 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1403
		_go_fuzz_dep_.CoverTab[72851]++
											return nil, 0, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1404
		// _ = "end of CoverTab[72851]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1405
		_go_fuzz_dep_.CoverTab[72852]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1405
		// _ = "end of CoverTab[72852]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1405
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1405
	// _ = "end of CoverTab[72849]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1405
	_go_fuzz_dep_.CoverTab[72850]++
										return p[4:], binary.BigEndian.Uint32(p[:4]), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1406
	// _ = "end of CoverTab[72850]"
}

type streamEnder interface {
	StreamEnded() bool
}

type headersEnder interface {
	HeadersEnded() bool
}

type headersOrContinuation interface {
	headersEnder
	HeaderBlockFragment() []byte
}

// A MetaHeadersFrame is the representation of one HEADERS frame and
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1422
// zero or more contiguous CONTINUATION frames and the decoding of
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1422
// their HPACK-encoded contents.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1422
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1422
// This type of frame does not appear on the wire and is only returned
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1422
// by the Framer when Framer.ReadMetaHeaders is set.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1428
type MetaHeadersFrame struct {
	*HeadersFrame

	// Fields are the fields contained in the HEADERS and
	// CONTINUATION frames. The underlying slice is owned by the
	// Framer and must not be retained after the next call to
	// ReadFrame.
	//
	// Fields are guaranteed to be in the correct http2 order and
	// not have unknown pseudo header fields or invalid header
	// field names or values. Required pseudo header fields may be
	// missing, however. Use the MetaHeadersFrame.Pseudo accessor
	// method access pseudo headers.
	Fields	[]hpack.HeaderField

	// Truncated is whether the max header list size limit was hit
	// and Fields is incomplete. The hpack decoder state is still
	// valid, however.
	Truncated	bool
}

// PseudoValue returns the given pseudo header field's value.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1449
// The provided pseudo field should not contain the leading colon.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1451
func (mh *MetaHeadersFrame) PseudoValue(pseudo string) string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1451
	_go_fuzz_dep_.CoverTab[72853]++
										for _, hf := range mh.Fields {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1452
		_go_fuzz_dep_.CoverTab[72855]++
											if !hf.IsPseudo() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1453
			_go_fuzz_dep_.CoverTab[72857]++
												return ""
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1454
			// _ = "end of CoverTab[72857]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1455
			_go_fuzz_dep_.CoverTab[72858]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1455
			// _ = "end of CoverTab[72858]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1455
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1455
		// _ = "end of CoverTab[72855]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1455
		_go_fuzz_dep_.CoverTab[72856]++
											if hf.Name[1:] == pseudo {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1456
			_go_fuzz_dep_.CoverTab[72859]++
												return hf.Value
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1457
			// _ = "end of CoverTab[72859]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1458
			_go_fuzz_dep_.CoverTab[72860]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1458
			// _ = "end of CoverTab[72860]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1458
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1458
		// _ = "end of CoverTab[72856]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1459
	// _ = "end of CoverTab[72853]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1459
	_go_fuzz_dep_.CoverTab[72854]++
										return ""
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1460
	// _ = "end of CoverTab[72854]"
}

// RegularFields returns the regular (non-pseudo) header fields of mh.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1463
// The caller does not own the returned slice.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1465
func (mh *MetaHeadersFrame) RegularFields() []hpack.HeaderField {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1465
	_go_fuzz_dep_.CoverTab[72861]++
										for i, hf := range mh.Fields {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1466
		_go_fuzz_dep_.CoverTab[72863]++
											if !hf.IsPseudo() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1467
			_go_fuzz_dep_.CoverTab[72864]++
												return mh.Fields[i:]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1468
			// _ = "end of CoverTab[72864]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1469
			_go_fuzz_dep_.CoverTab[72865]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1469
			// _ = "end of CoverTab[72865]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1469
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1469
		// _ = "end of CoverTab[72863]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1470
	// _ = "end of CoverTab[72861]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1470
	_go_fuzz_dep_.CoverTab[72862]++
										return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1471
	// _ = "end of CoverTab[72862]"
}

// PseudoFields returns the pseudo header fields of mh.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1474
// The caller does not own the returned slice.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1476
func (mh *MetaHeadersFrame) PseudoFields() []hpack.HeaderField {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1476
	_go_fuzz_dep_.CoverTab[72866]++
										for i, hf := range mh.Fields {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1477
		_go_fuzz_dep_.CoverTab[72868]++
											if !hf.IsPseudo() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1478
			_go_fuzz_dep_.CoverTab[72869]++
												return mh.Fields[:i]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1479
			// _ = "end of CoverTab[72869]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1480
			_go_fuzz_dep_.CoverTab[72870]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1480
			// _ = "end of CoverTab[72870]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1480
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1480
		// _ = "end of CoverTab[72868]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1481
	// _ = "end of CoverTab[72866]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1481
	_go_fuzz_dep_.CoverTab[72867]++
										return mh.Fields
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1482
	// _ = "end of CoverTab[72867]"
}

func (mh *MetaHeadersFrame) checkPseudos() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1485
	_go_fuzz_dep_.CoverTab[72871]++
										var isRequest, isResponse bool
										pf := mh.PseudoFields()
										for i, hf := range pf {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1488
		_go_fuzz_dep_.CoverTab[72874]++
											switch hf.Name {
		case ":method", ":path", ":scheme", ":authority":
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1490
			_go_fuzz_dep_.CoverTab[72876]++
												isRequest = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1491
			// _ = "end of CoverTab[72876]"
		case ":status":
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1492
			_go_fuzz_dep_.CoverTab[72877]++
												isResponse = true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1493
			// _ = "end of CoverTab[72877]"
		default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1494
			_go_fuzz_dep_.CoverTab[72878]++
												return pseudoHeaderError(hf.Name)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1495
			// _ = "end of CoverTab[72878]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1496
		// _ = "end of CoverTab[72874]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1496
		_go_fuzz_dep_.CoverTab[72875]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1500
		for _, hf2 := range pf[:i] {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1500
			_go_fuzz_dep_.CoverTab[72879]++
												if hf.Name == hf2.Name {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1501
				_go_fuzz_dep_.CoverTab[72880]++
													return duplicatePseudoHeaderError(hf.Name)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1502
				// _ = "end of CoverTab[72880]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1503
				_go_fuzz_dep_.CoverTab[72881]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1503
				// _ = "end of CoverTab[72881]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1503
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1503
			// _ = "end of CoverTab[72879]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1504
		// _ = "end of CoverTab[72875]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1505
	// _ = "end of CoverTab[72871]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1505
	_go_fuzz_dep_.CoverTab[72872]++
										if isRequest && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1506
		_go_fuzz_dep_.CoverTab[72882]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1506
		return isResponse
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1506
		// _ = "end of CoverTab[72882]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1506
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1506
		_go_fuzz_dep_.CoverTab[72883]++
											return errMixPseudoHeaderTypes
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1507
		// _ = "end of CoverTab[72883]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1508
		_go_fuzz_dep_.CoverTab[72884]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1508
		// _ = "end of CoverTab[72884]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1508
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1508
	// _ = "end of CoverTab[72872]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1508
	_go_fuzz_dep_.CoverTab[72873]++
										return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1509
	// _ = "end of CoverTab[72873]"
}

func (fr *Framer) maxHeaderStringLen() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1512
	_go_fuzz_dep_.CoverTab[72885]++
										v := fr.maxHeaderListSize()
										if uint32(int(v)) == v {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1514
		_go_fuzz_dep_.CoverTab[72887]++
											return int(v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1515
		// _ = "end of CoverTab[72887]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1516
		_go_fuzz_dep_.CoverTab[72888]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1516
		// _ = "end of CoverTab[72888]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1516
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1516
	// _ = "end of CoverTab[72885]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1516
	_go_fuzz_dep_.CoverTab[72886]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1519
	return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1519
	// _ = "end of CoverTab[72886]"
}

// readMetaFrame returns 0 or more CONTINUATION frames from fr and
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1522
// merge them into the provided hf and returns a MetaHeadersFrame
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1522
// with the decoded hpack values.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1525
func (fr *Framer) readMetaFrame(hf *HeadersFrame) (*MetaHeadersFrame, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1525
	_go_fuzz_dep_.CoverTab[72889]++
										if fr.AllowIllegalReads {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1526
		_go_fuzz_dep_.CoverTab[72897]++
											return nil, errors.New("illegal use of AllowIllegalReads with ReadMetaHeaders")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1527
		// _ = "end of CoverTab[72897]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1528
		_go_fuzz_dep_.CoverTab[72898]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1528
		// _ = "end of CoverTab[72898]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1528
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1528
	// _ = "end of CoverTab[72889]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1528
	_go_fuzz_dep_.CoverTab[72890]++
										mh := &MetaHeadersFrame{
		HeadersFrame: hf,
	}
	var remainSize = fr.maxHeaderListSize()
	var sawRegular bool

	var invalid error	// pseudo header field errors
	hdec := fr.ReadMetaHeaders
	hdec.SetEmitEnabled(true)
	hdec.SetMaxStringLength(fr.maxHeaderStringLen())
	hdec.SetEmitFunc(func(hf hpack.HeaderField) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1539
		_go_fuzz_dep_.CoverTab[72899]++
											if VerboseLogs && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1540
			_go_fuzz_dep_.CoverTab[72905]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1540
			return fr.logReads
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1540
			// _ = "end of CoverTab[72905]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1540
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1540
			_go_fuzz_dep_.CoverTab[72906]++
												fr.debugReadLoggerf("http2: decoded hpack field %+v", hf)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1541
			// _ = "end of CoverTab[72906]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1542
			_go_fuzz_dep_.CoverTab[72907]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1542
			// _ = "end of CoverTab[72907]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1542
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1542
		// _ = "end of CoverTab[72899]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1542
		_go_fuzz_dep_.CoverTab[72900]++
											if !httpguts.ValidHeaderFieldValue(hf.Value) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1543
			_go_fuzz_dep_.CoverTab[72908]++

												invalid = headerFieldValueError(hf.Name)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1545
			// _ = "end of CoverTab[72908]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1546
			_go_fuzz_dep_.CoverTab[72909]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1546
			// _ = "end of CoverTab[72909]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1546
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1546
		// _ = "end of CoverTab[72900]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1546
		_go_fuzz_dep_.CoverTab[72901]++
											isPseudo := strings.HasPrefix(hf.Name, ":")
											if isPseudo {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1548
			_go_fuzz_dep_.CoverTab[72910]++
												if sawRegular {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1549
				_go_fuzz_dep_.CoverTab[72911]++
													invalid = errPseudoAfterRegular
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1550
				// _ = "end of CoverTab[72911]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1551
				_go_fuzz_dep_.CoverTab[72912]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1551
				// _ = "end of CoverTab[72912]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1551
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1551
			// _ = "end of CoverTab[72910]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1552
			_go_fuzz_dep_.CoverTab[72913]++
												sawRegular = true
												if !validWireHeaderFieldName(hf.Name) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1554
				_go_fuzz_dep_.CoverTab[72914]++
													invalid = headerFieldNameError(hf.Name)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1555
				// _ = "end of CoverTab[72914]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1556
				_go_fuzz_dep_.CoverTab[72915]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1556
				// _ = "end of CoverTab[72915]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1556
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1556
			// _ = "end of CoverTab[72913]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1557
		// _ = "end of CoverTab[72901]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1557
		_go_fuzz_dep_.CoverTab[72902]++

											if invalid != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1559
			_go_fuzz_dep_.CoverTab[72916]++
												hdec.SetEmitEnabled(false)
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1561
			// _ = "end of CoverTab[72916]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1562
			_go_fuzz_dep_.CoverTab[72917]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1562
			// _ = "end of CoverTab[72917]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1562
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1562
		// _ = "end of CoverTab[72902]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1562
		_go_fuzz_dep_.CoverTab[72903]++

											size := hf.Size()
											if size > remainSize {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1565
			_go_fuzz_dep_.CoverTab[72918]++
												hdec.SetEmitEnabled(false)
												mh.Truncated = true
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1568
			// _ = "end of CoverTab[72918]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1569
			_go_fuzz_dep_.CoverTab[72919]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1569
			// _ = "end of CoverTab[72919]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1569
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1569
		// _ = "end of CoverTab[72903]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1569
		_go_fuzz_dep_.CoverTab[72904]++
											remainSize -= size

											mh.Fields = append(mh.Fields, hf)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1572
		// _ = "end of CoverTab[72904]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1573
	// _ = "end of CoverTab[72890]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1573
	_go_fuzz_dep_.CoverTab[72891]++

										defer hdec.SetEmitFunc(func(hf hpack.HeaderField) { _go_fuzz_dep_.CoverTab[72920]++; // _ = "end of CoverTab[72920]" })
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1575
	// _ = "end of CoverTab[72891]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1575
	_go_fuzz_dep_.CoverTab[72892]++

										var hc headersOrContinuation = hf
										for {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1578
		_go_fuzz_dep_.CoverTab[72921]++
											frag := hc.HeaderBlockFragment()
											if _, err := hdec.Write(frag); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1580
			_go_fuzz_dep_.CoverTab[72924]++
												return nil, ConnectionError(ErrCodeCompression)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1581
			// _ = "end of CoverTab[72924]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1582
			_go_fuzz_dep_.CoverTab[72925]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1582
			// _ = "end of CoverTab[72925]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1582
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1582
		// _ = "end of CoverTab[72921]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1582
		_go_fuzz_dep_.CoverTab[72922]++

											if hc.HeadersEnded() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1584
			_go_fuzz_dep_.CoverTab[72926]++
												break
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1585
			// _ = "end of CoverTab[72926]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1586
			_go_fuzz_dep_.CoverTab[72927]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1586
			// _ = "end of CoverTab[72927]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1586
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1586
		// _ = "end of CoverTab[72922]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1586
		_go_fuzz_dep_.CoverTab[72923]++
											if f, err := fr.ReadFrame(); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1587
			_go_fuzz_dep_.CoverTab[72928]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1588
			// _ = "end of CoverTab[72928]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1589
			_go_fuzz_dep_.CoverTab[72929]++
												hc = f.(*ContinuationFrame)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1590
			// _ = "end of CoverTab[72929]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1591
		// _ = "end of CoverTab[72923]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1592
	// _ = "end of CoverTab[72892]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1592
	_go_fuzz_dep_.CoverTab[72893]++

										mh.HeadersFrame.headerFragBuf = nil
										mh.HeadersFrame.invalidate()

										if err := hdec.Close(); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1597
		_go_fuzz_dep_.CoverTab[72930]++
											return nil, ConnectionError(ErrCodeCompression)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1598
		// _ = "end of CoverTab[72930]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1599
		_go_fuzz_dep_.CoverTab[72931]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1599
		// _ = "end of CoverTab[72931]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1599
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1599
	// _ = "end of CoverTab[72893]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1599
	_go_fuzz_dep_.CoverTab[72894]++
										if invalid != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1600
		_go_fuzz_dep_.CoverTab[72932]++
											fr.errDetail = invalid
											if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1602
			_go_fuzz_dep_.CoverTab[72934]++
												log.Printf("http2: invalid header: %v", invalid)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1603
			// _ = "end of CoverTab[72934]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1604
			_go_fuzz_dep_.CoverTab[72935]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1604
			// _ = "end of CoverTab[72935]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1604
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1604
		// _ = "end of CoverTab[72932]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1604
		_go_fuzz_dep_.CoverTab[72933]++
											return nil, StreamError{mh.StreamID, ErrCodeProtocol, invalid}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1605
		// _ = "end of CoverTab[72933]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1606
		_go_fuzz_dep_.CoverTab[72936]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1606
		// _ = "end of CoverTab[72936]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1606
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1606
	// _ = "end of CoverTab[72894]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1606
	_go_fuzz_dep_.CoverTab[72895]++
										if err := mh.checkPseudos(); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1607
		_go_fuzz_dep_.CoverTab[72937]++
											fr.errDetail = err
											if VerboseLogs {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1609
			_go_fuzz_dep_.CoverTab[72939]++
												log.Printf("http2: invalid pseudo headers: %v", err)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1610
			// _ = "end of CoverTab[72939]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1611
			_go_fuzz_dep_.CoverTab[72940]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1611
			// _ = "end of CoverTab[72940]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1611
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1611
		// _ = "end of CoverTab[72937]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1611
		_go_fuzz_dep_.CoverTab[72938]++
											return nil, StreamError{mh.StreamID, ErrCodeProtocol, err}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1612
		// _ = "end of CoverTab[72938]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1613
		_go_fuzz_dep_.CoverTab[72941]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1613
		// _ = "end of CoverTab[72941]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1613
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1613
	// _ = "end of CoverTab[72895]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1613
	_go_fuzz_dep_.CoverTab[72896]++
										return mh, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1614
	// _ = "end of CoverTab[72896]"
}

func summarizeFrame(f Frame) string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1617
	_go_fuzz_dep_.CoverTab[72942]++
										var buf bytes.Buffer
										f.Header().writeDebug(&buf)
										switch f := f.(type) {
	case *SettingsFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1621
		_go_fuzz_dep_.CoverTab[72944]++
											n := 0
											f.ForeachSetting(func(s Setting) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1623
			_go_fuzz_dep_.CoverTab[72953]++
												n++
												if n == 1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1625
				_go_fuzz_dep_.CoverTab[72955]++
													buf.WriteString(", settings:")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1626
				// _ = "end of CoverTab[72955]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1627
				_go_fuzz_dep_.CoverTab[72956]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1627
				// _ = "end of CoverTab[72956]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1627
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1627
			// _ = "end of CoverTab[72953]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1627
			_go_fuzz_dep_.CoverTab[72954]++
												fmt.Fprintf(&buf, " %v=%v,", s.ID, s.Val)
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1629
			// _ = "end of CoverTab[72954]"
		})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1630
		// _ = "end of CoverTab[72944]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1630
		_go_fuzz_dep_.CoverTab[72945]++
											if n > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1631
			_go_fuzz_dep_.CoverTab[72957]++
												buf.Truncate(buf.Len() - 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1632
			// _ = "end of CoverTab[72957]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1633
			_go_fuzz_dep_.CoverTab[72958]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1633
			// _ = "end of CoverTab[72958]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1633
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1633
		// _ = "end of CoverTab[72945]"
	case *DataFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1634
		_go_fuzz_dep_.CoverTab[72946]++
											data := f.Data()
											const max = 256
											if len(data) > max {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1637
			_go_fuzz_dep_.CoverTab[72959]++
												data = data[:max]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1638
			// _ = "end of CoverTab[72959]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1639
			_go_fuzz_dep_.CoverTab[72960]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1639
			// _ = "end of CoverTab[72960]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1639
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1639
		// _ = "end of CoverTab[72946]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1639
		_go_fuzz_dep_.CoverTab[72947]++
											fmt.Fprintf(&buf, " data=%q", data)
											if len(f.Data()) > max {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1641
			_go_fuzz_dep_.CoverTab[72961]++
												fmt.Fprintf(&buf, " (%d bytes omitted)", len(f.Data())-max)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1642
			// _ = "end of CoverTab[72961]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1643
			_go_fuzz_dep_.CoverTab[72962]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1643
			// _ = "end of CoverTab[72962]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1643
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1643
		// _ = "end of CoverTab[72947]"
	case *WindowUpdateFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1644
		_go_fuzz_dep_.CoverTab[72948]++
											if f.StreamID == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1645
			_go_fuzz_dep_.CoverTab[72963]++
												buf.WriteString(" (conn)")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1646
			// _ = "end of CoverTab[72963]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1647
			_go_fuzz_dep_.CoverTab[72964]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1647
			// _ = "end of CoverTab[72964]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1647
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1647
		// _ = "end of CoverTab[72948]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1647
		_go_fuzz_dep_.CoverTab[72949]++
											fmt.Fprintf(&buf, " incr=%v", f.Increment)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1648
		// _ = "end of CoverTab[72949]"
	case *PingFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1649
		_go_fuzz_dep_.CoverTab[72950]++
											fmt.Fprintf(&buf, " ping=%q", f.Data[:])
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1650
		// _ = "end of CoverTab[72950]"
	case *GoAwayFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1651
		_go_fuzz_dep_.CoverTab[72951]++
											fmt.Fprintf(&buf, " LastStreamID=%v ErrCode=%v Debug=%q",
			f.LastStreamID, f.ErrCode, f.debugData)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1653
		// _ = "end of CoverTab[72951]"
	case *RSTStreamFrame:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1654
		_go_fuzz_dep_.CoverTab[72952]++
											fmt.Fprintf(&buf, " ErrCode=%v", f.ErrCode)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1655
		// _ = "end of CoverTab[72952]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1656
	// _ = "end of CoverTab[72942]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1656
	_go_fuzz_dep_.CoverTab[72943]++
										return buf.String()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1657
	// _ = "end of CoverTab[72943]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1658
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/frame.go:1658
var _ = _go_fuzz_dep_.CoverTab
