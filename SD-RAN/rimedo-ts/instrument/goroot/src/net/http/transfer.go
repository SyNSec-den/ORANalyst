// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/http/transfer.go:5
package http

//line /usr/local/go/src/net/http/transfer.go:5
import (
//line /usr/local/go/src/net/http/transfer.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/transfer.go:5
)
//line /usr/local/go/src/net/http/transfer.go:5
import (
//line /usr/local/go/src/net/http/transfer.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/transfer.go:5
)

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http/httptrace"
	"net/http/internal"
	"net/http/internal/ascii"
	"net/textproto"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/http/httpguts"
)

// ErrLineTooLong is returned when reading request or response bodies
//line /usr/local/go/src/net/http/transfer.go:27
// with malformed chunked encoding.
//line /usr/local/go/src/net/http/transfer.go:29
var ErrLineTooLong = internal.ErrLineTooLong

type errorReader struct {
	err error
}

func (r errorReader) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/transfer.go:35
	_go_fuzz_dep_.CoverTab[43552]++
							return 0, r.err
//line /usr/local/go/src/net/http/transfer.go:36
	// _ = "end of CoverTab[43552]"
}

type byteReader struct {
	b	byte
	done	bool
}

func (br *byteReader) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/transfer.go:44
	_go_fuzz_dep_.CoverTab[43553]++
							if br.done {
//line /usr/local/go/src/net/http/transfer.go:45
		_go_fuzz_dep_.CoverTab[43556]++
								return 0, io.EOF
//line /usr/local/go/src/net/http/transfer.go:46
		// _ = "end of CoverTab[43556]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:47
		_go_fuzz_dep_.CoverTab[43557]++
//line /usr/local/go/src/net/http/transfer.go:47
		// _ = "end of CoverTab[43557]"
//line /usr/local/go/src/net/http/transfer.go:47
	}
//line /usr/local/go/src/net/http/transfer.go:47
	// _ = "end of CoverTab[43553]"
//line /usr/local/go/src/net/http/transfer.go:47
	_go_fuzz_dep_.CoverTab[43554]++
							if len(p) == 0 {
//line /usr/local/go/src/net/http/transfer.go:48
		_go_fuzz_dep_.CoverTab[43558]++
								return 0, nil
//line /usr/local/go/src/net/http/transfer.go:49
		// _ = "end of CoverTab[43558]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:50
		_go_fuzz_dep_.CoverTab[43559]++
//line /usr/local/go/src/net/http/transfer.go:50
		// _ = "end of CoverTab[43559]"
//line /usr/local/go/src/net/http/transfer.go:50
	}
//line /usr/local/go/src/net/http/transfer.go:50
	// _ = "end of CoverTab[43554]"
//line /usr/local/go/src/net/http/transfer.go:50
	_go_fuzz_dep_.CoverTab[43555]++
							br.done = true
							p[0] = br.b
							return 1, io.EOF
//line /usr/local/go/src/net/http/transfer.go:53
	// _ = "end of CoverTab[43555]"
}

// transferWriter inspects the fields of a user-supplied Request or Response,
//line /usr/local/go/src/net/http/transfer.go:56
// sanitizes them without changing the user object and provides methods for
//line /usr/local/go/src/net/http/transfer.go:56
// writing the respective header, body and trailer in wire format.
//line /usr/local/go/src/net/http/transfer.go:59
type transferWriter struct {
	Method			string
	Body			io.Reader
	BodyCloser		io.Closer
	ResponseToHEAD		bool
	ContentLength		int64	// -1 means unknown, 0 means exactly none
	Close			bool
	TransferEncoding	[]string
	Header			Header
	Trailer			Header
	IsResponse		bool
	bodyReadError		error	// any non-EOF error from reading Body

	FlushHeaders	bool		// flush headers to network before body
	ByteReadCh	chan readResult	// non-nil if probeRequestBody called
}

func newTransferWriter(r any) (t *transferWriter, err error) {
//line /usr/local/go/src/net/http/transfer.go:76
	_go_fuzz_dep_.CoverTab[43560]++
							t = &transferWriter{}

//line /usr/local/go/src/net/http/transfer.go:80
	atLeastHTTP11 := false
	switch rr := r.(type) {
	case *Request:
//line /usr/local/go/src/net/http/transfer.go:82
		_go_fuzz_dep_.CoverTab[43564]++
								if rr.ContentLength != 0 && func() bool {
//line /usr/local/go/src/net/http/transfer.go:83
			_go_fuzz_dep_.CoverTab[43570]++
//line /usr/local/go/src/net/http/transfer.go:83
			return rr.Body == nil
//line /usr/local/go/src/net/http/transfer.go:83
			// _ = "end of CoverTab[43570]"
//line /usr/local/go/src/net/http/transfer.go:83
		}() {
//line /usr/local/go/src/net/http/transfer.go:83
			_go_fuzz_dep_.CoverTab[43571]++
									return nil, fmt.Errorf("http: Request.ContentLength=%d with nil Body", rr.ContentLength)
//line /usr/local/go/src/net/http/transfer.go:84
			// _ = "end of CoverTab[43571]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:85
			_go_fuzz_dep_.CoverTab[43572]++
//line /usr/local/go/src/net/http/transfer.go:85
			// _ = "end of CoverTab[43572]"
//line /usr/local/go/src/net/http/transfer.go:85
		}
//line /usr/local/go/src/net/http/transfer.go:85
		// _ = "end of CoverTab[43564]"
//line /usr/local/go/src/net/http/transfer.go:85
		_go_fuzz_dep_.CoverTab[43565]++
								t.Method = valueOrDefault(rr.Method, "GET")
								t.Close = rr.Close
								t.TransferEncoding = rr.TransferEncoding
								t.Header = rr.Header
								t.Trailer = rr.Trailer
								t.Body = rr.Body
								t.BodyCloser = rr.Body
								t.ContentLength = rr.outgoingLength()
								if t.ContentLength < 0 && func() bool {
//line /usr/local/go/src/net/http/transfer.go:94
			_go_fuzz_dep_.CoverTab[43573]++
//line /usr/local/go/src/net/http/transfer.go:94
			return len(t.TransferEncoding) == 0
//line /usr/local/go/src/net/http/transfer.go:94
			// _ = "end of CoverTab[43573]"
//line /usr/local/go/src/net/http/transfer.go:94
		}() && func() bool {
//line /usr/local/go/src/net/http/transfer.go:94
			_go_fuzz_dep_.CoverTab[43574]++
//line /usr/local/go/src/net/http/transfer.go:94
			return t.shouldSendChunkedRequestBody()
//line /usr/local/go/src/net/http/transfer.go:94
			// _ = "end of CoverTab[43574]"
//line /usr/local/go/src/net/http/transfer.go:94
		}() {
//line /usr/local/go/src/net/http/transfer.go:94
			_go_fuzz_dep_.CoverTab[43575]++
									t.TransferEncoding = []string{"chunked"}
//line /usr/local/go/src/net/http/transfer.go:95
			// _ = "end of CoverTab[43575]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:96
			_go_fuzz_dep_.CoverTab[43576]++
//line /usr/local/go/src/net/http/transfer.go:96
			// _ = "end of CoverTab[43576]"
//line /usr/local/go/src/net/http/transfer.go:96
		}
//line /usr/local/go/src/net/http/transfer.go:96
		// _ = "end of CoverTab[43565]"
//line /usr/local/go/src/net/http/transfer.go:96
		_go_fuzz_dep_.CoverTab[43566]++

//line /usr/local/go/src/net/http/transfer.go:104
		if t.ContentLength != 0 && func() bool {
//line /usr/local/go/src/net/http/transfer.go:104
			_go_fuzz_dep_.CoverTab[43577]++
//line /usr/local/go/src/net/http/transfer.go:104
			return !isKnownInMemoryReader(t.Body)
//line /usr/local/go/src/net/http/transfer.go:104
			// _ = "end of CoverTab[43577]"
//line /usr/local/go/src/net/http/transfer.go:104
		}() {
//line /usr/local/go/src/net/http/transfer.go:104
			_go_fuzz_dep_.CoverTab[43578]++
									t.FlushHeaders = true
//line /usr/local/go/src/net/http/transfer.go:105
			// _ = "end of CoverTab[43578]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:106
			_go_fuzz_dep_.CoverTab[43579]++
//line /usr/local/go/src/net/http/transfer.go:106
			// _ = "end of CoverTab[43579]"
//line /usr/local/go/src/net/http/transfer.go:106
		}
//line /usr/local/go/src/net/http/transfer.go:106
		// _ = "end of CoverTab[43566]"
//line /usr/local/go/src/net/http/transfer.go:106
		_go_fuzz_dep_.CoverTab[43567]++

								atLeastHTTP11 = true
//line /usr/local/go/src/net/http/transfer.go:108
		// _ = "end of CoverTab[43567]"
	case *Response:
//line /usr/local/go/src/net/http/transfer.go:109
		_go_fuzz_dep_.CoverTab[43568]++
								t.IsResponse = true
								if rr.Request != nil {
//line /usr/local/go/src/net/http/transfer.go:111
			_go_fuzz_dep_.CoverTab[43580]++
									t.Method = rr.Request.Method
//line /usr/local/go/src/net/http/transfer.go:112
			// _ = "end of CoverTab[43580]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:113
			_go_fuzz_dep_.CoverTab[43581]++
//line /usr/local/go/src/net/http/transfer.go:113
			// _ = "end of CoverTab[43581]"
//line /usr/local/go/src/net/http/transfer.go:113
		}
//line /usr/local/go/src/net/http/transfer.go:113
		// _ = "end of CoverTab[43568]"
//line /usr/local/go/src/net/http/transfer.go:113
		_go_fuzz_dep_.CoverTab[43569]++
								t.Body = rr.Body
								t.BodyCloser = rr.Body
								t.ContentLength = rr.ContentLength
								t.Close = rr.Close
								t.TransferEncoding = rr.TransferEncoding
								t.Header = rr.Header
								t.Trailer = rr.Trailer
								atLeastHTTP11 = rr.ProtoAtLeast(1, 1)
								t.ResponseToHEAD = noResponseBodyExpected(t.Method)
//line /usr/local/go/src/net/http/transfer.go:122
		// _ = "end of CoverTab[43569]"
	}
//line /usr/local/go/src/net/http/transfer.go:123
	// _ = "end of CoverTab[43560]"
//line /usr/local/go/src/net/http/transfer.go:123
	_go_fuzz_dep_.CoverTab[43561]++

//line /usr/local/go/src/net/http/transfer.go:126
	if t.ResponseToHEAD {
//line /usr/local/go/src/net/http/transfer.go:126
		_go_fuzz_dep_.CoverTab[43582]++
								t.Body = nil
								if chunked(t.TransferEncoding) {
//line /usr/local/go/src/net/http/transfer.go:128
			_go_fuzz_dep_.CoverTab[43583]++
									t.ContentLength = -1
//line /usr/local/go/src/net/http/transfer.go:129
			// _ = "end of CoverTab[43583]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:130
			_go_fuzz_dep_.CoverTab[43584]++
//line /usr/local/go/src/net/http/transfer.go:130
			// _ = "end of CoverTab[43584]"
//line /usr/local/go/src/net/http/transfer.go:130
		}
//line /usr/local/go/src/net/http/transfer.go:130
		// _ = "end of CoverTab[43582]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:131
		_go_fuzz_dep_.CoverTab[43585]++
								if !atLeastHTTP11 || func() bool {
//line /usr/local/go/src/net/http/transfer.go:132
			_go_fuzz_dep_.CoverTab[43587]++
//line /usr/local/go/src/net/http/transfer.go:132
			return t.Body == nil
//line /usr/local/go/src/net/http/transfer.go:132
			// _ = "end of CoverTab[43587]"
//line /usr/local/go/src/net/http/transfer.go:132
		}() {
//line /usr/local/go/src/net/http/transfer.go:132
			_go_fuzz_dep_.CoverTab[43588]++
									t.TransferEncoding = nil
//line /usr/local/go/src/net/http/transfer.go:133
			// _ = "end of CoverTab[43588]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:134
			_go_fuzz_dep_.CoverTab[43589]++
//line /usr/local/go/src/net/http/transfer.go:134
			// _ = "end of CoverTab[43589]"
//line /usr/local/go/src/net/http/transfer.go:134
		}
//line /usr/local/go/src/net/http/transfer.go:134
		// _ = "end of CoverTab[43585]"
//line /usr/local/go/src/net/http/transfer.go:134
		_go_fuzz_dep_.CoverTab[43586]++
								if chunked(t.TransferEncoding) {
//line /usr/local/go/src/net/http/transfer.go:135
			_go_fuzz_dep_.CoverTab[43590]++
									t.ContentLength = -1
//line /usr/local/go/src/net/http/transfer.go:136
			// _ = "end of CoverTab[43590]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:137
			_go_fuzz_dep_.CoverTab[43591]++
//line /usr/local/go/src/net/http/transfer.go:137
			if t.Body == nil {
//line /usr/local/go/src/net/http/transfer.go:137
				_go_fuzz_dep_.CoverTab[43592]++
										t.ContentLength = 0
//line /usr/local/go/src/net/http/transfer.go:138
				// _ = "end of CoverTab[43592]"
			} else {
//line /usr/local/go/src/net/http/transfer.go:139
				_go_fuzz_dep_.CoverTab[43593]++
//line /usr/local/go/src/net/http/transfer.go:139
				// _ = "end of CoverTab[43593]"
//line /usr/local/go/src/net/http/transfer.go:139
			}
//line /usr/local/go/src/net/http/transfer.go:139
			// _ = "end of CoverTab[43591]"
//line /usr/local/go/src/net/http/transfer.go:139
		}
//line /usr/local/go/src/net/http/transfer.go:139
		// _ = "end of CoverTab[43586]"
	}
//line /usr/local/go/src/net/http/transfer.go:140
	// _ = "end of CoverTab[43561]"
//line /usr/local/go/src/net/http/transfer.go:140
	_go_fuzz_dep_.CoverTab[43562]++

//line /usr/local/go/src/net/http/transfer.go:143
	if !chunked(t.TransferEncoding) {
//line /usr/local/go/src/net/http/transfer.go:143
		_go_fuzz_dep_.CoverTab[43594]++
								t.Trailer = nil
//line /usr/local/go/src/net/http/transfer.go:144
		// _ = "end of CoverTab[43594]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:145
		_go_fuzz_dep_.CoverTab[43595]++
//line /usr/local/go/src/net/http/transfer.go:145
		// _ = "end of CoverTab[43595]"
//line /usr/local/go/src/net/http/transfer.go:145
	}
//line /usr/local/go/src/net/http/transfer.go:145
	// _ = "end of CoverTab[43562]"
//line /usr/local/go/src/net/http/transfer.go:145
	_go_fuzz_dep_.CoverTab[43563]++

							return t, nil
//line /usr/local/go/src/net/http/transfer.go:147
	// _ = "end of CoverTab[43563]"
}

// shouldSendChunkedRequestBody reports whether we should try to send a
//line /usr/local/go/src/net/http/transfer.go:150
// chunked request body to the server. In particular, the case we really
//line /usr/local/go/src/net/http/transfer.go:150
// want to prevent is sending a GET or other typically-bodyless request to a
//line /usr/local/go/src/net/http/transfer.go:150
// server with a chunked body when the body has zero bytes, since GETs with
//line /usr/local/go/src/net/http/transfer.go:150
// bodies (while acceptable according to specs), even zero-byte chunked
//line /usr/local/go/src/net/http/transfer.go:150
// bodies, are approximately never seen in the wild and confuse most
//line /usr/local/go/src/net/http/transfer.go:150
// servers. See Issue 18257, as one example.
//line /usr/local/go/src/net/http/transfer.go:150
//
//line /usr/local/go/src/net/http/transfer.go:150
// The only reason we'd send such a request is if the user set the Body to a
//line /usr/local/go/src/net/http/transfer.go:150
// non-nil value (say, io.NopCloser(bytes.NewReader(nil))) and didn't
//line /usr/local/go/src/net/http/transfer.go:150
// set ContentLength, or NewRequest set it to -1 (unknown), so then we assume
//line /usr/local/go/src/net/http/transfer.go:150
// there's bytes to send.
//line /usr/local/go/src/net/http/transfer.go:150
//
//line /usr/local/go/src/net/http/transfer.go:150
// This code tries to read a byte from the Request.Body in such cases to see
//line /usr/local/go/src/net/http/transfer.go:150
// whether the body actually has content (super rare) or is actually just
//line /usr/local/go/src/net/http/transfer.go:150
// a non-nil content-less ReadCloser (the more common case). In that more
//line /usr/local/go/src/net/http/transfer.go:150
// common case, we act as if their Body were nil instead, and don't send
//line /usr/local/go/src/net/http/transfer.go:150
// a body.
//line /usr/local/go/src/net/http/transfer.go:168
func (t *transferWriter) shouldSendChunkedRequestBody() bool {
//line /usr/local/go/src/net/http/transfer.go:168
	_go_fuzz_dep_.CoverTab[43596]++

//line /usr/local/go/src/net/http/transfer.go:171
	if t.ContentLength >= 0 || func() bool {
//line /usr/local/go/src/net/http/transfer.go:171
		_go_fuzz_dep_.CoverTab[43600]++
//line /usr/local/go/src/net/http/transfer.go:171
		return t.Body == nil
//line /usr/local/go/src/net/http/transfer.go:171
		// _ = "end of CoverTab[43600]"
//line /usr/local/go/src/net/http/transfer.go:171
	}() {
//line /usr/local/go/src/net/http/transfer.go:171
		_go_fuzz_dep_.CoverTab[43601]++
								return false
//line /usr/local/go/src/net/http/transfer.go:172
		// _ = "end of CoverTab[43601]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:173
		_go_fuzz_dep_.CoverTab[43602]++
//line /usr/local/go/src/net/http/transfer.go:173
		// _ = "end of CoverTab[43602]"
//line /usr/local/go/src/net/http/transfer.go:173
	}
//line /usr/local/go/src/net/http/transfer.go:173
	// _ = "end of CoverTab[43596]"
//line /usr/local/go/src/net/http/transfer.go:173
	_go_fuzz_dep_.CoverTab[43597]++
							if t.Method == "CONNECT" {
//line /usr/local/go/src/net/http/transfer.go:174
		_go_fuzz_dep_.CoverTab[43603]++
								return false
//line /usr/local/go/src/net/http/transfer.go:175
		// _ = "end of CoverTab[43603]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:176
		_go_fuzz_dep_.CoverTab[43604]++
//line /usr/local/go/src/net/http/transfer.go:176
		// _ = "end of CoverTab[43604]"
//line /usr/local/go/src/net/http/transfer.go:176
	}
//line /usr/local/go/src/net/http/transfer.go:176
	// _ = "end of CoverTab[43597]"
//line /usr/local/go/src/net/http/transfer.go:176
	_go_fuzz_dep_.CoverTab[43598]++
							if requestMethodUsuallyLacksBody(t.Method) {
//line /usr/local/go/src/net/http/transfer.go:177
		_go_fuzz_dep_.CoverTab[43605]++

//line /usr/local/go/src/net/http/transfer.go:181
		t.probeRequestBody()
								return t.Body != nil
//line /usr/local/go/src/net/http/transfer.go:182
		// _ = "end of CoverTab[43605]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:183
		_go_fuzz_dep_.CoverTab[43606]++
//line /usr/local/go/src/net/http/transfer.go:183
		// _ = "end of CoverTab[43606]"
//line /usr/local/go/src/net/http/transfer.go:183
	}
//line /usr/local/go/src/net/http/transfer.go:183
	// _ = "end of CoverTab[43598]"
//line /usr/local/go/src/net/http/transfer.go:183
	_go_fuzz_dep_.CoverTab[43599]++

//line /usr/local/go/src/net/http/transfer.go:188
	return true
//line /usr/local/go/src/net/http/transfer.go:188
	// _ = "end of CoverTab[43599]"
}

