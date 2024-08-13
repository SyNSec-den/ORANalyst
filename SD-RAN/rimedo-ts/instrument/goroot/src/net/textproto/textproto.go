// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/textproto/textproto.go:5
// Package textproto implements generic support for text-based request/response
//line /usr/local/go/src/net/textproto/textproto.go:5
// protocols in the style of HTTP, NNTP, and SMTP.
//line /usr/local/go/src/net/textproto/textproto.go:5
//
//line /usr/local/go/src/net/textproto/textproto.go:5
// The package provides:
//line /usr/local/go/src/net/textproto/textproto.go:5
//
//line /usr/local/go/src/net/textproto/textproto.go:5
// Error, which represents a numeric error response from
//line /usr/local/go/src/net/textproto/textproto.go:5
// a server.
//line /usr/local/go/src/net/textproto/textproto.go:5
//
//line /usr/local/go/src/net/textproto/textproto.go:5
// Pipeline, to manage pipelined requests and responses
//line /usr/local/go/src/net/textproto/textproto.go:5
// in a client.
//line /usr/local/go/src/net/textproto/textproto.go:5
//
//line /usr/local/go/src/net/textproto/textproto.go:5
// Reader, to read numeric response code lines,
//line /usr/local/go/src/net/textproto/textproto.go:5
// key: value headers, lines wrapped with leading spaces
//line /usr/local/go/src/net/textproto/textproto.go:5
// on continuation lines, and whole text blocks ending
//line /usr/local/go/src/net/textproto/textproto.go:5
// with a dot on a line by itself.
//line /usr/local/go/src/net/textproto/textproto.go:5
//
//line /usr/local/go/src/net/textproto/textproto.go:5
// Writer, to write dot-encoded text blocks.
//line /usr/local/go/src/net/textproto/textproto.go:5
//
//line /usr/local/go/src/net/textproto/textproto.go:5
// Conn, a convenient packaging of Reader, Writer, and Pipeline for use
//line /usr/local/go/src/net/textproto/textproto.go:5
// with a single network connection.
//line /usr/local/go/src/net/textproto/textproto.go:25
package textproto

//line /usr/local/go/src/net/textproto/textproto.go:25
import (
//line /usr/local/go/src/net/textproto/textproto.go:25
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/textproto/textproto.go:25
)
//line /usr/local/go/src/net/textproto/textproto.go:25
import (
//line /usr/local/go/src/net/textproto/textproto.go:25
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/textproto/textproto.go:25
)

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// An Error represents a numeric error response from a server.
type Error struct {
	Code	int
	Msg	string
}

func (e *Error) Error() string {
//line /usr/local/go/src/net/textproto/textproto.go:40
	_go_fuzz_dep_.CoverTab[34842]++
							return fmt.Sprintf("%03d %s", e.Code, e.Msg)
//line /usr/local/go/src/net/textproto/textproto.go:41
	// _ = "end of CoverTab[34842]"
}

// A ProtocolError describes a protocol violation such
//line /usr/local/go/src/net/textproto/textproto.go:44
// as an invalid response or a hung-up connection.
//line /usr/local/go/src/net/textproto/textproto.go:46
type ProtocolError string

func (p ProtocolError) Error() string {
//line /usr/local/go/src/net/textproto/textproto.go:48
	_go_fuzz_dep_.CoverTab[34843]++
							return string(p)
//line /usr/local/go/src/net/textproto/textproto.go:49
	// _ = "end of CoverTab[34843]"
}

// A Conn represents a textual network protocol connection.
//line /usr/local/go/src/net/textproto/textproto.go:52
// It consists of a Reader and Writer to manage I/O
//line /usr/local/go/src/net/textproto/textproto.go:52
// and a Pipeline to sequence concurrent requests on the connection.
//line /usr/local/go/src/net/textproto/textproto.go:52
// These embedded types carry methods with them;
//line /usr/local/go/src/net/textproto/textproto.go:52
// see the documentation of those types for details.
//line /usr/local/go/src/net/textproto/textproto.go:57
type Conn struct {
	Reader
	Writer
	Pipeline
	conn	io.ReadWriteCloser
}

