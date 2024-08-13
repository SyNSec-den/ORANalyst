// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:5
// Package socks provides a SOCKS version 5 client implementation.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:5
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:5
// SOCKS protocol version 5 is defined in RFC 1928.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:5
// Username/Password authentication for SOCKS version 5 is defined in
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:5
// RFC 1929.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:10
package socks

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:10
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:10
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:10
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:10
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:10
)

import (
	"context"
	"errors"
	"io"
	"net"
	"strconv"
)

// A Command represents a SOCKS command.
type Command int

func (cmd Command) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:23
	_go_fuzz_dep_.CoverTab[96826]++
											switch cmd {
	case CmdConnect:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:25
		_go_fuzz_dep_.CoverTab[96827]++
												return "socks connect"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:26
		// _ = "end of CoverTab[96827]"
	case cmdBind:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:27
		_go_fuzz_dep_.CoverTab[96828]++
												return "socks bind"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:28
		// _ = "end of CoverTab[96828]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:29
		_go_fuzz_dep_.CoverTab[96829]++
												return "socks " + strconv.Itoa(int(cmd))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:30
		// _ = "end of CoverTab[96829]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:31
	// _ = "end of CoverTab[96826]"
}

// An AuthMethod represents a SOCKS authentication method.
type AuthMethod int

// A Reply represents a SOCKS command reply code.
type Reply int

func (code Reply) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:40
	_go_fuzz_dep_.CoverTab[96830]++
											switch code {
	case StatusSucceeded:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:42
		_go_fuzz_dep_.CoverTab[96831]++
												return "succeeded"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:43
		// _ = "end of CoverTab[96831]"
	case 0x01:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:44
		_go_fuzz_dep_.CoverTab[96832]++
												return "general SOCKS server failure"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:45
		// _ = "end of CoverTab[96832]"
	case 0x02:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:46
		_go_fuzz_dep_.CoverTab[96833]++
												return "connection not allowed by ruleset"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:47
		// _ = "end of CoverTab[96833]"
	case 0x03:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:48
		_go_fuzz_dep_.CoverTab[96834]++
												return "network unreachable"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:49
		// _ = "end of CoverTab[96834]"
	case 0x04:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:50
		_go_fuzz_dep_.CoverTab[96835]++
												return "host unreachable"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:51
		// _ = "end of CoverTab[96835]"
	case 0x05:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:52
		_go_fuzz_dep_.CoverTab[96836]++
												return "connection refused"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:53
		// _ = "end of CoverTab[96836]"
	case 0x06:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:54
		_go_fuzz_dep_.CoverTab[96837]++
												return "TTL expired"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:55
		// _ = "end of CoverTab[96837]"
	case 0x07:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:56
		_go_fuzz_dep_.CoverTab[96838]++
												return "command not supported"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:57
		// _ = "end of CoverTab[96838]"
	case 0x08:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:58
		_go_fuzz_dep_.CoverTab[96839]++
												return "address type not supported"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:59
		// _ = "end of CoverTab[96839]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:60
		_go_fuzz_dep_.CoverTab[96840]++
												return "unknown code: " + strconv.Itoa(int(code))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:61
		// _ = "end of CoverTab[96840]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:62
	// _ = "end of CoverTab[96830]"
}

// Wire protocol constants.
const (
	Version5	= 0x05

	AddrTypeIPv4	= 0x01
	AddrTypeFQDN	= 0x03
	AddrTypeIPv6	= 0x04

	CmdConnect	Command	= 0x01	// establishes an active-open forward proxy connection
	cmdBind		Command	= 0x02	// establishes a passive-open forward proxy connection

	AuthMethodNotRequired		AuthMethod	= 0x00	// no authentication required
	AuthMethodUsernamePassword	AuthMethod	= 0x02	// use username/password
	AuthMethodNoAcceptableMethods	AuthMethod	= 0xff	// no acceptable authentication methods

	StatusSucceeded	Reply	= 0x00
)

// An Addr represents a SOCKS-specific address.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:83
// Either Name or IP is used exclusively.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:85
type Addr struct {
	Name	string	// fully-qualified domain name
	IP	net.IP
	Port	int
}

func (a *Addr) Network() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:91
	_go_fuzz_dep_.CoverTab[96841]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:91
	return "socks"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:91
	// _ = "end of CoverTab[96841]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:91
}

