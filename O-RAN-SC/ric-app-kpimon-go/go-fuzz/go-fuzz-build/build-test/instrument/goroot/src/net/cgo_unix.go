// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is called cgo_unix.go, but to allow syscalls-to-libc-based
// implementations to share the code, it does not use cgo directly.
// Instead of C.foo it uses _C_foo, which is defined in either
// cgo_unix_cgo.go or cgo_unix_syscall.go

//go:build !netgo && ((cgo && unix) || darwin)

//line /snap/go/10455/src/net/cgo_unix.go:12
package net

//line /snap/go/10455/src/net/cgo_unix.go:12
import (
//line /snap/go/10455/src/net/cgo_unix.go:12
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/cgo_unix.go:12
)
//line /snap/go/10455/src/net/cgo_unix.go:12
import (
//line /snap/go/10455/src/net/cgo_unix.go:12
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/cgo_unix.go:12
)

import (
	"context"
	"errors"
	"net/netip"
	"syscall"
	"unsafe"

	"golang.org/x/net/dns/dnsmessage"
)

// cgoAvailable set to true to indicate that the cgo resolver
//line /snap/go/10455/src/net/cgo_unix.go:24
// is available on this system.
//line /snap/go/10455/src/net/cgo_unix.go:26
const cgoAvailable = true

// An addrinfoErrno represents a getaddrinfo, getnameinfo-specific
//line /snap/go/10455/src/net/cgo_unix.go:28
// error number. It's a signed number and a zero value is a non-error
//line /snap/go/10455/src/net/cgo_unix.go:28
// by convention.
//line /snap/go/10455/src/net/cgo_unix.go:31
type addrinfoErrno int

func (eai addrinfoErrno) Error() string {
//line /snap/go/10455/src/net/cgo_unix.go:33
	_go_fuzz_dep_.CoverTab[4597]++
//line /snap/go/10455/src/net/cgo_unix.go:33
	return _C_gai_strerror(_C_int(eai))
//line /snap/go/10455/src/net/cgo_unix.go:33
	// _ = "end of CoverTab[4597]"
//line /snap/go/10455/src/net/cgo_unix.go:33
}
func (eai addrinfoErrno) Temporary() bool {
//line /snap/go/10455/src/net/cgo_unix.go:34
	_go_fuzz_dep_.CoverTab[4598]++
//line /snap/go/10455/src/net/cgo_unix.go:34
	return eai == _C_EAI_AGAIN
//line /snap/go/10455/src/net/cgo_unix.go:34
	// _ = "end of CoverTab[4598]"
//line /snap/go/10455/src/net/cgo_unix.go:34
}
func (eai addrinfoErrno) Timeout() bool {
//line /snap/go/10455/src/net/cgo_unix.go:35
	_go_fuzz_dep_.CoverTab[4599]++
//line /snap/go/10455/src/net/cgo_unix.go:35
	return false
//line /snap/go/10455/src/net/cgo_unix.go:35
	// _ = "end of CoverTab[4599]"
//line /snap/go/10455/src/net/cgo_unix.go:35
}

// isAddrinfoErrno is just for testing purposes.
func (eai addrinfoErrno) isAddrinfoErrno() {
//line /snap/go/10455/src/net/cgo_unix.go:38
	_go_fuzz_dep_.CoverTab[4600]++
//line /snap/go/10455/src/net/cgo_unix.go:38
	// _ = "end of CoverTab[4600]"
//line /snap/go/10455/src/net/cgo_unix.go:38
}

// doBlockingWithCtx executes a blocking function in a separate goroutine when the provided
//line /snap/go/10455/src/net/cgo_unix.go:40
// context is cancellable. It is intended for use with calls that don't support context
//line /snap/go/10455/src/net/cgo_unix.go:40
// cancellation (cgo, syscalls). blocking func may still be running after this function finishes.
//line /snap/go/10455/src/net/cgo_unix.go:43
func doBlockingWithCtx[T any](ctx context.Context, blocking func() (T, error)) (T, error) {
//line /snap/go/10455/src/net/cgo_unix.go:43
	_go_fuzz_dep_.CoverTab[4601]++
						if ctx.Done() == nil {
//line /snap/go/10455/src/net/cgo_unix.go:44
		_go_fuzz_dep_.CoverTab[527447]++
//line /snap/go/10455/src/net/cgo_unix.go:44
		_go_fuzz_dep_.CoverTab[4604]++
							return blocking()
//line /snap/go/10455/src/net/cgo_unix.go:45
		// _ = "end of CoverTab[4604]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:46
		_go_fuzz_dep_.CoverTab[527448]++
//line /snap/go/10455/src/net/cgo_unix.go:46
		_go_fuzz_dep_.CoverTab[4605]++
//line /snap/go/10455/src/net/cgo_unix.go:46
		// _ = "end of CoverTab[4605]"
//line /snap/go/10455/src/net/cgo_unix.go:46
	}
//line /snap/go/10455/src/net/cgo_unix.go:46
	// _ = "end of CoverTab[4601]"
//line /snap/go/10455/src/net/cgo_unix.go:46
	_go_fuzz_dep_.CoverTab[4602]++

						type result struct {
		res	T
		err	error
	}

						res := make(chan result, 1)
//line /snap/go/10455/src/net/cgo_unix.go:53
	_curRoutineNum3_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /snap/go/10455/src/net/cgo_unix.go:53
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum3_)
						go func() {
//line /snap/go/10455/src/net/cgo_unix.go:54
		_go_fuzz_dep_.CoverTab[4606]++
//line /snap/go/10455/src/net/cgo_unix.go:54
		defer func() {
//line /snap/go/10455/src/net/cgo_unix.go:54
			_go_fuzz_dep_.CoverTab[4607]++
//line /snap/go/10455/src/net/cgo_unix.go:54
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum3_)
//line /snap/go/10455/src/net/cgo_unix.go:54
			// _ = "end of CoverTab[4607]"
//line /snap/go/10455/src/net/cgo_unix.go:54
		}()
							var r result
							r.res, r.err = blocking()
							res <- r
//line /snap/go/10455/src/net/cgo_unix.go:57
		// _ = "end of CoverTab[4606]"
	}()
//line /snap/go/10455/src/net/cgo_unix.go:58
	// _ = "end of CoverTab[4602]"
//line /snap/go/10455/src/net/cgo_unix.go:58
	_go_fuzz_dep_.CoverTab[4603]++

						select {
	case r := <-res:
//line /snap/go/10455/src/net/cgo_unix.go:61
		_go_fuzz_dep_.CoverTab[4608]++
							return r.res, r.err
//line /snap/go/10455/src/net/cgo_unix.go:62
		// _ = "end of CoverTab[4608]"
	case <-ctx.Done():
//line /snap/go/10455/src/net/cgo_unix.go:63
		_go_fuzz_dep_.CoverTab[4609]++
							var zero T
							return zero, mapErr(ctx.Err())
//line /snap/go/10455/src/net/cgo_unix.go:65
		// _ = "end of CoverTab[4609]"
	}
//line /snap/go/10455/src/net/cgo_unix.go:66
	// _ = "end of CoverTab[4603]"
}

