// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows && !plan9

//line /usr/local/go/src/log/syslog/syslog.go:7
package syslog

//line /usr/local/go/src/log/syslog/syslog.go:7
import (
//line /usr/local/go/src/log/syslog/syslog.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/log/syslog/syslog.go:7
)
//line /usr/local/go/src/log/syslog/syslog.go:7
import (
//line /usr/local/go/src/log/syslog/syslog.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/log/syslog/syslog.go:7
)

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

// The Priority is a combination of the syslog facility and
//line /usr/local/go/src/log/syslog/syslog.go:20
// severity. For example, LOG_ALERT | LOG_FTP sends an alert severity
//line /usr/local/go/src/log/syslog/syslog.go:20
// message from the FTP facility. The default severity is LOG_EMERG;
//line /usr/local/go/src/log/syslog/syslog.go:20
// the default facility is LOG_KERN.
//line /usr/local/go/src/log/syslog/syslog.go:24
type Priority int

const severityMask = 0x07
const facilityMask = 0xf8

const (

//line /usr/local/go/src/log/syslog/syslog.go:32
	// From /usr/include/sys/syslog.h.
	// These are the same on Linux, BSD, and OS X.
	LOG_EMERG	Priority	= iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

const (

//line /usr/local/go/src/log/syslog/syslog.go:47
	// From /usr/include/sys/syslog.h.
	// These are the same up to LOG_FTP on Linux, BSD, and OS X.
	LOG_KERN	Priority	= iota << 3
	LOG_USER
	LOG_MAIL
	LOG_DAEMON
	LOG_AUTH
	LOG_SYSLOG
	LOG_LPR
	LOG_NEWS
	LOG_UUCP
	LOG_CRON
	LOG_AUTHPRIV
	LOG_FTP
	_	// unused
	_	// unused
	_	// unused
	_	// unused
	LOG_LOCAL0
	LOG_LOCAL1
	LOG_LOCAL2
	LOG_LOCAL3
	LOG_LOCAL4
	LOG_LOCAL5
	LOG_LOCAL6
	LOG_LOCAL7
)

// A Writer is a connection to a syslog server.
type Writer struct {
	priority	Priority
	tag		string
	hostname	string
	network		string
	raddr		string

	mu	sync.Mutex	// guards conn
	conn	serverConn
}

// This interface and the separate syslog_unix.go file exist for
//line /usr/local/go/src/log/syslog/syslog.go:87
// Solaris support as implemented by gccgo. On Solaris you cannot
//line /usr/local/go/src/log/syslog/syslog.go:87
// simply open a TCP connection to the syslog daemon. The gccgo
//line /usr/local/go/src/log/syslog/syslog.go:87
// sources have a syslog_solaris.go file that implements unixSyslog to
//line /usr/local/go/src/log/syslog/syslog.go:87
// return a type that satisfies this interface and simply calls the C
//line /usr/local/go/src/log/syslog/syslog.go:87
// library syslog function.
//line /usr/local/go/src/log/syslog/syslog.go:93
type serverConn interface {
	writeString(p Priority, hostname, tag, s, nl string) error
	close() error
}

type netConn struct {
	local	bool
	conn	net.Conn
}

// New establishes a new connection to the system log daemon. Each
//line /usr/local/go/src/log/syslog/syslog.go:103
// write to the returned writer sends a log message with the given
//line /usr/local/go/src/log/syslog/syslog.go:103
// priority (a combination of the syslog facility and severity) and
//line /usr/local/go/src/log/syslog/syslog.go:103
// prefix tag. If tag is empty, the os.Args[0] is used.
//line /usr/local/go/src/log/syslog/syslog.go:107
func New(priority Priority, tag string) (*Writer, error) {
//line /usr/local/go/src/log/syslog/syslog.go:107
	_go_fuzz_dep_.CoverTab[96004]++
							return Dial("", "", priority, tag)
//line /usr/local/go/src/log/syslog/syslog.go:108
	// _ = "end of CoverTab[96004]"
}

// Dial establishes a connection to a log daemon by connecting to
//line /usr/local/go/src/log/syslog/syslog.go:111
// address raddr on the specified network. Each write to the returned
//line /usr/local/go/src/log/syslog/syslog.go:111
// writer sends a log message with the facility and severity
//line /usr/local/go/src/log/syslog/syslog.go:111
// (from priority) and tag. If tag is empty, the os.Args[0] is used.
//line /usr/local/go/src/log/syslog/syslog.go:111
// If network is empty, Dial will connect to the local syslog server.
//line /usr/local/go/src/log/syslog/syslog.go:111
// Otherwise, see the documentation for net.Dial for valid values
//line /usr/local/go/src/log/syslog/syslog.go:111
// of network and raddr.
//line /usr/local/go/src/log/syslog/syslog.go:118
func Dial(network, raddr string, priority Priority, tag string) (*Writer, error) {
//line /usr/local/go/src/log/syslog/syslog.go:118
	_go_fuzz_dep_.CoverTab[96005]++
							if priority < 0 || func() bool {
//line /usr/local/go/src/log/syslog/syslog.go:119
		_go_fuzz_dep_.CoverTab[96009]++
//line /usr/local/go/src/log/syslog/syslog.go:119
		return priority > LOG_LOCAL7|LOG_DEBUG
//line /usr/local/go/src/log/syslog/syslog.go:119
		// _ = "end of CoverTab[96009]"
//line /usr/local/go/src/log/syslog/syslog.go:119
	}() {
//line /usr/local/go/src/log/syslog/syslog.go:119
		_go_fuzz_dep_.CoverTab[96010]++
								return nil, errors.New("log/syslog: invalid priority")
//line /usr/local/go/src/log/syslog/syslog.go:120
		// _ = "end of CoverTab[96010]"
	} else {
//line /usr/local/go/src/log/syslog/syslog.go:121
		_go_fuzz_dep_.CoverTab[96011]++
//line /usr/local/go/src/log/syslog/syslog.go:121
		// _ = "end of CoverTab[96011]"
//line /usr/local/go/src/log/syslog/syslog.go:121
	}
//line /usr/local/go/src/log/syslog/syslog.go:121
	// _ = "end of CoverTab[96005]"
//line /usr/local/go/src/log/syslog/syslog.go:121
	_go_fuzz_dep_.CoverTab[96006]++

							if tag == "" {
//line /usr/local/go/src/log/syslog/syslog.go:123
		_go_fuzz_dep_.CoverTab[96012]++
								tag = os.Args[0]
//line /usr/local/go/src/log/syslog/syslog.go:124
		// _ = "end of CoverTab[96012]"
	} else {
//line /usr/local/go/src/log/syslog/syslog.go:125
		_go_fuzz_dep_.CoverTab[96013]++
//line /usr/local/go/src/log/syslog/syslog.go:125
		// _ = "end of CoverTab[96013]"
//line /usr/local/go/src/log/syslog/syslog.go:125
	}
//line /usr/local/go/src/log/syslog/syslog.go:125
	// _ = "end of CoverTab[96006]"
//line /usr/local/go/src/log/syslog/syslog.go:125
	_go_fuzz_dep_.CoverTab[96007]++
							hostname, _ := os.Hostname()

							w := &Writer{
		priority:	priority,
		tag:		tag,
		hostname:	hostname,
		network:	network,
		raddr:		raddr,
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	err := w.connect()
	if err != nil {
//line /usr/local/go/src/log/syslog/syslog.go:140
		_go_fuzz_dep_.CoverTab[96014]++
								return nil, err
//line /usr/local/go/src/log/syslog/syslog.go:141
		// _ = "end of CoverTab[96014]"
	} else {
//line /usr/local/go/src/log/syslog/syslog.go:142
		_go_fuzz_dep_.CoverTab[96015]++
//line /usr/local/go/src/log/syslog/syslog.go:142
		// _ = "end of CoverTab[96015]"
//line /usr/local/go/src/log/syslog/syslog.go:142
	}
//line /usr/local/go/src/log/syslog/syslog.go:142
	// _ = "end of CoverTab[96007]"
//line /usr/local/go/src/log/syslog/syslog.go:142
	_go_fuzz_dep_.CoverTab[96008]++
							return w, err
//line /usr/local/go/src/log/syslog/syslog.go:143
	// _ = "end of CoverTab[96008]"
}

// connect makes a connection to the syslog server.
//line /usr/local/go/src/log/syslog/syslog.go:146
// It must be called with w.mu held.
//line /usr/local/go/src/log/syslog/syslog.go:148
func (w *Writer) connect() (err error) {
//line /usr/local/go/src/log/syslog/syslog.go:148
	_go_fuzz_dep_.CoverTab[96016]++
							if w.conn != nil {
//line /usr/local/go/src/log/syslog/syslog.go:149
		_go_fuzz_dep_.CoverTab[96019]++

								w.conn.close()
								w.conn = nil
//line /usr/local/go/src/log/syslog/syslog.go:152
		// _ = "end of CoverTab[96019]"
	} else {
//line /usr/local/go/src/log/syslog/syslog.go:153
		_go_fuzz_dep_.CoverTab[96020]++
//line /usr/local/go/src/log/syslog/syslog.go:153
		// _ = "end of CoverTab[96020]"
//line /usr/local/go/src/log/syslog/syslog.go:153
	}
//line /usr/local/go/src/log/syslog/syslog.go:153
	// _ = "end of CoverTab[96016]"
//line /usr/local/go/src/log/syslog/syslog.go:153
	_go_fuzz_dep_.CoverTab[96017]++

							if w.network == "" {
//line /usr/local/go/src/log/syslog/syslog.go:155
		_go_fuzz_dep_.CoverTab[96021]++
								w.conn, err = unixSyslog()
								if w.hostname == "" {
//line /usr/local/go/src/log/syslog/syslog.go:157
			_go_fuzz_dep_.CoverTab[96022]++
									w.hostname = "localhost"
//line /usr/local/go/src/log/syslog/syslog.go:158
			// _ = "end of CoverTab[96022]"
		} else {
//line /usr/local/go/src/log/syslog/syslog.go:159
			_go_fuzz_dep_.CoverTab[96023]++
//line /usr/local/go/src/log/syslog/syslog.go:159
			// _ = "end of CoverTab[96023]"
//line /usr/local/go/src/log/syslog/syslog.go:159
		}
//line /usr/local/go/src/log/syslog/syslog.go:159
		// _ = "end of CoverTab[96021]"
	} else {
//line /usr/local/go/src/log/syslog/syslog.go:160
		_go_fuzz_dep_.CoverTab[96024]++
								var c net.Conn
								c, err = net.Dial(w.network, w.raddr)
								if err == nil {
//line /usr/local/go/src/log/syslog/syslog.go:163
			_go_fuzz_dep_.CoverTab[96025]++
									w.conn = &netConn{
				conn:	c,
				local: w.network == "unixgram" || func() bool {
//line /usr/local/go/src/log/syslog/syslog.go:166
					_go_fuzz_dep_.CoverTab[96026]++
//line /usr/local/go/src/log/syslog/syslog.go:166
					return w.network == "unix"
//line /usr/local/go/src/log/syslog/syslog.go:166
					// _ = "end of CoverTab[96026]"
//line /usr/local/go/src/log/syslog/syslog.go:166
				}(),
			}
			if w.hostname == "" {
//line /usr/local/go/src/log/syslog/syslog.go:168
				_go_fuzz_dep_.CoverTab[96027]++
										w.hostname = c.LocalAddr().String()
//line /usr/local/go/src/log/syslog/syslog.go:169
				// _ = "end of CoverTab[96027]"
			} else {
//line /usr/local/go/src/log/syslog/syslog.go:170
				_go_fuzz_dep_.CoverTab[96028]++
//line /usr/local/go/src/log/syslog/syslog.go:170
				// _ = "end of CoverTab[96028]"
//line /usr/local/go/src/log/syslog/syslog.go:170
			}
//line /usr/local/go/src/log/syslog/syslog.go:170
			// _ = "end of CoverTab[96025]"
		} else {
//line /usr/local/go/src/log/syslog/syslog.go:171
			_go_fuzz_dep_.CoverTab[96029]++
//line /usr/local/go/src/log/syslog/syslog.go:171
			// _ = "end of CoverTab[96029]"
//line /usr/local/go/src/log/syslog/syslog.go:171
		}
//line /usr/local/go/src/log/syslog/syslog.go:171
		// _ = "end of CoverTab[96024]"
	}
//line /usr/local/go/src/log/syslog/syslog.go:172
	// _ = "end of CoverTab[96017]"
//line /usr/local/go/src/log/syslog/syslog.go:172
	_go_fuzz_dep_.CoverTab[96018]++
							return
//line /usr/local/go/src/log/syslog/syslog.go:173
	// _ = "end of CoverTab[96018]"
}

// Write sends a log message to the syslog daemon.
func (w *Writer) Write(b []byte) (int, error) {
//line /usr/local/go/src/log/syslog/syslog.go:177
	_go_fuzz_dep_.CoverTab[96030]++
							return w.writeAndRetry(w.priority, string(b))
//line /usr/local/go/src/log/syslog/syslog.go:178
	// _ = "end of CoverTab[96030]"
}

// Close closes a connection to the syslog daemon.
func (w *Writer) Close() error {
//line /usr/local/go/src/log/syslog/syslog.go:182
	_go_fuzz_dep_.CoverTab[96031]++
							w.mu.Lock()
							defer w.mu.Unlock()

							if w.conn != nil {
//line /usr/local/go/src/log/syslog/syslog.go:186
		_go_fuzz_dep_.CoverTab[96033]++
								err := w.conn.close()
								w.conn = nil
								return err
//line /usr/local/go/src/log/syslog/syslog.go:189
		// _ = "end of CoverTab[96033]"
	} else {
//line /usr/local/go/src/log/syslog/syslog.go:190
		_go_fuzz_dep_.CoverTab[96034]++
//line /usr/local/go/src/log/syslog/syslog.go:190
		// _ = "end of CoverTab[96034]"
//line /usr/local/go/src/log/syslog/syslog.go:190
	}
//line /usr/local/go/src/log/syslog/syslog.go:190
	// _ = "end of CoverTab[96031]"
//line /usr/local/go/src/log/syslog/syslog.go:190
	_go_fuzz_dep_.CoverTab[96032]++
							return nil
//line /usr/local/go/src/log/syslog/syslog.go:191
	// _ = "end of CoverTab[96032]"
}

// Emerg logs a message with severity LOG_EMERG, ignoring the severity
//line /usr/local/go/src/log/syslog/syslog.go:194
// passed to New.
//line /usr/local/go/src/log/syslog/syslog.go:196
func (w *Writer) Emerg(m string) error {
//line /usr/local/go/src/log/syslog/syslog.go:196
	_go_fuzz_dep_.CoverTab[96035]++
							_, err := w.writeAndRetry(LOG_EMERG, m)
							return err
//line /usr/local/go/src/log/syslog/syslog.go:198
	// _ = "end of CoverTab[96035]"
}

// Alert logs a message with severity LOG_ALERT, ignoring the severity
//line /usr/local/go/src/log/syslog/syslog.go:201
// passed to New.
//line /usr/local/go/src/log/syslog/syslog.go:203
func (w *Writer) Alert(m string) error {
//line /usr/local/go/src/log/syslog/syslog.go:203
	_go_fuzz_dep_.CoverTab[96036]++
							_, err := w.writeAndRetry(LOG_ALERT, m)
							return err
//line /usr/local/go/src/log/syslog/syslog.go:205
	// _ = "end of CoverTab[96036]"
}

// Crit logs a message with severity LOG_CRIT, ignoring the severity
//line /usr/local/go/src/log/syslog/syslog.go:208
// passed to New.
//line /usr/local/go/src/log/syslog/syslog.go:210
func (w *Writer) Crit(m string) error {
//line /usr/local/go/src/log/syslog/syslog.go:210
	_go_fuzz_dep_.CoverTab[96037]++
							_, err := w.writeAndRetry(LOG_CRIT, m)
							return err
//line /usr/local/go/src/log/syslog/syslog.go:212
	// _ = "end of CoverTab[96037]"
}

// Err logs a message with severity LOG_ERR, ignoring the severity
//line /usr/local/go/src/log/syslog/syslog.go:215
// passed to New.
//line /usr/local/go/src/log/syslog/syslog.go:217
func (w *Writer) Err(m string) error {
//line /usr/local/go/src/log/syslog/syslog.go:217
	_go_fuzz_dep_.CoverTab[96038]++
							_, err := w.writeAndRetry(LOG_ERR, m)
							return err
//line /usr/local/go/src/log/syslog/syslog.go:219
	// _ = "end of CoverTab[96038]"
}

// Warning logs a message with severity LOG_WARNING, ignoring the
//line /usr/local/go/src/log/syslog/syslog.go:222
// severity passed to New.
//line /usr/local/go/src/log/syslog/syslog.go:224
func (w *Writer) Warning(m string) error {
//line /usr/local/go/src/log/syslog/syslog.go:224
	_go_fuzz_dep_.CoverTab[96039]++
							_, err := w.writeAndRetry(LOG_WARNING, m)
							return err
//line /usr/local/go/src/log/syslog/syslog.go:226
	// _ = "end of CoverTab[96039]"
}

// Notice logs a message with severity LOG_NOTICE, ignoring the
//line /usr/local/go/src/log/syslog/syslog.go:229
// severity passed to New.
//line /usr/local/go/src/log/syslog/syslog.go:231
func (w *Writer) Notice(m string) error {
//line /usr/local/go/src/log/syslog/syslog.go:231
	_go_fuzz_dep_.CoverTab[96040]++
							_, err := w.writeAndRetry(LOG_NOTICE, m)
							return err
//line /usr/local/go/src/log/syslog/syslog.go:233
	// _ = "end of CoverTab[96040]"
}

// Info logs a message with severity LOG_INFO, ignoring the severity
//line /usr/local/go/src/log/syslog/syslog.go:236
// passed to New.
//line /usr/local/go/src/log/syslog/syslog.go:238
func (w *Writer) Info(m string) error {
//line /usr/local/go/src/log/syslog/syslog.go:238
	_go_fuzz_dep_.CoverTab[96041]++
							_, err := w.writeAndRetry(LOG_INFO, m)
							return err
//line /usr/local/go/src/log/syslog/syslog.go:240
	// _ = "end of CoverTab[96041]"
}

// Debug logs a message with severity LOG_DEBUG, ignoring the severity
//line /usr/local/go/src/log/syslog/syslog.go:243
// passed to New.
//line /usr/local/go/src/log/syslog/syslog.go:245
func (w *Writer) Debug(m string) error {
//line /usr/local/go/src/log/syslog/syslog.go:245
	_go_fuzz_dep_.CoverTab[96042]++
							_, err := w.writeAndRetry(LOG_DEBUG, m)
							return err
//line /usr/local/go/src/log/syslog/syslog.go:247
	// _ = "end of CoverTab[96042]"
}

func (w *Writer) writeAndRetry(p Priority, s string) (int, error) {
//line /usr/local/go/src/log/syslog/syslog.go:250
	_go_fuzz_dep_.CoverTab[96043]++
							pr := (w.priority & facilityMask) | (p & severityMask)

							w.mu.Lock()
							defer w.mu.Unlock()

							if w.conn != nil {
//line /usr/local/go/src/log/syslog/syslog.go:256
		_go_fuzz_dep_.CoverTab[96046]++
								if n, err := w.write(pr, s); err == nil {
//line /usr/local/go/src/log/syslog/syslog.go:257
			_go_fuzz_dep_.CoverTab[96047]++
									return n, nil
//line /usr/local/go/src/log/syslog/syslog.go:258
			// _ = "end of CoverTab[96047]"
		} else {
//line /usr/local/go/src/log/syslog/syslog.go:259
			_go_fuzz_dep_.CoverTab[96048]++
//line /usr/local/go/src/log/syslog/syslog.go:259
			// _ = "end of CoverTab[96048]"
//line /usr/local/go/src/log/syslog/syslog.go:259
		}
//line /usr/local/go/src/log/syslog/syslog.go:259
		// _ = "end of CoverTab[96046]"
	} else {
//line /usr/local/go/src/log/syslog/syslog.go:260
		_go_fuzz_dep_.CoverTab[96049]++
//line /usr/local/go/src/log/syslog/syslog.go:260
		// _ = "end of CoverTab[96049]"
//line /usr/local/go/src/log/syslog/syslog.go:260
	}
//line /usr/local/go/src/log/syslog/syslog.go:260
	// _ = "end of CoverTab[96043]"
//line /usr/local/go/src/log/syslog/syslog.go:260
	_go_fuzz_dep_.CoverTab[96044]++
							if err := w.connect(); err != nil {
//line /usr/local/go/src/log/syslog/syslog.go:261
		_go_fuzz_dep_.CoverTab[96050]++
								return 0, err
//line /usr/local/go/src/log/syslog/syslog.go:262
		// _ = "end of CoverTab[96050]"
	} else {
//line /usr/local/go/src/log/syslog/syslog.go:263
		_go_fuzz_dep_.CoverTab[96051]++
//line /usr/local/go/src/log/syslog/syslog.go:263
		// _ = "end of CoverTab[96051]"
//line /usr/local/go/src/log/syslog/syslog.go:263
	}
//line /usr/local/go/src/log/syslog/syslog.go:263
	// _ = "end of CoverTab[96044]"
//line /usr/local/go/src/log/syslog/syslog.go:263
	_go_fuzz_dep_.CoverTab[96045]++
							return w.write(pr, s)
//line /usr/local/go/src/log/syslog/syslog.go:264
	// _ = "end of CoverTab[96045]"
}

// write generates and writes a syslog formatted string. The
//line /usr/local/go/src/log/syslog/syslog.go:267
// format is as follows: <PRI>TIMESTAMP HOSTNAME TAG[PID]: MSG
//line /usr/local/go/src/log/syslog/syslog.go:269
func (w *Writer) write(p Priority, msg string) (int, error) {
//line /usr/local/go/src/log/syslog/syslog.go:269
	_go_fuzz_dep_.CoverTab[96052]++

							nl := ""
							if !strings.HasSuffix(msg, "\n") {
//line /usr/local/go/src/log/syslog/syslog.go:272
		_go_fuzz_dep_.CoverTab[96055]++
								nl = "\n"
//line /usr/local/go/src/log/syslog/syslog.go:273
		// _ = "end of CoverTab[96055]"
	} else {
//line /usr/local/go/src/log/syslog/syslog.go:274
		_go_fuzz_dep_.CoverTab[96056]++
//line /usr/local/go/src/log/syslog/syslog.go:274
		// _ = "end of CoverTab[96056]"
//line /usr/local/go/src/log/syslog/syslog.go:274
	}
//line /usr/local/go/src/log/syslog/syslog.go:274
	// _ = "end of CoverTab[96052]"
//line /usr/local/go/src/log/syslog/syslog.go:274
	_go_fuzz_dep_.CoverTab[96053]++

							err := w.conn.writeString(p, w.hostname, w.tag, msg, nl)
							if err != nil {
//line /usr/local/go/src/log/syslog/syslog.go:277
		_go_fuzz_dep_.CoverTab[96057]++
								return 0, err
//line /usr/local/go/src/log/syslog/syslog.go:278
		// _ = "end of CoverTab[96057]"
	} else {
//line /usr/local/go/src/log/syslog/syslog.go:279
		_go_fuzz_dep_.CoverTab[96058]++
//line /usr/local/go/src/log/syslog/syslog.go:279
		// _ = "end of CoverTab[96058]"
//line /usr/local/go/src/log/syslog/syslog.go:279
	}
//line /usr/local/go/src/log/syslog/syslog.go:279
	// _ = "end of CoverTab[96053]"
//line /usr/local/go/src/log/syslog/syslog.go:279
	_go_fuzz_dep_.CoverTab[96054]++

//line /usr/local/go/src/log/syslog/syslog.go:283
	return len(msg), nil
//line /usr/local/go/src/log/syslog/syslog.go:283
	// _ = "end of CoverTab[96054]"
}

func (n *netConn) writeString(p Priority, hostname, tag, msg, nl string) error {
//line /usr/local/go/src/log/syslog/syslog.go:286
	_go_fuzz_dep_.CoverTab[96059]++
							if n.local {
//line /usr/local/go/src/log/syslog/syslog.go:287
		_go_fuzz_dep_.CoverTab[96061]++

//line /usr/local/go/src/log/syslog/syslog.go:291
		timestamp := time.Now().Format(time.Stamp)
		_, err := fmt.Fprintf(n.conn, "<%d>%s %s[%d]: %s%s",
			p, timestamp,
			tag, os.Getpid(), msg, nl)
								return err
//line /usr/local/go/src/log/syslog/syslog.go:295
		// _ = "end of CoverTab[96061]"
	} else {
//line /usr/local/go/src/log/syslog/syslog.go:296
		_go_fuzz_dep_.CoverTab[96062]++
//line /usr/local/go/src/log/syslog/syslog.go:296
		// _ = "end of CoverTab[96062]"
//line /usr/local/go/src/log/syslog/syslog.go:296
	}
//line /usr/local/go/src/log/syslog/syslog.go:296
	// _ = "end of CoverTab[96059]"
//line /usr/local/go/src/log/syslog/syslog.go:296
	_go_fuzz_dep_.CoverTab[96060]++
							timestamp := time.Now().Format(time.RFC3339)
							_, err := fmt.Fprintf(n.conn, "<%d>%s %s %s[%d]: %s%s",
		p, timestamp, hostname,
		tag, os.Getpid(), msg, nl)
							return err
//line /usr/local/go/src/log/syslog/syslog.go:301
	// _ = "end of CoverTab[96060]"
}

func (n *netConn) close() error {
//line /usr/local/go/src/log/syslog/syslog.go:304
	_go_fuzz_dep_.CoverTab[96063]++
							return n.conn.Close()
//line /usr/local/go/src/log/syslog/syslog.go:305
	// _ = "end of CoverTab[96063]"
}

// NewLogger creates a log.Logger whose output is written to the
//line /usr/local/go/src/log/syslog/syslog.go:308
// system log service with the specified priority, a combination of
//line /usr/local/go/src/log/syslog/syslog.go:308
// the syslog facility and severity. The logFlag argument is the flag
//line /usr/local/go/src/log/syslog/syslog.go:308
// set passed through to log.New to create the Logger.
//line /usr/local/go/src/log/syslog/syslog.go:312
func NewLogger(p Priority, logFlag int) (*log.Logger, error) {
//line /usr/local/go/src/log/syslog/syslog.go:312
	_go_fuzz_dep_.CoverTab[96064]++
							s, err := New(p, "")
							if err != nil {
//line /usr/local/go/src/log/syslog/syslog.go:314
		_go_fuzz_dep_.CoverTab[96066]++
								return nil, err
//line /usr/local/go/src/log/syslog/syslog.go:315
		// _ = "end of CoverTab[96066]"
	} else {
//line /usr/local/go/src/log/syslog/syslog.go:316
		_go_fuzz_dep_.CoverTab[96067]++
//line /usr/local/go/src/log/syslog/syslog.go:316
		// _ = "end of CoverTab[96067]"
//line /usr/local/go/src/log/syslog/syslog.go:316
	}
//line /usr/local/go/src/log/syslog/syslog.go:316
	// _ = "end of CoverTab[96064]"
//line /usr/local/go/src/log/syslog/syslog.go:316
	_go_fuzz_dep_.CoverTab[96065]++
							return log.New(s, "", logFlag), nil
//line /usr/local/go/src/log/syslog/syslog.go:317
	// _ = "end of CoverTab[96065]"
}

//line /usr/local/go/src/log/syslog/syslog.go:318
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/log/syslog/syslog.go:318
var _ = _go_fuzz_dep_.CoverTab
