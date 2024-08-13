// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos

// Socket control messages

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:10
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:10
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:10
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:10
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:10
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:10
)

import (
	"unsafe"
)

// CmsgLen returns the value to store in the Len field of the Cmsghdr
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:16
// structure, taking into account any necessary alignment.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:18
func CmsgLen(datalen int) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:18
	_go_fuzz_dep_.CoverTab[45875]++
											return cmsgAlignOf(SizeofCmsghdr) + datalen
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:19
	// _ = "end of CoverTab[45875]"
}

// CmsgSpace returns the number of bytes an ancillary element with
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:22
// payload of the passed data length occupies.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:24
func CmsgSpace(datalen int) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:24
	_go_fuzz_dep_.CoverTab[45876]++
											return cmsgAlignOf(SizeofCmsghdr) + cmsgAlignOf(datalen)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:25
	// _ = "end of CoverTab[45876]"
}

func (h *Cmsghdr) data(offset uintptr) unsafe.Pointer {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:28
	_go_fuzz_dep_.CoverTab[45877]++
											return unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(cmsgAlignOf(SizeofCmsghdr)) + offset)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:29
	// _ = "end of CoverTab[45877]"
}

// SocketControlMessage represents a socket control message.
type SocketControlMessage struct {
	Header	Cmsghdr
	Data	[]byte
}

// ParseSocketControlMessage parses b as an array of socket control
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:38
// messages.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:40
func ParseSocketControlMessage(b []byte) ([]SocketControlMessage, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:40
	_go_fuzz_dep_.CoverTab[45878]++
											var msgs []SocketControlMessage
											i := 0
											for i+CmsgLen(0) <= len(b) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:43
		_go_fuzz_dep_.CoverTab[45880]++
												h, dbuf, err := socketControlMessageHeaderAndData(b[i:])
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:45
			_go_fuzz_dep_.CoverTab[45882]++
													return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:46
			// _ = "end of CoverTab[45882]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:47
			_go_fuzz_dep_.CoverTab[45883]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:47
			// _ = "end of CoverTab[45883]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:47
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:47
		// _ = "end of CoverTab[45880]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:47
		_go_fuzz_dep_.CoverTab[45881]++
												m := SocketControlMessage{Header: *h, Data: dbuf}
												msgs = append(msgs, m)
												i += cmsgAlignOf(int(h.Len))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:50
		// _ = "end of CoverTab[45881]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:51
	// _ = "end of CoverTab[45878]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:51
	_go_fuzz_dep_.CoverTab[45879]++
											return msgs, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:52
	// _ = "end of CoverTab[45879]"
}

// ParseOneSocketControlMessage parses a single socket control message from b, returning the message header,
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:55
// message data (a slice of b), and the remainder of b after that single message.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:55
// When there are no remaining messages, len(remainder) == 0.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:58
func ParseOneSocketControlMessage(b []byte) (hdr Cmsghdr, data []byte, remainder []byte, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:58
	_go_fuzz_dep_.CoverTab[45884]++
											h, dbuf, err := socketControlMessageHeaderAndData(b)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:60
		_go_fuzz_dep_.CoverTab[45887]++
												return Cmsghdr{}, nil, nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:61
		// _ = "end of CoverTab[45887]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:62
		_go_fuzz_dep_.CoverTab[45888]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:62
		// _ = "end of CoverTab[45888]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:62
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:62
	// _ = "end of CoverTab[45884]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:62
	_go_fuzz_dep_.CoverTab[45885]++
											if i := cmsgAlignOf(int(h.Len)); i < len(b) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:63
		_go_fuzz_dep_.CoverTab[45889]++
												remainder = b[i:]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:64
		// _ = "end of CoverTab[45889]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:65
		_go_fuzz_dep_.CoverTab[45890]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:65
		// _ = "end of CoverTab[45890]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:65
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:65
	// _ = "end of CoverTab[45885]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:65
	_go_fuzz_dep_.CoverTab[45886]++
											return *h, dbuf, remainder, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:66
	// _ = "end of CoverTab[45886]"
}

