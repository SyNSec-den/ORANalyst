// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/http/httputil/dump.go:5
package httputil

//line /usr/local/go/src/net/http/httputil/dump.go:5
import (
//line /usr/local/go/src/net/http/httputil/dump.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/http/httputil/dump.go:5
)
//line /usr/local/go/src/net/http/httputil/dump.go:5
import (
//line /usr/local/go/src/net/http/httputil/dump.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/http/httputil/dump.go:5
)

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// drainBody reads all of b to memory and then returns two equivalent
//line /usr/local/go/src/net/http/httputil/dump.go:20
// ReadClosers yielding the same bytes.
//line /usr/local/go/src/net/http/httputil/dump.go:20
//
//line /usr/local/go/src/net/http/httputil/dump.go:20
// It returns an error if the initial slurp of all bytes fails. It does not attempt
//line /usr/local/go/src/net/http/httputil/dump.go:20
// to make the returned ReadClosers have identical error-matching behavior.
//line /usr/local/go/src/net/http/httputil/dump.go:25
func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
//line /usr/local/go/src/net/http/httputil/dump.go:25
	_go_fuzz_dep_.CoverTab[76048]++
							if b == nil || func() bool {
//line /usr/local/go/src/net/http/httputil/dump.go:26
		_go_fuzz_dep_.CoverTab[76052]++
//line /usr/local/go/src/net/http/httputil/dump.go:26
		return b == http.NoBody
//line /usr/local/go/src/net/http/httputil/dump.go:26
		// _ = "end of CoverTab[76052]"
//line /usr/local/go/src/net/http/httputil/dump.go:26
	}() {
//line /usr/local/go/src/net/http/httputil/dump.go:26
		_go_fuzz_dep_.CoverTab[76053]++

								return http.NoBody, http.NoBody, nil
//line /usr/local/go/src/net/http/httputil/dump.go:28
		// _ = "end of CoverTab[76053]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:29
		_go_fuzz_dep_.CoverTab[76054]++
//line /usr/local/go/src/net/http/httputil/dump.go:29
		// _ = "end of CoverTab[76054]"
//line /usr/local/go/src/net/http/httputil/dump.go:29
	}
//line /usr/local/go/src/net/http/httputil/dump.go:29
	// _ = "end of CoverTab[76048]"
//line /usr/local/go/src/net/http/httputil/dump.go:29
	_go_fuzz_dep_.CoverTab[76049]++
							var buf bytes.Buffer
							if _, err = buf.ReadFrom(b); err != nil {
//line /usr/local/go/src/net/http/httputil/dump.go:31
		_go_fuzz_dep_.CoverTab[76055]++
								return nil, b, err
//line /usr/local/go/src/net/http/httputil/dump.go:32
		// _ = "end of CoverTab[76055]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:33
		_go_fuzz_dep_.CoverTab[76056]++
//line /usr/local/go/src/net/http/httputil/dump.go:33
		// _ = "end of CoverTab[76056]"
//line /usr/local/go/src/net/http/httputil/dump.go:33
	}
//line /usr/local/go/src/net/http/httputil/dump.go:33
	// _ = "end of CoverTab[76049]"
//line /usr/local/go/src/net/http/httputil/dump.go:33
	_go_fuzz_dep_.CoverTab[76050]++
							if err = b.Close(); err != nil {
//line /usr/local/go/src/net/http/httputil/dump.go:34
		_go_fuzz_dep_.CoverTab[76057]++
								return nil, b, err
//line /usr/local/go/src/net/http/httputil/dump.go:35
		// _ = "end of CoverTab[76057]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:36
		_go_fuzz_dep_.CoverTab[76058]++
//line /usr/local/go/src/net/http/httputil/dump.go:36
		// _ = "end of CoverTab[76058]"
//line /usr/local/go/src/net/http/httputil/dump.go:36
	}
//line /usr/local/go/src/net/http/httputil/dump.go:36
	// _ = "end of CoverTab[76050]"
//line /usr/local/go/src/net/http/httputil/dump.go:36
	_go_fuzz_dep_.CoverTab[76051]++
							return io.NopCloser(&buf), io.NopCloser(bytes.NewReader(buf.Bytes())), nil
//line /usr/local/go/src/net/http/httputil/dump.go:37
	// _ = "end of CoverTab[76051]"
}

// dumpConn is a net.Conn which writes to Writer and reads from Reader
type dumpConn struct {
	io.Writer
	io.Reader
}