func (a *Addr) String() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:93
	_go_fuzz_dep_.CoverTab[96842]++
											if a == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:94
		_go_fuzz_dep_.CoverTab[96845]++
												return "<nil>"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:95
		// _ = "end of CoverTab[96845]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:96
		_go_fuzz_dep_.CoverTab[96846]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:96
		// _ = "end of CoverTab[96846]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:96
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:96
	// _ = "end of CoverTab[96842]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:96
	_go_fuzz_dep_.CoverTab[96843]++
											port := strconv.Itoa(a.Port)
											if a.IP == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:98
		_go_fuzz_dep_.CoverTab[96847]++
												return net.JoinHostPort(a.Name, port)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:99
		// _ = "end of CoverTab[96847]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:100
		_go_fuzz_dep_.CoverTab[96848]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:100
		// _ = "end of CoverTab[96848]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:100
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:100
	// _ = "end of CoverTab[96843]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:100
	_go_fuzz_dep_.CoverTab[96844]++
											return net.JoinHostPort(a.IP.String(), port)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:101
	// _ = "end of CoverTab[96844]"
}

// A Conn represents a forward proxy connection.
type Conn struct {
	net.Conn

	boundAddr	net.Addr
}

// BoundAddr returns the address assigned by the proxy server for
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:111
// connecting to the command target address from the proxy server.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:113
func (c *Conn) BoundAddr() net.Addr {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:113
	_go_fuzz_dep_.CoverTab[96849]++
											if c == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:114
		_go_fuzz_dep_.CoverTab[96851]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:115
		// _ = "end of CoverTab[96851]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:116
		_go_fuzz_dep_.CoverTab[96852]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:116
		// _ = "end of CoverTab[96852]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:116
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:116
	// _ = "end of CoverTab[96849]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:116
	_go_fuzz_dep_.CoverTab[96850]++
											return c.boundAddr
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:117
	// _ = "end of CoverTab[96850]"
}

// A Dialer holds SOCKS-specific options.
type Dialer struct {
	cmd		Command	// either CmdConnect or cmdBind
	proxyNetwork	string	// network between a proxy server and a client
	proxyAddress	string	// proxy server address

	// ProxyDial specifies the optional dial function for
	// establishing the transport connection.
	ProxyDial	func(context.Context, string, string) (net.Conn, error)

	// AuthMethods specifies the list of request authentication
	// methods.
	// If empty, SOCKS client requests only AuthMethodNotRequired.
	AuthMethods	[]AuthMethod

	// Authenticate specifies the optional authentication
	// function. It must be non-nil when AuthMethods is not empty.
	// It must return an error when the authentication is failed.
	Authenticate	func(context.Context, io.ReadWriter, AuthMethod) error
}

// DialContext connects to the provided address on the provided
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:141
// network.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:141
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:141
// The returned error value may be a net.OpError. When the Op field of
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:141
// net.OpError contains "socks", the Source field contains a proxy
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:141
// server address and the Addr field contains a command target
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:141
// address.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:141
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:141
// See func Dial of the net package of standard library for a
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:141
// description of the network and address parameters.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:151
func (d *Dialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:151
	_go_fuzz_dep_.CoverTab[96853]++
											if err := d.validateTarget(network, address); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:152
		_go_fuzz_dep_.CoverTab[96859]++
												proxy, dst, _ := d.pathAddrs(address)
												return nil, &net.OpError{Op: d.cmd.String(), Net: network, Source: proxy, Addr: dst, Err: err}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:154
		// _ = "end of CoverTab[96859]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:155
		_go_fuzz_dep_.CoverTab[96860]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:155
		// _ = "end of CoverTab[96860]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:155
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:155
	// _ = "end of CoverTab[96853]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:155
	_go_fuzz_dep_.CoverTab[96854]++
											if ctx == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:156
		_go_fuzz_dep_.CoverTab[96861]++
												proxy, dst, _ := d.pathAddrs(address)
												return nil, &net.OpError{Op: d.cmd.String(), Net: network, Source: proxy, Addr: dst, Err: errors.New("nil context")}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:158
		// _ = "end of CoverTab[96861]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:159
		_go_fuzz_dep_.CoverTab[96862]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:159
		// _ = "end of CoverTab[96862]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:159
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:159
	// _ = "end of CoverTab[96854]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:159
	_go_fuzz_dep_.CoverTab[96855]++
											var err error
											var c net.Conn
											if d.ProxyDial != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:162
		_go_fuzz_dep_.CoverTab[96863]++
												c, err = d.ProxyDial(ctx, d.proxyNetwork, d.proxyAddress)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:163
		// _ = "end of CoverTab[96863]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:164
		_go_fuzz_dep_.CoverTab[96864]++
												var dd net.Dialer
												c, err = dd.DialContext(ctx, d.proxyNetwork, d.proxyAddress)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:166
		// _ = "end of CoverTab[96864]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:167
	// _ = "end of CoverTab[96855]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:167
	_go_fuzz_dep_.CoverTab[96856]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:168
		_go_fuzz_dep_.CoverTab[96865]++
												proxy, dst, _ := d.pathAddrs(address)
												return nil, &net.OpError{Op: d.cmd.String(), Net: network, Source: proxy, Addr: dst, Err: err}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:170
		// _ = "end of CoverTab[96865]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:171
		_go_fuzz_dep_.CoverTab[96866]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:171
		// _ = "end of CoverTab[96866]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:171
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:171
	// _ = "end of CoverTab[96856]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:171
	_go_fuzz_dep_.CoverTab[96857]++
											a, err := d.connect(ctx, c, address)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:173
		_go_fuzz_dep_.CoverTab[96867]++
												c.Close()
												proxy, dst, _ := d.pathAddrs(address)
												return nil, &net.OpError{Op: d.cmd.String(), Net: network, Source: proxy, Addr: dst, Err: err}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:176
		// _ = "end of CoverTab[96867]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:177
		_go_fuzz_dep_.CoverTab[96868]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:177
		// _ = "end of CoverTab[96868]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:177
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:177
	// _ = "end of CoverTab[96857]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:177
	_go_fuzz_dep_.CoverTab[96858]++
											return &Conn{Conn: c, boundAddr: a}, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:178
	// _ = "end of CoverTab[96858]"
}

