// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is called cgo_unix.go, but to allow syscalls-to-libc-based
// implementations to share the code, it does not use cgo directly.
// Instead of C.foo it uses _C_foo, which is defined in either
// cgo_unix_cgo.go or cgo_unix_syscall.go

//go:build !netgo && ((cgo && unix) || darwin)

//line /usr/local/go/src/net/cgo_unix.go:12
package net

//line /usr/local/go/src/net/cgo_unix.go:12
import (
//line /usr/local/go/src/net/cgo_unix.go:12
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/cgo_unix.go:12
)
//line /usr/local/go/src/net/cgo_unix.go:12
import (
//line /usr/local/go/src/net/cgo_unix.go:12
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/cgo_unix.go:12
)

import (
	"context"
	"errors"
	"syscall"
	"unsafe"

	"golang.org/x/net/dns/dnsmessage"
)

// An addrinfoErrno represents a getaddrinfo, getnameinfo-specific
//line /usr/local/go/src/net/cgo_unix.go:23
// error number. It's a signed number and a zero value is a non-error
//line /usr/local/go/src/net/cgo_unix.go:23
// by convention.
//line /usr/local/go/src/net/cgo_unix.go:26
type addrinfoErrno int

func (eai addrinfoErrno) Error() string {
//line /usr/local/go/src/net/cgo_unix.go:28
	_go_fuzz_dep_.CoverTab[12655]++
//line /usr/local/go/src/net/cgo_unix.go:28
	return _C_gai_strerror(_C_int(eai))
//line /usr/local/go/src/net/cgo_unix.go:28
	// _ = "end of CoverTab[12655]"
//line /usr/local/go/src/net/cgo_unix.go:28
}
func (eai addrinfoErrno) Temporary() bool {
//line /usr/local/go/src/net/cgo_unix.go:29
	_go_fuzz_dep_.CoverTab[12656]++
//line /usr/local/go/src/net/cgo_unix.go:29
	return eai == _C_EAI_AGAIN
//line /usr/local/go/src/net/cgo_unix.go:29
	// _ = "end of CoverTab[12656]"
//line /usr/local/go/src/net/cgo_unix.go:29
}
func (eai addrinfoErrno) Timeout() bool {
//line /usr/local/go/src/net/cgo_unix.go:30
	_go_fuzz_dep_.CoverTab[12657]++
//line /usr/local/go/src/net/cgo_unix.go:30
	return false
//line /usr/local/go/src/net/cgo_unix.go:30
	// _ = "end of CoverTab[12657]"
//line /usr/local/go/src/net/cgo_unix.go:30
}

type portLookupResult struct {
	port	int
	err	error
}

type ipLookupResult struct {
	addrs	[]IPAddr
	cname	string
	err	error
}

type reverseLookupResult struct {
	names	[]string
	err	error
}

func cgoLookupHost(ctx context.Context, name string) (hosts []string, err error, completed bool) {
//line /usr/local/go/src/net/cgo_unix.go:48
	_go_fuzz_dep_.CoverTab[12658]++
						addrs, err, completed := cgoLookupIP(ctx, "ip", name)
						for _, addr := range addrs {
//line /usr/local/go/src/net/cgo_unix.go:50
		_go_fuzz_dep_.CoverTab[12660]++
							hosts = append(hosts, addr.String())
//line /usr/local/go/src/net/cgo_unix.go:51
		// _ = "end of CoverTab[12660]"
	}
//line /usr/local/go/src/net/cgo_unix.go:52
	// _ = "end of CoverTab[12658]"
//line /usr/local/go/src/net/cgo_unix.go:52
	_go_fuzz_dep_.CoverTab[12659]++
						return
//line /usr/local/go/src/net/cgo_unix.go:53
	// _ = "end of CoverTab[12659]"
}

func cgoLookupPort(ctx context.Context, network, service string) (port int, err error, completed bool) {
//line /usr/local/go/src/net/cgo_unix.go:56
	_go_fuzz_dep_.CoverTab[12661]++
						var hints _C_struct_addrinfo
						switch network {
	case "":
//line /usr/local/go/src/net/cgo_unix.go:59
		_go_fuzz_dep_.CoverTab[12665]++
//line /usr/local/go/src/net/cgo_unix.go:59
		// _ = "end of CoverTab[12665]"
	case "tcp", "tcp4", "tcp6":
//line /usr/local/go/src/net/cgo_unix.go:60
		_go_fuzz_dep_.CoverTab[12666]++
							*_C_ai_socktype(&hints) = _C_SOCK_STREAM
							*_C_ai_protocol(&hints) = _C_IPPROTO_TCP
//line /usr/local/go/src/net/cgo_unix.go:62
		// _ = "end of CoverTab[12666]"
	case "udp", "udp4", "udp6":
//line /usr/local/go/src/net/cgo_unix.go:63
		_go_fuzz_dep_.CoverTab[12667]++
							*_C_ai_socktype(&hints) = _C_SOCK_DGRAM
							*_C_ai_protocol(&hints) = _C_IPPROTO_UDP
//line /usr/local/go/src/net/cgo_unix.go:65
		// _ = "end of CoverTab[12667]"
	default:
//line /usr/local/go/src/net/cgo_unix.go:66
		_go_fuzz_dep_.CoverTab[12668]++
							return 0, &DNSError{Err: "unknown network", Name: network + "/" + service}, true
//line /usr/local/go/src/net/cgo_unix.go:67
		// _ = "end of CoverTab[12668]"
	}
//line /usr/local/go/src/net/cgo_unix.go:68
	// _ = "end of CoverTab[12661]"
//line /usr/local/go/src/net/cgo_unix.go:68
	_go_fuzz_dep_.CoverTab[12662]++
						switch ipVersion(network) {
	case '4':
//line /usr/local/go/src/net/cgo_unix.go:70
		_go_fuzz_dep_.CoverTab[12669]++
							*_C_ai_family(&hints) = _C_AF_INET
//line /usr/local/go/src/net/cgo_unix.go:71
		// _ = "end of CoverTab[12669]"
	case '6':
//line /usr/local/go/src/net/cgo_unix.go:72
		_go_fuzz_dep_.CoverTab[12670]++
							*_C_ai_family(&hints) = _C_AF_INET6
//line /usr/local/go/src/net/cgo_unix.go:73
		// _ = "end of CoverTab[12670]"
//line /usr/local/go/src/net/cgo_unix.go:73
	default:
//line /usr/local/go/src/net/cgo_unix.go:73
		_go_fuzz_dep_.CoverTab[12671]++
//line /usr/local/go/src/net/cgo_unix.go:73
		// _ = "end of CoverTab[12671]"
	}
//line /usr/local/go/src/net/cgo_unix.go:74
	// _ = "end of CoverTab[12662]"
//line /usr/local/go/src/net/cgo_unix.go:74
	_go_fuzz_dep_.CoverTab[12663]++
						if ctx.Done() == nil {
//line /usr/local/go/src/net/cgo_unix.go:75
		_go_fuzz_dep_.CoverTab[12672]++
							port, err := cgoLookupServicePort(&hints, network, service)
							return port, err, true
//line /usr/local/go/src/net/cgo_unix.go:77
		// _ = "end of CoverTab[12672]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:78
		_go_fuzz_dep_.CoverTab[12673]++
//line /usr/local/go/src/net/cgo_unix.go:78
		// _ = "end of CoverTab[12673]"
//line /usr/local/go/src/net/cgo_unix.go:78
	}
//line /usr/local/go/src/net/cgo_unix.go:78
	// _ = "end of CoverTab[12663]"
//line /usr/local/go/src/net/cgo_unix.go:78
	_go_fuzz_dep_.CoverTab[12664]++
						result := make(chan portLookupResult, 1)
//line /usr/local/go/src/net/cgo_unix.go:79
	_curRoutineNum2_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/cgo_unix.go:79
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum2_)
						go func() {
//line /usr/local/go/src/net/cgo_unix.go:80
		_go_fuzz_dep_.CoverTab[12674]++
//line /usr/local/go/src/net/cgo_unix.go:80
		defer func() {
//line /usr/local/go/src/net/cgo_unix.go:80
			_go_fuzz_dep_.CoverTab[12675]++
//line /usr/local/go/src/net/cgo_unix.go:80
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum2_)
//line /usr/local/go/src/net/cgo_unix.go:80
			// _ = "end of CoverTab[12675]"
//line /usr/local/go/src/net/cgo_unix.go:80
		}()