// NewConn returns a new Conn using conn for I/O.
func NewConn(conn io.ReadWriteCloser) *Conn {
//line /usr/local/go/src/net/textproto/textproto.go:65
	_go_fuzz_dep_.CoverTab[34844]++
							return &Conn{
		Reader:	Reader{R: bufio.NewReader(conn)},
		Writer:	Writer{W: bufio.NewWriter(conn)},
		conn:	conn,
	}
//line /usr/local/go/src/net/textproto/textproto.go:70
	// _ = "end of CoverTab[34844]"
}

// Close closes the connection.
func (c *Conn) Close() error {
//line /usr/local/go/src/net/textproto/textproto.go:74
	_go_fuzz_dep_.CoverTab[34845]++
							return c.conn.Close()
//line /usr/local/go/src/net/textproto/textproto.go:75
	// _ = "end of CoverTab[34845]"
}

// Dial connects to the given address on the given network using net.Dial
//line /usr/local/go/src/net/textproto/textproto.go:78
// and then returns a new Conn for the connection.
//line /usr/local/go/src/net/textproto/textproto.go:80
func Dial(network, addr string) (*Conn, error) {
//line /usr/local/go/src/net/textproto/textproto.go:80
	_go_fuzz_dep_.CoverTab[34846]++
							c, err := net.Dial(network, addr)
							if err != nil {
//line /usr/local/go/src/net/textproto/textproto.go:82
		_go_fuzz_dep_.CoverTab[34848]++
								return nil, err
//line /usr/local/go/src/net/textproto/textproto.go:83
		// _ = "end of CoverTab[34848]"
	} else {
//line /usr/local/go/src/net/textproto/textproto.go:84
		_go_fuzz_dep_.CoverTab[34849]++
//line /usr/local/go/src/net/textproto/textproto.go:84
		// _ = "end of CoverTab[34849]"
//line /usr/local/go/src/net/textproto/textproto.go:84
	}
//line /usr/local/go/src/net/textproto/textproto.go:84
	// _ = "end of CoverTab[34846]"
//line /usr/local/go/src/net/textproto/textproto.go:84
	_go_fuzz_dep_.CoverTab[34847]++
							return NewConn(c), nil
//line /usr/local/go/src/net/textproto/textproto.go:85
	// _ = "end of CoverTab[34847]"
}

// Cmd is a convenience method that sends a command after
//line /usr/local/go/src/net/textproto/textproto.go:88
// waiting its turn in the pipeline. The command text is the
//line /usr/local/go/src/net/textproto/textproto.go:88
// result of formatting format with args and appending \r\n.
//line /usr/local/go/src/net/textproto/textproto.go:88
// Cmd returns the id of the command, for use with StartResponse and EndResponse.
//line /usr/local/go/src/net/textproto/textproto.go:88
//
//line /usr/local/go/src/net/textproto/textproto.go:88
// For example, a client might run a HELP command that returns a dot-body
//line /usr/local/go/src/net/textproto/textproto.go:88
// by using:
//line /usr/local/go/src/net/textproto/textproto.go:88
//
//line /usr/local/go/src/net/textproto/textproto.go:88
//	id, err := c.Cmd("HELP")
//line /usr/local/go/src/net/textproto/textproto.go:88
//	if err != nil {
//line /usr/local/go/src/net/textproto/textproto.go:88
//		return nil, err
//line /usr/local/go/src/net/textproto/textproto.go:88
//	}
//line /usr/local/go/src/net/textproto/textproto.go:88
//
//line /usr/local/go/src/net/textproto/textproto.go:88
//	c.StartResponse(id)
//line /usr/local/go/src/net/textproto/textproto.go:88
//	defer c.EndResponse(id)
//line /usr/local/go/src/net/textproto/textproto.go:88
//
//line /usr/local/go/src/net/textproto/textproto.go:88
//	if _, _, err = c.ReadCodeLine(110); err != nil {
//line /usr/local/go/src/net/textproto/textproto.go:88
//		return nil, err
//line /usr/local/go/src/net/textproto/textproto.go:88
//	}
//line /usr/local/go/src/net/textproto/textproto.go:88
//	text, err := c.ReadDotBytes()
//line /usr/local/go/src/net/textproto/textproto.go:88
//	if err != nil {
//line /usr/local/go/src/net/textproto/textproto.go:88
//		return nil, err
//line /usr/local/go/src/net/textproto/textproto.go:88
//	}
//line /usr/local/go/src/net/textproto/textproto.go:88
//	return c.ReadCodeLine(250)
//line /usr/local/go/src/net/textproto/textproto.go:112
func (c *Conn) Cmd(format string, args ...any) (id uint, err error) {
//line /usr/local/go/src/net/textproto/textproto.go:112
	_go_fuzz_dep_.CoverTab[34850]++
								id = c.Next()
								c.StartRequest(id)
								err = c.PrintfLine(format, args...)
								c.EndRequest(id)
								if err != nil {
//line /usr/local/go/src/net/textproto/textproto.go:117
		_go_fuzz_dep_.CoverTab[34852]++
									return 0, err
//line /usr/local/go/src/net/textproto/textproto.go:118
		// _ = "end of CoverTab[34852]"
	} else {
//line /usr/local/go/src/net/textproto/textproto.go:119
		_go_fuzz_dep_.CoverTab[34853]++
//line /usr/local/go/src/net/textproto/textproto.go:119
		// _ = "end of CoverTab[34853]"
//line /usr/local/go/src/net/textproto/textproto.go:119
	}
//line /usr/local/go/src/net/textproto/textproto.go:119
	// _ = "end of CoverTab[34850]"
//line /usr/local/go/src/net/textproto/textproto.go:119
	_go_fuzz_dep_.CoverTab[34851]++
								return id, nil
//line /usr/local/go/src/net/textproto/textproto.go:120
	// _ = "end of CoverTab[34851]"
}

