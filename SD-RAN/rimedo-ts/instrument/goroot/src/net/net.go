// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/net.go:5
/*
Package net provides a portable interface for network I/O, including
TCP/IP, UDP, domain name resolution, and Unix domain sockets.

Although the package provides access to low-level networking
primitives, most clients will need only the basic interface provided
by the Dial, Listen, and Accept functions and the associated
Conn and Listener interfaces. The crypto/tls package uses
the same interfaces and similar Dial and Listen functions.

The Dial function connects to a server:

	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	// ...

The Listen function creates servers:

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}

# Name Resolution

The method for resolving domain names, whether indirectly with functions like Dial
or directly with functions like LookupHost and LookupAddr, varies by operating system.

On Unix systems, the resolver has two options for resolving names.
It can use a pure Go resolver that sends DNS requests directly to the servers
listed in /etc/resolv.conf, or it can use a cgo-based resolver that calls C
library routines such as getaddrinfo and getnameinfo.

By default the pure Go resolver is used, because a blocked DNS request consumes
only a goroutine, while a blocked C call consumes an operating system thread.
When cgo is available, the cgo-based resolver is used instead under a variety of
conditions: on systems that do not let programs make direct DNS requests (OS X),
when the LOCALDOMAIN environment variable is present (even if empty),
when the RES_OPTIONS or HOSTALIASES environment variable is non-empty,
when the ASR_CONFIG environment variable is non-empty (OpenBSD only),
when /etc/resolv.conf or /etc/nsswitch.conf specify the use of features that the
Go resolver does not implement, and when the name being looked up ends in .local
or is an mDNS name.

The resolver decision can be overridden by setting the netdns value of the
GODEBUG environment variable (see package runtime) to go or cgo, as in:

	export GODEBUG=netdns=go    # force pure Go resolver
	export GODEBUG=netdns=cgo   # force native resolver (cgo, win32)

The decision can also be forced while building the Go source tree
by setting the netgo or netcgo build tag.

A numeric netdns setting, as in GODEBUG=netdns=1, causes the resolver
to print debugging information about its decisions.
To force a particular resolver while also printing debugging information,
join the two settings by a plus sign, as in GODEBUG=netdns=go+1.

On Plan 9, the resolver always accesses /net/cs and /net/dns.

On Windows, in Go 1.18.x and earlier, the resolver always used C
library functions, such as GetAddrInfo and DnsQuery.
*/
package net

//line /usr/local/go/src/net/net.go:79
import (
//line /usr/local/go/src/net/net.go:79
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/net.go:79
)
//line /usr/local/go/src/net/net.go:79
import (
//line /usr/local/go/src/net/net.go:79
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/net.go:79
)

import (
	"context"
	"errors"
	"internal/poll"
	"io"
	"os"
	"sync"
	"syscall"
	"time"
)

// netGo and netCgo contain the state of the build tags used
//line /usr/local/go/src/net/net.go:92
// to build this binary, and whether cgo is available.
//line /usr/local/go/src/net/net.go:92
// conf.go mirrors these into conf for easier testing.
//line /usr/local/go/src/net/net.go:95
var (
	netGo	bool	// set true in cgo_stub.go for build tag "netgo" (or no cgo)
	netCgo	bool	// set true in conf_netcgo.go for build tag "netcgo"
)

// Addr represents a network end point address.
//line /usr/local/go/src/net/net.go:100
//
//line /usr/local/go/src/net/net.go:100
// The two methods Network and String conventionally return strings
//line /usr/local/go/src/net/net.go:100
// that can be passed as the arguments to Dial, but the exact form
//line /usr/local/go/src/net/net.go:100
// and meaning of the strings is up to the implementation.
//line /usr/local/go/src/net/net.go:105
type Addr interface {
	Network() string	// name of the network (for example, "tcp", "udp")
	String() string		// string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
}

// Conn is a generic stream-oriented network connection.
//line /usr/local/go/src/net/net.go:110
//
//line /usr/local/go/src/net/net.go:110
// Multiple goroutines may invoke methods on a Conn simultaneously.
//line /usr/local/go/src/net/net.go:113
type Conn interface {
	// Read reads data from the connection.
	// Read can be made to time out and return an error after a fixed
	// time limit; see SetDeadline and SetReadDeadline.
	Read(b []byte) (n int, err error)

	// Write writes data to the connection.
	// Write can be made to time out and return an error after a fixed
	// time limit; see SetDeadline and SetWriteDeadline.
	Write(b []byte) (n int, err error)

	// Close closes the connection.
	// Any blocked Read or Write operations will be unblocked and return errors.
	Close() error

	// LocalAddr returns the local network address, if known.
	LocalAddr() Addr

	// RemoteAddr returns the remote network address, if known.
	RemoteAddr() Addr

	// SetDeadline sets the read and write deadlines associated
	// with the connection. It is equivalent to calling both
	// SetReadDeadline and SetWriteDeadline.
	//
	// A deadline is an absolute time after which I/O operations
	// fail instead of blocking. The deadline applies to all future
	// and pending I/O, not just the immediately following call to
	// Read or Write. After a deadline has been exceeded, the
	// connection can be refreshed by setting a deadline in the future.
	//
	// If the deadline is exceeded a call to Read or Write or to other
	// I/O methods will return an error that wraps os.ErrDeadlineExceeded.
	// This can be tested using errors.Is(err, os.ErrDeadlineExceeded).
	// The error's Timeout method will return true, but note that there
	// are other possible errors for which the Timeout method will
	// return true even if the deadline has not been exceeded.
	//
	// An idle timeout can be implemented by repeatedly extending
	// the deadline after successful Read or Write calls.
	//
	// A zero value for t means I/O operations will not time out.
	SetDeadline(t time.Time) error

	// SetReadDeadline sets the deadline for future Read calls
	// and any currently-blocked Read call.
	// A zero value for t means Read will not time out.
	SetReadDeadline(t time.Time) error

	// SetWriteDeadline sets the deadline for future Write calls
	// and any currently-blocked Write call.
	// Even if write times out, it may return n > 0, indicating that
	// some of the data was successfully written.
	// A zero value for t means Write will not time out.
	SetWriteDeadline(t time.Time) error
}

type conn struct {
	fd *netFD
}

func (c *conn) ok() bool {
//line /usr/local/go/src/net/net.go:174
	_go_fuzz_dep_.CoverTab[15503]++
//line /usr/local/go/src/net/net.go:174
	return c != nil && func() bool {
//line /usr/local/go/src/net/net.go:174
		_go_fuzz_dep_.CoverTab[15504]++
//line /usr/local/go/src/net/net.go:174
		return c.fd != nil
//line /usr/local/go/src/net/net.go:174
		// _ = "end of CoverTab[15504]"
//line /usr/local/go/src/net/net.go:174
	}()
//line /usr/local/go/src/net/net.go:174
	// _ = "end of CoverTab[15503]"
//line /usr/local/go/src/net/net.go:174
}

//line /usr/local/go/src/net/net.go:178
// Read implements the Conn Read method.
func (c *conn) Read(b []byte) (int, error) {
//line /usr/local/go/src/net/net.go:179
	_go_fuzz_dep_.CoverTab[15505]++
						if !c.ok() {
//line /usr/local/go/src/net/net.go:180
		_go_fuzz_dep_.CoverTab[15508]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/net.go:181
		// _ = "end of CoverTab[15508]"
	} else {
//line /usr/local/go/src/net/net.go:182
		_go_fuzz_dep_.CoverTab[15509]++
//line /usr/local/go/src/net/net.go:182
		// _ = "end of CoverTab[15509]"
//line /usr/local/go/src/net/net.go:182
	}
//line /usr/local/go/src/net/net.go:182
	// _ = "end of CoverTab[15505]"
//line /usr/local/go/src/net/net.go:182
	_go_fuzz_dep_.CoverTab[15506]++
						n, err := c.fd.Read(b)
						if err != nil && func() bool {
//line /usr/local/go/src/net/net.go:184
		_go_fuzz_dep_.CoverTab[15510]++
//line /usr/local/go/src/net/net.go:184
		return err != io.EOF
//line /usr/local/go/src/net/net.go:184
		// _ = "end of CoverTab[15510]"
//line /usr/local/go/src/net/net.go:184
	}() {
//line /usr/local/go/src/net/net.go:184
		_go_fuzz_dep_.CoverTab[15511]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/net.go:185
		// _ = "end of CoverTab[15511]"
	} else {
//line /usr/local/go/src/net/net.go:186
		_go_fuzz_dep_.CoverTab[15512]++
//line /usr/local/go/src/net/net.go:186
		// _ = "end of CoverTab[15512]"
//line /usr/local/go/src/net/net.go:186
	}
//line /usr/local/go/src/net/net.go:186
	// _ = "end of CoverTab[15506]"
//line /usr/local/go/src/net/net.go:186
	_go_fuzz_dep_.CoverTab[15507]++
						return n, err
//line /usr/local/go/src/net/net.go:187
	// _ = "end of CoverTab[15507]"
}