// probeRequestBody reads a byte from t.Body to see whether it's empty
//line /usr/local/go/src/net/http/transfer.go:191
// (returns io.EOF right away).
//line /usr/local/go/src/net/http/transfer.go:191
//
//line /usr/local/go/src/net/http/transfer.go:191
// But because we've had problems with this blocking users in the past
//line /usr/local/go/src/net/http/transfer.go:191
// (issue 17480) when the body is a pipe (perhaps waiting on the response
//line /usr/local/go/src/net/http/transfer.go:191
// headers before the pipe is fed data), we need to be careful and bound how
//line /usr/local/go/src/net/http/transfer.go:191
// long we wait for it. This delay will only affect users if all the following
//line /usr/local/go/src/net/http/transfer.go:191
// are true:
//line /usr/local/go/src/net/http/transfer.go:191
//   - the request body blocks
//line /usr/local/go/src/net/http/transfer.go:191
//   - the content length is not set (or set to -1)
//line /usr/local/go/src/net/http/transfer.go:191
//   - the method doesn't usually have a body (GET, HEAD, DELETE, ...)
//line /usr/local/go/src/net/http/transfer.go:191
//   - there is no transfer-encoding=chunked already set.
//line /usr/local/go/src/net/http/transfer.go:191
//
//line /usr/local/go/src/net/http/transfer.go:191
// In other words, this delay will not normally affect anybody, and there
//line /usr/local/go/src/net/http/transfer.go:191
// are workarounds if it does.
//line /usr/local/go/src/net/http/transfer.go:206
func (t *transferWriter) probeRequestBody() {
//line /usr/local/go/src/net/http/transfer.go:206
	_go_fuzz_dep_.CoverTab[43607]++
							t.ByteReadCh = make(chan readResult, 1)
//line /usr/local/go/src/net/http/transfer.go:207
	_curRoutineNum36_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/transfer.go:207
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum36_)
							go func(body io.Reader) {
//line /usr/local/go/src/net/http/transfer.go:208
		_go_fuzz_dep_.CoverTab[43609]++
//line /usr/local/go/src/net/http/transfer.go:208
		defer func() {
//line /usr/local/go/src/net/http/transfer.go:208
			_go_fuzz_dep_.CoverTab[43611]++
//line /usr/local/go/src/net/http/transfer.go:208
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum36_)
//line /usr/local/go/src/net/http/transfer.go:208
			// _ = "end of CoverTab[43611]"
//line /usr/local/go/src/net/http/transfer.go:208
		}()
								var buf [1]byte
								var rres readResult
								rres.n, rres.err = body.Read(buf[:])
								if rres.n == 1 {
//line /usr/local/go/src/net/http/transfer.go:212
			_go_fuzz_dep_.CoverTab[43612]++
									rres.b = buf[0]
//line /usr/local/go/src/net/http/transfer.go:213
			// _ = "end of CoverTab[43612]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:214
			_go_fuzz_dep_.CoverTab[43613]++
//line /usr/local/go/src/net/http/transfer.go:214
			// _ = "end of CoverTab[43613]"
//line /usr/local/go/src/net/http/transfer.go:214
		}
//line /usr/local/go/src/net/http/transfer.go:214
		// _ = "end of CoverTab[43609]"
//line /usr/local/go/src/net/http/transfer.go:214
		_go_fuzz_dep_.CoverTab[43610]++
								t.ByteReadCh <- rres
								close(t.ByteReadCh)
//line /usr/local/go/src/net/http/transfer.go:216
		// _ = "end of CoverTab[43610]"
	}(t.Body)
//line /usr/local/go/src/net/http/transfer.go:217
	// _ = "end of CoverTab[43607]"
//line /usr/local/go/src/net/http/transfer.go:217
	_go_fuzz_dep_.CoverTab[43608]++
							timer := time.NewTimer(200 * time.Millisecond)
							select {
	case rres := <-t.ByteReadCh:
//line /usr/local/go/src/net/http/transfer.go:220
		_go_fuzz_dep_.CoverTab[43614]++
								timer.Stop()
								if rres.n == 0 && func() bool {
//line /usr/local/go/src/net/http/transfer.go:222
			_go_fuzz_dep_.CoverTab[43616]++
//line /usr/local/go/src/net/http/transfer.go:222
			return rres.err == io.EOF
//line /usr/local/go/src/net/http/transfer.go:222
			// _ = "end of CoverTab[43616]"
//line /usr/local/go/src/net/http/transfer.go:222
		}() {
//line /usr/local/go/src/net/http/transfer.go:222
			_go_fuzz_dep_.CoverTab[43617]++

									t.Body = nil
									t.ContentLength = 0
//line /usr/local/go/src/net/http/transfer.go:225
			// _ = "end of CoverTab[43617]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:226
			_go_fuzz_dep_.CoverTab[43618]++
//line /usr/local/go/src/net/http/transfer.go:226
			if rres.n == 1 {
//line /usr/local/go/src/net/http/transfer.go:226
				_go_fuzz_dep_.CoverTab[43619]++
										if rres.err != nil {
//line /usr/local/go/src/net/http/transfer.go:227
					_go_fuzz_dep_.CoverTab[43620]++
											t.Body = io.MultiReader(&byteReader{b: rres.b}, errorReader{rres.err})
//line /usr/local/go/src/net/http/transfer.go:228
					// _ = "end of CoverTab[43620]"
				} else {
//line /usr/local/go/src/net/http/transfer.go:229
					_go_fuzz_dep_.CoverTab[43621]++
											t.Body = io.MultiReader(&byteReader{b: rres.b}, t.Body)
//line /usr/local/go/src/net/http/transfer.go:230
					// _ = "end of CoverTab[43621]"
				}
//line /usr/local/go/src/net/http/transfer.go:231
				// _ = "end of CoverTab[43619]"
			} else {
//line /usr/local/go/src/net/http/transfer.go:232
				_go_fuzz_dep_.CoverTab[43622]++
//line /usr/local/go/src/net/http/transfer.go:232
				if rres.err != nil {
//line /usr/local/go/src/net/http/transfer.go:232
					_go_fuzz_dep_.CoverTab[43623]++
											t.Body = errorReader{rres.err}
//line /usr/local/go/src/net/http/transfer.go:233
					// _ = "end of CoverTab[43623]"
				} else {
//line /usr/local/go/src/net/http/transfer.go:234
					_go_fuzz_dep_.CoverTab[43624]++
//line /usr/local/go/src/net/http/transfer.go:234
					// _ = "end of CoverTab[43624]"
//line /usr/local/go/src/net/http/transfer.go:234
				}
//line /usr/local/go/src/net/http/transfer.go:234
				// _ = "end of CoverTab[43622]"
//line /usr/local/go/src/net/http/transfer.go:234
			}
//line /usr/local/go/src/net/http/transfer.go:234
			// _ = "end of CoverTab[43618]"
//line /usr/local/go/src/net/http/transfer.go:234
		}
//line /usr/local/go/src/net/http/transfer.go:234
		// _ = "end of CoverTab[43614]"
	case <-timer.C:
//line /usr/local/go/src/net/http/transfer.go:235
		_go_fuzz_dep_.CoverTab[43615]++

//line /usr/local/go/src/net/http/transfer.go:240
		t.Body = io.MultiReader(finishAsyncByteRead{t}, t.Body)

//line /usr/local/go/src/net/http/transfer.go:244
		t.FlushHeaders = true
//line /usr/local/go/src/net/http/transfer.go:244
		// _ = "end of CoverTab[43615]"
	}
//line /usr/local/go/src/net/http/transfer.go:245
	// _ = "end of CoverTab[43608]"
}

func noResponseBodyExpected(requestMethod string) bool {
//line /usr/local/go/src/net/http/transfer.go:248
	_go_fuzz_dep_.CoverTab[43625]++
							return requestMethod == "HEAD"
//line /usr/local/go/src/net/http/transfer.go:249
	// _ = "end of CoverTab[43625]"
}

func (t *transferWriter) shouldSendContentLength() bool {
//line /usr/local/go/src/net/http/transfer.go:252
	_go_fuzz_dep_.CoverTab[43626]++
							if chunked(t.TransferEncoding) {
//line /usr/local/go/src/net/http/transfer.go:253
		_go_fuzz_dep_.CoverTab[43632]++
								return false
//line /usr/local/go/src/net/http/transfer.go:254
		// _ = "end of CoverTab[43632]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:255
		_go_fuzz_dep_.CoverTab[43633]++
//line /usr/local/go/src/net/http/transfer.go:255
		// _ = "end of CoverTab[43633]"
//line /usr/local/go/src/net/http/transfer.go:255
	}
//line /usr/local/go/src/net/http/transfer.go:255
	// _ = "end of CoverTab[43626]"
//line /usr/local/go/src/net/http/transfer.go:255
	_go_fuzz_dep_.CoverTab[43627]++
							if t.ContentLength > 0 {
//line /usr/local/go/src/net/http/transfer.go:256
		_go_fuzz_dep_.CoverTab[43634]++
								return true
//line /usr/local/go/src/net/http/transfer.go:257
		// _ = "end of CoverTab[43634]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:258
		_go_fuzz_dep_.CoverTab[43635]++
//line /usr/local/go/src/net/http/transfer.go:258
		// _ = "end of CoverTab[43635]"
//line /usr/local/go/src/net/http/transfer.go:258
	}
//line /usr/local/go/src/net/http/transfer.go:258
	// _ = "end of CoverTab[43627]"
//line /usr/local/go/src/net/http/transfer.go:258
	_go_fuzz_dep_.CoverTab[43628]++
							if t.ContentLength < 0 {
//line /usr/local/go/src/net/http/transfer.go:259
		_go_fuzz_dep_.CoverTab[43636]++
								return false
//line /usr/local/go/src/net/http/transfer.go:260
		// _ = "end of CoverTab[43636]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:261
		_go_fuzz_dep_.CoverTab[43637]++
//line /usr/local/go/src/net/http/transfer.go:261
		// _ = "end of CoverTab[43637]"
//line /usr/local/go/src/net/http/transfer.go:261
	}
//line /usr/local/go/src/net/http/transfer.go:261
	// _ = "end of CoverTab[43628]"
//line /usr/local/go/src/net/http/transfer.go:261
	_go_fuzz_dep_.CoverTab[43629]++

							if t.Method == "POST" || func() bool {
//line /usr/local/go/src/net/http/transfer.go:263
		_go_fuzz_dep_.CoverTab[43638]++
//line /usr/local/go/src/net/http/transfer.go:263
		return t.Method == "PUT"
//line /usr/local/go/src/net/http/transfer.go:263
		// _ = "end of CoverTab[43638]"
//line /usr/local/go/src/net/http/transfer.go:263
	}() || func() bool {
//line /usr/local/go/src/net/http/transfer.go:263
		_go_fuzz_dep_.CoverTab[43639]++
//line /usr/local/go/src/net/http/transfer.go:263
		return t.Method == "PATCH"
//line /usr/local/go/src/net/http/transfer.go:263
		// _ = "end of CoverTab[43639]"
//line /usr/local/go/src/net/http/transfer.go:263
	}() {
//line /usr/local/go/src/net/http/transfer.go:263
		_go_fuzz_dep_.CoverTab[43640]++
								return true
//line /usr/local/go/src/net/http/transfer.go:264
		// _ = "end of CoverTab[43640]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:265
		_go_fuzz_dep_.CoverTab[43641]++
//line /usr/local/go/src/net/http/transfer.go:265
		// _ = "end of CoverTab[43641]"
//line /usr/local/go/src/net/http/transfer.go:265
	}
//line /usr/local/go/src/net/http/transfer.go:265
	// _ = "end of CoverTab[43629]"
//line /usr/local/go/src/net/http/transfer.go:265
	_go_fuzz_dep_.CoverTab[43630]++
							if t.ContentLength == 0 && func() bool {
//line /usr/local/go/src/net/http/transfer.go:266
		_go_fuzz_dep_.CoverTab[43642]++
//line /usr/local/go/src/net/http/transfer.go:266
		return isIdentity(t.TransferEncoding)
//line /usr/local/go/src/net/http/transfer.go:266
		// _ = "end of CoverTab[43642]"
//line /usr/local/go/src/net/http/transfer.go:266
	}() {
//line /usr/local/go/src/net/http/transfer.go:266
		_go_fuzz_dep_.CoverTab[43643]++
								if t.Method == "GET" || func() bool {
//line /usr/local/go/src/net/http/transfer.go:267
			_go_fuzz_dep_.CoverTab[43645]++
//line /usr/local/go/src/net/http/transfer.go:267
			return t.Method == "HEAD"
//line /usr/local/go/src/net/http/transfer.go:267
			// _ = "end of CoverTab[43645]"
//line /usr/local/go/src/net/http/transfer.go:267
		}() {
//line /usr/local/go/src/net/http/transfer.go:267
			_go_fuzz_dep_.CoverTab[43646]++
									return false
//line /usr/local/go/src/net/http/transfer.go:268
			// _ = "end of CoverTab[43646]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:269
			_go_fuzz_dep_.CoverTab[43647]++
//line /usr/local/go/src/net/http/transfer.go:269
			// _ = "end of CoverTab[43647]"
//line /usr/local/go/src/net/http/transfer.go:269
		}
//line /usr/local/go/src/net/http/transfer.go:269
		// _ = "end of CoverTab[43643]"
//line /usr/local/go/src/net/http/transfer.go:269
		_go_fuzz_dep_.CoverTab[43644]++
								return true
//line /usr/local/go/src/net/http/transfer.go:270
		// _ = "end of CoverTab[43644]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:271
		_go_fuzz_dep_.CoverTab[43648]++
//line /usr/local/go/src/net/http/transfer.go:271
		// _ = "end of CoverTab[43648]"
//line /usr/local/go/src/net/http/transfer.go:271
	}
//line /usr/local/go/src/net/http/transfer.go:271
	// _ = "end of CoverTab[43630]"
//line /usr/local/go/src/net/http/transfer.go:271
	_go_fuzz_dep_.CoverTab[43631]++

							return false
//line /usr/local/go/src/net/http/transfer.go:273
	// _ = "end of CoverTab[43631]"
}

func (t *transferWriter) writeHeader(w io.Writer, trace *httptrace.ClientTrace) error {
//line /usr/local/go/src/net/http/transfer.go:276
	_go_fuzz_dep_.CoverTab[43649]++
							if t.Close && func() bool {
//line /usr/local/go/src/net/http/transfer.go:277
		_go_fuzz_dep_.CoverTab[43653]++
//line /usr/local/go/src/net/http/transfer.go:277
		return !hasToken(t.Header.get("Connection"), "close")
//line /usr/local/go/src/net/http/transfer.go:277
		// _ = "end of CoverTab[43653]"
//line /usr/local/go/src/net/http/transfer.go:277
	}() {
//line /usr/local/go/src/net/http/transfer.go:277
		_go_fuzz_dep_.CoverTab[43654]++
								if _, err := io.WriteString(w, "Connection: close\r\n"); err != nil {
//line /usr/local/go/src/net/http/transfer.go:278
			_go_fuzz_dep_.CoverTab[43656]++
									return err
//line /usr/local/go/src/net/http/transfer.go:279
			// _ = "end of CoverTab[43656]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:280
			_go_fuzz_dep_.CoverTab[43657]++
//line /usr/local/go/src/net/http/transfer.go:280
			// _ = "end of CoverTab[43657]"
//line /usr/local/go/src/net/http/transfer.go:280
		}
//line /usr/local/go/src/net/http/transfer.go:280
		// _ = "end of CoverTab[43654]"
//line /usr/local/go/src/net/http/transfer.go:280
		_go_fuzz_dep_.CoverTab[43655]++
								if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transfer.go:281
			_go_fuzz_dep_.CoverTab[43658]++
//line /usr/local/go/src/net/http/transfer.go:281
			return trace.WroteHeaderField != nil
//line /usr/local/go/src/net/http/transfer.go:281
			// _ = "end of CoverTab[43658]"
//line /usr/local/go/src/net/http/transfer.go:281
		}() {
//line /usr/local/go/src/net/http/transfer.go:281
			_go_fuzz_dep_.CoverTab[43659]++
									trace.WroteHeaderField("Connection", []string{"close"})
//line /usr/local/go/src/net/http/transfer.go:282
			// _ = "end of CoverTab[43659]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:283
			_go_fuzz_dep_.CoverTab[43660]++
//line /usr/local/go/src/net/http/transfer.go:283
			// _ = "end of CoverTab[43660]"
//line /usr/local/go/src/net/http/transfer.go:283
		}
//line /usr/local/go/src/net/http/transfer.go:283
		// _ = "end of CoverTab[43655]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:284
		_go_fuzz_dep_.CoverTab[43661]++
//line /usr/local/go/src/net/http/transfer.go:284
		// _ = "end of CoverTab[43661]"
//line /usr/local/go/src/net/http/transfer.go:284
	}
//line /usr/local/go/src/net/http/transfer.go:284
	// _ = "end of CoverTab[43649]"
//line /usr/local/go/src/net/http/transfer.go:284
	_go_fuzz_dep_.CoverTab[43650]++

