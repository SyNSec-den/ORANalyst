// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows && !plan9

//line /usr/local/go/src/log/syslog/syslog_unix.go:7
package syslog

//line /usr/local/go/src/log/syslog/syslog_unix.go:7
import (
//line /usr/local/go/src/log/syslog/syslog_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/log/syslog/syslog_unix.go:7
)
//line /usr/local/go/src/log/syslog/syslog_unix.go:7
import (
//line /usr/local/go/src/log/syslog/syslog_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/log/syslog/syslog_unix.go:7
)

import (
	"errors"
	"net"
)

//line /usr/local/go/src/log/syslog/syslog_unix.go:17
func unixSyslog() (conn serverConn, err error) {
//line /usr/local/go/src/log/syslog/syslog_unix.go:17
	_go_fuzz_dep_.CoverTab[96068]++
							logTypes := []string{"unixgram", "unix"}
							logPaths := []string{"/dev/log", "/var/run/syslog", "/var/run/log"}
							for _, network := range logTypes {
//line /usr/local/go/src/log/syslog/syslog_unix.go:20
		_go_fuzz_dep_.CoverTab[96070]++
								for _, path := range logPaths {
//line /usr/local/go/src/log/syslog/syslog_unix.go:21
			_go_fuzz_dep_.CoverTab[96071]++
									conn, err := net.Dial(network, path)
									if err == nil {
//line /usr/local/go/src/log/syslog/syslog_unix.go:23
				_go_fuzz_dep_.CoverTab[96072]++
										return &netConn{conn: conn, local: true}, nil
//line /usr/local/go/src/log/syslog/syslog_unix.go:24
				// _ = "end of CoverTab[96072]"
			} else {
//line /usr/local/go/src/log/syslog/syslog_unix.go:25
				_go_fuzz_dep_.CoverTab[96073]++
//line /usr/local/go/src/log/syslog/syslog_unix.go:25
				// _ = "end of CoverTab[96073]"
//line /usr/local/go/src/log/syslog/syslog_unix.go:25
			}
//line /usr/local/go/src/log/syslog/syslog_unix.go:25
			// _ = "end of CoverTab[96071]"
		}
//line /usr/local/go/src/log/syslog/syslog_unix.go:26
		// _ = "end of CoverTab[96070]"
	}
//line /usr/local/go/src/log/syslog/syslog_unix.go:27
	// _ = "end of CoverTab[96068]"
//line /usr/local/go/src/log/syslog/syslog_unix.go:27
	_go_fuzz_dep_.CoverTab[96069]++
							return nil, errors.New("Unix syslog delivery error")
//line /usr/local/go/src/log/syslog/syslog_unix.go:28
	// _ = "end of CoverTab[96069]"
}

//line /usr/local/go/src/log/syslog/syslog_unix.go:29
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/log/syslog/syslog_unix.go:29
var _ = _go_fuzz_dep_.CoverTab
