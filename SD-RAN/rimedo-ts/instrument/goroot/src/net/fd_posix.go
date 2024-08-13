// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || windows

//line /usr/local/go/src/net/fd_posix.go:7
package net

//line /usr/local/go/src/net/fd_posix.go:7
import (
//line /usr/local/go/src/net/fd_posix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/fd_posix.go:7
)
//line /usr/local/go/src/net/fd_posix.go:7
import (
//line /usr/local/go/src/net/fd_posix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/fd_posix.go:7
)

import (
	"internal/poll"
	"runtime"
	"syscall"
	"time"
)

// Network file descriptor.
type netFD struct {
	pfd	poll.FD

	// immutable until Close
	family		int
	sotype		int
	isConnected	bool	// handshake completed or use of association with peer
	net		string
	laddr		Addr
	raddr		Addr
}

func (fd *netFD) setAddr(laddr, raddr Addr) {
//line /usr/local/go/src/net/fd_posix.go:29
	_go_fuzz_dep_.CoverTab[13827]++
						fd.laddr = laddr
						fd.raddr = raddr
						runtime.SetFinalizer(fd, (*netFD).Close)
//line /usr/local/go/src/net/fd_posix.go:32
	// _ = "end of CoverTab[13827]"
}

func (fd *netFD) Close() error {
//line /usr/local/go/src/net/fd_posix.go:35
	_go_fuzz_dep_.CoverTab[13828]++
						runtime.SetFinalizer(fd, nil)
						return fd.pfd.Close()
//line /usr/local/go/src/net/fd_posix.go:37
	// _ = "end of CoverTab[13828]"
}

func (fd *netFD) shutdown(how int) error {
//line /usr/local/go/src/net/fd_posix.go:40
	_go_fuzz_dep_.CoverTab[13829]++
						err := fd.pfd.Shutdown(how)
						runtime.KeepAlive(fd)
						return wrapSyscallError("shutdown", err)
//line /usr/local/go/src/net/fd_posix.go:43
	// _ = "end of CoverTab[13829]"
}

func (fd *netFD) closeRead() error {
//line /usr/local/go/src/net/fd_posix.go:46
	_go_fuzz_dep_.CoverTab[13830]++
						return fd.shutdown(syscall.SHUT_RD)
//line /usr/local/go/src/net/fd_posix.go:47
	// _ = "end of CoverTab[13830]"
}

func (fd *netFD) closeWrite() error {
//line /usr/local/go/src/net/fd_posix.go:50
	_go_fuzz_dep_.CoverTab[13831]++
						return fd.shutdown(syscall.SHUT_WR)
//line /usr/local/go/src/net/fd_posix.go:51
	// _ = "end of CoverTab[13831]"
}

func (fd *netFD) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/fd_posix.go:54
	_go_fuzz_dep_.CoverTab[13832]++
						n, err = fd.pfd.Read(p)
						runtime.KeepAlive(fd)
						return n, wrapSyscallError(readSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:57
	// _ = "end of CoverTab[13832]"
}

