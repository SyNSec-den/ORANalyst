// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/net/net.go:5
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

On macOS, if Go code that uses the net package is built with
-buildmode=c-archive, linking the resulting archive into a C program
requires passing -lresolv when linking the C code.

On Plan 9, the resolver always accesses /net/cs and /net/dns.

On Windows, in Go 1.18.x and earlier, the resolver always used C
library functions, such as GetAddrInfo and DnsQuery.
*/
package net

//line /snap/go/10455/src/net/net.go:83
import (
//line /snap/go/10455/src/net/net.go:83
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/net.go:83
)
//line /snap/go/10455/src/net/net.go:83
import (
//line /snap/go/10455/src/net/net.go:83
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/net.go:83
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

// Addr represents a network end point address.
//line /snap/go/10455/src/net/net.go:96
//
//line /snap/go/10455/src/net/net.go:96
// The two methods Network and String conventionally return strings
//line /snap/go/10455/src/net/net.go:96
// that can be passed as the arguments to Dial, but the exact form
//line /snap/go/10455/src/net/net.go:96
// and meaning of the strings is up to the implementation.
//line /snap/go/10455/src/net/net.go:101
type Addr interface {
	Network() string	// name of the network (for example, "tcp", "udp")
	String() string		// string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
}

// Conn is a generic stream-oriented network connection.
//line /snap/go/10455/src/net/net.go:106
//
//line /snap/go/10455/src/net/net.go:106
// Multiple goroutines may invoke methods on a Conn simultaneously.
//line /snap/go/10455/src/net/net.go:109
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
//line /snap/go/10455/src/net/net.go:170
	_go_fuzz_dep_.CoverTab[7393]++
//line /snap/go/10455/src/net/net.go:170
	return c != nil && func() bool {
//line /snap/go/10455/src/net/net.go:170
		_go_fuzz_dep_.CoverTab[7394]++
//line /snap/go/10455/src/net/net.go:170
		return c.fd != nil
//line /snap/go/10455/src/net/net.go:170
		// _ = "end of CoverTab[7394]"
//line /snap/go/10455/src/net/net.go:170
	}()
//line /snap/go/10455/src/net/net.go:170
	// _ = "end of CoverTab[7393]"
//line /snap/go/10455/src/net/net.go:170
}

//line /snap/go/10455/src/net/net.go:174
// Read implements the Conn Read method.
func (c *conn) Read(b []byte) (int, error) {
//line /snap/go/10455/src/net/net.go:175
	_go_fuzz_dep_.CoverTab[7395]++
						if !c.ok() {
//line /snap/go/10455/src/net/net.go:176
		_go_fuzz_dep_.CoverTab[529240]++
//line /snap/go/10455/src/net/net.go:176
		_go_fuzz_dep_.CoverTab[7398]++
							return 0, syscall.EINVAL
//line /snap/go/10455/src/net/net.go:177
		// _ = "end of CoverTab[7398]"
	} else {
//line /snap/go/10455/src/net/net.go:178
		_go_fuzz_dep_.CoverTab[529241]++
//line /snap/go/10455/src/net/net.go:178
		_go_fuzz_dep_.CoverTab[7399]++
//line /snap/go/10455/src/net/net.go:178
		// _ = "end of CoverTab[7399]"
//line /snap/go/10455/src/net/net.go:178
	}
//line /snap/go/10455/src/net/net.go:178
	// _ = "end of CoverTab[7395]"
//line /snap/go/10455/src/net/net.go:178
	_go_fuzz_dep_.CoverTab[7396]++
						n, err := c.fd.Read(b)
						if err != nil && func() bool {
//line /snap/go/10455/src/net/net.go:180
		_go_fuzz_dep_.CoverTab[7400]++
//line /snap/go/10455/src/net/net.go:180
		return err != io.EOF
//line /snap/go/10455/src/net/net.go:180
		// _ = "end of CoverTab[7400]"
//line /snap/go/10455/src/net/net.go:180
	}() {
//line /snap/go/10455/src/net/net.go:180
		_go_fuzz_dep_.CoverTab[529242]++
//line /snap/go/10455/src/net/net.go:180
		_go_fuzz_dep_.CoverTab[7401]++
							err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/net.go:181
		// _ = "end of CoverTab[7401]"
	} else {
//line /snap/go/10455/src/net/net.go:182
		_go_fuzz_dep_.CoverTab[529243]++
//line /snap/go/10455/src/net/net.go:182
		_go_fuzz_dep_.CoverTab[7402]++
//line /snap/go/10455/src/net/net.go:182
		// _ = "end of CoverTab[7402]"
//line /snap/go/10455/src/net/net.go:182
	}
//line /snap/go/10455/src/net/net.go:182
	// _ = "end of CoverTab[7396]"
//line /snap/go/10455/src/net/net.go:182
	_go_fuzz_dep_.CoverTab[7397]++
						return n, err
//line /snap/go/10455/src/net/net.go:183
	// _ = "end of CoverTab[7397]"
}

// Write implements the Conn Write method.
func (c *conn) Write(b []byte) (int, error) {
//line /snap/go/10455/src/net/net.go:187
	_go_fuzz_dep_.CoverTab[7403]++
						if !c.ok() {
//line /snap/go/10455/src/net/net.go:188
		_go_fuzz_dep_.CoverTab[529244]++
//line /snap/go/10455/src/net/net.go:188
		_go_fuzz_dep_.CoverTab[7406]++
							return 0, syscall.EINVAL
//line /snap/go/10455/src/net/net.go:189
		// _ = "end of CoverTab[7406]"
	} else {
//line /snap/go/10455/src/net/net.go:190
		_go_fuzz_dep_.CoverTab[529245]++
//line /snap/go/10455/src/net/net.go:190
		_go_fuzz_dep_.CoverTab[7407]++
//line /snap/go/10455/src/net/net.go:190
		// _ = "end of CoverTab[7407]"
//line /snap/go/10455/src/net/net.go:190
	}
//line /snap/go/10455/src/net/net.go:190
	// _ = "end of CoverTab[7403]"
//line /snap/go/10455/src/net/net.go:190
	_go_fuzz_dep_.CoverTab[7404]++
						n, err := c.fd.Write(b)
						if err != nil {
//line /snap/go/10455/src/net/net.go:192
		_go_fuzz_dep_.CoverTab[529246]++
//line /snap/go/10455/src/net/net.go:192
		_go_fuzz_dep_.CoverTab[7408]++
							err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/net.go:193
		// _ = "end of CoverTab[7408]"
	} else {
//line /snap/go/10455/src/net/net.go:194
		_go_fuzz_dep_.CoverTab[529247]++
//line /snap/go/10455/src/net/net.go:194
		_go_fuzz_dep_.CoverTab[7409]++
//line /snap/go/10455/src/net/net.go:194
		// _ = "end of CoverTab[7409]"
//line /snap/go/10455/src/net/net.go:194
	}
//line /snap/go/10455/src/net/net.go:194
	// _ = "end of CoverTab[7404]"
//line /snap/go/10455/src/net/net.go:194
	_go_fuzz_dep_.CoverTab[7405]++
						return n, err
//line /snap/go/10455/src/net/net.go:195
	// _ = "end of CoverTab[7405]"
}