func socketControlMessageHeaderAndData(b []byte) (*Cmsghdr, []byte, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:69
	_go_fuzz_dep_.CoverTab[45891]++
											h := (*Cmsghdr)(unsafe.Pointer(&b[0]))
											if h.Len < SizeofCmsghdr || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:71
		_go_fuzz_dep_.CoverTab[45893]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:71
		return uint64(h.Len) > uint64(len(b))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:71
		// _ = "end of CoverTab[45893]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:71
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:71
		_go_fuzz_dep_.CoverTab[45894]++
												return nil, nil, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:72
		// _ = "end of CoverTab[45894]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:73
		_go_fuzz_dep_.CoverTab[45895]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:73
		// _ = "end of CoverTab[45895]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:73
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:73
	// _ = "end of CoverTab[45891]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:73
	_go_fuzz_dep_.CoverTab[45892]++
											return h, b[cmsgAlignOf(SizeofCmsghdr):h.Len], nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:74
	// _ = "end of CoverTab[45892]"
}

// UnixRights encodes a set of open file descriptors into a socket
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:77
// control message for sending to another process.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:79
func UnixRights(fds ...int) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:79
	_go_fuzz_dep_.CoverTab[45896]++
											datalen := len(fds) * 4
											b := make([]byte, CmsgSpace(datalen))
											h := (*Cmsghdr)(unsafe.Pointer(&b[0]))
											h.Level = SOL_SOCKET
											h.Type = SCM_RIGHTS
											h.SetLen(CmsgLen(datalen))
											for i, fd := range fds {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:86
		_go_fuzz_dep_.CoverTab[45898]++
												*(*int32)(h.data(4 * uintptr(i))) = int32(fd)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:87
		// _ = "end of CoverTab[45898]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:88
	// _ = "end of CoverTab[45896]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:88
	_go_fuzz_dep_.CoverTab[45897]++
											return b
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:89
	// _ = "end of CoverTab[45897]"
}

// ParseUnixRights decodes a socket control message that contains an
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:92
// integer array of open file descriptors from another process.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:94
func ParseUnixRights(m *SocketControlMessage) ([]int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:94
	_go_fuzz_dep_.CoverTab[45899]++
											if m.Header.Level != SOL_SOCKET {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:95
		_go_fuzz_dep_.CoverTab[45903]++
												return nil, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:96
		// _ = "end of CoverTab[45903]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:97
		_go_fuzz_dep_.CoverTab[45904]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:97
		// _ = "end of CoverTab[45904]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:97
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:97
	// _ = "end of CoverTab[45899]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:97
	_go_fuzz_dep_.CoverTab[45900]++
											if m.Header.Type != SCM_RIGHTS {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:98
		_go_fuzz_dep_.CoverTab[45905]++
												return nil, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:99
		// _ = "end of CoverTab[45905]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:100
		_go_fuzz_dep_.CoverTab[45906]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:100
		// _ = "end of CoverTab[45906]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:100
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:100
	// _ = "end of CoverTab[45900]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:100
	_go_fuzz_dep_.CoverTab[45901]++
											fds := make([]int, len(m.Data)>>2)
											for i, j := 0, 0; i < len(m.Data); i += 4 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:102
		_go_fuzz_dep_.CoverTab[45907]++
												fds[j] = int(*(*int32)(unsafe.Pointer(&m.Data[i])))
												j++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:104
		// _ = "end of CoverTab[45907]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:105
	// _ = "end of CoverTab[45901]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:105
	_go_fuzz_dep_.CoverTab[45902]++
											return fds, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:106
	// _ = "end of CoverTab[45902]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:107
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/sockcmsg_unix.go:107
var _ = _go_fuzz_dep_.CoverTab
