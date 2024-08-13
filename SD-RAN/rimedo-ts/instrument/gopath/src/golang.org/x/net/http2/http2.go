// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:5
// Package http2 implements the HTTP/2 protocol.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:5
// This package is low-level and intended to be used directly by very
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:5
// few people. Most users will use it indirectly through the automatic
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:5
// use by the net/http package (from Go 1.6 and later).
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:5
// For use in earlier Go versions see ConfigureServer. (Transport support
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:5
// requires Go 1.6 or later)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:5
// See https://http2.github.io/ for more information on HTTP/2.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:5
// See https://http2.golang.org/ for a test server running this code.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:16
package http2

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:16
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:16
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:16
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:16
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:16
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:16
)

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/net/http/httpguts"
)

var (
	VerboseLogs	bool
	logFrameWrites	bool
	logFrameReads	bool
	inTests		bool
)

func init() {
	e := os.Getenv("GODEBUG")
	if strings.Contains(e, "http2debug=1") {
		VerboseLogs = true
	}
	if strings.Contains(e, "http2debug=2") {
		VerboseLogs = true
		logFrameWrites = true
		logFrameReads = true
	}
}

const (
	// ClientPreface is the string that must be sent by new
	// connections from clients.
	ClientPreface	= "PRI * HTTP/2.0\r\n\r\nSM\r\n\r\n"

	// SETTINGS_MAX_FRAME_SIZE default
	// https://httpwg.org/specs/rfc7540.html#rfc.section.6.5.2
	initialMaxFrameSize	= 16384

	// NextProtoTLS is the NPN/ALPN protocol negotiated during
	// HTTP/2's TLS setup.
	NextProtoTLS	= "h2"

	// https://httpwg.org/specs/rfc7540.html#SettingValues
	initialHeaderTableSize	= 4096

	initialWindowSize	= 65535	// 6.9.2 Initial Flow Control Window Size

	defaultMaxReadFrameSize	= 1 << 20
)

var (
	clientPreface = []byte(ClientPreface)
)

type streamState int

// HTTP/2 stream states.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:79
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:79
// See http://tools.ietf.org/html/rfc7540#section-5.1.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:79
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:79
// For simplicity, the server code merges "reserved (local)" into
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:79
// "half-closed (remote)". This is one less state transition to track.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:79
// The only downside is that we send PUSH_PROMISEs slightly less
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:79
// liberally than allowable. More discussion here:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:79
// https://lists.w3.org/Archives/Public/ietf-http-wg/2016JulSep/0599.html
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:79
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:79
// "reserved (remote)" is omitted since the client code does not
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:79
// support server push.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:91
const (
	stateIdle	streamState	= iota
	stateOpen
	stateHalfClosedLocal
	stateHalfClosedRemote
	stateClosed
)

var stateName = [...]string{
	stateIdle:		"Idle",
	stateOpen:		"Open",
	stateHalfClosedLocal:	"HalfClosedLocal",
	stateHalfClosedRemote:	"HalfClosedRemote",
	stateClosed:		"Closed",
}

func (st streamState) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:107
	_go_fuzz_dep_.CoverTab[73057]++
										return stateName[st]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:108
	// _ = "end of CoverTab[73057]"
}

// Setting is a setting parameter: which setting it is, and its value.
type Setting struct {
	// ID is which setting is being set.
	// See https://httpwg.org/specs/rfc7540.html#SettingFormat
	ID	SettingID

	// Val is the value.
	Val	uint32
}

func (s Setting) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:121
	_go_fuzz_dep_.CoverTab[73058]++
										return fmt.Sprintf("[%v = %d]", s.ID, s.Val)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:122
	// _ = "end of CoverTab[73058]"
}