// Close closes the connection.
func (c *conn) Close() error {
//line /snap/go/10455/src/net/net.go:199
	_go_fuzz_dep_.CoverTab[7410]++
						if !c.ok() {
//line /snap/go/10455/src/net/net.go:200
		_go_fuzz_dep_.CoverTab[529248]++
//line /snap/go/10455/src/net/net.go:200
		_go_fuzz_dep_.CoverTab[7413]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/net.go:201
		// _ = "end of CoverTab[7413]"
	} else {
//line /snap/go/10455/src/net/net.go:202
		_go_fuzz_dep_.CoverTab[529249]++
//line /snap/go/10455/src/net/net.go:202
		_go_fuzz_dep_.CoverTab[7414]++
//line /snap/go/10455/src/net/net.go:202
		// _ = "end of CoverTab[7414]"
//line /snap/go/10455/src/net/net.go:202
	}
//line /snap/go/10455/src/net/net.go:202
	// _ = "end of CoverTab[7410]"
//line /snap/go/10455/src/net/net.go:202
	_go_fuzz_dep_.CoverTab[7411]++
						err := c.fd.Close()
						if err != nil {
//line /snap/go/10455/src/net/net.go:204
		_go_fuzz_dep_.CoverTab[529250]++
//line /snap/go/10455/src/net/net.go:204
		_go_fuzz_dep_.CoverTab[7415]++
							err = &OpError{Op: "close", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/net.go:205
		// _ = "end of CoverTab[7415]"
	} else {
//line /snap/go/10455/src/net/net.go:206
		_go_fuzz_dep_.CoverTab[529251]++
//line /snap/go/10455/src/net/net.go:206
		_go_fuzz_dep_.CoverTab[7416]++
//line /snap/go/10455/src/net/net.go:206
		// _ = "end of CoverTab[7416]"
//line /snap/go/10455/src/net/net.go:206
	}
//line /snap/go/10455/src/net/net.go:206
	// _ = "end of CoverTab[7411]"
//line /snap/go/10455/src/net/net.go:206
	_go_fuzz_dep_.CoverTab[7412]++
						return err
//line /snap/go/10455/src/net/net.go:207
	// _ = "end of CoverTab[7412]"
}

// LocalAddr returns the local network address.
//line /snap/go/10455/src/net/net.go:210
// The Addr returned is shared by all invocations of LocalAddr, so
//line /snap/go/10455/src/net/net.go:210
// do not modify it.
//line /snap/go/10455/src/net/net.go:213
func (c *conn) LocalAddr() Addr {
//line /snap/go/10455/src/net/net.go:213
	_go_fuzz_dep_.CoverTab[7417]++
						if !c.ok() {
//line /snap/go/10455/src/net/net.go:214
		_go_fuzz_dep_.CoverTab[529252]++
//line /snap/go/10455/src/net/net.go:214
		_go_fuzz_dep_.CoverTab[7419]++
							return nil
//line /snap/go/10455/src/net/net.go:215
		// _ = "end of CoverTab[7419]"
	} else {
//line /snap/go/10455/src/net/net.go:216
		_go_fuzz_dep_.CoverTab[529253]++
//line /snap/go/10455/src/net/net.go:216
		_go_fuzz_dep_.CoverTab[7420]++
//line /snap/go/10455/src/net/net.go:216
		// _ = "end of CoverTab[7420]"
//line /snap/go/10455/src/net/net.go:216
	}
//line /snap/go/10455/src/net/net.go:216
	// _ = "end of CoverTab[7417]"
//line /snap/go/10455/src/net/net.go:216
	_go_fuzz_dep_.CoverTab[7418]++
						return c.fd.laddr
//line /snap/go/10455/src/net/net.go:217
	// _ = "end of CoverTab[7418]"
}

// RemoteAddr returns the remote network address.
//line /snap/go/10455/src/net/net.go:220
// The Addr returned is shared by all invocations of RemoteAddr, so
//line /snap/go/10455/src/net/net.go:220
// do not modify it.
//line /snap/go/10455/src/net/net.go:223
func (c *conn) RemoteAddr() Addr {
//line /snap/go/10455/src/net/net.go:223
	_go_fuzz_dep_.CoverTab[7421]++
						if !c.ok() {
//line /snap/go/10455/src/net/net.go:224
		_go_fuzz_dep_.CoverTab[529254]++
//line /snap/go/10455/src/net/net.go:224
		_go_fuzz_dep_.CoverTab[7423]++
							return nil
//line /snap/go/10455/src/net/net.go:225
		// _ = "end of CoverTab[7423]"
	} else {
//line /snap/go/10455/src/net/net.go:226
		_go_fuzz_dep_.CoverTab[529255]++
//line /snap/go/10455/src/net/net.go:226
		_go_fuzz_dep_.CoverTab[7424]++
//line /snap/go/10455/src/net/net.go:226
		// _ = "end of CoverTab[7424]"
//line /snap/go/10455/src/net/net.go:226
	}
//line /snap/go/10455/src/net/net.go:226
	// _ = "end of CoverTab[7421]"
//line /snap/go/10455/src/net/net.go:226
	_go_fuzz_dep_.CoverTab[7422]++
						return c.fd.raddr
//line /snap/go/10455/src/net/net.go:227
	// _ = "end of CoverTab[7422]"
}

// SetDeadline implements the Conn SetDeadline method.
func (c *conn) SetDeadline(t time.Time) error {
//line /snap/go/10455/src/net/net.go:231
	_go_fuzz_dep_.CoverTab[7425]++
						if !c.ok() {
//line /snap/go/10455/src/net/net.go:232
		_go_fuzz_dep_.CoverTab[529256]++
//line /snap/go/10455/src/net/net.go:232
		_go_fuzz_dep_.CoverTab[7428]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/net.go:233
		// _ = "end of CoverTab[7428]"
	} else {
//line /snap/go/10455/src/net/net.go:234
		_go_fuzz_dep_.CoverTab[529257]++
//line /snap/go/10455/src/net/net.go:234
		_go_fuzz_dep_.CoverTab[7429]++
//line /snap/go/10455/src/net/net.go:234
		// _ = "end of CoverTab[7429]"
//line /snap/go/10455/src/net/net.go:234
	}
//line /snap/go/10455/src/net/net.go:234
	// _ = "end of CoverTab[7425]"
//line /snap/go/10455/src/net/net.go:234
	_go_fuzz_dep_.CoverTab[7426]++
						if err := c.fd.SetDeadline(t); err != nil {
//line /snap/go/10455/src/net/net.go:235
		_go_fuzz_dep_.CoverTab[529258]++
//line /snap/go/10455/src/net/net.go:235
		_go_fuzz_dep_.CoverTab[7430]++
							return &OpError{Op: "set", Net: c.fd.net, Source: nil, Addr: c.fd.laddr, Err: err}
//line /snap/go/10455/src/net/net.go:236
		// _ = "end of CoverTab[7430]"
	} else {
//line /snap/go/10455/src/net/net.go:237
		_go_fuzz_dep_.CoverTab[529259]++
//line /snap/go/10455/src/net/net.go:237
		_go_fuzz_dep_.CoverTab[7431]++
//line /snap/go/10455/src/net/net.go:237
		// _ = "end of CoverTab[7431]"
//line /snap/go/10455/src/net/net.go:237
	}
//line /snap/go/10455/src/net/net.go:237
	// _ = "end of CoverTab[7426]"
//line /snap/go/10455/src/net/net.go:237
	_go_fuzz_dep_.CoverTab[7427]++
						return nil
//line /snap/go/10455/src/net/net.go:238
	// _ = "end of CoverTab[7427]"
}

// SetReadDeadline implements the Conn SetReadDeadline method.
func (c *conn) SetReadDeadline(t time.Time) error {
//line /snap/go/10455/src/net/net.go:242
	_go_fuzz_dep_.CoverTab[7432]++
						if !c.ok() {
//line /snap/go/10455/src/net/net.go:243
		_go_fuzz_dep_.CoverTab[529260]++
//line /snap/go/10455/src/net/net.go:243
		_go_fuzz_dep_.CoverTab[7435]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/net.go:244
		// _ = "end of CoverTab[7435]"
	} else {
//line /snap/go/10455/src/net/net.go:245
		_go_fuzz_dep_.CoverTab[529261]++
//line /snap/go/10455/src/net/net.go:245
		_go_fuzz_dep_.CoverTab[7436]++
//line /snap/go/10455/src/net/net.go:245
		// _ = "end of CoverTab[7436]"
//line /snap/go/10455/src/net/net.go:245
	}
//line /snap/go/10455/src/net/net.go:245
	// _ = "end of CoverTab[7432]"
//line /snap/go/10455/src/net/net.go:245
	_go_fuzz_dep_.CoverTab[7433]++
						if err := c.fd.SetReadDeadline(t); err != nil {
//line /snap/go/10455/src/net/net.go:246
		_go_fuzz_dep_.CoverTab[529262]++
//line /snap/go/10455/src/net/net.go:246
		_go_fuzz_dep_.CoverTab[7437]++
							return &OpError{Op: "set", Net: c.fd.net, Source: nil, Addr: c.fd.laddr, Err: err}
//line /snap/go/10455/src/net/net.go:247
		// _ = "end of CoverTab[7437]"
	} else {
//line /snap/go/10455/src/net/net.go:248
		_go_fuzz_dep_.CoverTab[529263]++
//line /snap/go/10455/src/net/net.go:248
		_go_fuzz_dep_.CoverTab[7438]++
//line /snap/go/10455/src/net/net.go:248
		// _ = "end of CoverTab[7438]"
//line /snap/go/10455/src/net/net.go:248
	}
//line /snap/go/10455/src/net/net.go:248
	// _ = "end of CoverTab[7433]"
//line /snap/go/10455/src/net/net.go:248
	_go_fuzz_dep_.CoverTab[7434]++
						return nil
//line /snap/go/10455/src/net/net.go:249
	// _ = "end of CoverTab[7434]"
}

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (c *conn) SetWriteDeadline(t time.Time) error {
//line /snap/go/10455/src/net/net.go:253
	_go_fuzz_dep_.CoverTab[7439]++
						if !c.ok() {
//line /snap/go/10455/src/net/net.go:254
		_go_fuzz_dep_.CoverTab[529264]++
//line /snap/go/10455/src/net/net.go:254
		_go_fuzz_dep_.CoverTab[7442]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/net.go:255
		// _ = "end of CoverTab[7442]"
	} else {
//line /snap/go/10455/src/net/net.go:256
		_go_fuzz_dep_.CoverTab[529265]++
//line /snap/go/10455/src/net/net.go:256
		_go_fuzz_dep_.CoverTab[7443]++
//line /snap/go/10455/src/net/net.go:256
		// _ = "end of CoverTab[7443]"
//line /snap/go/10455/src/net/net.go:256
	}
//line /snap/go/10455/src/net/net.go:256
	// _ = "end of CoverTab[7439]"
//line /snap/go/10455/src/net/net.go:256
	_go_fuzz_dep_.CoverTab[7440]++
						if err := c.fd.SetWriteDeadline(t); err != nil {
//line /snap/go/10455/src/net/net.go:257
		_go_fuzz_dep_.CoverTab[529266]++
//line /snap/go/10455/src/net/net.go:257
		_go_fuzz_dep_.CoverTab[7444]++
							return &OpError{Op: "set", Net: c.fd.net, Source: nil, Addr: c.fd.laddr, Err: err}
//line /snap/go/10455/src/net/net.go:258
		// _ = "end of CoverTab[7444]"
	} else {
//line /snap/go/10455/src/net/net.go:259
		_go_fuzz_dep_.CoverTab[529267]++
//line /snap/go/10455/src/net/net.go:259
		_go_fuzz_dep_.CoverTab[7445]++
//line /snap/go/10455/src/net/net.go:259
		// _ = "end of CoverTab[7445]"
//line /snap/go/10455/src/net/net.go:259
	}
//line /snap/go/10455/src/net/net.go:259
	// _ = "end of CoverTab[7440]"
//line /snap/go/10455/src/net/net.go:259
	_go_fuzz_dep_.CoverTab[7441]++
						return nil
//line /snap/go/10455/src/net/net.go:260
	// _ = "end of CoverTab[7441]"
}

// SetReadBuffer sets the size of the operating system's
//line /snap/go/10455/src/net/net.go:263
// receive buffer associated with the connection.
//line /snap/go/10455/src/net/net.go:265
func (c *conn) SetReadBuffer(bytes int) error {
//line /snap/go/10455/src/net/net.go:265
	_go_fuzz_dep_.CoverTab[7446]++
						if !c.ok() {
//line /snap/go/10455/src/net/net.go:266
		_go_fuzz_dep_.CoverTab[529268]++
//line /snap/go/10455/src/net/net.go:266
		_go_fuzz_dep_.CoverTab[7449]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/net.go:267
		// _ = "end of CoverTab[7449]"
	} else {
//line /snap/go/10455/src/net/net.go:268
		_go_fuzz_dep_.CoverTab[529269]++
//line /snap/go/10455/src/net/net.go:268
		_go_fuzz_dep_.CoverTab[7450]++
//line /snap/go/10455/src/net/net.go:268
		// _ = "end of CoverTab[7450]"
//line /snap/go/10455/src/net/net.go:268
	}
//line /snap/go/10455/src/net/net.go:268
	// _ = "end of CoverTab[7446]"
//line /snap/go/10455/src/net/net.go:268
	_go_fuzz_dep_.CoverTab[7447]++
						if err := setReadBuffer(c.fd, bytes); err != nil {
//line /snap/go/10455/src/net/net.go:269
		_go_fuzz_dep_.CoverTab[529270]++
//line /snap/go/10455/src/net/net.go:269
		_go_fuzz_dep_.CoverTab[7451]++
							return &OpError{Op: "set", Net: c.fd.net, Source: nil, Addr: c.fd.laddr, Err: err}
//line /snap/go/10455/src/net/net.go:270
		// _ = "end of CoverTab[7451]"
	} else {
//line /snap/go/10455/src/net/net.go:271
		_go_fuzz_dep_.CoverTab[529271]++
//line /snap/go/10455/src/net/net.go:271
		_go_fuzz_dep_.CoverTab[7452]++
//line /snap/go/10455/src/net/net.go:271
		// _ = "end of CoverTab[7452]"
//line /snap/go/10455/src/net/net.go:271
	}
//line /snap/go/10455/src/net/net.go:271
	// _ = "end of CoverTab[7447]"
//line /snap/go/10455/src/net/net.go:271
	_go_fuzz_dep_.CoverTab[7448]++
						return nil
//line /snap/go/10455/src/net/net.go:272
	// _ = "end of CoverTab[7448]"
}

// SetWriteBuffer sets the size of the operating system's
//line /snap/go/10455/src/net/net.go:275
// transmit buffer associated with the connection.
//line /snap/go/10455/src/net/net.go:277
func (c *conn) SetWriteBuffer(bytes int) error {
//line /snap/go/10455/src/net/net.go:277
	_go_fuzz_dep_.CoverTab[7453]++
						if !c.ok() {
//line /snap/go/10455/src/net/net.go:278
		_go_fuzz_dep_.CoverTab[529272]++
//line /snap/go/10455/src/net/net.go:278
		_go_fuzz_dep_.CoverTab[7456]++
							return syscall.EINVAL
//line /snap/go/10455/src/net/net.go:279
		// _ = "end of CoverTab[7456]"
	} else {
//line /snap/go/10455/src/net/net.go:280
		_go_fuzz_dep_.CoverTab[529273]++
//line /snap/go/10455/src/net/net.go:280
		_go_fuzz_dep_.CoverTab[7457]++
//line /snap/go/10455/src/net/net.go:280
		// _ = "end of CoverTab[7457]"
//line /snap/go/10455/src/net/net.go:280
	}
//line /snap/go/10455/src/net/net.go:280
	// _ = "end of CoverTab[7453]"
//line /snap/go/10455/src/net/net.go:280
	_go_fuzz_dep_.CoverTab[7454]++
						if err := setWriteBuffer(c.fd, bytes); err != nil {
//line /snap/go/10455/src/net/net.go:281
		_go_fuzz_dep_.CoverTab[529274]++
//line /snap/go/10455/src/net/net.go:281
		_go_fuzz_dep_.CoverTab[7458]++
							return &OpError{Op: "set", Net: c.fd.net, Source: nil, Addr: c.fd.laddr, Err: err}
//line /snap/go/10455/src/net/net.go:282
		// _ = "end of CoverTab[7458]"
	} else {
//line /snap/go/10455/src/net/net.go:283
		_go_fuzz_dep_.CoverTab[529275]++
//line /snap/go/10455/src/net/net.go:283
		_go_fuzz_dep_.CoverTab[7459]++
//line /snap/go/10455/src/net/net.go:283
		// _ = "end of CoverTab[7459]"
//line /snap/go/10455/src/net/net.go:283
	}
//line /snap/go/10455/src/net/net.go:283
	// _ = "end of CoverTab[7454]"
//line /snap/go/10455/src/net/net.go:283
	_go_fuzz_dep_.CoverTab[7455]++
						return nil
//line /snap/go/10455/src/net/net.go:284
	// _ = "end of CoverTab[7455]"
}

// File returns a copy of the underlying os.File.
//line /snap/go/10455/src/net/net.go:287
// It is the caller's responsibility to close f when finished.
//line /snap/go/10455/src/net/net.go:287
// Closing c does not affect f, and closing f does not affect c.
//line /snap/go/10455/src/net/net.go:287
//
//line /snap/go/10455/src/net/net.go:287
// The returned os.File's file descriptor is different from the connection's.
//line /snap/go/10455/src/net/net.go:287
// Attempting to change properties of the original using this duplicate
//line /snap/go/10455/src/net/net.go:287
// may or may not have the desired effect.
//line /snap/go/10455/src/net/net.go:294
func (c *conn) File() (f *os.File, err error) {
//line /snap/go/10455/src/net/net.go:294
	_go_fuzz_dep_.CoverTab[7460]++
						f, err = c.fd.dup()
						if err != nil {
//line /snap/go/10455/src/net/net.go:296
		_go_fuzz_dep_.CoverTab[529276]++
//line /snap/go/10455/src/net/net.go:296
		_go_fuzz_dep_.CoverTab[7462]++
							err = &OpError{Op: "file", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
//line /snap/go/10455/src/net/net.go:297
		// _ = "end of CoverTab[7462]"
	} else {
//line /snap/go/10455/src/net/net.go:298
		_go_fuzz_dep_.CoverTab[529277]++
//line /snap/go/10455/src/net/net.go:298
		_go_fuzz_dep_.CoverTab[7463]++
//line /snap/go/10455/src/net/net.go:298
		// _ = "end of CoverTab[7463]"
//line /snap/go/10455/src/net/net.go:298
	}
//line /snap/go/10455/src/net/net.go:298
	// _ = "end of CoverTab[7460]"
//line /snap/go/10455/src/net/net.go:298
	_go_fuzz_dep_.CoverTab[7461]++
						return
//line /snap/go/10455/src/net/net.go:299
	// _ = "end of CoverTab[7461]"
}

// PacketConn is a generic packet-oriented network connection.
//line /snap/go/10455/src/net/net.go:302
//
//line /snap/go/10455/src/net/net.go:302
// Multiple goroutines may invoke methods on a PacketConn simultaneously.
//line /snap/go/10455/src/net/net.go:305
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
//line /snap/go/10455/src/net/net.go:372
	_go_fuzz_dep_.CoverTab[7464]++
						listenerBacklogCache.Do(func() {
//line /snap/go/10455/src/net/net.go:373
		_go_fuzz_dep_.CoverTab[7466]++
//line /snap/go/10455/src/net/net.go:373
		listenerBacklogCache.val = maxListenerBacklog()
//line /snap/go/10455/src/net/net.go:373
		// _ = "end of CoverTab[7466]"
//line /snap/go/10455/src/net/net.go:373
	})
//line /snap/go/10455/src/net/net.go:373
	// _ = "end of CoverTab[7464]"
//line /snap/go/10455/src/net/net.go:373
	_go_fuzz_dep_.CoverTab[7465]++
						return listenerBacklogCache.val
//line /snap/go/10455/src/net/net.go:374
	// _ = "end of CoverTab[7465]"
}

// A Listener is a generic network listener for stream-oriented protocols.
//line /snap/go/10455/src/net/net.go:377
//
//line /snap/go/10455/src/net/net.go:377
// Multiple goroutines may invoke methods on a Listener simultaneously.
//line /snap/go/10455/src/net/net.go:380
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
//line /snap/go/10455/src/net/net.go:416
// returned, while still being Is context.Canceled.
//line /snap/go/10455/src/net/net.go:418
type canceledError struct{}

func (canceledError) Error() string {
//line /snap/go/10455/src/net/net.go:420
	_go_fuzz_dep_.CoverTab[7467]++
//line /snap/go/10455/src/net/net.go:420
	return "operation was canceled"
//line /snap/go/10455/src/net/net.go:420
	// _ = "end of CoverTab[7467]"
//line /snap/go/10455/src/net/net.go:420
}

func (canceledError) Is(err error) bool {
//line /snap/go/10455/src/net/net.go:422
	_go_fuzz_dep_.CoverTab[7468]++
//line /snap/go/10455/src/net/net.go:422
	return err == context.Canceled
//line /snap/go/10455/src/net/net.go:422
	// _ = "end of CoverTab[7468]"
//line /snap/go/10455/src/net/net.go:422
}

// mapErr maps from the context errors to the historical internal net
//line /snap/go/10455/src/net/net.go:424
// error values.
//line /snap/go/10455/src/net/net.go:426
func mapErr(err error) error {
//line /snap/go/10455/src/net/net.go:426
	_go_fuzz_dep_.CoverTab[7469]++
						switch err {
	case context.Canceled:
//line /snap/go/10455/src/net/net.go:428
		_go_fuzz_dep_.CoverTab[529278]++
//line /snap/go/10455/src/net/net.go:428
		_go_fuzz_dep_.CoverTab[7470]++
							return errCanceled
//line /snap/go/10455/src/net/net.go:429
		// _ = "end of CoverTab[7470]"
	case context.DeadlineExceeded:
//line /snap/go/10455/src/net/net.go:430
		_go_fuzz_dep_.CoverTab[529279]++
//line /snap/go/10455/src/net/net.go:430
		_go_fuzz_dep_.CoverTab[7471]++
							return errTimeout
//line /snap/go/10455/src/net/net.go:431
		// _ = "end of CoverTab[7471]"
	default:
//line /snap/go/10455/src/net/net.go:432
		_go_fuzz_dep_.CoverTab[529280]++
//line /snap/go/10455/src/net/net.go:432
		_go_fuzz_dep_.CoverTab[7472]++
							return err
//line /snap/go/10455/src/net/net.go:433
		// _ = "end of CoverTab[7472]"
	}
//line /snap/go/10455/src/net/net.go:434
	// _ = "end of CoverTab[7469]"
}

// OpError is the error type usually returned by functions in the net
//line /snap/go/10455/src/net/net.go:437
// package. It describes the operation, network type, and address of
//line /snap/go/10455/src/net/net.go:437
// an error.
//line /snap/go/10455/src/net/net.go:440
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
//line /snap/go/10455/src/net/net.go:467
	_go_fuzz_dep_.CoverTab[7473]++
//line /snap/go/10455/src/net/net.go:467
	return e.Err
//line /snap/go/10455/src/net/net.go:467
	// _ = "end of CoverTab[7473]"
//line /snap/go/10455/src/net/net.go:467
}

