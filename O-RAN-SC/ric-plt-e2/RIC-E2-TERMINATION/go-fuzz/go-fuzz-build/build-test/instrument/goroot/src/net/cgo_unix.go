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
	_go_fuzz_dep_.CoverTab[4265]++
//line /usr/local/go/src/net/cgo_unix.go:28
	return _C_gai_strerror(_C_int(eai))
//line /usr/local/go/src/net/cgo_unix.go:28
	// _ = "end of CoverTab[4265]"
//line /usr/local/go/src/net/cgo_unix.go:28
}
func (eai addrinfoErrno) Temporary() bool {
//line /usr/local/go/src/net/cgo_unix.go:29
	_go_fuzz_dep_.CoverTab[4266]++
//line /usr/local/go/src/net/cgo_unix.go:29
	return eai == _C_EAI_AGAIN
//line /usr/local/go/src/net/cgo_unix.go:29
	// _ = "end of CoverTab[4266]"
//line /usr/local/go/src/net/cgo_unix.go:29
}
func (eai addrinfoErrno) Timeout() bool {
//line /usr/local/go/src/net/cgo_unix.go:30
	_go_fuzz_dep_.CoverTab[4267]++
//line /usr/local/go/src/net/cgo_unix.go:30
	return false
//line /usr/local/go/src/net/cgo_unix.go:30
	// _ = "end of CoverTab[4267]"
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
	_go_fuzz_dep_.CoverTab[4268]++
						addrs, err, completed := cgoLookupIP(ctx, "ip", name)
						for _, addr := range addrs {
//line /usr/local/go/src/net/cgo_unix.go:50
		_go_fuzz_dep_.CoverTab[4270]++
							hosts = append(hosts, addr.String())
//line /usr/local/go/src/net/cgo_unix.go:51
		// _ = "end of CoverTab[4270]"
	}
//line /usr/local/go/src/net/cgo_unix.go:52
	// _ = "end of CoverTab[4268]"
//line /usr/local/go/src/net/cgo_unix.go:52
	_go_fuzz_dep_.CoverTab[4269]++
						return
//line /usr/local/go/src/net/cgo_unix.go:53
	// _ = "end of CoverTab[4269]"
}

func cgoLookupPort(ctx context.Context, network, service string) (port int, err error, completed bool) {
//line /usr/local/go/src/net/cgo_unix.go:56
	_go_fuzz_dep_.CoverTab[4271]++
						var hints _C_struct_addrinfo
						switch network {
	case "":
//line /usr/local/go/src/net/cgo_unix.go:59
		_go_fuzz_dep_.CoverTab[4275]++
//line /usr/local/go/src/net/cgo_unix.go:59
		// _ = "end of CoverTab[4275]"
	case "tcp", "tcp4", "tcp6":
//line /usr/local/go/src/net/cgo_unix.go:60
		_go_fuzz_dep_.CoverTab[4276]++
							*_C_ai_socktype(&hints) = _C_SOCK_STREAM
							*_C_ai_protocol(&hints) = _C_IPPROTO_TCP
//line /usr/local/go/src/net/cgo_unix.go:62
		// _ = "end of CoverTab[4276]"
	case "udp", "udp4", "udp6":
//line /usr/local/go/src/net/cgo_unix.go:63
		_go_fuzz_dep_.CoverTab[4277]++
							*_C_ai_socktype(&hints) = _C_SOCK_DGRAM
							*_C_ai_protocol(&hints) = _C_IPPROTO_UDP
//line /usr/local/go/src/net/cgo_unix.go:65
		// _ = "end of CoverTab[4277]"
	default:
//line /usr/local/go/src/net/cgo_unix.go:66
		_go_fuzz_dep_.CoverTab[4278]++
							return 0, &DNSError{Err: "unknown network", Name: network + "/" + service}, true
//line /usr/local/go/src/net/cgo_unix.go:67
		// _ = "end of CoverTab[4278]"
	}
//line /usr/local/go/src/net/cgo_unix.go:68
	// _ = "end of CoverTab[4271]"
//line /usr/local/go/src/net/cgo_unix.go:68
	_go_fuzz_dep_.CoverTab[4272]++
						switch ipVersion(network) {
	case '4':
//line /usr/local/go/src/net/cgo_unix.go:70
		_go_fuzz_dep_.CoverTab[4279]++
							*_C_ai_family(&hints) = _C_AF_INET
//line /usr/local/go/src/net/cgo_unix.go:71
		// _ = "end of CoverTab[4279]"
	case '6':
//line /usr/local/go/src/net/cgo_unix.go:72
		_go_fuzz_dep_.CoverTab[4280]++
							*_C_ai_family(&hints) = _C_AF_INET6
//line /usr/local/go/src/net/cgo_unix.go:73
		// _ = "end of CoverTab[4280]"
//line /usr/local/go/src/net/cgo_unix.go:73
	default:
//line /usr/local/go/src/net/cgo_unix.go:73
		_go_fuzz_dep_.CoverTab[4281]++
//line /usr/local/go/src/net/cgo_unix.go:73
		// _ = "end of CoverTab[4281]"
	}
//line /usr/local/go/src/net/cgo_unix.go:74
	// _ = "end of CoverTab[4272]"
//line /usr/local/go/src/net/cgo_unix.go:74
	_go_fuzz_dep_.CoverTab[4273]++
						if ctx.Done() == nil {
//line /usr/local/go/src/net/cgo_unix.go:75
		_go_fuzz_dep_.CoverTab[4282]++
							port, err := cgoLookupServicePort(&hints, network, service)
							return port, err, true
//line /usr/local/go/src/net/cgo_unix.go:77
		// _ = "end of CoverTab[4282]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:78
		_go_fuzz_dep_.CoverTab[4283]++
//line /usr/local/go/src/net/cgo_unix.go:78
		// _ = "end of CoverTab[4283]"
//line /usr/local/go/src/net/cgo_unix.go:78
	}
//line /usr/local/go/src/net/cgo_unix.go:78
	// _ = "end of CoverTab[4273]"
//line /usr/local/go/src/net/cgo_unix.go:78
	_go_fuzz_dep_.CoverTab[4274]++
						result := make(chan portLookupResult, 1)
//line /usr/local/go/src/net/cgo_unix.go:79
	_curRoutineNum2_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/cgo_unix.go:79
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum2_)
						go func() {
//line /usr/local/go/src/net/cgo_unix.go:80
		_go_fuzz_dep_.CoverTab[4284]++
//line /usr/local/go/src/net/cgo_unix.go:80
		defer func() {
//line /usr/local/go/src/net/cgo_unix.go:80
			_go_fuzz_dep_.CoverTab[4285]++
//line /usr/local/go/src/net/cgo_unix.go:80
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum2_)
//line /usr/local/go/src/net/cgo_unix.go:80
			// _ = "end of CoverTab[4285]"
//line /usr/local/go/src/net/cgo_unix.go:80
		}()