func (c *dumpConn) Close() error {
//line /usr/local/go/src/net/http/httputil/dump.go:46
	_go_fuzz_dep_.CoverTab[76059]++
//line /usr/local/go/src/net/http/httputil/dump.go:46
	return nil
//line /usr/local/go/src/net/http/httputil/dump.go:46
	// _ = "end of CoverTab[76059]"
//line /usr/local/go/src/net/http/httputil/dump.go:46
}
func (c *dumpConn) LocalAddr() net.Addr {
//line /usr/local/go/src/net/http/httputil/dump.go:47
	_go_fuzz_dep_.CoverTab[76060]++
//line /usr/local/go/src/net/http/httputil/dump.go:47
	return nil
//line /usr/local/go/src/net/http/httputil/dump.go:47
	// _ = "end of CoverTab[76060]"
//line /usr/local/go/src/net/http/httputil/dump.go:47
}
func (c *dumpConn) RemoteAddr() net.Addr {
//line /usr/local/go/src/net/http/httputil/dump.go:48
	_go_fuzz_dep_.CoverTab[76061]++
//line /usr/local/go/src/net/http/httputil/dump.go:48
	return nil
//line /usr/local/go/src/net/http/httputil/dump.go:48
	// _ = "end of CoverTab[76061]"
//line /usr/local/go/src/net/http/httputil/dump.go:48
}
func (c *dumpConn) SetDeadline(t time.Time) error {
//line /usr/local/go/src/net/http/httputil/dump.go:49
	_go_fuzz_dep_.CoverTab[76062]++
//line /usr/local/go/src/net/http/httputil/dump.go:49
	return nil
//line /usr/local/go/src/net/http/httputil/dump.go:49
	// _ = "end of CoverTab[76062]"
//line /usr/local/go/src/net/http/httputil/dump.go:49
}
func (c *dumpConn) SetReadDeadline(t time.Time) error {
//line /usr/local/go/src/net/http/httputil/dump.go:50
	_go_fuzz_dep_.CoverTab[76063]++
//line /usr/local/go/src/net/http/httputil/dump.go:50
	return nil
//line /usr/local/go/src/net/http/httputil/dump.go:50
	// _ = "end of CoverTab[76063]"
//line /usr/local/go/src/net/http/httputil/dump.go:50
}
func (c *dumpConn) SetWriteDeadline(t time.Time) error {
//line /usr/local/go/src/net/http/httputil/dump.go:51
	_go_fuzz_dep_.CoverTab[76064]++
//line /usr/local/go/src/net/http/httputil/dump.go:51
	return nil
//line /usr/local/go/src/net/http/httputil/dump.go:51
	// _ = "end of CoverTab[76064]"
//line /usr/local/go/src/net/http/httputil/dump.go:51
}

type neverEnding byte

func (b neverEnding) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/http/httputil/dump.go:55
	_go_fuzz_dep_.CoverTab[76065]++
							for i := range p {
//line /usr/local/go/src/net/http/httputil/dump.go:56
		_go_fuzz_dep_.CoverTab[76067]++
								p[i] = byte(b)
//line /usr/local/go/src/net/http/httputil/dump.go:57
		// _ = "end of CoverTab[76067]"
	}
//line /usr/local/go/src/net/http/httputil/dump.go:58
	// _ = "end of CoverTab[76065]"
//line /usr/local/go/src/net/http/httputil/dump.go:58
	_go_fuzz_dep_.CoverTab[76066]++
							return len(p), nil
//line /usr/local/go/src/net/http/httputil/dump.go:59
	// _ = "end of CoverTab[76066]"
}

// outgoingLength is a copy of the unexported
//line /usr/local/go/src/net/http/httputil/dump.go:62
// (*http.Request).outgoingLength method.
//line /usr/local/go/src/net/http/httputil/dump.go:64
func outgoingLength(req *http.Request) int64 {
//line /usr/local/go/src/net/http/httputil/dump.go:64
	_go_fuzz_dep_.CoverTab[76068]++
							if req.Body == nil || func() bool {
//line /usr/local/go/src/net/http/httputil/dump.go:65
		_go_fuzz_dep_.CoverTab[76071]++
//line /usr/local/go/src/net/http/httputil/dump.go:65
		return req.Body == http.NoBody
//line /usr/local/go/src/net/http/httputil/dump.go:65
		// _ = "end of CoverTab[76071]"
//line /usr/local/go/src/net/http/httputil/dump.go:65
	}() {
//line /usr/local/go/src/net/http/httputil/dump.go:65
		_go_fuzz_dep_.CoverTab[76072]++
								return 0
//line /usr/local/go/src/net/http/httputil/dump.go:66
		// _ = "end of CoverTab[76072]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:67
		_go_fuzz_dep_.CoverTab[76073]++
//line /usr/local/go/src/net/http/httputil/dump.go:67
		// _ = "end of CoverTab[76073]"
//line /usr/local/go/src/net/http/httputil/dump.go:67
	}
//line /usr/local/go/src/net/http/httputil/dump.go:67
	// _ = "end of CoverTab[76068]"
//line /usr/local/go/src/net/http/httputil/dump.go:67
	_go_fuzz_dep_.CoverTab[76069]++
							if req.ContentLength != 0 {
//line /usr/local/go/src/net/http/httputil/dump.go:68
		_go_fuzz_dep_.CoverTab[76074]++
								return req.ContentLength
//line /usr/local/go/src/net/http/httputil/dump.go:69
		// _ = "end of CoverTab[76074]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:70
		_go_fuzz_dep_.CoverTab[76075]++
//line /usr/local/go/src/net/http/httputil/dump.go:70
		// _ = "end of CoverTab[76075]"
//line /usr/local/go/src/net/http/httputil/dump.go:70
	}
//line /usr/local/go/src/net/http/httputil/dump.go:70
	// _ = "end of CoverTab[76069]"
//line /usr/local/go/src/net/http/httputil/dump.go:70
	_go_fuzz_dep_.CoverTab[76070]++
							return -1
//line /usr/local/go/src/net/http/httputil/dump.go:71
	// _ = "end of CoverTab[76070]"
}