//line /usr/local/go/src/net/cgo_unix.go:80
		cgoPortLookup(result, &hints, network, service)
//line /usr/local/go/src/net/cgo_unix.go:80
		// _ = "end of CoverTab[12674]"
//line /usr/local/go/src/net/cgo_unix.go:80
	}()
						select {
	case r := <-result:
//line /usr/local/go/src/net/cgo_unix.go:82
		_go_fuzz_dep_.CoverTab[12676]++
							return r.port, r.err, true
//line /usr/local/go/src/net/cgo_unix.go:83
		// _ = "end of CoverTab[12676]"
	case <-ctx.Done():
//line /usr/local/go/src/net/cgo_unix.go:84
		_go_fuzz_dep_.CoverTab[12677]++

//line /usr/local/go/src/net/cgo_unix.go:87
		return 0, mapErr(ctx.Err()), false
//line /usr/local/go/src/net/cgo_unix.go:87
		// _ = "end of CoverTab[12677]"
	}
//line /usr/local/go/src/net/cgo_unix.go:88
	// _ = "end of CoverTab[12664]"
}

func cgoLookupServicePort(hints *_C_struct_addrinfo, network, service string) (port int, err error) {
//line /usr/local/go/src/net/cgo_unix.go:91
	_go_fuzz_dep_.CoverTab[12678]++
						cservice := make([]byte, len(service)+1)
						copy(cservice, service)

						for i, b := range cservice[:len(service)] {
//line /usr/local/go/src/net/cgo_unix.go:95
		_go_fuzz_dep_.CoverTab[12682]++
							cservice[i] = lowerASCII(b)
//line /usr/local/go/src/net/cgo_unix.go:96
		// _ = "end of CoverTab[12682]"
	}
//line /usr/local/go/src/net/cgo_unix.go:97
	// _ = "end of CoverTab[12678]"
//line /usr/local/go/src/net/cgo_unix.go:97
	_go_fuzz_dep_.CoverTab[12679]++
						var res *_C_struct_addrinfo
						gerrno, err := _C_getaddrinfo(nil, (*_C_char)(unsafe.Pointer(&cservice[0])), hints, &res)
						if gerrno != 0 {
//line /usr/local/go/src/net/cgo_unix.go:100
		_go_fuzz_dep_.CoverTab[12683]++
							isTemporary := false
							switch gerrno {
		case _C_EAI_SYSTEM:
//line /usr/local/go/src/net/cgo_unix.go:103
			_go_fuzz_dep_.CoverTab[12685]++
								if err == nil {
//line /usr/local/go/src/net/cgo_unix.go:104
				_go_fuzz_dep_.CoverTab[12687]++
									err = syscall.EMFILE
//line /usr/local/go/src/net/cgo_unix.go:105
				// _ = "end of CoverTab[12687]"
			} else {
//line /usr/local/go/src/net/cgo_unix.go:106
				_go_fuzz_dep_.CoverTab[12688]++
//line /usr/local/go/src/net/cgo_unix.go:106
				// _ = "end of CoverTab[12688]"
//line /usr/local/go/src/net/cgo_unix.go:106
			}
//line /usr/local/go/src/net/cgo_unix.go:106
			// _ = "end of CoverTab[12685]"
		default:
//line /usr/local/go/src/net/cgo_unix.go:107
			_go_fuzz_dep_.CoverTab[12686]++
								err = addrinfoErrno(gerrno)
								isTemporary = addrinfoErrno(gerrno).Temporary()
//line /usr/local/go/src/net/cgo_unix.go:109
			// _ = "end of CoverTab[12686]"
		}
//line /usr/local/go/src/net/cgo_unix.go:110
		// _ = "end of CoverTab[12683]"
//line /usr/local/go/src/net/cgo_unix.go:110
		_go_fuzz_dep_.CoverTab[12684]++
							return 0, &DNSError{Err: err.Error(), Name: network + "/" + service, IsTemporary: isTemporary}
//line /usr/local/go/src/net/cgo_unix.go:111
		// _ = "end of CoverTab[12684]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:112
		_go_fuzz_dep_.CoverTab[12689]++
//line /usr/local/go/src/net/cgo_unix.go:112
		// _ = "end of CoverTab[12689]"
//line /usr/local/go/src/net/cgo_unix.go:112
	}
//line /usr/local/go/src/net/cgo_unix.go:112
	// _ = "end of CoverTab[12679]"
//line /usr/local/go/src/net/cgo_unix.go:112
	_go_fuzz_dep_.CoverTab[12680]++
						defer _C_freeaddrinfo(res)

						for r := res; r != nil; r = *_C_ai_next(r) {
//line /usr/local/go/src/net/cgo_unix.go:115
		_go_fuzz_dep_.CoverTab[12690]++
							switch *_C_ai_family(r) {
		case _C_AF_INET:
//line /usr/local/go/src/net/cgo_unix.go:117
			_go_fuzz_dep_.CoverTab[12691]++
								sa := (*syscall.RawSockaddrInet4)(unsafe.Pointer(*_C_ai_addr(r)))
								p := (*[2]byte)(unsafe.Pointer(&sa.Port))
								return int(p[0])<<8 | int(p[1]), nil
//line /usr/local/go/src/net/cgo_unix.go:120
			// _ = "end of CoverTab[12691]"
		case _C_AF_INET6:
//line /usr/local/go/src/net/cgo_unix.go:121
			_go_fuzz_dep_.CoverTab[12692]++
								sa := (*syscall.RawSockaddrInet6)(unsafe.Pointer(*_C_ai_addr(r)))
								p := (*[2]byte)(unsafe.Pointer(&sa.Port))
								return int(p[0])<<8 | int(p[1]), nil
//line /usr/local/go/src/net/cgo_unix.go:124
			// _ = "end of CoverTab[12692]"
//line /usr/local/go/src/net/cgo_unix.go:124
		default:
//line /usr/local/go/src/net/cgo_unix.go:124
			_go_fuzz_dep_.CoverTab[12693]++
//line /usr/local/go/src/net/cgo_unix.go:124
			// _ = "end of CoverTab[12693]"
		}
//line /usr/local/go/src/net/cgo_unix.go:125
		// _ = "end of CoverTab[12690]"
	}
//line /usr/local/go/src/net/cgo_unix.go:126
	// _ = "end of CoverTab[12680]"
//line /usr/local/go/src/net/cgo_unix.go:126
	_go_fuzz_dep_.CoverTab[12681]++
						return 0, &DNSError{Err: "unknown port", Name: network + "/" + service}
//line /usr/local/go/src/net/cgo_unix.go:127
	// _ = "end of CoverTab[12681]"
}