//line /usr/local/go/src/net/cgo_unix.go:80
		cgoPortLookup(result, &hints, network, service)
//line /usr/local/go/src/net/cgo_unix.go:80
		// _ = "end of CoverTab[4284]"
//line /usr/local/go/src/net/cgo_unix.go:80
	}()
						select {
	case r := <-result:
//line /usr/local/go/src/net/cgo_unix.go:82
		_go_fuzz_dep_.CoverTab[4286]++
							return r.port, r.err, true
//line /usr/local/go/src/net/cgo_unix.go:83
		// _ = "end of CoverTab[4286]"
	case <-ctx.Done():
//line /usr/local/go/src/net/cgo_unix.go:84
		_go_fuzz_dep_.CoverTab[4287]++

//line /usr/local/go/src/net/cgo_unix.go:87
		return 0, mapErr(ctx.Err()), false
//line /usr/local/go/src/net/cgo_unix.go:87
		// _ = "end of CoverTab[4287]"
	}
//line /usr/local/go/src/net/cgo_unix.go:88
	// _ = "end of CoverTab[4274]"
}

func cgoLookupServicePort(hints *_C_struct_addrinfo, network, service string) (port int, err error) {
//line /usr/local/go/src/net/cgo_unix.go:91
	_go_fuzz_dep_.CoverTab[4288]++
						cservice := make([]byte, len(service)+1)
						copy(cservice, service)

						for i, b := range cservice[:len(service)] {
//line /usr/local/go/src/net/cgo_unix.go:95
		_go_fuzz_dep_.CoverTab[4292]++
							cservice[i] = lowerASCII(b)
//line /usr/local/go/src/net/cgo_unix.go:96
		// _ = "end of CoverTab[4292]"
	}
//line /usr/local/go/src/net/cgo_unix.go:97
	// _ = "end of CoverTab[4288]"
//line /usr/local/go/src/net/cgo_unix.go:97
	_go_fuzz_dep_.CoverTab[4289]++
						var res *_C_struct_addrinfo
						gerrno, err := _C_getaddrinfo(nil, (*_C_char)(unsafe.Pointer(&cservice[0])), hints, &res)
						if gerrno != 0 {
//line /usr/local/go/src/net/cgo_unix.go:100
		_go_fuzz_dep_.CoverTab[4293]++
							isTemporary := false
							switch gerrno {
		case _C_EAI_SYSTEM:
//line /usr/local/go/src/net/cgo_unix.go:103
			_go_fuzz_dep_.CoverTab[4295]++
								if err == nil {
//line /usr/local/go/src/net/cgo_unix.go:104
				_go_fuzz_dep_.CoverTab[4297]++
									err = syscall.EMFILE
//line /usr/local/go/src/net/cgo_unix.go:105
				// _ = "end of CoverTab[4297]"
			} else {
//line /usr/local/go/src/net/cgo_unix.go:106
				_go_fuzz_dep_.CoverTab[4298]++
//line /usr/local/go/src/net/cgo_unix.go:106
				// _ = "end of CoverTab[4298]"
//line /usr/local/go/src/net/cgo_unix.go:106
			}
//line /usr/local/go/src/net/cgo_unix.go:106
			// _ = "end of CoverTab[4295]"
		default:
//line /usr/local/go/src/net/cgo_unix.go:107
			_go_fuzz_dep_.CoverTab[4296]++
								err = addrinfoErrno(gerrno)
								isTemporary = addrinfoErrno(gerrno).Temporary()
//line /usr/local/go/src/net/cgo_unix.go:109
			// _ = "end of CoverTab[4296]"
		}
//line /usr/local/go/src/net/cgo_unix.go:110
		// _ = "end of CoverTab[4293]"
//line /usr/local/go/src/net/cgo_unix.go:110
		_go_fuzz_dep_.CoverTab[4294]++
							return 0, &DNSError{Err: err.Error(), Name: network + "/" + service, IsTemporary: isTemporary}
//line /usr/local/go/src/net/cgo_unix.go:111
		// _ = "end of CoverTab[4294]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:112
		_go_fuzz_dep_.CoverTab[4299]++
//line /usr/local/go/src/net/cgo_unix.go:112
		// _ = "end of CoverTab[4299]"
//line /usr/local/go/src/net/cgo_unix.go:112
	}
//line /usr/local/go/src/net/cgo_unix.go:112
	// _ = "end of CoverTab[4289]"
//line /usr/local/go/src/net/cgo_unix.go:112
	_go_fuzz_dep_.CoverTab[4290]++
						defer _C_freeaddrinfo(res)

						for r := res; r != nil; r = *_C_ai_next(r) {
//line /usr/local/go/src/net/cgo_unix.go:115
		_go_fuzz_dep_.CoverTab[4300]++
							switch *_C_ai_family(r) {
		case _C_AF_INET:
//line /usr/local/go/src/net/cgo_unix.go:117
			_go_fuzz_dep_.CoverTab[4301]++
								sa := (*syscall.RawSockaddrInet4)(unsafe.Pointer(*_C_ai_addr(r)))
								p := (*[2]byte)(unsafe.Pointer(&sa.Port))
								return int(p[0])<<8 | int(p[1]), nil
//line /usr/local/go/src/net/cgo_unix.go:120
			// _ = "end of CoverTab[4301]"
		case _C_AF_INET6:
//line /usr/local/go/src/net/cgo_unix.go:121
			_go_fuzz_dep_.CoverTab[4302]++
								sa := (*syscall.RawSockaddrInet6)(unsafe.Pointer(*_C_ai_addr(r)))
								p := (*[2]byte)(unsafe.Pointer(&sa.Port))
								return int(p[0])<<8 | int(p[1]), nil
//line /usr/local/go/src/net/cgo_unix.go:124
			// _ = "end of CoverTab[4302]"
//line /usr/local/go/src/net/cgo_unix.go:124
		default:
//line /usr/local/go/src/net/cgo_unix.go:124
			_go_fuzz_dep_.CoverTab[4303]++
//line /usr/local/go/src/net/cgo_unix.go:124
			// _ = "end of CoverTab[4303]"
		}
//line /usr/local/go/src/net/cgo_unix.go:125
		// _ = "end of CoverTab[4300]"
	}
//line /usr/local/go/src/net/cgo_unix.go:126
	// _ = "end of CoverTab[4290]"
//line /usr/local/go/src/net/cgo_unix.go:126
	_go_fuzz_dep_.CoverTab[4291]++
						return 0, &DNSError{Err: "unknown port", Name: network + "/" + service}
//line /usr/local/go/src/net/cgo_unix.go:127
	// _ = "end of CoverTab[4291]"
}

