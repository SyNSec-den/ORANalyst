// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Socket control messages

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:7
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:7
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:7
)

import "unsafe"

// UnixCredentials encodes credentials into a socket control message
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:11
// for sending to another process. This can be used for
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:11
// authentication.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:14
func UnixCredentials(ucred *Ucred) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:14
	_go_fuzz_dep_.CoverTab[45859]++
											b := make([]byte, CmsgSpace(SizeofUcred))
											h := (*Cmsghdr)(unsafe.Pointer(&b[0]))
											h.Level = SOL_SOCKET
											h.Type = SCM_CREDENTIALS
											h.SetLen(CmsgLen(SizeofUcred))
											*(*Ucred)(h.data(0)) = *ucred
											return b
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:21
	// _ = "end of CoverTab[45859]"
}

// ParseUnixCredentials decodes a socket control message that contains
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:24
// credentials in a Ucred structure. To receive such a message, the
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:24
// SO_PASSCRED option must be enabled on the socket.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:27
func ParseUnixCredentials(m *SocketControlMessage) (*Ucred, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:27
	_go_fuzz_dep_.CoverTab[45860]++
											if m.Header.Level != SOL_SOCKET {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:28
		_go_fuzz_dep_.CoverTab[45863]++
												return nil, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:29
		// _ = "end of CoverTab[45863]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:30
		_go_fuzz_dep_.CoverTab[45864]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:30
		// _ = "end of CoverTab[45864]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:30
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:30
	// _ = "end of CoverTab[45860]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:30
	_go_fuzz_dep_.CoverTab[45861]++
											if m.Header.Type != SCM_CREDENTIALS {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:31
		_go_fuzz_dep_.CoverTab[45865]++
												return nil, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:32
		// _ = "end of CoverTab[45865]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:33
		_go_fuzz_dep_.CoverTab[45866]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:33
		// _ = "end of CoverTab[45866]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:33
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:33
	// _ = "end of CoverTab[45861]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:33
	_go_fuzz_dep_.CoverTab[45862]++
											ucred := *(*Ucred)(unsafe.Pointer(&m.Data[0]))
											return &ucred, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:35
	// _ = "end of CoverTab[45862]"
}

// PktInfo4 encodes Inet4Pktinfo into a socket control message of type IP_PKTINFO.
func PktInfo4(info *Inet4Pktinfo) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:39
	_go_fuzz_dep_.CoverTab[45867]++
											b := make([]byte, CmsgSpace(SizeofInet4Pktinfo))
											h := (*Cmsghdr)(unsafe.Pointer(&b[0]))
											h.Level = SOL_IP
											h.Type = IP_PKTINFO
											h.SetLen(CmsgLen(SizeofInet4Pktinfo))
											*(*Inet4Pktinfo)(h.data(0)) = *info
											return b
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:46
	// _ = "end of CoverTab[45867]"
}

// PktInfo6 encodes Inet6Pktinfo into a socket control message of type IPV6_PKTINFO.
func PktInfo6(info *Inet6Pktinfo) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:50
	_go_fuzz_dep_.CoverTab[45868]++
											b := make([]byte, CmsgSpace(SizeofInet6Pktinfo))
											h := (*Cmsghdr)(unsafe.Pointer(&b[0]))
											h.Level = SOL_IPV6
											h.Type = IPV6_PKTINFO
											h.SetLen(CmsgLen(SizeofInet6Pktinfo))
											*(*Inet6Pktinfo)(h.data(0)) = *info
											return b
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:57
	// _ = "end of CoverTab[45868]"
}

// ParseOrigDstAddr decodes a socket control message containing the original
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:60
// destination address. To receive such a message the IP_RECVORIGDSTADDR or
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:60
// IPV6_RECVORIGDSTADDR option must be enabled on the socket.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:63
func ParseOrigDstAddr(m *SocketControlMessage) (Sockaddr, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:63
	_go_fuzz_dep_.CoverTab[45869]++
											switch {
	case m.Header.Level == SOL_IP && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:65
		_go_fuzz_dep_.CoverTab[45873]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:65
		return m.Header.Type == IP_ORIGDSTADDR
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:65
		// _ = "end of CoverTab[45873]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:65
	}():
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:65
		_go_fuzz_dep_.CoverTab[45870]++
												pp := (*RawSockaddrInet4)(unsafe.Pointer(&m.Data[0]))
												sa := new(SockaddrInet4)
												p := (*[2]byte)(unsafe.Pointer(&pp.Port))
												sa.Port = int(p[0])<<8 + int(p[1])
												sa.Addr = pp.Addr
												return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:71
		// _ = "end of CoverTab[45870]"

	case m.Header.Level == SOL_IPV6 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:73
		_go_fuzz_dep_.CoverTab[45874]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:73
		return m.Header.Type == IPV6_ORIGDSTADDR
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:73
		// _ = "end of CoverTab[45874]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:73
	}():
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:73
		_go_fuzz_dep_.CoverTab[45871]++
												pp := (*RawSockaddrInet6)(unsafe.Pointer(&m.Data[0]))
												sa := new(SockaddrInet6)
												p := (*[2]byte)(unsafe.Pointer(&pp.Port))
												sa.Port = int(p[0])<<8 + int(p[1])
												sa.ZoneId = pp.Scope_id
												sa.Addr = pp.Addr
												return sa, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:80
		// _ = "end of CoverTab[45871]"

	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:82
		_go_fuzz_dep_.CoverTab[45872]++
												return nil, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:83
		// _ = "end of CoverTab[45872]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:84
	// _ = "end of CoverTab[45869]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:85
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_linux.go:85
var _ = _go_fuzz_dep_.CoverTab
