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
	_go_fuzz_dep_.CoverTab[5437]++
						fd.laddr = laddr
						fd.raddr = raddr
						runtime.SetFinalizer(fd, (*netFD).Close)
//line /usr/local/go/src/net/fd_posix.go:32
	// _ = "end of CoverTab[5437]"
}

func (fd *netFD) Close() error {
//line /usr/local/go/src/net/fd_posix.go:35
	_go_fuzz_dep_.CoverTab[5438]++
						runtime.SetFinalizer(fd, nil)
						return fd.pfd.Close()
//line /usr/local/go/src/net/fd_posix.go:37
	// _ = "end of CoverTab[5438]"
}

func (fd *netFD) shutdown(how int) error {
//line /usr/local/go/src/net/fd_posix.go:40
	_go_fuzz_dep_.CoverTab[5439]++
						err := fd.pfd.Shutdown(how)
						runtime.KeepAlive(fd)
						return wrapSyscallError("shutdown", err)
//line /usr/local/go/src/net/fd_posix.go:43
	// _ = "end of CoverTab[5439]"
}

func (fd *netFD) closeRead() error {
//line /usr/local/go/src/net/fd_posix.go:46
	_go_fuzz_dep_.CoverTab[5440]++
						return fd.shutdown(syscall.SHUT_RD)
//line /usr/local/go/src/net/fd_posix.go:47
	// _ = "end of CoverTab[5440]"
}

func (fd *netFD) closeWrite() error {
//line /usr/local/go/src/net/fd_posix.go:50
	_go_fuzz_dep_.CoverTab[5441]++
						return fd.shutdown(syscall.SHUT_WR)
//line /usr/local/go/src/net/fd_posix.go:51
	// _ = "end of CoverTab[5441]"
}

func (fd *netFD) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/net/fd_posix.go:54
	_go_fuzz_dep_.CoverTab[5442]++
						n, err = fd.pfd.Read(p)
						runtime.KeepAlive(fd)
						return n, wrapSyscallError(readSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:57
	// _ = "end of CoverTab[5442]"
}