func cgoPortLookup(result chan<- portLookupResult, hints *_C_struct_addrinfo, network, service string) {
//line /usr/local/go/src/net/cgo_unix.go:130
	_go_fuzz_dep_.CoverTab[4304]++
						port, err := cgoLookupServicePort(hints, network, service)
						result <- portLookupResult{port, err}
//line /usr/local/go/src/net/cgo_unix.go:132
	// _ = "end of CoverTab[4304]"
}

func cgoLookupIPCNAME(network, name string) (addrs []IPAddr, cname string, err error) {
//line /usr/local/go/src/net/cgo_unix.go:135
	_go_fuzz_dep_.CoverTab[4305]++
						acquireThread()
						defer releaseThread()

						var hints _C_struct_addrinfo
						*_C_ai_flags(&hints) = cgoAddrInfoFlags
						*_C_ai_socktype(&hints) = _C_SOCK_STREAM
						*_C_ai_family(&hints) = _C_AF_UNSPEC
						switch ipVersion(network) {
	case '4':
//line /usr/local/go/src/net/cgo_unix.go:144
		_go_fuzz_dep_.CoverTab[4310]++
							*_C_ai_family(&hints) = _C_AF_INET
//line /usr/local/go/src/net/cgo_unix.go:145
		// _ = "end of CoverTab[4310]"
	case '6':
//line /usr/local/go/src/net/cgo_unix.go:146
		_go_fuzz_dep_.CoverTab[4311]++
							*_C_ai_family(&hints) = _C_AF_INET6
//line /usr/local/go/src/net/cgo_unix.go:147
		// _ = "end of CoverTab[4311]"
//line /usr/local/go/src/net/cgo_unix.go:147
	default:
//line /usr/local/go/src/net/cgo_unix.go:147
		_go_fuzz_dep_.CoverTab[4312]++
//line /usr/local/go/src/net/cgo_unix.go:147
		// _ = "end of CoverTab[4312]"
	}
//line /usr/local/go/src/net/cgo_unix.go:148
	// _ = "end of CoverTab[4305]"
//line /usr/local/go/src/net/cgo_unix.go:148
	_go_fuzz_dep_.CoverTab[4306]++

						h := make([]byte, len(name)+1)
						copy(h, name)
						var res *_C_struct_addrinfo
						gerrno, err := _C_getaddrinfo((*_C_char)(unsafe.Pointer(&h[0])), nil, &hints, &res)
						if gerrno != 0 {
//line /usr/local/go/src/net/cgo_unix.go:154
		_go_fuzz_dep_.CoverTab[4313]++
							isErrorNoSuchHost := false
							isTemporary := false
							switch gerrno {
		case _C_EAI_SYSTEM:
//line /usr/local/go/src/net/cgo_unix.go:158
			_go_fuzz_dep_.CoverTab[4315]++
								if err == nil {
//line /usr/local/go/src/net/cgo_unix.go:159
				_go_fuzz_dep_.CoverTab[4318]++

//line /usr/local/go/src/net/cgo_unix.go:167
				err = syscall.EMFILE
//line /usr/local/go/src/net/cgo_unix.go:167
				// _ = "end of CoverTab[4318]"
			} else {
//line /usr/local/go/src/net/cgo_unix.go:168
				_go_fuzz_dep_.CoverTab[4319]++
//line /usr/local/go/src/net/cgo_unix.go:168
				// _ = "end of CoverTab[4319]"
//line /usr/local/go/src/net/cgo_unix.go:168
			}
//line /usr/local/go/src/net/cgo_unix.go:168
			// _ = "end of CoverTab[4315]"
		case _C_EAI_NONAME:
//line /usr/local/go/src/net/cgo_unix.go:169
			_go_fuzz_dep_.CoverTab[4316]++
								err = errNoSuchHost
								isErrorNoSuchHost = true
//line /usr/local/go/src/net/cgo_unix.go:171
			// _ = "end of CoverTab[4316]"
		default:
//line /usr/local/go/src/net/cgo_unix.go:172
			_go_fuzz_dep_.CoverTab[4317]++
								err = addrinfoErrno(gerrno)
								isTemporary = addrinfoErrno(gerrno).Temporary()
//line /usr/local/go/src/net/cgo_unix.go:174
			// _ = "end of CoverTab[4317]"
		}
//line /usr/local/go/src/net/cgo_unix.go:175
		// _ = "end of CoverTab[4313]"
//line /usr/local/go/src/net/cgo_unix.go:175
		_go_fuzz_dep_.CoverTab[4314]++

							return nil, "", &DNSError{Err: err.Error(), Name: name, IsNotFound: isErrorNoSuchHost, IsTemporary: isTemporary}
//line /usr/local/go/src/net/cgo_unix.go:177
		// _ = "end of CoverTab[4314]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:178
		_go_fuzz_dep_.CoverTab[4320]++
//line /usr/local/go/src/net/cgo_unix.go:178
		// _ = "end of CoverTab[4320]"
//line /usr/local/go/src/net/cgo_unix.go:178
	}
//line /usr/local/go/src/net/cgo_unix.go:178
	// _ = "end of CoverTab[4306]"
//line /usr/local/go/src/net/cgo_unix.go:178
	_go_fuzz_dep_.CoverTab[4307]++
						defer _C_freeaddrinfo(res)

						if res != nil {
//line /usr/local/go/src/net/cgo_unix.go:181
		_go_fuzz_dep_.CoverTab[4321]++
							cname = _C_GoString(*_C_ai_canonname(res))
							if cname == "" {
//line /usr/local/go/src/net/cgo_unix.go:183
			_go_fuzz_dep_.CoverTab[4323]++
								cname = name
//line /usr/local/go/src/net/cgo_unix.go:184
			// _ = "end of CoverTab[4323]"
		} else {
//line /usr/local/go/src/net/cgo_unix.go:185
			_go_fuzz_dep_.CoverTab[4324]++
//line /usr/local/go/src/net/cgo_unix.go:185
			// _ = "end of CoverTab[4324]"
//line /usr/local/go/src/net/cgo_unix.go:185
		}
//line /usr/local/go/src/net/cgo_unix.go:185
		// _ = "end of CoverTab[4321]"
//line /usr/local/go/src/net/cgo_unix.go:185
		_go_fuzz_dep_.CoverTab[4322]++
							if len(cname) > 0 && func() bool {
//line /usr/local/go/src/net/cgo_unix.go:186
			_go_fuzz_dep_.CoverTab[4325]++
//line /usr/local/go/src/net/cgo_unix.go:186
			return cname[len(cname)-1] != '.'
//line /usr/local/go/src/net/cgo_unix.go:186
			// _ = "end of CoverTab[4325]"
//line /usr/local/go/src/net/cgo_unix.go:186
		}() {
//line /usr/local/go/src/net/cgo_unix.go:186
			_go_fuzz_dep_.CoverTab[4326]++
								cname += "."
//line /usr/local/go/src/net/cgo_unix.go:187
			// _ = "end of CoverTab[4326]"
		} else {
//line /usr/local/go/src/net/cgo_unix.go:188
			_go_fuzz_dep_.CoverTab[4327]++
//line /usr/local/go/src/net/cgo_unix.go:188
			// _ = "end of CoverTab[4327]"
//line /usr/local/go/src/net/cgo_unix.go:188
		}
//line /usr/local/go/src/net/cgo_unix.go:188
		// _ = "end of CoverTab[4322]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:189
		_go_fuzz_dep_.CoverTab[4328]++
//line /usr/local/go/src/net/cgo_unix.go:189
		// _ = "end of CoverTab[4328]"
//line /usr/local/go/src/net/cgo_unix.go:189
	}
//line /usr/local/go/src/net/cgo_unix.go:189
	// _ = "end of CoverTab[4307]"
//line /usr/local/go/src/net/cgo_unix.go:189
	_go_fuzz_dep_.CoverTab[4308]++
						for r := res; r != nil; r = *_C_ai_next(r) {
//line /usr/local/go/src/net/cgo_unix.go:190
		_go_fuzz_dep_.CoverTab[4329]++

							if *_C_ai_socktype(r) != _C_SOCK_STREAM {
//line /usr/local/go/src/net/cgo_unix.go:192
			_go_fuzz_dep_.CoverTab[4331]++
								continue
//line /usr/local/go/src/net/cgo_unix.go:193
			// _ = "end of CoverTab[4331]"
		} else {
//line /usr/local/go/src/net/cgo_unix.go:194
			_go_fuzz_dep_.CoverTab[4332]++
//line /usr/local/go/src/net/cgo_unix.go:194
			// _ = "end of CoverTab[4332]"
//line /usr/local/go/src/net/cgo_unix.go:194
		}
//line /usr/local/go/src/net/cgo_unix.go:194
		// _ = "end of CoverTab[4329]"
//line /usr/local/go/src/net/cgo_unix.go:194
		_go_fuzz_dep_.CoverTab[4330]++
							switch *_C_ai_family(r) {
		case _C_AF_INET:
//line /usr/local/go/src/net/cgo_unix.go:196
			_go_fuzz_dep_.CoverTab[4333]++
								sa := (*syscall.RawSockaddrInet4)(unsafe.Pointer(*_C_ai_addr(r)))
								addr := IPAddr{IP: copyIP(sa.Addr[:])}
								addrs = append(addrs, addr)
//line /usr/local/go/src/net/cgo_unix.go:199
			// _ = "end of CoverTab[4333]"
		case _C_AF_INET6:
//line /usr/local/go/src/net/cgo_unix.go:200
			_go_fuzz_dep_.CoverTab[4334]++
								sa := (*syscall.RawSockaddrInet6)(unsafe.Pointer(*_C_ai_addr(r)))
								addr := IPAddr{IP: copyIP(sa.Addr[:]), Zone: zoneCache.name(int(sa.Scope_id))}
								addrs = append(addrs, addr)
//line /usr/local/go/src/net/cgo_unix.go:203
			// _ = "end of CoverTab[4334]"
//line /usr/local/go/src/net/cgo_unix.go:203
		default:
//line /usr/local/go/src/net/cgo_unix.go:203
			_go_fuzz_dep_.CoverTab[4335]++
//line /usr/local/go/src/net/cgo_unix.go:203
			// _ = "end of CoverTab[4335]"
		}
//line /usr/local/go/src/net/cgo_unix.go:204
		// _ = "end of CoverTab[4330]"
	}
//line /usr/local/go/src/net/cgo_unix.go:205
	// _ = "end of CoverTab[4308]"
//line /usr/local/go/src/net/cgo_unix.go:205
	_go_fuzz_dep_.CoverTab[4309]++
						return addrs, cname, nil
//line /usr/local/go/src/net/cgo_unix.go:206
	// _ = "end of CoverTab[4309]"
}