func (e *OpError) Error() string {
//line /snap/go/10455/src/net/net.go:469
	_go_fuzz_dep_.CoverTab[7474]++
						if e == nil {
//line /snap/go/10455/src/net/net.go:470
		_go_fuzz_dep_.CoverTab[529281]++
//line /snap/go/10455/src/net/net.go:470
		_go_fuzz_dep_.CoverTab[7479]++
							return "<nil>"
//line /snap/go/10455/src/net/net.go:471
		// _ = "end of CoverTab[7479]"
	} else {
//line /snap/go/10455/src/net/net.go:472
		_go_fuzz_dep_.CoverTab[529282]++
//line /snap/go/10455/src/net/net.go:472
		_go_fuzz_dep_.CoverTab[7480]++
//line /snap/go/10455/src/net/net.go:472
		// _ = "end of CoverTab[7480]"
//line /snap/go/10455/src/net/net.go:472
	}
//line /snap/go/10455/src/net/net.go:472
	// _ = "end of CoverTab[7474]"
//line /snap/go/10455/src/net/net.go:472
	_go_fuzz_dep_.CoverTab[7475]++
						s := e.Op
						if e.Net != "" {
//line /snap/go/10455/src/net/net.go:474
		_go_fuzz_dep_.CoverTab[529283]++
//line /snap/go/10455/src/net/net.go:474
		_go_fuzz_dep_.CoverTab[7481]++
							s += " " + e.Net
//line /snap/go/10455/src/net/net.go:475
		// _ = "end of CoverTab[7481]"
	} else {
//line /snap/go/10455/src/net/net.go:476
		_go_fuzz_dep_.CoverTab[529284]++
//line /snap/go/10455/src/net/net.go:476
		_go_fuzz_dep_.CoverTab[7482]++
//line /snap/go/10455/src/net/net.go:476
		// _ = "end of CoverTab[7482]"
//line /snap/go/10455/src/net/net.go:476
	}
//line /snap/go/10455/src/net/net.go:476
	// _ = "end of CoverTab[7475]"
//line /snap/go/10455/src/net/net.go:476
	_go_fuzz_dep_.CoverTab[7476]++
						if e.Source != nil {
//line /snap/go/10455/src/net/net.go:477
		_go_fuzz_dep_.CoverTab[529285]++
//line /snap/go/10455/src/net/net.go:477
		_go_fuzz_dep_.CoverTab[7483]++
							s += " " + e.Source.String()
//line /snap/go/10455/src/net/net.go:478
		// _ = "end of CoverTab[7483]"
	} else {
//line /snap/go/10455/src/net/net.go:479
		_go_fuzz_dep_.CoverTab[529286]++
//line /snap/go/10455/src/net/net.go:479
		_go_fuzz_dep_.CoverTab[7484]++
//line /snap/go/10455/src/net/net.go:479
		// _ = "end of CoverTab[7484]"
//line /snap/go/10455/src/net/net.go:479
	}
//line /snap/go/10455/src/net/net.go:479
	// _ = "end of CoverTab[7476]"
//line /snap/go/10455/src/net/net.go:479
	_go_fuzz_dep_.CoverTab[7477]++
						if e.Addr != nil {
//line /snap/go/10455/src/net/net.go:480
		_go_fuzz_dep_.CoverTab[529287]++
//line /snap/go/10455/src/net/net.go:480
		_go_fuzz_dep_.CoverTab[7485]++
							if e.Source != nil {
//line /snap/go/10455/src/net/net.go:481
			_go_fuzz_dep_.CoverTab[529289]++
//line /snap/go/10455/src/net/net.go:481
			_go_fuzz_dep_.CoverTab[7487]++
								s += "->"
//line /snap/go/10455/src/net/net.go:482
			// _ = "end of CoverTab[7487]"
		} else {
//line /snap/go/10455/src/net/net.go:483
			_go_fuzz_dep_.CoverTab[529290]++
//line /snap/go/10455/src/net/net.go:483
			_go_fuzz_dep_.CoverTab[7488]++
								s += " "
//line /snap/go/10455/src/net/net.go:484
			// _ = "end of CoverTab[7488]"
		}
//line /snap/go/10455/src/net/net.go:485
		// _ = "end of CoverTab[7485]"
//line /snap/go/10455/src/net/net.go:485
		_go_fuzz_dep_.CoverTab[7486]++
							s += e.Addr.String()
//line /snap/go/10455/src/net/net.go:486
		// _ = "end of CoverTab[7486]"
	} else {
//line /snap/go/10455/src/net/net.go:487
		_go_fuzz_dep_.CoverTab[529288]++
//line /snap/go/10455/src/net/net.go:487
		_go_fuzz_dep_.CoverTab[7489]++
//line /snap/go/10455/src/net/net.go:487
		// _ = "end of CoverTab[7489]"
//line /snap/go/10455/src/net/net.go:487
	}
//line /snap/go/10455/src/net/net.go:487
	// _ = "end of CoverTab[7477]"
//line /snap/go/10455/src/net/net.go:487
	_go_fuzz_dep_.CoverTab[7478]++
						s += ": " + e.Err.Error()
						return s
//line /snap/go/10455/src/net/net.go:489
	// _ = "end of CoverTab[7478]"
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
//line /snap/go/10455/src/net/net.go:507
	_go_fuzz_dep_.CoverTab[7490]++
						if ne, ok := e.Err.(*os.SyscallError); ok {
//line /snap/go/10455/src/net/net.go:508
		_go_fuzz_dep_.CoverTab[529291]++
//line /snap/go/10455/src/net/net.go:508
		_go_fuzz_dep_.CoverTab[7492]++
							t, ok := ne.Err.(timeout)
							return ok && func() bool {
//line /snap/go/10455/src/net/net.go:510
			_go_fuzz_dep_.CoverTab[7493]++
//line /snap/go/10455/src/net/net.go:510
			return t.Timeout()
//line /snap/go/10455/src/net/net.go:510
			// _ = "end of CoverTab[7493]"
//line /snap/go/10455/src/net/net.go:510
		}()
//line /snap/go/10455/src/net/net.go:510
		// _ = "end of CoverTab[7492]"
	} else {
//line /snap/go/10455/src/net/net.go:511
		_go_fuzz_dep_.CoverTab[529292]++
//line /snap/go/10455/src/net/net.go:511
		_go_fuzz_dep_.CoverTab[7494]++
//line /snap/go/10455/src/net/net.go:511
		// _ = "end of CoverTab[7494]"
//line /snap/go/10455/src/net/net.go:511
	}
//line /snap/go/10455/src/net/net.go:511
	// _ = "end of CoverTab[7490]"
//line /snap/go/10455/src/net/net.go:511
	_go_fuzz_dep_.CoverTab[7491]++
						t, ok := e.Err.(timeout)
						return ok && func() bool {
//line /snap/go/10455/src/net/net.go:513
		_go_fuzz_dep_.CoverTab[7495]++
//line /snap/go/10455/src/net/net.go:513
		return t.Timeout()
//line /snap/go/10455/src/net/net.go:513
		// _ = "end of CoverTab[7495]"
//line /snap/go/10455/src/net/net.go:513
	}()
//line /snap/go/10455/src/net/net.go:513
	// _ = "end of CoverTab[7491]"
}