func cgoLookupHost(ctx context.Context, name string) (hosts []string, err error) {
//line /snap/go/10455/src/net/cgo_unix.go:69
	_go_fuzz_dep_.CoverTab[4610]++
						addrs, err := cgoLookupIP(ctx, "ip", name)
						if err != nil {
//line /snap/go/10455/src/net/cgo_unix.go:71
		_go_fuzz_dep_.CoverTab[527449]++
//line /snap/go/10455/src/net/cgo_unix.go:71
		_go_fuzz_dep_.CoverTab[4613]++
							return nil, err
//line /snap/go/10455/src/net/cgo_unix.go:72
		// _ = "end of CoverTab[4613]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:73
		_go_fuzz_dep_.CoverTab[527450]++
//line /snap/go/10455/src/net/cgo_unix.go:73
		_go_fuzz_dep_.CoverTab[4614]++
//line /snap/go/10455/src/net/cgo_unix.go:73
		// _ = "end of CoverTab[4614]"
//line /snap/go/10455/src/net/cgo_unix.go:73
	}
//line /snap/go/10455/src/net/cgo_unix.go:73
	// _ = "end of CoverTab[4610]"
//line /snap/go/10455/src/net/cgo_unix.go:73
	_go_fuzz_dep_.CoverTab[4611]++
//line /snap/go/10455/src/net/cgo_unix.go:73
	_go_fuzz_dep_.CoverTab[786640] = 0
						for _, addr := range addrs {
//line /snap/go/10455/src/net/cgo_unix.go:74
		if _go_fuzz_dep_.CoverTab[786640] == 0 {
//line /snap/go/10455/src/net/cgo_unix.go:74
			_go_fuzz_dep_.CoverTab[527521]++
//line /snap/go/10455/src/net/cgo_unix.go:74
		} else {
//line /snap/go/10455/src/net/cgo_unix.go:74
			_go_fuzz_dep_.CoverTab[527522]++
//line /snap/go/10455/src/net/cgo_unix.go:74
		}
//line /snap/go/10455/src/net/cgo_unix.go:74
		_go_fuzz_dep_.CoverTab[786640] = 1
//line /snap/go/10455/src/net/cgo_unix.go:74
		_go_fuzz_dep_.CoverTab[4615]++
							hosts = append(hosts, addr.String())
//line /snap/go/10455/src/net/cgo_unix.go:75
		// _ = "end of CoverTab[4615]"
	}
//line /snap/go/10455/src/net/cgo_unix.go:76
	if _go_fuzz_dep_.CoverTab[786640] == 0 {
//line /snap/go/10455/src/net/cgo_unix.go:76
		_go_fuzz_dep_.CoverTab[527523]++
//line /snap/go/10455/src/net/cgo_unix.go:76
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:76
		_go_fuzz_dep_.CoverTab[527524]++
//line /snap/go/10455/src/net/cgo_unix.go:76
	}
//line /snap/go/10455/src/net/cgo_unix.go:76
	// _ = "end of CoverTab[4611]"
//line /snap/go/10455/src/net/cgo_unix.go:76
	_go_fuzz_dep_.CoverTab[4612]++
						return hosts, nil
//line /snap/go/10455/src/net/cgo_unix.go:77
	// _ = "end of CoverTab[4612]"
}

func cgoLookupPort(ctx context.Context, network, service string) (port int, err error) {
//line /snap/go/10455/src/net/cgo_unix.go:80
	_go_fuzz_dep_.CoverTab[4616]++
						var hints _C_struct_addrinfo
						switch network {
	case "":
//line /snap/go/10455/src/net/cgo_unix.go:83
		_go_fuzz_dep_.CoverTab[527451]++
//line /snap/go/10455/src/net/cgo_unix.go:83
		_go_fuzz_dep_.CoverTab[4619]++
//line /snap/go/10455/src/net/cgo_unix.go:83
		// _ = "end of CoverTab[4619]"
	case "tcp", "tcp4", "tcp6":
//line /snap/go/10455/src/net/cgo_unix.go:84
		_go_fuzz_dep_.CoverTab[527452]++
//line /snap/go/10455/src/net/cgo_unix.go:84
		_go_fuzz_dep_.CoverTab[4620]++
							*_C_ai_socktype(&hints) = _C_SOCK_STREAM
							*_C_ai_protocol(&hints) = _C_IPPROTO_TCP
//line /snap/go/10455/src/net/cgo_unix.go:86
		// _ = "end of CoverTab[4620]"
	case "udp", "udp4", "udp6":
//line /snap/go/10455/src/net/cgo_unix.go:87
		_go_fuzz_dep_.CoverTab[527453]++
//line /snap/go/10455/src/net/cgo_unix.go:87
		_go_fuzz_dep_.CoverTab[4621]++
							*_C_ai_socktype(&hints) = _C_SOCK_DGRAM
							*_C_ai_protocol(&hints) = _C_IPPROTO_UDP
//line /snap/go/10455/src/net/cgo_unix.go:89
		// _ = "end of CoverTab[4621]"
	default:
//line /snap/go/10455/src/net/cgo_unix.go:90
		_go_fuzz_dep_.CoverTab[527454]++
//line /snap/go/10455/src/net/cgo_unix.go:90
		_go_fuzz_dep_.CoverTab[4622]++
							return 0, &DNSError{Err: "unknown network", Name: network + "/" + service}
//line /snap/go/10455/src/net/cgo_unix.go:91
		// _ = "end of CoverTab[4622]"
	}
//line /snap/go/10455/src/net/cgo_unix.go:92
	// _ = "end of CoverTab[4616]"
//line /snap/go/10455/src/net/cgo_unix.go:92
	_go_fuzz_dep_.CoverTab[4617]++
						switch ipVersion(network) {
	case '4':
//line /snap/go/10455/src/net/cgo_unix.go:94
		_go_fuzz_dep_.CoverTab[527455]++
//line /snap/go/10455/src/net/cgo_unix.go:94
		_go_fuzz_dep_.CoverTab[4623]++
							*_C_ai_family(&hints) = _C_AF_INET
//line /snap/go/10455/src/net/cgo_unix.go:95
		// _ = "end of CoverTab[4623]"
	case '6':
//line /snap/go/10455/src/net/cgo_unix.go:96
		_go_fuzz_dep_.CoverTab[527456]++
//line /snap/go/10455/src/net/cgo_unix.go:96
		_go_fuzz_dep_.CoverTab[4624]++
							*_C_ai_family(&hints) = _C_AF_INET6
//line /snap/go/10455/src/net/cgo_unix.go:97
		// _ = "end of CoverTab[4624]"
//line /snap/go/10455/src/net/cgo_unix.go:97
	default:
//line /snap/go/10455/src/net/cgo_unix.go:97
		_go_fuzz_dep_.CoverTab[527457]++
//line /snap/go/10455/src/net/cgo_unix.go:97
		_go_fuzz_dep_.CoverTab[4625]++
//line /snap/go/10455/src/net/cgo_unix.go:97
		// _ = "end of CoverTab[4625]"
	}
//line /snap/go/10455/src/net/cgo_unix.go:98
	// _ = "end of CoverTab[4617]"
//line /snap/go/10455/src/net/cgo_unix.go:98
	_go_fuzz_dep_.CoverTab[4618]++

						return doBlockingWithCtx(ctx, func() (int, error) {
//line /snap/go/10455/src/net/cgo_unix.go:100
		_go_fuzz_dep_.CoverTab[4626]++
							return cgoLookupServicePort(&hints, network, service)
//line /snap/go/10455/src/net/cgo_unix.go:101
		// _ = "end of CoverTab[4626]"
	})
//line /snap/go/10455/src/net/cgo_unix.go:102
	// _ = "end of CoverTab[4618]"
}