func cgoIPLookup(result chan<- ipLookupResult, network, name string) {
//line /usr/local/go/src/net/cgo_unix.go:209
	_go_fuzz_dep_.CoverTab[4336]++
						addrs, cname, err := cgoLookupIPCNAME(network, name)
						result <- ipLookupResult{addrs, cname, err}
//line /usr/local/go/src/net/cgo_unix.go:211
	// _ = "end of CoverTab[4336]"
}

func cgoLookupIP(ctx context.Context, network, name string) (addrs []IPAddr, err error, completed bool) {
//line /usr/local/go/src/net/cgo_unix.go:214
	_go_fuzz_dep_.CoverTab[4337]++
						if ctx.Done() == nil {
//line /usr/local/go/src/net/cgo_unix.go:215
		_go_fuzz_dep_.CoverTab[4339]++
							addrs, _, err = cgoLookupIPCNAME(network, name)
							return addrs, err, true
//line /usr/local/go/src/net/cgo_unix.go:217
		// _ = "end of CoverTab[4339]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:218
		_go_fuzz_dep_.CoverTab[4340]++
//line /usr/local/go/src/net/cgo_unix.go:218
		// _ = "end of CoverTab[4340]"
//line /usr/local/go/src/net/cgo_unix.go:218
	}
//line /usr/local/go/src/net/cgo_unix.go:218
	// _ = "end of CoverTab[4337]"
//line /usr/local/go/src/net/cgo_unix.go:218
	_go_fuzz_dep_.CoverTab[4338]++
						result := make(chan ipLookupResult, 1)
//line /usr/local/go/src/net/cgo_unix.go:219
	_curRoutineNum3_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/cgo_unix.go:219
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum3_)
						go func() {
//line /usr/local/go/src/net/cgo_unix.go:220
		_go_fuzz_dep_.CoverTab[4341]++
//line /usr/local/go/src/net/cgo_unix.go:220
		defer func() {
//line /usr/local/go/src/net/cgo_unix.go:220
			_go_fuzz_dep_.CoverTab[4342]++
//line /usr/local/go/src/net/cgo_unix.go:220
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum3_)
//line /usr/local/go/src/net/cgo_unix.go:220
			// _ = "end of CoverTab[4342]"
//line /usr/local/go/src/net/cgo_unix.go:220
		}()