type temporary interface {
	Temporary() bool
}

func (e *OpError) Temporary() bool {
//line /snap/go/10455/src/net/net.go:520
	_go_fuzz_dep_.CoverTab[7496]++

//line /snap/go/10455/src/net/net.go:523
	if e.Op == "accept" && func() bool {
//line /snap/go/10455/src/net/net.go:523
		_go_fuzz_dep_.CoverTab[7499]++
//line /snap/go/10455/src/net/net.go:523
		return isConnError(e.Err)
//line /snap/go/10455/src/net/net.go:523
		// _ = "end of CoverTab[7499]"
//line /snap/go/10455/src/net/net.go:523
	}() {
//line /snap/go/10455/src/net/net.go:523
		_go_fuzz_dep_.CoverTab[529293]++
//line /snap/go/10455/src/net/net.go:523
		_go_fuzz_dep_.CoverTab[7500]++
							return true
//line /snap/go/10455/src/net/net.go:524
		// _ = "end of CoverTab[7500]"
	} else {
//line /snap/go/10455/src/net/net.go:525
		_go_fuzz_dep_.CoverTab[529294]++
//line /snap/go/10455/src/net/net.go:525
		_go_fuzz_dep_.CoverTab[7501]++
//line /snap/go/10455/src/net/net.go:525
		// _ = "end of CoverTab[7501]"
//line /snap/go/10455/src/net/net.go:525
	}
//line /snap/go/10455/src/net/net.go:525
	// _ = "end of CoverTab[7496]"
//line /snap/go/10455/src/net/net.go:525
	_go_fuzz_dep_.CoverTab[7497]++

						if ne, ok := e.Err.(*os.SyscallError); ok {
//line /snap/go/10455/src/net/net.go:527
		_go_fuzz_dep_.CoverTab[529295]++
//line /snap/go/10455/src/net/net.go:527
		_go_fuzz_dep_.CoverTab[7502]++
							t, ok := ne.Err.(temporary)
							return ok && func() bool {
//line /snap/go/10455/src/net/net.go:529
			_go_fuzz_dep_.CoverTab[7503]++
//line /snap/go/10455/src/net/net.go:529
			return t.Temporary()
//line /snap/go/10455/src/net/net.go:529
			// _ = "end of CoverTab[7503]"
//line /snap/go/10455/src/net/net.go:529
		}()
//line /snap/go/10455/src/net/net.go:529
		// _ = "end of CoverTab[7502]"
	} else {
//line /snap/go/10455/src/net/net.go:530
		_go_fuzz_dep_.CoverTab[529296]++
//line /snap/go/10455/src/net/net.go:530
		_go_fuzz_dep_.CoverTab[7504]++
//line /snap/go/10455/src/net/net.go:530
		// _ = "end of CoverTab[7504]"
//line /snap/go/10455/src/net/net.go:530
	}
//line /snap/go/10455/src/net/net.go:530
	// _ = "end of CoverTab[7497]"
//line /snap/go/10455/src/net/net.go:530
	_go_fuzz_dep_.CoverTab[7498]++
						t, ok := e.Err.(temporary)
						return ok && func() bool {
//line /snap/go/10455/src/net/net.go:532
		_go_fuzz_dep_.CoverTab[7505]++
//line /snap/go/10455/src/net/net.go:532
		return t.Temporary()
//line /snap/go/10455/src/net/net.go:532
		// _ = "end of CoverTab[7505]"
//line /snap/go/10455/src/net/net.go:532
	}()
//line /snap/go/10455/src/net/net.go:532
	// _ = "end of CoverTab[7498]"
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
//line /snap/go/10455/src/net/net.go:545
	_go_fuzz_dep_.CoverTab[7506]++
//line /snap/go/10455/src/net/net.go:545
	return "invalid " + e.Type + ": " + e.Text
//line /snap/go/10455/src/net/net.go:545
	// _ = "end of CoverTab[7506]"
//line /snap/go/10455/src/net/net.go:545
}