// Write implements the Conn Write method.
func (c *conn) Write(b []byte) (int, error) {
//line /usr/local/go/src/net/net.go:191
	_go_fuzz_dep_.CoverTab[15513]++
						if !c.ok() {
//line /usr/local/go/src/net/net.go:192
		_go_fuzz_dep_.CoverTab[15516]++
							return 0, syscall.EINVAL
//line /usr/local/go/src/net/net.go:193
		// _ = "end of CoverTab[15516]"
	} else {
//line /usr/local/go/src/net/net.go:194
		_go_fuzz_dep_.CoverTab[15517]++
//line /usr/local/go/src/net/net.go:194
		// _ = "end of CoverTab[15517]"
//line /usr/local/go/src/net/net.go:194
	}
//line /usr/local/go/src/net/net.go:194
	// _ = "end of CoverTab[15513]"
//line /usr/local/go/src/net/net.go:194
	_go_fuzz_dep_.CoverTab[15514]++
						n, err := c.fd.Write(b)
						if err != nil {
//line /usr/local/go/src/net/net.go:196
		_go_fuzz_dep_.CoverTab[15518]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/net.go:197
		// _ = "end of CoverTab[15518]"
	} else {
//line /usr/local/go/src/net/net.go:198
		_go_fuzz_dep_.CoverTab[15519]++
//line /usr/local/go/src/net/net.go:198
		// _ = "end of CoverTab[15519]"
//line /usr/local/go/src/net/net.go:198
	}
//line /usr/local/go/src/net/net.go:198
	// _ = "end of CoverTab[15514]"
//line /usr/local/go/src/net/net.go:198
	_go_fuzz_dep_.CoverTab[15515]++
						return n, err
//line /usr/local/go/src/net/net.go:199
	// _ = "end of CoverTab[15515]"
}