func cgoLookupServicePort(hints *_C_struct_addrinfo, network, service string) (port int, err error) {
//line /snap/go/10455/src/net/cgo_unix.go:105
	_go_fuzz_dep_.CoverTab[4627]++
						cservice, err := syscall.ByteSliceFromString(service)
						if err != nil {
//line /snap/go/10455/src/net/cgo_unix.go:107
		_go_fuzz_dep_.CoverTab[527458]++
//line /snap/go/10455/src/net/cgo_unix.go:107
		_go_fuzz_dep_.CoverTab[4632]++
							return 0, &DNSError{Err: err.Error(), Name: network + "/" + service}
//line /snap/go/10455/src/net/cgo_unix.go:108
		// _ = "end of CoverTab[4632]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:109
		_go_fuzz_dep_.CoverTab[527459]++
//line /snap/go/10455/src/net/cgo_unix.go:109
		_go_fuzz_dep_.CoverTab[4633]++
//line /snap/go/10455/src/net/cgo_unix.go:109
		// _ = "end of CoverTab[4633]"
//line /snap/go/10455/src/net/cgo_unix.go:109
	}
//line /snap/go/10455/src/net/cgo_unix.go:109
	// _ = "end of CoverTab[4627]"
//line /snap/go/10455/src/net/cgo_unix.go:109
	_go_fuzz_dep_.CoverTab[4628]++
//line /snap/go/10455/src/net/cgo_unix.go:109
	_go_fuzz_dep_.CoverTab[786641] = 0

						for i, b := range cservice[:len(service)] {
//line /snap/go/10455/src/net/cgo_unix.go:111
		if _go_fuzz_dep_.CoverTab[786641] == 0 {
//line /snap/go/10455/src/net/cgo_unix.go:111
			_go_fuzz_dep_.CoverTab[527525]++
//line /snap/go/10455/src/net/cgo_unix.go:111
		} else {
//line /snap/go/10455/src/net/cgo_unix.go:111
			_go_fuzz_dep_.CoverTab[527526]++
//line /snap/go/10455/src/net/cgo_unix.go:111
		}
//line /snap/go/10455/src/net/cgo_unix.go:111
		_go_fuzz_dep_.CoverTab[786641] = 1
//line /snap/go/10455/src/net/cgo_unix.go:111
		_go_fuzz_dep_.CoverTab[4634]++
							cservice[i] = lowerASCII(b)
//line /snap/go/10455/src/net/cgo_unix.go:112
		// _ = "end of CoverTab[4634]"
	}
//line /snap/go/10455/src/net/cgo_unix.go:113
	if _go_fuzz_dep_.CoverTab[786641] == 0 {
//line /snap/go/10455/src/net/cgo_unix.go:113
		_go_fuzz_dep_.CoverTab[527527]++
//line /snap/go/10455/src/net/cgo_unix.go:113
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:113
		_go_fuzz_dep_.CoverTab[527528]++
//line /snap/go/10455/src/net/cgo_unix.go:113
	}
//line /snap/go/10455/src/net/cgo_unix.go:113
	// _ = "end of CoverTab[4628]"
//line /snap/go/10455/src/net/cgo_unix.go:113
	_go_fuzz_dep_.CoverTab[4629]++
						var res *_C_struct_addrinfo
						gerrno, err := _C_getaddrinfo(nil, (*_C_char)(unsafe.Pointer(&cservice[0])), hints, &res)
						if gerrno != 0 {
//line /snap/go/10455/src/net/cgo_unix.go:116
		_go_fuzz_dep_.CoverTab[527460]++
//line /snap/go/10455/src/net/cgo_unix.go:116
		_go_fuzz_dep_.CoverTab[4635]++
							isTemporary := false
							switch gerrno {
		case _C_EAI_SYSTEM:
//line /snap/go/10455/src/net/cgo_unix.go:119
			_go_fuzz_dep_.CoverTab[527462]++
//line /snap/go/10455/src/net/cgo_unix.go:119
			_go_fuzz_dep_.CoverTab[4637]++
								if err == nil {
//line /snap/go/10455/src/net/cgo_unix.go:120
				_go_fuzz_dep_.CoverTab[527464]++
//line /snap/go/10455/src/net/cgo_unix.go:120
				_go_fuzz_dep_.CoverTab[4639]++
									err = syscall.EMFILE
//line /snap/go/10455/src/net/cgo_unix.go:121
				// _ = "end of CoverTab[4639]"
			} else {
//line /snap/go/10455/src/net/cgo_unix.go:122
				_go_fuzz_dep_.CoverTab[527465]++
//line /snap/go/10455/src/net/cgo_unix.go:122
				_go_fuzz_dep_.CoverTab[4640]++
//line /snap/go/10455/src/net/cgo_unix.go:122
				// _ = "end of CoverTab[4640]"
//line /snap/go/10455/src/net/cgo_unix.go:122
			}
//line /snap/go/10455/src/net/cgo_unix.go:122
			// _ = "end of CoverTab[4637]"
		default:
//line /snap/go/10455/src/net/cgo_unix.go:123
			_go_fuzz_dep_.CoverTab[527463]++
//line /snap/go/10455/src/net/cgo_unix.go:123
			_go_fuzz_dep_.CoverTab[4638]++
								err = addrinfoErrno(gerrno)
								isTemporary = addrinfoErrno(gerrno).Temporary()
//line /snap/go/10455/src/net/cgo_unix.go:125
			// _ = "end of CoverTab[4638]"
		}
//line /snap/go/10455/src/net/cgo_unix.go:126
		// _ = "end of CoverTab[4635]"
//line /snap/go/10455/src/net/cgo_unix.go:126
		_go_fuzz_dep_.CoverTab[4636]++
							return 0, &DNSError{Err: err.Error(), Name: network + "/" + service, IsTemporary: isTemporary}
//line /snap/go/10455/src/net/cgo_unix.go:127
		// _ = "end of CoverTab[4636]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:128
		_go_fuzz_dep_.CoverTab[527461]++
//line /snap/go/10455/src/net/cgo_unix.go:128
		_go_fuzz_dep_.CoverTab[4641]++
//line /snap/go/10455/src/net/cgo_unix.go:128
		// _ = "end of CoverTab[4641]"
//line /snap/go/10455/src/net/cgo_unix.go:128
	}
//line /snap/go/10455/src/net/cgo_unix.go:128
	// _ = "end of CoverTab[4629]"
//line /snap/go/10455/src/net/cgo_unix.go:128
	_go_fuzz_dep_.CoverTab[4630]++
						defer _C_freeaddrinfo(res)
//line /snap/go/10455/src/net/cgo_unix.go:129
	_go_fuzz_dep_.CoverTab[786642] = 0

						for r := res; r != nil; r = *_C_ai_next(r) {
//line /snap/go/10455/src/net/cgo_unix.go:131
		if _go_fuzz_dep_.CoverTab[786642] == 0 {
//line /snap/go/10455/src/net/cgo_unix.go:131
			_go_fuzz_dep_.CoverTab[527529]++
//line /snap/go/10455/src/net/cgo_unix.go:131
		} else {
//line /snap/go/10455/src/net/cgo_unix.go:131
			_go_fuzz_dep_.CoverTab[527530]++
//line /snap/go/10455/src/net/cgo_unix.go:131
		}
//line /snap/go/10455/src/net/cgo_unix.go:131
		_go_fuzz_dep_.CoverTab[786642] = 1
//line /snap/go/10455/src/net/cgo_unix.go:131
		_go_fuzz_dep_.CoverTab[4642]++
							switch *_C_ai_family(r) {
		case _C_AF_INET:
//line /snap/go/10455/src/net/cgo_unix.go:133
			_go_fuzz_dep_.CoverTab[527466]++
//line /snap/go/10455/src/net/cgo_unix.go:133
			_go_fuzz_dep_.CoverTab[4643]++
								sa := (*syscall.RawSockaddrInet4)(unsafe.Pointer(*_C_ai_addr(r)))
								p := (*[2]byte)(unsafe.Pointer(&sa.Port))
								return int(p[0])<<8 | int(p[1]), nil
//line /snap/go/10455/src/net/cgo_unix.go:136
			// _ = "end of CoverTab[4643]"
		case _C_AF_INET6:
//line /snap/go/10455/src/net/cgo_unix.go:137
			_go_fuzz_dep_.CoverTab[527467]++
//line /snap/go/10455/src/net/cgo_unix.go:137
			_go_fuzz_dep_.CoverTab[4644]++
								sa := (*syscall.RawSockaddrInet6)(unsafe.Pointer(*_C_ai_addr(r)))
								p := (*[2]byte)(unsafe.Pointer(&sa.Port))
								return int(p[0])<<8 | int(p[1]), nil
//line /snap/go/10455/src/net/cgo_unix.go:140
			// _ = "end of CoverTab[4644]"
//line /snap/go/10455/src/net/cgo_unix.go:140
		default:
//line /snap/go/10455/src/net/cgo_unix.go:140
			_go_fuzz_dep_.CoverTab[527468]++
//line /snap/go/10455/src/net/cgo_unix.go:140
			_go_fuzz_dep_.CoverTab[4645]++
//line /snap/go/10455/src/net/cgo_unix.go:140
			// _ = "end of CoverTab[4645]"
		}
//line /snap/go/10455/src/net/cgo_unix.go:141
		// _ = "end of CoverTab[4642]"
	}
//line /snap/go/10455/src/net/cgo_unix.go:142
	if _go_fuzz_dep_.CoverTab[786642] == 0 {
//line /snap/go/10455/src/net/cgo_unix.go:142
		_go_fuzz_dep_.CoverTab[527531]++
//line /snap/go/10455/src/net/cgo_unix.go:142
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:142
		_go_fuzz_dep_.CoverTab[527532]++
//line /snap/go/10455/src/net/cgo_unix.go:142
	}
//line /snap/go/10455/src/net/cgo_unix.go:142
	// _ = "end of CoverTab[4630]"
//line /snap/go/10455/src/net/cgo_unix.go:142
	_go_fuzz_dep_.CoverTab[4631]++
						return 0, &DNSError{Err: "unknown port", Name: network + "/" + service}
//line /snap/go/10455/src/net/cgo_unix.go:143
	// _ = "end of CoverTab[4631]"
}