func (e *ParseError) Timeout() bool {
//line /snap/go/10455/src/net/net.go:547
	_go_fuzz_dep_.CoverTab[7507]++
//line /snap/go/10455/src/net/net.go:547
	return false
//line /snap/go/10455/src/net/net.go:547
	// _ = "end of CoverTab[7507]"
//line /snap/go/10455/src/net/net.go:547
}
func (e *ParseError) Temporary() bool {
//line /snap/go/10455/src/net/net.go:548
	_go_fuzz_dep_.CoverTab[7508]++
//line /snap/go/10455/src/net/net.go:548
	return false
//line /snap/go/10455/src/net/net.go:548
	// _ = "end of CoverTab[7508]"
//line /snap/go/10455/src/net/net.go:548
}

type AddrError struct {
	Err	string
	Addr	string
}

func (e *AddrError) Error() string {
//line /snap/go/10455/src/net/net.go:555
	_go_fuzz_dep_.CoverTab[7509]++
						if e == nil {
//line /snap/go/10455/src/net/net.go:556
		_go_fuzz_dep_.CoverTab[529297]++
//line /snap/go/10455/src/net/net.go:556
		_go_fuzz_dep_.CoverTab[7512]++
							return "<nil>"
//line /snap/go/10455/src/net/net.go:557
		// _ = "end of CoverTab[7512]"
	} else {
//line /snap/go/10455/src/net/net.go:558
		_go_fuzz_dep_.CoverTab[529298]++
//line /snap/go/10455/src/net/net.go:558
		_go_fuzz_dep_.CoverTab[7513]++
//line /snap/go/10455/src/net/net.go:558
		// _ = "end of CoverTab[7513]"
//line /snap/go/10455/src/net/net.go:558
	}
//line /snap/go/10455/src/net/net.go:558
	// _ = "end of CoverTab[7509]"
//line /snap/go/10455/src/net/net.go:558
	_go_fuzz_dep_.CoverTab[7510]++
						s := e.Err
						if e.Addr != "" {
//line /snap/go/10455/src/net/net.go:560
		_go_fuzz_dep_.CoverTab[529299]++
//line /snap/go/10455/src/net/net.go:560
		_go_fuzz_dep_.CoverTab[7514]++
							s = "address " + e.Addr + ": " + s
//line /snap/go/10455/src/net/net.go:561
		// _ = "end of CoverTab[7514]"
	} else {
//line /snap/go/10455/src/net/net.go:562
		_go_fuzz_dep_.CoverTab[529300]++
//line /snap/go/10455/src/net/net.go:562
		_go_fuzz_dep_.CoverTab[7515]++
//line /snap/go/10455/src/net/net.go:562
		// _ = "end of CoverTab[7515]"
//line /snap/go/10455/src/net/net.go:562
	}
//line /snap/go/10455/src/net/net.go:562
	// _ = "end of CoverTab[7510]"
//line /snap/go/10455/src/net/net.go:562
	_go_fuzz_dep_.CoverTab[7511]++
						return s
//line /snap/go/10455/src/net/net.go:563
	// _ = "end of CoverTab[7511]"
}