// Valid reports whether the setting is valid.
func (s Setting) Valid() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:126
	_go_fuzz_dep_.CoverTab[73059]++

										switch s.ID {
	case SettingEnablePush:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:129
		_go_fuzz_dep_.CoverTab[73061]++
											if s.Val != 1 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:130
			_go_fuzz_dep_.CoverTab[73065]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:130
			return s.Val != 0
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:130
			// _ = "end of CoverTab[73065]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:130
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:130
			_go_fuzz_dep_.CoverTab[73066]++
												return ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:131
			// _ = "end of CoverTab[73066]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:132
			_go_fuzz_dep_.CoverTab[73067]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:132
			// _ = "end of CoverTab[73067]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:132
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:132
		// _ = "end of CoverTab[73061]"
	case SettingInitialWindowSize:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:133
		_go_fuzz_dep_.CoverTab[73062]++
											if s.Val > 1<<31-1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:134
			_go_fuzz_dep_.CoverTab[73068]++
												return ConnectionError(ErrCodeFlowControl)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:135
			// _ = "end of CoverTab[73068]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:136
			_go_fuzz_dep_.CoverTab[73069]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:136
			// _ = "end of CoverTab[73069]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:136
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:136
		// _ = "end of CoverTab[73062]"
	case SettingMaxFrameSize:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:137
		_go_fuzz_dep_.CoverTab[73063]++
											if s.Val < 16384 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:138
			_go_fuzz_dep_.CoverTab[73070]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:138
			return s.Val > 1<<24-1
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:138
			// _ = "end of CoverTab[73070]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:138
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:138
			_go_fuzz_dep_.CoverTab[73071]++
												return ConnectionError(ErrCodeProtocol)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:139
			// _ = "end of CoverTab[73071]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:140
			_go_fuzz_dep_.CoverTab[73072]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:140
			// _ = "end of CoverTab[73072]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:140
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:140
		// _ = "end of CoverTab[73063]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:140
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:140
		_go_fuzz_dep_.CoverTab[73064]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:140
		// _ = "end of CoverTab[73064]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:141
	// _ = "end of CoverTab[73059]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:141
	_go_fuzz_dep_.CoverTab[73060]++
										return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:142
	// _ = "end of CoverTab[73060]"
}

// A SettingID is an HTTP/2 setting as defined in
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:145
// https://httpwg.org/specs/rfc7540.html#iana-settings
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:147
type SettingID uint16

const (
	SettingHeaderTableSize		SettingID	= 0x1
	SettingEnablePush		SettingID	= 0x2
	SettingMaxConcurrentStreams	SettingID	= 0x3
	SettingInitialWindowSize	SettingID	= 0x4
	SettingMaxFrameSize		SettingID	= 0x5
	SettingMaxHeaderListSize	SettingID	= 0x6
)

var settingName = map[SettingID]string{
	SettingHeaderTableSize:		"HEADER_TABLE_SIZE",
	SettingEnablePush:		"ENABLE_PUSH",
	SettingMaxConcurrentStreams:	"MAX_CONCURRENT_STREAMS",
	SettingInitialWindowSize:	"INITIAL_WINDOW_SIZE",
	SettingMaxFrameSize:		"MAX_FRAME_SIZE",
	SettingMaxHeaderListSize:	"MAX_HEADER_LIST_SIZE",
}

func (s SettingID) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:167
	_go_fuzz_dep_.CoverTab[73073]++
										if v, ok := settingName[s]; ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:168
		_go_fuzz_dep_.CoverTab[73075]++
											return v
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:169
		// _ = "end of CoverTab[73075]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:170
		_go_fuzz_dep_.CoverTab[73076]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:170
		// _ = "end of CoverTab[73076]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:170
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:170
	// _ = "end of CoverTab[73073]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:170
	_go_fuzz_dep_.CoverTab[73074]++
										return fmt.Sprintf("UNKNOWN_SETTING_%d", uint16(s))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:171
	// _ = "end of CoverTab[73074]"
}