// DialWithConn initiates a connection from SOCKS server to the target
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:181
// network and address using the connection c that is already
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:181
// connected to the SOCKS server.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:181
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:181
// It returns the connection's local address assigned by the SOCKS
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:181
// server.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:187
func (d *Dialer) DialWithConn(ctx context.Context, c net.Conn, network, address string) (net.Addr, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:187
	_go_fuzz_dep_.CoverTab[96869]++
											if err := d.validateTarget(network, address); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:188
		_go_fuzz_dep_.CoverTab[96873]++
												proxy, dst, _ := d.pathAddrs(address)
												return nil, &net.OpError{Op: d.cmd.String(), Net: network, Source: proxy, Addr: dst, Err: err}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:190
		// _ = "end of CoverTab[96873]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:191
		_go_fuzz_dep_.CoverTab[96874]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:191
		// _ = "end of CoverTab[96874]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:191
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:191
	// _ = "end of CoverTab[96869]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:191
	_go_fuzz_dep_.CoverTab[96870]++
											if ctx == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:192
		_go_fuzz_dep_.CoverTab[96875]++
												proxy, dst, _ := d.pathAddrs(address)
												return nil, &net.OpError{Op: d.cmd.String(), Net: network, Source: proxy, Addr: dst, Err: errors.New("nil context")}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:194
		// _ = "end of CoverTab[96875]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:195
		_go_fuzz_dep_.CoverTab[96876]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:195
		// _ = "end of CoverTab[96876]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:195
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:195
	// _ = "end of CoverTab[96870]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:195
	_go_fuzz_dep_.CoverTab[96871]++
											a, err := d.connect(ctx, c, address)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:197
		_go_fuzz_dep_.CoverTab[96877]++
												proxy, dst, _ := d.pathAddrs(address)
												return nil, &net.OpError{Op: d.cmd.String(), Net: network, Source: proxy, Addr: dst, Err: err}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:199
		// _ = "end of CoverTab[96877]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:200
		_go_fuzz_dep_.CoverTab[96878]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:200
		// _ = "end of CoverTab[96878]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:200
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:200
	// _ = "end of CoverTab[96871]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:200
	_go_fuzz_dep_.CoverTab[96872]++
											return a, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:201
	// _ = "end of CoverTab[96872]"
}