func cgoPortLookup(result chan<- portLookupResult, hints *_C_struct_addrinfo, network, service string) {
//line /usr/local/go/src/net/cgo_unix.go:130
	_go_fuzz_dep_.CoverTab[12694]++
						port, err := cgoLookupServicePort(hints, network, service)
						result <- portLookupResult{port, err}
//line /usr/local/go/src/net/cgo_unix.go:132
	// _ = "end of CoverTab[12694]"
}

func cgoLookupIPCNAME(network, name string) (addrs []IPAddr, cname string, err error) {
//line /usr/local/go/src/net/cgo_unix.go:135
	_go_fuzz_dep_.CoverTab[12695]++
						acquireThread()
						defer releaseThread()

						var hints _C_struct_addrinfo
						*_C_ai_flags(&hints) = cgoAddrInfoFlags
						*_C_ai_socktype(&hints) = _C_SOCK_STREAM
						*_C_ai_family(&hints) = _C_AF_UNSPEC
						switch ipVersion(network) {
	case '4':
//line /usr/local/go/src/net/cgo_unix.go:144
		_go_fuzz_dep_.CoverTab[12700]++
							*_C_ai_family(&hints) = _C_AF_INET
//line /usr/local/go/src/net/cgo_unix.go:145
		// _ = "end of CoverTab[12700]"
	case '6':
//line /usr/local/go/src/net/cgo_unix.go:146
		_go_fuzz_dep_.CoverTab[12701]++
							*_C_ai_family(&hints) = _C_AF_INET6
//line /usr/local/go/src/net/cgo_unix.go:147
		// _ = "end of CoverTab[12701]"
//line /usr/local/go/src/net/cgo_unix.go:147
	default:
//line /usr/local/go/src/net/cgo_unix.go:147
		_go_fuzz_dep_.CoverTab[12702]++
//line /usr/local/go/src/net/cgo_unix.go:147
		// _ = "end of CoverTab[12702]"
	}
//line /usr/local/go/src/net/cgo_unix.go:148
	// _ = "end of CoverTab[12695]"
//line /usr/local/go/src/net/cgo_unix.go:148
	_go_fuzz_dep_.CoverTab[12696]++

						h := make([]byte, len(name)+1)
						copy(h, name)
						var res *_C_struct_addrinfo
						gerrno, err := _C_getaddrinfo((*_C_char)(unsafe.Pointer(&h[0])), nil, &hints, &res)
						if gerrno != 0 {
//line /usr/local/go/src/net/cgo_unix.go:154
		_go_fuzz_dep_.CoverTab[12703]++
							isErrorNoSuchHost := false
							isTemporary := false
							switch gerrno {
		case _C_EAI_SYSTEM:
//line /usr/local/go/src/net/cgo_unix.go:158
			_go_fuzz_dep_.CoverTab[12705]++
								if err == nil {
//line /usr/local/go/src/net/cgo_unix.go:159
				_go_fuzz_dep_.CoverTab[12708]++

//line /usr/local/go/src/net/cgo_unix.go:167
				err = syscall.EMFILE
//line /usr/local/go/src/net/cgo_unix.go:167
				// _ = "end of CoverTab[12708]"
			} else {
//line /usr/local/go/src/net/cgo_unix.go:168
				_go_fuzz_dep_.CoverTab[12709]++
//line /usr/local/go/src/net/cgo_unix.go:168
				// _ = "end of CoverTab[12709]"
//line /usr/local/go/src/net/cgo_unix.go:168
			}
//line /usr/local/go/src/net/cgo_unix.go:168
			// _ = "end of CoverTab[12705]"
		case _C_EAI_NONAME:
//line /usr/local/go/src/net/cgo_unix.go:169
			_go_fuzz_dep_.CoverTab[12706]++
								err = errNoSuchHost
								isErrorNoSuchHost = true
//line /usr/local/go/src/net/cgo_unix.go:171
			// _ = "end of CoverTab[12706]"
		default:
//line /usr/local/go/src/net/cgo_unix.go:172
			_go_fuzz_dep_.CoverTab[12707]++
								err = addrinfoErrno(gerrno)
								isTemporary = addrinfoErrno(gerrno).Temporary()
//line /usr/local/go/src/net/cgo_unix.go:174
			// _ = "end of CoverTab[12707]"
		}
//line /usr/local/go/src/net/cgo_unix.go:175
		// _ = "end of CoverTab[12703]"
//line /usr/local/go/src/net/cgo_unix.go:175
		_go_fuzz_dep_.CoverTab[12704]++

							return nil, "", &DNSError{Err: err.Error(), Name: name, IsNotFound: isErrorNoSuchHost, IsTemporary: isTemporary}
//line /usr/local/go/src/net/cgo_unix.go:177
		// _ = "end of CoverTab[12704]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:178
		_go_fuzz_dep_.CoverTab[12710]++
//line /usr/local/go/src/net/cgo_unix.go:178
		// _ = "end of CoverTab[12710]"
//line /usr/local/go/src/net/cgo_unix.go:178
	}
//line /usr/local/go/src/net/cgo_unix.go:178
	// _ = "end of CoverTab[12696]"
//line /usr/local/go/src/net/cgo_unix.go:178
	_go_fuzz_dep_.CoverTab[12697]++
						defer _C_freeaddrinfo(res)

						if res != nil {
//line /usr/local/go/src/net/cgo_unix.go:181
		_go_fuzz_dep_.CoverTab[12711]++
							cname = _C_GoString(*_C_ai_canonname(res))
							if cname == "" {
//line /usr/local/go/src/net/cgo_unix.go:183
			_go_fuzz_dep_.CoverTab[12713]++
								cname = name
//line /usr/local/go/src/net/cgo_unix.go:184
			// _ = "end of CoverTab[12713]"
		} else {
//line /usr/local/go/src/net/cgo_unix.go:185
			_go_fuzz_dep_.CoverTab[12714]++
//line /usr/local/go/src/net/cgo_unix.go:185
			// _ = "end of CoverTab[12714]"
//line /usr/local/go/src/net/cgo_unix.go:185
		}
//line /usr/local/go/src/net/cgo_unix.go:185
		// _ = "end of CoverTab[12711]"
//line /usr/local/go/src/net/cgo_unix.go:185
		_go_fuzz_dep_.CoverTab[12712]++
							if len(cname) > 0 && func() bool {
//line /usr/local/go/src/net/cgo_unix.go:186
			_go_fuzz_dep_.CoverTab[12715]++
//line /usr/local/go/src/net/cgo_unix.go:186
			return cname[len(cname)-1] != '.'
//line /usr/local/go/src/net/cgo_unix.go:186
			// _ = "end of CoverTab[12715]"
//line /usr/local/go/src/net/cgo_unix.go:186
		}() {
//line /usr/local/go/src/net/cgo_unix.go:186
			_go_fuzz_dep_.CoverTab[12716]++
								cname += "."
//line /usr/local/go/src/net/cgo_unix.go:187
			// _ = "end of CoverTab[12716]"
		} else {
//line /usr/local/go/src/net/cgo_unix.go:188
			_go_fuzz_dep_.CoverTab[12717]++
//line /usr/local/go/src/net/cgo_unix.go:188
			// _ = "end of CoverTab[12717]"
//line /usr/local/go/src/net/cgo_unix.go:188
		}
//line /usr/local/go/src/net/cgo_unix.go:188
		// _ = "end of CoverTab[12712]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:189
		_go_fuzz_dep_.CoverTab[12718]++
//line /usr/local/go/src/net/cgo_unix.go:189
		// _ = "end of CoverTab[12718]"
//line /usr/local/go/src/net/cgo_unix.go:189
	}
//line /usr/local/go/src/net/cgo_unix.go:189
	// _ = "end of CoverTab[12697]"
//line /usr/local/go/src/net/cgo_unix.go:189
	_go_fuzz_dep_.CoverTab[12698]++
						for r := res; r != nil; r = *_C_ai_next(r) {
//line /usr/local/go/src/net/cgo_unix.go:190
		_go_fuzz_dep_.CoverTab[12719]++

							if *_C_ai_socktype(r) != _C_SOCK_STREAM {
//line /usr/local/go/src/net/cgo_unix.go:192
			_go_fuzz_dep_.CoverTab[12721]++
								continue
//line /usr/local/go/src/net/cgo_unix.go:193
			// _ = "end of CoverTab[12721]"
		} else {
//line /usr/local/go/src/net/cgo_unix.go:194
			_go_fuzz_dep_.CoverTab[12722]++
//line /usr/local/go/src/net/cgo_unix.go:194
			// _ = "end of CoverTab[12722]"
//line /usr/local/go/src/net/cgo_unix.go:194
		}
//line /usr/local/go/src/net/cgo_unix.go:194
		// _ = "end of CoverTab[12719]"
//line /usr/local/go/src/net/cgo_unix.go:194
		_go_fuzz_dep_.CoverTab[12720]++
							switch *_C_ai_family(r) {
		case _C_AF_INET:
//line /usr/local/go/src/net/cgo_unix.go:196
			_go_fuzz_dep_.CoverTab[12723]++
								sa := (*syscall.RawSockaddrInet4)(unsafe.Pointer(*_C_ai_addr(r)))
								addr := IPAddr{IP: copyIP(sa.Addr[:])}
								addrs = append(addrs, addr)
//line /usr/local/go/src/net/cgo_unix.go:199
			// _ = "end of CoverTab[12723]"
		case _C_AF_INET6:
//line /usr/local/go/src/net/cgo_unix.go:200
			_go_fuzz_dep_.CoverTab[12724]++
								sa := (*syscall.RawSockaddrInet6)(unsafe.Pointer(*_C_ai_addr(r)))
								addr := IPAddr{IP: copyIP(sa.Addr[:]), Zone: zoneCache.name(int(sa.Scope_id))}
								addrs = append(addrs, addr)
//line /usr/local/go/src/net/cgo_unix.go:203
			// _ = "end of CoverTab[12724]"
//line /usr/local/go/src/net/cgo_unix.go:203
		default:
//line /usr/local/go/src/net/cgo_unix.go:203
			_go_fuzz_dep_.CoverTab[12725]++
//line /usr/local/go/src/net/cgo_unix.go:203
			// _ = "end of CoverTab[12725]"
		}
//line /usr/local/go/src/net/cgo_unix.go:204
		// _ = "end of CoverTab[12720]"
	}
//line /usr/local/go/src/net/cgo_unix.go:205
	// _ = "end of CoverTab[12698]"
//line /usr/local/go/src/net/cgo_unix.go:205
	_go_fuzz_dep_.CoverTab[12699]++
						return addrs, cname, nil
//line /usr/local/go/src/net/cgo_unix.go:206
	// _ = "end of CoverTab[12699]"
}