//line /usr/local/go/src/net/http/transfer.go:289
	if t.shouldSendContentLength() {
//line /usr/local/go/src/net/http/transfer.go:289
		_go_fuzz_dep_.CoverTab[43662]++
								if _, err := io.WriteString(w, "Content-Length: "); err != nil {
//line /usr/local/go/src/net/http/transfer.go:290
			_go_fuzz_dep_.CoverTab[43665]++
									return err
//line /usr/local/go/src/net/http/transfer.go:291
			// _ = "end of CoverTab[43665]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:292
			_go_fuzz_dep_.CoverTab[43666]++
//line /usr/local/go/src/net/http/transfer.go:292
			// _ = "end of CoverTab[43666]"
//line /usr/local/go/src/net/http/transfer.go:292
		}
//line /usr/local/go/src/net/http/transfer.go:292
		// _ = "end of CoverTab[43662]"
//line /usr/local/go/src/net/http/transfer.go:292
		_go_fuzz_dep_.CoverTab[43663]++
								if _, err := io.WriteString(w, strconv.FormatInt(t.ContentLength, 10)+"\r\n"); err != nil {
//line /usr/local/go/src/net/http/transfer.go:293
			_go_fuzz_dep_.CoverTab[43667]++
									return err
//line /usr/local/go/src/net/http/transfer.go:294
			// _ = "end of CoverTab[43667]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:295
			_go_fuzz_dep_.CoverTab[43668]++
//line /usr/local/go/src/net/http/transfer.go:295
			// _ = "end of CoverTab[43668]"
//line /usr/local/go/src/net/http/transfer.go:295
		}
//line /usr/local/go/src/net/http/transfer.go:295
		// _ = "end of CoverTab[43663]"
//line /usr/local/go/src/net/http/transfer.go:295
		_go_fuzz_dep_.CoverTab[43664]++
								if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transfer.go:296
			_go_fuzz_dep_.CoverTab[43669]++
//line /usr/local/go/src/net/http/transfer.go:296
			return trace.WroteHeaderField != nil
//line /usr/local/go/src/net/http/transfer.go:296
			// _ = "end of CoverTab[43669]"
//line /usr/local/go/src/net/http/transfer.go:296
		}() {
//line /usr/local/go/src/net/http/transfer.go:296
			_go_fuzz_dep_.CoverTab[43670]++
									trace.WroteHeaderField("Content-Length", []string{strconv.FormatInt(t.ContentLength, 10)})
//line /usr/local/go/src/net/http/transfer.go:297
			// _ = "end of CoverTab[43670]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:298
			_go_fuzz_dep_.CoverTab[43671]++
//line /usr/local/go/src/net/http/transfer.go:298
			// _ = "end of CoverTab[43671]"
//line /usr/local/go/src/net/http/transfer.go:298
		}
//line /usr/local/go/src/net/http/transfer.go:298
		// _ = "end of CoverTab[43664]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:299
		_go_fuzz_dep_.CoverTab[43672]++
//line /usr/local/go/src/net/http/transfer.go:299
		if chunked(t.TransferEncoding) {
//line /usr/local/go/src/net/http/transfer.go:299
			_go_fuzz_dep_.CoverTab[43673]++
									if _, err := io.WriteString(w, "Transfer-Encoding: chunked\r\n"); err != nil {
//line /usr/local/go/src/net/http/transfer.go:300
				_go_fuzz_dep_.CoverTab[43675]++
										return err
//line /usr/local/go/src/net/http/transfer.go:301
				// _ = "end of CoverTab[43675]"
			} else {
//line /usr/local/go/src/net/http/transfer.go:302
				_go_fuzz_dep_.CoverTab[43676]++
//line /usr/local/go/src/net/http/transfer.go:302
				// _ = "end of CoverTab[43676]"
//line /usr/local/go/src/net/http/transfer.go:302
			}
//line /usr/local/go/src/net/http/transfer.go:302
			// _ = "end of CoverTab[43673]"
//line /usr/local/go/src/net/http/transfer.go:302
			_go_fuzz_dep_.CoverTab[43674]++
									if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transfer.go:303
				_go_fuzz_dep_.CoverTab[43677]++
//line /usr/local/go/src/net/http/transfer.go:303
				return trace.WroteHeaderField != nil
//line /usr/local/go/src/net/http/transfer.go:303
				// _ = "end of CoverTab[43677]"
//line /usr/local/go/src/net/http/transfer.go:303
			}() {
//line /usr/local/go/src/net/http/transfer.go:303
				_go_fuzz_dep_.CoverTab[43678]++
										trace.WroteHeaderField("Transfer-Encoding", []string{"chunked"})
//line /usr/local/go/src/net/http/transfer.go:304
				// _ = "end of CoverTab[43678]"
			} else {
//line /usr/local/go/src/net/http/transfer.go:305
				_go_fuzz_dep_.CoverTab[43679]++
//line /usr/local/go/src/net/http/transfer.go:305
				// _ = "end of CoverTab[43679]"
//line /usr/local/go/src/net/http/transfer.go:305
			}
//line /usr/local/go/src/net/http/transfer.go:305
			// _ = "end of CoverTab[43674]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:306
			_go_fuzz_dep_.CoverTab[43680]++
//line /usr/local/go/src/net/http/transfer.go:306
			// _ = "end of CoverTab[43680]"
//line /usr/local/go/src/net/http/transfer.go:306
		}
//line /usr/local/go/src/net/http/transfer.go:306
		// _ = "end of CoverTab[43672]"
//line /usr/local/go/src/net/http/transfer.go:306
	}
//line /usr/local/go/src/net/http/transfer.go:306
	// _ = "end of CoverTab[43650]"
//line /usr/local/go/src/net/http/transfer.go:306
	_go_fuzz_dep_.CoverTab[43651]++

//line /usr/local/go/src/net/http/transfer.go:309
	if t.Trailer != nil {
//line /usr/local/go/src/net/http/transfer.go:309
		_go_fuzz_dep_.CoverTab[43681]++
								keys := make([]string, 0, len(t.Trailer))
								for k := range t.Trailer {
//line /usr/local/go/src/net/http/transfer.go:311
			_go_fuzz_dep_.CoverTab[43683]++
									k = CanonicalHeaderKey(k)
									switch k {
			case "Transfer-Encoding", "Trailer", "Content-Length":
//line /usr/local/go/src/net/http/transfer.go:314
				_go_fuzz_dep_.CoverTab[43685]++
										return badStringError("invalid Trailer key", k)
//line /usr/local/go/src/net/http/transfer.go:315
				// _ = "end of CoverTab[43685]"
//line /usr/local/go/src/net/http/transfer.go:315
			default:
//line /usr/local/go/src/net/http/transfer.go:315
				_go_fuzz_dep_.CoverTab[43686]++
//line /usr/local/go/src/net/http/transfer.go:315
				// _ = "end of CoverTab[43686]"
			}
//line /usr/local/go/src/net/http/transfer.go:316
			// _ = "end of CoverTab[43683]"
//line /usr/local/go/src/net/http/transfer.go:316
			_go_fuzz_dep_.CoverTab[43684]++
									keys = append(keys, k)
//line /usr/local/go/src/net/http/transfer.go:317
			// _ = "end of CoverTab[43684]"
		}
//line /usr/local/go/src/net/http/transfer.go:318
		// _ = "end of CoverTab[43681]"
//line /usr/local/go/src/net/http/transfer.go:318
		_go_fuzz_dep_.CoverTab[43682]++
								if len(keys) > 0 {
//line /usr/local/go/src/net/http/transfer.go:319
			_go_fuzz_dep_.CoverTab[43687]++
									sort.Strings(keys)

//line /usr/local/go/src/net/http/transfer.go:323
			if _, err := io.WriteString(w, "Trailer: "+strings.Join(keys, ",")+"\r\n"); err != nil {
//line /usr/local/go/src/net/http/transfer.go:323
				_go_fuzz_dep_.CoverTab[43689]++
										return err
//line /usr/local/go/src/net/http/transfer.go:324
				// _ = "end of CoverTab[43689]"
			} else {
//line /usr/local/go/src/net/http/transfer.go:325
				_go_fuzz_dep_.CoverTab[43690]++
//line /usr/local/go/src/net/http/transfer.go:325
				// _ = "end of CoverTab[43690]"
//line /usr/local/go/src/net/http/transfer.go:325
			}
//line /usr/local/go/src/net/http/transfer.go:325
			// _ = "end of CoverTab[43687]"
//line /usr/local/go/src/net/http/transfer.go:325
			_go_fuzz_dep_.CoverTab[43688]++
									if trace != nil && func() bool {
//line /usr/local/go/src/net/http/transfer.go:326
				_go_fuzz_dep_.CoverTab[43691]++
//line /usr/local/go/src/net/http/transfer.go:326
				return trace.WroteHeaderField != nil
//line /usr/local/go/src/net/http/transfer.go:326
				// _ = "end of CoverTab[43691]"
//line /usr/local/go/src/net/http/transfer.go:326
			}() {
//line /usr/local/go/src/net/http/transfer.go:326
				_go_fuzz_dep_.CoverTab[43692]++
										trace.WroteHeaderField("Trailer", keys)
//line /usr/local/go/src/net/http/transfer.go:327
				// _ = "end of CoverTab[43692]"
			} else {
//line /usr/local/go/src/net/http/transfer.go:328
				_go_fuzz_dep_.CoverTab[43693]++
//line /usr/local/go/src/net/http/transfer.go:328
				// _ = "end of CoverTab[43693]"
//line /usr/local/go/src/net/http/transfer.go:328
			}
//line /usr/local/go/src/net/http/transfer.go:328
			// _ = "end of CoverTab[43688]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:329
			_go_fuzz_dep_.CoverTab[43694]++
//line /usr/local/go/src/net/http/transfer.go:329
			// _ = "end of CoverTab[43694]"
//line /usr/local/go/src/net/http/transfer.go:329
		}
//line /usr/local/go/src/net/http/transfer.go:329
		// _ = "end of CoverTab[43682]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:330
		_go_fuzz_dep_.CoverTab[43695]++
//line /usr/local/go/src/net/http/transfer.go:330
		// _ = "end of CoverTab[43695]"
//line /usr/local/go/src/net/http/transfer.go:330
	}
//line /usr/local/go/src/net/http/transfer.go:330
	// _ = "end of CoverTab[43651]"
//line /usr/local/go/src/net/http/transfer.go:330
	_go_fuzz_dep_.CoverTab[43652]++

							return nil
//line /usr/local/go/src/net/http/transfer.go:332
	// _ = "end of CoverTab[43652]"
}

// always closes t.BodyCloser
func (t *transferWriter) writeBody(w io.Writer) (err error) {
//line /usr/local/go/src/net/http/transfer.go:336
	_go_fuzz_dep_.CoverTab[43696]++
							var ncopy int64
							closed := false
							defer func() {
//line /usr/local/go/src/net/http/transfer.go:339
		_go_fuzz_dep_.CoverTab[43702]++
								if closed || func() bool {
//line /usr/local/go/src/net/http/transfer.go:340
			_go_fuzz_dep_.CoverTab[43704]++
//line /usr/local/go/src/net/http/transfer.go:340
			return t.BodyCloser == nil
//line /usr/local/go/src/net/http/transfer.go:340
			// _ = "end of CoverTab[43704]"
//line /usr/local/go/src/net/http/transfer.go:340
		}() {
//line /usr/local/go/src/net/http/transfer.go:340
			_go_fuzz_dep_.CoverTab[43705]++
									return
//line /usr/local/go/src/net/http/transfer.go:341
			// _ = "end of CoverTab[43705]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:342
			_go_fuzz_dep_.CoverTab[43706]++
//line /usr/local/go/src/net/http/transfer.go:342
			// _ = "end of CoverTab[43706]"
//line /usr/local/go/src/net/http/transfer.go:342
		}
//line /usr/local/go/src/net/http/transfer.go:342
		// _ = "end of CoverTab[43702]"
//line /usr/local/go/src/net/http/transfer.go:342
		_go_fuzz_dep_.CoverTab[43703]++
								if closeErr := t.BodyCloser.Close(); closeErr != nil && func() bool {
//line /usr/local/go/src/net/http/transfer.go:343
			_go_fuzz_dep_.CoverTab[43707]++
//line /usr/local/go/src/net/http/transfer.go:343
			return err == nil
//line /usr/local/go/src/net/http/transfer.go:343
			// _ = "end of CoverTab[43707]"
//line /usr/local/go/src/net/http/transfer.go:343
		}() {
//line /usr/local/go/src/net/http/transfer.go:343
			_go_fuzz_dep_.CoverTab[43708]++
									err = closeErr
//line /usr/local/go/src/net/http/transfer.go:344
			// _ = "end of CoverTab[43708]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:345
			_go_fuzz_dep_.CoverTab[43709]++
//line /usr/local/go/src/net/http/transfer.go:345
			// _ = "end of CoverTab[43709]"
//line /usr/local/go/src/net/http/transfer.go:345
		}
//line /usr/local/go/src/net/http/transfer.go:345
		// _ = "end of CoverTab[43703]"
	}()
//line /usr/local/go/src/net/http/transfer.go:346
	// _ = "end of CoverTab[43696]"
//line /usr/local/go/src/net/http/transfer.go:346
	_go_fuzz_dep_.CoverTab[43697]++

//line /usr/local/go/src/net/http/transfer.go:352
	if t.Body != nil {
//line /usr/local/go/src/net/http/transfer.go:352
		_go_fuzz_dep_.CoverTab[43710]++
								var body = t.unwrapBody()
								if chunked(t.TransferEncoding) {
//line /usr/local/go/src/net/http/transfer.go:354
			_go_fuzz_dep_.CoverTab[43712]++
									if bw, ok := w.(*bufio.Writer); ok && func() bool {
//line /usr/local/go/src/net/http/transfer.go:355
				_go_fuzz_dep_.CoverTab[43714]++
//line /usr/local/go/src/net/http/transfer.go:355
				return !t.IsResponse
//line /usr/local/go/src/net/http/transfer.go:355
				// _ = "end of CoverTab[43714]"
//line /usr/local/go/src/net/http/transfer.go:355
			}() {
//line /usr/local/go/src/net/http/transfer.go:355
				_go_fuzz_dep_.CoverTab[43715]++
										w = &internal.FlushAfterChunkWriter{Writer: bw}
//line /usr/local/go/src/net/http/transfer.go:356
				// _ = "end of CoverTab[43715]"
			} else {
//line /usr/local/go/src/net/http/transfer.go:357
				_go_fuzz_dep_.CoverTab[43716]++
//line /usr/local/go/src/net/http/transfer.go:357
				// _ = "end of CoverTab[43716]"
//line /usr/local/go/src/net/http/transfer.go:357
			}
//line /usr/local/go/src/net/http/transfer.go:357
			// _ = "end of CoverTab[43712]"
//line /usr/local/go/src/net/http/transfer.go:357
			_go_fuzz_dep_.CoverTab[43713]++
									cw := internal.NewChunkedWriter(w)
									_, err = t.doBodyCopy(cw, body)
									if err == nil {
//line /usr/local/go/src/net/http/transfer.go:360
				_go_fuzz_dep_.CoverTab[43717]++
										err = cw.Close()
//line /usr/local/go/src/net/http/transfer.go:361
				// _ = "end of CoverTab[43717]"
			} else {
//line /usr/local/go/src/net/http/transfer.go:362
				_go_fuzz_dep_.CoverTab[43718]++
//line /usr/local/go/src/net/http/transfer.go:362
				// _ = "end of CoverTab[43718]"
//line /usr/local/go/src/net/http/transfer.go:362
			}
//line /usr/local/go/src/net/http/transfer.go:362
			// _ = "end of CoverTab[43713]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:363
			_go_fuzz_dep_.CoverTab[43719]++
//line /usr/local/go/src/net/http/transfer.go:363
			if t.ContentLength == -1 {
//line /usr/local/go/src/net/http/transfer.go:363
				_go_fuzz_dep_.CoverTab[43720]++
										dst := w
										if t.Method == "CONNECT" {
//line /usr/local/go/src/net/http/transfer.go:365
					_go_fuzz_dep_.CoverTab[43722]++
											dst = bufioFlushWriter{dst}
//line /usr/local/go/src/net/http/transfer.go:366
					// _ = "end of CoverTab[43722]"
				} else {
//line /usr/local/go/src/net/http/transfer.go:367
					_go_fuzz_dep_.CoverTab[43723]++
//line /usr/local/go/src/net/http/transfer.go:367
					// _ = "end of CoverTab[43723]"
//line /usr/local/go/src/net/http/transfer.go:367
				}
//line /usr/local/go/src/net/http/transfer.go:367
				// _ = "end of CoverTab[43720]"
//line /usr/local/go/src/net/http/transfer.go:367
				_go_fuzz_dep_.CoverTab[43721]++
										ncopy, err = t.doBodyCopy(dst, body)
//line /usr/local/go/src/net/http/transfer.go:368
				// _ = "end of CoverTab[43721]"
			} else {
//line /usr/local/go/src/net/http/transfer.go:369
				_go_fuzz_dep_.CoverTab[43724]++
										ncopy, err = t.doBodyCopy(w, io.LimitReader(body, t.ContentLength))
										if err != nil {
//line /usr/local/go/src/net/http/transfer.go:371
					_go_fuzz_dep_.CoverTab[43726]++
											return err
//line /usr/local/go/src/net/http/transfer.go:372
					// _ = "end of CoverTab[43726]"
				} else {
//line /usr/local/go/src/net/http/transfer.go:373
					_go_fuzz_dep_.CoverTab[43727]++
//line /usr/local/go/src/net/http/transfer.go:373
					// _ = "end of CoverTab[43727]"
//line /usr/local/go/src/net/http/transfer.go:373
				}
//line /usr/local/go/src/net/http/transfer.go:373
				// _ = "end of CoverTab[43724]"
//line /usr/local/go/src/net/http/transfer.go:373
				_go_fuzz_dep_.CoverTab[43725]++
										var nextra int64
										nextra, err = t.doBodyCopy(io.Discard, body)
										ncopy += nextra
//line /usr/local/go/src/net/http/transfer.go:376
				// _ = "end of CoverTab[43725]"
			}
//line /usr/local/go/src/net/http/transfer.go:377
			// _ = "end of CoverTab[43719]"
//line /usr/local/go/src/net/http/transfer.go:377
		}
//line /usr/local/go/src/net/http/transfer.go:377
		// _ = "end of CoverTab[43710]"
//line /usr/local/go/src/net/http/transfer.go:377
		_go_fuzz_dep_.CoverTab[43711]++
								if err != nil {
//line /usr/local/go/src/net/http/transfer.go:378
			_go_fuzz_dep_.CoverTab[43728]++
									return err
//line /usr/local/go/src/net/http/transfer.go:379
			// _ = "end of CoverTab[43728]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:380
			_go_fuzz_dep_.CoverTab[43729]++
//line /usr/local/go/src/net/http/transfer.go:380
			// _ = "end of CoverTab[43729]"
//line /usr/local/go/src/net/http/transfer.go:380
		}
//line /usr/local/go/src/net/http/transfer.go:380
		// _ = "end of CoverTab[43711]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:381
		_go_fuzz_dep_.CoverTab[43730]++
//line /usr/local/go/src/net/http/transfer.go:381
		// _ = "end of CoverTab[43730]"
//line /usr/local/go/src/net/http/transfer.go:381
	}
//line /usr/local/go/src/net/http/transfer.go:381
	// _ = "end of CoverTab[43697]"
//line /usr/local/go/src/net/http/transfer.go:381
	_go_fuzz_dep_.CoverTab[43698]++
							if t.BodyCloser != nil {
//line /usr/local/go/src/net/http/transfer.go:382
		_go_fuzz_dep_.CoverTab[43731]++
								closed = true
								if err := t.BodyCloser.Close(); err != nil {
//line /usr/local/go/src/net/http/transfer.go:384
			_go_fuzz_dep_.CoverTab[43732]++
									return err
//line /usr/local/go/src/net/http/transfer.go:385
			// _ = "end of CoverTab[43732]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:386
			_go_fuzz_dep_.CoverTab[43733]++
//line /usr/local/go/src/net/http/transfer.go:386
			// _ = "end of CoverTab[43733]"
//line /usr/local/go/src/net/http/transfer.go:386
		}
//line /usr/local/go/src/net/http/transfer.go:386
		// _ = "end of CoverTab[43731]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:387
		_go_fuzz_dep_.CoverTab[43734]++
//line /usr/local/go/src/net/http/transfer.go:387
		// _ = "end of CoverTab[43734]"
//line /usr/local/go/src/net/http/transfer.go:387
	}
//line /usr/local/go/src/net/http/transfer.go:387
	// _ = "end of CoverTab[43698]"
//line /usr/local/go/src/net/http/transfer.go:387
	_go_fuzz_dep_.CoverTab[43699]++

							if !t.ResponseToHEAD && func() bool {
//line /usr/local/go/src/net/http/transfer.go:389
		_go_fuzz_dep_.CoverTab[43735]++
//line /usr/local/go/src/net/http/transfer.go:389
		return t.ContentLength != -1
//line /usr/local/go/src/net/http/transfer.go:389
		// _ = "end of CoverTab[43735]"
//line /usr/local/go/src/net/http/transfer.go:389
	}() && func() bool {
//line /usr/local/go/src/net/http/transfer.go:389
		_go_fuzz_dep_.CoverTab[43736]++
//line /usr/local/go/src/net/http/transfer.go:389
		return t.ContentLength != ncopy
//line /usr/local/go/src/net/http/transfer.go:389
		// _ = "end of CoverTab[43736]"
//line /usr/local/go/src/net/http/transfer.go:389
	}() {
//line /usr/local/go/src/net/http/transfer.go:389
		_go_fuzz_dep_.CoverTab[43737]++
								return fmt.Errorf("http: ContentLength=%d with Body length %d",
			t.ContentLength, ncopy)
//line /usr/local/go/src/net/http/transfer.go:391
		// _ = "end of CoverTab[43737]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:392
		_go_fuzz_dep_.CoverTab[43738]++
//line /usr/local/go/src/net/http/transfer.go:392
		// _ = "end of CoverTab[43738]"
//line /usr/local/go/src/net/http/transfer.go:392
	}
//line /usr/local/go/src/net/http/transfer.go:392
	// _ = "end of CoverTab[43699]"
//line /usr/local/go/src/net/http/transfer.go:392
	_go_fuzz_dep_.CoverTab[43700]++

							if chunked(t.TransferEncoding) {
//line /usr/local/go/src/net/http/transfer.go:394
		_go_fuzz_dep_.CoverTab[43739]++

								if t.Trailer != nil {
//line /usr/local/go/src/net/http/transfer.go:396
			_go_fuzz_dep_.CoverTab[43741]++
									if err := t.Trailer.Write(w); err != nil {
//line /usr/local/go/src/net/http/transfer.go:397
				_go_fuzz_dep_.CoverTab[43742]++
										return err
//line /usr/local/go/src/net/http/transfer.go:398
				// _ = "end of CoverTab[43742]"
			} else {
//line /usr/local/go/src/net/http/transfer.go:399
				_go_fuzz_dep_.CoverTab[43743]++
//line /usr/local/go/src/net/http/transfer.go:399
				// _ = "end of CoverTab[43743]"
//line /usr/local/go/src/net/http/transfer.go:399
			}
//line /usr/local/go/src/net/http/transfer.go:399
			// _ = "end of CoverTab[43741]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:400
			_go_fuzz_dep_.CoverTab[43744]++
//line /usr/local/go/src/net/http/transfer.go:400
			// _ = "end of CoverTab[43744]"
//line /usr/local/go/src/net/http/transfer.go:400
		}
//line /usr/local/go/src/net/http/transfer.go:400
		// _ = "end of CoverTab[43739]"
//line /usr/local/go/src/net/http/transfer.go:400
		_go_fuzz_dep_.CoverTab[43740]++

								_, err = io.WriteString(w, "\r\n")
//line /usr/local/go/src/net/http/transfer.go:402
		// _ = "end of CoverTab[43740]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:403
		_go_fuzz_dep_.CoverTab[43745]++
//line /usr/local/go/src/net/http/transfer.go:403
		// _ = "end of CoverTab[43745]"
//line /usr/local/go/src/net/http/transfer.go:403
	}
//line /usr/local/go/src/net/http/transfer.go:403
	// _ = "end of CoverTab[43700]"
//line /usr/local/go/src/net/http/transfer.go:403
	_go_fuzz_dep_.CoverTab[43701]++
							return err
//line /usr/local/go/src/net/http/transfer.go:404
	// _ = "end of CoverTab[43701]"
}