// validWireHeaderFieldName reports whether v is a valid header field
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:174
// name (key). See httpguts.ValidHeaderName for the base rules.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:174
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:174
// Further, http2 says:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:174
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:174
//	"Just as in HTTP/1.x, header field names are strings of ASCII
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:174
//	characters that are compared in a case-insensitive
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:174
//	fashion. However, header field names MUST be converted to
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:174
//	lowercase prior to their encoding in HTTP/2. "
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:183
func validWireHeaderFieldName(v string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:183
	_go_fuzz_dep_.CoverTab[73077]++
										if len(v) == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:184
		_go_fuzz_dep_.CoverTab[73080]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:185
		// _ = "end of CoverTab[73080]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:186
		_go_fuzz_dep_.CoverTab[73081]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:186
		// _ = "end of CoverTab[73081]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:186
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:186
	// _ = "end of CoverTab[73077]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:186
	_go_fuzz_dep_.CoverTab[73078]++
										for _, r := range v {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:187
		_go_fuzz_dep_.CoverTab[73082]++
											if !httpguts.IsTokenRune(r) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:188
			_go_fuzz_dep_.CoverTab[73084]++
												return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:189
			// _ = "end of CoverTab[73084]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:190
			_go_fuzz_dep_.CoverTab[73085]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:190
			// _ = "end of CoverTab[73085]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:190
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:190
		// _ = "end of CoverTab[73082]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:190
		_go_fuzz_dep_.CoverTab[73083]++
											if 'A' <= r && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:191
			_go_fuzz_dep_.CoverTab[73086]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:191
			return r <= 'Z'
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:191
			// _ = "end of CoverTab[73086]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:191
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:191
			_go_fuzz_dep_.CoverTab[73087]++
												return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:192
			// _ = "end of CoverTab[73087]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:193
			_go_fuzz_dep_.CoverTab[73088]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:193
			// _ = "end of CoverTab[73088]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:193
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:193
		// _ = "end of CoverTab[73083]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:194
	// _ = "end of CoverTab[73078]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:194
	_go_fuzz_dep_.CoverTab[73079]++
										return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:195
	// _ = "end of CoverTab[73079]"
}

func httpCodeString(code int) string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:198
	_go_fuzz_dep_.CoverTab[73089]++
										switch code {
	case 200:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:200
		_go_fuzz_dep_.CoverTab[73091]++
											return "200"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:201
		// _ = "end of CoverTab[73091]"
	case 404:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:202
		_go_fuzz_dep_.CoverTab[73092]++
											return "404"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:203
		// _ = "end of CoverTab[73092]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:203
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:203
		_go_fuzz_dep_.CoverTab[73093]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:203
		// _ = "end of CoverTab[73093]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:204
	// _ = "end of CoverTab[73089]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:204
	_go_fuzz_dep_.CoverTab[73090]++
										return strconv.Itoa(code)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:205
	// _ = "end of CoverTab[73090]"
}

// from pkg io
type stringWriter interface {
	WriteString(s string) (n int, err error)
}

// A gate lets two goroutines coordinate their activities.
type gate chan struct{}

func (g gate) Done() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:216
	_go_fuzz_dep_.CoverTab[73094]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:216
	g <- struct{}{}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:216
	// _ = "end of CoverTab[73094]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:216
}
func (g gate) Wait()	{ _go_fuzz_dep_.CoverTab[73095]++; <-g; // _ = "end of CoverTab[73095]" }

// A closeWaiter is like a sync.WaitGroup but only goes 1 to 0 (open to closed).
type closeWaiter chan struct{}

// Init makes a closeWaiter usable.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:222
// It exists because so a closeWaiter value can be placed inside a
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:222
// larger struct and have the Mutex and Cond's memory in the same
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:222
// allocation.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:226
func (cw *closeWaiter) Init() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:226
	_go_fuzz_dep_.CoverTab[73096]++
										*cw = make(chan struct{})
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:227
	// _ = "end of CoverTab[73096]"
}

// Close marks the closeWaiter as closed and unblocks any waiters.
func (cw closeWaiter) Close() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:231
	_go_fuzz_dep_.CoverTab[73097]++
										close(cw)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:232
	// _ = "end of CoverTab[73097]"
}

// Wait waits for the closeWaiter to become closed.
func (cw closeWaiter) Wait() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:236
	_go_fuzz_dep_.CoverTab[73098]++
										<-cw
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:237
	// _ = "end of CoverTab[73098]"
}

// bufferedWriter is a buffered writer that writes to w.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:240
// Its buffered writer is lazily allocated as needed, to minimize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:240
// idle memory usage with many connections.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:243
type bufferedWriter struct {
	_	incomparable
	w	io.Writer	// immutable
	bw	*bufio.Writer	// non-nil when data is buffered
}

func newBufferedWriter(w io.Writer) *bufferedWriter {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:249
	_go_fuzz_dep_.CoverTab[73099]++
										return &bufferedWriter{w: w}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:250
	// _ = "end of CoverTab[73099]"
}