// DumpRequestOut is like DumpRequest but for outgoing client requests. It
//line /usr/local/go/src/net/http/httputil/dump.go:74
// includes any headers that the standard http.Transport adds, such as
//line /usr/local/go/src/net/http/httputil/dump.go:74
// User-Agent.
//line /usr/local/go/src/net/http/httputil/dump.go:77
func DumpRequestOut(req *http.Request, body bool) ([]byte, error) {
//line /usr/local/go/src/net/http/httputil/dump.go:77
	_go_fuzz_dep_.CoverTab[76076]++
							save := req.Body
							dummyBody := false
							if !body {
//line /usr/local/go/src/net/http/httputil/dump.go:80
		_go_fuzz_dep_.CoverTab[76083]++
								contentLength := outgoingLength(req)
								if contentLength != 0 {
//line /usr/local/go/src/net/http/httputil/dump.go:82
			_go_fuzz_dep_.CoverTab[76084]++
									req.Body = io.NopCloser(io.LimitReader(neverEnding('x'), contentLength))
									dummyBody = true
//line /usr/local/go/src/net/http/httputil/dump.go:84
			// _ = "end of CoverTab[76084]"
		} else {
//line /usr/local/go/src/net/http/httputil/dump.go:85
			_go_fuzz_dep_.CoverTab[76085]++
//line /usr/local/go/src/net/http/httputil/dump.go:85
			// _ = "end of CoverTab[76085]"
//line /usr/local/go/src/net/http/httputil/dump.go:85
		}
//line /usr/local/go/src/net/http/httputil/dump.go:85
		// _ = "end of CoverTab[76083]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:86
		_go_fuzz_dep_.CoverTab[76086]++
								var err error
								save, req.Body, err = drainBody(req.Body)
								if err != nil {
//line /usr/local/go/src/net/http/httputil/dump.go:89
			_go_fuzz_dep_.CoverTab[76087]++
									return nil, err
//line /usr/local/go/src/net/http/httputil/dump.go:90
			// _ = "end of CoverTab[76087]"
		} else {
//line /usr/local/go/src/net/http/httputil/dump.go:91
			_go_fuzz_dep_.CoverTab[76088]++
//line /usr/local/go/src/net/http/httputil/dump.go:91
			// _ = "end of CoverTab[76088]"
//line /usr/local/go/src/net/http/httputil/dump.go:91
		}
//line /usr/local/go/src/net/http/httputil/dump.go:91
		// _ = "end of CoverTab[76086]"
	}
//line /usr/local/go/src/net/http/httputil/dump.go:92
	// _ = "end of CoverTab[76076]"
//line /usr/local/go/src/net/http/httputil/dump.go:92
	_go_fuzz_dep_.CoverTab[76077]++

//line /usr/local/go/src/net/http/httputil/dump.go:98
	reqSend := req
	if req.URL.Scheme == "https" {
//line /usr/local/go/src/net/http/httputil/dump.go:99
		_go_fuzz_dep_.CoverTab[76089]++
								reqSend = new(http.Request)
								*reqSend = *req
								reqSend.URL = new(url.URL)
								*reqSend.URL = *req.URL
								reqSend.URL.Scheme = "http"
//line /usr/local/go/src/net/http/httputil/dump.go:104
		// _ = "end of CoverTab[76089]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:105
		_go_fuzz_dep_.CoverTab[76090]++
//line /usr/local/go/src/net/http/httputil/dump.go:105
		// _ = "end of CoverTab[76090]"
//line /usr/local/go/src/net/http/httputil/dump.go:105
	}
//line /usr/local/go/src/net/http/httputil/dump.go:105
	// _ = "end of CoverTab[76077]"
//line /usr/local/go/src/net/http/httputil/dump.go:105
	_go_fuzz_dep_.CoverTab[76078]++

	// Use the actual Transport code to record what we would send
	// on the wire, but not using TCP.  Use a Transport with a
	// custom dialer that returns a fake net.Conn that waits
	// for the full input (and recording it), and then responds
	// with a dummy response.
	var buf bytes.Buffer	// records the output
	pr, pw := io.Pipe()
	defer pr.Close()
	defer pw.Close()
	dr := &delegateReader{c: make(chan io.Reader)}

	t := &http.Transport{
		Dial: func(net, addr string) (net.Conn, error) {
//line /usr/local/go/src/net/http/httputil/dump.go:119
			_go_fuzz_dep_.CoverTab[76091]++
									return &dumpConn{io.MultiWriter(&buf, pw), dr}, nil
//line /usr/local/go/src/net/http/httputil/dump.go:120
			// _ = "end of CoverTab[76091]"
		},
	}
//line /usr/local/go/src/net/http/httputil/dump.go:122
	// _ = "end of CoverTab[76078]"
//line /usr/local/go/src/net/http/httputil/dump.go:122
	_go_fuzz_dep_.CoverTab[76079]++
							defer t.CloseIdleConnections()

//line /usr/local/go/src/net/http/httputil/dump.go:128
	quitReadCh := make(chan struct{})
//line /usr/local/go/src/net/http/httputil/dump.go:128
	_curRoutineNum71_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/http/httputil/dump.go:128
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum71_)

							go func() {
//line /usr/local/go/src/net/http/httputil/dump.go:130
		_go_fuzz_dep_.CoverTab[76092]++
//line /usr/local/go/src/net/http/httputil/dump.go:130
		defer func() {
//line /usr/local/go/src/net/http/httputil/dump.go:130
			_go_fuzz_dep_.CoverTab[76094]++
//line /usr/local/go/src/net/http/httputil/dump.go:130
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum71_)
//line /usr/local/go/src/net/http/httputil/dump.go:130
			// _ = "end of CoverTab[76094]"
//line /usr/local/go/src/net/http/httputil/dump.go:130
		}()
								req, err := http.ReadRequest(bufio.NewReader(pr))
								if err == nil {
//line /usr/local/go/src/net/http/httputil/dump.go:132
			_go_fuzz_dep_.CoverTab[76095]++

//line /usr/local/go/src/net/http/httputil/dump.go:135
			io.Copy(io.Discard, req.Body)
									req.Body.Close()
//line /usr/local/go/src/net/http/httputil/dump.go:136
			// _ = "end of CoverTab[76095]"
		} else {
//line /usr/local/go/src/net/http/httputil/dump.go:137
			_go_fuzz_dep_.CoverTab[76096]++
//line /usr/local/go/src/net/http/httputil/dump.go:137
			// _ = "end of CoverTab[76096]"
//line /usr/local/go/src/net/http/httputil/dump.go:137
		}
