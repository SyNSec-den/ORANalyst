// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:5
package proxy

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:5
)

import (
	"context"
	"net"

	"golang.org/x/net/internal/socks"
)

// SOCKS5 returns a Dialer that makes SOCKSv5 connections to the given
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:14
// address with an optional username and password.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:14
// See RFC 1928 and RFC 1929.
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:17
func SOCKS5(network, address string, auth *Auth, forward Dialer) (Dialer, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:17
	_go_fuzz_dep_.CoverTab[97051]++
										d := socks.NewDialer(network, address)
										if forward != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:19
		_go_fuzz_dep_.CoverTab[97054]++
											if f, ok := forward.(ContextDialer); ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:20
			_go_fuzz_dep_.CoverTab[97055]++
												d.ProxyDial = func(ctx context.Context, network string, address string) (net.Conn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:21
				_go_fuzz_dep_.CoverTab[97056]++
													return f.DialContext(ctx, network, address)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:22
				// _ = "end of CoverTab[97056]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:23
			// _ = "end of CoverTab[97055]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:24
			_go_fuzz_dep_.CoverTab[97057]++
												d.ProxyDial = func(ctx context.Context, network string, address string) (net.Conn, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:25
				_go_fuzz_dep_.CoverTab[97058]++
													return dialContext(ctx, forward, network, address)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:26
				// _ = "end of CoverTab[97058]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:27
			// _ = "end of CoverTab[97057]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:28
		// _ = "end of CoverTab[97054]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:29
		_go_fuzz_dep_.CoverTab[97059]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:29
		// _ = "end of CoverTab[97059]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:29
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:29
	// _ = "end of CoverTab[97051]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:29
	_go_fuzz_dep_.CoverTab[97052]++
										if auth != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:30
		_go_fuzz_dep_.CoverTab[97060]++
											up := socks.UsernamePassword{
			Username:	auth.User,
			Password:	auth.Password,
		}
		d.AuthMethods = []socks.AuthMethod{
			socks.AuthMethodNotRequired,
			socks.AuthMethodUsernamePassword,
		}
											d.Authenticate = up.Authenticate
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:39
		// _ = "end of CoverTab[97060]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:40
		_go_fuzz_dep_.CoverTab[97061]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:40
		// _ = "end of CoverTab[97061]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:40
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:40
	// _ = "end of CoverTab[97052]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:40
	_go_fuzz_dep_.CoverTab[97053]++
										return d, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:41
	// _ = "end of CoverTab[97053]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:42
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/proxy/socks5.go:42
var _ = _go_fuzz_dep_.CoverTab