func (e *AddrError) Timeout() bool {
//line /snap/go/10455/src/net/net.go:566
	_go_fuzz_dep_.CoverTab[7516]++
//line /snap/go/10455/src/net/net.go:566
	return false
//line /snap/go/10455/src/net/net.go:566
	// _ = "end of CoverTab[7516]"
//line /snap/go/10455/src/net/net.go:566
}
func (e *AddrError) Temporary() bool {
//line /snap/go/10455/src/net/net.go:567
	_go_fuzz_dep_.CoverTab[7517]++
//line /snap/go/10455/src/net/net.go:567
	return false
//line /snap/go/10455/src/net/net.go:567
	// _ = "end of CoverTab[7517]"
//line /snap/go/10455/src/net/net.go:567
}

type UnknownNetworkError string

func (e UnknownNetworkError) Error() string {
//line /snap/go/10455/src/net/net.go:571
	_go_fuzz_dep_.CoverTab[7518]++
//line /snap/go/10455/src/net/net.go:571
	return "unknown network " + string(e)
//line /snap/go/10455/src/net/net.go:571
	// _ = "end of CoverTab[7518]"
//line /snap/go/10455/src/net/net.go:571
}
func (e UnknownNetworkError) Timeout() bool {
//line /snap/go/10455/src/net/net.go:572
	_go_fuzz_dep_.CoverTab[7519]++
//line /snap/go/10455/src/net/net.go:572
	return false
//line /snap/go/10455/src/net/net.go:572
	// _ = "end of CoverTab[7519]"
//line /snap/go/10455/src/net/net.go:572
}
func (e UnknownNetworkError) Temporary() bool {
//line /snap/go/10455/src/net/net.go:573
	_go_fuzz_dep_.CoverTab[7520]++
//line /snap/go/10455/src/net/net.go:573
	return false
//line /snap/go/10455/src/net/net.go:573
	// _ = "end of CoverTab[7520]"
//line /snap/go/10455/src/net/net.go:573
}

type InvalidAddrError string

func (e InvalidAddrError) Error() string {
//line /snap/go/10455/src/net/net.go:577
	_go_fuzz_dep_.CoverTab[7521]++
//line /snap/go/10455/src/net/net.go:577
	return string(e)
//line /snap/go/10455/src/net/net.go:577
	// _ = "end of CoverTab[7521]"
//line /snap/go/10455/src/net/net.go:577
}
func (e InvalidAddrError) Timeout() bool {
//line /snap/go/10455/src/net/net.go:578
	_go_fuzz_dep_.CoverTab[7522]++
//line /snap/go/10455/src/net/net.go:578
	return false
//line /snap/go/10455/src/net/net.go:578
	// _ = "end of CoverTab[7522]"
//line /snap/go/10455/src/net/net.go:578
}
func (e InvalidAddrError) Temporary() bool {
//line /snap/go/10455/src/net/net.go:579
	_go_fuzz_dep_.CoverTab[7523]++
//line /snap/go/10455/src/net/net.go:579
	return false
//line /snap/go/10455/src/net/net.go:579
	// _ = "end of CoverTab[7523]"
//line /snap/go/10455/src/net/net.go:579
}

// errTimeout exists to return the historical "i/o timeout" string
//line /snap/go/10455/src/net/net.go:581
// for context.DeadlineExceeded. See mapErr.
//line /snap/go/10455/src/net/net.go:581
// It is also used when Dialer.Deadline is exceeded.
//line /snap/go/10455/src/net/net.go:581
// error.Is(errTimeout, context.DeadlineExceeded) returns true.
//line /snap/go/10455/src/net/net.go:581
//
//line /snap/go/10455/src/net/net.go:581
// TODO(iant): We could consider changing this to os.ErrDeadlineExceeded
//line /snap/go/10455/src/net/net.go:581
// in the future, if we make
//line /snap/go/10455/src/net/net.go:581
//
//line /snap/go/10455/src/net/net.go:581
//	errors.Is(os.ErrDeadlineExceeded, context.DeadlineExceeded)
//line /snap/go/10455/src/net/net.go:581
//
//line /snap/go/10455/src/net/net.go:581
// return true.
//line /snap/go/10455/src/net/net.go:592
var errTimeout error = &timeoutError{}

type timeoutError struct{}

func (e *timeoutError) Error() string {
//line /snap/go/10455/src/net/net.go:596
	_go_fuzz_dep_.CoverTab[7524]++
//line /snap/go/10455/src/net/net.go:596
	return "i/o timeout"
//line /snap/go/10455/src/net/net.go:596
	// _ = "end of CoverTab[7524]"
//line /snap/go/10455/src/net/net.go:596
}
func (e *timeoutError) Timeout() bool {
//line /snap/go/10455/src/net/net.go:597
	_go_fuzz_dep_.CoverTab[7525]++
//line /snap/go/10455/src/net/net.go:597
	return true
//line /snap/go/10455/src/net/net.go:597
	// _ = "end of CoverTab[7525]"
//line /snap/go/10455/src/net/net.go:597
}
func (e *timeoutError) Temporary() bool {
//line /snap/go/10455/src/net/net.go:598
	_go_fuzz_dep_.CoverTab[7526]++
//line /snap/go/10455/src/net/net.go:598
	return true
//line /snap/go/10455/src/net/net.go:598
	// _ = "end of CoverTab[7526]"
//line /snap/go/10455/src/net/net.go:598
}

func (e *timeoutError) Is(err error) bool {
//line /snap/go/10455/src/net/net.go:600
	_go_fuzz_dep_.CoverTab[7527]++
						return err == context.DeadlineExceeded
//line /snap/go/10455/src/net/net.go:601
	// _ = "end of CoverTab[7527]"
}

// DNSConfigError represents an error reading the machine's DNS configuration.
//line /snap/go/10455/src/net/net.go:604
// (No longer used; kept for compatibility.)
//line /snap/go/10455/src/net/net.go:606
type DNSConfigError struct {
	Err error
}

func (e *DNSConfigError) Unwrap() error {
//line /snap/go/10455/src/net/net.go:610
	_go_fuzz_dep_.CoverTab[7528]++
//line /snap/go/10455/src/net/net.go:610
	return e.Err
//line /snap/go/10455/src/net/net.go:610
	// _ = "end of CoverTab[7528]"
//line /snap/go/10455/src/net/net.go:610
}
func (e *DNSConfigError) Error() string {
//line /snap/go/10455/src/net/net.go:611
	_go_fuzz_dep_.CoverTab[7529]++
//line /snap/go/10455/src/net/net.go:611
	return "error reading DNS config: " + e.Err.Error()
//line /snap/go/10455/src/net/net.go:611
	// _ = "end of CoverTab[7529]"
//line /snap/go/10455/src/net/net.go:611
}
func (e *DNSConfigError) Timeout() bool {
//line /snap/go/10455/src/net/net.go:612
	_go_fuzz_dep_.CoverTab[7530]++
//line /snap/go/10455/src/net/net.go:612
	return false
//line /snap/go/10455/src/net/net.go:612
	// _ = "end of CoverTab[7530]"
//line /snap/go/10455/src/net/net.go:612
}
func (e *DNSConfigError) Temporary() bool {
//line /snap/go/10455/src/net/net.go:613
	_go_fuzz_dep_.CoverTab[7531]++
//line /snap/go/10455/src/net/net.go:613
	return false
//line /snap/go/10455/src/net/net.go:613
	// _ = "end of CoverTab[7531]"
//line /snap/go/10455/src/net/net.go:613
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
//line /snap/go/10455/src/net/net.go:630
	_go_fuzz_dep_.CoverTab[7532]++
						if e == nil {
//line /snap/go/10455/src/net/net.go:631
		_go_fuzz_dep_.CoverTab[529301]++
//line /snap/go/10455/src/net/net.go:631
		_go_fuzz_dep_.CoverTab[7535]++
							return "<nil>"
//line /snap/go/10455/src/net/net.go:632
		// _ = "end of CoverTab[7535]"
	} else {
//line /snap/go/10455/src/net/net.go:633
		_go_fuzz_dep_.CoverTab[529302]++
//line /snap/go/10455/src/net/net.go:633
		_go_fuzz_dep_.CoverTab[7536]++
//line /snap/go/10455/src/net/net.go:633
		// _ = "end of CoverTab[7536]"
//line /snap/go/10455/src/net/net.go:633
	}
//line /snap/go/10455/src/net/net.go:633
	// _ = "end of CoverTab[7532]"
//line /snap/go/10455/src/net/net.go:633
	_go_fuzz_dep_.CoverTab[7533]++
						s := "lookup " + e.Name
						if e.Server != "" {
//line /snap/go/10455/src/net/net.go:635
		_go_fuzz_dep_.CoverTab[529303]++
//line /snap/go/10455/src/net/net.go:635
		_go_fuzz_dep_.CoverTab[7537]++
							s += " on " + e.Server
//line /snap/go/10455/src/net/net.go:636
		// _ = "end of CoverTab[7537]"
	} else {
//line /snap/go/10455/src/net/net.go:637
		_go_fuzz_dep_.CoverTab[529304]++
//line /snap/go/10455/src/net/net.go:637
		_go_fuzz_dep_.CoverTab[7538]++
//line /snap/go/10455/src/net/net.go:637
		// _ = "end of CoverTab[7538]"
//line /snap/go/10455/src/net/net.go:637
	}
//line /snap/go/10455/src/net/net.go:637
	// _ = "end of CoverTab[7533]"
//line /snap/go/10455/src/net/net.go:637
	_go_fuzz_dep_.CoverTab[7534]++
						s += ": " + e.Err
						return s
//line /snap/go/10455/src/net/net.go:639
	// _ = "end of CoverTab[7534]"
}