func cgoLookupHostIP(network, name string) (addrs []IPAddr, err error) {
//line /snap/go/10455/src/net/cgo_unix.go:146
	_go_fuzz_dep_.CoverTab[4646]++
						acquireThread()
						defer releaseThread()

						var hints _C_struct_addrinfo
						*_C_ai_flags(&hints) = cgoAddrInfoFlags
						*_C_ai_socktype(&hints) = _C_SOCK_STREAM
						*_C_ai_family(&hints) = _C_AF_UNSPEC
						switch ipVersion(network) {
	case '4':
//line /snap/go/10455/src/net/cgo_unix.go:155
		_go_fuzz_dep_.CoverTab[527469]++
//line /snap/go/10455/src/net/cgo_unix.go:155
		_go_fuzz_dep_.CoverTab[4651]++
							*_C_ai_family(&hints) = _C_AF_INET
//line /snap/go/10455/src/net/cgo_unix.go:156
		// _ = "end of CoverTab[4651]"
	case '6':
//line /snap/go/10455/src/net/cgo_unix.go:157
		_go_fuzz_dep_.CoverTab[527470]++
//line /snap/go/10455/src/net/cgo_unix.go:157
		_go_fuzz_dep_.CoverTab[4652]++
							*_C_ai_family(&hints) = _C_AF_INET6
//line /snap/go/10455/src/net/cgo_unix.go:158
		// _ = "end of CoverTab[4652]"
//line /snap/go/10455/src/net/cgo_unix.go:158
	default:
//line /snap/go/10455/src/net/cgo_unix.go:158
		_go_fuzz_dep_.CoverTab[527471]++
//line /snap/go/10455/src/net/cgo_unix.go:158
		_go_fuzz_dep_.CoverTab[4653]++
//line /snap/go/10455/src/net/cgo_unix.go:158
		// _ = "end of CoverTab[4653]"
	}
//line /snap/go/10455/src/net/cgo_unix.go:159
	// _ = "end of CoverTab[4646]"
//line /snap/go/10455/src/net/cgo_unix.go:159
	_go_fuzz_dep_.CoverTab[4647]++

						h, err := syscall.BytePtrFromString(name)
						if err != nil {
//line /snap/go/10455/src/net/cgo_unix.go:162
		_go_fuzz_dep_.CoverTab[527472]++
//line /snap/go/10455/src/net/cgo_unix.go:162
		_go_fuzz_dep_.CoverTab[4654]++
							return nil, &DNSError{Err: err.Error(), Name: name}
//line /snap/go/10455/src/net/cgo_unix.go:163
		// _ = "end of CoverTab[4654]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:164
		_go_fuzz_dep_.CoverTab[527473]++
//line /snap/go/10455/src/net/cgo_unix.go:164
		_go_fuzz_dep_.CoverTab[4655]++
//line /snap/go/10455/src/net/cgo_unix.go:164
		// _ = "end of CoverTab[4655]"
//line /snap/go/10455/src/net/cgo_unix.go:164
	}
//line /snap/go/10455/src/net/cgo_unix.go:164
	// _ = "end of CoverTab[4647]"
//line /snap/go/10455/src/net/cgo_unix.go:164
	_go_fuzz_dep_.CoverTab[4648]++
						var res *_C_struct_addrinfo
						gerrno, err := _C_getaddrinfo((*_C_char)(unsafe.Pointer(h)), nil, &hints, &res)
						if gerrno != 0 {
//line /snap/go/10455/src/net/cgo_unix.go:167
		_go_fuzz_dep_.CoverTab[527474]++
//line /snap/go/10455/src/net/cgo_unix.go:167
		_go_fuzz_dep_.CoverTab[4656]++
							isErrorNoSuchHost := false
							isTemporary := false
							switch gerrno {
		case _C_EAI_SYSTEM:
//line /snap/go/10455/src/net/cgo_unix.go:171
			_go_fuzz_dep_.CoverTab[527476]++
//line /snap/go/10455/src/net/cgo_unix.go:171
			_go_fuzz_dep_.CoverTab[4658]++
								if err == nil {
//line /snap/go/10455/src/net/cgo_unix.go:172
				_go_fuzz_dep_.CoverTab[527479]++
//line /snap/go/10455/src/net/cgo_unix.go:172
				_go_fuzz_dep_.CoverTab[4661]++

//line /snap/go/10455/src/net/cgo_unix.go:180
				err = syscall.EMFILE
//line /snap/go/10455/src/net/cgo_unix.go:180
				// _ = "end of CoverTab[4661]"
			} else {
//line /snap/go/10455/src/net/cgo_unix.go:181
				_go_fuzz_dep_.CoverTab[527480]++
//line /snap/go/10455/src/net/cgo_unix.go:181
				_go_fuzz_dep_.CoverTab[4662]++
//line /snap/go/10455/src/net/cgo_unix.go:181
				// _ = "end of CoverTab[4662]"
//line /snap/go/10455/src/net/cgo_unix.go:181
			}
//line /snap/go/10455/src/net/cgo_unix.go:181
			// _ = "end of CoverTab[4658]"
		case _C_EAI_NONAME, _C_EAI_NODATA:
//line /snap/go/10455/src/net/cgo_unix.go:182
			_go_fuzz_dep_.CoverTab[527477]++
//line /snap/go/10455/src/net/cgo_unix.go:182
			_go_fuzz_dep_.CoverTab[4659]++
								err = errNoSuchHost
								isErrorNoSuchHost = true
//line /snap/go/10455/src/net/cgo_unix.go:184
			// _ = "end of CoverTab[4659]"
		default:
//line /snap/go/10455/src/net/cgo_unix.go:185
			_go_fuzz_dep_.CoverTab[527478]++
//line /snap/go/10455/src/net/cgo_unix.go:185
			_go_fuzz_dep_.CoverTab[4660]++
								err = addrinfoErrno(gerrno)
								isTemporary = addrinfoErrno(gerrno).Temporary()
//line /snap/go/10455/src/net/cgo_unix.go:187
			// _ = "end of CoverTab[4660]"
		}
//line /snap/go/10455/src/net/cgo_unix.go:188
		// _ = "end of CoverTab[4656]"
//line /snap/go/10455/src/net/cgo_unix.go:188
		_go_fuzz_dep_.CoverTab[4657]++

							return nil, &DNSError{Err: err.Error(), Name: name, IsNotFound: isErrorNoSuchHost, IsTemporary: isTemporary}
//line /snap/go/10455/src/net/cgo_unix.go:190
		// _ = "end of CoverTab[4657]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:191
		_go_fuzz_dep_.CoverTab[527475]++
//line /snap/go/10455/src/net/cgo_unix.go:191
		_go_fuzz_dep_.CoverTab[4663]++
//line /snap/go/10455/src/net/cgo_unix.go:191
		// _ = "end of CoverTab[4663]"
//line /snap/go/10455/src/net/cgo_unix.go:191
	}
//line /snap/go/10455/src/net/cgo_unix.go:191
	// _ = "end of CoverTab[4648]"
//line /snap/go/10455/src/net/cgo_unix.go:191
	_go_fuzz_dep_.CoverTab[4649]++
						defer _C_freeaddrinfo(res)
//line /snap/go/10455/src/net/cgo_unix.go:192
	_go_fuzz_dep_.CoverTab[786643] = 0

						for r := res; r != nil; r = *_C_ai_next(r) {
//line /snap/go/10455/src/net/cgo_unix.go:194
		if _go_fuzz_dep_.CoverTab[786643] == 0 {
//line /snap/go/10455/src/net/cgo_unix.go:194
			_go_fuzz_dep_.CoverTab[527533]++
//line /snap/go/10455/src/net/cgo_unix.go:194
		} else {
//line /snap/go/10455/src/net/cgo_unix.go:194
			_go_fuzz_dep_.CoverTab[527534]++
//line /snap/go/10455/src/net/cgo_unix.go:194
		}
//line /snap/go/10455/src/net/cgo_unix.go:194
		_go_fuzz_dep_.CoverTab[786643] = 1
//line /snap/go/10455/src/net/cgo_unix.go:194
		_go_fuzz_dep_.CoverTab[4664]++

							if *_C_ai_socktype(r) != _C_SOCK_STREAM {
//line /snap/go/10455/src/net/cgo_unix.go:196
			_go_fuzz_dep_.CoverTab[527481]++
//line /snap/go/10455/src/net/cgo_unix.go:196
			_go_fuzz_dep_.CoverTab[4666]++
								continue
//line /snap/go/10455/src/net/cgo_unix.go:197
			// _ = "end of CoverTab[4666]"
		} else {
//line /snap/go/10455/src/net/cgo_unix.go:198
			_go_fuzz_dep_.CoverTab[527482]++
//line /snap/go/10455/src/net/cgo_unix.go:198
			_go_fuzz_dep_.CoverTab[4667]++
//line /snap/go/10455/src/net/cgo_unix.go:198
			// _ = "end of CoverTab[4667]"
//line /snap/go/10455/src/net/cgo_unix.go:198
		}
//line /snap/go/10455/src/net/cgo_unix.go:198
		// _ = "end of CoverTab[4664]"
//line /snap/go/10455/src/net/cgo_unix.go:198
		_go_fuzz_dep_.CoverTab[4665]++
							switch *_C_ai_family(r) {
		case _C_AF_INET:
//line /snap/go/10455/src/net/cgo_unix.go:200
			_go_fuzz_dep_.CoverTab[527483]++
//line /snap/go/10455/src/net/cgo_unix.go:200
			_go_fuzz_dep_.CoverTab[4668]++
								sa := (*syscall.RawSockaddrInet4)(unsafe.Pointer(*_C_ai_addr(r)))
								addr := IPAddr{IP: copyIP(sa.Addr[:])}
								addrs = append(addrs, addr)
//line /snap/go/10455/src/net/cgo_unix.go:203
			// _ = "end of CoverTab[4668]"
		case _C_AF_INET6:
//line /snap/go/10455/src/net/cgo_unix.go:204
			_go_fuzz_dep_.CoverTab[527484]++
//line /snap/go/10455/src/net/cgo_unix.go:204
			_go_fuzz_dep_.CoverTab[4669]++
								sa := (*syscall.RawSockaddrInet6)(unsafe.Pointer(*_C_ai_addr(r)))
								addr := IPAddr{IP: copyIP(sa.Addr[:]), Zone: zoneCache.name(int(sa.Scope_id))}
								addrs = append(addrs, addr)
//line /snap/go/10455/src/net/cgo_unix.go:207
			// _ = "end of CoverTab[4669]"
//line /snap/go/10455/src/net/cgo_unix.go:207
		default:
//line /snap/go/10455/src/net/cgo_unix.go:207
			_go_fuzz_dep_.CoverTab[527485]++
//line /snap/go/10455/src/net/cgo_unix.go:207
			_go_fuzz_dep_.CoverTab[4670]++
//line /snap/go/10455/src/net/cgo_unix.go:207
			// _ = "end of CoverTab[4670]"
		}
//line /snap/go/10455/src/net/cgo_unix.go:208
		// _ = "end of CoverTab[4665]"
	}
//line /snap/go/10455/src/net/cgo_unix.go:209
	if _go_fuzz_dep_.CoverTab[786643] == 0 {
//line /snap/go/10455/src/net/cgo_unix.go:209
		_go_fuzz_dep_.CoverTab[527535]++
//line /snap/go/10455/src/net/cgo_unix.go:209
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:209
		_go_fuzz_dep_.CoverTab[527536]++
//line /snap/go/10455/src/net/cgo_unix.go:209
	}
//line /snap/go/10455/src/net/cgo_unix.go:209
	// _ = "end of CoverTab[4649]"
//line /snap/go/10455/src/net/cgo_unix.go:209
	_go_fuzz_dep_.CoverTab[4650]++
						return addrs, nil
//line /snap/go/10455/src/net/cgo_unix.go:210
	// _ = "end of CoverTab[4650]"
}

