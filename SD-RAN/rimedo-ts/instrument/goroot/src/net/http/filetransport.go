// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/http/filetransport.go:5
package http

//line /usr/local/go/src/net/http/filetransport.go:5
import (
//line /usr/local/go/src/net/http/filetransport.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/filetransport.go:5
)
//line /usr/local/go/src/net/http/filetransport.go:5
import (
//line /usr/local/go/src/net/http/filetransport.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/filetransport.go:5
)

import (
	"fmt"
	"io"
)

// fileTransport implements RoundTripper for the 'file' protocol.
type fileTransport struct {
	fh fileHandler
}

// NewFileTransport returns a new RoundTripper, serving the provided
//line /usr/local/go/src/net/http/filetransport.go:17
// FileSystem. The returned RoundTripper ignores the URL host in its
//line /usr/local/go/src/net/http/filetransport.go:17
// incoming requests, as well as most other properties of the
//line /usr/local/go/src/net/http/filetransport.go:17
// request.
//line /usr/local/go/src/net/http/filetransport.go:17
//
//line /usr/local/go/src/net/http/filetransport.go:17
// The typical use case for NewFileTransport is to register the "file"
//line /usr/local/go/src/net/http/filetransport.go:17
// protocol with a Transport, as in:
//line /usr/local/go/src/net/http/filetransport.go:17
//
//line /usr/local/go/src/net/http/filetransport.go:17
//	t := &http.Transport{}
//line /usr/local/go/src/net/http/filetransport.go:17
//	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
//line /usr/local/go/src/net/http/filetransport.go:17
//	c := &http.Client{Transport: t}
//line /usr/local/go/src/net/http/filetransport.go:17
//	res, err := c.Get("file:///etc/passwd")
//line /usr/local/go/src/net/http/filetransport.go:17
//	...
//line /usr/local/go/src/net/http/filetransport.go:30
func NewFileTransport(fs FileSystem) RoundTripper {
//line /usr/local/go/src/net/http/filetransport.go:30
	_go_fuzz_dep_.CoverTab[37148]++
							return fileTransport{fileHandler{fs}}
//line /usr/local/go/src/net/http/filetransport.go:31
	// _ = "end of CoverTab[37148]"
}

func (t fileTransport) RoundTrip(req *Request) (resp *Response, err error) {
//line /usr/local/go/src/net/http/filetransport.go:34
	_go_fuzz_dep_.CoverTab[37149]++

//line /usr/local/go/src/net/http/filetransport.go:41
	rw, resc := newPopulateResponseWriter()
//line /usr/local/go/src/net/http/filetransport.go:41
	_curRoutineNum14_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/filetransport.go:41
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum14_)
							go func() {
//line /usr/local/go/src/net/http/filetransport.go:42
		_go_fuzz_dep_.CoverTab[37151]++
//line /usr/local/go/src/net/http/filetransport.go:42
		defer func() {
//line /usr/local/go/src/net/http/filetransport.go:42
			_go_fuzz_dep_.CoverTab[37152]++
//line /usr/local/go/src/net/http/filetransport.go:42
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum14_)
//line /usr/local/go/src/net/http/filetransport.go:42
			// _ = "end of CoverTab[37152]"
//line /usr/local/go/src/net/http/filetransport.go:42
		}()
								t.fh.ServeHTTP(rw, req)
								rw.finish()
//line /usr/local/go/src/net/http/filetransport.go:44
		// _ = "end of CoverTab[37151]"
	}()
//line /usr/local/go/src/net/http/filetransport.go:45
	// _ = "end of CoverTab[37149]"
//line /usr/local/go/src/net/http/filetransport.go:45
	_go_fuzz_dep_.CoverTab[37150]++
							return <-resc, nil
//line /usr/local/go/src/net/http/filetransport.go:46
	// _ = "end of CoverTab[37150]"
}

func newPopulateResponseWriter() (*populateResponse, <-chan *Response) {
//line /usr/local/go/src/net/http/filetransport.go:49
	_go_fuzz_dep_.CoverTab[37153]++
							pr, pw := io.Pipe()
							rw := &populateResponse{
		ch:	make(chan *Response),
		pw:	pw,
		res: &Response{
			Proto:		"HTTP/1.0",
			ProtoMajor:	1,
			Header:		make(Header),
			Close:		true,
			Body:		pr,
		},
	}
							return rw, rw.ch
//line /usr/local/go/src/net/http/filetransport.go:62
	// _ = "end of CoverTab[37153]"
}