//line /usr/local/go/src/net/http/httputil/dump.go:137
		// _ = "end of CoverTab[76092]"
//line /usr/local/go/src/net/http/httputil/dump.go:137
		_go_fuzz_dep_.CoverTab[76093]++
								select {
		case dr.c <- strings.NewReader("HTTP/1.1 204 No Content\r\nConnection: close\r\n\r\n"):
//line /usr/local/go/src/net/http/httputil/dump.go:139
			_go_fuzz_dep_.CoverTab[76097]++
//line /usr/local/go/src/net/http/httputil/dump.go:139
			// _ = "end of CoverTab[76097]"
		case <-quitReadCh:
//line /usr/local/go/src/net/http/httputil/dump.go:140
			_go_fuzz_dep_.CoverTab[76098]++

									close(dr.c)
//line /usr/local/go/src/net/http/httputil/dump.go:142
			// _ = "end of CoverTab[76098]"
		}
//line /usr/local/go/src/net/http/httputil/dump.go:143
		// _ = "end of CoverTab[76093]"
	}()
//line /usr/local/go/src/net/http/httputil/dump.go:144
	// _ = "end of CoverTab[76079]"
//line /usr/local/go/src/net/http/httputil/dump.go:144
	_go_fuzz_dep_.CoverTab[76080]++

							_, err := t.RoundTrip(reqSend)

							req.Body = save
							if err != nil {
//line /usr/local/go/src/net/http/httputil/dump.go:149
		_go_fuzz_dep_.CoverTab[76099]++
								pw.Close()
								dr.err = err
								close(quitReadCh)
								return nil, err
//line /usr/local/go/src/net/http/httputil/dump.go:153
		// _ = "end of CoverTab[76099]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:154
		_go_fuzz_dep_.CoverTab[76100]++
//line /usr/local/go/src/net/http/httputil/dump.go:154
		// _ = "end of CoverTab[76100]"
//line /usr/local/go/src/net/http/httputil/dump.go:154
	}
//line /usr/local/go/src/net/http/httputil/dump.go:154
	// _ = "end of CoverTab[76080]"
//line /usr/local/go/src/net/http/httputil/dump.go:154
	_go_fuzz_dep_.CoverTab[76081]++
							dump := buf.Bytes()

//line /usr/local/go/src/net/http/httputil/dump.go:162
	if dummyBody {
//line /usr/local/go/src/net/http/httputil/dump.go:162
		_go_fuzz_dep_.CoverTab[76101]++
								if i := bytes.Index(dump, []byte("\r\n\r\n")); i >= 0 {
//line /usr/local/go/src/net/http/httputil/dump.go:163
			_go_fuzz_dep_.CoverTab[76102]++
									dump = dump[:i+4]
//line /usr/local/go/src/net/http/httputil/dump.go:164
			// _ = "end of CoverTab[76102]"
		} else {
//line /usr/local/go/src/net/http/httputil/dump.go:165
			_go_fuzz_dep_.CoverTab[76103]++
//line /usr/local/go/src/net/http/httputil/dump.go:165
			// _ = "end of CoverTab[76103]"
//line /usr/local/go/src/net/http/httputil/dump.go:165
		}
//line /usr/local/go/src/net/http/httputil/dump.go:165
		// _ = "end of CoverTab[76101]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:166
		_go_fuzz_dep_.CoverTab[76104]++
//line /usr/local/go/src/net/http/httputil/dump.go:166
		// _ = "end of CoverTab[76104]"
//line /usr/local/go/src/net/http/httputil/dump.go:166
	}
//line /usr/local/go/src/net/http/httputil/dump.go:166
	// _ = "end of CoverTab[76081]"
//line /usr/local/go/src/net/http/httputil/dump.go:166
	_go_fuzz_dep_.CoverTab[76082]++
							return dump, nil
//line /usr/local/go/src/net/http/httputil/dump.go:167
	// _ = "end of CoverTab[76082]"
}

// delegateReader is a reader that delegates to another reader,
//line /usr/local/go/src/net/http/httputil/dump.go:170
// once it arrives on a channel.
//line /usr/local/go/src/net/http/httputil/dump.go:172
type delegateReader struct {
	c	chan io.Reader
	err	error		// only used if r is nil and c is closed.
	r	io.Reader	// nil until received from c
}

func (r *delegateReader) Read(p []byte) (int, error) {
//line /usr/local/go/src/net/http/httputil/dump.go:178
	_go_fuzz_dep_.CoverTab[76105]++
							if r.r == nil {
//line /usr/local/go/src/net/http/httputil/dump.go:179
		_go_fuzz_dep_.CoverTab[76107]++
								var ok bool
								if r.r, ok = <-r.c; !ok {
//line /usr/local/go/src/net/http/httputil/dump.go:181
			_go_fuzz_dep_.CoverTab[76108]++
									return 0, r.err
//line /usr/local/go/src/net/http/httputil/dump.go:182
			// _ = "end of CoverTab[76108]"
		} else {
//line /usr/local/go/src/net/http/httputil/dump.go:183
			_go_fuzz_dep_.CoverTab[76109]++
//line /usr/local/go/src/net/http/httputil/dump.go:183
			// _ = "end of CoverTab[76109]"
//line /usr/local/go/src/net/http/httputil/dump.go:183
		}
//line /usr/local/go/src/net/http/httputil/dump.go:183
		// _ = "end of CoverTab[76107]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:184
		_go_fuzz_dep_.CoverTab[76110]++
//line /usr/local/go/src/net/http/httputil/dump.go:184
		// _ = "end of CoverTab[76110]"
//line /usr/local/go/src/net/http/httputil/dump.go:184
	}
//line /usr/local/go/src/net/http/httputil/dump.go:184
	// _ = "end of CoverTab[76105]"
//line /usr/local/go/src/net/http/httputil/dump.go:184
	_go_fuzz_dep_.CoverTab[76106]++
							return r.r.Read(p)
//line /usr/local/go/src/net/http/httputil/dump.go:185
	// _ = "end of CoverTab[76106]"
}