func cgoIPLookup(result chan<- ipLookupResult, network, name string) {
//line /usr/local/go/src/net/cgo_unix.go:209
	_go_fuzz_dep_.CoverTab[12726]++
						addrs, cname, err := cgoLookupIPCNAME(network, name)
						result <- ipLookupResult{addrs, cname, err}
//line /usr/local/go/src/net/cgo_unix.go:211
	// _ = "end of CoverTab[12726]"
}

func cgoLookupIP(ctx context.Context, network, name string) (addrs []IPAddr, err error, completed bool) {
//line /usr/local/go/src/net/cgo_unix.go:214
	_go_fuzz_dep_.CoverTab[12727]++
						if ctx.Done() == nil {
//line /usr/local/go/src/net/cgo_unix.go:215
		_go_fuzz_dep_.CoverTab[12729]++
							addrs, _, err = cgoLookupIPCNAME(network, name)
							return addrs, err, true
//line /usr/local/go/src/net/cgo_unix.go:217
		// _ = "end of CoverTab[12729]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:218
		_go_fuzz_dep_.CoverTab[12730]++
//line /usr/local/go/src/net/cgo_unix.go:218
		// _ = "end of CoverTab[12730]"
//line /usr/local/go/src/net/cgo_unix.go:218
	}
//line /usr/local/go/src/net/cgo_unix.go:218
	// _ = "end of CoverTab[12727]"
//line /usr/local/go/src/net/cgo_unix.go:218
	_go_fuzz_dep_.CoverTab[12728]++
						result := make(chan ipLookupResult, 1)
//line /usr/local/go/src/net/cgo_unix.go:219
	_curRoutineNum3_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/cgo_unix.go:219
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum3_)
						go func() {
//line /usr/local/go/src/net/cgo_unix.go:220
		_go_fuzz_dep_.CoverTab[12731]++
//line /usr/local/go/src/net/cgo_unix.go:220
		defer func() {
//line /usr/local/go/src/net/cgo_unix.go:220
			_go_fuzz_dep_.CoverTab[12732]++
//line /usr/local/go/src/net/cgo_unix.go:220
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum3_)
//line /usr/local/go/src/net/cgo_unix.go:220
			// _ = "end of CoverTab[12732]"
//line /usr/local/go/src/net/cgo_unix.go:220
		}()