// doBodyCopy wraps a copy operation, with any resulting error also
//line /usr/local/go/src/net/http/transfer.go:407
// being saved in bodyReadError.
//line /usr/local/go/src/net/http/transfer.go:407
//
//line /usr/local/go/src/net/http/transfer.go:407
// This function is only intended for use in writeBody.
//line /usr/local/go/src/net/http/transfer.go:411
func (t *transferWriter) doBodyCopy(dst io.Writer, src io.Reader) (n int64, err error) {
//line /usr/local/go/src/net/http/transfer.go:411
	_go_fuzz_dep_.CoverTab[43746]++
							n, err = io.Copy(dst, src)
							if err != nil && func() bool {
//line /usr/local/go/src/net/http/transfer.go:413
		_go_fuzz_dep_.CoverTab[43748]++
//line /usr/local/go/src/net/http/transfer.go:413
		return err != io.EOF
//line /usr/local/go/src/net/http/transfer.go:413
		// _ = "end of CoverTab[43748]"
//line /usr/local/go/src/net/http/transfer.go:413
	}() {
//line /usr/local/go/src/net/http/transfer.go:413
		_go_fuzz_dep_.CoverTab[43749]++
								t.bodyReadError = err
//line /usr/local/go/src/net/http/transfer.go:414
		// _ = "end of CoverTab[43749]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:415
		_go_fuzz_dep_.CoverTab[43750]++
//line /usr/local/go/src/net/http/transfer.go:415
		// _ = "end of CoverTab[43750]"
//line /usr/local/go/src/net/http/transfer.go:415
	}
//line /usr/local/go/src/net/http/transfer.go:415
	// _ = "end of CoverTab[43746]"
//line /usr/local/go/src/net/http/transfer.go:415
	_go_fuzz_dep_.CoverTab[43747]++
							return
//line /usr/local/go/src/net/http/transfer.go:416
	// _ = "end of CoverTab[43747]"
}

// unwrapBodyReader unwraps the body's inner reader if it's a
//line /usr/local/go/src/net/http/transfer.go:419
// nopCloser. This is to ensure that body writes sourced from local
//line /usr/local/go/src/net/http/transfer.go:419
// files (*os.File types) are properly optimized.
//line /usr/local/go/src/net/http/transfer.go:419
//
//line /usr/local/go/src/net/http/transfer.go:419
// This function is only intended for use in writeBody.
//line /usr/local/go/src/net/http/transfer.go:424
func (t *transferWriter) unwrapBody() io.Reader {
//line /usr/local/go/src/net/http/transfer.go:424
	_go_fuzz_dep_.CoverTab[43751]++
							if r, ok := unwrapNopCloser(t.Body); ok {
//line /usr/local/go/src/net/http/transfer.go:425
		_go_fuzz_dep_.CoverTab[43754]++
								return r
//line /usr/local/go/src/net/http/transfer.go:426
		// _ = "end of CoverTab[43754]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:427
		_go_fuzz_dep_.CoverTab[43755]++
//line /usr/local/go/src/net/http/transfer.go:427
		// _ = "end of CoverTab[43755]"
//line /usr/local/go/src/net/http/transfer.go:427
	}
//line /usr/local/go/src/net/http/transfer.go:427
	// _ = "end of CoverTab[43751]"
//line /usr/local/go/src/net/http/transfer.go:427
	_go_fuzz_dep_.CoverTab[43752]++
							if r, ok := t.Body.(*readTrackingBody); ok {
//line /usr/local/go/src/net/http/transfer.go:428
		_go_fuzz_dep_.CoverTab[43756]++
								r.didRead = true
								return r.ReadCloser
//line /usr/local/go/src/net/http/transfer.go:430
		// _ = "end of CoverTab[43756]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:431
		_go_fuzz_dep_.CoverTab[43757]++
//line /usr/local/go/src/net/http/transfer.go:431
		// _ = "end of CoverTab[43757]"
//line /usr/local/go/src/net/http/transfer.go:431
	}
//line /usr/local/go/src/net/http/transfer.go:431
	// _ = "end of CoverTab[43752]"
//line /usr/local/go/src/net/http/transfer.go:431
	_go_fuzz_dep_.CoverTab[43753]++
							return t.Body
//line /usr/local/go/src/net/http/transfer.go:432
	// _ = "end of CoverTab[43753]"
}

type transferReader struct {
	// Input
	Header		Header
	StatusCode	int
	RequestMethod	string
	ProtoMajor	int
	ProtoMinor	int
	// Output
	Body		io.ReadCloser
	ContentLength	int64
	Chunked		bool
	Close		bool
	Trailer		Header
}

func (t *transferReader) protoAtLeast(m, n int) bool {
//line /usr/local/go/src/net/http/transfer.go:450
	_go_fuzz_dep_.CoverTab[43758]++
							return t.ProtoMajor > m || func() bool {
//line /usr/local/go/src/net/http/transfer.go:451
		_go_fuzz_dep_.CoverTab[43759]++
//line /usr/local/go/src/net/http/transfer.go:451
		return (t.ProtoMajor == m && func() bool {
//line /usr/local/go/src/net/http/transfer.go:451
			_go_fuzz_dep_.CoverTab[43760]++
//line /usr/local/go/src/net/http/transfer.go:451
			return t.ProtoMinor >= n
//line /usr/local/go/src/net/http/transfer.go:451
			// _ = "end of CoverTab[43760]"
//line /usr/local/go/src/net/http/transfer.go:451
		}())
//line /usr/local/go/src/net/http/transfer.go:451
		// _ = "end of CoverTab[43759]"
//line /usr/local/go/src/net/http/transfer.go:451
	}()
//line /usr/local/go/src/net/http/transfer.go:451
	// _ = "end of CoverTab[43758]"
}

// bodyAllowedForStatus reports whether a given response status code
//line /usr/local/go/src/net/http/transfer.go:454
// permits a body. See RFC 7230, section 3.3.
//line /usr/local/go/src/net/http/transfer.go:456
func bodyAllowedForStatus(status int) bool {
//line /usr/local/go/src/net/http/transfer.go:456
	_go_fuzz_dep_.CoverTab[43761]++
							switch {
	case status >= 100 && func() bool {
//line /usr/local/go/src/net/http/transfer.go:458
		_go_fuzz_dep_.CoverTab[43767]++
//line /usr/local/go/src/net/http/transfer.go:458
		return status <= 199
//line /usr/local/go/src/net/http/transfer.go:458
		// _ = "end of CoverTab[43767]"
//line /usr/local/go/src/net/http/transfer.go:458
	}():
//line /usr/local/go/src/net/http/transfer.go:458
		_go_fuzz_dep_.CoverTab[43763]++
								return false
//line /usr/local/go/src/net/http/transfer.go:459
		// _ = "end of CoverTab[43763]"
	case status == 204:
//line /usr/local/go/src/net/http/transfer.go:460
		_go_fuzz_dep_.CoverTab[43764]++
								return false
//line /usr/local/go/src/net/http/transfer.go:461
		// _ = "end of CoverTab[43764]"
	case status == 304:
//line /usr/local/go/src/net/http/transfer.go:462
		_go_fuzz_dep_.CoverTab[43765]++
								return false
//line /usr/local/go/src/net/http/transfer.go:463
		// _ = "end of CoverTab[43765]"
//line /usr/local/go/src/net/http/transfer.go:463
	default:
//line /usr/local/go/src/net/http/transfer.go:463
		_go_fuzz_dep_.CoverTab[43766]++
//line /usr/local/go/src/net/http/transfer.go:463
		// _ = "end of CoverTab[43766]"
	}
//line /usr/local/go/src/net/http/transfer.go:464
	// _ = "end of CoverTab[43761]"
//line /usr/local/go/src/net/http/transfer.go:464
	_go_fuzz_dep_.CoverTab[43762]++
							return true
//line /usr/local/go/src/net/http/transfer.go:465
	// _ = "end of CoverTab[43762]"
}

var (
	suppressedHeaders304	= []string{"Content-Type", "Content-Length", "Transfer-Encoding"}
	suppressedHeadersNoBody	= []string{"Content-Length", "Transfer-Encoding"}
	excludedHeadersNoBody	= map[string]bool{"Content-Length": true, "Transfer-Encoding": true}
)

func suppressedHeaders(status int) []string {
//line /usr/local/go/src/net/http/transfer.go:474
	_go_fuzz_dep_.CoverTab[43768]++
							switch {
	case status == 304:
//line /usr/local/go/src/net/http/transfer.go:476
		_go_fuzz_dep_.CoverTab[43770]++

								return suppressedHeaders304
//line /usr/local/go/src/net/http/transfer.go:478
		// _ = "end of CoverTab[43770]"
	case !bodyAllowedForStatus(status):
//line /usr/local/go/src/net/http/transfer.go:479
		_go_fuzz_dep_.CoverTab[43771]++
								return suppressedHeadersNoBody
//line /usr/local/go/src/net/http/transfer.go:480
		// _ = "end of CoverTab[43771]"
//line /usr/local/go/src/net/http/transfer.go:480
	default:
//line /usr/local/go/src/net/http/transfer.go:480
		_go_fuzz_dep_.CoverTab[43772]++
//line /usr/local/go/src/net/http/transfer.go:480
		// _ = "end of CoverTab[43772]"
	}
//line /usr/local/go/src/net/http/transfer.go:481
	// _ = "end of CoverTab[43768]"
//line /usr/local/go/src/net/http/transfer.go:481
	_go_fuzz_dep_.CoverTab[43769]++
							return nil
//line /usr/local/go/src/net/http/transfer.go:482
	// _ = "end of CoverTab[43769]"
}