// populateResponse is a ResponseWriter that populates the *Response
//line /usr/local/go/src/net/http/filetransport.go:65
// in res, and writes its body to a pipe connected to the response
//line /usr/local/go/src/net/http/filetransport.go:65
// body. Once writes begin or finish() is called, the response is sent
//line /usr/local/go/src/net/http/filetransport.go:65
// on ch.
//line /usr/local/go/src/net/http/filetransport.go:69
type populateResponse struct {
	res		*Response
	ch		chan *Response
	wroteHeader	bool
	hasContent	bool
	sentResponse	bool
	pw		*io.PipeWriter
}

func (pr *populateResponse) finish() {
//line /usr/local/go/src/net/http/filetransport.go:78
	_go_fuzz_dep_.CoverTab[37154]++
							if !pr.wroteHeader {
//line /usr/local/go/src/net/http/filetransport.go:79
		_go_fuzz_dep_.CoverTab[37157]++
								pr.WriteHeader(500)
//line /usr/local/go/src/net/http/filetransport.go:80
		// _ = "end of CoverTab[37157]"
	} else {
//line /usr/local/go/src/net/http/filetransport.go:81
		_go_fuzz_dep_.CoverTab[37158]++
//line /usr/local/go/src/net/http/filetransport.go:81
		// _ = "end of CoverTab[37158]"
//line /usr/local/go/src/net/http/filetransport.go:81
	}
//line /usr/local/go/src/net/http/filetransport.go:81
	// _ = "end of CoverTab[37154]"
//line /usr/local/go/src/net/http/filetransport.go:81
	_go_fuzz_dep_.CoverTab[37155]++
							if !pr.sentResponse {
//line /usr/local/go/src/net/http/filetransport.go:82
		_go_fuzz_dep_.CoverTab[37159]++
								pr.sendResponse()
//line /usr/local/go/src/net/http/filetransport.go:83
		// _ = "end of CoverTab[37159]"
	} else {
//line /usr/local/go/src/net/http/filetransport.go:84
		_go_fuzz_dep_.CoverTab[37160]++
//line /usr/local/go/src/net/http/filetransport.go:84
		// _ = "end of CoverTab[37160]"
//line /usr/local/go/src/net/http/filetransport.go:84
	}
//line /usr/local/go/src/net/http/filetransport.go:84
	// _ = "end of CoverTab[37155]"
//line /usr/local/go/src/net/http/filetransport.go:84
	_go_fuzz_dep_.CoverTab[37156]++
							pr.pw.Close()
//line /usr/local/go/src/net/http/filetransport.go:85
	// _ = "end of CoverTab[37156]"
}

func (pr *populateResponse) sendResponse() {
//line /usr/local/go/src/net/http/filetransport.go:88
	_go_fuzz_dep_.CoverTab[37161]++
							if pr.sentResponse {
//line /usr/local/go/src/net/http/filetransport.go:89
		_go_fuzz_dep_.CoverTab[37164]++
								return
//line /usr/local/go/src/net/http/filetransport.go:90
		// _ = "end of CoverTab[37164]"
	} else {
//line /usr/local/go/src/net/http/filetransport.go:91
		_go_fuzz_dep_.CoverTab[37165]++
//line /usr/local/go/src/net/http/filetransport.go:91
		// _ = "end of CoverTab[37165]"
//line /usr/local/go/src/net/http/filetransport.go:91
	}
//line /usr/local/go/src/net/http/filetransport.go:91
	// _ = "end of CoverTab[37161]"
//line /usr/local/go/src/net/http/filetransport.go:91
	_go_fuzz_dep_.CoverTab[37162]++
							pr.sentResponse = true

							if pr.hasContent {
//line /usr/local/go/src/net/http/filetransport.go:94
		_go_fuzz_dep_.CoverTab[37166]++
								pr.res.ContentLength = -1
//line /usr/local/go/src/net/http/filetransport.go:95
		// _ = "end of CoverTab[37166]"
	} else {
//line /usr/local/go/src/net/http/filetransport.go:96
		_go_fuzz_dep_.CoverTab[37167]++
//line /usr/local/go/src/net/http/filetransport.go:96
		// _ = "end of CoverTab[37167]"
//line /usr/local/go/src/net/http/filetransport.go:96
	}
//line /usr/local/go/src/net/http/filetransport.go:96
	// _ = "end of CoverTab[37162]"
//line /usr/local/go/src/net/http/filetransport.go:96
	_go_fuzz_dep_.CoverTab[37163]++
							pr.ch <- pr.res
//line /usr/local/go/src/net/http/filetransport.go:97
	// _ = "end of CoverTab[37163]"
}