// Timeout reports whether the DNS lookup is known to have timed out.
//line /snap/go/10455/src/net/net.go:642
// This is not always known; a DNS lookup may fail due to a timeout
//line /snap/go/10455/src/net/net.go:642
// and return a DNSError for which Timeout returns false.
//line /snap/go/10455/src/net/net.go:645
func (e *DNSError) Timeout() bool {
//line /snap/go/10455/src/net/net.go:645
	_go_fuzz_dep_.CoverTab[7539]++
//line /snap/go/10455/src/net/net.go:645
	return e.IsTimeout
//line /snap/go/10455/src/net/net.go:645
	// _ = "end of CoverTab[7539]"
//line /snap/go/10455/src/net/net.go:645
}

// Temporary reports whether the DNS error is known to be temporary.
//line /snap/go/10455/src/net/net.go:647
// This is not always known; a DNS lookup may fail due to a temporary
//line /snap/go/10455/src/net/net.go:647
// error and return a DNSError for which Temporary returns false.
//line /snap/go/10455/src/net/net.go:650
func (e *DNSError) Temporary() bool {
//line /snap/go/10455/src/net/net.go:650
	_go_fuzz_dep_.CoverTab[7540]++
//line /snap/go/10455/src/net/net.go:650
	return e.IsTimeout || func() bool {
//line /snap/go/10455/src/net/net.go:650
		_go_fuzz_dep_.CoverTab[7541]++
//line /snap/go/10455/src/net/net.go:650
		return e.IsTemporary
//line /snap/go/10455/src/net/net.go:650
		// _ = "end of CoverTab[7541]"
//line /snap/go/10455/src/net/net.go:650
	}()
//line /snap/go/10455/src/net/net.go:650
	// _ = "end of CoverTab[7540]"
//line /snap/go/10455/src/net/net.go:650
}

// errClosed exists just so that the docs for ErrClosed don't mention
//line /snap/go/10455/src/net/net.go:652
// the internal package poll.
//line /snap/go/10455/src/net/net.go:654
var errClosed = poll.ErrNetClosing

// ErrClosed is the error returned by an I/O call on a network
//line /snap/go/10455/src/net/net.go:656
// connection that has already been closed, or that is closed by
//line /snap/go/10455/src/net/net.go:656
// another goroutine before the I/O is completed. This may be wrapped
//line /snap/go/10455/src/net/net.go:656
// in another error, and should normally be tested using
//line /snap/go/10455/src/net/net.go:656
// errors.Is(err, net.ErrClosed).
//line /snap/go/10455/src/net/net.go:661
var ErrClosed error = errClosed

type writerOnly struct {
	io.Writer
}

// Fallback implementation of io.ReaderFrom's ReadFrom, when sendfile isn't
//line /snap/go/10455/src/net/net.go:667
// applicable.
//line /snap/go/10455/src/net/net.go:669
func genericReadFrom(w io.Writer, r io.Reader) (n int64, err error) {
//line /snap/go/10455/src/net/net.go:669
	_go_fuzz_dep_.CoverTab[7542]++

						return io.Copy(writerOnly{w}, r)
//line /snap/go/10455/src/net/net.go:671
	// _ = "end of CoverTab[7542]"
}

//line /snap/go/10455/src/net/net.go:680
var threadLimit chan struct{}

var threadOnce sync.Once

func acquireThread() {
//line /snap/go/10455/src/net/net.go:684
	_go_fuzz_dep_.CoverTab[7543]++
						threadOnce.Do(func() {
//line /snap/go/10455/src/net/net.go:685
		_go_fuzz_dep_.CoverTab[7545]++
							threadLimit = make(chan struct{}, concurrentThreadsLimit())
//line /snap/go/10455/src/net/net.go:686
		// _ = "end of CoverTab[7545]"
	})
//line /snap/go/10455/src/net/net.go:687
	// _ = "end of CoverTab[7543]"
//line /snap/go/10455/src/net/net.go:687
	_go_fuzz_dep_.CoverTab[7544]++
						threadLimit <- struct{}{}
//line /snap/go/10455/src/net/net.go:688
	// _ = "end of CoverTab[7544]"
}

func releaseThread() {
//line /snap/go/10455/src/net/net.go:691
	_go_fuzz_dep_.CoverTab[7546]++
						<-threadLimit
//line /snap/go/10455/src/net/net.go:692
	// _ = "end of CoverTab[7546]"
}

// buffersWriter is the interface implemented by Conns that support a
//line /snap/go/10455/src/net/net.go:695
// "writev"-like batch write optimization.
//line /snap/go/10455/src/net/net.go:695
// writeBuffers should fully consume and write all chunks from the
//line /snap/go/10455/src/net/net.go:695
// provided Buffers, else it should report a non-nil error.
//line /snap/go/10455/src/net/net.go:699
type buffersWriter interface {
	writeBuffers(*Buffers) (int64, error)
}

// Buffers contains zero or more runs of bytes to write.
//line /snap/go/10455/src/net/net.go:703
//
//line /snap/go/10455/src/net/net.go:703
// On certain machines, for certain types of connections, this is
//line /snap/go/10455/src/net/net.go:703
// optimized into an OS-specific batch write operation (such as
//line /snap/go/10455/src/net/net.go:703
// "writev").
//line /snap/go/10455/src/net/net.go:708
type Buffers [][]byte

var (
	_	io.WriterTo	= (*Buffers)(nil)
	_	io.Reader	= (*Buffers)(nil)
)

// WriteTo writes contents of the buffers to w.
//line /snap/go/10455/src/net/net.go:715
//
//line /snap/go/10455/src/net/net.go:715
// WriteTo implements io.WriterTo for Buffers.
//line /snap/go/10455/src/net/net.go:715
//
//line /snap/go/10455/src/net/net.go:715
// WriteTo modifies the slice v as well as v[i] for 0 <= i < len(v),
//line /snap/go/10455/src/net/net.go:715
// but does not modify v[i][j] for any i, j.
//line /snap/go/10455/src/net/net.go:721
func (v *Buffers) WriteTo(w io.Writer) (n int64, err error) {
//line /snap/go/10455/src/net/net.go:721
	_go_fuzz_dep_.CoverTab[7547]++
						if wv, ok := w.(buffersWriter); ok {
//line /snap/go/10455/src/net/net.go:722
		_go_fuzz_dep_.CoverTab[529305]++
//line /snap/go/10455/src/net/net.go:722
		_go_fuzz_dep_.CoverTab[7550]++
							return wv.writeBuffers(v)
//line /snap/go/10455/src/net/net.go:723
		// _ = "end of CoverTab[7550]"
	} else {
//line /snap/go/10455/src/net/net.go:724
		_go_fuzz_dep_.CoverTab[529306]++
//line /snap/go/10455/src/net/net.go:724
		_go_fuzz_dep_.CoverTab[7551]++
//line /snap/go/10455/src/net/net.go:724
		// _ = "end of CoverTab[7551]"
//line /snap/go/10455/src/net/net.go:724
	}
//line /snap/go/10455/src/net/net.go:724
	// _ = "end of CoverTab[7547]"
//line /snap/go/10455/src/net/net.go:724
	_go_fuzz_dep_.CoverTab[7548]++
//line /snap/go/10455/src/net/net.go:724
	_go_fuzz_dep_.CoverTab[786720] = 0
						for _, b := range *v {
//line /snap/go/10455/src/net/net.go:725
		if _go_fuzz_dep_.CoverTab[786720] == 0 {
//line /snap/go/10455/src/net/net.go:725
			_go_fuzz_dep_.CoverTab[529313]++
//line /snap/go/10455/src/net/net.go:725
		} else {
//line /snap/go/10455/src/net/net.go:725
			_go_fuzz_dep_.CoverTab[529314]++
//line /snap/go/10455/src/net/net.go:725
		}
//line /snap/go/10455/src/net/net.go:725
		_go_fuzz_dep_.CoverTab[786720] = 1
//line /snap/go/10455/src/net/net.go:725
		_go_fuzz_dep_.CoverTab[7552]++
							nb, err := w.Write(b)
							n += int64(nb)
							if err != nil {
//line /snap/go/10455/src/net/net.go:728
			_go_fuzz_dep_.CoverTab[529307]++
//line /snap/go/10455/src/net/net.go:728
			_go_fuzz_dep_.CoverTab[7553]++
								v.consume(n)
								return n, err
//line /snap/go/10455/src/net/net.go:730
			// _ = "end of CoverTab[7553]"
		} else {
//line /snap/go/10455/src/net/net.go:731
			_go_fuzz_dep_.CoverTab[529308]++
//line /snap/go/10455/src/net/net.go:731
			_go_fuzz_dep_.CoverTab[7554]++
//line /snap/go/10455/src/net/net.go:731
			// _ = "end of CoverTab[7554]"
//line /snap/go/10455/src/net/net.go:731
		}
//line /snap/go/10455/src/net/net.go:731
		// _ = "end of CoverTab[7552]"
	}
//line /snap/go/10455/src/net/net.go:732
	if _go_fuzz_dep_.CoverTab[786720] == 0 {
//line /snap/go/10455/src/net/net.go:732
		_go_fuzz_dep_.CoverTab[529315]++
//line /snap/go/10455/src/net/net.go:732
	} else {
//line /snap/go/10455/src/net/net.go:732
		_go_fuzz_dep_.CoverTab[529316]++
//line /snap/go/10455/src/net/net.go:732
	}
//line /snap/go/10455/src/net/net.go:732
	// _ = "end of CoverTab[7548]"
//line /snap/go/10455/src/net/net.go:732
	_go_fuzz_dep_.CoverTab[7549]++
						v.consume(n)
						return n, nil
//line /snap/go/10455/src/net/net.go:734
	// _ = "end of CoverTab[7549]"
}