func cgoLookupIP(ctx context.Context, network, name string) (addrs []IPAddr, err error) {
//line /snap/go/10455/src/net/cgo_unix.go:213
	_go_fuzz_dep_.CoverTab[4671]++
						return doBlockingWithCtx(ctx, func() ([]IPAddr, error) {
//line /snap/go/10455/src/net/cgo_unix.go:214
		_go_fuzz_dep_.CoverTab[4672]++
							return cgoLookupHostIP(network, name)
//line /snap/go/10455/src/net/cgo_unix.go:215
		// _ = "end of CoverTab[4672]"
	})
//line /snap/go/10455/src/net/cgo_unix.go:216
	// _ = "end of CoverTab[4671]"
}

// These are roughly enough for the following:
//line /snap/go/10455/src/net/cgo_unix.go:219
//
//line /snap/go/10455/src/net/cgo_unix.go:219
//	 Source		Encoding			Maximum length of single name entry
//line /snap/go/10455/src/net/cgo_unix.go:219
//	 Unicast DNS		ASCII or			<=253 + a NUL terminator
//line /snap/go/10455/src/net/cgo_unix.go:219
//				Unicode in RFC 5892		252 * total number of labels + delimiters + a NUL terminator
//line /snap/go/10455/src/net/cgo_unix.go:219
//	 Multicast DNS	UTF-8 in RFC 5198 or		<=253 + a NUL terminator
//line /snap/go/10455/src/net/cgo_unix.go:219
//				the same as unicast DNS ASCII	<=253 + a NUL terminator
//line /snap/go/10455/src/net/cgo_unix.go:219
//	 Local database	various				depends on implementation
//line /snap/go/10455/src/net/cgo_unix.go:227
const (
	nameinfoLen	= 64
	maxNameinfoLen	= 4096
)

func cgoLookupPTR(ctx context.Context, addr string) (names []string, err error) {
//line /snap/go/10455/src/net/cgo_unix.go:232
	_go_fuzz_dep_.CoverTab[4673]++
						ip, err := netip.ParseAddr(addr)
						if err != nil {
//line /snap/go/10455/src/net/cgo_unix.go:234
		_go_fuzz_dep_.CoverTab[527486]++
//line /snap/go/10455/src/net/cgo_unix.go:234
		_go_fuzz_dep_.CoverTab[4676]++
							return nil, &DNSError{Err: "invalid address", Name: addr}
//line /snap/go/10455/src/net/cgo_unix.go:235
		// _ = "end of CoverTab[4676]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:236
		_go_fuzz_dep_.CoverTab[527487]++
//line /snap/go/10455/src/net/cgo_unix.go:236
		_go_fuzz_dep_.CoverTab[4677]++
//line /snap/go/10455/src/net/cgo_unix.go:236
		// _ = "end of CoverTab[4677]"
//line /snap/go/10455/src/net/cgo_unix.go:236
	}
//line /snap/go/10455/src/net/cgo_unix.go:236
	// _ = "end of CoverTab[4673]"
//line /snap/go/10455/src/net/cgo_unix.go:236
	_go_fuzz_dep_.CoverTab[4674]++
						sa, salen := cgoSockaddr(IP(ip.AsSlice()), ip.Zone())
						if sa == nil {
//line /snap/go/10455/src/net/cgo_unix.go:238
		_go_fuzz_dep_.CoverTab[527488]++
//line /snap/go/10455/src/net/cgo_unix.go:238
		_go_fuzz_dep_.CoverTab[4678]++
							return nil, &DNSError{Err: "invalid address " + ip.String(), Name: addr}
//line /snap/go/10455/src/net/cgo_unix.go:239
		// _ = "end of CoverTab[4678]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:240
		_go_fuzz_dep_.CoverTab[527489]++
//line /snap/go/10455/src/net/cgo_unix.go:240
		_go_fuzz_dep_.CoverTab[4679]++
//line /snap/go/10455/src/net/cgo_unix.go:240
		// _ = "end of CoverTab[4679]"
//line /snap/go/10455/src/net/cgo_unix.go:240
	}
//line /snap/go/10455/src/net/cgo_unix.go:240
	// _ = "end of CoverTab[4674]"
//line /snap/go/10455/src/net/cgo_unix.go:240
	_go_fuzz_dep_.CoverTab[4675]++

						return doBlockingWithCtx(ctx, func() ([]string, error) {
//line /snap/go/10455/src/net/cgo_unix.go:242
		_go_fuzz_dep_.CoverTab[4680]++
							return cgoLookupAddrPTR(addr, sa, salen)
//line /snap/go/10455/src/net/cgo_unix.go:243
		// _ = "end of CoverTab[4680]"
	})
//line /snap/go/10455/src/net/cgo_unix.go:244
	// _ = "end of CoverTab[4675]"
}