//line /usr/local/go/src/net/cgo_unix.go:220
		cgoIPLookup(result, network, name)
//line /usr/local/go/src/net/cgo_unix.go:220
		// _ = "end of CoverTab[4341]"
//line /usr/local/go/src/net/cgo_unix.go:220
	}()
						select {
	case r := <-result:
//line /usr/local/go/src/net/cgo_unix.go:222
		_go_fuzz_dep_.CoverTab[4343]++
							return r.addrs, r.err, true
//line /usr/local/go/src/net/cgo_unix.go:223
		// _ = "end of CoverTab[4343]"
	case <-ctx.Done():
//line /usr/local/go/src/net/cgo_unix.go:224
		_go_fuzz_dep_.CoverTab[4344]++
							return nil, mapErr(ctx.Err()), false
//line /usr/local/go/src/net/cgo_unix.go:225
		// _ = "end of CoverTab[4344]"
	}
//line /usr/local/go/src/net/cgo_unix.go:226
	// _ = "end of CoverTab[4338]"
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
	_go_fuzz_dep_.CoverTab[4345]++
						var zone string
						ip := parseIPv4(addr)
						if ip == nil {
//line /usr/local/go/src/net/cgo_unix.go:245
		_go_fuzz_dep_.CoverTab[4350]++
							ip, zone = parseIPv6Zone(addr)
//line /usr/local/go/src/net/cgo_unix.go:246
		// _ = "end of CoverTab[4350]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:247
		_go_fuzz_dep_.CoverTab[4351]++
//line /usr/local/go/src/net/cgo_unix.go:247
		// _ = "end of CoverTab[4351]"
//line /usr/local/go/src/net/cgo_unix.go:247
	}
//line /usr/local/go/src/net/cgo_unix.go:247
	// _ = "end of CoverTab[4345]"
//line /usr/local/go/src/net/cgo_unix.go:247
	_go_fuzz_dep_.CoverTab[4346]++
						if ip == nil {
//line /usr/local/go/src/net/cgo_unix.go:248
		_go_fuzz_dep_.CoverTab[4352]++
							return nil, &DNSError{Err: "invalid address", Name: addr}, true
//line /usr/local/go/src/net/cgo_unix.go:249
		// _ = "end of CoverTab[4352]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:250
		_go_fuzz_dep_.CoverTab[4353]++
//line /usr/local/go/src/net/cgo_unix.go:250
		// _ = "end of CoverTab[4353]"
//line /usr/local/go/src/net/cgo_unix.go:250
	}
//line /usr/local/go/src/net/cgo_unix.go:250
	// _ = "end of CoverTab[4346]"
//line /usr/local/go/src/net/cgo_unix.go:250
	_go_fuzz_dep_.CoverTab[4347]++
						sa, salen := cgoSockaddr(ip, zone)
						if sa == nil {
//line /usr/local/go/src/net/cgo_unix.go:252
		_go_fuzz_dep_.CoverTab[4354]++
							return nil, &DNSError{Err: "invalid address " + ip.String(), Name: addr}, true
//line /usr/local/go/src/net/cgo_unix.go:253
		// _ = "end of CoverTab[4354]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:254
		_go_fuzz_dep_.CoverTab[4355]++
//line /usr/local/go/src/net/cgo_unix.go:254
		// _ = "end of CoverTab[4355]"
//line /usr/local/go/src/net/cgo_unix.go:254
	}
//line /usr/local/go/src/net/cgo_unix.go:254
	// _ = "end of CoverTab[4347]"
//line /usr/local/go/src/net/cgo_unix.go:254
	_go_fuzz_dep_.CoverTab[4348]++
						if ctx.Done() == nil {
//line /usr/local/go/src/net/cgo_unix.go:255
		_go_fuzz_dep_.CoverTab[4356]++
							names, err := cgoLookupAddrPTR(addr, sa, salen)
							return names, err, true
//line /usr/local/go/src/net/cgo_unix.go:257
		// _ = "end of CoverTab[4356]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:258
		_go_fuzz_dep_.CoverTab[4357]++
//line /usr/local/go/src/net/cgo_unix.go:258
		// _ = "end of CoverTab[4357]"
//line /usr/local/go/src/net/cgo_unix.go:258
	}
//line /usr/local/go/src/net/cgo_unix.go:258
	// _ = "end of CoverTab[4348]"
//line /usr/local/go/src/net/cgo_unix.go:258
	_go_fuzz_dep_.CoverTab[4349]++
						result := make(chan reverseLookupResult, 1)
//line /usr/local/go/src/net/cgo_unix.go:259
	_curRoutineNum4_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /usr/local/go/src/net/cgo_unix.go:259
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum4_)
						go func() {
//line /usr/local/go/src/net/cgo_unix.go:260
		_go_fuzz_dep_.CoverTab[4358]++
//line /usr/local/go/src/net/cgo_unix.go:260
		defer func() {
//line /usr/local/go/src/net/cgo_unix.go:260
			_go_fuzz_dep_.CoverTab[4359]++
//line /usr/local/go/src/net/cgo_unix.go:260
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum4_)
//line /usr/local/go/src/net/cgo_unix.go:260
			// _ = "end of CoverTab[4359]"
//line /usr/local/go/src/net/cgo_unix.go:260
		}()
//line /usr/local/go/src/net/cgo_unix.go:260
		cgoReverseLookup(result, addr, sa, salen)
//line /usr/local/go/src/net/cgo_unix.go:260
		// _ = "end of CoverTab[4358]"
//line /usr/local/go/src/net/cgo_unix.go:260
	}()
						select {
	case r := <-result:
//line /usr/local/go/src/net/cgo_unix.go:262
		_go_fuzz_dep_.CoverTab[4360]++
							return r.names, r.err, true
//line /usr/local/go/src/net/cgo_unix.go:263
		// _ = "end of CoverTab[4360]"
	case <-ctx.Done():
//line /usr/local/go/src/net/cgo_unix.go:264
		_go_fuzz_dep_.CoverTab[4361]++
							return nil, mapErr(ctx.Err()), false
//line /usr/local/go/src/net/cgo_unix.go:265
		// _ = "end of CoverTab[4361]"
	}
