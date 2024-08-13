// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:5
package socks

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:5
)

import (
	"context"
	"errors"
	"io"
	"net"
	"strconv"
	"time"
)

var (
	noDeadline	= time.Time{}
	aLongTimeAgo	= time.Unix(1, 0)
)

func (d *Dialer) connect(ctx context.Context, c net.Conn, address string) (_ net.Addr, ctxErr error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:21
	_go_fuzz_dep_.CoverTab[96729]++
											host, port, err := splitHostPort(address)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:23
		_go_fuzz_dep_.CoverTab[96749]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:24
		// _ = "end of CoverTab[96749]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:25
		_go_fuzz_dep_.CoverTab[96750]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:25
		// _ = "end of CoverTab[96750]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:25
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:25
	// _ = "end of CoverTab[96729]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:25
	_go_fuzz_dep_.CoverTab[96730]++
											if deadline, ok := ctx.Deadline(); ok && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:26
		_go_fuzz_dep_.CoverTab[96751]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:26
		return !deadline.IsZero()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:26
		// _ = "end of CoverTab[96751]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:26
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:26
		_go_fuzz_dep_.CoverTab[96752]++
												c.SetDeadline(deadline)
												defer c.SetDeadline(noDeadline)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:28
		// _ = "end of CoverTab[96752]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:29
		_go_fuzz_dep_.CoverTab[96753]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:29
		// _ = "end of CoverTab[96753]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:29
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:29
	// _ = "end of CoverTab[96730]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:29
	_go_fuzz_dep_.CoverTab[96731]++
											if ctx != context.Background() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:30
		_go_fuzz_dep_.CoverTab[96754]++
												errCh := make(chan error, 1)
												done := make(chan struct{})
												defer func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:33
			_go_fuzz_dep_.CoverTab[96756]++
													close(done)
													if ctxErr == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:35
				_go_fuzz_dep_.CoverTab[96757]++
														ctxErr = <-errCh
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:36
				// _ = "end of CoverTab[96757]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:37
				_go_fuzz_dep_.CoverTab[96758]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:37
				// _ = "end of CoverTab[96758]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:37
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:37
			// _ = "end of CoverTab[96756]"
		}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:38
		// _ = "end of CoverTab[96754]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:38
		_go_fuzz_dep_.CoverTab[96755]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:38
		_curRoutineNum112_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:38
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum112_)
												go func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:39
			_go_fuzz_dep_.CoverTab[96759]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:39
			defer func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:39
				_go_fuzz_dep_.CoverTab[96760]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:39
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum112_)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:39
				// _ = "end of CoverTab[96760]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:39
			}()
													select {
			case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:41
				_go_fuzz_dep_.CoverTab[96761]++
														c.SetDeadline(aLongTimeAgo)
														errCh <- ctx.Err()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:43
				// _ = "end of CoverTab[96761]"
			case <-done:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:44
				_go_fuzz_dep_.CoverTab[96762]++
														errCh <- nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:45
				// _ = "end of CoverTab[96762]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:46
			// _ = "end of CoverTab[96759]"
		}()
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:47
		// _ = "end of CoverTab[96755]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:48
		_go_fuzz_dep_.CoverTab[96763]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:48
		// _ = "end of CoverTab[96763]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:48
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:48
	// _ = "end of CoverTab[96731]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:48
	_go_fuzz_dep_.CoverTab[96732]++

											b := make([]byte, 0, 6+len(host))
											b = append(b, Version5)
											if len(d.AuthMethods) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:52
		_go_fuzz_dep_.CoverTab[96764]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:52
		return d.Authenticate == nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:52
		// _ = "end of CoverTab[96764]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:52
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:52
		_go_fuzz_dep_.CoverTab[96765]++
												b = append(b, 1, byte(AuthMethodNotRequired))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:53
		// _ = "end of CoverTab[96765]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:54
		_go_fuzz_dep_.CoverTab[96766]++
												ams := d.AuthMethods
												if len(ams) > 255 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:56
			_go_fuzz_dep_.CoverTab[96768]++
													return nil, errors.New("too many authentication methods")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:57
			// _ = "end of CoverTab[96768]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:58
			_go_fuzz_dep_.CoverTab[96769]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:58
			// _ = "end of CoverTab[96769]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:58
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:58
		// _ = "end of CoverTab[96766]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:58
		_go_fuzz_dep_.CoverTab[96767]++
												b = append(b, byte(len(ams)))
												for _, am := range ams {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:60
			_go_fuzz_dep_.CoverTab[96770]++
													b = append(b, byte(am))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:61
			// _ = "end of CoverTab[96770]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:62
		// _ = "end of CoverTab[96767]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:63
	// _ = "end of CoverTab[96732]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:63
	_go_fuzz_dep_.CoverTab[96733]++
											if _, ctxErr = c.Write(b); ctxErr != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:64
		_go_fuzz_dep_.CoverTab[96771]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:65
		// _ = "end of CoverTab[96771]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:66
		_go_fuzz_dep_.CoverTab[96772]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:66
		// _ = "end of CoverTab[96772]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:66
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:66
	// _ = "end of CoverTab[96733]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:66
	_go_fuzz_dep_.CoverTab[96734]++

											if _, ctxErr = io.ReadFull(c, b[:2]); ctxErr != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:68
		_go_fuzz_dep_.CoverTab[96773]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:69
		// _ = "end of CoverTab[96773]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:70
		_go_fuzz_dep_.CoverTab[96774]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:70
		// _ = "end of CoverTab[96774]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:70
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:70
	// _ = "end of CoverTab[96734]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:70
	_go_fuzz_dep_.CoverTab[96735]++
											if b[0] != Version5 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:71
		_go_fuzz_dep_.CoverTab[96775]++
												return nil, errors.New("unexpected protocol version " + strconv.Itoa(int(b[0])))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:72
		// _ = "end of CoverTab[96775]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:73
		_go_fuzz_dep_.CoverTab[96776]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:73
		// _ = "end of CoverTab[96776]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:73
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:73
	// _ = "end of CoverTab[96735]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:73
	_go_fuzz_dep_.CoverTab[96736]++
											am := AuthMethod(b[1])
											if am == AuthMethodNoAcceptableMethods {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:75
		_go_fuzz_dep_.CoverTab[96777]++
												return nil, errors.New("no acceptable authentication methods")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:76
		// _ = "end of CoverTab[96777]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:77
		_go_fuzz_dep_.CoverTab[96778]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:77
		// _ = "end of CoverTab[96778]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:77
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:77
	// _ = "end of CoverTab[96736]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:77
	_go_fuzz_dep_.CoverTab[96737]++
											if d.Authenticate != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:78
		_go_fuzz_dep_.CoverTab[96779]++
												if ctxErr = d.Authenticate(ctx, c, am); ctxErr != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:79
			_go_fuzz_dep_.CoverTab[96780]++
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:80
			// _ = "end of CoverTab[96780]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:81
			_go_fuzz_dep_.CoverTab[96781]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:81
			// _ = "end of CoverTab[96781]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:81
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:81
		// _ = "end of CoverTab[96779]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:82
		_go_fuzz_dep_.CoverTab[96782]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:82
		// _ = "end of CoverTab[96782]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:82
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:82
	// _ = "end of CoverTab[96737]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:82
	_go_fuzz_dep_.CoverTab[96738]++

											b = b[:0]
											b = append(b, Version5, byte(d.cmd), 0)
											if ip := net.ParseIP(host); ip != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:86
		_go_fuzz_dep_.CoverTab[96783]++
												if ip4 := ip.To4(); ip4 != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:87
			_go_fuzz_dep_.CoverTab[96784]++
													b = append(b, AddrTypeIPv4)
													b = append(b, ip4...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:89
			// _ = "end of CoverTab[96784]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:90
			_go_fuzz_dep_.CoverTab[96785]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:90
			if ip6 := ip.To16(); ip6 != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:90
				_go_fuzz_dep_.CoverTab[96786]++
														b = append(b, AddrTypeIPv6)
														b = append(b, ip6...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:92
				// _ = "end of CoverTab[96786]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:93
				_go_fuzz_dep_.CoverTab[96787]++
														return nil, errors.New("unknown address type")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:94
				// _ = "end of CoverTab[96787]"
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:95
			// _ = "end of CoverTab[96785]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:95
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:95
		// _ = "end of CoverTab[96783]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:96
		_go_fuzz_dep_.CoverTab[96788]++
												if len(host) > 255 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:97
			_go_fuzz_dep_.CoverTab[96790]++
													return nil, errors.New("FQDN too long")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:98
			// _ = "end of CoverTab[96790]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:99
			_go_fuzz_dep_.CoverTab[96791]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:99
			// _ = "end of CoverTab[96791]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:99
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:99
			// _ = "end of CoverTab[96788]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:99
			_go_fuzz_dep_.CoverTab[96789]++
													b = append(b, AddrTypeFQDN)
													b = append(b, byte(len(host)))
													b = append(b, host...)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:102
		// _ = "end of CoverTab[96789]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:103
	// _ = "end of CoverTab[96738]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:103
	_go_fuzz_dep_.CoverTab[96739]++
												b = append(b, byte(port>>8), byte(port))
												if _, ctxErr = c.Write(b); ctxErr != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:105
		_go_fuzz_dep_.CoverTab[96792]++
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:106
		// _ = "end of CoverTab[96792]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:107
		_go_fuzz_dep_.CoverTab[96793]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:107
		// _ = "end of CoverTab[96793]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:107
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:107
	// _ = "end of CoverTab[96739]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:107
	_go_fuzz_dep_.CoverTab[96740]++

												if _, ctxErr = io.ReadFull(c, b[:4]); ctxErr != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:109
		_go_fuzz_dep_.CoverTab[96794]++
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:110
		// _ = "end of CoverTab[96794]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:111
		_go_fuzz_dep_.CoverTab[96795]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:111
		// _ = "end of CoverTab[96795]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:111
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:111
	// _ = "end of CoverTab[96740]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:111
	_go_fuzz_dep_.CoverTab[96741]++
												if b[0] != Version5 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:112
		_go_fuzz_dep_.CoverTab[96796]++
													return nil, errors.New("unexpected protocol version " + strconv.Itoa(int(b[0])))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:113
		// _ = "end of CoverTab[96796]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:114
		_go_fuzz_dep_.CoverTab[96797]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:114
		// _ = "end of CoverTab[96797]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:114
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:114
	// _ = "end of CoverTab[96741]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:114
	_go_fuzz_dep_.CoverTab[96742]++
												if cmdErr := Reply(b[1]); cmdErr != StatusSucceeded {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:115
		_go_fuzz_dep_.CoverTab[96798]++
													return nil, errors.New("unknown error " + cmdErr.String())
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:116
		// _ = "end of CoverTab[96798]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:117
		_go_fuzz_dep_.CoverTab[96799]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:117
		// _ = "end of CoverTab[96799]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:117
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:117
	// _ = "end of CoverTab[96742]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:117
	_go_fuzz_dep_.CoverTab[96743]++
												if b[2] != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:118
		_go_fuzz_dep_.CoverTab[96800]++
													return nil, errors.New("non-zero reserved field")
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:119
		// _ = "end of CoverTab[96800]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:120
		_go_fuzz_dep_.CoverTab[96801]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:120
		// _ = "end of CoverTab[96801]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:120
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:120
	// _ = "end of CoverTab[96743]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:120
	_go_fuzz_dep_.CoverTab[96744]++
												l := 2
												var a Addr
												switch b[3] {
	case AddrTypeIPv4:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:124
		_go_fuzz_dep_.CoverTab[96802]++
													l += net.IPv4len
													a.IP = make(net.IP, net.IPv4len)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:126
		// _ = "end of CoverTab[96802]"
	case AddrTypeIPv6:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:127
		_go_fuzz_dep_.CoverTab[96803]++
													l += net.IPv6len
													a.IP = make(net.IP, net.IPv6len)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:129
		// _ = "end of CoverTab[96803]"
	case AddrTypeFQDN:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:130
		_go_fuzz_dep_.CoverTab[96804]++
													if _, err := io.ReadFull(c, b[:1]); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:131
			_go_fuzz_dep_.CoverTab[96807]++
														return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:132
			// _ = "end of CoverTab[96807]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:133
			_go_fuzz_dep_.CoverTab[96808]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:133
			// _ = "end of CoverTab[96808]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:133
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:133
		// _ = "end of CoverTab[96804]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:133
		_go_fuzz_dep_.CoverTab[96805]++
													l += int(b[0])
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:134
		// _ = "end of CoverTab[96805]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:135
		_go_fuzz_dep_.CoverTab[96806]++
													return nil, errors.New("unknown address type " + strconv.Itoa(int(b[3])))
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:136
		// _ = "end of CoverTab[96806]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:137
	// _ = "end of CoverTab[96744]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:137
	_go_fuzz_dep_.CoverTab[96745]++
												if cap(b) < l {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:138
		_go_fuzz_dep_.CoverTab[96809]++
													b = make([]byte, l)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:139
		// _ = "end of CoverTab[96809]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:140
		_go_fuzz_dep_.CoverTab[96810]++
													b = b[:l]
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:141
		// _ = "end of CoverTab[96810]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:142
	// _ = "end of CoverTab[96745]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:142
	_go_fuzz_dep_.CoverTab[96746]++
												if _, ctxErr = io.ReadFull(c, b); ctxErr != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:143
		_go_fuzz_dep_.CoverTab[96811]++
													return
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:144
		// _ = "end of CoverTab[96811]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:145
		_go_fuzz_dep_.CoverTab[96812]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:145
		// _ = "end of CoverTab[96812]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:145
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:145
	// _ = "end of CoverTab[96746]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:145
	_go_fuzz_dep_.CoverTab[96747]++
												if a.IP != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:146
		_go_fuzz_dep_.CoverTab[96813]++
													copy(a.IP, b)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:147
		// _ = "end of CoverTab[96813]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:148
		_go_fuzz_dep_.CoverTab[96814]++
													a.Name = string(b[:len(b)-2])
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:149
		// _ = "end of CoverTab[96814]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:150
	// _ = "end of CoverTab[96747]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:150
	_go_fuzz_dep_.CoverTab[96748]++
												a.Port = int(b[len(b)-2])<<8 | int(b[len(b)-1])
												return &a, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:152
	// _ = "end of CoverTab[96748]"
}

func splitHostPort(address string) (string, int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:155
	_go_fuzz_dep_.CoverTab[96815]++
												host, port, err := net.SplitHostPort(address)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:157
		_go_fuzz_dep_.CoverTab[96819]++
													return "", 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:158
		// _ = "end of CoverTab[96819]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:159
		_go_fuzz_dep_.CoverTab[96820]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:159
		// _ = "end of CoverTab[96820]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:159
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:159
	// _ = "end of CoverTab[96815]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:159
	_go_fuzz_dep_.CoverTab[96816]++
												portnum, err := strconv.Atoi(port)
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:161
		_go_fuzz_dep_.CoverTab[96821]++
													return "", 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:162
		// _ = "end of CoverTab[96821]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:163
		_go_fuzz_dep_.CoverTab[96822]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:163
		// _ = "end of CoverTab[96822]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:163
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:163
	// _ = "end of CoverTab[96816]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:163
	_go_fuzz_dep_.CoverTab[96817]++
												if 1 > portnum || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:164
		_go_fuzz_dep_.CoverTab[96823]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:164
		return portnum > 0xffff
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:164
		// _ = "end of CoverTab[96823]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:164
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:164
		_go_fuzz_dep_.CoverTab[96824]++
													return "", 0, errors.New("port number out of range " + port)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:165
		// _ = "end of CoverTab[96824]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:166
		_go_fuzz_dep_.CoverTab[96825]++
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:166
		// _ = "end of CoverTab[96825]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:166
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:166
	// _ = "end of CoverTab[96817]"
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:166
	_go_fuzz_dep_.CoverTab[96818]++
												return host, portnum, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:167
	// _ = "end of CoverTab[96818]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:168
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/net@v0.10.0/internal/socks/client.go:168
var _ = _go_fuzz_dep_.CoverTab