func cgoLookupAddrPTR(addr string, sa *_C_struct_sockaddr, salen _C_socklen_t) (names []string, err error) {
//line /snap/go/10455/src/net/cgo_unix.go:247
	_go_fuzz_dep_.CoverTab[4681]++
						acquireThread()
						defer releaseThread()

						var gerrno int
						var b []byte
//line /snap/go/10455/src/net/cgo_unix.go:252
	_go_fuzz_dep_.CoverTab[786644] = 0
						for l := nameinfoLen; l <= maxNameinfoLen; l *= 2 {
//line /snap/go/10455/src/net/cgo_unix.go:253
		if _go_fuzz_dep_.CoverTab[786644] == 0 {
//line /snap/go/10455/src/net/cgo_unix.go:253
			_go_fuzz_dep_.CoverTab[527537]++
//line /snap/go/10455/src/net/cgo_unix.go:253
		} else {
//line /snap/go/10455/src/net/cgo_unix.go:253
			_go_fuzz_dep_.CoverTab[527538]++
//line /snap/go/10455/src/net/cgo_unix.go:253
		}
//line /snap/go/10455/src/net/cgo_unix.go:253
		_go_fuzz_dep_.CoverTab[786644] = 1
//line /snap/go/10455/src/net/cgo_unix.go:253
		_go_fuzz_dep_.CoverTab[4685]++
							b = make([]byte, l)
							gerrno, err = cgoNameinfoPTR(b, sa, salen)
							if gerrno == 0 || func() bool {
//line /snap/go/10455/src/net/cgo_unix.go:256
			_go_fuzz_dep_.CoverTab[4686]++
//line /snap/go/10455/src/net/cgo_unix.go:256
			return gerrno != _C_EAI_OVERFLOW
//line /snap/go/10455/src/net/cgo_unix.go:256
			// _ = "end of CoverTab[4686]"
//line /snap/go/10455/src/net/cgo_unix.go:256
		}() {
//line /snap/go/10455/src/net/cgo_unix.go:256
			_go_fuzz_dep_.CoverTab[527490]++
//line /snap/go/10455/src/net/cgo_unix.go:256
			_go_fuzz_dep_.CoverTab[4687]++
								break
//line /snap/go/10455/src/net/cgo_unix.go:257
			// _ = "end of CoverTab[4687]"
		} else {
//line /snap/go/10455/src/net/cgo_unix.go:258
			_go_fuzz_dep_.CoverTab[527491]++
//line /snap/go/10455/src/net/cgo_unix.go:258
			_go_fuzz_dep_.CoverTab[4688]++
//line /snap/go/10455/src/net/cgo_unix.go:258
			// _ = "end of CoverTab[4688]"
//line /snap/go/10455/src/net/cgo_unix.go:258
		}
//line /snap/go/10455/src/net/cgo_unix.go:258
		// _ = "end of CoverTab[4685]"
	}
//line /snap/go/10455/src/net/cgo_unix.go:259
	if _go_fuzz_dep_.CoverTab[786644] == 0 {
//line /snap/go/10455/src/net/cgo_unix.go:259
		_go_fuzz_dep_.CoverTab[527539]++
//line /snap/go/10455/src/net/cgo_unix.go:259
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:259
		_go_fuzz_dep_.CoverTab[527540]++
//line /snap/go/10455/src/net/cgo_unix.go:259
	}
//line /snap/go/10455/src/net/cgo_unix.go:259
	// _ = "end of CoverTab[4681]"
//line /snap/go/10455/src/net/cgo_unix.go:259
	_go_fuzz_dep_.CoverTab[4682]++
						if gerrno != 0 {
//line /snap/go/10455/src/net/cgo_unix.go:260
		_go_fuzz_dep_.CoverTab[527492]++
//line /snap/go/10455/src/net/cgo_unix.go:260
		_go_fuzz_dep_.CoverTab[4689]++
							isErrorNoSuchHost := false
							isTemporary := false
							switch gerrno {
		case _C_EAI_SYSTEM:
//line /snap/go/10455/src/net/cgo_unix.go:264
			_go_fuzz_dep_.CoverTab[527494]++
//line /snap/go/10455/src/net/cgo_unix.go:264
			_go_fuzz_dep_.CoverTab[4691]++
								if err == nil {
//line /snap/go/10455/src/net/cgo_unix.go:265
				_go_fuzz_dep_.CoverTab[527497]++
//line /snap/go/10455/src/net/cgo_unix.go:265
				_go_fuzz_dep_.CoverTab[4694]++
									err = syscall.EMFILE
//line /snap/go/10455/src/net/cgo_unix.go:266
				// _ = "end of CoverTab[4694]"
			} else {
//line /snap/go/10455/src/net/cgo_unix.go:267
				_go_fuzz_dep_.CoverTab[527498]++
//line /snap/go/10455/src/net/cgo_unix.go:267
				_go_fuzz_dep_.CoverTab[4695]++
//line /snap/go/10455/src/net/cgo_unix.go:267
				// _ = "end of CoverTab[4695]"
//line /snap/go/10455/src/net/cgo_unix.go:267
			}
//line /snap/go/10455/src/net/cgo_unix.go:267
			// _ = "end of CoverTab[4691]"
		case _C_EAI_NONAME:
//line /snap/go/10455/src/net/cgo_unix.go:268
			_go_fuzz_dep_.CoverTab[527495]++
//line /snap/go/10455/src/net/cgo_unix.go:268
			_go_fuzz_dep_.CoverTab[4692]++
								err = errNoSuchHost
								isErrorNoSuchHost = true
//line /snap/go/10455/src/net/cgo_unix.go:270
			// _ = "end of CoverTab[4692]"
		default:
//line /snap/go/10455/src/net/cgo_unix.go:271
			_go_fuzz_dep_.CoverTab[527496]++
//line /snap/go/10455/src/net/cgo_unix.go:271
			_go_fuzz_dep_.CoverTab[4693]++
								err = addrinfoErrno(gerrno)
								isTemporary = addrinfoErrno(gerrno).Temporary()
//line /snap/go/10455/src/net/cgo_unix.go:273
			// _ = "end of CoverTab[4693]"
		}
//line /snap/go/10455/src/net/cgo_unix.go:274
		// _ = "end of CoverTab[4689]"
//line /snap/go/10455/src/net/cgo_unix.go:274
		_go_fuzz_dep_.CoverTab[4690]++
							return nil, &DNSError{Err: err.Error(), Name: addr, IsTemporary: isTemporary, IsNotFound: isErrorNoSuchHost}
//line /snap/go/10455/src/net/cgo_unix.go:275
		// _ = "end of CoverTab[4690]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:276
		_go_fuzz_dep_.CoverTab[527493]++
//line /snap/go/10455/src/net/cgo_unix.go:276
		_go_fuzz_dep_.CoverTab[4696]++
//line /snap/go/10455/src/net/cgo_unix.go:276
		// _ = "end of CoverTab[4696]"
//line /snap/go/10455/src/net/cgo_unix.go:276
	}
//line /snap/go/10455/src/net/cgo_unix.go:276
	// _ = "end of CoverTab[4682]"
//line /snap/go/10455/src/net/cgo_unix.go:276
	_go_fuzz_dep_.CoverTab[4683]++
//line /snap/go/10455/src/net/cgo_unix.go:276
	_go_fuzz_dep_.CoverTab[786645] = 0
						for i := 0; i < len(b); i++ {
//line /snap/go/10455/src/net/cgo_unix.go:277
		if _go_fuzz_dep_.CoverTab[786645] == 0 {
//line /snap/go/10455/src/net/cgo_unix.go:277
			_go_fuzz_dep_.CoverTab[527541]++
//line /snap/go/10455/src/net/cgo_unix.go:277
		} else {
//line /snap/go/10455/src/net/cgo_unix.go:277
			_go_fuzz_dep_.CoverTab[527542]++
//line /snap/go/10455/src/net/cgo_unix.go:277
		}
//line /snap/go/10455/src/net/cgo_unix.go:277
		_go_fuzz_dep_.CoverTab[786645] = 1
//line /snap/go/10455/src/net/cgo_unix.go:277
		_go_fuzz_dep_.CoverTab[4697]++
							if b[i] == 0 {
//line /snap/go/10455/src/net/cgo_unix.go:278
			_go_fuzz_dep_.CoverTab[527499]++
//line /snap/go/10455/src/net/cgo_unix.go:278
			_go_fuzz_dep_.CoverTab[4698]++
								b = b[:i]
								break
//line /snap/go/10455/src/net/cgo_unix.go:280
			// _ = "end of CoverTab[4698]"
		} else {
//line /snap/go/10455/src/net/cgo_unix.go:281
			_go_fuzz_dep_.CoverTab[527500]++
//line /snap/go/10455/src/net/cgo_unix.go:281
			_go_fuzz_dep_.CoverTab[4699]++
//line /snap/go/10455/src/net/cgo_unix.go:281
			// _ = "end of CoverTab[4699]"
//line /snap/go/10455/src/net/cgo_unix.go:281
		}
//line /snap/go/10455/src/net/cgo_unix.go:281
		// _ = "end of CoverTab[4697]"
	}
//line /snap/go/10455/src/net/cgo_unix.go:282
	if _go_fuzz_dep_.CoverTab[786645] == 0 {
//line /snap/go/10455/src/net/cgo_unix.go:282
		_go_fuzz_dep_.CoverTab[527543]++
//line /snap/go/10455/src/net/cgo_unix.go:282
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:282
		_go_fuzz_dep_.CoverTab[527544]++
//line /snap/go/10455/src/net/cgo_unix.go:282
	}
//line /snap/go/10455/src/net/cgo_unix.go:282
	// _ = "end of CoverTab[4683]"
//line /snap/go/10455/src/net/cgo_unix.go:282
	_go_fuzz_dep_.CoverTab[4684]++
						return []string{absDomainName(string(b))}, nil
//line /snap/go/10455/src/net/cgo_unix.go:283
	// _ = "end of CoverTab[4684]"
}