//line /usr/local/go/src/net/cgo_unix.go:266
	// _ = "end of CoverTab[4349]"
}

func cgoLookupAddrPTR(addr string, sa *_C_struct_sockaddr, salen _C_socklen_t) (names []string, err error) {
//line /usr/local/go/src/net/cgo_unix.go:269
	_go_fuzz_dep_.CoverTab[4362]++
						acquireThread()
						defer releaseThread()

						var gerrno int
						var b []byte
						for l := nameinfoLen; l <= maxNameinfoLen; l *= 2 {
//line /usr/local/go/src/net/cgo_unix.go:275
		_go_fuzz_dep_.CoverTab[4366]++
							b = make([]byte, l)
							gerrno, err = cgoNameinfoPTR(b, sa, salen)
							if gerrno == 0 || func() bool {
//line /usr/local/go/src/net/cgo_unix.go:278
			_go_fuzz_dep_.CoverTab[4367]++
//line /usr/local/go/src/net/cgo_unix.go:278
			return gerrno != _C_EAI_OVERFLOW
//line /usr/local/go/src/net/cgo_unix.go:278
			// _ = "end of CoverTab[4367]"
//line /usr/local/go/src/net/cgo_unix.go:278
		}() {
//line /usr/local/go/src/net/cgo_unix.go:278
			_go_fuzz_dep_.CoverTab[4368]++
								break
//line /usr/local/go/src/net/cgo_unix.go:279
			// _ = "end of CoverTab[4368]"
		} else {
//line /usr/local/go/src/net/cgo_unix.go:280
			_go_fuzz_dep_.CoverTab[4369]++
//line /usr/local/go/src/net/cgo_unix.go:280
			// _ = "end of CoverTab[4369]"
//line /usr/local/go/src/net/cgo_unix.go:280
		}
//line /usr/local/go/src/net/cgo_unix.go:280
		// _ = "end of CoverTab[4366]"
	}
//line /usr/local/go/src/net/cgo_unix.go:281
	// _ = "end of CoverTab[4362]"
//line /usr/local/go/src/net/cgo_unix.go:281
	_go_fuzz_dep_.CoverTab[4363]++
						if gerrno != 0 {
//line /usr/local/go/src/net/cgo_unix.go:282
		_go_fuzz_dep_.CoverTab[4370]++
							isTemporary := false
							switch gerrno {
		case _C_EAI_SYSTEM:
//line /usr/local/go/src/net/cgo_unix.go:285
			_go_fuzz_dep_.CoverTab[4372]++
								if err == nil {
//line /usr/local/go/src/net/cgo_unix.go:286
				_go_fuzz_dep_.CoverTab[4374]++
									err = syscall.EMFILE
//line /usr/local/go/src/net/cgo_unix.go:287
				// _ = "end of CoverTab[4374]"
			} else {
//line /usr/local/go/src/net/cgo_unix.go:288
				_go_fuzz_dep_.CoverTab[4375]++
//line /usr/local/go/src/net/cgo_unix.go:288
				// _ = "end of CoverTab[4375]"
//line /usr/local/go/src/net/cgo_unix.go:288
			}
//line /usr/local/go/src/net/cgo_unix.go:288
			// _ = "end of CoverTab[4372]"
		default:
//line /usr/local/go/src/net/cgo_unix.go:289
			_go_fuzz_dep_.CoverTab[4373]++
								err = addrinfoErrno(gerrno)
								isTemporary = addrinfoErrno(gerrno).Temporary()
//line /usr/local/go/src/net/cgo_unix.go:291
			// _ = "end of CoverTab[4373]"
		}
//line /usr/local/go/src/net/cgo_unix.go:292
		// _ = "end of CoverTab[4370]"
//line /usr/local/go/src/net/cgo_unix.go:292
		_go_fuzz_dep_.CoverTab[4371]++
							return nil, &DNSError{Err: err.Error(), Name: addr, IsTemporary: isTemporary}
//line /usr/local/go/src/net/cgo_unix.go:293
		// _ = "end of CoverTab[4371]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:294
		_go_fuzz_dep_.CoverTab[4376]++
//line /usr/local/go/src/net/cgo_unix.go:294
		// _ = "end of CoverTab[4376]"
//line /usr/local/go/src/net/cgo_unix.go:294
	}
//line /usr/local/go/src/net/cgo_unix.go:294
	// _ = "end of CoverTab[4363]"
//line /usr/local/go/src/net/cgo_unix.go:294
	_go_fuzz_dep_.CoverTab[4364]++
						for i := 0; i < len(b); i++ {
//line /usr/local/go/src/net/cgo_unix.go:295
		_go_fuzz_dep_.CoverTab[4377]++
							if b[i] == 0 {
//line /usr/local/go/src/net/cgo_unix.go:296
			_go_fuzz_dep_.CoverTab[4378]++
								b = b[:i]
								break
//line /usr/local/go/src/net/cgo_unix.go:298
			// _ = "end of CoverTab[4378]"
		} else {
//line /usr/local/go/src/net/cgo_unix.go:299
			_go_fuzz_dep_.CoverTab[4379]++
//line /usr/local/go/src/net/cgo_unix.go:299
			// _ = "end of CoverTab[4379]"
//line /usr/local/go/src/net/cgo_unix.go:299
		}
//line /usr/local/go/src/net/cgo_unix.go:299
		// _ = "end of CoverTab[4377]"
	}
//line /usr/local/go/src/net/cgo_unix.go:300
	// _ = "end of CoverTab[4364]"
//line /usr/local/go/src/net/cgo_unix.go:300
	_go_fuzz_dep_.CoverTab[4365]++
						return []string{absDomainName(string(b))}, nil
//line /usr/local/go/src/net/cgo_unix.go:301
	// _ = "end of CoverTab[4365]"
}

func cgoReverseLookup(result chan<- reverseLookupResult, addr string, sa *_C_struct_sockaddr, salen _C_socklen_t) {
//line /usr/local/go/src/net/cgo_unix.go:304
	_go_fuzz_dep_.CoverTab[4380]++
						names, err := cgoLookupAddrPTR(addr, sa, salen)
						result <- reverseLookupResult{names, err}
//line /usr/local/go/src/net/cgo_unix.go:306
	// _ = "end of CoverTab[4380]"
}