//line /usr/local/go/src/net/cgo_unix.go:220
		cgoIPLookup(result, network, name)
//line /usr/local/go/src/net/cgo_unix.go:220
		// _ = "end of CoverTab[12731]"
//line /usr/local/go/src/net/cgo_unix.go:220
	}()
						select {
	case r := <-result:
//line /usr/local/go/src/net/cgo_unix.go:222
		_go_fuzz_dep_.CoverTab[12733]++
							return r.addrs, r.err, true
//line /usr/local/go/src/net/cgo_unix.go:223
		// _ = "end of CoverTab[12733]"
	case <-ctx.Done():
//line /usr/local/go/src/net/cgo_unix.go:224
		_go_fuzz_dep_.CoverTab[12734]++
							return nil, mapErr(ctx.Err()), false
//line /usr/local/go/src/net/cgo_unix.go:225
		// _ = "end of CoverTab[12734]"
	}
//line /usr/local/go/src/net/cgo_unix.go:226
	// _ = "end of CoverTab[12728]"
}

// These are roughly enough for the following:
//line /usr/local/go/src/net/cgo_unix.go:229
//
//line /usr/local/go/src/net/cgo_unix.go:229
//	 Source		Encoding			Maximum length of single name entry
//line /usr/local/go/src/net/cgo_unix.go:229
//	 Unicast DNS		ASCII or			<=253 + a NUL terminator
//line /usr/local/go/src/net/cgo_unix.go:229
//				Unicode in RFC 5892		252 * total number of labels + delimiters + a NUL terminator
//line /usr/local/go/src/net/cgo_unix.go:229
//	 Multicast DNS	UTF-8 in RFC 5198 or		<=253 + a NUL terminator
//line /usr/local/go/src/net/cgo_unix.go:229
//				the same as unicast DNS ASCII	<=253 + a NUL terminator
//line /usr/local/go/src/net/cgo_unix.go:229
//	 Local database	various				depends on implementation
//line /usr/local/go/src/net/cgo_unix.go:237
const (
	nameinfoLen	= 64
	maxNameinfoLen	= 4096
)

func cgoLookupPTR(ctx context.Context, addr string) (names []string, err error, completed bool) {
//line /usr/local/go/src/net/cgo_unix.go:242
	_go_fuzz_dep_.CoverTab[12735]++
						var zone string
						ip := parseIPv4(addr)
						if ip == nil {
//line /usr/local/go/src/net/cgo_unix.go:245
		_go_fuzz_dep_.CoverTab[12740]++
							ip, zone = parseIPv6Zone(addr)
//line /usr/local/go/src/net/cgo_unix.go:246
		// _ = "end of CoverTab[12740]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:247
		_go_fuzz_dep_.CoverTab[12741]++
//line /usr/local/go/src/net/cgo_unix.go:247
		// _ = "end of CoverTab[12741]"
//line /usr/local/go/src/net/cgo_unix.go:247
	}
//line /usr/local/go/src/net/cgo_unix.go:247
	// _ = "end of CoverTab[12735]"
//line /usr/local/go/src/net/cgo_unix.go:247
	_go_fuzz_dep_.CoverTab[12736]++
						if ip == nil {
//line /usr/local/go/src/net/cgo_unix.go:248
		_go_fuzz_dep_.CoverTab[12742]++
							return nil, &DNSError{Err: "invalid address", Name: addr}, true
//line /usr/local/go/src/net/cgo_unix.go:249
		// _ = "end of CoverTab[12742]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:250
		_go_fuzz_dep_.CoverTab[12743]++
//line /usr/local/go/src/net/cgo_unix.go:250
		// _ = "end of CoverTab[12743]"
//line /usr/local/go/src/net/cgo_unix.go:250
	}
//line /usr/local/go/src/net/cgo_unix.go:250
	// _ = "end of CoverTab[12736]"
//line /usr/local/go/src/net/cgo_unix.go:250
	_go_fuzz_dep_.CoverTab[12737]++
						sa, salen := cgoSockaddr(ip, zone)
						if sa == nil {
//line /usr/local/go/src/net/cgo_unix.go:252
		_go_fuzz_dep_.CoverTab[12744]++
							return nil, &DNSError{Err: "invalid address " + ip.String(), Name: addr}, true
//line /usr/local/go/src/net/cgo_unix.go:253
		// _ = "end of CoverTab[12744]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:254
		_go_fuzz_dep_.CoverTab[12745]++
//line /usr/local/go/src/net/cgo_unix.go:254
		// _ = "end of CoverTab[12745]"
//line /usr/local/go/src/net/cgo_unix.go:254
	}
//line /usr/local/go/src/net/cgo_unix.go:254
	// _ = "end of CoverTab[12737]"
//line /usr/local/go/src/net/cgo_unix.go:254
	_go_fuzz_dep_.CoverTab[12738]++
						if ctx.Done() == nil {
//line /usr/local/go/src/net/cgo_unix.go:255
		_go_fuzz_dep_.CoverTab[12746]++
							names, err := cgoLookupAddrPTR(addr, sa, salen)
							return names, err, true
//line /usr/local/go/src/net/cgo_unix.go:257
		// _ = "end of CoverTab[12746]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:258
		_go_fuzz_dep_.CoverTab[12747]++
//line /usr/local/go/src/net/cgo_unix.go:258
		// _ = "end of CoverTab[12747]"
//line /usr/local/go/src/net/cgo_unix.go:258
	}
//line /usr/local/go/src/net/cgo_unix.go:258
	// _ = "end of CoverTab[12738]"
//line /usr/local/go/src/net/cgo_unix.go:258
	_go_fuzz_dep_.CoverTab[12739]++
						result := make(chan reverseLookupResult, 1)
//line /usr/local/go/src/net/cgo_unix.go:259
	_curRoutineNum4_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/cgo_unix.go:259
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum4_)
						go func() {
//line /usr/local/go/src/net/cgo_unix.go:260
		_go_fuzz_dep_.CoverTab[12748]++
//line /usr/local/go/src/net/cgo_unix.go:260
		defer func() {
//line /usr/local/go/src/net/cgo_unix.go:260
			_go_fuzz_dep_.CoverTab[12749]++
//line /usr/local/go/src/net/cgo_unix.go:260
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum4_)
//line /usr/local/go/src/net/cgo_unix.go:260
			// _ = "end of CoverTab[12749]"
//line /usr/local/go/src/net/cgo_unix.go:260
		}()
//line /usr/local/go/src/net/cgo_unix.go:260
		cgoReverseLookup(result, addr, sa, salen)
//line /usr/local/go/src/net/cgo_unix.go:260
		// _ = "end of CoverTab[12748]"