// msg is *Request or *Response.
func readTransfer(msg any, r *bufio.Reader) (err error) {
//line /usr/local/go/src/net/http/transfer.go:486
	_go_fuzz_dep_.CoverTab[43773]++
							t := &transferReader{RequestMethod: "GET"}

//line /usr/local/go/src/net/http/transfer.go:490
	isResponse := false
	switch rr := msg.(type) {
	case *Response:
//line /usr/local/go/src/net/http/transfer.go:492
		_go_fuzz_dep_.CoverTab[43783]++
								t.Header = rr.Header
								t.StatusCode = rr.StatusCode
								t.ProtoMajor = rr.ProtoMajor
								t.ProtoMinor = rr.ProtoMinor
								t.Close = shouldClose(t.ProtoMajor, t.ProtoMinor, t.Header, true)
								isResponse = true
								if rr.Request != nil {
//line /usr/local/go/src/net/http/transfer.go:499
			_go_fuzz_dep_.CoverTab[43786]++
									t.RequestMethod = rr.Request.Method
//line /usr/local/go/src/net/http/transfer.go:500
			// _ = "end of CoverTab[43786]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:501
			_go_fuzz_dep_.CoverTab[43787]++
//line /usr/local/go/src/net/http/transfer.go:501
			// _ = "end of CoverTab[43787]"
//line /usr/local/go/src/net/http/transfer.go:501
		}
//line /usr/local/go/src/net/http/transfer.go:501
		// _ = "end of CoverTab[43783]"
	case *Request:
//line /usr/local/go/src/net/http/transfer.go:502
		_go_fuzz_dep_.CoverTab[43784]++
								t.Header = rr.Header
								t.RequestMethod = rr.Method
								t.ProtoMajor = rr.ProtoMajor
								t.ProtoMinor = rr.ProtoMinor

//line /usr/local/go/src/net/http/transfer.go:509
		t.StatusCode = 200
								t.Close = rr.Close
//line /usr/local/go/src/net/http/transfer.go:510
		// _ = "end of CoverTab[43784]"
	default:
//line /usr/local/go/src/net/http/transfer.go:511
		_go_fuzz_dep_.CoverTab[43785]++
								panic("unexpected type")
//line /usr/local/go/src/net/http/transfer.go:512
		// _ = "end of CoverTab[43785]"
	}
//line /usr/local/go/src/net/http/transfer.go:513
	// _ = "end of CoverTab[43773]"
//line /usr/local/go/src/net/http/transfer.go:513
	_go_fuzz_dep_.CoverTab[43774]++

//line /usr/local/go/src/net/http/transfer.go:516
	if t.ProtoMajor == 0 && func() bool {
//line /usr/local/go/src/net/http/transfer.go:516
		_go_fuzz_dep_.CoverTab[43788]++
//line /usr/local/go/src/net/http/transfer.go:516
		return t.ProtoMinor == 0
//line /usr/local/go/src/net/http/transfer.go:516
		// _ = "end of CoverTab[43788]"
//line /usr/local/go/src/net/http/transfer.go:516
	}() {
//line /usr/local/go/src/net/http/transfer.go:516
		_go_fuzz_dep_.CoverTab[43789]++
								t.ProtoMajor, t.ProtoMinor = 1, 1
//line /usr/local/go/src/net/http/transfer.go:517
		// _ = "end of CoverTab[43789]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:518
		_go_fuzz_dep_.CoverTab[43790]++
//line /usr/local/go/src/net/http/transfer.go:518
		// _ = "end of CoverTab[43790]"
//line /usr/local/go/src/net/http/transfer.go:518
	}
//line /usr/local/go/src/net/http/transfer.go:518
	// _ = "end of CoverTab[43774]"
//line /usr/local/go/src/net/http/transfer.go:518
	_go_fuzz_dep_.CoverTab[43775]++

//line /usr/local/go/src/net/http/transfer.go:521
	if err := t.parseTransferEncoding(); err != nil {
//line /usr/local/go/src/net/http/transfer.go:521
		_go_fuzz_dep_.CoverTab[43791]++
								return err
//line /usr/local/go/src/net/http/transfer.go:522
		// _ = "end of CoverTab[43791]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:523
		_go_fuzz_dep_.CoverTab[43792]++
//line /usr/local/go/src/net/http/transfer.go:523
		// _ = "end of CoverTab[43792]"
//line /usr/local/go/src/net/http/transfer.go:523
	}
//line /usr/local/go/src/net/http/transfer.go:523
	// _ = "end of CoverTab[43775]"
//line /usr/local/go/src/net/http/transfer.go:523
	_go_fuzz_dep_.CoverTab[43776]++

							realLength, err := fixLength(isResponse, t.StatusCode, t.RequestMethod, t.Header, t.Chunked)
							if err != nil {
//line /usr/local/go/src/net/http/transfer.go:526
		_go_fuzz_dep_.CoverTab[43793]++
								return err
//line /usr/local/go/src/net/http/transfer.go:527
		// _ = "end of CoverTab[43793]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:528
		_go_fuzz_dep_.CoverTab[43794]++
//line /usr/local/go/src/net/http/transfer.go:528
		// _ = "end of CoverTab[43794]"
//line /usr/local/go/src/net/http/transfer.go:528
	}
//line /usr/local/go/src/net/http/transfer.go:528
	// _ = "end of CoverTab[43776]"
//line /usr/local/go/src/net/http/transfer.go:528
	_go_fuzz_dep_.CoverTab[43777]++
							if isResponse && func() bool {
//line /usr/local/go/src/net/http/transfer.go:529
		_go_fuzz_dep_.CoverTab[43795]++
//line /usr/local/go/src/net/http/transfer.go:529
		return t.RequestMethod == "HEAD"
//line /usr/local/go/src/net/http/transfer.go:529
		// _ = "end of CoverTab[43795]"
//line /usr/local/go/src/net/http/transfer.go:529
	}() {
//line /usr/local/go/src/net/http/transfer.go:529
		_go_fuzz_dep_.CoverTab[43796]++
								if n, err := parseContentLength(t.Header.get("Content-Length")); err != nil {
//line /usr/local/go/src/net/http/transfer.go:530
			_go_fuzz_dep_.CoverTab[43797]++
									return err
//line /usr/local/go/src/net/http/transfer.go:531
			// _ = "end of CoverTab[43797]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:532
			_go_fuzz_dep_.CoverTab[43798]++
									t.ContentLength = n
//line /usr/local/go/src/net/http/transfer.go:533
			// _ = "end of CoverTab[43798]"
		}
//line /usr/local/go/src/net/http/transfer.go:534
		// _ = "end of CoverTab[43796]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:535
		_go_fuzz_dep_.CoverTab[43799]++
								t.ContentLength = realLength
//line /usr/local/go/src/net/http/transfer.go:536
		// _ = "end of CoverTab[43799]"
	}
//line /usr/local/go/src/net/http/transfer.go:537
	// _ = "end of CoverTab[43777]"
//line /usr/local/go/src/net/http/transfer.go:537
	_go_fuzz_dep_.CoverTab[43778]++

//line /usr/local/go/src/net/http/transfer.go:540
	t.Trailer, err = fixTrailer(t.Header, t.Chunked)
	if err != nil {
//line /usr/local/go/src/net/http/transfer.go:541
		_go_fuzz_dep_.CoverTab[43800]++
								return err
//line /usr/local/go/src/net/http/transfer.go:542
		// _ = "end of CoverTab[43800]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:543
		_go_fuzz_dep_.CoverTab[43801]++
//line /usr/local/go/src/net/http/transfer.go:543
		// _ = "end of CoverTab[43801]"
//line /usr/local/go/src/net/http/transfer.go:543
	}
//line /usr/local/go/src/net/http/transfer.go:543
	// _ = "end of CoverTab[43778]"
//line /usr/local/go/src/net/http/transfer.go:543
	_go_fuzz_dep_.CoverTab[43779]++

//line /usr/local/go/src/net/http/transfer.go:548
	switch msg.(type) {
	case *Response:
//line /usr/local/go/src/net/http/transfer.go:549
		_go_fuzz_dep_.CoverTab[43802]++
								if realLength == -1 && func() bool {
//line /usr/local/go/src/net/http/transfer.go:550
			_go_fuzz_dep_.CoverTab[43803]++
//line /usr/local/go/src/net/http/transfer.go:550
			return !t.Chunked
//line /usr/local/go/src/net/http/transfer.go:550
			// _ = "end of CoverTab[43803]"
//line /usr/local/go/src/net/http/transfer.go:550
		}() && func() bool {
//line /usr/local/go/src/net/http/transfer.go:550
			_go_fuzz_dep_.CoverTab[43804]++
//line /usr/local/go/src/net/http/transfer.go:550
			return bodyAllowedForStatus(t.StatusCode)
//line /usr/local/go/src/net/http/transfer.go:550
			// _ = "end of CoverTab[43804]"
//line /usr/local/go/src/net/http/transfer.go:550
		}() {
//line /usr/local/go/src/net/http/transfer.go:550
			_go_fuzz_dep_.CoverTab[43805]++

									t.Close = true
//line /usr/local/go/src/net/http/transfer.go:552
			// _ = "end of CoverTab[43805]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:553
			_go_fuzz_dep_.CoverTab[43806]++
//line /usr/local/go/src/net/http/transfer.go:553
			// _ = "end of CoverTab[43806]"
//line /usr/local/go/src/net/http/transfer.go:553
		}
//line /usr/local/go/src/net/http/transfer.go:553
		// _ = "end of CoverTab[43802]"
	}
//line /usr/local/go/src/net/http/transfer.go:554
	// _ = "end of CoverTab[43779]"
//line /usr/local/go/src/net/http/transfer.go:554
	_go_fuzz_dep_.CoverTab[43780]++

//line /usr/local/go/src/net/http/transfer.go:558
	switch {
	case t.Chunked:
//line /usr/local/go/src/net/http/transfer.go:559
		_go_fuzz_dep_.CoverTab[43807]++
								if isResponse && func() bool {
//line /usr/local/go/src/net/http/transfer.go:560
			_go_fuzz_dep_.CoverTab[43811]++
//line /usr/local/go/src/net/http/transfer.go:560
			return (noResponseBodyExpected(t.RequestMethod) || func() bool {
//line /usr/local/go/src/net/http/transfer.go:560
				_go_fuzz_dep_.CoverTab[43812]++
//line /usr/local/go/src/net/http/transfer.go:560
				return !bodyAllowedForStatus(t.StatusCode)
//line /usr/local/go/src/net/http/transfer.go:560
				// _ = "end of CoverTab[43812]"
//line /usr/local/go/src/net/http/transfer.go:560
			}())
//line /usr/local/go/src/net/http/transfer.go:560
			// _ = "end of CoverTab[43811]"
//line /usr/local/go/src/net/http/transfer.go:560
		}() {
//line /usr/local/go/src/net/http/transfer.go:560
			_go_fuzz_dep_.CoverTab[43813]++
									t.Body = NoBody
//line /usr/local/go/src/net/http/transfer.go:561
			// _ = "end of CoverTab[43813]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:562
			_go_fuzz_dep_.CoverTab[43814]++
									t.Body = &body{src: internal.NewChunkedReader(r), hdr: msg, r: r, closing: t.Close}
//line /usr/local/go/src/net/http/transfer.go:563
			// _ = "end of CoverTab[43814]"
		}
//line /usr/local/go/src/net/http/transfer.go:564
		// _ = "end of CoverTab[43807]"
	case realLength == 0:
//line /usr/local/go/src/net/http/transfer.go:565
		_go_fuzz_dep_.CoverTab[43808]++
								t.Body = NoBody
//line /usr/local/go/src/net/http/transfer.go:566
		// _ = "end of CoverTab[43808]"
	case realLength > 0:
//line /usr/local/go/src/net/http/transfer.go:567
		_go_fuzz_dep_.CoverTab[43809]++
								t.Body = &body{src: io.LimitReader(r, realLength), closing: t.Close}
//line /usr/local/go/src/net/http/transfer.go:568
		// _ = "end of CoverTab[43809]"
	default:
//line /usr/local/go/src/net/http/transfer.go:569
		_go_fuzz_dep_.CoverTab[43810]++

								if t.Close {
//line /usr/local/go/src/net/http/transfer.go:571
			_go_fuzz_dep_.CoverTab[43815]++

									t.Body = &body{src: r, closing: t.Close}
//line /usr/local/go/src/net/http/transfer.go:573
			// _ = "end of CoverTab[43815]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:574
			_go_fuzz_dep_.CoverTab[43816]++

									t.Body = NoBody
//line /usr/local/go/src/net/http/transfer.go:576
			// _ = "end of CoverTab[43816]"
		}
//line /usr/local/go/src/net/http/transfer.go:577
		// _ = "end of CoverTab[43810]"
	}
//line /usr/local/go/src/net/http/transfer.go:578
	// _ = "end of CoverTab[43780]"
//line /usr/local/go/src/net/http/transfer.go:578
	_go_fuzz_dep_.CoverTab[43781]++

//line /usr/local/go/src/net/http/transfer.go:581
	switch rr := msg.(type) {
	case *Request:
//line /usr/local/go/src/net/http/transfer.go:582
		_go_fuzz_dep_.CoverTab[43817]++
								rr.Body = t.Body
								rr.ContentLength = t.ContentLength
								if t.Chunked {
//line /usr/local/go/src/net/http/transfer.go:585
			_go_fuzz_dep_.CoverTab[43821]++
									rr.TransferEncoding = []string{"chunked"}
//line /usr/local/go/src/net/http/transfer.go:586
			// _ = "end of CoverTab[43821]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:587
			_go_fuzz_dep_.CoverTab[43822]++
//line /usr/local/go/src/net/http/transfer.go:587
			// _ = "end of CoverTab[43822]"
//line /usr/local/go/src/net/http/transfer.go:587
		}
//line /usr/local/go/src/net/http/transfer.go:587
		// _ = "end of CoverTab[43817]"
//line /usr/local/go/src/net/http/transfer.go:587
		_go_fuzz_dep_.CoverTab[43818]++
								rr.Close = t.Close
								rr.Trailer = t.Trailer
//line /usr/local/go/src/net/http/transfer.go:589
		// _ = "end of CoverTab[43818]"
	case *Response:
//line /usr/local/go/src/net/http/transfer.go:590
		_go_fuzz_dep_.CoverTab[43819]++
								rr.Body = t.Body
								rr.ContentLength = t.ContentLength
								if t.Chunked {
//line /usr/local/go/src/net/http/transfer.go:593
			_go_fuzz_dep_.CoverTab[43823]++
									rr.TransferEncoding = []string{"chunked"}
//line /usr/local/go/src/net/http/transfer.go:594
			// _ = "end of CoverTab[43823]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:595
			_go_fuzz_dep_.CoverTab[43824]++
//line /usr/local/go/src/net/http/transfer.go:595
			// _ = "end of CoverTab[43824]"
//line /usr/local/go/src/net/http/transfer.go:595
		}
//line /usr/local/go/src/net/http/transfer.go:595
		// _ = "end of CoverTab[43819]"
//line /usr/local/go/src/net/http/transfer.go:595
		_go_fuzz_dep_.CoverTab[43820]++
								rr.Close = t.Close
								rr.Trailer = t.Trailer
//line /usr/local/go/src/net/http/transfer.go:597
		// _ = "end of CoverTab[43820]"
	}
//line /usr/local/go/src/net/http/transfer.go:598
	// _ = "end of CoverTab[43781]"
//line /usr/local/go/src/net/http/transfer.go:598
	_go_fuzz_dep_.CoverTab[43782]++

							return nil
//line /usr/local/go/src/net/http/transfer.go:600
	// _ = "end of CoverTab[43782]"
}

// Checks whether chunked is part of the encodings stack.
func chunked(te []string) bool {
//line /usr/local/go/src/net/http/transfer.go:604
	_go_fuzz_dep_.CoverTab[43825]++
//line /usr/local/go/src/net/http/transfer.go:604
	return len(te) > 0 && func() bool {
//line /usr/local/go/src/net/http/transfer.go:604
		_go_fuzz_dep_.CoverTab[43826]++
//line /usr/local/go/src/net/http/transfer.go:604
		return te[0] == "chunked"
//line /usr/local/go/src/net/http/transfer.go:604
		// _ = "end of CoverTab[43826]"
//line /usr/local/go/src/net/http/transfer.go:604
	}()
//line /usr/local/go/src/net/http/transfer.go:604
	// _ = "end of CoverTab[43825]"
//line /usr/local/go/src/net/http/transfer.go:604
}

// Checks whether the encoding is explicitly "identity".
func isIdentity(te []string) bool {
//line /usr/local/go/src/net/http/transfer.go:607
	_go_fuzz_dep_.CoverTab[43827]++
//line /usr/local/go/src/net/http/transfer.go:607
	return len(te) == 1 && func() bool {
//line /usr/local/go/src/net/http/transfer.go:607
		_go_fuzz_dep_.CoverTab[43828]++
//line /usr/local/go/src/net/http/transfer.go:607
		return te[0] == "identity"
//line /usr/local/go/src/net/http/transfer.go:607
		// _ = "end of CoverTab[43828]"
//line /usr/local/go/src/net/http/transfer.go:607
	}()
//line /usr/local/go/src/net/http/transfer.go:607
	// _ = "end of CoverTab[43827]"
//line /usr/local/go/src/net/http/transfer.go:607
}

// unsupportedTEError reports unsupported transfer-encodings.
type unsupportedTEError struct {
	err string
}

func (uste *unsupportedTEError) Error() string {
//line /usr/local/go/src/net/http/transfer.go:614
	_go_fuzz_dep_.CoverTab[43829]++
							return uste.err
//line /usr/local/go/src/net/http/transfer.go:615
	// _ = "end of CoverTab[43829]"
}

// isUnsupportedTEError checks if the error is of type
//line /usr/local/go/src/net/http/transfer.go:618
// unsupportedTEError. It is usually invoked with a non-nil err.
//line /usr/local/go/src/net/http/transfer.go:620
func isUnsupportedTEError(err error) bool {
//line /usr/local/go/src/net/http/transfer.go:620
	_go_fuzz_dep_.CoverTab[43830]++
							_, ok := err.(*unsupportedTEError)
							return ok
//line /usr/local/go/src/net/http/transfer.go:622
	// _ = "end of CoverTab[43830]"
}

// parseTransferEncoding sets t.Chunked based on the Transfer-Encoding header.
func (t *transferReader) parseTransferEncoding() error {
//line /usr/local/go/src/net/http/transfer.go:626
	_go_fuzz_dep_.CoverTab[43831]++
							raw, present := t.Header["Transfer-Encoding"]
							if !present {
//line /usr/local/go/src/net/http/transfer.go:628
		_go_fuzz_dep_.CoverTab[43836]++
								return nil
//line /usr/local/go/src/net/http/transfer.go:629
		// _ = "end of CoverTab[43836]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:630
		_go_fuzz_dep_.CoverTab[43837]++
//line /usr/local/go/src/net/http/transfer.go:630
		// _ = "end of CoverTab[43837]"
//line /usr/local/go/src/net/http/transfer.go:630
	}
//line /usr/local/go/src/net/http/transfer.go:630
	// _ = "end of CoverTab[43831]"
//line /usr/local/go/src/net/http/transfer.go:630
	_go_fuzz_dep_.CoverTab[43832]++
							delete(t.Header, "Transfer-Encoding")

//line /usr/local/go/src/net/http/transfer.go:634
	if !t.protoAtLeast(1, 1) {
//line /usr/local/go/src/net/http/transfer.go:634
		_go_fuzz_dep_.CoverTab[43838]++
								return nil
//line /usr/local/go/src/net/http/transfer.go:635
		// _ = "end of CoverTab[43838]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:636
		_go_fuzz_dep_.CoverTab[43839]++
//line /usr/local/go/src/net/http/transfer.go:636
		// _ = "end of CoverTab[43839]"
//line /usr/local/go/src/net/http/transfer.go:636
	}
//line /usr/local/go/src/net/http/transfer.go:636
	// _ = "end of CoverTab[43832]"
//line /usr/local/go/src/net/http/transfer.go:636
	_go_fuzz_dep_.CoverTab[43833]++

//line /usr/local/go/src/net/http/transfer.go:642
	if len(raw) != 1 {
//line /usr/local/go/src/net/http/transfer.go:642
		_go_fuzz_dep_.CoverTab[43840]++
								return &unsupportedTEError{fmt.Sprintf("too many transfer encodings: %q", raw)}
//line /usr/local/go/src/net/http/transfer.go:643
		// _ = "end of CoverTab[43840]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:644
		_go_fuzz_dep_.CoverTab[43841]++
//line /usr/local/go/src/net/http/transfer.go:644
		// _ = "end of CoverTab[43841]"
//line /usr/local/go/src/net/http/transfer.go:644
	}
//line /usr/local/go/src/net/http/transfer.go:644
	// _ = "end of CoverTab[43833]"
//line /usr/local/go/src/net/http/transfer.go:644
	_go_fuzz_dep_.CoverTab[43834]++
							if !ascii.EqualFold(raw[0], "chunked") {
//line /usr/local/go/src/net/http/transfer.go:645
		_go_fuzz_dep_.CoverTab[43842]++
								return &unsupportedTEError{fmt.Sprintf("unsupported transfer encoding: %q", raw[0])}
//line /usr/local/go/src/net/http/transfer.go:646
		// _ = "end of CoverTab[43842]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:647
		_go_fuzz_dep_.CoverTab[43843]++
//line /usr/local/go/src/net/http/transfer.go:647
		// _ = "end of CoverTab[43843]"
//line /usr/local/go/src/net/http/transfer.go:647
	}
//line /usr/local/go/src/net/http/transfer.go:647
	// _ = "end of CoverTab[43834]"
//line /usr/local/go/src/net/http/transfer.go:647
	_go_fuzz_dep_.CoverTab[43835]++

//line /usr/local/go/src/net/http/transfer.go:660
	delete(t.Header, "Content-Length")

							t.Chunked = true
							return nil
//line /usr/local/go/src/net/http/transfer.go:663
	// _ = "end of CoverTab[43835]"
}