func cgoSockaddr(ip IP, zone string) (*_C_struct_sockaddr, _C_socklen_t) {
//line /usr/local/go/src/net/cgo_unix.go:309
	_go_fuzz_dep_.CoverTab[4381]++
						if ip4 := ip.To4(); ip4 != nil {
//line /usr/local/go/src/net/cgo_unix.go:310
		_go_fuzz_dep_.CoverTab[4384]++
							return cgoSockaddrInet4(ip4), _C_socklen_t(syscall.SizeofSockaddrInet4)
//line /usr/local/go/src/net/cgo_unix.go:311
		// _ = "end of CoverTab[4384]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:312
		_go_fuzz_dep_.CoverTab[4385]++
//line /usr/local/go/src/net/cgo_unix.go:312
		// _ = "end of CoverTab[4385]"
//line /usr/local/go/src/net/cgo_unix.go:312
	}
//line /usr/local/go/src/net/cgo_unix.go:312
	// _ = "end of CoverTab[4381]"
//line /usr/local/go/src/net/cgo_unix.go:312
	_go_fuzz_dep_.CoverTab[4382]++
						if ip6 := ip.To16(); ip6 != nil {
//line /usr/local/go/src/net/cgo_unix.go:313
		_go_fuzz_dep_.CoverTab[4386]++
							return cgoSockaddrInet6(ip6, zoneCache.index(zone)), _C_socklen_t(syscall.SizeofSockaddrInet6)
//line /usr/local/go/src/net/cgo_unix.go:314
		// _ = "end of CoverTab[4386]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:315
		_go_fuzz_dep_.CoverTab[4387]++
//line /usr/local/go/src/net/cgo_unix.go:315
		// _ = "end of CoverTab[4387]"
//line /usr/local/go/src/net/cgo_unix.go:315
	}
//line /usr/local/go/src/net/cgo_unix.go:315
	// _ = "end of CoverTab[4382]"
//line /usr/local/go/src/net/cgo_unix.go:315
	_go_fuzz_dep_.CoverTab[4383]++
						return nil, 0
//line /usr/local/go/src/net/cgo_unix.go:316
	// _ = "end of CoverTab[4383]"
}

func cgoLookupCNAME(ctx context.Context, name string) (cname string, err error, completed bool) {
//line /usr/local/go/src/net/cgo_unix.go:319
	_go_fuzz_dep_.CoverTab[4388]++
						resources, err := resSearch(ctx, name, int(dnsmessage.TypeCNAME), int(dnsmessage.ClassINET))
						if err != nil {
//line /usr/local/go/src/net/cgo_unix.go:321
		_go_fuzz_dep_.CoverTab[4391]++
							return
//line /usr/local/go/src/net/cgo_unix.go:322
		// _ = "end of CoverTab[4391]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:323
		_go_fuzz_dep_.CoverTab[4392]++
//line /usr/local/go/src/net/cgo_unix.go:323
		// _ = "end of CoverTab[4392]"
//line /usr/local/go/src/net/cgo_unix.go:323
	}
//line /usr/local/go/src/net/cgo_unix.go:323
	// _ = "end of CoverTab[4388]"
//line /usr/local/go/src/net/cgo_unix.go:323
	_go_fuzz_dep_.CoverTab[4389]++
						cname, err = parseCNAMEFromResources(resources)
						if err != nil {
//line /usr/local/go/src/net/cgo_unix.go:325
		_go_fuzz_dep_.CoverTab[4393]++
							return "", err, false
//line /usr/local/go/src/net/cgo_unix.go:326
		// _ = "end of CoverTab[4393]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:327
		_go_fuzz_dep_.CoverTab[4394]++
//line /usr/local/go/src/net/cgo_unix.go:327
		// _ = "end of CoverTab[4394]"
//line /usr/local/go/src/net/cgo_unix.go:327
	}
//line /usr/local/go/src/net/cgo_unix.go:327
	// _ = "end of CoverTab[4389]"
//line /usr/local/go/src/net/cgo_unix.go:327
	_go_fuzz_dep_.CoverTab[4390]++
						return cname, nil, true
//line /usr/local/go/src/net/cgo_unix.go:328
	// _ = "end of CoverTab[4390]"
}

// resSearch will make a call to the 'res_nsearch' routine in the C library
//line /usr/local/go/src/net/cgo_unix.go:331
// and parse the output as a slice of DNS resources.
//line /usr/local/go/src/net/cgo_unix.go:333
func resSearch(ctx context.Context, hostname string, rtype, class int) ([]dnsmessage.Resource, error) {
//line /usr/local/go/src/net/cgo_unix.go:333
	_go_fuzz_dep_.CoverTab[4395]++
						if ctx.Done() == nil {
//line /usr/local/go/src/net/cgo_unix.go:334
		_go_fuzz_dep_.CoverTab[4398]++
							return cgoResSearch(hostname, rtype, class)
//line /usr/local/go/src/net/cgo_unix.go:335
		// _ = "end of CoverTab[4398]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:336
		_go_fuzz_dep_.CoverTab[4399]++
//line /usr/local/go/src/net/cgo_unix.go:336
		// _ = "end of CoverTab[4399]"
//line /usr/local/go/src/net/cgo_unix.go:336
	}
//line /usr/local/go/src/net/cgo_unix.go:336
	// _ = "end of CoverTab[4395]"
//line /usr/local/go/src/net/cgo_unix.go:336
	_go_fuzz_dep_.CoverTab[4396]++

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
		_go_fuzz_dep_.CoverTab[4400]++
//line /usr/local/go/src/net/cgo_unix.go:344
		defer func() {
//line /usr/local/go/src/net/cgo_unix.go:344
			_go_fuzz_dep_.CoverTab[4401]++
//line /usr/local/go/src/net/cgo_unix.go:344
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum5_)
//line /usr/local/go/src/net/cgo_unix.go:344
			// _ = "end of CoverTab[4401]"
//line /usr/local/go/src/net/cgo_unix.go:344
		}()
							r, err := cgoResSearch(hostname, rtype, class)
							res <- result{
			res:	r,
			err:	err,
		}
//line /usr/local/go/src/net/cgo_unix.go:349
		// _ = "end of CoverTab[4400]"
	}()
//line /usr/local/go/src/net/cgo_unix.go:350
	// _ = "end of CoverTab[4396]"
//line /usr/local/go/src/net/cgo_unix.go:350
	_go_fuzz_dep_.CoverTab[4397]++

						select {
	case res := <-res:
//line /usr/local/go/src/net/cgo_unix.go:353
		_go_fuzz_dep_.CoverTab[4402]++
							return res.res, res.err
//line /usr/local/go/src/net/cgo_unix.go:354
		// _ = "end of CoverTab[4402]"
	case <-ctx.Done():
//line /usr/local/go/src/net/cgo_unix.go:355
		_go_fuzz_dep_.CoverTab[4403]++
							return nil, mapErr(ctx.Err())
//line /usr/local/go/src/net/cgo_unix.go:356
		// _ = "end of CoverTab[4403]"
	}