// TrimString returns s without leading and trailing ASCII space.
func TrimString(s string) string {
//line /usr/local/go/src/net/textproto/textproto.go:124
	_go_fuzz_dep_.CoverTab[34854]++
								for len(s) > 0 && func() bool {
//line /usr/local/go/src/net/textproto/textproto.go:125
		_go_fuzz_dep_.CoverTab[34857]++
//line /usr/local/go/src/net/textproto/textproto.go:125
		return isASCIISpace(s[0])
//line /usr/local/go/src/net/textproto/textproto.go:125
		// _ = "end of CoverTab[34857]"
//line /usr/local/go/src/net/textproto/textproto.go:125
	}() {
//line /usr/local/go/src/net/textproto/textproto.go:125
		_go_fuzz_dep_.CoverTab[34858]++
									s = s[1:]
//line /usr/local/go/src/net/textproto/textproto.go:126
		// _ = "end of CoverTab[34858]"
	}
//line /usr/local/go/src/net/textproto/textproto.go:127
	// _ = "end of CoverTab[34854]"
//line /usr/local/go/src/net/textproto/textproto.go:127
	_go_fuzz_dep_.CoverTab[34855]++
								for len(s) > 0 && func() bool {
//line /usr/local/go/src/net/textproto/textproto.go:128
		_go_fuzz_dep_.CoverTab[34859]++
//line /usr/local/go/src/net/textproto/textproto.go:128
		return isASCIISpace(s[len(s)-1])
//line /usr/local/go/src/net/textproto/textproto.go:128
		// _ = "end of CoverTab[34859]"
//line /usr/local/go/src/net/textproto/textproto.go:128
	}() {
//line /usr/local/go/src/net/textproto/textproto.go:128
		_go_fuzz_dep_.CoverTab[34860]++
									s = s[:len(s)-1]
//line /usr/local/go/src/net/textproto/textproto.go:129
		// _ = "end of CoverTab[34860]"
	}
//line /usr/local/go/src/net/textproto/textproto.go:130
	// _ = "end of CoverTab[34855]"
//line /usr/local/go/src/net/textproto/textproto.go:130
	_go_fuzz_dep_.CoverTab[34856]++
								return s
//line /usr/local/go/src/net/textproto/textproto.go:131
	// _ = "end of CoverTab[34856]"
}