func cgoSockaddr(ip IP, zone string) (*_C_struct_sockaddr, _C_socklen_t) {
//line /snap/go/10455/src/net/cgo_unix.go:286
	_go_fuzz_dep_.CoverTab[4700]++
						if ip4 := ip.To4(); ip4 != nil {
//line /snap/go/10455/src/net/cgo_unix.go:287
		_go_fuzz_dep_.CoverTab[527501]++
//line /snap/go/10455/src/net/cgo_unix.go:287
		_go_fuzz_dep_.CoverTab[4703]++
							return cgoSockaddrInet4(ip4), _C_socklen_t(syscall.SizeofSockaddrInet4)
//line /snap/go/10455/src/net/cgo_unix.go:288
		// _ = "end of CoverTab[4703]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:289
		_go_fuzz_dep_.CoverTab[527502]++
//line /snap/go/10455/src/net/cgo_unix.go:289
		_go_fuzz_dep_.CoverTab[4704]++
//line /snap/go/10455/src/net/cgo_unix.go:289
		// _ = "end of CoverTab[4704]"
//line /snap/go/10455/src/net/cgo_unix.go:289
	}
//line /snap/go/10455/src/net/cgo_unix.go:289
	// _ = "end of CoverTab[4700]"
//line /snap/go/10455/src/net/cgo_unix.go:289
	_go_fuzz_dep_.CoverTab[4701]++
						if ip6 := ip.To16(); ip6 != nil {
//line /snap/go/10455/src/net/cgo_unix.go:290
		_go_fuzz_dep_.CoverTab[527503]++
//line /snap/go/10455/src/net/cgo_unix.go:290
		_go_fuzz_dep_.CoverTab[4705]++
							return cgoSockaddrInet6(ip6, zoneCache.index(zone)), _C_socklen_t(syscall.SizeofSockaddrInet6)
//line /snap/go/10455/src/net/cgo_unix.go:291
		// _ = "end of CoverTab[4705]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:292
		_go_fuzz_dep_.CoverTab[527504]++
//line /snap/go/10455/src/net/cgo_unix.go:292
		_go_fuzz_dep_.CoverTab[4706]++
//line /snap/go/10455/src/net/cgo_unix.go:292
		// _ = "end of CoverTab[4706]"
//line /snap/go/10455/src/net/cgo_unix.go:292
	}
//line /snap/go/10455/src/net/cgo_unix.go:292
	// _ = "end of CoverTab[4701]"
//line /snap/go/10455/src/net/cgo_unix.go:292
	_go_fuzz_dep_.CoverTab[4702]++
						return nil, 0
//line /snap/go/10455/src/net/cgo_unix.go:293
	// _ = "end of CoverTab[4702]"
}

func cgoLookupCNAME(ctx context.Context, name string) (cname string, err error, completed bool) {
//line /snap/go/10455/src/net/cgo_unix.go:296
	_go_fuzz_dep_.CoverTab[4707]++
						resources, err := resSearch(ctx, name, int(dnsmessage.TypeCNAME), int(dnsmessage.ClassINET))
						if err != nil {
//line /snap/go/10455/src/net/cgo_unix.go:298
		_go_fuzz_dep_.CoverTab[527505]++
//line /snap/go/10455/src/net/cgo_unix.go:298
		_go_fuzz_dep_.CoverTab[4710]++
							return
//line /snap/go/10455/src/net/cgo_unix.go:299
		// _ = "end of CoverTab[4710]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:300
		_go_fuzz_dep_.CoverTab[527506]++
//line /snap/go/10455/src/net/cgo_unix.go:300
		_go_fuzz_dep_.CoverTab[4711]++
//line /snap/go/10455/src/net/cgo_unix.go:300
		// _ = "end of CoverTab[4711]"
//line /snap/go/10455/src/net/cgo_unix.go:300
	}
//line /snap/go/10455/src/net/cgo_unix.go:300
	// _ = "end of CoverTab[4707]"
//line /snap/go/10455/src/net/cgo_unix.go:300
	_go_fuzz_dep_.CoverTab[4708]++
						cname, err = parseCNAMEFromResources(resources)
						if err != nil {
//line /snap/go/10455/src/net/cgo_unix.go:302
		_go_fuzz_dep_.CoverTab[527507]++
//line /snap/go/10455/src/net/cgo_unix.go:302
		_go_fuzz_dep_.CoverTab[4712]++
							return "", err, false
//line /snap/go/10455/src/net/cgo_unix.go:303
		// _ = "end of CoverTab[4712]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:304
		_go_fuzz_dep_.CoverTab[527508]++
//line /snap/go/10455/src/net/cgo_unix.go:304
		_go_fuzz_dep_.CoverTab[4713]++
//line /snap/go/10455/src/net/cgo_unix.go:304
		// _ = "end of CoverTab[4713]"
//line /snap/go/10455/src/net/cgo_unix.go:304
	}
//line /snap/go/10455/src/net/cgo_unix.go:304
	// _ = "end of CoverTab[4708]"
//line /snap/go/10455/src/net/cgo_unix.go:304
	_go_fuzz_dep_.CoverTab[4709]++
						return cname, nil, true
//line /snap/go/10455/src/net/cgo_unix.go:305
	// _ = "end of CoverTab[4709]"
}

// resSearch will make a call to the 'res_nsearch' routine in the C library
//line /snap/go/10455/src/net/cgo_unix.go:308
// and parse the output as a slice of DNS resources.
//line /snap/go/10455/src/net/cgo_unix.go:310
func resSearch(ctx context.Context, hostname string, rtype, class int) ([]dnsmessage.Resource, error) {
//line /snap/go/10455/src/net/cgo_unix.go:310
	_go_fuzz_dep_.CoverTab[4714]++
						return doBlockingWithCtx(ctx, func() ([]dnsmessage.Resource, error) {
//line /snap/go/10455/src/net/cgo_unix.go:311
		_go_fuzz_dep_.CoverTab[4715]++
							return cgoResSearch(hostname, rtype, class)
//line /snap/go/10455/src/net/cgo_unix.go:312
		// _ = "end of CoverTab[4715]"
	})
//line /snap/go/10455/src/net/cgo_unix.go:313
	// _ = "end of CoverTab[4714]"
}