// Return value if nonempty, def otherwise.
func valueOrDefault(value, def string) string {
//line /usr/local/go/src/net/http/httputil/dump.go:189
	_go_fuzz_dep_.CoverTab[76111]++
							if value != "" {
//line /usr/local/go/src/net/http/httputil/dump.go:190
		_go_fuzz_dep_.CoverTab[76113]++
								return value
//line /usr/local/go/src/net/http/httputil/dump.go:191
		// _ = "end of CoverTab[76113]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:192
		_go_fuzz_dep_.CoverTab[76114]++
//line /usr/local/go/src/net/http/httputil/dump.go:192
		// _ = "end of CoverTab[76114]"
//line /usr/local/go/src/net/http/httputil/dump.go:192
	}
//line /usr/local/go/src/net/http/httputil/dump.go:192
	// _ = "end of CoverTab[76111]"
//line /usr/local/go/src/net/http/httputil/dump.go:192
	_go_fuzz_dep_.CoverTab[76112]++
							return def
//line /usr/local/go/src/net/http/httputil/dump.go:193
	// _ = "end of CoverTab[76112]"
}

var reqWriteExcludeHeaderDump = map[string]bool{
	"Host":			true,
	"Transfer-Encoding":	true,
	"Trailer":		true,
}

// DumpRequest returns the given request in its HTTP/1.x wire
//line /usr/local/go/src/net/http/httputil/dump.go:202
// representation. It should only be used by servers to debug client
//line /usr/local/go/src/net/http/httputil/dump.go:202
// requests. The returned representation is an approximation only;
//line /usr/local/go/src/net/http/httputil/dump.go:202
// some details of the initial request are lost while parsing it into
//line /usr/local/go/src/net/http/httputil/dump.go:202
// an http.Request. In particular, the order and case of header field
//line /usr/local/go/src/net/http/httputil/dump.go:202
// names are lost. The order of values in multi-valued headers is kept
//line /usr/local/go/src/net/http/httputil/dump.go:202
// intact. HTTP/2 requests are dumped in HTTP/1.x form, not in their
//line /usr/local/go/src/net/http/httputil/dump.go:202
// original binary representations.
//line /usr/local/go/src/net/http/httputil/dump.go:202
//
//line /usr/local/go/src/net/http/httputil/dump.go:202
// If body is true, DumpRequest also returns the body. To do so, it
//line /usr/local/go/src/net/http/httputil/dump.go:202
// consumes req.Body and then replaces it with a new io.ReadCloser
//line /usr/local/go/src/net/http/httputil/dump.go:202
// that yields the same bytes. If DumpRequest returns an error,
//line /usr/local/go/src/net/http/httputil/dump.go:202
// the state of req is undefined.
//line /usr/local/go/src/net/http/httputil/dump.go:202
//
//line /usr/local/go/src/net/http/httputil/dump.go:202
// The documentation for http.Request.Write details which fields
//line /usr/local/go/src/net/http/httputil/dump.go:202
// of req are included in the dump.
//line /usr/local/go/src/net/http/httputil/dump.go:218
func DumpRequest(req *http.Request, body bool) ([]byte, error) {
//line /usr/local/go/src/net/http/httputil/dump.go:218
	_go_fuzz_dep_.CoverTab[76115]++
							var err error
							save := req.Body
							if !body || func() bool {
//line /usr/local/go/src/net/http/httputil/dump.go:221
		_go_fuzz_dep_.CoverTab[76123]++
//line /usr/local/go/src/net/http/httputil/dump.go:221
		return req.Body == nil
//line /usr/local/go/src/net/http/httputil/dump.go:221
		// _ = "end of CoverTab[76123]"
//line /usr/local/go/src/net/http/httputil/dump.go:221
	}() {
//line /usr/local/go/src/net/http/httputil/dump.go:221
		_go_fuzz_dep_.CoverTab[76124]++
								req.Body = nil
//line /usr/local/go/src/net/http/httputil/dump.go:222
		// _ = "end of CoverTab[76124]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:223
		_go_fuzz_dep_.CoverTab[76125]++
								save, req.Body, err = drainBody(req.Body)
								if err != nil {
//line /usr/local/go/src/net/http/httputil/dump.go:225
			_go_fuzz_dep_.CoverTab[76126]++
									return nil, err
//line /usr/local/go/src/net/http/httputil/dump.go:226
			// _ = "end of CoverTab[76126]"
		} else {
//line /usr/local/go/src/net/http/httputil/dump.go:227
			_go_fuzz_dep_.CoverTab[76127]++
//line /usr/local/go/src/net/http/httputil/dump.go:227
			// _ = "end of CoverTab[76127]"
//line /usr/local/go/src/net/http/httputil/dump.go:227
		}
//line /usr/local/go/src/net/http/httputil/dump.go:227
		// _ = "end of CoverTab[76125]"
	}
//line /usr/local/go/src/net/http/httputil/dump.go:228
	// _ = "end of CoverTab[76115]"
//line /usr/local/go/src/net/http/httputil/dump.go:228
	_go_fuzz_dep_.CoverTab[76116]++

							var b bytes.Buffer

//line /usr/local/go/src/net/http/httputil/dump.go:238
	reqURI := req.RequestURI
	if reqURI == "" {
//line /usr/local/go/src/net/http/httputil/dump.go:239
		_go_fuzz_dep_.CoverTab[76128]++
								reqURI = req.URL.RequestURI()
//line /usr/local/go/src/net/http/httputil/dump.go:240
		// _ = "end of CoverTab[76128]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:241
		_go_fuzz_dep_.CoverTab[76129]++
//line /usr/local/go/src/net/http/httputil/dump.go:241
		// _ = "end of CoverTab[76129]"
//line /usr/local/go/src/net/http/httputil/dump.go:241
	}
//line /usr/local/go/src/net/http/httputil/dump.go:241
	// _ = "end of CoverTab[76116]"
//line /usr/local/go/src/net/http/httputil/dump.go:241
	_go_fuzz_dep_.CoverTab[76117]++

							fmt.Fprintf(&b, "%s %s HTTP/%d.%d\r\n", valueOrDefault(req.Method, "GET"),
		reqURI, req.ProtoMajor, req.ProtoMinor)

	absRequestURI := strings.HasPrefix(req.RequestURI, "http://") || func() bool {
//line /usr/local/go/src/net/http/httputil/dump.go:246
		_go_fuzz_dep_.CoverTab[76130]++
//line /usr/local/go/src/net/http/httputil/dump.go:246
		return strings.HasPrefix(req.RequestURI, "https://")
//line /usr/local/go/src/net/http/httputil/dump.go:246
		// _ = "end of CoverTab[76130]"
//line /usr/local/go/src/net/http/httputil/dump.go:246
	}()
							if !absRequestURI {
//line /usr/local/go/src/net/http/httputil/dump.go:247
		_go_fuzz_dep_.CoverTab[76131]++
								host := req.Host
								if host == "" && func() bool {
//line /usr/local/go/src/net/http/httputil/dump.go:249
			_go_fuzz_dep_.CoverTab[76133]++
//line /usr/local/go/src/net/http/httputil/dump.go:249
			return req.URL != nil
//line /usr/local/go/src/net/http/httputil/dump.go:249
			// _ = "end of CoverTab[76133]"
//line /usr/local/go/src/net/http/httputil/dump.go:249
		}() {
//line /usr/local/go/src/net/http/httputil/dump.go:249
			_go_fuzz_dep_.CoverTab[76134]++
									host = req.URL.Host
//line /usr/local/go/src/net/http/httputil/dump.go:250
			// _ = "end of CoverTab[76134]"
		} else {
//line /usr/local/go/src/net/http/httputil/dump.go:251
			_go_fuzz_dep_.CoverTab[76135]++
//line /usr/local/go/src/net/http/httputil/dump.go:251
			// _ = "end of CoverTab[76135]"
//line /usr/local/go/src/net/http/httputil/dump.go:251
		}
//line /usr/local/go/src/net/http/httputil/dump.go:251
		// _ = "end of CoverTab[76131]"
//line /usr/local/go/src/net/http/httputil/dump.go:251
		_go_fuzz_dep_.CoverTab[76132]++
								if host != "" {
//line /usr/local/go/src/net/http/httputil/dump.go:252
			_go_fuzz_dep_.CoverTab[76136]++
									fmt.Fprintf(&b, "Host: %s\r\n", host)
//line /usr/local/go/src/net/http/httputil/dump.go:253
			// _ = "end of CoverTab[76136]"
		} else {
//line /usr/local/go/src/net/http/httputil/dump.go:254
			_go_fuzz_dep_.CoverTab[76137]++
//line /usr/local/go/src/net/http/httputil/dump.go:254
			// _ = "end of CoverTab[76137]"
//line /usr/local/go/src/net/http/httputil/dump.go:254
		}
//line /usr/local/go/src/net/http/httputil/dump.go:254
		// _ = "end of CoverTab[76132]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:255
		_go_fuzz_dep_.CoverTab[76138]++
//line /usr/local/go/src/net/http/httputil/dump.go:255
		// _ = "end of CoverTab[76138]"
//line /usr/local/go/src/net/http/httputil/dump.go:255
	}
//line /usr/local/go/src/net/http/httputil/dump.go:255
	// _ = "end of CoverTab[76117]"
//line /usr/local/go/src/net/http/httputil/dump.go:255
	_go_fuzz_dep_.CoverTab[76118]++

							chunked := len(req.TransferEncoding) > 0 && func() bool {
//line /usr/local/go/src/net/http/httputil/dump.go:257
		_go_fuzz_dep_.CoverTab[76139]++
//line /usr/local/go/src/net/http/httputil/dump.go:257
		return req.TransferEncoding[0] == "chunked"
//line /usr/local/go/src/net/http/httputil/dump.go:257
		// _ = "end of CoverTab[76139]"
//line /usr/local/go/src/net/http/httputil/dump.go:257
	}()
							if len(req.TransferEncoding) > 0 {
//line /usr/local/go/src/net/http/httputil/dump.go:258
		_go_fuzz_dep_.CoverTab[76140]++
								fmt.Fprintf(&b, "Transfer-Encoding: %s\r\n", strings.Join(req.TransferEncoding, ","))
//line /usr/local/go/src/net/http/httputil/dump.go:259
		// _ = "end of CoverTab[76140]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:260
		_go_fuzz_dep_.CoverTab[76141]++
//line /usr/local/go/src/net/http/httputil/dump.go:260
		// _ = "end of CoverTab[76141]"
//line /usr/local/go/src/net/http/httputil/dump.go:260
	}
//line /usr/local/go/src/net/http/httputil/dump.go:260
	// _ = "end of CoverTab[76118]"
//line /usr/local/go/src/net/http/httputil/dump.go:260
	_go_fuzz_dep_.CoverTab[76119]++

							err = req.Header.WriteSubset(&b, reqWriteExcludeHeaderDump)
							if err != nil {
//line /usr/local/go/src/net/http/httputil/dump.go:263
		_go_fuzz_dep_.CoverTab[76142]++
								return nil, err
//line /usr/local/go/src/net/http/httputil/dump.go:264
		// _ = "end of CoverTab[76142]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:265
		_go_fuzz_dep_.CoverTab[76143]++
//line /usr/local/go/src/net/http/httputil/dump.go:265
		// _ = "end of CoverTab[76143]"
//line /usr/local/go/src/net/http/httputil/dump.go:265
	}
//line /usr/local/go/src/net/http/httputil/dump.go:265
	// _ = "end of CoverTab[76119]"
//line /usr/local/go/src/net/http/httputil/dump.go:265
	_go_fuzz_dep_.CoverTab[76120]++

							io.WriteString(&b, "\r\n")

							if req.Body != nil {
//line /usr/local/go/src/net/http/httputil/dump.go:269
		_go_fuzz_dep_.CoverTab[76144]++
								var dest io.Writer = &b
								if chunked {
//line /usr/local/go/src/net/http/httputil/dump.go:271
			_go_fuzz_dep_.CoverTab[76146]++
									dest = NewChunkedWriter(dest)
//line /usr/local/go/src/net/http/httputil/dump.go:272
			// _ = "end of CoverTab[76146]"
		} else {
//line /usr/local/go/src/net/http/httputil/dump.go:273
			_go_fuzz_dep_.CoverTab[76147]++
//line /usr/local/go/src/net/http/httputil/dump.go:273
			// _ = "end of CoverTab[76147]"
//line /usr/local/go/src/net/http/httputil/dump.go:273
		}
//line /usr/local/go/src/net/http/httputil/dump.go:273
		// _ = "end of CoverTab[76144]"
//line /usr/local/go/src/net/http/httputil/dump.go:273
		_go_fuzz_dep_.CoverTab[76145]++
								_, err = io.Copy(dest, req.Body)
								if chunked {
//line /usr/local/go/src/net/http/httputil/dump.go:275
			_go_fuzz_dep_.CoverTab[76148]++
									dest.(io.Closer).Close()
									io.WriteString(&b, "\r\n")
//line /usr/local/go/src/net/http/httputil/dump.go:277
			// _ = "end of CoverTab[76148]"
		} else {
//line /usr/local/go/src/net/http/httputil/dump.go:278
			_go_fuzz_dep_.CoverTab[76149]++
//line /usr/local/go/src/net/http/httputil/dump.go:278
			// _ = "end of CoverTab[76149]"
//line /usr/local/go/src/net/http/httputil/dump.go:278
		}
//line /usr/local/go/src/net/http/httputil/dump.go:278
		// _ = "end of CoverTab[76145]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:279
		_go_fuzz_dep_.CoverTab[76150]++
//line /usr/local/go/src/net/http/httputil/dump.go:279
		// _ = "end of CoverTab[76150]"
//line /usr/local/go/src/net/http/httputil/dump.go:279
	}
//line /usr/local/go/src/net/http/httputil/dump.go:279
	// _ = "end of CoverTab[76120]"
//line /usr/local/go/src/net/http/httputil/dump.go:279
	_go_fuzz_dep_.CoverTab[76121]++

							req.Body = save
							if err != nil {
//line /usr/local/go/src/net/http/httputil/dump.go:282
		_go_fuzz_dep_.CoverTab[76151]++
								return nil, err
//line /usr/local/go/src/net/http/httputil/dump.go:283
		// _ = "end of CoverTab[76151]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:284
		_go_fuzz_dep_.CoverTab[76152]++
//line /usr/local/go/src/net/http/httputil/dump.go:284
		// _ = "end of CoverTab[76152]"
//line /usr/local/go/src/net/http/httputil/dump.go:284
	}
//line /usr/local/go/src/net/http/httputil/dump.go:284
	// _ = "end of CoverTab[76121]"
//line /usr/local/go/src/net/http/httputil/dump.go:284
	_go_fuzz_dep_.CoverTab[76122]++
							return b.Bytes(), nil
//line /usr/local/go/src/net/http/httputil/dump.go:285
	// _ = "end of CoverTab[76122]"
}