//line /usr/local/go/src/net/cgo_unix.go:260
	}()
						select {
	case r := <-result:
//line /usr/local/go/src/net/cgo_unix.go:262
		_go_fuzz_dep_.CoverTab[12750]++
							return r.names, r.err, true
//line /usr/local/go/src/net/cgo_unix.go:263
		// _ = "end of CoverTab[12750]"
	case <-ctx.Done():
//line /usr/local/go/src/net/cgo_unix.go:264
		_go_fuzz_dep_.CoverTab[12751]++
							return nil, mapErr(ctx.Err()), false
//line /usr/local/go/src/net/cgo_unix.go:265
		// _ = "end of CoverTab[12751]"
	}
//line /usr/local/go/src/net/cgo_unix.go:266
	// _ = "end of CoverTab[12739]"
}

func cgoLookupAddrPTR(addr string, sa *_C_struct_sockaddr, salen _C_socklen_t) (names []string, err error) {
//line /usr/local/go/src/net/cgo_unix.go:269
	_go_fuzz_dep_.CoverTab[12752]++
						acquireThread()
						defer releaseThread()

						var gerrno int
						var b []byte
						for l := nameinfoLen; l <= maxNameinfoLen; l *= 2 {
//line /usr/local/go/src/net/cgo_unix.go:275
		_go_fuzz_dep_.CoverTab[12756]++
							b = make([]byte, l)
							gerrno, err = cgoNameinfoPTR(b, sa, salen)
							if gerrno == 0 || func() bool {
//line /usr/local/go/src/net/cgo_unix.go:278
			_go_fuzz_dep_.CoverTab[12757]++
//line /usr/local/go/src/net/cgo_unix.go:278
			return gerrno != _C_EAI_OVERFLOW
//line /usr/local/go/src/net/cgo_unix.go:278
			// _ = "end of CoverTab[12757]"
//line /usr/local/go/src/net/cgo_unix.go:278
		}() {
//line /usr/local/go/src/net/cgo_unix.go:278
			_go_fuzz_dep_.CoverTab[12758]++
								break
//line /usr/local/go/src/net/cgo_unix.go:279
			// _ = "end of CoverTab[12758]"
		} else {
//line /usr/local/go/src/net/cgo_unix.go:280
			_go_fuzz_dep_.CoverTab[12759]++
//line /usr/local/go/src/net/cgo_unix.go:280
			// _ = "end of CoverTab[12759]"
//line /usr/local/go/src/net/cgo_unix.go:280
		}
//line /usr/local/go/src/net/cgo_unix.go:280
		// _ = "end of CoverTab[12756]"
	}
//line /usr/local/go/src/net/cgo_unix.go:281
	// _ = "end of CoverTab[12752]"
//line /usr/local/go/src/net/cgo_unix.go:281
	_go_fuzz_dep_.CoverTab[12753]++
						if gerrno != 0 {
//line /usr/local/go/src/net/cgo_unix.go:282
		_go_fuzz_dep_.CoverTab[12760]++
							isTemporary := false
							switch gerrno {
		case _C_EAI_SYSTEM:
//line /usr/local/go/src/net/cgo_unix.go:285
			_go_fuzz_dep_.CoverTab[12762]++
								if err == nil {
//line /usr/local/go/src/net/cgo_unix.go:286
				_go_fuzz_dep_.CoverTab[12764]++
									err = syscall.EMFILE
//line /usr/local/go/src/net/cgo_unix.go:287
				// _ = "end of CoverTab[12764]"
			} else {
//line /usr/local/go/src/net/cgo_unix.go:288
				_go_fuzz_dep_.CoverTab[12765]++
//line /usr/local/go/src/net/cgo_unix.go:288
				// _ = "end of CoverTab[12765]"
//line /usr/local/go/src/net/cgo_unix.go:288
			}
//line /usr/local/go/src/net/cgo_unix.go:288
			// _ = "end of CoverTab[12762]"
		default:
//line /usr/local/go/src/net/cgo_unix.go:289
			_go_fuzz_dep_.CoverTab[12763]++
								err = addrinfoErrno(gerrno)
								isTemporary = addrinfoErrno(gerrno).Temporary()
//line /usr/local/go/src/net/cgo_unix.go:291
			// _ = "end of CoverTab[12763]"
		}
//line /usr/local/go/src/net/cgo_unix.go:292
		// _ = "end of CoverTab[12760]"
//line /usr/local/go/src/net/cgo_unix.go:292
		_go_fuzz_dep_.CoverTab[12761]++
							return nil, &DNSError{Err: err.Error(), Name: addr, IsTemporary: isTemporary}
//line /usr/local/go/src/net/cgo_unix.go:293
		// _ = "end of CoverTab[12761]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:294
		_go_fuzz_dep_.CoverTab[12766]++
//line /usr/local/go/src/net/cgo_unix.go:294
		// _ = "end of CoverTab[12766]"
//line /usr/local/go/src/net/cgo_unix.go:294
	}
//line /usr/local/go/src/net/cgo_unix.go:294
	// _ = "end of CoverTab[12753]"
//line /usr/local/go/src/net/cgo_unix.go:294
	_go_fuzz_dep_.CoverTab[12754]++
						for i := 0; i < len(b); i++ {
//line /usr/local/go/src/net/cgo_unix.go:295
		_go_fuzz_dep_.CoverTab[12767]++
							if b[i] == 0 {
//line /usr/local/go/src/net/cgo_unix.go:296
			_go_fuzz_dep_.CoverTab[12768]++
								b = b[:i]
								break
//line /usr/local/go/src/net/cgo_unix.go:298
			// _ = "end of CoverTab[12768]"
		} else {
//line /usr/local/go/src/net/cgo_unix.go:299
			_go_fuzz_dep_.CoverTab[12769]++
//line /usr/local/go/src/net/cgo_unix.go:299
			// _ = "end of CoverTab[12769]"
//line /usr/local/go/src/net/cgo_unix.go:299
		}
//line /usr/local/go/src/net/cgo_unix.go:299
		// _ = "end of CoverTab[12767]"
	}
//line /usr/local/go/src/net/cgo_unix.go:300
	// _ = "end of CoverTab[12754]"
//line /usr/local/go/src/net/cgo_unix.go:300
	_go_fuzz_dep_.CoverTab[12755]++
						return []string{absDomainName(string(b))}, nil
//line /usr/local/go/src/net/cgo_unix.go:301
	// _ = "end of CoverTab[12755]"
}

func cgoReverseLookup(result chan<- reverseLookupResult, addr string, sa *_C_struct_sockaddr, salen _C_socklen_t) {
//line /usr/local/go/src/net/cgo_unix.go:304
	_go_fuzz_dep_.CoverTab[12770]++
						names, err := cgoLookupAddrPTR(addr, sa, salen)
						result <- reverseLookupResult{names, err}
//line /usr/local/go/src/net/cgo_unix.go:306
	// _ = "end of CoverTab[12770]"
}

