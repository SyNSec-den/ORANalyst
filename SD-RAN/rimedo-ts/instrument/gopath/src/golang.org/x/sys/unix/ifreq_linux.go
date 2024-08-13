// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux
// +build linux

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:8
)

import (
	"unsafe"
)

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:17
// An Ifreq is a type-safe wrapper around the raw ifreq struct. An Ifreq
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:17
// contains an interface name and a union of arbitrary data which can be
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:17
// accessed using the Ifreq's methods. To create an Ifreq, use the NewIfreq
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:17
// function.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:17
//
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:17
// Use the Name method to access the stored interface name. The union data
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:17
// fields can be get and set using the following methods:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:17
//   - Uint16/SetUint16: flags
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:17
//   - Uint32/SetUint32: ifindex, metric, mtu
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:26
type Ifreq struct{ raw ifreq }

// NewIfreq creates an Ifreq with the input network interface name after
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:28
// validating the name does not exceed IFNAMSIZ-1 (trailing NULL required)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:28
// bytes.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:31
func NewIfreq(name string) (*Ifreq, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:31
	_go_fuzz_dep_.CoverTab[45786]++

											if len(name) >= IFNAMSIZ {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:33
		_go_fuzz_dep_.CoverTab[45788]++
												return nil, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:34
		// _ = "end of CoverTab[45788]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:35
		_go_fuzz_dep_.CoverTab[45789]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:35
		// _ = "end of CoverTab[45789]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:35
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:35
	// _ = "end of CoverTab[45786]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:35
	_go_fuzz_dep_.CoverTab[45787]++

											var ifr ifreq
											copy(ifr.Ifrn[:], name)

											return &Ifreq{raw: ifr}, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:40
	// _ = "end of CoverTab[45787]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:45
// Name returns the interface name associated with the Ifreq.
func (ifr *Ifreq) Name() string {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:46
	_go_fuzz_dep_.CoverTab[45790]++
											return ByteSliceToString(ifr.raw.Ifrn[:])
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:47
	// _ = "end of CoverTab[45790]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:54
// Inet4Addr returns the Ifreq union data from an embedded sockaddr as a C
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:54
// in_addr/Go []byte (4-byte IPv4 address) value. If the sockaddr family is not
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:54
// AF_INET, an error is returned.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:57
func (ifr *Ifreq) Inet4Addr() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:57
	_go_fuzz_dep_.CoverTab[45791]++
											raw := *(*RawSockaddrInet4)(unsafe.Pointer(&ifr.raw.Ifru[:SizeofSockaddrInet4][0]))
											if raw.Family != AF_INET {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:59
		_go_fuzz_dep_.CoverTab[45793]++

												return nil, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:61
		// _ = "end of CoverTab[45793]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:62
		_go_fuzz_dep_.CoverTab[45794]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:62
		// _ = "end of CoverTab[45794]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:62
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:62
	// _ = "end of CoverTab[45791]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:62
	_go_fuzz_dep_.CoverTab[45792]++

											return raw.Addr[:], nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:64
	// _ = "end of CoverTab[45792]"
}

// SetInet4Addr sets a C in_addr/Go []byte (4-byte IPv4 address) value in an
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:67
// embedded sockaddr within the Ifreq's union data. v must be 4 bytes in length
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:67
// or an error will be returned.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:70
func (ifr *Ifreq) SetInet4Addr(v []byte) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:70
	_go_fuzz_dep_.CoverTab[45795]++
											if len(v) != 4 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:71
		_go_fuzz_dep_.CoverTab[45797]++
												return EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:72
		// _ = "end of CoverTab[45797]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:73
		_go_fuzz_dep_.CoverTab[45798]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:73
		// _ = "end of CoverTab[45798]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:73
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:73
	// _ = "end of CoverTab[45795]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:73
	_go_fuzz_dep_.CoverTab[45796]++

											var addr [4]byte
											copy(addr[:], v)

											ifr.clear()
											*(*RawSockaddrInet4)(
		unsafe.Pointer(&ifr.raw.Ifru[:SizeofSockaddrInet4][0]),
	) = RawSockaddrInet4{

		Family:	AF_INET,
		Addr:	addr,
	}

											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:87
	// _ = "end of CoverTab[45796]"
}

// Uint16 returns the Ifreq union data as a C short/Go uint16 value.
func (ifr *Ifreq) Uint16() uint16 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:91
	_go_fuzz_dep_.CoverTab[45799]++
											return *(*uint16)(unsafe.Pointer(&ifr.raw.Ifru[:2][0]))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:92
	// _ = "end of CoverTab[45799]"
}

// SetUint16 sets a C short/Go uint16 value as the Ifreq's union data.
func (ifr *Ifreq) SetUint16(v uint16) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:96
	_go_fuzz_dep_.CoverTab[45800]++
											ifr.clear()
											*(*uint16)(unsafe.Pointer(&ifr.raw.Ifru[:2][0])) = v
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:98
	// _ = "end of CoverTab[45800]"
}

// Uint32 returns the Ifreq union data as a C int/Go uint32 value.
func (ifr *Ifreq) Uint32() uint32 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:102
	_go_fuzz_dep_.CoverTab[45801]++
											return *(*uint32)(unsafe.Pointer(&ifr.raw.Ifru[:4][0]))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:103
	// _ = "end of CoverTab[45801]"
}

// SetUint32 sets a C int/Go uint32 value as the Ifreq's union data.
func (ifr *Ifreq) SetUint32(v uint32) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:107
	_go_fuzz_dep_.CoverTab[45802]++
											ifr.clear()
											*(*uint32)(unsafe.Pointer(&ifr.raw.Ifru[:4][0])) = v
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:109
	// _ = "end of CoverTab[45802]"
}

// clear zeroes the ifreq's union field to prevent trailing garbage data from
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:112
// being sent to the kernel if an ifreq is reused.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:114
func (ifr *Ifreq) clear() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:114
	_go_fuzz_dep_.CoverTab[45803]++
											for i := range ifr.raw.Ifru {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:115
		_go_fuzz_dep_.CoverTab[45804]++
												ifr.raw.Ifru[i] = 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:116
		// _ = "end of CoverTab[45804]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:117
	// _ = "end of CoverTab[45803]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:123
// An ifreqData is an Ifreq which carries pointer data. To produce an ifreqData,
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:123
// use the Ifreq.withData method.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:125
type ifreqData struct {
	name	[IFNAMSIZ]byte
	// A type separate from ifreq is required in order to comply with the
	// unsafe.Pointer rules since the "pointer-ness" of data would not be
	// preserved if it were cast into the byte array of a raw ifreq.
	data	unsafe.Pointer
	// Pad to the same size as ifreq.
	_	[len(ifreq{}.Ifru) - SizeofPtr]byte
}

// withData produces an ifreqData with the pointer p set for ioctls which require
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:135
// arbitrary pointer data.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:137
func (ifr Ifreq) withData(p unsafe.Pointer) ifreqData {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:137
	_go_fuzz_dep_.CoverTab[45805]++
											return ifreqData{
		name:	ifr.raw.Ifrn,
		data:	p,
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:141
	// _ = "end of CoverTab[45805]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:142
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ifreq_linux.go:142
var _ = _go_fuzz_dep_.CoverTab