// errNoBody is a sentinel error value used by failureToReadBody so we
//line /usr/local/go/src/net/http/httputil/dump.go:288
// can detect that the lack of body was intentional.
//line /usr/local/go/src/net/http/httputil/dump.go:290
var errNoBody = errors.New("sentinel error value")

// failureToReadBody is an io.ReadCloser that just returns errNoBody on
//line /usr/local/go/src/net/http/httputil/dump.go:292
// Read. It's swapped in when we don't actually want to consume
//line /usr/local/go/src/net/http/httputil/dump.go:292
// the body, but need a non-nil one, and want to distinguish the
//line /usr/local/go/src/net/http/httputil/dump.go:292
// error from reading the dummy body.
//line /usr/local/go/src/net/http/httputil/dump.go:296
type failureToReadBody struct{}

func (failureToReadBody) Read([]byte) (int, error) {
//line /usr/local/go/src/net/http/httputil/dump.go:298
	_go_fuzz_dep_.CoverTab[76153]++
//line /usr/local/go/src/net/http/httputil/dump.go:298
	return 0, errNoBody
//line /usr/local/go/src/net/http/httputil/dump.go:298
	// _ = "end of CoverTab[76153]"
//line /usr/local/go/src/net/http/httputil/dump.go:298
}
func (failureToReadBody) Close() error {
//line /usr/local/go/src/net/http/httputil/dump.go:299
	_go_fuzz_dep_.CoverTab[76154]++
//line /usr/local/go/src/net/http/httputil/dump.go:299
	return nil
//line /usr/local/go/src/net/http/httputil/dump.go:299
	// _ = "end of CoverTab[76154]"
//line /usr/local/go/src/net/http/httputil/dump.go:299
}