func cgoResSearch(hostname string, rtype, class int) ([]dnsmessage.Resource, error) {
//line /snap/go/10455/src/net/cgo_unix.go:316
	_go_fuzz_dep_.CoverTab[4716]++
						acquireThread()
						defer releaseThread()

						state := (*_C_struct___res_state)(_C_malloc(unsafe.Sizeof(_C_struct___res_state{})))
						defer _C_free(unsafe.Pointer(state))
						if err := _C_res_ninit(state); err != nil {
//line /snap/go/10455/src/net/cgo_unix.go:322
		_go_fuzz_dep_.CoverTab[527509]++
//line /snap/go/10455/src/net/cgo_unix.go:322
		_go_fuzz_dep_.CoverTab[4722]++
							return nil, errors.New("res_ninit failure: " + err.Error())
//line /snap/go/10455/src/net/cgo_unix.go:323
		// _ = "end of CoverTab[4722]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:324
		_go_fuzz_dep_.CoverTab[527510]++
//line /snap/go/10455/src/net/cgo_unix.go:324
		_go_fuzz_dep_.CoverTab[4723]++
//line /snap/go/10455/src/net/cgo_unix.go:324
		// _ = "end of CoverTab[4723]"
//line /snap/go/10455/src/net/cgo_unix.go:324
	}
//line /snap/go/10455/src/net/cgo_unix.go:324
	// _ = "end of CoverTab[4716]"
//line /snap/go/10455/src/net/cgo_unix.go:324
	_go_fuzz_dep_.CoverTab[4717]++
						defer _C_res_nclose(state)

//line /snap/go/10455/src/net/cgo_unix.go:335
	bufSize := maxDNSPacketSize
	buf := (*_C_uchar)(_C_malloc(uintptr(bufSize)))
	defer _C_free(unsafe.Pointer(buf))

	s, err := syscall.BytePtrFromString(hostname)
	if err != nil {
//line /snap/go/10455/src/net/cgo_unix.go:340
		_go_fuzz_dep_.CoverTab[527511]++
//line /snap/go/10455/src/net/cgo_unix.go:340
		_go_fuzz_dep_.CoverTab[4724]++
							return nil, err
//line /snap/go/10455/src/net/cgo_unix.go:341
		// _ = "end of CoverTab[4724]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:342
		_go_fuzz_dep_.CoverTab[527512]++
//line /snap/go/10455/src/net/cgo_unix.go:342
		_go_fuzz_dep_.CoverTab[4725]++
//line /snap/go/10455/src/net/cgo_unix.go:342
		// _ = "end of CoverTab[4725]"
//line /snap/go/10455/src/net/cgo_unix.go:342
	}
//line /snap/go/10455/src/net/cgo_unix.go:342
	// _ = "end of CoverTab[4717]"
//line /snap/go/10455/src/net/cgo_unix.go:342
	_go_fuzz_dep_.CoverTab[4718]++

						var size int
//line /snap/go/10455/src/net/cgo_unix.go:344
	_go_fuzz_dep_.CoverTab[786646] = 0
						for {
//line /snap/go/10455/src/net/cgo_unix.go:345
		if _go_fuzz_dep_.CoverTab[786646] == 0 {
//line /snap/go/10455/src/net/cgo_unix.go:345
			_go_fuzz_dep_.CoverTab[527545]++
//line /snap/go/10455/src/net/cgo_unix.go:345
		} else {
//line /snap/go/10455/src/net/cgo_unix.go:345
			_go_fuzz_dep_.CoverTab[527546]++
//line /snap/go/10455/src/net/cgo_unix.go:345
		}
//line /snap/go/10455/src/net/cgo_unix.go:345
		_go_fuzz_dep_.CoverTab[786646] = 1
//line /snap/go/10455/src/net/cgo_unix.go:345
		_go_fuzz_dep_.CoverTab[4726]++
							size, _ = _C_res_nsearch(state, (*_C_char)(unsafe.Pointer(s)), class, rtype, buf, bufSize)
							if size <= 0 || func() bool {
//line /snap/go/10455/src/net/cgo_unix.go:347
			_go_fuzz_dep_.CoverTab[4729]++
//line /snap/go/10455/src/net/cgo_unix.go:347
			return size > 0xffff
//line /snap/go/10455/src/net/cgo_unix.go:347
			// _ = "end of CoverTab[4729]"
//line /snap/go/10455/src/net/cgo_unix.go:347
		}() {
//line /snap/go/10455/src/net/cgo_unix.go:347
			_go_fuzz_dep_.CoverTab[527513]++
//line /snap/go/10455/src/net/cgo_unix.go:347
			_go_fuzz_dep_.CoverTab[4730]++
								return nil, errors.New("res_nsearch failure")
//line /snap/go/10455/src/net/cgo_unix.go:348
			// _ = "end of CoverTab[4730]"
		} else {
//line /snap/go/10455/src/net/cgo_unix.go:349
			_go_fuzz_dep_.CoverTab[527514]++
//line /snap/go/10455/src/net/cgo_unix.go:349
			_go_fuzz_dep_.CoverTab[4731]++
//line /snap/go/10455/src/net/cgo_unix.go:349
			// _ = "end of CoverTab[4731]"
//line /snap/go/10455/src/net/cgo_unix.go:349
		}
//line /snap/go/10455/src/net/cgo_unix.go:349
		// _ = "end of CoverTab[4726]"
//line /snap/go/10455/src/net/cgo_unix.go:349
		_go_fuzz_dep_.CoverTab[4727]++
							if size <= bufSize {
//line /snap/go/10455/src/net/cgo_unix.go:350
			_go_fuzz_dep_.CoverTab[527515]++
//line /snap/go/10455/src/net/cgo_unix.go:350
			_go_fuzz_dep_.CoverTab[4732]++
								break
//line /snap/go/10455/src/net/cgo_unix.go:351
			// _ = "end of CoverTab[4732]"
		} else {
//line /snap/go/10455/src/net/cgo_unix.go:352
			_go_fuzz_dep_.CoverTab[527516]++
//line /snap/go/10455/src/net/cgo_unix.go:352
			_go_fuzz_dep_.CoverTab[4733]++
//line /snap/go/10455/src/net/cgo_unix.go:352
			// _ = "end of CoverTab[4733]"
//line /snap/go/10455/src/net/cgo_unix.go:352
		}
//line /snap/go/10455/src/net/cgo_unix.go:352
		// _ = "end of CoverTab[4727]"
//line /snap/go/10455/src/net/cgo_unix.go:352
		_go_fuzz_dep_.CoverTab[4728]++

//line /snap/go/10455/src/net/cgo_unix.go:355
		_C_free(unsafe.Pointer(buf))
							bufSize = size
							buf = (*_C_uchar)(_C_malloc(uintptr(bufSize)))
//line /snap/go/10455/src/net/cgo_unix.go:357
		// _ = "end of CoverTab[4728]"
	}
//line /snap/go/10455/src/net/cgo_unix.go:358
	// _ = "end of CoverTab[4718]"
//line /snap/go/10455/src/net/cgo_unix.go:358
	_go_fuzz_dep_.CoverTab[4719]++

						var p dnsmessage.Parser
						if _, err := p.Start(unsafe.Slice((*byte)(unsafe.Pointer(buf)), size)); err != nil {
//line /snap/go/10455/src/net/cgo_unix.go:361
		_go_fuzz_dep_.CoverTab[527517]++
//line /snap/go/10455/src/net/cgo_unix.go:361
		_go_fuzz_dep_.CoverTab[4734]++
							return nil, err
//line /snap/go/10455/src/net/cgo_unix.go:362
		// _ = "end of CoverTab[4734]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:363
		_go_fuzz_dep_.CoverTab[527518]++
//line /snap/go/10455/src/net/cgo_unix.go:363
		_go_fuzz_dep_.CoverTab[4735]++
//line /snap/go/10455/src/net/cgo_unix.go:363
		// _ = "end of CoverTab[4735]"
//line /snap/go/10455/src/net/cgo_unix.go:363
	}
//line /snap/go/10455/src/net/cgo_unix.go:363
	// _ = "end of CoverTab[4719]"
//line /snap/go/10455/src/net/cgo_unix.go:363
	_go_fuzz_dep_.CoverTab[4720]++
						p.SkipAllQuestions()
						resources, err := p.AllAnswers()
						if err != nil {
//line /snap/go/10455/src/net/cgo_unix.go:366
		_go_fuzz_dep_.CoverTab[527519]++
//line /snap/go/10455/src/net/cgo_unix.go:366
		_go_fuzz_dep_.CoverTab[4736]++
							return nil, err
//line /snap/go/10455/src/net/cgo_unix.go:367
		// _ = "end of CoverTab[4736]"
	} else {
//line /snap/go/10455/src/net/cgo_unix.go:368
		_go_fuzz_dep_.CoverTab[527520]++
//line /snap/go/10455/src/net/cgo_unix.go:368
		_go_fuzz_dep_.CoverTab[4737]++
//line /snap/go/10455/src/net/cgo_unix.go:368
		// _ = "end of CoverTab[4737]"
//line /snap/go/10455/src/net/cgo_unix.go:368
	}
//line /snap/go/10455/src/net/cgo_unix.go:368
	// _ = "end of CoverTab[4720]"
//line /snap/go/10455/src/net/cgo_unix.go:368
	_go_fuzz_dep_.CoverTab[4721]++
						return resources, nil
//line /snap/go/10455/src/net/cgo_unix.go:369
	// _ = "end of CoverTab[4721]"
}

//line /snap/go/10455/src/net/cgo_unix.go:370
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/cgo_unix.go:370
var _ = _go_fuzz_dep_.CoverTab