// Close closes the connection.
func (c *conn) Close() error {
//line /usr/local/go/src/net/net.go:203
	_go_fuzz_dep_.CoverTab[15520]++
						if !c.ok() {
//line /usr/local/go/src/net/net.go:204
		_go_fuzz_dep_.CoverTab[15523]++
							return syscall.EINVAL
//line /usr/local/go/src/net/net.go:205
		// _ = "end of CoverTab[15523]"
	} else {
//line /usr/local/go/src/net/net.go:206
		_go_fuzz_dep_.CoverTab[15524]++
//line /usr/local/go/src/net/net.go:206
		// _ = "end of CoverTab[15524]"
//line /usr/local/go/src/net/net.go:206
	}
//line /usr/local/go/src/net/net.go:206
	// _ = "end of CoverTab[15520]"
//line /usr/local/go/src/net/net.go:206
	_go_fuzz_dep_.CoverTab[15521]++
						err := c.fd.Close()
						if err != nil {
//line /usr/local/go/src/net/net.go:208
		_go_fuzz_dep_.CoverTab[15525]++
							err = &OpError{Op: "close", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/net.go:209
		// _ = "end of CoverTab[15525]"
	} else {
//line /usr/local/go/src/net/net.go:210
		_go_fuzz_dep_.CoverTab[15526]++
//line /usr/local/go/src/net/net.go:210
		// _ = "end of CoverTab[15526]"
//line /usr/local/go/src/net/net.go:210
	}
//line /usr/local/go/src/net/net.go:210
	// _ = "end of CoverTab[15521]"
//line /usr/local/go/src/net/net.go:210
	_go_fuzz_dep_.CoverTab[15522]++
						return err
//line /usr/local/go/src/net/net.go:211
	// _ = "end of CoverTab[15522]"
}

// LocalAddr returns the local network address.
//line /usr/local/go/src/net/net.go:214
// The Addr returned is shared by all invocations of LocalAddr, so
//line /usr/local/go/src/net/net.go:214
// do not modify it.
//line /usr/local/go/src/net/net.go:217
func (c *conn) LocalAddr() Addr {
//line /usr/local/go/src/net/net.go:217
	_go_fuzz_dep_.CoverTab[15527]++
						if !c.ok() {
//line /usr/local/go/src/net/net.go:218
		_go_fuzz_dep_.CoverTab[15529]++
							return nil
//line /usr/local/go/src/net/net.go:219
		// _ = "end of CoverTab[15529]"
	} else {
//line /usr/local/go/src/net/net.go:220
		_go_fuzz_dep_.CoverTab[15530]++
//line /usr/local/go/src/net/net.go:220
		// _ = "end of CoverTab[15530]"
//line /usr/local/go/src/net/net.go:220
	}
//line /usr/local/go/src/net/net.go:220
	// _ = "end of CoverTab[15527]"
//line /usr/local/go/src/net/net.go:220
	_go_fuzz_dep_.CoverTab[15528]++
						return c.fd.laddr
//line /usr/local/go/src/net/net.go:221
	// _ = "end of CoverTab[15528]"
}

// RemoteAddr returns the remote network address.
//line /usr/local/go/src/net/net.go:224
// The Addr returned is shared by all invocations of RemoteAddr, so
//line /usr/local/go/src/net/net.go:224
// do not modify it.
//line /usr/local/go/src/net/net.go:227
func (c *conn) RemoteAddr() Addr {
//line /usr/local/go/src/net/net.go:227
	_go_fuzz_dep_.CoverTab[15531]++
						if !c.ok() {
//line /usr/local/go/src/net/net.go:228
		_go_fuzz_dep_.CoverTab[15533]++
							return nil
//line /usr/local/go/src/net/net.go:229
		// _ = "end of CoverTab[15533]"
	} else {
//line /usr/local/go/src/net/net.go:230
		_go_fuzz_dep_.CoverTab[15534]++
//line /usr/local/go/src/net/net.go:230
		// _ = "end of CoverTab[15534]"
//line /usr/local/go/src/net/net.go:230
	}
//line /usr/local/go/src/net/net.go:230
	// _ = "end of CoverTab[15531]"
//line /usr/local/go/src/net/net.go:230
	_go_fuzz_dep_.CoverTab[15532]++
						return c.fd.raddr
//line /usr/local/go/src/net/net.go:231
	// _ = "end of CoverTab[15532]"
}

// SetDeadline implements the Conn SetDeadline method.
func (c *conn) SetDeadline(t time.Time) error {
//line /usr/local/go/src/net/net.go:235
	_go_fuzz_dep_.CoverTab[15535]++
						if !c.ok() {
//line /usr/local/go/src/net/net.go:236
		_go_fuzz_dep_.CoverTab[15538]++
							return syscall.EINVAL
//line /usr/local/go/src/net/net.go:237
		// _ = "end of CoverTab[15538]"
	} else {
//line /usr/local/go/src/net/net.go:238
		_go_fuzz_dep_.CoverTab[15539]++
//line /usr/local/go/src/net/net.go:238
		// _ = "end of CoverTab[15539]"
//line /usr/local/go/src/net/net.go:238
	}
//line /usr/local/go/src/net/net.go:238
	// _ = "end of CoverTab[15535]"
//line /usr/local/go/src/net/net.go:238
	_go_fuzz_dep_.CoverTab[15536]++
						if err := c.fd.SetDeadline(t); err != nil {
//line /usr/local/go/src/net/net.go:239
		_go_fuzz_dep_.CoverTab[15540]++
							return &OpError{Op: "set", Net: c.fd.net, Source: nil, Addr: c.fd.laddr, Err: err}
//line /usr/local/go/src/net/net.go:240
		// _ = "end of CoverTab[15540]"
	} else {
//line /usr/local/go/src/net/net.go:241
		_go_fuzz_dep_.CoverTab[15541]++
//line /usr/local/go/src/net/net.go:241
		// _ = "end of CoverTab[15541]"
//line /usr/local/go/src/net/net.go:241
	}
//line /usr/local/go/src/net/net.go:241
	// _ = "end of CoverTab[15536]"
//line /usr/local/go/src/net/net.go:241
	_go_fuzz_dep_.CoverTab[15537]++
						return nil
//line /usr/local/go/src/net/net.go:242
	// _ = "end of CoverTab[15537]"
}

// SetReadDeadline implements the Conn SetReadDeadline method.
func (c *conn) SetReadDeadline(t time.Time) error {
//line /usr/local/go/src/net/net.go:246
	_go_fuzz_dep_.CoverTab[15542]++
						if !c.ok() {
//line /usr/local/go/src/net/net.go:247
		_go_fuzz_dep_.CoverTab[15545]++
							return syscall.EINVAL
//line /usr/local/go/src/net/net.go:248
		// _ = "end of CoverTab[15545]"
	} else {
//line /usr/local/go/src/net/net.go:249
		_go_fuzz_dep_.CoverTab[15546]++
//line /usr/local/go/src/net/net.go:249
		// _ = "end of CoverTab[15546]"
//line /usr/local/go/src/net/net.go:249
	}
//line /usr/local/go/src/net/net.go:249
	// _ = "end of CoverTab[15542]"
//line /usr/local/go/src/net/net.go:249
	_go_fuzz_dep_.CoverTab[15543]++
						if err := c.fd.SetReadDeadline(t); err != nil {
//line /usr/local/go/src/net/net.go:250
		_go_fuzz_dep_.CoverTab[15547]++
							return &OpError{Op: "set", Net: c.fd.net, Source: nil, Addr: c.fd.laddr, Err: err}
//line /usr/local/go/src/net/net.go:251
		// _ = "end of CoverTab[15547]"
	} else {
//line /usr/local/go/src/net/net.go:252
		_go_fuzz_dep_.CoverTab[15548]++
//line /usr/local/go/src/net/net.go:252
		// _ = "end of CoverTab[15548]"
//line /usr/local/go/src/net/net.go:252
	}
//line /usr/local/go/src/net/net.go:252
	// _ = "end of CoverTab[15543]"
//line /usr/local/go/src/net/net.go:252
	_go_fuzz_dep_.CoverTab[15544]++
						return nil
//line /usr/local/go/src/net/net.go:253
	// _ = "end of CoverTab[15544]"
}

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (c *conn) SetWriteDeadline(t time.Time) error {
//line /usr/local/go/src/net/net.go:257
	_go_fuzz_dep_.CoverTab[15549]++
						if !c.ok() {
//line /usr/local/go/src/net/net.go:258
		_go_fuzz_dep_.CoverTab[15552]++
							return syscall.EINVAL
//line /usr/local/go/src/net/net.go:259
		// _ = "end of CoverTab[15552]"
	} else {
//line /usr/local/go/src/net/net.go:260
		_go_fuzz_dep_.CoverTab[15553]++
//line /usr/local/go/src/net/net.go:260
		// _ = "end of CoverTab[15553]"
//line /usr/local/go/src/net/net.go:260
	}
//line /usr/local/go/src/net/net.go:260
	// _ = "end of CoverTab[15549]"
//line /usr/local/go/src/net/net.go:260
	_go_fuzz_dep_.CoverTab[15550]++
						if err := c.fd.SetWriteDeadline(t); err != nil {
//line /usr/local/go/src/net/net.go:261
		_go_fuzz_dep_.CoverTab[15554]++
							return &OpError{Op: "set", Net: c.fd.net, Source: nil, Addr: c.fd.laddr, Err: err}
//line /usr/local/go/src/net/net.go:262
		// _ = "end of CoverTab[15554]"
	} else {
//line /usr/local/go/src/net/net.go:263
		_go_fuzz_dep_.CoverTab[15555]++
//line /usr/local/go/src/net/net.go:263
		// _ = "end of CoverTab[15555]"
//line /usr/local/go/src/net/net.go:263
	}
//line /usr/local/go/src/net/net.go:263
	// _ = "end of CoverTab[15550]"
//line /usr/local/go/src/net/net.go:263
	_go_fuzz_dep_.CoverTab[15551]++
						return nil
//line /usr/local/go/src/net/net.go:264
	// _ = "end of CoverTab[15551]"
}

// SetReadBuffer sets the size of the operating system's
//line /usr/local/go/src/net/net.go:267
// receive buffer associated with the connection.
//line /usr/local/go/src/net/net.go:269
func (c *conn) SetReadBuffer(bytes int) error {
//line /usr/local/go/src/net/net.go:269
	_go_fuzz_dep_.CoverTab[15556]++
						if !c.ok() {
//line /usr/local/go/src/net/net.go:270
		_go_fuzz_dep_.CoverTab[15559]++
							return syscall.EINVAL
//line /usr/local/go/src/net/net.go:271
		// _ = "end of CoverTab[15559]"
	} else {
//line /usr/local/go/src/net/net.go:272
		_go_fuzz_dep_.CoverTab[15560]++
//line /usr/local/go/src/net/net.go:272
		// _ = "end of CoverTab[15560]"
//line /usr/local/go/src/net/net.go:272
	}
//line /usr/local/go/src/net/net.go:272
	// _ = "end of CoverTab[15556]"
//line /usr/local/go/src/net/net.go:272
	_go_fuzz_dep_.CoverTab[15557]++
						if err := setReadBuffer(c.fd, bytes); err != nil {
//line /usr/local/go/src/net/net.go:273
		_go_fuzz_dep_.CoverTab[15561]++
							return &OpError{Op: "set", Net: c.fd.net, Source: nil, Addr: c.fd.laddr, Err: err}
//line /usr/local/go/src/net/net.go:274
		// _ = "end of CoverTab[15561]"
	} else {
//line /usr/local/go/src/net/net.go:275
		_go_fuzz_dep_.CoverTab[15562]++
//line /usr/local/go/src/net/net.go:275
		// _ = "end of CoverTab[15562]"
//line /usr/local/go/src/net/net.go:275
	}
//line /usr/local/go/src/net/net.go:275
	// _ = "end of CoverTab[15557]"
//line /usr/local/go/src/net/net.go:275
	_go_fuzz_dep_.CoverTab[15558]++
						return nil
//line /usr/local/go/src/net/net.go:276
	// _ = "end of CoverTab[15558]"
}

// SetWriteBuffer sets the size of the operating system's
//line /usr/local/go/src/net/net.go:279
// transmit buffer associated with the connection.
//line /usr/local/go/src/net/net.go:281
func (c *conn) SetWriteBuffer(bytes int) error {
//line /usr/local/go/src/net/net.go:281
	_go_fuzz_dep_.CoverTab[15563]++
						if !c.ok() {
//line /usr/local/go/src/net/net.go:282
		_go_fuzz_dep_.CoverTab[15566]++
							return syscall.EINVAL
//line /usr/local/go/src/net/net.go:283
		// _ = "end of CoverTab[15566]"
	} else {
//line /usr/local/go/src/net/net.go:284
		_go_fuzz_dep_.CoverTab[15567]++
//line /usr/local/go/src/net/net.go:284
		// _ = "end of CoverTab[15567]"
//line /usr/local/go/src/net/net.go:284
	}
//line /usr/local/go/src/net/net.go:284
	// _ = "end of CoverTab[15563]"
//line /usr/local/go/src/net/net.go:284
	_go_fuzz_dep_.CoverTab[15564]++
						if err := setWriteBuffer(c.fd, bytes); err != nil {
//line /usr/local/go/src/net/net.go:285
		_go_fuzz_dep_.CoverTab[15568]++
							return &OpError{Op: "set", Net: c.fd.net, Source: nil, Addr: c.fd.laddr, Err: err}
//line /usr/local/go/src/net/net.go:286
		// _ = "end of CoverTab[15568]"
	} else {
//line /usr/local/go/src/net/net.go:287
		_go_fuzz_dep_.CoverTab[15569]++
//line /usr/local/go/src/net/net.go:287
		// _ = "end of CoverTab[15569]"
//line /usr/local/go/src/net/net.go:287
	}
//line /usr/local/go/src/net/net.go:287
	// _ = "end of CoverTab[15564]"
//line /usr/local/go/src/net/net.go:287
	_go_fuzz_dep_.CoverTab[15565]++
						return nil
//line /usr/local/go/src/net/net.go:288
	// _ = "end of CoverTab[15565]"
}

// File returns a copy of the underlying os.File.
//line /usr/local/go/src/net/net.go:291
// It is the caller's responsibility to close f when finished.
//line /usr/local/go/src/net/net.go:291
// Closing c does not affect f, and closing f does not affect c.
//line /usr/local/go/src/net/net.go:291
//
//line /usr/local/go/src/net/net.go:291
// The returned os.File's file descriptor is different from the connection's.
//line /usr/local/go/src/net/net.go:291
// Attempting to change properties of the original using this duplicate
//line /usr/local/go/src/net/net.go:291
// may or may not have the desired effect.
//line /usr/local/go/src/net/net.go:298
func (c *conn) File() (f *os.File, err error) {
//line /usr/local/go/src/net/net.go:298
	_go_fuzz_dep_.CoverTab[15570]++
						f, err = c.fd.dup()
						if err != nil {
//line /usr/local/go/src/net/net.go:300
		_go_fuzz_dep_.CoverTab[15572]++
							err = &OpError{Op: "file", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /usr/local/go/src/net/net.go:301
		// _ = "end of CoverTab[15572]"
	} else {
//line /usr/local/go/src/net/net.go:302
		_go_fuzz_dep_.CoverTab[15573]++
//line /usr/local/go/src/net/net.go:302
		// _ = "end of CoverTab[15573]"
//line /usr/local/go/src/net/net.go:302
	}
//line /usr/local/go/src/net/net.go:302
	// _ = "end of CoverTab[15570]"
//line /usr/local/go/src/net/net.go:302
	_go_fuzz_dep_.CoverTab[15571]++
						return
//line /usr/local/go/src/net/net.go:303
	// _ = "end of CoverTab[15571]"
}

// PacketConn is a generic packet-oriented network connection.
//line /usr/local/go/src/net/net.go:306
//
//line /usr/local/go/src/net/net.go:306
// Multiple goroutines may invoke methods on a PacketConn simultaneously.
//line /usr/local/go/src/net/net.go:309
type PacketConn interface {
	// ReadFrom reads a packet from the connection,
	// copying the payload into p. It returns the number of
	// bytes copied into p and the return address that
	// was on the packet.
	// It returns the number of bytes read (0 <= n <= len(p))
	// and any error encountered. Callers should always process
	// the n > 0 bytes returned before considering the error err.
	// ReadFrom can be made to time out and return an error after a
	// fixed time limit; see SetDeadline and SetReadDeadline.
	ReadFrom(p []byte) (n int, addr Addr, err error)

	// WriteTo writes a packet with payload p to addr.
	// WriteTo can be made to time out and return an Error after a
	// fixed time limit; see SetDeadline and SetWriteDeadline.
	// On packet-oriented connections, write timeouts are rare.
	WriteTo(p []byte, addr Addr) (n int, err error)

	// Close closes the connection.
	// Any blocked ReadFrom or WriteTo operations will be unblocked and return errors.
	Close() error

	// LocalAddr returns the local network address, if known.
	LocalAddr() Addr

	// SetDeadline sets the read and write deadlines associated
	// with the connection. It is equivalent to calling both
	// SetReadDeadline and SetWriteDeadline.
	//
	// A deadline is an absolute time after which I/O operations
	// fail instead of blocking. The deadline applies to all future
	// and pending I/O, not just the immediately following call to
	// Read or Write. After a deadline has been exceeded, the
	// connection can be refreshed by setting a deadline in the future.
	//
	// If the deadline is exceeded a call to Read or Write or to other
	// I/O methods will return an error that wraps os.ErrDeadlineExceeded.
	// This can be tested using errors.Is(err, os.ErrDeadlineExceeded).
	// The error's Timeout method will return true, but note that there
	// are other possible errors for which the Timeout method will
	// return true even if the deadline has not been exceeded.
	//
	// An idle timeout can be implemented by repeatedly extending
	// the deadline after successful ReadFrom or WriteTo calls.
	//
	// A zero value for t means I/O operations will not time out.
	SetDeadline(t time.Time) error

	// SetReadDeadline sets the deadline for future ReadFrom calls
	// and any currently-blocked ReadFrom call.
	// A zero value for t means ReadFrom will not time out.
	SetReadDeadline(t time.Time) error

	// SetWriteDeadline sets the deadline for future WriteTo calls
	// and any currently-blocked WriteTo call.
	// Even if write times out, it may return n > 0, indicating that
	// some of the data was successfully written.
	// A zero value for t means WriteTo will not time out.
	SetWriteDeadline(t time.Time) error
}

var listenerBacklogCache struct {
	sync.Once
	val	int
}

// listenerBacklog is a caching wrapper around maxListenerBacklog.
func listenerBacklog() int {
//line /usr/local/go/src/net/net.go:376
	_go_fuzz_dep_.CoverTab[15574]++
						listenerBacklogCache.Do(func() {
//line /usr/local/go/src/net/net.go:377
		_go_fuzz_dep_.CoverTab[15576]++
//line /usr/local/go/src/net/net.go:377
		listenerBacklogCache.val = maxListenerBacklog()
//line /usr/local/go/src/net/net.go:377
		// _ = "end of CoverTab[15576]"
//line /usr/local/go/src/net/net.go:377
	})
//line /usr/local/go/src/net/net.go:377
	// _ = "end of CoverTab[15574]"
//line /usr/local/go/src/net/net.go:377
	_go_fuzz_dep_.CoverTab[15575]++
						return listenerBacklogCache.val
//line /usr/local/go/src/net/net.go:378
	// _ = "end of CoverTab[15575]"
}

// A Listener is a generic network listener for stream-oriented protocols.
//line /usr/local/go/src/net/net.go:381
//
//line /usr/local/go/src/net/net.go:381
// Multiple goroutines may invoke methods on a Listener simultaneously.
//line /usr/local/go/src/net/net.go:384
type Listener interface {
	// Accept waits for and returns the next connection to the listener.
	Accept() (Conn, error)

	// Close closes the listener.
	// Any blocked Accept operations will be unblocked and return errors.
	Close() error

	// Addr returns the listener's network address.
	Addr() Addr
}

// An Error represents a network error.
type Error interface {
	error
	Timeout() bool	// Is the error a timeout?

	// Deprecated: Temporary errors are not well-defined.
	// Most "temporary" errors are timeouts, and the few exceptions are surprising.
	// Do not use this method.
	Temporary() bool
}

// Various errors contained in OpError.
var (
	// For connection setup operations.
	errNoSuitableAddress	= errors.New("no suitable address found")

	// For connection setup and write operations.
	errMissingAddress	= errors.New("missing address")

	// For both read and write operations.
	errCanceled		= canceledError{}
	ErrWriteToConnected	= errors.New("use of WriteTo with pre-connected connection")
)

// canceledError lets us return the same error string we have always
//line /usr/local/go/src/net/net.go:420
// returned, while still being Is context.Canceled.
//line /usr/local/go/src/net/net.go:422
type canceledError struct{}

func (canceledError) Error() string {
//line /usr/local/go/src/net/net.go:424
	_go_fuzz_dep_.CoverTab[15577]++
//line /usr/local/go/src/net/net.go:424
	return "operation was canceled"
//line /usr/local/go/src/net/net.go:424
	// _ = "end of CoverTab[15577]"
//line /usr/local/go/src/net/net.go:424
}

func (canceledError) Is(err error) bool {
//line /usr/local/go/src/net/net.go:426
	_go_fuzz_dep_.CoverTab[15578]++
//line /usr/local/go/src/net/net.go:426
	return err == context.Canceled
//line /usr/local/go/src/net/net.go:426
	// _ = "end of CoverTab[15578]"
//line /usr/local/go/src/net/net.go:426
}

// mapErr maps from the context errors to the historical internal net
//line /usr/local/go/src/net/net.go:428
// error values.
//line /usr/local/go/src/net/net.go:430
func mapErr(err error) error {
//line /usr/local/go/src/net/net.go:430
	_go_fuzz_dep_.CoverTab[15579]++
						switch err {
	case context.Canceled:
//line /usr/local/go/src/net/net.go:432
		_go_fuzz_dep_.CoverTab[15580]++
							return errCanceled
//line /usr/local/go/src/net/net.go:433
		// _ = "end of CoverTab[15580]"
	case context.DeadlineExceeded:
//line /usr/local/go/src/net/net.go:434
		_go_fuzz_dep_.CoverTab[15581]++
							return errTimeout
//line /usr/local/go/src/net/net.go:435
		// _ = "end of CoverTab[15581]"
	default:
//line /usr/local/go/src/net/net.go:436
		_go_fuzz_dep_.CoverTab[15582]++
							return err
//line /usr/local/go/src/net/net.go:437
		// _ = "end of CoverTab[15582]"
	}
//line /usr/local/go/src/net/net.go:438
	// _ = "end of CoverTab[15579]"
}

// OpError is the error type usually returned by functions in the net
//line /usr/local/go/src/net/net.go:441
// package. It describes the operation, network type, and address of
//line /usr/local/go/src/net/net.go:441
// an error.
//line /usr/local/go/src/net/net.go:444
type OpError struct {
	// Op is the operation which caused the error, such as
	// "read" or "write".
	Op	string

	// Net is the network type on which this error occurred,
	// such as "tcp" or "udp6".
	Net	string

	// For operations involving a remote network connection, like
	// Dial, Read, or Write, Source is the corresponding local
	// network address.
	Source	Addr

	// Addr is the network address for which this error occurred.
	// For local operations, like Listen or SetDeadline, Addr is
	// the address of the local endpoint being manipulated.
	// For operations involving a remote network connection, like
	// Dial, Read, or Write, Addr is the remote address of that
	// connection.
	Addr	Addr

	// Err is the error that occurred during the operation.
	// The Error method panics if the error is nil.
	Err	error
}

func (e *OpError) Unwrap() error {
//line /usr/local/go/src/net/net.go:471
	_go_fuzz_dep_.CoverTab[15583]++
//line /usr/local/go/src/net/net.go:471
	return e.Err
//line /usr/local/go/src/net/net.go:471
	// _ = "end of CoverTab[15583]"
//line /usr/local/go/src/net/net.go:471
}

func (e *OpError) Error() string {
//line /usr/local/go/src/net/net.go:473
	_go_fuzz_dep_.CoverTab[15584]++
						if e == nil {
//line /usr/local/go/src/net/net.go:474
		_go_fuzz_dep_.CoverTab[15589]++
							return "<nil>"
//line /usr/local/go/src/net/net.go:475
		// _ = "end of CoverTab[15589]"
	} else {
//line /usr/local/go/src/net/net.go:476
		_go_fuzz_dep_.CoverTab[15590]++
//line /usr/local/go/src/net/net.go:476
		// _ = "end of CoverTab[15590]"
//line /usr/local/go/src/net/net.go:476
	}
//line /usr/local/go/src/net/net.go:476
	// _ = "end of CoverTab[15584]"
//line /usr/local/go/src/net/net.go:476
	_go_fuzz_dep_.CoverTab[15585]++
						s := e.Op
						if e.Net != "" {
//line /usr/local/go/src/net/net.go:478
		_go_fuzz_dep_.CoverTab[15591]++
							s += " " + e.Net
//line /usr/local/go/src/net/net.go:479
		// _ = "end of CoverTab[15591]"
	} else {
//line /usr/local/go/src/net/net.go:480
		_go_fuzz_dep_.CoverTab[15592]++
//line /usr/local/go/src/net/net.go:480
		// _ = "end of CoverTab[15592]"
//line /usr/local/go/src/net/net.go:480
	}
//line /usr/local/go/src/net/net.go:480
	// _ = "end of CoverTab[15585]"
//line /usr/local/go/src/net/net.go:480
	_go_fuzz_dep_.CoverTab[15586]++
						if e.Source != nil {
//line /usr/local/go/src/net/net.go:481
		_go_fuzz_dep_.CoverTab[15593]++
							s += " " + e.Source.String()
//line /usr/local/go/src/net/net.go:482
		// _ = "end of CoverTab[15593]"
	} else {
//line /usr/local/go/src/net/net.go:483
		_go_fuzz_dep_.CoverTab[15594]++
//line /usr/local/go/src/net/net.go:483
		// _ = "end of CoverTab[15594]"
//line /usr/local/go/src/net/net.go:483
	}
//line /usr/local/go/src/net/net.go:483
	// _ = "end of CoverTab[15586]"
//line /usr/local/go/src/net/net.go:483
	_go_fuzz_dep_.CoverTab[15587]++
						if e.Addr != nil {
//line /usr/local/go/src/net/net.go:484
		_go_fuzz_dep_.CoverTab[15595]++
							if e.Source != nil {
//line /usr/local/go/src/net/net.go:485
			_go_fuzz_dep_.CoverTab[15597]++
								s += "->"
//line /usr/local/go/src/net/net.go:486
			// _ = "end of CoverTab[15597]"
		} else {
//line /usr/local/go/src/net/net.go:487
			_go_fuzz_dep_.CoverTab[15598]++
								s += " "
//line /usr/local/go/src/net/net.go:488
			// _ = "end of CoverTab[15598]"
		}
//line /usr/local/go/src/net/net.go:489
		// _ = "end of CoverTab[15595]"
//line /usr/local/go/src/net/net.go:489
		_go_fuzz_dep_.CoverTab[15596]++
							s += e.Addr.String()
//line /usr/local/go/src/net/net.go:490
		// _ = "end of CoverTab[15596]"
	} else {
//line /usr/local/go/src/net/net.go:491
		_go_fuzz_dep_.CoverTab[15599]++
//line /usr/local/go/src/net/net.go:491
		// _ = "end of CoverTab[15599]"
//line /usr/local/go/src/net/net.go:491
	}
//line /usr/local/go/src/net/net.go:491
	// _ = "end of CoverTab[15587]"
//line /usr/local/go/src/net/net.go:491
	_go_fuzz_dep_.CoverTab[15588]++
						s += ": " + e.Err.Error()
						return s
//line /usr/local/go/src/net/net.go:493
	// _ = "end of CoverTab[15588]"
}

var (
	// aLongTimeAgo is a non-zero time, far in the past, used for
	// immediate cancellation of dials.
	aLongTimeAgo	= time.Unix(1, 0)

	// noDeadline and noCancel are just zero values for
	// readability with functions taking too many parameters.
	noDeadline	= time.Time{}
	noCancel	= (chan struct{})(nil)
)

type timeout interface {
	Timeout() bool
}

func (e *OpError) Timeout() bool {
//line /usr/local/go/src/net/net.go:511
	_go_fuzz_dep_.CoverTab[15600]++
						if ne, ok := e.Err.(*os.SyscallError); ok {
//line /usr/local/go/src/net/net.go:512
		_go_fuzz_dep_.CoverTab[15602]++
							t, ok := ne.Err.(timeout)
							return ok && func() bool {
//line /usr/local/go/src/net/net.go:514
			_go_fuzz_dep_.CoverTab[15603]++
//line /usr/local/go/src/net/net.go:514
			return t.Timeout()
//line /usr/local/go/src/net/net.go:514
			// _ = "end of CoverTab[15603]"
//line /usr/local/go/src/net/net.go:514
		}()
//line /usr/local/go/src/net/net.go:514
		// _ = "end of CoverTab[15602]"
	} else {
//line /usr/local/go/src/net/net.go:515
		_go_fuzz_dep_.CoverTab[15604]++
//line /usr/local/go/src/net/net.go:515
		// _ = "end of CoverTab[15604]"
//line /usr/local/go/src/net/net.go:515
	}
//line /usr/local/go/src/net/net.go:515
	// _ = "end of CoverTab[15600]"
//line /usr/local/go/src/net/net.go:515
	_go_fuzz_dep_.CoverTab[15601]++
						t, ok := e.Err.(timeout)
						return ok && func() bool {
//line /usr/local/go/src/net/net.go:517
		_go_fuzz_dep_.CoverTab[15605]++
//line /usr/local/go/src/net/net.go:517
		return t.Timeout()
//line /usr/local/go/src/net/net.go:517
		// _ = "end of CoverTab[15605]"
//line /usr/local/go/src/net/net.go:517
	}()
//line /usr/local/go/src/net/net.go:517
	// _ = "end of CoverTab[15601]"
}

type temporary interface {
	Temporary() bool
}

func (e *OpError) Temporary() bool {
//line /usr/local/go/src/net/net.go:524
	_go_fuzz_dep_.CoverTab[15606]++

//line /usr/local/go/src/net/net.go:527
	if e.Op == "accept" && func() bool {
//line /usr/local/go/src/net/net.go:527
		_go_fuzz_dep_.CoverTab[15609]++
//line /usr/local/go/src/net/net.go:527
		return isConnError(e.Err)
//line /usr/local/go/src/net/net.go:527
		// _ = "end of CoverTab[15609]"
//line /usr/local/go/src/net/net.go:527
	}() {
//line /usr/local/go/src/net/net.go:527
		_go_fuzz_dep_.CoverTab[15610]++
							return true
//line /usr/local/go/src/net/net.go:528
		// _ = "end of CoverTab[15610]"
	} else {
//line /usr/local/go/src/net/net.go:529
		_go_fuzz_dep_.CoverTab[15611]++
//line /usr/local/go/src/net/net.go:529
		// _ = "end of CoverTab[15611]"
//line /usr/local/go/src/net/net.go:529
	}
//line /usr/local/go/src/net/net.go:529
	// _ = "end of CoverTab[15606]"
//line /usr/local/go/src/net/net.go:529
	_go_fuzz_dep_.CoverTab[15607]++

						if ne, ok := e.Err.(*os.SyscallError); ok {
//line /usr/local/go/src/net/net.go:531
		_go_fuzz_dep_.CoverTab[15612]++
							t, ok := ne.Err.(temporary)
							return ok && func() bool {
//line /usr/local/go/src/net/net.go:533
			_go_fuzz_dep_.CoverTab[15613]++
//line /usr/local/go/src/net/net.go:533
			return t.Temporary()
//line /usr/local/go/src/net/net.go:533
			// _ = "end of CoverTab[15613]"
//line /usr/local/go/src/net/net.go:533
		}()
//line /usr/local/go/src/net/net.go:533
		// _ = "end of CoverTab[15612]"
	} else {
//line /usr/local/go/src/net/net.go:534
		_go_fuzz_dep_.CoverTab[15614]++
//line /usr/local/go/src/net/net.go:534
		// _ = "end of CoverTab[15614]"
//line /usr/local/go/src/net/net.go:534
	}
//line /usr/local/go/src/net/net.go:534
	// _ = "end of CoverTab[15607]"
//line /usr/local/go/src/net/net.go:534
	_go_fuzz_dep_.CoverTab[15608]++
						t, ok := e.Err.(temporary)
						return ok && func() bool {
//line /usr/local/go/src/net/net.go:536
		_go_fuzz_dep_.CoverTab[15615]++
//line /usr/local/go/src/net/net.go:536
		return t.Temporary()
//line /usr/local/go/src/net/net.go:536
		// _ = "end of CoverTab[15615]"
//line /usr/local/go/src/net/net.go:536
	}()
//line /usr/local/go/src/net/net.go:536
	// _ = "end of CoverTab[15608]"
}

// A ParseError is the error type of literal network address parsers.
type ParseError struct {
	// Type is the type of string that was expected, such as
	// "IP address", "CIDR address".
	Type	string

	// Text is the malformed text string.
	Text	string
}

func (e *ParseError) Error() string {
//line /usr/local/go/src/net/net.go:549
	_go_fuzz_dep_.CoverTab[15616]++
//line /usr/local/go/src/net/net.go:549
	return "invalid " + e.Type + ": " + e.Text
//line /usr/local/go/src/net/net.go:549
	// _ = "end of CoverTab[15616]"
//line /usr/local/go/src/net/net.go:549
}

func (e *ParseError) Timeout() bool {
//line /usr/local/go/src/net/net.go:551
	_go_fuzz_dep_.CoverTab[15617]++
//line /usr/local/go/src/net/net.go:551
	return false
//line /usr/local/go/src/net/net.go:551
	// _ = "end of CoverTab[15617]"
//line /usr/local/go/src/net/net.go:551
}
func (e *ParseError) Temporary() bool {
//line /usr/local/go/src/net/net.go:552
	_go_fuzz_dep_.CoverTab[15618]++
//line /usr/local/go/src/net/net.go:552
	return false
//line /usr/local/go/src/net/net.go:552
	// _ = "end of CoverTab[15618]"
//line /usr/local/go/src/net/net.go:552
}

type AddrError struct {
	Err	string
	Addr	string
}

func (e *AddrError) Error() string {
//line /usr/local/go/src/net/net.go:559
	_go_fuzz_dep_.CoverTab[15619]++
						if e == nil {
//line /usr/local/go/src/net/net.go:560
		_go_fuzz_dep_.CoverTab[15622]++
							return "<nil>"
//line /usr/local/go/src/net/net.go:561
		// _ = "end of CoverTab[15622]"
	} else {
//line /usr/local/go/src/net/net.go:562
		_go_fuzz_dep_.CoverTab[15623]++
//line /usr/local/go/src/net/net.go:562
		// _ = "end of CoverTab[15623]"
//line /usr/local/go/src/net/net.go:562
	}
//line /usr/local/go/src/net/net.go:562
	// _ = "end of CoverTab[15619]"
//line /usr/local/go/src/net/net.go:562
	_go_fuzz_dep_.CoverTab[15620]++
						s := e.Err
						if e.Addr != "" {
//line /usr/local/go/src/net/net.go:564
		_go_fuzz_dep_.CoverTab[15624]++
							s = "address " + e.Addr + ": " + s
//line /usr/local/go/src/net/net.go:565
		// _ = "end of CoverTab[15624]"
	} else {
//line /usr/local/go/src/net/net.go:566
		_go_fuzz_dep_.CoverTab[15625]++
//line /usr/local/go/src/net/net.go:566
		// _ = "end of CoverTab[15625]"
//line /usr/local/go/src/net/net.go:566
	}
//line /usr/local/go/src/net/net.go:566
	// _ = "end of CoverTab[15620]"
//line /usr/local/go/src/net/net.go:566
	_go_fuzz_dep_.CoverTab[15621]++
						return s
//line /usr/local/go/src/net/net.go:567
	// _ = "end of CoverTab[15621]"
}

func (e *AddrError) Timeout() bool {
//line /usr/local/go/src/net/net.go:570
	_go_fuzz_dep_.CoverTab[15626]++
//line /usr/local/go/src/net/net.go:570
	return false
//line /usr/local/go/src/net/net.go:570
	// _ = "end of CoverTab[15626]"
//line /usr/local/go/src/net/net.go:570
}
func (e *AddrError) Temporary() bool {
//line /usr/local/go/src/net/net.go:571
	_go_fuzz_dep_.CoverTab[15627]++
//line /usr/local/go/src/net/net.go:571
	return false
//line /usr/local/go/src/net/net.go:571
	// _ = "end of CoverTab[15627]"
//line /usr/local/go/src/net/net.go:571
}

type UnknownNetworkError string

func (e UnknownNetworkError) Error() string {
//line /usr/local/go/src/net/net.go:575
	_go_fuzz_dep_.CoverTab[15628]++
//line /usr/local/go/src/net/net.go:575
	return "unknown network " + string(e)
//line /usr/local/go/src/net/net.go:575
	// _ = "end of CoverTab[15628]"
//line /usr/local/go/src/net/net.go:575
}
func (e UnknownNetworkError) Timeout() bool {
//line /usr/local/go/src/net/net.go:576
	_go_fuzz_dep_.CoverTab[15629]++
//line /usr/local/go/src/net/net.go:576
	return false
//line /usr/local/go/src/net/net.go:576
	// _ = "end of CoverTab[15629]"
//line /usr/local/go/src/net/net.go:576
}
func (e UnknownNetworkError) Temporary() bool {
//line /usr/local/go/src/net/net.go:577
	_go_fuzz_dep_.CoverTab[15630]++
//line /usr/local/go/src/net/net.go:577
	return false
//line /usr/local/go/src/net/net.go:577
	// _ = "end of CoverTab[15630]"
//line /usr/local/go/src/net/net.go:577
}

type InvalidAddrError string

func (e InvalidAddrError) Error() string {
//line /usr/local/go/src/net/net.go:581
	_go_fuzz_dep_.CoverTab[15631]++
//line /usr/local/go/src/net/net.go:581
	return string(e)
//line /usr/local/go/src/net/net.go:581
	// _ = "end of CoverTab[15631]"
//line /usr/local/go/src/net/net.go:581
}
func (e InvalidAddrError) Timeout() bool {
//line /usr/local/go/src/net/net.go:582
	_go_fuzz_dep_.CoverTab[15632]++
//line /usr/local/go/src/net/net.go:582
	return false
//line /usr/local/go/src/net/net.go:582
	// _ = "end of CoverTab[15632]"
//line /usr/local/go/src/net/net.go:582
}
func (e InvalidAddrError) Temporary() bool {
//line /usr/local/go/src/net/net.go:583
	_go_fuzz_dep_.CoverTab[15633]++
//line /usr/local/go/src/net/net.go:583
	return false
//line /usr/local/go/src/net/net.go:583
	// _ = "end of CoverTab[15633]"
//line /usr/local/go/src/net/net.go:583
}

// errTimeout exists to return the historical "i/o timeout" string
//line /usr/local/go/src/net/net.go:585
// for context.DeadlineExceeded. See mapErr.
//line /usr/local/go/src/net/net.go:585
// It is also used when Dialer.Deadline is exceeded.
//line /usr/local/go/src/net/net.go:585
// error.Is(errTimeout, context.DeadlineExceeded) returns true.
//line /usr/local/go/src/net/net.go:585
//
//line /usr/local/go/src/net/net.go:585
// TODO(iant): We could consider changing this to os.ErrDeadlineExceeded
//line /usr/local/go/src/net/net.go:585
// in the future, if we make
//line /usr/local/go/src/net/net.go:585
//
//line /usr/local/go/src/net/net.go:585
//	errors.Is(os.ErrDeadlineExceeded, context.DeadlineExceeded)
//line /usr/local/go/src/net/net.go:585
//
//line /usr/local/go/src/net/net.go:585
// return true.
//line /usr/local/go/src/net/net.go:596
var errTimeout error = &timeoutError{}

type timeoutError struct{}

func (e *timeoutError) Error() string {
//line /usr/local/go/src/net/net.go:600
	_go_fuzz_dep_.CoverTab[15634]++
//line /usr/local/go/src/net/net.go:600
	return "i/o timeout"
//line /usr/local/go/src/net/net.go:600
	// _ = "end of CoverTab[15634]"
//line /usr/local/go/src/net/net.go:600
}
func (e *timeoutError) Timeout() bool {
//line /usr/local/go/src/net/net.go:601
	_go_fuzz_dep_.CoverTab[15635]++
//line /usr/local/go/src/net/net.go:601
	return true
//line /usr/local/go/src/net/net.go:601
	// _ = "end of CoverTab[15635]"
//line /usr/local/go/src/net/net.go:601
}
func (e *timeoutError) Temporary() bool {
//line /usr/local/go/src/net/net.go:602
	_go_fuzz_dep_.CoverTab[15636]++
//line /usr/local/go/src/net/net.go:602
	return true
//line /usr/local/go/src/net/net.go:602
	// _ = "end of CoverTab[15636]"
//line /usr/local/go/src/net/net.go:602
}

func (e *timeoutError) Is(err error) bool {
//line /usr/local/go/src/net/net.go:604
	_go_fuzz_dep_.CoverTab[15637]++
						return err == context.DeadlineExceeded
//line /usr/local/go/src/net/net.go:605
	// _ = "end of CoverTab[15637]"
}

// DNSConfigError represents an error reading the machine's DNS configuration.
//line /usr/local/go/src/net/net.go:608
// (No longer used; kept for compatibility.)
//line /usr/local/go/src/net/net.go:610
type DNSConfigError struct {
	Err error
}

func (e *DNSConfigError) Unwrap() error {
//line /usr/local/go/src/net/net.go:614
	_go_fuzz_dep_.CoverTab[15638]++
//line /usr/local/go/src/net/net.go:614
	return e.Err
//line /usr/local/go/src/net/net.go:614
	// _ = "end of CoverTab[15638]"
//line /usr/local/go/src/net/net.go:614
}
func (e *DNSConfigError) Error() string {
//line /usr/local/go/src/net/net.go:615
	_go_fuzz_dep_.CoverTab[15639]++
//line /usr/local/go/src/net/net.go:615
	return "error reading DNS config: " + e.Err.Error()
//line /usr/local/go/src/net/net.go:615
	// _ = "end of CoverTab[15639]"
//line /usr/local/go/src/net/net.go:615
}
func (e *DNSConfigError) Timeout() bool {
//line /usr/local/go/src/net/net.go:616
	_go_fuzz_dep_.CoverTab[15640]++
//line /usr/local/go/src/net/net.go:616
	return false
//line /usr/local/go/src/net/net.go:616
	// _ = "end of CoverTab[15640]"
//line /usr/local/go/src/net/net.go:616
}
func (e *DNSConfigError) Temporary() bool {
//line /usr/local/go/src/net/net.go:617
	_go_fuzz_dep_.CoverTab[15641]++
//line /usr/local/go/src/net/net.go:617
	return false
//line /usr/local/go/src/net/net.go:617
	// _ = "end of CoverTab[15641]"
//line /usr/local/go/src/net/net.go:617
}

// Various errors contained in DNSError.
var (
	errNoSuchHost = errors.New("no such host")
)

// DNSError represents a DNS lookup error.
type DNSError struct {
	Err		string	// description of the error
	Name		string	// name looked for
	Server		string	// server used
	IsTimeout	bool	// if true, timed out; not all timeouts set this
	IsTemporary	bool	// if true, error is temporary; not all errors set this
	IsNotFound	bool	// if true, host could not be found
}

func (e *DNSError) Error() string {
//line /usr/local/go/src/net/net.go:634
	_go_fuzz_dep_.CoverTab[15642]++
						if e == nil {
//line /usr/local/go/src/net/net.go:635
		_go_fuzz_dep_.CoverTab[15645]++
							return "<nil>"
//line /usr/local/go/src/net/net.go:636
		// _ = "end of CoverTab[15645]"
	} else {
//line /usr/local/go/src/net/net.go:637
		_go_fuzz_dep_.CoverTab[15646]++
//line /usr/local/go/src/net/net.go:637
		// _ = "end of CoverTab[15646]"
//line /usr/local/go/src/net/net.go:637
	}
//line /usr/local/go/src/net/net.go:637
	// _ = "end of CoverTab[15642]"
//line /usr/local/go/src/net/net.go:637
	_go_fuzz_dep_.CoverTab[15643]++
						s := "lookup " + e.Name
						if e.Server != "" {
//line /usr/local/go/src/net/net.go:639
		_go_fuzz_dep_.CoverTab[15647]++
							s += " on " + e.Server
//line /usr/local/go/src/net/net.go:640
		// _ = "end of CoverTab[15647]"
	} else {
//line /usr/local/go/src/net/net.go:641
		_go_fuzz_dep_.CoverTab[15648]++
//line /usr/local/go/src/net/net.go:641
		// _ = "end of CoverTab[15648]"
//line /usr/local/go/src/net/net.go:641
	}
//line /usr/local/go/src/net/net.go:641
	// _ = "end of CoverTab[15643]"
//line /usr/local/go/src/net/net.go:641
	_go_fuzz_dep_.CoverTab[15644]++
						s += ": " + e.Err
						return s
//line /usr/local/go/src/net/net.go:643
	// _ = "end of CoverTab[15644]"
}

// Timeout reports whether the DNS lookup is known to have timed out.
//line /usr/local/go/src/net/net.go:646
// This is not always known; a DNS lookup may fail due to a timeout
//line /usr/local/go/src/net/net.go:646
// and return a DNSError for which Timeout returns false.
//line /usr/local/go/src/net/net.go:649
func (e *DNSError) Timeout() bool {
//line /usr/local/go/src/net/net.go:649
	_go_fuzz_dep_.CoverTab[15649]++
//line /usr/local/go/src/net/net.go:649
	return e.IsTimeout
//line /usr/local/go/src/net/net.go:649
	// _ = "end of CoverTab[15649]"
//line /usr/local/go/src/net/net.go:649
}

// Temporary reports whether the DNS error is known to be temporary.
//line /usr/local/go/src/net/net.go:651
// This is not always known; a DNS lookup may fail due to a temporary
//line /usr/local/go/src/net/net.go:651
// error and return a DNSError for which Temporary returns false.
//line /usr/local/go/src/net/net.go:654
func (e *DNSError) Temporary() bool {
//line /usr/local/go/src/net/net.go:654
	_go_fuzz_dep_.CoverTab[15650]++
//line /usr/local/go/src/net/net.go:654
	return e.IsTimeout || func() bool {
//line /usr/local/go/src/net/net.go:654
		_go_fuzz_dep_.CoverTab[15651]++
//line /usr/local/go/src/net/net.go:654
		return e.IsTemporary
//line /usr/local/go/src/net/net.go:654
		// _ = "end of CoverTab[15651]"
//line /usr/local/go/src/net/net.go:654
	}()
//line /usr/local/go/src/net/net.go:654
	// _ = "end of CoverTab[15650]"
//line /usr/local/go/src/net/net.go:654
}

// errClosed exists just so that the docs for ErrClosed don't mention
//line /usr/local/go/src/net/net.go:656
// the internal package poll.
//line /usr/local/go/src/net/net.go:658
var errClosed = poll.ErrNetClosing

// ErrClosed is the error returned by an I/O call on a network
//line /usr/local/go/src/net/net.go:660
// connection that has already been closed, or that is closed by
//line /usr/local/go/src/net/net.go:660
// another goroutine before the I/O is completed. This may be wrapped
//line /usr/local/go/src/net/net.go:660
// in another error, and should normally be tested using
//line /usr/local/go/src/net/net.go:660
// errors.Is(err, net.ErrClosed).
//line /usr/local/go/src/net/net.go:665
var ErrClosed error = errClosed

type writerOnly struct {
	io.Writer
}

// Fallback implementation of io.ReaderFrom's ReadFrom, when sendfile isn't
//line /usr/local/go/src/net/net.go:671
// applicable.
//line /usr/local/go/src/net/net.go:673
func genericReadFrom(w io.Writer, r io.Reader) (n int64, err error) {
//line /usr/local/go/src/net/net.go:673
	_go_fuzz_dep_.CoverTab[15652]++

						return io.Copy(writerOnly{w}, r)
//line /usr/local/go/src/net/net.go:675
	// _ = "end of CoverTab[15652]"
}

//line /usr/local/go/src/net/net.go:684
var threadLimit chan struct{}

var threadOnce sync.Once

func acquireThread() {
//line /usr/local/go/src/net/net.go:688
	_go_fuzz_dep_.CoverTab[15653]++
						threadOnce.Do(func() {
//line /usr/local/go/src/net/net.go:689
		_go_fuzz_dep_.CoverTab[15655]++
							threadLimit = make(chan struct{}, concurrentThreadsLimit())
//line /usr/local/go/src/net/net.go:690
		// _ = "end of CoverTab[15655]"
	})
//line /usr/local/go/src/net/net.go:691
	// _ = "end of CoverTab[15653]"
//line /usr/local/go/src/net/net.go:691
	_go_fuzz_dep_.CoverTab[15654]++
						threadLimit <- struct{}{}
//line /usr/local/go/src/net/net.go:692
	// _ = "end of CoverTab[15654]"
}

func releaseThread() {
//line /usr/local/go/src/net/net.go:695
	_go_fuzz_dep_.CoverTab[15656]++
						<-threadLimit
//line /usr/local/go/src/net/net.go:696
	// _ = "end of CoverTab[15656]"
}

// buffersWriter is the interface implemented by Conns that support a
//line /usr/local/go/src/net/net.go:699
// "writev"-like batch write optimization.
//line /usr/local/go/src/net/net.go:699
// writeBuffers should fully consume and write all chunks from the
//line /usr/local/go/src/net/net.go:699
// provided Buffers, else it should report a non-nil error.
//line /usr/local/go/src/net/net.go:703
type buffersWriter interface {
	writeBuffers(*Buffers) (int64, error)
}

// Buffers contains zero or more runs of bytes to write.
//line /usr/local/go/src/net/net.go:707
//
//line /usr/local/go/src/net/net.go:707
// On certain machines, for certain types of connections, this is
//line /usr/local/go/src/net/net.go:707
// optimized into an OS-specific batch write operation (such as
//line /usr/local/go/src/net/net.go:707
// "writev").
//line /usr/local/go/src/net/net.go:712
type Buffers [][]byte

var (
	_	io.WriterTo	= (*Buffers)(nil)
	_	io.Reader	= (*Buffers)(nil)
)

// WriteTo writes contents of the buffers to w.
//line /usr/local/go/src/net/net.go:719
//
//line /usr/local/go/src/net/net.go:719
// WriteTo implements io.WriterTo for Buffers.
//line /usr/local/go/src/net/net.go:719
//
//line /usr/local/go/src/net/net.go:719
// WriteTo modifies the slice v as well as v[i] for 0 <= i < len(v),
//line /usr/local/go/src/net/net.go:719
// but does not modify v[i][j] for any i, j.
//line /usr/local/go/src/net/net.go:725
func (v *Buffers) WriteTo(w io.Writer) (n int64, err error) {
//line /usr/local/go/src/net/net.go:725
	_go_fuzz_dep_.CoverTab[15657]++
						if wv, ok := w.(buffersWriter); ok {
//line /usr/local/go/src/net/net.go:726
		_go_fuzz_dep_.CoverTab[15660]++
							return wv.writeBuffers(v)
//line /usr/local/go/src/net/net.go:727
		// _ = "end of CoverTab[15660]"
	} else {
//line /usr/local/go/src/net/net.go:728
		_go_fuzz_dep_.CoverTab[15661]++
//line /usr/local/go/src/net/net.go:728
		// _ = "end of CoverTab[15661]"
//line /usr/local/go/src/net/net.go:728
	}
//line /usr/local/go/src/net/net.go:728
	// _ = "end of CoverTab[15657]"
//line /usr/local/go/src/net/net.go:728
	_go_fuzz_dep_.CoverTab[15658]++
						for _, b := range *v {
//line /usr/local/go/src/net/net.go:729
		_go_fuzz_dep_.CoverTab[15662]++
							nb, err := w.Write(b)
							n += int64(nb)
							if err != nil {
//line /usr/local/go/src/net/net.go:732
			_go_fuzz_dep_.CoverTab[15663]++
								v.consume(n)
								return n, err
//line /usr/local/go/src/net/net.go:734
			// _ = "end of CoverTab[15663]"
		} else {
//line /usr/local/go/src/net/net.go:735
			_go_fuzz_dep_.CoverTab[15664]++
//line /usr/local/go/src/net/net.go:735
			// _ = "end of CoverTab[15664]"
//line /usr/local/go/src/net/net.go:735
		}
//line /usr/local/go/src/net/net.go:735
		// _ = "end of CoverTab[15662]"
	}
//line /usr/local/go/src/net/net.go:736
	// _ = "end of CoverTab[15658]"
//line /usr/local/go/src/net/net.go:736
	_go_fuzz_dep_.CoverTab[15659]++
						v.consume(n)
						return n, nil
//line /usr/local/go/src/net/net.go:738
	// _ = "end of CoverTab[15659]"
}

// Read from the buffers.
//line /usr/local/go/src/net/net.go:741
//
//line /usr/local/go/src/net/net.go:741
// Read implements io.Reader for Buffers.
//line /usr/local/go/src/net/net.go:741
//
//line /usr/local/go/src/net/net.go:741
// Read modifies the slice v as well as v[i] for 0 <= i < len(v),
//line /usr/local/go/src/net/net.go:741
// but does not modify v[i][j] for any i, j.
//line /usr/local/go/src/net/net.go:747
func (v *Buffers) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/net.go:747
	_go_fuzz_dep_.CoverTab[15665]++
						for len(p) > 0 && func() bool {
//line /usr/local/go/src/net/net.go:748
		_go_fuzz_dep_.CoverTab[15668]++
//line /usr/local/go/src/net/net.go:748
		return len(*v) > 0
//line /usr/local/go/src/net/net.go:748
		// _ = "end of CoverTab[15668]"
//line /usr/local/go/src/net/net.go:748
	}() {
//line /usr/local/go/src/net/net.go:748
		_go_fuzz_dep_.CoverTab[15669]++
							n0 := copy(p, (*v)[0])
							v.consume(int64(n0))
							p = p[n0:]
							n += n0
//line /usr/local/go/src/net/net.go:752
		// _ = "end of CoverTab[15669]"
	}
//line /usr/local/go/src/net/net.go:753
	// _ = "end of CoverTab[15665]"
//line /usr/local/go/src/net/net.go:753
	_go_fuzz_dep_.CoverTab[15666]++
						if len(*v) == 0 {
//line /usr/local/go/src/net/net.go:754
		_go_fuzz_dep_.CoverTab[15670]++
							err = io.EOF
//line /usr/local/go/src/net/net.go:755
		// _ = "end of CoverTab[15670]"
	} else {
//line /usr/local/go/src/net/net.go:756
		_go_fuzz_dep_.CoverTab[15671]++
//line /usr/local/go/src/net/net.go:756
		// _ = "end of CoverTab[15671]"
//line /usr/local/go/src/net/net.go:756
	}
//line /usr/local/go/src/net/net.go:756
	// _ = "end of CoverTab[15666]"
//line /usr/local/go/src/net/net.go:756
	_go_fuzz_dep_.CoverTab[15667]++
						return
//line /usr/local/go/src/net/net.go:757
	// _ = "end of CoverTab[15667]"
}

func (v *Buffers) consume(n int64) {
//line /usr/local/go/src/net/net.go:760
	_go_fuzz_dep_.CoverTab[15672]++
						for len(*v) > 0 {
//line /usr/local/go/src/net/net.go:761
		_go_fuzz_dep_.CoverTab[15673]++
							ln0 := int64(len((*v)[0]))
							if ln0 > n {
//line /usr/local/go/src/net/net.go:763
			_go_fuzz_dep_.CoverTab[15675]++
								(*v)[0] = (*v)[0][n:]
								return
//line /usr/local/go/src/net/net.go:765
			// _ = "end of CoverTab[15675]"
		} else {
//line /usr/local/go/src/net/net.go:766
			_go_fuzz_dep_.CoverTab[15676]++
//line /usr/local/go/src/net/net.go:766
			// _ = "end of CoverTab[15676]"
//line /usr/local/go/src/net/net.go:766
		}
//line /usr/local/go/src/net/net.go:766
		// _ = "end of CoverTab[15673]"
//line /usr/local/go/src/net/net.go:766
		_go_fuzz_dep_.CoverTab[15674]++
							n -= ln0
							(*v)[0] = nil
							*v = (*v)[1:]
//line /usr/local/go/src/net/net.go:769
		// _ = "end of CoverTab[15674]"
	}
//line /usr/local/go/src/net/net.go:770
	// _ = "end of CoverTab[15672]"
}

//line /usr/local/go/src/net/net.go:771
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/net.go:771
var _ = _go_fuzz_dep_.CoverTab