// emptyBody is an instance of empty reader.
var emptyBody = io.NopCloser(strings.NewReader(""))

// DumpResponse is like DumpRequest but dumps a response.
func DumpResponse(resp *http.Response, body bool) ([]byte, error) {
//line /usr/local/go/src/net/http/httputil/dump.go:305
	_go_fuzz_dep_.CoverTab[76155]++
							var b bytes.Buffer
							var err error
							save := resp.Body
							savecl := resp.ContentLength

							if !body {
//line /usr/local/go/src/net/http/httputil/dump.go:311
		_go_fuzz_dep_.CoverTab[76159]++

//line /usr/local/go/src/net/http/httputil/dump.go:314
		if resp.ContentLength == 0 {
//line /usr/local/go/src/net/http/httputil/dump.go:314
			_go_fuzz_dep_.CoverTab[76160]++
									resp.Body = emptyBody
//line /usr/local/go/src/net/http/httputil/dump.go:315
			// _ = "end of CoverTab[76160]"
		} else {
//line /usr/local/go/src/net/http/httputil/dump.go:316
			_go_fuzz_dep_.CoverTab[76161]++
									resp.Body = failureToReadBody{}
//line /usr/local/go/src/net/http/httputil/dump.go:317
			// _ = "end of CoverTab[76161]"
		}
//line /usr/local/go/src/net/http/httputil/dump.go:318
		// _ = "end of CoverTab[76159]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:319
		_go_fuzz_dep_.CoverTab[76162]++
//line /usr/local/go/src/net/http/httputil/dump.go:319
		if resp.Body == nil {
//line /usr/local/go/src/net/http/httputil/dump.go:319
			_go_fuzz_dep_.CoverTab[76163]++
									resp.Body = emptyBody
//line /usr/local/go/src/net/http/httputil/dump.go:320
			// _ = "end of CoverTab[76163]"
		} else {
//line /usr/local/go/src/net/http/httputil/dump.go:321
			_go_fuzz_dep_.CoverTab[76164]++
									save, resp.Body, err = drainBody(resp.Body)
									if err != nil {
//line /usr/local/go/src/net/http/httputil/dump.go:323
				_go_fuzz_dep_.CoverTab[76165]++
										return nil, err
//line /usr/local/go/src/net/http/httputil/dump.go:324
				// _ = "end of CoverTab[76165]"
			} else {
//line /usr/local/go/src/net/http/httputil/dump.go:325
				_go_fuzz_dep_.CoverTab[76166]++
//line /usr/local/go/src/net/http/httputil/dump.go:325
				// _ = "end of CoverTab[76166]"
//line /usr/local/go/src/net/http/httputil/dump.go:325
			}
//line /usr/local/go/src/net/http/httputil/dump.go:325
			// _ = "end of CoverTab[76164]"
		}
//line /usr/local/go/src/net/http/httputil/dump.go:326
		// _ = "end of CoverTab[76162]"
//line /usr/local/go/src/net/http/httputil/dump.go:326
	}
//line /usr/local/go/src/net/http/httputil/dump.go:326
	// _ = "end of CoverTab[76155]"
//line /usr/local/go/src/net/http/httputil/dump.go:326
	_go_fuzz_dep_.CoverTab[76156]++
							err = resp.Write(&b)
							if err == errNoBody {
//line /usr/local/go/src/net/http/httputil/dump.go:328
		_go_fuzz_dep_.CoverTab[76167]++
								err = nil
//line /usr/local/go/src/net/http/httputil/dump.go:329
		// _ = "end of CoverTab[76167]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:330
		_go_fuzz_dep_.CoverTab[76168]++
//line /usr/local/go/src/net/http/httputil/dump.go:330
		// _ = "end of CoverTab[76168]"
//line /usr/local/go/src/net/http/httputil/dump.go:330
	}
//line /usr/local/go/src/net/http/httputil/dump.go:330
	// _ = "end of CoverTab[76156]"
//line /usr/local/go/src/net/http/httputil/dump.go:330
	_go_fuzz_dep_.CoverTab[76157]++
							resp.Body = save
							resp.ContentLength = savecl
							if err != nil {
//line /usr/local/go/src/net/http/httputil/dump.go:333
		_go_fuzz_dep_.CoverTab[76169]++
								return nil, err
//line /usr/local/go/src/net/http/httputil/dump.go:334
		// _ = "end of CoverTab[76169]"
	} else {
//line /usr/local/go/src/net/http/httputil/dump.go:335
		_go_fuzz_dep_.CoverTab[76170]++
//line /usr/local/go/src/net/http/httputil/dump.go:335
		// _ = "end of CoverTab[76170]"
//line /usr/local/go/src/net/http/httputil/dump.go:335
	}
//line /usr/local/go/src/net/http/httputil/dump.go:335
	// _ = "end of CoverTab[76157]"
//line /usr/local/go/src/net/http/httputil/dump.go:335
	_go_fuzz_dep_.CoverTab[76158]++
							return b.Bytes(), nil
//line /usr/local/go/src/net/http/httputil/dump.go:336
	// _ = "end of CoverTab[76158]"
}

//line /usr/local/go/src/net/http/httputil/dump.go:337
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/http/httputil/dump.go:337
var _ = _go_fuzz_dep_.CoverTab