// Read from the buffers.
//line /snap/go/10455/src/net/net.go:737
//
//line /snap/go/10455/src/net/net.go:737
// Read implements io.Reader for Buffers.
//line /snap/go/10455/src/net/net.go:737
//
//line /snap/go/10455/src/net/net.go:737
// Read modifies the slice v as well as v[i] for 0 <= i < len(v),
//line /snap/go/10455/src/net/net.go:737
// but does not modify v[i][j] for any i, j.
//line /snap/go/10455/src/net/net.go:743
func (v *Buffers) Read(p []byte) (n int, err error) {
//line /snap/go/10455/src/net/net.go:743
	_go_fuzz_dep_.CoverTab[7555]++
//line /snap/go/10455/src/net/net.go:743
	_go_fuzz_dep_.CoverTab[786721] = 0
						for len(p) > 0 && func() bool {
//line /snap/go/10455/src/net/net.go:744
		_go_fuzz_dep_.CoverTab[7558]++
//line /snap/go/10455/src/net/net.go:744
		return len(*v) > 0
//line /snap/go/10455/src/net/net.go:744
		// _ = "end of CoverTab[7558]"
//line /snap/go/10455/src/net/net.go:744
	}() {
//line /snap/go/10455/src/net/net.go:744
		if _go_fuzz_dep_.CoverTab[786721] == 0 {
//line /snap/go/10455/src/net/net.go:744
			_go_fuzz_dep_.CoverTab[529317]++
//line /snap/go/10455/src/net/net.go:744
		} else {
//line /snap/go/10455/src/net/net.go:744
			_go_fuzz_dep_.CoverTab[529318]++
//line /snap/go/10455/src/net/net.go:744
		}
//line /snap/go/10455/src/net/net.go:744
		_go_fuzz_dep_.CoverTab[786721] = 1
//line /snap/go/10455/src/net/net.go:744
		_go_fuzz_dep_.CoverTab[7559]++
							n0 := copy(p, (*v)[0])
							v.consume(int64(n0))
							p = p[n0:]
							n += n0
//line /snap/go/10455/src/net/net.go:748
		// _ = "end of CoverTab[7559]"
	}
//line /snap/go/10455/src/net/net.go:749
	if _go_fuzz_dep_.CoverTab[786721] == 0 {
//line /snap/go/10455/src/net/net.go:749
		_go_fuzz_dep_.CoverTab[529319]++
//line /snap/go/10455/src/net/net.go:749
	} else {
//line /snap/go/10455/src/net/net.go:749
		_go_fuzz_dep_.CoverTab[529320]++
//line /snap/go/10455/src/net/net.go:749
	}
//line /snap/go/10455/src/net/net.go:749
	// _ = "end of CoverTab[7555]"
//line /snap/go/10455/src/net/net.go:749
	_go_fuzz_dep_.CoverTab[7556]++
						if len(*v) == 0 {
//line /snap/go/10455/src/net/net.go:750
		_go_fuzz_dep_.CoverTab[529309]++
//line /snap/go/10455/src/net/net.go:750
		_go_fuzz_dep_.CoverTab[7560]++
							err = io.EOF
//line /snap/go/10455/src/net/net.go:751
		// _ = "end of CoverTab[7560]"
	} else {
//line /snap/go/10455/src/net/net.go:752
		_go_fuzz_dep_.CoverTab[529310]++
//line /snap/go/10455/src/net/net.go:752
		_go_fuzz_dep_.CoverTab[7561]++
//line /snap/go/10455/src/net/net.go:752
		// _ = "end of CoverTab[7561]"
//line /snap/go/10455/src/net/net.go:752
	}
//line /snap/go/10455/src/net/net.go:752
	// _ = "end of CoverTab[7556]"
//line /snap/go/10455/src/net/net.go:752
	_go_fuzz_dep_.CoverTab[7557]++
						return
//line /snap/go/10455/src/net/net.go:753
	// _ = "end of CoverTab[7557]"
}

func (v *Buffers) consume(n int64) {
//line /snap/go/10455/src/net/net.go:756
	_go_fuzz_dep_.CoverTab[7562]++
//line /snap/go/10455/src/net/net.go:756
	_go_fuzz_dep_.CoverTab[786722] = 0
						for len(*v) > 0 {
//line /snap/go/10455/src/net/net.go:757
		if _go_fuzz_dep_.CoverTab[786722] == 0 {
//line /snap/go/10455/src/net/net.go:757
			_go_fuzz_dep_.CoverTab[529321]++
//line /snap/go/10455/src/net/net.go:757
		} else {
//line /snap/go/10455/src/net/net.go:757
			_go_fuzz_dep_.CoverTab[529322]++
//line /snap/go/10455/src/net/net.go:757
		}
//line /snap/go/10455/src/net/net.go:757
		_go_fuzz_dep_.CoverTab[786722] = 1
//line /snap/go/10455/src/net/net.go:757
		_go_fuzz_dep_.CoverTab[7563]++
							ln0 := int64(len((*v)[0]))
							if ln0 > n {
//line /snap/go/10455/src/net/net.go:759
			_go_fuzz_dep_.CoverTab[529311]++
//line /snap/go/10455/src/net/net.go:759
			_go_fuzz_dep_.CoverTab[7565]++
								(*v)[0] = (*v)[0][n:]
								return
//line /snap/go/10455/src/net/net.go:761
			// _ = "end of CoverTab[7565]"
		} else {
//line /snap/go/10455/src/net/net.go:762
			_go_fuzz_dep_.CoverTab[529312]++
//line /snap/go/10455/src/net/net.go:762
			_go_fuzz_dep_.CoverTab[7566]++
//line /snap/go/10455/src/net/net.go:762
			// _ = "end of CoverTab[7566]"
//line /snap/go/10455/src/net/net.go:762
		}
//line /snap/go/10455/src/net/net.go:762
		// _ = "end of CoverTab[7563]"
//line /snap/go/10455/src/net/net.go:762
		_go_fuzz_dep_.CoverTab[7564]++
							n -= ln0
							(*v)[0] = nil
							*v = (*v)[1:]
//line /snap/go/10455/src/net/net.go:765
		// _ = "end of CoverTab[7564]"
	}
//line /snap/go/10455/src/net/net.go:766
	if _go_fuzz_dep_.CoverTab[786722] == 0 {
//line /snap/go/10455/src/net/net.go:766
		_go_fuzz_dep_.CoverTab[529323]++
//line /snap/go/10455/src/net/net.go:766
	} else {
//line /snap/go/10455/src/net/net.go:766
		_go_fuzz_dep_.CoverTab[529324]++
//line /snap/go/10455/src/net/net.go:766
	}
//line /snap/go/10455/src/net/net.go:766
	// _ = "end of CoverTab[7562]"
}

//line /snap/go/10455/src/net/net.go:767
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/net.go:767
var _ = _go_fuzz_dep_.CoverTab