func (pr *populateResponse) Header() Header {
//line /usr/local/go/src/net/http/filetransport.go:100
	_go_fuzz_dep_.CoverTab[37168]++
							return pr.res.Header
//line /usr/local/go/src/net/http/filetransport.go:101
	// _ = "end of CoverTab[37168]"
}

func (pr *populateResponse) WriteHeader(code int) {
//line /usr/local/go/src/net/http/filetransport.go:104
	_go_fuzz_dep_.CoverTab[37169]++
							if pr.wroteHeader {
//line /usr/local/go/src/net/http/filetransport.go:105
		_go_fuzz_dep_.CoverTab[37171]++
								return
//line /usr/local/go/src/net/http/filetransport.go:106
		// _ = "end of CoverTab[37171]"
	} else {
//line /usr/local/go/src/net/http/filetransport.go:107
		_go_fuzz_dep_.CoverTab[37172]++
//line /usr/local/go/src/net/http/filetransport.go:107
		// _ = "end of CoverTab[37172]"
//line /usr/local/go/src/net/http/filetransport.go:107
	}
//line /usr/local/go/src/net/http/filetransport.go:107
	// _ = "end of CoverTab[37169]"
//line /usr/local/go/src/net/http/filetransport.go:107
	_go_fuzz_dep_.CoverTab[37170]++
							pr.wroteHeader = true

							pr.res.StatusCode = code
							pr.res.Status = fmt.Sprintf("%d %s", code, StatusText(code))
//line /usr/local/go/src/net/http/filetransport.go:111
	// _ = "end of CoverTab[37170]"
}

func (pr *populateResponse) Write(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/filetransport.go:114
	_go_fuzz_dep_.CoverTab[37173]++
							if !pr.wroteHeader {
//line /usr/local/go/src/net/http/filetransport.go:115
		_go_fuzz_dep_.CoverTab[37176]++
								pr.WriteHeader(StatusOK)
//line /usr/local/go/src/net/http/filetransport.go:116
		// _ = "end of CoverTab[37176]"
	} else {
//line /usr/local/go/src/net/http/filetransport.go:117
		_go_fuzz_dep_.CoverTab[37177]++
//line /usr/local/go/src/net/http/filetransport.go:117
		// _ = "end of CoverTab[37177]"
//line /usr/local/go/src/net/http/filetransport.go:117
	}
//line /usr/local/go/src/net/http/filetransport.go:117
	// _ = "end of CoverTab[37173]"
//line /usr/local/go/src/net/http/filetransport.go:117
	_go_fuzz_dep_.CoverTab[37174]++
							pr.hasContent = true
							if !pr.sentResponse {
//line /usr/local/go/src/net/http/filetransport.go:119
		_go_fuzz_dep_.CoverTab[37178]++
								pr.sendResponse()
//line /usr/local/go/src/net/http/filetransport.go:120
		// _ = "end of CoverTab[37178]"
	} else {
//line /usr/local/go/src/net/http/filetransport.go:121
		_go_fuzz_dep_.CoverTab[37179]++
//line /usr/local/go/src/net/http/filetransport.go:121
		// _ = "end of CoverTab[37179]"
//line /usr/local/go/src/net/http/filetransport.go:121
	}
//line /usr/local/go/src/net/http/filetransport.go:121
	// _ = "end of CoverTab[37174]"
//line /usr/local/go/src/net/http/filetransport.go:121
	_go_fuzz_dep_.CoverTab[37175]++
							return pr.pw.Write(p)
//line /usr/local/go/src/net/http/filetransport.go:122
	// _ = "end of CoverTab[37175]"
}

//line /usr/local/go/src/net/http/filetransport.go:123
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/filetransport.go:123
var _ = _go_fuzz_dep_.CoverTab
