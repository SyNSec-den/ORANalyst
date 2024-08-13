// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/log/syslog/doc.go:5
// Package syslog provides a simple interface to the system log
//line /usr/local/go/src/log/syslog/doc.go:5
// service. It can send messages to the syslog daemon using UNIX
//line /usr/local/go/src/log/syslog/doc.go:5
// domain sockets, UDP or TCP.
//line /usr/local/go/src/log/syslog/doc.go:5
//
//line /usr/local/go/src/log/syslog/doc.go:5
// Only one call to Dial is necessary. On write failures,
//line /usr/local/go/src/log/syslog/doc.go:5
// the syslog client will attempt to reconnect to the server
//line /usr/local/go/src/log/syslog/doc.go:5
// and write again.
//line /usr/local/go/src/log/syslog/doc.go:5
//
//line /usr/local/go/src/log/syslog/doc.go:5
// The syslog package is frozen and is not accepting new features.
//line /usr/local/go/src/log/syslog/doc.go:5
// Some external packages provide more functionality. See:
//line /usr/local/go/src/log/syslog/doc.go:5
//
//line /usr/local/go/src/log/syslog/doc.go:5
//	https://godoc.org/?q=syslog
//line /usr/local/go/src/log/syslog/doc.go:17
package syslog

//line /usr/local/go/src/log/syslog/doc.go:17
import (
//line /usr/local/go/src/log/syslog/doc.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/log/syslog/doc.go:17
)
//line /usr/local/go/src/log/syslog/doc.go:17
import (
//line /usr/local/go/src/log/syslog/doc.go:17
	_atomic_ "sync/atomic"
//line /usr/local/go/src/log/syslog/doc.go:17
)

//line /usr/local/go/src/log/syslog/doc.go:17
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/log/syslog/doc.go:17
var _ = _go_fuzz_dep_.CoverTab