// Determine the expected body length, using RFC 7230 Section 3.3. This
//line /usr/local/go/src/net/http/transfer.go:666
// function is not a method, because ultimately it should be shared by
//line /usr/local/go/src/net/http/transfer.go:666
// ReadResponse and ReadRequest.
//line /usr/local/go/src/net/http/transfer.go:669
func fixLength(isResponse bool, status int, requestMethod string, header Header, chunked bool) (int64, error) {
//line /usr/local/go/src/net/http/transfer.go:669
	_go_fuzz_dep_.CoverTab[43844]++
							isRequest := !isResponse
							contentLens := header["Content-Length"]

//line /usr/local/go/src/net/http/transfer.go:674
	if len(contentLens) > 1 {
//line /usr/local/go/src/net/http/transfer.go:674
		_go_fuzz_dep_.CoverTab[43853]++

//line /usr/local/go/src/net/http/transfer.go:679
		first := textproto.TrimString(contentLens[0])
		for _, ct := range contentLens[1:] {
//line /usr/local/go/src/net/http/transfer.go:680
			_go_fuzz_dep_.CoverTab[43855]++
									if first != textproto.TrimString(ct) {
//line /usr/local/go/src/net/http/transfer.go:681
				_go_fuzz_dep_.CoverTab[43856]++
										return 0, fmt.Errorf("http: message cannot contain multiple Content-Length headers; got %q", contentLens)
//line /usr/local/go/src/net/http/transfer.go:682
				// _ = "end of CoverTab[43856]"
			} else {
//line /usr/local/go/src/net/http/transfer.go:683
				_go_fuzz_dep_.CoverTab[43857]++
//line /usr/local/go/src/net/http/transfer.go:683
				// _ = "end of CoverTab[43857]"
//line /usr/local/go/src/net/http/transfer.go:683
			}
//line /usr/local/go/src/net/http/transfer.go:683
			// _ = "end of CoverTab[43855]"
		}
//line /usr/local/go/src/net/http/transfer.go:684
		// _ = "end of CoverTab[43853]"
//line /usr/local/go/src/net/http/transfer.go:684
		_go_fuzz_dep_.CoverTab[43854]++

//line /usr/local/go/src/net/http/transfer.go:687
		header.Del("Content-Length")
								header.Add("Content-Length", first)

								contentLens = header["Content-Length"]
//line /usr/local/go/src/net/http/transfer.go:690
		// _ = "end of CoverTab[43854]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:691
		_go_fuzz_dep_.CoverTab[43858]++
//line /usr/local/go/src/net/http/transfer.go:691
		// _ = "end of CoverTab[43858]"
//line /usr/local/go/src/net/http/transfer.go:691
	}
//line /usr/local/go/src/net/http/transfer.go:691
	// _ = "end of CoverTab[43844]"
//line /usr/local/go/src/net/http/transfer.go:691
	_go_fuzz_dep_.CoverTab[43845]++

//line /usr/local/go/src/net/http/transfer.go:694
	if isResponse && func() bool {
//line /usr/local/go/src/net/http/transfer.go:694
		_go_fuzz_dep_.CoverTab[43859]++
//line /usr/local/go/src/net/http/transfer.go:694
		return noResponseBodyExpected(requestMethod)
//line /usr/local/go/src/net/http/transfer.go:694
		// _ = "end of CoverTab[43859]"
//line /usr/local/go/src/net/http/transfer.go:694
	}() {
//line /usr/local/go/src/net/http/transfer.go:694
		_go_fuzz_dep_.CoverTab[43860]++
								return 0, nil
//line /usr/local/go/src/net/http/transfer.go:695
		// _ = "end of CoverTab[43860]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:696
		_go_fuzz_dep_.CoverTab[43861]++
//line /usr/local/go/src/net/http/transfer.go:696
		// _ = "end of CoverTab[43861]"
//line /usr/local/go/src/net/http/transfer.go:696
	}
//line /usr/local/go/src/net/http/transfer.go:696
	// _ = "end of CoverTab[43845]"
//line /usr/local/go/src/net/http/transfer.go:696
	_go_fuzz_dep_.CoverTab[43846]++
							if status/100 == 1 {
//line /usr/local/go/src/net/http/transfer.go:697
		_go_fuzz_dep_.CoverTab[43862]++
								return 0, nil
//line /usr/local/go/src/net/http/transfer.go:698
		// _ = "end of CoverTab[43862]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:699
		_go_fuzz_dep_.CoverTab[43863]++
//line /usr/local/go/src/net/http/transfer.go:699
		// _ = "end of CoverTab[43863]"
//line /usr/local/go/src/net/http/transfer.go:699
	}
//line /usr/local/go/src/net/http/transfer.go:699
	// _ = "end of CoverTab[43846]"
//line /usr/local/go/src/net/http/transfer.go:699
	_go_fuzz_dep_.CoverTab[43847]++
							switch status {
	case 204, 304:
//line /usr/local/go/src/net/http/transfer.go:701
		_go_fuzz_dep_.CoverTab[43864]++
								return 0, nil
//line /usr/local/go/src/net/http/transfer.go:702
		// _ = "end of CoverTab[43864]"
//line /usr/local/go/src/net/http/transfer.go:702
	default:
//line /usr/local/go/src/net/http/transfer.go:702
		_go_fuzz_dep_.CoverTab[43865]++
//line /usr/local/go/src/net/http/transfer.go:702
		// _ = "end of CoverTab[43865]"
	}
//line /usr/local/go/src/net/http/transfer.go:703
	// _ = "end of CoverTab[43847]"
//line /usr/local/go/src/net/http/transfer.go:703
	_go_fuzz_dep_.CoverTab[43848]++

//line /usr/local/go/src/net/http/transfer.go:706
	if chunked {
//line /usr/local/go/src/net/http/transfer.go:706
		_go_fuzz_dep_.CoverTab[43866]++
								return -1, nil
//line /usr/local/go/src/net/http/transfer.go:707
		// _ = "end of CoverTab[43866]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:708
		_go_fuzz_dep_.CoverTab[43867]++
//line /usr/local/go/src/net/http/transfer.go:708
		// _ = "end of CoverTab[43867]"
//line /usr/local/go/src/net/http/transfer.go:708
	}
//line /usr/local/go/src/net/http/transfer.go:708
	// _ = "end of CoverTab[43848]"
//line /usr/local/go/src/net/http/transfer.go:708
	_go_fuzz_dep_.CoverTab[43849]++

	// Logic based on Content-Length
	var cl string
	if len(contentLens) == 1 {
//line /usr/local/go/src/net/http/transfer.go:712
		_go_fuzz_dep_.CoverTab[43868]++
								cl = textproto.TrimString(contentLens[0])
//line /usr/local/go/src/net/http/transfer.go:713
		// _ = "end of CoverTab[43868]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:714
		_go_fuzz_dep_.CoverTab[43869]++
//line /usr/local/go/src/net/http/transfer.go:714
		// _ = "end of CoverTab[43869]"
//line /usr/local/go/src/net/http/transfer.go:714
	}
//line /usr/local/go/src/net/http/transfer.go:714
	// _ = "end of CoverTab[43849]"
//line /usr/local/go/src/net/http/transfer.go:714
	_go_fuzz_dep_.CoverTab[43850]++
							if cl != "" {
//line /usr/local/go/src/net/http/transfer.go:715
		_go_fuzz_dep_.CoverTab[43870]++
								n, err := parseContentLength(cl)
								if err != nil {
//line /usr/local/go/src/net/http/transfer.go:717
			_go_fuzz_dep_.CoverTab[43872]++
									return -1, err
//line /usr/local/go/src/net/http/transfer.go:718
			// _ = "end of CoverTab[43872]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:719
			_go_fuzz_dep_.CoverTab[43873]++
//line /usr/local/go/src/net/http/transfer.go:719
			// _ = "end of CoverTab[43873]"
//line /usr/local/go/src/net/http/transfer.go:719
		}
//line /usr/local/go/src/net/http/transfer.go:719
		// _ = "end of CoverTab[43870]"
//line /usr/local/go/src/net/http/transfer.go:719
		_go_fuzz_dep_.CoverTab[43871]++
								return n, nil
//line /usr/local/go/src/net/http/transfer.go:720
		// _ = "end of CoverTab[43871]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:721
		_go_fuzz_dep_.CoverTab[43874]++
//line /usr/local/go/src/net/http/transfer.go:721
		// _ = "end of CoverTab[43874]"
//line /usr/local/go/src/net/http/transfer.go:721
	}
//line /usr/local/go/src/net/http/transfer.go:721
	// _ = "end of CoverTab[43850]"
//line /usr/local/go/src/net/http/transfer.go:721
	_go_fuzz_dep_.CoverTab[43851]++
							header.Del("Content-Length")

							if isRequest {
//line /usr/local/go/src/net/http/transfer.go:724
		_go_fuzz_dep_.CoverTab[43875]++

//line /usr/local/go/src/net/http/transfer.go:732
		return 0, nil
//line /usr/local/go/src/net/http/transfer.go:732
		// _ = "end of CoverTab[43875]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:733
		_go_fuzz_dep_.CoverTab[43876]++
//line /usr/local/go/src/net/http/transfer.go:733
		// _ = "end of CoverTab[43876]"
//line /usr/local/go/src/net/http/transfer.go:733
	}
//line /usr/local/go/src/net/http/transfer.go:733
	// _ = "end of CoverTab[43851]"
//line /usr/local/go/src/net/http/transfer.go:733
	_go_fuzz_dep_.CoverTab[43852]++

//line /usr/local/go/src/net/http/transfer.go:736
	return -1, nil
//line /usr/local/go/src/net/http/transfer.go:736
	// _ = "end of CoverTab[43852]"
}

// Determine whether to hang up after sending a request and body, or
//line /usr/local/go/src/net/http/transfer.go:739
// receiving a response and body
//line /usr/local/go/src/net/http/transfer.go:739
// 'header' is the request headers.
//line /usr/local/go/src/net/http/transfer.go:742
func shouldClose(major, minor int, header Header, removeCloseHeader bool) bool {
//line /usr/local/go/src/net/http/transfer.go:742
	_go_fuzz_dep_.CoverTab[43877]++
							if major < 1 {
//line /usr/local/go/src/net/http/transfer.go:743
		_go_fuzz_dep_.CoverTab[43881]++
								return true
//line /usr/local/go/src/net/http/transfer.go:744
		// _ = "end of CoverTab[43881]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:745
		_go_fuzz_dep_.CoverTab[43882]++
//line /usr/local/go/src/net/http/transfer.go:745
		// _ = "end of CoverTab[43882]"
//line /usr/local/go/src/net/http/transfer.go:745
	}
//line /usr/local/go/src/net/http/transfer.go:745
	// _ = "end of CoverTab[43877]"
//line /usr/local/go/src/net/http/transfer.go:745
	_go_fuzz_dep_.CoverTab[43878]++

							conv := header["Connection"]
							hasClose := httpguts.HeaderValuesContainsToken(conv, "close")
							if major == 1 && func() bool {
//line /usr/local/go/src/net/http/transfer.go:749
		_go_fuzz_dep_.CoverTab[43883]++
//line /usr/local/go/src/net/http/transfer.go:749
		return minor == 0
//line /usr/local/go/src/net/http/transfer.go:749
		// _ = "end of CoverTab[43883]"
//line /usr/local/go/src/net/http/transfer.go:749
	}() {
//line /usr/local/go/src/net/http/transfer.go:749
		_go_fuzz_dep_.CoverTab[43884]++
								return hasClose || func() bool {
//line /usr/local/go/src/net/http/transfer.go:750
			_go_fuzz_dep_.CoverTab[43885]++
//line /usr/local/go/src/net/http/transfer.go:750
			return !httpguts.HeaderValuesContainsToken(conv, "keep-alive")
//line /usr/local/go/src/net/http/transfer.go:750
			// _ = "end of CoverTab[43885]"
//line /usr/local/go/src/net/http/transfer.go:750
		}()
//line /usr/local/go/src/net/http/transfer.go:750
		// _ = "end of CoverTab[43884]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:751
		_go_fuzz_dep_.CoverTab[43886]++
//line /usr/local/go/src/net/http/transfer.go:751
		// _ = "end of CoverTab[43886]"
//line /usr/local/go/src/net/http/transfer.go:751
	}
//line /usr/local/go/src/net/http/transfer.go:751
	// _ = "end of CoverTab[43878]"
//line /usr/local/go/src/net/http/transfer.go:751
	_go_fuzz_dep_.CoverTab[43879]++

							if hasClose && func() bool {
//line /usr/local/go/src/net/http/transfer.go:753
		_go_fuzz_dep_.CoverTab[43887]++
//line /usr/local/go/src/net/http/transfer.go:753
		return removeCloseHeader
//line /usr/local/go/src/net/http/transfer.go:753
		// _ = "end of CoverTab[43887]"
//line /usr/local/go/src/net/http/transfer.go:753
	}() {
//line /usr/local/go/src/net/http/transfer.go:753
		_go_fuzz_dep_.CoverTab[43888]++
								header.Del("Connection")
//line /usr/local/go/src/net/http/transfer.go:754
		// _ = "end of CoverTab[43888]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:755
		_go_fuzz_dep_.CoverTab[43889]++
//line /usr/local/go/src/net/http/transfer.go:755
		// _ = "end of CoverTab[43889]"
//line /usr/local/go/src/net/http/transfer.go:755
	}
//line /usr/local/go/src/net/http/transfer.go:755
	// _ = "end of CoverTab[43879]"
//line /usr/local/go/src/net/http/transfer.go:755
	_go_fuzz_dep_.CoverTab[43880]++

							return hasClose
//line /usr/local/go/src/net/http/transfer.go:757
	// _ = "end of CoverTab[43880]"
}

// Parse the trailer header.
func fixTrailer(header Header, chunked bool) (Header, error) {
//line /usr/local/go/src/net/http/transfer.go:761
	_go_fuzz_dep_.CoverTab[43890]++
							vv, ok := header["Trailer"]
							if !ok {
//line /usr/local/go/src/net/http/transfer.go:763
		_go_fuzz_dep_.CoverTab[43896]++
								return nil, nil
//line /usr/local/go/src/net/http/transfer.go:764
		// _ = "end of CoverTab[43896]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:765
		_go_fuzz_dep_.CoverTab[43897]++
//line /usr/local/go/src/net/http/transfer.go:765
		// _ = "end of CoverTab[43897]"
//line /usr/local/go/src/net/http/transfer.go:765
	}
//line /usr/local/go/src/net/http/transfer.go:765
	// _ = "end of CoverTab[43890]"
//line /usr/local/go/src/net/http/transfer.go:765
	_go_fuzz_dep_.CoverTab[43891]++
							if !chunked {
//line /usr/local/go/src/net/http/transfer.go:766
		_go_fuzz_dep_.CoverTab[43898]++

//line /usr/local/go/src/net/http/transfer.go:774
		return nil, nil
//line /usr/local/go/src/net/http/transfer.go:774
		// _ = "end of CoverTab[43898]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:775
		_go_fuzz_dep_.CoverTab[43899]++
//line /usr/local/go/src/net/http/transfer.go:775
		// _ = "end of CoverTab[43899]"
//line /usr/local/go/src/net/http/transfer.go:775
	}
//line /usr/local/go/src/net/http/transfer.go:775
	// _ = "end of CoverTab[43891]"
//line /usr/local/go/src/net/http/transfer.go:775
	_go_fuzz_dep_.CoverTab[43892]++
							header.Del("Trailer")

							trailer := make(Header)
							var err error
							for _, v := range vv {
//line /usr/local/go/src/net/http/transfer.go:780
		_go_fuzz_dep_.CoverTab[43900]++
								foreachHeaderElement(v, func(key string) {
//line /usr/local/go/src/net/http/transfer.go:781
			_go_fuzz_dep_.CoverTab[43901]++
									key = CanonicalHeaderKey(key)
									switch key {
			case "Transfer-Encoding", "Trailer", "Content-Length":
//line /usr/local/go/src/net/http/transfer.go:784
				_go_fuzz_dep_.CoverTab[43903]++
										if err == nil {
//line /usr/local/go/src/net/http/transfer.go:785
					_go_fuzz_dep_.CoverTab[43905]++
											err = badStringError("bad trailer key", key)
											return
//line /usr/local/go/src/net/http/transfer.go:787
					// _ = "end of CoverTab[43905]"
				} else {
//line /usr/local/go/src/net/http/transfer.go:788
					_go_fuzz_dep_.CoverTab[43906]++
//line /usr/local/go/src/net/http/transfer.go:788
					// _ = "end of CoverTab[43906]"
//line /usr/local/go/src/net/http/transfer.go:788
				}
//line /usr/local/go/src/net/http/transfer.go:788
				// _ = "end of CoverTab[43903]"
//line /usr/local/go/src/net/http/transfer.go:788
			default:
//line /usr/local/go/src/net/http/transfer.go:788
				_go_fuzz_dep_.CoverTab[43904]++
//line /usr/local/go/src/net/http/transfer.go:788
				// _ = "end of CoverTab[43904]"
			}
//line /usr/local/go/src/net/http/transfer.go:789
			// _ = "end of CoverTab[43901]"
//line /usr/local/go/src/net/http/transfer.go:789
			_go_fuzz_dep_.CoverTab[43902]++
									trailer[key] = nil
//line /usr/local/go/src/net/http/transfer.go:790
			// _ = "end of CoverTab[43902]"
		})
//line /usr/local/go/src/net/http/transfer.go:791
		// _ = "end of CoverTab[43900]"
	}
//line /usr/local/go/src/net/http/transfer.go:792
	// _ = "end of CoverTab[43892]"
//line /usr/local/go/src/net/http/transfer.go:792
	_go_fuzz_dep_.CoverTab[43893]++
							if err != nil {
//line /usr/local/go/src/net/http/transfer.go:793
		_go_fuzz_dep_.CoverTab[43907]++
								return nil, err
//line /usr/local/go/src/net/http/transfer.go:794
		// _ = "end of CoverTab[43907]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:795
		_go_fuzz_dep_.CoverTab[43908]++
//line /usr/local/go/src/net/http/transfer.go:795
		// _ = "end of CoverTab[43908]"
//line /usr/local/go/src/net/http/transfer.go:795
	}
//line /usr/local/go/src/net/http/transfer.go:795
	// _ = "end of CoverTab[43893]"
//line /usr/local/go/src/net/http/transfer.go:795
	_go_fuzz_dep_.CoverTab[43894]++
							if len(trailer) == 0 {
//line /usr/local/go/src/net/http/transfer.go:796
		_go_fuzz_dep_.CoverTab[43909]++
								return nil, nil
//line /usr/local/go/src/net/http/transfer.go:797
		// _ = "end of CoverTab[43909]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:798
		_go_fuzz_dep_.CoverTab[43910]++
//line /usr/local/go/src/net/http/transfer.go:798
		// _ = "end of CoverTab[43910]"
//line /usr/local/go/src/net/http/transfer.go:798
	}
//line /usr/local/go/src/net/http/transfer.go:798
	// _ = "end of CoverTab[43894]"
//line /usr/local/go/src/net/http/transfer.go:798
	_go_fuzz_dep_.CoverTab[43895]++
							return trailer, nil
//line /usr/local/go/src/net/http/transfer.go:799
	// _ = "end of CoverTab[43895]"
}

// body turns a Reader into a ReadCloser.
//line /usr/local/go/src/net/http/transfer.go:802
// Close ensures that the body has been fully read
//line /usr/local/go/src/net/http/transfer.go:802
// and then reads the trailer if necessary.
//line /usr/local/go/src/net/http/transfer.go:805
type body struct {
	src		io.Reader
	hdr		any		// non-nil (Response or Request) value means read trailer
	r		*bufio.Reader	// underlying wire-format reader for the trailer
	closing		bool		// is the connection to be closed after reading body?
	doEarlyClose	bool		// whether Close should stop early

	mu		sync.Mutex	// guards following, and calls to Read and Close
	sawEOF		bool
	closed		bool
	earlyClose	bool	// Close called and we didn't read to the end of src
	onHitEOF	func()	// if non-nil, func to call when EOF is Read
}

// ErrBodyReadAfterClose is returned when reading a Request or Response
//line /usr/local/go/src/net/http/transfer.go:819
// Body after the body has been closed. This typically happens when the body is
//line /usr/local/go/src/net/http/transfer.go:819
// read after an HTTP Handler calls WriteHeader or Write on its
//line /usr/local/go/src/net/http/transfer.go:819
// ResponseWriter.
//line /usr/local/go/src/net/http/transfer.go:823
var ErrBodyReadAfterClose = errors.New("http: invalid Read on closed Body")

func (b *body) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/transfer.go:825
	_go_fuzz_dep_.CoverTab[43911]++
							b.mu.Lock()
							defer b.mu.Unlock()
							if b.closed {
//line /usr/local/go/src/net/http/transfer.go:828
		_go_fuzz_dep_.CoverTab[43913]++
								return 0, ErrBodyReadAfterClose
//line /usr/local/go/src/net/http/transfer.go:829
		// _ = "end of CoverTab[43913]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:830
		_go_fuzz_dep_.CoverTab[43914]++
//line /usr/local/go/src/net/http/transfer.go:830
		// _ = "end of CoverTab[43914]"
//line /usr/local/go/src/net/http/transfer.go:830
	}
//line /usr/local/go/src/net/http/transfer.go:830
	// _ = "end of CoverTab[43911]"
//line /usr/local/go/src/net/http/transfer.go:830
	_go_fuzz_dep_.CoverTab[43912]++
							return b.readLocked(p)
//line /usr/local/go/src/net/http/transfer.go:831
	// _ = "end of CoverTab[43912]"
}