// bufWriterPoolBufferSize is the size of bufio.Writer's
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:253
// buffers created using bufWriterPool.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:253
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:253
// TODO: pick a less arbitrary value? this is a bit under
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:253
// (3 x typical 1500 byte MTU) at least. Other than that,
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:253
// not much thought went into it.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:259
const bufWriterPoolBufferSize = 4 << 10

var bufWriterPool = sync.Pool{
	New: func() interface{} {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:262
		_go_fuzz_dep_.CoverTab[73100]++
											return bufio.NewWriterSize(nil, bufWriterPoolBufferSize)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:263
		// _ = "end of CoverTab[73100]"
	},
}

func (w *bufferedWriter) Available() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:267
	_go_fuzz_dep_.CoverTab[73101]++
										if w.bw == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:268
		_go_fuzz_dep_.CoverTab[73103]++
											return bufWriterPoolBufferSize
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:269
		// _ = "end of CoverTab[73103]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:270
		_go_fuzz_dep_.CoverTab[73104]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:270
		// _ = "end of CoverTab[73104]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:270
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:270
	// _ = "end of CoverTab[73101]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:270
	_go_fuzz_dep_.CoverTab[73102]++
										return w.bw.Available()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:271
	// _ = "end of CoverTab[73102]"
}

func (w *bufferedWriter) Write(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:274
	_go_fuzz_dep_.CoverTab[73105]++
										if w.bw == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:275
		_go_fuzz_dep_.CoverTab[73107]++
											bw := bufWriterPool.Get().(*bufio.Writer)
											bw.Reset(w.w)
											w.bw = bw
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:278
		// _ = "end of CoverTab[73107]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:279
		_go_fuzz_dep_.CoverTab[73108]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:279
		// _ = "end of CoverTab[73108]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:279
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:279
	// _ = "end of CoverTab[73105]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:279
	_go_fuzz_dep_.CoverTab[73106]++
										return w.bw.Write(p)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:280
	// _ = "end of CoverTab[73106]"
}

func (w *bufferedWriter) Flush() error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:283
	_go_fuzz_dep_.CoverTab[73109]++
										bw := w.bw
										if bw == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:285
		_go_fuzz_dep_.CoverTab[73111]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:286
		// _ = "end of CoverTab[73111]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:287
		_go_fuzz_dep_.CoverTab[73112]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:287
		// _ = "end of CoverTab[73112]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:287
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:287
	// _ = "end of CoverTab[73109]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:287
	_go_fuzz_dep_.CoverTab[73110]++
										err := bw.Flush()
										bw.Reset(nil)
										bufWriterPool.Put(bw)
										w.bw = nil
										return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:292
	// _ = "end of CoverTab[73110]"
}

func mustUint31(v int32) uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:295
	_go_fuzz_dep_.CoverTab[73113]++
										if v < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:296
		_go_fuzz_dep_.CoverTab[73115]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:296
		return v > 2147483647
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:296
		// _ = "end of CoverTab[73115]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:296
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:296
		_go_fuzz_dep_.CoverTab[73116]++
											panic("out of range")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:297
		// _ = "end of CoverTab[73116]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:298
		_go_fuzz_dep_.CoverTab[73117]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:298
		// _ = "end of CoverTab[73117]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:298
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:298
	// _ = "end of CoverTab[73113]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:298
	_go_fuzz_dep_.CoverTab[73114]++
										return uint32(v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:299
	// _ = "end of CoverTab[73114]"
}

// bodyAllowedForStatus reports whether a given response status code
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:302
// permits a body. See RFC 7230, section 3.3.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:304
func bodyAllowedForStatus(status int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:304
	_go_fuzz_dep_.CoverTab[73118]++
										switch {
	case status >= 100 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:306
		_go_fuzz_dep_.CoverTab[73124]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:306
		return status <= 199
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:306
		// _ = "end of CoverTab[73124]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:306
	}():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:306
		_go_fuzz_dep_.CoverTab[73120]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:307
		// _ = "end of CoverTab[73120]"
	case status == 204:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:308
		_go_fuzz_dep_.CoverTab[73121]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:309
		// _ = "end of CoverTab[73121]"
	case status == 304:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:310
		_go_fuzz_dep_.CoverTab[73122]++
											return false
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:311
		// _ = "end of CoverTab[73122]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:311
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:311
		_go_fuzz_dep_.CoverTab[73123]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:311
		// _ = "end of CoverTab[73123]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:312
	// _ = "end of CoverTab[73118]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:312
	_go_fuzz_dep_.CoverTab[73119]++
										return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:313
	// _ = "end of CoverTab[73119]"
}