// Dial connects to the provided address on the provided network.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:204
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:204
// Unlike DialContext, it returns a raw transport connection instead
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:204
// of a forward proxy connection.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:204
//
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:204
// Deprecated: Use DialContext or DialWithConn instead.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:210
func (d *Dialer) Dial(network, address string) (net.Conn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:210
	_go_fuzz_dep_.CoverTab[96879]++
											if err := d.validateTarget(network, address); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:211
		_go_fuzz_dep_.CoverTab[96884]++
												proxy, dst, _ := d.pathAddrs(address)
												return nil, &net.OpError{Op: d.cmd.String(), Net: network, Source: proxy, Addr: dst, Err: err}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:213
		// _ = "end of CoverTab[96884]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:214
		_go_fuzz_dep_.CoverTab[96885]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:214
		// _ = "end of CoverTab[96885]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:214
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:214
	// _ = "end of CoverTab[96879]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:214
	_go_fuzz_dep_.CoverTab[96880]++
											var err error
											var c net.Conn
											if d.ProxyDial != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:217
		_go_fuzz_dep_.CoverTab[96886]++
												c, err = d.ProxyDial(context.Background(), d.proxyNetwork, d.proxyAddress)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:218
		// _ = "end of CoverTab[96886]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:219
		_go_fuzz_dep_.CoverTab[96887]++
												c, err = net.Dial(d.proxyNetwork, d.proxyAddress)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:220
		// _ = "end of CoverTab[96887]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:221
	// _ = "end of CoverTab[96880]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:221
	_go_fuzz_dep_.CoverTab[96881]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:222
		_go_fuzz_dep_.CoverTab[96888]++
												proxy, dst, _ := d.pathAddrs(address)
												return nil, &net.OpError{Op: d.cmd.String(), Net: network, Source: proxy, Addr: dst, Err: err}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:224
		// _ = "end of CoverTab[96888]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:225
		_go_fuzz_dep_.CoverTab[96889]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:225
		// _ = "end of CoverTab[96889]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:225
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:225
	// _ = "end of CoverTab[96881]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:225
	_go_fuzz_dep_.CoverTab[96882]++
											if _, err := d.DialWithConn(context.Background(), c, network, address); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:226
		_go_fuzz_dep_.CoverTab[96890]++
												c.Close()
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:228
		// _ = "end of CoverTab[96890]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:229
		_go_fuzz_dep_.CoverTab[96891]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:229
		// _ = "end of CoverTab[96891]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:229
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:229
	// _ = "end of CoverTab[96882]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:229
	_go_fuzz_dep_.CoverTab[96883]++
											return c, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:230
	// _ = "end of CoverTab[96883]"
}