// Must hold b.mu.
func (b *body) readLocked(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/transfer.go:835
	_go_fuzz_dep_.CoverTab[43915]++
							if b.sawEOF {
//line /usr/local/go/src/net/http/transfer.go:836
		_go_fuzz_dep_.CoverTab[43920]++
								return 0, io.EOF
//line /usr/local/go/src/net/http/transfer.go:837
		// _ = "end of CoverTab[43920]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:838
		_go_fuzz_dep_.CoverTab[43921]++
//line /usr/local/go/src/net/http/transfer.go:838
		// _ = "end of CoverTab[43921]"
//line /usr/local/go/src/net/http/transfer.go:838
	}
//line /usr/local/go/src/net/http/transfer.go:838
	// _ = "end of CoverTab[43915]"
//line /usr/local/go/src/net/http/transfer.go:838
	_go_fuzz_dep_.CoverTab[43916]++
							n, err = b.src.Read(p)

							if err == io.EOF {
//line /usr/local/go/src/net/http/transfer.go:841
		_go_fuzz_dep_.CoverTab[43922]++
								b.sawEOF = true

								if b.hdr != nil {
//line /usr/local/go/src/net/http/transfer.go:844
			_go_fuzz_dep_.CoverTab[43923]++
									if e := b.readTrailer(); e != nil {
//line /usr/local/go/src/net/http/transfer.go:845
				_go_fuzz_dep_.CoverTab[43925]++
										err = e

//line /usr/local/go/src/net/http/transfer.go:851
				b.sawEOF = false
										b.closed = true
//line /usr/local/go/src/net/http/transfer.go:852
				// _ = "end of CoverTab[43925]"
			} else {
//line /usr/local/go/src/net/http/transfer.go:853
				_go_fuzz_dep_.CoverTab[43926]++
//line /usr/local/go/src/net/http/transfer.go:853
				// _ = "end of CoverTab[43926]"
//line /usr/local/go/src/net/http/transfer.go:853
			}
//line /usr/local/go/src/net/http/transfer.go:853
			// _ = "end of CoverTab[43923]"
//line /usr/local/go/src/net/http/transfer.go:853
			_go_fuzz_dep_.CoverTab[43924]++
									b.hdr = nil
//line /usr/local/go/src/net/http/transfer.go:854
			// _ = "end of CoverTab[43924]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:855
			_go_fuzz_dep_.CoverTab[43927]++

//line /usr/local/go/src/net/http/transfer.go:858
			if lr, ok := b.src.(*io.LimitedReader); ok && func() bool {
//line /usr/local/go/src/net/http/transfer.go:858
				_go_fuzz_dep_.CoverTab[43928]++
//line /usr/local/go/src/net/http/transfer.go:858
				return lr.N > 0
//line /usr/local/go/src/net/http/transfer.go:858
				// _ = "end of CoverTab[43928]"
//line /usr/local/go/src/net/http/transfer.go:858
			}() {
//line /usr/local/go/src/net/http/transfer.go:858
				_go_fuzz_dep_.CoverTab[43929]++
										err = io.ErrUnexpectedEOF
//line /usr/local/go/src/net/http/transfer.go:859
				// _ = "end of CoverTab[43929]"
			} else {
//line /usr/local/go/src/net/http/transfer.go:860
				_go_fuzz_dep_.CoverTab[43930]++
//line /usr/local/go/src/net/http/transfer.go:860
				// _ = "end of CoverTab[43930]"
//line /usr/local/go/src/net/http/transfer.go:860
			}
//line /usr/local/go/src/net/http/transfer.go:860
			// _ = "end of CoverTab[43927]"
		}
//line /usr/local/go/src/net/http/transfer.go:861
		// _ = "end of CoverTab[43922]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:862
		_go_fuzz_dep_.CoverTab[43931]++
//line /usr/local/go/src/net/http/transfer.go:862
		// _ = "end of CoverTab[43931]"
//line /usr/local/go/src/net/http/transfer.go:862
	}
//line /usr/local/go/src/net/http/transfer.go:862
	// _ = "end of CoverTab[43916]"
//line /usr/local/go/src/net/http/transfer.go:862
	_go_fuzz_dep_.CoverTab[43917]++

//line /usr/local/go/src/net/http/transfer.go:869
	if err == nil && func() bool {
//line /usr/local/go/src/net/http/transfer.go:869
		_go_fuzz_dep_.CoverTab[43932]++
//line /usr/local/go/src/net/http/transfer.go:869
		return n > 0
//line /usr/local/go/src/net/http/transfer.go:869
		// _ = "end of CoverTab[43932]"
//line /usr/local/go/src/net/http/transfer.go:869
	}() {
//line /usr/local/go/src/net/http/transfer.go:869
		_go_fuzz_dep_.CoverTab[43933]++
								if lr, ok := b.src.(*io.LimitedReader); ok && func() bool {
//line /usr/local/go/src/net/http/transfer.go:870
			_go_fuzz_dep_.CoverTab[43934]++
//line /usr/local/go/src/net/http/transfer.go:870
			return lr.N == 0
//line /usr/local/go/src/net/http/transfer.go:870
			// _ = "end of CoverTab[43934]"
//line /usr/local/go/src/net/http/transfer.go:870
		}() {
//line /usr/local/go/src/net/http/transfer.go:870
			_go_fuzz_dep_.CoverTab[43935]++
									err = io.EOF
									b.sawEOF = true
//line /usr/local/go/src/net/http/transfer.go:872
			// _ = "end of CoverTab[43935]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:873
			_go_fuzz_dep_.CoverTab[43936]++
//line /usr/local/go/src/net/http/transfer.go:873
			// _ = "end of CoverTab[43936]"
//line /usr/local/go/src/net/http/transfer.go:873
		}
//line /usr/local/go/src/net/http/transfer.go:873
		// _ = "end of CoverTab[43933]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:874
		_go_fuzz_dep_.CoverTab[43937]++
//line /usr/local/go/src/net/http/transfer.go:874
		// _ = "end of CoverTab[43937]"
//line /usr/local/go/src/net/http/transfer.go:874
	}
//line /usr/local/go/src/net/http/transfer.go:874
	// _ = "end of CoverTab[43917]"
//line /usr/local/go/src/net/http/transfer.go:874
	_go_fuzz_dep_.CoverTab[43918]++

							if b.sawEOF && func() bool {
//line /usr/local/go/src/net/http/transfer.go:876
		_go_fuzz_dep_.CoverTab[43938]++
//line /usr/local/go/src/net/http/transfer.go:876
		return b.onHitEOF != nil
//line /usr/local/go/src/net/http/transfer.go:876
		// _ = "end of CoverTab[43938]"
//line /usr/local/go/src/net/http/transfer.go:876
	}() {
//line /usr/local/go/src/net/http/transfer.go:876
		_go_fuzz_dep_.CoverTab[43939]++
								b.onHitEOF()
//line /usr/local/go/src/net/http/transfer.go:877
		// _ = "end of CoverTab[43939]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:878
		_go_fuzz_dep_.CoverTab[43940]++
//line /usr/local/go/src/net/http/transfer.go:878
		// _ = "end of CoverTab[43940]"
//line /usr/local/go/src/net/http/transfer.go:878
	}
//line /usr/local/go/src/net/http/transfer.go:878
	// _ = "end of CoverTab[43918]"
//line /usr/local/go/src/net/http/transfer.go:878
	_go_fuzz_dep_.CoverTab[43919]++

							return n, err
//line /usr/local/go/src/net/http/transfer.go:880
	// _ = "end of CoverTab[43919]"
}

var (
	singleCRLF	= []byte("\r\n")
	doubleCRLF	= []byte("\r\n\r\n")
)

func seeUpcomingDoubleCRLF(r *bufio.Reader) bool {
//line /usr/local/go/src/net/http/transfer.go:888
	_go_fuzz_dep_.CoverTab[43941]++
							for peekSize := 4; ; peekSize++ {
//line /usr/local/go/src/net/http/transfer.go:889
		_go_fuzz_dep_.CoverTab[43943]++

//line /usr/local/go/src/net/http/transfer.go:892
		buf, err := r.Peek(peekSize)
		if bytes.HasSuffix(buf, doubleCRLF) {
//line /usr/local/go/src/net/http/transfer.go:893
			_go_fuzz_dep_.CoverTab[43945]++
									return true
//line /usr/local/go/src/net/http/transfer.go:894
			// _ = "end of CoverTab[43945]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:895
			_go_fuzz_dep_.CoverTab[43946]++
//line /usr/local/go/src/net/http/transfer.go:895
			// _ = "end of CoverTab[43946]"
//line /usr/local/go/src/net/http/transfer.go:895
		}
//line /usr/local/go/src/net/http/transfer.go:895
		// _ = "end of CoverTab[43943]"
//line /usr/local/go/src/net/http/transfer.go:895
		_go_fuzz_dep_.CoverTab[43944]++
								if err != nil {
//line /usr/local/go/src/net/http/transfer.go:896
			_go_fuzz_dep_.CoverTab[43947]++
									break
//line /usr/local/go/src/net/http/transfer.go:897
			// _ = "end of CoverTab[43947]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:898
			_go_fuzz_dep_.CoverTab[43948]++
//line /usr/local/go/src/net/http/transfer.go:898
			// _ = "end of CoverTab[43948]"
//line /usr/local/go/src/net/http/transfer.go:898
		}
//line /usr/local/go/src/net/http/transfer.go:898
		// _ = "end of CoverTab[43944]"
	}
//line /usr/local/go/src/net/http/transfer.go:899
	// _ = "end of CoverTab[43941]"
//line /usr/local/go/src/net/http/transfer.go:899
	_go_fuzz_dep_.CoverTab[43942]++
							return false
//line /usr/local/go/src/net/http/transfer.go:900
	// _ = "end of CoverTab[43942]"
}

var errTrailerEOF = errors.New("http: unexpected EOF reading trailer")

func (b *body) readTrailer() error {
//line /usr/local/go/src/net/http/transfer.go:905
	_go_fuzz_dep_.CoverTab[43949]++

							buf, err := b.r.Peek(2)
							if bytes.Equal(buf, singleCRLF) {
//line /usr/local/go/src/net/http/transfer.go:908
		_go_fuzz_dep_.CoverTab[43956]++
								b.r.Discard(2)
								return nil
//line /usr/local/go/src/net/http/transfer.go:910
		// _ = "end of CoverTab[43956]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:911
		_go_fuzz_dep_.CoverTab[43957]++
//line /usr/local/go/src/net/http/transfer.go:911
		// _ = "end of CoverTab[43957]"
//line /usr/local/go/src/net/http/transfer.go:911
	}
//line /usr/local/go/src/net/http/transfer.go:911
	// _ = "end of CoverTab[43949]"
//line /usr/local/go/src/net/http/transfer.go:911
	_go_fuzz_dep_.CoverTab[43950]++
							if len(buf) < 2 {
//line /usr/local/go/src/net/http/transfer.go:912
		_go_fuzz_dep_.CoverTab[43958]++
								return errTrailerEOF
//line /usr/local/go/src/net/http/transfer.go:913
		// _ = "end of CoverTab[43958]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:914
		_go_fuzz_dep_.CoverTab[43959]++
//line /usr/local/go/src/net/http/transfer.go:914
		// _ = "end of CoverTab[43959]"
//line /usr/local/go/src/net/http/transfer.go:914
	}
//line /usr/local/go/src/net/http/transfer.go:914
	// _ = "end of CoverTab[43950]"
//line /usr/local/go/src/net/http/transfer.go:914
	_go_fuzz_dep_.CoverTab[43951]++
							if err != nil {
//line /usr/local/go/src/net/http/transfer.go:915
		_go_fuzz_dep_.CoverTab[43960]++
								return err
//line /usr/local/go/src/net/http/transfer.go:916
		// _ = "end of CoverTab[43960]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:917
		_go_fuzz_dep_.CoverTab[43961]++
//line /usr/local/go/src/net/http/transfer.go:917
		// _ = "end of CoverTab[43961]"
//line /usr/local/go/src/net/http/transfer.go:917
	}
//line /usr/local/go/src/net/http/transfer.go:917
	// _ = "end of CoverTab[43951]"
//line /usr/local/go/src/net/http/transfer.go:917
	_go_fuzz_dep_.CoverTab[43952]++

//line /usr/local/go/src/net/http/transfer.go:927
	if !seeUpcomingDoubleCRLF(b.r) {
//line /usr/local/go/src/net/http/transfer.go:927
		_go_fuzz_dep_.CoverTab[43962]++
								return errors.New("http: suspiciously long trailer after chunked body")
//line /usr/local/go/src/net/http/transfer.go:928
		// _ = "end of CoverTab[43962]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:929
		_go_fuzz_dep_.CoverTab[43963]++
//line /usr/local/go/src/net/http/transfer.go:929
		// _ = "end of CoverTab[43963]"
//line /usr/local/go/src/net/http/transfer.go:929
	}
//line /usr/local/go/src/net/http/transfer.go:929
	// _ = "end of CoverTab[43952]"
//line /usr/local/go/src/net/http/transfer.go:929
	_go_fuzz_dep_.CoverTab[43953]++

							hdr, err := textproto.NewReader(b.r).ReadMIMEHeader()
							if err != nil {
//line /usr/local/go/src/net/http/transfer.go:932
		_go_fuzz_dep_.CoverTab[43964]++
								if err == io.EOF {
//line /usr/local/go/src/net/http/transfer.go:933
			_go_fuzz_dep_.CoverTab[43966]++
									return errTrailerEOF
//line /usr/local/go/src/net/http/transfer.go:934
			// _ = "end of CoverTab[43966]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:935
			_go_fuzz_dep_.CoverTab[43967]++
//line /usr/local/go/src/net/http/transfer.go:935
			// _ = "end of CoverTab[43967]"
//line /usr/local/go/src/net/http/transfer.go:935
		}
//line /usr/local/go/src/net/http/transfer.go:935
		// _ = "end of CoverTab[43964]"
//line /usr/local/go/src/net/http/transfer.go:935
		_go_fuzz_dep_.CoverTab[43965]++
								return err
//line /usr/local/go/src/net/http/transfer.go:936
		// _ = "end of CoverTab[43965]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:937
		_go_fuzz_dep_.CoverTab[43968]++
//line /usr/local/go/src/net/http/transfer.go:937
		// _ = "end of CoverTab[43968]"
//line /usr/local/go/src/net/http/transfer.go:937
	}
//line /usr/local/go/src/net/http/transfer.go:937
	// _ = "end of CoverTab[43953]"
//line /usr/local/go/src/net/http/transfer.go:937
	_go_fuzz_dep_.CoverTab[43954]++
							switch rr := b.hdr.(type) {
	case *Request:
//line /usr/local/go/src/net/http/transfer.go:939
		_go_fuzz_dep_.CoverTab[43969]++
								mergeSetHeader(&rr.Trailer, Header(hdr))
//line /usr/local/go/src/net/http/transfer.go:940
		// _ = "end of CoverTab[43969]"
	case *Response:
//line /usr/local/go/src/net/http/transfer.go:941
		_go_fuzz_dep_.CoverTab[43970]++
								mergeSetHeader(&rr.Trailer, Header(hdr))
//line /usr/local/go/src/net/http/transfer.go:942
		// _ = "end of CoverTab[43970]"
	}
//line /usr/local/go/src/net/http/transfer.go:943
	// _ = "end of CoverTab[43954]"
//line /usr/local/go/src/net/http/transfer.go:943
	_go_fuzz_dep_.CoverTab[43955]++
							return nil
//line /usr/local/go/src/net/http/transfer.go:944
	// _ = "end of CoverTab[43955]"
}

func mergeSetHeader(dst *Header, src Header) {
//line /usr/local/go/src/net/http/transfer.go:947
	_go_fuzz_dep_.CoverTab[43971]++
							if *dst == nil {
//line /usr/local/go/src/net/http/transfer.go:948
		_go_fuzz_dep_.CoverTab[43973]++
								*dst = src
								return
//line /usr/local/go/src/net/http/transfer.go:950
		// _ = "end of CoverTab[43973]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:951
		_go_fuzz_dep_.CoverTab[43974]++
//line /usr/local/go/src/net/http/transfer.go:951
		// _ = "end of CoverTab[43974]"
//line /usr/local/go/src/net/http/transfer.go:951
	}
//line /usr/local/go/src/net/http/transfer.go:951
	// _ = "end of CoverTab[43971]"
//line /usr/local/go/src/net/http/transfer.go:951
	_go_fuzz_dep_.CoverTab[43972]++
							for k, vv := range src {
//line /usr/local/go/src/net/http/transfer.go:952
		_go_fuzz_dep_.CoverTab[43975]++
								(*dst)[k] = vv
//line /usr/local/go/src/net/http/transfer.go:953
		// _ = "end of CoverTab[43975]"
	}
//line /usr/local/go/src/net/http/transfer.go:954
	// _ = "end of CoverTab[43972]"
}

// unreadDataSizeLocked returns the number of bytes of unread input.
//line /usr/local/go/src/net/http/transfer.go:957
// It returns -1 if unknown.
//line /usr/local/go/src/net/http/transfer.go:957
// b.mu must be held.
//line /usr/local/go/src/net/http/transfer.go:960
func (b *body) unreadDataSizeLocked() int64 {
//line /usr/local/go/src/net/http/transfer.go:960
	_go_fuzz_dep_.CoverTab[43976]++
							if lr, ok := b.src.(*io.LimitedReader); ok {
//line /usr/local/go/src/net/http/transfer.go:961
		_go_fuzz_dep_.CoverTab[43978]++
								return lr.N
//line /usr/local/go/src/net/http/transfer.go:962
		// _ = "end of CoverTab[43978]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:963
		_go_fuzz_dep_.CoverTab[43979]++
//line /usr/local/go/src/net/http/transfer.go:963
		// _ = "end of CoverTab[43979]"
//line /usr/local/go/src/net/http/transfer.go:963
	}
//line /usr/local/go/src/net/http/transfer.go:963
	// _ = "end of CoverTab[43976]"
//line /usr/local/go/src/net/http/transfer.go:963
	_go_fuzz_dep_.CoverTab[43977]++
							return -1
//line /usr/local/go/src/net/http/transfer.go:964
	// _ = "end of CoverTab[43977]"
}