type httpError struct {
	_	incomparable
	msg	string
	timeout	bool
}

func (e *httpError) Error() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:322
	_go_fuzz_dep_.CoverTab[73125]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:322
	return e.msg
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:322
	// _ = "end of CoverTab[73125]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:322
}
func (e *httpError) Timeout() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:323
	_go_fuzz_dep_.CoverTab[73126]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:323
	return e.timeout
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:323
	// _ = "end of CoverTab[73126]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:323
}
func (e *httpError) Temporary() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:324
	_go_fuzz_dep_.CoverTab[73127]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:324
	return true
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:324
	// _ = "end of CoverTab[73127]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:324
}

var errTimeout error = &httpError{msg: "http2: timeout awaiting response headers", timeout: true}

type connectionStater interface {
	ConnectionState() tls.ConnectionState
}

var sorterPool = sync.Pool{New: func() interface{} {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:332
	_go_fuzz_dep_.CoverTab[73128]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:332
	return new(sorter)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:332
	// _ = "end of CoverTab[73128]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:332
}}

type sorter struct {
	v []string	// owned by sorter
}

func (s *sorter) Len() int {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:338
	_go_fuzz_dep_.CoverTab[73129]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:338
	return len(s.v)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:338
	// _ = "end of CoverTab[73129]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:338
}
func (s *sorter) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:339
	_go_fuzz_dep_.CoverTab[73130]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:339
	s.v[i], s.v[j] = s.v[j], s.v[i]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:339
	// _ = "end of CoverTab[73130]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:339
}
func (s *sorter) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:340
	_go_fuzz_dep_.CoverTab[73131]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:340
	return s.v[i] < s.v[j]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:340
	// _ = "end of CoverTab[73131]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:340
}

// Keys returns the sorted keys of h.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:342
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:342
// The returned slice is only valid until s used again or returned to
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:342
// its pool.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:346
func (s *sorter) Keys(h http.Header) []string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:346
	_go_fuzz_dep_.CoverTab[73132]++
										keys := s.v[:0]
										for k := range h {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:348
		_go_fuzz_dep_.CoverTab[73134]++
											keys = append(keys, k)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:349
		// _ = "end of CoverTab[73134]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:350
	// _ = "end of CoverTab[73132]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:350
	_go_fuzz_dep_.CoverTab[73133]++
										s.v = keys
										sort.Sort(s)
										return keys
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:353
	// _ = "end of CoverTab[73133]"
}

func (s *sorter) SortStrings(ss []string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:356
	_go_fuzz_dep_.CoverTab[73135]++

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:359
	save := s.v
										s.v = ss
										sort.Sort(s)
										s.v = save
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:362
	// _ = "end of CoverTab[73135]"
}

// validPseudoPath reports whether v is a valid :path pseudo-header
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:365
// value. It must be either:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:365
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:365
//   - a non-empty string starting with '/'
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:365
//   - the string '*', for OPTIONS requests.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:365
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:365
// For now this is only used a quick check for deciding when to clean
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:365
// up Opaque URLs before sending requests from the Transport.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:365
// See golang.org/issue/16847
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:365
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:365
// We used to enforce that the path also didn't start with "//", but
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:365
// Google's GFE accepts such paths and Chrome sends them, so ignore
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:365
// that part of the spec. See golang.org/issue/19103.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:378
func validPseudoPath(v string) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:378
	_go_fuzz_dep_.CoverTab[73136]++
										return (len(v) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:379
		_go_fuzz_dep_.CoverTab[73137]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:379
		return v[0] == '/'
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:379
		// _ = "end of CoverTab[73137]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:379
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:379
		_go_fuzz_dep_.CoverTab[73138]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:379
		return v == "*"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:379
		// _ = "end of CoverTab[73138]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:379
	}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:379
	// _ = "end of CoverTab[73136]"
}

// incomparable is a zero-width, non-comparable type. Adding it to a struct
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:382
// makes that struct also non-comparable, and generally doesn't add
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:382
// any size (as long as it's first).
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:385
type incomparable [0]func()

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:385
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/http2/http2.go:385
var _ = _go_fuzz_dep_.CoverTab