func cgoSockaddr(ip IP, zone string) (*_C_struct_sockaddr, _C_socklen_t) {
//line /usr/local/go/src/net/cgo_unix.go:309
	_go_fuzz_dep_.CoverTab[12771]++
						if ip4 := ip.To4(); ip4 != nil {
//line /usr/local/go/src/net/cgo_unix.go:310
		_go_fuzz_dep_.CoverTab[12774]++
							return cgoSockaddrInet4(ip4), _C_socklen_t(syscall.SizeofSockaddrInet4)
//line /usr/local/go/src/net/cgo_unix.go:311
		// _ = "end of CoverTab[12774]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:312
		_go_fuzz_dep_.CoverTab[12775]++
//line /usr/local/go/src/net/cgo_unix.go:312
		// _ = "end of CoverTab[12775]"
//line /usr/local/go/src/net/cgo_unix.go:312
	}
//line /usr/local/go/src/net/cgo_unix.go:312
	// _ = "end of CoverTab[12771]"
//line /usr/local/go/src/net/cgo_unix.go:312
	_go_fuzz_dep_.CoverTab[12772]++
						if ip6 := ip.To16(); ip6 != nil {
//line /usr/local/go/src/net/cgo_unix.go:313
		_go_fuzz_dep_.CoverTab[12776]++
							return cgoSockaddrInet6(ip6, zoneCache.index(zone)), _C_socklen_t(syscall.SizeofSockaddrInet6)
//line /usr/local/go/src/net/cgo_unix.go:314
		// _ = "end of CoverTab[12776]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:315
		_go_fuzz_dep_.CoverTab[12777]++
//line /usr/local/go/src/net/cgo_unix.go:315
		// _ = "end of CoverTab[12777]"
//line /usr/local/go/src/net/cgo_unix.go:315
	}
//line /usr/local/go/src/net/cgo_unix.go:315
	// _ = "end of CoverTab[12772]"
//line /usr/local/go/src/net/cgo_unix.go:315
	_go_fuzz_dep_.CoverTab[12773]++
						return nil, 0
//line /usr/local/go/src/net/cgo_unix.go:316
	// _ = "end of CoverTab[12773]"
}

func cgoLookupCNAME(ctx context.Context, name string) (cname string, err error, completed bool) {
//line /usr/local/go/src/net/cgo_unix.go:319
	_go_fuzz_dep_.CoverTab[12778]++
						resources, err := resSearch(ctx, name, int(dnsmessage.TypeCNAME), int(dnsmessage.ClassINET))
						if err != nil {
//line /usr/local/go/src/net/cgo_unix.go:321
		_go_fuzz_dep_.CoverTab[12781]++
							return
//line /usr/local/go/src/net/cgo_unix.go:322
		// _ = "end of CoverTab[12781]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:323
		_go_fuzz_dep_.CoverTab[12782]++
//line /usr/local/go/src/net/cgo_unix.go:323
		// _ = "end of CoverTab[12782]"
//line /usr/local/go/src/net/cgo_unix.go:323
	}
//line /usr/local/go/src/net/cgo_unix.go:323
	// _ = "end of CoverTab[12778]"
//line /usr/local/go/src/net/cgo_unix.go:323
	_go_fuzz_dep_.CoverTab[12779]++
						cname, err = parseCNAMEFromResources(resources)
						if err != nil {
//line /usr/local/go/src/net/cgo_unix.go:325
		_go_fuzz_dep_.CoverTab[12783]++
							return "", err, false
//line /usr/local/go/src/net/cgo_unix.go:326
		// _ = "end of CoverTab[12783]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:327
		_go_fuzz_dep_.CoverTab[12784]++
//line /usr/local/go/src/net/cgo_unix.go:327
		// _ = "end of CoverTab[12784]"
//line /usr/local/go/src/net/cgo_unix.go:327
	}
//line /usr/local/go/src/net/cgo_unix.go:327
	// _ = "end of CoverTab[12779]"
//line /usr/local/go/src/net/cgo_unix.go:327
	_go_fuzz_dep_.CoverTab[12780]++
						return cname, nil, true
//line /usr/local/go/src/net/cgo_unix.go:328
	// _ = "end of CoverTab[12780]"
}

// resSearch will make a call to the 'res_nsearch' routine in the C library
//line /usr/local/go/src/net/cgo_unix.go:331
// and parse the output as a slice of DNS resources.
//line /usr/local/go/src/net/cgo_unix.go:333
func resSearch(ctx context.Context, hostname string, rtype, class int) ([]dnsmessage.Resource, error) {
//line /usr/local/go/src/net/cgo_unix.go:333
	_go_fuzz_dep_.CoverTab[12785]++
						if ctx.Done() == nil {
//line /usr/local/go/src/net/cgo_unix.go:334
		_go_fuzz_dep_.CoverTab[12788]++
							return cgoResSearch(hostname, rtype, class)
//line /usr/local/go/src/net/cgo_unix.go:335
		// _ = "end of CoverTab[12788]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:336
		_go_fuzz_dep_.CoverTab[12789]++
//line /usr/local/go/src/net/cgo_unix.go:336
		// _ = "end of CoverTab[12789]"
//line /usr/local/go/src/net/cgo_unix.go:336
	}
//line /usr/local/go/src/net/cgo_unix.go:336
	// _ = "end of CoverTab[12785]"
//line /usr/local/go/src/net/cgo_unix.go:336
	_go_fuzz_dep_.CoverTab[12786]++

						type result struct {
		res	[]dnsmessage.Resource
		err	error
	}

						res := make(chan result, 1)
//line /usr/local/go/src/net/cgo_unix.go:343
	_curRoutineNum5_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/cgo_unix.go:343
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum5_)
						go func() {
//line /usr/local/go/src/net/cgo_unix.go:344
		_go_fuzz_dep_.CoverTab[12790]++
//line /usr/local/go/src/net/cgo_unix.go:344
		defer func() {
//line /usr/local/go/src/net/cgo_unix.go:344
			_go_fuzz_dep_.CoverTab[12791]++
//line /usr/local/go/src/net/cgo_unix.go:344
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum5_)
//line /usr/local/go/src/net/cgo_unix.go:344
			// _ = "end of CoverTab[12791]"
//line /usr/local/go/src/net/cgo_unix.go:344
		}()
							r, err := cgoResSearch(hostname, rtype, class)
							res <- result{
			res:	r,
			err:	err,
		}
//line /usr/local/go/src/net/cgo_unix.go:349
		// _ = "end of CoverTab[12790]"
	}()
//line /usr/local/go/src/net/cgo_unix.go:350
	// _ = "end of CoverTab[12786]"