func (b *body) Close() error {
//line /usr/local/go/src/net/http/transfer.go:967
	_go_fuzz_dep_.CoverTab[43980]++
							b.mu.Lock()
							defer b.mu.Unlock()
							if b.closed {
//line /usr/local/go/src/net/http/transfer.go:970
		_go_fuzz_dep_.CoverTab[43983]++
								return nil
//line /usr/local/go/src/net/http/transfer.go:971
		// _ = "end of CoverTab[43983]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:972
		_go_fuzz_dep_.CoverTab[43984]++
//line /usr/local/go/src/net/http/transfer.go:972
		// _ = "end of CoverTab[43984]"
//line /usr/local/go/src/net/http/transfer.go:972
	}
//line /usr/local/go/src/net/http/transfer.go:972
	// _ = "end of CoverTab[43980]"
//line /usr/local/go/src/net/http/transfer.go:972
	_go_fuzz_dep_.CoverTab[43981]++
							var err error
							switch {
	case b.sawEOF:
//line /usr/local/go/src/net/http/transfer.go:975
		_go_fuzz_dep_.CoverTab[43985]++
//line /usr/local/go/src/net/http/transfer.go:975
		// _ = "end of CoverTab[43985]"

	case b.hdr == nil && func() bool {
//line /usr/local/go/src/net/http/transfer.go:977
		_go_fuzz_dep_.CoverTab[43989]++
//line /usr/local/go/src/net/http/transfer.go:977
		return b.closing
//line /usr/local/go/src/net/http/transfer.go:977
		// _ = "end of CoverTab[43989]"
//line /usr/local/go/src/net/http/transfer.go:977
	}():
//line /usr/local/go/src/net/http/transfer.go:977
		_go_fuzz_dep_.CoverTab[43986]++
//line /usr/local/go/src/net/http/transfer.go:977
		// _ = "end of CoverTab[43986]"

//line /usr/local/go/src/net/http/transfer.go:980
	case b.doEarlyClose:
//line /usr/local/go/src/net/http/transfer.go:980
		_go_fuzz_dep_.CoverTab[43987]++

//line /usr/local/go/src/net/http/transfer.go:983
		if lr, ok := b.src.(*io.LimitedReader); ok && func() bool {
//line /usr/local/go/src/net/http/transfer.go:983
			_go_fuzz_dep_.CoverTab[43990]++
//line /usr/local/go/src/net/http/transfer.go:983
			return lr.N > maxPostHandlerReadBytes
//line /usr/local/go/src/net/http/transfer.go:983
			// _ = "end of CoverTab[43990]"
//line /usr/local/go/src/net/http/transfer.go:983
		}() {
//line /usr/local/go/src/net/http/transfer.go:983
			_go_fuzz_dep_.CoverTab[43991]++

//line /usr/local/go/src/net/http/transfer.go:986
			b.earlyClose = true
//line /usr/local/go/src/net/http/transfer.go:986
			// _ = "end of CoverTab[43991]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:987
			_go_fuzz_dep_.CoverTab[43992]++
									var n int64

//line /usr/local/go/src/net/http/transfer.go:991
			n, err = io.CopyN(io.Discard, bodyLocked{b}, maxPostHandlerReadBytes)
			if err == io.EOF {
//line /usr/local/go/src/net/http/transfer.go:992
				_go_fuzz_dep_.CoverTab[43994]++
										err = nil
//line /usr/local/go/src/net/http/transfer.go:993
				// _ = "end of CoverTab[43994]"
			} else {
//line /usr/local/go/src/net/http/transfer.go:994
				_go_fuzz_dep_.CoverTab[43995]++
//line /usr/local/go/src/net/http/transfer.go:994
				// _ = "end of CoverTab[43995]"
//line /usr/local/go/src/net/http/transfer.go:994
			}
//line /usr/local/go/src/net/http/transfer.go:994
			// _ = "end of CoverTab[43992]"
//line /usr/local/go/src/net/http/transfer.go:994
			_go_fuzz_dep_.CoverTab[43993]++
									if n == maxPostHandlerReadBytes {
//line /usr/local/go/src/net/http/transfer.go:995
				_go_fuzz_dep_.CoverTab[43996]++
										b.earlyClose = true
//line /usr/local/go/src/net/http/transfer.go:996
				// _ = "end of CoverTab[43996]"
			} else {
//line /usr/local/go/src/net/http/transfer.go:997
				_go_fuzz_dep_.CoverTab[43997]++
//line /usr/local/go/src/net/http/transfer.go:997
				// _ = "end of CoverTab[43997]"
//line /usr/local/go/src/net/http/transfer.go:997
			}
//line /usr/local/go/src/net/http/transfer.go:997
			// _ = "end of CoverTab[43993]"
		}
//line /usr/local/go/src/net/http/transfer.go:998
		// _ = "end of CoverTab[43987]"
	default:
//line /usr/local/go/src/net/http/transfer.go:999
		_go_fuzz_dep_.CoverTab[43988]++

//line /usr/local/go/src/net/http/transfer.go:1002
		_, err = io.Copy(io.Discard, bodyLocked{b})
//line /usr/local/go/src/net/http/transfer.go:1002
		// _ = "end of CoverTab[43988]"
	}
//line /usr/local/go/src/net/http/transfer.go:1003
	// _ = "end of CoverTab[43981]"
//line /usr/local/go/src/net/http/transfer.go:1003
	_go_fuzz_dep_.CoverTab[43982]++
							b.closed = true
							return err
//line /usr/local/go/src/net/http/transfer.go:1005
	// _ = "end of CoverTab[43982]"
}

func (b *body) didEarlyClose() bool {
//line /usr/local/go/src/net/http/transfer.go:1008
	_go_fuzz_dep_.CoverTab[43998]++
							b.mu.Lock()
							defer b.mu.Unlock()
							return b.earlyClose
//line /usr/local/go/src/net/http/transfer.go:1011
	// _ = "end of CoverTab[43998]"
}

// bodyRemains reports whether future Read calls might
//line /usr/local/go/src/net/http/transfer.go:1014
// yield data.
//line /usr/local/go/src/net/http/transfer.go:1016
func (b *body) bodyRemains() bool {
//line /usr/local/go/src/net/http/transfer.go:1016
	_go_fuzz_dep_.CoverTab[43999]++
							b.mu.Lock()
							defer b.mu.Unlock()
							return !b.sawEOF
//line /usr/local/go/src/net/http/transfer.go:1019
	// _ = "end of CoverTab[43999]"
}

func (b *body) registerOnHitEOF(fn func()) {
//line /usr/local/go/src/net/http/transfer.go:1022
	_go_fuzz_dep_.CoverTab[44000]++
							b.mu.Lock()
							defer b.mu.Unlock()
							b.onHitEOF = fn
//line /usr/local/go/src/net/http/transfer.go:1025
	// _ = "end of CoverTab[44000]"
}

// bodyLocked is an io.Reader reading from a *body when its mutex is
//line /usr/local/go/src/net/http/transfer.go:1028
// already held.
//line /usr/local/go/src/net/http/transfer.go:1030
type bodyLocked struct {
	b *body
}

func (bl bodyLocked) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/transfer.go:1034
	_go_fuzz_dep_.CoverTab[44001]++
							if bl.b.closed {
//line /usr/local/go/src/net/http/transfer.go:1035
		_go_fuzz_dep_.CoverTab[44003]++
								return 0, ErrBodyReadAfterClose
//line /usr/local/go/src/net/http/transfer.go:1036
		// _ = "end of CoverTab[44003]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:1037
		_go_fuzz_dep_.CoverTab[44004]++
//line /usr/local/go/src/net/http/transfer.go:1037
		// _ = "end of CoverTab[44004]"
//line /usr/local/go/src/net/http/transfer.go:1037
	}
//line /usr/local/go/src/net/http/transfer.go:1037
	// _ = "end of CoverTab[44001]"
//line /usr/local/go/src/net/http/transfer.go:1037
	_go_fuzz_dep_.CoverTab[44002]++
							return bl.b.readLocked(p)
//line /usr/local/go/src/net/http/transfer.go:1038
	// _ = "end of CoverTab[44002]"
}

// parseContentLength trims whitespace from s and returns -1 if no value
//line /usr/local/go/src/net/http/transfer.go:1041
// is set, or the value if it's >= 0.
//line /usr/local/go/src/net/http/transfer.go:1043
func parseContentLength(cl string) (int64, error) {
//line /usr/local/go/src/net/http/transfer.go:1043
	_go_fuzz_dep_.CoverTab[44005]++
							cl = textproto.TrimString(cl)
							if cl == "" {
//line /usr/local/go/src/net/http/transfer.go:1045
		_go_fuzz_dep_.CoverTab[44008]++
								return -1, nil
//line /usr/local/go/src/net/http/transfer.go:1046
		// _ = "end of CoverTab[44008]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:1047
		_go_fuzz_dep_.CoverTab[44009]++
//line /usr/local/go/src/net/http/transfer.go:1047
		// _ = "end of CoverTab[44009]"
//line /usr/local/go/src/net/http/transfer.go:1047
	}
//line /usr/local/go/src/net/http/transfer.go:1047
	// _ = "end of CoverTab[44005]"
//line /usr/local/go/src/net/http/transfer.go:1047
	_go_fuzz_dep_.CoverTab[44006]++
							n, err := strconv.ParseUint(cl, 10, 63)
							if err != nil {
//line /usr/local/go/src/net/http/transfer.go:1049
		_go_fuzz_dep_.CoverTab[44010]++
								return 0, badStringError("bad Content-Length", cl)
//line /usr/local/go/src/net/http/transfer.go:1050
		// _ = "end of CoverTab[44010]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:1051
		_go_fuzz_dep_.CoverTab[44011]++
//line /usr/local/go/src/net/http/transfer.go:1051
		// _ = "end of CoverTab[44011]"
//line /usr/local/go/src/net/http/transfer.go:1051
	}
//line /usr/local/go/src/net/http/transfer.go:1051
	// _ = "end of CoverTab[44006]"
//line /usr/local/go/src/net/http/transfer.go:1051
	_go_fuzz_dep_.CoverTab[44007]++
							return int64(n), nil
//line /usr/local/go/src/net/http/transfer.go:1052
	// _ = "end of CoverTab[44007]"

}

// finishAsyncByteRead finishes reading the 1-byte sniff
//line /usr/local/go/src/net/http/transfer.go:1056
// from the ContentLength==0, Body!=nil case.
//line /usr/local/go/src/net/http/transfer.go:1058
type finishAsyncByteRead struct {
	tw *transferWriter
}

func (fr finishAsyncByteRead) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/transfer.go:1062
	_go_fuzz_dep_.CoverTab[44012]++
							if len(p) == 0 {
//line /usr/local/go/src/net/http/transfer.go:1063
		_go_fuzz_dep_.CoverTab[44016]++
								return
//line /usr/local/go/src/net/http/transfer.go:1064
		// _ = "end of CoverTab[44016]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:1065
		_go_fuzz_dep_.CoverTab[44017]++
//line /usr/local/go/src/net/http/transfer.go:1065
		// _ = "end of CoverTab[44017]"
//line /usr/local/go/src/net/http/transfer.go:1065
	}
//line /usr/local/go/src/net/http/transfer.go:1065
	// _ = "end of CoverTab[44012]"
//line /usr/local/go/src/net/http/transfer.go:1065
	_go_fuzz_dep_.CoverTab[44013]++
							rres := <-fr.tw.ByteReadCh
							n, err = rres.n, rres.err
							if n == 1 {
//line /usr/local/go/src/net/http/transfer.go:1068
		_go_fuzz_dep_.CoverTab[44018]++
								p[0] = rres.b
//line /usr/local/go/src/net/http/transfer.go:1069
		// _ = "end of CoverTab[44018]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:1070
		_go_fuzz_dep_.CoverTab[44019]++
//line /usr/local/go/src/net/http/transfer.go:1070
		// _ = "end of CoverTab[44019]"
//line /usr/local/go/src/net/http/transfer.go:1070
	}
//line /usr/local/go/src/net/http/transfer.go:1070
	// _ = "end of CoverTab[44013]"
//line /usr/local/go/src/net/http/transfer.go:1070
	_go_fuzz_dep_.CoverTab[44014]++
							if err == nil {
//line /usr/local/go/src/net/http/transfer.go:1071
		_go_fuzz_dep_.CoverTab[44020]++
								err = io.EOF
//line /usr/local/go/src/net/http/transfer.go:1072
		// _ = "end of CoverTab[44020]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:1073
		_go_fuzz_dep_.CoverTab[44021]++
//line /usr/local/go/src/net/http/transfer.go:1073
		// _ = "end of CoverTab[44021]"
//line /usr/local/go/src/net/http/transfer.go:1073
	}
//line /usr/local/go/src/net/http/transfer.go:1073
	// _ = "end of CoverTab[44014]"
//line /usr/local/go/src/net/http/transfer.go:1073
	_go_fuzz_dep_.CoverTab[44015]++
							return
//line /usr/local/go/src/net/http/transfer.go:1074
	// _ = "end of CoverTab[44015]"
}

var nopCloserType = reflect.TypeOf(io.NopCloser(nil))
var nopCloserWriterToType = reflect.TypeOf(io.NopCloser(struct {
	io.Reader
	io.WriterTo
}{}))

// unwrapNopCloser return the underlying reader and true if r is a NopCloser
//line /usr/local/go/src/net/http/transfer.go:1083
// else it return false.
//line /usr/local/go/src/net/http/transfer.go:1085
func unwrapNopCloser(r io.Reader) (underlyingReader io.Reader, isNopCloser bool) {
//line /usr/local/go/src/net/http/transfer.go:1085
	_go_fuzz_dep_.CoverTab[44022]++
							switch reflect.TypeOf(r) {
	case nopCloserType, nopCloserWriterToType:
//line /usr/local/go/src/net/http/transfer.go:1087
		_go_fuzz_dep_.CoverTab[44023]++
								return reflect.ValueOf(r).Field(0).Interface().(io.Reader), true
//line /usr/local/go/src/net/http/transfer.go:1088
		// _ = "end of CoverTab[44023]"
	default:
//line /usr/local/go/src/net/http/transfer.go:1089
		_go_fuzz_dep_.CoverTab[44024]++
								return nil, false
//line /usr/local/go/src/net/http/transfer.go:1090
		// _ = "end of CoverTab[44024]"
	}
//line /usr/local/go/src/net/http/transfer.go:1091
	// _ = "end of CoverTab[44022]"
}

// isKnownInMemoryReader reports whether r is a type known to not
//line /usr/local/go/src/net/http/transfer.go:1094
// block on Read. Its caller uses this as an optional optimization to
//line /usr/local/go/src/net/http/transfer.go:1094
// send fewer TCP packets.
//line /usr/local/go/src/net/http/transfer.go:1097
func isKnownInMemoryReader(r io.Reader) bool {
//line /usr/local/go/src/net/http/transfer.go:1097
	_go_fuzz_dep_.CoverTab[44025]++
							switch r.(type) {
	case *bytes.Reader, *bytes.Buffer, *strings.Reader:
//line /usr/local/go/src/net/http/transfer.go:1099
		_go_fuzz_dep_.CoverTab[44029]++
								return true
//line /usr/local/go/src/net/http/transfer.go:1100
		// _ = "end of CoverTab[44029]"
	}
//line /usr/local/go/src/net/http/transfer.go:1101
	// _ = "end of CoverTab[44025]"
//line /usr/local/go/src/net/http/transfer.go:1101
	_go_fuzz_dep_.CoverTab[44026]++
							if r, ok := unwrapNopCloser(r); ok {
//line /usr/local/go/src/net/http/transfer.go:1102
		_go_fuzz_dep_.CoverTab[44030]++
								return isKnownInMemoryReader(r)
//line /usr/local/go/src/net/http/transfer.go:1103
		// _ = "end of CoverTab[44030]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:1104
		_go_fuzz_dep_.CoverTab[44031]++
//line /usr/local/go/src/net/http/transfer.go:1104
		// _ = "end of CoverTab[44031]"
//line /usr/local/go/src/net/http/transfer.go:1104
	}
//line /usr/local/go/src/net/http/transfer.go:1104
	// _ = "end of CoverTab[44026]"
//line /usr/local/go/src/net/http/transfer.go:1104
	_go_fuzz_dep_.CoverTab[44027]++
							if r, ok := r.(*readTrackingBody); ok {
//line /usr/local/go/src/net/http/transfer.go:1105
		_go_fuzz_dep_.CoverTab[44032]++
								return isKnownInMemoryReader(r.ReadCloser)
//line /usr/local/go/src/net/http/transfer.go:1106
		// _ = "end of CoverTab[44032]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:1107
		_go_fuzz_dep_.CoverTab[44033]++
//line /usr/local/go/src/net/http/transfer.go:1107
		// _ = "end of CoverTab[44033]"
//line /usr/local/go/src/net/http/transfer.go:1107
	}
//line /usr/local/go/src/net/http/transfer.go:1107
	// _ = "end of CoverTab[44027]"
//line /usr/local/go/src/net/http/transfer.go:1107
	_go_fuzz_dep_.CoverTab[44028]++
							return false
//line /usr/local/go/src/net/http/transfer.go:1108
	// _ = "end of CoverTab[44028]"
}

// bufioFlushWriter is an io.Writer wrapper that flushes all writes
//line /usr/local/go/src/net/http/transfer.go:1111
// on its wrapped writer if it's a *bufio.Writer.
//line /usr/local/go/src/net/http/transfer.go:1113
type bufioFlushWriter struct{ w io.Writer }

func (fw bufioFlushWriter) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/transfer.go:1115
	_go_fuzz_dep_.CoverTab[44034]++
							n, err = fw.w.Write(p)
							if bw, ok := fw.w.(*bufio.Writer); n > 0 && func() bool {
//line /usr/local/go/src/net/http/transfer.go:1117
		_go_fuzz_dep_.CoverTab[44036]++
//line /usr/local/go/src/net/http/transfer.go:1117
		return ok
//line /usr/local/go/src/net/http/transfer.go:1117
		// _ = "end of CoverTab[44036]"
//line /usr/local/go/src/net/http/transfer.go:1117
	}() {
//line /usr/local/go/src/net/http/transfer.go:1117
		_go_fuzz_dep_.CoverTab[44037]++
								ferr := bw.Flush()
								if ferr != nil && func() bool {
//line /usr/local/go/src/net/http/transfer.go:1119
			_go_fuzz_dep_.CoverTab[44038]++
//line /usr/local/go/src/net/http/transfer.go:1119
			return err == nil
//line /usr/local/go/src/net/http/transfer.go:1119
			// _ = "end of CoverTab[44038]"
//line /usr/local/go/src/net/http/transfer.go:1119
		}() {
//line /usr/local/go/src/net/http/transfer.go:1119
			_go_fuzz_dep_.CoverTab[44039]++
									err = ferr
//line /usr/local/go/src/net/http/transfer.go:1120
			// _ = "end of CoverTab[44039]"
		} else {
//line /usr/local/go/src/net/http/transfer.go:1121
			_go_fuzz_dep_.CoverTab[44040]++
//line /usr/local/go/src/net/http/transfer.go:1121
			// _ = "end of CoverTab[44040]"
//line /usr/local/go/src/net/http/transfer.go:1121
		}
//line /usr/local/go/src/net/http/transfer.go:1121
		// _ = "end of CoverTab[44037]"
	} else {
//line /usr/local/go/src/net/http/transfer.go:1122
		_go_fuzz_dep_.CoverTab[44041]++
//line /usr/local/go/src/net/http/transfer.go:1122
		// _ = "end of CoverTab[44041]"
//line /usr/local/go/src/net/http/transfer.go:1122
	}
//line /usr/local/go/src/net/http/transfer.go:1122
	// _ = "end of CoverTab[44034]"
//line /usr/local/go/src/net/http/transfer.go:1122
	_go_fuzz_dep_.CoverTab[44035]++
							return
//line /usr/local/go/src/net/http/transfer.go:1123
	// _ = "end of CoverTab[44035]"
}

//line /usr/local/go/src/net/http/transfer.go:1124
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/transfer.go:1124
var _ = _go_fuzz_dep_.CoverTab