func (fd *netFD) readFrom(p []byte) (n int, sa syscall.Sockaddr, err error) {
//line /usr/local/go/src/net/fd_posix.go:60
	_go_fuzz_dep_.CoverTab[13833]++
						n, sa, err = fd.pfd.ReadFrom(p)
						runtime.KeepAlive(fd)
						return n, sa, wrapSyscallError(readFromSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:63
	// _ = "end of CoverTab[13833]"
}
func (fd *netFD) readFromInet4(p []byte, from *syscall.SockaddrInet4) (n int, err error) {
//line /usr/local/go/src/net/fd_posix.go:65
	_go_fuzz_dep_.CoverTab[13834]++
						n, err = fd.pfd.ReadFromInet4(p, from)
						runtime.KeepAlive(fd)
						return n, wrapSyscallError(readFromSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:68
	// _ = "end of CoverTab[13834]"
}

func (fd *netFD) readFromInet6(p []byte, from *syscall.SockaddrInet6) (n int, err error) {
//line /usr/local/go/src/net/fd_posix.go:71
	_go_fuzz_dep_.CoverTab[13835]++
						n, err = fd.pfd.ReadFromInet6(p, from)
						runtime.KeepAlive(fd)
						return n, wrapSyscallError(readFromSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:74
	// _ = "end of CoverTab[13835]"
}

func (fd *netFD) readMsg(p []byte, oob []byte, flags int) (n, oobn, retflags int, sa syscall.Sockaddr, err error) {
//line /usr/local/go/src/net/fd_posix.go:77
	_go_fuzz_dep_.CoverTab[13836]++
						n, oobn, retflags, sa, err = fd.pfd.ReadMsg(p, oob, flags)
						runtime.KeepAlive(fd)
						return n, oobn, retflags, sa, wrapSyscallError(readMsgSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:80
	// _ = "end of CoverTab[13836]"
}

func (fd *netFD) readMsgInet4(p []byte, oob []byte, flags int, sa *syscall.SockaddrInet4) (n, oobn, retflags int, err error) {
//line /usr/local/go/src/net/fd_posix.go:83
	_go_fuzz_dep_.CoverTab[13837]++
						n, oobn, retflags, err = fd.pfd.ReadMsgInet4(p, oob, flags, sa)
						runtime.KeepAlive(fd)
						return n, oobn, retflags, wrapSyscallError(readMsgSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:86
	// _ = "end of CoverTab[13837]"
}

func (fd *netFD) readMsgInet6(p []byte, oob []byte, flags int, sa *syscall.SockaddrInet6) (n, oobn, retflags int, err error) {
//line /usr/local/go/src/net/fd_posix.go:89
	_go_fuzz_dep_.CoverTab[13838]++
						n, oobn, retflags, err = fd.pfd.ReadMsgInet6(p, oob, flags, sa)
						runtime.KeepAlive(fd)
						return n, oobn, retflags, wrapSyscallError(readMsgSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:92
	// _ = "end of CoverTab[13838]"
}

func (fd *netFD) Write(p []byte) (nn int, err error) {
//line /usr/local/go/src/net/fd_posix.go:95
	_go_fuzz_dep_.CoverTab[13839]++
						nn, err = fd.pfd.Write(p)
						runtime.KeepAlive(fd)
						return nn, wrapSyscallError(writeSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:98
	// _ = "end of CoverTab[13839]"
}

func (fd *netFD) writeTo(p []byte, sa syscall.Sockaddr) (n int, err error) {
//line /usr/local/go/src/net/fd_posix.go:101
	_go_fuzz_dep_.CoverTab[13840]++
						n, err = fd.pfd.WriteTo(p, sa)
						runtime.KeepAlive(fd)
						return n, wrapSyscallError(writeToSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:104
	// _ = "end of CoverTab[13840]"
}

func (fd *netFD) writeToInet4(p []byte, sa *syscall.SockaddrInet4) (n int, err error) {
//line /usr/local/go/src/net/fd_posix.go:107
	_go_fuzz_dep_.CoverTab[13841]++
						n, err = fd.pfd.WriteToInet4(p, sa)
						runtime.KeepAlive(fd)
						return n, wrapSyscallError(writeToSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:110
	// _ = "end of CoverTab[13841]"
}

func (fd *netFD) writeToInet6(p []byte, sa *syscall.SockaddrInet6) (n int, err error) {
//line /usr/local/go/src/net/fd_posix.go:113
	_go_fuzz_dep_.CoverTab[13842]++
						n, err = fd.pfd.WriteToInet6(p, sa)
						runtime.KeepAlive(fd)
						return n, wrapSyscallError(writeToSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:116
	// _ = "end of CoverTab[13842]"
}

func (fd *netFD) writeMsg(p []byte, oob []byte, sa syscall.Sockaddr) (n int, oobn int, err error) {
//line /usr/local/go/src/net/fd_posix.go:119
	_go_fuzz_dep_.CoverTab[13843]++
						n, oobn, err = fd.pfd.WriteMsg(p, oob, sa)
						runtime.KeepAlive(fd)
						return n, oobn, wrapSyscallError(writeMsgSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:122
	// _ = "end of CoverTab[13843]"
}

func (fd *netFD) writeMsgInet4(p []byte, oob []byte, sa *syscall.SockaddrInet4) (n int, oobn int, err error) {
//line /usr/local/go/src/net/fd_posix.go:125
	_go_fuzz_dep_.CoverTab[13844]++
						n, oobn, err = fd.pfd.WriteMsgInet4(p, oob, sa)
						runtime.KeepAlive(fd)
						return n, oobn, wrapSyscallError(writeMsgSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:128
	// _ = "end of CoverTab[13844]"
}

func (fd *netFD) writeMsgInet6(p []byte, oob []byte, sa *syscall.SockaddrInet6) (n int, oobn int, err error) {
//line /usr/local/go/src/net/fd_posix.go:131
	_go_fuzz_dep_.CoverTab[13845]++
						n, oobn, err = fd.pfd.WriteMsgInet6(p, oob, sa)
						runtime.KeepAlive(fd)
						return n, oobn, wrapSyscallError(writeMsgSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:134
	// _ = "end of CoverTab[13845]"
}

func (fd *netFD) SetDeadline(t time.Time) error {
//line /usr/local/go/src/net/fd_posix.go:137
	_go_fuzz_dep_.CoverTab[13846]++
						return fd.pfd.SetDeadline(t)
//line /usr/local/go/src/net/fd_posix.go:138
	// _ = "end of CoverTab[13846]"
}

func (fd *netFD) SetReadDeadline(t time.Time) error {
//line /usr/local/go/src/net/fd_posix.go:141
	_go_fuzz_dep_.CoverTab[13847]++
						return fd.pfd.SetReadDeadline(t)
//line /usr/local/go/src/net/fd_posix.go:142
	// _ = "end of CoverTab[13847]"
}

func (fd *netFD) SetWriteDeadline(t time.Time) error {
//line /usr/local/go/src/net/fd_posix.go:145
	_go_fuzz_dep_.CoverTab[13848]++
						return fd.pfd.SetWriteDeadline(t)
//line /usr/local/go/src/net/fd_posix.go:146
	// _ = "end of CoverTab[13848]"
}

//line /usr/local/go/src/net/fd_posix.go:147
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/fd_posix.go:147
var _ = _go_fuzz_dep_.CoverTab