// TrimBytes returns b without leading and trailing ASCII space.
func TrimBytes(b []byte) []byte {
//line /usr/local/go/src/net/textproto/textproto.go:135
	_go_fuzz_dep_.CoverTab[34861]++
								for len(b) > 0 && func() bool {
//line /usr/local/go/src/net/textproto/textproto.go:136
		_go_fuzz_dep_.CoverTab[34864]++
//line /usr/local/go/src/net/textproto/textproto.go:136
		return isASCIISpace(b[0])
//line /usr/local/go/src/net/textproto/textproto.go:136
		// _ = "end of CoverTab[34864]"
//line /usr/local/go/src/net/textproto/textproto.go:136
	}() {
//line /usr/local/go/src/net/textproto/textproto.go:136
		_go_fuzz_dep_.CoverTab[34865]++
									b = b[1:]
//line /usr/local/go/src/net/textproto/textproto.go:137
		// _ = "end of CoverTab[34865]"
	}
//line /usr/local/go/src/net/textproto/textproto.go:138
	// _ = "end of CoverTab[34861]"
//line /usr/local/go/src/net/textproto/textproto.go:138
	_go_fuzz_dep_.CoverTab[34862]++
								for len(b) > 0 && func() bool {
//line /usr/local/go/src/net/textproto/textproto.go:139
		_go_fuzz_dep_.CoverTab[34866]++
//line /usr/local/go/src/net/textproto/textproto.go:139
		return isASCIISpace(b[len(b)-1])
//line /usr/local/go/src/net/textproto/textproto.go:139
		// _ = "end of CoverTab[34866]"
//line /usr/local/go/src/net/textproto/textproto.go:139
	}() {
//line /usr/local/go/src/net/textproto/textproto.go:139
		_go_fuzz_dep_.CoverTab[34867]++
									b = b[:len(b)-1]
//line /usr/local/go/src/net/textproto/textproto.go:140
		// _ = "end of CoverTab[34867]"
	}
//line /usr/local/go/src/net/textproto/textproto.go:141
	// _ = "end of CoverTab[34862]"
//line /usr/local/go/src/net/textproto/textproto.go:141
	_go_fuzz_dep_.CoverTab[34863]++
								return b
//line /usr/local/go/src/net/textproto/textproto.go:142
	// _ = "end of CoverTab[34863]"
}

func isASCIISpace(b byte) bool {
//line /usr/local/go/src/net/textproto/textproto.go:145
	_go_fuzz_dep_.CoverTab[34868]++
								return b == ' ' || func() bool {
//line /usr/local/go/src/net/textproto/textproto.go:146
		_go_fuzz_dep_.CoverTab[34869]++
//line /usr/local/go/src/net/textproto/textproto.go:146
		return b == '\t'
//line /usr/local/go/src/net/textproto/textproto.go:146
		// _ = "end of CoverTab[34869]"
//line /usr/local/go/src/net/textproto/textproto.go:146
	}() || func() bool {
//line /usr/local/go/src/net/textproto/textproto.go:146
		_go_fuzz_dep_.CoverTab[34870]++
//line /usr/local/go/src/net/textproto/textproto.go:146
		return b == '\n'
//line /usr/local/go/src/net/textproto/textproto.go:146
		// _ = "end of CoverTab[34870]"
//line /usr/local/go/src/net/textproto/textproto.go:146
	}() || func() bool {
//line /usr/local/go/src/net/textproto/textproto.go:146
		_go_fuzz_dep_.CoverTab[34871]++
//line /usr/local/go/src/net/textproto/textproto.go:146
		return b == '\r'
//line /usr/local/go/src/net/textproto/textproto.go:146
		// _ = "end of CoverTab[34871]"
//line /usr/local/go/src/net/textproto/textproto.go:146
	}()
//line /usr/local/go/src/net/textproto/textproto.go:146
	// _ = "end of CoverTab[34868]"
}

func isASCIILetter(b byte) bool {
//line /usr/local/go/src/net/textproto/textproto.go:149
	_go_fuzz_dep_.CoverTab[34872]++
								b |= 0x20
								return 'a' <= b && func() bool {
//line /usr/local/go/src/net/textproto/textproto.go:151
		_go_fuzz_dep_.CoverTab[34873]++
//line /usr/local/go/src/net/textproto/textproto.go:151
		return b <= 'z'
//line /usr/local/go/src/net/textproto/textproto.go:151
		// _ = "end of CoverTab[34873]"
//line /usr/local/go/src/net/textproto/textproto.go:151
	}()
//line /usr/local/go/src/net/textproto/textproto.go:151
	// _ = "end of CoverTab[34872]"
}

//line /usr/local/go/src/net/textproto/textproto.go:152
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/textproto/textproto.go:152
var _ = _go_fuzz_dep_.CoverTab