func (fd *netFD) readFrom(p []byte) (n int, sa syscall.Sockaddr, err error) {
//line /usr/local/go/src/net/fd_posix.go:60
	_go_fuzz_dep_.CoverTab[5443]++
						n, sa, err = fd.pfd.ReadFrom(p)
						runtime.KeepAlive(fd)
						return n, sa, wrapSyscallError(readFromSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:63
	// _ = "end of CoverTab[5443]"
}
func (fd *netFD) readFromInet4(p []byte, from *syscall.SockaddrInet4) (n int, err error) {
//line /usr/local/go/src/net/fd_posix.go:65
	_go_fuzz_dep_.CoverTab[5444]++
						n, err = fd.pfd.ReadFromInet4(p, from)
						runtime.KeepAlive(fd)
						return n, wrapSyscallError(readFromSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:68
	// _ = "end of CoverTab[5444]"
}

func (fd *netFD) readFromInet6(p []byte, from *syscall.SockaddrInet6) (n int, err error) {
//line /usr/local/go/src/net/fd_posix.go:71
	_go_fuzz_dep_.CoverTab[5445]++
						n, err = fd.pfd.ReadFromInet6(p, from)
						runtime.KeepAlive(fd)
						return n, wrapSyscallError(readFromSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:74
	// _ = "end of CoverTab[5445]"
}

func (fd *netFD) readMsg(p []byte, oob []byte, flags int) (n, oobn, retflags int, sa syscall.Sockaddr, err error) {
//line /usr/local/go/src/net/fd_posix.go:77
	_go_fuzz_dep_.CoverTab[5446]++
						n, oobn, retflags, sa, err = fd.pfd.ReadMsg(p, oob, flags)
						runtime.KeepAlive(fd)
						return n, oobn, retflags, sa, wrapSyscallError(readMsgSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:80
	// _ = "end of CoverTab[5446]"
}

func (fd *netFD) readMsgInet4(p []byte, oob []byte, flags int, sa *syscall.SockaddrInet4) (n, oobn, retflags int, err error) {
//line /usr/local/go/src/net/fd_posix.go:83
	_go_fuzz_dep_.CoverTab[5447]++
						n, oobn, retflags, err = fd.pfd.ReadMsgInet4(p, oob, flags, sa)
						runtime.KeepAlive(fd)
						return n, oobn, retflags, wrapSyscallError(readMsgSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:86
	// _ = "end of CoverTab[5447]"
}

func (fd *netFD) readMsgInet6(p []byte, oob []byte, flags int, sa *syscall.SockaddrInet6) (n, oobn, retflags int, err error) {
//line /usr/local/go/src/net/fd_posix.go:89
	_go_fuzz_dep_.CoverTab[5448]++
						n, oobn, retflags, err = fd.pfd.ReadMsgInet6(p, oob, flags, sa)
						runtime.KeepAlive(fd)
						return n, oobn, retflags, wrapSyscallError(readMsgSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:92
	// _ = "end of CoverTab[5448]"
}

func (fd *netFD) Write(p []byte) (nn int, err error) {
//line /usr/local/go/src/net/fd_posix.go:95
	_go_fuzz_dep_.CoverTab[5449]++
						nn, err = fd.pfd.Write(p)
						runtime.KeepAlive(fd)
						return nn, wrapSyscallError(writeSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:98
	// _ = "end of CoverTab[5449]"
}

func (fd *netFD) writeTo(p []byte, sa syscall.Sockaddr) (n int, err error) {
//line /usr/local/go/src/net/fd_posix.go:101
	_go_fuzz_dep_.CoverTab[5450]++
						n, err = fd.pfd.WriteTo(p, sa)
						runtime.KeepAlive(fd)
						return n, wrapSyscallError(writeToSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:104
	// _ = "end of CoverTab[5450]"
}

func (fd *netFD) writeToInet4(p []byte, sa *syscall.SockaddrInet4) (n int, err error) {
//line /usr/local/go/src/net/fd_posix.go:107
	_go_fuzz_dep_.CoverTab[5451]++
						n, err = fd.pfd.WriteToInet4(p, sa)
						runtime.KeepAlive(fd)
						return n, wrapSyscallError(writeToSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:110
	// _ = "end of CoverTab[5451]"
}

func (fd *netFD) writeToInet6(p []byte, sa *syscall.SockaddrInet6) (n int, err error) {
//line /usr/local/go/src/net/fd_posix.go:113
	_go_fuzz_dep_.CoverTab[5452]++
						n, err = fd.pfd.WriteToInet6(p, sa)
						runtime.KeepAlive(fd)
						return n, wrapSyscallError(writeToSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:116
	// _ = "end of CoverTab[5452]"
}

func (fd *netFD) writeMsg(p []byte, oob []byte, sa syscall.Sockaddr) (n int, oobn int, err error) {
//line /usr/local/go/src/net/fd_posix.go:119
	_go_fuzz_dep_.CoverTab[5453]++
						n, oobn, err = fd.pfd.WriteMsg(p, oob, sa)
						runtime.KeepAlive(fd)
						return n, oobn, wrapSyscallError(writeMsgSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:122
	// _ = "end of CoverTab[5453]"
}

func (fd *netFD) writeMsgInet4(p []byte, oob []byte, sa *syscall.SockaddrInet4) (n int, oobn int, err error) {
//line /usr/local/go/src/net/fd_posix.go:125
	_go_fuzz_dep_.CoverTab[5454]++
						n, oobn, err = fd.pfd.WriteMsgInet4(p, oob, sa)
						runtime.KeepAlive(fd)
						return n, oobn, wrapSyscallError(writeMsgSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:128
	// _ = "end of CoverTab[5454]"
}

func (fd *netFD) writeMsgInet6(p []byte, oob []byte, sa *syscall.SockaddrInet6) (n int, oobn int, err error) {
//line /usr/local/go/src/net/fd_posix.go:131
	_go_fuzz_dep_.CoverTab[5455]++
						n, oobn, err = fd.pfd.WriteMsgInet6(p, oob, sa)
						runtime.KeepAlive(fd)
						return n, oobn, wrapSyscallError(writeMsgSyscallName, err)
//line /usr/local/go/src/net/fd_posix.go:134
	// _ = "end of CoverTab[5455]"
}

func (fd *netFD) SetDeadline(t time.Time) error {
//line /usr/local/go/src/net/fd_posix.go:137
	_go_fuzz_dep_.CoverTab[5456]++
						return fd.pfd.SetDeadline(t)
//line /usr/local/go/src/net/fd_posix.go:138
	// _ = "end of CoverTab[5456]"
}

func (fd *netFD) SetReadDeadline(t time.Time) error {
//line /usr/local/go/src/net/fd_posix.go:141
	_go_fuzz_dep_.CoverTab[5457]++
						return fd.pfd.SetReadDeadline(t)
//line /usr/local/go/src/net/fd_posix.go:142
	// _ = "end of CoverTab[5457]"
}

func (fd *netFD) SetWriteDeadline(t time.Time) error {
//line /usr/local/go/src/net/fd_posix.go:145
	_go_fuzz_dep_.CoverTab[5458]++
						return fd.pfd.SetWriteDeadline(t)
//line /usr/local/go/src/net/fd_posix.go:146
	// _ = "end of CoverTab[5458]"
}

//line /usr/local/go/src/net/fd_posix.go:147
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/fd_posix.go:147
var _ = _go_fuzz_dep_.CoverTab