//line /usr/local/go/src/net/cgo_unix.go:350
	_go_fuzz_dep_.CoverTab[12787]++

						select {
	case res := <-res:
//line /usr/local/go/src/net/cgo_unix.go:353
		_go_fuzz_dep_.CoverTab[12792]++
							return res.res, res.err
//line /usr/local/go/src/net/cgo_unix.go:354
		// _ = "end of CoverTab[12792]"
	case <-ctx.Done():
//line /usr/local/go/src/net/cgo_unix.go:355
		_go_fuzz_dep_.CoverTab[12793]++
							return nil, mapErr(ctx.Err())
//line /usr/local/go/src/net/cgo_unix.go:356
		// _ = "end of CoverTab[12793]"
	}
//line /usr/local/go/src/net/cgo_unix.go:357
	// _ = "end of CoverTab[12787]"
}

func cgoResSearch(hostname string, rtype, class int) ([]dnsmessage.Resource, error) {
//line /usr/local/go/src/net/cgo_unix.go:360
	_go_fuzz_dep_.CoverTab[12794]++
						acquireThread()
						defer releaseThread()

						state := (*_C_struct___res_state)(_C_malloc(unsafe.Sizeof(_C_struct___res_state{})))
						defer _C_free(unsafe.Pointer(state))
						if err := _C_res_ninit(state); err != nil {
//line /usr/local/go/src/net/cgo_unix.go:366
		_go_fuzz_dep_.CoverTab[12799]++
							return nil, errors.New("res_ninit failure: " + err.Error())
//line /usr/local/go/src/net/cgo_unix.go:367
		// _ = "end of CoverTab[12799]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:368
		_go_fuzz_dep_.CoverTab[12800]++
//line /usr/local/go/src/net/cgo_unix.go:368
		// _ = "end of CoverTab[12800]"
//line /usr/local/go/src/net/cgo_unix.go:368
	}
//line /usr/local/go/src/net/cgo_unix.go:368
	// _ = "end of CoverTab[12794]"
//line /usr/local/go/src/net/cgo_unix.go:368
	_go_fuzz_dep_.CoverTab[12795]++
						defer _C_res_nclose(state)

//line /usr/local/go/src/net/cgo_unix.go:379
	bufSize := maxDNSPacketSize
	buf := (*_C_uchar)(_C_malloc(uintptr(bufSize)))
	defer _C_free(unsafe.Pointer(buf))

	s := _C_CString(hostname)
	defer _C_FreeCString(s)

	var size int
	for {
//line /usr/local/go/src/net/cgo_unix.go:387
		_go_fuzz_dep_.CoverTab[12801]++
							size, _ = _C_res_nsearch(state, s, class, rtype, buf, bufSize)
							if size <= 0 || func() bool {
//line /usr/local/go/src/net/cgo_unix.go:389
			_go_fuzz_dep_.CoverTab[12804]++
//line /usr/local/go/src/net/cgo_unix.go:389
			return size > 0xffff
//line /usr/local/go/src/net/cgo_unix.go:389
			// _ = "end of CoverTab[12804]"
//line /usr/local/go/src/net/cgo_unix.go:389
		}() {
//line /usr/local/go/src/net/cgo_unix.go:389
			_go_fuzz_dep_.CoverTab[12805]++
								return nil, errors.New("res_nsearch failure")
//line /usr/local/go/src/net/cgo_unix.go:390
			// _ = "end of CoverTab[12805]"
		} else {
//line /usr/local/go/src/net/cgo_unix.go:391
			_go_fuzz_dep_.CoverTab[12806]++
//line /usr/local/go/src/net/cgo_unix.go:391
			// _ = "end of CoverTab[12806]"
//line /usr/local/go/src/net/cgo_unix.go:391
		}
//line /usr/local/go/src/net/cgo_unix.go:391
		// _ = "end of CoverTab[12801]"
//line /usr/local/go/src/net/cgo_unix.go:391
		_go_fuzz_dep_.CoverTab[12802]++
							if size <= bufSize {
//line /usr/local/go/src/net/cgo_unix.go:392
			_go_fuzz_dep_.CoverTab[12807]++
								break
//line /usr/local/go/src/net/cgo_unix.go:393
			// _ = "end of CoverTab[12807]"
		} else {
//line /usr/local/go/src/net/cgo_unix.go:394
			_go_fuzz_dep_.CoverTab[12808]++
//line /usr/local/go/src/net/cgo_unix.go:394
			// _ = "end of CoverTab[12808]"
//line /usr/local/go/src/net/cgo_unix.go:394
		}
//line /usr/local/go/src/net/cgo_unix.go:394
		// _ = "end of CoverTab[12802]"
//line /usr/local/go/src/net/cgo_unix.go:394
		_go_fuzz_dep_.CoverTab[12803]++

//line /usr/local/go/src/net/cgo_unix.go:397
		_C_free(unsafe.Pointer(buf))
							bufSize = size
							buf = (*_C_uchar)(_C_malloc(uintptr(bufSize)))
//line /usr/local/go/src/net/cgo_unix.go:399
		// _ = "end of CoverTab[12803]"
	}
//line /usr/local/go/src/net/cgo_unix.go:400
	// _ = "end of CoverTab[12795]"
//line /usr/local/go/src/net/cgo_unix.go:400
	_go_fuzz_dep_.CoverTab[12796]++

						var p dnsmessage.Parser
						if _, err := p.Start(unsafe.Slice((*byte)(unsafe.Pointer(buf)), size)); err != nil {
//line /usr/local/go/src/net/cgo_unix.go:403
		_go_fuzz_dep_.CoverTab[12809]++
							return nil, err
//line /usr/local/go/src/net/cgo_unix.go:404
		// _ = "end of CoverTab[12809]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:405
		_go_fuzz_dep_.CoverTab[12810]++
//line /usr/local/go/src/net/cgo_unix.go:405
		// _ = "end of CoverTab[12810]"
//line /usr/local/go/src/net/cgo_unix.go:405
	}
//line /usr/local/go/src/net/cgo_unix.go:405
	// _ = "end of CoverTab[12796]"
//line /usr/local/go/src/net/cgo_unix.go:405
	_go_fuzz_dep_.CoverTab[12797]++
						p.SkipAllQuestions()
						resources, err := p.AllAnswers()
						if err != nil {
//line /usr/local/go/src/net/cgo_unix.go:408
		_go_fuzz_dep_.CoverTab[12811]++
							return nil, err
//line /usr/local/go/src/net/cgo_unix.go:409
		// _ = "end of CoverTab[12811]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:410
		_go_fuzz_dep_.CoverTab[12812]++
//line /usr/local/go/src/net/cgo_unix.go:410
		// _ = "end of CoverTab[12812]"
//line /usr/local/go/src/net/cgo_unix.go:410
	}
//line /usr/local/go/src/net/cgo_unix.go:410
	// _ = "end of CoverTab[12797]"
//line /usr/local/go/src/net/cgo_unix.go:410
	_go_fuzz_dep_.CoverTab[12798]++
						return resources, nil
//line /usr/local/go/src/net/cgo_unix.go:411
	// _ = "end of CoverTab[12798]"
}

//line /usr/local/go/src/net/cgo_unix.go:412
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/cgo_unix.go:412
var _ = _go_fuzz_dep_.CoverTab