//line /usr/local/go/src/net/cgo_unix.go:357
	// _ = "end of CoverTab[4397]"
}

func cgoResSearch(hostname string, rtype, class int) ([]dnsmessage.Resource, error) {
//line /usr/local/go/src/net/cgo_unix.go:360
	_go_fuzz_dep_.CoverTab[4404]++
						acquireThread()
						defer releaseThread()

						state := (*_C_struct___res_state)(_C_malloc(unsafe.Sizeof(_C_struct___res_state{})))
						defer _C_free(unsafe.Pointer(state))
						if err := _C_res_ninit(state); err != nil {
//line /usr/local/go/src/net/cgo_unix.go:366
		_go_fuzz_dep_.CoverTab[4409]++
							return nil, errors.New("res_ninit failure: " + err.Error())
//line /usr/local/go/src/net/cgo_unix.go:367
		// _ = "end of CoverTab[4409]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:368
		_go_fuzz_dep_.CoverTab[4410]++
//line /usr/local/go/src/net/cgo_unix.go:368
		// _ = "end of CoverTab[4410]"
//line /usr/local/go/src/net/cgo_unix.go:368
	}
//line /usr/local/go/src/net/cgo_unix.go:368
	// _ = "end of CoverTab[4404]"
//line /usr/local/go/src/net/cgo_unix.go:368
	_go_fuzz_dep_.CoverTab[4405]++
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
		_go_fuzz_dep_.CoverTab[4411]++
							size, _ = _C_res_nsearch(state, s, class, rtype, buf, bufSize)
							if size <= 0 || func() bool {
//line /usr/local/go/src/net/cgo_unix.go:389
			_go_fuzz_dep_.CoverTab[4414]++
//line /usr/local/go/src/net/cgo_unix.go:389
			return size > 0xffff
//line /usr/local/go/src/net/cgo_unix.go:389
			// _ = "end of CoverTab[4414]"
//line /usr/local/go/src/net/cgo_unix.go:389
		}() {
//line /usr/local/go/src/net/cgo_unix.go:389
			_go_fuzz_dep_.CoverTab[4415]++
								return nil, errors.New("res_nsearch failure")
//line /usr/local/go/src/net/cgo_unix.go:390
			// _ = "end of CoverTab[4415]"
		} else {
//line /usr/local/go/src/net/cgo_unix.go:391
			_go_fuzz_dep_.CoverTab[4416]++
//line /usr/local/go/src/net/cgo_unix.go:391
			// _ = "end of CoverTab[4416]"
//line /usr/local/go/src/net/cgo_unix.go:391
		}
//line /usr/local/go/src/net/cgo_unix.go:391
		// _ = "end of CoverTab[4411]"
//line /usr/local/go/src/net/cgo_unix.go:391
		_go_fuzz_dep_.CoverTab[4412]++
							if size <= bufSize {
//line /usr/local/go/src/net/cgo_unix.go:392
			_go_fuzz_dep_.CoverTab[4417]++
								break
//line /usr/local/go/src/net/cgo_unix.go:393
			// _ = "end of CoverTab[4417]"
		} else {
//line /usr/local/go/src/net/cgo_unix.go:394
			_go_fuzz_dep_.CoverTab[4418]++
//line /usr/local/go/src/net/cgo_unix.go:394
			// _ = "end of CoverTab[4418]"
//line /usr/local/go/src/net/cgo_unix.go:394
		}
//line /usr/local/go/src/net/cgo_unix.go:394
		// _ = "end of CoverTab[4412]"
//line /usr/local/go/src/net/cgo_unix.go:394
		_go_fuzz_dep_.CoverTab[4413]++

//line /usr/local/go/src/net/cgo_unix.go:397
		_C_free(unsafe.Pointer(buf))
							bufSize = size
							buf = (*_C_uchar)(_C_malloc(uintptr(bufSize)))
//line /usr/local/go/src/net/cgo_unix.go:399
		// _ = "end of CoverTab[4413]"
	}
//line /usr/local/go/src/net/cgo_unix.go:400
	// _ = "end of CoverTab[4405]"
//line /usr/local/go/src/net/cgo_unix.go:400
	_go_fuzz_dep_.CoverTab[4406]++

						var p dnsmessage.Parser
						if _, err := p.Start(unsafe.Slice((*byte)(unsafe.Pointer(buf)), size)); err != nil {
//line /usr/local/go/src/net/cgo_unix.go:403
		_go_fuzz_dep_.CoverTab[4419]++
							return nil, err
//line /usr/local/go/src/net/cgo_unix.go:404
		// _ = "end of CoverTab[4419]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:405
		_go_fuzz_dep_.CoverTab[4420]++
//line /usr/local/go/src/net/cgo_unix.go:405
		// _ = "end of CoverTab[4420]"
//line /usr/local/go/src/net/cgo_unix.go:405
	}
//line /usr/local/go/src/net/cgo_unix.go:405
	// _ = "end of CoverTab[4406]"
//line /usr/local/go/src/net/cgo_unix.go:405
	_go_fuzz_dep_.CoverTab[4407]++
						p.SkipAllQuestions()
						resources, err := p.AllAnswers()
						if err != nil {
//line /usr/local/go/src/net/cgo_unix.go:408
		_go_fuzz_dep_.CoverTab[4421]++
							return nil, err
//line /usr/local/go/src/net/cgo_unix.go:409
		// _ = "end of CoverTab[4421]"
	} else {
//line /usr/local/go/src/net/cgo_unix.go:410
		_go_fuzz_dep_.CoverTab[4422]++
//line /usr/local/go/src/net/cgo_unix.go:410
		// _ = "end of CoverTab[4422]"
//line /usr/local/go/src/net/cgo_unix.go:410
	}
//line /usr/local/go/src/net/cgo_unix.go:410
	// _ = "end of CoverTab[4407]"
//line /usr/local/go/src/net/cgo_unix.go:410
	_go_fuzz_dep_.CoverTab[4408]++
						return resources, nil
//line /usr/local/go/src/net/cgo_unix.go:411
	// _ = "end of CoverTab[4408]"
}

//line /usr/local/go/src/net/cgo_unix.go:412
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/cgo_unix.go:412
var _ = _go_fuzz_dep_.CoverTab
