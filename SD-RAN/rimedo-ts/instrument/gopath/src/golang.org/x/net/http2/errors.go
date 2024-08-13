// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:5
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:5
)

import (
	"errors"
	"fmt"
)

// An ErrCode is an unsigned 32-bit error code as defined in the HTTP/2 spec.
type ErrCode uint32

const (
	ErrCodeNo			ErrCode	= 0x0
	ErrCodeProtocol			ErrCode	= 0x1
	ErrCodeInternal			ErrCode	= 0x2
	ErrCodeFlowControl		ErrCode	= 0x3
	ErrCodeSettingsTimeout		ErrCode	= 0x4
	ErrCodeStreamClosed		ErrCode	= 0x5
	ErrCodeFrameSize		ErrCode	= 0x6
	ErrCodeRefusedStream		ErrCode	= 0x7
	ErrCodeCancel			ErrCode	= 0x8
	ErrCodeCompression		ErrCode	= 0x9
	ErrCodeConnect			ErrCode	= 0xa
	ErrCodeEnhanceYourCalm		ErrCode	= 0xb
	ErrCodeInadequateSecurity	ErrCode	= 0xc
	ErrCodeHTTP11Required		ErrCode	= 0xd
)

var errCodeName = map[ErrCode]string{
	ErrCodeNo:			"NO_ERROR",
	ErrCodeProtocol:		"PROTOCOL_ERROR",
	ErrCodeInternal:		"INTERNAL_ERROR",
	ErrCodeFlowControl:		"FLOW_CONTROL_ERROR",
	ErrCodeSettingsTimeout:		"SETTINGS_TIMEOUT",
	ErrCodeStreamClosed:		"STREAM_CLOSED",
	ErrCodeFrameSize:		"FRAME_SIZE_ERROR",
	ErrCodeRefusedStream:		"REFUSED_STREAM",
	ErrCodeCancel:			"CANCEL",
	ErrCodeCompression:		"COMPRESSION_ERROR",
	ErrCodeConnect:			"CONNECT_ERROR",
	ErrCodeEnhanceYourCalm:		"ENHANCE_YOUR_CALM",
	ErrCodeInadequateSecurity:	"INADEQUATE_SECURITY",
	ErrCodeHTTP11Required:		"HTTP_1_1_REQUIRED",
}

func (e ErrCode) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:49
	_go_fuzz_dep_.CoverTab[72381]++
										if s, ok := errCodeName[e]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:50
		_go_fuzz_dep_.CoverTab[72383]++
											return s
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:51
		// _ = "end of CoverTab[72383]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:52
		_go_fuzz_dep_.CoverTab[72384]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:52
		// _ = "end of CoverTab[72384]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:52
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:52
	// _ = "end of CoverTab[72381]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:52
	_go_fuzz_dep_.CoverTab[72382]++
										return fmt.Sprintf("unknown error code 0x%x", uint32(e))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:53
	// _ = "end of CoverTab[72382]"
}

func (e ErrCode) stringToken() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:56
	_go_fuzz_dep_.CoverTab[72385]++
										if s, ok := errCodeName[e]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:57
		_go_fuzz_dep_.CoverTab[72387]++
											return s
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:58
		// _ = "end of CoverTab[72387]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:59
		_go_fuzz_dep_.CoverTab[72388]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:59
		// _ = "end of CoverTab[72388]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:59
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:59
	// _ = "end of CoverTab[72385]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:59
	_go_fuzz_dep_.CoverTab[72386]++
										return fmt.Sprintf("ERR_UNKNOWN_%d", uint32(e))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:60
	// _ = "end of CoverTab[72386]"
}

// ConnectionError is an error that results in the termination of the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:63
// entire connection.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:65
type ConnectionError ErrCode

func (e ConnectionError) Error() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:67
	_go_fuzz_dep_.CoverTab[72389]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:67
	return fmt.Sprintf("connection error: %s", ErrCode(e))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:67
	// _ = "end of CoverTab[72389]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:67
}