func (d *Dialer) validateTarget(network, address string) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:233
	_go_fuzz_dep_.CoverTab[96892]++
											switch network {
	case "tcp", "tcp6", "tcp4":
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:235
		_go_fuzz_dep_.CoverTab[96895]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:235
		// _ = "end of CoverTab[96895]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:236
		_go_fuzz_dep_.CoverTab[96896]++
												return errors.New("network not implemented")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:237
		// _ = "end of CoverTab[96896]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:238
	// _ = "end of CoverTab[96892]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:238
	_go_fuzz_dep_.CoverTab[96893]++
											switch d.cmd {
	case CmdConnect, cmdBind:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:240
		_go_fuzz_dep_.CoverTab[96897]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:240
		// _ = "end of CoverTab[96897]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:241
		_go_fuzz_dep_.CoverTab[96898]++
												return errors.New("command not implemented")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:242
		// _ = "end of CoverTab[96898]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:243
	// _ = "end of CoverTab[96893]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:243
	_go_fuzz_dep_.CoverTab[96894]++
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:244
	// _ = "end of CoverTab[96894]"
}

func (d *Dialer) pathAddrs(address string) (proxy, dst net.Addr, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:247
	_go_fuzz_dep_.CoverTab[96899]++
											for i, s := range []string{d.proxyAddress, address} {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:248
		_go_fuzz_dep_.CoverTab[96901]++
												host, port, err := splitHostPort(s)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:250
			_go_fuzz_dep_.CoverTab[96904]++
													return nil, nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:251
			// _ = "end of CoverTab[96904]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:252
			_go_fuzz_dep_.CoverTab[96905]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:252
			// _ = "end of CoverTab[96905]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:252
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:252
		// _ = "end of CoverTab[96901]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:252
		_go_fuzz_dep_.CoverTab[96902]++
												a := &Addr{Port: port}
												a.IP = net.ParseIP(host)
												if a.IP == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:255
			_go_fuzz_dep_.CoverTab[96906]++
													a.Name = host
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:256
			// _ = "end of CoverTab[96906]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:257
			_go_fuzz_dep_.CoverTab[96907]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:257
			// _ = "end of CoverTab[96907]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:257
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:257
		// _ = "end of CoverTab[96902]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:257
		_go_fuzz_dep_.CoverTab[96903]++
												if i == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:258
			_go_fuzz_dep_.CoverTab[96908]++
													proxy = a
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:259
			// _ = "end of CoverTab[96908]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:260
			_go_fuzz_dep_.CoverTab[96909]++
													dst = a
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:261
			// _ = "end of CoverTab[96909]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:262
		// _ = "end of CoverTab[96903]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:263
	// _ = "end of CoverTab[96899]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:263
	_go_fuzz_dep_.CoverTab[96900]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:264
	// _ = "end of CoverTab[96900]"
}

// NewDialer returns a new Dialer that dials through the provided
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:267
// proxy server's network and address.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:269
func NewDialer(network, address string) *Dialer {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:269
	_go_fuzz_dep_.CoverTab[96910]++
											return &Dialer{proxyNetwork: network, proxyAddress: address, cmd: CmdConnect}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:270
	// _ = "end of CoverTab[96910]"
}

const (
	authUsernamePasswordVersion	= 0x01
	authStatusSucceeded		= 0x00
)

// UsernamePassword are the credentials for the username/password
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:278
// authentication method.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:280
type UsernamePassword struct {
	Username	string
	Password	string
}

// Authenticate authenticates a pair of username and password with the
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:285
// proxy server.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:287
func (up *UsernamePassword) Authenticate(ctx context.Context, rw io.ReadWriter, auth AuthMethod) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:287
	_go_fuzz_dep_.CoverTab[96911]++
											switch auth {
	case AuthMethodNotRequired:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:289
		_go_fuzz_dep_.CoverTab[96913]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:290
		// _ = "end of CoverTab[96913]"
	case AuthMethodUsernamePassword:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:291
		_go_fuzz_dep_.CoverTab[96914]++
												if len(up.Username) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:292
			_go_fuzz_dep_.CoverTab[96921]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:292
			return len(up.Username) > 255
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:292
			// _ = "end of CoverTab[96921]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:292
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:292
			_go_fuzz_dep_.CoverTab[96922]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:292
			return len(up.Password) > 255
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:292
			// _ = "end of CoverTab[96922]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:292
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:292
			_go_fuzz_dep_.CoverTab[96923]++
													return errors.New("invalid username/password")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:293
			// _ = "end of CoverTab[96923]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:294
			_go_fuzz_dep_.CoverTab[96924]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:294
			// _ = "end of CoverTab[96924]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:294
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:294
		// _ = "end of CoverTab[96914]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:294
		_go_fuzz_dep_.CoverTab[96915]++
												b := []byte{authUsernamePasswordVersion}
												b = append(b, byte(len(up.Username)))
												b = append(b, up.Username...)
												b = append(b, byte(len(up.Password)))
												b = append(b, up.Password...)

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:302
		if _, err := rw.Write(b); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:302
			_go_fuzz_dep_.CoverTab[96925]++
													return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:303
			// _ = "end of CoverTab[96925]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:304
			_go_fuzz_dep_.CoverTab[96926]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:304
			// _ = "end of CoverTab[96926]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:304
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:304
		// _ = "end of CoverTab[96915]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:304
		_go_fuzz_dep_.CoverTab[96916]++
												if _, err := io.ReadFull(rw, b[:2]); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:305
			_go_fuzz_dep_.CoverTab[96927]++
													return err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:306
			// _ = "end of CoverTab[96927]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:307
			_go_fuzz_dep_.CoverTab[96928]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:307
			// _ = "end of CoverTab[96928]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:307
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:307
		// _ = "end of CoverTab[96916]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:307
		_go_fuzz_dep_.CoverTab[96917]++
												if b[0] != authUsernamePasswordVersion {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:308
			_go_fuzz_dep_.CoverTab[96929]++
													return errors.New("invalid username/password version")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:309
			// _ = "end of CoverTab[96929]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:310
			_go_fuzz_dep_.CoverTab[96930]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:310
			// _ = "end of CoverTab[96930]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:310
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:310
		// _ = "end of CoverTab[96917]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:310
		_go_fuzz_dep_.CoverTab[96918]++
												if b[1] != authStatusSucceeded {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:311
			_go_fuzz_dep_.CoverTab[96931]++
													return errors.New("username/password authentication failed")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:312
			// _ = "end of CoverTab[96931]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:313
			_go_fuzz_dep_.CoverTab[96932]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:313
			// _ = "end of CoverTab[96932]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:313
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:313
		// _ = "end of CoverTab[96918]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:313
		_go_fuzz_dep_.CoverTab[96919]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:314
		// _ = "end of CoverTab[96919]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:314
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:314
		_go_fuzz_dep_.CoverTab[96920]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:314
		// _ = "end of CoverTab[96920]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:315
	// _ = "end of CoverTab[96911]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:315
	_go_fuzz_dep_.CoverTab[96912]++
											return errors.New("unsupported authentication method " + strconv.Itoa(int(auth)))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:316
	// _ = "end of CoverTab[96912]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:317
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/socks.go:317
var _ = _go_fuzz_dep_.CoverTab