// StreamError is an error that only affects one stream within an
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:69
// HTTP/2 connection.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:71
type StreamError struct {
	StreamID	uint32
	Code		ErrCode
	Cause		error	// optional additional detail
}

// errFromPeer is a sentinel error value for StreamError.Cause to
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:77
// indicate that the StreamError was sent from the peer over the wire
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:77
// and wasn't locally generated in the Transport.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:80
var errFromPeer = errors.New("received from peer")

func streamError(id uint32, code ErrCode) StreamError {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:82
	_go_fuzz_dep_.CoverTab[72390]++
										return StreamError{StreamID: id, Code: code}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:83
	// _ = "end of CoverTab[72390]"
}

func (e StreamError) Error() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:86
	_go_fuzz_dep_.CoverTab[72391]++
										if e.Cause != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:87
		_go_fuzz_dep_.CoverTab[72393]++
											return fmt.Sprintf("stream error: stream ID %d; %v; %v", e.StreamID, e.Code, e.Cause)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:88
		// _ = "end of CoverTab[72393]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:89
		_go_fuzz_dep_.CoverTab[72394]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:89
		// _ = "end of CoverTab[72394]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:89
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:89
	// _ = "end of CoverTab[72391]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:89
	_go_fuzz_dep_.CoverTab[72392]++
										return fmt.Sprintf("stream error: stream ID %d; %v", e.StreamID, e.Code)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:90
	// _ = "end of CoverTab[72392]"
}

// 6.9.1 The Flow Control Window
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:93
// "If a sender receives a WINDOW_UPDATE that causes a flow control
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:93
// window to exceed this maximum it MUST terminate either the stream
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:93
// or the connection, as appropriate. For streams, [...]; for the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:93
// connection, a GOAWAY frame with a FLOW_CONTROL_ERROR code."
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:98
type goAwayFlowError struct{}

func (goAwayFlowError) Error() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:100
	_go_fuzz_dep_.CoverTab[72395]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:100
	return "connection exceeded flow control window size"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:100
	// _ = "end of CoverTab[72395]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:100
}

// connError represents an HTTP/2 ConnectionError error code, along
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:102
// with a string (for debugging) explaining why.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:102
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:102
// Errors of this type are only returned by the frame parser functions
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:102
// and converted into ConnectionError(Code), after stashing away
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:102
// the Reason into the Framer's errDetail field, accessible via
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:102
// the (*Framer).ErrorDetail method.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:109
type connError struct {
	Code	ErrCode	// the ConnectionError error code
	Reason	string	// additional reason
}

func (e connError) Error() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:114
	_go_fuzz_dep_.CoverTab[72396]++
										return fmt.Sprintf("http2: connection error: %v: %v", e.Code, e.Reason)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:115
	// _ = "end of CoverTab[72396]"
}

type pseudoHeaderError string

func (e pseudoHeaderError) Error() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:120
	_go_fuzz_dep_.CoverTab[72397]++
										return fmt.Sprintf("invalid pseudo-header %q", string(e))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:121
	// _ = "end of CoverTab[72397]"
}

type duplicatePseudoHeaderError string

func (e duplicatePseudoHeaderError) Error() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:126
	_go_fuzz_dep_.CoverTab[72398]++
										return fmt.Sprintf("duplicate pseudo-header %q", string(e))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:127
	// _ = "end of CoverTab[72398]"
}

type headerFieldNameError string

func (e headerFieldNameError) Error() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:132
	_go_fuzz_dep_.CoverTab[72399]++
										return fmt.Sprintf("invalid header field name %q", string(e))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:133
	// _ = "end of CoverTab[72399]"
}

type headerFieldValueError string

func (e headerFieldValueError) Error() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:138
	_go_fuzz_dep_.CoverTab[72400]++
										return fmt.Sprintf("invalid header field value for %q", string(e))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:139
	// _ = "end of CoverTab[72400]"
}

var (
	errMixPseudoHeaderTypes	= errors.New("mix of request and response pseudo headers")
	errPseudoAfterRegular	= errors.New("pseudo header field after regular")
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:145
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/errors.go:145
var _ = _go_fuzz_dep_.CoverTab
